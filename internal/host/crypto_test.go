package host

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/shibukawa/pgmem/internal/vfs"
)

func mustHex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// cipherOnce drives pgmem_cipher_create/run/free the way the .inc does and
// returns the output, or nil when the run fails.
func cipherOnce(t *testing.T, algo, mode int32, key, iv, data []byte, encrypt, padding bool) []byte {
	t.Helper()
	create, _ := Lookup("env", "pgmem_cipher_create")
	run, _ := Lookup("env", "pgmem_cipher_run")
	free, _ := Lookup("env", "pgmem_cipher_free")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 4096))
	copy(mem[0:], key)
	copy(mem[128:], iv)
	copy(mem[256:], data)
	if len(iv) == 0 {
		iv = make([]byte, 16)
	}
	handle := int32(create.Call(h, mem, []uint64{uint64(algo), uint64(mode), 0, uint64(len(key)), 128, uint64(len(iv))}))
	if handle <= 0 {
		return nil
	}
	e, p := uint64(0), uint64(0)
	if encrypt {
		e = 1
	}
	if padding {
		p = 1
	}
	n := int32(run.Call(h, mem, []uint64{uint64(handle), e, p, 256, uint64(len(data)), 1024, uint64(len(data) + 16)}))
	free.Call(h, mem, []uint64{uint64(handle)})
	if n < 0 {
		return nil
	}
	return append([]byte(nil), mem[1024:1024+int(n)]...)
}

// Published single-block vectors: FIPS-197 C.1 (AES), the RFC 2144 B.1
// set including the 80- and 40-bit keys that select 12 rounds (CAST5),
// Schneier's Blowfish vector and the classic DES vector.
func TestCipherKnownAnswers(t *testing.T) {
	cases := []struct {
		name        string
		algo        int32
		key, pt, ct string
	}{
		{"aes-128", cipherAES, "000102030405060708090a0b0c0d0e0f", "00112233445566778899aabbccddeeff", "69c4e0d86a7b0430d8cdb78070b4c55a"},
		{"aes-256", cipherAES, "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", "00112233445566778899aabbccddeeff", "8ea2b7ca516745bfeafc49904b496089"},
		{"cast5-128", cipherCAST5, "0123456712345678234567893456789a", "0123456789abcdef", "238b4fe5847e44b2"},
		{"cast5-80", cipherCAST5, "01234567123456782345", "0123456789abcdef", "eb6a711a2c02271b"},
		{"cast5-40", cipherCAST5, "0123456712", "0123456789abcdef", "7ac816d16e9b302e"},
		{"blowfish", cipherBF, "0000000000000000", "0000000000000000", "4ef997456198dd78"},
		{"des", cipherDES, "133457799bbcdff1", "0123456789abcdef", "85e813540f0ab405"},
		{"des3", cipherDES3, "0123456789abcdef23456789abcdef01456789abcdef0123", "6bc1bee22e409f96", "714772f339841d34"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			key, pt, want := mustHex(t, c.key), mustHex(t, c.pt), mustHex(t, c.ct)
			got := cipherOnce(t, c.algo, modeECB, key, nil, pt, true, false)
			if !bytes.Equal(got, want) {
				t.Fatalf("encrypt = %x, want %x", got, want)
			}
			back := cipherOnce(t, c.algo, modeECB, key, nil, want, false, false)
			if !bytes.Equal(back, pt) {
				t.Fatalf("decrypt = %x, want %x", back, pt)
			}
		})
	}
}

