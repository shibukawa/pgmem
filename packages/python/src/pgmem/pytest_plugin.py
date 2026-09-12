"""pytest fixtures for pgmem.

Session scope: ``pgmem_process`` (the process), ``pgmem_server`` (its default
template) and ``pgmem_snapshot`` (a snapshot of the template). Override
``pgmem_snapshot`` in your conftest to run migrations first::

    @pytest.fixture(scope="session")
    def pgmem_snapshot(pgmem_server):
        run_migrations(pgmem_server.dsn)
        return pgmem_server.snapshot()

Per test: ``pgmem_dsn`` gives the DSN of a fresh fork that is closed when
the test ends. ``pgmem_class_dsn`` shares one fork across a test class,
for read-only tests. ``pgmem_options`` (session) returns the keyword
arguments for :func:`pgmem.start`; override it to change database name,
server parameters or logging.
"""

import pytest

import pgmem as _pgmem


@pytest.fixture(scope="session")
def pgmem_options():
    return {}


@pytest.fixture(scope="session")
def pgmem_process(pgmem_options):
    with _pgmem.start(**pgmem_options) as pg:
        yield pg


@pytest.fixture(scope="session")
def pgmem_server(pgmem_process):
    return pgmem_process.template


@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server):
    return pgmem_server.snapshot()


@pytest.fixture
def pgmem_fork(pgmem_snapshot):
    with pgmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture
def pgmem_dsn(pgmem_fork):
    return pgmem_fork.dsn


@pytest.fixture(scope="class")
def pgmem_class_fork(pgmem_snapshot):
    with pgmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture(scope="class")
def pgmem_class_dsn(pgmem_class_fork):
    return pgmem_class_fork.dsn
