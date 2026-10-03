package pgmemprocess_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/shibukawa/pgmem/pgmemprocess"
)

var binary string

func TestMain(m *testing.M) {
	if os.Getenv("PGMEM_CLIENT_TEST_HELPER") == "hang" {
		for {
			time.Sleep(time.Second)
		}
	}
	if os.Getenv("PGMEM_CLIENT_TEST_HELPER") == "ready-hang" {
		fmt.Println(`{"event":"ready","protocol":1,"server":{"id":"template","dsn":"postgres://postgres@127.0.0.1:1/postgres?sslmode=disable"}}`)
		for {
			time.Sleep(time.Second)
		}
	}
	binary = os.Getenv("PGMEM_BINARY")
	var dir string
	if binary == "" {
		var err error
		dir, err = os.MkdirTemp("", "pgmem-client-test-")
		if err != nil {
			panic(err)
		}
		binary = filepath.Join(dir, "pgmem")
		if runtime.GOOS == "windows" {
			binary += ".exe"
		}
		cmd := exec.Command("go", "build", "-o", binary, "./cmd/pgmem")
		cmd.Dir = ".."
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
	code := m.Run()
	if dir != "" {
		os.RemoveAll(dir)
	}
	os.Exit(code)
}

func socketParent() string {
	if runtime.GOOS == "windows" {
		return os.TempDir()
	}
	return "/tmp"
}

func prepare(ctx context.Context, db *sql.DB, _ string) error {
	_, err := db.ExecContext(ctx, "CREATE TABLE t(v int); INSERT INTO t VALUES (1),(2)")
	return err
}
func count(t *testing.T, db *sql.DB) int {
	t.Helper()
	var n int
	if err := db.QueryRow("SELECT count(*) FROM t").Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestFixtureTransports(t *testing.T) {
	for _, transport := range []string{"tcp", "unix"} {
		t.Run(transport, func(t *testing.T) {
			opts := pgmemprocess.FixtureOptions{Options: pgmemprocess.Options{Binary: binary, Transport: transport, Database: "app"}, Prepare: prepare, MaxForks: 3, SharedReadOnly: true}
			if transport == "unix" {
				dir, err := os.MkdirTemp(socketParent(), "pgm space-")
				if err != nil {
					t.Fatal(err)
				}
				defer os.RemoveAll(dir)
				opts.SocketDir = dir
			}
			f, err := pgmemprocess.NewFixture(t.Context(), opts)
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if err := f.Close(); err != nil {
					t.Error(err)
				}
			})
			shared, err := f.SharedDB()
			if err != nil {
				t.Fatal(err)
			}
			if count(t, shared) != 2 {
				t.Fatal("unprepared shared fork")
			}
			if _, err := shared.Exec("INSERT INTO t VALUES (99)"); err == nil {
				t.Fatal("shared DB write escaped guard")
			}
			sharedPool, err := f.SharedPgxPool()
			if err != nil {
				t.Fatal(err)
			}
			if _, err := sharedPool.Exec(t.Context(), "INSERT INTO t VALUES (99)"); err == nil {
				t.Fatal("shared pool write escaped guard")
			}
			t.Run("same database and stable reset", func(t *testing.T) {
				d := f.For(t)
				db := d.DB()
				pool := d.PgxPool()
				if _, err := db.Exec("INSERT INTO t VALUES (3)"); err != nil {
					t.Fatal(err)
				}
				var n int
				if err := pool.QueryRow(t.Context(), "SELECT count(*) FROM t").Scan(&n); err != nil || n != 3 {
					t.Fatalf("pool observes same fork: %d %v", n, err)
				}
				if d.DSN() == f.Template().DSN() {
					t.Fatal("test uses template")
				}
				if count(t, shared) != 2 {
					t.Fatal("write escaped fork")
				}
				if err := d.Reset(t.Context()); err != nil {
					t.Fatal(err)
				}
				if count(t, db) != 2 {
					t.Fatal("reset did not restore open connection")
				}
			})
			t.Run("parallel writes", func(t *testing.T) {
				for i := 0; i < 2; i++ {
					t.Run(fmt.Sprintf("writer-%d", i), func(t *testing.T) {
						t.Parallel()
						db := f.For(t).DB()
						if _, err := db.Exec("CREATE TABLE isolated(v int PRIMARY KEY); INSERT INTO isolated VALUES (1); INSERT INTO t VALUES (3)"); err != nil {
							t.Fatal(err)
						}
						if count(t, db) != 3 {
							t.Fatal("another writer changed this fork")
						}
					})
				}
			})
			a, err := f.Snapshot().Fork(t.Context())
			if err != nil {
				t.Fatal(err)
			}
			b, err := f.Snapshot().Fork(t.Context())
			if err != nil {
				t.Fatal(err)
			}
			if a.DSN() == b.DSN() || a.DSN() == f.Template().DSN() {
				t.Fatal("fork endpoints collide")
			}
			ctx, cancel := context.WithTimeout(t.Context(), 30*time.Millisecond)
			_, err = f.Snapshot().Fork(ctx)
			cancel()
			if err == nil {
				t.Fatal("full pool did not time out")
			}
			if err := a.Close(); err != nil {
				t.Fatal(err)
			}
			c, err := f.Snapshot().Fork(t.Context())
			if err != nil {
				t.Fatal(err)
			}
			c.Close()
			b.Close()
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
			if transport == "unix" {
				entries, err := os.ReadDir(opts.SocketDir)
				if err != nil || len(entries) != 0 {
					t.Fatalf("socket directories leaked: %v %v", entries, err)
				}
			}
		})
	}
}

