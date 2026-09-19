package io.github.shibukawa.pgmem;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.Map;
import javax.sql.DataSource;

/** A listening PostgreSQL server inside the pgmem process: a template or a {@link Fork}. */
public class Server implements AutoCloseable {
    final Pgmem pgmem;
    private final Endpoint endpoint;
    private volatile boolean closed;

    Server(Pgmem pgmem, Endpoint endpoint) {
        this.pgmem = pgmem;
        this.endpoint = endpoint;
    }

    public Endpoint endpoint() { return endpoint; }
    public String id() { return endpoint.id(); }
    public String host() { return endpoint.host(); }
    public int port() { return endpoint.port(); }
    public String user() { return endpoint.user(); }
    public String database() { return endpoint.database(); }
    public String dsn() { return endpoint.dsn(); }
    public String jdbcUrl() { return endpoint.jdbcUrl(); }

    /** A DriverManager-backed DataSource for this server (needs pgjdbc on the classpath). */
    public DataSource dataSource() { return new PgmemDataSource(jdbcUrl()); }

    /** Checkpoints and copies this server's state; forks start from the copy. */
    public Snapshot snapshot() { return snapshot(0); }

    /**
     * @param maxForks forks alive at once before {@link Snapshot#fork()} blocks; 0 = derived from the server's memory (a quarter of its limit divided by a fork's cost)
     */
    public Snapshot snapshot(int maxForks) { return snapshot(maxForks, Duration.ofSeconds(30)); }

    /**
     * The snapshot waits for open transactions to end, so commit or close every
     * connection first; after {@code timeout} (null = forever) it fails with a
     * {@link ProtocolException} whose code is {@code busy}.
     */
    public Snapshot snapshot(int maxForks, Duration timeout) {
        Map<String, Object> req = new LinkedHashMap<>();
        req.put("server", id());
        if (maxForks > 0) req.put("max_forks", maxForks);
        if (timeout != null) req.put("timeout_ms", timeout.toMillis());
        Map<String, Object> res = pgmem.request("snapshot", req);
        return new Snapshot(pgmem, (String) res.get("snapshot"), this);
    }

    /** Stops this server. Idempotent. */
    @Override public void close() {
        if (closed) return;
        closed = true;
        pgmem.request("close", Map.of("server", id()));
    }

    public boolean isClosed() { return closed; }

    @Override public String toString() { return getClass().getSimpleName() + "(" + endpoint + ")"; }
}
