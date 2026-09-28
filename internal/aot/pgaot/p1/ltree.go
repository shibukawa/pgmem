package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(0)
	if v11 == v12 {
		v28 = v12
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
		if v15 == int32(0) {
			v28 = v12
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v18 != int32(7) {
				v28 = v12
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				if v21 != int32(17) {
					v28 = v12
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+32)))
					v28 = v24 ^ int32(1)
				}
			}
		}
	}
	if v28&int32(1) != 0 {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v32 = F_get_fn_opclass_options(m, v31)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
			v38 = v36
			v40 = Fn14312(m, base.I32_wrap_i64(v6), base.I32_wrap_i64(v9), v38, int32(2))
			mBase = m.M
			*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = base.F32_convert_i32_s(v40)
			return v3
		}
	} else {
		v38 = int32(28)
		v40 = Fn14312(m, base.I32_wrap_i64(v6), base.I32_wrap_i64(v9), v38, int32(2))
		mBase = m.M
		*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v3)))) = base.F32_convert_i32_s(v40)
		return v3
	}
}
func F__ltree_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v767 int32
	_ = v767
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v800 int32
	_ = v800
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v828 int32
	_ = v828
	v2 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v29 = base.I32_wrap_i64(v28)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == v2 {
		v48 = v2
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
		if v35 == int32(0) {
			v48 = v2
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v38 != int32(7) {
				v48 = v2
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				if v41 != int32(17) {
					v48 = v2
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
					v48 = v44 ^ int32(1)
				}
			}
		}
	}
	if v48&int32(1) != 0 {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v52 = F_get_fn_opclass_options(m, v51)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int64(0)
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
			v57 = v56
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v62 = (v58 + int32(_a_F__ltree_picksplit_0)) & int32(_a_F__ltree_picksplit_1)
			v66 = v62<<(uint(int32(1))%32) + int32(4)
			v67 = F_palloc(m, v66)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = v67
				v70 = F_palloc(m, v66)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v70
					if base.Ui32(int32(2)) <= base.Ui32(v62) {
						v76 = v27 + int32(8)
						v80 = int32(-1)
						v81 = v2
						v86 = int32(1)
						v92 = v2
						for {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v76+v86*int32(24))))
							v110 = v86 + int32(1)
							v111 = v110
							v112 = v80
							v113 = v81
							v119 = v110
							v124 = v92
							for {
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v76+v119*int32(24))))
								v142 = Fn14312(m, v108, v140, v57, int32(2))
								mBase = m.M
								v143 = base.B2i32(v112 < v142)
								if v112 < v142 {
									v144 = v142
								} else {
									v144 = v112
								}
								if v112 < v142 {
									v145 = v111
								} else {
									v145 = v124
								}
								if v112 < v142 {
									v146 = v86
								} else {
									v146 = v113
								}
								v148 = v111 + int32(1)
								v150 = v148 & int32(_a_F__ltree_picksplit_1)
								if base.Ui32(v150) <= base.Ui32(v62) {
									v111 = v148
									v112 = v144
									v113 = v146
									v119 = v150
									v124 = v145
									continue
								} else {
									break
								}
								break
							}
							if v110 != v62 {
								v80 = v144
								v81 = v146
								v86 = v110
								v92 = v145
								continue
							} else {
								break
							}
							break
						}
						v153 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v156 = v146
						v166 = v153
						v167 = v145
					} else {
						v156 = v2
						v166 = v70
						v167 = v2
					}
					v182 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v182
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v182
					v186 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v188 = v27 + int32(8)
					v190 = int32(_a_F__ltree_picksplit_1)
					v198 = base.B2i32(v156&v190 == v182) | base.B2i32(v167&v190 == v182)
					if v198 != 0 {
						v199 = int32(1)
					} else {
						v199 = v156
					}
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v188+v199&int32(_a_F__ltree_picksplit_1)*int32(24))))
					v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
					v213 = int32(0)
					v215 = F_ltree_gist_alloc(m, int32(base.Ui32(v206&int32(2))>>(uint(int32(1))%32)), v205+int32(8), v57, v213, v213)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int64(0)
					} else {
						if v198 != 0 {
							v218 = int32(2)
						} else {
							v218 = v167
						}
						v224 = *(*int32)(unsafe.Add(mBase, uint32(v188+v218&int32(_a_F__ltree_picksplit_1)*int32(24))))
						v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
						v232 = int32(0)
						v234 = F_ltree_gist_alloc(m, int32(base.Ui32(v225&int32(2))>>(uint(int32(1))%32)), v224+int32(8), v57, v232, v232)
						mBase = m.M
						v235 = m.ExcPending
						if v235 != 0 {
							return int64(0)
						} else {
							v237 = int32(_a_F__ltree_picksplit_1)
							v238 = v58 + v237
							v240 = v238 & v237
							v241 = F_palloc_mul(m, int32(8), v240)
							mBase = m.M
							v242 = m.ExcPending
							if v242 != 0 {
								return int64(0)
							} else {
								if v58&int32(_a_F__ltree_picksplit_1) == int32(1) {
									F_pg_qsort(m, v241, v240, int32(8), int32(_a_F__ltree_picksplit_2))
									mBase = m.M
									v250 = m.ExcPending
									if v250 != 0 {
										return int64(0)
									} else {
										v809 = v186
										v814 = v166
										v828 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v809))) = uint16(v828)
										*(*uint16)(unsafe.Add(mBase, uint32(v814))) = uint16(v828)
										*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = base.I64_extend_i32_u(v234)
										*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = base.I64_extend_i32_u(v215)
										return v28 & int64(4294967295)
									}
								} else {
									v251 = int32(1)
									v253 = v251
									v254 = v251
									for {
										v281 = v241 + v253<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v281-int32(8)))) = uint16(v254)
										v290 = *(*int32)(unsafe.Add(mBase, uint32(v188+v253*int32(24))))
										v292 = Fn14312(m, v215, v290, v57, int32(2))
										mBase = m.M
										v294 = Fn14312(m, v234, v290, v57, int32(2))
										mBase = m.M
										v295 = v292 - v294
										v297 = v295 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v281-int32(4)))) = v295 ^ v297 - v297
										v302 = v254 + int32(1)
										v303 = int32(_a_F__ltree_picksplit_1)
										v304 = v302 & v303
										if base.Ui32(v304) <= base.Ui32(v238&v303) {
											v253 = v304
											v254 = v302
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v241, v240, int32(8), int32(_a_F__ltree_picksplit_2))
									mBase = m.M
									v311 = m.ExcPending
									if v311 != 0 {
										return int64(0)
									} else {
										v312 = int32(1)
										if base.Ui32(v240) <= base.Ui32(v312) {
											v315 = v312
										} else {
											v315 = v240
										}
										v317 = v57 & int32(2147483644)
										v319 = v57 & int32(3)
										v320 = int32(8)
										v321 = v234 + v320
										v323 = v215 + v320
										v328 = base.B2i32(base.Ui32(v57) < base.Ui32(int32(4)))
										v329 = int32(0)
										v336 = v186
										v341 = v166
										for {
											v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v329<<(uint(int32(3))%32)))))
											if v199&int32(_a_F__ltree_picksplit_1) == v358 {
												*(*uint16)(unsafe.Add(mBase, uint32(v336))) = uint16(v199)
												v361 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v361 + int32(1)
												v780 = v336 + int32(2)
												v785 = v341
											} else {
												if v218&int32(_a_F__ltree_picksplit_1) == v358 {
													*(*uint16)(unsafe.Add(mBase, uint32(v341))) = uint16(v218)
													v767 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v767 + int32(1)
													v780 = v336
													v785 = v341 + int32(2)
												} else {
													v374 = *(*int32)(unsafe.Add(mBase, uint32(v188+v358*int32(24))))
													v376 = Fn14312(m, v215, v374, v57, int32(2))
													mBase = m.M
													v379 = Fn14312(m, v234, v374, v57, int32(2))
													mBase = m.M
													v381 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													v382 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
													v383 = v381 - v382
													if base.F64_lt(base.F64_convert_i32_s(v376), base.F64_add(base.F64_convert_i32_s(v379), base.F64_mul(base.F64_convert_i32_s(v383*v383*v383), float64(-1e-05)))) != 0 {
														v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
														if v391&int32(2) != 0 {
														} else {
															v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+4)))
															if v394&int32(2) != 0 {
																if v57 == int32(0) {
																} else {
																	base.MemoryFill(m, v323, int32(255), v57)
																}
															} else {
																if v57 <= int32(0) {
																} else {
																	v404 = v374 + int32(8)
																	v405 = int32(0)
																	if v328 == v405 {
																		v411 = v405
																		v412 = v405
																		for {
																			v436 = v411 + v323
																			v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
																			v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v404))))
																			v440 = v437 | v439
																			*(*uint8)(unsafe.Add(mBase, uint32(v436))) = uint8(v440)
																			v443 = v411 | int32(1)
																			v444 = v323 + v443
																			v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
																			v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443+v404))))
																			v448 = v445 | v447
																			*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v448)
																			v451 = v411 | int32(2)
																			v452 = v323 + v451
																			v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
																			v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v404))))
																			v456 = v453 | v455
																			*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v456)
																			v459 = v411 | int32(3)
																			v460 = v323 + v459
																			v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
																			v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v404))))
																			v464 = v461 | v463
																			*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v464)
																			v466 = int32(4)
																			v467 = v411 + v466
																			v469 = v412 + v466
																			if v469 != v317 {
																				v411 = v467
																				v412 = v469
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v319 == int32(0) {
																		} else {
																			v475 = v467
																			v501 = v475
																			v507 = v405
																			for {
																				v525 = v501 + v323
																				v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																				v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v404))))
																				v529 = v526 | v528
																				*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																				v531 = int32(1)
																				v534 = v507 + v531
																				if v534 != v319 {
																					v501 = v501 + v531
																					v507 = v534
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v475 = v405
																		v501 = v475
																		v507 = v405
																		for {
																			v525 = v501 + v323
																			v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																			v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v404))))
																			v529 = v526 | v528
																			*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																			v531 = int32(1)
																			v534 = v507 + v531
																			if v534 != v319 {
																				v501 = v501 + v531
																				v507 = v534
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
														*(*uint16)(unsafe.Add(mBase, uint32(v336))) = uint16(v358)
														v563 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v563 + int32(1)
														v780 = v336 + int32(2)
														v785 = v341
													} else {
														v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+4)))
														if v569&int32(2) != 0 {
														} else {
															v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+4)))
															if v572&int32(2) != 0 {
																if v57 == int32(0) {
																} else {
																	base.MemoryFill(m, v321, int32(255), v57)
																}
															} else {
																if v57 <= int32(0) {
																} else {
																	v582 = v374 + int32(8)
																	v583 = int32(0)
																	if v328 == v583 {
																		v589 = v583
																		v590 = v583
																		for {
																			v614 = v589 + v321
																			v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
																			v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v582))))
																			v618 = v615 | v617
																			*(*uint8)(unsafe.Add(mBase, uint32(v614))) = uint8(v618)
																			v621 = v589 | int32(1)
																			v622 = v321 + v621
																			v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
																			v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v582))))
																			v626 = v623 | v625
																			*(*uint8)(unsafe.Add(mBase, uint32(v622))) = uint8(v626)
																			v629 = v589 | int32(2)
																			v630 = v321 + v629
																			v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
																			v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+v582))))
																			v634 = v631 | v633
																			*(*uint8)(unsafe.Add(mBase, uint32(v630))) = uint8(v634)
																			v637 = v589 | int32(3)
																			v638 = v321 + v637
																			v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
																			v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v582))))
																			v642 = v639 | v641
																			*(*uint8)(unsafe.Add(mBase, uint32(v638))) = uint8(v642)
																			v644 = int32(4)
																			v645 = v589 + v644
																			v647 = v590 + v644
																			if v647 != v317 {
																				v589 = v645
																				v590 = v647
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v319 == int32(0) {
																		} else {
																			v653 = v645
																			v679 = v653
																			v685 = v583
																			for {
																				v703 = v679 + v321
																				v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
																				v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v582))))
																				v707 = v704 | v706
																				*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v707)
																				v709 = int32(1)
																				v712 = v685 + v709
																				if v712 != v319 {
																					v679 = v679 + v709
																					v685 = v712
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v653 = v583
																		v679 = v653
																		v685 = v583
																		for {
																			v703 = v679 + v321
																			v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
																			v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v582))))
																			v707 = v704 | v706
																			*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v707)
																			v709 = int32(1)
																			v712 = v685 + v709
																			if v712 != v319 {
																				v679 = v679 + v709
																				v685 = v712
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
														*(*uint16)(unsafe.Add(mBase, uint32(v341))) = uint16(v358)
														v767 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v767 + int32(1)
														v780 = v336
														v785 = v341 + int32(2)
													}
												}
											}
											v800 = v329 + int32(1)
											if v800 != v315 {
												v329 = v800
												v336 = v780
												v341 = v785
												continue
											} else {
												break
											}
											break
										}
										v809 = v780
										v814 = v785
										v828 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v809))) = uint16(v828)
										*(*uint16)(unsafe.Add(mBase, uint32(v814))) = uint16(v828)
										*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = base.I64_extend_i32_u(v234)
										*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = base.I64_extend_i32_u(v215)
										return v28 & int64(4294967295)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v57 = int32(28)
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		v62 = (v58 + int32(_a_F__ltree_picksplit_0)) & int32(_a_F__ltree_picksplit_1)
		v66 = v62<<(uint(int32(1))%32) + int32(4)
		v67 = F_palloc(m, v66)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v67
			v70 = F_palloc(m, v66)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v70
				if base.Ui32(int32(2)) <= base.Ui32(v62) {
					v76 = v27 + int32(8)
					v80 = int32(-1)
					v81 = v2
					v86 = int32(1)
					v92 = v2
					for {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v76+v86*int32(24))))
						v110 = v86 + int32(1)
						v111 = v110
						v112 = v80
						v113 = v81
						v119 = v110
						v124 = v92
						for {
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v76+v119*int32(24))))
							v142 = Fn14312(m, v108, v140, v57, int32(2))
							mBase = m.M
							v143 = base.B2i32(v112 < v142)
							if v112 < v142 {
								v144 = v142
							} else {
								v144 = v112
							}
							if v112 < v142 {
								v145 = v111
							} else {
								v145 = v124
							}
							if v112 < v142 {
								v146 = v86
							} else {
								v146 = v113
							}
							v148 = v111 + int32(1)
							v150 = v148 & int32(_a_F__ltree_picksplit_1)
							if base.Ui32(v150) <= base.Ui32(v62) {
								v111 = v148
								v112 = v144
								v113 = v146
								v119 = v150
								v124 = v145
								continue
							} else {
								break
							}
							break
						}
						if v110 != v62 {
							v80 = v144
							v81 = v146
							v86 = v110
							v92 = v145
							continue
						} else {
							break
						}
						break
					}
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					v156 = v146
					v166 = v153
					v167 = v145
				} else {
					v156 = v2
					v166 = v70
					v167 = v2
				}
				v182 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v182
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v182
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v188 = v27 + int32(8)
				v190 = int32(_a_F__ltree_picksplit_1)
				v198 = base.B2i32(v156&v190 == v182) | base.B2i32(v167&v190 == v182)
				if v198 != 0 {
					v199 = int32(1)
				} else {
					v199 = v156
				}
				v205 = *(*int32)(unsafe.Add(mBase, uint32(v188+v199&int32(_a_F__ltree_picksplit_1)*int32(24))))
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
				v213 = int32(0)
				v215 = F_ltree_gist_alloc(m, int32(base.Ui32(v206&int32(2))>>(uint(int32(1))%32)), v205+int32(8), v57, v213, v213)
				mBase = m.M
				v216 = m.ExcPending
				if v216 != 0 {
					return int64(0)
				} else {
					if v198 != 0 {
						v218 = int32(2)
					} else {
						v218 = v167
					}
					v224 = *(*int32)(unsafe.Add(mBase, uint32(v188+v218&int32(_a_F__ltree_picksplit_1)*int32(24))))
					v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
					v232 = int32(0)
					v234 = F_ltree_gist_alloc(m, int32(base.Ui32(v225&int32(2))>>(uint(int32(1))%32)), v224+int32(8), v57, v232, v232)
					mBase = m.M
					v235 = m.ExcPending
					if v235 != 0 {
						return int64(0)
					} else {
						v237 = int32(_a_F__ltree_picksplit_1)
						v238 = v58 + v237
						v240 = v238 & v237
						v241 = F_palloc_mul(m, int32(8), v240)
						mBase = m.M
						v242 = m.ExcPending
						if v242 != 0 {
							return int64(0)
						} else {
							if v58&int32(_a_F__ltree_picksplit_1) == int32(1) {
								F_pg_qsort(m, v241, v240, int32(8), int32(_a_F__ltree_picksplit_2))
								mBase = m.M
								v250 = m.ExcPending
								if v250 != 0 {
									return int64(0)
								} else {
									v809 = v186
									v814 = v166
									v828 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v809))) = uint16(v828)
									*(*uint16)(unsafe.Add(mBase, uint32(v814))) = uint16(v828)
									*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = base.I64_extend_i32_u(v234)
									*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = base.I64_extend_i32_u(v215)
									return v28 & int64(4294967295)
								}
							} else {
								v251 = int32(1)
								v253 = v251
								v254 = v251
								for {
									v281 = v241 + v253<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v281-int32(8)))) = uint16(v254)
									v290 = *(*int32)(unsafe.Add(mBase, uint32(v188+v253*int32(24))))
									v292 = Fn14312(m, v215, v290, v57, int32(2))
									mBase = m.M
									v294 = Fn14312(m, v234, v290, v57, int32(2))
									mBase = m.M
									v295 = v292 - v294
									v297 = v295 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v281-int32(4)))) = v295 ^ v297 - v297
									v302 = v254 + int32(1)
									v303 = int32(_a_F__ltree_picksplit_1)
									v304 = v302 & v303
									if base.Ui32(v304) <= base.Ui32(v238&v303) {
										v253 = v304
										v254 = v302
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v241, v240, int32(8), int32(_a_F__ltree_picksplit_2))
								mBase = m.M
								v311 = m.ExcPending
								if v311 != 0 {
									return int64(0)
								} else {
									v312 = int32(1)
									if base.Ui32(v240) <= base.Ui32(v312) {
										v315 = v312
									} else {
										v315 = v240
									}
									v317 = v57 & int32(2147483644)
									v319 = v57 & int32(3)
									v320 = int32(8)
									v321 = v234 + v320
									v323 = v215 + v320
									v328 = base.B2i32(base.Ui32(v57) < base.Ui32(int32(4)))
									v329 = int32(0)
									v336 = v186
									v341 = v166
									for {
										v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241+v329<<(uint(int32(3))%32)))))
										if v199&int32(_a_F__ltree_picksplit_1) == v358 {
											*(*uint16)(unsafe.Add(mBase, uint32(v336))) = uint16(v199)
											v361 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v361 + int32(1)
											v780 = v336 + int32(2)
											v785 = v341
										} else {
											if v218&int32(_a_F__ltree_picksplit_1) == v358 {
												*(*uint16)(unsafe.Add(mBase, uint32(v341))) = uint16(v218)
												v767 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v767 + int32(1)
												v780 = v336
												v785 = v341 + int32(2)
											} else {
												v374 = *(*int32)(unsafe.Add(mBase, uint32(v188+v358*int32(24))))
												v376 = Fn14312(m, v215, v374, v57, int32(2))
												mBase = m.M
												v379 = Fn14312(m, v234, v374, v57, int32(2))
												mBase = m.M
												v381 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												v382 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
												v383 = v381 - v382
												if base.F64_lt(base.F64_convert_i32_s(v376), base.F64_add(base.F64_convert_i32_s(v379), base.F64_mul(base.F64_convert_i32_s(v383*v383*v383), float64(-1e-05)))) != 0 {
													v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
													if v391&int32(2) != 0 {
													} else {
														v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+4)))
														if v394&int32(2) != 0 {
															if v57 == int32(0) {
															} else {
																base.MemoryFill(m, v323, int32(255), v57)
															}
														} else {
															if v57 <= int32(0) {
															} else {
																v404 = v374 + int32(8)
																v405 = int32(0)
																if v328 == v405 {
																	v411 = v405
																	v412 = v405
																	for {
																		v436 = v411 + v323
																		v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
																		v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v404))))
																		v440 = v437 | v439
																		*(*uint8)(unsafe.Add(mBase, uint32(v436))) = uint8(v440)
																		v443 = v411 | int32(1)
																		v444 = v323 + v443
																		v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
																		v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443+v404))))
																		v448 = v445 | v447
																		*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v448)
																		v451 = v411 | int32(2)
																		v452 = v323 + v451
																		v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
																		v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v404))))
																		v456 = v453 | v455
																		*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v456)
																		v459 = v411 | int32(3)
																		v460 = v323 + v459
																		v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
																		v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v404))))
																		v464 = v461 | v463
																		*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v464)
																		v466 = int32(4)
																		v467 = v411 + v466
																		v469 = v412 + v466
																		if v469 != v317 {
																			v411 = v467
																			v412 = v469
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v319 == int32(0) {
																	} else {
																		v475 = v467
																		v501 = v475
																		v507 = v405
																		for {
																			v525 = v501 + v323
																			v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																			v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v404))))
																			v529 = v526 | v528
																			*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																			v531 = int32(1)
																			v534 = v507 + v531
																			if v534 != v319 {
																				v501 = v501 + v531
																				v507 = v534
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v475 = v405
																	v501 = v475
																	v507 = v405
																	for {
																		v525 = v501 + v323
																		v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
																		v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v404))))
																		v529 = v526 | v528
																		*(*uint8)(unsafe.Add(mBase, uint32(v525))) = uint8(v529)
																		v531 = int32(1)
																		v534 = v507 + v531
																		if v534 != v319 {
																			v501 = v501 + v531
																			v507 = v534
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
													*(*uint16)(unsafe.Add(mBase, uint32(v336))) = uint16(v358)
													v563 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v563 + int32(1)
													v780 = v336 + int32(2)
													v785 = v341
												} else {
													v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+4)))
													if v569&int32(2) != 0 {
													} else {
														v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+4)))
														if v572&int32(2) != 0 {
															if v57 == int32(0) {
															} else {
																base.MemoryFill(m, v321, int32(255), v57)
															}
														} else {
															if v57 <= int32(0) {
															} else {
																v582 = v374 + int32(8)
																v583 = int32(0)
																if v328 == v583 {
																	v589 = v583
																	v590 = v583
																	for {
																		v614 = v589 + v321
																		v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
																		v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v582))))
																		v618 = v615 | v617
																		*(*uint8)(unsafe.Add(mBase, uint32(v614))) = uint8(v618)
																		v621 = v589 | int32(1)
																		v622 = v321 + v621
																		v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
																		v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v582))))
																		v626 = v623 | v625
																		*(*uint8)(unsafe.Add(mBase, uint32(v622))) = uint8(v626)
																		v629 = v589 | int32(2)
																		v630 = v321 + v629
																		v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
																		v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+v582))))
																		v634 = v631 | v633
																		*(*uint8)(unsafe.Add(mBase, uint32(v630))) = uint8(v634)
																		v637 = v589 | int32(3)
																		v638 = v321 + v637
																		v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
																		v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637+v582))))
																		v642 = v639 | v641
																		*(*uint8)(unsafe.Add(mBase, uint32(v638))) = uint8(v642)
																		v644 = int32(4)
																		v645 = v589 + v644
																		v647 = v590 + v644
																		if v647 != v317 {
																			v589 = v645
																			v590 = v647
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v319 == int32(0) {
																	} else {
																		v653 = v645
																		v679 = v653
																		v685 = v583
																		for {
																			v703 = v679 + v321
																			v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
																			v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v582))))
																			v707 = v704 | v706
																			*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v707)
																			v709 = int32(1)
																			v712 = v685 + v709
																			if v712 != v319 {
																				v679 = v679 + v709
																				v685 = v712
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v653 = v583
																	v679 = v653
																	v685 = v583
																	for {
																		v703 = v679 + v321
																		v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
																		v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679+v582))))
																		v707 = v704 | v706
																		*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v707)
																		v709 = int32(1)
																		v712 = v685 + v709
																		if v712 != v319 {
																			v679 = v679 + v709
																			v685 = v712
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
													*(*uint16)(unsafe.Add(mBase, uint32(v341))) = uint16(v358)
													v767 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v767 + int32(1)
													v780 = v336
													v785 = v341 + int32(2)
												}
											}
										}
										v800 = v329 + int32(1)
										if v800 != v315 {
											v329 = v800
											v336 = v780
											v341 = v785
											continue
										} else {
											break
										}
										break
									}
									v809 = v780
									v814 = v785
									v828 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v809))) = uint16(v828)
									*(*uint16)(unsafe.Add(mBase, uint32(v814))) = uint16(v828)
									*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = base.I64_extend_i32_u(v234)
									*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = base.I64_extend_i32_u(v215)
									return v28 & int64(4294967295)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F__ltree_r_isparent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltree_r_isparent_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F__ltree_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int64
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v257 int32
	_ = v257
	v2 = int32(0)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
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
	return int64(0)
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
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = int32(base.Ui32(v257) >> (uint(int32(2)) % 32))
	return base.I64_extend_i32_u(v51)
