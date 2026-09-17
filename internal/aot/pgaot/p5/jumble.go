package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AppendJumble(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
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
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
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
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
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
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int64
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(1024)-v383) < base.Ui32(l2) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v383 = v12
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v15-int32(1021)) < base.Ui32(v13) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372
	v383 = v372
	goto L1
L6:
	;
	v25 = v15
	v27 = v13
	v29 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v14))) = v9
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v372 = v366 + int32(4)
	goto L5
L9:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v25) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v372 = v362
	goto L5
L11:
	;
	v32 = int32(1024)
	v39 = int32(-1636607408)
	goto L16
L12:
	;
	v354 = v25
	goto L13
L13:
	;
	v356 = int32(1024) - v354
	if base.Ui32(v27) < base.Ui32(v356) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v349 = F_Int64GetDatum(m, base.I64_extend_i32_u(v339)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v339^v331-base.I32_rotl(v339, int32(24))))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L57
	} else {
		goto L58
	}
L15:
	;
	if v14&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	goto L15
L19:
	;
	v317 = int32(14)
	v319 = v313 ^ v314 - base.I32_rotl(v313, v317)
	v323 = v319 ^ v312 - base.I32_rotl(v319, int32(11))
	v327 = v323 ^ v313 - base.I32_rotl(v323, int32(25))
	v331 = v327 ^ v319 - base.I32_rotl(v327, int32(16))
	v335 = v331 ^ v323 - base.I32_rotl(v331, int32(4))
	v339 = v335 ^ v327 - base.I32_rotl(v335, v317)
	goto L14
L20:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v312 = v304 + v307
	v313 = v305
	v314 = v306
	goto L19
L21:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)))
	v304 = v300<<(uint(int32(8))%32) + v297
	v305 = v298
	v306 = v299
	goto L20
L22:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+2)))
	v297 = v293<<(uint(int32(16))%32) + v290
	v298 = v291
	v299 = v292
	goto L21
L23:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+3)))
	v290 = v286<<(uint(int32(24))%32) + v122
	v291 = v284
	v292 = v285
	goto L22
L24:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)))
	v284 = v280 + v282
	v285 = v281
	goto L23
L25:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+5)))
	v280 = v276<<(uint(int32(8))%32) + v274
	v281 = v275
	goto L24
L26:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+6)))
	v274 = v270<<(uint(int32(16))%32) + v268
	v275 = v269
	goto L25
L27:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+7)))
	v268 = v264<<(uint(int32(24))%32) + v123
	v269 = v263
	goto L26
L28:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)))
	v263 = v259<<(uint(int32(8))%32) + v258
	goto L27
L29:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+9)))
	v258 = v254<<(uint(int32(16))%32) + v253
	goto L28
L30:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+10)))
	v253 = v249<<(uint(int32(24))%32) + v127
	goto L29
L31:
	;
	goto L34
L32:
	;
	goto L33
L33:
	;
	goto L40
L34:
	;
	v85 = v14
	v86 = v32
	v88 = v39
	v89 = v39
	v90 = v39
	goto L37
L36:
	;
	switch v131 - int32(1) {
	case 0:
		v304 = v122
		v305 = v123
		v306 = v127
		goto L20
	case 1:
		v297 = v122
		v298 = v123
		v299 = v127
		goto L21
	case 2:
		v290 = v122
		v291 = v123
		v292 = v127
		goto L22
	case 3:
		v284 = v123
		v285 = v127
		goto L23
	case 4:
		v280 = v123
		v281 = v127
		goto L24
	case 5:
		v274 = v123
		v275 = v127
		goto L25
	case 6:
		v268 = v123
		v269 = v127
		goto L26
	case 7:
		v263 = v127
		goto L27
	case 8:
		v258 = v127
		goto L28
	case 9:
		v253 = v127
		goto L29
	case 10:
		goto L30
	default:
		v312 = v122
		v313 = v123
		v314 = v127
		goto L19
	}
