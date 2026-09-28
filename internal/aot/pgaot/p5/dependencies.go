package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_collectDependenciesOfExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v3
	v14 = F_list_make1_impl(m, int32(1), v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v14
		v19 = F_find_expr_references_walker(m, l1, v6+int32(8))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
