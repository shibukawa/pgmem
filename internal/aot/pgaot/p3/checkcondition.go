package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_bit_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v4 != 0 {
		v29 = int32(2)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = int32(1)
		v15 = base.I32_rem_u_s(v7, v8<<(uint(v9)%32)&int32(-8)+int32(-64))
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v15)>>(uint(int32(3))%32)))+8)))
		v29 = (int32(0) - int32(base.Ui32(v19)>>(uint(v15&int32(7))%32))&v9) & int32(2)
	}
	return v29
}
