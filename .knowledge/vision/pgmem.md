---
id: vision:pgmem
type: vision
title: pgmem Vision
---
Real PostgreSQL 18 running fully in memory, reachable through the normal wire protocol, with no Docker, binaries to download, or disk: a test database that costs one dependency and forks in milliseconds.

```yaml
summary:
  primary_use: automated tests (requirement:test-fixture-fork, requirement:in-memory-only)
  scope_now:
    - Go library (api:go-server, api:clone) and pgmemtest helpers
    - Python, Java and Node.js packages bundling the binary (requirement:multi-language-wrapper, api:node-wrapper per decision:node-test-integration)
  positioning: concept:alternatives
  constraints:
    - rule:process-per-connection
    - concept:limits
```
