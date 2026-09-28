package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomBoolVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	v5 = l4
	v7 = int32(0)
	v9 = F_init_custom_variable(m, l0, l1, l2, l5, v7, v7)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+116)) = uint8(v5)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+100)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = l3
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v14
		F_define_custom_variable(m, v9)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
