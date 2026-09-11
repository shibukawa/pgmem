package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeDefElemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_palloc0(m, int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(93)
		return v8
	}
}
