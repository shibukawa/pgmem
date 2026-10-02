# Go browser E2E tests

Use the prepared `pgmemtest.Fixture` from [unit tests](unit.md). Call `fx.ShadowPG(t)` before launching the app with `exec.CommandContext`. A nil `Cmd.Env` inherits the fork's `DATABASE_URL` and libpq `PG*` variables. If setting `Cmd.Env` explicitly, build it from `os.Environ()` **after** `ShadowPG`. Give the app a distinct port, wait for its health endpoint, drive the browser, and stop/wait for the child at cleanup.

An already running child cannot be retargeted. A worker may keep one child and fork alive and reset sequentially, but must finish requests and transactions before each reset. Run separate package test processes or disjoint `go test -run ...` processes for parallel workers; each needs its own fork, app, port, and browser session. `t.Parallel()` is incompatible with `ShadowPG`.

Full example: `website/src/content/docs/guides/go/e2e-testing.mdx`.
