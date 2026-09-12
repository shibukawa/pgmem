---
id: concept:seeding
type: concept
title: Seed Data on pgmem
---
Load fixture rows once into the template after migrations so every fork starts with the same data (requirement:test-fixture-fork).

```yaml
summary:
  where: the same prepare step as migrations (concept:migration-tools); after it, api:snapshot
  raw_sql: INSERT statements from a seed.sql file executed on the template DSN
  copy: 'COPY t FROM STDIN with the driver''s copy API (pgx CopyFrom, psycopg copy, pgjdbc CopyManager); a connection owns the session for the whole COPY'
  libraries:
    go: go-testfixtures loads YAML files with a *sql.DB, run it in Options.Prepare
    python: factory_boy or plain fixtures against pgmem_server.dsn inside the pgmem_snapshot override
    java: DbUnit or a hand-written loader against t.jdbcUrl() in prepare
  per_test_data: rows a single test needs go into the fork, not the template; the fork is discarded (decision:fork-or-not)
  several_seed_sets: one template per set (api:control-protocol op start, Java builder().template)
  identity: sequences advance in the template and continue in every fork from the same value, so ids in forks are deterministic
  size: seed data grows the vfs and therefore every fork copy (metric:fork-cost); keep templates small and load bulk data per test only when needed
```
