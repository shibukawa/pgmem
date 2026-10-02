# SQLAlchemy unit tests

Create a SQLAlchemy engine in the `@shadow_pg` scope using a supported PostgreSQL driver: psycopg 3 (`postgresql+psycopg://`), psycopg2 (`postgresql://`), or asyncpg (`postgresql+asyncpg://`). Calls that open new connections route to the test fork while application connection arguments remain unchanged. Dispose the engine and close sessions before teardown; a pooled connection established earlier keeps its prior destination. See [unit lifecycle](python-unit.md) and driver tabs in `website/src/content/docs/guides/python/testing.mdx`.
