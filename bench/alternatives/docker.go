package main

import (
	"context"
	"fmt"
	"strings"
)

// dockerTarget starts the official image with docker run, the way a
// Makefile or a CI service container does.
type dockerTarget struct {
	image string
	id    string
	port  string
	n     int
}

func (d *dockerTarget) start(ctx context.Context) error {
	id, err := run(ctx, "docker", "run", "-d", "--rm", "-p", "127.0.0.1::5432",
		"-e", "POSTGRES_HOST_AUTH_METHOD=trust", "-e", "POSTGRES_DB=app", d.image)
	if err != nil {
		return err
	}
	d.id = id
	out, err := run(ctx, "docker", "port", id, "5432/tcp")
	if err != nil {
		return err
	}
	line := strings.Split(out, "\n")[0]
	d.port = line[strings.LastIndex(line, ":")+1:]
	return waitReady(ctx, d.url())
}

func (d *dockerTarget) dbURL(db string) string {
	return fmt.Sprintf("postgres://postgres@127.0.0.1:%s/%s?sslmode=disable", d.port, db)
}

func (d *dockerTarget) url() string { return d.dbURL("app") }

func (d *dockerTarget) prepareIsolation(context.Context) error { return nil }

func (d *dockerTarget) isolate(ctx context.Context, fn func(string) error) error {
	return templateIsolation(ctx, d.dbURL, d.dbURL, &d.n, fn)
}

func (d *dockerTarget) memMB(ctx context.Context) (float64, error) { return dockerMemMB(ctx, d.id) }

func (d *dockerTarget) stop(ctx context.Context) {
	if d.id != "" {
		run(ctx, "docker", "rm", "-f", d.id)
	}
}
