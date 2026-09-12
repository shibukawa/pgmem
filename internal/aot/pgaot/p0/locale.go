package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_locale_time(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[808])) = uint8(v4)
	return
}
func F_check_locale_messages(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v5 == int32(0) {
		return base.B2i32(l2 == int32(0))
	} else {
		v13 = F_check_locale(m, int32(5), v4, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v13
		}
	}
}
