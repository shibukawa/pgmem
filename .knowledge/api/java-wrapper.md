---
id: api:java-wrapper
type: api
title: Java Wrapper
---
Maven artifacts in packages/java: a core client for api:control-protocol and a JUnit 5 extension that injects a fresh fork per test.

```yaml
api:
  package: packages/java (Gradle multi-module, wrapper 9.5.1, options.release 17); java >= 17; core has no third-party runtime dependency (Json.java is a minimal codec)
  status: implemented 2026-09-12; tests use pgjdbc against the binary that the root buildBinary task compiles (system property pgmem.binary)
  group_id: io.github.shibukawa.pgmem  # jp.shibu until 2026-09-14; the namespace io.github.shibukawa is verified through the GitHub account, as for osmem
  artifacts:
    pgmem: core client
    pgmem-junit5: JUnit 5 extension
    pgmem-spring: '@ShadowPg composite SpringBootTest annotation, context customizer, and early TestExecutionListener reset; Spring Boot supplied by the application'
    pgmem-micronaut: MicronautTestApp helper for an HTTP test server on a pgmem application fork; Micronaut runtime is compileOnly
    pgmem-native: binary per classifier, built by the root buildBinary Exec task (go build, -Pgoos/-Pgoarch cross-compile, PGMEM_BINARY copies a prebuilt one) (policy:binary-distribution)
  core:
    - 'Pgmem.builder().database("app").user(..).param(k, v).log(bool).binary(path).start() -> Pgmem, AutoCloseable'
    - 'Pgmem: template() -> Server (default); startServer(ServerOptions) -> Server (op start); close() shuts the process down'
    - 'Server: id(), jdbcUrl(), dsn(), user(), host(), port(), dataSource(), snapshot(), snapshot(maxForks), snapshot(maxForks, Duration timeout) -> Snapshot; close()'
    - 'Server.reset() or reset(snapshot, Duration timeout): restore in place; default reset waits 5s; explicit snapshots must belong to the same process'
    - 'Snapshot.fork(Duration timeout) -> pool_timeout instead of waiting forever'
    - 'BinaryLocator: pgmem.binary property, PGMEM_BINARY, classpath resource /io/github/shibukawa/pgmem/native/<classifier>/pgmem extracted to PGMEM_CACHE_DIR or ~/.cache/pgmem/<Implementation-Version>/, then PATH'
    - 'Snapshot: fork() -> Fork, AutoCloseable; close()'
    - 'Fork: jdbcUrl(), dataSource(), reset(), reset(Snapshot), reset(Snapshot, Duration timeout=5s), close(); reset keeps JDBC URL and open connections valid'
    - 'dataSource(): javax.sql.DataSource over DriverManager so core does not compile against pgjdbc'
  junit5:
    guard: builder sharedReadOnly(true) guards CLASS-scoped and declarative fork=false TestDatabase/DataSource/JDBC URL injection; raw Fork connection access stays explicit; core dsn(boolean)/jdbcUrl(boolean)/dataSource(boolean) expose guard views
    cleanup: attempts all fork and process closes; suppressed exceptions retain original test plus cleanup errors; EOF bounds process shutdown
    registration: '@RegisterExtension static PgmemExtension pg = PgmemExtension.builder().prepare(url -> migrate(url)).build()'
    several_templates: 'builder().database("app").prepare(t -> migrate(t.jdbcUrl())).template("audit", t -> migrateAudit(t.jdbcUrl())); the default template comes first unless only named ones are configured; @PgmemFork("audit") selects, unnamed uses the first'
    lifecycle: BeforeAll start, prepare, snapshot; fork per scope; close at scope end (decision:fork-release-detection); AfterAll shutdown
    fork_scope: METHOD default (fresh fork per test); CLASS via builder forkScope(CLASS) or @PgmemFork(scope = CLASS) so read-only tests share one fork; closed in afterEach / afterAll
    store_gotcha: ExtensionContext.Store lookups fall through to parent contexts, so fork sets are keyed by the scope context's unique id, never by a shared class key
    registration_rule: static @RegisterExtension (or @ExtendWith for defaults); nested classes reuse the outer process and only the owner class's afterAll stops it
    injection: ParameterResolver for Fork, DataSource, TestDatabase and String jdbcUrl; TestDatabase offers fork(), dataSource(), jdbcUrl(), dsn(), reset() for one automatically owned fork; pg.fork() for manual use
    declarative: '@PgmemTest(fork=true) method or class annotation selects fresh method fork by default; fork=false shares a prepared read-only fork for unannotated parameters'
    application: 'applicationFork([templateName]) lazily starts prepared templates even before BeforeAll (for Spring DynamicPropertySource), returns a stable fork that can reset before sequential HTTP/browser cases, closes after owner class'
    parallel: safe with junit.jupiter.execution.parallel; forks are independent (policy:fork-pool-limit, default from the server's memory budget)
    forked_jvms: each test JVM (Gradle maxParallelForks, surefire forkCount) spawns its own concept:server-process
  micronaut:
    helper: 'MicronautTestApp.start(Fork[, Map<String, ?> overrides]) starts ApplicationContext.run(EmbeddedServer.class, properties, "test") with a random port, fork JDBC URL and user, PostgreSQL driver, and Flyway disabled; uri(), server(), close()'
    lifecycle: 'Start once after applicationFork(); reset the same fork before each sequential API or browser test; close the app before the extension closes its fork'
    shadow_pg: '@ShadowPg(schema=classpath SQL, seed=classpath SQL) JUnit annotation creates a pgmem fork, starts the app with its JDBC URL, injects the test instance, captures post-startup baseline, resets before each method, closes server before database'
  spring:
    shadow_pg: '@ShadowPg(schema=classpath SQL, seed=classpath SQL) meta SpringBootTest annotation with ContextCustomizerFactory; injects JDBC properties before DataSource creation, captures post-startup baseline, resets via early TestExecutionListener before test transaction starts, closes database on Spring context close; @DirtiesContext(AFTER_CLASS) prevents reuse'
  shadow_pg_shared:
    runtime: 'ShadowPgDatabase starts template and fork, runs schema and seed before snapshot; if schema omitted, seed is deferred until after app startup migrations; captures an application baseline and resets a stable fork without changing pooled connections'
  threads: one daemon reader thread; CompletableFuture per request id (rule:non-blocking-control-channel)
  later: Spring Boot test auto-configuration (zonky style), Testcontainers JdbcDatabaseContainer adapter
```
