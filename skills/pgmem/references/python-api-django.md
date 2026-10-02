# Django API client

Use plain pytest functions marked `@shadow_pg(live=True)`. In a function-scoped fixture, set `DJANGO_SETTINGS_MODULE` to the application's test settings if needed, call `django.setup()`, and return `django.test.Client()`. The test settings must use Django's normal PostgreSQL backend; initialize before Django has established a connection to another database.

Do not combine this lifecycle with Django `TestCase`, pytest-django `db`/`transactional_db`, or `live_server`: they create, roll back, or flush a separate test database. Application `transaction.atomic()` and savepoints remain usable. For an external Django server, use the [E2E lifecycle](python-e2e.md) with `runserver --noreload` and settings that read the worker's fork DSN.

See [API lifecycle](python-api.md) and `website/src/content/docs/guides/python/api-testing.mdx`.
