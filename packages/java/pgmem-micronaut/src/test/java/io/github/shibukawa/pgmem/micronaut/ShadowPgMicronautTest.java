package io.github.shibukawa.pgmem.micronaut;

import static org.junit.jupiter.api.Assertions.assertEquals;

import io.micronaut.runtime.server.EmbeddedServer;
import jakarta.inject.Inject;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import javax.sql.DataSource;
import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;

@ShadowPg(schema = "shadow-pg-schema.sql", seed = "shadow-pg-seed.sql",
        properties = "datasources.default.url=jdbc:postgresql://127.0.0.1:1/wrong")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class ShadowPgMicronautTest {
    @Inject EmbeddedServer server;
    @Inject DataSource dataSource;

    @Test @Order(1) void appReadsAndWritesPgmem() throws Exception {
        assertEquals("1", countOverHttp());
        try (var connection = dataSource.getConnection();
             var statement = connection.createStatement()) {
            statement.execute("INSERT INTO orders VALUES (2, 'test')");
        }
        assertEquals("2", countOverHttp());
    }

    @Test @Order(2) void nextCaseStartsFromSeed() throws Exception {
        assertEquals("1", countOverHttp());
    }

    private String countOverHttp() throws Exception {
        HttpRequest request = HttpRequest.newBuilder(server.getURI().resolve("/orders/count")).build();
        return HttpClient.newHttpClient().send(request, HttpResponse.BodyHandlers.ofString()).body();
    }
}
