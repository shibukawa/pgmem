package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_bit_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v5 != 0 {
		v27 = int32(2)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = int32(1)
		v14 = base.I32_rem_u_s(v6, v7<<(uint(v8)%32)&int32(-8)+int32(-64))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v14)>>(uint(int32(3))%32)))+8)))
		v27 = int32(base.Ui32(v18)>>(uint(v14&int32(7))%32)) & v8 << (uint(v8) % 32)
	}
	return v27
}
