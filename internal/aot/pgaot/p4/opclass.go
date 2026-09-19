package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_opclass_input_type(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14014(m, l0, int32(_a_F_get_opclass_input_type_0), int32(1312), int32(_a_F_get_opclass_input_type_1), int32(_a_F_get_opclass_input_type_2), int32(14))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