L14:
	;
	v56 = int32(8)
	v57 = v51 + v56
	v61 = v46 & int32(3)
	v68 = v53
	v77 = v2
	goto L15
L15:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v18+v56+v77*int32(24))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v88&int32(2) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v236 | int32(2)
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
	v96 = v87 + int32(8)
	v97 = int32(0)
	if base.B2i32(base.Ui32(v46) < base.Ui32(int32(4))) == v97 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v215 = v68
	goto L22
L22:
	;
	v232 = v77 + int32(1)
	if v232 < v215 {
		v68 = v215
		v77 = v232
		goto L15
	} else {
		goto L34
	}
L23:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v215 = v214
	goto L22
L24:
	;
	v102 = v97
	v103 = v97
	goto L27
L25:
	;
	v156 = v97
	goto L26
L26:
	;
	v172 = v156
	v181 = v97
	goto L31
L27:
	;
	v118 = v102 + v57
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v96))))
	v122 = v119 | v121
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v122)
	v125 = v102 | int32(1)
	v126 = v57 + v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+v96))))
	v130 = v127 | v129
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v130)
	v133 = v102 | int32(2)
	v134 = v57 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v96))))
	v138 = v135 | v137
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v138)
	v141 = v102 | int32(3)
	v142 = v57 + v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v96))))
	v146 = v143 | v145
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v146)
	v148 = int32(4)
	v149 = v102 + v148
	v151 = v103 + v148
	if v151 != v46&int32(2147483644) {
		v102 = v149
		v103 = v151
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
	v156 = v149
	goto L26
L31:
	;
	v187 = v172 + v57
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172+v96))))
	v191 = v188 | v190
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v191)
	v193 = int32(1)
	v196 = v181 + v193
	if v196 != v61 {
		v172 = v172 + v193
		v181 = v196
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
func F_ltree_cmp(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v159 != v14 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v157 = v22 - v21
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v40 = v21
	v41 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113 != 0 {
		v157 = v113
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
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
	if v48 != v49 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v48 - v49
	goto L4
L32:
	;
	goto L33
L33:
	;
	if v41 < int32(2) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(1)
	v120 = int32(9)
	v122 = int32(_a_F_ltree_cmp_0)
	if v118 < v40 {
		v35 = v35 + (v48+v120)&v122
		v36 = v36 + (v49+v120)&v122
		v40 = v40 - v118
		v41 = v41 - v118
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v163 != v19 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v19)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.I64_extend_i32_s(v157)
L43:
	;
	goto L42
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
					F_errfinish(m, int32(_a_F_ltree_concat_3), int32(382), int32(_a_F_ltree_concat_4))
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
func F_ltree_consistent(m *base.Module, l0 int32) int64 {
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
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v766 int32
	_ = v766
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v870 int32
	_ = v870
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1044 int32
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1148 int32
	_ = v1148
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1224 int32
	_ = v1224
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1354 int32
	_ = v1354
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1445 int32
	_ = v1445
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1478 int64
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1493 int32
	_ = v1493
	var v1499 int32
	_ = v1499
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1524 int32
	_ = v1524
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1543 int32
	_ = v1543
	var v1551 int32
	_ = v1551
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1608 int32
	_ = v1608
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1689 int32
	_ = v1689
	var v1695 int32
	_ = v1695
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1780 int32
	_ = v1780
	var v1786 int32
	_ = v1786
	var v1813 int32
	_ = v1813
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1845 int64
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1853 int32
	_ = v1853
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1879 int64
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1904 int32
	_ = v1904
	var v1910 int32
	_ = v1910
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1966 int32
	_ = v1966
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1985 int32
	_ = v1985
	var v1993 int32
	_ = v1993
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2050 int32
	_ = v2050
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2131 int32
	_ = v2131
	var v2137 int32
	_ = v2137
	var v2151 int32
	_ = v2151
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2190 int32
	_ = v2190
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2222 int32
	_ = v2222
	var v2228 int32
	_ = v2228
	var v2255 int32
	_ = v2255
	var v2270 int32
	_ = v2270
	var v2290 int32
	_ = v2290
	var v2298 int32
	_ = v2298
	var v2305 int32
	_ = v2305
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2317 int32
	_ = v2317
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2416 int32
	_ = v2416
	var v2426 int32
	_ = v2426
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2466 int32
	_ = v2466
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v26 == v2 {
		v43 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v43&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v30 == int32(0) {
		v43 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v33 != int32(7) {
		v43 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v36 != int32(17) {
		v43 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+32)))
	v43 = v39 ^ int32(1)
	goto L2
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = F_get_fn_opclass_options(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v52 = int32(8)
	goto L9
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v54)
	v57 = base.I32_wrap_i64(v23) & int32(_a_F_ltree_consistent_0)
	switch v57 - int32(1) {
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
	return int64(0)
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v52 = v51
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L10
	} else {
		goto L527
	}
L13:
	;
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2445 != v2431 {
		goto L523
	} else {
		goto L524
	}
L14:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v2335&int32(3) != 0 {
		goto L505
	} else {
		goto L506
	}
L15:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1865 = F_pg_detoast_datum(m, v1864)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L10
	} else {
		goto L410
	}
L16:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1830 = F_pg_detoast_datum(m, v1829)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L10
	} else {
		goto L403
	}
L17:
	;
	v1464 = v53 + int32(8)
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1466 = F_pg_detoast_datum(m, v1465)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L10
	} else {
		goto L333
	}
L18:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1175 = F_pg_detoast_datum(m, v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L10
	} else {
		goto L255
	}
L19:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v884 = F_pg_detoast_datum_copy(m, v883)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L10
	} else {
		goto L188
	}
L20:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v664 = F_pg_detoast_datum(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L10
	} else {
		goto L142
	}
L21:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v554 = F_pg_detoast_datum(m, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L120
	}
L22:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v259 = F_pg_detoast_datum(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L61
	}
L23:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v161 = F_pg_detoast_datum(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L42
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v61 = F_pg_detoast_datum(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+16)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v64)+12)))
	if v66&int32(1) == int32(0) {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v73 = int32(0)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(8))+4)))
	if base.B2i32(v80 == v73)|base.B2i32(v83 == v73) != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v2431 = v61
	v2434 = base.B2i32(int32(0) < v157)
	goto L13
