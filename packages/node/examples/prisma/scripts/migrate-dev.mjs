// `npm run migrate:dev -- --name add_something` authors a migration without
// a PostgreSQL install: pgmem applies the existing migrations to an empty
// in-memory database, and Prisma creates its shadow database on the same
// server to diff the schema against them.
import { spawnSync } from "node:child_process";
import { PgmemServer } from "@pgmem/core";

const server = await PgmemServer.start({ database: "app", control: false });
try {
  const run = spawnSync("npx", ["prisma", "migrate", "dev", ...process.argv.slice(2)], {
    stdio: "inherit",
    env: { ...process.env, DATABASE_URL: server.url },
    shell: process.platform === "win32",
  });
  process.exitCode = run.status ?? 1;
} finally {
  await server.close();
}
