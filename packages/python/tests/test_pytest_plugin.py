import pg8000.dbapi


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
