package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineCustomRealVariable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 float64, l6 float64, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	v11 = F_init_custom_variable(m, l0, l1, l2, l7, int32(0), int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v11)+144)) = l4
		*(*float64)(unsafe.Add(mBase, uint32(v11)+104)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = l3
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v16
		*(*float64)(unsafe.Add(mBase, uint32(v11)+120)) = l6
		*(*float64)(unsafe.Add(mBase, uint32(v11)+112)) = l5
		F_define_custom_variable(m, v11)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			return
		}
	}
}
