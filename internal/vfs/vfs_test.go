package vfs

import (
	"bytes"
	"testing"
)

// A clone and its source share file contents until one of them writes;
// every write path must give the writer a private copy first.
func TestCloneCopyOnWrite(t *testing.T) {
	src := New()
	if e := src.WriteFile("/data/a", []byte("hello"), 0o644); e != OK {
		t.Fatal(e)
	}
	if e := src.WriteFile("/data/b", bytes.Repeat([]byte{7}, 100), 0o644); e != OK {
		t.Fatal(e)
	}
	c := src.Clone()

	// shared until written
	sa, _ := src.Lookup("/data/a")
	ca, _ := c.Lookup("/data/a")
	if &sa.data[0] != &ca.data[0] || !sa.shared || !ca.shared {
		t.Fatal("clone should share the bytes")
	}

	// in-place write in the clone
	fd, e := c.Openat(AT_FDCWD, "/data/a", O_WRONLY, 0)
	if e != OK {
		t.Fatal(e)
	}
	if _, e := c.Write(fd, []byte("J")); e != OK {
		t.Fatal(e)
	}
	c.Close(fd)
	check := func(fs *FS, p, want string) {
		t.Helper()
		got, e := fs.ReadFile(p)
		if e != OK || string(got) != want {
			t.Fatalf("%s: %q (%v), want %q", p, got, e, want)
		}
	}
	check(c, "/data/a", "Jello")
	check(src, "/data/a", "hello")
	if ca.shared || !sa.shared {
		t.Fatalf("after the clone wrote: clone shared=%v source shared=%v", ca.shared, sa.shared)
	}

	// append within capacity in the source must not leak into the clone
	fd, _ = src.Openat(AT_FDCWD, "/data/b", O_WRONLY|O_APPEND, 0)
	src.Write(fd, []byte{9})
	src.Close(fd)
	sb, _ := src.Lookup("/data/b")
	cb, _ := c.Lookup("/data/b")
	if len(sb.data) != 101 || len(cb.data) != 100 || cb.data[99] != 7 {
		t.Fatalf("append leaked: src %d clone %d", len(sb.data), len(cb.data))
	}
	if sb.shared {
		t.Fatal("source should own its data after writing")
	}

	// O_TRUNC in the clone drops the shared bytes without copying them
	fd, _ = c.Openat(AT_FDCWD, "/data/b", O_WRONLY|O_TRUNC, 0)
	c.Write(fd, []byte("x"))
	c.Close(fd)
	check(c, "/data/b", "x")
	if got, _ := src.ReadFile("/data/b"); len(got) != 101 {
		t.Fatalf("source truncated through the clone: %d bytes", len(got))
	}

	// truncate: shrink keeps sharing, the later write copies the prefix only
	c2 := src.Clone()
	if e := c2.Truncate("/data/b", 3); e != OK {
		t.Fatal(e)
	}
	fd, _ = c2.Openat(AT_FDCWD, "/data/b", O_WRONLY, 0)
	c2.Write(fd, []byte{1})
	c2.Close(fd)
	if got, _ := c2.ReadFile("/data/b"); !bytes.Equal(got, []byte{1, 7, 7}) {
		t.Fatalf("c2: %v", got)
	}
	if got, _ := src.ReadFile("/data/b"); len(got) != 101 || got[0] != 7 {
		t.Fatalf("source changed by clone truncate: %v", got[:4])
	}

	// growing past capacity in the clone
	c3 := src.Clone()
	fd, _ = c3.Openat(AT_FDCWD, "/data/b", O_WRONLY, 0)
	c3.Seek(fd, 200, SEEK_SET)
	c3.Write(fd, []byte{5})
	c3.Close(fd)
	if got, _ := c3.ReadFile("/data/b"); len(got) != 201 || got[200] != 5 || got[150] != 0 {
		t.Fatalf("c3 grow: %d bytes", len(got))
	}
	if got, _ := src.ReadFile("/data/b"); len(got) != 101 {
		t.Fatalf("source grew: %d", len(got))
	}

	// a shared file whose array has spare capacity (grown by earlier
	// writes) is extended into a private array, not in place
	big := New()
	fd, _ = big.Openat(AT_FDCWD, "/w", O_WRONLY|O_CREAT, 0o644)
	for i := 0; i < 6; i++ {
		big.Write(fd, bytes.Repeat([]byte{byte(i)}, 8192)) // 48 KB, capacity well above
	}
	big.Close(fd)
	bn, _ := big.Lookup("/w")
	if cap(bn.data) <= len(bn.data) {
		t.Fatalf("test needs spare capacity: len %d cap %d", len(bn.data), cap(bn.data))
	}
	bc := big.Clone()
	fd, _ = bc.Openat(AT_FDCWD, "/w", O_WRONLY|O_APPEND, 0)
	if _, e := bc.Write(fd, bytes.Repeat([]byte{9}, 8192)); e != OK {
		t.Fatal(e)
	}
	bc.Close(fd)
	fd, _ = big.Openat(AT_FDCWD, "/w", O_WRONLY|O_APPEND, 0)
	big.Write(fd, bytes.Repeat([]byte{8}, 8192))
	big.Close(fd)
	gc, _ := bc.ReadFile("/w")
	gs, _ := big.ReadFile("/w")
	if len(gc) != 7*8192 || gc[6*8192] != 9 || len(gs) != 7*8192 || gs[6*8192] != 8 {
		t.Fatalf("shared growth leaked: clone %d/%d source %d/%d", len(gc), gc[6*8192], len(gs), gs[6*8192])
	}

	// WriteFile and PutFile replace shared data with private data
	c4 := src.Clone()
	c4.WriteFile("/data/a", []byte("new"), 0o644)
	c4.PutFile("/data/b", []byte("put"), 0o644)
	check(c4, "/data/a", "new")
	check(c4, "/data/b", "put")
	check(src, "/data/a", "hello")

	// cloning a clone that was never written shares the same bytes, and a
	// second clone of the source does not write to the source's nodes
	c5 := c4.Clone()
	n4, _ := c4.Lookup("/data/a")
	n5, _ := c5.Lookup("/data/a")
	if &n4.data[0] != &n5.data[0] {
		t.Fatal("clone of a clone should share")
	}
}
