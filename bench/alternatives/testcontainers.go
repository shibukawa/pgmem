package main

import (
	"context"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Snapshot and Restore open database/sql with this driver
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// tcTarget uses the testcontainers-go postgres module with the wait
// strategy its documentation recommends. A fresh process also starts the
// Ryuk reaper container, which is part of what a test binary pays.
type tcTarget struct {
	image       string
	fresh       bool
	vm          orbVM
	processBase float64
	c           *postgres.PostgresContainer
	dsn         string
}

func (t *tcTarget) prepare(ctx context.Context) error {
	if !t.fresh {
		return nil
	}
	if err := t.vm.restart(ctx); err != nil {
		return err
	}
	var err error
	t.processBase, err = footprintMB(ctx, os.Getpid())
	return err
}

func (t *tcTarget) start(ctx context.Context) error {
	var err error
	t.c, err = postgres.Run(ctx, t.image,
		postgres.WithDatabase("app"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithSQLDriver("pgx"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(time.Minute)),
		// postgres:18 moved PGDATA from /var/lib/postgresql/data to
		// /var/lib/postgresql/18/docker. Keep database I/O in memory for this
		// benchmark by mounting the actual data directory as tmpfs.
		testcontainers.WithTmpfs(map[string]string{
			"/var/lib/postgresql/18/docker": "rw",
		}),
	)
	if err != nil {
		return err
	}
	t.dsn, err = t.c.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return err
	}
	return waitReady(ctx, t.dsn)
}

func (t *tcTarget) url() string { return t.dsn }

// prepareIsolation uses the module's own Snapshot, which copies the
// database into a template; Restore recreates it before each test.
func (t *tcTarget) prepareIsolation(ctx context.Context) error {
	return t.c.Snapshot(ctx)
}

func (t *tcTarget) isolate(ctx context.Context, fn func(string) error) error {
	if err := t.c.Restore(ctx); err != nil {
		return err
	}
	return fn(t.dsn)
}

// mem includes the Ryuk reaper container that testcontainers starts.
func (t *tcTarget) mem(ctx context.Context) (memory, error) {
	ids := []string{t.c.GetContainerID()}
	if ryuk, err := run(ctx, "docker", "ps", "-q", "--filter", "label=org.testcontainers.ryuk=true"); err == nil && ryuk != "" {
		ids = append(ids, strings.Fields(ryuk)...)
	}
	m, err := containerMem(ctx, t.fresh, &t.vm, ids...)
	if err != nil || !t.fresh {
		return m, err
	}
	return addBenchmarkProcessGrowth(ctx, m, t.processBase)
}

func (t *tcTarget) stop(ctx context.Context) {
	if t.c != nil {
		testcontainers.TerminateContainer(t.c)
	}
}
