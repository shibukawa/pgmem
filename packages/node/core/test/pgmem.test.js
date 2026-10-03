import { after, before, test } from "node:test";
import assert from "node:assert/strict";
import { execFileSync, spawn } from "node:child_process";
import { mkdtempSync, writeFileSync } from "node:fs";
import { createRequire } from "node:module";
import { tmpdir } from "node:os";
import { dirname, join } from "node:path";
import { createInterface } from "node:readline";
import { fileURLToPath, pathToFileURL } from "node:url";
import pg from "pg";
import { PgmemClient, PgmemError, PgmemServer, pgmemTest, shadowPg, startTestApp, withShadowPg, withTestDatabase } from "../index.js";
import { connectWire } from "./wire.js";

const pkg = join(dirname(fileURLToPath(import.meta.url)), "..");
const repo = join(pkg, "..", "..", "..");

let server;

before(async () => {
  if (!process.env.PGMEM_BINARY) {
    const bin = join(mkdtempSync(join(tmpdir(), "pgmem-node-")), process.platform === "win32" ? "pgmem.exe" : "pgmem");
    execFileSync("go", ["build", "-o", bin, "./cmd/pgmem"], { cwd: repo, stdio: "inherit" });
    process.env.PGMEM_BINARY = bin;
  }
  server = await PgmemServer.start({
    database: "app",
    async prepare(template) {
      const db = await connectWire(template.url);
      await db.query("CREATE TABLE t (v int); INSERT INTO t VALUES (1), (2)");
      await db.close();
    },
  });
});

after(() => server?.close());

async function count(url) {
  const db = await connectWire(url);
  try {
    return Number((await db.query("SELECT count(*) FROM t"))[0][0]);
  } finally {
    await db.close();
  }
}

async function refused(url, withinMs = 10000) {
  for (const deadline = Date.now() + withinMs; Date.now() < deadline; await new Promise((r) => setTimeout(r, 50))) {
    try {
      const db = await connectWire(url);
      await db.close();
    } catch {
      return true;
    }
  }
  return false;
}

function runNode(args, env) {
  return new Promise((resolve) => {
    const child = spawn(process.execPath, args, { env: { ...process.env, ...env }, stdio: ["ignore", "pipe", "pipe"] });
    let stdout = "";
    let stderr = "";
    child.stdout.on("data", (d) => (stdout += d));
    child.stderr.on("data", (d) => (stderr += d));
    child.on("exit", (code) => resolve({ code, stdout, stderr }));
  });
}

test("start prepares the template and snapshots it", async () => {
  assert.match(server.url, /^postgres:\/\/postgres@127\.0\.0\.1:\d+\/app\?sslmode=disable$/);
  assert.equal(server.template.database, "app");
  assert.match(server.controlUrl, /^pgmem-control:\/\/[0-9a-f]{32}@127\.0\.0\.1:\d+$/);
  assert.deepEqual(server.env(), { PGMEM_CONTROL: server.controlUrl, PGMEM_SNAPSHOT: server.snapshot.id });
  assert.equal(await count(server.url), 2);
});

test("forks are isolated and describe themselves as environment variables", async () => {
  const a = await server.fork();
  const b = await server.fork();
  try {
    const db = await connectWire(a.url);
    await db.query("INSERT INTO t VALUES (3)");
    await db.close();
    assert.equal(await count(a.url), 3);
    assert.equal(await count(b.url), 2);
    assert.equal(await count(server.url), 2);
    assert.deepEqual(a.env(), { DATABASE_URL: a.url });
    assert.deepEqual(a.env(["DIRECT_URL", "PGHOST", "PGPORT", "PGUSER", "PGDATABASE"]), {
      DIRECT_URL: a.url,
      PGHOST: "127.0.0.1",
      PGPORT: String(a.port),
      PGUSER: "postgres",
      PGDATABASE: "app",
    });
  } finally {
    await a.close();
    await b.close();
  }
});

test('read-only connection URLs reject persistent writes and keep reads available', async () => {
  await server.withFork(async (fork) => {
    const db = await connectWire(fork.env(undefined, { readOnly: true }).DATABASE_URL);
    try {
      assert.equal((await db.query('SELECT count(*) FROM t'))[0][0], '2');
      await assert.rejects(db.query('INSERT INTO t VALUES (99)'), (err) => err.code === '25006');
      await fork.reset();
      await assert.rejects(db.query('INSERT INTO t VALUES (100)'), (err) => err.code === '25006');
      assert.match(fork.env(['PGHOST'], { readOnly: true }).PGOPTIONS, /default_transaction_read_only=on/);
    } finally { await db.close(); }
  });
});

