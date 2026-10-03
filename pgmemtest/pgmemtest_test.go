package pgmemtest_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shibukawa/pgmem"
	"github.com/shibukawa/pgmem/pgmemtest"
)

var fx *pgmemtest.Fixture

func TestMain(m *testing.M) {
	os.Exit(pgmemtest.Run(m, pgmemtest.Options{
		Options:  pgmem.Options{Database: "app", User: "tester"},
		MaxForks: 2,
		Prepare: func(ctx context.Context, db *sql.DB, dsn string) error {
			if dsn == "" {
				panic("empty dsn")
			}
			_, err := db.ExecContext(ctx, `
				CREATE TABLE users(id serial PRIMARY KEY, name text NOT NULL);
				INSERT INTO users(name) VALUES ('alice'), ('bob');
			`)
			return err
		},
	}, func(f *pgmemtest.Fixture) { fx = f }))
}

func userCount(t *testing.T, db *sql.DB) int {
	t.Helper()
	var n int
	if err := db.QueryRow(`SELECT count(*) FROM users`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

// Every test starts from the seed and its writes stay private, even when
// tests run in parallel and MaxForks makes some of them wait.
func TestIsolation(t *testing.T) {
	for _, name := range []string{"a", "b", "c", "d", "e"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			db := fx.DB(t)
			if n := userCount(t, db); n != 2 {
				t.Fatalf("start: %d users, want 2", n)
			}
			if _, err := db.Exec(`INSERT INTO users(name) VALUES ($1)`, name); err != nil {
				t.Fatal(err)
			}
			if n := userCount(t, db); n != 3 {
				t.Fatalf("after insert: %d users, want 3", n)
			}
		})
	}
}

func TestPgxConn(t *testing.T) {
	conn := fx.PgxConn(t)
	var db, user string
	if err := conn.QueryRow(t.Context(), `SELECT current_database(), current_user`).Scan(&db, &user); err != nil || db != "app" || user != "tester" {
		t.Fatalf("db = %q user = %q %v", db, user, err)
	}
}

func TestPgxPoolDefaultSize(t *testing.T) {
	pool := fx.PgxPool(t)
	ctx := t.Context()
	var wg sync.WaitGroup
	errs := make(chan error, 8)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tx, err := pool.Begin(ctx)
			if err != nil {
				errs <- err
				return
			}
			if _, err := tx.Exec(ctx, `INSERT INTO users(name) VALUES ('x')`); err != nil {
				errs <- err
				return
			}
			errs <- tx.Commit(ctx)
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
	var n int
	if err := pool.QueryRow(ctx, `SELECT count(*) FROM users`).Scan(&n); err != nil || n != 10 {
		t.Fatalf("count = %d, %v", n, err)
	}
}

func TestTemplateUntouched(t *testing.T) {
	conn, err := sql.Open("pgx", fx.Template().DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if n := userCount(t, conn); n != 2 {
		t.Fatalf("template has %d users, want 2", n)
	}
}

func TestObjectHandlesShareOneFork(t *testing.T) {
	d := fx.For(t)
	db := d.DB()
	if _, err := db.Exec("INSERT INTO users(name) VALUES ('shared-handles')"); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := d.PgxPool().QueryRow(t.Context(), "SELECT count(*) FROM users").Scan(&n); err != nil || n != 3 {
		t.Fatalf("different fork: %d %v", n, err)
	}
	if err := d.Reset(t.Context()); err != nil {
		t.Fatal(err)
	}
	if userCount(t, db) != 2 {
		t.Fatal("reset failed")
	}
}

func TestSharedHandles(t *testing.T) {
	db, err := fx.SharedDB()
	if err != nil {
		t.Fatal(err)
	}
	pool, err := fx.SharedPgxPool()
	if err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(t.Context(), "SELECT count(*) FROM users").Scan(&n); err != nil || n != userCount(t, db) {
		t.Fatalf("shared handles: %d %v", n, err)
	}
}

func TestNetworkHelpers(t *testing.T) {
	for _, transport := range []string{"tcp", "unix"} {
		t.Run(transport, func(t *testing.T) {
			f, err := pgmemtest.New(t.Context(), pgmemtest.Options{Transport: transport})
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { f.Close() })
			conn := f.For(t).PgxConn()
			if got := conn.PgConn().Conn().RemoteAddr().Network(); got != transport {
				t.Fatalf("transport %s, got %s", transport, got)
			}
		})
	}
}

