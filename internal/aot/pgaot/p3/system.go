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
	if base.Ui32(l0) < base.Ui32(int32(_a_F_IsSystemClass_0)) {
		v18 = v4
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		if v7 == int32(99) {
			v18 = v4
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_IsSystemClass[0]))
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
func F_system_rows_initsamplescan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_palloc0(m, int32(40))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v4
		return
	}
}
func F_system_rows_nextsampletuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
	if v5 < v7 {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)))
		v11 = v9 + int32(1)
		if base.Ui32(v11&int32(_a_F_system_rows_nextsampletuple_0)) <= base.Ui32(l2) {
			v16 = v11
		} else {
			v16 = int32(0)
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)) = uint16(v16)
		v19 = v16
	} else {
		v19 = int32(0)
	}
	return v19 & int32(_a_F_system_rows_nextsampletuple_0)
}
func F_system_time_initsamplescan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_palloc0(m, int32(48))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v4
		return
	}
}
