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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v15 - int32(43) {
	case 0, 2:
		*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0])) = int32(0)
		v26 = F_strtol(m, l0+int32(1), v12+int32(12), int32(10))
		mBase = m.M
		v27 = int32(-5)
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0]))
		if v29 == int32(68) {
			v108 = v27
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			if v33 != 0 {
				if v33 != int32(58) {
					v74 = int32(0)
					v75 = v26
					v77 = v3
					v81 = int32(59)
					if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v75))|base.B2i32(base.Ui32(v81) < base.Ui32(v74))|base.B2i32(base.Ui32(v81) < base.Ui32(v77)) != 0 {
						v108 = v27
					} else {
						v87 = int32(60)
						v92 = (v75*v87+v74)*v87 + v77
						v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						if v95 == int32(45) {
							v98 = v92
						} else {
							v98 = int32(0) - v92
						}
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
						if v103 != 0 {
							v104 = int32(-1)
						} else {
							v104 = int32(0)
						}
						v108 = v104
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0])) = int32(0)
					v43 = v12 + int32(12)
					v45 = F_strtol(m, v32+int32(1), v43, int32(10))
					mBase = m.M
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0]))
					if v47 == int32(68) {
						v108 = v27
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
						if v51 != int32(58) {
							v74 = v45
							v75 = v26
							v77 = v3
							v81 = int32(59)
							if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v75))|base.B2i32(base.Ui32(v81) < base.Ui32(v74))|base.B2i32(base.Ui32(v81) < base.Ui32(v77)) != 0 {
								v108 = v27
							} else {
								v87 = int32(60)
								v92 = (v75*v87+v74)*v87 + v77
								v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if v95 == int32(45) {
									v98 = v92
								} else {
									v98 = int32(0) - v92
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
								if v103 != 0 {
									v104 = int32(-1)
								} else {
									v104 = int32(0)
								}
								v108 = v104
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0])) = int32(0)
							v60 = F_strtol(m, v50+int32(1), v43, int32(10))
							mBase = m.M
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezone[0]))
							if v62 != int32(68) {
								v74 = v45
								v75 = v26
								v77 = v60
								v81 = int32(59)
								if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v75))|base.B2i32(base.Ui32(v81) < base.Ui32(v74))|base.B2i32(base.Ui32(v81) < base.Ui32(v77)) != 0 {
									v108 = v27
								} else {
									v87 = int32(60)
									v92 = (v75*v87+v74)*v87 + v77
									v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
									if v95 == int32(45) {
										v98 = v92
									} else {
										v98 = int32(0) - v92
									}
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
									if v103 != 0 {
										v104 = int32(-1)
									} else {
										v104 = int32(0)
									}
									v108 = v104
								}
							} else {
								v108 = v27
							}
						}
					}
				}
			} else {
				v65 = F_strlen(m, l0)
				mBase = m.M
				if base.Ui32(v65) < base.Ui32(int32(4)) {
					v74 = int32(0)
					v75 = v26
					v77 = v3
				} else {
					v69 = int32(100)
					v70 = base.I32_div_s(v26, v69)
					v74 = v26 - v70*v69
					v75 = v70
					v77 = v3
				}
				v81 = int32(59)
				if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v75))|base.B2i32(base.Ui32(v81) < base.Ui32(v74))|base.B2i32(base.Ui32(v81) < base.Ui32(v77)) != 0 {
					v108 = v27
				} else {
					v87 = int32(60)
					v92 = (v75*v87+v74)*v87 + v77
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v95 == int32(45) {
						v98 = v92
					} else {
						v98 = int32(0) - v92
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v98
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
					if v103 != 0 {
						v104 = int32(-1)
					} else {
						v104 = int32(0)
					}
					v108 = v104
				}
			}
		}
	default:
		v108 = int32(-1)
	}
	m.G0 = v12 + int32(16)
	return v108
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v784 int32
	_ = v784
	v14 = m.G0
	v16 = v14 - int32(272)
	m.G0 = v16
	v19 = l0 * int32(20)
	v21 = v19 + int32(_a_F_DecodeTimezoneAbbrev_0)
	goto L5
L1:
	;
	m.G0 = v16 + int32(272)
	return v784
L2:
	;
	v784 = int32(0)
	goto L1
