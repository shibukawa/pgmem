---
id: api:go-test-fixture
type: api
title: Go Test Database Helpers
---
One test database object supplies all driver handles for one fork. Implemented 2026-10-03 across pgmemtest, pgmemfixture and api:go-process.

```yaml
api:
  identity: fx.For(t) -> pgmemfixture.TestDB; DB(), PgxConn(), PgxPool(), DSN() refer to one fork; another For(t) explicitly acquires another fork
  cleanup: t.Cleanup closes driver handles before the fork; Reset(ctx) preserves the endpoint
  embedded: pgmemtest.Options.Transport selects inprocess default, tcp or unix; SnapshotTimeout and ForkTimeout default 30s for all fixture helpers; existing direct helpers still create one fork per call
  shared: SharedDB() and SharedPgxPool() use one read-only-by-convention fork, owned by Fixture.Close; acquire during setup where possible; shared fork consumes one slot
  dependency_boundary: pgmemfixture imports drivers but not api:go-server or its embedded engine; pgmemtest adapts embedded Server; pgmemprocess supplies network endpoints
  driver_config: variadic callbacks on TestDB.DB/PgxConn/PgxPool; preserve Host/Port/Database/User/Dial; customize RuntimeParams, pool limits and driver hooks
  guard: Options.SharedReadOnly guards SharedDB/SharedPgxPool and shared ShadowPG environment with default_transaction_read_only; session default rather than SQL privileges
  failures: cleanup attempts every release and joins errors; Run reports teardown failure as a nonzero test exit
  preparation: Prepare(ctx, db, dsn) initializes one baseline; automatic snapshot follows; failure cleans up process/server and driver handle
```
