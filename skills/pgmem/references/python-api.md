# Python HTTP API tests

Prepare the template with `pgmem_prepare` as in [unit tests](python-unit.md). A pytest worker owns one session-scoped `pgmem_live_fork`; `@shadow_pg(live=True)` resets it before each marked test and routes new connections to its stable endpoint. Make the HTTP client fixture function-scoped so lifespan and pool construction happen under the marker. Tests sharing a worker run sequentially. `pgmem_live_snapshots` can map names to snapshots; select one with `@pytest.mark.pgmem_dataset("name")`.

Read the matching framework reference: [FastAPI or Starlette](python-api-fastapi-starlette.md), [Flask](python-api-flask.md), [Django](python-api-django.md). These describe client setup; use `@shadow_pg(live=True)` on the test that requests that client.

The live fork occupies a snapshot slot for the session. Leave capacity for any per-test forks. Commit/close open transactions and finish background jobs before the next reset. For parallel API cases, use pytest-xdist processes, each with its own app client and pgmem process; a shared external app defeats worker isolation.

Full guide: `website/src/content/docs/guides/python/api-testing.mdx`.
