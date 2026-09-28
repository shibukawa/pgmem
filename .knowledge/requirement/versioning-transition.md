---
id: requirement:versioning-transition
type: requirement
title: Versioning Transition
---
Move the current PostgreSQL 18 stable line to 1.18.0 and introduce the PostgreSQL 19 Beta 4 build as the 0.19 beta line (policy:versioning, flow:release).

```yaml
requirement:
  outcomes:
    - current stable release and all pgmem-owned artifacts use 1.18.0
    - beta versions use 0.X.Y, with X matching the bundled PostgreSQL major
    - PostgreSQL 19 Beta 4 is the base for pgmem 0.19.0; 0.19 is shorthand for this beta line
    - PostgreSQL 19 stable releases start at 1.19.0
    - release tags, binaries, wrappers, package metadata and version examples carry the same full X.Y.Z version
  acceptance:
    - release/build inputs select PostgreSQL 19 Beta 4 for the 0.19 beta build
    - the beta build identifies the exact upstream PostgreSQL version in its release notes and server version response
    - PostgreSQL 18 stable artifacts and examples use 1.18.0
    - flow:release accepts and publishes the complete 0.19.0 version consistently
```
