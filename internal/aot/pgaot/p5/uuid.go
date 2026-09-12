package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v231 float64
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l0 < int32(10000) {
		v326 = v3
		m.G0 = v10 + int32(96)
		return v326
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		if v15 < int64(10000) {
			v326 = v3
			m.G0 = v10 + int32(96)
			return v326
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
			if v18 != int32(1) {
				v326 = v3
				m.G0 = v10 + int32(96)
				return v326
			} else {
				v22 = v14 + int32(16)
				v23 = float64(0)
				v25 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v32 != 0 {
					v33 = int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					if v32 == v33 {
						v69 = v23
						v71 = v25
					} else {
						v41 = v23
						v43 = v25
						v45 = v25
						for {
							v50 = float64(1)
							v51 = v43 + v35
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
							v53 = F_ldexp(m, v50, v52)
							mBase = m.M
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v57 = F_ldexp(m, v50, v56)
							mBase = m.M
							v62 = base.F64_add(base.F64_add(v41, base.F64_div(v50, v57)), base.F64_div(v50, v53))
							v63 = int32(2)
							v64 = v43 + v63
							v66 = v45 + v63
							if v66 != v32&int32(-2) {
								v41 = v62
								v43 = v64
								v45 = v66
								continue
							} else {
								break
							}
							break
						}
						v69 = v62
						v71 = v64
					}
					if v32&v33 != 0 {
						v78 = float64(1)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v35))))
						v82 = F_ldexp(m, v78, v81)
						mBase = m.M
						v85 = base.F64_add(v69, base.F64_div(v78, v82))
					} else {
						v85 = v69
					}
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v87 = base.F64_div(v86, v85)
					v88 = base.F64_convert_i32_u(v32)
					if base.F64_le(v87, base.F64_mul(v88, float64(2.5))) == int32(0) {
						v191 = v87
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
						v231 = v212
					} else {
						v95 = v32 & int32(3)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						v97 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v108 = v97
							v109 = int32(0)
							v110 = v97
							for {
								v115 = v108 + v96
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
								v117 = int32(0)
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
								v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
								v131 = v110 + base.B2i32(v116 == v117) + base.B2i32(v120 == v117) + base.B2i32(v124 == v117) + base.B2i32(v128 == v117)
								v132 = int32(4)
								v133 = v108 + v132
								v135 = v109 + v132
								if v135 != v32&int32(-4) {
									v108 = v133
									v109 = v135
									v110 = v131
									continue
								} else {
									break
								}
								break
							}
							v140 = v133
							v142 = v131
						} else {
							v140 = v97
							v142 = v97
						}
						if v95 != 0 {
							v150 = v140
							v152 = v142
							v153 = v97
							for {
								v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v96))))
								v161 = v152 + base.B2i32(v158 == int32(0))
								v162 = int32(1)
								v165 = v153 + v162
								if v165 != v95 {
									v150 = v150 + v162
									v152 = v161
									v153 = v165
									continue
								} else {
									break
								}
								break
							}
							v172 = v161
						} else {
							v172 = v142
						}
						if v172 == int32(0) {
							v212 = v87
							v231 = v212
						} else {
							v181 = F_log(m, base.F64_div(v88, base.F64_convert_i32_s(v172)))
							mBase = m.M
							v231 = base.F64_mul(v181, v88)
						}
					}
				} else {
					v183 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v185 = base.F64_div(v183, float64(0))
					if base.F64_le(v185, base.F64_mul(base.F64_convert_i32_u(v32), float64(2.5))) != 0 {
						v212 = v185
					} else {
						v191 = v185
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
					}
					v231 = v212
				}
				if base.F64_gt(v231, float64(100000)) != 0 {
					v235 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1107])))
					if v235 != int32(1) {
						v259 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
						v326 = v3
						m.G0 = v10 + int32(96)
						return v326
					} else {
						v240 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							if v240 == int32(0) {
								v259 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
								v326 = v3
								m.G0 = v10 + int32(96)
								return v326
							} else {
								v246 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v246
								*(*float64)(unsafe.Add(mBase, uint32(v10))) = v231
								F_errmsg_internal(m, int32(706268), v10)
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(526319), int32(356), int32(87455))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										v259 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
										v326 = v3
										m.G0 = v10 + int32(96)
										return v326
									}
								}
							}
						}
					}
				} else {
					v262 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1107])))
					v263 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v263), float64(2000)), float64(0.5)), v231) != 0 {
						v270 = int32(1)
						if v262&v270 == int32(0) {
							v326 = v270
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v277 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v278 = m.ExcPending
							if v278 != 0 {
								return int32(0)
							} else {
								if v277 == int32(0) {
									v326 = v270
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v281 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v281
									*(*float64)(unsafe.Add(mBase, uint32(v10)+32)) = v231
									*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v281), float64(2000)), float64(0.5))
									F_errmsg_internal(m, int32(706860), v10+int32(32))
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(526319), int32(374), int32(87455))
										mBase = m.M
										v300 = m.ExcPending
										if v300 != 0 {
											return int32(0)
										} else {
											v326 = v270
											m.G0 = v10 + int32(96)
											return v326
										}
									}
								}
							}
						}
					} else {
						if v262&int32(1) == int32(0) {
							v326 = v3
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v307 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v308 = m.ExcPending
							if v308 != 0 {
								return int32(0)
							} else {
								if v307 == int32(0) {
									v326 = v3
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v311 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v311
									*(*float64)(unsafe.Add(mBase, uint32(v10)+64)) = v231
									F_errmsg_internal(m, int32(706539), v10-int32(-64))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(526319), int32(381), int32(87455))
										mBase = m.M
										v324 = m.ExcPending
										if v324 != 0 {
											return int32(0)
										} else {
											v326 = v3
											m.G0 = v10 + int32(96)
											return v326
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_uuid_generate_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	switch l0 {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L8
	case 3, 5:
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L44
	} else {
		goto L153
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L44
	} else {
		goto L140
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L44
	} else {
		goto L127
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L44
	} else {
		goto L114
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L44
	} else {
		goto L101
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L44
	} else {
		goto L88
	}
L7:
	;
	v302 = F_DirectFunctionCall1Coll(m, int32(3395), int32(0), v9+int32(112))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L44
	} else {
		goto L87
	}
L8:
	;
	v289 = v9 + int32(96)
	F_uuid_generate_random(m, v289)
	mBase = m.M
	F_uuid_unparse(m, v289, v9+int32(112))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L44
	} else {
		goto L86
	}
L9:
	;
	if l0 == int32(3) {
		goto L70
	} else {
		goto L71
	}
L10:
	;
	v129 = v9 + int32(96)
	F_uuid_generate_time(m, v129)
	mBase = m.M
	F_uuid_unparse(m, v129, v9+int32(112))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L44
	} else {
		goto L45
	}
L11:
	;
	v12 = v9 + int32(112)
	goto L15
L12:
	;
	goto L7
L13:
	;
	v125 = F_strlen(m, v114)
	mBase = m.M
	goto L12
L15:
	;
	goto L16
L16:
	;
	v19 = int32(36)
	if (v12^l2)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v118)
	goto L13
