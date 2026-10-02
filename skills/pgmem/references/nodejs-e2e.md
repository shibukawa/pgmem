# Node.js E2E tests with Playwright

Use worker-scoped Playwright fixtures. Start `PgmemServer.start({database: 'app', prepare: ({url}) => migrateAndSeed(url)})`, fork it, and keep both open for the worker. Start the application's child process with `startTestApp({fork, args: ['./src/server.js'], healthPath: '/health'})`; expose `app.url` as worker-scoped `appUrl`. An automatic test-scoped fixture calls `fork.reset()` before each browser case. Close app, fork, and server in reverse order at worker teardown.

`startTestApp` selects a loopback port, passes `PORT` and the fork's `DATABASE_URL`, waits for health, and shuts the child down. The app must read that DSN when constructing its pool and listen on `PORT`. Options include `command`, `portEnv`, `databaseEnv`, `healthPath`, and `timeoutMs`. For another baseline, seed and snapshot the worker fork, then reset from that snapshot; close it after use.

Set bounded Playwright workers. Each worker needs a distinct pgmem server/fork, app process, port, and browser context. A shared external app URL breaks isolation. Wait for HTTP and background database work before reset. Full fixture: `website/src/content/docs/guides/nodejs/e2e-testing.mdx`.
