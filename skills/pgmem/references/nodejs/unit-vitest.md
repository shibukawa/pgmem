# Vitest setup

Use `test.globalSetup: ['./test/global-setup.ts']` to start `PgmemServer` with migrations/seed in `prepare`, then `Object.assign(process.env, server.env())`; return `() => server.close()`. Put `@pgmem/core/register` first in `test.setupFiles` so each test file forks before application imports. Tests within a file share that file fork; use `shadowPg` or reset for writing cases, and keep cases sharing a Pool sequential. See [unit lifecycle](unit.md) and `website/src/content/docs/guides/nodejs/testing.mdx`.
