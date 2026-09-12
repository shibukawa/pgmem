package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomIntVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	v11 = F_init_custom_variable(m, l0, l1, l2, int32(6), int32(0), int32(1), int32(128))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = l3
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+116)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = int32(1)
		F_define_custom_variable(m, v11)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExecCustomScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v3 != 0 {
		F_ProcessInterrupts(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v10 = m.T0[v9].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
