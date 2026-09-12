package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RunNamespaceSearchHook(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v2 = l1
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v8)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[229]))
	m.T0[v17].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(3), int32(2615), l0, int32(0), v6+int32(14))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		m.G0 = v6 + int32(16)
		return v22
	}
}
