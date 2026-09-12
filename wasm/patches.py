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

# LISTEN/NOTIFY: report changes to the session's listen set to the host at
# commit time (op 1 = LISTEN, 0 = UNLISTEN, 2 = UNLISTEN *). The single
# backend session is shared by every client connection; pgmem keeps a
# per-connection registry from these calls and routes NotifyResponse
# messages to the connections that actually listen.
patch('src/backend/commands/async.c',
'''static void
Exec_ListenCommit(const char *channel)
{
	MemoryContext oldcontext;

	/* Do nothing if we are already listening on this channel */
	if (IsListeningOn(channel))
		return;
''',
'''#ifdef __PGMEM__
extern void pgmem_listen(const char *channel, int op)
	__attribute__((import_module("env"), import_name("pgmem_listen")));
#else
#define pgmem_listen(channel, op) ((void) 0)
#endif

static void
Exec_ListenCommit(const char *channel)
{
	MemoryContext oldcontext;

	pgmem_listen(channel, 1);
	/* Do nothing if we are already listening on this channel */
	if (IsListeningOn(channel))
		return;
''',
'pgmem_listen(channel, 1)')

patch('src/backend/commands/async.c',
'''	if (Trace_notify)
		elog(DEBUG1, "Exec_UnlistenCommit(%s,%d)", channel, MyProcPid);
''',
'''	if (Trace_notify)
		elog(DEBUG1, "Exec_UnlistenCommit(%s,%d)", channel, MyProcPid);
	pgmem_listen(channel, 0);
''',
'pgmem_listen(channel, 0)')

patch('src/backend/commands/async.c',
'''	if (Trace_notify)
		elog(DEBUG1, "Exec_UnlistenAllCommit(%d)", MyProcPid);
''',
'''	if (Trace_notify)
		elog(DEBUG1, "Exec_UnlistenAllCommit(%d)", MyProcPid);
	pgmem_listen("", 2);
''',
'pgmem_listen("", 2)')

# pgcrypto: the OpenSSL-backed files (digests, ciphers, OpenPGP bignum
# arithmetic) and the zlib-backed OpenPGP compression are replaced by
# host-backed bodies; see the .inc files next to build.sh. The originals stay
# for builds without __PGMEM__.
for rel, inc in [('contrib/pgcrypto/openssl.c', 'pgmem_pgcrypto_openssl.inc'),
                 ('contrib/pgcrypto/pgp-mpi-openssl.c', 'pgmem_pgcrypto_mpi.inc'),
                 ('contrib/pgcrypto/pgp-compress.c', 'pgmem_pgcrypto_compress.inc')]:
    p = os.path.join(src, rel)
    s = open(p).read()
    if inc not in s:
        open(p, 'w').write(f'#ifdef __PGMEM__\n#include "{inc}"\n#else\n' + s + '\n#endif /* __PGMEM__ */\n')
        print(f'{rel}: patched')
    else:
        print(f'{rel}: already patched')

# ReadyForQuery after an error: PGlite's pgl_longjmp decides whether to send
# ReadyForQuery before PostgresMainLongJmp has set ignore_till_sync, so an
# error inside an extended-protocol batch is followed by a ReadyForQuery
# that real PostgreSQL withholds until Sync (the client then sees two).
# Move the decision after the recovery block, where PostgresMain has it.
patch('pglite/src/pglitec/pglitec.c',
"""        // reset this as it is expected
        if (!ignore_till_sync)
		    send_ready_for_query = true;	/* initially, or after error */
""",
"""#ifndef __PGMEM__
        // reset this as it is expected
        if (!ignore_till_sync)
		    send_ready_for_query = true;	/* initially, or after error */
#endif
""",
'#ifndef __PGMEM__\n        // reset this as it is expected')

patch('src/backend/tcop/postgres.c',
"""		/* Now we can allow interrupts again */
		RESUME_INTERRUPTS();
}
""",
"""		/* Now we can allow interrupts again */
		RESUME_INTERRUPTS();

#ifdef __PGMEM__
		/* Same order as PostgresMain: decided after ignore_till_sync is set. */
		if (!ignore_till_sync)
			send_ready_for_query = true;	/* after error */
#endif
}
""",
'decided after ignore_till_sync is set')
