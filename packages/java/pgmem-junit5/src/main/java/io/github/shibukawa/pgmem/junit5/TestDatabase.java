package io.github.shibukawa.pgmem.junit5;

import io.github.shibukawa.pgmem.Fork;
import javax.sql.DataSource;
import java.time.Duration;

/** An injected test database; all handles refer to the same automatically managed fork. */
public final class TestDatabase {
    private final Fork fork;
    private final boolean readOnly;
    TestDatabase(Fork fork, boolean readOnly) { this.fork = fork; this.readOnly = readOnly; }
    public Fork fork() { return fork; }
    public String dsn() { return fork.dsn(readOnly); }
    public String jdbcUrl() { return fork.jdbcUrl(readOnly); }
    public DataSource dataSource() { return fork.dataSource(readOnly); }
    public void reset() { fork.reset(); }
    public void reset(Duration timeout) { fork.reset(null, timeout); }
}
