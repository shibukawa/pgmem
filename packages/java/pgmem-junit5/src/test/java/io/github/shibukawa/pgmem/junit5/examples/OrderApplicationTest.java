package io.github.shibukawa.pgmem.junit5.examples;

import static org.junit.jupiter.api.Assertions.assertEquals;
import java.sql.SQLException;
import javax.sql.DataSource;
import io.github.shibukawa.pgmem.junit5.PgmemExtension;
import io.github.shibukawa.pgmem.junit5.TestDatabase;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class OrderApplicationTest {
    @RegisterExtension static PgmemExtension pg = PgmemExtension.builder()
            .prepare(server -> {
                try (var connection = server.dataSource().getConnection(); var statement = connection.createStatement()) {
                    statement.execute("CREATE TABLE orders(id int PRIMARY KEY)");
                } catch (SQLException error) { throw new RuntimeException(error); }
            }).build();

    static final class OrderApplication {
        private final DataSource database;
        OrderApplication(DataSource database) { this.database = database; }
        void createOrder(int id) throws SQLException {
            try (var connection = database.getConnection(); var statement = connection.prepareStatement("INSERT INTO orders(id) VALUES (?)")) {
                statement.setInt(1, id);
                statement.executeUpdate();
            }
        }
    }

    @Test void applicationAndAssertionsUseTheSameFork(TestDatabase database) throws Exception {
        var application = new OrderApplication(database.dataSource());
        application.createOrder(42);
        try (var connection = database.dataSource().getConnection();
             var statement = connection.createStatement(); var rows = statement.executeQuery("SELECT id FROM orders")) {
            rows.next();
            assertEquals(42, rows.getInt(1));
        }
    }
}
