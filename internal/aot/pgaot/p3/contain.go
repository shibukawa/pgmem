package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_contain_dml_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(67) {
			v11 = int32(1)
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12 != v11 {
				v26 = v11
				return v26
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
				if v15 != 0 {
					v26 = v11
					return v26
				} else {
					v18 = F_query_tree_walker_impl(m, l0, int32(842), l1, int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v18
					}
				}
			}
		} else {
			v24 = F_expression_tree_walker_impl(m, l0, int32(842), l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = v24
				return v26
			}
		}
	}
}
func F_contain_placeholder_walker(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13868(m, l0, l1, int32(1488), int32(319))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
