# node:test setup

On Node.js 24+, provide a module exporting `globalSetup()` and `globalTeardown()`. Start `PgmemServer` in setup, run migrations/seed in `prepare`, assign `server.env()` to `process.env`, and close the server in teardown. Run `node --test --test-global-setup=./test/global-setup.mjs --import @pgmem/core/register`. Each test file process gets its own fork before imports. Keep writing cases that share a pool sequential. See [unit lifecycle](nodejs-unit.md).
