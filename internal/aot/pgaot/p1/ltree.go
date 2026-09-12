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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
		if v14 == int32(0) {
			v27 = v2
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != int32(7) {
				v27 = v2
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				if v20 != int32(17) {
					v27 = v2
				} else {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
					v27 = v23 ^ int32(1)
				}
			}
		}
	}
	if v27&int32(1) != 0 {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v31 = F_get_fn_opclass_options(m, v30)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = v35
			v37 = F_hemdist_2(m, v7, v9, v36)
			mBase = m.M
			*(*float32)(unsafe.Add(mBase, uint32(v5))) = base.F32_convert_i32_s(v37)
			return v5
		}
	} else {
		v36 = int32(28)
		v37 = F_hemdist_2(m, v7, v9, v36)
		mBase = m.M
		*(*float32)(unsafe.Add(mBase, uint32(v5))) = base.F32_convert_i32_s(v37)
		return v5
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v477 int32
	_ = v477
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v574 int32
	_ = v574
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
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
	var v620 int32
	_ = v620
	var v646 int32
	_ = v646
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v706 int32
	_ = v706
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v761 int32
	_ = v761
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
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
			v59 = (v55 + int32(65534)) & int32(65535)
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
						v82 = v2
						v84 = int32(-1)
						v85 = v2
						v86 = int32(1)
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
							v105 = v86 + int32(1)
							v106 = v105
							v112 = v82
							v113 = v105
							v114 = v84
							v115 = v85
							for {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v113<<(uint(int32(4))%32))))
								v134 = F_hemdist_2(m, v103, v133, v54)
								mBase = m.M
								v135 = base.B2i32(v114 < v134)
								if v114 < v134 {
									v136 = v134
								} else {
									v136 = v114
								}
								if v114 < v134 {
									v137 = v106
								} else {
									v137 = v112
								}
								if v114 < v134 {
									v138 = v86
								} else {
									v138 = v115
								}
								v140 = v106 + int32(1)
								v142 = v140 & int32(65535)
								if base.Ui32(v142) <= base.Ui32(v59) {
									v106 = v140
									v112 = v137
									v113 = v142
									v114 = v136
									v115 = v138
									continue
								} else {
									break
								}
								break
							}
							if v59 != v105 {
								v82 = v137
								v84 = v136
								v85 = v138
								v86 = v105
								continue
							} else {
								break
							}
							break
						}
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v152 = v137
						v155 = v138
						v161 = v145
					} else {
						v152 = v2
						v155 = v2
						v161 = v67
					}
					v170 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v170
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v170
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v176 = v26 + int32(4)
					v178 = int32(65535)
					v186 = base.B2i32(v155&v178 == v170) | base.B2i32(v152&v178 == v170)
					if v186 != 0 {
						v187 = int32(1)
					} else {
						v187 = v155
					}
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v176+v187&int32(65535)<<(uint(int32(4))%32))))
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
					v201 = int32(0)
					v203 = F_ltree_gist_alloc(m, int32(base.Ui32(v194&int32(2))>>(uint(int32(1))%32)), v193+int32(8), v54, v201, v201)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						if v186 != 0 {
							v206 = int32(2)
						} else {
							v206 = v152
						}
						v212 = *(*int32)(unsafe.Add(mBase, uint32(v176+v206&int32(65535)<<(uint(int32(4))%32))))
						v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
						v220 = int32(0)
						v222 = F_ltree_gist_alloc(m, int32(base.Ui32(v213&int32(2))>>(uint(int32(1))%32)), v212+int32(8), v54, v220, v220)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							v224 = int32(65535)
							v225 = v55 + v224
							v227 = v225 & v224
							v230 = F_palloc(m, v227<<(uint(int32(3))%32))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return int32(0)
							} else {
								if v55&int32(65535) == int32(1) {
									F_pg_qsort(m, v230, v227, int32(8), int32(6997))
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return int32(0)
									} else {
										v775 = v174
										v778 = v161
										v787 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v775))) = uint16(v787)
										*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v787)
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v222
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v203
										return v25
									}
								} else {
									v240 = int32(1)
									v242 = v240
									v250 = v240
									for {
										v268 = v230 + v242<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v268-int32(8)))) = uint16(v250)
										v272 = int32(4)
										v277 = *(*int32)(unsafe.Add(mBase, uint32(v176+v242<<(uint(v272)%32))))
										v278 = F_hemdist_2(m, v203, v277, v54)
										mBase = m.M
										v279 = F_hemdist_2(m, v222, v277, v54)
										mBase = m.M
										v280 = v278 - v279
										v282 = v280 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v268-v272))) = v280 ^ v282 - v282
										v287 = v250 + int32(1)
										v288 = int32(65535)
										v289 = v287 & v288
										if base.Ui32(v289) <= base.Ui32(v225&v288) {
											v242 = v289
											v250 = v287
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v230, v227, int32(8), int32(6997))
									mBase = m.M
									v296 = m.ExcPending
									if v296 != 0 {
										return int32(0)
									} else {
										v297 = int32(1)
										if base.Ui32(v227) <= base.Ui32(v297) {
											v300 = v297
										} else {
											v300 = v227
										}
										v302 = v54 & int32(2147483644)
										v304 = v54 & int32(3)
										v305 = int32(8)
										v306 = v222 + v305
										v308 = v203 + v305
										v320 = int32(0)
										v324 = v174
										v327 = v161
										for {
											v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+v320<<(uint(int32(3))%32)))))
											if v187&int32(65535) == v339 {
												*(*uint16)(unsafe.Add(mBase, uint32(v324))) = uint16(v187)
												v342 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v342 + int32(1)
												v748 = v324 + int32(2)
												v751 = v327
											} else {
												if v206&int32(65535) == v339 {
													*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v206)
													v352 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v352 + int32(1)
													v748 = v324
													v751 = v327 + int32(2)
												} else {
													v359 = *(*int32)(unsafe.Add(mBase, uint32(v176+v339<<(uint(int32(4))%32))))
													v360 = F_hemdist_2(m, v203, v359, v54)
													mBase = m.M
													v362 = F_hemdist_2(m, v222, v359, v54)
													mBase = m.M
													v364 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													v365 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v366 = v364 - v365
													if base.F64_lt(base.F64_convert_i32_s(v360), base.F64_add(base.F64_convert_i32_s(v362), base.F64_mul(base.F64_convert_i32_s(v366*v366*v366), float64(-1e-05)))) != 0 {
														v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)))
														if v374&int32(2) != 0 {
														} else {
															v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
															if v377&int32(2) != 0 {
																v382 = F__emscripten_memset_bulkmem(m, v308, base.I32_extend8_s(int32(255)), v54)
																mBase = m.M
															} else {
																if v54 <= int32(0) {
																} else {
																	v386 = v359 + int32(8)
																	v387 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v392 = v387
																		v405 = v387
																		for {
																			v416 = v392 + v308
																			v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
																			v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392+v386))))
																			v420 = v417 | v419
																			*(*uint8)(unsafe.Add(mBase, uint32(v416))) = uint8(v420)
																			v423 = v392 | int32(1)
																			v424 = v308 + v423
																			v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
																			v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v423))))
																			v428 = v425 | v427
																			*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v428)
																			v431 = v392 | int32(2)
																			v432 = v308 + v431
																			v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
																			v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v431))))
																			v436 = v433 | v435
																			*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v436)
																			v439 = v392 | int32(3)
																			v440 = v308 + v439
																			v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
																			v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v439))))
																			v444 = v441 | v443
																			*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v444)
																			v446 = int32(4)
																			v447 = v392 + v446
																			v449 = v405 + v446
																			if v449 != v302 {
																				v392 = v447
																				v405 = v449
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v451 = v447
																	} else {
																		v451 = v387
																	}
																	if v304 == int32(0) {
																	} else {
																		v477 = v451
																		v494 = v387
																		for {
																			v501 = v477 + v308
																			v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
																			v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v386))))
																			v505 = v502 | v504
																			*(*uint8)(unsafe.Add(mBase, uint32(v501))) = uint8(v505)
																			v507 = int32(1)
																			v510 = v494 + v507
																			if v510 != v304 {
																				v477 = v477 + v507
																				v494 = v510
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
														*(*uint16)(unsafe.Add(mBase, uint32(v324))) = uint16(v339)
														v537 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v537 + int32(1)
														v748 = v324 + int32(2)
														v751 = v327
													} else {
														v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
														if v543&int32(2) != 0 {
														} else {
															v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
															if v546&int32(2) != 0 {
																v551 = F__emscripten_memset_bulkmem(m, v306, base.I32_extend8_s(int32(255)), v54)
																mBase = m.M
															} else {
																if v54 <= int32(0) {
																} else {
																	v555 = v359 + int32(8)
																	v556 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v54) {
																		v561 = v556
																		v574 = v556
																		for {
																			v585 = v561 + v306
																			v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
																			v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561+v555))))
																			v589 = v586 | v588
																			*(*uint8)(unsafe.Add(mBase, uint32(v585))) = uint8(v589)
																			v592 = v561 | int32(1)
																			v593 = v306 + v592
																			v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
																			v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v592))))
																			v597 = v594 | v596
																			*(*uint8)(unsafe.Add(mBase, uint32(v593))) = uint8(v597)
																			v600 = v561 | int32(2)
																			v601 = v306 + v600
																			v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
																			v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v600))))
																			v605 = v602 | v604
																			*(*uint8)(unsafe.Add(mBase, uint32(v601))) = uint8(v605)
																			v608 = v561 | int32(3)
																			v609 = v306 + v608
																			v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
																			v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v608))))
																			v613 = v610 | v612
																			*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v613)
																			v615 = int32(4)
																			v616 = v561 + v615
																			v618 = v574 + v615
																			if v618 != v302 {
																				v561 = v616
																				v574 = v618
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v620 = v616
																	} else {
																		v620 = v556
																	}
																	if v304 == int32(0) {
																	} else {
																		v646 = v620
																		v663 = v556
																		for {
																			v670 = v646 + v306
																			v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
																			v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v555))))
																			v674 = v671 | v673
																			*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v674)
																			v676 = int32(1)
																			v679 = v663 + v676
																			if v679 != v304 {
																				v646 = v646 + v676
																				v663 = v679
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
														*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v339)
														v706 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v706 + int32(1)
														v748 = v324
														v751 = v327 + int32(2)
													}
												}
											}
											v761 = v320 + int32(1)
											if v761 != v300 {
												v320 = v761
												v324 = v748
												v327 = v751
												continue
											} else {
												break
											}
											break
										}
										v775 = v748
										v778 = v751
										v787 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v775))) = uint16(v787)
										*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v787)
										*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v222
										*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v203
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
		v59 = (v55 + int32(65534)) & int32(65535)
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
					v82 = v2
					v84 = int32(-1)
					v85 = v2
					v86 = int32(1)
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
						v105 = v86 + int32(1)
						v106 = v105
						v112 = v82
						v113 = v105
						v114 = v84
						v115 = v85
						for {
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v113<<(uint(int32(4))%32))))
							v134 = F_hemdist_2(m, v103, v133, v54)
							mBase = m.M
							v135 = base.B2i32(v114 < v134)
							if v114 < v134 {
								v136 = v134
							} else {
								v136 = v114
							}
							if v114 < v134 {
								v137 = v106
							} else {
								v137 = v112
							}
							if v114 < v134 {
								v138 = v86
							} else {
								v138 = v115
							}
							v140 = v106 + int32(1)
							v142 = v140 & int32(65535)
							if base.Ui32(v142) <= base.Ui32(v59) {
								v106 = v140
								v112 = v137
								v113 = v142
								v114 = v136
								v115 = v138
								continue
							} else {
								break
							}
							break
						}
						if v59 != v105 {
							v82 = v137
							v84 = v136
							v85 = v138
							v86 = v105
							continue
						} else {
							break
						}
						break
					}
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v152 = v137
					v155 = v138
					v161 = v145
				} else {
					v152 = v2
					v155 = v2
					v161 = v67
				}
				v170 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v170
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v170
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v176 = v26 + int32(4)
				v178 = int32(65535)
				v186 = base.B2i32(v155&v178 == v170) | base.B2i32(v152&v178 == v170)
				if v186 != 0 {
					v187 = int32(1)
				} else {
					v187 = v155
				}
				v193 = *(*int32)(unsafe.Add(mBase, uint32(v176+v187&int32(65535)<<(uint(int32(4))%32))))
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
				v201 = int32(0)
				v203 = F_ltree_gist_alloc(m, int32(base.Ui32(v194&int32(2))>>(uint(int32(1))%32)), v193+int32(8), v54, v201, v201)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					if v186 != 0 {
						v206 = int32(2)
					} else {
						v206 = v152
					}
					v212 = *(*int32)(unsafe.Add(mBase, uint32(v176+v206&int32(65535)<<(uint(int32(4))%32))))
					v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
					v220 = int32(0)
					v222 = F_ltree_gist_alloc(m, int32(base.Ui32(v213&int32(2))>>(uint(int32(1))%32)), v212+int32(8), v54, v220, v220)
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return int32(0)
					} else {
						v224 = int32(65535)
						v225 = v55 + v224
						v227 = v225 & v224
						v230 = F_palloc(m, v227<<(uint(int32(3))%32))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return int32(0)
						} else {
							if v55&int32(65535) == int32(1) {
								F_pg_qsort(m, v230, v227, int32(8), int32(6997))
								mBase = m.M
								v239 = m.ExcPending
								if v239 != 0 {
									return int32(0)
								} else {
									v775 = v174
									v778 = v161
									v787 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v775))) = uint16(v787)
									*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v787)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v222
									*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v203
									return v25
								}
							} else {
								v240 = int32(1)
								v242 = v240
								v250 = v240
								for {
									v268 = v230 + v242<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v268-int32(8)))) = uint16(v250)
									v272 = int32(4)
									v277 = *(*int32)(unsafe.Add(mBase, uint32(v176+v242<<(uint(v272)%32))))
									v278 = F_hemdist_2(m, v203, v277, v54)
									mBase = m.M
									v279 = F_hemdist_2(m, v222, v277, v54)
									mBase = m.M
									v280 = v278 - v279
									v282 = v280 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v268-v272))) = v280 ^ v282 - v282
									v287 = v250 + int32(1)
									v288 = int32(65535)
									v289 = v287 & v288
									if base.Ui32(v289) <= base.Ui32(v225&v288) {
										v242 = v289
										v250 = v287
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v230, v227, int32(8), int32(6997))
								mBase = m.M
								v296 = m.ExcPending
								if v296 != 0 {
									return int32(0)
								} else {
									v297 = int32(1)
									if base.Ui32(v227) <= base.Ui32(v297) {
										v300 = v297
									} else {
										v300 = v227
									}
									v302 = v54 & int32(2147483644)
									v304 = v54 & int32(3)
									v305 = int32(8)
									v306 = v222 + v305
									v308 = v203 + v305
									v320 = int32(0)
									v324 = v174
									v327 = v161
									for {
										v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+v320<<(uint(int32(3))%32)))))
										if v187&int32(65535) == v339 {
											*(*uint16)(unsafe.Add(mBase, uint32(v324))) = uint16(v187)
											v342 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v342 + int32(1)
											v748 = v324 + int32(2)
											v751 = v327
										} else {
											if v206&int32(65535) == v339 {
												*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v206)
												v352 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v352 + int32(1)
												v748 = v324
												v751 = v327 + int32(2)
											} else {
												v359 = *(*int32)(unsafe.Add(mBase, uint32(v176+v339<<(uint(int32(4))%32))))
												v360 = F_hemdist_2(m, v203, v359, v54)
												mBase = m.M
												v362 = F_hemdist_2(m, v222, v359, v54)
												mBase = m.M
												v364 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
												v365 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												v366 = v364 - v365
												if base.F64_lt(base.F64_convert_i32_s(v360), base.F64_add(base.F64_convert_i32_s(v362), base.F64_mul(base.F64_convert_i32_s(v366*v366*v366), float64(-1e-05)))) != 0 {
													v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)))
													if v374&int32(2) != 0 {
													} else {
														v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
														if v377&int32(2) != 0 {
															v382 = F__emscripten_memset_bulkmem(m, v308, base.I32_extend8_s(int32(255)), v54)
															mBase = m.M
														} else {
															if v54 <= int32(0) {
															} else {
																v386 = v359 + int32(8)
																v387 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v392 = v387
																	v405 = v387
																	for {
																		v416 = v392 + v308
																		v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
																		v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392+v386))))
																		v420 = v417 | v419
																		*(*uint8)(unsafe.Add(mBase, uint32(v416))) = uint8(v420)
																		v423 = v392 | int32(1)
																		v424 = v308 + v423
																		v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
																		v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v423))))
																		v428 = v425 | v427
																		*(*uint8)(unsafe.Add(mBase, uint32(v424))) = uint8(v428)
																		v431 = v392 | int32(2)
																		v432 = v308 + v431
																		v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
																		v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v431))))
																		v436 = v433 | v435
																		*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v436)
																		v439 = v392 | int32(3)
																		v440 = v308 + v439
																		v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
																		v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386+v439))))
																		v444 = v441 | v443
																		*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v444)
																		v446 = int32(4)
																		v447 = v392 + v446
																		v449 = v405 + v446
																		if v449 != v302 {
																			v392 = v447
																			v405 = v449
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v451 = v447
																} else {
																	v451 = v387
																}
																if v304 == int32(0) {
																} else {
																	v477 = v451
																	v494 = v387
																	for {
																		v501 = v477 + v308
																		v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
																		v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v386))))
																		v505 = v502 | v504
																		*(*uint8)(unsafe.Add(mBase, uint32(v501))) = uint8(v505)
																		v507 = int32(1)
																		v510 = v494 + v507
																		if v510 != v304 {
																			v477 = v477 + v507
																			v494 = v510
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
													*(*uint16)(unsafe.Add(mBase, uint32(v324))) = uint16(v339)
													v537 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v537 + int32(1)
													v748 = v324 + int32(2)
													v751 = v327
												} else {
													v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
													if v543&int32(2) != 0 {
													} else {
														v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
														if v546&int32(2) != 0 {
															v551 = F__emscripten_memset_bulkmem(m, v306, base.I32_extend8_s(int32(255)), v54)
															mBase = m.M
														} else {
															if v54 <= int32(0) {
															} else {
																v555 = v359 + int32(8)
																v556 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v54) {
																	v561 = v556
																	v574 = v556
																	for {
																		v585 = v561 + v306
																		v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
																		v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561+v555))))
																		v589 = v586 | v588
																		*(*uint8)(unsafe.Add(mBase, uint32(v585))) = uint8(v589)
																		v592 = v561 | int32(1)
																		v593 = v306 + v592
																		v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
																		v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v592))))
																		v597 = v594 | v596
																		*(*uint8)(unsafe.Add(mBase, uint32(v593))) = uint8(v597)
																		v600 = v561 | int32(2)
																		v601 = v306 + v600
																		v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
																		v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v600))))
																		v605 = v602 | v604
																		*(*uint8)(unsafe.Add(mBase, uint32(v601))) = uint8(v605)
																		v608 = v561 | int32(3)
																		v609 = v306 + v608
																		v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
																		v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v608))))
																		v613 = v610 | v612
																		*(*uint8)(unsafe.Add(mBase, uint32(v609))) = uint8(v613)
																		v615 = int32(4)
																		v616 = v561 + v615
																		v618 = v574 + v615
																		if v618 != v302 {
																			v561 = v616
																			v574 = v618
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v620 = v616
																} else {
																	v620 = v556
																}
																if v304 == int32(0) {
																} else {
																	v646 = v620
																	v663 = v556
																	for {
																		v670 = v646 + v306
																		v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
																		v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v555))))
																		v674 = v671 | v673
																		*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v674)
																		v676 = int32(1)
																		v679 = v663 + v676
																		if v679 != v304 {
																			v646 = v646 + v676
																			v663 = v679
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
													*(*uint16)(unsafe.Add(mBase, uint32(v327))) = uint16(v339)
													v706 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v706 + int32(1)
													v748 = v324
													v751 = v327 + int32(2)
												}
											}
										}
										v761 = v320 + int32(1)
										if v761 != v300 {
											v320 = v761
											v324 = v748
											v327 = v751
											continue
										} else {
											break
										}
										break
									}
									v775 = v748
									v778 = v751
									v787 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v775))) = uint16(v787)
									*(*uint16)(unsafe.Add(mBase, uint32(v778))) = uint16(v787)
									*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v222
									*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v203
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5594), int32(0), v4, v5)
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
	var v75 int32
	_ = v75
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
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v252 int32
	_ = v252
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
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(base.Ui32(v252) >> (uint(int32(2)) % 32))
	return v51
