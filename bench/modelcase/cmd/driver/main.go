// Command driver runs the model-case test binary on one execution model
// and prints one JSON line per run: the server boot time it measured, the
// test binary's wall-clock time, and the trace the binary wrote.
//
//	driver -backend pgmem|devbox|docker -runs 5 -testbin bin/store.test
//
// pgmem boots inside the test binary, so there is no external boot span.
// devbox starts a nix-installed PostgreSQL on an initialized cluster with
// pg_ctl, docker starts the official image; both are timed until SELECT 1
// succeeds on the app database, then the test binary runs against it.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/shibukawa/pgmem/bench/modelcase/internal/trace"
)

type Result struct {
	Model     string       `json:"model"`
	Run       int          `json:"run"`
	Parallel  int          `json:"parallel"`
	BootMS    float64      `json:"boot_ms"`
	ProcessMS float64      `json:"process_ms"`
	TotalMS   float64      `json:"total_ms"`
	Version   string       `json:"version"`
	Spans     []trace.Span `json:"spans"`
	Time      time.Time    `json:"time"`
}

type server interface {
	// start boots the server and returns the DSN of the empty app database.
	start(ctx context.Context) (string, error)
	stop(ctx context.Context)
}

func main() {
	backend := flag.String("backend", "pgmem", "pgmem | devbox | docker")
	runs := flag.Int("runs", 5, "how many times to run the suite")
	warmup := flag.Int("warmup", 1, "untimed runs first: the first execution of a fresh binary pays for page-in and code signing")
	testbin := flag.String("testbin", "bin/store.test", "the compiled test binary (go test -c ./store)")
	parallel := flag.Int("parallel", runtime.GOMAXPROCS(0), "-test.parallel for the binary")
	image := flag.String("image", "postgres:18-alpine", "image for docker")
	devboxDir := flag.String("devbox", "../alternatives/devbox", "directory with devbox.json")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	var s server
	switch *backend {
	case "pgmem":
		s = nil
	case "devbox":
		d := &devboxServer{dir: *devboxDir}
		if err := d.prepare(ctx); err != nil {
			fatal(err)
		}
		defer d.cleanup()
		s = d
	case "docker":
		s = &dockerServer{image: *image}
	default:
		fatal(fmt.Errorf("unknown backend %q", *backend))
	}
	enc := json.NewEncoder(os.Stdout)
	for i := 1; i <= *warmup; i++ {
		fmt.Fprintf(os.Stderr, "warm-up %d %s\n", i, *backend)
		if _, err := once(ctx, *backend, s, *testbin, *parallel); err != nil {
			fatal(err)
		}
	}
	for i := 1; i <= *runs; i++ {
		fmt.Fprintf(os.Stderr, "run %d %s\n", i, *backend)
		r, err := once(ctx, *backend, s, *testbin, *parallel)
		if err != nil {
			fatal(err)
		}
		r.Run = i
		enc.Encode(r)
	}
}

