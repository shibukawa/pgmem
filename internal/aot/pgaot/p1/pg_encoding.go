package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_max_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_max_length[0])))
		v8 = v6
	} else {
		v8 = int32(1)
	}
	return v8
}
