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
	var v61 int64
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v440 int64
	_ = v440
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int64
	_ = v777
	var v780 int64
	_ = v780
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v939 int32
	_ = v939
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int64
	_ = v964
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
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
					if base.Ui32(v30) <= base.Ui32(int32(49151)) {
						v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
						if base.Ui32(int32(49151)) < base.Ui32(v33) {
							v42 = v33
							if v42&int32(65535) == int32(49152) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1000 = m.ExcPending
								if v1000 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v1003 = m.ExcPending
									if v1003 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(529241), int32(0))
										mBase = m.M
										v1007 = m.ExcPending
										if v1007 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(501522), int32(1991), int32(492002))
											mBase = m.M
											v1012 = m.ExcPending
											if v1012 != 0 {
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
								if v48 == int32(49152) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1000 = m.ExcPending
									if v1000 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v1003 = m.ExcPending
										if v1003 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(529241), int32(0))
											mBase = m.M
											v1007 = m.ExcPending
											if v1007 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(501522), int32(1991), int32(492002))
												mBase = m.M
												v1012 = m.ExcPending
												if v1012 != 0 {
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
									if v47&int32(57343) == int32(53248) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1016 = m.ExcPending
										if v1016 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1019 = m.ExcPending
											if v1019 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(351481), int32(0))
												mBase = m.M
												v1023 = m.ExcPending
												if v1023 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(501522), int32(1996), int32(492002))
													mBase = m.M
													v1028 = m.ExcPending
													if v1028 != 0 {
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
										if v48&int32(57343) == int32(53248) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v1016 = m.ExcPending
											if v1016 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(386138242))
												mBase = m.M
												v1019 = m.ExcPending
												if v1019 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(351481), int32(0))
													mBase = m.M
													v1023 = m.ExcPending
													if v1023 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(501522), int32(1996), int32(492002))
														mBase = m.M
														v1028 = m.ExcPending
														if v1028 != 0 {
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
											v61 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v61
											v68 = F_palloc(m, int32(12))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
												v71 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v71)
												*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
												v79 = v71
												v84 = v68 + int32(12)
												v88 = base.I64_extend_i32_u(v27)
												for {
													v91 = v84 - int32(2)
													v93 = base.I64_div_u_s(v88, int64(10000))
													v96 = v93*int64(55536) + v88
													*(*uint16)(unsafe.Add(mBase, uint32(v91))) = uint16(v96)
													v99 = v79 + int32(1)
													if base.Ui64(int64(9999)) < base.Ui64(v88) {
														v79 = v99
														v84 = v91
														v88 = v93
														continue
													} else {
														break
													}
													break
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v79
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v91
												v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v119 = int32(49152)
												v120 = v118 & v119
												if v120 == v119 {
													if v118 != int32(53248) {
														if v118 != int32(49152) {
															if v117 != int32(61440) {
																v139 = int32(-1)
															} else {
																v139 = int32(0)
															}
															v258 = v139
														} else {
															v258 = base.B2i32(v117 != int32(49152))
														}
													} else {
														if v117 == int32(49152) {
															v134 = int32(-1)
														} else {
															v134 = base.B2i32(v117 != int32(53248))
														}
														v258 = v134
													}
												} else {
													if base.Ui32(int32(49152)) <= base.Ui32(v117) {
														if v117 == int32(61440) {
															v146 = int32(1)
														} else {
															v146 = int32(-1)
														}
														v258 = v146
													} else {
														v148 = v22 + int32(6)
														v153 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
														if int32(0) <= base.I32_extend16_s(v118) {
															v154 = int32(-8)
														} else {
															v154 = int32(-6)
														}
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														if int32(0) <= base.I32_extend16_s(v118) {
															v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
															v168 = v158
														} else {
															v168 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
														}
														v169 = v154 + int32(base.Ui32(v155)>>(uint(int32(2))%32))
														v171 = v25 + int32(6)
														v176 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
														if int32(0) <= base.I32_extend16_s(v117) {
															v177 = int32(-8)
														} else {
															v177 = int32(-6)
														}
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														if int32(0) <= base.I32_extend16_s(v117) {
															v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v171))))
															v191 = v181
														} else {
															v191 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
														}
														v192 = v177 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
														v198 = v117 & int32(49152)
														if v198 == int32(32768) {
															v201 = v117 << (uint(int32(1)) % 32) & int32(16384)
														} else {
															v201 = v198
														}
														if base.Ui32(v169) <= base.Ui32(int32(1)) {
															if base.Ui32(v192) < base.Ui32(int32(2)) {
																v258 = int32(0)
															} else {
																if v201 == int32(16384) {
																	v211 = int32(1)
																} else {
																	v211 = int32(-1)
																}
																v258 = v211
															}
														} else {
															if v120 == int32(32768) {
																v218 = v118 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v218 = v120
															}
															if base.Ui32(v192) <= base.Ui32(int32(1)) {
																if v218 != 0 {
																	v223 = int32(-1)
																} else {
																	v223 = int32(1)
																}
																v258 = v223
															} else {
																if int32(0) <= base.I32_extend16_s(v118) {
																	v226 = v22 + int32(8)
																} else {
																	v226 = v148
																}
																v228 = int32(base.Ui32(v169) >> (uint(int32(1)) % 32))
																if int32(0) <= base.I32_extend16_s(v117) {
																	v231 = v25 + int32(8)
																} else {
																	v231 = v171
																}
																v233 = int32(base.Ui32(v192) >> (uint(int32(1)) % 32))
																if v218 == int32(0) {
																	if v201 == int32(16384) {
																		v258 = int32(1)
																	} else {
																		v239 = F_cmp_abs_common(m, v226, v228, v168, v231, v233, v191)
																		mBase = m.M
																		v258 = v239
																	}
																} else {
																	if v201 == int32(0) {
																		v258 = int32(-1)
																	} else {
																		v243 = F_cmp_abs_common(m, v231, v233, v191, v226, v228, v168)
																		mBase = m.M
																		v258 = v243
																	}
																}
															}
														}
													}
												}
												switch v258 {
												case 0:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v262 = m.ExcPending
													if v262 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(386138242))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(425786), int32(0))
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(501522), int32(2010), int32(492002))
																mBase = m.M
																v274 = m.ExcPending
																if v274 != 0 {
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
													v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v629 = int32(49152)
													v630 = v628 & v629
													if v630 == v629 {
														if v628 != int32(53248) {
															if v628 != int32(49152) {
																if v627 != int32(61440) {
																	v649 = int32(-1)
																} else {
																	v649 = int32(0)
																}
																v768 = v649
															} else {
																v768 = base.B2i32(v627 != int32(49152))
															}
														} else {
															if v627 == int32(49152) {
																v644 = int32(-1)
															} else {
																v644 = base.B2i32(v627 != int32(53248))
															}
															v768 = v644
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v627) {
															if v627 == int32(61440) {
																v656 = int32(1)
															} else {
																v656 = int32(-1)
															}
															v768 = v656
														} else {
															v658 = v17 + int32(6)
															v663 = base.B2i32(int32(0) <= base.I32_extend16_s(v628))
															if int32(0) <= base.I32_extend16_s(v628) {
																v664 = int32(-8)
															} else {
																v664 = int32(-6)
															}
															v665 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v628) {
																v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
																v678 = v668
															} else {
																v678 = v628<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v628&int32(63)
															}
															v679 = v664 + int32(base.Ui32(v665)>>(uint(int32(2))%32))
															v681 = v22 + int32(6)
															v686 = base.B2i32(int32(0) <= base.I32_extend16_s(v627))
															if int32(0) <= base.I32_extend16_s(v627) {
																v687 = int32(-8)
															} else {
																v687 = int32(-6)
															}
															v688 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v627) {
																v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681))))
																v701 = v691
															} else {
																v701 = v627<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v627&int32(63)
															}
															v702 = v687 + int32(base.Ui32(v688)>>(uint(int32(2))%32))
															v708 = v627 & int32(49152)
															if v708 == int32(32768) {
																v711 = v627 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v711 = v708
															}
															if base.Ui32(v679) <= base.Ui32(int32(1)) {
																if base.Ui32(v702) < base.Ui32(int32(2)) {
																	v768 = int32(0)
																} else {
																	if v711 == int32(16384) {
																		v721 = int32(1)
																	} else {
																		v721 = int32(-1)
																	}
																	v768 = v721
																}
															} else {
																if v630 == int32(32768) {
																	v728 = v628 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v728 = v630
																}
																if base.Ui32(v702) <= base.Ui32(int32(1)) {
																	if v728 != 0 {
																		v733 = int32(-1)
																	} else {
																		v733 = int32(1)
																	}
																	v768 = v733
																} else {
																	if int32(0) <= base.I32_extend16_s(v628) {
																		v736 = v17 + int32(8)
																	} else {
																		v736 = v658
																	}
																	v738 = int32(base.Ui32(v679) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v627) {
																		v741 = v22 + int32(8)
																	} else {
																		v741 = v681
																	}
																	v743 = int32(base.Ui32(v702) >> (uint(int32(1)) % 32))
																	if v728 == int32(0) {
																		if v711 == int32(16384) {
																			v768 = int32(1)
																		} else {
																			v749 = F_cmp_abs_common(m, v736, v738, v678, v741, v743, v701)
																			mBase = m.M
																			v768 = v749
																		}
																	} else {
																		if v711 == int32(0) {
																			v768 = int32(-1)
																		} else {
																			v753 = F_cmp_abs_common(m, v741, v743, v701, v736, v738, v678)
																			mBase = m.M
																			v768 = v753
																		}
																	}
																}
															}
														}
													}
													if int32(0) < v768 {
														v772 = F_palloc(m, int32(2))
														mBase = m.M
														v773 = m.ExcPending
														if v773 != 0 {
															return int32(0)
														} else {
															v774 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v772))) = uint16(v774)
															v777 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v777
															v780 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v780
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v772 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v772
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v800 = int32(49152)
														v801 = v799 & v800
														if v801 == v800 {
															if v799 != int32(53248) {
																if v799 != int32(49152) {
																	if v798 != int32(61440) {
																		v820 = int32(-1)
																	} else {
																		v820 = int32(0)
																	}
																	v939 = v820
																} else {
																	v939 = base.B2i32(v798 != int32(49152))
																}
															} else {
																if v798 == int32(49152) {
																	v815 = int32(-1)
																} else {
																	v815 = base.B2i32(v798 != int32(53248))
																}
																v939 = v815
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v798) {
																if v798 == int32(61440) {
																	v827 = int32(1)
																} else {
																	v827 = int32(-1)
																}
																v939 = v827
															} else {
																v829 = v17 + int32(6)
																v834 = base.B2i32(int32(0) <= base.I32_extend16_s(v799))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v835 = int32(-8)
																} else {
																	v835 = int32(-6)
																}
																v836 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v829))))
																	v849 = v839
																} else {
																	v849 = v799<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v799&int32(63)
																}
																v850 = v835 + int32(base.Ui32(v836)>>(uint(int32(2))%32))
																v852 = v25 + int32(6)
																v857 = base.B2i32(int32(0) <= base.I32_extend16_s(v798))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v858 = int32(-8)
																} else {
																	v858 = int32(-6)
																}
																v859 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v852))))
																	v872 = v862
																} else {
																	v872 = v798<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v798&int32(63)
																}
																v873 = v858 + int32(base.Ui32(v859)>>(uint(int32(2))%32))
																v879 = v798 & int32(49152)
																if v879 == int32(32768) {
																	v882 = v798 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v882 = v879
																}
																if base.Ui32(v850) <= base.Ui32(int32(1)) {
																	if base.Ui32(v873) < base.Ui32(int32(2)) {
																		v939 = int32(0)
																	} else {
																		if v882 == int32(16384) {
																			v892 = int32(1)
																		} else {
																			v892 = int32(-1)
																		}
																		v939 = v892
																	}
																} else {
																	if v801 == int32(32768) {
																		v899 = v799 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v899 = v801
																	}
																	if base.Ui32(v873) <= base.Ui32(int32(1)) {
																		if v899 != 0 {
																			v904 = int32(-1)
																		} else {
																			v904 = int32(1)
																		}
																		v939 = v904
																	} else {
																		if int32(0) <= base.I32_extend16_s(v799) {
																			v907 = v17 + int32(8)
																		} else {
																			v907 = v829
																		}
																		v909 = int32(base.Ui32(v850) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v798) {
																			v912 = v25 + int32(8)
																		} else {
																			v912 = v852
																		}
																		v914 = int32(base.Ui32(v873) >> (uint(int32(1)) % 32))
																		if v899 == int32(0) {
																			if v882 == int32(16384) {
																				v939 = int32(1)
																			} else {
																				v920 = F_cmp_abs_common(m, v907, v909, v849, v912, v914, v872)
																				mBase = m.M
																				v939 = v920
																			}
																		} else {
																			if v882 == int32(0) {
																				v939 = int32(-1)
																			} else {
																				v924 = F_cmp_abs_common(m, v912, v914, v872, v907, v909, v849)
																				mBase = m.M
																				v939 = v924
																			}
																		}
																	}
																}
															}
														}
														if v939 <= int32(0) {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v948 = m.ExcPending
															if v948 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v954 = m.ExcPending
															if v954 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												default:
													v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v289 = int32(49152)
													v290 = v288 & v289
													if v290 == v289 {
														if v288 != int32(53248) {
															if v288 != int32(49152) {
																if v287 != int32(61440) {
																	v309 = int32(-1)
																} else {
																	v309 = int32(0)
																}
																v428 = v309
															} else {
																v428 = base.B2i32(v287 != int32(49152))
															}
														} else {
															if v287 == int32(49152) {
																v304 = int32(-1)
															} else {
																v304 = base.B2i32(v287 != int32(53248))
															}
															v428 = v304
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v287) {
															if v287 == int32(61440) {
																v316 = int32(1)
															} else {
																v316 = int32(-1)
															}
															v428 = v316
														} else {
															v318 = v17 + int32(6)
															v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v288))
															if int32(0) <= base.I32_extend16_s(v288) {
																v324 = int32(-8)
															} else {
																v324 = int32(-6)
															}
															v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v288) {
																v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v318))))
																v338 = v328
															} else {
																v338 = v288<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v288&int32(63)
															}
															v339 = v324 + int32(base.Ui32(v325)>>(uint(int32(2))%32))
															v341 = v22 + int32(6)
															v346 = base.B2i32(int32(0) <= base.I32_extend16_s(v287))
															if int32(0) <= base.I32_extend16_s(v287) {
																v347 = int32(-8)
															} else {
																v347 = int32(-6)
															}
															v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v287) {
																v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341))))
																v361 = v351
															} else {
																v361 = v287<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v287&int32(63)
															}
															v362 = v347 + int32(base.Ui32(v348)>>(uint(int32(2))%32))
															v368 = v287 & int32(49152)
															if v368 == int32(32768) {
																v371 = v287 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v371 = v368
															}
															if base.Ui32(v339) <= base.Ui32(int32(1)) {
																if base.Ui32(v362) < base.Ui32(int32(2)) {
																	v428 = int32(0)
																} else {
																	if v371 == int32(16384) {
																		v381 = int32(1)
																	} else {
																		v381 = int32(-1)
																	}
																	v428 = v381
																}
															} else {
																if v290 == int32(32768) {
																	v388 = v288 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v388 = v290
																}
																if base.Ui32(v362) <= base.Ui32(int32(1)) {
																	if v388 != 0 {
																		v393 = int32(-1)
																	} else {
																		v393 = int32(1)
																	}
																	v428 = v393
																} else {
																	if int32(0) <= base.I32_extend16_s(v288) {
																		v396 = v17 + int32(8)
																	} else {
																		v396 = v318
																	}
																	v398 = int32(base.Ui32(v339) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v287) {
																		v401 = v22 + int32(8)
																	} else {
																		v401 = v341
																	}
																	v403 = int32(base.Ui32(v362) >> (uint(int32(1)) % 32))
																	if v388 == int32(0) {
																		if v371 == int32(16384) {
																			v428 = int32(1)
																		} else {
																			v409 = F_cmp_abs_common(m, v396, v398, v338, v401, v403, v361)
																			mBase = m.M
																			v428 = v409
																		}
																	} else {
																		if v371 == int32(0) {
																			v428 = int32(-1)
																		} else {
																			v413 = F_cmp_abs_common(m, v401, v403, v361, v396, v398, v338)
																			mBase = m.M
																			v428 = v413
																		}
																	}
																}
															}
														}
													}
													if v428 < int32(0) {
														v432 = F_palloc(m, int32(2))
														mBase = m.M
														v433 = m.ExcPending
														if v433 != 0 {
															return int32(0)
														} else {
															v434 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v432))) = uint16(v434)
															v437 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v437
															v440 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v440
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v432 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v432
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v460 = int32(49152)
														v461 = v459 & v460
														if v461 == v460 {
															if v459 != int32(53248) {
																if v459 != int32(49152) {
																	if v458 != int32(61440) {
																		v480 = int32(-1)
																	} else {
																		v480 = int32(0)
																	}
																	v599 = v480
																} else {
																	v599 = base.B2i32(v458 != int32(49152))
																}
															} else {
																if v458 == int32(49152) {
																	v475 = int32(-1)
																} else {
																	v475 = base.B2i32(v458 != int32(53248))
																}
																v599 = v475
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v458) {
																if v458 == int32(61440) {
																	v487 = int32(1)
																} else {
																	v487 = int32(-1)
																}
																v599 = v487
															} else {
																v489 = v17 + int32(6)
																v494 = base.B2i32(int32(0) <= base.I32_extend16_s(v459))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v495 = int32(-8)
																} else {
																	v495 = int32(-6)
																}
																v496 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489))))
																	v509 = v499
																} else {
																	v509 = v459<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v459&int32(63)
																}
																v510 = v495 + int32(base.Ui32(v496)>>(uint(int32(2))%32))
																v512 = v25 + int32(6)
																v517 = base.B2i32(int32(0) <= base.I32_extend16_s(v458))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v518 = int32(-8)
																} else {
																	v518 = int32(-6)
																}
																v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v512))))
																	v532 = v522
																} else {
																	v532 = v458<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v458&int32(63)
																}
																v533 = v518 + int32(base.Ui32(v519)>>(uint(int32(2))%32))
																v539 = v458 & int32(49152)
																if v539 == int32(32768) {
																	v542 = v458 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v542 = v539
																}
																if base.Ui32(v510) <= base.Ui32(int32(1)) {
																	if base.Ui32(v533) < base.Ui32(int32(2)) {
																		v599 = int32(0)
																	} else {
																		if v542 == int32(16384) {
																			v552 = int32(1)
																		} else {
																			v552 = int32(-1)
																		}
																		v599 = v552
																	}
																} else {
																	if v461 == int32(32768) {
																		v559 = v459 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v559 = v461
																	}
																	if base.Ui32(v533) <= base.Ui32(int32(1)) {
																		if v559 != 0 {
																			v564 = int32(-1)
																		} else {
																			v564 = int32(1)
																		}
																		v599 = v564
																	} else {
																		if int32(0) <= base.I32_extend16_s(v459) {
																			v567 = v17 + int32(8)
																		} else {
																			v567 = v489
																		}
																		v569 = int32(base.Ui32(v510) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v458) {
																			v572 = v25 + int32(8)
																		} else {
																			v572 = v512
																		}
																		v574 = int32(base.Ui32(v533) >> (uint(int32(1)) % 32))
																		if v559 == int32(0) {
																			if v542 == int32(16384) {
																				v599 = int32(1)
																			} else {
																				v580 = F_cmp_abs_common(m, v567, v569, v509, v572, v574, v532)
																				mBase = m.M
																				v599 = v580
																			}
																		} else {
																			if v542 == int32(0) {
																				v599 = int32(-1)
																			} else {
																				v584 = F_cmp_abs_common(m, v572, v574, v532, v567, v569, v509)
																				mBase = m.M
																				v599 = v584
																			}
																		}
																	}
																}
															}
														}
														if int32(0) <= v599 {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v608 = m.ExcPending
															if v608 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v614 = m.ExcPending
															if v614 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
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
							v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
							if base.Ui32(int32(49151)) < base.Ui32(v36) {
								v47 = v33
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
								if v48 == int32(49152) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1000 = m.ExcPending
									if v1000 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v1003 = m.ExcPending
										if v1003 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(529241), int32(0))
											mBase = m.M
											v1007 = m.ExcPending
											if v1007 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(501522), int32(1991), int32(492002))
												mBase = m.M
												v1012 = m.ExcPending
												if v1012 != 0 {
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
									if v47&int32(57343) == int32(53248) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1016 = m.ExcPending
										if v1016 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1019 = m.ExcPending
											if v1019 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(351481), int32(0))
												mBase = m.M
												v1023 = m.ExcPending
												if v1023 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(501522), int32(1996), int32(492002))
													mBase = m.M
													v1028 = m.ExcPending
													if v1028 != 0 {
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
										if v48&int32(57343) == int32(53248) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v1016 = m.ExcPending
											if v1016 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(386138242))
												mBase = m.M
												v1019 = m.ExcPending
												if v1019 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(351481), int32(0))
													mBase = m.M
													v1023 = m.ExcPending
													if v1023 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(501522), int32(1996), int32(492002))
														mBase = m.M
														v1028 = m.ExcPending
														if v1028 != 0 {
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
											v61 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v61
											v68 = F_palloc(m, int32(12))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
												v71 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v71)
												*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
												v79 = v71
												v84 = v68 + int32(12)
												v88 = base.I64_extend_i32_u(v27)
												for {
													v91 = v84 - int32(2)
													v93 = base.I64_div_u_s(v88, int64(10000))
													v96 = v93*int64(55536) + v88
													*(*uint16)(unsafe.Add(mBase, uint32(v91))) = uint16(v96)
													v99 = v79 + int32(1)
													if base.Ui64(int64(9999)) < base.Ui64(v88) {
														v79 = v99
														v84 = v91
														v88 = v93
														continue
													} else {
														break
													}
													break
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v79
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v91
												v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v119 = int32(49152)
												v120 = v118 & v119
												if v120 == v119 {
													if v118 != int32(53248) {
														if v118 != int32(49152) {
															if v117 != int32(61440) {
																v139 = int32(-1)
															} else {
																v139 = int32(0)
															}
															v258 = v139
														} else {
															v258 = base.B2i32(v117 != int32(49152))
														}
													} else {
														if v117 == int32(49152) {
															v134 = int32(-1)
														} else {
															v134 = base.B2i32(v117 != int32(53248))
														}
														v258 = v134
													}
												} else {
													if base.Ui32(int32(49152)) <= base.Ui32(v117) {
														if v117 == int32(61440) {
															v146 = int32(1)
														} else {
															v146 = int32(-1)
														}
														v258 = v146
													} else {
														v148 = v22 + int32(6)
														v153 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
														if int32(0) <= base.I32_extend16_s(v118) {
															v154 = int32(-8)
														} else {
															v154 = int32(-6)
														}
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														if int32(0) <= base.I32_extend16_s(v118) {
															v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
															v168 = v158
														} else {
															v168 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
														}
														v169 = v154 + int32(base.Ui32(v155)>>(uint(int32(2))%32))
														v171 = v25 + int32(6)
														v176 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
														if int32(0) <= base.I32_extend16_s(v117) {
															v177 = int32(-8)
														} else {
															v177 = int32(-6)
														}
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														if int32(0) <= base.I32_extend16_s(v117) {
															v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v171))))
															v191 = v181
														} else {
															v191 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
														}
														v192 = v177 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
														v198 = v117 & int32(49152)
														if v198 == int32(32768) {
															v201 = v117 << (uint(int32(1)) % 32) & int32(16384)
														} else {
															v201 = v198
														}
														if base.Ui32(v169) <= base.Ui32(int32(1)) {
															if base.Ui32(v192) < base.Ui32(int32(2)) {
																v258 = int32(0)
															} else {
																if v201 == int32(16384) {
																	v211 = int32(1)
																} else {
																	v211 = int32(-1)
																}
																v258 = v211
															}
														} else {
															if v120 == int32(32768) {
																v218 = v118 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v218 = v120
															}
															if base.Ui32(v192) <= base.Ui32(int32(1)) {
																if v218 != 0 {
																	v223 = int32(-1)
																} else {
																	v223 = int32(1)
																}
																v258 = v223
															} else {
																if int32(0) <= base.I32_extend16_s(v118) {
																	v226 = v22 + int32(8)
																} else {
																	v226 = v148
																}
																v228 = int32(base.Ui32(v169) >> (uint(int32(1)) % 32))
																if int32(0) <= base.I32_extend16_s(v117) {
																	v231 = v25 + int32(8)
																} else {
																	v231 = v171
																}
																v233 = int32(base.Ui32(v192) >> (uint(int32(1)) % 32))
																if v218 == int32(0) {
																	if v201 == int32(16384) {
																		v258 = int32(1)
																	} else {
																		v239 = F_cmp_abs_common(m, v226, v228, v168, v231, v233, v191)
																		mBase = m.M
																		v258 = v239
																	}
																} else {
																	if v201 == int32(0) {
																		v258 = int32(-1)
																	} else {
																		v243 = F_cmp_abs_common(m, v231, v233, v191, v226, v228, v168)
																		mBase = m.M
																		v258 = v243
																	}
																}
															}
														}
													}
												}
												switch v258 {
												case 0:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v262 = m.ExcPending
													if v262 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(386138242))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(425786), int32(0))
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(501522), int32(2010), int32(492002))
																mBase = m.M
																v274 = m.ExcPending
																if v274 != 0 {
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
													v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v629 = int32(49152)
													v630 = v628 & v629
													if v630 == v629 {
														if v628 != int32(53248) {
															if v628 != int32(49152) {
																if v627 != int32(61440) {
																	v649 = int32(-1)
																} else {
																	v649 = int32(0)
																}
																v768 = v649
															} else {
																v768 = base.B2i32(v627 != int32(49152))
															}
														} else {
															if v627 == int32(49152) {
																v644 = int32(-1)
															} else {
																v644 = base.B2i32(v627 != int32(53248))
															}
															v768 = v644
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v627) {
															if v627 == int32(61440) {
																v656 = int32(1)
															} else {
																v656 = int32(-1)
															}
															v768 = v656
														} else {
															v658 = v17 + int32(6)
															v663 = base.B2i32(int32(0) <= base.I32_extend16_s(v628))
															if int32(0) <= base.I32_extend16_s(v628) {
																v664 = int32(-8)
															} else {
																v664 = int32(-6)
															}
															v665 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v628) {
																v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
																v678 = v668
															} else {
																v678 = v628<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v628&int32(63)
															}
															v679 = v664 + int32(base.Ui32(v665)>>(uint(int32(2))%32))
															v681 = v22 + int32(6)
															v686 = base.B2i32(int32(0) <= base.I32_extend16_s(v627))
															if int32(0) <= base.I32_extend16_s(v627) {
																v687 = int32(-8)
															} else {
																v687 = int32(-6)
															}
															v688 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v627) {
																v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681))))
																v701 = v691
															} else {
																v701 = v627<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v627&int32(63)
															}
															v702 = v687 + int32(base.Ui32(v688)>>(uint(int32(2))%32))
															v708 = v627 & int32(49152)
															if v708 == int32(32768) {
																v711 = v627 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v711 = v708
															}
															if base.Ui32(v679) <= base.Ui32(int32(1)) {
																if base.Ui32(v702) < base.Ui32(int32(2)) {
																	v768 = int32(0)
																} else {
																	if v711 == int32(16384) {
																		v721 = int32(1)
																	} else {
																		v721 = int32(-1)
																	}
																	v768 = v721
																}
															} else {
																if v630 == int32(32768) {
																	v728 = v628 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v728 = v630
																}
																if base.Ui32(v702) <= base.Ui32(int32(1)) {
																	if v728 != 0 {
																		v733 = int32(-1)
																	} else {
																		v733 = int32(1)
																	}
																	v768 = v733
																} else {
																	if int32(0) <= base.I32_extend16_s(v628) {
																		v736 = v17 + int32(8)
																	} else {
																		v736 = v658
																	}
																	v738 = int32(base.Ui32(v679) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v627) {
																		v741 = v22 + int32(8)
																	} else {
																		v741 = v681
																	}
																	v743 = int32(base.Ui32(v702) >> (uint(int32(1)) % 32))
																	if v728 == int32(0) {
																		if v711 == int32(16384) {
																			v768 = int32(1)
																		} else {
																			v749 = F_cmp_abs_common(m, v736, v738, v678, v741, v743, v701)
																			mBase = m.M
																			v768 = v749
																		}
																	} else {
																		if v711 == int32(0) {
																			v768 = int32(-1)
																		} else {
																			v753 = F_cmp_abs_common(m, v741, v743, v701, v736, v738, v678)
																			mBase = m.M
																			v768 = v753
																		}
																	}
																}
															}
														}
													}
													if int32(0) < v768 {
														v772 = F_palloc(m, int32(2))
														mBase = m.M
														v773 = m.ExcPending
														if v773 != 0 {
															return int32(0)
														} else {
															v774 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v772))) = uint16(v774)
															v777 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v777
															v780 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v780
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v772 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v772
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v800 = int32(49152)
														v801 = v799 & v800
														if v801 == v800 {
															if v799 != int32(53248) {
																if v799 != int32(49152) {
																	if v798 != int32(61440) {
																		v820 = int32(-1)
																	} else {
																		v820 = int32(0)
																	}
																	v939 = v820
																} else {
																	v939 = base.B2i32(v798 != int32(49152))
																}
															} else {
																if v798 == int32(49152) {
																	v815 = int32(-1)
																} else {
																	v815 = base.B2i32(v798 != int32(53248))
																}
																v939 = v815
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v798) {
																if v798 == int32(61440) {
																	v827 = int32(1)
																} else {
																	v827 = int32(-1)
																}
																v939 = v827
															} else {
																v829 = v17 + int32(6)
																v834 = base.B2i32(int32(0) <= base.I32_extend16_s(v799))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v835 = int32(-8)
																} else {
																	v835 = int32(-6)
																}
																v836 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v829))))
																	v849 = v839
																} else {
																	v849 = v799<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v799&int32(63)
																}
																v850 = v835 + int32(base.Ui32(v836)>>(uint(int32(2))%32))
																v852 = v25 + int32(6)
																v857 = base.B2i32(int32(0) <= base.I32_extend16_s(v798))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v858 = int32(-8)
																} else {
																	v858 = int32(-6)
																}
																v859 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v852))))
																	v872 = v862
																} else {
																	v872 = v798<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v798&int32(63)
																}
																v873 = v858 + int32(base.Ui32(v859)>>(uint(int32(2))%32))
																v879 = v798 & int32(49152)
																if v879 == int32(32768) {
																	v882 = v798 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v882 = v879
																}
																if base.Ui32(v850) <= base.Ui32(int32(1)) {
																	if base.Ui32(v873) < base.Ui32(int32(2)) {
																		v939 = int32(0)
																	} else {
																		if v882 == int32(16384) {
																			v892 = int32(1)
																		} else {
																			v892 = int32(-1)
																		}
																		v939 = v892
																	}
																} else {
																	if v801 == int32(32768) {
																		v899 = v799 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v899 = v801
																	}
																	if base.Ui32(v873) <= base.Ui32(int32(1)) {
																		if v899 != 0 {
																			v904 = int32(-1)
																		} else {
																			v904 = int32(1)
																		}
																		v939 = v904
																	} else {
																		if int32(0) <= base.I32_extend16_s(v799) {
																			v907 = v17 + int32(8)
																		} else {
																			v907 = v829
																		}
																		v909 = int32(base.Ui32(v850) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v798) {
																			v912 = v25 + int32(8)
																		} else {
																			v912 = v852
																		}
																		v914 = int32(base.Ui32(v873) >> (uint(int32(1)) % 32))
																		if v899 == int32(0) {
																			if v882 == int32(16384) {
																				v939 = int32(1)
																			} else {
																				v920 = F_cmp_abs_common(m, v907, v909, v849, v912, v914, v872)
																				mBase = m.M
																				v939 = v920
																			}
																		} else {
																			if v882 == int32(0) {
																				v939 = int32(-1)
																			} else {
																				v924 = F_cmp_abs_common(m, v912, v914, v872, v907, v909, v849)
																				mBase = m.M
																				v939 = v924
																			}
																		}
																	}
																}
															}
														}
														if v939 <= int32(0) {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v948 = m.ExcPending
															if v948 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v954 = m.ExcPending
															if v954 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												default:
													v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v289 = int32(49152)
													v290 = v288 & v289
													if v290 == v289 {
														if v288 != int32(53248) {
															if v288 != int32(49152) {
																if v287 != int32(61440) {
																	v309 = int32(-1)
																} else {
																	v309 = int32(0)
																}
																v428 = v309
															} else {
																v428 = base.B2i32(v287 != int32(49152))
															}
														} else {
															if v287 == int32(49152) {
																v304 = int32(-1)
															} else {
																v304 = base.B2i32(v287 != int32(53248))
															}
															v428 = v304
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v287) {
															if v287 == int32(61440) {
																v316 = int32(1)
															} else {
																v316 = int32(-1)
															}
															v428 = v316
														} else {
															v318 = v17 + int32(6)
															v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v288))
															if int32(0) <= base.I32_extend16_s(v288) {
																v324 = int32(-8)
															} else {
																v324 = int32(-6)
															}
															v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v288) {
																v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v318))))
																v338 = v328
															} else {
																v338 = v288<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v288&int32(63)
															}
															v339 = v324 + int32(base.Ui32(v325)>>(uint(int32(2))%32))
															v341 = v22 + int32(6)
															v346 = base.B2i32(int32(0) <= base.I32_extend16_s(v287))
															if int32(0) <= base.I32_extend16_s(v287) {
																v347 = int32(-8)
															} else {
																v347 = int32(-6)
															}
															v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v287) {
																v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341))))
																v361 = v351
															} else {
																v361 = v287<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v287&int32(63)
															}
															v362 = v347 + int32(base.Ui32(v348)>>(uint(int32(2))%32))
															v368 = v287 & int32(49152)
															if v368 == int32(32768) {
																v371 = v287 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v371 = v368
															}
															if base.Ui32(v339) <= base.Ui32(int32(1)) {
																if base.Ui32(v362) < base.Ui32(int32(2)) {
																	v428 = int32(0)
																} else {
																	if v371 == int32(16384) {
																		v381 = int32(1)
																	} else {
																		v381 = int32(-1)
																	}
																	v428 = v381
																}
															} else {
																if v290 == int32(32768) {
																	v388 = v288 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v388 = v290
																}
																if base.Ui32(v362) <= base.Ui32(int32(1)) {
																	if v388 != 0 {
																		v393 = int32(-1)
																	} else {
																		v393 = int32(1)
																	}
																	v428 = v393
																} else {
																	if int32(0) <= base.I32_extend16_s(v288) {
																		v396 = v17 + int32(8)
																	} else {
																		v396 = v318
																	}
																	v398 = int32(base.Ui32(v339) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v287) {
																		v401 = v22 + int32(8)
																	} else {
																		v401 = v341
																	}
																	v403 = int32(base.Ui32(v362) >> (uint(int32(1)) % 32))
																	if v388 == int32(0) {
																		if v371 == int32(16384) {
																			v428 = int32(1)
																		} else {
																			v409 = F_cmp_abs_common(m, v396, v398, v338, v401, v403, v361)
																			mBase = m.M
																			v428 = v409
																		}
																	} else {
																		if v371 == int32(0) {
																			v428 = int32(-1)
																		} else {
																			v413 = F_cmp_abs_common(m, v401, v403, v361, v396, v398, v338)
																			mBase = m.M
																			v428 = v413
																		}
																	}
																}
															}
														}
													}
													if v428 < int32(0) {
														v432 = F_palloc(m, int32(2))
														mBase = m.M
														v433 = m.ExcPending
														if v433 != 0 {
															return int32(0)
														} else {
															v434 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v432))) = uint16(v434)
															v437 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v437
															v440 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v440
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v432 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v432
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v460 = int32(49152)
														v461 = v459 & v460
														if v461 == v460 {
															if v459 != int32(53248) {
																if v459 != int32(49152) {
																	if v458 != int32(61440) {
																		v480 = int32(-1)
																	} else {
																		v480 = int32(0)
																	}
																	v599 = v480
																} else {
																	v599 = base.B2i32(v458 != int32(49152))
																}
															} else {
																if v458 == int32(49152) {
																	v475 = int32(-1)
																} else {
																	v475 = base.B2i32(v458 != int32(53248))
																}
																v599 = v475
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v458) {
																if v458 == int32(61440) {
																	v487 = int32(1)
																} else {
																	v487 = int32(-1)
																}
																v599 = v487
															} else {
																v489 = v17 + int32(6)
																v494 = base.B2i32(int32(0) <= base.I32_extend16_s(v459))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v495 = int32(-8)
																} else {
																	v495 = int32(-6)
																}
																v496 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489))))
																	v509 = v499
																} else {
																	v509 = v459<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v459&int32(63)
																}
																v510 = v495 + int32(base.Ui32(v496)>>(uint(int32(2))%32))
																v512 = v25 + int32(6)
																v517 = base.B2i32(int32(0) <= base.I32_extend16_s(v458))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v518 = int32(-8)
																} else {
																	v518 = int32(-6)
																}
																v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v512))))
																	v532 = v522
																} else {
																	v532 = v458<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v458&int32(63)
																}
																v533 = v518 + int32(base.Ui32(v519)>>(uint(int32(2))%32))
																v539 = v458 & int32(49152)
																if v539 == int32(32768) {
																	v542 = v458 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v542 = v539
																}
																if base.Ui32(v510) <= base.Ui32(int32(1)) {
																	if base.Ui32(v533) < base.Ui32(int32(2)) {
																		v599 = int32(0)
																	} else {
																		if v542 == int32(16384) {
																			v552 = int32(1)
																		} else {
																			v552 = int32(-1)
																		}
																		v599 = v552
																	}
																} else {
																	if v461 == int32(32768) {
																		v559 = v459 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v559 = v461
																	}
																	if base.Ui32(v533) <= base.Ui32(int32(1)) {
																		if v559 != 0 {
																			v564 = int32(-1)
																		} else {
																			v564 = int32(1)
																		}
																		v599 = v564
																	} else {
																		if int32(0) <= base.I32_extend16_s(v459) {
																			v567 = v17 + int32(8)
																		} else {
																			v567 = v489
																		}
																		v569 = int32(base.Ui32(v510) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v458) {
																			v572 = v25 + int32(8)
																		} else {
																			v572 = v512
																		}
																		v574 = int32(base.Ui32(v533) >> (uint(int32(1)) % 32))
																		if v559 == int32(0) {
																			if v542 == int32(16384) {
																				v599 = int32(1)
																			} else {
																				v580 = F_cmp_abs_common(m, v567, v569, v509, v572, v574, v532)
																				mBase = m.M
																				v599 = v580
																			}
																		} else {
																			if v542 == int32(0) {
																				v599 = int32(-1)
																			} else {
																				v584 = F_cmp_abs_common(m, v572, v574, v532, v567, v569, v509)
																				mBase = m.M
																				v599 = v584
																			}
																		}
																	}
																}
															}
														}
														if int32(0) <= v599 {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v608 = m.ExcPending
															if v608 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v614 = m.ExcPending
															if v614 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
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
								v61 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v61
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v61
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v61
								v68 = F_palloc(m, int32(12))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
									v71 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v71)
									*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
									v79 = v71
									v84 = v68 + int32(12)
									v88 = base.I64_extend_i32_u(v27)
									for {
										v91 = v84 - int32(2)
										v93 = base.I64_div_u_s(v88, int64(10000))
										v96 = v93*int64(55536) + v88
										*(*uint16)(unsafe.Add(mBase, uint32(v91))) = uint16(v96)
										v99 = v79 + int32(1)
										if base.Ui64(int64(9999)) < base.Ui64(v88) {
											v79 = v99
											v84 = v91
											v88 = v93
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v79
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v99
									*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v91
									v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
									v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
									v119 = int32(49152)
									v120 = v118 & v119
									if v120 == v119 {
										if v118 != int32(53248) {
											if v118 != int32(49152) {
												if v117 != int32(61440) {
													v139 = int32(-1)
												} else {
													v139 = int32(0)
												}
												v258 = v139
											} else {
												v258 = base.B2i32(v117 != int32(49152))
											}
										} else {
											if v117 == int32(49152) {
												v134 = int32(-1)
											} else {
												v134 = base.B2i32(v117 != int32(53248))
											}
											v258 = v134
										}
									} else {
										if base.Ui32(int32(49152)) <= base.Ui32(v117) {
											if v117 == int32(61440) {
												v146 = int32(1)
											} else {
												v146 = int32(-1)
											}
											v258 = v146
										} else {
											v148 = v22 + int32(6)
											v153 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
											if int32(0) <= base.I32_extend16_s(v118) {
												v154 = int32(-8)
											} else {
												v154 = int32(-6)
											}
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
											if int32(0) <= base.I32_extend16_s(v118) {
												v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
												v168 = v158
											} else {
												v168 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
											}
											v169 = v154 + int32(base.Ui32(v155)>>(uint(int32(2))%32))
											v171 = v25 + int32(6)
											v176 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
											if int32(0) <= base.I32_extend16_s(v117) {
												v177 = int32(-8)
											} else {
												v177 = int32(-6)
											}
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
											if int32(0) <= base.I32_extend16_s(v117) {
												v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v171))))
												v191 = v181
											} else {
												v191 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
											}
											v192 = v177 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
											v198 = v117 & int32(49152)
											if v198 == int32(32768) {
												v201 = v117 << (uint(int32(1)) % 32) & int32(16384)
											} else {
												v201 = v198
											}
											if base.Ui32(v169) <= base.Ui32(int32(1)) {
												if base.Ui32(v192) < base.Ui32(int32(2)) {
													v258 = int32(0)
												} else {
													if v201 == int32(16384) {
														v211 = int32(1)
													} else {
														v211 = int32(-1)
													}
													v258 = v211
												}
											} else {
												if v120 == int32(32768) {
													v218 = v118 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v218 = v120
												}
												if base.Ui32(v192) <= base.Ui32(int32(1)) {
													if v218 != 0 {
														v223 = int32(-1)
													} else {
														v223 = int32(1)
													}
													v258 = v223
												} else {
													if int32(0) <= base.I32_extend16_s(v118) {
														v226 = v22 + int32(8)
													} else {
														v226 = v148
													}
													v228 = int32(base.Ui32(v169) >> (uint(int32(1)) % 32))
													if int32(0) <= base.I32_extend16_s(v117) {
														v231 = v25 + int32(8)
													} else {
														v231 = v171
													}
													v233 = int32(base.Ui32(v192) >> (uint(int32(1)) % 32))
													if v218 == int32(0) {
														if v201 == int32(16384) {
															v258 = int32(1)
														} else {
															v239 = F_cmp_abs_common(m, v226, v228, v168, v231, v233, v191)
															mBase = m.M
															v258 = v239
														}
													} else {
														if v201 == int32(0) {
															v258 = int32(-1)
														} else {
															v243 = F_cmp_abs_common(m, v231, v233, v191, v226, v228, v168)
															mBase = m.M
															v258 = v243
														}
													}
												}
											}
										}
									}
									switch v258 {
									case 0:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v262 = m.ExcPending
										if v262 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v265 = m.ExcPending
											if v265 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(425786), int32(0))
												mBase = m.M
												v269 = m.ExcPending
												if v269 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(501522), int32(2010), int32(492002))
													mBase = m.M
													v274 = m.ExcPending
													if v274 != 0 {
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
										v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										v629 = int32(49152)
										v630 = v628 & v629
										if v630 == v629 {
											if v628 != int32(53248) {
												if v628 != int32(49152) {
													if v627 != int32(61440) {
														v649 = int32(-1)
													} else {
														v649 = int32(0)
													}
													v768 = v649
												} else {
													v768 = base.B2i32(v627 != int32(49152))
												}
											} else {
												if v627 == int32(49152) {
													v644 = int32(-1)
												} else {
													v644 = base.B2i32(v627 != int32(53248))
												}
												v768 = v644
											}
										} else {
											if base.Ui32(int32(49152)) <= base.Ui32(v627) {
												if v627 == int32(61440) {
													v656 = int32(1)
												} else {
													v656 = int32(-1)
												}
												v768 = v656
											} else {
												v658 = v17 + int32(6)
												v663 = base.B2i32(int32(0) <= base.I32_extend16_s(v628))
												if int32(0) <= base.I32_extend16_s(v628) {
													v664 = int32(-8)
												} else {
													v664 = int32(-6)
												}
												v665 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												if int32(0) <= base.I32_extend16_s(v628) {
													v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
													v678 = v668
												} else {
													v678 = v628<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v628&int32(63)
												}
												v679 = v664 + int32(base.Ui32(v665)>>(uint(int32(2))%32))
												v681 = v22 + int32(6)
												v686 = base.B2i32(int32(0) <= base.I32_extend16_s(v627))
												if int32(0) <= base.I32_extend16_s(v627) {
													v687 = int32(-8)
												} else {
													v687 = int32(-6)
												}
												v688 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												if int32(0) <= base.I32_extend16_s(v627) {
													v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681))))
													v701 = v691
												} else {
													v701 = v627<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v627&int32(63)
												}
												v702 = v687 + int32(base.Ui32(v688)>>(uint(int32(2))%32))
												v708 = v627 & int32(49152)
												if v708 == int32(32768) {
													v711 = v627 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v711 = v708
												}
												if base.Ui32(v679) <= base.Ui32(int32(1)) {
													if base.Ui32(v702) < base.Ui32(int32(2)) {
														v768 = int32(0)
													} else {
														if v711 == int32(16384) {
															v721 = int32(1)
														} else {
															v721 = int32(-1)
														}
														v768 = v721
													}
												} else {
													if v630 == int32(32768) {
														v728 = v628 << (uint(int32(1)) % 32) & int32(16384)
													} else {
														v728 = v630
													}
													if base.Ui32(v702) <= base.Ui32(int32(1)) {
														if v728 != 0 {
															v733 = int32(-1)
														} else {
															v733 = int32(1)
														}
														v768 = v733
													} else {
														if int32(0) <= base.I32_extend16_s(v628) {
															v736 = v17 + int32(8)
														} else {
															v736 = v658
														}
														v738 = int32(base.Ui32(v679) >> (uint(int32(1)) % 32))
														if int32(0) <= base.I32_extend16_s(v627) {
															v741 = v22 + int32(8)
														} else {
															v741 = v681
														}
														v743 = int32(base.Ui32(v702) >> (uint(int32(1)) % 32))
														if v728 == int32(0) {
															if v711 == int32(16384) {
																v768 = int32(1)
															} else {
																v749 = F_cmp_abs_common(m, v736, v738, v678, v741, v743, v701)
																mBase = m.M
																v768 = v749
															}
														} else {
															if v711 == int32(0) {
																v768 = int32(-1)
															} else {
																v753 = F_cmp_abs_common(m, v741, v743, v701, v736, v738, v678)
																mBase = m.M
																v768 = v753
															}
														}
													}
												}
											}
										}
										if int32(0) < v768 {
											v772 = F_palloc(m, int32(2))
											mBase = m.M
											v773 = m.ExcPending
											if v773 != 0 {
												return int32(0)
											} else {
												v774 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v772))) = uint16(v774)
												v777 = *(*int64)(unsafe.Add(mBase, _consts[985]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v777
												v780 = *(*int64)(unsafe.Add(mBase, _consts[986]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v780
												*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v772 + int32(2)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v772
												v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v961 = m.ExcPending
												if v961 != 0 {
													return int32(0)
												} else {
													if v960 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1033 = m.ExcPending
														if v1033 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1036 = m.ExcPending
															if v1036 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(403555), int32(0))
																mBase = m.M
																v1040 = m.ExcPending
																if v1040 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(501522), int32(2040), int32(492002))
																	mBase = m.M
																	v1045 = m.ExcPending
																	if v1045 != 0 {
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
														v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															F_pfree(m, v68)
															mBase = m.M
															v970 = m.ExcPending
															if v970 != 0 {
																return int32(0)
															} else {
																v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v971 != 0 {
																	F_pfree(m, v971)
																	mBase = m.M
																	v973 = m.ExcPending
																	if v973 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return base.I32_wrap_i64(v964)
																}
															}
														}
													}
												}
											}
										} else {
											v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											v800 = int32(49152)
											v801 = v799 & v800
											if v801 == v800 {
												if v799 != int32(53248) {
													if v799 != int32(49152) {
														if v798 != int32(61440) {
															v820 = int32(-1)
														} else {
															v820 = int32(0)
														}
														v939 = v820
													} else {
														v939 = base.B2i32(v798 != int32(49152))
													}
												} else {
													if v798 == int32(49152) {
														v815 = int32(-1)
													} else {
														v815 = base.B2i32(v798 != int32(53248))
													}
													v939 = v815
												}
											} else {
												if base.Ui32(int32(49152)) <= base.Ui32(v798) {
													if v798 == int32(61440) {
														v827 = int32(1)
													} else {
														v827 = int32(-1)
													}
													v939 = v827
												} else {
													v829 = v17 + int32(6)
													v834 = base.B2i32(int32(0) <= base.I32_extend16_s(v799))
													if int32(0) <= base.I32_extend16_s(v799) {
														v835 = int32(-8)
													} else {
														v835 = int32(-6)
													}
													v836 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													if int32(0) <= base.I32_extend16_s(v799) {
														v839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v829))))
														v849 = v839
													} else {
														v849 = v799<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v799&int32(63)
													}
													v850 = v835 + int32(base.Ui32(v836)>>(uint(int32(2))%32))
													v852 = v25 + int32(6)
													v857 = base.B2i32(int32(0) <= base.I32_extend16_s(v798))
													if int32(0) <= base.I32_extend16_s(v798) {
														v858 = int32(-8)
													} else {
														v858 = int32(-6)
													}
													v859 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													if int32(0) <= base.I32_extend16_s(v798) {
														v862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v852))))
														v872 = v862
													} else {
														v872 = v798<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v798&int32(63)
													}
													v873 = v858 + int32(base.Ui32(v859)>>(uint(int32(2))%32))
													v879 = v798 & int32(49152)
													if v879 == int32(32768) {
														v882 = v798 << (uint(int32(1)) % 32) & int32(16384)
													} else {
														v882 = v879
													}
													if base.Ui32(v850) <= base.Ui32(int32(1)) {
														if base.Ui32(v873) < base.Ui32(int32(2)) {
															v939 = int32(0)
														} else {
															if v882 == int32(16384) {
																v892 = int32(1)
															} else {
																v892 = int32(-1)
															}
															v939 = v892
														}
													} else {
														if v801 == int32(32768) {
															v899 = v799 << (uint(int32(1)) % 32) & int32(16384)
														} else {
															v899 = v801
														}
														if base.Ui32(v873) <= base.Ui32(int32(1)) {
															if v899 != 0 {
																v904 = int32(-1)
															} else {
																v904 = int32(1)
															}
															v939 = v904
														} else {
															if int32(0) <= base.I32_extend16_s(v799) {
																v907 = v17 + int32(8)
															} else {
																v907 = v829
															}
															v909 = int32(base.Ui32(v850) >> (uint(int32(1)) % 32))
															if int32(0) <= base.I32_extend16_s(v798) {
																v912 = v25 + int32(8)
															} else {
																v912 = v852
															}
															v914 = int32(base.Ui32(v873) >> (uint(int32(1)) % 32))
															if v899 == int32(0) {
																if v882 == int32(16384) {
																	v939 = int32(1)
																} else {
																	v920 = F_cmp_abs_common(m, v907, v909, v849, v912, v914, v872)
																	mBase = m.M
																	v939 = v920
																}
															} else {
																if v882 == int32(0) {
																	v939 = int32(-1)
																} else {
																	v924 = F_cmp_abs_common(m, v912, v914, v872, v907, v909, v849)
																	mBase = m.M
																	v939 = v924
																}
															}
														}
													}
												}
											}
											if v939 <= int32(0) {
												F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
												mBase = m.M
												v948 = m.ExcPending
												if v948 != 0 {
													return int32(0)
												} else {
													v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v961 = m.ExcPending
													if v961 != 0 {
														return int32(0)
													} else {
														if v960 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1036 = m.ExcPending
																	if v1036 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(403555), int32(0))
																		mBase = m.M
																		v1040 = m.ExcPending
																		if v1040 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(501522), int32(2040), int32(492002))
																			mBase = m.M
																			v1045 = m.ExcPending
																			if v1045 != 0 {
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
																F_pfree(m, v68)
																mBase = m.M
																v970 = m.ExcPending
																if v970 != 0 {
																	return int32(0)
																} else {
																	v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v971 != 0 {
																		F_pfree(m, v971)
																		mBase = m.M
																		v973 = m.ExcPending
																		if v973 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v964)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
																	}
																}
															}
														}
													}
												}
											} else {
												F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
												mBase = m.M
												v954 = m.ExcPending
												if v954 != 0 {
													return int32(0)
												} else {
													v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v961 = m.ExcPending
													if v961 != 0 {
														return int32(0)
													} else {
														if v960 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1036 = m.ExcPending
																	if v1036 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(403555), int32(0))
																		mBase = m.M
																		v1040 = m.ExcPending
																		if v1040 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(501522), int32(2040), int32(492002))
																			mBase = m.M
																			v1045 = m.ExcPending
																			if v1045 != 0 {
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
																F_pfree(m, v68)
																mBase = m.M
																v970 = m.ExcPending
																if v970 != 0 {
																	return int32(0)
																} else {
																	v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v971 != 0 {
																		F_pfree(m, v971)
																		mBase = m.M
																		v973 = m.ExcPending
																		if v973 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v964)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
																	}
																}
															}
														}
													}
												}
											}
										}
									default:
										v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
										v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
										v289 = int32(49152)
										v290 = v288 & v289
										if v290 == v289 {
											if v288 != int32(53248) {
												if v288 != int32(49152) {
													if v287 != int32(61440) {
														v309 = int32(-1)
													} else {
														v309 = int32(0)
													}
													v428 = v309
												} else {
													v428 = base.B2i32(v287 != int32(49152))
												}
											} else {
												if v287 == int32(49152) {
													v304 = int32(-1)
												} else {
													v304 = base.B2i32(v287 != int32(53248))
												}
												v428 = v304
											}
										} else {
											if base.Ui32(int32(49152)) <= base.Ui32(v287) {
												if v287 == int32(61440) {
													v316 = int32(1)
												} else {
													v316 = int32(-1)
												}
												v428 = v316
											} else {
												v318 = v17 + int32(6)
												v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v288))
												if int32(0) <= base.I32_extend16_s(v288) {
													v324 = int32(-8)
												} else {
													v324 = int32(-6)
												}
												v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												if int32(0) <= base.I32_extend16_s(v288) {
													v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v318))))
													v338 = v328
												} else {
													v338 = v288<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v288&int32(63)
												}
												v339 = v324 + int32(base.Ui32(v325)>>(uint(int32(2))%32))
												v341 = v22 + int32(6)
												v346 = base.B2i32(int32(0) <= base.I32_extend16_s(v287))
												if int32(0) <= base.I32_extend16_s(v287) {
													v347 = int32(-8)
												} else {
													v347 = int32(-6)
												}
												v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
												if int32(0) <= base.I32_extend16_s(v287) {
													v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341))))
													v361 = v351
												} else {
													v361 = v287<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v287&int32(63)
												}
												v362 = v347 + int32(base.Ui32(v348)>>(uint(int32(2))%32))
												v368 = v287 & int32(49152)
												if v368 == int32(32768) {
													v371 = v287 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v371 = v368
												}
												if base.Ui32(v339) <= base.Ui32(int32(1)) {
													if base.Ui32(v362) < base.Ui32(int32(2)) {
														v428 = int32(0)
													} else {
														if v371 == int32(16384) {
															v381 = int32(1)
														} else {
															v381 = int32(-1)
														}
														v428 = v381
													}
												} else {
													if v290 == int32(32768) {
														v388 = v288 << (uint(int32(1)) % 32) & int32(16384)
													} else {
														v388 = v290
													}
													if base.Ui32(v362) <= base.Ui32(int32(1)) {
														if v388 != 0 {
															v393 = int32(-1)
														} else {
															v393 = int32(1)
														}
														v428 = v393
													} else {
														if int32(0) <= base.I32_extend16_s(v288) {
															v396 = v17 + int32(8)
														} else {
															v396 = v318
														}
														v398 = int32(base.Ui32(v339) >> (uint(int32(1)) % 32))
														if int32(0) <= base.I32_extend16_s(v287) {
															v401 = v22 + int32(8)
														} else {
															v401 = v341
														}
														v403 = int32(base.Ui32(v362) >> (uint(int32(1)) % 32))
														if v388 == int32(0) {
															if v371 == int32(16384) {
																v428 = int32(1)
															} else {
																v409 = F_cmp_abs_common(m, v396, v398, v338, v401, v403, v361)
																mBase = m.M
																v428 = v409
															}
														} else {
															if v371 == int32(0) {
																v428 = int32(-1)
															} else {
																v413 = F_cmp_abs_common(m, v401, v403, v361, v396, v398, v338)
																mBase = m.M
																v428 = v413
															}
														}
													}
												}
											}
										}
										if v428 < int32(0) {
											v432 = F_palloc(m, int32(2))
											mBase = m.M
											v433 = m.ExcPending
											if v433 != 0 {
												return int32(0)
											} else {
												v434 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v432))) = uint16(v434)
												v437 = *(*int64)(unsafe.Add(mBase, _consts[985]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v437
												v440 = *(*int64)(unsafe.Add(mBase, _consts[986]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v440
												*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v432 + int32(2)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v432
												v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
												mBase = m.M
												v961 = m.ExcPending
												if v961 != 0 {
													return int32(0)
												} else {
													if v960 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v1033 = m.ExcPending
														if v1033 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v1036 = m.ExcPending
															if v1036 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(403555), int32(0))
																mBase = m.M
																v1040 = m.ExcPending
																if v1040 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(501522), int32(2040), int32(492002))
																	mBase = m.M
																	v1045 = m.ExcPending
																	if v1045 != 0 {
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
														v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
														if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															F_pfree(m, v68)
															mBase = m.M
															v970 = m.ExcPending
															if v970 != 0 {
																return int32(0)
															} else {
																v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																if v971 != 0 {
																	F_pfree(m, v971)
																	mBase = m.M
																	v973 = m.ExcPending
																	if v973 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
																	}
																} else {
																	m.G0 = v14 - int32(-64)
																	return base.I32_wrap_i64(v964)
																}
															}
														}
													}
												}
											}
										} else {
											v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
											v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											v460 = int32(49152)
											v461 = v459 & v460
											if v461 == v460 {
												if v459 != int32(53248) {
													if v459 != int32(49152) {
														if v458 != int32(61440) {
															v480 = int32(-1)
														} else {
															v480 = int32(0)
														}
														v599 = v480
													} else {
														v599 = base.B2i32(v458 != int32(49152))
													}
												} else {
													if v458 == int32(49152) {
														v475 = int32(-1)
													} else {
														v475 = base.B2i32(v458 != int32(53248))
													}
													v599 = v475
												}
											} else {
												if base.Ui32(int32(49152)) <= base.Ui32(v458) {
													if v458 == int32(61440) {
														v487 = int32(1)
													} else {
														v487 = int32(-1)
													}
													v599 = v487
												} else {
													v489 = v17 + int32(6)
													v494 = base.B2i32(int32(0) <= base.I32_extend16_s(v459))
													if int32(0) <= base.I32_extend16_s(v459) {
														v495 = int32(-8)
													} else {
														v495 = int32(-6)
													}
													v496 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													if int32(0) <= base.I32_extend16_s(v459) {
														v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489))))
														v509 = v499
													} else {
														v509 = v459<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v459&int32(63)
													}
													v510 = v495 + int32(base.Ui32(v496)>>(uint(int32(2))%32))
													v512 = v25 + int32(6)
													v517 = base.B2i32(int32(0) <= base.I32_extend16_s(v458))
													if int32(0) <= base.I32_extend16_s(v458) {
														v518 = int32(-8)
													} else {
														v518 = int32(-6)
													}
													v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
													if int32(0) <= base.I32_extend16_s(v458) {
														v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v512))))
														v532 = v522
													} else {
														v532 = v458<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v458&int32(63)
													}
													v533 = v518 + int32(base.Ui32(v519)>>(uint(int32(2))%32))
													v539 = v458 & int32(49152)
													if v539 == int32(32768) {
														v542 = v458 << (uint(int32(1)) % 32) & int32(16384)
													} else {
														v542 = v539
													}
													if base.Ui32(v510) <= base.Ui32(int32(1)) {
														if base.Ui32(v533) < base.Ui32(int32(2)) {
															v599 = int32(0)
														} else {
															if v542 == int32(16384) {
																v552 = int32(1)
															} else {
																v552 = int32(-1)
															}
															v599 = v552
														}
													} else {
														if v461 == int32(32768) {
															v559 = v459 << (uint(int32(1)) % 32) & int32(16384)
														} else {
															v559 = v461
														}
														if base.Ui32(v533) <= base.Ui32(int32(1)) {
															if v559 != 0 {
																v564 = int32(-1)
															} else {
																v564 = int32(1)
															}
															v599 = v564
														} else {
															if int32(0) <= base.I32_extend16_s(v459) {
																v567 = v17 + int32(8)
															} else {
																v567 = v489
															}
															v569 = int32(base.Ui32(v510) >> (uint(int32(1)) % 32))
															if int32(0) <= base.I32_extend16_s(v458) {
																v572 = v25 + int32(8)
															} else {
																v572 = v512
															}
															v574 = int32(base.Ui32(v533) >> (uint(int32(1)) % 32))
															if v559 == int32(0) {
																if v542 == int32(16384) {
																	v599 = int32(1)
																} else {
																	v580 = F_cmp_abs_common(m, v567, v569, v509, v572, v574, v532)
																	mBase = m.M
																	v599 = v580
																}
															} else {
																if v542 == int32(0) {
																	v599 = int32(-1)
																} else {
																	v584 = F_cmp_abs_common(m, v572, v574, v532, v567, v569, v509)
																	mBase = m.M
																	v599 = v584
																}
															}
														}
													}
												}
											}
											if int32(0) <= v599 {
												F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
												mBase = m.M
												v608 = m.ExcPending
												if v608 != 0 {
													return int32(0)
												} else {
													v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v961 = m.ExcPending
													if v961 != 0 {
														return int32(0)
													} else {
														if v960 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1036 = m.ExcPending
																	if v1036 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(403555), int32(0))
																		mBase = m.M
																		v1040 = m.ExcPending
																		if v1040 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(501522), int32(2040), int32(492002))
																			mBase = m.M
																			v1045 = m.ExcPending
																			if v1045 != 0 {
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
																F_pfree(m, v68)
																mBase = m.M
																v970 = m.ExcPending
																if v970 != 0 {
																	return int32(0)
																} else {
																	v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v971 != 0 {
																		F_pfree(m, v971)
																		mBase = m.M
																		v973 = m.ExcPending
																		if v973 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v964)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
																	}
																}
															}
														}
													}
												}
											} else {
												F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
												mBase = m.M
												v614 = m.ExcPending
												if v614 != 0 {
													return int32(0)
												} else {
													v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
													mBase = m.M
													v961 = m.ExcPending
													if v961 != 0 {
														return int32(0)
													} else {
														if v960 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v1033 = m.ExcPending
															if v1033 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v1036 = m.ExcPending
																if v1036 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(403555), int32(0))
																	mBase = m.M
																	v1040 = m.ExcPending
																	if v1040 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(501522), int32(2040), int32(492002))
																		mBase = m.M
																		v1045 = m.ExcPending
																		if v1045 != 0 {
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
															v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
															if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v1033 = m.ExcPending
																if v1033 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v1036 = m.ExcPending
																	if v1036 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(403555), int32(0))
																		mBase = m.M
																		v1040 = m.ExcPending
																		if v1040 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(501522), int32(2040), int32(492002))
																			mBase = m.M
																			v1045 = m.ExcPending
																			if v1045 != 0 {
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
																F_pfree(m, v68)
																mBase = m.M
																v970 = m.ExcPending
																if v970 != 0 {
																	return int32(0)
																} else {
																	v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																	if v971 != 0 {
																		F_pfree(m, v971)
																		mBase = m.M
																		v973 = m.ExcPending
																		if v973 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v14 - int32(-64)
																			return base.I32_wrap_i64(v964)
																		}
																	} else {
																		m.G0 = v14 - int32(-64)
																		return base.I32_wrap_i64(v964)
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
						if v30 == int32(49152) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v1000 = m.ExcPending
							if v1000 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(386138242))
								mBase = m.M
								v1003 = m.ExcPending
								if v1003 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(529241), int32(0))
									mBase = m.M
									v1007 = m.ExcPending
									if v1007 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(501522), int32(1991), int32(492002))
										mBase = m.M
										v1012 = m.ExcPending
										if v1012 != 0 {
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
							if v42&int32(65535) == int32(49152) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v1000 = m.ExcPending
								if v1000 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(386138242))
									mBase = m.M
									v1003 = m.ExcPending
									if v1003 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(529241), int32(0))
										mBase = m.M
										v1007 = m.ExcPending
										if v1007 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(501522), int32(1991), int32(492002))
											mBase = m.M
											v1012 = m.ExcPending
											if v1012 != 0 {
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
								if v48 == int32(49152) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1000 = m.ExcPending
									if v1000 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(386138242))
										mBase = m.M
										v1003 = m.ExcPending
										if v1003 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(529241), int32(0))
											mBase = m.M
											v1007 = m.ExcPending
											if v1007 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(501522), int32(1991), int32(492002))
												mBase = m.M
												v1012 = m.ExcPending
												if v1012 != 0 {
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
									if v47&int32(57343) == int32(53248) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v1016 = m.ExcPending
										if v1016 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(386138242))
											mBase = m.M
											v1019 = m.ExcPending
											if v1019 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(351481), int32(0))
												mBase = m.M
												v1023 = m.ExcPending
												if v1023 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(501522), int32(1996), int32(492002))
													mBase = m.M
													v1028 = m.ExcPending
													if v1028 != 0 {
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
										if v48&int32(57343) == int32(53248) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v1016 = m.ExcPending
											if v1016 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(386138242))
												mBase = m.M
												v1019 = m.ExcPending
												if v1019 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(351481), int32(0))
													mBase = m.M
													v1023 = m.ExcPending
													if v1023 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(501522), int32(1996), int32(492002))
														mBase = m.M
														v1028 = m.ExcPending
														if v1028 != 0 {
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
											v61 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v61
											*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v61
											v68 = F_palloc(m, int32(12))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v68
												v71 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v68))) = uint16(v71)
												*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(0)
												v79 = v71
												v84 = v68 + int32(12)
												v88 = base.I64_extend_i32_u(v27)
												for {
													v91 = v84 - int32(2)
													v93 = base.I64_div_u_s(v88, int64(10000))
													v96 = v93*int64(55536) + v88
													*(*uint16)(unsafe.Add(mBase, uint32(v91))) = uint16(v96)
													v99 = v79 + int32(1)
													if base.Ui64(int64(9999)) < base.Ui64(v88) {
														v79 = v99
														v84 = v91
														v88 = v93
														continue
													} else {
														break
													}
													break
												}
												*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v79
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v99
												*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v91
												v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
												v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
												v119 = int32(49152)
												v120 = v118 & v119
												if v120 == v119 {
													if v118 != int32(53248) {
														if v118 != int32(49152) {
															if v117 != int32(61440) {
																v139 = int32(-1)
															} else {
																v139 = int32(0)
															}
															v258 = v139
														} else {
															v258 = base.B2i32(v117 != int32(49152))
														}
													} else {
														if v117 == int32(49152) {
															v134 = int32(-1)
														} else {
															v134 = base.B2i32(v117 != int32(53248))
														}
														v258 = v134
													}
												} else {
													if base.Ui32(int32(49152)) <= base.Ui32(v117) {
														if v117 == int32(61440) {
															v146 = int32(1)
														} else {
															v146 = int32(-1)
														}
														v258 = v146
													} else {
														v148 = v22 + int32(6)
														v153 = base.B2i32(int32(0) <= base.I32_extend16_s(v118))
														if int32(0) <= base.I32_extend16_s(v118) {
															v154 = int32(-8)
														} else {
															v154 = int32(-6)
														}
														v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
														if int32(0) <= base.I32_extend16_s(v118) {
															v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v148))))
															v168 = v158
														} else {
															v168 = v118<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v118&int32(63)
														}
														v169 = v154 + int32(base.Ui32(v155)>>(uint(int32(2))%32))
														v171 = v25 + int32(6)
														v176 = base.B2i32(int32(0) <= base.I32_extend16_s(v117))
														if int32(0) <= base.I32_extend16_s(v117) {
															v177 = int32(-8)
														} else {
															v177 = int32(-6)
														}
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
														if int32(0) <= base.I32_extend16_s(v117) {
															v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v171))))
															v191 = v181
														} else {
															v191 = v117<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v117&int32(63)
														}
														v192 = v177 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
														v198 = v117 & int32(49152)
														if v198 == int32(32768) {
															v201 = v117 << (uint(int32(1)) % 32) & int32(16384)
														} else {
															v201 = v198
														}
														if base.Ui32(v169) <= base.Ui32(int32(1)) {
															if base.Ui32(v192) < base.Ui32(int32(2)) {
																v258 = int32(0)
															} else {
																if v201 == int32(16384) {
																	v211 = int32(1)
																} else {
																	v211 = int32(-1)
																}
																v258 = v211
															}
														} else {
															if v120 == int32(32768) {
																v218 = v118 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v218 = v120
															}
															if base.Ui32(v192) <= base.Ui32(int32(1)) {
																if v218 != 0 {
																	v223 = int32(-1)
																} else {
																	v223 = int32(1)
																}
																v258 = v223
															} else {
																if int32(0) <= base.I32_extend16_s(v118) {
																	v226 = v22 + int32(8)
																} else {
																	v226 = v148
																}
																v228 = int32(base.Ui32(v169) >> (uint(int32(1)) % 32))
																if int32(0) <= base.I32_extend16_s(v117) {
																	v231 = v25 + int32(8)
																} else {
																	v231 = v171
																}
																v233 = int32(base.Ui32(v192) >> (uint(int32(1)) % 32))
																if v218 == int32(0) {
																	if v201 == int32(16384) {
																		v258 = int32(1)
																	} else {
																		v239 = F_cmp_abs_common(m, v226, v228, v168, v231, v233, v191)
																		mBase = m.M
																		v258 = v239
																	}
																} else {
																	if v201 == int32(0) {
																		v258 = int32(-1)
																	} else {
																		v243 = F_cmp_abs_common(m, v231, v233, v191, v226, v228, v168)
																		mBase = m.M
																		v258 = v243
																	}
																}
															}
														}
													}
												}
												switch v258 {
												case 0:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v262 = m.ExcPending
													if v262 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(386138242))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(425786), int32(0))
															mBase = m.M
															v269 = m.ExcPending
															if v269 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(501522), int32(2010), int32(492002))
																mBase = m.M
																v274 = m.ExcPending
																if v274 != 0 {
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
													v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v629 = int32(49152)
													v630 = v628 & v629
													if v630 == v629 {
														if v628 != int32(53248) {
															if v628 != int32(49152) {
																if v627 != int32(61440) {
																	v649 = int32(-1)
																} else {
																	v649 = int32(0)
																}
																v768 = v649
															} else {
																v768 = base.B2i32(v627 != int32(49152))
															}
														} else {
															if v627 == int32(49152) {
																v644 = int32(-1)
															} else {
																v644 = base.B2i32(v627 != int32(53248))
															}
															v768 = v644
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v627) {
															if v627 == int32(61440) {
																v656 = int32(1)
															} else {
																v656 = int32(-1)
															}
															v768 = v656
														} else {
															v658 = v17 + int32(6)
															v663 = base.B2i32(int32(0) <= base.I32_extend16_s(v628))
															if int32(0) <= base.I32_extend16_s(v628) {
																v664 = int32(-8)
															} else {
																v664 = int32(-6)
															}
															v665 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v628) {
																v668 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
																v678 = v668
															} else {
																v678 = v628<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v628&int32(63)
															}
															v679 = v664 + int32(base.Ui32(v665)>>(uint(int32(2))%32))
															v681 = v22 + int32(6)
															v686 = base.B2i32(int32(0) <= base.I32_extend16_s(v627))
															if int32(0) <= base.I32_extend16_s(v627) {
																v687 = int32(-8)
															} else {
																v687 = int32(-6)
															}
															v688 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v627) {
																v691 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681))))
																v701 = v691
															} else {
																v701 = v627<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v627&int32(63)
															}
															v702 = v687 + int32(base.Ui32(v688)>>(uint(int32(2))%32))
															v708 = v627 & int32(49152)
															if v708 == int32(32768) {
																v711 = v627 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v711 = v708
															}
															if base.Ui32(v679) <= base.Ui32(int32(1)) {
																if base.Ui32(v702) < base.Ui32(int32(2)) {
																	v768 = int32(0)
																} else {
																	if v711 == int32(16384) {
																		v721 = int32(1)
																	} else {
																		v721 = int32(-1)
																	}
																	v768 = v721
																}
															} else {
																if v630 == int32(32768) {
																	v728 = v628 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v728 = v630
																}
																if base.Ui32(v702) <= base.Ui32(int32(1)) {
																	if v728 != 0 {
																		v733 = int32(-1)
																	} else {
																		v733 = int32(1)
																	}
																	v768 = v733
																} else {
																	if int32(0) <= base.I32_extend16_s(v628) {
																		v736 = v17 + int32(8)
																	} else {
																		v736 = v658
																	}
																	v738 = int32(base.Ui32(v679) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v627) {
																		v741 = v22 + int32(8)
																	} else {
																		v741 = v681
																	}
																	v743 = int32(base.Ui32(v702) >> (uint(int32(1)) % 32))
																	if v728 == int32(0) {
																		if v711 == int32(16384) {
																			v768 = int32(1)
																		} else {
																			v749 = F_cmp_abs_common(m, v736, v738, v678, v741, v743, v701)
																			mBase = m.M
																			v768 = v749
																		}
																	} else {
																		if v711 == int32(0) {
																			v768 = int32(-1)
																		} else {
																			v753 = F_cmp_abs_common(m, v741, v743, v701, v736, v738, v678)
																			mBase = m.M
																			v768 = v753
																		}
																	}
																}
															}
														}
													}
													if int32(0) < v768 {
														v772 = F_palloc(m, int32(2))
														mBase = m.M
														v773 = m.ExcPending
														if v773 != 0 {
															return int32(0)
														} else {
															v774 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v772))) = uint16(v774)
															v777 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v777
															v780 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v780
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v772 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v772
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v800 = int32(49152)
														v801 = v799 & v800
														if v801 == v800 {
															if v799 != int32(53248) {
																if v799 != int32(49152) {
																	if v798 != int32(61440) {
																		v820 = int32(-1)
																	} else {
																		v820 = int32(0)
																	}
																	v939 = v820
																} else {
																	v939 = base.B2i32(v798 != int32(49152))
																}
															} else {
																if v798 == int32(49152) {
																	v815 = int32(-1)
																} else {
																	v815 = base.B2i32(v798 != int32(53248))
																}
																v939 = v815
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v798) {
																if v798 == int32(61440) {
																	v827 = int32(1)
																} else {
																	v827 = int32(-1)
																}
																v939 = v827
															} else {
																v829 = v17 + int32(6)
																v834 = base.B2i32(int32(0) <= base.I32_extend16_s(v799))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v835 = int32(-8)
																} else {
																	v835 = int32(-6)
																}
																v836 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v799) {
																	v839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v829))))
																	v849 = v839
																} else {
																	v849 = v799<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v799&int32(63)
																}
																v850 = v835 + int32(base.Ui32(v836)>>(uint(int32(2))%32))
																v852 = v25 + int32(6)
																v857 = base.B2i32(int32(0) <= base.I32_extend16_s(v798))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v858 = int32(-8)
																} else {
																	v858 = int32(-6)
																}
																v859 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v798) {
																	v862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v852))))
																	v872 = v862
																} else {
																	v872 = v798<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v798&int32(63)
																}
																v873 = v858 + int32(base.Ui32(v859)>>(uint(int32(2))%32))
																v879 = v798 & int32(49152)
																if v879 == int32(32768) {
																	v882 = v798 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v882 = v879
																}
																if base.Ui32(v850) <= base.Ui32(int32(1)) {
																	if base.Ui32(v873) < base.Ui32(int32(2)) {
																		v939 = int32(0)
																	} else {
																		if v882 == int32(16384) {
																			v892 = int32(1)
																		} else {
																			v892 = int32(-1)
																		}
																		v939 = v892
																	}
																} else {
																	if v801 == int32(32768) {
																		v899 = v799 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v899 = v801
																	}
																	if base.Ui32(v873) <= base.Ui32(int32(1)) {
																		if v899 != 0 {
																			v904 = int32(-1)
																		} else {
																			v904 = int32(1)
																		}
																		v939 = v904
																	} else {
																		if int32(0) <= base.I32_extend16_s(v799) {
																			v907 = v17 + int32(8)
																		} else {
																			v907 = v829
																		}
																		v909 = int32(base.Ui32(v850) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v798) {
																			v912 = v25 + int32(8)
																		} else {
																			v912 = v852
																		}
																		v914 = int32(base.Ui32(v873) >> (uint(int32(1)) % 32))
																		if v899 == int32(0) {
																			if v882 == int32(16384) {
																				v939 = int32(1)
																			} else {
																				v920 = F_cmp_abs_common(m, v907, v909, v849, v912, v914, v872)
																				mBase = m.M
																				v939 = v920
																			}
																		} else {
																			if v882 == int32(0) {
																				v939 = int32(-1)
																			} else {
																				v924 = F_cmp_abs_common(m, v912, v914, v872, v907, v909, v849)
																				mBase = m.M
																				v939 = v924
																			}
																		}
																	}
																}
															}
														}
														if v939 <= int32(0) {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v948 = m.ExcPending
															if v948 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v954 = m.ExcPending
															if v954 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												default:
													v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
													v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													v289 = int32(49152)
													v290 = v288 & v289
													if v290 == v289 {
														if v288 != int32(53248) {
															if v288 != int32(49152) {
																if v287 != int32(61440) {
																	v309 = int32(-1)
																} else {
																	v309 = int32(0)
																}
																v428 = v309
															} else {
																v428 = base.B2i32(v287 != int32(49152))
															}
														} else {
															if v287 == int32(49152) {
																v304 = int32(-1)
															} else {
																v304 = base.B2i32(v287 != int32(53248))
															}
															v428 = v304
														}
													} else {
														if base.Ui32(int32(49152)) <= base.Ui32(v287) {
															if v287 == int32(61440) {
																v316 = int32(1)
															} else {
																v316 = int32(-1)
															}
															v428 = v316
														} else {
															v318 = v17 + int32(6)
															v323 = base.B2i32(int32(0) <= base.I32_extend16_s(v288))
															if int32(0) <= base.I32_extend16_s(v288) {
																v324 = int32(-8)
															} else {
																v324 = int32(-6)
															}
															v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															if int32(0) <= base.I32_extend16_s(v288) {
																v328 = int32(*(*int16)(unsafe.Add(mBase, uint32(v318))))
																v338 = v328
															} else {
																v338 = v288<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v288&int32(63)
															}
															v339 = v324 + int32(base.Ui32(v325)>>(uint(int32(2))%32))
															v341 = v22 + int32(6)
															v346 = base.B2i32(int32(0) <= base.I32_extend16_s(v287))
															if int32(0) <= base.I32_extend16_s(v287) {
																v347 = int32(-8)
															} else {
																v347 = int32(-6)
															}
															v348 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
															if int32(0) <= base.I32_extend16_s(v287) {
																v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341))))
																v361 = v351
															} else {
																v361 = v287<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v287&int32(63)
															}
															v362 = v347 + int32(base.Ui32(v348)>>(uint(int32(2))%32))
															v368 = v287 & int32(49152)
															if v368 == int32(32768) {
																v371 = v287 << (uint(int32(1)) % 32) & int32(16384)
															} else {
																v371 = v368
															}
															if base.Ui32(v339) <= base.Ui32(int32(1)) {
																if base.Ui32(v362) < base.Ui32(int32(2)) {
																	v428 = int32(0)
																} else {
																	if v371 == int32(16384) {
																		v381 = int32(1)
																	} else {
																		v381 = int32(-1)
																	}
																	v428 = v381
																}
															} else {
																if v290 == int32(32768) {
																	v388 = v288 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v388 = v290
																}
																if base.Ui32(v362) <= base.Ui32(int32(1)) {
																	if v388 != 0 {
																		v393 = int32(-1)
																	} else {
																		v393 = int32(1)
																	}
																	v428 = v393
																} else {
																	if int32(0) <= base.I32_extend16_s(v288) {
																		v396 = v17 + int32(8)
																	} else {
																		v396 = v318
																	}
																	v398 = int32(base.Ui32(v339) >> (uint(int32(1)) % 32))
																	if int32(0) <= base.I32_extend16_s(v287) {
																		v401 = v22 + int32(8)
																	} else {
																		v401 = v341
																	}
																	v403 = int32(base.Ui32(v362) >> (uint(int32(1)) % 32))
																	if v388 == int32(0) {
																		if v371 == int32(16384) {
																			v428 = int32(1)
																		} else {
																			v409 = F_cmp_abs_common(m, v396, v398, v338, v401, v403, v361)
																			mBase = m.M
																			v428 = v409
																		}
																	} else {
																		if v371 == int32(0) {
																			v428 = int32(-1)
																		} else {
																			v413 = F_cmp_abs_common(m, v401, v403, v361, v396, v398, v338)
																			mBase = m.M
																			v428 = v413
																		}
																	}
																}
															}
														}
													}
													if v428 < int32(0) {
														v432 = F_palloc(m, int32(2))
														mBase = m.M
														v433 = m.ExcPending
														if v433 != 0 {
															return int32(0)
														} else {
															v434 = int32(0)
															*(*uint16)(unsafe.Add(mBase, uint32(v432))) = uint16(v434)
															v437 = *(*int64)(unsafe.Add(mBase, _consts[985]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v437
															v440 = *(*int64)(unsafe.Add(mBase, _consts[986]))
															*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v440
															*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v432 + int32(2)
															*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v432
															v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
															mBase = m.M
															v961 = m.ExcPending
															if v961 != 0 {
																return int32(0)
															} else {
																if v960 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v1033 = m.ExcPending
																	if v1033 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(50331778))
																		mBase = m.M
																		v1036 = m.ExcPending
																		if v1036 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(403555), int32(0))
																			mBase = m.M
																			v1040 = m.ExcPending
																			if v1040 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(501522), int32(2040), int32(492002))
																				mBase = m.M
																				v1045 = m.ExcPending
																				if v1045 != 0 {
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
																	v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																	if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		F_pfree(m, v68)
																		mBase = m.M
																		v970 = m.ExcPending
																		if v970 != 0 {
																			return int32(0)
																		} else {
																			v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																			if v971 != 0 {
																				F_pfree(m, v971)
																				mBase = m.M
																				v973 = m.ExcPending
																				if v973 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			} else {
																				m.G0 = v14 - int32(-64)
																				return base.I32_wrap_i64(v964)
																			}
																		}
																	}
																}
															}
														}
													} else {
														v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
														v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
														v460 = int32(49152)
														v461 = v459 & v460
														if v461 == v460 {
															if v459 != int32(53248) {
																if v459 != int32(49152) {
																	if v458 != int32(61440) {
																		v480 = int32(-1)
																	} else {
																		v480 = int32(0)
																	}
																	v599 = v480
																} else {
																	v599 = base.B2i32(v458 != int32(49152))
																}
															} else {
																if v458 == int32(49152) {
																	v475 = int32(-1)
																} else {
																	v475 = base.B2i32(v458 != int32(53248))
																}
																v599 = v475
															}
														} else {
															if base.Ui32(int32(49152)) <= base.Ui32(v458) {
																if v458 == int32(61440) {
																	v487 = int32(1)
																} else {
																	v487 = int32(-1)
																}
																v599 = v487
															} else {
																v489 = v17 + int32(6)
																v494 = base.B2i32(int32(0) <= base.I32_extend16_s(v459))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v495 = int32(-8)
																} else {
																	v495 = int32(-6)
																}
																v496 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																if int32(0) <= base.I32_extend16_s(v459) {
																	v499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v489))))
																	v509 = v499
																} else {
																	v509 = v459<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v459&int32(63)
																}
																v510 = v495 + int32(base.Ui32(v496)>>(uint(int32(2))%32))
																v512 = v25 + int32(6)
																v517 = base.B2i32(int32(0) <= base.I32_extend16_s(v458))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v518 = int32(-8)
																} else {
																	v518 = int32(-6)
																}
																v519 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
																if int32(0) <= base.I32_extend16_s(v458) {
																	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v512))))
																	v532 = v522
																} else {
																	v532 = v458<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v458&int32(63)
																}
																v533 = v518 + int32(base.Ui32(v519)>>(uint(int32(2))%32))
																v539 = v458 & int32(49152)
																if v539 == int32(32768) {
																	v542 = v458 << (uint(int32(1)) % 32) & int32(16384)
																} else {
																	v542 = v539
																}
																if base.Ui32(v510) <= base.Ui32(int32(1)) {
																	if base.Ui32(v533) < base.Ui32(int32(2)) {
																		v599 = int32(0)
																	} else {
																		if v542 == int32(16384) {
																			v552 = int32(1)
																		} else {
																			v552 = int32(-1)
																		}
																		v599 = v552
																	}
																} else {
																	if v461 == int32(32768) {
																		v559 = v459 << (uint(int32(1)) % 32) & int32(16384)
																	} else {
																		v559 = v461
																	}
																	if base.Ui32(v533) <= base.Ui32(int32(1)) {
																		if v559 != 0 {
																			v564 = int32(-1)
																		} else {
																			v564 = int32(1)
																		}
																		v599 = v564
																	} else {
																		if int32(0) <= base.I32_extend16_s(v459) {
																			v567 = v17 + int32(8)
																		} else {
																			v567 = v489
																		}
																		v569 = int32(base.Ui32(v510) >> (uint(int32(1)) % 32))
																		if int32(0) <= base.I32_extend16_s(v458) {
																			v572 = v25 + int32(8)
																		} else {
																			v572 = v512
																		}
																		v574 = int32(base.Ui32(v533) >> (uint(int32(1)) % 32))
																		if v559 == int32(0) {
																			if v542 == int32(16384) {
																				v599 = int32(1)
																			} else {
																				v580 = F_cmp_abs_common(m, v567, v569, v509, v572, v574, v532)
																				mBase = m.M
																				v599 = v580
																			}
																		} else {
																			if v542 == int32(0) {
																				v599 = int32(-1)
																			} else {
																				v584 = F_cmp_abs_common(m, v572, v574, v532, v567, v569, v509)
																				mBase = m.M
																				v599 = v584
																			}
																		}
																	}
																}
															}
														}
														if int32(0) <= v599 {
															F_add_var(m, v12+int32(-32), int32(1741928), v12+int32(-56))
															mBase = m.M
															v608 = m.ExcPending
															if v608 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_compute_bucket(m, v17, v22, v25, v12+int32(-32), v12+int32(-56))
															mBase = m.M
															v614 = m.ExcPending
															if v614 != 0 {
																return int32(0)
															} else {
																v960 = F_numericvar_to_int64(m, v12+int32(-56), v12+int32(-8))
																mBase = m.M
																v961 = m.ExcPending
																if v961 != 0 {
																	return int32(0)
																} else {
																	if v960 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v1033 = m.ExcPending
																		if v1033 != 0 {
																			return int32(0)
																		} else {
																			F_errcode(m, int32(50331778))
																			mBase = m.M
																			v1036 = m.ExcPending
																			if v1036 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(403555), int32(0))
																				mBase = m.M
																				v1040 = m.ExcPending
																				if v1040 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(501522), int32(2040), int32(492002))
																					mBase = m.M
																					v1045 = m.ExcPending
																					if v1045 != 0 {
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
																		v964 = *(*int64)(unsafe.Add(mBase, uint32(v14)+56))
																		if base.Ui64(v964-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v1033 = m.ExcPending
																			if v1033 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(50331778))
																				mBase = m.M
																				v1036 = m.ExcPending
																				if v1036 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg(m, int32(403555), int32(0))
																					mBase = m.M
																					v1040 = m.ExcPending
																					if v1040 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(501522), int32(2040), int32(492002))
																						mBase = m.M
																						v1045 = m.ExcPending
																						if v1045 != 0 {
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
																			F_pfree(m, v68)
																			mBase = m.M
																			v970 = m.ExcPending
																			if v970 != 0 {
																				return int32(0)
																			} else {
																				v971 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
																				if v971 != 0 {
																					F_pfree(m, v971)
																					mBase = m.M
																					v973 = m.ExcPending
																					if v973 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v14 - int32(-64)
																						return base.I32_wrap_i64(v964)
																					}
																				} else {
																					m.G0 = v14 - int32(-64)
																					return base.I32_wrap_i64(v964)
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
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v982 = m.ExcPending
					if v982 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(386138242))
						mBase = m.M
						v985 = m.ExcPending
						if v985 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(240472), int32(0))
							mBase = m.M
							v989 = m.ExcPending
							if v989 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(501522), int32(1980), int32(492002))
								mBase = m.M
								v994 = m.ExcPending
								if v994 != 0 {
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
