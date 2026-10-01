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


def shadow_pg(func=None, *, fork=True, prepare=False):
    """Route a pytest test to a fork, or a preparation fixture to the template.

    ``fork=False`` shares one fork for read-only tests. Construct database
    clients after test setup; clients made during test-module import keep the
    URL they captured at import time. ``prepare=True`` wraps a session-scoped
    ``pgmem_prepare(pgmem_server)`` fixture before the snapshot is taken.
    """
    if not isinstance(fork, bool):
        raise TypeError("fork must be a bool")
    if not isinstance(prepare, bool):
        raise TypeError("prepare must be a bool")
    import pytest

    if prepare:
        if not fork:
            raise TypeError("fork=False cannot be used while preparing the template")
        import functools
        import inspect

        from ._shadow import install_driver_routes, install_sqlalchemy_pool_guard, route_environment

        def decorate_fixture(fn):
            signature = inspect.signature(fn)
            if "pgmem_server" not in signature.parameters:
                raise TypeError("shadow_pg(prepare=True) requires a pgmem_server fixture parameter")
            if inspect.isgeneratorfunction(fn) or inspect.iscoroutinefunction(fn):
                raise TypeError("pgmem_prepare must be a synchronous fixture without yield")

            @functools.wraps(fn)
            def wrapped(*args, **kwargs):
                server = signature.bind(*args, **kwargs).arguments["pgmem_server"]
                with pytest.MonkeyPatch.context() as patch:
                    route_environment(patch, server)
                    install_driver_routes(patch, server)
                    remove_pool_guard = install_sqlalchemy_pool_guard(server)
                    try:
                        result = fn(*args, **kwargs)
                        if inspect.isgenerator(result) or inspect.isawaitable(result):
                            raise TypeError("pgmem_prepare must be a synchronous fixture without yield")
                        return result
                    finally:
                        remove_pool_guard()

            return wrapped

        return decorate_fixture(func) if func is not None else decorate_fixture

    marker = pytest.mark.shadow_pg(fork=fork)
    return marker(func) if func is not None else marker

__all__ = [
    "Fork",
    "Pgmem",
    "PgmemError",
    "ProtocolError",
    "Server",
    "ServerExited",
    "Snapshot",
    "find_binary",
    "shadow_pg",
    "start",
]
