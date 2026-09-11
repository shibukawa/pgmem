package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_recordDependencyOnOwner(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v24 int32
	_ = v24
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(1260)
	F_recordSharedDependencyOn(m, v7+int32(20), v7+int32(8), int32(111))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		m.G0 = v7 + int32(32)
		return
	}
}
