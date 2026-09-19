//go:build !linux && !darwin && !windows

package pgmem

func physicalMemory() uint64 { return 0 }
