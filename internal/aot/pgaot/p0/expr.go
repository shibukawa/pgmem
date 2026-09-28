package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecPushExprSetupSteps(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v480 int64
	_ = v480
	var v482 int64
	_ = v482
	var v484 int64
	_ = v484
	var v486 int64
	_ = v486
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int64
	_ = v635
	var v637 int64
	_ = v637
	var v639 int64
	_ = v639
	var v641 int64
	_ = v641
	var v643 int64
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int64
	_ = v792
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v798 int64
	_ = v798
	var v800 int64
	_ = v800
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v176 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(2)
	v30 = v9 + int32(8)
	v34 = m.G0
	v36 = v34 - int32(16)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v24)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	if v128 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	m.G0 = v36 + int32(16)
	goto L3
L5:
	;
	v128 = int32(1)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v103
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v102
	if v103 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L5
	} else {
		goto L30
	}
L7:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v112)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
	goto L5
L8:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)))
	if base.B2i32(v102 == int32(0))|base.B2i32(v106 != int32(1)) != 0 {
		goto L7
	} else {
		goto L28
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(base.B2i32(v42 != int32(0)))
	v102 = v41
	v103 = v42
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v48 - int32(2) {
	case 0:
		goto L15
	case 1:
		goto L14
	default:
		goto L13
	}
L13:
	;
	if base.Ui32(int32(2)) < base.Ui32(v48-int32(4)) {
		goto L7
	} else {
		goto L26
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+101)))
	if v72 != int32(1) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+102)))
	if v52 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v51 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+98)))
	if v55 != int32(1) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+88))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v61 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v102 = v63
	v103 = v58
	goto L8
L20:
	;
	v69 = F_ExecGetResultSlotOps(m, v51, v36+int32(15))
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v102 = v70
	v103 = v69
	goto L8
L21:
	;
	if v71 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+97)))
	if v75 != int32(1) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
	if v78 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v102 = v83
	v103 = v78
	goto L8
L25:
	;
	v89 = F_ExecGetResultSlotOps(m, v71, v36+int32(15))
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v102 = v90
	v103 = v89
	goto L8
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+100)))
	if v97 != int32(1) {
		v102 = v96
		v103 = v95
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v100)
	v102 = v96
	v103 = v95
	goto L8
L28:
	;
	if v103 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	v128 = int32(0)
	goto L4
L31:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v134 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v157 + int32(1)
	v163 = v156 + v157*int32(40)
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+32)) = v164
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+24)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+16)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v172
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v154
	v156 = v154
	goto L32
L34:
	;
	v137 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v137
	v141 = F_palloc_mul(m, int32(40), v137)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v143 != v134 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	return
L38:
	;
	v154 = v141
	goto L33
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v156 = v145
	goto L32
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v134 << (uint(int32(1)) % 32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v152 = F_repalloc(m, v149, v134*int32(80))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v154 = v152
	goto L33
L43:
	;
	v333 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v333 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L44:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v179)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v187 = v9 + int32(8)
	v191 = m.G0
	v193 = v191 - int32(16)
	m.G0 = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)) = uint8(v179)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v187)+24))
	if v198 != 0 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v285 == int32(0) {
		goto L43
	} else {
		goto L73
	}
L46:
	;
	m.G0 = v193 + int32(16)
	goto L45
L47:
	;
	v285 = int32(1)
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+28)) = v260
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+20)) = uint8(v274)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+24)) = v259
	if v260 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L47
	} else {
		goto L72
	}
L49:
	;
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+20)) = uint8(v269)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+24)) = int64(0)
	goto L47
L50:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)))
	if base.B2i32(v259 == int32(0))|base.B2i32(v263 != int32(1)) != 0 {
		goto L49
	} else {
		goto L70
	}
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v187)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)) = uint8(base.B2i32(v199 != int32(0)))
	v259 = v198
	v260 = v199
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v195 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	switch v205 - int32(2) {
	case 0:
		goto L57
	case 1:
		goto L56
	default:
		goto L55
	}
L55:
	;
	if base.Ui32(int32(2)) < base.Ui32(v205-int32(4)) {
		goto L49
	} else {
		goto L68
	}
L56:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v195)+36))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+101)))
	if v229 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v195)+40))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+102)))
	if v209 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v208 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L59:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+98)))
	if v212 != int32(1) {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v195)+88))
	if v215 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)) = uint8(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v208)+56))
	v259 = v220
	v260 = v215
	goto L50
