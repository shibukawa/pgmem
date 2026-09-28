package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildGroupingEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
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
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v469 int64
	_ = v469
	var v471 int64
	_ = v471
	var v476 int32
	_ = v476
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int64
	_ = v599
	var v601 int64
	_ = v601
	var v603 int64
	_ = v603
	var v605 int64
	_ = v605
	var v607 int64
	_ = v607
	var v612 int32
	_ = v612
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v660 int64
	_ = v660
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int64
	_ = v702
	var v704 int64
	_ = v704
	var v706 int64
	_ = v706
	var v708 int64
	_ = v708
	var v710 int64
	_ = v710
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v752 int64
	_ = v752
	var v754 int64
	_ = v754
	var v756 int64
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v781 int32
	_ = v781
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int64
	_ = v860
	var v862 int64
	_ = v862
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v884 int32
	_ = v884
	v10 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v25 = F_palloc0(m, int32(72))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(386)
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v31
	if l4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v22 + int32(48)
	return v884
L4:
	;
	v884 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = l8
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)) = uint8(v43)
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v45
	v48 = v25 + int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v48
	v51 = v25 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v51
	if l4 <= v45 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(2)
	v179 = v22 + int32(8)
	v180 = int32(0)
	v183 = m.G0
	v185 = v183 - int32(16)
	m.G0 = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)) = uint8(v180)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	if v190 != 0 {
		goto L42
	} else {
		goto L43
	}
L8:
	;
	v162 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v57 = l4 & int32(3)
	v58 = int32(-1)
	if base.Ui32(int32(4)) <= base.Ui32(l4) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v71 = v58
	v74 = v10
	v79 = v10
	goto L14
L12:
	;
	v112 = v58
	v115 = v10
	goto L13
L13:
	;
	v132 = v112
	v135 = v115
	v136 = int32(0)
	goto L30
L14:
	;
	v84 = l5 + v74<<(uint(int32(1))%32)
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84))))
	if v85 < v71 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v57 == int32(0) {
		v162 = v96
		goto L7
	} else {
		goto L29
	}
L16:
	;
	v87 = v71
	goto L18
L17:
	;
	v87 = v85
	goto L18
L18:
	;
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+2)))
	if v88 < v87 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = v87
	goto L21
L20:
	;
	v90 = v88
	goto L21
L21:
	;
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+4)))
	if v91 < v90 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v93 = v90
	goto L24
L23:
	;
	v93 = v91
	goto L24
L24:
	;
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+6)))
	if v94 < v93 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = v93
	goto L27
L26:
	;
	v96 = v94
	goto L27
L27:
	;
	v97 = int32(4)
	v98 = v74 + v97
	v100 = v79 + v97
	if v100 != l4&int32(2147483644) {
		v71 = v96
		v74 = v98
		v79 = v100
		goto L14
	} else {
		goto L28
	}
L28:
	;
	goto L15
L29:
	;
	v112 = v96
	v115 = v98
	goto L13
L30:
	;
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v135<<(uint(int32(1))%32)))))
	if v146 < v132 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v162 = v148
	goto L7
L32:
	;
	v148 = v132
	goto L34
L33:
	;
	v148 = v146
	goto L34
L34:
	;
	v149 = int32(1)
	v152 = v136 + v149
	if v152 != v57 {
		v132 = v148
		v135 = v135 + v149
		v136 = v152
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	if v277 != 0 {
		goto L64
	} else {
		goto L65
	}
L37:
	;
	m.G0 = v185 + int32(16)
	goto L36
L38:
	;
	v277 = int32(1)
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+28)) = v252
	v266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+24)) = v251
	if v252 != int32(_a_F_ExecBuildGroupingEqual_0) {
		goto L38
	} else {
		goto L63
	}
L40:
	;
	v261 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v261)
	*(*int64)(unsafe.Add(mBase, uint32(v179)+24)) = int64(0)
	goto L38
