package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// devboxTarget runs the nix-built PostgreSQL that devbox installs. The
// devbox environment is resolved once before timing starts, as for a test
// suite running inside devbox shell; startup is initdb plus pg_ctl start
// of a fresh cluster.
type devboxTarget struct {
	dir     string
	env     []string
	dataDir string
	port    int
	n       int
}

func (d *devboxTarget) prepare(ctx context.Context) error {
	path, err := run(ctx, "devbox", "run", "-q", "-c", d.dir, "--", "printenv", "PATH")
	if err != nil {
		return err
	}
	lines := strings.Split(path, "\n")
	for _, kv := range os.Environ() {
		if !strings.HasPrefix(kv, "PATH=") {
			d.env = append(d.env, kv)
		}
	}
	d.env = append(d.env, "PATH="+lines[len(lines)-1])
	return nil
}

func (d *devboxTarget) start(ctx context.Context) error {
	var err error
	d.dataDir, err = os.MkdirTemp("", "pgbench-devbox-")
	if err != nil {
		return err
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	d.port = ln.Addr().(*net.TCPAddr).Port
	ln.Close()

	if err := d.cmd(ctx, "initdb", "-D", filepath.Join(d.dataDir, "data"), "-U", "postgres", "--auth=trust", "-N"); err != nil {
		return err
	}
	opts := fmt.Sprintf("-p %d -k %s -c listen_addresses=127.0.0.1", d.port, d.dataDir)
	if err := d.cmd(ctx, "pg_ctl", "-D", filepath.Join(d.dataDir, "data"), "-o", opts, "-l", filepath.Join(d.dataDir, "log"), "start"); err != nil {
		return err
	}
	if err := waitReady(ctx, d.dbURL("postgres")); err != nil {
		return err
	}
	return d.createApp(ctx)
}

func (d *devboxTarget) cmd(ctx context.Context, name string, args ...string) error {
	p := lookPath(d.env, name)
	if p == "" {
		return fmt.Errorf("%s not found in the devbox PATH", name)
	}
	c := exec.CommandContext(ctx, p, args...)
	c.Env = d.env
	out, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %v: %s", name, err, out)
	}
	return nil
}

func lookPath(env []string, name string) string {
	for _, kv := range env {
		if strings.HasPrefix(kv, "PATH=") {
			for _, dir := range filepath.SplitList(kv[5:]) {
				p := filepath.Join(dir, name)
				if st, err := os.Stat(p); err == nil && !st.IsDir() {
					return p
				}
			}
		}
	}
	return ""
}

func (d *devboxTarget) dbURL(db string) string {
	return fmt.Sprintf("postgres://postgres@127.0.0.1:%d/%s?sslmode=disable", d.port, db)
}

func (d *devboxTarget) url() string { return d.dbURL("app") }

func (d *devboxTarget) prepareIsolation(context.Context) error { return nil }

func (d *devboxTarget) isolate(ctx context.Context, fn func(string) error) error {
	return templateIsolation(ctx, d.dbURL, d.dbURL, &d.n, fn)
}

// mem sums the postmaster and every child it forked: footprint for the
// host figure, RSS for comparison (RSS counts shared buffers once per
// process that touched them).
func (d *devboxTarget) mem(ctx context.Context) (memory, error) {
	b, err := os.ReadFile(filepath.Join(d.dataDir, "data", "postmaster.pid"))
	if err != nil {
		return memory{}, err
	}
	pm, err := strconv.Atoi(strings.SplitN(string(b), "\n", 2)[0])
	if err != nil {
		return memory{}, err
	}
	out, err := run(ctx, "ps", "-axo", "pid=,ppid=")
	if err != nil {
		return memory{}, err
	}
	pids := []int{pm}
	for _, line := range strings.Split(out, "\n") {
		f := strings.Fields(line)
		if len(f) == 2 && f[1] == strconv.Itoa(pm) {
			if pid, err := strconv.Atoi(f[0]); err == nil {
				pids = append(pids, pid)
			}
		}
	}
	rss, err := rssMB(ctx, pids...)
	if err != nil {
		return memory{}, err
	}
	fp, err := footprintMB(ctx, pids...)
	return memory{Host: fp, RSS: rss, Process: fp, Service: fp}, err
}

func (d *devboxTarget) stop(ctx context.Context) {
	if d.dataDir == "" {
		return
	}
	d.cmd(ctx, "pg_ctl", "-D", filepath.Join(d.dataDir, "data"), "-m", "immediate", "stop")
	os.RemoveAll(d.dataDir)
}

// createApp creates the app database that docker and testcontainers get
// from POSTGRES_DB.
func (d *devboxTarget) createApp(ctx context.Context) error {
	c, err := pgx.Connect(ctx, d.dbURL("postgres"))
	if err != nil {
		return err
	}
	defer c.Close(ctx)
	_, err = c.Exec(ctx, "CREATE DATABASE app")
	return err
}