L62:
	;
	v226 = F_ExecGetResultSlotOps(m, v208, v193+int32(15))
	mBase = m.M
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v208)+56))
	v259 = v227
	v260 = v226
	goto L50
L63:
	;
	if v228 == int32(0) {
		goto L49
	} else {
		goto L67
	}
L64:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+97)))
	if v232 != int32(1) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v195)+84))
	if v235 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)) = uint8(v238)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v228)+56))
	v259 = v240
	v260 = v235
	goto L50
L67:
	;
	v246 = F_ExecGetResultSlotOps(m, v228, v193+int32(15))
	mBase = m.M
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v228)+56))
	v259 = v247
	v260 = v246
	goto L50
L68:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v195)+80))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v195)+76))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+100)))
	if v254 != int32(1) {
		v259 = v253
		v260 = v252
		goto L50
	} else {
		goto L69
	}
L69:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+15)) = uint8(v257)
	v259 = v253
	v260 = v252
	goto L50
L70:
	;
	if v260 != 0 {
		goto L48
	} else {
		goto L71
	}
L71:
	;
	goto L49
L72:
	;
	v285 = int32(0)
	goto L46
L73:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v291 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v314 + int32(1)
	v320 = v313 + v314*int32(40)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+32)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+24)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v329
	goto L43
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v311
	v313 = v311
	goto L74
L76:
	;
	v294 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v294
	v298 = F_palloc_mul(m, int32(40), v294)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L37
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v300 != v291 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v311 = v298
	goto L75
L80:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v313 = v302
	goto L74
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v291 << (uint(int32(1)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v309 = F_repalloc(m, v306, v291*int32(80))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	v311 = v309
	goto L75
L84:
	;
	v490 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v490 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L85:
	;
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v344 = v9 + int32(8)
	v348 = m.G0
	v350 = v348 - int32(16)
	m.G0 = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)) = uint8(v336)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v344)+24))
	if v355 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if v442 == int32(0) {
		goto L84
	} else {
		goto L114
	}
L87:
	;
	m.G0 = v350 + int32(16)
	goto L86
L88:
	;
	v442 = int32(1)
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+28)) = v417
	v431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+20)) = uint8(v431)
	*(*int32)(unsafe.Add(mBase, uint32(v344)+24)) = v416
	if v417 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L88
	} else {
		goto L113
	}
L90:
	;
	v426 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v344)+20)) = uint8(v426)
	*(*int64)(unsafe.Add(mBase, uint32(v344)+24)) = int64(0)
	goto L88
L91:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)))
	if base.B2i32(v416 == int32(0))|base.B2i32(v420 != int32(1)) != 0 {
		goto L90
	} else {
		goto L111
	}
L92:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v344)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)) = uint8(base.B2i32(v356 != int32(0)))
	v416 = v355
	v417 = v356
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v352 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	switch v362 - int32(2) {
	case 0:
		goto L98
	case 1:
		goto L97
	default:
		goto L96
	}
L96:
	;
	if base.Ui32(int32(2)) < base.Ui32(v362-int32(4)) {
		goto L90
	} else {
		goto L109
	}
L97:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v352)+36))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+101)))
	if v386 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v352)+40))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+102)))
	if v366 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v365 == int32(0) {
		goto L90
	} else {
		goto L103
	}
L100:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+98)))
	if v369 != int32(1) {
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v352)+88))
	if v372 == int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v375 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v365)+56))
	v416 = v377
	v417 = v372
	goto L91
L103:
	;
	v383 = F_ExecGetResultSlotOps(m, v365, v350+int32(15))
	mBase = m.M
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v365)+56))
	v416 = v384
	v417 = v383
	goto L91
L104:
	;
	if v385 == int32(0) {
		goto L90
	} else {
		goto L108
	}
L105:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+97)))
	if v389 != int32(1) {
		goto L90
	} else {
		goto L106
	}
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v352)+84))
	if v392 == int32(0) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)) = uint8(v395)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v385)+56))
	v416 = v397
	v417 = v392
	goto L91
L108:
	;
	v403 = F_ExecGetResultSlotOps(m, v385, v350+int32(15))
	mBase = m.M
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v385)+56))
	v416 = v404
	v417 = v403
	goto L91
L109:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v352)+80))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v352)+76))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+100)))
	if v411 != int32(1) {
		v416 = v410
		v417 = v409
		goto L91
	} else {
		goto L110
	}
L110:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v350)+15)) = uint8(v414)
	v416 = v410
	v417 = v409
	goto L91
