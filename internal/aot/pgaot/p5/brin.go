package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_brinRevmapTerminate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_ReleaseBuffer(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 != 0 {
			F_ReleaseBuffer(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_brin_bloom_add_value(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v46 float64
	_ = v46
	var v49 float64
	_ = v49
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_get_fn_opclass_options(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int64(0)
	} else {
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
		if v32 == int32(1) {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+180))
			if v36 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				v42 = base.F64_convert_i32_u(v37 * int32(291))
			} else {
				v42 = float64(37248)
			}
			v43 = float64(-0.1)
			if v26 == int32(0) {
				v49 = v43
			} else {
				v46 = *(*float64)(unsafe.Add(mBase, uint32(v26)+8))
				if base.F64_eq(v46, float64(0)) != 0 {
					v49 = v43
				} else {
					v49 = v46
				}
			}
			if base.F64_lt(v49, float64(0)) != 0 {
				v55 = base.F64_mul(v42, base.F64_neg(v49))
			} else {
				v55 = v49
			}
			v56 = float64(16)
			if base.F64_gt(v55, v56) != 0 {
				v59 = v55
			} else {
				v59 = v56
			}
			if base.F64_lt(v59, v42) != 0 {
				v61 = v59
			} else {
				v61 = v42
			}
			v63 = float64(-4.605170185988091)
			if v26 == int32(0) {
				v70 = v63
			} else {
				v66 = *(*float64)(unsafe.Add(mBase, uint32(v26)+16))
				if base.F64_eq(v66, float64(0)) != 0 {
					v70 = v63
				} else {
					v69 = F_log(m, v66)
					mBase = m.M
					v70 = v69
				}
			}
			v72 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v61))
			v81 = base.I32_div_s(base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_div(base.F64_mul(v70, v72), float64(-0.4804530139182014))))+int32(7), int32(8))
			if base.Ui32(int32(_a_F_brin_bloom_add_value_0)) <= base.Ui32(v81) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v379 = m.ExcPending
				if v379 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(_a_F_brin_bloom_add_value_1)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v81
					F_errmsg_internal(m, int32(_a_F_brin_bloom_add_value_2), v20)
					mBase = m.M
					v385 = m.ExcPending
					if v385 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_bloom_add_value_3), int32(346), int32(_a_F_brin_bloom_add_value_4))
						mBase = m.M
						v390 = m.ExcPending
						if v390 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v85 = v81 + int32(16)
				v86 = F_palloc0(m, v85)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int64(0)
				} else {
					v89 = v81 << (uint(int32(3)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v89
					v91 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)) = uint16(v91)
					*(*int32)(unsafe.Add(mBase, uint32(v86))) = v85 << (uint(int32(2)) % 32)
					v99 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v89), float64(0.6931471805599453)), v72)
					v101 = base.F64_floor(v99)
					if base.F64_ge(base.F64_sub(v99, v101), float64(0.5)) != 0 {
						v105 = base.F64_ceil(v99)
					} else {
						v105 = v101
					}
					v106 = base.I32_trunc_sat_f64_s(v105)
					*(*uint8)(unsafe.Add(mBase, uint32(v86)+6)) = uint8(v106)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v108))) = base.I64_extend_i32_u(v86)
					v111 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)) = uint8(v111)
					v117 = v86
					v123 = F_bloom_get_procinfo(m, v23, v30)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int64(0)
					} else {
						v125 = F_FunctionCall1Coll(m, v123, v31, v22)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							v127 = base.I32_wrap_i64(v125)
							v138 = base.I32_wrap_i64(int64(1910056111))
							v140 = v138 + int32(1021750440)
							v145 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v151 = v138 - v145 - int32(1636608428) ^ base.I32_rotl(v145, int32(6))
							v155 = v140 - v151 ^ base.I32_rotl(v151, int32(8))
							v156 = v145 + v140
							v157 = v151 + v156
							v158 = v155 + v157
							v162 = v156 - v155 ^ base.I32_rotl(v155, int32(16))
							v166 = v157 - v162 ^ base.I32_rotl(v162, int32(19))
							v171 = v162 + v158
							v172 = v166 + v171
							v179 = int32(14)
							v181 = v158 - v166 ^ base.I32_rotl(v166, int32(4)) ^ v172 - base.I32_rotl(v172, v179)
							v186 = v181 ^ (v127 + v171) - base.I32_rotl(v181, int32(11))
							v190 = v172 ^ v186 - base.I32_rotl(v186, int32(25))
							v194 = v190 ^ v181 - base.I32_rotl(v190, int32(16))
							v198 = v194 ^ v186 - base.I32_rotl(v194, int32(4))
							v202 = v198 ^ v190 - base.I32_rotl(v198, v179)
							v212 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
							v213 = base.I64_rem_u_s(base.I64_extend_i32_u(v202)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v202^v194-base.I32_rotl(v202, int32(24))), v212)
							v224 = base.I32_wrap_i64(int64(3125326612))
							v226 = v224 + int32(1021750440)
							v231 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v237 = v224 - v231 - int32(1636608428) ^ base.I32_rotl(v231, int32(6))
							v241 = v226 - v237 ^ base.I32_rotl(v237, int32(8))
							v242 = v231 + v226
							v243 = v237 + v242
							v244 = v241 + v243
							v248 = v242 - v241 ^ base.I32_rotl(v241, int32(16))
							v252 = v243 - v248 ^ base.I32_rotl(v248, int32(19))
							v257 = v248 + v244
							v258 = v252 + v257
							v265 = int32(14)
							v267 = v244 - v252 ^ base.I32_rotl(v252, int32(4)) ^ v258 - base.I32_rotl(v258, v265)
							v272 = v267 ^ (v127 + v257) - base.I32_rotl(v267, int32(11))
							v276 = v258 ^ v272 - base.I32_rotl(v272, int32(25))
							v280 = v276 ^ v267 - base.I32_rotl(v276, int32(16))
							v284 = v280 ^ v272 - base.I32_rotl(v280, int32(4))
							v288 = v284 ^ v276 - base.I32_rotl(v284, v265)
							v298 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
							v299 = base.I64_rem_u_s(base.I64_extend_i32_u(v288)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v288^v280-base.I32_rotl(v288, int32(24))), v298)
							v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
							if v300 != int64(0) {
								v316 = v32
								v318 = int64(0)
								for {
									v326 = base.I64_rem_u_s(v318*v299+v213, v298)
									v327 = base.I32_wrap_i64(v326)
									v330 = int32(1) << (uint(v327&int32(7)) % 32)
									v333 = v117 + int32(16) + int32(base.Ui32(v327)>>(uint(int32(3))%32))
									v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
									if v330&v334 == int32(0) {
										v338 = v330 | v334
										*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v338)
										v340 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
										v341 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v340 + v341
										v345 = v341
									} else {
										v345 = v316
									}
									v347 = v318 + int64(1)
									if base.Ui64(v347) < base.Ui64(v300) {
										v316 = v345
										v318 = v347
										continue
									} else {
										break
									}
									break
								}
								v359 = v345
							} else {
								v359 = v32
							}
							v366 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v366))) = base.I64_extend_i32_u(v117)
							m.G0 = v20 + int32(16)
							return base.I64_extend_i32_u(v359) & int64(1)
						}
					}
				}
			}
		} else {
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
			v115 = F_pg_detoast_datum(m, v114)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int64(0)
			} else {
				v117 = v115
				v123 = F_bloom_get_procinfo(m, v23, v30)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int64(0)
				} else {
					v125 = F_FunctionCall1Coll(m, v123, v31, v22)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int64(0)
					} else {
						v127 = base.I32_wrap_i64(v125)
						v138 = base.I32_wrap_i64(int64(1910056111))
						v140 = v138 + int32(1021750440)
						v145 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v151 = v138 - v145 - int32(1636608428) ^ base.I32_rotl(v145, int32(6))
						v155 = v140 - v151 ^ base.I32_rotl(v151, int32(8))
						v156 = v145 + v140
						v157 = v151 + v156
						v158 = v155 + v157
						v162 = v156 - v155 ^ base.I32_rotl(v155, int32(16))
						v166 = v157 - v162 ^ base.I32_rotl(v162, int32(19))
						v171 = v162 + v158
						v172 = v166 + v171
						v179 = int32(14)
						v181 = v158 - v166 ^ base.I32_rotl(v166, int32(4)) ^ v172 - base.I32_rotl(v172, v179)
						v186 = v181 ^ (v127 + v171) - base.I32_rotl(v181, int32(11))
						v190 = v172 ^ v186 - base.I32_rotl(v186, int32(25))
						v194 = v190 ^ v181 - base.I32_rotl(v190, int32(16))
						v198 = v194 ^ v186 - base.I32_rotl(v194, int32(4))
						v202 = v198 ^ v190 - base.I32_rotl(v198, v179)
						v212 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
						v213 = base.I64_rem_u_s(base.I64_extend_i32_u(v202)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v202^v194-base.I32_rotl(v202, int32(24))), v212)
						v224 = base.I32_wrap_i64(int64(3125326612))
						v226 = v224 + int32(1021750440)
						v231 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v237 = v224 - v231 - int32(1636608428) ^ base.I32_rotl(v231, int32(6))
						v241 = v226 - v237 ^ base.I32_rotl(v237, int32(8))
						v242 = v231 + v226
						v243 = v237 + v242
						v244 = v241 + v243
						v248 = v242 - v241 ^ base.I32_rotl(v241, int32(16))
						v252 = v243 - v248 ^ base.I32_rotl(v248, int32(19))
						v257 = v248 + v244
						v258 = v252 + v257
						v265 = int32(14)
						v267 = v244 - v252 ^ base.I32_rotl(v252, int32(4)) ^ v258 - base.I32_rotl(v258, v265)
						v272 = v267 ^ (v127 + v257) - base.I32_rotl(v267, int32(11))
						v276 = v258 ^ v272 - base.I32_rotl(v272, int32(25))
						v280 = v276 ^ v267 - base.I32_rotl(v276, int32(16))
						v284 = v280 ^ v272 - base.I32_rotl(v280, int32(4))
						v288 = v284 ^ v276 - base.I32_rotl(v284, v265)
						v298 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
						v299 = base.I64_rem_u_s(base.I64_extend_i32_u(v288)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v288^v280-base.I32_rotl(v288, int32(24))), v298)
						v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
						if v300 != int64(0) {
							v316 = v32
							v318 = int64(0)
							for {
								v326 = base.I64_rem_u_s(v318*v299+v213, v298)
								v327 = base.I32_wrap_i64(v326)
								v330 = int32(1) << (uint(v327&int32(7)) % 32)
								v333 = v117 + int32(16) + int32(base.Ui32(v327)>>(uint(int32(3))%32))
								v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
								if v330&v334 == int32(0) {
									v338 = v330 | v334
									*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v338)
									v340 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
									v341 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v340 + v341
									v345 = v341
								} else {
									v345 = v316
								}
								v347 = v318 + int64(1)
								if base.Ui64(v347) < base.Ui64(v300) {
									v316 = v345
									v318 = v347
									continue
								} else {
									break
								}
								break
							}
							v359 = v345
						} else {
							v359 = v32
						}
						v366 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v366))) = base.I64_extend_i32_u(v117)
						m.G0 = v20 + int32(16)
						return base.I64_extend_i32_u(v359) & int64(1)
					}
				}
			}
		}
	}
}
func F_brin_inclusion_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v178 int64
	_ = v178
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int64(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v19 != int64(0) {
		v178 = v16
		m.G0 = v14 + int32(16)
		return v178
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)+48))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
		v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
		switch v29 - int32(1) {
		case 0:
			v165 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(4))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return int64(0)
			} else {
				v167 = F_FunctionCall2Coll(m, v165, v22, v24, v26)
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v167 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 1:
			v33 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(5))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				v37 = F_FunctionCall2Coll(m, v33, v22, v24, v26)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v37 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 2, 6, 15, 23, 24:
			v90 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, v29)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int64(0)
			} else {
				v92 = F_FunctionCall2Coll(m, v90, v22, v24, v26)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int64(0)
				} else {
					v178 = v92
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 3:
			v43 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(1))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = F_FunctionCall2Coll(m, v43, v22, v24, v26)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v45 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 4:
			v51 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(2))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v53 = F_FunctionCall2Coll(m, v51, v22, v24, v26)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v53 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 5, 17:
			v125 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(7))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int64(0)
			} else {
				v127 = F_FunctionCall2Coll(m, v125, v22, v24, v26)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int64(0)
				} else {
					if v127 != int64(0) {
						v178 = v16
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)+16))
						v178 = v132
					}
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 7, 25, 26:
			v95 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(3))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int64(0)
			} else {
				v97 = F_FunctionCall2Coll(m, v95, v22, v24, v26)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					if v97 != int64(0) {
						v178 = v16
					} else {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v101)+16))
						v178 = v102
					}
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 8:
			v67 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(11))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				v69 = F_FunctionCall2Coll(m, v67, v22, v24, v26)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v69 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 9:
			v59 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(12))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int64(0)
			} else {
				v61 = F_FunctionCall2Coll(m, v59, v22, v24, v26)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v61 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 10:
			v83 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(9))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int64(0)
			} else {
				v85 = F_FunctionCall2Coll(m, v83, v22, v24, v26)
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v85 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 11:
			v75 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(10))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int64(0)
			} else {
				v77 = F_FunctionCall2Coll(m, v75, v22, v24, v26)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v77 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return int64(0)
			} else {
				v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v154
				F_errmsg_internal(m, int32(_a_F_brin_inclusion_consistent_0), v14)
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_brin_inclusion_consistent_1), int32(462), int32(_a_F_brin_inclusion_consistent_2))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 16:
			v104 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(3))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int64(0)
			} else {
				v106 = F_FunctionCall2Coll(m, v104, v22, v24, v26)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int64(0)
				} else {
					if v106 != int64(0) {
						v178 = v16
						m.G0 = v14 + int32(16)
						return v178
					} else {
						v111 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(17))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int64(0)
						} else {
							v113 = F_FunctionCall2Coll(m, v111, v22, v24, v26)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int64(0)
							} else {
								v178 = v113
								m.G0 = v14 + int32(16)
								return v178
							}
						}
					}
				}
			}
		case 19, 20:
			v116 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(5))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return int64(0)
			} else {
				v118 = F_FunctionCall2Coll(m, v116, v22, v24, v26)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int64(0)
				} else {
					if v118 == int64(0) {
						v178 = v16
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v122)+16))
						v178 = v123
					}
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 21:
			v143 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(1))
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return int64(0)
			} else {
				v145 = F_FunctionCall2Coll(m, v143, v22, v24, v26)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int64(0)
				} else {
					v178 = base.I64_extend_i32_u(base.B2i32(v145 == int64(0)))
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		case 22:
			v134 = F_inclusion_get_strategy_procinfo(m, v23, v28, v27, int32(1))
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int64(0)
			} else {
				v136 = F_FunctionCall2Coll(m, v134, v22, v24, v26)
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return int64(0)
				} else {
					if v136 == int64(0) {
						v178 = v16
					} else {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v140)+16))
						v178 = v141
					}
					m.G0 = v14 + int32(16)
					return v178
				}
			}
		}
	}
}
func F_brin_inclusion_opcinfo(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_lookup_type_cache(m, int32(16), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = F_palloc0(m, int32(984))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)) = uint8(v14)
			v16 = int32(3)
			*(*uint16)(unsafe.Add(mBase, uint32(v12))) = uint16(v16)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = (v12 + int32(27)) & int32(-8)
			v24 = F_lookup_type_cache(m, v4, int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v24
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
func F_brin_metapage_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v4 = int32(_a_F_brin_metapage_init_0)
	v6 = int32(0)
	if v6|(l0&int32(3)|int32(1)) == v6 {
		v22 = l0 + v4
		v24 = l0 + int32(4)
		if base.Ui32(v24) < base.Ui32(v22) {
			v26 = v22
		} else {
			v26 = v24
		}
		v31 = (l0^int32(-1)+v26)&int32(-4) + int32(4)
		if v31 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v31)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_brin_metapage_init_1)
	v45 = int32(_a_F_brin_metapage_init_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v45)
	v51 = int32(_a_F_brin_metapage_init_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v51)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v51)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v56 = int32(_a_F_brin_metapage_init_4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v54)+6)) = uint16(v56)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(-1475306246)
	v64 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v64)
	return
}
func F_brin_minmax_multi_distance_float4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int64
	_ = v16
	var v27 int64
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = int32(2147483647)
	v6 = v4 & v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(int32(2139095041)) <= base.Ui32(v7&v5) {
		if base.Ui32(v6) <= base.Ui32(int32(2139095040)) {
			v16 = int64(9218868437227405312)
		} else {
			v16 = int64(0)
		}
		return v16
	} else {
		if base.Ui32(int32(2139095040)) < base.Ui32(v6) {
			v27 = int64(9218868437227405312)
		} else {
			v27 = base.I64_reinterpret_f64(base.F64_sub(base.F64_promote_f32(base.F32_reinterpret_i32(v4)), base.F64_promote_f32(base.F32_reinterpret_i32(v7))))
		}
		return v27
	}
}
func F_brin_minmax_multi_distance_macaddr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
	v10 = float64(0.00390625)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	return base.I64_reinterpret_f64(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_sub(base.F64_convert_i32_u(v4), base.F64_convert_i32_u(v7)), v10), base.F64_sub(base.F64_convert_i32_u(v12), base.F64_convert_i32_u(v14))), v10), base.F64_sub(base.F64_convert_i32_u(v20), base.F64_convert_i32_u(v22))), v10), base.F64_sub(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(v30))), v10), base.F64_sub(base.F64_convert_i32_u(v36), base.F64_convert_i32_u(v38))), v10), base.F64_sub(base.F64_convert_i32_u(v44), base.F64_convert_i32_u(v46))), v10))
}
func F_brin_minmax_opcinfo(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc0(m, int32(160))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v9)
		v11 = int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(v5))) = uint16(v11)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = (v5 + int32(23)) & int32(-8)
		v19 = F_lookup_type_cache(m, v3, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v19
			return base.I64_extend_i32_u(v5)
		}
	}
}
func F_brin_summarize_new_values(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(4), int32(0), v4, int64(4294967295))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_verify_brin_page(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = F_get_page_from_raw(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)))
		if v14 != 0 {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
			v16 = int32(8)
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
			if (v15<<(uint(v16)%32)-v18)&int32(_a_F_verify_brin_page_0) != v16 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a_F_verify_brin_page_1)
						F_errmsg(m, int32(_a_F_verify_brin_page_2), v6+int32(-16))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)))
							v48 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = (v47<<(uint(v48)%32) - v46) & int32(_a_F_verify_brin_page_0)
							v59 = F_errdetail(m, int32(_a_F_verify_brin_page_3), v6+int32(-32))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_brin_page_4), int32(108), int32(_a_F_verify_brin_page_5))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+v18)+6)))
				if v25 != l1 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
							F_errmsg(m, int32(_a_F_verify_brin_page_6), v6+int32(-48))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
								v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10+v79)+6)))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v81
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
								v85 = F_errdetail(m, int32(_a_F_verify_brin_page_7), v8)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_brin_page_4), int32(116), int32(_a_F_verify_brin_page_5))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
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
					m.G0 = v8 - int32(-64)
					return v10
				}
			}
		} else {
			m.G0 = v8 - int32(-64)
			return v10
		}
	}
}
