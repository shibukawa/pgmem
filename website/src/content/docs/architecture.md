---
title: Architecture
description: How PostgreSQL 18 ends up running in memory inside a Go process, and how forks and the wrappers work.
---

pgmem is PostgreSQL 18.3 compiled to WebAssembly once, translated to Go ahead of time, and run in single-user mode against an in-memory file system. A small bridge exposes it on a loopback TCP port.

## Layers

| Layer | What it is |
|---|---|
| PostgreSQL source | The PGlite fork (`electric-sql/postgres-pglite`), pinned by commit and checksum in `wasm/postgres-pglite.lock`, patched at build time by `wasm/patches.py`. |
| WebAssembly module | `wasm/build.sh` links one static module with Emscripten. Loadable modules and extensions are linked in, not loaded dynamically. There is no OpenSSL, ICU or zlib. |
| Generated Go | wasm2go translates the module into the pure-Go package `internal/aot/pgaot`. The `.wasm` file is a build intermediate; only Go ships, so there is no wasm runtime and no cgo. |
| Host | `internal/host` implements the system calls, clock, memory growth and exit unwinding the module imports. Entropy, CRC-32C, hashes and pgcrypto's ciphers run on Go's `crypto` and `hash` packages. |
| File system | `internal/vfs` is an in-memory POSIX-like file system holding the data directory, WAL and share files. Nothing is written to the host disk. |
| Engine | `internal/engine` runs `postgres --single` the way PGlite does: protocol bytes flow through in-memory buffers, and `ereport(ERROR)` unwinds are handled without a postmaster. |
| Bridge | `pgmem.go` accepts TCP connections on `127.0.0.1` (or an in-process `net.Pipe`), multiplexes them onto the single session at transaction boundaries, and routes `NOTIFY` to the connections that listen. |
| Embedded data | `internal/pgdata/pgdata.tar.zst` is the output of `initdb`, produced at build time. Unpacking it into the file system is why a server starts in about a tenth of a second. |

## A query's path

1. The driver connects to the DSN over loopback TCP.
2. The bridge reads the startup packet. When it names another database of the server, the backend is restarted on that database first; a database that does not exist is refused with SQLSTATE 3D000.
3. The connection takes the backend for one message batch, or for the whole transaction after `BEGIN`.
4. The bytes are fed to PostgreSQL's main loop in generated Go.
5. Page reads and WAL writes hit the in-memory file system.
6. Responses are copied back to the socket.

## Snapshots and forks

A **snapshot** runs `CHECKPOINT` on a prepared server and deep-copies its in-memory data directory. A **fork** copies that image again, starts a new backend on the copy and listens on a new port. A **reset** does the same for a running server in place: the port and client connections stay, so a URL an application captured at import time remains valid. Startup replays almost no WAL, so a fork takes milliseconds, and the loaded engine code is shared by every server in the process.

Forks are full backends. They are independent of each other, so tests that write can run in parallel, and each one holds its own buffer cache and data directory copy. `MaxForks` bounds how many exist at once; asking for another blocks until one is closed.

## Other languages

The Python, Java and Node.js packages bundle `pgmem`, a standalone binary built from `cmd/pgmem`. A wrapper spawns it, reads one JSON line with the template's DSN, and sends `snapshot`, `fork` and `close` requests as JSON lines on stdin. When the parent exits, stdin closes and the binary shuts everything down, so a crashed test run cannot leave servers behind. Node.js test workers, which are separate processes, fork and reset through a loopback control socket instead; forks made on a socket connection close when that connection ends.

Every test process (a pytest-xdist worker, a forked test JVM) starts its own binary. Within it, forks behave exactly as in Go.
