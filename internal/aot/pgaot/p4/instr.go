package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrJitAgg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3 + v4
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v7 + v8
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v11 + v12
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v15 + v16
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v19 + v20
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v23 + v24
	return
}