func once(ctx context.Context, model string, s server, testbin string, parallel int) (*Result, error) {
	r := &Result{Model: model, Parallel: parallel, Time: time.Now().UTC()}
	var dsn string
	if s != nil {
		begin := time.Now()
		var err error
		dsn, err = s.start(ctx)
		if err != nil {
			return nil, fmt.Errorf("start: %w", err)
		}
		defer s.stop(context.Background())
		r.BootMS = ms(time.Since(begin))
	}
	dir, err := os.MkdirTemp("", "modelcase-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	tracePath := filepath.Join(dir, "trace.json")

	cmd := exec.CommandContext(ctx, testbin, "-test.parallel", strconv.Itoa(parallel))
	cmd.Env = append(os.Environ(), "MODELCASE_TRACE="+tracePath)
	if dsn != "" {
		cmd.Env = append(cmd.Env, "MODELCASE_DSN="+dsn)
	}
	begin := time.Now()
	out, err := cmd.CombinedOutput()
	r.ProcessMS = ms(time.Since(begin))
	if err != nil {
		return nil, fmt.Errorf("%s: %v\n%s", testbin, err, out)
	}
	b, err := os.ReadFile(tracePath)
	if err != nil {
		return nil, err
	}
	var t trace.Trace
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	// The trace's clock starts at TestMain; put the external boot before
	// it and the binary's own startup after it.
	for i := range t.Spans {
		t.Spans[i].StartMS += r.BootMS
		t.Spans[i].EndMS += r.BootMS
	}
	if r.BootMS > 0 {
		t.Spans = append([]trace.Span{{Name: "boot", Kind: "boot", Lane: "setup", StartMS: 0, EndMS: r.BootMS}}, t.Spans...)
	}
	r.Spans = t.Spans
	r.Version = t.Meta["server_version"]
	r.TotalMS = r.BootMS + r.ProcessMS
	return r, nil
}

// dockerServer: docker run the official image, the way a CI service
// container or a Makefile target does.
type dockerServer struct {
	image string
	id    string
}

func (d *dockerServer) start(ctx context.Context) (string, error) {
	id, err := run(ctx, "docker", "run", "-d", "--rm", "-p", "127.0.0.1::5432",
		"-e", "POSTGRES_HOST_AUTH_METHOD=trust", "-e", "POSTGRES_DB=app", d.image)
	if err != nil {
		return "", err
	}
	d.id = id
	out, err := run(ctx, "docker", "port", id, "5432/tcp")
	if err != nil {
		return "", err
	}
	line := strings.Split(out, "\n")[0]
	port := line[strings.LastIndex(line, ":")+1:]
	dsn := fmt.Sprintf("postgres://postgres@127.0.0.1:%s/app?sslmode=disable", port)
	return dsn, waitReady(ctx, dsn)
}

func (d *dockerServer) stop(ctx context.Context) {
	if d.id != "" {
		run(ctx, "docker", "rm", "-f", d.id)
		d.id = ""
	}
}

// devboxServer: the nix-installed PostgreSQL that devbox provides, run
// directly with pg_ctl on a cluster initialized once outside the timing.
// Boot is pg_ctl start plus CREATE DATABASE app, until SELECT 1 succeeds.
type devboxServer struct {
	dir     string
	bin     string
	dataDir string
	port    int
}

func (d *devboxServer) prepare(ctx context.Context) error {
	out, err := run(ctx, "devbox", "run", "-q", "-c", d.dir, "--", "sh", "-c", "command -v pg_ctl")
	if err != nil {
		return err
	}
	d.bin = filepath.Dir(strings.TrimSpace(out))
	d.dataDir, err = os.MkdirTemp("", "modelcase-devbox-")
	if err != nil {
		return err
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	d.port = ln.Addr().(*net.TCPAddr).Port
	ln.Close()
	_, err = run(ctx, filepath.Join(d.bin, "initdb"), "-D", d.pgData(), "-U", "postgres", "--auth=trust", "-N")
	return err
}

func (d *devboxServer) pgData() string { return filepath.Join(d.dataDir, "data") }

func (d *devboxServer) dbURL(db string) string {
	return fmt.Sprintf("postgres://postgres@127.0.0.1:%d/%s?sslmode=disable", d.port, db)
}

func (d *devboxServer) start(ctx context.Context) (string, error) {
	opts := fmt.Sprintf("-p %d -k %s -c listen_addresses=127.0.0.1", d.port, d.dataDir)
	if _, err := run(ctx, filepath.Join(d.bin, "pg_ctl"), "-D", d.pgData(), "-l", filepath.Join(d.dataDir, "log"), "-o", opts, "-w", "-s", "start"); err != nil {
		return "", err
	}
	if err := waitReady(ctx, d.dbURL("postgres")); err != nil {
		return "", err
	}
	c, err := pgx.Connect(ctx, d.dbURL("postgres"))
	if err != nil {
		return "", err
	}
	_, err = c.Exec(ctx, "CREATE DATABASE app")
	c.Close(ctx)
	if err != nil {
		return "", err
	}
	return d.dbURL("app"), waitReady(ctx, d.dbURL("app"))
}

func (d *devboxServer) stop(ctx context.Context) {
	if c, err := pgx.Connect(ctx, d.dbURL("postgres")); err == nil {
		c.Exec(ctx, "DROP DATABASE app")
		c.Close(ctx)
	}
	run(ctx, filepath.Join(d.bin, "pg_ctl"), "-D", d.pgData(), "-m", "fast", "-w", "-s", "stop")
}

func (d *devboxServer) cleanup() {
	if d.dataDir != "" {
		os.RemoveAll(d.dataDir)
	}
}

// waitReady polls url until SELECT 1 succeeds.
func waitReady(ctx context.Context, url string) error {
	var last error
	for {
		c, err := pgx.Connect(ctx, url)
		if err == nil {
			var one int
			err = c.QueryRow(ctx, "SELECT 1").Scan(&one)
			c.Close(ctx)
			if err == nil {
				return nil
			}
		}
		last = err
		select {
		case <-ctx.Done():
			return fmt.Errorf("server not ready: %w (last error %v)", ctx.Err(), last)
		case <-time.After(10 * time.Millisecond):
		}
	}
}

func run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("%s %s: %v: %s", name, strings.Join(args, " "), err, ee.Stderr)
		}
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func ms(d time.Duration) float64 { return float64(d.Microseconds()) / 1000 }

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
