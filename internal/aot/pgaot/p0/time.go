package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdjustTimeForTypmod(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	if base.Ui32(l1) <= base.Ui32(int32(6)) {
		v9 = l1 << (uint(int32(3)) % 32)
		v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_AdjustTimeForTypmod[0])))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_AdjustTimeForTypmod[1])))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if int64(0) <= v12 {
			v15 = v11 + v12
			v16 = base.I64_rem_s(v15, v10)
			v22 = v15 - v16
		} else {
			v18 = v11 - v12
			v19 = base.I64_rem_s(v18, v10)
			v22 = v19 - v18
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v22
	} else {
	}
	return
}
func F_make_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v30 float64
	_ = v30
	var v37 int64
	_ = v37
	var v46 int64
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = base.I32_wrap_i64(v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = base.I32_wrap_i64(v17)
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v14))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(v18))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807))) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
				F_errmsg(m, int32(_a_F_make_time_0), v10)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_make_time_1), int32(1694), int32(_a_F_make_time_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
		v30 = base.F64_nearest(base.F64_mul(v12, float64(1e+06)))
		if base.F64_lt(v30, float64(0))|base.F64_gt(v30, float64(6e+07)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
					F_errmsg(m, int32(_a_F_make_time_0), v10)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_make_time_1), int32(1694), int32(_a_F_make_time_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
			v37 = int64(60)
			v46 = base.I64_trunc_sat_f64_s(v30) + (v13*v37+v17)*v37&int64(4294967292)*int64(1000000)
			if v46 < int64(86400000001) {
				m.G0 = v10 + int32(16)
				return v46
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v18
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
						F_errmsg(m, int32(_a_F_make_time_0), v10)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_make_time_1), int32(1694), int32(_a_F_make_time_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
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
	}
}
func F_time_mi_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 != int32(-2147483648) {
		if v7 == int32(2147483647) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			if base.B2i32(v19 != int32(2147483647))|base.B2i32(v22 != int64(9223372036854775807)) != 0 {
				v45 = v22
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v5-v45, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				return v53
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_time_mi_interval_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_time_mi_interval_1), int32(2201), int32(_a_F_time_mi_interval_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v45 = v12
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v5-v45, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			return v53
		}
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v14 != int32(-2147483648) {
			v45 = v13
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v5-v45, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			return v53
		} else {
			if v13 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_time_mi_interval_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_time_mi_interval_1), int32(2201), int32(_a_F_time_mi_interval_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
				v45 = v13
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v5-v45, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				return v53
			}
		}
	}
}
func F_time_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = base.I64_div_s(v9, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+140)) = uint32(v11)
	v16 = v9 + base.I64_extend32_s(v11)*int64(-3600000000)
	v18 = base.I64_div_s(v16, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+136)) = uint32(v18)
	v23 = base.I64_extend32_s(v18)*int64(-60000000) + v16
	v25 = base.I64_div_s(v23, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+132)) = uint32(v25)
	v28 = v7 + int32(132)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v38 = int32(2)
	v39 = F_pg_ultostr_zeropad(m, v7, v37, v38)
	mBase = m.M
	v40 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v40)
	v42 = int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v46 = F_pg_ultostr_zeropad(m, v39+v42, v44, v38)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v53 = F_AppendSeconds(m, v46+v42, v51, base.I32_wrap_i64(v25*int64(4293967296)+v23), v42)
	mBase = m.M
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v56)
	v58 = F_pstrdup(m, v7)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(176)
		return base.I64_extend_i32_u(v58)
	}
}
