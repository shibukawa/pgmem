package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_on_dsm_detach(m *base.Module, l0 int32, l1 int32, l2 int64) {
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_on_dsm_detach[0]))
	v8 = F_MemoryContextAlloc(m, v6, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v8 + int32(16)
		return
	}
}
func F_on_pb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
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
	var v25 int64
	_ = v25
	v4 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.F64_le(v6, v8) == int32(0) {
		v25 = v4
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		if base.F64_le(v12, v6) == int32(0) {
			v25 = v4
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			if base.F64_le(v16, v17) == int32(0) {
				v25 = v4
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
				v25 = base.I64_extend_i32_u(base.F64_le(v21, v16))
			}
		}
	}
	return v25
}
