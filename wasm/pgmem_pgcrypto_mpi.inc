/*
 * pgmem_pgcrypto_mpi.inc
 *   Replacement body for contrib/pgcrypto/pgp-mpi-openssl.c: the modular
 *   arithmetic behind OpenPGP public-key operations runs on the host
 *   (Go's math/big) over big-endian byte strings.
 *
 * Derived from contrib/pgcrypto/pgp-mpi-openssl.c:
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

#define PGMEM_IMPORT(name) __attribute__((import_module("env"), import_name(name)))

/*
 * pgmem_bn_op computes one of
 *   1: a^b mod m      2: a*b mod m      3: a^-1 mod m (b unused)
 * over big-endian unsigned integers and writes the result, without leading
 * zero bytes, to out (capacity cap). It returns the result length, 0 for a
 * zero result, or -1 when the operation fails (m == 0, no inverse, cap too
 * small).
 */
PGMEM_IMPORT("pgmem_bn_op") extern int pgmem_bn_op(int op, const void *a, int alen, const void *b, int blen, const void *m, int mlen, void *out, int cap);
/*
 * pgmem_bn_rand writes a random integer of exactly bits bits (top bit set,
 * like BN_rand(.., top=0, bottom=0)) and returns its byte length, or -1.
 */
PGMEM_IMPORT("pgmem_bn_rand") extern int pgmem_bn_rand(int bits, void *out, int cap);

#define PGMEM_BN_MODEXP	1
#define PGMEM_BN_MODMUL	2
#define PGMEM_BN_MODINV	3

/* Number of significant bits in a big-endian byte string. */
static int
bytes_num_bits(const uint8 *data, int bytes)
{
	int			i;
	int			bits;
	uint8		top;

	for (i = 0; i < bytes && data[i] == 0; i++)
		;
	if (i == bytes)
		return 0;
	bits = (bytes - i - 1) * 8;
	for (top = data[i]; top; top >>= 1)
		bits++;
	return bits;
}

/*
 * mpi_check mirrors mpi_to_bn(): an MPI whose declared bit count does not
 * match its data is rejected.
 */
static bool
mpi_check(PGP_MPI *n)
{
	if (bytes_num_bits(n->data, n->bytes) != n->bits)
	{
		px_debug("mpi_check: bignum conversion failed: mpi=%d, bn=%d",
				 n->bits, bytes_num_bits(n->data, n->bytes));
		return false;
	}
	return true;
}

/*
 * bytes_to_mpi mirrors bn_to_mpi(): the host's minimal big-endian result
 * becomes an MPI with an exact bit count.
 */
static PGP_MPI *
bytes_to_mpi(const uint8 *data, int len)
{
	int			res;
	int			bits = bytes_num_bits(data, len);
	PGP_MPI    *n;

	res = pgp_mpi_alloc(bits, &n);
	if (res < 0)
		return NULL;

	if (len != n->bytes)
	{
		px_debug("bytes_to_mpi: bignum conversion failed: bn=%d, mpi=%d",
				 len, n->bytes);
		pgp_mpi_free(n);
		return NULL;
	}
	memcpy(n->data, data, len);
	return n;
}

/* A scratch buffer that is wiped before it is freed. */
static uint8 *
bn_buf_alloc(int len)
{
	return palloc(len > 0 ? len : 1);
}

static void
bn_buf_free(uint8 *buf, int len)
{
	if (buf == NULL)
		return;
	px_memset(buf, 0, len > 0 ? len : 1);
	pfree(buf);
}

/*
 * Decide the number of bits in the random component k
 *
 * It should be in the same range as p for signing (which
 * is deprecated), but can be much smaller for encrypting.
 *
 * Until I research it further, I just mimic gpg behaviour.
 * It has a special mapping table, for values <= 5120,
 * above that it uses 'arbitrary high number'.  Following
 * algorithm hovers 10-70 bits above gpg values.  And for
 * larger p, it uses gpg's algorithm.
 *
 * The point is - if k gets large, encryption will be
 * really slow.  It does not matter for decryption.
 */
static int
decide_k_bits(int p_bits)
{
	if (p_bits <= 5120)
		return p_bits / 10 + 160;
	else
		return (p_bits / 8 + 200) * 3 / 2;
}

