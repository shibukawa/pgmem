import os
import subprocess
import sys
from pathlib import Path

import pytest
from pgmem import shadow_pg

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
@shadow_pg(prepare=True)
def pgmem_prepare(pgmem_server):
    import psycopg

    with psycopg.connect("postgresql://bad@127.0.0.1:1/production") as conn:
        conn.execute("CREATE TABLE t (id int)")
        conn.execute("INSERT INTO t VALUES (1), (2)")


@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server, pgmem_prepare):
    return pgmem_server.snapshot(max_forks=3)
