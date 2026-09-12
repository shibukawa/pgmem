/*
 * pgmem_pgcrypto_compress.inc
 *   Replacement body for contrib/pgcrypto/pgp-compress.c: OpenPGP
 *   compressed packets (ZIP = raw deflate, ZLIB) are handled by the host
 *   (Go's compress/flate and compress/zlib) instead of zlib. Compression
 *   streams through a host writer; decompression hands the host the whole
 *   compressed stream and reads the result back in pieces.
 *
 * Derived from contrib/pgcrypto/pgp-compress.c:
 * Copyright (c) 2005 Marko Kreen
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *	  notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *	  notice, this list of conditions and the following disclaimer in the
 *	  documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#include "postgres.h"

#include "pgp.h"
#include "px.h"
#include "utils/memutils.h"
#include "utils/resowner.h"

#define PGMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))

/* Handles > 0; 0 or -1 on failure. See internal/host/crypto.go. */
PGMEM_IMPORT("pgmem_deflate_create") extern int pgmem_deflate_create(int raw, int level);
PGMEM_IMPORT("pgmem_deflate_write") extern int pgmem_deflate_write(int handle, const void *data, int len);
PGMEM_IMPORT("pgmem_deflate_finish") extern int pgmem_deflate_finish(int handle);
PGMEM_IMPORT("pgmem_inflate_all") extern int pgmem_inflate_all(int raw, const void *data, int len);
PGMEM_IMPORT("pgmem_zstream_read") extern int pgmem_zstream_read(int handle, void *out, int cap);
PGMEM_IMPORT("pgmem_zstream_free") extern void pgmem_zstream_free(int handle);

#define ZIP_OUT_BUF 8192
#define ZIP_IN_BLOCK 8192

/*
 * A host stream handle, released through the ResourceOwner on abort the
 * way the cipher handles are (zlib's state was palloc'd and went away with
 * the memory context; a host handle needs an explicit free).
 */
typedef struct PGMEMZHandle
{
	int			handle;
	ResourceOwner owner;
} PGMEMZHandle;

static void ResOwnerReleasePGMEMZHandle(Datum res);

static const ResourceOwnerDesc pgmemzhandle_resowner_desc =
{
	.name = "pgcrypto compression handle",
	.release_phase = RESOURCE_RELEASE_BEFORE_LOCKS,
	.release_priority = RELEASE_PRIO_FIRST,
	.ReleaseResource = ResOwnerReleasePGMEMZHandle,
	.DebugPrint = NULL,
};

static PGMEMZHandle *
zhandle_remember(int handle)
{
	PGMEMZHandle *zh;

	ResourceOwnerEnlarge(CurrentResourceOwner);
	zh = MemoryContextAlloc(TopMemoryContext, sizeof(*zh));
	zh->handle = handle;
	zh->owner = CurrentResourceOwner;
	ResourceOwnerRemember(zh->owner, PointerGetDatum(zh), &pgmemzhandle_resowner_desc);
	return zh;
}

static void
zhandle_free(PGMEMZHandle *zh)
{
	if (zh == NULL)
		return;
	pgmem_zstream_free(zh->handle);
	if (zh->owner != NULL)
		ResourceOwnerForget(zh->owner, PointerGetDatum(zh), &pgmemzhandle_resowner_desc);
	pfree(zh);
}

static void
ResOwnerReleasePGMEMZHandle(Datum res)
{
	PGMEMZHandle *zh = (PGMEMZHandle *) DatumGetPointer(res);

	zh->owner = NULL;
	zhandle_free(zh);
}

/* Move the host's pending output to next until it is drained. */
static int
drain_to(PushFilter *next, PGMEMZHandle *zh, uint8 *buf, int buflen)
{
	int			n,
				res;

	for (;;)
	{
		n = pgmem_zstream_read(zh->handle, buf, buflen);
		if (n < 0)
			return PXE_PGP_COMPRESSION_ERROR;
		if (n == 0)
			return 0;
		res = pushf_write(next, buf, n);
		if (res < 0)
			return res;
	}
}

/*
 * Compressed pkt writer
 */

struct ZipStat
{
	PGMEMZHandle *zh;
	int			buf_len;
	uint8		buf[ZIP_OUT_BUF];
};

static int
compress_init(PushFilter *next, void *init_arg, void **priv_p)
{
	int			handle;
	struct ZipStat *st;
	PGP_Context *ctx = init_arg;
	uint8		type = ctx->compress_algo;

	if (type != PGP_COMPR_ZLIB && type != PGP_COMPR_ZIP)
		return PXE_PGP_UNSUPPORTED_COMPR;

	handle = pgmem_deflate_create(type == PGP_COMPR_ZIP, ctx->compress_level);
	if (handle <= 0)
		return PXE_PGP_COMPRESSION_ERROR;

	st = palloc0(sizeof(*st));
	st->buf_len = ZIP_OUT_BUF;
	st->zh = zhandle_remember(handle);
	*priv_p = st;

	return ZIP_IN_BLOCK;
}

