package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimezone(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v14 - int32(43) {
	case 0, 2:
		*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
		v25 = F_strtol(m, l0+int32(1), v11+int32(12), int32(10))
		mBase = m.M
		v26 = int32(-5)
		v28 = *(*int32)(unsafe.Add(mBase, _consts[43]))
		if v28 == int32(68) {
			v104 = v26
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
			if v32 != 0 {
				if v32 != int32(58) {
					v75 = int32(0)
					v76 = v25
					v78 = v3
					if base.Ui32(int32(15)) < base.Ui32(v76) {
						v104 = v26
					} else {
						if base.Ui32(int32(59)) < base.Ui32(v75) {
							v104 = v26
						} else {
							if base.Ui32(int32(59)) < base.Ui32(v78) {
								v104 = v26
							} else {
								v85 = int32(60)
								v90 = (v76*v85+v75)*v85 + v78
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v93 == int32(45) {
									v96 = v90
								} else {
									v96 = int32(0) - v90
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
								if v101 != 0 {
									v102 = int32(-1)
								} else {
									v102 = int32(0)
								}
								v104 = v102
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
					v44 = F_strtol(m, v31+int32(1), v11+int32(12), int32(10))
					mBase = m.M
					v46 = *(*int32)(unsafe.Add(mBase, _consts[43]))
					if v46 == int32(68) {
						v104 = v26
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
						if v50 != int32(58) {
							v75 = v44
							v76 = v25
							v78 = v3
							if base.Ui32(int32(15)) < base.Ui32(v76) {
								v104 = v26
							} else {
								if base.Ui32(int32(59)) < base.Ui32(v75) {
									v104 = v26
								} else {
									if base.Ui32(int32(59)) < base.Ui32(v78) {
										v104 = v26
									} else {
										v85 = int32(60)
										v90 = (v76*v85+v75)*v85 + v78
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
										if v93 == int32(45) {
											v96 = v90
										} else {
											v96 = int32(0) - v90
										}
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
										if v101 != 0 {
											v102 = int32(-1)
										} else {
											v102 = int32(0)
										}
										v104 = v102
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
							v61 = F_strtol(m, v49+int32(1), v11+int32(12), int32(10))
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, _consts[43]))
							if v63 != int32(68) {
								v75 = v44
								v76 = v25
								v78 = v61
								if base.Ui32(int32(15)) < base.Ui32(v76) {
									v104 = v26
								} else {
									if base.Ui32(int32(59)) < base.Ui32(v75) {
										v104 = v26
									} else {
										if base.Ui32(int32(59)) < base.Ui32(v78) {
											v104 = v26
										} else {
											v85 = int32(60)
											v90 = (v76*v85+v75)*v85 + v78
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
											if v93 == int32(45) {
												v96 = v90
											} else {
												v96 = int32(0) - v90
											}
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
											v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
											if v101 != 0 {
												v102 = int32(-1)
											} else {
												v102 = int32(0)
											}
											v104 = v102
										}
									}
								}
							} else {
								v104 = v26
							}
						}
					}
				}
			} else {
				v66 = F_strlen(m, l0)
				mBase = m.M
				if base.Ui32(v66) < base.Ui32(int32(4)) {
					v75 = int32(0)
					v76 = v25
					v78 = v3
				} else {
					v70 = int32(100)
					v71 = base.I32_div_s(v25, v70)
					v75 = v25 - v71*v70
					v76 = v71
					v78 = v3
				}
				if base.Ui32(int32(15)) < base.Ui32(v76) {
					v104 = v26
				} else {
					if base.Ui32(int32(59)) < base.Ui32(v75) {
						v104 = v26
					} else {
						if base.Ui32(int32(59)) < base.Ui32(v78) {
							v104 = v26
						} else {
							v85 = int32(60)
							v90 = (v76*v85+v75)*v85 + v78
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
							if v93 == int32(45) {
								v96 = v90
							} else {
								v96 = int32(0) - v90
							}
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
							if v101 != 0 {
								v102 = int32(-1)
							} else {
								v102 = int32(0)
							}
							v104 = v102
						}
					}
				}
			}
		}
	default:
		v104 = int32(-1)
	}
	m.G0 = v11 + int32(16)
	return v104
}
func F_DecodeTimezoneAbbrev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
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
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v661 int32
	_ = v661
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v790 int32
	_ = v790
	v14 = m.G0
	v16 = v14 - int32(272)
	m.G0 = v16
	v19 = l0 * int32(20)
	v21 = v19 + int32(4532400)
	goto L5
L1:
	;
	m.G0 = v16 + int32(272)
	return v790
L2:
	;
	v790 = int32(0)
	goto L1
L3:
	;
	if v58-v59 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L6
L6:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = l1
	v30 = v21
	v31 = int32(10)
	v32 = v28
	goto L11
L8:
	;
	v54 = v21
	v58 = int32(0)
	goto L9
L9:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	goto L3
L10:
	;
	v54 = v49
	v58 = v51
	goto L9
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v32 != v34 {
		v49 = v30
		v51 = v32
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v49 = v43
	v51 = int32(0)
	goto L10
L13:
	;
	if v34 == int32(0) {
		v49 = v30
		v51 = v32
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v39 = v31 - int32(1)
	if v39 == int32(0) {
		v49 = v30
		v51 = v32
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v42 = int32(1)
	v43 = v30 + v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v44 != 0 {
		v29 = v29 + v42
		v30 = v43
		v31 = v39
		v32 = v44
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[960]))))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[961])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[962])))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v73
	goto L2
L18:
	;
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _consts[963]))
	if v508 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L21:
	;
	v80 = v16 + int32(16)
	goto L25
