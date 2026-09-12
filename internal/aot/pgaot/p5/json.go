package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonEncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int64
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l0 == int32(0) {
		v15 = F_palloc(m, int32(129))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = v15
			if l2 <= int32(1183) {
				switch l2 - int32(1082) {
				case 0:
					if base.Ui32(int32(1)) < base.Ui32(l1-int32(2147483647)) {
						v231 = l1 + int32(2483589)
						v232 = int32(146097)
						v233 = base.I32_div_u_s(v231, v232)
						v234 = int32(3)
						v240 = int32(2)
						v245 = base.I32_div_u_s((v233*int32(1073595727)+v231)<<(uint(v240)%32)|v234, v232)
						v248 = l1 + int32(2451545) + v233*v234 + v245 + int32(32104)
						v249 = int32(1461)
						v250 = base.I32_div_u_s(v248, v249)
						v253 = v250*int32(-1461) + v248
						v255 = v253 << (uint(v240) % 32)
						if base.Ui32(v249) <= base.Ui32(v255) {
							v261 = base.I32_rem_u_s(v253+int32(305), int32(365))
							v266 = v261
						} else {
							v265 = base.I32_rem_u_s(v253+int32(306), int32(366))
							v266 = v265
						}
						v268 = base.I32_div_u_s(v255, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-24)))) = v268 + v250<<(uint(int32(2))%32) - int32(4800)
						v276 = v266 + int32(123)
						v280 = int32(base.Ui32(v276*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-32)))) = v276 - int32(base.Ui32(v280*int32(7834))>>(uint(int32(8))%32))
						v290 = base.I32_rem_u_s(v280+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-28)))) = v290 + int32(1)
						v295 = v8 + int32(-44)
						switch int32(3) {
						case 0, 3:
							v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
							if int32(0) < v299 {
								v304 = v299
							} else {
								v304 = int32(1) - v299
							}
							v306 = F_pg_ultostr_zeropad(m, v19, v304, int32(4))
							mBase = m.M
							v307 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v306))) = uint8(v307)
							v309 = int32(1)
							v311 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
							v312 = int32(2)
							v313 = F_pg_ultostr_zeropad(m, v306+v309, v311, v312)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v307)
							v318 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
							v320 = F_pg_ultostr_zeropad(m, v313+v309, v318, v312)
							mBase = m.M
							v413 = v320
						case 1:
							v324 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
							v326 = base.B2i32(v324 == int32(1))
							if v324 == int32(1) {
								v327 = int32(12)
							} else {
								v327 = int32(16)
							}
							v329 = *(*int32)(unsafe.Add(mBase, uint32(v295+v327)))
							v331 = F_pg_ultostr_zeropad(m, v19, v329, int32(2))
							mBase = m.M
							v332 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v332)
							if v324 == int32(1) {
								v338 = int32(16)
							} else {
								v338 = int32(12)
							}
							v340 = *(*int32)(unsafe.Add(mBase, uint32(v295+v338)))
							v342 = F_pg_ultostr_zeropad(m, v331+int32(1), v340, int32(2))
							mBase = m.M
							v343 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v343)
							v345 = int32(1)
							v347 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
							if int32(0) < v347 {
								v352 = v347
							} else {
								v352 = v345 - v347
							}
							v354 = F_pg_ultostr_zeropad(m, v342+v345, v352, int32(4))
							mBase = m.M
							v413 = v354
						case 2:
							v355 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
							v356 = int32(2)
							v357 = F_pg_ultostr_zeropad(m, v19, v355, v356)
							mBase = m.M
							v358 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v358)
							v360 = int32(1)
							v362 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
							v364 = F_pg_ultostr_zeropad(m, v357+v360, v362, v356)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v358)
							v369 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
							if int32(0) < v369 {
								v374 = v369
							} else {
								v374 = v360 - v369
							}
							v376 = F_pg_ultostr_zeropad(m, v364+v360, v374, int32(4))
							mBase = m.M
							v413 = v376
						default:
							v380 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
							v382 = base.B2i32(v380 == int32(1))
							if v380 == int32(1) {
								v383 = int32(12)
							} else {
								v383 = int32(16)
							}
							v385 = *(*int32)(unsafe.Add(mBase, uint32(v295+v383)))
							v387 = F_pg_ultostr_zeropad(m, v19, v385, int32(2))
							mBase = m.M
							v388 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v388)
							if v380 == int32(1) {
								v394 = int32(16)
							} else {
								v394 = int32(12)
							}
							v396 = *(*int32)(unsafe.Add(mBase, uint32(v295+v394)))
							v398 = F_pg_ultostr_zeropad(m, v387+int32(1), v396, int32(2))
							mBase = m.M
							v399 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v399)
							v401 = int32(1)
							v403 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
							if int32(0) < v403 {
								v408 = v403
							} else {
								v408 = v401 - v403
							}
							v410 = F_pg_ultostr_zeropad(m, v398+v401, v408, int32(4))
							mBase = m.M
							v413 = v410
						}
						v414 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
						if v414 <= int32(0) {
							v418 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1067])))
							*(*uint8)(unsafe.Add(mBase, uint32(v413)+2)) = uint8(v418)
							v421 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1068])))
							*(*uint16)(unsafe.Add(mBase, uint32(v413))) = uint16(v421)
							v425 = v413 + int32(3)
						} else {
							v425 = v413
						}
						v426 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v426)
						m.G0 = v10 - int32(-64)
						return v19
					} else {
						F_EncodeSpecialDate(m, l1, v19)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return v19
						}
					}
				case 1:
					v92 = v8 + int32(-44)
					v93 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					v95 = base.I64_div_s(v93, int64(3600000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v92)+8)) = uint32(v95)
					v100 = base.I64_extend32_s(v95)*int64(-3600000000) + v93
					v102 = base.I64_div_s(v100, int64(60000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v92)+4)) = uint32(v102)
					v107 = base.I64_extend32_s(v102)*int64(-60000000) + v100
					v109 = base.I64_div_s(v107, int64(1000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v92))) = uint32(v109)
					v115 = v109*int64(4293967296) + v107
					*(*uint32)(unsafe.Add(mBase, uint32(v8+int32(-48)))) = uint32(v115)
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
					v124 = int32(2)
					v125 = F_pg_ultostr_zeropad(m, v19, v123, v124)
					mBase = m.M
					v126 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v126)
					v128 = int32(1)
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
					v132 = F_pg_ultostr_zeropad(m, v125+v128, v130, v124)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v126)
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					v139 = F_AppendSeconds(m, v132+v128, v137, v119, v128)
					mBase = m.M
					v142 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v142)
					m.G0 = v10 - int32(-64)
					return v19
				default:
					if l2 == int32(1114) {
						v428 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
						if base.Ui64(v428-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
							F_EncodeSpecialTimestamp(m, v428, v19)
							mBase = m.M
							v434 = m.ExcPending
							if v434 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 - int32(-64)
								return v19
							}
						} else {
							v435 = int32(0)
							v442 = F_timestamp2tm(m, v428, v435, v8+int32(-44), v8+int32(-48), v435, v435)
							mBase = m.M
							v443 = m.ExcPending
							if v443 != 0 {
								return int32(0)
							} else {
								if v442 == int32(0) {
									v448 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v449 = int32(0)
									F_EncodeDateTime(m, v8+int32(-44), v448, v449, v449, v449, int32(4), v19)
									mBase = m.M
									v454 = m.ExcPending
									if v454 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 - int32(-64)
										return v19
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v458 = m.ExcPending
									if v458 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v461 = m.ExcPending
										if v461 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(402169), int32(0))
											mBase = m.M
											v465 = m.ExcPending
											if v465 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495940), int32(375), int32(375787))
												mBase = m.M
												v470 = m.ExcPending
												if v470 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v210 = m.ExcPending
						if v210 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
							F_errmsg_internal(m, int32(53983), v10)
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495940), int32(418), int32(375787))
								mBase = m.M
								v219 = m.ExcPending
								if v219 != 0 {
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
				if l2 == int32(1184) {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
					v146 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					if l3 != 0 {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v147
						v154 = base.I64_extend_i32_s(v147)*int64(-1000000) + v146
					} else {
						v154 = v146
					}
					if base.Ui64(v154-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, v154, v19)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return v19
						}
					} else {
						if l3 != 0 {
							v164 = int32(0)
						} else {
							v164 = v8 + int32(-48)
						}
						if l3 != 0 {
							v172 = int32(0)
						} else {
							v172 = v8 + int32(-56)
						}
						v174 = F_timestamp2tm(m, v154, v164, v8+int32(-44), v8+int32(-52), v172, int32(0))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							if v174 == int32(0) {
								if l3 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(1)
								} else {
								}
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								F_EncodeDateTime(m, v8+int32(-44), v182, int32(1), v184, v185, int32(4), v19)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 - int32(-64)
									return v19
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(402169), int32(0))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495940), int32(414), int32(375787))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
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
				} else {
					if l2 != int32(1266) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v210 = m.ExcPending
						if v210 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
							F_errmsg_internal(m, int32(53983), v10)
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495940), int32(418), int32(375787))
								mBase = m.M
								v219 = m.ExcPending
								if v219 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = v8 + int32(-44)
						v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
						v32 = base.I64_div_s(v30, int64(3600000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v29)+8)) = uint32(v32)
						v37 = v30 + base.I64_extend32_s(v32)*int64(-3600000000)
						v39 = base.I64_div_s(v37, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v29)+4)) = uint32(v39)
						v44 = base.I64_extend32_s(v39)*int64(-60000000) + v37
						v46 = base.I64_div_s(v44, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v29))) = uint32(v46)
						v52 = v46*int64(4293967296) + v44
						*(*uint32)(unsafe.Add(mBase, uint32(v8+int32(-48)))) = uint32(v52)
						v55 = v8 + int32(-52)
						if v55 != 0 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
						} else {
						}
						v59 = v8 + int32(-44)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
						v61 = int32(1)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
						v65 = int32(2)
						v66 = F_pg_ultostr_zeropad(m, v19, v64, v65)
						mBase = m.M
						v67 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v67)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						v73 = F_pg_ultostr_zeropad(m, v66+v61, v71, v65)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v67)
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						v80 = F_AppendSeconds(m, v73+v61, v78, v60, v61)
						mBase = m.M
						v81 = F_EncodeTimezone(m, v80, v62, int32(4))
						mBase = m.M
						v83 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v83)
						m.G0 = v10 - int32(-64)
						return v19
					}
				}
			}
		}
	} else {
		v19 = l0
		if l2 <= int32(1183) {
			switch l2 - int32(1082) {
			case 0:
				if base.Ui32(int32(1)) < base.Ui32(l1-int32(2147483647)) {
					v231 = l1 + int32(2483589)
					v232 = int32(146097)
					v233 = base.I32_div_u_s(v231, v232)
					v234 = int32(3)
					v240 = int32(2)
					v245 = base.I32_div_u_s((v233*int32(1073595727)+v231)<<(uint(v240)%32)|v234, v232)
					v248 = l1 + int32(2451545) + v233*v234 + v245 + int32(32104)
					v249 = int32(1461)
					v250 = base.I32_div_u_s(v248, v249)
					v253 = v250*int32(-1461) + v248
					v255 = v253 << (uint(v240) % 32)
					if base.Ui32(v249) <= base.Ui32(v255) {
						v261 = base.I32_rem_u_s(v253+int32(305), int32(365))
						v266 = v261
					} else {
						v265 = base.I32_rem_u_s(v253+int32(306), int32(366))
						v266 = v265
					}
					v268 = base.I32_div_u_s(v255, int32(1461))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-24)))) = v268 + v250<<(uint(int32(2))%32) - int32(4800)
					v276 = v266 + int32(123)
					v280 = int32(base.Ui32(v276*int32(2141)) >> (uint(int32(16)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-32)))) = v276 - int32(base.Ui32(v280*int32(7834))>>(uint(int32(8))%32))
					v290 = base.I32_rem_u_s(v280+int32(10), int32(12))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-28)))) = v290 + int32(1)
					v295 = v8 + int32(-44)
					switch int32(3) {
					case 0, 3:
						v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
						if int32(0) < v299 {
							v304 = v299
						} else {
							v304 = int32(1) - v299
						}
						v306 = F_pg_ultostr_zeropad(m, v19, v304, int32(4))
						mBase = m.M
						v307 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v306))) = uint8(v307)
						v309 = int32(1)
						v311 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
						v312 = int32(2)
						v313 = F_pg_ultostr_zeropad(m, v306+v309, v311, v312)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v307)
						v318 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
						v320 = F_pg_ultostr_zeropad(m, v313+v309, v318, v312)
						mBase = m.M
						v413 = v320
					case 1:
						v324 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
						v326 = base.B2i32(v324 == int32(1))
						if v324 == int32(1) {
							v327 = int32(12)
						} else {
							v327 = int32(16)
						}
						v329 = *(*int32)(unsafe.Add(mBase, uint32(v295+v327)))
						v331 = F_pg_ultostr_zeropad(m, v19, v329, int32(2))
						mBase = m.M
						v332 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v332)
						if v324 == int32(1) {
							v338 = int32(16)
						} else {
							v338 = int32(12)
						}
						v340 = *(*int32)(unsafe.Add(mBase, uint32(v295+v338)))
						v342 = F_pg_ultostr_zeropad(m, v331+int32(1), v340, int32(2))
						mBase = m.M
						v343 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v343)
						v345 = int32(1)
						v347 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
						if int32(0) < v347 {
							v352 = v347
						} else {
							v352 = v345 - v347
						}
						v354 = F_pg_ultostr_zeropad(m, v342+v345, v352, int32(4))
						mBase = m.M
						v413 = v354
					case 2:
						v355 = *(*int32)(unsafe.Add(mBase, uint32(v295)+12))
						v356 = int32(2)
						v357 = F_pg_ultostr_zeropad(m, v19, v355, v356)
						mBase = m.M
						v358 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v358)
						v360 = int32(1)
						v362 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
						v364 = F_pg_ultostr_zeropad(m, v357+v360, v362, v356)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v364))) = uint8(v358)
						v369 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
						if int32(0) < v369 {
							v374 = v369
						} else {
							v374 = v360 - v369
						}
						v376 = F_pg_ultostr_zeropad(m, v364+v360, v374, int32(4))
						mBase = m.M
						v413 = v376
					default:
						v380 = *(*int32)(unsafe.Add(mBase, _consts[1066]))
						v382 = base.B2i32(v380 == int32(1))
						if v380 == int32(1) {
							v383 = int32(12)
						} else {
							v383 = int32(16)
						}
						v385 = *(*int32)(unsafe.Add(mBase, uint32(v295+v383)))
						v387 = F_pg_ultostr_zeropad(m, v19, v385, int32(2))
						mBase = m.M
						v388 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v388)
						if v380 == int32(1) {
							v394 = int32(16)
						} else {
							v394 = int32(12)
						}
						v396 = *(*int32)(unsafe.Add(mBase, uint32(v295+v394)))
						v398 = F_pg_ultostr_zeropad(m, v387+int32(1), v396, int32(2))
						mBase = m.M
						v399 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v398))) = uint8(v399)
						v401 = int32(1)
						v403 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
						if int32(0) < v403 {
							v408 = v403
						} else {
							v408 = v401 - v403
						}
						v410 = F_pg_ultostr_zeropad(m, v398+v401, v408, int32(4))
						mBase = m.M
						v413 = v410
					}
					v414 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
					if v414 <= int32(0) {
						v418 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1067])))
						*(*uint8)(unsafe.Add(mBase, uint32(v413)+2)) = uint8(v418)
						v421 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1068])))
						*(*uint16)(unsafe.Add(mBase, uint32(v413))) = uint16(v421)
						v425 = v413 + int32(3)
					} else {
						v425 = v413
					}
					v426 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v426)
					m.G0 = v10 - int32(-64)
					return v19
				} else {
					F_EncodeSpecialDate(m, l1, v19)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 - int32(-64)
						return v19
					}
				}
			case 1:
				v92 = v8 + int32(-44)
				v93 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v95 = base.I64_div_s(v93, int64(3600000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v92)+8)) = uint32(v95)
				v100 = base.I64_extend32_s(v95)*int64(-3600000000) + v93
				v102 = base.I64_div_s(v100, int64(60000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v92)+4)) = uint32(v102)
				v107 = base.I64_extend32_s(v102)*int64(-60000000) + v100
				v109 = base.I64_div_s(v107, int64(1000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v92))) = uint32(v109)
				v115 = v109*int64(4293967296) + v107
				*(*uint32)(unsafe.Add(mBase, uint32(v8+int32(-48)))) = uint32(v115)
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
				v124 = int32(2)
				v125 = F_pg_ultostr_zeropad(m, v19, v123, v124)
				mBase = m.M
				v126 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v126)
				v128 = int32(1)
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
				v132 = F_pg_ultostr_zeropad(m, v125+v128, v130, v124)
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v126)
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				v139 = F_AppendSeconds(m, v132+v128, v137, v119, v128)
				mBase = m.M
				v142 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v142)
				m.G0 = v10 - int32(-64)
				return v19
			default:
				if l2 == int32(1114) {
					v428 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					if base.Ui64(v428-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, v428, v19)
						mBase = m.M
						v434 = m.ExcPending
						if v434 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return v19
						}
					} else {
						v435 = int32(0)
						v442 = F_timestamp2tm(m, v428, v435, v8+int32(-44), v8+int32(-48), v435, v435)
						mBase = m.M
						v443 = m.ExcPending
						if v443 != 0 {
							return int32(0)
						} else {
							if v442 == int32(0) {
								v448 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v449 = int32(0)
								F_EncodeDateTime(m, v8+int32(-44), v448, v449, v449, v449, int32(4), v19)
								mBase = m.M
								v454 = m.ExcPending
								if v454 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 - int32(-64)
									return v19
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v458 = m.ExcPending
								if v458 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v461 = m.ExcPending
									if v461 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(402169), int32(0))
										mBase = m.M
										v465 = m.ExcPending
										if v465 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495940), int32(375), int32(375787))
											mBase = m.M
											v470 = m.ExcPending
											if v470 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(53983), v10)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495940), int32(418), int32(375787))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
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
			if l2 == int32(1184) {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
				v146 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if l3 != 0 {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v147
					v154 = base.I64_extend_i32_s(v147)*int64(-1000000) + v146
				} else {
					v154 = v146
				}
				if base.Ui64(v154-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
					F_EncodeSpecialTimestamp(m, v154, v19)
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 - int32(-64)
						return v19
					}
				} else {
					if l3 != 0 {
						v164 = int32(0)
					} else {
						v164 = v8 + int32(-48)
					}
					if l3 != 0 {
						v172 = int32(0)
					} else {
						v172 = v8 + int32(-56)
					}
					v174 = F_timestamp2tm(m, v154, v164, v8+int32(-44), v8+int32(-52), v172, int32(0))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						if v174 == int32(0) {
							if l3 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(1)
							} else {
							}
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							F_EncodeDateTime(m, v8+int32(-44), v182, int32(1), v184, v185, int32(4), v19)
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 - int32(-64)
								return v19
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(402169), int32(0))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495940), int32(414), int32(375787))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
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
			} else {
				if l2 != int32(1266) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(53983), v10)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495940), int32(418), int32(375787))
							mBase = m.M
							v219 = m.ExcPending
							if v219 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = v8 + int32(-44)
					v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					v32 = base.I64_div_s(v30, int64(3600000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v29)+8)) = uint32(v32)
					v37 = v30 + base.I64_extend32_s(v32)*int64(-3600000000)
					v39 = base.I64_div_s(v37, int64(60000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v29)+4)) = uint32(v39)
					v44 = base.I64_extend32_s(v39)*int64(-60000000) + v37
					v46 = base.I64_div_s(v44, int64(1000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v29))) = uint32(v46)
					v52 = v46*int64(4293967296) + v44
					*(*uint32)(unsafe.Add(mBase, uint32(v8+int32(-48)))) = uint32(v52)
					v55 = v8 + int32(-52)
					if v55 != 0 {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
					} else {
					}
					v59 = v8 + int32(-44)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v61 = int32(1)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
					v65 = int32(2)
					v66 = F_pg_ultostr_zeropad(m, v19, v64, v65)
					mBase = m.M
					v67 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v67)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
					v73 = F_pg_ultostr_zeropad(m, v66+v61, v71, v65)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v67)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v80 = F_AppendSeconds(m, v73+v61, v78, v60, v61)
					mBase = m.M
					v81 = F_EncodeTimezone(m, v80, v62, int32(4))
					mBase = m.M
					v83 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v83)
					m.G0 = v10 - int32(-64)
					return v19
				}
			}
		}
	}
}
func F_JsonTableResetNestedPlan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = l0
	goto L1
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(51) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v6 != int32(50) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+52))
	F_JsonTableResetNestedPlan(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L11
	}
