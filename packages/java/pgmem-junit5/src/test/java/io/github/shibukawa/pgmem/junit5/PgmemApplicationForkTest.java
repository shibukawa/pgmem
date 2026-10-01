package io.github.shibukawa.pgmem.junit5;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertSame;

import io.github.shibukawa.pgmem.Fork;
import com.sun.net.httpserver.HttpServer;
import java.net.InetSocketAddress;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class PgmemApplicationForkTest {
    @RegisterExtension
    static PgmemExtension pg = PgmemExtension.builder()
            .database("app")
            .prepare(t -> PgmemExtensionTest.exec(t.jdbcUrl(), "CREATE TABLE t (id int); INSERT INTO t VALUES (1), (2)"))
            .build();

    // Simulates a framework reading its JDBC URL before JUnit's BeforeAll callback.
    static final Fork appDatabase = pg.applicationFork();

    @Test void appKeepsItsConnectionAcrossReset() throws SQLException {
        assertSame(appDatabase, pg.applicationFork());
        String url = appDatabase.jdbcUrl();
        try (Connection connection = appDatabase.dataSource().getConnection()) {
            PgmemExtensionTest.exec(url, "INSERT INTO t VALUES (3)");
            assertEquals(3, count(connection));
            appDatabase.reset();
            assertEquals(url, appDatabase.jdbcUrl());
            assertEquals(2, count(connection));
        }
    }

    @Test void httpAppSeesRestoredDataWithoutRestarting() throws Exception {
        HttpServer http = HttpServer.create(new InetSocketAddress("127.0.0.1", 0), 0);
        try (Connection connection = appDatabase.dataSource().getConnection()) {
            http.createContext("/count", exchange -> {
                try {
                    byte[] body = String.valueOf(count(connection)).getBytes(StandardCharsets.UTF_8);
                    exchange.sendResponseHeaders(200, body.length);
                    try (var out = exchange.getResponseBody()) { out.write(body); }
                } catch (SQLException e) {
                    exchange.sendResponseHeaders(500, -1);
                } finally {
                    exchange.close();
                }
            });
            http.start();
            String url = "http://127.0.0.1:" + http.getAddress().getPort() + "/count";
            HttpClient client = HttpClient.newHttpClient();
            HttpRequest request = HttpRequest.newBuilder(URI.create(url)).build();
            assertEquals("2", client.send(request, HttpResponse.BodyHandlers.ofString()).body());
            PgmemExtensionTest.exec(appDatabase.jdbcUrl(), "INSERT INTO t VALUES (3)");
            assertEquals("3", client.send(request, HttpResponse.BodyHandlers.ofString()).body());
            appDatabase.reset();
            assertEquals("2", client.send(request, HttpResponse.BodyHandlers.ofString()).body());
        } finally {
            http.stop(0);
        }
    }

    private static int count(Connection connection) throws SQLException {
        try (ResultSet rs = connection.createStatement().executeQuery("SELECT count(*) FROM t")) {
            rs.next();
            return rs.getInt(1);
        }
    }
}
