package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_opfamily_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13923(m, l0, l1, l2, l3, int32(4))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
