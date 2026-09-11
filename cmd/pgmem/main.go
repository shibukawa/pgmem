// Command pgmem runs an in-memory PostgreSQL server as a standalone process,
// for test suites in other languages that start it as a subprocess.
//
// It prints one JSON line on stdout once the server is ready, for example
//
//	{"port":54321,"host":"127.0.0.1","user":"postgres","database":"app","dsn":"postgres://postgres@127.0.0.1:54321/app?sslmode=disable","pid":1234}
//
// and keeps running until stdin is closed (the portable way for a parent
// process to end a child on every platform) or SIGINT/SIGTERM arrives.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shibukawa/pgmem"
)

func main() {
	var (
		port     = flag.Int("port", 0, "TCP port on 127.0.0.1 (0 = pick a free one)")
		database = flag.String("database", "postgres", "database to create and expose")
		user     = flag.String("user", "postgres", "superuser name")
		params   = flag.String("params", "", "extra postgres -c settings, comma separated (e.g. shared_buffers=32MB,log_statement=all)")
		verbose  = flag.Bool("log", false, "print the server log to stderr")
		noStdin  = flag.Bool("no-stdin", false, "do not exit when stdin is closed")
	)
	flag.Parse()

	opts := pgmem.Options{Port: *port, Database: *database, User: *user}
	for _, p := range strings.Split(*params, ",") {
		if p = strings.TrimSpace(p); p != "" {
			opts.Params = append(opts.Params, "-c", p)
		}
	}
	if *verbose {
		opts.Log = func(format string, args ...any) { fmt.Fprintf(os.Stderr, format+"\n", args...) }
	}
	s, err := pgmem.Start(context.Background(), opts)
	if err != nil {
		log.Fatalf("pgmem: %v", err)
	}
	ready := map[string]any{
		"host": "127.0.0.1", "port": s.Port(), "user": *user, "database": *database,
		"dsn": s.DSN(), "pid": os.Getpid(),
	}
	if err := json.NewEncoder(os.Stdout).Encode(ready); err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{}, 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() { <-sig; done <- struct{}{} }()
	if !*noStdin {
		go func() {
			io.Copy(io.Discard, os.Stdin) // returns when the parent closes our stdin
			done <- struct{}{}
		}()
	}
	<-done
	s.Close()
}
