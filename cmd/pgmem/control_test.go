package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/shibukawa/pgmem"
)

// harness runs a controller on pipes and matches responses by id.
type harness struct {
	t     *testing.T
	in    io.WriteCloser
	mu    sync.Mutex
	seq   int
	wait  map[int]chan map[string]any
	done  chan struct{}
	ready map[string]any
}

func newHarness(t *testing.T) *harness {
	t.Helper()
	s, err := pgmem.Start(context.Background(), pgmem.Options{Database: "app"})
	if err != nil {
		t.Fatal(err)
	}
	inR, inW := io.Pipe()
	outR, outW := io.Pipe()
	c := newController(pgmem.Options{}, outW)
	tmpl := c.add("template", s, "postgres", "app")
	h := &harness{t: t, in: inW, wait: map[int]chan map[string]any{}, done: make(chan struct{})}
	// io.Pipe writes block until read, so the reader must run before ready.
	readyCh := make(chan map[string]any, 1)
	go func() {
		sc := bufio.NewScanner(outR)
		sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
		if !sc.Scan() {
			close(readyCh)
			return
		}
		var ready map[string]any
		if err := json.Unmarshal(sc.Bytes(), &ready); err != nil {
			panic(err)
		}
		readyCh <- ready
		for sc.Scan() {
			var m map[string]any
			if err := json.Unmarshal(sc.Bytes(), &m); err != nil {
				panic(err)
			}
			id, _ := m["id"].(float64)
			h.mu.Lock()
			ch := h.wait[int(id)]
			delete(h.wait, int(id))
			h.mu.Unlock()
			if ch != nil {
				ch <- m
			}
		}
	}()
	if err := c.ready(1, "test", tmpl); err != nil {
		t.Fatal(err)
	}
	go func() { c.serve(inR); outW.Close(); close(h.done) }()
	var ok bool
	if h.ready, ok = <-readyCh; !ok {
		t.Fatal("no ready line")
	}
	t.Cleanup(func() { inW.Close(); <-h.done })
	return h
}

// send writes a request and returns a channel for its response.
func (h *harness) send(req map[string]any) chan map[string]any {
	h.mu.Lock()
	h.seq++
	id := h.seq
	ch := make(chan map[string]any, 1)
	h.wait[id] = ch
	h.mu.Unlock()
	req["id"] = id
	b, _ := json.Marshal(req)
	if _, err := h.in.Write(append(b, '\n')); err != nil {
		h.t.Fatal(err)
	}
	return ch
}

func (h *harness) call(req map[string]any) map[string]any {
	h.t.Helper()
	select {
	case res := <-h.send(req):
		return res
	case <-time.After(30 * time.Second):
		h.t.Fatalf("timeout waiting for %v", req)
		return nil
	}
}

func (h *harness) ok(req map[string]any) map[string]any {
	h.t.Helper()
	res := h.call(req)
	if res["ok"] != true {
		h.t.Fatalf("%v -> %v", req, res)
	}
	return res
}

func dsnOf(res map[string]any) string {
	return res["server"].(map[string]any)["dsn"].(string)
}

func idOf(res map[string]any) string {
	return res["server"].(map[string]any)["id"].(string)
}

func exec(t *testing.T, dsn, sql string) {
	t.Helper()
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(), sql); err != nil {
		t.Fatal(err)
	}
}

func count(t *testing.T, dsn string) int {
	t.Helper()
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(context.Background())
	var n int
	if err := conn.QueryRow(context.Background(), "SELECT count(*) FROM t").Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}

func TestReadyLine(t *testing.T) {
	h := newHarness(t)
	if h.ready["event"] != "ready" || h.ready["protocol"] != float64(protocolVersion) {
		t.Fatalf("ready = %v", h.ready)
	}
	// legacy flat fields and the nested endpoint agree
	srv := h.ready["server"].(map[string]any)
	if h.ready["dsn"] != srv["dsn"] || h.ready["port"] != srv["port"] || h.ready["database"] != "app" {
		t.Fatalf("ready = %v", h.ready)
	}
}

