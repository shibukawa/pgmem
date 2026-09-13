# Prisma 7 with @pgmem/core

A small Prisma 7 app tested with Vitest against pgmem. Nothing in `src/`
knows about pgmem: the Prisma client reads `DATABASE_URL` when `src/db.ts`
is imported.

- `test/global-setup.ts` starts pgmem once, runs `prisma migrate deploy`
  against the template database and snapshots it.
- `@pgmem/core/register` (a setup file) gives every test file its own copy
  and writes its URL to `DATABASE_URL` before the file is imported, so the
  two test files run in parallel on separate databases.
- `test/reset.ts` puts the copy back before every test without
  disconnecting the Prisma client.
- `npm run migrate:dev -- --name <name>` runs `prisma migrate dev` against
  a throwaway pgmem database; the migrations in `prisma/migrations` were
  created that way. Prisma's shadow database lives on the same server.

```sh
npm install
npm test
```

Until the `@pgmem/<platform>` binary packages are published, point
`PGMEM_BINARY` at a build of `cmd/pgmem` (`go build -o pgmem ./cmd/pgmem` in
the repository root).