L37:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v93 = v92 + v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v97 = v96 + v90
	v99 = int32(4)
	v101 = v94 + v88 - v97 ^ base.I32_rotl(v97, v99)
	v105 = v93 - v101 ^ base.I32_rotl(v101, int32(6))
	v106 = v97 + v93
	v107 = v101 + v106
	v108 = v105 + v107
	v112 = v106 - v105 ^ base.I32_rotl(v105, int32(8))
	v116 = v107 - v112 ^ base.I32_rotl(v112, int32(16))
	v120 = v108 - v116 ^ base.I32_rotl(v116, int32(19))
	v121 = v112 + v108
	v122 = v116 + v121
	v123 = v120 + v122
	v127 = v121 - v120 ^ base.I32_rotl(v120, v99)
	v128 = int32(12)
	v129 = v85 + v128
	v131 = v86 - v128
	if base.Ui32(int32(11)) < base.Ui32(v131) {
		v85 = v129
		v86 = v131
		v88 = v122
		v89 = v123
		v90 = v127
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v145 = v14
	v146 = v32
	v148 = v39
	v149 = v39
	v150 = v39
	goto L43
L42:
	;
	switch v191 - int32(1) {
	case 0:
		v246 = v182
		goto L46
	case 1:
		v241 = v182
		goto L47
	case 2:
		goto L48
	case 3:
		v234 = v183
		goto L49
	case 4:
		v231 = v183
		goto L50
	case 5:
		v226 = v183
		goto L51
	case 6:
		goto L52
	case 7:
		v217 = v187
		goto L53
	case 8:
		v212 = v187
		goto L54
	case 9:
		v207 = v187
		goto L55
	case 10:
		goto L56
	default:
		v312 = v182
		v313 = v183
		v314 = v187
		goto L19
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v153 = v152 + v149
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v157 = v156 + v150
	v159 = int32(4)
	v161 = v154 + v148 - v157 ^ base.I32_rotl(v157, v159)
	v165 = v153 - v161 ^ base.I32_rotl(v161, int32(6))
	v166 = v157 + v153
	v167 = v161 + v166
	v168 = v165 + v167
	v172 = v166 - v165 ^ base.I32_rotl(v165, int32(8))
	v176 = v167 - v172 ^ base.I32_rotl(v172, int32(16))
	v180 = v168 - v176 ^ base.I32_rotl(v176, int32(19))
	v181 = v172 + v168
	v182 = v176 + v181
	v183 = v180 + v182
	v187 = v181 - v180 ^ base.I32_rotl(v180, v159)
	v188 = int32(12)
	v189 = v145 + v188
	v191 = v146 - v188
	if base.Ui32(int32(11)) < base.Ui32(v191) {
		v145 = v189
		v146 = v191
		v148 = v182
		v149 = v183
		v150 = v187
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v312 = v246 + v247
	v313 = v183
	v314 = v187
	goto L19
L47:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v246 = v242<<(uint(int32(8))%32) + v241
	goto L46
L48:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+2)))
	v241 = v237<<(uint(int32(16))%32) + v182
	goto L47
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v312 = v235 + v182
	v313 = v234
	v314 = v187
	goto L19
L50:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
	v234 = v231 + v232
	goto L49
L51:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)))
	v231 = v227<<(uint(int32(8))%32) + v226
	goto L50
L52:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+6)))
	v226 = v222<<(uint(int32(16))%32) + v183
	goto L51
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v312 = v218 + v182
	v313 = v220 + v183
	v314 = v217
	goto L19
L54:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)))
	v217 = v213<<(uint(int32(8))%32) + v212
	goto L53
L55:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+9)))
	v212 = v208<<(uint(int32(16))%32) + v207
	goto L54
L56:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+10)))
	v207 = v203<<(uint(int32(24))%32) + v187
	goto L55
L57:
	;
	return
L58:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v349)))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v351
	v354 = int32(8)
	goto L13
L59:
	;
	v358 = v27
	goto L61
L60:
	;
	v358 = v356
	goto L61
L61:
	;
	if v358 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	base.MemoryCopy(m, v354+v14, v29, v358)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v362 = v354 + v358
	v363 = v27 - v358
	if v363 != 0 {
		v25 = v362
		v27 = v363
		v29 = v358 + v29
		goto L9
	} else {
		goto L65
	}
L65:
	;
	goto L10
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v741
	return
L67:
	;
	v393 = l1
	v394 = l2
	v395 = v383
	goto L70
L68:
	;
	goto L69
L69:
	;
	if l2 != 0 {
		goto L126
	} else {
		goto L127
	}
L70:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v395) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v402 = int32(1024)
	v409 = int32(-1636607408)
	goto L77
L73:
	;
	v724 = v395
	goto L74
