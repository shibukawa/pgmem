---
id: api:reset
type: api
title: Reset Server In Place
---
Op that returns a running server to a snapshot while keeping its port and client connections, so a URL captured at import time stays valid across tests (decision:node-test-integration).

```yaml
api:
  status: implemented 2026-09-13 (reset.go; reset_test.go, reset_internal_test.go; op reset tested by cmd/pgmem TestResetOp)
  op: 'reset in {"server":"f1","snapshot":"s1"?,"timeout_ms":N?} out {}; snapshot defaults to the one the fork came from'
  steps:
    - acquire the backend; wait for open transactions; busy after timeout (rule:single-session-per-backend)
    - clone the snapshot vfs, start a new backend, run startup for the configured user and database
    - keep client TCP connections and their statement-name prefixes
    - re-issue LISTEN for registered channels
    - replay recorded Parse messages of each connection's named prepared statements (node-postgres never re-sends Parse for a name it parsed on that connection)
  lost: session SET values, temp tables, portals; same as a fresh session
  cost: close to api:clone (metric:fork-cost)
  skip: no restart when nothing ran since the last reset
  go: Server.Reset(ctx) to the origin snapshot, Server.Restore(ctx, snap) to any snapshot
  rejected: dropping every client connection on reset, because pools without an 'error' listener crash the test process (system:node-orms)
```
