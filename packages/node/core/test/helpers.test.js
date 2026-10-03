import { test } from 'node:test';
import assert from 'node:assert/strict';
import { createRequire } from 'node:module';
import { withTestReset } from '../index.js';
import { installServiceEnv, ServiceEnvConflict } from '../environment.cjs';

test('service environment ownership detects collisions before changing any variable', () => {
  const scope = {};
  const target = { DATABASE_URL: 'existing' };
  const release = installServiceEnv('pgmem', { DATABASE_URL: 'test', APP_DB: 'test' }, { scope, target });
  assert.throws(() => installServiceEnv('valkey', { VALKEY_URL: 'cache', DATABASE_URL: 'wrong' }, { scope, target }),
    (e) => e instanceof ServiceEnvConflict && e.owner === 'pgmem' && e.service === 'valkey' && e.variable === 'DATABASE_URL');
  assert.equal(target.VALKEY_URL, undefined);
  const releaseCache = installServiceEnv('valkey', { VALKEY_URL: 'cache' }, { scope, target });
  const renewed = installServiceEnv('pgmem', { DATABASE_URL: 'next' }, { scope, target });
  release();
  assert.equal(target.DATABASE_URL, 'next');
  renewed(); renewed(); releaseCache();
  assert.deepEqual(target, { DATABASE_URL: 'existing' });
});

test('environment installation rolls back earlier writes when assignment fails', () => {
  const values = { A: 'previous' };
  const target = new Proxy(values, { set(object, name, value) {
    if (name === 'B' && value === 'fail') throw new Error('assignment failed');
    object[name] = value; return true;
  } });
  const scope = {};
  assert.throws(() => installServiceEnv('pgmem', { A: 'next', B: 'fail' }, { target, scope }), /assignment failed/);
  assert.deepEqual(values, { A: 'previous' });
  const release = installServiceEnv('other', { A: 'other' }, { target, scope });
  release();
  assert.equal(values.A, 'previous');
});

test('Jest factory separates acquisition and reset deadlines and restores environment', async () => {
  const require = createRequire(import.meta.url);
  const pgmem = require('../index.cjs');
  const { withPgmemEnvironment } = require('../jest-environment-factory.cjs');
  const original = pgmem.fork;
  const events = [];
  pgmem.fork = async (options) => {
    events.push(['fork', options.timeoutMs]);
    return { env: () => ({ DATABASE_URL: 'fork-url' }), reset: async (options) => events.push(['reset', options.timeoutMs]), close: async () => {} };
  };
  class Base {
    global = { process: { env: { DATABASE_URL: 'previous' } } };
    async setup() {}
    async teardown() { events.push(['base cleanup']); }
  }
  try {
    const Environment = withPgmemEnvironment(Base, { forkTimeoutMs: 123, resetTimeoutMs: 456 });
    const first = new Environment();
    await first.setup(); await first.teardown();
    assert.equal(first.global.process.env.DATABASE_URL, 'previous');
    const second = new Environment();
    await second.setup(); await second.teardown();
    assert.deepEqual(events.slice(0, 3), [['fork', 123], ['base cleanup'], ['reset', 456]]);
  } finally {
    class Failure extends Base { async setup() { throw new Error('clear cached fork'); } }
    await assert.rejects(new (withPgmemEnvironment(Failure))().setup());
    pgmem.fork = original;
  }
});

test('reset options, failure evidence and overlapping callbacks', async () => {
  const key = Symbol.for('pgmem.currentFork');
  const old = globalThis[key];
  const resetError = new Error('reset failed');
  const testError = new Error('assertion failed');
  const options = { snapshot: 'seeded', timeoutMs: 123 };
  try {
    globalThis[key] = { reset: async (got) => { assert.equal(got, options); throw resetError; } };
    await assert.rejects(withTestReset(async () => { throw testError; }, options)(),
      (err) => err instanceof AggregateError && err.errors[0] === testError && err.errors[1] === resetError);
    let resume;
    globalThis[key] = { reset: async () => {} };
    const waiting = withTestReset(() => new Promise((resolve) => { resume = resolve; }))();
    await assert.rejects(withTestReset(async () => {})(), /overlapping/);
    resume(); await waiting;
    assert.equal(await withTestReset(async (n) => n + 1)(2), 3);
  } finally {
    if (old === undefined) delete globalThis[key]; else globalThis[key] = old;
  }
});

test('environment factory cleans the base environment when setup fails', async () => {
  const { withPgmemEnvironment } = createRequire(import.meta.url)('../jest-environment-factory.cjs');
  let cleaned = 0;
  class Base {
    async setup() { throw new Error('other service failed'); }
    async teardown() { cleaned++; }
  }
  const Environment = withPgmemEnvironment(Base);
  await assert.rejects(new Environment().setup(), /other service failed/);
  assert.equal(cleaned, 1);
});

test('environment release attempts every variable after a restore failure', () => {
  const values = { A: 'old-a', B: 'old-b' };
  let failRestore = false;
  const target = new Proxy(values, { set(object, name, value) {
    if (failRestore && name === 'A') throw new Error('restore failed');
    object[name] = value; return true;
  } });
  const release = installServiceEnv('pgmem', { A: 'next-a', B: 'next-b' }, { target, scope: {} });
  failRestore = true;
  assert.throws(release, (error) => error instanceof AggregateError && error.errors[0].message === 'restore failed');
  assert.equal(values.B, 'old-b');
});

test('a claimed Jest endpoint closes the acquired fork and cleans the base', async () => {
  const require = createRequire(import.meta.url);
  const pgmem = require('../index.cjs');
  const { withPgmemEnvironment } = require('../jest-environment-factory.cjs');
  const original = pgmem.fork;
  let closed = 0;
  let cleaned = 0;
  let requestedTimeout;
  pgmem.fork = async (options) => {
    requestedTimeout = options.timeoutMs;
    return { env: () => ({ DATABASE_URL: 'fork-url' }), close: async () => { closed++; } };
  };
  class Base {
    global = { process: { env: {} } };
    async setup() {
      this.release = installServiceEnv('valkey', { DATABASE_URL: 'wrong-name' }, {
        target: this.global.process.env, scope: this.global,
      });
    }
    async teardown() { this.release(); cleaned++; }
  }
  try {
    const env = new (withPgmemEnvironment(Base, { forkTimeoutMs: 30000 }))();
    await assert.rejects(env.setup(), (error) => error instanceof ServiceEnvConflict && error.owner === 'valkey');
    assert.equal(closed, 1);
    assert.equal(cleaned, 1);
    assert.equal(requestedTimeout, 30000);
    assert.deepEqual(env.global.process.env, {});
  } finally { pgmem.fork = original; }
});
