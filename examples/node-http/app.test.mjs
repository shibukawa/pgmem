import { test, after, before } from 'node:test';
import assert from 'node:assert/strict';
import { once } from 'node:events';
import pg from 'pg';
import { withTestReset } from '@pgmem/core';
import { createApp, pool } from './app.mjs';

const app = createApp();
const assertions = new pg.Pool({ connectionString: process.env.DATABASE_URL });
before(async () => { app.listen(0, '127.0.0.1'); await once(app, 'listening'); });
after(async () => {
  const failures = [];
  try { await new Promise((resolve, reject) => app.close((error) => error ? reject(error) : resolve())); }
  catch (error) { failures.push(error); }
  const results = await Promise.allSettled([pool.end(), assertions.end()]);
  failures.push(...results.filter((r) => r.status === 'rejected').map((r) => r.reason));
  if (failures.length) throw new AggregateError(failures, 'application cleanup failed');
});

test('the application and assertions use the same fork', withTestReset(async () => {
  const response = await fetch(`http://127.0.0.1:${app.address().port}/orders`, {
    method: 'POST', headers: { 'content-type': 'application/json' }, body: JSON.stringify({ id: 42 }),
  });
  await response.arrayBuffer();
  assert.equal(response.status, 201);
  assert.deepEqual((await assertions.query('SELECT id FROM orders')).rows, [{ id: 42 }]);
}));

test('the same application client sees the restored baseline', async () => {
  assert.equal((await pool.query('SELECT count(*)::int AS n FROM orders')).rows[0].n, 0);
});
