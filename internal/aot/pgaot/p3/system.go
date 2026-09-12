package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsSystemClass(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	v4 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(12000)) {
		v18 = v4
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		if v7 == int32(99) {
			v18 = v4
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[136]))
			v18 = base.B2i32(v12 != int32(0)) & base.B2i32(v7 == v12)
		}
	}
	return v18
}
func F_system_initsamplescan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_palloc0(m, int32(24))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v4
		return
	}
}
