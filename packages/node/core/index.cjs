"use strict";
// @pgmem/core: starts the pgmem binary (a real PostgreSQL 18 that keeps
// everything in memory) and drives its control protocol. The
// implementation is CommonJS so that Jest can load it inside test files;
// index.js re-exports it for ESM.
const { spawn } = require("node:child_process");
const { AsyncLocalStorage } = require("node:async_hooks");
const { existsSync } = require("node:fs");
const Module = require("node:module");
const { createRequire } = Module;
const net = require("node:net");
const { join } = require("node:path");
const { createInterface } = require("node:readline");

/** Control protocol version this package speaks; the binary must match. */
const PROTOCOL = 1;
const asyncDispose = Symbol.asyncDispose ?? Symbol.for("Symbol.asyncDispose");
const CURRENT = Symbol.for("pgmem.currentFork");
const { installServiceEnv } = require("./environment.cjs");
const { fixtureOptions } = require("./fixture-config.cjs");

async function runWithCleanup(fn, cleanup, message) {
  let failed = false;
  let failure;
  let result;
  try { result = await fn(); }
  catch (err) { failed = true; failure = err; }
  try { await cleanup(); }
  catch (err) {
    if (failed) throw new AggregateError([failure, err], message);
    throw err;
  }
  if (failed) throw failure;
  return result;
}

const shadowContext = new AsyncLocalStorage();

/** An error answered by pgmem; code is the protocol error code (busy, pool_timeout, ...). */
class PgmemError extends Error {
  constructor(code, message) {
    super(`pgmem: ${code}: ${message}`);
    this.name = "PgmemError";
    this.code = code;
  }
}

/** Resolve the pgmem binary: the argument, PGMEM_BINARY, then the platform package. */
function resolveBinary(binary) {
  if (binary) return binary;
  if (process.env.PGMEM_BINARY) return process.env.PGMEM_BINARY;
  const pkg = `@pgmem/${process.platform}-${process.arch}`;
  let pkgJson;
  try {
    pkgJson = require.resolve(`${pkg}/package.json`);
  } catch {
    throw new Error(
      `pgmem: no binary for ${process.platform}-${process.arch}: install ${pkg} (an optional dependency of @pgmem/core) or set PGMEM_BINARY`,
    );
  }
  const bin = join(pkgJson, "..", "bin", process.platform === "win32" ? "pgmem.exe" : "pgmem");
  if (!existsSync(bin)) throw new Error(`pgmem: binary missing at ${bin}`);
  return bin;
}

/**
 * One control channel (the child's stdio, or a control socket): JSON lines
 * with responses matched by id, so a fork waiting for a slot never delays
 * the close that frees one. The handle is referenced only while a request
 * is pending, so an idle channel never keeps the process alive.
 */
class Channel {
  #write;
  #ref;
  #unref;
  #seq = 0;
  #pending = new Map();
  #error = null;

