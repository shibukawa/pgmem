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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	if l2 == int32(0) {
		if l0 < int32(0) {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[0]))
			v12 = l0 ^ int32(-1)
			v15 = v10 + v12<<(uint(int32(6))%32)
			v17 = F_StartLocalBufferIO(m, v15, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 == int32(0) {
					return
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[1]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v12<<(uint(int32(2))%32))))
					v30 = F__emscripten_memset_bulkmem(m, v26, base.I32_extend8_s(int32(0)), int32(_a_F_ZeroAndLockBuffer_0))
					mBase = m.M
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v34&int32(-134217729) | int32(16777216)
					return
				}
			}
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
			v49 = v46 + l0<<(uint(int32(6))%32)
			v51 = v49 + int32(-64)
			v54 = F_StartBufferIO(m, v51, int32(1), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				if v54 == int32(0) {
					if l1 == int32(1) {
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
						v94 = F_LWLockAcquire(m, v87+l0<<(uint(int32(6))%32)-int32(16), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							return
						}
					} else {
						F_LockBufferForCleanup(m, l0)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[3]))
					v68 = F__emscripten_memset_bulkmem(m, v59+l0<<(uint(int32(13))%32)+int32(-8192), base.I32_extend8_s(int32(0)), int32(_a_F_ZeroAndLockBuffer_0))
					mBase = m.M
					v72 = F_LWLockAcquire(m, v49-int32(16), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						v74 = int32(0)
						F_TerminateBufferIO(m, v51, v74, int32(16777216), int32(1), v74)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_ZeroAndLockBuffer[2]))
				v94 = F_LWLockAcquire(m, v87+l0<<(uint(int32(6))%32)-int32(16), int32(0))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					return
				}
			} else {
				F_LockBufferForCleanup(m, l0)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
