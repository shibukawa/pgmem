import threading
import time

import pg8000.dbapi
import psycopg
import pytest

import pgmem


def connect(server):
    return pg8000.dbapi.connect(user=server.user, host=server.host, port=server.port, database=server.database)


def count(server):
    conn = connect(server)
    cur = conn.cursor()
    cur.execute("SELECT count(*) FROM t")
    (n,) = cur.fetchone()
    conn.close()
    return n


def test_start_and_query():
    with pgmem.start(database="app") as pg:
        assert pg.template.database == "app"
        assert pg.template.dsn.startswith("postgres://postgres@127.0.0.1:")
        assert pg.pid > 0
        conn = connect(pg.template)
        cur = conn.cursor()
        cur.execute("SELECT 1 + 1")
        assert tuple(cur.fetchone()) == (2,)
        conn.close()
    assert pg._proc.returncode == 0


def test_psycopg_dsn():
    with pgmem.start() as pg:
        with psycopg.connect(pg.template.dsn) as conn:
            assert conn.execute("SELECT current_database()").fetchone() == ("postgres",)


def test_snapshot_fork_isolation():
    with pgmem.start(database="app") as pg:
        conn = connect(pg.template)
        cur = conn.cursor()
        cur.execute("CREATE TABLE t (id int)")
        cur.execute("INSERT INTO t VALUES (1), (2)")
        conn.commit()
        conn.close()
        snap = pg.template.snapshot(max_forks=2)
        with snap.fork() as f1, snap.fork() as f2:
            assert f1.id != f2.id and f1.port != f2.port
            conn = connect(f1)
            conn.cursor().execute("INSERT INTO t VALUES (3)")
            conn.commit()
            conn.close()
            assert count(f1) == 3
            assert count(f2) == 2
            assert count(pg.template) == 2

            with pytest.raises(pgmem.ProtocolError) as ei:
                snap.fork(timeout=0.2)
            assert ei.value.code == "pool_timeout"

            # a blocked fork is released by a close from another thread
            got = {}

            def blocked():
                with snap.fork() as f3:
                    got["n"] = count(f3)

            th = threading.Thread(target=blocked)
            th.start()
            time.sleep(0.2)
            assert th.is_alive()
            f1.close()
            th.join(30)
            assert got == {"n": 2}
        f1.close()  # idempotent
        snap.close()
        with pytest.raises(pgmem.ProtocolError) as ei:
            snap.fork()
        assert ei.value.code == "unknown_id"


def test_start_server_extra_template():
    with pgmem.start(database="app") as pg:
        audit = pg.start_server(database="audit", params=["log_statement=all"])
        assert audit.database == "audit" and audit.id != pg.template.id
        conn = connect(audit)
        conn.cursor().execute("CREATE TABLE t (id int)")
        conn.commit()
        conn.close()
        snap = audit.snapshot()
        with snap.fork() as f:
            assert f.database == "audit"
            assert count(f) == 0


def test_server_exit_fails_pending():
    pg = pgmem.start()
    pg._proc.kill()
    pg._proc.wait()
    with pytest.raises(pgmem.ServerExited):
        pg.template.snapshot()
    pg.close()


def test_missing_binary(monkeypatch, tmp_path):
    monkeypatch.setenv("PGMEM_BINARY", str(tmp_path / "nope"))
    monkeypatch.setenv("PATH", str(tmp_path))
    monkeypatch.setattr(pgmem._binary, "__file__", str(tmp_path / "pkg" / "_binary.py"))
    with pytest.raises(FileNotFoundError):
        pgmem.find_binary()


def test_snapshot_busy_while_transaction_open():
    with pgmem.start() as pg:
        conn = connect(pg.template)
        conn.cursor().execute("SELECT 1")  # pg8000 opened a transaction and keeps it open
        with pytest.raises(pgmem.ProtocolError) as ei:
            pg.template.snapshot(timeout=0.2)
        assert ei.value.code == "busy"
        conn.commit()
        conn.close()
        pg.template.snapshot(timeout=5).close()
