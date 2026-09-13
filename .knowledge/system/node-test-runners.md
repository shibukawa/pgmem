---
id: system:node-test-runners
type: system
title: Node.js Test Runners
---
Isolation and environment behavior of Node.js test runners that bounds how a database URL reaches application code (checked 2026-09-13).

```yaml
system:
  vitest_5:
    default: pool forks with isolate true, so a child process per test file; maxWorkers = available parallelism (half in watch)
    global_setup: runs before workers; process.env assignments reach workers; config test.env wins on conflicts; project.provide and inject pass values
    setup_files: awaited before the test file imports run; top-level await works; hooks registered there apply to every test
    fixtures: test.extend with scope file or worker since 3.2
    ids: VITEST_POOL_ID in 1..maxWorkers, reused
  jest_30:
    workers: long-lived; environment and process.env copied per test file; JEST_WORKER_ID from 1
    global_setup: process.env assignments reach workers
    per_file_hook: async testEnvironment setup() runs before test file imports; setupFiles cannot await
    esm: still experimental, so a library ships CJS
  node_test:
    isolation: a child process per test file by default
    global_setup: --test-global-setup since Node 24 (early development); env reaches children (verified Node 26.8)
    per_file_hook: --import preload with top-level await (verified Node 26.8)
  bun_1_4: one process and module registry by default; --isolate (1.3.13+) reruns preloads per file; no global setup
  playwright: webServer starts before globalSetup, so a URL computed there cannot reach webServer.env
  node_versions: 24 active LTS, 22 maintenance, 26 current; await using is native from 24
  package_managers: pnpm 10+, Yarn 4.14, Bun and npm 11.19 hold dependency install scripts by default (npm observed 2026-09-13)
```