L74:
	;
	v726 = int32(1024) - v724
	if base.Ui32(v394) < base.Ui32(v726) {
		goto L119
	} else {
		goto L120
	}
L75:
	;
	v719 = F_Int64GetDatum(m, base.I64_extend_i32_u(v709)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v709^v701-base.I32_rotl(v709, int32(24))))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L57
	} else {
		goto L118
	}
L76:
	;
	if v388&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L77:
	;
	goto L76
L80:
	;
	v687 = int32(14)
	v689 = v683 ^ v684 - base.I32_rotl(v683, v687)
	v693 = v689 ^ v682 - base.I32_rotl(v689, int32(11))
	v697 = v693 ^ v683 - base.I32_rotl(v693, int32(25))
	v701 = v697 ^ v689 - base.I32_rotl(v697, int32(16))
	v705 = v701 ^ v693 - base.I32_rotl(v701, int32(4))
	v709 = v705 ^ v697 - base.I32_rotl(v705, v687)
	goto L75
L81:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	v682 = v674 + v677
	v683 = v675
	v684 = v676
	goto L80
L82:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+1)))
	v674 = v670<<(uint(int32(8))%32) + v667
	v675 = v668
	v676 = v669
	goto L81
L83:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+2)))
	v667 = v663<<(uint(int32(16))%32) + v660
	v668 = v661
	v669 = v662
	goto L82
L84:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+3)))
	v660 = v656<<(uint(int32(24))%32) + v492
	v661 = v654
	v662 = v655
	goto L83
L85:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+4)))
	v654 = v650 + v652
	v655 = v651
	goto L84
L86:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+5)))
	v650 = v646<<(uint(int32(8))%32) + v644
	v651 = v645
	goto L85
L87:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+6)))
	v644 = v640<<(uint(int32(16))%32) + v638
	v645 = v639
	goto L86
L88:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+7)))
	v638 = v634<<(uint(int32(24))%32) + v493
	v639 = v633
	goto L87
L89:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+8)))
	v633 = v629<<(uint(int32(8))%32) + v628
	goto L88
L90:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+9)))
	v628 = v624<<(uint(int32(16))%32) + v623
	goto L89
L91:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+10)))
	v623 = v619<<(uint(int32(24))%32) + v497
	goto L90
L92:
	;
	goto L95
L93:
	;
	goto L94
L94:
	;
	goto L101
L95:
	;
	v455 = v388
	v456 = v402
	v458 = v409
	v459 = v409
	v460 = v409
	goto L98
L97:
	;
	switch v501 - int32(1) {
	case 0:
		v674 = v492
		v675 = v493
		v676 = v497
		goto L81
	case 1:
		v667 = v492
		v668 = v493
		v669 = v497
		goto L82
	case 2:
		v660 = v492
		v661 = v493
		v662 = v497
		goto L83
	case 3:
		v654 = v493
		v655 = v497
		goto L84
	case 4:
		v650 = v493
		v651 = v497
		goto L85
	case 5:
		v644 = v493
		v645 = v497
		goto L86
	case 6:
		v638 = v493
		v639 = v497
		goto L87
	case 7:
		v633 = v497
		goto L88
	case 8:
		v628 = v497
		goto L89
	case 9:
		v623 = v497
		goto L90
	case 10:
		goto L91
	default:
		v682 = v492
		v683 = v493
		v684 = v497
		goto L80
	}