L14:
	;
	v57 = v51 + int32(8)
	v61 = v46 & int32(3)
	v66 = v53
	v75 = v2
	goto L15
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(4)+v75<<(uint(int32(4))%32))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v86&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v232 | int32(2)
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
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v211 = v66
	goto L22
L22:
	;
	v228 = v75 + int32(1)
	if v228 < v211 {
		v66 = v211
		v75 = v228
		goto L15
	} else {
		goto L35
	}
L23:
	;
	v100 = v95
	v107 = v95
	goto L26
L24:
	;
	v151 = v95
	goto L25
L25:
	;
	if v61 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
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
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v151 = v147
	goto L25
L28:
	;
	goto L27
L29:
	;
	v167 = v151
	v178 = v95
	goto L32
L30:
	;
	goto L31
L31:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v211 = v210
	goto L22
L32:
	;
	v183 = v167 + v57
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v94))))
	v187 = v184 | v186
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v187)
	v189 = int32(1)
	v192 = v178 + v189
	if v192 != v61 {
		v167 = v167 + v189
		v178 = v192
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L31
L34:
	;
	goto L33
L35:
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
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
	if v22 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v179 != v14 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v178 = (v149 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	if v21 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(8)
	v33 = v14 + v27
	v35 = v22
	v38 = v19 + v27
	v42 = v21
	goto L8
L8:
	;
	v43 = int32(2)
	v44 = v33 + v43
	v46 = v38 + v43
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v47) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v149 = v131
	goto L5
L10:
	;
	v131 = v35 - int32(1)
	if v35 < int32(2) {
		v149 = v131
		goto L5
	} else {
		goto L39
	}
L11:
	;
	v50 = v47
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v112 = int32(0)
	goto L14
L16:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L26
L17:
	;
	if (v44|v46)&int32(3) != 0 {
		v81 = v44
		v82 = v46
		v83 = v50
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v74 = v44
	v75 = v46
	v76 = v50
	goto L19
L19:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v58 = v44
	v59 = v46
	v60 = v50
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L19
L23:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L16
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = v91 - v92
	goto L14
L28:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
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
	if v47 == v48 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v112 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v116 = int32(10)
	v178 = (v35*v116 + v116) * (v47 - v48)
	goto L4
L36:
	;
	v128 = int32(-10)
	goto L38
L37:
	;
	v128 = int32(10)
	goto L38
L38:
	;
	v178 = (v35 + int32(1)) * v128
	goto L4
L39:
	;
	v134 = int32(9)
	v136 = int32(131064)
	v144 = int32(1)
	if v144 < v42 {
		v33 = v33 + (v47+v134)&v136
		v35 = v131
		v38 = v38 + (v48+v134)&v136
		v42 = v42 - v144
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L9
L41:
	;
	F_pfree(m, v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 != v19 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return v178
L48:
	;
	goto L47
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = v12 + v13
	if base.Ui32(int32(65536)) <= base.Ui32(v14) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(65535)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
				F_errmsg(m, int32(681844), v10)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497293), int32(353), int32(112878))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v42 = int32(2)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v50 = F_palloc0(m, int32(base.Ui32(v41)>>(uint(v42)%32))+int32(base.Ui32(v44)>>(uint(v42)%32))-int32(8))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v14)
			v55 = int32(-4)
			*(*int32)(unsafe.Add(mBase, uint32(v50))) = (v52+v53&v55)&v55 - int32(32)
			v63 = int32(8)
			v64 = v50 + v63
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v71 = int32(base.Ui32(v67)>>(uint(int32(2))%32)) - v63
			if v71 != 0 {
				v72 = F__emscripten_memcpy_bulkmem(m, v64, l0+v63, v71)
				mBase = m.M
				v73 = v72
			} else {
				v73 = v64
			}
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v75 = int32(2)
			v78 = int32(8)
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v86 = int32(base.Ui32(v82)>>(uint(v75)%32)) - v78
			if v86 != 0 {
				v87 = F__emscripten_memcpy_bulkmem(m, v73+int32(base.Ui32(v74)>>(uint(v75)%32))-v78, l1+v78, v86)
				mBase = m.M
			} else {
			}
			m.G0 = v10 + int32(16)
			return v50
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v843 int32
	_ = v843
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1321 int32
	_ = v1321
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1445 int32
	_ = v1445
	var v1451 int32
	_ = v1451
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1610 int32
	_ = v1610
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1668 int32
	_ = v1668
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1806 int32
	_ = v1806
	var v1817 int32
	_ = v1817
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1898 int32
	_ = v1898
	var v1909 int32
	_ = v1909
	var v1930 int32
	_ = v1930
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1968 int32
	_ = v1968
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2055 int32
	_ = v2055
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2092 int32
	_ = v2092
	var v2098 int32
	_ = v2098
	var v2106 int32
	_ = v2106
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2163 int32
	_ = v2163
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2244 int32
	_ = v2244
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2336 int32
	_ = v2336
	var v2347 int32
	_ = v2347
	var v2368 int32
	_ = v2368
	var v2383 int32
	_ = v2383
	var v2403 int32
	_ = v2403
	var v2411 int32
	_ = v2411
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2427 int32
	_ = v2427
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2447 int32
	_ = v2447
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2501 int32
	_ = v2501
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2529 int32
	_ = v2529
	var v2535 int32
	_ = v2535
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2587 int32
	_ = v2587
	var v2593 int32
	_ = v2593
	var v2600 int32
	_ = v2600
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
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L10
	} else {
		goto L581
	}
