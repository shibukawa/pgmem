---
id: concept:architecture
type: concept
title: Architecture
---
PostgreSQL 18.3 (PGlite fork) compiled to wasm once, translated to Go ahead of time, and run as a real cluster (postmaster, auxiliary processes, a backend process per connection, each a module instance) against an in-memory filesystem; a Go bridge exposes it on loopback TCP.

```yaml
summary:
  layers:
    - name: postgres source
      what: electric-sql/postgres-pglite pinned by commit and sha256 in wasm/postgres-pglite.lock; patched at build time by wasm/patches.py (CRC-32C, cryptohash, UUID entropy, LISTEN commit hook routed to the host, pgcrypto OpenSSL and zlib layers per decision:pgcrypto-on-host, PGlite ReadyForQuery fixes, EXEC_BACKEND spawn through pgmem_spawn, pg_usleep on the host, 32-bit BackendParameters size, checkpoint requests to the checkpointer)
    - name: wasm module
      what: wasm/build.sh links one static module with Emscripten (devbox 6.0.8) with -matomics and -DEXEC_BACKEND; libc calls that touch other processes, shared memory, semaphores and sockets are renamed to pgmem_* wrappers in wasm/pgmem_shim.c that call host imports; no dynamic linking; plpgsql, contrib extensions and pgvector statically linked (concept:static-modules, policy:bundled-extensions); no OpenSSL, ICU, zlib
    - name: generated Go
      what: wasm2go pure-Go backend (shibukawa/wasm2go-fork branch pgmem, pinned by wasm/wasm2go.lock) turns postgres.wasm into internal/aot/pgaot (104 MB source, ~35 s compile once); functions are named after PostgreSQL symbols and grouped into files by subject so a rebuild changes only the functions that changed; the wasm file is a build intermediate and does not ship
    - name: host table
      what: internal/host implements the Emscripten env and WASI imports (syscalls, clock, memory growth, exit unwinding) and, in cluster.go, the kernel of the process model (rule:process-per-connection); entropy, CRC-32C and hashes run on Go crypto and hash packages
    - name: vfs
      what: internal/vfs in-memory POSIX-like filesystem holding the data directory, WAL and share files, shared by the processes of a cluster under one mutex (per-process fd tables, ForkFDs inheritance, sockets, pipe EOF); FS.Clone gives concept:vfs-snapshot
    - name: engine
      what: internal/engine boots initdb output and runs the postmaster (engine/cluster.go: process goroutines, readiness from postmaster.pid, Kill for Close and Restore, clean Shutdown for Snapshot); the setup child still runs postgres --single
    - name: bridge
      what: pgmem.go accepts TCP on 127.0.0.1 (or api:in-process-dialer) and hands each connection to the postmaster; cluster.go keeps a per-connection proxy so sessions survive the cluster restart of api:reset (rule:process-per-connection)
    - name: embedded data
      what: internal/pgdata/pgdata.tar.zst is initdb output produced at build time by cmd/pgmem-mkdata (initdb wasm under wazero); unpacking it is why Start takes ~0.1 s (metric:server-footprint)
    - name: process and wrappers
      what: cmd/pgmem wraps the library as concept:server-process for api:python-wrapper, api:java-wrapper and api:node-wrapper over api:control-protocol, plus api:control-socket for test workers
  query_path:
    - client driver connects to the DSN over loopback TCP
    - pgmem.go hands the connection to the postmaster's accept queue; the postmaster spawns a backend process (a new module instance) that reads the startup packet itself
    - the bytes flow to that backend's socket (host.ConnSock) and PostgresMain runs in generated Go on the process's goroutine
    - locks, NOTIFY and the other sessions go through the shared memory segments
    - page reads and WAL writes hit internal/vfs; no host file is touched (requirement:in-memory-only)
    - response bytes are copied back to the socket; NotifyResponse is routed to listening connections
  repo_map:
    pgmem.go: Start, Options, Server; cluster.go: per-connection proxy that survives restarts
    internal/host/cluster.go: processes, signals, shared memory, semaphores, sockets, poll
    snapshot.go: api:snapshot and api:clone
    pgmemtest/: Go test helpers (decision:clone-release-style)
    cmd/pgmem: standalone binary and control protocol
    cmd/pgmem-mkdata: regenerates the embedded data directory
    wasm/: build scripts, patches, static module table, lock files for postgres-pglite, pgvector and wasm2go
    testdata/regress: upstream contrib regression files replayed by regress_test.go
    packages/python, packages/java, packages/node: wrappers (requirement:multi-language-wrapper)
  rebuild: devbox run -- ./wasm/build.sh (~5 min), ./wasm/gen-aot.sh (~1.5 min, pinned wasm2go fork), go run ./cmd/pgmem-mkdata, go test ./...
  aot_backend_choice: pure Go beats the wasm2go asm backend on size (104 vs 281 MB) and speed (95 vs 122 ms sort of 200k rows); asm stays available with ASM=1
```
