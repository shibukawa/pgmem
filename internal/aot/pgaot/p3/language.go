package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_has_language_privilege_name_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13917(m, l0, int32(_a_F_has_language_privilege_name_id_0), int32(2612))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
