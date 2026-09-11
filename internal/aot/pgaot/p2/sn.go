package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SN_set_current(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = F_replace_s(m, l0, v4, v5, l1, l2, v4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
		return v7
	}
}
func F_sn_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v4 == int32(1) {
		v7 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v7)
		return v7
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if l2 == int32(1) {
			F_escape_json(m, v11, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			F_appendStringInfoString(m, v11, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
