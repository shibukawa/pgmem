---
id: requirement:declarative-test-target
type: requirement
title: Declarative Test Database Target
---
The Python pytest adapter lets a case declare pgmem routing and isolation without changing SQL or the PostgreSQL driver.

```yaml
requirements:
  default: give each test isolated prepared state; use a fresh fork and close it after the test (requirement:test-fixture-fork)
  opt_out: 'fork=false selects one prepared shared fork for tests that only read; writes are not database-blocked and affect later shared tests; never expose the mutable preparation template'
  routing: pytest marker intercepts new psycopg/psycopg2/asyncpg connections and their SQLAlchemy dialects; it also sets PostgreSQL URL and libpq environment variables for the marked test
  scope: pytest decorator/marker
  cleanup: restore any changed test-process configuration and close owned fork even when the test fails
  parallel: isolated forks may run concurrently; shared-target tests must not mutate shared state
  constraint: import-time clients and already checked-out connections cannot be rerouted; a long-lived HTTP app uses a stable fork and api:reset (requirement:live-http-test-reset)
  existing: api:python-wrapper
```

Implemented as the Python `pgmem.shadow_pg` pytest marker. Read-only enforcement is not implemented.
