package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomStringVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	v11 = F_init_custom_variable(m, l0, l1, int32(0), l4, l5, int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(v11)+100)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l2
		F_define_custom_variable(m, v11)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			return
		}
	}
}
