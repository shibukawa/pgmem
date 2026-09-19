package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_language_handler_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_language_handler_out_0), int32(368), int32(_a_F_language_handler_out_1), int32(_a_F_language_handler_out_2), int32(_a_F_language_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
