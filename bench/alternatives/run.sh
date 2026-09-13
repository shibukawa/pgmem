#!/bin/bash
# Measure pgmem next to docker run, testcontainers-go and devbox.
#   RUNS=5 ./run.sh            -> results/raw.jsonl, results/sizes.json, results/summary.json
# Requires a running Docker daemon with postgres:18-alpine pulled, and devbox.
set -euo pipefail
cd "$(dirname "$0")"
RUNS=${RUNS:-5}
IMAGE=${IMAGE:-postgres:18-alpine}
OUT=results
mkdir -p "$OUT" bin
go build -o bin/alternatives .
go build -C ../.. -ldflags='-s -w' -o bench/alternatives/bin/pgmem ./cmd/pgmem
docker pull -q "$IMAGE" >/dev/null
devbox install -q -c devbox >/dev/null 2>&1

: > "$OUT/raw.jsonl"
for i in $(seq "$RUNS"); do
  for t in pgmem docker testcontainers devbox; do
    echo "run $i $t" >&2
    bin/alternatives -target "$t" -image "$IMAGE" >> "$OUT/raw.jsonl"
  done
  # the standalone binary that the Python, Java and Node.js wrappers spawn
  python3 - bin/pgmem >> "$OUT/raw.jsonl" <<'PY'
import json, subprocess, sys, time
t0 = time.perf_counter()
p = subprocess.Popen([sys.argv[1], "-database", "app"], stdin=subprocess.PIPE, stdout=subprocess.PIPE)
p.stdout.readline()
ms = (time.perf_counter() - t0) * 1000
rss = int(subprocess.check_output(["ps", "-o", "rss=", "-p", str(p.pid)]).strip()) / 1024
p.stdin.close(); p.wait(timeout=10)
print(json.dumps({"target": "binary", "startup_ms": ms, "mem_idle_mb": rss}))
PY
done

./sizes.sh "$IMAGE" > "$OUT/sizes.json"
python3 summarize.py "$OUT/raw.jsonl" "$OUT/sizes.json" > "$OUT/summary.json"
cp "$OUT/summary.json" ../../website/src/data/benchmarks.json
cat "$OUT/summary.json"
