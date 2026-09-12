package jp.shibu.pgmem;

import java.util.Map;

/** Where one server inside the pgmem process listens. */
public final class Endpoint {
    private final String id;
    private final String host;
    private final int port;
    private final String user;
    private final String database;
    private final String dsn;

    Endpoint(Map<String, Object> m) {
        this.id = (String) m.get("id");
        this.host = (String) m.get("host");
        this.port = ((Number) m.get("port")).intValue();
        this.user = (String) m.get("user");
        this.database = (String) m.get("database");
        this.dsn = (String) m.get("dsn");
    }

    public String id() { return id; }
    public String host() { return host; }
    public int port() { return port; }
    public String user() { return user; }
    public String database() { return database; }

    /** libpq-style URL, e.g. {@code postgres://postgres@127.0.0.1:5432/app?sslmode=disable}. */
    public String dsn() { return dsn; }

    /** pgjdbc URL with user and sslmode already set. */
    public String jdbcUrl() {
        return "jdbc:postgresql://" + host + ":" + port + "/" + database + "?user=" + user + "&sslmode=disable";
    }

    @Override public String toString() { return id + " " + dsn; }
}
