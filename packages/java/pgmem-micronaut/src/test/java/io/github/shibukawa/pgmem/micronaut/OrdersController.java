package io.github.shibukawa.pgmem.micronaut;

import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Get;
import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import javax.sql.DataSource;

@Controller("/orders")
public final class OrdersController {
    private final DataSource dataSource;

    public OrdersController(DataSource dataSource) { this.dataSource = dataSource; }

    @Get("/count") public int count() throws SQLException {
        try (Connection connection = dataSource.getConnection();
             var statement = connection.createStatement();
             ResultSet rows = statement.executeQuery("SELECT count(*) FROM orders")) {
            rows.next();
            return rows.getInt(1);
        }
    }
}
