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
										F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1720), int32(_a_F_FindFKPeriodOpers_3))
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
							F_errdetail(m, int32(_a_F_FindFKPeriodOpers_6), int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1682), int32(_a_F_FindFKPeriodOpers_3))
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
					F_errfinish(m, int32(_a_F_FindFKPeriodOpers_2), int32(1686), int32(_a_F_FindFKPeriodOpers_3))
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
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
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int64
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
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
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
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
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int64
	_ = v405
	var v406 int32
	_ = v406
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v483 int32
	_ = v483
	var v486 int64
	_ = v486
	var v487 int32
	_ = v487
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
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
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[0]))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[1])))
	if v33 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_before_shmem_exit(m, int32(393), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v48 = F_LWLockAcquire(m, v44+int32(2304), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[1])) = uint8(v41)
	goto L3
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if int32(0) < v52 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L4
	} else {
		goto L178
	}
L8:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L4
	} else {
		goto L175
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L170
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L4
	} else {
		goto L166
	}
L11:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v211
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[5])) = v86
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v216+int32(2304))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L45
	}
L12:
	;
	v60 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v187+int32(2304))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L40
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v51+int32(8)+v60<<(uint(int32(2))%32))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+44)))
	if v87 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v159 = v60 + int32(1)
	if v159 != v52 {
		v60 = v159
		goto L15
	} else {
		goto L39
	}
L18:
	;
	v91 = v86 + int32(47)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v94 == int32(0))|base.B2i32(v94 != v97) != 0 {
		v115 = v94
		v116 = v97
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v115-v116 != 0 {
		goto L17
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v100 = v91
	v101 = l0
	goto L22
L22:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v105
		v116 = v104
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v115 = v105
	v116 = v104
	goto L20
L24:
	;
	v108 = int32(1)
	if v105 == v104 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v118 != int32(-1) {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[6]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	if v125 != v31 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v127 = F_superuser_arg(m, v31)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[7]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v124+v121*int32(640))+60))
	if v132 == v136 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	if v127 == int32(0) {
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
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_0), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errhint(m, int32(_a_F_FinishPreparedTransaction_1), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(599), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = l0
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_4), v26+int32(-16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(615), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
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
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[6]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v86)+32))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+45)))
	if v226 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)+40))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+36))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)+32))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v240)+28))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240)+54)))
	v254 = v240 + (v247+int32(7))&int32(_a_F_FinishPreparedTransaction_5) + int32(72)
	v258 = v246 - int32(1)
	if v258 < int32(0) {
		v326 = v225
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v230 = F_ReadTwoPhaseFile(m, v225, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v86)+16))
	F_XlogReadTwoPhaseData(m, v233, v26+int32(-4), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v230
	v240 = v230
	goto L46
L51:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v240 = v239
	goto L46
L52:
	;
	v331 = int32(_a_F_FinishPreparedTransaction_6)
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8])) = v333 + int32(1)
	v339 = int32(7)
	v341 = int32(-8)
	v343 = v254 + (v246<<(uint(int32(2))%32)+v339)&v341
	v344 = int32(12)
	v350 = v343 + (v245*v344+v339)&v341
	v357 = v350 + (v244*v344+v339)&v341
	v358 = int32(4)
	v360 = v357 + v243<<(uint(v358)%32)
	v363 = v360 + v242<<(uint(v358)%32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v240)+28))
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
	if v246&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v263 = int32(2)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v254+v258<<(uint(v263)%32))))
	if base.B2i32(base.Ui32(v263) < base.Ui32(v266))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v225)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v281 = v225
	v283 = v258
	goto L57
L57:
	;
	if v258 == int32(0) {
		v326 = v281
		goto L53
	} else {
		goto L65
	}
L58:
	;
	v281 = v278
	v283 = v246 - int32(2)
	goto L57
L59:
	;
	v278 = v266
	goto L58
L60:
	;
	if base.Ui32(v225) < base.Ui32(v266) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if int32(0) <= v225-v266 {
		v278 = v225
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v278 = v225
	goto L58
L64:
	;
	goto L59
L65:
	;
	v286 = v281
	v287 = v283
	goto L66
L66:
	;
	v291 = int32(3)
	v295 = v254 + v287<<(uint(int32(2))%32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if base.B2i32(base.Ui32(v286) < base.Ui32(v291))|base.B2i32(base.Ui32(v296) < base.Ui32(v291)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v326 = v321
	goto L53
L68:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v295-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v309))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v306)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v306 = v296
	goto L68
L70:
	;
	if v286-v296 < int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(v296) <= base.Ui32(v286) {
		v306 = v286
		goto L68
	} else {
		goto L74
	}
L73:
	;
	v306 = v286
	goto L68
L74:
	;
	goto L69
L75:
	;
	if int32(1) < v287 {
		v286 = v321
		v287 = v287 - int32(2)
		goto L66
	} else {
		goto L82
	}
L76:
	;
	v321 = v309
	goto L75
L77:
	;
	if base.Ui32(v306) < base.Ui32(v309) {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int32(0) <= v306-v309 {
		v321 = v306
		goto L75
	} else {
		goto L81
	}
L80:
	;
	v321 = v306
	goto L75
L81:
	;
	goto L76
L82:
	;
	goto L67
L83:
	;
	v528 = v363 + v241<<(uint(int32(4))%32)
	F_ProcArrayRemove(m, v224+v221*int32(640), v326)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L4
	} else {
		goto L110
	}
L84:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+52)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v240)+40))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v240)+32))
	v372 = m.G0
	v373 = int32(16)
	v374 = v372 - v373
	m.G0 = v374
	F_gettimeofday(m, v374)
	mBase = m.M
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v374)))
	v378 = int64(*(*int32)(unsafe.Add(mBase, uint32(v374)+8)))
	m.G0 = v374 + v373
	v386 = v378 + v377*int64(1000000) - int64(946684800000000)
	goto L87
L85:
	;
	goto L86
L86:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v240)+36))
	v456 = F_TransactionIdDidCommit(m, v225)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L99
	}
L87:
	;
	v387 = int32(_a_F_FinishPreparedTransaction_7)
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10]))
	v390 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])) = v389 + v390
	v394 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])))
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[11]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+120)) = v397 | v390
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[12]))
	v405 = F_XactLogCommitRecord(m, v386, v364, v254, v368, v343, v367, v357, v366, v363, v365, v402|int32(2), v225, l0)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v394+int32(1))&int32(_a_F_FinishPreparedTransaction_8)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[9])))
	F_TransactionTreeSetCommitTsData(m, v225, v364, v254, v426, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L95
	}
L90:
	;
	v414 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[13]))
	v416 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[14]))
	F_replorigin_session_advance(m, v414, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[15])) = v386
	v426 = v386
	goto L89
L93:
	;
	v420 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[15]))
	if v420 != int64(0) {
		v426 = v420
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	F_XLogFlush(m, v405)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_TransactionIdCommitTree(m, v225, v364, v254)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[11]))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+120)) = v439 & int32(-2)
	v443 = int32(_a_F_FinishPreparedTransaction_7)
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10]))
	v446 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])) = v445 - v446
	F_SyncRepWaitForLSN(m, v405, v446)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v515 = v240 + int32(32)
	v517 = v343
	goto L83
L99:
	;
	if v456 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	v458 = int32(_a_F_FinishPreparedTransaction_7)
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])) = v460 + int32(1)
	v467 = m.G0
	v468 = int32(16)
	v469 = v467 - v468
	m.G0 = v469
	F_gettimeofday(m, v469)
	mBase = m.M
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v469)))
	v473 = int64(*(*int32)(unsafe.Add(mBase, uint32(v469)+8)))
	m.G0 = v469 + v468
	goto L101
L101:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[12]))
	v486 = F_XactLogAbortRecord(m, v473+v472*int64(1000000)-int64(946684800000000), v364, v254, v455, v350, v454, v360, v483|int32(2), v225, l0)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32((v453-int32(1))&int32(_a_F_FinishPreparedTransaction_8)) <= base.Ui32(int32(_a_F_FinishPreparedTransaction_9)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v495 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[13]))
	v497 = *(*int64)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[14]))
	F_replorigin_session_advance(m, v495, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_XLogFlush(m, v486)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	F_TransactionIdAbortTree(m, v225, v364, v254)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v506 = int32(_a_F_FinishPreparedTransaction_7)
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[10])) = v508 - int32(1)
	F_SyncRepWaitForLSN(m, v486, int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v515 = v240 + int32(36)
	v517 = v350
	goto L83
L110:
	;
	v534 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+44)) = uint8(v534)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	F_DropRelationFiles(m, v517, v536, v534)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
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
	v690 = m.G0
	v692 = v690 - int32(16)
	m.G0 = v692
	*(*int32)(unsafe.Add(mBase, uint32(v692)+12)) = v225
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v700 = F_LWLockAcquire(m, v696+int32(3584), int32(1))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L145
	}
L113:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v240)+40))
	F_pgstat_execute_transactional_drops(m, v540, v357)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	F_pgstat_execute_transactional_drops(m, v609, v360)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L4
	} else {
		goto L135
	}
L116:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+52)))
	if v543 == int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	F_SIInsertDataEntries(m, v363, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+52)))
	if v551 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v561 = F_LWLockAcquire(m, v557+int32(2304), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+4)))
	if v563 == int32(0) {
		goto L112
	} else {
		goto L127
	}
L127:
	;
	v568 = v528
	v569 = v563
	goto L128
L128:
	;
	v592 = v568 + int32(8)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v569&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_FinishPreparedTransaction[16])))
	if v597 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L112
L130:
	;
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v568)+6)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	m.T0[v597].(func(*base.Module, int32, int32, int32, int32))(m, v225, v598, v592, v599)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	v607 = v592 + (v602+int32(7))&int32(-8)
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+4)))
	if v608 != 0 {
		v568 = v607
		v569 = v608
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
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	v617 = F_LWLockAcquire(m, v613+int32(2304), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+4)))
	if v619 == int32(0) {
		goto L112
	} else {
		goto L137
	}
