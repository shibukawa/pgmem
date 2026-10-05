import pg8000.dbapi
import os
import asyncio
import subprocess
import sys

import asyncpg
import psycopg
from sqlalchemy import create_engine, text
from sqlalchemy.ext.asyncio import create_async_engine

from pgmem import shadow_pg
from pgmem._shadow import install_driver_routes, install_sqlalchemy_pool_guard
import pytest


_cold_engine = create_engine("postgresql+psycopg://bad@127.0.0.1:1/production")


def count(dsn_server):
    conn = pg8000.dbapi.connect(user=dsn_server.user, host=dsn_server.host,
                                port=dsn_server.port, database=dsn_server.database)
    cur = conn.cursor()
    cur.execute("SELECT count(*) FROM t")
    (n,) = cur.fetchone()
    conn.close()
    return n


def test_fork_is_prepared(pgmem_fork, pgmem_dsn):
    assert pgmem_dsn == pgmem_fork.dsn
    assert count(pgmem_fork) == 2
    conn = pg8000.dbapi.connect(user=pgmem_fork.user, host=pgmem_fork.host,
                                port=pgmem_fork.port, database=pgmem_fork.database)
    conn.cursor().execute("INSERT INTO t VALUES (99)")
    conn.commit()
    conn.close()
    assert count(pgmem_fork) == 3


def test_next_test_gets_a_fresh_fork(pgmem_fork):
    assert count(pgmem_fork) == 2


def test_live_endpoint_restores_data_without_changing_port(pgmem_live_db, pgmem_live_fork):
    assert pgmem_live_db is pgmem_live_fork
    port = pgmem_live_db.port
    assert count(pgmem_live_db) == 2
    with psycopg.connect(pgmem_live_db.dsn) as db:
        db.execute("INSERT INTO t VALUES (99)")
    assert count(pgmem_live_db) == 3
    pgmem_live_db.reset()
    assert pgmem_live_db.port == port
    assert count(pgmem_live_db) == 2


class TestClassScoped:
    def test_write(self, pgmem_class_fork):
        conn = pg8000.dbapi.connect(user=pgmem_class_fork.user, host=pgmem_class_fork.host,
                                    port=pgmem_class_fork.port, database=pgmem_class_fork.database)
        conn.cursor().execute("INSERT INTO t VALUES (3)")
        conn.commit()
        conn.close()
        assert count(pgmem_class_fork) == 3

    def test_sees_previous_write(self, pgmem_class_fork, pgmem_class_dsn):
        assert pgmem_class_dsn == pgmem_class_fork.dsn
        assert count(pgmem_class_fork) == 3


@shadow_pg
def test_shadow_routes_to_fresh_fork(pgmem_fork):
    assert os.environ.get("DATABASE_URL") != pgmem_fork.dsn
    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        assert conn.execute("SELECT count(*) FROM t").fetchone() == (2,)


@shadow_pg(fork=False)
def test_shadow_can_share_prepared_target(pgmem_shared_fork):
    assert os.environ.get("DATABASE_URL") != pgmem_shared_fork.dsn
    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        assert conn.execute("SELECT count(*) FROM t").fetchone() == (2,)


@shadow_pg
def test_shadow_intercepts_psycopg_without_application_dsn():
    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        assert conn.execute("SELECT count(*) FROM t").fetchone() == (2,)


@shadow_pg(live=True)
def test_shadow_live_endpoint_without_application_dsn(pgmem_live_fork):
    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        assert conn.execute("SELECT count(*) FROM t").fetchone() == (2,)
        assert conn.info.port == pgmem_live_fork.port
        conn.execute("INSERT INTO t VALUES (99)")


@shadow_pg(live=True)
def test_shadow_live_restores_before_next_test():
    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        assert conn.execute("SELECT count(*) FROM t").fetchone() == (2,)


