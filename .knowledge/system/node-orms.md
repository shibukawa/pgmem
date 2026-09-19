---
id: system:node-orms
type: system
title: Node.js ORMs and Drivers
---
Where Node.js ORMs read the database URL and how they behave against pgmem; compatibility PoC run 2026-09-13 against a cmd/pgmem build.

```yaml
system:
  prisma_7_10:
    runtime: new PrismaClient({adapter: new PrismaPg({connectionString})}); PrismaPg also takes a pg.Pool; adapter-pg ignores ?schema=
    cli: Rust schema-engine; URL from prisma.config.ts datasource.url; no .env autoload
    verified: migrate deploy 0.6s, migrate status, migrate diff config datasource vs schema empty, db push twice (second already in sync), runtime CRUD, nested create, interactive transaction and rollback, raw pgvector, 30 concurrent creates, P2002, groupBy
    migrate_dev: works (since 2026-09-13; every database of a server can be connected to); Prisma creates prisma_migrate_shadow_db_<uuid> on the same server, uses and drops it; verified init 1.3s, a second migration 0.6s, an already-in-sync run, migrate status, db pull; before that it failed with P1003
    shadow_config: migrate diff --from-migrations requires shadowDatabaseUrl, which may name another database of the same server (verified)
    migrate_reset: verified 2026-09-13 with the user's consent (Prisma refuses it for AI agents otherwise, PRISMA_USER_CONSENT_FOR_DANGEROUS_AI_ACTION); on a throwaway pgmem database 3 users and 2 posts became 0, both migrations were re-applied in 662ms, migrate status up to date
    example: packages/node/examples/prisma (2026-09-13) authored its init migration with migrate dev through pgmem and passes 4 Vitest tests over per-file forks with per-test reset
    npm: prisma@latest is 8.0.0-rc, so pin 7
  prisma_8_rc: runtime postgres({contractJson, url}) over pg.Pool; TypeScript migrations with --db <url>; no shadow database
  drizzle_0_45:
    runtime: drizzle(url) builds a pg.Pool with no 'error' listener; close with $client.end()
    verified: drizzle-kit generate and migrate, push reports no changes, programmatic migrate 21ms, node-postgres and postgres-js CRUD, pgvector cosineDistance, named .prepare x20 across the pool, transaction rollback, 50 concurrent inserts
    crash: closing a fork while the pool holds an idle client kills Node with an unhandled 'error' (Connection terminated unexpectedly)
    note: 1.0 is rc and changes the migrations folder layout
  typeorm_1_1:
    runtime: new DataSource({type postgres, url}); its SELECT version() regex matches pgmem; CREATE EXTENSION only for column types in use; attaches a pool error handler
    verified: synchronize with uuid-ossp, citext, hstore, ltree 91ms; schema diff after synchronize empty; relations, hstore object, text[], version column, numeric SUM, 50 concurrent saves, runMigrations on a fork
  pg_8_23: sslmode prefer or require means TLS required and fails on pgmem's 'N'; keep sslmode=disable; Parse is sent once per statement name per connection
  postgres_js_3_4: prepare true by default; re-prepares once on 26000 or 0A000 using the ErrorResponse routine field
  hazard_outer_client_in_transaction: db.transaction(async tx => { await db.select() }) works since rule:process-per-connection (2026-09-19), as on real PostgreSQL; in the removed single-session model it ended with 55P03
  env_loading: dotenv and NestJS ConfigModule do not override variables that are already set
  with_reset_2026_09_13: through @pgmem/core, Prisma's pooled client, drizzle named .prepare and postgres.js keep working across api:reset; the outer-client-in-transaction case ended with 55P03 in the single-session model (gone since 2026-09-19); a fork closed under drizzle's idle pool no longer crashes the worker
```
