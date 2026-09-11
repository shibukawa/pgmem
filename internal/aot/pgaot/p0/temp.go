package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetTempTableNamespace(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(2615)
		F_performDeletion(m, v5+int32(4), int32(1), int32(29))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_isTempToastNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	return base.B2i32(v4 != int32(0)) & base.B2i32(l0 == v4)
}
