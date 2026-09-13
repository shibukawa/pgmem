#!/usr/bin/env bash
# Packs the binaries in dist/<goos>-<goarch>/ into the GitHub Release assets
# dist/release/pgmem-<version>-<goos>-<goarch>.tar.gz (.zip for Windows),
# each with LICENSE and NOTICE, and writes dist/release/SHA256SUMS.
#
#   VERSION=0.1.0 scripts/package-archives.sh
set -euo pipefail
cd "$(dirname "$0")/.."
root=$PWD
v=${VERSION:?set VERSION to X.Y.Z}
export COPYFILE_DISABLE=1 # no AppleDouble files in tarballs made on macOS

all="darwin-amd64 darwin-arm64 linux-amd64 linux-arm64 windows-amd64 windows-arm64"
rm -rf dist/release
mkdir -p dist/release
for target in $all; do
  name=pgmem-$v-$target
  stage=$(mktemp -d)
  mkdir "$stage/$name"
  cp LICENSE NOTICE "$stage/$name/"
  if [ "${target%-*}" = windows ]; then
    cp "dist/$target/pgmem.exe" "$stage/$name/"
    (cd "$stage" && zip -qr "$root/dist/release/$name.zip" "$name")
  else
    cp "dist/$target/pgmem" "$stage/$name/"
    chmod 755 "$stage/$name/pgmem"
    tar -C "$stage" -czf "dist/release/$name.tar.gz" "$name"
  fi
  rm -rf "$stage"
done
if command -v sha256sum >/dev/null; then sum=sha256sum; else sum="shasum -a 256"; fi
(cd dist/release && $sum pgmem-* > SHA256SUMS)
cat dist/release/SHA256SUMS
