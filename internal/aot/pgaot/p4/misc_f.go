package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FindFKPeriodOpers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v5
	v19 = F_get_opclass_opfamily_and_input_type(m, l0, v9+int32(28), v9+int32(24))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		if v19 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			if base.B2i32(v21 == int32(3831))|base.B2i32(v21 == int32(_a_F_FindFKPeriodOpers_0)) != 0 {
				v65 = v9 + int32(22)
				F_GetOperatorFromCompareType(m, l0, int32(0), int32(8), l1, v65)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_GetOperatorFromCompareType(m, l0, int32(_a_F_FindFKPeriodOpers_0), int32(8), l2, v65)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
						if v72 != int32(3831) {
							if v72 != int32(_a_F_FindFKPeriodOpers_0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v88
									F_errmsg_internal(m, int32(_a_F_FindFKPeriodOpers_1), v9)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1725), int32(_a_F_FindFKPeriodOpers_3))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v79 = int32(_a_F_FindFKPeriodOpers_4)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v79
								m.G0 = v9 + int32(32)
								return
							}
						} else {
							v79 = int32(3900)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v79
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_FindFKPeriodOpers_5), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v40 = F_errdetail(m, int32(_a_F_FindFKPeriodOpers_6), int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1687), int32(_a_F_FindFKPeriodOpers_3))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
				F_errmsg_internal(m, int32(_a_F_FindFKPeriodOpers_7), v9+int32(16))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1691), int32(_a_F_FindFKPeriodOpers_3))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
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
func F_FinishPreparedTransaction(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
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
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int64
	_ = v399
	var v400 int64
	_ = v400
	var v408 int64
	_ = v408
	var v410 int32
	_ = v410
	var v413 int64
	_ = v413
	var v414 int32
	_ = v414
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int32
	_ = v426
	var v428 int64
	_ = v428
	var v434 int64
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int64
	_ = v480
	var v481 int64
	_ = v481
	var v491 int32
	_ = v491
	var v494 int64
	_ = v494
	var v495 int32
	_ = v495
	var v503 int64
	_ = v503
	var v505 int64
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[0]))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[1])))
	if v34 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_before_shmem_exit(m, int32(413), int64(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v49 = F_LWLockAcquire(m, v45+int32(2304), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[1])) = uint8(v42)
	goto L3
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if int32(0) < v53 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L178
	}
L8:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L4
	} else {
		goto L175
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L170
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L166
	}
L11:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v214
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[5])) = v88
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v219+int32(2304))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L45
	}
L12:
	;
	v61 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v190+int32(2304))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L40
	}
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(8)+v61<<(uint(int32(2))%32))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+48)))
	if v89 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v161 = v61 + int32(1)
	if v161 != v53 {
		v61 = v161
		goto L15
	} else {
		goto L39
	}
L18:
	;
	v93 = v88 + int32(51)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v96 == int32(0))|base.B2i32(v96 != v99) != 0 {
		v117 = v96
		v118 = v99
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v117-v118 != 0 {
		goto L17
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v102 = v93
	v103 = l0
	goto L22
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v107
		v118 = v106
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v117 = v107
	v118 = v106
	goto L20
L24:
	;
	v110 = int32(1)
	if v107 == v106 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+44))
	if v120 != int32(-1) {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[6]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v88)+40))
	if v127 != v32 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v129 = F_superuser_arg(m, v32)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[7]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126+v123*int32(768))+20))
	if v134 == v138 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	if v129 == int32(0) {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_0), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errhint(m, int32(_a_F_FinishPreparedTransaction_1), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(607), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	goto L16
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = l0
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_4), v27+int32(-16))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(623), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[6]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v88)+32))
	v229 = base.I32_wrap_i64(v228)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+49)))
	if v230 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v244)+40))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+36))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)+28))
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244)+54)))
	v258 = v244 + (v251+int32(7))&int32(_a_F_FinishPreparedTransaction_5) + int32(72)
	v262 = v250 - int32(1)
	if v262 < int32(0) {
		v330 = v229
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v234 = F_ReadTwoPhaseFile(m, v228, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v88)+16))
	F_XlogReadTwoPhaseData(m, v237, v27+int32(-4), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v234
	v244 = v234
	goto L46
L51:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v244 = v243
	goto L46
L52:
	;
	v335 = int32(_a_F_FinishPreparedTransaction_6)
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8])) = v337 + int32(1)
	v343 = int32(7)
	v345 = int32(-8)
	v347 = v258 + (v250<<(uint(int32(2))%32)+v343)&v345
	v348 = int32(12)
	v354 = v347 + (v249*v348+v343)&v345
	v361 = v354 + (v248*v348+v343)&v345
	v362 = int32(4)
	v364 = v361 + v247<<(uint(v362)%32)
	v367 = v364 + v246<<(uint(v362)%32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v244)+28))
	if l1 != 0 {
		goto L84
	} else {
		goto L85
	}
L53:
	;
	goto L52
L54:
	;
	if v250&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v267 = int32(2)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v258+v262<<(uint(v267)%32))))
	if base.B2i32(base.Ui32(v267) < base.Ui32(v270))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v229)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v285 = v229
	v287 = v262
	goto L57
L57:
	;
	if v262 == int32(0) {
		v330 = v285
		goto L53
	} else {
		goto L65
	}
L58:
	;
	v285 = v282
	v287 = v250 - int32(2)
	goto L57
L59:
	;
	v282 = v270
	goto L58
L60:
	;
	if base.Ui32(v229) < base.Ui32(v270) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if int32(0) <= v229-v270 {
		v282 = v229
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v282 = v229
	goto L58
L64:
	;
	goto L59
L65:
	;
	v290 = v285
	v291 = v287
	goto L66
L66:
	;
	v295 = int32(3)
	v299 = v258 + v291<<(uint(int32(2))%32)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if base.B2i32(base.Ui32(v290) < base.Ui32(v295))|base.B2i32(base.Ui32(v300) < base.Ui32(v295)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v330 = v325
	goto L53
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v299-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v313))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v310)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v310 = v300
	goto L68
L70:
	;
	if v290-v300 < int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(v300) <= base.Ui32(v290) {
		v310 = v290
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v310 = v290
	goto L68
L74:
	;
	goto L69
L75:
	;
	if int32(1) < v291 {
		v290 = v325
		v291 = v291 - int32(2)
		goto L66
	} else {
		goto L82
	}
L76:
	;
	v325 = v313
	goto L75
L77:
	;
	if base.Ui32(v310) < base.Ui32(v313) {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int32(0) <= v310-v313 {
		v325 = v310
		goto L75
	} else {
		goto L81
	}
L80:
	;
	v325 = v310
	goto L75
L81:
	;
	goto L76
L82:
	;
	goto L67
L83:
	;
	v536 = v367 + v245<<(uint(int32(4))%32)
	F_ProcArrayRemove(m, v227+v224*int32(768), v330)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L110
	}
L84:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+52)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v244)+40))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v373 = int32(_a_F_FinishPreparedTransaction_7)
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])) = v375 + int32(1)
	v380 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])))
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[11]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+336)) = v383 | int32(5)
	v387 = int32(0)
	v390 = base.AtomicRmwOr32(m, v387, int32(_a_F_FinishPreparedTransaction_8), v387)
	v394 = m.G0
	v395 = int32(16)
	v396 = v394 - v395
	m.G0 = v396
	F_gettimeofday(m, v396)
	mBase = m.M
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v396)))
	v400 = int64(*(*int32)(unsafe.Add(mBase, uint32(v396)+8)))
	m.G0 = v396 + v395
	v408 = v400 + v399*int64(1000000) - int64(946684800000000)
	goto L87
L85:
	;
	goto L86
L86:
	;
	v461 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v244)+36))
	v464 = F_TransactionIdDidCommit(m, v229)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L99
	}
L87:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[12]))
	v413 = F_XactLogCommitRecord(m, v408, v368, v258, v372, v347, v371, v361, v370, v367, v369, v410|int32(2), v229, l0)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v380+int32(1))&int32(_a_F_FinishPreparedTransaction_9)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])))
	F_TransactionTreeSetCommitTsData(m, v229, v368, v258, v434, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L95
	}
L90:
	;
	v422 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[13]))
	v424 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[14]))
	F_replorigin_session_advance(m, v422, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[15])) = v408
	v434 = v408
	goto L89
L93:
	;
	v428 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[15]))
	if v428 != int64(0) {
		v434 = v428
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	F_XLogFlush(m, v413)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_TransactionIdCommitTree(m, v229, v368, v258)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[11]))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v446)+336)) = v447 & int32(-6)
	v451 = int32(_a_F_FinishPreparedTransaction_7)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9]))
	v454 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])) = v453 - v454
	F_SyncRepWaitForLSN(m, v413, v454)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v523 = v244 + int32(32)
	v525 = v347
	goto L83
L99:
	;
	if v464 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v466 = int32(_a_F_FinishPreparedTransaction_7)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])) = v468 + int32(1)
	v475 = m.G0
	v476 = int32(16)
	v477 = v475 - v476
	m.G0 = v477
	F_gettimeofday(m, v477)
	mBase = m.M
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v477)))
	v481 = int64(*(*int32)(unsafe.Add(mBase, uint32(v477)+8)))
	m.G0 = v477 + v476
	goto L101
L101:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[12]))
	v494 = F_XactLogAbortRecord(m, v481+v480*int64(1000000)-int64(946684800000000), v368, v258, v463, v354, v462, v364, v491|int32(2), v229, l0)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32((v461-int32(1))&int32(_a_F_FinishPreparedTransaction_9)) <= base.Ui32(int32(_a_F_FinishPreparedTransaction_10)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v503 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[13]))
	v505 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[14]))
	F_replorigin_session_advance(m, v503, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_XLogFlush(m, v494)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	F_TransactionIdAbortTree(m, v229, v368, v258)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v514 = int32(_a_F_FinishPreparedTransaction_7)
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])) = v516 - int32(1)
	F_SyncRepWaitForLSN(m, v494, int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v523 = v244 + int32(36)
	v525 = v354
	goto L83
L110:
	;
	v542 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+48)) = uint8(v542)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	F_DropRelationFiles(m, v525, v544, v542)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	if l1 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v697 = m.G0
	v699 = v697 - int32(16)
	m.G0 = v699
	*(*uint32)(unsafe.Add(mBase, uint32(v699)+12)) = uint32(v228)
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v707 = F_LWLockAcquire(m, v703+int32(3584), int32(1))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L145
	}
L113:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v244)+40))
	F_pgstat_execute_transactional_drops(m, v548, v361)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
	F_pgstat_execute_transactional_drops(m, v616, v364)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L135
	}
L116:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+52)))
	if v551 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	F_SIInsertDataEntries(m, v367, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+52)))
	if v559 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v569 = F_LWLockAcquire(m, v565+int32(2304), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L4
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536)+4)))
	if v571 == int32(0) {
		goto L112
	} else {
		goto L127
	}
L127:
	;
	v574 = v571
	v576 = v536
	goto L128
L128:
	;
	v601 = v576 + int32(8)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v574<<(uint(int32(2))%32))+uint32(_c_F_FinishPreparedTransaction[16])))
	if v604 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L112
L130:
	;
	v605 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v576)+6)))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	m.T0[v604].(func(*base.Module, int64, int32, int32, int32))(m, v228, v605, v601, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v614 = v601 + (v609+int32(7))&int32(-8)
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+4)))
	if v615 != 0 {
		v574 = v615
		v576 = v614
		goto L128
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	goto L129
L135:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v624 = F_LWLockAcquire(m, v620+int32(2304), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536)+4)))
	if v626 == int32(0) {
		goto L112
	} else {
		goto L137
	}
L137:
	;
	v629 = v626
	v631 = v536
	goto L138
L138:
	;
	v656 = v631 + int32(8)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v629<<(uint(int32(2))%32))+uint32(_c_F_FinishPreparedTransaction[17])))
	if v659 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L112
L140:
	;
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v631)+6)))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	m.T0[v659].(func(*base.Module, int64, int32, int32, int32))(m, v228, v660, v656, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	v669 = v656 + (v664+int32(7))&int32(-8)
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+4)))
	if v670 != 0 {
		v629 = v670
		v631 = v669
		goto L138
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	goto L139
L145:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[18]))
	v713 = int32(0)
	v715 = F_hash_search(m, v710, v699+int32(12), v713, v713)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v718+int32(3584))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v715 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	v725 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[19])) = uint8(v725)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[20])) = v723
	F_ReleasePredicateLocks(m, l1, int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	m.G0 = v699 + int32(16)
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+4))
	if v738 <= int32(0) {
		goto L7
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+49)))
	v743 = v737 + int32(8)
	v747 = int32(0)
	goto L153
L153:
	;
	v773 = v743 + v747<<(uint(int32(2))%32)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	if v774 != v88 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v780 = v738 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v780
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v743+v780<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v773))) = v785
	v787 = int32(_a_F_FinishPreparedTransaction_11)
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v789
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v792))) = v88
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v795+int32(2304))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L159
	}
L155:
	;
	v777 = v747 + int32(1)
	if v738 != v777 {
		v747 = v777
		goto L153
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	goto L154
L158:
	;
	goto L7
L159:
	;
	F_AtEOXact_PgStat(m, l1, int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	if v741&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	F_RemoveTwoPhaseFile(m, v228, int32(1))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v808 = int32(_a_F_FinishPreparedTransaction_6)
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8])) = v810 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[5])) = int32(0)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	F_pfree(m, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	m.G0 = v29 - int32(-64)
	return
L166:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l0
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_12), v27+int32(-32))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(589), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_13), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errhint(m, int32(_a_F_FinishPreparedTransaction_14), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(595), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v229
	F_errmsg_internal(m, int32(_a_F_FinishPreparedTransaction_15), v27+int32(-48))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(2463), int32(_a_F_FinishPreparedTransaction_16))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v88
	F_errmsg_internal(m, int32(_a_F_FinishPreparedTransaction_17), v29)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(658), int32(_a_F_FinishPreparedTransaction_18))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FlushOneBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_FlushOneBuffer[0]))
	v4 = int32(56)
	F_FlushBuffer(m, v3+l0*v4-v4, int32(0), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_ForgetInplace_Inval(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_ForgetInplace_Inval[0])) = int32(0)
	return
}
func F_FreeSpaceMapVacuumRange(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(l1) < base.Ui32(l2) {
		v11 = *(*int64)(unsafe.Add(mBase, _c_F_FreeSpaceMapVacuumRange[0]))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
		v15 = F_fsm_vacuum_page(m, l0, v7, l1, l2, v7+int32(15))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F___fstat(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	if l0 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F___fstat[0])) = int32(8)
		return int32(-1)
	} else {
		v17 = F___fstatat(m, l0, int32(_a_F___fstat_0), l1, int32(_a_F___fstat_1))
		mBase = m.M
		return v17
	}
}
func F_fdw_handler_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_fdw_handler_in_0), int32(369), int32(_a_F_fdw_handler_in_1), int32(_a_F_fdw_handler_in_2), int32(_a_F_fdw_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_findNotNullConstraint(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = F_get_attnum(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 <= int32(0) {
			return int32(0)
		} else {
			v11 = F_findNotNullConstraintAttnum(m, l0, v3)
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_findTargetlistEntrySQL99(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	v9 = F_transformExpr(m, l0, l1, l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v119
L4:
	;
	v86 = m.G0
	v88 = v86 - int32(16)
	m.G0 = v88
	if v9 != 0 {
		v98 = v9
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(0)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v26<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L4
L9:
	;
	v72 = F_equal(m, v9, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L30
	}
L10:
	;
	goto L9
L11:
	;
	v33 = v32
	goto L14
L12:
	;
	goto L13
L13:
	;
	v71 = int32(0)
	goto L10
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	switch v34 - int32(15) {
	case 0:
		goto L22
	default:
		v71 = v33
		goto L10
	case 12:
		goto L21
	case 13:
		goto L20
	case 14:
		goto L19
	case 15:
		goto L18
	case 40:
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v68 != 0 {
		v33 = v68
		goto L14
	} else {
		goto L29
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	if v62 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L28
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v57 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L27
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v52 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L26
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v47 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L25
	}
L21:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	if v42 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L24
	}
L22:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v37 != int32(2) {
		v71 = v33
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v67 = v41
	goto L16
L24:
	;
	v67 = v33 + int32(4)
	goto L16
L25:
	;
	v67 = v33 + int32(4)
	goto L16
L26:
	;
	v67 = v33 + int32(4)
	goto L16
L27:
	;
	v67 = v33 + int32(4)
	goto L16
L28:
	;
	v67 = v33 + int32(4)
	goto L16
L29:
	;
	goto L15
L30:
	;
	if v72 != 0 {
		v119 = v31
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v75 = v26 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v75 < v76 {
		v26 = v75
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v100 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v99 + v100
	v106 = F_makeTargetEntry(m, v98, base.I32_extend16_s(v99), int32(0), v100)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L40
	}
L34:
	;
	if l3 == int32(16) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v92 == int32(57) {
		v98 = l1
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v96 = F_transformExpr(m, l0, l1, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v98 = v96
	goto L33
L40:
	;
	m.G0 = v88 + int32(16)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v112 = F_lappend(m, v111, v106)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v112
	v119 = v106
	goto L3
}
func F_find_appinfos_by_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = int64(0)
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v60 = F_palloc_mul(m, int32(4), v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v59 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = l1 + int32(8)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v59 = base.I32_popcnt(v24)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v27 = v21 << (uint(int32(2)) % 32)
	if v27 <= int32(7) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v59 = base.I32_wrap_i64(v54)
	goto L1
L9:
	;
	if v27 == int32(0) {
		v54 = v15
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = F_pg_popcount_optimized(m, v20, v27)
	mBase = m.M
	v54 = v51
	goto L8
L12:
	;
	v32 = v27
	v33 = v20
	v34 = v15
	goto L13
L13:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+3)))
	v36 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_find_appinfos_by_relids[0]))))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+2)))
	v38 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_find_appinfos_by_relids[0]))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v40 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_find_appinfos_by_relids[0]))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v42 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_find_appinfos_by_relids[0]))))
	v46 = v36 + (v38 + (v40 + (v34 + v42)))
	v47 = int32(4)
	v50 = v32 - v47
	if v50 != 0 {
		v32 = v50
		v33 = v33 + v47
		v34 = v46
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v54 = v46
	goto L8
L15:
	;
	goto L14
L16:
	;
	return int32(0)
L17:
	;
	if l1 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if int32(0) <= v120 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v120 = base.I32_ctz(v106) | v107<<(uint(int32(5))%32)
	goto L18
L20:
	;
	v120 = int32(-2)
	goto L18
L21:
	;
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 <= v71 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v77 = l1 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v84 = v81 & int32(-1)
	if v84 != 0 {
		v106 = v84
		v107 = v71
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(1)
	if v85 == v74 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v89 = v85
	goto L25
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77+v89<<(uint(int32(2))%32))))
	if v96 != 0 {
		v106 = v96
		v107 = v89
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L20
L27:
	;
	v98 = v89 + int32(1)
	if v98 != v74 {
		v89 = v98
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v126 = v120
	v128 = v4
	goto L32
L30:
	;
	v225 = v4
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
	m.G0 = v11 + int32(16)
	return v60
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v126<<(uint(int32(2))%32))))
	if v135 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v225 = v161
	goto L31
L34:
	;
	if l1 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v138 = F_find_base_rel_ignore_join(m, l0, v126)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L16
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60+v128<<(uint(int32(2))%32)))) = v135
	v161 = v128 + int32(1)
	goto L34
L38:
	;
	if v138 == int32(0) {
		v161 = v128
		goto L34
	} else {
		goto L39
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v126
	F_errmsg_internal(m, int32(_a_F_find_appinfos_by_relids_0), v11)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L16
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_find_appinfos_by_relids_1), int32(829), int32(_a_F_find_appinfos_by_relids_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	if int32(0) <= v217 {
		v126 = v217
		v128 = v161
		goto L32
	} else {
		goto L54
	}
L44:
	;
	v217 = base.I32_ctz(v203) | v204<<(uint(int32(5))%32)
	goto L43
L45:
	;
	v217 = int32(-2)
	goto L43
L46:
	;
	v168 = v126 + int32(1)
	v170 = int32(base.Ui32(v168) >> (uint(int32(5)) % 32))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v171 <= v170 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v174 = l1 + int32(8)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v170<<(uint(int32(2))%32))))
	v181 = v178 & (int32(-1) << (uint(v168) % 32))
	if v181 != 0 {
		v203 = v181
		v204 = v170
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v183 = v170 + int32(1)
	if v183 == v171 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v186 = v183
	goto L50
L50:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174+v186<<(uint(int32(2))%32))))
	if v193 != 0 {
		v203 = v193
		v204 = v186
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L45
L52:
	;
	v195 = v186 + int32(1)
	if v195 != v171 {
		v186 = v195
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L33
}
func F_find_cols_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
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
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v3 - int32(6) {
		case 0:
			v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v7 == int32(1) {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v11 = F_bms_add_member(m, v10, v6)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v11
					return int32(0)
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v19 = F_bms_add_member(m, v18, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v19
					return int32(0)
				}
			}
		default:
			v34 = F_expression_tree_walker_impl(m, l0, int32(739), l1)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v37 = v34
				return v37
			}
		case 3:
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v24)
			v27 = F_expression_tree_walker_impl(m, l0, int32(739), l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v29)
				return v29
			}
		}
	} else {
		v37 = int32(0)
		return v37
	}
}
func F_find_duplicate_ors(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int64
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int64
	_ = v261
	var v268 int64
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v410
L2:
	;
	v410 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 != int32(21) {
		v410 = l0
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v20 {
	case 0:
		goto L11
	case 1:
		goto L12
	default:
		v410 = l0
		goto L1
	}
L6:
	;
	v406 = F_pull_ands(m, v399)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L20
	} else {
		goto L144
	}
L7:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v392 != int32(1) {
		v399 = v385
		goto L6
	} else {
		goto L143
	}
L8:
	;
	v381 = F_make_orclause(m, v84)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L20
	} else {
		goto L142
	}
L9:
	;
	if v220 == int32(0) {
		goto L8
	} else {
		goto L115
	}
L10:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v410 = v308
	goto L1
