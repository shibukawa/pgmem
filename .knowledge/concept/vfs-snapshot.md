---
id: concept:vfs-snapshot
type: concept
title: VFS Snapshot Mechanism
---
A fork is a deep copy of the in-memory vfs data directory plus a fresh single-user backend instance started on that copy.

```yaml
summary:
  existing_code:
    - internal/vfs FS.Clone() deep copy (used by snapshot.go)
    - internal/engine Engine.Start(fs, opts) boots a backend on any vfs
    - internal/engine Tar/Untar for optional export of a snapshot
  snapshot_steps:
    - issue CHECKPOINT on template backend (single-user mode runs it synchronously)
    - FS.Clone() the vfs; keep the copy as immutable template image
  fork_steps:
    - FS.Clone() the template image
    - Engine.Start on the copy
    - net.Listen on a new loopback port
  startup_on_fork: short WAL recovery from the checkpoint record; near-empty
  optimization_later: copy-on-write data slices in cloneNode; write paths are writeAt and truncateNode only
  gotcha: io_method must be sync; PGlite sets IsUnderPostmaster so worker AIO waits forever on batched read_stream reads (pgmem.go withDefaults)
  rejected: PostgreSQL CREATE DATABASE TEMPLATE per test (the command works in a live session since io_method=sync, and concept:database-switching serves the copy, but every alternation between databases restarts the backend and forks run in parallel while databases do not)
```
