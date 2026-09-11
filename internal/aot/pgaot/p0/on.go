package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v8 = F_MemoryContextAlloc(m, v6, int32(12))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v8 + int32(8)
		return
	}
}
func F_on_pb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.F64_ge(v6, v8) == v2 {
		v23 = v2
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
		if base.F64_le(v12, v8) == int32(0) {
			v23 = v2
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
			if base.F64_le(v16, v17) == int32(0) {
				v23 = v2
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v5)+24))
				v23 = base.F64_le(v21, v16)
			}
		}
	}
	return v23
}
