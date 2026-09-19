package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == v2 {
		v25 = v2
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		if v12 == int32(0) {
			v25 = v2
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if v15 != int32(7) {
				v25 = v2
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				if v18 != int32(17) {
					v25 = v2
				} else {
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
					v25 = v21 ^ int32(1)
				}
			}
		}
	}
	if v25&int32(1) != 0 {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = F_get_fn_opclass_options(m, v28)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v35 = v33
			v37 = Fn13945(m, v5, v7, v35, int32(2))
			mBase = m.M
			*(*float32)(unsafe.Add(mBase, uint32(v3))) = base.F32_convert_i32_s(v37)
			return v3
		}
	} else {
		v35 = int32(28)
		v37 = Fn13945(m, v5, v7, v35, int32(2))
		mBase = m.M
		*(*float32)(unsafe.Add(mBase, uint32(v3))) = base.F32_convert_i32_s(v37)
		return v3
	}
}
func F__ltree_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v733 int32
	_ = v733
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v764 int32
	_ = v764
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v790 int32
	_ = v790
	v2 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 == v2 {
		v45 = v2
	} else {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
		if v32 == int32(0) {
			v45 = v2
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			if v35 != int32(7) {
				v45 = v2
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				if v38 != int32(17) {
					v45 = v2
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
					v45 = v41 ^ int32(1)
				}
			}
		}
	}
	if v45&int32(1) != 0 {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v49 = F_get_fn_opclass_options(m, v48)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
			v54 = v53
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v59 = (v55 + int32(_a_F__ltree_picksplit_0)) & int32(_a_F__ltree_picksplit_1)
			v63 = v59<<(uint(int32(1))%32) + int32(4)
			v64 = F_palloc(m, v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = v64
				v67 = F_palloc(m, v63)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v67
					if base.Ui32(int32(2)) <= base.Ui32(v59) {
						v73 = v26 + int32(4)
						v77 = int32(-1)
						v78 = v2
						v83 = int32(1)
						v89 = v2
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v83<<(uint(int32(4))%32))))
							v105 = v83 + int32(1)
							v106 = v105
							v107 = v77
							v108 = v78
							v114 = v105
							v119 = v89
							for {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v114<<(uint(int32(4))%32))))
								v135 = Fn13945(m, v103, v133, v54, int32(2))
								mBase = m.M
								v136 = base.B2i32(v107 < v135)
								if v107 < v135 {
									v137 = v135
								} else {
									v137 = v107
								}
								if v107 < v135 {
									v138 = v106
								} else {
									v138 = v119
								}
								if v107 < v135 {
									v139 = v83
								} else {
									v139 = v108
								}
								v141 = v106 + int32(1)
								v143 = v141 & int32(_a_F__ltree_picksplit_1)
								if base.Ui32(v143) <= base.Ui32(v59) {
									v106 = v141
									v107 = v137
									v108 = v139
									v114 = v143
									v119 = v138
									continue
								} else {
									break
								}
								break
							}
							if v105 != v59 {
								v77 = v137
								v78 = v139
								v83 = v105
								v89 = v138
								continue
							} else {
								break
							}
							break
						}
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v149 = v139
						v159 = v146
						v160 = v138
					} else {
						v149 = v2
						v159 = v67
						v160 = v2
					}
					v171 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v171
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v171
					v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v177 = v26 + int32(4)
					v179 = int32(_a_F__ltree_picksplit_1)
					v187 = base.B2i32(v149&v179 == v171) | base.B2i32(v160&v179 == v171)
					if v187 != 0 {
						v188 = int32(1)
					} else {
						v188 = v149
					}
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v177+v188&int32(_a_F__ltree_picksplit_1)<<(uint(int32(4))%32))))
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
					v202 = int32(0)
					v204 = F_ltree_gist_alloc(m, int32(base.Ui32(v195&int32(2))>>(uint(int32(1))%32)), v194+int32(8), v54, v202, v202)
					mBase = m.M
					v205 = m.ExcPending
					if v205 != 0 {
						return int32(0)
					} else {
						if v187 != 0 {
							v207 = int32(2)
						} else {
							v207 = v160
						}
						v213 = *(*int32)(unsafe.Add(mBase, uint32(v177+v207&int32(_a_F__ltree_picksplit_1)<<(uint(int32(4))%32))))
						v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
						v221 = int32(0)
						v223 = F_ltree_gist_alloc(m, int32(base.Ui32(v214&int32(2))>>(uint(int32(1))%32)), v213+int32(8), v54, v221, v221)
						mBase = m.M
						v224 = m.ExcPending
						if v224 != 0 {
							return int32(0)
						} else {
							v225 = int32(_a_F__ltree_picksplit_1)
							v226 = v55 + v225
							v228 = v226 & v225
							v231 = F_palloc(m, v228<<(uint(int32(3))%32))
							mBase = m.M
							v232 = m.ExcPending
							if v232 != 0 {
								return int32(0)
							} else {
								if v55&int32(_a_F__ltree_picksplit_1) == int32(1) {
									F_pg_qsort(m, v231, v228, int32(8), int32(_a_F__ltree_picksplit_2))
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										v773 = v175
										v778 = v159
										v790 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v773))) = uint16(v790)
										*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v790)
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v223
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v204
										return v25
									}
								} else {
									v241 = int32(1)
									v243 = v241
									v244 = v241
									for {
										v269 = v231 + v243<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v269-int32(8)))) = uint16(v244)
										v273 = int32(4)
										v278 = *(*int32)(unsafe.Add(mBase, uint32(v177+v243<<(uint(v273)%32))))
										v280 = Fn13945(m, v204, v278, v54, int32(2))
										mBase = m.M
										v282 = Fn13945(m, v223, v278, v54, int32(2))
										mBase = m.M
										v283 = v280 - v282
										v285 = v283 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v269-v273))) = v283 ^ v285 - v285
										v290 = v244 + int32(1)
										v291 = int32(_a_F__ltree_picksplit_1)
										v292 = v290 & v291
										if base.Ui32(v292) <= base.Ui32(v226&v291) {
											v243 = v292
											v244 = v290
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v231, v228, int32(8), int32(_a_F__ltree_picksplit_2))
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return int32(0)
									} else {
										v300 = int32(1)
										if base.Ui32(v228) <= base.Ui32(v300) {
											v303 = v300
										} else {
											v303 = v228
										}
										v305 = v54 & int32(2147483644)
										v307 = v54 & int32(3)
										v308 = int32(8)
										v309 = v223 + v308
										v311 = v204 + v308
										v315 = int32(0)
										v322 = v175
										v327 = v159
										for {
											v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231+v315<<(uint(int32(3))%32)))))
											if v188&int32(_a_F__ltree_picksplit_1) == v342 {
												*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v188)
												v345 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v345 + int32(1)
												v746 = v322 + int32(2)
												v751 = v327
											} else {
												if v207&int32(_a_F__ltree_picksplit_1) == v342 {
													*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v207)
													v733 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v733 + int32(1)
													v746 = v322
													v751 = v327 + int32(2)
												} else {
													v358 = *(*int32)(unsafe.Add(mBase, uint32(v177+v342<<(uint(int32(4))%32))))
													v360 = Fn13945(m, v204, v358, v54, int32(2))
													mBase = m.M
													v363 = Fn13945(m, v223, v358, v54, int32(2))
													mBase = m.M
													v365 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													v366 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v367 = v365 - v366
													if base.F64_lt(base.F64_convert_i32_s(v360), base.F64_add(base.F64_convert_i32_s(v363), base.F64_mul(base.F64_convert_i32_s(v367*v367*v367), float64(-1e-05)))) != 0 {
														v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+4)))
														if v375&int32(2) != 0 {
														} else {
															v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+4)))
															if v378&int32(2) != 0 {
																if v54 == int32(0) {
																} else {
																	base.MemoryFill(m, v311, int32(255), v54)
																}
															} else {
																if v54 <= int32(0) {
																} else {
																	v388 = v358 + int32(8)
																	v389 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v395 = v389
																		v396 = v389
																		for {
																			v418 = v395 + v311
																			v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
																			v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v388))))
																			v422 = v419 | v421
																			*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v422)
																			v425 = v395 | int32(1)
																			v426 = v311 + v425
																			v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
																			v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v388))))
																			v430 = v427 | v429
																			*(*uint8)(unsafe.Add(mBase, uint32(v426))) = uint8(v430)
																			v433 = v395 | int32(2)
																			v434 = v311 + v433
																			v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
																			v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v388))))
																			v438 = v435 | v437
																			*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v438)
																			v441 = v395 | int32(3)
																			v442 = v311 + v441
																			v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
																			v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441+v388))))
																			v446 = v443 | v445
																			*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v446)
																			v448 = int32(4)
																			v449 = v395 + v448
																			v451 = v396 + v448
																			if v451 != v305 {
																				v395 = v449
																				v396 = v451
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v307 == int32(0) {
																		} else {
																			v457 = v449
																			v481 = v457
																			v487 = v389
																			for {
																				v503 = v481 + v311
																				v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
																				v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v388))))
																				v507 = v504 | v506
																				*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v507)
																				v509 = int32(1)
																				v512 = v487 + v509
																				if v512 != v307 {
																					v481 = v481 + v509
																					v487 = v512
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v457 = v389
																		v481 = v457
																		v487 = v389
																		for {
																			v503 = v481 + v311
																			v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
																			v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v388))))
																			v507 = v504 | v506
																			*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v507)
																			v509 = int32(1)
																			v512 = v487 + v509
																			if v512 != v307 {
																				v481 = v481 + v509
																				v487 = v512
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v342)
														v539 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v539 + int32(1)
														v746 = v322 + int32(2)
														v751 = v327
													} else {
														v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+4)))
														if v545&int32(2) != 0 {
														} else {
															v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+4)))
															if v548&int32(2) != 0 {
																if v54 == int32(0) {
																} else {
																	base.MemoryFill(m, v309, int32(255), v54)
																}
															} else {
																if v54 <= int32(0) {
																} else {
																	v558 = v358 + int32(8)
																	v559 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v565 = v559
																		v566 = v559
																		for {
																			v588 = v565 + v309
																			v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
																			v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v558))))
																			v592 = v589 | v591
																			*(*uint8)(unsafe.Add(mBase, uint32(v588))) = uint8(v592)
																			v595 = v565 | int32(1)
																			v596 = v309 + v595
																			v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
																			v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595+v558))))
																			v600 = v597 | v599
																			*(*uint8)(unsafe.Add(mBase, uint32(v596))) = uint8(v600)
																			v603 = v565 | int32(2)
																			v604 = v309 + v603
																			v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
																			v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603+v558))))
																			v608 = v605 | v607
																			*(*uint8)(unsafe.Add(mBase, uint32(v604))) = uint8(v608)
																			v611 = v565 | int32(3)
																			v612 = v309 + v611
																			v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
																			v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611+v558))))
																			v616 = v613 | v615
																			*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v616)
																			v618 = int32(4)
																			v619 = v565 + v618
																			v621 = v566 + v618
																			if v621 != v305 {
																				v565 = v619
																				v566 = v621
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v307 == int32(0) {
																		} else {
																			v627 = v619
																			v651 = v627
																			v657 = v559
																			for {
																				v673 = v651 + v309
																				v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
																				v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v558))))
																				v677 = v674 | v676
																				*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v677)
																				v679 = int32(1)
																				v682 = v657 + v679
																				if v682 != v307 {
																					v651 = v651 + v679
																					v657 = v682
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v627 = v559
																		v651 = v627
																		v657 = v559
																		for {
																			v673 = v651 + v309
																			v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
																			v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v558))))
																			v677 = v674 | v676
																			*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v677)
																			v679 = int32(1)
																			v682 = v657 + v679
																			if v682 != v307 {
																				v651 = v651 + v679
																				v657 = v682
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v342)
														v733 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v733 + int32(1)
														v746 = v322
														v751 = v327 + int32(2)
													}
												}
											}
											v764 = v315 + int32(1)
											if v764 != v303 {
												v315 = v764
												v322 = v746
												v327 = v751
												continue
											} else {
												break
											}
											break
										}
										v773 = v746
										v778 = v751
										v790 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v773))) = uint16(v790)
										*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v790)
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v223
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v204
										return v25
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v54 = int32(28)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v59 = (v55 + int32(_a_F__ltree_picksplit_0)) & int32(_a_F__ltree_picksplit_1)
		v63 = v59<<(uint(int32(1))%32) + int32(4)
		v64 = F_palloc(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v64
			v67 = F_palloc(m, v63)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v67
				if base.Ui32(int32(2)) <= base.Ui32(v59) {
					v73 = v26 + int32(4)
					v77 = int32(-1)
					v78 = v2
					v83 = int32(1)
					v89 = v2
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v83<<(uint(int32(4))%32))))
						v105 = v83 + int32(1)
						v106 = v105
						v107 = v77
						v108 = v78
						v114 = v105
						v119 = v89
						for {
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v114<<(uint(int32(4))%32))))
							v135 = Fn13945(m, v103, v133, v54, int32(2))
							mBase = m.M
							v136 = base.B2i32(v107 < v135)
							if v107 < v135 {
								v137 = v135
							} else {
								v137 = v107
							}
							if v107 < v135 {
								v138 = v106
							} else {
								v138 = v119
							}
							if v107 < v135 {
								v139 = v83
							} else {
								v139 = v108
							}
							v141 = v106 + int32(1)
							v143 = v141 & int32(_a_F__ltree_picksplit_1)
							if base.Ui32(v143) <= base.Ui32(v59) {
								v106 = v141
								v107 = v137
								v108 = v139
								v114 = v143
								v119 = v138
								continue
							} else {
								break
							}
							break
						}
						if v105 != v59 {
							v77 = v137
							v78 = v139
							v83 = v105
							v89 = v138
							continue
						} else {
							break
						}
						break
					}
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v149 = v139
					v159 = v146
					v160 = v138
				} else {
					v149 = v2
					v159 = v67
					v160 = v2
				}
				v171 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v171
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v171
				v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v177 = v26 + int32(4)
				v179 = int32(_a_F__ltree_picksplit_1)
				v187 = base.B2i32(v149&v179 == v171) | base.B2i32(v160&v179 == v171)
				if v187 != 0 {
					v188 = int32(1)
				} else {
					v188 = v149
				}
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v177+v188&int32(_a_F__ltree_picksplit_1)<<(uint(int32(4))%32))))
				v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
				v202 = int32(0)
				v204 = F_ltree_gist_alloc(m, int32(base.Ui32(v195&int32(2))>>(uint(int32(1))%32)), v194+int32(8), v54, v202, v202)
				mBase = m.M
				v205 = m.ExcPending
				if v205 != 0 {
					return int32(0)
				} else {
					if v187 != 0 {
						v207 = int32(2)
					} else {
						v207 = v160
					}
					v213 = *(*int32)(unsafe.Add(mBase, uint32(v177+v207&int32(_a_F__ltree_picksplit_1)<<(uint(int32(4))%32))))
					v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
					v221 = int32(0)
					v223 = F_ltree_gist_alloc(m, int32(base.Ui32(v214&int32(2))>>(uint(int32(1))%32)), v213+int32(8), v54, v221, v221)
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						v225 = int32(_a_F__ltree_picksplit_1)
						v226 = v55 + v225
						v228 = v226 & v225
						v231 = F_palloc(m, v228<<(uint(int32(3))%32))
						mBase = m.M
						v232 = m.ExcPending
						if v232 != 0 {
							return int32(0)
						} else {
							if v55&int32(_a_F__ltree_picksplit_1) == int32(1) {
								F_pg_qsort(m, v231, v228, int32(8), int32(_a_F__ltree_picksplit_2))
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									v773 = v175
									v778 = v159
									v790 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v773))) = uint16(v790)
									*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v790)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v223
									*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v204
									return v25
								}
							} else {
								v241 = int32(1)
								v243 = v241
								v244 = v241
								for {
									v269 = v231 + v243<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v269-int32(8)))) = uint16(v244)
									v273 = int32(4)
									v278 = *(*int32)(unsafe.Add(mBase, uint32(v177+v243<<(uint(v273)%32))))
									v280 = Fn13945(m, v204, v278, v54, int32(2))
									mBase = m.M
									v282 = Fn13945(m, v223, v278, v54, int32(2))
									mBase = m.M
									v283 = v280 - v282
									v285 = v283 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v269-v273))) = v283 ^ v285 - v285
									v290 = v244 + int32(1)
									v291 = int32(_a_F__ltree_picksplit_1)
									v292 = v290 & v291
									if base.Ui32(v292) <= base.Ui32(v226&v291) {
										v243 = v292
										v244 = v290
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v231, v228, int32(8), int32(_a_F__ltree_picksplit_2))
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return int32(0)
								} else {
									v300 = int32(1)
									if base.Ui32(v228) <= base.Ui32(v300) {
										v303 = v300
									} else {
										v303 = v228
									}
									v305 = v54 & int32(2147483644)
									v307 = v54 & int32(3)
									v308 = int32(8)
									v309 = v223 + v308
									v311 = v204 + v308
									v315 = int32(0)
									v322 = v175
									v327 = v159
									for {
										v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231+v315<<(uint(int32(3))%32)))))
										if v188&int32(_a_F__ltree_picksplit_1) == v342 {
											*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v188)
											v345 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v345 + int32(1)
											v746 = v322 + int32(2)
											v751 = v327
										} else {
											if v207&int32(_a_F__ltree_picksplit_1) == v342 {
												*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v207)
												v733 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v733 + int32(1)
												v746 = v322
												v751 = v327 + int32(2)
											} else {
												v358 = *(*int32)(unsafe.Add(mBase, uint32(v177+v342<<(uint(int32(4))%32))))
												v360 = Fn13945(m, v204, v358, v54, int32(2))
												mBase = m.M
												v363 = Fn13945(m, v223, v358, v54, int32(2))
												mBase = m.M
												v365 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v366 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												v367 = v365 - v366
												if base.F64_lt(base.F64_convert_i32_s(v360), base.F64_add(base.F64_convert_i32_s(v363), base.F64_mul(base.F64_convert_i32_s(v367*v367*v367), float64(-1e-05)))) != 0 {
													v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+4)))
													if v375&int32(2) != 0 {
													} else {
														v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+4)))
														if v378&int32(2) != 0 {
															if v54 == int32(0) {
															} else {
																base.MemoryFill(m, v311, int32(255), v54)
															}
														} else {
															if v54 <= int32(0) {
															} else {
																v388 = v358 + int32(8)
																v389 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v395 = v389
																	v396 = v389
																	for {
																		v418 = v395 + v311
																		v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
																		v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v388))))
																		v422 = v419 | v421
																		*(*uint8)(unsafe.Add(mBase, uint32(v418))) = uint8(v422)
																		v425 = v395 | int32(1)
																		v426 = v311 + v425
																		v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
																		v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v388))))
																		v430 = v427 | v429
																		*(*uint8)(unsafe.Add(mBase, uint32(v426))) = uint8(v430)
																		v433 = v395 | int32(2)
																		v434 = v311 + v433
																		v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
																		v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v388))))
																		v438 = v435 | v437
																		*(*uint8)(unsafe.Add(mBase, uint32(v434))) = uint8(v438)
																		v441 = v395 | int32(3)
																		v442 = v311 + v441
																		v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
																		v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441+v388))))
																		v446 = v443 | v445
																		*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v446)
																		v448 = int32(4)
																		v449 = v395 + v448
																		v451 = v396 + v448
																		if v451 != v305 {
																			v395 = v449
																			v396 = v451
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v307 == int32(0) {
																	} else {
																		v457 = v449
																		v481 = v457
																		v487 = v389
																		for {
																			v503 = v481 + v311
																			v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
																			v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v388))))
																			v507 = v504 | v506
																			*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v507)
																			v509 = int32(1)
																			v512 = v487 + v509
																			if v512 != v307 {
																				v481 = v481 + v509
																				v487 = v512
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v457 = v389
																	v481 = v457
																	v487 = v389
																	for {
																		v503 = v481 + v311
																		v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
																		v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v388))))
																		v507 = v504 | v506
																		*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v507)
																		v509 = int32(1)
																		v512 = v487 + v509
																		if v512 != v307 {
																			v481 = v481 + v509
																			v487 = v512
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v322))) = uint16(v342)
													v539 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v539 + int32(1)
													v746 = v322 + int32(2)
													v751 = v327
												} else {
													v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+4)))
													if v545&int32(2) != 0 {
													} else {
														v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+4)))
														if v548&int32(2) != 0 {
															if v54 == int32(0) {
															} else {
																base.MemoryFill(m, v309, int32(255), v54)
															}
														} else {
															if v54 <= int32(0) {
															} else {
																v558 = v358 + int32(8)
																v559 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v565 = v559
																	v566 = v559
																	for {
																		v588 = v565 + v309
																		v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588))))
																		v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565+v558))))
																		v592 = v589 | v591
																		*(*uint8)(unsafe.Add(mBase, uint32(v588))) = uint8(v592)
																		v595 = v565 | int32(1)
																		v596 = v309 + v595
																		v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
																		v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595+v558))))
																		v600 = v597 | v599
																		*(*uint8)(unsafe.Add(mBase, uint32(v596))) = uint8(v600)
																		v603 = v565 | int32(2)
																		v604 = v309 + v603
																		v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
																		v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603+v558))))
																		v608 = v605 | v607
																		*(*uint8)(unsafe.Add(mBase, uint32(v604))) = uint8(v608)
																		v611 = v565 | int32(3)
																		v612 = v309 + v611
																		v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612))))
																		v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611+v558))))
																		v616 = v613 | v615
																		*(*uint8)(unsafe.Add(mBase, uint32(v612))) = uint8(v616)
																		v618 = int32(4)
																		v619 = v565 + v618
																		v621 = v566 + v618
																		if v621 != v305 {
																			v565 = v619
																			v566 = v621
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v307 == int32(0) {
																	} else {
																		v627 = v619
																		v651 = v627
																		v657 = v559
																		for {
																			v673 = v651 + v309
																			v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
																			v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v558))))
																			v677 = v674 | v676
																			*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v677)
																			v679 = int32(1)
																			v682 = v657 + v679
																			if v682 != v307 {
																				v651 = v651 + v679
																				v657 = v682
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v627 = v559
																	v651 = v627
																	v657 = v559
																	for {
																		v673 = v651 + v309
																		v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
																		v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v558))))
																		v677 = v674 | v676
																		*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v677)
																		v679 = int32(1)
																		v682 = v657 + v679
																		if v682 != v307 {
																			v651 = v651 + v679
																			v657 = v682
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v342)
													v733 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v733 + int32(1)
													v746 = v322
													v751 = v327 + int32(2)
												}
											}
										}
										v764 = v315 + int32(1)
										if v764 != v303 {
											v315 = v764
											v322 = v746
											v327 = v751
											continue
										} else {
											break
										}
										break
									}
									v773 = v746
									v778 = v751
									v790 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v773))) = uint16(v790)
									*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v790)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v223
									*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v204
									return v25
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__ltree_r_isparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltree_r_isparent_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F__ltree_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v254 int32
	_ = v254
	v2 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == v2 {
		v37 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v37&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v24 == int32(0) {
		v37 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v27 != int32(7) {
		v37 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v30 != int32(17) {
		v37 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
	v37 = v33 ^ int32(1)
	goto L2
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = F_get_fn_opclass_options(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v46 = int32(28)
	goto L9
L9:
	;
	v47 = int32(0)
	v51 = F_ltree_gist_alloc(m, v47, v47, v46, v47, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v46 = v45
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v53 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(base.Ui32(v254) >> (uint(int32(2)) % 32))
	return v51
L14:
	;
	v57 = v51 + int32(8)
	v61 = v46 & int32(3)
	v66 = v53
	v76 = v2
	goto L15
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(4)+v76<<(uint(int32(4))%32))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v86&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v234 | int32(2)
	goto L13
L17:
	;
	if base.B2i32(v46 <= int32(0)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v94 = v85 + int32(8)
	v95 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v46) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v213 = v66
	goto L22
L22:
	;
	v230 = v76 + int32(1)
	if v230 < v213 {
		v66 = v213
		v76 = v230
		goto L15
	} else {
		goto L34
	}
L23:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v213 = v212
	goto L22
L24:
	;
	v100 = v95
	v107 = v95
	goto L27
L25:
	;
	v153 = v95
	goto L26
L26:
	;
	v169 = v153
	v180 = v95
	goto L31
L27:
	;
	v116 = v100 + v57
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+v94))))
	v120 = v117 | v119
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v120)
	v123 = v100 | int32(1)
	v124 = v57 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v123))))
	v128 = v125 | v127
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v128)
	v131 = v100 | int32(2)
	v132 = v57 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v131))))
	v136 = v133 | v135
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v136)
	v139 = v100 | int32(3)
	v140 = v57 + v139
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v139))))
	v144 = v141 | v143
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v144)
	v146 = int32(4)
	v147 = v100 + v146
	v149 = v107 + v146
	if v149 != v46&int32(2147483644) {
		v100 = v147
		v107 = v149
		goto L27
	} else {
		goto L29
	}
