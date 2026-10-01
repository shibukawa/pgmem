package io.github.shibukawa.pgmem;

import java.time.Duration;
import java.util.LinkedHashMap;
import java.util.Map;

/** A server started from a {@link Snapshot}; closing it frees its pool slot. */
public final class Fork extends Server {
    Fork(Pgmem pgmem, Endpoint endpoint) { super(pgmem, endpoint); }

    /** Restore the snapshot this fork was created from, keeping its JDBC URL and connections valid. */
    public void reset() { reset(null, Duration.ofSeconds(5)); }

    /** Restore a chosen snapshot without changing this fork's JDBC URL. */
    public void reset(Snapshot snapshot) { reset(snapshot, Duration.ofSeconds(5)); }

    /** Restore data, waiting at most {@code timeout} for open transactions; {@code null} waits forever. */
    public void reset(Snapshot snapshot, Duration timeout) {
        Map<String, Object> req = new LinkedHashMap<>();
        req.put("server", id());
        if (snapshot != null) req.put("snapshot", snapshot.id());
        if (timeout != null) req.put("timeout_ms", timeout.toMillis());
        pgmem.request("reset", req);
    }
}
