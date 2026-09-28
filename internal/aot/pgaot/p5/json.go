package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonEncodeDateTime(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
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
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
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
	var v144 int32
	_ = v144
	var v150 int64
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if l0 == int32(0) {
		v14 = F_palloc(m, int32(129))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = v14
			if l2 <= int32(1183) {
				switch l2 - int32(1082) {
				case 0:
					v85 = base.I32_wrap_i64(l1)
					if base.Ui32(int32(1)) < base.Ui32(v85-int32(2147483647)) {
						v228 = v85 + int32(_a_F_JsonEncodeDateTime_0)
						v229 = int32(_a_F_JsonEncodeDateTime_1)
						v230 = base.I32_div_u_s(v228, v229)
						v231 = int32(3)
						v237 = int32(2)
						v242 = base.I32_div_u_s((v230*int32(1073595727)+v228)<<(uint(v237)%32)|v231, v229)
						v245 = v85 + int32(_a_F_JsonEncodeDateTime_2) + v230*v231 + v242 + int32(_a_F_JsonEncodeDateTime_3)
						v246 = int32(1461)
						v247 = base.I32_div_u_s(v245, v246)
						v250 = v247*int32(-1461) + v245
						v252 = v250 << (uint(v237) % 32)
						if base.Ui32(v246) <= base.Ui32(v252) {
							v258 = base.I32_rem_u_s(v250+int32(305), int32(365))
							v263 = v258
						} else {
							v262 = base.I32_rem_u_s(v250+int32(306), int32(366))
							v263 = v262
						}
						v265 = base.I32_div_u_s(v252, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-24)))) = v265 + v247<<(uint(int32(2))%32) - int32(_a_F_JsonEncodeDateTime_4)
						v273 = v263 + int32(123)
						v277 = int32(base.Ui32(v273*int32(2141)) >> (uint(int32(16)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-32)))) = v273 - int32(base.Ui32(v277*int32(_a_F_JsonEncodeDateTime_5))>>(uint(int32(8))%32))
						v287 = base.I32_rem_u_s(v277+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-28)))) = v287 + int32(1)
						v292 = v7 + int32(-44)
						switch int32(3) {
						case 0, 3:
							v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
							if int32(0) < v296 {
								v301 = v296
							} else {
								v301 = int32(1) - v296
							}
							v303 = F_pg_ultostr_zeropad(m, v18, v301, int32(4))
							mBase = m.M
							v304 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v303))) = uint8(v304)
							v306 = int32(1)
							v308 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
							v309 = int32(2)
							v310 = F_pg_ultostr_zeropad(m, v303+v306, v308, v309)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v304)
							v315 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
							v317 = F_pg_ultostr_zeropad(m, v310+v306, v315, v309)
							mBase = m.M
							v410 = v317
						case 1:
							v321 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
							v323 = base.B2i32(v321 == int32(1))
							if v321 == int32(1) {
								v324 = int32(12)
							} else {
								v324 = int32(16)
							}
							v326 = *(*int32)(unsafe.Add(mBase, uint32(v292+v324)))
							v328 = F_pg_ultostr_zeropad(m, v18, v326, int32(2))
							mBase = m.M
							v329 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v328))) = uint8(v329)
							if v321 == int32(1) {
								v335 = int32(16)
							} else {
								v335 = int32(12)
							}
							v337 = *(*int32)(unsafe.Add(mBase, uint32(v292+v335)))
							v339 = F_pg_ultostr_zeropad(m, v328+int32(1), v337, int32(2))
							mBase = m.M
							v340 = int32(47)
							*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v340)
							v342 = int32(1)
							v344 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
							if int32(0) < v344 {
								v349 = v344
							} else {
								v349 = v342 - v344
							}
							v351 = F_pg_ultostr_zeropad(m, v339+v342, v349, int32(4))
							mBase = m.M
							v410 = v351
						case 2:
							v352 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
							v353 = int32(2)
							v354 = F_pg_ultostr_zeropad(m, v18, v352, v353)
							mBase = m.M
							v355 = int32(46)
							*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v355)
							v357 = int32(1)
							v359 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
							v361 = F_pg_ultostr_zeropad(m, v354+v357, v359, v353)
							mBase = m.M
							*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v355)
							v366 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
							if int32(0) < v366 {
								v371 = v366
							} else {
								v371 = v357 - v366
							}
							v373 = F_pg_ultostr_zeropad(m, v361+v357, v371, int32(4))
							mBase = m.M
							v410 = v373
						default:
							v377 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
							v379 = base.B2i32(v377 == int32(1))
							if v377 == int32(1) {
								v380 = int32(12)
							} else {
								v380 = int32(16)
							}
							v382 = *(*int32)(unsafe.Add(mBase, uint32(v292+v380)))
							v384 = F_pg_ultostr_zeropad(m, v18, v382, int32(2))
							mBase = m.M
							v385 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v385)
							if v377 == int32(1) {
								v391 = int32(16)
							} else {
								v391 = int32(12)
							}
							v393 = *(*int32)(unsafe.Add(mBase, uint32(v292+v391)))
							v395 = F_pg_ultostr_zeropad(m, v384+int32(1), v393, int32(2))
							mBase = m.M
							v396 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v395))) = uint8(v396)
							v398 = int32(1)
							v400 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
							if int32(0) < v400 {
								v405 = v400
							} else {
								v405 = v398 - v400
							}
							v407 = F_pg_ultostr_zeropad(m, v395+v398, v405, int32(4))
							mBase = m.M
							v410 = v407
						}
						v411 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
						if v411 <= int32(0) {
							v415 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[1])))
							*(*uint8)(unsafe.Add(mBase, uint32(v410)+2)) = uint8(v415)
							v418 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[2])))
							*(*uint16)(unsafe.Add(mBase, uint32(v410))) = uint16(v418)
							v422 = v410 + int32(3)
						} else {
							v422 = v410
						}
						v423 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v422))) = uint8(v423)
						m.G0 = v9 - int32(-64)
						return v18
					} else {
						F_EncodeSpecialDate(m, v85, v18)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 - int32(-64)
							return v18
						}
					}
				case 1:
					v93 = v7 + int32(-44)
					v95 = base.I64_div_s(l1, int64(3600000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v93)+8)) = uint32(v95)
					v100 = base.I64_extend32_s(v95)*int64(-3600000000) + l1
					v102 = base.I64_div_s(v100, int64(60000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v93)+4)) = uint32(v102)
					v107 = base.I64_extend32_s(v102)*int64(-60000000) + v100
					v109 = base.I64_div_s(v107, int64(1000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v93))) = uint32(v109)
					v115 = v109*int64(4293967296) + v107
					*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(-48)))) = uint32(v115)
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
					v122 = int32(2)
					v123 = F_pg_ultostr_zeropad(m, v18, v121, v122)
					mBase = m.M
					v124 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
					v126 = int32(1)
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
					v130 = F_pg_ultostr_zeropad(m, v123+v126, v128, v122)
					mBase = m.M
					*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v124)
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
					v137 = F_AppendSeconds(m, v130+v126, v135, v117, v126)
					mBase = m.M
					v140 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
					m.G0 = v9 - int32(-64)
					return v18
				default:
					if l2 == int32(1114) {
						if base.Ui64(l1-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
							F_EncodeSpecialTimestamp(m, l1, v18)
							mBase = m.M
							v430 = m.ExcPending
							if v430 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 - int32(-64)
								return v18
							}
						} else {
							v431 = int32(0)
							v433 = v7 + int32(-44)
							v438 = F_timestamp2tm(m, l1, v431, v433, v7+int32(-48), v431, v431)
							mBase = m.M
							v439 = m.ExcPending
							if v439 != 0 {
								return int32(0)
							} else {
								if v438 == int32(0) {
									v442 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									v443 = int32(0)
									F_EncodeDateTime(m, v433, v442, v443, v443, v443, int32(4), v18)
									mBase = m.M
									v448 = m.ExcPending
									if v448 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 - int32(-64)
										return v18
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v452 = m.ExcPending
									if v452 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v455 = m.ExcPending
										if v455 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
											mBase = m.M
											v459 = m.ExcPending
											if v459 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(374), int32(_a_F_JsonEncodeDateTime_8))
												mBase = m.M
												v464 = m.ExcPending
												if v464 != 0 {
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
						v207 = m.ExcPending
						if v207 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v9)
							mBase = m.M
							v211 = m.ExcPending
							if v211 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(417), int32(_a_F_JsonEncodeDateTime_8))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
					if l3 != 0 {
						v144 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v144
						v150 = base.I64_extend_i32_s(v144)*int64(-1000000) + l1
					} else {
						v150 = l1
					}
					if base.Ui64(v150-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, v150, v18)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 - int32(-64)
							return v18
						}
					} else {
						if l3 != 0 {
							v161 = int32(0)
						} else {
							v161 = v7 + int32(-48)
						}
						if l3 != 0 {
							v169 = int32(0)
						} else {
							v169 = v7 + int32(-56)
						}
						v171 = F_timestamp2tm(m, v150, v161, v7+int32(-44), v7+int32(-52), v169, int32(0))
						mBase = m.M
						v172 = m.ExcPending
						if v172 != 0 {
							return int32(0)
						} else {
							if v171 == int32(0) {
								if l3 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(1)
								} else {
								}
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
								v181 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
								v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								F_EncodeDateTime(m, v7+int32(-44), v179, int32(1), v181, v182, int32(4), v18)
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 - int32(-64)
									return v18
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v192 = m.ExcPending
									if v192 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(413), int32(_a_F_JsonEncodeDateTime_8))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
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
						v207 = m.ExcPending
						if v207 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
							F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v9)
							mBase = m.M
							v211 = m.ExcPending
							if v211 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(417), int32(_a_F_JsonEncodeDateTime_8))
								mBase = m.M
								v216 = m.ExcPending
								if v216 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = v7 + int32(-44)
						v29 = base.I32_wrap_i64(l1)
						v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
						v32 = base.I64_div_s(v30, int64(3600000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)) = uint32(v32)
						v37 = v30 + base.I64_extend32_s(v32)*int64(-3600000000)
						v39 = base.I64_div_s(v37, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)) = uint32(v39)
						v44 = base.I64_extend32_s(v39)*int64(-60000000) + v37
						v46 = base.I64_div_s(v44, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v28))) = uint32(v46)
						v52 = v46*int64(4293967296) + v44
						*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(-48)))) = uint32(v52)
						v55 = v7 + int32(-52)
						if v55 != 0 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
						} else {
						}
						v59 = v7 + int32(-44)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
						v61 = int32(1)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
						v65 = int32(2)
						v66 = F_pg_ultostr_zeropad(m, v18, v64, v65)
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
						m.G0 = v9 - int32(-64)
						return v18
					}
				}
			}
		}
	} else {
		v18 = l0
		if l2 <= int32(1183) {
			switch l2 - int32(1082) {
			case 0:
				v85 = base.I32_wrap_i64(l1)
				if base.Ui32(int32(1)) < base.Ui32(v85-int32(2147483647)) {
					v228 = v85 + int32(_a_F_JsonEncodeDateTime_0)
					v229 = int32(_a_F_JsonEncodeDateTime_1)
					v230 = base.I32_div_u_s(v228, v229)
					v231 = int32(3)
					v237 = int32(2)
					v242 = base.I32_div_u_s((v230*int32(1073595727)+v228)<<(uint(v237)%32)|v231, v229)
					v245 = v85 + int32(_a_F_JsonEncodeDateTime_2) + v230*v231 + v242 + int32(_a_F_JsonEncodeDateTime_3)
					v246 = int32(1461)
					v247 = base.I32_div_u_s(v245, v246)
					v250 = v247*int32(-1461) + v245
					v252 = v250 << (uint(v237) % 32)
					if base.Ui32(v246) <= base.Ui32(v252) {
						v258 = base.I32_rem_u_s(v250+int32(305), int32(365))
						v263 = v258
					} else {
						v262 = base.I32_rem_u_s(v250+int32(306), int32(366))
						v263 = v262
					}
					v265 = base.I32_div_u_s(v252, int32(1461))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-24)))) = v265 + v247<<(uint(int32(2))%32) - int32(_a_F_JsonEncodeDateTime_4)
					v273 = v263 + int32(123)
					v277 = int32(base.Ui32(v273*int32(2141)) >> (uint(int32(16)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-32)))) = v273 - int32(base.Ui32(v277*int32(_a_F_JsonEncodeDateTime_5))>>(uint(int32(8))%32))
					v287 = base.I32_rem_u_s(v277+int32(10), int32(12))
					*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-28)))) = v287 + int32(1)
					v292 = v7 + int32(-44)
					switch int32(3) {
					case 0, 3:
						v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
						if int32(0) < v296 {
							v301 = v296
						} else {
							v301 = int32(1) - v296
						}
						v303 = F_pg_ultostr_zeropad(m, v18, v301, int32(4))
						mBase = m.M
						v304 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v303))) = uint8(v304)
						v306 = int32(1)
						v308 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
						v309 = int32(2)
						v310 = F_pg_ultostr_zeropad(m, v303+v306, v308, v309)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v304)
						v315 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
						v317 = F_pg_ultostr_zeropad(m, v310+v306, v315, v309)
						mBase = m.M
						v410 = v317
					case 1:
						v321 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
						v323 = base.B2i32(v321 == int32(1))
						if v321 == int32(1) {
							v324 = int32(12)
						} else {
							v324 = int32(16)
						}
						v326 = *(*int32)(unsafe.Add(mBase, uint32(v292+v324)))
						v328 = F_pg_ultostr_zeropad(m, v18, v326, int32(2))
						mBase = m.M
						v329 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v328))) = uint8(v329)
						if v321 == int32(1) {
							v335 = int32(16)
						} else {
							v335 = int32(12)
						}
						v337 = *(*int32)(unsafe.Add(mBase, uint32(v292+v335)))
						v339 = F_pg_ultostr_zeropad(m, v328+int32(1), v337, int32(2))
						mBase = m.M
						v340 = int32(47)
						*(*uint8)(unsafe.Add(mBase, uint32(v339))) = uint8(v340)
						v342 = int32(1)
						v344 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
						if int32(0) < v344 {
							v349 = v344
						} else {
							v349 = v342 - v344
						}
						v351 = F_pg_ultostr_zeropad(m, v339+v342, v349, int32(4))
						mBase = m.M
						v410 = v351
					case 2:
						v352 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
						v353 = int32(2)
						v354 = F_pg_ultostr_zeropad(m, v18, v352, v353)
						mBase = m.M
						v355 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v355)
						v357 = int32(1)
						v359 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
						v361 = F_pg_ultostr_zeropad(m, v354+v357, v359, v353)
						mBase = m.M
						*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v355)
						v366 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
						if int32(0) < v366 {
							v371 = v366
						} else {
							v371 = v357 - v366
						}
						v373 = F_pg_ultostr_zeropad(m, v361+v357, v371, int32(4))
						mBase = m.M
						v410 = v373
					default:
						v377 = *(*int32)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[0]))
						v379 = base.B2i32(v377 == int32(1))
						if v377 == int32(1) {
							v380 = int32(12)
						} else {
							v380 = int32(16)
						}
						v382 = *(*int32)(unsafe.Add(mBase, uint32(v292+v380)))
						v384 = F_pg_ultostr_zeropad(m, v18, v382, int32(2))
						mBase = m.M
						v385 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v385)
						if v377 == int32(1) {
							v391 = int32(16)
						} else {
							v391 = int32(12)
						}
						v393 = *(*int32)(unsafe.Add(mBase, uint32(v292+v391)))
						v395 = F_pg_ultostr_zeropad(m, v384+int32(1), v393, int32(2))
						mBase = m.M
						v396 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v395))) = uint8(v396)
						v398 = int32(1)
						v400 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
						if int32(0) < v400 {
							v405 = v400
						} else {
							v405 = v398 - v400
						}
						v407 = F_pg_ultostr_zeropad(m, v395+v398, v405, int32(4))
						mBase = m.M
						v410 = v407
					}
					v411 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
					if v411 <= int32(0) {
						v415 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[1])))
						*(*uint8)(unsafe.Add(mBase, uint32(v410)+2)) = uint8(v415)
						v418 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_JsonEncodeDateTime[2])))
						*(*uint16)(unsafe.Add(mBase, uint32(v410))) = uint16(v418)
						v422 = v410 + int32(3)
					} else {
						v422 = v410
					}
					v423 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v422))) = uint8(v423)
					m.G0 = v9 - int32(-64)
					return v18
				} else {
					F_EncodeSpecialDate(m, v85, v18)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 - int32(-64)
						return v18
					}
				}
			case 1:
				v93 = v7 + int32(-44)
				v95 = base.I64_div_s(l1, int64(3600000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v93)+8)) = uint32(v95)
				v100 = base.I64_extend32_s(v95)*int64(-3600000000) + l1
				v102 = base.I64_div_s(v100, int64(60000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v93)+4)) = uint32(v102)
				v107 = base.I64_extend32_s(v102)*int64(-60000000) + v100
				v109 = base.I64_div_s(v107, int64(1000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v93))) = uint32(v109)
				v115 = v109*int64(4293967296) + v107
				*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(-48)))) = uint32(v115)
				v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				v121 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
				v122 = int32(2)
				v123 = F_pg_ultostr_zeropad(m, v18, v121, v122)
				mBase = m.M
				v124 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
				v126 = int32(1)
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
				v130 = F_pg_ultostr_zeropad(m, v123+v126, v128, v122)
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v124)
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
				v137 = F_AppendSeconds(m, v130+v126, v135, v117, v126)
				mBase = m.M
				v140 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
				m.G0 = v9 - int32(-64)
				return v18
			default:
				if l2 == int32(1114) {
					if base.Ui64(l1-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
						F_EncodeSpecialTimestamp(m, l1, v18)
						mBase = m.M
						v430 = m.ExcPending
						if v430 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 - int32(-64)
							return v18
						}
					} else {
						v431 = int32(0)
						v433 = v7 + int32(-44)
						v438 = F_timestamp2tm(m, l1, v431, v433, v7+int32(-48), v431, v431)
						mBase = m.M
						v439 = m.ExcPending
						if v439 != 0 {
							return int32(0)
						} else {
							if v438 == int32(0) {
								v442 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
								v443 = int32(0)
								F_EncodeDateTime(m, v433, v442, v443, v443, v443, int32(4), v18)
								mBase = m.M
								v448 = m.ExcPending
								if v448 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 - int32(-64)
									return v18
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v452 = m.ExcPending
								if v452 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v455 = m.ExcPending
									if v455 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
										mBase = m.M
										v459 = m.ExcPending
										if v459 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(374), int32(_a_F_JsonEncodeDateTime_8))
											mBase = m.M
											v464 = m.ExcPending
											if v464 != 0 {
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
					v207 = m.ExcPending
					if v207 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v9)
						mBase = m.M
						v211 = m.ExcPending
						if v211 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(417), int32(_a_F_JsonEncodeDateTime_8))
							mBase = m.M
							v216 = m.ExcPending
							if v216 != 0 {
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
				*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(0)
				if l3 != 0 {
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v144
					v150 = base.I64_extend_i32_s(v144)*int64(-1000000) + l1
				} else {
					v150 = l1
				}
				if base.Ui64(v150-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
					F_EncodeSpecialTimestamp(m, v150, v18)
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 - int32(-64)
						return v18
					}
				} else {
					if l3 != 0 {
						v161 = int32(0)
					} else {
						v161 = v7 + int32(-48)
					}
					if l3 != 0 {
						v169 = int32(0)
					} else {
						v169 = v7 + int32(-56)
					}
					v171 = F_timestamp2tm(m, v150, v161, v7+int32(-44), v7+int32(-52), v169, int32(0))
					mBase = m.M
					v172 = m.ExcPending
					if v172 != 0 {
						return int32(0)
					} else {
						if v171 == int32(0) {
							if l3 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = int32(1)
							} else {
							}
							v179 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							v181 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
							F_EncodeDateTime(m, v7+int32(-44), v179, int32(1), v181, v182, int32(4), v18)
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 - int32(-64)
								return v18
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_JsonEncodeDateTime_6), int32(0))
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(413), int32(_a_F_JsonEncodeDateTime_8))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
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
					v207 = m.ExcPending
					if v207 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						F_errmsg_internal(m, int32(_a_F_JsonEncodeDateTime_9), v9)
						mBase = m.M
						v211 = m.ExcPending
						if v211 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonEncodeDateTime_7), int32(417), int32(_a_F_JsonEncodeDateTime_8))
							mBase = m.M
							v216 = m.ExcPending
							if v216 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = v7 + int32(-44)
					v29 = base.I32_wrap_i64(l1)
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
					v32 = base.I64_div_s(v30, int64(3600000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)) = uint32(v32)
					v37 = v30 + base.I64_extend32_s(v32)*int64(-3600000000)
					v39 = base.I64_div_s(v37, int64(60000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v28)+4)) = uint32(v39)
					v44 = base.I64_extend32_s(v39)*int64(-60000000) + v37
					v46 = base.I64_div_s(v44, int64(1000000))
					*(*uint32)(unsafe.Add(mBase, uint32(v28))) = uint32(v46)
					v52 = v46*int64(4293967296) + v44
					*(*uint32)(unsafe.Add(mBase, uint32(v7+int32(-48)))) = uint32(v52)
					v55 = v7 + int32(-52)
					if v55 != 0 {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
					} else {
					}
					v59 = v7 + int32(-44)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					v61 = int32(1)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
					v65 = int32(2)
					v66 = F_pg_ultostr_zeropad(m, v18, v64, v65)
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
					m.G0 = v9 - int32(-64)
					return v18
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
	var v13 int64
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v3)+128))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+136))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+112)))
	if v12 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v11)+104))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v3)+132))
	v3 = v20
	goto L1
}
func F_JsonTableSetDocument(m *base.Module, l0 int32, l1 int64) {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
	v16 = l0 + int32(4)
	v18 = F_findJsonbValueFromContainer(m, v16, int32(536870912), v9)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(-1)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = v26
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(18)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v31 == v26 {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v37 == int32(18) {
					v40 = int32(16)
				} else {
					v40 = int32(0)
				}
				if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v47 = int32(4)
				} else {
					v47 = v40
				}
				v60 = v47
			} else {
				v48 = int32(1)
				if v31&v48 != 0 {
					v60 = int32(base.Ui32(v31)>>(uint(v48)%32)) - v48
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v60
		}
		m.G0 = v9 + int32(32)
		return v18
	}
}
func F_json_agg_strict_transfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_json_agg_transfn_worker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_json_array_element(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14318(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_json_build_array(m *base.Module, l0 int32) int64 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_extract_variadic_args(m, l0, v7+int32(12), v7+int32(4), v7+int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 < int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v30 = int64(0)
			m.G0 = v7 + int32(16)
			return v30
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v28 = F_json_build_array_worker(m, v15, v24, v25, v26, int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = v28
				m.G0 = v7 + int32(16)
				return v30
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
func F_json_object_agg_unique_strict_transfn(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_json_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_json_object_field_text(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14319(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_json_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_text_to_cstring(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_json_to_tsvector(m *base.Module, l0 int32) int64 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = F_parse_jsonb_index_flags(m, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_getTSCurrentConfig(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
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
						return int64(0)
					} else {
						v35 = F_make_tsvector(m, v29)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v37 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v41 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(32)
											return base.I64_extend_i32_u(v35)
										}
									} else {
										m.G0 = v9 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v41 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								} else {
									m.G0 = v9 + int32(32)
									return base.I64_extend_i32_u(v35)
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
			*(*int64)(unsafe.Add(mBase, uint32(v8)+176)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+160)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+184)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+168)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+144)) = v28
			*(*int64)(unsafe.Add(mBase, uint32(v8)+152)) = int64(51539607564)
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_json_validate[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+180)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = int32(1446)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = int32(1447)
			v55 = F_hash_create(m, int32(_a_F_json_validate_0), int64(32), v8+int32(144), int32(1224))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(1448)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(1449)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1450)
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
												F_errfinish(m, int32(_a_F_json_validate_3), int32(1820), int32(_a_F_json_validate_4))
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
											F_errfinish(m, int32(_a_F_json_validate_3), int32(1820), int32(_a_F_json_validate_4))
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(32)
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
			v17 = F_getBaseType(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v17
				if v17 == int32(17) {
					v22 = F_exprLocation(m, v11)
					mBase = m.M
					v23 = F_getJsonEncodingConst(m, l2)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v11
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v23
						v35 = F_list_make2_impl(m, v9+int32(12), v9+int32(8))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = int32(0)
							v39 = F_makeFuncExpr(m, int32(1714), int32(25), v35, v37, v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v22
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
								v44 = F_makeJsonValueExpr(m, v11, v39, l2)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v75 = v44
									m.G0 = v9 + int32(32)
									return v75
								}
							}
						}
					}
				} else {
					F_get_type_category_preferred(m, v17, v9+int32(28), v9+int32(24))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if v52 != int32(705) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
							if v55 != int32(83) {
								v71 = v11
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if v73 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											F_parser_errposition(m, l0, v115)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
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
									v75 = v71
									m.G0 = v9 + int32(32)
									return v75
								}
							} else {
								v58 = F_exprLocation(m, v11)
								mBase = m.M
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v61 = int32(-1)
								v65 = F_coerce_to_target_type(m, l0, v11, v59, int32(25), v61, int32(0), int32(2), v61)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									if v65 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(101744772))
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
												v89 = F_format_type_be(m, v88)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return int32(0)
												} else {
													v92 = F_format_type_be(m, int32(25))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v92
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v89
														F_errmsg(m, int32(_a_F_transformJsonParseArg_4), v9+int32(16))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int32(0)
														} else {
															F_parser_errposition(m, l0, v58)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_5), int32(_a_F_transformJsonParseArg_3))
																mBase = m.M
																v107 = m.ExcPending
																if v107 != 0 {
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
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
										v71 = v65
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if v73 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													F_parser_errposition(m, l0, v115)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
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
											v75 = v71
											m.G0 = v9 + int32(32)
											return v75
										}
									}
								}
							}
						} else {
							v58 = F_exprLocation(m, v11)
							mBase = m.M
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v61 = int32(-1)
							v65 = F_coerce_to_target_type(m, l0, v11, v59, int32(25), v61, int32(0), int32(2), v61)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								if v65 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(101744772))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
											v89 = F_format_type_be(m, v88)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v92 = F_format_type_be(m, int32(25))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v92
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v89
													F_errmsg(m, int32(_a_F_transformJsonParseArg_4), v9+int32(16))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, v58)
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_5), int32(_a_F_transformJsonParseArg_3))
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
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
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(25)
									v71 = v65
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v73 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												F_parser_errposition(m, l0, v115)
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_transformJsonParseArg_0), int32(0))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_transformJsonParseArg_1), int32(_a_F_transformJsonParseArg_2), int32(_a_F_transformJsonParseArg_3))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
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
										v75 = v71
										m.G0 = v9 + int32(32)
										return v75
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