L22:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v196 != 0 {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	v193 = F_strlen(m, v182)
	mBase = m.M
	goto L22
L25:
	;
	goto L26
L26:
	;
	v87 = int32(255)
	if (v80^l1)&int32(3) != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v186)
	goto L23
L28:
	;
	v167 = v162
	v168 = v163
	v169 = v164
	goto L50
L29:
	;
	if v157 == int32(0) {
		v182 = v155
		v183 = v156
		goto L27
	} else {
		goto L49
	}
L30:
	;
	v155 = l1
	v156 = v80
	v157 = v87
	goto L29
L31:
	;
	goto L32
L32:
	;
	if l1&int32(3) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v124 == int32(0) {
		v182 = v121
		v183 = v122
		goto L27
	} else {
		goto L42
	}
L34:
	;
	v121 = l1
	v122 = v80
	v123 = v87
	v124 = int32(1)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v100 = l1
	v101 = v80
	v102 = v87
	goto L37
L37:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v104)
	if v104 == int32(0) {
		v162 = v100
		v163 = v101
		v164 = v102
		goto L28
	} else {
		goto L39
	}
L38:
	;
	v121 = v115
	v122 = v109
	v123 = v111
	v124 = v113
	goto L33
L39:
	;
	v108 = int32(1)
	v109 = v101 + v108
	v111 = v102 - v108
	v112 = int32(0)
	v113 = base.B2i32(v111 != v112)
	v115 = v100 + v108
	if v115&int32(3) == v112 {
		v121 = v115
		v122 = v109
		v123 = v111
		v124 = v113
		goto L33
	} else {
		goto L40
	}
L40:
	;
	if v111 != 0 {
		v100 = v115
		v101 = v109
		v102 = v111
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if v127 == int32(0) {
		v155 = v121
		v156 = v122
		v157 = v123
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v123) < base.Ui32(int32(4)) {
		v155 = v121
		v156 = v122
		v157 = v123
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v133 = v121
	v134 = v122
	v135 = v123
	goto L45
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v141 = int32(-2139062144)
	if (int32(16843008)-v138|v138)&v141 != v141 {
		v162 = v133
		v163 = v134
		v164 = v135
		goto L28
	} else {
		goto L47
	}
L46:
	;
	v155 = v149
	v156 = v147
	v157 = v151
	goto L29
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v138
	v146 = int32(4)
	v147 = v134 + v146
	v149 = v133 + v146
	v151 = v135 - v146
	if base.Ui32(int32(3)) < base.Ui32(v151) {
		v133 = v149
		v134 = v147
		v135 = v151
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v162 = v155
	v163 = v156
	v164 = v157
	goto L28
L50:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
	if v171 == int32(0) {
		v182 = v167
		v183 = v168
		goto L27
	} else {
		goto L52
	}
L51:
	;
	v182 = v178
	v183 = v176
	goto L27
L52:
	;
	v175 = int32(1)
	v176 = v168 + v175
	v178 = v167 + v175
	v180 = v169 - v175
	if v180 != 0 {
		v167 = v178
		v168 = v176
		v169 = v180
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v199 = v16 + int32(16)
	v205 = v196
	goto L57
L55:
	;
	goto L56
L56:
	;
	v245 = v16 + int32(11)
	v247 = v16 + int32(12)
	v249 = v16 + int32(4)
	v250 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v76)+268))
	if v256 <= v250 {
		v344 = v250
		goto L65
	} else {
		goto L66
	}
L57:
	;
	v212 = int32(255)
	v213 = v205 & v212
	if base.Ui32((v213-int32(97))&v212) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L56
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v224)
	v227 = v199 + int32(1)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v228 != 0 {
		v199 = v227
		v205 = v228
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v222 = v213 - int32(32)
	goto L62
L61:
	;
	v222 = v213
	goto L62
L62:
	;
	v224 = v222 & int32(255)
	goto L59
L63:
	;
	goto L58
L64:
	;
	if v344 == int32(0) {
		goto L20
	} else {
		goto L89
	}
L65:
	;
	goto L64
L66:
	;
	v260 = v76 + int32(22376)
	v270 = v250
	goto L67
L67:
	;
	v273 = F_strcmp(m, v16+int32(16), v260+v270)
	mBase = m.M
	if v273 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v76)+264))
	if v290 <= int32(0) {
		v344 = v250
		goto L65
	} else {
		goto L76
	}
