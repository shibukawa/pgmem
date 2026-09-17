package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xidLogicalComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v4))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3)) == int32(0) {
		v16 = base.B2i32(base.Ui32(v3) < base.Ui32(v4))
	} else {
		v16 = int32(base.Ui32(v3-v4) >> (uint(int32(31)) % 32))
	}
	if v16 != 0 {
		v30 = int32(-1)
	} else {
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v4)) == int32(0) {
			v29 = base.B2i32(base.Ui32(v4) < base.Ui32(v3))
		} else {
			v29 = int32(base.Ui32(v4-v3) >> (uint(int32(31)) % 32))
		}
		v30 = v29
	}
	return v30
}
func F_xidin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_uint32in_subr(m, v2, int32(0), int32(_a_F_xidin_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
