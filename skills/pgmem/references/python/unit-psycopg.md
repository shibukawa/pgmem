# psycopg unit tests

With `@shadow_pg`, ordinary `psycopg.connect(...)` and `psycopg2.connect(...)` calls made during the marked test reach its fork, regardless of the application's unchanged connection string. Use context managers or `finally` to commit/close connections before the marker exits. SQLAlchemy's default `postgresql://` dialect uses psycopg2 and is also covered. Use `fork=False` only for read-only tests. See [unit lifecycle](unit.md).
