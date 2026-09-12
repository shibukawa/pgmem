package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_gin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = base.I32_div_s(l1-v6, int32(3))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5+v9)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v11))))
	if v13 == int32(1) {
		v16 = int32(2)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v19 != 0 {
			v20 = v16
		} else {
			v20 = int32(1)
		}
		if l2 != 0 {
			v21 = v16
		} else {
			v21 = v20
		}
		v22 = v21
	} else {
		v22 = v13
	}
	return base.I32_extend8_s(v22)
}
