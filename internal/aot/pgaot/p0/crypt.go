package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__crypt_gensalt_sha512_rn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l4 != 0 {
		base.MemoryFill(m, l3, int32(0), l4)
	} else {
	}
	v8 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+2)) = uint8(v8)
	v10 = int32(_a_F__crypt_gensalt_sha512_rn_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v10)
	v12 = F__crypt_gensalt_sha(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
