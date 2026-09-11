/*
 * pgmem_cryptohash.inc
 *   Replacement body for src/common/cryptohash.c: the hash state lives on
 *   the host (Go's crypto/md5, crypto/sha1, crypto/sha256, crypto/sha512)
 *   and is addressed by an integer handle. Same API and error semantics
 *   as the in-tree implementation.
 */
#ifndef FRONTEND
#include "postgres.h"
#else
#include "postgres_fe.h"
#endif

#include "common/cryptohash.h"

#ifndef FRONTEND
#define ALLOC(size) palloc(size)
#define FREE(ptr) pfree(ptr)
#else
#define ALLOC(size) malloc(size)
#define FREE(ptr) free(ptr)
#endif

#define PGMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))
/* Returns a handle > 0, or 0 if the type is unknown. */
PGMEM_IMPORT("pgmem_hash_create") extern int pgmem_hash_create(int type);
PGMEM_IMPORT("pgmem_hash_reset") extern int pgmem_hash_reset(int handle);
PGMEM_IMPORT("pgmem_hash_update") extern int pgmem_hash_update(int handle, const void *data, size_t len);
/* Writes the digest if len >= digest size; returns the digest size, or -1. */
PGMEM_IMPORT("pgmem_hash_final") extern int pgmem_hash_final(int handle, void *dest, size_t len);
PGMEM_IMPORT("pgmem_hash_free") extern void pgmem_hash_free(int handle);

typedef enum pg_cryptohash_errno
{
	PG_CRYPTOHASH_ERROR_NONE = 0,
	PG_CRYPTOHASH_ERROR_DEST_LEN,
	PG_CRYPTOHASH_ERROR_OOM,
} pg_cryptohash_errno;

struct pg_cryptohash_ctx
{
	pg_cryptohash_type type;
	pg_cryptohash_errno error;
	int			handle;
};

pg_cryptohash_ctx *
pg_cryptohash_create(pg_cryptohash_type type)
{
	pg_cryptohash_ctx *ctx;

	ctx = ALLOC(sizeof(pg_cryptohash_ctx));
	if (ctx == NULL)
		return NULL;
	memset(ctx, 0, sizeof(pg_cryptohash_ctx));
	ctx->type = type;
	ctx->error = PG_CRYPTOHASH_ERROR_NONE;
	ctx->handle = pgmem_hash_create((int) type);
	if (ctx->handle <= 0)
	{
		explicit_bzero(ctx, sizeof(pg_cryptohash_ctx));
		FREE(ctx);
		return NULL;
	}
	return ctx;
}

int
pg_cryptohash_init(pg_cryptohash_ctx *ctx)
{
	if (ctx == NULL)
		return -1;
	pgmem_hash_reset(ctx->handle);
	return 0;
}

int
pg_cryptohash_update(pg_cryptohash_ctx *ctx, const uint8 *data, size_t len)
{
	if (ctx == NULL)
		return -1;
	pgmem_hash_update(ctx->handle, data, len);
	return 0;
}

int
pg_cryptohash_final(pg_cryptohash_ctx *ctx, uint8 *dest, size_t len)
{
	if (ctx == NULL)
		return -1;
	if (pgmem_hash_final(ctx->handle, dest, len) < 0)
	{
		ctx->error = PG_CRYPTOHASH_ERROR_DEST_LEN;
		return -1;
	}
	return 0;
}

void
pg_cryptohash_free(pg_cryptohash_ctx *ctx)
{
	if (ctx == NULL)
		return;
	pgmem_hash_free(ctx->handle);
	explicit_bzero(ctx, sizeof(pg_cryptohash_ctx));
	FREE(ctx);
}

const char *
pg_cryptohash_error(pg_cryptohash_ctx *ctx)
{
	if (ctx == NULL)
		return _("out of memory");
	switch (ctx->error)
	{
		case PG_CRYPTOHASH_ERROR_NONE:
			return _("success");
		case PG_CRYPTOHASH_ERROR_DEST_LEN:
			return _("destination buffer too small");
		case PG_CRYPTOHASH_ERROR_OOM:
			return _("out of memory");
	}
	Assert(false);
	return _("success");
}
