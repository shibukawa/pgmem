---
id: requirement:in-memory-only
type: requirement
title: In-Memory Only Operation
---
Because the target is testing, every server, snapshot, and fork keeps all state in process memory; nothing is written to the host filesystem.

```yaml
summary:
  state_locations:
    - vfs tree (data directory, WAL, share files)
    - wasm linear memory (live backend session)
  persistence: optional explicit export only (concept:vfs-snapshot)
  implication: fork count is bounded by memory (policy:fork-pool-limit, metric:fork-cost)
```
