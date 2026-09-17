package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_make_timestamptz_at_timezone(m *base.Module, l0 int32) int32 {
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
	var v22 float64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
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
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int64
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	v12 = m.G0
	v14 = v12 - int32(352)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = F_pg_detoast_datum_packed(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = F_make_timestamp_internal(m, v20, v19, v18, v17, v16, v22)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = base.I64_div_s(v28, int64(86400000000))
			if base.Ui64(int64(172799999999)) <= base.Ui64(v28+int64(86399999999)) {
				v39 = v31 * int64(-86400000000)
			} else {
				v39 = int64(0)
			}
			v40 = v39 + v28
			v43 = v40>>(uint(int64(63))%64) + v31
			if int64(-2451546) < v43 {
				v46 = base.I32_wrap_i64(v43)
				v58 = v46 + int32(_a_F_make_timestamptz_at_timezone_0)
				v59 = int32(_a_F_make_timestamptz_at_timezone_1)
				v60 = base.I32_div_u_s(v58, v59)
				v61 = int32(3)
				v67 = int32(2)
				v72 = base.I32_div_u_s((v60*int32(1073595727)+v58)<<(uint(v67)%32)|v61, v59)
				v75 = v46 + int32(_a_F_make_timestamptz_at_timezone_2) + v60*v61 + v72 + int32(_a_F_make_timestamptz_at_timezone_3)
				v76 = int32(1461)
				v77 = base.I32_div_u_s(v75, v76)
				v80 = v77*int32(-1461) + v75
				v82 = v80 << (uint(v67) % 32)
				if base.Ui32(v76) <= base.Ui32(v82) {
					v88 = base.I32_rem_u_s(v80+int32(305), int32(365))
					v93 = v88
				} else {
					v92 = base.I32_rem_u_s(v80+int32(306), int32(366))
					v93 = v92
				}
				v95 = base.I32_div_u_s(v82, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(60)))) = v95 + v77<<(uint(int32(2))%32) - int32(_a_F_make_timestamptz_at_timezone_4)
				v103 = v93 + int32(123)
				v107 = int32(base.Ui32(v103*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(52)))) = v103 - int32(base.Ui32(v107*int32(_a_F_make_timestamptz_at_timezone_5))>>(uint(int32(8))%32))
				v117 = base.I32_rem_u_s(v107+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(56)))) = v117 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = int64(4294967295)
				if v40 < int64(0) {
					v129 = v40 + int64(86400000000)
				} else {
					v129 = v40
				}
				v131 = base.I64_div_s(v129, int64(3600000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+48)) = uint32(v131)
				v136 = base.I64_extend32_s(v131)*int64(-3600000000) + v129
				v138 = base.I64_div_s(v136, int64(60000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+44)) = uint32(v138)
				v145 = base.I64_div_s(base.I64_extend32_s(v138)*int64(-60000000)+v136, int64(1000000))
				*(*uint32)(unsafe.Add(mBase, uint32(v14)+40)) = uint32(v145)
				v148 = v14 + int32(96)
				F_text_to_cstring_buffer(m, v24, v148, int32(256))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+96)))
					if base.Ui32((v152-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v417 = m.ExcPending
						if v417 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v420 = m.ExcPending
							if v420 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_make_timestamptz_at_timezone_6)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v14 + int32(96)
								F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_7), v14)
								mBase = m.M
								v428 = m.ExcPending
								if v428 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_make_timestamptz_at_timezone_8), int32(0))
									mBase = m.M
									v432 = m.ExcPending
									if v432 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(514), int32(_a_F_make_timestamptz_at_timezone_10))
										mBase = m.M
										v437 = m.ExcPending
										if v437 != 0 {
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
						v159 = int32(0)
						v169 = m.G0
						v171 = v169 - int32(16)
						m.G0 = v171
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
						switch v174 - int32(43) {
						case 0, 2:
							v177 = int32(_a_F_make_timestamptz_at_timezone_11)
							*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
							v184 = F_strtoint(m, v14+int32(97), v171+int32(12))
							mBase = m.M
							v185 = int32(-5)
							v187 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
							if v187 == int32(68) {
								v264 = v185
							} else {
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
								v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
								if v191 != 0 {
									if v191 != int32(58) {
										v230 = int32(0)
										v231 = v184
										v233 = v159
										v237 = int32(59)
										if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v231))|base.B2i32(base.Ui32(v237) < base.Ui32(v230))|base.B2i32(base.Ui32(v237) < base.Ui32(v233)) != 0 {
											v264 = v185
										} else {
											v243 = int32(60)
											v248 = (v231*v243+v230)*v243 + v233
											v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
											if v251 == int32(45) {
												v254 = v248
											} else {
												v254 = int32(0) - v248
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
											v258 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
											v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
											if v259 != 0 {
												v260 = int32(-1)
											} else {
												v260 = int32(0)
											}
											v264 = v260
										}
									} else {
										v195 = int32(_a_F_make_timestamptz_at_timezone_11)
										*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
										v201 = v171 + int32(12)
										v202 = F_strtoint(m, v190+int32(1), v201)
										mBase = m.M
										v204 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
										if v204 == int32(68) {
											v264 = v185
										} else {
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
											v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
											if v208 != int32(58) {
												v230 = v202
												v231 = v184
												v233 = v159
												v237 = int32(59)
												if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v231))|base.B2i32(base.Ui32(v237) < base.Ui32(v230))|base.B2i32(base.Ui32(v237) < base.Ui32(v233)) != 0 {
													v264 = v185
												} else {
													v243 = int32(60)
													v248 = (v231*v243+v230)*v243 + v233
													v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
													if v251 == int32(45) {
														v254 = v248
													} else {
														v254 = int32(0) - v248
													}
													*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
													v258 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
													v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
													if v259 != 0 {
														v260 = int32(-1)
													} else {
														v260 = int32(0)
													}
													v264 = v260
												}
											} else {
												v211 = int32(_a_F_make_timestamptz_at_timezone_11)
												*(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0])) = int32(0)
												v216 = F_strtoint(m, v207+int32(1), v201)
												mBase = m.M
												v218 = *(*int32)(unsafe.Add(mBase, _c_F_make_timestamptz_at_timezone[0]))
												if v218 != int32(68) {
													v230 = v202
													v231 = v184
													v233 = v216
													v237 = int32(59)
													if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v231))|base.B2i32(base.Ui32(v237) < base.Ui32(v230))|base.B2i32(base.Ui32(v237) < base.Ui32(v233)) != 0 {
														v264 = v185
													} else {
														v243 = int32(60)
														v248 = (v231*v243+v230)*v243 + v233
														v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
														if v251 == int32(45) {
															v254 = v248
														} else {
															v254 = int32(0) - v248
														}
														*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
														v258 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
														v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
														if v259 != 0 {
															v260 = int32(-1)
														} else {
															v260 = int32(0)
														}
														v264 = v260
													}
												} else {
													v264 = v185
												}
											}
										}
									}
								} else {
									v221 = F_strlen(m, v148)
									mBase = m.M
									if base.Ui32(v221) < base.Ui32(int32(4)) {
										v230 = int32(0)
										v231 = v184
										v233 = v159
									} else {
										v225 = int32(100)
										v226 = base.I32_div_s(v184, v225)
										v230 = v184 - v226*v225
										v231 = v226
										v233 = v159
									}
									v237 = int32(59)
									if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v231))|base.B2i32(base.Ui32(v237) < base.Ui32(v230))|base.B2i32(base.Ui32(v237) < base.Ui32(v233)) != 0 {
										v264 = v185
									} else {
										v243 = int32(60)
										v248 = (v231*v243+v230)*v243 + v233
										v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
										if v251 == int32(45) {
											v254 = v248
										} else {
											v254 = int32(0) - v248
										}
										*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
										v258 = *(*int32)(unsafe.Add(mBase, uint32(v171)+12))
										v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
										if v259 != 0 {
											v260 = int32(-1)
										} else {
											v260 = int32(0)
										}
										v264 = v260
									}
								}
							}
						default:
							v264 = int32(-1)
						}
						m.G0 = v171 + int32(16)
						if v264 == int32(0) {
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
							v382 = v273
							v387 = base.I64_extend_i32_s(v159-v382)*int64(-1000000) + v28
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v387+int64(211813488000000000)) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v461 = m.ExcPending
								if v461 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v464 = m.ExcPending
									if v464 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
										mBase = m.M
										v468 = m.ExcPending
										if v468 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(716), int32(_a_F_make_timestamptz_at_timezone_13))
											mBase = m.M
											v473 = m.ExcPending
											if v473 != 0 {
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
								v392 = F_Int64GetDatum(m, v387)
								mBase = m.M
								v393 = m.ExcPending
								if v393 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(352)
									return v392
								}
							}
						} else {
							if v264 != int32(-1) {
								if v264 == int32(-5) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v441 = m.ExcPending
									if v441 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v444 = m.ExcPending
										if v444 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(96)
											F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_14), v14+int32(32))
											mBase = m.M
											v452 = m.ExcPending
											if v452 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(526), int32(_a_F_make_timestamptz_at_timezone_10))
												mBase = m.M
												v457 = m.ExcPending
												if v457 != 0 {
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
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v281 = m.ExcPending
									if v281 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v284 = m.ExcPending
										if v284 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(96)
											F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_15), v14+int32(16))
											mBase = m.M
											v292 = m.ExcPending
											if v292 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(530), int32(_a_F_make_timestamptz_at_timezone_10))
												mBase = m.M
												v297 = m.ExcPending
												if v297 != 0 {
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
								v304 = F_DecodeTimezoneName(m, v14+int32(96), v14+int32(88), v14+int32(84))
								mBase = m.M
								v305 = m.ExcPending
								if v305 != 0 {
									return int32(0)
								} else {
									switch v304 {
									case 0:
										v307 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
										v382 = int32(0) - v307
									case 1:
										v310 = v14 + int32(40)
										v313 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v318 = m.G0
										v320 = v318 - int32(288)
										m.G0 = v320
										v324 = F_DetermineTimeZoneOffsetInternal(m, v310, v313, v320+int32(280))
										mBase = m.M
										v326 = v320 + int32(16)
										v328 = F_strlcpy(m, v326, v14+int32(96), int32(256))
										mBase = m.M
										v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+16)))
										if v329 != 0 {
											v331 = v326
											v335 = v329
											for {
												v337 = F_pg_toupper(m, v335)
												mBase = m.M
												*(*uint8)(unsafe.Add(mBase, uint32(v331))) = uint8(v337)
												v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+1)))
												if v339 != 0 {
													v331 = v331 + int32(1)
													v335 = v339
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v357 = F_pg_interpret_timezone_abbrev(m, v320+int32(16), v320+int32(280), v320+int32(12), v320+int32(8), v313)
										mBase = m.M
										if v357 != 0 {
											v358 = *(*int32)(unsafe.Add(mBase, uint32(v320)+12))
											v359 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v310)+32)) = v359
											v364 = int32(0) - v358
										} else {
											v364 = v324
										}
										m.G0 = v320 + int32(288)
										v382 = v364
									default:
										v370 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v372 = m.G0
										v373 = int32(16)
										v374 = v372 - v373
										m.G0 = v374
										v378 = F_DetermineTimeZoneOffsetInternal(m, v14+int32(40), v370, v374+int32(8))
										mBase = m.M
										m.G0 = v374 + v373
										v382 = v378
									}
									v387 = base.I64_extend_i32_s(v159-v382)*int64(-1000000) + v28
									if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v387+int64(211813488000000000)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v461 = m.ExcPending
										if v461 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v464 = m.ExcPending
											if v464 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
												mBase = m.M
												v468 = m.ExcPending
												if v468 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(716), int32(_a_F_make_timestamptz_at_timezone_13))
													mBase = m.M
													v473 = m.ExcPending
													if v473 != 0 {
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
										v392 = F_Int64GetDatum(m, v387)
										mBase = m.M
										v393 = m.ExcPending
										if v393 != 0 {
											return int32(0)
										} else {
											m.G0 = v14 + int32(352)
											return v392
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
				v401 = m.ExcPending
				if v401 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v404 = m.ExcPending
					if v404 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_make_timestamptz_at_timezone_12), int32(0))
						mBase = m.M
						v408 = m.ExcPending
						if v408 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_make_timestamptz_at_timezone_9), int32(707), int32(_a_F_make_timestamptz_at_timezone_13))
							mBase = m.M
							v413 = m.ExcPending
							if v413 != 0 {
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
func F_timestamptz_ne_date(m *base.Module, l0 int32) int32 {
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
				F_j2date(m, v2+int32(_a_F_timestamptz_ne_date_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ne_date[0]))
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
	return base.B2i32(v65 != int32(0))
}
func F_timestamptz_ne_timestamp(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
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
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		m.G0 = v7 + int32(16)
		return base.B2i32(v19 != int32(0)) | base.B2i32(v10 != v15)
	}
}
func F_timestamptz_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
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
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v11
		if base.Ui64(v11-int64(9223372036854775807)) < base.Ui64(int64(2)) {
			F_AdjustTimestampForTypmod(m, v5+int32(-8), v9, int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
				v59 = F_Int64GetDatum(m, v58)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 - int32(-64)
					return v59
				}
			}
		} else {
			v26 = int32(0)
			v28 = F_timestamp2tm(m, v11, v5+int32(-12), v5+int32(-56), v5+int32(-60), v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if base.B2i32(v28 == int32(0))&base.B2i32(base.Ui64(v11+int64(211813488000000000)) <= base.Ui64(int64(-9011559254509551617))) != 0 {
					F_AdjustTimestampForTypmod(m, v5+int32(-8), v9, int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
						v59 = F_Int64GetDatum(m, v58)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 - int32(-64)
							return v59
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz_recv_0), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_timestamptz_recv_1), int32(827), int32(_a_F_timestamptz_recv_2))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
func F_timestamptz_trunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_trunc[0]))
		v11 = F_timestamptz_trunc_internal(m, v3, v8, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_Int64GetDatum(m, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
