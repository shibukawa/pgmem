// Package pgmemfixture provides driver helpers shared by embedded and
// subprocess test fixtures. It does not import the PostgreSQL engine.
package pgmemfixture

import (
	"context"
	"database/sql"
	"net"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// Endpoint is one database, reached through an embedded or network dialer.
type Endpoint interface {
	DSN() string
	Dial(context.Context, string, string) (net.Conn, error)
	Reset(context.Context) error
	Close() error
}

// TestDB holds one fork. All its handles use the same database.
type TestDB struct {
	t        testing.TB
	endpoint Endpoint
}

// New registers fork cleanup. Handles obtained later close before the fork.
func New(t testing.TB, endpoint Endpoint) *TestDB {
	t.Helper()
	t.Cleanup(func() {
		if err := endpoint.Close(); err != nil {
			t.Errorf("pgmem: close test database: %v", err)
		}
	})
	return &TestDB{t: t, endpoint: endpoint}
}

func (d *TestDB) DSN() string                     { return d.endpoint.DSN() }
func (d *TestDB) Reset(ctx context.Context) error { return d.endpoint.Reset(ctx) }

// Dial connects using the fixture's selected transport, for custom driver configs.
func (d *TestDB) Dial(ctx context.Context, network, address string) (net.Conn, error) {
	return d.endpoint.Dial(ctx, network, address)
}

// OpenDB opens an in-process, TCP or Unix socket connection as configured by endpoint.
func OpenDB(endpoint Endpoint, configure ...func(*pgx.ConnConfig)) (*sql.DB, error) {
	cfg, err := pgx.ParseConfig(endpoint.DSN())
	if err != nil {
		return nil, err
	}
	configureConnection(cfg, endpoint, configure)
	return stdlib.OpenDB(*cfg), nil
}

func (d *TestDB) DB(configure ...func(*pgx.ConnConfig)) *sql.DB {
	d.t.Helper()
	db, err := OpenDB(d.endpoint, configure...)
	if err != nil {
		d.t.Fatalf("pgmem: DB: %v", err)
	}
	d.t.Cleanup(func() {
		if err := db.Close(); err != nil {
			d.t.Errorf("pgmem: close DB: %v", err)
		}
	})
	return db
}

func (d *TestDB) PgxConn(configure ...func(*pgx.ConnConfig)) *pgx.Conn {
	d.t.Helper()
	cfg, err := pgx.ParseConfig(d.DSN())
	if err != nil {
		d.t.Fatal(err)
	}
	configureConnection(cfg, d.endpoint, configure)
	conn, err := pgx.ConnectConfig(d.t.Context(), cfg)
	if err != nil {
		d.t.Fatalf("pgmem: connect: %v", err)
	}
	d.t.Cleanup(func() {
		if err := conn.Close(context.Background()); err != nil {
			d.t.Errorf("pgmem: close connection: %v", err)
		}
	})
	return conn
}

func (d *TestDB) PgxPool(configure ...func(*pgxpool.Config)) *pgxpool.Pool {
	d.t.Helper()
	cfg, err := pgxpool.ParseConfig(d.DSN())
	if err != nil {
		d.t.Fatal(err)
	}
	identity := cfg.ConnConfig.Copy()
	for _, fn := range configure {
		if fn != nil {
			fn(cfg)
		}
	}
	if cfg.ConnConfig == nil {
		d.t.Fatal("pgmem: pool configuration must keep ConnConfig")
	}
	restoreConnectionIdentity(cfg.ConnConfig, identity, d.endpoint)
	pool, err := pgxpool.NewWithConfig(d.t.Context(), cfg)
	if err != nil {
		d.t.Fatalf("pgmem: pool: %v", err)
	}
	d.t.Cleanup(pool.Close)
	return pool
}

// ReadOnly returns a connection view with a guard against accidental persistent writes.
// This sets a session default, not a privilege boundary; callers can deliberately disable it.
func ReadOnly(endpoint Endpoint) Endpoint { return readOnlyEndpoint{endpoint} }

type readOnlyEndpoint struct{ Endpoint }

func (e readOnlyEndpoint) DSN() string {
	u, err := url.Parse(e.Endpoint.DSN())
	if err != nil {
		return e.Endpoint.DSN()
	}
	q := u.Query()
	options := q.Get("options")
	if options != "" {
		options += " "
	}
	q.Set("options", options+"-c default_transaction_read_only=on")
	u.RawQuery = strings.ReplaceAll(q.Encode(), "+", "%20")
	return u.String()
}

func configureConnection(cfg *pgx.ConnConfig, endpoint Endpoint, callbacks []func(*pgx.ConnConfig)) {
	identity := cfg.Copy()
	for _, fn := range callbacks {
		if fn != nil {
			fn(cfg)
		}
	}
	restoreConnectionIdentity(cfg, identity, endpoint)
}

// Application callbacks configure the driver; the fixture still chooses the database and transport.
func restoreConnectionIdentity(cfg, identity *pgx.ConnConfig, endpoint Endpoint) {
	cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password = identity.Host, identity.Port, identity.Database, identity.User, identity.Password
	cfg.Fallbacks = identity.Fallbacks
	cfg.DialFunc = endpoint.Dial
}
