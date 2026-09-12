---
id: requirement:nodejs-wrapper
type: requirement
title: Node.js Wrapper
---
An npm package that gives Drizzle, Prisma and TypeORM tests a DATABASE_URL backed by pgmem, positioned as a test database rather than a PGlite replacement.

```yaml
summary:
  status: planned (direction stated 2026-09-12); after api:python-wrapper and api:java-wrapper
  targets: [Drizzle, Prisma, TypeORM, node-postgres, postgres.js]
  interface: DSN only; no per-ORM adapter; connection pools work; a fork per test worker (vitest, jest)
  positioning: PGlite is in-process JavaScript with per-ORM adapters and no wire protocol; pgmem is a real listener any ORM connects to (concept:alternatives)
  distribution: optionalDependencies per platform, esbuild style (policy:binary-distribution); avoid the npm name pg-mem (existing emulator)
  blockers:
    - Prisma migrate dev needs a shadow database, so multi-database serving must land first (concept:limits)
    - extensions common in cloud PostgreSQL that Prisma and Drizzle schemas assume (requirement:cloud-common-extensions)
  mechanism: same concept:server-process and api:control-protocol; lifecycle flow:wrapper-test-lifecycle
```
