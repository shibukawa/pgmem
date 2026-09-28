---
title: Versioning
description: Stable pgmem releases use 1.X.Y and beta releases use 0.X.Y, with X matching the bundled PostgreSQL major version.
---

Stable pgmem releases use **1.X.Y**. Beta releases use **0.X.Y**; the leading `0` marks a pre-release.

| Part | Stable track | Beta track |
|---|---|---|
| Leading number | `1` marks stable releases. | `0` marks beta releases. |
| `X` | Major version of the bundled PostgreSQL; a new stable line starts at `1.<X>.0`. | Major version of the bundled PostgreSQL; a beta line starts at `0.<X>.0`. |
| `Y` | pgmem's release counter within the PostgreSQL major line. | pgmem's release counter within the PostgreSQL beta line. |

The current PostgreSQL 18 stable line starts at **1.18.0** and currently bundles PostgreSQL 18.3. The PostgreSQL 19 Beta 4 build starts the beta line at **0.19.0** (`0.19` is shorthand for that line). When PostgreSQL 19 is generally available, its stable pgmem line starts at **1.19.0**.

- A PostgreSQL point release within one major line increments pgmem's `Y`. The exact PostgreSQL version is in the release notes and in `SELECT version()`.
- This is not semantic versioning. The Go module stays on the `v1` import path, and breaking API changes are announced in the release notes.
- The Go module tag, the binary, the PyPI package, the Maven artifacts and the npm packages always carry the same number. Wrappers refuse a binary whose version differs from their own.