L11:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 == int32(0) {
		v288 = v3
		goto L84
	} else {
		goto L85
	}
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 == int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v84 = F_pull_ors(m, v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L36
	}
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= int32(0) {
		v80 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v29 = v3
	v32 = v3
	goto L16
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v41 = F_find_duplicate_ors(m, v40, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v80 = v70
	goto L13
L18:
	;
	v72 = v29 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v72 < v73 {
		v29 = v72
		v32 = v70
		goto L16
	} else {
		goto L35
	}
L19:
	;
	v67 = F_lappend(m, v32, v41)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L34
	}
L20:
	;
	return int32(0)
L21:
	;
	if v41 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v47 != int32(7) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+32)))
	if l1 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v50&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v50&int32(1) != 0 {
		v70 = v32
		goto L18
	} else {
		goto L32
	}
L27:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
	if v55 == int64(0) {
		v70 = v32
		goto L18
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v60 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v410 = v60
	goto L1
L32:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
	if v64 == int64(0) {
		v70 = v32
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v410 = v41
	goto L1
L34:
	;
	v70 = v67
	goto L18
L35:
	;
	goto L17
L36:
	;
	if v84 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v88 = int32(0)
	v90 = F_makeBoolConst(m, v88, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L20
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v92 == int32(1) {
		goto L10
	} else {
		goto L41
	}
L40:
	;
	v410 = v90
	goto L1
L41:
	;
	v95 = int32(0)
	if v92 <= v95 {
		v141 = v95
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v149 = F_list_union(m, v141)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L20
	} else {
		goto L61
	}
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v99 = int32(0)
	v102 = v95
	v104 = v99
	v106 = v99
	goto L44
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v98+v104<<(uint(int32(2))%32))))
	if v113 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v141 = v135
	goto L42
L46:
	;
	v134 = base.B2i32(v102 == int32(0)) | base.B2i32(v130 < v106)
	if v134 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v130 = v129
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v113
	v127 = F_list_make1_impl(m, int32(1), v12+int32(8))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L20
	} else {
		goto L53
	}
L49:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v116 != int32(21) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v119 != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	if v120 != 0 {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v130 = int32(0)
	goto L46
L53:
	;
	v141 = v127
	goto L42
L54:
	;
	v135 = v120
	goto L56
L55:
	;
	v135 = v102
	goto L56
L56:
	;
	if v134 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v136 = v130
	goto L59
L58:
	;
	v136 = v106
	goto L59
L59:
	;
	v138 = v104 + int32(1)
	if v138 != v92 {
		v102 = v135
		v104 = v138
		v106 = v136
		goto L44
	} else {
		goto L60
	}
L60:
	;
	goto L45
L61:
	;
	if v149 == int32(0) {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v153 <= int32(0) {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v156 = int32(0)
	v160 = v156
	v163 = v156
	goto L64
L64:
	;
	v167 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v163<<(uint(int32(2))%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v167 < v173 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L9
L66:
	;
	v228 = v163 + int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v228 < v229 {
		v160 = v220
		v163 = v228
		goto L64
	} else {
		goto L83
	}
L67:
	;
	v176 = v167
	goto L70
L68:
	;
	goto L69
L69:
	;
	v216 = F_lappend(m, v160, v172)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L20
	} else {
		goto L82
	}
L70:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v176<<(uint(int32(2))%32))))
	if v189 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L69
L72:
	;
	v204 = v176 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v204 < v205 {
		v176 = v204
		goto L70
	} else {
		goto L81
	}
L73:
	;
	v199 = F_equal(m, v172, v189)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L20
	} else {
		goto L79
	}
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v192 != int32(21) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if v195 != 0 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v197 = F_list_member(m, v196, v172)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L77
	}
L77:
	;
	if v197 != 0 {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	v220 = v160
	goto L66
L79:
	;
	if v199 == int32(0) {
		v220 = v160
		goto L66
	} else {
		goto L80
	}
L80:
	;
	goto L72
L81:
	;
	goto L71
L82:
	;
	v220 = v216
	goto L66
L83:
	;
	goto L65
L84:
	;
	v292 = F_pull_ands(m, v288)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L20
	} else {
		goto L106
	}
L85:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v234 <= int32(0) {
		v288 = v3
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v239 = v3
	v242 = v3
	goto L87
L87:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246+v239<<(uint(int32(2))%32))))
	v251 = F_find_duplicate_ors(m, v250, l1)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L91
	}
L88:
	;
	v288 = v278
	goto L84
L89:
	;
	v280 = v239 + int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v280 < v281 {
		v239 = v280
		v242 = v278
		goto L87
	} else {
		goto L105
	}
L90:
	;
	v275 = F_lappend(m, v242, v251)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L20
	} else {
		goto L104
	}
L91:
	;
	if v251 == int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v255 != int32(7) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+32)))
	if l1 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if v258&int32(1) != 0 {
		v278 = v242
		goto L89
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v258&int32(1) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v251)+24))
	if v261 != int64(0) {
		v278 = v242
		goto L89
	} else {
		goto L98
	}
L98:
	;
	v410 = v251
	goto L1
L99:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v251)+24))
	if v268 != int64(0) {
		v278 = v242
		goto L89
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v271 = int32(0)
	v273 = F_makeBoolConst(m, v271, v271)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v410 = v273
	goto L1
L104:
	;
	v278 = v275
	goto L89
L105:
	;
	goto L88
L106:
	;
	if v292 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v298 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L20
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v300 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v410 = v298
	goto L1
L111:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v410 = v304
	goto L1
L112:
	;
	goto L113
L113:
	;
	v305 = F_make_andclause(m, v292)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	v410 = v305
	goto L1
L115:
	;
	v311 = int32(0)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v312 <= v311 {
		v385 = v220
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v317 = int32(0)
	v319 = v311
	goto L117
L117:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325+v317<<(uint(int32(2))%32))))
	if v329 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	if v351 == int32(0) {
		v385 = v220
		goto L7
	} else {
		goto L133
	}
L119:
	;
	v351 = F_lappend(m, v319, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L20
	} else {
		goto L131
	}
L120:
	;
	v348 = F_make_andclause(m, v337)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L20
	} else {
		goto L130
	}
L121:
	;
	v346 = F_list_member(m, v220, v329)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L20
	} else {
		goto L128
	}
L122:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if v332 != int32(21) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v335 != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v337 = F_list_difference(m, v336, v220)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	if v337 == int32(0) {
		v385 = v220
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v341 != int32(1) {
		goto L120
	} else {
		goto L127
	}
L127:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v350 = v345
	goto L119
L128:
	;
	if v346 != 0 {
		v385 = v220
		goto L7
	} else {
		goto L129
	}
L129:
	;
	v350 = v329
	goto L119
L130:
	;
	v350 = v348
	goto L119
L131:
	;
	v354 = v317 + int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v354 < v355 {
		v317 = v354
		v319 = v351
		goto L117
	} else {
		goto L132
	}
L132:
	;
	goto L118
L133:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v359 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v369 = F_lappend(m, v220, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L20
	} else {
		goto L140
	}
L135:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v368 = v363
	goto L134
L136:
	;
	goto L137
L137:
	;
	v364 = F_pull_ors(m, v351)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	v366 = F_make_orclause(m, v364)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L20
	} else {
		goto L139
	}
L139:
	;
	v368 = v366
	goto L134
L140:
	;
	if v369 != 0 {
		v385 = v369
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v399 = int32(0)
	goto L6
L142:
	;
	v410 = v381
	goto L1
L143:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v410 = v396
	goto L1
L144:
	;
	v408 = F_make_andclause(m, v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L20
	} else {
		goto L145
	}
L145:
	;
	v410 = v408
	goto L1
}
func F_find_forced_null_vars(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v71 int32
	_ = v71
	v2 = int32(0)
	if l0 == v2 {
		v71 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v71
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 != int32(21) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v71 = v2
	goto L1
L5:
	;
	switch v9 - int32(52) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L11
	}
L6:
	;
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v67 != 0 {
		v71 = v2
		goto L1
	} else {
		goto L30
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+8)))
	v64 = F_mbms_add_member(m, v60, v61+int32(7))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L16
	} else {
		goto L29
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v49 != int32(4) {
		v71 = v2
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v38 != 0 {
		v71 = v2
		goto L1
	} else {
		goto L20
	}
L11:
	;
	if v9 != int32(1) {
		v71 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v16 <= int32(0) {
		v71 = v2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v21 = int32(0)
	v22 = v2
	goto L14
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v21<<(uint(int32(2))%32))))
	v28 = F_find_forced_null_vars(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v71 = v32
	goto L1
L16:
	;
	return int32(0)
L17:
	;
	v32 = F_mbms_add_members(m, v22, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v35 = v21 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v35 < v36 {
		v21 = v35
		v22 = v32
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
	if v39 != 0 {
		v71 = v2
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v40 == int32(0) {
		v71 = v2
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v43 != int32(6) {
		v71 = v2
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if v46 == int32(0) {
		v59 = v40
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v71 = v2
	goto L1
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v52 == int32(0) {
		v71 = v2
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v55 != int32(6) {
		v71 = v2
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	if v58 != 0 {
		v71 = v2
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v59 = v52
	goto L8
L29:
	;
	return v64
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v68 != 0 {
		v6 = v68
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L4
}
func F_find_indexpath_quals(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v11 - int32(283) {
	case 0:
		goto L3
	default:
		goto L1
	case 3:
		goto L5
	case 4:
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L27
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v60 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v37 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v23 = v4
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v23<<(uint(int32(2))%32))))
	F_find_indexpath_quals(m, v30, l1, l2)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	v34 = v23 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v34 < v35 {
		v23 = v34
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v46 = v4
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v46<<(uint(int32(2))%32))))
	F_find_indexpath_quals(m, v53, l1, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L2
L17:
	;
	v57 = v46 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v57 < v58 {
		v46 = v57
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+88))
	v96 = F_list_concat(m, v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L26
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v69 = v4
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v69<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = F_lappend(m, v72, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
	v84 = v69 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v84 < v85 {
		v69 = v84
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96
	goto L2
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v112
	F_errmsg_internal(m, int32(_a_F_find_indexpath_quals_0), v9)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_find_indexpath_quals_1), int32(2189), int32(_a_F_find_indexpath_quals_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_inheritance_children(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_find_inheritance_children_extended(m, l0, int32(1), l1, v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_find_inheritance_children_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int64
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v21 = base.I64_extend_i32_u(l0)
	v22 = F_SearchSysCache1(m, int32(57), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v27)+126)))
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L76
	}
L6:
	;
	if v29 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = F_palloc(m, int32(128))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v258 = v6
	goto L9
L9:
	;
	m.G0 = v18 + int32(96)
	return v258
L10:
	;
	v39 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = v18 + int32(32)
	F_ScanKeyInit(m, v42, int32(2), int32(3), int32(184), v21)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(1)
	v52 = F_systable_beginscan(m, v39, int32(2187), v49, int32(0), v49, v42)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v54 = F_systable_getnext(m, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = v54
	v64 = v35
	v65 = int32(32)
	v66 = v6
	goto L18
L16:
	;
	v172 = v35
	v174 = v6
	goto L17
L17:
	;
	F_systable_endscan(m, v52)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L55
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v73)+12)))
	if v75 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v172 = v159
	v174 = v161
	goto L17
L20:
	;
	v163 = F_systable_getnext(m, v52)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L53
	}
L21:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+22)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138+v139)))
	if v65 <= v66 {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v78)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if l1 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_find_inheritance_children_extended[0]))
	goto L27
L27:
	;
	if base.B2i32(v83 != int32(0)) == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+20)))
	v91 = int32(768)
	if v90&v91 != v91 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v96 = v95
	goto L31
L30:
	;
	v96 = int32(2)
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_find_inheritance_children_extended[0]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	goto L32
L32:
	;
	v100 = F_XidInMVCCSnapshot(m, v96, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v100 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	if l4 == int32(0) {
		v159 = v64
		v160 = v65
		v161 = v66
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v104 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v96
	v159 = v64
	v160 = v65
	v161 = v66
	goto L20
L37:
	;
	v109 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v109 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_find_inheritance_children_extended_0), v18+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v122 = int32(3)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.B2i32(base.Ui32(v96) < base.Ui32(v122))|base.B2i32(base.Ui32(v124) < base.Ui32(v122)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_errfinish(m, int32(_a_F_find_inheritance_children_extended_1), int32(168), int32(_a_F_find_inheritance_children_extended_2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if int32(0) < v96-v124 {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v96) <= base.Ui32(v124) {
		v159 = v64
		v160 = v65
		v161 = v66
		goto L20
	} else {
		goto L48
	}
L47:
	;
	v159 = v64
	v160 = v65
	v161 = v66
	goto L20
L48:
	;
	goto L36
L49:
	;
	v145 = F_repalloc(m, v64, v65<<(uint(int32(3))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v149 = v64
	v150 = v65
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149+v66<<(uint(int32(2))%32)))) = v141
	v159 = v149
	v160 = v150
	v161 = v66 + int32(1)
	goto L20
L52:
	;
	v149 = v145
	v150 = v65 << (uint(int32(1)) % 32)
	goto L51
L53:
	;
	if v163 != 0 {
		v62 = v163
		v64 = v159
		v65 = v160
		v66 = v161
		goto L18
	} else {
		goto L54
	}
L54:
	;
	goto L19
L55:
	;
	F_relation_close(m, v39, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if int32(2) <= v174 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_pfree(m, v172)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L75
	}
L58:
	;
	v195 = int32(0)
	v202 = v195
	v203 = v195
	goto L64
L59:
	;
	F_pg_qsort(m, v172, v174, int32(4), int32(506))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v174 != int32(1) {
		v241 = int32(0)
		goto L57
	} else {
		goto L63
	}
L62:
	;
	goto L58
L63:
	;
	goto L58
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v172+v202<<(uint(int32(2))%32))))
	if l2 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v241 = v231
	goto L57
L66:
	;
	v233 = v202 + int32(1)
	if v233 != v174 {
		v202 = v233
		v203 = v231
		goto L64
	} else {
		goto L74
	}
L67:
	;
	v229 = F_lappend_oid(m, v203, v215)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L73
	}
L68:
	;
	F_LockRelationOid(m, v215, l2)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v222 = int64(0)
	v225 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(v215), v222, v222, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v225 != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	F_UnlockRelationOid(m, v215, l2)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v231 = v203
	goto L66
L73:
	;
	v231 = v229
	goto L66
L74:
	;
	goto L65
L75:
	;
	v258 = v241
	goto L9
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_find_inheritance_children_extended_3), v18)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_find_inheritance_children_extended_1), int32(363), int32(_a_F_find_inheritance_children_extended_4))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_matching_subplans_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if v30 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v30 = v28
	goto L3
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v20 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v18 = F_get_matching_partitions(m, l1+int32(32), v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v30 = v18
	goto L3
L10:
	;
	v25 = F_get_matching_partitions(m, l1+int32(76), v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v30 = v25
	goto L3
L12:
	;
	if int32(0) <= v87 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v87 = base.I32_ctz(v73) | v74<<(uint(int32(5))%32)
	goto L12
L14:
	;
	v87 = int32(-2)
	goto L12
L15:
	;
	v38 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v41 <= v38 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v44 = v30 + int32(8)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v51 = v48 & int32(-1)
	if v51 != 0 {
		v73 = v51
		v74 = v38
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(1)
	if v52 == v41 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v56 = v52
	goto L19
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44+v56<<(uint(int32(2))%32))))
	if v63 != 0 {
		v73 = v63
		v74 = v56
		goto L13
	} else {
		goto L21
	}
L20:
	;
	goto L14
L21:
	;
	v65 = v56 + int32(1)
	if v65 != v41 {
		v56 = v65
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v97 = v87
	goto L26
L24:
	;
	goto L25
L25:
	;
	return
L26:
	;
	v103 = v97 << (uint(int32(2)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103+v104)))
	if int32(0) <= v106 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L25
L28:
	;
	if v30 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v110 = F_bms_add_member(m, v109, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124+v103)))
	if v126 < int32(0) {
		goto L28
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v110
	if l4 == int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115+v103)))
	if v117 == int32(0) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v121 = F_bms_add_member(m, v120, v117)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v121
	goto L28
L36:
	;
	F_find_matching_subplans_recurse(m, l0, l0+int32(4)+v126*int32(120), l2, l3, l4)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L28
L38:
	;
	if int32(0) <= v190 {
		v97 = v190
		goto L26
	} else {
		goto L49
	}
L39:
	;
	v190 = base.I32_ctz(v176) | v177<<(uint(int32(5))%32)
	goto L38
L40:
	;
	v190 = int32(-2)
	goto L38
L41:
	;
	v141 = v97 + int32(1)
	v143 = int32(base.Ui32(v141) >> (uint(int32(5)) % 32))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v144 <= v143 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v147 = v30 + int32(8)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v143<<(uint(int32(2))%32))))
	v154 = v151 & (int32(-1) << (uint(v141) % 32))
	if v154 != 0 {
		v176 = v154
		v177 = v143
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v156 = v143 + int32(1)
	if v156 == v144 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v159 = v156
	goto L45
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v147+v159<<(uint(int32(2))%32))))
	if v166 != 0 {
		v176 = v166
		v177 = v159
		goto L39
	} else {
		goto L47
	}
L46:
	;
	goto L40
L47:
	;
	v168 = v159 + int32(1)
	if v168 != v144 {
		v159 = v168
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L27
}
func F_find_nonnullable_rels_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		v230 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v230
L2:
	;
	v14 = l0
	v15 = l1
	v16 = l1
	goto L3
L3:
	;
	v21 = int32(4)
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	switch v23 - int32(1) {
	case 0:
		goto L16
	case 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 15, 17, 18, 21, 23, 24, 25, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50:
		v230 = v22
		goto L1
	case 5:
		goto L15
	case 14:
		goto L14
	case 16:
		goto L13
	case 19:
		goto L12
	case 20:
		goto L11
	case 22:
		goto L8
	case 26, 27, 28, 29, 30:
		v221 = v16
		v223 = v21
		goto L5
	case 51:
		goto L10
	case 52:
		goto L9
	default:
		goto L7
	}
L4:
	;
	v230 = int32(0)
	goto L1
L5:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14+v223)))
	if v225 != 0 {
		v14 = v225
		v15 = v221
		v16 = v221
		goto L3
	} else {
		goto L84
	}
L6:
	;
	v221 = v218
	v223 = int32(28)
	goto L5
L7:
	;
	if v23 != int32(321) {
		v230 = v22
		goto L1
	} else {
		goto L63
	}
L8:
	;
	v150 = int32(8)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v151 == int32(2))&v15 != 0 {
		goto L59
	} else {
		goto L60
	}
L9:
	;
	if v15&int32(1) == int32(0) {
		v230 = v22
		goto L1
	} else {
		goto L56
	}
