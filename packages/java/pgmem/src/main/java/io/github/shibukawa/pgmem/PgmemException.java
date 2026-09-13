package io.github.shibukawa.pgmem;

/** Base class for pgmem failures. */
public class PgmemException extends RuntimeException {
    public PgmemException(String message) { super(message); }
    public PgmemException(String message, Throwable cause) { super(message, cause); }
}