L28:
	;
	v157 = v147
	goto L27
L29:
	;
	v147 = v80 - v83
	goto L28
L30:
	;
	v91 = v61 + int32(8)
	v92 = v53 + int32(16)
	v95 = v83
	v96 = v80
	goto L31
L31:
	;
	v100 = int32(2)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92))))
	if base.Ui32(v104) < base.Ui32(v105) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L29
L33:
	;
	v107 = v104
	goto L35
L34:
	;
	v107 = v105
	goto L35
L35:
	;
	v108 = F_memcmp(m, v91+v100, v92+v100, v107)
	mBase = m.M
	if v108 != 0 {
		v147 = v108
		goto L28
	} else {
		goto L36
	}
L36:
	;
	if v104 != v105 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v157 = v104 - v105
	goto L27
L38:
	;
	goto L39
L39:
	;
	if v96 < int32(2) {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v113 = int32(1)
	v115 = int32(9)
	v117 = int32(_a_F_ltree_consistent_1)
	if v113 < v95 {
		v91 = v91 + (v104+v115)&v117
		v92 = v92 + (v105+v115)&v117
		v95 = v95 - v113
		v96 = v96 - v113
		goto L31
	} else {
		goto L41
	}
L41:
	;
	goto L32
L42:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v164&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v167 = int32(0)
	goto L45
L44:
	;
	v167 = v52
	goto L45
L45:
	;
	v168 = v53 + v167
	v171 = int32(0)
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+4)))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168+int32(8))+4)))
	if base.B2i32(v178 == v171)|base.B2i32(v181 == v171) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v2431 = v161
	v2434 = base.B2i32(int32(0) <= v255)
	goto L13
