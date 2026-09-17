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
func F_make_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v28 float64
	_ = v28
	var v35 int32
	_ = v35
	var v43 int64
	_ = v43
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v13))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(v16))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807))) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
				F_errmsg(m, int32(_a_F_make_time_0), v9)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_make_time_1), int32(1654), int32(_a_F_make_time_2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
		v28 = base.F64_nearest(base.F64_mul(v12, float64(1e+06)))
		if base.F64_lt(v28, float64(0))|base.F64_gt(v28, float64(6e+07)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
					F_errmsg(m, int32(_a_F_make_time_0), v9)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_make_time_1), int32(1654), int32(_a_F_make_time_2))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
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
			v35 = int32(60)
			v43 = base.I64_trunc_sat_f64_s(v28) + base.I64_extend_i32_u((v13*v35+v16)*v35)*int64(1000000)
			if v43 < int64(86400000001) {
				v68 = F_Int64GetDatum(m, v43)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v68
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
						F_errmsg(m, int32(_a_F_make_time_0), v9)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_make_time_1), int32(1654), int32(_a_F_make_time_2))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
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
	}
}
func F_time_mi_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v8 != int32(-2147483648) {
		if v8 == int32(2147483647) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			if base.B2i32(v20 != int32(2147483647))|base.B2i32(v23 != int64(9223372036854775807)) != 0 {
				v46 = v23
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v6-v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				v55 = F_Int64GetDatum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v55
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_time_mi_interval_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_time_mi_interval_1), int32(2149), int32(_a_F_time_mi_interval_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v46 = v13
			v48 = int64(86400000000)
			v49 = base.I64_rem_s(v6-v46, v48)
			if v49 < int64(0) {
				v54 = v49 + v48
			} else {
				v54 = v49
			}
			v55 = F_Int64GetDatum(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return v55
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v15 != int32(-2147483648) {
			v46 = v14
			v48 = int64(86400000000)
			v49 = base.I64_rem_s(v6-v46, v48)
			if v49 < int64(0) {
				v54 = v49 + v48
			} else {
				v54 = v49
			}
			v55 = F_Int64GetDatum(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return v55
			}
		} else {
			if v14 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_time_mi_interval_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_time_mi_interval_1), int32(2149), int32(_a_F_time_mi_interval_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
				v46 = v14
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v6-v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				v55 = F_Int64GetDatum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v55
				}
			}
		}
	}
}
func F_time_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v12 = base.I64_div_s(v10, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+140)) = uint32(v12)
	v17 = v10 + base.I64_extend32_s(v12)*int64(-3600000000)
	v19 = base.I64_div_s(v17, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+136)) = uint32(v19)
	v24 = base.I64_extend32_s(v19)*int64(-60000000) + v17
	v26 = base.I64_div_s(v24, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v7)+132)) = uint32(v26)
	v29 = v7 + int32(132)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v39 = int32(2)
	v40 = F_pg_ultostr_zeropad(m, v7, v38, v39)
	mBase = m.M
	v41 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v41)
	v43 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v47 = F_pg_ultostr_zeropad(m, v40+v43, v45, v39)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v41)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v54 = F_AppendSeconds(m, v47+v43, v52, base.I32_wrap_i64(v26*int64(4293967296)+v24), v43)
	mBase = m.M
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v57)
	v59 = F_pstrdup(m, v7)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(176)
		return v59
	}
}
