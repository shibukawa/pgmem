package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

// devboxTarget measures the built-in PostgreSQL service workflow. prepare
// resolves the Devbox environment and prepares a persistent cluster before
// timing; start measures devbox services up -b through a successful SQL probe.
type devboxTarget struct {
	dir     string
	dataDir string
	port    int
	n       int
}

func (d *devboxTarget) prepare(ctx context.Context) error {
	// Resolve/install the environment outside the startup interval.
	if _, err := run(ctx, "devbox", "run", "-q", "-c", d.dir, "--", "true"); err != nil {
		return err
	}

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

	if err := d.initDB(ctx); err != nil {
		return err
	}
	// Devbox services expect an initialized persistent PGDATA. Create the
	// benchmark database in an untimed service launch, then stop it so start
	// measures a real service restart with the database already in place.
	if err := d.servicesUp(ctx); err != nil {
		return err
	}
	if err := waitReady(ctx, d.dbURL("postgres")); err != nil {
		return err
	}
	if err := d.createApp(ctx); err != nil {
		return err
	}
	return d.servicesStop(ctx)
}

func (d *devboxTarget) start(ctx context.Context) error {
	if err := d.servicesUp(ctx); err != nil {
		return err
	}
	return waitReady(ctx, d.url())
}

func (d *devboxTarget) initDB(ctx context.Context) error {
	args := []string{
		"run", "-q", "-c", d.dir,
		"-e", "PGDATA=" + d.pgData(),
		"-e", "PGHOST=" + d.dataDir,
		"-e", fmt.Sprintf("PGPORT=%d", d.port),
		"--", "initdb", "-D", d.pgData(), "-U", "postgres", "--auth=trust", "-N",
	}
	_, err := run(ctx, "devbox", args...)
	return err
}

func (d *devboxTarget) serviceArgs(action string) []string {
	args := []string{
		"services", action, "-q", "-c", d.dir,
		"-e", "PGDATA=" + d.pgData(),
		"-e", "PGHOST=" + d.dataDir,
		"-e", fmt.Sprintf("PGPORT=%d", d.port),
	}
	if action == "up" {
		args = append(args, "-b")
	}
	return args
}

func (d *devboxTarget) servicesUp(ctx context.Context) error {
	_, err := run(ctx, "devbox", d.serviceArgs("up")...)
	return err
}

func (d *devboxTarget) servicesStop(ctx context.Context) error {
	_, err := run(ctx, "devbox", d.serviceArgs("stop")...)
	return err
}

func (d *devboxTarget) pgData() string {
	return filepath.Join(d.dataDir, "data")
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
	b, err := os.ReadFile(filepath.Join(d.pgData(), "postmaster.pid"))
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
	_ = d.servicesStop(ctx)
	_ = os.RemoveAll(d.dataDir)
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
