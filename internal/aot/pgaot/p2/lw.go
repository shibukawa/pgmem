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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	if v8 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
		F_s_lock(m, v12, int32(_a_F_LWLockNewTrancheId_0), int32(622), int32(_a_F_LWLockNewTrancheId_1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = v5 - int32(4)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 + int32(1)
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(0)
			return v22
		}
	} else {
		v21 = v5 - int32(4)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 + int32(1)
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(0)
		return v22
	}
}
func F_LWLockRegisterTranche(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	if int32(95) <= l0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[0]))
		v12 = l0 - int32(95)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[1]))
		if v14 <= v12 {
			v18 = int32(102)
			if base.Ui32(l0) <= base.Ui32(v18) {
				v21 = v18
			} else {
				v21 = l0
			}
			v23 = v21 - int32(94)
			if v23&(v21-int32(95)) != 0 {
				v30 = int32(1) << (uint(int32(32)-base.I32_clz(v23)) % 32)
			} else {
				v30 = v23
			}
			v32 = v30 << (uint(int32(2)) % 32)
			if v10 == int32(0) {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[2]))
				v37 = F_MemoryContextAllocZero(m, v36, v32)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v43 = v37
					*(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[1])) = v30
					*(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[0])) = v43
					v49 = v43
					*(*int32)(unsafe.Add(mBase, uint32(v49+v12<<(uint(int32(2))%32)))) = l1
					return
				}
			} else {
				v41 = F_repalloc0(m, v10, v14<<(uint(int32(2))%32), v32)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = v41
					*(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[1])) = v30
					*(*int32)(unsafe.Add(mBase, _c_F_LWLockRegisterTranche[0])) = v43
					v49 = v43
					*(*int32)(unsafe.Add(mBase, uint32(v49+v12<<(uint(int32(2))%32)))) = l1
					return
				}
			}
		} else {
			v49 = v10
			*(*int32)(unsafe.Add(mBase, uint32(v49+v12<<(uint(int32(2))%32)))) = l1
			return
		}
	} else {
		return
	}
}
