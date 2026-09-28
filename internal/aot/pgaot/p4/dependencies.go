package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_dependencies_array_start(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14255(m, l0, int32(_a_F_dependencies_array_start_0), int32(271), int32(_a_F_dependencies_array_start_1), int32(_a_F_dependencies_array_start_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
