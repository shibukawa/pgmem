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
	var v28 int32
	_ = v28
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
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
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
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int64
	_ = v722
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
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(1024)-v384) < base.Ui32(l2) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v384 = v12
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v373
	v384 = v373
	goto L1
L6:
	;
	v25 = v15
	v28 = v13
	v29 = l0 + int32(28)
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v14))) = v9
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v373 = v367 + int32(4)
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
	v373 = v363
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
	v357 = int32(1024) - v354
	if base.Ui32(v28) < base.Ui32(v357) {
		goto L59
	} else {
		goto L60
	}
L14:
	;
	v349 = F_Int64GetDatum(m, base.I64_extend_i32_u(v339)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v331^v339-base.I32_rotl(v339, int32(24))))
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
	v359 = v28
	goto L61
L60:
	;
	v359 = v357
	goto L61
L61:
	;
	if v359 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v363 = v354 + v359
	v364 = v28 - v359
	if v364 != 0 {
		v25 = v363
		v28 = v364
		v29 = v359 + v29
		goto L9
	} else {
		goto L66
	}
L63:
	;
	v360 = F__emscripten_memcpy_bulkmem(m, v354+v14, v29, v359)
	mBase = m.M
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L10
L67:
	;
	v394 = l1
	v395 = l2
	v396 = v384
	goto L70
L68:
	;
	goto L69
L69:
	;
	if l2 != 0 {
		goto L128
	} else {
		goto L129
	}
L70:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v396) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v734
	return
L72:
	;
	v403 = int32(1024)
	v410 = int32(-1636607408)
	goto L77
L73:
	;
	v725 = v396
	goto L74
L74:
	;
	v728 = int32(1024) - v725
	if base.Ui32(v395) < base.Ui32(v728) {
		goto L119
	} else {
		goto L120
	}
L75:
	;
	v720 = F_Int64GetDatum(m, base.I64_extend_i32_u(v710)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v702^v710-base.I32_rotl(v710, int32(24))))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L57
	} else {
		goto L118
	}
L76:
	;
	if v389&int32(3) != 0 {
		goto L92
	} else {
		goto L93
	}
L77:
	;
	goto L76
L80:
	;
	v688 = int32(14)
	v690 = v684 ^ v685 - base.I32_rotl(v684, v688)
	v694 = v690 ^ v683 - base.I32_rotl(v690, int32(11))
	v698 = v694 ^ v684 - base.I32_rotl(v694, int32(25))
	v702 = v698 ^ v690 - base.I32_rotl(v698, int32(16))
	v706 = v702 ^ v694 - base.I32_rotl(v702, int32(4))
	v710 = v706 ^ v698 - base.I32_rotl(v706, v688)
	goto L75
L81:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	v683 = v675 + v678
	v684 = v676
	v685 = v677
	goto L80
L82:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+1)))
	v675 = v671<<(uint(int32(8))%32) + v668
	v676 = v669
	v677 = v670
	goto L81
L83:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+2)))
	v668 = v664<<(uint(int32(16))%32) + v661
	v669 = v662
	v670 = v663
	goto L82
L84:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+3)))
	v661 = v657<<(uint(int32(24))%32) + v493
	v662 = v655
	v663 = v656
	goto L83
L85:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+4)))
	v655 = v651 + v653
	v656 = v652
	goto L84
L86:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+5)))
	v651 = v647<<(uint(int32(8))%32) + v645
	v652 = v646
	goto L85
L87:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+6)))
	v645 = v641<<(uint(int32(16))%32) + v639
	v646 = v640
	goto L86
L88:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+7)))
	v639 = v635<<(uint(int32(24))%32) + v494
	v640 = v634
	goto L87
L89:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+8)))
	v634 = v630<<(uint(int32(8))%32) + v629
	goto L88
L90:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+9)))
	v629 = v625<<(uint(int32(16))%32) + v624
	goto L89
L91:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+10)))
	v624 = v620<<(uint(int32(24))%32) + v498
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
	v456 = v389
	v457 = v403
	v459 = v410
	v460 = v410
	v461 = v410
	goto L98
L97:
	;
	switch v502 - int32(1) {
	case 0:
		v675 = v493
		v676 = v494
		v677 = v498
		goto L81
	case 1:
		v668 = v493
		v669 = v494
		v670 = v498
		goto L82
	case 2:
		v661 = v493
		v662 = v494
		v663 = v498
		goto L83
	case 3:
		v655 = v494
		v656 = v498
		goto L84
	case 4:
		v651 = v494
		v652 = v498
		goto L85
	case 5:
		v645 = v494
		v646 = v498
		goto L86
	case 6:
		v639 = v494
		v640 = v498
		goto L87
	case 7:
		v634 = v498
		goto L88
	case 8:
		v629 = v498
		goto L89
	case 9:
		v624 = v498
		goto L90
	case 10:
		goto L91
	default:
		v683 = v493
		v684 = v494
		v685 = v498
		goto L80
	}
