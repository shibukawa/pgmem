package host

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"io"
	"math/big"

	"github.com/shibukawa/pgmem/internal/host/cast5"
	"golang.org/x/crypto/blowfish"
)

// pgcrypto's OpenSSL layer (contrib/pgcrypto/openssl.c and
// pgp-mpi-openssl.c) is replaced in the wasm build by host calls declared
// in wasm/pgmem_pgcrypto_openssl.inc and wasm/pgmem_pgcrypto_mpi.inc. The
// digests reuse the cryptohash handles in host.go; this file holds the
// symmetric ciphers and the OpenPGP big-integer arithmetic.

// Cipher algorithm and mode codes, shared with the .inc file.
const (
	cipherBF    = 1
	cipherDES   = 2
	cipherDES3  = 3
	cipherCAST5 = 4
	cipherAES   = 5

	modeECB = 1
	modeCBC = 2
	modeCFB = 3
)

// Big-integer operation codes, shared with the .inc file.
const (
	bnModExp = 1
	bnModMul = 2
	bnModInv = 3
)

// hostCipher is one cipher context. Like an OpenSSL EVP context it is
// bound to one direction by its first use and keeps the CBC/CFB chaining
// state between calls.
type hostCipher struct {
	block   cipher.Block
	mode    int32
	iv      []byte
	started bool
	encrypt bool
	cbc     cipher.BlockMode
	cfb     cipher.Stream
}

func newBlockCipher(algo int32, key []byte) (cipher.Block, error) {
	switch algo {
	case cipherBF:
		return blowfish.NewCipher(key)
	case cipherDES:
		return des.NewCipher(key)
	case cipherDES3:
		return des.NewTripleDESCipher(key)
	case cipherCAST5:
		return cast5.NewCipher(key)
	case cipherAES:
		return aes.NewCipher(key)
	}
	return nil, errors.New("unknown cipher algorithm")
}

func newHostCipher(algo, mode int32, key, iv []byte) (*hostCipher, error) {
	block, err := newBlockCipher(algo, key)
	if err != nil {
		return nil, err
	}
	if mode != modeECB && mode != modeCBC && mode != modeCFB {
		return nil, errors.New("unknown cipher mode")
	}
	if len(iv) < block.BlockSize() {
		return nil, errors.New("iv too short")
	}
	return &hostCipher{block: block, mode: mode, iv: append([]byte(nil), iv[:block.BlockSize()]...)}, nil
}

// run applies the cipher to data with EVP_{En,De}cryptUpdate +
// EVP_{En,De}cryptFinal_ex semantics: PKCS#7 padding when padding is set
// (ignored for CFB), a failure when the length is not a whole number of
// blocks otherwise, and a failure on bad padding when decrypting.
func (c *hostCipher) run(encrypt, padding bool, data []byte) ([]byte, bool) {
	if !c.started {
		c.started = true
		c.encrypt = encrypt
		switch c.mode {
		case modeCBC:
			if encrypt {
				c.cbc = cipher.NewCBCEncrypter(c.block, c.iv)
			} else {
				c.cbc = cipher.NewCBCDecrypter(c.block, c.iv)
			}
		case modeCFB:
			c.cfb = newCFB(c.block, c.iv, encrypt)
		}
	} else if c.encrypt != encrypt {
		return nil, false
	}
	if c.mode == modeCFB {
		out := make([]byte, len(data))
		c.cfb.XORKeyStream(out, data)
		return out, true
	}
	bs := c.block.BlockSize()
	if encrypt {
		in := data
		if padding {
			pad := bs - len(data)%bs
			in = make([]byte, len(data)+pad)
			copy(in, data)
			for i := len(data); i < len(in); i++ {
				in[i] = byte(pad)
			}
		} else if len(data)%bs != 0 {
			return nil, false
		}
		out := make([]byte, len(in))
		c.cryptBlocks(out, in)
		return out, true
	}
	if len(data)%bs != 0 || (padding && len(data) == 0) {
		return nil, false
	}
	out := make([]byte, len(data))
	c.cryptBlocks(out, data)
	if padding {
		pad := int(out[len(out)-1])
		if pad < 1 || pad > bs {
			return nil, false
		}
		for _, b := range out[len(out)-pad:] {
			if int(b) != pad {
				return nil, false
			}
		}
		out = out[:len(out)-pad]
	}
	return out, true
}

func (c *hostCipher) cryptBlocks(dst, src []byte) {
	if c.mode == modeCBC {
		c.cbc.CryptBlocks(dst, src)
		return
	}
	bs := c.block.BlockSize()
	for i := 0; i+bs <= len(src); i += bs {
		if c.encrypt {
			c.block.Encrypt(dst[i:i+bs], src[i:i+bs])
		} else {
			c.block.Decrypt(dst[i:i+bs], src[i:i+bs])
		}
	}
}

// cfbStream is full-block CFB (CFB-64 for 8-byte ciphers, CFB-128 for
// AES, what EVP_*_cfb means): the keystream for each block is the block
// cipher applied to the previous ciphertext block, the IV at first.
// crypto/cipher has the same mode but deprecates it, so it lives here.
type cfbStream struct {
	block   cipher.Block
	reg     []byte // shift register: previous ciphertext block
	ks      []byte // keystream for the current block
	used    int    // bytes of ks consumed
	encrypt bool
}

