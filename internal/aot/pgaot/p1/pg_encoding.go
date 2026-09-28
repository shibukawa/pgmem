package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_max_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		v13 = int32(1)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_max_length[0])))
		v13 = v12
	}
	return v13
}
