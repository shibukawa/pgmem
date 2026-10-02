# Bun test setup

`bun test` has no global setup and shares a process by default. Use a small launcher script: start `PgmemServer` with `prepare`, then spawn `bun test --isolate --preload @pgmem/core/register` with `{...process.env, ...server.env()}`; close the server and forward Bun's exit status. The preload creates a fork before test imports. Keep writing cases sharing one pool sequential. See [unit lifecycle](nodejs-unit.md) and `website/src/content/docs/guides/nodejs/testing.mdx`.
