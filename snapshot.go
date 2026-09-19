package pgmem

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shibukawa/pgmem/internal/engine"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// SnapshotOptions configures Server.Snapshot.
type SnapshotOptions struct {
	// MaxForks caps the number of forks alive at once; Fork blocks until a
	// fork is closed when the cap is reached. Each fork is a full backend
	// (its own buffer cache and a copy-on-write view of the data
	// directory), so this bounds memory. 0 derives the cap from memory: a
	// quarter of the process's memory limit (GOMEMLIMIT, the cgroup limit
	// or the physical memory, whichever is smallest) divided by the cost
	// of one fork (shared_buffers plus about 32 MB), which is 28 on a 7 GB
	// CI runner with the default shared_buffers. A slot costs nothing until
	// a fork occupies it, so this is a ceiling, not a reservation. When the
	// memory cannot be determined the cap is runtime.GOMAXPROCS(0).
	MaxForks int
}

// Snapshot is a frozen copy of a server's state that Fork starts new
// servers from. Take one after loading the schema and seed data, then give
// every test its own fork.
type Snapshot struct {
	e      *engine.Engine
	opts   Options
	fs     *vfs.FS
	slots  chan struct{}
	closed atomic.Bool
	forks  sync.WaitGroup
}

// Snapshot checkpoints the server and clones its data directory (file
// contents are shared copy-on-write, so this is cheap). The
// server keeps running and later writes to it do not affect the snapshot.
func (s *Server) Snapshot(ctx context.Context, opts SnapshotOptions) (*Snapshot, error) {
	if s.closed.Load() {
		return nil, errors.New("pgmem: server is closed")
	}
	if opts.MaxForks <= 0 {
		opts.MaxForks = defaultMaxForks(s.opts.Params, memoryLimit(), runtime.GOMAXPROCS(0))
	}
	fsOpts := s.opts
	fsOpts.Port = 0
	if s.cl != nil {
		// A postmaster's data directory is only consistent when it is shut
		// down: stop the cluster (a fast shutdown checkpoints), copy the
		// directory and start it again. Client sessions get a new backend
		// on their next message (see cluster.go).
		fs, err := s.cloneStopped(ctx)
		if err != nil {
			return nil, err
		}
		return &Snapshot{e: s.e, opts: fsOpts, fs: fs, slots: make(chan struct{}, opts.MaxForks)}, nil
	}
	// acquire waits for any open transaction to end, so the copy is taken
	// between transactions; CHECKPOINT then flushes every dirty page. A
	// connection left idle in a transaction would block this forever, so
	// ctx bounds the wait.
	if err := s.acquire(ctx, 0, false); err != nil {
		return nil, fmt.Errorf("pgmem: snapshot waited for an open transaction to end (commit or close every connection first): %w", err)
	}
	defer s.release()
	out, err := s.b.Exec(simpleQuery("CHECKPOINT"))
	if err != nil {
		return nil, fmt.Errorf("checkpoint: %w", err)
	}
	if err := checkNoError(out); err != nil {
		return nil, fmt.Errorf("checkpoint: %w", err)
	}
	return &Snapshot{
		e:     s.e,
		opts:  fsOpts,
		fs:    s.fs.Clone(),
		slots: make(chan struct{}, opts.MaxForks),
	}, nil
}

// forkWarnAfter is how long Fork waits for a free slot before logging.
const forkWarnAfter = 5 * time.Second

// Fork starts a new server on a copy of the snapshot. It blocks while
// MaxForks forks are alive; ctx cancels the wait. Close the returned
// server to free its slot.
func (sn *Snapshot) Fork(ctx context.Context) (*Server, error) {
	if sn.closed.Load() {
		return nil, errors.New("pgmem: snapshot is closed")
	}
	select {
	case sn.slots <- struct{}{}:
	default:
		t := time.NewTimer(forkWarnAfter)
		select {
		case sn.slots <- struct{}{}:
			t.Stop()
		case <-t.C:
			sn.logf("pgmem: Fork has waited %s for a free slot (%d forks alive; close forks or raise MaxForks)", forkWarnAfter, cap(sn.slots))
			select {
			case sn.slots <- struct{}{}:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		case <-ctx.Done():
			t.Stop()
			return nil, ctx.Err()
		}
	}
	sn.forks.Add(1)
	srv, err := boot(sn.e, sn.opts, sn.fs.Clone(), sn, func() {
		<-sn.slots
		sn.forks.Done()
	})
	if err != nil {
		<-sn.slots
		sn.forks.Done()
		return nil, err
	}
	return srv, nil
}

// Close releases the snapshot. Forks already running keep working; they
// own their own copies.
func (sn *Snapshot) Close() error {
	sn.closed.Store(true)
	return nil
}

// Wait blocks until every fork has been closed.
func (sn *Snapshot) Wait() { sn.forks.Wait() }

// MaxForks is the cap on forks alive at once: SnapshotOptions.MaxForks, or
// the memory-derived default when it was 0.
func (sn *Snapshot) MaxForks() int { return cap(sn.slots) }

func (sn *Snapshot) logf(format string, args ...any) {
	if sn.opts.Log != nil {
		sn.opts.Log(format, args...)
	}
}

// cloneStopped stops the cluster, copies its data directory and starts the
// cluster again on the original. Runs under bmu.
func (s *Server) cloneStopped(ctx context.Context) (*vfs.FS, error) {
	if err := s.waitSessionsIdle(ctx); err != nil {
		return nil, fmt.Errorf("pgmem: snapshot waited for an open transaction to end (commit or close every connection first): %w", err)
	}
	s.bmu.Lock()
	defer s.bmu.Unlock()
	if s.closed.Load() {
		return nil, errServerClosed
	}
	s.muteSessions()
	if err := s.cl.Shutdown(ctx); err != nil {
		s.logf("pgmem: shutdown: %v", err)
	}
	fs := s.fs.Clone()
	cl, err := s.startCluster(s.fs)
	if err != nil {
		return nil, fmt.Errorf("pgmem: restart after snapshot: %w", err)
	}
	s.cl = cl
	s.reattachSessions()
	return fs, nil
}
