import pg from 'pg';
import { PgmemServer } from '@pgmem/core';
import { installServiceEnv } from '@pgmem/core/environment';

let server;
let release;
export async function globalSetup() {
  server = await PgmemServer.start({
    async prepare({ url }) {
      const pool = new pg.Pool({ connectionString: url });
      try { await pool.query('CREATE TABLE orders(id int PRIMARY KEY)'); }
      finally { await pool.end(); }
    },
  });
  try { release = installServiceEnv('pgmem:control', server.env()); }
  catch (error) {
    try { await server.close(); }
    catch (cleanup) { throw new AggregateError([error, cleanup], 'setup and cleanup failed'); }
    throw error;
  }
}
export async function globalTeardown() {
  const failures = [];
  try { release?.(); } catch (error) { failures.push(error); }
  try { await server?.close(); } catch (error) { failures.push(error); }
  if (failures.length) throw new AggregateError(failures, 'global cleanup failed');
}
