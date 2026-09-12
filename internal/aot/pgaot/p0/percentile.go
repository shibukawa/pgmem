package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_percentile_cont_float8_final(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_percentile_cont_final_common(m, l0, int32(1475))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
