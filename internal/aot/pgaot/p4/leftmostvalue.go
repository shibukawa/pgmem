package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_leftmostvalue_inet(m *base.Module) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1678), int32(0), int64(615723))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_leftmostvalue_macaddr(m *base.Module) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_palloc0(m, int32(6))
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v2)
	}
}
func F_leftmostvalue_macaddr8(m *base.Module) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_palloc0(m, int32(8))
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v2)
	}
}
func F_leftmostvalue_oid(m *base.Module) int64 {
	return int64(0)
}
