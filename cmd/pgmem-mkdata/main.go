// Command pgmem-mkdata runs initdb once inside the wasm module and writes
// the resulting data directory as internal/pgdata/pgdata.tar.zst, so that
// pgmem.Start can skip initdb at runtime.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/klauspost/compress/zstd"

	"github.com/shibukawa/pgmem/internal/engine"
	"github.com/shibukawa/pgmem/internal/wzr"
)

func main() {
	out := flag.String("o", "internal/pgdata/pgdata.tar.zst", "output file")
	initdbWasm := flag.String("initdb", "wasm/out/initdb.exnref.wasm", "initdb wasm module in the exnref encoding (built by wasm/build.sh)")
	flag.Parse()
	ctx := context.Background()
	wasm, err := os.ReadFile(*initdbWasm)
	if err != nil {
		log.Fatalf("read initdb module: %v (run ./wasm/build.sh first)", err)
	}
	rt, err := wzr.NewRuntime(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rt.Close(ctx)
	initdb, err := rt.Compile(ctx, wasm)
	if err != nil {
		log.Fatal(err)
	}
	e, err := engine.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	fs, err := e.BaseFS()
	if err != nil {
		log.Fatal(err)
	}
	if err := e.Initdb(fs, engine.InitdbOptions{Initdb: initdb}); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	zw, err := zstd.NewWriter(f, zstd.WithEncoderLevel(zstd.SpeedBestCompression), zstd.WithWindowSize(64<<20))
	if err != nil {
		log.Fatal(err)
	}
	if err := engine.Tar(fs, engine.PGData, zw); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	st, _ := os.Stat(*out)
	fmt.Printf("wrote %s (%d bytes)\n", *out, st.Size())
}
