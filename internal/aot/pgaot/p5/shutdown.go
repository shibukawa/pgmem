package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownAuxiliaryProcess(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	F_LWLockReleaseAll(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_ConditionVariableCancelSleep(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, _consts[97]))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
			return
		}
	}
}
