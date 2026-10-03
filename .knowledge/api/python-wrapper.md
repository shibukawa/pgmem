---
id: api:python-wrapper
type: api
title: Python Wrapper
---
PyPI package pgmem in packages/python: a small synchronous client for api:control-protocol plus a pytest plugin that exposes it as fixtures.

```yaml
api:
  package: packages/python (src layout, hatchling); PyPI name pgmem (available as of 2026-09-12); import pgmem; python >= 3.9; no runtime dependencies
  status: implemented 2026-09-12; tests in packages/python/tests run against a go build of cmd/pgmem (conftest builds it)
  core:
    - 'pgmem.start(database="postgres", user="postgres", params=None, log=False, binary=None) -> Pgmem; the process, context manager'
    - 'Pgmem: template -> Server (default); start_server(database, user="postgres", params=None) -> Server (op start); close() shuts the process down'
    - 'Server: id, dsn, host, port, user, database; snapshot(max_forks=None, timeout=30.0) -> Snapshot; close(); context manager'
    - 'Server.reset(snapshot=None, timeout=5.0): restore data in place, retaining endpoint and connections; optional named snapshot for dataset switching (requirement:live-http-test-reset)'
    - 'Snapshot: fork(timeout=None) -> Fork; close(); context manager'
    - 'Fork: dsn, host, port; close(); context manager'
    - 'Server.reset(snapshot=None, timeout=5.0): reset a fork to its origin or restore an explicit snapshot; keeps endpoint and pooled sockets'
    - 'binary lookup: argument, then PGMEM_BINARY env, then bundled pgmem/_bin/pgmem (policy:binary-distribution)'
  pytest_plugin:
    entry_point: pytest11 = pgmem.pytest_plugin
    cleanup: context managers preserve original and close failures in CleanupError (primary_error, cleanup_errors); process shutdown uses EOF with a bounded wait
    connection_guard: readonly_dsn(dsn); fixture option shared_read_only guards selected shared and class DSNs; reset-mode pgmem_shared_dsn remains writable
    fixtures:
      pgmem_options: session; kwargs for pgmem.start(); override to set database, params, log
      pgmem_process: session; the Pgmem process handle (not named pgmem, which would shadow the module)
      pgmem_server: session; pgmem_process.template; suites with several seed sets call start_server() in their own session fixtures
      pgmem_snapshot: session; default invokes pgmem_prepare then snapshots the template; advanced suites may override it directly
      pgmem_prepare: session; callback receives template Server; default snapshot fixture invokes it once, then snapshots automatically
      pgmem_fixture_options: session; isolation fork default, max_forks, fork_timeout 30s, snapshot_timeout 30s, reset_timeout 5s
      pgmem_fork / pgmem_dsn: function; fork then close in teardown (decision:fork-release-detection)
      pgmem_class_fork / pgmem_class_dsn: class; one fork shared by read-only tests in a class (parity with api:java-wrapper fork_scope CLASS)
      pgmem_test_dsn: function; isolated by default, @pytest.mark.pgmem(fork=False) selects the prepared template before application client fixtures connect (requirement:python-selective-isolation)
      pgmem_shared_fork / pgmem_shared_dsn: session; stable fork for persistent clients; @pytest.mark.pgmem(isolation="reset") resets after each marked serial test, including setup or test failure
      pgmem_live_fork: session; stable endpoint for a long-lived HTTP application per pytest worker
      pgmem_live_snapshots: session; name-to-snapshot mapping, default prepared snapshot
      pgmem_live_db: function; restore selected snapshot before sequential HTTP test, chosen by pgmem_dataset(name) marker
    marker: 'pgmem.shadow_pg (pytest marker/decorator): default fresh pgmem_fork; fork=False selects pgmem_shared_fork; prepare=True wraps pgmem_prepare fixture while it targets the template before snapshot; routes DATABASE_URL and PG* variables plus newly created psycopg 3, psycopg2 and asyncpg connections (requirement:declarative-test-target)'
    sqlalchemy: 'PostgreSQL dialects using patched drivers route to pgmem; pooled connections from earlier forks are invalidated on checkout; already checked-out connections and prebuilt asyncpg pools do not switch'
    example: |
      @pytest.fixture(scope="session")
      def pgmem_snapshot(pgmem_server):
          run_migrations(pgmem_server.dsn)
          return pgmem_server.snapshot()

      def test_orders(pgmem_dsn):
          with psycopg.connect(pgmem_dsn) as conn: ...
  returns_dsn_not_connection: driver neutral (psycopg, asyncpg, SQLAlchemy); a per-driver fixture builds on pgmem_dsn
  threads: one reader thread per process (rule:non-blocking-control-channel); pytest-xdist workers each get their own process
  wheel: hatch_build.py copies PGMEM_BINARY or runs go build with GOOS/GOARCH into pgmem/_bin, tags py3-none-<platform>, force-includes the gitignored binary; editable installs skip it (policy:binary-distribution)
  async: later; asyncio fixtures run the sync client in a thread for now
```
