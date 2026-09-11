package p3

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F___time(m *base.Module) int64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	v2 = m.Env.Emscripten_date_now(m)
	v4 = base.F64_div(v2, float64(1000))
	if base.F64_lt(base.F64_abs(v4), float64(9.223372036854776e+18)) != 0 {
		v8 = base.I64_trunc_f64_s(v4)
		v10 = v8
	} else {
		v10 = int64(-9223372036854775807 - 1)
	}
	return v10
}
func F_time_overflows(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	v6 = int32(1)
	if base.Ui32(int32(24)) < base.Ui32(l0) {
		v28 = v6
	} else {
		if base.Ui32(int32(59)) < base.Ui32(l1) {
			v28 = v6
		} else {
			if base.Ui32(int32(60)) < base.Ui32(l2) {
				v28 = v6
			} else {
				if base.Ui32(int32(1000000)) < base.Ui32(l3) {
					v28 = v6
				} else {
					v16 = int32(60)
					v28 = base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(l3)+base.I64_extend_i32_u((l0*v16+l1)*v16+l2)*int64(1000000)))
				}
			}
		}
	}
	return v28
}