L13:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2576 != v2562 {
		goto L577
	} else {
		goto L578
	}
L14:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v2456&int32(3) != 0 {
		goto L555
	} else {
		goto L556
	}
L15:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1980 = F_pg_detoast_datum(m, v1979)
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L10
	} else {
		goto L460
	}
L16:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1947 = F_pg_detoast_datum(m, v1946)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L10
	} else {
		goto L453
	}
L17:
	;
	v1583 = v52 + int32(8)
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1585 = F_pg_detoast_datum(m, v1584)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L10
	} else {
		goto L383
	}
L18:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1272 = F_pg_detoast_datum(m, v1271)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L10
	} else {
		goto L295
	}
L19:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v961 = F_pg_detoast_datum_copy(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L10
	} else {
		goto L220
	}
L20:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v721 = F_pg_detoast_datum(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L10
	} else {
		goto L166
	}
L21:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v601 = F_pg_detoast_datum(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L10
	} else {
		goto L140
	}
L22:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v276 = F_pg_detoast_datum(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L69
	}
L23:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v168 = F_pg_detoast_datum(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L10
	} else {
		goto L46
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
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	if v78 == int32(0) {
		v142 = v78
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v2562 = v58
	v2565 = base.B2i32(int32(0) < v164)
	goto L13
L28:
	;
	v164 = (v142 + int32(1)) * (v78 - v77) * int32(10)
	goto L27
L29:
	;
	if v77 == int32(0) {
		v142 = v78
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v87 = v58 + int32(8)
	v88 = v52 + int32(16)
	v91 = v78
	v95 = v77
	goto L31
L31:
	;
	v96 = int32(2)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
	if base.Ui32(v100) < base.Ui32(v101) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v142 = v123
	goto L28
L33:
	;
	v123 = v91 - int32(1)
	if v91 < int32(2) {
		v142 = v123
		goto L28
	} else {
		goto L44
	}
L34:
	;
	v103 = v100
	goto L36
L35:
	;
	v103 = v101
	goto L36
L36:
	;
	v104 = F_memcmp(m, v87+v96, v88+v96, v103)
	mBase = m.M
	if v104 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v100 == v101 {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v104 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v108 = int32(10)
	v164 = (v91*v108 + v108) * (v100 - v101)
	goto L27
L41:
	;
	v120 = int32(-10)
	goto L43
L42:
	;
	v120 = int32(10)
	goto L43
L43:
	;
	v164 = (v91 + int32(1)) * v120
	goto L27
L44:
	;
	v126 = int32(9)
	v128 = int32(131064)
	v136 = int32(1)
	if v136 < v95 {
		v87 = v87 + (v100+v126)&v128
		v88 = v88 + (v101+v126)&v128
		v91 = v123
		v95 = v95 - v136
		goto L31
	} else {
		goto L45
	}
L45:
	;
	goto L32
L46:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v171&int32(3) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v174 = int32(0)
	goto L49
L48:
	;
	v174 = v51
	goto L49
L49:
	;
	v175 = v52 + v174
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175+int32(8))+4)))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+4)))
	if v186 == int32(0) {
		v250 = v186
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v2562 = v168
	v2565 = base.B2i32(int32(0) <= v272)
	goto L13
L51:
	;
	v272 = (v250 + int32(1)) * (v186 - v185) * int32(10)
	goto L50
L52:
	;
	if v185 == int32(0) {
		v250 = v186
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v195 = v168 + int32(8)
	v196 = v175 + int32(16)
	v199 = v186
	v203 = v185
	goto L54
L54:
	;
	v204 = int32(2)
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195))))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196))))
	if base.Ui32(v208) < base.Ui32(v209) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v250 = v231
	goto L51
L56:
	;
	v231 = v199 - int32(1)
	if v199 < int32(2) {
		v250 = v231
		goto L51
	} else {
		goto L67
	}
L57:
	;
	v211 = v208
	goto L59
L58:
	;
	v211 = v209
	goto L59
L59:
	;
	v212 = F_memcmp(m, v195+v204, v196+v204, v211)
	mBase = m.M
	if v212 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v208 == v209 {
		goto L56
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v212 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v216 = int32(10)
	v272 = (v199*v216 + v216) * (v208 - v209)
	goto L50
L64:
	;
	v228 = int32(-10)
	goto L66
L65:
	;
	v228 = int32(10)
	goto L66
L66:
	;
	v272 = (v199 + int32(1)) * v228
	goto L50
L67:
	;
	v234 = int32(9)
	v236 = int32(131064)
	v244 = int32(1)
	if v244 < v203 {
		v195 = v195 + (v208+v234)&v236
		v196 = v196 + (v209+v234)&v236
		v199 = v231
		v203 = v203 - v244
		goto L54
	} else {
		goto L68
	}
L68:
	;
	goto L55
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+16)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v279)+12)))
	if v281&int32(1) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+4)))
	if v294 == int32(0) {
		v358 = v294
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v384 = v52 + int32(8)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v386&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L73:
	;
	v2562 = v276
	v2565 = base.B2i32(v380 == int32(0))
	goto L13
L74:
	;
	v380 = (v358 + int32(1)) * (v294 - v293) * int32(10)
	goto L73
L75:
	;
	if v293 == int32(0) {
		v358 = v294
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v303 = v276 + int32(8)
	v304 = v52 + int32(16)
	v307 = v294
	v311 = v293
	goto L77
L77:
	;
	v312 = int32(2)
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303))))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304))))
	if base.Ui32(v316) < base.Ui32(v317) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v358 = v339
	goto L74
L79:
	;
	v339 = v307 - int32(1)
	if v307 < int32(2) {
		v358 = v339
		goto L74
	} else {
		goto L90
	}
L80:
	;
	v319 = v316
	goto L82
L81:
	;
	v319 = v317
	goto L82
