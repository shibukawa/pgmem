/*
 * pgmem_pgcrypto_openssl.inc
 *   Replacement body for contrib/pgcrypto/openssl.c: digests and ciphers
 *   run on the host (Go's crypto packages and golang.org/x/crypto) and are addressed
 *   by integer handles. Same PX_MD / PX_Cipher contract, error codes and
 *   ResourceOwner cleanup as the OpenSSL implementation.
 *
 * Derived from contrib/pgcrypto/openssl.c:
 * Copyright (c) 2001 Marko Kreen
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

#include "px.h"
#include "utils/memutils.h"
#include "utils/resowner.h"

#define PGMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))

/* Digest contexts: the same host functions src/common/cryptohash.c uses. */
PGMEM_IMPORT("pgmem_hash_create") extern int pgmem_hash_create(int type);
PGMEM_IMPORT("pgmem_hash_reset") extern int pgmem_hash_reset(int handle);
PGMEM_IMPORT("pgmem_hash_update") extern int pgmem_hash_update(int handle, const void *data, size_t len);
PGMEM_IMPORT("pgmem_hash_final") extern int pgmem_hash_final(int handle, void *dest, size_t len);
PGMEM_IMPORT("pgmem_hash_free") extern void pgmem_hash_free(int handle);
/* Returns block_size << 16 | result_size, or -1 for an unknown handle. */
PGMEM_IMPORT("pgmem_hash_info") extern int pgmem_hash_info(int handle);

/*
 * Cipher contexts. create returns a handle > 0 or 0 when the host cannot
 * build the cipher from that key (wrong length for the algorithm). run
 * writes the output to res (capacity rescap) and returns its length, or
 * -1 when the operation fails (bad padding, length not a multiple of the
 * block size). Like an OpenSSL EVP context, a handle is bound to one
 * direction by its first run and keeps chaining state between runs.
 */
PGMEM_IMPORT("pgmem_cipher_create") extern int pgmem_cipher_create(int algo, int mode, const void *key, int klen, const void *iv, int ivlen);
PGMEM_IMPORT("pgmem_cipher_run") extern int pgmem_cipher_run(int handle, int encrypt, int padding, const void *data, int dlen, void *res, int rescap);
PGMEM_IMPORT("pgmem_cipher_free") extern void pgmem_cipher_free(int handle);

/* Algorithm and mode codes shared with the host (internal/host). */
#define PGMEM_CIPHER_BF		1
#define PGMEM_CIPHER_DES	2
#define PGMEM_CIPHER_DES3	3
#define PGMEM_CIPHER_CAST5	4
#define PGMEM_CIPHER_AES	5

#define PGMEM_MODE_ECB		1
#define PGMEM_MODE_CBC		2
#define PGMEM_MODE_CFB		3

/*
 * Max lengths we might want to handle.
 */
#define MAX_KEY		(512/8)
#define MAX_IV		(128/8)

/*
 * Hashes
 */

/*
 * Digest names as OpenSSL's EVP_get_digestbyname() accepts them, mapped to
 * the host hash types (0-5 are pg_cryptohash_type, the rest are pgmem's).
 */
static const struct
{
	const char *name;
	int			type;
}			pgmem_digest_names[] = {
	{"md5", 0},
	{"sha1", 1},
	{"sha-1", 1},
	{"sha224", 2},
	{"sha-224", 2},
	{"sha2-224", 2},
	{"sha256", 3},
	{"sha-256", 3},
	{"sha2-256", 3},
	{"sha384", 4},
	{"sha-384", 4},
	{"sha2-384", 4},
	{"sha512", 5},
	{"sha-512", 5},
	{"sha2-512", 5},
	{"ripemd160", 6},
	{"ripemd-160", 6},
	{"rmd160", 6},
	{"sha3-224", 7},
	{"sha3-256", 8},
	{"sha3-384", 9},
	{"sha3-512", 10},
	{NULL, 0}
};

/*
 * To make sure we don't leak host handles, we use the ResourceOwner
 * mechanism to free them on abort.
 */
typedef struct PGMEMDigest
{
	int			handle;
	unsigned	result_size;
	unsigned	block_size;

	ResourceOwner owner;
} PGMEMDigest;

/* ResourceOwner callbacks to hold digest handles */
static void ResOwnerReleasePGMEMDigest(Datum res);

static const ResourceOwnerDesc pgmemdigest_resowner_desc =
{
	.name = "pgcrypto digest handle",
	.release_phase = RESOURCE_RELEASE_BEFORE_LOCKS,
	.release_priority = RELEASE_PRIO_FIRST,
	.ReleaseResource = ResOwnerReleasePGMEMDigest,
	.DebugPrint = NULL,			/* default message is fine */
};

