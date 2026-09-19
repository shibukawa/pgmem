package aot

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"sync"
)

// Placeholder bookkeeping for the Windows linear memory.
//
// Unix keeps one anonymous mapping per instance and puts shared segments
// into it with mmap(MAP_FIXED). Windows has no MAP_FIXED; its equivalent
// is a placeholder reservation, which can be split into pieces, each piece
// replaced by committed private memory or by a view of a file mapping, and
// each given back to a placeholder. Which piece is what has to be tracked,
// because every call must name a piece exactly (a view replaces exactly one
// placeholder of its size, adjacent placeholders only merge on request).
// That tracking is platform-independent and lives here, so it can be
// tested anywhere; mem_windows.go supplies the system calls.

type pieceKind uint8

const (
	pieceHolder  pieceKind = iota // a placeholder: reserved, inaccessible
	piecePrivate                  // committed private memory
	pieceView                     // a view of a shared segment's file
)

type piece struct {
	start, end uint64 // offsets into the region
	kind       pieceKind
}

// pageOps are the system calls the bookkeeping drives; addresses are
// offsets into the region.
type pageOps interface {
	// split divides the placeholder starting at addr: its first size
	// bytes stay, the rest becomes a placeholder of its own.
	split(addr, size uint64) error
	// coalesce merges the adjacent placeholders that exactly cover
	// [addr, addr+size) into one.
	coalesce(addr, size uint64) error
	// commit replaces the placeholder [addr, addr+size) with private,
	// readable and writable memory.
	commit(addr, size uint64) error
	// mapView replaces the placeholder [addr, addr+size) with a view of f.
	mapView(addr, size uint64, f *os.File) error
	// toHolder gives a private or view piece back to a placeholder.
	toHolder(addr, size uint64, kind pieceKind) error
	// free releases a piece for good.
	free(addr uint64, kind pieceKind) error
}

// region is the bookkeeping of one reservation.
type region struct {
	mu     sync.Mutex
	size   uint64
	pieces []piece // sorted, contiguous, covering [0, size)
	ops    pageOps
}

func newRegion(size uint64, ops pageOps) *region {
	return &region{size: size, pieces: []piece{{0, size, pieceHolder}}, ops: ops}
}

// Commit makes [from, to) private read/write memory.
func (r *region) Commit(from, to uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.makeHolder(from, to); err != nil {
		return err
	}
	if err := r.ops.commit(from, to-from); err != nil {
		return err
	}
	r.setKind(from, to, piecePrivate)
	return nil
}

// Map puts a view of f over [off, off+size).
func (r *region) Map(off, size uint64, f *os.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.makeHolder(off, off+size); err != nil {
		return err
	}
	if err := r.ops.mapView(off, size, f); err != nil {
		return err
	}
	r.setKind(off, off+size, pieceView)
	return nil
}

// Unmap gives the view at [off, off+size) back to a placeholder.
func (r *region) Unmap(off, size uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	i := r.find(off)
	if i < 0 || r.pieces[i].kind != pieceView || r.pieces[i].start != off || r.pieces[i].end != off+size {
		return fmt.Errorf("aot: no view at [%#x, %#x)", off, off+size)
	}
	if err := r.ops.toHolder(off, size, pieceView); err != nil {
		return err
	}
	r.pieces[i].kind = pieceHolder
	return nil
}

// Release frees every piece.
func (r *region) Release() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	var errs []error
	for _, p := range r.pieces {
		if err := r.ops.free(p.start, p.kind); err != nil {
			errs = append(errs, err)
		}
	}
	r.pieces = nil
	return errors.Join(errs...)
}

// find returns the index of the piece containing off, or -1.
func (r *region) find(off uint64) int {
	i := sort.Search(len(r.pieces), func(i int) bool { return r.pieces[i].end > off })
	if i == len(r.pieces) || r.pieces[i].start > off {
		return -1
	}
	return i
}

// overlapping returns the index range [i, j) of pieces overlapping [start, end).
func (r *region) overlapping(start, end uint64) (int, int) {
	i := sort.Search(len(r.pieces), func(i int) bool { return r.pieces[i].end > start })
	j := i
	for j < len(r.pieces) && r.pieces[j].start < end {
		j++
	}
	return i, j
}

// makeHolder makes [start, end) exactly one placeholder piece: pieces
// inside it are given back, the pieces at its edges are split, and what
// remains is merged. Locked.
func (r *region) makeHolder(start, end uint64) error {
	if start >= end || end > r.size {
		return fmt.Errorf("aot: bad range [%#x, %#x)", start, end)
	}
	i, j := r.overlapping(start, end)
	for k := i; k < j; k++ {
		p := r.pieces[k]
		if p.kind == pieceHolder {
			continue
		}
		if p.start < start || p.end > end {
			return fmt.Errorf("aot: [%#x, %#x) overlaps an allocation at [%#x, %#x)", start, end, p.start, p.end)
		}
		if err := r.ops.toHolder(p.start, p.end-p.start, p.kind); err != nil {
			return err
		}
		r.pieces[k].kind = pieceHolder
	}
	// trim the edges. VirtualFree(MEM_PRESERVE_PLACEHOLDER) splits a
	// placeholder from its start: the first size bytes stay, the rest
	// becomes a placeholder of its own. So a piece that begins before start
	// is split at start, and then the piece that begins at start (or the
	// last one) is split at end.
	if p := r.pieces[i]; p.start < start {
		if err := r.ops.split(p.start, start-p.start); err != nil {
			return err
		}
		r.splitAt(start)
		i, j = r.overlapping(start, end)
	}
	if p := r.pieces[j-1]; p.end > end {
		if err := r.ops.split(p.start, end-p.start); err != nil {
			return err
		}
		r.splitAt(end)
		i, j = r.overlapping(start, end)
	}
	if j-i > 1 {
		if err := r.ops.coalesce(start, end-start); err != nil {
			return err
		}
		merged := piece{start, end, pieceHolder}
		r.pieces = append(r.pieces[:i], append([]piece{merged}, r.pieces[j:]...)...)
	}
	return nil
}

// splitAt divides the piece containing off at off (bookkeeping only).
func (r *region) splitAt(off uint64) {
	i := r.find(off)
	if i < 0 || r.pieces[i].start == off {
		return
	}
	p := r.pieces[i]
	r.pieces = append(r.pieces[:i], append([]piece{{p.start, off, p.kind}, {off, p.end, p.kind}}, r.pieces[i+1:]...)...)
}

// setKind records the kind of the single piece [start, end). Locked.
func (r *region) setKind(start, end uint64, kind pieceKind) {
	i := r.find(start)
	if i >= 0 && r.pieces[i].end == end {
		r.pieces[i].kind = kind
	}
}

// regions maps a reservation's base address to its bookkeeping.
var (
	regionsMu sync.Mutex
	regions   = map[uintptr]*region{}
)

func registerRegion(base uintptr, r *region) {
	regionsMu.Lock()
	regions[base] = r
	regionsMu.Unlock()
}

func lookupRegion(base uintptr) *region {
	regionsMu.Lock()
	defer regionsMu.Unlock()
	return regions[base]
}

func forgetRegion(base uintptr) {
	regionsMu.Lock()
	delete(regions, base)
	regionsMu.Unlock()
}
