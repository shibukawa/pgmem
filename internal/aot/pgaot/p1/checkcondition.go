package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_QueryOperand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = base.I32_div_s(l1-v7-int32(8), int32(12))
	v15 = v6 + v12*int32(32776)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != int32(1) {
		v41 = int32(0)
	} else {
		if l2 == int32(0) {
			v41 = int32(1)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v24 = v15 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
			v27 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v28 != v27 {
				v41 = v27
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24 + (int32(16384)-v32)<<(uint(v34)%32)
				v41 = v34
			}
		}
	}
	return v41
}
