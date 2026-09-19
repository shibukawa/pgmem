package engine

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/shibukawa/pgmem/internal/guest"
	"github.com/shibukawa/pgmem/internal/host"
	"github.com/shibukawa/pgmem/internal/vfs"
)

// Shared memory window in every process's linear memory: the heap grows
// below shmBase, the System V segments (the main segment, dynamic shared
// memory for parallel workers and the stats system) are placed from
// shmBase up. 512 MiB of segments and 1.5 GiB of private heap per process.
const (
	shmBase  = 0x6000_0000
	shmLimit = 0x7FFF_0000
)

// Cluster is a running postmaster with its child processes: PostgreSQL's
// own multi-process model, each process a module instance on its own
// goroutine, sharing memory through host.Cluster.
type Cluster struct {
	e    *Engine
	fs   *vfs.FS
	c    *host.Cluster
	pm   *host.Process
	port int

	mu  sync.Mutex
	log bytes.Buffer
	// Stderr receives the server log (nil = keep in Log()).
	Stderr func([]byte)
}

// clusterPort is the port number inside the cluster: it only names the
// Unix socket file (/tmp/.s.PGSQL.<port>); clients arrive through pgmem.
const clusterPort = 5432

// StartCluster boots a postmaster on a filesystem that already holds a
// data directory and waits until it accepts connections.
func (e *Engine) StartCluster(ctx context.Context, fs *vfs.FS, opts StartOptions, stderr func([]byte)) (*Cluster, error) {
	if opts.User == "" {
		opts.User = "postgres"
	}
	if opts.Database == "" {
		opts.Database = "postgres"
	}
	cl := &Cluster{e: e, fs: fs, port: clusterPort, Stderr: stderr}
	pfs := fs.Fork()
	pfs.Stderr = cl.stderr
	pfs.Stdout = cl.stderr
	c := host.NewCluster(shmBase, shmLimit)
	c.Log = e.Log
	c.Exec = cl.exec
	cl.c = c
	argv := []string{
		BinDir + "/postgres",
		"-D", PGData,
		"-c", "listen_addresses=",
		"-c", "unix_socket_directories=/tmp",
		"-c", fmt.Sprintf("port=%d", cl.port),
	}
	argv = append(argv, opts.Params...)
	cl.pm = c.Start(pfs, argv, guestEnv(opts.User, opts.Database))
	if err := cl.waitReady(ctx); err != nil {
		cl.Shutdown(context.Background())
		return nil, err
	}
	return cl, nil
}

func (cl *Cluster) stderr(p []byte) {
	if cl.Stderr != nil {
		cl.Stderr(p)
		return
	}
	cl.mu.Lock()
	cl.log.Write(p)
	cl.mu.Unlock()
}

// Log returns and clears the accumulated server log.
func (cl *Cluster) Log() string {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	s := cl.log.String()
	cl.log.Reset()
	return s
}

// exec runs one process on its own goroutine (host.Cluster.Exec).
func (cl *Cluster) exec(p *host.Process) {
	go func() {
		e := cl.e
		status := int32(0)
		var mod guest.Instance
		defer func() {
			// Exit first: a Kill waiting for every process need not wait
			// for the unmapping of 2 GiB of address space per process
			p.Exit(status)
			if mod != nil {
				mod.Close(e.ctx)
			}
		}()
		// LISTEN/UNLISTEN at commit go to the client connection's owner
		p.H.Listen = func(channel string, op int) {
			if cs := p.H.ClientSock(); cs != nil && cs.Listen != nil {
				cs.Listen(channel, op)
			}
		}
		// popen/system from a process (the postmaster runs "postgres -V"
		// to check the executable): a throwaway single-user child
		p.H.Run = func(cmd, stdinPath, stdoutPath string) int {
			// "postgres -V" is the same answer every time: run it once
			if strings.HasSuffix(strings.TrimSpace(cmd), " -V") && stdoutPath != "" {
				e.versionMu.Lock()
				defer e.versionMu.Unlock()
				if e.versionOut == nil {
					var stderr bytes.Buffer
					if code := e.runChild(p.FS, p.H.Env, cmd, stdinPath, stdoutPath, &stderr); code != 0 {
						return code
					}
					out, err := p.FS.ReadFile(stdoutPath)
					if err != vfs.OK {
						return 1
					}
					e.versionOut = out
					return 0
				}
				if err := p.FS.WriteFile(stdoutPath, e.versionOut, 0o644); err != vfs.OK {
					return 1
				}
				return 0
			}
			var stderr bytes.Buffer
			code := e.runChild(p.FS, p.H.Env, cmd, stdinPath, stdoutPath, &stderr)
			if stderr.Len() > 0 {
				cl.stderr(stderr.Bytes())
			}
			return code
		}
		m, err := e.postgres.Instantiate(e.ctx, p.H)
		if err != nil {
			e.logf("process %d: instantiate: %v", p.Pid, err)
			status = 127 << 8
			return
		}
		mod = m
		if _, err := mod.Call("pgmem_init"); err != nil {
			e.logf("process %d: pgmem_init: %v", p.Pid, err)
			status = 127 << 8
			return
		}
		code, _, err := callMain(mod, p.Argv)
		if err != nil {
			// a trap or abort: report it the way a crashed process would
			e.logf("process %d %v: %v", p.Pid, p.Argv[1:], err)
			cl.stderr([]byte(fmt.Sprintf("pgmem: process %d %v: %v\n", p.Pid, p.Argv[1:], err)))
			status = host.SIGABRT
			return
		}
		status = (code & 0xff) << 8
	}()
}