L18:
	;
	v99 = v94
	v100 = v95
	v101 = v96
	goto L40
L19:
	;
	if v89 == int32(0) {
		v114 = v87
		v115 = v88
		goto L17
	} else {
		goto L39
	}
L20:
	;
	v87 = l2
	v88 = v12
	v89 = v19
	goto L19
L21:
	;
	goto L22
L22:
	;
	if l2&int32(3) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v56 == int32(0) {
		v114 = v53
		v115 = v54
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v53 = l2
	v54 = v12
	v55 = v19
	v56 = int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v32 = l2
	v33 = v12
	v34 = v19
	goto L27
L27:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
	if v36 == int32(0) {
		v94 = v32
		v95 = v33
		v96 = v34
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v53 = v47
	v54 = v41
	v55 = v43
	v56 = v45
	goto L23
L29:
	;
	v40 = int32(1)
	v41 = v33 + v40
	v43 = v34 - v40
	v44 = int32(0)
	v45 = base.B2i32(v43 != v44)
	v47 = v32 + v40
	if v47&int32(3) == v44 {
		v53 = v47
		v54 = v41
		v55 = v43
		v56 = v45
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v43 != 0 {
		v32 = v47
		v33 = v41
		v34 = v43
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == int32(0) {
		v87 = v53
		v88 = v54
		v89 = v55
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v55) < base.Ui32(int32(4)) {
		v87 = v53
		v88 = v54
		v89 = v55
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v65 = v53
	v66 = v54
	v67 = v55
	goto L35
L35:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v73 = int32(-2139062144)
	if (int32(16843008)-v70|v70)&v73 != v73 {
		v94 = v65
		v95 = v66
		v96 = v67
		goto L18
	} else {
		goto L37
	}
L36:
	;
	v87 = v81
	v88 = v79
	v89 = v83
	goto L19
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v70
	v78 = int32(4)
	v79 = v66 + v78
	v81 = v65 + v78
	v83 = v67 - v78
	if base.Ui32(int32(3)) < base.Ui32(v83) {
		v65 = v81
		v66 = v79
		v67 = v83
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v94 = v87
	v95 = v88
	v96 = v89
	goto L18
L40:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
	if v103 == int32(0) {
		v114 = v99
		v115 = v100
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v114 = v110
	v115 = v108
	goto L17
L42:
	;
	v107 = int32(1)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if v112 != 0 {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	return int32(0)
L45:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	if int32(36) < l3 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v145 = v9 - l3 + int32(148)
	if (l2^v145)&int32(3) != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L7
L49:
	;
	goto L48
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v199)
	if v199&int32(255) == int32(0) {
		goto L49
	} else {
		goto L65
	}
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v198 = l2
	v199 = v151
	v200 = v145
	goto L50
L52:
	;
	goto L53
L53:
	;
	if l2&int32(3) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v155 = l2
	v157 = v145
	goto L57
L55:
	;
	v169 = l2
	v171 = v145
	goto L56
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v176 = int32(-2139062144)
	if (int32(16843008)-v173|v173)&v176 != v176 {
		v198 = v169
		v199 = v173
		v200 = v171
		goto L50
	} else {
		goto L61
	}
L57:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v158)
	if v158 == int32(0) {
		goto L49
	} else {
		goto L59
	}
L58:
	;
	v169 = v165
	v171 = v163
	goto L56
L59:
	;
	v162 = int32(1)
	v163 = v157 + v162
	v165 = v155 + v162
	if v165&int32(3) != 0 {
		v155 = v165
		v157 = v163
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v181 = v169
	v182 = v173
	v183 = v171
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v182
	v185 = int32(4)
	v186 = v183 + v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v189 = v181 + v185
	v193 = int32(-2139062144)
	if (v187|(int32(16843008)-v187))&v193 == v193 {
		v181 = v189
		v182 = v187
		v183 = v186
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v198 = v189
	v199 = v187
	v200 = v186
	goto L50
L64:
	;
	goto L63
L65:
	;
	v207 = v198
	v209 = v200
	goto L66
L66:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)) = uint8(v210)
	v212 = int32(1)
	if v210 != 0 {
		v207 = v207 + v212
		v209 = v209 + v212
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L49
L68:
	;
	goto L67
L69:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+104)))
	v273 = v269&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+104)) = uint8(v273)
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+102)))
	v280 = v275&int32(65295) | l0<<(uint(int32(4))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+102)) = uint16(v280)
	F_uuid_unparse(m, v9+int32(96), v9+int32(112))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L44
	} else {
		goto L85
	}
