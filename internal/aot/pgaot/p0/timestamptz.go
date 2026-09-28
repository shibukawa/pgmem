package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamptz_age(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v238 int64
	_ = v238
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v377 int64
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int64
	_ = v388
	var v397 int64
	_ = v397
	var v398 int64
	_ = v398
	var v403 int64
	_ = v403
	var v406 int64
	_ = v406
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v417 int64
	_ = v417
	var v424 int64
	_ = v424
	var v435 int64
	_ = v435
	var v437 int64
	_ = v437
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	v19 = m.G0
	v21 = v19 - int32(128)
	m.G0 = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_palloc(m, int32(16))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		if v23 != int64(9223372036854775807) {
			v33 = int64(-9223372036854775807 - 1)
			if v23 != v33 {
				if v25 == int64(-9223372036854775807-1) {
					v81 = int64(9223372036854775807)
					v82 = int32(2147483647)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v82
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v82
					*(*int64)(unsafe.Add(mBase, uint32(v27))) = v81
					m.G0 = v21 + int32(128)
					return base.I64_extend_i32_u(v27)
				} else {
					if v25 != int64(9223372036854775807) {
						v92 = int32(0)
						v94 = F_timestamp2tm(m, v23, v21+int32(28), v21+int32(76), v21+int32(124), v92, v92)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int64(0)
						} else {
							if v94 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v496 = m.ExcPending
								if v496 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v499 = m.ExcPending
									if v499 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz_age_0), int32(0))
										mBase = m.M
										v503 = m.ExcPending
										if v503 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_2), int32(_a_F_timestamptz_age_3))
											mBase = m.M
											v508 = m.ExcPending
											if v508 != 0 {
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
								v102 = int32(0)
								v104 = F_timestamp2tm(m, v25, v21+int32(24), v21+int32(32), v21+int32(120), v102, v102)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int64(0)
								} else {
									if v104 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v496 = m.ExcPending
										if v496 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v499 = m.ExcPending
											if v499 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz_age_0), int32(0))
												mBase = m.M
												v503 = m.ExcPending
												if v503 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_2), int32(_a_F_timestamptz_age_3))
													mBase = m.M
													v508 = m.ExcPending
													if v508 != 0 {
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
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
										v108 = v106 - v107
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
										v111 = v109 - v110
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
										v114 = v112 - v113
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
										v117 = v115 - v116
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
										v120 = v118 - v119
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+124))
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v21)+120))
										v123 = v121 - v122
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
										v127 = base.I64_extend_i32_s(v124 - v125)
										v128 = base.B2i32(v25 <= v23)
										if v128 == int32(0) {
											v131 = int32(0)
											v145 = v131 - v111
											v146 = int64(0) - v127
											v147 = v131 - v117
											v148 = v131 - v120
											v149 = v131 - v123
											v150 = v131 - v114
											v151 = v131 - v108
										} else {
											v145 = v111
											v146 = v127
											v147 = v117
											v148 = v120
											v149 = v123
											v150 = v114
											v151 = v108
										}
										if v149 < int32(0) {
											v154 = int32(-1000000)
											if base.Ui32(v149) <= base.Ui32(v154) {
												v157 = v154
											} else {
												v157 = v149
											}
											v159 = base.B2i32(base.Ui32(v149) < base.Ui32(int32(-1000000)))
											v162 = int32(_a_F_timestamptz_age_4)
											v163 = base.I32_div_u_s(v157-(v149+v159), v162)
											v164 = v163 + v159
											v173 = v148 + (v164 ^ int32(-1))
											v174 = v149 + v164*v162 + v162
										} else {
											v173 = v148
											v174 = v149
										}
										if v173 < int32(0) {
											v178 = int32(-60)
											if base.Ui32(v173) <= base.Ui32(v178) {
												v181 = v178
											} else {
												v181 = v173
											}
											v183 = base.B2i32(base.Ui32(v173) < base.Ui32(int32(-60)))
											v186 = int32(60)
											v187 = base.I32_div_u_s(v181-(v173+v183), v186)
											v188 = v187 + v183
											v197 = v147 + (v188 ^ int32(-1))
											v198 = v173 + v188*v186 + v186
										} else {
											v197 = v147
											v198 = v173
										}
										if v197 < int32(0) {
											v202 = int32(-60)
											if base.Ui32(v197) <= base.Ui32(v202) {
												v205 = v202
											} else {
												v205 = v197
											}
											v207 = base.B2i32(base.Ui32(v197) < base.Ui32(int32(-60)))
											v210 = int32(60)
											v211 = base.I32_div_u_s(v205-(v197+v207), v210)
											v212 = v211 + v207
											v222 = v146 + base.I64_extend_i32_s(v212^int32(-1))
											v223 = v197 + v212*v210 + v210
										} else {
											v222 = v146
											v223 = v197
										}
										if v222 < int64(0) {
											v227 = int64(-24)
											if base.Ui64(v222) <= base.Ui64(v227) {
												v230 = v227
											} else {
												v230 = v222
											}
											v233 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v222) < base.Ui64(int64(-24))))
											v236 = int64(24)
											v237 = base.I64_div_u_s(v230-(v222+v233), v236)
											v238 = v237 + v233
											v248 = v222 + v238*v236 + v236
											v250 = v150 + (base.I32_wrap_i64(v238) ^ int32(-1))
										} else {
											v248 = v222
											v250 = v150
										}
										if v250 < int32(0) {
											v259 = int32(0)
											v262 = base.I32_rem_s(v106, int32(100))
											v267 = base.I32_rem_s(v106, int32(400))
											if v267 != 0 {
												v268 = base.B2i32(v262 != v259)
											} else {
												v268 = int32(1)
											}
											v279 = int32(0)
											v282 = base.I32_rem_s(v107, int32(100))
											v287 = base.I32_rem_s(v107, int32(400))
											if v287 != 0 {
												v288 = base.B2i32(v282 != v279)
											} else {
												v288 = int32(1)
											}
											if v23 < v25 {
												v294 = v109<<(uint(int32(2))%32) + int32(_a_F_timestamptz_age_5) + base.B2i32(v106&int32(3) == v259)&v268*int32(52)
											} else {
												v294 = v110<<(uint(int32(2))%32) + int32(_a_F_timestamptz_age_5) + base.B2i32(v107&int32(3) == v279)&v288*int32(52)
											}
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v294-int32(4))))
											v298 = v145
											v307 = v250
											for {
												v317 = v298 - int32(1)
												v318 = v307 + v297
												if v318 < int32(0) {
													v298 = v317
													v307 = v318
													continue
												} else {
													break
												}
												break
											}
											v321 = v317
											v330 = v318
										} else {
											v321 = v145
											v330 = v250
										}
										if v321 < int32(0) {
											v341 = int32(-12)
											if base.Ui32(v321) <= base.Ui32(v341) {
												v344 = v341
											} else {
												v344 = v321
											}
											v346 = base.B2i32(base.Ui32(v321) < base.Ui32(int32(-12)))
											v349 = int32(12)
											v350 = base.I32_div_u_s(v344-(v321+v346), v349)
											v351 = v350 + v346
											v360 = v321 + v351*v349 + v349
											v362 = v151 + (v351 ^ int32(-1))
										} else {
											v360 = v321
											v362 = v151
										}
										if v25 <= v23 {
											v377 = v248
											v378 = v223
											v379 = v198
											v380 = v174
											v381 = v330
											v382 = v362
											v383 = v360
										} else {
											v363 = int32(0)
											v377 = int64(0) - v248
											v378 = v363 - v223
											v379 = v363 - v198
											v380 = v363 - v174
											v381 = v363 - v330
											v382 = v363 - v362
											v383 = v363 - v360
										}
										v388 = base.I64_extend_i32_s(v383) + base.I64_extend_i32_s(v382)*int64(12)
										if base.Ui64(v388-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v480 = m.ExcPending
											if v480 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v483 = m.ExcPending
												if v483 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
													mBase = m.M
													v487 = m.ExcPending
													if v487 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
														mBase = m.M
														v492 = m.ExcPending
														if v492 != 0 {
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
											*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v381
											*(*uint32)(unsafe.Add(mBase, uint32(v27)+12)) = uint32(v388)
											v397 = int64(3600000000)
											v398 = int64(0)
											v403 = int64(32)
											v406 = int64(base.Ui64(v377) >> (uint(v403) % 64))
											v409 = int64(4294967295)
											v412 = v377 & v409
											v413 = v397 * v412
											v417 = int64(base.Ui64(v413)>>(uint(v403)%64)) + v397*v406
											v424 = v412*v398 + v417&v409
											*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v377*v398 + v377>>(uint(int64(63))%64)*v397 + v398*v406 + int64(base.Ui64(v417)>>(uint(v403)%64)) + int64(base.Ui64(v424)>>(uint(v403)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v21))) = v413&v409 | v424<<(uint(v403)%64)
											v435 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
											*(*int64)(unsafe.Add(mBase, uint32(v27))) = v435
											v437 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
											if v437 != v435>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v480 = m.ExcPending
												if v480 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v483 = m.ExcPending
													if v483 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
														mBase = m.M
														v487 = m.ExcPending
														if v487 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
															mBase = m.M
															v492 = m.ExcPending
															if v492 != 0 {
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
												v443 = base.I64_extend_i32_s(v378) * int64(60000000)
												v444 = v435 + v443
												*(*int64)(unsafe.Add(mBase, uint32(v27))) = v444
												if base.B2i32(v443 < int64(0))^base.B2i32(v444 < v435) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v480 = m.ExcPending
													if v480 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v483 = m.ExcPending
														if v483 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
															mBase = m.M
															v487 = m.ExcPending
															if v487 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
																mBase = m.M
																v492 = m.ExcPending
																if v492 != 0 {
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
													v452 = base.I64_extend_i32_s(v379) * int64(1000000)
													v453 = v444 + v452
													*(*int64)(unsafe.Add(mBase, uint32(v27))) = v453
													if base.B2i32(v452 < int64(0))^base.B2i32(v453 < v444) != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v480 = m.ExcPending
														if v480 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v483 = m.ExcPending
															if v483 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
																mBase = m.M
																v487 = m.ExcPending
																if v487 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
																	mBase = m.M
																	v492 = m.ExcPending
																	if v492 != 0 {
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
														v459 = base.I64_extend_i32_s(v380)
														v460 = v453 + v459
														*(*int64)(unsafe.Add(mBase, uint32(v27))) = v460
														if base.B2i32(v459 < int64(0))^base.B2i32(v460 < v453) != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v480 = m.ExcPending
															if v480 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v483 = m.ExcPending
																if v483 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
																	mBase = m.M
																	v487 = m.ExcPending
																	if v487 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
																		mBase = m.M
																		v492 = m.ExcPending
																		if v492 != 0 {
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
															if base.B2i32(v381 != int32(2147483647))|base.B2i32(v388 != int64(2147483647))|base.B2i32(v460 != int64(9223372036854775807)) != 0 {
																m.G0 = v21 + int32(128)
																return base.I64_extend_i32_u(v27)
															} else {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v480 = m.ExcPending
																if v480 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(134217858))
																	mBase = m.M
																	v483 = m.ExcPending
																	if v483 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
																		mBase = m.M
																		v487 = m.ExcPending
																		if v487 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_7), int32(_a_F_timestamptz_age_3))
																			mBase = m.M
																			v492 = m.ExcPending
																			if v492 != 0 {
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
											}
										}
									}
								}
							}
						}
					} else {
						v81 = v33
						v82 = int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v82
						*(*int64)(unsafe.Add(mBase, uint32(v27))) = v81
						m.G0 = v21 + int32(128)
						return base.I64_extend_i32_u(v27)
					}
				}
			} else {
				if v25 != int64(-9223372036854775807-1) {
					v81 = v33
					v82 = int32(-2147483648)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v82
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v82
					*(*int64)(unsafe.Add(mBase, uint32(v27))) = v81
					m.G0 = v21 + int32(128)
					return base.I64_extend_i32_u(v27)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_8), int32(_a_F_timestamptz_age_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
		} else {
			if v25 != int64(9223372036854775807) {
				v81 = int64(9223372036854775807)
				v82 = int32(2147483647)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v82
				*(*int64)(unsafe.Add(mBase, uint32(v27))) = v81
				m.G0 = v21 + int32(128)
				return base.I64_extend_i32_u(v27)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamptz_age_6), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamptz_age_1), int32(_a_F_timestamptz_age_9), int32(_a_F_timestamptz_age_3))
							mBase = m.M
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
			}
		}
	}
}
func F_timestamptz_eq_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v35 int64
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_eq_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_eq_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 == int32(1) {
			if base.Ui64(v17-int64(9223372036854775807)) < base.Ui64(int64(2)) {
				v35 = int64(0)
			} else {
				v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) == int32(0)))
			}
		} else {
			v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) == int32(0)))
		}
		m.G0 = v7 + int32(16)
		return v35
	}
}
func F_timestamptz_ge_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ge_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_ge_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v17 <= v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v17 <= v9)
				} else {
					v33 = base.B2i32(v9 == int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 != int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamptz_part(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamptz_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timestamptz_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int64
	_ = v61
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v8-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v61 = int64(0)
		m.G0 = v6 - int32(-64)
		return v61
	} else {
		v22 = int32(0)
		v24 = F_timestamp2tm(m, v8, v4+int32(-48), v4+int32(-44), v4+int32(-52), v22, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			if v24 != 0 {
				v28 = int64(0)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = F_errsave_start(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					if v30 == int32(0) {
						v61 = v28
						m.G0 = v6 - int32(-64)
						return v61
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_time_0), int32(0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v29, int32(_a_F_timestamptz_time_1), int32(2061), int32(_a_F_timestamptz_time_2))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									v61 = v28
									m.G0 = v6 - int32(-64)
									return v61
								}
							}
						}
					}
				}
			} else {
				v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+12)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
				v50 = int32(60)
				v61 = v46 + base.I64_extend_i32_s(v47+(v48+v49*v50)*v50)*int64(1000000)
				m.G0 = v6 - int32(-64)
				return v61
			}
		}
	}
}
