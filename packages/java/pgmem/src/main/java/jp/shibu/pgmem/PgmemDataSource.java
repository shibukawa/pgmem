package jp.shibu.pgmem;

import java.io.PrintWriter;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.SQLFeatureNotSupportedException;
import java.util.logging.Logger;
import javax.sql.DataSource;

/**
 * A {@link DataSource} over {@link DriverManager} for one server, so the core
 * artifact needs no compile-time dependency on pgjdbc. Wrap it in a pool if
 * you want one; pooled connections are serialized at transaction boundaries.
 */
final class PgmemDataSource implements DataSource {
    private final String url;
    private PrintWriter logWriter;
    private int loginTimeout;

    PgmemDataSource(String url) { this.url = url; }

    @Override public Connection getConnection() throws SQLException { return DriverManager.getConnection(url); }

    @Override public Connection getConnection(String username, String password) throws SQLException {
        return DriverManager.getConnection(url, username, password);
    }

    @Override public PrintWriter getLogWriter() { return logWriter; }
    @Override public void setLogWriter(PrintWriter out) { this.logWriter = out; }
    @Override public void setLoginTimeout(int seconds) { this.loginTimeout = seconds; }
    @Override public int getLoginTimeout() { return loginTimeout; }
    @Override public Logger getParentLogger() throws SQLFeatureNotSupportedException {
        throw new SQLFeatureNotSupportedException();
    }
    @Override public <T> T unwrap(Class<T> iface) throws SQLException {
        if (iface.isInstance(this)) return iface.cast(this);
        throw new SQLException("not a wrapper for " + iface);
    }
    @Override public boolean isWrapperFor(Class<?> iface) { return iface.isInstance(this); }
    @Override public String toString() { return "PgmemDataSource(" + url + ")"; }
}
