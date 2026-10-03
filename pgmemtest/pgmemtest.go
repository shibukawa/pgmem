// Package pgmemtest gives every test its own PostgreSQL, forked from a
// snapshot that TestMain prepares once.
//
//	var fx *pgmemtest.Fixture
//
//	func TestMain(m *testing.M) {
//		os.Exit(pgmemtest.Run(m, pgmemtest.Options{
//			Prepare: func(ctx context.Context, db *sql.DB, dsn string) error {
//				// run migrations, load seed data
//				_, err := db.ExecContext(ctx, schema)
//				return err
//			},
//		}, func(f *pgmemtest.Fixture) { fx = f }))
//	}
//
//	func TestOrders(t *testing.T) {
//		t.Parallel()
//		db := fx.DB(t) // a fresh copy of the prepared database, closed when t ends
//		// ...
//	}
//
// Forks are full backends, so tests that use them can run in parallel;
// Options.MaxForks bounds how many exist at once.
package pgmemtest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/shibukawa/pgmem"
	"github.com/shibukawa/pgmem/pgmemfixture"
)

// Options configures New and Run.
type Options struct {
	// Options are the template server's options (database, user, params).
	pgmem.Options
	// Prepare runs once against the template server before the snapshot
	// is taken: load the schema and seed data here. db uses Transport;
	// dsn is for tools that need a connection string.
	Prepare func(ctx context.Context, db *sql.DB, dsn string) error
	// MaxForks caps the forks alive at once; Fork and the handle helpers
	// block until one is closed when the cap is reached. 0 means the pgmem
	// default, a quarter of the memory limit divided by a fork's cost (see
	// pgmem.SnapshotOptions). go test itself runs at most -parallel tests
	// at once (default GOMAXPROCS), so raise that too when tests wait on
	// something other than the CPU.
	MaxForks int
	// SharedReadOnly guards connections created by SharedDB and SharedPgxPool.
	SharedReadOnly bool
	// ForkTimeout bounds all fixture fork waits. Zero uses 30 seconds.
	ForkTimeout time.Duration
	// SnapshotTimeout bounds the wait for open template transactions. Zero uses 30 seconds.
	SnapshotTimeout time.Duration
	// Transport is "inprocess" (default), "tcp" or "unix" for helper connections.
	// Unix selects SocketDir (/tmp on Unix, os.TempDir() on Windows).
	Transport string
}

// Fixture is a prepared snapshot that tests fork from.
type Fixture struct {
	template       *pgmem.Server
	snap           *pgmem.Snapshot
	sharedMu       sync.Mutex
	shared         *pgmem.Server
	sharedDB       *sql.DB
	sharedPool     *pgxpool.Pool
	forkTimeout    time.Duration
	network        bool
	sharedReadOnly bool
}

// ShadowOptions controls ShadowPG. The default is a fresh fork per test.
type ShadowOptions struct {
	// Shared reuses one fork. Use it only for read-only tests: writes persist.
	Shared bool
	// ExtraEnv sets additional application-specific DSN variables.
	ExtraEnv []string
}

// New starts a template server, runs Options.Prepare against it and takes
// the snapshot. Close it when the tests are done.
func New(ctx context.Context, opts Options) (fixture *Fixture, err error) {
	switch opts.Transport {
	case "", "inprocess":
	case "tcp":
		if opts.SocketDir != "" {
			return nil, errors.New("pgmemtest: TCP transport cannot use SocketDir")
		}
	case "unix":
		if opts.SocketDir == "" {
			opts.SocketDir = "/tmp"
			if runtime.GOOS == "windows" {
				opts.SocketDir = os.TempDir()
			}
		}
	default:
		return nil, fmt.Errorf("pgmemtest: unknown transport %q", opts.Transport)
	}
	srv, err := pgmem.Start(ctx, opts.Options)
	if err != nil {
		return nil, err
	}
	ready := false
	defer func() {
		if !ready {
			err = errors.Join(err, srv.Close())
		}
	}()
	if opts.Prepare != nil {
		var endpoint pgmemfixture.Endpoint = srv
		if opts.Transport == "tcp" || opts.Transport == "unix" {
			endpoint = networkServer{srv}
		}
		db, err := pgmemfixture.OpenDB(endpoint)
		if err != nil {
			return nil, err
		}
		err = func() (prepareErr error) {
			defer func() { prepareErr = errors.Join(prepareErr, db.Close()) }()
			return opts.Prepare(ctx, db, srv.DSN())
		}()
		if err != nil {
			return nil, fmt.Errorf("pgmemtest: prepare: %w", err)
		}
	}
	duration := opts.SnapshotTimeout
	if duration <= 0 {
		duration = 30 * time.Second
	}
	snapshotCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	snap, err := srv.Snapshot(snapshotCtx, pgmem.SnapshotOptions{MaxForks: opts.MaxForks})
	if err != nil {
		return nil, err
	}
	wait := opts.ForkTimeout
	if wait <= 0 {
		wait = 30 * time.Second
	}
	ready = true
	return &Fixture{template: srv, snap: snap, forkTimeout: wait, network: opts.Transport == "tcp" || opts.Transport == "unix", sharedReadOnly: opts.SharedReadOnly}, nil
}

