package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_collation_actual_version(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	switch l0 - int32(98) {
	case 0:
		v6 = F_get_collation_actual_version_builtin(m, l1)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	case 1:
		v12 = F_pg_strcasecmp(m, int32(_a_F_get_collation_actual_version_0), l1)
		if v12 == int32(0) {
		} else {
			v17 = F_pg_strncasecmp(m, int32(_a_F_get_collation_actual_version_1), l1, int32(2))
			if v17 == int32(0) {
			} else {
				v21 = F_pg_strcasecmp(m, int32(_a_F_get_collation_actual_version_2), l1)
			}
		}
		v23 = int32(0)
		return v23
	default:
		v23 = int32(0)
		return v23
	}
}
func F_get_collation_name(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13917(m, l0, int32(16))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
