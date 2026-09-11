package pgmem_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/shibukawa/pgmem"
)

func startServer(t *testing.T, opts pgmem.Options) *pgmem.Server {
	t.Helper()
	if os.Getenv("PGMEM_DEBUG") != "" {
		opts.Log = t.Logf
	}
	start := time.Now()
	s, err := pgmem.Start(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("server up in %s at %s", time.Since(start), s.DSN())
	t.Cleanup(func() { s.Close() })
	return s
}

func TestPgx(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	var version string
	if err := conn.QueryRow(ctx, "SELECT version()").Scan(&version); err != nil {
		t.Fatal(err)
	}
	t.Log(version)

	if _, err := conn.Exec(ctx, `CREATE TABLE users (id serial PRIMARY KEY, name text NOT NULL, created timestamptz DEFAULT now())`); err != nil {
		t.Fatal(err)
	}
	for _, n := range []string{"alice", "bob", "carol"} {
		if _, err := conn.Exec(ctx, `INSERT INTO users(name) VALUES ($1)`, n); err != nil {
			t.Fatal(err)
		}
	}
	rows, err := conn.Query(ctx, `SELECT id, name FROM users WHERE name <> $1 ORDER BY id`, "bob")
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			t.Fatal(err)
		}
		got = append(got, name)
	}
	rows.Close()
	if len(got) != 2 || got[0] != "alice" || got[1] != "carol" {
		t.Fatalf("got %v", got)
	}

	// Errors are reported and the session survives.
	if _, err := conn.Exec(ctx, `SELECT * FROM missing`); err == nil {
		t.Fatal("expected error")
	}
	var n int
	if err := conn.QueryRow(ctx, `SELECT count(*) FROM users`).Scan(&n); err != nil || n != 3 {
		t.Fatalf("count: %d %v", n, err)
	}

	// Transactions.
	tx, err := conn.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `DELETE FROM users`); err != nil {
		t.Fatal(err)
	}
	if err := tx.Rollback(ctx); err != nil {
		t.Fatal(err)
	}
	if err := conn.QueryRow(ctx, `SELECT count(*) FROM users`).Scan(&n); err != nil || n != 3 {
		t.Fatalf("count after rollback: %d %v", n, err)
	}

	// plpgsql (statically linked module).
	if _, err := conn.Exec(ctx, `CREATE FUNCTION add_one(x int) RETURNS int LANGUAGE plpgsql AS $$ BEGIN RETURN x + 1; END $$`); err != nil {
		t.Fatal(err)
	}
	if err := conn.QueryRow(ctx, `SELECT add_one(41)`).Scan(&n); err != nil || n != 42 {
		t.Fatalf("plpgsql: %d %v", n, err)
	}

	// JSON and a few types.
	var j map[string]any
	var ts time.Time
	if err := conn.QueryRow(ctx, `SELECT '{"a":1}'::jsonb, now()`).Scan(&j, &ts); err != nil {
		t.Fatal(err)
	}
	if j["a"] != float64(1) || ts.IsZero() {
		t.Fatalf("json/ts: %v %v", j, ts)
	}
}

func TestSecondConnectionGetsFreshSession(t *testing.T) {
	s := startServer(t, pgmem.Options{Database: "app", User: "tester"})
	ctx := context.Background()
	c1, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c1.Exec(ctx, `CREATE TEMP TABLE tt(x int); SET application_name = 'one'`); err != nil {
		t.Fatal(err)
	}
	var db, user string
	if err := c1.QueryRow(ctx, `SELECT current_database(), current_user`).Scan(&db, &user); err != nil || db != "app" || user != "tester" {
		t.Fatalf("db = %q user = %q %v", db, user, err)
	}
	c1.Close(ctx)

	c2, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer c2.Close(ctx)
	if _, err := c2.Exec(ctx, `SELECT * FROM tt`); err == nil {
		t.Fatal("temp table leaked into the next session")
	}
	var app string
	if err := c2.QueryRow(ctx, `SHOW application_name`).Scan(&app); err != nil {
		t.Fatal(err)
	}
	if app == "one" {
		t.Fatal("SET leaked into the next session")
	}
}

