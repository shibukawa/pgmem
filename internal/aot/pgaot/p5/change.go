package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

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
