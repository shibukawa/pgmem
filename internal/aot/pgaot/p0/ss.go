package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SS_process_sublinks(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = l2
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	v13 = F_process_sublinks_mutator(m, l1, v7+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v13
	}
}
