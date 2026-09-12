---
id: decision:pgcrypto-on-host
type: decision
title: pgcrypto Crypto Runs on the Host
---
pgcrypto ships without OpenSSL or zlib: its OpenSSL-backed and zlib-backed files are replaced by host calls into Go's crypto and compress packages, chosen over porting those libraries to wasm.

```yaml
summary:
  chosen: replace contrib/pgcrypto/openssl.c, pgp-mpi-openssl.c and pgp-compress.c (wasm/patches.py) with wasm/pgmem_pgcrypto_openssl.inc, pgmem_pgcrypto_mpi.inc and pgmem_pgcrypto_compress.inc; host side in internal/host/crypto.go
  rejected: compiling OpenSSL and zlib to wasm (size, build complexity, no TLS need)
  host_imports:
    - pgmem_hash_info plus the existing pgmem_hash_* handles for digests (md5, sha1, sha2, ripemd160, sha3)
    - pgmem_cipher_create/run/free for AES, Blowfish, DES, 3DES, CAST5 in ECB/CBC/CFB with EVP padding semantics
    - pgmem_bn_op/pgmem_bn_rand for OpenPGP RSA and ElGamal (math/big)
    - pgmem_deflate_create/write/finish, pgmem_inflate_all, pgmem_zstream_read/free for OpenPGP ZIP (raw deflate) and ZLIB packets (compress/flate, compress/zlib); decompression is one-shot over the whole packet
  fidelity:
    - CAST5 is vendored (internal/host/cast5) to add OpenSSL's zero-padded short keys and 12-round variant
    - CFB is a 30-line implementation in crypto.go because crypto/cipher deprecates NewCFBEncrypter (Go 1.24)
    - compressed bytes differ from zlib's (Go's deflate); upstream tests only check decryption and round trips
    - AES keys are rounded up to 16/24/32 bytes with zero fill, as ossl_aes_init did
    - verified by replaying upstream contrib/pgcrypto regression files (testdata/regress, TestPgcryptoRegress) with a psql-aligned-format emulator
  limits:
    - bzip2 packets are unsupported, as in upstream pgcrypto
    - fips_mode() is always false; pgcrypto.builtin_crypto_enabled=fips behaves like on
  adding_more_contrib: append the name to CONTRIB_MODULES in wasm/build.sh; control and SQL files are copied into the share tree
```
