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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
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
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int64
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v464 int64
	_ = v464
	var v466 int64
	_ = v466
	var v468 int64
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v611 int64
	_ = v611
	var v613 int64
	_ = v613
	var v615 int64
	_ = v615
	var v617 int64
	_ = v617
	var v619 int64
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int64
	_ = v762
	var v764 int64
	_ = v764
	var v766 int64
	_ = v766
	var v768 int64
	_ = v768
	var v770 int64
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
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
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v170 <= int32(0) {
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v24)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	if v123 == int32(0) {
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
	v123 = int32(1)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v99
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v98
	if v99 != int32(1601812) {
		goto L5
	} else {
		goto L30
	}
L7:
	;
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+20)) = uint8(v107)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = int64(0)
	goto L5
L8:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)))
	if v100 != int32(1) {
		goto L7
	} else {
		goto L27
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(base.B2i32(v42 != int32(0)))
	v98 = v41
	v99 = v42
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
	case 2, 3, 4:
		goto L13
	default:
		goto L7
	}
L13:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+100)))
	if v93 != int32(1) {
		v98 = v92
		v99 = v91
		goto L8
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
	v98 = v63
	v99 = v58
	goto L8
L20:
	;
	v69 = F_ExecGetResultSlotOps(m, v51, v36+int32(15))
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	v98 = v70
	v99 = v69
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
	v98 = v83
	v99 = v78
	goto L8
L25:
	;
	v89 = F_ExecGetResultSlotOps(m, v71, v36+int32(15))
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v98 = v90
	v99 = v89
	goto L8
L26:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+15)) = uint8(v96)
	v98 = v92
	v99 = v91
	goto L8
L27:
	;
	if v98 == int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	if v99 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L7
L30:
	;
	v123 = int32(0)
	goto L4
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v129 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v151 + int32(1)
	v157 = v150 + v151*int32(40)
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+32)) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+24)) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+16)) = v162
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = v164
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v157))) = v166
	goto L1
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v148
	v150 = v148
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v135 = F_palloc(m, int32(640))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v137 != v129 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	return
L38:
	;
	v148 = v135
	goto L33
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v150 = v139
	goto L32
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v129 << (uint(int32(1)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v146 = F_repalloc(m, v143, v129*int32(80))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v148 = v146
	goto L33
L43:
	;
	v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v321 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L44:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v181 = v9 + int32(8)
	v185 = m.G0
	v187 = v185 - int32(16)
	m.G0 = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)) = uint8(v173)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	if v192 != 0 {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v274 == int32(0) {
		goto L43
	} else {
		goto L73
	}
L46:
	;
	m.G0 = v187 + int32(16)
	goto L45
L47:
	;
	v274 = int32(1)
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+28)) = v250
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+20)) = uint8(v263)
	*(*int32)(unsafe.Add(mBase, uint32(v181)+24)) = v249
	if v250 != int32(1601812) {
		goto L47
	} else {
		goto L72
	}
L49:
	;
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+20)) = uint8(v258)
	*(*int64)(unsafe.Add(mBase, uint32(v181)+24)) = int64(0)
	goto L47
L50:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)))
	if v251 != int32(1) {
		goto L49
	} else {
		goto L69
	}
L51:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)) = uint8(base.B2i32(v193 != int32(0)))
	v249 = v192
	v250 = v193
	goto L50
L52:
	;
	goto L53
L53:
	;
	if v189 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	switch v199 - int32(2) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2, 3, 4:
		goto L55
	default:
		goto L49
	}
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v189)+80))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v189)+76))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+100)))
	if v244 != int32(1) {
		v249 = v243
		v250 = v242
		goto L50
	} else {
		goto L68
	}
L56:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v189)+36))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+101)))
	if v223 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v189)+40))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+102)))
	if v203 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v202 == int32(0) {
		goto L49
	} else {
		goto L62
	}
L59:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+98)))
	if v206 != int32(1) {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v189)+88))
	if v209 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)) = uint8(v212)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v202)+56))
	v249 = v214
	v250 = v209
	goto L50
L62:
	;
	v220 = F_ExecGetResultSlotOps(m, v202, v187+int32(15))
	mBase = m.M
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v202)+56))
	v249 = v221
	v250 = v220
	goto L50
L63:
	;
	if v222 == int32(0) {
		goto L49
	} else {
		goto L67
	}
