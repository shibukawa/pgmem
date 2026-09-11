package engine

import (
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"testing"

	"github.com/shibukawa/pgmem/internal/guest"
	"github.com/shibukawa/pgmem/internal/wzr"
)

// initdbFactory loads wasm/out/initdb.wasm (a build output) under wazero;
// the test is skipped when it has not been built.
func initdbFactory(t *testing.T) guest.Factory {
	t.Helper()
	wasm, err := os.ReadFile("../../wasm/out/initdb.exnref.wasm")
	if err != nil {
		t.Skip("wasm/out/initdb.exnref.wasm not built")
	}
	rt, err := wzr.NewRuntime(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { rt.Close(context.Background()) })
	f, err := rt.Compile(context.Background(), wasm)
	if err != nil {
		t.Fatal(err)
	}
	return f
}

func TestParseCommand(t *testing.T) {
	argv, in, out := parseCommand(`"/pglite/bin/postgres" --check -F -c log_checkpoints=false < "/dev/null" > "/dev/null" 2>&1`)
	if len(argv) != 5 || argv[0] != "/pglite/bin/postgres" || argv[4] != "log_checkpoints=false" {
		t.Fatalf("argv = %q", argv)
	}
	if in != "/dev/null" || out != "/dev/null" {
		t.Fatalf("redirs = %q %q", in, out)
	}
}

func newEngine(t *testing.T) *Engine {
	t.Helper()
	e, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if os.Getenv("PGMEM_DEBUG") != "" {
		e.Log = t.Logf
	}
	t.Cleanup(func() { e.Close() })
	return e
}

func startupPacket(params map[string]string) []byte {
	var body bytes.Buffer
	binary.Write(&body, binary.BigEndian, uint32(196608)) // protocol 3.0
	for k, v := range params {
		body.WriteString(k)
		body.WriteByte(0)
		body.WriteString(v)
		body.WriteByte(0)
	}
	body.WriteByte(0)
	var pkt bytes.Buffer
	binary.Write(&pkt, binary.BigEndian, uint32(body.Len()+4))
	pkt.Write(body.Bytes())
	return pkt.Bytes()
}

func simpleQuery(q string) []byte {
	var pkt bytes.Buffer
	pkt.WriteByte('Q')
	binary.Write(&pkt, binary.BigEndian, uint32(len(q)+1+4))
	pkt.WriteString(q)
	pkt.WriteByte(0)
	return pkt.Bytes()
}

func TestInitdbAndQuery(t *testing.T) {
	e := newEngine(t)
	fs, err := e.BaseFS()
	if err != nil {
		t.Fatal(err)
	}
	if err := e.Initdb(fs, InitdbOptions{Initdb: initdbFactory(t)}); err != nil {
		t.Fatal(err)
	}
	if !fs.Exists(PGData + "/PG_VERSION") {
		t.Fatal("PG_VERSION missing after initdb")
	}
	b, err := e.Start(fs, StartOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	resp, err := b.Startup(startupPacket(map[string]string{"user": "postgres", "database": "postgres"}))
	if err != nil {
		t.Fatalf("startup: %v", err)
	}
	if len(resp) == 0 || resp[0] != 'R' {
		t.Fatalf("unexpected startup response %q\nlog: %s", resp, b.Log())
	}
	out, err := b.Exec(simpleQuery("SELECT 1+1 AS two, version()"))
	if err != nil {
		t.Fatalf("exec: %v", err)
	}
	if !bytes.Contains(out, []byte("PostgreSQL")) || out[0] != 'T' {
		t.Fatalf("unexpected response %q\nlog: %s", out, b.Log())
	}
	// An error must not kill the session.
	out, err = b.Exec(simpleQuery("SELECT * FROM no_such_table"))
	if err != nil {
		t.Fatalf("exec error query: %v", err)
	}
	if out[0] != 'E' {
		t.Fatalf("expected ErrorResponse, got %q", out)
	}
	out, err = b.Exec(simpleQuery("CREATE TABLE t(id int, name text); INSERT INTO t VALUES (1,'a'),(2,'b'); SELECT count(*) FROM t"))
	if err != nil {
		t.Fatalf("exec ddl: %v", err)
	}
	if !bytes.Contains(out, []byte("2")) || bytes.Contains(out, []byte{'E', 0, 0}) {
		t.Fatalf("unexpected ddl response %q\nlog: %s", out, b.Log())
	}
	t.Logf("server log:\n%s", b.Log())
}
