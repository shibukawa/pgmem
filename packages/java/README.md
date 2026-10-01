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

Each connection to a fork gets its own PostgreSQL backend, including pooled
connections. Commit or close transactions before `snapshot()` or `reset()`:
they wait for open transactions and fail with code `busy` after their timeout.

## Building

`./gradlew build` compiles the Go binary for the host (`-Pgoos=linux
-Pgoarch=amd64` cross-compiles) into `pgmem-native`, and runs the tests
against it. Java 17 or newer.
