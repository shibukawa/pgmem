package io.github.shibukawa.pgmem.junit5;

import static org.junit.jupiter.api.Assertions.*;
import java.sql.SQLException;
import javax.sql.DataSource;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.RegisterExtension;

class DeclarativeReadOnlyExtensionTest {
    @RegisterExtension static PgmemExtension pg = PgmemExtension.builder().sharedReadOnly(true)
            .prepare(server -> PgmemExtensionTest.exec(server.jdbcUrl(), "CREATE TABLE guarded(v int)")).build();

    @Test @PgmemTest(fork = false)
    void sharedSelectionIsGuardedWithDefaultMethodScope(TestDatabase db, DataSource dataSource) throws Exception {
        for (DataSource source : new DataSource[] { db.dataSource(), dataSource }) {
            try (var connection = source.getConnection(); var statement = connection.createStatement()) {
                SQLException error = assertThrows(SQLException.class, () -> statement.execute("INSERT INTO guarded VALUES (1)"));
                assertEquals("25006", error.getSQLState());
            }
        }
    }
}
