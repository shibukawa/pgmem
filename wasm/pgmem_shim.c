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

/* ---- host imports ----
 *
 * Every call that touches another process, the shared memory or a socket
 * goes to the host: the process model is emulated there (one module
 * instance per PostgreSQL process, one goroutine each). Results follow
 * the kernel convention: a negative value is -errno.
 */
/* Run "postgres ..." as a child. stdin_path/stdout_path may be "" (inherit). Returns exit code. */
PGMEM_IMPORT("pgmem_run") extern int pgmem_host_run(const char *cmd, const char *stdin_path, const char *stdout_path);
PGMEM_IMPORT("pgmem_getpid") extern int pgmem_host_getpid(void);
PGMEM_IMPORT("pgmem_kill") extern int pgmem_host_kill(int pid, int sig);
PGMEM_IMPORT("pgmem_waitpid") extern int pgmem_host_waitpid(int pid, int *status, int options);
/* Sleep for us microseconds; returns early when a signal is delivered. */
PGMEM_IMPORT("pgmem_usleep") extern void pgmem_host_usleep(int us);
PGMEM_IMPORT("pgmem_shmget") extern int pgmem_host_shmget(int key, int size, int flags);
/* Returns the address the segment is mapped at (always below 2 GiB). */
PGMEM_IMPORT("pgmem_shmat") extern int pgmem_host_shmat(int id, int addr, int flags);
PGMEM_IMPORT("pgmem_shmdt") extern int pgmem_host_shmdt(int addr);
PGMEM_IMPORT("pgmem_shmctl") extern int pgmem_host_shmctl(int id, int cmd, int *segsz, int *nattch);
/* op: 0 init(value) 1 destroy 2 wait 3 trywait 4 post; sem is the sem_t address. */
PGMEM_IMPORT("pgmem_sem") extern int pgmem_host_sem(int op, int sem, int arg);
PGMEM_IMPORT("pgmem_sock_recv") extern int pgmem_host_sock_recv(int fd, void *buf, int n);
PGMEM_IMPORT("pgmem_sock_send") extern int pgmem_host_sock_send(int fd, const void *buf, int n);
/* poll(2) over an array of struct pollfd; returns the ready count. */
PGMEM_IMPORT("pgmem_poll") extern int pgmem_host_poll(void *fds, int nfds, int timeout);

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

/* ---- process model bridge ----
 *
 * PostgreSQL objects are compiled with -Dkill=pgmem_kill and friends
 * (build.sh), so these wrappers see the same prototypes the libc ones
 * have. This file is compiled without the -D overrides and can use the
 * real libc names (raise, sigprocmask).
 */
#include <errno.h>
#include <poll.h>
#include <semaphore.h>
#include <signal.h>
#include <stdint.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <time.h>

static int host_ret(int r)
{
	if (r < 0)
	{
		errno = -r;
		return -1;
	}
	return r;
}

pid_t pgmem_getpid(void) { return pgmem_host_getpid(); }

int pgmem_kill(pid_t pid, int sig) { return host_ret(pgmem_host_kill((int) pid, sig)); }

pid_t pgmem_waitpid(pid_t pid, int *status, int options)
{
	int			st = 0;
	int			r = pgmem_host_waitpid((int) pid, &st, options);

	if (status)
		*status = st;
	return host_ret(r);
}

int pgmem_nanosleep(const struct timespec *req, struct timespec *rem)
{
	long		us = req->tv_sec * 1000000L + req->tv_nsec / 1000;

	pgmem_host_usleep((int) us);
	if (rem)
	{
		rem->tv_sec = 0;
		rem->tv_nsec = 0;
	}
	return 0;
}

int pgmem_shmget(key_t key, size_t size, int flags)
{
	return host_ret(pgmem_host_shmget((int) key, (int) size, flags));
}