/* Convenience wrappers over ResourceOwnerRemember/Forget */
static inline void
ResourceOwnerRememberPGMEMDigest(ResourceOwner owner, PGMEMDigest *digest)
{
	ResourceOwnerRemember(owner, PointerGetDatum(digest), &pgmemdigest_resowner_desc);
}
static inline void
ResourceOwnerForgetPGMEMDigest(ResourceOwner owner, PGMEMDigest *digest)
{
	ResourceOwnerForget(owner, PointerGetDatum(digest), &pgmemdigest_resowner_desc);
}

static void
free_pgmem_digest(PGMEMDigest *digest)
{
	pgmem_hash_free(digest->handle);
	if (digest->owner != NULL)
		ResourceOwnerForgetPGMEMDigest(digest->owner, digest);
	pfree(digest);
}

static unsigned
digest_result_size(PX_MD *h)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	return digest->result_size;
}

static unsigned
digest_block_size(PX_MD *h)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	return digest->block_size;
}

static void
digest_reset(PX_MD *h)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	if (pgmem_hash_reset(digest->handle) != 0)
		elog(ERROR, "pgmem_hash_reset() failed");
}

static void
digest_update(PX_MD *h, const uint8 *data, unsigned dlen)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	if (pgmem_hash_update(digest->handle, data, dlen) != 0)
		elog(ERROR, "pgmem_hash_update() failed");
}

static void
digest_finish(PX_MD *h, uint8 *dst)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	if (pgmem_hash_final(digest->handle, dst, digest->result_size) < 0)
		elog(ERROR, "pgmem_hash_final() failed");
}

static void
digest_free(PX_MD *h)
{
	PGMEMDigest *digest = (PGMEMDigest *) h->p.ptr;

	free_pgmem_digest(digest);
	pfree(h);
}

/* PUBLIC functions */

int
px_find_digest(const char *name, PX_MD **res)
{
	int			i;
	int			handle;
	int			info;
	PX_MD	   *h;
	PGMEMDigest *digest;

	for (i = 0; pgmem_digest_names[i].name; i++)
		if (pg_strcasecmp(pgmem_digest_names[i].name, name) == 0)
			break;
	if (pgmem_digest_names[i].name == NULL)
		return PXE_NO_HASH;

	ResourceOwnerEnlarge(CurrentResourceOwner);

	/*
	 * Create a PGMEMDigest object, a host hash handle, and a PX_MD object.
	 * The order is crucial, to make sure we don't leak anything on
	 * out-of-memory or other error.
	 */
	digest = MemoryContextAlloc(TopMemoryContext, sizeof(*digest));

	handle = pgmem_hash_create(pgmem_digest_names[i].type);
	if (handle <= 0)
	{
		pfree(digest);
		return PXE_NO_HASH;
	}
	info = pgmem_hash_info(handle);
	if (info < 0)
	{
		pgmem_hash_free(handle);
		pfree(digest);
		return PXE_CIPHER_INIT;
	}

	digest->handle = handle;
	digest->result_size = info & 0xffff;
	digest->block_size = (info >> 16) & 0xffff;
	digest->owner = CurrentResourceOwner;
	ResourceOwnerRememberPGMEMDigest(digest->owner, digest);

	/* The PX_MD object is allocated in the current memory context. */
	h = palloc(sizeof(*h));
	h->result_size = digest_result_size;
	h->block_size = digest_block_size;
	h->reset = digest_reset;
	h->update = digest_update;
	h->finish = digest_finish;
	h->free = digest_free;
	h->p.ptr = digest;

	*res = h;
	return 0;
}

/* ResourceOwner callbacks for PGMEMDigest */

static void
ResOwnerReleasePGMEMDigest(Datum res)
{
	PGMEMDigest *digest = (PGMEMDigest *) DatumGetPointer(res);

	digest->owner = NULL;
	free_pgmem_digest(digest);
}

/*
 * Ciphers
 */

/*
 * pgmem_cipher contains the static information about each cipher.
 */
struct pgmem_cipher
{
	int			(*init) (PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv);
	int			algo;
	int			mode;
	int			block_size;
	int			max_key_size;
};

/*
 * PGMEMCipher contains the state for using a cipher. A separate PGMEMCipher
 * object is allocated in each px_find_cipher() call. The host handle is
 * created lazily by the first encrypt/decrypt, like the EVP context was.
 *
 * To make sure we don't leak host handles, we use the ResourceOwner
 * mechanism to free them on abort.
 */