L111:
	;
	if v417 != 0 {
		goto L89
	} else {
		goto L112
	}
L112:
	;
	goto L90
L113:
	;
	v442 = int32(0)
	goto L87
L114:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v448 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v471 + int32(1)
	v477 = v470 + v471*int32(40)
	v478 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v477)+32)) = v478
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v477)+24)) = v480
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v477)+16)) = v482
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v477)+8)) = v484
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v477))) = v486
	goto L84
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v468
	v470 = v468
	goto L115
L117:
	;
	v451 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v451
	v455 = F_palloc_mul(m, int32(40), v451)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L37
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v457 != v448 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v468 = v455
	goto L116
L121:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v470 = v459
	goto L115
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v448 << (uint(int32(1)) % 32)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v466 = F_repalloc(m, v463, v448*int32(80))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L37
	} else {
		goto L124
	}
L124:
	;
	v468 = v466
	goto L116
L125:
	;
	v647 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v647 <= int32(0) {
		goto L166
	} else {
		goto L167
	}
L126:
	;
	v493 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v493)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(5)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v501 = v9 + int32(8)
	v505 = m.G0
	v507 = v505 - int32(16)
	m.G0 = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)) = uint8(v493)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v501)+24))
	if v512 != 0 {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	if v599 == int32(0) {
		goto L125
	} else {
		goto L155
	}
L128:
	;
	m.G0 = v507 + int32(16)
	goto L127
L129:
	;
	v599 = int32(1)
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501)+28)) = v574
	v588 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v501)+20)) = uint8(v588)
	*(*int32)(unsafe.Add(mBase, uint32(v501)+24)) = v573
	if v574 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L129
	} else {
		goto L154
	}
L131:
	;
	v583 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v501)+20)) = uint8(v583)
	*(*int64)(unsafe.Add(mBase, uint32(v501)+24)) = int64(0)
	goto L129
L132:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)))
	if base.B2i32(v573 == int32(0))|base.B2i32(v577 != int32(1)) != 0 {
		goto L131
	} else {
		goto L152
	}
L133:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v501)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)) = uint8(base.B2i32(v513 != int32(0)))
	v573 = v512
	v574 = v513
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v509 == int32(0) {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	switch v519 - int32(2) {
	case 0:
		goto L139
	case 1:
		goto L138
	default:
		goto L137
	}
L137:
	;
	if base.Ui32(int32(2)) < base.Ui32(v519-int32(4)) {
		goto L131
	} else {
		goto L150
	}
L138:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+101)))
	if v543 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v509)+40))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+102)))
	if v523 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v522 == int32(0) {
		goto L131
	} else {
		goto L144
	}
L141:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+98)))
	if v526 != int32(1) {
		goto L131
	} else {
		goto L142
	}
L142:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v509)+88))
	if v529 == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)) = uint8(v532)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v522)+56))
	v573 = v534
	v574 = v529
	goto L132
L144:
	;
	v540 = F_ExecGetResultSlotOps(m, v522, v507+int32(15))
	mBase = m.M
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v522)+56))
	v573 = v541
	v574 = v540
	goto L132
L145:
	;
	if v542 == int32(0) {
		goto L131
	} else {
		goto L149
	}
L146:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+97)))
	if v546 != int32(1) {
		goto L131
	} else {
		goto L147
	}
L147:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)+84))
	if v549 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v552 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)) = uint8(v552)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v542)+56))
	v573 = v554
	v574 = v549
	goto L132
L149:
	;
	v560 = F_ExecGetResultSlotOps(m, v542, v507+int32(15))
	mBase = m.M
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v542)+56))
	v573 = v561
	v574 = v560
	goto L132
L150:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v509)+80))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v509)+76))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+100)))
	if v568 != int32(1) {
		v573 = v567
		v574 = v566
		goto L132
	} else {
		goto L151
	}
L151:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v507)+15)) = uint8(v571)
	v573 = v567
	v574 = v566
	goto L132
L152:
	;
	if v574 != 0 {
		goto L130
	} else {
		goto L153
	}
L153:
	;
	goto L131
L154:
	;
	v599 = int32(0)
	goto L128
L155:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v605 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v628 + int32(1)
	v634 = v627 + v628*int32(40)
	v635 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v634)+32)) = v635
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v634)+24)) = v637
	v639 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v634)+16)) = v639
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v634)+8)) = v641
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v634))) = v643
	goto L125
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v625
	v627 = v625
	goto L156
L158:
	;
	v608 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v608
	v612 = F_palloc_mul(m, int32(40), v608)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L37
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v614 != v605 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v625 = v612
	goto L157
