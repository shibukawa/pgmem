---
id: requirement:test-api-ergonomics
type: requirement
title: Test API Ergonomics
---
User authorized all proposed API and documentation improvements on 2026-10-03.

```yaml
requirements:
  automatic_deadlines: all Go fixture acquisitions use ForkTimeout; Node register/useFork/Jest default fork 30s and reset 5s with independent configuration; manual snapshot fork APIs retain explicit unbounded waits
  cleanup: attempt every owned release after a failure; preserve primary and cleanup failures via Go errors.Join, JS AggregateError, Java suppressed exceptions, Python CleanupError for Python 3.9 compatibility
  driver_config: Go TestDB DB/PgxConn/PgxPool accept callbacks; endpoint identity and fixture Dial remain fixed; driver runtime params, pool sizes and hooks are configurable
  service_env: lightweight Node environment subpath registers named claims across services; preflight conflicts, rollback partial assignment, idempotent release restoring prior unclaimed values; unregistered services cannot be attributed
  shared_guard: opt-in default_transaction_read_only on convenience connection URLs; protects ordinary persistent writes; not a privilege boundary; isolated writing and reset clients remain writable
  choice_guide: complete English/Japanese lifecycle decision page with scopes, parallel constraints, deadlines, and cleanup owners (concept:test-isolation-choice)
  app_examples: runnable Go and Node HTTP, Python SQLAlchemy, Java DataSource applications; test observes application writes through the same fork; examples executed by package CI
  reset_contract: guarantee matrix explicitly separates database state, client caches, application memory, Valkey keys and OpenSearch indexes
implementation:
  Go: api:go-test-fixture; api:go-process
  Python: api:python-wrapper
  Node: api:node-wrapper; requirement:node-test-composition
  Java: api:java-wrapper
```
