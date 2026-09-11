package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_fn_opclass_options(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = int32(-1)
	v5 = int32(0)
	v10 = F_makeConst(m, int32(17), v4, v5, v4, l1, base.B2i32(l1 == v5), v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v10
		return
	}
}
