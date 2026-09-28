package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_width_bucket_numeric(m *base.Module, l0 int32) int64 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v435 int64
	_ = v435
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int64
	_ = v772
	var v775 int64
	_ = v775
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v934 int32
	_ = v934
	var v943 int32
	_ = v943
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int64
	_ = v959
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v25 = F_pg_detoast_datum(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				if int32(0) < base.I32_wrap_i64(v27) {
					v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
					if base.Ui32(v31) <= base.Ui32(int32(_a_F_width_bucket_numeric_0)) {
						v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
						if base.Ui32(int32(_a_F_width_bucket_numeric_0)) < base.Ui32(v34) {
							v40 = v34
							if v40&int32(_a_F_width_bucket_numeric_1) == int32(_a_F_width_bucket_numeric_2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v993 = m.ExcPending
								if v993 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v996 = m.ExcPending
									if v996 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
										mBase = m.M
										v1000 = m.ExcPending
										if v1000 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1978), int32(_a_F_width_bucket_numeric_5))
											mBase = m.M
											v1005 = m.ExcPending
											if v1005 != 0 {
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
								v45 = int32(_a_F_width_bucket_numeric_6)
								v47 = int32(_a_F_width_bucket_numeric_7)
								if base.B2i32(v31&v45 == v47)|base.B2i32(v40&v45 == v47) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1009 = m.ExcPending
									if v1009 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v1012 = m.ExcPending
										if v1012 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_8), int32(0))
											mBase = m.M
											v1016 = m.ExcPending
											if v1016 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1983), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v1021 = m.ExcPending
												if v1021 != 0 {
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
									v55 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v55
									*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v55
									*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v55
									v62 = F_palloc(m, int32(12))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v62
										v65 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
										*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
										v74 = v65
										v81 = v62 + int32(12)
										v83 = v27 & int64(2147483647)
										for {
											v86 = v81 - int32(2)
											v88 = base.I64_div_u_s(v83, int64(10000))
											v91 = v88*int64(55536) + v83
											*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v91)
											v94 = v74 + int32(1)
											if base.Ui64(int64(9999)) < base.Ui64(v83) {
												v74 = v94
												v81 = v86
												v83 = v88
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94
										*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v86
										v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
										v112 = base.I32_extend16_s(v111)
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v113) {
											if v113 != int32(_a_F_width_bucket_numeric_7) {
												if v113 != int32(_a_F_width_bucket_numeric_2) {
													if v112 != int32(-4096) {
														v132 = int32(-1)
													} else {
														v132 = int32(0)
													}
													v253 = v132
												} else {
													v253 = base.B2i32(v112 != int32(-16384))
												}
											} else {
												if v112 == int32(-16384) {
													v127 = int32(-1)
												} else {
													v127 = base.B2i32(v112 != int32(-12288))
												}
												v253 = v127
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v112) {
												if v112 == int32(-4096) {
													v139 = int32(1)
												} else {
													v139 = int32(-1)
												}
												v253 = v139
											} else {
												v141 = v22 + int32(6)
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												v147 = base.I32_extend16_s(v113)
												v149 = base.B2i32(int32(0) <= v147)
												if int32(0) <= v147 {
													v150 = int32(-8)
												} else {
													v150 = int32(-6)
												}
												if int32(0) <= v147 {
													v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141))))
													v162 = v152
												} else {
													v162 = v113<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v113&int32(63)
												}
												v164 = int32(base.Ui32(int32(base.Ui32(v142)>>(uint(int32(2))%32))+v150) >> (uint(int32(1)) % 32))
												v166 = v25 + int32(6)
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
												v173 = base.B2i32(int32(0) <= v112)
												if int32(0) <= v112 {
													v174 = int32(-8)
												} else {
													v174 = int32(-6)
												}
												if int32(0) <= v112 {
													v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166))))
													v186 = v176
												} else {
													v186 = v111<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v111&int32(63)
												}
												v187 = int32(1)
												v188 = int32(base.Ui32(int32(base.Ui32(v167)>>(uint(int32(2))%32))+v174) >> (uint(v187) % 32))
												v194 = v111 & int32(_a_F_width_bucket_numeric_2)
												if v194 == int32(_a_F_width_bucket_numeric_9) {
													v197 = v111 << (uint(v187) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v197 = v194
												}
												if v164 == int32(0) {
													if v188 == int32(0) {
														v253 = int32(0)
													} else {
														if v197 == int32(_a_F_width_bucket_numeric_10) {
															v207 = int32(1)
														} else {
															v207 = int32(-1)
														}
														v253 = v207
													}
												} else {
													v213 = v113 & int32(_a_F_width_bucket_numeric_2)
													if v213 == int32(_a_F_width_bucket_numeric_9) {
														v216 = v113 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v216 = v213
													}
													if v188 == int32(0) {
														if v216 != 0 {
															v221 = int32(-1)
														} else {
															v221 = int32(1)
														}
														v253 = v221
													} else {
														if v147 < int32(0) {
															v226 = v141
														} else {
															v226 = v22 + int32(8)
														}
														if v112 < int32(0) {
															v231 = v166
														} else {
															v231 = v25 + int32(8)
														}
														if v216 == int32(0) {
															if v197 == int32(_a_F_width_bucket_numeric_10) {
																v253 = int32(1)
															} else {
																v237 = F_cmp_abs_common(m, v226, v164, v162, v231, v188, v186)
																mBase = m.M
																v253 = v237
															}
														} else {
															if v197 == int32(0) {
																v253 = int32(-1)
															} else {
																v241 = F_cmp_abs_common(m, v231, v188, v186, v226, v164, v162)
																mBase = m.M
																v253 = v241
															}
														}
													}
												}
											}
										}
										switch v253 {
										case 0:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(386138242))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_width_bucket_numeric_11), int32(0))
													mBase = m.M
													v264 = m.ExcPending
													if v264 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1997), int32(_a_F_width_bucket_numeric_5))
														mBase = m.M
														v269 = m.ExcPending
														if v269 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										case 1:
											v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v622 = base.I32_extend16_s(v621)
											v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v623) {
												if v623 != int32(_a_F_width_bucket_numeric_7) {
													if v623 != int32(_a_F_width_bucket_numeric_2) {
														if v622 != int32(-4096) {
															v642 = int32(-1)
														} else {
															v642 = int32(0)
														}
														v763 = v642
													} else {
														v763 = base.B2i32(v622 != int32(-16384))
													}
												} else {
													if v622 == int32(-16384) {
														v637 = int32(-1)
													} else {
														v637 = base.B2i32(v622 != int32(-12288))
													}
													v763 = v637
												}
											} else {
												if base.Ui32(int32(-16384)) <= base.Ui32(v622) {
													if v622 == int32(-4096) {
														v649 = int32(1)
													} else {
														v649 = int32(-1)
													}
													v763 = v649
												} else {
													v651 = v17 + int32(6)
													v652 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v657 = base.I32_extend16_s(v623)
													v659 = base.B2i32(int32(0) <= v657)
													if int32(0) <= v657 {
														v660 = int32(-8)
													} else {
														v660 = int32(-6)
													}
													if int32(0) <= v657 {
														v662 = int32(*(*int16)(unsafe.Add(mBase, uint32(v651))))
														v672 = v662
													} else {
														v672 = v623<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v623&int32(63)
													}
													v674 = int32(base.Ui32(int32(base.Ui32(v652)>>(uint(int32(2))%32))+v660) >> (uint(int32(1)) % 32))
													v676 = v22 + int32(6)
													v677 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v683 = base.B2i32(int32(0) <= v622)
													if int32(0) <= v622 {
														v684 = int32(-8)
													} else {
														v684 = int32(-6)
													}
													if int32(0) <= v622 {
														v686 = int32(*(*int16)(unsafe.Add(mBase, uint32(v676))))
														v696 = v686
													} else {
														v696 = v621<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v621&int32(63)
													}
													v697 = int32(1)
													v698 = int32(base.Ui32(int32(base.Ui32(v677)>>(uint(int32(2))%32))+v684) >> (uint(v697) % 32))
													v704 = v621 & int32(_a_F_width_bucket_numeric_2)
													if v704 == int32(_a_F_width_bucket_numeric_9) {
														v707 = v621 << (uint(v697) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v707 = v704
													}
													if v674 == int32(0) {
														if v698 == int32(0) {
															v763 = int32(0)
														} else {
															if v707 == int32(_a_F_width_bucket_numeric_10) {
																v717 = int32(1)
															} else {
																v717 = int32(-1)
															}
															v763 = v717
														}
													} else {
														v723 = v623 & int32(_a_F_width_bucket_numeric_2)
														if v723 == int32(_a_F_width_bucket_numeric_9) {
															v726 = v623 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v726 = v723
														}
														if v698 == int32(0) {
															if v726 != 0 {
																v731 = int32(-1)
															} else {
																v731 = int32(1)
															}
															v763 = v731
														} else {
															if v657 < int32(0) {
																v736 = v651
															} else {
																v736 = v17 + int32(8)
															}
															if v622 < int32(0) {
																v741 = v676
															} else {
																v741 = v22 + int32(8)
															}
															if v726 == int32(0) {
																if v707 == int32(_a_F_width_bucket_numeric_10) {
																	v763 = int32(1)
																} else {
																	v747 = F_cmp_abs_common(m, v736, v674, v672, v741, v698, v696)
																	mBase = m.M
																	v763 = v747
																}
															} else {
																if v707 == int32(0) {
																	v763 = int32(-1)
																} else {
																	v751 = F_cmp_abs_common(m, v741, v698, v696, v736, v674, v672)
																	mBase = m.M
																	v763 = v751
																}
															}
														}
													}
												}
											}
											if int32(0) < v763 {
												v767 = F_palloc(m, int32(2))
												mBase = m.M
												v768 = m.ExcPending
												if v768 != 0 {
													return int64(0)
												} else {
													v769 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(v767))) = uint16(v769)
													v772 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v772
													v775 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v775
													*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v767 + int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v767
													v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v956 = m.ExcPending
													if v956 != 0 {
														return int64(0)
													} else {
														if v955 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																F_pfree(m, v62)
																mBase = m.M
																v965 = m.ExcPending
																if v965 != 0 {
																	return int64(0)
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v966 != 0 {
																		F_pfree(m, v966)
																		mBase = m.M
																		v968 = m.ExcPending
																		if v968 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																}
															}
														}
													}
												}
											} else {
												v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v793 = base.I32_extend16_s(v792)
												v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v794) {
													if v794 != int32(_a_F_width_bucket_numeric_7) {
														if v794 != int32(_a_F_width_bucket_numeric_2) {
															if v793 != int32(-4096) {
																v813 = int32(-1)
															} else {
																v813 = int32(0)
															}
															v934 = v813
														} else {
															v934 = base.B2i32(v793 != int32(-16384))
														}
													} else {
														if v793 == int32(-16384) {
															v808 = int32(-1)
														} else {
															v808 = base.B2i32(v793 != int32(-12288))
														}
														v934 = v808
													}
												} else {
													if base.Ui32(int32(-16384)) <= base.Ui32(v793) {
														if v793 == int32(-4096) {
															v820 = int32(1)
														} else {
															v820 = int32(-1)
														}
														v934 = v820
													} else {
														v822 = v17 + int32(6)
														v823 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v828 = base.I32_extend16_s(v794)
														v830 = base.B2i32(int32(0) <= v828)
														if int32(0) <= v828 {
															v831 = int32(-8)
														} else {
															v831 = int32(-6)
														}
														if int32(0) <= v828 {
															v833 = int32(*(*int16)(unsafe.Add(mBase, uint32(v822))))
															v843 = v833
														} else {
															v843 = v794<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v794&int32(63)
														}
														v845 = int32(base.Ui32(int32(base.Ui32(v823)>>(uint(int32(2))%32))+v831) >> (uint(int32(1)) % 32))
														v847 = v25 + int32(6)
														v848 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														v854 = base.B2i32(int32(0) <= v793)
														if int32(0) <= v793 {
															v855 = int32(-8)
														} else {
															v855 = int32(-6)
														}
														if int32(0) <= v793 {
															v857 = int32(*(*int16)(unsafe.Add(mBase, uint32(v847))))
															v867 = v857
														} else {
															v867 = v792<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v792&int32(63)
														}
														v868 = int32(1)
														v869 = int32(base.Ui32(int32(base.Ui32(v848)>>(uint(int32(2))%32))+v855) >> (uint(v868) % 32))
														v875 = v792 & int32(_a_F_width_bucket_numeric_2)
														if v875 == int32(_a_F_width_bucket_numeric_9) {
															v878 = v792 << (uint(v868) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v878 = v875
														}
														if v845 == int32(0) {
															if v869 == int32(0) {
																v934 = int32(0)
															} else {
																if v878 == int32(_a_F_width_bucket_numeric_10) {
																	v888 = int32(1)
																} else {
																	v888 = int32(-1)
																}
																v934 = v888
															}
														} else {
															v894 = v794 & int32(_a_F_width_bucket_numeric_2)
															if v894 == int32(_a_F_width_bucket_numeric_9) {
																v897 = v794 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
															} else {
																v897 = v894
															}
															if v869 == int32(0) {
																if v897 != 0 {
																	v902 = int32(-1)
																} else {
																	v902 = int32(1)
																}
																v934 = v902
															} else {
																if v828 < int32(0) {
																	v907 = v822
																} else {
																	v907 = v17 + int32(8)
																}
																if v793 < int32(0) {
																	v912 = v847
																} else {
																	v912 = v25 + int32(8)
																}
																if v897 == int32(0) {
																	if v878 == int32(_a_F_width_bucket_numeric_10) {
																		v934 = int32(1)
																	} else {
																		v918 = F_cmp_abs_common(m, v907, v845, v843, v912, v869, v867)
																		mBase = m.M
																		v934 = v918
																	}
																} else {
																	if v878 == int32(0) {
																		v934 = int32(-1)
																	} else {
																		v922 = F_cmp_abs_common(m, v912, v869, v867, v907, v845, v843)
																		mBase = m.M
																		v934 = v922
																	}
																}
															}
														}
													}
												}
												if v934 <= int32(0) {
													F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
													mBase = m.M
													v943 = m.ExcPending
													if v943 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												} else {
													F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
													mBase = m.M
													v949 = m.ExcPending
													if v949 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												}
											}
										default:
											v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v282 = base.I32_extend16_s(v281)
											v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v283) {
												if v283 != int32(_a_F_width_bucket_numeric_7) {
													if v283 != int32(_a_F_width_bucket_numeric_2) {
														if v282 != int32(-4096) {
															v302 = int32(-1)
														} else {
															v302 = int32(0)
														}
														v423 = v302
													} else {
														v423 = base.B2i32(v282 != int32(-16384))
													}
												} else {
													if v282 == int32(-16384) {
														v297 = int32(-1)
													} else {
														v297 = base.B2i32(v282 != int32(-12288))
													}
													v423 = v297
												}
											} else {
												if base.Ui32(int32(-16384)) <= base.Ui32(v282) {
													if v282 == int32(-4096) {
														v309 = int32(1)
													} else {
														v309 = int32(-1)
													}
													v423 = v309
												} else {
													v311 = v17 + int32(6)
													v312 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v317 = base.I32_extend16_s(v283)
													v319 = base.B2i32(int32(0) <= v317)
													if int32(0) <= v317 {
														v320 = int32(-8)
													} else {
														v320 = int32(-6)
													}
													if int32(0) <= v317 {
														v322 = int32(*(*int16)(unsafe.Add(mBase, uint32(v311))))
														v332 = v322
													} else {
														v332 = v283<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v283&int32(63)
													}
													v334 = int32(base.Ui32(int32(base.Ui32(v312)>>(uint(int32(2))%32))+v320) >> (uint(int32(1)) % 32))
													v336 = v22 + int32(6)
													v337 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v343 = base.B2i32(int32(0) <= v282)
													if int32(0) <= v282 {
														v344 = int32(-8)
													} else {
														v344 = int32(-6)
													}
													if int32(0) <= v282 {
														v346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336))))
														v356 = v346
													} else {
														v356 = v281<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v281&int32(63)
													}
													v357 = int32(1)
													v358 = int32(base.Ui32(int32(base.Ui32(v337)>>(uint(int32(2))%32))+v344) >> (uint(v357) % 32))
													v364 = v281 & int32(_a_F_width_bucket_numeric_2)
													if v364 == int32(_a_F_width_bucket_numeric_9) {
														v367 = v281 << (uint(v357) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v367 = v364
													}
													if v334 == int32(0) {
														if v358 == int32(0) {
															v423 = int32(0)
														} else {
															if v367 == int32(_a_F_width_bucket_numeric_10) {
																v377 = int32(1)
															} else {
																v377 = int32(-1)
															}
															v423 = v377
														}
													} else {
														v383 = v283 & int32(_a_F_width_bucket_numeric_2)
														if v383 == int32(_a_F_width_bucket_numeric_9) {
															v386 = v283 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v386 = v383
														}
														if v358 == int32(0) {
															if v386 != 0 {
																v391 = int32(-1)
															} else {
																v391 = int32(1)
															}
															v423 = v391
														} else {
															if v317 < int32(0) {
																v396 = v311
															} else {
																v396 = v17 + int32(8)
															}
															if v282 < int32(0) {
																v401 = v336
															} else {
																v401 = v22 + int32(8)
															}
															if v386 == int32(0) {
																if v367 == int32(_a_F_width_bucket_numeric_10) {
																	v423 = int32(1)
																} else {
																	v407 = F_cmp_abs_common(m, v396, v334, v332, v401, v358, v356)
																	mBase = m.M
																	v423 = v407
																}
															} else {
																if v367 == int32(0) {
																	v423 = int32(-1)
																} else {
																	v411 = F_cmp_abs_common(m, v401, v358, v356, v396, v334, v332)
																	mBase = m.M
																	v423 = v411
																}
															}
														}
													}
												}
											}
											if v423 < int32(0) {
												v427 = F_palloc(m, int32(2))
												mBase = m.M
												v428 = m.ExcPending
												if v428 != 0 {
													return int64(0)
												} else {
													v429 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
													v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
													v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
													*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
													v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v956 = m.ExcPending
													if v956 != 0 {
														return int64(0)
													} else {
														if v955 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																F_pfree(m, v62)
																mBase = m.M
																v965 = m.ExcPending
																if v965 != 0 {
																	return int64(0)
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v966 != 0 {
																		F_pfree(m, v966)
																		mBase = m.M
																		v968 = m.ExcPending
																		if v968 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																}
															}
														}
													}
												}
											} else {
												v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v453 = base.I32_extend16_s(v452)
												v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v454) {
													if v454 != int32(_a_F_width_bucket_numeric_7) {
														if v454 != int32(_a_F_width_bucket_numeric_2) {
															if v453 != int32(-4096) {
																v473 = int32(-1)
															} else {
																v473 = int32(0)
															}
															v594 = v473
														} else {
															v594 = base.B2i32(v453 != int32(-16384))
														}
													} else {
														if v453 == int32(-16384) {
															v468 = int32(-1)
														} else {
															v468 = base.B2i32(v453 != int32(-12288))
														}
														v594 = v468
													}
												} else {
													if base.Ui32(int32(-16384)) <= base.Ui32(v453) {
														if v453 == int32(-4096) {
															v480 = int32(1)
														} else {
															v480 = int32(-1)
														}
														v594 = v480
													} else {
														v482 = v17 + int32(6)
														v483 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v488 = base.I32_extend16_s(v454)
														v490 = base.B2i32(int32(0) <= v488)
														if int32(0) <= v488 {
															v491 = int32(-8)
														} else {
															v491 = int32(-6)
														}
														if int32(0) <= v488 {
															v493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v482))))
															v503 = v493
														} else {
															v503 = v454<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v454&int32(63)
														}
														v505 = int32(base.Ui32(int32(base.Ui32(v483)>>(uint(int32(2))%32))+v491) >> (uint(int32(1)) % 32))
														v507 = v25 + int32(6)
														v508 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														v514 = base.B2i32(int32(0) <= v453)
														if int32(0) <= v453 {
															v515 = int32(-8)
														} else {
															v515 = int32(-6)
														}
														if int32(0) <= v453 {
															v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507))))
															v527 = v517
														} else {
															v527 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
														}
														v528 = int32(1)
														v529 = int32(base.Ui32(int32(base.Ui32(v508)>>(uint(int32(2))%32))+v515) >> (uint(v528) % 32))
														v535 = v452 & int32(_a_F_width_bucket_numeric_2)
														if v535 == int32(_a_F_width_bucket_numeric_9) {
															v538 = v452 << (uint(v528) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v538 = v535
														}
														if v505 == int32(0) {
															if v529 == int32(0) {
																v594 = int32(0)
															} else {
																if v538 == int32(_a_F_width_bucket_numeric_10) {
																	v548 = int32(1)
																} else {
																	v548 = int32(-1)
																}
																v594 = v548
															}
														} else {
															v554 = v454 & int32(_a_F_width_bucket_numeric_2)
															if v554 == int32(_a_F_width_bucket_numeric_9) {
																v557 = v454 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
															} else {
																v557 = v554
															}
															if v529 == int32(0) {
																if v557 != 0 {
																	v562 = int32(-1)
																} else {
																	v562 = int32(1)
																}
																v594 = v562
															} else {
																if v488 < int32(0) {
																	v567 = v482
																} else {
																	v567 = v17 + int32(8)
																}
																if v453 < int32(0) {
																	v572 = v507
																} else {
																	v572 = v25 + int32(8)
																}
																if v557 == int32(0) {
																	if v538 == int32(_a_F_width_bucket_numeric_10) {
																		v594 = int32(1)
																	} else {
																		v578 = F_cmp_abs_common(m, v567, v505, v503, v572, v529, v527)
																		mBase = m.M
																		v594 = v578
																	}
																} else {
																	if v538 == int32(0) {
																		v594 = int32(-1)
																	} else {
																		v582 = F_cmp_abs_common(m, v572, v529, v527, v567, v505, v503)
																		mBase = m.M
																		v594 = v582
																	}
																}
															}
														}
													}
												}
												if int32(0) <= v594 {
													F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
													mBase = m.M
													v603 = m.ExcPending
													if v603 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												} else {
													F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
													mBase = m.M
													v609 = m.ExcPending
													if v609 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
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
							v55 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v55
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v55
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v55
							v62 = F_palloc(m, int32(12))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v62
								v65 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
								*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
								v74 = v65
								v81 = v62 + int32(12)
								v83 = v27 & int64(2147483647)
								for {
									v86 = v81 - int32(2)
									v88 = base.I64_div_u_s(v83, int64(10000))
									v91 = v88*int64(55536) + v83
									*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v91)
									v94 = v74 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v83) {
										v74 = v94
										v81 = v86
										v83 = v88
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v74
								*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94
								*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v86
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
								v112 = base.I32_extend16_s(v111)
								v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
								if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v113) {
									if v113 != int32(_a_F_width_bucket_numeric_7) {
										if v113 != int32(_a_F_width_bucket_numeric_2) {
											if v112 != int32(-4096) {
												v132 = int32(-1)
											} else {
												v132 = int32(0)
											}
											v253 = v132
										} else {
											v253 = base.B2i32(v112 != int32(-16384))
										}
									} else {
										if v112 == int32(-16384) {
											v127 = int32(-1)
										} else {
											v127 = base.B2i32(v112 != int32(-12288))
										}
										v253 = v127
									}
								} else {
									if base.Ui32(int32(-16384)) <= base.Ui32(v112) {
										if v112 == int32(-4096) {
											v139 = int32(1)
										} else {
											v139 = int32(-1)
										}
										v253 = v139
									} else {
										v141 = v22 + int32(6)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
										v147 = base.I32_extend16_s(v113)
										v149 = base.B2i32(int32(0) <= v147)
										if int32(0) <= v147 {
											v150 = int32(-8)
										} else {
											v150 = int32(-6)
										}
										if int32(0) <= v147 {
											v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141))))
											v162 = v152
										} else {
											v162 = v113<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v113&int32(63)
										}
										v164 = int32(base.Ui32(int32(base.Ui32(v142)>>(uint(int32(2))%32))+v150) >> (uint(int32(1)) % 32))
										v166 = v25 + int32(6)
										v167 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
										v173 = base.B2i32(int32(0) <= v112)
										if int32(0) <= v112 {
											v174 = int32(-8)
										} else {
											v174 = int32(-6)
										}
										if int32(0) <= v112 {
											v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166))))
											v186 = v176
										} else {
											v186 = v111<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v111&int32(63)
										}
										v187 = int32(1)
										v188 = int32(base.Ui32(int32(base.Ui32(v167)>>(uint(int32(2))%32))+v174) >> (uint(v187) % 32))
										v194 = v111 & int32(_a_F_width_bucket_numeric_2)
										if v194 == int32(_a_F_width_bucket_numeric_9) {
											v197 = v111 << (uint(v187) % 32) & int32(_a_F_width_bucket_numeric_10)
										} else {
											v197 = v194
										}
										if v164 == int32(0) {
											if v188 == int32(0) {
												v253 = int32(0)
											} else {
												if v197 == int32(_a_F_width_bucket_numeric_10) {
													v207 = int32(1)
												} else {
													v207 = int32(-1)
												}
												v253 = v207
											}
										} else {
											v213 = v113 & int32(_a_F_width_bucket_numeric_2)
											if v213 == int32(_a_F_width_bucket_numeric_9) {
												v216 = v113 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
											} else {
												v216 = v213
											}
											if v188 == int32(0) {
												if v216 != 0 {
													v221 = int32(-1)
												} else {
													v221 = int32(1)
												}
												v253 = v221
											} else {
												if v147 < int32(0) {
													v226 = v141
												} else {
													v226 = v22 + int32(8)
												}
												if v112 < int32(0) {
													v231 = v166
												} else {
													v231 = v25 + int32(8)
												}
												if v216 == int32(0) {
													if v197 == int32(_a_F_width_bucket_numeric_10) {
														v253 = int32(1)
													} else {
														v237 = F_cmp_abs_common(m, v226, v164, v162, v231, v188, v186)
														mBase = m.M
														v253 = v237
													}
												} else {
													if v197 == int32(0) {
														v253 = int32(-1)
													} else {
														v241 = F_cmp_abs_common(m, v231, v188, v186, v226, v164, v162)
														mBase = m.M
														v253 = v241
													}
												}
											}
										}
									}
								}
								switch v253 {
								case 0:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v260 = m.ExcPending
										if v260 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_11), int32(0))
											mBase = m.M
											v264 = m.ExcPending
											if v264 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1997), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v269 = m.ExcPending
												if v269 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								case 1:
									v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
									v622 = base.I32_extend16_s(v621)
									v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
									if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v623) {
										if v623 != int32(_a_F_width_bucket_numeric_7) {
											if v623 != int32(_a_F_width_bucket_numeric_2) {
												if v622 != int32(-4096) {
													v642 = int32(-1)
												} else {
													v642 = int32(0)
												}
												v763 = v642
											} else {
												v763 = base.B2i32(v622 != int32(-16384))
											}
										} else {
											if v622 == int32(-16384) {
												v637 = int32(-1)
											} else {
												v637 = base.B2i32(v622 != int32(-12288))
											}
											v763 = v637
										}
									} else {
										if base.Ui32(int32(-16384)) <= base.Ui32(v622) {
											if v622 == int32(-4096) {
												v649 = int32(1)
											} else {
												v649 = int32(-1)
											}
											v763 = v649
										} else {
											v651 = v17 + int32(6)
											v652 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
											v657 = base.I32_extend16_s(v623)
											v659 = base.B2i32(int32(0) <= v657)
											if int32(0) <= v657 {
												v660 = int32(-8)
											} else {
												v660 = int32(-6)
											}
											if int32(0) <= v657 {
												v662 = int32(*(*int16)(unsafe.Add(mBase, uint32(v651))))
												v672 = v662
											} else {
												v672 = v623<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v623&int32(63)
											}
											v674 = int32(base.Ui32(int32(base.Ui32(v652)>>(uint(int32(2))%32))+v660) >> (uint(int32(1)) % 32))
											v676 = v22 + int32(6)
											v677 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
											v683 = base.B2i32(int32(0) <= v622)
											if int32(0) <= v622 {
												v684 = int32(-8)
											} else {
												v684 = int32(-6)
											}
											if int32(0) <= v622 {
												v686 = int32(*(*int16)(unsafe.Add(mBase, uint32(v676))))
												v696 = v686
											} else {
												v696 = v621<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v621&int32(63)
											}
											v697 = int32(1)
											v698 = int32(base.Ui32(int32(base.Ui32(v677)>>(uint(int32(2))%32))+v684) >> (uint(v697) % 32))
											v704 = v621 & int32(_a_F_width_bucket_numeric_2)
											if v704 == int32(_a_F_width_bucket_numeric_9) {
												v707 = v621 << (uint(v697) % 32) & int32(_a_F_width_bucket_numeric_10)
											} else {
												v707 = v704
											}
											if v674 == int32(0) {
												if v698 == int32(0) {
													v763 = int32(0)
												} else {
													if v707 == int32(_a_F_width_bucket_numeric_10) {
														v717 = int32(1)
													} else {
														v717 = int32(-1)
													}
													v763 = v717
												}
											} else {
												v723 = v623 & int32(_a_F_width_bucket_numeric_2)
												if v723 == int32(_a_F_width_bucket_numeric_9) {
													v726 = v623 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v726 = v723
												}
												if v698 == int32(0) {
													if v726 != 0 {
														v731 = int32(-1)
													} else {
														v731 = int32(1)
													}
													v763 = v731
												} else {
													if v657 < int32(0) {
														v736 = v651
													} else {
														v736 = v17 + int32(8)
													}
													if v622 < int32(0) {
														v741 = v676
													} else {
														v741 = v22 + int32(8)
													}
													if v726 == int32(0) {
														if v707 == int32(_a_F_width_bucket_numeric_10) {
															v763 = int32(1)
														} else {
															v747 = F_cmp_abs_common(m, v736, v674, v672, v741, v698, v696)
															mBase = m.M
															v763 = v747
														}
													} else {
														if v707 == int32(0) {
															v763 = int32(-1)
														} else {
															v751 = F_cmp_abs_common(m, v741, v698, v696, v736, v674, v672)
															mBase = m.M
															v763 = v751
														}
													}
												}
											}
										}
									}
									if int32(0) < v763 {
										v767 = F_palloc(m, int32(2))
										mBase = m.M
										v768 = m.ExcPending
										if v768 != 0 {
											return int64(0)
										} else {
											v769 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v767))) = uint16(v769)
											v772 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
											*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v772
											v775 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
											*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v775
											*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v767 + int32(2)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v767
											v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
											mBase = m.M
											v956 = m.ExcPending
											if v956 != 0 {
												return int64(0)
											} else {
												if v955 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v1026 = m.ExcPending
													if v1026 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v1029 = m.ExcPending
														if v1029 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																mBase = m.M
																v1038 = m.ExcPending
																if v1038 != 0 {
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
													v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
													if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														F_pfree(m, v62)
														mBase = m.M
														v965 = m.ExcPending
														if v965 != 0 {
															return int64(0)
														} else {
															v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
															if v966 != 0 {
																F_pfree(m, v966)
																mBase = m.M
																v968 = m.ExcPending
																if v968 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
																}
															} else {
																m.G0 = v14 - int32(-64)
																return v959
															}
														}
													}
												}
											}
										}
									} else {
										v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
										v793 = base.I32_extend16_s(v792)
										v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v794) {
											if v794 != int32(_a_F_width_bucket_numeric_7) {
												if v794 != int32(_a_F_width_bucket_numeric_2) {
													if v793 != int32(-4096) {
														v813 = int32(-1)
													} else {
														v813 = int32(0)
													}
													v934 = v813
												} else {
													v934 = base.B2i32(v793 != int32(-16384))
												}
											} else {
												if v793 == int32(-16384) {
													v808 = int32(-1)
												} else {
													v808 = base.B2i32(v793 != int32(-12288))
												}
												v934 = v808
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v793) {
												if v793 == int32(-4096) {
													v820 = int32(1)
												} else {
													v820 = int32(-1)
												}
												v934 = v820
											} else {
												v822 = v17 + int32(6)
												v823 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v828 = base.I32_extend16_s(v794)
												v830 = base.B2i32(int32(0) <= v828)
												if int32(0) <= v828 {
													v831 = int32(-8)
												} else {
													v831 = int32(-6)
												}
												if int32(0) <= v828 {
													v833 = int32(*(*int16)(unsafe.Add(mBase, uint32(v822))))
													v843 = v833
												} else {
													v843 = v794<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v794&int32(63)
												}
												v845 = int32(base.Ui32(int32(base.Ui32(v823)>>(uint(int32(2))%32))+v831) >> (uint(int32(1)) % 32))
												v847 = v25 + int32(6)
												v848 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
												v854 = base.B2i32(int32(0) <= v793)
												if int32(0) <= v793 {
													v855 = int32(-8)
												} else {
													v855 = int32(-6)
												}
												if int32(0) <= v793 {
													v857 = int32(*(*int16)(unsafe.Add(mBase, uint32(v847))))
													v867 = v857
												} else {
													v867 = v792<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v792&int32(63)
												}
												v868 = int32(1)
												v869 = int32(base.Ui32(int32(base.Ui32(v848)>>(uint(int32(2))%32))+v855) >> (uint(v868) % 32))
												v875 = v792 & int32(_a_F_width_bucket_numeric_2)
												if v875 == int32(_a_F_width_bucket_numeric_9) {
													v878 = v792 << (uint(v868) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v878 = v875
												}
												if v845 == int32(0) {
													if v869 == int32(0) {
														v934 = int32(0)
													} else {
														if v878 == int32(_a_F_width_bucket_numeric_10) {
															v888 = int32(1)
														} else {
															v888 = int32(-1)
														}
														v934 = v888
													}
												} else {
													v894 = v794 & int32(_a_F_width_bucket_numeric_2)
													if v894 == int32(_a_F_width_bucket_numeric_9) {
														v897 = v794 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v897 = v894
													}
													if v869 == int32(0) {
														if v897 != 0 {
															v902 = int32(-1)
														} else {
															v902 = int32(1)
														}
														v934 = v902
													} else {
														if v828 < int32(0) {
															v907 = v822
														} else {
															v907 = v17 + int32(8)
														}
														if v793 < int32(0) {
															v912 = v847
														} else {
															v912 = v25 + int32(8)
														}
														if v897 == int32(0) {
															if v878 == int32(_a_F_width_bucket_numeric_10) {
																v934 = int32(1)
															} else {
																v918 = F_cmp_abs_common(m, v907, v845, v843, v912, v869, v867)
																mBase = m.M
																v934 = v918
															}
														} else {
															if v878 == int32(0) {
																v934 = int32(-1)
															} else {
																v922 = F_cmp_abs_common(m, v912, v869, v867, v907, v845, v843)
																mBase = m.M
																v934 = v922
															}
														}
													}
												}
											}
										}
										if v934 <= int32(0) {
											F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
											mBase = m.M
											v943 = m.ExcPending
											if v943 != 0 {
												return int64(0)
											} else {
												v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v956 = m.ExcPending
												if v956 != 0 {
													return int64(0)
												} else {
													if v955 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															F_pfree(m, v62)
															mBase = m.M
															v965 = m.ExcPending
															if v965 != 0 {
																return int64(0)
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v966 != 0 {
																	F_pfree(m, v966)
																	mBase = m.M
																	v968 = m.ExcPending
																	if v968 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
																}
															}
														}
													}
												}
											}
										} else {
											F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
											mBase = m.M
											v949 = m.ExcPending
											if v949 != 0 {
												return int64(0)
											} else {
												v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v956 = m.ExcPending
												if v956 != 0 {
													return int64(0)
												} else {
													if v955 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															F_pfree(m, v62)
															mBase = m.M
															v965 = m.ExcPending
															if v965 != 0 {
																return int64(0)
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v966 != 0 {
																	F_pfree(m, v966)
																	mBase = m.M
																	v968 = m.ExcPending
																	if v968 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
																}
															}
														}
													}
												}
											}
										}
									}
								default:
									v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
									v282 = base.I32_extend16_s(v281)
									v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
									if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v283) {
										if v283 != int32(_a_F_width_bucket_numeric_7) {
											if v283 != int32(_a_F_width_bucket_numeric_2) {
												if v282 != int32(-4096) {
													v302 = int32(-1)
												} else {
													v302 = int32(0)
												}
												v423 = v302
											} else {
												v423 = base.B2i32(v282 != int32(-16384))
											}
										} else {
											if v282 == int32(-16384) {
												v297 = int32(-1)
											} else {
												v297 = base.B2i32(v282 != int32(-12288))
											}
											v423 = v297
										}
									} else {
										if base.Ui32(int32(-16384)) <= base.Ui32(v282) {
											if v282 == int32(-4096) {
												v309 = int32(1)
											} else {
												v309 = int32(-1)
											}
											v423 = v309
										} else {
											v311 = v17 + int32(6)
											v312 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
											v317 = base.I32_extend16_s(v283)
											v319 = base.B2i32(int32(0) <= v317)
											if int32(0) <= v317 {
												v320 = int32(-8)
											} else {
												v320 = int32(-6)
											}
											if int32(0) <= v317 {
												v322 = int32(*(*int16)(unsafe.Add(mBase, uint32(v311))))
												v332 = v322
											} else {
												v332 = v283<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v283&int32(63)
											}
											v334 = int32(base.Ui32(int32(base.Ui32(v312)>>(uint(int32(2))%32))+v320) >> (uint(int32(1)) % 32))
											v336 = v22 + int32(6)
											v337 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
											v343 = base.B2i32(int32(0) <= v282)
											if int32(0) <= v282 {
												v344 = int32(-8)
											} else {
												v344 = int32(-6)
											}
											if int32(0) <= v282 {
												v346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336))))
												v356 = v346
											} else {
												v356 = v281<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v281&int32(63)
											}
											v357 = int32(1)
											v358 = int32(base.Ui32(int32(base.Ui32(v337)>>(uint(int32(2))%32))+v344) >> (uint(v357) % 32))
											v364 = v281 & int32(_a_F_width_bucket_numeric_2)
											if v364 == int32(_a_F_width_bucket_numeric_9) {
												v367 = v281 << (uint(v357) % 32) & int32(_a_F_width_bucket_numeric_10)
											} else {
												v367 = v364
											}
											if v334 == int32(0) {
												if v358 == int32(0) {
													v423 = int32(0)
												} else {
													if v367 == int32(_a_F_width_bucket_numeric_10) {
														v377 = int32(1)
													} else {
														v377 = int32(-1)
													}
													v423 = v377
												}
											} else {
												v383 = v283 & int32(_a_F_width_bucket_numeric_2)
												if v383 == int32(_a_F_width_bucket_numeric_9) {
													v386 = v283 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v386 = v383
												}
												if v358 == int32(0) {
													if v386 != 0 {
														v391 = int32(-1)
													} else {
														v391 = int32(1)
													}
													v423 = v391
												} else {
													if v317 < int32(0) {
														v396 = v311
													} else {
														v396 = v17 + int32(8)
													}
													if v282 < int32(0) {
														v401 = v336
													} else {
														v401 = v22 + int32(8)
													}
													if v386 == int32(0) {
														if v367 == int32(_a_F_width_bucket_numeric_10) {
															v423 = int32(1)
														} else {
															v407 = F_cmp_abs_common(m, v396, v334, v332, v401, v358, v356)
															mBase = m.M
															v423 = v407
														}
													} else {
														if v367 == int32(0) {
															v423 = int32(-1)
														} else {
															v411 = F_cmp_abs_common(m, v401, v358, v356, v396, v334, v332)
															mBase = m.M
															v423 = v411
														}
													}
												}
											}
										}
									}
									if v423 < int32(0) {
										v427 = F_palloc(m, int32(2))
										mBase = m.M
										v428 = m.ExcPending
										if v428 != 0 {
											return int64(0)
										} else {
											v429 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
											v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
											*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
											v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
											*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
											*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
											v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
											mBase = m.M
											v956 = m.ExcPending
											if v956 != 0 {
												return int64(0)
											} else {
												if v955 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v1026 = m.ExcPending
													if v1026 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v1029 = m.ExcPending
														if v1029 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																mBase = m.M
																v1038 = m.ExcPending
																if v1038 != 0 {
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
													v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
													if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														F_pfree(m, v62)
														mBase = m.M
														v965 = m.ExcPending
														if v965 != 0 {
															return int64(0)
														} else {
															v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
															if v966 != 0 {
																F_pfree(m, v966)
																mBase = m.M
																v968 = m.ExcPending
																if v968 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
																}
															} else {
																m.G0 = v14 - int32(-64)
																return v959
															}
														}
													}
												}
											}
										}
									} else {
										v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
										v453 = base.I32_extend16_s(v452)
										v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v454) {
											if v454 != int32(_a_F_width_bucket_numeric_7) {
												if v454 != int32(_a_F_width_bucket_numeric_2) {
													if v453 != int32(-4096) {
														v473 = int32(-1)
													} else {
														v473 = int32(0)
													}
													v594 = v473
												} else {
													v594 = base.B2i32(v453 != int32(-16384))
												}
											} else {
												if v453 == int32(-16384) {
													v468 = int32(-1)
												} else {
													v468 = base.B2i32(v453 != int32(-12288))
												}
												v594 = v468
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v453) {
												if v453 == int32(-4096) {
													v480 = int32(1)
												} else {
													v480 = int32(-1)
												}
												v594 = v480
											} else {
												v482 = v17 + int32(6)
												v483 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v488 = base.I32_extend16_s(v454)
												v490 = base.B2i32(int32(0) <= v488)
												if int32(0) <= v488 {
													v491 = int32(-8)
												} else {
													v491 = int32(-6)
												}
												if int32(0) <= v488 {
													v493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v482))))
													v503 = v493
												} else {
													v503 = v454<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v454&int32(63)
												}
												v505 = int32(base.Ui32(int32(base.Ui32(v483)>>(uint(int32(2))%32))+v491) >> (uint(int32(1)) % 32))
												v507 = v25 + int32(6)
												v508 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
												v514 = base.B2i32(int32(0) <= v453)
												if int32(0) <= v453 {
													v515 = int32(-8)
												} else {
													v515 = int32(-6)
												}
												if int32(0) <= v453 {
													v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507))))
													v527 = v517
												} else {
													v527 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
												}
												v528 = int32(1)
												v529 = int32(base.Ui32(int32(base.Ui32(v508)>>(uint(int32(2))%32))+v515) >> (uint(v528) % 32))
												v535 = v452 & int32(_a_F_width_bucket_numeric_2)
												if v535 == int32(_a_F_width_bucket_numeric_9) {
													v538 = v452 << (uint(v528) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v538 = v535
												}
												if v505 == int32(0) {
													if v529 == int32(0) {
														v594 = int32(0)
													} else {
														if v538 == int32(_a_F_width_bucket_numeric_10) {
															v548 = int32(1)
														} else {
															v548 = int32(-1)
														}
														v594 = v548
													}
												} else {
													v554 = v454 & int32(_a_F_width_bucket_numeric_2)
													if v554 == int32(_a_F_width_bucket_numeric_9) {
														v557 = v454 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v557 = v554
													}
													if v529 == int32(0) {
														if v557 != 0 {
															v562 = int32(-1)
														} else {
															v562 = int32(1)
														}
														v594 = v562
													} else {
														if v488 < int32(0) {
															v567 = v482
														} else {
															v567 = v17 + int32(8)
														}
														if v453 < int32(0) {
															v572 = v507
														} else {
															v572 = v25 + int32(8)
														}
														if v557 == int32(0) {
															if v538 == int32(_a_F_width_bucket_numeric_10) {
																v594 = int32(1)
															} else {
																v578 = F_cmp_abs_common(m, v567, v505, v503, v572, v529, v527)
																mBase = m.M
																v594 = v578
															}
														} else {
															if v538 == int32(0) {
																v594 = int32(-1)
															} else {
																v582 = F_cmp_abs_common(m, v572, v529, v527, v567, v505, v503)
																mBase = m.M
																v594 = v582
															}
														}
													}
												}
											}
										}
										if int32(0) <= v594 {
											F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
											mBase = m.M
											v603 = m.ExcPending
											if v603 != 0 {
												return int64(0)
											} else {
												v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v956 = m.ExcPending
												if v956 != 0 {
													return int64(0)
												} else {
													if v955 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															F_pfree(m, v62)
															mBase = m.M
															v965 = m.ExcPending
															if v965 != 0 {
																return int64(0)
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v966 != 0 {
																	F_pfree(m, v966)
																	mBase = m.M
																	v968 = m.ExcPending
																	if v968 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
																}
															}
														}
													}
												}
											}
										} else {
											F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
											mBase = m.M
											v609 = m.ExcPending
											if v609 != 0 {
												return int64(0)
											} else {
												v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v956 = m.ExcPending
												if v956 != 0 {
													return int64(0)
												} else {
													if v955 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1026 = m.ExcPending
														if v1026 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1029 = m.ExcPending
															if v1029 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1038 = m.ExcPending
																	if v1038 != 0 {
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
														v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															F_pfree(m, v62)
															mBase = m.M
															v965 = m.ExcPending
															if v965 != 0 {
																return int64(0)
															} else {
																v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v966 != 0 {
																	F_pfree(m, v966)
																	mBase = m.M
																	v968 = m.ExcPending
																	if v968 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return v959
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
						if v31 == int32(_a_F_width_bucket_numeric_2) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v993 = m.ExcPending
							if v993 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(386138242))
								mBase = m.M
								v996 = m.ExcPending
								if v996 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
									mBase = m.M
									v1000 = m.ExcPending
									if v1000 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1978), int32(_a_F_width_bucket_numeric_5))
										mBase = m.M
										v1005 = m.ExcPending
										if v1005 != 0 {
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
							v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
							v40 = v39
							if v40&int32(_a_F_width_bucket_numeric_1) == int32(_a_F_width_bucket_numeric_2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v993 = m.ExcPending
								if v993 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v996 = m.ExcPending
									if v996 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
										mBase = m.M
										v1000 = m.ExcPending
										if v1000 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1978), int32(_a_F_width_bucket_numeric_5))
											mBase = m.M
											v1005 = m.ExcPending
											if v1005 != 0 {
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
								v45 = int32(_a_F_width_bucket_numeric_6)
								v47 = int32(_a_F_width_bucket_numeric_7)
								if base.B2i32(v31&v45 == v47)|base.B2i32(v40&v45 == v47) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1009 = m.ExcPending
									if v1009 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v1012 = m.ExcPending
										if v1012 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_8), int32(0))
											mBase = m.M
											v1016 = m.ExcPending
											if v1016 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1983), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v1021 = m.ExcPending
												if v1021 != 0 {
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
									v55 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v55
									*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v55
									*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v55
									v62 = F_palloc(m, int32(12))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v62
										v65 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
										*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
										v74 = v65
										v81 = v62 + int32(12)
										v83 = v27 & int64(2147483647)
										for {
											v86 = v81 - int32(2)
											v88 = base.I64_div_u_s(v83, int64(10000))
											v91 = v88*int64(55536) + v83
											*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v91)
											v94 = v74 + int32(1)
											if base.Ui64(int64(9999)) < base.Ui64(v83) {
												v74 = v94
												v81 = v86
												v83 = v88
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94
										*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v86
										v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
										v112 = base.I32_extend16_s(v111)
										v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v113) {
											if v113 != int32(_a_F_width_bucket_numeric_7) {
												if v113 != int32(_a_F_width_bucket_numeric_2) {
													if v112 != int32(-4096) {
														v132 = int32(-1)
													} else {
														v132 = int32(0)
													}
													v253 = v132
												} else {
													v253 = base.B2i32(v112 != int32(-16384))
												}
											} else {
												if v112 == int32(-16384) {
													v127 = int32(-1)
												} else {
													v127 = base.B2i32(v112 != int32(-12288))
												}
												v253 = v127
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v112) {
												if v112 == int32(-4096) {
													v139 = int32(1)
												} else {
													v139 = int32(-1)
												}
												v253 = v139
											} else {
												v141 = v22 + int32(6)
												v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												v147 = base.I32_extend16_s(v113)
												v149 = base.B2i32(int32(0) <= v147)
												if int32(0) <= v147 {
													v150 = int32(-8)
												} else {
													v150 = int32(-6)
												}
												if int32(0) <= v147 {
													v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141))))
													v162 = v152
												} else {
													v162 = v113<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v113&int32(63)
												}
												v164 = int32(base.Ui32(int32(base.Ui32(v142)>>(uint(int32(2))%32))+v150) >> (uint(int32(1)) % 32))
												v166 = v25 + int32(6)
												v167 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
												v173 = base.B2i32(int32(0) <= v112)
												if int32(0) <= v112 {
													v174 = int32(-8)
												} else {
													v174 = int32(-6)
												}
												if int32(0) <= v112 {
													v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v166))))
													v186 = v176
												} else {
													v186 = v111<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v111&int32(63)
												}
												v187 = int32(1)
												v188 = int32(base.Ui32(int32(base.Ui32(v167)>>(uint(int32(2))%32))+v174) >> (uint(v187) % 32))
												v194 = v111 & int32(_a_F_width_bucket_numeric_2)
												if v194 == int32(_a_F_width_bucket_numeric_9) {
													v197 = v111 << (uint(v187) % 32) & int32(_a_F_width_bucket_numeric_10)
												} else {
													v197 = v194
												}
												if v164 == int32(0) {
													if v188 == int32(0) {
														v253 = int32(0)
													} else {
														if v197 == int32(_a_F_width_bucket_numeric_10) {
															v207 = int32(1)
														} else {
															v207 = int32(-1)
														}
														v253 = v207
													}
												} else {
													v213 = v113 & int32(_a_F_width_bucket_numeric_2)
													if v213 == int32(_a_F_width_bucket_numeric_9) {
														v216 = v113 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v216 = v213
													}
													if v188 == int32(0) {
														if v216 != 0 {
															v221 = int32(-1)
														} else {
															v221 = int32(1)
														}
														v253 = v221
													} else {
														if v147 < int32(0) {
															v226 = v141
														} else {
															v226 = v22 + int32(8)
														}
														if v112 < int32(0) {
															v231 = v166
														} else {
															v231 = v25 + int32(8)
														}
														if v216 == int32(0) {
															if v197 == int32(_a_F_width_bucket_numeric_10) {
																v253 = int32(1)
															} else {
																v237 = F_cmp_abs_common(m, v226, v164, v162, v231, v188, v186)
																mBase = m.M
																v253 = v237
															}
														} else {
															if v197 == int32(0) {
																v253 = int32(-1)
															} else {
																v241 = F_cmp_abs_common(m, v231, v188, v186, v226, v164, v162)
																mBase = m.M
																v253 = v241
															}
														}
													}
												}
											}
										}
										switch v253 {
										case 0:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(386138242))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_width_bucket_numeric_11), int32(0))
													mBase = m.M
													v264 = m.ExcPending
													if v264 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1997), int32(_a_F_width_bucket_numeric_5))
														mBase = m.M
														v269 = m.ExcPending
														if v269 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										case 1:
											v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v622 = base.I32_extend16_s(v621)
											v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v623) {
												if v623 != int32(_a_F_width_bucket_numeric_7) {
													if v623 != int32(_a_F_width_bucket_numeric_2) {
														if v622 != int32(-4096) {
															v642 = int32(-1)
														} else {
															v642 = int32(0)
														}
														v763 = v642
													} else {
														v763 = base.B2i32(v622 != int32(-16384))
													}
												} else {
													if v622 == int32(-16384) {
														v637 = int32(-1)
													} else {
														v637 = base.B2i32(v622 != int32(-12288))
													}
													v763 = v637
												}
											} else {
												if base.Ui32(int32(-16384)) <= base.Ui32(v622) {
													if v622 == int32(-4096) {
														v649 = int32(1)
													} else {
														v649 = int32(-1)
													}
													v763 = v649
												} else {
													v651 = v17 + int32(6)
													v652 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v657 = base.I32_extend16_s(v623)
													v659 = base.B2i32(int32(0) <= v657)
													if int32(0) <= v657 {
														v660 = int32(-8)
													} else {
														v660 = int32(-6)
													}
													if int32(0) <= v657 {
														v662 = int32(*(*int16)(unsafe.Add(mBase, uint32(v651))))
														v672 = v662
													} else {
														v672 = v623<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v623&int32(63)
													}
													v674 = int32(base.Ui32(int32(base.Ui32(v652)>>(uint(int32(2))%32))+v660) >> (uint(int32(1)) % 32))
													v676 = v22 + int32(6)
													v677 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v683 = base.B2i32(int32(0) <= v622)
													if int32(0) <= v622 {
														v684 = int32(-8)
													} else {
														v684 = int32(-6)
													}
													if int32(0) <= v622 {
														v686 = int32(*(*int16)(unsafe.Add(mBase, uint32(v676))))
														v696 = v686
													} else {
														v696 = v621<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v621&int32(63)
													}
													v697 = int32(1)
													v698 = int32(base.Ui32(int32(base.Ui32(v677)>>(uint(int32(2))%32))+v684) >> (uint(v697) % 32))
													v704 = v621 & int32(_a_F_width_bucket_numeric_2)
													if v704 == int32(_a_F_width_bucket_numeric_9) {
														v707 = v621 << (uint(v697) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v707 = v704
													}
													if v674 == int32(0) {
														if v698 == int32(0) {
															v763 = int32(0)
														} else {
															if v707 == int32(_a_F_width_bucket_numeric_10) {
																v717 = int32(1)
															} else {
																v717 = int32(-1)
															}
															v763 = v717
														}
													} else {
														v723 = v623 & int32(_a_F_width_bucket_numeric_2)
														if v723 == int32(_a_F_width_bucket_numeric_9) {
															v726 = v623 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v726 = v723
														}
														if v698 == int32(0) {
															if v726 != 0 {
																v731 = int32(-1)
															} else {
																v731 = int32(1)
															}
															v763 = v731
														} else {
															if v657 < int32(0) {
																v736 = v651
															} else {
																v736 = v17 + int32(8)
															}
															if v622 < int32(0) {
																v741 = v676
															} else {
																v741 = v22 + int32(8)
															}
															if v726 == int32(0) {
																if v707 == int32(_a_F_width_bucket_numeric_10) {
																	v763 = int32(1)
																} else {
																	v747 = F_cmp_abs_common(m, v736, v674, v672, v741, v698, v696)
																	mBase = m.M
																	v763 = v747
																}
															} else {
																if v707 == int32(0) {
																	v763 = int32(-1)
																} else {
																	v751 = F_cmp_abs_common(m, v741, v698, v696, v736, v674, v672)
																	mBase = m.M
																	v763 = v751
																}
															}
														}
													}
												}
											}
											if int32(0) < v763 {
												v767 = F_palloc(m, int32(2))
												mBase = m.M
												v768 = m.ExcPending
												if v768 != 0 {
													return int64(0)
												} else {
													v769 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(v767))) = uint16(v769)
													v772 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v772
													v775 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v775
													*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v767 + int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v767
													v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v956 = m.ExcPending
													if v956 != 0 {
														return int64(0)
													} else {
														if v955 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																F_pfree(m, v62)
																mBase = m.M
																v965 = m.ExcPending
																if v965 != 0 {
																	return int64(0)
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v966 != 0 {
																		F_pfree(m, v966)
																		mBase = m.M
																		v968 = m.ExcPending
																		if v968 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																}
															}
														}
													}
												}
											} else {
												v792 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v793 = base.I32_extend16_s(v792)
												v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v794) {
													if v794 != int32(_a_F_width_bucket_numeric_7) {
														if v794 != int32(_a_F_width_bucket_numeric_2) {
															if v793 != int32(-4096) {
																v813 = int32(-1)
															} else {
																v813 = int32(0)
															}
															v934 = v813
														} else {
															v934 = base.B2i32(v793 != int32(-16384))
														}
													} else {
														if v793 == int32(-16384) {
															v808 = int32(-1)
														} else {
															v808 = base.B2i32(v793 != int32(-12288))
														}
														v934 = v808
													}
												} else {
													if base.Ui32(int32(-16384)) <= base.Ui32(v793) {
														if v793 == int32(-4096) {
															v820 = int32(1)
														} else {
															v820 = int32(-1)
														}
														v934 = v820
													} else {
														v822 = v17 + int32(6)
														v823 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v828 = base.I32_extend16_s(v794)
														v830 = base.B2i32(int32(0) <= v828)
														if int32(0) <= v828 {
															v831 = int32(-8)
														} else {
															v831 = int32(-6)
														}
														if int32(0) <= v828 {
															v833 = int32(*(*int16)(unsafe.Add(mBase, uint32(v822))))
															v843 = v833
														} else {
															v843 = v794<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v794&int32(63)
														}
														v845 = int32(base.Ui32(int32(base.Ui32(v823)>>(uint(int32(2))%32))+v831) >> (uint(int32(1)) % 32))
														v847 = v25 + int32(6)
														v848 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														v854 = base.B2i32(int32(0) <= v793)
														if int32(0) <= v793 {
															v855 = int32(-8)
														} else {
															v855 = int32(-6)
														}
														if int32(0) <= v793 {
															v857 = int32(*(*int16)(unsafe.Add(mBase, uint32(v847))))
															v867 = v857
														} else {
															v867 = v792<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v792&int32(63)
														}
														v868 = int32(1)
														v869 = int32(base.Ui32(int32(base.Ui32(v848)>>(uint(int32(2))%32))+v855) >> (uint(v868) % 32))
														v875 = v792 & int32(_a_F_width_bucket_numeric_2)
														if v875 == int32(_a_F_width_bucket_numeric_9) {
															v878 = v792 << (uint(v868) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v878 = v875
														}
														if v845 == int32(0) {
															if v869 == int32(0) {
																v934 = int32(0)
															} else {
																if v878 == int32(_a_F_width_bucket_numeric_10) {
																	v888 = int32(1)
																} else {
																	v888 = int32(-1)
																}
																v934 = v888
															}
														} else {
															v894 = v794 & int32(_a_F_width_bucket_numeric_2)
															if v894 == int32(_a_F_width_bucket_numeric_9) {
																v897 = v794 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
															} else {
																v897 = v894
															}
															if v869 == int32(0) {
																if v897 != 0 {
																	v902 = int32(-1)
																} else {
																	v902 = int32(1)
																}
																v934 = v902
															} else {
																if v828 < int32(0) {
																	v907 = v822
																} else {
																	v907 = v17 + int32(8)
																}
																if v793 < int32(0) {
																	v912 = v847
																} else {
																	v912 = v25 + int32(8)
																}
																if v897 == int32(0) {
																	if v878 == int32(_a_F_width_bucket_numeric_10) {
																		v934 = int32(1)
																	} else {
																		v918 = F_cmp_abs_common(m, v907, v845, v843, v912, v869, v867)
																		mBase = m.M
																		v934 = v918
																	}
																} else {
																	if v878 == int32(0) {
																		v934 = int32(-1)
																	} else {
																		v922 = F_cmp_abs_common(m, v912, v869, v867, v907, v845, v843)
																		mBase = m.M
																		v934 = v922
																	}
																}
															}
														}
													}
												}
												if v934 <= int32(0) {
													F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
													mBase = m.M
													v943 = m.ExcPending
													if v943 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												} else {
													F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
													mBase = m.M
													v949 = m.ExcPending
													if v949 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												}
											}
										default:
											v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v282 = base.I32_extend16_s(v281)
											v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v283) {
												if v283 != int32(_a_F_width_bucket_numeric_7) {
													if v283 != int32(_a_F_width_bucket_numeric_2) {
														if v282 != int32(-4096) {
															v302 = int32(-1)
														} else {
															v302 = int32(0)
														}
														v423 = v302
													} else {
														v423 = base.B2i32(v282 != int32(-16384))
													}
												} else {
													if v282 == int32(-16384) {
														v297 = int32(-1)
													} else {
														v297 = base.B2i32(v282 != int32(-12288))
													}
													v423 = v297
												}
											} else {
												if base.Ui32(int32(-16384)) <= base.Ui32(v282) {
													if v282 == int32(-4096) {
														v309 = int32(1)
													} else {
														v309 = int32(-1)
													}
													v423 = v309
												} else {
													v311 = v17 + int32(6)
													v312 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v317 = base.I32_extend16_s(v283)
													v319 = base.B2i32(int32(0) <= v317)
													if int32(0) <= v317 {
														v320 = int32(-8)
													} else {
														v320 = int32(-6)
													}
													if int32(0) <= v317 {
														v322 = int32(*(*int16)(unsafe.Add(mBase, uint32(v311))))
														v332 = v322
													} else {
														v332 = v283<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v283&int32(63)
													}
													v334 = int32(base.Ui32(int32(base.Ui32(v312)>>(uint(int32(2))%32))+v320) >> (uint(int32(1)) % 32))
													v336 = v22 + int32(6)
													v337 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v343 = base.B2i32(int32(0) <= v282)
													if int32(0) <= v282 {
														v344 = int32(-8)
													} else {
														v344 = int32(-6)
													}
													if int32(0) <= v282 {
														v346 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336))))
														v356 = v346
													} else {
														v356 = v281<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v281&int32(63)
													}
													v357 = int32(1)
													v358 = int32(base.Ui32(int32(base.Ui32(v337)>>(uint(int32(2))%32))+v344) >> (uint(v357) % 32))
													v364 = v281 & int32(_a_F_width_bucket_numeric_2)
													if v364 == int32(_a_F_width_bucket_numeric_9) {
														v367 = v281 << (uint(v357) % 32) & int32(_a_F_width_bucket_numeric_10)
													} else {
														v367 = v364
													}
													if v334 == int32(0) {
														if v358 == int32(0) {
															v423 = int32(0)
														} else {
															if v367 == int32(_a_F_width_bucket_numeric_10) {
																v377 = int32(1)
															} else {
																v377 = int32(-1)
															}
															v423 = v377
														}
													} else {
														v383 = v283 & int32(_a_F_width_bucket_numeric_2)
														if v383 == int32(_a_F_width_bucket_numeric_9) {
															v386 = v283 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v386 = v383
														}
														if v358 == int32(0) {
															if v386 != 0 {
																v391 = int32(-1)
															} else {
																v391 = int32(1)
															}
															v423 = v391
														} else {
															if v317 < int32(0) {
																v396 = v311
															} else {
																v396 = v17 + int32(8)
															}
															if v282 < int32(0) {
																v401 = v336
															} else {
																v401 = v22 + int32(8)
															}
															if v386 == int32(0) {
																if v367 == int32(_a_F_width_bucket_numeric_10) {
																	v423 = int32(1)
																} else {
																	v407 = F_cmp_abs_common(m, v396, v334, v332, v401, v358, v356)
																	mBase = m.M
																	v423 = v407
																}
															} else {
																if v367 == int32(0) {
																	v423 = int32(-1)
																} else {
																	v411 = F_cmp_abs_common(m, v401, v358, v356, v396, v334, v332)
																	mBase = m.M
																	v423 = v411
																}
															}
														}
													}
												}
											}
											if v423 < int32(0) {
												v427 = F_palloc(m, int32(2))
												mBase = m.M
												v428 = m.ExcPending
												if v428 != 0 {
													return int64(0)
												} else {
													v429 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
													v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
													v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
													*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
													*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
													v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v956 = m.ExcPending
													if v956 != 0 {
														return int64(0)
													} else {
														if v955 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1026 = m.ExcPending
															if v1026 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1029 = m.ExcPending
																if v1029 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1038 = m.ExcPending
																		if v1038 != 0 {
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
															v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																F_pfree(m, v62)
																mBase = m.M
																v965 = m.ExcPending
																if v965 != 0 {
																	return int64(0)
																} else {
																	v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v966 != 0 {
																		F_pfree(m, v966)
																		mBase = m.M
																		v968 = m.ExcPending
																		if v968 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return v959
																	}
																}
															}
														}
													}
												}
											} else {
												v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v453 = base.I32_extend16_s(v452)
												v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v454) {
													if v454 != int32(_a_F_width_bucket_numeric_7) {
														if v454 != int32(_a_F_width_bucket_numeric_2) {
															if v453 != int32(-4096) {
																v473 = int32(-1)
															} else {
																v473 = int32(0)
															}
															v594 = v473
														} else {
															v594 = base.B2i32(v453 != int32(-16384))
														}
													} else {
														if v453 == int32(-16384) {
															v468 = int32(-1)
														} else {
															v468 = base.B2i32(v453 != int32(-12288))
														}
														v594 = v468
													}
												} else {
													if base.Ui32(int32(-16384)) <= base.Ui32(v453) {
														if v453 == int32(-4096) {
															v480 = int32(1)
														} else {
															v480 = int32(-1)
														}
														v594 = v480
													} else {
														v482 = v17 + int32(6)
														v483 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v488 = base.I32_extend16_s(v454)
														v490 = base.B2i32(int32(0) <= v488)
														if int32(0) <= v488 {
															v491 = int32(-8)
														} else {
															v491 = int32(-6)
														}
														if int32(0) <= v488 {
															v493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v482))))
															v503 = v493
														} else {
															v503 = v454<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v454&int32(63)
														}
														v505 = int32(base.Ui32(int32(base.Ui32(v483)>>(uint(int32(2))%32))+v491) >> (uint(int32(1)) % 32))
														v507 = v25 + int32(6)
														v508 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														v514 = base.B2i32(int32(0) <= v453)
														if int32(0) <= v453 {
															v515 = int32(-8)
														} else {
															v515 = int32(-6)
														}
														if int32(0) <= v453 {
															v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v507))))
															v527 = v517
														} else {
															v527 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
														}
														v528 = int32(1)
														v529 = int32(base.Ui32(int32(base.Ui32(v508)>>(uint(int32(2))%32))+v515) >> (uint(v528) % 32))
														v535 = v452 & int32(_a_F_width_bucket_numeric_2)
														if v535 == int32(_a_F_width_bucket_numeric_9) {
															v538 = v452 << (uint(v528) % 32) & int32(_a_F_width_bucket_numeric_10)
														} else {
															v538 = v535
														}
														if v505 == int32(0) {
															if v529 == int32(0) {
																v594 = int32(0)
															} else {
																if v538 == int32(_a_F_width_bucket_numeric_10) {
																	v548 = int32(1)
																} else {
																	v548 = int32(-1)
																}
																v594 = v548
															}
														} else {
															v554 = v454 & int32(_a_F_width_bucket_numeric_2)
															if v554 == int32(_a_F_width_bucket_numeric_9) {
																v557 = v454 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_10)
															} else {
																v557 = v554
															}
															if v529 == int32(0) {
																if v557 != 0 {
																	v562 = int32(-1)
																} else {
																	v562 = int32(1)
																}
																v594 = v562
															} else {
																if v488 < int32(0) {
																	v567 = v482
																} else {
																	v567 = v17 + int32(8)
																}
																if v453 < int32(0) {
																	v572 = v507
																} else {
																	v572 = v25 + int32(8)
																}
																if v557 == int32(0) {
																	if v538 == int32(_a_F_width_bucket_numeric_10) {
																		v594 = int32(1)
																	} else {
																		v578 = F_cmp_abs_common(m, v567, v505, v503, v572, v529, v527)
																		mBase = m.M
																		v594 = v578
																	}
																} else {
																	if v538 == int32(0) {
																		v594 = int32(-1)
																	} else {
																		v582 = F_cmp_abs_common(m, v572, v529, v527, v567, v505, v503)
																		mBase = m.M
																		v594 = v582
																	}
																}
															}
														}
													}
												}
												if int32(0) <= v594 {
													F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_13), v12+int32(-56))
													mBase = m.M
													v603 = m.ExcPending
													if v603 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
																		}
																	}
																}
															}
														}
													}
												} else {
													F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
													mBase = m.M
													v609 = m.ExcPending
													if v609 != 0 {
														return int64(0)
													} else {
														v955 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v956 = m.ExcPending
														if v956 != 0 {
															return int64(0)
														} else {
															if v955 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1029 = m.ExcPending
																	if v1029 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1038 = m.ExcPending
																			if v1038 != 0 {
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
																v959 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v959-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int64(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1029 = m.ExcPending
																		if v1029 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2027), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1038 = m.ExcPending
																				if v1038 != 0 {
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
																	F_pfree(m, v62)
																	mBase = m.M
																	v965 = m.ExcPending
																	if v965 != 0 {
																		return int64(0)
																	} else {
																		v966 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v966 != 0 {
																			F_pfree(m, v966)
																			mBase = m.M
																			v968 = m.ExcPending
																			if v968 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return v959
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return v959
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
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v976 = m.ExcPending
					if v976 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(386138242))
						mBase = m.M
						v979 = m.ExcPending
						if v979 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_width_bucket_numeric_14), int32(0))
							mBase = m.M
							v983 = m.ExcPending
							if v983 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1971), int32(_a_F_width_bucket_numeric_5))
								mBase = m.M
								v988 = m.ExcPending
								if v988 != 0 {
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