L98:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v463 = v462 + v459
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v455)+8))
	v467 = v466 + v460
	v469 = int32(4)
	v471 = v464 + v458 - v467 ^ base.I32_rotl(v467, v469)
	v475 = v463 - v471 ^ base.I32_rotl(v471, int32(6))
	v476 = v467 + v463
	v477 = v471 + v476
	v478 = v475 + v477
	v482 = v476 - v475 ^ base.I32_rotl(v475, int32(8))
	v486 = v477 - v482 ^ base.I32_rotl(v482, int32(16))
	v490 = v478 - v486 ^ base.I32_rotl(v486, int32(19))
	v491 = v482 + v478
	v492 = v486 + v491
	v493 = v490 + v492
	v497 = v491 - v490 ^ base.I32_rotl(v490, v469)
	v498 = int32(12)
	v499 = v455 + v498
	v501 = v456 - v498
	if base.Ui32(int32(11)) < base.Ui32(v501) {
		v455 = v499
		v456 = v501
		v458 = v492
		v459 = v493
		v460 = v497
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	goto L99
L101:
	;
	v515 = v388
	v516 = v402
	v518 = v409
	v519 = v409
	v520 = v409
	goto L104
L103:
	;
	switch v561 - int32(1) {
	case 0:
		v616 = v552
		goto L107
	case 1:
		v611 = v552
		goto L108
	case 2:
		goto L109
	case 3:
		v604 = v553
		goto L110
	case 4:
		v601 = v553
		goto L111
	case 5:
		v596 = v553
		goto L112
	case 6:
		goto L113
	case 7:
		v587 = v557
		goto L114
	case 8:
		v582 = v557
		goto L115
	case 9:
		v577 = v557
		goto L116
	case 10:
		goto L117
	default:
		v682 = v552
		v683 = v553
		v684 = v557
		goto L80
	}
L104:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	v523 = v522 + v519
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	v527 = v526 + v520
	v529 = int32(4)
	v531 = v524 + v518 - v527 ^ base.I32_rotl(v527, v529)
	v535 = v523 - v531 ^ base.I32_rotl(v531, int32(6))
	v536 = v527 + v523
	v537 = v531 + v536
	v538 = v535 + v537
	v542 = v536 - v535 ^ base.I32_rotl(v535, int32(8))
	v546 = v537 - v542 ^ base.I32_rotl(v542, int32(16))
	v550 = v538 - v546 ^ base.I32_rotl(v546, int32(19))
	v551 = v542 + v538
	v552 = v546 + v551
	v553 = v550 + v552
	v557 = v551 - v550 ^ base.I32_rotl(v550, v529)
	v558 = int32(12)
	v559 = v515 + v558
	v561 = v516 - v558
	if base.Ui32(int32(11)) < base.Ui32(v561) {
		v515 = v559
		v516 = v561
		v518 = v552
		v519 = v553
		v520 = v557
		goto L104
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	goto L105
L107:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559))))
	v682 = v616 + v617
	v683 = v553
	v684 = v557
	goto L80
L108:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+1)))
	v616 = v612<<(uint(int32(8))%32) + v611
	goto L107
L109:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+2)))
	v611 = v607<<(uint(int32(16))%32) + v552
	goto L108
L110:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v682 = v605 + v552
	v683 = v604
	v684 = v557
	goto L80
L111:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+4)))
	v604 = v601 + v602
	goto L110
L112:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+5)))
	v601 = v597<<(uint(int32(8))%32) + v596
	goto L111
L113:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+6)))
	v596 = v592<<(uint(int32(16))%32) + v553
	goto L112
L114:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	v682 = v588 + v552
	v683 = v590 + v553
	v684 = v587
	goto L80
L115:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+8)))
	v587 = v583<<(uint(int32(8))%32) + v582
	goto L114
L116:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+9)))
	v582 = v578<<(uint(int32(16))%32) + v577
	goto L115
L117:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559)+10)))
	v577 = v573<<(uint(int32(24))%32) + v557
	goto L116
L118:
	;
	v721 = *(*int64)(unsafe.Add(mBase, uint32(v719)))
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = v721
	v724 = int32(8)
	goto L74
L119:
	;
	v728 = v394
	goto L121
L120:
	;
	v728 = v726
	goto L121
L121:
	;
	if v728 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	base.MemoryCopy(m, v724+v388, v393, v728)
	goto L124
L123:
	;
	goto L124
L124:
	;
	v732 = v724 + v728
	v733 = v394 - v728
	if v733 != 0 {
		v393 = v393 + v728
		v394 = v733
		v395 = v732
		goto L70
	} else {
		goto L125
	}
L125:
	;
	v741 = v732
	goto L66
L126:
	;
	base.MemoryCopy(m, v383+v388, l1, l2)
	goto L128
L127:
	;
	goto L128
L128:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v741 = v736 + l2
	goto L66
}
func F_JumbleQuery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
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
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
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
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_palloc(m, int32(1024))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
	v23 = F_palloc(m, int32(384))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v25)
	F__jumbleNode(m, v10, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = int32(4)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.Ui32(v37-int32(1021)) < base.Ui32(v35) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v410 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v393
	goto L8
L10:
	;
	v46 = v37
	v48 = v35
	v50 = v10 + int32(28)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v36))) = v34
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v393 = v388 + int32(4)
	goto L9
