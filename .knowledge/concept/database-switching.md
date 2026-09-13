---
id: concept:database-switching
type: concept
title: Database Switching
---
A server serves every database of its data directory, one at a time, by restarting the backend on the database a connection names (databases.go); added 2026-09-13 so Prisma's shadow database works with only DATABASE_URL.

```yaml
summary:
  trigger: a startup or batch from a session whose database is not the served one, once it holds the backend (rule:single-session-per-backend)
  switch: shutdown checkpoint of the old backend, a new backend on the same vfs, replay of that database's first startup packet, SET SESSION AUTHORIZATION, its LISTENs re-issued and Parse replayed for its live sessions; api:reset shares this restart
  cost: about 7ms per switch (TestOtherDatabases, 10 alternating queries in 68ms)
  per_database_state: first startup packet and response, live session count; listeners keyed by database and channel
  missing_database: backend start fails in about 5ms, the previous database is started again, the connection gets 3D000; the same for a database dropped under a live connection
  lost_on_switch: SET values, temp tables and session advisory locks of the database that stops being served
  verified: TestOtherDatabases, TestRestoreWhileServingAnotherDatabase; Prisma 7.10 migrate dev without shadowDatabaseUrl (system:node-orms)
  replaces: refusing other databases with 3D000 (2026-09-12)
```
