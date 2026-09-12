#!/bin/bash
# Build PostgreSQL (PGlite fork) to a self-contained wasm module for pgmem.
#   SJLJ=wasm|emscripten   setjmp/longjmp strategy (default: wasm = wasm EH, legacy try/catch)
#   JOBS=N                 parallelism
set -euo pipefail
HERE=$(cd "$(dirname "$0")" && pwd)
ROOT=$(dirname "$HERE")
source "$ROOT/toolchain/emsdk/emsdk_env.sh" >/dev/null 2>&1
OUT=$HERE/out
PREFIX=$OUT/install
SJLJ=${SJLJ:-wasm}
JOBS=${JOBS:-8}
mkdir -p "$OUT/src"

# PostgreSQL source: a pinned commit of the PGlite fork, downloaded once
# into wasm/out/src (gitignored) and patched in place by patches.py.
source "$HERE/postgres-pglite.lock"
SRC=$OUT/src/postgres-pglite-$commit
if [ ! -f "$SRC/configure" ]; then
  ARCHIVE=$OUT/src/postgres-pglite-$commit.tar.gz
  if [ ! -f "$ARCHIVE" ]; then
    echo "== fetching $repo@$commit"
    curl -fsSL -o "$ARCHIVE" "https://github.com/$repo/archive/$commit.tar.gz"
  fi
  echo "$sha256  $ARCHIVE" | shasum -a 256 -c - >/dev/null || { echo "error: checksum mismatch for $ARCHIVE"; exit 2; }
  tar xzf "$ARCHIVE" -C "$OUT/src"
fi

# pgmem-specific source patches (idempotent); before anything is compiled,
# pglitec.c included.
python3 "$HERE/patches.py" "$SRC"

BASE_CFLAGS="-m32 -O2 -sWASM_BIGINT -sSUPPORT_LONGJMP=$SJLJ \
 -Wno-declaration-after-statement -Wno-macro-redefined -Wno-unused-function \
 -Wno-missing-prototypes -Wno-incompatible-pointer-types"

# Objects that must see the real libc names (no -D overrides).
emcc $BASE_CFLAGS -D__PGMEM__ -c -o "$OUT/pglitec.o" "$SRC/pglite/src/pglitec/pglitec.c"
emcc $BASE_CFLAGS -c -o "$OUT/pgmem_shim.o" "$HERE/pgmem_shim.c"

PG_CFLAGS="$BASE_CFLAGS \
-D__PGLITE__ -D__PGMEM__ -I$HERE \
-Dsystem=pgl_system -Dpopen=pgl_popen -Dpclose=pgl_pclose \
-Dgeteuid=pgl_geteuid -Dgetuid=pgl_getuid -Dgetpwuid=pgl_getpwuid \
-Dexit=pgl_exit \
-Dmunmap=pgl_munmap \
-Dfcntl=pgl_fcntl \
-Datexit=pgl_atexit \
-Dsetsockopt=pgl_setsockopt -Dgetsockopt=pgl_getsockopt -Dgetsockname=pgl_getsockname \
-Drecv=pgl_recv -Dsend=pgl_send -Dconnect=pgl_connect \
-Dpoll=pgmem_poll \
-Dshmget=pgl_shmget -Dshmat=pgl_shmat -Dshmdt=pgl_shmdt -Dshmctl=pgl_shmctl \
-Dlongjmp=pgl_longjmp -Dsiglongjmp=pgl_siglongjmp \
-Ddlopen=pgmem_dlopen -Ddlsym=pgmem_dlsym -Ddlclose=pgmem_dlclose -Ddlerror=pgmem_dlerror"

# Symbols every executable (initdb, postgres) exports to the host.
EXPORTS_COMMON=_main,_pgmem_init,_pgmem_main,_pgmem_call_sighandler,_pgl_freopen,_pgl_run_atexit_funcs,_pgl_getPGliteExitStatus,_pgl_setPGliteExitStatus,_pgl_setPGliteActive,_malloc,_free,_fflush,___errno_location,_strerror,_emscripten_stack_get_current,__emscripten_stack_restore,_emscripten_builtin_memalign,_emscripten_builtin_free,__emscripten_timeout

