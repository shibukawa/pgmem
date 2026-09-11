package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetCreate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
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
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
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
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int64
	_ = v715
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	v7 = m.G0
	v9 = v7 - int32(4144)
	m.G0 = v9
	v14 = l0 + int32(12)
	if l1&int32(3) == int32(0) {
		v38 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v77 = v71 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v71 = v63 - l1
	goto L1
L3:
	;
	v42 = v38
	goto L12
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v71 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v27 = l1
	goto L8
L8:
	;
	v31 = v27 + int32(1)
	if v31&int32(3) == int32(0) {
		v38 = v31
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v63 = v31
	goto L2
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 != 0 {
		v27 = v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v51 = int32(-2139062144)
	if (int32(16843008)-v48|v48)&v51 == v51 {
		v42 = v42 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v57 = v42
	goto L15
L14:
	;
	goto L13
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v61 != 0 {
		v57 = v57 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v63 = v57
	goto L2
L17:
	;
	goto L16
L18:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v337 = base.I32_rem_u_s(v331^v323-base.I32_rotl(v331, int32(24)), v336)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v14+v337<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v9+int32(3120), v341)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L58
	} else {
		goto L59
	}
L19:
	;
	v309 = int32(14)
	v311 = v305 ^ v306 - base.I32_rotl(v305, v309)
	v315 = v311 ^ v304 - base.I32_rotl(v311, int32(11))
	v319 = v315 ^ v305 - base.I32_rotl(v315, int32(25))
	v323 = v319 ^ v311 - base.I32_rotl(v319, int32(16))
	v327 = v323 ^ v315 - base.I32_rotl(v323, int32(4))
	v331 = v327 ^ v319 - base.I32_rotl(v327, v309)
	goto L18
L20:
	;
	switch v235 - int32(1) {
	case 0:
		v297 = v236
		v298 = v237
		v299 = v238
		goto L47
	case 1:
		v290 = v236
		v291 = v237
		v292 = v238
		goto L48
	case 2:
		v283 = v236
		v284 = v237
		v285 = v238
		goto L49
	case 3:
		v277 = v237
		v278 = v238
		goto L50
	case 4:
		v273 = v237
		v274 = v238
		goto L51
	case 5:
		v267 = v237
		v268 = v238
		goto L52
	case 6:
		v261 = v237
		v262 = v238
		goto L53
	case 7:
		v256 = v238
		goto L54
	case 8:
		v251 = v238
		goto L55
	case 9:
		v246 = v238
		goto L56
	case 10:
		goto L57
	default:
		v304 = v236
		v305 = v237
		v306 = v238
		goto L19
	}
L21:
	;
	v186 = l1
	v187 = v71
	v188 = v77
	v189 = v77
	v190 = v77
	goto L44
L22:
	;
	if base.Ui32(int32(11)) < base.Ui32(v71) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v71) < base.Ui32(int32(12)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v234 = l1
	v235 = v71
	v236 = v77
	v237 = v77
	v238 = v77
	goto L20
L26:
	;
	switch v133 - int32(1) {
	case 0:
		v183 = v134
		goto L33
	case 1:
		v178 = v134
		goto L34
	case 2:
		goto L35
	case 3:
		v171 = v135
		goto L36
	case 4:
		v168 = v135
		goto L37
	case 5:
		v163 = v135
		goto L38
	case 6:
		goto L39
	case 7:
		v154 = v136
		goto L40
	case 8:
		v149 = v136
		goto L41
	case 9:
		v144 = v136
		goto L42
	case 10:
		goto L43
	default:
		v304 = v134
		v305 = v135
		v306 = v136
		goto L19
	}
L27:
	;
	v132 = l1
	v133 = v71
	v134 = v77
	v135 = v77
	v136 = v77
	goto L26
L28:
	;
	goto L29
L29:
	;
	v84 = l1
	v85 = v71
	v86 = v77
	v87 = v77
	v88 = v77
	goto L30
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v91 = v90 + v87
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v95 = v94 + v88
	v97 = int32(4)
	v99 = v92 + v86 - v95 ^ base.I32_rotl(v95, v97)
	v103 = v91 - v99 ^ base.I32_rotl(v99, int32(6))
	v104 = v95 + v91
	v105 = v99 + v104
	v106 = v103 + v105
	v110 = v104 - v103 ^ base.I32_rotl(v103, int32(8))
	v114 = v105 - v110 ^ base.I32_rotl(v110, int32(16))
	v118 = v106 - v114 ^ base.I32_rotl(v114, int32(19))
	v119 = v110 + v106
	v120 = v114 + v119
	v121 = v118 + v120
	v125 = v119 - v118 ^ base.I32_rotl(v118, v97)
	v126 = int32(12)
	v127 = v84 + v126
	v129 = v85 - v126
	if base.Ui32(int32(11)) < base.Ui32(v129) {
		v84 = v127
		v85 = v129
		v86 = v120
		v87 = v121
		v88 = v125
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v132 = v127
	v133 = v129
	v134 = v120
	v135 = v121
	v136 = v125
	goto L26
L32:
	;
	goto L31
L33:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v304 = v183 + v184
	v305 = v135
	v306 = v136
	goto L19
L34:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	v183 = v179<<(uint(int32(8))%32) + v178
	goto L33
L35:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+2)))
	v178 = v174<<(uint(int32(16))%32) + v134
	goto L34
L36:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v304 = v172 + v134
	v305 = v171
	v306 = v136
	goto L19
L37:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+4)))
	v171 = v168 + v169
	goto L36
