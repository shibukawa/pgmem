package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_leftmostvalue_int2(m *base.Module) int64 {
	return int64(-32768)
}
func F_leftmostvalue_int4(m *base.Module) int64 {
	return int64(-2147483648)
}
func F_leftmostvalue_uuid(m *base.Module) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_palloc0(m, int32(16))
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v2)
	}
}
