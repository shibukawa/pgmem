package jp.shibu.pgmem;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.Map;

/** A frozen copy of a server's state that forks start from. */
public final class Snapshot implements AutoCloseable {
    private final Pgmem pgmem;
    private final String id;
    private final Server origin;
    private volatile boolean closed;

    Snapshot(Pgmem pgmem, String id, Server origin) {
        this.pgmem = pgmem;
        this.id = id;
        this.origin = origin;
    }

    public String id() { return id; }

    /** The server this snapshot was taken from. */
    public Server origin() { return origin; }

    /** Starts a fresh server on a copy of the snapshot, blocking while maxForks forks are alive. */
    public Fork fork() { return fork(null); }

    /** Like {@link #fork()} but fails with code {@code pool_timeout} after {@code timeout}. */
    public Fork fork(Duration timeout) {
        Map<String, Object> req = new LinkedHashMap<>();
        req.put("snapshot", id);
        if (timeout != null) req.put("timeout_ms", timeout.toMillis());
        Map<String, Object> res = pgmem.request("fork", req);
        return new Fork(pgmem, Pgmem.endpoint(res));
    }

    /** Rejects new forks; live forks keep running. Idempotent. */
    @Override public void close() {
        if (closed) return;
        closed = true;
        pgmem.request("close", Map.of("snapshot", id));
    }

    @Override public String toString() { return "Snapshot(" + id + " of " + origin.id() + ")"; }
}
