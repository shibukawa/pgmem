package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_utf8_to_johab(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14035(m, l0, int32(40), v3, v3, v3, int32(_a_F_utf8_to_johab_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_koi8u(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14035(m, l0, int32(34), v3, v3, v3, int32(_a_F_utf8_to_koi8u_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
