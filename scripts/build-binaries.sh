#!/usr/bin/env bash
# Cross-compiles cmd/pgmem into dist/<goos>-<goarch>/pgmem (pgmem.exe on
# Windows). The module is pure Go, so any host builds every target. VERSION
# (X.Y.Z) is stamped into the binary as vX.Y.Z and reported in the ready
# event; without it the binary reports what the go command recorded.
#
#   scripts/build-binaries.sh                     # every release target
#   scripts/build-binaries.sh host                # this machine only
#   VERSION=0.1.0 scripts/build-binaries.sh linux-amd64 darwin-arm64
set -euo pipefail
cd "$(dirname "$0")/.."

all="darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64"
case "${1:-}" in
  "") targets=$all ;;
  host) targets="$(go env GOOS)-$(go env GOARCH)" ;;
  *) targets="$*" ;;
esac

ldflags="-s -w"
if [ -n "${VERSION:-}" ]; then
  ldflags="$ldflags -X main.version=v$VERSION"
fi
for target in $targets; do
  goos=${target%-*}
  goarch=${target#*-}
  exe=pgmem
  if [ "$goos" = windows ]; then exe=pgmem.exe; fi
  mkdir -p "dist/$target"
  CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -trimpath -ldflags="$ldflags" -o "dist/$target/$exe" ./cmd/pgmem
  echo "dist/$target/$exe"
done
