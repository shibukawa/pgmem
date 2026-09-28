package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_leftmostvalue_float8(m *base.Module) int64 {
	return int64(-4503599627370496)
}
func F_leftmostvalue_varbit(m *base.Module) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall3Coll(m, int32(2851), int32(0), int64(4166504), int64(0), int64(-1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
