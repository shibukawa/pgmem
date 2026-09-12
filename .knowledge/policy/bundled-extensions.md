---
id: policy:bundled-extensions
type: policy
title: Bundled Extensions
---
Which PostgreSQL extensions ship inside pgmem and how one is added.

```yaml
policy:
  bundled:
    - plpgsql  # from initdb
    - pgcrypto  # decision:pgcrypto-on-host
    - citext
    - pg_trgm
    - hstore
    - ltree
    - btree_gist
    - btree_gin
    - unaccent
  add_one:
    - append the contrib name to CONTRIB_MODULES in wasm/build.sh (control, SQL and DATA_TSEARCH files are copied into the share tree)
    - ./wasm/build.sh then ./wasm/gen-aot.sh (the symbol-named split keeps the pgaot diff to the new functions)
    - vendor contrib/<name>/sql and expected into testdata/regress/<name> and list the REGRESS order in contribRegress (regress_test.go)
    - one extension per commit
  selection: extensions most managed services (AWS RDS, Google Cloud SQL) and PGlite offer, ORM-facing ones first
  out_of_scope: background workers (pg_cron, pg_partman), outbound network (dblink, postgres_fdw), other language runtimes (plv8, plperl), PostGIS
```
