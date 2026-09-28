package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_leftmostvalue_bit(m *base.Module) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_DirectFunctionCall3Coll(m, int32(525), int32(0), int64(4166504), int64(0), int64(-1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_leftmostvalue_interval(m *base.Module) int64 {
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
		*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(-9223372034707292160)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(-9223372036854775807 - 1)
		return base.I64_extend_i32_u(v3)
	}
}
