//go:build !unix && !windows

package aot

import (
	"errors"
	"os"
)

// reserveMemory falls back to the Go heap where no lazy mapping is
// available.
func reserveMemory(n int) (mem []byte, release func(), err error) {
	return make([]byte, n), func() {}, nil
}

func commitMemory(mem []byte, from, to uint64) error { return nil }

func mapShared(mem []byte, off uint32, f *os.File, size uint32) error {
	return errors.New("aot: shared memory segments are not supported on this platform")
}

func unmapShared(mem []byte, off uint32, size uint32) error { return nil }
