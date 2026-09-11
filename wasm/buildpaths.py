#!/usr/bin/env python3
"""Keep the build directory out of the compiled module.

usage: buildpaths.py <postgres-pglite dir> <repo root> <emsdk root>

pgmem ships the wasm as generated Go source, so anything the compiler bakes
into static data is committed. PostgreSQL bakes in two sets of build paths:

  * src/common/Makefile passes CC / CFLAGS / LDFLAGS_EX / configure_args to
    config_info.c as -DVAL_* string literals — that is what pg_config
    reports — and they carry -I<root>/wasm, the shim objects and the emsdk
    location;
  * src/port/pg_config_paths.h holds the install prefix (PGBINDIR,
    PGSHAREDIR, ...), generated from Makefile.global.

Both are informational here: pgmem's guest finds its share tree through the
/pglite layout, not through PGSHAREDIR. So the roots are replaced with the
fixed tokens /pgmem-build and /emsdk (a path that does not exist in the
guest filesystem, as the build directory did not), and two checkouts at
different paths build
byte-identical wasm. __FILE__ is handled separately, by -ffile-prefix-map in
build.sh's BASE_CFLAGS.

The VAL_* values are make variables the build itself needs (CFLAGS really has
to keep the include path), so the reference is wrapped in $(subst ...) rather
than the variable rewritten. Idempotent; run after configure on every build.
"""
import glob
import os
import re
import sys

src, root, emsdk = sys.argv[1], sys.argv[2], sys.argv[3]


def real(p):
    try:
        return os.path.realpath(p)
    except OSError:
        return p


# Longest first: the emsdk can live under the repo root.
subs = []
for path, token in ((emsdk, '/emsdk'), (root, '/pgmem-build')):
    for p in dict.fromkeys((real(path), path)):
        subs.append((p, token))
subs.sort(key=lambda s: -len(s[0]))


def scrub(text):
    for old, new in subs:
        text = text.replace(old, new)
    return text


# configure records its own argument list (with the prefix, CC and CFLAGS) in
# pg_config.h as CONFIGURE_ARGS, which config_info.c reports as well.
config_h = os.path.join(src, 'src/include/pg_config.h')
if os.path.exists(config_h):
    s = open(config_h).read()
    t = scrub(s)
    if t != s:
        open(config_h, 'w').write(t)
        print('pg_config.h: build paths scrubbed')
    os.utime(config_h, None)

header = os.path.join(src, 'src/port/pg_config_paths.h')
if os.path.exists(header):
    s = open(header).read()
    t = scrub(s)
    if t != s:
        open(header, 'w').write(t)
        print('pg_config_paths.h: install paths scrubbed')
    # Keep it newer than Makefile.global so make does not regenerate it.
    os.utime(header, None)

# config_info.o does not depend on the Makefile that carries its -DVAL_*
# values, so make would keep an object compiled before the wrap below. Both
# variants go: the backend links the _srv one.
# The archives are what the links actually consume, and they already hold a
# config_info member compiled before the wrap, so drop them as well; make
# recompiles the one missing object and rebuilds them.
for stale in (glob.glob(os.path.join(src, 'src/common/config_info*.o'))
              + glob.glob(os.path.join(src, 'src/common/libpgcommon*.a'))):
    os.remove(stale)

makefile = os.path.join(src, 'src/common/Makefile')
s = open(makefile).read()
if 'pgmem: VAL_* paths scrubbed' not in s:
    def wrap(m):
        expr = m.group(2)
        for old, new in subs:
            expr = '$(subst %s,%s,%s)' % (old, new, expr)
        return m.group(1) + expr + m.group(3)

    # override CPPFLAGS += -DVAL_CFLAGS="\"$(CFLAGS)\""
    t, n = re.subn(r'(-DVAL_[A-Z_]+="\\")(\$\([A-Z_]+\))(\\"")', wrap, s)
    if n:
        open(makefile, 'w').write('# pgmem: VAL_* paths scrubbed\n' + t)
        print(f'src/common/Makefile: {n} VAL_* values scrubbed')
