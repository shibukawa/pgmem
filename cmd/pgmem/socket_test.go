package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/shibukawa/pgmem"
)

type lockedBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (l *lockedBuffer) Write(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.b.Write(p)
}

func (l *lockedBuffer) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.b.String()
}

// startWithControl runs a controller whose template holds table t (1, 2)
// and that listens on a control socket; it returns the ready event.
func startWithControl(t *testing.T) map[string]any {
	t.Helper()
	s, err := pgmem.Start(context.Background(), pgmem.Options{Database: "app"})
	if err != nil {
		t.Fatal(err)
	}
	exec(t, s.DSN(), "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)")
	var out lockedBuffer
	c := newController(pgmem.Options{}, &out)
	tmpl := c.add("template", s, "postgres", "app")
	if err := c.listen("127.0.0.1:0"); err != nil {
		t.Fatal(err)
	}
	if err := c.ready(1, "test", tmpl); err != nil {
		t.Fatal(err)
	}
	inR, inW := io.Pipe()
	done := make(chan struct{})
	go func() { c.serve(inR); close(done) }()
	t.Cleanup(func() { inW.Close(); <-done })
	var ready map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(out.String())), &ready); err != nil {
		t.Fatal(err)
	}
	return ready
}

// sock is a sequential control-socket client.
type sock struct {
	t    *testing.T
	conn net.Conn
	sc   *bufio.Scanner
	seq  int
}

func dialControl(t *testing.T, ready map[string]any) *sock {
	t.Helper()
	conn, err := net.Dial("tcp", ready["control"].(map[string]any)["addr"].(string))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })
	return &sock{t: t, conn: conn, sc: bufio.NewScanner(conn)}
}

func hello(s *sock, ready map[string]any) map[string]any {
	s.t.Helper()
	return s.ok(map[string]any{"op": "hello", "token": ready["control"].(map[string]any)["token"]})
}

// call sends one request and returns its response, or nil when the server
// closed the connection instead of answering.
func (s *sock) call(req map[string]any) map[string]any {
	s.t.Helper()
	s.seq++
	req["id"] = s.seq
	b, _ := json.Marshal(req)
	if _, err := s.conn.Write(append(b, '\n')); err != nil {
		return nil
	}
	s.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	if !s.sc.Scan() {
		return nil
	}
	var res map[string]any
	if err := json.Unmarshal(s.sc.Bytes(), &res); err != nil {
		s.t.Fatal(err)
	}
	return res
}

func (s *sock) ok(req map[string]any) map[string]any {
	s.t.Helper()
	res := s.call(req)
	if res == nil || res["ok"] != true {
		s.t.Fatalf("%v -> %v", req, res)
	}
	return res
}

func errCode(res map[string]any) string {
	if res == nil || res["ok"] != false {
		return ""
	}
	return res["error"].(map[string]any)["code"].(string)
}

// refusedWithin reports whether dsn stops accepting connections within d.
func refusedWithin(dsn string, d time.Duration) bool {
	for deadline := time.Now().Add(d); time.Now().Before(deadline); time.Sleep(50 * time.Millisecond) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		conn, err := pgx.Connect(ctx, dsn)
		cancel()
		if err != nil {
			return true
		}
		conn.Close(context.Background())
	}
	return false
}

func TestControlSocketAuth(t *testing.T) {
	ready := startWithControl(t)
	ctl := ready["control"].(map[string]any)
	token := ctl["token"].(string)
	if len(token) != 32 || ctl["url"] != "pgmem-control://"+token+"@"+ctl["addr"].(string) {
		t.Fatalf("control = %v", ctl)
	}

	s := dialControl(t, ready)
	if code := errCode(s.call(map[string]any{"op": "snapshot"})); code != "unauthorized" {
		t.Fatalf("request before hello: code %q", code)
	}
	if res := s.call(map[string]any{"op": "hello", "token": token}); res != nil {
		t.Fatalf("connection survived a request before hello: %v", res)
	}

	s = dialControl(t, ready)
	if code := errCode(s.call(map[string]any{"op": "hello", "token": "nope"})); code != "unauthorized" {
		t.Fatalf("wrong token: code %q", code)
	}
	if res := s.call(map[string]any{"op": "hello", "token": token}); res != nil {
		t.Fatalf("connection survived a wrong token: %v", res)
	}

	s = dialControl(t, ready)
	if res := hello(s, ready); res["protocol"] != float64(protocolVersion) {
		t.Fatalf("hello -> %v", res)
	}
	if code := errCode(s.call(map[string]any{"op": "shutdown"})); code != "forbidden" {
		t.Fatalf("shutdown over the socket: code %q", code)
	}
}