test('register times out on a full fork pool using its configured acquisition deadline', async () => {
  const limited = await PgmemServer.start({ maxForks: 1 });
  const occupied = await limited.fork();
  try {
    const started = Date.now();
    const result = await runNode(['--import', join(pkg, 'register.js'), '-e', 'console.log("should not run")'], {
      ...limited.env(), PGMEM_FORK_TIMEOUT_MS: '40',
    });
    assert.notEqual(result.code, 0);
    assert.match(result.stderr, /pool_timeout/);
    assert.ok(Date.now() - started < 5000);
  } finally { await occupied.close(); await limited.close(); }
});

test("reset restores the data under an open connection", async () => {
  await server.withFork(async (fork) => {
    const db = await connectWire(fork.url);
    try {
      await db.query("INSERT INTO t VALUES (3)");
      await fork.reset();
      assert.equal((await db.query("SELECT count(*) FROM t"))[0][0], "2");

      await db.query("INSERT INTO t VALUES (4)");
      const seeded = await fork.snapshot();
      await db.query("INSERT INTO t VALUES (5)");
      await fork.reset({ snapshot: seeded });
      assert.equal((await db.query("SELECT count(*) FROM t"))[0][0], "3");

      await db.query("BEGIN");
      await assert.rejects(fork.reset({ timeoutMs: 200 }), (err) => err instanceof PgmemError && err.code === "busy");
      await db.query("ROLLBACK");
    } finally {
      await db.close();
    }
  });
});

test("withFork closes the fork", async () => {
  let url;
  await server.withFork((fork) => {
    url = fork.url;
  });
  assert.ok(await refused(url));
});

test("withFork retains a callback failure when its child exits before cleanup", async () => {
  const pg = await PgmemServer.start();
  const failure = new Error('application assertion');
  try {
    await assert.rejects(pg.withFork(async () => {
      await pg.close();
      throw failure;
    }), (err) => err === failure);
  } finally { await pg.close(); }
});

test("withTestDatabase routes and restores the URL", async () => {
  const previousControl = process.env.PGMEM_CONTROL;
  const previousSnapshot = process.env.PGMEM_SNAPSHOT;
  const previousUrl = process.env.DATABASE_URL;
  Object.assign(process.env, server.env());
  process.env.DATABASE_URL = "outside";
  try {
    let freshUrl;
    await withTestDatabase(async (target) => {
      freshUrl = target.url;
      assert.equal(process.env.DATABASE_URL, target.url);
      assert.equal(await count(target.url), 2);
      const db = await connectWire(target.url);
      await db.query("INSERT INTO t VALUES (3)");
      await db.close();
    });
    assert.equal(process.env.DATABASE_URL, "outside");
    assert.ok(await refused(freshUrl));
    await assert.rejects(withTestDatabase(async () => {
      assert.notEqual(process.env.DATABASE_URL, "outside");
      throw new Error("case failed");
    }), /case failed/);
    assert.equal(process.env.DATABASE_URL, "outside");
    let sharedUrl;
    await withTestDatabase(async (target) => { sharedUrl = target.url; assert.equal(await count(target.url), 2); }, { fork: false });
    await withTestDatabase(async (target) => { assert.equal(target.url, sharedUrl); }, { fork: false });
  } finally {
    for (const [name, value] of Object.entries({ PGMEM_CONTROL: previousControl, PGMEM_SNAPSHOT: previousSnapshot, DATABASE_URL: previousUrl })) {
      if (value === undefined) delete process.env[name];
      else process.env[name] = value;
    }
  }
});

test("pgmemTest registers a callback with test context", async () => {
  const oldControl = process.env.PGMEM_CONTROL;
  const oldSnapshot = process.env.PGMEM_SNAPSHOT;
  Object.assign(process.env, server.env());
  try {
    let registered;
    pgmemTest((name, callback) => { assert.equal(name, "reads seed"); registered = callback; }, "reads seed", async (context) => {
      assert.equal(context.id, 7);
      assert.equal(await count(process.env.DATABASE_URL), 2);
    });
    await registered({ id: 7 });
  } finally {
    if (oldControl === undefined) delete process.env.PGMEM_CONTROL;
    else process.env.PGMEM_CONTROL = oldControl;
    if (oldSnapshot === undefined) delete process.env.PGMEM_SNAPSHOT;
    else process.env.PGMEM_SNAPSHOT = oldSnapshot;
  }
});

