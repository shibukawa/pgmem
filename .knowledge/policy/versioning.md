---
id: policy:versioning
type: policy
title: Versioning
---
Stable pgmem releases use 1.X.Y; beta releases use 0.X.Y. X is the major version of the bundled PostgreSQL, and Y counts pgmem's releases within that PostgreSQL line.

```yaml
policy:
  format:
    stable: 1.X.Y
    beta: 0.X.Y
  parts:
    "1": stable track; fixed for stable releases, so the number reads as a PostgreSQL version at a glance
    "0": beta track; every 0.x release is pre-release, not stable
    X: PostgreSQL major version in the build; each stable line starts at 1.<X>.0 and each beta line at 0.<X>.0
    Y: pgmem release counter within that PostgreSQL major and track, incremented for every change of the Go module, binary or wrappers; no separate patch level
  stable_baseline: 1.18.0 for PostgreSQL 18; use this as the current stable version
  postgres_19_beta: PostgreSQL 19 Beta 4 is the requested beta input; its initial pgmem version is 0.19.0 (0.19 is shorthand for the beta line; tags and manifests keep X.Y.Z)
  postgres_19_stable: when PostgreSQL 19 is stable, start its pgmem stable line at 1.19.0
  minor_postgres_updates: a PostgreSQL 18.x point release only bumps Y; the exact 18.x is stated in the release notes and in the ready event
  semver_note: not semantic versioning; the Go module tag v1.18.Y stays a v1 module path (no /v2), and breaking API changes are announced in release notes instead of a major bump
  same_number_everywhere:  # the tag is the source; scripts/set-version.sh stamps the manifests during flow:release
    - Go module tag vX.Y.Z
    - cmd/pgmem binary version in the ready event (api:control-protocol), -X main.version=vX.Y.Z from scripts/build-binaries.sh
    - PyPI pgmem version in packages/python/pyproject.toml
    - Maven io.github.shibukawa.pgmem artifacts and Implementation-Version, from packages/java/gradle.properties
    - npm @pgmem/core, @pgmem/<platform> and core's optionalDependencies pins
  compatibility: wrappers require an exactly matching binary version (policy:binary-distribution); the control protocol integer is versioned separately
```
