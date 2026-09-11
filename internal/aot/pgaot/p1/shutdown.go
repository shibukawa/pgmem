package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownPostgres(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	F_AbortOutOfAnyTransaction(m)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_LockReleaseAll(m, int32(2), int32(1))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_shutdown_validator_library(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_shutdown_validator_library[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_shutdown_validator_library[1]))
		m.T0[v4].(func(*base.Module, int32))(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