L3:
	;
	if v59-v60 == int32(0) {
		goto L16
	} else {
		goto L17
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
	v55 = v21
	v59 = int32(0)
	goto L9
L9:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	goto L3
L10:
	;
	v55 = v50
	v59 = v52
	goto L9
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if base.B2i32(v32 != v34)|base.B2i32(v34 == int32(0)) != 0 {
		v50 = v30
		v52 = v32
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v50 = v44
	v52 = int32(0)
	goto L10
L13:
	;
	v40 = v31 - int32(1)
	if v40 == int32(0) {
		v50 = v30
		v52 = v32
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v43 = int32(1)
	v44 = v30 + v43
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v45 != 0 {
		v29 = v29 + v43
		v30 = v44
		v31 = v40
		v32 = v45
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[0]))))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[1])))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[2])))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v74
	goto L2
L17:
	;
	goto L18
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezoneAbbrev[3]))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezoneAbbrev[4]))
	if v498 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L20:
	;
	v81 = v16 + int32(16)
	goto L24
L21:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v201 != 0 {
		goto L52
	} else {
		goto L53
	}
L22:
	;
	v198 = F_strlen(m, v187)
	mBase = m.M
	goto L21
L24:
	;
	goto L25
L25:
	;
	v88 = int32(255)
	if (v81^l1)&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v191)
	goto L22
L27:
	;
	v172 = v167
	v173 = v168
	v174 = v169
	goto L48
L28:
	;
	if v162 == int32(0) {
		v187 = v160
		v188 = v161
		goto L26
	} else {
		goto L47
	}
L29:
	;
	v160 = l1
	v161 = v81
	v162 = v88
	goto L28
L30:
	;
	goto L31
L31:
	;
	v92 = int32(0)
	if base.B2i32(l1&int32(3) == v92)|int32(0) == v92 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v128 == int32(0) {
		v187 = v125
		v188 = v126
		goto L26
	} else {
		goto L41
	}
L33:
	;
	v104 = l1
	v105 = v81
	v106 = v88
	goto L36
L34:
	;
	goto L35
L35:
	;
	v125 = l1
	v126 = v81
	v127 = v88
	v128 = int32(1)
	goto L32
L36:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	if v108 == int32(0) {
		v167 = v104
		v168 = v105
		v169 = v106
		goto L27
	} else {
		goto L38
	}
L37:
	;
	v125 = v119
	v126 = v113
	v127 = v115
	v128 = v117
	goto L32
L38:
	;
	v112 = int32(1)
	v113 = v105 + v112
	v115 = v106 - v112
	v116 = int32(0)
	v117 = base.B2i32(v115 != v116)
	v119 = v104 + v112
	if v119&int32(3) == v116 {
		v125 = v119
		v126 = v113
		v127 = v115
		v128 = v117
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v115 != 0 {
		v104 = v119
		v105 = v113
		v106 = v115
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if base.B2i32(v131 == int32(0))|base.B2i32(base.Ui32(v127) < base.Ui32(int32(4))) != 0 {
		v160 = v125
		v161 = v126
		v162 = v127
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v138 = v125
	v139 = v126
	v140 = v127
	goto L43
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v146 = int32(-2139062144)
	if (int32(16843008)-v143|v143)&v146 != v146 {
		v167 = v138
		v168 = v139
		v169 = v140
		goto L27
	} else {
		goto L45
	}
L44:
	;
	v160 = v154
	v161 = v152
	v162 = v156
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v143
	v151 = int32(4)
	v152 = v139 + v151
	v154 = v138 + v151
	v156 = v140 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v138 = v154
		v139 = v152
		v140 = v156
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L27
L48:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v176)
	if v176 == int32(0) {
		v187 = v172
		v188 = v173
		goto L26
	} else {
		goto L50
	}
L49:
	;
	v187 = v183
	v188 = v181
	goto L26
L50:
	;
	v180 = int32(1)
	v181 = v173 + v180
	v183 = v172 + v180
	v185 = v174 - v180
	if v185 != 0 {
		v172 = v183
		v173 = v181
		v174 = v185
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v202 = v81
	v208 = v201
	goto L55
L53:
	;
	goto L54
L54:
	;
	v246 = v16 + int32(11)
	v248 = v16 + int32(12)
	v250 = v16 + int32(4)
	v251 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v77)+268))
	if v256 <= v251 {
		v332 = v251
		goto L63
	} else {
		goto L64
	}
