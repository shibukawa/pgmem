from sqlalchemy import Column, Integer, MetaData, Table, insert

metadata = MetaData()
orders = Table("orders", metadata, Column("id", Integer, primary_key=True))


class OrderApplication:
    def __init__(self, engine):
        self.engine = engine

    def create_order(self, order_id):
        with self.engine.begin() as connection:
            connection.execute(insert(orders).values(id=order_id))
