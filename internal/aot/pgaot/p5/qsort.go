package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_qsort_interruptible_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v8 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l1, l2, l4)
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v8 < int32(0) {
				if v12 < int32(0) {
					v31 = l1
					return v31
				} else {
					v18 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l4)
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 < int32(0) {
							v22 = l2
						} else {
							v22 = l0
						}
						return v22
					}
				}
			} else {
				if int32(0) < v12 {
					v31 = l1
					return v31
				} else {
					v26 = m.T0[l3].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, l4)
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 < int32(0) {
							v30 = l0
						} else {
							v30 = l2
						}
						v31 = v30
						return v31
					}
				}
			}
		}
	}
}