def test_prepare_fixture_is_snapshotted_without_snapshot_override(tmp_path):
    (tmp_path / "conftest.py").write_text(
        'import pytest\n'
        'from pgmem import shadow_pg\n'
        '@pytest.fixture(scope="session")\n'
        '@shadow_pg(prepare=True)\n'
        'def pgmem_prepare(pgmem_server):\n'
        '    import psycopg\n'
        '    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as db:\n'
        '        db.execute("CREATE TABLE prepared (id int)")\n'
        '        db.execute("INSERT INTO prepared VALUES (7)")\n'
    )
    (tmp_path / "test_prepared.py").write_text(
        'import psycopg\n'
        'from pgmem import shadow_pg\n'
        '@shadow_pg\n'
        'def test_prepared():\n'
        '    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as db:\n'
        '        assert db.execute("SELECT id FROM prepared").fetchone() == (7,)\n'
    )
    result = subprocess.run([sys.executable, "-m", "pytest", "-q"],
                            cwd=tmp_path, capture_output=True, text=True)
    assert result.returncode == 0, result.stdout + result.stderr


@shadow_pg
def test_shadow_intercepts_asyncpg_and_pool():
    async def run():
        conn = await asyncpg.connect("postgresql://bad@127.0.0.1:1/production")
        try:
            assert await conn.fetchval("SELECT count(*) FROM t") == 2
        finally:
            await conn.close()
        async with asyncpg.create_pool("postgresql://bad@127.0.0.1:1/production", min_size=1, max_size=1) as pool:
            assert await pool.fetchval("SELECT count(*) FROM t") == 2

    asyncio.run(run())


@shadow_pg
def test_shadow_intercepts_psycopg_async_connection():
    async def run():
        async with await psycopg.AsyncConnection.connect("postgresql://bad@127.0.0.1:1/production") as conn:
            assert (await (await conn.execute("SELECT count(*) FROM t")).fetchone()) == (2,)

    asyncio.run(run())


@shadow_pg
def test_shadow_intercepts_sqlalchemy_sync_engine():
    engine = create_engine("postgresql+psycopg://bad@127.0.0.1:1/production")
    try:
        with engine.connect() as conn:
            assert conn.execute(text("SELECT count(*) FROM t")).scalar_one() == 2
    finally:
        engine.dispose()


@shadow_pg
def test_shadow_intercepts_engine_constructed_before_fixture():
    with _cold_engine.connect() as conn:
        assert conn.execute(text("SELECT count(*) FROM t")).scalar_one() == 2
    _cold_engine.dispose()


def test_reused_sqlalchemy_pool_reconnects_to_next_fork(pgmem_snapshot):
    engine = create_engine("postgresql+psycopg://bad@127.0.0.1:1/production")
    try:
        with pgmem_snapshot.fork() as first:
            with pytest.MonkeyPatch.context() as patch:
                install_driver_routes(patch, first)
                remove = install_sqlalchemy_pool_guard(first)
                try:
                    with engine.connect() as conn:
                        assert conn.execute(text("SELECT count(*) FROM t")).scalar_one() == 2
                finally:
                    remove()
        with pgmem_snapshot.fork() as second:
            with pytest.MonkeyPatch.context() as patch:
                install_driver_routes(patch, second)
                remove = install_sqlalchemy_pool_guard(second)
                try:
                    with engine.connect() as conn:
                        assert conn.execute(text("SELECT count(*) FROM t")).scalar_one() == 2
                finally:
                    remove()
    finally:
        engine.dispose()


@shadow_pg
def test_shadow_intercepts_sqlalchemy_default_psycopg2_engine():
    engine = create_engine("postgresql://bad@127.0.0.1:1/production")
    try:
        with engine.connect() as conn:
            assert conn.execute(text("SELECT count(*) FROM t")).scalar_one() == 2
    finally:
        engine.dispose()


@shadow_pg
def test_shadow_intercepts_sqlalchemy_async_engine():
    async def run():
        engine = create_async_engine("postgresql+asyncpg://bad@127.0.0.1:1/production")
        try:
            async with engine.connect() as conn:
                assert (await conn.execute(text("SELECT count(*) FROM t"))).scalar_one() == 2
        finally:
            await engine.dispose()

    asyncio.run(run())