L28:
	;
	if v61 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v153 = v147
	goto L26
L31:
	;
	v185 = v169 + v57
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v94))))
	v189 = v186 | v188
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v189)
	v191 = int32(1)
	v194 = v180 + v191
	if v194 != v61 {
		v169 = v169 + v191
		v180 = v194
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L23
L33:
	;
	goto L32
L34:
	;
	goto L13
}
func F_ltree_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		v152 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v14 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v179 = (v152 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	v28 = int32(8)
	v33 = v22
	v37 = v14 + v28
	v38 = v19 + v28
	v43 = v21
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v37 + v44
	v47 = v38 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v152 = v132
	goto L5
L9:
	;
	v132 = v33 - int32(1)
	if v33 < int32(2) {
		v152 = v132
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v51 = v48
	goto L12
L11:
	;
	v51 = v49
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v113 = int32(0)
	goto L13
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L25
L16:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L18
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L22:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L15
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = v92 - v93
	goto L13
L27:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v48 == v49 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v113 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(10)
	v179 = (v33*v117 + v117) * (v48 - v49)
	goto L4
L35:
	;
	v129 = int32(-10)
	goto L37
L36:
	;
	v129 = int32(10)
	goto L37
L37:
	;
	v179 = (v33 + int32(1)) * v129
	goto L4
L38:
	;
	v135 = int32(9)
	v137 = int32(_a_F_ltree_cmp_0)
	v145 = int32(1)
	if v145 < v43 {
		v33 = v132
		v37 = v37 + (v48+v135)&v137
		v38 = v38 + (v49+v135)&v137
		v43 = v43 - v145
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v184 != v19 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v19)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	return v179
L47:
	;
	goto L46
}
func F_ltree_concat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = v12 + v13
	if base.Ui32(v14) < base.Ui32(int32(_a_F_ltree_concat_0)) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = int32(2)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v26 = F_palloc0(m, int32(base.Ui32(v17)>>(uint(v18)%32))+int32(base.Ui32(v20)>>(uint(v18)%32))-int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)) = uint16(v14)
			v33 = int32(-4)
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = (v30+v31&v33)&v33 - int32(32)
			v41 = int32(8)
			v42 = v26 + v41
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v47 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - v41
			if v47 != 0 {
				base.MemoryCopy(m, v42, l0+int32(8), v47)
			} else {
			}
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v55 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(8)
			if v55 != 0 {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v60 = int32(8)
				base.MemoryCopy(m, v42+int32(base.Ui32(v56)>>(uint(int32(2))%32))-v60, l1+v60, v55)
			} else {
			}
			m.G0 = v10 + int32(16)
			return v26
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(_a_F_ltree_concat_1)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
				F_errmsg(m, int32(_a_F_ltree_concat_2), v10)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ltree_concat_3), int32(353), int32(_a_F_ltree_concat_4))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
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
func F_ltree_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1018 int32
	_ = v1018
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1237 int32
	_ = v1237
	var v1244 int32
	_ = v1244
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1331 int32
	_ = v1331
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1454 int32
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1556 int32
	_ = v1556
	var v1563 int32
	_ = v1563
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1678 int32
	_ = v1678
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1791 int32
	_ = v1791
	var v1816 int32
	_ = v1816
	var v1822 int32
	_ = v1822
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1907 int32
	_ = v1907
	var v1913 int32
	_ = v1913
	var v1940 int32
	_ = v1940
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2027 int32
	_ = v2027
	var v2033 int32
	_ = v2033
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2058 int32
	_ = v2058
	var v2066 int32
	_ = v2066
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2082 int32
	_ = v2082
	var v2089 int32
	_ = v2089
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2108 int32
	_ = v2108
	var v2116 int32
	_ = v2116
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2229 int32
	_ = v2229
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2345 int32
	_ = v2345
	var v2351 int32
	_ = v2351
	var v2378 int32
	_ = v2378
	var v2393 int32
	_ = v2393
	var v2413 int32
	_ = v2413
	var v2421 int32
	_ = v2421
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2465 int32
	_ = v2465
	var v2472 int32
	_ = v2472
	var v2475 int32
	_ = v2475
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2532 int32
	_ = v2532
	var v2539 int32
	_ = v2539
	var v2560 int32
	_ = v2560
	var v2565 int32
	_ = v2565
	var v2568 int32
	_ = v2568
	var v2579 int32
	_ = v2579
	var v2582 int32
	_ = v2582
	var v2590 int32
	_ = v2590
	var v2594 int32
	_ = v2594
	var v2599 int32
	_ = v2599
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == v2 {
		v42 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v42&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	if v29 == int32(0) {
		v42 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 != int32(7) {
		v42 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v35 != int32(17) {
		v42 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	v42 = v38 ^ int32(1)
	goto L2
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = F_get_fn_opclass_options(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v51 = int32(8)
	goto L9
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v53)
	switch v22 - int32(1) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	default:
		goto L12
	case 9:
		goto L19
	case 10:
		goto L18
	case 11, 12:
		goto L17
	case 13, 14:
		goto L16
	case 15, 16:
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v51 = v50
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L10
	} else {
		goto L566
	}
L13:
	;
	v2579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2579 != v2565 {
		goto L562
	} else {
		goto L563
	}
L14:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v2458&int32(3) != 0 {
		goto L541
	} else {
		goto L542
	}
L15:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1990 = F_pg_detoast_datum(m, v1989)
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L10
	} else {
		goto L446
	}
L16:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1957 = F_pg_detoast_datum(m, v1956)
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L10
	} else {
		goto L439
	}
L17:
	;
	v1593 = v52 + int32(8)
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1595 = F_pg_detoast_datum(m, v1594)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L10
	} else {
		goto L369
	}
