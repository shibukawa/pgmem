---
id: api:reset
type: api
title: Reset Server In Place
---
Op that returns a running server to a snapshot while keeping its port and client connections, so a URL captured at import time stays valid across tests (decision:node-test-integration).

```yaml
api:
  status: implemented 2026-09-13 (reset.go; reset_test.go; op reset tested by cmd/pgmem TestResetOp); on the cluster model since 2026-09-19
  op: 'reset in {"server":"f1","snapshot":"s1"?,"timeout_ms":N?} out {}; snapshot defaults to the one the fork came from'
  steps:
    - wait until no session is inside a transaction or a batch; busy after timeout
    - mute the backends, kill the cluster (the old directory is discarded), start one on a clone of the snapshot
    - every session re-attaches at once: startup packet replayed on a new backend, LISTEN re-issued, recorded Parse messages of named prepared statements replayed (node-postgres never re-sends Parse for a name it parsed on that connection) (rule:process-per-connection)
  lost: session SET values, temp tables, portals; same as a fresh session
  cost: ~15 ms (metric:fork-cost)
  go: Server.Reset(ctx) to the origin snapshot, Server.Restore(ctx, snap) to any snapshot
  rejected: dropping every client connection on reset, because pools without an 'error' listener crash the test process (system:node-orms)
```
