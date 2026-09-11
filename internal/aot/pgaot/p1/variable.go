package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_setup_parse_variable_parameters(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v6 = F_palloc(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = int32(493)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(494)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v6
		return
	}
}