func newCFB(block cipher.Block, iv []byte, encrypt bool) *cfbStream {
	bs := block.BlockSize()
	return &cfbStream{
		block:   block,
		reg:     append([]byte(nil), iv[:bs]...),
		ks:      make([]byte, bs),
		used:    bs,
		encrypt: encrypt,
	}
}

func (c *cfbStream) XORKeyStream(dst, src []byte) {
	bs := len(c.ks)
	for i := range src {
		if c.used == bs {
			c.block.Encrypt(c.ks, c.reg)
			c.used = 0
		}
		in := src[i]
		out := in ^ c.ks[c.used]
		if c.encrypt {
			c.reg[c.used] = out
		} else {
			c.reg[c.used] = in
		}
		dst[i] = out
		c.used++
	}
}

// zstream is a compression or decompression context: pending output is
// read out by the guest in pieces.
type zstream struct {
	w   io.WriteCloser // deflate side; nil after finish or for inflate
	out bytes.Buffer
}

// newDeflate starts a deflate stream in raw (OpenPGP "ZIP") or zlib
// framing at the given level (0-9, or -1 for the default).
func newDeflate(raw bool, level int) (*zstream, error) {
	z := &zstream{}
	var err error
	if raw {
		z.w, err = flate.NewWriter(&z.out, level)
	} else {
		z.w, err = zlib.NewWriterLevel(&z.out, level)
	}
	if err != nil {
		return nil, err
	}
	return z, nil
}

// inflateAll decompresses one complete stream. Bytes after the end of the
// stream are an error, as pgcrypto treats a compressed packet with
// trailing data as corrupt.
func inflateAll(raw bool, data []byte) (*zstream, error) {
	src := bytes.NewReader(data)
	var r io.ReadCloser
	if raw {
		r = flate.NewReader(src)
	} else {
		var err error
		if r, err = zlib.NewReader(src); err != nil {
			return nil, err
		}
	}
	z := &zstream{}
	if _, err := io.Copy(&z.out, r); err != nil {
		return nil, err
	}
	if src.Len() > 0 {
		return nil, errors.New("extra bytes after end of stream")
	}
	return z, nil
}

// bnOp computes a^b mod m, a*b mod m or a^-1 mod m over big-endian
// unsigned integers; the result has no leading zero bytes (empty for 0).
func bnOp(op int32, a, b, m []byte) ([]byte, bool) {
	A := new(big.Int).SetBytes(a)
	M := new(big.Int).SetBytes(m)
	if M.Sign() == 0 {
		return nil, false
	}
	var r *big.Int
	switch op {
	case bnModExp:
		r = new(big.Int).Exp(A, new(big.Int).SetBytes(b), M)
	case bnModMul:
		r = new(big.Int).Mul(A, new(big.Int).SetBytes(b))
		r.Mod(r, M)
	case bnModInv:
		r = new(big.Int).ModInverse(A, M)
		if r == nil {
			return nil, false
		}
	default:
		return nil, false
	}
	return r.Bytes(), true
}

// bnRand returns a random integer of exactly bits bits (top bit set), the
// way BN_rand(bn, bits, BN_RAND_TOP_ONE, BN_RAND_BOTTOM_ANY) does.
func bnRand(bits int32) ([]byte, bool) {
	if bits <= 0 {
		return nil, false
	}
	n := (int(bits) + 7) / 8
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, false
	}
	b[0] &= byte(0xff >> uint(n*8-int(bits)))
	b[0] |= 1 << uint((int(bits)-1)%8)
	return b, true
}

func init() {
	table = append(table, cryptoTable...)
}

