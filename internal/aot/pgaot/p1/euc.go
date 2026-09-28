package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_euc_tw_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(4), v3, v3, v3, int32(_a_F_euc_tw_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
