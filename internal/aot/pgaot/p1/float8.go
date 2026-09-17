package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v10 = base.F64_abs(base.F64_sub(v6, v8))
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(v10, v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float8_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v23 int64
	_ = v23
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 float64
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v10, float64(0)) != 0 {
				v23 = int64(-9223372036854775807 - 1)
			} else {
				v23 = int64(9223372036854775807)
			}
			v39 = v23
			v40 = F_Int64GetDatum(m, v39)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(32)
				return v40
			}
		} else {
			if base.F64_lt(v10, float64(-2.108668032e+11))|base.F64_ge(v10, float64(9.224318016e+12)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v7))) = v10
						F_errmsg(m, int32(_a_F_float8_timestamptz_0), v7)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(753), int32(_a_F_float8_timestamptz_2))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v34 = base.I64_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(base.F64_add(v10, float64(-9.466848e+08)), float64(1e+06))))
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v34+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v88 = *(*float64)(unsafe.Add(mBase, uint32(v87)))
							*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = v88
							F_errmsg(m, int32(_a_F_float8_timestamptz_0), v7+int32(16))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(766), int32(_a_F_float8_timestamptz_2))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v39 = v34
					v40 = F_Int64GetDatum(m, v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v40
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_float8_timestamptz_3), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_timestamptz_1), int32(735), int32(_a_F_float8_timestamptz_2))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
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
