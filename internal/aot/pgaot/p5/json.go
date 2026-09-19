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
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v152 int64
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int64
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
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
						v229 = l1 + int32(_a_F_JsonEncodeDateTime_0)
						v230 = int32(_a_F_JsonEncodeDateTime_1)
						v231 = base.I32_div_u_s(v229, v230)
						v232 = int32(3)
						v238 = int32(2)
						v243 = base.I32_div_u_s((v231*int32(1073595727)+v229)<<(uint(v238)%32)|v232, v230)
						v246 = l1 + int32(_a_F_JsonEncodeDateTime_2) + v231*v232 + v243 + int32(_a_F_JsonEncodeDateTime_3)
						v247 = int32(1461)
						v248 = base.I32_div_u_s(v246, v247)
						v251 = v248*int32(-1461) + v246
						v253 = v251 << (uint(v238) % 32)
						if base.Ui32(v247) <= base.Ui32(v253) {
							v259 = base.I32_rem_u_s(v251+int32(305), int32(365))
							v264 = v259
						} else {
							v263 = base.I32_rem_u_s(v251+int32(306), int32(366))
							v264 = v263
						}
						v266 = base.I32_div_u_s(v253, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-24)))) = v266 + v248<<(uint(int32(2))%32) - int32(_a_F_JsonEncodeDateTime_4)
						v274 = v264 + int32(123)
						v278 = int32(base.Ui32(v274*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-32)))) = v274 - int32(base.Ui32(v278*int32(_a_F_JsonEncodeDateTime_5))>>(uint(int32(8))%32))
						v288 = base.I32_rem_u_s(v278+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-28)))) = v288 + int32(1)
						v293 = v8 + int32(-44)
						switch int32(3) {
						case 0, 3:
							v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
							if int32(0) < v297 {
								v302 = v297
							} else {
								v302 = int32(1) - v297
							}
							v304 = F_pg_ultostr_zeropad(m, v19, v302, int32(4))
							mBase = m.M
							v305 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v305)
							v307 = int32(1)
							v309 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
							v310 = int32(2)
							v311 = F_pg_ultostr_zeropad(m, v304+v307, v309, v310)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v305)
							v316 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
							v318 = F_pg_ultostr_zeropad(m, v311+v307, v316, v310)
							mBase = m.M
							v411 = v318
						case 1:
							v322 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
							v324 = base.B2i32(v322 == int32(1))
							if v322 == int32(1) {
								v325 = int32(12)
							} else {
								v325 = int32(16)
							}
							v327 = *(*int32)(unsafe.Add(mBase, uint32(v293+v325)))
							v329 = F_pg_ultostr_zeropad(m, v19, v327, int32(2))
							mBase = m.M
							v330 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v330)
							if v322 == int32(1) {
								v336 = int32(16)
							} else {
								v336 = int32(12)
							}
							v338 = *(*int32)(unsafe.Add(mBase, uint32(v293+v336)))
							v340 = F_pg_ultostr_zeropad(m, v329+int32(1), v338, int32(2))
							mBase = m.M
							v341 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v341)
							v343 = int32(1)
							v345 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
							if int32(0) < v345 {
								v350 = v345
							} else {
								v350 = v343 - v345
							}
							v352 = F_pg_ultostr_zeropad(m, v340+v343, v350, int32(4))
							mBase = m.M
							v411 = v352
						case 2:
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
							v354 = int32(2)
							v355 = F_pg_ultostr_zeropad(m, v19, v353, v354)
							mBase = m.M
							v356 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v355))) = uint8(v356)
							v358 = int32(1)
							v360 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
							v362 = F_pg_ultostr_zeropad(m, v355+v358, v360, v354)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v362))) = uint8(v356)
							v367 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
							if int32(0) < v367 {
								v372 = v367
							} else {
								v372 = v358 - v367
							}
							v374 = F_pg_ultostr_zeropad(m, v362+v358, v372, int32(4))
							mBase = m.M
							v411 = v374
						default:
							v378 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
							v380 = base.B2i32(v378 == int32(1))
							if v378 == int32(1) {
								v381 = int32(12)
							} else {
								v381 = int32(16)
							}
							v383 = *(*int32)(unsafe.Add(mBase, uint32(v293+v381)))
							v385 = F_pg_ultostr_zeropad(m, v19, v383, int32(2))
							mBase = m.M
							v386 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v386)
							if v378 == int32(1) {
								v392 = int32(16)
							} else {
								v392 = int32(12)
							}
							v394 = *(*int32)(unsafe.Add(mBase, uint32(v293+v392)))
							v396 = F_pg_ultostr_zeropad(m, v385+int32(1), v394, int32(2))
							mBase = m.M
							v397 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v397)
							v399 = int32(1)
							v401 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
							if int32(0) < v401 {
								v406 = v401
							} else {
								v406 = v399 - v401
							}
							v408 = F_pg_ultostr_zeropad(m, v396+v399, v406, int32(4))
							mBase = m.M
							v411 = v408
						}
						v412 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
						if v412 <= int32(0) {
							v416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v411)+2)) = uint8(v416)
							v419 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v411))) = uint16(v419)
							v423 = v411 + int32(3)
						} else {
							v423 = v411
						}
						v424 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v423))) = uint8(v424)
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
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
					v122 = int32(2)
					v123 = F_pg_ultostr_zeropad(m, v19, v121, v122)
					mBase = m.M
					v124 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
					v126 = int32(1)
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
					v130 = F_pg_ultostr_zeropad(m, v123+v126, v128, v122)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v124)
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
					v137 = F_AppendSeconds(m, v130+v126, v135, v117, v126)
					mBase = m.M
					v140 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
					m.G0 = v10 - int32(-64)
					return v19
				default:
					if l2 == int32(1114) {
						v426 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
						if base.Ui64(v426-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
							F_EncodeSpecialTimestamp(m, v426, v19)
							mBase = m.M
							v432 = m.ExcPending
							if v432 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 - int32(-64)
								return v19
							}
						} else {
							v433 = int32(0)
							v435 = v8 + int32(-44)
							v440 = F_timestamp2tm(m, v426, v433, v435, v8+int32(-48), v433, v433)
							mBase = m.M
							v441 = m.ExcPending
							if v441 != 0 {
								return int32(0)
							} else {
								if v440 == int32(0) {
									v444 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v445 = int32(0)
									F_EncodeDateTime(m, v435, v444, v445, v445, v445, int32(4), v19)
									mBase = m.M
									v450 = m.ExcPending
									if v450 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 - int32(-64)
										return v19
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v454 = m.ExcPending
									if v454 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v457 = m.ExcPending
										if v457 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
											mBase = m.M
											v461 = m.ExcPending
											if v461 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(375), int32(_a_F_JsonEncodeDateTime_8))
												mBase = m.M
												v466 = m.ExcPending
												if v466 != 0 {
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
						v208 = m.ExcPending
						if v208 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
							F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v10)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(418), int32(_a_F_JsonEncodeDateTime_8))
								mBase = m.M
								v217 = m.ExcPending
								if v217 != 0 {
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
					v144 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					if l3 != 0 {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v145
						v152 = base.I64_extend_i32_s(v145)*int64(-1000000) + v144
					} else {
						v152 = v144
					}
					if base.Ui64(v152-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, v152, v19)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return v19
						}
					} else {
						if l3 != 0 {
							v162 = int32(0)
						} else {
							v162 = v8 + int32(-48)
						}
						if l3 != 0 {
							v170 = int32(0)
						} else {
							v170 = v8 + int32(-56)
						}
						v172 = F_timestamp2tm(m, v152, v162, v8+int32(-44), v8+int32(-52), v170, int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return int32(0)
						} else {
							if v172 == int32(0) {
								if l3 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(1)
								} else {
								}
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								F_EncodeDateTime(m, v8+int32(-44), v180, int32(1), v182, v183, int32(4), v19)
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 - int32(-64)
									return v19
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(414), int32(_a_F_JsonEncodeDateTime_8))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
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
						v208 = m.ExcPending
						if v208 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
							F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v10)
							mBase = m.M
							v212 = m.ExcPending
							if v212 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(418), int32(_a_F_JsonEncodeDateTime_8))
								mBase = m.M
								v217 = m.ExcPending
								if v217 != 0 {
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
					v229 = l1 + int32(_a_F_JsonEncodeDateTime_0)
					v230 = int32(_a_F_JsonEncodeDateTime_1)
					v231 = base.I32_div_u_s(v229, v230)
					v232 = int32(3)
					v238 = int32(2)
					v243 = base.I32_div_u_s((v231*int32(1073595727)+v229)<<(uint(v238)%32)|v232, v230)
					v246 = l1 + int32(_a_F_JsonEncodeDateTime_2) + v231*v232 + v243 + int32(_a_F_JsonEncodeDateTime_3)
					v247 = int32(1461)
					v248 = base.I32_div_u_s(v246, v247)
					v251 = v248*int32(-1461) + v246
					v253 = v251 << (uint(v238) % 32)
					if base.Ui32(v247) <= base.Ui32(v253) {
						v259 = base.I32_rem_u_s(v251+int32(305), int32(365))
						v264 = v259
					} else {
						v263 = base.I32_rem_u_s(v251+int32(306), int32(366))
						v264 = v263
					}
					v266 = base.I32_div_u_s(v253, int32(1461))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-24)))) = v266 + v248<<(uint(int32(2))%32) - int32(_a_F_JsonEncodeDateTime_4)
					v274 = v264 + int32(123)
					v278 = int32(base.Ui32(v274*int32(2141)) >> (uint(int32(16)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-32)))) = v274 - int32(base.Ui32(v278*int32(_a_F_JsonEncodeDateTime_5))>>(uint(int32(8))%32))
					v288 = base.I32_rem_u_s(v278+int32(10), int32(12))
					*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-28)))) = v288 + int32(1)
					v293 = v8 + int32(-44)
					switch int32(3) {
					case 0, 3:
						v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
						if int32(0) < v297 {
							v302 = v297
						} else {
							v302 = int32(1) - v297
						}
						v304 = F_pg_ultostr_zeropad(m, v19, v302, int32(4))
						mBase = m.M
						v305 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v305)
						v307 = int32(1)
						v309 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
						v310 = int32(2)
						v311 = F_pg_ultostr_zeropad(m, v304+v307, v309, v310)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v305)
						v316 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
						v318 = F_pg_ultostr_zeropad(m, v311+v307, v316, v310)
						mBase = m.M
						v411 = v318
					case 1:
						v322 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
						v324 = base.B2i32(v322 == int32(1))
						if v322 == int32(1) {
							v325 = int32(12)
						} else {
							v325 = int32(16)
						}
						v327 = *(*int32)(unsafe.Add(mBase, uint32(v293+v325)))
						v329 = F_pg_ultostr_zeropad(m, v19, v327, int32(2))
						mBase = m.M
						v330 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v330)
						if v322 == int32(1) {
							v336 = int32(16)
						} else {
							v336 = int32(12)
						}
						v338 = *(*int32)(unsafe.Add(mBase, uint32(v293+v336)))
						v340 = F_pg_ultostr_zeropad(m, v329+int32(1), v338, int32(2))
						mBase = m.M
						v341 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v341)
						v343 = int32(1)
						v345 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
						if int32(0) < v345 {
							v350 = v345
						} else {
							v350 = v343 - v345
						}
						v352 = F_pg_ultostr_zeropad(m, v340+v343, v350, int32(4))
						mBase = m.M
						v411 = v352
					case 2:
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
						v354 = int32(2)
						v355 = F_pg_ultostr_zeropad(m, v19, v353, v354)
						mBase = m.M
						v356 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v355))) = uint8(v356)
						v358 = int32(1)
						v360 = *(*int32)(unsafe.Add(mBase, uint32(v293)+16))
						v362 = F_pg_ultostr_zeropad(m, v355+v358, v360, v354)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v362))) = uint8(v356)
						v367 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
						if int32(0) < v367 {
							v372 = v367
						} else {
							v372 = v358 - v367
						}
						v374 = F_pg_ultostr_zeropad(m, v362+v358, v372, int32(4))
						mBase = m.M
						v411 = v374
					default:
						v378 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
						v380 = base.B2i32(v378 == int32(1))
						if v378 == int32(1) {
							v381 = int32(12)
						} else {
							v381 = int32(16)
						}
						v383 = *(*int32)(unsafe.Add(mBase, uint32(v293+v381)))
						v385 = F_pg_ultostr_zeropad(m, v19, v383, int32(2))
						mBase = m.M
						v386 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v385))) = uint8(v386)
						if v378 == int32(1) {
							v392 = int32(16)
						} else {
							v392 = int32(12)
						}
						v394 = *(*int32)(unsafe.Add(mBase, uint32(v293+v392)))
						v396 = F_pg_ultostr_zeropad(m, v385+int32(1), v394, int32(2))
						mBase = m.M
						v397 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v397)
						v399 = int32(1)
						v401 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
						if int32(0) < v401 {
							v406 = v401
						} else {
							v406 = v399 - v401
						}
						v408 = F_pg_ultostr_zeropad(m, v396+v399, v406, int32(4))
						mBase = m.M
						v411 = v408
					}
					v412 = *(*int32)(unsafe.Add(mBase, uint32(v293)+20))
					if v412 <= int32(0) {
						v416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[1])))
						*(*uint8)(unsafe.Add(mBase, uint32(v411)+2)) = uint8(v416)
						v419 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[2])))
						*(*uint16)(unsafe.Add(mBase, uint32(v411))) = uint16(v419)
						v423 = v411 + int32(3)
					} else {
						v423 = v411
					}
					v424 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v423))) = uint8(v424)
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
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				v121 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
				v122 = int32(2)
				v123 = F_pg_ultostr_zeropad(m, v19, v121, v122)
				mBase = m.M
				v124 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
				v126 = int32(1)
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
				v130 = F_pg_ultostr_zeropad(m, v123+v126, v128, v122)
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v124)
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				v137 = F_AppendSeconds(m, v130+v126, v135, v117, v126)
				mBase = m.M
				v140 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
				m.G0 = v10 - int32(-64)
				return v19
			default:
				if l2 == int32(1114) {
					v426 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
					if base.Ui64(v426-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, v426, v19)
						mBase = m.M
						v432 = m.ExcPending
						if v432 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 - int32(-64)
							return v19
						}
					} else {
						v433 = int32(0)
						v435 = v8 + int32(-44)
						v440 = F_timestamp2tm(m, v426, v433, v435, v8+int32(-48), v433, v433)
						mBase = m.M
						v441 = m.ExcPending
						if v441 != 0 {
							return int32(0)
						} else {
							if v440 == int32(0) {
								v444 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
								v445 = int32(0)
								F_EncodeDateTime(m, v435, v444, v445, v445, v445, int32(4), v19)
								mBase = m.M
								v450 = m.ExcPending
								if v450 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 - int32(-64)
									return v19
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v454 = m.ExcPending
								if v454 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v457 = m.ExcPending
									if v457 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
										mBase = m.M
										v461 = m.ExcPending
										if v461 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(375), int32(_a_F_JsonEncodeDateTime_8))
											mBase = m.M
											v466 = m.ExcPending
											if v466 != 0 {
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
					v208 = m.ExcPending
					if v208 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v10)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(418), int32(_a_F_JsonEncodeDateTime_8))
							mBase = m.M
							v217 = m.ExcPending
							if v217 != 0 {
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
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if l3 != 0 {
					v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v145
					v152 = base.I64_extend_i32_s(v145)*int64(-1000000) + v144
				} else {
					v152 = v144
				}
				if base.Ui64(v152-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
					F_EncodeSpecialTimestamp(m, v152, v19)
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 - int32(-64)
						return v19
					}
				} else {
					if l3 != 0 {
						v162 = int32(0)
					} else {
						v162 = v8 + int32(-48)
					}
					if l3 != 0 {
						v170 = int32(0)
					} else {
						v170 = v8 + int32(-56)
					}
					v172 = F_timestamp2tm(m, v152, v162, v8+int32(-44), v8+int32(-52), v170, int32(0))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						if v172 == int32(0) {
							if l3 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = int32(1)
							} else {
							}
							v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							F_EncodeDateTime(m, v8+int32(-44), v180, int32(1), v182, v183, int32(4), v19)
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 - int32(-64)
								return v19
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v190 = m.ExcPending
							if v190 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(414), int32(_a_F_JsonEncodeDateTime_8))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
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
					v208 = m.ExcPending
					if v208 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v10)
						mBase = m.M
						v212 = m.ExcPending
						if v212 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(418), int32(_a_F_JsonEncodeDateTime_8))
							mBase = m.M
							v217 = m.ExcPending
							if v217 != 0 {
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
	v4 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableSetDocument_0))
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v39 == int32(18) {
					v42 = int32(16)
				} else {
					v42 = int32(0)
				}
				if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v49 = int32(4)
				} else {
					v49 = v42
				}
				v62 = v49
			} else {
				v50 = int32(1)
				if v33&v50 != 0 {
					v62 = int32(base.Ui32(v33)>>(uint(v50)%32)) - v50
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v62 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v62
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13952(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
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
	if base.B2i32(l0 == int32(16))|base.B2i32(l1 == int32(_a_F_json_errdetail_0)) != 0 {
		v182 = int32(_a_F_json_errdetail_1)
		m.G0 = v9 + int32(208)
		return v182
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v19)
			*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v19
			switch l0 - int32(2) {
			case 0:
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
				if v35 != 0 {
					v36 = int32(_a_F_json_errdetail_2)
				} else {
					v36 = int32(_a_F_json_errdetail_3)
				}
				v182 = v36
				m.G0 = v9 + int32(208)
				return v182
			case 1:
				v182 = int32(_a_F_json_errdetail_4)
				m.G0 = v9 + int32(208)
				return v182
			case 2:
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v162
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v161 - v162
				F_appendStringInfo(m, v160, int32(_a_F_json_errdetail_5), v9+int32(32))
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v39
				F_appendStringInfo(m, v37, int32(_a_F_json_errdetail_7), v9+int32(48))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v58 - v59
				F_appendStringInfo(m, v57, int32(_a_F_json_errdetail_8), v9+int32(80))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v69 - v70
				F_appendStringInfo(m, v68, int32(_a_F_json_errdetail_9), v9+int32(96))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v81
				*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v80 - v81
				F_appendStringInfo(m, v79, int32(_a_F_json_errdetail_10), v9+int32(112))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v47 - v48
				F_appendStringInfo(m, v46, int32(_a_F_json_errdetail_11), v9-int32(-64))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v92
				*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v91 - v92
				F_appendStringInfo(m, v90, int32(_a_F_json_errdetail_12), v9+int32(128))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v182 = int32(_a_F_json_errdetail_13)
				m.G0 = v9 + int32(208)
				return v182
			case 10:
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v104
				*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v103 - v104
				F_appendStringInfo(m, v102, int32(_a_F_json_errdetail_14), v9+int32(144))
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v115
				*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v114 - v115
				F_appendStringInfo(m, v113, int32(_a_F_json_errdetail_15), v9+int32(160))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v126
				*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v125 - v126
				F_appendStringInfo(m, v124, int32(_a_F_json_errdetail_16), v9+int32(176))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v137
				*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v136 - v137
				F_appendStringInfo(m, v135, int32(_a_F_json_errdetail_17), v9+int32(192))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
				v182 = int32(_a_F_json_errdetail_18)
				m.G0 = v9 + int32(208)
				return v182
			case 16:
				v182 = int32(_a_F_json_errdetail_19)
				m.G0 = v9 + int32(208)
				return v182
			case 17:
				v182 = int32(_a_F_json_errdetail_20)
				m.G0 = v9 + int32(208)
				return v182
			case 18:
				v150 = *(*int32)(unsafe.Add(mBase, _c_F_json_errdetail[0]))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v151
				v156 = F_psprintf(m, int32(_a_F_json_errdetail_21), v9+int32(16))
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					v182 = v156
					m.G0 = v9 + int32(208)
					return v182
				}
			case 19:
				v182 = int32(_a_F_json_errdetail_22)
				m.G0 = v9 + int32(208)
				return v182
			case 20:
				v182 = int32(_a_F_json_errdetail_23)
				m.G0 = v9 + int32(208)
				return v182
			}
		} else {
			v25 = F_makeStringInfo(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v25
				switch l0 - int32(2) {
				case 0:
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
					if v35 != 0 {
						v36 = int32(_a_F_json_errdetail_2)
					} else {
						v36 = int32(_a_F_json_errdetail_3)
					}
					v182 = v36
					m.G0 = v9 + int32(208)
					return v182
				case 1:
					v182 = int32(_a_F_json_errdetail_4)
					m.G0 = v9 + int32(208)
					return v182
				case 2:
					v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v162
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v161 - v162
					F_appendStringInfo(m, v160, int32(_a_F_json_errdetail_5), v9+int32(32))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v39
					F_appendStringInfo(m, v37, int32(_a_F_json_errdetail_7), v9+int32(48))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v59
					*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v58 - v59
					F_appendStringInfo(m, v57, int32(_a_F_json_errdetail_8), v9+int32(80))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v69 - v70
					F_appendStringInfo(m, v68, int32(_a_F_json_errdetail_9), v9+int32(96))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v80 - v81
					F_appendStringInfo(m, v79, int32(_a_F_json_errdetail_10), v9+int32(112))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v47 - v48
					F_appendStringInfo(m, v46, int32(_a_F_json_errdetail_11), v9-int32(-64))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v92
					*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v91 - v92
					F_appendStringInfo(m, v90, int32(_a_F_json_errdetail_12), v9+int32(128))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v182 = int32(_a_F_json_errdetail_13)
					m.G0 = v9 + int32(208)
					return v182
				case 10:
					v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v104
					*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v103 - v104
					F_appendStringInfo(m, v102, int32(_a_F_json_errdetail_14), v9+int32(144))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v115
					*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v114 - v115
					F_appendStringInfo(m, v113, int32(_a_F_json_errdetail_15), v9+int32(160))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v126
					*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v125 - v126
					F_appendStringInfo(m, v124, int32(_a_F_json_errdetail_16), v9+int32(176))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v137
					*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v136 - v137
					F_appendStringInfo(m, v135, int32(_a_F_json_errdetail_17), v9+int32(192))
					mBase = m.M
					v145 = m.ExcPending
					if v145 != 0 {
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
							F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
						F_appendStringInfo(m, v173, int32(_a_F_json_errdetail_6), v9)
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
					v182 = int32(_a_F_json_errdetail_18)
					m.G0 = v9 + int32(208)
					return v182
				case 16:
					v182 = int32(_a_F_json_errdetail_19)
					m.G0 = v9 + int32(208)
					return v182
				case 17:
					v182 = int32(_a_F_json_errdetail_20)
					m.G0 = v9 + int32(208)
					return v182
				case 18:
					v150 = *(*int32)(unsafe.Add(mBase, _c_F_json_errdetail[0]))
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v151
					v156 = F_psprintf(m, int32(_a_F_json_errdetail_21), v9+int32(16))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int32(0)
					} else {
						v182 = v156
						m.G0 = v9 + int32(208)
						return v182
					}
				case 19:
					v182 = int32(_a_F_json_errdetail_22)
					m.G0 = v9 + int32(208)
					return v182
				case 20:
					v182 = int32(_a_F_json_errdetail_23)
					m.G0 = v9 + int32(208)
					return v182
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
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_json_manifest_array_start_0)
		m.T0[v13].(func(*base.Module, int32, int32, int32))(m, v12, int32(_a_F_json_manifest_array_start_1), v6)
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
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_json_manifest_object_start_0)
		m.T0[v24].(func(*base.Module, int32, int32, int32))(m, v23, int32(_a_F_json_manifest_object_start_1), v6)
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13953(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_parse_jsonb_index_flags(m, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_getTSCurrentConfig(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v21
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v24
					v29 = v9 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v29
					F_iterate_json_values(m, v12, v19, v9+int32(24))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = F_make_tsvector(m, v29)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v37 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v41 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(32)
											return v35
										}
									} else {
										m.G0 = v9 + int32(32)
										return v35
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v41 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(32)
										return v35
									}
								} else {
									m.G0 = v9 + int32(32)
									return v35
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v10
	v21 = v8 + int32(76)
	F_makeJsonLexContext(m, v21, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v26)
			v28 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v8)+184)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+168)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+152)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+144)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = int64(51539607564)
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_json_validate[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = int32(1305)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = int32(1306)
			v55 = F_hash_create(m, int32(_a_F_json_validate_0), int32(32), v8+int32(144), int32(1224))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(1307)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(1308)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1309)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(12)
				v70 = v8 + int32(32)
				v71 = F_pg_parse_json(m, v21, v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if v71 != 0 {
						v73 = int32(0)
						if l2 == v73 {
							v112 = v73
							m.G0 = v8 + int32(192)
							return v112
						} else {
							F_json_errsave_error(m, v71, v8+int32(76), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								v112 = v73
								m.G0 = v8 + int32(192)
								return v112
							}
						}
					} else {
						if l1 == int32(0) {
							v105 = int32(1)
							if l1 == int32(0) {
								v112 = v105
								m.G0 = v8 + int32(192)
								return v112
							} else {
								F_freeJsonLexContext(m, v8+int32(76))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = v105
									m.G0 = v8 + int32(192)
									return v112
								}
							}
						} else {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
							if v83&int32(1) != 0 {
								v105 = int32(1)
								if l1 == int32(0) {
									v112 = v105
									m.G0 = v8 + int32(192)
									return v112
								} else {
									F_freeJsonLexContext(m, v8+int32(76))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v112 = v105
										m.G0 = v8 + int32(192)
										return v112
									}
								}
							} else {
								if l2 == int32(0) {
									v112 = int32(0)
									m.G0 = v8 + int32(192)
									return v112
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_json_validate_1))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_json_validate_2), int32(0))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_json_validate_3), int32(1850), int32(_a_F_json_validate_4))
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
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
		} else {
			v70 = int32(_a_F_json_validate_5)
			v71 = F_pg_parse_json(m, v21, v70)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				if v71 != 0 {
					v73 = int32(0)
					if l2 == v73 {
						v112 = v73
						m.G0 = v8 + int32(192)
						return v112
					} else {
						F_json_errsave_error(m, v71, v8+int32(76), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v112 = v73
							m.G0 = v8 + int32(192)
							return v112
						}
					}
				} else {
					if l1 == int32(0) {
						v105 = int32(1)
						if l1 == int32(0) {
							v112 = v105
							m.G0 = v8 + int32(192)
							return v112
						} else {
							F_freeJsonLexContext(m, v8+int32(76))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = v105
								m.G0 = v8 + int32(192)
								return v112
							}
						}
					} else {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
						if v83&int32(1) != 0 {
							v105 = int32(1)
							if l1 == int32(0) {
								v112 = v105
								m.G0 = v8 + int32(192)
								return v112
							} else {
								F_freeJsonLexContext(m, v8+int32(76))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = v105
									m.G0 = v8 + int32(192)
									return v112
								}
							}
						} else {
							if l2 == int32(0) {
								v112 = int32(0)
								m.G0 = v8 + int32(192)
								return v112
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(_a_F_json_validate_1))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_json_validate_2), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_json_validate_3), int32(1850), int32(_a_F_json_validate_4))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
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
											F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
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
												F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
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
											F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
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
