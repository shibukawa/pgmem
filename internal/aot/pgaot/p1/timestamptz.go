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
	var v21 float64
	_ = v21
	var v22 int32
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
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
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int64
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = F_pg_detoast_datum_packed(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = F_make_timestamp_internal(m, v22, v19, v18, v17, v16, v21)
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
				v58 = v46 + int32(2483589)
				v59 = int32(146097)
				v60 = base.I32_div_u_s(v58, v59)
				v61 = int32(3)
				v67 = int32(2)
				v72 = base.I32_div_u_s((v60*int32(1073595727)+v58)<<(uint(v67)%32)|v61, v59)
				v75 = v46 + int32(2451545) + v60*v61 + v72 + int32(32104)
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
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(60)))) = v95 + v77<<(uint(int32(2))%32) - int32(4800)
				v103 = v93 + int32(123)
				v107 = int32(base.Ui32(v103*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v14+int32(52)))) = v103 - int32(base.Ui32(v107*int32(7834))>>(uint(int32(8))%32))
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
				F_text_to_cstring_buffer(m, v24, v14+int32(96), int32(256))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+96)))
					if base.Ui32((v152-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v414 = m.ExcPending
						if v414 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(370290)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v14 + int32(96)
								F_errmsg(m, int32(706670), v14)
								mBase = m.M
								v425 = m.ExcPending
								if v425 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(590797), int32(0))
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(492756), int32(514), int32(369972))
										mBase = m.M
										v434 = m.ExcPending
										if v434 != 0 {
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
						v161 = v14 + int32(96)
						v170 = m.G0
						v172 = v170 - int32(16)
						m.G0 = v172
						v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
						switch v175 - int32(43) {
						case 0, 2:
							v178 = int32(4651076)
							*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
							v185 = F_strtoint(m, v14+int32(97), v172+int32(12))
							mBase = m.M
							v186 = int32(-5)
							v188 = *(*int32)(unsafe.Add(mBase, _consts[137]))
							if v188 == int32(68) {
								v262 = v186
							} else {
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
								v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
								if v192 != 0 {
									if v192 != int32(58) {
										v233 = int32(0)
										v234 = v185
										v236 = v159
										if base.Ui32(int32(15)) < base.Ui32(v234) {
											v262 = v186
										} else {
											if base.Ui32(int32(59)) < base.Ui32(v233) {
												v262 = v186
											} else {
												if base.Ui32(int32(59)) < base.Ui32(v236) {
													v262 = v186
												} else {
													v243 = int32(60)
													v248 = (v234*v243+v233)*v243 + v236
													v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
													if v251 == int32(45) {
														v254 = v248
													} else {
														v254 = int32(0) - v248
													}
													*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
													v258 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
													v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
													if v259 != 0 {
														v260 = int32(-1)
													} else {
														v260 = int32(0)
													}
													v262 = v260
												}
											}
										}
									} else {
										v196 = int32(4651076)
										*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
										v203 = F_strtoint(m, v191+int32(1), v172+int32(12))
										mBase = m.M
										v205 = *(*int32)(unsafe.Add(mBase, _consts[137]))
										if v205 == int32(68) {
											v262 = v186
										} else {
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
											v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
											if v209 != int32(58) {
												v233 = v203
												v234 = v185
												v236 = v159
												if base.Ui32(int32(15)) < base.Ui32(v234) {
													v262 = v186
												} else {
													if base.Ui32(int32(59)) < base.Ui32(v233) {
														v262 = v186
													} else {
														if base.Ui32(int32(59)) < base.Ui32(v236) {
															v262 = v186
														} else {
															v243 = int32(60)
															v248 = (v234*v243+v233)*v243 + v236
															v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
															if v251 == int32(45) {
																v254 = v248
															} else {
																v254 = int32(0) - v248
															}
															*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
															v258 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
															v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
															if v259 != 0 {
																v260 = int32(-1)
															} else {
																v260 = int32(0)
															}
															v262 = v260
														}
													}
												}
											} else {
												v212 = int32(4651076)
												*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
												v219 = F_strtoint(m, v208+int32(1), v172+int32(12))
												mBase = m.M
												v221 = *(*int32)(unsafe.Add(mBase, _consts[137]))
												if v221 != int32(68) {
													v233 = v203
													v234 = v185
													v236 = v219
													if base.Ui32(int32(15)) < base.Ui32(v234) {
														v262 = v186
													} else {
														if base.Ui32(int32(59)) < base.Ui32(v233) {
															v262 = v186
														} else {
															if base.Ui32(int32(59)) < base.Ui32(v236) {
																v262 = v186
															} else {
																v243 = int32(60)
																v248 = (v234*v243+v233)*v243 + v236
																v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
																if v251 == int32(45) {
																	v254 = v248
																} else {
																	v254 = int32(0) - v248
																}
																*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
																v258 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
																v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
																if v259 != 0 {
																	v260 = int32(-1)
																} else {
																	v260 = int32(0)
																}
																v262 = v260
															}
														}
													}
												} else {
													v262 = v186
												}
											}
										}
									}
								} else {
									v224 = F_strlen(m, v161)
									mBase = m.M
									if base.Ui32(v224) < base.Ui32(int32(4)) {
										v233 = int32(0)
										v234 = v185
										v236 = v159
									} else {
										v228 = int32(100)
										v229 = base.I32_div_s(v185, v228)
										v233 = v185 - v229*v228
										v234 = v229
										v236 = v159
									}
									if base.Ui32(int32(15)) < base.Ui32(v234) {
										v262 = v186
									} else {
										if base.Ui32(int32(59)) < base.Ui32(v233) {
											v262 = v186
										} else {
											if base.Ui32(int32(59)) < base.Ui32(v236) {
												v262 = v186
											} else {
												v243 = int32(60)
												v248 = (v234*v243+v233)*v243 + v236
												v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
												if v251 == int32(45) {
													v254 = v248
												} else {
													v254 = int32(0) - v248
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14+int32(92)))) = v254
												v258 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
												v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
												if v259 != 0 {
													v260 = int32(-1)
												} else {
													v260 = int32(0)
												}
												v262 = v260
											}
										}
									}
								}
							}
						default:
							v262 = int32(-1)
						}
						m.G0 = v172 + int32(16)
						if v262 == int32(0) {
							v271 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
							v379 = v271
							v384 = base.I64_extend_i32_s(v159-v379)*int64(-1000000) + v28
							if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v384+int64(211813488000000000)) {
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
										F_errmsg(m, int32(400033), int32(0))
										mBase = m.M
										v465 = m.ExcPending
										if v465 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(492756), int32(716), int32(369930))
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
							} else {
								v389 = F_Int64GetDatum(m, v384)
								mBase = m.M
								v390 = m.ExcPending
								if v390 != 0 {
									return int32(0)
								} else {
									m.G0 = v14 + int32(352)
									return v389
								}
							}
						} else {
							if v262 != int32(-1) {
								if v262 == int32(-5) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v438 = m.ExcPending
									if v438 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v441 = m.ExcPending
										if v441 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(96)
											F_errmsg(m, int32(400171), v14+int32(32))
											mBase = m.M
											v449 = m.ExcPending
											if v449 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492756), int32(526), int32(369972))
												mBase = m.M
												v454 = m.ExcPending
												if v454 != 0 {
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
									v279 = m.ExcPending
									if v279 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v282 = m.ExcPending
										if v282 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(96)
											F_errmsg(m, int32(435557), v14+int32(16))
											mBase = m.M
											v290 = m.ExcPending
											if v290 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(492756), int32(530), int32(369972))
												mBase = m.M
												v295 = m.ExcPending
												if v295 != 0 {
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
								v302 = F_DecodeTimezoneName(m, v14+int32(96), v14+int32(88), v14+int32(84))
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return int32(0)
								} else {
									switch v302 {
									case 0:
										v305 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
										v379 = int32(0) - v305
									case 1:
										v308 = v14 + int32(40)
										v311 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v315 = m.G0
										v317 = v315 - int32(288)
										m.G0 = v317
										v321 = F_DetermineTimeZoneOffsetInternal(m, v308, v311, v317+int32(280))
										mBase = m.M
										v325 = F_strlcpy(m, v317+int32(16), v14+int32(96), int32(256))
										mBase = m.M
										v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+16)))
										if v326 != 0 {
											v330 = v317 + int32(16)
											v334 = v326
											for {
												v335 = F_pg_toupper(m, v334)
												mBase = m.M
												*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v335)
												v338 = v330 + int32(1)
												v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
												if v339 != 0 {
													v330 = v338
													v334 = v339
													continue
												} else {
													break
												}
												break
											}
										} else {
										}
										v354 = F_pg_interpret_timezone_abbrev(m, v317+int32(16), v317+int32(280), v317+int32(12), v317+int32(8), v311)
										mBase = m.M
										if v354 != 0 {
											v355 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
											v356 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v308)+32)) = v356
											v361 = int32(0) - v355
										} else {
											v361 = v321
										}
										m.G0 = v317 + int32(288)
										v379 = v361
									default:
										v367 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v369 = m.G0
										v370 = int32(16)
										v371 = v369 - v370
										m.G0 = v371
										v375 = F_DetermineTimeZoneOffsetInternal(m, v14+int32(40), v367, v371+int32(8))
										mBase = m.M
										m.G0 = v371 + v370
										v379 = v375
									}
									v384 = base.I64_extend_i32_s(v159-v379)*int64(-1000000) + v28
									if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v384+int64(211813488000000000)) {
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
												F_errmsg(m, int32(400033), int32(0))
												mBase = m.M
												v465 = m.ExcPending
												if v465 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(492756), int32(716), int32(369930))
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
									} else {
										v389 = F_Int64GetDatum(m, v384)
										mBase = m.M
										v390 = m.ExcPending
										if v390 != 0 {
											return int32(0)
										} else {
											m.G0 = v14 + int32(352)
											return v389
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
				v398 = m.ExcPending
				if v398 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v401 = m.ExcPending
					if v401 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(400033), int32(0))
						mBase = m.M
						v405 = m.ExcPending
						if v405 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492756), int32(707), int32(369930))
							mBase = m.M
							v410 = m.ExcPending
							if v410 != 0 {
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
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
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
							F_errmsg(m, int32(400033), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492756), int32(827), int32(35985))
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
		v10 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
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
