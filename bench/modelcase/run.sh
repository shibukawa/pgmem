#!/bin/bash
# Run the model-case suite on pgmem and on conventional servers.
#   RUNS=5 ./run.sh                 -> results/raw.jsonl, results/summary.json,
#                                      website/src/data/modelcase.json
#   TARGETS="pgmem devbox" ./run.sh -> skip docker
# devbox needs devbox; docker needs a running daemon with the image pulled
# (a target whose tool is missing is skipped with a note).
set -euo pipefail
cd "$(dirname "$0")"
RUNS=${RUNS:-5}
TARGETS=${TARGETS:-"pgmem devbox docker"}
IMAGE=${IMAGE:-postgres:18-alpine}
OUT=results
mkdir -p "$OUT" bin
go vet ./...
go test -c -o bin/store.test ./store
go build -o bin/driver ./cmd/driver

: > "$OUT/raw.jsonl"
for t in $TARGETS; do
  case $t in
    docker)
      if ! docker info >/dev/null 2>&1; then echo "skip docker: no daemon" >&2; continue; fi
      docker pull -q "$IMAGE" >/dev/null
      ;;
    devbox)
      if ! command -v devbox >/dev/null; then echo "skip devbox: not installed" >&2; continue; fi
      devbox install -q -c ../alternatives/devbox >/dev/null 2>&1
      ;;
  esac
  bin/driver -backend "$t" -runs "$RUNS" -image "$IMAGE" >> "$OUT/raw.jsonl"
done

python3 summarize.py "$OUT/raw.jsonl" > "$OUT/summary.json"
cp "$OUT/summary.json" ../../website/src/data/modelcase.json
