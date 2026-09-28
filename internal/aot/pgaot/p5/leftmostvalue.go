package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_leftmostvalue_int8(m *base.Module) int64 {
	return int64(-9223372036854775807 - 1)
}
func F_leftmostvalue_name(m *base.Module) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_palloc0(m, int32(64))
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v2)
	}
}
func F_leftmostvalue_text(m *base.Module) int64 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text_with_len(m, int32(_a_F_leftmostvalue_text_0), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_leftmostvalue_timetz(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(-86400)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
		return base.I64_extend_i32_u(v3)
	}
}