def test_live_baseline_captures_the_current_state(pgmem_snapshot):
    from pgmem.pytest_plugin import LiveBaseline

    with pgmem_snapshot.fork() as fork:
        baseline = LiveBaseline(fork)
        assert baseline.captured is None
        with psycopg.connect(fork.dsn) as db:  # what an application writes while starting
            db.execute("CREATE TABLE started (id int)")
            db.execute("INSERT INTO t VALUES (3)")
        captured = baseline.capture()
        assert baseline.capture() is captured
        with psycopg.connect(fork.dsn) as db:
            db.execute("INSERT INTO t VALUES (4)")
        assert count(fork) == 4
        fork.reset(snapshot=baseline.captured)
        assert count(fork) == 3
        with psycopg.connect(fork.dsn) as db:
            assert db.execute("SELECT count(*) FROM started").fetchone() == (0,)

        with psycopg.connect(fork.dsn) as db:
            db.execute("INSERT INTO t VALUES (5)")
        baseline.recapture()
        with psycopg.connect(fork.dsn) as db:
            db.execute("INSERT INTO t VALUES (6)")
        fork.reset(snapshot=baseline.captured)
        assert count(fork) == 4

        baseline.close()
        assert baseline.captured is None
        fork.reset()  # the snapshot the fork came from, without the application's schema
        assert count(fork) == 2
        with psycopg.connect(fork.dsn) as db:
            assert db.execute("SELECT count(*) FROM pg_tables WHERE tablename = 'started'").fetchone() == (0,)


def test_live_baseline_keeps_application_startup_schema(tmp_path):
    (tmp_path / "conftest.py").write_text(
        'import psycopg\n'
        'import pytest\n'
        '\n'
        '@pytest.fixture(scope="session")\n'
        'def pgmem_options():\n'
        '    return {"database": "app"}\n'
        '\n'
        '@pytest.fixture(scope="session")\n'
        'def pgmem_prepare():\n'
        '    def prepare(server):\n'
        '        with psycopg.connect(server.dsn) as db:\n'
        '            db.execute("CREATE TABLE prepared (id int)")\n'
        '            db.execute("INSERT INTO prepared VALUES (1)")\n'
        '    return prepare\n'
        '\n'
        '@pytest.fixture(scope="session", autouse=True)\n'
        'def app(pgmem_live_fork, pgmem_live_baseline):\n'
        '    with psycopg.connect(pgmem_live_fork.dsn) as db:\n'
        '        db.execute("CREATE TABLE app_migrated (id int)")\n'
        '        db.execute("INSERT INTO app_migrated VALUES (7)")\n'
        '    pgmem_live_baseline.capture()\n'
        '    return pgmem_live_fork\n'
    )
    (tmp_path / "test_live.py").write_text(
        'import psycopg\n'
        'from pgmem import shadow_pg\n'
        '\n'
        'def rows(dsn, table):\n'
        '    with psycopg.connect(dsn) as db:\n'
        '        return db.execute(f"SELECT count(*) FROM {table}").fetchone()[0]\n'
        '\n'
        '@shadow_pg(live=True)\n'
        'def test_startup_schema_is_the_baseline(pgmem_live_fork):\n'
        '    assert rows(pgmem_live_fork.dsn, "app_migrated") == 1\n'
        '    assert rows(pgmem_live_fork.dsn, "prepared") == 1\n'
        '    with psycopg.connect(pgmem_live_fork.dsn) as db:\n'
        '        db.execute("INSERT INTO app_migrated VALUES (8)")\n'
        '    assert rows(pgmem_live_fork.dsn, "app_migrated") == 2\n'
        '\n'
        '@shadow_pg(live=True)\n'
        'def test_next_test_starts_from_the_captured_baseline(pgmem_live_fork):\n'
        '    assert rows(pgmem_live_fork.dsn, "app_migrated") == 1\n'
    )
    result = subprocess.run([sys.executable, "-m", "pytest", "-q"],
                            cwd=tmp_path, capture_output=True, text=True)
    assert result.returncode == 0, result.stdout + result.stderr
