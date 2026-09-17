package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomStringVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v10 = F_init_custom_variable(m, l0, l1, int32(0), int32(6), int32(1), int32(3), int32(120))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = int32(_a_F_DefineCustomStringVariable_0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(_a_F_DefineCustomStringVariable_1)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = l2
		F_define_custom_variable(m, v10)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
