package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_width_bucket_numeric(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int64
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
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
	var v455 int32
	_ = v455
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v925 int32
	_ = v925
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int64
	_ = v950
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v25 = F_pg_detoast_datum(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if int32(0) < v27 {
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
					if base.Ui32(v30) <= base.Ui32(int32(_a_F_width_bucket_numeric_0)) {
						v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
						if base.Ui32(int32(_a_F_width_bucket_numeric_0)) < base.Ui32(v33) {
							v42 = v33
							if v42&int32(_a_F_width_bucket_numeric_1) == int32(_a_F_width_bucket_numeric_2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v986 = m.ExcPending
								if v986 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v989 = m.ExcPending
									if v989 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
										mBase = m.M
										v993 = m.ExcPending
										if v993 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
											mBase = m.M
											v998 = m.ExcPending
											if v998 != 0 {
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
								v47 = v42
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
								if v48 == int32(_a_F_width_bucket_numeric_2) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v986 = m.ExcPending
									if v986 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v989 = m.ExcPending
										if v989 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
											mBase = m.M
											v993 = m.ExcPending
											if v993 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v998 = m.ExcPending
												if v998 != 0 {
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
									v51 = int32(_a_F_width_bucket_numeric_6)
									v53 = int32(_a_F_width_bucket_numeric_7)
									if base.B2i32(v47&v51 == v53)|base.B2i32(v48&v51 == v53) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1002 = m.ExcPending
										if v1002 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1005 = m.ExcPending
											if v1005 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_width_bucket_numeric_8), int32(0))
												mBase = m.M
												v1009 = m.ExcPending
												if v1009 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1996), int32(_a_F_width_bucket_numeric_5))
													mBase = m.M
													v1014 = m.ExcPending
													if v1014 != 0 {
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
										v62 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v62
										v69 = F_palloc(m, int32(12))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v69
											v72 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v72)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
											v80 = v72
											v87 = v69 + int32(12)
											v89 = base.I64_extend_i32_u(v27)
											for {
												v92 = v87 - int32(2)
												v94 = base.I64_div_u_s(v89, int64(10000))
												v97 = v94*int64(55536) + v89
												*(*uint16)(unsafe.Add(mBase, uint32(v92))) = uint16(v97)
												v100 = v80 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v89) {
													v80 = v100
													v87 = v92
													v89 = v94
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v80
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v100
											*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v92
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v119 = int32(_a_F_width_bucket_numeric_2)
											v120 = v118 & v119
											if v120 == v119 {
												if v118 != int32(_a_F_width_bucket_numeric_7) {
													if v118 != int32(_a_F_width_bucket_numeric_2) {
														if v117 != int32(_a_F_width_bucket_numeric_9) {
															v139 = int32(-1)
														} else {
															v139 = int32(0)
														}
														v256 = v139
													} else {
														v256 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_2))
													}
												} else {
													if v117 == int32(_a_F_width_bucket_numeric_2) {
														v134 = int32(-1)
													} else {
														v134 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_7))
													}
													v256 = v134
												}
											} else {
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v117) {
													if v117 == int32(_a_F_width_bucket_numeric_9) {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v256 = v146
												} else {
													v148 = v22 + int32(6)
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v156 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
													if int32(0) <= base.I32_extend16_s(v118) {
														v157 = int32(-8)
													} else {
														v157 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v118) {
														v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
														v169 = v159
													} else {
														v169 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
													}
													v171 = int32(base.Ui32(int32(base.Ui32(v149)>>(uint(int32(2))%32))+v157) >> (uint(int32(1)) % 32))
													v173 = v25 + int32(6)
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													v181 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
													if int32(0) <= base.I32_extend16_s(v117) {
														v182 = int32(-8)
													} else {
														v182 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v117) {
														v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173))))
														v194 = v184
													} else {
														v194 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
													}
													v195 = int32(1)
													v196 = int32(base.Ui32(int32(base.Ui32(v174)>>(uint(int32(2))%32))+v182) >> (uint(v195) % 32))
													v202 = v117 & int32(_a_F_width_bucket_numeric_2)
													if v202 == int32(_a_F_width_bucket_numeric_10) {
														v205 = v117 << (uint(v195) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v205 = v202
													}
													if v171 == int32(0) {
														if v196 == int32(0) {
															v256 = int32(0)
														} else {
															if v205 == int32(_a_F_width_bucket_numeric_11) {
																v215 = int32(1)
															} else {
																v215 = int32(-1)
															}
															v256 = v215
														}
													} else {
														if v120 == int32(_a_F_width_bucket_numeric_10) {
															v222 = v118 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v222 = v120
														}
														if v196 == int32(0) {
															if v222 != 0 {
																v227 = int32(-1)
															} else {
																v227 = int32(1)
															}
															v256 = v227
														} else {
															if int32(0) <= base.I32_extend16_s(v118) {
																v230 = v22 + int32(8)
															} else {
																v230 = v148
															}
															if int32(0) <= base.I32_extend16_s(v117) {
																v233 = v25 + int32(8)
															} else {
																v233 = v173
															}
															if v222 == int32(0) {
																if v205 == int32(_a_F_width_bucket_numeric_11) {
																	v256 = int32(1)
																} else {
																	v239 = F_cmp_abs_common(m, v230, v171, v169, v233, v196, v194)
																	mBase = m.M
																	v256 = v239
																}
															} else {
																if v205 == int32(0) {
																	v256 = int32(-1)
																} else {
																	v243 = F_cmp_abs_common(m, v233, v196, v194, v230, v171, v169)
																	mBase = m.M
																	v256 = v243
																}
															}
														}
													}
												}
											}
											switch v256 {
											case 0:
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(386138242))
													mBase = m.M
													v263 = m.ExcPending
													if v263 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2010), int32(_a_F_width_bucket_numeric_5))
															mBase = m.M
															v272 = m.ExcPending
															if v272 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											case 1:
												v618 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v620 = int32(_a_F_width_bucket_numeric_2)
												v621 = v619 & v620
												if v621 == v620 {
													if v619 != int32(_a_F_width_bucket_numeric_7) {
														if v619 != int32(_a_F_width_bucket_numeric_2) {
															if v618 != int32(_a_F_width_bucket_numeric_9) {
																v640 = int32(-1)
															} else {
																v640 = int32(0)
															}
															v757 = v640
														} else {
															v757 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v618 == int32(_a_F_width_bucket_numeric_2) {
															v635 = int32(-1)
														} else {
															v635 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_7))
														}
														v757 = v635
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v618) {
														if v618 == int32(_a_F_width_bucket_numeric_9) {
															v647 = int32(1)
														} else {
															v647 = int32(-1)
														}
														v757 = v647
													} else {
														v649 = v17 + int32(6)
														v650 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v657 = base.B2i32(int32(0) <= base.I32_extend16_s(v619))
														if int32(0) <= base.I32_extend16_s(v619) {
															v658 = int32(-8)
														} else {
															v658 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v619) {
															v660 = int32(*(*int16)(unsafe.Add(mBase, uint32(v649))))
															v670 = v660
														} else {
															v670 = v619<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v619&int32(63)
														}
														v672 = int32(base.Ui32(int32(base.Ui32(v650)>>(uint(int32(2))%32))+v658) >> (uint(int32(1)) % 32))
														v674 = v22 + int32(6)
														v675 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v682 = base.B2i32(int32(0) <= base.I32_extend16_s(v618))
														if int32(0) <= base.I32_extend16_s(v618) {
															v683 = int32(-8)
														} else {
															v683 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v618) {
															v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674))))
															v695 = v685
														} else {
															v695 = v618<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v618&int32(63)
														}
														v696 = int32(1)
														v697 = int32(base.Ui32(int32(base.Ui32(v675)>>(uint(int32(2))%32))+v683) >> (uint(v696) % 32))
														v703 = v618 & int32(_a_F_width_bucket_numeric_2)
														if v703 == int32(_a_F_width_bucket_numeric_10) {
															v706 = v618 << (uint(v696) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v706 = v703
														}
														if v672 == int32(0) {
															if v697 == int32(0) {
																v757 = int32(0)
															} else {
																if v706 == int32(_a_F_width_bucket_numeric_11) {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v757 = v716
															}
														} else {
															if v621 == int32(_a_F_width_bucket_numeric_10) {
																v723 = v619 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v723 = v621
															}
															if v697 == int32(0) {
																if v723 != 0 {
																	v728 = int32(-1)
																} else {
																	v728 = int32(1)
																}
																v757 = v728
															} else {
																if int32(0) <= base.I32_extend16_s(v619) {
																	v731 = v17 + int32(8)
																} else {
																	v731 = v649
																}
																if int32(0) <= base.I32_extend16_s(v618) {
																	v734 = v22 + int32(8)
																} else {
																	v734 = v674
																}
																if v723 == int32(0) {
																	if v706 == int32(_a_F_width_bucket_numeric_11) {
																		v757 = int32(1)
																	} else {
																		v740 = F_cmp_abs_common(m, v731, v672, v670, v734, v697, v695)
																		mBase = m.M
																		v757 = v740
																	}
																} else {
																	if v706 == int32(0) {
																		v757 = int32(-1)
																	} else {
																		v744 = F_cmp_abs_common(m, v734, v697, v695, v731, v672, v670)
																		mBase = m.M
																		v757 = v744
																	}
																}
															}
														}
													}
												}
												if int32(0) < v757 {
													v761 = F_palloc(m, int32(2))
													mBase = m.M
													v762 = m.ExcPending
													if v762 != 0 {
														return int32(0)
													} else {
														v763 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v763)
														v766 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v766
														v769 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v769
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v761 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v761
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v788 = int32(_a_F_width_bucket_numeric_2)
													v789 = v787 & v788
													if v789 == v788 {
														if v787 != int32(_a_F_width_bucket_numeric_7) {
															if v787 != int32(_a_F_width_bucket_numeric_2) {
																if v786 != int32(_a_F_width_bucket_numeric_9) {
																	v808 = int32(-1)
																} else {
																	v808 = int32(0)
																}
																v925 = v808
															} else {
																v925 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v786 == int32(_a_F_width_bucket_numeric_2) {
																v803 = int32(-1)
															} else {
																v803 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_7))
															}
															v925 = v803
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v786) {
															if v786 == int32(_a_F_width_bucket_numeric_9) {
																v815 = int32(1)
															} else {
																v815 = int32(-1)
															}
															v925 = v815
														} else {
															v817 = v17 + int32(6)
															v818 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v825 = base.B2i32(int32(0) <= base.I32_extend16_s(v787))
															if int32(0) <= base.I32_extend16_s(v787) {
																v826 = int32(-8)
															} else {
																v826 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v787) {
																v828 = int32(*(*int16)(unsafe.Add(mBase, uint32(v817))))
																v838 = v828
															} else {
																v838 = v787<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v787&int32(63)
															}
															v840 = int32(base.Ui32(int32(base.Ui32(v818)>>(uint(int32(2))%32))+v826) >> (uint(int32(1)) % 32))
															v842 = v25 + int32(6)
															v843 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v850 = base.B2i32(int32(0) <= base.I32_extend16_s(v786))
															if int32(0) <= base.I32_extend16_s(v786) {
																v851 = int32(-8)
															} else {
																v851 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v786) {
																v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v842))))
																v863 = v853
															} else {
																v863 = v786<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v786&int32(63)
															}
															v864 = int32(1)
															v865 = int32(base.Ui32(int32(base.Ui32(v843)>>(uint(int32(2))%32))+v851) >> (uint(v864) % 32))
															v871 = v786 & int32(_a_F_width_bucket_numeric_2)
															if v871 == int32(_a_F_width_bucket_numeric_10) {
																v874 = v786 << (uint(v864) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v874 = v871
															}
															if v840 == int32(0) {
																if v865 == int32(0) {
																	v925 = int32(0)
																} else {
																	if v874 == int32(_a_F_width_bucket_numeric_11) {
																		v884 = int32(1)
																	} else {
																		v884 = int32(-1)
																	}
																	v925 = v884
																}
															} else {
																if v789 == int32(_a_F_width_bucket_numeric_10) {
																	v891 = v787 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v891 = v789
																}
																if v865 == int32(0) {
																	if v891 != 0 {
																		v896 = int32(-1)
																	} else {
																		v896 = int32(1)
																	}
																	v925 = v896
																} else {
																	if int32(0) <= base.I32_extend16_s(v787) {
																		v899 = v17 + int32(8)
																	} else {
																		v899 = v817
																	}
																	if int32(0) <= base.I32_extend16_s(v786) {
																		v902 = v25 + int32(8)
																	} else {
																		v902 = v842
																	}
																	if v891 == int32(0) {
																		if v874 == int32(_a_F_width_bucket_numeric_11) {
																			v925 = int32(1)
																		} else {
																			v908 = F_cmp_abs_common(m, v899, v840, v838, v902, v865, v863)
																			mBase = m.M
																			v925 = v908
																		}
																	} else {
																		if v874 == int32(0) {
																			v925 = int32(-1)
																		} else {
																			v912 = F_cmp_abs_common(m, v902, v865, v863, v899, v840, v838)
																			mBase = m.M
																			v925 = v912
																		}
																	}
																}
															}
														}
													}
													if v925 <= int32(0) {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v934 = m.ExcPending
														if v934 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v940 = m.ExcPending
														if v940 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													}
												}
											default:
												v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v286 = int32(_a_F_width_bucket_numeric_2)
												v287 = v285 & v286
												if v287 == v286 {
													if v285 != int32(_a_F_width_bucket_numeric_7) {
														if v285 != int32(_a_F_width_bucket_numeric_2) {
															if v284 != int32(_a_F_width_bucket_numeric_9) {
																v306 = int32(-1)
															} else {
																v306 = int32(0)
															}
															v423 = v306
														} else {
															v423 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v284 == int32(_a_F_width_bucket_numeric_2) {
															v301 = int32(-1)
														} else {
															v301 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_7))
														}
														v423 = v301
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v284) {
														if v284 == int32(_a_F_width_bucket_numeric_9) {
															v313 = int32(1)
														} else {
															v313 = int32(-1)
														}
														v423 = v313
													} else {
														v315 = v17 + int32(6)
														v316 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v285))
														if int32(0) <= base.I32_extend16_s(v285) {
															v324 = int32(-8)
														} else {
															v324 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v285) {
															v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v315))))
															v336 = v326
														} else {
															v336 = v285<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v285&int32(63)
														}
														v338 = int32(base.Ui32(int32(base.Ui32(v316)>>(uint(int32(2))%32))+v324) >> (uint(int32(1)) % 32))
														v340 = v22 + int32(6)
														v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v348 = base.B2i32(int32(0) <= base.I32_extend16_s(v284))
														if int32(0) <= base.I32_extend16_s(v284) {
															v349 = int32(-8)
														} else {
															v349 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v284) {
															v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v340))))
															v361 = v351
														} else {
															v361 = v284<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v284&int32(63)
														}
														v362 = int32(1)
														v363 = int32(base.Ui32(int32(base.Ui32(v341)>>(uint(int32(2))%32))+v349) >> (uint(v362) % 32))
														v369 = v284 & int32(_a_F_width_bucket_numeric_2)
														if v369 == int32(_a_F_width_bucket_numeric_10) {
															v372 = v284 << (uint(v362) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v372 = v369
														}
														if v338 == int32(0) {
															if v363 == int32(0) {
																v423 = int32(0)
															} else {
																if v372 == int32(_a_F_width_bucket_numeric_11) {
																	v382 = int32(1)
																} else {
																	v382 = int32(-1)
																}
																v423 = v382
															}
														} else {
															if v287 == int32(_a_F_width_bucket_numeric_10) {
																v389 = v285 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v389 = v287
															}
															if v363 == int32(0) {
																if v389 != 0 {
																	v394 = int32(-1)
																} else {
																	v394 = int32(1)
																}
																v423 = v394
															} else {
																if int32(0) <= base.I32_extend16_s(v285) {
																	v397 = v17 + int32(8)
																} else {
																	v397 = v315
																}
																if int32(0) <= base.I32_extend16_s(v284) {
																	v400 = v22 + int32(8)
																} else {
																	v400 = v340
																}
																if v389 == int32(0) {
																	if v372 == int32(_a_F_width_bucket_numeric_11) {
																		v423 = int32(1)
																	} else {
																		v406 = F_cmp_abs_common(m, v397, v338, v336, v400, v363, v361)
																		mBase = m.M
																		v423 = v406
																	}
																} else {
																	if v372 == int32(0) {
																		v423 = int32(-1)
																	} else {
																		v410 = F_cmp_abs_common(m, v400, v363, v361, v397, v338, v336)
																		mBase = m.M
																		v423 = v410
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
														return int32(0)
													} else {
														v429 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
														v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
														v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v454 = int32(_a_F_width_bucket_numeric_2)
													v455 = v453 & v454
													if v455 == v454 {
														if v453 != int32(_a_F_width_bucket_numeric_7) {
															if v453 != int32(_a_F_width_bucket_numeric_2) {
																if v452 != int32(_a_F_width_bucket_numeric_9) {
																	v474 = int32(-1)
																} else {
																	v474 = int32(0)
																}
																v591 = v474
															} else {
																v591 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v452 == int32(_a_F_width_bucket_numeric_2) {
																v469 = int32(-1)
															} else {
																v469 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_7))
															}
															v591 = v469
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v452) {
															if v452 == int32(_a_F_width_bucket_numeric_9) {
																v481 = int32(1)
															} else {
																v481 = int32(-1)
															}
															v591 = v481
														} else {
															v483 = v17 + int32(6)
															v484 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v491 = base.B2i32(int32(0) <= base.I32_extend16_s(v453))
															if int32(0) <= base.I32_extend16_s(v453) {
																v492 = int32(-8)
															} else {
																v492 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v453) {
																v494 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483))))
																v504 = v494
															} else {
																v504 = v453<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v453&int32(63)
															}
															v506 = int32(base.Ui32(int32(base.Ui32(v484)>>(uint(int32(2))%32))+v492) >> (uint(int32(1)) % 32))
															v508 = v25 + int32(6)
															v509 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v516 = base.B2i32(int32(0) <= base.I32_extend16_s(v452))
															if int32(0) <= base.I32_extend16_s(v452) {
																v517 = int32(-8)
															} else {
																v517 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v452) {
																v519 = int32(*(*int16)(unsafe.Add(mBase, uint32(v508))))
																v529 = v519
															} else {
																v529 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
															}
															v530 = int32(1)
															v531 = int32(base.Ui32(int32(base.Ui32(v509)>>(uint(int32(2))%32))+v517) >> (uint(v530) % 32))
															v537 = v452 & int32(_a_F_width_bucket_numeric_2)
															if v537 == int32(_a_F_width_bucket_numeric_10) {
																v540 = v452 << (uint(v530) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v540 = v537
															}
															if v506 == int32(0) {
																if v531 == int32(0) {
																	v591 = int32(0)
																} else {
																	if v540 == int32(_a_F_width_bucket_numeric_11) {
																		v550 = int32(1)
																	} else {
																		v550 = int32(-1)
																	}
																	v591 = v550
																}
															} else {
																if v455 == int32(_a_F_width_bucket_numeric_10) {
																	v557 = v453 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v557 = v455
																}
																if v531 == int32(0) {
																	if v557 != 0 {
																		v562 = int32(-1)
																	} else {
																		v562 = int32(1)
																	}
																	v591 = v562
																} else {
																	if int32(0) <= base.I32_extend16_s(v453) {
																		v565 = v17 + int32(8)
																	} else {
																		v565 = v483
																	}
																	if int32(0) <= base.I32_extend16_s(v452) {
																		v568 = v25 + int32(8)
																	} else {
																		v568 = v508
																	}
																	if v557 == int32(0) {
																		if v540 == int32(_a_F_width_bucket_numeric_11) {
																			v591 = int32(1)
																		} else {
																			v574 = F_cmp_abs_common(m, v565, v506, v504, v568, v531, v529)
																			mBase = m.M
																			v591 = v574
																		}
																	} else {
																		if v540 == int32(0) {
																			v591 = int32(-1)
																		} else {
																			v578 = F_cmp_abs_common(m, v568, v531, v529, v565, v506, v504)
																			mBase = m.M
																			v591 = v578
																		}
																	}
																}
															}
														}
													}
													if int32(0) <= v591 {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v600 = m.ExcPending
														if v600 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v606 = m.ExcPending
														if v606 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
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
							v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
							if base.Ui32(int32(_a_F_width_bucket_numeric_0)) < base.Ui32(v36) {
								v47 = v33
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
								if v48 == int32(_a_F_width_bucket_numeric_2) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v986 = m.ExcPending
									if v986 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v989 = m.ExcPending
										if v989 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
											mBase = m.M
											v993 = m.ExcPending
											if v993 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v998 = m.ExcPending
												if v998 != 0 {
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
									v51 = int32(_a_F_width_bucket_numeric_6)
									v53 = int32(_a_F_width_bucket_numeric_7)
									if base.B2i32(v47&v51 == v53)|base.B2i32(v48&v51 == v53) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1002 = m.ExcPending
										if v1002 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1005 = m.ExcPending
											if v1005 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_width_bucket_numeric_8), int32(0))
												mBase = m.M
												v1009 = m.ExcPending
												if v1009 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1996), int32(_a_F_width_bucket_numeric_5))
													mBase = m.M
													v1014 = m.ExcPending
													if v1014 != 0 {
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
										v62 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v62
										v69 = F_palloc(m, int32(12))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v69
											v72 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v72)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
											v80 = v72
											v87 = v69 + int32(12)
											v89 = base.I64_extend_i32_u(v27)
											for {
												v92 = v87 - int32(2)
												v94 = base.I64_div_u_s(v89, int64(10000))
												v97 = v94*int64(55536) + v89
												*(*uint16)(unsafe.Add(mBase, uint32(v92))) = uint16(v97)
												v100 = v80 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v89) {
													v80 = v100
													v87 = v92
													v89 = v94
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v80
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v100
											*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v92
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v119 = int32(_a_F_width_bucket_numeric_2)
											v120 = v118 & v119
											if v120 == v119 {
												if v118 != int32(_a_F_width_bucket_numeric_7) {
													if v118 != int32(_a_F_width_bucket_numeric_2) {
														if v117 != int32(_a_F_width_bucket_numeric_9) {
															v139 = int32(-1)
														} else {
															v139 = int32(0)
														}
														v256 = v139
													} else {
														v256 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_2))
													}
												} else {
													if v117 == int32(_a_F_width_bucket_numeric_2) {
														v134 = int32(-1)
													} else {
														v134 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_7))
													}
													v256 = v134
												}
											} else {
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v117) {
													if v117 == int32(_a_F_width_bucket_numeric_9) {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v256 = v146
												} else {
													v148 = v22 + int32(6)
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v156 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
													if int32(0) <= base.I32_extend16_s(v118) {
														v157 = int32(-8)
													} else {
														v157 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v118) {
														v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
														v169 = v159
													} else {
														v169 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
													}
													v171 = int32(base.Ui32(int32(base.Ui32(v149)>>(uint(int32(2))%32))+v157) >> (uint(int32(1)) % 32))
													v173 = v25 + int32(6)
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													v181 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
													if int32(0) <= base.I32_extend16_s(v117) {
														v182 = int32(-8)
													} else {
														v182 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v117) {
														v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173))))
														v194 = v184
													} else {
														v194 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
													}
													v195 = int32(1)
													v196 = int32(base.Ui32(int32(base.Ui32(v174)>>(uint(int32(2))%32))+v182) >> (uint(v195) % 32))
													v202 = v117 & int32(_a_F_width_bucket_numeric_2)
													if v202 == int32(_a_F_width_bucket_numeric_10) {
														v205 = v117 << (uint(v195) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v205 = v202
													}
													if v171 == int32(0) {
														if v196 == int32(0) {
															v256 = int32(0)
														} else {
															if v205 == int32(_a_F_width_bucket_numeric_11) {
																v215 = int32(1)
															} else {
																v215 = int32(-1)
															}
															v256 = v215
														}
													} else {
														if v120 == int32(_a_F_width_bucket_numeric_10) {
															v222 = v118 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v222 = v120
														}
														if v196 == int32(0) {
															if v222 != 0 {
																v227 = int32(-1)
															} else {
																v227 = int32(1)
															}
															v256 = v227
														} else {
															if int32(0) <= base.I32_extend16_s(v118) {
																v230 = v22 + int32(8)
															} else {
																v230 = v148
															}
															if int32(0) <= base.I32_extend16_s(v117) {
																v233 = v25 + int32(8)
															} else {
																v233 = v173
															}
															if v222 == int32(0) {
																if v205 == int32(_a_F_width_bucket_numeric_11) {
																	v256 = int32(1)
																} else {
																	v239 = F_cmp_abs_common(m, v230, v171, v169, v233, v196, v194)
																	mBase = m.M
																	v256 = v239
																}
															} else {
																if v205 == int32(0) {
																	v256 = int32(-1)
																} else {
																	v243 = F_cmp_abs_common(m, v233, v196, v194, v230, v171, v169)
																	mBase = m.M
																	v256 = v243
																}
															}
														}
													}
												}
											}
											switch v256 {
											case 0:
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(386138242))
													mBase = m.M
													v263 = m.ExcPending
													if v263 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2010), int32(_a_F_width_bucket_numeric_5))
															mBase = m.M
															v272 = m.ExcPending
															if v272 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											case 1:
												v618 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v620 = int32(_a_F_width_bucket_numeric_2)
												v621 = v619 & v620
												if v621 == v620 {
													if v619 != int32(_a_F_width_bucket_numeric_7) {
														if v619 != int32(_a_F_width_bucket_numeric_2) {
															if v618 != int32(_a_F_width_bucket_numeric_9) {
																v640 = int32(-1)
															} else {
																v640 = int32(0)
															}
															v757 = v640
														} else {
															v757 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v618 == int32(_a_F_width_bucket_numeric_2) {
															v635 = int32(-1)
														} else {
															v635 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_7))
														}
														v757 = v635
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v618) {
														if v618 == int32(_a_F_width_bucket_numeric_9) {
															v647 = int32(1)
														} else {
															v647 = int32(-1)
														}
														v757 = v647
													} else {
														v649 = v17 + int32(6)
														v650 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v657 = base.B2i32(int32(0) <= base.I32_extend16_s(v619))
														if int32(0) <= base.I32_extend16_s(v619) {
															v658 = int32(-8)
														} else {
															v658 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v619) {
															v660 = int32(*(*int16)(unsafe.Add(mBase, uint32(v649))))
															v670 = v660
														} else {
															v670 = v619<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v619&int32(63)
														}
														v672 = int32(base.Ui32(int32(base.Ui32(v650)>>(uint(int32(2))%32))+v658) >> (uint(int32(1)) % 32))
														v674 = v22 + int32(6)
														v675 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v682 = base.B2i32(int32(0) <= base.I32_extend16_s(v618))
														if int32(0) <= base.I32_extend16_s(v618) {
															v683 = int32(-8)
														} else {
															v683 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v618) {
															v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674))))
															v695 = v685
														} else {
															v695 = v618<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v618&int32(63)
														}
														v696 = int32(1)
														v697 = int32(base.Ui32(int32(base.Ui32(v675)>>(uint(int32(2))%32))+v683) >> (uint(v696) % 32))
														v703 = v618 & int32(_a_F_width_bucket_numeric_2)
														if v703 == int32(_a_F_width_bucket_numeric_10) {
															v706 = v618 << (uint(v696) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v706 = v703
														}
														if v672 == int32(0) {
															if v697 == int32(0) {
																v757 = int32(0)
															} else {
																if v706 == int32(_a_F_width_bucket_numeric_11) {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v757 = v716
															}
														} else {
															if v621 == int32(_a_F_width_bucket_numeric_10) {
																v723 = v619 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v723 = v621
															}
															if v697 == int32(0) {
																if v723 != 0 {
																	v728 = int32(-1)
																} else {
																	v728 = int32(1)
																}
																v757 = v728
															} else {
																if int32(0) <= base.I32_extend16_s(v619) {
																	v731 = v17 + int32(8)
																} else {
																	v731 = v649
																}
																if int32(0) <= base.I32_extend16_s(v618) {
																	v734 = v22 + int32(8)
																} else {
																	v734 = v674
																}
																if v723 == int32(0) {
																	if v706 == int32(_a_F_width_bucket_numeric_11) {
																		v757 = int32(1)
																	} else {
																		v740 = F_cmp_abs_common(m, v731, v672, v670, v734, v697, v695)
																		mBase = m.M
																		v757 = v740
																	}
																} else {
																	if v706 == int32(0) {
																		v757 = int32(-1)
																	} else {
																		v744 = F_cmp_abs_common(m, v734, v697, v695, v731, v672, v670)
																		mBase = m.M
																		v757 = v744
																	}
																}
															}
														}
													}
												}
												if int32(0) < v757 {
													v761 = F_palloc(m, int32(2))
													mBase = m.M
													v762 = m.ExcPending
													if v762 != 0 {
														return int32(0)
													} else {
														v763 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v763)
														v766 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v766
														v769 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v769
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v761 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v761
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v788 = int32(_a_F_width_bucket_numeric_2)
													v789 = v787 & v788
													if v789 == v788 {
														if v787 != int32(_a_F_width_bucket_numeric_7) {
															if v787 != int32(_a_F_width_bucket_numeric_2) {
																if v786 != int32(_a_F_width_bucket_numeric_9) {
																	v808 = int32(-1)
																} else {
																	v808 = int32(0)
																}
																v925 = v808
															} else {
																v925 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v786 == int32(_a_F_width_bucket_numeric_2) {
																v803 = int32(-1)
															} else {
																v803 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_7))
															}
															v925 = v803
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v786) {
															if v786 == int32(_a_F_width_bucket_numeric_9) {
																v815 = int32(1)
															} else {
																v815 = int32(-1)
															}
															v925 = v815
														} else {
															v817 = v17 + int32(6)
															v818 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v825 = base.B2i32(int32(0) <= base.I32_extend16_s(v787))
															if int32(0) <= base.I32_extend16_s(v787) {
																v826 = int32(-8)
															} else {
																v826 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v787) {
																v828 = int32(*(*int16)(unsafe.Add(mBase, uint32(v817))))
																v838 = v828
															} else {
																v838 = v787<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v787&int32(63)
															}
															v840 = int32(base.Ui32(int32(base.Ui32(v818)>>(uint(int32(2))%32))+v826) >> (uint(int32(1)) % 32))
															v842 = v25 + int32(6)
															v843 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v850 = base.B2i32(int32(0) <= base.I32_extend16_s(v786))
															if int32(0) <= base.I32_extend16_s(v786) {
																v851 = int32(-8)
															} else {
																v851 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v786) {
																v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v842))))
																v863 = v853
															} else {
																v863 = v786<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v786&int32(63)
															}
															v864 = int32(1)
															v865 = int32(base.Ui32(int32(base.Ui32(v843)>>(uint(int32(2))%32))+v851) >> (uint(v864) % 32))
															v871 = v786 & int32(_a_F_width_bucket_numeric_2)
															if v871 == int32(_a_F_width_bucket_numeric_10) {
																v874 = v786 << (uint(v864) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v874 = v871
															}
															if v840 == int32(0) {
																if v865 == int32(0) {
																	v925 = int32(0)
																} else {
																	if v874 == int32(_a_F_width_bucket_numeric_11) {
																		v884 = int32(1)
																	} else {
																		v884 = int32(-1)
																	}
																	v925 = v884
																}
															} else {
																if v789 == int32(_a_F_width_bucket_numeric_10) {
																	v891 = v787 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v891 = v789
																}
																if v865 == int32(0) {
																	if v891 != 0 {
																		v896 = int32(-1)
																	} else {
																		v896 = int32(1)
																	}
																	v925 = v896
																} else {
																	if int32(0) <= base.I32_extend16_s(v787) {
																		v899 = v17 + int32(8)
																	} else {
																		v899 = v817
																	}
																	if int32(0) <= base.I32_extend16_s(v786) {
																		v902 = v25 + int32(8)
																	} else {
																		v902 = v842
																	}
																	if v891 == int32(0) {
																		if v874 == int32(_a_F_width_bucket_numeric_11) {
																			v925 = int32(1)
																		} else {
																			v908 = F_cmp_abs_common(m, v899, v840, v838, v902, v865, v863)
																			mBase = m.M
																			v925 = v908
																		}
																	} else {
																		if v874 == int32(0) {
																			v925 = int32(-1)
																		} else {
																			v912 = F_cmp_abs_common(m, v902, v865, v863, v899, v840, v838)
																			mBase = m.M
																			v925 = v912
																		}
																	}
																}
															}
														}
													}
													if v925 <= int32(0) {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v934 = m.ExcPending
														if v934 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v940 = m.ExcPending
														if v940 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													}
												}
											default:
												v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v286 = int32(_a_F_width_bucket_numeric_2)
												v287 = v285 & v286
												if v287 == v286 {
													if v285 != int32(_a_F_width_bucket_numeric_7) {
														if v285 != int32(_a_F_width_bucket_numeric_2) {
															if v284 != int32(_a_F_width_bucket_numeric_9) {
																v306 = int32(-1)
															} else {
																v306 = int32(0)
															}
															v423 = v306
														} else {
															v423 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v284 == int32(_a_F_width_bucket_numeric_2) {
															v301 = int32(-1)
														} else {
															v301 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_7))
														}
														v423 = v301
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v284) {
														if v284 == int32(_a_F_width_bucket_numeric_9) {
															v313 = int32(1)
														} else {
															v313 = int32(-1)
														}
														v423 = v313
													} else {
														v315 = v17 + int32(6)
														v316 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v285))
														if int32(0) <= base.I32_extend16_s(v285) {
															v324 = int32(-8)
														} else {
															v324 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v285) {
															v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v315))))
															v336 = v326
														} else {
															v336 = v285<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v285&int32(63)
														}
														v338 = int32(base.Ui32(int32(base.Ui32(v316)>>(uint(int32(2))%32))+v324) >> (uint(int32(1)) % 32))
														v340 = v22 + int32(6)
														v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v348 = base.B2i32(int32(0) <= base.I32_extend16_s(v284))
														if int32(0) <= base.I32_extend16_s(v284) {
															v349 = int32(-8)
														} else {
															v349 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v284) {
															v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v340))))
															v361 = v351
														} else {
															v361 = v284<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v284&int32(63)
														}
														v362 = int32(1)
														v363 = int32(base.Ui32(int32(base.Ui32(v341)>>(uint(int32(2))%32))+v349) >> (uint(v362) % 32))
														v369 = v284 & int32(_a_F_width_bucket_numeric_2)
														if v369 == int32(_a_F_width_bucket_numeric_10) {
															v372 = v284 << (uint(v362) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v372 = v369
														}
														if v338 == int32(0) {
															if v363 == int32(0) {
																v423 = int32(0)
															} else {
																if v372 == int32(_a_F_width_bucket_numeric_11) {
																	v382 = int32(1)
																} else {
																	v382 = int32(-1)
																}
																v423 = v382
															}
														} else {
															if v287 == int32(_a_F_width_bucket_numeric_10) {
																v389 = v285 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v389 = v287
															}
															if v363 == int32(0) {
																if v389 != 0 {
																	v394 = int32(-1)
																} else {
																	v394 = int32(1)
																}
																v423 = v394
															} else {
																if int32(0) <= base.I32_extend16_s(v285) {
																	v397 = v17 + int32(8)
																} else {
																	v397 = v315
																}
																if int32(0) <= base.I32_extend16_s(v284) {
																	v400 = v22 + int32(8)
																} else {
																	v400 = v340
																}
																if v389 == int32(0) {
																	if v372 == int32(_a_F_width_bucket_numeric_11) {
																		v423 = int32(1)
																	} else {
																		v406 = F_cmp_abs_common(m, v397, v338, v336, v400, v363, v361)
																		mBase = m.M
																		v423 = v406
																	}
																} else {
																	if v372 == int32(0) {
																		v423 = int32(-1)
																	} else {
																		v410 = F_cmp_abs_common(m, v400, v363, v361, v397, v338, v336)
																		mBase = m.M
																		v423 = v410
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
														return int32(0)
													} else {
														v429 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
														v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
														v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v454 = int32(_a_F_width_bucket_numeric_2)
													v455 = v453 & v454
													if v455 == v454 {
														if v453 != int32(_a_F_width_bucket_numeric_7) {
															if v453 != int32(_a_F_width_bucket_numeric_2) {
																if v452 != int32(_a_F_width_bucket_numeric_9) {
																	v474 = int32(-1)
																} else {
																	v474 = int32(0)
																}
																v591 = v474
															} else {
																v591 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v452 == int32(_a_F_width_bucket_numeric_2) {
																v469 = int32(-1)
															} else {
																v469 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_7))
															}
															v591 = v469
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v452) {
															if v452 == int32(_a_F_width_bucket_numeric_9) {
																v481 = int32(1)
															} else {
																v481 = int32(-1)
															}
															v591 = v481
														} else {
															v483 = v17 + int32(6)
															v484 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v491 = base.B2i32(int32(0) <= base.I32_extend16_s(v453))
															if int32(0) <= base.I32_extend16_s(v453) {
																v492 = int32(-8)
															} else {
																v492 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v453) {
																v494 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483))))
																v504 = v494
															} else {
																v504 = v453<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v453&int32(63)
															}
															v506 = int32(base.Ui32(int32(base.Ui32(v484)>>(uint(int32(2))%32))+v492) >> (uint(int32(1)) % 32))
															v508 = v25 + int32(6)
															v509 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v516 = base.B2i32(int32(0) <= base.I32_extend16_s(v452))
															if int32(0) <= base.I32_extend16_s(v452) {
																v517 = int32(-8)
															} else {
																v517 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v452) {
																v519 = int32(*(*int16)(unsafe.Add(mBase, uint32(v508))))
																v529 = v519
															} else {
																v529 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
															}
															v530 = int32(1)
															v531 = int32(base.Ui32(int32(base.Ui32(v509)>>(uint(int32(2))%32))+v517) >> (uint(v530) % 32))
															v537 = v452 & int32(_a_F_width_bucket_numeric_2)
															if v537 == int32(_a_F_width_bucket_numeric_10) {
																v540 = v452 << (uint(v530) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v540 = v537
															}
															if v506 == int32(0) {
																if v531 == int32(0) {
																	v591 = int32(0)
																} else {
																	if v540 == int32(_a_F_width_bucket_numeric_11) {
																		v550 = int32(1)
																	} else {
																		v550 = int32(-1)
																	}
																	v591 = v550
																}
															} else {
																if v455 == int32(_a_F_width_bucket_numeric_10) {
																	v557 = v453 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v557 = v455
																}
																if v531 == int32(0) {
																	if v557 != 0 {
																		v562 = int32(-1)
																	} else {
																		v562 = int32(1)
																	}
																	v591 = v562
																} else {
																	if int32(0) <= base.I32_extend16_s(v453) {
																		v565 = v17 + int32(8)
																	} else {
																		v565 = v483
																	}
																	if int32(0) <= base.I32_extend16_s(v452) {
																		v568 = v25 + int32(8)
																	} else {
																		v568 = v508
																	}
																	if v557 == int32(0) {
																		if v540 == int32(_a_F_width_bucket_numeric_11) {
																			v591 = int32(1)
																		} else {
																			v574 = F_cmp_abs_common(m, v565, v506, v504, v568, v531, v529)
																			mBase = m.M
																			v591 = v574
																		}
																	} else {
																		if v540 == int32(0) {
																			v591 = int32(-1)
																		} else {
																			v578 = F_cmp_abs_common(m, v568, v531, v529, v565, v506, v504)
																			mBase = m.M
																			v591 = v578
																		}
																	}
																}
															}
														}
													}
													if int32(0) <= v591 {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v600 = m.ExcPending
														if v600 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v606 = m.ExcPending
														if v606 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
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
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v62
								v69 = F_palloc(m, int32(12))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v69
									v72 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v72)
									*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
									v80 = v72
									v87 = v69 + int32(12)
									v89 = base.I64_extend_i32_u(v27)
									for {
										v92 = v87 - int32(2)
										v94 = base.I64_div_u_s(v89, int64(10000))
										v97 = v94*int64(55536) + v89
										*(*uint16)(unsafe.Add(mBase, uint32(v92))) = uint16(v97)
										v100 = v80 + int32(1)
										if base.Ui64(int64(9999)) < base.Ui64(v89) {
											v80 = v100
											v87 = v92
											v89 = v94
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v80
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v100
									*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v92
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
									v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
									v119 = int32(_a_F_width_bucket_numeric_2)
									v120 = v118 & v119
									if v120 == v119 {
										if v118 != int32(_a_F_width_bucket_numeric_7) {
											if v118 != int32(_a_F_width_bucket_numeric_2) {
												if v117 != int32(_a_F_width_bucket_numeric_9) {
													v139 = int32(-1)
												} else {
													v139 = int32(0)
												}
												v256 = v139
											} else {
												v256 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_2))
											}
										} else {
											if v117 == int32(_a_F_width_bucket_numeric_2) {
												v134 = int32(-1)
											} else {
												v134 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_7))
											}
											v256 = v134
										}
									} else {
										if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v117) {
											if v117 == int32(_a_F_width_bucket_numeric_9) {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v256 = v146
										} else {
											v148 = v22 + int32(6)
											v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
											v156 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
											if int32(0) <= base.I32_extend16_s(v118) {
												v157 = int32(-8)
											} else {
												v157 = int32(-6)
											}
											if int32(0) <= base.I32_extend16_s(v118) {
												v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
												v169 = v159
											} else {
												v169 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
											}
											v171 = int32(base.Ui32(int32(base.Ui32(v149)>>(uint(int32(2))%32))+v157) >> (uint(int32(1)) % 32))
											v173 = v25 + int32(6)
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
											v181 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
											if int32(0) <= base.I32_extend16_s(v117) {
												v182 = int32(-8)
											} else {
												v182 = int32(-6)
											}
											if int32(0) <= base.I32_extend16_s(v117) {
												v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173))))
												v194 = v184
											} else {
												v194 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
											}
											v195 = int32(1)
											v196 = int32(base.Ui32(int32(base.Ui32(v174)>>(uint(int32(2))%32))+v182) >> (uint(v195) % 32))
											v202 = v117 & int32(_a_F_width_bucket_numeric_2)
											if v202 == int32(_a_F_width_bucket_numeric_10) {
												v205 = v117 << (uint(v195) % 32) & int32(_a_F_width_bucket_numeric_11)
											} else {
												v205 = v202
											}
											if v171 == int32(0) {
												if v196 == int32(0) {
													v256 = int32(0)
												} else {
													if v205 == int32(_a_F_width_bucket_numeric_11) {
														v215 = int32(1)
													} else {
														v215 = int32(-1)
													}
													v256 = v215
												}
											} else {
												if v120 == int32(_a_F_width_bucket_numeric_10) {
													v222 = v118 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
												} else {
													v222 = v120
												}
												if v196 == int32(0) {
													if v222 != 0 {
														v227 = int32(-1)
													} else {
														v227 = int32(1)
													}
													v256 = v227
												} else {
													if int32(0) <= base.I32_extend16_s(v118) {
														v230 = v22 + int32(8)
													} else {
														v230 = v148
													}
													if int32(0) <= base.I32_extend16_s(v117) {
														v233 = v25 + int32(8)
													} else {
														v233 = v173
													}
													if v222 == int32(0) {
														if v205 == int32(_a_F_width_bucket_numeric_11) {
															v256 = int32(1)
														} else {
															v239 = F_cmp_abs_common(m, v230, v171, v169, v233, v196, v194)
															mBase = m.M
															v256 = v239
														}
													} else {
														if v205 == int32(0) {
															v256 = int32(-1)
														} else {
															v243 = F_cmp_abs_common(m, v233, v196, v194, v230, v171, v169)
															mBase = m.M
															v256 = v243
														}
													}
												}
											}
										}
									}
									switch v256 {
									case 0:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v260 = m.ExcPending
										if v260 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v263 = m.ExcPending
											if v263 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
												mBase = m.M
												v267 = m.ExcPending
												if v267 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2010), int32(_a_F_width_bucket_numeric_5))
													mBase = m.M
													v272 = m.ExcPending
													if v272 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									case 1:
										v618 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										v620 = int32(_a_F_width_bucket_numeric_2)
										v621 = v619 & v620
										if v621 == v620 {
											if v619 != int32(_a_F_width_bucket_numeric_7) {
												if v619 != int32(_a_F_width_bucket_numeric_2) {
													if v618 != int32(_a_F_width_bucket_numeric_9) {
														v640 = int32(-1)
													} else {
														v640 = int32(0)
													}
													v757 = v640
												} else {
													v757 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_2))
												}
											} else {
												if v618 == int32(_a_F_width_bucket_numeric_2) {
													v635 = int32(-1)
												} else {
													v635 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_7))
												}
												v757 = v635
											}
										} else {
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v618) {
												if v618 == int32(_a_F_width_bucket_numeric_9) {
													v647 = int32(1)
												} else {
													v647 = int32(-1)
												}
												v757 = v647
											} else {
												v649 = v17 + int32(6)
												v650 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v657 = base.B2i32(int32(0) <= base.I32_extend16_s(v619))
												if int32(0) <= base.I32_extend16_s(v619) {
													v658 = int32(-8)
												} else {
													v658 = int32(-6)
												}
												if int32(0) <= base.I32_extend16_s(v619) {
													v660 = int32(*(*int16)(unsafe.Add(mBase, uint32(v649))))
													v670 = v660
												} else {
													v670 = v619<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v619&int32(63)
												}
												v672 = int32(base.Ui32(int32(base.Ui32(v650)>>(uint(int32(2))%32))+v658) >> (uint(int32(1)) % 32))
												v674 = v22 + int32(6)
												v675 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												v682 = base.B2i32(int32(0) <= base.I32_extend16_s(v618))
												if int32(0) <= base.I32_extend16_s(v618) {
													v683 = int32(-8)
												} else {
													v683 = int32(-6)
												}
												if int32(0) <= base.I32_extend16_s(v618) {
													v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674))))
													v695 = v685
												} else {
													v695 = v618<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v618&int32(63)
												}
												v696 = int32(1)
												v697 = int32(base.Ui32(int32(base.Ui32(v675)>>(uint(int32(2))%32))+v683) >> (uint(v696) % 32))
												v703 = v618 & int32(_a_F_width_bucket_numeric_2)
												if v703 == int32(_a_F_width_bucket_numeric_10) {
													v706 = v618 << (uint(v696) % 32) & int32(_a_F_width_bucket_numeric_11)
												} else {
													v706 = v703
												}
												if v672 == int32(0) {
													if v697 == int32(0) {
														v757 = int32(0)
													} else {
														if v706 == int32(_a_F_width_bucket_numeric_11) {
															v716 = int32(1)
														} else {
															v716 = int32(-1)
														}
														v757 = v716
													}
												} else {
													if v621 == int32(_a_F_width_bucket_numeric_10) {
														v723 = v619 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v723 = v621
													}
													if v697 == int32(0) {
														if v723 != 0 {
															v728 = int32(-1)
														} else {
															v728 = int32(1)
														}
														v757 = v728
													} else {
														if int32(0) <= base.I32_extend16_s(v619) {
															v731 = v17 + int32(8)
														} else {
															v731 = v649
														}
														if int32(0) <= base.I32_extend16_s(v618) {
															v734 = v22 + int32(8)
														} else {
															v734 = v674
														}
														if v723 == int32(0) {
															if v706 == int32(_a_F_width_bucket_numeric_11) {
																v757 = int32(1)
															} else {
																v740 = F_cmp_abs_common(m, v731, v672, v670, v734, v697, v695)
																mBase = m.M
																v757 = v740
															}
														} else {
															if v706 == int32(0) {
																v757 = int32(-1)
															} else {
																v744 = F_cmp_abs_common(m, v734, v697, v695, v731, v672, v670)
																mBase = m.M
																v757 = v744
															}
														}
													}
												}
											}
										}
										if int32(0) < v757 {
											v761 = F_palloc(m, int32(2))
											mBase = m.M
											v762 = m.ExcPending
											if v762 != 0 {
												return int32(0)
											} else {
												v763 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v763)
												v766 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v766
												v769 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v769
												*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v761 + int32(2)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v761
												v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v947 = m.ExcPending
												if v947 != 0 {
													return int32(0)
												} else {
													if v946 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1019 = m.ExcPending
														if v1019 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1022 = m.ExcPending
															if v1022 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1031 = m.ExcPending
																	if v1031 != 0 {
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
														v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															F_pfree(m, v69)
															mBase = m.M
															v956 = m.ExcPending
															if v956 != 0 {
																return int32(0)
															} else {
																v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v957 != 0 {
																	F_pfree(m, v957)
																	mBase = m.M
																	v959 = m.ExcPending
																	if v959 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return base.I32_wrap_i64(v950)
																}
															}
														}
													}
												}
											}
										} else {
											v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											v788 = int32(_a_F_width_bucket_numeric_2)
											v789 = v787 & v788
											if v789 == v788 {
												if v787 != int32(_a_F_width_bucket_numeric_7) {
													if v787 != int32(_a_F_width_bucket_numeric_2) {
														if v786 != int32(_a_F_width_bucket_numeric_9) {
															v808 = int32(-1)
														} else {
															v808 = int32(0)
														}
														v925 = v808
													} else {
														v925 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_2))
													}
												} else {
													if v786 == int32(_a_F_width_bucket_numeric_2) {
														v803 = int32(-1)
													} else {
														v803 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_7))
													}
													v925 = v803
												}
											} else {
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v786) {
													if v786 == int32(_a_F_width_bucket_numeric_9) {
														v815 = int32(1)
													} else {
														v815 = int32(-1)
													}
													v925 = v815
												} else {
													v817 = v17 + int32(6)
													v818 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v825 = base.B2i32(int32(0) <= base.I32_extend16_s(v787))
													if int32(0) <= base.I32_extend16_s(v787) {
														v826 = int32(-8)
													} else {
														v826 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v787) {
														v828 = int32(*(*int16)(unsafe.Add(mBase, uint32(v817))))
														v838 = v828
													} else {
														v838 = v787<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v787&int32(63)
													}
													v840 = int32(base.Ui32(int32(base.Ui32(v818)>>(uint(int32(2))%32))+v826) >> (uint(int32(1)) % 32))
													v842 = v25 + int32(6)
													v843 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													v850 = base.B2i32(int32(0) <= base.I32_extend16_s(v786))
													if int32(0) <= base.I32_extend16_s(v786) {
														v851 = int32(-8)
													} else {
														v851 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v786) {
														v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v842))))
														v863 = v853
													} else {
														v863 = v786<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v786&int32(63)
													}
													v864 = int32(1)
													v865 = int32(base.Ui32(int32(base.Ui32(v843)>>(uint(int32(2))%32))+v851) >> (uint(v864) % 32))
													v871 = v786 & int32(_a_F_width_bucket_numeric_2)
													if v871 == int32(_a_F_width_bucket_numeric_10) {
														v874 = v786 << (uint(v864) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v874 = v871
													}
													if v840 == int32(0) {
														if v865 == int32(0) {
															v925 = int32(0)
														} else {
															if v874 == int32(_a_F_width_bucket_numeric_11) {
																v884 = int32(1)
															} else {
																v884 = int32(-1)
															}
															v925 = v884
														}
													} else {
														if v789 == int32(_a_F_width_bucket_numeric_10) {
															v891 = v787 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v891 = v789
														}
														if v865 == int32(0) {
															if v891 != 0 {
																v896 = int32(-1)
															} else {
																v896 = int32(1)
															}
															v925 = v896
														} else {
															if int32(0) <= base.I32_extend16_s(v787) {
																v899 = v17 + int32(8)
															} else {
																v899 = v817
															}
															if int32(0) <= base.I32_extend16_s(v786) {
																v902 = v25 + int32(8)
															} else {
																v902 = v842
															}
															if v891 == int32(0) {
																if v874 == int32(_a_F_width_bucket_numeric_11) {
																	v925 = int32(1)
																} else {
																	v908 = F_cmp_abs_common(m, v899, v840, v838, v902, v865, v863)
																	mBase = m.M
																	v925 = v908
																}
															} else {
																if v874 == int32(0) {
																	v925 = int32(-1)
																} else {
																	v912 = F_cmp_abs_common(m, v902, v865, v863, v899, v840, v838)
																	mBase = m.M
																	v925 = v912
																}
															}
														}
													}
												}
											}
											if v925 <= int32(0) {
												F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
												mBase = m.M
												v934 = m.ExcPending
												if v934 != 0 {
													return int32(0)
												} else {
													v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v947 = m.ExcPending
													if v947 != 0 {
														return int32(0)
													} else {
														if v946 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																F_pfree(m, v69)
																mBase = m.M
																v956 = m.ExcPending
																if v956 != 0 {
																	return int32(0)
																} else {
																	v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v957 != 0 {
																		F_pfree(m, v957)
																		mBase = m.M
																		v959 = m.ExcPending
																		if v959 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
																	}
																}
															}
														}
													}
												}
											} else {
												F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
												mBase = m.M
												v940 = m.ExcPending
												if v940 != 0 {
													return int32(0)
												} else {
													v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v947 = m.ExcPending
													if v947 != 0 {
														return int32(0)
													} else {
														if v946 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																F_pfree(m, v69)
																mBase = m.M
																v956 = m.ExcPending
																if v956 != 0 {
																	return int32(0)
																} else {
																	v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v957 != 0 {
																		F_pfree(m, v957)
																		mBase = m.M
																		v959 = m.ExcPending
																		if v959 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
																	}
																}
															}
														}
													}
												}
											}
										}
									default:
										v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										v286 = int32(_a_F_width_bucket_numeric_2)
										v287 = v285 & v286
										if v287 == v286 {
											if v285 != int32(_a_F_width_bucket_numeric_7) {
												if v285 != int32(_a_F_width_bucket_numeric_2) {
													if v284 != int32(_a_F_width_bucket_numeric_9) {
														v306 = int32(-1)
													} else {
														v306 = int32(0)
													}
													v423 = v306
												} else {
													v423 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_2))
												}
											} else {
												if v284 == int32(_a_F_width_bucket_numeric_2) {
													v301 = int32(-1)
												} else {
													v301 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_7))
												}
												v423 = v301
											}
										} else {
											if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v284) {
												if v284 == int32(_a_F_width_bucket_numeric_9) {
													v313 = int32(1)
												} else {
													v313 = int32(-1)
												}
												v423 = v313
											} else {
												v315 = v17 + int32(6)
												v316 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v285))
												if int32(0) <= base.I32_extend16_s(v285) {
													v324 = int32(-8)
												} else {
													v324 = int32(-6)
												}
												if int32(0) <= base.I32_extend16_s(v285) {
													v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v315))))
													v336 = v326
												} else {
													v336 = v285<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v285&int32(63)
												}
												v338 = int32(base.Ui32(int32(base.Ui32(v316)>>(uint(int32(2))%32))+v324) >> (uint(int32(1)) % 32))
												v340 = v22 + int32(6)
												v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												v348 = base.B2i32(int32(0) <= base.I32_extend16_s(v284))
												if int32(0) <= base.I32_extend16_s(v284) {
													v349 = int32(-8)
												} else {
													v349 = int32(-6)
												}
												if int32(0) <= base.I32_extend16_s(v284) {
													v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v340))))
													v361 = v351
												} else {
													v361 = v284<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v284&int32(63)
												}
												v362 = int32(1)
												v363 = int32(base.Ui32(int32(base.Ui32(v341)>>(uint(int32(2))%32))+v349) >> (uint(v362) % 32))
												v369 = v284 & int32(_a_F_width_bucket_numeric_2)
												if v369 == int32(_a_F_width_bucket_numeric_10) {
													v372 = v284 << (uint(v362) % 32) & int32(_a_F_width_bucket_numeric_11)
												} else {
													v372 = v369
												}
												if v338 == int32(0) {
													if v363 == int32(0) {
														v423 = int32(0)
													} else {
														if v372 == int32(_a_F_width_bucket_numeric_11) {
															v382 = int32(1)
														} else {
															v382 = int32(-1)
														}
														v423 = v382
													}
												} else {
													if v287 == int32(_a_F_width_bucket_numeric_10) {
														v389 = v285 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v389 = v287
													}
													if v363 == int32(0) {
														if v389 != 0 {
															v394 = int32(-1)
														} else {
															v394 = int32(1)
														}
														v423 = v394
													} else {
														if int32(0) <= base.I32_extend16_s(v285) {
															v397 = v17 + int32(8)
														} else {
															v397 = v315
														}
														if int32(0) <= base.I32_extend16_s(v284) {
															v400 = v22 + int32(8)
														} else {
															v400 = v340
														}
														if v389 == int32(0) {
															if v372 == int32(_a_F_width_bucket_numeric_11) {
																v423 = int32(1)
															} else {
																v406 = F_cmp_abs_common(m, v397, v338, v336, v400, v363, v361)
																mBase = m.M
																v423 = v406
															}
														} else {
															if v372 == int32(0) {
																v423 = int32(-1)
															} else {
																v410 = F_cmp_abs_common(m, v400, v363, v361, v397, v338, v336)
																mBase = m.M
																v423 = v410
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
												return int32(0)
											} else {
												v429 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
												v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
												v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
												*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
												v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v947 = m.ExcPending
												if v947 != 0 {
													return int32(0)
												} else {
													if v946 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1019 = m.ExcPending
														if v1019 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1022 = m.ExcPending
															if v1022 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																mBase = m.M
																v1026 = m.ExcPending
																if v1026 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																	mBase = m.M
																	v1031 = m.ExcPending
																	if v1031 != 0 {
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
														v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															F_pfree(m, v69)
															mBase = m.M
															v956 = m.ExcPending
															if v956 != 0 {
																return int32(0)
															} else {
																v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v957 != 0 {
																	F_pfree(m, v957)
																	mBase = m.M
																	v959 = m.ExcPending
																	if v959 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return base.I32_wrap_i64(v950)
																}
															}
														}
													}
												}
											}
										} else {
											v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											v454 = int32(_a_F_width_bucket_numeric_2)
											v455 = v453 & v454
											if v455 == v454 {
												if v453 != int32(_a_F_width_bucket_numeric_7) {
													if v453 != int32(_a_F_width_bucket_numeric_2) {
														if v452 != int32(_a_F_width_bucket_numeric_9) {
															v474 = int32(-1)
														} else {
															v474 = int32(0)
														}
														v591 = v474
													} else {
														v591 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_2))
													}
												} else {
													if v452 == int32(_a_F_width_bucket_numeric_2) {
														v469 = int32(-1)
													} else {
														v469 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_7))
													}
													v591 = v469
												}
											} else {
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v452) {
													if v452 == int32(_a_F_width_bucket_numeric_9) {
														v481 = int32(1)
													} else {
														v481 = int32(-1)
													}
													v591 = v481
												} else {
													v483 = v17 + int32(6)
													v484 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v491 = base.B2i32(int32(0) <= base.I32_extend16_s(v453))
													if int32(0) <= base.I32_extend16_s(v453) {
														v492 = int32(-8)
													} else {
														v492 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v453) {
														v494 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483))))
														v504 = v494
													} else {
														v504 = v453<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v453&int32(63)
													}
													v506 = int32(base.Ui32(int32(base.Ui32(v484)>>(uint(int32(2))%32))+v492) >> (uint(int32(1)) % 32))
													v508 = v25 + int32(6)
													v509 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													v516 = base.B2i32(int32(0) <= base.I32_extend16_s(v452))
													if int32(0) <= base.I32_extend16_s(v452) {
														v517 = int32(-8)
													} else {
														v517 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v452) {
														v519 = int32(*(*int16)(unsafe.Add(mBase, uint32(v508))))
														v529 = v519
													} else {
														v529 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
													}
													v530 = int32(1)
													v531 = int32(base.Ui32(int32(base.Ui32(v509)>>(uint(int32(2))%32))+v517) >> (uint(v530) % 32))
													v537 = v452 & int32(_a_F_width_bucket_numeric_2)
													if v537 == int32(_a_F_width_bucket_numeric_10) {
														v540 = v452 << (uint(v530) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v540 = v537
													}
													if v506 == int32(0) {
														if v531 == int32(0) {
															v591 = int32(0)
														} else {
															if v540 == int32(_a_F_width_bucket_numeric_11) {
																v550 = int32(1)
															} else {
																v550 = int32(-1)
															}
															v591 = v550
														}
													} else {
														if v455 == int32(_a_F_width_bucket_numeric_10) {
															v557 = v453 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v557 = v455
														}
														if v531 == int32(0) {
															if v557 != 0 {
																v562 = int32(-1)
															} else {
																v562 = int32(1)
															}
															v591 = v562
														} else {
															if int32(0) <= base.I32_extend16_s(v453) {
																v565 = v17 + int32(8)
															} else {
																v565 = v483
															}
															if int32(0) <= base.I32_extend16_s(v452) {
																v568 = v25 + int32(8)
															} else {
																v568 = v508
															}
															if v557 == int32(0) {
																if v540 == int32(_a_F_width_bucket_numeric_11) {
																	v591 = int32(1)
																} else {
																	v574 = F_cmp_abs_common(m, v565, v506, v504, v568, v531, v529)
																	mBase = m.M
																	v591 = v574
																}
															} else {
																if v540 == int32(0) {
																	v591 = int32(-1)
																} else {
																	v578 = F_cmp_abs_common(m, v568, v531, v529, v565, v506, v504)
																	mBase = m.M
																	v591 = v578
																}
															}
														}
													}
												}
											}
											if int32(0) <= v591 {
												F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
												mBase = m.M
												v600 = m.ExcPending
												if v600 != 0 {
													return int32(0)
												} else {
													v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v947 = m.ExcPending
													if v947 != 0 {
														return int32(0)
													} else {
														if v946 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																F_pfree(m, v69)
																mBase = m.M
																v956 = m.ExcPending
																if v956 != 0 {
																	return int32(0)
																} else {
																	v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v957 != 0 {
																		F_pfree(m, v957)
																		mBase = m.M
																		v959 = m.ExcPending
																		if v959 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
																	}
																}
															}
														}
													}
												}
											} else {
												F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
												mBase = m.M
												v606 = m.ExcPending
												if v606 != 0 {
													return int32(0)
												} else {
													v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v947 = m.ExcPending
													if v947 != 0 {
														return int32(0)
													} else {
														if v946 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1019 = m.ExcPending
															if v1019 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1022 = m.ExcPending
																if v1022 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																	mBase = m.M
																	v1026 = m.ExcPending
																	if v1026 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																		mBase = m.M
																		v1031 = m.ExcPending
																		if v1031 != 0 {
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
															v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																F_pfree(m, v69)
																mBase = m.M
																v956 = m.ExcPending
																if v956 != 0 {
																	return int32(0)
																} else {
																	v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v957 != 0 {
																		F_pfree(m, v957)
																		mBase = m.M
																		v959 = m.ExcPending
																		if v959 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v950)
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
						if v30 == int32(_a_F_width_bucket_numeric_2) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v986 = m.ExcPending
							if v986 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(386138242))
								mBase = m.M
								v989 = m.ExcPending
								if v989 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
									mBase = m.M
									v993 = m.ExcPending
									if v993 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
										mBase = m.M
										v998 = m.ExcPending
										if v998 != 0 {
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
							v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
							v42 = v41
							if v42&int32(_a_F_width_bucket_numeric_1) == int32(_a_F_width_bucket_numeric_2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v986 = m.ExcPending
								if v986 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v989 = m.ExcPending
									if v989 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
										mBase = m.M
										v993 = m.ExcPending
										if v993 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
											mBase = m.M
											v998 = m.ExcPending
											if v998 != 0 {
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
								v47 = v42
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
								if v48 == int32(_a_F_width_bucket_numeric_2) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v986 = m.ExcPending
									if v986 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v989 = m.ExcPending
										if v989 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_width_bucket_numeric_3), int32(0))
											mBase = m.M
											v993 = m.ExcPending
											if v993 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1991), int32(_a_F_width_bucket_numeric_5))
												mBase = m.M
												v998 = m.ExcPending
												if v998 != 0 {
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
									v51 = int32(_a_F_width_bucket_numeric_6)
									v53 = int32(_a_F_width_bucket_numeric_7)
									if base.B2i32(v47&v51 == v53)|base.B2i32(v48&v51 == v53) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1002 = m.ExcPending
										if v1002 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1005 = m.ExcPending
											if v1005 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_width_bucket_numeric_8), int32(0))
												mBase = m.M
												v1009 = m.ExcPending
												if v1009 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1996), int32(_a_F_width_bucket_numeric_5))
													mBase = m.M
													v1014 = m.ExcPending
													if v1014 != 0 {
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
										v62 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v62
										*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v62
										v69 = F_palloc(m, int32(12))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v69
											v72 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v72)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
											v80 = v72
											v87 = v69 + int32(12)
											v89 = base.I64_extend_i32_u(v27)
											for {
												v92 = v87 - int32(2)
												v94 = base.I64_div_u_s(v89, int64(10000))
												v97 = v94*int64(55536) + v89
												*(*uint16)(unsafe.Add(mBase, uint32(v92))) = uint16(v97)
												v100 = v80 + int32(1)
												if base.Ui64(int64(9999)) < base.Ui64(v89) {
													v80 = v100
													v87 = v92
													v89 = v94
													continue
												} else {
													break
												}
												break
											}
											*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v80
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v100
											*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v92
											v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
											v119 = int32(_a_F_width_bucket_numeric_2)
											v120 = v118 & v119
											if v120 == v119 {
												if v118 != int32(_a_F_width_bucket_numeric_7) {
													if v118 != int32(_a_F_width_bucket_numeric_2) {
														if v117 != int32(_a_F_width_bucket_numeric_9) {
															v139 = int32(-1)
														} else {
															v139 = int32(0)
														}
														v256 = v139
													} else {
														v256 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_2))
													}
												} else {
													if v117 == int32(_a_F_width_bucket_numeric_2) {
														v134 = int32(-1)
													} else {
														v134 = base.B2i32(v117 != int32(_a_F_width_bucket_numeric_7))
													}
													v256 = v134
												}
											} else {
												if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v117) {
													if v117 == int32(_a_F_width_bucket_numeric_9) {
														v146 = int32(1)
													} else {
														v146 = int32(-1)
													}
													v256 = v146
												} else {
													v148 = v22 + int32(6)
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v156 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
													if int32(0) <= base.I32_extend16_s(v118) {
														v157 = int32(-8)
													} else {
														v157 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v118) {
														v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
														v169 = v159
													} else {
														v169 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
													}
													v171 = int32(base.Ui32(int32(base.Ui32(v149)>>(uint(int32(2))%32))+v157) >> (uint(int32(1)) % 32))
													v173 = v25 + int32(6)
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													v181 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
													if int32(0) <= base.I32_extend16_s(v117) {
														v182 = int32(-8)
													} else {
														v182 = int32(-6)
													}
													if int32(0) <= base.I32_extend16_s(v117) {
														v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173))))
														v194 = v184
													} else {
														v194 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
													}
													v195 = int32(1)
													v196 = int32(base.Ui32(int32(base.Ui32(v174)>>(uint(int32(2))%32))+v182) >> (uint(v195) % 32))
													v202 = v117 & int32(_a_F_width_bucket_numeric_2)
													if v202 == int32(_a_F_width_bucket_numeric_10) {
														v205 = v117 << (uint(v195) % 32) & int32(_a_F_width_bucket_numeric_11)
													} else {
														v205 = v202
													}
													if v171 == int32(0) {
														if v196 == int32(0) {
															v256 = int32(0)
														} else {
															if v205 == int32(_a_F_width_bucket_numeric_11) {
																v215 = int32(1)
															} else {
																v215 = int32(-1)
															}
															v256 = v215
														}
													} else {
														if v120 == int32(_a_F_width_bucket_numeric_10) {
															v222 = v118 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v222 = v120
														}
														if v196 == int32(0) {
															if v222 != 0 {
																v227 = int32(-1)
															} else {
																v227 = int32(1)
															}
															v256 = v227
														} else {
															if int32(0) <= base.I32_extend16_s(v118) {
																v230 = v22 + int32(8)
															} else {
																v230 = v148
															}
															if int32(0) <= base.I32_extend16_s(v117) {
																v233 = v25 + int32(8)
															} else {
																v233 = v173
															}
															if v222 == int32(0) {
																if v205 == int32(_a_F_width_bucket_numeric_11) {
																	v256 = int32(1)
																} else {
																	v239 = F_cmp_abs_common(m, v230, v171, v169, v233, v196, v194)
																	mBase = m.M
																	v256 = v239
																}
															} else {
																if v205 == int32(0) {
																	v256 = int32(-1)
																} else {
																	v243 = F_cmp_abs_common(m, v233, v196, v194, v230, v171, v169)
																	mBase = m.M
																	v256 = v243
																}
															}
														}
													}
												}
											}
											switch v256 {
											case 0:
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(386138242))
													mBase = m.M
													v263 = m.ExcPending
													if v263 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_width_bucket_numeric_12), int32(0))
														mBase = m.M
														v267 = m.ExcPending
														if v267 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2010), int32(_a_F_width_bucket_numeric_5))
															mBase = m.M
															v272 = m.ExcPending
															if v272 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											case 1:
												v618 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v620 = int32(_a_F_width_bucket_numeric_2)
												v621 = v619 & v620
												if v621 == v620 {
													if v619 != int32(_a_F_width_bucket_numeric_7) {
														if v619 != int32(_a_F_width_bucket_numeric_2) {
															if v618 != int32(_a_F_width_bucket_numeric_9) {
																v640 = int32(-1)
															} else {
																v640 = int32(0)
															}
															v757 = v640
														} else {
															v757 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v618 == int32(_a_F_width_bucket_numeric_2) {
															v635 = int32(-1)
														} else {
															v635 = base.B2i32(v618 != int32(_a_F_width_bucket_numeric_7))
														}
														v757 = v635
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v618) {
														if v618 == int32(_a_F_width_bucket_numeric_9) {
															v647 = int32(1)
														} else {
															v647 = int32(-1)
														}
														v757 = v647
													} else {
														v649 = v17 + int32(6)
														v650 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v657 = base.B2i32(int32(0) <= base.I32_extend16_s(v619))
														if int32(0) <= base.I32_extend16_s(v619) {
															v658 = int32(-8)
														} else {
															v658 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v619) {
															v660 = int32(*(*int16)(unsafe.Add(mBase, uint32(v649))))
															v670 = v660
														} else {
															v670 = v619<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v619&int32(63)
														}
														v672 = int32(base.Ui32(int32(base.Ui32(v650)>>(uint(int32(2))%32))+v658) >> (uint(int32(1)) % 32))
														v674 = v22 + int32(6)
														v675 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v682 = base.B2i32(int32(0) <= base.I32_extend16_s(v618))
														if int32(0) <= base.I32_extend16_s(v618) {
															v683 = int32(-8)
														} else {
															v683 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v618) {
															v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v674))))
															v695 = v685
														} else {
															v695 = v618<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v618&int32(63)
														}
														v696 = int32(1)
														v697 = int32(base.Ui32(int32(base.Ui32(v675)>>(uint(int32(2))%32))+v683) >> (uint(v696) % 32))
														v703 = v618 & int32(_a_F_width_bucket_numeric_2)
														if v703 == int32(_a_F_width_bucket_numeric_10) {
															v706 = v618 << (uint(v696) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v706 = v703
														}
														if v672 == int32(0) {
															if v697 == int32(0) {
																v757 = int32(0)
															} else {
																if v706 == int32(_a_F_width_bucket_numeric_11) {
																	v716 = int32(1)
																} else {
																	v716 = int32(-1)
																}
																v757 = v716
															}
														} else {
															if v621 == int32(_a_F_width_bucket_numeric_10) {
																v723 = v619 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v723 = v621
															}
															if v697 == int32(0) {
																if v723 != 0 {
																	v728 = int32(-1)
																} else {
																	v728 = int32(1)
																}
																v757 = v728
															} else {
																if int32(0) <= base.I32_extend16_s(v619) {
																	v731 = v17 + int32(8)
																} else {
																	v731 = v649
																}
																if int32(0) <= base.I32_extend16_s(v618) {
																	v734 = v22 + int32(8)
																} else {
																	v734 = v674
																}
																if v723 == int32(0) {
																	if v706 == int32(_a_F_width_bucket_numeric_11) {
																		v757 = int32(1)
																	} else {
																		v740 = F_cmp_abs_common(m, v731, v672, v670, v734, v697, v695)
																		mBase = m.M
																		v757 = v740
																	}
																} else {
																	if v706 == int32(0) {
																		v757 = int32(-1)
																	} else {
																		v744 = F_cmp_abs_common(m, v734, v697, v695, v731, v672, v670)
																		mBase = m.M
																		v757 = v744
																	}
																}
															}
														}
													}
												}
												if int32(0) < v757 {
													v761 = F_palloc(m, int32(2))
													mBase = m.M
													v762 = m.ExcPending
													if v762 != 0 {
														return int32(0)
													} else {
														v763 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v761))) = uint16(v763)
														v766 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v766
														v769 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v769
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v761 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v761
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v788 = int32(_a_F_width_bucket_numeric_2)
													v789 = v787 & v788
													if v789 == v788 {
														if v787 != int32(_a_F_width_bucket_numeric_7) {
															if v787 != int32(_a_F_width_bucket_numeric_2) {
																if v786 != int32(_a_F_width_bucket_numeric_9) {
																	v808 = int32(-1)
																} else {
																	v808 = int32(0)
																}
																v925 = v808
															} else {
																v925 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v786 == int32(_a_F_width_bucket_numeric_2) {
																v803 = int32(-1)
															} else {
																v803 = base.B2i32(v786 != int32(_a_F_width_bucket_numeric_7))
															}
															v925 = v803
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v786) {
															if v786 == int32(_a_F_width_bucket_numeric_9) {
																v815 = int32(1)
															} else {
																v815 = int32(-1)
															}
															v925 = v815
														} else {
															v817 = v17 + int32(6)
															v818 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v825 = base.B2i32(int32(0) <= base.I32_extend16_s(v787))
															if int32(0) <= base.I32_extend16_s(v787) {
																v826 = int32(-8)
															} else {
																v826 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v787) {
																v828 = int32(*(*int16)(unsafe.Add(mBase, uint32(v817))))
																v838 = v828
															} else {
																v838 = v787<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v787&int32(63)
															}
															v840 = int32(base.Ui32(int32(base.Ui32(v818)>>(uint(int32(2))%32))+v826) >> (uint(int32(1)) % 32))
															v842 = v25 + int32(6)
															v843 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v850 = base.B2i32(int32(0) <= base.I32_extend16_s(v786))
															if int32(0) <= base.I32_extend16_s(v786) {
																v851 = int32(-8)
															} else {
																v851 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v786) {
																v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v842))))
																v863 = v853
															} else {
																v863 = v786<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v786&int32(63)
															}
															v864 = int32(1)
															v865 = int32(base.Ui32(int32(base.Ui32(v843)>>(uint(int32(2))%32))+v851) >> (uint(v864) % 32))
															v871 = v786 & int32(_a_F_width_bucket_numeric_2)
															if v871 == int32(_a_F_width_bucket_numeric_10) {
																v874 = v786 << (uint(v864) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v874 = v871
															}
															if v840 == int32(0) {
																if v865 == int32(0) {
																	v925 = int32(0)
																} else {
																	if v874 == int32(_a_F_width_bucket_numeric_11) {
																		v884 = int32(1)
																	} else {
																		v884 = int32(-1)
																	}
																	v925 = v884
																}
															} else {
																if v789 == int32(_a_F_width_bucket_numeric_10) {
																	v891 = v787 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v891 = v789
																}
																if v865 == int32(0) {
																	if v891 != 0 {
																		v896 = int32(-1)
																	} else {
																		v896 = int32(1)
																	}
																	v925 = v896
																} else {
																	if int32(0) <= base.I32_extend16_s(v787) {
																		v899 = v17 + int32(8)
																	} else {
																		v899 = v817
																	}
																	if int32(0) <= base.I32_extend16_s(v786) {
																		v902 = v25 + int32(8)
																	} else {
																		v902 = v842
																	}
																	if v891 == int32(0) {
																		if v874 == int32(_a_F_width_bucket_numeric_11) {
																			v925 = int32(1)
																		} else {
																			v908 = F_cmp_abs_common(m, v899, v840, v838, v902, v865, v863)
																			mBase = m.M
																			v925 = v908
																		}
																	} else {
																		if v874 == int32(0) {
																			v925 = int32(-1)
																		} else {
																			v912 = F_cmp_abs_common(m, v902, v865, v863, v899, v840, v838)
																			mBase = m.M
																			v925 = v912
																		}
																	}
																}
															}
														}
													}
													if v925 <= int32(0) {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v934 = m.ExcPending
														if v934 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v940 = m.ExcPending
														if v940 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													}
												}
											default:
												v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
												v286 = int32(_a_F_width_bucket_numeric_2)
												v287 = v285 & v286
												if v287 == v286 {
													if v285 != int32(_a_F_width_bucket_numeric_7) {
														if v285 != int32(_a_F_width_bucket_numeric_2) {
															if v284 != int32(_a_F_width_bucket_numeric_9) {
																v306 = int32(-1)
															} else {
																v306 = int32(0)
															}
															v423 = v306
														} else {
															v423 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_2))
														}
													} else {
														if v284 == int32(_a_F_width_bucket_numeric_2) {
															v301 = int32(-1)
														} else {
															v301 = base.B2i32(v284 != int32(_a_F_width_bucket_numeric_7))
														}
														v423 = v301
													}
												} else {
													if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v284) {
														if v284 == int32(_a_F_width_bucket_numeric_9) {
															v313 = int32(1)
														} else {
															v313 = int32(-1)
														}
														v423 = v313
													} else {
														v315 = v17 + int32(6)
														v316 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
														v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v285))
														if int32(0) <= base.I32_extend16_s(v285) {
															v324 = int32(-8)
														} else {
															v324 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v285) {
															v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v315))))
															v336 = v326
														} else {
															v336 = v285<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v285&int32(63)
														}
														v338 = int32(base.Ui32(int32(base.Ui32(v316)>>(uint(int32(2))%32))+v324) >> (uint(int32(1)) % 32))
														v340 = v22 + int32(6)
														v341 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														v348 = base.B2i32(int32(0) <= base.I32_extend16_s(v284))
														if int32(0) <= base.I32_extend16_s(v284) {
															v349 = int32(-8)
														} else {
															v349 = int32(-6)
														}
														if int32(0) <= base.I32_extend16_s(v284) {
															v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v340))))
															v361 = v351
														} else {
															v361 = v284<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v284&int32(63)
														}
														v362 = int32(1)
														v363 = int32(base.Ui32(int32(base.Ui32(v341)>>(uint(int32(2))%32))+v349) >> (uint(v362) % 32))
														v369 = v284 & int32(_a_F_width_bucket_numeric_2)
														if v369 == int32(_a_F_width_bucket_numeric_10) {
															v372 = v284 << (uint(v362) % 32) & int32(_a_F_width_bucket_numeric_11)
														} else {
															v372 = v369
														}
														if v338 == int32(0) {
															if v363 == int32(0) {
																v423 = int32(0)
															} else {
																if v372 == int32(_a_F_width_bucket_numeric_11) {
																	v382 = int32(1)
																} else {
																	v382 = int32(-1)
																}
																v423 = v382
															}
														} else {
															if v287 == int32(_a_F_width_bucket_numeric_10) {
																v389 = v285 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v389 = v287
															}
															if v363 == int32(0) {
																if v389 != 0 {
																	v394 = int32(-1)
																} else {
																	v394 = int32(1)
																}
																v423 = v394
															} else {
																if int32(0) <= base.I32_extend16_s(v285) {
																	v397 = v17 + int32(8)
																} else {
																	v397 = v315
																}
																if int32(0) <= base.I32_extend16_s(v284) {
																	v400 = v22 + int32(8)
																} else {
																	v400 = v340
																}
																if v389 == int32(0) {
																	if v372 == int32(_a_F_width_bucket_numeric_11) {
																		v423 = int32(1)
																	} else {
																		v406 = F_cmp_abs_common(m, v397, v338, v336, v400, v363, v361)
																		mBase = m.M
																		v423 = v406
																	}
																} else {
																	if v372 == int32(0) {
																		v423 = int32(-1)
																	} else {
																		v410 = F_cmp_abs_common(m, v400, v363, v361, v397, v338, v336)
																		mBase = m.M
																		v423 = v410
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
														return int32(0)
													} else {
														v429 = int32(0)
														*(*uint16)(unsafe.Add(mBase, uint32(v427))) = uint16(v429)
														v432 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[0]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v432
														v435 = *(*int64)(unsafe.Add(mBase, _c_F_width_bucket_numeric[1]))
														*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v435
														*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v427 + int32(2)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v427
														v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
														mBase = m.M
														v947 = m.ExcPending
														if v947 != 0 {
															return int32(0)
														} else {
															if v946 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1019 = m.ExcPending
																if v1019 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1022 = m.ExcPending
																	if v1022 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																		mBase = m.M
																		v1026 = m.ExcPending
																		if v1026 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																			mBase = m.M
																			v1031 = m.ExcPending
																			if v1031 != 0 {
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
																v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	F_pfree(m, v69)
																	mBase = m.M
																	v956 = m.ExcPending
																	if v956 != 0 {
																		return int32(0)
																	} else {
																		v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																		if v957 != 0 {
																			F_pfree(m, v957)
																			mBase = m.M
																			v959 = m.ExcPending
																			if v959 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v950)
																		}
																	}
																}
															}
														}
													}
												} else {
													v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
													v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v454 = int32(_a_F_width_bucket_numeric_2)
													v455 = v453 & v454
													if v455 == v454 {
														if v453 != int32(_a_F_width_bucket_numeric_7) {
															if v453 != int32(_a_F_width_bucket_numeric_2) {
																if v452 != int32(_a_F_width_bucket_numeric_9) {
																	v474 = int32(-1)
																} else {
																	v474 = int32(0)
																}
																v591 = v474
															} else {
																v591 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_2))
															}
														} else {
															if v452 == int32(_a_F_width_bucket_numeric_2) {
																v469 = int32(-1)
															} else {
																v469 = base.B2i32(v452 != int32(_a_F_width_bucket_numeric_7))
															}
															v591 = v469
														}
													} else {
														if base.Ui32(int32(_a_F_width_bucket_numeric_2)) <= base.Ui32(v452) {
															if v452 == int32(_a_F_width_bucket_numeric_9) {
																v481 = int32(1)
															} else {
																v481 = int32(-1)
															}
															v591 = v481
														} else {
															v483 = v17 + int32(6)
															v484 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v491 = base.B2i32(int32(0) <= base.I32_extend16_s(v453))
															if int32(0) <= base.I32_extend16_s(v453) {
																v492 = int32(-8)
															} else {
																v492 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v453) {
																v494 = int32(*(*int16)(unsafe.Add(mBase, uint32(v483))))
																v504 = v494
															} else {
																v504 = v453<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v453&int32(63)
															}
															v506 = int32(base.Ui32(int32(base.Ui32(v484)>>(uint(int32(2))%32))+v492) >> (uint(int32(1)) % 32))
															v508 = v25 + int32(6)
															v509 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
															v516 = base.B2i32(int32(0) <= base.I32_extend16_s(v452))
															if int32(0) <= base.I32_extend16_s(v452) {
																v517 = int32(-8)
															} else {
																v517 = int32(-6)
															}
															if int32(0) <= base.I32_extend16_s(v452) {
																v519 = int32(*(*int16)(unsafe.Add(mBase, uint32(v508))))
																v529 = v519
															} else {
																v529 = v452<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v452&int32(63)
															}
															v530 = int32(1)
															v531 = int32(base.Ui32(int32(base.Ui32(v509)>>(uint(int32(2))%32))+v517) >> (uint(v530) % 32))
															v537 = v452 & int32(_a_F_width_bucket_numeric_2)
															if v537 == int32(_a_F_width_bucket_numeric_10) {
																v540 = v452 << (uint(v530) % 32) & int32(_a_F_width_bucket_numeric_11)
															} else {
																v540 = v537
															}
															if v506 == int32(0) {
																if v531 == int32(0) {
																	v591 = int32(0)
																} else {
																	if v540 == int32(_a_F_width_bucket_numeric_11) {
																		v550 = int32(1)
																	} else {
																		v550 = int32(-1)
																	}
																	v591 = v550
																}
															} else {
																if v455 == int32(_a_F_width_bucket_numeric_10) {
																	v557 = v453 << (uint(int32(1)) % 32) & int32(_a_F_width_bucket_numeric_11)
																} else {
																	v557 = v455
																}
																if v531 == int32(0) {
																	if v557 != 0 {
																		v562 = int32(-1)
																	} else {
																		v562 = int32(1)
																	}
																	v591 = v562
																} else {
																	if int32(0) <= base.I32_extend16_s(v453) {
																		v565 = v17 + int32(8)
																	} else {
																		v565 = v483
																	}
																	if int32(0) <= base.I32_extend16_s(v452) {
																		v568 = v25 + int32(8)
																	} else {
																		v568 = v508
																	}
																	if v557 == int32(0) {
																		if v540 == int32(_a_F_width_bucket_numeric_11) {
																			v591 = int32(1)
																		} else {
																			v574 = F_cmp_abs_common(m, v565, v506, v504, v568, v531, v529)
																			mBase = m.M
																			v591 = v574
																		}
																	} else {
																		if v540 == int32(0) {
																			v591 = int32(-1)
																		} else {
																			v578 = F_cmp_abs_common(m, v568, v531, v529, v565, v506, v504)
																			mBase = m.M
																			v591 = v578
																		}
																	}
																}
															}
														}
													}
													if int32(0) <= v591 {
														F_add_var(m, v12+int32(-32), int32(_a_F_width_bucket_numeric_14), v12+int32(-56))
														mBase = m.M
														v600 = m.ExcPending
														if v600 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
																			}
																		}
																	}
																}
															}
														}
													} else {
														F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
														mBase = m.M
														v606 = m.ExcPending
														if v606 != 0 {
															return int32(0)
														} else {
															v946 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v947 = m.ExcPending
															if v947 != 0 {
																return int32(0)
															} else {
																if v946 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1019 = m.ExcPending
																	if v1019 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1022 = m.ExcPending
																		if v1022 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																			mBase = m.M
																			v1026 = m.ExcPending
																			if v1026 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																				mBase = m.M
																				v1031 = m.ExcPending
																				if v1031 != 0 {
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
																	v950 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v950-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1019 = m.ExcPending
																		if v1019 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1022 = m.ExcPending
																			if v1022 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(_a_F_width_bucket_numeric_13), int32(0))
																				mBase = m.M
																				v1026 = m.ExcPending
																				if v1026 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(2040), int32(_a_F_width_bucket_numeric_5))
																					mBase = m.M
																					v1031 = m.ExcPending
																					if v1031 != 0 {
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
																		F_pfree(m, v69)
																		mBase = m.M
																		v956 = m.ExcPending
																		if v956 != 0 {
																			return int32(0)
																		} else {
																			v957 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v957 != 0 {
																				F_pfree(m, v957)
																				mBase = m.M
																				v959 = m.ExcPending
																				if v959 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v950)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v950)
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
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v968 = m.ExcPending
					if v968 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(386138242))
						mBase = m.M
						v971 = m.ExcPending
						if v971 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_width_bucket_numeric_15), int32(0))
							mBase = m.M
							v975 = m.ExcPending
							if v975 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_width_bucket_numeric_4), int32(1980), int32(_a_F_width_bucket_numeric_5))
								mBase = m.M
								v980 = m.ExcPending
								if v980 != 0 {
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
