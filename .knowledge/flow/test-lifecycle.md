---
id: flow:test-lifecycle
type: flow
title: Test Suite Lifecycle
---
How a Go test package uses pgmem from TestMain through each test case.

```yaml
flow:
  - step: start template
    actor: TestMain
    action: pgmem.Start(ctx, Options) -> template server (concept:template-server)
  - step: prepare
    actor: TestMain
    action: run migrations and seed data through template DSN with an ordinary driver
  - step: snapshot
    actor: TestMain
    action: api:snapshot freezes template state
  - step: run tests
    actor: TestMain
    action: m.Run()
  - step: fork
    actor: test case
    action: api:clone acquires a slot (policy:fork-pool-limit), clones vfs, starts backend, returns handle
  - step: exercise
    actor: test case
    action: use handle as *sql.DB or pgx.Conn (decision:clone-return-type)
  - step: release
    actor: test case
    action: fork closed automatically (decision:clone-release-style); slot returned to pool
  - step: teardown
    actor: TestMain
    action: close snapshot and template server; os.Exit(code)
```
