package pgmem

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func physicalMemory() uint64 {
	var si unix.Sysinfo_t
	if err := unix.Sysinfo(&si); err != nil {
		return 0
	}
	return uint64(si.Totalram) * uint64(si.Unit)
}

// cgroupMemoryLimit is the memory limit of the process's cgroup (v2, then
// v1), which is how containers cap memory; 0 when there is none. Inside a
// container the cgroup namespace makes these files the container's own;
// on a bare host they belong to the root cgroup, which has no limit.
func cgroupMemoryLimit() uint64 {
	for _, path := range []string{"/sys/fs/cgroup/memory.max", "/sys/fs/cgroup/memory/memory.limit_in_bytes"} {
		b, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		n, err := strconv.ParseUint(strings.TrimSpace(string(b)), 10, 64)
		if err != nil {
			continue // v2 writes "max" for no limit
		}
		if n >= 1<<60 {
			continue // v1 writes a value near 2^63 for no limit
		}
		return n
	}
	return 0
}
