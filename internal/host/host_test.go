package host

import (
	"fmt"
	"testing"

	"github.com/shibukawa/pgmem/internal/vfs"
)

type sliceMem []byte

func (s sliceMem) Read(off, n uint32) ([]byte, bool) {
	if uint64(off)+uint64(n) > uint64(len(s)) {
		return nil, false
	}
	return s[off : off+n], true
}
func (s sliceMem) Write(off uint32, b []byte) bool { return copy(s[off:], b) == len(b) }
func (s sliceMem) Size() uint32                    { return uint32(len(s)) }
func (s sliceMem) Grow(uint32) (uint32, bool)      { return 0, false }

// PostgreSQL computes CRC-32C as INIT (0xFFFFFFFF), raw COMP updates, FIN
// (xor 0xFFFFFFFF). The check value of CRC-32C over "123456789" is
// 0xE3069283, so the host's raw update must reproduce it when driven the
// PostgreSQL way, including when the data is fed in pieces.
func TestCRC32CMatchesPostgresConvention(t *testing.T) {
	fn, ok := Lookup("env", "pgmem_crc32c")
	if !ok {
		t.Fatal("pgmem_crc32c not in table")
	}
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 64))
	copy(mem, "123456789")
	crc := uint32(0xFFFFFFFF)
	crc = uint32(fn.Call(h, mem, []uint64{uint64(crc), 0, 4}))
	crc = uint32(fn.Call(h, mem, []uint64{uint64(crc), 4, 5}))
	if got := crc ^ 0xFFFFFFFF; got != 0xE3069283 {
		t.Fatalf("crc32c = %#x, want 0xE3069283", got)
	}
}

func TestRandomBytesFillsGuestMemory(t *testing.T) {
	fn, _ := Lookup("env", "pgmem_random_bytes")
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 32))
	if r := fn.Call(h, mem, []uint64{8, 16}); r != 1 {
		t.Fatalf("returned %d", r)
	}
	zero := true
	for _, b := range mem[8:24] {
		if b != 0 {
			zero = false
		}
	}
	if zero {
		t.Fatal("no random bytes written")
	}
}

func TestHashHandles(t *testing.T) {
	h := New(vfs.New())
	mem := sliceMem(make([]byte, 128))
	copy(mem, "abc")
	create, _ := Lookup("env", "pgmem_hash_create")
	update, _ := Lookup("env", "pgmem_hash_update")
	final, _ := Lookup("env", "pgmem_hash_final")
	free, _ := Lookup("env", "pgmem_hash_free")
	want := map[int32]string{
		0: "900150983cd24fb0d6963f7d28e17f72",                                 // md5
		1: "a9993e364706816aba3e25717850c26c9cd0d89d",                         // sha1
		3: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", // sha256
	}
	for typ, hex := range want {
		hd := int32(create.Call(h, mem, []uint64{uint64(uint32(typ))}))
		if hd <= 0 {
			t.Fatalf("type %d: no handle", typ)
		}
		update.Call(h, mem, []uint64{uint64(uint32(hd)), 0, 3})
		if n := int32(final.Call(h, mem, []uint64{uint64(uint32(hd)), 64, 64})); n != int32(len(hex)/2) {
			t.Fatalf("type %d: final returned %d", typ, n)
		}
		if got := fmt.Sprintf("%x", mem[64:64+len(hex)/2]); got != hex {
			t.Fatalf("type %d: %s != %s", typ, got, hex)
		}
		// Too small a destination is an error, as in PostgreSQL.
		if n := int32(final.Call(h, mem, []uint64{uint64(uint32(hd)), 64, 4})); n != -1 {
			t.Fatalf("type %d: short dest accepted", typ)
		}
		free.Call(h, mem, []uint64{uint64(uint32(hd))})
	}
	if create.Call(h, mem, []uint64{99}) != 0 {
		t.Fatal("unknown type accepted")
	}
}