L38:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+5)))
	v168 = v164<<(uint(int32(8))%32) + v163
	goto L37
L39:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+6)))
	v163 = v159<<(uint(int32(16))%32) + v135
	goto L38
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v304 = v155 + v134
	v305 = v157 + v135
	v306 = v154
	goto L19
L41:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+8)))
	v154 = v150<<(uint(int32(8))%32) + v149
	goto L40
L42:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+9)))
	v149 = v145<<(uint(int32(16))%32) + v144
	goto L41
L43:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+10)))
	v144 = v140<<(uint(int32(24))%32) + v136
	goto L42
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	v193 = v192 + v189
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v197 = v196 + v190
	v199 = int32(4)
	v201 = v194 + v188 - v197 ^ base.I32_rotl(v197, v199)
	v205 = v193 - v201 ^ base.I32_rotl(v201, int32(6))
	v206 = v197 + v193
	v207 = v201 + v206
	v208 = v205 + v207
	v212 = v206 - v205 ^ base.I32_rotl(v205, int32(8))
	v216 = v207 - v212 ^ base.I32_rotl(v212, int32(16))
	v220 = v208 - v216 ^ base.I32_rotl(v216, int32(19))
	v221 = v212 + v208
	v222 = v216 + v221
	v223 = v220 + v222
	v227 = v221 - v220 ^ base.I32_rotl(v220, v199)
	v228 = int32(12)
	v229 = v186 + v228
	v231 = v187 - v228
	if base.Ui32(int32(11)) < base.Ui32(v231) {
		v186 = v229
		v187 = v231
		v188 = v222
		v189 = v223
		v190 = v227
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v234 = v229
	v235 = v231
	v236 = v222
	v237 = v223
	v238 = v227
	goto L20
L46:
	;
	goto L45
L47:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	v304 = v297 + v300
	v305 = v298
	v306 = v299
	goto L19
L48:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	v297 = v293<<(uint(int32(8))%32) + v290
	v298 = v291
	v299 = v292
	goto L47
L49:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+2)))
	v290 = v286<<(uint(int32(16))%32) + v283
	v291 = v284
	v292 = v285
	goto L48
L50:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+3)))
	v283 = v279<<(uint(int32(24))%32) + v236
	v284 = v277
	v285 = v278
	goto L49
L51:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+4)))
	v277 = v273 + v275
	v278 = v274
	goto L50
L52:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+5)))
	v273 = v269<<(uint(int32(8))%32) + v267
	v274 = v268
	goto L51
L53:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+6)))
	v267 = v263<<(uint(int32(16))%32) + v261
	v268 = v262
	goto L52
L54:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+7)))
	v261 = v257<<(uint(int32(24))%32) + v237
	v262 = v256
	goto L53
L55:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+8)))
	v256 = v252<<(uint(int32(8))%32) + v251
	goto L54
L56:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+9)))
	v251 = v247<<(uint(int32(16))%32) + v246
	goto L55
L57:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+10)))
	v246 = v242<<(uint(int32(24))%32) + v238
	goto L56
L58:
	;
	return int32(0)