L13:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v46) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v393 = v384
	goto L9
L15:
	;
	v54 = int32(1024)
	v61 = int32(-1636607408)
	goto L20
L16:
	;
	v376 = v46
	goto L17
L17:
	;
	v378 = int32(1024) - v376
	if base.Ui32(v48) < base.Ui32(v378) {
		goto L62
	} else {
		goto L63
	}
L18:
	;
	v371 = F_Int64GetDatum(m, base.I64_extend_i32_u(v361)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v361^v353-base.I32_rotl(v361, int32(24))))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L61
	}
L19:
	;
	if v36&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	goto L19
L23:
	;
	v339 = int32(14)
	v341 = v335 ^ v336 - base.I32_rotl(v335, v339)
	v345 = v341 ^ v334 - base.I32_rotl(v341, int32(11))
	v349 = v345 ^ v335 - base.I32_rotl(v345, int32(25))
	v353 = v349 ^ v341 - base.I32_rotl(v349, int32(16))
	v357 = v353 ^ v345 - base.I32_rotl(v353, int32(4))
	v361 = v357 ^ v349 - base.I32_rotl(v357, v339)
	goto L18
L24:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v334 = v326 + v329
	v335 = v327
	v336 = v328
	goto L23
L25:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+1)))
	v326 = v322<<(uint(int32(8))%32) + v319
	v327 = v320
	v328 = v321
	goto L24
L26:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)))
	v319 = v315<<(uint(int32(16))%32) + v312
	v320 = v313
	v321 = v314
	goto L25
L27:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
	v312 = v308<<(uint(int32(24))%32) + v144
	v313 = v306
	v314 = v307
	goto L26
L28:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
	v306 = v302 + v304
	v307 = v303
	goto L27
L29:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+5)))
	v302 = v298<<(uint(int32(8))%32) + v296
	v303 = v297
	goto L28
L30:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+6)))
	v296 = v292<<(uint(int32(16))%32) + v290
	v297 = v291
	goto L29
L31:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+7)))
	v290 = v286<<(uint(int32(24))%32) + v145
	v291 = v285
	goto L30
L32:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
	v285 = v281<<(uint(int32(8))%32) + v280
	goto L31
L33:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+9)))
	v280 = v276<<(uint(int32(16))%32) + v275
	goto L32
L34:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+10)))
	v275 = v271<<(uint(int32(24))%32) + v149
	goto L33
L35:
	;
	goto L38
L36:
	;
	goto L37
L37:
	;
	goto L44
L38:
	;
	v107 = v36
	v108 = v54
	v110 = v61
	v111 = v61
	v112 = v61
	goto L41
