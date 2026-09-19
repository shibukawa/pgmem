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

The selection rule is simple: extensions that most managed services (Amazon RDS, Google Cloud SQL) and PGlite offer, ORM-facing ones first. The same steps build a custom pgmem with an extension the release does not ship.

### With an agent

The repository carries a skill, `add-pgmem-extension`, that walks a coding agent through the whole procedure. In a pgmem checkout, Claude Code finds it as `/add-pgmem-extension` on its own. For another agent, or a fork, install it from the repository:

```bash
npx skills add shibukawa/pgmem --skill add-pgmem-extension
```

Then ask in plain words, naming the contrib module:

> Add pg_surgery to pgmem.

The skill points the agent at a driver script, `skills/add-pgmem-extension/add-extension.sh`, and the agent does this:

1. `add-extension.sh status` lists what is bundled and which upstream contrib modules still have regression tests, downloading the pinned PostgreSQL source into `wasm/out/src` when it is not there.
2. `add-extension.sh add <name>` appends the name to `CONTRIB_MODULES` in `wasm/build.sh`, copies `contrib/<name>/sql`, `expected` and `data` to `testdata/regress/<name>`, and registers the Makefile's `REGRESS` list in `regress_test.go`. It reports psql commands the regression runner cannot emulate, upstream `REGRESS_OPTS` to mirror, and host libraries the module links.
3. The agent edits the extension lists by hand: `README.md`, `.knowledge/policy/bundled-extensions.md` and this page in both languages.
4. `add-extension.sh build` runs `wasm/build.sh` (Emscripten through Devbox: `postgres.wasm` and `internal/assets/share.tar.gz`) and `wasm/gen-aot.sh` (the wasm2go fork: `internal/aot/pgaot`). This takes a few minutes and is the only step that needs the wasm toolchain.
5. `add-extension.sh smoke <name>` starts pgmem, runs `CREATE EXTENSION` and lists the objects it installed; `add-extension.sh test <name>` replays the vendored regression files and diffs the transcript against upstream's expected output.
6. One commit per extension, `pgmem: bundle <name>`.

An extension outside the PostgreSQL tree follows pgvector: a lock file with repository, version and checksum, and its own compile step in `build.sh`. The skill describes that pattern but the driver does not automate it.

### By hand

The same steps without the driver:

1. Append the contrib name to `CONTRIB_MODULES` in `wasm/build.sh`. Its control, SQL and data files are copied into the share tree that ships with the module.
2. For an extension outside the PostgreSQL tree, add a lock file with repository, version and checksum, and a compile step in `build.sh`, as for pgvector.
3. Rebuild with `./wasm/build.sh` and regenerate the Go with `./wasm/gen-aot.sh`.
4. Vendor the extension's upstream regression files into `testdata/regress/<name>` and list them in `regress_test.go`, so the suite replays them against pgmem.
5. Commit one extension per commit.

Each module is compiled with its `_PG_init` and magic-block symbols renamed, and a generated table lets a static replacement for `dlopen` find them. That is what makes `CREATE EXTENSION` work without dynamic linking.
