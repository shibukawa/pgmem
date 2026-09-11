---
id: decision:clone-return-type
type: decision
title: Clone Return Type
---
Which handle a fork exposes: DSN, *sql.DB, pgx.Conn, or pgxpool.Pool.

```yaml
decision:
  resolution: pending
  options:
    dsn_only: caller opens its own driver; must remember rule:single-session-per-backend
    sql_db: "*sql.DB; matches most application code; pool size is free since rule:single-session-per-backend serializes at transaction boundaries"
    pgx_conn: "*pgx.Conn; one connection maps exactly to one backend session and is fastest"
    pgx_pool: "*pgxpool.Pool; for code that takes a pool"
  recommendation:
    - Fork always exposes DSN()
    - convenience constructors DB(), PgxConn(), PgxPool(); no connection cap needed since transaction-scoped locking landed
    - handles are closed by Fork.Close so callers close nothing else
```
