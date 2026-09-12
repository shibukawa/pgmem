---
id: decision:fork-release-detection
type: decision
title: Fork Release Detection
---
How concept:server-process learns that a fork, or the whole process, is no longer needed when the owner is a test in another process.

```yaml
decision:
  resolution: explicit close op as the primary path, stdin EOF as the process-level safety net (implemented 2026-09-12; both paths covered by cmd/pgmem/control_test.go TestShutdownAndStdinEOF)
  layers:
    explicit_close:
      who: fixture teardown (pytest yield fixture, JUnit AfterEach, try-with-resources)
      guarantee: same as decision:clone-release-style testing_tb; forgetting is impossible through the provided fixtures
    parent_death:
      signal: stdin EOF in concept:server-process; the kernel closes the pipe when the parent exits or crashes, on every OS
      effect: all forks, snapshots and the template close and the process exits
      note: independent of decision:data-plane-transport; the control pipe carries liveness, not the SQL socket
    child_death:
      signal: wrapper reader thread sees stdout EOF or process exit; pending waiters fail with server_exited, the next fork raises
  rejected:
    last_connection_closed:
      shape: close a fork when its client connection count drops to zero
      why: pools connect lazily and tests reconnect; zero connections is normal between statements, and a grace timer only makes it flaky
    lease_timeout:
      why: slow tests would lose their database; no good default exists
  uds_question: a uds listener has no peer so nothing is detectable there; per-connection close is visible on uds and TCP alike, but it is the wrong signal (last_connection_closed)
```
