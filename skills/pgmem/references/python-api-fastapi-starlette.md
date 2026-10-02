# FastAPI and Starlette API clients

Create a function-scoped `client` fixture. In it, use `with TestClient(create_app()) as client: yield client`; import `TestClient` from `fastapi.testclient` or `starlette.testclient`. The context manager runs application lifespan within `@shadow_pg(live=True)`. Then test normal routes using `client.get(...)`, `client.post(..., json=...)`, and assertions on responses.

Create the app and its pool only after the test's marker is active. A module-global pool may hold connections to a previous destination. Keep marked methods sequential within one worker. See [API lifecycle](python-api.md) and `website/src/content/docs/guides/python/api-testing.mdx`.
