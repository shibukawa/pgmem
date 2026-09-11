package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_info_copy(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	v6 = l0 + int32(16)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
	return
}
