package jp.shibu.pgmem.junit5;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import javax.sql.DataSource;
import jp.shibu.pgmem.Fork;
import jp.shibu.pgmem.Server;
import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;
import org.junit.jupiter.api.extension.RegisterExtension;

@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class PgmemExtensionTest {

    @RegisterExtension
    static PgmemExtension pg = PgmemExtension.builder()
            .database("app")
            .maxForks(3)
            .prepare(t -> exec(t.jdbcUrl(), "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)"))
            .template("audit", t -> exec(t.jdbcUrl(), "CREATE TABLE a (id int)"))
            .build();

    static void exec(String url, String sql) {
        try (Connection c = DriverManager.getConnection(url); Statement st = c.createStatement()) {
            st.execute(sql);
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }

    static int count(String url, String table) {
        try (Connection c = DriverManager.getConnection(url);
             Statement st = c.createStatement();
             ResultSet rs = st.executeQuery("SELECT count(*) FROM " + table)) {
            rs.next();
            return rs.getInt(1);
        } catch (SQLException e) {
            throw new RuntimeException(e);
        }
    }

    static String classForkId;

    @Test @Order(1) void forkIsPrepared(Fork fork, DataSource ds, @PgmemFork String url) throws SQLException {
        assertEquals(fork.jdbcUrl(), url);
        assertEquals(2, count(url, "t"));
        exec(url, "INSERT INTO t VALUES (99)");
        try (Connection c = ds.getConnection(); ResultSet rs = c.createStatement().executeQuery("SELECT count(*) FROM t")) {
            rs.next();
            assertEquals(3, rs.getInt(1)); // same fork behind all three parameters
        }
    }

    @Test @Order(2) void nextTestGetsFreshFork(Fork fork) {
        assertEquals(2, count(fork.jdbcUrl(), "t"));
    }

    @Test @Order(3) void namedTemplate(@PgmemFork("audit") Fork audit, Fork app) {
        assertEquals("audit", audit.database());
        assertEquals(0, count(audit.jdbcUrl(), "a"));
        assertEquals("app", app.database());
        assertNotEquals(audit.id(), app.id());
    }

    @Test @Order(4) void classScopedWrite(@PgmemFork(scope = ForkScope.CLASS) Fork fork) {
        classForkId = fork.id();
        exec(fork.jdbcUrl(), "INSERT INTO t VALUES (3)");
        assertEquals(3, count(fork.jdbcUrl(), "t"));
    }

    @Test @Order(5) void classScopedSeesPreviousWrite(@PgmemFork(scope = ForkScope.CLASS) Fork fork, Fork fresh) {
        assertEquals(classForkId, fork.id());
        assertEquals(3, count(fork.jdbcUrl(), "t"));
        assertEquals(2, count(fresh.jdbcUrl(), "t"));
    }

    @Test @Order(6) void manualFork() {
        try (Fork f = pg.fork()) {
            assertEquals(2, count(f.jdbcUrl(), "t"));
        }
        Server template = pg.template();
        assertEquals(2, count(template.jdbcUrl(), "t"));
        assertTrue(pg.pgmem().pid() > 0);
    }

    @Nested
    class Inner {
        @Test void nestedClassUsesSameProcess(Fork fork) {
            assertEquals(2, count(fork.jdbcUrl(), "t"));
        }
    }
}