L70:
	;
	v223 = F_pg_cryptohash_create(m, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L44
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v244 = F_pg_cryptohash_create(m, int32(1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L44
	} else {
		goto L79
	}
L73:
	;
	v225 = F_pg_cryptohash_init(m, v223)
	mBase = m.M
	if v225 < int32(0) {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v229 = F_pg_cryptohash_update(m, v223, l1, int32(16))
	mBase = m.M
	if v229 < int32(0) {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v232 = F_pg_cryptohash_update(m, v223, l2, l3)
	mBase = m.M
	if v232 < int32(0) {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v238 = F_pg_cryptohash_final(m, v223, v9+int32(96), int32(16))
	mBase = m.M
	if v238 < int32(0) {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_pg_cryptohash_free(m, v223)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L44
	} else {
		goto L78
	}
L78:
	;
	goto L69
L79:
	;
	v246 = F_pg_cryptohash_init(m, v244)
	mBase = m.M
	if v246 < int32(0) {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	v250 = F_pg_cryptohash_update(m, v244, l1, int32(16))
	mBase = m.M
	if v250 < int32(0) {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v253 = F_pg_cryptohash_update(m, v244, l2, l3)
	mBase = m.M
	if v253 < int32(0) {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v259 = F_pg_cryptohash_final(m, v244, v9+int32(112), int32(20))
	mBase = m.M
	if v259 < int32(0) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_pg_cryptohash_free(m, v244)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L44
	} else {
		goto L84
	}
L84:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v9)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v9)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v266
	goto L69
L85:
	;
	goto L7
L86:
	;
	goto L7
L87:
	;
	m.G0 = v9 + int32(160)
	return v302
L88:
	;
	if v223 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(586003)
	F_errmsg_internal(m, int32(210370), v9)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L44
	} else {
		goto L99
	}
L90:
	;
	v326 = int32(14086)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v318 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v321 = int32(320083)
	goto L95
L94:
	;
	v321 = int32(139009)
	goto L95
L95:
	;
	if v318 == int32(2) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v324 = int32(14086)
	goto L98
L97:
	;
	v324 = v321
	goto L98
L98:
	;
	v326 = v324
	goto L89
L99:
	;
	F_errfinish(m, int32(521510), int32(338), int32(328282))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L44
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	if v223 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(586003)
	F_errmsg_internal(m, int32(210406), v9+int32(16))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L44
	} else {
		goto L112
	}
L103:
	;
	v361 = int32(14086)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v353 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v356 = int32(320083)
	goto L108
L107:
	;
	v356 = int32(139009)
	goto L108
L108:
	;
	if v353 == int32(2) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v359 = int32(14086)
	goto L111
L110:
	;
	v359 = v356
	goto L111
L111:
	;
	v361 = v359
	goto L102
L112:
	;
	F_errfinish(m, int32(521510), int32(342), int32(328282))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L44
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	if v223 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(586003)
	F_errmsg_internal(m, int32(210336), v9+int32(32))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L44
	} else {
		goto L125
	}
L116:
	;
	v398 = int32(14086)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v390 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v393 = int32(320083)
	goto L121
L120:
	;
	v393 = int32(139009)
	goto L121
L121:
	;
	if v390 == int32(2) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v396 = int32(14086)
	goto L124
L123:
	;
	v396 = v393
	goto L124
L124:
	;
	v398 = v396
	goto L115
L125:
	;
	F_errfinish(m, int32(521510), int32(347), int32(328282))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L44
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	if v244 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(590632)
	F_errmsg_internal(m, int32(210370), v9+int32(48))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L44
	} else {
		goto L138
	}
L129:
	;
	v435 = int32(14086)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v427 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v430 = int32(320083)
	goto L134
L133:
	;
	v430 = int32(139009)
	goto L134
L134:
	;
	if v427 == int32(2) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v433 = int32(14086)
	goto L137
L136:
	;
	v433 = v430
	goto L137
L137:
	;
	v435 = v433
	goto L128
L138:
	;
	F_errfinish(m, int32(521510), int32(357), int32(328282))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L44
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	if v244 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = int32(590632)
	F_errmsg_internal(m, int32(210406), v9-int32(-64))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L44
	} else {
		goto L151
	}