L55:
	;
	if base.Ui32((v208-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L54
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v225)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v227 != 0 {
		v202 = v202 + int32(1)
		v208 = v227
		goto L55
	} else {
		goto L61
	}
L58:
	;
	v223 = v208 - int32(32)
	goto L60
L59:
	;
	v223 = v208
	goto L60
L60:
	;
	v225 = v223 & int32(255)
	goto L57
L61:
	;
	goto L56
L62:
	;
	if v332 == int32(0) {
		goto L19
	} else {
		goto L84
	}
L63:
	;
	goto L62
L64:
	;
	v269 = v251
	goto L65
L65:
	;
	v271 = v77 + int32(_a_F_DecodeTimezoneAbbrev_1) + v269
	v272 = F_strcmp(m, v16+int32(16), v271)
	mBase = m.M
	if v272 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v77)+264))
	if v278 <= int32(0) {
		v332 = v251
		goto L63
	} else {
		goto L71
	}
L67:
	;
	v273 = F_strlen(m, v271)
	mBase = m.M
	v276 = v273 + v269 + int32(1)
	if v276 < v256 {
		v269 = v276
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	v332 = v251
	goto L63
L71:
	;
	v289 = int32(0)
	v290 = v278
	v291 = v251
	goto L72
L72:
	;
	v296 = v77 + int32(_a_F_DecodeTimezoneAbbrev_2) + v289<<(uint(int32(4))%32)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	if v297 != v269 {
		v320 = v290
		v321 = v291
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v332 = v321
	goto L63
L74:
	;
	v323 = v289 + int32(1)
	if v323 < v320 {
		v289 = v323
		v290 = v320
		v291 = v321
		goto L72
	} else {
		goto L83
	}
L75:
	;
	if v291 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v301 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v301)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v304
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v77)+264))
	v320 = v308
	v321 = v301
	goto L74
L77:
	;
	goto L78
L78:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v309 == v310 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+4)))
	if v313 == v314 {
		v320 = v290
		v321 = int32(1)
		goto L74
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v317)
	v332 = int32(1)
	goto L63
L82:
	;
	goto L81
L83:
	;
	goto L73
L84:
	;
	v337 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v337 - v339
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v344 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v345 = int32(6)
	goto L87
L86:
	;
	v345 = int32(5)
	goto L87
L87:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v347 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v348 = v345
	goto L90
L89:
	;
	v348 = int32(7)
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v348
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimezoneAbbrev[3]))
	if v347 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v353 = int32(0)
	goto L93
L92:
	;
	v353 = v352
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v353
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0) - v356
	goto L97
L94:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[0]))) = uint8(v478)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[1]))) = v480
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[2]))) = v482
	v784 = v337
	goto L1
L95:
	;
	v475 = F_strlen(m, v464)
	mBase = m.M
	goto L94
L97:
	;
	goto L98
L98:
	;
	v365 = int32(10)
	if (v21^l1)&int32(3) != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v468 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v468)
	goto L95
L100:
	;
	v449 = v444
	v450 = v445
	v451 = v446
	goto L121
L101:
	;
	if v439 == int32(0) {
		v464 = v437
		v465 = v438
		goto L99
	} else {
		goto L120
	}
L102:
	;
	v437 = l1
	v438 = v21
	v439 = v365
	goto L101
L103:
	;
	goto L104
L104:
	;
	v369 = int32(0)
	if base.B2i32(l1&int32(3) == v369)|int32(0) == v369 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v405 == int32(0) {
		v464 = v402
		v465 = v403
		goto L99
	} else {
		goto L114
	}
L106:
	;
	v381 = l1
	v382 = v21
	v383 = v365
	goto L109
L107:
	;
	goto L108
L108:
	;
	v402 = l1
	v403 = v21
	v404 = v365
	v405 = int32(1)
	goto L105
L109:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v385)
	if v385 == int32(0) {
		v444 = v381
		v445 = v382
		v446 = v383
		goto L100
	} else {
		goto L111
	}
L110:
	;
	v402 = v396
	v403 = v390
	v404 = v392
	v405 = v394
	goto L105
L111:
	;
	v389 = int32(1)
	v390 = v382 + v389
	v392 = v383 - v389
	v393 = int32(0)
	v394 = base.B2i32(v392 != v393)
	v396 = v381 + v389
	if v396&int32(3) == v393 {
		v402 = v396
		v403 = v390
		v404 = v392
		v405 = v394
		goto L105
	} else {
		goto L112
	}
