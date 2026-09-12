package jp.shibu.pgmem;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.time.Duration;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;
import org.junit.jupiter.api.Test;

class PgmemTest {

    static void exec(Server s, String sql) throws SQLException {
        try (Connection c = DriverManager.getConnection(s.jdbcUrl()); Statement st = c.createStatement()) {
            st.execute(sql);
        }
    }

    static int count(Server s) throws SQLException {
        try (Connection c = s.dataSource().getConnection();
             Statement st = c.createStatement();
             ResultSet rs = st.executeQuery("SELECT count(*) FROM t")) {
            rs.next();
            return rs.getInt(1);
        }
    }

    @Test void startAndQuery() throws Exception {
        try (Pgmem pg = Pgmem.builder().database("app").start()) {
            assertEquals("app", pg.template().database());
            assertTrue(pg.pid() > 0);
            assertTrue(pg.template().jdbcUrl().startsWith("jdbc:postgresql://127.0.0.1:"));
            try (Connection c = DriverManager.getConnection(pg.template().jdbcUrl());
                 ResultSet rs = c.createStatement().executeQuery("SELECT 1 + 1")) {
                rs.next();
                assertEquals(2, rs.getInt(1));
            }
        }
    }

    @Test void snapshotForkIsolation() throws Exception {
        try (Pgmem pg = Pgmem.builder().database("app").start()) {
            exec(pg.template(), "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)");
            Snapshot snap = pg.template().snapshot(2);
            try (Fork f1 = snap.fork(); Fork f2 = snap.fork()) {
                assertNotEquals(f1.id(), f2.id());
                exec(f1, "INSERT INTO t VALUES (3)");
                assertEquals(3, count(f1));
                assertEquals(2, count(f2));
                assertEquals(2, count(pg.template()));

                ProtocolException e = assertThrows(ProtocolException.class, () -> snap.fork(Duration.ofMillis(200)));
                assertEquals("pool_timeout", e.code());

                // a blocked fork is released by a close from another thread
                CompletableFuture<Integer> blocked = CompletableFuture.supplyAsync(() -> {
                    try (Fork f3 = snap.fork()) { return count(f3); } catch (SQLException ex) { throw new RuntimeException(ex); }
                });
                Thread.sleep(200);
                assertFalse(blocked.isDone());
                f1.close();
                assertEquals(2, blocked.get(30, TimeUnit.SECONDS));
                f1.close(); // idempotent
            }
            snap.close();
            ProtocolException e = assertThrows(ProtocolException.class, snap::fork);
            assertEquals("unknown_id", e.code());
        }
    }

    @Test void snapshotBusyWhileTransactionOpen() throws Exception {
        try (Pgmem pg = Pgmem.start()) {
            try (Connection c = DriverManager.getConnection(pg.template().jdbcUrl())) {
                c.setAutoCommit(false);
                c.createStatement().execute("SELECT 1");
                ProtocolException e = assertThrows(ProtocolException.class,
                        () -> pg.template().snapshot(0, Duration.ofMillis(200)));
                assertEquals("busy", e.code());
                c.commit();
            }
            pg.template().snapshot(0, Duration.ofSeconds(5)).close();
        }
    }

    @Test void startServerExtraTemplate() throws Exception {
        try (Pgmem pg = Pgmem.start()) {
            Server audit = pg.startServer("audit", "postgres", List.of("log_statement=all"));
            assertEquals("audit", audit.database());
            assertNotEquals(pg.template().id(), audit.id());
            exec(audit, "CREATE TABLE t (id int)");
            try (Fork f = audit.snapshot().fork()) {
                assertEquals("audit", f.database());
                assertEquals(0, count(f));
            }
        }
    }

    @Test void serverExitFailsRequests() throws Exception {
        Pgmem pg = Pgmem.start();
        ProcessHandle.of(pg.pid()).orElseThrow().destroyForcibly();
        Thread.sleep(500);
        assertThrows(ServerExitedException.class, () -> pg.template().snapshot());
        pg.close();
    }

    @Test void jsonRoundTrip() {
        java.util.Map<String, Object> in = new java.util.LinkedHashMap<>();
        in.put("id", 1L); in.put("op", "fork"); in.put("s", "a\"b\n"); in.put("l", List.of(1L, true)); in.put("n", null);
        String s = Json.write(in);
        Map<String, Object> m = Json.parseObject(s);
        assertEquals(1L, m.get("id"));
        assertEquals("a\"b\n", m.get("s"));
        assertEquals(List.of(1L, true), m.get("l"));
        assertTrue(m.containsKey("n"));
        assertEquals(1.5, Json.parse("1.5"));
        assertEquals(Map.of(), Json.parse(" {} "));
    }

    @Test void classifierMatchesBuild() {
        assertTrue(BinaryLocator.classifier().matches("(linux|darwin|windows)-(x86_64|arm64)"));
    }
}
