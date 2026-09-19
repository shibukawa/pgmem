package pgmem_test

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shibukawa/pgmem"
)

// forkWithTable starts a template whose table t holds 1 and 2, snapshots
// it and returns the snapshot and a fork.
func forkWithTable(t *testing.T, opts pgmem.Options) (*pgmem.Snapshot, *pgmem.Server) {
	t.Helper()
	s := startServer(t, opts)
	ctx := context.Background()
	c, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Exec(ctx, `CREATE TABLE t(v int); INSERT INTO t VALUES (1), (2)`); err != nil {
		t.Fatal(err)
	}
	c.Close(ctx)
	snap, err := s.Snapshot(ctx, pgmem.SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { snap.Close() })
	f, err := snap.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return snap, f
}

// Reset puts the data back while pooled connections stay connected, and
// their named prepared statements and LISTEN registrations keep working.
func TestResetKeepsConnections(t *testing.T) {
	_, f := forkWithTable(t, pgmem.Options{Database: "app", User: "tester"})
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, f.DSN()) // pgx prepares and caches a named statement per query
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()
	conn, err := pgx.Connect(ctx, f.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	listener, err := pgx.Connect(ctx, f.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close(ctx)

	if _, err := conn.Prepare(ctx, "total", `SELECT count(*) FROM t`); err != nil {
		t.Fatal(err)
	}
	if _, err := listener.Exec(ctx, `LISTEN ch`); err != nil {
		t.Fatal(err)
	}
	if _, err := conn.Exec(ctx, `SET application_name = 'before'`); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		if _, err := pool.Exec(ctx, `INSERT INTO t VALUES ($1)`, 10+i); err != nil {
			t.Fatal(err)
		}
	}
	var n int
	if err := conn.QueryRow(ctx, "total").Scan(&n); err != nil || n != 7 {
		t.Fatalf("before reset: count = %d, %v", n, err)
	}
	connects := pool.Stat().NewConnsCount()

	start := time.Now()
	if err := f.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("reset in %s", time.Since(start))

	if err := conn.QueryRow(ctx, "total").Scan(&n); err != nil || n != 2 {
		t.Fatalf("prepared statement after reset: count = %d, %v", n, err)
	}
	for i := 0; i < 10; i++ {
		if err := pool.QueryRow(ctx, `SELECT count(*) FROM t WHERE v > $1`, 0).Scan(&n); err != nil || n != 2 {
			t.Fatalf("pool after reset: count = %d, %v", n, err)
		}
	}
	if got := pool.Stat().NewConnsCount(); got != connects {
		t.Fatalf("pool opened %d new connections across the reset", got-connects)
	}
	var app, db, user string
	if err := conn.QueryRow(ctx, `SELECT current_setting('application_name'), current_database(), current_user`).Scan(&app, &db, &user); err != nil {
		t.Fatal(err)
	}
	if app == "before" || db != "app" || user != "tester" {
		t.Fatalf("session after reset: application_name=%q database=%q user=%q", app, db, user)
	}
	if _, err := pool.Exec(ctx, `NOTIFY ch, 'after'`); err != nil {
		t.Fatal(err)
	}
	if n := waitNotification(t, listener, 5*time.Second); n == nil || n.Payload != "after" {
		t.Fatalf("LISTEN after reset: %+v", n)
	}
}

// Reset waits for an open transaction and gives up when ctx ends.
func TestResetWaitsForOpenTransaction(t *testing.T) {
	_, f := forkWithTable(t, pgmem.Options{})
	ctx := context.Background()
	c, err := pgx.Connect(ctx, f.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)
	if _, err := c.Exec(ctx, `BEGIN; INSERT INTO t VALUES (3)`); err != nil {
		t.Fatal(err)
	}
	short, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	if err := f.Reset(short); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("reset during a transaction: %v", err)
	}
	if _, err := c.Exec(ctx, `COMMIT`); err != nil {
		t.Fatal(err)
	}
	if err := f.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	if n := count(t, f.DSN()); n != 2 {
		t.Fatalf("count after reset = %d, want 2", n)
	}
}

// Restore takes any server back to a snapshot of the same database; Reset
// is for forks.
func TestRestoreTemplate(t *testing.T) {
	s := startServer(t, pgmem.Options{Database: "app"})
	ctx := context.Background()
	c, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Exec(ctx, `CREATE TABLE t(v int); INSERT INTO t VALUES (1)`); err != nil {
		t.Fatal(err)
	}
	snap, err := s.Snapshot(ctx, pgmem.SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer snap.Close()
	if _, err := c.Exec(ctx, `INSERT INTO t VALUES (2), (3)`); err != nil {
		t.Fatal(err)
	}
	if err := s.Reset(ctx); err == nil {
		t.Fatal("Reset on a server that is not a fork succeeded")
	}
	if err := s.Restore(ctx, snap); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := c.QueryRow(ctx, `SELECT count(*) FROM t`).Scan(&n); err != nil || n != 1 {
		t.Fatalf("count after restore = %d, %v", n, err)
	}
	c.Close(ctx)

	other := startServer(t, pgmem.Options{Database: "other"})
	if err := other.Restore(ctx, snap); err == nil || !strings.Contains(err.Error(), `"app"`) {
		t.Fatalf("restoring a snapshot of another database: %v", err)
	}
}

// A closed server leaves an idle client connection open and answers the
// client's next message with 57P01, so a pool holding the connection never
// sees its socket drop while idle.
func TestCloseLeavesIdleConnectionsToTheClient(t *testing.T) {
	_, f := forkWithTable(t, pgmem.Options{})
	ctx := context.Background()
	c, err := pgx.Connect(ctx, f.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)
	if _, err := c.Exec(ctx, `SELECT 1`); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	raw := c.PgConn().Conn()
	raw.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
	var one [1]byte
	_, err = raw.Read(one[:])
	var ne net.Error
	if !errors.As(err, &ne) || !ne.Timeout() {
		t.Fatalf("idle connection of a closed server: read = %v, want a timeout", err)
	}
	raw.SetReadDeadline(time.Time{})
	_, err = c.Exec(ctx, `SELECT 1`)
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) || pgErr.Code != "57P01" {
		t.Fatalf("query on a closed server: %v, want 57P01", err)
	}
	if _, err := pgx.Connect(ctx, f.DSN()); err == nil {
		t.Fatal("closed server still accepts connections")
	}
}
