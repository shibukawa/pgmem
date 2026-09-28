package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BackendStatusShmemAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemAttach[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemAttach[1]))
	v9 = F_mul_size(m, v4, v6+int32(38))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_BackendStatusShmemAttach[2])) = v9
		return
	}
}