L10:
	;
	if v15&int32(1) == int32(0) {
		v230 = v22
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v74 = int32(8)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	switch v75 {
	case 0:
		goto L35
	case 1:
		v80 = v15
		goto L34
	case 2:
		v221 = int32(0)
		v223 = v74
		goto L5
	default:
		goto L33
	}
L12:
	;
	v71 = F_is_strict_saop(m, v14, v15&int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L31
	}
L13:
	;
	F_set_opfuncid(m, v14)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L20
	} else {
		goto L28
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v60 = F_func_strict(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L20
	} else {
		goto L26
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v54 != 0 {
		v230 = v22
		goto L1
	} else {
		goto L24
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v26 <= int32(0) {
		v230 = v22
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v33 = v22
	v34 = int32(0)
	goto L18
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v34<<(uint(int32(2))%32))))
	v44 = F_find_nonnullable_rels_walker(m, v41, v15&int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v230 = v48
	goto L1
L20:
	;
	return int32(0)
L21:
	;
	v48 = F_bms_join(m, v33, v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v51 = v34 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v51 < v52 {
		v33 = v48
		v34 = v51
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v56 = F_bms_make_singleton(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v230 = v56
	goto L1
L26:
	;
	if v60 != 0 {
		v218 = int32(0)
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v230 = v22
	goto L1
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v66 = F_func_strict(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	if v66 != 0 {
		v218 = int32(0)
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v230 = v22
	goto L1
L31:
	;
	if v71 != 0 {
		v218 = int32(0)
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v230 = v22
	goto L1
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L20
	} else {
		goto L50
	}
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v82 == int32(0) {
		v230 = v22
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v76 = int32(1)
	if v15&v76 != 0 {
		v221 = v76
		v223 = v74
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v80 = int32(0)
	goto L34
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v85 <= int32(0) {
		v230 = v22
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v91 = int32(0)
	v94 = v22
	goto L39
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v91<<(uint(int32(2))%32))))
	v103 = F_find_nonnullable_rels_walker(m, v102, v80&int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L41
	}
L40:
	;
	v230 = int32(0)
	goto L1
L41:
	;
	if v94 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v105 = F_bms_int_members(m, v94, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L45
	}
L43:
	;
	v107 = v103
	goto L44
L44:
	;
	if v107 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v107 = v105
	goto L44
L46:
	;
	v109 = v91 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v110 <= v109 {
		v230 = v107
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L40
L49:
	;
	v91 = v109
	v94 = v107
	goto L39
L50:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v117
	F_errmsg_internal(m, int32(_a_F_find_nonnullable_rels_walker_0), v10)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_find_nonnullable_rels_walker_1), int32(1686), int32(_a_F_find_nonnullable_rels_walker_2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v131 != int32(1) {
		v230 = v22
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v134 = int32(0)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v135 == v134 {
		v221 = v134
		v223 = v21
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v230 = v22
	goto L1
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(int32(5)) < base.Ui32(v143) {
		v230 = v22
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if int32(1)<<(uint(v143)%32)&int32(37) != 0 {
		v221 = int32(0)
		v223 = v21
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v230 = v22
	goto L1
L59:
	;
	v221 = v15
	v223 = v150
	goto L5
L60:
	;
	goto L61
L61:
	;
	if v151 == int32(3) {
		v221 = v15
		v223 = v150
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v230 = int32(0)
	goto L1
L63:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v163 = F_find_nonnullable_rels_walker(m, v160, v15&int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v165 != 0 {
		v230 = v163
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v167 = int32(0)
	if v166 == v167 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v212 != int32(1) {
		v230 = v163
		goto L1
	} else {
		goto L82
	}
L67:
	;
	v212 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v175 = int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v176 <= v175 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v179 = v175
	goto L72
L71:
	;
	v179 = v176
	goto L72
L72:
	;
	v183 = int32(0)
	v185 = v167
	goto L73
L73:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v166+int32(8)+v183<<(uint(int32(2))%32))))
	if v192 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v212 = v204
	goto L66
L75:
	;
	goto L74
L76:
	;
	v193 = int32(2)
	if v185 != 0 {
		v204 = v193
		goto L75
	} else {
		goto L79
	}
L77:
	;
	v199 = v185
	goto L78
L78:
	;
	v201 = v183 + int32(1)
	if v201 != v179 {
		v183 = v201
		v185 = v199
		goto L73
	} else {
		goto L81
	}
L79:
	;
	v194 = int32(1)
	if base.Ui32(v194) < base.Ui32(base.I32_popcnt(v192)) {
		v204 = v193
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v199 = v194
	goto L78
L81:
	;
	v204 = v199
	goto L75
L82:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v216 = F_bms_add_members(m, v163, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	v230 = v216
	goto L1
L84:
	;
	goto L4
}
func F_find_placeholder_info(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui32(v7) < base.Ui32(v8) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L52
	}
L2:
	;
	return v150
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v7<<(uint(int32(2))%32))))
	if v14 != 0 {
		v150 = v14
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+337)))
	if v16 == int32(1) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v20 = F_palloc0(m, int32(28))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(326)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v26
	v28 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = F_pull_varnos(m, l0, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v37 = F_bms_difference(m, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_bms_int_members(m, v34, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v41
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v47 = F_bms_copy(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v53 = F_exprType(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v47
	goto L16
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v56 = F_exprTypmod(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v58 = F_get_typavgwidth(m, v53, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v62 = F_lappend(m, v61, v20)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if base.Ui32(v65) < base.Ui32(v66) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v102<<(uint(int32(2))%32)))) = v20
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v111 = F_pull_var_clause(m, v109, int32(26))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L41
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v99 = v68
	v102 = v65
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v66 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v72 = v66 << (uint(int32(1)) % 32)
	goto L28
L27:
	;
	v72 = int32(8)
	goto L28
L28:
	;
	v74 = v72
	goto L29
L29:
	;
	if base.Ui32(v74) <= base.Ui32(v65) {
		v74 = v74 << (uint(int32(1)) % 32)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v82 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v94
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v99 = v94
	v102 = v97
	goto L22
L33:
	;
	v84 = F_mul_size(m, int32(4), v66)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v92 = F_palloc0_mul(m, int32(4), v74)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L39
	}
L36:
	;
	v87 = F_mul_size(m, int32(4), v74)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v89 = F_repalloc0(m, v82, v84, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v94 = v89
	goto L32
L39:
	;
	v94 = v92
	goto L32
L40:
	;
	F_list_free(m, v111)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L51
	}
L41:
	;
	if v111 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v115 <= int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v120 = int32(0)
	goto L44
L44:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v120<<(uint(int32(2))%32))))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v130 == int32(321) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L40
L46:
	;
	v133 = F_find_placeholder_info(m, l0, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v136 = v120 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v136 < v137 {
		v120 = v136
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L45
L51:
	;
	v150 = v20
	goto L2
L52:
	;
	F_errmsg_internal(m, int32(_a_F_find_placeholder_info_0), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_find_placeholder_info_1), int32(106), int32(_a_F_find_placeholder_info_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_placeholders_recurse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L12
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v13 - int32(63) {
	case 0:
		goto L2
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L1
	}
L4:
	;
	F_list_free(m, v114)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L37
	}
L5:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_find_placeholders_recurse(m, l0, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L25
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v47 = F_pull_var_clause(m, v45, int32(26))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L15
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v24 = v3
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	F_find_placeholders_recurse(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	return
L13:
	;
	v36 = v24 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v36 < v37 {
		v24 = v36
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v47 == int32(0) {
		v114 = v47
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v51 <= int32(0) {
		v114 = v47
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v56 = v51
	v58 = int32(0)
	goto L18
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v58<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 == int32(321) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v114 = v47
	goto L4
L20:
	;
	v69 = F_find_placeholder_info(m, l0, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L23
	}
L21:
	;
	v72 = v56
	goto L22
L22:
	;
	v74 = v58 + int32(1)
	if v74 < v72 {
		v56 = v72
		v58 = v74
		goto L18
	} else {
		goto L24
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v72 = v71
	goto L22
L24:
	;
	goto L19
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_find_placeholders_recurse(m, l0, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v84 = F_pull_var_clause(m, v82, int32(26))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v84 == int32(0) {
		v114 = v84
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v88 <= int32(0) {
		v114 = v84
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v92 = v88
	v94 = v3
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v94<<(uint(int32(2))%32))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 == int32(321) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v114 = v84
	goto L4
L32:
	;
	v105 = F_find_placeholder_info(m, l0, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L35
	}
L33:
	;
	v108 = v92
	goto L34
L34:
	;
	v110 = v94 + int32(1)
	if v110 < v108 {
		v92 = v108
		v94 = v110
		goto L30
	} else {
		goto L36
	}
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v108 = v107
	goto L34
L36:
	;
	goto L31
L37:
	;
	goto L2
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v133
	F_errmsg_internal(m, int32(_a_F_find_placeholders_recurse_0), v9)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_find_placeholders_recurse_1), int32(250), int32(_a_F_find_placeholders_recurse_2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_tabstat_entry(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_find_tabstat_entry[0]))
	v9 = base.I64_extend_i32_u(l0)
	v10 = F_pgstat_fetch_pending_entry(m, int32(2), v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v18 = F_pgstat_fetch_pending_entry(m, int32(2), int32(0), v9)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v56 = int32(0)
					return v56
				} else {
					v22 = v18
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v25 = F_palloc(m, int32(136))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						base.MemoryCopy(m, v25, v23, int32(136))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(0)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
						if v31 != 0 {
							v32 = *(*int64)(unsafe.Add(mBase, uint32(v25)+56))
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v25)+48))
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
							v36 = v31
							v37 = v32
							v38 = v33
							v39 = v34
							for {
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
								v41 = v39 + v40
								*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v41
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
								v44 = v38 + v43
								*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v44
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
								v47 = v37 + v46
								*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v47
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
								if v49 != 0 {
									v36 = v49
									v37 = v47
									v38 = v44
									v39 = v41
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v56 = v25
						return v56
					}
				}
			}
		} else {
			v22 = v10
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
			v25 = F_palloc(m, int32(136))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				base.MemoryCopy(m, v25, v23, int32(136))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(0)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				if v31 != 0 {
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v25)+56))
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v25)+48))
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
					v36 = v31
					v37 = v32
					v38 = v33
					v39 = v34
					for {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
						v41 = v39 + v40
						*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v41
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
						v44 = v38 + v43
						*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v44
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
						v47 = v37 + v46
						*(*int64)(unsafe.Add(mBase, uint32(v25)+56)) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
						if v49 != 0 {
							v36 = v49
							v37 = v47
							v38 = v44
							v39 = v41
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v56 = v25
				return v56
			}
		}
	}
}
func F_findoprnd_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	F_check_stack_depth(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v9) < base.Ui32(l2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = v9
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v19 = l0 + v15*int32(12)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	switch v20 - int32(1) {
	case 0:
		v43 = v15
		goto L9
	default:
		goto L11
	case 2:
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43 + int32(1)
	return
L10:
	;
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v43 = v42
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v23 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v26 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v28 + v26
	goto L8
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15 + int32(1)
	F_findoprnd_recurse(m, l0, l1, l2, l3)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v37 - v15
	goto L8
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v49) < base.Ui32(l2) {
		v15 = v49
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	F_errmsg_internal(m, int32(_a_F_findoprnd_recurse_0), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_findoprnd_recurse_1), int32(732), int32(_a_F_findoprnd_recurse_2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_finnish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v453 int32
	_ = v453
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1655 int32
	_ = v1655
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1773 int32
	_ = v1773
	var v1789 int32
	_ = v1789
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1904 int32
	_ = v1904
	var v1920 int32
	_ = v1920
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1952 int32
	_ = v1952
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1992 int32
	_ = v1992
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2008 int32
	_ = v2008
	var v2023 int32
	_ = v2023
	var v2032 int32
	_ = v2032
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2114 int32
	_ = v2114
	var v2122 int32
	_ = v2122
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2247 int32
	_ = v2247
	var v2263 int32
	_ = v2263
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2315 int32
	_ = v2315
	var v2318 int32
	_ = v2318
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = v12
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v506 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v506)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v508 < v510 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v129 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v129 = v101
	goto L2
L4:
	;
	if v9 <= v34 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v129 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v41 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v26))))
	if base.Ui32(v43) < base.Ui32(int32(192)) {
		v100 = v43
		v101 = v41
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(246) < v100 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v47 = v34 + int32(1)
	if v47 == v9 {
		v100 = v43
		v101 = v41
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v26))))
	v52 = v50 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v43) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v26))))
	v68 = v66 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v43) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v56 = v34 + int32(2)
	if v56 != v9 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
	v101 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v72))))
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
	v101 = int32(4)
	goto L9
L18:
	;
	v72 = v34 + int32(3)
	if v72 != v9 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v100 = v43<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v52<<(uint(int32(6))%32) | v68
	v101 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v118 = v101 + v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	v34 = v118
	goto L4
L23:
	;
	v105 = v100 - int32(97)
	if v105 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = v132 + v129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = v133
	goto L30
L28:
	;
	if v252 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v252 = v223
	goto L28
L30:
	;
	if v147 <= v156 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v252 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v163 = int32(1)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v148))))
	if base.Ui32(v165) < base.Ui32(int32(192)) {
		v222 = v165
		v223 = v163
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(246) < v222 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v169 = v156 + int32(1)
	if v169 == v147 {
		v222 = v165
		v223 = v163
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v148))))
	v174 = v172 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v165) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v148))))
	v190 = v188 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v165) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v178 = v156 + int32(2)
	if v178 != v147 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v222 = v165<<(uint(int32(6))%32)&int32(1984) | v174
	v223 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v194))))
	v222 = v207&int32(63) | (v165<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v174<<(uint(int32(12))%32) | v190<<(uint(int32(6))%32))
	v223 = int32(4)
	goto L35
L44:
	;
	v194 = v156 + int32(3)
	if v194 != v147 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v222 = v165<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v174<<(uint(int32(6))%32) | v190
	v223 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v227 = v222 - int32(97)
	if v227 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v227)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v233)>>(uint(v227&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v241 = v223 + v156
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241
	v156 = v241
	goto L30
L52:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v256 = v255 + v252
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v256
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = v256
	goto L55
L53:
	;
	if v375 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v375 = v347
	goto L53
L55:
	;
	if v271 <= v280 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v375 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v287 = int32(1)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v272))))
	if base.Ui32(v289) < base.Ui32(int32(192)) {
		v346 = v289
		v347 = v287
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(246) < v346 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v293 = v280 + int32(1)
	if v293 == v271 {
		v346 = v289
		v347 = v287
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v272))))
	v298 = v296 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v289) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v272))))
	v314 = v312 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v289) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v302 = v280 + int32(2)
	if v302 != v271 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v346 = v289<<(uint(int32(6))%32)&int32(1984) | v298
	v347 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v318))))
	v346 = v331&int32(63) | (v289<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v298<<(uint(int32(12))%32) | v314<<(uint(int32(6))%32))
	v347 = int32(4)
	goto L60
L69:
	;
	v318 = v280 + int32(3)
	if v318 != v271 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v346 = v289<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v298<<(uint(int32(6))%32) | v314
	v347 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v364 = v347 + v280
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v364
	v280 = v364
	goto L55
L74:
	;
	v351 = v346 - int32(97)
	if v351 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v357)>>(uint(v351&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v379 = v378 + v375
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v379
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v402 = v379
	goto L81
L79:
	;
	if v498 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v498 = v469
	goto L79
L81:
	;
	if v393 <= v402 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v498 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v409 = int32(1)
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v394))))
	if base.Ui32(v411) < base.Ui32(int32(192)) {
		v468 = v411
		v469 = v409
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(246) < v468 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v415 = v402 + int32(1)
	if v415 == v393 {
		v468 = v411
		v469 = v409
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v394))))
	v420 = v418 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v411) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v394))))
	v436 = v434 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v411) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v424 = v402 + int32(2)
	if v424 != v393 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v468 = v411<<(uint(int32(6))%32)&int32(1984) | v420
	v469 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394+v440))))
	v468 = v453&int32(63) | (v411<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v420<<(uint(int32(12))%32) | v436<<(uint(int32(6))%32))
	v469 = int32(4)
	goto L86
L95:
	;
	v440 = v402 + int32(3)
	if v440 != v393 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v468 = v411<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v420<<(uint(int32(6))%32) | v436
	v469 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v473 = v468 - int32(97)
	if v473 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v473)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v479)>>(uint(v473&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v487 = v469 + v402
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v487
	v402 = v487
	goto L81
L103:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v501 + v498
	goto L1
L104:
	;
	return v2318
L105:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v667
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v667 < v669 {
		goto L141
	} else {
		goto L142
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v510
	v517 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_2), int32(10), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	return int32(0)
L108:
	;
	if v517 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	goto L105
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v525
	switch v517 - int32(1) {
	case 0:
		goto L114
	case 1:
		goto L113
	default:
		goto L112
	}
L112:
	;
	v662 = F_slice_del(m, l0)
	mBase = m.M
	if v662 < int32(0) {
		v2318 = v662
		goto L104
	} else {
		goto L140
	}
L113:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v525 < v660 {
		goto L105
	} else {
		goto L139
	}
L114:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L117
L115:
	;
	if v657 == int32(0) {
		goto L112
	} else {
		goto L138
	}
L116:
	;
	v657 = v650
	goto L115
L117:
	;
	if v541 <= v542 {
		v650 = int32(-1)
		goto L116
	} else {
		goto L119
	}
L118:
	;
	v650 = int32(0)
	goto L116
L119:
	;
	v559 = int32(1)
	v560 = v541 - v559
	v562 = int32(*(*int8)(unsafe.Add(mBase, uint32(v543+v560))))
	v564 = v562 & int32(255)
	if base.B2i32(v560 == v542)|base.B2i32(int32(0) <= v562) != 0 {
		v622 = v564
		v626 = v559
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if int32(246) < v622 {
		goto L128
	} else {
		goto L129
	}
L121:
	;
	v571 = v564 & int32(63)
	v573 = v541 - int32(2)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v573))))
	v577 = v575 << (uint(int32(6)) % 32)
	if base.B2i32(v573 != v542)&base.B2i32(base.Ui32(v575) < base.Ui32(int32(192))) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v622 = v577&int32(1984) | v571
	v626 = int32(2)
	goto L120
L123:
	;
	goto L124
L124:
	;
	v590 = v577&int32(4032) | v571
	v592 = v541 - int32(3)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v592))))
	if base.B2i32(v592 != v542)&base.B2i32(base.Ui32(v594) < base.Ui32(int32(224))) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v622 = v594<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v590
	v626 = int32(3)
	goto L120
L126:
	;
	goto L127
L127:
	;
	v612 = int32(4)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+v543-v612))))
	v622 = v594<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v614&int32(7)<<(uint(int32(18))%32) | v590
	v626 = v612
	goto L120
L128:
	;
	v657 = v626
	goto L115
L129:
	;
	goto L130
L130:
	;
	v628 = v622 - int32(97)
	if v628 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v657 = v626
	goto L115
L132:
	;
	goto L133
L133:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v628)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[1]))))
	if int32(base.Ui32(v634)>>(uint(v628&int32(7))%32))&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v657 = v626
	goto L115
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541 - v626
	goto L137
L137:
	;
	goto L118
L138:
	;
	goto L105
L139:
	;
	goto L112
L140:
	;
	goto L105
L141:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v793
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v793 < v795 {
		goto L182
	} else {
		goto L183
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v667
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v669
	v677 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_4), int32(9), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L107
	} else {
		goto L143
	}
L143:
	;
	if v677 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v672
	goto L141
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v672
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v683
	switch v677 - int32(1) {
	case 0:
		goto L152
	case 1:
		goto L151
	case 2:
		goto L150
	case 3:
		goto L149
	case 4:
		goto L148
	case 5:
		goto L147
	default:
		goto L141
	}
L147:
	;
	if v683-int32(2) <= v672 {
		goto L141
	} else {
		goto L177
	}
L148:
	;
	if v683-int32(2) <= v672 {
		goto L141
	} else {
		goto L172
	}
L149:
	;
	v733 = v683 - int32(1)
	if v733 <= v672 {
		goto L141
	} else {
		goto L167
	}
L150:
	;
	v729 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v729 {
		goto L141
	} else {
		goto L166
	}
L151:
	;
	v698 = F_slice_del(m, l0)
	mBase = m.M
	if v698 < int32(0) {
		v2318 = v698
		goto L104
	} else {
		goto L158
	}
L152:
	;
	if v672 < v683 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688+v683-int32(1)))))
	if v692 == int32(107) {
		goto L141
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v695 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v695 {
		goto L141
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v2318 = v695
	goto L104
L158:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v701
	v703 = int32(3)
	v705 = int32(0)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v701-v708 < v703 {
		v718 = v705
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v718 == int32(0) {
		goto L141
	} else {
		goto L163
	}
L160:
	;
	goto L159
L161:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = F_memcmp(m, v711+v701-v703, int32(_a_F_finnish_UTF_8_stem_5), v703)
	mBase = m.M
	if v714 != 0 {
		v718 = v705
		goto L160
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v701 - v703
	v718 = int32(1)
	goto L160
L163:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v721
	v725 = F_slice_from_s(m, l0, int32(3), int32(_a_F_finnish_UTF_8_stem_6))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L107
	} else {
		goto L164
	}
L164:
	;
	if int32(0) <= v725 {
		goto L141
	} else {
		goto L165
	}
L165:
	;
	v2318 = v725
	goto L104
L166:
	;
	v2318 = v729
	goto L104
L167:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735+v733))))
	if v737 != int32(97) {
		goto L141
	} else {
		goto L168
	}
L168:
	;
	v743 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_7), int32(6), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L107
	} else {
		goto L169
	}
L169:
	;
	if v743 == int32(0) {
		goto L141
	} else {
		goto L170
	}
L170:
	;
	v747 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v747 {
		goto L141
	} else {
		goto L171
	}
L171:
	;
	v2318 = v747
	goto L104
L172:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v683-int32(1)))))
	if v757 != int32(164) {
		goto L141
	} else {
		goto L173
	}
L173:
	;
	v763 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_8), int32(6), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L107
	} else {
		goto L174
	}
L174:
	;
	if v763 == int32(0) {
		goto L141
	} else {
		goto L175
	}
L175:
	;
	v767 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v767 {
		goto L141
	} else {
		goto L176
	}
L176:
	;
	v2318 = v767
	goto L104
L177:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v683-int32(1)))))
	if v777 != int32(101) {
		goto L141
	} else {
		goto L178
	}
L178:
	;
	v783 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_9), int32(2), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L107
	} else {
		goto L179
	}
L179:
	;
	if v783 == int32(0) {
		goto L141
	} else {
		goto L180
	}
L180:
	;
	v787 = F_slice_del(m, l0)
	mBase = m.M
	if v787 < int32(0) {
		v2318 = v787
		goto L104
	} else {
		goto L181
	}
L181:
	;
	goto L141
L182:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1250
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1250 < v1252 {
		goto L294
	} else {
		goto L295
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v793
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v795
	v803 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_10), int32(30), int32(_a_F_finnish_UTF_8_stem_11))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L107
	} else {
		goto L184
	}
L184:
	;
	if v803 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v798
	goto L182
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v798
	v809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v809
	switch v803 - int32(1) {
	case 0:
		goto L196
	case 1:
		goto L195
	case 2:
		goto L194
	case 3:
		goto L193
	case 4:
		goto L192
	case 5:
		goto L191
	case 6:
		goto L190
	case 7:
		goto L189
	default:
		goto L188
	}
L188:
	;
	v1242 = F_slice_del(m, l0)
	mBase = m.M
	if v1242 < int32(0) {
		v2318 = v1242
		goto L104
	} else {
		goto L293
	}
L189:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L247
L190:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v890 = v889 - v809
	v894 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_12), int32(7), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L107
	} else {
		goto L216
	}
L191:
	;
	v873 = int32(2)
	v875 = int32(0)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v877-v878 < v873 {
		v888 = v875
		goto L211
	} else {
		goto L212
	}
L192:
	;
	v857 = int32(2)
	v859 = int32(0)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v861-v862 < v857 {
		v872 = v859
		goto L206
	} else {
		goto L207
	}
L193:
	;
	if v809 <= v798 {
		goto L182
	} else {
		goto L203
	}
L194:
	;
	if v809 <= v798 {
		goto L182
	} else {
		goto L201
	}
L195:
	;
	if v809 <= v798 {
		goto L182
	} else {
		goto L199
	}
L196:
	;
	if v809 <= v798 {
		goto L182
	} else {
		goto L197
	}
