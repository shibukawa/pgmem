package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cstring_to_text_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = l1 + int32(4)
	v7 = F_palloc(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v6 << (uint(int32(2)) % 32)
		if l1 != 0 {
			base.MemoryCopy(m, v7+int32(4), l0, l1)
		} else {
		}
		return v7
	}
}