L64:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+97)))
	if v226 != int32(1) {
		goto L49
	} else {
		goto L65
	}
L65:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v189)+84))
	if v229 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)) = uint8(v232)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v222)+56))
	v249 = v234
	v250 = v229
	goto L50
L67:
	;
	v240 = F_ExecGetResultSlotOps(m, v222, v187+int32(15))
	mBase = m.M
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v222)+56))
	v249 = v241
	v250 = v240
	goto L50
L68:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+15)) = uint8(v247)
	v249 = v243
	v250 = v242
	goto L50
L69:
	;
	if v249 == int32(0) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	if v250 != 0 {
		goto L48
	} else {
		goto L71
	}
L71:
	;
	goto L49
L72:
	;
	v274 = int32(0)
	goto L46
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v280 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v302 + int32(1)
	v308 = v301 + v302*int32(40)
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+32)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+24)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+16)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v308)+8)) = v315
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v308))) = v317
	goto L43
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v299
	v301 = v299
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v286 = F_palloc(m, int32(640))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L37
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v288 != v280 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v299 = v286
	goto L75
L80:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v301 = v290
	goto L74
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v280 << (uint(int32(1)) % 32)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v297 = F_repalloc(m, v294, v280*int32(80))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	v299 = v297
	goto L75
L84:
	;
	v472 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v472 <= int32(0) {
		goto L125
	} else {
		goto L126
	}
L85:
	;
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v332 = v9 + int32(8)
	v336 = m.G0
	v338 = v336 - int32(16)
	m.G0 = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)) = uint8(v324)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v332)+24))
	if v343 != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if v425 == int32(0) {
		goto L84
	} else {
		goto L114
	}
L87:
	;
	m.G0 = v338 + int32(16)
	goto L86
L88:
	;
	v425 = int32(1)
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332)+28)) = v401
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v332)+20)) = uint8(v414)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+24)) = v400
	if v401 != int32(1601812) {
		goto L88
	} else {
		goto L113
	}
L90:
	;
	v409 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v332)+20)) = uint8(v409)
	*(*int64)(unsafe.Add(mBase, uint32(v332)+24)) = int64(0)
	goto L88
L91:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)))
	if v402 != int32(1) {
		goto L90
	} else {
		goto L110
	}
L92:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v332)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)) = uint8(base.B2i32(v344 != int32(0)))
	v400 = v343
	v401 = v344
	goto L91
L93:
	;
	goto L94
L94:
	;
	if v340 == int32(0) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	switch v350 - int32(2) {
	case 0:
		goto L98
	case 1:
		goto L97
	case 2, 3, 4:
		goto L96
	default:
		goto L90
	}
L96:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v340)+80))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v340)+76))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+100)))
	if v395 != int32(1) {
		v400 = v394
		v401 = v393
		goto L91
	} else {
		goto L109
	}
L97:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v340)+36))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+101)))
	if v374 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v340)+40))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+102)))
	if v354 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v353 == int32(0) {
		goto L90
	} else {
		goto L103
	}
L100:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+98)))
	if v357 != int32(1) {
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v340)+88))
	if v360 == int32(0) {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)) = uint8(v363)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v353)+56))
	v400 = v365
	v401 = v360
	goto L91
L103:
	;
	v371 = F_ExecGetResultSlotOps(m, v353, v338+int32(15))
	mBase = m.M
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v353)+56))
	v400 = v372
	v401 = v371
	goto L91
L104:
	;
	if v373 == int32(0) {
		goto L90
	} else {
		goto L108
	}
L105:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+97)))
	if v377 != int32(1) {
		goto L90
	} else {
		goto L106
	}
L106:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v340)+84))
	if v380 == int32(0) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)) = uint8(v383)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v373)+56))
	v400 = v385
	v401 = v380
	goto L91
L108:
	;
	v391 = F_ExecGetResultSlotOps(m, v373, v338+int32(15))
	mBase = m.M
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v373)+56))
	v400 = v392
	v401 = v391
	goto L91
L109:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v338)+15)) = uint8(v398)
	v400 = v394
	v401 = v393
	goto L91
L110:
	;
	if v400 == int32(0) {
		goto L90
	} else {
		goto L111
	}
L111:
	;
	if v401 != 0 {
		goto L89
	} else {
		goto L112
	}
L112:
	;
	goto L90
L113:
	;
	v425 = int32(0)
	goto L87
