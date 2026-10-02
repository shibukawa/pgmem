# Node.js HTTP API tests

Prepare the template and register each test file's fork as in [unit tests](unit.md). For an in-process HTTP client, `shadowPg(test, name, callback)` routes new `pg` connections through the callback's async context, even if the app's Pool was constructed earlier. For a long-lived app and pool in a file, `beforeEach(() => currentFork().reset())` restores the snapshot while the endpoint remains stable; close the app/pool in `afterAll`. Do not combine concurrent cases with a shared app/reset.

Choose [Express + Supertest](api-express.md) or [Fastify inject](api-fastify.md) for request syntax and lifecycle. If the HTTP app runs in another process, use the [E2E worker model](e2e.md) instead; the test process's async context does not cross a process boundary.

To keep an extra dataset, seed the file fork, `snapshot()` it, then `reset({snapshot: seeded})` in each case; close the extra snapshot after the suite. A reset waits for open transactions, so finish request and background work first. Parallelize by worker/file, with each worker owning a fork and app instance.

Full guide: `website/src/content/docs/guides/nodejs/api-testing.mdx`.
