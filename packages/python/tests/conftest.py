import os
import subprocess
import sys
from pathlib import Path

import pytest

REPO = Path(__file__).resolve().parents[3]
BUILD_DIR = Path(__file__).resolve().parents[1] / "src" / "pgmem" / "_bin"


def pytest_configure(config):
    """Build cmd/pgmem into the package's _bin directory unless PGMEM_BINARY is set."""
    if os.environ.get("PGMEM_BINARY"):
        return
    name = "pgmem.exe" if sys.platform == "win32" else "pgmem"
    out = BUILD_DIR / name
    subprocess.run(["go", "build", "-o", str(out), "./cmd/pgmem"], cwd=REPO, check=True)
    os.environ["PGMEM_BINARY"] = str(out)


@pytest.fixture(scope="session")
def pgmem_options():
    return {"database": "app"}


@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server):
    import pg8000.dbapi

    conn = pg8000.dbapi.connect(user=pgmem_server.user, host=pgmem_server.host,
                                port=pgmem_server.port, database=pgmem_server.database)
    cur = conn.cursor()
    cur.execute("CREATE TABLE t (id int)")
    cur.execute("INSERT INTO t VALUES (1), (2)")
    conn.commit()
    conn.close()
    return pgmem_server.snapshot(max_forks=2)
