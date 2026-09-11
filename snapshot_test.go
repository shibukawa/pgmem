package pgmem_test

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/shibukawa/pgmem"
)

func count(t *testing.T, dsn string) int {
	t.Helper()
	ctx := context.Background()
	c, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)
	var n int
	if err := c.QueryRow(ctx, `SELECT count(*) FROM t`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func insert(t *testing.T, dsn string, v int) {
	t.Helper()
	ctx := context.Background()
	c, err := pgx.Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)
	if _, err := c.Exec(ctx, `INSERT INTO t VALUES ($1)`, v); err != nil {
		t.Fatal(err)
	}
}

// Forks start from the snapshot's state and are isolated from each other
// and from the template server.
func TestSnapshotFork(t *testing.T) {
	s := startServer(t, pgmem.Options{Database: "app", User: "tester"})
	ctx := context.Background()
	c, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Exec(ctx, `CREATE TABLE t(v int); INSERT INTO t VALUES (1), (2)`); err != nil {
		t.Fatal(err)
	}
	c.Close(ctx)

	start := time.Now()
	snap, err := s.Snapshot(ctx, pgmem.SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("snapshot in %s", time.Since(start))
	defer snap.Close()

	// The template keeps working and diverges from the snapshot.
	insert(t, s.DSN(), 3)

	start = time.Now()
	a, err := snap.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("fork in %s", time.Since(start))
	defer a.Close()
	b, err := snap.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()

	if n := count(t, a.DSN()); n != 2 {
		t.Fatalf("fork a starts with %d rows, want 2", n)
	}
	insert(t, a.DSN(), 10)
	insert(t, a.DSN(), 11)
	if n := count(t, a.DSN()); n != 4 {
		t.Fatalf("fork a has %d rows, want 4", n)
	}
	if n := count(t, b.DSN()); n != 2 {
		t.Fatalf("fork b sees %d rows, want 2", n)
	}
	if n := count(t, s.DSN()); n != 3 {
		t.Fatalf("template has %d rows, want 3", n)
	}
	// A fork keeps the configured database and user.
	var db, user string
	cc, err := pgx.Connect(ctx, a.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer cc.Close(ctx)
	if err := cc.QueryRow(ctx, `SELECT current_database(), current_user`).Scan(&db, &user); err != nil || db != "app" || user != "tester" {
		t.Fatalf("db = %q user = %q %v", db, user, err)
	}
	// A fork can itself be snapshotted.
	snap2, err := a.Snapshot(ctx, pgmem.SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer snap2.Close()
	c2, err := snap2.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close()
	if n := count(t, c2.DSN()); n != 4 {
		t.Fatalf("fork of fork has %d rows, want 4", n)
	}
}

// Fork blocks at MaxForks until a fork is closed, and honors ctx.
func TestSnapshotForkLimit(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	snap, err := s.Snapshot(ctx, pgmem.SnapshotOptions{MaxForks: 1})
	if err != nil {
		t.Fatal(err)
	}
	defer snap.Close()
	first, err := snap.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	short, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	if _, err := snap.Fork(short); err != context.DeadlineExceeded {
		t.Fatalf("second fork: err = %v, want deadline exceeded", err)
	}
	done := make(chan *pgmem.Server, 1)
	go func() {
		f, err := snap.Fork(ctx)
		if err != nil {
			t.Error(err)
		}
		done <- f
	}()
	time.Sleep(100 * time.Millisecond)
	first.Close()
	select {
	case f := <-done:
		if f != nil {
			f.Close()
		}
	case <-time.After(10 * time.Second):
		t.Fatal("fork still blocked after the first was closed")
	}
	snap.Wait()
}
