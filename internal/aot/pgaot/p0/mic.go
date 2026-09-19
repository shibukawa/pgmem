package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_mic_to_latin2(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13968(m, l0, int32(9), int32(130))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