L197:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v814+v809-int32(1)))))
	if v818 != int32(97) {
		goto L182
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	goto L188
L199:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825+v809-int32(1)))))
	if v829 != int32(101) {
		goto L182
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	goto L188
L201:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836+v809-int32(1)))))
	if v840 != int32(105) {
		goto L182
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	goto L188
L203:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847+v809-int32(1)))))
	if v851 != int32(111) {
		goto L182
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v809 - int32(1)
	goto L188
L205:
	;
	if v872 != 0 {
		goto L188
	} else {
		goto L209
	}
L206:
	;
	goto L205
L207:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v868 = F_memcmp(m, v865+v861-v857, int32(_a_F_finnish_UTF_8_stem_13), v857)
	mBase = m.M
	if v868 != 0 {
		v872 = v859
		goto L206
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v861 - v857
	v872 = int32(1)
	goto L206
L209:
	;
	goto L182
L210:
	;
	if v888 != 0 {
		goto L188
	} else {
		goto L214
	}
L211:
	;
	goto L210
L212:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v884 = F_memcmp(m, v881+v877-v873, int32(_a_F_finnish_UTF_8_stem_14), v873)
	mBase = m.M
	if v884 != 0 {
		v888 = v875
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v877 - v873
	v888 = int32(1)
	goto L211
L214:
	;
	goto L182
L215:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v919 = v918 - v890
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v919
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L225
L216:
	;
	if v894 != 0 {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v897 = v896 - v890
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v897
	v899 = int32(2)
	v901 = int32(0)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v897-v904 < v899 {
		v914 = v901
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if v914 != 0 {
		goto L215
	} else {
		goto L222
	}
L219:
	;
	goto L218
L220:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v910 = F_memcmp(m, v907+v897-v899, int32(_a_F_finnish_UTF_8_stem_15), v899)
	mBase = m.M
	if v910 != 0 {
		v914 = v901
		goto L219
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v897 - v899
	v914 = int32(1)
	goto L219
L222:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v915 - v890
	goto L188
L223:
	;
	if v974 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L225:
	;
	goto L226
L226:
	;
	goto L227
L227:
	;
	v929 = v919
	v931 = int32(1)
	goto L230
L229:
	;
	v974 = v956
	goto L223
L230:
	;
	if v929 <= v922 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L229
L232:
	;
	v974 = int32(-1)
	goto L223
L233:
	;
	goto L234
L234:
	;
	v936 = v929 - int32(1)
	v938 = int32(*(*int8)(unsafe.Add(mBase, uint32(v921+v936))))
	if base.B2i32(int32(0) <= v938)|base.B2i32(v936 <= v922) != 0 {
		v956 = v936
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v960 = int32(1)
	if v960 < v931 {
		v929 = v956
		v931 = v931 - v960
		goto L230
	} else {
		goto L241
	}
L236:
	;
	v944 = v936
	goto L237
L237:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921+v944))))
	if base.Ui32(int32(191)) < base.Ui32(v949) {
		v956 = v944
		goto L235
	} else {
		goto L239
	}
L238:
	;
	v956 = v922
	goto L235
L239:
	;
	v953 = v944 - int32(1)
	if v922 < v953 {
		v944 = v953
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	goto L231
L242:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977 - v890
	goto L188
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v974
	goto L188
L245:
	;
	if v1110 != 0 {
		goto L182
	} else {
		goto L268
	}
L246:
	;
	v1110 = v1103
	goto L245
L247:
	;
	if v994 <= v995 {
		v1103 = int32(-1)
		goto L246
	} else {
		goto L249
	}
L248:
	;
	v1103 = int32(0)
	goto L246
L249:
	;
	v1012 = int32(1)
	v1013 = v994 - v1012
	v1015 = int32(*(*int8)(unsafe.Add(mBase, uint32(v996+v1013))))
	v1017 = v1015 & int32(255)
	if base.B2i32(v1013 == v995)|base.B2i32(int32(0) <= v1015) != 0 {
		v1075 = v1017
		v1079 = v1012
		goto L250
	} else {
		goto L251
	}
L250:
	;
	if int32(246) < v1075 {
		goto L258
	} else {
		goto L259
	}
L251:
	;
	v1024 = v1017 & int32(63)
	v1026 = v994 - int32(2)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996+v1026))))
	v1030 = v1028 << (uint(int32(6)) % 32)
	if base.B2i32(v1026 != v995)&base.B2i32(base.Ui32(v1028) < base.Ui32(int32(192))) == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1075 = v1030&int32(1984) | v1024
	v1079 = int32(2)
	goto L250
L253:
	;
	goto L254
L254:
	;
	v1043 = v1030&int32(4032) | v1024
	v1045 = v994 - int32(3)
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996+v1045))))
	if base.B2i32(v1045 != v995)&base.B2i32(base.Ui32(v1047) < base.Ui32(int32(224))) == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1075 = v1047<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1043
	v1079 = int32(3)
	goto L250
L256:
	;
	goto L257
L257:
	;
	v1065 = int32(4)
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994+v996-v1065))))
	v1075 = v1047<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1067&int32(7)<<(uint(int32(18))%32) | v1043
	v1079 = v1065
	goto L250
L258:
	;
	v1110 = v1079
	goto L245
L259:
	;
	goto L260
L260:
	;
	v1081 = v1075 - int32(97)
	if v1081 < int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1110 = v1079
	goto L245
L262:
	;
	goto L263
L263:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1081)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1087)>>(uint(v1081&int32(7))%32))&int32(1) == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1110 = v1079
	goto L245
L265:
	;
	goto L266
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v994 - v1079
	goto L267
L267:
	;
	goto L248
L268:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L271
L269:
	;
	if v1239 != 0 {
		goto L182
	} else {
		goto L292
	}
L270:
	;
	v1239 = v1232
	goto L269
L271:
	;
	if v1123 <= v1124 {
		v1232 = int32(-1)
		goto L270
	} else {
		goto L273
	}
L272:
	;
	v1232 = int32(0)
	goto L270
L273:
	;
	v1141 = int32(1)
	v1142 = v1123 - v1141
	v1144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1125+v1142))))
	v1146 = v1144 & int32(255)
	if base.B2i32(v1142 == v1124)|base.B2i32(int32(0) <= v1144) != 0 {
		v1204 = v1146
		v1208 = v1141
		goto L274
	} else {
		goto L275
	}
L274:
	;
	if int32(122) < v1204 {
		goto L282
	} else {
		goto L283
	}
L275:
	;
	v1153 = v1146 & int32(63)
	v1155 = v1123 - int32(2)
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125+v1155))))
	v1159 = v1157 << (uint(int32(6)) % 32)
	if base.B2i32(v1155 != v1124)&base.B2i32(base.Ui32(v1157) < base.Ui32(int32(192))) == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1204 = v1159&int32(1984) | v1153
	v1208 = int32(2)
	goto L274
L277:
	;
	goto L278
L278:
	;
	v1172 = v1159&int32(4032) | v1153
	v1174 = v1123 - int32(3)
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125+v1174))))
	if base.B2i32(v1174 != v1124)&base.B2i32(base.Ui32(v1176) < base.Ui32(int32(224))) == int32(0) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1204 = v1176<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1172
	v1208 = int32(3)
	goto L274
L280:
	;
	goto L281
L281:
	;
	v1194 = int32(4)
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123+v1125-v1194))))
	v1204 = v1176<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1196&int32(7)<<(uint(int32(18))%32) | v1172
	v1208 = v1194
	goto L274
L282:
	;
	v1239 = v1208
	goto L269
L283:
	;
	goto L284
L284:
	;
	v1210 = v1204 - int32(98)
	if v1210 < int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1239 = v1208
	goto L269
L286:
	;
	goto L287
L287:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1210)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1216)>>(uint(v1210&int32(7))%32))&int32(1) == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1239 = v1208
	goto L269
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1123 - v1208
	goto L291
L291:
	;
	goto L272
L292:
	;
	goto L188
L293:
	;
	v1245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1245)
	goto L182
L294:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1297
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v1299 != 0 {
		goto L310
	} else {
		goto L311
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1250
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1252
	v1260 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_16), int32(14), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L107
	} else {
		goto L296
	}
L296:
	;
	if v1260 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1255
	goto L294
L298:
	;
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1255
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1266
	if v1260 == int32(1) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1271 = int32(2)
	v1273 = int32(0)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1275-v1276 < v1271 {
		v1286 = v1273
		goto L304
	} else {
		goto L305
	}
L301:
	;
	goto L302
L302:
	;
	v1292 = F_slice_del(m, l0)
	mBase = m.M
	if v1292 < int32(0) {
		v2318 = v1292
		goto L104
	} else {
		goto L308
	}
L303:
	;
	if v1286 != 0 {
		goto L294
	} else {
		goto L307
	}
L304:
	;
	goto L303
L305:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1282 = F_memcmp(m, v1279+v1275-v1271, int32(_a_F_finnish_UTF_8_stem_17), v1271)
	mBase = m.M
	if v1282 != 0 {
		v1286 = v1273
		goto L304
	} else {
		goto L306
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1275 - v1271
	v1286 = int32(1)
	goto L304
L307:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1287 + (v1266 - v1270)
	goto L302
L308:
	;
	goto L294
L309:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1581
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1581 < v1585 {
		v2306 = int32(0)
		goto L383
	} else {
		goto L384
	}
L310:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1304 <= v1303 {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	goto L312
L312:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1347 < v1348 {
		v1570 = int32(0)
		goto L327
	} else {
		goto L328
	}
L313:
	;
	if int32(0) <= v1344 {
		goto L309
	} else {
		goto L325
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1303
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1304
	if v1304 < v1303 {
		goto L318
	} else {
		goto L319
	}
L315:
	;
	v1340 = int32(0)
	goto L316
L316:
	;
	v1344 = v1340
	goto L313
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1307
	v1324 = int32(1)
	v1325 = v1303 - v1324
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1325
	v1329 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1329 {
		goto L322
	} else {
		goto L323
	}
L318:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1310+v1303-int32(1)))))
	if base.Ui32((v1314-int32(105))&int32(255)) < base.Ui32(int32(2)) {
		goto L317
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1307
	v1344 = int32(0)
	goto L313
L321:
	;
	goto L320
L322:
	;
	v1335 = v1324
	goto L324
L323:
	;
	v1335 = v1329 >> (uint(int32(31)) % 32) & v1329
	goto L324
L324:
	;
	v1340 = v1335
	goto L316
L325:
	;
	v2318 = v1344
	goto L104
L326:
	;
	if v1574 < int32(0) {
		v2318 = v1574
		goto L104
	} else {
		goto L382
	}
L327:
	;
	v1574 = v1570
	goto L326
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1347
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1348
	if v1348 < v1347 {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	v1364 = v1347 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1364
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1364
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L336
L330:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1354+v1347-int32(1)))))
	if v1358 == int32(116) {
		goto L329
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1574 = int32(0)
	goto L326
L333:
	;
	goto L332
L334:
	;
	if v1496 != 0 {
		goto L357
	} else {
		goto L358
	}
L335:
	;
	v1496 = v1489
	goto L334
L336:
	;
	if v1364 <= v1381 {
		v1489 = int32(-1)
		goto L335
	} else {
		goto L338
	}
L337:
	;
	v1489 = int32(0)
	goto L335
L338:
	;
	v1398 = int32(1)
	v1399 = v1364 - v1398
	v1401 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1382+v1399))))
	v1403 = v1401 & int32(255)
	if base.B2i32(v1399 == v1381)|base.B2i32(int32(0) <= v1401) != 0 {
		v1461 = v1403
		v1465 = v1398
		goto L339
	} else {
		goto L340
	}
L339:
	;
	if int32(246) < v1461 {
		goto L347
	} else {
		goto L348
	}
L340:
	;
	v1410 = v1403 & int32(63)
	v1412 = v1364 - int32(2)
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1412))))
	v1416 = v1414 << (uint(int32(6)) % 32)
	if base.B2i32(v1412 != v1381)&base.B2i32(base.Ui32(v1414) < base.Ui32(int32(192))) == int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1461 = v1416&int32(1984) | v1410
	v1465 = int32(2)
	goto L339
L342:
	;
	goto L343
L343:
	;
	v1429 = v1416&int32(4032) | v1410
	v1431 = v1364 - int32(3)
	v1433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1431))))
	if base.B2i32(v1431 != v1381)&base.B2i32(base.Ui32(v1433) < base.Ui32(int32(224))) == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1461 = v1433<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1429
	v1465 = int32(3)
	goto L339
L345:
	;
	goto L346
L346:
	;
	v1451 = int32(4)
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364+v1382-v1451))))
	v1461 = v1433<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1453&int32(7)<<(uint(int32(18))%32) | v1429
	v1465 = v1451
	goto L339
L347:
	;
	v1496 = v1465
	goto L334
L348:
	;
	goto L349
L349:
	;
	v1467 = v1461 - int32(97)
	if v1467 < int32(0) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1496 = v1465
	goto L334
L351:
	;
	goto L352
L352:
	;
	v1473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1467)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1473)>>(uint(v1467&int32(7))%32))&int32(1) == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1496 = v1465
	goto L334
L354:
	;
	goto L355
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1364 - v1465
	goto L356
L356:
	;
	goto L337
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1574 = int32(0)
	goto L326
L358:
	;
	goto L359
L359:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499 + (v1364 - v1367)
	v1503 = F_slice_del(m, l0)
	mBase = m.M
	if v1503 < int32(0) {
		v1570 = v1503
		goto L327
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1508 < v1509 {
		v1574 = int32(0)
		goto L326
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1508
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1509
	if v1509 < v1508-int32(2) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1528 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_18), int32(2), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L107
	} else {
		goto L367
	}
L363:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1516+v1508-int32(1)))))
	if v1520 == int32(97) {
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1574 = int32(0)
	goto L326
L366:
	;
	goto L365
L367:
	;
	if v1528 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1574 = int32(0)
	goto L326
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1351
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1535
	if v1528 == int32(1) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1540 = int32(0)
	v1541 = int32(2)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1545-v1546 < v1541 {
		v1556 = v1540
		goto L375
	} else {
		goto L376
	}
L372:
	;
	goto L373
L373:
	;
	v1564 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1564 {
		goto L379
	} else {
		goto L380
	}
L374:
	;
	if v1556 != 0 {
		v1570 = v1540
		goto L327
	} else {
		goto L378
	}
L375:
	;
	goto L374
L376:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1552 = F_memcmp(m, v1549+v1545-v1541, int32(_a_F_finnish_UTF_8_stem_19), v1541)
	mBase = m.M
	if v1552 != 0 {
		v1556 = v1540
		goto L375
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1545 - v1541
	v1556 = int32(1)
	goto L375
L378:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1557 + (v1535 - v1539)
	goto L373
L379:
	;
	v1567 = int32(1)
	goto L381
L380:
	;
	v1567 = v1564
	goto L381
L381:
	;
	v1570 = v1567
	goto L327
L382:
	;
	goto L309
L383:
	;
	if v2306 < int32(0) {
		v2318 = v2306
		goto L104
	} else {
		goto L532
	}
L384:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1585
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1590 = v1589 - v1581
	v1594 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_12), int32(7), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L107
	} else {
		goto L386
	}
L385:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1665 = v1664 - v1590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1665
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1665
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L412
L386:
	;
	if v1594 == int32(0) {
		goto L385
	} else {
		goto L387
	}
L387:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1599 = v1598 - v1590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1599
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L390
L388:
	;
	if v1655 < int32(0) {
		goto L385
	} else {
		goto L407
	}
L390:
	;
	goto L391
L391:
	;
	goto L392
L392:
	;
	v1610 = v1599
	v1612 = int32(1)
	goto L395
L394:
	;
	v1655 = v1637
	goto L388
L395:
	;
	if v1610 <= v1603 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L394
L397:
	;
	v1655 = int32(-1)
	goto L388
L398:
	;
	goto L399
L399:
	;
	v1617 = v1610 - int32(1)
	v1619 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1602+v1617))))
	if base.B2i32(int32(0) <= v1619)|base.B2i32(v1617 <= v1603) != 0 {
		v1637 = v1617
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1641 = int32(1)
	if v1641 < v1612 {
		v1610 = v1637
		v1612 = v1612 - v1641
		goto L395
	} else {
		goto L406
	}
L401:
	;
	v1625 = v1617
	goto L402
L402:
	;
	v1630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1602+v1625))))
	if base.Ui32(int32(191)) < base.Ui32(v1630) {
		v1637 = v1625
		goto L400
	} else {
		goto L404
	}
L403:
	;
	v1637 = v1603
	goto L400
L404:
	;
	v1634 = v1625 - int32(1)
	if v1603 < v1634 {
		v1625 = v1634
		goto L402
	} else {
		goto L405
	}
L405:
	;
	goto L403
L406:
	;
	goto L396
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1655
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1655
	v1660 = F_slice_del(m, l0)
	mBase = m.M
	if v1660 < int32(0) {
		v2306 = v1660
		goto L383
	} else {
		goto L408
	}
L408:
	;
	goto L385
L409:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1933 = v1932 - v1590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1933
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1933
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1933 <= v1936 {
		v1970 = v1933
		v1972 = v1936
		goto L459
	} else {
		goto L460
	}
L410:
	;
	if v1796 != 0 {
		goto L409
	} else {
		goto L433
	}
L411:
	;
	v1796 = v1789
	goto L410
L412:
	;
	if v1665 <= v1681 {
		v1789 = int32(-1)
		goto L411
	} else {
		goto L414
	}
L413:
	;
	v1789 = int32(0)
	goto L411
L414:
	;
	v1698 = int32(1)
	v1699 = v1665 - v1698
	v1701 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1682+v1699))))
	v1703 = v1701 & int32(255)
	if base.B2i32(v1699 == v1681)|base.B2i32(int32(0) <= v1701) != 0 {
		v1761 = v1703
		v1765 = v1698
		goto L415
	} else {
		goto L416
	}
L415:
	;
	if int32(228) < v1761 {
		goto L423
	} else {
		goto L424
	}
L416:
	;
	v1710 = v1703 & int32(63)
	v1712 = v1665 - int32(2)
	v1714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1682+v1712))))
	v1716 = v1714 << (uint(int32(6)) % 32)
	if base.B2i32(v1712 != v1681)&base.B2i32(base.Ui32(v1714) < base.Ui32(int32(192))) == int32(0) {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v1761 = v1716&int32(1984) | v1710
	v1765 = int32(2)
	goto L415
L418:
	;
	goto L419
L419:
	;
	v1729 = v1716&int32(4032) | v1710
	v1731 = v1665 - int32(3)
	v1733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1682+v1731))))
	if base.B2i32(v1731 != v1681)&base.B2i32(base.Ui32(v1733) < base.Ui32(int32(224))) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1761 = v1733<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1729
	v1765 = int32(3)
	goto L415
L421:
	;
	goto L422
L422:
	;
	v1751 = int32(4)
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1665+v1682-v1751))))
	v1761 = v1733<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1753&int32(7)<<(uint(int32(18))%32) | v1729
	v1765 = v1751
	goto L415
L423:
	;
	v1796 = v1765
	goto L410
L424:
	;
	goto L425
L425:
	;
	v1767 = v1761 - int32(97)
	if v1767 < int32(0) {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v1796 = v1765
	goto L410
L427:
	;
	goto L428
L428:
	;
	v1773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1767)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[3]))))
	if int32(base.Ui32(v1773)>>(uint(v1767&int32(7))%32))&int32(1) == int32(0) {
		goto L429
	} else {
		goto L430
	}
L429:
	;
	v1796 = v1765
	goto L410
L430:
	;
	goto L431
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1665 - v1765
	goto L432
L432:
	;
	goto L413
L433:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1797
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L436
L434:
	;
	if v1927 != 0 {
		goto L409
	} else {
		goto L457
	}
L435:
	;
	v1927 = v1920
	goto L434
L436:
	;
	if v1797 <= v1812 {
		v1920 = int32(-1)
		goto L435
	} else {
		goto L438
	}
L437:
	;
	v1920 = int32(0)
	goto L435
L438:
	;
	v1829 = int32(1)
	v1830 = v1797 - v1829
	v1832 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1813+v1830))))
	v1834 = v1832 & int32(255)
	if base.B2i32(v1830 == v1812)|base.B2i32(int32(0) <= v1832) != 0 {
		v1892 = v1834
		v1896 = v1829
		goto L439
	} else {
		goto L440
	}
L439:
	;
	if int32(122) < v1892 {
		goto L447
	} else {
		goto L448
	}
L440:
	;
	v1841 = v1834 & int32(63)
	v1843 = v1797 - int32(2)
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1813+v1843))))
	v1847 = v1845 << (uint(int32(6)) % 32)
	if base.B2i32(v1843 != v1812)&base.B2i32(base.Ui32(v1845) < base.Ui32(int32(192))) == int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v1892 = v1847&int32(1984) | v1841
	v1896 = int32(2)
	goto L439
L442:
	;
	goto L443
L443:
	;
	v1860 = v1847&int32(4032) | v1841
	v1862 = v1797 - int32(3)
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1813+v1862))))
	if base.B2i32(v1862 != v1812)&base.B2i32(base.Ui32(v1864) < base.Ui32(int32(224))) == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1892 = v1864<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1860
	v1896 = int32(3)
	goto L439
L445:
	;
	goto L446
L446:
	;
	v1882 = int32(4)
	v1884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1797+v1813-v1882))))
	v1892 = v1864<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1884&int32(7)<<(uint(int32(18))%32) | v1860
	v1896 = v1882
	goto L439
L447:
	;
	v1927 = v1896
	goto L434
L448:
	;
	goto L449
L449:
	;
	v1898 = v1892 - int32(98)
	if v1898 < int32(0) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1927 = v1896
	goto L434
L451:
	;
	goto L452
L452:
	;
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1898)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1904)>>(uint(v1898&int32(7))%32))&int32(1) == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1927 = v1896
	goto L434
L454:
	;
	goto L455
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1797 - v1896
	goto L456
L456:
	;
	goto L437
L457:
	;
	v1928 = F_slice_del(m, l0)
	mBase = m.M
	if v1928 < int32(0) {
		v2306 = v1928
		goto L383
	} else {
		goto L458
	}
L458:
	;
	goto L409
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1970
	if v1970 <= v1972 {
		v2003 = v1970
		goto L468
	} else {
		goto L469
	}
L460:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1939 = v1938 + v1933
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1939-int32(1)))))
	if v1942 != int32(106) {
		v1970 = v1933
		v1972 = v1936
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1946 = v1933 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1946
	if v1946 <= v1936 {
		v1970 = v1933
		v1972 = v1936
		goto L459
	} else {
		goto L462
	}
L462:
	;
	v1952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1939-int32(2)))))
	if v1952 != int32(111) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1946+v1938-int32(1)))))
	if v1958 != int32(117) {
		v1970 = v1933
		v1972 = v1936
		goto L459
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1933 - int32(2)
	v1964 = F_slice_del(m, l0)
	mBase = m.M
	if v1964 < int32(0) {
		v2306 = v1964
		goto L383
	} else {
		goto L467
	}
L466:
	;
	goto L465
L467:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1970 = v1967 - v1590
	v1972 = v1969
	goto L459
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1587
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2003
	v2008 = int32(0)
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2032 = v2003
	goto L476
L469:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1979 = v1978 + v1970
	v1982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1979-int32(1)))))
	if v1982 != int32(111) {
		v2003 = v1970
		goto L468
	} else {
		goto L470
	}
L470:
	;
	v1986 = v1970 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1986
	if v1986 <= v1972 {
		v2003 = v1970
		goto L468
	} else {
		goto L471
	}
