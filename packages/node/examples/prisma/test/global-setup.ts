import { execFileSync } from "node:child_process";
import { PgmemServer } from "@pgmem/core";

// Once per run: start pgmem, apply the migrations to the template database
// and snapshot it. Test files fork the snapshot through PGMEM_CONTROL.
export async function setup() {
  const pg = await PgmemServer.start({
    database: "app",
    prepare({ url }) {
      execFileSync("npx", ["prisma", "migrate", "deploy"], {
        stdio: "inherit",
        env: { ...process.env, DATABASE_URL: url },
        shell: process.platform === "win32",
      });
    },
  });
  Object.assign(process.env, pg.env());
  return () => pg.close();
}
