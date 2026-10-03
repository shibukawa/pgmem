import pytest
from sqlalchemy import create_engine
from app import metadata, OrderApplication


def engine_for(dsn):
    return create_engine(dsn.replace("postgres://", "postgresql+psycopg://", 1))


@pytest.fixture(scope="session")
def pgmem_prepare():
    def prepare(server):
        engine = engine_for(server.dsn)
        try:
            metadata.create_all(engine)
        finally:
            engine.dispose()
    return prepare


@pytest.fixture
def engine(pgmem_test_dsn):
    value = engine_for(pgmem_test_dsn)
    try:
        yield value
    finally:
        value.dispose()


@pytest.fixture
def app(engine):
    return OrderApplication(engine)
