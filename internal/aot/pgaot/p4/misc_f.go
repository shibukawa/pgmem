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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
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
			if v21 == int32(3831) {
				F_GetOperatorFromCompareType(m, l0, int32(0), int32(8), l1, v9+int32(22))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_GetOperatorFromCompareType(m, l0, int32(4537), int32(8), l2, v9+int32(22))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
						if v74 != int32(3831) {
							if v74 != int32(4537) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v89
									F_errmsg_internal(m, int32(58531), v9)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_errfinish(m, int32(492793), int32(1720), int32(133613))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v80 = int32(4394)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
								m.G0 = v9 + int32(32)
								return
							}
						} else {
							v80 = int32(3900)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			} else {
				if v21 == int32(4537) {
					F_GetOperatorFromCompareType(m, l0, int32(0), int32(8), l1, v9+int32(22))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_GetOperatorFromCompareType(m, l0, int32(4537), int32(8), l2, v9+int32(22))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
							if v74 != int32(3831) {
								if v74 != int32(4537) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v89
										F_errmsg_internal(m, int32(58531), v9)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											F_errfinish(m, int32(492793), int32(1720), int32(133613))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v80 = int32(4394)
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
									m.G0 = v9 + int32(32)
									return
								}
							} else {
								v80 = int32(3900)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v80
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_errmsg(m, int32(22015), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								F_errdetail(m, int32(647393), int32(0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									F_errfinish(m, int32(492793), int32(1682), int32(133613))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
				F_errmsg_internal(m, int32(42264), v9+int32(16))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(492793), int32(1686), int32(133613))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
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
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v237 int32
	_ = v237
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
	var v253 int32
	_ = v253
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
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
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
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v593 int32
	_ = v593
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
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
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
	var v669 int32
	_ = v669
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v811 int32
	_ = v811
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[169])))
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
	v44 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[169])) = uint8(v41)
	goto L3
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[170]))
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
	v896 = m.ExcPending
	if v896 != 0 {
		goto L4
	} else {
		goto L179
	}
L8:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L176
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L171
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L167
	}
L11:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v210
	*(*int32)(unsafe.Add(mBase, _consts[167])) = v86
	v215 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v215+int32(2304))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L46
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
	v186 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v186+int32(2304))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L41
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
	v158 = v60 + int32(1)
	if v158 != v52 {
		v60 = v158
		goto L15
	} else {
		goto L40
	}
L18:
	;
	v91 = v86 + int32(47)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 == int32(0) {
		v114 = v94
		v115 = v95
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v115-v114 != 0 {
		goto L17
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	if v94 != v95 {
		v114 = v94
		v115 = v95
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v99 = v91
	v100 = l0
	goto L23
L23:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v104 == int32(0) {
		v114 = v103
		v115 = v104
		goto L20
	} else {
		goto L25
	}
L24:
	;
	v114 = v103
	v115 = v104
	goto L20
L25:
	;
	v107 = int32(1)
	if v103 == v104 {
		v99 = v99 + v107
		v100 = v100 + v107
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v86)+40))
	if v117 != int32(-1) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v122 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	if v124 != v31 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v126 = F_superuser_arg(m, v31)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123+v120*int32(640))+60))
	if v131 == v135 {
		goto L11
	} else {
		goto L34
	}
L32:
	;
	if v126 == int32(0) {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(361645), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errhint(m, int32(581166), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(498441), int32(599), int32(111835))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	goto L16
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = l0
	F_errmsg(m, int32(71047), v26+int32(-16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(498441), int32(615), int32(111835))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v222 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v86)+32))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+45)))
	if v225 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+48))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v239)+44))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)+32))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)+28))
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+54)))
	v253 = v239 + (v246+int32(7))&int32(131064) + int32(72)
	v258 = v245 - int32(1)
	if v258 < int32(0) {
		v325 = v224
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v229 = F_ReadTwoPhaseFile(m, v224, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v86)+16))
	F_XlogReadTwoPhaseData(m, v232, v26+int32(-4), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v229
	v239 = v229
	goto L47
L52:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v239 = v238
	goto L47
L53:
	;
	v331 = int32(4509772)
	v333 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v333 + int32(1)
	v339 = int32(7)
	v341 = int32(-8)
	v343 = v253 + (v245<<(uint(int32(2))%32)+v339)&v341
	v344 = int32(12)
	v350 = v343 + (v244*v344+v339)&v341
	v357 = v350 + (v243*v344+v339)&v341
	v358 = int32(4)
	v360 = v357 + v242<<(uint(v358)%32)
	v363 = v360 + v241<<(uint(v358)%32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v239)+28))
	if l1 != 0 {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	goto L53
L55:
	;
	if v245&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v263 = int32(2)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v253+v258<<(uint(v263)%32))))
	if base.B2i32(base.Ui32(v263) < base.Ui32(v266))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v224)) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v281 = v224
	v283 = v258
	goto L58
L58:
	;
	if v258 == int32(0) {
		v325 = v281
		goto L54
	} else {
		goto L66
	}
L59:
	;
	v281 = v278
	v283 = v245 - int32(2)
	goto L58
L60:
	;
	v278 = v266
	goto L59
L61:
	;
	if base.Ui32(v224) < base.Ui32(v266) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if int32(0) <= v224-v266 {
		v278 = v224
		goto L59
	} else {
		goto L65
	}
L64:
	;
	v278 = v224
	goto L59
L65:
	;
	goto L60
L66:
	;
	v288 = v281
	v289 = v283
	goto L67
L67:
	;
	v295 = v289 << (uint(int32(2)) % 32)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v253+v295)))
	if base.Ui32(v288) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v325 = v320
	goto L54
L69:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v295+(v253-int32(4)))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v308))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v306)) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v306 = v297
	goto L69
L71:
	;
	if base.Ui32(v297) <= base.Ui32(v288) {
		v306 = v288
		goto L69
	} else {
		goto L75
	}
L72:
	;
	if base.Ui32(v297) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v288-v297 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v306 = v288
	goto L69
L75:
	;
	goto L70
L76:
	;
	if int32(1) < v289 {
		v288 = v320
		v289 = v289 - int32(2)
		goto L67
	} else {
		goto L83
	}
L77:
	;
	v320 = v308
	goto L76
L78:
	;
	if base.Ui32(v306) < base.Ui32(v308) {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if int32(0) <= v306-v308 {
		v320 = v306
		goto L76
	} else {
		goto L82
	}
L81:
	;
	v320 = v306
	goto L76
L82:
	;
	goto L77
L83:
	;
	goto L68
L84:
	;
	v529 = v363 + v240<<(uint(int32(4))%32)
	F_ProcArrayRemove(m, v223+v220*int32(640), v325)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L111
	}
L85:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+52)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v239)+48))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v239)+32))
	v372 = m.G0
	v373 = int32(16)
	v374 = v372 - v373
	m.G0 = v374
	F___gettimeofday(m, v374)
	mBase = m.M
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v374)))
	v378 = int64(*(*int32)(unsafe.Add(mBase, uint32(v374)+8)))
	m.G0 = v374 + v373
	v386 = v378 + v377*int64(1000000) - int64(946684800000000)
	goto L88
L86:
	;
	goto L87
L87:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, _consts[172])))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v239)+44))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v239)+36))
	v456 = F_TransactionIdDidCommit(m, v224)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L100
	}
L88:
	;
	v387 = int32(4509780)
	v389 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v390 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v389 + v390
	v394 = int32(*(*uint16)(unsafe.Add(mBase, _consts[172])))
	v396 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+120)) = v397 | v390
	v402 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v405 = F_XactLogCommitRecord(m, v386, v364, v253, v368, v343, v367, v357, v366, v363, v365, v402|int32(2), v224, l0)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v394+int32(1))&int32(65535)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, _consts[172])))
	F_TransactionTreeSetCommitTsData(m, v224, v364, v253, v426, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L96
	}
L91:
	;
	v414 = *(*int64)(unsafe.Add(mBase, _consts[174]))
	v416 = *(*int64)(unsafe.Add(mBase, _consts[175]))
	F_replorigin_session_advance(m, v414, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	*(*int64)(unsafe.Add(mBase, _consts[176])) = v386
	v426 = v386
	goto L90
L94:
	;
	v420 = *(*int64)(unsafe.Add(mBase, _consts[176]))
	if v420 != int64(0) {
		v426 = v420
		goto L90
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	F_XLogFlush(m, v405)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	F_TransactionIdCommitTree(m, v224, v364, v253)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+120)) = v439 & int32(-2)
	v443 = int32(4509780)
	v445 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v446 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v445 - v446
	F_SyncRepWaitForLSN(m, v405, v446)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v515 = v239 + int32(32)
	v517 = v343
	goto L84
L100:
	;
	if v456 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v458 = int32(4509780)
	v460 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v460 + int32(1)
	v467 = m.G0
	v468 = int32(16)
	v469 = v467 - v468
	m.G0 = v469
	F___gettimeofday(m, v469)
	mBase = m.M
	v472 = *(*int64)(unsafe.Add(mBase, uint32(v469)))
	v473 = int64(*(*int32)(unsafe.Add(mBase, uint32(v469)+8)))
	m.G0 = v469 + v468
	goto L102
L102:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	v486 = F_XactLogAbortRecord(m, v473+v472*int64(1000000)-int64(946684800000000), v364, v253, v455, v350, v454, v360, v483|int32(2), v224, l0)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	if base.Ui32((v453-int32(1))&int32(65535)) <= base.Ui32(int32(65533)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v495 = *(*int64)(unsafe.Add(mBase, _consts[174]))
	v497 = *(*int64)(unsafe.Add(mBase, _consts[175]))
	F_replorigin_session_advance(m, v495, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_XLogFlush(m, v486)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	F_TransactionIdAbortTree(m, v224, v364, v253)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v506 = int32(4509780)
	v508 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v508 - int32(1)
	F_SyncRepWaitForLSN(m, v486, int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v515 = v239 + int32(36)
	v517 = v350
	goto L84
L111:
	;
	v535 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+44)) = uint8(v535)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	F_DropRelationFiles(m, v517, v537, v535)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	if l1 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v695 = m.G0
	v697 = v695 - int32(16)
	m.G0 = v697
	*(*int32)(unsafe.Add(mBase, uint32(v697)+12)) = v224
	v701 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v705 = F_LWLockAcquire(m, v701+int32(3584), int32(1))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L4
	} else {
		goto L146
	}
L114:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v239)+40))
	F_pgstat_execute_transactional_drops(m, v541, v357)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v239)+44))
	F_pgstat_execute_transactional_drops(m, v612, v360)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L136
	}
L117:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+52)))
	if v544 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v239)+48))
	F_SendSharedInvalidMessages(m, v363, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+52)))
	if v552 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v562 = F_LWLockAcquire(m, v558+int32(2304), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+4)))
	if v564 == int32(0) {
		goto L113
	} else {
		goto L128
	}
L128:
	;
	v569 = v529
	v570 = v564
	goto L129
L129:
	;
	v593 = v569 + int32(8)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v570&int32(255)<<(uint(int32(2))%32))+uint32(_consts[177])))
	if v600 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L113
L131:
	;
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v569)+6)))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	m.T0[v600].(func(*base.Module, int32, int32, int32, int32))(m, v224, v601, v593, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v569)))
	v610 = v593 + (v605+int32(7))&int32(-8)
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610)+4)))
	if v611 != 0 {
		v569 = v610
		v570 = v611
		goto L129
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	goto L130
L136:
	;
	v616 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v620 = F_LWLockAcquire(m, v616+int32(2304), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+4)))
	if v622 == int32(0) {
		goto L113
	} else {
		goto L138
	}
L138:
	;
	v627 = v529
	v628 = v622
	goto L139
L139:
	;
	v651 = v627 + int32(8)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v628&int32(255)<<(uint(int32(2))%32))+uint32(_consts[178])))
	if v658 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L113
L141:
	;
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v627)+6)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	m.T0[v658].(func(*base.Module, int32, int32, int32, int32))(m, v224, v659, v651, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	v668 = v651 + (v663+int32(7))&int32(-8)
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+4)))
	if v669 != 0 {
		v627 = v668
		v628 = v669
		goto L139
	} else {
		goto L145
	}
L144:
	;
	goto L143
L145:
	;
	goto L140
L146:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v711 = int32(0)
	v713 = F_hash_search(m, v708, v697+int32(12), v711, v711)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v716+int32(3584))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	if v713 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	v723 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[179])) = uint8(v723)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = v721
	F_ReleasePredicateLocks(m, l1, int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	m.G0 = v697 + int32(16)
	v735 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	if v736 <= int32(0) {
		goto L7
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+45)))
	v741 = v735 + int32(8)
	v745 = int32(0)
	goto L154
L154:
	;
	v770 = v741 + v745<<(uint(int32(2))%32)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	if v771 != v86 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v777 = v736 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v735)+4)) = v777
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v741+v777<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v770))) = v782
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v784
	*(*int32)(unsafe.Add(mBase, uint32(v735))) = v86
	v788 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v788+int32(2304))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L160
	}
L156:
	;
	v774 = v745 + int32(1)
	if v736 != v774 {
		v745 = v774
		goto L154
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	goto L7
L160:
	;
	F_AtEOXact_PgStat(m, l1, int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v739&int32(1) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	F_RemoveTwoPhaseFile(m, v224, int32(1))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v801 = int32(4509772)
	v803 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v803 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[167])) = int32(0)
	F_pfree(m, v239)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	m.G0 = v28 - int32(-64)
	return
L167:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l0
	F_errmsg(m, int32(11896), v26+int32(-32))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(498441), int32(581), int32(111835))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(257207), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_errhint(m, int32(614020), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(498441), int32(587), int32(111835))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v224
	F_errmsg_internal(m, int32(440521), v26+int32(-48))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(498441), int32(2419), int32(451038))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v86
	F_errmsg_internal(m, int32(25149), v28)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(498441), int32(650), int32(111845))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	*(*int32)(unsafe.Add(mBase, _consts[117])) = int32(0)
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
		v11 = *(*int64)(unsafe.Add(mBase, _consts[958]))
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
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(8)
		return int32(-1)
	} else {
		v17 = F___fstatat(m, l0, int32(756936), l1, int32(4096))
		mBase = m.M
		return v17
	}
}
func F_fdw_handler_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(219165)
			F_errmsg(m, int32(192350), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493865), int32(369), int32(279132))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
	var v120 int32
	_ = v120
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
	return v120
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
		v120 = v31
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
	v120 = v106
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
	var v118 int32
	_ = v118
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
	var v214 int32
	_ = v214
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
	v117 = v4
	v118 = v111
	goto L30
L28:
	;
	v214 = v4
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v214
	m.G0 = v11 + int32(16)
	return v51
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v118<<(uint(int32(2))%32))))
	if v126 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v214 = v152
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
	v129 = F_find_base_rel_ignore_join(m, l0, v118)
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
	*(*int32)(unsafe.Add(mBase, uint32(v51+v117<<(uint(int32(2))%32)))) = v126
	v152 = v117 + int32(1)
	goto L32
L36:
	;
	if v129 == int32(0) {
		v152 = v117
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v118
	F_errmsg_internal(m, int32(24429), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(495896), int32(778), int32(173223))
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
		v117 = v152
		v118 = v208
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
	v159 = v118 + int32(1)
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v4 - int32(6) {
		case 0:
			v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v8 == int32(1) {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v12 = F_bms_add_member(m, v11, v7)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v12
					return int32(0)
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v20 = F_bms_add_member(m, v19, v7)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v20
					return int32(0)
				}
			}
		default:
			v35 = F_expression_tree_walker_impl(m, l0, int32(693), l1)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = v35
				return v37
			}
		case 3:
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v25)
			v28 = F_expression_tree_walker_impl(m, l0, int32(693), l1)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v30)
				return v30
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
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
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v427
L2:
	;
	v427 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != int32(21) {
		v427 = l0
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v21 {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		v427 = l0
		goto L1
	}
L6:
	;
	v425 = F_make_orclause(m, v87)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L22
	} else {
		goto L152
	}
L7:
	;
	if v231 == int32(0) {
		goto L6
	} else {
		goto L119
	}
L8:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v427 = v315
	goto L1
L9:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v240 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v87 = F_pull_ors(m, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L22
	} else {
		goto L38
	}
L12:
	;
	v80 = v3
	goto L11
L13:
	;
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v80 = v3
	goto L11
L16:
	;
	goto L17
L17:
	;
	v31 = v3
	v32 = v3
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v32<<(uint(int32(2))%32))))
	v43 = F_find_duplicate_ors(m, v42, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v80 = v71
	goto L11
L20:
	;
	v74 = v32 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v74 < v75 {
		v31 = v71
		v32 = v74
		goto L18
	} else {
		goto L37
	}
L21:
	;
	v69 = F_lappend(m, v31, v43)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L22
	} else {
		goto L36
	}
L22:
	;
	return int32(0)
L23:
	;
	if v43 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v49 != int32(7) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
	if l1 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v52&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v52&int32(1) != 0 {
		v71 = v31
		goto L20
	} else {
		goto L34
	}
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v57 == int32(0) {
		v71 = v31
		goto L20
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v62 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L22
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v427 = v62
	goto L1
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v66 == int32(0) {
		v71 = v31
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v427 = v43
	goto L1
L36:
	;
	v71 = v69
	goto L20
L37:
	;
	goto L19
L38:
	;
	if v87 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v91 = int32(0)
	v93 = F_makeBoolConst(m, v91, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v95 == int32(1) {
		goto L8
	} else {
		goto L43
	}
L42:
	;
	v427 = v93
	goto L1
L43:
	;
	v98 = int32(0)
	if v95 <= v98 {
		v145 = v98
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v154 = F_list_union(m, v145)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L22
	} else {
		goto L63
	}
L45:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v102 = int32(0)
	v105 = v98
	v106 = v102
	v107 = v102
	goto L46
L46:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v101+v106<<(uint(int32(2))%32))))
	if v117 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v145 = v139
	goto L44
L48:
	;
	v138 = base.B2i32(v105 == int32(0)) | base.B2i32(v134 < v107)
	if v138 != 0 {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v134 = v133
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v117
	v131 = F_list_make1_impl(m, int32(1), v13+int32(8))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L22
	} else {
		goto L55
	}
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v120 != int32(21) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v123 != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v124 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v134 = int32(0)
	goto L48
L55:
	;
	v145 = v131
	goto L44
L56:
	;
	v139 = v124
	goto L58
L57:
	;
	v139 = v105
	goto L58
L58:
	;
	if v138 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v140 = v134
	goto L61
L60:
	;
	v140 = v107
	goto L61
L61:
	;
	v142 = v106 + int32(1)
	if v142 != v95 {
		v105 = v139
		v106 = v142
		v107 = v140
		goto L46
	} else {
		goto L62
	}
L62:
	;
	goto L47
L63:
	;
	if v154 == int32(0) {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v158 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v159 <= v158 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v166 = v158
	v168 = int32(0)
	goto L66
L66:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v166<<(uint(int32(2))%32))))
	v178 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v178 < v179 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L7
L68:
	;
	v237 = v166 + int32(1)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v237 < v238 {
		v166 = v237
		v168 = v231
		goto L66
	} else {
		goto L85
	}
L69:
	;
	v182 = v178
	goto L72
L70:
	;
	goto L71
L71:
	;
	v224 = F_lappend(m, v168, v177)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L22
	} else {
		goto L84
	}
L72:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v182<<(uint(int32(2))%32))))
	if v196 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L71
L74:
	;
	v211 = v182 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v211 < v212 {
		v182 = v211
		goto L72
	} else {
		goto L83
	}
L75:
	;
	v206 = F_equal(m, v177, v196)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L22
	} else {
		goto L81
	}
L76:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v199 != int32(21) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v202 != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v204 = F_list_member(m, v203, v177)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L22
	} else {
		goto L79
	}
L79:
	;
	if v204 != 0 {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v231 = v168
	goto L68
L81:
	;
	if v206 == int32(0) {
		v231 = v168
		goto L68
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	goto L73
L84:
	;
	v231 = v224
	goto L68
L85:
	;
	goto L67
L86:
	;
	v299 = F_pull_ands(m, v292)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L22
	} else {
		goto L110
	}
L87:
	;
	v292 = v3
	goto L86
L88:
	;
	goto L89
L89:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v243 <= int32(0) {
		v292 = v3
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v249 = v3
	v250 = v3
	goto L91
L91:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v250<<(uint(int32(2))%32))))
	v261 = F_find_duplicate_ors(m, v260, l1)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L22
	} else {
		goto L95
	}
L92:
	;
	v292 = v283
	goto L86
L93:
	;
	v286 = v250 + int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v286 < v287 {
		v249 = v283
		v250 = v286
		goto L91
	} else {
		goto L109
	}
L94:
	;
	v281 = F_lappend(m, v249, v261)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L22
	} else {
		goto L108
	}
L95:
	;
	if v261 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v265 != int32(7) {
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+24)))
	if l1 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v268&int32(1) != 0 {
		v283 = v249
		goto L93
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v268&int32(1) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	if v271 != 0 {
		v283 = v249
		goto L93
	} else {
		goto L102
	}
L102:
	;
	v427 = v261
	goto L1
L103:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	if v276 != 0 {
		v283 = v249
		goto L93
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v277 = int32(0)
	v279 = F_makeBoolConst(m, v277, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L22
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v427 = v279
	goto L1
L108:
	;
	v283 = v281
	goto L93
L109:
	;
	goto L92
L110:
	;
	if v299 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v305 = F_makeBoolConst(m, int32(1), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L22
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if v307 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v427 = v305
	goto L1
L115:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v427 = v311
	goto L1
L116:
	;
	goto L117
L117:
	;
	v312 = F_make_andclause(m, v299)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L22
	} else {
		goto L118
	}
L118:
	;
	v427 = v312
	goto L1
L119:
	;
	v318 = int32(0)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v319 <= v318 {
		v391 = v231
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v411 = F_pull_ands(m, v406)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L22
	} else {
		goto L150
	}
L121:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v396 != int32(1) {
		v406 = v391
		goto L120
	} else {
		goto L149
	}
L122:
	;
	v324 = v318
	v325 = int32(0)
	goto L124
L123:
	;
	if v384 != 0 {
		v391 = v384
		goto L121
	} else {
		goto L148
	}
L124:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333+v324<<(uint(int32(2))%32))))
	if v337 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	if v361 == int32(0) {
		v384 = v231
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v361 = F_lappend(m, v325, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L22
	} else {
		goto L138
	}
L127:
	;
	v358 = F_make_andclause(m, v345)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L22
	} else {
		goto L137
	}
L128:
	;
	v354 = F_list_member(m, v231, v337)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L22
	} else {
		goto L135
	}
L129:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	if v340 != int32(21) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v343 != 0 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v337)+8))
	v345 = F_list_difference(m, v344, v231)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L22
	} else {
		goto L132
	}