L112:
	;
	if v392 != 0 {
		v381 = v396
		v382 = v390
		v383 = v392
		goto L109
	} else {
		goto L113
	}
L113:
	;
	goto L110
L114:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	if base.B2i32(v408 == int32(0))|base.B2i32(base.Ui32(v404) < base.Ui32(int32(4))) != 0 {
		v437 = v402
		v438 = v403
		v439 = v404
		goto L101
	} else {
		goto L115
	}
L115:
	;
	v415 = v402
	v416 = v403
	v417 = v404
	goto L116
L116:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v423 = int32(-2139062144)
	if (int32(16843008)-v420|v420)&v423 != v423 {
		v444 = v415
		v445 = v416
		v446 = v417
		goto L100
	} else {
		goto L118
	}
L117:
	;
	v437 = v431
	v438 = v429
	v439 = v433
	goto L101
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v420
	v428 = int32(4)
	v429 = v416 + v428
	v431 = v415 + v428
	v433 = v417 - v428
	if base.Ui32(int32(3)) < base.Ui32(v433) {
		v415 = v431
		v416 = v429
		v417 = v433
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v444 = v437
	v445 = v438
	v446 = v439
	goto L100
L121:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v453)
	if v453 == int32(0) {
		v464 = v449
		v465 = v450
		goto L99
	} else {
		goto L123
	}
L122:
	;
	v464 = v460
	v465 = v458
	goto L99
L123:
	;
	v457 = int32(1)
	v458 = v450 + v457
	v460 = v449 + v457
	v462 = v451 - v457
	if v462 != 0 {
		v449 = v460
		v450 = v458
		v451 = v462
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v613 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530)+11)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v613
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+11)))
	if v615 == int32(7) {
		goto L157
	} else {
		goto L158
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(31)
	v608 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v608
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v608
	v784 = v608
	goto L1
L127:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v501 <= int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v505 = v498 + int32(8)
	v511 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	v512 = v505
	v522 = v505 + v501<<(uint(int32(4))%32) - int32(16)
	goto L129
L129:
	;
	v530 = v512 + (v522-v512)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530))))
	v532 = v511 - v531
	if v532 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L126
L131:
	;
	goto L136
L132:
	;
	v583 = v532
	goto L133
L133:
	;
	v587 = base.B2i32(v583 < int32(0))
	if v583 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L134:
	;
	if v574 == int32(0) {
		goto L125
	} else {
		goto L147
	}
L136:
	;
	goto L137
L137:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v541 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v542 = l1
	v543 = v530
	v544 = int32(10)
	v545 = v541
	goto L142
L139:
	;
	v568 = v530
	v572 = int32(0)
	goto L140
L140:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568))))
	v574 = v572 - v573
	goto L134
L141:
	;
	v568 = v563
	v572 = v565
	goto L140
L142:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if base.B2i32(v545 != v547)|base.B2i32(v547 == int32(0)) != 0 {
		v563 = v543
		v565 = v545
		goto L141
	} else {
		goto L144
	}
L143:
	;
	v563 = v557
	v565 = int32(0)
	goto L141
L144:
	;
	v553 = v544 - int32(1)
	if v553 == int32(0) {
		v563 = v543
		v565 = v545
		goto L141
	} else {
		goto L145
	}
L145:
	;
	v556 = int32(1)
	v557 = v543 + v556
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+1)))
	if v558 != 0 {
		v542 = v542 + v556
		v543 = v557
		v544 = v553
		v545 = v558
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v583 = v574
	goto L133
L148:
	;
	v588 = v530 - int32(16)
	goto L150
L149:
	;
	v588 = v522
	goto L150
L150:
	;
	if v583 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v591 = v512
	goto L153
L152:
	;
	v591 = v530 + int32(16)
	goto L153
L153:
	;
	if base.Ui32(v591) <= base.Ui32(v588) {
		v512 = v591
		v522 = v588
		goto L129
	} else {
		goto L154
	}
L154:
	;
	goto L130
L155:
	;
	goto L168
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v638
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v530)+12))
	v621 = v498 + v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	if v622 != 0 {
		v638 = v622
		goto L156
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v530)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v635
	v638 = int32(0)
	goto L156