L137:
	;
	v624 = v528
	v625 = v619
	goto L138
L138:
	;
	v648 = v624 + int32(8)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v625&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_FinishPreparedTransaction[17])))
	if v653 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L112
L140:
	;
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v624)+6)))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	m.T0[v653].(func(*base.Module, int32, int32, int32, int32))(m, v225, v654, v648, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	v663 = v648 + (v658+int32(7))&int32(-8)
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+4)))
	if v664 != 0 {
		v624 = v663
		v625 = v664
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
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[18]))
	v706 = int32(0)
	v708 = F_hash_search(m, v703, v692+int32(12), v706, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v711 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v711+int32(3584))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	if v708 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	v718 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[19])) = uint8(v718)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[20])) = v716
	F_ReleasePredicateLocks(m, l1, int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	m.G0 = v692 + int32(16)
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[3]))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	if v731 <= int32(0) {
		goto L7
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+45)))
	v736 = v730 + int32(8)
	v740 = int32(0)
	goto L153
L153:
	;
	v765 = v736 + v740<<(uint(int32(2))%32)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	if v766 != v86 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v772 = v731 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v730)+4)) = v772
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v736+v772<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v765))) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v779
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v86
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[2]))
	F_LWLockRelease(m, v783+int32(2304))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L159
	}
L155:
	;
	v769 = v740 + int32(1)
	if v731 != v769 {
		v740 = v769
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
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	if v734&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	F_RemoveTwoPhaseFile(m, v225, int32(1))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L4
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v796 = int32(_a_F_FinishPreparedTransaction_6)
	v798 = *(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[8])) = v798 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_FinishPreparedTransaction[5])) = int32(0)
	F_pfree(m, v240)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	m.G0 = v28 - int32(-64)
	return
L166:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l0
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_10), v26+int32(-32))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(581), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
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
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_FinishPreparedTransaction_11), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errhint(m, int32(_a_F_FinishPreparedTransaction_12), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(587), int32(_a_F_FinishPreparedTransaction_3))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v225
	F_errmsg_internal(m, int32(_a_F_FinishPreparedTransaction_13), v26+int32(-48))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(2419), int32(_a_F_FinishPreparedTransaction_14))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v86
	F_errmsg_internal(m, int32(_a_F_FinishPreparedTransaction_15), v28)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_FinishPreparedTransaction_2), int32(650), int32(_a_F_FinishPreparedTransaction_16))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
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
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_FlushOneBuffer[0]))
	F_FlushBuffer(m, v3+l0<<(uint(int32(6))%32)+int32(-64), int32(0), int32(3))
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
func F_fdw_handler_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_fdw_handler_in_0), int32(369), int32(_a_F_fdw_handler_in_1), int32(_a_F_fdw_handler_in_2), int32(_a_F_fdw_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v100 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v99 + v100
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v51 = F_palloc(m, v48<<(uint(int32(2))%32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v48 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 <= v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v20
	goto L7
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v28 = int32(0)
	v30 = v4
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v28<<(uint(int32(2))%32))))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v48 = v39
	goto L1
L10:
	;
	v39 = v30 + base.I32_popcnt(v36)
	goto L12
L11:
	;
	v39 = v30
	goto L12
L12:
	;
	v41 = v28 + int32(1)
	if v41 != v24 {
		v28 = v41
		v30 = v39
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	return int32(0)
L15:
	;
	if l1 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if int32(0) <= v111 {
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v111 = base.I32_ctz(v97) | v98<<(uint(int32(5))%32)
	goto L16
L18:
	;
	v111 = int32(-2)
	goto L16
L19:
	;
	v64 = base.I32_div_s(int32(0), int32(32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v65 <= v64 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = l1 + int32(8)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
	v75 = v72 & int32(-1)
	if v75 != 0 {
		v97 = v75
		v98 = v64
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v77 = v64 + int32(1)
	if v77 == v65 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v80 = v77
	goto L23
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v68+v80<<(uint(int32(2))%32))))
	if v87 != 0 {
		v97 = v87
		v98 = v80
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L18
L25:
	;
	v89 = v80 + int32(1)
	if v89 != v65 {
		v80 = v89
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v117 = v111
	v119 = v4
	goto L30
L28:
	;
	v216 = v4
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v216
	m.G0 = v11 + int32(16)
	return v51
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v117<<(uint(int32(2))%32))))
	if v126 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v216 = v152
	goto L29
L32:
	;
	if l1 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	v129 = F_find_base_rel_ignore_join(m, l0, v117)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v119<<(uint(int32(2))%32)))) = v126
	v152 = v119 + int32(1)
	goto L32
L36:
	;
	if v129 == int32(0) {
		v152 = v119
		goto L32
	} else {
		goto L37
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v117
	F_errmsg_internal(m, int32(_a_F_find_appinfos_by_relids_0), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_find_appinfos_by_relids_1), int32(778), int32(_a_F_find_appinfos_by_relids_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	if int32(0) <= v208 {
		v117 = v208
		v119 = v152
		goto L30
	} else {
		goto L52
	}
L42:
	;
	v208 = base.I32_ctz(v194) | v195<<(uint(int32(5))%32)
	goto L41
L43:
	;
	v208 = int32(-2)
	goto L41
L44:
	;
	v159 = v117 + int32(1)
	v161 = base.I32_div_s(v159, int32(32))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v162 <= v161 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v165 = l1 + int32(8)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+v161<<(uint(int32(2))%32))))
	v172 = v169 & (int32(-1) << (uint(v159) % 32))
	if v172 != 0 {
		v194 = v172
		v195 = v161
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v174 = v161 + int32(1)
	if v174 == v162 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v177 = v174
	goto L48
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v165+v177<<(uint(int32(2))%32))))
	if v184 != 0 {
		v194 = v184
		v195 = v177
		goto L42
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v186 = v177 + int32(1)
	if v186 != v162 {
		v177 = v186
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L31
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
			v34 = F_expression_tree_walker_impl(m, l0, int32(693), l1)
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
			v27 = F_expression_tree_walker_impl(m, l0, int32(693), l1)
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
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
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
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
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
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
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
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
	return v406
L2:
	;
	v406 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 != int32(21) {
		v406 = l0
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
		v406 = l0
		goto L1
	}
L6:
	;
	v402 = F_pull_ands(m, v395)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L20
	} else {
		goto L144
	}
L7:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v388 != int32(1) {
		v395 = v381
		goto L6
	} else {
		goto L143
	}
L8:
	;
	v377 = F_make_orclause(m, v84)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
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
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v406 = v304
	goto L1
L11:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 == int32(0) {
		v284 = v3
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
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
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
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v55 == int32(0) {
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
	v406 = v60
	goto L1
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v64 == int32(0) {
		v70 = v32
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v406 = v41
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
	v406 = v90
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
	v288 = F_pull_ands(m, v284)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L20
	} else {
		goto L106
	}
L85:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v234 <= int32(0) {
		v284 = v3
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
	v284 = v274
	goto L84
L89:
	;
	v276 = v239 + int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v276 < v277 {
		v239 = v276
		v242 = v274
		goto L87
	} else {
		goto L105
	}
L90:
	;
	v271 = F_lappend(m, v242, v251)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
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
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+24)))
	if l1 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if v258&int32(1) != 0 {
		v274 = v242
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
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v251)+20))
	if v261 != 0 {
		v274 = v242
		goto L89
	} else {
		goto L98
	}
L98:
	;
	v406 = v251
	goto L1
L99:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v251)+20))
	if v266 != 0 {
		v274 = v242
		goto L89
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v267 = int32(0)
	v269 = F_makeBoolConst(m, v267, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L20
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v406 = v269
	goto L1
L104:
	;
	v274 = v271
	goto L89
L105:
	;
	goto L88
L106:
	;
	if v288 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v294 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L20
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v296 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v406 = v294
	goto L1
L111:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v406 = v300
	goto L1
L112:
	;
	goto L113
L113:
	;
	v301 = F_make_andclause(m, v288)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L20
	} else {
		goto L114
	}
L114:
	;
	v406 = v301
	goto L1
L115:
	;
	v307 = int32(0)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v308 <= v307 {
		v381 = v220
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v313 = int32(0)
	v315 = v307
	goto L117
L117:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321+v313<<(uint(int32(2))%32))))
	if v325 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	if v347 == int32(0) {
		v381 = v220
		goto L7
	} else {
		goto L133
	}
L119:
	;
	v347 = F_lappend(m, v315, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L20
	} else {
		goto L131
	}
L120:
	;
	v344 = F_make_andclause(m, v333)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L20
	} else {
		goto L130
	}
L121:
	;
	v342 = F_list_member(m, v220, v325)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L20
	} else {
		goto L128
	}
L122:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	if v328 != int32(21) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v331 != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	v333 = F_list_difference(m, v332, v220)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	if v333 == int32(0) {
		v381 = v220
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	if v337 != int32(1) {
		goto L120
	} else {
		goto L127
	}
L127:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v346 = v341
	goto L119
L128:
	;
	if v342 != 0 {
		v381 = v220
		goto L7
	} else {
		goto L129
	}
L129:
	;
	v346 = v325
	goto L119
L130:
	;
	v346 = v344
	goto L119
L131:
	;
	v350 = v313 + int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v350 < v351 {
		v313 = v350
		v315 = v347
		goto L117
	} else {
		goto L132
	}
L132:
	;
	goto L118
L133:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v355 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v365 = F_lappend(m, v220, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L20
	} else {
		goto L140
	}
L135:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v347)+12))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v364 = v359
	goto L134
L136:
	;
	goto L137
L137:
	;
	v360 = F_pull_ors(m, v347)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	v362 = F_make_orclause(m, v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L20
	} else {
		goto L139
	}
L139:
	;
	v364 = v362
	goto L134
