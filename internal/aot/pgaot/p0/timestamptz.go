package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamptz_age(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int64
	_ = v392
	var v400 int32
	_ = v400
	var v403 int64
	_ = v403
	var v404 int64
	_ = v404
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v415 int64
	_ = v415
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v423 int64
	_ = v423
	var v430 int64
	_ = v430
	var v441 int64
	_ = v441
	var v443 int64
	_ = v443
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v458 int64
	_ = v458
	var v459 int64
	_ = v459
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	v19 = m.G0
	v21 = v19 - int32(128)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	v29 = F_palloc(m, int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		return int32(0)
	} else {
		if v24 != int64(9223372036854775807) {
			v35 = int64(-9223372036854775807 - 1)
			if v24 != v35 {
				if v27 == int64(-9223372036854775807-1) {
					v83 = int64(9223372036854775807)
					v84 = int32(2147483647)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
					m.G0 = v21 + int32(128)
					return v29
				} else {
					if v27 != int64(9223372036854775807) {
						v94 = int32(0)
						v96 = F_timestamp2tm(m, v24, v21+int32(28), v21+int32(76), v21+int32(124), v94, v94)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							if v96 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v500 = m.ExcPending
								if v500 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v503 = m.ExcPending
									if v503 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(411195), int32(0))
										mBase = m.M
										v507 = m.ExcPending
										if v507 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(507443), int32(4591), int32(417408))
											mBase = m.M
											v512 = m.ExcPending
											if v512 != 0 {
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
								v104 = int32(0)
								v106 = F_timestamp2tm(m, v27, v21+int32(24), v21+int32(32), v21+int32(120), v104, v104)
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									if v106 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v500 = m.ExcPending
										if v500 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v503 = m.ExcPending
											if v503 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(411195), int32(0))
												mBase = m.M
												v507 = m.ExcPending
												if v507 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(507443), int32(4591), int32(417408))
													mBase = m.M
													v512 = m.ExcPending
													if v512 != 0 {
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
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
										v110 = v108 - v109
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+92))
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
										v113 = v111 - v112
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
										v115 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
										v116 = v114 - v115
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
										v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
										v119 = v117 - v118
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
										v122 = v120 - v121
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v21)+124))
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+120))
										v125 = v123 - v124
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
										v129 = base.I64_extend_i32_s(v126 - v127)
										v130 = base.B2i32(v27 <= v24)
										if v130 == int32(0) {
											v133 = int32(0)
											v147 = int64(0) - v129
											v148 = v133 - v113
											v149 = v133 - v119
											v150 = v133 - v122
											v151 = v133 - v125
											v152 = v133 - v116
											v153 = v133 - v110
										} else {
											v147 = v129
											v148 = v113
											v149 = v119
											v150 = v122
											v151 = v125
											v152 = v116
											v153 = v110
										}
										if v151 < int32(0) {
											v156 = int32(-1000000)
											if base.Ui32(v151) <= base.Ui32(v156) {
												v159 = v156
											} else {
												v159 = v151
											}
											v161 = base.B2i32(base.Ui32(v151) < base.Ui32(int32(-1000000)))
											v164 = int32(1000000)
											v165 = base.I32_div_u_s(v159-(v151+v161), v164)
											v166 = v165 + v161
											v175 = v150 + (v166 ^ int32(-1))
											v176 = v151 + v166*v164 + v164
										} else {
											v175 = v150
											v176 = v151
										}
										if v175 < int32(0) {
											v180 = int32(-60)
											if base.Ui32(v175) <= base.Ui32(v180) {
												v183 = v180
											} else {
												v183 = v175
											}
											v185 = base.B2i32(base.Ui32(v175) < base.Ui32(int32(-60)))
											v188 = int32(60)
											v189 = base.I32_div_u_s(v183-(v175+v185), v188)
											v190 = v189 + v185
											v199 = v149 + (v190 ^ int32(-1))
											v200 = v175 + v190*v188 + v188
										} else {
											v199 = v149
											v200 = v175
										}
										if v199 < int32(0) {
											v204 = int32(-60)
											if base.Ui32(v199) <= base.Ui32(v204) {
												v207 = v204
											} else {
												v207 = v199
											}
											v209 = base.B2i32(base.Ui32(v199) < base.Ui32(int32(-60)))
											v212 = int32(60)
											v213 = base.I32_div_u_s(v207-(v199+v209), v212)
											v214 = v213 + v209
											v224 = v147 + base.I64_extend_i32_s(v214^int32(-1))
											v225 = v199 + v214*v212 + v212
										} else {
											v224 = v147
											v225 = v199
										}
										if v224 < int64(0) {
											v229 = int64(-24)
											if base.Ui64(v224) <= base.Ui64(v229) {
												v232 = v229
											} else {
												v232 = v224
											}
											v235 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v224) < base.Ui64(int64(-24))))
											v238 = int64(24)
											v239 = base.I64_div_u_s(v232-(v224+v235), v238)
											v240 = v239 + v235
											v250 = v224 + v240*v238 + v238
											v252 = v152 + (base.I32_wrap_i64(v240) ^ int32(-1))
										} else {
											v250 = v224
											v252 = v152
										}
										if v252 < int32(0) {
											v256 = base.I32_rem_s(v108, int32(400))
											if v256 != 0 {
												v258 = base.I32_rem_s(v108, int32(100))
												v262 = base.B2i32(v258 != int32(0))
											} else {
												v262 = int32(1)
											}
											v276 = base.I32_rem_s(v109, int32(400))
											if v276 != 0 {
												v278 = base.I32_rem_s(v109, int32(100))
												v282 = base.B2i32(v278 != int32(0))
											} else {
												v282 = int32(1)
											}
											if v24 < v27 {
												v296 = v262&base.B2i32(v108&int32(3) == int32(0))*int32(52) + int32(1674576) + v111<<(uint(int32(2))%32)
											} else {
												v296 = v282&base.B2i32(v109&int32(3) == int32(0))*int32(52) + int32(1674576) + v112<<(uint(int32(2))%32)
											}
											v299 = *(*int32)(unsafe.Add(mBase, uint32(v296-int32(4))))
											v306 = v148
											v310 = v252
											for {
												v319 = v306 - int32(1)
												v320 = v310 + v299
												if v320 < int32(0) {
													v306 = v319
													v310 = v320
													continue
												} else {
													break
												}
												break
											}
											v329 = v319
											v333 = v320
										} else {
											v329 = v148
											v333 = v252
										}
										if v329 < int32(0) {
											v343 = int32(-12)
											if base.Ui32(v329) <= base.Ui32(v343) {
												v346 = v343
											} else {
												v346 = v329
											}
											v348 = base.B2i32(base.Ui32(v329) < base.Ui32(int32(-12)))
											v351 = int32(12)
											v352 = base.I32_div_u_s(v346-(v329+v348), v351)
											v353 = v352 + v348
											v362 = v329 + v353*v351 + v351
											v364 = v153 + (v353 ^ int32(-1))
										} else {
											v362 = v329
											v364 = v153
										}
										if v130 == int32(0) {
											v367 = int32(0)
											v381 = int64(0) - v250
											v382 = v367 - v362
											v383 = v367 - v225
											v384 = v367 - v200
											v385 = v367 - v176
											v386 = v367 - v333
											v387 = v367 - v364
										} else {
											v381 = v250
											v382 = v362
											v383 = v225
											v384 = v200
											v385 = v176
											v386 = v333
											v387 = v364
										}
										v392 = base.I64_extend_i32_s(v382) + base.I64_extend_i32_s(v387)*int64(12)
										if base.Ui64(v392-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v484 = m.ExcPending
											if v484 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v487 = m.ExcPending
												if v487 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(411258), int32(0))
													mBase = m.M
													v491 = m.ExcPending
													if v491 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(507443), int32(4586), int32(417408))
														mBase = m.M
														v496 = m.ExcPending
														if v496 != 0 {
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
											*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v386
											*(*uint32)(unsafe.Add(mBase, uint32(v29)+12)) = uint32(v392)
											v400 = v21 + int32(8)
											v403 = int64(3600000000)
											v404 = int64(0)
											v409 = int64(32)
											v412 = int64(base.Ui64(v381) >> (uint(v409) % 64))
											v415 = int64(4294967295)
											v418 = v381 & v415
											v419 = v403 * v418
											v423 = int64(base.Ui64(v419)>>(uint(v409)%64)) + v403*v412
											v430 = v418*v404 + v423&v415
											*(*int64)(unsafe.Add(mBase, uint32(v400)+8)) = v381*v404 + v381>>(uint(int64(63))%64)*v403 + v404*v412 + int64(base.Ui64(v423)>>(uint(v409)%64)) + int64(base.Ui64(v430)>>(uint(v409)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v400))) = v419&v415 | v430<<(uint(v409)%64)
											v441 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v29))) = v441
											v443 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
											if v443 != v441>>(uint(int64(63))%64) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v484 = m.ExcPending
												if v484 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v487 = m.ExcPending
													if v487 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(411258), int32(0))
														mBase = m.M
														v491 = m.ExcPending
														if v491 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(507443), int32(4586), int32(417408))
															mBase = m.M
															v496 = m.ExcPending
															if v496 != 0 {
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
												v449 = base.I64_extend_i32_s(v383) * int64(60000000)
												v450 = v441 + v449
												*(*int64)(unsafe.Add(mBase, uint32(v29))) = v450
												if base.B2i32(v449 < int64(0))^base.B2i32(v450 < v441) != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v484 = m.ExcPending
													if v484 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v487 = m.ExcPending
														if v487 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(411258), int32(0))
															mBase = m.M
															v491 = m.ExcPending
															if v491 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(507443), int32(4586), int32(417408))
																mBase = m.M
																v496 = m.ExcPending
																if v496 != 0 {
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
													v458 = base.I64_extend_i32_s(v384) * int64(1000000)
													v459 = v450 + v458
													*(*int64)(unsafe.Add(mBase, uint32(v29))) = v459
													if base.B2i32(v458 < int64(0))^base.B2i32(v459 < v450) != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v484 = m.ExcPending
														if v484 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v487 = m.ExcPending
															if v487 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(411258), int32(0))
																mBase = m.M
																v491 = m.ExcPending
																if v491 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(507443), int32(4586), int32(417408))
																	mBase = m.M
																	v496 = m.ExcPending
																	if v496 != 0 {
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
														v465 = base.I64_extend_i32_s(v385)
														v466 = v459 + v465
														*(*int64)(unsafe.Add(mBase, uint32(v29))) = v466
														if base.B2i32(v465 < int64(0))^base.B2i32(v466 < v459) != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v484 = m.ExcPending
															if v484 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(134217858))
																mBase = m.M
																v487 = m.ExcPending
																if v487 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(411258), int32(0))
																	mBase = m.M
																	v491 = m.ExcPending
																	if v491 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(507443), int32(4586), int32(417408))
																		mBase = m.M
																		v496 = m.ExcPending
																		if v496 != 0 {
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
															if v392 != int64(2147483647) {
																m.G0 = v21 + int32(128)
																return v29
															} else {
																if v386 != int32(2147483647) {
																	m.G0 = v21 + int32(128)
																	return v29
																} else {
																	if v466 != int64(9223372036854775807) {
																		m.G0 = v21 + int32(128)
																		return v29
																	} else {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v484 = m.ExcPending
																		if v484 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(134217858))
																			mBase = m.M
																			v487 = m.ExcPending
																			if v487 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(411258), int32(0))
																				mBase = m.M
																				v491 = m.ExcPending
																				if v491 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(507443), int32(4586), int32(417408))
																					mBase = m.M
																					v496 = m.ExcPending
																					if v496 != 0 {
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
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v83 = v35
						v84 = int32(-2147483648)
						*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
						*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
						*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
						m.G0 = v21 + int32(128)
						return v29
					}
				}
			} else {
				if v27 != int64(-9223372036854775807-1) {
					v83 = v35
					v84 = int32(-2147483648)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
					*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
					*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
					m.G0 = v21 + int32(128)
					return v29
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(411258), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(507443), int32(4481), int32(417408))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
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
		} else {
			if v27 != int64(9223372036854775807) {
				v83 = int64(9223372036854775807)
				v84 = int32(2147483647)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v84
				*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v84
				*(*int64)(unsafe.Add(mBase, uint32(v29))) = v83
				m.G0 = v21 + int32(128)
				return v29
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(411258), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(507443), int32(4490), int32(417408))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
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
func F_timestamptz_eq_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[775]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(v65 == int32(0))
}
func F_timestamptz_ge_timestamp(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 == int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 != int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v15 <= v10)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamptz_part(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamptz_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timestamptz_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	if base.Ui64(v9-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v45 = int32(0)
		m.G0 = v6 - int32(-64)
		return v45
	} else {
		v23 = int32(0)
		v25 = F_timestamp2tm(m, v9, v4+int32(-48), v4+int32(-44), v4+int32(-52), v23, v23)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v25 != 0 {
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
						F_errmsg(m, int32(411195), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(510641), int32(2015), int32(383021))
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
			} else {
				v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6)+12)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
				v33 = int32(60)
				v43 = F_Int64GetDatum(m, v29+base.I64_extend_i32_s(v30+(v31+v32*v33)*v33)*int64(1000000))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = v43
					m.G0 = v6 - int32(-64)
					return v45
				}
			}
		}
	}
}
