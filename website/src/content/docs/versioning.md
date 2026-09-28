---
title: Versioning
description: Stable pgmem versions use 1.X.Y; beta versions use 0.X.Y. The table maps pgmem releases to bundled PostgreSQL versions.
---

Stable pgmem versions use `1.X.Y`; beta versions use `0.X.Y`. `X` identifies the bundled PostgreSQL major version. `Y` counts pgmem releases for that major and track. Every release artifact uses the same full `X.Y.Z` version.

| pgmem version | Bundled PostgreSQL | Track | Meaning |
|---|---|---|---|
| `1.18.x` | PostgreSQL `18.x` (`1.18.0` uses `18.3`) | Stable | Current stable line; `Y` advances for each pgmem release. |
| `0.19.x` | PostgreSQL 19 prereleases (initially `Beta 4`) | Beta | The initial build, `0.19.0`, uses `Beta 4`. |
| `1.19.x` | PostgreSQL `19.x` after general availability | Stable | Stable PostgreSQL 19 line, starting at `1.19.0`. |
