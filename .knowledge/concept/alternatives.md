---
id: concept:alternatives
type: concept
title: Alternatives and Positioning
---
How pgmem compares to the other ways of getting a PostgreSQL for tests; the basis of the why-pgmem page.

```yaml
summary:
  pgmem:
    shape: real PostgreSQL 18 in the test process (Go) or in one child process (other languages); no daemon, no disk, no download (policy:binary-distribution)
    strengths: [start ~0.1 s, fork ~20 ms with prepared schema and seed (api:clone), works on any CI runner including ones without Docker, one dependency, same wire protocol so every driver and ORM works]
    weaknesses: [one session per fork with transaction-mode multiplexing (rule:single-session-per-backend), extensions limited to the bundled set, no background workers or PostGIS (policy:bundled-extensions), ~150 MB resident per process, ~36 MB binary]
  testcontainers_docker:
    shape: Docker daemon runs the official postgres image; the test framework starts and stops containers
    strengths: [exact production image and version, any extension including PostGIS and pg_cron, real multi-session server, mature ecosystem in every language]
    weaknesses: [Docker daemon required on developer machines and CI (privileged or docker-in-docker runners), image pull on first use, seconds per container start, per-test isolation left to the user (template database or transaction rollback), VM memory on macOS and Windows]
    numbers: metric:alternative-comparison
  embedded_postgres_binaries:
    shape: zonky embedded-postgres (Java), fergusstrange/embedded-postgres (Go), pytest-postgresql: unpack a native postgres build into a temp dir and run initdb plus a postmaster
    strengths: [no Docker, real multi-session server, extensions from the bundled build]
    weaknesses: [writes to disk, initdb and postmaster start take about a second, platform binaries per package, isolation again by database per test]
  pglite:
    shape: PostgreSQL wasm in the Node or browser process, driven through a TypeScript API
    strengths: [in-process for JavaScript, browser support, pgvector and contrib bundled]
    weaknesses: [no wire protocol by default (pglite-socket is a bridge), per-ORM adapters, single connection, JavaScript only]
    relation: pgmem uses the same PGlite fork of PostgreSQL but ships it as a Go library and a binary that speaks the wire protocol (concept:architecture)
  sql_emulators:
    shape: pg-mem (JavaScript) or SQLite standing in for PostgreSQL
    weaknesses: [not PostgreSQL: dialect, planner and extension gaps surface only in production]
  choose_pgmem_when: [tests must run without Docker, suites want per-test isolation cheaper than a database per test, developers want one dependency per language]
  choose_docker_when: [tests need PostGIS or background-worker extensions, tests exercise concurrency between sessions, an exact server build must be pinned]
```
