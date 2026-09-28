package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_div_int64(m *base.Module, l0 int64, l1 int64) int64 {
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v4 = int64(1)
	v5 = l1 + v4
	if base.Ui64(v5) <= base.Ui64(v4) {
		if base.I32_wrap_i64(v5) == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(33816706))
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_cash_div_int64_0), int32(0))
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_cash_div_int64_1), int32(162), int32(_a_F_cash_div_int64_2))
						v28 = m.ExcPending
						if v28 != 0 {
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
			if l0 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_cash_div_int64_3), int32(0))
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cash_div_int64_1), int32(175), int32(_a_F_cash_div_int64_2))
							v51 = m.ExcPending
							if v51 != 0 {
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
				return int64(0) - l0
			}
		}
	} else {
		v34 = base.I64_div_s(l0, l1)
		return v34
	}
}
func F_cash_mul_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v5 float64
	_ = v5
	var v8 float64
	_ = v8
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v22 float64
	_ = v22
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v8 = base.F64_mul(l1, base.F64_convert_i64_s(l0))
	if base.F64_eq(base.F64_abs(l1), v5)|base.F64_ne(base.F64_abs(v8), v5) == int32(0) {
		v16 = F_float_overflow_error_ext(m, int32(0))
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v31 = v16
			v32 = base.F64_nearest(v31)
			v40 = int32(0)
			if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v32, float64(-9.223372036854776e+18)) == v40) == v40)&base.F64_lt(v32, float64(9.223372036854776e+18)) == v40 {
				F_errstart_cold(m, int32(21), int32(0))
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_cash_mul_float8_0), int32(0))
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cash_mul_float8_1), int32(125), int32(_a_F_cash_mul_float8_2))
							v65 = m.ExcPending
							if v65 != 0 {
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
				return base.I64_trunc_sat_f64_s(v32)
			}
		}
	} else {
		v22 = float64(0)
		if base.B2i32(l0 == int64(0))|base.F64_ne(v8, v22)|base.F64_eq(l1, v22) != 0 {
			v31 = v8
			v32 = base.F64_nearest(v31)
			v40 = int32(0)
			if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v32, float64(-9.223372036854776e+18)) == v40) == v40)&base.F64_lt(v32, float64(9.223372036854776e+18)) == v40 {
				F_errstart_cold(m, int32(21), int32(0))
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_cash_mul_float8_0), int32(0))
						v60 = m.ExcPending
						if v60 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cash_mul_float8_1), int32(125), int32(_a_F_cash_mul_float8_2))
							v65 = m.ExcPending
							if v65 != 0 {
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
				return base.I64_trunc_sat_f64_s(v32)
			}
		} else {
			v29 = F_float_underflow_error_ext(m, int32(0))
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = v29
				v32 = base.F64_nearest(v31)
				v40 = int32(0)
				if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v32, float64(-9.223372036854776e+18)) == v40) == v40)&base.F64_lt(v32, float64(9.223372036854776e+18)) == v40 {
					F_errstart_cold(m, int32(21), int32(0))
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_cash_mul_float8_0), int32(0))
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_cash_mul_float8_1), int32(125), int32(_a_F_cash_mul_float8_2))
								v65 = m.ExcPending
								if v65 != 0 {
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
					return base.I64_trunc_sat_f64_s(v32)
				}
			}
		}
	}
}
func F_cash_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
