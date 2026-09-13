---
id: api:node-wrapper
type: api
title: Node.js Wrapper
---
npm package driving concept:server-process, shaped like @osmem/core and implementing decision:node-test-integration.

```yaml
api:
  status: implemented 2026-09-13 in packages/node, unpublished; node --test suite in core/test (9 tests); end to end in a scratch project with Vitest 5 (Prisma 7.10, Drizzle 0.45 with postgres.js, TypeORM 1.1) and Jest 30
  package: '@pgmem/core plus @pgmem/<platform> (darwin-arm64, darwin-x64, linux-arm64, linux-x64, win32-arm64, win32-x64); the npm org pgmem belongs to the user (confirmed 2026-09-13); unscoped pgmem is rejected by npm name similarity to pg-mem'
  core:
    - 'PgmemServer.start({database, user, params, prepare(template), maxForks, control=true, waitTimeoutMs, log, binary}) -> server; spawns pgmem -control 127.0.0.1:0, runs prepare, snapshots the template'
    - 'server: url, template, snapshot, controlUrl, env() -> {PGMEM_CONTROL, PGMEM_SNAPSHOT}, fork(), withFork(fn), close()'
    - 'fork: id, url, host, port, user, database, env(names = [DATABASE_URL]; PGHOST PGPORT PGUSER PGDATABASE PGSSLMODE get parts), reset({snapshot, timeoutMs = 5000}), snapshot(), close()'
    - 'PgmemClient.connect({controlUrl = PGMEM_CONTROL, snapshot = PGMEM_SNAPSHOT}); connect(), fork(), withFork(fn) use one process-wide client'
    - 'useFork({env}): fork once per process (reset when it already has one) and write env; PGMEM_ENV lists names'
    - 'currentFork(): globalThis[Symbol.for("pgmem.currentFork")], so it works across module instances and Jest realms'
    - 'PgmemError.code carries the api:control-protocol error code'
  entries:
    register: ESM, top-level await useFork(); Vitest setupFiles, node --test --import, bun test --isolate --preload
    jest_environment: extends jest-environment-node TestEnvironment; one fork per Jest worker, reset between test files, URL written to this.global.process.env
  format: implementation in index.cjs so Jest test files can require it (Jest vm has no dynamic import without experimental flags); index.js re-exports for ESM; one index.d.ts
  liveness: channel handles are referenced only while a request is pending, so the control socket never keeps a test process alive; api:control-socket closes a worker's forks when it exits
  binary: resolveBinary(option, PGMEM_BINARY, @pgmem/<platform>/bin/pgmem); scripts/build-npm.sh fills platforms/*/bin (policy:binary-distribution)
  docs: packages/node/core/README.md (Vitest, Jest, node:test, Bun, ORM prepare recipes); packages/node/examples/prisma (Prisma 7 with Vitest, migrate deploy in globalSetup, per-test reset, migrate dev through pgmem)
```
