package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v9 = int64(1000000)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v12 = base.I64_extend_i32_s(v7)*v9 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v19 = base.I64_extend_i32_s(v14)*v9 + v18
	if v19 < v12 {
		return int32(0)
	} else {
		if v12 < v19 {
			return int32(1)
		} else {
			return base.B2i32(v7 <= v14)
		}
	}
}
func F_timetz_ne(m *base.Module, l0 int32) int32 {
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
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v12 = int64(1000000)
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	return base.B2i32(v6 != v8) | base.B2i32(v10+base.I64_extend_i32_s(v6)*v12 != v15+base.I64_extend_i32_s(v8)*v12)
}
