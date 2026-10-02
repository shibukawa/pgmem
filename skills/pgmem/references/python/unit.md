# Python unit tests

Installing `pgmem` registers a pytest plugin. Override session fixtures in `conftest.py`: `pgmem_options` returns `{"database": "app"}`; `pgmem_prepare(pgmem_server)` runs migrations and seed. Decorate that session fixture with `@pytest.fixture(scope="session")` outside `@shadow_pg(prepare=True)` so ordinary psycopg, psycopg2, asyncpg, or SQLAlchemy connections reach the template. Close/commit all preparation connections before the plugin snapshots it.

Decorate a test with `@shadow_pg` to create and clean up a fresh fork without changing the application's connection settings. It intercepts new psycopg 3, psycopg2, and asyncpg connections, including SQLAlchemy engines using those drivers. Read the relevant driver note: [psycopg](unit-psycopg.md), [asyncpg](unit-asyncpg.md), or [SQLAlchemy](unit-sqlalchemy.md). `@shadow_pg(fork=False)` shares one prepared fork only for read-only cases. Direct fixture alternatives are `pgmem_fork`/`pgmem_dsn` per function and `pgmem_class_fork`/`pgmem_class_dsn` per class.

Use pytest-xdist workers for parallelism: each process has its own template and forks. The driver monkeypatch is process-wide within a worker, so do not run marked cases concurrently in threads; wait for background database work before decorator cleanup.

Full driver-specific examples: `website/src/content/docs/guides/python/testing.mdx`.
