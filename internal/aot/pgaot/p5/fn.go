package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_fn_expr_rettype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 == v2 {
		v13 = v2
		return v13
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v5 == int32(0) {
			v13 = v2
			return v13
		} else {
			v8 = F_exprType(m, v5)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v13 = v8
				return v13
			}
		}
	}
}