L132:
	;
	if v345 == int32(0) {
		v384 = v231
		goto L123
	} else {
		goto L133
	}
L133:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v349 != int32(1) {
		goto L127
	} else {
		goto L134
	}
L134:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	v360 = v353
	goto L126
L135:
	;
	if v354 == int32(0) {
		v360 = v337
		goto L126
	} else {
		goto L136
	}
L136:
	;
	v384 = v231
	goto L123
L137:
	;
	v360 = v358
	goto L126
L138:
	;
	v364 = v324 + int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v364 < v365 {
		v324 = v364
		v325 = v361
		goto L124
	} else {
		goto L139
	}
L139:
	;
	goto L125
L140:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	if v369 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v379 = F_lappend(m, v231, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L22
	} else {
		goto L147
	}
L142:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v378 = v373
	goto L141
L143:
	;
	goto L144
L144:
	;
	v374 = F_pull_ors(m, v361)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L22
	} else {
		goto L145
	}
L145:
	;
	v376 = F_make_orclause(m, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L22
	} else {
		goto L146
	}
L146:
	;
	v378 = v376
	goto L141
L147:
	;
	v384 = v379
	goto L123
L148:
	;
	v406 = int32(0)
	goto L120
L149:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v427 = v400
	goto L1
L150:
	;
	v413 = F_make_andclause(m, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L22
	} else {
		goto L151
	}
L151:
	;
	v427 = v413
	goto L1
L152:
	;
	v427 = v425
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
	var v24 int32
	_ = v24
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
	var v47 int32
	_ = v47
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
	var v70 int32
	_ = v70
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
	v24 = v4
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
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
	v34 = v24 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v34 < v35 {
		v24 = v34
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
	v47 = v4
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v47<<(uint(int32(2))%32))))
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
	v57 = v47 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v57 < v58 {
		v47 = v57
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
	v70 = v4
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v70<<(uint(int32(2))%32))))
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
	v84 = v70 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v84 < v85 {
		v70 = v84
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
	F_errmsg_internal(m, int32(485725), v9)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(497533), int32(2192), int32(153331))
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
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
	v268 = m.ExcPending
	if v268 != 0 {
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
	v257 = v6
	goto L9
L9:
	;
	m.G0 = v17 + int32(80)
	return v257
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
	F_ScanKeyInit(m, v17+int32(32), int32(2), int32(3), int32(184), l0)
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
	v52 = F_systable_beginscan(m, v37, int32(2187), v47, int32(0), v47, v17+int32(32))
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
	v65 = v6
	v66 = v33
	v68 = int32(32)
	goto L18
L16:
	;
	v172 = v6
	v173 = v33
	goto L17
L17:
	;
	F_systable_endscan(m, v52)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L56
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v72)+12)))
	if v74 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v172 = v158
	v173 = v159
	goto L17
L20:
	;
	v162 = F_systable_getnext(m, v52)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L54
	}
L21:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+22)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137+v138)))
	if v68 <= v65 {
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
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v77)
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
	v82 = *(*int32)(unsafe.Add(mBase, _consts[263]))
	goto L27
L27:
	;
	if base.B2i32(v82 != int32(0)) == int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88)+20)))
	v90 = int32(768)
	if v89&v90 != v90 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v95 = v94
	goto L31
L30:
	;
	v95 = int32(2)
	goto L31
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[263]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	goto L32
L32:
	;
	v99 = F_XidInMVCCSnapshot(m, v95, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v99 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	if l4 == int32(0) {
		v158 = v65
		v159 = v66
		v161 = v68
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v103 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v106 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v95
	v158 = v65
	v159 = v66
	v161 = v68
	goto L20
L39:
	;
	if v106 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l0
	F_errmsg_internal(m, int32(56527), v17+int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v119))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v95)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	F_errfinish(m, int32(493356), int32(167), int32(460391))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v131 == int32(0) {
		v158 = v65
		v159 = v66
		v161 = v68
		goto L20
	} else {
		goto L49
	}
L46:
	;
	v131 = base.B2i32(base.Ui32(v119) < base.Ui32(v95))
	goto L45
L47:
	;
	goto L48
L48:
	;
	v131 = base.B2i32(int32(0) < v95-v119)
	goto L45
L49:
	;
	goto L38
L50:
	;
	v144 = F_repalloc(m, v66, v68<<(uint(int32(3))%32))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v148 = v66
	v149 = v68
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v65<<(uint(int32(2))%32)))) = v140
	v158 = v65 + int32(1)
	v159 = v148
	v161 = v149
	goto L20
L53:
	;
	v148 = v144
	v149 = v68 << (uint(int32(1)) % 32)
	goto L52
L54:
	;
	if v162 != 0 {
		v62 = v162
		v65 = v158
		v66 = v159
		v68 = v161
		goto L18
	} else {
		goto L55
	}
L55:
	;
	goto L19
L56:
	;
	F_sequence_close(m, v37, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if int32(2) <= v172 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	F_pfree(m, v173)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L76
	}
L59:
	;
	v193 = int32(0)
	v201 = v193
	v205 = v193
	goto L65
L60:
	;
	F_pg_qsort(m, v173, v172, int32(4), int32(471))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v172 != int32(1) {
		v241 = int32(0)
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
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v173+v201<<(uint(int32(2))%32))))
	if l2 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v241 = v227
	goto L58
L67:
	;
	v229 = v201 + int32(1)
	if v229 != v172 {
		v201 = v229
		v205 = v227
		goto L65
	} else {
		goto L75
	}
L68:
	;
	v225 = F_lappend_oid(m, v205, v212)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L74
	}
L69:
	;
	F_LockRelationOid(m, v212, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v218 = int32(0)
	v221 = F_SearchSysCacheExists(m, int32(57), v212, v218, v218, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v221 != 0 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	F_UnlockRelationOid(m, v212, l2)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v227 = v205
	goto L67
L74:
	;
	v227 = v225
	goto L67
L75:
	;
	goto L66
L76:
	;
	v257 = v241
	goto L9
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(46249), v17)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(493356), int32(362), int32(130450))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
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
	var v91 int32
	_ = v91
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v228
L2:
	;
	v228 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v14 = l0
	v15 = l1
	v17 = l1
	goto L5
L5:
	;
	v21 = int32(4)
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	switch v23 - int32(1) {
	case 0:
		goto L18
	case 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 15, 17, 18, 21, 23, 24, 25, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50:
		v228 = v22
		goto L1
	case 5:
		goto L17
	case 14:
		goto L16
	case 16:
		goto L15
	case 19:
		goto L14
	case 20:
		goto L13
	case 22:
		goto L10
	case 26, 27, 28, 29, 30:
		v221 = v17
		v222 = v21
		goto L7
	case 51:
		goto L12
	case 52:
		goto L11
	default:
		goto L9
	}
L6:
	;
	v228 = int32(0)
	goto L1
L7:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14+v222)))
	if v224 != 0 {
		v14 = v224
		v15 = v221
		v17 = v221
		goto L5
	} else {
		goto L98
	}
L8:
	;
	v221 = v217
	v222 = int32(28)
	goto L7
L9:
	;
	if v23 != int32(319) {
		v228 = v22
		goto L1
	} else {
		goto L77
	}
L10:
	;
	v148 = int32(8)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v149 == int32(2))&v15 != 0 {
		goto L73
	} else {
		goto L74
	}
L11:
	;
	if v15&int32(1) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L12:
	;
	if v15&int32(1) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v72 = int32(8)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	switch v73 {
	case 0:
		goto L37
	case 1:
		v78 = v15
		goto L36
	case 2:
		v221 = int32(0)
		v222 = v72
		goto L7
	default:
		goto L35
	}
L14:
	;
	v69 = F_is_strict_saop(m, v14)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L22
	} else {
		goto L33
	}
L15:
	;
	F_set_opfuncid(m, v14)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L22
	} else {
		goto L30
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v60 = F_func_strict(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L22
	} else {
		goto L28
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v54 != 0 {
		v228 = v22
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v26 <= int32(0) {
		v228 = v22
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v32 = v22
	v34 = int32(0)
	goto L20
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v34<<(uint(int32(2))%32))))
	v44 = F_find_nonnullable_rels_walker(m, v41, v15&int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v228 = v48
	goto L1
L22:
	;
	return int32(0)
L23:
	;
	v48 = F_bms_join(m, v32, v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v51 = v34 + int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v51 < v52 {
		v32 = v48
		v34 = v51
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v56 = F_bms_make_singleton(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v228 = v56
	goto L1
L28:
	;
	if v60 != 0 {
		v217 = int32(0)
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v228 = v22
	goto L1
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v66 = F_func_strict(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	if v66 != 0 {
		v217 = int32(0)
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v228 = v22
	goto L1
L33:
	;
	if v69 != 0 {
		v217 = int32(0)
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v228 = v22
	goto L1
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L22
	} else {
		goto L56
	}
L36:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v80 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v74 = int32(1)
	if v15&v74 != 0 {
		v221 = v74
		v222 = v72
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v78 = int32(0)
	goto L36
L39:
	;
	v228 = v22
	goto L1
L40:
	;
	goto L41
L41:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v83 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v228 = v22
	goto L1
L43:
	;
	goto L44
L44:
	;
	v89 = int32(0)
	v91 = v22
	goto L45
L45:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v89<<(uint(int32(2))%32))))
	v101 = F_find_nonnullable_rels_walker(m, v100, v78&int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L22
	} else {
		goto L47
	}
L46:
	;
	v228 = int32(0)
	goto L1
L47:
	;
	if v91 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v103 = F_bms_int_members(m, v91, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L22
	} else {
		goto L51
	}
L49:
	;
	v105 = v101
	goto L50
L50:
	;
	if v105 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v105 = v103
	goto L50
L52:
	;
	v107 = v89 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v108 <= v107 {
		v228 = v105
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L46
L55:
	;
	v89 = v107
	v91 = v105
	goto L45
L56:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v115
	F_errmsg_internal(m, int32(482460), v10)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(493833), int32(1572), int32(221383))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v228 = v22
	goto L1
L60:
	;
	goto L61
L61:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v129 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v228 = v22
	goto L1
L63:
	;
	goto L64
L64:
	;
	v132 = int32(0)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v133 == v132 {
		v221 = v132
		v222 = v21
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v228 = v22
	goto L1
L66:
	;
	v228 = v22
	goto L1
L67:
	;
	goto L68
L68:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(int32(5)) < base.Ui32(v141) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v228 = v22
	goto L1
L70:
	;
	goto L71
L71:
	;
	if int32(1)<<(uint(v141)%32)&int32(37) != 0 {
		v221 = int32(0)
		v222 = v21
		goto L7
	} else {
		goto L72
	}
L72:
	;
	v228 = v22
	goto L1
L73:
	;
	v221 = v15
	v222 = v148
	goto L7
L74:
	;
	goto L75
L75:
	;
	if v149 == int32(3) {
		v221 = v15
		v222 = v148
		goto L7
	} else {
		goto L76
	}
L76:
	;
	v228 = int32(0)
	goto L1
L77:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v161 = F_find_nonnullable_rels_walker(m, v158, v15&int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L22
	} else {
		goto L78
	}
L78:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if v163 != 0 {
		v228 = v161
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v164 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v211 != int32(1) {
		v228 = v161
		goto L1
	} else {
		goto L96
	}
L81:
	;
	v211 = int32(0)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v173 = int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v174 <= v173 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v177 = v173
	goto L86
L85:
	;
	v177 = v174
	goto L86
L86:
	;
	v180 = int32(0)
	v182 = v180
	v183 = v180
	goto L87
L87:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v164+int32(8)+v182<<(uint(int32(2))%32))))
	if v191 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v211 = v204
	goto L80
L89:
	;
	goto L88
L90:
	;
	v192 = int32(2)
	if v183 != 0 {
		v204 = v192
		goto L89
	} else {
		goto L93
	}
L91:
	;
	v197 = v183
	goto L92
L92:
	;
	v200 = v182 + int32(1)
	if v200 != v177 {
		v182 = v200
		v183 = v197
		goto L87
	} else {
		goto L95
	}
L93:
	;
	v193 = int32(1)
	if base.Ui32(v193) < base.Ui32(base.I32_popcnt(v191)) {
		v204 = v192
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v197 = v193
	goto L92
L95:
	;
	v204 = v197
	goto L89
L96:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v215 = F_bms_add_members(m, v161, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L22
	} else {
		goto L97
	}
L97:
	;
	v228 = v215
	goto L1
L98:
	;
	goto L6
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
	var v75 int32
	_ = v75
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
	var v99 int32
	_ = v99
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
	*(*int32)(unsafe.Add(mBase, uint32(v99+v101<<(uint(int32(2))%32)))) = v20
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
	v99 = v68
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
	v75 = v72
	goto L29
L29:
	;
	if base.Ui32(v75) <= base.Ui32(v65) {
		v75 = v75 << (uint(int32(1)) % 32)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v99 = v93
	v101 = v96
	goto L22
L33:
	;
	v83 = int32(2)
	v87 = F_repalloc0(m, v82, v66<<(uint(v83)%32), v75<<(uint(v83)%32))
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
	v91 = F_palloc0(m, v75<<(uint(int32(2))%32))
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
	F_errmsg_internal(m, int32(242176), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(495262), int32(104), int32(241552))
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
	F_errmsg_internal(m, int32(485725), v9)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(495262), int32(248), int32(360093))
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
					v57 = int32(0)
					return v57
				} else {
					v22 = v18
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v25 = F_palloc(m, int32(136))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v28 = F__emscripten_memcpy_bulkmem(m, v25, v23, int32(136))
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
						if v32 != 0 {
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
							v37 = v32
							v38 = v33
							v39 = v34
							v40 = v35
							for {
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
								v42 = v40 + v41
								*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v42
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
								v45 = v39 + v44
								*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v45
								v47 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
								v48 = v38 + v47
								*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v48
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
								if v50 != 0 {
									v37 = v50
									v38 = v48
									v39 = v45
									v40 = v42
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v57 = v28
						return v57
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
				v28 = F__emscripten_memcpy_bulkmem(m, v25, v23, int32(136))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				if v32 != 0 {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v28)+56))
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
					v37 = v32
					v38 = v33
					v39 = v34
					v40 = v35
					for {
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						v42 = v40 + v41
						*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v42
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
						v45 = v39 + v44
						*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v45
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
						v48 = v38 + v47
						*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
						if v50 != 0 {
							v37 = v50
							v38 = v48
							v39 = v45
							v40 = v42
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v57 = v28
				return v57
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
	F_errmsg_internal(m, int32(423624), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(492089), int32(732), int32(360326))
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
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
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v647 int32
	_ = v647
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
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
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
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
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
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
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
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
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1437 int32
	_ = v1437
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1472 int32
	_ = v1472
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1643 int32
	_ = v1643
	var v1647 int32
	_ = v1647
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1716 int32
	_ = v1716
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1735 int32
	_ = v1735
	var v1751 int32
	_ = v1751
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1847 int32
	_ = v1847
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1882 int32
	_ = v1882
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1908 int32
	_ = v1908
	var v1914 int32
	_ = v1914
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1947 int32
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1976 int32
	_ = v1976
	var v1989 int32
	_ = v1989
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2049 int32
	_ = v2049
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2076 int32
	_ = v2076
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2182 int32
	_ = v2182
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2217 int32
	_ = v2217
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v9
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = v23
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v501)+8)) = int32(0)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	if v506 < v508 {
		goto L105
	} else {
		goto L106
	}
L2:
	;
	if v128 < int32(0) {
		goto L1
	} else {
		goto L27
	}
L3:
	;
	v128 = v100
	goto L2
L4:
	;
	if v24 <= v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v128 = int32(-1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v40 = int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v25))))
	if base.Ui32(v42) < base.Ui32(int32(192)) {
		v99 = v42
		v100 = v40
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if int32(246) < v99 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	v46 = v33 + int32(1)
	if v46 == v24 {
		v99 = v42
		v100 = v40
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v25))))
	v51 = v49 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v42) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v25))))
	v67 = v65 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v42) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v55 = v33 + int32(2)
	if v55 != v24 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v99 = v42<<(uint(int32(6))%32)&int32(1984) | v51
	v100 = int32(2)
	goto L9
L16:
	;
	goto L15
L17:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v71))))
	v99 = v84&int32(63) | (v42<<(uint(int32(18))%32)&int32(1835008) | v51<<(uint(int32(12))%32) | v67<<(uint(int32(6))%32))
	v100 = int32(4)
	goto L9
L18:
	;
	v71 = v33 + int32(3)
	if v71 != v24 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v99 = v42<<(uint(int32(12))%32)&int32(61440) | v51<<(uint(int32(6))%32) | v67
	v100 = int32(3)
	goto L9
L21:
	;
	goto L20
L22:
	;
	v117 = v100 + v33
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v33 = v117
	goto L4
L23:
	;
	v104 = v99 - int32(97)
	if v104 < int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v104)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v110)>>(uint(v104&int32(7))%32))&int32(1) != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L22
L27:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = v142
	goto L30
L28:
	;
	if v248 < int32(0) {
		goto L1
	} else {
		goto L52
	}
L29:
	;
	v248 = v219
	goto L28
L30:
	;
	if v143 <= v152 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v248 = int32(-1)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v159 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v144))))
	if base.Ui32(v161) < base.Ui32(int32(192)) {
		v218 = v161
		v219 = v159
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(246) < v218 {
		goto L29
	} else {
		goto L48
	}
L36:
	;
	v165 = v152 + int32(1)
	if v165 == v143 {
		v218 = v161
		v219 = v159
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v144))))
	v170 = v168 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v161) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v144))))
	v186 = v184 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v161) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v174 = v152 + int32(2)
	if v174 != v143 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v218 = v161<<(uint(int32(6))%32)&int32(1984) | v170
	v219 = int32(2)
	goto L35
L42:
	;
	goto L41
L43:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v190))))
	v218 = v203&int32(63) | (v161<<(uint(int32(18))%32)&int32(1835008) | v170<<(uint(int32(12))%32) | v186<<(uint(int32(6))%32))
	v219 = int32(4)
	goto L35
L44:
	;
	v190 = v152 + int32(3)
	if v190 != v143 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v218 = v161<<(uint(int32(12))%32)&int32(61440) | v170<<(uint(int32(6))%32) | v186
	v219 = int32(3)
	goto L35
L47:
	;
	goto L46
L48:
	;
	v223 = v218 - int32(97)
	if v223 < int32(0) {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v223)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v229)>>(uint(v223&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v237 = v219 + v152
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v237
	v152 = v237
	goto L30
L52:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v252 = v251 + v248
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v252
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v277 = v267
	goto L55
L53:
	;
	if v372 < int32(0) {
		goto L1
	} else {
		goto L78
	}
L54:
	;
	v372 = v344
	goto L53
L55:
	;
	if v268 <= v277 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v372 = int32(-1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v284 = int32(1)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v269))))
	if base.Ui32(v286) < base.Ui32(int32(192)) {
		v343 = v286
		v344 = v284
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if int32(246) < v343 {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v290 = v277 + int32(1)
	if v290 == v268 {
		v343 = v286
		v344 = v284
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v269))))
	v295 = v293 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v286) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v269))))
	v311 = v309 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v286) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v299 = v277 + int32(2)
	if v299 != v268 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v343 = v286<<(uint(int32(6))%32)&int32(1984) | v295
	v344 = int32(2)
	goto L60
L67:
	;
	goto L66
L68:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v315))))
	v343 = v328&int32(63) | (v286<<(uint(int32(18))%32)&int32(1835008) | v295<<(uint(int32(12))%32) | v311<<(uint(int32(6))%32))
	v344 = int32(4)
	goto L60
L69:
	;
	v315 = v277 + int32(3)
	if v315 != v268 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v343 = v286<<(uint(int32(12))%32)&int32(61440) | v295<<(uint(int32(6))%32) | v311
	v344 = int32(3)
	goto L60
L72:
	;
	goto L71
L73:
	;
	v361 = v344 + v277
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v277 = v361
	goto L55
L74:
	;
	v348 = v343 - int32(97)
	if v348 < int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v348)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v354)>>(uint(v348&int32(7))%32))&int32(1) != 0 {
		goto L54
	} else {
		goto L76
	}
L76:
	;
	goto L73
L78:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v396 = v386
	goto L81
L79:
	;
	if v492 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L80:
	;
	v492 = v463
	goto L79
L81:
	;
	if v387 <= v396 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v492 = int32(-1)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v403 = int32(1)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v388))))
	if base.Ui32(v405) < base.Ui32(int32(192)) {
		v462 = v405
		v463 = v403
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if int32(246) < v462 {
		goto L80
	} else {
		goto L99
	}
L87:
	;
	v409 = v396 + int32(1)
	if v409 == v387 {
		v462 = v405
		v463 = v403
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409+v388))))
	v414 = v412 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v405) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418+v388))))
	v430 = v428 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v405) {
		goto L95
	} else {
		goto L96
	}
L90:
	;
	v418 = v396 + int32(2)
	if v418 != v387 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v462 = v405<<(uint(int32(6))%32)&int32(1984) | v414
	v463 = int32(2)
	goto L86
L93:
	;
	goto L92
L94:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388+v434))))
	v462 = v447&int32(63) | (v405<<(uint(int32(18))%32)&int32(1835008) | v414<<(uint(int32(12))%32) | v430<<(uint(int32(6))%32))
	v463 = int32(4)
	goto L86
L95:
	;
	v434 = v396 + int32(3)
	if v434 != v387 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v462 = v405<<(uint(int32(12))%32)&int32(61440) | v414<<(uint(int32(6))%32) | v430
	v463 = int32(3)
	goto L86
L98:
	;
	goto L97
L99:
	;
	v467 = v462 - int32(97)
	if v467 < int32(0) {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v467)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v473)>>(uint(v467&int32(7))%32))&int32(1) == int32(0) {
		goto L80
	} else {
		goto L101
	}
L101:
	;
	v481 = v463 + v396
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v481
	v396 = v481
	goto L81
L103:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v496 + v492
	goto L1
L104:
	;
	return v2269
L105:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v666
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	if v666 < v669 {
		goto L143
	} else {
		goto L144
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v508
	v514 = F_find_among_b(m, l0, int32(4269408), int32(10))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	return int32(0)
L108:
	;
	if v514 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v504
	goto L105
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v504
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v522
	switch v514 - int32(1) {
	case 0:
		goto L114
	case 1:
		goto L113
	default:
		goto L112
	}
L112:
	;
	v660 = F_slice_del(m, l0)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L107
	} else {
		goto L141
	}
L113:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	if v522 < v658 {
		goto L105
	} else {
		goto L140
	}
L114:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L117
L115:
	;
	if v654 == int32(0) {
		goto L112
	} else {
		goto L139
	}
L116:
	;
	v654 = v647
	goto L115
L117:
	;
	if v542 <= v543 {
		v647 = int32(-1)
		goto L116
	} else {
		goto L119
	}