int
pgp_elgamal_encrypt(PGP_PubKey *pk, PGP_MPI *_m,
					PGP_MPI **c1_p, PGP_MPI **c2_p)
{
	int			res = PXE_PGP_MATH_FAILED;
	int			k_bits;
	PGP_MPI    *p = pk->pub.elg.p;
	PGP_MPI    *g = pk->pub.elg.g;
	PGP_MPI    *y = pk->pub.elg.y;
	int			cap = p->bytes;
	uint8	   *k = NULL;
	uint8	   *yk = NULL;
	uint8	   *c1 = NULL;
	uint8	   *c2 = NULL;
	int			klen,
				yklen,
				c1len,
				c2len;
	int			kcap;

	if (!mpi_check(_m) || !mpi_check(p) || !mpi_check(g) || !mpi_check(y))
		goto err;

	/*
	 * generate k
	 */
	k_bits = decide_k_bits(p->bits);
	kcap = (k_bits + 7) / 8;
	k = bn_buf_alloc(kcap);
	klen = pgmem_bn_rand(k_bits, k, kcap);
	if (klen < 0)
		goto err;

	/*
	 * c1 = g^k c2 = m * y^k
	 */
	c1 = bn_buf_alloc(cap);
	yk = bn_buf_alloc(cap);
	c2 = bn_buf_alloc(cap);
	c1len = pgmem_bn_op(PGMEM_BN_MODEXP, g->data, g->bytes, k, klen, p->data, p->bytes, c1, cap);
	if (c1len < 0)
		goto err;
	yklen = pgmem_bn_op(PGMEM_BN_MODEXP, y->data, y->bytes, k, klen, p->data, p->bytes, yk, cap);
	if (yklen < 0)
		goto err;
	c2len = pgmem_bn_op(PGMEM_BN_MODMUL, _m->data, _m->bytes, yk, yklen, p->data, p->bytes, c2, cap);
	if (c2len < 0)
		goto err;

	/* result */
	*c1_p = bytes_to_mpi(c1, c1len);
	*c2_p = bytes_to_mpi(c2, c2len);
	if (*c1_p && *c2_p)
		res = 0;
err:
	bn_buf_free(c2, cap);
	bn_buf_free(yk, cap);
	bn_buf_free(c1, cap);
	if (k)
		bn_buf_free(k, kcap);
	return res;
}

int
pgp_elgamal_decrypt(PGP_PubKey *pk, PGP_MPI *_c1, PGP_MPI *_c2,
					PGP_MPI **msg_p)
{
	int			res = PXE_PGP_MATH_FAILED;
	PGP_MPI    *p = pk->pub.elg.p;
	PGP_MPI    *x = pk->sec.elg.x;
	int			cap = p->bytes;
	uint8	   *c1x = NULL;
	uint8	   *div = NULL;
	uint8	   *m = NULL;
	int			c1xlen,
				divlen,
				mlen;

	if (!mpi_check(_c1) || !mpi_check(_c2) || !mpi_check(p) || !mpi_check(x))
		goto err;

	/*
	 * m = c2 / (c1^x)
	 */
	c1x = bn_buf_alloc(cap);
	div = bn_buf_alloc(cap);
	m = bn_buf_alloc(cap);
	c1xlen = pgmem_bn_op(PGMEM_BN_MODEXP, _c1->data, _c1->bytes, x->data, x->bytes, p->data, p->bytes, c1x, cap);
	if (c1xlen < 0)
		goto err;
	divlen = pgmem_bn_op(PGMEM_BN_MODINV, c1x, c1xlen, NULL, 0, p->data, p->bytes, div, cap);
	if (divlen < 0)
		goto err;
	mlen = pgmem_bn_op(PGMEM_BN_MODMUL, _c2->data, _c2->bytes, div, divlen, p->data, p->bytes, m, cap);
	if (mlen < 0)
		goto err;

	/* result */
	*msg_p = bytes_to_mpi(m, mlen);
	if (*msg_p)
		res = 0;
err:
	bn_buf_free(m, cap);
	bn_buf_free(div, cap);
	bn_buf_free(c1x, cap);
	return res;
}

int
pgp_rsa_encrypt(PGP_PubKey *pk, PGP_MPI *_m, PGP_MPI **c_p)
{
	int			res = PXE_PGP_MATH_FAILED;
	PGP_MPI    *e = pk->pub.rsa.e;
	PGP_MPI    *n = pk->pub.rsa.n;
	int			cap = n->bytes;
	uint8	   *c = NULL;
	int			clen;

	if (!mpi_check(_m) || !mpi_check(e) || !mpi_check(n))
		goto err;

	/*
	 * c = m ^ e
	 */
	c = bn_buf_alloc(cap);
	clen = pgmem_bn_op(PGMEM_BN_MODEXP, _m->data, _m->bytes, e->data, e->bytes, n->data, n->bytes, c, cap);
	if (clen < 0)
		goto err;

	*c_p = bytes_to_mpi(c, clen);
	if (*c_p)
		res = 0;
err:
	bn_buf_free(c, cap);
	return res;
}

int
pgp_rsa_decrypt(PGP_PubKey *pk, PGP_MPI *_c, PGP_MPI **m_p)
{
	int			res = PXE_PGP_MATH_FAILED;
	PGP_MPI    *d = pk->sec.rsa.d;
	PGP_MPI    *n = pk->pub.rsa.n;
	int			cap = n->bytes;
	uint8	   *m = NULL;
	int			mlen;

	if (!mpi_check(_c) || !mpi_check(d) || !mpi_check(n))
		goto err;

	/*
	 * m = c ^ d
	 */
	m = bn_buf_alloc(cap);
	mlen = pgmem_bn_op(PGMEM_BN_MODEXP, _c->data, _c->bytes, d->data, d->bytes, n->data, n->bytes, m, cap);
	if (mlen < 0)
		goto err;

	*m_p = bytes_to_mpi(m, mlen);
	if (*m_p)
		res = 0;
err:
	bn_buf_free(m, cap);
	return res;
}
