---
id: concept:architecture
type: concept
title: Architecture
---
PostgreSQL 18.3 (PGlite fork) compiled to wasm once, translated to Go ahead of time, and run in single-user mode against an in-memory filesystem; a Go bridge exposes it on loopback TCP.

```yaml
summary:
  layers:
    - name: postgres source
      what: electric-sql/postgres-pglite pinned by commit and sha256 in wasm/postgres-pglite.lock; patched at build time by wasm/patches.py (CRC-32C, cryptohash, UUID entropy, LISTEN commit hook routed to the host)
    - name: wasm module
      what: wasm/build.sh links one static module with Emscripten 3.1.74; no dynamic linking, loadable modules statically linked (concept:static-modules); no OpenSSL, ICU, zlib
    - name: generated Go
      what: wasm2go pure-Go backend turns postgres.wasm into internal/aot/pgaot (102 MB source, ~35 s compile once); the wasm file is a build intermediate and does not ship
    - name: host table
      what: internal/host implements the Emscripten env and WASI imports (syscalls, clock, memory growth, exit unwinding); entropy, CRC-32C and hashes run on Go crypto and hash packages
    - name: vfs
      what: internal/vfs in-memory POSIX-like filesystem holding the data directory, WAL and share files; FS.Clone gives concept:vfs-snapshot
    - name: engine
      what: internal/engine boots initdb output and runs postgres --single the way PGlite does; frontend and backend protocol flows through in-memory buffers; ereport(ERROR) unwinds use the PGlite exit trick
    - name: bridge
      what: pgmem.go accepts TCP on 127.0.0.1 (or api:in-process-dialer), multiplexes connections onto the one session (rule:single-session-per-backend), routes NOTIFY per connection
    - name: embedded data
      what: internal/pgdata/pgdata.tar.zst is initdb output produced at build time by cmd/pgmem-mkdata (initdb wasm under wazero); unpacking it is why Start takes ~0.1 s (metric:server-footprint)
    - name: process and wrappers
      what: cmd/pgmem wraps the library as concept:server-process for api:python-wrapper and api:java-wrapper over api:control-protocol
  query_path:
    - client driver connects to the DSN over loopback TCP
    - pgmem.go serve parses the startup packet; only Options.Database is served
    - connection acquires the backend semaphore for one message batch, or for the whole transaction after BEGIN
    - engine.Backend feeds the bytes to PostgresMainLoopOnce in generated Go
    - page reads and WAL writes hit internal/vfs; no host file is touched (requirement:in-memory-only)
    - response bytes are copied back to the socket; NotifyResponse is routed to listening connections
  repo_map:
    pgmem.go: Start, Options, Server, connection multiplexing
    snapshot.go: api:snapshot and api:clone
    pgmemtest/: Go test helpers (decision:clone-release-style)
    cmd/pgmem: standalone binary and control protocol
    cmd/pgmem-mkdata: regenerates the embedded data directory
    wasm/: build scripts, patches, static module table
    packages/python, packages/java: wrappers (requirement:multi-language-wrapper)
  rebuild: emsdk 3.1.74 at toolchain/emsdk; ./wasm/build.sh, ./wasm/gen-aot.sh, go run ./cmd/pgmem-mkdata, go test ./...
  aot_backend_choice: pure Go beats the wasm2go asm backend on size (102 vs 281 MB) and speed (95 vs 122 ms sort of 200k rows); asm stays available with ASM=1
```
