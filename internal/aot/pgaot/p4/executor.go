package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[500]))
	if v4 != 0 {
		m.T0[v4].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ExecutorEnd(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExecutorRewind(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = int32(4562096)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_ExecReScan(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v4
		return
	}
}
