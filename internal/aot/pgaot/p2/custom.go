package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomBoolVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	v4 = l3
	v6 = int32(0)
	v10 = F_init_custom_variable(m, l0, l1, v6, l4, v6, v6, int32(120))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+112)) = uint8(v4)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+96)) = uint8(v4)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l2
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v15
		F_define_custom_variable(m, v10)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			return
		}
	}
}
