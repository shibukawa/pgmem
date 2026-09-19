#!/bin/bash
# add-extension.sh: bundle one PostgreSQL contrib extension into pgmem.
#
#   add-extension.sh status              bundled vs. available contrib modules
#   add-extension.sh add <contrib-name>  edit wasm/build.sh, vendor the upstream
#                                        regression files, register them in
#                                        regress_test.go, print what is left
#   add-extension.sh build               wasm then aot
#   add-extension.sh wasm                ./wasm/build.sh (through devbox when emcc
#                                        is not on PATH): postgres.wasm, share.tar.gz
#   add-extension.sh aot                 ./wasm/gen-aot.sh: internal/aot/pgaot
#   add-extension.sh smoke <name>        start pgmem, CREATE EXTENSION <name>,
#                                        list what it installed
#   add-extension.sh test <name>         replay the vendored regression files
#
# Runs from anywhere inside the repository. The PostgreSQL source (the PGlite
# fork pinned by wasm/postgres-pglite.lock) is fetched into wasm/out/src when
# it is not there yet, exactly as wasm/build.sh does.
set -euo pipefail

# repository root: the closest ancestor of this script that has wasm/build.sh
ROOT=$(cd "$(dirname "$0")" && pwd)
while [ ! -f "$ROOT/wasm/build.sh" ]; do
  [ "$ROOT" = / ] && { echo "error: wasm/build.sh not found above $0" >&2; exit 1; }
  ROOT=$(dirname "$ROOT")
done
cd "$ROOT"
OUT=wasm/out

# fetch_source: make sure wasm/out/src/postgres-pglite-<commit> exists (the
# archive part of wasm/build.sh, without the toolchain requirement) and set SRC.
fetch_source() {
  # shellcheck disable=SC1091
  source wasm/postgres-pglite.lock
  SRC=$OUT/src/postgres-pglite-$commit
  if [ ! -f "$SRC/configure" ]; then
    mkdir -p "$OUT/src"
    ARCHIVE=$OUT/src/postgres-pglite-$commit.tar.gz
    if [ ! -f "$ARCHIVE" ]; then
      echo "== fetching $repo@$commit"
      curl -fsSL -o "$ARCHIVE" "https://github.com/$repo/archive/$commit.tar.gz"
    fi
    echo "$sha256  $ARCHIVE" | shasum -a 256 -c - >/dev/null || { echo "error: checksum mismatch for $ARCHIVE" >&2; exit 2; }
    tar xzf "$ARCHIVE" -C "$OUT/src"
  fi
}

bundled() {
  sed -n 's/^CONTRIB_MODULES="\(.*\)"$/\1/p' wasm/build.sh
}

# make_var FILE VAR: the value of a Makefile variable, continuation lines joined
make_var() {
  python3 - "$1" "$2" <<'PY'
import re, sys
text = open(sys.argv[1]).read().replace("\\\n", " ")
m = re.search(r"^%s[ \t]*[:?]?=[ \t]*(.*)$" % re.escape(sys.argv[2]), text, re.M)
print(" ".join(m.group(1).split()) if m else "")
PY
}