L82:
	;
	v320 = F_memcmp(m, v303+v312, v304+v312, v319)
	mBase = m.M
	if v320 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v316 == v317 {
		goto L79
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v320 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v324 = int32(10)
	v380 = (v307*v324 + v324) * (v316 - v317)
	goto L73
L87:
	;
	v336 = int32(-10)
	goto L89
L88:
	;
	v336 = int32(10)
	goto L89
L89:
	;
	v380 = (v307 + int32(1)) * v336
	goto L73
L90:
	;
	v342 = int32(9)
	v344 = int32(131064)
	v352 = int32(1)
	if v352 < v311 {
		v303 = v303 + (v316+v342)&v344
		v304 = v304 + (v317+v342)&v344
		v307 = v339
		v311 = v311 - v352
		goto L77
	} else {
		goto L91
	}
L91:
	;
	goto L78
L92:
	;
	v389 = int32(0)
	goto L94
L93:
	;
	v389 = v51
	goto L94
L94:
	;
	v390 = v384 + v389
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+4)))
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+4)))
	if v399 == int32(0) {
		v463 = v399
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v485 < int32(0) {
		v2562 = v276
		v2565 = v2
		goto L13
	} else {
		goto L114
	}
L96:
	;
	v485 = (v463 + int32(1)) * (v399 - v398) * int32(10)
	goto L95
L97:
	;
	if v398 == int32(0) {
		v463 = v399
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v404 = int32(8)
	v408 = v276 + v404
	v409 = v390 + v404
	v412 = v399
	v416 = v398
	goto L99
L99:
	;
	v417 = int32(2)
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408))))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v409))))
	if base.Ui32(v421) < base.Ui32(v422) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v463 = v444
	goto L96
L101:
	;
	v444 = v412 - int32(1)
	if v412 < int32(2) {
		v463 = v444
		goto L96
	} else {
		goto L112
	}
L102:
	;
	v424 = v421
	goto L104
L103:
	;
	v424 = v422
	goto L104
L104:
	;
	v425 = F_memcmp(m, v408+v417, v409+v417, v424)
	mBase = m.M
	if v425 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v421 == v422 {
		goto L101
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v425 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v429 = int32(10)
	v485 = (v412*v429 + v429) * (v421 - v422)
	goto L95
L109:
	;
	v441 = int32(-10)
	goto L111
L110:
	;
	v441 = int32(10)
	goto L111
L111:
	;
	v485 = (v412 + int32(1)) * v441
	goto L95
L112:
	;
	v447 = int32(9)
	v449 = int32(131064)
	v457 = int32(1)
	if v457 < v416 {
		v408 = v408 + (v421+v447)&v449
		v409 = v409 + (v422+v447)&v449
		v412 = v444
		v416 = v416 - v457
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v488&int32(1) != 0 {
		v502 = v384
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v502)+4)))
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+4)))
	if v511 == int32(0) {
		v575 = v511
		goto L122
	} else {
		goto L123
	}
L116:
	;
	if v488&int32(2) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v494 = int32(0)
	goto L119
L118:
	;
	v494 = v51
	goto L119
L119:
	;
	v495 = v384 + v494
	if v488&int32(4) != 0 {
		v502 = v495
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v502 = v495 + int32(base.Ui32(v498)>>(uint(int32(2))%32))
	goto L115
L121:
	;
	v2562 = v276
	v2565 = base.B2i32(v597 <= int32(0))
	goto L13
L122:
	;
	v597 = (v575 + int32(1)) * (v511 - v510) * int32(10)
	goto L121
L123:
	;
	if v510 == int32(0) {
		v575 = v511
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v516 = int32(8)
	v520 = v276 + v516
	v521 = v502 + v516
	v524 = v511
	v528 = v510
	goto L125
L125:
	;
	v529 = int32(2)
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v520))))
	v534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521))))
	if base.Ui32(v533) < base.Ui32(v534) {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v575 = v556
	goto L122
L127:
	;
	v556 = v524 - int32(1)
	if v524 < int32(2) {
		v575 = v556
		goto L122
	} else {
		goto L138
	}
L128:
	;
	v536 = v533
	goto L130
L129:
	;
	v536 = v534
	goto L130
L130:
	;
	v537 = F_memcmp(m, v520+v529, v521+v529, v536)
	mBase = m.M
	if v537 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v533 == v534 {
		goto L127
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if v537 < int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v541 = int32(10)
	v597 = (v524*v541 + v541) * (v533 - v534)
	goto L121
L135:
	;
	v553 = int32(-10)
	goto L137
L136:
	;
	v553 = int32(10)
	goto L137
L137:
	;
	v597 = (v524 + int32(1)) * v553
	goto L121
L138:
	;
	v559 = int32(9)
	v561 = int32(131064)
	v569 = int32(1)
	if v569 < v528 {
		v520 = v520 + (v533+v559)&v561
		v521 = v521 + (v534+v559)&v561
		v524 = v556
		v528 = v528 - v569
		goto L125
	} else {
		goto L139
	}
L139:
	;
	goto L126
L140:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v605&int32(1) != 0 {
		v622 = v52 + int32(8)
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v622)+4)))
	v631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601)+4)))
	if v631 == int32(0) {
		v695 = v631
		goto L148
	} else {
		goto L149
	}
L142:
	;
	if v605&int32(2) != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v611 = int32(0)
	goto L145
L144:
	;
	v611 = v51
	goto L145
L145:
	;
	v614 = v52 + v611 + int32(8)
	if v605&int32(4) != 0 {
		v622 = v614
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	v622 = v614 + int32(base.Ui32(v617)>>(uint(int32(2))%32))
	goto L141
L147:
	;
	v2562 = v601
	v2565 = base.B2i32(v717 <= int32(0))
	goto L13
L148:
	;
	v717 = (v695 + int32(1)) * (v631 - v630) * int32(10)
	goto L147
L149:
	;
	if v630 == int32(0) {
		v695 = v631
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v636 = int32(8)
	v640 = v601 + v636
	v641 = v622 + v636
	v644 = v631
	v648 = v630
	goto L151
L151:
	;
	v649 = int32(2)
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v640))))
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v641))))
	if base.Ui32(v653) < base.Ui32(v654) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v695 = v676
	goto L148
L153:
	;
	v676 = v644 - int32(1)
	if v644 < int32(2) {
		v695 = v676
		goto L148
	} else {
		goto L164
	}
L154:
	;
	v656 = v653
	goto L156
L155:
	;
	v656 = v654
	goto L156
L156:
	;
	v657 = F_memcmp(m, v640+v649, v641+v649, v656)
	mBase = m.M
	if v657 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	if v653 == v654 {
		goto L153
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if v657 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v661 = int32(10)
	v717 = (v644*v661 + v661) * (v653 - v654)
	goto L147
L161:
	;
	v673 = int32(-10)
	goto L163
L162:
	;
	v673 = int32(10)
	goto L163
L163:
	;
	v717 = (v644 + int32(1)) * v673
	goto L147
L164:
	;
	v679 = int32(9)
	v681 = int32(131064)
	v689 = int32(1)
	if v689 < v648 {
		v640 = v640 + (v653+v679)&v681
		v641 = v641 + (v654+v679)&v681
		v644 = v676
		v648 = v648 - v689
		goto L151
	} else {
		goto L165
	}
L165:
	;
	goto L152
L166:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v724 = int32(1)
	v725 = v723 & v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v726)+16)))
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726+v727)+12)))
	if v729&v724 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if v725 != 0 {
		v748 = v52 + int32(8)
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	if v725 != 0 {
		v862 = v52 + int32(8)
		goto L195
	} else {
		goto L196
	}
L170:
	;
	v756 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v748)+4)))
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v721)+4)))
	if v757 == int32(0) {
		v821 = v757
		goto L177
	} else {
		goto L178
	}
L171:
	;
	if v723&int32(2) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v737 = int32(0)
	goto L174
L173:
	;
	v737 = v51
	goto L174
L174:
	;
	v740 = v52 + v737 + int32(8)
	if v723&int32(4) != 0 {
		v748 = v740
		goto L170
	} else {
		goto L175
	}
L175:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v748 = v740 + int32(base.Ui32(v743)>>(uint(int32(2))%32))
	goto L170
L176:
	;
	v2562 = v721
	v2565 = int32(base.Ui32(v843) >> (uint(int32(31)) % 32))
	goto L13
L177:
	;
	v843 = (v821 + int32(1)) * (v757 - v756) * int32(10)
	goto L176
L178:
	;
	if v756 == int32(0) {
		v821 = v757
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v762 = int32(8)
	v766 = v721 + v762
	v767 = v748 + v762
	v770 = v757
	v774 = v756
	goto L180
L180:
	;
	v775 = int32(2)
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v766))))
	v780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v767))))
	if base.Ui32(v779) < base.Ui32(v780) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v821 = v802
	goto L177
L182:
	;
	v802 = v770 - int32(1)
	if v770 < int32(2) {
		v821 = v802
		goto L177
	} else {
		goto L193
	}
L183:
	;
	v782 = v779
	goto L185
L184:
	;
	v782 = v780
	goto L185
L185:
	;
	v783 = F_memcmp(m, v766+v775, v767+v775, v782)
	mBase = m.M
	if v783 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	if v779 == v780 {
		goto L182
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if v783 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v787 = int32(10)
	v843 = (v770*v787 + v787) * (v779 - v780)
	goto L176
L190:
	;
	v799 = int32(-10)
	goto L192
L191:
	;
	v799 = int32(10)
	goto L192
L192:
	;
	v843 = (v770 + int32(1)) * v799
	goto L176
L193:
	;
	v805 = int32(9)
	v807 = int32(131064)
	v815 = int32(1)
	if v815 < v774 {
		v766 = v766 + (v779+v805)&v807
		v767 = v767 + (v780+v805)&v807
		v770 = v802
		v774 = v774 - v815
		goto L180
	} else {
		goto L194
	}
L194:
	;
	goto L181
L195:
	;
	v870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v862)+4)))
	v871 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v721)+4)))
	if v871 == int32(0) {
		v935 = v871
		goto L202
	} else {
		goto L203
	}
L196:
	;
	if v723&int32(2) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v851 = int32(0)
	goto L199
L198:
	;
	v851 = v51
	goto L199
L199:
	;
	v854 = v52 + v851 + int32(8)
	if v723&int32(4) != 0 {
		v862 = v854
		goto L195
	} else {
		goto L200
	}
L200:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v854)))
	v862 = v854 + int32(base.Ui32(v857)>>(uint(int32(2))%32))
	goto L195
L201:
	;
	v2562 = v721
	v2565 = base.B2i32(v957 <= int32(0))
	goto L13
