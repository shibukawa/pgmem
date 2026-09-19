package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_y_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
	if base.F64_gt(v8, v10) != 0 {
		v12 = int32(1)
	} else {
		v12 = int32(-1)
	}
	if base.F64_ne(v8, v10) != 0 {
		v15 = v12
	} else {
		v15 = int32(0)
	}
	return v15
}
func F_yy_fatal_error_3(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14026(m, l0, int32(_a_F_yy_fatal_error_3_0), int32(38), int32(_a_F_yy_fatal_error_3_1), int32(_a_F_yy_fatal_error_3_2))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