L114:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v431 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v453 + int32(1)
	v459 = v452 + v453*int32(40)
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v459)+32)) = v460
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v459)+24)) = v462
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v459)+16)) = v464
	v466 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v459)+8)) = v466
	v468 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v459))) = v468
	goto L84
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v450
	v452 = v450
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v437 = F_palloc(m, int32(640))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L37
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v439 != v431 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v450 = v437
	goto L116
L121:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v452 = v441
	goto L115
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v431 << (uint(int32(1)) % 32)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v448 = F_repalloc(m, v445, v431*int32(80))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L37
	} else {
		goto L124
	}
L124:
	;
	v450 = v448
	goto L116
L125:
	;
	v623 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v623 <= int32(0) {
		goto L166
	} else {
		goto L167
	}
L126:
	;
	v475 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v475)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(5)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v483 = v9 + int32(8)
	v487 = m.G0
	v489 = v487 - int32(16)
	m.G0 = v489
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)) = uint8(v475)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v483)+24))
	if v494 != 0 {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	if v576 == int32(0) {
		goto L125
	} else {
		goto L155
	}
L128:
	;
	m.G0 = v489 + int32(16)
	goto L127
L129:
	;
	v576 = int32(1)
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+28)) = v552
	v565 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+20)) = uint8(v565)
	*(*int32)(unsafe.Add(mBase, uint32(v483)+24)) = v551
	if v552 != int32(1601812) {
		goto L129
	} else {
		goto L154
	}
L131:
	;
	v560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+20)) = uint8(v560)
	*(*int64)(unsafe.Add(mBase, uint32(v483)+24)) = int64(0)
	goto L129
L132:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)))
	if v553 != int32(1) {
		goto L131
	} else {
		goto L151
	}
L133:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v483)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)) = uint8(base.B2i32(v495 != int32(0)))
	v551 = v494
	v552 = v495
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v491 == int32(0) {
		goto L131
	} else {
		goto L136
	}
L136:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	switch v501 - int32(2) {
	case 0:
		goto L139
	case 1:
		goto L138
	case 2, 3, 4:
		goto L137
	default:
		goto L131
	}
L137:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v491)+80))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v491)+76))
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+100)))
	if v546 != int32(1) {
		v551 = v545
		v552 = v544
		goto L132
	} else {
		goto L150
	}
L138:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v491)+36))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+101)))
	if v525 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v491)+40))
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+102)))
	if v505 != int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v504 == int32(0) {
		goto L131
	} else {
		goto L144
	}
L141:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+98)))
	if v508 != int32(1) {
		goto L131
	} else {
		goto L142
	}
L142:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v491)+88))
	if v511 == int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v514 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)) = uint8(v514)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v551 = v516
	v552 = v511
	goto L132
L144:
	;
	v522 = F_ExecGetResultSlotOps(m, v504, v489+int32(15))
	mBase = m.M
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v504)+56))
	v551 = v523
	v552 = v522
	goto L132
L145:
	;
	if v524 == int32(0) {
		goto L131
	} else {
		goto L149
	}
L146:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+97)))
	if v528 != int32(1) {
		goto L131
	} else {
		goto L147
	}
L147:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v491)+84))
	if v531 == int32(0) {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)) = uint8(v534)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v524)+56))
	v551 = v536
	v552 = v531
	goto L132
L149:
	;
	v542 = F_ExecGetResultSlotOps(m, v524, v489+int32(15))
	mBase = m.M
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v524)+56))
	v551 = v543
	v552 = v542
	goto L132
L150:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v489)+15)) = uint8(v549)
	v551 = v545
	v552 = v544
	goto L132
L151:
	;
	if v551 == int32(0) {
		goto L131
	} else {
		goto L152
	}
L152:
	;
	if v552 != 0 {
		goto L130
	} else {
		goto L153
	}
L153:
	;
	goto L131
L154:
	;
	v576 = int32(0)
	goto L128
L155:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v582 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v604 + int32(1)
	v610 = v603 + v604*int32(40)
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+32)) = v611
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+24)) = v613
	v615 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+16)) = v615
	v617 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v610)+8)) = v617
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v610))) = v619
	goto L125
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v601
	v603 = v601
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v588 = F_palloc(m, int32(640))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L37
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v590 != v582 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v601 = v588
	goto L157
L162:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v603 = v592
	goto L156
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v582 << (uint(int32(1)) % 32)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v599 = F_repalloc(m, v596, v582*int32(80))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L37
	} else {
		goto L165
	}
