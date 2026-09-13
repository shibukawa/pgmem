---
id: policy:versioning
type: policy
title: Versioning
---
Every pgmem release is numbered 1.X.Y: 1 is fixed, X is the major version of the bundled PostgreSQL, Y counts pgmem's own releases for that PostgreSQL line.

```yaml
policy:
  format: 1.X.Y
  parts:
    "1": fixed; never bumped, so the number reads as a PostgreSQL version at a glance
    X: PostgreSQL major version in the build (18 today, pinned by wasm/postgres-pglite.lock); a new PostgreSQL line starts at 1.<X>.0
    Y: pgmem release counter within one PostgreSQL line, incremented for every change of the Go module, binary or wrappers; no separate patch level
  first_release: 1.18.0
  bootstrap: 0.1.x releases (decided 2026-09-14) bring flow:release up on every registry before 1.18.0; one number everywhere, but it does not encode PostgreSQL; manifests and ready-event examples say 0.1.0 until then
  minor_postgres_updates: a PostgreSQL 18.x point release only bumps Y; the exact 18.x is stated in the release notes and in the ready event
  semver_note: not semantic versioning; the Go module tag v1.18.Y stays a v1 module path (no /v2), and breaking API changes are announced in release notes instead of a major bump
  same_number_everywhere:  # the tag is the source; scripts/set-version.sh stamps the manifests during flow:release
    - Go module tag vX.Y.Z
    - cmd/pgmem binary version in the ready event (api:control-protocol), -X main.version=vX.Y.Z from scripts/build-binaries.sh
    - PyPI pgmem version in packages/python/pyproject.toml
    - Maven jp.shibu artifacts and Implementation-Version, from packages/java/gradle.properties
    - npm @pgmem/core, @pgmem/<platform> and core's optionalDependencies pins
  compatibility: wrappers require an exactly matching binary version (policy:binary-distribution); the control protocol integer is versioned separately
```
