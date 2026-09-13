#!/usr/bin/env bash
# Copies the binaries from dist/<goos>-<goarch>/ (built on demand by
# scripts/build-binaries.sh) into the npm platform packages and packs every
# package into dist/npm/. Publish the platform packages first, then
# @pgmem/core, whose optionalDependencies pin them. Pass platform names
# (darwin-arm64 ...) to fill and pack only those platform packages.
set -euo pipefail
cd "$(dirname "$0")/.."
root=$PWD

all="darwin-arm64 darwin-x64 linux-arm64 linux-x64 win32-arm64 win32-x64"
rm -rf dist/npm
mkdir -p dist/npm
for platform in ${*:-$all}; do
  os=${platform%%-*}
  cpu=${platform##*-}
  goos=$os
  if [ "$os" = win32 ]; then goos=windows; fi
  goarch=$cpu
  if [ "$cpu" = x64 ]; then goarch=amd64; fi
  exe=pgmem
  if [ "$goos" = windows ]; then exe=pgmem.exe; fi
  [ -f "dist/$goos-$goarch/$exe" ] || scripts/build-binaries.sh "$goos-$goarch"
  dir=packages/node/platforms/$platform
  rm -rf "$dir/bin"
  mkdir -p "$dir/bin"
  cp "dist/$goos-$goarch/$exe" "$dir/bin/$exe"
  chmod 755 "$dir/bin/$exe"
  cp LICENSE NOTICE "$dir/"
  (cd "$dir" && npm pack --silent --pack-destination "$root/dist/npm")
done
cp LICENSE NOTICE packages/node/core/
(cd packages/node/core && npm pack --silent --pack-destination "$root/dist/npm")
ls -l dist/npm
