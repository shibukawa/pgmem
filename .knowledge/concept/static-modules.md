---
id: concept:static-modules
type: concept
title: Static Module Loading
---
How loadable modules and extensions are linked into the one pgmem binary and how dlopen is replaced so CREATE EXTENSION and LOAD resolve against a static table; the list itself is policy:bundled-extensions.

```yaml
summary:
  module_kinds:
    core: plpgsql, dict_snowball, every encoding conversion module (src/backend/utils/mb/conversion_procs)
    contrib: CONTRIB_MODULES in wasm/build.sh (policy:bundled-extensions)
    external: pgvector, pinned by wasm/pgvector.lock and compiled against the installed server headers like PGXS
  not_compiled_in: [OpenSSL and TLS, ICU collations, zlib]  # pgcrypto gets crypto and compression from Go instead (decision:pgcrypto-on-host)
  mechanism:
    build: each module is compiled with -DPg_magic_func=Pg_magic_func_<name> -D_PG_init=_PG_init_<name> so many modules share one binary
    table: wasm/gen_modules.py runs llvm-nm over the module objects and emits pgmem_modules_gen.c, a name to address table per module
    resolver: wasm/pgmem_dl.c implements pgmem_dlopen, pgmem_dlsym, pgmem_dlclose, pgmem_dlerror; -Ddlopen=pgmem_dlopen in PG_CFLAGS routes dfmgr.c to it; the library basename without suffix selects the module
    placeholders: pgmem_module_name lets the host create empty .so placeholder files in the vfs so PostgreSQL's file existence checks pass
    share_files: control, SQL and tsearch data files are copied into the share tree and ship in internal/assets/share.tar.gz, unpacked into the vfs at Start
  error_when_missing: "dlopen: <name>: module is not linked into this pgmem build"
  cost_per_module: more generated Go (concept:architecture) and compile time; symbol-named output keeps the diff to the new functions
  verification: upstream contrib regression files replayed from testdata/regress (regress_test.go)
```
