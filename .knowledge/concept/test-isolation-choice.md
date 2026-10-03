---
id: concept:test-isolation-choice
type: concept
title: Choose Test Isolation
---
Choose isolation from application client lifetime and whether test cases run concurrently.

```yaml
choices:
  recreate_clients: fork per test; bind application and assertions to one database object or selected DSN; writers may run concurrently
  persistent_clients: stable file/session endpoint and reset after each serial writer; no overlapping readers during reset
  shared_reads: shared prepared endpoint; optional connection read-only guard; shared resources occupy fork slots where a shared fork is used
  execution: Go embedded vs owned subprocess independent of inprocess/TCP/Unix SQL transport; keep external imports free of embedded engine
ownership:
  fixture: process, template, snapshot and forks
  application: clients, pools, ORM sessions, HTTP servers, caches
  ordering: stop serving requests; close application clients; release fork; close snapshot/process; attempt all releases even on error
reset_guarantees:
  restored: database rows, schema and sequences from the selected snapshot
  stable: endpoint and supported open client sockets (api:reset)
  external: application memory, ORM identity maps, Valkey keys and OpenSearch indexes require separate lifecycle
```
Related: requirement:test-api-ergonomics, flow:test-lifecycle, flow:wrapper-test-lifecycle.