L471:
	;
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1979-int32(2)))))
	if v1992 != int32(106) {
		v2003 = v1970
		goto L468
	} else {
		goto L472
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1970 - int32(2)
	v1998 = F_slice_del(m, l0)
	mBase = m.M
	if v1998 < int32(0) {
		v2306 = v1998
		goto L383
	} else {
		goto L473
	}
L473:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2003 = v2001 - v1590
	goto L468
L474:
	;
	if v2137 < int32(0) {
		v2306 = v2008
		goto L383
	} else {
		goto L497
	}
L475:
	;
	v2137 = int32(-1)
	goto L474
L476:
	;
	if v2032 <= v1587 {
		goto L475
	} else {
		goto L478
	}
L478:
	;
	v2039 = int32(1)
	v2040 = v2032 - v2039
	v2042 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2023+v2040))))
	v2044 = v2042 & int32(255)
	if base.B2i32(v2040 == v1587)|base.B2i32(int32(0) <= v2042) != 0 {
		v2102 = v2044
		v2106 = v2039
		goto L479
	} else {
		goto L480
	}
L479:
	;
	if int32(246) < v2102 {
		goto L487
	} else {
		goto L488
	}
L480:
	;
	v2051 = v2044 & int32(63)
	v2053 = v2032 - int32(2)
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2023+v2053))))
	v2057 = v2055 << (uint(int32(6)) % 32)
	if base.B2i32(v2053 != v1587)&base.B2i32(base.Ui32(v2055) < base.Ui32(int32(192))) == int32(0) {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v2102 = v2057&int32(1984) | v2051
	v2106 = int32(2)
	goto L479
L482:
	;
	goto L483
L483:
	;
	v2070 = v2057&int32(4032) | v2051
	v2072 = v2032 - int32(3)
	v2074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2023+v2072))))
	if base.B2i32(v2072 != v1587)&base.B2i32(base.Ui32(v2074) < base.Ui32(int32(224))) == int32(0) {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2102 = v2074<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v2070
	v2106 = int32(3)
	goto L479
L485:
	;
	goto L486
L486:
	;
	v2092 = int32(4)
	v2094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032+v2023-v2092))))
	v2102 = v2074<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v2094&int32(7)<<(uint(int32(18))%32) | v2070
	v2106 = v2092
	goto L479
L487:
	;
	v2137 = v2106
	goto L474
L488:
	;
	goto L489
L489:
	;
	v2108 = v2102 - int32(97)
	if v2108 < int32(0) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2137 = v2106
	goto L474
L491:
	;
	goto L492
L492:
	;
	v2114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2108)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v2114)>>(uint(v2108&int32(7))%32))&int32(1) == int32(0) {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v2137 = v2106
	goto L474
L494:
	;
	goto L495
L495:
	;
	v2122 = v2032 - v2106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2122
	v2032 = v2122
	goto L476
L497:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2140
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L500
L498:
	;
	if v2270 != 0 {
		v2306 = v2008
		goto L383
	} else {
		goto L521
	}
L499:
	;
	v2270 = v2263
	goto L498
L500:
	;
	if v2140 <= v2155 {
		v2263 = int32(-1)
		goto L499
	} else {
		goto L502
	}
L501:
	;
	v2263 = int32(0)
	goto L499
L502:
	;
	v2172 = int32(1)
	v2173 = v2140 - v2172
	v2175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2156+v2173))))
	v2177 = v2175 & int32(255)
	if base.B2i32(v2173 == v2155)|base.B2i32(int32(0) <= v2175) != 0 {
		v2235 = v2177
		v2239 = v2172
		goto L503
	} else {
		goto L504
	}
L503:
	;
	if int32(122) < v2235 {
		goto L511
	} else {
		goto L512
	}
L504:
	;
	v2184 = v2177 & int32(63)
	v2186 = v2140 - int32(2)
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156+v2186))))
	v2190 = v2188 << (uint(int32(6)) % 32)
	if base.B2i32(v2186 != v2155)&base.B2i32(base.Ui32(v2188) < base.Ui32(int32(192))) == int32(0) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2235 = v2190&int32(1984) | v2184
	v2239 = int32(2)
	goto L503
L506:
	;
	goto L507
L507:
	;
	v2203 = v2190&int32(4032) | v2184
	v2205 = v2140 - int32(3)
	v2207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156+v2205))))
	if base.B2i32(v2205 != v2155)&base.B2i32(base.Ui32(v2207) < base.Ui32(int32(224))) == int32(0) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2235 = v2207<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v2203
	v2239 = int32(3)
	goto L503
L509:
	;
	goto L510
L510:
	;
	v2225 = int32(4)
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140+v2156-v2225))))
	v2235 = v2207<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v2227&int32(7)<<(uint(int32(18))%32) | v2203
	v2239 = v2225
	goto L503
L511:
	;
	v2270 = v2239
	goto L498
L512:
	;
	goto L513
L513:
	;
	v2241 = v2235 - int32(98)
	if v2241 < int32(0) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2270 = v2239
	goto L498
L515:
	;
	goto L516
L516:
	;
	v2247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2241)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v2247)>>(uint(v2241&int32(7))%32))&int32(1) == int32(0) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2270 = v2239
	goto L498
L518:
	;
	goto L519
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2140 - v2239
	goto L520
L520:
	;
	goto L501
L521:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2271
	v2275 = F_slice_to(m, l0, l0+int32(40))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L107
	} else {
		goto L522
	}
L522:
	;
	if v2275 < int32(0) {
		v2306 = v2275
		goto L383
	} else {
		goto L523
	}
L523:
	;
	v2279 = int32(0)
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2280-int32(4))))
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2287-v2288 < v2286 {
		v2298 = v2279
		goto L525
	} else {
		goto L526
	}
L524:
	;
	if v2298 == int32(0) {
		v2306 = v2279
		goto L383
	} else {
		goto L528
	}
L525:
	;
	goto L524
L526:
	;
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2294 = F_memcmp(m, v2291+v2287-v2286, v2280, v2286)
	mBase = m.M
	if v2294 != 0 {
		v2298 = v2279
		goto L525
	} else {
		goto L527
	}
L527:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2287 - v2286
	v2298 = int32(1)
	goto L525
L528:
	;
	v2302 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2302 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2305 = int32(1)
	goto L531
L530:
	;
	v2305 = v2302
	goto L531
L531:
	;
	v2306 = v2305
	goto L383
L532:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2315
	v2318 = int32(1)
	goto L104
}
func F_fireRIRonSubLink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(22) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v13 = F_fireRIRrules(m, v11, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)))
				v20 = v18 | v19
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v20)
				v24 = F_expression_tree_walker_impl(m, l0, int32(1120), l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					return v24
				}
			}
		} else {
			v24 = F_expression_tree_walker_impl(m, l0, int32(1120), l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_fixup_selfjoin_quals(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	v4 = int32(0)
	if l1 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v4
	v22 = v4
	goto L7
L5:
	;
	v105 = v4
	goto L6
L6:
	;
	return v105
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = F_pull_varnos(m, l0, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v105 = v94
	goto L6
L9:
	;
	v97 = v22 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v97 < v98 {
		v21 = v94
		v22 = v97
		goto L7
	} else {
		goto L37
	}
L10:
	;
	return int32(0)
L11:
	;
	v33 = F_bms_is_member(m, l2, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v28 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v87 = v28
	goto L15
L15:
	;
	v90 = F_lappend(m, v21, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L36
	}
L16:
	;
	v85 = F_list_member(m, v21, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L34
	}
L17:
	;
	v82 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v38 != int32(17) {
		v82 = v28
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v41 == int32(0) {
		v82 = v28
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 != int32(2) {
		v82 = v28
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = F_equal(m, v48, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v50 == int32(0) {
		v82 = v28
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_set_opfuncid(m, v28)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v57 = F_func_strict(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	if v57 == int32(0) {
		v82 = v28
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v62 = F_exprType(m, v48)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v64 = F_op_mergejoinable(m, v61, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v64 == int32(0) {
		v82 = v28
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v68 = F_contain_volatile_functions(m, v48)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	if v68 != 0 {
		v82 = v28
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v71 = F_palloc0(m, int32(20))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = int32(-1)
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+12)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(52)
	v82 = v71
	goto L16
L34:
	;
	if v85 != 0 {
		v94 = v21
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v87 = v82
	goto L15
L36:
	;
	v94 = v90
	goto L9
L37:
	;
	goto L8
}
func F_flatten_rtes_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v3 = int32(0)
	if l0 == v3 {
		v37 = v3
		return v37
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(67) {
			if v7 != int32(101) {
				v34 = F_expression_tree_walker_impl(m, l0, int32(882), l1)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v37 = v34
					return v37
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				switch v12 {
				case 0:
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
					F_add_rte_to_flat_rtable(m, v16, v18, l0)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				case 1:
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v13 == int32(0) {
						v37 = v3
						return v37
					} else {
						v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
						F_add_rte_to_flat_rtable(m, v16, v18, l0)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				default:
					v37 = v3
					return v37
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
			v29 = F_query_tree_walker_impl(m, l0, int32(882), l1, int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v25
				return v29
			}
		}
	}
}
func F_float48lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float64
	_ = v11
	var v22 int64
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
		v22 = base.I64_extend_i32_u(base.F64_lt(v5, v11) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807))))
	} else {
		v22 = int64(0)
	}
	return v22
}
func F_float48mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v26 float64
	_ = v26
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.F64_promote_f32(v5)
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = base.F64_sub(v6, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v7), v10)|base.F64_eq(base.F64_abs(v6), v10) != 0 {
		v26 = v8
		return base.I64_reinterpret_f64(v26)
	} else {
		v21 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v26 = float64(0)
			return base.I64_reinterpret_f64(v26)
		}
	}
}
func F_float4mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v9 float32
	_ = v9
	var v24 int32
	_ = v24
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = base.F32_sub(v5, v6)
	v9 = math.Float32frombits(uint32(0x7f800000))
	if base.F32_ne(base.F32_abs(v7), v9)|base.F32_eq(base.F32_abs(v5), v9)|base.F32_eq(base.F32_abs(v6), v9) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I64_extend_i32_s(base.I32_reinterpret_f32(v7))
	}
}
func F_float4out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_float4out[0]))
		if int32(0) < v11 {
			v14 = F_float_to_shortest_decimal_bufn(m, v4, v6)
			mBase = m.M
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14+v6))) = uint8(v16)
			return base.I64_extend_i32_u(v6)
		} else {
			F_pg_strfromd(m, v6, v11+int32(6), base.F64_promote_f32(v4))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v6)
			}
		}
	}
}
func F_float4smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = base.I32_wrap_i64(v6)
	if base.Ui32(v7&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v12 = base.I32_wrap_i64(v5)
		if base.Ui32(int32(2139095040)) < base.Ui32(v12&int32(2147483647)) {
			v17 = v6
		} else {
			v17 = v5
		}
		if base.F32_lt(base.F32_reinterpret_i32(v7), base.F32_reinterpret_i32(v12)) != 0 {
			v21 = v6
		} else {
			v21 = v17
		}
		v23 = v21
	} else {
		v23 = v5
	}
	return base.I64_extend32_s(v23)
}
func F_float84mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 float32
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v10 float64
	_ = v10
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v26 float64
	_ = v26
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = base.F64_promote_f32(v6)
	v8 = base.F64_sub(v5, v7)
	v10 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v8), v10)|base.F64_eq(base.F64_abs(v5), v10)|base.F64_eq(base.F64_abs(v7), v10) != 0 {
		v26 = v8
		return base.I64_reinterpret_f64(v26)
	} else {
		v21 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v26 = float64(0)
			return base.I64_reinterpret_f64(v26)
		}
	}
}
func F_float8eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 float64
	_ = v9
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(9223372036854775807)
	v8 = base.I64_reinterpret_f64(v5) & v7
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v9)&v7) {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v8)))
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.Ui64(v8) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v5, v9))
	}
}
func F_float8ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float64
	_ = v10
	var v21 int64
	_ = v21
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
		v21 = base.I64_extend_i32_u(base.F64_ge(v4, v10) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))))
	} else {
		v21 = int64(1)
	}
	return v21
}
func F_float8lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v10 float64
	_ = v10
	var v21 int64
	_ = v21
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
		v21 = base.I64_extend_i32_u(base.F64_lt(v4, v10) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807))))
	} else {
		v21 = int64(0)
	}
	return v21
}
func F_float8mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v22 float64
	_ = v22
	var v25 int32
	_ = v25
	var v26 float64
	_ = v26
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v38 float64
	_ = v38
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = base.F64_mul(v5, v6)
	v9 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v7), v9)|base.F64_eq(base.F64_abs(v5), v9)|base.F64_eq(base.F64_abs(v6), v9) == int32(0) {
		v22 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v38 = float64(0)
			return base.I64_reinterpret_f64(v38)
		}
	} else {
		v26 = float64(0)
		if base.F64_eq(v5, v26)|base.F64_ne(v7, v26)|base.F64_eq(v6, v26) != 0 {
			v38 = v7
			return base.I64_reinterpret_f64(v38)
		} else {
			v35 = F_float_underflow_error_ext(m, int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				v38 = float64(0)
				return base.I64_reinterpret_f64(v38)
			}
		}
	}
}
func F_float8out_internal(m *base.Module, l0 float64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_float8out_internal[0]))
		if int32(0) < v10 {
			F_double_to_shortest_decimal_buf(m, l0, v5)
			mBase = m.M
			return v5
		} else {
			F_pg_strfromd(m, v5, v10+int32(15), l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v5
			}
		}
	}
}
func F_float8send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 float64
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_pq_sendfloat8(m, v6, v8)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v16)
		}
	}
}
func F_fmod(m *base.Module, l0 float64) float64 {
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 float64
	_ = v15
	var v19 int64
	_ = v19
	var v26 float64
	_ = v26
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v136 int64
	_ = v136
	v6 = base.I64_reinterpret_f64(l0)
	v10 = int32(2047)
	v11 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(52))%64))) & v10
	if v11 == v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = base.F64_mul(l0, float64(360))
	return base.F64_div(v15, v15)
L2:
	;
	goto L3
