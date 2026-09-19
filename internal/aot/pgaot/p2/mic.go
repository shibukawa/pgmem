package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_mic_to_latin4(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13968(m, l0, int32(11), int32(132))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_mic_to_win1250(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13967(m, l0, int32(_a_F_mic_to_win1250_0), int32(29), int32(130))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
