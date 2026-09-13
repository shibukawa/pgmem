#!/usr/bin/env bash
# Stamps one release version into every package manifest: the npm packages
# (including @pgmem/core's pins on its platform packages), the Python
# project and the Gradle build. The binary gets the same number through
# VERSION in scripts/build-binaries.sh.
#
#   scripts/set-version.sh 0.1.0
set -euo pipefail
cd "$(dirname "$0")/.."
python3 - "${1:?usage: scripts/set-version.sh X.Y.Z}" <<'PY'
import json
import pathlib
import re
import sys

v = sys.argv[1]
if not re.fullmatch(r"\d+\.\d+\.\d+", v):
    sys.exit(f"set-version: {v!r} is not X.Y.Z")

node = pathlib.Path("packages/node")
for p in [node / "core" / "package.json", *sorted(node.glob("platforms/*/package.json"))]:
    d = json.loads(p.read_text())
    d["version"] = v
    for dep in d.get("optionalDependencies", {}):
        d["optionalDependencies"][dep] = v
    p.write_text(json.dumps(d, indent=2, ensure_ascii=False) + "\n")


def stamp(path, pattern, line):
    p = pathlib.Path(path)
    text, n = re.subn(pattern, line, p.read_text(), count=1, flags=re.M)
    if n != 1:
        sys.exit(f"set-version: no version line in {path}")
    p.write_text(text)


stamp("packages/python/pyproject.toml", r'^version = ".*"$', f'version = "{v}"')
stamp("packages/java/gradle.properties", r"^version=.*$", f"version={v}")
print("version set to", v)
PY
