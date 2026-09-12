---
id: flow:add-extension
type: flow
title: Add an Extension to the Build
---
Steps to statically link a contrib or third-party extension into pgmem so CREATE EXTENSION works without dynamic loading (concept:static-modules).

```yaml
flow:
  - step: obtain source
    actor: maintainer
    action: contrib modules are already in the pinned tree; out-of-tree ones (pgvector) are fetched by wasm/build.sh with a pinned commit and sha256 next to postgres-pglite.lock, and patched by wasm/patches.py if they touch OpenSSL, ICU, zlib or threads
  - step: register the module
    actor: maintainer
    action: append name=dir to MODULE_DIRS in wasm/build.sh (contrib/<name> for contrib); the loop rebuilds it with the per-module Pg_magic_func and _PG_init renames and adds its objects to MODULE_OBJS
  - step: install share files
    actor: build
    action: emmake make install in the module dir so <name>.control and <name>--<ver>.sql land in share/extension and ship in internal/assets/share.tar.gz
  - step: regenerate
    actor: maintainer
    action: ./wasm/build.sh, ./wasm/gen-aot.sh, go run ./cmd/pgmem-mkdata, then go test ./...
  - step: verify
    actor: test
    action: a Go test runs CREATE EXTENSION <name> on a fresh server and one representative query; add it to pgmem_test.go
  - step: document
    actor: maintainer
    action: add the module to concept:static-modules linked_now with any caveat (no TLS, no ICU)
  constraints:
    - modules needing shared_preload_libraries or background workers run only if they work under postgres --single
    - modules calling OpenSSL need a host-side replacement like pgmem_cryptohash.inc or stay out
    - every module adds to the 102 MB generated Go and to compile time; keep the list to requirement:cloud-common-extensions
    - a static dlsym lookup cannot see symbols the module never exported; check gen_modules.py output counts when CREATE EXTENSION fails with undefined symbol
```