L59:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(220933)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(3120)
	v359 = F_pg_snprintf(m, v9+int32(2096), int32(1024), int32(97391), v9+int32(32))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(2096)
	v371 = F_pg_snprintf(m, v9+int32(1072), int32(1024), int32(165202), v9+int32(16))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v376 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	if v376 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if l1&int32(3) == int32(0) {
		v405 = l1
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v805 = v376
	goto L65
L65:
	;
	m.G0 = v9 + int32(4144)
	return v805
L66:
	;
	v444 = v438 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L67:
	;
	v438 = v430 - l1
	goto L66
L68:
	;
	v409 = v405
	goto L77
L69:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v389 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v438 = int32(0)
	goto L66
L71:
	;
	goto L72
L72:
	;
	v394 = l1
	goto L73
L73:
	;
	v398 = v394 + int32(1)
	if v398&int32(3) == int32(0) {
		v405 = v398
		goto L68
	} else {
		goto L75
	}
L74:
	;
	v430 = v398
	goto L67
L75:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398))))
	if v403 != 0 {
		v394 = v398
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v418 = int32(-2139062144)
	if (int32(16843008)-v415|v415)&v418 == v418 {
		v409 = v409 + int32(4)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v424 = v409
	goto L80
L79:
	;
	goto L78
L80:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v428 != 0 {
		v424 = v424 + int32(1)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v430 = v424
	goto L67
L82:
	;
	goto L81
L83:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v704 = base.I32_rem_u_s(v698^v690-base.I32_rotl(v698, int32(24)), v703)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v14+v704<<(uint(int32(2))%32))))
	F_TempTablespacePath(m, v9+int32(2096), v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L58
	} else {
		goto L123
	}
L84:
	;
	v676 = int32(14)
	v678 = v672 ^ v673 - base.I32_rotl(v672, v676)
	v682 = v678 ^ v671 - base.I32_rotl(v678, int32(11))
	v686 = v682 ^ v672 - base.I32_rotl(v682, int32(25))
	v690 = v686 ^ v678 - base.I32_rotl(v686, int32(16))
	v694 = v690 ^ v682 - base.I32_rotl(v690, int32(4))
	v698 = v694 ^ v686 - base.I32_rotl(v694, v676)
	goto L83
L85:
	;
	switch v602 - int32(1) {
	case 0:
		v664 = v603
		v665 = v604
		v666 = v605
		goto L112
	case 1:
		v657 = v603
		v658 = v604
		v659 = v605
		goto L113
	case 2:
		v650 = v603
		v651 = v604
		v652 = v605
		goto L114
	case 3:
		v644 = v604
		v645 = v605
		goto L115
	case 4:
		v640 = v604
		v641 = v605
		goto L116
	case 5:
		v634 = v604
		v635 = v605
		goto L117
	case 6:
		v628 = v604
		v629 = v605
		goto L118
	case 7:
		v623 = v605
		goto L119
	case 8:
		v618 = v605
		goto L120
	case 9:
		v613 = v605
		goto L121
	case 10:
		goto L122
	default:
		v671 = v603
		v672 = v604
		v673 = v605
		goto L84
	}
L86:
	;
	v553 = l1
	v554 = v438
	v555 = v444
	v556 = v444
	v557 = v444
	goto L109
L87:
	;
	if base.Ui32(int32(11)) < base.Ui32(v438) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v438) < base.Ui32(int32(12)) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v601 = l1
	v602 = v438
	v603 = v444
	v604 = v444
	v605 = v444
	goto L85
L91:
	;
	switch v500 - int32(1) {
	case 0:
		v550 = v501
		goto L98
	case 1:
		v545 = v501
		goto L99
	case 2:
		goto L100
	case 3:
		v538 = v502
		goto L101
	case 4:
		v535 = v502
		goto L102
	case 5:
		v530 = v502
		goto L103
	case 6:
		goto L104
	case 7:
		v521 = v503
		goto L105
	case 8:
		v516 = v503
		goto L106
	case 9:
		v511 = v503
		goto L107
	case 10:
		goto L108
	default:
		v671 = v501
		v672 = v502
		v673 = v503
		goto L84
	}
L92:
	;
	v499 = l1
	v500 = v438
	v501 = v444
	v502 = v444
	v503 = v444
	goto L91
L93:
	;
	goto L94
L94:
	;
	v451 = l1
	v452 = v438
	v453 = v444
	v454 = v444
	v455 = v444
	goto L95
