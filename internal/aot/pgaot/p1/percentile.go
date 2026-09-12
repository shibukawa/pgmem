package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_percentile_cont_interval_multi_final(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_percentile_cont_multi_final_common(m, l0, int32(1186), int32(16), int32(1473))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
