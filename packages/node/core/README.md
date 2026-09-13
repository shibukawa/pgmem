# @pgmem/core

A real PostgreSQL 18 that keeps everything in memory, packaged as a single
binary and driven from Node.js. Prepare the schema once per test run, give
every test file its own copy through `DATABASE_URL`, and put that copy back
between tests without reconnecting.

Application code does not change: Prisma, Drizzle, TypeORM, Kysely, `pg`
and `postgres` connect to the URL like to any PostgreSQL, pools included.

## How it fits together

- **Once per run** (global setup): `PgmemServer.start()` starts pgmem, runs
  your migrations against the template database and snapshots it. Export
  `server.env()` so test processes can reach it.
- **Per test file**: `@pgmem/core/register` forks the snapshot (about 15 ms)
  and writes the fork's URL to `DATABASE_URL` before the test file is
  imported, so an ORM client created at import time uses the fork. Test
  files run in parallel, each on its own fork.
- **Per test** (optional): `currentFork().reset()` puts the fork back to the
  snapshot in place. The URL and pooled connections stay valid.
- **Concurrent tests in one file**: `withFork(fn)` hands a separate fork to
  code that takes a URL.

Setting `DATABASE_URL` per test case would not help: an ORM client reads
the URL once, when it is created, and keeps it.

## Vitest

```ts
// test/global-setup.ts
import { execFileSync } from "node:child_process";
import { PgmemServer } from "@pgmem/core";

export async function setup() {
  const pg = await PgmemServer.start({
    database: "app",
    prepare({ url }) {
      execFileSync("npx", ["prisma", "migrate", "deploy"], {
        stdio: "inherit",
        env: { ...process.env, DATABASE_URL: url },
      });
    },
  });
  Object.assign(process.env, pg.env()); // PGMEM_CONTROL, PGMEM_SNAPSHOT
  return () => pg.close();
}
```

```ts
// vitest.config.ts
import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    globalSetup: ["./test/global-setup.ts"],
    setupFiles: ["@pgmem/core/register", "./test/reset.ts"],
  },
});
```

```ts
// test/reset.ts: only if every test should start from the snapshot
import { beforeEach } from "vitest";
import { currentFork } from "@pgmem/core";

beforeEach(() => currentFork().reset());
```

The global setup is an ordinary function, so servers of other kinds
(osmem, valkeymem, ...) start next to pgmem in the same file.

## Jest

```js
// jest.config.js
module.exports = {
  globalSetup: "./test/global-setup.js",
  globalTeardown: "./test/global-teardown.js",
  testEnvironment: "@pgmem/core/jest-environment",
  setupFilesAfterEnv: ["./test/reset.js"], // optional, per-test reset
};

// test/global-setup.js
const { PgmemServer } = require("@pgmem/core");
module.exports = async () => {
  globalThis.pgmem = await PgmemServer.start({ database: "app", prepare: ({ url }) => migrate(url) });
  Object.assign(process.env, globalThis.pgmem.env());
};

// test/global-teardown.js
module.exports = () => globalThis.pgmem.close();

// test/reset.js
const { currentFork } = require("@pgmem/core");
beforeEach(() => currentFork().reset());
```

The environment gives each Jest worker one fork and resets it between the
test files the worker runs. It needs `jest-environment-node`, which Jest
installs.

## node:test

```sh
node --test --test-global-setup=./test/global-setup.mjs --import @pgmem/core/register
```

```js
// test/global-setup.mjs (Node 24 or newer)
import { PgmemServer } from "@pgmem/core";

let pg;
export async function globalSetup() {
  pg = await PgmemServer.start({ database: "app", prepare: ({ url }) => migrate(url) });
  Object.assign(process.env, pg.env());
}
export async function globalTeardown() {
  await pg.close();
}
```

Each test file runs in its own process, so `--import` gives every file its
fork. For per-test resets add `beforeEach(() => currentFork().reset())`.

## Bun

`bun test` has no global setup and shares one process by default. Start
pgmem in a small script and run the tests with `--isolate`:

```js
// test/run.mjs
import { spawnSync } from "node:child_process";
import { PgmemServer } from "@pgmem/core";

const pg = await PgmemServer.start({ database: "app", prepare: ({ url }) => migrate(url) });
const run = spawnSync("bun", ["test", "--isolate", "--preload", "@pgmem/core/register"], {
  stdio: "inherit",
  env: { ...process.env, ...pg.env() },
});
await pg.close();
process.exit(run.status ?? 1);
```

## Preparing the schema

`prepare` receives the template server (`url`, `host`, `port`, ...). Close
or commit what it opens: the snapshot waits for open transactions.

- **Prisma 7**: run `prisma migrate deploy` or `prisma db push` with
  `DATABASE_URL` in the child's environment (`prisma.config.ts` reads it with
  `env("DATABASE_URL")`, and `dotenv` does not override a variable that is
  already set). `prisma migrate dev` works as well: Prisma creates its
  shadow database on the same pgmem server. [`examples/prisma`](../examples/prisma)
  runs both, and has a script that authors migrations against pgmem
  instead of a local PostgreSQL.
- **Drizzle**: `const db = drizzle(url); await migrate(db, { migrationsFolder: "drizzle" }); await db.$client.end();`
- **TypeORM**: `const ds = await new DataSource({ ...options, url }).initialize(); await ds.runMigrations(); await ds.destroy();`

## Things to know

- A fork is one PostgreSQL session shared by all its connections, the way
  a transaction-mode pooler shares one: pools work, but `SET`, temp tables
  and advisory locks are shared (use `SET LOCAL`).
- Inside a transaction callback, query through the transaction handle. A
  query on another pooled connection would wait for the transaction to end;
  pgmem ends that wait after `waitTimeoutMs` (default 2 s) with SQLSTATE
  55P03 and a message naming both connections.
- `reset` waits for open transactions and fails with code `busy` after
  `timeoutMs` (default 5 s).
- Keep `sslmode=disable` in the URL: `pg` treats `prefer` and `require` as
  "TLS required", and pgmem has no TLS.
- Forks made through `@pgmem/core/register`, `fork()` or `withFork()` belong
  to the process that made them and close when it exits. Closing a fork
  leaves idle pooled connections open until their pool closes them, so a
  pool without an `error` listener does not crash the test process.
- `PGMEM_BINARY` points at a local build of
  [`cmd/pgmem`](https://github.com/shibukawa/pgmem) for platforms without a
  binary package.

## API

| | |
|---|---|
| `PgmemServer.start(options)` | `database`, `user`, `params`, `prepare`, `maxForks`, `waitTimeoutMs`, `log`, `binary` |
| `server.url`, `server.template`, `server.snapshot` | the prepared template and its snapshot |
| `server.env()` | `{ PGMEM_CONTROL, PGMEM_SNAPSHOT }` for test processes |
| `server.fork()`, `server.withFork(fn)`, `server.close()` | |
| `useFork()` | what `@pgmem/core/register` runs: fork (or reset) and write `DATABASE_URL`, or the names in `PGMEM_ENV` |
| `currentFork()` | the fork of this test file |
| `fork()`, `withFork(fn)`, `connect()` | fork through `PGMEM_CONTROL` from any process |
| `fork.url`, `fork.env(names)` | `DATABASE_URL` by default; `PGHOST`, `PGPORT`, `PGUSER`, `PGDATABASE` get the parts |
| `fork.reset({ snapshot, timeoutMs })` | back to the snapshot in place |
| `fork.snapshot()` | a snapshot of this fork to reset to later |
| `fork.close()` | |
