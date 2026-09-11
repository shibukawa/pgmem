package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CacheInvalidateSmgr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v8 = int32(253)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)) = uint16(v10)
	v13 = int32(base.Ui32(v10) >> (uint(v5) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v17
	F_SendSharedInvalidMessages(m, v6, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
