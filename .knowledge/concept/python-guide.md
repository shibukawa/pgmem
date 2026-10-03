---
id: concept:python-guide
type: concept
title: Python Guide
---
Using pgmem from Python tests: install the wheel, the pytest fixtures, prepare-once hooks and automatic snapshots, and fork styles.

```yaml
summary:
  install: pip install pgmem (platform wheel with the binary, policy:binary-distribution); PGMEM_BINARY for other platforms; python >= 3.9
  basic_test: |
    import pgmem, psycopg
    with pgmem.start(database="app") as pg:
        with psycopg.connect(pg.template.dsn) as conn:
            conn.execute("select 1")
  initialization:
    fixture: override session pgmem_prepare to return a callback accepting the template Server; run migrations and seed via its DSN; default pgmem_snapshot invokes it then snapshots automatically (api:python-wrapper)
    advanced: override pgmem_snapshot for multiple snapshots or custom preparation
    raw_sql: conn.execute(open("schema.sql").read()) on the template DSN
    tools: Alembic, SQLAlchemy create_all, yoyo; Django to verify (concept:migration-tools); seed with factory_boy or plain SQL (concept:seeding)
  test_case_styles:
    fork_per_test: pgmem_dsn or pgmem_fork function fixtures; closed in teardown
    fork_per_class: pgmem_class_dsn for read-only test classes
    selective: pgmem_test_dsn uses @pytest.mark.pgmem(fork=False) or isolation=shared for read-only template access; default isolated fork
    persistent_clients: pgmem_shared_dsn plus isolation=reset keeps a stable endpoint and resets in teardown; serial within worker
    guidance: decision:fork-or-not; requirement:python-selective-isolation implemented without driver connection management
  drivers: psycopg, asyncpg, pg8000, SQLAlchemy all take the DSN (system:postgres-drivers); build a per-driver fixture on pgmem_dsn
  async: asyncpg works against the DSN; the client itself is sync and runs in a thread from asyncio fixtures
  parallel: pytest-xdist workers each spawn their own binary (concept:server-process); inside one worker forks are independent
  pools: SQLAlchemy pools work as on a server (rule:process-per-connection); commit or close before snapshot() or it fails with busy after 30 s
  limits: concept:limits
```