L6:
	;
	return
L7:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
	if v12 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	F_JsonTableResetRowPattern(m, v3, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	goto L6
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v3)+56))
	v3 = v20
	goto L1
}
func F_JsonTableSetDocument(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = F_GetJsonTableExecContext(m, l0, int32(94719))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		F_JsonTableResetRowPattern(m, v6, l1)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_getJsonPathVariableFromJsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
	v16 = l0 + int32(4)
	v20 = F_findJsonbValueFromContainer(m, v16, int32(536870912), v9+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		if v20 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(-1)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v28
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(18)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v33 == v28 {
				v36 = int32(4)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v38&int32(254) == int32(2) {
					v47 = v36
				} else {
					v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
				}
				if v38 == int32(1) {
					v50 = v36
				} else {
					v50 = v47
				}
				v63 = v50
			} else {
				v51 = int32(1)
				if v33&v51 != 0 {
					v63 = int32(base.Ui32(v33)>>(uint(v51)%32)) - v51
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v63
		}
		m.G0 = v9 + int32(32)
		return v20
	}
}
func F_json_agg_strict_transfn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_json_agg_transfn_worker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_json_array_element(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v13
		v15 = int32(0)
		v20 = F_get_worker(m, v9, v15, v6+int32(12), int32(1), v15)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int32(0)
			} else {
				v27 = v20
			}
			m.G0 = v6 + int32(16)
			return v27
		}
	}
}
func F_json_build_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_extract_variadic_args(m, l0, v7+int32(12), v7+int32(4), v7+int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v29 = int32(0)
			m.G0 = v7 + int32(16)
			return v29
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v27 = F_json_build_array_worker(m, v15, v23, v24, v25, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				m.G0 = v7 + int32(16)
				return v29
			}
		}
	}
}
func F_json_errdetail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v7 = m.G0
	v9 = v7 - int32(208)
	m.G0 = v9
	v11 = int32(13904)
	if l0 == int32(16) {
		v182 = v11
		m.G0 = v9 + int32(208)
		return v182
	} else {
		if l1 == int32(4516584) {
			v182 = v11
			m.G0 = v9 + int32(208)
			return v182
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			if v16 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v18
				switch l0 - int32(2) {
				case 0:
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
					if v34 != 0 {
						v35 = int32(607332)
					} else {
						v35 = int32(607285)
					}
					v182 = v35
					m.G0 = v9 + int32(208)
					return v182
				case 1:
					v182 = int32(659793)
					m.G0 = v9 + int32(208)
					return v182
				case 2:
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v161
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v160 - v161
					F_appendStringInfo(m, v159, int32(645873), v9+int32(32))
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 3:
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v38
					F_appendStringInfo(m, v36, int32(649079), v9+int32(48))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 4:
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v58
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v57 - v58
					F_appendStringInfo(m, v56, int32(663795), v9+int32(80))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 5:
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v69
					*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v68 - v69
					F_appendStringInfo(m, v67, int32(663844), v9+int32(96))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 6:
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v80
					*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v79 - v80
					F_appendStringInfo(m, v78, int32(663883), v9+int32(112))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 7:
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v47
					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v46 - v47
					F_appendStringInfo(m, v45, int32(663599), v9-int32(-64))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 8:
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v91
					*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v90 - v91
					F_appendStringInfo(m, v89, int32(663675), v9+int32(128))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 9:
					v182 = int32(575392)
					m.G0 = v9 + int32(208)
					return v182
				case 10:
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v103
					*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v102 - v103
					F_appendStringInfo(m, v101, int32(663714), v9+int32(144))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 11:
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v114
					*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v113 - v114
					F_appendStringInfo(m, v112, int32(663756), v9+int32(160))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 12:
					v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v125
					*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v124 - v125
					F_appendStringInfo(m, v123, int32(663640), v9+int32(176))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				case 13:
					v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v136
					*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v135 - v136
					F_appendStringInfo(m, v134, int32(645909), v9+int32(192))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
						return int32(0)
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					}
				default:
					v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
					if v174 != 0 {
						v180 = v173
						v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
						v182 = v181
						m.G0 = v9 + int32(208)
						return v182
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_appendStringInfo(m, v173, int32(484720), v9)
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v180 = v179
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						}
					}
				case 15:
					v182 = int32(577796)
					m.G0 = v9 + int32(208)
					return v182
				case 16:
					v182 = int32(586182)
					m.G0 = v9 + int32(208)
					return v182
				case 17:
					v182 = int32(659480)
					m.G0 = v9 + int32(208)
					return v182
				case 18:
					v149 = *(*int32)(unsafe.Add(mBase, _consts[251]))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v150
					v155 = F_psprintf(m, int32(604807), v9+int32(16))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v182 = v155
						m.G0 = v9 + int32(208)
						return v182
					}
				case 19:
					v182 = int32(630873)
					m.G0 = v9 + int32(208)
					return v182
				case 20:
					v182 = int32(630821)
					m.G0 = v9 + int32(208)
					return v182
				}
			} else {
				v24 = F_makeStringInfo(m)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v24
					switch l0 - int32(2) {
					case 0:
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
						if v34 != 0 {
							v35 = int32(607332)
						} else {
							v35 = int32(607285)
						}
						v182 = v35
						m.G0 = v9 + int32(208)
						return v182
					case 1:
						v182 = int32(659793)
						m.G0 = v9 + int32(208)
						return v182
					case 2:
						v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v161
						*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v160 - v161
						F_appendStringInfo(m, v159, int32(645873), v9+int32(32))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 3:
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v38
						F_appendStringInfo(m, v36, int32(649079), v9+int32(48))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 4:
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v57 - v58
						F_appendStringInfo(m, v56, int32(663795), v9+int32(80))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 5:
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v68 - v69
						F_appendStringInfo(m, v67, int32(663844), v9+int32(96))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 6:
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v79 - v80
						F_appendStringInfo(m, v78, int32(663883), v9+int32(112))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 7:
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v46 - v47
						F_appendStringInfo(m, v45, int32(663599), v9-int32(-64))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 8:
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v91
						*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v90 - v91
						F_appendStringInfo(m, v89, int32(663675), v9+int32(128))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 9:
						v182 = int32(575392)
						m.G0 = v9 + int32(208)
						return v182
					case 10:
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v103
						*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v102 - v103
						F_appendStringInfo(m, v101, int32(663714), v9+int32(144))
						mBase = m.M
						v111 = m.ExcPending
						if v111 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 11:
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v114
						*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v113 - v114
						F_appendStringInfo(m, v112, int32(663756), v9+int32(160))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 12:
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v125
						*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v124 - v125
						F_appendStringInfo(m, v123, int32(663640), v9+int32(176))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					case 13:
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v136
						*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v135 - v136
						F_appendStringInfo(m, v134, int32(645909), v9+int32(192))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							if v174 != 0 {
								v180 = v173
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v173, int32(484720), v9)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
									v180 = v179
									v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
									v182 = v181
									m.G0 = v9 + int32(208)
									return v182
								}
							}
						}
					default:
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
						if v174 != 0 {
							v180 = v173
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
							v182 = v181
							m.G0 = v9 + int32(208)
							return v182
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v173, int32(484720), v9)
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
								v180 = v179
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
								v182 = v181
								m.G0 = v9 + int32(208)
								return v182
							}
						}
					case 15:
						v182 = int32(577796)
						m.G0 = v9 + int32(208)
						return v182
					case 16:
						v182 = int32(586182)
						m.G0 = v9 + int32(208)
						return v182
					case 17:
						v182 = int32(659480)
						m.G0 = v9 + int32(208)
						return v182
					case 18:
						v149 = *(*int32)(unsafe.Add(mBase, _consts[251]))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v150
						v155 = F_psprintf(m, int32(604807), v9+int32(16))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							v182 = v155
							m.G0 = v9 + int32(208)
							return v182
						}
					case 19:
						v182 = int32(630873)
						m.G0 = v9 + int32(208)
						return v182
					case 20:
						v182 = int32(630821)
						m.G0 = v9 + int32(208)
						return v182
					}
				}
			}
		}
	}
}
func F_json_manifest_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 - int32(5) {
	case 0:
		v22 = int32(6)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(82544)
		m.T0[v13].(func(*base.Module, int32, int32, int32))(m, v12, int32(199305), v6)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 4:
		v22 = int32(10)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_json_manifest_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v8 {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		m.G0 = v6 + int32(16)
		return int32(0)
	default:
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(82602)
		m.T0[v24].(func(*base.Module, int32, int32, int32))(m, v23, int32(199305), v6)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 6:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
		v13 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v13
		*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v13
		m.G0 = v6 + int32(16)
		return int32(0)
	case 10:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(11)
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_json_object_agg_unique_strict_transfn(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_json_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_object_field_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_text_to_cstring(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v16
				v22 = int32(1)
				v24 = F_get_worker(m, v9, v6+int32(12), int32(0), v22, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v31 = int32(0)
					} else {
						v31 = v24
					}
					m.G0 = v6 + int32(16)
					return v31
				}
			}
		}
	}
}
func F_json_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_text_to_cstring(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_json_to_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = l0 + int32(28)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = F_parse_jsonb_index_flags(m, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_getTSCurrentConfig(m)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v23
					v26 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v9 + int32(8)
					F_iterate_json_values(m, v12, v21, v9+int32(24))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v39 = F_make_tsvector(m, v9+int32(8))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v41 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									if v45 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(32)
											return v39
										}
									} else {
										m.G0 = v9 + int32(32)
										return v39
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								if v45 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(32)
										return v39
									}
								} else {
									m.G0 = v9 + int32(32)
									return v39
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_json_unique_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v4 == int32(1) {
		v8 = F_palloc(m, int32(8))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8
			return int32(0)
		}
	} else {
		return int32(0)
	}
}
func F_json_validate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = m.G0
	v7 = v5 - int32(192)
	m.G0 = v7
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7-int32(-64)))) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v11
	F_makeJsonLexContext(m, v7+int32(76), l0, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			v27 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v27)
			v29 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v7 + int32(76)
			v35 = v7 + int32(184)
			*(*int64)(unsafe.Add(mBase, uint32(v35))) = v29
			v39 = v7 + int32(168)
			*(*int64)(unsafe.Add(mBase, uint32(v39))) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v7)+176)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v7)+152)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v7)+160)) = int64(51539607564)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = v49
			*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(1321)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+144)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v7)+172)) = int32(1322)
			v62 = F_hash_create(m, int32(391616), int32(32), v7+int32(144), int32(1224))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v62
				*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = int32(1323)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(1324)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = int32(1325)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v7 + int32(12)
				v78 = F_pg_parse_json(m, v7+int32(76), v7+int32(32))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					if v78 != 0 {
						v89 = v78
						v91 = int32(0)
						if l2 == v91 {
							v123 = v91
							m.G0 = v7 + int32(192)
							return v123
						} else {
							F_json_errsave_error(m, v89, v7+int32(76), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								v123 = v91
								m.G0 = v7 + int32(192)
								return v123
							}
						}
					} else {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
						if v80 != 0 {
							F_freeJsonLexContext(m, v7+int32(76))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v123 = int32(1)
								m.G0 = v7 + int32(192)
								return v123
							}
						} else {
							if l2 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(786562))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(345200), int32(0))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495940), int32(1850), int32(356470))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
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
								v123 = int32(0)
								m.G0 = v7 + int32(192)
								return v123
							}
						}
					}
				}
			}
		} else {
			v85 = F_pg_parse_json(m, v7+int32(76), int32(1855488))
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				if v85 == int32(0) {
					v123 = int32(1)
					m.G0 = v7 + int32(192)
					return v123
				} else {
					v89 = v85
					v91 = int32(0)
					if l2 == v91 {
						v123 = v91
						m.G0 = v7 + int32(192)
						return v123
					} else {
						F_json_errsave_error(m, v89, v7+int32(76), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							v123 = v91
							m.G0 = v7 + int32(192)
							return v123
						}
					}
				}
			}
		}
	}
}
func F_transformJsonParseArg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_transformExprRecurse(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_exprType(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
			if v15 == int32(17) {
				v20 = F_exprLocation(m, v11)
				mBase = m.M
				v21 = F_getJsonEncodingConst(m, l2)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v11
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
					v31 = F_list_make2_impl(m, v9+int32(4), v9)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = int32(0)
						v35 = F_makeFuncExpr(m, int32(1714), int32(25), v31, v33, v33)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v20
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
							v40 = F_makeJsonValueExpr(m, v11, v35, l2)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v66 = v40
								m.G0 = v9 + int32(16)
								return v66
							}
						}
					}
				}
			} else {
				F_get_type_category_preferred(m, v15, v9+int32(12), v9+int32(8))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v48 != int32(705) {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
						if v51 != int32(83) {
							v63 = v11
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v64 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										F_parser_errposition(m, l0, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(162010), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494770), int32(4102), int32(326947))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
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
								v66 = v63
								m.G0 = v9 + int32(16)
								return v66
							}
						} else {
							v55 = int32(-1)
							v59 = F_coerce_to_target_type(m, l0, v11, v48, int32(25), v55, int32(0), int32(2), v55)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
								v63 = v59
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if v64 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											F_parser_errposition(m, l0, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(162010), int32(0))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494770), int32(4102), int32(326947))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
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
									v66 = v63
									m.G0 = v9 + int32(16)
									return v66
								}
							}
						}
					} else {
						v55 = int32(-1)
						v59 = F_coerce_to_target_type(m, l0, v11, v48, int32(25), v55, int32(0), int32(2), v55)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
							v63 = v59
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v64 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										F_parser_errposition(m, l0, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(162010), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494770), int32(4102), int32(326947))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
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
								v66 = v63
								m.G0 = v9 + int32(16)
								return v66
							}
						}
					}
				}
			}
		}
	}
}
