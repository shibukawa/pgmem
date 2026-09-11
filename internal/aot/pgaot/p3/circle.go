package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_circle_radius(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+16))
	v4 = F_Float8GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
