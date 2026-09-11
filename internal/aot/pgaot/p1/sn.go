package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sn_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v8 <= v5+int32(1) {
		F_appendStringInfoChar(m, v4, int32(93))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v19 = int32(93)
		*(*uint8)(unsafe.Add(mBase, uint32(v17+v5))) = uint8(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
		v24 = v22 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v26+v24))) = uint8(v28)
		return v28
	}
}
