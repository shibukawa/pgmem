package io.github.shibukawa.pgmem;

/** A server started from a {@link Snapshot}; closing it frees its pool slot. */
public final class Fork extends Server {
    Fork(Pgmem pgmem, Endpoint endpoint) { super(pgmem, endpoint); }
}
