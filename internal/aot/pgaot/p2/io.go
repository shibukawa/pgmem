package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_io_combine_limit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_assign_io_combine_limit[0]))
	if v5 < l0 {
		v7 = v5
	} else {
		v7 = l0
	}
	*(*int32)(unsafe.Add(mBase, _c_F_assign_io_combine_limit[1])) = v7
	return
}
