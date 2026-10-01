---
id: requirement:live-http-test-reset
type: requirement
title: Stable Endpoint for HTTP Tests
---
Long-lived API and browser tests keep one application and pgmem endpoint per pytest worker. Before each sequential test, restore a selected snapshot in place instead of switching the application's pool to a new fork. See api:python-wrapper and api:control-protocol.

```yaml
requirements:
  lifecycle: one session-scoped fork per worker; function-scoped reset before each HTTP test
  dataset: map names to snapshots; default is the prepared template snapshot; marker chooses another
  isolation: tests within one worker are sequential; xdist workers own independent pgmem processes and application servers
  connection: endpoint and pooled client connections survive restore; open transactions cause busy after timeout
  frameworks: API examples for FastAPI, Starlette, Flask, Django; Django uses ordinary PostgreSQL backend and plain pytest tests without Django TestCase or pytest-django DB lifecycle fixtures
  docs: English and Japanese Python guides split into unit, API, and Playwright E2E pages
  limits: persistent fork consumes one slot in the snapshot fork pool
```