L118:
	;
	v647 = int32(0)
	goto L116
L119:
	;
	v560 = int32(1)
	v561 = v542 - v560
	v563 = int32(*(*int8)(unsafe.Add(mBase, uint32(v539+v561))))
	v565 = v563 & int32(255)
	if v561 == v543 {
		v620 = v565
		v621 = v560
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if int32(246) < v620 {
		goto L129
	} else {
		goto L130
	}
L121:
	;
	if int32(0) <= v563 {
		v620 = v565
		v621 = v560
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v571 = v565 & int32(63)
	v573 = v542 - int32(2)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539+v573))))
	v577 = v575 << (uint(int32(6)) % 32)
	if base.B2i32(v573 != v543)&base.B2i32(base.Ui32(v575) < base.Ui32(int32(192))) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v620 = v577&int32(1984) | v571
	v621 = int32(2)
	goto L120
L124:
	;
	goto L125
L125:
	;
	v590 = v577&int32(4032) | v571
	v592 = v542 - int32(3)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539+v592))))
	if base.B2i32(v592 != v543)&base.B2i32(base.Ui32(v594) < base.Ui32(int32(224))) == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v620 = v594<<(uint(int32(12))%32)&int32(61440) | v590
	v621 = int32(3)
	goto L120
L127:
	;
	goto L128
L128:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542+(v539-int32(4))))))
	v620 = v594<<(uint(int32(12))%32)&int32(258048) | v612&int32(7)<<(uint(int32(18))%32) | v590
	v621 = int32(4)
	goto L120
L129:
	;
	v654 = v621
	goto L115
L130:
	;
	goto L131
L131:
	;
	v625 = v620 - int32(97)
	if v625 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v654 = v621
	goto L115
L133:
	;
	goto L134
L134:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v625)>>(uint(int32(3))%32)))+uint32(_consts[1484]))))
	if int32(base.Ui32(v631)>>(uint(v625&int32(7))%32))&int32(1) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v654 = v621
	goto L115
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v542 - v621
	goto L138
L138:
	;
	goto L118
L139:
	;
	goto L105
L140:
	;
	goto L112
L141:
	;
	if v660 < int32(0) {
		v2269 = v660
		goto L104
	} else {
		goto L142
	}
L142:
	;
	goto L105
L143:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v795
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+4))
	if v795 < v798 {
		goto L190
	} else {
		goto L191
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v666
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v669
	v676 = F_find_among_b(m, l0, int32(4269616), int32(9))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L107
	} else {
		goto L145
	}
L145:
	;
	if v676 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v672
	goto L143
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v672
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v682
	switch v676 - int32(1) {
	case 0:
		goto L154
	case 1:
		goto L153
	case 2:
		goto L152
	case 3:
		goto L151
	case 4:
		goto L150
	case 5:
		goto L149
	default:
		goto L143
	}
L149:
	;
	if v682-int32(2) <= v672 {
		goto L143
	} else {
		goto L184
	}
L150:
	;
	if v682-int32(2) <= v672 {
		goto L143
	} else {
		goto L178
	}
L151:
	;
	v735 = v682 - int32(1)
	if v735 <= v672 {
		goto L143
	} else {
		goto L172
	}
L152:
	;
	v730 = F_slice_del(m, l0)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L107
	} else {
		goto L170
	}
L153:
	;
	v698 = F_slice_del(m, l0)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L107
	} else {
		goto L161
	}
L154:
	;
	if v672 < v682 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687+v682-int32(1)))))
	if v691 == int32(107) {
		goto L143
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v694 = F_slice_del(m, l0)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L107
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	if int32(0) <= v694 {
		goto L143
	} else {
		goto L160
	}
L160:
	;
	v2269 = v694
	goto L104
L161:
	;
	if v698 < int32(0) {
		v2269 = v698
		goto L104
	} else {
		goto L162
	}
L162:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v702
	v704 = int32(3)
	v706 = int32(0)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v702-v709 < v704 {
		v719 = v706
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v719 == int32(0) {
		goto L143
	} else {
		goto L167
	}
L164:
	;
	goto L163
L165:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v715 = F_memcmp(m, v712+v702-v704, int32(2193124), v704)
	mBase = m.M
	if v715 != 0 {
		v719 = v706
		goto L164
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v702 - v704
	v719 = int32(1)
	goto L164
L167:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v722
	v726 = F_slice_from_s(m, l0, int32(3), int32(2193127))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L107
	} else {
		goto L168
	}
L168:
	;
	if int32(0) <= v726 {
		goto L143
	} else {
		goto L169
	}
L169:
	;
	v2269 = v726
	goto L104
L170:
	;
	if int32(0) <= v730 {
		goto L143
	} else {
		goto L171
	}
L171:
	;
	v2269 = v730
	goto L104
L172:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737+v735))))
	if v739 != int32(97) {
		goto L143
	} else {
		goto L173
	}
L173:
	;
	v744 = F_find_among_b(m, l0, int32(4269808), int32(6))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L107
	} else {
		goto L174
	}
L174:
	;
	if v744 == int32(0) {
		goto L143
	} else {
		goto L175
	}
L175:
	;
	v748 = F_slice_del(m, l0)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L107
	} else {
		goto L176
	}
L176:
	;
	if int32(0) <= v748 {
		goto L143
	} else {
		goto L177
	}
L177:
	;
	v2269 = v748
	goto L104
L178:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755+v682-int32(1)))))
	if v759 != int32(164) {
		goto L143
	} else {
		goto L179
	}
L179:
	;
	v764 = F_find_among_b(m, l0, int32(4269936), int32(6))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L107
	} else {
		goto L180
	}
L180:
	;
	if v764 == int32(0) {
		goto L143
	} else {
		goto L181
	}
L181:
	;
	v768 = F_slice_del(m, l0)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L107
	} else {
		goto L182
	}
L182:
	;
	if int32(0) <= v768 {
		goto L143
	} else {
		goto L183
	}
L183:
	;
	v2269 = v768
	goto L104
L184:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775+v682-int32(1)))))
	if v779 != int32(101) {
		goto L143
	} else {
		goto L185
	}
L185:
	;
	v784 = F_find_among_b(m, l0, int32(4270064), int32(2))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L107
	} else {
		goto L186
	}
L186:
	;
	if v784 == int32(0) {
		goto L143
	} else {
		goto L187
	}
L187:
	;
	v788 = F_slice_del(m, l0)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L107
	} else {
		goto L188
	}
L188:
	;
	if v788 < int32(0) {
		v2269 = v788
		goto L104
	} else {
		goto L189
	}
L189:
	;
	goto L143
L190:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1252
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	if v1252 < v1255 {
		goto L306
	} else {
		goto L307
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v795
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v798
	v805 = F_find_among_b(m, l0, int32(4270112), int32(30))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L107
	} else {
		goto L192
	}
L192:
	;
	if v805 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v801
	goto L190
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v801
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v811
	switch v805 - int32(1) {
	case 0:
		goto L204
	case 1:
		goto L203
	case 2:
		goto L202
	case 3:
		goto L201
	case 4:
		goto L200
	case 5:
		goto L199
	case 6:
		goto L198
	case 7:
		goto L197
	default:
		goto L196
	}
L196:
	;
	v1242 = F_slice_del(m, l0)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L107
	} else {
		goto L304
	}
L197:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L256
L198:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v892 = v891 - v811
	v895 = F_find_among_b(m, l0, int32(4270720), int32(7))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L107
	} else {
		goto L224
	}
L199:
	;
	v875 = int32(2)
	v877 = int32(0)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v879-v880 < v875 {
		v890 = v877
		goto L219
	} else {
		goto L220
	}
L200:
	;
	v859 = int32(2)
	v861 = int32(0)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v863-v864 < v859 {
		v874 = v861
		goto L214
	} else {
		goto L215
	}
L201:
	;
	if v811 <= v801 {
		goto L190
	} else {
		goto L211
	}
L202:
	;
	if v811 <= v801 {
		goto L190
	} else {
		goto L209
	}
L203:
	;
	if v811 <= v801 {
		goto L190
	} else {
		goto L207
	}
L204:
	;
	if v811 <= v801 {
		goto L190
	} else {
		goto L205
	}
L205:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816+v811-int32(1)))))
	if v820 != int32(97) {
		goto L190
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811 - int32(1)
	goto L196
L207:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827+v811-int32(1)))))
	if v831 != int32(101) {
		goto L190
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811 - int32(1)
	goto L196
L209:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v811-int32(1)))))
	if v842 != int32(105) {
		goto L190
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811 - int32(1)
	goto L196
L211:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849+v811-int32(1)))))
	if v853 != int32(111) {
		goto L190
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811 - int32(1)
	goto L196
L213:
	;
	if v874 != 0 {
		goto L196
	} else {
		goto L217
	}
L214:
	;
	goto L213
L215:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v870 = F_memcmp(m, v867+v863-v859, int32(2193198), v859)
	mBase = m.M
	if v870 != 0 {
		v874 = v861
		goto L214
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v863 - v859
	v874 = int32(1)
	goto L214
L217:
	;
	goto L190
L218:
	;
	if v890 != 0 {
		goto L196
	} else {
		goto L222
	}
L219:
	;
	goto L218
L220:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v886 = F_memcmp(m, v883+v879-v875, int32(2193200), v875)
	mBase = m.M
	if v886 != 0 {
		v890 = v877
		goto L219
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v879 - v875
	v890 = int32(1)
	goto L219
L222:
	;
	goto L190
L223:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v920 = v919 - v892
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L233
L224:
	;
	if v895 != 0 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v898 = v897 - v892
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v898
	v900 = int32(2)
	v902 = int32(0)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v898-v905 < v900 {
		v915 = v902
		goto L227
	} else {
		goto L228
	}
L226:
	;
	if v915 != 0 {
		goto L223
	} else {
		goto L230
	}
L227:
	;
	goto L226
L228:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v911 = F_memcmp(m, v908+v898-v900, int32(2193202), v900)
	mBase = m.M
	if v911 != 0 {
		v915 = v902
		goto L227
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v898 - v900
	v915 = int32(1)
	goto L227
L230:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v916 - v892
	goto L196
L231:
	;
	if v974 < int32(0) {
		goto L251
	} else {
		goto L252
	}
L233:
	;
	goto L234
L234:
	;
	goto L235
L235:
	;
	v930 = v920
	v932 = int32(1)
	goto L238
L237:
	;
	v974 = v956
	goto L231
L238:
	;
	if v930 <= v923 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	goto L237
L240:
	;
	v974 = int32(-1)
	goto L231
L241:
	;
	goto L242
L242:
	;
	v937 = v930 - int32(1)
	v939 = int32(*(*int8)(unsafe.Add(mBase, uint32(v922+v937))))
	if int32(0) <= v939 {
		v956 = v937
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v960 = int32(1)
	if v960 < v932 {
		v930 = v956
		v932 = v932 - v960
		goto L238
	} else {
		goto L250
	}
L244:
	;
	if v937 <= v923 {
		v956 = v937
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v944 = v937
	goto L246
L246:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922+v944))))
	if base.Ui32(int32(191)) < base.Ui32(v949) {
		v956 = v944
		goto L243
	} else {
		goto L248
	}
L247:
	;
	v956 = v923
	goto L243
L248:
	;
	v953 = v944 - int32(1)
	if v923 < v953 {
		v944 = v953
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	goto L239
L251:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977 - v892
	goto L196
L252:
	;
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v974
	goto L196
L254:
	;
	if v1110 != 0 {
		goto L190
	} else {
		goto L278
	}
L255:
	;
	v1110 = v1103
	goto L254
L256:
	;
	if v998 <= v999 {
		v1103 = int32(-1)
		goto L255
	} else {
		goto L258
	}
L257:
	;
	v1103 = int32(0)
	goto L255
L258:
	;
	v1016 = int32(1)
	v1017 = v998 - v1016
	v1019 = int32(*(*int8)(unsafe.Add(mBase, uint32(v995+v1017))))
	v1021 = v1019 & int32(255)
	if v1017 == v999 {
		v1076 = v1021
		v1077 = v1016
		goto L259
	} else {
		goto L260
	}
L259:
	;
	if int32(246) < v1076 {
		goto L268
	} else {
		goto L269
	}
L260:
	;
	if int32(0) <= v1019 {
		v1076 = v1021
		v1077 = v1016
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1027 = v1021 & int32(63)
	v1029 = v998 - int32(2)
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995+v1029))))
	v1033 = v1031 << (uint(int32(6)) % 32)
	if base.B2i32(v1029 != v999)&base.B2i32(base.Ui32(v1031) < base.Ui32(int32(192))) == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1076 = v1033&int32(1984) | v1027
	v1077 = int32(2)
	goto L259
L263:
	;
	goto L264
L264:
	;
	v1046 = v1033&int32(4032) | v1027
	v1048 = v998 - int32(3)
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995+v1048))))
	if base.B2i32(v1048 != v999)&base.B2i32(base.Ui32(v1050) < base.Ui32(int32(224))) == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1076 = v1050<<(uint(int32(12))%32)&int32(61440) | v1046
	v1077 = int32(3)
	goto L259
L266:
	;
	goto L267
L267:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998+(v995-int32(4))))))
	v1076 = v1050<<(uint(int32(12))%32)&int32(258048) | v1068&int32(7)<<(uint(int32(18))%32) | v1046
	v1077 = int32(4)
	goto L259
L268:
	;
	v1110 = v1077
	goto L254
L269:
	;
	goto L270
L270:
	;
	v1081 = v1076 - int32(97)
	if v1081 < int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1110 = v1077
	goto L254
L272:
	;
	goto L273
L273:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1081)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v1087)>>(uint(v1081&int32(7))%32))&int32(1) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1110 = v1077
	goto L254
L275:
	;
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v998 - v1077
	goto L277
L277:
	;
	goto L257
L278:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L281
L279:
	;
	if v1239 != 0 {
		goto L190
	} else {
		goto L303
	}
L280:
	;
	v1239 = v1232
	goto L279
L281:
	;
	if v1127 <= v1128 {
		v1232 = int32(-1)
		goto L280
	} else {
		goto L283
	}
L282:
	;
	v1232 = int32(0)
	goto L280
L283:
	;
	v1145 = int32(1)
	v1146 = v1127 - v1145
	v1148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1124+v1146))))
	v1150 = v1148 & int32(255)
	if v1146 == v1128 {
		v1205 = v1150
		v1206 = v1145
		goto L284
	} else {
		goto L285
	}
L284:
	;
	if int32(122) < v1205 {
		goto L293
	} else {
		goto L294
	}
L285:
	;
	if int32(0) <= v1148 {
		v1205 = v1150
		v1206 = v1145
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1156 = v1150 & int32(63)
	v1158 = v1127 - int32(2)
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1158))))
	v1162 = v1160 << (uint(int32(6)) % 32)
	if base.B2i32(v1158 != v1128)&base.B2i32(base.Ui32(v1160) < base.Ui32(int32(192))) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1205 = v1162&int32(1984) | v1156
	v1206 = int32(2)
	goto L284
L288:
	;
	goto L289
L289:
	;
	v1175 = v1162&int32(4032) | v1156
	v1177 = v1127 - int32(3)
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1177))))
	if base.B2i32(v1177 != v1128)&base.B2i32(base.Ui32(v1179) < base.Ui32(int32(224))) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1205 = v1179<<(uint(int32(12))%32)&int32(61440) | v1175
	v1206 = int32(3)
	goto L284
L291:
	;
	goto L292
L292:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127+(v1124-int32(4))))))
	v1205 = v1179<<(uint(int32(12))%32)&int32(258048) | v1197&int32(7)<<(uint(int32(18))%32) | v1175
	v1206 = int32(4)
	goto L284
L293:
	;
	v1239 = v1206
	goto L279
L294:
	;
	goto L295
L295:
	;
	v1210 = v1205 - int32(98)
	if v1210 < int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1239 = v1206
	goto L279
L297:
	;
	goto L298
L298:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1210)>>(uint(int32(3))%32)))+uint32(_consts[1485]))))
	if int32(base.Ui32(v1216)>>(uint(v1210&int32(7))%32))&int32(1) == int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1239 = v1206
	goto L279
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1127 - v1206
	goto L302
L302:
	;
	goto L282
L303:
	;
	goto L196
L304:
	;
	if v1242 < int32(0) {
		v2269 = v1242
		goto L104
	} else {
		goto L305
	}
L305:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1246)+8)) = int32(1)
	goto L190
L306:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1301
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+4))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+8))
	if v1305 != 0 {
		goto L324
	} else {
		goto L325
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1252
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1255
	v1262 = F_find_among_b(m, l0, int32(4270864), int32(14))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L107
	} else {
		goto L308
	}
L308:
	;
	if v1262 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1258
	goto L306
L310:
	;
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1258
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1268
	if v1262 == int32(1) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1273 = int32(2)
	v1275 = int32(0)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1277-v1278 < v1273 {
		v1288 = v1275
		goto L316
	} else {
		goto L317
	}
L313:
	;
	goto L314
L314:
	;
	v1294 = F_slice_del(m, l0)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L107
	} else {
		goto L320
	}
L315:
	;
	if v1288 != 0 {
		goto L306
	} else {
		goto L319
	}
L316:
	;
	goto L315
L317:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1284 = F_memcmp(m, v1281+v1277-v1273, int32(2193349), v1273)
	mBase = m.M
	if v1284 != 0 {
		v1288 = v1275
		goto L316
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1277 - v1273
	v1288 = int32(1)
	goto L316
L319:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1289 + (v1268 - v1272)
	goto L314
L320:
	;
	if v1294 < int32(0) {
		v2269 = v1294
		goto L104
	} else {
		goto L321
	}
L321:
	;
	goto L306
L322:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1549
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+4))
	if v1549 < v1552 {
		goto L379
	} else {
		goto L380
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1544
	goto L322
L324:
	;
	if v1301 < v1304 {
		goto L322
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if v1301 < v1304 {
		goto L322
	} else {
		goto L334
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1301
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1304
	if v1301 <= v1304 {
		v1544 = v1308
		goto L323
	} else {
		goto L328
	}
L328:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1313 = int32(1)
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311+v1301-v1313))))
	if base.Ui32(v1313) < base.Ui32((v1315-int32(105))&int32(255)) {
		v1544 = v1308
		goto L323
	} else {
		goto L329
	}
L329:
	;
	v1324 = F_find_among_b(m, l0, int32(4271152), int32(2))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L107
	} else {
		goto L330
	}
L330:
	;
	if v1324 == int32(0) {
		v1544 = v1308
		goto L323
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1308
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1329
	v1331 = F_slice_del(m, l0)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L107
	} else {
		goto L332
	}
L332:
	;
	if int32(0) <= v1331 {
		goto L322
	} else {
		goto L333
	}
L333:
	;
	v2269 = v1331
	goto L104
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1301
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1304
	if v1301 <= v1304 {
		v1544 = v1337
		goto L323
	} else {
		goto L335
	}
L335:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1340+v1301-int32(1)))))
	if v1344 != int32(116) {
		v1544 = v1337
		goto L323
	} else {
		goto L336
	}
L336:
	;
	v1348 = v1301 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1348
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1348
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L339
L337:
	;
	if v1479 != 0 {
		v1544 = v1337
		goto L323
	} else {
		goto L361
	}
L338:
	;
	v1479 = v1472
	goto L337
L339:
	;
	if v1348 <= v1368 {
		v1472 = int32(-1)
		goto L338
	} else {
		goto L341
	}
L340:
	;
	v1472 = int32(0)
	goto L338
L341:
	;
	v1385 = int32(1)
	v1386 = v1348 - v1385
	v1388 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1364+v1386))))
	v1390 = v1388 & int32(255)
	if v1386 == v1368 {
		v1445 = v1390
		v1446 = v1385
		goto L342
	} else {
		goto L343
	}
L342:
	;
	if int32(246) < v1445 {
		goto L351
	} else {
		goto L352
	}
L343:
	;
	if int32(0) <= v1388 {
		v1445 = v1390
		v1446 = v1385
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1396 = v1390 & int32(63)
	v1398 = v1348 - int32(2)
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364+v1398))))
	v1402 = v1400 << (uint(int32(6)) % 32)
	if base.B2i32(v1398 != v1368)&base.B2i32(base.Ui32(v1400) < base.Ui32(int32(192))) == int32(0) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1445 = v1402&int32(1984) | v1396
	v1446 = int32(2)
	goto L342
L346:
	;
	goto L347
L347:
	;
	v1415 = v1402&int32(4032) | v1396
	v1417 = v1348 - int32(3)
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364+v1417))))
	if base.B2i32(v1417 != v1368)&base.B2i32(base.Ui32(v1419) < base.Ui32(int32(224))) == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1445 = v1419<<(uint(int32(12))%32)&int32(61440) | v1415
	v1446 = int32(3)
	goto L342
L349:
	;
	goto L350
L350:
	;
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1348+(v1364-int32(4))))))
	v1445 = v1419<<(uint(int32(12))%32)&int32(258048) | v1437&int32(7)<<(uint(int32(18))%32) | v1415
	v1446 = int32(4)
	goto L342
L351:
	;
	v1479 = v1446
	goto L337
L352:
	;
	goto L353
L353:
	;
	v1450 = v1445 - int32(97)
	if v1450 < int32(0) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1479 = v1446
	goto L337
L355:
	;
	goto L356
L356:
	;
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1450)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v1456)>>(uint(v1450&int32(7))%32))&int32(1) == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1479 = v1446
	goto L337
L358:
	;
	goto L359
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1348 - v1446
	goto L360
L360:
	;
	goto L340
L361:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1480 - int32(1)
	v1484 = F_slice_del(m, l0)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L107
	} else {
		goto L362
	}
L362:
	;
	if v1484 < int32(0) {
		v2269 = v1484
		goto L104
	} else {
		goto L363
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1337
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)))
	if v1489 < v1491 {
		goto L322
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1489
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1491
	if v1489-int32(2) <= v1491 {
		v1544 = v1337
		goto L323
	} else {
		goto L365
	}
L365:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498+v1489-int32(1)))))
	if v1502 != int32(97) {
		v1544 = v1337
		goto L323
	} else {
		goto L366
	}
L366:
	;
	v1507 = F_find_among_b(m, l0, int32(4271200), int32(2))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L107
	} else {
		goto L367
	}
L367:
	;
	if v1507 == int32(0) {
		v1544 = v1337
		goto L323
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1337
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1512
	if v1507 == int32(1) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1517 = int32(2)
	v1519 = int32(0)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1521-v1522 < v1517 {
		v1532 = v1519
		goto L373
	} else {
		goto L374
	}
L370:
	;
	goto L371
L371:
	;
	v1538 = F_slice_del(m, l0)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L107
	} else {
		goto L377
	}
L372:
	;
	if v1532 != 0 {
		goto L322
	} else {
		goto L376
	}
L373:
	;
	goto L372
