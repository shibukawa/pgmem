# Micronaut integration

Add test dependency `io.github.shibukawa.pgmem:pgmem-micronaut` plus the PostgreSQL JDBC driver and platform `pgmem-native` dependency (see `website/src/content/docs/guides/java/basics.mdx`). Import `io.github.shibukawa.pgmem.micronaut.ShadowPg`; annotate a test class `@ShadowPg(schema = "db/schema.sql")`. It starts an `EmbeddedServer`, directs the app's normal DataSource to pgmem, and resets the post-startup baseline before each method. Omit schema if the application migrates on startup; add `seed` if needed.

For unit tests, `@Inject` the normal repository bean. For HTTP tests, `@Inject EmbeddedServer app` and send requests to `app.getURI().resolve("/orders")`. For E2E, add `@UsePlaywright` and navigate a `Page` to that URL. Keep methods sequential within a class because they share a context and fork. See [unit](java-unit.md), [API](java-api.md), and [E2E](java-e2e.md).