L18:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1282 = F_pg_detoast_datum(m, v1281)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L10
	} else {
		goto L285
	}
L19:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v969 = F_pg_detoast_datum_copy(m, v968)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L10
	} else {
		goto L212
	}
L20:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v727 = F_pg_detoast_datum(m, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L10
	} else {
		goto L160
	}
L21:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v606 = F_pg_detoast_datum(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L10
	} else {
		goto L135
	}
L22:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v278 = F_pg_detoast_datum(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L10
	} else {
		goto L67
	}
L23:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v169 = F_pg_detoast_datum(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L45
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v58 = F_pg_detoast_datum(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+16)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v61)+12)))
	if v63&int32(1) == int32(0) {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v70 = int32(0)
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	if base.B2i32(v77 == v70)|base.B2i32(v80 == v70) != 0 {
		v144 = v77
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v2565 = v58
	v2568 = base.B2i32(int32(0) < v165)
	goto L13
L28:
	;
	v165 = (v144 + int32(1)) * (v77 - v80) * int32(10)
	goto L27
L29:
	;
	v88 = v58 + int32(8)
	v89 = v52 + int32(16)
	v92 = v77
	v95 = v80
	goto L30
L30:
	;
	v97 = int32(2)
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
	if base.Ui32(v101) < base.Ui32(v102) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v144 = v124
	goto L28
L32:
	;
	v124 = v92 - int32(1)
	if v92 < int32(2) {
		v144 = v124
		goto L28
	} else {
		goto L43
	}
L33:
	;
	v104 = v101
	goto L35
L34:
	;
	v104 = v102
	goto L35
L35:
	;
	v105 = F_memcmp(m, v88+v97, v89+v97, v104)
	mBase = m.M
	if v105 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v101 == v102 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v105 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v109 = int32(10)
	v165 = (v92*v109 + v109) * (v101 - v102)
	goto L27
L40:
	;
	v121 = int32(-10)
	goto L42
L41:
	;
	v121 = int32(10)
	goto L42
L42:
	;
	v165 = (v92 + int32(1)) * v121
	goto L27
L43:
	;
	v127 = int32(9)
	v129 = int32(_a_F_ltree_consistent_0)
	v137 = int32(1)
	if v137 < v95 {
		v88 = v88 + (v101+v127)&v129
		v89 = v89 + (v102+v127)&v129
		v92 = v124
		v95 = v95 - v137
		goto L30
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v172&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v175 = int32(0)
	goto L48
L47:
	;
	v175 = v51
	goto L48
L48:
	;
	v176 = v52 + v175
	v179 = int32(0)
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+int32(8))+4)))
	if base.B2i32(v186 == v179)|base.B2i32(v189 == v179) != 0 {
		v253 = v186
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v2565 = v169
	v2568 = base.B2i32(int32(0) <= v274)
	goto L13
L50:
	;
	v274 = (v253 + int32(1)) * (v186 - v189) * int32(10)
	goto L49
L51:
	;
	v197 = v169 + int32(8)
	v198 = v176 + int32(16)
	v201 = v186
	v204 = v189
	goto L52
L52:
	;
	v206 = int32(2)
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197))))
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	if base.Ui32(v210) < base.Ui32(v211) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v253 = v233
	goto L50