L374:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1528 = F_memcmp(m, v1525+v1521-v1517, int32(2193406), v1517)
	mBase = m.M
	if v1528 != 0 {
		v1532 = v1519
		goto L373
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1521 - v1517
	v1532 = int32(1)
	goto L373
L376:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1533 + (v1512 - v1516)
	goto L371
L377:
	;
	if int32(0) <= v1538 {
		goto L322
	} else {
		goto L378
	}
L378:
	;
	v2269 = v1538
	goto L104
L379:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2266
	v2269 = int32(1)
	goto L104
L380:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1552
	v1558 = F_find_among_b(m, l0, int32(4270720), int32(7))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L107
	} else {
		goto L382
	}
L381:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1627
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1627
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L410
L382:
	;
	if v1558 == int32(0) {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1562
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L386
L384:
	;
	if v1617 < int32(0) {
		goto L381
	} else {
		goto L404
	}
L386:
	;
	goto L387
L387:
	;
	goto L388
L388:
	;
	v1573 = v1562
	v1575 = int32(1)
	goto L391
L390:
	;
	v1617 = v1599
	goto L384
L391:
	;
	if v1573 <= v1566 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	goto L390
L393:
	;
	v1617 = int32(-1)
	goto L384
L394:
	;
	goto L395
L395:
	;
	v1580 = v1573 - int32(1)
	v1582 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1565+v1580))))
	if int32(0) <= v1582 {
		v1599 = v1580
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1603 = int32(1)
	if v1603 < v1575 {
		v1573 = v1599
		v1575 = v1575 - v1603
		goto L391
	} else {
		goto L403
	}
L397:
	;
	if v1580 <= v1566 {
		v1599 = v1580
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1587 = v1580
	goto L399
L399:
	;
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565+v1587))))
	if base.Ui32(int32(191)) < base.Ui32(v1592) {
		v1599 = v1587
		goto L396
	} else {
		goto L401
	}
L400:
	;
	v1599 = v1566
	goto L396
L401:
	;
	v1596 = v1587 - int32(1)
	if v1566 < v1596 {
		v1587 = v1596
		goto L399
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	goto L392
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1617
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1617
	v1622 = F_slice_del(m, l0)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L107
	} else {
		goto L405
	}
L405:
	;
	if v1622 < int32(0) {
		v2269 = v1622
		goto L104
	} else {
		goto L406
	}
L406:
	;
	goto L381
L407:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1895
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1895
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1895 <= v1898 {
		v1926 = v1895
		v1927 = v1898
		goto L460
	} else {
		goto L461
	}
L408:
	;
	if v1758 != 0 {
		goto L407
	} else {
		goto L432
	}
L409:
	;
	v1758 = v1751
	goto L408
L410:
	;
	if v1627 <= v1647 {
		v1751 = int32(-1)
		goto L409
	} else {
		goto L412
	}
L411:
	;
	v1751 = int32(0)
	goto L409
L412:
	;
	v1664 = int32(1)
	v1665 = v1627 - v1664
	v1667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1643+v1665))))
	v1669 = v1667 & int32(255)
	if v1665 == v1647 {
		v1724 = v1669
		v1725 = v1664
		goto L413
	} else {
		goto L414
	}
L413:
	;
	if int32(228) < v1724 {
		goto L422
	} else {
		goto L423
	}
L414:
	;
	if int32(0) <= v1667 {
		v1724 = v1669
		v1725 = v1664
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1675 = v1669 & int32(63)
	v1677 = v1627 - int32(2)
	v1679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643+v1677))))
	v1681 = v1679 << (uint(int32(6)) % 32)
	if base.B2i32(v1677 != v1647)&base.B2i32(base.Ui32(v1679) < base.Ui32(int32(192))) == int32(0) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1724 = v1681&int32(1984) | v1675
	v1725 = int32(2)
	goto L413
L417:
	;
	goto L418
L418:
	;
	v1694 = v1681&int32(4032) | v1675
	v1696 = v1627 - int32(3)
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1643+v1696))))
	if base.B2i32(v1696 != v1647)&base.B2i32(base.Ui32(v1698) < base.Ui32(int32(224))) == int32(0) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1724 = v1698<<(uint(int32(12))%32)&int32(61440) | v1694
	v1725 = int32(3)
	goto L413
L420:
	;
	goto L421
L421:
	;
	v1716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1627+(v1643-int32(4))))))
	v1724 = v1698<<(uint(int32(12))%32)&int32(258048) | v1716&int32(7)<<(uint(int32(18))%32) | v1694
	v1725 = int32(4)
	goto L413
L422:
	;
	v1758 = v1725
	goto L408
L423:
	;
	goto L424
L424:
	;
	v1729 = v1724 - int32(97)
	if v1729 < int32(0) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1758 = v1725
	goto L408
L426:
	;
	goto L427
L427:
	;
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1729)>>(uint(int32(3))%32)))+uint32(_consts[1486]))))
	if int32(base.Ui32(v1735)>>(uint(v1729&int32(7))%32))&int32(1) == int32(0) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1758 = v1725
	goto L408
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1627 - v1725
	goto L431
L431:
	;
	goto L411
L432:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1759
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L435
L433:
	;
	if v1889 != 0 {
		goto L407
	} else {
		goto L457
	}
L434:
	;
	v1889 = v1882
	goto L433
L435:
	;
	if v1759 <= v1778 {
		v1882 = int32(-1)
		goto L434
	} else {
		goto L437
	}
L436:
	;
	v1882 = int32(0)
	goto L434
L437:
	;
	v1795 = int32(1)
	v1796 = v1759 - v1795
	v1798 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1774+v1796))))
	v1800 = v1798 & int32(255)
	if v1796 == v1778 {
		v1855 = v1800
		v1856 = v1795
		goto L438
	} else {
		goto L439
	}
L438:
	;
	if int32(122) < v1855 {
		goto L447
	} else {
		goto L448
	}
L439:
	;
	if int32(0) <= v1798 {
		v1855 = v1800
		v1856 = v1795
		goto L438
	} else {
		goto L440
	}
L440:
	;
	v1806 = v1800 & int32(63)
	v1808 = v1759 - int32(2)
	v1810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1808))))
	v1812 = v1810 << (uint(int32(6)) % 32)
	if base.B2i32(v1808 != v1778)&base.B2i32(base.Ui32(v1810) < base.Ui32(int32(192))) == int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v1855 = v1812&int32(1984) | v1806
	v1856 = int32(2)
	goto L438
L442:
	;
	goto L443
L443:
	;
	v1825 = v1812&int32(4032) | v1806
	v1827 = v1759 - int32(3)
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1827))))
	if base.B2i32(v1827 != v1778)&base.B2i32(base.Ui32(v1829) < base.Ui32(int32(224))) == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1855 = v1829<<(uint(int32(12))%32)&int32(61440) | v1825
	v1856 = int32(3)
	goto L438
L445:
	;
	goto L446
L446:
	;
	v1847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759+(v1774-int32(4))))))
	v1855 = v1829<<(uint(int32(12))%32)&int32(258048) | v1847&int32(7)<<(uint(int32(18))%32) | v1825
	v1856 = int32(4)
	goto L438
L447:
	;
	v1889 = v1856
	goto L433
L448:
	;
	goto L449
L449:
	;
	v1860 = v1855 - int32(98)
	if v1860 < int32(0) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1889 = v1856
	goto L433
L451:
	;
	goto L452
L452:
	;
	v1866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1860)>>(uint(int32(3))%32)))+uint32(_consts[1485]))))
	if int32(base.Ui32(v1866)>>(uint(v1860&int32(7))%32))&int32(1) == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1889 = v1856
	goto L433
L454:
	;
	goto L455
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1759 - v1856
	goto L456
L456:
	;
	goto L436
L457:
	;
	v1890 = F_slice_del(m, l0)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L107
	} else {
		goto L458
	}
L458:
	;
	if v1890 < int32(0) {
		v2269 = v1890
		goto L104
	} else {
		goto L459
	}
L459:
	;
	goto L407
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1926
	if v1926 <= v1927 {
		v1958 = v1926
		goto L467
	} else {
		goto L468
	}
L461:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1901 = v1900 + v1895
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1901-int32(1)))))
	if v1904 != int32(106) {
		v1926 = v1895
		v1927 = v1898
		goto L460
	} else {
		goto L462
	}
L462:
	;
	v1908 = v1895 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1908
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1908
	if v1908 <= v1898 {
		v1926 = v1895
		v1927 = v1898
		goto L460
	} else {
		goto L463
	}
L463:
	;
	v1914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1901-int32(2)))))
	switch v1914 - int32(111) {
	case 0, 6:
		goto L464
	default:
		v1926 = v1895
		v1927 = v1898
		goto L460
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1895 - int32(2)
	v1920 = F_slice_del(m, l0)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L107
	} else {
		goto L465
	}
L465:
	;
	if v1920 < int32(0) {
		v2269 = v1920
		goto L104
	} else {
		goto L466
	}
L466:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1926 = v1925
	v1927 = v1924
	goto L460
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1554
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1958
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1989 = v1958
	goto L476
L468:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1934 = v1933 + v1926
	v1937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1934-int32(1)))))
	if v1937 != int32(111) {
		v1958 = v1926
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v1941 = v1926 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1941
	if v1941 <= v1927 {
		v1958 = v1926
		goto L467
	} else {
		goto L470
	}
L470:
	;
	v1947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1934-int32(2)))))
	if v1947 != int32(106) {
		v1958 = v1926
		goto L467
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1926 - int32(2)
	v1953 = F_slice_del(m, l0)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L107
	} else {
		goto L472
	}
L472:
	;
	if v1953 < int32(0) {
		v2269 = v1953
		goto L104
	} else {
		goto L473
	}
L473:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1958 = v1957
	goto L467
L474:
	;
	if v2091 < int32(0) {
		goto L379
	} else {
		goto L498
	}
L475:
	;
	v2091 = int32(-1)
	goto L474
L476:
	;
	if v1989 <= v1554 {
		goto L475
	} else {
		goto L478
	}
L478:
	;
	v1997 = int32(1)
	v1998 = v1989 - v1997
	v2000 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1976+v1998))))
	v2002 = v2000 & int32(255)
	if v1998 == v1554 {
		v2057 = v2002
		v2058 = v1997
		goto L479
	} else {
		goto L480
	}
L479:
	;
	if int32(246) < v2057 {
		goto L488
	} else {
		goto L489
	}
L480:
	;
	if int32(0) <= v2000 {
		v2057 = v2002
		v2058 = v1997
		goto L479
	} else {
		goto L481
	}
L481:
	;
	v2008 = v2002 & int32(63)
	v2010 = v1989 - int32(2)
	v2012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1976+v2010))))
	v2014 = v2012 << (uint(int32(6)) % 32)
	if base.B2i32(v2010 != v1554)&base.B2i32(base.Ui32(v2012) < base.Ui32(int32(192))) == int32(0) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2057 = v2014&int32(1984) | v2008
	v2058 = int32(2)
	goto L479
L483:
	;
	goto L484
L484:
	;
	v2027 = v2014&int32(4032) | v2008
	v2029 = v1989 - int32(3)
	v2031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1976+v2029))))
	if base.B2i32(v2029 != v1554)&base.B2i32(base.Ui32(v2031) < base.Ui32(int32(224))) == int32(0) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2057 = v2031<<(uint(int32(12))%32)&int32(61440) | v2027
	v2058 = int32(3)
	goto L479
L486:
	;
	goto L487
L487:
	;
	v2049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1989+(v1976-int32(4))))))
	v2057 = v2031<<(uint(int32(12))%32)&int32(258048) | v2049&int32(7)<<(uint(int32(18))%32) | v2027
	v2058 = int32(4)
	goto L479
L488:
	;
	v2091 = v2058
	goto L474
L489:
	;
	goto L490
L490:
	;
	v2062 = v2057 - int32(97)
	if v2062 < int32(0) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v2091 = v2058
	goto L474
L492:
	;
	goto L493
L493:
	;
	v2068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2062)>>(uint(int32(3))%32)))+uint32(_consts[1483]))))
	if int32(base.Ui32(v2068)>>(uint(v2062&int32(7))%32))&int32(1) == int32(0) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2091 = v2058
	goto L474
L495:
	;
	goto L496
L496:
	;
	v2076 = v1989 - v2058
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2076
	v1989 = v2076
	goto L476
L498:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2094
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L501
L499:
	;
	if v2224 != 0 {
		goto L379
	} else {
		goto L523
	}
L500:
	;
	v2224 = v2217
	goto L499
L501:
	;
	if v2094 <= v2113 {
		v2217 = int32(-1)
		goto L500
	} else {
		goto L503
	}
L502:
	;
	v2217 = int32(0)
	goto L500
L503:
	;
	v2130 = int32(1)
	v2131 = v2094 - v2130
	v2133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2109+v2131))))
	v2135 = v2133 & int32(255)
	if v2131 == v2113 {
		v2190 = v2135
		v2191 = v2130
		goto L504
	} else {
		goto L505
	}
L504:
	;
	if int32(122) < v2190 {
		goto L513
	} else {
		goto L514
	}
L505:
	;
	if int32(0) <= v2133 {
		v2190 = v2135
		v2191 = v2130
		goto L504
	} else {
		goto L506
	}
L506:
	;
	v2141 = v2135 & int32(63)
	v2143 = v2094 - int32(2)
	v2145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109+v2143))))
	v2147 = v2145 << (uint(int32(6)) % 32)
	if base.B2i32(v2143 != v2113)&base.B2i32(base.Ui32(v2145) < base.Ui32(int32(192))) == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2190 = v2147&int32(1984) | v2141
	v2191 = int32(2)
	goto L504
L508:
	;
	goto L509
L509:
	;
	v2160 = v2147&int32(4032) | v2141
	v2162 = v2094 - int32(3)
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2109+v2162))))
	if base.B2i32(v2162 != v2113)&base.B2i32(base.Ui32(v2164) < base.Ui32(int32(224))) == int32(0) {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2190 = v2164<<(uint(int32(12))%32)&int32(61440) | v2160
	v2191 = int32(3)
	goto L504
L511:
	;
	goto L512
L512:
	;
	v2182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094+(v2109-int32(4))))))
	v2190 = v2164<<(uint(int32(12))%32)&int32(258048) | v2182&int32(7)<<(uint(int32(18))%32) | v2160
	v2191 = int32(4)
	goto L504
L513:
	;
	v2224 = v2191
	goto L499
L514:
	;
	goto L515
L515:
	;
	v2195 = v2190 - int32(98)
	if v2195 < int32(0) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v2224 = v2191
	goto L499
L517:
	;
	goto L518
L518:
	;
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2195)>>(uint(int32(3))%32)))+uint32(_consts[1485]))))
	if int32(base.Ui32(v2201)>>(uint(v2195&int32(7))%32))&int32(1) == int32(0) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v2224 = v2191
	goto L499
L520:
	;
	goto L521
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2094 - v2191
	goto L522
L522:
	;
	goto L502
L523:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2225
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2227)))
	v2229 = F_slice_to(m, l0, v2228)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L107
	} else {
		goto L524
	}
L524:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2231))) = v2229
	if v2229 == int32(0) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	return int32(-1)
L526:
	;
	goto L527
L527:
	;
	v2237 = int32(0)
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2229-int32(4))))
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2243-v2244 < v2242 {
		v2254 = v2237
		goto L529
	} else {
		goto L530
	}
L528:
	;
	if v2254 == int32(0) {
		goto L379
	} else {
		goto L532
	}
L529:
	;
	goto L528
L530:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2250 = F_memcmp(m, v2247+v2243-v2242, v2229, v2242)
	mBase = m.M
	if v2250 != 0 {
		v2254 = v2237
		goto L529
	} else {
		goto L531
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2243 - v2242
	v2254 = int32(1)
	goto L529
L532:
	;
	v2257 = F_slice_del(m, l0)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L107
	} else {
		goto L533
	}
L533:
	;
	if v2257 < int32(0) {
		v2269 = v2257
		goto L104
	} else {
		goto L534
	}
L534:
	;
	goto L379
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
	var v36 int32
	_ = v36
	v3 = int32(0)
	if l0 == v3 {
		v36 = v3
		return v36
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
					v36 = v34
					return v36
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
						v36 = v3
						return v36
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
					v36 = v3
					return v36
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
	var v21 int32
	_ = v21
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v21 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v21 = int32(0)
	}
	return v21
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = base.F64_promote_f32(v5)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_sub(v6, v8)
	if base.F64_ne(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v23 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	} else {
		if base.F64_eq(base.F64_abs(v6), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v23 = F_Float8GetDatum(m, v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		} else {
			if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v23 = F_Float8GetDatum(m, v9)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
func F_float4mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v20 int32
	_ = v20
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = base.F32_sub(v5, v6)
	if base.F32_ne(base.F32_abs(v7), math.Float32frombits(uint32(0x7f800000))) != 0 {
		return base.I32_reinterpret_f32(v7)
	} else {
		if base.F32_eq(base.F32_abs(v5), math.Float32frombits(uint32(0x7f800000))) != 0 {
			return base.I32_reinterpret_f32(v7)
		} else {
			if base.F32_eq(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000))) != 0 {
				return base.I32_reinterpret_f32(v7)
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
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
func F_float4out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 float32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v412 int64
	_ = v412
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v685 int32
	_ = v685
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v728 int32
	_ = v728
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int64
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	v2 = int32(0)
	v21 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_palloc(m, int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[1268]))
	if int32(0) < v28 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = base.I32_reinterpret_f32(v21)
	v33 = v31 & int32(8388607)
	v36 = int32(255)
	v37 = int32(base.Ui32(v31)>>(uint(int32(23))%32)) & v36
	if v33|v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	F_pg_strfromd(m, v23, v28+int32(6), base.F64_promote_f32(v21))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L173
	}
L6:
	;
	v1154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1152+v23))) = uint8(v1154)
	return v23
L7:
	;
	v42 = base.B2i32(v37 != v36)
	goto L9
L8:
	;
	v42 = int32(0)
	goto L9
L9:
	;
	if v42 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(int32(23)) < base.Ui32(v37-int32(127)) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1269])))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v46)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1270])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23))) = uint16(v49)
	v1152 = int32(3)
	goto L6
L14:
	;
	goto L15
L15:
	;
	if v31 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v54)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v58 = v23 + int32(base.Ui32(v31)>>(uint(int32(31))%32))
	if v37 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = int64(8751735898823355977)
	if v31 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v66 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v66)
	if v31 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v65 = int32(9)
	goto L24
L23:
	;
	v65 = int32(8)
	goto L24
L24:
	;
	v1152 = v65
	goto L6
L25:
	;
	v72 = int32(2)
	goto L27
L26:
	;
	v72 = int32(1)
	goto L27
L27:
	;
	v1152 = v72
	goto L6
L28:
	;
	v737 = v736 + v728
	v738 = int32(0)
	if v31 < v738 {
		goto L98
	} else {
		goto L99
	}
L29:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(8)
		goto L28
	} else {
		goto L89
	}
L30:
	;
	v90 = int32(2)
	v96 = v33 << (uint(v90) % 32)
	if v37 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v77 = int32(-1)
	v79 = int32(150) - v37
	if v33&(v77<<(uint(v79)%32)^v77) != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v674 = int32(base.Ui32(v33|int32(8388608)) >> (uint(v79) % 32))
	v685 = v2
	goto L29
L33:
	;
	v99 = v96 | int32(33554432)
	goto L35
L34:
	;
	v99 = v96
	goto L35
L35:
	;
	v100 = base.B2i32(v33 != int32(0)) | base.B2i32(base.Ui32(v37) < base.Ui32(v90)) ^ int32(-1) + v99
	v102 = v99 | int32(2)
	if v37 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v668 = v648 + v653
	v669 = v667 + v650
	if base.Ui32(v669) <= base.Ui32(int32(99999999)) {
		v674 = v669
		v685 = v668
		goto L29
	} else {
		goto L88
	}
L37:
	;
	v583 = int32(0)
	v584 = int32(10)
	v585 = base.I32_div_u_s(v571, v584)
	v587 = base.I32_div_u_s(v574, v584)
	if base.Ui32(v587) < base.Ui32(v585) {
		goto L82
	} else {
		goto L83
	}
L38:
	;
	v484 = int32(0)
	v485 = int32(10)
	v486 = base.I32_div_u_s(v472, v485)
	v488 = base.I32_div_u_s(v475, v485)
	if base.Ui32(v486) <= base.Ui32(v488) {
		goto L76
	} else {
		goto L77
	}
L39:
	;
	v106 = v37 - int32(152)
	goto L41
L40:
	;
	v106 = int32(-151)
	goto L41