L142:
	;
	v472 = int32(14086)
	goto L141
L143:
	;
	goto L144
L144:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v464 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v467 = int32(320083)
	goto L147
L146:
	;
	v467 = int32(139009)
	goto L147
L147:
	;
	if v464 == int32(2) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v470 = int32(14086)
	goto L150
L149:
	;
	v470 = v467
	goto L150
L150:
	;
	v472 = v470
	goto L141
L151:
	;
	F_errfinish(m, int32(521510), int32(361), int32(328282))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L44
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	if v244 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(590632)
	F_errmsg_internal(m, int32(210336), v9+int32(80))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L44
	} else {
		goto L164
	}
L155:
	;
	v509 = int32(14086)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v501 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v504 = int32(320083)
	goto L160
L159:
	;
	v504 = int32(139009)
	goto L160
L160:
	;
	if v501 == int32(2) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v507 = int32(14086)
	goto L163
L162:
	;
	v507 = v504
	goto L163
L163:
	;
	v509 = v507
	goto L154
L164:
	;
	F_errfinish(m, int32(521510), int32(364), int32(328282))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L44
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_uuid_generate_time(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1472])))
	if v11 == int32(0) {
		v14 = int32(4726688)
		v16 = m.Env.Pgmem_random_bytes(m, v14, int32(6))
		mBase = m.M
		v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1473])))
		v20 = int32(1)
		v21 = v19 | v20
		*(*uint8)(unsafe.Add(mBase, _consts[1473])) = uint8(v21)
		v23 = int32(4726698)
		v25 = m.Env.Pgmem_random_bytes(m, v23, int32(2))
		mBase = m.M
		*(*uint8)(unsafe.Add(mBase, _consts[1472])) = uint8(v20)
		v31 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1474])))
		v33 = v31 & int32(16383)
		*(*uint16)(unsafe.Add(mBase, _consts[1474])) = uint16(v33)
	} else {
	}
	F___gettimeofday(m, v8)
	mBase = m.M
	v36 = int32(4726704)
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	v45 = v37*int64(10000000) + v40*int64(10) + int64(122192928000000000)
	v47 = *(*int64)(unsafe.Add(mBase, _consts[1475]))
	if base.Ui64(v47) < base.Ui64(v45) {
		v51 = v45
	} else {
		v51 = v47 + int64(1)
	}
	*(*int64)(unsafe.Add(mBase, _consts[1475])) = v51
	v54 = int64(base.Ui64(v51) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v54)
	v57 = int64(base.Ui64(v51) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v57)
	v60 = int64(base.Ui64(v51) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v60)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v51)
	v64 = int64(base.Ui64(v51) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v64)
	v67 = int64(base.Ui64(v51) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v67)
	v70 = int64(base.Ui64(v51) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v70)
	v77 = int32(16)
	v78 = base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(56))%64)))&int32(15) | v77
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v78)
	v81 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1474])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)) = uint8(v81)
	v88 = int32(base.Ui32(v81)>>(uint(int32(8))%32))&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v88)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[1473]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = v91
	v94 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1476])))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v94)
	m.G0 = v8 + v77
	return
}
func F_uuid_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = int32(0)
	v26 = v13 + base.B2i32(v19 == int32(123))
	goto L5