L98:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	v464 = v463 + v460
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	v468 = v467 + v461
	v470 = int32(4)
	v472 = v465 + v459 - v468 ^ base.I32_rotl(v468, v470)
	v476 = v464 - v472 ^ base.I32_rotl(v472, int32(6))
	v477 = v468 + v464
	v478 = v472 + v477
	v479 = v476 + v478
	v483 = v477 - v476 ^ base.I32_rotl(v476, int32(8))
	v487 = v478 - v483 ^ base.I32_rotl(v483, int32(16))
	v491 = v479 - v487 ^ base.I32_rotl(v487, int32(19))
	v492 = v483 + v479
	v493 = v487 + v492
	v494 = v491 + v493
	v498 = v492 - v491 ^ base.I32_rotl(v491, v470)
	v499 = int32(12)
	v500 = v456 + v499
	v502 = v457 - v499
	if base.Ui32(int32(11)) < base.Ui32(v502) {
		v456 = v500
		v457 = v502
		v459 = v493
		v460 = v494
		v461 = v498
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
	v516 = v389
	v517 = v403
	v519 = v410
	v520 = v410
	v521 = v410
	goto L104
L103:
	;
	switch v562 - int32(1) {
	case 0:
		v617 = v553
		goto L107
	case 1:
		v612 = v553
		goto L108
	case 2:
		goto L109
	case 3:
		v605 = v554
		goto L110
	case 4:
		v602 = v554
		goto L111
	case 5:
		v597 = v554
		goto L112
	case 6:
		goto L113
	case 7:
		v588 = v558
		goto L114
	case 8:
		v583 = v558
		goto L115
	case 9:
		v578 = v558
		goto L116
	case 10:
		goto L117
	default:
		v683 = v553
		v684 = v554
		v685 = v558
		goto L80
	}
L104:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v524 = v523 + v520
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	v528 = v527 + v521
	v530 = int32(4)
	v532 = v525 + v519 - v528 ^ base.I32_rotl(v528, v530)
	v536 = v524 - v532 ^ base.I32_rotl(v532, int32(6))
	v537 = v528 + v524
	v538 = v532 + v537
	v539 = v536 + v538
	v543 = v537 - v536 ^ base.I32_rotl(v536, int32(8))
	v547 = v538 - v543 ^ base.I32_rotl(v543, int32(16))
	v551 = v539 - v547 ^ base.I32_rotl(v547, int32(19))
	v552 = v543 + v539
	v553 = v547 + v552
	v554 = v551 + v553
	v558 = v552 - v551 ^ base.I32_rotl(v551, v530)
	v559 = int32(12)
	v560 = v516 + v559
	v562 = v517 - v559
	if base.Ui32(int32(11)) < base.Ui32(v562) {
		v516 = v560
		v517 = v562
		v519 = v553
		v520 = v554
		v521 = v558
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
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	v683 = v617 + v618
	v684 = v554
	v685 = v558
	goto L80
L108:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+1)))
	v617 = v613<<(uint(int32(8))%32) + v612
	goto L107
L109:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+2)))
	v612 = v608<<(uint(int32(16))%32) + v553
	goto L108
L110:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v683 = v606 + v553
	v684 = v605
	v685 = v558
	goto L80
L111:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+4)))
	v605 = v602 + v603
	goto L110
L112:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+5)))
	v602 = v598<<(uint(int32(8))%32) + v597
	goto L111
L113:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+6)))
	v597 = v593<<(uint(int32(16))%32) + v554
	goto L112
L114:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v683 = v589 + v553
	v684 = v591 + v554
	v685 = v588
	goto L80
L115:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+8)))
	v588 = v584<<(uint(int32(8))%32) + v583
	goto L114
L116:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+9)))
	v583 = v579<<(uint(int32(16))%32) + v578
	goto L115
L117:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+10)))
	v578 = v574<<(uint(int32(24))%32) + v558
	goto L116
L118:
	;
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v720)))
	*(*int64)(unsafe.Add(mBase, uint32(v389))) = v722
	v725 = int32(8)
	goto L74
L119:
	;
	v730 = v395
	goto L121
L120:
	;
	v730 = v728
	goto L121
L121:
	;
	if v730 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v734 = v725 + v730
	v735 = v395 - v730
	if v735 != 0 {
		v394 = v394 + v730
		v395 = v735
		v396 = v734
		goto L70
	} else {
		goto L126
	}
