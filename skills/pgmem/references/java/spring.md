# Spring Boot integration

Add test dependency `io.github.shibukawa.pgmem:pgmem-spring` plus the PostgreSQL JDBC driver and platform `pgmem-native` dependency (see `website/src/content/docs/guides/java/basics.mdx`). Import `io.github.shibukawa.pgmem.spring.ShadowPg`; annotate a test class `@ShadowPg(schema = "db/schema.sql")`. It supplies the fork URL through normal Spring DataSource configuration, starts a random-port application, and captures post-startup data as reset baseline. Omit schema if Flyway/Liquibase or the app migrates on startup; add `seed` when required.

For unit tests, `@Autowired` the normal repository bean. For HTTP tests, inject `@LocalServerPort int port` and send to `http://127.0.0.1:<port>`. For E2E, add `@UsePlaywright` and navigate a `Page` to that URL. Keep methods sequential because they share an application context and fork. See [unit](unit.md), [API](api.md), and [E2E](e2e.md).
