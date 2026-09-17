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
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 float64
	_ = v100
	var v102 float64
	_ = v102
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_get_fn_opclass_options(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int32(0)
	} else {
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)))
		if v33 == int32(1) {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
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
			v64 = float64(-4.605170185988091)
			if v27 == int32(0) {
				v71 = v64
			} else {
				v67 = *(*float64)(unsafe.Add(mBase, uint32(v27)+16))
				if base.F64_eq(v67, float64(0)) != 0 {
					v71 = v64
				} else {
					v70 = F_log(m, v67)
					mBase = m.M
					v71 = v70
				}
			}
			v73 = base.F64_convert_i32_s(base.I32_trunc_sat_f64_s(v62))
			v82 = base.I32_div_s(base.I32_trunc_sat_f64_s(base.F64_ceil(base.F64_div(base.F64_mul(v71, v73), float64(-0.4804530139182014))))+int32(7), int32(8))
			if base.Ui32(int32(_a_F_brin_bloom_add_value_0)) <= base.Ui32(v82) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v377 = m.ExcPending
				if v377 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(_a_F_brin_bloom_add_value_1)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v82
					F_errmsg_internal(m, int32(_a_F_brin_bloom_add_value_2), v21)
					mBase = m.M
					v383 = m.ExcPending
					if v383 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_brin_bloom_add_value_3), int32(344), int32(_a_F_brin_bloom_add_value_4))
						mBase = m.M
						v388 = m.ExcPending
						if v388 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v86 = v82 + int32(16)
				v87 = F_palloc0(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					v90 = v82 << (uint(int32(3)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v90
					v92 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)) = uint16(v92)
					*(*int32)(unsafe.Add(mBase, uint32(v87))) = v86 << (uint(int32(2)) % 32)
					v100 = base.F64_div(base.F64_mul(base.F64_convert_i32_u(v90), float64(0.6931471805599453)), v73)
					v102 = base.F64_floor(v100)
					if base.F64_ge(base.F64_sub(v100, v102), float64(0.5)) != 0 {
						v106 = base.F64_ceil(v100)
					} else {
						v106 = v102
					}
					v107 = base.I32_trunc_sat_f64_s(v106)
					*(*uint8)(unsafe.Add(mBase, uint32(v87)+6)) = uint8(v107)
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v109))) = v87
					v111 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+3)) = uint8(v111)
					v117 = v87
					v123 = F_bloom_get_procinfo(m, v24, v31)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						v125 = F_FunctionCall1Coll(m, v123, v32, v23)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v137 = base.I32_wrap_i64(int64(1910056111))
							v139 = v137 + int32(1021750440)
							v144 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v150 = v137 - v144 - int32(1636608428) ^ base.I32_rotl(v144, int32(6))
							v154 = v139 - v150 ^ base.I32_rotl(v150, int32(8))
							v155 = v144 + v139
							v156 = v150 + v155
							v157 = v154 + v156
							v161 = v155 - v154 ^ base.I32_rotl(v154, int32(16))
							v165 = v156 - v161 ^ base.I32_rotl(v161, int32(19))
							v170 = v161 + v157
							v171 = v165 + v170
							v178 = int32(14)
							v180 = v157 - v165 ^ base.I32_rotl(v165, int32(4)) ^ v171 - base.I32_rotl(v171, v178)
							v185 = v180 ^ (v125 + v170) - base.I32_rotl(v180, int32(11))
							v189 = v171 ^ v185 - base.I32_rotl(v185, int32(25))
							v193 = v189 ^ v180 - base.I32_rotl(v189, int32(16))
							v197 = v193 ^ v185 - base.I32_rotl(v193, int32(4))
							v201 = v197 ^ v189 - base.I32_rotl(v197, v178)
							v211 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
							v212 = base.I64_rem_u_s(base.I64_extend_i32_u(v201)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v201^v193-base.I32_rotl(v201, int32(24))), v211)
							v223 = base.I32_wrap_i64(int64(3125326612))
							v225 = v223 + int32(1021750440)
							v230 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
							v236 = v223 - v230 - int32(1636608428) ^ base.I32_rotl(v230, int32(6))
							v240 = v225 - v236 ^ base.I32_rotl(v236, int32(8))
							v241 = v230 + v225
							v242 = v236 + v241
							v243 = v240 + v242
							v247 = v241 - v240 ^ base.I32_rotl(v240, int32(16))
							v251 = v242 - v247 ^ base.I32_rotl(v247, int32(19))
							v256 = v247 + v243
							v257 = v251 + v256
							v264 = int32(14)
							v266 = v243 - v251 ^ base.I32_rotl(v251, int32(4)) ^ v257 - base.I32_rotl(v257, v264)
							v271 = v266 ^ (v125 + v256) - base.I32_rotl(v266, int32(11))
							v275 = v257 ^ v271 - base.I32_rotl(v271, int32(25))
							v279 = v275 ^ v266 - base.I32_rotl(v275, int32(16))
							v283 = v279 ^ v271 - base.I32_rotl(v279, int32(4))
							v287 = v283 ^ v275 - base.I32_rotl(v283, v264)
							v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
							v298 = base.I64_rem_u_s(base.I64_extend_i32_u(v287)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v287^v279-base.I32_rotl(v287, int32(24))), v297)
							v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
							if v299 != int64(0) {
								v314 = v33
								v317 = int64(0)
								for {
									v325 = base.I64_rem_u_s(v317*v298+v212, v297)
									v326 = base.I32_wrap_i64(v325)
									v329 = int32(1) << (uint(v326&int32(7)) % 32)
									v332 = v117 + int32(16) + int32(base.Ui32(v326)>>(uint(int32(3))%32))
									v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
									if v329&v333 == int32(0) {
										v337 = v329 | v333
										*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v337)
										v339 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
										v340 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v339 + v340
										v344 = v340
									} else {
										v344 = v314
									}
									v346 = v317 + int64(1)
									if base.Ui64(v346) < base.Ui64(v299) {
										v314 = v344
										v317 = v346
										continue
									} else {
										break
									}
									break
								}
								v358 = v344
							} else {
								v358 = v33
							}
							v366 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v366))) = v117
							m.G0 = v21 + int32(16)
							return v358 & int32(1)
						}
					}
				}
			}
		} else {
			v113 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
			v115 = F_pg_detoast_datum(m, v114)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				v117 = v115
				v123 = F_bloom_get_procinfo(m, v24, v31)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = F_FunctionCall1Coll(m, v123, v32, v23)
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v137 = base.I32_wrap_i64(int64(1910056111))
						v139 = v137 + int32(1021750440)
						v144 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v150 = v137 - v144 - int32(1636608428) ^ base.I32_rotl(v144, int32(6))
						v154 = v139 - v150 ^ base.I32_rotl(v150, int32(8))
						v155 = v144 + v139
						v156 = v150 + v155
						v157 = v154 + v156
						v161 = v155 - v154 ^ base.I32_rotl(v154, int32(16))
						v165 = v156 - v161 ^ base.I32_rotl(v161, int32(19))
						v170 = v161 + v157
						v171 = v165 + v170
						v178 = int32(14)
						v180 = v157 - v165 ^ base.I32_rotl(v165, int32(4)) ^ v171 - base.I32_rotl(v171, v178)
						v185 = v180 ^ (v125 + v170) - base.I32_rotl(v180, int32(11))
						v189 = v171 ^ v185 - base.I32_rotl(v185, int32(25))
						v193 = v189 ^ v180 - base.I32_rotl(v189, int32(16))
						v197 = v193 ^ v185 - base.I32_rotl(v193, int32(4))
						v201 = v197 ^ v189 - base.I32_rotl(v197, v178)
						v211 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
						v212 = base.I64_rem_u_s(base.I64_extend_i32_u(v201)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v201^v193-base.I32_rotl(v201, int32(24))), v211)
						v223 = base.I32_wrap_i64(int64(3125326612))
						v225 = v223 + int32(1021750440)
						v230 = base.I32_wrap_i64(int64(0)) ^ int32(-415931063)
						v236 = v223 - v230 - int32(1636608428) ^ base.I32_rotl(v230, int32(6))
						v240 = v225 - v236 ^ base.I32_rotl(v236, int32(8))
						v241 = v230 + v225
						v242 = v236 + v241
						v243 = v240 + v242
						v247 = v241 - v240 ^ base.I32_rotl(v240, int32(16))
						v251 = v242 - v247 ^ base.I32_rotl(v247, int32(19))
						v256 = v247 + v243
						v257 = v251 + v256
						v264 = int32(14)
						v266 = v243 - v251 ^ base.I32_rotl(v251, int32(4)) ^ v257 - base.I32_rotl(v257, v264)
						v271 = v266 ^ (v125 + v256) - base.I32_rotl(v266, int32(11))
						v275 = v257 ^ v271 - base.I32_rotl(v271, int32(25))
						v279 = v275 ^ v266 - base.I32_rotl(v275, int32(16))
						v283 = v279 ^ v271 - base.I32_rotl(v279, int32(4))
						v287 = v283 ^ v275 - base.I32_rotl(v283, v264)
						v297 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v117)+8)))
						v298 = base.I64_rem_u_s(base.I64_extend_i32_u(v287)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v287^v279-base.I32_rotl(v287, int32(24))), v297)
						v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+6)))
						if v299 != int64(0) {
							v314 = v33
							v317 = int64(0)
							for {
								v325 = base.I64_rem_u_s(v317*v298+v212, v297)
								v326 = base.I32_wrap_i64(v325)
								v329 = int32(1) << (uint(v326&int32(7)) % 32)
								v332 = v117 + int32(16) + int32(base.Ui32(v326)>>(uint(int32(3))%32))
								v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
								if v329&v333 == int32(0) {
									v337 = v329 | v333
									*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v337)
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
									v340 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v339 + v340
									v344 = v340
								} else {
									v344 = v314
								}
								v346 = v317 + int64(1)
								if base.Ui64(v346) < base.Ui64(v299) {
									v314 = v344
									v317 = v346
									continue
								} else {
									break
								}
								break
							}
							v358 = v344
						} else {
							v358 = v33
						}
						v366 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v366))) = v117
						m.G0 = v21 + int32(16)
						return v358 & int32(1)
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
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)))
		switch v27 - int32(1) {
		case 0:
			v149 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(4))
			mBase = m.M
			v150 = m.ExcPending
			if v150 != 0 {
				return int32(0)
			} else {
				v151 = F_FunctionCall2Coll(m, v149, v20, v22, v24)
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
			v31 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(5))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = F_FunctionCall2Coll(m, v31, v20, v22, v24)
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
			v81 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, v27)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = F_FunctionCall2Coll(m, v81, v20, v22, v24)
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
			v40 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(1))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = F_FunctionCall2Coll(m, v40, v20, v22, v24)
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
			v47 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(2))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = F_FunctionCall2Coll(m, v47, v20, v22, v24)
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
			v112 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(7))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				v114 = F_FunctionCall2Coll(m, v112, v20, v22, v24)
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
			v86 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(3))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				v88 = F_FunctionCall2Coll(m, v86, v20, v22, v24)
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
			v61 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(11))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = F_FunctionCall2Coll(m, v61, v20, v22, v24)
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
			v54 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(12))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v56 = F_FunctionCall2Coll(m, v54, v20, v22, v24)
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
			v75 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(9))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v77 = F_FunctionCall2Coll(m, v75, v20, v22, v24)
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
			v68 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(10))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = F_FunctionCall2Coll(m, v68, v20, v22, v24)
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
				v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+6)))
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v138
				F_errmsg_internal(m, int32(_a_F_brin_inclusion_consistent_0), v14)
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_brin_inclusion_consistent_1), int32(462), int32(_a_F_brin_inclusion_consistent_2))
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
			v93 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(3))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = F_FunctionCall2Coll(m, v93, v20, v22, v24)
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
						v98 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(17))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = F_FunctionCall2Coll(m, v98, v20, v22, v24)
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
			v103 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(5))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v105 = F_FunctionCall2Coll(m, v103, v20, v22, v24)
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
			v128 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(1))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v130 = F_FunctionCall2Coll(m, v128, v20, v22, v24)
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
			v119 = F_inclusion_get_strategy_procinfo(m, v21, v26, v25, int32(1))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				v121 = F_FunctionCall2Coll(m, v119, v20, v22, v24)
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
							F_errdetail(m, int32(_a_F_verify_brin_page_3), v6+int32(-32))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_verify_brin_page_4), int32(107), int32(_a_F_verify_brin_page_5))
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
								F_errdetail(m, int32(_a_F_verify_brin_page_7), v8)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_verify_brin_page_4), int32(115), int32(_a_F_verify_brin_page_5))
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
