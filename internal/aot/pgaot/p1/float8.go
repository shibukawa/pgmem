package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 float64
	_ = v9
	var v10 int64
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v20 int32
	_ = v20
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v45 float64
	_ = v45
	var v53 int32
	_ = v53
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = base.F64_reinterpret_i64(v8)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = base.F64_reinterpret_i64(v10)
	v12 = base.F64_sub(v9, v11)
	v13 = base.F64_abs(v12)
	v14 = math.Float64frombits(uint64(0x7ff0000000000000))
	v20 = int32(0)
	if base.B2i32(base.F64_ne(v13, v14)|base.F64_eq(base.F64_abs(v9), v14) == v20)&base.F64_ne(base.F64_abs(v11), v14) == v20 {
		if base.Ui64(base.I64_reinterpret_f64(v13)) < base.Ui64(int64(9218868437227405313)) {
			v45 = v12
		} else {
			v31 = int64(9223372036854775807)
			v32 = v10 & v31
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v8&v31) {
				if base.Ui64(v32) <= base.Ui64(int64(9218868437227405312)) {
					v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				} else {
					v45 = float64(0)
				}
			} else {
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v32) {
					v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				} else {
					v45 = float64(0)
				}
			}
		}
		return base.I64_reinterpret_f64(v45) & int64(9223372036854775807)
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float8_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v14 float64
	_ = v14
	var v22 int64
	_ = v22
	var v33 int64
	_ = v33
	var v38 int64
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v9&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v14 = base.F64_reinterpret_i64(v9)
		if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v14, float64(0)) != 0 {
				v22 = int64(-9223372036854775807 - 1)
			} else {
				v22 = int64(9223372036854775807)
			}
			v38 = v22
			m.G0 = v7 + int32(32)
			return v38
		} else {
			if base.F64_lt(v14, float64(-2.108668032e+11))|base.F64_ge(v14, float64(9.224318016e+12)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(_a_F_float8_timestamptz_0), v7)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(749), int32(_a_F_float8_timestamptz_2))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
				v33 = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_add(v14, float64(-9.466848e+08)), float64(1e+06))))
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v33+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int64(0)
						} else {
							v84 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = v84
							F_errmsg(m, int32(_a_F_float8_timestamptz_0), v7+int32(16))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(762), int32(_a_F_float8_timestamptz_2))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
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
					v38 = v33
					m.G0 = v7 + int32(32)
					return v38
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_float8_timestamptz_3), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(731), int32(_a_F_float8_timestamptz_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
