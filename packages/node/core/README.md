# @pgmem/core

A real PostgreSQL that keeps everything in memory, packaged as a single
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

Setting `DATABASE_URL` per test case does not move an ORM client that was
already created; it keeps the URL it read during construction.

For a client built **inside** each test callback, `pgmemTest(test, name, fn)`
routes the callback to a fresh fork by default and closes it afterwards.
`withTestDatabase(fn, { fork: false })` reuses a shared fork for read-only
cases. Both helpers temporarily change `process.env`, serialize their own
callbacks, and require clients to be created inside the callback. They do
not move an already connected pool or coordinate with unrelated concurrent
code in the same process.

For API tests, register a fork before importing the app, keep the app and
its pool alive, and call `currentFork().reset()` before each sequential
case. For Playwright E2E tests, start a pgmem server, fork, and app process
per Playwright worker; `startTestApp({ fork, args: ['./src/server.js'] })`
chooses a port, passes the fork URL to the app, waits for `/health`, and
closes the process after use. Reset the fork
before each browser test. See the [unit](https://shibukawa.github.io/pgmem/guides/nodejs/testing/),
[API](https://shibukawa.github.io/pgmem/guides/nodejs/api-testing/), and
[E2E](https://shibukawa.github.io/pgmem/guides/nodejs/e2e-testing/) guides.

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

## Compose test services

Wrap a sequential writing callback with `withTestReset(fn, { snapshot, timeoutMs })`.
It restores the file fork after success or failure and preserves both errors
if reset also fails. Compose it with other callback wrappers in your own helper.
Overlapping wrapped callbacks on the shared fork fail explicitly; concurrent
writing cases use `withFork` and their own clients.

Jest can retain a custom environment:

```js
const { TestEnvironment } = require("jest-environment-node");
const { withPgmemEnvironment } = require("@pgmem/core/jest-environment-factory");
module.exports = withPgmemEnvironment(TestEnvironment);
```

Another service's environment can be the base class or wrap the returned class.
Use distinct endpoint variable names and inject them before application imports.
The factory accepts `{ env: ["APP_DATABASE_URL"], forkTimeoutMs: 30000, resetTimeoutMs: 5000, readOnly: false }`. `timeoutMs` is a deprecated alias for the reset deadline.

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

- A fork is a PostgreSQL cluster of its own: every connection gets its
  own backend process, so pools, locks between connections and deadlock
  detection behave as on a server.
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
| `PgmemServer.start(options)` | `database`, `user`, `params`, `prepare`, `maxForks`, `log`, `binary` |
| `server.url`, `server.template`, `server.snapshot` | the prepared template and its snapshot |
| `server.env()` | `{ PGMEM_CONTROL, PGMEM_SNAPSHOT }` for test processes |
| `server.fork()`, `server.withFork(fn)`, `server.close()` | |
| `useFork()` | what `@pgmem/core/register` runs: fork (or reset) and write `DATABASE_URL`, or the names in `PGMEM_ENV` |
| `withTestReset(fn, options)` | restore after a sequential writing callback |
| `currentFork()` | the fork of this test file |
| `withTestDatabase(fn, { fork, env })` | route a callback to a fresh fork (default) or a shared read-only fork |
| `pgmemTest(test, name, fn, options)` | register a test using `withTestDatabase` |
| `startTestApp({ fork, command, args, healthPath, ... })` | launch an HTTP app with the fork URL and free port; wait for readiness; close it afterward |
| `fork()`, `withFork(fn)`, `connect()` | fork through `PGMEM_CONTROL` from any process |
| `fork.url`, `fork.env(names)` | `DATABASE_URL` by default; `PGHOST`, `PGPORT`, `PGUSER`, `PGDATABASE` get the parts |
| `fork.reset({ snapshot, timeoutMs })` | back to the snapshot in place |
| `fork.snapshot()` | a snapshot of this fork to reset to later |
| `fork.close()` | |

## Automatic waits and service ownership / 自動待機とサービスの所有

`useFork`, registration and Jest acquisition default to 30 seconds; set
`forkTimeoutMs` or `PGMEM_FORK_TIMEOUT_MS`. Reused worker resets default to
5 seconds; set `resetTimeoutMs` or `PGMEM_RESET_TIMEOUT_MS`. Explicit options
take precedence over environment variables. Manual `fork()` remains unbounded
when `timeoutMs` is omitted.

`@pgmem/core/environment` exports `installServiceEnv(service, values, options)`
and `ServiceEnvConflict`. Registered services cannot claim the same variable;
installation rolls back partial changes and release restores prior values.
Use `{ target: this.global.process.env, scope: this.global }` in a Jest realm.
Direct assignments from unregistered services cannot be attributed.

`readOnly: true` or `PGMEM_READ_ONLY=true` guards injected connection URLs with
`default_transaction_read_only`. This session default can be deliberately
changed; use writable forks for writing tests. Cleanup preserves failures from
both the test and teardown with `AggregateError`.

自動取得の期限は既定 30 秒で、`forkTimeoutMs` または
`PGMEM_FORK_TIMEOUT_MS` で指定します。再利用する worker の reset は既定 5 秒で、
`resetTimeoutMs` または `PGMEM_RESET_TIMEOUT_MS` を使います。明示した option が
環境変数より優先します。手動の `fork()` は `timeoutMs` を省略すると無期限です。

`@pgmem/core/environment` の `installServiceEnv(service, values, options)` で
所有を登録すると、別の登録済みサービスによる同じ変数の取得を
`ServiceEnvConflict` で拒否します。途中の変更は失敗時に戻し、解放時は元の値を
復元します。Jest 内では `{ target: this.global.process.env, scope: this.global }`
を使います。登録せず直接代入するサービスの所有は識別できません。

`readOnly: true` または `PGMEM_READ_ONLY=true` は注入する URL に
`default_transaction_read_only` を付けます。意図的に変更できるセッションの既定値
なので、書き込みテストは書き込み可能なフォークを使います。テストと後片付けの
両方の失敗は `AggregateError` に残します。

Runnable application examples / 実行可能なアプリのサンプル:
[examples](../../../examples/README.md).