var cryptoTable = []Fn{
	// pgmem_hash_info(handle): block_size << 16 | result_size of a live
	// cryptohash handle, -1 for an unknown one.
	{"env", "pgmem_hash_info", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		hh := h.hashes[i32(a[0])]
		if hh == nil {
			return ret32(-1)
		}
		return ret32(int32(hh.BlockSize())<<16 | int32(hh.Size()))
	}},
	// pgmem_cipher_create(algo, mode, key, klen, iv, ivlen): a handle > 0,
	// or 0 when the key length does not suit the algorithm.
	{"env", "pgmem_cipher_create", "iiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		key, ok := m.Read(u32(a[2]), u32(a[3]))
		if !ok {
			return 0
		}
		iv, ok := m.Read(u32(a[4]), u32(a[5]))
		if !ok {
			return 0
		}
		c, err := newHostCipher(i32(a[0]), i32(a[1]), key, iv)
		if err != nil {
			return 0
		}
		if h.ciphers == nil {
			h.ciphers = map[int32]*hostCipher{}
		}
		h.nextCipher++
		if h.nextCipher <= 0 {
			h.nextCipher = 1
		}
		h.ciphers[h.nextCipher] = c
		return ret32(h.nextCipher)
	}},
	// pgmem_cipher_run(handle, encrypt, padding, data, dlen, res, rescap):
	// output length, or -1 on failure (bad length, bad padding, wrong
	// direction for this handle, output larger than rescap).
	{"env", "pgmem_cipher_run", "iiiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		c := h.ciphers[i32(a[0])]
		if c == nil {
			return ret32(-1)
		}
		data, ok := m.Read(u32(a[3]), u32(a[4]))
		if !ok {
			return ret32(-1)
		}
		out, ok := c.run(i32(a[1]) != 0, i32(a[2]) != 0, data)
		if !ok || uint32(len(out)) > u32(a[6]) {
			return ret32(-1)
		}
		if !m.Write(u32(a[5]), out) {
			return ret32(-1)
		}
		return ret32(int32(len(out)))
	}},
	{"env", "pgmem_cipher_free", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		delete(h.ciphers, i32(a[0]))
		return 0
	}},
	// pgmem_bn_op(op, a, alen, b, blen, m, mlen, out, cap): result length
	// (0 for zero), or -1 when the operation fails or cap is too small.
	{"env", "pgmem_bn_op", "iiiiiiiii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		x, ok1 := m.Read(u32(a[1]), u32(a[2]))
		y, ok2 := m.Read(u32(a[3]), u32(a[4]))
		mod, ok3 := m.Read(u32(a[5]), u32(a[6]))
		if !ok1 || !ok2 || !ok3 {
			return ret32(-1)
		}
		out, ok := bnOp(i32(a[0]), x, y, mod)
		if !ok || uint32(len(out)) > u32(a[8]) {
			return ret32(-1)
		}
		if !m.Write(u32(a[7]), out) {
			return ret32(-1)
		}
		return ret32(int32(len(out)))
	}},
	// ===== pgcrypto's zlib (wasm/pgmem_pgcrypto_compress.inc) =====
	// pgmem_deflate_create(raw, level): a handle > 0, or 0 on a bad level.
	{"env", "pgmem_deflate_create", "ii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		z, err := newDeflate(i32(a[0]) != 0, int(i32(a[1])))
		if err != nil {
			return 0
		}
		return ret32(h.addZstream(z))
	}},
	// pgmem_deflate_write(handle, data, len): 0, or -1 on failure.
	{"env", "pgmem_deflate_write", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		z := h.zstreams[i32(a[0])]
		data, ok := m.Read(u32(a[1]), u32(a[2]))
		if z == nil || z.w == nil || !ok {
			return ret32(-1)
		}
		if _, err := z.w.Write(data); err != nil {
			return ret32(-1)
		}
		return 0
	}},
	// pgmem_deflate_finish(handle): ends the stream; 0, or -1 on failure.
	{"env", "pgmem_deflate_finish", "i", "i", func(h *Host, m Memory, a []uint64) uint64 {
		z := h.zstreams[i32(a[0])]
		if z == nil || z.w == nil {
			return ret32(-1)
		}
		err := z.w.Close()
		z.w = nil
		if err != nil {
			return ret32(-1)
		}
		return 0
	}},
	// pgmem_inflate_all(raw, data, len): a handle > 0 whose output is the
	// whole decompressed stream, or -1 when the data is not a valid stream
	// or has trailing bytes.
	{"env", "pgmem_inflate_all", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		data, ok := m.Read(u32(a[1]), u32(a[2]))
		if !ok {
			return ret32(-1)
		}
		z, err := inflateAll(i32(a[0]) != 0, data)
		if err != nil {
			return ret32(-1)
		}
		return ret32(h.addZstream(z))
	}},
	// pgmem_zstream_read(handle, out, cap): moves up to cap bytes of pending
	// output to the guest and returns the count (0 when drained), or -1.
	{"env", "pgmem_zstream_read", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		z := h.zstreams[i32(a[0])]
		if z == nil {
			return ret32(-1)
		}
		n := int(u32(a[2]))
		if n > z.out.Len() {
			n = z.out.Len()
		}
		if n == 0 {
			return 0
		}
		if !m.Write(u32(a[1]), z.out.Next(n)) {
			return ret32(-1)
		}
		return ret32(int32(n))
	}},
	{"env", "pgmem_zstream_free", "i", "", func(h *Host, m Memory, a []uint64) uint64 {
		delete(h.zstreams, i32(a[0]))
		return 0
	}},
	// pgmem_bn_rand(bits, out, cap): byte length of the number, or -1.
	{"env", "pgmem_bn_rand", "iii", "i", func(h *Host, m Memory, a []uint64) uint64 {
		out, ok := bnRand(i32(a[0]))
		if !ok || uint32(len(out)) > u32(a[2]) {
			return ret32(-1)
		}
		if !m.Write(u32(a[1]), out) {
			return ret32(-1)
		}
		return ret32(int32(len(out)))
	}},
}

func (h *Host) addZstream(z *zstream) int32 {
	if h.zstreams == nil {
		h.zstreams = map[int32]*zstream{}
	}
	h.nextZstream++
	if h.nextZstream <= 0 {
		h.nextZstream = 1
	}
	h.zstreams[h.nextZstream] = z
	return h.nextZstream
}