// waitReady polls the pid file for the postmaster's "ready" status.
func (cl *Cluster) waitReady(ctx context.Context) error {
	deadline := time.Now().Add(60 * time.Second)
	for {
		select {
		case <-cl.pm.Done:
			return fmt.Errorf("postmaster exited with status %d\n%s", cl.pm.Status, cl.Log())
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if data, err := cl.fs.ReadFile(PGData + "/postmaster.pid"); err == vfs.OK {
			lines := strings.Split(string(data), "\n")
			if len(lines) >= 8 && strings.HasPrefix(lines[7], "ready") {
				return nil
			}
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("postmaster did not become ready\n%s", cl.Log())
		}
		time.Sleep(time.Millisecond)
	}
}

// Connect hands a client connection to the postmaster. Bytes the backend
// sends go to w; onClose runs when the backend side is closed.
func (cl *Cluster) Connect(w io.Writer, onClose func()) (*host.ConnSock, error) {
	s := host.NewConnSock(w)
	s.OnClose = onClose
	if err := cl.c.Connect(s); err != nil {
		return nil, err
	}
	return s, nil
}

// Dead reports whether the postmaster has exited.
func (cl *Cluster) Dead() bool { return cl.c.Dead() }

// Kill ends every process at once, without a shutdown: for a data
// directory that is discarded anyway (Close, Restore), the checkpoint and
// the orderly exits of a fast shutdown would be wasted work. Every process
// exits at its next host call (the blocked ones wake), so this takes well
// under a millisecond unless a process is deep in a computation.
func (cl *Cluster) Kill() {
	c := cl.c
	c.KillAll(host.SIGKILL)
	deadline := time.After(5 * time.Second)
	for _, p := range c.Processes() {
		select {
		case <-p.Done:
		case <-deadline:
			cl.e.logf("pgmem: process %d did not exit after SIGKILL", p.Pid)
		}
	}
	// releasing the segment files (tens of MB of pages) takes a few
	// milliseconds the caller need not wait for
	go c.Close()
}

// Shutdown stops the cluster: a fast shutdown (SIGINT), then an immediate
// one, then SIGKILL, waiting for every process to exit.
func (cl *Cluster) Shutdown(ctx context.Context) error {
	c := cl.c
	steps := []struct {
		sig  int32
		wait time.Duration
	}{{host.SIGINT, 20 * time.Second}, {host.SIGQUIT, 5 * time.Second}, {host.SIGKILL, 5 * time.Second}}
	var err error
	for _, st := range steps {
		if c.Dead() {
			break
		}
		if st.sig == host.SIGKILL {
			c.KillAll(st.sig)
		} else {
			c.Kill(nil, cl.pm.Pid, st.sig)
		}
		select {
		case <-cl.pm.Done:
		case <-time.After(st.wait):
			err = fmt.Errorf("pgmem: postmaster did not stop after signal %d", st.sig)
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
		err = nil
	}
	// the children notice the postmaster's death and exit on their own
	deadline := time.After(10 * time.Second)
	for _, p := range c.Processes() {
		select {
		case <-p.Done:
		case <-deadline:
			c.KillAll(host.SIGKILL)
			if err == nil {
				err = errors.New("pgmem: some processes did not exit")
			}
		}
	}
	go c.Close()
	return err
}
