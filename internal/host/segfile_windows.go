//go:build windows

package host

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"

	"golang.org/x/sys/windows"
)

var segSeq atomic.Int64

// newSegmentFile creates the file behind a shared memory segment. Windows
// refuses to delete a file that is open, so it is created delete-on-close
// instead of unlinked: it goes when the last handle, including the file
// mappings the instances hold, is closed.
func newSegmentFile(size uint32) (*os.File, error) {
	name := filepath.Join(os.TempDir(), fmt.Sprintf("pgmem-shm-%d-%d", os.Getpid(), segSeq.Add(1)))
	p, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}
	h, err := windows.CreateFile(p, windows.GENERIC_READ|windows.GENERIC_WRITE, 0, nil, windows.CREATE_NEW,
		windows.FILE_ATTRIBUTE_TEMPORARY|windows.FILE_FLAG_DELETE_ON_CLOSE, 0)
	if err != nil {
		return nil, fmt.Errorf("create %s: %w", name, err)
	}
	f := os.NewFile(uintptr(h), name)
	if err := f.Truncate(int64(size)); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}
