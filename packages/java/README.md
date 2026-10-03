# pgmem for Java

A real PostgreSQL that runs entirely in memory, packaged as a single
binary and driven from the JVM. Prepare the schema once, then give every
test its own fork of that state in about 20 ms.

```java
try (Pgmem pg = Pgmem.builder().database("app").start()) {
    migrate(pg.template().jdbcUrl());
    Snapshot snap = pg.template().snapshot();
    try (Fork fork = snap.fork()) {                 // private copy of the prepared database
        DriverManager.getConnection(fork.jdbcUrl()) ...
    }
}
```

## Dependencies

```kotlin
testImplementation("io.github.shibukawa.pgmem:pgmem-junit5:1.18.0")           // pulls in io.github.shibukawa.pgmem:pgmem
testRuntimeOnly("io.github.shibukawa.pgmem:pgmem-native:1.18.0:darwin-arm64") // the binary for your platform
testRuntimeOnly("org.postgresql:postgresql:42.7.7")                          // your usual driver
```

Classifiers: `linux-x86_64`, `linux-arm64`, `darwin-arm64`, `windows-x86_64`,
`windows-arm64`. Pick one with `os-maven-plugin` or Gradle's
`osdetector` if the build runs on several platforms. Without the native
artifact, `PGMEM_BINARY` (or the `pgmem.binary` system property) points at
a local build of `github.com/shibukawa/pgmem/cmd/pgmem`. The binary is
extracted once into `~/.cache/pgmem/<version>/`.

## JUnit 5

```java
@RegisterExtension
static PgmemExtension pg = PgmemExtension.builder()
        .database("app")
        .prepare(t -> migrate(t.jdbcUrl()))        // runs once, then snapshot
        .template("audit", t -> migrateAudit(t.jdbcUrl()))
        .build();

@Test void orders(Fork db) { ... }                                  // fresh copy per test
@Test @PgmemTest(fork = false) void report(DataSource ds) { ... }  // shared read-only copy
@Test void classReport(@PgmemFork(scope = ForkScope.CLASS) DataSource ds) { ... } // shared by the class
@Test void audit(@PgmemFork("audit") String jdbcUrl) { ... }        // a named template
```

`Fork`, `DataSource` and `@PgmemFork String` parameters are resolved; several
parameters in one test share one fork. Forks are closed after each test
(or after the class for `ForkScope.CLASS`). Tests may run in parallel;
`maxForks` bounds how many forks exist at once and `forkTimeout` turns a
full pool into a test failure instead of a wait.

For an API or browser test, call `pg.applicationFork()` before starting the
HTTP application and give its JDBC URL to the app's pool. The extension may
start early when Spring's `@DynamicPropertySource` requests that URL. Call
`pg.applicationFork().reset()` before each sequential case: the JDBC URL
and existing connections remain valid. A chosen dataset can be restored with
`fork.reset(snapshot)`. See the [unit](https://shibukawa.github.io/pgmem/guides/java/testing/),
[API](https://shibukawa.github.io/pgmem/guides/java/api-testing/), and
[E2E](https://shibukawa.github.io/pgmem/guides/java/e2e-testing/) guides.

For Micronaut HTTP tests, add `testImplementation("io.github.shibukawa.pgmem:pgmem-micronaut:1.18.0")`
alongside your application's Micronaut server dependencies. Then start the app
with `MicronautTestApp.start(pg.applicationFork())`. The helper selects an
available port, supplies the fork's JDBC URL, disables duplicate Flyway
migrations, and exposes `uri()` and `close()`. Pass a property map as the
second argument for application-specific settings.

For Spring Boot or Micronaut tests with no manual connection setup, use the
framework module's `@ShadowPg` annotation instead:

```java
@ShadowPg(schema = "db/schema.sql", seed = "db/seed.sql")
class OrderRepositoryTest {
    @Autowired OrderRepository orders; // use @Inject with Micronaut
    @Test void createsOrder() { orders.create("book"); }
}
```

Add `pgmem-spring` for Spring Boot or `pgmem-micronaut` for Micronaut as a test
dependency alongside the application's usual Spring Boot test or Micronaut
runtime dependencies. The SQL files are classpath resources in `src/test/resources`.
The annotation starts the application with a pgmem JDBC URL before its
DataSource is built, captures the post-startup state, and resets it before
each method. Omit `schema` if the application runs its own migrations. The
same annotation starts a random-port HTTP server for API or browser tests.
Each class owns and closes its server and fork. Test methods in one class
must run sequentially because they share the same application database.

Each connection to a fork gets its own PostgreSQL backend, including pooled
connections. Commit or close transactions before `snapshot()` or `reset()`:
they wait for open transactions and fail with code `busy` after their timeout.

## Building

`./gradlew build` compiles the Go binary for the host (`-Pgoos=linux
-Pgoarch=amd64` cross-compiles) into `pgmem-native`, and runs the tests
against it. Java 17 or newer.

## One database object per test

JUnit methods can inject `TestDatabase` from `io.github.shibukawa.pgmem.junit5`.
Its `dataSource()`, `jdbcUrl()`, `dsn()` and `fork()` refer to the same fork;
other pgmem parameters with the same template and scope share it. The extension
owns fork cleanup. `@PgmemFork(scope = ForkScope.CLASS)` shares it for read-only
classes. `TestDatabase.reset()` or `Server.reset(snapshot, timeout)` restores it
in place after active transactions end. Application pools should close before
their owning fork.

The [testing quickstart](../../website/src/content/docs/guides/java/testing.mdx)
is a complete JUnit example. The [dependency guide](../../website/src/content/docs/guides/java/basics.mdx)
includes Gradle host detection and Maven platform profiles for local/CI builds.

## Shared read guard and cleanup / 共有接続のガードと後片付け

`.sharedReadOnly(true)` guards CLASS-scoped `TestDatabase`, `DataSource` and
JDBC URL injection. METHOD scope remains writable; raw `Fork` access bypasses
the guard. Core `dsn(true)`, `jdbcUrl(true)` and `dataSource(true)` also expose
guarded connection views. The setting is a session default that callers can
change. Cleanup attempts all owned forks and the process, retaining test and
cleanup failures as suppressed exceptions. Configure `.forkTimeout(Duration)`
to bound automatic acquisition; it is unbounded by default.

`.sharedReadOnly(true)` は CLASS スコープの `TestDatabase`、`DataSource`、JDBC URL
注入を保護します。METHOD スコープと生の `Fork` は書き込み可能です。core の
`dsn(true)`、`jdbcUrl(true)`、`dataSource(true)` でもガード付きの接続先を取得できます。
呼び出し側が変更可能なセッションの既定値です。後片付けは管理するすべての
フォークとプロセスの解放を試み、テストと解放の失敗を suppressed exception に
残します。自動取得の期限は `.forkTimeout(Duration)` で指定します。既定は無期限です。

Runnable DataSource application / 実行可能な DataSource アプリ:
[examples](../../examples/README.md).
