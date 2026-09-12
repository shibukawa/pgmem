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

The wheel registers a plugin. Override `pgmem_snapshot` to run migrations,
then use `pgmem_dsn` (or `pgmem_fork`) in tests:

```python
# conftest.py
@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server):
    run_migrations(pgmem_server.dsn)
    return pgmem_server.snapshot()

# test_orders.py
def test_orders(pgmem_dsn):
    with psycopg.connect(pgmem_dsn) as conn: ...
```

`pgmem_class_dsn` shares one fork across a test class for read-only tests.
`pgmem_options` returns the keyword arguments passed to `pgmem.start`.

Any driver works (psycopg, asyncpg, pg8000, SQLAlchemy): the wrapper only
hands out DSNs. The server is one PostgreSQL session per fork; pooled
connections are serialized at transaction boundaries.

## Binary

Platform wheels bundle the `pgmem` binary. `PGMEM_BINARY` overrides the
lookup for platforms without a wheel or for local builds
(`go build ./cmd/pgmem`).
