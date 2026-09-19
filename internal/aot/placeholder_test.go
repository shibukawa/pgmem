package aot

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// fakeOps records the calls the bookkeeping makes, in the shape the
// Windows implementation would issue them.
type fakeOps struct{ calls []string }

func (f *fakeOps) log(format string, a ...any) { f.calls = append(f.calls, fmt.Sprintf(format, a...)) }
func (f *fakeOps) split(addr, size uint64) error {
	f.log("split %#x+%#x", addr, size)
	return nil
}
func (f *fakeOps) coalesce(addr, size uint64) error {
	f.log("coalesce %#x+%#x", addr, size)
	return nil
}
func (f *fakeOps) commit(addr, size uint64) error {
	f.log("commit %#x+%#x", addr, size)
	return nil
}
func (f *fakeOps) mapView(addr, size uint64, _ *os.File) error {
	f.log("map %#x+%#x", addr, size)
	return nil
}
func (f *fakeOps) toHolder(addr, size uint64, kind pieceKind) error {
	f.log("holder %#x+%#x kind=%d", addr, size, kind)
	return nil
}
func (f *fakeOps) free(addr uint64, kind pieceKind) error {
	f.log("free %#x kind=%d", addr, kind)
	return nil
}

const k64 = 64 << 10

func layout(r *region) string {
	var sb strings.Builder
	for _, p := range r.pieces {
		fmt.Fprintf(&sb, "[%#x,%#x)%d ", p.start, p.end, p.kind)
	}
	return strings.TrimSpace(sb.String())
}

func TestRegionHeapGrowthAndSegments(t *testing.T) {
	ops := &fakeOps{}
	r := newRegion(16*k64, ops)
	// the heap grows twice from the bottom
	if err := r.Commit(0, 2*k64); err != nil {
		t.Fatal(err)
	}
	if err := r.Commit(2*k64, 4*k64); err != nil {
		t.Fatal(err)
	}
	// a segment in the middle, then detached and a larger one over it
	if err := r.Map(8*k64, 2*k64, nil); err != nil {
		t.Fatal(err)
	}
	if err := r.Unmap(8*k64, 2*k64); err != nil {
		t.Fatal(err)
	}
	if err := r.Map(8*k64, 4*k64, nil); err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("[0x0,%#x)1 [%#x,%#x)1 [%#x,%#x)0 [%#x,%#x)2 [%#x,%#x)0",
		2*k64, 2*k64, 4*k64, 4*k64, 8*k64, 8*k64, 12*k64, 12*k64, 16*k64)
	if got := layout(r); got != want {
		t.Fatalf("layout\n got %s\nwant %s", got, want)
	}
	wantCalls := []string{
		// Commit(0, 2*k64): the single placeholder is split at 2*k64
		fmt.Sprintf("split 0x0+%#x", 2*k64),
		fmt.Sprintf("commit 0x0+%#x", 2*k64),
		// Commit(2*k64, 4*k64): the remaining placeholder starts at 2*k64, split at 4*k64
		fmt.Sprintf("split %#x+%#x", 2*k64, 2*k64),
		fmt.Sprintf("commit %#x+%#x", 2*k64, 2*k64),
		// Map(8*k64, 2*k64): inside [4*k64, 16*k64): one split makes three pieces
		fmt.Sprintf("split %#x+%#x", 8*k64, 2*k64),
		fmt.Sprintf("map %#x+%#x", 8*k64, 2*k64),
		fmt.Sprintf("holder %#x+%#x kind=2", 8*k64, 2*k64),
		// Map(8*k64, 4*k64): [8,10) is a placeholder, [10,16) must be split at 12, then merged
		fmt.Sprintf("split %#x+%#x", 10*k64, 2*k64),
		fmt.Sprintf("coalesce %#x+%#x", 8*k64, 4*k64),
		fmt.Sprintf("map %#x+%#x", 8*k64, 4*k64),
	}
	if strings.Join(ops.calls, "\n") != strings.Join(wantCalls, "\n") {
		t.Fatalf("calls\n got:\n%s\nwant:\n%s", strings.Join(ops.calls, "\n"), strings.Join(wantCalls, "\n"))
	}
	// mapping over committed memory is refused
	if err := r.Map(3*k64, 2*k64, nil); err == nil {
		t.Fatal("mapping over private memory succeeded")
	}
	ops.calls = nil
	if err := r.Release(); err != nil {
		t.Fatal(err)
	}
	if len(ops.calls) != 5 {
		t.Fatalf("release freed %d pieces, want 5: %v", len(ops.calls), ops.calls)
	}
}

func TestRegionExactPlaceholderNeedsNoCalls(t *testing.T) {
	ops := &fakeOps{}
	r := newRegion(4*k64, ops)
	if err := r.Commit(0, 4*k64); err != nil {
		t.Fatal(err)
	}
	if len(ops.calls) != 1 || !strings.HasPrefix(ops.calls[0], "commit") {
		t.Fatalf("whole-region commit issued %v", ops.calls)
	}
}
