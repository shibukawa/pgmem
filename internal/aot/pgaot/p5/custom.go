package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomStringVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	v12 = F_init_custom_variable(m, l0, l1, int32(0), int32(6), int32(1), int32(3), int32(120))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = l2
		F_define_custom_variable(m, v12)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
