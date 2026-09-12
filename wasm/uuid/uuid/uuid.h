/*
 * uuid/uuid.h - the slice of libuuid's interface contrib/uuid-ossp uses
 * (HAVE_UUID_E2FS), implemented by wasm/pgmem_uuid.c on the host's
 * entropy. Built into the uuid-ossp module by wasm/build.sh.
 */
#ifndef PGMEM_UUID_H
#define PGMEM_UUID_H

typedef unsigned char uuid_t[16];

extern void uuid_generate_random(uuid_t out);
extern void uuid_generate_time(uuid_t out);
extern void uuid_unparse(const uuid_t uu, char *out);

#endif
