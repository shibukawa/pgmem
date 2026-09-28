package p5

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pg_try_advisory_lock_shared_int8(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14351(m, l0, int32(1), int32(5))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