L41:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)))
	if base.B2i32(v251 == int32(0))|base.B2i32(v255 != int32(1)) != 0 {
		goto L40
	} else {
		goto L61
	}
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)) = uint8(base.B2i32(v191 != int32(0)))
	v251 = v190
	v252 = v191
	goto L41
L43:
	;
	goto L44
L44:
	;
	if v187 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	switch v197 - int32(2) {
	case 0:
		goto L48
	case 1:
		goto L47
	default:
		goto L46
	}
L46:
	;
	if base.Ui32(int32(2)) < base.Ui32(v197-int32(4)) {
		goto L40
	} else {
		goto L59
	}
L47:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v187)+36))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+101)))
	if v221 != int32(1) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v187)+40))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+102)))
	if v201 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v200 == int32(0) {
		goto L40
	} else {
		goto L53
	}
L50:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+98)))
	if v204 != int32(1) {
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v187)+88))
	if v207 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)) = uint8(v210)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v200)+56))
	v251 = v212
	v252 = v207
	goto L41
L53:
	;
	v218 = F_ExecGetResultSlotOps(m, v200, v185+int32(15))
	mBase = m.M
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v200)+56))
	v251 = v219
	v252 = v218
	goto L41
L54:
	;
	if v220 == int32(0) {
		goto L40
	} else {
		goto L58
	}
L55:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+97)))
	if v224 != int32(1) {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v187)+84))
	if v227 == int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)) = uint8(v230)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)+56))
	v251 = v232
	v252 = v227
	goto L41
L58:
	;
	v238 = F_ExecGetResultSlotOps(m, v220, v185+int32(15))
	mBase = m.M
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v220)+56))
	v251 = v239
	v252 = v238
	goto L41
L59:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v187)+80))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v187)+76))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+100)))
	if v246 != int32(1) {
		v251 = v245
		v252 = v244
		goto L41
	} else {
		goto L60
	}
L60:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+15)) = uint8(v249)
	v251 = v245
	v252 = v244
	goto L41
L61:
	;
	if v252 != 0 {
		goto L39
	} else {
		goto L62
	}
L62:
	;
	goto L40
L63:
	;
	v277 = int32(0)
	goto L37
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v281 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l1
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)) = uint8(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
	v331 = v22 + int32(8)
	v335 = m.G0
	v337 = v335 - int32(16)
	m.G0 = v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)) = uint8(v325)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v331)+24))
	if v342 != 0 {
		goto L83
	} else {
		goto L84
	}
L67:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v304 + int32(1)
	v310 = v303 + v304*int32(40)
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+32)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+24)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+16)) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v317
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v310))) = v319
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v301
	v303 = v301
	goto L67
L69:
	;
	v284 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v284
	v288 = F_palloc_mul(m, int32(40), v284)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v290 != v281 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v301 = v288
	goto L68
L73:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v303 = v292
	goto L67
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v281 << (uint(int32(1)) % 32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v299 = F_repalloc(m, v296, v281*int32(80))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v301 = v299
	goto L68
L77:
	;
	if v429 != 0 {
		goto L105
	} else {
		goto L106
	}
L78:
	;
	m.G0 = v337 + int32(16)
	goto L77
L79:
	;
	v429 = int32(1)
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+28)) = v404
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+20)) = uint8(v418)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+24)) = v403
	if v404 != int32(_a_F_ExecBuildGroupingEqual_0) {
		goto L79
	} else {
		goto L104
	}
L81:
	;
	v413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v331)+20)) = uint8(v413)
	*(*int64)(unsafe.Add(mBase, uint32(v331)+24)) = int64(0)
	goto L79
L82:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)))
	if base.B2i32(v403 == int32(0))|base.B2i32(v407 != int32(1)) != 0 {
		goto L81
	} else {
		goto L102
	}
L83:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v331)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)) = uint8(base.B2i32(v343 != int32(0)))
	v403 = v342
	v404 = v343
	goto L82
L84:
	;
	goto L85