func TestDriverConfigurationKeepsForkAndTransport(t *testing.T) {
	d := fx.For(t)
	configure := func(cfg *pgx.ConnConfig) {
		cfg.RuntimeParams["application_name"] = "configured"
		cfg.Host = "wrong.invalid"
		cfg.DialFunc = func(context.Context, string, string) (net.Conn, error) { return nil, errors.New("must be replaced") }
	}
	db := d.DB(configure)
	var name string
	if err := db.QueryRow("SHOW application_name").Scan(&name); err != nil || name != "configured" {
		t.Fatalf("DB config: %q %v", name, err)
	}
	conn := d.PgxConn(configure)
	if err := conn.QueryRow(t.Context(), "SHOW application_name").Scan(&name); err != nil || name != "configured" {
		t.Fatalf("connection config: %q %v", name, err)
	}
	pool := d.PgxPool(func(cfg *pgxpool.Config) { cfg.MaxConns = 1; configure(cfg.ConnConfig) })
	if pool.Config().MaxConns != 1 {
		t.Fatal("pool configuration ignored")
	}
	if err := pool.QueryRow(t.Context(), "SHOW application_name").Scan(&name); err != nil || name != "configured" {
		t.Fatalf("pool config: %q %v", name, err)
	}
	if _, err := db.Exec("INSERT INTO users(name) VALUES ('driver config')"); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := pool.QueryRow(t.Context(), "SELECT count(*) FROM users").Scan(&n); err != nil || n != 3 {
		t.Fatalf("handles no longer share fork: %d %v", n, err)
	}
}

func TestSharedReadOnlyGuard(t *testing.T) {
	f, err := pgmemtest.New(t.Context(), pgmemtest.Options{SharedReadOnly: true, Prepare: func(ctx context.Context, db *sql.DB, _ string) error {
		_, err := db.ExecContext(ctx, "CREATE TABLE guarded(v int)")
		return err
	}})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Error(err)
		}
	})
	db, err := f.SharedDB()
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO guarded VALUES (1)")
	var pgerr *pgconn.PgError
	if !errors.As(err, &pgerr) || pgerr.Code != "25006" {
		t.Fatalf("shared DB write was not guarded: %v", err)
	}
	pool, err := f.SharedPgxPool()
	if err != nil {
		t.Fatal(err)
	}
	_, err = pool.Exec(t.Context(), "INSERT INTO guarded VALUES (2)")
	if !errors.As(err, &pgerr) || pgerr.Code != "25006" {
		t.Fatalf("shared pool write was not guarded: %v", err)
	}
	if _, err := f.For(t).DB().Exec("INSERT INTO guarded VALUES (3)"); err != nil {
		t.Fatalf("isolated write was guarded: %v", err)
	}
}

func TestDirectHelperForkDeadline(t *testing.T) {
	if os.Getenv("PGMEM_FORK_DEADLINE_HELPER") == "true" {
		f, err := pgmemtest.New(t.Context(), pgmemtest.Options{MaxForks: 1, ForkTimeout: 40 * time.Millisecond})
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { f.Close() })
		occupied, err := f.Snapshot().Fork(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		defer occupied.Close()
		f.DB(t) // must fail through the direct helper's bounded fork acquisition
		t.Fatal("full pool did not fail")
	}
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, self, "-test.run=^TestDirectHelperForkDeadline$")
	cmd.Env = append(os.Environ(), "PGMEM_FORK_DEADLINE_HELPER=true")
	out, err := cmd.CombinedOutput()
	if err == nil || ctx.Err() != nil || !strings.Contains(string(out), "context deadline exceeded") {
		t.Fatalf("helper deadline not applied: %v %s", err, out)
	}
}

// The application keeps its ordinary driver and environment-based config.
type userRepository struct{ db *sql.DB }

