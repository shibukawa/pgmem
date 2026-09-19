---
id: concept:vfs-snapshot
type: concept
title: VFS Snapshot Mechanism
---
A fork is a copy-on-write clone of the in-memory vfs data directory (the tree is copied, file bytes are shared until written) plus a fresh single-user backend instance started on that clone.

```yaml
summary:
  existing_code:
    - internal/vfs FS.Clone() copies the node tree and shares file bytes; Node.shared + Node.own() copy a file on its first in-place write (writeAt), O_TRUNC drops shared bytes, WriteFile/PutFile replace them (2026-09-19)
    - internal/engine Engine.Start(fs, opts) boots a backend on any vfs
    - internal/engine Tar/Untar for optional export of a snapshot
  snapshot_steps:
    - stop the template cluster cleanly (the shutdown checkpoint), clone, start it again; sessions re-attach (rule:process-per-connection)
    - FS.Clone() the vfs; keep the copy as immutable template image
  fork_steps:
    - FS.Clone() the template image
    - Engine.Start on the copy
    - net.Listen on a new loopback port
  startup_on_fork: short WAL recovery from the checkpoint record; near-empty
  cow_cost: a clone costs the node count (~2,100 nodes, well under 1 ms) instead of the 50 MB the data directory holds; the first write to the 16 MB WAL segment copies it once per fork
  cow_race_note: many Forks clone the same snapshot fs concurrently; cloneNode marks the source node shared only when it is not already, and every data node of a snapshot (itself a clone) already is, so the snapshot tree is only read
  anonymous_mmap: wasm/pgmem_shim.c overrides mmap for MAP_ANONYMOUS so PostgreSQL's ~37 MB shared-memory block is not memset when it comes from fresh sbrk memory (dlmalloc never trims, so bytes at or above the old break are zero); only the part below the old break is cleared (2026-09-19)
  gotcha: io_method must be sync; PGlite sets IsUnderPostmaster so worker AIO waits forever on batched read_stream reads (pgmem.go withDefaults)
  rejected: PostgreSQL CREATE DATABASE TEMPLATE per test (the command works in a live session since io_method=sync, but a fork is a whole cluster with its own buffer cache and processes, and forks are what the fixtures hand out)
```