L40:
	;
	switch v153 - int32(1) {
	case 0:
		v326 = v144
		v327 = v145
		v328 = v149
		goto L24
	case 1:
		v319 = v144
		v320 = v145
		v321 = v149
		goto L25
	case 2:
		v312 = v144
		v313 = v145
		v314 = v149
		goto L26
	case 3:
		v306 = v145
		v307 = v149
		goto L27
	case 4:
		v302 = v145
		v303 = v149
		goto L28
	case 5:
		v296 = v145
		v297 = v149
		goto L29
	case 6:
		v290 = v145
		v291 = v149
		goto L30
	case 7:
		v285 = v149
		goto L31
	case 8:
		v280 = v149
		goto L32
	case 9:
		v275 = v149
		goto L33
	case 10:
		goto L34
	default:
		v334 = v144
		v335 = v145
		v336 = v149
		goto L23
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v115 = v114 + v111
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v119 = v118 + v112
	v121 = int32(4)
	v123 = v116 + v110 - v119 ^ base.I32_rotl(v119, v121)
	v127 = v115 - v123 ^ base.I32_rotl(v123, int32(6))
	v128 = v119 + v115
	v129 = v123 + v128
	v130 = v127 + v129
	v134 = v128 - v127 ^ base.I32_rotl(v127, int32(8))
	v138 = v129 - v134 ^ base.I32_rotl(v134, int32(16))
	v142 = v130 - v138 ^ base.I32_rotl(v138, int32(19))
	v143 = v134 + v130
	v144 = v138 + v143
	v145 = v142 + v144
	v149 = v143 - v142 ^ base.I32_rotl(v142, v121)
	v150 = int32(12)
	v151 = v107 + v150
	v153 = v108 - v150
	if base.Ui32(int32(11)) < base.Ui32(v153) {
		v107 = v151
		v108 = v153
		v110 = v144
		v111 = v145
		v112 = v149
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	goto L42
L44:
	;
	v167 = v36
	v168 = v54
	v170 = v61
	v171 = v61
	v172 = v61
	goto L47
L46:
	;
	switch v213 - int32(1) {
	case 0:
		v268 = v204
		goto L50
	case 1:
		v263 = v204
		goto L51
	case 2:
		goto L52
	case 3:
		v256 = v205
		goto L53
	case 4:
		v253 = v205
		goto L54
	case 5:
		v248 = v205
		goto L55
	case 6:
		goto L56
	case 7:
		v239 = v209
		goto L57
	case 8:
		v234 = v209
		goto L58
	case 9:
		v229 = v209
		goto L59
	case 10:
		goto L60
	default:
		v334 = v204
		v335 = v205
		v336 = v209
		goto L23
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v175 = v174 + v171
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v179 = v178 + v172
	v181 = int32(4)
	v183 = v176 + v170 - v179 ^ base.I32_rotl(v179, v181)
	v187 = v175 - v183 ^ base.I32_rotl(v183, int32(6))
	v188 = v179 + v175
	v189 = v183 + v188
	v190 = v187 + v189
	v194 = v188 - v187 ^ base.I32_rotl(v187, int32(8))
	v198 = v189 - v194 ^ base.I32_rotl(v194, int32(16))
	v202 = v190 - v198 ^ base.I32_rotl(v198, int32(19))
	v203 = v194 + v190
	v204 = v198 + v203
	v205 = v202 + v204
	v209 = v203 - v202 ^ base.I32_rotl(v202, v181)
	v210 = int32(12)
	v211 = v167 + v210
	v213 = v168 - v210
	if base.Ui32(int32(11)) < base.Ui32(v213) {
		v167 = v211
		v168 = v213
		v170 = v204
		v171 = v205
		v172 = v209
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	goto L48
L50:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v334 = v268 + v269
	v335 = v205
	v336 = v209
	goto L23
L51:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	v268 = v264<<(uint(int32(8))%32) + v263
	goto L50
L52:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+2)))
	v263 = v259<<(uint(int32(16))%32) + v204
	goto L51
L53:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v334 = v257 + v204
	v335 = v256
	v336 = v209
	goto L23
L54:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
	v256 = v253 + v254
	goto L53
L55:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+5)))
	v253 = v249<<(uint(int32(8))%32) + v248
	goto L54
L56:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+6)))
	v248 = v244<<(uint(int32(16))%32) + v205
	goto L55
L57:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v334 = v240 + v204
	v335 = v242 + v205
	v336 = v239
	goto L23
L58:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+8)))
	v239 = v235<<(uint(int32(8))%32) + v234
	goto L57
L59:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+9)))
	v234 = v230<<(uint(int32(16))%32) + v229
	goto L58
L60:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+10)))
	v229 = v225<<(uint(int32(24))%32) + v209
	goto L59
L61:
	;
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v373
	v376 = int32(8)
	goto L17
L62:
	;
	v380 = v48
	goto L64
L63:
	;
	v380 = v378
	goto L64
L64:
	;
	if v380 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	base.MemoryCopy(m, v376+v36, v50, v380)
	goto L67
L66:
	;
	goto L67
L67:
	;
	v384 = v376 + v380
	v385 = v48 - v380
	if v385 != 0 {
		v46 = v384
		v48 = v385
		v50 = v380 + v50
		goto L13
	} else {
		goto L68
	}
L68:
	;
	goto L14
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v423 = v416 - int32(1636608432)
	goto L74
L72:
	;
	v733 = F_Int64GetDatum(m, base.I64_extend_i32_u(v723)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v723^v715-base.I32_rotl(v723, int32(24))))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L115
	}
L73:
	;
	if v415&int32(3) != 0 {
		goto L89
	} else {
		goto L90
	}
L74:
	;
	goto L73