cmd_status() {
  fetch_source
  echo "bundled (CONTRIB_MODULES in wasm/build.sh):"
  echo "  $(bundled)"
  echo "not bundled, with upstream regression tests:"
  for d in "$SRC"/contrib/*/; do
    n=$(basename "$d")
    case " $(bundled) " in *" $n "*) continue;; esac
    [ -d "$d/sql" ] || continue
    echo "  $n: $(make_var "$d/Makefile" REGRESS)"
  done
}

cmd_add() {
  name=${1:?usage: add-extension.sh add <contrib-name>}
  fetch_source
  dir=$SRC/contrib/$name
  if [ ! -f "$dir/Makefile" ]; then
    echo "error: no contrib/$name in the PostgreSQL source; run 'add-extension.sh status'" >&2
    exit 1
  fi
  lib=$(make_var "$dir/Makefile" MODULE_big); [ -n "$lib" ] || lib=$(make_var "$dir/Makefile" MODULES)
  if [ -z "$lib" ]; then
    echo "warning: contrib/$name has no MODULE_big/MODULES (SQL-only extension); the object loop in wasm/build.sh expects .o files" >&2
  fi
  for v in SHLIB_LINK PG_CPPFLAGS PG_CFLAGS; do
    val=$(make_var "$dir/Makefile" $v)
    [ -n "$val" ] && echo "note: $v = $val (a host library; see uuid-ossp in wasm/build.sh for how one is replaced)"
  done

  # 1. CONTRIB_MODULES in wasm/build.sh
  case " $(bundled) " in
    *" $name "*) echo "== $name is already in CONTRIB_MODULES";;
    *) python3 - "$name" <<'PY'
import re, sys
name = sys.argv[1]
p = "wasm/build.sh"; s = open(p).read()
s2, n = re.subn(r'^(CONTRIB_MODULES=".*)"$', lambda m: m.group(1) + " " + name + '"', s, count=1, flags=re.M)
assert n == 1, "CONTRIB_MODULES line not found"
open(p, "w").write(s2)
print("== added %s to CONTRIB_MODULES in wasm/build.sh" % name)
PY
  esac

  # 2. vendor the regression files
  regress=$(make_var "$dir/Makefile" REGRESS)
  if [ -z "$regress" ]; then
    echo "note: contrib/$name has no REGRESS tests upstream; nothing vendored (write a smoke test of your own)"
  else
    dest=testdata/regress/$name
    mkdir -p "$dest"
    for sub in sql expected data; do
      [ -d "$dir/$sub" ] && { rm -rf "$dest/$sub"; cp -R "$dir/$sub" "$dest/$sub"; }
    done
    echo "== vendored contrib/$name/{sql,expected$( [ -d "$dir/data" ] && printf ',data')} into $dest"
    # the psql emulation in regress_test.go knows a fixed set of meta-commands
    unsupported=$(grep -hoE '^\\[a-zA-Z_]+' "$dest"/sql/*.sql | sort -u \
      | grep -vE '^\\(if|elif|else|endif|quit|q|set|unset|x|echo|copy|gset)$' || true)
    if [ -n "$unsupported" ]; then
      echo "!! meta-commands the regress runner does not emulate: $(echo $unsupported)"
      echo "   cut those lines (and their expected output) or leave the file out of the list:"
      grep -nE "^\\\\($(echo $unsupported | sed 's/\\//g; s/ /|/g'))" "$dest"/sql/*.sql | sed 's/^/   /'
    fi
    # 3. contribRegress in regress_test.go
    python3 - "$name" $regress <<'PY'
import re, sys
name, files = sys.argv[1], sys.argv[2:]
p = "regress_test.go"; s = open(p).read()
if re.search(r'^\t"%s":' % re.escape(name), s, re.M):
    print("== %s is already in contribRegress (regress_test.go)" % name)
    sys.exit()
start = s.index("var contribRegress = map[string][]string{")
end = s.index("\n}\n", start)
entry = '\t"%s": {%s},\n' % (name, ", ".join('"%s"' % f for f in files))
s = s[:end + 1] + entry + s[end + 1:]
open(p, "w").write(s)
print("== registered %s in contribRegress (regress_test.go): %s" % (name, " ".join(files)))
PY
    gofmt -w regress_test.go
    opts=$(make_var "$dir/Makefile" REGRESS_OPTS)
    [ -n "$opts" ] && echo "!! REGRESS_OPTS = $opts: mirror --temp-config settings as params and --load-extension as setup in regressSuites (regress_test.go)"
  fi

  cat <<TXT

next:
  1. edit the extension lists by hand: README.md ("- Extensions:" line),
     .knowledge/policy/bundled-extensions.md, website/src/content/docs/extensions.md
     and website/src/content/docs/ja/extensions.md
  2. $0 build          (rebuilds postgres.wasm, share.tar.gz and internal/aot/pgaot)
  3. $0 smoke $name
  4. $0 test $name
  5. commit everything as one commit: "pgmem: bundle $name"
TXT
}

cmd_wasm() {
  if command -v emcc >/dev/null 2>&1 || [ -f toolchain/emsdk/emsdk_env.sh ]; then
    run=()
  elif command -v devbox >/dev/null 2>&1; then
    run=(devbox run -q --)
  else
    echo "error: neither emcc nor devbox found; see README.md 'Rebuilding the wasm module'" >&2
    exit 1
  fi
  echo "== wasm/build.sh"; "${run[@]}" ./wasm/build.sh
}

cmd_aot() {
  # gen-aot.sh needs git (to clone the wasm2go fork), go, python3 and
  # wasm-opt. Devbox has wasm-opt but its environment breaks the host's git
  # ("error: tool 'git' not found"), so run natively and only borrow the
  # wasm-opt path from Devbox when the host has none.
  export AOT_WASM_OPT=${AOT_WASM_OPT:-$(command -v wasm-opt || true)}
  if [ -z "$AOT_WASM_OPT" ] && command -v devbox >/dev/null 2>&1; then
    AOT_WASM_OPT=$(devbox run -q -- sh -c 'command -v wasm-opt')
  fi
  if [ -z "$AOT_WASM_OPT" ]; then
    echo "error: wasm-opt not found (brew install binaryen / apt-get install binaryen, or devbox install)" >&2
    exit 1
  fi
  echo "== wasm/gen-aot.sh (wasm-opt: $AOT_WASM_OPT)"; ./wasm/gen-aot.sh
}

cmd_build() {
  cmd_wasm
  cmd_aot
}

cmd_smoke() {
  name=${1:?usage: add-extension.sh smoke <extension>}
  go run "$(dirname "$0")" "$name"
}

cmd_test() {
  name=${1:?usage: add-extension.sh test <extension>}
  go test . -run "TestContribRegress/$name\$" -count=1 -v 2>&1 | grep -E '^(=== RUN|--- |PASS|FAIL|ok|\s+regress)' || go test . -run "TestContribRegress/$name\$" -count=1
}

case "${1:-}" in
  status) cmd_status;;
  add) shift; cmd_add "$@";;
  build) cmd_build;;
  wasm) cmd_wasm;;
  aot) cmd_aot;;
  smoke) shift; cmd_smoke "$@";;
  test) shift; cmd_test "$@";;
  *) sed -n '2,19p' "$0"; exit 1;;
esac
