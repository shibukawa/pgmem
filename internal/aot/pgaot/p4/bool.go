package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_decrement(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 == v4))
	return v4
}
