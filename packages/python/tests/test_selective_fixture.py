import pytest


def test_default_uses_fork(pgmem_test_dsn, pgmem_server):
    assert pgmem_test_dsn != pgmem_server.dsn


@pytest.mark.pgmem(fork=False)
def test_marker_shares_prepared_template(pgmem_test_dsn, pgmem_server):
    assert pgmem_test_dsn == pgmem_server.dsn


@pytest.mark.pgmem(fork=True)
def test_marker_requests_fork(pgmem_test_dsn, pgmem_server):
    assert pgmem_test_dsn != pgmem_server.dsn
