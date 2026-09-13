---
title: Limits
description: What pgmem does differently from a PostgreSQL server started by a postmaster.
---

## One session per server

PostgreSQL runs in single-user mode, so every connection to a server or fork shares one backend session. pgmem multiplexes connections the way a transaction-mode pooler does: a connection owns the backend from `BEGIN` (or an unsynced pipelined message, or `COPY FROM STDIN`) until the transaction ends, and other connections wait. Connection pools of any size work, but they serialize instead of running in parallel. Parallelism comes from forks.

The consequences are the usual transaction-pooling ones:

- **Session state is shared** between live connections: `SET`, temporary tables and advisory locks. Use `SET LOCAL`. Prepared statements are safe, because their names are prefixed per connection. When a connection starts and no other is alive, the session is reset as a new backend would be.
- **`LISTEN` and `NOTIFY` work per connection.** Notifications are routed to the connections that listen on the channel.
- **Waits behind an idle transaction end.** A connection cannot run while another is inside a transaction. When the holder sits idle in its transaction while another connection waits, for example code that queries through the pool instead of the transaction handle inside a transaction callback, the waiting connection is ended after `WaitTimeout` (default two seconds) with SQLSTATE 55P03 and a message naming the holder. A statement that is merely slow is waited for.
- **`snapshot` waits for open transactions.** Commit or close every connection to the template before taking a snapshot; the wrappers fail with `busy` after 30 seconds.

## Databases

Every database of a server can be connected to, and `CREATE DATABASE` works. The backend serves one database at a time: a connection to another database restarts the backend on it, in well under 10 ms. Connections of the database that stopped being served get their prepared statements and `LISTEN` registrations back when it is served again, but not their `SET` values or temporary tables. Prisma's shadow database needs nothing more; connections that keep alternating between databases pay a restart each time. A database that does not exist is refused with SQLSTATE 3D000.

## Closing a server

Closing a server or fork does not drop client connections. Each stays open until its client closes it, sends a message, which is answered with SQLSTATE 57P01, or 30 seconds pass. Pools such as node-postgres's raise an unhandled error when an idle connection is dropped under them.

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
