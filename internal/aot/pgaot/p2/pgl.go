package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_pclose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1298]))
	if v4 != 0 {
		v5 = m.T0[v4].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(52)
		return int32(-1)
	}
}
func F_pgl_recv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1299]))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