L95:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v458 = v457 + v454
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	v462 = v461 + v455
	v464 = int32(4)
	v466 = v459 + v453 - v462 ^ base.I32_rotl(v462, v464)
	v470 = v458 - v466 ^ base.I32_rotl(v466, int32(6))
	v471 = v462 + v458
	v472 = v466 + v471
	v473 = v470 + v472
	v477 = v471 - v470 ^ base.I32_rotl(v470, int32(8))
	v481 = v472 - v477 ^ base.I32_rotl(v477, int32(16))
	v485 = v473 - v481 ^ base.I32_rotl(v481, int32(19))
	v486 = v477 + v473
	v487 = v481 + v486
	v488 = v485 + v487
	v492 = v486 - v485 ^ base.I32_rotl(v485, v464)
	v493 = int32(12)
	v494 = v451 + v493
	v496 = v452 - v493
	if base.Ui32(int32(11)) < base.Ui32(v496) {
		v451 = v494
		v452 = v496
		v453 = v487
		v454 = v488
		v455 = v492
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v499 = v494
	v500 = v496
	v501 = v487
	v502 = v488
	v503 = v492
	goto L91
L97:
	;
	goto L96
L98:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v671 = v550 + v551
	v672 = v502
	v673 = v503
	goto L84
L99:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	v550 = v546<<(uint(int32(8))%32) + v545
	goto L98
L100:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+2)))
	v545 = v541<<(uint(int32(16))%32) + v501
	goto L99
L101:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v671 = v539 + v501
	v672 = v538
	v673 = v503
	goto L84
L102:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+4)))
	v538 = v535 + v536
	goto L101
L103:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+5)))
	v535 = v531<<(uint(int32(8))%32) + v530
	goto L102
L104:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+6)))
	v530 = v526<<(uint(int32(16))%32) + v502
	goto L103
L105:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	v671 = v522 + v501
	v672 = v524 + v502
	v673 = v521
	goto L84
L106:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+8)))
	v521 = v517<<(uint(int32(8))%32) + v516
	goto L105
L107:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+9)))
	v516 = v512<<(uint(int32(16))%32) + v511
	goto L106
L108:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+10)))
	v511 = v507<<(uint(int32(24))%32) + v503
	goto L107
L109:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v560 = v559 + v556
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v564 = v563 + v557
	v566 = int32(4)
	v568 = v561 + v555 - v564 ^ base.I32_rotl(v564, v566)
	v572 = v560 - v568 ^ base.I32_rotl(v568, int32(6))
	v573 = v564 + v560
	v574 = v568 + v573
	v575 = v572 + v574
	v579 = v573 - v572 ^ base.I32_rotl(v572, int32(8))
	v583 = v574 - v579 ^ base.I32_rotl(v579, int32(16))
	v587 = v575 - v583 ^ base.I32_rotl(v583, int32(19))
	v588 = v579 + v575
	v589 = v583 + v588
	v590 = v587 + v589
	v594 = v588 - v587 ^ base.I32_rotl(v587, v566)
	v595 = int32(12)
	v596 = v553 + v595
	v598 = v554 - v595
	if base.Ui32(int32(11)) < base.Ui32(v598) {
		v553 = v596
		v554 = v598
		v555 = v589
		v556 = v590
		v557 = v594
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v601 = v596
	v602 = v598
	v603 = v589
	v604 = v590
	v605 = v594
	goto L85
L111:
	;
	goto L110
L112:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v671 = v664 + v667
	v672 = v665
	v673 = v666
	goto L84
L113:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	v664 = v660<<(uint(int32(8))%32) + v657
	v665 = v658
	v666 = v659
	goto L112
L114:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+2)))
	v657 = v653<<(uint(int32(16))%32) + v650
	v658 = v651
	v659 = v652
	goto L113
L115:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+3)))
	v650 = v646<<(uint(int32(24))%32) + v603
	v651 = v644
	v652 = v645
	goto L114
L116:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+4)))
	v644 = v640 + v642
	v645 = v641
	goto L115
L117:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+5)))
	v640 = v636<<(uint(int32(8))%32) + v634
	v641 = v635
	goto L116
L118:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+6)))
	v634 = v630<<(uint(int32(16))%32) + v628
	v635 = v629
	goto L117
L119:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+7)))
	v628 = v624<<(uint(int32(24))%32) + v604
	v629 = v623
	goto L118