L47:
	;
	v255 = v245
	goto L46
L48:
	;
	v245 = v178 - v181
	goto L47
L49:
	;
	v189 = v161 + int32(8)
	v190 = v168 + int32(16)
	v193 = v181
	v194 = v178
	goto L50
L50:
	;
	v198 = int32(2)
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
	if base.Ui32(v202) < base.Ui32(v203) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L48
L52:
	;
	v205 = v202
	goto L54
L53:
	;
	v205 = v203
	goto L54
L54:
	;
	v206 = F_memcmp(m, v189+v198, v190+v198, v205)
	mBase = m.M
	if v206 != 0 {
		v245 = v206
		goto L47
	} else {
		goto L55
	}
L55:
	;
	if v202 != v203 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v255 = v202 - v203
	goto L46
L57:
	;
	goto L58
L58:
	;
	if v194 < int32(2) {
		goto L48
	} else {
		goto L59
	}
L59:
	;
	v211 = int32(1)
	v213 = int32(9)
	v215 = int32(_a_F_ltree_consistent_1)
	if v211 < v193 {
		v189 = v189 + (v202+v213)&v215
		v190 = v190 + (v203+v213)&v215
		v193 = v193 - v211
		v194 = v194 - v211
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v262)+12)))
	if v264&int32(1) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v269 = int32(0)
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+4)))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(8))+4)))
	if base.B2i32(v276 == v269)|base.B2i32(v279 == v269) != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	goto L64
L64:
	;
	v357 = v53 + int32(8)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v359&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L65:
	;
	v2431 = v259
	v2434 = base.B2i32(v353 == int32(0))
	goto L13
L66:
	;
	v353 = v343
	goto L65
L67:
	;
	v343 = v276 - v279
	goto L66
L68:
	;
	v287 = v259 + int32(8)
	v288 = v53 + int32(16)
	v291 = v279
	v292 = v276
	goto L69
L69:
	;
	v296 = int32(2)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287))))
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288))))
	if base.Ui32(v300) < base.Ui32(v301) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L67
L71:
	;
	v303 = v300
	goto L73
L72:
	;
	v303 = v301
	goto L73
L73:
	;
	v304 = F_memcmp(m, v287+v296, v288+v296, v303)
	mBase = m.M
	if v304 != 0 {
		v343 = v304
		goto L66
	} else {
		goto L74
	}
L74:
	;
	if v300 != v301 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v353 = v300 - v301
	goto L65
L76:
	;
	goto L77
L77:
	;
	if v292 < int32(2) {
		goto L67
	} else {
		goto L78
	}
L78:
	;
	v309 = int32(1)
	v311 = int32(9)
	v313 = int32(_a_F_ltree_consistent_1)
	if v309 < v291 {
		v287 = v287 + (v300+v311)&v313
		v288 = v288 + (v301+v311)&v313
		v291 = v291 - v309
		v292 = v292 - v309
		goto L69
	} else {
		goto L79
	}
L79:
	;
	goto L70
L80:
	;
	v362 = int32(0)
	goto L82
L81:
	;
	v362 = v52
	goto L82
L82:
	;
	v363 = v357 + v362
	v364 = int32(0)
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+4)))
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363)+4)))
	if base.B2i32(v371 == v364)|base.B2i32(v374 == v364) != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	if v448 < int32(0) {
		v2431 = v259
		v2434 = v2
		goto L13
	} else {
		goto L98
	}
L84:
	;
	v448 = v438
	goto L83
L85:
	;
	v438 = v371 - v374
	goto L84
L86:
	;
	v378 = int32(8)
	v382 = v259 + v378
	v383 = v363 + v378
	v386 = v374
	v387 = v371
	goto L87
L87:
	;
	v391 = int32(2)
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382))))
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v383))))
	if base.Ui32(v395) < base.Ui32(v396) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L85
L89:
	;
	v398 = v395
	goto L91
L90:
	;
	v398 = v396
	goto L91
L91:
	;
	v399 = F_memcmp(m, v382+v391, v383+v391, v398)
	mBase = m.M
	if v399 != 0 {
		v438 = v399
		goto L84
	} else {
		goto L92
	}
L92:
	;
	if v395 != v396 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v448 = v395 - v396
	goto L83
L94:
	;
	goto L95
L95:
	;
	if v387 < int32(2) {
		goto L85
	} else {
		goto L96
	}
L96:
	;
	v404 = int32(1)
	v406 = int32(9)
	v408 = int32(_a_F_ltree_consistent_1)
	if v404 < v386 {
		v382 = v382 + (v395+v406)&v408
		v383 = v383 + (v396+v406)&v408
		v386 = v386 - v404
		v387 = v387 - v404
		goto L87
	} else {
		goto L97
	}
L97:
	;
	goto L88
L98:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v451&int32(1) != 0 {
		v465 = v357
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v466 = int32(0)
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+4)))
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v465)+4)))
	if base.B2i32(v473 == v466)|base.B2i32(v476 == v466) != 0 {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	if v451&int32(2) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v457 = int32(0)
	goto L103
L102:
	;
	v457 = v52
	goto L103
L103:
	;
	v458 = v357 + v457
	if v451&int32(4) != 0 {
		v465 = v458
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	v465 = v458 + int32(base.Ui32(v461)>>(uint(int32(2))%32))
	goto L99
L105:
	;
	v2431 = v259
	v2434 = base.B2i32(v550 <= int32(0))
	goto L13
L106:
	;
	v550 = v540
	goto L105
L107:
	;
	v540 = v473 - v476
	goto L106
L108:
	;
	v480 = int32(8)
	v484 = v259 + v480
	v485 = v465 + v480
	v488 = v476
	v489 = v473
	goto L109
L109:
	;
	v493 = int32(2)
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v484))))
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v485))))
	if base.Ui32(v497) < base.Ui32(v498) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L107
L111:
	;
	v500 = v497
	goto L113
L112:
	;
	v500 = v498
	goto L113
L113:
	;
	v501 = F_memcmp(m, v484+v493, v485+v493, v500)
	mBase = m.M
	if v501 != 0 {
		v540 = v501
		goto L106
	} else {
		goto L114
	}
L114:
	;
	if v497 != v498 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v550 = v497 - v498
	goto L105
L116:
	;
	goto L117
L117:
	;
	if v489 < int32(2) {
		goto L107
	} else {
		goto L118
	}
L118:
	;
	v506 = int32(1)
	v508 = int32(9)
	v510 = int32(_a_F_ltree_consistent_1)
	if v506 < v488 {
		v484 = v484 + (v497+v508)&v510
		v485 = v485 + (v498+v508)&v510
		v488 = v488 - v506
		v489 = v489 - v506
		goto L109
	} else {
		goto L119
	}
L119:
	;
	goto L110
L120:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v558&int32(1) != 0 {
		v575 = v53 + int32(8)
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v576 = int32(0)
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v554)+4)))
	v586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+4)))
	if base.B2i32(v583 == v576)|base.B2i32(v586 == v576) != 0 {
		goto L129
	} else {
		goto L130
	}
L122:
	;
	if v558&int32(2) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v564 = int32(0)
	goto L125
L124:
	;
	v564 = v52
	goto L125
L125:
	;
	v567 = v53 + v564 + int32(8)
	if v558&int32(4) != 0 {
		v575 = v567
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v575 = v567 + int32(base.Ui32(v570)>>(uint(int32(2))%32))
	goto L121
L127:
	;
	v2431 = v554
	v2434 = base.B2i32(v660 <= int32(0))
	goto L13
L128:
	;
	v660 = v650
	goto L127
L129:
	;
	v650 = v583 - v586
	goto L128
L130:
	;
	v590 = int32(8)
	v594 = v554 + v590
	v595 = v575 + v590
	v598 = v586
	v599 = v583
	goto L131
L131:
	;
	v603 = int32(2)
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v594))))
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595))))
	if base.Ui32(v607) < base.Ui32(v608) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L129
