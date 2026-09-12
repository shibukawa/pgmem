---
id: concept:go-guide
type: concept
title: Go Guide
---
Using pgmem from Go tests: install, a first test, initialization through pgmemtest, fork styles and parallelism.

```yaml
summary:
  install: go get github.com/shibukawa/pgmem; first go build compiles 102 MB of generated Go once (~35 s), cached after that
  basic_test: |
    s, err := pgmem.Start(ctx, pgmem.Options{Database: "app"})  // api:go-server
    defer s.Close()
    db, _ := sql.Open("pgx", s.DSN())
    db.Exec("create table t(id int)")
  initialization:
    helper: pgmemtest.Run in TestMain with Options.Prepare(ctx, db, dsn) running migrations and seed once (flow:test-lifecycle)
    raw_sql: db.ExecContext(ctx, schemaSQL) inside Prepare
    tools: golang-migrate, goose, atlas, tern (concept:migration-tools); seed with go-testfixtures (concept:seeding)
  test_case_styles:
    fork_per_test: fx.DB(t), fx.PgxConn(t), fx.PgxPool(t), fx.DSN(t), fx.Fork(t); closed by t.Cleanup (api:clone, decision:clone-release-style)
    no_fork: use s.DSN() from a package-level server with BEGIN and ROLLBACK per test; sequential only
    server_per_test: pgmem.Start inside the test for schema-per-test cases
    guidance: decision:fork-or-not
  parallel: t.Parallel() is safe with forks; MaxForks defaults to GOMAXPROCS and Fork blocks when full (policy:fork-pool-limit)
  pools: pgxpool on a fork works; connections serialize at transaction boundaries (rule:single-session-per-backend)
  speed: Fixture.DB uses api:in-process-dialer; ~3x faster small queries than TCP
  logging: Options.Log or pgmemtest stderr forwarding; PGMEM_TRACE=1 for host calls
  limits: concept:limits
```