L202:
	;
	v957 = (v935 + int32(1)) * (v871 - v870) * int32(10)
	goto L201
L203:
	;
	if v870 == int32(0) {
		v935 = v871
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v876 = int32(8)
	v880 = v721 + v876
	v881 = v862 + v876
	v884 = v871
	v888 = v870
	goto L205
L205:
	;
	v889 = int32(2)
	v893 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v880))))
	v894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v881))))
	if base.Ui32(v893) < base.Ui32(v894) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v935 = v916
	goto L202
L207:
	;
	v916 = v884 - int32(1)
	if v884 < int32(2) {
		v935 = v916
		goto L202
	} else {
		goto L218
	}
L208:
	;
	v896 = v893
	goto L210
L209:
	;
	v896 = v894
	goto L210
L210:
	;
	v897 = F_memcmp(m, v880+v889, v881+v889, v896)
	mBase = m.M
	if v897 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v893 == v894 {
		goto L207
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if v897 < int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v901 = int32(10)
	v957 = (v884*v901 + v901) * (v893 - v894)
	goto L201
L215:
	;
	v913 = int32(-10)
	goto L217
L216:
	;
	v913 = int32(10)
	goto L217
L217:
	;
	v957 = (v884 + int32(1)) * v913
	goto L201
L218:
	;
	v919 = int32(9)
	v921 = int32(131064)
	v929 = int32(1)
	if v929 < v888 {
		v880 = v880 + (v893+v919)&v921
		v881 = v881 + (v894+v919)&v921
		v884 = v916
		v888 = v888 - v929
		goto L205
	} else {
		goto L219
	}
L219:
	;
	goto L206
L220:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v963)+16)))
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963+v964)+12)))
	if v966&int32(1) != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v974 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	v975 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)))
	if base.Ui32(v975) < base.Ui32(v974) {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	goto L223
L223:
	;
	v1022 = v52 + int32(8)
	v1023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)))
	v1025 = v1023
	goto L240
L224:
	;
	v2562 = v961
	v2565 = v1020
	goto L13
L225:
	;
	v1020 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	if v974 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1020 = int32(1)
	goto L224
L229:
	;
	goto L230
L230:
	;
	v985 = v961 + int32(8)
	v986 = v52 + int32(16)
	v987 = v974
	goto L231
L231:
	;
	v990 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985))))
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v986))))
	if v990 != v991 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1020 = int32(1)
	goto L224
L233:
	;
	v1020 = int32(0)
	goto L224
L234:
	;
	goto L235
L235:
	;
	v994 = int32(2)
	v998 = F_memcmp(m, v985+v994, v986+v994, v990)
	mBase = m.M
	if v998 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1020 = int32(0)
	goto L224
L237:
	;
	goto L238
L238:
	;
	v1000 = int32(9)
	v1002 = int32(131064)
	v1010 = int32(1)
	if v1010 < v987 {
		v985 = v985 + (v990+v1000)&v1002
		v986 = v986 + (v991+v1000)&v1002
		v987 = v987 - v1010
		goto L231
	} else {
		goto L239
	}
L239:
	;
	goto L232
L240:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)) = uint16(v1025)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1042&int32(3) != 0 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)) = uint16(v1023)
	v2562 = v961
	v2565 = v1268
	goto L13
L242:
	;
	goto L241
L243:
	;
	v1045 = int32(0)
	goto L245
L244:
	;
	v1045 = v51
	goto L245
L245:
	;
	v1046 = v1022 + v1045
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+4)))
	v1055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)))
	if v1055 == int32(0) {
		v1119 = v1055
		goto L247
	} else {
		goto L248
	}
L246:
	;
	if int32(0) <= v1141 {
		goto L265
	} else {
		goto L266
	}
L247:
	;
	v1141 = (v1119 + int32(1)) * (v1055 - v1054) * int32(10)
	goto L246
L248:
	;
	if v1054 == int32(0) {
		v1119 = v1055
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1060 = int32(8)
	v1064 = v961 + v1060
	v1065 = v1046 + v1060
	v1068 = v1055
	v1072 = v1054
	goto L250
L250:
	;
	v1073 = int32(2)
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1064))))
	v1078 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1065))))
	if base.Ui32(v1077) < base.Ui32(v1078) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	v1119 = v1100
	goto L247
L252:
	;
	v1100 = v1068 - int32(1)
	if v1068 < int32(2) {
		v1119 = v1100
		goto L247
	} else {
		goto L263
	}
L253:
	;
	v1080 = v1077
	goto L255
L254:
	;
	v1080 = v1078
	goto L255
L255:
	;
	v1081 = F_memcmp(m, v1064+v1073, v1065+v1073, v1080)
	mBase = m.M
	if v1081 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if v1077 == v1078 {
		goto L252
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v1081 < int32(0) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1085 = int32(10)
	v1141 = (v1068*v1085 + v1085) * (v1077 - v1078)
	goto L246
L260:
	;
	v1097 = int32(-10)
	goto L262
L261:
	;
	v1097 = int32(10)
	goto L262
L262:
	;
	v1141 = (v1068 + int32(1)) * v1097
	goto L246
L263:
	;
	v1103 = int32(9)
	v1105 = int32(131064)
	v1113 = int32(1)
	if v1113 < v1072 {
		v1064 = v1064 + (v1077+v1103)&v1105
		v1065 = v1065 + (v1078+v1103)&v1105
		v1068 = v1100
		v1072 = v1072 - v1113
		goto L250
	} else {
		goto L264
	}
L264:
	;
	goto L251
L265:
	;
	v1144 = int32(1)
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1145&v1144 != 0 {
		v1160 = v1022
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L267
L267:
	;
	v1261 = int32(0)
	if v1261 < v1025 {
		v1025 = v1025 - int32(1)
		goto L240
	} else {
		goto L294
	}
L268:
	;
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1160)+4)))
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)))
	if v1169 == int32(0) {
		v1233 = v1169
		goto L275
	} else {
		goto L276
	}
L269:
	;
	if v1145&int32(2) != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1151 = int32(0)
	goto L272
L271:
	;
	v1151 = v51
	goto L272
L272:
	;
	v1152 = v1022 + v1151
	if v1145&int32(4) != 0 {
		v1160 = v1152
		goto L268
	} else {
		goto L273
	}
L273:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1160 = v1152 + int32(base.Ui32(v1155)>>(uint(int32(2))%32))
	goto L268
L274:
	;
	if v1255 <= int32(0) {
		v1268 = v1144
		goto L242
	} else {
		goto L293
	}
L275:
	;
	v1255 = (v1233 + int32(1)) * (v1169 - v1168) * int32(10)
	goto L274
L276:
	;
	if v1168 == int32(0) {
		v1233 = v1169
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1174 = int32(8)
	v1178 = v961 + v1174
	v1179 = v1160 + v1174
	v1182 = v1169
	v1186 = v1168
	goto L278
L278:
	;
	v1187 = int32(2)
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1178))))
	v1192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1179))))
	if base.Ui32(v1191) < base.Ui32(v1192) {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v1233 = v1214
	goto L275
L280:
	;
	v1214 = v1182 - int32(1)
	if v1182 < int32(2) {
		v1233 = v1214
		goto L275
	} else {
		goto L291
	}
L281:
	;
	v1194 = v1191
	goto L283
L282:
	;
	v1194 = v1192
	goto L283
L283:
	;
	v1195 = F_memcmp(m, v1178+v1187, v1179+v1187, v1194)
	mBase = m.M
	if v1195 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	if v1191 == v1192 {
		goto L280
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	if v1195 < int32(0) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1199 = int32(10)
	v1255 = (v1182*v1199 + v1199) * (v1191 - v1192)
	goto L274
L288:
	;
	v1211 = int32(-10)
	goto L290
L289:
	;
	v1211 = int32(10)
	goto L290
L290:
	;
	v1255 = (v1182 + int32(1)) * v1211
	goto L274
L291:
	;
	v1217 = int32(9)
	v1219 = int32(131064)
	v1227 = int32(1)
	if v1227 < v1186 {
		v1178 = v1178 + (v1191+v1217)&v1219
		v1179 = v1179 + (v1192+v1217)&v1219
		v1182 = v1214
		v1186 = v1186 - v1227
		goto L278
	} else {
		goto L292
	}
L292:
	;
	goto L279
L293:
	;
	goto L267
L294:
	;
	v1268 = v1261
	goto L242
L295:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1274)+16)))
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274+v1275)+12)))
	if v1277&int32(1) != 0 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+4)))
	v1286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+int32(8))+4)))
	if base.Ui32(v1286) < base.Ui32(v1285) {
		goto L300
	} else {
		goto L301
	}
L297:
	;
	goto L298
L298:
	;
	v1333 = v52 + int32(8)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1335&int32(3) != 0 {
		goto L315
	} else {
		goto L316
	}
L299:
	;
	v2562 = v1272
	v2565 = v1331
	goto L13
L300:
	;
	v1331 = int32(0)
	goto L299
L301:
	;
	goto L302
L302:
	;
	if v1285 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1331 = int32(1)
	goto L299
L304:
	;
	goto L305
L305:
	;
	v1296 = v52 + int32(16)
	v1297 = v1272 + int32(8)
	v1298 = v1285
	goto L306
L306:
	;
	v1301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1296))))
	v1302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1297))))
	if v1301 != v1302 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1331 = int32(1)
	goto L299
L308:
	;
	v1331 = int32(0)
	goto L299
L309:
	;
	goto L310
L310:
	;
	v1305 = int32(2)
	v1309 = F_memcmp(m, v1296+v1305, v1297+v1305, v1301)
	mBase = m.M
	if v1309 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1331 = int32(0)
	goto L299
L312:
	;
	goto L313
L313:
	;
	v1311 = int32(9)
	v1313 = int32(131064)
	v1321 = int32(1)
	if v1321 < v1298 {
		v1296 = v1296 + (v1301+v1311)&v1313
		v1297 = v1297 + (v1302+v1311)&v1313
		v1298 = v1298 - v1321
		goto L306
	} else {
		goto L314
	}
L314:
	;
	goto L307
L315:
	;
	v1338 = int32(0)
	goto L317
L316:
	;
	v1338 = v51
	goto L317
L317:
	;
	v1339 = v1333 + v1338
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1339)))
	v1343 = F_palloc0(m, int32(base.Ui32(v1340)>>(uint(int32(2))%32)))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L10
	} else {
		goto L318
	}
