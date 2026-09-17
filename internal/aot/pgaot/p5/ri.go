package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_setdefault_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setdefault_upd_0), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(0), int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_setnull_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setnull_del_0), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(1), int32(3))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_trigger_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(1644)
	if base.Ui32(v3) <= base.Ui32(int32(11)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_RI_FKey_trigger_type[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
