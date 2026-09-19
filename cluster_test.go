package pgmem

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// startCluster starts a server in the default (multi-process) model and
// returns it with its log.
func startClusterServer(t *testing.T, opts Options) (*Server, *strings.Builder) {
	t.Helper()
	var mu sync.Mutex
	var logs strings.Builder
	opts.Log = func(f string, a ...any) {
		mu.Lock()
		logs.WriteString(strings.TrimSpace(fmt.Sprintf(f, a...)) + "\n")
		mu.Unlock()
	}
	s, err := Start(context.Background(), opts)
	if err != nil {
		t.Fatalf("start: %v\n%s", err, logs.String())
	}
	t.Cleanup(func() { s.Close() })
	return s, &logs
}

func TestClusterBasic(t *testing.T) {
	s, logs := startClusterServer(t, Options{Database: "app"})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatalf("connect: %v\n%s", err, logs.String())
	}
	defer conn.Close(ctx)
	var pid int
	if err := conn.QueryRow(ctx, "select pg_backend_pid()").Scan(&pid); err != nil {
		t.Fatalf("query: %v\n%s", err, logs.String())
	}
	if _, err := conn.Exec(ctx, "create table t(id int primary key, v text); insert into t values (1,'a'),(2,'b')"); err != nil {
		t.Fatalf("ddl: %v\n%s", err, logs.String())
	}
	var n int
	if err := conn.QueryRow(ctx, "select count(*) from t").Scan(&n); err != nil || n != 2 {
		t.Fatalf("count: %v %d\n%s", err, n, logs.String())
	}
	// two sessions at once: the second is not blocked by the first's
	// open transaction
	c2, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatalf("connect 2: %v\n%s", err, logs.String())
	}
	defer c2.Close(ctx)
	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, "update t set v='x' where id=1"); err != nil {
		t.Fatal(err)
	}
	var pid2 int
	if err := c2.QueryRow(ctx, "select pg_backend_pid()").Scan(&pid2); err != nil {
		t.Fatalf("second session while first is in a transaction: %v\n%s", err, logs.String())
	}
	if pid2 == pid {
		t.Fatalf("both connections got backend %d", pid)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
}

// TestClusterDeadlock: two sessions lock rows crosswise; PostgreSQL's
// deadlock detector must end one of them with 40P01.
func TestClusterDeadlock(t *testing.T) {
	s, logs := startClusterServer(t, Options{Database: "app", Params: []string{"-c", "deadlock_timeout=200ms"}})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	c1, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatalf("connect: %v\n%s", err, logs.String())
	}
	defer c1.Close(ctx)
	c2, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close(ctx)
	if _, err := c1.Exec(ctx, "create table d(id int primary key); insert into d values (1),(2)"); err != nil {
		t.Fatal(err)
	}
	tx1, _ := c1.Begin(ctx)
	tx2, _ := c2.Begin(ctx)
	if _, err := tx1.Exec(ctx, "update d set id=id where id=1"); err != nil {
		t.Fatal(err)
	}
	if _, err := tx2.Exec(ctx, "update d set id=id where id=2"); err != nil {
		t.Fatal(err)
	}
	errs := make(chan error, 2)
	go func() { _, err := tx1.Exec(ctx, "update d set id=id where id=2"); errs <- err }()
	time.Sleep(100 * time.Millisecond)
	go func() { _, err := tx2.Exec(ctx, "update d set id=id where id=1"); errs <- err }()
	var got []error
	for i := 0; i < 2; i++ {
		got = append(got, <-errs)
	}
	deadlocks := 0
	for _, err := range got {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "40P01" {
			deadlocks++
		} else if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
	if deadlocks != 1 {
		t.Fatalf("want exactly one 40P01, got %v\n%s", got, logs.String())
	}
	tx1.Rollback(ctx)
	tx2.Rollback(ctx)
}

// TestClusterLockWait: a row lock held by one session makes another
// wait, then proceed when it is released (single mode could only time
// out here).
func TestClusterLockWait(t *testing.T) {
	s, logs := startClusterServer(t, Options{Database: "app"})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	c1, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatalf("connect: %v\n%s", err, logs.String())
	}
	defer c1.Close(ctx)
	c2, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close(ctx)
	if _, err := c1.Exec(ctx, "create table w(id int primary key, n int); insert into w values (1,0)"); err != nil {
		t.Fatal(err)
	}
	tx1, _ := c1.Begin(ctx)
	if _, err := tx1.Exec(ctx, "update w set n=n+1 where id=1"); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { _, err := c2.Exec(ctx, "update w set n=n+10 where id=1"); done <- err }()
	select {
	case err := <-done:
		t.Fatalf("second update did not wait for the lock: %v", err)
	case <-time.After(300 * time.Millisecond):
	}
	if err := tx1.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	if err := <-done; err != nil {
		t.Fatalf("second update: %v\n%s", err, logs.String())
	}
	var n int
	if err := c1.QueryRow(ctx, "select n from w where id=1").Scan(&n); err != nil || n != 11 {
		t.Fatalf("n=%d err=%v", n, err)
	}
}

// TestClusterTerminateBackend: a signal to an idle backend ends it.
func TestClusterTerminateBackend(t *testing.T) {
	s, logs := startClusterServer(t, Options{Database: "app"})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	victim, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatalf("connect: %v\n%s", err, logs.String())
	}
	defer victim.Close(ctx)
	var pid int
	if err := victim.QueryRow(ctx, "select pg_backend_pid()").Scan(&pid); err != nil {
		t.Fatal(err)
	}
	killer, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer killer.Close(ctx)
	start := time.Now()
	var ok bool
	if err := killer.QueryRow(ctx, "select pg_terminate_backend($1)", pid).Scan(&ok); err != nil || !ok {
		t.Fatalf("terminate: %v %v\n%s", ok, err, logs.String())
	}
	_, err = victim.Exec(ctx, "select 1")
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) || pgErr.Code != "57P01" {
		t.Fatalf("victim after terminate: %v, want 57P01 (took %s)\n%s", err, time.Since(start), logs.String())
	}
	t.Logf("terminated in %s", time.Since(start))
	deadline := time.Now().Add(5 * time.Second)
	for {
		var n int
		if err := killer.QueryRow(ctx, "select count(*) from pg_stat_activity where pid=$1", pid).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n == 0 {
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("backend %d still in pg_stat_activity\n%s", pid, logs.String())
		}
		time.Sleep(20 * time.Millisecond)
	}
}
