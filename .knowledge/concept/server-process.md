---
id: concept:server-process
type: concept
title: Server Process
---
One pgmem binary instance spawned per test process; it hosts the template server, snapshots and forks, and its lifetime is bound to the stdin pipe from the parent.

```yaml
summary:
  binary: cmd/pgmem; main.go starts the template and hands stdin/stdout to the controller in control.go
  hosts: the default concept:template-server plus any started by op start, their api:snapshot objects and every api:clone fork, all in one OS process
  why_one_process: forks share the loaded engine (pgmem.go sharedEngine), so a fork costs ~20ms (metric:fork-cost) instead of ~0.1s and ~150MB per extra process
  stdin: control requests (api:control-protocol); EOF means the parent exited or closed us, so close everything and exit
  stdout: ready event, responses, events; nothing else may ever be written there
  stderr: server log with -log; the wrapper inherits it
  exit_codes: 0 normal; non-zero on startup failure with the message on stderr
  ownership: exactly one parent; multi-process test runners (pytest-xdist, Gradle forked JVMs) spawn one process each and prepare each one
  memory: template plus up to MaxForks backends, same bound as Go (policy:fork-pool-limit)
  compat: ready-line-only wrappers from docs/subprocess.md keep working because they never write to stdin
```
