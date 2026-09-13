# pgmem for Java

A real PostgreSQL 18 that runs entirely in memory, packaged as a single
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
testImplementation("jp.shibu:pgmem-junit5:0.1.0")          // pulls in jp.shibu:pgmem
testRuntimeOnly("jp.shibu:pgmem-native:0.1.0:darwin-arm64") // the binary for your platform
testRuntimeOnly("org.postgresql:postgresql:42.7.7")         // your usual driver
```

Classifiers: `linux-x86_64`, `linux-arm64`, `darwin-x86_64`, `darwin-arm64`,
`windows-x86_64`, `windows-arm64`. Pick one with `os-maven-plugin` or Gradle's
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
@Test void report(@PgmemFork(scope = ForkScope.CLASS) DataSource ds) { ... } // shared by the class
@Test void audit(@PgmemFork("audit") String jdbcUrl) { ... }        // a named template
```

`Fork`, `DataSource` and `@PgmemFork String` parameters are resolved; several
parameters in one test share one fork. Forks are closed after each test
(or after the class for `ForkScope.CLASS`). Tests may run in parallel;
`maxForks` bounds how many forks exist at once and `forkTimeout` turns a
full pool into a test failure instead of a wait.

The server is one PostgreSQL session per fork; pooled connections
(HikariCP and friends) are serialized at transaction boundaries. Commit or
close every connection before `snapshot()`: it waits for open transactions
and fails with code `busy` after 30 s.

## Building

`./gradlew build` compiles the Go binary for the host (`-Pgoos=linux
-Pgoarch=amd64` cross-compiles) into `pgmem-native`, and runs the tests
against it. Java 17 or newer.
