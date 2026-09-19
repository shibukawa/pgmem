# pgmem

A real PostgreSQL 18 server that runs entirely inside your Go test process,
keeps everything in memory, and speaks the normal wire protocol on a loopback
port. No Docker, no external binaries, no files on disk. Works wherever Go
and wazero run (macOS, Linux, Windows; amd64 and arm64).

```go
func TestSomething(t *testing.T) {
    s, err := pgmem.Start(context.Background(), pgmem.Options{Database: "app"})
    if err != nil {
        t.Fatal(err)
    }
    defer s.Close()

    db, _ := sql.Open("pgx", s.DSN()) // any driver works, pools included
    // ...
}
```

### Per-test forks

Load the schema and seed data once, snapshot, and give every test its own
copy. A fork is a full backend on a copy of the data directory, so tests
can run in parallel; `MaxForks` bounds how many exist at once (by default a
quarter of the available memory divided by a fork's cost, 28 on a 7 GB CI
runner) and `Fork` blocks when the cap is reached.

```go
var fx *pgmemtest.Fixture

func TestMain(m *testing.M) {
    os.Exit(pgmemtest.Run(m, pgmemtest.Options{
        Prepare: func(ctx context.Context, db *sql.DB, dsn string) error {
            _, err := db.ExecContext(ctx, schema) // migrations, seed data
            return err
        },
    }, func(f *pgmemtest.Fixture) { fx = f }))
}

func TestOrders(t *testing.T) {
    t.Parallel()
    db := fx.DB(t) // fresh copy of the prepared database, closed when t ends
    // fx.PgxConn(t), fx.PgxPool(t), fx.DSN(t) and fx.Fork(t) also exist
}
```

The same primitives are available without the test helpers:
`s.Snapshot(ctx, opts)` checkpoints and copies the server's state, and
`snap.Fork(ctx)` starts a new server from it. Snapshotting takes about
10 ms and a fork about 20 ms.

A fork can also be put back in place: `srv.Reset(ctx)` returns it to its
snapshot while it keeps its port and its client connections, so a pool or
a connection string captured once (an ORM client built at import time,
say) stays valid from test to test. Pooled connections continue in a fresh
session; pgmem re-creates their named prepared statements and LISTEN
registrations. `srv.Restore(ctx, snap)` does the same with any snapshot of
the same database, and a reset with nothing to undo returns at once.

Connecting over loopback TCP works with every driver. For pgx you can
skip the kernel entirely with the in-process dialer, which makes small
queries about three times faster:

```go
cfg, _ := pgx.ParseConfig(s.DSN())
cfg.DialFunc = s.Dial // net.Pipe, no TCP
conn, _ := pgx.ConnectConfig(ctx, cfg)
// database/sql: stdlib.RegisterConnConfig(cfg) then sql.Open("pgx", name)
```

The module is pure Go: PostgreSQL was compiled to WebAssembly once and
translated to Go ahead of time, so there is no wasm runtime, no cgo and
nothing to download at run time.

| first `Start` in a process | ~0.1 s |
|---|---|
| further servers in the process | ~0.1 s |
| simple indexed `SELECT` via pgx over TCP | ~27 µs |
| same, in-process via `Server.Dial` | ~8.5 µs |
| sort + count over 200k generated rows | ~95 ms |
| what ships | 115 MB generated Go (~35 s to compile once), 1.4 MB of AOT data |
| minimal program, `-ldflags="-s -w"` | 36.7 MB |

Reference measurements and their conditions are in
[docs/benchmarks.md](docs/benchmarks.md).

Every server has its own filesystem and its own database state.

## From other languages

The same prepare-once, fork-per-test workflow is available without Go:

- **Python**: `pip install pgmem` and use the `pgmem_dsn` pytest fixture
  (override `pgmem_snapshot` to run migrations). See
  [packages/python](packages/python/README.md).
- **Java**: `io.github.shibukawa.pgmem:pgmem-junit5` injects a fresh
  `Fork` or `DataSource` into every test method, plus
  `io.github.shibukawa.pgmem:pgmem-native` with your platform's
  classifier. See [packages/java](packages/java/README.md).
