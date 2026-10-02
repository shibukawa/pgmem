# Node.js unit tests

Use `@pgmem/core`. In runner global setup, `PgmemServer.start({database: 'app', prepare: ({url}) => migrateAndSeed(url)})` prepares a template once; export `server.env()` to workers, then close the server in global teardown. Close migration/seed connections before `prepare` returns. Register `@pgmem/core/register` before application imports so each test file gets a fork and `DATABASE_URL`.

For a Vitest repository test, use `shadowPg(test, 'creates a user', async () => { ... })`. It routes new `pg` Client/Pool connections, including `@prisma/adapter-pg`, to a fresh fork without changing application connection settings. Existing checked-out connections do not move. `{ fork: false }` is for read-only cases. `pgmemTest` instead changes `DATABASE_URL` temporarily; create clients inside its callback, and note that it serializes callbacks because environment is process-wide. Test repository methods, then close pools in teardown.

Read the runner setup only when needed: [Vitest](unit-vitest.md), [Jest](unit-jest.md), [node:test](unit-node-test.md), or [Bun](unit-bun.md). Other PostgreSQL drivers require their own routing adapter.

Parallelize separate files in separate workers. Cases sharing a Pool/fork must run sequentially; `AsyncLocalStorage` cannot move a checked-out connection. Full examples, including Bun: `website/src/content/docs/guides/nodejs/testing.mdx`.
