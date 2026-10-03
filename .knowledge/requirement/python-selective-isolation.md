---
id: requirement:python-selective-isolation
type: requirement
title: Python Selective Test Isolation
---
Permit a pytest decorator/marker argument to select or suppress per-test isolation without requiring each test to request a different DSN or manage connection lifetimes.

```yaml
requirements:
  selection: declarative per-test opt-in or suppression; default and marker precedence must be explicit
  setup: selection is resolved before driver fixtures or application clients connect
  reuse: read-only cases may share prepared state; writes must not leak into later cases
  compatibility: preserve explicit pgmem_dsn and pgmem_class_dsn fixtures (api:python-wrapper)
  driver_neutral: work with psycopg, asyncpg and SQLAlchemy without taking ownership of their connections
  cleanup: release any test fork even when setup, test or teardown fails
  concurrency: isolated writing cases remain independent under pytest-xdist
  cost: avoid a connection-routing layer merely to implement decorator selection
implementation:
  default: fork; configurable with pgmem_fixture_options; closest marker selects fork/shared/reset
  per_test_clients: application fixture depends on pgmem_test_dsn; marker is resolved before client construction
  persistent_clients: session fixture depends on pgmem_shared_dsn; isolation reset uses automatic teardown after marked serial tests
  shared_contract: shared template is read-only by convention; accidental writes can affect other shared tests; no SQL permission enforcement
  preparation: pgmem_prepare callback, automatic snapshot, configurable capacity and deadlines
```

Related: decision:fork-or-not, concept:python-guide.
