package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// dockerTarget starts the official image with docker run, the way a
// Makefile or a CI service container does.
type dockerTarget struct {
	image       string
	fresh       bool
	vm          orbVM
	processBase float64
	id          string
	port        string
	n           int
}

func (d *dockerTarget) prepare(ctx context.Context) error {
	if !d.fresh {
		return nil
	}
	if err := d.vm.restart(ctx); err != nil {
		return err
	}
	var err error
	d.processBase, err = footprintMB(ctx, os.Getpid())
	return err
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

func (d *dockerTarget) mem(ctx context.Context) (memory, error) {
	m, err := containerMem(ctx, d.fresh, &d.vm, d.id)
	if err != nil || !d.fresh {
		return m, err
	}
	return addBenchmarkProcessGrowth(ctx, m, d.processBase)
}

// containerMem reads docker stats for the containers and, on a freshly
// restarted VM, records the VM's host-side net change separately.
func containerMem(ctx context.Context, fresh bool, vm *orbVM, ids ...string) (memory, error) {
	var m memory
	for _, id := range ids {
		c, err := dockerMemMB(ctx, id)
		if err != nil {
			return m, err
		}
		m.Container += c
	}
	m.Host = m.Container
	m.Service = m.Container
	if fresh {
		growth, total, err := vm.usage(ctx)
		if err != nil {
			return m, err
		}
		m.Host, m.VMTotal = growth, total
	}
	return m, nil
}

func (d *dockerTarget) stop(ctx context.Context) {
	if d.id != "" {
		run(ctx, "docker", "rm", "-f", d.id)
	}
}
