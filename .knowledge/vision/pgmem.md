---
id: vision:pgmem
type: vision
title: pgmem Vision
---
Real PostgreSQL 18 running fully in memory inside a Go test process, reachable through the normal wire protocol, with no Docker, binaries, or disk.

```yaml
summary:
  primary_use: automated tests
  scope_now: Go test API (requirement:test-fixture-fork, requirement:in-memory-only)
  scope_next: requirement:multi-language-wrapper (Python, Java via concept:server-process)
  constraints:
    - rule:single-session-per-backend
```