L133:
	;
	v610 = v607
	goto L135
L134:
	;
	v610 = v608
	goto L135
L135:
	;
	v611 = F_memcmp(m, v594+v603, v595+v603, v610)
	mBase = m.M
	if v611 != 0 {
		v650 = v611
		goto L128
	} else {
		goto L136
	}
L136:
	;
	if v607 != v608 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v660 = v607 - v608
	goto L127
L138:
	;
	goto L139
L139:
	;
	if v599 < int32(2) {
		goto L129
	} else {
		goto L140
	}
L140:
	;
	v616 = int32(1)
	v618 = int32(9)
	v620 = int32(_a_F_ltree_consistent_1)
	if v616 < v598 {
		v594 = v594 + (v607+v618)&v620
		v595 = v595 + (v608+v618)&v620
		v598 = v598 - v616
		v599 = v599 - v616
		goto L131
	} else {
		goto L141
	}
L141:
	;
	goto L132
L142:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v667 = int32(1)
	v668 = v666 & v667
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v669)+16)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669+v670)+12)))
	if v672&v667 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if v668 != 0 {
		v691 = v53 + int32(8)
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	if v668 != 0 {
		v795 = v53 + int32(8)
		goto L167
	} else {
		goto L168
	}
L146:
	;
	v692 = int32(0)
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+4)))
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v691)+4)))
	if base.B2i32(v699 == v692)|base.B2i32(v702 == v692) != 0 {
		goto L154
	} else {
		goto L155
	}
L147:
	;
	if v666&int32(2) != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v680 = int32(0)
	goto L150
L149:
	;
	v680 = v52
	goto L150
L150:
	;
	v683 = v53 + v680 + int32(8)
	if v666&int32(4) != 0 {
		v691 = v683
		goto L146
	} else {
		goto L151
	}
L151:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	v691 = v683 + int32(base.Ui32(v686)>>(uint(int32(2))%32))
	goto L146
L152:
	;
	v2431 = v664
	v2434 = int32(base.Ui32(v776) >> (uint(int32(31)) % 32))
	goto L13
L153:
	;
	v776 = v766
	goto L152
L154:
	;
	v766 = v699 - v702
	goto L153
L155:
	;
	v706 = int32(8)
	v710 = v664 + v706
	v711 = v691 + v706
	v714 = v702
	v715 = v699
	goto L156
L156:
	;
	v719 = int32(2)
	v723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v710))))
	v724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v711))))
	if base.Ui32(v723) < base.Ui32(v724) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L154
L158:
	;
	v726 = v723
	goto L160
L159:
	;
	v726 = v724
	goto L160
L160:
	;
	v727 = F_memcmp(m, v710+v719, v711+v719, v726)
	mBase = m.M
	if v727 != 0 {
		v766 = v727
		goto L153
	} else {
		goto L161
	}
L161:
	;
	if v723 != v724 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v776 = v723 - v724
	goto L152
L163:
	;
	goto L164
L164:
	;
	if v715 < int32(2) {
		goto L154
	} else {
		goto L165
	}
L165:
	;
	v732 = int32(1)
	v734 = int32(9)
	v736 = int32(_a_F_ltree_consistent_1)
	if v732 < v714 {
		v710 = v710 + (v723+v734)&v736
		v711 = v711 + (v724+v734)&v736
		v714 = v714 - v732
		v715 = v715 - v732
		goto L156
	} else {
		goto L166
	}
L166:
	;
	goto L157
L167:
	;
	v796 = int32(0)
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+4)))
	v806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v795)+4)))
	if base.B2i32(v803 == v796)|base.B2i32(v806 == v796) != 0 {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	if v666&int32(2) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v784 = int32(0)
	goto L171
L170:
	;
	v784 = v52
	goto L171
L171:
	;
	v787 = v53 + v784 + int32(8)
	if v666&int32(4) != 0 {
		v795 = v787
		goto L167
	} else {
		goto L172
	}
L172:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)))
	v795 = v787 + int32(base.Ui32(v790)>>(uint(int32(2))%32))
	goto L167
L173:
	;
	v2431 = v664
	v2434 = base.B2i32(v880 <= int32(0))
	goto L13
L174:
	;
	v880 = v870
	goto L173
L175:
	;
	v870 = v803 - v806
	goto L174
L176:
	;
	v810 = int32(8)
	v814 = v664 + v810
	v815 = v795 + v810
	v818 = v806
	v819 = v803
	goto L177
L177:
	;
	v823 = int32(2)
	v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v814))))
	v828 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815))))
	if base.Ui32(v827) < base.Ui32(v828) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L175
L179:
	;
	v830 = v827
	goto L181
L180:
	;
	v830 = v828
	goto L181
L181:
	;
	v831 = F_memcmp(m, v814+v823, v815+v823, v830)
	mBase = m.M
	if v831 != 0 {
		v870 = v831
		goto L174
	} else {
		goto L182
	}
L182:
	;
	if v827 != v828 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v880 = v827 - v828
	goto L173
L184:
	;
	goto L185
L185:
	;
	if v819 < int32(2) {
		goto L175
	} else {
		goto L186
	}
L186:
	;
	v836 = int32(1)
	v838 = int32(9)
	v840 = int32(_a_F_ltree_consistent_1)
	if v836 < v818 {
		v814 = v814 + (v827+v838)&v840
		v815 = v815 + (v828+v838)&v840
		v818 = v818 - v836
		v819 = v819 - v836
		goto L177
	} else {
		goto L187
	}
L187:
	;
	goto L178
L188:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v886)+16)))
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886+v887)+12)))
	if v889&int32(1) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v897 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(8))+4)))
	v898 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)))
	if base.Ui32(v898) < base.Ui32(v897) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v945 = v53 + int32(8)
	v946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)))
	v948 = v946
	goto L208
L192:
	;
	v2431 = v884
	v2434 = v943
	goto L13
L193:
	;
	v943 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	if v897 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v943 = int32(1)
	goto L192
L197:
	;
	goto L198
L198:
	;
	v908 = v884 + int32(8)
	v909 = v53 + int32(16)
	v910 = v897
	goto L199
L199:
	;
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v908))))
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909))))
	if v913 != v914 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v943 = int32(1)
	goto L192
L201:
	;
	v943 = int32(0)
	goto L192
L202:
	;
	goto L203
L203:
	;
	v917 = int32(2)
	v921 = F_memcmp(m, v908+v917, v909+v917, v913)
	mBase = m.M
	if v921 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v943 = int32(0)
	goto L192
L205:
	;
	goto L206
L206:
	;
	v923 = int32(9)
	v925 = int32(_a_F_ltree_consistent_1)
	v933 = int32(1)
	if v933 < v910 {
		v908 = v908 + (v913+v923)&v925
		v909 = v909 + (v914+v923)&v925
		v910 = v910 - v933
		goto L199
	} else {
		goto L207
	}
L207:
	;
	goto L200
L208:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)) = uint16(v948)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v965&int32(3) != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)) = uint16(v946)
	v2431 = v884
	v2434 = v1170
	goto L13
L210:
	;
	goto L209
L211:
	;
	v968 = int32(0)
	goto L213
L212:
	;
	v968 = v52
	goto L213
L213:
	;
	v969 = v945 + v968
	v970 = int32(0)
	v977 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)))
	v980 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v969)+4)))
	if base.B2i32(v977 == v970)|base.B2i32(v980 == v970) != 0 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	if int32(0) <= v1054 {
		goto L229
	} else {
		goto L230
	}
L215:
	;
	v1054 = v1044
	goto L214
L216:
	;
	v1044 = v977 - v980
	goto L215
L217:
	;
	v984 = int32(8)
	v988 = v884 + v984
	v989 = v969 + v984
	v992 = v980
	v993 = v977
	goto L218
L218:
	;
	v997 = int32(2)
	v1001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v988))))
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v989))))
	if base.Ui32(v1001) < base.Ui32(v1002) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L216
L220:
	;
	v1004 = v1001
	goto L222
L221:
	;
	v1004 = v1002
	goto L222
L222:
	;
	v1005 = F_memcmp(m, v988+v997, v989+v997, v1004)
	mBase = m.M
	if v1005 != 0 {
		v1044 = v1005
		goto L215
	} else {
		goto L223
	}
L223:
	;
	if v1001 != v1002 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1054 = v1001 - v1002
	goto L214
L225:
	;
	goto L226
L226:
	;
	if v993 < int32(2) {
		goto L216
	} else {
		goto L227
	}
L227:
	;
	v1010 = int32(1)
	v1012 = int32(9)
	v1014 = int32(_a_F_ltree_consistent_1)
	if v1010 < v992 {
		v988 = v988 + (v1001+v1012)&v1014
		v989 = v989 + (v1002+v1012)&v1014
		v992 = v992 - v1010
		v993 = v993 - v1010
		goto L218
	} else {
		goto L228
	}
L228:
	;
	goto L219
L229:
	;
	v1057 = int32(1)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1058&v1057 != 0 {
		v1073 = v945
		goto L232
	} else {
		goto L233
	}
L230:
	;
	goto L231
L231:
	;
	v1164 = int32(0)
	if v1164 < v948 {
		v948 = v948 - int32(1)
		goto L208
	} else {
		goto L254
	}
L232:
	;
	v1074 = int32(0)
	v1081 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v884)+4)))
	v1084 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1073)+4)))
	if base.B2i32(v1081 == v1074)|base.B2i32(v1084 == v1074) != 0 {
		goto L240
	} else {
		goto L241
	}
L233:
	;
	if v1058&int32(2) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1064 = int32(0)
	goto L236
L235:
	;
	v1064 = v52
	goto L236
L236:
	;
	v1065 = v945 + v1064
	if v1058&int32(4) != 0 {
		v1073 = v1065
		goto L232
	} else {
		goto L237
	}
L237:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1065)))
	v1073 = v1065 + int32(base.Ui32(v1068)>>(uint(int32(2))%32))
	goto L232
L238:
	;
	if v1158 <= int32(0) {
		v1170 = v1057
		goto L210
	} else {
		goto L253
	}
L239:
	;
	v1158 = v1148
	goto L238
L240:
	;
	v1148 = v1081 - v1084
	goto L239
L241:
	;
	v1088 = int32(8)
	v1092 = v884 + v1088
	v1093 = v1073 + v1088
	v1096 = v1084
	v1097 = v1081
	goto L242
