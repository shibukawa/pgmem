---
id: requirement:test-fixture-fork
type: requirement
title: Per-Test Forked Fixture
---
A test suite prepares schema and seed data once, then every test case receives an isolated copy of that prepared state that is discarded when the test ends.

```yaml
summary:
  prepare_once: TestMain runs migrations and seed loading against a template server (flow:test-lifecycle)
  per_test: api:clone gives an isolated database handle
  isolation: writes in one test never reach the template or other tests
  cleanup: automatic; a test must not be able to leak a fork (decision:clone-release-style)
  parallel: forks are independent clusters so t.Parallel() is safe (rule:process-per-connection)
  budget: api:clone blocks instead of over-allocating (policy:fork-pool-limit)
```
