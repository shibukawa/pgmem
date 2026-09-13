---
title: Versioning
description: pgmem release numbers are 1.X.Y, where X is the bundled PostgreSQL major version.
---

Every pgmem release is numbered **1.X.Y**.

| Part | Meaning |
|---|---|
| `1` | Fixed. It never changes, so the number reads like a PostgreSQL version at a glance. |
| `X` | The major version of the bundled PostgreSQL. Today that is `18`. A new PostgreSQL line starts at `1.<X>.0`. |
| `Y` | pgmem's own release counter within that PostgreSQL line. Any change to the Go module, the binary or a wrapper increments it. |

The first release under this scheme is **1.18.0**. The **0.1.x** releases before it bring the release pipeline up on every registry (GitHub Releases, the Go module proxy, PyPI, Maven Central and npm). They also carry one number everywhere, but that number says nothing about the bundled PostgreSQL.

- A PostgreSQL point release (18.3 to 18.4) only increments `Y`. The exact PostgreSQL version is in the release notes and in `SELECT version()`.
- This is not semantic versioning. The Go module stays on the `v1` import path, and breaking API changes are announced in the release notes.
- The Go module tag, the binary, the PyPI package, the Maven artifacts and the npm packages always carry the same number. Wrappers refuse a binary whose version differs from their own.
