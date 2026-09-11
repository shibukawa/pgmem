package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecReScanHash(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+52))
	if v3 == int32(0) {
		F_ExecReScan(m, v2)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
