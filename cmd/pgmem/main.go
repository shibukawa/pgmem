// Command pgmem runs an in-memory PostgreSQL server as a standalone process,
// for test suites in other languages that start it as a subprocess.
//
// It prints one JSON line on stdout once the server is ready, for example
//
//	{"event":"ready","protocol":1,"version":"v0.1.0","pid":1234,"server":{"id":"template","host":"127.0.0.1","port":54321,...},"port":54321,"host":"127.0.0.1","user":"postgres","database":"app","dsn":"postgres://postgres@127.0.0.1:54321/app?sslmode=disable"}
//
// and then serves the control protocol: one JSON request per line on stdin
// (start, snapshot, fork, reset, close, shutdown), one JSON response per
// line on stdout, matched by the request's "id". With -control it also
// serves the protocol on a loopback socket, so that other processes (test
// workers) can fork and reset; the ready line then carries the socket's
// address and token. It keeps running until stdin is closed (the portable
// way for a parent process to end a child on every platform), a shutdown
// request arrives, or SIGINT/SIGTERM is received.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/shibukawa/pgmem"
)

func main() {
	var (
		port     = flag.Int("port", 0, "TCP port on 127.0.0.1 for the template server (0 = pick a free one)")
		database = flag.String("database", "postgres", "database to create and expose")
		user     = flag.String("user", "postgres", "superuser name")
		params   = flag.String("params", "", "extra postgres -c settings, comma separated (e.g. shared_buffers=32MB,log_statement=all)")
		verbose  = flag.Bool("log", false, "print the server log to stderr")
		noStdin  = flag.Bool("no-stdin", false, "do not read control requests from stdin and do not exit when it is closed")
		control  = flag.String("control", "", "also serve the control protocol on this loopback address (e.g. 127.0.0.1:0) for other processes")
	)
	// accepted for wrappers that still pass it: connections no longer share
	// a session, so there is no wait to bound
	flag.Duration("wait-timeout", 0, "deprecated; ignored")
	flag.Parse()

	base := pgmem.Options{}
	for _, p := range strings.Split(*params, ",") {
		if p = strings.TrimSpace(p); p != "" {
			base.Params = append(base.Params, "-c", p)
		}
	}
	if *verbose {
		base.Log = func(format string, args ...any) { fmt.Fprintf(os.Stderr, format+"\n", args...) }
	}
	opts := base
	opts.Port, opts.Database, opts.User = *port, *database, *user
	s, err := pgmem.Start(context.Background(), opts)
	if err != nil {
		log.Fatalf("pgmem: %v", err)
	}
	c := newController(base, os.Stdout)
	tmpl := c.add("template", s, *user, *database)
	if *control != "" {
		if err := c.listen(*control); err != nil {
			log.Fatalf("pgmem: control socket: %v", err)
		}
	}
	if err := c.ready(os.Getpid(), buildVersion(), tmpl); err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	if *noStdin {
		<-sig
		c.closeAll()
		return
	}
	done := make(chan struct{})
	go func() {
		c.serve(os.Stdin) // returns on EOF or shutdown, after closing everything
		close(done)
	}()
	select {
	case <-done:
	case <-sig:
		c.closeAll()
	}
}

// version is stamped by release builds (-ldflags "-X main.version=v0.1.0",
// scripts/build-binaries.sh). Other builds report the module version the go
// command recorded, which is a real version only under go install pkg@version.
var version string

func buildVersion() string {
	if version != "" {
		return version
	}
	if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "" {
		return bi.Main.Version
	}
	return "(devel)"
}
