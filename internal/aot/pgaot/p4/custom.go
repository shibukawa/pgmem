package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomRealVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64, l6 float64, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	v12 = F_init_custom_variable(m, l0, l1, l2, l7, int32(0), int32(2), int32(152))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v12)+136)) = l4
		*(*float64)(unsafe.Add(mBase, uint32(v12)+96)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = l3
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v12)+124)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v17
		*(*float64)(unsafe.Add(mBase, uint32(v12)+112)) = l6
		*(*float64)(unsafe.Add(mBase, uint32(v12)+104)) = l5
		F_define_custom_variable(m, v12)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			return
		}
	}
}