L3:
	;
	v19 = v6 << (uint(int64(1)) % 64)
	if base.Ui64(v19) <= base.Ui64(int64(-9156662467374350336)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v19 == int64(-9156662467374350336) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if v11 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v26 = base.F64_mul(l0, float64(0))
	goto L9
L8:
	;
	v26 = l0
	goto L9
L9:
	;
	return v26
L10:
	;
	if int32(1031) < v63 {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v30 = int32(0)
	v32 = v6 << (uint(int64(12)) % 64)
	if int64(0) <= v32 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v63 = v11
	v64 = v6&int64(4503599627370495) | int64(4503599627370496)
	goto L10
L14:
	;
	v36 = v32
	v39 = v30
	goto L17
L15:
	;
	v50 = v30
	goto L16
L16:
	;
	v63 = v50
	v64 = v6 << (uint(base.I64_extend_i32_u(int32(1)-v50)) % 64)
	goto L10
L17:
	;
	v41 = v39 - int32(1)
	v43 = v36 << (uint(int64(1)) % 64)
	if int64(0) <= v43 {
		v36 = v43
		v39 = v41
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v50 = v41
	goto L16
L19:
	;
	goto L18
L20:
	;
	v68 = v64
	v71 = v63
	goto L23
L21:
	;
	v90 = v64
	v93 = v63
	goto L22
L22:
	;
	v95 = v90 - int64(6333186975989760)
	if v95 < int64(0) {
		v103 = v90
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v73 = v68 - int64(6333186975989760)
	if v73 < int64(0) {
		v81 = v68
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v90 = v83
	v93 = int32(1031)
	goto L22
L25:
	;
	v83 = v81 << (uint(int64(1)) % 64)
	v85 = v71 - int32(1)
	if int32(1031) < v85 {
		v68 = v83
		v71 = v85
		goto L23
	} else {
		goto L28
	}
L26:
	;
	if v73 != int64(0) {
		v81 = v73
		goto L25
	} else {
		goto L27
	}
L27:
	;
	return base.F64_mul(l0, float64(0))
L28:
	;
	goto L24
L29:
	;
	if base.Ui64(v103) <= base.Ui64(int64(4503599627370495)) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	if v95 != int64(0) {
		v103 = v95
		goto L29
	} else {
		goto L31
	}
L31:
	;
	return base.F64_mul(l0, float64(0))
L32:
	;
	v107 = v103
	v110 = v93
	goto L35
L33:
	;
	v118 = v103
	v121 = v93
	goto L34
L34:
	;
	if int32(0) < v121 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v112 = v110 - int32(1)
	v114 = v107 << (uint(int64(1)) % 64)
	if base.Ui64(v107) < base.Ui64(int64(2251799813685248)) {
		v107 = v114
		v110 = v112
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v118 = v114
	v121 = v112
	goto L34
L37:
	;
	goto L36
L38:
	;
	v136 = v118 - int64(4503599627370496) | base.I64_extend_i32_u(v121)<<(uint(int64(52))%64)
	goto L40
L39:
	;
	v136 = int64(base.Ui64(v118) >> (uint(base.I64_extend_i32_u(int32(1)-v121)) % 64))
	goto L40
L40:
	;
	return base.F64_reinterpret_i64(v6&int64(-9223372036854775807-1) | v136)
}
func F_fmt_fp(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v40 float64
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v119 int32
	_ = v119
	var v129 float64
	_ = v129
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 float64
	_ = v168
	var v174 int32
	_ = v174
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 float64
	_ = v202
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v269 int64
	_ = v269
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v324 int32
	_ = v324
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v548 int32
	_ = v548
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v611 int32
	_ = v611
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v744 float64
	_ = v744
	var v751 int32
	_ = v751
	var v758 float64
	_ = v758
	var v763 float64
	_ = v763
	var v766 int32
	_ = v766
	var v768 float64
	_ = v768
	var v770 float64
	_ = v770
	var v771 int32
	_ = v771
	var v776 float64
	_ = v776
	var v777 float64
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v950 int32
	_ = v950
	var v966 int32
	_ = v966
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1071 int32
	_ = v1071
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1142 int32
	_ = v1142
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1177 int64
	_ = v1177
	var v1183 int64
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1190 int64
	_ = v1190
	var v1191 int64
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1201 int64
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1246 int32
	_ = v1246
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1281 int32
	_ = v1281
	var v1304 int32
	_ = v1304
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1336 int32
	_ = v1336
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1397 int64
	_ = v1397
	var v1403 int64
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1411 int64
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1421 int64
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1467 int32
	_ = v1467
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1571 int64
	_ = v1571
	var v1577 int64
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int64
	_ = v1584
	var v1585 int64
	_ = v1585
	var v1591 int32
	_ = v1591
	var v1595 int64
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1640 int32
	_ = v1640
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1675 int32
	_ = v1675
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1748 int64
	_ = v1748
	var v1754 int64
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1761 int64
	_ = v1761
	var v1762 int64
	_ = v1762
	var v1768 int32
	_ = v1768
	var v1772 int64
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1824 int32
	_ = v1824
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1864 int32
	_ = v1864
	var v1871 int32
	_ = v1871
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1914 int32
	_ = v1914
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1948 int32
	_ = v1948
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2024 float64
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2032 float64
	_ = v2032
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2048 float64
	_ = v2048
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2061 float64
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2069 float64
	_ = v2069
	var v2078 float64
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2085 int64
	_ = v2085
	var v2091 int64
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2097 int32
	_ = v2097
	var v2098 int64
	_ = v2098
	var v2099 int64
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2109 int64
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2180 float64
	_ = v2180
	var v2186 int32
	_ = v2186
	var v2207 int32
	_ = v2207
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2216 float64
	_ = v2216
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2294 int32
	_ = v2294
	v8 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(560)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+556)) = v8
	v35 = base.I64_reinterpret_f64(l1)
	if v35 < int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v58 = int64(9218868437227405312)
	if v57&v58 == v58 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v40 = base.F64_neg(l1)
	v53 = v40
	v54 = int32(1)
	v55 = int32(_a_F_fmt_fp_0)
	v56 = v8
	v57 = base.I64_reinterpret_f64(v40)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l4&int32(2048) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v53 = l1
	v54 = int32(1)
	v55 = int32(_a_F_fmt_fp_1)
	v56 = v8
	v57 = v35
	goto L1
L6:
	;
	goto L7
L7:
	;
	v49 = l4 & int32(1)
	if v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = int32(_a_F_fmt_fp_2)
	goto L10
L9:
	;
	v50 = int32(_a_F_fmt_fp_3)
	goto L10
L10:
	;
	v53 = l1
	v54 = v49
	v55 = v50
	v56 = base.B2i32(v49 == int32(0))
	v57 = v35
	goto L1
L11:
	;
	m.G0 = v31 + int32(560)
	return v2294
L12:
	;
	v64 = v54 + int32(3)
	F_pad(m, l0, int32(32), l2, v64, l4&int32(-65537))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v94 = v31 + int32(528)
	v96 = v31 + int32(556)
	v99 = base.I64_reinterpret_f64(v53)
	v103 = int32(2047)
	v104 = base.I32_wrap_i64(int64(base.Ui64(v99)>>(uint(int64(52))%64))) & v103
	if v104 != v103 {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	return int32(0)
L16:
	;
	F_out(m, l0, v55, v54)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v76 = l5 & int32(32)
	if v76 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v77 = int32(_a_F_fmt_fp_4)
	goto L20
L19:
	;
	v77 = int32(_a_F_fmt_fp_5)
	goto L20
L20:
	;
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = int32(_a_F_fmt_fp_6)
	goto L23
L22:
	;
	v80 = int32(_a_F_fmt_fp_7)
	goto L23
L23:
	;
	if base.F64_ne(v53, v53) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v82 = v77
	goto L26
L25:
	;
	v82 = v80
	goto L26
L26:
	;
	F_out(m, l0, v82, int32(3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_pad(m, l0, int32(32), l2, v64, l4^int32(_a_F_fmt_fp_8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	if v64 < l2 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = l2
	goto L31
L30:
	;
	v92 = v64
	goto L31
L31:
	;
	v2294 = v92
	goto L11
L32:
	;
	v2020 = v55 + l5<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(9)
	if base.Ui32(int32(12)) < base.Ui32(l3) {
		v2078 = v132
		goto L375
	} else {
		goto L376
	}
L33:
	;
	if l3 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L34:
	;
	v149 = v135 - int32(29)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+556)) = v149
	v153 = base.F64_mul(v132, float64(2.68435456e+08))
	v155 = v149
	v156 = v140
	goto L33
L35:
	;
	v132 = base.F64_add(v131, v131)
	if base.F64_ne(v132, float64(0)) != 0 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	if v104 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v129 = v53
	goto L38
L38:
	;
	v131 = v129
	goto L35
L39:
	;
	if base.F64_eq(v53, float64(0)) != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v104 - int32(1022)
	v129 = base.F64_reinterpret_i64(v99&int64(-9218868437227405313) | int64(4602678819172646912))
	goto L38
L42:
	;
	v118 = v53
	v119 = int32(0)
	goto L44
L43:
	;
	v114 = F_frexp(m, base.F64_mul(v53, float64(1.8446744073709552e+19)), v96)
	mBase = m.M
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v118 = v114
	v119 = v115 + int32(-64)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v119
	v131 = v118
	goto L35
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+556)) = v135 - int32(1)
	v140 = l5 | int32(32)
	if v140 != int32(97) {
		goto L34
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v144 = l5 | int32(32)
	if v144 == int32(97) {
		goto L32
	} else {
		goto L49
	}
L48:
	;
	goto L32
L49:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	v153 = v132
	v155 = v147
	v156 = v144
	goto L33
L50:
	;
	v160 = int32(6)
	goto L52
L51:
	;
	v160 = l3
	goto L52
L52:
	;
	v162 = int32(0)
	if v162 <= v155 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v165 = int32(488)
	goto L55
L54:
	;
	v165 = v162
	goto L55
L55:
	;
	v166 = v31 + v165
	v168 = v153
	v174 = v166
	goto L56
L56:
	;
	v195 = base.I32_trunc_sat_f64_u(v168)
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v195
	v198 = v174 + int32(4)
	v202 = base.F64_mul(base.F64_sub(v168, base.F64_convert_i32_u(v195)), float64(1e+09))
	if base.F64_ne(v202, float64(0)) != 0 {
		v168 = v202
		v174 = v198
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v155 <= int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L57
L59:
	;
	if v366 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L60:
	;
	v363 = v198
	v365 = v166
	v366 = v155
	goto L59
L61:
	;
	goto L62
L62:
	;
	v214 = v198
	v215 = v166
	v216 = v155
	goto L63
L63:
	;
	v235 = int32(29)
	if base.Ui32(v235) <= base.Ui32(v216) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v363 = v324
	v365 = v297
	v366 = v353
	goto L59
L65:
	;
	v238 = v235
	goto L67
L66:
	;
	v238 = v216
	goto L67
L67:
	;
	v240 = v214 - int32(4)
	if base.Ui32(v240) < base.Ui32(v215) {
		v297 = v215
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v324 = v214
	goto L74
L69:
	;
	v250 = v240
	v269 = int64(0)
	goto L70
L70:
	;
	v272 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v250))))
	v274 = v272<<(uint(base.I64_extend_i32_u(v238))%64) + v269
	v275 = int64(1000000000)
	v276 = base.I64_div_u_s(v274, v275)
	v279 = v274 - v276*v275
	*(*uint32)(unsafe.Add(mBase, uint32(v250))) = uint32(v279)
	v282 = v250 - int32(4)
	if base.Ui32(v215) <= base.Ui32(v282) {
		v250 = v282
		v269 = v276
		goto L70
	} else {
		goto L72
	}
L71:
	;
	if base.Ui64(v274) < base.Ui64(int64(1000000000)) {
		v297 = v215
		goto L68
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	v287 = v215 - int32(4)
	*(*uint32)(unsafe.Add(mBase, uint32(v287))) = uint32(v276)
	v297 = v287
	goto L68
L74:
	;
	if base.Ui32(v297) < base.Ui32(v324) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	v353 = v352 - v238
	*(*int32)(unsafe.Add(mBase, uint32(v31)+556)) = v353
	if int32(0) < v353 {
		v214 = v324
		v215 = v297
		v216 = v353
		goto L63
	} else {
		goto L80
	}
L76:
	;
	v347 = v324 - int32(4)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	if v348 == int32(0) {
		v324 = v347
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L78
L80:
	;
	goto L64
L81:
	;
	v390 = base.I32_div_u_s(v160+int32(25), int32(9))
	v392 = v390 + int32(1)
	v401 = v363
	v403 = v365
	v404 = v366
	goto L84
L82:
	;
	v536 = v363
	v538 = v365
	v548 = v8
	goto L83
L83:
	;
	if base.Ui32(v536) <= base.Ui32(v538) {
		v611 = int32(0)
		goto L110
	} else {
		goto L111
	}
L84:
	;
	v423 = int32(9)
	v425 = int32(0) - v404
	if base.Ui32(v423) <= base.Ui32(v425) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v536 = v527
	v538 = v518
	v548 = v392
	goto L83
L86:
	;
	v428 = v423
	goto L88
L87:
	;
	v428 = v425
	goto L88
L88:
	;
	if base.Ui32(v401) <= base.Ui32(v403) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	v516 = v515 + v428
	*(*int32)(unsafe.Add(mBase, uint32(v31)+556)) = v516
	v518 = v494 + v403
	if v156 == int32(102) {
		goto L103
	} else {
		goto L104
	}
L90:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v432 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v436 = int32(-1)
	v448 = v403
	v450 = int32(0)
	goto L96
L93:
	;
	v433 = int32(0)
	goto L95
L94:
	;
	v433 = int32(4)
	goto L95
L95:
	;
	v493 = v401
	v494 = v433
	goto L89
L96:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = int32(base.Ui32(v469)>>(uint(v428)%32)) + v450
	v474 = v469 & (v436<<(uint(v428)%32) ^ v436) * int32(base.Ui32(int32(1000000000))>>(uint(v428)%32))
	v476 = v448 + int32(4)
	if base.Ui32(v476) < base.Ui32(v401) {
		v448 = v476
		v450 = v474
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v480 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L97
L99:
	;
	v481 = int32(0)
	goto L101
L100:
	;
	v481 = int32(4)
	goto L101
L101:
	;
	if v474 == int32(0) {
		v493 = v401
		v494 = v481
		goto L89
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = v474
	v493 = v401 + int32(4)
	v494 = v481
	goto L89
L103:
	;
	v519 = v166
	goto L105
L104:
	;
	v519 = v518
	goto L105
L105:
	;
	v520 = int32(2)
	if v392 < (v493-v519)>>(uint(v520)%32) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v527 = v519 + v392<<(uint(v520)%32)
	goto L108
L107:
	;
	v527 = v493
	goto L108
L108:
	;
	if v516 < int32(0) {
		v401 = v527
		v403 = v518
		v404 = v516
		goto L84
	} else {
		goto L109
	}
L109:
	;
	goto L85
L110:
	;
	if v156 != int32(102) {
		goto L116
	} else {
		goto L117
	}
L111:
	;
	v564 = (v166 - v538) >> (uint(int32(2)) % 32) * int32(9)
	v565 = int32(10)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if base.Ui32(v566) < base.Ui32(v565) {
		v611 = v564
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v576 = v565
	v578 = v564
	goto L113
L113:
	;
	v598 = v578 + int32(1)
	v600 = v576 * int32(10)
	if base.Ui32(v600) <= base.Ui32(v566) {
		v576 = v600
		v578 = v598
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v611 = v598
	goto L110
L115:
	;
	goto L114
L116:
	;
	v633 = v611
	goto L118
L117:
	;
	v633 = int32(0)
	goto L118
L118:
	;
	v640 = v160 - v633 - base.B2i32(v156 == int32(103))&base.B2i32(v160 != int32(0))
	v644 = int32(9)
	if v640 < (v536-v166)>>(uint(int32(2))%32)*v644-v644 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if v155 < int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v938 = v536
	v940 = v538
	v941 = v611
	v950 = v548
	goto L121
L121:
	;
	v966 = v938
	goto L168
L122:
	;
	v653 = int32(-4092)
	goto L124
L123:
	;
	v653 = int32(-3604)
	goto L124
L124:
	;
	v656 = v640 + int32(_a_F_fmt_fp_9)
	v657 = int32(9)
	v658 = base.I32_div_s(v656, v657)
	v661 = v31 + v653 + v658<<(uint(int32(2))%32)
	v662 = int32(10)
	v665 = v656 - v658*v657
	if v665 <= int32(7) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v671 = v665
	v675 = v662
	goto L128
L126:
	;
	v709 = v662
	goto L127
L127:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	v731 = base.I32_div_u_s(v730, v709)
	v733 = v730 - v731*v709
	v737 = v661 + int32(4)
	if base.B2i32(v733 == int32(0))&base.B2i32(v737 == v536) != 0 {
		v908 = v538
		v909 = v611
		v913 = v661
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v697 = v675 * int32(10)
	v699 = v671 + int32(1)
	if v699 != int32(8) {
		v671 = v699
		v675 = v697
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v709 = v697
	goto L127
L130:
	;
	goto L129
L131:
	;
	v929 = v913 + int32(4)
	if base.Ui32(v929) < base.Ui32(v536) {
		goto L165
	} else {
		goto L166
	}
L132:
	;
	if v731&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v737 == v536 {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v744 = float64(9.007199254740992e+15)
	if base.B2i32(v709 != int32(1000000000))|base.B2i32(base.Ui32(v661) <= base.Ui32(v538)) != 0 {
		v758 = v744
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v758 = float64(9.007199254740994e+15)
	goto L133
L137:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661-int32(4)))))
	if v751&int32(1) == int32(0) {
		v758 = v744
		goto L133
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v763 = float64(1)
	goto L141
L140:
	;
	v763 = float64(1.5)
	goto L141
L141:
	;
	v766 = int32(base.Ui32(v709) >> (uint(int32(1)) % 32))
	if v733 == v766 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v768 = v763
	goto L144
L143:
	;
	v768 = float64(1.5)
	goto L144
L144:
	;
	if base.Ui32(v733) < base.Ui32(v766) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v770 = float64(0.5)
	goto L147
L146:
	;
	v770 = v768
	goto L147
L147:
	;
	if v56 != 0 {
		v776 = v758
		v777 = v770
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v778 = v730 - v733
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v778
	if base.F64_eq(base.F64_add(v776, v777), v776) != 0 {
		v908 = v538
		v909 = v611
		v913 = v661
		goto L131
	} else {
		goto L151
	}
L149:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v771 != int32(45) {
		v776 = v758
		v777 = v770
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v776 = base.F64_neg(v758)
	v777 = base.F64_neg(v770)
	goto L148
L151:
	;
	v782 = v778 + v709
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v782
	if base.Ui32(int32(1000000000)) <= base.Ui32(v782) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v794 = v538
	v799 = v661
	goto L155
L153:
	;
	v838 = v538
	v843 = v661
	goto L154
L154:
	;
	v862 = (v166 - v838) >> (uint(int32(2)) % 32) * int32(9)
	v863 = int32(10)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	if base.Ui32(v864) < base.Ui32(v863) {
		v908 = v838
		v909 = v862
		v913 = v843
		goto L131
	} else {
		goto L161
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v799))) = int32(0)
	v817 = v799 - int32(4)
	if base.Ui32(v817) < base.Ui32(v794) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v838 = v823
	v843 = v817
	goto L154
L157:
	;
	v820 = v794 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v820))) = int32(0)
	v823 = v820
	goto L159
L158:
	;
	v823 = v794
	goto L159
L159:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v826 = v824 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v817))) = v826
	if base.Ui32(int32(999999999)) < base.Ui32(v826) {
		v794 = v823
		v799 = v817
		goto L155
	} else {
		goto L160
	}
L160:
	;
	goto L156
L161:
	;
	v874 = v863
	v876 = v862
	goto L162
L162:
	;
	v896 = v876 + int32(1)
	v898 = v874 * int32(10)
	if base.Ui32(v898) <= base.Ui32(v864) {
		v874 = v898
		v876 = v896
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v908 = v838
	v909 = v896
	v913 = v843
	goto L131
L164:
	;
	goto L163
L165:
	;
	v931 = v929
	goto L167
L166:
	;
	v931 = v536
	goto L167
L167:
	;
	v938 = v931
	v940 = v908
	v941 = v909
	v950 = v731
	goto L121
L168:
	;
	v988 = base.B2i32(base.Ui32(v966) <= base.Ui32(v940))
	if v988 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v156 != int32(103) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	v992 = v966 - int32(4)
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	if v993 == int32(0) {
		v966 = v992
		goto L168
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	goto L169
L173:
	;
	goto L172
L174:
	;
	v1151 = int32(-1)
	v1154 = v1133 | v1142
	if v1154 != 0 {
		goto L210
	} else {
		goto L211
	}
L175:
	;
	v1128 = l5
	v1133 = v160
	v1142 = l4 & int32(8)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1001 = int32(-1)
	if v160 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1005 = v160
	goto L180
L179:
	;
	v1005 = int32(1)
	goto L180
L180:
	;
	v1009 = base.B2i32(v941 < v1005) & base.B2i32(int32(-5) < v941)
	if v1009 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1010 = v941 ^ v1001
	goto L183
L182:
	;
	v1010 = v1001
	goto L183
L183:
	;
	v1011 = v1010 + v1005
	if v1009 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1014 = int32(-1)
	goto L186
L185:
	;
	v1014 = int32(-2)
	goto L186
L186:
	;
	v1015 = v1014 + l5
	v1017 = l4 & int32(8)
	if v1017 != 0 {
		v1128 = v1015
		v1133 = v1011
		v1142 = v1017
		goto L174
	} else {
		goto L187
	}
L187:
	;
	v1018 = int32(-9)
	if base.Ui32(v966) <= base.Ui32(v940) {
		v1071 = v1018
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1097 = (v966 - v166) >> (uint(int32(2)) % 32) * int32(9)
	if v1015&int32(-33) == int32(70) {
		goto L195
	} else {
		goto L196
	}
L189:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v966-int32(4))))
	if v1021 == int32(0) {
		v1071 = v1018
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1024 = int32(10)
	v1025 = int32(0)
	v1027 = base.I32_rem_u_s(v1021, v1024)
	if v1027 != 0 {
		v1071 = v1025
		goto L188
	} else {
		goto L191
	}
L191:
	;
	v1031 = v1024
	v1034 = v1025
	goto L192
L192:
	;
	v1059 = v1031 * int32(10)
	v1060 = base.I32_rem_u_s(v1021, v1059)
	if v1060 == int32(0) {
		v1031 = v1059
		v1034 = v1034 + int32(1)
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v1071 = v1034 ^ int32(-1)
	goto L188
L194:
	;
	goto L193
L195:
	;
	v1102 = int32(0)
	v1105 = v1097 + v1071 - int32(9)
	if v1102 < v1105 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	v1112 = int32(0)
	v1116 = v1097 + v941 + v1071 - int32(9)
	if v1112 < v1116 {
		goto L204
	} else {
		goto L205
	}
L198:
	;
	v1109 = v1105
	goto L200
L199:
	;
	v1109 = v1102
	goto L200
L200:
	;
	if v1011 < v1109 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1111 = v1011
	goto L203
L202:
	;
	v1111 = v1109
	goto L203
L203:
	;
	v1128 = v1015
	v1133 = v1111
	v1142 = v1102
	goto L174
L204:
	;
	v1120 = v1116
	goto L206
L205:
	;
	v1120 = v1112
	goto L206
L206:
	;
	if v1011 < v1120 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1122 = v1011
	goto L209
L208:
	;
	v1122 = v1120
	goto L209
L209:
	;
	v1128 = v1015
	v1133 = v1122
	v1142 = v1112
	goto L174
L210:
	;
	v1155 = int32(2147483645)
	goto L212
L211:
	;
	v1155 = int32(2147483646)
	goto L212
L212:
	;
	if v1155 < v1133 {
		v2294 = v1151
		goto L11
	} else {
		goto L213
	}
L213:
	;
	v1161 = v1133 + base.B2i32(v1154 != int32(0)) + int32(1)
	v1163 = v1128 & int32(-33)
	if v1163 == int32(70) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v1346 = v1324 + v1161
	if v54^int32(2147483647) < v1346 {
		v2294 = v1151
		goto L11
	} else {
		goto L248
	}
L215:
	;
	if v1161^int32(2147483647) < v941 {
		v2294 = v1151
		goto L11
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v1174 = v941 >> (uint(int32(31)) % 32)
	v1177 = base.I64_extend_i32_u(v941 ^ v1174 - v1174)
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1177) {
		goto L223
	} else {
		goto L224
	}
L218:
	;
	v1169 = int32(0)
	if v1169 < v941 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1172 = v941
	goto L221
L220:
	;
	v1172 = v1169
	goto L221
L221:
	;
	v1324 = v1172
	v1336 = v950
	goto L214
L222:
	;
	if v94-v1236 <= int32(1) {
		goto L238
	} else {
		goto L239
	}
L223:
	;
	v1183 = v1177
	v1184 = v94
	goto L226
L224:
	;
	v1201 = v1177
	v1202 = v94
	goto L225
L225:
	;
	v1206 = base.I32_wrap_i64(v1201)
	if base.Ui64(int64(10)) <= base.Ui64(v1201) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v1189 = v1184 - int32(1)
	v1190 = int64(10)
	v1191 = base.I64_div_u_s(v1183, v1190)
	v1197 = base.I32_wrap_i64(v1183-v1191*v1190) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1189))) = uint8(v1197)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1183) {
		v1183 = v1191
		v1184 = v1189
		goto L226
	} else {
		goto L228
	}
L227:
	;
	v1201 = v1191
	v1202 = v1189
	goto L225
L228:
	;
	goto L227
L229:
	;
	v1210 = v1202
	v1211 = v1206
	goto L232
L230:
	;
	v1227 = v1202
	v1228 = v1206
	goto L231
L231:
	;
	if v1228 != 0 {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v1215 = v1210 - int32(1)
	v1216 = int32(10)
	v1217 = base.I32_div_u_s(v1211, v1216)
	v1222 = v1211 - v1217*v1216 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1215))) = uint8(v1222)
	if base.Ui32(int32(99)) < base.Ui32(v1211) {
		v1210 = v1215
		v1211 = v1217
		goto L232
	} else {
		goto L234
	}
L233:
	;
	v1227 = v1215
	v1228 = v1217
	goto L231
L234:
	;
	goto L233
L235:
	;
	v1232 = v1227 - int32(1)
	v1234 = v1228 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1232))) = uint8(v1234)
	v1236 = v1232
	goto L237
L236:
	;
	v1236 = v1227
	goto L237
L237:
	;
	goto L222
L238:
	;
	v1246 = v1236
	goto L241
L239:
	;
	v1281 = v1236
	goto L240
L240:
	;
	v1304 = v1281 - int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v1304))) = uint8(v1128)
	if v941 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v1269 = v1246 - int32(1)
	v1270 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1269))) = uint8(v1270)
	if v94-v1269 < int32(2) {
		v1246 = v1269
		goto L241
	} else {
		goto L243
	}
L242:
	;
	v1281 = v1269
	goto L240
L243:
	;
	goto L242
L244:
	;
	v1312 = int32(45)
	goto L246
L245:
	;
	v1312 = int32(43)
	goto L246
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1281-int32(1)))) = uint8(v1312)
	v1314 = v94 - v1304
	if v1161^int32(2147483647) < v1314 {
		v2294 = v1151
		goto L11
	} else {
		goto L247
	}
L247:
	;
	v1324 = v1314
	v1336 = v1304
	goto L214
L248:
	;
	v1351 = v1346 + v54
	F_pad(m, l0, int32(32), l2, v1351, l4)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L15
	} else {
		goto L249
	}
L249:
	;
	F_out(m, l0, v55, v54)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L15
	} else {
		goto L250
	}
L250:
	;
	F_pad(m, l0, int32(48), l2, v1351, l4^int32(_a_F_fmt_fp_10))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L15
	} else {
		goto L251
	}
L251:
	;
	if v1163 == int32(70) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	F_pad(m, l0, int32(32), l2, v1351, l4^int32(_a_F_fmt_fp_8))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L15
	} else {
		goto L371
	}
L253:
	;
	v1971 = int32(9)
	F_pad(m, l0, int32(48), v1948+v1971, v1971, int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L15
	} else {
		goto L370
	}
L254:
	;
	v1948 = v1133
	goto L253
L255:
	;
	v1366 = v31 + int32(528) | int32(9)
	if base.Ui32(v166) < base.Ui32(v940) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	goto L257
L257:
	;
	if v1133 < int32(0) {
		v1914 = v1133
		goto L325
	} else {
		goto L326
	}
L258:
	;
	v1368 = v166
	goto L260
L259:
	;
	v1368 = v940
	goto L260
L260:
	;
	v1377 = v1368
	goto L261
L261:
	;
	v1397 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1377))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1397) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	if v1154 != 0 {
		goto L290
	} else {
		goto L291
	}
L263:
	;
	if v1368 != v1377 {
		goto L280
	} else {
		goto L281
	}
L264:
	;
	v1403 = v1397
	v1404 = v1366
	goto L267
L265:
	;
	v1421 = v1397
	v1422 = v1366
	goto L266
L266:
	;
	v1426 = base.I32_wrap_i64(v1421)
	if base.Ui64(int64(10)) <= base.Ui64(v1421) {
		goto L270
	} else {
		goto L271
	}
L267:
	;
	v1409 = v1404 - int32(1)
	v1410 = int64(10)
	v1411 = base.I64_div_u_s(v1403, v1410)
	v1417 = base.I32_wrap_i64(v1403-v1411*v1410) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1409))) = uint8(v1417)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1403) {
		v1403 = v1411
		v1404 = v1409
		goto L267
	} else {
		goto L269
	}
L268:
	;
	v1421 = v1411
	v1422 = v1409
	goto L266
L269:
	;
	goto L268
L270:
	;
	v1430 = v1422
	v1431 = v1426
	goto L273
L271:
	;
	v1447 = v1422
	v1448 = v1426
	goto L272
L272:
	;
	if v1448 != 0 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1435 = v1430 - int32(1)
	v1436 = int32(10)
	v1437 = base.I32_div_u_s(v1431, v1436)
	v1442 = v1431 - v1437*v1436 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1435))) = uint8(v1442)
	if base.Ui32(int32(99)) < base.Ui32(v1431) {
		v1430 = v1435
		v1431 = v1437
		goto L273
	} else {
		goto L275
	}
L274:
	;
	v1447 = v1435
	v1448 = v1437
	goto L272
L275:
	;
	goto L274
L276:
	;
	v1452 = v1447 - int32(1)
	v1454 = v1448 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1452))) = uint8(v1454)
	v1456 = v1452
	goto L278
L277:
	;
	v1456 = v1447
	goto L278
L278:
	;
	goto L263
