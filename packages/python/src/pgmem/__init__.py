"""In-memory PostgreSQL for tests.

    import pgmem

    with pgmem.start(database="app") as pg:
        run_migrations(pg.template.dsn)
        snap = pg.template.snapshot()
        with snap.fork() as fork:
            conn = psycopg.connect(fork.dsn)   # a private copy of the prepared database

With pytest installed the ``pgmem_dsn`` fixture does the fork/close per
test; override ``pgmem_snapshot`` to run migrations first.
"""

from ._client import (
    Fork,
    Pgmem,
    PgmemError,
    ProtocolError,
    Server,
    ServerExited,
    Snapshot,
    start,
)
from ._binary import find_binary

__all__ = [
    "Fork",
    "Pgmem",
    "PgmemError",
    "ProtocolError",
    "Server",
    "ServerExited",
    "Snapshot",
    "find_binary",
    "start",
]
