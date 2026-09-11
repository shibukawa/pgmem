package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetFuncArgCurrent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+64))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+l1<<(uint(int32(2))%32))))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v17 = m.T0[v16].(func(*base.Module, int32, int32, int32) int32)(m, v15, v7, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		return v17
	}
}
