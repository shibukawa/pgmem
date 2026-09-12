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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v8 = F__emscripten_memset_bulkmem(m, l3, base.I32_extend8_s(int32(0)), l4)
	mBase = m.M
	v9 = int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)) = uint8(v9)
	v11 = int32(13860)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v11)
	v13 = F__crypt_gensalt_sha(m, l0, l1, l2, v8, l4)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		return v13
	}
}
