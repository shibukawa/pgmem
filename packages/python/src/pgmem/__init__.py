"""In-memory PostgreSQL for tests.

    import pgmem

    with pgmem.start(database="app") as pg:
        run_migrations(pg.template.dsn)
        snap = pg.template.snapshot()
        with snap.fork() as fork:
            conn = psycopg.connect(fork.dsn)   # a private copy of the prepared database

With pytest installed, ``pgmem_test_dsn`` selects the isolation mode per
test. Override ``pgmem_prepare`` to run migrations once before snapshotting.
"""

from ._client import (
    CleanupError,
    readonly_dsn,
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


def shadow_pg(func=None, *, fork=True, prepare=False, live=False):
    """Route a pytest test to a fork, or a preparation fixture to the template.

    ``fork=False`` shares one fork for read-only tests. ``live=True`` resets a
    stable endpoint before each in-process HTTP test. New psycopg and asyncpg
    connections are routed without changing application configuration.
    ``prepare=True`` wraps a session-scoped ``pgmem_prepare(pgmem_server)``
    fixture before the snapshot is taken.
    """
    if not isinstance(fork, bool):
        raise TypeError("fork must be a bool")
    if not isinstance(prepare, bool):
        raise TypeError("prepare must be a bool")
    if not isinstance(live, bool):
        raise TypeError("live must be a bool")
    if live and (prepare or not fork):
        raise TypeError("live=True cannot be combined with prepare=True or fork=False")
    import pytest

    if prepare:
        if not fork:
            raise TypeError("fork=False cannot be used while preparing the template")
        import functools
        import inspect

        from ._shadow import install_driver_routes, install_sqlalchemy_pool_guard

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

    marker = pytest.mark.shadow_pg(fork=fork, live=live)
    return marker(func) if func is not None else marker

__all__ = [
    "CleanupError",
    "readonly_dsn",
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
