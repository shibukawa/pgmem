package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCommandTagNameAndLen(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_GetCommandTagNameAndLen[0]))))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	return
}
