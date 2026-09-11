package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_system_nextsampletuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)))
	v7 = v5 + int32(1)
	if base.Ui32(v7&int32(65535)) <= base.Ui32(l2) {
		v12 = v7
	} else {
		v12 = int32(0)
	}
	*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)) = uint16(v12)
	return v12 & int32(65535)
}
func F_system_user(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1189]))
	if v4 != 0 {
		v5 = F_cstring_to_text(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		return int32(0)
	}
}
