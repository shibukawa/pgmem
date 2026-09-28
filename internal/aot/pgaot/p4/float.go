package p4

import base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"

func F_float_overflow_error(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(50331778))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_float_overflow_error_0), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_float_overflow_error_1), int32(107), int32(_a_F_float_overflow_error_2))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_float_overflow_error_ext(m *base.Module, l0 int32) float64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = F_errsave_start(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return float64(0)
	} else {
		if v2 != 0 {
			F_errcode(m, int32(50331778))
			v8 = m.ExcPending
			if v8 != 0 {
				return float64(0)
			} else {
				F_errmsg(m, int32(_a_F_float_overflow_error_ext_0), int32(0))
				v12 = m.ExcPending
				if v12 != 0 {
					return float64(0)
				} else {
					F_errsave_finish(m, l0, int32(_a_F_float_overflow_error_ext_1), int32(131), int32(_a_F_float_overflow_error_ext_2))
					v17 = m.ExcPending
					if v17 != 0 {
						return float64(0)
					} else {
						return float64(0)
					}
				}
			}
		} else {
			return float64(0)
		}
	}
}
func F_float_zero_divide_error(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(33816706))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_float_zero_divide_error_0), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_float_zero_divide_error_1), int32(123), int32(_a_F_float_zero_divide_error_2))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
