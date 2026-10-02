# Java unit tests with JUnit 5

For an application-managed repository, use the framework's one-annotation integration: [Spring Boot](java-spring.md) or [Micronaut](java-micronaut.md). `@ShadowPg(schema = "db/schema.sql")` creates the app context against pgmem and restores its prepared data before each method. Omit `schema` if the application migrates at startup; `seed = "db/seed.sql"` adds initial rows. Retrieve the normal repository bean and assert through its methods.

For a plain repository, register `static PgmemExtension pg = PgmemExtension.builder().database("app").prepare(t -> migrateAndSeed(t.jdbcUrl())).build()` with `@RegisterExtension`. Inject a method-scoped `DataSource` parameter into tests and build the repository from it. `@PgmemTest(fork = false)` shares a prepared fork only for read-only tests. A pool made before injection cannot be redirected. Close preparation transactions before snapshot.

Method-scoped forks with method-owned repositories can run in parallel. `@ShadowPg` apps and class-scoped writable forks must run methods sequentially; parallel classes need distinct contexts, forks, and ports. `maxForks(n)` bounds snapshot fanout; `forkTimeout(Duration)` bounds waiting for a slot.

Full Java/Kotlin examples: `website/src/content/docs/guides/java/testing.mdx`.
