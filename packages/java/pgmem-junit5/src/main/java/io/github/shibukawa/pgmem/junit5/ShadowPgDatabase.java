package io.github.shibukawa.pgmem.junit5;

import io.github.shibukawa.pgmem.Fork;
import io.github.shibukawa.pgmem.Pgmem;
import io.github.shibukawa.pgmem.Snapshot;
import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.Objects;

/** A prepared pgmem database and stable application fork owned by a test class. */
public final class ShadowPgDatabase implements AutoCloseable {
    private final Pgmem pgmem;
    private final Snapshot baseline;
    private final Fork fork;
    private final Class<?> testClass;
    private final String deferredSeed;
    private Snapshot applicationBaseline;
    private boolean closed;

    private ShadowPgDatabase(Pgmem pgmem, Snapshot baseline, Fork fork, Class<?> testClass, String deferredSeed) {
        this.pgmem = pgmem;
        this.baseline = baseline;
        this.fork = fork;
        this.testClass = testClass;
        this.deferredSeed = deferredSeed;
    }

    /** Start pgmem, run optional classpath SQL files, and fork the prepared snapshot. */
    public static ShadowPgDatabase start(Class<?> testClass, String database, String schema, String seed) {
        Objects.requireNonNull(testClass, "testClass");
        Pgmem pg = Pgmem.builder().database(database).start();
        Snapshot snapshot = null;
        Fork appFork = null;
        try {
            executeResource(testClass, pg.template().jdbcUrl(), schema);
            if (!schema.isEmpty()) executeResource(testClass, pg.template().jdbcUrl(), seed);
            snapshot = pg.template().snapshot();
            appFork = snapshot.fork();
            return new ShadowPgDatabase(pg, snapshot, appFork, testClass, schema.isEmpty() ? seed : "");
        } catch (RuntimeException e) {
            if (appFork != null) appFork.close();
            if (snapshot != null) snapshot.close();
            pg.close();
            throw e;
        }
    }

    private static void executeResource(Class<?> testClass, String url, String resource) {
        if (resource.isEmpty()) return;
        String name = resource.startsWith("/") ? resource.substring(1) : resource;
        ClassLoader loader = testClass.getClassLoader();
        try (InputStream input = loader.getResourceAsStream(name)) {
            if (input == null) throw new IllegalArgumentException("ShadowPg SQL resource not found: " + resource);
            String sql = new String(input.readAllBytes(), StandardCharsets.UTF_8);
            try (Connection connection = DriverManager.getConnection(url);
                 Statement statement = connection.createStatement()) {
                statement.execute(sql);
            }
        } catch (IOException | SQLException e) {
            throw new IllegalStateException("Cannot prepare ShadowPg SQL resource: " + resource, e);
        }
    }

    public Fork fork() { return fork; }

    /** Capture changes made by application startup, including its own migrations. */
    public synchronized void captureApplicationBaseline() {
        if (applicationBaseline != null) return;
        executeResource(testClass, fork.jdbcUrl(), deferredSeed);
        applicationBaseline = fork.snapshot();
    }

    /** Restore the prepared application baseline without changing pooled connections. */
    public synchronized void reset() { fork.reset(applicationBaseline == null ? baseline : applicationBaseline); }

    @Override public synchronized void close() {
        if (closed) return;
        closed = true;
        try { fork.close(); }
        finally {
            try {
                if (applicationBaseline != null) applicationBaseline.close();
            } finally {
                try { baseline.close(); }
                finally { pgmem.close(); }
            }
        }
    }
}
