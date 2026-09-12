---
id: concept:static-modules
type: concept
title: Bundled Loadable Modules
---
Which PostgreSQL loadable modules are linked into the pgmem build and how dlopen is replaced so CREATE EXTENSION and LOAD resolve against a static table.

```yaml
summary:
  linked_now:
    plpgsql: installed by initdb; the only extension available at run time today
    dict_snowball: full-text stemming dictionaries for every snowball language
    encoding_conversions: all src/backend/utils/mb/conversion_procs modules (euc_jp, utf8_and_sjis, ...)
  not_compiled_in: [contrib modules, pgcrypto, OpenSSL and TLS, ICU collations, zlib, pg_stat_statements]
  mechanism:
    build: wasm/build.sh rebuilds each module directory with -DPg_magic_func=Pg_magic_func_<name> -D_PG_init=_PG_init_<name> so several modules can share one binary
    table: wasm/gen_modules.py runs llvm-nm over the module objects and emits pgmem_modules_gen.c, a name to address table per module
    resolver: wasm/pgmem_dl.c implements pgmem_dlopen, pgmem_dlsym, pgmem_dlclose, pgmem_dlerror; -Ddlopen=pgmem_dlopen in PG_CFLAGS routes dfmgr.c to it; the library basename without suffix selects the module
    placeholders: pgmem_module_name lets the host create empty .so placeholder files in the vfs so PostgreSQL's file existence checks pass
    share_files: control and sql files come from make install into internal/assets/share.tar.gz, unpacked into the vfs at Start
  error_when_missing: "dlopen: <name>: module is not linked into this pgmem build"
  add_one: flow:add-extension
  roadmap: requirement:cloud-common-extensions
```
