package p0

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pg_logical_slot_peek_binary_changes(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_logical_slot_get_changes_guts(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