L54:
	;
	v233 = v201 - int32(1)
	if v201 < int32(2) {
		v253 = v233
		goto L50
	} else {
		goto L65
	}
L55:
	;
	v213 = v210
	goto L57
L56:
	;
	v213 = v211
	goto L57
L57:
	;
	v214 = F_memcmp(m, v197+v206, v198+v206, v213)
	mBase = m.M
	if v214 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v210 == v211 {
		goto L54
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v214 < int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v218 = int32(10)
	v274 = (v201*v218 + v218) * (v210 - v211)
	goto L49
L62:
	;
	v230 = int32(-10)
	goto L64
L63:
	;
	v230 = int32(10)
	goto L64
L64:
	;
	v274 = (v201 + int32(1)) * v230
	goto L49
L65:
	;
	v236 = int32(9)
	v238 = int32(_a_F_ltree_consistent_0)
	v246 = int32(1)
	if v246 < v204 {
		v197 = v197 + (v210+v236)&v238
		v198 = v198 + (v211+v236)&v238
		v201 = v233
		v204 = v204 - v246
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+16)))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v281)+12)))
	if v283&int32(1) != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v288 = int32(0)
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+4)))
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	if base.B2i32(v295 == v288)|base.B2i32(v298 == v288) != 0 {
		v362 = v295
		goto L72
	} else {
		goto L73
	}
L69:
	;
	goto L70
L70:
	;
	v387 = v52 + int32(8)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v389&int32(3) != 0 {
		goto L89
	} else {
		goto L90
	}
L71:
	;
	v2565 = v278
	v2568 = base.B2i32(v383 == int32(0))
	goto L13
L72:
	;
	v383 = (v362 + int32(1)) * (v295 - v298) * int32(10)
	goto L71
L73:
	;
	v306 = v278 + int32(8)
	v307 = v52 + int32(16)
	v310 = v295
	v313 = v298
	goto L74
L74:
	;
	v315 = int32(2)
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306))))
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307))))
	if base.Ui32(v319) < base.Ui32(v320) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v362 = v342
	goto L72
L76:
	;
	v342 = v310 - int32(1)
	if v310 < int32(2) {
		v362 = v342
		goto L72
	} else {
		goto L87
	}
L77:
	;
	v322 = v319
	goto L79
L78:
	;
	v322 = v320
	goto L79
L79:
	;
	v323 = F_memcmp(m, v306+v315, v307+v315, v322)
	mBase = m.M
	if v323 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v319 == v320 {
		goto L76
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v323 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v327 = int32(10)
	v383 = (v310*v327 + v327) * (v319 - v320)
	goto L71
L84:
	;
	v339 = int32(-10)
	goto L86
L85:
	;
	v339 = int32(10)
	goto L86
L86:
	;
	v383 = (v310 + int32(1)) * v339
	goto L71
L87:
	;
	v345 = int32(9)
	v347 = int32(_a_F_ltree_consistent_0)
	v355 = int32(1)
	if v355 < v313 {
		v306 = v306 + (v319+v345)&v347
		v307 = v307 + (v320+v345)&v347
		v310 = v342
		v313 = v313 - v355
		goto L74
	} else {
		goto L88
	}
L88:
	;
	goto L75
L89:
	;
	v392 = int32(0)
	goto L91
L90:
	;
	v392 = v51
	goto L91
L91:
	;
	v393 = v387 + v392
	v394 = int32(0)
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+4)))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+4)))
	if base.B2i32(v401 == v394)|base.B2i32(v404 == v394) != 0 {
		v468 = v401
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v489 < int32(0) {
		v2565 = v278
		v2568 = v2
		goto L13
	} else {
		goto L110
	}
L93:
	;
	v489 = (v468 + int32(1)) * (v401 - v404) * int32(10)
	goto L92
L94:
	;
	v408 = int32(8)
	v412 = v278 + v408
	v413 = v393 + v408
	v416 = v401
	v419 = v404
	goto L95
L95:
	;
	v421 = int32(2)
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412))))
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413))))
	if base.Ui32(v425) < base.Ui32(v426) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v468 = v448
	goto L93
L97:
	;
	v448 = v416 - int32(1)
	if v416 < int32(2) {
		v468 = v448
		goto L93
	} else {
		goto L108
	}
L98:
	;
	v428 = v425
	goto L100
L99:
	;
	v428 = v426
	goto L100
L100:
	;
	v429 = F_memcmp(m, v412+v421, v413+v421, v428)
	mBase = m.M
	if v429 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v425 == v426 {
		goto L97
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v429 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v433 = int32(10)
	v489 = (v416*v433 + v433) * (v425 - v426)
	goto L92
L105:
	;
	v445 = int32(-10)
	goto L107
L106:
	;
	v445 = int32(10)
	goto L107
L107:
	;
	v489 = (v416 + int32(1)) * v445
	goto L92
L108:
	;
	v451 = int32(9)
	v453 = int32(_a_F_ltree_consistent_0)
	v461 = int32(1)
	if v461 < v419 {
		v412 = v412 + (v425+v451)&v453
		v413 = v413 + (v426+v451)&v453
		v416 = v448
		v419 = v419 - v461
		goto L95
	} else {
		goto L109
	}
L109:
	;
	goto L96
L110:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v492&int32(1) != 0 {
		v506 = v387
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v507 = int32(0)
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+4)))
	v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506)+4)))
	if base.B2i32(v514 == v507)|base.B2i32(v517 == v507) != 0 {
		v581 = v514
		goto L118
	} else {
		goto L119
	}
L112:
	;
	if v492&int32(2) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v498 = int32(0)
	goto L115
L114:
	;
	v498 = v51
	goto L115
L115:
	;
	v499 = v387 + v498
	if v492&int32(4) != 0 {
		v506 = v499
		goto L111
	} else {
		goto L116
	}
L116:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v506 = v499 + int32(base.Ui32(v502)>>(uint(int32(2))%32))
	goto L111
L117:
	;
	v2565 = v278
	v2568 = base.B2i32(v602 <= int32(0))
	goto L13
L118:
	;
	v602 = (v581 + int32(1)) * (v514 - v517) * int32(10)
	goto L117
L119:
	;
	v521 = int32(8)
	v525 = v278 + v521
	v526 = v506 + v521
	v529 = v514
	v532 = v517
	goto L120
L120:
	;
	v534 = int32(2)
	v538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525))))
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v526))))
	if base.Ui32(v538) < base.Ui32(v539) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v581 = v561
	goto L118
L122:
	;
	v561 = v529 - int32(1)
	if v529 < int32(2) {
		v581 = v561
		goto L118
	} else {
		goto L133
	}
L123:
	;
	v541 = v538
	goto L125
L124:
	;
	v541 = v539
	goto L125
L125:
	;
	v542 = F_memcmp(m, v525+v534, v526+v534, v541)
	mBase = m.M
	if v542 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if v538 == v539 {
		goto L122
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v542 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v546 = int32(10)
	v602 = (v529*v546 + v546) * (v538 - v539)
	goto L117
L130:
	;
	v558 = int32(-10)
	goto L132
L131:
	;
	v558 = int32(10)
	goto L132
L132:
	;
	v602 = (v529 + int32(1)) * v558
	goto L117
L133:
	;
	v564 = int32(9)
	v566 = int32(_a_F_ltree_consistent_0)
	v574 = int32(1)
	if v574 < v532 {
		v525 = v525 + (v538+v564)&v566
		v526 = v526 + (v539+v564)&v566
		v529 = v561
		v532 = v532 - v574
		goto L120
	} else {
		goto L134
	}
L134:
	;
	goto L121
L135:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v610&int32(1) != 0 {
		v627 = v52 + int32(8)
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v628 = int32(0)
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v606)+4)))
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v627)+4)))
	if base.B2i32(v635 == v628)|base.B2i32(v638 == v628) != 0 {
		v702 = v635
		goto L143
	} else {
		goto L144
	}
L137:
	;
	if v610&int32(2) != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v616 = int32(0)
	goto L140
L139:
	;
	v616 = v51
	goto L140
L140:
	;
	v619 = v52 + v616 + int32(8)
	if v610&int32(4) != 0 {
		v627 = v619
		goto L136
	} else {
		goto L141
	}
L141:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v627 = v619 + int32(base.Ui32(v622)>>(uint(int32(2))%32))
	goto L136
L142:
	;
	v2565 = v606
	v2568 = base.B2i32(v723 <= int32(0))
	goto L13
L143:
	;
	v723 = (v702 + int32(1)) * (v635 - v638) * int32(10)
	goto L142
L144:
	;
	v642 = int32(8)
	v646 = v606 + v642
	v647 = v627 + v642
	v650 = v635
	v653 = v638
	goto L145
L145:
	;
	v655 = int32(2)
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646))))
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v647))))
	if base.Ui32(v659) < base.Ui32(v660) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v702 = v682
	goto L143
L147:
	;
	v682 = v650 - int32(1)
	if v650 < int32(2) {
		v702 = v682
		goto L143
	} else {
		goto L158
	}
L148:
	;
	v662 = v659
	goto L150
L149:
	;
	v662 = v660
	goto L150
L150:
	;
	v663 = F_memcmp(m, v646+v655, v647+v655, v662)
	mBase = m.M
	if v663 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v659 == v660 {
		goto L147
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v663 < int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v667 = int32(10)
	v723 = (v650*v667 + v667) * (v659 - v660)
	goto L142
L155:
	;
	v679 = int32(-10)
	goto L157
L156:
	;
	v679 = int32(10)
	goto L157
L157:
	;
	v723 = (v650 + int32(1)) * v679
	goto L142
L158:
	;
	v685 = int32(9)
	v687 = int32(_a_F_ltree_consistent_0)
	v695 = int32(1)
	if v695 < v653 {
		v646 = v646 + (v659+v685)&v687
		v647 = v647 + (v660+v685)&v687
		v650 = v682
		v653 = v653 - v695
		goto L145
	} else {
		goto L159
	}
L159:
	;
	goto L146
L160:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v730 = int32(1)
	v731 = v729 & v730
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v732)+16)))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732+v733)+12)))
	if v735&v730 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v731 != 0 {
		v754 = v52 + int32(8)
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	if v731 != 0 {
		v869 = v52 + int32(8)
		goto L188
	} else {
		goto L189
	}
L164:
	;
	v755 = int32(0)
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+4)))
	v765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v754)+4)))
	if base.B2i32(v762 == v755)|base.B2i32(v765 == v755) != 0 {
		v829 = v762
		goto L171
	} else {
		goto L172
	}