L162:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v627 = v616
	goto L156
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v605 << (uint(int32(1)) % 32)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v623 = F_repalloc(m, v620, v605*int32(80))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L37
	} else {
		goto L165
	}
L165:
	;
	v625 = v623
	goto L157
L166:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v804 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L167:
	;
	v650 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v650)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v647
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v658 = v9 + int32(8)
	v662 = m.G0
	v664 = v662 - int32(16)
	m.G0 = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)) = uint8(v650)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v658)+24))
	if v669 != 0 {
		goto L174
	} else {
		goto L175
	}
L168:
	;
	if v756 == int32(0) {
		goto L166
	} else {
		goto L196
	}
L169:
	;
	m.G0 = v664 + int32(16)
	goto L168
L170:
	;
	v756 = int32(1)
	goto L169
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v658)+28)) = v731
	v745 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+20)) = uint8(v745)
	*(*int32)(unsafe.Add(mBase, uint32(v658)+24)) = v730
	if v731 != int32(_a_F_ExecPushExprSetupSteps_0) {
		goto L170
	} else {
		goto L195
	}
L172:
	;
	v740 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+20)) = uint8(v740)
	*(*int64)(unsafe.Add(mBase, uint32(v658)+24)) = int64(0)
	goto L170
L173:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)))
	if base.B2i32(v730 == int32(0))|base.B2i32(v734 != int32(1)) != 0 {
		goto L172
	} else {
		goto L193
	}
L174:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v658)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)) = uint8(base.B2i32(v670 != int32(0)))
	v730 = v669
	v731 = v670
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v666 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	switch v676 - int32(2) {
	case 0:
		goto L180
	case 1:
		goto L179
	default:
		goto L178
	}
L178:
	;
	if base.Ui32(int32(2)) < base.Ui32(v676-int32(4)) {
		goto L172
	} else {
		goto L191
	}
L179:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v666)+36))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+101)))
	if v700 != int32(1) {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v666)+40))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+102)))
	if v680 != int32(1) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v679 == int32(0) {
		goto L172
	} else {
		goto L185
	}
L182:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+98)))
	if v683 != int32(1) {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v666)+88))
	if v686 == int32(0) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	v689 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)) = uint8(v689)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v679)+56))
	v730 = v691
	v731 = v686
	goto L173
L185:
	;
	v697 = F_ExecGetResultSlotOps(m, v679, v664+int32(15))
	mBase = m.M
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v679)+56))
	v730 = v698
	v731 = v697
	goto L173
L186:
	;
	if v699 == int32(0) {
		goto L172
	} else {
		goto L190
	}
L187:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+97)))
	if v703 != int32(1) {
		goto L172
	} else {
		goto L188
	}
L188:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v666)+84))
	if v706 == int32(0) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v709 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)) = uint8(v709)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v699)+56))
	v730 = v711
	v731 = v706
	goto L173
L190:
	;
	v717 = F_ExecGetResultSlotOps(m, v699, v664+int32(15))
	mBase = m.M
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v699)+56))
	v730 = v718
	v731 = v717
	goto L173
L191:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v666)+80))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v666)+76))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+100)))
	if v725 != int32(1) {
		v730 = v724
		v731 = v723
		goto L173
	} else {
		goto L192
	}
L192:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v664)+15)) = uint8(v728)
	v730 = v724
	v731 = v723
	goto L173
L193:
	;
	if v731 != 0 {
		goto L171
	} else {
		goto L194
	}
L194:
	;
	goto L172
L195:
	;
	v756 = int32(0)
	goto L169
L196:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v762 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v785 + int32(1)
	v791 = v784 + v785*int32(40)
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+32)) = v792
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+24)) = v794
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+16)) = v796
	v798 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v791)+8)) = v798
	v800 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v791))) = v800
	goto L166
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v782
	v784 = v782
	goto L197
L199:
	;
	v765 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v765
	v769 = F_palloc_mul(m, int32(40), v765)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L37
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v771 != v762 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v782 = v769
	goto L198
L203:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v784 = v773
	goto L197
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v762 << (uint(int32(1)) % 32)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v780 = F_repalloc(m, v777, v762*int32(80))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L37
	} else {
		goto L206
	}
L206:
	;
	v782 = v780
	goto L198
L207:
	;
	m.G0 = v9 + int32(48)
	return
L208:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v807 <= int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v816 = int32(0)
	goto L210
