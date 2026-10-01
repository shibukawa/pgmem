package io.github.shibukawa.pgmem.micronaut;

import io.github.shibukawa.pgmem.Fork;
import io.micronaut.context.ApplicationContext;
import io.micronaut.runtime.server.EmbeddedServer;
import java.net.URI;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Objects;

/** A Micronaut HTTP server connected to an existing pgmem fork for API and browser tests. */
public final class MicronautTestApp implements AutoCloseable {
    private final EmbeddedServer server;

    private MicronautTestApp(EmbeddedServer server) { this.server = server; }

    /** Start on an available port in the test environment. The caller owns and resets the fork. */
    public static MicronautTestApp start(Fork fork) { return start(fork, Map.of()); }

    /** Extra Micronaut properties override defaults except the fork's datasource URL. */
    public static MicronautTestApp start(Fork fork, Map<String, ?> overrides) {
        Objects.requireNonNull(fork, "fork");
        Objects.requireNonNull(overrides, "overrides");
        Map<String, Object> properties = new LinkedHashMap<>();
        properties.put("micronaut.server.port", -1);
        properties.put("datasources.default.username", fork.user());
        properties.put("datasources.default.password", "");
        properties.put("datasources.default.driverClassName", "org.postgresql.Driver");
        // The template was migrated before it was snapshotted.
        properties.put("flyway.datasources.default.enabled", false);
        properties.putAll(overrides);
        properties.put("datasources.default.url", fork.jdbcUrl());
        return new MicronautTestApp(ApplicationContext.run(EmbeddedServer.class, properties, "test"));
    }

    /** Base URI of the running HTTP server. */
    public URI uri() { return server.getURI(); }

    /** The underlying Micronaut server, for framework-specific assertions. */
    public EmbeddedServer server() { return server; }

    /** Stop Micronaut and its connection pool before the fork is closed. */
    @Override public void close() { server.close(); }
}
