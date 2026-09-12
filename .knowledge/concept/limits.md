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
  databases: only Options.Database is served; the README of 2026-09-12 still says CREATE DATABASE hangs, a later probe showed it completes in ~10 ms, and refusing connections to other databases with 3D000 plus multi-database serving are in progress
  extensions: plpgsql only (concept:static-modules); no ICU, OpenSSL, zlib, pgcrypto
  io_method: forced to sync; worker AIO would wait on IO workers that do not exist
  memory: ~150 MB resident per server (metric:server-footprint); Options.Params can raise shared_buffers when a test needs cache
  timeouts: statement_timeout fires between protocol messages and on clock reads or sleeps, not inside a pure CPU loop
  address_space: 32-bit wasm, at most 4 GiB per server
  tls_and_auth: trust auth, no password, no TLS; DSNs carry sslmode=disable (system:postgres-drivers)
  persistence: nothing on disk (requirement:in-memory-only); export of a snapshot is a later feature
```