L85:
	;
	if v339 == int32(0) {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	switch v349 - int32(2) {
	case 0:
		goto L89
	case 1:
		goto L88
	default:
		goto L87
	}
L87:
	;
	if base.Ui32(int32(2)) < base.Ui32(v349-int32(4)) {
		goto L81
	} else {
		goto L100
	}
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v339)+36))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+101)))
	if v373 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v339)+40))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+102)))
	if v353 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v352 == int32(0) {
		goto L81
	} else {
		goto L94
	}
L91:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+98)))
	if v356 != int32(1) {
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v339)+88))
	if v359 == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)) = uint8(v362)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v352)+56))
	v403 = v364
	v404 = v359
	goto L82
L94:
	;
	v370 = F_ExecGetResultSlotOps(m, v352, v337+int32(15))
	mBase = m.M
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v352)+56))
	v403 = v371
	v404 = v370
	goto L82
L95:
	;
	if v372 == int32(0) {
		goto L81
	} else {
		goto L99
	}
L96:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+97)))
	if v376 != int32(1) {
		goto L81
	} else {
		goto L97
	}
L97:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v339)+84))
	if v379 == int32(0) {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v382 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)) = uint8(v382)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v372)+56))
	v403 = v384
	v404 = v379
	goto L82
L99:
	;
	v390 = F_ExecGetResultSlotOps(m, v372, v337+int32(15))
	mBase = m.M
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v372)+56))
	v403 = v391
	v404 = v390
	goto L82
L100:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v339)+80))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v339)+76))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+100)))
	if v398 != int32(1) {
		v403 = v397
		v404 = v396
		goto L82
	} else {
		goto L101
	}
L101:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v337)+15)) = uint8(v401)
	v403 = v397
	v404 = v396
	goto L82
L102:
	;
	if v404 != 0 {
		goto L80
	} else {
		goto L103
	}
L103:
	;
	goto L81
L104:
	;
	v429 = int32(0)
	goto L78
L105:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v433 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	goto L107
L107:
	;
	v476 = l4 - int32(1)
	if v476 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L108:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v456 + int32(1)
	v462 = v455 + v456*int32(40)
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+32)) = v463
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+24)) = v465
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+16)) = v467
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v462)+8)) = v469
	v471 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v462))) = v471
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v453
	v455 = v453
	goto L108
L110:
	;
	v436 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v436
	v440 = F_palloc_mul(m, int32(40), v436)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v442 != v433 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v453 = v440
	goto L109
L114:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v455 = v444
	goto L108
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v433 << (uint(int32(1)) % 32)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v451 = F_repalloc(m, v448, v433*int32(80))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v453 = v451
	goto L109
L118:
	;
	v826 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v826
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(0)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v830 == v826 {
		goto L184
	} else {
		goto L185
	}
L119:
	;
	v484 = int32(0)
	v488 = v476
	goto L120
L120:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v500 = int32(3)
	v502 = int32(1)
	v505 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v488<<(uint(v502)%32)))))
	v507 = v505 - v502
	v509 = v507 * int32(100)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v515 = v488 << (uint(int32(2)) % 32)
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l7+v515)))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l6+v515)))
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildGroupingEqual[0]))
	v524 = F_object_aclcheck(m, int32(1255), v520, v522, int64(128))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	if v763 == int32(0) {
		goto L118
	} else {
		goto L177
	}