test("shadowPg redirects an existing pg Pool without changing DATABASE_URL", async () => {
  const oldControl = process.env.PGMEM_CONTROL;
  const oldSnapshot = process.env.PGMEM_SNAPSHOT;
  const oldUrl = process.env.DATABASE_URL;
  Object.assign(process.env, server.env());
  process.env.DATABASE_URL = "postgres://bad@127.0.0.1:1/production";
  const pool = new pg.Pool({ connectionString: process.env.DATABASE_URL, max: 1 });
  try {
    await withShadowPg(async () => {
      assert.equal((await pool.query("SELECT count(*) AS n FROM t")).rows[0].n, "2");
      await pool.query("INSERT INTO t VALUES (3)");
      assert.equal((await pool.query("SELECT count(*) AS n FROM t")).rows[0].n, "3");
    });
    await withShadowPg(async () => {
      assert.equal(process.env.DATABASE_URL, "postgres://bad@127.0.0.1:1/production");
      assert.equal((await pool.query("SELECT count(*) AS n FROM t")).rows[0].n, "2");
    });
  } finally {
    await pool.end();
    for (const [name, value] of Object.entries({ PGMEM_CONTROL: oldControl, PGMEM_SNAPSHOT: oldSnapshot, DATABASE_URL: oldUrl })) {
      if (value === undefined) delete process.env[name];
      else process.env[name] = value;
    }
  }
});

test("shadowPg registers a case and redirects pg Client", async () => {
  const oldControl = process.env.PGMEM_CONTROL;
  const oldSnapshot = process.env.PGMEM_SNAPSHOT;
  Object.assign(process.env, server.env());
  try {
    let registered;
    shadowPg((name, callback) => { assert.equal(name, "uses application client"); registered = callback; },
      "uses application client", async () => {
        const client = new pg.Client("postgres://bad@127.0.0.1:1/production");
        await client.connect();
        try { assert.equal((await client.query("SELECT count(*) AS n FROM t")).rows[0].n, "2"); }
        finally { await client.end(); }
      });
    await registered();
  } finally {
    if (oldControl === undefined) delete process.env.PGMEM_CONTROL;
    else process.env.PGMEM_CONTROL = oldControl;
    if (oldSnapshot === undefined) delete process.env.PGMEM_SNAPSHOT;
    else process.env.PGMEM_SNAPSHOT = oldSnapshot;
  }
});

test("a long-lived HTTP process keeps its connection while tests reset the data", async () => {
  const target = await server.fork();
  const wireModule = pathToFileURL(join(pkg, "test", "wire.js")).href;
  const script = `
    import { createServer } from "node:http";
    import { connectWire } from ${JSON.stringify(wireModule)};
    const db = await connectWire(process.env.DATABASE_URL);
    const app = createServer(async (_req, res) => {
      try {
        const rows = await db.query("SELECT count(*) FROM t");
        res.setHeader("content-type", "application/json");
        res.end(JSON.stringify({ count: Number(rows[0][0]) }));
      } catch (err) {
        res.statusCode = 500;
        res.end(String(err));
      }
    });
    app.listen(Number(process.env.PORT), "127.0.0.1");
  `;
  const app = await startTestApp({ fork: target, args: ["--input-type=module", "-e", script], healthPath: "/" });
  try {
    const getCount = async () => {
      const response = await fetch(app.url);
      const body = await response.text();
      assert.equal(response.status, 200, body);
      return JSON.parse(body).count;
    };
    assert.equal(await getCount(), 2);
    const db = await connectWire(target.url);
    await db.query("INSERT INTO t VALUES (3)");
    await db.close();
    assert.equal(await getCount(), 3);
    await target.reset();
    assert.equal(await getCount(), 2);
  } finally {
    await app.close();
    await target.close();
  }
});

test("startTestApp reports a child that exits before readiness", async () => {
  const target = await server.fork();
  try {
    await assert.rejects(
      startTestApp({ fork: target, args: ["-e", "process.exit(3)"], timeoutMs: 2000 }),
      /app exited before/,
    );
  } finally {
    await target.close();
  }
});

test("a worker process forks through the control socket and its fork ends with it", async () => {
  const script = `
    const { useFork } = await import(${JSON.stringify(pathToFileURL(join(pkg, "index.js")).href)});
    await useFork();
    console.log(JSON.stringify({ url: process.env.DATABASE_URL, direct: process.env.DIRECT_URL }));
    setInterval(() => {}, 1000);
  `;
  const child = spawn(process.execPath, ["--input-type=module", "-e", script], {
    env: { ...process.env, ...server.env(), PGMEM_ENV: "DATABASE_URL, DIRECT_URL" },
    stdio: ["ignore", "pipe", "inherit"],
  });
  const line = await new Promise((resolve, reject) => {
    createInterface({ input: child.stdout }).once("line", resolve);
    child.once("exit", (code) => reject(new Error(`worker exited with ${code}`)));
  });
  const { url, direct } = JSON.parse(line);
  assert.equal(direct, url);
  assert.notEqual(url, server.url);
  assert.equal(await count(url), 2);
  child.kill();
  assert.ok(await refused(url), "fork still accepts connections after its worker died");
});

