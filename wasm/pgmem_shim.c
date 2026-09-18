/*
 * pgmem_shim.c
 *   Host bridge for pgmem. Replaces the JS-side callbacks that PGlite
 *   installs via pgl_set_*() with C functions that forward to imports
 *   implemented by the Go host. Compiled WITHOUT the -D overrides that
 *   PostgreSQL sources get, so plain libc names are used here.
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/types.h>

#define PGMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))
#define PGMEM_EXPORT __attribute__((used, visibility("default")))

/* ---- host imports ---- */
PGMEM_IMPORT("pgmem_recv") extern int pgmem_host_recv(void *buf, int max_len);
PGMEM_IMPORT("pgmem_send") extern int pgmem_host_send(const void *buf, int len);
/* Run "postgres ..." as a child. stdin_path/stdout_path may be "" (inherit). Returns exit code. */
PGMEM_IMPORT("pgmem_run") extern int pgmem_host_run(const char *cmd, const char *stdin_path, const char *stdout_path);
/* Sleep for timeout_ms (-1 = "forever", the host decides how long). */
PGMEM_IMPORT("pgmem_poll") extern int pgmem_host_poll(int timeout_ms);

/* ---- pglitec.c hooks ---- */
typedef ssize_t (*pgl_read_t)(void *buffer, size_t max_length);
typedef ssize_t (*pgl_write_t)(void *buffer, size_t length);
extern void pgl_set_rw_cbs(pgl_read_t read_cb, pgl_write_t write_cb);
typedef ssize_t (*pglite_system_t)(const char *command);
extern void pgl_set_system_fn(pglite_system_t system_fn);
typedef FILE *(*pglite_popen_t)(const char *command, const char *mode);
extern void pgl_set_popen_fn(pglite_popen_t popen_fn);
typedef int (*pglite_pclose_t)(FILE *stream);
extern void pgl_set_pclose_fn(pglite_pclose_t pclose_fn);

#define PGMEM_STDIN_PATH "/pgmem/pgstdin"
#define PGMEM_STDOUT_PATH "/pgmem/pgstdout"

static ssize_t shim_read(void *buf, size_t n) { return pgmem_host_recv(buf, (int) n); }
static ssize_t shim_write(void *buf, size_t n) { return pgmem_host_send(buf, (int) n); }
/* popen/system report a wait(2) status: exit code in bits 8..15. */
static int wait_status(int rc) { return (rc & 0xff) << 8; }
static ssize_t shim_system(const char *cmd) { return wait_status(pgmem_host_run(cmd, "", "")); }

static FILE *popen_stream = NULL;
static char popen_cmd[8192];
static int popen_pending = 0; /* mode "w": run the command on pclose */
static int popen_last_rc = 0;

static FILE *shim_popen(const char *cmd, const char *mode)
{
	strncpy(popen_cmd, cmd, sizeof(popen_cmd) - 1);
	popen_cmd[sizeof(popen_cmd) - 1] = 0;
	if (mode[0] == 'w')
	{
		popen_pending = 1;
		popen_stream = fopen(PGMEM_STDIN_PATH, "w");
		return popen_stream;
	}
	popen_pending = 0;
	popen_last_rc = wait_status(pgmem_host_run(cmd, "", PGMEM_STDOUT_PATH));
	popen_stream = fopen(PGMEM_STDOUT_PATH, "r");
	return popen_stream;
}

static int shim_pclose(FILE *f)
{
	int rc;
	if (f == NULL || f != popen_stream)
		return -1;
	fclose(f);
	popen_stream = NULL;
	if (popen_pending)
	{
		popen_pending = 0;
		rc = wait_status(pgmem_host_run(popen_cmd, PGMEM_STDIN_PATH, ""));
		return rc;
	}
	return popen_last_rc;
}

/*
 * poll(2) replacement. In single-user mode the only pollers are latch
 * waits (pg_sleep, timeouts); no fd can ever become ready, so sleep for
 * the timeout on the host and report a timeout. PostgreSQL's WaitLatch
 * loops re-check their conditions, so a shorter host sleep is also fine.
 */
struct pgmem_pollfd
{
	int			fd;
	short		events;
	short		revents;
};

PGMEM_EXPORT int pgmem_poll(struct pgmem_pollfd *fds, unsigned long nfds, int timeout)
{
	unsigned long i;

	for (i = 0; i < nfds; i++)
		fds[i].revents = 0;
	pgmem_host_poll(timeout);
	return 0;
}

PGMEM_EXPORT void pgmem_init(void)
{
	pgl_set_rw_cbs(shim_read, shim_write);
	pgl_set_system_fn(shim_system);
	pgl_set_popen_fn(shim_popen);
	pgl_set_pclose_fn(shim_pclose);
}

extern int main(int argc, char **argv);

/* args: NUL-separated argv strings, len bytes total (including trailing NUL). */
PGMEM_EXPORT int pgmem_main(char *args, int len)
{
	int argc = 0, i, k = 0;
	char **argv;
	char *p = args;
	for (i = 0; i < len; i++)
		if (args[i] == 0)
			argc++;
	argv = (char **) malloc(sizeof(char *) * (argc + 1));
	for (i = 0; i < argc; i++)
	{
		argv[i] = p;
		p += strlen(p) + 1;
	}
	argv[argc] = NULL;
	(void) k;
	return main(argc, argv);
}

/* Called by the host to implement the env.__call_sighandler import. */
PGMEM_EXPORT void pgmem_call_sighandler(int fp, int sig)
{
	((void (*)(int)) fp)(sig);
}

/* ---- anonymous mmap without the zero fill ----
 *
 * PostgreSQL maps its shared memory (shared_buffers and the rest, about
 * 40 MB) with one anonymous mmap while a backend starts. Emscripten's
 * mmap takes that block from dlmalloc and memsets all of it, which is the
 * largest single cost of starting a backend, so of every fork, although
 * the memory is fresh from memory.grow and already zero.
 *
 * dlmalloc in Emscripten is built with MORECORE_CANNOT_TRIM, so the
 * program break only ever grows: every byte at or above it has never been
 * handed out and is still zero, and only the part of a block below the
 * old break (the tail of dlmalloc's top chunk, or a recycled chunk) can
 * hold old data. That part is cleared, the rest is used as is.
 *
 * munmap is pgl_munmap in the PostgreSQL objects (-Dmunmap in build.sh),
 * which frees nothing, so nothing has to find these blocks again; file
 * mappings keep going through the libc implementation.
 */
#include <errno.h>
#include <stdint.h>
#include <sys/mman.h>
#include <emscripten/heap.h>

extern void *__mmap(void *addr, size_t len, int prot, int flags, int fd, off_t off);

void *mmap(void *addr, size_t len, int prot, int flags, int fd, off_t off)
{
	if (!(flags & MAP_ANONYMOUS) || addr != NULL)
		return __mmap(addr, len, prot, flags, fd, off);
	uintptr_t brk = (uintptr_t) sbrk(0);
	size_t alloc = (len + 0xFFFF) & ~(size_t) 0xFFFF;
	char *p = emscripten_builtin_memalign(65536, alloc);
	if (p == NULL) {
		errno = ENOMEM;
		return MAP_FAILED;
	}
	if ((uintptr_t) p < brk) {
		size_t dirty = brk - (uintptr_t) p;
		if (dirty > alloc)
			dirty = alloc;
		memset(p, 0, dirty);
	}
	return p;
}
