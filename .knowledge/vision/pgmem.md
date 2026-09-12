---
id: vision:pgmem
type: vision
title: pgmem Vision
---
Real PostgreSQL 18 running fully in memory, reachable through the normal wire protocol, with no Docker, binaries to download, or disk: a test database that costs one dependency and forks in 20 ms.

```yaml
summary:
  primary_use: automated tests (requirement:test-fixture-fork, requirement:in-memory-only)
  scope_now:
    - Go library (api:go-server, api:clone) and pgmemtest helpers
    - Python and Java packages bundling the binary (requirement:multi-language-wrapper)
  scope_next:
    - requirement:nodejs-wrapper
  positioning: concept:alternatives
  constraints:
    - rule:single-session-per-backend
    - concept:limits
```