L69:
	;
	v279 = v270
	goto L72
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	v287 = v279 + int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v260))))
	if v288 != 0 {
		v279 = v287
		goto L72
	} else {
		goto L74
	}
L73:
	;
	if v287 < v256 {
		v270 = v287
		goto L67
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v344 = v250
	goto L65
L76:
	;
	v301 = int32(0)
	v302 = v250
	v304 = v290
	goto L77
L77:
	;
	v309 = v76 + int32(18280) + v301<<(uint(int32(4))%32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v310 != v270 {
		v333 = v302
		v334 = v304
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v344 = v333
	goto L65
L79:
	;
	v336 = v301 + int32(1)
	if v336 < v334 {
		v301 = v336
		v302 = v333
		v304 = v334
		goto L77
	} else {
		goto L88
	}
L80:
	;
	if v302 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v314)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v317
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v76)+264))
	v333 = v314
	v334 = v321
	goto L79
L82:
	;
	goto L83
L83:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	if v322 == v323 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+4)))
	if v326 == v327 {
		v333 = int32(1)
		v334 = v304
		goto L79
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v330)
	v344 = int32(1)
	goto L65
L87:
	;
	goto L86
L88:
	;
	goto L78
L89:
	;
	v351 = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v351 - v353
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v358 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v359 = int32(6)
	goto L92
L91:
	;
	v359 = int32(5)
	goto L92
L92:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v361 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v362 = v359
	goto L95
L94:
	;
	v362 = int32(7)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v362
	v366 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v361 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v367 = int32(0)
	goto L98
L97:
	;
	v367 = v366
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v367
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0) - v370
	goto L102
L99:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[960]))) = uint8(v488)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[961]))) = v490
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[962]))) = v492
	v790 = v351
	goto L1
L100:
	;
	v485 = F_strlen(m, v474)
	mBase = m.M
	goto L99
L102:
	;
	goto L103
L103:
	;
	v379 = int32(10)
	if (v21^l1)&int32(3) != 0 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v478 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v475))) = uint8(v478)
	goto L100
L105:
	;
	v459 = v454
	v460 = v455
	v461 = v456
	goto L127
L106:
	;
	if v449 == int32(0) {
		v474 = v447
		v475 = v448
		goto L104
	} else {
		goto L126
	}
L107:
	;
	v447 = l1
	v448 = v21
	v449 = v379
	goto L106
L108:
	;
	goto L109
L109:
	;
	if l1&int32(3) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v416 == int32(0) {
		v474 = v413
		v475 = v414
		goto L104
	} else {
		goto L119
	}
L111:
	;
	v413 = l1
	v414 = v21
	v415 = v379
	v416 = int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v392 = l1
	v393 = v21
	v394 = v379
	goto L114
L114:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	*(*uint8)(unsafe.Add(mBase, uint32(v393))) = uint8(v396)
	if v396 == int32(0) {
		v454 = v392
		v455 = v393
		v456 = v394
		goto L105
	} else {
		goto L116
	}
L115:
	;
	v413 = v407
	v414 = v401
	v415 = v403
	v416 = v405
	goto L110
