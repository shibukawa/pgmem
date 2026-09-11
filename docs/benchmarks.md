# Backend benchmarks (reference values)

Measured 2026-09-11 on an Apple M-series laptop (arm64, 8 cores), Go 1.27,
wasm2go fork [shibukawa/wasm2go-fork](https://github.com/shibukawa/wasm2go-fork) branch `pgmem` (pinned in `wasm/wasm2go.lock`), Emscripten 3.1.74, PostgreSQL
18.3 (PGlite fork). Numbers are the representative value of three
`-count=3` runs; each run stayed within about 1% of the others. Treat them
as an indication of relative cost, not as absolute figures: amd64 was not
measured, and the simple-query case is dominated by the TCP round trip
and protocol handling in pgx and the bridge, not by the database engine.

Commands: `make bench`; `ASM=1 ./wasm/gen-aot.sh` / `./wasm/gen-aot.sh`
switch the generated backend between asm and pure Go. The wazero rows
are historical: the wazero backend was removed from the library once the
generated code became the only backend (wazero now only runs initdb in
`pgmem-mkdata`).

## Query benchmarks

`BenchmarkSimpleQueries`: `SELECT v FROM bench WHERE id = $1` over a
1000-row table with a primary key, one pgx `QueryRow` per iteration.

`BenchmarkCPUHeavyQuery`: `SELECT count(*) FROM (SELECT md5(i::text) AS h
FROM generate_series(1, 200000) i ORDER BY h) t`; almost all time is spent
inside the engine (hashing, sorting).

| backend | simple SELECT (TCP) | simple SELECT (in-process `Dial`) | sort + count 200k rows |
|---|---|---|---|
| wazero (compiler engine, exnref EH) | 73.3 µs | | 504 ms |
| AOT, wasm2go pure-Go backend | 26.5 µs | 8.5 µs | 95 ms |
| AOT, wasm2go asm backend (arm64.s) | 27.9 µs | | 122 ms |

Relative to wazero the AOT pure-Go build is about 2.8x faster on the
round-trip-bound query and about 5.3x faster on the engine-bound query.
The asm backend is about 28% slower than pure Go on the engine-bound
query. A plausible reading is that gc's own optimizer produces better
code for this memory-access-heavy C-derived code than wasm2go's
assembly register allocator; that is a hypothesis, not a measurement.

## Profile (AOT pure Go, `go test -cpuprofile`; the generated functions carry PostgreSQL's symbol names)

Simple SELECT over TCP, whole process (client and server in one
process): 45% raw syscalls, 24% `kevent`, 16% goroutine wake-ups
(`pthread_cond_*`), 4% `madvise`; `PostgresMainLoopOnce` and everything
under it is 5.3%. The in-process dialer removes the syscall and poller
share.

Sort + count 200k rows, inside `Backend.Exec`: `ExecInterpExpr` 21% cum,
`qsort_tuple_unsigned` 11% flat, `varstrfastcmp_c` + `varstr_abbrev_convert`
about 5%, `crypto/md5` (host) 4%, `heap_form_minimal_tuple`,
`tuplestore_puttuple_common`, `MemoryContextReset` a few percent each.
No libc string routine shows up above 1%.

Linear memory: the module now starts with a 32 MB memory (Emscripten
`INITIAL_MEMORY`) and grows on demand. On the AOT backend the memory is
an anonymous `mmap` (`VirtualAlloc` on Windows) of the full 2 GiB range
outside the Go heap: pages materialize on first touch, growth is
bookkeeping, and `Close` unmaps at once. Keeping it on the Go heap was
measured to be the dominant cost of a running server (reused spans are
zero-filled, freed instances linger until the GC scavenges).

## Resident memory

`pgmem` binary, darwin/arm64, RSS after startup (`ps`), default
`shared_buffers=32MB`:

| | before | after |
|---|---|---|
| default server | 248 MB | 152 MB |
| with `-database app -user tester` (setup child) | 525 MB | 161 MB |
| `-params shared_buffers=128MB` (initdb's default) | 352 MB | ~230 MB |

Breakdown of the 152 MB: linear memory 65 MB (touched), VFS contents 39 MB
(the initdb output uncompressed: three catalog copies and a 16 MB WAL
segment), Go heap for the rest of the VFS structure and the decompressed
share tree, and the binary's own pages.

## Startup

| | wazero | AOT |
|---|---|---|
| first `Start` in a process, cold (no compilation cache) | ~2.5 s | ~0.13 s |
| first `Start` in a process, compilation cache warm | ~0.3 s | ~0.13 s |
| further `Start` in the same process | ~0.2 s | ~0.1 s |

wazero's cost is compiling the 9.6 MB module; the cache lives under the
user cache directory (`pgmem/wazero`). The AOT figure is the cost of
unpacking the embedded share tree and data directory into the in-memory
filesystem plus backend boot.

## Binary size

A minimal program that starts one server and prints its DSN, built with
`go build -ldflags="-s -w"` on darwin/arm64:

| backend | stripped binary | of which embedded assets |
|---|---|---|
| AOT pure Go (default) | 36.2 MB | 1.2 MB pgdata (zstd) + 0.3 MB share |
| wazero (`-tags pgmem_wazero`) | 16.0 MB | 9.1 MB wasm + 1.2 MB pgdata + 0.3 MB share |

The default build does not embed the wasm module. The initdb output is a
zstd-compressed tar (1.2 MB; it was 4.3 MB as gzip, because the three
catalog copies in template0/template1/postgres compress well against
each other only with a large window).

## Size and build cost

| backend | source in the repository | compile once (`go build`) |
|---|---|---|
| wazero | 9.6 MB wasm + 1.2 MB pgdata | seconds |
| AOT pure Go | 104 MB Go | ~35 s |
| AOT asm | 281 MB (84 MB amd64.s + 84 MB arm64.s + pure fallback) | not measured separately |

Generation time with the fork: pure Go about 25 s, asm about 3.3 min
(includes wasm2go's own `go build -gcflags=-S` capture of every package
for both architectures).

## wasm2go fork notes

The fork lives at [shibukawa/wasm2go-fork](https://github.com/shibukawa/wasm2go-fork),
branch `pgmem` on top of upstream `main`; `wasm/wasm2go.lock` pins the
commit gen-aot.sh builds. The asm backend needed two fixes to get through
this module:

- Go 1.27 prints data symbols in `-S` listings as `... size=N align=0xM`;
  the capture regex rejected the suffix, so no jump table was captured
  and the bundle aborted with "jump table ... not captured".
- Jump-table shapes the transform does not recognize (for example a spill
  scheduled between the table `LEAQ` and the indirect `JMP`) now fall
  back to the pure-Go body for that function only, instead of aborting.
  With the regex fix in place this module needed zero such fallbacks.

wasm2go's own test suite fails on this machine at the unmodified upstream
commit as well (Go 1.27 environment), so those failures are unrelated to
the fork changes; the gcasm package tests, including the new regression
test for the listing format, pass.

The fork also adds `-symbol-names` (with `-chunks N`), which pgmem uses
for the pure-Go backend. Stock wasm2go names every function `Fn<wasm
function index>` and packs functions into the chunk packages by
first-fit-decreasing bin packing on body size. Both are unstable under
small source changes: commit d71bf05 added one host import
(`pgmem_listen`), every defined function index moved by one, and the
generated tree showed a 200,000-line diff for no code change. With the
flag the parser reads the wasm name section (`--profiling-funcs` at link
time keeps it), functions are named `F_<symbol>` (duplicate static names
get a 1-based ordinal), and each function's chunk is `fnv32a(name) mod
N`, so a function keeps its name and its package for as long as its
symbol exists. Binaryen's own uniquifier for repeated static names
(`heap_getattr_1883`, the number being a pre-optimization index that
shifts too) is undone by the same rule.

Measured on this module by linking one extra object that adds one host
import and one exported function, then regenerating both ways
(`git diff --numstat` over the generated tree, 5.9 M lines):

| naming | files changed | lines changed |
|---|---|---|
| `Fn<index>` (stock) | 21 | 394,052 |
| `-symbol-names -chunks 6` | 4 | 10 |

The ten lines are the new import's interface method and glue, and the
new function's body. The mode is pure-Go only: the asm bundle derives
`Fn<index>` symbols itself.

`-group-files` (fork release 2) then splits each package into files
named after the functions' subject — the first non-verb token of the
symbol (`heap.go`, `relation.go`, `pg_stat.go`), small groups pooled
into `misc_<letters>.go`, a function over 128 KiB alone in
`<group>_<name>.go` — instead of one 13-21 MB `pN.go`. The module comes
out as about 1,800 files, most under 100 KB, so a diff is reviewable
file by file. Every split derives from names, so a rebuild keeps the
files it does not touch.