func TestCipherPaddingSemantics(t *testing.T) {
	key := mustHex(t, "000102030405060708090a0b0c0d0e0f")
	iv := mustHex(t, "0f0e0d0c0b0a09080706050403020100")
	data := []byte("hello")
	// padding on: one block out, and it decrypts back with the pad removed
	ct := cipherOnce(t, cipherAES, modeCBC, key, iv, data, true, true)
	if len(ct) != 16 {
		t.Fatalf("padded ciphertext is %d bytes", len(ct))
	}
	if pt := cipherOnce(t, cipherAES, modeCBC, key, iv, ct, false, true); !bytes.Equal(pt, data) {
		t.Fatalf("decrypt = %q", pt)
	}
	// padding on, empty input: a whole block of padding
	if ct := cipherOnce(t, cipherAES, modeCBC, key, iv, nil, true, true); len(ct) != 16 {
		t.Fatalf("empty input gave %d bytes", len(ct))
	}
	// padding off: partial block is an error, whole blocks pass through
	if ct := cipherOnce(t, cipherAES, modeCBC, key, iv, data, true, false); ct != nil {
		t.Fatalf("unpadded partial block accepted: %x", ct)
	}
	if ct := cipherOnce(t, cipherAES, modeCBC, key, iv, make([]byte, 32), true, false); len(ct) != 32 {
		t.Fatalf("unpadded whole blocks gave %d bytes", len(ct))
	}
	// decrypt: partial block, empty input with padding, and bad padding fail
	if pt := cipherOnce(t, cipherAES, modeCBC, key, iv, ct[:5], false, false); pt != nil {
		t.Fatal("partial block decrypt accepted")
	}
	if pt := cipherOnce(t, cipherAES, modeCBC, key, iv, nil, false, true); pt != nil {
		t.Fatal("empty padded decrypt accepted")
	}
	garbage := cipherOnce(t, cipherAES, modeCBC, key, iv, make([]byte, 16), true, false)
	if pt := cipherOnce(t, cipherAES, modeCBC, key, iv, garbage, false, true); pt != nil && len(pt) == 16 {
		t.Fatal("bad padding not detected")
	}
	// CFB is a stream: any length, padding ignored
	if ct := cipherOnce(t, cipherAES, modeCFB, key, iv, data, true, true); len(ct) != len(data) {
		t.Fatalf("cfb gave %d bytes", len(ct))
	}
	if ct := cipherOnce(t, cipherBF, modeCFB, key[:8], iv[:8], data, true, false); len(ct) != len(data) {
		t.Fatalf("bf-cfb gave %d bytes", len(ct))
	}
}

func TestCipherRejectsBadKeys(t *testing.T) {
	if cipherOnce(t, cipherBF, modeECB, nil, nil, make([]byte, 8), true, false) != nil {
		t.Fatal("empty blowfish key accepted")
	}
	if cipherOnce(t, cipherCAST5, modeECB, make([]byte, 17), nil, make([]byte, 8), true, false) != nil {
		t.Fatal("17-byte cast5 key accepted")
	}
	if cipherOnce(t, cipherAES, modeECB, make([]byte, 20), nil, make([]byte, 16), true, false) != nil {
		t.Fatal("20-byte aes key accepted")
	}
}

func TestCipherKeepsChainingState(t *testing.T) {
	// two runs on one handle must equal one run over the concatenation
	create, _ := Lookup("env", "pgmem_cipher_create")
	run, _ := Lookup("env", "pgmem_cipher_run")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 4096))
	key := mustHex(t, "000102030405060708090a0b0c0d0e0f")
	copy(mem[0:], key)
	copy(mem[256:], bytes.Repeat([]byte{7}, 32))
	handle := create.Call(h, mem, []uint64{cipherAES, modeCBC, 0, 16, 128, 16})
	n1 := run.Call(h, mem, []uint64{handle, 1, 0, 256, 16, 1024, 32})
	n2 := run.Call(h, mem, []uint64{handle, 1, 0, 272, 16, 1040, 32})
	if int32(n1) != 16 || int32(n2) != 16 {
		t.Fatalf("runs returned %d, %d", int32(n1), int32(n2))
	}
	whole := cipherOnce(t, cipherAES, modeCBC, key, make([]byte, 16), bytes.Repeat([]byte{7}, 32), true, false)
	if !bytes.Equal(mem[1024:1056], whole) {
		t.Fatal("chained runs differ from a single run")
	}
	// a handle bound to encryption refuses to decrypt
	if r := int32(run.Call(h, mem, []uint64{handle, 0, 0, 256, 16, 1024, 32})); r != -1 {
		t.Fatalf("direction change returned %d", r)
	}
}