func TestSnapshotForkClose(t *testing.T) {
	h := newHarness(t)
	tmpl := h.ready["dsn"].(string)
	exec(t, tmpl, "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)")

	snap := h.ok(map[string]any{"op": "snapshot", "max_forks": 2})["snapshot"].(string)
	if !strings.HasPrefix(snap, "s") {
		t.Fatalf("snapshot id %q", snap)
	}
	f1 := h.ok(map[string]any{"op": "fork", "snapshot": snap})
	f2 := h.ok(map[string]any{"op": "fork", "snapshot": snap})
	if !strings.HasPrefix(idOf(f1), "f") || idOf(f1) == idOf(f2) {
		t.Fatalf("fork ids %q %q", idOf(f1), idOf(f2))
	}
	exec(t, dsnOf(f1), "INSERT INTO t VALUES (3)")
	if got := count(t, dsnOf(f1)); got != 3 {
		t.Fatalf("fork1 count = %d", got)
	}
	if got := count(t, dsnOf(f2)); got != 2 {
		t.Fatalf("fork2 count = %d, writes leaked", got)
	}
	if got := count(t, tmpl); got != 2 {
		t.Fatalf("template count = %d, writes leaked", got)
	}

	// pool is full: a third fork times out
	res := h.call(map[string]any{"op": "fork", "snapshot": snap, "timeout_ms": 200})
	if res["ok"] != false || res["error"].(map[string]any)["code"] != "pool_timeout" {
		t.Fatalf("expected pool_timeout, got %v", res)
	}
	// a blocked fork is released by a close that arrives later (out of order)
	pending := h.send(map[string]any{"op": "fork", "snapshot": snap})
	select {
	case res := <-pending:
		t.Fatalf("fork should block, got %v", res)
	case <-time.After(200 * time.Millisecond):
	}
	h.ok(map[string]any{"op": "close", "server": idOf(f1)})
	select {
	case res := <-pending:
		if res["ok"] != true {
			t.Fatalf("fork after close: %v", res)
		}
		if got := count(t, dsnOf(res)); got != 2 {
			t.Fatalf("fork3 count = %d", got)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("fork did not unblock after close")
	}
	// close is idempotent
	h.ok(map[string]any{"op": "close", "server": idOf(f1)})

	h.ok(map[string]any{"op": "close", "snapshot": snap})
	res = h.call(map[string]any{"op": "fork", "snapshot": snap})
	if res["ok"] != false || res["error"].(map[string]any)["code"] != "unknown_id" {
		t.Fatalf("fork on closed snapshot: %v", res)
	}
}

func TestSnapshotBusy(t *testing.T) {
	h := newHarness(t)
	conn, err := pgx.Connect(context.Background(), h.ready["dsn"].(string))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(context.Background())
	tx, err := conn.Begin(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	res := h.call(map[string]any{"op": "snapshot", "timeout_ms": 200})
	if res["ok"] != false || res["error"].(map[string]any)["code"] != "busy" {
		t.Fatalf("snapshot during open transaction: %v", res)
	}
	if err := tx.Rollback(context.Background()); err != nil {
		t.Fatal(err)
	}
	h.ok(map[string]any{"op": "snapshot", "timeout_ms": 5000})
}

func TestStartExtraTemplate(t *testing.T) {
	h := newHarness(t)
	res := h.ok(map[string]any{"op": "start", "database": "audit", "params": []string{"log_statement=all"}})
	ep := res["server"].(map[string]any)
	if ep["database"] != "audit" || !strings.HasPrefix(idOf(res), "t") {
		t.Fatalf("start -> %v", res)
	}
	exec(t, ep["dsn"].(string), "CREATE TABLE t (id int)")
	snap := h.ok(map[string]any{"op": "snapshot", "server": idOf(res)})["snapshot"].(string)
	f := h.ok(map[string]any{"op": "fork", "snapshot": snap})
	if got := count(t, dsnOf(f)); got != 0 {
		t.Fatalf("count = %d", got)
	}
	if !strings.Contains(dsnOf(f), "/audit?") {
		t.Fatalf("fork dsn %q should use the extra template's database", dsnOf(f))
	}
}

func TestErrors(t *testing.T) {
	h := newHarness(t)
	for _, tc := range []struct {
		req  map[string]any
		code string
	}{
		{map[string]any{"op": "nope"}, "unknown_op"},
		{map[string]any{}, "protocol"},
		{map[string]any{"op": "snapshot", "server": "zzz"}, "unknown_id"},
		{map[string]any{"op": "fork", "snapshot": "zzz"}, "unknown_id"},
		{map[string]any{"op": "close"}, "protocol"},
	} {
		res := h.call(tc.req)
		if res["ok"] != false || res["error"].(map[string]any)["code"] != tc.code {
			t.Errorf("%v -> %v, want %s", tc.req, res, tc.code)
		}
	}
}

func TestShutdownAndStdinEOF(t *testing.T) {
	h := newHarness(t)
	snap := h.ok(map[string]any{"op": "snapshot"})["snapshot"].(string)
	f := h.ok(map[string]any{"op": "fork", "snapshot": snap})
	h.ok(map[string]any{"op": "shutdown"})
	select {
	case <-h.done:
	case <-time.After(30 * time.Second):
		t.Fatal("serve did not return after shutdown")
	}
	if _, err := pgx.Connect(context.Background(), dsnOf(f)); err == nil {
		t.Fatal("fork still accepting connections after shutdown")
	}

	// stdin EOF alone also tears everything down
	h2 := newHarness(t)
	snap = h2.ok(map[string]any{"op": "snapshot"})["snapshot"].(string)
	f = h2.ok(map[string]any{"op": "fork", "snapshot": snap})
	h2.in.Close()
	select {
	case <-h2.done:
	case <-time.After(30 * time.Second):
		t.Fatal("serve did not return after EOF")
	}
	if _, err := pgx.Connect(context.Background(), dsnOf(f)); err == nil {
		t.Fatal("fork still accepting connections after EOF")
	}
	_ = fmt.Sprint
}
