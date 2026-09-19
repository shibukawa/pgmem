//go:build unix

package aot

import (
	"os"
	"syscall"
	"unsafe"
)

// reserveMemory maps n bytes of anonymous memory outside the Go heap. Pages
// are materialized by the OS on first touch, so a large reservation costs
// nothing until the guest uses it, and release returns everything at once.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	mem, err = syscall.Mmap(-1, 0, n, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, nil, err
	}
	return mem, func() { syscall.Munmap(mem) }, nil
}

// commitMemory is a no-op on unix: anonymous mappings commit lazily.
func commitMemory(mem []byte, from, to uint64) error { return nil }

// mapShared maps f over [off, off+size) of mem, shared with every other
// mapping of f (System V shared memory between the processes of a
// cluster). off and size are multiples of 64 KiB, which satisfies every
// page size in use.
func mapShared(mem []byte, off uint32, f *os.File, size uint32) error {
	base := uintptr(unsafe.Pointer(unsafe.SliceData(mem))) + uintptr(off)
	_, _, e := syscall.Syscall6(syscall.SYS_MMAP, base, uintptr(size),
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED|syscall.MAP_FIXED, f.Fd(), 0)
	if e != 0 {
		return e
	}
	return nil
}

// unmapShared puts private zero pages back over [off, off+size).
func unmapShared(mem []byte, off uint32, size uint32) error {
	base := uintptr(unsafe.Pointer(unsafe.SliceData(mem))) + uintptr(off)
	_, _, e := syscall.Syscall6(syscall.SYS_MMAP, base, uintptr(size),
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE|syscall.MAP_ANON|syscall.MAP_FIXED, ^uintptr(0), 0)
	if e != 0 {
		return e
	}
	return nil
}
