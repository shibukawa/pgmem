package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpecialFURL(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v3)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v6 - v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11 - v12
	return
}
func F_SpecialVerVersion(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = v4 - v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9 - v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v2
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v2
	return
}
