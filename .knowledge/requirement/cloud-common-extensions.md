---
id: requirement:cloud-common-extensions
type: requirement
title: Cloud-Common Extensions
---
Ship the extensions that managed PostgreSQL (AWS RDS, Google Cloud SQL, per pgextensions.org) and PGlite all offer, so a test against pgmem exercises the same SQL as production.

```yaml
summary:
  status: planned (stated 2026-09-12); nothing beyond plpgsql is linked yet (concept:static-modules)
  selection_rule: intersection of RDS, Cloud SQL and PGlite lists, contrib first because the source is already in the pinned tree
  candidates:
    tier1_contrib: [pg_trgm, uuid-ossp, hstore, citext, btree_gist, btree_gin, ltree, unaccent, tablefunc, fuzzystrmatch, intarray, cube, earthdistance, pg_stat_statements]
    tier1_external: [pgvector]
    needs_crypto_host_shim: [pgcrypto]  # OpenSSL is not linked; md5 and sha families already run on the host
    excluded: [postgis, timescaledb, plv8, plpython]  # size, background workers or a second runtime
  per_module_gate: flow:add-extension verify step must pass under postgres --single
  why: ORMs and app code use gen_random_uuid, similarity, vector distance; a test database that lacks them forces mocks
```
