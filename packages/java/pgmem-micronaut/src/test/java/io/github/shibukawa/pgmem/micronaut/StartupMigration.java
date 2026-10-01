package io.github.shibukawa.pgmem.micronaut;

import io.micronaut.context.annotation.Context;
import io.micronaut.context.annotation.Requires;
import java.sql.SQLException;
import javax.sql.DataSource;

@Context
@Requires(property = "shadow.migration", value = "true")
public final class StartupMigration {
    public StartupMigration(DataSource dataSource) throws SQLException {
        try (var connection = dataSource.getConnection();
             var statement = connection.createStatement()) {
            statement.execute("CREATE TABLE startup_data (id integer)");
        }
    }
}
