package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_timezone_abbreviations(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, _c_F_assign_timezone_abbreviations[0])) = l1
		base.MemoryFill(m, int32(_a_F_assign_timezone_abbreviations_0), int32(0), int32(500))
	} else {
	}
	return
}
