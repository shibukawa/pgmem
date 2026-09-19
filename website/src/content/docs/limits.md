---
title: Limits
description: What pgmem does differently from a PostgreSQL server started by a postmaster.
---

## Process model

PostgreSQL runs as it does on a server: a postmaster starts a backend process for every connection, plus its auxiliary processes, and they share memory. Every process is a WebAssembly instance on its own goroutine. Sessions are therefore real: locks between connections wait, deadlocks are detected (SQLSTATE 40P01 after `deadlock_timeout`), `pg_stat_activity` and `pg_locks` show the other sessions, and session state (`SET`, temporary tables, advisory locks) belongs to one connection.

- **`snapshot` and `reset` wait for open transactions.** Both stop the cluster and start it again; client connections keep their sockets and get a new backend on their next message, with named prepared statements and `LISTEN` registrations re-created. Commit or close transactions first; the wrappers fail with `busy` after their timeout.
- **A backend PostgreSQL itself ends** (`pg_terminate_backend`, `DROP DATABASE ... WITH (FORCE)`, a crash) ends the connection with its FATAL message, as on a server.
- **Cancelling a statement that makes no system call** waits until it does.

## Databases

Every database of a server can be connected to, and `CREATE DATABASE` works. A database that does not exist is refused with SQLSTATE 3D000.

## Closing a server

Closing a server or fork does not drop client connections. Each stays open until its client closes it, sends a message, which is answered with SQLSTATE 57P01, or 30 seconds pass. Pools such as node-postgres's raise an unhandled error when an idle connection is dropped under them.

## Background processes

- Parallel query and parallel index builds work; the workers are processes of the cluster.
- Extensions that need their own background worker are not bundled. See [extensions](../extensions/).
- `io_method` defaults to `sync` (PostgreSQL 18's `worker` would add three I/O worker processes for no gain).

## Resources

- A server is about 150 MB resident with the default `shared_buffers=32MB`, plus a few MB per process (five auxiliary processes and one per connection). Each live fork adds its own copy. Raise `shared_buffers` through the server parameters when a test needs a larger cache.
- The server is 32-bit WebAssembly and addresses at most 4 GiB.
- `statement_timeout` fires between protocol messages and whenever the backend reads the clock or sleeps, but not inside a CPU loop that does neither.

## Fork capacity and waiting

Each snapshot has its own pool of fork slots. `MaxForks` limits live forks from that snapshot; the template server does not use a slot. The default comes from memory: a quarter of the process's memory limit (`GOMEMLIMIT`, the cgroup limit or the physical memory, whichever is smallest) divided by the cost of one fork (`shared_buffers` plus about 32 MB), which is 28 on a 7 GB CI runner with the default `shared_buffers=32MB`. A slot costs nothing until a fork occupies it. When the memory cannot be determined, the default is the number of CPUs. Each fork owns another data-directory copy and buffer cache, so the cap also limits memory use. When all slots are occupied, another fork request waits until a fork closes and releases its slot.

| API | Set the limit | Limit a wait for a free slot |
|---|---|---|
| Go | `SnapshotOptions.MaxForks` or `pgmemtest.Options.MaxForks` | `Snapshot.Fork(ctx)` waits until its context is cancelled or reaches its deadline. If the server logger is enabled, a wait longer than 5 seconds is logged. |
| Python | `server.snapshot(max_forks=n)` | `snapshot.fork(timeout=seconds)`; omitted or `None` waits indefinitely. Expiration raises `ProtocolError` with code `pool_timeout`. |
| Java | `Server.snapshot(maxForks)` or JUnit's `.maxForks(n)` | `Snapshot.fork(Duration)` or JUnit's `.forkTimeout(Duration)`; no timeout is the default. Expiration raises `pool_timeout`. |
| Node.js | `PgmemServer.start({ maxForks: n })` | `fork({ timeoutMs })` or `withFork(fn, { timeoutMs })`; omitting it waits indefinitely. Expiration raises `PgmemError` with code `pool_timeout`. |

The JUnit extension applies its cap to each registered template. A snapshot's `close()` refuses new forks but leaves live forks running; Go's `Snapshot.Wait()` waits for those forks to close.

## Timeouts govern different waits

The timeout for a fork slot does not limit snapshot creation or server start-up. Those are separate waits with separate settings:

| Wait | Setting | What happens at the deadline |
|---|---|---|
| Snapshot creation waits for open transactions to finish | Go: the `Snapshot` context. Python: `server.snapshot(timeout=30.0)`. Java: `Server.snapshot(maxForks, timeout)` (30 seconds by default). Node.js: `snapshotTimeoutMs` (30,000 ms by default). | Snapshot creation fails with `busy`; commit or close connections to the template first. |
| A server process starts | Python: `pgmem.start(timeout=30)`. Java: `.readyTimeout(Duration)` (30 seconds by default). Node.js: `startupTimeoutMs` (30,000 ms by default). Go: the `Start` context. | Startup fails if the process does not become ready in time. |
| A fork is reset to a snapshot | Go: `Server.Reset(ctx)` or `Restore(ctx)`. Node.js: `fork.reset({ timeoutMs })` (5,000 ms by default). | Reset fails with `busy` if an open transaction does not finish. |

## Security

There is no TLS and no password authentication. Every server listens on `127.0.0.1` only, and URLs carry `sslmode=disable`. ICU collations are not available.
