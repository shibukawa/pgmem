---
id: requirement:test-suite-preparation
type: requirement
title: Test Suite Preparation
---
Python pytest integration prepares a reusable database state before test cases run.

```yaml
requirements:
  lifecycle: one setup per pytest session
  order: start concept:server-process; run migrations/schema and seed data with the normal driver; close or commit preparation connections; create api:snapshot
  extensibility: user supplies preparation callback or setup fixture; preparation may call existing migration tools
  failure: preparation or snapshot failure stops dependent tests and releases owned resources
  reuse: test cases select the prepared state through requirement:declarative-test-target
  existing: pgmem_prepare fixture with @shadow_pg(prepare=True), followed by pgmem_snapshot
```