L122:
	;
	if v524 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v527 = F_get_func_name(m, v520)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildGroupingEqual[1]))
	if v532 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	F_aclcheck_error(m, v524, int32(19), v527)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	F_RunFunctionExecuteHook(m, v520)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v539 = F_palloc0(m, int32(28))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	v542 = F_palloc0(m, int32(56))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_fmgr_info(m, v520, v539)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v546 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v539)+24)) = v546
	v548 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+18)) = uint16(v548)
	*(*uint8)(unsafe.Add(mBase, uint32(v542)+16)) = uint8(v546)
	*(*int32)(unsafe.Add(mBase, uint32(v542)+12)) = v517
	*(*int64)(unsafe.Add(mBase, uint32(v542)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(7)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v509+(l0+v510<<(uint(v500)%32)))+96))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v542 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v542 + int32(24)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v569 == v546 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v592 + int32(1)
	v596 = int32(40)
	v598 = v591 + v592*v596
	v599 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v598)+32)) = v599
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v598)+24)) = v601
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v598)+16)) = v603
	v605 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v598)+8)) = v605
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v598))) = v607
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(8)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1+v499<<(uint(v500)%32)+v509)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v612
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v542 + v596
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v542 + int32(48)
	v620 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v622 == v620 {
		goto L147
	} else {
		goto L148
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v589
	v591 = v589
	goto L135
L137:
	;
	v572 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v572
	v576 = F_palloc_mul(m, int32(40), v572)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v578 != v569 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v589 = v576
	goto L136
L141:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v591 = v580
	goto L135
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v569 << (uint(int32(1)) % 32)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v587 = F_repalloc(m, v584, v569*int32(80))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v589 = v587
	goto L136
L145:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v645 + int32(1)
	v651 = v644 + v645*int32(40)
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+32)) = v652
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+24)) = v654
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+16)) = v656
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+8)) = v658
	v660 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v651))) = v660
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(62)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v539
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(2)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v672 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v642
	v644 = v642
	goto L145
L147:
	;
	v625 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v625
	v629 = F_palloc_mul(m, int32(40), v625)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v631 != v622 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v642 = v629
	goto L146
L151:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v644 = v633
	goto L145
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v622 << (uint(int32(1)) % 32)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v640 = F_repalloc(m, v637, v622*int32(80))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v642 = v640
	goto L146
L155:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v695 + int32(1)
	v701 = v694 + v695*int32(40)
	v702 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v701)+32)) = v702
	v704 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v701)+24)) = v704
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v701)+16)) = v706
	v708 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v701)+8)) = v708
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v701))) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v51
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	if v718 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v692
	v694 = v692
	goto L155
L157:
	;
	v675 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v675
	v679 = F_palloc_mul(m, int32(40), v675)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v681 != v672 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v692 = v679
	goto L156
L161:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v694 = v683
	goto L155
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v672 << (uint(int32(1)) % 32)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v690 = F_repalloc(m, v687, v672*int32(80))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v692 = v690
	goto L156
L165:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v742 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v741 + v742
	v747 = v740 + v741*int32(40)
	v748 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v747)+32)) = v748
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v747)+24)) = v750
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v747)+16)) = v752
	v754 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v747)+8)) = v754
	v756 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v747))) = v756
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v763 = F_lappend_int(m, v484, v760-v742)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L175
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v738
	v740 = v738
	goto L165
L167:
	;
	v721 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v721
	v725 = F_palloc_mul(m, int32(40), v721)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v727 != v718 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v738 = v725
	goto L166
L171:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v740 = v729
	goto L165
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v718 << (uint(int32(1)) % 32)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v736 = F_repalloc(m, v733, v718*int32(80))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v738 = v736
	goto L166
L175:
	;
	if int32(0) < v488 {
		v484 = v763
		v488 = v488 - v742
		goto L120
	} else {
		goto L176
	}
L176:
	;
	goto L121
L177:
	;
	v769 = int32(0)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	if v770 <= v769 {
		goto L118
	} else {
		goto L178
	}
L178:
	;
	v781 = v769
	goto L179
L179:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v763)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793+v781<<(uint(int32(2))%32))))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v792+v797*int32(40))+16)) = v801
	v804 = v781 + int32(1)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	if v804 < v805 {
		v781 = v804
		goto L179
	} else {
		goto L181
	}
L180:
	;
	goto L118
L181:
	;
	goto L180
L182:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v853 + int32(1)
	v859 = v852 + v853*int32(40)
	v860 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v859)+32)) = v860
	v862 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v859)+24)) = v862
	v864 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v859)+16)) = v864
	v866 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v859)+8)) = v866
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v859))) = v868
	v870 = F_jit_compile_expr(m, v25)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L192
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v850
	v852 = v850
	goto L182
