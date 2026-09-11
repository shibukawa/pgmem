package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetReindexState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[0]))
	if l0 <= v3 {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[1])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[2])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[3])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[0])) = v6
	} else {
	}
	return
}
