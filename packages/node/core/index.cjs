"use strict";
// @pgmem/core: starts the pgmem binary (a real PostgreSQL 18 that keeps
// everything in memory) and drives its control protocol. The
// implementation is CommonJS so that Jest can load it inside test files;
// index.js re-exports it for ESM.
const { spawn } = require("node:child_process");
const { existsSync } = require("node:fs");
const net = require("node:net");
const { join } = require("node:path");
const { createInterface } = require("node:readline");

/** Control protocol version this package speaks; the binary must match. */
const PROTOCOL = 1;
const asyncDispose = Symbol.asyncDispose ?? Symbol.for("Symbol.asyncDispose");
const CURRENT = Symbol.for("pgmem.currentFork");

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
  env(names = ["DATABASE_URL"]) {
    const out = {};
    for (const name of [].concat(names)) out[name] = envParts[name]?.(this) ?? this.url;
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
    try {
      return await fn(fork);
    } finally {
      await fork.close();
    }
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
    await this.#channel.request("close", { server: this.id }).catch(() => {});
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
      await server.close();
      throw err;
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
    await this.#channel.request("shutdown").catch(() => {});
    child.stdin.end();
    const killer = setTimeout(() => child.kill("SIGKILL"), 10000);
    await this.#exited;
    clearTimeout(killer);
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
async function useFork({ env } = {}) {
  const names = env ?? process.env.PGMEM_ENV?.split(",").map((s) => s.trim()).filter(Boolean);
  let current = globalThis[CURRENT];
  if (current) await current.reset();
  else {
    current = await fork();
    globalThis[CURRENT] = current;
  }
  Object.assign(process.env, current.env(names?.length ? names : undefined));
  return current;
}

/** The fork of this test file, made by @pgmem/core/register or @pgmem/core/jest-environment. */
function currentFork() {
  const current = globalThis[CURRENT];
  if (!current) {
    throw new Error("pgmem: no fork for this test file: add @pgmem/core/register to the setup files (Jest: testEnvironment @pgmem/core/jest-environment)");
  }
  return current;
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
  resolveBinary,
};
