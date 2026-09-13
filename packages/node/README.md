# npm packages

- `core/`: `@pgmem/core`, the launcher, the control-socket client and the
  test-runner entries (`@pgmem/core/register`,
  `@pgmem/core/jest-environment`). Pure JavaScript, no dependencies. See
  [core/README.md](core/README.md).
- `platforms/<platform>/` (darwin-arm64, darwin-x64, linux-arm64, linux-x64,
  win32-arm64, win32-x64): `@pgmem/<platform>` packages that contain only
  the binary. `@pgmem/core` lists them as optional dependencies, so the
  package manager installs the one matching the host; nothing runs at
  install time.

`scripts/build-npm.sh` cross-compiles `cmd/pgmem` into
`platforms/<platform>/bin/` (all platforms, or the ones named). Publish the
platform packages first, then `@pgmem/core`.

`npm test` in `core/` builds the binary for the host (or uses
`PGMEM_BINARY`) and runs the tests with `node --test`.
