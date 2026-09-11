#!/usr/bin/env python3
"""Idempotent pgmem source patches for the PostgreSQL (PGlite fork) checkout.

usage: patches.py <postgres-pglite dir>

Everything is guarded by __PGMEM__ so the tree still builds unpatched.
"""
import sys, os

src = sys.argv[1]

def patch(rel, old, new, marker):
    p = os.path.join(src, rel)
    s = open(p).read()
    if marker in s:
        print(f'{rel}: already patched')
        return
    if old not in s:
        sys.exit(f'{rel}: anchor not found')
    open(p, 'w').write(s.replace(old, new, 1))
    print(f'{rel}: patched')

# CRC-32C: the wasm build has no hardware CRC; hand the raw state update to
# the host (Go's hash/crc32), keeping PostgreSQL's INIT/FIN conventions.
patch('src/include/port/pg_crc32c.h',
'''#define COMP_CRC32C(crc, data, len) \\
	((crc) = pg_comp_crc32c_sb8((crc), (data), (len)))
#ifdef WORDS_BIGENDIAN''',
'''#ifdef __PGMEM__
extern pg_crc32c pgmem_crc32c(pg_crc32c crc, const void *data, size_t len)
	__attribute__((import_module("env"), import_name("pgmem_crc32c")));
#define COMP_CRC32C(crc, data, len) \\
	((crc) = pgmem_crc32c((crc), (data), (len)))
#else
#define COMP_CRC32C(crc, data, len) \\
	((crc) = pg_comp_crc32c_sb8((crc), (data), (len)))
#endif
#ifdef WORDS_BIGENDIAN''',
'pgmem_crc32c')

# Crypto hashes (md5(), sha256() ..., SCRAM's HMAC): the whole cryptohash
# implementation moves to the host (Go's crypto/*); the original body is
# kept for builds without __PGMEM__. The .inc lives next to build.sh, which
# passes -I for it.
p = os.path.join(src, 'src/common/cryptohash.c')
s = open(p).read()
if 'pgmem_cryptohash.inc' not in s:
    open(p, 'w').write('#ifdef __PGMEM__\n#include "pgmem_cryptohash.inc"\n#else\n' + s + '\n#endif /* __PGMEM__ */\n')
    print('src/common/cryptohash.c: patched')
else:
    print('src/common/cryptohash.c: already patched')

# UUID: draw the random bytes straight from the host's crypto/rand instead
# of the emulated /dev/urandom.
patch('src/backend/utils/adt/uuid.c',
'''#include "utils/uuid.h"
''',
'''#include "utils/uuid.h"

#ifdef __PGMEM__
extern int pgmem_random_bytes(void *buf, size_t len)
	__attribute__((import_module("env"), import_name("pgmem_random_bytes")));
#define pg_strong_random(buf, len) (pgmem_random_bytes((buf), (len)) != 0)
#endif
''',
'pgmem_random_bytes')
