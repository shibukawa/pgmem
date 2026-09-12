package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_dsa_create_in_place_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v5 = int32(0)
	v9 = F_create_internal(m, l0, l1, l2, v5, v5, int32(1048576), int32(134217728))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l3 != 0 {
			F_on_dsm_detach(m, l3, int32(1786), l0)
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v9
			}
		} else {
			return v9
		}
	}
}