typedef struct PGMEMCipher
{
	int			handle;
	uint8		key[MAX_KEY];
	uint8		iv[MAX_IV];
	unsigned	klen;
	unsigned	init;
	const struct pgmem_cipher *ciph;

	ResourceOwner owner;
} PGMEMCipher;

/* ResourceOwner callbacks to hold cipher state */
static void ResOwnerReleasePGMEMCipher(Datum res);

static const ResourceOwnerDesc pgmemcipher_resowner_desc =
{
	.name = "pgcrypto cipher handle",
	.release_phase = RESOURCE_RELEASE_BEFORE_LOCKS,
	.release_priority = RELEASE_PRIO_FIRST,
	.ReleaseResource = ResOwnerReleasePGMEMCipher,
	.DebugPrint = NULL,			/* default message is fine */
};

/* Convenience wrappers over ResourceOwnerRemember/Forget */
static inline void
ResourceOwnerRememberPGMEMCipher(ResourceOwner owner, PGMEMCipher *od)
{
	ResourceOwnerRemember(owner, PointerGetDatum(od), &pgmemcipher_resowner_desc);
}
static inline void
ResourceOwnerForgetPGMEMCipher(ResourceOwner owner, PGMEMCipher *od)
{
	ResourceOwnerForget(owner, PointerGetDatum(od), &pgmemcipher_resowner_desc);
}

static void
free_pgmem_cipher(PGMEMCipher *od)
{
	if (od->handle > 0)
		pgmem_cipher_free(od->handle);
	if (od->owner != NULL)
		ResourceOwnerForgetPGMEMCipher(od->owner, od);
	px_memset(od, 0, sizeof(*od));
	pfree(od);
}

/* Common routines for all algorithms */

static unsigned
gen_pgmem_block_size(PX_Cipher *c)
{
	PGMEMCipher *od = (PGMEMCipher *) c->ptr;

	return od->ciph->block_size;
}

static unsigned
gen_pgmem_key_size(PX_Cipher *c)
{
	PGMEMCipher *od = (PGMEMCipher *) c->ptr;

	return od->ciph->max_key_size;
}

static unsigned
gen_pgmem_iv_size(PX_Cipher *c)
{
	unsigned	ivlen;
	PGMEMCipher *od = (PGMEMCipher *) c->ptr;

	ivlen = od->ciph->block_size;
	return ivlen;
}

static void
gen_pgmem_free(PX_Cipher *c)
{
	PGMEMCipher *od = (PGMEMCipher *) c->ptr;

	free_pgmem_cipher(od);
	pfree(c);
}

static int
gen_pgmem_init_handle(PGMEMCipher *od)
{
	int			handle;

	if (od->init)
		return 0;
	handle = pgmem_cipher_create(od->ciph->algo, od->ciph->mode,
								 od->key, od->klen,
								 od->iv, od->ciph->block_size);
	if (handle <= 0)
		return PXE_CIPHER_INIT;
	od->handle = handle;
	od->init = true;
	return 0;
}

static int
gen_pgmem_decrypt(PX_Cipher *c, int padding, const uint8 *data, unsigned dlen,
				  uint8 *res, unsigned *rlen)
{
	PGMEMCipher *od = c->ptr;
	int			err;
	int			outlen;

	err = gen_pgmem_init_handle(od);
	if (err)
		return err;

	/* Decryption never produces more than it consumes. */
	outlen = pgmem_cipher_run(od->handle, 0, padding, data, dlen, res, dlen);
	if (outlen < 0)
		return PXE_DECRYPT_FAILED;
	*rlen = outlen;

	return 0;
}

static int
gen_pgmem_encrypt(PX_Cipher *c, int padding, const uint8 *data, unsigned dlen,
				  uint8 *res, unsigned *rlen)
{
	PGMEMCipher *od = c->ptr;
	int			err;
	int			outlen;

	err = gen_pgmem_init_handle(od);
	if (err)
		return err;

	/* Padding adds at most one block (EVP's documented bound). */
	outlen = pgmem_cipher_run(od->handle, 1, padding, data, dlen, res,
							  dlen + od->ciph->block_size);
	if (outlen < 0)
		return PXE_ENCRYPT_FAILED;
	*rlen = outlen;

	return 0;
}

/* Blowfish */

