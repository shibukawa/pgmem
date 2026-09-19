---
id: concept:limits
type: concept
title: Limits
---
What pgmem cannot do or does differently from a postmaster-based PostgreSQL; the limits page and the caveat box in every guide.

```yaml
summary:
  process_model: rule:process-per-connection; sessions, locks, deadlock detection (40P01) and session state are PostgreSQL's own; parallelism within a fork as on a server, across forks for isolation (requirement:test-fixture-fork)
  restart_survival: Snapshot and Reset restart the cluster; client sockets survive and re-attach with prepared statements and LISTEN replayed; a backend PostgreSQL itself ends (FATAL) closes the connection
  close_keeps_clients: Server.Close leaves client connections open until the client closes, sends a message (57P01) or 30 s pass, so pools without an error listener do not crash
  databases: every database of a server can be connected to, each connection to its own backend; a missing database gets 3D000
  extensions: the bundled set only (policy:bundled-extensions); no ICU, OpenSSL or zlib, pgcrypto runs its crypto on Go (decision:pgcrypto-on-host), fips_mode() is always false
  parallel_query: works; workers are processes of the cluster
  io_method: defaults to sync; worker would add three I/O worker processes for no gain
  memory: ~150 MB resident per server plus a few MB per process (metric:server-footprint); Options.Params can raise shared_buffers when a test needs cache
  timeouts: statement_timeout and query cancel are delivered at host calls (clock reads, I/O, waits), not inside a pure CPU loop
  address_space: 32-bit wasm, at most 4 GiB per server
  tls_and_auth: trust auth, no password, no TLS; DSNs carry sslmode=disable (system:postgres-drivers)
  persistence: nothing on disk (requirement:in-memory-only); export of a snapshot is a later feature
```
