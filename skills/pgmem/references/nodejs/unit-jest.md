# Jest setup

Set `globalSetup`, `globalTeardown`, and `testEnvironment: '@pgmem/core/jest-environment'` in Jest configuration. Global setup starts `PgmemServer`, exports `server.env()` to workers, and stores the server for teardown; teardown closes it. Each Jest worker has a fork and resets between files it runs. `jest-environment-node` must be available. Within a file, isolate writing cases and keep shared Pool use sequential. See [unit lifecycle](unit.md) and `website/src/content/docs/guides/nodejs/testing.mdx`.