L165:
	;
	if v729&int32(2) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v743 = int32(0)
	goto L168
L167:
	;
	v743 = v51
	goto L168
L168:
	;
	v746 = v52 + v743 + int32(8)
	if v729&int32(4) != 0 {
		v754 = v746
		goto L164
	} else {
		goto L169
	}
L169:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v754 = v746 + int32(base.Ui32(v749)>>(uint(int32(2))%32))
	goto L164
L170:
	;
	v2565 = v727
	v2568 = int32(base.Ui32(v850) >> (uint(int32(31)) % 32))
	goto L13
L171:
	;
	v850 = (v829 + int32(1)) * (v762 - v765) * int32(10)
	goto L170
L172:
	;
	v769 = int32(8)
	v773 = v727 + v769
	v774 = v754 + v769
	v777 = v762
	v780 = v765
	goto L173
L173:
	;
	v782 = int32(2)
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v773))))
	v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v774))))
	if base.Ui32(v786) < base.Ui32(v787) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v829 = v809
	goto L171
L175:
	;
	v809 = v777 - int32(1)
	if v777 < int32(2) {
		v829 = v809
		goto L171
	} else {
		goto L186
	}
L176:
	;
	v789 = v786
	goto L178
L177:
	;
	v789 = v787
	goto L178
L178:
	;
	v790 = F_memcmp(m, v773+v782, v774+v782, v789)
	mBase = m.M
	if v790 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if v786 == v787 {
		goto L175
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v790 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v794 = int32(10)
	v850 = (v777*v794 + v794) * (v786 - v787)
	goto L170
L183:
	;
	v806 = int32(-10)
	goto L185
L184:
	;
	v806 = int32(10)
	goto L185
L185:
	;
	v850 = (v777 + int32(1)) * v806
	goto L170
L186:
	;
	v812 = int32(9)
	v814 = int32(_a_F_ltree_consistent_0)
	v822 = int32(1)
	if v822 < v780 {
		v773 = v773 + (v786+v812)&v814
		v774 = v774 + (v787+v812)&v814
		v777 = v809
		v780 = v780 - v822
		goto L173
	} else {
		goto L187
	}
L187:
	;
	goto L174
L188:
	;
	v870 = int32(0)
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v727)+4)))
	v880 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v869)+4)))
	if base.B2i32(v877 == v870)|base.B2i32(v880 == v870) != 0 {
		v944 = v877
		goto L195
	} else {
		goto L196
	}
L189:
	;
	if v729&int32(2) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v858 = int32(0)
	goto L192
L191:
	;
	v858 = v51
	goto L192
L192:
	;
	v861 = v52 + v858 + int32(8)
	if v729&int32(4) != 0 {
		v869 = v861
		goto L188
	} else {
		goto L193
	}
L193:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v869 = v861 + int32(base.Ui32(v864)>>(uint(int32(2))%32))
	goto L188
L194:
	;
	v2565 = v727
	v2568 = base.B2i32(v965 <= int32(0))
	goto L13
L195:
	;
	v965 = (v944 + int32(1)) * (v877 - v880) * int32(10)
	goto L194
L196:
	;
	v884 = int32(8)
	v888 = v727 + v884
	v889 = v869 + v884
	v892 = v877
	v895 = v880
	goto L197
L197:
	;
	v897 = int32(2)
	v901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v888))))
	v902 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v889))))
	if base.Ui32(v901) < base.Ui32(v902) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v944 = v924
	goto L195
L199:
	;
	v924 = v892 - int32(1)
	if v892 < int32(2) {
		v944 = v924
		goto L195
	} else {
		goto L210
	}
L200:
	;
	v904 = v901
	goto L202
L201:
	;
	v904 = v902
	goto L202
L202:
	;
	v905 = F_memcmp(m, v888+v897, v889+v897, v904)
	mBase = m.M
	if v905 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	if v901 == v902 {
		goto L199
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	if v905 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v909 = int32(10)
	v965 = (v892*v909 + v909) * (v901 - v902)
	goto L194
L207:
	;
	v921 = int32(-10)
	goto L209
L208:
	;
	v921 = int32(10)
	goto L209
L209:
	;
	v965 = (v892 + int32(1)) * v921
	goto L194
L210:
	;
	v927 = int32(9)
	v929 = int32(_a_F_ltree_consistent_0)
	v937 = int32(1)
	if v937 < v895 {
		v888 = v888 + (v901+v927)&v929
		v889 = v889 + (v902+v927)&v929
		v892 = v924
		v895 = v895 - v937
		goto L197
	} else {
		goto L211
	}
L211:
	;
	goto L198
L212:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v971)+16)))
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971+v972)+12)))
	if v974&int32(1) != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)))
	if base.Ui32(v983) < base.Ui32(v982) {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	goto L215
L215:
	;
	v1030 = v52 + int32(8)
	v1031 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)))
	v1033 = v1031
	goto L232
L216:
	;
	v2565 = v969
	v2568 = v1028
	goto L13
L217:
	;
	v1028 = int32(0)
	goto L216
L218:
	;
	goto L219
L219:
	;
	if v982 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1028 = int32(1)
	goto L216
L221:
	;
	goto L222
L222:
	;
	v993 = v969 + int32(8)
	v994 = v52 + int32(16)
	v995 = v982
	goto L223
L223:
	;
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993))))
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v994))))
	if v998 != v999 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1028 = int32(1)
	goto L216
L225:
	;
	v1028 = int32(0)
	goto L216
L226:
	;
	goto L227
L227:
	;
	v1002 = int32(2)
	v1006 = F_memcmp(m, v993+v1002, v994+v1002, v998)
	mBase = m.M
	if v1006 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1028 = int32(0)
	goto L216
L229:
	;
	goto L230
L230:
	;
	v1008 = int32(9)
	v1010 = int32(_a_F_ltree_consistent_0)
	v1018 = int32(1)
	if v1018 < v995 {
		v993 = v993 + (v998+v1008)&v1010
		v994 = v994 + (v999+v1008)&v1010
		v995 = v995 - v1018
		goto L223
	} else {
		goto L231
	}
L231:
	;
	goto L224
L232:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)) = uint16(v1033)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1050&int32(3) != 0 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)) = uint16(v1031)
	v2565 = v969
	v2568 = v1277
	goto L13
L234:
	;
	goto L233
L235:
	;
	v1053 = int32(0)
	goto L237
L236:
	;
	v1053 = v51
	goto L237
L237:
	;
	v1054 = v1030 + v1053
	v1055 = int32(0)
	v1062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)))
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1054)+4)))
	if base.B2i32(v1062 == v1055)|base.B2i32(v1065 == v1055) != 0 {
		v1129 = v1062
		goto L239
	} else {
		goto L240
	}
L238:
	;
	if int32(0) <= v1150 {
		goto L256
	} else {
		goto L257
	}
L239:
	;
	v1150 = (v1129 + int32(1)) * (v1062 - v1065) * int32(10)
	goto L238
L240:
	;
	v1069 = int32(8)
	v1073 = v969 + v1069
	v1074 = v1054 + v1069
	v1077 = v1062
	v1080 = v1065
	goto L241
L241:
	;
	v1082 = int32(2)
	v1086 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1073))))
	v1087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1074))))
	if base.Ui32(v1086) < base.Ui32(v1087) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v1129 = v1109
	goto L239
L243:
	;
	v1109 = v1077 - int32(1)
	if v1077 < int32(2) {
		v1129 = v1109
		goto L239
	} else {
		goto L254
	}
L244:
	;
	v1089 = v1086
	goto L246
L245:
	;
	v1089 = v1087
	goto L246
L246:
	;
	v1090 = F_memcmp(m, v1073+v1082, v1074+v1082, v1089)
	mBase = m.M
	if v1090 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	if v1086 == v1087 {
		goto L243
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	if v1090 < int32(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1094 = int32(10)
	v1150 = (v1077*v1094 + v1094) * (v1086 - v1087)
	goto L238
L251:
	;
	v1106 = int32(-10)
	goto L253
L252:
	;
	v1106 = int32(10)
	goto L253
L253:
	;
	v1150 = (v1077 + int32(1)) * v1106
	goto L238
L254:
	;
	v1112 = int32(9)
	v1114 = int32(_a_F_ltree_consistent_0)
	v1122 = int32(1)
	if v1122 < v1080 {
		v1073 = v1073 + (v1086+v1112)&v1114
		v1074 = v1074 + (v1087+v1112)&v1114
		v1077 = v1109
		v1080 = v1080 - v1122
		goto L241
	} else {
		goto L255
	}
L255:
	;
	goto L242
L256:
	;
	v1153 = int32(1)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1154&v1153 != 0 {
		v1169 = v1030
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v1271 = int32(0)
	if v1271 < v1033 {
		v1033 = v1033 - int32(1)
		goto L232
	} else {
		goto L284
	}
L259:
	;
	v1170 = int32(0)
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)))
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169)+4)))
	if base.B2i32(v1177 == v1170)|base.B2i32(v1180 == v1170) != 0 {
		v1244 = v1177
		goto L266
	} else {
		goto L267
	}
L260:
	;
	if v1154&int32(2) != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1160 = int32(0)
	goto L263
L262:
	;
	v1160 = v51
	goto L263
L263:
	;
	v1161 = v1030 + v1160
	if v1154&int32(4) != 0 {
		v1169 = v1161
		goto L259
	} else {
		goto L264
	}
L264:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1161)))
	v1169 = v1161 + int32(base.Ui32(v1164)>>(uint(int32(2))%32))
	goto L259
L265:
	;
	if v1265 <= int32(0) {
		v1277 = v1153
		goto L234
	} else {
		goto L283
	}
L266:
	;
	v1265 = (v1244 + int32(1)) * (v1177 - v1180) * int32(10)
	goto L265
L267:
	;
	v1184 = int32(8)
	v1188 = v969 + v1184
	v1189 = v1169 + v1184
	v1192 = v1177
	v1195 = v1180
	goto L268