L318:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1339)))
	v1347 = int32(base.Ui32(v1345) >> (uint(int32(2)) % 32))
	if v1347 != 0 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1350&int32(1) != 0 {
		v1364 = v1333
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v1348 = F__emscripten_memcpy_bulkmem(m, v1343, v1339, v1347)
	mBase = m.M
	v1349 = v1348
	goto L322
L321:
	;
	v1349 = v1343
	goto L322
L322:
	;
	goto L319
L323:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	v1368 = F_palloc0(m, int32(base.Ui32(v1365)>>(uint(int32(2))%32)))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L10
	} else {
		goto L329
	}
L324:
	;
	if v1350&int32(2) != 0 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1356 = int32(0)
	goto L327
L326:
	;
	v1356 = v51
	goto L327
L327:
	;
	v1357 = v1333 + v1356
	if v1350&int32(4) != 0 {
		v1364 = v1357
		goto L323
	} else {
		goto L328
	}
L328:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1357)))
	v1364 = v1357 + int32(base.Ui32(v1360)>>(uint(int32(2))%32))
	goto L323
L329:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	v1372 = int32(base.Ui32(v1370) >> (uint(int32(2)) % 32))
	if v1372 != 0 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+4)))
	v1376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1349)+4)))
	if base.Ui32(v1375) < base.Ui32(v1376) {
		goto L334
	} else {
		goto L335
	}
L331:
	;
	v1373 = F__emscripten_memcpy_bulkmem(m, v1368, v1364, v1372)
	mBase = m.M
	v1374 = v1373
	goto L333
L332:
	;
	v1374 = v1368
	goto L333
L333:
	;
	goto L330
L334:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1349)+4)) = uint16(v1375)
	goto L336
L335:
	;
	goto L336
L336:
	;
	v1386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1349)+4)))
	v1387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+4)))
	if v1387 == int32(0) {
		v1451 = v1387
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v1474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+4)))
	v1475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1374)+4)))
	if base.Ui32(v1474) < base.Ui32(v1475) {
		goto L356
	} else {
		goto L357
	}
L338:
	;
	v1473 = (v1451 + int32(1)) * (v1387 - v1386) * int32(10)
	goto L337
L339:
	;
	if v1386 == int32(0) {
		v1451 = v1387
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1392 = int32(8)
	v1396 = v1272 + v1392
	v1397 = v1349 + v1392
	v1400 = v1387
	v1404 = v1386
	goto L341
L341:
	;
	v1405 = int32(2)
	v1409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1396))))
	v1410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1397))))
	if base.Ui32(v1409) < base.Ui32(v1410) {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	v1451 = v1432
	goto L338
L343:
	;
	v1432 = v1400 - int32(1)
	if v1400 < int32(2) {
		v1451 = v1432
		goto L338
	} else {
		goto L354
	}
L344:
	;
	v1412 = v1409
	goto L346
L345:
	;
	v1412 = v1410
	goto L346
L346:
	;
	v1413 = F_memcmp(m, v1396+v1405, v1397+v1405, v1412)
	mBase = m.M
	if v1413 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	if v1409 == v1410 {
		goto L343
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	if v1413 < int32(0) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v1417 = int32(10)
	v1473 = (v1400*v1417 + v1417) * (v1409 - v1410)
	goto L337
L351:
	;
	v1429 = int32(-10)
	goto L353
L352:
	;
	v1429 = int32(10)
	goto L353
L353:
	;
	v1473 = (v1400 + int32(1)) * v1429
	goto L337
L354:
	;
	v1435 = int32(9)
	v1437 = int32(131064)
	v1445 = int32(1)
	if v1445 < v1404 {
		v1396 = v1396 + (v1409+v1435)&v1437
		v1397 = v1397 + (v1410+v1435)&v1437
		v1400 = v1432
		v1404 = v1404 - v1445
		goto L341
	} else {
		goto L355
	}
L355:
	;
	goto L342
L356:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1374)+4)) = uint16(v1474)
	goto L358
L357:
	;
	goto L358
L358:
	;
	if int32(0) <= v1473 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1374)+4)))
	v1488 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+4)))
	if v1488 == int32(0) {
		v1552 = v1488
		goto L363
	} else {
		goto L364
	}
L360:
	;
	v1577 = v2
	goto L361
L361:
	;
	F_pfree(m, v1349)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L10
	} else {
		goto L381
	}
L362:
	;
	v1577 = base.B2i32(v1574 <= int32(0))
	goto L361
L363:
	;
	v1574 = (v1552 + int32(1)) * (v1488 - v1487) * int32(10)
	goto L362
L364:
	;
	if v1487 == int32(0) {
		v1552 = v1488
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1493 = int32(8)
	v1497 = v1272 + v1493
	v1498 = v1374 + v1493
	v1501 = v1488
	v1505 = v1487
	goto L366
L366:
	;
	v1506 = int32(2)
	v1510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1497))))
	v1511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498))))
	if base.Ui32(v1510) < base.Ui32(v1511) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	v1552 = v1533
	goto L363
L368:
	;
	v1533 = v1501 - int32(1)
	if v1501 < int32(2) {
		v1552 = v1533
		goto L363
	} else {
		goto L379
	}
L369:
	;
	v1513 = v1510
	goto L371
L370:
	;
	v1513 = v1511
	goto L371
L371:
	;
	v1514 = F_memcmp(m, v1497+v1506, v1498+v1506, v1513)
	mBase = m.M
	if v1514 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	if v1510 == v1511 {
		goto L368
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	if v1514 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1518 = int32(10)
	v1574 = (v1501*v1518 + v1518) * (v1510 - v1511)
	goto L362
L376:
	;
	v1530 = int32(-10)
	goto L378
L377:
	;
	v1530 = int32(10)
	goto L378
L378:
	;
	v1574 = (v1501 + int32(1)) * v1530
	goto L362
L379:
	;
	v1536 = int32(9)
	v1538 = int32(131064)
	v1546 = int32(1)
	if v1546 < v1505 {
		v1497 = v1497 + (v1510+v1536)&v1538
		v1498 = v1498 + (v1511+v1536)&v1538
		v1501 = v1533
		v1505 = v1505 - v1546
		goto L366
	} else {
		goto L380
	}
L380:
	;
	goto L367
L381:
	;
	F_pfree(m, v1374)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L10
	} else {
		goto L382
	}
L382:
	;
	v2562 = v1272
	v2565 = v1577
	goto L13
L383:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1587)+16)))
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587+v1588)+12)))
	if v1590&int32(1) != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1595 = F_DirectFunctionCall2Coll(m, int32(5617), int32(0), v1583, v1585)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L10
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v1599&int32(2) != 0 {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	v2562 = v1585
	v2565 = base.B2i32(v1595 != int32(0))
	goto L13
L388:
	;
	v1725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1585)+6)))
	if v1725 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L389:
	;
	v1602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1585)+4)))
	if v1602 == int32(0) {
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v1610 = v1602
	v1618 = v1585 + int32(16)
	goto L391
L391:
	;
	v1625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1618)+4)))
	if v1625 == int32(0) {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	goto L388
L393:
	;
	v1688 = int32(1)
	v1690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1618))))
	if v1688 < v1610 {
		v1610 = v1610 - v1688
		v1618 = v1618 + (v1690+int32(7))&int32(131064)
		goto L391
	} else {
		goto L400
	}
L394:
	;
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1618)+2)))
	if v1628&int32(21) != 0 {
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1634 = v1618 + int32(16)
	v1636 = v1625
	goto L396
L396:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1634)))
	v1650 = base.I32_rem_u_s(v1649, v51<<(uint(int32(3))%32))
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583+int32(base.Ui32(v1650)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v1654)>>(uint(v1650&int32(7))%32))&int32(1) != 0 {
		goto L393
	} else {
		goto L398
	}
L397:
	;
	v2562 = v1585
	v2565 = v2
	goto L13
L398:
	;
	v1660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1634)+4)))
	v1668 = int32(1)
	if v1668 < v1636 {
		v1634 = v1634 + (v1660+int32(7))&int32(131064) + int32(8)
		v1636 = v1636 - v1668
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	goto L392
L401:
	;
	v2562 = v1585
	v2565 = v1945
	goto L13
L402:
	;
	v1945 = int32(1)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v1730 = v52 + int32(8)
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v1732&int32(2) != 0 {
		goto L407
	} else {
		goto L408
	}
L405:
	;
	v1822 = int32(0)
	if v1817 <= v1822 {
		goto L428
	} else {
		goto L429
	}
L406:
	;
	if base.Ui32(v1740) < base.Ui32(v1725) {
		goto L425
	} else {
		goto L426
	}
L407:
	;
	v1735 = int32(0)
	goto L409
L408:
	;
	v1735 = v51
	goto L409
L409:
	;
	v1736 = v1730 + v1735
	v1738 = v1732 & int32(1)
	if v1738 != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1739 = v1730
	goto L412
L411:
	;
	v1739 = v1736
	goto L412
L412:
	;
	v1740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1739)+4)))
	if v1740 == int32(0) {
		goto L406
	} else {
		goto L413
	}
L413:
	;
	v1747 = v1585 + int32(16)
	v1749 = v1739 + int32(8)
	v1753 = v1725
	v1754 = v1740
	goto L414
L414:
	;
	v1765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1749))))
	v1766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1747)+20)))
	if base.Ui32(v1765) < base.Ui32(v1766) {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	goto L406
L416:
	;
	v1768 = v1765
	goto L418
L417:
	;
	v1768 = v1766
	goto L418
L418:
	;
	v1769 = F_memcmp(m, v1749+int32(2), v1747+int32(23), v1768)
	mBase = m.M
	if v1769 != 0 {
		v1817 = v1769
		goto L405
	} else {
		goto L419
	}
L419:
	;
	if v1765 != v1766 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1817 = v1765 - v1766
	goto L405
L421:
	;
	goto L422
L422:
	;
	if v1754 < int32(2) {
		goto L406
	} else {
		goto L423
	}
L423:
	;
	v1774 = int32(1)
	v1778 = int32(131064)
	v1781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1747))))
	if v1774 < v1753 {
		v1747 = v1747 + (v1781+int32(7))&v1778
		v1749 = v1749 + (v1765+int32(9))&v1778
		v1753 = v1753 - v1774
		v1754 = v1754 - v1774
		goto L414
	} else {
		goto L424
	}
L424:
	;
	goto L415
L425:
	;
	v1806 = v1740
	goto L427
L426:
	;
	v1806 = v1725
	goto L427
L427:
	;
	v1817 = v1806 - v1725
	goto L405
L428:
	;
	if v1738 != 0 {
		v1831 = v1730
		goto L431
	} else {
		goto L432
	}
L429:
	;
	v1930 = v1822
	goto L430
L430:
	;
	v1945 = v1930
	goto L401
L431:
	;
	v1832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1831)+4)))
	if v1832 == int32(0) {
		goto L437
	} else {
		goto L438
	}
