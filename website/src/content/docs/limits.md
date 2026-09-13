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

## Fork capacity and waiting

Each snapshot has its own pool of fork slots. `MaxForks` limits live forks from that snapshot; the template server does not use a slot. The default is the number of available CPUs. Each fork owns another data-directory copy and buffer cache, so the cap also limits memory use. When all slots are occupied, another fork request waits until a fork closes and releases its slot.

| API | Set the limit | Limit a wait for a free slot |
|---|---|---|
| Go | `SnapshotOptions.MaxForks` or `pgmemtest.Options.MaxForks` | `Snapshot.Fork(ctx)` waits until its context is cancelled or reaches its deadline. If the server logger is enabled, a wait longer than 5 seconds is logged. |
| Python | `server.snapshot(max_forks=n)` | `snapshot.fork(timeout=seconds)`; omitted or `None` waits indefinitely. Expiration raises `ProtocolError` with code `pool_timeout`. |
| Java | `Server.snapshot(maxForks)` or JUnit's `.maxForks(n)` | `Snapshot.fork(Duration)` or JUnit's `.forkTimeout(Duration)`; no timeout is the default. Expiration raises `pool_timeout`. |
| Node.js | `PgmemServer.start({ maxForks: n })` | `fork({ timeoutMs })` or `withFork(fn, { timeoutMs })`; omitting it waits indefinitely. Expiration raises `PgmemError` with code `pool_timeout`. |

The JUnit extension applies its cap to each registered template. A snapshot's `close()` refuses new forks but leaves live forks running; Go's `Snapshot.Wait()` waits for those forks to close.

## Timeouts govern different waits

The timeout for a fork slot does not limit snapshot creation or a database connection's wait behind another connection. Those are separate waits with separate settings:

| Wait | Setting | What happens at the deadline |
|---|---|---|
| Snapshot creation waits for open transactions to finish | Go: the `Snapshot` context. Python: `server.snapshot(timeout=30.0)`. Java: `Server.snapshot(maxForks, timeout)` (30 seconds by default). Node.js: `snapshotTimeoutMs` (30,000 ms by default). | Snapshot creation fails with `busy`; commit or close connections to the template first. |
| A connection waits behind another connection's idle transaction | Go: `Options.WaitTimeout`; Node.js: `waitTimeoutMs` (2,000 ms by default). | The waiting connection ends with SQLSTATE 55P03. This is backend-session contention, not a full fork pool. |
| A server process starts | Python: `pgmem.start(timeout=30)`. Java: `.readyTimeout(Duration)` (30 seconds by default). Node.js: `startupTimeoutMs` (30,000 ms by default). Go: the `Start` context. | Startup fails if the process does not become ready in time. |
| A fork is reset to a snapshot | Go: `Server.Reset(ctx)` or `Restore(ctx)`. Node.js: `fork.reset({ timeoutMs })` (5,000 ms by default). | Reset fails with `busy` if an open transaction does not finish. |

## Security

There is no TLS and no password authentication. Every server listens on `127.0.0.1` only, and URLs carry `sslmode=disable`. ICU collations are not available.
