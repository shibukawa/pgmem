package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_predicate_refuted_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = int32(0)
	if base.B2i32(l0 == v4)|base.B2i32(l1 == v4) != 0 {
		v28 = int32(0)
		return v28
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v11 == int32(1) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v16 = v15
		} else {
			v16 = l0
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v17 == int32(1) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v22 = v21
		} else {
			v22 = l1
		}
		v23 = F_predicate_refuted_by_recurse(m, v22, v16, l2)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = v23
			return v28
		}
	}
}
