# pgmem for Python

A real PostgreSQL 18 that runs entirely in memory, packaged as a single
binary and driven from Python. Prepare the schema once, then give every
test its own fork of that state in about 20 ms.

```python
import pgmem, psycopg

with pgmem.start(database="app") as pg:
    with psycopg.connect(pg.template.dsn) as conn:
        conn.execute(open("schema.sql").read())
    snap = pg.template.snapshot()
    with snap.fork() as fork:            # private copy of the prepared database
        psycopg.connect(fork.dsn) ...
```

## pytest

The wheel registers a plugin. Define a session-scoped `pgmem_prepare` fixture
to run migrations and seed data once. `@shadow_pg(prepare=True)` routes its
normal PostgreSQL SDK calls to the template; the plugin snapshots it after
the fixture returns:

```python
# conftest.py
from pgmem import shadow_pg

@pytest.fixture(scope="session")
@shadow_pg(prepare=True)
def pgmem_prepare(pgmem_server):
    run_migrations()  # may use psycopg or SQLAlchemy; use asyncio.run for asyncpg
    load_seed()

# test_orders.py
def test_orders(pgmem_dsn):
    with psycopg.connect(pgmem_dsn) as conn: ...
```

Keep `@pytest.fixture` outside `@shadow_pg(prepare=True)` as shown. The
`pgmem_server` parameter supplies the template to the decorator; preparation
code can ignore it. Return from the fixture after committing or closing its
connections. Override `pgmem_snapshot` directly when you need custom snapshot
options such as `max_forks`.

`pgmem_class_dsn` shares one fork across a test class for read-only tests.
`pgmem_options` returns the keyword arguments passed to `pgmem.start`.

Any driver works (psycopg, asyncpg, pg8000, SQLAlchemy): the wrapper only
hands out DSNs. Each connection has its own PostgreSQL backend; pools, transactions and locks
behave as on a server.

## Select isolation without changing your application fixture

Make a function-scoped application fixture depend on `pgmem_test_dsn`.
It uses a fresh fork by default. `@pytest.mark.pgmem(fork=False)` selects
the shared prepared template for a read-only test.

Override session-scoped `pgmem_prepare` with a callback accepting the template
server to migrate and seed once; the plugin then creates the snapshot.
`pgmem_snapshot` remains available for custom snapshot management.

For a session-scoped client, use `pgmem_shared_dsn` and mark its serial tests
with `@pytest.mark.pgmem(isolation="reset")`. The plugin restores its shared
fork after each marked test, including failed tests. Commit or roll back open
transactions before reset. `Server.reset(snapshot=None, timeout=5.0)` is also
available for manual lifecycles.

`pgmem_fixture_options` configures isolation, `max_forks`, `fork_timeout`
(default 30 seconds), `snapshot_timeout` (30 seconds) and `reset_timeout`
(5 seconds). See the [complete pytest quickstart](../../website/src/content/docs/guides/python/testing.mdx).

### Route a test with `@shadow_pg`

Prepare the schema and seeds once in `pgmem_prepare` as above. The marker
intercepts new psycopg 3, psycopg2 and asyncpg connections, including
SQLAlchemy engines using those drivers. Application code can keep its usual
connection arguments:

```python
import psycopg
from pgmem import shadow_pg

@shadow_pg
def test_create_order():
    with psycopg.connect("postgresql://app@db.example/app") as db:
        ...  # connects to a fresh pgmem fork

@shadow_pg(fork=False)
def test_lookup():
    with psycopg.connect("postgresql://app@db.example/app") as db:
        ...  # reads from one prepared fork shared in this pytest session
```

It also sets `DATABASE_URL` and libpq's `PGHOST`, `PGPORT`, `PGUSER`,
`PGDATABASE`, and `PGSSLMODE` for code that reads those variables; pytest
restores them afterward. `fork=False` is for read-only tests: writes to the
shared fork affect later tests. The database does not enforce read-only
access. SQLAlchemy can reuse an Engine created earlier; pooled connections
from a previous fork are invalidated on checkout. Connections already
checked out, asyncpg pools created before the test, and functions copied
with `from psycopg import connect` before the marker runs cannot be rerouted.
Use an explicit `pgmem_dsn` fixture for those patterns.

Other drivers still work through the `pgmem_dsn` fixture. The wrapper does
not replace PostgreSQL's wire protocol or the drivers themselves.

### Long-lived HTTP application

For FastAPI, Starlette, Django, or another application that keeps a connection
pool alive across tests, use one stable endpoint per pytest worker. Configure
the application once with `pgmem_live_fork.dsn` and request `pgmem_live_db`
in every sequential HTTP test. The latter calls `reset()` before the test,
restoring the prepared snapshot without changing the endpoint or client
connections. Override `pgmem_live_snapshots` with a mapping of names to
snapshots and select one with `@pytest.mark.pgmem_dataset("name")`.

If the application creates its schema while it starts, with Alembic in a
lifespan handler, Django `migrate`, or SQLAlchemy `create_all`, request
`pgmem_live_baseline` where the application is started and call `capture()`
once it is up. Later resets restore that state instead of the prepared
snapshot, and named datasets keep their own snapshots. The captured snapshot
holds one more copy of the data directory for the session.

An open transaction blocks reset and raises `ProtocolError` with code `busy`
after the timeout. The live fork holds one snapshot fork slot for the session;
raise `max_forks` if other forks coexist. For Django, use plain pytest tests
with its request Client or a separately launched app; avoid `TestCase` and
pytest-django database fixtures that manage rollback or flushing themselves.

## Binary

Platform wheels bundle the `pgmem` binary. `PGMEM_BINARY` overrides the
lookup for platforms without a wheel or for local builds
(`go build ./cmd/pgmem`).

## Shared read guard and cleanup / 共有接続のガードと後片付け

`pgmem_fixture_options` accepts `shared_read_only: True` to guard selected
shared and class DSNs. Isolated writing forks and the stable `pgmem_shared_dsn`
used by reset mode remain writable. `readonly_dsn(dsn)` creates a guarded URL
explicitly. This sets a session default, which callers can deliberately disable.
Context managers preserve a primary error and cleanup failures through
`CleanupError.primary_error` and `.cleanup_errors` on Python 3.9+.

`pgmem_fixture_options` に `shared_read_only: True` を指定すると、選択した共有 DSN
とクラス用 DSN を保護します。独立した書き込み用フォークと reset 用の
`pgmem_shared_dsn` は書き込み可能です。URL を自分で作る場合は `readonly_dsn(dsn)`
を使えます。意図的に無効化できるセッションの既定値です。context manager は
Python 3.9 以降で、元のエラーと解放時の失敗を `CleanupError.primary_error` と
`.cleanup_errors` に残します。

Runnable SQLAlchemy application / 実行可能な SQLAlchemy アプリ:
[examples](../../examples/README.md).
