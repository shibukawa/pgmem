/*
 * pgmem_uuid.c
 *   libuuid replacement for contrib/uuid-ossp: version 4 from the host's
 *   crypto/rand, version 1 from the clock with a random multicast node
 *   (there is no MAC address to read), like libuuid falls back to.
 */
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <sys/time.h>

#include "uuid/uuid.h"

extern int pgmem_random_bytes(void *buf, size_t len)
	__attribute__((import_module("env"), import_name("pgmem_random_bytes")));

void
uuid_generate_random(uuid_t out)
{
	pgmem_random_bytes(out, 16);
	out[6] = (out[6] & 0x0f) | 0x40;	/* version 4 */
	out[8] = (out[8] & 0x3f) | 0x80;	/* RFC 4122 variant */
}

/* 100 ns intervals between 1582-10-15 and 1970-01-01 */
#define UUID_EPOCH_OFFSET UINT64_C(0x01B21DD213814000)

void
uuid_generate_time(uuid_t out)
{
	static unsigned char node[6];
	static int	node_init = 0;
	static uint16_t clock_seq;
	static uint64_t last_ts;
	struct timeval tv;
	uint64_t	ts;

	if (!node_init)
	{
		pgmem_random_bytes(node, sizeof(node));
		node[0] |= 0x01;		/* multicast bit: not a real MAC */
		pgmem_random_bytes(&clock_seq, sizeof(clock_seq));
		clock_seq &= 0x3fff;
		node_init = 1;
	}
	gettimeofday(&tv, NULL);
	ts = (uint64_t) tv.tv_sec * 10000000 + (uint64_t) tv.tv_usec * 10 + UUID_EPOCH_OFFSET;
	if (ts <= last_ts)
		ts = last_ts + 1;		/* keep them unique within the resolution */
	last_ts = ts;

	out[0] = (unsigned char) (ts >> 24);	/* time_low */
	out[1] = (unsigned char) (ts >> 16);
	out[2] = (unsigned char) (ts >> 8);
	out[3] = (unsigned char) ts;
	out[4] = (unsigned char) (ts >> 40);	/* time_mid */
	out[5] = (unsigned char) (ts >> 32);
	out[6] = (unsigned char) (((ts >> 56) & 0x0f) | 0x10);	/* time_hi + version 1 */
	out[7] = (unsigned char) (ts >> 48);
	out[8] = (unsigned char) (((clock_seq >> 8) & 0x3f) | 0x80);	/* variant */
	out[9] = (unsigned char) clock_seq;
	memcpy(out + 10, node, 6);
}

void
uuid_unparse(const uuid_t uu, char *out)
{
	snprintf(out, 37,
			 "%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
			 uu[0], uu[1], uu[2], uu[3], uu[4], uu[5], uu[6], uu[7],
			 uu[8], uu[9], uu[10], uu[11], uu[12], uu[13], uu[14], uu[15]);
}