LDFLAGS="-sWASM_BIGINT -sUSE_PTHREADS=0 -sSUPPORT_LONGJMP=$SJLJ"
LDFLAGS_EX="-sINITIAL_MEMORY=64MB -sALLOW_MEMORY_GROWTH=1 -sSTACK_SIZE=8MB \
 -sEXIT_RUNTIME=1 -sINVOKE_RUN=0 -sENVIRONMENT=node \
 -sERROR_ON_UNDEFINED_SYMBOLS=0 \
 -sEXPORTED_FUNCTIONS=$EXPORTS_COMMON \
 $OUT/pglitec.o $OUT/pgmem_shim.o"

CONFIGURE_PARAMS="\
ac_cv_exeext=.js \
--host wasm32-unknown-linux-gnu \
--without-llvm \
--without-pam \
--disable-largefile \
--with-openssl=no \
--without-readline \
--without-icu \
--without-zlib \
--with-template=emscripten \
--prefix=$PREFIX"

cd "$SRC"
CONF_SIG="$CONFIGURE_PARAMS|$PG_CFLAGS|$LDFLAGS|$LDFLAGS_EX"
if [ ! -f config.status ] || [ "$(cat "$OUT/configure.sig" 2>/dev/null)" != "$CONF_SIG" ]; then
  echo "== configure"
  LDFLAGS="$LDFLAGS" LDFLAGS_EX="$LDFLAGS_EX" CFLAGS="$PG_CFLAGS" \
    emconfigure ./configure $CONFIGURE_PARAMS
  printf '%s' "$CONF_SIG" > "$OUT/configure.sig"
  # make does not track CFLAGS changes; start from clean objects.
  emmake make PORTNAME=emscripten clean >/dev/null || true
fi
# Executables embed pglitec.o/pgmem_shim.o via LDFLAGS_EX without make
# dependencies, so always relink them.
rm -f src/bin/initdb/initdb.js src/bin/initdb/initdb.wasm src/backend/postgres.js src/backend/postgres.wasm

echo "== make"
emmake make PORTNAME=emscripten -j"$JOBS"
echo "== make install"
emmake make PORTNAME=emscripten install

echo "== rebuild loadable modules for static linking"
# Each module gets unique Pg_magic_func/_PG_init symbols; pgmem_dl.c maps
# them back when dfmgr.c asks for the generic names.
LLVM_NM="$ROOT/toolchain/emsdk/upstream/bin/llvm-nm"
MODULE_DIRS="plpgsql=src/pl/plpgsql/src dict_snowball=src/backend/snowball"
for d in src/backend/utils/mb/conversion_procs/*/; do
  n=$(basename "$d")
  MODULE_DIRS="$MODULE_DIRS $n=src/backend/utils/mb/conversion_procs/$n"
done
# Contrib extensions linked in the same way. Their control and SQL files go
# into the share tree below so CREATE EXTENSION finds them. pgcrypto's
# OpenSSL-backed files are replaced by host-backed ones (patches.py).
CONTRIB_MODULES="pgcrypto citext pg_trgm hstore ltree btree_gist btree_gin unaccent tablefunc intarray fuzzystrmatch cube earthdistance seg bloom isn dict_int"
for n in $CONTRIB_MODULES; do
  # the library name (what $libdir/<name> in the extension's SQL refers to)
  # is the Makefile's MODULE_big or MODULES, not always the directory name
  # (intarray builds _int)
  lib=$(sed -n 's/^MODULE_big *= *//p;s/^MODULES *= *//p' "$SRC/contrib/$n/Makefile" | head -1)
  MODULE_DIRS="$MODULE_DIRS ${lib:-$n}=contrib/$n"
