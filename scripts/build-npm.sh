#!/usr/bin/env bash
# Cross-compiles cmd/pgmem for every npm platform package and puts the
# binary into packages/node/platforms/<platform>/bin/. Publish the platform
# packages first, then @pgmem/core. Pass platform names (darwin-arm64 ...)
# to build only those.
set -euo pipefail
cd "$(dirname "$0")/.."
all="darwin-arm64 darwin-x64 linux-arm64 linux-x64 win32-arm64 win32-x64"
for platform in ${*:-$all}; do
  os=${platform%%-*}
  cpu=${platform##*-}
  goos=$os; [ "$os" = win32 ] && goos=windows
  goarch=$cpu; [ "$cpu" = x64 ] && goarch=amd64
  exe=pgmem; [ "$goos" = windows ] && exe=pgmem.exe
  out="packages/node/platforms/$platform/bin/$exe"
  mkdir -p "$(dirname "$out")"
  CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -trimpath -ldflags="-s -w" -o "$out" ./cmd/pgmem
  echo "$out"
done
