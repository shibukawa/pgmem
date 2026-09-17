package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_info_copy(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	return
}
