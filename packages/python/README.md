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

An open transaction blocks reset and raises `ProtocolError` with code `busy`
after the timeout. The live fork holds one snapshot fork slot for the session;
raise `max_forks` if other forks coexist. For Django, use plain pytest tests
with its request Client or a separately launched app; avoid `TestCase` and
pytest-django database fixtures that manage rollback or flushing themselves.

## Binary

Platform wheels bundle the `pgmem` binary. `PGMEM_BINARY` overrides the
lookup for platforms without a wheel or for local builds
(`go build ./cmd/pgmem`).
