# Go HTTP API tests

Use the prepared `pgmemtest.Fixture` from [unit tests](unit.md). In each sequential test, call `fx.ShadowPG(t)` **before** `NewApp()` creates its pool, then start an `httptest.Server` for that app. Exercise routes with `server.Client()` and close both app and server with `t.Cleanup`. Add `ShadowOptions{ExtraEnv: ...}` if the application reads a custom DSN variable.

Each test receives a fork of the prepared schema and seed. A pre-existing pool retains its old destination. Finish handlers, background jobs, and transactions before cleanup. Do not use `t.Parallel()` with `ShadowPG`; use separate package processes for worker-level parallelism. If methods must overlap, inject a separate `fx.DB(t)` or `fx.PgxPool(t)` into each app instead.

Full example: `website/src/content/docs/guides/go/api-testing.mdx`.
