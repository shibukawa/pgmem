# Fastify inject

Create `const app = buildApp()` after the register hook. Use `beforeEach(() => currentFork().reset())`, `afterAll(() => app.close())`, then call `await app.inject({method: 'POST', url: '/orders', payload: {sku: 'book'}})` and assert on `statusCode` and `json()`.

For per-case forks, wrap `app.inject(...)` in `shadowPg(test, name, callback)`; its async context reaches `pg` when inject runs in-process. Do not overlap cases using one app and Pool. See [API lifecycle](nodejs-api.md) and `website/src/content/docs/guides/nodejs/api-testing.mdx`.
