//go:build !windows

package host

import "os"

// newSegmentFile creates the anonymous file behind a shared memory
// segment: unlinked at once, it lives as long as a descriptor or a mapping
// of it does.
func newSegmentFile(size uint32) (*os.File, error) {
	f, err := os.CreateTemp("", "pgmem-shm-")
	if err != nil {
		return nil, err
	}
	os.Remove(f.Name())
	if err := f.Truncate(int64(size)); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}
