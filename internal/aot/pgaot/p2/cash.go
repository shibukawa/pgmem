package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_div_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v4 float64
	_ = v4
	var v15 float64
	_ = v15
	var v30 float64
	_ = v30
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v4 = base.F64_convert_i64_s(l0)
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(l1, float64(0)) == int32(0) {
		v15 = base.F64_div(v4, l1)
		if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_float_overflow_error(m)
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			if base.B2i32(base.B2i32(l0 == int64(0))|base.F64_ne(v15, float64(0)) == int32(0))&base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_underflow_error(m)
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v30 = base.F64_nearest(v15)
				v38 = int32(0)
				if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v30)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v30, float64(-9.223372036854776e+18)) == v38)|base.B2i32(base.F64_lt(v30, float64(9.223372036854776e+18)) == v38) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_cash_div_float8_0), int32(0))
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_cash_div_float8_1), int32(137), int32(_a_F_cash_div_float8_2))
								v71 = m.ExcPending
								if v71 != 0 {
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
					return base.I64_trunc_sat_f64_s(v30)
				}
			}
		}
	} else {
		F_float_zero_divide_error(m)
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
func F_cash_div_flt4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 float32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_cash_div_float8(m, v3, base.F64_promote_f32(v4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_cash_lt(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v3 < v5)
}
func F_cash_mi(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13881(m, l0, int32(_a_F_cash_mi_0), int32(111), int32(_a_F_cash_mi_1), int32(_a_F_cash_mi_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_cash_mul_int2(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13882(m, l0, int32(_a_F_cash_mul_int2_0), int32(150), int32(_a_F_cash_mul_int2_1), int32(_a_F_cash_mul_int2_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_cash_mul_int8(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13884(m, l0, int32(_a_F_cash_mul_int8_0), int32(150), int32(_a_F_cash_mul_int8_1), int32(_a_F_cash_mul_int8_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_cash_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v19 = m.G0
	v21 = v19 - int32(416)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v25 = F_PGLC_localeconv(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+41)))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
		v40 = int32(_a_F_cash_out_0)
		v41 = int32(46)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
		if v43 == int32(0) {
			v52 = v40
			v53 = v41
		} else {
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
			if v46 != 0 {
				v52 = v40
				v53 = v41
			} else {
				if v43 == int32(44) {
					v51 = int32(_a_F_cash_out_1)
				} else {
					v51 = int32(_a_F_cash_out_0)
				}
				v52 = v51
				v53 = v43
			}
		}
		if base.Ui32(int32(10)) < base.Ui32(v29) {
			v55 = int32(2)
		} else {
			v55 = v29
		}
		if base.Ui32((v33-int32(7))&int32(255)) < base.Ui32(int32(250)) {
			v57 = int32(3)
		} else {
			v57 = v33
		}
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
		if v24 < int64(0) {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
			if v69 != 0 {
				v70 = v67
			} else {
				v70 = int32(_a_F_cash_out_2)
			}
			v76 = int32(44)
			v77 = int32(45)
			v78 = int32(47)
			v79 = v70
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
			v76 = int32(42)
			v77 = int32(43)
			v78 = int32(46)
			v79 = v74
		}
		if v59 != 0 {
			v83 = v58
		} else {
			v83 = int32(_a_F_cash_out_3)
		}
		if v61 != 0 {
			v84 = v60
		} else {
			v84 = v52
		}
		v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v78))))
		v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v76))))
		v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v77))))
		v91 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v21)+415)) = uint8(v91)
		v94 = v24 >> (uint(int64(63)) % 64)
		v99 = base.I32_extend8_s(v55)
		v101 = v21 + int32(415)
		v116 = v24 ^ v94 - v94
		for {
			v117 = int32(0)
			if base.B2i32(v55 == v117)|v99 == v117 {
				v123 = v101 - int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v53)
				v133 = v123
			} else {
				if int32(0) <= v99 {
					v133 = v101
				} else {
					v127 = base.I32_rem_s(v99, base.I32_extend8_s(v57))
					if v127 != 0 {
						v133 = v101
					} else {
						v128 = F_strlen(m, v84)
						mBase = m.M
						v129 = v101 - v128
						if v128 == int32(0) {
							v133 = v129
						} else {
							base.MemoryCopy(m, v129, v84, v128)
							v133 = v129
						}
					}
				}
			}
			v135 = int32(1)
			v136 = v133 - v135
			v137 = int64(10)
			v138 = base.I64_div_u_s(v116, v137)
			v144 = base.I32_wrap_i64(v116-v138*v137) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v144)
			v147 = v99 - v135
			if base.B2i32(int32(0) <= v147)|base.B2i32(base.Ui64(int64(9)) < base.Ui64(v116)) != 0 {
				v99 = v147
				v101 = v136
				v116 = v138
				continue
			} else {
				break
			}
			break
		}
		switch v86 {
		case 0:
			if v90 == int32(1) {
				v157 = int32(_a_F_cash_out_4)
			} else {
				v157 = int32(_a_F_cash_out_5)
			}
			if v88 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v157
				*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v83
				v164 = F_psprintf(m, int32(_a_F_cash_out_6), v21+int32(80))
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					v314 = v164
					m.G0 = v21 + int32(416)
					return v314
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v157
				*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v136
				v172 = F_psprintf(m, int32(_a_F_cash_out_6), v21-int32(-64))
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v314 = v172
					m.G0 = v21 + int32(416)
					return v314
				}
			}
		default:
			if v90 == int32(1) {
				v178 = int32(_a_F_cash_out_4)
			} else {
				v178 = int32(_a_F_cash_out_5)
			}
			if v90 == int32(2) {
				v183 = int32(_a_F_cash_out_4)
			} else {
				v183 = int32(_a_F_cash_out_5)
			}
			if v88 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v178
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v183
				*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v79
				v192 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(32))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return int32(0)
				} else {
					v314 = v192
					m.G0 = v21 + int32(416)
					return v314
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v178
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v183
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v79
				v200 = F_psprintf(m, int32(_a_F_cash_out_7), v21)
				mBase = m.M
				v201 = m.ExcPending
				if v201 != 0 {
					return int32(0)
				} else {
					v314 = v200
					m.G0 = v21 + int32(416)
					return v314
				}
			}
		case 2:
			if v90 == int32(2) {
				v206 = int32(_a_F_cash_out_4)
			} else {
				v206 = int32(_a_F_cash_out_5)
			}
			if v90 == int32(1) {
				v211 = int32(_a_F_cash_out_4)
			} else {
				v211 = int32(_a_F_cash_out_5)
			}
			if v88 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v206
				*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v211
				*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v83
				v220 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(192))
				mBase = m.M
				v221 = m.ExcPending
				if v221 != 0 {
					return int32(0)
				} else {
					v314 = v220
					m.G0 = v21 + int32(416)
					return v314
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v206
				*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v211
				*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v136
				v230 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(160))
				mBase = m.M
				v231 = m.ExcPending
				if v231 != 0 {
					return int32(0)
				} else {
					v314 = v230
					m.G0 = v21 + int32(416)
					return v314
				}
			}
		case 3:
			if v88 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+272)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+264)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = v79
				if v90 == int32(1) {
					v239 = int32(_a_F_cash_out_4)
				} else {
					v239 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+268)) = v239
				if v90 == int32(2) {
					v245 = int32(_a_F_cash_out_4)
				} else {
					v245 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+260)) = v245
				v250 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(256))
				mBase = m.M
				v251 = m.ExcPending
				if v251 != 0 {
					return int32(0)
				} else {
					v314 = v250
					m.G0 = v21 + int32(416)
					return v314
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v136
				if v90 == int32(2) {
					v259 = int32(_a_F_cash_out_4)
				} else {
					v259 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v259
				if v90 == int32(1) {
					v265 = int32(_a_F_cash_out_4)
				} else {
					v265 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = v265
				v270 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(224))
				mBase = m.M
				v271 = m.ExcPending
				if v271 != 0 {
					return int32(0)
				} else {
					v314 = v270
					m.G0 = v21 + int32(416)
					return v314
				}
			}
		case 4:
			if v88 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v136
				*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v83
				if v90 == int32(1) {
					v279 = int32(_a_F_cash_out_4)
				} else {
					v279 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+140)) = v279
				if v90 == int32(2) {
					v285 = int32(_a_F_cash_out_4)
				} else {
					v285 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v285
				v290 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(128))
				mBase = m.M
				v291 = m.ExcPending
				if v291 != 0 {
					return int32(0)
				} else {
					v314 = v290
					m.G0 = v21 + int32(416)
					return v314
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v136
				if v90 == int32(2) {
					v299 = int32(_a_F_cash_out_4)
				} else {
					v299 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v299
				if v90 == int32(1) {
					v305 = int32(_a_F_cash_out_4)
				} else {
					v305 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v305
				v310 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(96))
				mBase = m.M
				v311 = m.ExcPending
				if v311 != 0 {
					return int32(0)
				} else {
					v314 = v310
					m.G0 = v21 + int32(416)
					return v314
				}
			}
		}
	}
}
func F_cash_pl(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13885(m, l0, int32(_a_F_cash_pl_0), int32(98), int32(_a_F_cash_pl_1), int32(_a_F_cash_pl_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