done
GEN_ARGS=""
MODULE_OBJS=""
for spec in $MODULE_DIRS; do
  name=${spec%%=*}; dir=${spec#*=}
  emmake make -C "$dir" clean >/dev/null
  emmake make -C "$dir" -j"$JOBS" CFLAGS="$PG_CFLAGS -DPg_magic_func=Pg_magic_func_$name -D_PG_init=_PG_init_$name" all >/dev/null
  objs=$(ls "$SRC/$dir"/*.o | tr '\n' ',')
  GEN_ARGS="$GEN_ARGS $name=$objs"
  MODULE_OBJS="$MODULE_OBJS $(ls "$SRC/$dir"/*.o | tr '\n' ' ')"
done
python3 "$HERE/gen_modules.py" "$LLVM_NM" "$OUT/pgmem_modules_gen.c" $GEN_ARGS
mkdir -p "$PREFIX/share/postgresql/extension"
mkdir -p "$PREFIX/share/postgresql/tsearch_data"
for n in $CONTRIB_MODULES; do
  cp "$SRC/contrib/$n"/*.control "$SRC/contrib/$n"/*.sql "$PREFIX/share/postgresql/extension/"
  # text search dictionaries and rules (DATA_TSEARCH in the module's Makefile)
  for f in $(sed -n 's/^DATA_TSEARCH *= *//p' "$SRC/contrib/$n/Makefile"); do
    cp "$SRC/contrib/$n/$f" "$PREFIX/share/postgresql/tsearch_data/"
  done
done
emcc $BASE_CFLAGS -c -o "$OUT/pgmem_dl.o" "$HERE/pgmem_dl.c"
emcc $BASE_CFLAGS -c -o "$OUT/pgmem_modules_gen.o" "$OUT/pgmem_modules_gen.c"

echo "== link backend (pglite target)"
BACKEND_EXPORTS=$EXPORTS_COMMON,_PostgresMainLoopOnce,_PostgresMainLongJmp,_PostgresSendReadyForQueryIfNecessary,_ProcessStartupPacket,_IsTransactionBlock,_pgl_startPGlite,_pgl_getMyProcPort,_pgl_sendConnData,_pgl_pq_flush,_pq_buffer_remaining_data,_pgmem_module_name
rm -f src/backend/pglite.wasm src/backend/pglite.js
# A small initial memory: the heap grows on demand (emscripten_resize_heap)
# and every initial byte is resident memory the host has to zero.
# --profiling-funcs keeps the function names (a "name" custom section,
# ~0.5 MB, not shipped): gen-aot.sh names the generated Go functions
# after them so a rebuild only changes the functions that changed.
POSTGRES_PGLITE_FLAGS="-sINITIAL_MEMORY=32MB --profiling-funcs -sEXPORTED_FUNCTIONS=$BACKEND_EXPORTS $OUT/pgmem_dl.o $OUT/pgmem_modules_gen.o $MODULE_OBJS" \
  emmake make PORTNAME=emscripten -C src/backend -j"$JOBS" pglite

cp src/backend/pglite.wasm "$OUT/postgres.wasm"
cp src/bin/initdb/initdb.wasm "$OUT/initdb.wasm"
ls -la "$OUT"/*.wasm

echo "== translate to the exnref exception encoding (for wazero)"
for n in postgres initdb; do
  wasm-opt --enable-exception-handling --enable-reference-types --enable-bulk-memory \
    --enable-sign-ext --enable-mutable-globals --enable-nontrapping-float-to-int --enable-multivalue \
    --translate-to-exnref "$OUT/$n.wasm" -o "$OUT/$n.exnref.wasm"
done
# Only the share tree ships with the Go module; the wasm files stay in
# wasm/out as inputs for gen-aot.sh (postgres) and pgmem-mkdata (initdb).
ASSETS="$ROOT/internal/assets"
mkdir -p "$ASSETS"
python3 "$HERE/share-tarball.py" "$PREFIX/share/postgresql" "$ASSETS/share.tar.gz"
ls -la "$ASSETS"
echo "== done"