func TestBignumOps(t *testing.T) {
	op, _ := Lookup("env", "pgmem_bn_op")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 4096))
	call := func(code int32, a, b, m []byte) ([]byte, int32) {
		copy(mem[0:], a)
		copy(mem[256:], b)
		copy(mem[512:], m)
		n := int32(op.Call(h, mem, []uint64{uint64(code), 0, uint64(len(a)), 256, uint64(len(b)), 512, uint64(len(m)), 1024, 256}))
		if n < 0 {
			return nil, n
		}
		return append([]byte(nil), mem[1024:1024+int(n)]...), n
	}
	// 4^13 mod 497 = 445
	if r, _ := call(bnModExp, []byte{4}, []byte{13}, []byte{1, 241}); !bytes.Equal(r, []byte{1, 189}) {
		t.Fatalf("modexp = %x", r)
	}
	// 3 * 4 mod 5 = 2
	if r, _ := call(bnModMul, []byte{3}, []byte{4}, []byte{5}); !bytes.Equal(r, []byte{2}) {
		t.Fatalf("modmul = %x", r)
	}
	// 3^-1 mod 11 = 4; 2 has no inverse mod 4
	if r, _ := call(bnModInv, []byte{3}, nil, []byte{11}); !bytes.Equal(r, []byte{4}) {
		t.Fatalf("modinv = %x", r)
	}
	if _, n := call(bnModInv, []byte{2}, nil, []byte{4}); n != -1 {
		t.Fatalf("modinv without inverse returned %d", n)
	}
	// zero result has zero length; zero modulus fails
	if r, n := call(bnModMul, []byte{5}, []byte{0}, []byte{7}); n != 0 || len(r) != 0 {
		t.Fatalf("zero result: n=%d r=%x", n, r)
	}
	if _, n := call(bnModExp, []byte{5}, []byte{1}, []byte{0}); n != -1 {
		t.Fatalf("zero modulus returned %d", n)
	}
	// an RSA-sized round trip: (m^e)^d mod n == m
	p, _ := new(big.Int).SetString("170141183460469231731687303715884105727", 10) // 2^127-1
	q, _ := new(big.Int).SetString("618970019642690137449562111", 10)             // 2^89-1
	n := new(big.Int).Mul(p, q)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	e := big.NewInt(65537)
	d := new(big.Int).ModInverse(e, phi)
	msg := big.NewInt(123456789)
	c, _ := call(bnModExp, msg.Bytes(), e.Bytes(), n.Bytes())
	back, _ := call(bnModExp, c, d.Bytes(), n.Bytes())
	if !bytes.Equal(back, msg.Bytes()) {
		t.Fatalf("rsa round trip = %x", back)
	}
}

func TestBignumRand(t *testing.T) {
	rnd, _ := Lookup("env", "pgmem_bn_rand")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 4096))
	for _, bits := range []int{1, 7, 8, 9, 100, 262} {
		n := int32(rnd.Call(h, mem, []uint64{uint64(bits), 0, 256}))
		if int(n) != (bits+7)/8 {
			t.Fatalf("bits=%d: length %d", bits, n)
		}
		v := new(big.Int).SetBytes(mem[:n])
		if v.BitLen() != bits {
			t.Fatalf("bits=%d: got %d bits", bits, v.BitLen())
		}
	}
	if r := int32(rnd.Call(h, mem, []uint64{0, 0, 256})); r != -1 {
		t.Fatalf("bits=0 returned %d", r)
	}
	if r := int32(rnd.Call(h, mem, []uint64{4096, 0, 8})); r != -1 {
		t.Fatalf("cap too small returned %d", r)
	}
}

