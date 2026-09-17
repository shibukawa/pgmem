package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_mul_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v17 float64
	_ = v17
	var v25 float64
	_ = v25
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v5 = base.F64_mul(l1, base.F64_convert_i64_s(l0))
	v7 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v5), v7)&base.F64_ne(base.F64_abs(l1), v7) == int32(0) {
		v17 = float64(0)
		if base.B2i32(base.B2i32(l0 == int64(0))|base.F64_ne(v5, v17) == int32(0))&base.F64_ne(l1, v17) != 0 {
			F_float_underflow_error(m)
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v25 = base.F64_nearest(v5)
			v33 = int32(0)
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v25)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v25, float64(-9.223372036854776e+18)) == v33)|base.B2i32(base.F64_lt(v25, float64(9.223372036854776e+18)) == v33) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_cash_mul_float8_0), int32(0))
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cash_mul_float8_1), int32(124), int32(_a_F_cash_mul_float8_2))
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				return base.I64_trunc_sat_f64_s(v25)
			}
		}
	} else {
		F_float_overflow_error(m)
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cash_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 != v5)
}
