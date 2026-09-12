package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_float_compare_desc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_gt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F32_lt(v6, v7)
	}
	return v10
}