// Run is a TestMain helper: it builds the fixture, hands it to setup (store
// it in a package variable), runs the tests and tears everything down.
// It returns the exit code for os.Exit. A fixture that fails to start is
// reported on stderr and exits with 1.
func Run(m *testing.M, opts Options, setup func(*Fixture)) (code int) {
	f, err := New(context.Background(), opts)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(errOut, err)
			code = 1
		}
	}()
	setup(f)
	return m.Run()
}

// Close releases the snapshot and the template server. Forks still alive
// keep working until they are closed.
func (f *Fixture) Close() error {
	f.sharedMu.Lock()
	defer f.sharedMu.Unlock()
	var failures []error
	if f.sharedDB != nil {
		failures = append(failures, f.sharedDB.Close())
	}
	if f.sharedPool != nil {
		f.sharedPool.Close()
	}
	if f.shared != nil {
		failures = append(failures, f.shared.Close())
	}
	return errors.Join(append(failures, f.snap.Close(), f.template.Close())...)
}

// For acquires one fork; all handles on the returned object refer to it.
// Call For again only when the test deliberately needs another database.
func (f *Fixture) For(t testing.TB) *pgmemfixture.TestDB {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), f.forkTimeout)
	defer cancel()
	srv, err := f.snap.Fork(ctx)
	if err != nil {
		t.Fatalf("pgmemtest: fork: %v", err)
	}
	return pgmemfixture.New(t, f.endpoint(srv))
}

// SharedDB shares one prepared fork for read-only tests. Fixture.Close owns cleanup.
func (f *Fixture) SharedDB() (*sql.DB, error) {
	f.sharedMu.Lock()
	defer f.sharedMu.Unlock()
	if f.shared == nil {
		ctx, cancel := context.WithTimeout(context.Background(), f.forkTimeout)
		defer cancel()
		srv, err := f.snap.Fork(ctx)
		if err != nil {
			return nil, err
		}
		f.shared = srv
	}
	if f.sharedDB == nil {
		var err error
		f.sharedDB, err = pgmemfixture.OpenDB(f.sharedEndpoint())
		if err != nil {
			return nil, err
		}
	}
	return f.sharedDB, nil
}

// SharedPgxPool shares the same read-only fork as SharedDB.
func (f *Fixture) SharedPgxPool() (*pgxpool.Pool, error) {
	f.sharedMu.Lock()
	defer f.sharedMu.Unlock()
	if f.shared == nil {
		ctx, cancel := context.WithTimeout(context.Background(), f.forkTimeout)
		defer cancel()
		srv, err := f.snap.Fork(ctx)
		if err != nil {
			return nil, err
		}
		f.shared = srv
	}
	if f.sharedPool == nil {
		cfg, err := pgxpool.ParseConfig(f.sharedEndpoint().DSN())
		if err != nil {
			return nil, err
		}
		cfg.ConnConfig.DialFunc = f.sharedEndpoint().Dial
		f.sharedPool, err = pgxpool.NewWithConfig(context.Background(), cfg)
		if err != nil {
			return nil, err
		}
	}
	return f.sharedPool, nil
}

// Template is the server the snapshot was taken from. It keeps running;
// writes to it after New do not reach the forks.
func (f *Fixture) Template() *pgmem.Server { return f.template }

// Snapshot is the underlying snapshot, for callers that manage forks
// themselves.
func (f *Fixture) Snapshot() *pgmem.Snapshot { return f.snap }

// Fork starts a server for t on a fresh copy of the snapshot and closes it
// when t ends. It blocks while Options.MaxForks forks are alive.
func (f *Fixture) Fork(t testing.TB) *pgmem.Server {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), f.forkTimeout)
	defer cancel()
	srv, err := f.snap.Fork(ctx)
	if err != nil {
		t.Fatalf("pgmemtest: fork: %v", err)
	}
	t.Cleanup(func() {
		if err := srv.Close(); err != nil {
			t.Errorf("pgmemtest: close fork: %v", err)
		}
	})
	return srv
}

// DB returns a database/sql handle (pgx driver, in-process connection) on
// a fresh fork. Both are closed when t ends. Pools of any size are fine:
// every connection is its own backend process.
func (f *Fixture) DB(t testing.TB) *sql.DB {
	t.Helper()
	db, err := pgmemfixture.OpenDB(f.endpoint(f.Fork(t)))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("pgmemtest: close DB: %v", err)
		}
	})
	return db
}

