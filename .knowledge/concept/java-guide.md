---
id: concept:java-guide
type: concept
title: Java Guide
---
Using pgmem from JVM tests: dependencies with the platform classifier, the JUnit 5 extension, Flyway or Liquibase in prepare, and fork scopes.

```yaml
summary:
  install: |
    testImplementation("jp.shibu:pgmem-junit5:0.1.0")
    testRuntimeOnly("jp.shibu:pgmem-native:0.1.0:darwin-arm64")   // classifier per platform, policy:binary-distribution
    testRuntimeOnly("org.postgresql:postgresql:42.7.7")
  basic_test: |
    try (Pgmem pg = Pgmem.builder().database("app").start();
         Connection c = DriverManager.getConnection(pg.template().jdbcUrl())) {
        c.createStatement().execute("select 1");
    }
  initialization:
    extension: '@RegisterExtension static PgmemExtension pg = PgmemExtension.builder().database("app").prepare(t -> migrate(t.jdbcUrl())).build()  # prepare runs once, then snapshot (api:java-wrapper)'
    raw_sql: Statement.execute of the schema file contents inside prepare
    tools: Flyway, Liquibase, Hibernate hbm2ddl (concept:migration-tools); seed with DbUnit or SQL (concept:seeding)
  test_case_styles:
    fork_per_test: 'Fork, DataSource or @PgmemFork String jdbcUrl parameters; closed after each test'
    fork_per_class: '@PgmemFork(scope = ForkScope.CLASS) or builder forkScope(CLASS) for read-only classes'
    no_fork: pg.template().jdbcUrl() with a transaction rolled back per test; sequential
    named_templates: 'builder().template("audit", t -> ...) and @PgmemFork("audit")'
    guidance: decision:fork-or-not
  parallel: junit.jupiter.execution.parallel works; maxForks bounds live forks and forkTimeout turns a full pool into a failure (policy:fork-pool-limit); Gradle maxParallelForks JVMs each start their own binary (concept:server-process)
  pools: HikariCP on a fork works; connections serialize at transaction boundaries (rule:single-session-per-backend); close connections before snapshot or it fails with busy after 30 s
  spring_boot: auto-configuration is planned; today set spring.datasource.url from a fork's jdbcUrl in a DynamicPropertySource
  limits: concept:limits
```