func openUserRepository() (*userRepository, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &userRepository{db: db}, nil
}

func (r *userRepository) count() (int, error) {
	var n int
	err := r.db.QueryRow("SELECT count(*) FROM users").Scan(&n)
	return n, err
}

func TestShadowPGRoutesUnchangedRepository(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://bad@127.0.0.1:1/production")
	fx.ShadowPG(t)
	repo, err := openUserRepository()
	if err != nil {
		t.Fatal(err)
	}
	defer repo.db.Close()
	if n, err := repo.count(); err != nil || n != 2 {
		t.Fatalf("count = %d, %v; want 2", n, err)
	}
	if _, err := repo.db.Exec("INSERT INTO users(name) VALUES ('shadow')"); err != nil {
		t.Fatal(err)
	}
}

func TestShadowPGStartsFromPreparedSnapshot(t *testing.T) {
	fx.ShadowPG(t)
	repo, err := openUserRepository()
	if err != nil {
		t.Fatal(err)
	}
	defer repo.db.Close()
	if n, err := repo.count(); err != nil || n != 2 {
		t.Fatalf("count = %d, %v; want 2", n, err)
	}
}

func TestShadowPGUsesLibpqVariables(t *testing.T) {
	fx.ShadowPG(t)
	db, err := sql.Open("pgx", "")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if n := userCount(t, db); n != 2 {
		t.Fatalf("count = %d; want 2", n)
	}
}

func TestShadowPGSharedReadOnlyTarget(t *testing.T) {
	var port int
	for _, name := range []string{"first", "second"} {
		t.Run(name, func(t *testing.T) {
			srv := fx.ShadowPG(t, pgmemtest.ShadowOptions{Shared: true})
			if port != 0 && srv.Port() != port {
				t.Fatalf("shared port = %d; want %d", srv.Port(), port)
			}
			port = srv.Port()
			repo, err := openUserRepository()
			if err != nil {
				t.Fatal(err)
			}
			defer repo.db.Close()
			if n, err := repo.count(); err != nil || n != 2 {
				t.Fatalf("count = %d, %v; want 2", n, err)
			}
		})
	}
}

func TestShadowPGAPIClientUsesUnchangedRepository(t *testing.T) {
	fx.ShadowPG(t, pgmemtest.ShadowOptions{ExtraEnv: []string{"APP_DATABASE_URL"}})
	if os.Getenv("APP_DATABASE_URL") != os.Getenv("DATABASE_URL") {
		t.Fatal("custom application DSN was not routed")
	}
	repo, err := openUserRepository()
	if err != nil {
		t.Fatal(err)
	}
	defer repo.db.Close()
	app := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n, err := repo.count()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, n)
	}))
	defer app.Close()
	resp, err := app.Client().Get(app.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d; want 200", resp.StatusCode)
	}
	var n int
	if _, err := fmt.Fscan(resp.Body, &n); err != nil || n != 2 {
		t.Fatalf("count = %d, %v; want 2", n, err)
	}
}

func TestSharedShadowUsesUnixTransportAndReadGuard(t *testing.T) {
	f, err := pgmemtest.New(t.Context(), pgmemtest.Options{
		Transport: "unix", SharedReadOnly: true,
		Prepare: func(ctx context.Context, db *sql.DB, _ string) error {
			_, err := db.ExecContext(ctx, "CREATE TABLE shadow_guarded(v int)")
			return err
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Error(err)
		}
	})
	f.ShadowPG(t, pgmemtest.ShadowOptions{Shared: true})
	db, err := sql.Open("pgx", "") // application uses ordinary libpq environment settings
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var n int
	if err := db.QueryRow("SELECT count(*) FROM shadow_guarded").Scan(&n); err != nil || n != 0 {
		t.Fatalf("Unix shadow connection: %d %v", n, err)
	}
	_, err = db.Exec("INSERT INTO shadow_guarded VALUES (1)")
	var failure *pgconn.PgError
	if !errors.As(err, &failure) || failure.Code != "25006" {
		t.Fatalf("shadow read guard: %v", err)
	}
}
