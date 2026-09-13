---
id: rule:single-session-per-backend
type: rule
title: Single Session per Backend
---
PostgreSQL runs in single-user mode, so each backend serves exactly one session; concurrent connections to the same backend would interleave state.

```yaml
rule:
  sessions_per_backend: 1
  multiplexing: transaction-mode pooling in pgmem.go; a connection owns the backend from BEGIN (or unsynced pipeline, or COPY IN) until ReadyForQuery reports idle; others wait
  pools: any size works but serializes; api:clone handles no longer need max 1 connection
  shared_session_state: SET, temp tables, advisory locks leak between live connections; prepared statements are prefixed per connection and dropped at its end
  listen_notify: per connection; async.c commit hook (wasm/patches.py, host import pgmem_listen) feeds a channel->sessions registry in pgmem.go, NotifyResponse messages are routed to listening sessions through a per-session queue, and LISTEN is re-issued when one connection's UNLISTEN drops a channel others still want
  fresh_session: at connection start, when no other connection is alive, pgmem_reset_session (C, exported from postgres.c) aborts any transaction, does what DISCARD ALL does and forgets the temp namespace; no statement runs, so pg_stat_statements and log_statement do not see it
  deadlock: a transaction waiting on another connection's work waits forever; diagnostic logged after 5s; the common ORM form is using the outer client inside a transaction callback (system:node-orms); Options.WaitTimeout (default 2s, pgmem -wait-timeout) ends the waiter with 55P03 once the holder has sat idle in its transaction that long, while slow statements are waited for (implemented 2026-09-13)
  parallelism: obtained across forks, not within a fork (requirement:test-fixture-fork)
```
