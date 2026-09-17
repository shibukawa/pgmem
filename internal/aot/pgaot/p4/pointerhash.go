package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pointerhash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13997(m, l0, l1, l2, int32(_a_F_pointerhash_create_0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