L242:
	;
	v1101 = int32(2)
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1092))))
	v1106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093))))
	if base.Ui32(v1105) < base.Ui32(v1106) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	goto L240
L244:
	;
	v1108 = v1105
	goto L246
L245:
	;
	v1108 = v1106
	goto L246
L246:
	;
	v1109 = F_memcmp(m, v1092+v1101, v1093+v1101, v1108)
	mBase = m.M
	if v1109 != 0 {
		v1148 = v1109
		goto L239
	} else {
		goto L247
	}
L247:
	;
	if v1105 != v1106 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1158 = v1105 - v1106
	goto L238
L249:
	;
	goto L250
L250:
	;
	if v1097 < int32(2) {
		goto L240
	} else {
		goto L251
	}
L251:
	;
	v1114 = int32(1)
	v1116 = int32(9)
	v1118 = int32(_a_F_ltree_consistent_1)
	if v1114 < v1096 {
		v1092 = v1092 + (v1105+v1116)&v1118
		v1093 = v1093 + (v1106+v1116)&v1118
		v1096 = v1096 - v1114
		v1097 = v1097 - v1114
		goto L242
	} else {
		goto L252
	}
L252:
	;
	goto L243
L253:
	;
	goto L231
L254:
	;
	v1170 = v1164
	goto L210
L255:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1177)+16)))
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1177+v1178)+12)))
	if v1180&int32(1) != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	v1189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53+int32(8))+4)))
	if base.Ui32(v1189) < base.Ui32(v1188) {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	goto L258
L258:
	;
	v1236 = v53 + int32(8)
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1238&int32(3) != 0 {
		goto L275
	} else {
		goto L276
	}
L259:
	;
	v2431 = v1175
	v2434 = v1234
	goto L13
L260:
	;
	v1234 = int32(0)
	goto L259
L261:
	;
	goto L262
L262:
	;
	if v1188 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1234 = int32(1)
	goto L259
L264:
	;
	goto L265
L265:
	;
	v1199 = v53 + int32(16)
	v1200 = v1175 + int32(8)
	v1201 = v1188
	goto L266
L266:
	;
	v1204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199))))
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1200))))
	if v1204 != v1205 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v1234 = int32(1)
	goto L259
L268:
	;
	v1234 = int32(0)
	goto L259
L269:
	;
	goto L270
L270:
	;
	v1208 = int32(2)
	v1212 = F_memcmp(m, v1199+v1208, v1200+v1208, v1204)
	mBase = m.M
	if v1212 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1234 = int32(0)
	goto L259
L272:
	;
	goto L273
L273:
	;
	v1214 = int32(9)
	v1216 = int32(_a_F_ltree_consistent_1)
	v1224 = int32(1)
	if v1224 < v1201 {
		v1199 = v1199 + (v1204+v1214)&v1216
		v1200 = v1200 + (v1205+v1214)&v1216
		v1201 = v1201 - v1224
		goto L266
	} else {
		goto L274
	}
L274:
	;
	goto L267
L275:
	;
	v1241 = int32(0)
	goto L277
L276:
	;
	v1241 = v52
	goto L277
L277:
	;
	v1242 = v1236 + v1241
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	v1246 = F_palloc0(m, int32(base.Ui32(v1243)>>(uint(int32(2))%32)))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L10
	} else {
		goto L278
	}
L278:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1242)))
	v1250 = int32(base.Ui32(v1248) >> (uint(int32(2)) % 32))
	if v1250 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	base.MemoryCopy(m, v1246, v1242, v1250)
	goto L281
L280:
	;
	goto L281
L281:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1252&int32(1) != 0 {
		v1266 = v1236
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1266)))
	v1270 = F_palloc0(m, int32(base.Ui32(v1267)>>(uint(int32(2))%32)))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L10
	} else {
		goto L288
	}
L283:
	;
	if v1252&int32(2) != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1258 = int32(0)
	goto L286
L285:
	;
	v1258 = v52
	goto L286
L286:
	;
	v1259 = v1236 + v1258
	if v1252&int32(4) != 0 {
		v1266 = v1259
		goto L282
	} else {
		goto L287
	}
L287:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1259)))
	v1266 = v1259 + int32(base.Ui32(v1262)>>(uint(int32(2))%32))
	goto L282
L288:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1266)))
	v1274 = int32(base.Ui32(v1272) >> (uint(int32(2)) % 32))
	if v1274 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	base.MemoryCopy(m, v1270, v1266, v1274)
	goto L291
L290:
	;
	goto L291
L291:
	;
	v1276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	v1277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1246)+4)))
	if base.Ui32(v1276) < base.Ui32(v1277) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1246)+4)) = uint16(v1276)
	goto L294
L293:
	;
	goto L294
L294:
	;
	v1280 = int32(0)
	v1287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	v1290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1246)+4)))
	if base.B2i32(v1287 == v1280)|base.B2i32(v1290 == v1280) != 0 {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	v1365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1270)+4)))
	if base.Ui32(v1365) < base.Ui32(v1366) {
		goto L310
	} else {
		goto L311
	}
L296:
	;
	v1364 = v1354
	goto L295
L297:
	;
	v1354 = v1287 - v1290
	goto L296
L298:
	;
	v1294 = int32(8)
	v1298 = v1175 + v1294
	v1299 = v1246 + v1294
	v1302 = v1290
	v1303 = v1287
	goto L299
L299:
	;
	v1307 = int32(2)
	v1311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1298))))
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1299))))
	if base.Ui32(v1311) < base.Ui32(v1312) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	goto L297
L301:
	;
	v1314 = v1311
	goto L303
L302:
	;
	v1314 = v1312
	goto L303
L303:
	;
	v1315 = F_memcmp(m, v1298+v1307, v1299+v1307, v1314)
	mBase = m.M
	if v1315 != 0 {
		v1354 = v1315
		goto L296
	} else {
		goto L304
	}
L304:
	;
	if v1311 != v1312 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1364 = v1311 - v1312
	goto L295
L306:
	;
	goto L307
L307:
	;
	if v1303 < int32(2) {
		goto L297
	} else {
		goto L308
	}
L308:
	;
	v1320 = int32(1)
	v1322 = int32(9)
	v1324 = int32(_a_F_ltree_consistent_1)
	if v1320 < v1302 {
		v1298 = v1298 + (v1311+v1322)&v1324
		v1299 = v1299 + (v1312+v1322)&v1324
		v1302 = v1302 - v1320
		v1303 = v1303 - v1320
		goto L299
	} else {
		goto L309
	}
L309:
	;
	goto L300
L310:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1270)+4)) = uint16(v1365)
	goto L312
L311:
	;
	goto L312
L312:
	;
	if int32(0) <= v1364 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1371 = int32(0)
	v1378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1175)+4)))
	v1381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1270)+4)))
	if base.B2i32(v1378 == v1371)|base.B2i32(v1381 == v1371) != 0 {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	v1458 = v2
	goto L315
L315:
	;
	F_pfree(m, v1246)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L10
	} else {
		goto L331
	}
L316:
	;
	v1458 = base.B2i32(v1455 <= int32(0))
	goto L315
L317:
	;
	v1455 = v1445
	goto L316
L318:
	;
	v1445 = v1378 - v1381
	goto L317
L319:
	;
	v1385 = int32(8)
	v1389 = v1175 + v1385
	v1390 = v1270 + v1385
	v1393 = v1381
	v1394 = v1378
	goto L320
L320:
	;
	v1398 = int32(2)
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1389))))
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1390))))
	if base.Ui32(v1402) < base.Ui32(v1403) {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	goto L318
L322:
	;
	v1405 = v1402
	goto L324
L323:
	;
	v1405 = v1403
	goto L324
L324:
	;
	v1406 = F_memcmp(m, v1389+v1398, v1390+v1398, v1405)
	mBase = m.M
	if v1406 != 0 {
		v1445 = v1406
		goto L317
	} else {
		goto L325
	}
L325:
	;
	if v1402 != v1403 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1455 = v1402 - v1403
	goto L316
L327:
	;
	goto L328
L328:
	;
	if v1394 < int32(2) {
		goto L318
	} else {
		goto L329
	}
L329:
	;
	v1411 = int32(1)
	v1413 = int32(9)
	v1415 = int32(_a_F_ltree_consistent_1)
	if v1411 < v1393 {
		v1389 = v1389 + (v1402+v1413)&v1415
		v1390 = v1390 + (v1403+v1413)&v1415
		v1393 = v1393 - v1411
		v1394 = v1394 - v1411
		goto L320
	} else {
		goto L330
	}
L330:
	;
	goto L321
L331:
	;
	F_pfree(m, v1270)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L10
	} else {
		goto L332
	}
L332:
	;
	v2431 = v1175
	v2434 = v1458
	goto L13
L333:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1468)+16)))
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1468+v1469)+12)))
	if v1471&int32(1) != 0 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1478 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_2), int32(0), base.I64_extend_i32_u(v1464), base.I64_extend_i32_u(v1466))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L10
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v1482&int32(2) != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2431 = v1466
	v2434 = base.B2i32(v1478 != int64(0))
	goto L13
L338:
	;
	v1608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1466)+6)))
	if v1608 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L339:
	;
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1466)+4)))
	if v1485 == int32(0) {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v1493 = v1485
	v1499 = v1466 + int32(16)
	goto L341
L341:
	;
	v1508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1499)+4)))
	if v1508 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	goto L338
L343:
	;
	v1571 = int32(1)
	v1573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1499))))
	if v1571 < v1493 {
		v1493 = v1493 - v1571
		v1499 = v1499 + (v1573+int32(7))&int32(_a_F_ltree_consistent_1)
		goto L341
	} else {
		goto L350
	}
L344:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1499)+2)))
	if v1511&int32(21) != 0 {
		goto L343
	} else {
		goto L345
	}
L345:
	;
	v1517 = v1499 + int32(16)
	v1524 = v1508
	goto L346
L346:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1517)))
	v1533 = base.I32_rem_u_s(v1532, v52<<(uint(int32(3))%32))
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464+int32(base.Ui32(v1533)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v1537)>>(uint(v1533&int32(7))%32))&int32(1) != 0 {
		goto L343
	} else {
		goto L348
	}
L347:
	;
	v2431 = v1466
	v2434 = v2
	goto L13