static int
bf_init(PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv)
{
	PGMEMCipher *od = c->ptr;
	unsigned	bs = gen_pgmem_block_size(c);

	/* Go's blowfish takes the full 448-bit key range; no strength probe. */
	od->klen = klen;
	memcpy(od->key, key, klen);

	if (iv)
		memcpy(od->iv, iv, bs);
	else
		memset(od->iv, 0, bs);
	return 0;
}

/* DES */

static int
pgmem_des_init(PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv)
{
	PGMEMCipher *od = c->ptr;
	unsigned	bs = gen_pgmem_block_size(c);

	od->klen = 8;
	memset(od->key, 0, 8);
	memcpy(od->key, key, klen > 8 ? 8 : klen);

	if (iv)
		memcpy(od->iv, iv, bs);
	else
		memset(od->iv, 0, bs);
	return 0;
}

/* DES3 */

static int
pgmem_des3_init(PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv)
{
	PGMEMCipher *od = c->ptr;
	unsigned	bs = gen_pgmem_block_size(c);

	od->klen = 24;
	memset(od->key, 0, 24);
	memcpy(od->key, key, klen > 24 ? 24 : klen);

	if (iv)
		memcpy(od->iv, iv, bs);
	else
		memset(od->iv, 0, bs);
	return 0;
}

/* CAST5 */

static int
pgmem_cast_init(PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv)
{
	PGMEMCipher *od = c->ptr;
	unsigned	bs = gen_pgmem_block_size(c);

	od->klen = klen;
	memcpy(od->key, key, klen);

	if (iv)
		memcpy(od->iv, iv, bs);
	else
		memset(od->iv, 0, bs);
	return 0;
}

/* AES */

static int
pgmem_aes_init(PX_Cipher *c, const uint8 *key, unsigned klen, const uint8 *iv)
{
	PGMEMCipher *od = c->ptr;
	unsigned	bs = gen_pgmem_block_size(c);

	/*
	 * The key is rounded up to the next AES key size; the bytes beyond klen
	 * are zero because od was allocated zeroed (OpenSSL saw the same).
	 */
	if (klen <= 128 / 8)
		od->klen = 128 / 8;
	else if (klen <= 192 / 8)
		od->klen = 192 / 8;
	else if (klen <= 256 / 8)
		od->klen = 256 / 8;
	else
		return PXE_KEY_TOO_BIG;

	memcpy(od->key, key, klen);

	if (iv)
		memcpy(od->iv, iv, bs);
	else
		memset(od->iv, 0, bs);

	return 0;
}

/*
 * aliases
 */

static PX_Alias pgmem_aliases[] = {
	{"bf", "bf-cbc"},
	{"blowfish", "bf-cbc"},
	{"blowfish-cbc", "bf-cbc"},
	{"blowfish-ecb", "bf-ecb"},
	{"blowfish-cfb", "bf-cfb"},
	{"des", "des-cbc"},
	{"3des", "des3-cbc"},
	{"3des-ecb", "des3-ecb"},
	{"3des-cbc", "des3-cbc"},
	{"cast5", "cast5-cbc"},
	{"aes", "aes-cbc"},
	{"rijndael", "aes-cbc"},
	{"rijndael-cbc", "aes-cbc"},
	{"rijndael-ecb", "aes-ecb"},
	{"rijndael-cfb", "aes-cfb"},
	{NULL}
};

static const struct pgmem_cipher pgmem_bf_cbc = {
	bf_init, PGMEM_CIPHER_BF, PGMEM_MODE_CBC, 64 / 8, 448 / 8
};

static const struct pgmem_cipher pgmem_bf_ecb = {
	bf_init, PGMEM_CIPHER_BF, PGMEM_MODE_ECB, 64 / 8, 448 / 8
};

static const struct pgmem_cipher pgmem_bf_cfb = {
	bf_init, PGMEM_CIPHER_BF, PGMEM_MODE_CFB, 64 / 8, 448 / 8
};

static const struct pgmem_cipher pgmem_des_ecb = {
	pgmem_des_init, PGMEM_CIPHER_DES, PGMEM_MODE_ECB, 64 / 8, 64 / 8
};

static const struct pgmem_cipher pgmem_des_cbc = {
	pgmem_des_init, PGMEM_CIPHER_DES, PGMEM_MODE_CBC, 64 / 8, 64 / 8
};

static const struct pgmem_cipher pgmem_des3_ecb = {
	pgmem_des3_init, PGMEM_CIPHER_DES3, PGMEM_MODE_ECB, 64 / 8, 192 / 8
};

static const struct pgmem_cipher pgmem_des3_cbc = {
	pgmem_des3_init, PGMEM_CIPHER_DES3, PGMEM_MODE_CBC, 64 / 8, 192 / 8
};

