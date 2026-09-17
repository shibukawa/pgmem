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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
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
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v463 int64
	_ = v463
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v469 int64
	_ = v469
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int64
	_ = v596
	var v598 int64
	_ = v598
	var v600 int64
	_ = v600
	var v602 int64
	_ = v602
	var v604 int64
	_ = v604
	var v609 int32
	_ = v609
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int64
	_ = v648
	var v650 int64
	_ = v650
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int64
	_ = v656
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
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
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v699 int64
	_ = v699
	var v701 int64
	_ = v701
	var v703 int64
	_ = v703
	var v705 int64
	_ = v705
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	var v744 int64
	_ = v744
	var v746 int64
	_ = v746
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v853 int64
	_ = v853
	var v855 int64
	_ = v855
	var v857 int64
	_ = v857
	var v859 int64
	_ = v859
	var v861 int64
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v877 int32
	_ = v877
	v10 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v25 = F_palloc0(m, int32(68))
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
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(380)
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
	return v877
L4:
	;
	v877 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = l8
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)) = uint8(v43)
	v45 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v45
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
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
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
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
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
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)) = uint8(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3)
	v330 = v22 + int32(8)
	v334 = m.G0
	v336 = v334 - int32(16)
	m.G0 = v336
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v25)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)) = uint8(v324)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v330)+24))
	if v341 != 0 {
		goto L83
	} else {
		goto L84
	}
L67:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v303 + int32(1)
	v309 = v302 + v303*int32(40)
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+32)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+24)) = v312
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+16)) = v314
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v309))) = v318
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v300
	v302 = v300
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v287 = F_palloc(m, int32(640))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v289 != v281 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v300 = v287
	goto L68
L73:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v302 = v291
	goto L67
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v281 << (uint(int32(1)) % 32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v298 = F_repalloc(m, v295, v281*int32(80))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v300 = v298
	goto L68
L77:
	;
	if v428 != 0 {
		goto L105
	} else {
		goto L106
	}
L78:
	;
	m.G0 = v336 + int32(16)
	goto L77
L79:
	;
	v428 = int32(1)
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+28)) = v403
	v417 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v330)+20)) = uint8(v417)
	*(*int32)(unsafe.Add(mBase, uint32(v330)+24)) = v402
	if v403 != int32(_a_F_ExecBuildGroupingEqual_0) {
		goto L79
	} else {
		goto L104
	}
L81:
	;
	v412 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v330)+20)) = uint8(v412)
	*(*int64)(unsafe.Add(mBase, uint32(v330)+24)) = int64(0)
	goto L79
L82:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)))
	if base.B2i32(v402 == int32(0))|base.B2i32(v406 != int32(1)) != 0 {
		goto L81
	} else {
		goto L102
	}
L83:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v330)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)) = uint8(base.B2i32(v342 != int32(0)))
	v402 = v341
	v403 = v342
	goto L82
L84:
	;
	goto L85
L85:
	;
	if v338 == int32(0) {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	switch v348 - int32(2) {
	case 0:
		goto L89
	case 1:
		goto L88
	default:
		goto L87
	}
L87:
	;
	if base.Ui32(int32(2)) < base.Ui32(v348-int32(4)) {
		goto L81
	} else {
		goto L100
	}
L88:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v338)+36))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+101)))
	if v372 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v338)+40))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+102)))
	if v352 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v351 == int32(0) {
		goto L81
	} else {
		goto L94
	}
L91:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+98)))
	if v355 != int32(1) {
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v338)+88))
	if v358 == int32(0) {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v361 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)) = uint8(v361)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v351)+56))
	v402 = v363
	v403 = v358
	goto L82
L94:
	;
	v369 = F_ExecGetResultSlotOps(m, v351, v336+int32(15))
	mBase = m.M
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v351)+56))
	v402 = v370
	v403 = v369
	goto L82
L95:
	;
	if v371 == int32(0) {
		goto L81
	} else {
		goto L99
	}
L96:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+97)))
	if v375 != int32(1) {
		goto L81
	} else {
		goto L97
	}
L97:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v338)+84))
	if v378 == int32(0) {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v381 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)) = uint8(v381)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v371)+56))
	v402 = v383
	v403 = v378
	goto L82
L99:
	;
	v389 = F_ExecGetResultSlotOps(m, v371, v336+int32(15))
	mBase = m.M
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v371)+56))
	v402 = v390
	v403 = v389
	goto L82
L100:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v338)+80))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v338)+76))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+100)))
	if v397 != int32(1) {
		v402 = v396
		v403 = v395
		goto L82
	} else {
		goto L101
	}
L101:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v336)+15)) = uint8(v400)
	v402 = v396
	v403 = v395
	goto L82
L102:
	;
	if v403 != 0 {
		goto L80
	} else {
		goto L103
	}
L103:
	;
	goto L81
L104:
	;
	v428 = int32(0)
	goto L78
L105:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v432 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	goto L107
L107:
	;
	v474 = l4 - int32(1)
	if v474 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L108:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v454 + int32(1)
	v460 = v453 + v454*int32(40)
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v460)+32)) = v461
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v460)+24)) = v463
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v460)+16)) = v465
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v460)+8)) = v467
	v469 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v460))) = v469
	goto L107
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v451
	v453 = v451
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v438 = F_palloc(m, int32(640))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v440 != v432 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v451 = v438
	goto L109
L114:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v453 = v442
	goto L108
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v432 << (uint(int32(1)) % 32)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v449 = F_repalloc(m, v446, v432*int32(80))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v451 = v449
	goto L109