L268:
	;
	v1197 = int32(2)
	v1201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1188))))
	v1202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1189))))
	if base.Ui32(v1201) < base.Ui32(v1202) {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	v1244 = v1224
	goto L266
L270:
	;
	v1224 = v1192 - int32(1)
	if v1192 < int32(2) {
		v1244 = v1224
		goto L266
	} else {
		goto L281
	}
L271:
	;
	v1204 = v1201
	goto L273
L272:
	;
	v1204 = v1202
	goto L273
L273:
	;
	v1205 = F_memcmp(m, v1188+v1197, v1189+v1197, v1204)
	mBase = m.M
	if v1205 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	if v1201 == v1202 {
		goto L270
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	if v1205 < int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1209 = int32(10)
	v1265 = (v1192*v1209 + v1209) * (v1201 - v1202)
	goto L265
L278:
	;
	v1221 = int32(-10)
	goto L280
L279:
	;
	v1221 = int32(10)
	goto L280
L280:
	;
	v1265 = (v1192 + int32(1)) * v1221
	goto L265
L281:
	;
	v1227 = int32(9)
	v1229 = int32(_a_F_ltree_consistent_0)
	v1237 = int32(1)
	if v1237 < v1195 {
		v1188 = v1188 + (v1201+v1227)&v1229
		v1189 = v1189 + (v1202+v1227)&v1229
		v1192 = v1224
		v1195 = v1195 - v1237
		goto L268
	} else {
		goto L282
	}
L282:
	;
	goto L269
L283:
	;
	goto L258
L284:
	;
	v1277 = v1271
	goto L234
L285:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1284)+16)))
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1284+v1285)+12)))
	if v1287&int32(1) != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+4)))
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	if base.Ui32(v1296) < base.Ui32(v1295) {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	goto L288
L288:
	;
	v1343 = v52 + int32(8)
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1345&int32(3) != 0 {
		goto L305
	} else {
		goto L306
	}
L289:
	;
	v2565 = v1282
	v2568 = v1341
	goto L13
L290:
	;
	v1341 = int32(0)
	goto L289
L291:
	;
	goto L292
L292:
	;
	if v1295 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1341 = int32(1)
	goto L289
L294:
	;
	goto L295
L295:
	;
	v1306 = v52 + int32(16)
	v1307 = v1282 + int32(8)
	v1308 = v1295
	goto L296
L296:
	;
	v1311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1306))))
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1307))))
	if v1311 != v1312 {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1341 = int32(1)
	goto L289
L298:
	;
	v1341 = int32(0)
	goto L289
L299:
	;
	goto L300
L300:
	;
	v1315 = int32(2)
	v1319 = F_memcmp(m, v1306+v1315, v1307+v1315, v1311)
	mBase = m.M
	if v1319 != 0 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1341 = int32(0)
	goto L289
L302:
	;
	goto L303
L303:
	;
	v1321 = int32(9)
	v1323 = int32(_a_F_ltree_consistent_0)
	v1331 = int32(1)
	if v1331 < v1308 {
		v1306 = v1306 + (v1311+v1321)&v1323
		v1307 = v1307 + (v1312+v1321)&v1323
		v1308 = v1308 - v1331
		goto L296
	} else {
		goto L304
	}
L304:
	;
	goto L297
L305:
	;
	v1348 = int32(0)
	goto L307
L306:
	;
	v1348 = v51
	goto L307
L307:
	;
	v1349 = v1343 + v1348
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	v1353 = F_palloc0(m, int32(base.Ui32(v1350)>>(uint(int32(2))%32)))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L10
	} else {
		goto L308
	}
L308:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1349)))
	v1357 = int32(base.Ui32(v1355) >> (uint(int32(2)) % 32))
	if v1357 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	base.MemoryCopy(m, v1353, v1349, v1357)
	goto L311
L310:
	;
	goto L311
L311:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1359&int32(1) != 0 {
		v1373 = v1343
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	v1377 = F_palloc0(m, int32(base.Ui32(v1374)>>(uint(int32(2))%32)))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L10
	} else {
		goto L318
	}
L313:
	;
	if v1359&int32(2) != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1365 = int32(0)
	goto L316
L315:
	;
	v1365 = v51
	goto L316
L316:
	;
	v1366 = v1343 + v1365
	if v1359&int32(4) != 0 {
		v1373 = v1366
		goto L312
	} else {
		goto L317
	}
L317:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1366)))
	v1373 = v1366 + int32(base.Ui32(v1369)>>(uint(int32(2))%32))
	goto L312
L318:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	v1381 = int32(base.Ui32(v1379) >> (uint(int32(2)) % 32))
	if v1381 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	base.MemoryCopy(m, v1377, v1373, v1381)
	goto L321
L320:
	;
	goto L321
L321:
	;
	v1383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+4)))
	v1384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1353)+4)))
	if base.Ui32(v1383) < base.Ui32(v1384) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1353)+4)) = uint16(v1383)
	goto L324
L323:
	;
	goto L324
L324:
	;
	v1387 = int32(0)
	v1394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+4)))
	v1397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1353)+4)))
	if base.B2i32(v1394 == v1387)|base.B2i32(v1397 == v1387) != 0 {
		v1461 = v1394
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v1483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+4)))
	v1484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+4)))
	if base.Ui32(v1483) < base.Ui32(v1484) {
		goto L343
	} else {
		goto L344
	}
L326:
	;
	v1482 = (v1461 + int32(1)) * (v1394 - v1397) * int32(10)
	goto L325
L327:
	;
	v1401 = int32(8)
	v1405 = v1282 + v1401
	v1406 = v1353 + v1401
	v1409 = v1394
	v1412 = v1397
	goto L328
L328:
	;
	v1414 = int32(2)
	v1418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1405))))
	v1419 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1406))))
	if base.Ui32(v1418) < base.Ui32(v1419) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	v1461 = v1441
	goto L326
L330:
	;
	v1441 = v1409 - int32(1)
	if v1409 < int32(2) {
		v1461 = v1441
		goto L326
	} else {
		goto L341
	}
L331:
	;
	v1421 = v1418
	goto L333
L332:
	;
	v1421 = v1419
	goto L333
L333:
	;
	v1422 = F_memcmp(m, v1405+v1414, v1406+v1414, v1421)
	mBase = m.M
	if v1422 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	if v1418 == v1419 {
		goto L330
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	if v1422 < int32(0) {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v1426 = int32(10)
	v1482 = (v1409*v1426 + v1426) * (v1418 - v1419)
	goto L325
L338:
	;
	v1438 = int32(-10)
	goto L340
L339:
	;
	v1438 = int32(10)
	goto L340
L340:
	;
	v1482 = (v1409 + int32(1)) * v1438
	goto L325
L341:
	;
	v1444 = int32(9)
	v1446 = int32(_a_F_ltree_consistent_0)
	v1454 = int32(1)
	if v1454 < v1412 {
		v1405 = v1405 + (v1418+v1444)&v1446
		v1406 = v1406 + (v1419+v1444)&v1446
		v1409 = v1441
		v1412 = v1412 - v1454
		goto L328
	} else {
		goto L342
	}
L342:
	;
	goto L329
L343:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1377)+4)) = uint16(v1483)
	goto L345
L344:
	;
	goto L345
L345:
	;
	if int32(0) <= v1482 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1489 = int32(0)
	v1496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+4)))
	v1499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1377)+4)))
	if base.B2i32(v1496 == v1489)|base.B2i32(v1499 == v1489) != 0 {
		v1563 = v1496
		goto L350
	} else {
		goto L351
	}
L347:
	;
	v1587 = v2
	goto L348
L348:
	;
	F_pfree(m, v1353)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L10
	} else {
		goto L367
	}
L349:
	;
	v1587 = base.B2i32(v1584 <= int32(0))
	goto L348
L350:
	;
	v1584 = (v1563 + int32(1)) * (v1496 - v1499) * int32(10)
	goto L349
L351:
	;
	v1503 = int32(8)
	v1507 = v1282 + v1503
	v1508 = v1377 + v1503
	v1511 = v1496
	v1514 = v1499
	goto L352
L352:
	;
	v1516 = int32(2)
	v1520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1507))))
	v1521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1508))))
	if base.Ui32(v1520) < base.Ui32(v1521) {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	v1563 = v1543
	goto L350
L354:
	;
	v1543 = v1511 - int32(1)
	if v1511 < int32(2) {
		v1563 = v1543
		goto L350
	} else {
		goto L365
	}
L355:
	;
	v1523 = v1520
	goto L357
L356:
	;
	v1523 = v1521
	goto L357
L357:
	;
	v1524 = F_memcmp(m, v1507+v1516, v1508+v1516, v1523)
	mBase = m.M
	if v1524 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	if v1520 == v1521 {
		goto L354
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	if v1524 < int32(0) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1528 = int32(10)
	v1584 = (v1511*v1528 + v1528) * (v1520 - v1521)
	goto L349
L362:
	;
	v1540 = int32(-10)
	goto L364
L363:
	;
	v1540 = int32(10)
	goto L364
L364:
	;
	v1584 = (v1511 + int32(1)) * v1540
	goto L349
L365:
	;
	v1546 = int32(9)
	v1548 = int32(_a_F_ltree_consistent_0)
	v1556 = int32(1)
	if v1556 < v1514 {
		v1507 = v1507 + (v1520+v1546)&v1548
		v1508 = v1508 + (v1521+v1546)&v1548
		v1511 = v1543
		v1514 = v1514 - v1556
		goto L352
	} else {
		goto L366
	}
L366:
	;
	goto L353
L367:
	;
	F_pfree(m, v1377)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L10
	} else {
		goto L368
	}
L368:
	;
	v2565 = v1282
	v2568 = v1587
	goto L13
L369:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1597)+16)))
	v1600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1597+v1598)+12)))
	if v1600&int32(1) != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1605 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_1), int32(0), v1593, v1595)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L10
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v1609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v1609&int32(2) != 0 {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	v2565 = v1595
	v2568 = base.B2i32(v1605 != int32(0))
	goto L13
L374:
	;
	v1735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1595)+6)))
	if v1735 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L375:
	;
	v1612 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1595)+4)))
	if v1612 == int32(0) {
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1620 = v1612
	v1626 = v1595 + int32(16)
	goto L377
L377:
	;
	v1635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1626)+4)))
	if v1635 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	goto L374
L379:
	;
	v1698 = int32(1)
	v1700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1626))))
	if v1698 < v1620 {
		v1620 = v1620 - v1698
		v1626 = v1626 + (v1700+int32(7))&int32(_a_F_ltree_consistent_0)
		goto L377
	} else {
		goto L386
	}
L380:
	;
	v1638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626)+2)))
	if v1638&int32(21) != 0 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1644 = v1626 + int32(16)
	v1651 = v1635
	goto L382
L382:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1644)))
	v1660 = base.I32_rem_u_s(v1659, v51<<(uint(int32(3))%32))
	v1664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1593+int32(base.Ui32(v1660)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v1664)>>(uint(v1660&int32(7))%32))&int32(1) != 0 {
		goto L379
	} else {
		goto L384
	}
L383:
	;
	v2565 = v1595
	v2568 = v2
	goto L13
L384:
	;
	v1670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1644)+4)))
	v1678 = int32(1)
	if v1678 < v1651 {
		v1644 = v1644 + (v1670+int32(7))&int32(_a_F_ltree_consistent_0) + int32(8)
		v1651 = v1651 - v1678
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	goto L378
L387:
	;
	v2565 = v1595
	v2568 = v1955
	goto L13
L388:
	;
	v1955 = int32(1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1740 = v52 + int32(8)
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1742&int32(2) != 0 {
		goto L393
	} else {
		goto L394
	}
L391:
	;
	if v1822 <= int32(0) {
		goto L414
	} else {
		goto L415
	}
L392:
	;
	if base.Ui32(v1750) < base.Ui32(v1735) {
		goto L411
	} else {
		goto L412
	}
L393:
	;
	v1745 = int32(0)
	goto L395
L394:
	;
	v1745 = v51
	goto L395
L395:
	;
	v1746 = v1740 + v1745
	v1748 = v1742 & int32(1)
	if v1748 != 0 {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1749 = v1740
	goto L398
L397:
	;
	v1749 = v1746
	goto L398
L398:
	;
	v1750 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1749)+4)))
	if v1750 == int32(0) {
		goto L392
	} else {
		goto L399
	}
