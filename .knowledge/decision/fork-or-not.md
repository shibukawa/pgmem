---
id: decision:fork-or-not
type: decision
title: Fork per Test or Not
---
Which isolation style a test suite should use with pgmem: a fork per test, a fork per class, one shared server with rollback, or a server per test.

```yaml
decision:
  resolution: default to one fork per test case; share a fork per class only for read-only tests; drop to a single server with transaction rollback only when memory is tight and tests are sequential (2026-09-12)
  options:
    fork_per_test:
      cost: ~20 ms per test plus ~100 MB per live fork, bounded by policy:fork-pool-limit (default CPU count)
      pros: [full isolation including commits, DDL and sequences, parallel tests are free, teardown is a close]
      cons: [memory per live fork, a test holding two forks under a full pool deadlocks]
      how: Go Fixture.DB(t) (api:clone); Python pgmem_dsn (api:python-wrapper); Java Fork parameter (api:java-wrapper)
    fork_per_class:
      cost: one fork per test class
      pros: [cheaper for many small read-only tests]
      cons: [a writing test leaks state into its siblings; order-dependent failures]
      how: Python pgmem_class_dsn; Java @PgmemFork(scope = CLASS); Go a package-level fork created in TestMain after the snapshot
    shared_server_rollback:
      cost: no fork; one BEGIN and ROLLBACK per test
      pros: [lowest memory, no snapshot needed]
      cons: [code under test must not COMMIT or open a second connection; sequential only because the session is one (rule:single-session-per-backend); sequences and DDL side effects survive rollback]
      how: open a transaction on the template DSN in the test setup and roll back in teardown
    shared_server_truncate:
      cost: TRUNCATE ... RESTART IDENTITY CASCADE between tests
      pros: [works with committing code]
      cons: [sequential, seed data must be reloaded, slower than a fork on wide schemas]
    server_per_test:
      cost: ~0.1 s and ~150 MB per test (metric:server-footprint)
      pros: [no shared state at all, schema can differ per test]
      cons: [no prepared data unless each test migrates, memory if parallel]
      how: pgmem.Start in the test (api:go-server) or pgmem.start() / Pgmem.builder().start() per test
    several_templates:
      when: suites with more than one seed set or schema
      how: Go one pgmemtest.Fixture per set; Python start_server plus a second session snapshot fixture; Java builder().template(name, prepare) and @PgmemFork("name") (api:control-protocol op start)
  rule_of_thumb: if the test commits or the suite runs in parallel, fork; if it only reads, share a class fork; snapshot once after migrations and seeding (concept:migration-tools, concept:seeding)
```