L279:
	;
	F_out(m, l0, v1507, v1366-v1507)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L15
	} else {
		goto L288
	}
L280:
	;
	if base.Ui32(v1456) <= base.Ui32(v31+int32(528)) {
		v1507 = v1456
		goto L279
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	if v1366 != v1456 {
		v1507 = v1456
		goto L279
	} else {
		goto L287
	}
L283:
	;
	v1467 = v1456
	goto L284
L284:
	;
	v1490 = v1467 - int32(1)
	v1491 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1490))) = uint8(v1491)
	if base.Ui32(v31+int32(528)) < base.Ui32(v1490) {
		v1467 = v1490
		goto L284
	} else {
		goto L286
	}
L285:
	;
	v1507 = v1490
	goto L279
L286:
	;
	goto L285
L287:
	;
	v1498 = v1456 - int32(1)
	v1499 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1498))) = uint8(v1499)
	v1507 = v1498
	goto L279
L288:
	;
	v1533 = v1377 + int32(4)
	if base.Ui32(v1533) <= base.Ui32(v166) {
		v1377 = v1533
		goto L261
	} else {
		goto L289
	}
L289:
	;
	goto L262
L290:
	;
	F_out(m, l0, int32(_a_F_fmt_fp_11), int32(1))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L15
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	if base.B2i32(v1133 <= int32(0))|base.B2i32(base.Ui32(v966) <= base.Ui32(v1533)) != 0 {
		goto L254
	} else {
		goto L294
	}
L293:
	;
	goto L292
L294:
	;
	v1551 = v1533
	v1553 = v1133
	goto L295
L295:
	;
	v1571 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1551))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1571) {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	v1948 = v1704
	goto L253
L297:
	;
	if base.Ui32(v31+int32(528)) < base.Ui32(v1630) {
		goto L313
	} else {
		goto L314
	}
L298:
	;
	v1577 = v1571
	v1578 = v1366
	goto L301
L299:
	;
	v1595 = v1571
	v1596 = v1366
	goto L300
L300:
	;
	v1600 = base.I32_wrap_i64(v1595)
	if base.Ui64(int64(10)) <= base.Ui64(v1595) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v1583 = v1578 - int32(1)
	v1584 = int64(10)
	v1585 = base.I64_div_u_s(v1577, v1584)
	v1591 = base.I32_wrap_i64(v1577-v1585*v1584) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1583))) = uint8(v1591)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1577) {
		v1577 = v1585
		v1578 = v1583
		goto L301
	} else {
		goto L303
	}
L302:
	;
	v1595 = v1585
	v1596 = v1583
	goto L300
L303:
	;
	goto L302
L304:
	;
	v1604 = v1596
	v1605 = v1600
	goto L307
L305:
	;
	v1621 = v1596
	v1622 = v1600
	goto L306
L306:
	;
	if v1622 != 0 {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	v1609 = v1604 - int32(1)
	v1610 = int32(10)
	v1611 = base.I32_div_u_s(v1605, v1610)
	v1616 = v1605 - v1611*v1610 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1609))) = uint8(v1616)
	if base.Ui32(int32(99)) < base.Ui32(v1605) {
		v1604 = v1609
		v1605 = v1611
		goto L307
	} else {
		goto L309
	}
L308:
	;
	v1621 = v1609
	v1622 = v1611
	goto L306
L309:
	;
	goto L308
L310:
	;
	v1626 = v1621 - int32(1)
	v1628 = v1622 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1626))) = uint8(v1628)
	v1630 = v1626
	goto L312
L311:
	;
	v1630 = v1621
	goto L312
L312:
	;
	goto L297
L313:
	;
	v1640 = v1630
	goto L316
L314:
	;
	v1675 = v1630
	goto L315
L315:
	;
	v1697 = int32(9)
	if v1697 <= v1553 {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	v1663 = v1640 - int32(1)
	v1664 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1663))) = uint8(v1664)
	if base.Ui32(v31+int32(528)) < base.Ui32(v1663) {
		v1640 = v1663
		goto L316
	} else {
		goto L318
	}
L317:
	;
	v1675 = v1663
	goto L315
L318:
	;
	goto L317
L319:
	;
	v1700 = v1697
	goto L321
L320:
	;
	v1700 = v1553
	goto L321
L321:
	;
	F_out(m, l0, v1675, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L15
	} else {
		goto L322
	}
L322:
	;
	v1704 = v1553 - int32(9)
	v1706 = v1551 + int32(4)
	if base.Ui32(v966) <= base.Ui32(v1706) {
		v1948 = v1704
		goto L253
	} else {
		goto L323
	}
L323:
	;
	if int32(9) < v1553 {
		v1551 = v1706
		v1553 = v1704
		goto L295
	} else {
		goto L324
	}
L324:
	;
	goto L296
L325:
	;
	v1933 = int32(18)
	F_pad(m, l0, int32(48), v1914+v1933, v1933, int32(0))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L15
	} else {
		goto L368
	}
L326:
	;
	if base.Ui32(v940) < base.Ui32(v966) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1715 = v966
	goto L329
L328:
	;
	v1715 = v940 + int32(4)
	goto L329
L329:
	;
	v1719 = v31 + int32(528) | int32(9)
	v1727 = v940
	v1730 = v1133
	goto L330
L330:
	;
	v1748 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1727))))
	if base.Ui64(int64(4294967296)) <= base.Ui64(v1748) {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	v1914 = v1898
	goto L325
L332:
	;
	if v1719 == v1807 {
		goto L348
	} else {
		goto L349
	}
L333:
	;
	v1754 = v1748
	v1755 = v1719
	goto L336
L334:
	;
	v1772 = v1748
	v1773 = v1719
	goto L335
L335:
	;
	v1777 = base.I32_wrap_i64(v1772)
	if base.Ui64(int64(10)) <= base.Ui64(v1772) {
		goto L339
	} else {
		goto L340
	}
L336:
	;
	v1760 = v1755 - int32(1)
	v1761 = int64(10)
	v1762 = base.I64_div_u_s(v1754, v1761)
	v1768 = base.I32_wrap_i64(v1754-v1762*v1761) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1760))) = uint8(v1768)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1754) {
		v1754 = v1762
		v1755 = v1760
		goto L336
	} else {
		goto L338
	}
L337:
	;
	v1772 = v1762
	v1773 = v1760
	goto L335
L338:
	;
	goto L337
L339:
	;
	v1781 = v1773
	v1782 = v1777
	goto L342
L340:
	;
	v1798 = v1773
	v1799 = v1777
	goto L341
L341:
	;
	if v1799 != 0 {
		goto L345
	} else {
		goto L346
	}
L342:
	;
	v1786 = v1781 - int32(1)
	v1787 = int32(10)
	v1788 = base.I32_div_u_s(v1782, v1787)
	v1793 = v1782 - v1788*v1787 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1786))) = uint8(v1793)
	if base.Ui32(int32(99)) < base.Ui32(v1782) {
		v1781 = v1786
		v1782 = v1788
		goto L342
	} else {
		goto L344
	}
L343:
	;
	v1798 = v1786
	v1799 = v1788
	goto L341
L344:
	;
	goto L343
L345:
	;
	v1803 = v1798 - int32(1)
	v1805 = v1799 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1803))) = uint8(v1805)
	v1807 = v1803
	goto L347
L346:
	;
	v1807 = v1798
	goto L347
L347:
	;
	goto L332
L348:
	;
	v1810 = v1807 - int32(1)
	v1811 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1810))) = uint8(v1811)
	v1813 = v1810
	goto L350
L349:
	;
	v1813 = v1807
	goto L350
L350:
	;
	if v1727 != v940 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1893 = v1719 - v1871
	if v1893 < v1730 {
		goto L362
	} else {
		goto L363
	}
L352:
	;
	if base.Ui32(v1813) <= base.Ui32(v31+int32(528)) {
		v1871 = v1813
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	F_out(m, l0, v1813, int32(1))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L15
	} else {
		goto L359
	}
L355:
	;
	v1824 = v1813
	goto L356
L356:
	;
	v1847 = v1824 - int32(1)
	v1848 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1847))) = uint8(v1848)
	if base.Ui32(v31+int32(528)) < base.Ui32(v1847) {
		v1824 = v1847
		goto L356
	} else {
		goto L358
	}
L357:
	;
	v1871 = v1847
	goto L351
L358:
	;
	goto L357
L359:
	;
	v1857 = v1813 + int32(1)
	if v1730|v1142 == int32(0) {
		v1871 = v1857
		goto L351
	} else {
		goto L360
	}
L360:
	;
	F_out(m, l0, int32(_a_F_fmt_fp_11), int32(1))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L15
	} else {
		goto L361
	}
L361:
	;
	v1871 = v1857
	goto L351
L362:
	;
	v1895 = v1893
	goto L364
L363:
	;
	v1895 = v1730
	goto L364
L364:
	;
	F_out(m, l0, v1871, v1895)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L15
	} else {
		goto L365
	}
L365:
	;
	v1898 = v1730 - v1893
	v1900 = v1727 + int32(4)
	if base.Ui32(v1715) <= base.Ui32(v1900) {
		v1914 = v1898
		goto L325
	} else {
		goto L366
	}
L366:
	;
	if int32(0) <= v1898 {
		v1727 = v1900
		v1730 = v1898
		goto L330
	} else {
		goto L367
	}
L367:
	;
	goto L331
L368:
	;
	F_out(m, l0, v1336, v94-v1336)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L15
	} else {
		goto L369
	}
L369:
	;
	goto L252
L370:
	;
	goto L252
L371:
	;
	if v1351 < l2 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v2011 = l2
	goto L374
L373:
	;
	v2011 = v1351
	goto L374
L374:
	;
	v2294 = v2011
	goto L11
L375:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	v2082 = v2080 >> (uint(int32(31)) % 32)
	v2085 = base.I64_extend_i32_u(v2080 ^ v2082 - v2082)
	if base.Ui64(int64(4294967296)) <= base.Ui64(v2085) {
		goto L399
	} else {
		goto L400
	}
L376:
	;
	v2023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2020))))
	v2024 = float64(1)
	v2028 = int32(52) - l3<<(uint(int32(2))%32)
	if int32(1024) <= v2028 {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	if v2023 == int32(45) {
		goto L395
	} else {
		goto L396
	}
L378:
	;
	v2069 = base.F64_mul(v2061, base.F64_reinterpret_i64(base.I64_extend_i32_u(v2062+int32(1023))<<(uint(int64(52))%64)))
	goto L377
L379:
	;
	v2032 = base.F64_mul(v2024, float64(8.98846567431158e+307))
	if base.Ui32(v2028) < base.Ui32(int32(2047)) {
		goto L382
	} else {
		goto L383
	}
L380:
	;
	goto L381
L381:
	;
	if int32(-1023) < v2028 {
		v2061 = v2024
		v2062 = v2028
		goto L378
	} else {
		goto L388
	}
L382:
	;
	v2061 = v2032
	v2062 = v2028 - int32(1023)
	goto L378
L383:
	;
	goto L384
L384:
	;
	v2039 = int32(3069)
	if base.Ui32(v2039) <= base.Ui32(v2028) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2042 = v2039
	goto L387
L386:
	;
	v2042 = v2028
	goto L387
L387:
	;
	v2061 = base.F64_mul(v2032, float64(8.98846567431158e+307))
	v2062 = v2042 - int32(2046)
	goto L378
L388:
	;
	v2048 = base.F64_mul(v2024, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v2028) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v2061 = v2048
	v2062 = v2028 + int32(969)
	goto L378
L390:
	;
	goto L391
L391:
	;
	v2055 = int32(-2960)
	if base.Ui32(v2028) <= base.Ui32(v2055) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2058 = v2055
	goto L394
L393:
	;
	v2058 = v2028
	goto L394
L394:
	;
	v2061 = base.F64_mul(v2048, float64(2.004168360008973e-292))
	v2062 = v2058 + int32(1938)
	goto L378
L395:
	;
	v2078 = base.F64_neg(base.F64_add(v2069, base.F64_sub(base.F64_neg(v132), v2069)))
	goto L375
L396:
	;
	goto L397
L397:
	;
	v2078 = base.F64_sub(base.F64_add(v132, v2069), v2069)
	goto L375
L398:
	;
	if v94 == v2144 {
		goto L414
	} else {
		goto L415
	}
L399:
	;
	v2091 = v2085
	v2092 = v94
	goto L402
L400:
	;
	v2109 = v2085
	v2110 = v94
	goto L401
L401:
	;
	v2114 = base.I32_wrap_i64(v2109)
	if base.Ui64(int64(10)) <= base.Ui64(v2109) {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v2097 = v2092 - int32(1)
	v2098 = int64(10)
	v2099 = base.I64_div_u_s(v2091, v2098)
	v2105 = base.I32_wrap_i64(v2091-v2099*v2098) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2097))) = uint8(v2105)
	if base.Ui64(int64(42949672959)) < base.Ui64(v2091) {
		v2091 = v2099
		v2092 = v2097
		goto L402
	} else {
		goto L404
	}
L403:
	;
	v2109 = v2099
	v2110 = v2097
	goto L401
L404:
	;
	goto L403
L405:
	;
	v2118 = v2110
	v2119 = v2114
	goto L408
L406:
	;
	v2135 = v2110
	v2136 = v2114
	goto L407
L407:
	;
	if v2136 != 0 {
		goto L411
	} else {
		goto L412
	}
L408:
	;
	v2123 = v2118 - int32(1)
	v2124 = int32(10)
	v2125 = base.I32_div_u_s(v2119, v2124)
	v2130 = v2119 - v2125*v2124 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2123))) = uint8(v2130)
	if base.Ui32(int32(99)) < base.Ui32(v2119) {
		v2118 = v2123
		v2119 = v2125
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v2135 = v2123
	v2136 = v2125
	goto L407
L410:
	;
	goto L409
L411:
	;
	v2140 = v2135 - int32(1)
	v2142 = v2136 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2140))) = uint8(v2142)
	v2144 = v2140
	goto L413
L412:
	;
	v2144 = v2135
	goto L413
L413:
	;
	goto L398
L414:
	;
	v2147 = v2144 - int32(1)
	v2148 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2147))) = uint8(v2148)
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v31)+556))
	v2151 = v2147
	v2152 = v2150
	goto L416
L415:
	;
	v2151 = v2144
	v2152 = v2080
	goto L416
L416:
	;
	v2153 = int32(2)
	v2154 = v54 | v2153
	v2158 = v2151 - v2153
	v2160 = l5 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v2158))) = uint8(v2160)
	if v2152 < int32(0) {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2168 = int32(45)
	goto L419
L418:
	;
	v2168 = int32(43)
	goto L419
L419:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2151-int32(1)))) = uint8(v2168)
	v2172 = int32(0)
	v2180 = v2078
	v2186 = v31 + int32(528)
	goto L420
L420:
	;
	v2207 = base.I32_trunc_sat_f64_s(v2180)
	v2210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2207)+uint32(_c_F_fmt_fp[0]))))
	v2211 = v2210 | l5&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2186))) = uint8(v2211)
	v2216 = base.F64_mul(base.F64_sub(v2180, base.F64_convert_i32_s(v2207)), float64(16))
	v2220 = int32(1)
	v2221 = v2186 + v2220
	if base.F64_eq(v2216, float64(0))&(base.B2i32(l4&int32(8) == v2172)&base.B2i32(l3 <= v2172))|base.B2i32(v2221-(v31+int32(528)) != v2220) == int32(0) {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v2239 = v94 - v2158
	v2240 = v2154 + v2239
	if int32(2147483645)-v2240 < l3 {
		v2294 = int32(-1)
		goto L11
	} else {
		goto L426
	}
L422:
	;
	v2230 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2186)+1)) = uint8(v2230)
	v2234 = v2186 + int32(2)
	goto L424
L423:
	;
	v2234 = v2221
	goto L424
L424:
	;
	if base.F64_ne(v2216, float64(0)) != 0 {
		v2180 = v2216
		v2186 = v2234
		goto L420
	} else {
		goto L425
	}
L425:
	;
	goto L421
L426:
	;
	v2244 = int32(2)
	v2247 = v31 + int32(528)
	v2248 = v2234 - v2247
	if v2248-v2244 < l3 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v2252 = l3 + v2244
	goto L429
L428:
	;
	v2252 = v2248
	goto L429
L429:
	;
	if l3 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2253 = v2252
	goto L432
L431:
	;
	v2253 = v2248
	goto L432
L432:
	;
	v2254 = v2240 + v2253
	F_pad(m, l0, int32(32), l2, v2254, l4)
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L15
	} else {
		goto L433
	}
L433:
	;
	F_out(m, l0, v2020, v2154)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L15
	} else {
		goto L434
	}
L434:
	;
	F_pad(m, l0, int32(48), l2, v2254, l4^int32(_a_F_fmt_fp_10))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L15
	} else {
		goto L435
	}
L435:
	;
	F_out(m, l0, v2247, v2248)
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L15
	} else {
		goto L436
	}
L436:
	;
	v2268 = int32(0)
	F_pad(m, l0, int32(48), v2253-v2248, v2268, v2268)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L15
	} else {
		goto L437
	}
L437:
	;
	F_out(m, l0, v2158, v2239)
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L15
	} else {
		goto L438
	}
L438:
	;
	F_pad(m, l0, int32(32), l2, v2254, l4^int32(_a_F_fmt_fp_8))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L15
	} else {
		goto L439
	}
L439:
	;
	if v2254 < l2 {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v2280 = l2
	goto L442
L441:
	;
	v2280 = v2254
	goto L442
L442:
	;
	v2294 = v2280
	goto L11
}
func F_fopen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v12 = F___strchrnul(m, int32(_a_F_fopen_0), v11)
	mBase = m.M
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 == v11&int32(255) {
		v18 = v12
	} else {
		v18 = v3
	}
	if v18 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_fopen[0])) = int32(28)
		v83 = int32(0)
	} else {
		v27 = F_strchr(m, l1, int32(43))
		mBase = m.M
		if v27 == int32(0) {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v33 = base.B2i32(v30 != int32(114))
		} else {
			v33 = int32(2)
		}
		v37 = F_strchr(m, l1, int32(120))
		mBase = m.M
		if v37 != 0 {
			v38 = v33 | int32(128)
		} else {
			v38 = v33
		}
		v42 = F_strchr(m, l1, int32(101))
		mBase = m.M
		if v42 != 0 {
			v43 = v38 | int32(_a_F_fopen_1)
		} else {
			v43 = v38
		}
		v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v46 == int32(114) {
			v49 = v43
		} else {
			v49 = v43 | int32(64)
		}
		if v46 == int32(119) {
			v54 = v49 | int32(512)
		} else {
			v54 = v49
		}
		if v46 == int32(97) {
			v59 = v54 | int32(1024)
		} else {
			v59 = v54
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(438)
		v65 = m.Env.X__syscall_openat(m, int32(-100), l0, v59|int32(_a_F_fopen_2), v8)
		mBase = m.M
		if base.Ui32(int32(-4095)) <= base.Ui32(v65) {
			*(*int32)(unsafe.Add(mBase, _c_F_fopen[0])) = int32(0) - v65
			v73 = int32(-1)
		} else {
			v73 = v65
		}
		if v73 < int32(0) {
			v83 = v3
		} else {
			v76 = F___fdopen(m, v73, l1)
			mBase = m.M
			if v76 != 0 {
				v83 = v76
			} else {
				v77 = m.Wasi_snapshot_preview1.Fd_close(m, v73)
				mBase = m.M
				v83 = int32(0)
			}
		}
	}
	m.G0 = v8 + int32(16)
	return v83
}
func F_forget_invalid_pages(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_forget_invalid_pages[0]))
	if v15 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L27
	}
L2:
	;
	m.G0 = v12 + int32(112)
	return
