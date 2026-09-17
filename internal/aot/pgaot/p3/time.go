package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_time_overflows(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v12 int32
	_ = v12
	v12 = int32(60)
	return base.B2i32(base.Ui32(int32(24)) < base.Ui32(l0)) | base.B2i32(base.Ui32(int32(59)) < base.Ui32(l1)) | (base.B2i32(base.Ui32(v12) < base.Ui32(l2)) | base.B2i32(base.Ui32(int32(_a_F_time_overflows_0)) < base.Ui32(l3))) | base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(l3)+base.I64_extend_i32_u((l0*v12+l1)*v12+l2)*int64(1000000)))
}
