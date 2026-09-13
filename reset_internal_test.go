package pgmem

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
)

// Reset restarts the backend only when a client ran something since the
// last restore.
func TestResetSkipsRestartWhenNothingRan(t *testing.T) {
	ctx := context.Background()
	s, err := Start(ctx, Options{})
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	snap, err := s.Snapshot(ctx, SnapshotOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer snap.Close()
	f, err := snap.Fork(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	backend := func() any {
		f.bmu.Lock()
		defer f.bmu.Unlock()
		return f.b
	}

	b := backend()
	if err := f.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	if backend() != b {
		t.Fatal("Reset restarted a fork nobody used")
	}
	c, err := pgx.Connect(ctx, f.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)
	if backend() != b {
		t.Fatal("connecting alone counted as a change")
	}
	if _, err := c.Exec(ctx, `SELECT 1`); err != nil {
		t.Fatal(err)
	}
	if err := f.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	if backend() == b {
		t.Fatal("Reset after a statement kept the old backend")
	}
	b = backend()
	if err := f.Reset(ctx); err != nil {
		t.Fatal(err)
	}
	if backend() != b {
		t.Fatal("a second Reset in a row restarted again")
	}
}