L210:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v804)+12))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v821+v816<<(uint(int32(2))%32))))
	F_ExecInitSubPlanExpr(m, v825, l0, l0+int32(8), l0+int32(5))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L37
	} else {
		goto L212
	}
L211:
	;
	goto L207
L212:
	;
	v829 = v816 + int32(1)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v829 < v830 {
		v816 = v829
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
}
func F_assign_expr_collations(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	v15 = F_assign_collations_walker(m, l1, v6+int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(32)
		return
	}
}
func F_checkExprIsVarFree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_contain_vars_of_level(m, l1, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_checkExprIsVarFree_0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
					F_errmsg(m, int32(_a_F_checkExprIsVarFree_1), v7)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v24 = F_locate_var_of_level(m, l1, int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_parser_errposition(m, l0, v24)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_checkExprIsVarFree_2), int32(1937), int32(_a_F_checkExprIsVarFree_3))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
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
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_exprInputCollation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v2 = int32(0)
	if l0 == v2 {
		v30 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = v7 - int32(9)
		if base.Ui32(int32(30)) < base.Ui32(v9) {
			v30 = v2
		} else {
			v13 = int32(1) << (uint(v9) % 32)
			if v13&int32(3904) == int32(0) {
				if v13&int32(5) != 0 {
					v25 = int32(16)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
					v30 = v27
				} else {
					if v9 != int32(30) {
						v30 = v2
					} else {
						v25 = int32(12)
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
						v30 = v27
					}
				}
			} else {
				v25 = int32(24)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0)))
				v30 = v27
			}
		}
	}
	return v30
}
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 float64
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v191 float64
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	v3 = int32(0)
	if l0 == v3 {
		v219 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v229 = F_copyObjectImpl(m, v109)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L10
	} else {
		goto L53
	}
L2:
	;
	return v219
L3:
	;
	v15 = l0
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v27 - int32(6) {
	case 0:
		goto L9
	default:
		v125 = v27
		goto L6
	case 2:
		goto L8
	case 3:
		goto L7
	}
L5:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_fix_expr_common(m, v210, v15)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L51
	}
L6:
	;
	if v125 != int32(24) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+296))
	if v65 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v61 = F_fix_param_node(m, v60, v15)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L16
	}
L9:
	;
	v31 = F_palloc(m, int32(48))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v35
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(0) <= v47 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v50 + v47
	goto L14
L13:
	;
	goto L14
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v53 == int32(0) {
		v219 = v31
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v56 + v53
	return v31
L16:
	;
	return v61
L17:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v125 = v122
	goto L6
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if v68 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 != int32(1) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v74 <= int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v82 = int32(0)
	v85 = v74
	goto L22
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v82<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v97 == v98 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	if v109 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v102 = F_equal(m, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	v105 = v85
	goto L27
L27:
	;
	v107 = v82 + int32(1)
	if v107 < v105 {
		v82 = v107
		v85 = v105
		goto L22
	} else {
		goto L30
	}
L28:
	;
	if v102 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v105 = v104
	goto L27
L30:
	;
	goto L17
L31:
	;
	goto L17
L32:
	;
	goto L5
L33:
	;
	if v125 != int32(321) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v150 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v152 = int32(0)
	v155 = v152
	v159 = v152
	v164 = float64(0)
	goto L42
L36:
	;
	if v125 != int32(58) {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v148 != 0 {
		v15 = v148
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v141 = F_copyObjectImpl(m, v15)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v143 + v144
	return v141
L41:
	;
	v219 = v3
	goto L2
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v159<<(uint(int32(2))%32))))
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v171)+64))
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v171)+56))
	v175 = base.F64_add(base.F64_mul(v150, v172), v174)
	if v155 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v151)+380))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v204-v206))) = uint8(v206)
	if v188 != 0 {
		v15 = v188
		goto L4
	} else {
		goto L50
	}
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v151)+376))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192+v193-v195))) = uint8(v195)
	v200 = v159 + v195
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v200 < v201 {
		v155 = v188
		v159 = v200
		v164 = v191
		goto L42
	} else {
		goto L49
	}