L123:
	;
	v731 = F__emscripten_memcpy_bulkmem(m, v725+v389, v394, v730)
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	goto L71
L127:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v740 + l2
	return
L128:
	;
	v738 = F__emscripten_memcpy_bulkmem(m, v384+v389, l1, l2)
	mBase = m.M
	goto L130
L129:
	;
	goto L130
L130:
	;
	goto L127
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
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
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
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
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
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
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
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
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
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int64
	_ = v736
	var v742 int32
	_ = v742
	var v743 int64
	_ = v743
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
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
	if v411 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v394
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
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v394 = v389 + int32(4)
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
	v394 = v385
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
	v379 = int32(1024) - v376
	if base.Ui32(v48) < base.Ui32(v379) {
		goto L62
	} else {
		goto L63
	}
L18:
	;
	v371 = F_Int64GetDatum(m, base.I64_extend_i32_u(v361)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v353^v361-base.I32_rotl(v361, int32(24))))
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
	v381 = v48
	goto L64
L63:
	;
	v381 = v379
	goto L64
L64:
	;
	if v381 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v385 = v376 + v381
	v386 = v48 - v381
	if v386 != 0 {
		v46 = v385
		v48 = v386
		v50 = v381 + v50
		goto L13
	} else {
		goto L69
	}
L66:
	;
	v382 = F__emscripten_memcpy_bulkmem(m, v376+v36, v50, v381)
	mBase = m.M
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	goto L14
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(0)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v424 = v417 - int32(1636608432)
	goto L75
L73:
	;
	v734 = F_Int64GetDatum(m, base.I64_extend_i32_u(v724)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v716^v724-base.I32_rotl(v724, int32(24))))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L116
	}
L74:
	;
	if v416&int32(3) != 0 {
		goto L90
	} else {
		goto L91
	}
L75:
	;
	goto L74
L78:
	;
	v702 = int32(14)
	v704 = v698 ^ v699 - base.I32_rotl(v698, v702)
	v708 = v704 ^ v697 - base.I32_rotl(v704, int32(11))
	v712 = v708 ^ v698 - base.I32_rotl(v708, int32(25))
	v716 = v712 ^ v704 - base.I32_rotl(v712, int32(16))
	v720 = v716 ^ v708 - base.I32_rotl(v716, int32(4))
	v724 = v720 ^ v712 - base.I32_rotl(v720, v702)
	goto L73
L79:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	v697 = v689 + v692
	v698 = v690
	v699 = v691
	goto L78
L80:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+1)))
	v689 = v685<<(uint(int32(8))%32) + v682
	v690 = v683
	v691 = v684
	goto L79
L81:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+2)))
	v682 = v678<<(uint(int32(16))%32) + v675
	v683 = v676
	v684 = v677
	goto L80
L82:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+3)))
	v675 = v671<<(uint(int32(24))%32) + v522
	v676 = v669
	v677 = v670
	goto L81
L83:
	;
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+4)))
	v669 = v665 + v667
	v670 = v666
	goto L82
L84:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+5)))
	v665 = v661<<(uint(int32(8))%32) + v659
	v666 = v660
	goto L83
L85:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+6)))
	v659 = v655<<(uint(int32(16))%32) + v653
	v660 = v654
	goto L84
L86:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+7)))
	v653 = v649<<(uint(int32(24))%32) + v523
	v654 = v648
	goto L85
L87:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+8)))
	v648 = v644<<(uint(int32(8))%32) + v643
	goto L86
L88:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+9)))
	v643 = v639<<(uint(int32(16))%32) + v638
	goto L87
L89:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+10)))
	v638 = v634<<(uint(int32(24))%32) + v524
	goto L88
