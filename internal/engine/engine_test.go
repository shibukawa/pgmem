package engine

import (
	"context"
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
	// a standalone child, the way the setup step creates databases
	if err := e.ExecStandalone(fs, "postgres", "CREATE TABLE t(id int, name text);\n\nINSERT INTO t VALUES (1,'a'),(2,'b');\n"); err != nil {
		t.Fatalf("standalone: %v", err)
	}
	// the postmaster and its processes come up on the same directory
	cl, err := e.StartCluster(context.Background(), fs, StartOptions{}, nil)
	if err != nil {
		t.Fatalf("cluster: %v", err)
	}
	if cl.Dead() {
		t.Fatal("postmaster exited")
	}
	cl.Kill()
	if !cl.Dead() {
		t.Fatalf("postmaster still alive after Kill\n%s", cl.Log())
	}
}
