package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamptz_at_timezone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 float64
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
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int64
	_ = v386
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
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
	v12 = m.G0
	v14 = v12 - int32(352)
	m.G0 = v14
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l0)+104))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v23 = F_pg_detoast_datum_packed(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v27 = F_make_timestamp_internal(m, v21, v20, v19, v18, v17, v16)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v30 = base.I64_div_s(v27, int64(86400000000))
			if base.Ui64(int64(172799999999)) <= base.Ui64(v27+int64(86399999999)) {
				v38 = v30 * int64(-86400000000)
			} else {
				v38 = int64(0)
			}
			v39 = v38 + v27
			v42 = v39>>(uint(int64(63))%64) + v30
			if int64(-2451546) < v42 {
				v45 = base.I32_wrap_i64(v42)
				v57 = v45 + int32(_a_F_make_timestamptz_at_timezone_0)
				v58 = int32(_a_F_make_timestamptz_at_timezone_1)
				v59 = base.I32_div_u_s(v57, v58)
				v60 = int32(3)
				v66 = int32(2)
				v71 = base.I32_div_u_s((v59*int32(1073595727)+v57)<<(uint(v66)%32)|v60, v58)
				v74 = v45 + int32(_a_F_make_timestamptz_at_timezone_2) + v59*v60 + v71 + int32(_a_F_make_timestamptz_at_timezone_3)
				v75 = int32(1461)
				v76 = base.I32_div_u_s(v74, v75)
				v79 = v76*int32(-1461) + v74
				v81 = v79 << (uint(v66) % 32)
				if base.Ui32(v75) <= base.Ui32(v81) {
					v87 = base.I32_rem_u_s(v79+int32(305), int32(365))
					v92 = v87
				} else {
					v91 = base.I32_rem_u_s(v79+int32(306), int32(366))
					v92 = v91
				}
				v94 = base.I32_div_u_s(v81, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(60)))) = v94 + v76<<(uint(int32(2))%32) - int32(_a_F_make_timestamptz_at_timezone_4)
				v102 = v92 + int32(123)
				v106 = int32(base.Ui32(v102*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(52)))) = v102 - int32(base.Ui32(v106*int32(_a_F_make_timestamptz_at_timezone_5))>>(uint(int32(8))%32))
				v116 = base.I32_rem_u_s(v106+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(56)))) = v116 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = int64(4294967295)
				if v39 < int64(0) {
					v128 = v39 + int64(86400000000)
				} else {
					v128 = v39
				}
				v130 = base.I64_div_s(v128, int64(3600000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+48)) = uint32(v130)
				v135 = base.I64_extend32_s(v130)*int64(-3600000000) + v128
				v137 = base.I64_div_s(v135, int64(60000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+44)) = uint32(v137)
				v144 = base.I64_div_s(base.I64_extend32_s(v137)*int64(-60000000)+v135, int64(1000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+40)) = uint32(v144)
				v147 = v14 + int32(96)
				F_text_to_cstring_buffer(m, v23, v147, int32(256))
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int64(0)
				} else {
					v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+96)))
					if base.Ui32((v151-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v414 = m.ExcPending
						if v414 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_make_timestamptz_at_timezone_6)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v14 + int32(96)
								F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_7), v14)
								mBase = m.M
								v425 = m.ExcPending
								if v425 != 0 {
									return int64(0)
								} else {
									F_errhint(m, int32(_a_F_make_timestamptz_at_timezone_8), int32(0))
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(510), int32(_a_F_make_timestamptz_at_timezone_10))
										mBase = m.M
										v434 = m.ExcPending
										if v434 != 0 {
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
					} else {
						v158 = int32(0)
						v168 = m.G0
						v170 = v168 - int32(16)
						m.G0 = v170
						v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
						switch v173 - int32(43) {
						case 0, 2:
							v176 = int32(_a_F_make_timestamptz_at_timezone_11)
							*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
							v183 = F_strtoint(m, v14+int32(97), v170+int32(12))
							mBase = m.M
							v184 = int32(-5)
							v186 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
							if v186 == int32(68) {
								v263 = v184
							} else {
								v189 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
								v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
								if v190 != 0 {
									if v190 != int32(58) {
										v229 = int32(0)
										v230 = v183
										v232 = v158
										v236 = int32(59)
										if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v230))|base.B2i32(base.Ui32(v236) < base.Ui32(v229))|base.B2i32(base.Ui32(v236) < base.Ui32(v232)) != 0 {
											v263 = v184
										} else {
											v242 = int32(60)
											v247 = (v230*v242+v229)*v242 + v232
											v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
											if v250 == int32(45) {
												v253 = v247
											} else {
												v253 = int32(0) - v247
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v253
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
											v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
											if v258 != 0 {
												v259 = int32(-1)
											} else {
												v259 = int32(0)
											}
											v263 = v259
										}
									} else {
										v194 = int32(_a_F_make_timestamptz_at_timezone_11)
										*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
										v200 = v170 + int32(12)
										v201 = F_strtoint(m, v189+int32(1), v200)
										mBase = m.M
										v203 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
										if v203 == int32(68) {
											v263 = v184
										} else {
											v206 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
											v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
											if v207 != int32(58) {
												v229 = v201
												v230 = v183
												v232 = v158
												v236 = int32(59)
												if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v230))|base.B2i32(base.Ui32(v236) < base.Ui32(v229))|base.B2i32(base.Ui32(v236) < base.Ui32(v232)) != 0 {
													v263 = v184
												} else {
													v242 = int32(60)
													v247 = (v230*v242+v229)*v242 + v232
													v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
													if v250 == int32(45) {
														v253 = v247
													} else {
														v253 = int32(0) - v247
													}
													*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v253
													v257 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
													v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
													if v258 != 0 {
														v259 = int32(-1)
													} else {
														v259 = int32(0)
													}
													v263 = v259
												}
											} else {
												v210 = int32(_a_F_make_timestamptz_at_timezone_11)
												*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
												v215 = F_strtoint(m, v206+int32(1), v200)
												mBase = m.M
												v217 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
												if v217 != int32(68) {
													v229 = v201
													v230 = v183
													v232 = v215
													v236 = int32(59)
													if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v230))|base.B2i32(base.Ui32(v236) < base.Ui32(v229))|base.B2i32(base.Ui32(v236) < base.Ui32(v232)) != 0 {
														v263 = v184
													} else {
														v242 = int32(60)
														v247 = (v230*v242+v229)*v242 + v232
														v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
														if v250 == int32(45) {
															v253 = v247
														} else {
															v253 = int32(0) - v247
														}
														*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v253
														v257 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
														v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
														if v258 != 0 {
															v259 = int32(-1)
														} else {
															v259 = int32(0)
														}
														v263 = v259
													}
												} else {
													v263 = v184
												}
											}
										}
									}
								} else {
									v220 = F_strlen(m, v147)
									mBase = m.M
									if base.Ui32(v220) < base.Ui32(int32(4)) {
										v229 = int32(0)
										v230 = v183
										v232 = v158
									} else {
										v224 = int32(100)
										v225 = base.I32_div_s(v183, v224)
										v229 = v183 - v225*v224
										v230 = v225
										v232 = v158
									}
									v236 = int32(59)
									if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v230))|base.B2i32(base.Ui32(v236) < base.Ui32(v229))|base.B2i32(base.Ui32(v236) < base.Ui32(v232)) != 0 {
										v263 = v184
									} else {
										v242 = int32(60)
										v247 = (v230*v242+v229)*v242 + v232
										v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
										if v250 == int32(45) {
											v253 = v247
										} else {
											v253 = int32(0) - v247
										}
										*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v253
										v257 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
										v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
										if v258 != 0 {
											v259 = int32(-1)
										} else {
											v259 = int32(0)
										}
										v263 = v259
									}
								}
							}
						default:
							v263 = int32(-1)
						}
						m.G0 = v170 + int32(16)
						if v263 == int32(0) {
							v272 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
							v381 = v272
							v386 = base.I64_extend_i32_s(v158-v381)*int64(-1000000) + v27
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v386+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v458 = m.ExcPending
								if v458 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v461 = m.ExcPending
									if v461 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
										mBase = m.M
										v465 = m.ExcPending
										if v465 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(712), int32(_a_F_make_timestamptz_at_timezone_13))
											mBase = m.M
											v470 = m.ExcPending
											if v470 != 0 {
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
								m.G0 = v14 + int32(352)
								return v386
							}
						} else {
							if v263 != int32(-1) {
								if v263 == int32(-5) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v438 = m.ExcPending
									if v438 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v441 = m.ExcPending
										if v441 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(96)
											F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_14), v14+int32(32))
											mBase = m.M
											v449 = m.ExcPending
											if v449 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(522), int32(_a_F_make_timestamptz_at_timezone_10))
												mBase = m.M
												v454 = m.ExcPending
												if v454 != 0 {
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
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v280 = m.ExcPending
									if v280 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v283 = m.ExcPending
										if v283 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(96)
											F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_15), v14+int32(16))
											mBase = m.M
											v291 = m.ExcPending
											if v291 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(526), int32(_a_F_make_timestamptz_at_timezone_10))
												mBase = m.M
												v296 = m.ExcPending
												if v296 != 0 {
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
							} else {
								v303 = F_DecodeTimezoneName(m, v14+int32(96), v14+int32(88), v14+int32(84))
								mBase = m.M
								v304 = m.ExcPending
								if v304 != 0 {
									return int64(0)
								} else {
									switch v303 {
									case 0:
										v306 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
										v381 = int32(0) - v306
									case 1:
										v309 = v14 + int32(40)
										v312 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v317 = m.G0
										v319 = v317 - int32(288)
										m.G0 = v319
										v323 = F_DetermineTimeZoneOffsetInternal(m, v309, v312, v319+int32(280))
										mBase = m.M
										v325 = v319 + int32(16)
										v327 = F_strlcpy(m, v325, v14+int32(96), int32(256))
										mBase = m.M
										v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+16)))
										if v328 != 0 {
											v330 = v325
											v334 = v328
											for {
												v336 = F_pg_toupper(m, v334)
												mBase = m.M
												*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v336)
												v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
												if v338 != 0 {
													v330 = v330 + int32(1)
													v334 = v338
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v356 = F_pg_interpret_timezone_abbrev(m, v319+int32(16), v319+int32(280), v319+int32(12), v319+int32(8), v312)
										mBase = m.M
										if v356 != 0 {
											v357 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
											v358 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v309)+32)) = v358
											v363 = int32(0) - v357
										} else {
											v363 = v323
										}
										m.G0 = v319 + int32(288)
										v381 = v363
									default:
										v369 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v371 = m.G0
										v372 = int32(16)
										v373 = v371 - v372
										m.G0 = v373
										v377 = F_DetermineTimeZoneOffsetInternal(m, v14+int32(40), v369, v373+int32(8))
										mBase = m.M
										m.G0 = v373 + v372
										v381 = v377
									}
									v386 = base.I64_extend_i32_s(v158-v381)*int64(-1000000) + v27
									if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v386+int64(211813488000000000)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v458 = m.ExcPending
										if v458 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v461 = m.ExcPending
											if v461 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
												mBase = m.M
												v465 = m.ExcPending
												if v465 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(712), int32(_a_F_make_timestamptz_at_timezone_13))
													mBase = m.M
													v470 = m.ExcPending
													if v470 != 0 {
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
										m.G0 = v14 + int32(352)
										return v386
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v398 = m.ExcPending
				if v398 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v401 = m.ExcPending
					if v401 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
						mBase = m.M
						v405 = m.ExcPending
						if v405 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(703), int32(_a_F_make_timestamptz_at_timezone_13))
							mBase = m.M
							v410 = m.ExcPending
							if v410 != 0 {
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
func F_timestamptz_ne_date(m *base.Module, l0 int32) int64 {
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
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ne_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_ne_date[1]))
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
				v35 = int64(1)
			} else {
				v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) != int32(0)))
			}
		} else {
			v35 = base.I64_extend_i32_u(base.B2i32(base.B2i32(v9 < v17)-base.B2i32(v17 < v9) != int32(0)))
		}
		m.G0 = v7 + int32(16)
		return v35
	}
}
func F_timestamptz_ne_timestamp(m *base.Module, l0 int32) int64 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ne_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_ne_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v9, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v21&base.B2i32(base.Ui64(v17-int64(9223372036854775807)) < base.Ui64(int64(2))) | base.B2i32(v17 != v10))
	}
}
func F_timestamptz_recv(m *base.Module, l0 int32) int64 {
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
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v11
		if base.Ui64(v11-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			v57 = F_AdjustTimestampForTypmod(m, v5+int32(-8), base.I32_wrap_i64(v9), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
				m.G0 = v7 - int32(-64)
				return v59
			}
		} else {
			v26 = int32(0)
			v28 = F_timestamp2tm(m, v11, v5+int32(-12), v5+int32(-56), v5+int32(-60), v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				if base.B2i32(v28 == int32(0))&base.B2i32(base.Ui64(v11+int64(211813488000000000)) <= base.Ui64(int64(-9011559254509551617))) != 0 {
					v57 = F_AdjustTimestampForTypmod(m, v5+int32(-8), base.I32_wrap_i64(v9), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
						m.G0 = v7 - int32(-64)
						return v59
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_recv_0), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamptz_recv_1), int32(824), int32(_a_F_timestamptz_recv_2))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
func F_timestamptz_trunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_trunc[0]))
		v10 = F_timestamptz_trunc_internal(m, v3, v7, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return v10
		}
	}
}
