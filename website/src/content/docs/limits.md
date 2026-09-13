---
title: Limits
description: What pgmem does differently from a PostgreSQL server started by a postmaster.
---

## One session per server

PostgreSQL runs in single-user mode, so every connection to a server or fork shares one backend session. pgmem multiplexes connections the way a transaction-mode pooler does: a connection owns the backend from `BEGIN` (or an unsynced pipelined message, or `COPY FROM STDIN`) until the transaction ends, and other connections wait. Connection pools of any size work, but they serialize instead of running in parallel. Parallelism comes from forks.

The consequences are the usual transaction-pooling ones:

- **Session state is shared** between live connections: `SET`, temporary tables and advisory locks. Use `SET LOCAL`. Prepared statements are safe, because their names are prefixed per connection. When a connection starts and no other is alive, the session is reset as a new backend would be.
- **`LISTEN` and `NOTIFY` work per connection.** Notifications are routed to the connections that listen on the channel.
- **Cross-connection waits hang.** A connection that waits inside a transaction for work another connection must do first, such as a row lock or an advisory lock, waits forever. pgmem logs a diagnostic after five seconds.
- **`snapshot` waits for open transactions.** Commit or close every connection to the template before taking a snapshot; the wrappers fail with `busy` after 30 seconds.

## Databases

`CREATE DATABASE` and `DROP DATABASE` work, but a server serves only the database it was started with. A connection that names another database is refused with SQLSTATE 3D000. Tools that create and connect to a second database, such as a shadow database for migration diffing, need a second server.

## Features that need background processes

- Parallel query and parallel index builds are off (`max_parallel_workers=0`).
- Extensions that run background workers are not available. See [extensions](../extensions/).
- `io_method` is forced to `sync`.

## Resources

- A server is about 150 MB resident with the default `shared_buffers=32MB`. Each live fork adds its own copy. Raise `shared_buffers` through the server parameters when a test needs a larger cache.
- The server is 32-bit WebAssembly and addresses at most 4 GiB.
- `statement_timeout` fires between protocol messages and whenever the backend reads the clock or sleeps, but not inside a CPU loop that does neither.

## Security

There is no TLS and no password authentication. Every server listens on `127.0.0.1` only, and URLs carry `sslmode=disable`. ICU collations are not available.
