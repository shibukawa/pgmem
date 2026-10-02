# Django E2E server

Start one child per pytest-xdist worker after the live fork exists: `python -m django runserver --noreload 127.0.0.1:PORT`. Point `DJANGO_SETTINGS_MODULE` to test settings that configure the normal PostgreSQL backend from the worker's fork DSN before app import. Each browser test requests `pgmem_live_db` before navigation. Use plain pytest browser tests; Django `TestCase` and pytest-django `live_server` manage another test database lifecycle. See [E2E lifecycle](e2e.md).
