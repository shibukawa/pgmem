package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_preprocess_subquery_phvs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v3 = int32(0)
	if l0 == v3 {
		v60 = v3
		return v60
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 != int32(321) {
			if v8 != int32(67) {
				v56 = F_expression_tree_walker_impl(m, l0, int32(879), l1)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v60 = v56
					return v60
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v13 + int32(1)
				v19 = F_query_tree_walker_impl(m, l0, int32(879), l1, int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v23 - int32(1)
					return v19
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v29) < base.Ui32(v28) {
				v60 = v3
				return v60
			} else {
				if base.B2i32(v28 != v29)|base.B2i32(v29 <= int32(0)) != 0 {
					v56 = F_expression_tree_walker_impl(m, l0, int32(879), l1)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v60 = v56
						return v60
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v36 = F_copyObjectImpl(m, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = int32(0)
						F_IncrementVarSublevelsUp(m, v36, v38-v28, v38)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v45 = F_preprocess_expression(m, v43, v36, int32(8))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_IncrementVarSublevelsUp(m, v45, v28, int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
									return int32(0)
								}
							}
						}
					}
				}
			}
		}
	}
}
