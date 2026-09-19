#!/bin/bash
# Regenerate the ahead-of-time compiled backend (internal/aot/pgaot) from
# wasm/out/postgres.wasm with the forked wasm2go, then the glue code.
#   WASM2GO=path   wasm2go binary to use instead of the pinned fork build
#   WASM2GO_SOURCE=path  local wasm2go fork source to build when WASM2GO is unset
#   AOT_WASM_OPT_LEVEL=Oz|O3|0  Binaryen optimization level for AOT input
#   AOT_WASM_OPT=path           wasm-opt binary to use
#   AOT_EXPORTS_FILE=path      newline-delimited AOT export allowlist
# The fork is pinned by wasm/wasm2go.lock: it is cloned into
# toolchain/wasm2go-src (gitignored) at the locked commit and built into
# toolchain/wasm2go-<commit>. Bump the lock to pick up a new fork commit.
set -euo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
ROOT=$(dirname "$HERE")
source "$HERE/wasm2go.lock"
WASM2GO_SOURCE=${WASM2GO_SOURCE:-}
if [ -n "$WASM2GO_SOURCE" ]; then
  if [ ! -d "$WASM2GO_SOURCE" ] || [ ! -f "$WASM2GO_SOURCE/go.mod" ]; then
    echo "error: wasm2go source is not a Go module: $WASM2GO_SOURCE" >&2
    exit 1
  fi
  WASM2GO=${WASM2GO:-$ROOT/toolchain/wasm2go-local}
  echo "== building wasm2go from $WASM2GO_SOURCE"
  (cd "$WASM2GO_SOURCE" && go build -o "$WASM2GO" ./cmd/wasm2go)
elif [ -z "${WASM2GO:-}" ]; then
  WASM2GO=$ROOT/toolchain/wasm2go-${commit:0:12}
fi
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

# Optimize a separate copy of the module before translating it.  The
# exnref-encoded module is for wazero and cannot be used here: wasm2go expects
# the original exception representation.  Keep names enabled because the
# generated Go names are part of the stable diff workflow.
AOT_WASM="$HERE/out/postgres.wasm"
AOT_WASM_OPT_FILE=""
if [ "${AOT_WASM_OPT_LEVEL:-Oz}" != "0" ]; then
  AOT_WASM_OPT=${AOT_WASM_OPT:-wasm-opt}
  if ! command -v "$AOT_WASM_OPT" >/dev/null 2>&1 && [ ! -x "$AOT_WASM_OPT" ]; then
    echo "error: wasm-opt not found; set AOT_WASM_OPT or AOT_WASM_OPT_LEVEL=0" >&2
    exit 1
  fi
  AOT_WASM_OPT_FILE=$(mktemp "$HERE/out/postgres.aot.XXXXXX.wasm")
  trap 'if [ -n "$AOT_WASM_OPT_FILE" ]; then rm -f -- "$AOT_WASM_OPT_FILE"; fi' EXIT
  "$AOT_WASM_OPT" \
    --enable-exception-handling \
    --enable-reference-types \
    --enable-bulk-memory \
    --enable-threads \
    --enable-sign-ext \
    --enable-mutable-globals \
    --enable-nontrapping-float-to-int \
    --enable-multivalue \
    --debuginfo \
    "-${AOT_WASM_OPT_LEVEL:-Oz}" \
    --dce \
    --remove-unused-module-elements \
    "$AOT_WASM" -o "$AOT_WASM_OPT_FILE"
  AOT_WASM="$AOT_WASM_OPT_FILE"
fi

# Default is wasm2go's pure-Go backend: on this workload it is both smaller
# (115 MB vs 281 MB) and faster than the asm backend (see README). ASM=1
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
EXPORTS_FILE=${AOT_EXPORTS_FILE:-$HERE/aot-exports.txt}
if [ ! -f "$EXPORTS_FILE" ]; then
  echo "error: AOT export allowlist not found: $EXPORTS_FILE" >&2
  exit 1
fi
ENTRY_EXPORTS=$(sed -e 's/[[:space:]]*#.*//' -e '/^[[:space:]]*$/d' "$EXPORTS_FILE" | paste -sd, -)
if [ -z "$ENTRY_EXPORTS" ]; then
  echo "error: AOT export allowlist is empty: $EXPORTS_FILE" >&2
  exit 1
fi
"$WASM2GO" $MODE -i "$AOT_WASM" -out-dir "$ROOT/internal/aot/pgaot" \
  -entry-exports="$ENTRY_EXPORTS" \
  -pkg pgaot -import github.com/shibukawa/pgmem/internal/aot/pgaot 2>&1 | grep -v -E 'fixpoint cap|slab=' || true
(cd "$ROOT/internal/aot" && python3 gen.py "$AOT_WASM" "$EXPORTS_FILE")

# wasm2go emits the linear-memory data segment as raw bytes.  Compress it in
# place after generation; data_compressed.go transparently expands it before
# any module constructor can copy the segments into guest memory.
python3 "$HERE/compress-aot-data.py" \
  "$ROOT/internal/aot/pgaot/data.bin" \
  "$ROOT/internal/aot/pgaot/data.bin.tmp"
mv "$ROOT/internal/aot/pgaot/data.bin.tmp" "$ROOT/internal/aot/pgaot/data.bin"
cp "$HERE/aot-data.go.txt" "$ROOT/internal/aot/pgaot/data_compressed.go"
gofmt -w "$ROOT/internal/aot"/*_gen.go
(cd "$ROOT" && go build ./internal/aot)
echo "== done: $(du -sh "$ROOT/internal/aot/pgaot" | cut -f1) of generated Go"