L348:
	;
	v1543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1517)+4)))
	v1551 = int32(1)
	if v1551 < v1524 {
		v1517 = v1517 + (v1543+int32(7))&int32(_a_F_ltree_consistent_1) + int32(8)
		v1524 = v1524 - v1551
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	goto L342
L351:
	;
	v2431 = v1466
	v2434 = v1828
	goto L13
L352:
	;
	v1828 = int32(1)
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1613 = v53 + int32(8)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v1615&int32(2) != 0 {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	if v1695 <= int32(0) {
		goto L378
	} else {
		goto L379
	}
L356:
	;
	if base.Ui32(v1623) < base.Ui32(v1608) {
		goto L375
	} else {
		goto L376
	}
L357:
	;
	v1618 = int32(0)
	goto L359
L358:
	;
	v1618 = v52
	goto L359
L359:
	;
	v1619 = v1613 + v1618
	v1621 = v1615 & int32(1)
	if v1621 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1622 = v1613
	goto L362
L361:
	;
	v1622 = v1619
	goto L362
L362:
	;
	v1623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1622)+4)))
	if v1623 == int32(0) {
		goto L356
	} else {
		goto L363
	}
L363:
	;
	v1630 = v1466 + int32(16)
	v1632 = v1622 + int32(8)
	v1637 = v1608
	v1638 = v1623
	goto L364
L364:
	;
	v1648 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1632))))
	v1649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1630)+20)))
	if base.Ui32(v1648) < base.Ui32(v1649) {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	goto L356
L366:
	;
	v1651 = v1648
	goto L368
L367:
	;
	v1651 = v1649
	goto L368
L368:
	;
	v1652 = F_memcmp(m, v1632+int32(2), v1630+int32(23), v1651)
	mBase = m.M
	if v1652 != 0 {
		v1695 = v1652
		goto L355
	} else {
		goto L369
	}
L369:
	;
	if v1648 != v1649 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1695 = v1648 - v1649
	goto L355
L371:
	;
	goto L372
L372:
	;
	if v1638 < int32(2) {
		goto L356
	} else {
		goto L373
	}
L373:
	;
	v1657 = int32(1)
	v1661 = int32(_a_F_ltree_consistent_1)
	v1664 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1630))))
	if v1657 < v1637 {
		v1630 = v1630 + (v1664+int32(7))&v1661
		v1632 = v1632 + (v1648+int32(9))&v1661
		v1637 = v1637 - v1657
		v1638 = v1638 - v1657
		goto L364
	} else {
		goto L374
	}
L374:
	;
	goto L365
L375:
	;
	v1689 = v1623
	goto L377
L376:
	;
	v1689 = v1608
	goto L377
L377:
	;
	v1695 = v1689 - v1608
	goto L355
L378:
	;
	if v1621 != 0 {
		v1713 = v1613
		goto L381
	} else {
		goto L382
	}
L379:
	;
	v1813 = int32(0)
	goto L380
L380:
	;
	v1828 = v1813
	goto L351
L381:
	;
	v1714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1713)+4)))
	if v1714 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L382:
	;
	if v1615&int32(4) != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1713 = v1619
	goto L381
L384:
	;
	goto L385
L385:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1619)))
	v1713 = v1619 + int32(base.Ui32(v1709)>>(uint(int32(2))%32))
	goto L381
L386:
	;
	v1813 = base.B2i32(int32(0) <= v1786)
	goto L380
L387:
	;
	if base.Ui32(v1714) < base.Ui32(v1608) {
		goto L400
	} else {
		goto L401
	}
L388:
	;
	v1721 = v1466 + int32(16)
	v1723 = v1713 + int32(8)
	v1728 = v1608
	v1729 = v1714
	goto L389
L389:
	;
	v1739 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1723))))
	v1740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1721)+20)))
	if base.Ui32(v1739) < base.Ui32(v1740) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	goto L387
L391:
	;
	v1742 = v1739
	goto L393
L392:
	;
	v1742 = v1740
	goto L393
L393:
	;
	v1743 = F_memcmp(m, v1723+int32(2), v1721+int32(23), v1742)
	mBase = m.M
	if v1743 != 0 {
		v1786 = v1743
		goto L386
	} else {
		goto L394
	}
L394:
	;
	if v1739 != v1740 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1786 = v1739 - v1740
	goto L386
L396:
	;
	goto L397
L397:
	;
	if v1729 < int32(2) {
		goto L387
	} else {
		goto L398
	}
L398:
	;
	v1748 = int32(1)
	v1752 = int32(_a_F_ltree_consistent_1)
	v1755 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1721))))
	if v1748 < v1728 {
		v1721 = v1721 + (v1755+int32(7))&v1752
		v1723 = v1723 + (v1739+int32(9))&v1752
		v1728 = v1728 - v1748
		v1729 = v1729 - v1748
		goto L389
	} else {
		goto L399
	}
L399:
	;
	goto L390
L400:
	;
	v1780 = v1714
	goto L402
L401:
	;
	v1780 = v1608
	goto L402
L402:
	;
	v1786 = v1780 - v1608
	goto L386
L403:
	;
	v1832 = int32(1)
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1833)+16)))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1833+v1834)+12)))
	if v1836&v1832 != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1845 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_3), int32(0), base.I64_extend_i32_u(v53+int32(8)), base.I64_extend_i32_u(v1830))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L10
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	if v1849&int32(2) != 0 {
		v2431 = v1830
		v2434 = v1832
		goto L13
	} else {
		goto L408
	}
L407:
	;
	v2431 = v1830
	v2434 = base.B2i32(v1845 != int64(0))
	goto L13
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v52
	v1853 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v53 + v1853
	v1862 = F_ltree_execute(m, v1830+v1853, v19+v1853, int32(0), int32(_a_F_ltree_consistent_4))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L10
	} else {
		goto L409
	}
L409:
	;
	v2431 = v1830
	v2434 = v1862
	goto L13
L410:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1867)+16)))
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1867+v1868)+12)))
	if v1870&int32(1) != 0 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1879 = F_DirectFunctionCall2Coll(m, int32(_a_F_ltree_consistent_5), int32(0), base.I64_extend_i32_u(v53+int32(8)), base.I64_extend_i32_u(v1865))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L10
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+8))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	v1887 = F_ArrayGetNItemsSafe(m, v1884, v1865+int32(16))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L10
	} else {
		goto L415
	}
L414:
	;
	v2431 = v1865
	v2434 = base.B2i32(v1879 != int64(0))
	goto L13
L415:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1865)+4))
	if v1889 < int32(2) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1892 = F_array_contains_nulls(m, v1865)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L10
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L10
	} else {
		goto L501
	}
L419:
	;
	if v1892 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	if v1887 <= int32(0) {
		v2431 = v1865
		v2434 = v2
		goto L13
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L10
	} else {
		goto L497
	}
L423:
	;
	if v1883 != 0 {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1904 = v1883
	goto L426
L425:
	;
	v1904 = (v1884<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L426
L426:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v1920 = v1865 + v1904
	v1924 = v1887
	goto L427
L427:
	;
	if v1910&int32(2) != 0 {
		goto L430
	} else {
		goto L431
	}
L428:
	;
	v2431 = v1865
	v2434 = v2
	goto L13
L429:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1920)))
	v2298 = int32(1)
	if v2298 < v1924 {
		v1920 = v1920 + (int32(base.Ui32(v2290)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v1924 = v1924 - v2298
		goto L427
	} else {
		goto L496
	}
L430:
	;
	v2050 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	if v2050 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L431:
	;
	v1929 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	if v1929 == int32(0) {
		goto L430
	} else {
		goto L432
	}
L432:
	;
	v1935 = v1929
	v1943 = v1920 + int32(16)
	goto L433
L433:
	;
	v1950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1943)+4)))
	if v1950 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	goto L430
L435:
	;
	v2013 = int32(1)
	v2015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1943))))
	if v2013 < v1935 {
		v1935 = v1935 - v2013
		v1943 = v1943 + (v2015+int32(7))&int32(_a_F_ltree_consistent_1)
		goto L433
	} else {
		goto L442
	}
L436:
	;
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1943)+2)))
	if v1953&int32(21) != 0 {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1959 = v1943 + int32(16)
	v1966 = v1950
	goto L438
L438:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1959)))
	v1975 = base.I32_rem_u_s(v1974, v52<<(uint(int32(3))%32))
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+int32(8)+int32(base.Ui32(v1975)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v1979)>>(uint(v1975&int32(7))%32))&int32(1) != 0 {
		goto L435
	} else {
		goto L440
	}
L439:
	;
	goto L429
L440:
	;
	v1985 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1959)+4)))
	v1993 = int32(1)
	if v1993 < v1966 {
		v1959 = v1959 + (v1985+int32(7))&int32(_a_F_ltree_consistent_1) + int32(8)
		v1966 = v1966 - v1993
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	goto L434
L443:
	;
	if v2270 == int32(0) {
		goto L429
	} else {
		goto L495
	}
L444:
	;
	v2270 = int32(1)
	goto L443
L445:
	;
	goto L446
L446:
	;
	v2055 = v53 + int32(8)
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v2057&int32(2) != 0 {
		goto L449
	} else {
		goto L450
	}
L447:
	;
	if v2137 <= int32(0) {
		goto L470
	} else {
		goto L471
	}
L448:
	;
	if base.Ui32(v2065) < base.Ui32(v2050) {
		goto L467
	} else {
		goto L468
	}
L449:
	;
	v2060 = int32(0)
	goto L451
L450:
	;
	v2060 = v52
	goto L451
L451:
	;
	v2061 = v2055 + v2060
	v2063 = v2057 & int32(1)
	if v2063 != 0 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v2064 = v2055
	goto L454
L453:
	;
	v2064 = v2061
	goto L454
L454:
	;
	v2065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2064)+4)))
	if v2065 == int32(0) {
		goto L448
	} else {
		goto L455
	}
L455:
	;
	v2072 = v1920 + int32(16)
	v2074 = v2064 + int32(8)
	v2079 = v2050
	v2080 = v2065
	goto L456
L456:
	;
	v2090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2074))))
	v2091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2072)+20)))
	if base.Ui32(v2090) < base.Ui32(v2091) {
		goto L458
	} else {
		goto L459
	}
L457:
	;
	goto L448
L458:
	;
	v2093 = v2090
	goto L460
L459:
	;
	v2093 = v2091
	goto L460