/* writes compressed data packet */

/* can handle zero-len incoming data, but shouldn't */
static int
compress_process(PushFilter *next, void *priv, const uint8 *data, int len)
{
	struct ZipStat *st = priv;

	if (pgmem_deflate_write(st->zh->handle, data, len) < 0)
		return PXE_PGP_COMPRESSION_ERROR;
	return drain_to(next, st->zh, st->buf, st->buf_len);
}

static int
compress_flush(PushFilter *next, void *priv)
{
	struct ZipStat *st = priv;

	if (pgmem_deflate_finish(st->zh->handle) < 0)
		return PXE_PGP_COMPRESSION_ERROR;
	return drain_to(next, st->zh, st->buf, st->buf_len);
}

static void
compress_free(void *priv)
{
	struct ZipStat *st = priv;

	zhandle_free(st->zh);
	px_memset(st, 0, sizeof(*st));
	pfree(st);
}

static const PushFilterOps
			compress_filter = {
	compress_init, compress_process, compress_flush, compress_free
};

int
pgp_compress_filter(PushFilter **res, PGP_Context *ctx, PushFilter *dst)
{
	return pushf_create(res, &compress_filter, ctx, dst);
}

/*
 * Decompress
 */
struct DecomprData
{
	int			buf_len;		/* = ZIP_OUT_BUF */
	int			buf_data;		/* available data */
	uint8	   *pos;
	int			raw;
	PGMEMZHandle *zh;			/* NULL until the stream has been read */
	int			eof;
	uint8		buf[ZIP_OUT_BUF];
};

static int
decompress_init(void **priv_p, void *arg, PullFilter *src)
{
	PGP_Context *ctx = arg;
	struct DecomprData *dec;

	if (ctx->compress_algo != PGP_COMPR_ZLIB
		&& ctx->compress_algo != PGP_COMPR_ZIP)
		return PXE_PGP_UNSUPPORTED_COMPR;

	dec = palloc0(sizeof(*dec));
	dec->buf_len = ZIP_OUT_BUF;
	dec->raw = (ctx->compress_algo == PGP_COMPR_ZIP);
	*priv_p = dec;

	return 0;
}

/*
 * Read the whole compressed stream from src and inflate it on the host.
 * The compressed packet is the tail of the message, so this consumes the
 * same bytes the streaming reader would, including the check that
 * nothing follows the end of the stream.
 */
static int
decompress_load(struct DecomprData *dec, PullFilter *src)
{
	StringInfoData all;
	uint8	   *tmp;
	int			res;
	int			handle;

	initStringInfo(&all);
	for (;;)
	{
		res = pullf_read(src, 8192, &tmp);
		if (res < 0)
		{
			pfree(all.data);
			return res;
		}
		if (res == 0)
			break;
		appendBinaryStringInfo(&all, (const char *) tmp, res);
	}
	handle = pgmem_inflate_all(dec->raw, all.data, all.len);
	px_memset(all.data, 0, all.len);
	pfree(all.data);
	if (handle <= 0)
	{
		px_debug("decompress_read: inflate error");
		return PXE_PGP_CORRUPT_DATA;
	}
	dec->zh = zhandle_remember(handle);
	return 0;
}

static int
decompress_read(void *priv, PullFilter *src, int len,
				uint8 **data_p, uint8 *buf, int buflen)
{
	int			res;
	struct DecomprData *dec = priv;

restart:
	if (dec->buf_data > 0)
	{
		if (len > dec->buf_data)
			len = dec->buf_data;
		*data_p = dec->pos;
		dec->pos += len;
		dec->buf_data -= len;
		return len;
	}

	if (dec->eof)
		return 0;

	if (dec->zh == NULL)
	{
		res = decompress_load(dec, src);
		if (res < 0)
			return res;
	}

	res = pgmem_zstream_read(dec->zh->handle, dec->buf, dec->buf_len);
	if (res < 0)
		return PXE_PGP_CORRUPT_DATA;
	dec->pos = dec->buf;
	dec->buf_data = res;
	if (res == 0)
		dec->eof = 1;
	goto restart;
}

static void
decompress_free(void *priv)
{
	struct DecomprData *dec = priv;

	zhandle_free(dec->zh);
	px_memset(dec, 0, sizeof(*dec));
	pfree(dec);
}

static const PullFilterOps
			decompress_filter = {
	decompress_init, decompress_read, decompress_free
};

int
pgp_decompress_filter(PullFilter **res, PGP_Context *ctx, PullFilter *src)
{
	return pullf_create(res, &decompress_filter, ctx, src);
}
