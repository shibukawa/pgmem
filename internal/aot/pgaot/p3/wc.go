package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_wc_isalnum_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	if base.Ui32(int32(10)) <= base.Ui32(l0-int32(48)) {
		v8 = F_iswalpha(m, l0)
		v12 = base.B2i32(v8 != int32(0))
	} else {
		v12 = int32(1)
	}
	return base.B2i32(v12 != int32(0))
}
func F_wc_isalnum_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v20 int32
	_ = v20
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v20 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	} else {
		v20 = int32(0)
	}
	return v20
}
func F_wc_isdigit_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10))) != int32(0))
}
func F_wc_isxdigit_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v20 int32
	_ = v20
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v20 = base.B2i32(base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(6))) != int32(0))
	} else {
		v20 = int32(0)
	}
	return v20
}
