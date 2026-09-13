---
id: decision:node-test-integration
type: decision
title: Node.js Test Integration Shape
---
How Node.js suites get a database URL per test file and isolation per test from one pgmem process; agreed with the user 2026-09-13.

```yaml
decision:
  resolution: shared concept:server-process started from the user's own globalSetup + api:control-socket; URL env var per test file; api:reset per test; withFork for concurrent tests (api:node-wrapper)
  why_env_per_file: system:node-test-runners rebuild the process or process.env per file, and ORMs capture the URL when the client is built (system:node-orms)
  layers:
    run: PgmemServer.start({prepare}) runs migrations on the template, snapshots, exports PGMEM_CONTROL; env is the only channel that reaches workers in every runner
    file: setup entry (Vitest setupFiles, node:test --import, Bun --preload, Jest testEnvironment) forks via PGMEM_CONTROL and writes DATABASE_URL (names configurable) before test imports run
    test: fork.reset() in beforeEach keeps URL and pools; withFork(fn) hands a separate URL to code that accepts one
  composition: constructors plus close, never own the runner or its globalSetup, so pgmem sits next to osmem in one setup file; pgmem exec only as a fallback for runners without a global setup
  api_shape: mirrors @osmem/core (start, withClone, close) with pgmem's fork vocabulary
  rejected:
    env_per_test_case: rewriting process.env after import does not reach built clients (verified Node 26.8); vi.resetModules plus re-import is slow and breaks concurrent tests
    process_per_file_from_snapshot_file: osmem-like, but writes the prepared database to disk (requirement:in-memory-only) and costs ~166 MB per file process versus ~100 MB per fork (metric:fork-cost)
    runner_owning_wrapper: does not compose with sibling helpers
  constraints:
    - reset assumes tests in a file run sequentially (runner default); concurrent tests use withFork
    - data from beforeAll survives reset only if the file snapshots after beforeAll and resets to that snapshot
    - Vitest isolate false reuses workers and module caches, so the setup entry resets its fork instead of forking again
  poc: 2026-09-13 Vitest 5 with pre-forked servers per VITEST_POOL_ID; globalSetup env reached workers and app modules read the injected URL at import (4 files passed)
  verified: 2026-09-13 api:node-wrapper end to end; Vitest 5 run with Prisma, Drizzle and TypeORM schemas prepared in 1.9s, 4 files on separate forks, reset per test, 11 tests in 5.1s; Jest 30 environment 3 files passed
```
