package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_restore_relation_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)) = uint8(v2)
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
	v14 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+26)) = uint16(v14)
	v17 = v6 + int32(8)
	v19 = F_stats_fill_fcinfo_from_arg_pairs(m, l0, v17, int32(_a_F_pg_restore_relation_stats_0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = F_relation_statistics_update(m, v17)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(80)
			return v19 & v23
		}
	}
}