L140:
	;
	if v365 != 0 {
		v381 = v365
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v395 = int32(0)
	goto L6
L142:
	;
	v406 = v377
	goto L1
L143:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v406 = v392
	goto L1
L144:
	;
	v404 = F_make_andclause(m, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L20
	} else {
		goto L145
	}
L145:
	;
	v406 = v404
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
	switch v11 - int32(280) {
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
	F_errfinish(m, int32(_a_F_find_indexpath_quals_1), int32(2192), int32(_a_F_find_indexpath_quals_2))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v20 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v25)+126)))
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L77
	}
L6:
	;
	if v27 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = F_palloc(m, int32(128))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v251 = v6
	goto L9
L9:
	;
	m.G0 = v17 + int32(80)
	return v251
L10:
	;
	v37 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = v17 + int32(32)
	F_ScanKeyInit(m, v40, int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v47 = int32(1)
	v50 = F_systable_beginscan(m, v37, int32(2187), v47, int32(0), v47, v40)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v52 = F_systable_getnext(m, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v52 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v33
	v62 = v52
	v63 = v6
	v65 = int32(32)
	goto L18
L16:
	;
	v167 = v33
	v169 = v6
	goto L17
L17:
	;
	F_systable_endscan(m, v50)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v70)+12)))
	if v72 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v167 = v155
	v169 = v156
	goto L17
L20:
	;
	v159 = F_systable_getnext(m, v50)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L54
	}
L21:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+22)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135+v136)))
	if v65 <= v63 {
		goto L50
	} else {
		goto L51
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
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v75)
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
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_find_inheritance_children_extended[0]))
	goto L27
L27:
	;
	if base.B2i32(v80 != int32(0)) == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+20)))
	v88 = int32(768)
	if v87&v88 != v88 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v93 = v92
	goto L31
L30:
	;
	v93 = int32(2)
	goto L31
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_find_inheritance_children_extended[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	goto L32
L32:
	;
	v97 = F_XidInMVCCSnapshot(m, v93, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v97 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	if l4 == int32(0) {
		v155 = v61
		v156 = v63
		v157 = v65
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v104 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v93
	v155 = v61
	v156 = v63
	v157 = v65
	goto L20
L39:
	;
	if v104 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_find_inheritance_children_extended_0), v17+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v117))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v93)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_errfinish(m, int32(_a_F_find_inheritance_children_extended_1), int32(167), int32(_a_F_find_inheritance_children_extended_2))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v129 == int32(0) {
		v155 = v61
		v156 = v63
		v157 = v65
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v129 = base.B2i32(base.Ui32(v117) < base.Ui32(v93))
	goto L45
L47:
	;
	goto L48
L48:
	;
	v129 = base.B2i32(int32(0) < v93-v117)
	goto L45
L49:
	;
	goto L38
L50:
	;
	v142 = F_repalloc(m, v61, v65<<(uint(int32(3))%32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v146 = v61
	v147 = v65
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146+v63<<(uint(int32(2))%32)))) = v138
	v155 = v146
	v156 = v63 + int32(1)
	v157 = v147
	goto L20
L53:
	;
	v146 = v142
	v147 = v65 << (uint(int32(1)) % 32)
	goto L52
L54:
	;
	if v159 != 0 {
		v61 = v155
		v62 = v159
		v63 = v156
		v65 = v157
		goto L18
	} else {
		goto L55
	}
L55:
	;
	goto L19
L56:
	;
	F_relation_close(m, v37, int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if int32(2) <= v169 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	F_pfree(m, v167)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L76
	}
L59:
	;
	v190 = int32(0)
	v197 = v190
	v199 = v190
	goto L65
L60:
	;
	F_pg_qsort(m, v167, v169, int32(4), int32(471))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v169 != int32(1) {
		v235 = int32(0)
		goto L58
	} else {
		goto L64
	}
L63:
	;
	goto L59
L64:
	;
	goto L59
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v167+v197<<(uint(int32(2))%32))))
	if l2 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v235 = v224
	goto L58
L67:
	;
	v226 = v197 + int32(1)
	if v226 != v169 {
		v197 = v226
		v199 = v224
		goto L65
	} else {
		goto L75
	}
L68:
	;
	v222 = F_lappend_oid(m, v199, v209)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L74
	}
L69:
	;
	F_LockRelationOid(m, v209, l2)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v215 = int32(0)
	v218 = F_SearchSysCacheExists(m, int32(57), v209, v215, v215, v215)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v218 != 0 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	F_UnlockRelationOid(m, v209, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v224 = v199
	goto L67
L74:
	;
	v224 = v222
	goto L67
L75:
	;
	goto L66
L76:
	;
	v251 = v235
	goto L9
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(_a_F_find_inheritance_children_extended_3), v17)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_find_inheritance_children_extended_1), int32(362), int32(_a_F_find_inheritance_children_extended_4))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	v40 = base.I32_div_s(int32(0), int32(32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v41 <= v40 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v44 = v30 + int32(8)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32))))
	v51 = v48 & int32(-1)
	if v51 != 0 {
		v73 = v51
		v74 = v40
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v53 = v40 + int32(1)
	if v53 == v41 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v56 = v53
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
	v143 = base.I32_div_s(v141, int32(32))
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		v228 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v228
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
		v228 = v22
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
		v219 = v16
		v221 = v21
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
	v228 = int32(0)
	goto L1
L5:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14+v221)))
	if v223 != 0 {
		v14 = v223
		v15 = v219
		v16 = v219
		goto L3
	} else {
		goto L84
	}
L6:
	;
	v219 = v216
	v221 = int32(28)
	goto L5
L7:
	;
	if v23 != int32(319) {
		v228 = v22
		goto L1
	} else {
		goto L63
	}
L8:
	;
	v148 = int32(8)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v149 == int32(2))&v15 != 0 {
		goto L59
	} else {
		goto L60
	}
L9:
	;
	if v15&int32(1) == int32(0) {
		v228 = v22
		goto L1
	} else {
		goto L56
	}
L10:
	;
	if v15&int32(1) == int32(0) {
		v228 = v22
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v72 = int32(8)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	switch v73 {
	case 0:
		goto L35
	case 1:
		v78 = v15
		goto L34
	case 2:
		v219 = int32(0)
		v221 = v72
		goto L5
	default:
		goto L33
	}
L12:
	;
	v69 = F_is_strict_saop(m, v14)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
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
		v228 = v22
		goto L1
	} else {
		goto L24
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v26 <= int32(0) {
		v228 = v22
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
	v228 = v48
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
	v228 = v56
	goto L1
L26:
	;
	if v60 != 0 {
		v216 = int32(0)
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v228 = v22
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
		v216 = int32(0)
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v228 = v22
	goto L1
L31:
	;
	if v69 != 0 {
		v216 = int32(0)
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v228 = v22
	goto L1
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L20
	} else {
		goto L50
	}
L34:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v80 == int32(0) {
		v228 = v22
		goto L1
	} else {
		goto L37
	}
L35:
	;
	v74 = int32(1)
	if v15&v74 != 0 {
		v219 = v74
		v221 = v72
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v78 = int32(0)
	goto L34
L37:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v83 <= int32(0) {
		v228 = v22
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v89 = int32(0)
	v92 = v22
	goto L39
L39:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v89<<(uint(int32(2))%32))))
	v101 = F_find_nonnullable_rels_walker(m, v100, v78&int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L41
	}
L40:
	;
	v228 = int32(0)
	goto L1
L41:
	;
	if v92 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v103 = F_bms_int_members(m, v92, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L20
	} else {
		goto L45
	}
L43:
	;
	v105 = v101
	goto L44
L44:
	;
	if v105 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v105 = v103
	goto L44
L46:
	;
	v107 = v89 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v108 <= v107 {
		v228 = v105
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
	v89 = v107
	v92 = v105
	goto L39
L50:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v115
	F_errmsg_internal(m, int32(_a_F_find_nonnullable_rels_walker_0), v10)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L20
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_find_nonnullable_rels_walker_1), int32(1572), int32(_a_F_find_nonnullable_rels_walker_2))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v129 != int32(1) {
		v228 = v22
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v132 = int32(0)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v133 == v132 {
		v219 = v132
		v221 = v21
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v228 = v22
	goto L1
L56:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(int32(5)) < base.Ui32(v141) {
		v228 = v22
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if int32(1)<<(uint(v141)%32)&int32(37) != 0 {
		v219 = int32(0)
		v221 = v21
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v228 = v22
	goto L1
L59:
	;
	v219 = v15
	v221 = v148
	goto L5
L60:
	;
	goto L61
L61:
	;
	if v149 == int32(3) {
		v219 = v15
		v221 = v148
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v228 = int32(0)
	goto L1
L63:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v161 = F_find_nonnullable_rels_walker(m, v158, v15&int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v163 != 0 {
		v228 = v161
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v165 = int32(0)
	if v164 == v165 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v210 != int32(1) {
		v228 = v161
		goto L1
	} else {
		goto L82
	}
L67:
	;
	v210 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v173 = int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v174 <= v173 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v177 = v173
	goto L72
L71:
	;
	v177 = v174
	goto L72
L72:
	;
	v181 = int32(0)
	v183 = v165
	goto L73
L73:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(8)+v181<<(uint(int32(2))%32))))
	if v190 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v210 = v202
	goto L66
L75:
	;
	goto L74
L76:
	;
	v191 = int32(2)
	if v183 != 0 {
		v202 = v191
		goto L75
	} else {
		goto L79
	}
L77:
	;
	v197 = v183
	goto L78
L78:
	;
	v199 = v181 + int32(1)
	if v199 != v177 {
		v181 = v199
		v183 = v197
		goto L73
	} else {
		goto L81
	}
L79:
	;
	v192 = int32(1)
	if base.Ui32(v192) < base.Ui32(base.I32_popcnt(v190)) {
		v202 = v191
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v197 = v192
	goto L78
L81:
	;
	v202 = v197
	goto L75
L82:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v214 = F_bms_add_members(m, v161, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	v228 = v214
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v7) < base.Ui32(v8) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L8
	} else {
		goto L50
	}
L2:
	;
	return v149
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v7<<(uint(int32(2))%32))))
	if v14 != 0 {
		v149 = v14
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+321)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(324)
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v65) < base.Ui32(v66) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98+v101<<(uint(int32(2))%32)))) = v20
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v110 = F_pull_var_clause(m, v108, int32(26))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L39
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v98 = v68
	v101 = v65
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v98 = v93
	v101 = v96
	goto L22