L41:
	;
	if int32(0) <= v106 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v112 = int32(base.Ui32(v106*int32(78913)) >> (uint(int32(18)) % 32))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v112<<(uint(int32(3))%32))+uint32(_consts[1271])))
	v119 = v117 & int64(4294967295)
	v120 = base.I64_extend_i32_u(v100)
	v122 = int64(32)
	v124 = base.I32_wrap_i64(int64(base.Ui64(v119*v120) >> (uint(v122) % 64)))
	v126 = int64(base.Ui64(v117) >> (uint(v122) % 64))
	v127 = v126 * v120
	v129 = v124 + base.I32_wrap_i64(v127)
	v136 = v112 - v106
	v141 = v136 + int32(base.Ui32(v112*int32(1217359))>>(uint(int32(19))%32))
	v142 = int32(5) - v141
	v145 = v141 + int32(27)
	v147 = (base.B2i32(base.Ui32(v129) < base.Ui32(v124))+base.I32_wrap_i64(int64(base.Ui64(v127)>>(uint(v122)%64))))<<(uint(v142)%32) | int32(base.Ui32(v129)>>(uint(v145)%32))
	v148 = base.I64_extend_i32_u(v102)
	v152 = base.I32_wrap_i64(int64(base.Ui64(v119*v148) >> (uint(v122) % 64)))
	v153 = v148 * v126
	v155 = v152 + base.I32_wrap_i64(v153)
	v163 = (base.B2i32(base.Ui32(v155) < base.Ui32(v152))+base.I32_wrap_i64(int64(base.Ui64(v153)>>(uint(v122)%64))))<<(uint(v142)%32) | int32(base.Ui32(v155)>>(uint(v145)%32))
	v164 = base.I64_extend_i32_u(v99)
	v168 = base.I32_wrap_i64(int64(base.Ui64(v119*v164) >> (uint(v122) % 64)))
	v169 = v164 * v126
	v171 = v168 + base.I32_wrap_i64(v169)
	v179 = (base.B2i32(base.Ui32(v171) < base.Ui32(v168))+base.I32_wrap_i64(int64(base.Ui64(v169)>>(uint(v122)%64))))<<(uint(v142)%32) | int32(base.Ui32(v171)>>(uint(v145)%32))
	v180 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v106) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v328 = v106 * int32(-732923)
	v330 = int32(base.Ui32(v328) >> (uint(int32(20)) % 32))
	v331 = v330 + v106
	v335 = *(*int64)(unsafe.Add(mBase, uint32(int32(1843744)-v331<<(uint(int32(3))%32))))
	v337 = v335 & int64(4294967295)
	v338 = base.I64_extend_i32_u(v100)
	v340 = int64(32)
	v342 = base.I32_wrap_i64(int64(base.Ui64(v337*v338) >> (uint(v340) % 64)))
	v344 = int64(base.Ui64(v335) >> (uint(v340) % 64))
	v345 = v344 * v338
	v347 = v342 + base.I32_wrap_i64(v345)
	v358 = v330 - int32(base.Ui32(v331*int32(-1217359))>>(uint(int32(19))%32))
	v359 = int32(4) - v358
	v362 = v358 + int32(28)
	v364 = (base.B2i32(base.Ui32(v347) < base.Ui32(v342))+base.I32_wrap_i64(int64(base.Ui64(v345)>>(uint(v340)%64))))<<(uint(v359)%32) | int32(base.Ui32(v347)>>(uint(v362)%32))
	v365 = base.I64_extend_i32_u(v99)
	v369 = base.I32_wrap_i64(int64(base.Ui64(v337*v365) >> (uint(v340) % 64)))
	v370 = v365 * v344
	v372 = v369 + base.I32_wrap_i64(v370)
	v380 = (base.B2i32(base.Ui32(v372) < base.Ui32(v369))+base.I32_wrap_i64(int64(base.Ui64(v370)>>(uint(v340)%64))))<<(uint(v359)%32) | int32(base.Ui32(v372)>>(uint(v362)%32))
	v381 = base.I64_extend_i32_u(v102)
	v385 = base.I32_wrap_i64(int64(base.Ui64(v337*v381) >> (uint(v340) % 64)))
	v386 = v344 * v381
	v388 = v385 + base.I32_wrap_i64(v386)
	v396 = (base.B2i32(base.Ui32(v388) < base.Ui32(v385))+base.I32_wrap_i64(int64(base.Ui64(v386)>>(uint(v340)%64))))<<(uint(v359)%32) | int32(base.Ui32(v388)>>(uint(v362)%32))
	v398 = v396 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v328) {
		goto L66
	} else {
		goto L67
	}
L45:
	;
	v186 = int32(10)
	v187 = base.I32_div_u_s(v163-int32(1), v186)
	v189 = base.I32_div_u_s(v147, v186)
	if base.Ui32(v187) <= base.Ui32(v189) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v234 = v180
	goto L47
L47:
	;
	v239 = base.I32_rem_u_s(v99, int32(5))
	if v239 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v192 = v112 - int32(1)
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v192<<(uint(int32(3))%32))+uint32(_consts[1271])))
	v201 = int64(32)
	v203 = base.I32_wrap_i64(int64(base.Ui64(v197&int64(4294967295)*v164) >> (uint(v201) % 64)))
	v206 = int64(base.Ui64(v197)>>(uint(v201)%64)) * v164
	v208 = v203 + base.I32_wrap_i64(v206)
	v219 = v136 + int32(base.Ui32(v192*int32(1217359))>>(uint(int32(19))%32))
	v227 = base.I32_rem_u_s((base.B2i32(base.Ui32(v208) < base.Ui32(v203))+base.I32_wrap_i64(int64(base.Ui64(v206)>>(uint(v201)%64))))<<(uint(int32(6)-v219)%32)|int32(base.Ui32(v208)>>(uint(v219+int32(26))%32)), int32(10))
	v228 = v227
	goto L50
L49:
	;
	v228 = v180
	goto L50
L50:
	;
	if base.Ui32(int32(33)) < base.Ui32(v106) {
		v565 = v179
		v567 = v228
		v569 = v112
		v571 = v163
		v574 = v147
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v234 = v228
	goto L47
L52:
	;
	v243 = v99
	v245 = v180
	goto L55
L53:
	;
	goto L54
L54:
	;
	v271 = int32(0)
	v273 = base.I32_rem_u_s(v102, int32(5))
	if v273 == v271 {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v263 = v245 + int32(1)
	v264 = int32(5)
	v265 = base.I32_div_u_s(v243, v264)
	v267 = base.I32_rem_u_s(v265, v264)
	if v267 == int32(0) {
		v243 = v265
		v245 = v263
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if base.Ui32(v263) < base.Ui32(v112) {
		v565 = v179
		v567 = v234
		v569 = v112
		v571 = v163
		v574 = v147
		goto L37
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v466 = v179
	v468 = v234
	v470 = v112
	v472 = v163
	v475 = v147
	goto L38
L59:
	;
	v277 = v271
	v281 = v102
	goto L62
L60:
	;
	v305 = v271
	goto L61
L61:
	;
	v565 = v179
	v567 = v234
	v569 = v112
	v571 = v163 - base.B2i32(base.Ui32(v112) <= base.Ui32(v305))
	v574 = v147
	goto L37
L62:
	;
	v297 = v277 + int32(1)
	v298 = int32(5)
	v299 = base.I32_div_u_s(v281, v298)
	v301 = base.I32_rem_u_s(v299, v298)
	if v301 == int32(0) {
		v277 = v297
		v281 = v299
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v305 = v297
	goto L61
L64:
	;
	goto L63
L65:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v328) {
		v565 = v380
		v567 = v445
		v569 = v331
		v571 = v396
		v574 = v364
		goto L37
	} else {
		goto L73
	}
L66:
	;
	v401 = int32(10)
	v402 = base.I32_div_u_s(v398, v401)
	v404 = base.I32_div_u_s(v364, v401)
	if base.Ui32(v402) <= base.Ui32(v404) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v451 = v2
	goto L68
L68:
	;
	v466 = v380
	v468 = v451
	v470 = v331
	v472 = v398
	v475 = v364
	goto L38
L69:
	;
	v407 = int32(1) - v331
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v407<<(uint(int32(3))%32))+uint32(_consts[1272])))
	v416 = int64(32)
	v418 = base.I32_wrap_i64(int64(base.Ui64(v412&int64(4294967295)*v365) >> (uint(v416) % 64)))
	v421 = int64(base.Ui64(v412)>>(uint(v416)%64)) * v365
	v423 = v418 + base.I32_wrap_i64(v421)
	v436 = v330 + (int32(base.Ui32(v407*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v444 = base.I32_rem_u_s((base.B2i32(base.Ui32(v423) < base.Ui32(v418))+base.I32_wrap_i64(int64(base.Ui64(v421)>>(uint(v416)%64))))<<(uint(int32(4)-v436)%32)|int32(base.Ui32(v423)>>(uint(v436+int32(28))%32)), int32(10))
	v445 = v444
	goto L71
L70:
	;
	v445 = v2
	goto L71
L71:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v328) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v451 = v445
	goto L68
L73:
	;
	v457 = int32(-1)
	if v99&(v457<<(uint(v330-int32(1))%32)^v457) != 0 {
		v565 = v380
		v567 = v445
		v569 = v331
		v571 = v396
		v574 = v364
		goto L37
	} else {
		goto L74
	}
L74:
	;
	v466 = v380
	v468 = v445
	v470 = v331
	v472 = v396
	v475 = v364
	goto L38
L75:
	;
	v553 = v536 & int32(255)
	v648 = v532
	v650 = v534
	v653 = v470
	v667 = (v551|base.B2i32(v553 != int32(5))|v534)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v553)) | base.B2i32(v534 == v542)
	goto L36
L76:
	;
	v532 = v484
	v534 = v466
	v536 = v468
	v542 = v475
	v551 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v493 = v484
	v494 = v466
	v496 = v468
	v499 = v486
	v500 = v488
	v502 = int32(1)
	goto L79
L79:
	;
	v513 = v493 + int32(1)
	v514 = int32(10)
	v515 = base.I32_div_u_s(v494, v514)
	v518 = v494 - v515*v514
	v523 = v502 & base.B2i32(v496&int32(255) == int32(0))
	v525 = base.I32_div_u_s(v499, v514)
	v527 = base.I32_div_u_s(v500, v514)
	if base.Ui32(v527) < base.Ui32(v525) {
		v493 = v513
		v494 = v515
		v496 = v518
		v499 = v525
		v500 = v527
		v502 = v523
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v532 = v513
	v534 = v515
	v536 = v518
	v542 = v500
	v551 = v523 ^ int32(1)
	goto L75
L81:
	;
	goto L80
L82:
	;
	v590 = v583
	v591 = v565
	v592 = v585
	v594 = v587
	goto L85
L83:
	;
	v622 = v583
	v623 = v565
	v625 = v567
	v632 = v574
	goto L84
L84:
	;
	v648 = v622
	v650 = v623
	v653 = v569
	v667 = base.B2i32(v623 == v632) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v625&int32(255)))
	goto L36
L85:
	;
	v610 = v590 + int32(1)
	v611 = int32(10)
	v612 = base.I32_div_u_s(v591, v611)
	v614 = base.I32_div_u_s(v592, v611)
	v616 = base.I32_div_u_s(v594, v611)
	if base.Ui32(v616) < base.Ui32(v614) {
		v590 = v610
		v591 = v612
		v592 = v614
		v594 = v616
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v622 = v610
	v623 = v612
	v625 = v591 - v612*int32(10)
	v632 = v594
	goto L84
L87:
	;
	goto L86
L88:
	;
	v717 = v669
	v728 = v668
	v736 = int32(9)
	goto L28
L89:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(7)
		goto L28
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(6)
		goto L28
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(5)
		goto L28
	} else {
		goto L92
	}
L92:
	;
	if base.Ui32(int32(999)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(4)
		goto L28
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(int32(99)) < base.Ui32(v674) {
		v717 = v674
		v728 = v685
		v736 = int32(3)
		goto L28
	} else {
		goto L94
	}
L94:
	;
	if base.Ui32(int32(9)) < base.Ui32(v674) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v715 = int32(2)
	goto L97
L96:
	;
	v715 = int32(1)
	goto L97
L97:
	;
	v717 = v674
	v728 = v685
	v736 = v715
	goto L28
L98:
	;
	v741 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v741)
	v744 = int32(1)
	goto L100
L99:
	;
	v744 = v738
	goto L100
L100:
	;
	if base.Ui32(v737+int32(3)) <= base.Ui32(int32(9)) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v969 = int32(0)
	if base.Ui32(v717) < base.Ui32(int32(10000)) {
		goto L143
	} else {
		goto L144
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v749))) = v965
	v967 = v964
	goto L101
L103:
	;
	v749 = v744 + v23
	v750 = int32(0)
	if v737 <= v750 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if v728 != 0 {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v964 = int32(2) - v737
	v965 = int64(3472328296227679792)
	goto L102
L107:
	;
	goto L108
L108:
	;
	if int32(0) <= v728 {
		v964 = v750
		v965 = int64(3472328296227680304)
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v967 = int32(1)
	goto L101
L110:
	;
	v811 = int32(0)
	if base.Ui32(v793) < base.Ui32(int32(10000)) {
		goto L119
	} else {
		goto L120
	}
L111:
	;
	v793 = v717
	v799 = v736
	goto L110
L112:
	;
	goto L113
L113:
	;
	v761 = v717
	v767 = v736
	goto L114
L114:
	;
	if v761&int32(1) != 0 {
		v793 = v761
		v799 = v767
		goto L110
	} else {
		goto L116
	}
L115:
	;
	v793 = v761
	v799 = v767
	goto L110
L116:
	;
	v787 = base.I32_div_u_s(v761, int32(10))
	if int32(0)-v761 == v787*int32(-10) {
		v761 = v787
		v767 = v767 - int32(1)
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if base.Ui32(v870) < base.Ui32(int32(100)) {
		goto L126
	} else {
		goto L127
	}
L119:
	;
	v868 = v811
	v870 = v793
	goto L118
L120:
	;
	goto L121
L121:
	;
	v817 = v811
	v818 = v793
	goto L122
L122:
	;
	v836 = v744 + v23 + v799 - v817
	v840 = base.I32_div_u_s(v818, int32(10000))
	v843 = v840*int32(-10000) + v818
	v844 = int32(100)
	v845 = base.I32_div_u_s(v843, v844)
	v846 = int32(1)
	v850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v845<<(uint(v846)%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v836-int32(3)))) = uint16(v850)
	v861 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v843-v845*v844)<<(uint(v846)%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v836-v846))) = uint16(v861)
	v864 = v817 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v818) {
		v817 = v864
		v818 = v840
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v868 = v864
	v870 = v840
	goto L118
L124:
	;
	goto L123
L125:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v912) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v911 = v868
	v912 = v870
	goto L125
L127:
	;
	goto L128
L128:
	;
	v894 = int32(65535)
	v896 = int32(100)
	v897 = base.I32_div_u_s(v870&v894, v896)
	v907 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v870-v897*v896)&v894<<(uint(int32(1))%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v744+v23+v799+(v868^int32(-1))))) = uint16(v907)
	v911 = v868 | int32(2)
	v912 = v897
	goto L125
L129:
	;
	v932 = v737 - int32(1)
	v933 = v744 + v23
	*(*uint8)(unsafe.Add(mBase, uint32(v933))) = uint8(v930)
	if base.Ui32(int32(2)) <= base.Ui32(v799) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v919 = v912 << (uint(int32(1)) % 32)
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+uint32(_consts[1274]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+(v744+v799-v911)))) = uint8(v922)
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+uint32(_consts[1273]))))
	v930 = v926
	goto L129
L131:
	;
	goto L132
L132:
	;
	v930 = v912 | int32(48)
	goto L129
L133:
	;
	v937 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)) = uint8(v937)
	v942 = v799 + int32(1)
	goto L135
L134:
	;
	v942 = int32(1)
	goto L135
L135:
	;
	v943 = v942 + v744
	v944 = v23 + v943
	v945 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v944))) = uint8(v945)
	v950 = base.B2i32(v932 < int32(0))
	if v932 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v951 = int32(45)
	goto L138
L137:
	;
	v951 = int32(43)
	goto L138
L138:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v944)+1)) = uint8(v951)
	if v932 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v955 = int32(1) - v737
	goto L141
L140:
	;
	v955 = v932
	goto L141
L141:
	;
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v955<<(uint(int32(1))%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v944)+2)) = uint16(v960)
	v1152 = v943 + int32(4)
	goto L6
L142:
	;
	if base.Ui32(v1028) < base.Ui32(int32(100)) {
		goto L150
	} else {
		goto L151
	}
L143:
	;
	v1027 = v969
	v1028 = v717
	goto L142
L144:
	;
	goto L145
L145:
	;
	v975 = v717
	v976 = v969
	goto L146
L146:
	;
	v994 = v967 + v749 + v736 - v976
	v995 = int32(4)
	v998 = base.I32_div_u_s(v975, int32(10000))
	v1001 = v998*int32(-10000) + v975
	v1002 = int32(100)
	v1003 = base.I32_div_u_s(v1001, v1002)
	v1004 = int32(1)
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1003<<(uint(v1004)%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v994-v995))) = uint16(v1008)
	v1019 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1001-v1003*v1002)<<(uint(v1004)%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v994-int32(2)))) = uint16(v1019)
	v1022 = v976 + v995
	if base.Ui32(int32(99999999)) < base.Ui32(v975) {
		v975 = v998
		v976 = v1022
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v1027 = v1022
	v1028 = v998
	goto L142
L148:
	;
	goto L147
L149:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1069) {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v1069 = v1028
	v1070 = v1027
	goto L149
L151:
	;
	goto L152
L152:
	;
	v1050 = int32(2)
	v1052 = int32(65535)
	v1054 = int32(100)
	v1055 = base.I32_div_u_s(v1028&v1052, v1054)
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1028-v1055*v1054)&v1052<<(uint(int32(1))%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v967+v749+v736-v1027-v1050))) = uint16(v1065)
	v1069 = v1055
	v1070 = v1027 | v1050
	goto L149
L153:
	;
	v1088 = int32(1)
	if v967 == v1088 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v1082 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1069<<(uint(int32(1))%32))+uint32(_consts[1273]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v967+v749+v736-v1070-int32(2)))) = uint16(v1082)
	goto L153
L155:
	;
	goto L156
L156:
	;
	v1086 = v1069 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v967+v749))) = uint8(v1086)
	goto L153
L157:
	;
	v1152 = v1128 + int32(base.Ui32(v31)>>(uint(int32(31))%32))
	goto L6
L158:
	;
	if v737&int32(4) != 0 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	if v728 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L161:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v749)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v1093
	v1096 = int32(5)
	goto L163
L162:
	;
	v1096 = v1088
	goto L163
L163:
	;
	if v737&int32(2) != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1099 = v1096 + v749
	v1102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1099-int32(1)))) = uint16(v1102)
	v1107 = v1096 | int32(2)
	goto L166
L165:
	;
	v1107 = v1096
	goto L166
L166:
	;
	if v737&int32(1) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1110 = v1107 + v749
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1110-int32(1)))) = uint8(v1113)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v1117 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v749+v737))) = uint8(v1117)
	v1128 = v736 + int32(1)
	goto L157
L170:
	;
	v1125 = int32(2) - v728
	goto L172
L171:
	;
	v1125 = v737
	goto L172
L172:
	;
	v1128 = v1125
	goto L157