L165:
	;
	v601 = v599
	goto L157
L166:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v774 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L167:
	;
	v626 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v626)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(0)
	v634 = v9 + int32(8)
	v638 = m.G0
	v640 = v638 - int32(16)
	m.G0 = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)) = uint8(v626)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v634)+24))
	if v645 != 0 {
		goto L174
	} else {
		goto L175
	}
L168:
	;
	if v727 == int32(0) {
		goto L166
	} else {
		goto L196
	}
L169:
	;
	m.G0 = v640 + int32(16)
	goto L168
L170:
	;
	v727 = int32(1)
	goto L169
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v634)+28)) = v703
	v716 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v634)+20)) = uint8(v716)
	*(*int32)(unsafe.Add(mBase, uint32(v634)+24)) = v702
	if v703 != int32(1601812) {
		goto L170
	} else {
		goto L195
	}
L172:
	;
	v711 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v634)+20)) = uint8(v711)
	*(*int64)(unsafe.Add(mBase, uint32(v634)+24)) = int64(0)
	goto L170
L173:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)))
	if v704 != int32(1) {
		goto L172
	} else {
		goto L192
	}
L174:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v634)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)) = uint8(base.B2i32(v646 != int32(0)))
	v702 = v645
	v703 = v646
	goto L173
L175:
	;
	goto L176
L176:
	;
	if v642 == int32(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	switch v652 - int32(2) {
	case 0:
		goto L180
	case 1:
		goto L179
	case 2, 3, 4:
		goto L178
	default:
		goto L172
	}
L178:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v642)+80))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v642)+76))
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+100)))
	if v697 != int32(1) {
		v702 = v696
		v703 = v695
		goto L173
	} else {
		goto L191
	}
L179:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v642)+36))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+101)))
	if v676 != int32(1) {
		goto L186
	} else {
		goto L187
	}
L180:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v642)+40))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+102)))
	if v656 != int32(1) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v655 == int32(0) {
		goto L172
	} else {
		goto L185
	}
L182:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+98)))
	if v659 != int32(1) {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v642)+88))
	if v662 == int32(0) {
		goto L181
	} else {
		goto L184
	}
L184:
	;
	v665 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)) = uint8(v665)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v655)+56))
	v702 = v667
	v703 = v662
	goto L173
L185:
	;
	v673 = F_ExecGetResultSlotOps(m, v655, v640+int32(15))
	mBase = m.M
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v655)+56))
	v702 = v674
	v703 = v673
	goto L173
L186:
	;
	if v675 == int32(0) {
		goto L172
	} else {
		goto L190
	}
L187:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+97)))
	if v679 != int32(1) {
		goto L172
	} else {
		goto L188
	}
L188:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v642)+84))
	if v682 == int32(0) {
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v685 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)) = uint8(v685)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v675)+56))
	v702 = v687
	v703 = v682
	goto L173
L190:
	;
	v693 = F_ExecGetResultSlotOps(m, v675, v640+int32(15))
	mBase = m.M
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v675)+56))
	v702 = v694
	v703 = v693
	goto L173
L191:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v640)+15)) = uint8(v700)
	v702 = v696
	v703 = v695
	goto L173
L192:
	;
	if v702 == int32(0) {
		goto L172
	} else {
		goto L193
	}
L193:
	;
	if v703 != 0 {
		goto L171
	} else {
		goto L194
	}
L194:
	;
	goto L172
L195:
	;
	v727 = int32(0)
	goto L169
L196:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v733 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v755 + int32(1)
	v761 = v754 + v755*int32(40)
	v762 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+32)) = v762
	v764 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+24)) = v764
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+16)) = v766
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v761)+8)) = v768
	v770 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v761))) = v770
	goto L166
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v752
	v754 = v752
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
	v739 = F_palloc(m, int32(640))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L37
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v741 != v733 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v752 = v739
	goto L198
L203:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v754 = v743
	goto L197
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v733 << (uint(int32(1)) % 32)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v750 = F_repalloc(m, v747, v733*int32(80))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L37
	} else {
		goto L206
	}
L206:
	;
	v752 = v750
	goto L198
L207:
	;
	m.G0 = v9 + int32(48)
	return
L208:
	;
	v777 = int32(0)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	if v778 <= v777 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v786 = v777
	goto L210