L33:
	;
	v83 = int32(2)
	v87 = F_repalloc0(m, v82, v66<<(uint(v83)%32), v74<<(uint(v83)%32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v91 = F_palloc0(m, v74<<(uint(int32(2))%32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L37
	}
L36:
	;
	v93 = v87
	goto L32
L37:
	;
	v93 = v91
	goto L32
L38:
	;
	F_list_free(m, v110)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L49
	}
L39:
	;
	if v110 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v114 <= int32(0) {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v119 = int32(0)
	goto L42
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v119<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129 == int32(319) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L38
L44:
	;
	v132 = F_find_placeholder_info(m, l0, v128)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v135 = v119 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v135 < v136 {
		v119 = v135
		goto L42
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L43
L49:
	;
	v149 = v20
	goto L2
L50:
	;
	F_errmsg_internal(m, int32(_a_F_find_placeholder_info_0), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_find_placeholder_info_1), int32(104), int32(_a_F_find_placeholder_info_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
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
	if v66 == int32(319) {
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
	if v102 == int32(319) {
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
	F_errfinish(m, int32(_a_F_find_placeholders_recurse_1), int32(248), int32(_a_F_find_placeholders_recurse_2))
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1235 int32
	_ = v1235
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1483 int32
	_ = v1483
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1666 int32
	_ = v1666
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1785 int32
	_ = v1785
	var v1801 int32
	_ = v1801
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1932 int32
	_ = v1932
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1971 int32
	_ = v1971
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2006 int32
	_ = v2006
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2038 int32
	_ = v2038
	var v2047 int32
	_ = v2047
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2129 int32
	_ = v2129
	var v2137 int32
	_ = v2137
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2262 int32
	_ = v2262
	var v2278 int32
	_ = v2278
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = v25
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v503)+8)) = int32(0)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v508
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v503)+4))
	if v508 < v510 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v130 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v130 = v102
	goto L2
L4:
	;
	if v26 <= v35 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v130 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v42 = int32(1)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v27))))
	if base.Ui32(v44) < base.Ui32(int32(192)) {
		v101 = v44
		v102 = v42
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(246) < v101 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v48 = v35 + int32(1)
	if v48 == v26 {
		v101 = v44
		v102 = v42
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v27))))
	v53 = v51 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v44) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+v27))))
	v69 = v67 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v44) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v57 = v35 + int32(2)
	if v57 != v26 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v101 = v44<<(uint(int32(6))%32)&int32(1984) | v53
	v102 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v73))))
	v101 = v86&int32(63) | (v44<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v53<<(uint(int32(12))%32) | v69<<(uint(int32(6))%32))
	v102 = int32(4)
	goto L9
L18:
	;
	v73 = v35 + int32(3)
	if v73 != v26 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v101 = v44<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v53<<(uint(int32(6))%32) | v69
	v102 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v119 = v102 + v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
	v35 = v119
	goto L4
L23:
	;
	v106 = v101 - int32(97)
	if v106 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v106)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v112)>>(uint(v106&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = v144
	goto L30
L28:
	;
	if v250 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v250 = v221
	goto L28
L30:
	;
	if v145 <= v154 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v250 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v161 = int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v146))))
	if base.Ui32(v163) < base.Ui32(int32(192)) {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(246) < v220 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v167 = v154 + int32(1)
	if v167 == v145 {
		v220 = v163
		v221 = v161
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v146))))
	v172 = v170 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v163) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v146))))
	v188 = v186 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v163) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v176 = v154 + int32(2)
	if v176 != v145 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v220 = v163<<(uint(int32(6))%32)&int32(1984) | v172
	v221 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v192))))
	v220 = v205&int32(63) | (v163<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v172<<(uint(int32(12))%32) | v188<<(uint(int32(6))%32))
	v221 = int32(4)
	goto L35
L44:
	;
	v192 = v154 + int32(3)
	if v192 != v145 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v220 = v163<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v172<<(uint(int32(6))%32) | v188
	v221 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v225 = v220 - int32(97)
	if v225 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v225)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v231)>>(uint(v225&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v239 = v221 + v154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v239
	v154 = v239
	goto L30
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v254 = v253 + v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v254
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v279 = v269
	goto L55
L53:
	;
	if v374 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v374 = v346
	goto L53
L55:
	;
	if v270 <= v279 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v374 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v286 = int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v271))))
	if base.Ui32(v288) < base.Ui32(int32(192)) {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(246) < v345 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v292 = v279 + int32(1)
	if v292 == v270 {
		v345 = v288
		v346 = v286
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v271))))
	v297 = v295 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v288) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v271))))
	v313 = v311 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v288) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v301 = v279 + int32(2)
	if v301 != v270 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v345 = v288<<(uint(int32(6))%32)&int32(1984) | v297
	v346 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271+v317))))
	v345 = v330&int32(63) | (v288<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v297<<(uint(int32(12))%32) | v313<<(uint(int32(6))%32))
	v346 = int32(4)
	goto L60
L69:
	;
	v317 = v279 + int32(3)
	if v317 != v270 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v345 = v288<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v297<<(uint(int32(6))%32) | v313
	v346 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v363 = v346 + v279
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v363
	v279 = v363
	goto L55
L74:
	;
	v350 = v345 - int32(97)
	if v350 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v350)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v356)>>(uint(v350&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v398 = v388
	goto L81
L79:
	;
	if v494 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v494 = v465
	goto L79
L81:
	;
	if v389 <= v398 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v494 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v405 = int32(1)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v390))))
	if base.Ui32(v407) < base.Ui32(int32(192)) {
		v464 = v407
		v465 = v405
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(246) < v464 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v411 = v398 + int32(1)
	if v411 == v389 {
		v464 = v407
		v465 = v405
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v390))))
	v416 = v414 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v407) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420+v390))))
	v432 = v430 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v407) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v420 = v398 + int32(2)
	if v420 != v389 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v464 = v407<<(uint(int32(6))%32)&int32(1984) | v416
	v465 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v436))))
	v464 = v449&int32(63) | (v407<<(uint(int32(18))%32)&int32(_a_F_finnish_UTF_8_stem_0) | v416<<(uint(int32(12))%32) | v432<<(uint(int32(6))%32))
	v465 = int32(4)
	goto L86
L95:
	;
	v436 = v398 + int32(3)
	if v436 != v389 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v464 = v407<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v416<<(uint(int32(6))%32) | v432
	v465 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v469 = v464 - int32(97)
	if v469 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v469)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v475)>>(uint(v469&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v483 = v465 + v398
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483
	v398 = v483
	goto L81
L103:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v498 + v494
	goto L1
L104:
	;
	return v2341
L105:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v668
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if v668 < v671 {
		goto L142
	} else {
		goto L143
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v510
	v516 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_2), int32(10))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	return int32(0)
L108:
	;
	if v516 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v506
	goto L105
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v506
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v524
	switch v516 - int32(1) {
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
	v663 = m.ExcPending
	if v663 != 0 {
		goto L107
	} else {
		goto L140
	}
L113:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	if v524 < v660 {
		goto L105
	} else {
		goto L139
	}
L114:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L117
L115:
	;
	if v656 == int32(0) {
		goto L112
	} else {
		goto L138
	}
L116:
	;
	v656 = v649
	goto L115
L117:
	;
	if v540 <= v541 {
		v649 = int32(-1)
		goto L116
	} else {
		goto L119
	}
L118:
	;
	v649 = int32(0)
	goto L116
L119:
	;
	v558 = int32(1)
	v559 = v540 - v558
	v561 = int32(*(*int8)(unsafe.Add(mBase, uint32(v542+v559))))
	v563 = v561 & int32(255)
	if base.B2i32(v559 == v541)|base.B2i32(int32(0) <= v561) != 0 {
		v621 = v563
		v625 = v558
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if int32(246) < v621 {
		goto L128
	} else {
		goto L129
	}
L121:
	;
	v570 = v563 & int32(63)
	v572 = v540 - int32(2)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+v572))))
	v576 = v574 << (uint(int32(6)) % 32)
	if base.B2i32(v572 != v541)&base.B2i32(base.Ui32(v574) < base.Ui32(int32(192))) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v621 = v576&int32(1984) | v570
	v625 = int32(2)
	goto L120
L123:
	;
	goto L124
L124:
	;
	v589 = v576&int32(4032) | v570
	v591 = v540 - int32(3)
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+v591))))
	if base.B2i32(v591 != v541)&base.B2i32(base.Ui32(v593) < base.Ui32(int32(224))) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v621 = v593<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v589
	v625 = int32(3)
	goto L120
L126:
	;
	goto L127
L127:
	;
	v611 = int32(4)
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540+v542-v611))))
	v621 = v593<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v613&int32(7)<<(uint(int32(18))%32) | v589
	v625 = v611
	goto L120
L128:
	;
	v656 = v625
	goto L115
L129:
	;
	goto L130
L130:
	;
	v627 = v621 - int32(97)
	if v627 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v656 = v625
	goto L115
L132:
	;
	goto L133
L133:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v627)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[1]))))
	if int32(base.Ui32(v633)>>(uint(v627&int32(7))%32))&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v656 = v625
	goto L115
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v540 - v625
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
	if v662 < int32(0) {
		v2341 = v662
		goto L104
	} else {
		goto L141
	}
L141:
	;
	goto L105
L142:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v797
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	if v797 < v800 {
		goto L189
	} else {
		goto L190
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v668
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v671
	v678 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_4), int32(9))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L107
	} else {
		goto L144
	}
