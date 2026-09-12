package jp.shibu.pgmem;

/** The server answered a request with an error; {@link #code()} is its snake_case code. */
public class ProtocolException extends PgmemException {
    private final String code;

    public ProtocolException(String code, String message) {
        super(code + ": " + message);
        this.code = code;
    }

    /** One of unknown_op, unknown_id, snapshot_closed, pool_timeout, protocol, internal. */
    public String code() { return code; }
}
