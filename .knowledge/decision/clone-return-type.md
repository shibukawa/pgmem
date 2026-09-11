---
id: decision:clone-return-type
type: decision
title: Clone Return Type
---
Which handle a fork exposes: DSN, *sql.DB, pgx.Conn, or pgxpool.Pool.

```yaml
decision:
  resolution: all of them, as pgmemtest.Fixture methods over one Fork (2026-09-12)
  shapes:
    - Fork(t) *pgmem.Server exposes DSN() and Dial for anything else
    - DB(t) *sql.DB via pgx stdlib with the in-process dialer
    - PgxConn(t) *pgx.Conn
    - PgxPool(t) *pgxpool.Pool with default size; rule:single-session-per-backend serializes at transaction boundaries
    - DSN(t) string
  note: no connection cap needed since transaction-scoped locking landed
```
