package io.github.shibukawa.pgmem.junit5;

import static org.junit.jupiter.api.Assertions.*;
import java.sql.Connection;
import java.sql.SQLException;
import javax.sql.DataSource;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class ReadOnlyExtensionTest {
    @RegisterExtension static PgmemExtension pg = PgmemExtension.builder()
            .sharedReadOnly(true).forkScope(ForkScope.CLASS)
            .prepare(server -> PgmemExtensionTest.exec(server.jdbcUrl(), "CREATE TABLE guarded(v int)")).build();

    @Test void classHandlesGuardWrites(TestDatabase db, DataSource ds) throws Exception {
        for (DataSource source : new DataSource[] { db.dataSource(), ds }) {
            try (Connection connection = source.getConnection(); var statement = connection.createStatement()) {
                try (var rows = statement.executeQuery("SELECT count(*) FROM guarded")) { rows.next(); assertEquals(0, rows.getInt(1)); }
                SQLException failure = assertThrows(SQLException.class, () -> statement.execute("INSERT INTO guarded VALUES (1)"));
                assertEquals("25006", failure.getSQLState());
            }
        }
    }

    @Test void methodForkStillAllowsWrites(@PgmemFork(scope = ForkScope.METHOD) TestDatabase db) {
        PgmemExtensionTest.exec(db.jdbcUrl(), "INSERT INTO guarded VALUES (1)");
    }
}