L432:
	;
	if v1732&int32(4) != 0 {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1831 = v1736
	goto L431
L434:
	;
	goto L435
L435:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1736)))
	v1831 = v1736 + int32(base.Ui32(v1827)>>(uint(int32(2))%32))
	goto L431
L436:
	;
	v1930 = base.B2i32(int32(0) <= v1909)
	goto L430
L437:
	;
	if base.Ui32(v1832) < base.Ui32(v1725) {
		goto L450
	} else {
		goto L451
	}
L438:
	;
	v1839 = v1585 + int32(16)
	v1841 = v1831 + int32(8)
	v1845 = v1725
	v1846 = v1832
	goto L439
L439:
	;
	v1857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1841))))
	v1858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1839)+20)))
	if base.Ui32(v1857) < base.Ui32(v1858) {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	goto L437
L441:
	;
	v1860 = v1857
	goto L443
L442:
	;
	v1860 = v1858
	goto L443
L443:
	;
	v1861 = F_memcmp(m, v1841+int32(2), v1839+int32(23), v1860)
	mBase = m.M
	if v1861 != 0 {
		v1909 = v1861
		goto L436
	} else {
		goto L444
	}
L444:
	;
	if v1857 != v1858 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1909 = v1857 - v1858
	goto L436
L446:
	;
	goto L447
L447:
	;
	if v1846 < int32(2) {
		goto L437
	} else {
		goto L448
	}
L448:
	;
	v1866 = int32(1)
	v1870 = int32(131064)
	v1873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1839))))
	if v1866 < v1845 {
		v1839 = v1839 + (v1873+int32(7))&v1870
		v1841 = v1841 + (v1857+int32(9))&v1870
		v1845 = v1845 - v1866
		v1846 = v1846 - v1866
		goto L439
	} else {
		goto L449
	}
L449:
	;
	goto L440
L450:
	;
	v1898 = v1832
	goto L452
L451:
	;
	v1898 = v1725
	goto L452
L452:
	;
	v1909 = v1898 - v1725
	goto L436
L453:
	;
	v1949 = int32(1)
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1951 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1950)+16)))
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1950+v1951)+12)))
	if v1953&v1949 != 0 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1960 = F_DirectFunctionCall2Coll(m, int32(5654), int32(0), v52+int32(8), v1947)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L10
	} else {
		goto L457
	}
L455:
	;
	goto L456
L456:
	;
	v1964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v1964&int32(2) != 0 {
		v2562 = v1947
		v2565 = v1949
		goto L13
	} else {
		goto L458
	}
L457:
	;
	v2562 = v1947
	v2565 = base.B2i32(v1960 != int32(0))
	goto L13
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v51
	v1968 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v52 + v1968
	v1977 = F_ltree_execute(m, v1947+v1968, v19+v1968, int32(0), int32(7000))
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L10
	} else {
		goto L459
	}
L459:
	;
	v2562 = v1947
	v2565 = v1977
	goto L13
L460:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1982)+16)))
	v1985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1982+v1983)+12)))
	if v1985&int32(1) != 0 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1992 = F_DirectFunctionCall2Coll(m, int32(5615), int32(0), v52+int32(8), v1980)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L10
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+8))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	v2000 = F_ArrayGetNItems(m, v1997, v1980+int32(16))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L10
	} else {
		goto L465
	}
L464:
	;
	v2562 = v1980
	v2565 = base.B2i32(v1992 != int32(0))
	goto L13
L465:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1980)+4))
	if v2002 < int32(2) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2005 = F_array_contains_nulls(m, v1980)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L10
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L10
	} else {
		goto L551
	}
L469:
	;
	if v2005 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	if v2000 <= int32(0) {
		v2562 = v1980
		v2565 = v2
		goto L13
	} else {
		goto L473
	}
L471:
	;
	goto L472
L472:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L10
	} else {
		goto L547
	}
L473:
	;
	if v1996 != 0 {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2017 = v1996
	goto L476
L475:
	;
	v2017 = (v1997<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L476
L476:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v2037 = v1980 + v2017
	v2039 = v2000
	goto L477
L477:
	;
	if v2023&int32(2) != 0 {
		goto L480
	} else {
		goto L481
	}
L478:
	;
	v2562 = v1980
	v2565 = v2
	goto L13
L479:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2037)))
	v2411 = int32(1)
	if v2411 < v2039 {
		v2037 = v2037 + (int32(base.Ui32(v2403)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v2039 = v2039 - v2411
		goto L477
	} else {
		goto L546
	}
L480:
	;
	v2163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2037)+6)))
	if v2163 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L481:
	;
	v2042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2037)+4)))
	if v2042 == int32(0) {
		goto L480
	} else {
		goto L482
	}
L482:
	;
	v2048 = v2042
	v2055 = v2037 + int32(16)
	goto L483
L483:
	;
	v2063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2055)+4)))
	if v2063 == int32(0) {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	goto L480
L485:
	;
	v2126 = int32(1)
	v2128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2055))))
	if v2126 < v2048 {
		v2048 = v2048 - v2126
		v2055 = v2055 + (v2128+int32(7))&int32(131064)
		goto L483
	} else {
		goto L492
	}
L486:
	;
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055)+2)))
	if v2066&int32(21) != 0 {
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2072 = v2055 + int32(16)
	v2074 = v2063
	goto L488
L488:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2072)))
	v2088 = base.I32_rem_u_s(v2087, v51<<(uint(int32(3))%32))
	v2092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(8)+int32(base.Ui32(v2088)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v2092)>>(uint(v2088&int32(7))%32))&int32(1) != 0 {
		goto L485
	} else {
		goto L490
	}
L489:
	;
	goto L479
L490:
	;
	v2098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2072)+4)))
	v2106 = int32(1)
	if v2106 < v2074 {
		v2072 = v2072 + (v2098+int32(7))&int32(131064) + int32(8)
		v2074 = v2074 - v2106
		goto L488
	} else {
		goto L491
	}
L491:
	;
	goto L489
L492:
	;
	goto L484
L493:
	;
	if v2383 == int32(0) {
		goto L479
	} else {
		goto L545
	}
L494:
	;
	v2383 = int32(1)
	goto L493
L495:
	;
	goto L496
L496:
	;
	v2168 = v52 + int32(8)
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v2170&int32(2) != 0 {
		goto L499
	} else {
		goto L500
	}
L497:
	;
	v2260 = int32(0)
	if v2255 <= v2260 {
		goto L520
	} else {
		goto L521
	}
L498:
	;
	if base.Ui32(v2178) < base.Ui32(v2163) {
		goto L517
	} else {
		goto L518
	}
L499:
	;
	v2173 = int32(0)
	goto L501
L500:
	;
	v2173 = v51
	goto L501
L501:
	;
	v2174 = v2168 + v2173
	v2176 = v2170 & int32(1)
	if v2176 != 0 {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2177 = v2168
	goto L504
L503:
	;
	v2177 = v2174
	goto L504
L504:
	;
	v2178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2177)+4)))
	if v2178 == int32(0) {
		goto L498
	} else {
		goto L505
	}
L505:
	;
	v2185 = v2037 + int32(16)
	v2187 = v2177 + int32(8)
	v2191 = v2163
	v2192 = v2178
	goto L506
L506:
	;
	v2203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2187))))
	v2204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2185)+20)))
	if base.Ui32(v2203) < base.Ui32(v2204) {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	goto L498
L508:
	;
	v2206 = v2203
	goto L510
L509:
	;
	v2206 = v2204
	goto L510
L510:
	;
	v2207 = F_memcmp(m, v2187+int32(2), v2185+int32(23), v2206)
	mBase = m.M
	if v2207 != 0 {
		v2255 = v2207
		goto L497
	} else {
		goto L511
	}
L511:
	;
	if v2203 != v2204 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v2255 = v2203 - v2204
	goto L497
L513:
	;
	goto L514
L514:
	;
	if v2192 < int32(2) {
		goto L498
	} else {
		goto L515
	}
L515:
	;
	v2212 = int32(1)
	v2216 = int32(131064)
	v2219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2185))))
	if v2212 < v2191 {
		v2185 = v2185 + (v2219+int32(7))&v2216
		v2187 = v2187 + (v2203+int32(9))&v2216
		v2191 = v2191 - v2212
		v2192 = v2192 - v2212
		goto L506
	} else {
		goto L516
	}
L516:
	;
	goto L507
L517:
	;
	v2244 = v2178
	goto L519
L518:
	;
	v2244 = v2163
	goto L519
L519:
	;
	v2255 = v2244 - v2163
	goto L497
L520:
	;
	if v2176 != 0 {
		v2269 = v2168
		goto L523
	} else {
		goto L524
	}
L521:
	;
	v2368 = v2260
	goto L522
L522:
	;
	v2383 = v2368
	goto L493
L523:
	;
	v2270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2269)+4)))
	if v2270 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L524:
	;
	if v2170&int32(4) != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2269 = v2174
	goto L523
L526:
	;
	goto L527
L527:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2174)))
	v2269 = v2174 + int32(base.Ui32(v2265)>>(uint(int32(2))%32))
	goto L523
L528:
	;
	v2368 = base.B2i32(int32(0) <= v2347)
	goto L522
L529:
	;
	if base.Ui32(v2270) < base.Ui32(v2163) {
		goto L542
	} else {
		goto L543
	}
L530:
	;
	v2277 = v2037 + int32(16)
	v2279 = v2269 + int32(8)
	v2283 = v2163
	v2284 = v2270
	goto L531
L531:
	;
	v2295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2279))))
	v2296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2277)+20)))
	if base.Ui32(v2295) < base.Ui32(v2296) {
		goto L533
	} else {
		goto L534
	}
L532:
	;
	goto L529
L533:
	;
	v2298 = v2295
	goto L535
L534:
	;
	v2298 = v2296
	goto L535
L535:
	;
	v2299 = F_memcmp(m, v2279+int32(2), v2277+int32(23), v2298)
	mBase = m.M
	if v2299 != 0 {
		v2347 = v2299
		goto L528
	} else {
		goto L536
	}
L536:
	;
	if v2295 != v2296 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2347 = v2295 - v2296
	goto L528
L538:
	;
	goto L539
L539:
	;
	if v2284 < int32(2) {
		goto L529
	} else {
		goto L540
	}
L540:
	;
	v2304 = int32(1)
	v2308 = int32(131064)
	v2311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2277))))
	if v2304 < v2283 {
		v2277 = v2277 + (v2311+int32(7))&v2308
		v2279 = v2279 + (v2295+int32(9))&v2308
		v2283 = v2283 - v2304
		v2284 = v2284 - v2304
		goto L531
	} else {
		goto L541
	}
L541:
	;
	goto L532
L542:
	;
	v2336 = v2270
	goto L544
