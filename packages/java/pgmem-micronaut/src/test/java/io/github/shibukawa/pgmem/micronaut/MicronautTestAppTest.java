package io.github.shibukawa.pgmem.micronaut;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import io.github.shibukawa.pgmem.Fork;
import io.github.shibukawa.pgmem.Pgmem;
import io.github.shibukawa.pgmem.Snapshot;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.Map;
import org.junit.jupiter.api.Test;

class MicronautTestAppTest {
    @Test void startsHttpServerAgainstForkAndClosesIt() throws Exception {
        try (Pgmem pg = Pgmem.builder().database("app").start()) {
            try (Connection c = DriverManager.getConnection(pg.template().jdbcUrl());
                 Statement st = c.createStatement()) {
                st.execute("CREATE TABLE t (id int); INSERT INTO t VALUES (1)");
            }
            try (Snapshot baseline = pg.template().snapshot();
                 Fork fork = baseline.fork();
                 MicronautTestApp app = MicronautTestApp.start(fork, Map.of("micronaut.server.host", "127.0.0.1"))) {
                assertTrue(app.uri().getPort() > 0);
                assertEquals(fork.jdbcUrl(), app.server().getApplicationContext().getEnvironment()
                        .getProperty("datasources.default.url", String.class).orElseThrow());
                HttpRequest request = HttpRequest.newBuilder(app.uri().resolve("/missing")).build();
                assertEquals(404, HttpClient.newHttpClient()
                        .send(request, HttpResponse.BodyHandlers.discarding()).statusCode());

                try (Connection pooled = DriverManager.getConnection(fork.jdbcUrl())) {
                    try (Connection writer = DriverManager.getConnection(fork.jdbcUrl());
                         Statement st = writer.createStatement()) {
                        st.execute("INSERT INTO t VALUES (2)");
                    }
                    assertEquals(2, count(pooled));
                    fork.reset();
                    assertEquals(1, count(pooled));
                }
            }
        }
    }

    private static int count(Connection connection) throws Exception {
        try (Statement st = connection.createStatement();
             ResultSet rs = st.executeQuery("SELECT count(*) FROM t")) {
            rs.next();
            return rs.getInt(1);
        }
    }
}