L144:
	;
	if v678 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v674
	goto L142
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v674
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v684
	switch v678 - int32(1) {
	case 0:
		goto L153
	case 1:
		goto L152
	case 2:
		goto L151
	case 3:
		goto L150
	case 4:
		goto L149
	case 5:
		goto L148
	default:
		goto L142
	}
L148:
	;
	if v684-int32(2) <= v674 {
		goto L142
	} else {
		goto L183
	}
L149:
	;
	if v684-int32(2) <= v674 {
		goto L142
	} else {
		goto L177
	}
L150:
	;
	v737 = v684 - int32(1)
	if v737 <= v674 {
		goto L142
	} else {
		goto L171
	}
L151:
	;
	v732 = F_slice_del(m, l0)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L107
	} else {
		goto L169
	}
L152:
	;
	v700 = F_slice_del(m, l0)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L107
	} else {
		goto L160
	}
L153:
	;
	if v674 < v684 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689+v684-int32(1)))))
	if v693 == int32(107) {
		goto L142
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v696 = F_slice_del(m, l0)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L107
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	if int32(0) <= v696 {
		goto L142
	} else {
		goto L159
	}
L159:
	;
	v2341 = v696
	goto L104
L160:
	;
	if v700 < int32(0) {
		v2341 = v700
		goto L104
	} else {
		goto L161
	}
L161:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v704
	v706 = int32(3)
	v708 = int32(0)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v704-v711 < v706 {
		v721 = v708
		goto L163
	} else {
		goto L164
	}
L162:
	;
	if v721 == int32(0) {
		goto L142
	} else {
		goto L166
	}
L163:
	;
	goto L162
L164:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v717 = F_memcmp(m, v714+v704-v706, int32(_a_F_finnish_UTF_8_stem_5), v706)
	mBase = m.M
	if v717 != 0 {
		v721 = v708
		goto L163
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v704 - v706
	v721 = int32(1)
	goto L163
L166:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v724
	v728 = F_slice_from_s(m, l0, int32(3), int32(_a_F_finnish_UTF_8_stem_6))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L107
	} else {
		goto L167
	}
L167:
	;
	if int32(0) <= v728 {
		goto L142
	} else {
		goto L168
	}
L168:
	;
	v2341 = v728
	goto L104
L169:
	;
	if int32(0) <= v732 {
		goto L142
	} else {
		goto L170
	}
L170:
	;
	v2341 = v732
	goto L104
L171:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739+v737))))
	if v741 != int32(97) {
		goto L142
	} else {
		goto L172
	}
L172:
	;
	v746 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_7), int32(6))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L107
	} else {
		goto L173
	}
L173:
	;
	if v746 == int32(0) {
		goto L142
	} else {
		goto L174
	}
L174:
	;
	v750 = F_slice_del(m, l0)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L107
	} else {
		goto L175
	}
L175:
	;
	if int32(0) <= v750 {
		goto L142
	} else {
		goto L176
	}
L176:
	;
	v2341 = v750
	goto L104
L177:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757+v684-int32(1)))))
	if v761 != int32(164) {
		goto L142
	} else {
		goto L178
	}
L178:
	;
	v766 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_8), int32(6))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L107
	} else {
		goto L179
	}
L179:
	;
	if v766 == int32(0) {
		goto L142
	} else {
		goto L180
	}
L180:
	;
	v770 = F_slice_del(m, l0)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L107
	} else {
		goto L181
	}
L181:
	;
	if int32(0) <= v770 {
		goto L142
	} else {
		goto L182
	}
L182:
	;
	v2341 = v770
	goto L104
L183:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777+v684-int32(1)))))
	if v781 != int32(101) {
		goto L142
	} else {
		goto L184
	}
L184:
	;
	v786 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_9), int32(2))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L107
	} else {
		goto L185
	}
L185:
	;
	if v786 == int32(0) {
		goto L142
	} else {
		goto L186
	}
L186:
	;
	v790 = F_slice_del(m, l0)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L107
	} else {
		goto L187
	}
L187:
	;
	if v790 < int32(0) {
		v2341 = v790
		goto L104
	} else {
		goto L188
	}
L188:
	;
	goto L142
L189:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1255
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1257)))
	if v1255 < v1258 {
		goto L302
	} else {
		goto L303
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v797
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v800
	v807 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_10), int32(30))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L107
	} else {
		goto L191
	}
L191:
	;
	if v807 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v803
	goto L189
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v803
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v813
	switch v807 - int32(1) {
	case 0:
		goto L203
	case 1:
		goto L202
	case 2:
		goto L201
	case 3:
		goto L200
	case 4:
		goto L199
	case 5:
		goto L198
	case 6:
		goto L197
	case 7:
		goto L196
	default:
		goto L195
	}
L195:
	;
	v1245 = F_slice_del(m, l0)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L107
	} else {
		goto L300
	}
L196:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L254
L197:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v894 = v893 - v813
	v897 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_11), int32(7))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L107
	} else {
		goto L223
	}
L198:
	;
	v877 = int32(2)
	v879 = int32(0)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v881-v882 < v877 {
		v892 = v879
		goto L218
	} else {
		goto L219
	}
L199:
	;
	v861 = int32(2)
	v863 = int32(0)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v865-v866 < v861 {
		v876 = v863
		goto L213
	} else {
		goto L214
	}
L200:
	;
	if v813 <= v803 {
		goto L189
	} else {
		goto L210
	}
L201:
	;
	if v813 <= v803 {
		goto L189
	} else {
		goto L208
	}
L202:
	;
	if v813 <= v803 {
		goto L189
	} else {
		goto L206
	}
L203:
	;
	if v813 <= v803 {
		goto L189
	} else {
		goto L204
	}
L204:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818+v813-int32(1)))))
	if v822 != int32(97) {
		goto L189
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - int32(1)
	goto L195
L206:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v813-int32(1)))))
	if v833 != int32(101) {
		goto L189
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - int32(1)
	goto L195
L208:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v813-int32(1)))))
	if v844 != int32(105) {
		goto L189
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - int32(1)
	goto L195
L210:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851+v813-int32(1)))))
	if v855 != int32(111) {
		goto L189
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - int32(1)
	goto L195
L212:
	;
	if v876 != 0 {
		goto L195
	} else {
		goto L216
	}
L213:
	;
	goto L212
L214:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v872 = F_memcmp(m, v869+v865-v861, int32(_a_F_finnish_UTF_8_stem_12), v861)
	mBase = m.M
	if v872 != 0 {
		v876 = v863
		goto L213
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v865 - v861
	v876 = int32(1)
	goto L213
L216:
	;
	goto L189
L217:
	;
	if v892 != 0 {
		goto L195
	} else {
		goto L221
	}
L218:
	;
	goto L217
L219:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v888 = F_memcmp(m, v885+v881-v877, int32(_a_F_finnish_UTF_8_stem_13), v877)
	mBase = m.M
	if v888 != 0 {
		v892 = v879
		goto L218
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v881 - v877
	v892 = int32(1)
	goto L218
L221:
	;
	goto L189
L222:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v922 = v921 - v894
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v922
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L232
L223:
	;
	if v897 != 0 {
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v900 = v899 - v894
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v900
	v902 = int32(2)
	v904 = int32(0)
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v900-v907 < v902 {
		v917 = v904
		goto L226
	} else {
		goto L227
	}
L225:
	;
	if v917 != 0 {
		goto L222
	} else {
		goto L229
	}
L226:
	;
	goto L225
L227:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v913 = F_memcmp(m, v910+v900-v902, int32(_a_F_finnish_UTF_8_stem_14), v902)
	mBase = m.M
	if v913 != 0 {
		v917 = v904
		goto L226
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v900 - v902
	v917 = int32(1)
	goto L226
L229:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v918 - v894
	goto L195
L230:
	;
	if v977 < int32(0) {
		goto L249
	} else {
		goto L250
	}
L232:
	;
	goto L233
L233:
	;
	goto L234
L234:
	;
	v932 = v922
	v934 = int32(1)
	goto L237
L236:
	;
	v977 = v959
	goto L230
L237:
	;
	if v932 <= v925 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	goto L236
L239:
	;
	v977 = int32(-1)
	goto L230
L240:
	;
	goto L241
L241:
	;
	v939 = v932 - int32(1)
	v941 = int32(*(*int8)(unsafe.Add(mBase, uint32(v924+v939))))
	if base.B2i32(int32(0) <= v941)|base.B2i32(v939 <= v925) != 0 {
		v959 = v939
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v963 = int32(1)
	if v963 < v934 {
		v932 = v959
		v934 = v934 - v963
		goto L237
	} else {
		goto L248
	}
L243:
	;
	v947 = v939
	goto L244
L244:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924+v947))))
	if base.Ui32(int32(191)) < base.Ui32(v952) {
		v959 = v947
		goto L242
	} else {
		goto L246
	}
L245:
	;
	v959 = v925
	goto L242
L246:
	;
	v956 = v947 - int32(1)
	if v925 < v956 {
		v947 = v956
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	goto L238
L249:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v980 - v894
	goto L195
L250:
	;
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977
	goto L195
L252:
	;
	if v1113 != 0 {
		goto L189
	} else {
		goto L275
	}
L253:
	;
	v1113 = v1106
	goto L252
L254:
	;
	if v997 <= v998 {
		v1106 = int32(-1)
		goto L253
	} else {
		goto L256
	}
L255:
	;
	v1106 = int32(0)
	goto L253
L256:
	;
	v1015 = int32(1)
	v1016 = v997 - v1015
	v1018 = int32(*(*int8)(unsafe.Add(mBase, uint32(v999+v1016))))
	v1020 = v1018 & int32(255)
	if base.B2i32(v1016 == v998)|base.B2i32(int32(0) <= v1018) != 0 {
		v1078 = v1020
		v1082 = v1015
		goto L257
	} else {
		goto L258
	}
L257:
	;
	if int32(246) < v1078 {
		goto L265
	} else {
		goto L266
	}
L258:
	;
	v1027 = v1020 & int32(63)
	v1029 = v997 - int32(2)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999+v1029))))
	v1033 = v1031 << (uint(int32(6)) % 32)
	if base.B2i32(v1029 != v998)&base.B2i32(base.Ui32(v1031) < base.Ui32(int32(192))) == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1078 = v1033&int32(1984) | v1027
	v1082 = int32(2)
	goto L257
