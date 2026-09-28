package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_float8_lerp(m *base.Module, l0 int64, l1 int64, l2 float64) int64 {
	var v5 float64
	_ = v5
	v5 = base.F64_reinterpret_i64(l0)
	return base.I64_reinterpret_f64(base.F64_add(base.F64_mul(l2, base.F64_sub(base.F64_reinterpret_i64(l1), v5)), v5))
}
func F_float8_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v10&int64(9223372036854775807)) {
		v17 = F_make_result_safe(m, int32(_a_F_float8_numeric_0), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v72 = v17
			m.G0 = v8 + int32(176)
			return base.I64_extend_i32_u(v72)
		}
	} else {
		v21 = base.F64_reinterpret_i64(v10)
		if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_lt(v21, float64(0)) != 0 {
				v29 = F_make_result_safe(m, int32(_a_F_float8_numeric_1), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v72 = v29
					m.G0 = v8 + int32(176)
					return base.I64_extend_i32_u(v72)
				}
			} else {
				v33 = F_make_result_safe(m, int32(_a_F_float8_numeric_2), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v72 = v33
					m.G0 = v8 + int32(176)
					return base.I64_extend_i32_u(v72)
				}
			}
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(15)
			v39 = v8 + int32(32)
			v42 = F_pg_snprintf(m, v39, int32(115), int32(_a_F_float8_numeric_3), v8)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				v44 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v8)+168)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v8)+152)) = v44
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v55 = F_set_var_from_str(m, v39, v39, v8+int32(152), v8+int32(28), v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int64(0)
				} else {
					if v55 == int32(0) {
						v59 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
						v72 = int32(0)
						m.G0 = v8 + int32(176)
						return base.I64_extend_i32_u(v72)
					} else {
						v65 = F_make_result_safe(m, v8+int32(152), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+168))
							if v67 == int32(0) {
								v72 = v65
								m.G0 = v8 + int32(176)
								return base.I64_extend_i32_u(v72)
							} else {
								F_pfree(m, v67)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									v72 = v65
									m.G0 = v8 + int32(176)
									return base.I64_extend_i32_u(v72)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_intercept(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v41 int32
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v51 int64
	_ = v51
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v71 int64
	_ = v71
	var v73 float64
	_ = v73
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v145 float64
	_ = v145
	var v146 int32
	_ = v146
	var v156 float64
	_ = v156
	var v158 float64
	_ = v158
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 float64
	_ = v179
	var v180 int32
	_ = v180
	var v183 float64
	_ = v183
	var v184 int32
	_ = v184
	var v194 float64
	_ = v194
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 float64
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 float64
	_ = v222
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 float64
	_ = v235
	var v236 int32
	_ = v236
	var v245 float64
	_ = v245
	var v257 int64
	_ = v257
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		if v22 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v265 = m.ExcPending
			if v265 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_regr_intercept_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_intercept_1), v15)
				mBase = m.M
				v272 = m.ExcPending
				if v272 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_intercept_2), int32(2985), int32(_a_F_float8_regr_intercept_3))
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v25 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_regr_intercept_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_intercept_1), v15)
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_intercept_2), int32(2985), int32(_a_F_float8_regr_intercept_3))
						mBase = m.M
						v277 = m.ExcPending
						if v277 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				if v28 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_regr_intercept_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_intercept_1), v15)
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_intercept_2), int32(2985), int32(_a_F_float8_regr_intercept_3))
							mBase = m.M
							v277 = m.ExcPending
							if v277 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
					if v29 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v265 = m.ExcPending
						if v265 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_float8_regr_intercept_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_intercept_1), v15)
							mBase = m.M
							v272 = m.ExcPending
							if v272 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_intercept_2), int32(2985), int32(_a_F_float8_regr_intercept_3))
								mBase = m.M
								v277 = m.ExcPending
								if v277 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
						if base.F64_lt(v32, float64(1)) != 0 {
							v35 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
							v257 = int64(0)
						} else {
							v38 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
							if base.F64_eq(v38, float64(0)) != 0 {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
								v257 = int64(0)
							} else {
								v44 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
								v45 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
								v46 = *(*float64)(unsafe.Add(mBase, uint32(v18)+64))
								v48 = base.F64_div(base.F64_mul(v45, v46), v38)
								v51 = base.I64_reinterpret_f64(v48) & int64(9223372036854775807)
								if base.B2i32(base.Ui64(v51-int64(1)) < base.Ui64(int64(4503599627370495)))|base.B2i32(base.Ui64(v51-int64(4503599627370496)) < base.Ui64(int64(9214364837600034816)))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v51)) != 0 {
									v245 = v48
								} else {
									v64 = base.F64_abs(v45)
									if base.F64_eq(v64, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v245 = v48
									} else {
										v67 = base.F64_abs(v46)
										if base.F64_eq(v67, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											v245 = v48
										} else {
											v71 = int64(9218868437227405312)
											v73 = base.F64_abs(v38)
											if base.B2i32(base.Ui64(v71) < base.Ui64(base.I64_reinterpret_f64(v64)))|base.F64_eq(v73, math.Float64frombits(uint64(0x7ff0000000000000)))|(base.B2i32(base.Ui64(v71) < base.Ui64(base.I64_reinterpret_f64(v67)))|base.B2i32(base.Ui64(v71) < base.Ui64(base.I64_reinterpret_f64(v73)))) != 0 {
												v245 = v48
											} else {
												v86 = v15 + int32(28)
												v89 = base.I64_reinterpret_f64(v45)
												v93 = int32(2047)
												v94 = base.I32_wrap_i64(int64(base.Ui64(v89)>>(uint(int64(52))%64))) & v93
												if v94 != v93 {
													if v94 == int32(0) {
														if base.F64_eq(v45, float64(0)) != 0 {
															v108 = v45
															v109 = int32(0)
														} else {
															v104 = F_frexp(m, base.F64_mul(v45, float64(1.8446744073709552e+19)), v86)
															mBase = m.M
															v105 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
															v108 = v104
															v109 = v105 + int32(-64)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v86))) = v109
														v121 = v108
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v86))) = v94 - int32(1022)
														v119 = base.F64_reinterpret_i64(v89&int64(-9218868437227405313) | int64(4602678819172646912))
														v121 = v119
													}
												} else {
													v119 = v45
													v121 = v119
												}
												v123 = v15 + int32(24)
												v126 = base.I64_reinterpret_f64(v46)
												v130 = int32(2047)
												v131 = base.I32_wrap_i64(int64(base.Ui64(v126)>>(uint(int64(52))%64))) & v130
												if v131 != v130 {
													if v131 == int32(0) {
														if base.F64_eq(v46, float64(0)) != 0 {
															v145 = v46
															v146 = int32(0)
														} else {
															v141 = F_frexp(m, base.F64_mul(v46, float64(1.8446744073709552e+19)), v123)
															mBase = m.M
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
															v145 = v141
															v146 = v142 + int32(-64)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v123))) = v146
														v158 = v145
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v123))) = v131 - int32(1022)
														v156 = base.F64_reinterpret_i64(v126&int64(-9218868437227405313) | int64(4602678819172646912))
														v158 = v156
													}
												} else {
													v156 = v46
													v158 = v156
												}
												v161 = v15 + int32(20)
												v164 = base.I64_reinterpret_f64(v38)
												v168 = int32(2047)
												v169 = base.I32_wrap_i64(int64(base.Ui64(v164)>>(uint(int64(52))%64))) & v168
												if v169 != v168 {
													if v169 == int32(0) {
														if base.F64_eq(v38, float64(0)) != 0 {
															v183 = v38
															v184 = int32(0)
														} else {
															v179 = F_frexp(m, base.F64_mul(v38, float64(1.8446744073709552e+19)), v161)
															mBase = m.M
															v180 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
															v183 = v179
															v184 = v180 + int32(-64)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v161))) = v184
														v196 = v183
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v161))) = v169 - int32(1022)
														v194 = base.F64_reinterpret_i64(v164&int64(-9218868437227405313) | int64(4602678819172646912))
														v196 = v194
													}
												} else {
													v194 = v38
													v196 = v194
												}
												v197 = base.F64_div(base.F64_mul(v121, v158), v196)
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
												v199 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
												v202 = v198 + v199 - v201
												if int32(1024) <= v202 {
													v206 = base.F64_mul(v197, float64(8.98846567431158e+307))
													if base.Ui32(v202) < base.Ui32(int32(2047)) {
														v235 = v206
														v236 = v202 - int32(1023)
													} else {
														v213 = int32(3069)
														if base.Ui32(v213) <= base.Ui32(v202) {
															v216 = v213
														} else {
															v216 = v202
														}
														v235 = base.F64_mul(v206, float64(8.98846567431158e+307))
														v236 = v216 - int32(2046)
													}
												} else {
													if int32(-1023) < v202 {
														v235 = v197
														v236 = v202
													} else {
														v222 = base.F64_mul(v197, float64(2.004168360008973e-292))
														if base.Ui32(int32(-1992)) < base.Ui32(v202) {
															v235 = v222
															v236 = v202 + int32(969)
														} else {
															v229 = int32(-2960)
															if base.Ui32(v202) <= base.Ui32(v229) {
																v232 = v229
															} else {
																v232 = v202
															}
															v235 = base.F64_mul(v222, float64(2.004168360008973e-292))
															v236 = v232 + int32(1938)
														}
													}
												}
												v245 = base.F64_mul(v235, base.F64_reinterpret_i64(base.I64_extend_i32_u(v236+int32(1023))<<(uint(int64(52))%64)))
											}
										}
									}
								}
								v257 = base.I64_reinterpret_f64(base.F64_div(base.F64_sub(v44, v245), v32))
							}
						}
						m.G0 = v15 + int32(32)
						return v257
					}
				}
			}
		}
	}
}
func F_float8_regr_r2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v32 int32
	_ = v32
	var v35 float64
	_ = v35
	var v38 int32
	_ = v38
	var v42 float64
	_ = v42
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v56 float64
	_ = v56
	var v59 int64
	_ = v59
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v81 float64
	_ = v81
	var v89 int64
	_ = v89
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		if v19 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_float8_regr_r2_0)
				F_errmsg_internal(m, int32(_a_F_float8_regr_r2_1), v12)
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_regr_r2_2), int32(2985), int32(_a_F_float8_regr_r2_3))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			if v22 != int32(8) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_float8_regr_r2_0)
					F_errmsg_internal(m, int32(_a_F_float8_regr_r2_1), v12)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_regr_r2_2), int32(2985), int32(_a_F_float8_regr_r2_3))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_float8_regr_r2_0)
						F_errmsg_internal(m, int32(_a_F_float8_regr_r2_1), v12)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_regr_r2_2), int32(2985), int32(_a_F_float8_regr_r2_3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					if v26 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_float8_regr_r2_0)
							F_errmsg_internal(m, int32(_a_F_float8_regr_r2_1), v12)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_regr_r2_2), int32(2985), int32(_a_F_float8_regr_r2_3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
						if base.F64_lt(v29, float64(1)) != 0 {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							v89 = int64(0)
						} else {
							v35 = *(*float64)(unsafe.Add(mBase, uint32(v15)+40))
							if base.F64_eq(v35, float64(0)) != 0 {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
								v89 = int64(0)
							} else {
								v42 = *(*float64)(unsafe.Add(mBase, uint32(v15)+56))
								if base.F64_eq(v42, float64(0)) != 0 {
									v89 = int64(4607182418800017408)
								} else {
									v46 = *(*float64)(unsafe.Add(mBase, uint32(v15)+64))
									v47 = base.F64_mul(v46, v46)
									if base.F64_eq(v47, float64(0))|base.F64_eq(base.F64_abs(v47), math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
										v56 = base.F64_mul(v35, v42)
										v59 = base.I64_reinterpret_f64(v56) & int64(9223372036854775807)
										if base.B2i32(v59 == int64(0))|base.B2i32(v59 == int64(9218868437227405312)) == int32(0) {
											v78 = base.F64_div(v47, v56)
										} else {
											v72 = base.F64_div(v46, base.F64_mul(base.F64_sqrt(v35), base.F64_sqrt(v42)))
											v78 = base.F64_mul(v72, v72)
										}
									} else {
										v72 = base.F64_div(v46, base.F64_mul(base.F64_sqrt(v35), base.F64_sqrt(v42)))
										v78 = base.F64_mul(v72, v72)
									}
									if base.F64_gt(v78, float64(1)) != 0 {
										v81 = float64(1)
									} else {
										v81 = v78
									}
									v89 = base.I64_reinterpret_f64(v81)
								}
							}
						}
						m.G0 = v12 + int32(16)
						return v89
					}
				}
			}
		}
	}
}
func F_float8_var_samp(m *base.Module, l0 int32) int64 {
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
	var v36 int64
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
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
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_samp_0)
				F_errmsg_internal(m, int32(_a_F_float8_var_samp_1), v8)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_float8_var_samp_2), int32(2985), int32(_a_F_float8_var_samp_3))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
			if v18 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_samp_0)
					F_errmsg_internal(m, int32(_a_F_float8_var_samp_1), v8)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_float8_var_samp_2), int32(2985), int32(_a_F_float8_var_samp_3))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_samp_0)
						F_errmsg_internal(m, int32(_a_F_float8_var_samp_1), v8)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_float8_var_samp_2), int32(2985), int32(_a_F_float8_var_samp_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_float8_var_samp_0)
							F_errmsg_internal(m, int32(_a_F_float8_var_samp_1), v8)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_float8_var_samp_2), int32(2985), int32(_a_F_float8_var_samp_3))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
						if base.F64_le(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v36 = int64(0)
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							v36 = base.I64_reinterpret_f64(base.F64_div(v31, base.F64_add(v25, float64(-1))))
						}
						m.G0 = v8 + int32(16)
						return v36
					}
				}
			}
		}
	}
}
