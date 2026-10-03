---
id: api:go-process
type: api
title: Go Subprocess Client
---
Lightweight Go client for requirement:go-external-test-backend, implemented 2026-10-03 in pgmemprocess. Does not import the embedded engine.

```yaml
api:
  start: Start(ctx, Options) -> Process; startup context does not own the returned process lifetime
  options: [Binary, Database, User, Params, Transport, SocketDir, Log, StartupTimeout, ShutdownTimeout]
  binary: explicit Binary, PGMEM_BINARY, then pgmem on PATH; no runtime download
  transports: tcp default; unix explicitly selected on AF_UNIX stream hosts including modern Windows; private socket directories under SocketDir, default /tmp on Unix or os.TempDir() on Windows
  process: Template, Version, PID(), StartServer(ctx, database, user), Close()
  server: DSN(), Host(), Port(), Dial(ctx, network, addr), Snapshot(ctx, maxForks), Reset(ctx), Restore(ctx, snapshot), Close()
  snapshot: Fork(ctx), Close()
  channel: api:control-protocol over stdio; concurrent responses dispatched by id; cancelled allocations close on late completion
  defaults: startup 30s, shutdown 10s; low-level operations use caller context
  fixture: NewFixture(ctx, FixtureOptions), Run(m, opts, setup), For(t), SharedDB(), SharedPgxPool(), Template(), Snapshot(), Close() (api:go-test-fixture)
  fixture_options: Options plus Prepare, MaxForks, ForkTimeout and SnapshotTimeout; fork and snapshot defaults 30s
  guard: FixtureOptions.SharedReadOnly guards shared driver connections; writable test forks unchanged
  cleanup: EOF shuts down the owned child; bounded shutdown kills a stuck child; graceful Unix listener close removes owned paths
  tests: TCP and Unix SQL, same-fork handles, isolation, reset under open connections, capacity, startup deadline, child death, prepare failure and socket cleanup; client race test
```
