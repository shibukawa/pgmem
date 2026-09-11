package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dist_bs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_box_closept_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_dist_sp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_lseg_closept_point(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
