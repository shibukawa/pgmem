package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockNewTrancheId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[810]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	if v8 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[810]))
		F_s_lock(m, v12, int32(518076), int32(622), int32(483494))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = v5 - int32(4)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 + int32(1)
			v27 = *(*int32)(unsafe.Add(mBase, _consts[810]))
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(0)
			return v22
		}
	} else {
		v21 = v5 - int32(4)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 + int32(1)
		v27 = *(*int32)(unsafe.Add(mBase, _consts[810]))
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(0)
		return v22
	}
}
func F_LWLockRegisterTranche(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	if int32(95) <= l0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[808]))
		v11 = l0 - int32(95)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[809]))
		if v13 <= v11 {
			v17 = int32(102)
			if base.Ui32(l0) <= base.Ui32(v17) {
				v20 = v17
			} else {
				v20 = l0
			}
			v22 = v20 - int32(94)
			if v22&(v20-int32(95)) != 0 {
				v29 = int32(1) << (uint(int32(32)-base.I32_clz(v22)) % 32)
			} else {
				v29 = v22
			}
			v31 = v29 << (uint(int32(2)) % 32)
			if v9 == int32(0) {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[11]))
				v36 = F_MemoryContextAllocZero(m, v35, v31)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v42 = v36
					*(*int32)(unsafe.Add(mBase, _consts[809])) = v29
					*(*int32)(unsafe.Add(mBase, _consts[808])) = v42
					v48 = v42
					*(*int32)(unsafe.Add(mBase, uint32(v48+v11<<(uint(int32(2))%32)))) = int32(448270)
					return
				}
			} else {
				v40 = F_repalloc0(m, v9, v13<<(uint(int32(2))%32), v31)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = v40
					*(*int32)(unsafe.Add(mBase, _consts[809])) = v29
					*(*int32)(unsafe.Add(mBase, _consts[808])) = v42
					v48 = v42
					*(*int32)(unsafe.Add(mBase, uint32(v48+v11<<(uint(int32(2))%32)))) = int32(448270)
					return
				}
			}
		} else {
			v48 = v9
			*(*int32)(unsafe.Add(mBase, uint32(v48+v11<<(uint(int32(2))%32)))) = int32(448270)
			return
		}
	} else {
		return
	}
}
