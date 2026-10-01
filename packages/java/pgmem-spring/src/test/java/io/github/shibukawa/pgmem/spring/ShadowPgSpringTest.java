package io.github.shibukawa.pgmem.spring;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;
import org.junit.jupiter.api.MethodOrderer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringBootConfiguration;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.test.web.server.LocalServerPort;
import org.springframework.context.annotation.Import;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@ShadowPg(classes = ShadowPgSpringTest.Application.class,
        schema = "shadow-pg-schema.sql", seed = "shadow-pg-seed.sql",
        properties = "spring.datasource.url=jdbc:postgresql://127.0.0.1:1/wrong")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class ShadowPgSpringTest {
    @SpringBootConfiguration @EnableAutoConfiguration @Import(Orders.class)
    static class Application {}

    @RestController
    static class Orders {
        private final JdbcTemplate jdbc;
        Orders(JdbcTemplate jdbc) { this.jdbc = jdbc; }
        @GetMapping("/orders/count") int count() {
            return jdbc.queryForObject("SELECT count(*) FROM orders", Integer.class);
        }
    }

    @Autowired JdbcTemplate jdbc;
    @LocalServerPort int port;

    @Test @Order(1) void appReadsAndWritesPgmem() throws Exception {
        try (var connection = jdbc.getDataSource().getConnection()) {
            assertTrue(connection.getMetaData().getURL().contains("127.0.0.1"));
        }
        assertEquals("1", countOverHttp());
        jdbc.update("INSERT INTO orders VALUES (2, 'test')");
        assertEquals("2", countOverHttp());
    }

    @Test @Order(2) void nextCaseStartsFromSeed() throws Exception {
        assertEquals("1", countOverHttp());
    }

    private String countOverHttp() throws Exception {
        HttpRequest request = HttpRequest.newBuilder(URI.create("http://127.0.0.1:" + port + "/orders/count")).build();
        return HttpClient.newHttpClient().send(request, HttpResponse.BodyHandlers.ofString()).body();
    }
}