L184:
	;
	v833 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v833
	v837 = F_palloc_mul(m, int32(40), v833)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v839 != v830 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v850 = v837
	goto L183
L188:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v852 = v841
	goto L182
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v830 << (uint(int32(1)) % 32)
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v848 = F_repalloc(m, v845, v830*int32(80))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v850 = v848
	goto L183
L192:
	;
	if v870 != 0 {
		v884 = v25
		goto L3
	} else {
		goto L193
	}
L193:
	;
	F_ExecReadyInterpretedExpr(m, v25)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v884 = v25
	goto L3
}
func F_flatten_grouping_sets(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	v4 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		v173 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v173
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(1) {
		v129 = l0
		v130 = l1
		v131 = l2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v137 <= v136 {
		v173 = v136
		goto L3
	} else {
		goto L59
	}
L6:
	;
	if v14 != int32(107) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if v73 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L8:
	;
	if v14 != int32(36) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L11:
	;
	return l0
L12:
	;
	goto L13
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 != int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return l0
L15:
	;
	goto L16
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v26 == int32(0) {
		v173 = v4
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v32 = v26
	goto L19
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v39 != int32(36) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v173 = v4
	goto L3
L21:
	;
	if v39 == int32(107) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v48 != int32(2) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v66 = v32
	v70 = int32(1)
	goto L7
L25:
	;
	goto L26
L26:
	;
	if v39 != int32(1) {
		v173 = v32
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v129 = v32
	v130 = int32(0)
	v131 = int32(0)
	goto L5
L28:
	;
	return v32
L29:
	;
	goto L30
L30:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v52 != 0 {
		v32 = v52
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 == v62 {
		v173 = v62
		goto L3
	} else {
		goto L39
	}
L34:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	if l1 != 0 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if l1 == int32(0) {
		v66 = l0
		v70 = int32(1)
		goto L7
	} else {
		goto L38
	}
L37:
	;
	v66 = l0
	v70 = v55
	goto L7
L38:
	;
	goto L33
L39:
	;
	v66 = l0
	v70 = v62
	goto L7
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if base.B2i32(v120 == int32(4))&v70 != 0 {
		goto L55
	} else {
		goto L56
	}
L41:
	;
	v115 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v78 <= v77 {
		v115 = v77
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v84 = v77
	v85 = int32(0)
	goto L45
L45:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v85<<(uint(int32(2))%32))))
	v94 = int32(0)
	v96 = F_flatten_grouping_sets(m, v93, v94, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v115 = v108
	goto L40
L47:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v98 != int32(107) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v110 = v85 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v110 < v111 {
		v84 = v108
		v85 = v110
		goto L45
	} else {
		goto L54
	}
L49:
	;
	v106 = F_lappend(m, v84, v96)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v101 != int32(4) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v104 = F_list_concat(m, v84, v96)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v108 = v104
	goto L48
L53:
	;
	v108 = v106
	goto L48
L54:
	;
	goto L46
L55:
	;
	return v115
L56:
	;
	goto L57
L57:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v126 = F_makeGroupingSet(m, v120, v115, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	return v126
L59:
	;
	v144 = int32(0)
	v145 = v136
	goto L60
L60:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v144<<(uint(int32(2))%32))))
	v153 = F_flatten_grouping_sets(m, v152, v130, v131)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v173 = v164
	goto L3
L62:
	;
	v166 = v144 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v166 < v167 {
		v144 = v166
		v145 = v164
		goto L60
	} else {
		goto L70
	}
L63:
	;
	if v153 == int32(0) {
		v164 = v145
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v157 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v160 = F_list_concat(m, v145, v153)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v162 = F_lappend(m, v145, v153)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v164 = v160
	goto L62
L69:
	;
	v164 = v162
	goto L62
L70:
	;
	goto L61
}