L160:
	;
	v624 = v621 + int32(4)
	v625 = F_pg_tzset(m, v624)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	return int32(0)
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v625
	if v625 != 0 {
		v638 = v625
		goto L156
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v624
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v632
	if v632 != 0 {
		goto L155
	} else {
		goto L164
	}
L164:
	;
	v784 = int32(-7)
	goto L1
L165:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[0]))) = uint8(v764)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[1]))) = v766
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_DecodeTimezoneAbbrev[2]))) = v768
	goto L2
L166:
	;
	v761 = F_strlen(m, v750)
	mBase = m.M
	goto L165
L168:
	;
	goto L169
L169:
	;
	v651 = int32(10)
	if (v21^l1)&int32(3) != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	v754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v751))) = uint8(v754)
	goto L166
L171:
	;
	v735 = v730
	v736 = v731
	v737 = v732
	goto L192
L172:
	;
	if v725 == int32(0) {
		v750 = v723
		v751 = v724
		goto L170
	} else {
		goto L191
	}
L173:
	;
	v723 = l1
	v724 = v21
	v725 = v651
	goto L172
L174:
	;
	goto L175
L175:
	;
	v655 = int32(0)
	if base.B2i32(l1&int32(3) == v655)|int32(0) == v655 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	if v691 == int32(0) {
		v750 = v688
		v751 = v689
		goto L170
	} else {
		goto L185
	}
L177:
	;
	v667 = l1
	v668 = v21
	v669 = v651
	goto L180
L178:
	;
	goto L179
L179:
	;
	v688 = l1
	v689 = v21
	v690 = v651
	v691 = int32(1)
	goto L176
L180:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	*(*uint8)(unsafe.Add(mBase, uint32(v668))) = uint8(v671)
	if v671 == int32(0) {
		v730 = v667
		v731 = v668
		v732 = v669
		goto L171
	} else {
		goto L182
	}
L181:
	;
	v688 = v682
	v689 = v676
	v690 = v678
	v691 = v680
	goto L176
L182:
	;
	v675 = int32(1)
	v676 = v668 + v675
	v678 = v669 - v675
	v679 = int32(0)
	v680 = base.B2i32(v678 != v679)
	v682 = v667 + v675
	if v682&int32(3) == v679 {
		v688 = v682
		v689 = v676
		v690 = v678
		v691 = v680
		goto L176
	} else {
		goto L183
	}
L183:
	;
	if v678 != 0 {
		v667 = v682
		v668 = v676
		v669 = v678
		goto L180
	} else {
		goto L184
	}
L184:
	;
	goto L181
L185:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688))))
	if base.B2i32(v694 == int32(0))|base.B2i32(base.Ui32(v690) < base.Ui32(int32(4))) != 0 {
		v723 = v688
		v724 = v689
		v725 = v690
		goto L172
	} else {
		goto L186
	}
L186:
	;
	v701 = v688
	v702 = v689
	v703 = v690
	goto L187
L187:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v709 = int32(-2139062144)
	if (int32(16843008)-v706|v706)&v709 != v709 {
		v730 = v701
		v731 = v702
		v732 = v703
		goto L171
	} else {
		goto L189
	}
L188:
	;
	v723 = v717
	v724 = v715
	v725 = v719
	goto L172
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v706
	v714 = int32(4)
	v715 = v702 + v714
	v717 = v701 + v714
	v719 = v703 - v714
	if base.Ui32(int32(3)) < base.Ui32(v719) {
		v701 = v717
		v702 = v715
		v703 = v719
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v730 = v723
	v731 = v724
	v732 = v725
	goto L171
L192:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	*(*uint8)(unsafe.Add(mBase, uint32(v736))) = uint8(v739)
	if v739 == int32(0) {
		v750 = v735
		v751 = v736
		goto L170
	} else {
		goto L194
	}
L193:
	;
	v750 = v746
	v751 = v744
	goto L170
L194:
	;
	v743 = int32(1)
	v744 = v736 + v743
	v746 = v735 + v743
	v748 = v737 - v743
	if v748 != 0 {
		v735 = v746
		v736 = v744
		v737 = v748
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
}
func F_show_timezone(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_show_timezone[0]))
	if v3 != 0 {
		v5 = v3
	} else {
		v5 = int32(_a_F_show_timezone_0)
	}
	return v5
}
