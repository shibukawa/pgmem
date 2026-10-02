# Express + Supertest

With Vitest, import `request` from `supertest`, `currentFork` from `@pgmem/core`, and the application's exported `app` and `pool`. Use `beforeEach(() => currentFork().reset())` and `afterAll(() => pool.end())`. A test sends `await request(app).post('/orders').send({sku: 'book'})` and checks status/body; the next test begins from the prepared snapshot. The app keeps its normal datasource configuration.

Alternatively wrap cases in `shadowPg(test, name, async () => { await request(app)... })` for a fresh fork per case. Keep one shared Pool's cases sequential, including callbacks wrapped by `shadowPg`. See [API lifecycle](api.md) and `website/src/content/docs/guides/nodejs/api-testing.mdx`.
