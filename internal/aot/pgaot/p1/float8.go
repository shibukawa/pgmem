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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v10 = base.F64_abs(base.F64_sub(v6, v8))
	if base.F64_ne(v10, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v23 = F_Float8GetDatum(m, v10)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	} else {
		if base.F64_eq(base.F64_abs(v6), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v23 = F_Float8GetDatum(m, v10)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		} else {
			if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v23 = F_Float8GetDatum(m, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
	var v32 float64
	_ = v32
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
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
			v44 = v23
			v45 = F_Int64GetDatum(m, v44)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(32)
				return v45
			}
		} else {
			if base.F64_lt(v10, float64(-2.108668032e+11)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v7))) = v10
						F_errmsg(m, int32(712899), v7)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495274), int32(753), int32(7522))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
				if base.F64_ge(v10, float64(9.224318016e+12)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v7))) = v10
							F_errmsg(m, int32(712899), v7)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495274), int32(753), int32(7522))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
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
					v32 = base.F64_nearest(base.F64_mul(base.F64_add(v10, float64(-9.466848e+08)), float64(1e+06)))
					if base.F64_lt(base.F64_abs(v32), float64(9.223372036854776e+18)) != 0 {
						v36 = base.I64_trunc_f64_s(v32)
						v38 = v36
					} else {
						v38 = int64(-9223372036854775807 - 1)
					}
					if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v38+int64(211813488000000000)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v93 = *(*float64)(unsafe.Add(mBase, uint32(v92)))
								*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = v93
								F_errmsg(m, int32(712899), v7+int32(16))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495274), int32(766), int32(7522))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
						v44 = v38
						v45 = F_Int64GetDatum(m, v44)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(32)
							return v45
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(526930), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495274), int32(735), int32(7522))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
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