// PgxConn returns a pgx connection on a fresh fork, closed when t ends.
func (f *Fixture) PgxConn(t testing.TB) *pgx.Conn {
	t.Helper()
	srv := f.Fork(t)
	cfg, err := pgx.ParseConfig(srv.DSN())
	if err != nil {
		t.Fatal(err)
	}
	cfg.DialFunc = f.endpoint(srv).Dial
	conn, err := pgx.ConnectConfig(t.Context(), cfg)
	if err != nil {
		t.Fatalf("pgmemtest: connect: %v", err)
	}
	t.Cleanup(func() {
		if err := conn.Close(context.Background()); err != nil {
			t.Errorf("pgmemtest: close connection: %v", err)
		}
	})
	return conn
}

// PgxPool returns a pgxpool on a fresh fork, closed when t ends.
func (f *Fixture) PgxPool(t testing.TB) *pgxpool.Pool {
	t.Helper()
	srv := f.Fork(t)
	cfg, err := pgxpool.ParseConfig(srv.DSN())
	if err != nil {
		t.Fatal(err)
	}
	cfg.ConnConfig.DialFunc = f.endpoint(srv).Dial
	pool, err := pgxpool.NewWithConfig(t.Context(), cfg)
	if err != nil {
		t.Fatalf("pgmemtest: pool: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// DSN returns the connection string of a fresh fork that is closed when t
// ends, for code that opens its own driver.
func (f *Fixture) DSN(t testing.TB) string {
	t.Helper()
	return f.Fork(t).DSN()
}

// ShadowPG routes applications that read DATABASE_URL or libpq's PG*
// environment variables to pgmem without changing their datasource code.
// Call it before constructing the application or its connection pool.
// It uses testing.TB.Setenv, so the test and its ancestors cannot be parallel.
// Go cannot replace an already-open *sql.DB or pgx pool through this helper.
func (f *Fixture) ShadowPG(t testing.TB, options ...ShadowOptions) *pgmem.Server {
	t.Helper()
	if len(options) > 1 {
		t.Fatal("pgmemtest: ShadowPG accepts at most one ShadowOptions")
	}
	var opts ShadowOptions
	if len(options) == 1 {
		opts = options[0]
	}
	var srv *pgmem.Server
	if opts.Shared {
		f.sharedMu.Lock()
		if f.shared == nil {
			var err error
			ctx, cancel := context.WithTimeout(t.Context(), f.forkTimeout)
			f.shared, err = f.snap.Fork(ctx)
			cancel()
			if err != nil {
				f.sharedMu.Unlock()
				t.Fatalf("pgmemtest: shared shadow fork: %v", err)
			}
		}
		srv = f.shared
		f.sharedMu.Unlock()
	} else {
		srv = f.Fork(t)
	}
	dsn := srv.DSN()
	if opts.Shared && f.sharedReadOnly {
		dsn = pgmemfixture.ReadOnly(srv).DSN()
	}
	u, err := url.Parse(dsn)
	if err != nil {
		t.Fatalf("pgmemtest: shadow DSN: %v", err)
	}
	vars := map[string]string{
		"DATABASE_URL": dsn,
		"PGHOST":       srv.Host(),
		"PGHOSTADDR":   "",
		"PGPORT":       strconv.Itoa(srv.Port()),
		"PGUSER":       u.User.Username(),
		"PGDATABASE":   u.Path[1:],
		"PGPASSWORD":   "",
		"PGSERVICE":    "",
		"PGSSLMODE":    "disable",
	}
	if opts.Shared && f.sharedReadOnly {
		vars["PGOPTIONS"] = u.Query().Get("options")
	}
	for _, name := range opts.ExtraEnv {
		if name == "" {
			t.Fatal("pgmemtest: ExtraEnv contains an empty name")
		}
		if _, reserved := vars[name]; reserved {
			t.Fatalf("pgmemtest: ExtraEnv duplicates built-in variable %q", name)
		}
		vars[name] = dsn
	}
	for name, value := range vars {
		t.Setenv(name, value)
	}
	return srv
}

func openDB(srv *pgmem.Server) *sql.DB {
	cfg, err := pgx.ParseConfig(srv.DSN())
	if err != nil {
		panic(err) // DSN is generated by pgmem
	}
	cfg.DialFunc = srv.Dial
	return stdlib.OpenDB(*cfg)
}

type networkServer struct{ *pgmem.Server }

func (s networkServer) Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	return (&net.Dialer{}).DialContext(ctx, network, addr)
}
func (f *Fixture) endpoint(s *pgmem.Server) pgmemfixture.Endpoint {
	if f.network {
		return networkServer{s}
	}
	return s
}

func (f *Fixture) sharedEndpoint() pgmemfixture.Endpoint {
	ep := f.endpoint(f.shared)
	if f.sharedReadOnly {
		return pgmemfixture.ReadOnly(ep)
	}
	return ep
}
