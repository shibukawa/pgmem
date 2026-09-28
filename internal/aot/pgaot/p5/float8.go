package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_corr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v40 int32
	_ = v40
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v55 float64
	_ = v55
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v67 float64
	_ = v67
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if v17 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_float8_corr_0)
				F_errmsg_internal(m, int32(_a_F_float8_corr_1), v10)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_corr_2), int32(2985), int32(_a_F_float8_corr_3))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			if v20 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_float8_corr_0)
					F_errmsg_internal(m, int32(_a_F_float8_corr_1), v10)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_corr_2), int32(2985), int32(_a_F_float8_corr_3))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v23 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_float8_corr_0)
						F_errmsg_internal(m, int32(_a_F_float8_corr_1), v10)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_corr_2), int32(2985), int32(_a_F_float8_corr_3))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					if v24 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_float8_corr_0)
							F_errmsg_internal(m, int32(_a_F_float8_corr_1), v10)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_corr_2), int32(2985), int32(_a_F_float8_corr_3))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v27 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
						if base.F64_lt(v27, float64(1)) != 0 {
							v30 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
							v73 = int64(0)
						} else {
							v33 = *(*float64)(unsafe.Add(mBase, uint32(v13)+40))
							if base.F64_ne(v33, float64(0)) != 0 {
								v36 = *(*float64)(unsafe.Add(mBase, uint32(v13)+56))
								if base.F64_ne(v36, float64(0)) != 0 {
									v44 = *(*float64)(unsafe.Add(mBase, uint32(v13)+64))
									v47 = base.F64_mul(base.F64_sqrt(v33), base.F64_sqrt(v36))
									v48 = base.F64_mul(v33, v36)
									v52 = base.I64_reinterpret_f64(v48) & int64(9223372036854775807)
									if v52 == int64(9218868437227405312) {
										v55 = v47
									} else {
										v55 = base.F64_sqrt(v48)
									}
									if v52 == int64(0) {
										v58 = v47
									} else {
										v58 = v55
									}
									v59 = base.F64_div(v44, v58)
									if base.F64_lt(v59, float64(-1)) != 0 {
										v67 = float64(-1)
									} else {
										if base.F64_gt(v59, float64(1)) == int32(0) {
											v67 = v59
										} else {
											v67 = float64(1)
										}
									}
									v73 = base.I64_reinterpret_f64(v67)
								} else {
									v40 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
									v73 = int64(0)
								}
							} else {
								v40 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
								v73 = int64(0)
							}
						}
						m.G0 = v10 + int32(16)
						return v73
					}
				}
			}
		}
	}
}
func F_float8_mul(m *base.Module, l0 float64, l1 float64) float64 {
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v19 float64
	_ = v19
	var v22 int32
	_ = v22
	var v24 float64
	_ = v24
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v7 = base.F64_mul(l0, l1)
	if base.F64_eq(base.F64_abs(l0), v5)|base.F64_ne(base.F64_abs(v7), v5)|base.F64_eq(base.F64_abs(l1), v5) == int32(0) {
		v19 = F_float_overflow_error_ext(m, int32(0))
		v22 = m.ExcPending
		if v22 != 0 {
			return float64(0)
		} else {
			return v19
		}
	} else {
		v24 = float64(0)
		if base.F64_eq(l0, v24)|base.F64_ne(v7, v24)|base.F64_eq(l1, v24) != 0 {
			v35 = v7
			return v35
		} else {
			v33 = F_float_underflow_error_ext(m, int32(0))
			v34 = m.ExcPending
			if v34 != 0 {
				return float64(0)
			} else {
				v35 = v33
				return v35
			}
		}
	}
}
func F_float8_qsort_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_ne(v7, v8)
	}
	return v11
}
func F_float8_regr_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 float64
	_ = v44
	var v45 int64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v49 int64
	_ = v49
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v70 float64
	_ = v70
	var v72 int64
	_ = v72
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v98 float64
	_ = v98
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v112 float64
	_ = v112
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v138 float64
	_ = v138
	var v147 int32
	_ = v147
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v180 float64
	_ = v180
	var v189 int32
	_ = v189
	var v193 float64
	_ = v193
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v202 float64
	_ = v202
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v225 int32
	_ = v225
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v230 float64
	_ = v230
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v274 int32
	_ = v274
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		if v34 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v306 = m.ExcPending
			if v306 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(_a_F_float8_regr_accum_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v27)
				mBase = m.M
				v313 = m.ExcPending
				if v313 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2985), int32(_a_F_float8_regr_accum_3))
					mBase = m.M
					v318 = m.ExcPending
					if v318 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
			if v37 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v306 = m.ExcPending
				if v306 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(_a_F_float8_regr_accum_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v27)
					mBase = m.M
					v313 = m.ExcPending
					if v313 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2985), int32(_a_F_float8_regr_accum_3))
						mBase = m.M
						v318 = m.ExcPending
						if v318 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
				if v40 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v306 = m.ExcPending
					if v306 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(_a_F_float8_regr_accum_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v27)
						mBase = m.M
						v313 = m.ExcPending
						if v313 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2985), int32(_a_F_float8_regr_accum_3))
							mBase = m.M
							v318 = m.ExcPending
							if v318 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
					if v41 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v306 = m.ExcPending
						if v306 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(_a_F_float8_regr_accum_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_accum_1), v27)
							mBase = m.M
							v313 = m.ExcPending
							if v313 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_accum_2), int32(2985), int32(_a_F_float8_regr_accum_3))
								mBase = m.M
								v318 = m.ExcPending
								if v318 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v30)+48))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v46 = base.F64_reinterpret_i64(v45)
						v47 = base.F64_add(v44, v46)
						v48 = *(*float64)(unsafe.Add(mBase, uint32(v30)+32))
						v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
						v50 = base.F64_reinterpret_i64(v49)
						v51 = base.F64_add(v48, v50)
						v52 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
						v54 = base.F64_add(v52, float64(1))
						v55 = *(*float64)(unsafe.Add(mBase, uint32(v30)+64))
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v30)+56))
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v30)+40))
						if base.F64_gt(v52, float64(0)) != 0 {
							v61 = base.F64_sub(base.F64_mul(v46, v54), v47)
							v65 = base.F64_div(float64(1), base.F64_mul(v52, v54))
							v68 = math.Float64frombits(uint64(0x7ff8000000000000))
							v70 = *(*float64)(unsafe.Add(mBase, uint32(v30)+80))
							v72 = v45 & int64(9223372036854775807)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v72) {
								v75 = v68
							} else {
								v75 = v70
							}
							if base.F64_ne(v70, v46) != 0 {
								v77 = v68
							} else {
								v77 = v75
							}
							v82 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
							if base.Ui64(base.I64_reinterpret_f64(v77)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v83 = v56
							} else {
								v83 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), v65), v56)
							}
							v85 = base.F64_sub(base.F64_mul(v50, v54), v51)
							v89 = math.Float64frombits(uint64(0x7ff8000000000000))
							v91 = *(*float64)(unsafe.Add(mBase, uint32(v30)+72))
							v95 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v49&int64(9223372036854775807)))
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v49&int64(9223372036854775807)) {
								v96 = v89
							} else {
								v96 = v91
							}
							if base.F64_ne(v91, v50) != 0 {
								v98 = v89
							} else {
								v98 = v96
							}
							v103 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
							if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v104 = v57
							} else {
								v104 = base.F64_add(base.F64_mul(base.F64_mul(v85, v85), v65), v57)
							}
							if v82|v103 == int32(0) {
								v125 = base.F64_add(base.F64_mul(base.F64_mul(v85, v61), v65), v55)
							} else {
								v112 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.B2i32(v95|base.F64_eq(base.F64_abs(v50), v112)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)) == int32(0))&base.F64_ne(base.F64_abs(v46), v112) != 0 {
									v125 = v55
								} else {
									v125 = math.Float64frombits(uint64(0x7ff8000000000000))
								}
							}
							v126 = base.F64_abs(v125)
							v128 = math.Float64frombits(uint64(0x7ff0000000000000))
							v130 = base.F64_abs(v104)
							v133 = base.F64_eq(base.F64_abs(v51), v128) | base.F64_eq(v130, v128)
							v134 = base.F64_abs(v47)
							v138 = base.F64_abs(v83)
							if base.B2i32(v133|base.F64_eq(v134, v128)|base.F64_eq(v138, v128) == int32(0))&base.F64_ne(v126, v128) != 0 {
								v234 = v83
								v235 = v125
								v236 = v77
								v237 = v98
								v238 = v104
								v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v246 == int32(0) {
									v274 = int32(0)
								} else {
									v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
									switch v249 - int32(435) {
									case 0:
										v274 = int32(1)
									case 1:
										v274 = int32(2)
									default:
										v274 = int32(0)
									}
								}
								if v274 != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v30)+80)) = v236
									*(*float64)(unsafe.Add(mBase, uint32(v30)+72)) = v237
									*(*float64)(unsafe.Add(mBase, uint32(v30)+64)) = v235
									*(*float64)(unsafe.Add(mBase, uint32(v30)+56)) = v234
									*(*float64)(unsafe.Add(mBase, uint32(v30)+48)) = v47
									*(*float64)(unsafe.Add(mBase, uint32(v30)+40)) = v238
									*(*float64)(unsafe.Add(mBase, uint32(v30)+32)) = v51
									*(*float64)(unsafe.Add(mBase, uint32(v30)+24)) = v54
									v297 = v30
									m.G0 = v27 + int32(80)
									return base.I64_extend_i32_u(v297)
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v27)+72)) = v236
									*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = v237
									*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v235
									*(*float64)(unsafe.Add(mBase, uint32(v27)+48)) = v234
									*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = v47
									*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v238
									*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v51
									*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v54
									v295 = F_construct_array_builtin(m, v27+int32(16), int32(8), int32(701))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return int64(0)
									} else {
										v297 = v295
										m.G0 = v27 + int32(80)
										return base.I64_extend_i32_u(v297)
									}
								}
							} else {
								v147 = int32(0)
								v150 = math.Float64frombits(uint64(0x7ff0000000000000))
								v151 = base.F64_eq(base.F64_abs(v48), v150)
								v161 = base.F64_eq(base.F64_abs(v44), v150)
								if base.B2i32(base.B2i32(v133 == v147)|v151 == v147)&base.F64_ne(base.F64_abs(v50), v150)|base.B2i32(v161|base.F64_ne(v134, v150)&base.F64_ne(v138, v150) == v147)&base.F64_ne(base.F64_abs(v46), v150) == v147 {
									if base.F64_ne(v126, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										if base.F64_eq(v138, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v193 = math.Float64frombits(uint64(0x7ff8000000000000))
										} else {
											v193 = v83
										}
										if base.F64_eq(v130, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v197 = math.Float64frombits(uint64(0x7ff8000000000000))
										} else {
											v197 = v104
										}
										v234 = v193
										v235 = v125
										v236 = v77
										v237 = v98
										v238 = v197
										v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v246 == int32(0) {
											v274 = int32(0)
										} else {
											v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
											switch v249 - int32(435) {
											case 0:
												v274 = int32(1)
											case 1:
												v274 = int32(2)
											default:
												v274 = int32(0)
											}
										}
										if v274 != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v30)+80)) = v236
											*(*float64)(unsafe.Add(mBase, uint32(v30)+72)) = v237
											*(*float64)(unsafe.Add(mBase, uint32(v30)+64)) = v235
											*(*float64)(unsafe.Add(mBase, uint32(v30)+56)) = v234
											*(*float64)(unsafe.Add(mBase, uint32(v30)+48)) = v47
											*(*float64)(unsafe.Add(mBase, uint32(v30)+40)) = v238
											*(*float64)(unsafe.Add(mBase, uint32(v30)+32)) = v51
											*(*float64)(unsafe.Add(mBase, uint32(v30)+24)) = v54
											v297 = v30
											m.G0 = v27 + int32(80)
											return base.I64_extend_i32_u(v297)
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v27)+72)) = v236
											*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = v237
											*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v235
											*(*float64)(unsafe.Add(mBase, uint32(v27)+48)) = v234
											*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = v47
											*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v238
											*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v51
											*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v54
											v295 = F_construct_array_builtin(m, v27+int32(16), int32(8), int32(701))
											mBase = m.M
											v296 = m.ExcPending
											if v296 != 0 {
												return int64(0)
											} else {
												v297 = v295
												m.G0 = v27 + int32(80)
												return base.I64_extend_i32_u(v297)
											}
										}
									} else {
										v180 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v50), v180)|v151|(base.F64_eq(base.F64_abs(v46), v180)|v161) != 0 {
											v198 = math.Float64frombits(uint64(0x7ff8000000000000))
											if base.F64_eq(v138, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v202 = v198
											} else {
												v202 = v83
											}
											if base.F64_eq(v130, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												v206 = math.Float64frombits(uint64(0x7ff8000000000000))
											} else {
												v206 = v104
											}
											v234 = v202
											v235 = v198
											v236 = v77
											v237 = v98
											v238 = v206
											v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											if v246 == int32(0) {
												v274 = int32(0)
											} else {
												v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
												switch v249 - int32(435) {
												case 0:
													v274 = int32(1)
												case 1:
													v274 = int32(2)
												default:
													v274 = int32(0)
												}
											}
											if v274 != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(v30)+80)) = v236
												*(*float64)(unsafe.Add(mBase, uint32(v30)+72)) = v237
												*(*float64)(unsafe.Add(mBase, uint32(v30)+64)) = v235
												*(*float64)(unsafe.Add(mBase, uint32(v30)+56)) = v234
												*(*float64)(unsafe.Add(mBase, uint32(v30)+48)) = v47
												*(*float64)(unsafe.Add(mBase, uint32(v30)+40)) = v238
												*(*float64)(unsafe.Add(mBase, uint32(v30)+32)) = v51
												*(*float64)(unsafe.Add(mBase, uint32(v30)+24)) = v54
												v297 = v30
												m.G0 = v27 + int32(80)
												return base.I64_extend_i32_u(v297)
											} else {
												*(*float64)(unsafe.Add(mBase, uint32(v27)+72)) = v236
												*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = v237
												*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v235
												*(*float64)(unsafe.Add(mBase, uint32(v27)+48)) = v234
												*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = v47
												*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v238
												*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v51
												*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v54
												v295 = F_construct_array_builtin(m, v27+int32(16), int32(8), int32(701))
												mBase = m.M
												v296 = m.ExcPending
												if v296 != 0 {
													return int64(0)
												} else {
													v297 = v295
													m.G0 = v27 + int32(80)
													return base.I64_extend_i32_u(v297)
												}
											}
										} else {
											F_float_overflow_error(m)
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_float_overflow_error(m)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v207 = math.Float64frombits(uint64(0x7ff8000000000000))
							v216 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v49&int64(9223372036854775807))) | base.F64_eq(base.F64_abs(v50), math.Float64frombits(uint64(0x7ff0000000000000)))
							if v216 != 0 {
								v217 = v207
							} else {
								v217 = v55
							}
							v225 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v45&int64(9223372036854775807))) | base.F64_eq(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000)))
							if v225 != 0 {
								v226 = v207
							} else {
								v226 = v217
							}
							if v225 != 0 {
								v228 = math.Float64frombits(uint64(0x7ff8000000000000))
							} else {
								v228 = v56
							}
							if v216 != 0 {
								v230 = math.Float64frombits(uint64(0x7ff8000000000000))
							} else {
								v230 = v57
							}
							v234 = v228
							v235 = v226
							v236 = v46
							v237 = v50
							v238 = v230
							v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v246 == int32(0) {
								v274 = int32(0)
							} else {
								v249 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
								switch v249 - int32(435) {
								case 0:
									v274 = int32(1)
								case 1:
									v274 = int32(2)
								default:
									v274 = int32(0)
								}
							}
							if v274 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v30)+80)) = v236
								*(*float64)(unsafe.Add(mBase, uint32(v30)+72)) = v237
								*(*float64)(unsafe.Add(mBase, uint32(v30)+64)) = v235
								*(*float64)(unsafe.Add(mBase, uint32(v30)+56)) = v234
								*(*float64)(unsafe.Add(mBase, uint32(v30)+48)) = v47
								*(*float64)(unsafe.Add(mBase, uint32(v30)+40)) = v238
								*(*float64)(unsafe.Add(mBase, uint32(v30)+32)) = v51
								*(*float64)(unsafe.Add(mBase, uint32(v30)+24)) = v54
								v297 = v30
								m.G0 = v27 + int32(80)
								return base.I64_extend_i32_u(v297)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v27)+72)) = v236
								*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = v237
								*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v235
								*(*float64)(unsafe.Add(mBase, uint32(v27)+48)) = v234
								*(*float64)(unsafe.Add(mBase, uint32(v27)+40)) = v47
								*(*float64)(unsafe.Add(mBase, uint32(v27)+32)) = v238
								*(*float64)(unsafe.Add(mBase, uint32(v27)+24)) = v51
								*(*float64)(unsafe.Add(mBase, uint32(v27)+16)) = v54
								v295 = F_construct_array_builtin(m, v27+int32(16), int32(8), int32(701))
								mBase = m.M
								v296 = m.ExcPending
								if v296 != 0 {
									return int64(0)
								} else {
									v297 = v295
									m.G0 = v27 + int32(80)
									return base.I64_extend_i32_u(v297)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_avgy(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v35 float64
	_ = v35
	var v38 int64
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgy_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v8)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2985), int32(_a_F_float8_regr_avgy_3))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgy_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2985), int32(_a_F_float8_regr_avgy_3))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgy_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v8)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2985), int32(_a_F_float8_regr_avgy_3))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_avgy_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_avgy_1), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_avgy_2), int32(2985), int32(_a_F_float8_regr_avgy_3))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v38 = int64(0)
						} else {
							v30 = *(*int64)(unsafe.Add(mBase, uint32(v11)+80))
							if base.Ui64(v30&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v38 = v30
							} else {
								v35 = *(*float64)(unsafe.Add(mBase, uint32(v11)+48))
								v38 = base.I64_reinterpret_f64(base.F64_div(v35, v25))
							}
						}
						m.G0 = v8 + int32(16)
						return v38
					}
				}
			}
		}
	}
}
func F_float8_regr_slope(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v41 int64
	_ = v41
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_slope_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v8)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2985), int32(_a_F_float8_regr_slope_3))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_slope_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v8)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2985), int32(_a_F_float8_regr_slope_3))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_slope_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v8)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2985), int32(_a_F_float8_regr_slope_3))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_regr_slope_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_slope_1), v8)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_slope_2), int32(2985), int32(_a_F_float8_regr_slope_3))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v41 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							if base.F64_eq(v31, float64(0)) != 0 {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v41 = int64(0)
							} else {
								v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)+64))
								v41 = base.I64_reinterpret_f64(base.F64_div(v37, v31))
							}
						}
						m.G0 = v8 + int32(16)
						return v41
					}
				}
			}
		}
	}
}