  constructor(lines, write, { ref, unref }) {
    this.#write = write;
    this.#ref = ref;
    this.#unref = unref;
    lines.on("line", (line) => this.#receive(line));
    lines.on("close", () => this.fail(new Error("pgmem: control channel closed")));
    unref();
  }

  #receive(line) {
    let msg;
    try {
      msg = JSON.parse(line);
    } catch {
      return;
    }
    if (msg.id == null) {
      if (msg.event === "fatal") process.stderr.write(`pgmem: fatal: ${msg.message}\n`);
      return;
    }
    const waiter = this.#pending.get(msg.id);
    if (!waiter) return;
    this.#pending.delete(msg.id);
    if (this.#pending.size === 0) this.#unref();
    if (msg.ok) waiter.resolve(msg);
    else waiter.reject(new PgmemError(msg.error?.code ?? "internal", msg.error?.message ?? "unknown error"));
  }

  request(op, fields = {}) {
    if (this.#error) return Promise.reject(this.#error);
    const id = ++this.#seq;
    return new Promise((resolve, reject) => {
      this.#pending.set(id, { resolve, reject });
      if (this.#pending.size === 1) this.#ref();
      this.#write(JSON.stringify({ id, op, ...fields }) + "\n");
    });
  }

  fail(err) {
    this.#error ??= err;
    for (const waiter of this.#pending.values()) waiter.reject(this.#error);
    this.#pending.clear();
    this.#unref();
  }
}

const envParts = {
  PGHOST: (ep) => ep.host,
  PGPORT: (ep) => String(ep.port),
  PGUSER: (ep) => ep.user,
  PGDATABASE: (ep) => ep.database,
  PGSSLMODE: () => "disable",
};

/** A listening server: its URL and the parts of it. */
class Endpoint {
  constructor(ep) {
    this.id = ep.id;
    this.url = ep.dsn;
    this.host = ep.host;
    this.port = ep.port;
    this.user = ep.user;
    this.database = ep.database;
  }

  /**
   * Environment variables pointing at this server: DATABASE_URL by default,
   * or the given names. PGHOST, PGPORT, PGUSER, PGDATABASE and PGSSLMODE get
   * the matching part; any other name gets the URL.
   */
  env(names = ["DATABASE_URL"], { readOnly = false } = {}) {
    const out = {};
    const url = new URL(this.url);
    if (readOnly) url.searchParams.set("options", [url.searchParams.get("options"), "-c default_transaction_read_only=on"].filter(Boolean).join(" "));
    if (readOnly) url.search = url.searchParams.toString().replace(/\+/g, "%20");
    for (const name of [].concat(names)) out[name] = envParts[name]?.(this) ?? url.toString();
    if (readOnly && Object.keys(out).some((name) => name in envParts)) out.PGOPTIONS = url.searchParams.get("options");
    return out;
  }
}

/** A frozen copy of a server's data that forks start from. */
class PgmemSnapshot {
  #channel;

  constructor(channel, id) {
    this.#channel = channel;
    this.id = id;
  }

  /** Start a server on a fresh copy; waits while maxForks forks are alive. */
  async fork({ timeoutMs } = {}) {
    const res = await this.#channel.request("fork", { snapshot: this.id, ...(timeoutMs != null && { timeout_ms: timeoutMs }) });
    return new PgmemFork(this.#channel, res.server);
  }

  /** Run fn with a fresh fork and close the fork afterwards. */
  async withFork(fn, options) {
    const fork = await this.fork(options);
    return runWithCleanup(() => fn(fork), () => fork.close(), "pgmem: callback and fork cleanup both failed");
  }

  /** Refuse further forks; running forks keep working. */
  async close() {
    await this.#channel.request("close", { snapshot: this.id }).catch(() => {});
  }
}

/** A server started from a snapshot. */
class PgmemFork extends Endpoint {
  #channel;
  #closed = false;

  constructor(channel, ep) {
    super(ep);
    this.#channel = channel;
  }

  /**
   * Put the data back to the snapshot the fork came from (or the given one)
   * in place. The URL and open connections stay valid; pooled connections
   * continue in a fresh session with their prepared statements intact.
   * Fails with code busy when a connection keeps a transaction open.
   */
  async reset({ snapshot, timeoutMs = 5000 } = {}) {
    await this.#channel.request("reset", {
      server: this.id,
      ...(snapshot && { snapshot: typeof snapshot === "string" ? snapshot : snapshot.id }),
      timeout_ms: timeoutMs,
    });
  }

  /** Snapshot this fork, for example after beforeAll seeding, to reset to later. */
  async snapshot({ maxForks, timeoutMs = 30000 } = {}) {
    const res = await this.#channel.request("snapshot", { server: this.id, timeout_ms: timeoutMs, ...(maxForks && { max_forks: maxForks }) });
    return new PgmemSnapshot(this.#channel, res.snapshot);
  }

  /** Stop the fork and free its slot. Idle client connections are left to their pools. */
  async close() {
    if (this.#closed) return;
    this.#closed = true;
    await this.#channel.request("close", { server: this.id }).catch((err) => {
      // A vanished child or already released id owns no resource to clean up.
      if (err instanceof PgmemError && err.code !== "unknown_id") throw err;
    });
  }

  [asyncDispose]() {
    return this.close();
  }
}

/**
 * A pgmem process owned by this process: the template server, prepared
 * once and snapshotted, and the control socket other processes fork from.
 *
 *   const server = await PgmemServer.start({ database: "app", prepare: ({ url }) => migrate(url) });
 *   Object.assign(process.env, server.env()); // PGMEM_CONTROL, PGMEM_SNAPSHOT for test workers
 */
class PgmemServer {
  #child;
  #channel;
  #exited;
  #closed = false;

  constructor(child, lines, ready) {
    this.#child = child;
    const handles = [child, child.stdin, child.stdout];
    this.#channel = new Channel(lines, (line) => child.stdin.write(line), {
      ref: () => handles.forEach((h) => h.ref?.()),
      unref: () => handles.forEach((h) => h.unref?.()),
    });
    child.stdin.on("error", () => {});
    this.#exited = new Promise((resolve) => child.once("exit", resolve));
    child.once("exit", () => this.#channel.fail(new Error("pgmem: process exited")));
    this.pid = ready.pid;
    this.version = ready.version;
    this.template = new Endpoint(ready.server);
    this.url = this.template.url;
    this.controlUrl = ready.control?.url;
    this.snapshot = undefined;
  }

  /**
   * Start pgmem, run prepare against the template (migrations, seed data;
   * close or commit every connection it opens) and snapshot it.
   */
  static async start(options = {}) {
    const bin = resolveBinary(options.binary);
    const args = [`-database=${options.database ?? "postgres"}`, `-user=${options.user ?? "postgres"}`];
    const params = Object.entries(options.params ?? {}).map(([k, v]) => `${k}=${v}`);
    if (params.length) args.push(`-params=${params.join(",")}`);
    if (options.control !== false) args.push("-control=127.0.0.1:0");
    // waitTimeoutMs is accepted for compatibility: connections no longer share a session
    if (options.log) args.push("-log");
    const child = spawn(bin, args, { stdio: ["pipe", "pipe", "inherit"], windowsHide: true });
    const lines = createInterface({ input: child.stdout });
    const ready = await new Promise((resolve, reject) => {
      const timeoutMs = options.startupTimeoutMs ?? 30000;
      const onError = (err) => fail(new Error(`pgmem: cannot start ${bin}: ${err.message}`));
      const onExit = (code) => fail(new Error(`pgmem: ${bin} exited with ${code} before it was ready`));
      const timer = setTimeout(() => fail(new Error(`pgmem: ${bin} was not ready within ${timeoutMs} ms`)), timeoutMs);
      const settle = () => {
        clearTimeout(timer);
        child.off("error", onError);
        child.off("exit", onExit);
      };
      const fail = (err) => {
        settle();
        child.kill();
        reject(err);
      };
      child.once("error", onError);
      child.once("exit", onExit);
      lines.once("line", (line) => {
        let msg;
        try {
          msg = JSON.parse(line);
        } catch {
          return fail(new Error(`pgmem: unexpected first line from ${bin}: ${line}`));
        }
        if (msg.event !== "ready" || !msg.server) return fail(new Error(`pgmem: unexpected first line from ${bin}: ${line}`));
        if (msg.protocol !== PROTOCOL) return fail(new Error(`pgmem: ${bin} speaks protocol ${msg.protocol}, @pgmem/core needs ${PROTOCOL}`));
        settle();
        resolve(msg);
      });
    });
    const server = new PgmemServer(child, lines, ready);
    try {
      if (options.prepare) await options.prepare(server.template);
      const res = await server.#channel.request("snapshot", {
        server: "template",
        timeout_ms: options.snapshotTimeoutMs ?? 30000,
        ...(options.maxForks && { max_forks: options.maxForks }),
      });
      server.snapshot = new PgmemSnapshot(server.#channel, res.snapshot);
    } catch (err) {
      return runWithCleanup(() => { throw err; }, () => server.close(), "pgmem: preparation and process cleanup both failed");
    }
    return server;
  }

  /** PGMEM_CONTROL and PGMEM_SNAPSHOT, for the processes that fork from this server. */
  env() {
    if (!this.controlUrl) throw new Error("pgmem: started with control: false, so other processes cannot fork");
    return { PGMEM_CONTROL: this.controlUrl, PGMEM_SNAPSHOT: this.snapshot.id };
  }

  /** Fork the prepared snapshot. */
  fork(options) {
    return this.snapshot.fork(options);
  }

  /** Run fn with a fork of the prepared snapshot and close it afterwards. */
  withFork(fn, options) {
    return this.snapshot.withFork(fn, options);
  }

  /** Shut the process down with every fork. Idempotent. */
  async close() {
    if (this.#closed) return;
    this.#closed = true;
    const child = this.#child;
    if (child.exitCode !== null || child.signalCode !== null) return;
    [child, child.stdin, child.stdout].forEach((h) => h.ref?.());
    child.stdin.end();
    let timedOut = false;
    const killer = setTimeout(() => { timedOut = true; child.kill("SIGKILL"); }, 10000);
    await this.#exited;
    clearTimeout(killer);
    if (timedOut) throw new Error("pgmem: process did not exit within 10000 ms; the child was killed");
  }

  [asyncDispose]() {
    return this.close();
  }
}

/**
 * A connection to a PgmemServer's control socket from another process, such
 * as a test worker. Forks made through it close when it closes, including
 * when the process exits.
 */
class PgmemClient {
  #socket;
  #channel;

  constructor(socket, snapshot) {
    this.#socket = socket;
    this.#channel = new Channel(createInterface({ input: socket }), (line) => socket.write(line), {
      ref: () => socket.ref(),
      unref: () => socket.unref(),
    });
    socket.on("error", (err) => this.#channel.fail(err));
    this.snapshot = snapshot ? new PgmemSnapshot(this.#channel, snapshot) : undefined;
  }

  /** Connect to controlUrl (default PGMEM_CONTROL) and fork from snapshot (default PGMEM_SNAPSHOT). */
  static async connect({ controlUrl = process.env.PGMEM_CONTROL, snapshot = process.env.PGMEM_SNAPSHOT } = {}) {
    if (!controlUrl) {
      throw new Error("pgmem: PGMEM_CONTROL is not set: start PgmemServer in the test runner's global setup and export server.env()");
    }
    const url = new URL(controlUrl);
    const socket = net.connect({ host: url.hostname, port: Number(url.port) });
    await new Promise((resolve, reject) => {
      socket.once("connect", resolve);
      socket.once("error", reject);
    });
    socket.setNoDelay(true);
    const client = new PgmemClient(socket, snapshot);
    try {
      const hello = await client.#channel.request("hello", { token: decodeURIComponent(url.username) });
      if (hello.protocol !== PROTOCOL) throw new Error(`pgmem: server speaks protocol ${hello.protocol}, @pgmem/core needs ${PROTOCOL}`);
    } catch (err) {
      socket.destroy();
      throw err;
    }
    return client;
  }

  #snapshot() {
    if (!this.snapshot) throw new Error("pgmem: no snapshot to fork: PGMEM_SNAPSHOT is not set");
    return this.snapshot;
  }

  /** Fork the snapshot. */
  fork(options) {
    return this.#snapshot().fork(options);
  }

  /** Run fn with a fork and close it afterwards. */
  withFork(fn, options) {
    return this.#snapshot().withFork(fn, options);
  }

  /** Close the connection; forks made through it are closed with it. */
  close() {
    this.#socket.end();
  }
}

let shared;

/** Connect to the control socket: with options, a new client; without, the process-wide one. */
function connect(options) {
  if (options) return PgmemClient.connect(options);
  shared ??= PgmemClient.connect().catch((err) => {
    shared = undefined;
    throw err;
  });
  return shared;
}

/** Fork PGMEM_SNAPSHOT through the process-wide client. */
async function fork(options) {
  return (await connect()).fork(options);
}

/** Run fn with a fork of PGMEM_SNAPSHOT and close it afterwards. */
async function withFork(fn, options) {
  return (await connect()).withFork(fn, options);
}

/**
 * Give this process its fork and write the fork's URL to process.env
 * (DATABASE_URL, or the comma-separated names in PGMEM_ENV). A process that
 * already has one resets it instead, so a worker reused across test files
 * keeps the URL its modules captured.
 */
async function useFork(options = {}) {
  const cfg = fixtureOptions(options);
  let current = globalThis[CURRENT];
  const acquired = !current;
  if (current) await current.reset({ timeoutMs: cfg.resetTimeoutMs });
  else current = await fork({ timeoutMs: cfg.forkTimeoutMs });
  try {
    const release = installServiceEnv("pgmem", current.env(cfg.env?.length ? cfg.env : undefined, { readOnly: cfg.readOnly }));
    globalThis[Symbol.for("pgmem.releaseEnv")] = release;
    globalThis[CURRENT] = current;
    return current;
  } catch (err) {
    if (acquired) return runWithCleanup(() => { throw err; }, () => current.close(), "pgmem: environment and fork cleanup both failed");
    throw err;
  }
}

/** The fork of this test file, made by @pgmem/core/register or @pgmem/core/jest-environment. */
function currentFork() {
  const current = globalThis[CURRENT];
  if (!current) {
    throw new Error("pgmem: no fork for this test file: add @pgmem/core/register to the setup files (Jest: testEnvironment @pgmem/core/jest-environment)");
  }
  return current;
}

/** Reset the file fork after a sequential test; composes with other callback wrappers. */
function withTestReset(fn, options) {
  if (typeof fn !== "function") throw new TypeError("pgmem: withTestReset expects a test callback");
  return async function (...args) {
    const fork = currentFork();
    const active = globalThis[Symbol.for("pgmem.resetActiveForks")] ??= new WeakSet();
    if (active.has(fork)) throw new Error("pgmem: overlapping withTestReset callbacks share a fork; use withFork for concurrent tests");
    active.add(fork);
    try {
      return await runWithCleanup(() => fn.apply(this, args), () => fork.reset(options), "pgmem: test and reset both failed");
    } finally {
      active.delete(fork);
    }
  };
}

let sharedTestFork;
let testRouteTail = Promise.resolve();

/** Route a test callback to prepared state; constructs clients inside fn. */
function withTestDatabase(fn, { fork: forkEnabled = true, env } = {}) {
  if (typeof forkEnabled !== "boolean") throw new TypeError("pgmem: fork must be a boolean");
  const run = async () => {
    const target = forkEnabled ? await fork() : await (sharedTestFork ??= fork().catch((err) => {
      sharedTestFork = undefined;
      throw err;
    }));
    const names = env ?? process.env.PGMEM_ENV?.split(",").map((s) => s.trim()).filter(Boolean);
    const values = target.env(names?.length ? names : undefined);
    const previous = Object.fromEntries(Object.keys(values).map((name) => [name, process.env[name]]));
    Object.assign(process.env, values);
    try {
      return await fn(target);
    } finally {
      for (const [name, value] of Object.entries(previous)) {
        if (value === undefined) delete process.env[name];
        else process.env[name] = value;
      }
      if (forkEnabled) await target.close();
    }
  };
  // process.env is process-wide; serialize callbacks that change it.
  const result = testRouteTail.then(run);
  testRouteTail = result.catch(() => {});
  return result;
}

/** Register a node:test, Vitest or Jest case with pgmem routing. */
function pgmemTest(test, name, fn, options) {
  return test(name, (...args) => withTestDatabase(() => fn(...args), options));
}

const patchedPg = new WeakSet();
let shadowInstalled = false;
let sharedShadowFork;

/** Redirect new pg connections without changing the application's connection string. */
function patchPg(pg) {
  if (!pg || typeof pg.Client !== "function" || typeof pg.Pool !== "function" || patchedPg.has(pg)) return;
  patchedPg.add(pg);
  const clientConnect = pg.Client.prototype.connect;
  pg.Client.prototype.connect = function (...args) {
    const target = shadowContext.getStore();
    if (target) {
      const endpoint = new URL(target.url);
      const params = this.connectionParameters;
      if (!params) throw new Error("pgmem: unsupported pg Client without connectionParameters");
      params.host = this.host = endpoint.hostname;
      params.port = this.port = Number(endpoint.port);
      params.user = this.user = decodeURIComponent(endpoint.username);
      params.database = this.database = decodeURIComponent(endpoint.pathname.slice(1));
      params.password = this.password = decodeURIComponent(endpoint.password);
      params.ssl = this.ssl = false;
      if (this.connection) this.connection.ssl = false;
      this.__pgmemShadowTarget = target.id;
    }
    return clientConnect.apply(this, args);
  };

  const poolConnect = pg.Pool.prototype.connect;
  pg.Pool.prototype.connect = function (...args) {
    const target = shadowContext.getStore();
    if (target && Array.isArray(this._clients) && Array.isArray(this._idle)) {
      (target.__pgmemShadowPools ??= new Set()).add(this);
      for (const idle of [...this._idle]) {
        if (idle.client.__pgmemShadowTarget !== target.id) this._remove(idle.client);
      }
      const idleClients = new Set(this._idle.map((item) => item.client));
      if (this._clients.some((client) => !idleClients.has(client) && client.__pgmemShadowTarget !== target.id)) {
        throw new Error("pgmem: a pg Pool still has a checked-out connection from another target");
      }
    }
    return poolConnect.apply(this, args);
  };
}

/** Install a passive pg hook; outside withShadowPg, ordinary connections are unchanged. */
function installPgShadow() {
  if (shadowInstalled) return;
  shadowInstalled = true;
  // ESM imports of a CommonJS package can bypass Module._load (notably under
  // Vitest), so patch the application's resolved pg export eagerly as well.
  try { patchPg(createRequire(join(process.cwd(), "package.json"))("pg")); }
  catch (error) { if (error.code !== "MODULE_NOT_FOUND") throw error; }
  try { patchPg(require("pg")); }
  catch (error) { if (error.code !== "MODULE_NOT_FOUND") throw error; }
  for (const loaded of Object.values(require.cache)) patchPg(loaded?.exports);
  const load = Module._load;
  Module._load = function (request, parent, isMain) {
    const exported = load.call(this, request, parent, isMain);
    if (request === "pg") patchPg(exported);
    return exported;
  };
}

/** Run a test against a fork while pg Clients and Pools retain their normal configuration. */
async function withShadowPg(fn, { fork: forkEnabled = true } = {}) {
  if (typeof fn !== "function") throw new TypeError("pgmem: withShadowPg requires a function");
  if (typeof forkEnabled !== "boolean") throw new TypeError("pgmem: fork must be a boolean");
  installPgShadow();
  const target = forkEnabled ? await fork() : await (sharedShadowFork ??= fork().catch((err) => {
    sharedShadowFork = undefined;
    throw err;
  }));
  try {
    return await shadowContext.run(target, fn);
  } finally {
    if (forkEnabled) {
      // Drain idle connections before closing the fork: otherwise pg reports
      // the server shutdown as an asynchronous error on an application pool.
      for (const pool of target.__pgmemShadowPools ?? []) {
        await Promise.all([...pool._idle]
          .filter((item) => item.client.__pgmemShadowTarget === target.id)
          .map((item) => new Promise((resolve) => pool._remove(item.client, resolve))));
      }
      await target.close();
    }
  }
}

/** Register a node:test, Vitest or Jest case using pg driver interception. */
function shadowPg(test, name, fn, options) {
  return test(name, (...args) => withShadowPg(() => fn(...args), options));
}

/** Bind a loopback socket to obtain an available port for an app child. */
function freePort() {
  return new Promise((resolve, reject) => {
    const listener = net.createServer();
    listener.once("error", reject);
    listener.listen(0, "127.0.0.1", () => {
      const port = listener.address().port;
      listener.close((err) => err ? reject(err) : resolve(port));
    });
  });
}

/** Launch an HTTP app against a fork and wait until it is ready. */
async function startTestApp({ fork: target, command = process.execPath, args = [], cwd, env = {}, healthPath = "/health", timeoutMs = 10000, portEnv = "PORT", databaseEnv = "DATABASE_URL", stdio = "inherit" }) {
  if (!target?.url) throw new TypeError("pgmem: startTestApp requires a fork");
  if (!Array.isArray(args)) throw new TypeError("pgmem: startTestApp args must be an array");
  const port = await freePort();
  const url = `http://127.0.0.1:${port}`;
  const child = spawn(command, args, {
    cwd,
    env: { ...process.env, ...env, [databaseEnv]: target.url, [portEnv]: String(port) },
    stdio,
  });
  let childError;
  child.once("error", (err) => { childError = err; });
  const stopped = new Promise((resolve) => child.once("close", resolve));
  const close = async () => {
    if (child.exitCode === null && child.signalCode === null && !childError) {
      child.kill("SIGTERM");
      let timer;
      const result = await Promise.race([
        stopped.then(() => "closed"),
        new Promise((resolve) => { timer = setTimeout(() => resolve("timeout"), 5000); }),
      ]);
      clearTimeout(timer);
      if (result === "timeout") child.kill("SIGKILL");
    }
    await stopped;
  };
  try {
    const healthUrl = new URL(healthPath, url);
    const deadline = Date.now() + timeoutMs;
    while (Date.now() < deadline) {
      if (childError) throw childError;
      if (child.exitCode !== null || child.signalCode !== null) throw new Error(`pgmem: app exited before ${healthUrl} was ready`);
      try {
        const response = await fetch(healthUrl, { signal: AbortSignal.timeout(Math.min(1000, Math.max(1, deadline - Date.now()))) });
        if (response.ok) return { url, port, process: child, close };
      } catch { /* The app can still be starting. */ }
      await new Promise((resolve) => setTimeout(resolve, 100));
    }
    throw new Error(`pgmem: app did not become ready at ${healthUrl} within ${timeoutMs} ms`);
  } catch (err) {
    await close();
    throw err;
  }
}

module.exports = {
  PROTOCOL,
  PgmemError,
  PgmemServer,
  PgmemClient,
  PgmemSnapshot,
  PgmemFork,
  connect,
  fork,
  withFork,
  useFork,
  currentFork,
  withTestReset,

  withTestDatabase,
  pgmemTest,
  withShadowPg,
  shadowPg,
  startTestApp,
  resolveBinary,
};
