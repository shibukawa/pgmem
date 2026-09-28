package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	v4 = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 == v4))
	return v4
}
func F_bool_int4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return base.I64_extend_i32_u(base.B2i32(v2 != int64(0)))
}
