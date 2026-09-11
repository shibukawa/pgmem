package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_xidLogicalComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v5)) == int32(0) {
		v18 = base.B2i32(base.Ui32(v5) < base.Ui32(v6))
	} else {
		v18 = int32(base.Ui32(v5-v6) >> (uint(int32(31)) % 32))
	}
	if v18 != 0 {
		v31 = int32(-1)
	} else {
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v5))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v6)) == int32(0) {
			v30 = base.B2i32(base.Ui32(v6) < base.Ui32(v5))
		} else {
			v30 = int32(base.Ui32(v6-v5) >> (uint(int32(31)) % 32))
		}
		v31 = v30
	}
	return v31
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
	v6 = F_uint32in_subr(m, v2, int32(0), int32(406622), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
