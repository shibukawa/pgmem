---
id: concept:limits
type: concept
title: Limits
---
What pgmem cannot do or does differently from a postmaster-based PostgreSQL; the limits page and the caveat box in every guide.

```yaml
summary:
  one_session: rule:single-session-per-backend; pools work but serialize at transaction boundaries; parallelism comes from forks (requirement:test-fixture-fork)
  session_state: SET, temp tables and advisory locks are shared between live connections; use SET LOCAL; prepared statements are safe (per-connection name prefix)
  cross_connection_waits: a transaction waiting for another connection's row lock or advisory lock waits forever; diagnostic logged after 5 s
  databases: CREATE DATABASE and DROP DATABASE work, but only Options.Database is served; a connection naming another database is refused with SQLSTATE 3D000; multi-database serving (Prisma shadow database) is still open
  extensions: the bundled set only (policy:bundled-extensions); no ICU, OpenSSL or zlib, pgcrypto runs its crypto on Go (decision:pgcrypto-on-host), fips_mode() is always false
  parallel_query: max_parallel_workers=0; no postmaster starts workers, so parallel plans and parallel index builds are off
  fresh_session: a connection that starts while no other is alive gets a reset session (everything DISCARD ALL does plus no temp namespace) done in C, invisible to pg_stat_statements
  io_method: forced to sync; worker AIO would wait on IO workers that do not exist
  memory: ~150 MB resident per server (metric:server-footprint); Options.Params can raise shared_buffers when a test needs cache
  timeouts: statement_timeout fires between protocol messages and on clock reads or sleeps, not inside a pure CPU loop
  address_space: 32-bit wasm, at most 4 GiB per server
  tls_and_auth: trust auth, no password, no TLS; DSNs carry sslmode=disable (system:postgres-drivers)
  persistence: nothing on disk (requirement:in-memory-only); export of a snapshot is a later feature
```
