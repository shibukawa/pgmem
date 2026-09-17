package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_language_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13899(m, l0, l1, int32(_a_F_get_language_oid_0), int32(235), int32(_a_F_get_language_oid_1), int32(_a_F_get_language_oid_2), int32(67137668), int32(35))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
