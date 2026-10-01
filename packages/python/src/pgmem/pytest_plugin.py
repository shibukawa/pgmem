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
server parameters or logging. For a long-lived HTTP application,
``pgmem_live_fork`` keeps one endpoint for the session and
``pgmem_live_db`` restores it before each test.
"""

import pytest

import pgmem as _pgmem
from ._shadow import install_driver_routes, install_sqlalchemy_pool_guard, route_environment


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
def pgmem_prepare():
    """Override with a session fixture to migrate and seed the template."""


@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server, pgmem_prepare):
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


@pytest.fixture(scope="session")
def pgmem_shared_fork(pgmem_snapshot):
    """One prepared target shared by read-only ``fork=False`` tests."""
    with pgmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture(scope="session")
def pgmem_live_fork(pgmem_snapshot):
    """Stable endpoint for a long-lived application in one pytest worker."""
    with pgmem_snapshot.fork() as fork:
        yield fork


@pytest.fixture(scope="session")
def pgmem_live_snapshots(pgmem_snapshot):
    """Override with a mapping of dataset names to snapshots."""
    return {"default": pgmem_snapshot}


@pytest.fixture
def pgmem_live_db(request, pgmem_live_fork, pgmem_live_snapshots):
    """Restore the selected dataset before a sequential HTTP test."""
    marker = request.node.get_closest_marker("pgmem_dataset")
    if marker is None:
        name = "default"
    elif len(marker.args) == 1 and isinstance(marker.args[0], str) and not marker.kwargs:
        name = marker.args[0]
    else:
        raise pytest.UsageError("pgmem_dataset requires exactly one dataset name")
    try:
        snapshot = pgmem_live_snapshots[name]
    except KeyError as exc:
        raise pytest.UsageError(f"unknown pgmem dataset: {name}") from exc
    pgmem_live_fork.reset(snapshot=snapshot)
    return pgmem_live_fork


@pytest.fixture(autouse=True)
def _shadow_pg_route(request, monkeypatch):
    marker = request.node.get_closest_marker("shadow_pg")
    if marker is None:
        yield
        return
    if marker.args or set(marker.kwargs) - {"fork"}:
        raise pytest.UsageError("shadow_pg accepts only fork=True or fork=False")
    fork_enabled = marker.kwargs.get("fork", True)
    if not isinstance(fork_enabled, bool):
        raise pytest.UsageError("shadow_pg fork must be a bool")
    target = request.getfixturevalue("pgmem_fork" if fork_enabled else "pgmem_shared_fork")
    route_environment(monkeypatch, target)
    install_driver_routes(monkeypatch, target)
    remove_pool_guard = install_sqlalchemy_pool_guard(target)
    try:
        yield
    finally:
        remove_pool_guard()


def pytest_configure(config):
    config.addinivalue_line(
        "markers", "shadow_pg(fork=True): route this test to a prepared pgmem fork; fork=False shares a target for read-only tests"
    )
    config.addinivalue_line(
        "markers", "pgmem_dataset(name): select a snapshot for pgmem_live_db before an HTTP test"
    )
