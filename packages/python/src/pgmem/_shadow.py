"""Temporary PostgreSQL driver routing for the pytest marker.

Patch only installed drivers. Connections created while the marker is active
use its prepared fork, even when the caller supplies a production DSN.
"""

import importlib


_PSYCOPG_TARGET_KEYS = {
    "host", "hostaddr", "port", "dbname", "database", "user", "password",
    "service", "servicefile", "passfile", "sslmode", "sslcert", "sslkey",
    "sslrootcert", "sslcrl", "sslcrldir", "target_session_attrs",
}
_ASYNCPG_TARGET_KEYS = {
    "host", "port", "user", "password", "passfile", "service", "servicefile",
    "database", "ssl", "direct_tls", "target_session_attrs",
}


def _drop_target_options(kwargs, keys):
    return {key: value for key, value in kwargs.items() if key not in keys}


def route_environment(monkeypatch, target):
    """Set PostgreSQL URL and libpq variables for one target."""
    monkeypatch.setenv("DATABASE_URL", target.dsn)
    monkeypatch.setenv("PGHOST", target.host)
    monkeypatch.setenv("PGPORT", str(target.port))
    monkeypatch.setenv("PGUSER", target.user)
    monkeypatch.setenv("PGDATABASE", target.database)
    monkeypatch.setenv("PGSSLMODE", "disable")


def install_driver_routes(monkeypatch, target):
    """Route newly created psycopg, psycopg2 and asyncpg connections to target."""
    try:
        psycopg = importlib.import_module("psycopg")
    except ModuleNotFoundError as exc:
        if exc.name != "psycopg":
            raise
    else:
        original_sync = psycopg.Connection.connect.__func__
        original_async = psycopg.AsyncConnection.connect.__func__

        def sync_connect(cls, conninfo="", **kwargs):
            return original_sync(cls, target.dsn, **_drop_target_options(kwargs, _PSYCOPG_TARGET_KEYS))

        async def async_connect(cls, conninfo="", **kwargs):
            return await original_async(cls, target.dsn, **_drop_target_options(kwargs, _PSYCOPG_TARGET_KEYS))

        monkeypatch.setattr(psycopg.Connection, "connect", classmethod(sync_connect))
        monkeypatch.setattr(psycopg.AsyncConnection, "connect", classmethod(async_connect))
        monkeypatch.setattr(psycopg, "connect", psycopg.Connection.connect)

    try:
        psycopg2 = importlib.import_module("psycopg2")
    except ModuleNotFoundError as exc:
        if exc.name != "psycopg2":
            raise
    else:
        original = psycopg2.connect

        def connect(dsn=None, **kwargs):
            return original(target.dsn, **_drop_target_options(kwargs, _PSYCOPG_TARGET_KEYS))

        monkeypatch.setattr(psycopg2, "connect", connect)

    try:
        asyncpg = importlib.import_module("asyncpg")
    except ModuleNotFoundError as exc:
        if exc.name != "asyncpg":
            raise
    else:
        original = asyncpg.connection.connect

        async def connect(dsn=None, **kwargs):
            return await original(target.dsn, **_drop_target_options(kwargs, _ASYNCPG_TARGET_KEYS))

        monkeypatch.setattr(asyncpg.connection, "connect", connect)
        monkeypatch.setattr(asyncpg, "connect", connect)


def install_sqlalchemy_pool_guard(target):
    """Invalidate pooled PostgreSQL connections from an earlier test fork."""
    try:
        sqlalchemy = importlib.import_module("sqlalchemy")
    except ModuleNotFoundError as exc:
        if exc.name != "sqlalchemy":
            raise
        return lambda: None

    def is_postgres(connection):
        module = type(connection).__module__
        return module.startswith(("psycopg", "sqlalchemy.dialects.postgresql.asyncpg"))

    def on_connect(connection, record):
        if is_postgres(connection):
            record.info["shadow_pg_target"] = target.id

    def on_checkout(connection, record, proxy):
        if is_postgres(connection) and record.info.get("shadow_pg_target") != target.id:
            raise sqlalchemy.exc.DisconnectionError("pgmem: pooled connection belongs to another test")

    sqlalchemy.event.listen(sqlalchemy.pool.Pool, "connect", on_connect)
    sqlalchemy.event.listen(sqlalchemy.pool.Pool, "checkout", on_checkout)

    def remove():
        sqlalchemy.event.remove(sqlalchemy.pool.Pool, "checkout", on_checkout)
        sqlalchemy.event.remove(sqlalchemy.pool.Pool, "connect", on_connect)

    return remove
