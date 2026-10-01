package io.github.shibukawa.pgmem.spring;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.MethodOrderer;
import org.junit.jupiter.api.Order;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestMethodOrder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringBootConfiguration;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.test.annotation.Rollback;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.transaction.support.TransactionSynchronizationManager;

@ShadowPg(classes = ShadowPgApplicationMigrationTest.Application.class, seed = "startup-seed.sql")
@TestMethodOrder(MethodOrderer.OrderAnnotation.class)
class ShadowPgApplicationMigrationTest {
    @SpringBootConfiguration @EnableAutoConfiguration
    static class Application {
        @Bean ApplicationRunner migrate(JdbcTemplate jdbc) {
            return args -> jdbc.execute("CREATE TABLE startup_data (id integer)");
        }
    }

    @Autowired JdbcTemplate jdbc;

    @Test @Order(1) @Transactional @Rollback(false) void appMigrationIsInBaseline() {
        assertTrue(TransactionSynchronizationManager.isActualTransactionActive());
        assertEquals(1, count());
        jdbc.update("INSERT INTO startup_data VALUES (2)");
    }

    @Test @Order(2) void restoresAppMigrationAndSeed() { assertEquals(1, count()); }

    private int count() { return jdbc.queryForObject("SELECT count(*) FROM startup_data", Integer.class); }
}
