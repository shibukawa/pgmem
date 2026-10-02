# Java E2E tests with Playwright

Use Playwright's JUnit integration, `@UsePlaywright`, together with framework `@ShadowPg(schema = "db/schema.sql")`: [Spring Boot](spring.md) or [Micronaut](micronaut.md). The annotation starts the real HTTP app with a pgmem fork and resets it before each method. Accept a Playwright `Page` parameter, navigate to the injected random-port URL, and assert UI behavior. The server and pool remain alive for the class.

If the application migrates on startup, omit `schema`; that post-startup state becomes the reset baseline. Use `seed = "db/seed.sql"` for initial rows. An external JVM must receive the fork JDBC URL before creating its pool.

Run methods in one class sequentially. Parallel classes need separate context, fork, random port, and browser context; `junit.jupiter.execution.parallel.mode.default = same_thread` plus `...mode.classes.default = concurrent` is suitable. Finish browser-triggered requests and transactions before the next reset. Full examples: `website/src/content/docs/guides/java/e2e-testing.mdx`.
