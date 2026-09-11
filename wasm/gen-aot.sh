#!/bin/bash
# Regenerate the ahead-of-time compiled backend (internal/aot/pgaot) from
# wasm/out/postgres.wasm with the forked wasm2go, then the glue code.
#   WASM2GO=path   wasm2go binary to use instead of the pinned fork build
# The fork is pinned by wasm/wasm2go.lock: it is cloned into
# toolchain/wasm2go-src (gitignored) at the locked commit and built into
# toolchain/wasm2go-<commit>. Bump the lock to pick up a new fork commit.
set -euo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
ROOT=$(dirname "$HERE")
source "$HERE/wasm2go.lock"
WASM2GO=${WASM2GO:-$ROOT/toolchain/wasm2go-${commit:0:12}}
if [ ! -x "$WASM2GO" ]; then
  SRC=$ROOT/toolchain/wasm2go-src
  if [ ! -d "$SRC/.git" ]; then
    echo "== cloning $repo"
    mkdir -p "$ROOT/toolchain"
    git clone -q --branch "$branch" "$repo" "$SRC"
  fi
  if ! git -C "$SRC" cat-file -e "$commit^{commit}" 2>/dev/null; then
    git -C "$SRC" fetch -q origin "$branch"
  fi
  git -C "$SRC" checkout -q --detach "$commit"
  echo "== building wasm2go at $commit"
  (cd "$SRC" && go build -o "$WASM2GO" ./cmd/wasm2go)
fi
rm -rf "$ROOT/internal/aot/pgaot"
mkdir -p "$ROOT/internal/aot/pgaot"
# Default is wasm2go's pure-Go backend: on this workload it is both smaller
# (102 MB vs 281 MB) and faster than the asm backend (see README). ASM=1
# selects the asm backend (amd64.s + arm64.s + pure fallback).
# -symbol-names names the generated functions after PostgreSQL's symbols
# (F_heap_insert, not Fn3908) and assigns them to the p0..p5 packages by
# name hash, so a rebuilt wasm only changes the Go of the functions that
# actually changed instead of renumbering and repacking all of them.
# -group-files then writes each package as files named after the
# functions' subject (heap.go, relation.go, pg_stat.go; huge functions
# alone) instead of one 20 MB pN.go, so diffs stay reviewable.
# Pure-Go backend only; -chunks pins the package count (6 = what the
# module's ~5.7 MB of code derives to) so growth never reshuffles it.
# -addr-consts names the constants that point into the static data, so
# adding a string literal no longer rewrites every body that carries an
# address (see the fork's README_ABOUT_FORK.md).
MODE="-pure -symbol-names -chunks 6 -group-files -addr-consts"
if [ "${ASM:-0}" = "1" ]; then MODE=""; fi
"$WASM2GO" $MODE -i "$HERE/out/postgres.wasm" -out-dir "$ROOT/internal/aot/pgaot" \
  -pkg pgaot -import github.com/shibukawa/pgmem/internal/aot/pgaot 2>&1 | grep -v -E 'fixpoint cap|slab=' || true
(cd "$ROOT/internal/aot" && python3 gen.py "$HERE/out/postgres.wasm")
gofmt -w "$ROOT/internal/aot"/*_gen.go
(cd "$ROOT" && go build ./internal/aot)
echo "== done: $(du -sh "$ROOT/internal/aot/pgaot" | cut -f1) of generated Go"
