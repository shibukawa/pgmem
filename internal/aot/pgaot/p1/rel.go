package p1

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_get_rel_tablespace(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13896(m, l0, int32(57))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
