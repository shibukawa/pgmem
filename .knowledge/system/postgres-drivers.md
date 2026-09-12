---
id: system:postgres-drivers
type: system
title: PostgreSQL Client Drivers
---
Driver capabilities in each target language that constrain decision:data-plane-transport and the wrapper return types.

```yaml
system:
  python:
    psycopg3: DSN URL; uds via host=/dir; pure-python or libpq backend
    asyncpg: DSN URL; uds via host=/dir
    sqlalchemy: takes the URL, postgresql+psycopg://
  java:
    pgjdbc: 'jdbc:postgresql://host:port/db?user=..'; no built-in uds (junixsocket socketFactory needed); DriverManager or PGSimpleDataSource
    hikaricp: one pool per fork is fine (rule:single-session-per-backend serializes at transaction boundaries)
  auth: trust, no password; URLs need sslmode=disable because the build has no TLS
  implication: loopback TCP is the only transport every listed driver supports without extra dependencies
```
