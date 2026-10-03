import pytest
from sqlalchemy import select, func
from app import orders


def test_application_write_is_visible_to_assertions(app, engine):
    app.create_order(42)
    with engine.connect() as connection:
        assert connection.execute(select(orders.c.id)).scalar_one() == 42


@pytest.mark.pgmem(fork=False)
def test_read_the_prepared_baseline(app, engine):
    with engine.connect() as connection:
        assert connection.execute(select(func.count()).select_from(orders)).scalar_one() == 0