L3:
	;
	v19 = v12 + int32(92)
	F_hash_seq_init(m, v19, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v22 = F_hash_seq_search(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v22 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = v22
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v38 != v28 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v85 = F_hash_seq_search(m, v12+int32(92))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v40 != v27 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v42 != v26 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v44 != l1 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if base.Ui32(v46) < base.Ui32(l2) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v50 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v50 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v54 = v12 + int32(20)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	F_GetRelationPath(m, v54, v55, v56, v57, int32(-1), l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_forget_invalid_pages[0]))
	v77 = F_hash_search(m, v74, v32, int32(2), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v54
	F_errmsg_internal(m, int32(_a_F_forget_invalid_pages_0), v12)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_forget_invalid_pages_1), int32(184), int32(_a_F_forget_invalid_pages_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	if v77 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	if v85 != 0 {
		v32 = v85
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L9
L27:
	;
	F_errmsg_internal(m, int32(_a_F_forget_invalid_pages_3), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_forget_invalid_pages_1), int32(189), int32(_a_F_forget_invalid_pages_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+72)) = v8 - int32(1) | v8
	v13 = l1 * l2
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v14 == v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v193 = l0
	v195 = v13
	goto L3
L2:
	;
	v17 = v15 - v14
	if base.Ui32(v17) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v195 != 0 {
		goto L54
	} else {
		goto L55
	}
L4:
	;
	v19 = v17
	goto L6
L5:
	;
	v19 = v13
	goto L6
L6:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v19) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v19 + v14
	v193 = l0 + v19
	v195 = v13 - v19
	goto L3
L8:
	;
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v26 = l0 + v19
	if (l0^v14)&int32(3) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	base.MemoryCopy(m, l0, v14, v19)
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L7
L14:
	;
	if base.Ui32(v158) < base.Ui32(v26) {
		goto L48
	} else {
		goto L49
	}
L15:
	;
	if l0&int32(3) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	if base.Ui32(v26) < base.Ui32(int32(4)) {
		goto L39
	} else {
		goto L40
	}
L18:
	;
	v62 = v26 & int32(-4)
	if base.Ui32(v26) < base.Ui32(int32(64)) {
		v112 = v56
		v113 = v57
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v56 = v14
	v57 = l0
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v19 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v56 = v14
	v57 = l0
	goto L18
L23:
	;
	goto L24
L24:
	;
	v39 = v14
	v40 = l0
	goto L25
L25:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v44)
	v46 = int32(1)
	v47 = v39 + v46
	v49 = v40 + v46
	if v49&int32(3) == int32(0) {
		v56 = v47
		v57 = v49
		goto L18
	} else {
		goto L27
	}
L26:
	;
	v56 = v47
	v57 = v49
	goto L18
L27:
	;
	if base.Ui32(v49) < base.Ui32(v26) {
		v39 = v47
		v40 = v49
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if base.Ui32(v62) <= base.Ui32(v113) {
		v157 = v112
		v158 = v113
		goto L14
	} else {
		goto L35
	}
L30:
	;
	v66 = v62 + int32(-64)
	if base.Ui32(v66) < base.Ui32(v57) {
		v112 = v56
		v113 = v57
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v69 = v56
	v70 = v57
	goto L32
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v69)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+32)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v69)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+40)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+44)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+48)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+52)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+56)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v69)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+60)) = v104
	v106 = int32(-64)
	v107 = v69 - v106
	v109 = v70 - v106
	if base.Ui32(v109) <= base.Ui32(v66) {
		v69 = v107
		v70 = v109
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v112 = v107
	v113 = v109
	goto L29
L34:
	;
	goto L33
L35:
	;
	v119 = v112
	v120 = v113
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v124
	v126 = int32(4)
	v127 = v119 + v126
	v129 = v120 + v126
	if base.Ui32(v129) < base.Ui32(v62) {
		v119 = v127
		v120 = v129
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v157 = v127
	v158 = v129
	goto L14
L38:
	;
	goto L37
L39:
	;
	v157 = v14
	v158 = l0
	goto L14
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v19) < base.Ui32(int32(4)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v157 = v14
	v158 = l0
	goto L14
L43:
	;
	goto L44
L44:
	;
	v138 = v14
	v139 = l0
	goto L45
L45:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v143)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)) = uint8(v145)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+2)) = uint8(v147)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+3)) = uint8(v149)
	v151 = int32(4)
	v152 = v138 + v151
	v154 = v139 + v151
	if base.Ui32(v154) <= base.Ui32(v26-int32(4)) {
		v138 = v152
		v139 = v154
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v157 = v152
	v158 = v154
	goto L14
L47:
	;
	goto L46
L48:
	;
	v164 = v157
	v165 = v158
	goto L51
L49:
	;
	goto L50
L50:
	;
	goto L7
L51:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v169)
	v171 = int32(1)
	v174 = v165 + v171
	if v174 != v26 {
		v164 = v164 + v171
		v165 = v174
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	v196 = v193
	v200 = v195
	goto L57
L55:
	;
	goto L56
L56:
	;
	if l1 != 0 {
		goto L68
	} else {
		goto L69
	}
L57:
	;
	v203 = F___toread(m, l3)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L56
L59:
	;
	v217 = v200 - v210
	if v217 != 0 {
		v196 = v196 + v210
		v200 = v217
		goto L57
	} else {
		goto L67
	}
L60:
	;
	return int32(0)
L61:
	;
	if v203 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v210 = m.T0[v209].(func(*base.Module, int32, int32, int32) int32)(m, l3, v196, v200)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v214 = base.I32_div_u_s(v13-v200, l1)
	return v214
L65:
	;
	if v210 != 0 {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	goto L58
L68:
	;
	v226 = l2
	goto L70
L69:
	;
	v226 = int32(0)
	goto L70
L70:
	;
	return v226
}
func F_free_conversion_map(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_free_attrmap(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_pfree(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					F_pfree(m, v14)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_fsm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 <= int32(0) {
		v64 = v13
	} else {
		v18 = v14 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v14) {
			v25 = int32(0)
			v27 = v13
			for {
				v33 = v27 * int32(-1799145247)
				v35 = v25 + int32(8)
				if v35 != v14&int32(2147483640) {
					v25 = v35
					v27 = v33
					continue
				} else {
					break
				}
				break
			}
			if v18 == int32(0) {
				v64 = v33
			} else {
				v42 = v33
				v49 = int32(0)
				v51 = v42
				for {
					v57 = v51 * int32(4069)
					v59 = v49 + int32(1)
					if v59 != v18 {
						v49 = v59
						v51 = v57
						continue
					} else {
						break
					}
					break
				}
				v64 = v57
			}
		} else {
			v42 = v13
			v49 = int32(0)
			v51 = v42
			for {
				v57 = v51 * int32(4069)
				v59 = v49 + int32(1)
				if v59 != v18 {
					v49 = v59
					v51 = v57
					continue
				} else {
					break
				}
				break
			}
			v64 = v57
		}
	}
	v70 = base.I32_div_u_s(v64, int32(4069))
	v73 = base.I32_div_u_s(v64, int32(16556761))
	v76 = v64 + v70 + v73 + int32(3)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v79 == int32(0) {
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v83
		v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v85
		v89 = F_smgropen(m, v11+int32(24), v82)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+72))
			if v95 != 0 {
				v103 = v95
			} else {
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+76))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v97
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v97))) = v99
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v89)+72))
				v103 = v101
			}
			*(*int32)(unsafe.Add(mBase, uint32(v89)+72)) = v103 + int32(1)
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v108 = v107
			v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
			v112 = v76 + (v14 ^ int32(-1))
			if base.B2i32(v109 != int32(-1))&base.B2i32(base.Ui32(v112) < base.Ui32(v109)) != 0 {
				v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					v153 = v150
					if v153 < int32(0) {
						v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
						v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
						v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
						if v164 != 0 {
							v252 = v153
							m.G0 = v11 + int32(48)
							return v252
						} else {
							F_LockBufferInternal(m, v153, int32(3))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return int32(0)
							} else {
								v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
								v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
								v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
								if v172 == int32(0) {
									v195 = v171
									v196 = int32(_a_F_fsm_readbuf_0)
									v197 = int32(0)
									if v197|(v195&int32(3)|int32(1)) == v197 {
										v214 = v195 + v196
										v216 = v195 + int32(4)
										if base.Ui32(v216) < base.Ui32(v214) {
											v218 = v214
										} else {
											v218 = v216
										}
										v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
										if v223 == int32(0) {
										} else {
											base.MemoryFill(m, v195, int32(0), v223)
										}
									} else {
										base.MemoryFill(m, v195, int32(0), v196)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
									v237 = int32(_a_F_fsm_readbuf_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
									v243 = int32(_a_F_fsm_readbuf_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
								} else {
								}
								F_UnlockBuffer(m, v153)
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return int32(0)
								} else {
									v252 = v153
									m.G0 = v11 + int32(48)
									return v252
								}
							}
						}
					} else {
						v176 = v153 << (uint(int32(13)) % 32)
						v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
						v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
						if v182 != 0 {
							v252 = v153
							m.G0 = v11 + int32(48)
							return v252
						} else {
							F_LockBufferInternal(m, v153, int32(3))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int32(0)
							} else {
								v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
								v188 = v187 + v176
								v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
								if v191 != 0 {
								} else {
									v195 = v188 + int32(-8192)
									v196 = int32(_a_F_fsm_readbuf_0)
									v197 = int32(0)
									if v197|(v195&int32(3)|int32(1)) == v197 {
										v214 = v195 + v196
										v216 = v195 + int32(4)
										if base.Ui32(v216) < base.Ui32(v214) {
											v218 = v214
										} else {
											v218 = v216
										}
										v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
										if v223 == int32(0) {
										} else {
											base.MemoryFill(m, v195, int32(0), v223)
										}
									} else {
										base.MemoryFill(m, v195, int32(0), v196)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
									v237 = int32(_a_F_fsm_readbuf_2)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
									v243 = int32(_a_F_fsm_readbuf_0)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
									*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
								}
								F_UnlockBuffer(m, v153)
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return int32(0)
								} else {
									v252 = v153
									m.G0 = v11 + int32(48)
									return v252
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(-1)
				v118 = F_smgrexists(m, v108, int32(1))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					if v118 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(0)
						v129 = int32(0)
						if l2 == v129 {
							v252 = v129
							m.G0 = v11 + int32(48)
							return v252
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
							v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
							v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return int32(0)
							} else {
								v153 = v145
								if v153 < int32(0) {
									v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
									v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
									if v164 != 0 {
										v252 = v153
										m.G0 = v11 + int32(48)
										return v252
									} else {
										F_LockBufferInternal(m, v153, int32(3))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
											v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
											if v172 == int32(0) {
												v195 = v171
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											} else {
											}
											F_UnlockBuffer(m, v153)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return int32(0)
											} else {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											}
										}
									}
								} else {
									v176 = v153 << (uint(int32(13)) % 32)
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
									v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
									if v182 != 0 {
										v252 = v153
										m.G0 = v11 + int32(48)
										return v252
									} else {
										F_LockBufferInternal(m, v153, int32(3))
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v188 = v187 + v176
											v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
											if v191 != 0 {
											} else {
												v195 = v188 + int32(-8192)
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											}
											F_UnlockBuffer(m, v153)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return int32(0)
											} else {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											}
										}
									}
								}
							}
						}
					} else {
						v125 = F_smgrnblocks(m, v108, int32(1))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
							if base.Ui32(v112) < base.Ui32(v127) {
								v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									v153 = v150
									if v153 < int32(0) {
										v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
										v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
										v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
										if v164 != 0 {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										} else {
											F_LockBufferInternal(m, v153, int32(3))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
												v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
												if v172 == int32(0) {
													v195 = v171
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												} else {
												}
												F_UnlockBuffer(m, v153)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return int32(0)
												} else {
													v252 = v153
													m.G0 = v11 + int32(48)
													return v252
												}
											}
										}
									} else {
										v176 = v153 << (uint(int32(13)) % 32)
										v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
										if v182 != 0 {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										} else {
											F_LockBufferInternal(m, v153, int32(3))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return int32(0)
											} else {
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
												v188 = v187 + v176
												v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
												if v191 != 0 {
												} else {
													v195 = v188 + int32(-8192)
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												}
												F_UnlockBuffer(m, v153)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return int32(0)
												} else {
													v252 = v153
													m.G0 = v11 + int32(48)
													return v252
												}
											}
										}
									}
								}
							} else {
								v129 = int32(0)
								if l2 == v129 {
									v252 = v129
									m.G0 = v11 + int32(48)
									return v252
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
									v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
									*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
									v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										v153 = v145
										if v153 < int32(0) {
											v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
											v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
											v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
											if v164 != 0 {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											} else {
												F_LockBufferInternal(m, v153, int32(3))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
													v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
													v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
													if v172 == int32(0) {
														v195 = v171
														v196 = int32(_a_F_fsm_readbuf_0)
														v197 = int32(0)
														if v197|(v195&int32(3)|int32(1)) == v197 {
															v214 = v195 + v196
															v216 = v195 + int32(4)
															if base.Ui32(v216) < base.Ui32(v214) {
																v218 = v214
															} else {
																v218 = v216
															}
															v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
															if v223 == int32(0) {
															} else {
																base.MemoryFill(m, v195, int32(0), v223)
															}
														} else {
															base.MemoryFill(m, v195, int32(0), v196)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
														v237 = int32(_a_F_fsm_readbuf_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
														v243 = int32(_a_F_fsm_readbuf_0)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
													} else {
													}
													F_UnlockBuffer(m, v153)
													mBase = m.M
													v249 = m.ExcPending
													if v249 != 0 {
														return int32(0)
													} else {
														v252 = v153
														m.G0 = v11 + int32(48)
														return v252
													}
												}
											}
										} else {
											v176 = v153 << (uint(int32(13)) % 32)
											v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
											if v182 != 0 {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											} else {
												F_LockBufferInternal(m, v153, int32(3))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return int32(0)
												} else {
													v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
													v188 = v187 + v176
													v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
													if v191 != 0 {
													} else {
														v195 = v188 + int32(-8192)
														v196 = int32(_a_F_fsm_readbuf_0)
														v197 = int32(0)
														if v197|(v195&int32(3)|int32(1)) == v197 {
															v214 = v195 + v196
															v216 = v195 + int32(4)
															if base.Ui32(v216) < base.Ui32(v214) {
																v218 = v214
															} else {
																v218 = v216
															}
															v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
															if v223 == int32(0) {
															} else {
																base.MemoryFill(m, v195, int32(0), v223)
															}
														} else {
															base.MemoryFill(m, v195, int32(0), v196)
														}
														*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
														v237 = int32(_a_F_fsm_readbuf_2)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
														v243 = int32(_a_F_fsm_readbuf_0)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
														*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
													}
													F_UnlockBuffer(m, v153)
													mBase = m.M
													v249 = m.ExcPending
													if v249 != 0 {
														return int32(0)
													} else {
														v252 = v153
														m.G0 = v11 + int32(48)
														return v252
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
		v108 = v79
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
		v112 = v76 + (v14 ^ int32(-1))
		if base.B2i32(v109 != int32(-1))&base.B2i32(base.Ui32(v112) < base.Ui32(v109)) != 0 {
			v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int32(0)
			} else {
				v153 = v150
				if v153 < int32(0) {
					v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
					v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
					v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
					if v164 != 0 {
						v252 = v153
						m.G0 = v11 + int32(48)
						return v252
					} else {
						F_LockBufferInternal(m, v153, int32(3))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
							v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
							if v172 == int32(0) {
								v195 = v171
								v196 = int32(_a_F_fsm_readbuf_0)
								v197 = int32(0)
								if v197|(v195&int32(3)|int32(1)) == v197 {
									v214 = v195 + v196
									v216 = v195 + int32(4)
									if base.Ui32(v216) < base.Ui32(v214) {
										v218 = v214
									} else {
										v218 = v216
									}
									v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
									if v223 == int32(0) {
									} else {
										base.MemoryFill(m, v195, int32(0), v223)
									}
								} else {
									base.MemoryFill(m, v195, int32(0), v196)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
								v237 = int32(_a_F_fsm_readbuf_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
								v243 = int32(_a_F_fsm_readbuf_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
							} else {
							}
							F_UnlockBuffer(m, v153)
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								v252 = v153
								m.G0 = v11 + int32(48)
								return v252
							}
						}
					}
				} else {
					v176 = v153 << (uint(int32(13)) % 32)
					v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
					v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
					if v182 != 0 {
						v252 = v153
						m.G0 = v11 + int32(48)
						return v252
					} else {
						F_LockBufferInternal(m, v153, int32(3))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int32(0)
						} else {
							v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
							v188 = v187 + v176
							v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
							if v191 != 0 {
							} else {
								v195 = v188 + int32(-8192)
								v196 = int32(_a_F_fsm_readbuf_0)
								v197 = int32(0)
								if v197|(v195&int32(3)|int32(1)) == v197 {
									v214 = v195 + v196
									v216 = v195 + int32(4)
									if base.Ui32(v216) < base.Ui32(v214) {
										v218 = v214
									} else {
										v218 = v216
									}
									v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
									if v223 == int32(0) {
									} else {
										base.MemoryFill(m, v195, int32(0), v223)
									}
								} else {
									base.MemoryFill(m, v195, int32(0), v196)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
								v237 = int32(_a_F_fsm_readbuf_2)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
								v243 = int32(_a_F_fsm_readbuf_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
								*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
							}
							F_UnlockBuffer(m, v153)
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								v252 = v153
								m.G0 = v11 + int32(48)
								return v252
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(-1)
			v118 = F_smgrexists(m, v108, int32(1))
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				if v118 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(0)
					v129 = int32(0)
					if l2 == v129 {
						v252 = v129
						m.G0 = v11 + int32(48)
						return v252
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
						v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
						v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							v153 = v145
							if v153 < int32(0) {
								v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
								v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
								v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
								if v164 != 0 {
									v252 = v153
									m.G0 = v11 + int32(48)
									return v252
								} else {
									F_LockBufferInternal(m, v153, int32(3))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
										v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
										if v172 == int32(0) {
											v195 = v171
											v196 = int32(_a_F_fsm_readbuf_0)
											v197 = int32(0)
											if v197|(v195&int32(3)|int32(1)) == v197 {
												v214 = v195 + v196
												v216 = v195 + int32(4)
												if base.Ui32(v216) < base.Ui32(v214) {
													v218 = v214
												} else {
													v218 = v216
												}
												v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
												if v223 == int32(0) {
												} else {
													base.MemoryFill(m, v195, int32(0), v223)
												}
											} else {
												base.MemoryFill(m, v195, int32(0), v196)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
											v237 = int32(_a_F_fsm_readbuf_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
											v243 = int32(_a_F_fsm_readbuf_0)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
										} else {
										}
										F_UnlockBuffer(m, v153)
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return int32(0)
										} else {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										}
									}
								}
							} else {
								v176 = v153 << (uint(int32(13)) % 32)
								v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
								v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
								if v182 != 0 {
									v252 = v153
									m.G0 = v11 + int32(48)
									return v252
								} else {
									F_LockBufferInternal(m, v153, int32(3))
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v188 = v187 + v176
										v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
										if v191 != 0 {
										} else {
											v195 = v188 + int32(-8192)
											v196 = int32(_a_F_fsm_readbuf_0)
											v197 = int32(0)
											if v197|(v195&int32(3)|int32(1)) == v197 {
												v214 = v195 + v196
												v216 = v195 + int32(4)
												if base.Ui32(v216) < base.Ui32(v214) {
													v218 = v214
												} else {
													v218 = v216
												}
												v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
												if v223 == int32(0) {
												} else {
													base.MemoryFill(m, v195, int32(0), v223)
												}
											} else {
												base.MemoryFill(m, v195, int32(0), v196)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
											v237 = int32(_a_F_fsm_readbuf_2)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
											v243 = int32(_a_F_fsm_readbuf_0)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
											*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
										}
										F_UnlockBuffer(m, v153)
										mBase = m.M
										v249 = m.ExcPending
										if v249 != 0 {
											return int32(0)
										} else {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										}
									}
								}
							}
						}
					}
				} else {
					v125 = F_smgrnblocks(m, v108, int32(1))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
						if base.Ui32(v112) < base.Ui32(v127) {
							v150 = F_ReadBufferExtended(m, l0, int32(1), v112, int32(3), int32(0))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								v153 = v150
								if v153 < int32(0) {
									v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
									v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
									if v164 != 0 {
										v252 = v153
										m.G0 = v11 + int32(48)
										return v252
									} else {
										F_LockBufferInternal(m, v153, int32(3))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return int32(0)
										} else {
											v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
											v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
											v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
											if v172 == int32(0) {
												v195 = v171
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											} else {
											}
											F_UnlockBuffer(m, v153)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return int32(0)
											} else {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											}
										}
									}
								} else {
									v176 = v153 << (uint(int32(13)) % 32)
									v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
									v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
									if v182 != 0 {
										v252 = v153
										m.G0 = v11 + int32(48)
										return v252
									} else {
										F_LockBufferInternal(m, v153, int32(3))
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
											v188 = v187 + v176
											v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
											if v191 != 0 {
											} else {
												v195 = v188 + int32(-8192)
												v196 = int32(_a_F_fsm_readbuf_0)
												v197 = int32(0)
												if v197|(v195&int32(3)|int32(1)) == v197 {
													v214 = v195 + v196
													v216 = v195 + int32(4)
													if base.Ui32(v216) < base.Ui32(v214) {
														v218 = v214
													} else {
														v218 = v216
													}
													v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
													if v223 == int32(0) {
													} else {
														base.MemoryFill(m, v195, int32(0), v223)
													}
												} else {
													base.MemoryFill(m, v195, int32(0), v196)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
												v237 = int32(_a_F_fsm_readbuf_2)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
												v243 = int32(_a_F_fsm_readbuf_0)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
												*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
											}
											F_UnlockBuffer(m, v153)
											mBase = m.M
											v249 = m.ExcPending
											if v249 != 0 {
												return int32(0)
											} else {
												v252 = v153
												m.G0 = v11 + int32(48)
												return v252
											}
										}
									}
								}
							}
						} else {
							v129 = int32(0)
							if l2 == v129 {
								v252 = v129
								m.G0 = v11 + int32(48)
								return v252
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
								v135 = *(*int64)(unsafe.Add(mBase, uint32(v11)+36))
								*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v135
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v137
								v145 = F_ExtendBufferedRelTo(m, v11+int32(8), int32(1), int32(20), v76-v14, int32(3))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									v153 = v145
									if v153 < int32(0) {
										v159 = (v153 ^ int32(-1)) << (uint(int32(2)) % 32)
										v161 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v161)))
										v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
										if v164 != 0 {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										} else {
											F_LockBufferInternal(m, v153, int32(3))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[0]))
												v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v159)))
												v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+14)))
												if v172 == int32(0) {
													v195 = v171
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												} else {
												}
												F_UnlockBuffer(m, v153)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return int32(0)
												} else {
													v252 = v153
													m.G0 = v11 + int32(48)
													return v252
												}
											}
										}
									} else {
										v176 = v153 << (uint(int32(13)) % 32)
										v178 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
										v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v178-int32(_a_F_fsm_readbuf_3)))))
										if v182 != 0 {
											v252 = v153
											m.G0 = v11 + int32(48)
											return v252
										} else {
											F_LockBufferInternal(m, v153, int32(3))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return int32(0)
											} else {
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_readbuf[1]))
												v188 = v187 + v176
												v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188-int32(_a_F_fsm_readbuf_3)))))
												if v191 != 0 {
												} else {
													v195 = v188 + int32(-8192)
													v196 = int32(_a_F_fsm_readbuf_0)
													v197 = int32(0)
													if v197|(v195&int32(3)|int32(1)) == v197 {
														v214 = v195 + v196
														v216 = v195 + int32(4)
														if base.Ui32(v216) < base.Ui32(v214) {
															v218 = v214
														} else {
															v218 = v216
														}
														v223 = (v195^int32(-1)+v218)&int32(-4) + int32(4)
														if v223 == int32(0) {
														} else {
															base.MemoryFill(m, v195, int32(0), v223)
														}
													} else {
														base.MemoryFill(m, v195, int32(0), v196)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v195)+10)) = int32(_a_F_fsm_readbuf_1)
													v237 = int32(_a_F_fsm_readbuf_2)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+18)) = uint16(v237)
													v243 = int32(_a_F_fsm_readbuf_0)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+16)) = uint16(v243)
													*(*uint16)(unsafe.Add(mBase, uint32(v195)+14)) = uint16(v243)
												}
												F_UnlockBuffer(m, v153)
												mBase = m.M
												v249 = m.ExcPending
												if v249 != 0 {
													return int32(0)
												} else {
													v252 = v153
													m.G0 = v11 + int32(48)
													return v252
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
func F_ftruncate(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_ftruncate64(m, l0, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v3) {
		*(*int32)(unsafe.Add(mBase, _c_F_ftruncate[0])) = int32(0) - v3
		v11 = int32(-1)
	} else {
		v11 = v3
	}
	return v11
}
