package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_abort(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+73)))
	if v3 == int32(1) {
		F_errdetail(m, int32(109076), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
