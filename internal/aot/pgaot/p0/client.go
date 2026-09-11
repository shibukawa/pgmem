package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ClientCheckTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ClientCheckTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_ClientCheckTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ClientCheckTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