func TestPrepareFailureCleansSockets(t *testing.T) {
	dir, err := os.MkdirTemp(socketParent(), "pgm-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	failure := errors.New("migration failed")
	_, err = pgmemprocess.NewFixture(t.Context(), pgmemprocess.FixtureOptions{Options: pgmemprocess.Options{Binary: binary, Transport: "unix", SocketDir: dir}, Prepare: func(context.Context, *sql.DB, string) error { return failure }})
	if !errors.Is(err, failure) {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(dir)
	if err != nil || len(entries) != 0 {
		t.Fatalf("failed prepare leaked sockets: %v %v", entries, err)
	}
}

func TestChildExitReleasesRequests(t *testing.T) {
	p, err := pgmemprocess.Start(t.Context(), pgmemprocess.Options{Binary: binary})
	if err != nil {
		t.Fatal(err)
	}
	process, err := os.FindProcess(p.PID())
	if err != nil {
		t.Fatal(err)
	}
	process.Kill()
	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()
	if _, err := p.Template.Snapshot(ctx, 1); err == nil {
		t.Fatal("dead child accepted request")
	}
	p.Close()
}

func TestInvalidOptions(t *testing.T) {
	for _, opts := range []pgmemprocess.Options{{Binary: "/missing/pgmem"}, {Binary: binary, Transport: "invalid"}, {Binary: binary, SocketDir: "/tmp"}} {
		if _, err := pgmemprocess.Start(t.Context(), opts); err == nil {
			t.Fatalf("accepted %+v", opts)
		}
	}
}

func TestStartupDeadline(t *testing.T) {
	t.Setenv("PGMEM_CLIENT_TEST_HELPER", "hang")
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	start := time.Now()
	_, err = pgmemprocess.Start(t.Context(), pgmemprocess.Options{Binary: self, StartupTimeout: 30 * time.Millisecond, ShutdownTimeout: time.Second})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	if time.Since(start) > 2*time.Second {
		t.Fatal("startup shutdown exceeded deadline")
	}
}

func TestShutdownDeadline(t *testing.T) {
	t.Setenv("PGMEM_CLIENT_TEST_HELPER", "ready-hang")
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	p, err := pgmemprocess.Start(t.Context(), pgmemprocess.Options{Binary: self, ShutdownTimeout: 30 * time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	start := time.Now()
	if err := p.Close(); err == nil {
		t.Fatal("stuck child was not killed")
	}
	if time.Since(start) > 2*time.Second {
		t.Fatal("shutdown exceeded deadline")
	}
}

func TestUnixStartupDiagnostics(t *testing.T) {
	_, err := pgmemprocess.Start(t.Context(), pgmemprocess.Options{Binary: binary, Transport: "unix", SocketDir: filepath.Join(socketParent(), strings.Repeat("x", 110))})
	if err == nil || !strings.Contains(err.Error(), "shorter SocketDir") {
		t.Fatal(err)
	}
	_, err = pgmemprocess.Start(t.Context(), pgmemprocess.Options{Binary: binary, Transport: "unix", SocketDir: filepath.Join(socketParent(), "pgmem-missing-42")})
	if err == nil || !strings.Contains(err.Error(), "socket directory") {
		t.Fatal(err)
	}
}