L118:
	;
	v820 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v820
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = int64(0)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v824 == v820 {
		goto L184
	} else {
		goto L185
	}
L119:
	;
	v482 = int32(0)
	v486 = v474
	goto L120
L120:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v498 = int32(4)
	v500 = int32(1)
	v503 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v486<<(uint(v500)%32)))))
	v505 = v503 - v500
	v507 = v505 * int32(100)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v513 = v486 << (uint(int32(2)) % 32)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l7+v513)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l6+v513)))
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildGroupingEqual[0]))
	v522 = F_object_aclcheck(m, int32(1255), v518, v520, int64(128))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	if v757 == int32(0) {
		goto L118
	} else {
		goto L177
	}
L122:
	;
	if v522 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v525 = F_get_func_name(m, v518)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_ExecBuildGroupingEqual[1]))
	if v530 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	F_aclcheck_error(m, v522, int32(19), v525)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	F_RunFunctionExecuteHook(m, v518)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v537 = F_palloc0(m, int32(28))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	v540 = F_palloc0(m, int32(36))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_fmgr_info(m, v518, v537)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v544 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+24)) = v544
	v546 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v540)+18)) = uint16(v546)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+16)) = uint8(v544)
	*(*int32)(unsafe.Add(mBase, uint32(v540)+12)) = v515
	*(*int64)(unsafe.Add(mBase, uint32(v540)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(7)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v507+(l0+v508<<(uint(v498)%32)))+88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v540 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v540 + int32(20)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v567 == v544 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v589 + int32(1)
	v595 = v588 + v589*int32(40)
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+32)) = v596
	v598 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+24)) = v598
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+16)) = v600
	v602 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v595)+8)) = v602
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v595))) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(8)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1+v497<<(uint(v498)%32)+v507)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v609
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v540 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v540 + int32(32)
	v617 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v617
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v619 == v617 {
		goto L147
	} else {
		goto L148
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v586
	v588 = v586
	goto L135
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v573 = F_palloc(m, int32(640))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v575 != v567 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v586 = v573
	goto L136
L141:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v588 = v577
	goto L135
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v567 << (uint(int32(1)) % 32)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v584 = F_repalloc(m, v581, v567*int32(80))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v586 = v584
	goto L136
L145:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v641 + int32(1)
	v647 = v640 + v641*int32(40)
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v647)+32)) = v648
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v647)+24)) = v650
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v647)+16)) = v652
	v654 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v647)+8)) = v654
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v647))) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(62)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v537
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v662
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(2)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v668 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v638
	v640 = v638
	goto L145
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v625 = F_palloc(m, int32(640))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v627 != v619 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v638 = v625
	goto L146
L151:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v640 = v629
	goto L145
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v619 << (uint(int32(1)) % 32)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v636 = F_repalloc(m, v633, v619*int32(80))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v638 = v636
	goto L146
L155:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v690 + int32(1)
	v696 = v689 + v690*int32(40)
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v696)+32)) = v697
	v699 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v696)+24)) = v699
	v701 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v696)+16)) = v701
	v703 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v696)+8)) = v703
	v705 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v696))) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(39)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v51
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	if v713 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v687
	v689 = v687
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v674 = F_palloc(m, int32(640))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v676 != v668 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v687 = v674
	goto L156
L161:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v689 = v678
	goto L155
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v668 << (uint(int32(1)) % 32)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v685 = F_repalloc(m, v682, v668*int32(80))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v687 = v685
	goto L156
L165:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v736 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v735 + v736
	v741 = v734 + v735*int32(40)
	v742 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v741)+32)) = v742
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v741)+24)) = v744
	v746 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v741)+16)) = v746
	v748 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v741)+8)) = v748
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v741))) = v750
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	v757 = F_lappend_int(m, v482, v754-v736)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L175
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v732
	v734 = v732
	goto L165
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v719 = F_palloc(m, int32(640))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v721 != v713 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v732 = v719
	goto L166
L171:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v734 = v723
	goto L165
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v713 << (uint(int32(1)) % 32)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v730 = F_repalloc(m, v727, v713*int32(80))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v732 = v730
	goto L166
L175:
	;
	if int32(0) < v486 {
		v482 = v757
		v486 = v486 - v736
		goto L120
	} else {
		goto L176
	}
L176:
	;
	goto L121
L177:
	;
	v763 = int32(0)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v757)+4))
	if v764 <= v763 {
		goto L118
	} else {
		goto L178
	}
L178:
	;
	v775 = v763
	goto L179
L179:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v757)+12))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v787+v775<<(uint(int32(2))%32))))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v786+v791*int32(40))+16)) = v795
	v798 = v775 + int32(1)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v757)+4))
	if v798 < v799 {
		v775 = v798
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
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v846 + int32(1)
	v852 = v845 + v846*int32(40)
	v853 = *(*int64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v852)+32)) = v853
	v855 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v852)+24)) = v855
	v857 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v852)+16)) = v857
	v859 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v852)+8)) = v859
	v861 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v852))) = v861
	v863 = F_jit_compile_expr(m, v25)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L192
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v843
	v845 = v843
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(16)
	v830 = F_palloc(m, int32(640))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	if v832 != v824 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v843 = v830
	goto L183
L188:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v845 = v834
	goto L182
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v824 << (uint(int32(1)) % 32)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v841 = F_repalloc(m, v838, v824*int32(80))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v843 = v841
	goto L183
L192:
	;
	if v863 != 0 {
		v877 = v25
		goto L3
	} else {
		goto L193
	}
L193:
	;
	F_ExecReadyInterpretedExpr(m, v25)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v877 = v25
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
