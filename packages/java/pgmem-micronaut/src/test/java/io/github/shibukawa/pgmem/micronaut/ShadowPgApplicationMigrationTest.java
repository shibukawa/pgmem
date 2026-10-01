package io.github.shibukawa.pgmem.micronaut;

import static org.junit.jupiter.api.Assertions.assertEquals;

import jakarta.inject.Inject;
import javax.sql.DataSource;
import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;

@ShadowPg(seed = "startup-seed.sql", properties = "shadow.migration=true")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class ShadowPgApplicationMigrationTest {
    @Inject DataSource dataSource;

    @Test @Order(1) void appMigrationIsInBaseline() throws Exception {
        assertEquals(1, count());
        try (var connection = dataSource.getConnection();
             var statement = connection.createStatement()) {
            statement.execute("INSERT INTO startup_data VALUES (2)");
        }
    }

    @Test @Order(2) void restoresAppMigrationAndSeed() throws Exception { assertEquals(1, count()); }

    private int count() throws Exception {
        try (var connection = dataSource.getConnection();
             var statement = connection.createStatement();
             var rows = statement.executeQuery("SELECT count(*) FROM startup_data")) {
            rows.next();
            return rows.getInt(1);
        }
    }
}