- **Node.js**: `@pgmem/core` gives every test file its own copy through
  `DATABASE_URL` (a Vitest setup file, a Jest environment, `node --test
  --import`) and resets it in place between tests, so Prisma, Drizzle and
  TypeORM clients created at import time work unchanged. See
  [packages/node](packages/node/core/README.md).

All three bundle the `pgmem` binary (`go build ./cmd/pgmem`): a standalone
process that prints a JSON line with the DSN when ready, takes snapshot,
fork and reset requests as JSON lines on stdin (and, with `-control`, on a
loopback socket that test workers can use), and exits when its stdin is
closed.
[docs/subprocess.md](docs/subprocess.md) documents the protocol for other
languages.

## How it works

The server is PostgreSQL 18.3 from the [PGlite](https://pglite.dev) fork
(`electric-sql/postgres-pglite`), rebuilt as a self-contained WebAssembly
module and translated to Go with wasm2go. The wasm file is a build
intermediate; only the generated Go is in the module. (The initdb program
is still executed as wasm, under [wazero](https://wazero.io), but only by
`cmd/pgmem-mkdata` when regenerating the embedded data directory.)

- `wasm/build.sh` builds the module with Emscripten. Unlike the PGlite
  distribution it is linked statically: no dynamic linking, no JS-based
  longjmp, loadable modules (plpgsql, snowball dictionaries, encoding
  conversions) are linked in and resolved through a static `dlopen`
  replacement (`wasm/pgmem_dl.c`, table generated by `wasm/gen_modules.py`).
- `internal/vfs` is an in-memory POSIX-ish filesystem. The data directory,
  WAL and share files live there.
- `internal/host` implements the Emscripten `env` and WASI imports the
  module needs (syscalls, time, memory growth, exit unwinding) as an
  engine-agnostic table. A few PostgreSQL internals are routed to Go on
  purpose: all entropy (`/dev/urandom`, `getentropy`, and UUID generation)
  comes from `crypto/rand`, WAL CRC-32C is computed with `hash/crc32`,
  MD5 and SHA-1/SHA-2 (`md5()`, `sha256()` ..., SCRAM's HMAC) run on Go's
  `crypto/*` through a host-side replacement of `src/common/cryptohash.c`,
  and `poll()` sleeps on the host so `pg_sleep` and timeouts do not spin.
  There is no OpenSSL: pgcrypto's OpenSSL layer (`openssl.c`,
  `pgp-mpi-openssl.c`) is replaced by `wasm/pgmem_pgcrypto_*.inc`, which
  hand digests, the symmetric ciphers (AES, Blowfish, DES, 3DES, CAST5 in
  ECB/CBC/CFB) and the OpenPGP RSA/ElGamal arithmetic to Go
  (`internal/host/crypto.go`: `crypto/*`, `golang.org/x/crypto`,
  `math/big`), and `pgp-compress.c` by `wasm/pgmem_pgcrypto_compress.inc`,
  which does OpenPGP ZIP/ZLIB compression on `compress/flate` and
  `compress/zlib`. Its own `crypt()` algorithms and the OpenPGP packet
  code are the upstream C. There is no TLS code.
- Contrib extensions are linked in statically like plpgsql
  (`CONTRIB_MODULES` in `wasm/build.sh`); `CREATE EXTENSION pgcrypto` works
  out of the box. `testdata/regress` replays every bundled extension's
  upstream regression suite against it (`TestContribRegress`). Adding one
  is scripted: `skills/add-pgmem-extension/add-extension.sh` (the
  `add-pgmem-extension` skill, also installable with `npx skills add
  shibukawa/pgmem --skill add-pgmem-extension`).
- `internal/aot` binds that table to the generated Go code;
  `internal/wzr` binds it to wazero for the initdb step of `pgmem-mkdata`.
- `internal/engine` drives initdb and runs the cluster: PostgreSQL is
  built with `EXEC_BACKEND` (the Windows way of starting children, a new
  program image per process that re-attaches to the shared memory), so
  every process is a fresh module instance on its own goroutine.
  `internal/host/cluster.go` is the kernel they share: pids and signals
  (delivered when the target next enters a host call, where a kernel would
  interrupt it too), System V shared memory (a file per segment, mapped
  with `MAP_FIXED` at the same address into every instance's reserved
  linear memory; on Windows a placeholder reservation whose pieces are
  replaced by `MapViewOfFile3` views, which needs Windows 10 1803 or
  later), POSIX semaphores, and the sockets that carry client
  connections to the postmaster. The module is compiled with `-matomics`
  so spinlocks and `pg_atomic_*` are real atomic instructions. Every
  instance's memory starts as a copy-on-write view of one shared image of
  the module's data segments, so a new process costs no copying.
  The setup step that creates the database and user still runs a
  throwaway `postgres --single` child the way initdb does.
- `pgmem.go` bridges TCP connections to that backend.

Startup is fast because `internal/pgdata/pgdata.tar.zst` contains a data
directory produced by initdb at build time (`go run ./cmd/pgmem-mkdata`),
compressed with zstd (pure Go, `klauspost/compress`).

pgvector is not part of the PostgreSQL tree: `wasm/pgvector.lock` pins
its release and checksum, and `build.sh` compiles it against the
installed server headers like a PGXS build would (without
`-march=native`).

### Ahead-of-time backend

`internal/aot/pgaot` is the same wasm module translated to Go by
[wasm2go](https://github.com/goccy/wasm2go). The pinned
[shibukawa/wasm2go-fork](https://github.com/shibukawa/wasm2go-fork) commit in
`wasm/wasm2go.lock` is used by default. Set `WASM2GO_SOURCE` to a checked-out
local fork, for example `/path/to/wasm2go-fork`, then run
`devbox run -- ./wasm/gen-aot.sh`. `internal/aot`
plugs the generated package's import interfaces into the same host table.
Regenerate with `./wasm/gen-aot.sh` after rebuilding the wasm.

wasm2go has two backends. The fork ([shibukawa/wasm2go-fork](https://github.com/shibukawa/wasm2go-fork),
branch `pgmem`, pinned by `wasm/wasm2go.lock`) makes both work on this module (it
accepts Go 1.27's `-S` listing format and falls back per function on
jump-table shapes the transform does not understand). Measured here the
pure-Go backend wins on both size and speed, so it is the default:

| backend | generated source | sort + count 200k rows | simple `SELECT` |
|---|---|---|---|
| pure Go (`gen-aot.sh`) | 115 MB | 95 ms | 26.5 µs |
| asm (`ASM=1 gen-aot.sh`) | 281 MB | 122 ms | 27.9 µs |

`wasm/aot-exports.txt` is the single allowlist for the supported host-facing
AOT API. It is used both for wasm2go's `-entry-exports` roots and for the Go
dispatcher, so the generated dispatcher has 25 cases instead of all 2,070
function exports in the unpruned wasm. The current `postgres.wasm` contains
34 export entries in total, including the memory/table and Emscripten runtime
exports required by the module. With `wasm-opt -Oz`, compressed AOT data, and
the limited dispatcher, the stripped darwin/arm64 binary is 36.7 MB.

The release build uses `-trimpath -ldflags='-s -w'`: DWARF is removed, while
Go's compact pclntab is retained for panic and traceback support. Disabling
inlining was measured as a line-table reduction but made the final binary
larger, so no separate line-information compression flag is enabled.

The generated functions are named after PostgreSQL's symbols
(`F_heap_insert`), spread over the `p0`..`p5` packages by name hash, and
written to files named after their subject (`heap.go`, `relation.go`,
`pg_stat.go`; a huge function gets its own file), so rebuilding the wasm
after a source change only rewrites the Go of the functions that changed
and the diff is readable file by file. The fork's `-symbol-names` and
`-group-files` flags do this; with the stock index-based names
(`Fn3908`) one added import renumbers every function and the whole tree
shows up in the diff.

The asm backend's register allocator seems to lose to gc's own optimizer
on this memory-bound code (arm64; amd64 not measured).

## Limits

- **Process model.** PostgreSQL runs as it does on a server: a
  postmaster starts a backend process for every connection plus its
  auxiliary processes (checkpointer, background writer, WAL writer,
  autovacuum launcher), and they share memory. Every process is a module
  instance on its own goroutine; the shared memory is a segment mapped at
  the same address into each of them. So sessions are real: row and
  advisory locks make other sessions wait, deadlocks are detected
  (SQLSTATE 40P01 after `deadlock_timeout`), `pg_terminate_backend`,
  `pg_stat_activity` and `pg_locks` show the other sessions, and session
  state is per connection. A `Restore`/`Reset` stops the cluster and starts
  it on the copy; client connections keep their sockets and get a new
  backend on their next message, with their named prepared statements and
  `LISTEN` registrations re-created (a backend PostgreSQL itself ends, with
  a FATAL, ends the connection as it would anywhere). Query cancellation of
  a statement that never makes a system call waits until it does.
- `Close` does not drop client connections: each stays open until its
  client closes it, sends a message (answered with SQLSTATE 57P01) or 30 s
  pass, because pools such as node-postgres's raise an unhandled error
  when an idle connection is dropped under them.
- Every database of a server can be connected to (`CREATE DATABASE`
  works). A database that does not exist is refused with SQLSTATE 3D000.
- Extensions: plpgsql, pgcrypto, citext, pg_trgm, hstore, ltree, btree_gist, btree_gin, unaccent, tablefunc, intarray, fuzzystrmatch, cube, earthdistance, seg, bloom, isn, dict_int, dict_xsyn, lo, tsm_system_rows, tsm_system_time, pgstattuple, uuid-ossp, amcheck, pg_visibility, pageinspect, pg_buffercache, pg_freespacemap, pg_prewarm, pg_stat_statements, auto_explain (a `LOAD`-able module rather than an extension), and pgvector 0.8.6 as `vector`. ICU, OpenSSL and zlib are not compiled
  in; pgcrypto gets its crypto and its OpenPGP compression from Go instead
  (so `compress-algo=1|2` and messages made by GnuPG work), and
  `fips_mode()` is always false.
- Parallel query and parallel index builds work: the workers are
  processes of the cluster like any other. `io_method` defaults to `sync`
  (PostgreSQL 18's `worker` would add three I/O worker processes for no
  gain in memory).
- `io_method` is forced to `sync`. PGlite runs the backend as if under a
  postmaster, so PostgreSQL 18's default `worker` method would hand
  batched reads to IO workers that do not exist; with an in-memory
  filesystem there is nothing to overlap anyway.
- A running server is about 150 MB resident (65 MB PostgreSQL memory with
  the default `shared_buffers=32MB`, 40 MB of in-memory data directory,
  the rest binary and Go heap). `Options.Params` can raise
  `shared_buffers` when a test needs a larger cache.
- `statement_timeout` and friends fire between protocol messages and
  whenever the backend reads the clock or sleeps (so they interrupt
  `pg_sleep`), but not inside a pure CPU loop that never does either.
- 32-bit wasm: a server can address at most 4 GiB.

## Rebuilding the wasm module

Requires bash, python3, curl, binaryen (`wasm-opt`), and Emscripten. The
pinned emsdk can be checked out at `toolchain/emsdk` with version 3.1.74
installed and activated, or a Devbox environment can provide `emcc` and
`emmake` directly on `PATH`. The PostgreSQL source (the PGlite fork, pinned
by commit and sha256 in `wasm/postgres-pglite.lock`) is downloaded by the
build script into `wasm/out/src` and patched there by `wasm/patches.py`;
nothing under version control is modified.

The repository also includes a reproducible Devbox environment. Its lockfile
currently uses Emscripten 6.0.8, Binaryen, Bison, Flex, and GNU Make:

```bash
devbox install
devbox run -- ./wasm/build.sh
devbox run -- ./wasm/gen-aot.sh
devbox run -- go run ./cmd/pgmem-mkdata
devbox run -- go test . ./cmd/... ./internal/...
```

The checked-in generated tree and size figures above were regenerated with
this Devbox environment on 2026-09-17.

The equivalent commands without Devbox are:

```bash
git clone --depth 1 https://github.com/emscripten-core/emsdk.git toolchain/emsdk
(cd toolchain/emsdk && ./emsdk install 3.1.74 && ./emsdk activate 3.1.74)
./wasm/build.sh                  # downloads + patches the source, writes wasm/out/*.wasm and internal/assets/share.tar.gz
./wasm/gen-aot.sh                # optimizes + regenerates internal/aot/pgaot from postgres.wasm
go run ./cmd/pgmem-mkdata        # runs wasm/out/initdb.exnref.wasm under wazero, writes internal/pgdata/pgdata.tar.zst
go test . ./cmd/... ./internal/...
```

The AOT step runs `wasm-opt -Oz` before wasm2go and gzip-compresses the
generated linear-memory data segment. `AOT_WASM_OPT_LEVEL=O3` selects the
speed-oriented Binaryen passes; `AOT_WASM_OPT_LEVEL=0` skips the Wasm
optimization for a baseline comparison. The compressed data is expanded once
when the generated package is initialized.

The build emits legacy wasm exception handling (what wasm2go consumes for
setjmp/longjmp) and also an `exnref`-encoded copy via `wasm-opt`, which
is what wazero needs for the initdb step.

## Where the time goes

Profiling the ahead-of-time build (`docs/benchmarks.md` has the numbers)
shows that a small indexed `SELECT` over TCP spends about 5% of its time in
PostgreSQL and the rest in the kernel socket path, the poller and goroutine
wake-ups; hence the in-process dialer. An engine-bound query (sort of 200k
rows) is entirely PostgreSQL proper (expression interpreter, `qsort`,
string comparison); libc-level string routines are a few percent, and
`memcpy`/`memset` are already Go `copy`/`clear`, so there is no
SIMD-in-Go lever there.

## Development

`make vet` vets and formats everything except the generated package;
`make test` and `make bench` run the suite and the benchmarks.

## Releasing

Every artifact of a release carries one version, taken from a `vX.Y.Z`
tag (the numbering is explained on the
[versioning](https://shibukawa.github.io/pgmem/versioning/) page; the
0.1.x releases bring the pipeline up before 1.18.0). Pushing the tag runs
`.github/workflows/release.yml`. It cross-compiles `cmd/pgmem` for five
platforms (`scripts/build-binaries.sh`), stamps the version into every
manifest (`scripts/set-version.sh`), packs the archives, npm packages,
wheels and Maven bundle (`scripts/package-archives.sh`,
`scripts/build-npm.sh`, `scripts/build-python-wheels.sh`,
`scripts/build-maven-bundle.sh`) and publishes them to GitHub Releases,
npm, PyPI and Maven Central. It also fetches the tag through
proxy.golang.org so that pkg.go.dev lists it. CI runs the same packaging
scripts on every push with placeholder binaries and a throwaway signing
key, and `.github/workflows/docs.yml` deploys `website/` to GitHub Pages.

Rehearse on the commit you are going to tag. Nothing is published: npm
runs `publish --dry-run`, and the Central Publisher Portal validates the
Maven bundle, which is then dropped.

```bash
gh workflow run release.yml -f version=0.1.0
```

The registries need a one-time setup:

- **PyPI**: a pending trusted publisher for the project `pgmem` (owner
  `shibukawa`, repository `pgmem`, workflow `release.yml`, environment
  `release`). It works before the first upload.
- **npm**: a trusted publisher can only be added to a package that
  exists, so the first version of the five `@pgmem/<platform>` packages
  and of `@pgmem/core` is published by hand, with 2FA, from the tarballs
  of a rehearsal (`gh run download <run-id> -n packages -D dist`, then
  `npm publish dist/npm/<tarball> --access public`, `pgmem-core` last).
  Each package then gets the trusted publisher `shibukawa/pgmem`,
  `release.yml`, environment `release`, with direct publishing allowed
  (package settings on npmjs.com, or `npm trust github`). The workflow
  skips versions that are already on npm.
- **Maven Central**: the namespace `io.github.shibukawa`, which the
  Central Publisher Portal verifies when you sign in with the GitHub
  account `shibukawa` (the artifacts are `io.github.shibukawa.pgmem:*`),
  a Portal user token as `CENTRAL_USERNAME` / `CENTRAL_PASSWORD`, and an
  armored GPG secret key as `GPG_PRIVATE_KEY` / `GPG_PASSPHRASE` whose
  public key is on keys.openpgp.org or keyserver.ubuntu.com, all four as
  secrets of the `release` environment.

To publish a tagged version to some registries again, for instance after
a registry-side failure, run the workflow on the tag:

```bash
gh workflow run release.yml --ref v0.1.0 -f version=0.1.0 -f registries=maven -f dry_run=false
```

## Debugging

- `Options.Log` receives the server log.
- `PGMEM_TRACE=1` prints every host call (syscalls etc.) to stderr.
