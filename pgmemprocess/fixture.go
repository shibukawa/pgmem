package pgmemprocess

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shibukawa/pgmem/pgmemfixture"
)

type FixtureOptions struct {
	Options
	Prepare         func(context.Context, *sql.DB, string) error
	MaxForks        int
	SharedReadOnly  bool          // guard connections made by shared helpers
	ForkTimeout     time.Duration // default 30s
	SnapshotTimeout time.Duration // default 30s
}

type Fixture struct {
	process        *Process
	snapshot       *Snapshot
	forkTimeout    time.Duration
	mu             sync.Mutex
	shared         *Server
	sharedDB       *sql.DB
	sharedPool     *pgxpool.Pool
	sharedReadOnly bool
}

func NewFixture(ctx context.Context, opts FixtureOptions) (fixture *Fixture, err error) {
	p, err := Start(ctx, opts.Options)
	if err != nil {
		return nil, err
	}
	ok := false
	defer func() {
		if !ok {
			err = errors.Join(err, p.Close())
		}
	}()
	if opts.Prepare != nil {
		db, err := pgmemfixture.OpenDB(p.Template)
		if err != nil {
			return nil, err
		}
		err = func() (prepareErr error) {
			defer func() { prepareErr = errors.Join(prepareErr, db.Close()) }()
			return opts.Prepare(ctx, db, p.Template.DSN())
		}()
		if err != nil {
			return nil, fmt.Errorf("pgmem: prepare: %w", err)
		}
	}
	duration := opts.SnapshotTimeout
	if duration <= 0 {
		duration = 30 * time.Second
	}
	snapshotCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	sn, err := p.Template.Snapshot(snapshotCtx, opts.MaxForks)
	if err != nil {
		return nil, err
	}
	forkTimeout := opts.ForkTimeout
	if forkTimeout <= 0 {
		forkTimeout = 30 * time.Second
	}
	ok = true
	return &Fixture{process: p, snapshot: sn, forkTimeout: forkTimeout, sharedReadOnly: opts.SharedReadOnly}, nil
}

func Run(m *testing.M, opts FixtureOptions, setup func(*Fixture)) (code int) {
	f, err := NewFixture(context.Background(), opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			code = 1
		}
	}()
	setup(f)
	return m.Run()
}

func (f *Fixture) For(t testing.TB) *pgmemfixture.TestDB {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), f.forkTimeout)
	defer cancel()
	srv, err := f.snapshot.Fork(ctx)
	if err != nil {
		t.Fatalf("pgmem: acquire test database (close unused forks or raise MaxForks): %v", err)
	}
	return pgmemfixture.New(t, srv)
}
func (f *Fixture) Template() *Server   { return f.process.Template }
func (f *Fixture) Snapshot() *Snapshot { return f.snapshot }

func (f *Fixture) sharedServer() (*Server, error) {
	if f.shared != nil {
		return f.shared, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), f.forkTimeout)
	defer cancel()
	srv, err := f.snapshot.Fork(ctx)
	if err == nil {
		f.shared = srv
	}
	return srv, err
}
func (f *Fixture) SharedDB() (*sql.DB, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	srv, err := f.sharedServer()
	if err != nil {
		return nil, err
	}
	if f.sharedDB == nil {
		f.sharedDB, err = pgmemfixture.OpenDB(f.sharedEndpoint(srv))
	}
	return f.sharedDB, err
}
func (f *Fixture) SharedPgxPool() (*pgxpool.Pool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	srv, err := f.sharedServer()
	if err != nil {
		return nil, err
	}
	if f.sharedPool == nil {
		f.sharedPool, err = pgxpool.New(context.Background(), f.sharedEndpoint(srv).DSN())
	}
	return f.sharedPool, err
}
func (f *Fixture) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	var failures []error
	if f.sharedDB != nil {
		failures = append(failures, f.sharedDB.Close())
	}
	if f.sharedPool != nil {
		f.sharedPool.Close()
	}
	// The process closes every remaining server and snapshot, even if a test failed.
	return errors.Join(append(failures, f.process.Close())...)
}

func (f *Fixture) sharedEndpoint(srv *Server) pgmemfixture.Endpoint {
	if f.sharedReadOnly {
		return pgmemfixture.ReadOnly(srv)
	}
	return srv
}
