//go:build windows

package aot

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Windows linear memory: one placeholder reservation per instance
// (VirtualAlloc2 with MEM_RESERVE_PLACEHOLDER). The heap is committed into
// it piece by piece and the shared segments of a cluster are views of
// their files mapped into it at fixed addresses (MapViewOfFile3 with
// MEM_REPLACE_PLACEHOLDER), which is what mmap(MAP_FIXED) does on Unix.
// The bookkeeping of the pieces is in placeholder.go; these are the calls.
// Needs Windows 10 1803 / Server 2016 or later.

const (
	memCoalescePlaceholders = 0x00000001
	memPreservePlaceholder  = 0x00000002
	memReplacePlaceholder   = 0x00004000
	memReservePlaceholder   = 0x00040000
)

var (
	kernelbase           = windows.NewLazySystemDLL("kernelbase.dll")
	procVirtualAlloc2    = kernelbase.NewProc("VirtualAlloc2")
	procMapViewOfFile3   = kernelbase.NewProc("MapViewOfFile3")
	procUnmapViewOfFile2 = kernelbase.NewProc("UnmapViewOfFile2")
)

func virtualAlloc2(addr, size uintptr, allocType, protect uint32) (uintptr, error) {
	r, _, e := procVirtualAlloc2.Call(uintptr(windows.CurrentProcess()), addr, size, uintptr(allocType), uintptr(protect), 0, 0)
	if r == 0 {
		return 0, fmt.Errorf("VirtualAlloc2(%#x, %#x, %#x): %w", addr, size, allocType, e)
	}
	return r, nil
}

func mapViewOfFile3(mapping windows.Handle, addr, size uintptr, allocType, protect uint32) error {
	r, _, e := procMapViewOfFile3.Call(uintptr(mapping), uintptr(windows.CurrentProcess()), addr, 0, size, uintptr(allocType), uintptr(protect), 0, 0)
	if r == 0 {
		return fmt.Errorf("MapViewOfFile3(%#x, %#x): %w", addr, size, e)
	}
	return nil
}

func unmapViewOfFile2(addr uintptr, flags uint32) error {
	r, _, e := procUnmapViewOfFile2.Call(uintptr(windows.CurrentProcess()), addr, uintptr(flags))
	if r == 0 {
		return fmt.Errorf("UnmapViewOfFile2(%#x, %#x): %w", addr, flags, e)
	}
	return nil
}

// winOps implements pageOps for one reservation at base.
type winOps struct{ base uintptr }

func (w winOps) split(addr, size uint64) error {
	if err := windows.VirtualFree(w.base+uintptr(addr), uintptr(size), windows.MEM_RELEASE|memPreservePlaceholder); err != nil {
		return fmt.Errorf("split placeholder at %#x+%#x: %w", addr, size, err)
	}
	return nil
}

func (w winOps) coalesce(addr, size uint64) error {
	if err := windows.VirtualFree(w.base+uintptr(addr), uintptr(size), windows.MEM_RELEASE|memCoalescePlaceholders); err != nil {
		return fmt.Errorf("coalesce placeholders at %#x+%#x: %w", addr, size, err)
	}
	return nil
}

func (w winOps) commit(addr, size uint64) error {
	_, err := virtualAlloc2(w.base+uintptr(addr), uintptr(size), windows.MEM_RESERVE|windows.MEM_COMMIT|memReplacePlaceholder, windows.PAGE_READWRITE)
	return err
}

func (w winOps) mapView(addr, size uint64, f *os.File) error {
	h, err := windows.CreateFileMapping(windows.Handle(f.Fd()), nil, windows.PAGE_READWRITE, uint32(size>>32), uint32(size), nil)
	if err != nil {
		return fmt.Errorf("CreateFileMapping: %w", err)
	}
	defer windows.CloseHandle(h) // the view keeps the section alive
	return mapViewOfFile3(h, w.base+uintptr(addr), uintptr(size), memReplacePlaceholder, windows.PAGE_READWRITE)
}

func (w winOps) toHolder(addr, size uint64, kind pieceKind) error {
	if kind == pieceView {
		return unmapViewOfFile2(w.base+uintptr(addr), memPreservePlaceholder)
	}
	// "Frees an allocation back to a placeholder": the documentation does
	// not say whether dwSize is the allocation's size (as for a split) or
	// 0 (as for MEM_RELEASE); accept either answer.
	err := windows.VirtualFree(w.base+uintptr(addr), uintptr(size), windows.MEM_RELEASE|memPreservePlaceholder)
	if err != nil {
		err = windows.VirtualFree(w.base+uintptr(addr), 0, windows.MEM_RELEASE|memPreservePlaceholder)
	}
	if err != nil {
		return fmt.Errorf("release to placeholder at %#x+%#x: %w", addr, size, err)
	}
	return nil
}

func (w winOps) free(addr uint64, kind pieceKind) error {
	if kind == pieceView {
		return unmapViewOfFile2(w.base+uintptr(addr), 0)
	}
	return windows.VirtualFree(w.base+uintptr(addr), 0, windows.MEM_RELEASE)
}

// reserveMemory reserves n bytes of address space as a placeholder; pages
// are committed by commitMemory as the guest grows its memory.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	addr, err := virtualAlloc2(0, uintptr(n), windows.MEM_RESERVE|memReservePlaceholder, windows.PAGE_NOACCESS)
	if err != nil {
		return nil, nil, err
	}
	// addr comes straight from VirtualAlloc2; unsafe.Add keeps vet quiet
	// about the uintptr-to-pointer conversion, which is intended here.
	mem = unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(nil), addr)), n)
	r := newRegion(uint64(n), winOps{base: addr})
	registerRegion(addr, r)
	return mem, func() {
		forgetRegion(addr)
		r.Release()
	}, nil
}

func regionOf(mem []byte) (*region, error) {
	base := uintptr(unsafe.Pointer(unsafe.SliceData(mem)))
	r := lookupRegion(base)
	if r == nil {
		return nil, fmt.Errorf("aot: no reservation at %#x", base)
	}
	return r, nil
}

// commitMemory makes [from, to) readable and writable.
func commitMemory(mem []byte, from, to uint64) error {
	if to <= from {
		return nil
	}
	r, err := regionOf(mem)
	if err != nil {
		return err
	}
	return r.Commit(from, to)
}

// mapShared maps f over [off, off+size) of mem, shared with every other
// mapping of f. off and size are multiples of 64 KiB, the allocation
// granularity.
func mapShared(mem []byte, off uint32, f *os.File, size uint32) error {
	r, err := regionOf(mem)
	if err != nil {
		return err
	}
	return r.Map(uint64(off), uint64(size), f)
}

// unmapShared gives [off, off+size) back to a placeholder. Unlike the Unix
// version it leaves the range inaccessible, which a detached segment is
// never touched through anyway.
func unmapShared(mem []byte, off uint32, size uint32) error {
	r, err := regionOf(mem)
	if err != nil {
		return err
	}
	return r.Unmap(uint64(off), uint64(size))
}
