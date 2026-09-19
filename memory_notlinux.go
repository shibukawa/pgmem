//go:build !linux

package pgmem

func cgroupMemoryLimit() uint64 { return 0 }
