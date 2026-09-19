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
'''#include <emscripten/emscripten.h>
#else
''',
'''#include <emscripten/emscripten.h>
#ifdef __PGMEM__
/* Keep the pgmem ABI explicit; the full PGlite callback surface is
 * internal to the statically linked backend. */
#undef EMSCRIPTEN_KEEPALIVE
#define EMSCRIPTEN_KEEPALIVE
#endif
#else
''',
'Keep the pgmem ABI explicit')

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

# Session authorization: single-user mode never gives the
# session_authorization GUC its reset value (upstream says so in the
# comment), so RESET SESSION AUTHORIZATION, SET ... DEFAULT and the DISCARD
# ALL pgmem runs between connections left a SET SESSION AUTHORIZATION in
# place for the next client of the shared backend. Set it like
# InitializeSessionUserId() does.
patch('src/backend/utils/init/miscinit.c',
"""	SetSessionAuthorization(BOOTSTRAP_SUPERUSERID, true);

	/* We could do SetConfigOption("role"), but let's be consistent */
	SetCurrentRoleId(InvalidOid, false);
}
""",
"""	SetSessionAuthorization(BOOTSTRAP_SUPERUSERID, true);

#ifdef __PGMEM__
	/*
	 * pgmem: every client shares this session, and RESET SESSION
	 * AUTHORIZATION (also DISCARD ALL between clients) must bring it back
	 * to the superuser, so give the GUC its reset value after all. Not in
	 * bootstrap mode (initdb): no catalog to read the name from yet.
	 */
	if (!IsBootstrapProcessingMode())
	{
		char	   *rname = GetUserNameFromId(BOOTSTRAP_SUPERUSERID, true);

		if (rname != NULL)
			SetConfigOption("session_authorization", rname,
							PGC_BACKEND, PGC_S_OVERRIDE);
	}
#endif

	/* We could do SetConfigOption("role"), but let's be consistent */
	SetCurrentRoleId(InvalidOid, false);
}
""",
'give the GUC its reset value after all')

# Session cleanup between clients: pgmem used to run "ROLLBACK; DISCARD ALL"
# as SQL when a new connection found the backend free, which showed up in
# pg_stat_statements (like a pooler's server_reset_query does) and in the
# log. pgmem_reset_session does the same work as C calls, so nothing is
# executed as a statement.
patch('src/backend/tcop/postgres.c',
"""void PostgresMainLoopOnce() {
""",
"""#ifdef __PGMEM__
#include "commands/discard.h"

/*
 * pgmem_reset_session: what a client expects from a fresh session, done
 * without running a statement (so pg_stat_statements does not count it):
 * abort anything left open, then, with discard, everything DISCARD ALL
 * does plus forgetting the temp namespace.
 */
extern void pgmem_forget_temp_namespace(void);

void
pgmem_reset_session(int discard)
{
	DiscardStmt stmt;

	AbortOutOfAnyTransaction();
	if (!discard)
		return;
	StartTransactionCommand();
	stmt.type = T_DiscardStmt;
	stmt.target = DISCARD_ALL;
	DiscardCommand(&stmt, true);
	CommitTransactionCommand();
	/* DISCARD TEMP emptied it; a new client starts without one, like a new backend */
	pgmem_forget_temp_namespace();
}
#endif

void PostgresMainLoopOnce() {
""",
'pgmem_reset_session')

# Temp namespace: DISCARD TEMP empties the session's pg_temp_N but the
# session keeps it assigned, so a later client's first CREATE TEMP TABLE
# does not change the effective search_path the way it does in a new
# backend (which is what revalidates cached plans). Forgetting it makes
# the next client start like a new backend; pg_temp_N is found and
# cleaned again on first use.
patch('src/backend/catalog/namespace.c',
"""/*
 * GetTempToastNamespace - get the OID of my temporary-toast-table namespace,
""",
"""#ifdef __PGMEM__
void
pgmem_forget_temp_namespace(void)
{
	myTempNamespace = InvalidOid;
	myTempToastNamespace = InvalidOid;
	myTempNamespaceSubID = InvalidSubTransactionId;
	/* the cached search paths still carry the old pg_temp_N */
	baseSearchPathValid = false;
	searchPathCacheValid = false;
}
#endif

/*
 * GetTempToastNamespace - get the OID of my temporary-toast-table namespace,
""",
'pgmem_forget_temp_namespace')

# ---- multi-process model (EXEC_BACKEND on the pgmem host) ----

# Process creation: EXEC_BACKEND's internal_forkexec has written the
# child's variables to a file; instead of fork()+execv() the host starts a
# new module instance running "postgres --forkchild=<kind> <file>" and
# returns its pid.
patch('src/backend/postmaster/launch_backend.c',
'''	/* Fire off execv in child */
	if ((pid = fork_process()) == 0)
	{
		if (execv(postgres_exec_path, argv) < 0)
		{
			ereport(LOG,
					(errmsg("could not execute server process \\"%s\\": %m",
							postgres_exec_path)));
			/* We're already in the child process here, can't return */
			exit(1);
		}
	}

	return pid;					/* Parent returns pid, or -1 on fork failure */
''',
'''#ifdef __PGMEM__
	{
		extern int pgmem_spawn(const char *forkarg, const char *paramfile)
			__attribute__((import_module("env"), import_name("pgmem_spawn")));

		pid = pgmem_spawn(argv[1], argv[2]);
		if (pid < 0)
		{
			errno = -pid;
			pid = -1;
		}
	}
#else
	/* Fire off execv in child */
	if ((pid = fork_process()) == 0)
	{
		if (execv(postgres_exec_path, argv) < 0)
		{
			ereport(LOG,
					(errmsg("could not execute server process \\"%s\\": %m",
							postgres_exec_path)));
			/* We're already in the child process here, can't return */
			exit(1);
		}
	}
#endif

	return pid;					/* Parent returns pid, or -1 on fork failure */
''',
'pgmem_spawn')