L260:
	;
	goto L261
L261:
	;
	v1046 = v1033&int32(4032) | v1027
	v1048 = v997 - int32(3)
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999+v1048))))
	if base.B2i32(v1048 != v998)&base.B2i32(base.Ui32(v1050) < base.Ui32(int32(224))) == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1078 = v1050<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1046
	v1082 = int32(3)
	goto L257
L263:
	;
	goto L264
L264:
	;
	v1068 = int32(4)
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v997+v999-v1068))))
	v1078 = v1050<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1070&int32(7)<<(uint(int32(18))%32) | v1046
	v1082 = v1068
	goto L257
L265:
	;
	v1113 = v1082
	goto L252
L266:
	;
	goto L267
L267:
	;
	v1084 = v1078 - int32(97)
	if v1084 < int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1113 = v1082
	goto L252
L269:
	;
	goto L270
L270:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1084)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1090)>>(uint(v1084&int32(7))%32))&int32(1) == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1113 = v1082
	goto L252
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v997 - v1082
	goto L274
L274:
	;
	goto L255
L275:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L278
L276:
	;
	if v1242 != 0 {
		goto L189
	} else {
		goto L299
	}
L277:
	;
	v1242 = v1235
	goto L276
L278:
	;
	if v1126 <= v1127 {
		v1235 = int32(-1)
		goto L277
	} else {
		goto L280
	}
L279:
	;
	v1235 = int32(0)
	goto L277
L280:
	;
	v1144 = int32(1)
	v1145 = v1126 - v1144
	v1147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1128+v1145))))
	v1149 = v1147 & int32(255)
	if base.B2i32(v1145 == v1127)|base.B2i32(int32(0) <= v1147) != 0 {
		v1207 = v1149
		v1211 = v1144
		goto L281
	} else {
		goto L282
	}
L281:
	;
	if int32(122) < v1207 {
		goto L289
	} else {
		goto L290
	}
L282:
	;
	v1156 = v1149 & int32(63)
	v1158 = v1126 - int32(2)
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128+v1158))))
	v1162 = v1160 << (uint(int32(6)) % 32)
	if base.B2i32(v1158 != v1127)&base.B2i32(base.Ui32(v1160) < base.Ui32(int32(192))) == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1207 = v1162&int32(1984) | v1156
	v1211 = int32(2)
	goto L281
L284:
	;
	goto L285
L285:
	;
	v1175 = v1162&int32(4032) | v1156
	v1177 = v1126 - int32(3)
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128+v1177))))
	if base.B2i32(v1177 != v1127)&base.B2i32(base.Ui32(v1179) < base.Ui32(int32(224))) == int32(0) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1207 = v1179<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1175
	v1211 = int32(3)
	goto L281
L287:
	;
	goto L288
L288:
	;
	v1197 = int32(4)
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126+v1128-v1197))))
	v1207 = v1179<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1199&int32(7)<<(uint(int32(18))%32) | v1175
	v1211 = v1197
	goto L281
L289:
	;
	v1242 = v1211
	goto L276
L290:
	;
	goto L291
L291:
	;
	v1213 = v1207 - int32(98)
	if v1213 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1242 = v1211
	goto L276
L293:
	;
	goto L294
L294:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1213)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1219)>>(uint(v1213&int32(7))%32))&int32(1) == int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1242 = v1211
	goto L276
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1126 - v1211
	goto L298
L298:
	;
	goto L279
L299:
	;
	goto L195
L300:
	;
	if v1245 < int32(0) {
		v2341 = v1245
		goto L104
	} else {
		goto L301
	}
L301:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1249)+8)) = int32(1)
	goto L189
L302:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1303
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+8))
	if v1306 != 0 {
		goto L319
	} else {
		goto L320
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1255
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1258
	v1265 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_15), int32(14))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L107
	} else {
		goto L304
	}
L304:
	;
	if v1265 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1261
	goto L302
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1261
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1271
	if v1265 == int32(1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1276 = int32(2)
	v1278 = int32(0)
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1280-v1281 < v1276 {
		v1291 = v1278
		goto L312
	} else {
		goto L313
	}
L309:
	;
	goto L310
L310:
	;
	v1297 = F_slice_del(m, l0)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L107
	} else {
		goto L316
	}
L311:
	;
	if v1291 != 0 {
		goto L302
	} else {
		goto L315
	}
L312:
	;
	goto L311
L313:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1287 = F_memcmp(m, v1284+v1280-v1276, int32(_a_F_finnish_UTF_8_stem_16), v1276)
	mBase = m.M
	if v1287 != 0 {
		v1291 = v1278
		goto L312
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1280 - v1276
	v1291 = int32(1)
	goto L312
L315:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1292 + (v1271 - v1275)
	goto L310
L316:
	;
	if v1297 < int32(0) {
		v2341 = v1297
		goto L104
	} else {
		goto L317
	}
L317:
	;
	goto L302
L318:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1592
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+4))
	if v1592 < v1597 {
		v2335 = int32(0)
		goto L395
	} else {
		goto L396
	}
L319:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+4))
	if v1309 <= v1307 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	goto L321
L321:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1357)+4))
	if v1356 < v1358 {
		v1578 = int32(0)
		goto L342
	} else {
		goto L343
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1307
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1309
	if v1309 < v1307 {
		goto L327
	} else {
		goto L328
	}
L323:
	;
	v1353 = int32(0)
	goto L324
L324:
	;
	if int32(0) <= v1353 {
		goto L318
	} else {
		goto L339
	}
L325:
	;
	v1353 = v1349
	goto L324
L326:
	;
	v1330 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_17), int32(2))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L107
	} else {
		goto L331
	}
L327:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315+v1307-int32(1)))))
	if base.Ui32((v1319-int32(105))&int32(255)) < base.Ui32(int32(2)) {
		goto L326
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1312
	v1349 = int32(0)
	goto L325
L330:
	;
	goto L329
L331:
	;
	if v1330 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1312
	v1349 = int32(0)
	goto L325
L333:
	;
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1312
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1337
	v1340 = F_slice_del(m, l0)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L107
	} else {
		goto L335
	}
L335:
	;
	if int32(0) <= v1340 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1347 = int32(1)
	goto L338
L337:
	;
	v1347 = v1340 >> (uint(int32(31)) % 32) & v1340
	goto L338
L338:
	;
	v1349 = v1347
	goto L325
L339:
	;
	v2341 = v1353
	goto L104
L340:
	;
	if v1585 < int32(0) {
		v2341 = v1585
		goto L104
	} else {
		goto L394
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1361
	v1585 = int32(0)
	goto L340
L342:
	;
	v1585 = v1578
	goto L340
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1356
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1358
	if v1358 < v1356 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1374 = v1356 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1374
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L351
L345:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364+v1356-int32(1)))))
	if v1368 == int32(116) {
		goto L344
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1361
	v1585 = int32(0)
	goto L340
L348:
	;
	goto L347
L349:
	;
	if v1506 != 0 {
		goto L372
	} else {
		goto L373
	}
L350:
	;
	v1506 = v1499
	goto L349
L351:
	;
	if v1374 <= v1391 {
		v1499 = int32(-1)
		goto L350
	} else {
		goto L353
	}
L352:
	;
	v1499 = int32(0)
	goto L350
L353:
	;
	v1408 = int32(1)
	v1409 = v1374 - v1408
	v1411 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1392+v1409))))
	v1413 = v1411 & int32(255)
	if base.B2i32(v1409 == v1391)|base.B2i32(int32(0) <= v1411) != 0 {
		v1471 = v1413
		v1475 = v1408
		goto L354
	} else {
		goto L355
	}
L354:
	;
	if int32(246) < v1471 {
		goto L362
	} else {
		goto L363
	}
L355:
	;
	v1420 = v1413 & int32(63)
	v1422 = v1374 - int32(2)
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1422))))
	v1426 = v1424 << (uint(int32(6)) % 32)
	if base.B2i32(v1422 != v1391)&base.B2i32(base.Ui32(v1424) < base.Ui32(int32(192))) == int32(0) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1471 = v1426&int32(1984) | v1420
	v1475 = int32(2)
	goto L354
L357:
	;
	goto L358
L358:
	;
	v1439 = v1426&int32(4032) | v1420
	v1441 = v1374 - int32(3)
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1441))))
	if base.B2i32(v1441 != v1391)&base.B2i32(base.Ui32(v1443) < base.Ui32(int32(224))) == int32(0) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1471 = v1443<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1439
	v1475 = int32(3)
	goto L354
L360:
	;
	goto L361
L361:
	;
	v1461 = int32(4)
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1374+v1392-v1461))))
	v1471 = v1443<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1463&int32(7)<<(uint(int32(18))%32) | v1439
	v1475 = v1461
	goto L354
L362:
	;
	v1506 = v1475
	goto L349
L363:
	;
	goto L364
L364:
	;
	v1477 = v1471 - int32(97)
	if v1477 < int32(0) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1506 = v1475
	goto L349
L366:
	;
	goto L367
L367:
	;
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1477)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v1483)>>(uint(v1477&int32(7))%32))&int32(1) == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1506 = v1475
	goto L349
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1374 - v1475
	goto L371
L371:
	;
	goto L352
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1361
	v1585 = int32(0)
	goto L340
L373:
	;
	goto L374
L374:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1509 + (v1374 - v1377)
	v1513 = F_slice_del(m, l0)
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L107
	} else {
		goto L375
	}
