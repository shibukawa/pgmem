package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCommandTagNameAndLen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v4 = l0 << (uint(int32(3)) % 32)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[832]))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+uint32(_consts[465])))
	return v11
}
