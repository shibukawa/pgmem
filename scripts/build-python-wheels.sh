#!/usr/bin/env bash
# Builds one wheel per platform into dist/wheels/, each carrying the binary
# from dist/<goos>-<goarch>/ (built on demand by scripts/build-binaries.sh).
# packages/python/hatch_build.py derives the platform tag from GOOS/GOARCH.
#
#   scripts/build-python-wheels.sh                 # every release target
#   scripts/build-python-wheels.sh darwin-arm64
set -euo pipefail
cd "$(dirname "$0")/.."
root=$PWD

all="darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64"
rm -rf dist/wheels
mkdir -p dist/wheels
cp LICENSE NOTICE packages/python/
for target in ${*:-$all}; do
  goos=${target%-*}
  exe=pgmem
  if [ "$goos" = windows ]; then exe=pgmem.exe; fi
  [ -f "dist/$target/$exe" ] || scripts/build-binaries.sh "$target"
  (cd packages/python && GOOS=$goos GOARCH=${target#*-} PGMEM_BINARY="$root/dist/$target/$exe" \
    uv build --wheel --out-dir "$root/dist/wheels")
done
rm -f packages/python/src/pgmem/_bin/pgmem packages/python/src/pgmem/_bin/pgmem.exe
ls -l dist/wheels
