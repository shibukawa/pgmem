package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
func F_brin_bloom_add_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v116 float64
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int64
	_ = v365
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_get_fn_opclass_options(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
		if v33 == int32(1) {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+180))
			if v37 != 0 {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
				v43 = base.F64_convert_i32_u(v38 * int32(291))
			} else {
				v43 = float64(37248)
			}
			v44 = float64(-0.1)
			if v27 == int32(0) {
				v50 = v44
			} else {
				v47 = *(*float64)(unsafe.Add(mBase, uint32(v27)+8))
				if base.F64_eq(v47, float64(0)) != 0 {
					v50 = v44
				} else {
					v50 = v47
				}
			}
			if base.F64_lt(v50, float64(0)) != 0 {
				v56 = base.F64_mul(v43, base.F64_neg(v50))
			} else {
				v56 = v50
			}
			v57 = float64(16)
			if base.F64_gt(v56, v57) != 0 {
				v60 = v56
			} else {
				v60 = v57
			}
			if base.F64_lt(v60, v43) != 0 {
				v62 = v60
			} else {
				v62 = v43
			}
			if base.F64_lt(base.F64_abs(v62), float64(2.147483648e+09)) != 0 {
				v66 = base.I32_trunc_f64_s(v62)
				v68 = v66
			} else {
				v68 = int32(-2147483648)
			}
			v69 = float64(0.01)
			if v27 == int32(0) {
				v75 = v69
			} else {
				v72 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
				if base.F64_eq(v72, float64(0)) != 0 {
					v75 = v69
				} else {
					v75 = v72
				}
			}
			v77 = F_log(m, v75)
			mBase = m.M
			v78 = base.F64_convert_i32_s(v68)
			v82 = base.F64_ceil(base.F64_div(base.F64_mul(v77, v78), float64(-0.4804530139182014)))
			if base.F64_lt(base.F64_abs(v82), float64(2.147483648e+09)) != 0 {
				v86 = base.I32_trunc_f64_s(v82)
				v88 = v86
			} else {
				v88 = int32(-2147483648)
			}
			v92 = base.I32_div_s(v88+int32(7), int32(8))
			if base.Ui32(int32(8145)) <= base.Ui32(v92) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(8144)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v92
					F_errmsg_internal(m, int32(633886), v21)
					mBase = m.M
					v402 = m.ExcPending
					if v402 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(477903), int32(344), int32(94789))
						mBase = m.M
						v407 = m.ExcPending
						if v407 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v96 = v92 + int32(16)
				v97 = F_palloc0(m, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v100 = v92 << (uint(int32(3)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v100
					v102 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v102)
					*(*int32)(unsafe.Add(mBase, uint32(v97))) = v96 << (uint(int32(2)) % 32)
					v110 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v100), float64(0.6931471805599453)), v78)
					v112 = base.F64_floor(v110)
					if base.F64_ge(base.F64_sub(v110, v112), float64(0.5)) != 0 {
						v116 = base.F64_ceil(v110)
					} else {
						v116 = v112
					}
					if base.F64_lt(base.F64_abs(v116), float64(2.147483648e+09)) != 0 {
						v120 = base.I32_trunc_f64_s(v116)
						v122 = v120
					} else {
						v122 = int32(-2147483648)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)) = uint8(v122)
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v124))) = v97
					v126 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)) = uint8(v126)
					v135 = v97
					v138 = F_bloom_get_procinfo(m, v25, v31)
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						v140 = F_FunctionCall1Coll(m, v138, v32, v23)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							v153 = base.I32_wrap_i64(int64(1910056111))
							v158 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v164 = v153 - v158 - int32(1636608428) ^ base.I32_rotl(v158, int32(6))
							v166 = v153 + int32(1021750440)
							v167 = v158 + v166
							v168 = v164 + v167
							v172 = v166 - v164 ^ base.I32_rotl(v164, int32(8))
							v176 = v167 - v172 ^ base.I32_rotl(v172, int32(16))
							v180 = v168 - v176 ^ base.I32_rotl(v176, int32(19))
							v181 = v172 + v168
							v182 = v176 + v181
							v183 = v180 + v182
							v195 = int32(14)
							v197 = v183 ^ (v181 - v180 ^ base.I32_rotl(v180, int32(4))) - base.I32_rotl(v183, v195)
							v202 = v197 ^ (v140 + v182) - base.I32_rotl(v197, int32(11))
							v206 = v202 ^ v183 - base.I32_rotl(v202, int32(25))
							v210 = v206 ^ v197 - base.I32_rotl(v206, int32(16))
							v214 = v210 ^ v202 - base.I32_rotl(v210, int32(4))
							v218 = v214 ^ v206 - base.I32_rotl(v214, v195)
							v228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v135)+8)))
							v229 = base.I64_rem_u_s(base.I64_extend_i32_u(v218)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v218^v210-base.I32_rotl(v218, int32(24))), v228)
							v241 = base.I32_wrap_i64(int64(3125326612))
							v246 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v252 = v241 - v246 - int32(1636608428) ^ base.I32_rotl(v246, int32(6))
							v254 = v241 + int32(1021750440)
							v255 = v246 + v254
							v256 = v252 + v255
							v260 = v254 - v252 ^ base.I32_rotl(v252, int32(8))
							v264 = v255 - v260 ^ base.I32_rotl(v260, int32(16))
							v268 = v256 - v264 ^ base.I32_rotl(v264, int32(19))
							v269 = v260 + v256
							v270 = v264 + v269
							v271 = v268 + v270
							v283 = int32(14)
							v285 = v271 ^ (v269 - v268 ^ base.I32_rotl(v268, int32(4))) - base.I32_rotl(v271, v283)
							v290 = v285 ^ (v140 + v270) - base.I32_rotl(v285, int32(11))
							v294 = v290 ^ v271 - base.I32_rotl(v290, int32(25))
							v298 = v294 ^ v285 - base.I32_rotl(v294, int32(16))
							v302 = v298 ^ v290 - base.I32_rotl(v298, int32(4))
							v306 = v302 ^ v294 - base.I32_rotl(v302, v283)
							v316 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v135)+8)))
							v317 = base.I64_rem_u_s(base.I64_extend_i32_u(v306)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v306^v298-base.I32_rotl(v306, int32(24))), v316)
							v318 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+6)))
							if v318 != int64(0) {
								v333 = v33
								v336 = int64(0)
								for {
									v344 = base.I64_rem_u_s(v336*v317+v229, v316)
									v345 = base.I32_wrap_i64(v344)
									v348 = int32(1) << (uint(v345&int32(7)) % 32)
									v351 = v135 + int32(16) + int32(base.Ui32(v345)>>(uint(int32(3))%32))
									v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
									if v348&v352 == int32(0) {
										v356 = v348 | v352
										*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v356)
										v358 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
										v359 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v358 + v359
										v363 = v359
									} else {
										v363 = v333
									}
									v365 = v336 + int64(1)
									if base.Ui64(v365) < base.Ui64(v318) {
										v333 = v363
										v336 = v365
										continue
									} else {
										break
									}
									break
								}
								v377 = v363
							} else {
								v377 = v33
							}
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v385))) = v135
							m.G0 = v21 + int32(16)
							return v377 & int32(1)
						}
					}
				}
			}
		} else {
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
			v130 = F_pg_detoast_datum(m, v129)
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return int32(0)
			} else {
				v135 = v130
				v138 = F_bloom_get_procinfo(m, v25, v31)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return int32(0)
				} else {
					v140 = F_FunctionCall1Coll(m, v138, v32, v23)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						v153 = base.I32_wrap_i64(int64(1910056111))
						v158 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v164 = v153 - v158 - int32(1636608428) ^ base.I32_rotl(v158, int32(6))
						v166 = v153 + int32(1021750440)
						v167 = v158 + v166
						v168 = v164 + v167
						v172 = v166 - v164 ^ base.I32_rotl(v164, int32(8))
						v176 = v167 - v172 ^ base.I32_rotl(v172, int32(16))
						v180 = v168 - v176 ^ base.I32_rotl(v176, int32(19))
						v181 = v172 + v168
						v182 = v176 + v181
						v183 = v180 + v182
						v195 = int32(14)
						v197 = v183 ^ (v181 - v180 ^ base.I32_rotl(v180, int32(4))) - base.I32_rotl(v183, v195)
						v202 = v197 ^ (v140 + v182) - base.I32_rotl(v197, int32(11))
						v206 = v202 ^ v183 - base.I32_rotl(v202, int32(25))
						v210 = v206 ^ v197 - base.I32_rotl(v206, int32(16))
						v214 = v210 ^ v202 - base.I32_rotl(v210, int32(4))
						v218 = v214 ^ v206 - base.I32_rotl(v214, v195)
						v228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v135)+8)))
						v229 = base.I64_rem_u_s(base.I64_extend_i32_u(v218)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v218^v210-base.I32_rotl(v218, int32(24))), v228)
						v241 = base.I32_wrap_i64(int64(3125326612))
						v246 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v252 = v241 - v246 - int32(1636608428) ^ base.I32_rotl(v246, int32(6))
						v254 = v241 + int32(1021750440)
						v255 = v246 + v254
						v256 = v252 + v255
						v260 = v254 - v252 ^ base.I32_rotl(v252, int32(8))
						v264 = v255 - v260 ^ base.I32_rotl(v260, int32(16))
						v268 = v256 - v264 ^ base.I32_rotl(v264, int32(19))
						v269 = v260 + v256
						v270 = v264 + v269
						v271 = v268 + v270
						v283 = int32(14)
						v285 = v271 ^ (v269 - v268 ^ base.I32_rotl(v268, int32(4))) - base.I32_rotl(v271, v283)
						v290 = v285 ^ (v140 + v270) - base.I32_rotl(v285, int32(11))
						v294 = v290 ^ v271 - base.I32_rotl(v290, int32(25))
						v298 = v294 ^ v285 - base.I32_rotl(v294, int32(16))
						v302 = v298 ^ v290 - base.I32_rotl(v298, int32(4))
						v306 = v302 ^ v294 - base.I32_rotl(v302, v283)
						v316 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v135)+8)))
						v317 = base.I64_rem_u_s(base.I64_extend_i32_u(v306)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v306^v298-base.I32_rotl(v306, int32(24))), v316)
						v318 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+6)))
						if v318 != int64(0) {
							v333 = v33
							v336 = int64(0)
							for {
								v344 = base.I64_rem_u_s(v336*v317+v229, v316)
								v345 = base.I32_wrap_i64(v344)
								v348 = int32(1) << (uint(v345&int32(7)) % 32)
								v351 = v135 + int32(16) + int32(base.Ui32(v345)>>(uint(int32(3))%32))
								v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
								if v348&v352 == int32(0) {
									v356 = v348 | v352
									*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v356)
									v358 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
									v359 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v358 + v359
									v363 = v359
								} else {
									v363 = v333
								}
								v365 = v336 + int64(1)
								if base.Ui64(v365) < base.Ui64(v318) {
									v333 = v363
									v336 = v365
									continue
								} else {
									break
								}
								break
							}
							v377 = v363
						} else {
							v377 = v33
						}
						v385 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v385))) = v135
						m.G0 = v21 + int32(16)
						return v377 & int32(1)
					}
				}
			}
		}
	}
}
func F_brin_inclusion_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 != 0 {
		v157 = v16
		m.G0 = v14 + int32(16)
		return v157
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
		switch v27 - int32(1) {
		case 0:
			v149 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(4))
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				v151 = F_FunctionCall2Coll(m, v149, v20, v26, v23)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v151 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 1:
			v31 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(5))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_FunctionCall2Coll(m, v31, v20, v26, v23)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v35 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 2, 6, 15, 23, 24:
			v81 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, v27)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = F_FunctionCall2Coll(m, v81, v20, v26, v23)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v157 = v83
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 3:
			v40 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = F_FunctionCall2Coll(m, v40, v20, v26, v23)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v42 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 4:
			v47 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(2))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = F_FunctionCall2Coll(m, v47, v20, v26, v23)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v49 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 5, 17:
			v112 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(7))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				v114 = F_FunctionCall2Coll(m, v112, v20, v26, v23)
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					if v114 != 0 {
						v157 = v16
					} else {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
						v157 = v117
					}
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 7, 25, 26:
			v86 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(3))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				v88 = F_FunctionCall2Coll(m, v86, v20, v26, v23)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					if v88 != 0 {
						v157 = v16
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
						v157 = v91
					}
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 8:
			v61 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(11))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = F_FunctionCall2Coll(m, v61, v20, v26, v23)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v63 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 9:
			v54 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(12))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v56 = F_FunctionCall2Coll(m, v54, v20, v26, v23)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v56 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 10:
			v75 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(9))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v77 = F_FunctionCall2Coll(m, v75, v20, v26, v23)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v77 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 11:
			v68 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(10))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = F_FunctionCall2Coll(m, v68, v20, v26, v23)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v70 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return int32(0)
			} else {
				v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v138
				F_errmsg_internal(m, int32(453370), v14)
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(477208), int32(462), int32(88141))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 16:
			v93 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(3))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = F_FunctionCall2Coll(m, v93, v20, v26, v23)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					if v95 != 0 {
						v157 = v16
						m.G0 = v14 + int32(16)
						return v157
					} else {
						v98 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(17))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = F_FunctionCall2Coll(m, v98, v20, v26, v23)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v157 = v100
								m.G0 = v14 + int32(16)
								return v157
							}
						}
					}
				}
			}
		case 19, 20:
			v103 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(5))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v105 = F_FunctionCall2Coll(m, v103, v20, v26, v23)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					if v105 == int32(0) {
						v157 = v16
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
						v157 = v110
					}
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 21:
			v128 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(1))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v130 = F_FunctionCall2Coll(m, v128, v20, v26, v23)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					v157 = base.B2i32(v130 == int32(0))
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		case 22:
			v119 = F_inclusion_get_strategy_procinfo(m, v21, v25, v24, int32(1))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				v121 = F_FunctionCall2Coll(m, v119, v20, v26, v23)
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					if v121 == int32(0) {
						v157 = v16
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
						v157 = v126
					}
					m.G0 = v14 + int32(16)
					return v157
				}
			}
		}
	}
}
func F_brin_inclusion_opcinfo(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_lookup_type_cache(m, int32(16), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = F_palloc0(m, int32(984))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
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
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v24
				return v12
			}
		}
	}
}
func F_brin_metapage_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	if l0&int32(3) != 0 {
	} else {
	}
	v30 = F___memset(m, l0, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v36 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v36)
	v42 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v42)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v42)
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v47 = int32(61585)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v45)+6)) = uint16(v47)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(-1475306246)
	v55 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v55)
	return
}
func F_brin_minmax_multi_distance_float4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(2147483647)
	v7 = v5 & v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2139095041)) <= base.Ui32(v8&v6) {
		if base.Ui32(v7) <= base.Ui32(int32(2139095040)) {
			v18 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v18
			}
		} else {
			v28 = float64(0)
			v29 = F_Float8GetDatum(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		}
	} else {
		if base.Ui32(v7) < base.Ui32(int32(2139095041)) {
			v28 = base.F64_sub(base.F64_promote_f32(base.F32_reinterpret_i32(v5)), base.F64_promote_f32(base.F32_reinterpret_i32(v8)))
			v29 = F_Float8GetDatum(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				return v29
			}
		} else {
			v18 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_brin_minmax_multi_distance_macaddr(m *base.Module, l0 int32) int32 {
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v52 = F_Float8GetDatum(m, base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_sub(base.F64_convert_i32_u(v4), base.F64_convert_i32_u(v7)), v10), base.F64_sub(base.F64_convert_i32_u(v12), base.F64_convert_i32_u(v14))), v10), base.F64_sub(base.F64_convert_i32_u(v20), base.F64_convert_i32_u(v22))), v10), base.F64_sub(base.F64_convert_i32_u(v28), base.F64_convert_i32_u(v30))), v10), base.F64_sub(base.F64_convert_i32_u(v36), base.F64_convert_i32_u(v38))), v10), base.F64_sub(base.F64_convert_i32_u(v44), base.F64_convert_i32_u(v46))), v10))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		return int32(0)
	} else {
		return v52
	}
}
func F_brin_minmax_opcinfo(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(160))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
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
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v19
			return v5
		}
	}
}
func F_brin_summarize_new_values(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_Int64GetDatum(m, int64(4294967295))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_DirectFunctionCall2Coll(m, int32(16), int32(0), v4, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