# pg_usleep: Emscripten's nanosleep spins on the clock; sleep on the host
# instead, waking early when a signal arrives like nanosleep does.
patch('src/port/pgsleep.c',
'''		(void) nanosleep(&delay, NULL);
''',
'''#ifdef __PGMEM__
		{
			extern void pgmem_usleep(int us)
				__attribute__((import_module("env"), import_name("pgmem_usleep")));

			(void) delay;
			pgmem_usleep((int) microsec);
		}
#else
		(void) nanosleep(&delay, NULL);
#endif
''',
'pgmem_usleep')

# ReadyForQuery at session start: PGlite split PostgresMain's loop body
# into PostgresMainLoopOnce, which reads the global send_ready_for_query
# (pglitec.c, initially false), while PostgresMain still declares a local
# of the same name and sets that. In single mode the host sends the first
# ReadyForQuery; a real backend must set the global or the client never
# gets one.
patch('src/backend/tcop/postgres.c',
'''	/* these must be volatile to ensure state is preserved across longjmp: */
	volatile bool send_ready_for_query = true;
	volatile bool idle_in_transaction_timeout_enabled = false;
''',
'''	/* these must be volatile to ensure state is preserved across longjmp: */
#ifndef __PGMEM__
	volatile bool send_ready_for_query = true;
#endif
	volatile bool idle_in_transaction_timeout_enabled = false;
''',
'#ifndef __PGMEM__\n\tvolatile bool send_ready_for_query = true;')

patch('src/backend/tcop/postgres.c',
'''	if (!ignore_till_sync)
		send_ready_for_query = true;	/* initially, or after error */

	/*
	 * Non-error queries loop here.
''',
'''#ifdef __PGMEM__
	/* the global PostgresMainLoopOnce reads; in single mode the host sends it */
	if (!ignore_till_sync && !is_pglite_active)
		send_ready_for_query = true;
#else
	if (!ignore_till_sync)
		send_ready_for_query = true;	/* initially, or after error */
#endif

	/*
	 * Non-error queries loop here.
''',
'in single mode the host sends it')

# 32-bit EXEC_BACKEND: internal_forkexec writes SizeOfBackendParameters(len)
# bytes (offsetof the flexible startup_data member plus the data), but the
# reader asks for sizeof(BackendParameters), which on wasm32 is 4 bytes more
# than the offset (the struct is padded to its 8-byte alignment after a
# 4-byte size_t). A child without startup data then reads short and dies.
patch('src/backend/postmaster/launch_backend.c',
'''	if (fread(&param, sizeof(param), 1, fp) != 1)
	{
		write_stderr("could not read from backend variables file \\"%s\\": %m\\n", id);
		exit(1);
	}
''',
'''#ifdef __PGMEM__
	/* what the writer wrote: up to the flexible member, not the padded sizeof */
	if (fread(&param, offsetof(BackendParameters, startup_data), 1, fp) != 1)
#else
	if (fread(&param, sizeof(param), 1, fp) != 1)
#endif
	{
		write_stderr("could not read from backend variables file \\"%s\\": %m\\n", id);
		exit(1);
	}
''',
'not the padded sizeof')

# Checkpoints in a cluster: PGlite made RequestCheckpoint run the checkpoint
# in the calling backend (single-user mode has no checkpointer). A backend
# of a cluster must ask the checkpointer instead (a backend cannot process
# sync requests). Keep the local checkpoint for the shared single session.
patch('src/backend/postmaster/checkpointer.c',
'''#ifndef __PGLITE__	
	if (!IsPostmasterEnvironment)
#endif
	{''',
'''#if defined(__PGMEM__)
	{
		extern int is_pglite_active;

		if (!IsPostmasterEnvironment || is_pglite_active)
		{
			CreateCheckPoint(flags | CHECKPOINT_IMMEDIATE);
			smgrdestroyall();
			return;
		}
	}
	if (false)
#elif !defined(__PGLITE__)
	if (!IsPostmasterEnvironment)
#endif
	{''',
'!IsPostmasterEnvironment || is_pglite_active')

# WAL-driven checkpoints: PGlite disabled the checkpoint request on a WAL
# segment switch; a cluster has a checkpointer to take it.
patch('src/backend/access/transam/xlog.c',
'''#ifndef __PGLITE__					
				if (IsUnderPostmaster && XLogCheckpointNeeded(openLogSegNo))''',
'''#if !defined(__PGLITE__) || defined(__PGMEM__)
				if (IsUnderPostmaster && XLogCheckpointNeeded(openLogSegNo))''',
'#if !defined(__PGLITE__) || defined(__PGMEM__)\n\t\t\t\tif (IsUnderPostmaster && XLogCheckpointNeeded')
