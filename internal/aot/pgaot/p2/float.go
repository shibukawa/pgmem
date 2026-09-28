package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_float_compare_desc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_gt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F32_lt(v6, v7)
	}
	return v10
}
func F_float_underflow_error_ext(m *base.Module, l0 int32) float64 {
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
				F_errmsg(m, int32(_a_F_float_underflow_error_ext_0), int32(0))
				v12 = m.ExcPending
				if v12 != 0 {
					return float64(0)
				} else {
					F_errsave_finish(m, l0, int32(_a_F_float_underflow_error_ext_1), int32(139), int32(_a_F_float_underflow_error_ext_2))
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
func F_float_zero_divide_error_ext(m *base.Module, l0 int32) float64 {
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
			F_errcode(m, int32(33816706))
			v8 = m.ExcPending
			if v8 != 0 {
				return float64(0)
			} else {
				F_errmsg(m, int32(_a_F_float_zero_divide_error_ext_0), int32(0))
				v12 = m.ExcPending
				if v12 != 0 {
					return float64(0)
				} else {
					F_errsave_finish(m, l0, int32(_a_F_float_zero_divide_error_ext_1), int32(147), int32(_a_F_float_zero_divide_error_ext_2))
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