L375:
	;
	if v1513 < int32(0) {
		v1578 = v1513
		goto L342
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1361
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1520)))
	if v1519 < v1521 {
		v1585 = int32(0)
		goto L340
	} else {
		goto L377
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1519
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1521
	if v1519-int32(2) <= v1521 {
		goto L341
	} else {
		goto L378
	}
L378:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1528+v1519-int32(1)))))
	if v1532 != int32(97) {
		goto L341
	} else {
		goto L379
	}
L379:
	;
	v1537 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_18), int32(2))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L107
	} else {
		goto L380
	}
L380:
	;
	if v1537 == int32(0) {
		goto L341
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1361
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1542
	if v1537 == int32(1) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1547 = int32(0)
	v1548 = int32(2)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1552-v1553 < v1548 {
		v1563 = v1547
		goto L386
	} else {
		goto L387
	}
L383:
	;
	goto L384
L384:
	;
	v1571 = F_slice_del(m, l0)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L107
	} else {
		goto L390
	}
L385:
	;
	if v1563 != 0 {
		v1578 = v1547
		goto L342
	} else {
		goto L389
	}
L386:
	;
	goto L385
L387:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1559 = F_memcmp(m, v1556+v1552-v1548, int32(_a_F_finnish_UTF_8_stem_19), v1548)
	mBase = m.M
	if v1559 != 0 {
		v1563 = v1547
		goto L386
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1552 - v1548
	v1563 = int32(1)
	goto L386
L389:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1564 + (v1542 - v1546)
	goto L384
L390:
	;
	if int32(0) <= v1571 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1575 = int32(1)
	goto L393
L392:
	;
	v1575 = v1571
	goto L393
L393:
	;
	v1578 = v1575
	goto L342
L394:
	;
	goto L318
L395:
	;
	if v2335 < int32(0) {
		v2341 = v2335
		goto L104
	} else {
		goto L550
	}
L396:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1597
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1602 = v1601 - v1592
	v1605 = F_find_among_b(m, l0, int32(_a_F_finnish_UTF_8_stem_11), int32(7))
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L107
	} else {
		goto L399
	}
L397:
	;
	v2335 = v2323
	goto L395
L398:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1677 = v1676 - v1602
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1677
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1677
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L426
L399:
	;
	if v1605 == int32(0) {
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1610 = v1609 - v1602
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1610
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L403
L401:
	;
	if v1666 < int32(0) {
		goto L398
	} else {
		goto L420
	}
L403:
	;
	goto L404
L404:
	;
	goto L405
L405:
	;
	v1621 = v1610
	v1623 = int32(1)
	goto L408
L407:
	;
	v1666 = v1648
	goto L401
L408:
	;
	if v1621 <= v1614 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v1666 = int32(-1)
	goto L401
L411:
	;
	goto L412
L412:
	;
	v1628 = v1621 - int32(1)
	v1630 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1613+v1628))))
	if base.B2i32(int32(0) <= v1630)|base.B2i32(v1628 <= v1614) != 0 {
		v1648 = v1628
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1652 = int32(1)
	if v1652 < v1623 {
		v1621 = v1648
		v1623 = v1623 - v1652
		goto L408
	} else {
		goto L419
	}
L414:
	;
	v1636 = v1628
	goto L415
L415:
	;
	v1641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613+v1636))))
	if base.Ui32(int32(191)) < base.Ui32(v1641) {
		v1648 = v1636
		goto L413
	} else {
		goto L417
	}
L416:
	;
	v1648 = v1614
	goto L413
L417:
	;
	v1645 = v1636 - int32(1)
	if v1614 < v1645 {
		v1636 = v1645
		goto L415
	} else {
		goto L418
	}
L418:
	;
	goto L416
L419:
	;
	goto L409
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1666
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1666
	v1671 = F_slice_del(m, l0)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L107
	} else {
		goto L421
	}
L421:
	;
	if v1671 < int32(0) {
		v2323 = v1671
		goto L397
	} else {
		goto L422
	}
L422:
	;
	goto L398
L423:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1946 = v1945 - v1602
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1946
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1946 <= v1949 {
		v1984 = v1946
		v1985 = v1949
		goto L474
	} else {
		goto L475
	}
L424:
	;
	if v1808 != 0 {
		goto L423
	} else {
		goto L447
	}
L425:
	;
	v1808 = v1801
	goto L424
L426:
	;
	if v1677 <= v1693 {
		v1801 = int32(-1)
		goto L425
	} else {
		goto L428
	}
L427:
	;
	v1801 = int32(0)
	goto L425
L428:
	;
	v1710 = int32(1)
	v1711 = v1677 - v1710
	v1713 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1694+v1711))))
	v1715 = v1713 & int32(255)
	if base.B2i32(v1711 == v1693)|base.B2i32(int32(0) <= v1713) != 0 {
		v1773 = v1715
		v1777 = v1710
		goto L429
	} else {
		goto L430
	}
L429:
	;
	if int32(228) < v1773 {
		goto L437
	} else {
		goto L438
	}
L430:
	;
	v1722 = v1715 & int32(63)
	v1724 = v1677 - int32(2)
	v1726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694+v1724))))
	v1728 = v1726 << (uint(int32(6)) % 32)
	if base.B2i32(v1724 != v1693)&base.B2i32(base.Ui32(v1726) < base.Ui32(int32(192))) == int32(0) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1773 = v1728&int32(1984) | v1722
	v1777 = int32(2)
	goto L429
L432:
	;
	goto L433
L433:
	;
	v1741 = v1728&int32(4032) | v1722
	v1743 = v1677 - int32(3)
	v1745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694+v1743))))
	if base.B2i32(v1743 != v1693)&base.B2i32(base.Ui32(v1745) < base.Ui32(int32(224))) == int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1773 = v1745<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1741
	v1777 = int32(3)
	goto L429
L435:
	;
	goto L436
L436:
	;
	v1763 = int32(4)
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1677+v1694-v1763))))
	v1773 = v1745<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1765&int32(7)<<(uint(int32(18))%32) | v1741
	v1777 = v1763
	goto L429
L437:
	;
	v1808 = v1777
	goto L424
L438:
	;
	goto L439
L439:
	;
	v1779 = v1773 - int32(97)
	if v1779 < int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v1808 = v1777
	goto L424
L441:
	;
	goto L442
L442:
	;
	v1785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1779)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[3]))))
	if int32(base.Ui32(v1785)>>(uint(v1779&int32(7))%32))&int32(1) == int32(0) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1808 = v1777
	goto L424
L444:
	;
	goto L445
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1677 - v1777
	goto L446
L446:
	;
	goto L427
L447:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1809
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L450
L448:
	;
	if v1939 != 0 {
		goto L423
	} else {
		goto L471
	}
L449:
	;
	v1939 = v1932
	goto L448
L450:
	;
	if v1809 <= v1824 {
		v1932 = int32(-1)
		goto L449
	} else {
		goto L452
	}
L451:
	;
	v1932 = int32(0)
	goto L449
L452:
	;
	v1841 = int32(1)
	v1842 = v1809 - v1841
	v1844 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1825+v1842))))
	v1846 = v1844 & int32(255)
	if base.B2i32(v1842 == v1824)|base.B2i32(int32(0) <= v1844) != 0 {
		v1904 = v1846
		v1908 = v1841
		goto L453
	} else {
		goto L454
	}
L453:
	;
	if int32(122) < v1904 {
		goto L461
	} else {
		goto L462
	}
L454:
	;
	v1853 = v1846 & int32(63)
	v1855 = v1809 - int32(2)
	v1857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1825+v1855))))
	v1859 = v1857 << (uint(int32(6)) % 32)
	if base.B2i32(v1855 != v1824)&base.B2i32(base.Ui32(v1857) < base.Ui32(int32(192))) == int32(0) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1904 = v1859&int32(1984) | v1853
	v1908 = int32(2)
	goto L453
L456:
	;
	goto L457
L457:
	;
	v1872 = v1859&int32(4032) | v1853
	v1874 = v1809 - int32(3)
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1825+v1874))))
	if base.B2i32(v1874 != v1824)&base.B2i32(base.Ui32(v1876) < base.Ui32(int32(224))) == int32(0) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v1904 = v1876<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v1872
	v1908 = int32(3)
	goto L453
L459:
	;
	goto L460
L460:
	;
	v1894 = int32(4)
	v1896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1809+v1825-v1894))))
	v1904 = v1876<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v1896&int32(7)<<(uint(int32(18))%32) | v1872
	v1908 = v1894
	goto L453
L461:
	;
	v1939 = v1908
	goto L448
L462:
	;
	goto L463
L463:
	;
	v1910 = v1904 - int32(98)
	if v1910 < int32(0) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1939 = v1908
	goto L448
L465:
	;
	goto L466
L466:
	;
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1910)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v1916)>>(uint(v1910&int32(7))%32))&int32(1) == int32(0) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v1939 = v1908
	goto L448
L468:
	;
	goto L469
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1809 - v1908
	goto L470
L470:
	;
	goto L451
L471:
	;
	v1940 = F_slice_del(m, l0)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L107
	} else {
		goto L472
	}
L472:
	;
	if v1940 < int32(0) {
		v2323 = v1940
		goto L397
	} else {
		goto L473
	}
L473:
	;
	goto L423
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1984
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1984
	if v1984 <= v1985 {
		v2018 = v1984
		goto L484
	} else {
		goto L485
	}
L475:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1952 = v1951 + v1946
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1952-int32(1)))))
	if v1955 != int32(106) {
		v1984 = v1946
		v1985 = v1949
		goto L474
	} else {
		goto L476
	}
L476:
	;
	v1959 = v1946 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1959
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1959
	if v1959 <= v1949 {
		v1984 = v1946
		v1985 = v1949
		goto L474
	} else {
		goto L477
	}
L477:
	;
	v1965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1952-int32(2)))))
	if v1965 != int32(111) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959+v1951-int32(1)))))
	if v1971 != int32(117) {
		v1984 = v1946
		v1985 = v1949
		goto L474
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1946 - int32(2)
	v1977 = F_slice_del(m, l0)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L107
	} else {
		goto L482
	}