func TestDatabaseSQL(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	db, err := sql.Open("pgx", s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)
	if _, err := db.Exec(`CREATE TABLE kv(k text PRIMARY KEY, v int)`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO kv VALUES ($1, $2), ($3, $4)`, "a", 1, "b", 2); err != nil {
		t.Fatal(err)
	}
	var sum int
	if err := db.QueryRow(`SELECT sum(v) FROM kv`).Scan(&sum); err != nil || sum != 3 {
		t.Fatalf("sum = %d, %v", sum, err)
	}
	// Full text search uses the statically linked snowball dictionary.
	var ok bool
	if err := db.QueryRow(`SELECT to_tsvector('english', 'running quickly') @@ to_tsquery('english', 'run')`).Scan(&ok); err != nil || !ok {
		t.Fatalf("tsearch: %v %v", ok, err)
	}
}

func BenchmarkSimpleQueries(b *testing.B) {
	s, err := pgmem.Start(context.Background(), pgmem.Options{})
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close(ctx)
	if _, err := conn.Exec(ctx, `CREATE TABLE bench(id serial PRIMARY KEY, v text); INSERT INTO bench(v) SELECT md5(i::text) FROM generate_series(1, 1000) i`); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v string
		if err := conn.QueryRow(ctx, `SELECT v FROM bench WHERE id = $1`, i%1000+1).Scan(&v); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSimpleQueriesInProcess(b *testing.B) {
	s, err := pgmem.Start(context.Background(), pgmem.Options{})
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	ctx := context.Background()
	cfg, err := pgx.ParseConfig(s.DSN())
	if err != nil {
		b.Fatal(err)
	}
	cfg.DialFunc = s.Dial
	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close(ctx)
	if _, err := conn.Exec(ctx, `CREATE TABLE bench(id serial PRIMARY KEY, v text); INSERT INTO bench(v) SELECT md5(i::text) FROM generate_series(1, 1000) i`); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v string
		if err := conn.QueryRow(ctx, `SELECT v FROM bench WHERE id = $1`, i%1000+1).Scan(&v); err != nil {
			b.Fatal(err)
		}
	}
}

func TestDialInProcess(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	cfg, err := pgx.ParseConfig(s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	cfg.DialFunc = s.Dial
	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	var n int
	if err := conn.QueryRow(ctx, `SELECT 40 + 2`).Scan(&n); err != nil || n != 42 {
		t.Fatalf("%d %v", n, err)
	}
}

func BenchmarkCPUHeavyQuery(b *testing.B) {
	s, err := pgmem.Start(context.Background(), pgmem.Options{})
	if err != nil {
		b.Fatal(err)
	}
	defer s.Close()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close(ctx)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var n int64
		// sort + aggregate over 200k generated rows: almost all time is inside the guest
		if err := conn.QueryRow(ctx, `SELECT count(*) FROM (SELECT md5(i::text) AS h FROM generate_series(1, 200000) i ORDER BY h) t`).Scan(&n); err != nil {
			b.Fatal(err)
		}
	}
}

func TestUUID(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	var u4, u7 [16]byte
	var v4, v7 int
	if err := conn.QueryRow(ctx, `SELECT gen_random_uuid(), uuidv7(), uuid_extract_version(gen_random_uuid()), uuid_extract_version(uuidv7())`).Scan(&u4, &u7, &v4, &v7); err != nil {
		t.Fatal(err)
	}
	if v4 != 4 || v7 != 7 || u4 == u7 {
		t.Fatalf("uuid versions %d %d, %x %x", v4, v7, u4, u7)
	}
	if _, err := conn.Exec(ctx, `CREATE TABLE u(id uuid PRIMARY KEY DEFAULT uuidv7(), n int); INSERT INTO u(n) VALUES (1),(2)`); err != nil {
		t.Fatal(err)
	}
	var n int
	if err := conn.QueryRow(ctx, `SELECT count(DISTINCT id) FROM u`).Scan(&n); err != nil || n != 2 {
		t.Fatalf("%d %v", n, err)
	}
}

func TestSleepAndStatementTimeout(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	// pg_sleep must actually wait (the host sleeps instead of spinning).
	start := time.Now()
	if _, err := conn.Exec(ctx, `SELECT pg_sleep(0.3)`); err != nil {
		t.Fatal(err)
	}
	if el := time.Since(start); el < 300*time.Millisecond || el > 2*time.Second {
		t.Fatalf("pg_sleep(0.3) took %s", el)
	}

	// statement_timeout must interrupt a running pg_sleep, not wait for
	// the next protocol message.
	if _, err := conn.Exec(ctx, `SET statement_timeout = '200ms'`); err != nil {
		t.Fatal(err)
	}
	start = time.Now()
	_, err = conn.Exec(ctx, `SELECT pg_sleep(5)`)
	el := time.Since(start)
	if err == nil {
		t.Fatal("expected statement timeout error")
	}
	if el > 2*time.Second {
		t.Fatalf("timeout fired after %s", el)
	}
	if _, err := conn.Exec(ctx, `RESET statement_timeout`); err != nil {
		t.Fatal(err)
	}
	// The session is still usable afterwards.
	var one int
	if err := conn.QueryRow(ctx, `SELECT 1`).Scan(&one); err != nil || one != 1 {
		t.Fatalf("after timeout: %d %v", one, err)
	}
}

func TestCryptoHashes(t *testing.T) {
	s := startServer(t, pgmem.Options{})
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)
	var md5s, sha224s, sha256s, sha384s, sha512s string
	err = conn.QueryRow(ctx, `SELECT md5('abc'), encode(sha224('abc'), 'hex'), encode(sha256('abc'), 'hex'), encode(sha384('abc'), 'hex'), encode(sha512('abc'), 'hex')`).Scan(&md5s, &sha224s, &sha256s, &sha384s, &sha512s)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		md5s:    "900150983cd24fb0d6963f7d28e17f72",
		sha224s: "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7",
		sha256s: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		sha384s: "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7",
		sha512s: "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f",
	}
	for got, exp := range want {
		if got != exp {
			t.Fatalf("hash mismatch: got %s want %s", got, exp)
		}
	}
	// A large input goes through several update() calls.
	var big string
	if err := conn.QueryRow(ctx, `SELECT md5(repeat('x', 1000000))`).Scan(&big); err != nil || big != "ec78dbd963d2fc01e51176ed4dec299e" {
		t.Fatalf("md5(repeat x 1e6) = %s, %v", big, err)
	}
}
