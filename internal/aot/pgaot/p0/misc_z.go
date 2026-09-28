package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ZeroAndLockBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	if l2 == int32(0) {
		if l0 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[0]))
			v12 = l0 ^ int32(-1)
			v15 = v10 + v12*int32(56)
			v18 = F_StartLocalBufferIO(m, v15, int32(1), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 != int32(2) {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[1]))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v12<<(uint(int32(2))%32))))
					base.MemoryFill(m, v27, int32(0), int32(_a_F_ZeroAndLockBuffer_0))
					v33 = int64(0)
					v37 = base.AtomicRmwCmpxchg64(m, v15, int32(24), v33, v33)
					*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v37&int64(-134217729) | int64(16777216)
					return
				}
			}
		} else {
			v48 = int32(56)
			v49 = l0 * v48
			v51 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
			v54 = v49 + v51 - v48
			v55 = int32(1)
			v58 = F_StartSharedBufferIO(m, v54, v55, v55, int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				if v58 != int32(2) {
					if l1 == int32(1) {
						v93 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
						v94 = int32(56)
						F_BufferLockAcquire(m, l0, v93+l0*v94-v94, int32(3))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							return
						}
					} else {
						F_LockBufferForCleanup(m, l0)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[3]))
					base.MemoryFill(m, v63+l0<<(uint(int32(13))%32)+int32(-8192), int32(0), int32(_a_F_ZeroAndLockBuffer_0))
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
					F_BufferLockAcquire(m, l0, v73+v49-int32(56), int32(3))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v80 = int32(0)
						F_TerminateBufferIO(m, v54, v80, int64(16777216), int32(1), v80)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	} else {
		if l0 < int32(0) {
			return
		} else {
			if l1 == int32(1) {
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
				v94 = int32(56)
				F_BufferLockAcquire(m, l0, v93+l0*v94-v94, int32(3))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					return
				}
			} else {
				F_LockBufferForCleanup(m, l0)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
