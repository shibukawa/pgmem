---
id: rule:process-per-connection
type: rule
title: Process per Connection
---
PostgreSQL runs its real multi-process model: a postmaster starts a backend process per connection plus the auxiliary processes, all sharing memory, so sessions, locks and their interactions are PostgreSQL's own. Replaced rule:single-session-per-backend on 2026-09-19.

```yaml
rule:
  process: one module instance (goroutine) per PostgreSQL process; EXEC_BACKEND build, the postmaster starts children through the pgmem_spawn host import (internal/host/cluster.go)
  shared_memory: System V segments are files MAP_FIXED-aliased at the same address (0x6000_0000..) into every instance's reserved linear memory; the heap is capped below; atomics are real (-matomics, wasm2go forceContendedAtomics)
  signals: kill() queues on the target Host and is delivered at its next host import (or wakes it out of poll/semaphore/sleep) through the pgmem_raise export; a killed process exits at its next host call
  sessions: per connection, as on a server; locks between connections wait; deadlocks end with 40P01 after deadlock_timeout; pg_terminate_backend, pg_stat_activity and pg_locks see the other sessions
  restart_survival: a Restore/Snapshot restarts the cluster; pgmem's per-connection proxy (cluster.go) keeps the client socket, re-attaches to a new backend on the next message, replays the startup packet, named Parse messages and LISTEN (pgmem_listen hook, ignored while the backend is muted)
  postgres_ended_backends: a FATAL from PostgreSQL itself (pg_terminate_backend, DROP DATABASE WITH FORCE, crash) closes the client connection like a server does
  standalone: initdb's children and the setup child still run postgres --single with local shm/semaphores (internal/host/single.go)
  cancel: a statement that makes no system call is cancelled at its next one
  cost: metric:fork-cost (fork 3.8 ms, first connection 2.4 ms, close 0.1 ms; one test 7.7 ms against 6.2 ms in the removed single-user model)
  windows: the reservation is a placeholder (VirtualAlloc2 MEM_RESERVE_PLACEHOLDER); heap pieces are committed and segment views mapped with MEM_REPLACE_PLACEHOLDER (MapViewOfFile3), pieces tracked in internal/aot/placeholder.go (tested with a fake); segment files are delete-on-close; Windows 10 1803 / Server 2019 or later
```