L45:
	;
	v188 = v171
	v191 = v175
	goto L44
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)+52))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155)+52))
	if v178 < v179 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.B2i32(base.F64_ge(v164, v175) == int32(0))|base.B2i32(v178 != v179) != 0 {
		v188 = v155
		v191 = v164
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L43
L50:
	;
	v219 = v3
	goto L2
L51:
	;
	v214 = F_expression_tree_mutator_impl(m, v15, int32(887), l1)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v219 = v214
	goto L2
L53:
	;
	return v229
}
func F_format_expr_params(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v350 int32
	_ = v350
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(_a_F_format_expr_params_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v19
	v22 = v12 + int32(56)
	F_initStringInfo(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v350 = int32(0)
	goto L3
L3:
	;
	m.G0 = v12 + int32(80)
	return v350
L4:
	;
	return int32(0)
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v16
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v350 = v339
	goto L3
L7:
	;
	if v84 < int32(0) {
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v84 = base.I32_ctz(v70) | v71<<(uint(int32(5))%32)
	goto L7
L9:
	;
	v84 = int32(-2)
	goto L7
L10:
	;
	v35 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 <= v35 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v41 = v27 + int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v48 = v45 & int32(-1)
	if v48 != 0 {
		v70 = v48
		v71 = v35
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(1)
	if v49 == v38 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v53 = v49
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+v53<<(uint(int32(2))%32))))
	if v60 != 0 {
		v70 = v60
		v71 = v53
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v62 = v53 + int32(1)
	if v62 != v38 {
		v53 = v62
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v91, v12+int32(44), v12+int32(36), v12+int32(48), v12+int32(43))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_format_expr_params_1)
	F_appendStringInfo(m, v22, int32(_a_F_format_expr_params_2), v12+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+43)))
	if v111 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v144 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v116 = int32(_a_F_format_expr_params_0)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v120
	F_getTypeOutputInfo(m, v115, v12+int32(76), v12+int32(75))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_appendStringInfoString(m, v12+int32(56), int32(_a_F_format_expr_params_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L28
	}
L25:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v129 = F_OidOutputFunctionCall(m, v128, v114)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v117
	F_appendStringInfoStringQuoted(m, v22, v129, int32(-1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	goto L21
L29:
	;
	if v200 < int32(0) {
		goto L6
	} else {
		goto L40
	}
L30:
	;
	v200 = base.I32_ctz(v186) | v187<<(uint(int32(5))%32)
	goto L29
L31:
	;
	v200 = int32(-2)
	goto L29
L32:
	;
	v151 = v84 + int32(1)
	v153 = int32(base.Ui32(v151) >> (uint(int32(5)) % 32))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v154 <= v153 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v157 = v144 + int32(8)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v153<<(uint(int32(2))%32))))
	v164 = v161 & (int32(-1) << (uint(v151) % 32))
	if v164 != 0 {
		v186 = v164
		v187 = v153
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v166 = v153 + int32(1)
	if v166 == v154 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v169 = v166
	goto L36
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v157+v169<<(uint(int32(2))%32))))
	if v176 != 0 {
		v186 = v176
		v187 = v169
		goto L30
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	v178 = v169 + int32(1)
	if v178 != v154 {
		v169 = v178
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v207 = v200
	goto L41
L41:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212+v207<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v216, v12+int32(44), v12+int32(36), v12+int32(48), v12+int32(43))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L6
L43:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_format_expr_params_4)
	v232 = v12 + int32(56)
	F_appendStringInfo(m, v232, int32(_a_F_format_expr_params_2), v12)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+43)))
	if v236 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v269 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	F_appendStringInfoString(m, v232, int32(_a_F_format_expr_params_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v244 = int32(_a_F_format_expr_params_0)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v248
	F_getTypeOutputInfo(m, v243, v12+int32(76), v12+int32(75))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	v257 = F_OidOutputFunctionCall(m, v256, v242)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_format_expr_params[0])) = v245
	F_appendStringInfoStringQuoted(m, v12+int32(56), v257, int32(-1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	if int32(0) <= v325 {
		v207 = v325
		goto L41
	} else {
		goto L64
	}
L54:
	;
	v325 = base.I32_ctz(v311) | v312<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v325 = int32(-2)
	goto L53
L56:
	;
	v276 = v207 + int32(1)
	v278 = int32(base.Ui32(v276) >> (uint(int32(5)) % 32))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v279 <= v278 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v282 = v269 + int32(8)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282+v278<<(uint(int32(2))%32))))
	v289 = v286 & (int32(-1) << (uint(v276) % 32))
	if v289 != 0 {
		v311 = v289
		v312 = v278
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v291 = v278 + int32(1)
	if v291 == v279 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v294 = v291
	goto L60
L60:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v282+v294<<(uint(int32(2))%32))))
	if v301 != 0 {
		v311 = v301
		v312 = v294
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v303 = v294 + int32(1)
	if v303 != v279 {
		v294 = v303
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L42
}
