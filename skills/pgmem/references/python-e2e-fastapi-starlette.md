# FastAPI / Starlette E2E server

Start one Uvicorn child per pytest-xdist worker after `pgmem_live_fork` exists: `python -m uvicorn myapp.main:app --host 127.0.0.1 --port PORT` with `DATABASE_URL=pgmem_live_fork.dsn` in its environment. Disable reload and multiple app workers. Wait for readiness, expose the worker's `app_url`, and terminate/wait for the child at fixture teardown. Each browser test requests `pgmem_live_db` before navigation. See [E2E lifecycle](python-e2e.md) and the launch fixture in `website/src/content/docs/guides/python/e2e-testing.mdx`.
