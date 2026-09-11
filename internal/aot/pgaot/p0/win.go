package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetCurrentPosition(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)+176))
	return v3
}
func F_WinGetPartitionLocalMemory(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+360))
		v9 = F_MemoryContextAllocZero(m, v8, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
			v14 = v9
			return v14
		}
	} else {
		v14 = v4
		return v14
	}
}
