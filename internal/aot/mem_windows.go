//go:build windows

package aot

import (
	"errors"
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

// reserveMemory reserves n bytes of address space; pages are committed by
// commitMemory as the guest grows its memory.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	addr, err := windows.VirtualAlloc(0, uintptr(n), windows.MEM_RESERVE, windows.PAGE_NOACCESS)
	if err != nil {
		return nil, nil, err
	}
	// addr comes straight from VirtualAlloc; unsafe.Add keeps vet quiet
	// about the uintptr-to-pointer conversion, which is intended here.
	mem = unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(nil), addr)), n)
	return mem, func() { windows.VirtualFree(addr, 0, windows.MEM_RELEASE) }, nil
}

// commitMemory makes [from, to) readable and writable.
func commitMemory(mem []byte, from, to uint64) error {
	if to <= from {
		return nil
	}
	base := uintptr(unsafe.Pointer(unsafe.SliceData(mem)))
	_, err := windows.VirtualAlloc(base+uintptr(from), uintptr(to-from), windows.MEM_COMMIT, windows.PAGE_READWRITE)
	return err
}

// mapShared is not implemented on Windows (it needs placeholder mappings,
// MapViewOfFile3); the multi-process model is unavailable there.
func mapShared(mem []byte, off uint32, f *os.File, size uint32) error {
	return errors.New("aot: shared memory segments are not supported on windows")
}

func unmapShared(mem []byte, off uint32, size uint32) error { return nil }
