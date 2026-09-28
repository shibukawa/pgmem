package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_pg_copy_logical_replication_slot_a(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_copy_replication_slot(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pg_logical_slot_get_binary_changes(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = int32(1)
	F_pg_logical_slot_get_changes_guts(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_logical_slot_get_changes(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_logical_slot_get_changes_guts(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
