---
id: flow:wrapper-test-lifecycle
type: flow
title: Wrapper Test Lifecycle
---
How a Python or Java test suite uses pgmem through concept:server-process; mirrors flow:test-lifecycle across the process boundary.

```yaml
flow:
  - step: spawn
    actor: wrapper session fixture or JUnit BeforeAll
    action: locate the binary (policy:binary-distribution), spawn with stdin and stdout pipes, stderr inherited
  - step: ready
    actor: wrapper
    action: read the first line, check protocol (api:control-protocol), expose the template DSN; fail after 30s
  - step: prepare
    actor: user code
    action: migrations and seed through the template DSN with the normal driver (system:postgres-drivers); op start first when a suite needs more than one template
  - step: snapshot
    actor: wrapper
    action: op snapshot per template returns a snapshot id (api:snapshot); busy error if a connection is still in a transaction
  - step: fork
    actor: per-test fixture
    action: op fork returns an endpoint; blocks under policy:fork-pool-limit (rule:non-blocking-control-channel)
  - step: exercise
    actor: test
    action: connect with the DSN or JDBC URL; pools allowed (rule:single-session-per-backend)
  - step: release
    actor: per-test fixture teardown
    action: op close on the fork (decision:fork-release-detection)
  - step: teardown
    actor: wrapper
    action: op shutdown, close stdin, wait 10s for exit, then kill
```
