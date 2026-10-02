# Flask API client

Create a function-scoped fixture: `app = create_app(); app.config.update(TESTING=True); with app.test_client() as client: yield client`. Mark tests requesting `client` with `@shadow_pg(live=True)`. Flask's WSGI client runs in-process; new PostgreSQL driver connections are routed to the stable live fork, reset before each test. Keep the app's normal database driver and configuration.

Close application database work before the next reset; run concurrent cases in distinct pytest-xdist processes with separate app instances. See [API lifecycle](api.md) and `website/src/content/docs/guides/python/api-testing.mdx`.