L173:
	;
	return v23
}
func F_float4smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float32
	_ = v5
	var v15 float32
	_ = v15
	var v17 float32
	_ = v17
	var v18 float32
	_ = v18
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(base.I32_reinterpret_f32(v5)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(base.F32_abs(v4))) {
			v15 = v5
		} else {
			v15 = v4
		}
		if base.F32_gt(v4, v5) != 0 {
			v17 = v5
		} else {
			v17 = v15
		}
		v18 = v17
	} else {
		v18 = v4
	}
	return base.I32_reinterpret_f32(v18)
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = base.F64_promote_f32(v7)
	v9 = base.F64_sub(v6, v8)
	if base.F64_ne(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v23 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	} else {
		if base.F64_eq(base.F64_abs(v6), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v23 = F_Float8GetDatum(m, v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		} else {
			if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v23 = F_Float8GetDatum(m, v9)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
		v23 = base.F64_ge(v7, v14) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v23 = int32(1)
	}
	return v23
}
func F_float8lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
		v22 = base.F64_lt(v6, v13) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)))
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_mul(v6, v8)
	if base.F64_ne(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		if base.F64_ne(v9, float64(0)) != 0 {
			v25 = F_Float8GetDatum(m, v9)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				return v25
			}
		} else {
			if base.F64_eq(v6, float64(0)) != 0 {
				v25 = F_Float8GetDatum(m, v9)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					return v25
				}
			} else {
				if base.F64_ne(v8, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v25 = F_Float8GetDatum(m, v9)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v25
					}
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v6), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_ne(v9, float64(0)) != 0 {
				v25 = F_Float8GetDatum(m, v9)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					return v25
				}
			} else {
				if base.F64_eq(v6, float64(0)) != 0 {
					v25 = F_Float8GetDatum(m, v9)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v25
					}
				} else {
					if base.F64_ne(v8, float64(0)) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v25 = F_Float8GetDatum(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							return v25
						}
					}
				}
			}
		} else {
			if base.F64_ne(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_ne(v9, float64(0)) != 0 {
					v25 = F_Float8GetDatum(m, v9)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v25
					}
				} else {
					if base.F64_eq(v6, float64(0)) != 0 {
						v25 = F_Float8GetDatum(m, v9)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							return v25
						}
					} else {
						if base.F64_ne(v8, float64(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v25 = F_Float8GetDatum(m, v9)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								return v25
							}
						}
					}
				}
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
		v10 = *(*int32)(unsafe.Add(mBase, _consts[1268]))
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
func F_float_overflow_error(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(50331778))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(31496), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(493244), int32(90), int32(211410))
				v16 = m.ExcPending
				if v16 != 0 {
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
func F_float_zero_divide_error(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(33816706))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(239436), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errfinish(m, int32(493244), int32(106), int32(211880))
				v16 = m.ExcPending
				if v16 != 0 {
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
func F_fmod(m *base.Module, l0 float64) float64 {
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 float64
	_ = v16
	var v20 int64
	_ = v20
	var v27 float64
	_ = v27
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v78 int64
	_ = v78
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v142 int64
	_ = v142
	v7 = base.I64_reinterpret_f64(l0)
	v11 = int32(2047)
	v12 = base.I32_wrap_i64(int64(base.Ui64(v7)>>(uint(int64(52))%64))) & v11
	if v12 != v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v7 << (uint(int64(1)) % 64)
	if base.Ui64(v20) <= base.Ui64(int64(-9156662467374350336)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v16 = base.F64_mul(l0, float64(360))
	return base.F64_div(v16, v16)
L3:
	;
	if v20 == int64(-9156662467374350336) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v12 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v27 = base.F64_mul(l0, float64(0))
	goto L8
L7:
	;
	v27 = l0
	goto L8
L8:
	;
	return v27
L9:
	;
	if int32(1031) < v63 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	v31 = int32(0)
	v33 = v7 << (uint(int64(12)) % 64)
	if int64(0) <= v33 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v63 = v12
	v68 = v7&int64(4503599627370495) | int64(4503599627370496)
	goto L9
L13:
	;
	v37 = v31
	v39 = v33
	goto L16
L14:
	;
	v49 = v31
	goto L15
L15:
	;
	v63 = v49
	v68 = v7 << (uint(base.I64_extend_i32_u(int32(1)-v49)) % 64)
	goto L9
L16:
	;
	v43 = v37 - int32(1)
	v45 = v39 << (uint(int64(1)) % 64)
	if int64(0) <= v45 {
		v37 = v43
		v39 = v45
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v49 = v43
	goto L15
L18:
	;
	goto L17
L19:
	;
	v72 = v63
	v74 = v68
	goto L22
L20:
	;
	v95 = v63
	v97 = v68
	goto L21
L21:
	;
	v101 = v97 - int64(6333186975989760)
	if v101 < int64(0) {
		v109 = v97
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v78 = v74 - int64(6333186975989760)
	if v78 < int64(0) {
		v86 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v95 = int32(1031)
	v97 = v88
	goto L21
L24:
	;
	v88 = v86 << (uint(int64(1)) % 64)
	v90 = v72 - int32(1)
	if int32(1031) < v90 {
		v72 = v90
		v74 = v88
		goto L22
	} else {
		goto L27
	}
L25:
	;
	if v78 != int64(0) {
		v86 = v78
		goto L24
	} else {
		goto L26
	}
L26:
	;
	return base.F64_mul(l0, float64(0))
L27:
	;
	goto L23
L28:
	;
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v109) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v101 != int64(0) {
		v109 = v101
		goto L28
	} else {
		goto L30
	}
L30:
	;
	return base.F64_mul(l0, float64(0))
L31:
	;
	if int32(0) < v125 {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v125 = v95
	v128 = v109
	goto L31
L33:
	;
	goto L34
L34:
	;
	v113 = v95
	v115 = v109
	goto L35
L35:
	;
	v119 = v113 - int32(1)
	v123 = v115 << (uint(int64(1)) % 64)
	if base.Ui64(v115) < base.Ui64(int64(2251799813685248)) {
		v113 = v119
		v115 = v123
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v125 = v119
	v128 = v123
	goto L31
L37:
	;
	goto L36
L38:
	;
	v142 = v128 - int64(4503599627370496) | base.I64_extend_i32_u(v125)<<(uint(int64(52))%64)
	goto L40
L39:
	;
	v142 = int64(base.Ui64(v128) >> (uint(base.I64_extend_i32_u(int32(1)-v125)) % 64))
	goto L40
L40:
	;
	return base.F64_reinterpret_i64(v142 | v7&int64(-9223372036854775807-1))
}
func F_fmt_fp(m *base.Module, l0 int32, l1 float64, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 float64
	_ = v175
	var v181 int32
	_ = v181
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 float64
	_ = v216
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v283 int64
	_ = v283
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v295 int64
	_ = v295
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v340 int32
	_ = v340
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v729 int32
	_ = v729
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v764 float64
	_ = v764
	var v770 int32
	_ = v770
	var v777 float64
	_ = v777
	var v782 float64
	_ = v782
	var v785 int32
	_ = v785
	var v787 float64
	_ = v787
	var v789 float64
	_ = v789
	var v790 int32
	_ = v790
	var v795 float64
	_ = v795
	var v796 float64
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v985 int32
	_ = v985
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1090 int32
	_ = v1090
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int64
	_ = v1196
	var v1203 int64
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1211 int64
	_ = v1211
	var v1212 int64
	_ = v1212
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1227 int64
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1264 int32
	_ = v1264
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1299 int32
	_ = v1299
	var v1322 int32
	_ = v1322
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1342 int32
	_ = v1342
	var v1350 int32
	_ = v1350
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1396 int32
	_ = v1396
	var v1415 int64
	_ = v1415
	var v1422 int64
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1429 int32
	_ = v1429
	var v1430 int64
	_ = v1430
	var v1431 int64
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1446 int64
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1484 int32
	_ = v1484
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1524 int32
	_ = v1524
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1587 int64
	_ = v1587
	var v1594 int64
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1601 int32
	_ = v1601
	var v1602 int64
	_ = v1602
	var v1603 int64
	_ = v1603
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1618 int64
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1655 int32
	_ = v1655
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1690 int32
	_ = v1690
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1763 int64
	_ = v1763
	var v1770 int64
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1777 int32
	_ = v1777
	var v1778 int64
	_ = v1778
	var v1779 int64
	_ = v1779
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1794 int64
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1838 int32
	_ = v1838
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1878 int32
	_ = v1878
	var v1885 int32
	_ = v1885
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1929 int32
	_ = v1929
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1962 int32
	_ = v1962
	var v1985 int32
	_ = v1985
	var v1990 int32
	_ = v1990
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2034 int32
	_ = v2034
	var v2046 int32
	_ = v2046
	var v2064 float64
	_ = v2064
	var v2069 float64
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2082 float64
	_ = v2082
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2114 int64
	_ = v2114
	var v2121 int64
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2128 int32
	_ = v2128
	var v2129 int64
	_ = v2129
	var v2130 int64
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2141 int32
	_ = v2141
	var v2145 int64
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2168 int32
	_ = v2168
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2196 int32
	_ = v2196
	var v2200 int32
	_ = v2200
	var v2208 float64
	_ = v2208
	var v2214 int32
	_ = v2214
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2249 float64
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2293 int32
	_ = v2293
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2323 int32
	_ = v2323
	v7 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(560)
	m.G0 = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v7
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
	v55 = int32(29763)
	v56 = v7
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
	v55 = int32(29766)
	v56 = v7
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
	v50 = int32(29769)
	goto L10
L9:
	;
	v50 = int32(29764)
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
	return v2323
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
	v94 = v31 + int32(16)
	v96 = v31 + int32(44)
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
	v77 = int32(282543)
	goto L20
L19:
	;
	v77 = int32(530557)
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
	v80 = int32(338547)
	goto L23
L22:
	;
	v80 = int32(536739)
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
	F_pad(m, l0, int32(32), l2, v64, l4^int32(8192))
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
	v2323 = v92
	goto L11
L32:
	;
	v2034 = v55 + l5<<(uint(int32(26))%32)>>(uint(int32(31))%32)&int32(9)
	if base.Ui32(int32(11)) < base.Ui32(l3) {
		v2082 = v132
		goto L370
	} else {
		goto L371
	}
L33:
	;
	v169 = int32(0)
	if v169 <= v163 {
		goto L56
	} else {
		goto L57
	}
L34:
	;
	v153 = v135 - int32(29)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v153
	if l3 < int32(0) {
		goto L53
	} else {
		goto L54
	}
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
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v135 - int32(1)
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
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	if l3 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v151 = int32(6)
	goto L52
L51:
	;
	v151 = l3
	goto L52
L52:
	;
	v161 = v132
	v163 = v147
	v164 = v144
	v165 = v151
	goto L33
L53:
	;
	v160 = int32(6)
	goto L55
L54:
	;
	v160 = l3
	goto L55
L55:
	;
	v161 = base.F64_mul(v132, float64(2.68435456e+08))
	v163 = v153
	v164 = v140
	v165 = v160
	goto L33
L56:
	;
	v172 = int32(288)
	goto L58
L57:
	;
	v172 = v169
	goto L58
L58:
	;
	v173 = v31 + int32(48) + v172
	v175 = v161
	v181 = v173
	goto L59
L59:
	;
	if base.F64_lt(v175, float64(4.294967296e+09))&base.F64_ge(v175, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if v163 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v209
	v212 = v181 + int32(4)
	v216 = base.F64_mul(base.F64_sub(v175, base.F64_convert_i32_u(v209)), float64(1e+09))
	if base.F64_ne(v216, float64(0)) != 0 {
		v175 = v216
		v181 = v212
		goto L59
	} else {
		goto L65
	}
L62:
	;
	v207 = base.I32_trunc_f64_u(v175)
	v209 = v207
	goto L61
L63:
	;
	goto L64
L64:
	;
	v209 = int32(0)
	goto L61
L65:
	;
	goto L60
L66:
	;
	if v381 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L67:
	;
	v379 = v212
	v381 = v163
	v382 = v173
	goto L66
L68:
	;
	goto L69
L69:
	;
	v228 = v212
	v229 = v163
	v230 = v173
	goto L70
L70:
	;
	v249 = int32(29)
	if base.Ui32(v249) <= base.Ui32(v229) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v379 = v340
	v381 = v369
	v382 = v314
	goto L66
L72:
	;
	v252 = v249
	goto L74
L73:
	;
	v252 = v229
	goto L74
L74:
	;
	v254 = v228 - int32(4)
	if base.Ui32(v254) < base.Ui32(v230) {
		v314 = v230
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v340 = v228
	goto L81
L76:
	;
	v264 = v254
	v283 = int64(0)
	goto L77
L77:
	;
	v288 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v264))))
	v290 = v283&int64(4294967295) + v288<<(uint(base.I64_extend_i32_u(v252))%64)
	v291 = int64(1000000000)
	v292 = base.I64_div_u_s(v290, v291)
	v295 = v290 - v292*v291
	*(*uint32)(unsafe.Add(mBase, uint32(v264))) = uint32(v295)
	v298 = v264 - int32(4)
	if base.Ui32(v230) <= base.Ui32(v298) {
		v264 = v298
		v283 = v292
		goto L77
	} else {
		goto L79
	}
L78:
	;
	if base.Ui64(v290) < base.Ui64(int64(1000000000)) {
		v314 = v230
		goto L75
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v303 = v230 - int32(4)
	*(*uint32)(unsafe.Add(mBase, uint32(v303))) = uint32(v292)
	v314 = v303
	goto L75
L81:
	;
	if base.Ui32(v314) < base.Ui32(v340) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v369 = v368 - v252
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v369
	if int32(0) < v369 {
		v228 = v340
		v229 = v369
		v230 = v314
		goto L70
	} else {
		goto L87
	}
L83:
	;
	v363 = v340 - int32(4)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	if v364 == int32(0) {
		v340 = v363
		goto L81
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	goto L85
L87:
	;
	goto L71
L88:
	;
	v406 = base.I32_div_u_s(v165+int32(25), int32(9))
	v408 = v406 + int32(1)
	v417 = v379
	v419 = v381
	v420 = v382
	goto L91
L89:
	;
	v554 = v379
	v557 = v382
	v562 = v7
	goto L90
L90:
	;
	if base.Ui32(v554) <= base.Ui32(v557) {
		v628 = int32(0)
		goto L111
	} else {
		goto L112
	}
L91:
	;
	v439 = int32(9)
	v441 = int32(0) - v419
	if base.Ui32(v439) <= base.Ui32(v441) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v554 = v545
	v557 = v536
	v562 = v408
	goto L90
L93:
	;
	v444 = v439
	goto L95
L94:
	;
	v444 = v441
	goto L95
L95:
	;
	if base.Ui32(v417) <= base.Ui32(v420) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v534 = v533 + v444
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = v534
	v536 = v512 + v420
	if v164 == int32(102) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v511 = v417
	v512 = base.B2i32(v446 == int32(0)) << (uint(int32(2)) % 32)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v453 = int32(-1)
	v465 = v420
	v466 = int32(0)
	goto L100
L100:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	*(*int32)(unsafe.Add(mBase, uint32(v465))) = int32(base.Ui32(v486)>>(uint(v444)%32)) + v466
	v491 = v486 & (v453<<(uint(v444)%32) ^ v453) * int32(base.Ui32(int32(1000000000))>>(uint(v444)%32))
	v493 = v465 + int32(4)
	if base.Ui32(v493) < base.Ui32(v417) {
		v465 = v493
		v466 = v491
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v496 = int32(0)
	v499 = base.B2i32(v495 == v496) << (uint(int32(2)) % 32)
	if v491 == v496 {
		v511 = v417
		v512 = v499
		goto L96
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = v491
	v511 = v417 + int32(4)
	v512 = v499
	goto L96
L104:
	;
	v537 = v173
	goto L106
L105:
	;
	v537 = v536
	goto L106
L106:
	;
	v538 = int32(2)
	if v408 < (v511-v537)>>(uint(v538)%32) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v545 = v537 + v408<<(uint(v538)%32)
	goto L109
L108:
	;
	v545 = v511
	goto L109
L109:
	;
	if v534 < int32(0) {
		v417 = v545
		v419 = v534
		v420 = v536
		goto L91
	} else {
		goto L110
	}
L110:
	;
	goto L92
L111:
	;
	if v164 != int32(102) {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v582 = (v173 - v557) >> (uint(int32(2)) % 32) * int32(9)
	v583 = int32(10)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	if base.Ui32(v584) < base.Ui32(v583) {
		v628 = v582
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v594 = v583
	v595 = v582
	goto L114
L114:
	;
	v616 = v595 + int32(1)
	v618 = v594 * int32(10)
	if base.Ui32(v618) <= base.Ui32(v584) {
		v594 = v618
		v595 = v616
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v628 = v616
	goto L111
L116:
	;
	goto L115
L117:
	;
	v651 = v628
	goto L119
L118:
	;
	v651 = int32(0)
	goto L119
L119:
	;
	v658 = v165 - v651 - base.B2i32(v164 == int32(103))&base.B2i32(v165 != int32(0))
	v662 = int32(9)
	if v658 < (v554-v173)>>(uint(int32(2))%32)*v662-v662 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if v163 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v957 = v554
	v959 = v628
	v960 = v557
	v965 = v562
	goto L122
L122:
	;
	v985 = v957
	goto L170
L123:
	;
	v673 = int32(-4092)
	goto L125
L124:
	;
	v673 = int32(-3804)
	goto L125
L125:
	;
	v676 = v658 + int32(9216)
	v677 = int32(9)
	v678 = base.I32_div_s(v676, v677)
	v681 = v31 + int32(48) + v673 + v678<<(uint(int32(2))%32)
	v682 = int32(10)
	v685 = v676 - v678*v677
	if v685 <= int32(7) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v691 = v685
	v695 = v682
	goto L129
L127:
	;
	v729 = v682
	goto L128
L128:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v751 = base.I32_div_u_s(v750, v729)
	v753 = v750 - v751*v729
	v757 = v681 + int32(4)
	if base.B2i32(v753 == int32(0))&base.B2i32(v757 == v554) != 0 {
		v927 = v628
		v928 = v557
		v929 = v681
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v717 = v695 * int32(10)
	v719 = v691 + int32(1)
	if v719 != int32(8) {
		v691 = v719
		v695 = v717
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v729 = v717
	goto L128
L131:
	;
	goto L130
L132:
	;
	v948 = v929 + int32(4)
	if base.Ui32(v948) < base.Ui32(v554) {
		goto L167
	} else {
		goto L168
	}
L133:
	;
	if v751&int32(1) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v554 == v757 {
		goto L141
	} else {
		goto L142
	}
L135:
	;
	v764 = float64(9.007199254740992e+15)
	if v729 != int32(1000000000) {
		v777 = v764
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v777 = float64(9.007199254740994e+15)
	goto L134
L138:
	;
	if base.Ui32(v681) <= base.Ui32(v557) {
		v777 = v764
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681-int32(4)))))
	if v770&int32(1) == int32(0) {
		v777 = v764
		goto L134
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v782 = float64(1)
	goto L143
L142:
	;
	v782 = float64(1.5)
	goto L143
L143:
	;
	v785 = int32(base.Ui32(v729) >> (uint(int32(1)) % 32))
	if v753 == v785 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v787 = v782
	goto L146
L145:
	;
	v787 = float64(1.5)
	goto L146
L146:
	;
	if base.Ui32(v753) < base.Ui32(v785) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v789 = float64(0.5)
	goto L149
L148:
	;
	v789 = v787
	goto L149
L149:
	;
	if v56 != 0 {
		v795 = v777
		v796 = v789
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v797 = v750 - v753
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v797
	if base.F64_eq(base.F64_add(v795, v796), v795) != 0 {
		v927 = v628
		v928 = v557
		v929 = v681
		goto L132
	} else {
		goto L153
	}
L151:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v790 != int32(45) {
		v795 = v777
		v796 = v789
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v795 = base.F64_neg(v777)
	v796 = base.F64_neg(v789)
	goto L150
L153:
	;
	v801 = v797 + v729
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = v801
	if base.Ui32(int32(1000000000)) <= base.Ui32(v801) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v814 = v557
	v815 = v681
	goto L157
L155:
	;
	v858 = v557
	v859 = v681
	goto L156
L156:
	;
	v881 = (v173 - v858) >> (uint(int32(2)) % 32) * int32(9)
	v882 = int32(10)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	if base.Ui32(v883) < base.Ui32(v882) {
		v927 = v881
		v928 = v858
		v929 = v859
		goto L132
	} else {
		goto L163
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v815))) = int32(0)
	v836 = v815 - int32(4)
	if base.Ui32(v836) < base.Ui32(v814) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v858 = v842
	v859 = v836
	goto L156
L159:
	;
	v839 = v814 - int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = int32(0)
	v842 = v839
	goto L161
L160:
	;
	v842 = v814
	goto L161
L161:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v836)))
	v845 = v843 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v836))) = v845
	if base.Ui32(int32(999999999)) < base.Ui32(v845) {
		v814 = v842
		v815 = v836
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v893 = v882
	v894 = v881
	goto L164
L164:
	;
	v915 = v894 + int32(1)
	v917 = v893 * int32(10)
	if base.Ui32(v917) <= base.Ui32(v883) {
		v893 = v917
		v894 = v915
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v927 = v915
	v928 = v858
	v929 = v859
	goto L132
L166:
	;
	goto L165
L167:
	;
	v950 = v948
	goto L169
L168:
	;
	v950 = v554
	goto L169
L169:
	;
	v957 = v950
	v959 = v927
	v960 = v928
	v965 = v751
	goto L122
L170:
	;
	v1007 = base.B2i32(base.Ui32(v985) <= base.Ui32(v960))
	if v1007 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v164 != int32(103) {
		goto L177
	} else {
		goto L178
	}
L172:
	;
	v1011 = v985 - int32(4)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	if v1012 == int32(0) {
		v985 = v1011
		goto L170
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	goto L171
L175:
	;
	goto L174
L176:
	;
	v1170 = int32(-1)
	v1173 = v1153 | v1155
	if v1173 != 0 {
		goto L212
	} else {
		goto L213
	}
L177:
	;
	v1147 = l5
	v1153 = v165
	v1155 = l4 & int32(8)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v1020 = int32(-1)
	if v165 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1024 = v165
	goto L182
L181:
	;
	v1024 = int32(1)
	goto L182
L182:
	;
	v1028 = base.B2i32(v959 < v1024) & base.B2i32(int32(-5) < v959)
	if v1028 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1029 = v959 ^ v1020
	goto L185
L184:
	;
	v1029 = v1020
	goto L185
L185:
	;
	v1030 = v1029 + v1024
	if v1028 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1033 = int32(-1)
	goto L188
L187:
	;
	v1033 = int32(-2)
	goto L188
L188:
	;
	v1034 = v1033 + l5
	v1036 = l4 & int32(8)
	if v1036 != 0 {
		v1147 = v1034
		v1153 = v1030
		v1155 = v1036
		goto L176
	} else {
		goto L189
	}
L189:
	;
	v1037 = int32(-9)
	if base.Ui32(v985) <= base.Ui32(v960) {
		v1090 = v1037
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1116 = (v985 - v173) >> (uint(int32(2)) % 32) * int32(9)
	if v1034&int32(-33) == int32(70) {
		goto L197
	} else {
		goto L198
	}
L191:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v985-int32(4))))
	if v1040 == int32(0) {
		v1090 = v1037
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v1043 = int32(10)
	v1044 = int32(0)
	v1046 = base.I32_rem_u_s(v1040, v1043)
	if v1046 != 0 {
		v1090 = v1044
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v1050 = v1043
	v1053 = v1044
	goto L194
L194:
	;
	v1078 = v1050 * int32(10)
	v1079 = base.I32_rem_u_s(v1040, v1078)
	if v1079 == int32(0) {
		v1050 = v1078
		v1053 = v1053 + int32(1)
		goto L194
	} else {
		goto L196
	}
L195:
	;
	v1090 = v1053 ^ int32(-1)
	goto L190
L196:
	;
	goto L195
L197:
	;
	v1121 = int32(0)
	v1124 = v1116 + v1090 - int32(9)
	if v1121 < v1124 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v1131 = int32(0)
	v1135 = v1116 + v959 + v1090 - int32(9)
	if v1131 < v1135 {
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v1128 = v1124
	goto L202
L201:
	;
	v1128 = v1121
	goto L202
L202:
	;
	if v1030 < v1128 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1130 = v1030
	goto L205
L204:
	;
	v1130 = v1128
	goto L205
L205:
	;
	v1147 = v1034
	v1153 = v1130
	v1155 = v1121
	goto L176
L206:
	;
	v1139 = v1135
	goto L208
L207:
	;
	v1139 = v1131
	goto L208
L208:
	;
	if v1030 < v1139 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1141 = v1030
	goto L211
L210:
	;
	v1141 = v1139
	goto L211
L211:
	;
	v1147 = v1034
	v1153 = v1141
	v1155 = v1131
	goto L176
L212:
	;
	v1174 = int32(2147483645)
	goto L214
L213:
	;
	v1174 = int32(2147483646)
	goto L214
L214:
	;
	if v1174 < v1153 {
		v2323 = v1170
		goto L11
	} else {
		goto L215
	}
L215:
	;
	v1180 = v1153 + base.B2i32(v1173 != int32(0)) + int32(1)
	v1182 = v1147 & int32(-33)
	if v1182 == int32(70) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1364 = v1180 + v1342
	if v54^int32(2147483647) < v1364 {
		v2323 = v1170
		goto L11
	} else {
		goto L248
	}
L217:
	;
	if v1180^int32(2147483647) < v959 {
		v2323 = v1170
		goto L11
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1193 = v959 >> (uint(int32(31)) % 32)
	v1196 = base.I64_extend_i32_u(v959 ^ v1193 - v1193)
	if base.Ui64(v1196) < base.Ui64(int64(4294967296)) {
		goto L226
	} else {
		goto L227
	}
L220:
	;
	v1188 = int32(0)
	if v1188 < v959 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1191 = v959
	goto L223
L222:
	;
	v1191 = v1188
	goto L223
L223:
	;
	v1342 = v1191
	v1350 = v965
	goto L216
L224:
	;
	if v94-v1250 <= int32(1) {
		goto L238
	} else {
		goto L239
	}
L225:
	;
	if v1227 != int64(0) {
		goto L232
	} else {
		goto L233
	}
L226:
	;
	v1223 = v94
	v1227 = v1196
	goto L225
L227:
	;
	goto L228
L228:
	;
	v1203 = v1196
	v1204 = v94
	goto L229
L229:
	;
	v1210 = v1204 - int32(1)
	v1211 = int64(10)
	v1212 = base.I64_div_u_s(v1203, v1211)
	v1218 = base.I32_wrap_i64(v1203-v1212*v1211) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1210))) = uint8(v1218)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1203) {
		v1203 = v1212
		v1204 = v1210
		goto L229
	} else {
		goto L231
	}
L230:
	;
	v1223 = v1210
	v1227 = v1212
	goto L225
L231:
	;
	goto L230
L232:
	;
	v1232 = v1223
	v1233 = base.I32_wrap_i64(v1227)
	goto L235
L233:
	;
	v1250 = v1223
	goto L234
L234:
	;
	goto L224
L235:
	;
	v1238 = v1232 - int32(1)
	v1239 = int32(10)
	v1240 = base.I32_div_u_s(v1233, v1239)
	v1245 = v1233 - v1240*v1239 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1238))) = uint8(v1245)
	if base.Ui32(int32(9)) < base.Ui32(v1233) {
		v1232 = v1238
		v1233 = v1240
		goto L235
	} else {
		goto L237
	}
L236:
	;
	v1250 = v1238
	goto L234
L237:
	;
	goto L236
L238:
	;
	v1264 = v1250
	goto L241
L239:
	;
	v1299 = v1250
	goto L240
L240:
	;
	v1322 = v1299 - int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v1322))) = uint8(v1147)
	if v959 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v1287 = v1264 - int32(1)
	v1288 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1287))) = uint8(v1288)
	if v94-v1287 < int32(2) {
		v1264 = v1287
		goto L241
	} else {
		goto L243
	}
L242:
	;
	v1299 = v1287
	goto L240
L243:
	;
	goto L242
L244:
	;
	v1330 = int32(45)
	goto L246