L116:
	;
	v400 = int32(1)
	v401 = v393 + v400
	v403 = v394 - v400
	v404 = int32(0)
	v405 = base.B2i32(v403 != v404)
	v407 = v392 + v400
	if v407&int32(3) == v404 {
		v413 = v407
		v414 = v401
		v415 = v403
		v416 = v405
		goto L110
	} else {
		goto L117
	}
L117:
	;
	if v403 != 0 {
		v392 = v407
		v393 = v401
		v394 = v403
		goto L114
	} else {
		goto L118
	}
L118:
	;
	goto L115
L119:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	if v419 == int32(0) {
		v447 = v413
		v448 = v414
		v449 = v415
		goto L106
	} else {
		goto L120
	}
L120:
	;
	if base.Ui32(v415) < base.Ui32(int32(4)) {
		v447 = v413
		v448 = v414
		v449 = v415
		goto L106
	} else {
		goto L121
	}
L121:
	;
	v425 = v413
	v426 = v414
	v427 = v415
	goto L122
L122:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v433 = int32(-2139062144)
	if (int32(16843008)-v430|v430)&v433 != v433 {
		v454 = v425
		v455 = v426
		v456 = v427
		goto L105
	} else {
		goto L124
	}
L123:
	;
	v447 = v441
	v448 = v439
	v449 = v443
	goto L106
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v430
	v438 = int32(4)
	v439 = v426 + v438
	v441 = v425 + v438
	v443 = v427 - v438
	if base.Ui32(int32(3)) < base.Ui32(v443) {
		v425 = v441
		v426 = v439
		v427 = v443
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v454 = v447
	v455 = v448
	v456 = v449
	goto L105
L127:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v463)
	if v463 == int32(0) {
		v474 = v459
		v475 = v460
		goto L104
	} else {
		goto L129
	}
L128:
	;
	v474 = v470
	v475 = v468
	goto L104
L129:
	;
	v467 = int32(1)
	v468 = v460 + v467
	v470 = v459 + v467
	v472 = v461 - v467
	if v472 != 0 {
		v459 = v470
		v460 = v468
		v461 = v472
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v623 = int32(*(*int8)(unsafe.Add(mBase, uint32(v541)+11)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v623
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+11)))
	if v625 == int32(7) {
		goto L165
	} else {
		goto L166
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(31)
	v618 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v618
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v618
	v790 = v618
	goto L1
L133:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	if v511 <= int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v515 = v508 + int32(8)
	v520 = v515 + v511<<(uint(int32(4))%32) - int32(16)
	if base.Ui32(v520) < base.Ui32(v515) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v522 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v523 = v515
	v532 = v520
	goto L136
L136:
	;
	v541 = v523 + (v532-v523)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v542 = int32(*(*int8)(unsafe.Add(mBase, uint32(v541))))
	v543 = v522 - v542
	if v543 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L132
L138:
	;
	goto L143
L139:
	;
	v593 = v543
	goto L140
L140:
	;
	v597 = base.B2i32(v593 < int32(0))
	if v593 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L141:
	;
	if v584 == int32(0) {
		goto L131
	} else {
		goto L155
	}
L143:
	;
	goto L144
L144:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v552 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v553 = l1
	v554 = v541
	v555 = int32(10)
	v556 = v552
	goto L149
L146:
	;
	v578 = v541
	v582 = int32(0)
	goto L147
L147:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v584 = v582 - v583
	goto L141
L148:
	;
	v578 = v573
	v582 = v575
	goto L147
L149:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554))))
	if v556 != v558 {
		v573 = v554
		v575 = v556
		goto L148
	} else {
		goto L151
	}
L150:
	;
	v573 = v567
	v575 = int32(0)
	goto L148
L151:
	;
	if v558 == int32(0) {
		v573 = v554
		v575 = v556
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v563 = v555 - int32(1)
	if v563 == int32(0) {
		v573 = v554
		v575 = v556
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v566 = int32(1)
	v567 = v554 + v566
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)))
	if v568 != 0 {
		v553 = v553 + v566
		v554 = v567
		v555 = v563
		v556 = v568
		goto L149
	} else {
		goto L154
	}
L154:
	;
	goto L150
L155:
	;
	v593 = v584
	goto L140
L156:
	;
	v598 = v541 - int32(16)
	goto L158
L157:
	;
	v598 = v532
	goto L158
L158:
	;
	if v593 < int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v601 = v523
	goto L161
L160:
	;
	v601 = v541 + int32(16)
	goto L161
