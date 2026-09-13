---
id: policy:third-party-notices
type: policy
title: Third-Party Notices
---
Every distributed form of pgmem carries LICENSE (MIT, pgmem's own code) and NOTICE (the notices of upstream code compiled or translated into the module and binaries).

```yaml
policy:
  why: internal/aot/pgaot is PostgreSQL translated to Go, so the Go module and every binary are copies under the PostgreSQL License, whose notice must appear in all copies; the MIT and BSD notices below ask the same
  notice_sections: [PostgreSQL via postgres-pglite COPYRIGHT, pgvector LICENSE, Emscripten LICENSE, musl COPYRIGHT, wasm2go MIT, Go BSD-3 (same text for golang.org/x/crypto and x/sys), github.com/klauspost/compress LICENSE]
  carried_by:
    go_module: LICENSE and NOTICE at the module root
    archives: inside pgmem-<version>-<goos>-<goarch>/ (scripts/package-archives.sh)
    wheels: .dist-info/licenses through hatchling's default license-files globs; scripts/build-python-wheels.sh copies both into packages/python
    npm: LICENSE is always packed and NOTICE is in files; scripts/build-npm.sh copies both into each package
    jars: META-INF of every jar (packages/java/build.gradle.kts)
  source_files: NOTICE was assembled 2026-09-14 from wasm/out/src (postgres-pglite COPYRIGHT, pgvector tarball), toolchain/emsdk (Emscripten LICENSE, musl COPYRIGHT), ../wasm2go-fork/LICENSE and the GOMODCACHE licenses of the modules linked into cmd/pgmem
  revisit: when a bundled extension (policy:bundled-extensions) or a linked Go module brings another license
  manifests: package license fields stay MIT (pyproject.toml, package.json, POM); NOTICE covers the bundled parts
```