L245:
	;
	v1330 = int32(43)
	goto L246
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1299-int32(1)))) = uint8(v1330)
	v1332 = v94 - v1322
	if v1180^int32(2147483647) < v1332 {
		v2323 = v1170
		goto L11
	} else {
		goto L247
	}
L247:
	;
	v1342 = v1332
	v1350 = v1322
	goto L216
L248:
	;
	v1369 = v1364 + v54
	F_pad(m, l0, int32(32), l2, v1369, l4)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L15
	} else {
		goto L249
	}
L249:
	;
	F_out(m, l0, v55, v54)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L15
	} else {
		goto L250
	}
L250:
	;
	F_pad(m, l0, int32(48), l2, v1369, l4^int32(65536))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L15
	} else {
		goto L251
	}
L251:
	;
	if v1182 == int32(70) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	F_pad(m, l0, int32(32), l2, v1369, l4^int32(8192))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L15
	} else {
		goto L366
	}
L253:
	;
	v1985 = int32(9)
	F_pad(m, l0, int32(48), v1962+v1985, v1985, int32(0))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L15
	} else {
		goto L365
	}
L254:
	;
	v1962 = v1153
	goto L253
L255:
	;
	v1384 = v31 + int32(16) | int32(9)
	if base.Ui32(v173) < base.Ui32(v960) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	goto L257
L257:
	;
	if v1153 < int32(0) {
		v1929 = v1153
		goto L322
	} else {
		goto L323
	}
L258:
	;
	v1386 = v173
	goto L260
L259:
	;
	v1386 = v960
	goto L260
L260:
	;
	v1396 = v1386
	goto L261
L261:
	;
	v1415 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1396))))
	if base.Ui64(v1415) < base.Ui64(int64(4294967296)) {
		goto L265
	} else {
		goto L266
	}
L262:
	;
	if v1173 != 0 {
		goto L288
	} else {
		goto L289
	}
L263:
	;
	if v1386 != v1396 {
		goto L278
	} else {
		goto L279
	}
L264:
	;
	if v1446 != int64(0) {
		goto L271
	} else {
		goto L272
	}
L265:
	;
	v1442 = v1384
	v1446 = v1415
	goto L264
L266:
	;
	goto L267
L267:
	;
	v1422 = v1415
	v1423 = v1384
	goto L268
L268:
	;
	v1429 = v1423 - int32(1)
	v1430 = int64(10)
	v1431 = base.I64_div_u_s(v1422, v1430)
	v1437 = base.I32_wrap_i64(v1422-v1431*v1430) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1429))) = uint8(v1437)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1422) {
		v1422 = v1431
		v1423 = v1429
		goto L268
	} else {
		goto L270
	}
L269:
	;
	v1442 = v1429
	v1446 = v1431
	goto L264
L270:
	;
	goto L269
L271:
	;
	v1451 = v1442
	v1452 = base.I32_wrap_i64(v1446)
	goto L274
L272:
	;
	v1469 = v1442
	goto L273
L273:
	;
	goto L263
L274:
	;
	v1457 = v1451 - int32(1)
	v1458 = int32(10)
	v1459 = base.I32_div_u_s(v1452, v1458)
	v1464 = v1452 - v1459*v1458 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1457))) = uint8(v1464)
	if base.Ui32(int32(9)) < base.Ui32(v1452) {
		v1451 = v1457
		v1452 = v1459
		goto L274
	} else {
		goto L276
	}
L275:
	;
	v1469 = v1457
	goto L273
L276:
	;
	goto L275
L277:
	;
	F_out(m, l0, v1524, v1384-v1524)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L15
	} else {
		goto L286
	}
L278:
	;
	if base.Ui32(v1469) <= base.Ui32(v31+int32(16)) {
		v1524 = v1469
		goto L277
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	if v1469 != v1384 {
		v1524 = v1469
		goto L277
	} else {
		goto L285
	}
L281:
	;
	v1484 = v1469
	goto L282
L282:
	;
	v1507 = v1484 - int32(1)
	v1508 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1507))) = uint8(v1508)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1507) {
		v1484 = v1507
		goto L282
	} else {
		goto L284
	}
L283:
	;
	v1524 = v1507
	goto L277
L284:
	;
	goto L283
L285:
	;
	v1515 = v1469 - int32(1)
	v1516 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1515))) = uint8(v1516)
	v1524 = v1515
	goto L277
L286:
	;
	v1550 = v1396 + int32(4)
	if base.Ui32(v1550) <= base.Ui32(v173) {
		v1396 = v1550
		goto L261
	} else {
		goto L287
	}
L287:
	;
	goto L262
L288:
	;
	F_out(m, l0, int32(669525), int32(1))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L15
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	if base.Ui32(v985) <= base.Ui32(v1550) {
		goto L254
	} else {
		goto L292
	}
L291:
	;
	goto L290
L292:
	;
	if v1153 <= int32(0) {
		goto L254
	} else {
		goto L293
	}
L293:
	;
	v1568 = v1550
	v1570 = v1153
	goto L294
L294:
	;
	v1587 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1568))))
	if base.Ui64(v1587) < base.Ui64(int64(4294967296)) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v1962 = v1719
	goto L253
L296:
	;
	if base.Ui32(v31+int32(16)) < base.Ui32(v1641) {
		goto L310
	} else {
		goto L311
	}
L297:
	;
	if v1618 != int64(0) {
		goto L304
	} else {
		goto L305
	}
L298:
	;
	v1614 = v1384
	v1618 = v1587
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1594 = v1587
	v1595 = v1384
	goto L301
L301:
	;
	v1601 = v1595 - int32(1)
	v1602 = int64(10)
	v1603 = base.I64_div_u_s(v1594, v1602)
	v1609 = base.I32_wrap_i64(v1594-v1603*v1602) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1601))) = uint8(v1609)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1594) {
		v1594 = v1603
		v1595 = v1601
		goto L301
	} else {
		goto L303
	}
L302:
	;
	v1614 = v1601
	v1618 = v1603
	goto L297
L303:
	;
	goto L302
L304:
	;
	v1623 = v1614
	v1624 = base.I32_wrap_i64(v1618)
	goto L307
L305:
	;
	v1641 = v1614
	goto L306
L306:
	;
	goto L296
L307:
	;
	v1629 = v1623 - int32(1)
	v1630 = int32(10)
	v1631 = base.I32_div_u_s(v1624, v1630)
	v1636 = v1624 - v1631*v1630 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1629))) = uint8(v1636)
	if base.Ui32(int32(9)) < base.Ui32(v1624) {
		v1623 = v1629
		v1624 = v1631
		goto L307
	} else {
		goto L309
	}
L308:
	;
	v1641 = v1629
	goto L306
L309:
	;
	goto L308
L310:
	;
	v1655 = v1641
	goto L313
L311:
	;
	v1690 = v1641
	goto L312
L312:
	;
	v1712 = int32(9)
	if v1712 <= v1570 {
		goto L316
	} else {
		goto L317
	}
L313:
	;
	v1678 = v1655 - int32(1)
	v1679 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1678))) = uint8(v1679)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1678) {
		v1655 = v1678
		goto L313
	} else {
		goto L315
	}
L314:
	;
	v1690 = v1678
	goto L312
L315:
	;
	goto L314
L316:
	;
	v1715 = v1712
	goto L318
L317:
	;
	v1715 = v1570
	goto L318
L318:
	;
	F_out(m, l0, v1690, v1715)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L15
	} else {
		goto L319
	}
L319:
	;
	v1719 = v1570 - int32(9)
	v1721 = v1568 + int32(4)
	if base.Ui32(v985) <= base.Ui32(v1721) {
		v1962 = v1719
		goto L253
	} else {
		goto L320
	}
L320:
	;
	if int32(9) < v1570 {
		v1568 = v1721
		v1570 = v1719
		goto L294
	} else {
		goto L321
	}
L321:
	;
	goto L295
L322:
	;
	v1947 = int32(18)
	F_pad(m, l0, int32(48), v1929+v1947, v1947, int32(0))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L15
	} else {
		goto L363
	}
L323:
	;
	if base.Ui32(v960) < base.Ui32(v985) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1730 = v985
	goto L326
L325:
	;
	v1730 = v960 + int32(4)
	goto L326
L326:
	;
	v1734 = v31 + int32(16) | int32(9)
	v1742 = v960
	v1746 = v1153
	goto L327
L327:
	;
	v1763 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1742))))
	if base.Ui64(v1763) < base.Ui64(int64(4294967296)) {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	v1929 = v1912
	goto L322
L329:
	;
	if v1734 == v1817 {
		goto L343
	} else {
		goto L344
	}
L330:
	;
	if v1794 != int64(0) {
		goto L337
	} else {
		goto L338
	}
L331:
	;
	v1790 = v1734
	v1794 = v1763
	goto L330
L332:
	;
	goto L333
L333:
	;
	v1770 = v1763
	v1771 = v1734
	goto L334
L334:
	;
	v1777 = v1771 - int32(1)
	v1778 = int64(10)
	v1779 = base.I64_div_u_s(v1770, v1778)
	v1785 = base.I32_wrap_i64(v1770-v1779*v1778) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1777))) = uint8(v1785)
	if base.Ui64(int64(42949672959)) < base.Ui64(v1770) {
		v1770 = v1779
		v1771 = v1777
		goto L334
	} else {
		goto L336
	}
L335:
	;
	v1790 = v1777
	v1794 = v1779
	goto L330
L336:
	;
	goto L335
L337:
	;
	v1799 = v1790
	v1800 = base.I32_wrap_i64(v1794)
	goto L340
L338:
	;
	v1817 = v1790
	goto L339
L339:
	;
	goto L329
L340:
	;
	v1805 = v1799 - int32(1)
	v1806 = int32(10)
	v1807 = base.I32_div_u_s(v1800, v1806)
	v1812 = v1800 - v1807*v1806 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1805))) = uint8(v1812)
	if base.Ui32(int32(9)) < base.Ui32(v1800) {
		v1799 = v1805
		v1800 = v1807
		goto L340
	} else {
		goto L342
	}
L341:
	;
	v1817 = v1805
	goto L339
L342:
	;
	goto L341
L343:
	;
	v1824 = v1817 - int32(1)
	v1825 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1824))) = uint8(v1825)
	v1827 = v1824
	goto L345
L344:
	;
	v1827 = v1817
	goto L345
L345:
	;
	if v1742 != v960 {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1907 = v1734 - v1885
	if v1907 < v1746 {
		goto L357
	} else {
		goto L358
	}
L347:
	;
	if base.Ui32(v1827) <= base.Ui32(v31+int32(16)) {
		v1885 = v1827
		goto L346
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	F_out(m, l0, v1827, int32(1))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L15
	} else {
		goto L354
	}
L350:
	;
	v1838 = v1827
	goto L351
L351:
	;
	v1861 = v1838 - int32(1)
	v1862 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1861))) = uint8(v1862)
	if base.Ui32(v31+int32(16)) < base.Ui32(v1861) {
		v1838 = v1861
		goto L351
	} else {
		goto L353
	}
L352:
	;
	v1885 = v1861
	goto L346
L353:
	;
	goto L352
L354:
	;
	v1871 = v1827 + int32(1)
	if v1746|v1155 == int32(0) {
		v1885 = v1871
		goto L346
	} else {
		goto L355
	}
L355:
	;
	F_out(m, l0, int32(669525), int32(1))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L15
	} else {
		goto L356
	}
L356:
	;
	v1885 = v1871
	goto L346
L357:
	;
	v1909 = v1907
	goto L359
L358:
	;
	v1909 = v1746
	goto L359
L359:
	;
	F_out(m, l0, v1885, v1909)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L15
	} else {
		goto L360
	}
L360:
	;
	v1912 = v1746 - v1907
	v1914 = v1742 + int32(4)
	if base.Ui32(v1730) <= base.Ui32(v1914) {
		v1929 = v1912
		goto L322
	} else {
		goto L361
	}
L361:
	;
	if int32(0) <= v1912 {
		v1742 = v1914
		v1746 = v1912
		goto L327
	} else {
		goto L362
	}
L362:
	;
	goto L328
L363:
	;
	F_out(m, l0, v1350, v94-v1350)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L15
	} else {
		goto L364
	}
L364:
	;
	goto L252
L365:
	;
	goto L252
L366:
	;
	if v1369 < l2 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v2025 = l2
	goto L369
L368:
	;
	v2025 = v1369
	goto L369
L369:
	;
	v2323 = v2025
	goto L11
L370:
	;
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v2111 = v2109 >> (uint(int32(31)) % 32)
	v2114 = base.I64_extend_i32_u(v2109 ^ v2111 - v2111)
	if base.Ui64(v2114) < base.Ui64(int64(4294967296)) {
		goto L380
	} else {
		goto L381
	}
L371:
	;
	v2046 = int32(12) - l3
	v2064 = float64(16)
	goto L372
L372:
	;
	v2069 = base.F64_mul(v2064, float64(16))
	v2071 = v2046 - int32(1)
	if v2071 != 0 {
		v2046 = v2071
		v2064 = v2069
		goto L372
	} else {
		goto L374
	}
L373:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034))))
	if v2072 == int32(45) {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	goto L373
L375:
	;
	v2082 = base.F64_neg(base.F64_add(v2069, base.F64_sub(base.F64_neg(v132), v2069)))
	goto L370
L376:
	;
	goto L377
L377:
	;
	v2082 = base.F64_sub(base.F64_add(v132, v2069), v2069)
	goto L370
L378:
	;
	if v94 == v2168 {
		goto L392
	} else {
		goto L393
	}
L379:
	;
	if v2145 != int64(0) {
		goto L386
	} else {
		goto L387
	}
L380:
	;
	v2141 = v94
	v2145 = v2114
	goto L379
L381:
	;
	goto L382
L382:
	;
	v2121 = v2114
	v2122 = v94
	goto L383
L383:
	;
	v2128 = v2122 - int32(1)
	v2129 = int64(10)
	v2130 = base.I64_div_u_s(v2121, v2129)
	v2136 = base.I32_wrap_i64(v2121-v2130*v2129) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2128))) = uint8(v2136)
	if base.Ui64(int64(42949672959)) < base.Ui64(v2121) {
		v2121 = v2130
		v2122 = v2128
		goto L383
	} else {
		goto L385
	}
L384:
	;
	v2141 = v2128
	v2145 = v2130
	goto L379
L385:
	;
	goto L384
L386:
	;
	v2150 = v2141
	v2151 = base.I32_wrap_i64(v2145)
	goto L389
L387:
	;
	v2168 = v2141
	goto L388
L388:
	;
	goto L378
L389:
	;
	v2156 = v2150 - int32(1)
	v2157 = int32(10)
	v2158 = base.I32_div_u_s(v2151, v2157)
	v2163 = v2151 - v2158*v2157 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2156))) = uint8(v2163)
	if base.Ui32(int32(9)) < base.Ui32(v2151) {
		v2150 = v2156
		v2151 = v2158
		goto L389
	} else {
		goto L391
	}
L390:
	;
	v2168 = v2156
	goto L388
L391:
	;
	goto L390
L392:
	;
	v2175 = v2168 - int32(1)
	v2176 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2175))) = uint8(v2176)
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v2179 = v2175
	v2180 = v2178
	goto L394
L393:
	;
	v2179 = v2168
	v2180 = v2109
	goto L394
L394:
	;
	v2181 = int32(2)
	v2182 = v54 | v2181
	v2186 = v2179 - v2181
	v2188 = l5 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v2186))) = uint8(v2188)
	if v2180 < int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2196 = int32(45)
	goto L397
L396:
	;
	v2196 = int32(43)
	goto L397
L397:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2179-int32(1)))) = uint8(v2196)
	v2200 = int32(0)
	v2208 = v2082
	v2214 = v31 + int32(16)
	goto L398
L398:
	;
	if base.F64_lt(base.F64_abs(v2208), float64(2.147483648e+09)) != 0 {
		goto L401
	} else {
		goto L402
	}
L399:
	;
	v2269 = v94 - v2186
	v2270 = v2182 + v2269
	if int32(2147483645)-v2270 < l3 {
		v2323 = int32(-1)
		goto L11
	} else {
		goto L408
	}
L400:
	;
	v2243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2240)+uint32(_consts[1626]))))
	v2244 = v2243 | l5&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2214))) = uint8(v2244)
	v2249 = base.F64_mul(base.F64_sub(v2208, base.F64_convert_i32_s(v2240)), float64(16))
	v2250 = int32(1)
	v2251 = v2214 + v2250
	if v2251-(v31+int32(16)) != v2250 {
		v2264 = v2251
		goto L404
	} else {
		goto L405
	}
L401:
	;
	v2238 = base.I32_trunc_f64_s(v2208)
	v2240 = v2238
	goto L400
L402:
	;
	goto L403
L403:
	;
	v2240 = int32(-2147483648)
	goto L400
L404:
	;
	if base.F64_ne(v2249, float64(0)) != 0 {
		v2208 = v2249
		v2214 = v2264
		goto L398
	} else {
		goto L407
	}
L405:
	;
	if base.F64_eq(v2249, float64(0))&(base.B2i32(l4&int32(8) == v2200)&base.B2i32(l3 <= v2200)) != 0 {
		v2264 = v2251
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v2260 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2214)+1)) = uint8(v2260)
	v2264 = v2214 + int32(2)
	goto L404
L407:
	;
	goto L399
L408:
	;
	v2274 = int32(2)
	v2278 = v2264 - (v31 + int32(16))
	if v2278-v2274 < l3 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v2282 = l3 + v2274
	goto L411
L410:
	;
	v2282 = v2278
	goto L411
L411:
	;
	if l3 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2283 = v2282
	goto L414
L413:
	;
	v2283 = v2278
	goto L414
L414:
	;
	v2284 = v2270 + v2283
	F_pad(m, l0, int32(32), l2, v2284, l4)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L15
	} else {
		goto L415
	}
L415:
	;
	F_out(m, l0, v2034, v2182)
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L15
	} else {
		goto L416
	}
L416:
	;
	F_pad(m, l0, int32(48), l2, v2284, l4^int32(65536))
	mBase = m.M
	v2293 = m.ExcPending
	if v2293 != 0 {
		goto L15
	} else {
		goto L417
	}
L417:
	;
	F_out(m, l0, v31+int32(16), v2278)
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L15
	} else {
		goto L418
	}
L418:
	;
	v2300 = int32(0)
	F_pad(m, l0, int32(48), v2283-v2278, v2300, v2300)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L15
	} else {
		goto L419
	}
L419:
	;
	F_out(m, l0, v2186, v2269)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L15
	} else {
		goto L420
	}
L420:
	;
	F_pad(m, l0, int32(32), l2, v2284, l4^int32(8192))
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L15
	} else {
		goto L421
	}
L421:
	;
	if v2284 < l2 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v2312 = l2
	goto L424
L423:
	;
	v2312 = v2284
	goto L424
L424:
	;
	v2323 = v2312
	goto L11
}
func F_fopen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v13 = F___strchrnul(m, int32(503627), v12)
	mBase = m.M
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v15 == v12&int32(255) {
		v19 = v13
	} else {
		v19 = v3
	}
	if v19 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
		v199 = int32(0)
	} else {
		v28 = F_strchr(m, l1, int32(43))
		mBase = m.M
		if v28 == int32(0) {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v34 = base.B2i32(v31 != int32(114))
		} else {
			v34 = int32(2)
		}
		v38 = F_strchr(m, l1, int32(120))
		mBase = m.M
		if v38 != 0 {
			v39 = v34 | int32(128)
		} else {
			v39 = v34
		}
		v43 = F_strchr(m, l1, int32(101))
		mBase = m.M
		if v43 != 0 {
			v44 = v39 | int32(524288)
		} else {
			v44 = v39
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v47 == int32(114) {
			v50 = v44
		} else {
			v50 = v44 | int32(64)
		}
		if v47 == int32(119) {
			v55 = v50 | int32(512)
		} else {
			v55 = v50
		}
		if v47 == int32(97) {
			v60 = v55 | int32(1024)
		} else {
			v60 = v55
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(438)
		v66 = m.Env.X__syscall_openat(m, int32(-100), l0, v60|int32(32768), v9)
		mBase = m.M
		if base.Ui32(int32(-4095)) <= base.Ui32(v66) {
			*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0) - v66
			v74 = int32(-1)
		} else {
			v74 = v66
		}
		if v74 < int32(0) {
			v199 = v3
		} else {
			v77 = m.G0
			v79 = v77 - int32(32)
			m.G0 = v79
			v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
			v83 = F___strchrnul(m, int32(503627), v82)
			mBase = m.M
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
			if v85 == v82&int32(255) {
				v89 = v83
			} else {
				v89 = int32(0)
			}
			if v89 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
				v185 = int32(0)
			} else {
				v96 = F_emscripten_builtin_malloc(m, int32(1176))
				mBase = m.M
				if v96 != 0 {
					v102 = F__emscripten_memset_bulkmem(m, v96, base.I32_extend8_s(int32(0)), int32(144))
					mBase = m.M
					v103 = int32(43)
					v104 = F___strchrnul(m, l1, v103)
					mBase = m.M
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
					if v106 == v103 {
						v110 = v104
					} else {
						v110 = int32(0)
					}
					if v110 == int32(0) {
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
						if v115 == int32(114) {
							v118 = int32(8)
						} else {
							v118 = int32(4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v118
					} else {
					}
					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
					if v120 != int32(97) {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						v143 = v123
					} else {
						v125 = int32(0)
						v126 = m.Env.X__syscall_fcntl64(m, v74, int32(3), v125)
						mBase = m.M
						if v126&int32(1024) == v125 {
							*(*int64)(unsafe.Add(mBase, uint32(v79)+16)) = base.I64_extend_i32_s(v126 | int32(1024))
							v138 = m.Env.X__syscall_fcntl64(m, v74, int32(4), v79+int32(16))
							mBase = m.M
						} else {
						}
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
						v141 = v139 | int32(128)
						*(*int32)(unsafe.Add(mBase, uint32(v96))) = v141
						v143 = v141
					}
					*(*int32)(unsafe.Add(mBase, uint32(v96)+80)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v96)+48)) = int32(1024)
					*(*int32)(unsafe.Add(mBase, uint32(v96)+60)) = v74
					*(*int32)(unsafe.Add(mBase, uint32(v96)+44)) = v96 + int32(152)
					if v143&int32(8) != 0 {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v79))) = base.I64_extend_i32_u(v79 + int32(24))
						v159 = m.Env.X__syscall_ioctl(m, v74, int32(21523), v79)
						mBase = m.M
						if v159 != 0 {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v96)+80)) = int32(10)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v96)+40)) = int32(7104)
					*(*int32)(unsafe.Add(mBase, uint32(v96)+36)) = int32(7105)
					*(*int32)(unsafe.Add(mBase, uint32(v96)+32)) = int32(7106)
					*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = int32(7107)
					v171 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1579])))
					if v171 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v96)+76)) = int32(-1)
					} else {
					}
					v177 = *(*int32)(unsafe.Add(mBase, _consts[1580]))
					*(*int32)(unsafe.Add(mBase, uint32(v96)+56)) = v177
					if v177 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v177)+52)) = v96
					} else {
					}
					*(*int32)(unsafe.Add(mBase, _consts[1580])) = v96
					v185 = v96
				} else {
					v185 = int32(0)
				}
			}
			m.G0 = v79 + int32(32)
			if v185 != 0 {
				v199 = v185
			} else {
				v189 = m.Wasi_snapshot_preview1.Fd_close(m, v74)
				mBase = m.M
				v199 = int32(0)
			}
		}
	}
	m.G0 = v9 + int32(16)
	return v199
}
func F_forget_invalid_pages(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L27
	}
