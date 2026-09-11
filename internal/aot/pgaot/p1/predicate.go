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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v4 = int32(0)
	if l0 == v4 {
		v26 = v4
		return v26
	} else {
		if l1 == int32(0) {
			v26 = v4
			return v26
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v9 == int32(1) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v14 = v13
			} else {
				v14 = l0
			}
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v15 == int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v20 = v19
			} else {
				v20 = l1
			}
			v21 = F_predicate_refuted_by_recurse(m, v20, v14, l2)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = v21
				return v26
			}
		}
	}
}
