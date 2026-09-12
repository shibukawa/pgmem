package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_timezone_abbreviations(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, _consts[421])) = l1
		v9 = F__emscripten_memset_bulkmem(m, int32(4427232), base.I32_extend8_s(int32(0)), int32(500))
		mBase = m.M
	} else {
	}
	return
}
