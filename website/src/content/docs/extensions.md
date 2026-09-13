---
title: Extensions
description: The PostgreSQL extensions bundled with pgmem and how a new one is added.
---

Extensions are linked into the pgmem build statically. `CREATE EXTENSION` works for every name below without installing anything.

## Bundled

| Group | Extensions |
|---|---|
| Procedural language | `plpgsql` |
| Types and operators | `citext`, `hstore`, `ltree`, `cube`, `seg`, `isn`, `intarray`, `lo`, `earthdistance` |
| Text search and fuzzy matching | `pg_trgm`, `fuzzystrmatch`, `unaccent`, `dict_int`, `dict_xsyn` |
| Indexes | `btree_gist`, `btree_gin`, `bloom` |
| Vectors | `vector` (pgvector 0.8.6) |
| Cryptography and IDs | `pgcrypto`, `uuid-ossp` |
| Table functions and sampling | `tablefunc`, `tsm_system_rows`, `tsm_system_time` |
| Inspection and statistics | `pg_stat_statements`, `pgstattuple`, `pageinspect`, `pg_buffercache`, `pg_freespacemap`, `pg_visibility`, `amcheck`, `pg_prewarm` |
| Loadable module | `auto_explain` (`LOAD 'auto_explain'`, not `CREATE EXTENSION`) |

Snowball stemming dictionaries and every encoding conversion are linked in as well.

**pgcrypto without OpenSSL.** pgcrypto's OpenSSL and zlib layers are replaced by calls into Go's `crypto`, `math/big` and `compress` packages. Digests, AES, Blowfish, DES, 3DES and CAST5, OpenPGP encryption with RSA or ElGamal keys, and OpenPGP compression all work, and messages made by GnuPG decrypt. `fips_mode()` is always false, and bzip2-compressed OpenPGP packets are unsupported, as upstream.

## Not bundled

Extensions outside that list fail with `module is not linked into this pgmem build`. Some are out of scope by design:

- background workers, such as `pg_cron` and `pg_partman`'s worker, because single-user mode has no postmaster to start them;
- outbound network access, such as `dblink` and `postgres_fdw`;
- other language runtimes, such as `plv8`, `plperl` and `plpython`;
- PostGIS, for its size and its dependency chain.

Use a container for tests that need them.

## How an extension is added

The selection rule is simple: extensions that most managed services (Amazon RDS, Google Cloud SQL) and PGlite offer, ORM-facing ones first.

1. Append the contrib name to `CONTRIB_MODULES` in `wasm/build.sh`. Its control, SQL and data files are copied into the share tree that ships with the module.
2. For an extension outside the PostgreSQL tree, add a lock file with repository, version and checksum, and a compile step in `build.sh`, as for pgvector.
3. Rebuild with `./wasm/build.sh` and regenerate the Go with `./wasm/gen-aot.sh`.
4. Vendor the extension's upstream regression files into `testdata/regress/<name>` and list them in `regress_test.go`, so the suite replays them against pgmem.
5. Commit one extension per commit.

Each module is compiled with its `_PG_init` and magic-block symbols renamed, and a generated table lets a static replacement for `dlopen` find them. That is what makes `CREATE EXTENSION` work without dynamic linking.