L399:
	;
	v1757 = v1595 + int32(16)
	v1759 = v1749 + int32(8)
	v1764 = v1735
	v1765 = v1750
	goto L400
L400:
	;
	v1775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759))))
	v1776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1757)+20)))
	if base.Ui32(v1775) < base.Ui32(v1776) {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	goto L392
L402:
	;
	v1778 = v1775
	goto L404
L403:
	;
	v1778 = v1776
	goto L404
L404:
	;
	v1779 = F_memcmp(m, v1759+int32(2), v1757+int32(23), v1778)
	mBase = m.M
	if v1779 != 0 {
		v1822 = v1779
		goto L391
	} else {
		goto L405
	}
L405:
	;
	if v1775 != v1776 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1822 = v1775 - v1776
	goto L391
L407:
	;
	goto L408
L408:
	;
	if v1765 < int32(2) {
		goto L392
	} else {
		goto L409
	}
L409:
	;
	v1784 = int32(1)
	v1788 = int32(_a_F_ltree_consistent_0)
	v1791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1757))))
	if v1784 < v1764 {
		v1757 = v1757 + (v1791+int32(7))&v1788
		v1759 = v1759 + (v1775+int32(9))&v1788
		v1764 = v1764 - v1784
		v1765 = v1765 - v1784
		goto L400
	} else {
		goto L410
	}
L410:
	;
	goto L401
L411:
	;
	v1816 = v1750
	goto L413
L412:
	;
	v1816 = v1735
	goto L413
L413:
	;
	v1822 = v1816 - v1735
	goto L391
L414:
	;
	if v1748 != 0 {
		v1840 = v1740
		goto L417
	} else {
		goto L418
	}
L415:
	;
	v1940 = int32(0)
	goto L416
L416:
	;
	v1955 = v1940
	goto L387
L417:
	;
	v1841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840)+4)))
	if v1841 == int32(0) {
		goto L423
	} else {
		goto L424
	}
L418:
	;
	if v1742&int32(4) != 0 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1840 = v1746
	goto L417
L420:
	;
	goto L421
L421:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1746)))
	v1840 = v1746 + int32(base.Ui32(v1836)>>(uint(int32(2))%32))
	goto L417
L422:
	;
	v1940 = base.B2i32(int32(0) <= v1913)
	goto L416
L423:
	;
	if base.Ui32(v1841) < base.Ui32(v1735) {
		goto L436
	} else {
		goto L437
	}
L424:
	;
	v1848 = v1595 + int32(16)
	v1850 = v1840 + int32(8)
	v1855 = v1735
	v1856 = v1841
	goto L425
L425:
	;
	v1866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1850))))
	v1867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1848)+20)))
	if base.Ui32(v1866) < base.Ui32(v1867) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	goto L423
L427:
	;
	v1869 = v1866
	goto L429
L428:
	;
	v1869 = v1867
	goto L429
L429:
	;
	v1870 = F_memcmp(m, v1850+int32(2), v1848+int32(23), v1869)
	mBase = m.M
	if v1870 != 0 {
		v1913 = v1870
		goto L422
	} else {
		goto L430
	}
L430:
	;
	if v1866 != v1867 {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1913 = v1866 - v1867
	goto L422
L432:
	;
	goto L433
L433:
	;
	if v1856 < int32(2) {
		goto L423
	} else {
		goto L434
	}
L434:
	;
	v1875 = int32(1)
	v1879 = int32(_a_F_ltree_consistent_0)
	v1882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1848))))
	if v1875 < v1855 {
		v1848 = v1848 + (v1882+int32(7))&v1879
		v1850 = v1850 + (v1866+int32(9))&v1879
		v1855 = v1855 - v1875
		v1856 = v1856 - v1875
		goto L425
	} else {
		goto L435
	}
L435:
	;
	goto L426
L436:
	;
	v1907 = v1841
	goto L438
L437:
	;
	v1907 = v1735
	goto L438
L438:
	;
	v1913 = v1907 - v1735
	goto L422
L439:
	;
	v1959 = int32(1)
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1960)+16)))
	v1963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1960+v1961)+12)))
	if v1963&v1959 != 0 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1970 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_2), int32(0), v52+int32(8), v1957)
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L10
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	v1974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v1974&int32(2) != 0 {
		v2565 = v1957
		v2568 = v1959
		goto L13
	} else {
		goto L444
	}
L443:
	;
	v2565 = v1957
	v2568 = base.B2i32(v1970 != int32(0))
	goto L13
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v51
	v1978 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v52 + v1978
	v1987 = F_ltree_execute(m, v1957+v1978, v19+v1978, int32(0), int32(_a_F_ltree_consistent_3))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L10
	} else {
		goto L445
	}
L445:
	;
	v2565 = v1957
	v2568 = v1987
	goto L13
L446:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1992)+16)))
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1992+v1993)+12)))
	if v1995&int32(1) != 0 {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2002 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_4), int32(0), v52+int32(8), v1990)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L10
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+8))
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+4))
	v2010 = F_ArrayGetNItemsSafe(m, v2007, v1990+int32(16))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L10
	} else {
		goto L451
	}
L450:
	;
	v2565 = v1990
	v2568 = base.B2i32(v2002 != int32(0))
	goto L13
L451:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+4))
	if v2012 < int32(2) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v2015 = F_array_contains_nulls(m, v1990)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L10
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L10
	} else {
		goto L537
	}
L455:
	;
	if v2015 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	if v2010 <= int32(0) {
		v2565 = v1990
		v2568 = v2
		goto L13
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L10
	} else {
		goto L533
	}
L459:
	;
	if v2006 != 0 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2027 = v2006
	goto L462
L461:
	;
	v2027 = (v2007<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L462
L462:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v2043 = v1990 + v2027
	v2047 = v2010
	goto L463
L463:
	;
	if v2033&int32(2) != 0 {
		goto L466
	} else {
		goto L467
	}
L464:
	;
	v2565 = v1990
	v2568 = v2
	goto L13
L465:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2043)))
	v2421 = int32(1)
	if v2421 < v2047 {
		v2043 = v2043 + (int32(base.Ui32(v2413)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v2047 = v2047 - v2421
		goto L463
	} else {
		goto L532
	}
L466:
	;
	v2173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2043)+6)))
	if v2173 == int32(0) {
		goto L480
	} else {
		goto L481
	}
L467:
	;
	v2052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2043)+4)))
	if v2052 == int32(0) {
		goto L466
	} else {
		goto L468
	}
L468:
	;
	v2058 = v2052
	v2066 = v2043 + int32(16)
	goto L469
L469:
	;
	v2073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2066)+4)))
	if v2073 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	goto L466
L471:
	;
	v2136 = int32(1)
	v2138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2066))))
	if v2136 < v2058 {
		v2058 = v2058 - v2136
		v2066 = v2066 + (v2138+int32(7))&int32(_a_F_ltree_consistent_0)
		goto L469
	} else {
		goto L478
	}
L472:
	;
	v2076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2066)+2)))
	if v2076&int32(21) != 0 {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v2082 = v2066 + int32(16)
	v2089 = v2073
	goto L474
L474:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2082)))
	v2098 = base.I32_rem_u_s(v2097, v51<<(uint(int32(3))%32))
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(8)+int32(base.Ui32(v2098)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v2102)>>(uint(v2098&int32(7))%32))&int32(1) != 0 {
		goto L471
	} else {
		goto L476
	}
L475:
	;
	goto L465
L476:
	;
	v2108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2082)+4)))
	v2116 = int32(1)
	if v2116 < v2089 {
		v2082 = v2082 + (v2108+int32(7))&int32(_a_F_ltree_consistent_0) + int32(8)
		v2089 = v2089 - v2116
		goto L474
	} else {
		goto L477
	}
L477:
	;
	goto L475
L478:
	;
	goto L470
L479:
	;
	if v2393 == int32(0) {
		goto L465
	} else {
		goto L531
	}
L480:
	;
	v2393 = int32(1)
	goto L479
L481:
	;
	goto L482
L482:
	;
	v2178 = v52 + int32(8)
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v2180&int32(2) != 0 {
		goto L485
	} else {
		goto L486
	}
L483:
	;
	if v2260 <= int32(0) {
		goto L506
	} else {
		goto L507
	}
L484:
	;
	if base.Ui32(v2188) < base.Ui32(v2173) {
		goto L503
	} else {
		goto L504
	}
L485:
	;
	v2183 = int32(0)
	goto L487
L486:
	;
	v2183 = v51
	goto L487
L487:
	;
	v2184 = v2178 + v2183
	v2186 = v2180 & int32(1)
	if v2186 != 0 {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v2187 = v2178
	goto L490
L489:
	;
	v2187 = v2184
	goto L490
L490:
	;
	v2188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2187)+4)))
	if v2188 == int32(0) {
		goto L484
	} else {
		goto L491
	}
L491:
	;
	v2195 = v2043 + int32(16)
	v2197 = v2187 + int32(8)
	v2202 = v2173
	v2203 = v2188
	goto L492
L492:
	;
	v2213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2197))))
	v2214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2195)+20)))
	if base.Ui32(v2213) < base.Ui32(v2214) {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	goto L484
L494:
	;
	v2216 = v2213
	goto L496
L495:
	;
	v2216 = v2214
	goto L496
L496:
	;
	v2217 = F_memcmp(m, v2197+int32(2), v2195+int32(23), v2216)
	mBase = m.M
	if v2217 != 0 {
		v2260 = v2217
		goto L483
	} else {
		goto L497
	}
L497:
	;
	if v2213 != v2214 {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v2260 = v2213 - v2214
	goto L483
L499:
	;
	goto L500
L500:
	;
	if v2203 < int32(2) {
		goto L484
	} else {
		goto L501
	}
L501:
	;
	v2222 = int32(1)
	v2226 = int32(_a_F_ltree_consistent_0)
	v2229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2195))))
	if v2222 < v2202 {
		v2195 = v2195 + (v2229+int32(7))&v2226
		v2197 = v2197 + (v2213+int32(9))&v2226
		v2202 = v2202 - v2222
		v2203 = v2203 - v2222
		goto L492
	} else {
		goto L502
	}
L502:
	;
	goto L493
L503:
	;
	v2254 = v2188
	goto L505
L504:
	;
	v2254 = v2173
	goto L505
L505:
	;
	v2260 = v2254 - v2173
	goto L483
L506:
	;
	if v2186 != 0 {
		v2278 = v2178
		goto L509
	} else {
		goto L510
	}
L507:
	;
	v2378 = int32(0)
	goto L508
L508:
	;
	v2393 = v2378
	goto L479
L509:
	;
	v2279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2278)+4)))
	if v2279 == int32(0) {
		goto L515
	} else {
		goto L516
	}
L510:
	;
	if v2180&int32(4) != 0 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v2278 = v2184
	goto L509
L512:
	;
	goto L513
L513:
	;
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v2184)))
	v2278 = v2184 + int32(base.Ui32(v2274)>>(uint(int32(2))%32))
	goto L509
L514:
	;
	v2378 = base.B2i32(int32(0) <= v2351)
	goto L508
