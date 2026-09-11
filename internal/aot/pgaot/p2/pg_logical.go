package p2

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pg_logical_slot_peek_changes(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(0)
	F_pg_logical_slot_get_changes_guts(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
