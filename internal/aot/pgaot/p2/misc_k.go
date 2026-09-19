package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_koi8r_to_win866(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13958(m, l0, int32(_a_F_koi8r_to_win866_0), int32(20), int32(22))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_koi8u_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13870(m, l0, int32(34), v3, v3, v3, int32(_a_F_koi8u_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
