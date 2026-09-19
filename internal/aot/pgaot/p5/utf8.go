package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_utf8_to_euc_cn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(2), v3, v3, v3, int32(_a_F_utf8_to_euc_cn_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_euc_tw(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(4), v3, v3, v3, int32(_a_F_utf8_to_euc_tw_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_gb18030(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(0)
	v7 = Fn14015(m, l0, int32(39), int32(_a_F_utf8_to_gb18030_0), v4, v4, int32(_a_F_utf8_to_gb18030_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
