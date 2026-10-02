# Flask E2E server

Start one child per pytest-xdist worker after the live fork exists: `python -m flask --app myapp run --no-reload --host 127.0.0.1 --port PORT`. Flask does not read `DATABASE_URL` automatically; its existing test configuration must consume the fork DSN before importing the app. Pass the DSN in the child's environment, wait for readiness, and stop the child at worker teardown. Each browser test requests `pgmem_live_db` before navigation. See [E2E lifecycle](python-e2e.md).
