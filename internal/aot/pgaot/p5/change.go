package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChangeVarNodesExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	if l0 == v5 {
		v70 = F_ChangeVarNodes_walker(m, l0, v10)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v19 != int32(67) {
			v70 = F_ChangeVarNodes_walker(m, l0, v10)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if l1 == v22 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l2
			} else {
			}
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if l1 == v25 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l2
			} else {
			}
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v28 == int32(0) {
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
				if v31 != l1 {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l2
				}
			}
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			if v34 == int32(0) {
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				if v37 <= int32(0) {
				} else {
					v45 = v5
					for {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v45<<(uint(int32(2))%32))))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
						if l1 == v52 {
							*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l2
						} else {
						}
						v56 = v45 + int32(1)
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						if v56 < v57 {
							v45 = v56
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v68 = F_query_tree_walker_impl(m, l0, int32(1049), v10, int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_changeDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v7 = F_table_open(m, int32(1214), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		F_shdepChangeDep(m, v7, l0, l1, int32(1260), l2, int32(111))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_shdepDropDependency(m, v7, l0, l1, int32(0), int32(1), int32(1260), l2, int32(97))
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_relation_close(m, v7, int32(3))
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