L460:
	;
	v2094 = F_memcmp(m, v2074+int32(2), v2072+int32(23), v2093)
	mBase = m.M
	if v2094 != 0 {
		v2137 = v2094
		goto L447
	} else {
		goto L461
	}
L461:
	;
	if v2090 != v2091 {
		goto L462
	} else {
		goto L463
	}
L462:
	;
	v2137 = v2090 - v2091
	goto L447
L463:
	;
	goto L464
L464:
	;
	if v2080 < int32(2) {
		goto L448
	} else {
		goto L465
	}
L465:
	;
	v2099 = int32(1)
	v2103 = int32(_a_F_ltree_consistent_1)
	v2106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2072))))
	if v2099 < v2079 {
		v2072 = v2072 + (v2106+int32(7))&v2103
		v2074 = v2074 + (v2090+int32(9))&v2103
		v2079 = v2079 - v2099
		v2080 = v2080 - v2099
		goto L456
	} else {
		goto L466
	}
L466:
	;
	goto L457
L467:
	;
	v2131 = v2065
	goto L469
L468:
	;
	v2131 = v2050
	goto L469
L469:
	;
	v2137 = v2131 - v2050
	goto L447
L470:
	;
	if v2063 != 0 {
		v2155 = v2055
		goto L473
	} else {
		goto L474
	}
L471:
	;
	v2255 = int32(0)
	goto L472
L472:
	;
	v2270 = v2255
	goto L443
L473:
	;
	v2156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2155)+4)))
	if v2156 == int32(0) {
		goto L479
	} else {
		goto L480
	}
L474:
	;
	if v2057&int32(4) != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v2155 = v2061
	goto L473
L476:
	;
	goto L477
L477:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v2061)))
	v2155 = v2061 + int32(base.Ui32(v2151)>>(uint(int32(2))%32))
	goto L473
L478:
	;
	v2255 = base.B2i32(int32(0) <= v2228)
	goto L472
L479:
	;
	if base.Ui32(v2156) < base.Ui32(v2050) {
		goto L492
	} else {
		goto L493
	}
L480:
	;
	v2163 = v1920 + int32(16)
	v2165 = v2155 + int32(8)
	v2170 = v2050
	v2171 = v2156
	goto L481
L481:
	;
	v2181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2165))))
	v2182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2163)+20)))
	if base.Ui32(v2181) < base.Ui32(v2182) {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	goto L479
L483:
	;
	v2184 = v2181
	goto L485
L484:
	;
	v2184 = v2182
	goto L485
L485:
	;
	v2185 = F_memcmp(m, v2165+int32(2), v2163+int32(23), v2184)
	mBase = m.M
	if v2185 != 0 {
		v2228 = v2185
		goto L478
	} else {
		goto L486
	}
L486:
	;
	if v2181 != v2182 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2228 = v2181 - v2182
	goto L478
L488:
	;
	goto L489
L489:
	;
	if v2171 < int32(2) {
		goto L479
	} else {
		goto L490
	}
L490:
	;
	v2190 = int32(1)
	v2194 = int32(_a_F_ltree_consistent_1)
	v2197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2163))))
	if v2190 < v2170 {
		v2163 = v2163 + (v2197+int32(7))&v2194
		v2165 = v2165 + (v2181+int32(9))&v2194
		v2170 = v2170 - v2190
		v2171 = v2171 - v2190
		goto L481
	} else {
		goto L491
	}
L491:
	;
	goto L482
L492:
	;
	v2222 = v2156
	goto L494
L493:
	;
	v2222 = v2050
	goto L494
L494:
	;
	v2228 = v2222 - v2050
	goto L478
L495:
	;
	v2431 = v1865
	v2434 = int32(1)
	goto L13
L496:
	;
	goto L428
L497:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L10
	} else {
		goto L498
	}
L498:
	;
	F_errmsg(m, int32(_a_F_ltree_consistent_6), int32(0))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L10
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_7), int32(604), int32(_a_F_ltree_consistent_8))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L10
	} else {
		goto L500
	}
L500:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L10
	} else {
		goto L502
	}
L502:
	;
	F_errmsg(m, int32(_a_F_ltree_consistent_9), int32(0))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L10
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_7), int32(600), int32(_a_F_ltree_consistent_8))
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L10
	} else {
		goto L504
	}
L504:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L505:
	;
	v2338 = int32(0)
	goto L507
L506:
	;
	v2338 = v52
	goto L507
L507:
	;
	v2339 = v53 + v2338
	v2342 = int32(0)
	v2349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	v2352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2339+int32(8))+4)))
	if base.B2i32(v2349 == v2342)|base.B2i32(v2352 == v2342) != 0 {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	v2431 = v61
	v2434 = base.B2i32(int32(0) <= v2426)
	goto L13
L509:
	;
	v2426 = v2416
	goto L508
L510:
	;
	v2416 = v2349 - v2352
	goto L509
L511:
	;
	v2360 = v61 + int32(8)
	v2361 = v2339 + int32(16)
	v2364 = v2352
	v2365 = v2349
	goto L512
L512:
	;
	v2369 = int32(2)
	v2373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2360))))
	v2374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2361))))
	if base.Ui32(v2373) < base.Ui32(v2374) {
		goto L514
	} else {
		goto L515
	}
L513:
	;
	goto L510
L514:
	;
	v2376 = v2373
	goto L516
L515:
	;
	v2376 = v2374
	goto L516
L516:
	;
	v2377 = F_memcmp(m, v2360+v2369, v2361+v2369, v2376)
	mBase = m.M
	if v2377 != 0 {
		v2416 = v2377
		goto L509
	} else {
		goto L517
	}
L517:
	;
	if v2373 != v2374 {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	v2426 = v2373 - v2374
	goto L508
L519:
	;
	goto L520
L520:
	;
	if v2365 < int32(2) {
		goto L510
	} else {
		goto L521
	}
L521:
	;
	v2382 = int32(1)
	v2384 = int32(9)
	v2386 = int32(_a_F_ltree_consistent_1)
	if v2382 < v2364 {
		v2360 = v2360 + (v2373+v2384)&v2386
		v2361 = v2361 + (v2374+v2384)&v2386
		v2364 = v2364 - v2382
		v2365 = v2365 - v2382
		goto L512
	} else {
		goto L522
	}
L522:
	;
	goto L513
L523:
	;
	F_pfree(m, v2431)
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L10
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	m.G0 = v19 + int32(16)
	return base.I64_extend_i32_u(v2434)
L526:
	;
	goto L525
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v57
	F_errmsg_internal(m, int32(_a_F_ltree_consistent_10), v19)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L10
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_ltree_consistent_7), int32(716), int32(_a_F_ltree_consistent_11))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L10
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ltree_eq(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v156 != v14 {
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v155 = base.I64_extend_i32_u(base.B2i32(v21 == v22))
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v39 = v21
	v40 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113|base.B2i32(v48 != v49) != 0 {
		v155 = int64(0)
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
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
	if v40 < int32(2) {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v118 = int32(1)
	v123 = (v48 + int32(9)) & int32(_a_F_ltree_eq_0)
	if v118 < v39 {
		v35 = v123 + v35
		v36 = v36 + v123
		v39 = v39 - v118
		v40 = v40 - v118
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
L33:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v160 != v19 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	F_pfree(m, v19)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	return v155
L40:
	;
	goto L39
}
func F_ltree_gist_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_ltree_gist_out_0), int32(36), int32(_a_F_ltree_gist_out_1), int32(_a_F_ltree_gist_out_2), int32(_a_F_ltree_gist_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ltree_lt(m *base.Module, l0 int32) int64 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v159 != v14 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v157 = v22 - v21
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v40 = v21
	v41 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113 != 0 {
		v157 = v113
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
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
	if v48 != v49 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v48 - v49
	goto L4
L32:
	;
	goto L33
L33:
	;
	if v41 < int32(2) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(1)
	v120 = int32(9)
	v122 = int32(_a_F_ltree_lt_0)
	if v118 < v40 {
		v35 = v35 + (v48+v120)&v122
		v36 = v36 + (v49+v120)&v122
		v40 = v40 - v118
		v41 = v41 - v118
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v163 != v19 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v19)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.I64_extend_i32_u(int32(base.Ui32(v157) >> (uint(int32(31)) % 32)))
L43:
	;
	goto L42
}
func F_ltree_risparent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v9 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v136 != v11 {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
	if base.Ui32(v19) < base.Ui32(v18) {
		v135 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = int32(8)
	v27 = v18
	v28 = v16 + v21
	v29 = v11 + v21
	goto L9
L7:
	;
	goto L8
L8:
	;
	v135 = int64(1)
	goto L3
L9:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	if v34 != v35 {
		v135 = v9
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v37 = int32(2)
	v38 = v29 + v37
	v40 = v28 + v37
	if base.Ui32(int32(4)) <= base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v102 != 0 {
		v135 = v9
		goto L3
	} else {
		goto L30
	}
L13:
	;
	v102 = int32(0)
	goto L12
L14:
	;
	v76 = v71
	v77 = v72
	v78 = v73
	goto L24
L15:
	;
	if (v38|v40)&int32(3) != 0 {
		v71 = v38
		v72 = v40
		v73 = v34
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v64 = v38
	v65 = v40
	v66 = v34
	goto L17
L17:
	;
	if v66 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v48 = v38
	v49 = v40
	v50 = v34
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v53 != v54 {
		v71 = v48
		v72 = v49
		v73 = v50
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v64 = v59
	v65 = v57
	v66 = v61
	goto L17
L21:
	;
	v56 = int32(4)
	v57 = v49 + v56
	v59 = v48 + v56
	v61 = v50 - v56
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		v48 = v59
		v49 = v57
		v50 = v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v71 = v64
	v72 = v65
	v73 = v66
	goto L14
L24:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 == v82 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v102 = v81 - v82
	goto L12
L26:
	;
	v84 = int32(1)
	v89 = v78 - v84
	if v89 != 0 {
		v76 = v76 + v84
		v77 = v77 + v84
		v78 = v89
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
	v103 = int32(9)
	v105 = int32(_a_F_ltree_risparent_0)
	v113 = int32(1)
	if v113 < v27 {
		v27 = v27 - v113
		v28 = v28 + (v35+v103)&v105
		v29 = v29 + (v34+v103)&v105
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	F_pfree(m, v11)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v140 != v16 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v16)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	return v135
L39:
	;
	goto L38
}