void *pgmem_shmat(int id, const void *addr, int flags)
{
	int			r = pgmem_host_shmat(id, (int) (uintptr_t) addr, flags);

	if (r <= 0)
	{
		errno = r < 0 ? -r : EINVAL;
		return (void *) -1;
	}
	return (void *) (uintptr_t) r;
}

int pgmem_shmdt(const void *addr) { return host_ret(pgmem_host_shmdt((int) (uintptr_t) addr)); }

int pgmem_shmctl(int id, int cmd, struct shmid_ds *buf)
{
	int			segsz = 0, nattch = 0;
	int			r = pgmem_host_shmctl(id, cmd, &segsz, &nattch);

	if (r < 0)
		return host_ret(r);
	if (cmd == IPC_STAT && buf != NULL)
	{
		memset(buf, 0, sizeof(*buf));
		buf->shm_segsz = (size_t) segsz;
		buf->shm_nattch = (unsigned long) nattch;
	}
	return 0;
}

int pgmem_sem_init(sem_t *s, int pshared, unsigned value)
{
	return host_ret(pgmem_host_sem(0, (int) (uintptr_t) s, (int) value));
}
int pgmem_sem_destroy(sem_t *s) { return host_ret(pgmem_host_sem(1, (int) (uintptr_t) s, 0)); }
int pgmem_sem_wait(sem_t *s) { return host_ret(pgmem_host_sem(2, (int) (uintptr_t) s, 0)); }
int pgmem_sem_trywait(sem_t *s) { return host_ret(pgmem_host_sem(3, (int) (uintptr_t) s, 0)); }
int pgmem_sem_post(sem_t *s) { return host_ret(pgmem_host_sem(4, (int) (uintptr_t) s, 0)); }

ssize_t pgmem_recv(int fd, void *buf, size_t n, int flags)
{
	return host_ret(pgmem_host_sock_recv(fd, buf, (int) n));
}

ssize_t pgmem_send(int fd, const void *buf, size_t n, int flags)
{
	return host_ret(pgmem_host_sock_send(fd, buf, (int) n));
}

/*
 * poll(2): the host knows which fds are pipes, sockets and files (this is
 * where a process blocks, so it is also where the host delivers signals
 * and timer expiries before returning EINTR).
 */
int pgmem_poll(struct pollfd *fds, nfds_t nfds, int timeout)
{
	return host_ret(pgmem_host_poll(fds, (int) nfds, timeout));
}

/*
 * Signal delivery. The host calls pgmem_raise(sig) on the target instance
 * at a point where it is inside a host call, which is when a real kernel
 * would run the handler too. raise() applies the process's signal mask
 * (sigprocmask is emulated by Emscripten's libc: a blocked signal stays
 * pending until it is unblocked).
 *
 * A signal that arrives before the process has installed its handler
 * would take the default action, which for SIGUSR1 or SIGTERM is to
 * terminate. A real child inherits the postmaster's blocked mask across
 * exec and gets the signal once it unblocks, after installing handlers;
 * a fresh module instance starts with an empty mask, so hold such signals
 * here and re-raise them at the next sigprocmask call, which is how
 * PostgreSQL unblocks signals once its handlers are in place.
 */
static unsigned early_pending;

PGMEM_EXPORT void pgmem_raise(int sig)
{
	struct sigaction sa;

	if (sig <= 0 || sig >= 32)
		return;
	if (sigaction(sig, NULL, &sa) == 0 && !(sa.sa_flags & SA_SIGINFO) && sa.sa_handler == SIG_DFL)
	{
		early_pending |= 1u << sig;
		return;
	}
	raise(sig);
}

int pgmem_sigprocmask(int how, const sigset_t *set, sigset_t *old)
{
	int			r = sigprocmask(how, set, old);
	unsigned	p = early_pending;

	early_pending = 0;
	for (int sig = 1; sig < 32 && p != 0; sig++)
	{
		if (p & (1u << sig))
		{
			p &= ~(1u << sig);
			pgmem_raise(sig);
		}
	}
	return r;
}

PGMEM_EXPORT void pgmem_init(void)
{
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
