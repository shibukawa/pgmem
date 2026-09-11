package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_show_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v9 = F_set_deparse_context_plan(m, v7, v8, l3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = F_deparse_expression(m, l0, v9, l4, int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_ExplainPropertyText(m, l1, v12, l5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	}
}
