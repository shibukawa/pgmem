package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_xact_id_if_assigned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_xact_id_if_assigned[0]))
	if base.I32_wrap_i64(v4) != 0 {
		v9 = v4
	} else {
		v6 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v6)
		v9 = int64(0)
	}
	return v9
}
