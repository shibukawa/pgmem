#!/bin/bash
# Download and link sizes, printed as JSON (bytes).
set -euo pipefail
cd "$(dirname "$0")"
IMAGE=${1:-postgres:18-alpine}
ARCH=$(go env GOARCH)

layers() { # image -> compressed layer bytes for linux/$ARCH
  docker manifest inspect -v "$1" | python3 -c "
import json, sys
d = json.load(sys.stdin); d = d if isinstance(d, list) else [d]
for m in d:
    p = m.get('Descriptor', {}).get('platform', {})
    if p.get('os') == 'linux' and p.get('architecture') == '$ARCH':
        print(sum(l['size'] for l in m.get('OCIManifest', m.get('SchemaV2Manifest', {})).get('layers', [])))
        break"
}
image_disk() { docker image inspect -f '{{.Size}}' "$1"; }

RYUK=$(docker images --format '{{.Repository}}:{{.Tag}}' | grep '^testcontainers/ryuk:' | sort -V | tail -1)

PG=$(devbox run -q -c devbox -- bash -c 'readlink -f "$(command -v postgres)"' 2>/dev/null | grep '^/nix' | tail -1)
STORE=$(echo "$PG" | cut -d/ -f1-4)
NIX=$(nix --extra-experimental-features nix-command path-info -r --json --json-format 1 --store https://cache.nixos.org "$STORE" 2>/dev/null | python3 -c "
import json, sys
d = json.load(sys.stdin); v = [x for x in (d.values() if isinstance(d, dict) else d) if x]
print(sum(x.get('downloadSize', 0) for x in v), sum(x.get('narSize', 0) for x in v))")

(cd linksize && go mod tidy >/dev/null 2>&1 && go build -o ../bin/base ./base && go build -ldflags='-s -w' -o ../bin/base.s ./base \
  && go build -o ../bin/withpgmem ./withpgmem && go build -ldflags='-s -w' -o ../bin/withpgmem.s ./withpgmem)
sz() { stat -f %z "$1" 2>/dev/null || stat -c %s "$1"; }
rm -f bin/pgmem.zip && (cd bin && zip -q -9 pgmem.zip pgmem)

cat <<JSON
{
  "arch": "$ARCH",
  "image": "$IMAGE",
  "image_download": $(layers "$IMAGE"),
  "image_disk": $(image_disk "$IMAGE"),
  "ryuk": "$RYUK",
  "ryuk_download": $(layers "$RYUK"),
  "devbox_postgres": "$(basename "$STORE")",
  "devbox_download": ${NIX% *},
  "devbox_disk": ${NIX#* },
  "go_link_added": $(( $(sz bin/withpgmem) - $(sz bin/base) )),
  "go_link_added_stripped": $(( $(sz bin/withpgmem.s) - $(sz bin/base.s) )),
  "binary": $(sz bin/pgmem),
  "binary_compressed": $(sz bin/pgmem.zip)
}
JSON
