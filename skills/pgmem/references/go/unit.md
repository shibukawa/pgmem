# Go unit tests

`github.com/shibukawa/pgmem/pgmemtest` provides `Fixture`. In `TestMain`, call `pgmemtest.Run(m, pgmemtest.Options{Options: pgmem.Options{Database: "app"}, Prepare: func(ctx context.Context, db *sql.DB, dsn string) error { ... }}, func(f *pgmemtest.Fixture) { fx = f })`. Run migrations and seed in `Prepare`; close transactions before it returns. The resulting snapshot feeds each fork.

For code that already reads `DATABASE_URL` or libpq `PG*` variables, call `fx.ShadowPG(t)` before constructing the repository or pool. It creates a fresh fork, sets those variables with `t.Setenv`, and restores them at cleanup. `ShadowOptions{ExtraEnv: []string{"APP_DATABASE_URL"}}` covers a custom DSN variable. `Shared: true` reuses a fork only for read-only tests. It cannot retarget a `*sql.DB` or pgx pool created before the call.

When the application accepts a database handle, use `fx.DB(t)` (`database/sql`) or `fx.PgxPool(t)` for in-process connections; `fx.PgxConn(t)` supplies one pgx connection, and `fx.DSN(t)` supplies a TCP URL. Each per-test helper owns cleanup.

`ShadowPG` tests cannot call `t.Parallel()` because environment changes are process-wide, even with `-parallel 1`. Parallelize packages with `go test -p 4 ./...`, or separate same-package test selections into processes. In-process parallel cases need separately injected handles and forks.

Full examples: `website/src/content/docs/guides/go/testing.mdx`.
