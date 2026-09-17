package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pg_try_advisory_lock_shared_int8(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13967(m, l0, int32(1), int32(5))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