test("the register entry sets DATABASE_URL before the entry point and lets the process exit", async () => {
  const res = await runNode(["--import", pathToFileURL(join(pkg, "register.js")).href, "-e", "console.log(process.env.DATABASE_URL)"], server.env());
  assert.equal(res.code, 0, res.stderr);
  assert.match(res.stdout.trim(), /^postgres:\/\/postgres@127\.0\.0\.1:\d+\/app\?sslmode=disable$/);
});

test("node:test wrappers compose in either order and both clean up", async () => {
  const entry = pathToFileURL(join(pkg, "index.js")).href;
  const script = `
    import { test } from 'node:test';
    import { withTestReset } from ${JSON.stringify(entry)};
    let otherCleanups = 0;
    const other = (fn) => async (...args) => { try { return await fn(...args); } finally { otherCleanups++; } };
    let calls = 0;
    const fork = (await import(${JSON.stringify(entry)})).currentFork();
    const reset = fork.reset.bind(fork);
    fork.reset = async (...args) => { calls++; return reset(...args); };
    test('other outside pgmem', other(withTestReset(async () => {})));
    test('pgmem outside other', withTestReset(other(async () => {})));
    test('cleanup ran', () => {
      if (calls !== 2 || otherCleanups !== 2) throw Error('cleanup counts: ' + calls + ', ' + otherCleanups);
    });
  `;
  const file = join(mkdtempSync(join(tmpdir(), "pgmem-compose-")), "compose.test.mjs");
  writeFileSync(file, script);
  const res = await runNode(["--test", "--import", pathToFileURL(join(pkg, "register.js")).href, file], server.env());
  assert.equal(res.code, 0, res.stderr);
});

test("the CommonJS entry works without the ESM one", async () => {
  const cjs = createRequire(import.meta.url)("../index.cjs");
  const client = await cjs.PgmemClient.connect({ controlUrl: server.controlUrl, snapshot: server.snapshot.id });
  try {
    await client.withFork(async (fork) => assert.equal(await count(fork.url), 2));
  } finally {
    client.close();
  }
  assert.throws(() => cjs.currentFork(), /no fork for this test file/);
});

test("Jest environment factory composes around another service in either order", async () => {
  const factory = pathToFileURL(join(pkg, 'jest-environment-factory.cjs')).href;
  const environment = pathToFileURL(join(pkg, 'environment.cjs')).href;
  const script = `
    import assert from 'node:assert/strict';
    import { withPgmemEnvironment } from ${JSON.stringify(factory)};
    import { installServiceEnv } from ${JSON.stringify(environment)};
    let cleaned = 0;
    class Base {
      global = { process: { env: {} } };
      async setup() {}
      async teardown() {}
    }
    const other = (Parent) => class extends Parent {
      async setup() {
        await super.setup();
        this.releaseValkey = installServiceEnv('valkey', { VALKEY_URL: 'mock://valkey' }, {
          target: this.global.process.env, scope: this.global,
        });
      }
      async teardown() { cleaned++; try { this.releaseValkey?.(); } finally { await super.teardown(); } }
    };
    for (const Env of [other(withPgmemEnvironment(Base)), withPgmemEnvironment(other(Base))]) {
      const env = new Env();
      await env.setup();
      assert.match(env.global.process.env.DATABASE_URL, /^postgres:/);
      assert.equal(env.global.process.env.VALKEY_URL, 'mock://valkey');
      await env.teardown();
      assert.deepEqual(env.global.process.env, {});
    }
    assert.equal(cleaned, 2);
  `;
  const res = await runNode(['--input-type=module', '-e', script], server.env());
  assert.equal(res.code, 0, res.stderr);
});

test("a wrong control token is refused", async () => {
  const url = new URL(server.controlUrl);
  url.username = "0".repeat(32);
  await assert.rejects(PgmemClient.connect({ controlUrl: url.href, snapshot: server.snapshot.id }), (err) => err.code === "unauthorized");
});

test("a failing prepare stops the process", async () => {
  await assert.rejects(
    PgmemServer.start({
      prepare() {
        throw new Error("migration failed");
      },
    }),
    /migration failed/,
  );
});
