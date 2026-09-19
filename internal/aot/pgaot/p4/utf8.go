package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_utf8_to_gbk(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(37), v3, v3, v3, int32(_a_F_utf8_to_gbk_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_shift_jis_2004(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14015(m, l0, int32(41), int32(0), int32(25), int32(_a_F_utf8_to_shift_jis_2004_0), int32(_a_F_utf8_to_shift_jis_2004_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_uhc(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(38), v3, v3, v3, int32(_a_F_utf8_to_uhc_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_win(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14016(m, l0, int32(_a_F_utf8_to_win_0), int32(_a_F_utf8_to_win_1), int32(150), int32(_a_F_utf8_to_win_2), int32(_a_F_utf8_to_win_3), int32(_a_F_utf8_to_win_4), int32(15), int32(18))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
