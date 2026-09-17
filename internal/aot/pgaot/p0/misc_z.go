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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	if l2 == int32(0) {
		if int32(0) <= l0 {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[0]))
			v13 = v10 + l0<<(uint(int32(6))%32)
			v15 = v13 + int32(-64)
			v18 = F_StartBufferIO(m, v15, int32(1), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 == int32(0) {
					if l1 == int32(1) {
						v85 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[0]))
						v92 = F_LWLockAcquire(m, v85+l0<<(uint(int32(6))%32)-int32(16), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							return
						}
					} else {
						F_LockBufferForCleanup(m, l0)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[1]))
					v29 = int32(0)
					base.MemoryFill(m, v23+l0<<(uint(int32(13))%32)+int32(-8192), v29, int32(_a_F_ZeroAndLockBuffer_0))
					v35 = F_LWLockAcquire(m, v13-int32(16), v29)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = int32(0)
						F_TerminateBufferIO(m, v15, v37, int32(16777216), int32(1), v37)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
			v46 = l0 ^ int32(-1)
			v49 = v44 + v46<<(uint(int32(6))%32)
			v51 = F_StartLocalBufferIO(m, v49, int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				if v51 == int32(0) {
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[3]))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v46<<(uint(int32(2))%32))))
					base.MemoryFill(m, v60, int32(0), int32(_a_F_ZeroAndLockBuffer_0))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v67&int32(-134217729) | int32(16777216)
					return
				}
			}
		}
	} else {
		if l0 < int32(0) {
			return
		} else {
			if l1 == int32(1) {
				v85 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[0]))
				v92 = F_LWLockAcquire(m, v85+l0<<(uint(int32(6))%32)-int32(16), int32(0))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					return
				}
			} else {
				F_LockBufferForCleanup(m, l0)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