L77:
	;
	v701 = int32(14)
	v703 = v697 ^ v698 - base.I32_rotl(v697, v701)
	v707 = v703 ^ v696 - base.I32_rotl(v703, int32(11))
	v711 = v707 ^ v697 - base.I32_rotl(v707, int32(25))
	v715 = v711 ^ v703 - base.I32_rotl(v711, int32(16))
	v719 = v715 ^ v707 - base.I32_rotl(v715, int32(4))
	v723 = v719 ^ v711 - base.I32_rotl(v719, v701)
	goto L72
L78:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	v696 = v688 + v691
	v697 = v689
	v698 = v690
	goto L77
L79:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	v688 = v684<<(uint(int32(8))%32) + v681
	v689 = v682
	v690 = v683
	goto L78
L80:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+2)))
	v681 = v677<<(uint(int32(16))%32) + v674
	v682 = v675
	v683 = v676
	goto L79
L81:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+3)))
	v674 = v670<<(uint(int32(24))%32) + v521
	v675 = v668
	v676 = v669
	goto L80
L82:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+4)))
	v668 = v664 + v666
	v669 = v665
	goto L81
L83:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+5)))
	v664 = v660<<(uint(int32(8))%32) + v658
	v665 = v659
	goto L82
L84:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+6)))
	v658 = v654<<(uint(int32(16))%32) + v652
	v659 = v653
	goto L83
L85:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+7)))
	v652 = v648<<(uint(int32(24))%32) + v522
	v653 = v647
	goto L84
L86:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+8)))
	v647 = v643<<(uint(int32(8))%32) + v642
	goto L85
L87:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+9)))
	v642 = v638<<(uint(int32(16))%32) + v637
	goto L86
L88:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+10)))
	v637 = v633<<(uint(int32(24))%32) + v523
	goto L87
L89:
	;
	if base.Ui32(int32(11)) < base.Ui32(v416) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v416) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	v469 = v415
	v470 = v416
	v472 = v423
	v473 = v423
	v474 = v423
	goto L95
L93:
	;
	v518 = v415
	v519 = v416
	v521 = v423
	v522 = v423
	v523 = v423
	goto L94
L94:
	;
	switch v519 - int32(1) {
	case 0:
		v688 = v521
		v689 = v522
		v690 = v523
		goto L78
	case 1:
		v681 = v521
		v682 = v522
		v683 = v523
		goto L79
	case 2:
		v674 = v521
		v675 = v522
		v676 = v523
		goto L80
	case 3:
		v668 = v522
		v669 = v523
		goto L81
	case 4:
		v664 = v522
		v665 = v523
		goto L82
	case 5:
		v658 = v522
		v659 = v523
		goto L83
	case 6:
		v652 = v522
		v653 = v523
		goto L84
	case 7:
		v647 = v523
		goto L85
	case 8:
		v642 = v523
		goto L86
	case 9:
		v637 = v523
		goto L87
	case 10:
		goto L88
	default:
		v696 = v521
		v697 = v522
		v698 = v523
		goto L77
	}
