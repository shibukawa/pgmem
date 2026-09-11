package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlabGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	v3 = l0 - int32(8)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v3-base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(34))%64)))&int32(1073741822))))
	return v11
}