L3:
	;
	m.G0 = v11 + int32(16)
	return v15
L4:
	;
	v114 = F_errsave_start(m, v23)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	if v19 == int32(123) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v39)
	v42 = v39 & int32(255)
	goto L9
L9:
	;
	if base.B2i32(base.Ui32(v42-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v42|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	goto L11
L11:
	;
	if base.B2i32(base.Ui32(v57-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v57|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v71)
	v79 = F_strtox_2(m, v11+int32(12), v71, int32(16), int64(4294967295))
	mBase = m.M
	v80 = base.I32_wrap_i64(v79)
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25+v15))) = uint8(v80)
	v83 = v26 + int32(2)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != int32(45) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v97 = v25 + int32(1)
	if v97 != int32(16) {
		v25 = v97
		v26 = v95
		goto L5
	} else {
		goto L24
	}
L15:
	;
	v95 = v83
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v25&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = v26 + int32(3)
	goto L20
L19:
	;
	v91 = v83
	goto L20
L20:
	;
	if v25 != int32(15) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v94 = v91
	goto L23
L22:
	;
	v94 = v83
	goto L23
L23:
	;
	v95 = v94
	goto L14
L24:
	;
	goto L6
L25:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v102 != int32(125) {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	v107 = v95
	goto L27
L27:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	v107 = v95 + int32(1)
	goto L27
L29:
	;
	goto L4
L30:
	;
	if v114 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(455167)
	F_errmsg(m, int32(762095), v11)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, v23, int32(526319), int32(183), int32(455053))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L3
}
func F_uuid_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v15 = v6 + int32(15)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v16 != int32(255) {
			v130 = v16
			v131 = v15
			v133 = v130 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
			v135 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
			return v6
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v19)
			v22 = v6 + int32(14)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			if v23 != int32(255) {
				v130 = v23
				v131 = v22
				v133 = v130 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
				v135 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
				return v6
			} else {
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v26)
				v29 = v6 + int32(13)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				if v30 != int32(255) {
					v130 = v30
					v131 = v29
					v133 = v130 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
					v135 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
					return v6
				} else {
					v33 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v33)
					v36 = v6 + int32(12)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
					if v37 != int32(255) {
						v130 = v37
						v131 = v36
						v133 = v130 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
						v135 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
						return v6
					} else {
						v40 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
						v43 = v6 + int32(11)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
						if v44 != int32(255) {
							v130 = v44
							v131 = v43
							v133 = v130 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
							v135 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
							return v6
						} else {
							v47 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v47)
							v50 = v6 + int32(10)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
							if v51 != int32(255) {
								v130 = v51
								v131 = v50
								v133 = v130 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
								v135 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
								return v6
							} else {
								v54 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v54)
								v57 = v6 + int32(9)
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
								if v58 != int32(255) {
									v130 = v58
									v131 = v57
									v133 = v130 + int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
									v135 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
									return v6
								} else {
									v61 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v61)
									v64 = v6 + int32(8)
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
									if v65 != int32(255) {
										v130 = v65
										v131 = v64
										v133 = v130 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
										v135 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
										return v6
									} else {
										v68 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v68)
										v71 = v6 + int32(7)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
										if v72 != int32(255) {
											v130 = v72
											v131 = v71
											v133 = v130 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
											v135 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
											return v6
										} else {
											v75 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v75)
											v78 = v6 + int32(6)
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
											if v79 != int32(255) {
												v130 = v79
												v131 = v78
												v133 = v130 + int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
												v135 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
												return v6
											} else {
												v82 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v82)
												v85 = v6 + int32(5)
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
												if v86 != int32(255) {
													v130 = v86
													v131 = v85
													v133 = v130 + int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
													v135 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
													return v6
												} else {
													v89 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v89)
													v92 = v6 + int32(4)
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
													if v93 != int32(255) {
														v130 = v93
														v131 = v92
														v133 = v130 + int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
														v135 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
														return v6
													} else {
														v96 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v96)
														v99 = v6 + int32(3)
														v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
														if v100 != int32(255) {
															v130 = v100
															v131 = v99
															v133 = v130 + int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
															v135 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
															return v6
														} else {
															v103 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v103)
															v106 = v6 + int32(2)
															v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
															if v107 != int32(255) {
																v130 = v107
																v131 = v106
																v133 = v130 + int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																v135 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																return v6
															} else {
																v110 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v110)
																v113 = v6 + int32(1)
																v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
																if v114 != int32(255) {
																	v130 = v114
																	v131 = v113
																	v133 = v130 + int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																	v135 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																	return v6
																} else {
																	v117 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v117)
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
																	if v119 != int32(255) {
																		v130 = v119
																		v131 = v6
																		v133 = v130 + int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																		v135 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																		return v6
																	} else {
																		v122 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v122)
																		F_pfree(m, v6)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			v126 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v126)
																			return int32(0)
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_uuid_ns_x500(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = int32(584152)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[1467]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+29)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, _consts[1468]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, _consts[1469]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, _consts[1470]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, _consts[1471]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v16
	v20 = F_DirectFunctionCall1Coll(m, int32(3395), int32(0), v5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(48)
		return v20
	}
}
func F_uuid_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(1546)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v9 == int32(1) {
		v12 = int32(4562096)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[10]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, _consts[10])) = v15
		v18 = F_palloc(m, int32(40))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v22)
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
			F_initHyperLogLog(m, v18+int32(16), int32(10))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(1546)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(1547)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1548)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v18
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v13
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
