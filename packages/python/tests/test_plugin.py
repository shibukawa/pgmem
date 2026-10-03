pytest_plugins = ["pytester"]


def configure(pytester):
    pytester.makeconftest('''
import pytest
import psycopg

@pytest.fixture(scope="session")
def pgmem_prepare():
    def prepare(server):
        with psycopg.connect(server.dsn) as conn:
            conn.execute("CREATE TABLE items(v int); INSERT INTO items VALUES (1),(2)")
    return prepare

@pytest.fixture(scope="session")
def db(pgmem_shared_dsn):
    with psycopg.connect(pgmem_shared_dsn, autocommit=True) as conn:
        yield conn
''')


def test_prepare_and_stable_client_reset(pytester):
    configure(pytester)
    pytester.makepyfile('''
import pytest
import psycopg

@pytest.mark.pgmem(isolation="reset")
class TestStable:
    def test_write(self, db):
        db.execute("INSERT INTO items VALUES (3)")
        assert db.execute("SELECT count(*) FROM items").fetchone()[0] == 3
    def test_restored(self, db):
        assert db.execute("SELECT count(*) FROM items").fetchone()[0] == 2

@pytest.mark.pgmem(fork=False)
def test_shared(pgmem_test_dsn, pgmem_server):
    assert pgmem_test_dsn == pgmem_server.dsn
    with psycopg.connect(pgmem_test_dsn) as conn:
        assert conn.execute("SELECT count(*) FROM items").fetchone()[0] == 2
''')
    pytester.runpytest_subprocess("-q").assert_outcomes(passed=3)


def test_reset_runs_after_test_and_fixture_failures(pytester):
    configure(pytester)
    pytester.makepyfile('''
import pytest

pytestmark = pytest.mark.pgmem(isolation="reset")

def test_write_failure(db):
    db.execute("INSERT INTO items VALUES (3)")
    assert False, "expected assertion failure"

@pytest.fixture
def broken(db):
    db.execute("INSERT INTO items VALUES (4)")
    raise RuntimeError("expected setup failure")

def test_setup_failure(broken):
    pass

def test_clean_baseline(db):
    assert db.execute("SELECT count(*) FROM items").fetchone()[0] == 2
''')
    pytester.runpytest_subprocess("-q").assert_outcomes(passed=1, failed=1, errors=1)


def test_shared_read_only_guard_keeps_isolated_and_reset_writes(pytester):
    configure(pytester)
    pytester.makeconftest(pytester.path.joinpath("conftest.py").read_text() + '''
@pytest.fixture(scope="session")
def pgmem_fixture_options():
    return {"shared_read_only": True}
''')
    pytester.makepyfile('''
import pytest
import psycopg

@pytest.mark.pgmem(fork=False)
def test_guard(pgmem_test_dsn):
    with psycopg.connect(pgmem_test_dsn, autocommit=True) as conn:
        assert conn.execute("SELECT count(*) FROM items").fetchone()[0] == 2
        with pytest.raises(psycopg.errors.ReadOnlySqlTransaction):
            conn.execute("INSERT INTO items VALUES (3)")

class TestClassGuard:
    def test_guard(self, pgmem_class_dsn):
        with psycopg.connect(pgmem_class_dsn, autocommit=True) as conn:
            with pytest.raises(psycopg.errors.ReadOnlySqlTransaction):
                conn.execute("INSERT INTO items VALUES (3)")

def test_isolated(pgmem_test_dsn):
    with psycopg.connect(pgmem_test_dsn) as conn:
        conn.execute("INSERT INTO items VALUES (3)")

@pytest.mark.pgmem(isolation="reset")
def test_reset_mode_can_write(db):
    db.execute("INSERT INTO items VALUES (3)")
''')
    pytester.runpytest_subprocess("-q").assert_outcomes(passed=4)