L210:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v774)+12))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v791+v786<<(uint(int32(2))%32))))
	F_ExecInitSubPlanExpr(m, v795, l0, l0+int32(8), l0+int32(5))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L37
	} else {
		goto L212
	}
L211:
	;
	goto L207
L212:
	;
	v799 = v786 + int32(1)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v774)+4))
	if v799 < v800 {
		v786 = v799
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
				F_errcode(m, int32(393348))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
					F_errmsg(m, int32(166649), v7)
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
								F_errfinish(m, int32(498071), int32(1935), int32(410082))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v2 = int32(0)
	if l0 == v2 {
		v28 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = v6 - int32(9)
		if base.Ui32(int32(30)) < base.Ui32(v8) {
			v28 = v2
		} else {
			v12 = int32(1) << (uint(v8) % 32)
			if v12&int32(3904) == int32(0) {
				if v12&int32(5) != 0 {
					v24 = int32(16)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
					v28 = v26
				} else {
					if v8 != int32(30) {
						v28 = v2
					} else {
						v24 = int32(12)
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
						v28 = v26
					}
				}
			} else {
				v24 = int32(24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0+v24)))
				v28 = v26
			}
		}
	}
	return v28
}
func F_fix_scan_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 float64
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 float64
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	if l0 == int32(0) {
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
	v16 = l0
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v27 - int32(6) {
	case 0:
		goto L12
	default:
		v123 = v27
		goto L9
	case 2:
		goto L11
	case 3:
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	if v219 != 0 {
		v16 = v219
		goto L4
	} else {
		goto L56
	}
L7:
	;
	v216 = F_copyObjectImpl(m, v108)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L55
	}
L8:
	;
	return v206
L9:
	;
	if v123 != int32(24) {
		goto L36
	} else {
		goto L37
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+276))
	if v65 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v61 = F_fix_param_node(m, v60, v16)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L19
	}
L12:
	;
	v31 = F_palloc(m, int32(48))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v35
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if int32(0) <= v47 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v50 + v47
	goto L17
L16:
	;
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v53 == int32(0) {
		v206 = v31
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v56 + v53
	return v31
L19:
	;
	return v61
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v123 = v120
	goto L9
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v68 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 != int32(1) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v74 <= int32(0) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v82 = int32(0)
	v84 = v74
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v82<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v96 == v97 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	if v108 != 0 {
		goto L7
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v101 = F_equal(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L31
	}
L29:
	;
	v104 = v84
	goto L30
L30:
	;
	v106 = v82 + int32(1)
	if v106 < v104 {
		v82 = v106
		v84 = v104
		goto L25
	} else {
		goto L33
	}
L31:
	;
	if v101 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v104 = v103
	goto L30
L33:
	;
	goto L20
L34:
	;
	goto L20
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_fix_expr_common(m, v198, v16)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L53
	}
L36:
	;
	if v123 != int32(319) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v147 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v149 = int32(0)
	v152 = v149
	v155 = v149
	v159 = float64(0)
	goto L44
L39:
	;
	if v123 != int32(58) {
		goto L35
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v219 = v145
	goto L6
L42:
	;
	v138 = F_copyObjectImpl(m, v16)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v140 + v141
	return v138
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v155<<(uint(int32(2))%32))))
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v167)+56))
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v167)+64))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v148)+360))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v170+v171-v173))) = uint8(v173)
	v178 = base.F64_add(v168, base.F64_mul(v147, v169))
	v180 = int32(0)
	v184 = base.B2i32(base.F64_ge(v159, v178) == v180) & base.B2i32(v152 != v180)
	if v184 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v148)+364))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191+v192-v194))) = uint8(v194)
	v219 = v186
	goto L6
L46:
	;
	v185 = v159
	goto L48
L47:
	;
	v185 = v178
	goto L48
L48:
	;
	if v184 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v186 = v152
	goto L51
L50:
	;
	v186 = v167
	goto L51