static const struct pgmem_cipher pgmem_cast_ecb = {
	pgmem_cast_init, PGMEM_CIPHER_CAST5, PGMEM_MODE_ECB, 64 / 8, 128 / 8
};

static const struct pgmem_cipher pgmem_cast_cbc = {
	pgmem_cast_init, PGMEM_CIPHER_CAST5, PGMEM_MODE_CBC, 64 / 8, 128 / 8
};

static const struct pgmem_cipher pgmem_aes_ecb = {
	pgmem_aes_init, PGMEM_CIPHER_AES, PGMEM_MODE_ECB, 128 / 8, 256 / 8
};

static const struct pgmem_cipher pgmem_aes_cbc = {
	pgmem_aes_init, PGMEM_CIPHER_AES, PGMEM_MODE_CBC, 128 / 8, 256 / 8
};

static const struct pgmem_cipher pgmem_aes_cfb = {
	pgmem_aes_init, PGMEM_CIPHER_AES, PGMEM_MODE_CFB, 128 / 8, 256 / 8
};

/*
 * Special handlers
 */
struct pgmem_cipher_lookup
{
	const char *name;
	const struct pgmem_cipher *ciph;
};

static const struct pgmem_cipher_lookup pgmem_cipher_types[] = {
	{"bf-cbc", &pgmem_bf_cbc},
	{"bf-ecb", &pgmem_bf_ecb},
	{"bf-cfb", &pgmem_bf_cfb},
	{"des-ecb", &pgmem_des_ecb},
	{"des-cbc", &pgmem_des_cbc},
	{"des3-ecb", &pgmem_des3_ecb},
	{"des3-cbc", &pgmem_des3_cbc},
	{"cast5-ecb", &pgmem_cast_ecb},
	{"cast5-cbc", &pgmem_cast_cbc},
	{"aes-ecb", &pgmem_aes_ecb},
	{"aes-cbc", &pgmem_aes_cbc},
	{"aes-cfb", &pgmem_aes_cfb},
	{NULL}
};

/* PUBLIC functions */

int
px_find_cipher(const char *name, PX_Cipher **res)
{
	const struct pgmem_cipher_lookup *i;
	PX_Cipher  *c = NULL;
	PGMEMCipher *od;

	name = px_resolve_alias(pgmem_aliases, name);
	for (i = pgmem_cipher_types; i->name; i++)
		if (strcmp(i->name, name) == 0)
			break;
	if (i->name == NULL)
		return PXE_NO_CIPHER;

	ResourceOwnerEnlarge(CurrentResourceOwner);

	/*
	 * Create a PGMEMCipher object and a PX_Cipher. The host handle comes
	 * later, once the key and direction are known.
	 */
	od = MemoryContextAllocZero(TopMemoryContext, sizeof(*od));
	od->ciph = i->ciph;
	od->owner = CurrentResourceOwner;
	ResourceOwnerRememberPGMEMCipher(od->owner, od);

	/* The PX_Cipher is allocated in current memory context */
	c = palloc(sizeof(*c));
	c->block_size = gen_pgmem_block_size;
	c->key_size = gen_pgmem_key_size;
	c->iv_size = gen_pgmem_iv_size;
	c->free = gen_pgmem_free;
	c->init = od->ciph->init;
	c->encrypt = gen_pgmem_encrypt;
	c->decrypt = gen_pgmem_decrypt;
	c->ptr = od;

	*res = c;
	return 0;
}

/* ResourceOwner callbacks for PGMEMCipher */

static void
ResOwnerReleasePGMEMCipher(Datum res)
{
	PGMEMCipher *od = (PGMEMCipher *) DatumGetPointer(res);

	od->owner = NULL;
	free_pgmem_cipher(od);
}

/*
 * CheckFIPSMode
 *
 * There is no OpenSSL underneath, so there is no FIPS mode either.
 */
bool
CheckFIPSMode(void)
{
	return false;
}

/*
 * CheckBuiltinCryptoMode
 *
 * Function for erroring out in case built-in crypto is executed when the user
 * has disabled it. "fips" never errors here since FIPS mode is never on.
 */
void
CheckBuiltinCryptoMode(void)
{
	if (builtin_crypto_enabled == BC_ON)
		return;

	if (builtin_crypto_enabled == BC_OFF)
		ereport(ERROR,
				errmsg("use of built-in crypto functions is disabled"));

	Assert(builtin_crypto_enabled == BC_FIPS);
}
