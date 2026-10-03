import { createServer } from 'node:http';
import pg from 'pg';

// register must inject DATABASE_URL before this module is imported.
export const pool = new pg.Pool({ connectionString: process.env.DATABASE_URL });
export function createApp() {
  return createServer(async (request, response) => {
    if (request.method !== 'POST' || request.url !== '/orders') {
      response.writeHead(404).end();
      return;
    }
    try {
      let body = '';
      for await (const chunk of request) body += chunk;
      const { id } = JSON.parse(body);
      if (!Number.isSafeInteger(id) || id <= 0) { response.writeHead(400).end(); return; }
      await pool.query('INSERT INTO orders(id) VALUES ($1)', [id]);
      response.writeHead(201).end();
    } catch { response.writeHead(500).end(); }
  });
}