L515:
	;
	if base.Ui32(v2279) < base.Ui32(v2173) {
		goto L528
	} else {
		goto L529
	}
L516:
	;
	v2286 = v2043 + int32(16)
	v2288 = v2278 + int32(8)
	v2293 = v2173
	v2294 = v2279
	goto L517
L517:
	;
	v2304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2288))))
	v2305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2286)+20)))
	if base.Ui32(v2304) < base.Ui32(v2305) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	goto L515
L519:
	;
	v2307 = v2304
	goto L521
L520:
	;
	v2307 = v2305
	goto L521
L521:
	;
	v2308 = F_memcmp(m, v2288+int32(2), v2286+int32(23), v2307)
	mBase = m.M
	if v2308 != 0 {
		v2351 = v2308
		goto L514
	} else {
		goto L522
	}
L522:
	;
	if v2304 != v2305 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v2351 = v2304 - v2305
	goto L514
L524:
	;
	goto L525
L525:
	;
	if v2294 < int32(2) {
		goto L515
	} else {
		goto L526
	}
L526:
	;
	v2313 = int32(1)
	v2317 = int32(_a_F_ltree_consistent_0)
	v2320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2286))))
	if v2313 < v2293 {
		v2286 = v2286 + (v2320+int32(7))&v2317
		v2288 = v2288 + (v2304+int32(9))&v2317
		v2293 = v2293 - v2313
		v2294 = v2294 - v2313
		goto L517
	} else {
		goto L527
	}
L527:
	;
	goto L518
L528:
	;
	v2345 = v2279
	goto L530
L529:
	;
	v2345 = v2173
	goto L530
L530:
	;
	v2351 = v2345 - v2173
	goto L514
L531:
	;
	v2565 = v1990
	v2568 = int32(1)
	goto L13
L532:
	;
	goto L464
L533:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L10
	} else {
		goto L534
	}
L534:
	;
	F_errmsg(m, int32(_a_F_ltree_consistent_5), int32(0))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L10
	} else {
		goto L535
	}
L535:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_6), int32(604), int32(_a_F_ltree_consistent_7))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L10
	} else {
		goto L536
	}
L536:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L537:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L10
	} else {
		goto L538
	}
L538:
	;
	F_errmsg(m, int32(_a_F_ltree_consistent_8), int32(0))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L10
	} else {
		goto L539
	}
L539:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_6), int32(600), int32(_a_F_ltree_consistent_7))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L10
	} else {
		goto L540
	}
L540:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L541:
	;
	v2461 = int32(0)
	goto L543
L542:
	;
	v2461 = v51
	goto L543
L543:
	;
	v2462 = v52 + v2461
	v2465 = int32(0)
	v2472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	v2475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2462+int32(8))+4)))
	if base.B2i32(v2472 == v2465)|base.B2i32(v2475 == v2465) != 0 {
		v2539 = v2472
		goto L545
	} else {
		goto L546
	}
L544:
	;
	v2565 = v58
	v2568 = base.B2i32(int32(0) <= v2560)
	goto L13
L545:
	;
	v2560 = (v2539 + int32(1)) * (v2472 - v2475) * int32(10)
	goto L544
L546:
	;
	v2483 = v58 + int32(8)
	v2484 = v2462 + int32(16)
	v2487 = v2472
	v2490 = v2475
	goto L547
L547:
	;
	v2492 = int32(2)
	v2496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2483))))
	v2497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2484))))
	if base.Ui32(v2496) < base.Ui32(v2497) {
		goto L550
	} else {
		goto L551
	}
L548:
	;
	v2539 = v2519
	goto L545
L549:
	;
	v2519 = v2487 - int32(1)
	if v2487 < int32(2) {
		v2539 = v2519
		goto L545
	} else {
		goto L560
	}
L550:
	;
	v2499 = v2496
	goto L552
L551:
	;
	v2499 = v2497
	goto L552
L552:
	;
	v2500 = F_memcmp(m, v2483+v2492, v2484+v2492, v2499)
	mBase = m.M
	if v2500 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	if v2496 == v2497 {
		goto L549
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	if v2500 < int32(0) {
		goto L557
	} else {
		goto L558
	}
L556:
	;
	v2504 = int32(10)
	v2560 = (v2487*v2504 + v2504) * (v2496 - v2497)
	goto L544
L557:
	;
	v2516 = int32(-10)
	goto L559
L558:
	;
	v2516 = int32(10)
	goto L559
L559:
	;
	v2560 = (v2487 + int32(1)) * v2516
	goto L544
L560:
	;
	v2522 = int32(9)
	v2524 = int32(_a_F_ltree_consistent_0)
	v2532 = int32(1)
	if v2532 < v2490 {
		v2483 = v2483 + (v2496+v2522)&v2524
		v2484 = v2484 + (v2497+v2522)&v2524
		v2487 = v2519
		v2490 = v2490 - v2532
		goto L547
	} else {
		goto L561
	}
L561:
	;
	goto L548
L562:
	;
	F_pfree(m, v2565)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L10
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	m.G0 = v19 + int32(16)
	return v2568
L565:
	;
	goto L564
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
	F_errmsg_internal(m, int32(_a_F_ltree_consistent_9), v19)
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L10
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_6), int32(715), int32(_a_F_ltree_consistent_10))
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L10
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ltree_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v24 = int32(0)
	if base.B2i32(v23 == v24)|base.B2i32(v22 == v24) != 0 {
		v152 = v23
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v174 != v15 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v172 = (v152 + int32(1)) * (v23 - v22) * int32(10)
	goto L4
L6:
	;
	v29 = int32(8)
	v37 = v15 + v29
	v38 = v20 + v29
	v39 = v23
	v46 = v22
	goto L7
L7:
	;
	v47 = int32(2)
	v48 = v37 + v47
	v50 = v38 + v47
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v51) < base.Ui32(v52) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v125
	goto L5
L9:
	;
	v54 = v51
	goto L11
L10:
	;
	v54 = v52
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v116 != 0 {
		v172 = int32(1)
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v116 = int32(0)
	goto L12
L14:
	;
	v90 = v85
	v91 = v86
	v92 = v87
	goto L24
L15:
	;
	if (v48|v50)&int32(3) != 0 {
		v85 = v48
		v86 = v50
		v87 = v54
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v78 = v48
	v79 = v50
	v80 = v54
	goto L17
L17:
	;
	if v80 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v62 = v48
	v63 = v50
	v64 = v54
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v67 != v68 {
		v85 = v62
		v86 = v63
		v87 = v64
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v78 = v73
	v79 = v71
	v80 = v75
	goto L17
L21:
	;
	v70 = int32(4)
	v71 = v63 + v70
	v73 = v62 + v70
	v75 = v64 - v70
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		v62 = v73
		v63 = v71
		v64 = v75
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v85 = v78
	v86 = v79
	v87 = v80
	goto L14
L24:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 == v96 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v116 = v95 - v96
	goto L12
L26:
	;
	v98 = int32(1)
	v103 = v92 - v98
	if v103 != 0 {
		v90 = v90 + v98
		v91 = v91 + v98
		v92 = v103
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	if v51 != v52 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = int32(10)
	v172 = (v39*v118 + v118) * (v51 - v52)
	goto L4
L32:
	;
	goto L33
L33:
	;
	v125 = v39 - int32(1)
	if v39 < int32(2) {
		v152 = v125
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v128 = int32(9)
	v130 = int32(_a_F_ltree_eq_0)
	v138 = int32(1)
	if v138 < v46 {
		v37 = v37 + (v51+v128)&v130
		v38 = v38 + (v52+v128)&v130
		v39 = v125
		v46 = v46 - v138
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v15)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v178 != v20 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v20)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.B2i32(v172 == int32(0))
L43:
	;
	goto L42
}
func F_ltree_gist_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_ltree_gist_out_0), int32(36), int32(_a_F_ltree_gist_out_1), int32(_a_F_ltree_gist_out_2), int32(_a_F_ltree_gist_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ltree_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		v152 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v14 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v179 = (v152 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	v28 = int32(8)
	v33 = v22
	v37 = v14 + v28
	v38 = v19 + v28
	v43 = v21
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v37 + v44
	v47 = v38 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v152 = v132
	goto L5
L9:
	;
	v132 = v33 - int32(1)
	if v33 < int32(2) {
		v152 = v132
		goto L5
	} else {
		goto L38
	}
L10:
	;
	v51 = v48
	goto L12
L11:
	;
	v51 = v49
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v113 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v113 = int32(0)
	goto L13
L15:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L25
L16:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L18
L18:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L20
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L18
L22:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L15
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = v92 - v93
	goto L13
L27:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v48 == v49 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v113 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(10)
	v179 = (v33*v117 + v117) * (v48 - v49)
	goto L4
L35:
	;
	v129 = int32(-10)
	goto L37
L36:
	;
	v129 = int32(10)
	goto L37
L37:
	;
	v179 = (v33 + int32(1)) * v129
	goto L4
L38:
	;
	v135 = int32(9)
	v137 = int32(_a_F_ltree_lt_0)
	v145 = int32(1)
	if v145 < v43 {
		v33 = v132
		v37 = v37 + (v48+v135)&v137
		v38 = v38 + (v49+v135)&v137
		v43 = v43 - v145
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L8
L40:
	;
	F_pfree(m, v14)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v184 != v19 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_pfree(m, v19)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	return int32(base.Ui32(v179) >> (uint(int32(31)) % 32))
L47:
	;
	goto L46
}
func F_ltree_risparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v129 != v10 {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	if base.Ui32(v18) < base.Ui32(v17) {
		v122 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v122 = int32(1)
	goto L3
L7:
	;
	goto L8
L8:
	;
	v23 = int32(8)
	v30 = v17
	v31 = v15 + v23
	v32 = v10 + v23
	goto L9
L9:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
	if v35 != v36 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v122 = v116
	goto L3
L11:
	;
	v122 = int32(0)
	goto L3
L12:
	;
	goto L13
L13:
	;
	v39 = int32(2)
	v40 = v32 + v39
	v42 = v31 + v39
	if base.Ui32(int32(4)) <= base.Ui32(v35) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v104 != 0 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v104 = int32(0)
	goto L14
L16:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L26
L17:
	;
	if (v40|v42)&int32(3) != 0 {
		v73 = v40
		v74 = v42
		v75 = v35
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v66 = v40
	v67 = v42
	v68 = v35
	goto L19
L19:
	;
	if v68 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v50 = v40
	v51 = v42
	v52 = v35
	goto L21
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L19
L23:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L16
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v104 = v83 - v84
	goto L14
L28:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	v122 = int32(0)
	goto L3
L33:
	;
	goto L34
L34:
	;
	v106 = int32(9)
	v108 = int32(_a_F_ltree_risparent_0)
	v116 = int32(1)
	if v116 < v30 {
		v30 = v30 - v116
		v31 = v31 + (v36+v106)&v108
		v32 = v32 + (v35+v106)&v108
		goto L9
	} else {
		goto L35
	}
L35:
	;
	goto L10
L36:
	;
	F_pfree(m, v10)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v133 != v15 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return v122
L43:
	;
	goto L42
}
