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
  minor_postgres_updates: a PostgreSQL 18.x point release only bumps Y; the exact 18.x is stated in the release notes and in the ready event
  semver_note: not semantic versioning; the Go module tag v1.18.Y stays a v1 module path (no /v2), and breaking API changes are announced in release notes instead of a major bump
  same_number_everywhere:
    - Go module tag v1.X.Y
    - cmd/pgmem binary version in the ready event (api:control-protocol)
    - PyPI pgmem version in pyproject.toml
    - Maven jp.shibu artifacts and Implementation-Version in build.gradle.kts
  compatibility: wrappers require an exactly matching binary version (policy:binary-distribution); the control protocol integer is versioned separately
  todo_2026_09_12: current sources still say 0.1.0 (pyproject.toml, build.gradle.kts, packages/java/README.md) and v0.1.0 in the ready-event examples (cmd/pgmem/main.go, docs/subprocess.md); switch them to 1.18.0 before the first publish
```
