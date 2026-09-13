import { defineConfig } from "prisma/config";

export default defineConfig({
  schema: "prisma/schema.prisma",
  migrations: { path: "prisma/migrations" },
  // pgmem hands the URL over through the environment; generate needs none
  datasource: { url: process.env.DATABASE_URL ?? "" },
});
