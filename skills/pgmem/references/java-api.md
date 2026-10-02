# Java HTTP API tests

For an application-managed HTTP server, choose [Spring Boot](java-spring.md) or [Micronaut](java-micronaut.md). Annotate the JUnit 5 class with framework `@ShadowPg(schema = "db/schema.sql")`; it starts an application on a random port with its normal DataSource routed to pgmem. The app and pool remain alive through the class. Each test method starts from the prepared baseline after reset. Send HTTP requests to the injected port/`EmbeddedServer` URI.

For an application accepting a `DataSource`, build it with `pg.applicationFork().dataSource()` before serving requests. For a child JVM, pass `pg.applicationFork().jdbcUrl()` before its pool starts. A suite-specific baseline can use `Snapshot seeded = pg.applicationFork().snapshot()` after seeding and `pg.applicationFork().reset(seeded)` before each test; close `seeded` afterward.

Methods sharing one app/fork must be sequential. Parallel classes need separate application contexts, forks, ports, and no static shared `DataSource`. Complete requests, jobs, and transactions before reset. Full guide: `website/src/content/docs/guides/java/api-testing.mdx`.
