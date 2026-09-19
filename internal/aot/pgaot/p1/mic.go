package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_mic_to_win1251(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13947(m, l0, int32(_a_F_mic_to_win1251_0), int32(23), int32(139))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
