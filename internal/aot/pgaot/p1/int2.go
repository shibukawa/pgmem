package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int2_sum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3 == int32(1) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v6 == int32(1) {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
			return int64(0)
		} else {
			v13 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
			return v13
		}
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v16 != 0 {
			v19 = v15
		} else {
			v17 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
			v19 = v17 + v15
		}
		return v19
	}
}