// A worker forks and resets over the socket; its forks close when its
// connection does, and other workers' forks are not affected.
func TestControlSocketForkResetAndRelease(t *testing.T) {
	ready := startWithControl(t)
	w1 := dialControl(t, ready)
	hello(w1, ready)
	snap := w1.ok(map[string]any{"op": "snapshot", "server": "template"})["snapshot"].(string)
	f1 := w1.ok(map[string]any{"op": "fork", "snapshot": snap})
	exec(t, dsnOf(f1), "INSERT INTO t VALUES (3)")
	if got := count(t, dsnOf(f1)); got != 3 {
		t.Fatalf("fork count = %d", got)
	}
	w1.ok(map[string]any{"op": "reset", "server": idOf(f1)})
	if got := count(t, dsnOf(f1)); got != 2 {
		t.Fatalf("count after reset = %d", got)
	}

	w2 := dialControl(t, ready)
	hello(w2, ready)
	f2 := w2.ok(map[string]any{"op": "fork", "snapshot": snap})

	w1.conn.Close()
	if !refusedWithin(dsnOf(f1), 10*time.Second) {
		t.Fatal("fork still accepting connections after its control connection closed")
	}
	if got := count(t, dsnOf(f2)); got != 2 {
		t.Fatalf("other worker's fork: count = %d", got)
	}
}

func TestResetOp(t *testing.T) {
	h := newHarness(t)
	ctx := context.Background()
	tmpl := h.ready["dsn"].(string)
	exec(t, tmpl, "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)")
	snap := h.ok(map[string]any{"op": "snapshot"})["snapshot"].(string)
	f := h.ok(map[string]any{"op": "fork", "snapshot": snap})

	conn, err := pgx.Connect(ctx, dsnOf(f)) // stays connected across resets
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	if _, err := conn.Exec(ctx, "INSERT INTO t VALUES (3)"); err != nil {
		t.Fatal(err)
	}
	h.ok(map[string]any{"op": "reset", "server": idOf(f)})
	var n int
	if err := conn.QueryRow(ctx, "SELECT count(*) FROM t").Scan(&n); err != nil || n != 2 {
		t.Fatalf("count after reset = %d, %v", n, err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, "INSERT INTO t VALUES (4)"); err != nil {
		t.Fatal(err)
	}
	if code := errCode(h.call(map[string]any{"op": "reset", "server": idOf(f), "timeout_ms": 200})); code != "busy" {
		t.Fatalf("reset during a transaction: code %q", code)
	}
	if err := tx.Rollback(ctx); err != nil {
		t.Fatal(err)
	}

	exec(t, tmpl, "INSERT INTO t VALUES (9)")
	if code := errCode(h.call(map[string]any{"op": "reset", "server": "template"})); code != "protocol" {
		t.Fatalf("reset of a template without snapshot: code %q", code)
	}
	h.ok(map[string]any{"op": "reset", "server": "template", "snapshot": snap})
	if got := count(t, tmpl); got != 2 {
		t.Fatalf("template count after reset = %d", got)
	}
	if code := errCode(h.call(map[string]any{"op": "reset", "server": "zzz"})); code != "unknown_id" {
		t.Fatalf("unknown server: code %q", code)
	}
	if code := errCode(h.call(map[string]any{"op": "reset", "server": idOf(f), "snapshot": "zzz"})); code != "unknown_id" {
		t.Fatalf("unknown snapshot: code %q", code)
	}
}