L543:
	;
	v2336 = v2163
	goto L544
L544:
	;
	v2347 = v2336 - v2163
	goto L528
L545:
	;
	v2562 = v1980
	v2565 = int32(1)
	goto L13
L546:
	;
	goto L478
L547:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L10
	} else {
		goto L548
	}
L548:
	;
	F_errmsg(m, int32(153131), int32(0))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L10
	} else {
		goto L549
	}
L549:
	;
	F_errfinish(m, int32(494159), int32(604), int32(147830))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L10
	} else {
		goto L550
	}
L550:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L551:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L10
	} else {
		goto L552
	}
L552:
	;
	F_errmsg(m, int32(313999), int32(0))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L10
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(494159), int32(600), int32(147830))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L10
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	v2459 = int32(0)
	goto L557
L556:
	;
	v2459 = v51
	goto L557
L557:
	;
	v2460 = v52 + v2459
	v2470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2460+int32(8))+4)))
	v2471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	if v2471 == int32(0) {
		v2535 = v2471
		goto L559
	} else {
		goto L560
	}
L558:
	;
	v2562 = v58
	v2565 = base.B2i32(int32(0) <= v2557)
	goto L13
L559:
	;
	v2557 = (v2535 + int32(1)) * (v2471 - v2470) * int32(10)
	goto L558
L560:
	;
	if v2470 == int32(0) {
		v2535 = v2471
		goto L559
	} else {
		goto L561
	}
L561:
	;
	v2480 = v58 + int32(8)
	v2481 = v2460 + int32(16)
	v2484 = v2471
	v2488 = v2470
	goto L562
L562:
	;
	v2489 = int32(2)
	v2493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2480))))
	v2494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2481))))
	if base.Ui32(v2493) < base.Ui32(v2494) {
		goto L565
	} else {
		goto L566
	}
L563:
	;
	v2535 = v2516
	goto L559
L564:
	;
	v2516 = v2484 - int32(1)
	if v2484 < int32(2) {
		v2535 = v2516
		goto L559
	} else {
		goto L575
	}
L565:
	;
	v2496 = v2493
	goto L567
L566:
	;
	v2496 = v2494
	goto L567
L567:
	;
	v2497 = F_memcmp(m, v2480+v2489, v2481+v2489, v2496)
	mBase = m.M
	if v2497 == int32(0) {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	if v2493 == v2494 {
		goto L564
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	if v2497 < int32(0) {
		goto L572
	} else {
		goto L573
	}
L571:
	;
	v2501 = int32(10)
	v2557 = (v2484*v2501 + v2501) * (v2493 - v2494)
	goto L558
L572:
	;
	v2513 = int32(-10)
	goto L574
L573:
	;
	v2513 = int32(10)
	goto L574
L574:
	;
	v2557 = (v2484 + int32(1)) * v2513
	goto L558
L575:
	;
	v2519 = int32(9)
	v2521 = int32(131064)
	v2529 = int32(1)
	if v2529 < v2488 {
		v2480 = v2480 + (v2493+v2519)&v2521
		v2481 = v2481 + (v2494+v2519)&v2521
		v2484 = v2516
		v2488 = v2488 - v2529
		goto L562
	} else {
		goto L576
	}
L576:
	;
	goto L563
L577:
	;
	F_pfree(m, v2562)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L10
	} else {
		goto L580
	}
L578:
	;
	goto L579
L579:
	;
	m.G0 = v19 + int32(16)
	return v2565
L580:
	;
	goto L579
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
	F_errmsg_internal(m, int32(483927), v19)
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L10
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(494159), int32(715), int32(93269))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L10
	} else {
		goto L583
	}
L583:
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
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
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
	if v23 == int32(0) {
		v145 = v23
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v173 != v15 {
		goto L37
	} else {
		goto L38
	}
L5:
	;
	v171 = (v145 + int32(1)) * (v23 - v22) * int32(10)
	goto L4
L6:
	;
	if v22 == int32(0) {
		v145 = v23
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(8)
	v38 = v15 + v28
	v39 = v20 + v28
	v40 = v23
	v45 = v22
	goto L8
L8:
	;
	v46 = int32(2)
	v47 = v38 + v46
	v49 = v39 + v46
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32(v50) < base.Ui32(v51) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v145 = v124
	goto L5
L10:
	;
	v53 = v50
	goto L12
L11:
	;
	v53 = v51
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v115 != 0 {
		v171 = int32(1)
		goto L4
	} else {
		goto L31
	}
L14:
	;
	v115 = int32(0)
	goto L13
L15:
	;
	v89 = v84
	v90 = v85
	v91 = v86
	goto L25
L16:
	;
	if (v47|v49)&int32(3) != 0 {
		v84 = v47
		v85 = v49
		v86 = v53
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v77 = v47
	v78 = v49
	v79 = v53
	goto L18
L18:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v61 = v47
	v62 = v49
	v63 = v53
	goto L20
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v66 != v67 {
		v84 = v61
		v85 = v62
		v86 = v63
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v77 = v72
	v78 = v70
	v79 = v74
	goto L18
L22:
	;
	v69 = int32(4)
	v70 = v62 + v69
	v72 = v61 + v69
	v74 = v63 - v69
	if base.Ui32(int32(3)) < base.Ui32(v74) {
		v61 = v72
		v62 = v70
		v63 = v74
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v84 = v77
	v85 = v78
	v86 = v79
	goto L15
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 == v95 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v115 = v94 - v95
	goto L13
L27:
	;
	v97 = int32(1)
	v102 = v91 - v97
	if v102 != 0 {
		v89 = v89 + v97
		v90 = v90 + v97
		v91 = v102
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
	if v50 != v51 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = int32(10)
	v171 = (v40*v117 + v117) * (v50 - v51)
	goto L4
L33:
	;
	goto L34
L34:
	;
	v124 = v40 - int32(1)
	if v40 < int32(2) {
		v145 = v124
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v127 = int32(9)
	v129 = int32(131064)
	v137 = int32(1)
	if v137 < v45 {
		v38 = v38 + (v50+v127)&v129
		v39 = v39 + (v51+v127)&v129
		v40 = v124
		v45 = v45 - v137
		goto L8
	} else {
		goto L36
	}
L36:
	;
	goto L9
L37:
	;
	F_pfree(m, v15)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v177 != v20 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	F_pfree(m, v20)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	return base.B2i32(v171 == int32(0))
L44:
	;
	goto L43
}
func F_ltree_gist_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(76078)
			F_errmsg(m, int32(192969), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494159), int32(36), int32(67074))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
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
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
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
	if v22 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v179 != v14 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v178 = (v149 + int32(1)) * (v22 - v21) * int32(10)
	goto L4
L6:
	;
	if v21 == int32(0) {
		v149 = v22
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v27 = int32(8)
	v33 = v14 + v27
	v35 = v22
	v38 = v19 + v27
	v42 = v21
	goto L8
L8:
	;
	v43 = int32(2)
	v44 = v33 + v43
	v46 = v38 + v43
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if base.Ui32(v47) < base.Ui32(v48) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v149 = v131
	goto L5
L10:
	;
	v131 = v35 - int32(1)
	if v35 < int32(2) {
		v149 = v131
		goto L5
	} else {
		goto L39
	}
L11:
	;
	v50 = v47
	goto L13
L12:
	;
	v50 = v48
	goto L13
L13:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v50) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v112 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v112 = int32(0)
	goto L14
L16:
	;
	v86 = v81
	v87 = v82
	v88 = v83
	goto L26
L17:
	;
	if (v44|v46)&int32(3) != 0 {
		v81 = v44
		v82 = v46
		v83 = v50
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v74 = v44
	v75 = v46
	v76 = v50
	goto L19
L19:
	;
	if v76 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v58 = v44
	v59 = v46
	v60 = v50
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v63 != v64 {
		v81 = v58
		v82 = v59
		v83 = v60
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v74 = v69
	v75 = v67
	v76 = v71
	goto L19
L23:
	;
	v66 = int32(4)
	v67 = v59 + v66
	v69 = v58 + v66
	v71 = v60 - v66
	if base.Ui32(int32(3)) < base.Ui32(v71) {
		v58 = v69
		v59 = v67
		v60 = v71
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v81 = v74
	v82 = v75
	v83 = v76
	goto L16
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v112 = v91 - v92
	goto L14
L28:
	;
	v94 = int32(1)
	v99 = v88 - v94
	if v99 != 0 {
		v86 = v86 + v94
		v87 = v87 + v94
		v88 = v99
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
	if v47 == v48 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v112 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v116 = int32(10)
	v178 = (v35*v116 + v116) * (v47 - v48)
	goto L4
L36:
	;
	v128 = int32(-10)
	goto L38
L37:
	;
	v128 = int32(10)
	goto L38
L38:
	;
	v178 = (v35 + int32(1)) * v128
	goto L4
L39:
	;
	v134 = int32(9)
	v136 = int32(131064)
	v144 = int32(1)
	if v144 < v42 {
		v33 = v33 + (v47+v134)&v136
		v35 = v131
		v38 = v38 + (v48+v134)&v136
		v42 = v42 - v144
		goto L8
	} else {
		goto L40
	}
L40:
	;
	goto L9
L41:
	;
	F_pfree(m, v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v183 != v19 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return int32(base.Ui32(v178) >> (uint(int32(31)) % 32))
L48:
	;
	goto L47
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
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	if base.Ui32(v18) < base.Ui32(v17) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v122 = int32(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v122 = int32(1)
	goto L3
L9:
	;
	goto L10
L10:
	;
	v23 = int32(8)
	v30 = v17
	v31 = v15 + v23
	v32 = v10 + v23
	goto L11
L11:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
	if v35 != v36 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v122 = v116
	goto L3
L13:
	;
	v122 = int32(0)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v39 = int32(2)
	v40 = v32 + v39
	v42 = v31 + v39
	if base.Ui32(int32(4)) <= base.Ui32(v35) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v40|v42)&int32(3) != 0 {
		v73 = v40
		v74 = v42
		v75 = v35
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v40
	v67 = v42
	v68 = v35
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v40
	v51 = v42
	v52 = v35
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	v122 = int32(0)
	goto L3
L35:
	;
	goto L36
L36:
	;
	v106 = int32(9)
	v108 = int32(131064)
	v116 = int32(1)
	if v116 < v30 {
		v30 = v30 - v116
		v31 = v31 + (v36+v106)&v108
		v32 = v32 + (v35+v106)&v108
		goto L11
	} else {
		goto L37
	}
L37:
	;
	goto L12
L38:
	;
	F_pfree(m, v10)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v133 != v15 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	F_pfree(m, v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	return v122
L45:
	;
	goto L44
}