L90:
	;
	if base.Ui32(int32(11)) < base.Ui32(v417) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v417) {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	v470 = v416
	v471 = v417
	v473 = v424
	v474 = v424
	v475 = v424
	goto L96
L94:
	;
	v519 = v416
	v520 = v417
	v522 = v424
	v523 = v424
	v524 = v424
	goto L95
L95:
	;
	switch v520 - int32(1) {
	case 0:
		v689 = v522
		v690 = v523
		v691 = v524
		goto L79
	case 1:
		v682 = v522
		v683 = v523
		v684 = v524
		goto L80
	case 2:
		v675 = v522
		v676 = v523
		v677 = v524
		goto L81
	case 3:
		v669 = v523
		v670 = v524
		goto L82
	case 4:
		v665 = v523
		v666 = v524
		goto L83
	case 5:
		v659 = v523
		v660 = v524
		goto L84
	case 6:
		v653 = v523
		v654 = v524
		goto L85
	case 7:
		v648 = v524
		goto L86
	case 8:
		v643 = v524
		goto L87
	case 9:
		v638 = v524
		goto L88
	case 10:
		goto L89
	default:
		v697 = v522
		v698 = v523
		v699 = v524
		goto L78
	}
L96:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	v478 = v477 + v474
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v470)+8))
	v482 = v481 + v475
	v484 = int32(4)
	v486 = v479 + v473 - v482 ^ base.I32_rotl(v482, v484)
	v490 = v478 - v486 ^ base.I32_rotl(v486, int32(6))
	v491 = v482 + v478
	v492 = v486 + v491
	v493 = v490 + v492
	v497 = v491 - v490 ^ base.I32_rotl(v490, int32(8))
	v501 = v492 - v497 ^ base.I32_rotl(v497, int32(16))
	v505 = v493 - v501 ^ base.I32_rotl(v501, int32(19))
	v506 = v497 + v493
	v507 = v501 + v506
	v508 = v505 + v507
	v512 = v506 - v505 ^ base.I32_rotl(v505, v484)
	v513 = int32(12)
	v514 = v470 + v513
	v516 = v471 - v513
	if base.Ui32(int32(11)) < base.Ui32(v516) {
		v470 = v514
		v471 = v516
		v473 = v507
		v474 = v508
		v475 = v512
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v519 = v514
	v520 = v516
	v522 = v507
	v523 = v508
	v524 = v512
	goto L95
L98:
	;
	goto L97
L99:
	;
	v530 = v416
	v531 = v417
	v533 = v424
	v534 = v424
	v535 = v424
	goto L102
L100:
	;
	v579 = v416
	v580 = v417
	v582 = v424
	v583 = v424
	v584 = v424
	goto L101
L101:
	;
	switch v580 - int32(1) {
	case 0:
		v631 = v582
		goto L105
	case 1:
		v626 = v582
		goto L106
	case 2:
		goto L107
	case 3:
		v619 = v583
		goto L108
	case 4:
		v616 = v583
		goto L109
	case 5:
		v611 = v583
		goto L110
	case 6:
		goto L111
	case 7:
		v602 = v584
		goto L112
	case 8:
		v597 = v584
		goto L113
	case 9:
		v592 = v584
		goto L114
	case 10:
		goto L115
	default:
		v697 = v582
		v698 = v583
		v699 = v584
		goto L78
	}
L102:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	v538 = v537 + v534
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v530)+8))
	v542 = v541 + v535
	v544 = int32(4)
	v546 = v539 + v533 - v542 ^ base.I32_rotl(v542, v544)
	v550 = v538 - v546 ^ base.I32_rotl(v546, int32(6))
	v551 = v542 + v538
	v552 = v546 + v551
	v553 = v550 + v552
	v557 = v551 - v550 ^ base.I32_rotl(v550, int32(8))
	v561 = v552 - v557 ^ base.I32_rotl(v557, int32(16))
	v565 = v553 - v561 ^ base.I32_rotl(v561, int32(19))
	v566 = v557 + v553
	v567 = v561 + v566
	v568 = v565 + v567
	v572 = v566 - v565 ^ base.I32_rotl(v565, v544)
	v573 = int32(12)
	v574 = v530 + v573
	v576 = v531 - v573
	if base.Ui32(int32(11)) < base.Ui32(v576) {
		v530 = v574
		v531 = v576
		v533 = v567
		v534 = v568
		v535 = v572
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v579 = v574
	v580 = v576
	v582 = v567
	v583 = v568
	v584 = v572
	goto L101
L104:
	;
	goto L103
L105:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579))))
	v697 = v631 + v632
	v698 = v583
	v699 = v584
	goto L78
L106:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+1)))
	v631 = v627<<(uint(int32(8))%32) + v626
	goto L105
L107:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+2)))
	v626 = v622<<(uint(int32(16))%32) + v582
	goto L106
L108:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v697 = v620 + v582
	v698 = v619
	v699 = v584
	goto L78
L109:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+4)))
	v619 = v616 + v617
	goto L108
L110:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+5)))
	v616 = v612<<(uint(int32(8))%32) + v611
	goto L109
L111:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+6)))
	v611 = v607<<(uint(int32(16))%32) + v583
	goto L110
L112:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	v697 = v603 + v582
	v698 = v605 + v583
	v699 = v602
	goto L78
L113:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+8)))
	v602 = v598<<(uint(int32(8))%32) + v597
	goto L112
L114:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+9)))
	v597 = v593<<(uint(int32(16))%32) + v592
	goto L113
L115:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+10)))
	v592 = v588<<(uint(int32(24))%32) + v584
	goto L114
L116:
	;
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v734)))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v736
	if v736 == int64(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v742 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	return v10
L120:
	;
	v743 = int64(2)
	goto L122
L121:
	;
	v743 = int64(1)
	goto L122
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v743
	goto L119
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
