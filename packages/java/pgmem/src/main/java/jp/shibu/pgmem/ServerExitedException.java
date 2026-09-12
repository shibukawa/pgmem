package jp.shibu.pgmem;

/** The pgmem process ended while a request was pending or before it could be sent. */
public class ServerExitedException extends PgmemException {
    public ServerExitedException(String message) { super(message); }
    public ServerExitedException(String message, Throwable cause) { super(message, cause); }
}
