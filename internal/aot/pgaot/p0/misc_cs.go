package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cstring_to_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	v5 = F_strlen(m, l0)
	mBase = m.M
	v7 = v5 + int32(4)
	v8 = F_palloc(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v7 << (uint(int32(2)) % 32)
		if v5 != 0 {
			v17 = F__emscripten_memcpy_bulkmem(m, v8+int32(4), l0, v5)
			mBase = m.M
		} else {
		}
		return v8
	}
}
