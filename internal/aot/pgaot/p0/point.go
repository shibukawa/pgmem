package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_point_below(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_point_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v13 float64
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v10
		*(*float64)(unsafe.Add(mBase, uint32(v6))) = v10
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v6)+24)) = v13
		*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		return v6
	}
}
func F_point_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = F_pair_decode(m, v4, v6, v6+int32(8), int32(0), int32(88893), v4, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v6
		}
	}
}
