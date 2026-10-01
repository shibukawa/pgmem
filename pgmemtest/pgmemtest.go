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
	"net/url"
	"strconv"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/shibukawa/pgmem"
)

// Options configures New and Run.
type Options struct {
	// Options are the template server's options (database, user, params).
	pgmem.Options
	// Prepare runs once against the template server before the snapshot
	// is taken: load the schema and seed data here. db is connected
	// in-process; dsn is for tools that need a connection string.
	Prepare func(ctx context.Context, db *sql.DB, dsn string) error
	// MaxForks caps the forks alive at once; Fork and the handle helpers
	// block until one is closed when the cap is reached. 0 means the pgmem
	// default, a quarter of the memory limit divided by a fork's cost (see
	// pgmem.SnapshotOptions). go test itself runs at most -parallel tests
	// at once (default GOMAXPROCS), so raise that too when tests wait on
	// something other than the CPU.
	MaxForks int
}

// Fixture is a prepared snapshot that tests fork from.
type Fixture struct {
	template *pgmem.Server
	snap     *pgmem.Snapshot
	sharedMu sync.Mutex
	shared   *pgmem.Server
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
func New(ctx context.Context, opts Options) (*Fixture, error) {
	srv, err := pgmem.Start(ctx, opts.Options)
	if err != nil {
		return nil, err
	}
	if opts.Prepare != nil {
		db := openDB(srv)
		err := opts.Prepare(ctx, db, srv.DSN())
		db.Close()
		if err != nil {
			srv.Close()
			return nil, fmt.Errorf("pgmemtest: prepare: %w", err)
		}
	}
	snap, err := srv.Snapshot(ctx, pgmem.SnapshotOptions{MaxForks: opts.MaxForks})
	if err != nil {
		srv.Close()
		return nil, err
	}
	return &Fixture{template: srv, snap: snap}, nil
}

// Run is a TestMain helper: it builds the fixture, hands it to setup (store
// it in a package variable), runs the tests and tears everything down.
// It returns the exit code for os.Exit. A fixture that fails to start is
// reported on stderr and exits with 1.
func Run(m *testing.M, opts Options, setup func(*Fixture)) int {
	f, err := New(context.Background(), opts)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}
	setup(f)
	code := m.Run()
	f.Close()
	return code
}

// Close releases the snapshot and the template server. Forks still alive
// keep working until they are closed.
func (f *Fixture) Close() error {
	f.sharedMu.Lock()
	shared := f.shared
	f.shared = nil
	f.sharedMu.Unlock()
	var sharedErr error
	if shared != nil {
		sharedErr = shared.Close()
	}
	return errors.Join(sharedErr, f.snap.Close(), f.template.Close())
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
	srv, err := f.snap.Fork(t.Context())
	if err != nil {
		t.Fatalf("pgmemtest: fork: %v", err)
	}
	t.Cleanup(func() { srv.Close() })
	return srv
}

// DB returns a database/sql handle (pgx driver, in-process connection) on
// a fresh fork. Both are closed when t ends. Pools of any size are fine:
// every connection is its own backend process.
func (f *Fixture) DB(t testing.TB) *sql.DB {
	t.Helper()
	db := openDB(f.Fork(t))
	t.Cleanup(func() { db.Close() })
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
	cfg.DialFunc = srv.Dial
	conn, err := pgx.ConnectConfig(t.Context(), cfg)
	if err != nil {
		t.Fatalf("pgmemtest: connect: %v", err)
	}
	t.Cleanup(func() { conn.Close(context.Background()) })
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
	cfg.ConnConfig.DialFunc = srv.Dial
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
			f.shared, err = f.snap.Fork(t.Context())
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
	u, err := url.Parse(srv.DSN())
	if err != nil {
		t.Fatalf("pgmemtest: shadow DSN: %v", err)
	}
	vars := map[string]string{
		"DATABASE_URL": srv.DSN(),
		"PGHOST":       u.Hostname(),
		"PGHOSTADDR":   u.Hostname(),
		"PGPORT":       strconv.Itoa(srv.Port()),
		"PGUSER":       u.User.Username(),
		"PGDATABASE":   u.Path[1:],
		"PGPASSWORD":   "",
		"PGSERVICE":    "",
		"PGSSLMODE":    "disable",
	}
	for _, name := range opts.ExtraEnv {
		if name == "" {
			t.Fatal("pgmemtest: ExtraEnv contains an empty name")
		}
		if _, reserved := vars[name]; reserved {
			t.Fatalf("pgmemtest: ExtraEnv duplicates built-in variable %q", name)
		}
		vars[name] = srv.DSN()
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