L51:
	;
	v188 = v155 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v188 < v189 {
		v152 = v186
		v155 = v188
		v159 = v185
		goto L44
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	v202 = F_expression_tree_mutator_impl(m, v16, int32(839), l1)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v206 = v202
	goto L8
L55:
	;
	return v216
L56:
	;
	goto L5
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(4489440)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
	F_initStringInfo(m, v10+int32(-24))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v359 = int32(0)
	goto L3
L3:
	;
	m.G0 = v12 - int32(-64)
	return v359
L4:
	;
	return int32(0)
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v359 = v348
	goto L3
L7:
	;
	if v83 < int32(0) {
		goto L6
	} else {
		goto L18
	}
L8:
	;
	v83 = base.I32_ctz(v69) | v70<<(uint(int32(5))%32)
	goto L7
L9:
	;
	v83 = int32(-2)
	goto L7
L10:
	;
	v36 = base.I32_div_s(int32(0), int32(32))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v37 <= v36 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v40 = v26 + int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v36<<(uint(int32(2))%32))))
	v47 = v44 & int32(-1)
	if v47 != 0 {
		v69 = v47
		v70 = v36
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v49 = v36 + int32(1)
	if v49 == v37 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v52 = v49
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v40+v52<<(uint(int32(2))%32))))
	if v59 != 0 {
		v69 = v59
		v70 = v52
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L9
L16:
	;
	v61 = v52 + int32(1)
	if v61 != v37 {
		v52 = v61
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v83<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v90, v10+int32(-32), v10+int32(-40), v10+int32(-28), v10+int32(-33))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(741336)
	F_appendStringInfo(m, v10+int32(-24), int32(729150), v10+int32(-48))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v115 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v149 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v120 = int32(4489440)
	v121 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v123
	F_getTypeOutputInfo(m, v119, v10+int32(-4), v10+int32(-5))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_appendStringInfoString(m, v10+int32(-24), int32(532435))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L28
	}
L25:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v132 = F_OidOutputFunctionCall(m, v131, v118)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v121
	F_appendStringInfoStringQuoted(m, v10+int32(-24), v132, int32(-1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	if v205 < int32(0) {
		goto L6
	} else {
		goto L40
	}
L30:
	;
	v205 = base.I32_ctz(v191) | v192<<(uint(int32(5))%32)
	goto L29
L31:
	;
	v205 = int32(-2)
	goto L29
L32:
	;
	v156 = v83 + int32(1)
	v158 = base.I32_div_s(v156, int32(32))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v159 <= v158 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v162 = v149 + int32(8)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v158<<(uint(int32(2))%32))))
	v169 = v166 & (int32(-1) << (uint(v156) % 32))
	if v169 != 0 {
		v191 = v169
		v192 = v158
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v171 = v158 + int32(1)
	if v171 == v159 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v174 = v171
	goto L36
L36:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v162+v174<<(uint(int32(2))%32))))
	if v181 != 0 {
		v191 = v181
		v192 = v174
		goto L30
	} else {
		goto L38
	}
L37:
	;
	goto L31
L38:
	;
	v183 = v174 + int32(1)
	if v183 != v159 {
		v174 = v183
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v212 = v205
	goto L41
L41:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217+v212<<(uint(int32(2))%32))))
	F_exec_eval_datum(m, l0, v221, v10+int32(-32), v10+int32(-40), v10+int32(-28), v10+int32(-33))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L6
L43:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(730094)
	F_appendStringInfo(m, v10+int32(-24), int32(729150), v12)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v244 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v278 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	F_appendStringInfoString(m, v10+int32(-24), int32(532435))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v254 = int32(4489440)
	v255 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v257
	F_getTypeOutputInfo(m, v253, v10+int32(-4), v10+int32(-5))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v266 = F_OidOutputFunctionCall(m, v265, v252)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v255
	F_appendStringInfoStringQuoted(m, v10+int32(-24), v266, int32(-1))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	goto L45
L53:
	;
	if int32(0) <= v334 {
		v212 = v334
		goto L41
	} else {
		goto L64
	}
L54:
	;
	v334 = base.I32_ctz(v320) | v321<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v334 = int32(-2)
	goto L53
L56:
	;
	v285 = v212 + int32(1)
	v287 = base.I32_div_s(v285, int32(32))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v288 <= v287 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v291 = v278 + int32(8)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v287<<(uint(int32(2))%32))))
	v298 = v295 & (int32(-1) << (uint(v285) % 32))
	if v298 != 0 {
		v320 = v298
		v321 = v287
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v300 = v287 + int32(1)
	if v300 == v288 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v303 = v300
	goto L60
L60:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v291+v303<<(uint(int32(2))%32))))
	if v310 != 0 {
		v320 = v310
		v321 = v303
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v312 = v303 + int32(1)
	if v312 != v288 {
		v303 = v312
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
