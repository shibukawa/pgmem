package pgmemtest_test

import (
	"context"
	"database/sql"
	"os"
	"sync"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"

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
