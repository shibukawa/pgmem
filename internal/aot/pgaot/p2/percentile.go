package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_percentile_cont_float8_multi_final(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_percentile_cont_multi_final_common(m, l0, int32(701), int32(8), int32(1), int32(1602))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