L120:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+8)))
	v623 = v619<<(uint(int32(8))%32) + v618
	goto L119
L121:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+9)))
	v618 = v614<<(uint(int32(16))%32) + v613
	goto L120
L122:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+10)))
	v613 = v609<<(uint(int32(24))%32) + v605
	goto L121
L123:
	;
	F_TempTablespacePath(m, v9+int32(3120), v708)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L58
	} else {
		goto L124
	}
L124:
	;
	v715 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(220933)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(3120)
	v726 = F_pg_snprintf(m, v9+int32(48), int32(1024), int32(97391), v9)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L58
	} else {
		goto L125
	}
L125:
	;
	v729 = v9 + int32(2096)
	v730 = m.G0
	v732 = v730 - int32(32)
	m.G0 = v732
	v735 = v9 + int32(48)
	v737 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v738 = F_mkdir(m, v735, v737)
	mBase = m.M
	if int32(0) <= v738 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v801 = F_PathNameCreateTemporaryFile(m, v9+int32(1072), int32(1))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L58
	} else {
		goto L146
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L58
	} else {
		goto L142
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L58
	} else {
		goto L138
	}
L129:
	;
	m.G0 = v732 + int32(32)
	goto L126
L130:
	;
	v742 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v742 == int32(20) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v747 = F_mkdir(m, v729, v746)
	mBase = m.M
	if v747 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v751 != int32(20) {
		goto L128
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _consts[383]))
	v756 = F_mkdir(m, v735, v755)
	mBase = m.M
	if int32(0) <= v756 {
		goto L129
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v760 != int32(20) {
		goto L127
	} else {
		goto L137
	}
L137:
	;
	goto L129
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L58
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+16)) = v729
	F_errmsg(m, int32(278105), v732+int32(16))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L58
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(469068), int32(1685), int32(200866))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L58
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L58
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732))) = v735
	F_errmsg(m, int32(278059), v732)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L58
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(469068), int32(1692), int32(200866))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L58
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	v805 = v801
	goto L65
}
func F_FileSetInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v13
	v15 = int32(4341328)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[760]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16
	v22 = base.I32_rem_u_s(v16+int32(1), int32(2147483647))
	*(*int32)(unsafe.Add(mBase, _consts[760])) = v22
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		v27 = l0 + int32(12)
		v28 = int32(8)
		v30 = *(*int32)(unsafe.Add(mBase, _consts[405]))
		if v28 < v30 {
			v33 = v28
		} else {
			v33 = v30
		}
		if int32(0) < v33 {
			v37 = *(*int32)(unsafe.Add(mBase, _consts[406]))
			if base.Ui32(int32(4)) <= base.Ui32(v33) {
				v44 = v2
				v50 = v2
				for {
					v54 = v44 << (uint(int32(2)) % 32)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v54))) = v57
					v59 = int32(4)
					v60 = v54 | v59
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v37+v60)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v60))) = v63
					v66 = v54 | int32(8)
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v37+v66)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v66))) = v69
					v72 = v54 | int32(12)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v72+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v72))) = v75
					v78 = v44 + v59
					v80 = v50 + v59
					if v80 != v33&int32(2147483644) {
						v44 = v78
						v50 = v80
						continue
					} else {
						break
					}
					break
				}
				v84 = v78
			} else {
				v84 = v2
			}
			v94 = v33 & int32(3)
			if v94 != 0 {
				v97 = v84
				v104 = v2
				for {
					v107 = v97 << (uint(int32(2)) % 32)
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v27+v107))) = v110
					v112 = int32(1)
					v115 = v104 + v112
					if v115 != v94 {
						v97 = v97 + v112
						v104 = v115
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v140 = v33
		} else {
			v140 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
		if v140 != 0 {
			if int32(0) < v140 {
				v145 = v140
				v151 = v2
				for {
					v157 = v27 + v151<<(uint(int32(2))%32)
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
					if v158 == int32(0) {
						v162 = *(*int32)(unsafe.Add(mBase, _consts[169]))
						*(*int32)(unsafe.Add(mBase, uint32(v157))) = v162
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v165 = v164
					} else {
						v165 = v145
					}
					v167 = v151 + int32(1)
					if v167 < v165 {
						v145 = v165
						v151 = v167
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return
		} else {
			v181 = *(*int32)(unsafe.Add(mBase, _consts[169]))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
			return
		}
	}
}