func TestHashInfoAndExtraDigests(t *testing.T) {
	create, _ := Lookup("env", "pgmem_hash_create")
	info, _ := Lookup("env", "pgmem_hash_info")
	update, _ := Lookup("env", "pgmem_hash_update")
	final, _ := Lookup("env", "pgmem_hash_final")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 4096))
	copy(mem[0:], "abc")
	cases := []struct {
		typ         int32
		block, size int32
		abc         string
	}{
		{0, 64, 16, "900150983cd24fb0d6963f7d28e17f72"},
		{3, 64, 32, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{5, 128, 64, "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
		{6, 64, 20, "8eb208f7e05d987a9b044a8e98c6b087f15a0bfc"},
		{8, 136, 32, "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532"},
	}
	for _, c := range cases {
		id := create.Call(h, mem, []uint64{uint64(c.typ)})
		if id == 0 {
			t.Fatalf("type %d: no handle", c.typ)
		}
		got := int32(info.Call(h, mem, []uint64{id}))
		if got>>16 != c.block || got&0xffff != c.size {
			t.Fatalf("type %d: info = %d/%d", c.typ, got>>16, got&0xffff)
		}
		update.Call(h, mem, []uint64{id, 0, 3})
		n := final.Call(h, mem, []uint64{id, 1024, 64})
		if hex.EncodeToString(mem[1024:1024+n]) != c.abc {
			t.Fatalf("type %d: digest(abc) = %x", c.typ, mem[1024:1024+n])
		}
	}
	if r := int32(info.Call(h, mem, []uint64{9999})); r != -1 {
		t.Fatalf("unknown handle info = %d", r)
	}
}

// NIST SP 800-38A F.3.13/F.3.14: CFB128-AES128 encrypt and decrypt, two
// blocks so the shift register is exercised, plus a Blowfish CFB64
// round trip through partial-block calls.
func TestCFBKnownAnswer(t *testing.T) {
	key := mustHex(t, "2b7e151628aed2a6abf7158809cf4f3c")
	iv := mustHex(t, "000102030405060708090a0b0c0d0e0f")
	pt := mustHex(t, "6bc1bee22e409f96e93d7e117393172aae2d8a571e03ac9c9eb76fac45af8e51")
	want := mustHex(t, "3b3fd92eb72dad20333449f8e83cfb4ac8a64537a0b3a93fcde3cdad9f1ce58b")
	if got := cipherOnce(t, cipherAES, modeCFB, key, iv, pt, true, false); !bytes.Equal(got, want) {
		t.Fatalf("encrypt = %x", got)
	}
	if got := cipherOnce(t, cipherAES, modeCFB, key, iv, want, false, false); !bytes.Equal(got, pt) {
		t.Fatalf("decrypt = %x", got)
	}
	// odd-sized pieces on one stream must equal one call
	block, _ := newBlockCipher(cipherBF, []byte("bfkey"))
	msg := []byte("The quick brown fox jumps over the lazy dog")
	whole := make([]byte, len(msg))
	newCFB(block, make([]byte, 8), true).XORKeyStream(whole, msg)
	pieces := make([]byte, len(msg))
	st := newCFB(block, make([]byte, 8), true)
	for i, n := 0, 0; i < len(msg); i += n {
		n = 1 + i%5
		if i+n > len(msg) {
			n = len(msg) - i
		}
		st.XORKeyStream(pieces[i:i+n], msg[i:i+n])
	}
	if !bytes.Equal(whole, pieces) {
		t.Fatal("piecewise CFB differs")
	}
	back := make([]byte, len(msg))
	newCFB(block, make([]byte, 8), false).XORKeyStream(back, whole)
	if !bytes.Equal(back, msg) {
		t.Fatalf("round trip = %q", back)
	}
}

func TestZlibHandles(t *testing.T) {
	create, _ := Lookup("env", "pgmem_deflate_create")
	write, _ := Lookup("env", "pgmem_deflate_write")
	finish, _ := Lookup("env", "pgmem_deflate_finish")
	inflate, _ := Lookup("env", "pgmem_inflate_all")
	read, _ := Lookup("env", "pgmem_zstream_read")
	free, _ := Lookup("env", "pgmem_zstream_free")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 1<<16))
	msg := bytes.Repeat([]byte("compress me, please. "), 200)
	copy(mem[0:], msg)
	drain := func(handle uint64) []byte {
		var out []byte
		for {
			n := int32(read.Call(h, mem, []uint64{handle, 8192, 100}))
			if n < 0 {
				t.Fatal("read failed")
			}
			if n == 0 {
				return out
			}
			out = append(out, mem[8192:8192+int(n)]...)
		}
	}
	for _, raw := range []uint64{0, 1} {
		d := create.Call(h, mem, []uint64{raw, 6})
		if d == 0 {
			t.Fatal("no deflate handle")
		}
		// two writes, output collected before and after finish
		if r := int32(write.Call(h, mem, []uint64{d, 0, 1000})); r != 0 {
			t.Fatal("write 1")
		}
		comp := drain(d)
		if r := int32(write.Call(h, mem, []uint64{d, 1000, uint64(len(msg) - 1000)})); r != 0 {
			t.Fatal("write 2")
		}
		if r := int32(finish.Call(h, mem, []uint64{d})); r != 0 {
			t.Fatal("finish")
		}
		comp = append(comp, drain(d)...)
		free.Call(h, mem, []uint64{d})
		if len(comp) == 0 || len(comp) > len(msg)/4 {
			t.Fatalf("raw=%d: compressed to %d bytes", raw, len(comp))
		}
		copy(mem[20000:], comp)
		i := inflate.Call(h, mem, []uint64{raw, 20000, uint64(len(comp))})
		if int32(i) <= 0 {
			t.Fatalf("raw=%d: inflate failed", raw)
		}
		if got := drain(i); !bytes.Equal(got, msg) {
			t.Fatalf("raw=%d: round trip differs (%d bytes)", raw, len(got))
		}
		free.Call(h, mem, []uint64{i})
		// trailing garbage and corrupt input are rejected
		mem[20000+len(comp)] = 0x42
		if r := int32(inflate.Call(h, mem, []uint64{raw, 20000, uint64(len(comp) + 1)})); r != -1 {
			t.Fatalf("raw=%d: trailing byte accepted (%d)", raw, r)
		}
		if r := int32(inflate.Call(h, mem, []uint64{raw, 0, 50})); r != -1 {
			t.Fatalf("raw=%d: garbage accepted (%d)", raw, r)
		}
	}
	if r := create.Call(h, mem, []uint64{1, 42}); r != 0 {
		t.Fatalf("level 42 accepted: %d", r)
	}
}
