package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_max_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[356])))
		v11 = v10
	} else {
		v11 = int32(1)
	}
	return v11
}