L161:
	;
	if base.Ui32(v601) <= base.Ui32(v598) {
		v523 = v601
		v532 = v598
		goto L136
	} else {
		goto L162
	}
L162:
	;
	goto L137
L163:
	;
	goto L176
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v648
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v631 = v508 + v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	if v632 != 0 {
		v648 = v632
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v645
	v648 = int32(0)
	goto L164
L168:
	;
	v634 = v631 + int32(4)
	v635 = F_pg_tzset(m, v634)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	return int32(0)
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631))) = v635
	if v635 != 0 {
		v648 = v635
		goto L164
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v634
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v642
	if v642 != 0 {
		goto L163
	} else {
		goto L172
	}
L172:
	;
	v790 = int32(-7)
	goto L1
L173:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[960]))) = uint8(v770)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[961]))) = v772
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[962]))) = v774
	goto L2
L174:
	;
	v767 = F_strlen(m, v756)
	mBase = m.M
	goto L173
L176:
	;
	goto L177
L177:
	;
	v661 = int32(10)
	if (v21^l1)&int32(3) != 0 {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	v760 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v757))) = uint8(v760)
	goto L174
L179:
	;
	v741 = v736
	v742 = v737
	v743 = v738
	goto L201
L180:
	;
	if v731 == int32(0) {
		v756 = v729
		v757 = v730
		goto L178
	} else {
		goto L200
	}
L181:
	;
	v729 = l1
	v730 = v21
	v731 = v661
	goto L180
L182:
	;
	goto L183
L183:
	;
	if l1&int32(3) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v698 == int32(0) {
		v756 = v695
		v757 = v696
		goto L178
	} else {
		goto L193
	}
L185:
	;
	v695 = l1
	v696 = v21
	v697 = v661
	v698 = int32(1)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v674 = l1
	v675 = v21
	v676 = v661
	goto L188
L188:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	*(*uint8)(unsafe.Add(mBase, uint32(v675))) = uint8(v678)
	if v678 == int32(0) {
		v736 = v674
		v737 = v675
		v738 = v676
		goto L179
	} else {
		goto L190
	}
L189:
	;
	v695 = v689
	v696 = v683
	v697 = v685
	v698 = v687
	goto L184
L190:
	;
	v682 = int32(1)
	v683 = v675 + v682
	v685 = v676 - v682
	v686 = int32(0)
	v687 = base.B2i32(v685 != v686)
	v689 = v674 + v682
	if v689&int32(3) == v686 {
		v695 = v689
		v696 = v683
		v697 = v685
		v698 = v687
		goto L184
	} else {
		goto L191
	}
L191:
	;
	if v685 != 0 {
		v674 = v689
		v675 = v683
		v676 = v685
		goto L188
	} else {
		goto L192
	}
L192:
	;
	goto L189
L193:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	if v701 == int32(0) {
		v729 = v695
		v730 = v696
		v731 = v697
		goto L180
	} else {
		goto L194
	}
L194:
	;
	if base.Ui32(v697) < base.Ui32(int32(4)) {
		v729 = v695
		v730 = v696
		v731 = v697
		goto L180
	} else {
		goto L195
	}
L195:
	;
	v707 = v695
	v708 = v696
	v709 = v697
	goto L196
L196:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v715 = int32(-2139062144)
	if (int32(16843008)-v712|v712)&v715 != v715 {
		v736 = v707
		v737 = v708
		v738 = v709
		goto L179
	} else {
		goto L198
	}
L197:
	;
	v729 = v723
	v730 = v721
	v731 = v725
	goto L180
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = v712
	v720 = int32(4)
	v721 = v708 + v720
	v723 = v707 + v720
	v725 = v709 - v720
	if base.Ui32(int32(3)) < base.Ui32(v725) {
		v707 = v723
		v708 = v721
		v709 = v725
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v736 = v729
	v737 = v730
	v738 = v731
	goto L179
L201:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	*(*uint8)(unsafe.Add(mBase, uint32(v742))) = uint8(v745)
	if v745 == int32(0) {
		v756 = v741
		v757 = v742
		goto L178
	} else {
		goto L203
	}
L202:
	;
	v756 = v752
	v757 = v750
	goto L178
L203:
	;
	v749 = int32(1)
	v750 = v742 + v749
	v752 = v741 + v749
	v754 = v743 - v749
	if v754 != 0 {
		v741 = v752
		v742 = v750
		v743 = v754
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
}
func F_show_timezone(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(253980)
	}
	return v5
}