L95:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v477 = v476 + v473
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v469)+8))
	v481 = v480 + v474
	v483 = int32(4)
	v485 = v478 + v472 - v481 ^ base.I32_rotl(v481, v483)
	v489 = v477 - v485 ^ base.I32_rotl(v485, int32(6))
	v490 = v481 + v477
	v491 = v485 + v490
	v492 = v489 + v491
	v496 = v490 - v489 ^ base.I32_rotl(v489, int32(8))
	v500 = v491 - v496 ^ base.I32_rotl(v496, int32(16))
	v504 = v492 - v500 ^ base.I32_rotl(v500, int32(19))
	v505 = v496 + v492
	v506 = v500 + v505
	v507 = v504 + v506
	v511 = v505 - v504 ^ base.I32_rotl(v504, v483)
	v512 = int32(12)
	v513 = v469 + v512
	v515 = v470 - v512
	if base.Ui32(int32(11)) < base.Ui32(v515) {
		v469 = v513
		v470 = v515
		v472 = v506
		v473 = v507
		v474 = v511
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v518 = v513
	v519 = v515
	v521 = v506
	v522 = v507
	v523 = v511
	goto L94
L97:
	;
	goto L96
L98:
	;
	v529 = v415
	v530 = v416
	v532 = v423
	v533 = v423
	v534 = v423
	goto L101
L99:
	;
	v578 = v415
	v579 = v416
	v581 = v423
	v582 = v423
	v583 = v423
	goto L100
L100:
	;
	switch v579 - int32(1) {
	case 0:
		v630 = v581
		goto L104
	case 1:
		v625 = v581
		goto L105
	case 2:
		goto L106
	case 3:
		v618 = v582
		goto L107
	case 4:
		v615 = v582
		goto L108
	case 5:
		v610 = v582
		goto L109
	case 6:
		goto L110
	case 7:
		v601 = v583
		goto L111
	case 8:
		v596 = v583
		goto L112
	case 9:
		v591 = v583
		goto L113
	case 10:
		goto L114
	default:
		v696 = v581
		v697 = v582
		v698 = v583
		goto L77
	}
L101:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v529)+4))
	v537 = v536 + v533
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v529)+8))
	v541 = v540 + v534
	v543 = int32(4)
	v545 = v538 + v532 - v541 ^ base.I32_rotl(v541, v543)
	v549 = v537 - v545 ^ base.I32_rotl(v545, int32(6))
	v550 = v541 + v537
	v551 = v545 + v550
	v552 = v549 + v551
	v556 = v550 - v549 ^ base.I32_rotl(v549, int32(8))
	v560 = v551 - v556 ^ base.I32_rotl(v556, int32(16))
	v564 = v552 - v560 ^ base.I32_rotl(v560, int32(19))
	v565 = v556 + v552
	v566 = v560 + v565
	v567 = v564 + v566
	v571 = v565 - v564 ^ base.I32_rotl(v564, v543)
	v572 = int32(12)
	v573 = v529 + v572
	v575 = v530 - v572
	if base.Ui32(int32(11)) < base.Ui32(v575) {
		v529 = v573
		v530 = v575
		v532 = v566
		v533 = v567
		v534 = v571
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v578 = v573
	v579 = v575
	v581 = v566
	v582 = v567
	v583 = v571
	goto L100
L103:
	;
	goto L102
L104:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v696 = v630 + v631
	v697 = v582
	v698 = v583
	goto L77
L105:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+1)))
	v630 = v626<<(uint(int32(8))%32) + v625
	goto L104
L106:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+2)))
	v625 = v621<<(uint(int32(16))%32) + v581
	goto L105
L107:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	v696 = v619 + v581
	v697 = v618
	v698 = v583
	goto L77
L108:
	;
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+4)))
	v618 = v615 + v616
	goto L107
L109:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+5)))
	v615 = v611<<(uint(int32(8))%32) + v610
	goto L108
L110:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+6)))
	v610 = v606<<(uint(int32(16))%32) + v582
	goto L109
L111:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v696 = v602 + v581
	v697 = v604 + v582
	v698 = v601
	goto L77
L112:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+8)))
	v601 = v597<<(uint(int32(8))%32) + v596
	goto L111
L113:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+9)))
	v596 = v592<<(uint(int32(16))%32) + v591
	goto L112
L114:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+10)))
	v591 = v587<<(uint(int32(24))%32) + v583
	goto L113
L115:
	;
	v735 = *(*int64)(unsafe.Add(mBase, uint32(v733)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v735
	if v735 == int64(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v741 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	return v10
L119:
	;
	v742 = int64(2)
	goto L121
L120:
	;
	v742 = int64(1)
	goto L121
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v742
	goto L118
}
func F__jumbleAlterUserMappingStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v7 != 0 {
			v8 = F_strlen(m, v7)
			mBase = m.M
			F_AppendJumble(m, l0, v7, v8+int32(1))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v13 + int32(1)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleArrayCoerceExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_AppendJumble32(m, l0, l1+int32(12))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleBoolExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		F_AppendJumble(m, l0, v4, v5+int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_AppendJumble8(m, l0, l1+int32(8))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10 + int32(1)
		F_AppendJumble8(m, l0, l1+int32(8))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleCreateSchemaStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		F_AppendJumble(m, l0, v4, v5+int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			F__jumbleNode(m, l0, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_AppendJumble8(m, l0, l1+int32(16))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_AppendJumble8(m, l0, l1+int32(16))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleFieldStore(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleFuncExpr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		F__jumbleNode(m, l0, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F__jumbleJsonObjectConstructor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F__jumbleNode(m, l0, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_AppendJumble8(m, l0, l1+int32(12))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_AppendJumble8(m, l0, l1+int32(13))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__jumbleWithClause(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F__jumbleNode(m, l0, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_AppendJumble8(m, l0, l1+int32(8))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
