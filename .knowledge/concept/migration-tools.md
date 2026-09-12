---
id: concept:migration-tools
type: concept
title: Migration Tools on pgmem
---
Any migration tool that accepts a DSN or JDBC URL runs unchanged against the template server before the snapshot; this lists the call for each popular tool.

```yaml
summary:
  principle: migrate the template once (flow:test-lifecycle prepare step), then api:snapshot; forks inherit the migrated schema; commit or close every connection before snapshot or it reports busy
  raw_sql: read the schema file and execute it on the template DSN with the normal driver; multi-statement strings work; COPY FROM STDIN works
  go:
    golang_migrate: 'm, _ := migrate.New("file://migrations", dsn); m.Up()  # inside Options.Prepare, dsn argument'
    goose: 'goose.SetDialect("postgres"); goose.Up(db, "migrations")  # db is the Prepare *sql.DB'
    atlas: 'exec atlas migrate apply --url dsn from Prepare, or the atlas Go SDK'
    tern: 'tern migrate --conn-string dsn or the tern migrate package'
    orm_auto: 'GORM db.AutoMigrate, ent client.Schema.Create(ctx), sqlc has no migrations'
  python:
    alembic: 'cfg = Config("alembic.ini"); cfg.set_main_option("sqlalchemy.url", dsn); command.upgrade(cfg, "head")  # inside the overridden pgmem_snapshot fixture'
    sqlalchemy: 'Base.metadata.create_all(create_engine(dsn))'
    yoyo: 'backend = get_backend(dsn); backend.apply_migrations(backend.to_apply(read_migrations("migrations")))'
    django: 'status to verify: the test runner creates test_<name> with CREATE DATABASE and connects to it, which pgmem does not serve yet (concept:limits); until then call_command("migrate") with DATABASES pointing at the template and run tests against pgmem_dsn with --reuse-db or a custom runner'
  java:
    flyway: 'Flyway.configure().dataSource(url, user, null).load().migrate()  # inside PgmemExtension.builder().prepare(t -> ...)'
    liquibase: 'new CommandScope("update").addArgumentValue("url", url).addArgumentValue("username", user).addArgumentValue("changelogFile", "db/changelog.xml").execute()'
    hibernate: 'hibernate.hbm2ddl.auto=create against the template, or Flyway under Spring Boot once auto-configuration exists (api:java-wrapper later)'
  nodejs: 'planned (requirement:nodejs-wrapper): Prisma migrate deploy with DATABASE_URL (shadow database needs multi-database serving), drizzle-kit migrate, TypeORM runMigrations, knex.migrate.latest'
  gotchas:
    - tools that open a second connection inside a transaction (advisory-lock based locking in golang-migrate and Flyway) work because the lock and the migration run on the same connection; a tool that waits on another connection would hang (rule:single-session-per-backend)
    - CREATE EXTENSION works for the bundled set only (policy:bundled-extensions); pgvector is available as vector
    - migration history tables are part of the snapshot, so forks report the schema as up to date
```
