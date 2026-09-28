package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_cash_div_float8(m *base.Module, l0 int64, l1 float64) int64 {
	var v4 float64
	_ = v4
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = base.F64_convert_i64_s(l0)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)))|base.F64_ne(l1, float64(0)) == int32(0) {
		v16 = F_float_zero_divide_error_ext(m, int32(0))
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v39 = v16
			v40 = base.F64_nearest(v39)
			v48 = int32(0)
			if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v40, float64(-9.223372036854776e+18)) == v48) == v48)&base.F64_lt(v40, float64(9.223372036854776e+18)) == v48 {
				F_errstart_cold(m, int32(21), int32(0))
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50331778))
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_cash_div_float8_0), int32(0))
						v68 = m.ExcPending
						if v68 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_cash_div_float8_1), int32(138), int32(_a_F_cash_div_float8_2))
							v73 = m.ExcPending
							if v73 != 0 {
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
				return base.I64_trunc_sat_f64_s(v40)
			}
		}
	} else {
		v20 = base.F64_div(v4, l1)
		if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v25 = F_float_overflow_error_ext(m, int32(0))
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v39 = v25
				v40 = base.F64_nearest(v39)
				v48 = int32(0)
				if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v40, float64(-9.223372036854776e+18)) == v48) == v48)&base.F64_lt(v40, float64(9.223372036854776e+18)) == v48 {
					F_errstart_cold(m, int32(21), int32(0))
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_cash_div_float8_0), int32(0))
							v68 = m.ExcPending
							if v68 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_cash_div_float8_1), int32(138), int32(_a_F_cash_div_float8_2))
								v73 = m.ExcPending
								if v73 != 0 {
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
					return base.I64_trunc_sat_f64_s(v40)
				}
			}
		} else {
			if base.B2i32(l0 == int64(0))|base.F64_ne(v20, float64(0))|base.F64_eq(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v39 = v20
				v40 = base.F64_nearest(v39)
				v48 = int32(0)
				if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v40, float64(-9.223372036854776e+18)) == v48) == v48)&base.F64_lt(v40, float64(9.223372036854776e+18)) == v48 {
					F_errstart_cold(m, int32(21), int32(0))
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_cash_div_float8_0), int32(0))
							v68 = m.ExcPending
							if v68 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_cash_div_float8_1), int32(138), int32(_a_F_cash_div_float8_2))
								v73 = m.ExcPending
								if v73 != 0 {
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
					return base.I64_trunc_sat_f64_s(v40)
				}
			} else {
				v37 = F_float_underflow_error_ext(m, int32(0))
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					v39 = v37
					v40 = base.F64_nearest(v39)
					v48 = int32(0)
					if base.B2i32(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)))|base.B2i32(base.F64_ge(v40, float64(-9.223372036854776e+18)) == v48) == v48)&base.F64_lt(v40, float64(9.223372036854776e+18)) == v48 {
						F_errstart_cold(m, int32(21), int32(0))
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							v64 = m.ExcPending
							if v64 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_cash_div_float8_0), int32(0))
								v68 = m.ExcPending
								if v68 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_cash_div_float8_1), int32(138), int32(_a_F_cash_div_float8_2))
									v73 = m.ExcPending
									if v73 != 0 {
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
						return base.I64_trunc_sat_f64_s(v40)
					}
				}
			}
		}
	}
}
func F_cash_div_flt4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 float32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_cash_div_float8(m, v2, base.F64_promote_f32(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_cash_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
func F_cash_mi(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14243(m, l0, int32(_a_F_cash_mi_0), int32(112), int32(_a_F_cash_mi_1), int32(_a_F_cash_mi_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_cash_mul_int2(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14244(m, l0, int32(_a_F_cash_mul_int2_0), int32(151), int32(_a_F_cash_mul_int2_1), int32(_a_F_cash_mul_int2_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_cash_mul_int8(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14246(m, l0, int32(_a_F_cash_mul_int8_0), int32(151), int32(_a_F_cash_mul_int8_1), int32(_a_F_cash_mul_int8_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_cash_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	v19 = m.G0
	v21 = v19 - int32(416)
	m.G0 = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = F_PGLC_localeconv(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+41)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
		v39 = int32(_a_F_cash_out_0)
		v40 = int32(46)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
		if v42 == int32(0) {
			v51 = v39
			v52 = v40
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
			if v45 != 0 {
				v51 = v39
				v52 = v40
			} else {
				if v42 == int32(44) {
					v50 = int32(_a_F_cash_out_1)
				} else {
					v50 = int32(_a_F_cash_out_0)
				}
				v51 = v50
				v52 = v42
			}
		}
		if base.Ui32(int32(10)) < base.Ui32(v28) {
			v54 = int32(2)
		} else {
			v54 = v28
		}
		if base.Ui32((v32-int32(7))&int32(255)) < base.Ui32(int32(250)) {
			v56 = int32(3)
		} else {
			v56 = v32
		}
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
		v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
		if v23 < int64(0) {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
			if v68 != 0 {
				v69 = v66
			} else {
				v69 = int32(_a_F_cash_out_2)
			}
			v75 = int32(44)
			v76 = int32(45)
			v77 = int32(47)
			v78 = v69
		} else {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
			v75 = int32(42)
			v76 = int32(43)
			v77 = int32(46)
			v78 = v73
		}
		if v58 != 0 {
			v82 = v57
		} else {
			v82 = int32(_a_F_cash_out_3)
		}
		if v60 != 0 {
			v83 = v59
		} else {
			v83 = v51
		}
		v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v77))))
		v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v75))))
		v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v76))))
		v90 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v21)+415)) = uint8(v90)
		v93 = v23 >> (uint(int64(63)) % 64)
		v98 = base.I32_extend8_s(v54)
		v100 = v21 + int32(415)
		v115 = v23 ^ v93 - v93
		for {
			v116 = int32(0)
			if base.B2i32(v54 == v116)|v98 == v116 {
				v122 = v100 - int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v52)
				v132 = v122
			} else {
				if int32(0) <= v98 {
					v132 = v100
				} else {
					v126 = base.I32_rem_s(v98, base.I32_extend8_s(v56))
					if v126 != 0 {
						v132 = v100
					} else {
						v127 = F_strlen(m, v83)
						mBase = m.M
						v128 = v100 - v127
						if v127 == int32(0) {
							v132 = v128
						} else {
							base.MemoryCopy(m, v128, v83, v127)
							v132 = v128
						}
					}
				}
			}
			v134 = int32(1)
			v135 = v132 - v134
			v136 = int64(10)
			v137 = base.I64_div_u_s(v115, v136)
			v143 = base.I32_wrap_i64(v115-v137*v136) | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v143)
			v146 = v98 - v134
			if base.B2i32(int32(0) <= v146)|base.B2i32(base.Ui64(int64(9)) < base.Ui64(v115)) != 0 {
				v98 = v146
				v100 = v135
				v115 = v137
				continue
			} else {
				break
			}
			break
		}
		switch v85 {
		case 0:
			if v89 == int32(1) {
				v156 = int32(_a_F_cash_out_4)
			} else {
				v156 = int32(_a_F_cash_out_5)
			}
			if v87 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v156
				*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v82
				v163 = F_psprintf(m, int32(_a_F_cash_out_6), v21+int32(80))
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int64(0)
				} else {
					v313 = v163
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v156
				*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v135
				v171 = F_psprintf(m, int32(_a_F_cash_out_6), v21-int32(-64))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int64(0)
				} else {
					v313 = v171
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			}
		default:
			if v89 == int32(1) {
				v177 = int32(_a_F_cash_out_4)
			} else {
				v177 = int32(_a_F_cash_out_5)
			}
			if v89 == int32(2) {
				v182 = int32(_a_F_cash_out_4)
			} else {
				v182 = int32(_a_F_cash_out_5)
			}
			if v87 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v177
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v182
				*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v78
				v191 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(32))
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return int64(0)
				} else {
					v313 = v191
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v177
				*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v182
				*(*int32)(unsafe.Add(mBase, uint32(v21))) = v78
				v199 = F_psprintf(m, int32(_a_F_cash_out_7), v21)
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int64(0)
				} else {
					v313 = v199
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			}
		case 2:
			if v89 == int32(2) {
				v205 = int32(_a_F_cash_out_4)
			} else {
				v205 = int32(_a_F_cash_out_5)
			}
			if v89 == int32(1) {
				v210 = int32(_a_F_cash_out_4)
			} else {
				v210 = int32(_a_F_cash_out_5)
			}
			if v87 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v205
				*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v210
				*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v82
				v219 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(192))
				mBase = m.M
				v220 = m.ExcPending
				if v220 != 0 {
					return int64(0)
				} else {
					v313 = v219
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v205
				*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v210
				*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v135
				v229 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(160))
				mBase = m.M
				v230 = m.ExcPending
				if v230 != 0 {
					return int64(0)
				} else {
					v313 = v229
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			}
		case 3:
			if v87 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+272)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+264)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = v78
				if v89 == int32(1) {
					v238 = int32(_a_F_cash_out_4)
				} else {
					v238 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+268)) = v238
				if v89 == int32(2) {
					v244 = int32(_a_F_cash_out_4)
				} else {
					v244 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+260)) = v244
				v249 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(256))
				mBase = m.M
				v250 = m.ExcPending
				if v250 != 0 {
					return int64(0)
				} else {
					v313 = v249
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v135
				if v89 == int32(2) {
					v258 = int32(_a_F_cash_out_4)
				} else {
					v258 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = v258
				if v89 == int32(1) {
					v264 = int32(_a_F_cash_out_4)
				} else {
					v264 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = v264
				v269 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(224))
				mBase = m.M
				v270 = m.ExcPending
				if v270 != 0 {
					return int64(0)
				} else {
					v313 = v269
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			}
		case 4:
			if v87 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v135
				*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v82
				if v89 == int32(1) {
					v278 = int32(_a_F_cash_out_4)
				} else {
					v278 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+140)) = v278
				if v89 == int32(2) {
					v284 = int32(_a_F_cash_out_4)
				} else {
					v284 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v284
				v289 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(128))
				mBase = m.M
				v290 = m.ExcPending
				if v290 != 0 {
					return int64(0)
				} else {
					v313 = v289
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v78
				*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v135
				if v89 == int32(2) {
					v298 = int32(_a_F_cash_out_4)
				} else {
					v298 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v298
				if v89 == int32(1) {
					v304 = int32(_a_F_cash_out_4)
				} else {
					v304 = int32(_a_F_cash_out_5)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v304
				v309 = F_psprintf(m, int32(_a_F_cash_out_7), v21+int32(96))
				mBase = m.M
				v310 = m.ExcPending
				if v310 != 0 {
					return int64(0)
				} else {
					v313 = v309
					m.G0 = v21 + int32(416)
					return base.I64_extend_i32_u(v313)
				}
			}
		}
	}
}
func F_cash_pl(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14247(m, l0, int32(_a_F_cash_pl_0), int32(99), int32(_a_F_cash_pl_1), int32(_a_F_cash_pl_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
