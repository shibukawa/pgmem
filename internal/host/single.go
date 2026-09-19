package host

import "github.com/shibukawa/pgmem/internal/vfs"

// Single-user mode has one process and no cluster, but the same
// PostgreSQL build: with EXEC_BACKEND its shared memory is System V
// (shmget) and its semaphores are POSIX ones, so serve both from the
// process's own heap, the way PGlite's shims did.

type localSeg struct {
	key  int32
	size uint32
	addr uint32
}

func (h *Host) localShmget(m Memory, key int32, size uint32, flags int32) int32 {
	if key != 0 {
		for id, seg := range h.localSegs {
			if seg.key == key {
				if flags&(ipcCreat|ipcExcl) == ipcCreat|ipcExcl {
					return -int32(vfs.EEXIST)
				}
				return id
			}
		}
		if flags&ipcCreat == 0 {
			return -int32(vfs.ENOENT)
		}
	}
	if h.Guest == nil {
		return -int32(vfs.ENOSYS)
	}
	addr, err := h.Guest.Memalign(65536, (size+65535)&^65535)
	if err != nil || addr == 0 {
		return -int32(vfs.ENOMEM)
	}
	if h.localSegs == nil {
		h.localSegs = map[int32]*localSeg{}
	}
	h.nextLocalSeg++
	h.localSegs[h.nextLocalSeg] = &localSeg{key: key, size: size, addr: addr}
	return h.nextLocalSeg
}

func (h *Host) localShmat(id int32) int32 {
	seg := h.localSegs[id]
	if seg == nil {
		return -int32(vfs.EINVAL)
	}
	return int32(seg.addr)
}

func (h *Host) localShmctl(id int32, cmd int32) (segsz, nattch, ret int32) {
	seg := h.localSegs[id]
	if seg == nil {
		return 0, 0, -int32(vfs.EINVAL)
	}
	switch cmd {
	case ipcRMID:
		delete(h.localSegs, id)
		return 0, 0, 0
	case ipcStat:
		return int32(seg.size), 1, 0
	}
	return 0, 0, 0
}

// localSem is a semaphore nobody else can post: a wait that would block
// fails instead of hanging the only process.
func (h *Host) localSem(op int32, addr uint32, arg int32) int32 {
	if h.localSems == nil {
		h.localSems = map[uint32]int32{}
	}
	switch op {
	case semInit:
		h.localSems[addr] = arg
		return 0
	case semDestroy:
		delete(h.localSems, addr)
		return 0
	case semPost:
		h.localSems[addr]++
		return 0
	case semWait, semTryWait:
		if h.localSems[addr] > 0 {
			h.localSems[addr]--
			return 0
		}
		return -int32(vfs.EAGAIN)
	}
	return -int32(vfs.EINVAL)
}
