# Python browser E2E tests

Use `pgmem_prepare` for schema/seed and a session-scoped `pgmem_live_fork` per pytest worker. Start a real application process **after** obtaining `pgmem_live_fork.dsn`, passing it through the app's existing test configuration. Keep one app process and endpoint per worker. Each browser test requests `pgmem_live_db`, which resets the stable endpoint before navigation; a named dataset may be selected through `pgmem_live_snapshots` and `@pytest.mark.pgmem_dataset`.

Use pytest-playwright's `page` fixture for browser actions. Read the matching launch note: [FastAPI/Starlette](python-e2e-fastapi-starlette.md), [Flask](python-e2e-flask.md), or [Django](python-e2e-django.md).

Use pytest-xdist for parallel E2E tests. Each worker needs its own pgmem process, app process, HTTP port, and browser context. Finish requests/background jobs and transactions before the next reset. Avoid Django `TestCase` and pytest-django `live_server` with this lifecycle.

Full launch fixture and readiness loop: `website/src/content/docs/guides/python/e2e-testing.mdx`.
