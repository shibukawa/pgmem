---
id: requirement:in-memory-only
type: requirement
title: In-Memory Only Operation
---
Because the target is testing, every server, snapshot, and fork keeps database state in process memory; database files are not written to the host filesystem.

```yaml
summary:
  state_locations:
    - vfs tree (data directory, WAL, share files)
    - wasm linear memory (live backend session)
  persistence: optional explicit export only (concept:vfs-snapshot)
  transport_paths: explicit Unix socket mode creates owned socket entries and runtime directories, but no database data or snapshot files (requirement:go-external-test-backend)
  implication: fork count is bounded by memory (policy:fork-pool-limit, metric:fork-cost)
```
