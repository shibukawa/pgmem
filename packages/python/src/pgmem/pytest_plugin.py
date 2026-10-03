"""pytest fixtures for pgmem.

Session scope: ``pgmem_process`` (the process), ``pgmem_server`` (its default
template) and ``pgmem_snapshot`` (a snapshot of the template). Override
``pgmem_prepare`` in your conftest to run migrations before the automatic snapshot::

    @pytest.fixture(scope="session")
    def pgmem_prepare():
        return lambda server: run_migrations(server.dsn)

Per test: ``pgmem_dsn`` gives the DSN of a fresh fork that is closed when
the test ends. ``pgmem_class_dsn`` shares one fork across a test class,
for read-only tests. ``pgmem_options`` (session) returns the keyword
arguments for :func:`pgmem.start`; override it to change database name,
server parameters or logging.

``pgmem_test_dsn`` selects fork, shared or reset isolation using the pgmem
marker. ``pgmem_shared_dsn`` keeps a stable endpoint for persistent clients
in serial reset-mode tests. ``pgmem_fixture_options`` sets isolation and deadlines.
"""

import pytest

import pgmem as _pgmem
from ._shadow import install_driver_routes, install_sqlalchemy_pool_guard


def pytest_configure(config):
    config.addinivalue_line("markers", "shadow_pg(fork=True, live=False): route a test through the normal driver; live=True uses a stable application endpoint")
    config.addinivalue_line("markers", "pgmem_dataset(name): select a baseline for live application tests")
    config.addinivalue_line(
        "markers", "pgmem(fork=True, isolation='fork'): select fork, shared or reset isolation; use either fork or isolation"
    )


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
    """Override with a callback that migrates and seeds a Server once."""
    return lambda server: None


@pytest.fixture(scope="session")
def pgmem_fixture_options():
    """Isolation and deadlines, independent of process startup options."""
    return {"isolation": "fork", "fork_timeout": 30.0, "snapshot_timeout": 30.0, "reset_timeout": 5.0, "shared_read_only": False}


@pytest.fixture(scope="session")
def pgmem_snapshot(pgmem_server, pgmem_prepare, pgmem_fixture_options):
    # A preparation fixture may perform setup itself or return a Server callback.
    if pgmem_prepare is not None:
        if not callable(pgmem_prepare):
            raise pytest.UsageError("pgmem_prepare must perform setup or return a callable")
        pgmem_prepare(pgmem_server)
    return pgmem_server.snapshot(
        max_forks=pgmem_fixture_options.get("max_forks"),
        timeout=pgmem_fixture_options.get("snapshot_timeout", 30.0),
    )


@pytest.fixture
def pgmem_fork(pgmem_snapshot, pgmem_fixture_options):
    with pgmem_snapshot.fork(timeout=pgmem_fixture_options.get("fork_timeout", 30.0)) as fork:
        yield fork


@pytest.fixture
def pgmem_dsn(pgmem_fork):
    return pgmem_fork.dsn


@pytest.fixture(scope="class")
def pgmem_class_fork(pgmem_snapshot, pgmem_fixture_options):
    with pgmem_snapshot.fork(timeout=pgmem_fixture_options.get("fork_timeout", 30.0)) as fork:
        yield fork


@pytest.fixture(scope="class")
def pgmem_class_dsn(pgmem_class_fork, pgmem_fixture_options):
    dsn = pgmem_class_fork.dsn
    return _pgmem.readonly_dsn(dsn) if pgmem_fixture_options.get("shared_read_only", False) else dsn


def _isolation(request):
    mode = request.getfixturevalue("pgmem_fixture_options").get("isolation", "fork")
    marker = request.node.get_closest_marker("pgmem")
    if marker is not None:
        if marker.args or len(marker.kwargs) != 1:
            raise pytest.UsageError("@pytest.mark.pgmem requires fork=<bool> or isolation='fork'|'shared'|'reset'")
        if "fork" in marker.kwargs and isinstance(marker.kwargs["fork"], bool):
            mode = "fork" if marker.kwargs["fork"] else "shared"
        elif "isolation" in marker.kwargs:
            mode = marker.kwargs["isolation"]
        else:
            raise pytest.UsageError("@pytest.mark.pgmem requires fork=<bool> or isolation='fork'|'shared'|'reset'")
    if mode not in ("fork", "shared", "reset"):
        raise pytest.UsageError("pgmem isolation must be fork, shared or reset")
    return mode


