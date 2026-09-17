package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_opfamily_method(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13828(m, l0, int32(_a_F_get_opfamily_method_0), int32(1384), int32(_a_F_get_opfamily_method_1), int32(_a_F_get_opfamily_method_2), int32(42))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