L481:
	;
	goto L480
L482:
	;
	if v1977 < int32(0) {
		v2323 = v1977
		goto L397
	} else {
		goto L483
	}
L483:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1984 = v1981 - v1602
	v1985 = v1983
	goto L474
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2018
	v2023 = int32(0)
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2047 = v2018
	goto L493
L485:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1993 = v1992 + v1984
	v1996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993-int32(1)))))
	if v1996 != int32(111) {
		v2018 = v1984
		goto L484
	} else {
		goto L486
	}
L486:
	;
	v2000 = v1984 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2000
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2000
	if v2000 <= v1985 {
		v2018 = v1984
		goto L484
	} else {
		goto L487
	}
L487:
	;
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993-int32(2)))))
	if v2006 != int32(106) {
		v2018 = v1984
		goto L484
	} else {
		goto L488
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1984 - int32(2)
	v2012 = F_slice_del(m, l0)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L107
	} else {
		goto L489
	}
L489:
	;
	if v2012 < int32(0) {
		v2323 = v2012
		goto L397
	} else {
		goto L490
	}
L490:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2018 = v2016 - v1602
	goto L484
L491:
	;
	if v2152 < int32(0) {
		v2323 = v2023
		goto L397
	} else {
		goto L514
	}
L492:
	;
	v2152 = int32(-1)
	goto L491
L493:
	;
	if v2047 <= v1599 {
		goto L492
	} else {
		goto L495
	}
L495:
	;
	v2054 = int32(1)
	v2055 = v2047 - v2054
	v2057 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2038+v2055))))
	v2059 = v2057 & int32(255)
	if base.B2i32(v2055 == v1599)|base.B2i32(int32(0) <= v2057) != 0 {
		v2117 = v2059
		v2121 = v2054
		goto L496
	} else {
		goto L497
	}
L496:
	;
	if int32(246) < v2117 {
		goto L504
	} else {
		goto L505
	}
L497:
	;
	v2066 = v2059 & int32(63)
	v2068 = v2047 - int32(2)
	v2070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2038+v2068))))
	v2072 = v2070 << (uint(int32(6)) % 32)
	if base.B2i32(v2068 != v1599)&base.B2i32(base.Ui32(v2070) < base.Ui32(int32(192))) == int32(0) {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v2117 = v2072&int32(1984) | v2066
	v2121 = int32(2)
	goto L496
L499:
	;
	goto L500
L500:
	;
	v2085 = v2072&int32(4032) | v2066
	v2087 = v2047 - int32(3)
	v2089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2038+v2087))))
	if base.B2i32(v2087 != v1599)&base.B2i32(base.Ui32(v2089) < base.Ui32(int32(224))) == int32(0) {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v2117 = v2089<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v2085
	v2121 = int32(3)
	goto L496
L502:
	;
	goto L503
L503:
	;
	v2107 = int32(4)
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2047+v2038-v2107))))
	v2117 = v2089<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v2109&int32(7)<<(uint(int32(18))%32) | v2085
	v2121 = v2107
	goto L496
L504:
	;
	v2152 = v2121
	goto L491
L505:
	;
	goto L506
L506:
	;
	v2123 = v2117 - int32(97)
	if v2123 < int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2152 = v2121
	goto L491
L508:
	;
	goto L509
L509:
	;
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2123)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[0]))))
	if int32(base.Ui32(v2129)>>(uint(v2123&int32(7))%32))&int32(1) == int32(0) {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2152 = v2121
	goto L491
L511:
	;
	goto L512
L512:
	;
	v2137 = v2047 - v2121
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2137
	v2047 = v2137
	goto L493
L514:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2155
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L517
L515:
	;
	if v2285 != 0 {
		v2323 = v2023
		goto L397
	} else {
		goto L538
	}
L516:
	;
	v2285 = v2278
	goto L515
L517:
	;
	if v2155 <= v2170 {
		v2278 = int32(-1)
		goto L516
	} else {
		goto L519
	}
L518:
	;
	v2278 = int32(0)
	goto L516
L519:
	;
	v2187 = int32(1)
	v2188 = v2155 - v2187
	v2190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2171+v2188))))
	v2192 = v2190 & int32(255)
	if base.B2i32(v2188 == v2170)|base.B2i32(int32(0) <= v2190) != 0 {
		v2250 = v2192
		v2254 = v2187
		goto L520
	} else {
		goto L521
	}
L520:
	;
	if int32(122) < v2250 {
		goto L528
	} else {
		goto L529
	}
L521:
	;
	v2199 = v2192 & int32(63)
	v2201 = v2155 - int32(2)
	v2203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2171+v2201))))
	v2205 = v2203 << (uint(int32(6)) % 32)
	if base.B2i32(v2201 != v2170)&base.B2i32(base.Ui32(v2203) < base.Ui32(int32(192))) == int32(0) {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v2250 = v2205&int32(1984) | v2199
	v2254 = int32(2)
	goto L520
L523:
	;
	goto L524
L524:
	;
	v2218 = v2205&int32(4032) | v2199
	v2220 = v2155 - int32(3)
	v2222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2171+v2220))))
	if base.B2i32(v2220 != v2170)&base.B2i32(base.Ui32(v2222) < base.Ui32(int32(224))) == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2250 = v2222<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_1) | v2218
	v2254 = int32(3)
	goto L520
L526:
	;
	goto L527
L527:
	;
	v2240 = int32(4)
	v2242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2155+v2171-v2240))))
	v2250 = v2222<<(uint(int32(12))%32)&int32(_a_F_finnish_UTF_8_stem_3) | v2242&int32(7)<<(uint(int32(18))%32) | v2218
	v2254 = v2240
	goto L520
L528:
	;
	v2285 = v2254
	goto L515
L529:
	;
	goto L530
L530:
	;
	v2256 = v2250 - int32(98)
	if v2256 < int32(0) {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2285 = v2254
	goto L515
L532:
	;
	goto L533
L533:
	;
	v2262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2256)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_UTF_8_stem[2]))))
	if int32(base.Ui32(v2262)>>(uint(v2256&int32(7))%32))&int32(1) == int32(0) {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v2285 = v2254
	goto L515
L535:
	;
	goto L536
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2155 - v2254
	goto L537
L537:
	;
	goto L518
L538:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2286
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v2288)))
	v2290 = F_slice_to(m, l0, v2289)
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L107
	} else {
		goto L539
	}
L539:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2292))) = v2290
	if v2290 == int32(0) {
		v2335 = int32(-1)
		goto L395
	} else {
		goto L540
	}
L540:
	;
	v2297 = int32(0)
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v2290-int32(4))))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2303-v2304 < v2302 {
		v2314 = v2297
		goto L542
	} else {
		goto L543
	}
L541:
	;
	if v2314 == int32(0) {
		v2323 = v2023
		goto L397
	} else {
		goto L545
	}
L542:
	;
	goto L541
L543:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2310 = F_memcmp(m, v2307+v2303-v2302, v2290, v2302)
	mBase = m.M
	if v2310 != 0 {
		v2314 = v2297
		goto L542
	} else {
		goto L544
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2303 - v2302
	v2314 = int32(1)
	goto L542
L545:
	;
	v2318 = F_slice_del(m, l0)
	mBase = m.M
	v2319 = m.ExcPending
	if v2319 != 0 {
		goto L107
	} else {
		goto L546
	}
L546:
	;
	if int32(0) <= v2318 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2322 = int32(1)
	goto L549
L548:
	;
	v2322 = v2318
	goto L549
L549:
	;
	v2323 = v2322
	goto L397
L550:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2338
	v2341 = int32(1)
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
				v24 = F_expression_tree_walker_impl(m, l0, int32(1042), l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					return v24
				}
			}
		} else {
			v24 = F_expression_tree_walker_impl(m, l0, int32(1042), l1)
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
				v34 = F_expression_tree_walker_impl(m, l0, int32(834), l1)
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
			v29 = F_query_tree_walker_impl(m, l0, int32(834), l1, int32(16))
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
func F_float48lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float48mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = base.F64_promote_f32(v5)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_sub(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float4mi(m *base.Module, l0 int32) int32 {
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
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = base.F32_sub(v5, v6)
	v9 = math.Float32frombits(uint32(0x7f800000))
	if base.F32_ne(base.F32_abs(v7), v9)|base.F32_eq(base.F32_abs(v5), v9)|base.F32_eq(base.F32_abs(v6), v9) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I32_reinterpret_f32(v7)
	}
}
func F_float4out(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_float4out[0]))
		if int32(0) < v11 {
			v14 = F_float_to_shortest_decimal_bufn(m, v4, v6)
			mBase = m.M
			v16 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14+v6))) = uint8(v16)
			return v6
		} else {
			F_pg_strfromd(m, v6, v11+int32(6), base.F64_promote_f32(v4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F_float4smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	var v16 float32
	_ = v16
	var v18 float32
	_ = v18
	var v19 float32
	_ = v19
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(base.I32_reinterpret_f32(v5)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) {
			v16 = v5
		} else {
			v16 = v4
		}
		if base.F32_gt(v4, v5) != 0 {
			v18 = v5
		} else {
			v18 = v16
		}
		v19 = v18
	} else {
		v19 = v4
	}
	return base.I32_reinterpret_f32(v19)
}
func F_float84mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = base.F64_promote_f32(v7)
	v9 = base.F64_sub(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float8eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9))
	} else {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v6, v11)
	}
}
func F_float8ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_float8lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float8mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v17 int32
	_ = v17
	var v25 float64
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_mul(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	v17 = int32(0)
	if base.B2i32(base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11) == v17)&base.F64_ne(base.F64_abs(v8), v11) == v17 {
		v25 = float64(0)
		if base.B2i32(base.F64_eq(v6, v25)|base.F64_ne(v9, v25) == int32(0))&base.F64_ne(v8, v25) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v35 = F_Float8GetDatum(m, v9)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				return v35
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
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
func F_float8send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		F_pq_sendfloat8(m, v6, v9)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v17
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