@pytest.fixture(scope="session")
def pgmem_shared_fork(pgmem_snapshot, pgmem_fixture_options):
    with pgmem_snapshot.fork(timeout=pgmem_fixture_options.get("fork_timeout", 30.0)) as fork:
        yield fork


@pytest.fixture(scope="session")
def pgmem_shared_dsn(pgmem_shared_fork):
    """Stable endpoint for session-scoped clients in serial reset-mode tests."""
    return pgmem_shared_fork.dsn


@pytest.fixture(autouse=True)
def _pgmem_reset(request):
    # Do not start pgmem for tests that do not use it.
    if request.node.get_closest_marker("pgmem") is None and not any(
        name in request.fixturenames for name in ("pgmem_test_dsn", "pgmem_shared_dsn", "pgmem_shared_fork")
    ):
        yield
        return
    if _isolation(request) != "reset":
        yield
        return
    fork = request.getfixturevalue("pgmem_shared_fork")
    try:
        yield
    finally:
        fork.reset(timeout=request.getfixturevalue("pgmem_fixture_options").get("reset_timeout", 5.0))


@pytest.fixture
def pgmem_test_dsn(request):
    """Selected DSN: isolated by default; @pytest.mark.pgmem(fork=False) shares the template.

    Application-owned client fixtures should depend on this fixture, so selection
    happens before they create a connection or pool. Explicit pgmem_dsn remains
    available for tests that always require a fork.
    """
    mode = _isolation(request)
    if mode == "shared":
        request.getfixturevalue("pgmem_snapshot")
        dsn = request.getfixturevalue("pgmem_server").dsn
        options = request.getfixturevalue("pgmem_fixture_options")
        return _pgmem.readonly_dsn(dsn) if options.get("shared_read_only", False) else dsn
    if mode == "reset":
        return request.getfixturevalue("pgmem_shared_dsn")
    return request.getfixturevalue("pgmem_dsn")


@pytest.fixture(scope="session")
def pgmem_live_fork(pgmem_snapshot, pgmem_fixture_options):
    """Stable endpoint for a long-lived application in one pytest worker."""
    with pgmem_snapshot.fork(timeout=pgmem_fixture_options.get("fork_timeout", 30.0)) as fork:
        yield fork


@pytest.fixture(scope="session")
def pgmem_live_snapshots(pgmem_snapshot):
    """Override with a mapping of dataset names to snapshots."""
    return {"default": pgmem_snapshot}


@pytest.fixture
def pgmem_live_db(request, pgmem_live_fork, pgmem_live_snapshots, pgmem_fixture_options):
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
    pgmem_live_fork.reset(snapshot=snapshot, timeout=pgmem_fixture_options.get("reset_timeout", 5.0))
    return pgmem_live_fork


@pytest.fixture(autouse=True)
def _shadow_pg_route(request, monkeypatch):
    marker = request.node.get_closest_marker("shadow_pg")
    if marker is None:
        yield
        return
    if marker.args or set(marker.kwargs) - {"fork", "live"}:
        raise pytest.UsageError("shadow_pg accepts only fork and live options")
    fork_enabled = marker.kwargs.get("fork", True)
    live = marker.kwargs.get("live", False)
    if not isinstance(fork_enabled, bool) or not isinstance(live, bool) or (live and not fork_enabled):
        raise pytest.UsageError("shadow_pg fork and live must be bool; live requires fork=True")
    target = request.getfixturevalue("pgmem_live_db" if live else "pgmem_fork" if fork_enabled else "pgmem_shared_fork")
    install_driver_routes(monkeypatch, target)
    remove_pool_guard = install_sqlalchemy_pool_guard(target)
    try:
        yield
    finally:
        remove_pool_guard()