L2:
	;
	m.G0 = v11 + int32(112)
	return
L3:
	;
	F_hash_seq_init(m, v11+int32(92), v14)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v23 = F_hash_seq_search(m, v11+int32(92))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = v23
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v38 != v29 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v85 = F_hash_seq_search(m, v11+int32(92))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v40 != v28 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v42 != v27 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v44 != l1 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_GetRelationPath(m, v11+int32(20), v55, v56, v57, int32(-1), l1)
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v78 = F_hash_search(m, v75, v33, int32(2), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v11 + int32(20)
	F_errmsg_internal(m, int32(451469), v11)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(493655), int32(184), int32(170385))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	if v78 == int32(0) {
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
		v33 = v85
		goto L8
	} else {
		goto L26
	}
L26:
	;
	goto L9
L27:
	;
	F_errmsg_internal(m, int32(444755), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(493655), int32(189), int32(170385))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
func F_fputc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	v1 = l0
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if int32(0) <= v6 {
		if v6 == int32(0) {
			v31 = l1 + int32(76)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			if v32 != 0 {
				v34 = v32
			} else {
				v34 = int32(1073741823)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
			v37 = v1 & int32(255)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
			if v37 == v38 {
				F___overflow(m, l1, v37)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
					return
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v40 == v41 {
					F___overflow(m, l1, v37)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v1)
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
					return
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[1581]))
			if v12 != v6&int32(1073741823) {
				v31 = l1 + int32(76)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				if v32 != 0 {
					v34 = v32
				} else {
					v34 = int32(1073741823)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
				v37 = v1 & int32(255)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				if v37 == v38 {
					F___overflow(m, l1, v37)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					if v40 == v41 {
						F___overflow(m, l1, v37)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v1)
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				}
			} else {
				v17 = v1 & int32(255)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				if v17 == v18 {
					F___overflow(m, l1, v17)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					if v20 == v21 {
						F___overflow(m, l1, v17)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v20 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
						return
					}
				}
			}
		}
	} else {
		v17 = v1 & int32(255)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
		if v17 == v18 {
			F___overflow(m, l1, v17)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				return
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v20 == v21 {
				F___overflow(m, l1, v17)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v20 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
				return
			}
		}
	}
}
func F_fread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	v9 = l1 * l2
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+72)) = v10 - int32(1) | v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v15 == v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = l0
	v30 = v9
	goto L3
L2:
	;
	v18 = v16 - v15
	if base.Ui32(v18) < base.Ui32(v9) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v20 = v18
	goto L6
L5:
	;
	v20 = v9
	goto L6
L6:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v23 + v20
	v28 = l0 + v20
	v30 = v9 - v20
	goto L3
L8:
	;
	v21 = F__emscripten_memcpy_bulkmem(m, l0, v15, v20)
	mBase = m.M
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	v31 = v28
	v36 = v30
	goto L14
L12:
	;
	goto L13
L13:
	;
	if l1 != 0 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v38 = F___toread(m, l3)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	v52 = v36 - v45
	if v52 != 0 {
		v31 = v31 + v45
		v36 = v52
		goto L14
	} else {
		goto L24
	}
L17:
	;
	return int32(0)
L18:
	;
	if v38 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v45 = m.T0[v44].(func(*base.Module, int32, int32, int32) int32)(m, l3, v31, v36)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v49 = base.I32_div_u_s(v9-v36, l1)
	return v49
L22:
	;
	if v45 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L15
L25:
	;
	v61 = l2
	goto L27
L26:
	;
	v61 = int32(0)
	goto L27
L27:
	;
	return v61
}
func F_freeGISTstate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v15 <= int32(0) {
		v68 = v14
	} else {
		v19 = v15 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v15) {
			v26 = int32(0)
			v28 = v14
			for {
				v35 = v28 * int32(-1799145247)
				v37 = v26 + int32(8)
				if v37 != v15&int32(2147483640) {
					v26 = v37
					v28 = v35
					continue
				} else {
					break
				}
				break
			}
			v42 = v35
		} else {
			v42 = v14
		}
		if v19 == int32(0) {
			v68 = v42
		} else {
			v52 = int32(0)
			v54 = v42
			for {
				v61 = v54 * int32(4069)
				v63 = v52 + int32(1)
				if v63 != v19 {
					v52 = v63
					v54 = v61
					continue
				} else {
					break
				}
				break
			}
			v68 = v61
		}
	}
	v75 = base.I32_div_u_s(v68, int32(4069))
	v78 = base.I32_div_u_s(v68, int32(16556761))
	v81 = v68 + v75 + v78 + int32(3)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 == int32(0) {
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v88
		v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v90
		v94 = F_smgropen(m, v12+int32(24), v87)
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v94
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
			if v100 != 0 {
				v108 = v100
			} else {
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)+76))
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v102
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v94)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v102))) = v104
				v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+72))
				v108 = v106
			}
			*(*int32)(unsafe.Add(mBase, uint32(v94)+72)) = v108 + int32(1)
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v113 = v112
			v115 = v113 + int32(24)
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
			v119 = v15 ^ int32(-1) + v81
			if base.B2i32(v116 != int32(-1))&base.B2i32(base.Ui32(v119) < base.Ui32(v116)) == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(-1)
				v127 = F_smgrexists(m, v113, int32(1))
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int32(0)
				} else {
					if v127 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = int32(0)
						v145 = int32(0)
						if l2 == v145 {
							v253 = v145
							m.G0 = v12 + int32(48)
							return v253
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
							v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
							*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
							v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								v165 = v161
								if v165 < int32(0) {
									v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
									v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
									v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
									if v176 != 0 {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v165, int32(2))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int32(0)
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
											v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
											if v184 == int32(0) {
												v206 = v183
												if v206&int32(3) != 0 {
												} else {
												}
												v233 = F___memset(m, v206, int32(0), int32(8192))
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
												v239 = int32(8196)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
												v245 = int32(8192)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
											} else {
											}
											F_LockBuffer(m, v165, int32(0))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int32(0)
											} else {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
											}
										}
									}
								} else {
									v188 = v165 << (uint(int32(13)) % 32)
									v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
									if v194 != 0 {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v165, int32(2))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											v200 = v199 + v188
											v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
											if v203 != 0 {
											} else {
												v206 = v200 + int32(-8192)
												if v206&int32(3) != 0 {
												} else {
												}
												v233 = F___memset(m, v206, int32(0), int32(8192))
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
												v239 = int32(8196)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
												v245 = int32(8192)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
											}
											F_LockBuffer(m, v165, int32(0))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int32(0)
											} else {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
											}
										}
									}
								}
							}
						}
					} else {
						v132 = F_smgrnblocks(m, v113, int32(1))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
							v135 = v134
							if base.Ui32(v135) <= base.Ui32(v119) {
								v145 = int32(0)
								if l2 == v145 {
									v253 = v145
									m.G0 = v12 + int32(48)
									return v253
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
									v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
									*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
									v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return int32(0)
									} else {
										v165 = v161
										if v165 < int32(0) {
											v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
											v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
											v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
											if v176 != 0 {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
											} else {
												F_LockBuffer(m, v165, int32(2))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int32(0)
												} else {
													v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
													v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
													v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
													if v184 == int32(0) {
														v206 = v183
														if v206&int32(3) != 0 {
														} else {
														}
														v233 = F___memset(m, v206, int32(0), int32(8192))
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
														v239 = int32(8196)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
														v245 = int32(8192)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
													} else {
													}
													F_LockBuffer(m, v165, int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														v253 = v165
														m.G0 = v12 + int32(48)
														return v253
													}
												}
											}
										} else {
											v188 = v165 << (uint(int32(13)) % 32)
											v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
											if v194 != 0 {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
											} else {
												F_LockBuffer(m, v165, int32(2))
												mBase = m.M
												v197 = m.ExcPending
												if v197 != 0 {
													return int32(0)
												} else {
													v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
													v200 = v199 + v188
													v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
													if v203 != 0 {
													} else {
														v206 = v200 + int32(-8192)
														if v206&int32(3) != 0 {
														} else {
														}
														v233 = F___memset(m, v206, int32(0), int32(8192))
														mBase = m.M
														*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
														v239 = int32(8196)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
														v245 = int32(8192)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
														*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
													}
													F_LockBuffer(m, v165, int32(0))
													mBase = m.M
													v251 = m.ExcPending
													if v251 != 0 {
														return int32(0)
													} else {
														v253 = v165
														m.G0 = v12 + int32(48)
														return v253
													}
												}
											}
										}
									}
								}
							} else {
								v140 = F_ReadBufferExtended(m, l0, int32(1), v119, int32(3), int32(0))
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v165 = v140
									if v165 < int32(0) {
										v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
										v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
										if v176 != 0 {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v165, int32(2))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return int32(0)
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
												v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
												v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
												if v184 == int32(0) {
													v206 = v183
													if v206&int32(3) != 0 {
													} else {
													}
													v233 = F___memset(m, v206, int32(0), int32(8192))
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
													v239 = int32(8196)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
													v245 = int32(8192)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
												} else {
												}
												F_LockBuffer(m, v165, int32(0))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													v253 = v165
													m.G0 = v12 + int32(48)
													return v253
												}
											}
										}
									} else {
										v188 = v165 << (uint(int32(13)) % 32)
										v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
										if v194 != 0 {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v165, int32(2))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
												v200 = v199 + v188
												v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
												if v203 != 0 {
												} else {
													v206 = v200 + int32(-8192)
													if v206&int32(3) != 0 {
													} else {
													}
													v233 = F___memset(m, v206, int32(0), int32(8192))
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
													v239 = int32(8196)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
													v245 = int32(8192)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
												}
												F_LockBuffer(m, v165, int32(0))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													v253 = v165
													m.G0 = v12 + int32(48)
													return v253
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
				v135 = v116
				if base.Ui32(v135) <= base.Ui32(v119) {
					v145 = int32(0)
					if l2 == v145 {
						v253 = v145
						m.G0 = v12 + int32(48)
						return v253
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
						v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
						v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							v165 = v161
							if v165 < int32(0) {
								v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
								v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
								v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
								if v176 != 0 {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v165, int32(2))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int32(0)
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
										v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
										if v184 == int32(0) {
											v206 = v183
											if v206&int32(3) != 0 {
											} else {
											}
											v233 = F___memset(m, v206, int32(0), int32(8192))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
											v239 = int32(8196)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
											v245 = int32(8192)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
										} else {
										}
										F_LockBuffer(m, v165, int32(0))
										mBase = m.M
										v251 = m.ExcPending
										if v251 != 0 {
											return int32(0)
										} else {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										}
									}
								}
							} else {
								v188 = v165 << (uint(int32(13)) % 32)
								v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
								if v194 != 0 {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v165, int32(2))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v200 = v199 + v188
										v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
										if v203 != 0 {
										} else {
											v206 = v200 + int32(-8192)
											if v206&int32(3) != 0 {
											} else {
											}
											v233 = F___memset(m, v206, int32(0), int32(8192))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
											v239 = int32(8196)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
											v245 = int32(8192)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
										}
										F_LockBuffer(m, v165, int32(0))
										mBase = m.M
										v251 = m.ExcPending
										if v251 != 0 {
											return int32(0)
										} else {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										}
									}
								}
							}
						}
					}
				} else {
					v140 = F_ReadBufferExtended(m, l0, int32(1), v119, int32(3), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						v165 = v140
						if v165 < int32(0) {
							v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
							v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
							v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
							if v176 != 0 {
								v253 = v165
								m.G0 = v12 + int32(48)
								return v253
							} else {
								F_LockBuffer(m, v165, int32(2))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
									v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
									if v184 == int32(0) {
										v206 = v183
										if v206&int32(3) != 0 {
										} else {
										}
										v233 = F___memset(m, v206, int32(0), int32(8192))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
										v239 = int32(8196)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
										v245 = int32(8192)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
									} else {
									}
									F_LockBuffer(m, v165, int32(0))
									mBase = m.M
									v251 = m.ExcPending
									if v251 != 0 {
										return int32(0)
									} else {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									}
								}
							}
						} else {
							v188 = v165 << (uint(int32(13)) % 32)
							v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
							if v194 != 0 {
								v253 = v165
								m.G0 = v12 + int32(48)
								return v253
							} else {
								F_LockBuffer(m, v165, int32(2))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v200 = v199 + v188
									v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
									if v203 != 0 {
									} else {
										v206 = v200 + int32(-8192)
										if v206&int32(3) != 0 {
										} else {
										}
										v233 = F___memset(m, v206, int32(0), int32(8192))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
										v239 = int32(8196)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
										v245 = int32(8192)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
									}
									F_LockBuffer(m, v165, int32(0))
									mBase = m.M
									v251 = m.ExcPending
									if v251 != 0 {
										return int32(0)
									} else {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v113 = v84
		v115 = v113 + int32(24)
		v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
		v119 = v15 ^ int32(-1) + v81
		if base.B2i32(v116 != int32(-1))&base.B2i32(base.Ui32(v119) < base.Ui32(v116)) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(-1)
			v127 = F_smgrexists(m, v113, int32(1))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return int32(0)
			} else {
				if v127 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v113)+24)) = int32(0)
					v145 = int32(0)
					if l2 == v145 {
						v253 = v145
						m.G0 = v12 + int32(48)
						return v253
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
						v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
						v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							v165 = v161
							if v165 < int32(0) {
								v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
								v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
								v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
								if v176 != 0 {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v165, int32(2))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int32(0)
									} else {
										v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
										v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
										if v184 == int32(0) {
											v206 = v183
											if v206&int32(3) != 0 {
											} else {
											}
											v233 = F___memset(m, v206, int32(0), int32(8192))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
											v239 = int32(8196)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
											v245 = int32(8192)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
										} else {
										}
										F_LockBuffer(m, v165, int32(0))
										mBase = m.M
										v251 = m.ExcPending
										if v251 != 0 {
											return int32(0)
										} else {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										}
									}
								}
							} else {
								v188 = v165 << (uint(int32(13)) % 32)
								v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
								if v194 != 0 {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
								} else {
									F_LockBuffer(m, v165, int32(2))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v200 = v199 + v188
										v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
										if v203 != 0 {
										} else {
											v206 = v200 + int32(-8192)
											if v206&int32(3) != 0 {
											} else {
											}
											v233 = F___memset(m, v206, int32(0), int32(8192))
											mBase = m.M
											*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
											v239 = int32(8196)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
											v245 = int32(8192)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
											*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
										}
										F_LockBuffer(m, v165, int32(0))
										mBase = m.M
										v251 = m.ExcPending
										if v251 != 0 {
											return int32(0)
										} else {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										}
									}
								}
							}
						}
					}
				} else {
					v132 = F_smgrnblocks(m, v113, int32(1))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
						v135 = v134
						if base.Ui32(v135) <= base.Ui32(v119) {
							v145 = int32(0)
							if l2 == v145 {
								v253 = v145
								m.G0 = v12 + int32(48)
								return v253
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
								v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
								*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
								v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									v165 = v161
									if v165 < int32(0) {
										v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
										v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
										v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
										if v176 != 0 {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v165, int32(2))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return int32(0)
											} else {
												v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
												v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
												v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
												if v184 == int32(0) {
													v206 = v183
													if v206&int32(3) != 0 {
													} else {
													}
													v233 = F___memset(m, v206, int32(0), int32(8192))
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
													v239 = int32(8196)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
													v245 = int32(8192)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
												} else {
												}
												F_LockBuffer(m, v165, int32(0))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													v253 = v165
													m.G0 = v12 + int32(48)
													return v253
												}
											}
										}
									} else {
										v188 = v165 << (uint(int32(13)) % 32)
										v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
										if v194 != 0 {
											v253 = v165
											m.G0 = v12 + int32(48)
											return v253
										} else {
											F_LockBuffer(m, v165, int32(2))
											mBase = m.M
											v197 = m.ExcPending
											if v197 != 0 {
												return int32(0)
											} else {
												v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
												v200 = v199 + v188
												v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
												if v203 != 0 {
												} else {
													v206 = v200 + int32(-8192)
													if v206&int32(3) != 0 {
													} else {
													}
													v233 = F___memset(m, v206, int32(0), int32(8192))
													mBase = m.M
													*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
													v239 = int32(8196)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
													v245 = int32(8192)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
													*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
												}
												F_LockBuffer(m, v165, int32(0))
												mBase = m.M
												v251 = m.ExcPending
												if v251 != 0 {
													return int32(0)
												} else {
													v253 = v165
													m.G0 = v12 + int32(48)
													return v253
												}
											}
										}
									}
								}
							}
						} else {
							v140 = F_ReadBufferExtended(m, l0, int32(1), v119, int32(3), int32(0))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								v165 = v140
								if v165 < int32(0) {
									v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
									v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
									v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
									if v176 != 0 {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v165, int32(2))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int32(0)
										} else {
											v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
											v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
											v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
											if v184 == int32(0) {
												v206 = v183
												if v206&int32(3) != 0 {
												} else {
												}
												v233 = F___memset(m, v206, int32(0), int32(8192))
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
												v239 = int32(8196)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
												v245 = int32(8192)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
											} else {
											}
											F_LockBuffer(m, v165, int32(0))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int32(0)
											} else {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
											}
										}
									}
								} else {
									v188 = v165 << (uint(int32(13)) % 32)
									v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
									if v194 != 0 {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									} else {
										F_LockBuffer(m, v165, int32(2))
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int32(0)
										} else {
											v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											v200 = v199 + v188
											v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
											if v203 != 0 {
											} else {
												v206 = v200 + int32(-8192)
												if v206&int32(3) != 0 {
												} else {
												}
												v233 = F___memset(m, v206, int32(0), int32(8192))
												mBase = m.M
												*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
												v239 = int32(8196)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
												v245 = int32(8192)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
												*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
											}
											F_LockBuffer(m, v165, int32(0))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int32(0)
											} else {
												v253 = v165
												m.G0 = v12 + int32(48)
												return v253
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
			v135 = v116
			if base.Ui32(v135) <= base.Ui32(v119) {
				v145 = int32(0)
				if l2 == v145 {
					v253 = v145
					m.G0 = v12 + int32(48)
					return v253
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v150
					*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l0
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v12)+36))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
					v161 = F_ExtendBufferedRelTo(m, v12+int32(8), int32(1), int32(20), v81-v15, int32(3))
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						v165 = v161
						if v165 < int32(0) {
							v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
							v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
							v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
							if v176 != 0 {
								v253 = v165
								m.G0 = v12 + int32(48)
								return v253
							} else {
								F_LockBuffer(m, v165, int32(2))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return int32(0)
								} else {
									v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
									v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
									if v184 == int32(0) {
										v206 = v183
										if v206&int32(3) != 0 {
										} else {
										}
										v233 = F___memset(m, v206, int32(0), int32(8192))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
										v239 = int32(8196)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
										v245 = int32(8192)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
									} else {
									}
									F_LockBuffer(m, v165, int32(0))
									mBase = m.M
									v251 = m.ExcPending
									if v251 != 0 {
										return int32(0)
									} else {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									}
								}
							}
						} else {
							v188 = v165 << (uint(int32(13)) % 32)
							v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
							if v194 != 0 {
								v253 = v165
								m.G0 = v12 + int32(48)
								return v253
							} else {
								F_LockBuffer(m, v165, int32(2))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v200 = v199 + v188
									v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
									if v203 != 0 {
									} else {
										v206 = v200 + int32(-8192)
										if v206&int32(3) != 0 {
										} else {
										}
										v233 = F___memset(m, v206, int32(0), int32(8192))
										mBase = m.M
										*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
										v239 = int32(8196)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
										v245 = int32(8192)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
										*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
									}
									F_LockBuffer(m, v165, int32(0))
									mBase = m.M
									v251 = m.ExcPending
									if v251 != 0 {
										return int32(0)
									} else {
										v253 = v165
										m.G0 = v12 + int32(48)
										return v253
									}
								}
							}
						}
					}
				}
			} else {
				v140 = F_ReadBufferExtended(m, l0, int32(1), v119, int32(3), int32(0))
				mBase = m.M
				v141 = m.ExcPending
				if v141 != 0 {
					return int32(0)
				} else {
					v165 = v140
					if v165 < int32(0) {
						v171 = (v165 ^ int32(-1)) << (uint(int32(2)) % 32)
						v173 = *(*int32)(unsafe.Add(mBase, _consts[1]))
						v175 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173)))
						v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)))
						if v176 != 0 {
							v253 = v165
							m.G0 = v12 + int32(48)
							return v253
						} else {
							F_LockBuffer(m, v165, int32(2))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								v181 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v181+v171)))
								v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
								if v184 == int32(0) {
									v206 = v183
									if v206&int32(3) != 0 {
									} else {
									}
									v233 = F___memset(m, v206, int32(0), int32(8192))
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
									v239 = int32(8196)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
									v245 = int32(8192)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
								} else {
								}
								F_LockBuffer(m, v165, int32(0))
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return int32(0)
								} else {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
								}
							}
						}
					} else {
						v188 = v165 << (uint(int32(13)) % 32)
						v190 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v190-int32(8178)))))
						if v194 != 0 {
							v253 = v165
							m.G0 = v12 + int32(48)
							return v253
						} else {
							F_LockBuffer(m, v165, int32(2))
							mBase = m.M
							v197 = m.ExcPending
							if v197 != 0 {
								return int32(0)
							} else {
								v199 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v200 = v199 + v188
								v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200-int32(8178)))))
								if v203 != 0 {
								} else {
									v206 = v200 + int32(-8192)
									if v206&int32(3) != 0 {
									} else {
									}
									v233 = F___memset(m, v206, int32(0), int32(8192))
									mBase = m.M
									*(*int32)(unsafe.Add(mBase, uint32(v206)+10)) = int32(1572864)
									v239 = int32(8196)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)) = uint16(v239)
									v245 = int32(8192)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+16)) = uint16(v245)
									*(*uint16)(unsafe.Add(mBase, uint32(v206)+14)) = uint16(v245)
								}
								F_LockBuffer(m, v165, int32(0))
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return int32(0)
								} else {
									v253 = v165
									m.G0 = v12 + int32(48)
									return v253
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
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0) - v3
		v11 = int32(-1)
	} else {
		v11 = v3
	}
	return v11
}
