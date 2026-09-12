package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expandRTE(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
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
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	v9 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	if l6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v30 {
	case 0:
		goto L13
	case 1:
		goto L16
	case 2:
		goto L10
	case 3:
		goto L15
	case 4, 5, 6, 7:
		goto L11
	case 8, 9:
		goto L7
	default:
		goto L12
	}
L7:
	;
	m.G0 = v24 + int32(32)
	return
L8:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v643 == int32(0) {
		goto L7
	} else {
		goto L190
	}
L9:
	;
	v630 = v343 + int32(1)
	goto L8
L10:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	v523 = v9
	goto L151
L11:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	if v393 != 0 {
		goto L111
	} else {
		goto L112
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L31
	} else {
		goto L108
	}
L13:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v368 = F_relation_open(m, v366, int32(1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L31
	} else {
		goto L105
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L31
	} else {
		goto L102
	}
L15:
	;
	v125 = int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v126 == int32(0) {
		v630 = v125
		goto L8
	} else {
		goto L47
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v32 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v35 = v33
	goto L19
L18:
	;
	v35 = int32(0)
	goto L19
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	if v37 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 <= int32(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v52 = int32(0)
	v55 = v35
	v56 = v9
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v52<<(uint(int32(2))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+26)))
	if v70 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L7
L24:
	;
	if v55 == int32(0) {
		goto L14
	} else {
		goto L27
	}
L25:
	;
	v119 = v55
	v120 = v56
	goto L26
L26:
	;
	v122 = v52 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v122 < v123 {
		v52 = v122
		v55 = v119
		v56 = v120
		goto L22
	} else {
		goto L46
	}
L27:
	;
	if l6 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v78 = F_pstrdup(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v86 = v56 + int32(1)
	if l7 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	return
L32:
	;
	v80 = F_makeString(m, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v82 = F_lappend(m, v75, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v82
	goto L30
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v89 = F_exprType(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v107 = v55 + int32(4)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if base.Ui32(v107) < base.Ui32(v111+v112<<(uint(int32(2))%32)) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v92 = F_exprTypmod(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v95 = F_exprCollation(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v97 = F_makeVar(m, l1, base.I32_extend16_s(v86), v89, v92, v95, l2)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = l3
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v102 = F_lappend(m, v101, v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v102
	goto L37
L43:
	;
	v117 = v107
	goto L45
L44:
	;
	v117 = int32(0)
	goto L45
L45:
	;
	v119 = v117
	v120 = v86
	goto L26
L46:
	;
	goto L23
L47:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v129 <= int32(0) {
		v630 = v125
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v148 = v9
	v152 = v9
	goto L49
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v152<<(uint(int32(2))%32))))
	v158 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v158
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	if v162 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L9
L51:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v343 = v342 + v148
	v345 = v152 + int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v345 < v346 {
		v148 = v343
		v152 = v345
		goto L49
	} else {
		goto L101
	}
L52:
	;
	if l6 != 0 {
		goto L73
	} else {
		goto L74
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v168 = F_get_expr_result_type(m, v163, v24+int32(28), v24+int32(24))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	v170 = int32(1)
	if base.Ui32(v168-v170) <= base.Ui32(v170) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	F_expandTupleDesc(m, v174, v175, v176, v148, l1, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L31
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	switch v168 {
	case 0:
		goto L60
	default:
		goto L59
	case 3:
		goto L52
	}
L58:
	;
	goto L51
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L31
	} else {
		goto L70
	}
L60:
	;
	if l6 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v148<<(uint(int32(2))%32))))
	v187 = F_lappend(m, v179, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L31
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if l7 == int32(0) {
		goto L51
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v187
	goto L63
L65:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v197 = F_exprTypmod(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v200 = F_exprCollation(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	v202 = F_makeVar(m, l1, base.I32_extend16_s(v148+int32(1)), v195, v197, v200, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v202)+32)) = l3
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v207 = F_lappend(m, v206, v202)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v207
	goto L51
L70:
	;
	F_errmsg_internal(m, int32(350950), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(473353), int32(2926), int32(514103))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v226 = F_list_copy_tail(m, v225, v148)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L31
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if l7 == int32(0) {
		goto L51
	} else {
		goto L85
	}
L76:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v229 = int32(0)
	if v226 == v229 {
		v237 = v229
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v239 = F_list_concat(m, v238, v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L31
	} else {
		goto L84
	}
L78:
	;
	goto L77
L79:
	;
	if v228 <= int32(0) {
		v237 = v229
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v228 < v234 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v228
	goto L83
L82:
	;
	goto L83
L83:
	;
	v237 = v226
	goto L78
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v239
	goto L75
L85:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v257 = int32(0)
	v267 = v148
	goto L86
L86:
	;
	v270 = int32(0)
	if v247 == v270 {
		v280 = v270
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v281 = int32(0)
	if v246 == v281 {
		v290 = v281
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if v274 <= v257 {
		v280 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v280 = v276 + v257<<(uint(int32(2))%32)
	goto L88
L91:
	;
	if v245 == int32(0) {
		goto L51
	} else {
		goto L94
	}
L92:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v284 <= v257 {
		v290 = v281
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v290 = v286 + v257<<(uint(int32(2))%32)
	goto L91
L94:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v293 <= v257 {
		goto L51
	} else {
		goto L95
	}
L95:
	;
	if v280 == int32(0) {
		goto L51
	} else {
		goto L96
	}
L96:
	;
	if v290 == int32(0) {
		goto L51
	} else {
		goto L97
	}
L97:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v245)+12))
	v302 = v299 + v257<<(uint(int32(2))%32)
	if v302 == int32(0) {
		goto L51
	} else {
		goto L98
	}
L98:
	;
	v306 = v267 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v311 = F_makeVar(m, l1, base.I32_extend16_s(v306), v308, v309, v310, l2)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L31
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v311)+32)) = l3
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v316 = F_lappend(m, v315, v311)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L31
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v316
	v257 = v257 + int32(1)
	v267 = v306
	goto L86
L101:
	;
	goto L50
L102:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v353
	F_errmsg_internal(m, int32(167514), v24+int32(16))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L31
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(473353), int32(2797), int32(514103))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L31
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v368)+52))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	F_expandTupleDesc(m, v370, v365, v371, int32(0), l1, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L31
	} else {
		goto L106
	}
L106:
	;
	F_relation_close(m, v368, int32(1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L31
	} else {
		goto L107
	}
L107:
	;
	goto L7
L108:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v382
	F_errmsg_internal(m, int32(465096), v24)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L31
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(473353), int32(3110), int32(514103))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L31
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v395 = v394
	goto L113
L112:
	;
	v395 = v9
	goto L113
L113:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v408 = int32(0)
	v415 = v395
	goto L114
L114:
	;
	v421 = int32(0)
	if v398 == v421 {
		v432 = v421
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v397 == int32(0) {
		v441 = v421
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v426 <= v408 {
		v432 = int32(0)
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	v432 = v428 + v408<<(uint(int32(2))%32)
	goto L116
L119:
	;
	if v396 == int32(0) {
		goto L7
	} else {
		goto L122
	}
L120:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v435 <= v408 {
		v441 = v421
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v441 = v437 + v408<<(uint(int32(2))%32)
	goto L119
L122:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v444 <= v408 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	if v432 == int32(0) {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	if v441 == int32(0) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v453 = v450 + v408<<(uint(int32(2))%32)
	if v453 == int32(0) {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if l6 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v458 != 0 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v487 = v415
	goto L129
L129:
	;
	v489 = v408 + int32(1)
	if l7 == int32(0) {
		v408 = v489
		v415 = v487
		goto L114
	} else {
		goto L142
	}
L130:
	;
	v475 = v415 + int32(4)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if base.Ui32(v475) < base.Ui32(v479+v480<<(uint(int32(2))%32)) {
		goto L139
	} else {
		goto L140
	}
L131:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v466 = F_pstrdup(m, v464)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L31
	} else {
		goto L136
	}
L132:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v464 = v460
	goto L131
L133:
	;
	goto L134
L134:
	;
	if l5 == int32(0) {
		goto L130
	} else {
		goto L135
	}
L135:
	;
	v464 = int32(715480)
	goto L131
L136:
	;
	v468 = F_makeString(m, v466)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L31
	} else {
		goto L137
	}
L137:
	;
	v470 = F_lappend(m, v465, v468)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L31
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v470
	goto L130
L139:
	;
	v485 = v475
	goto L141
L140:
	;
	v485 = int32(0)
	goto L141
L141:
	;
	v487 = v485
	goto L129
L142:
	;
	if v458 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v493 = F_makeVar(m, l1, base.I32_extend16_s(v489), v458, v457, v456, l2)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L31
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if l5 == int32(0) {
		v408 = v489
		v415 = v487
		goto L114
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v493)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v493)+32)) = l3
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v498 = F_lappend(m, v497, v493)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L31
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v498
	v408 = v489
	v415 = v487
	goto L114
L148:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v507 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L31
	} else {
		goto L149
	}
L149:
	;
	v509 = F_lappend(m, v503, v507)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L31
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v509
	v408 = v489
	v415 = v487
	goto L114
L151:
	;
	v536 = int32(0)
	if v514 == v536 {
		v546 = v536
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if v512 == int32(0) {
		goto L7
	} else {
		goto L156
	}
L154:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v540 <= v523 {
		v546 = int32(0)
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v514)+12))
	v546 = v542 + v523<<(uint(int32(2))%32)
	goto L153
L156:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v549 <= v523 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	if v546 == int32(0) {
		goto L7
	} else {
		goto L158
	}
L158:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v512)+12))
	v556 = v553 + v523<<(uint(int32(2))%32)
	if v556 == int32(0) {
		goto L7
	} else {
		goto L159
	}
L159:
	;
	v560 = v523 + int32(1)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	if v561 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if l5 == int32(0) {
		v523 = v560
		goto L151
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if l6 != 0 {
		goto L173
	} else {
		goto L174
	}
L163:
	;
	if l6 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v568 = F_pstrdup(m, int32(715480))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L31
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if l7 == int32(0) {
		v523 = v560
		goto L151
	} else {
		goto L170
	}
L167:
	;
	v570 = F_makeString(m, v568)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L31
	} else {
		goto L168
	}
L168:
	;
	v572 = F_lappend(m, v566, v570)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L31
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v572
	goto L166
L170:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v581 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L31
	} else {
		goto L171
	}
L171:
	;
	v583 = F_lappend(m, v577, v581)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L31
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v583
	v523 = v560
	goto L151
L173:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	v589 = F_pstrdup(m, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L31
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if l7 == int32(0) {
		v523 = v560
		goto L151
	} else {
		goto L179
	}
L176:
	;
	v591 = F_makeString(m, v589)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L31
	} else {
		goto L177
	}
L177:
	;
	v593 = F_lappend(m, v586, v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L31
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v593
	goto L175
L179:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v561)))
	if v598 == int32(6) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v613)+32)) = l3
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v617 = F_lappend(m, v616, v613)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L31
	} else {
		goto L189
	}
L181:
	;
	v601 = F_copyObjectImpl(m, v561)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L31
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v605 = F_exprType(m, v561)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L31
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+28)) = l2
	v613 = v601
	goto L180
L185:
	;
	v607 = F_exprTypmod(m, v561)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L31
	} else {
		goto L186
	}
L186:
	;
	v609 = F_exprCollation(m, v561)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L31
	} else {
		goto L187
	}
L187:
	;
	v611 = F_makeVar(m, l1, base.I32_extend16_s(v560), v605, v607, v609, l2)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L31
	} else {
		goto L188
	}
L188:
	;
	v613 = v611
	goto L180
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v617
	v523 = v560
	goto L151
L190:
	;
	if l6 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+8))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+12))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v649+v650<<(uint(int32(2))%32)-int32(4))))
	v657 = F_lappend(m, v646, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L31
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	if l7 == int32(0) {
		goto L7
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v657
	goto L193
L195:
	;
	v667 = F_makeVar(m, l1, base.I32_extend16_s(v630), int32(20), int32(-1), int32(0), l2)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L31
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+32)) = l3
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v671 = F_lappend(m, v670, v667)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L31
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v671
	goto L7
}
func F_get_rte_alias(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = l1 << (uint(int32(2)) % 32)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13+v18-int32(4))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L34
	}
L2:
	;
	m.G0 = v10 + int32(32)
	return
L3:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v93 = F_quote_identifier(m, v22)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L29
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v13-int32(4))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)))
	if v30 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v31 {
	case 0:
		goto L7
	case 1, 3, 5:
		goto L3
	default:
		goto L2
	case 6:
		goto L6
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v33 = F_get_rel_name(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v33 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v40 == int32(0) {
		v59 = v39
		v60 = v40
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v60-v59 != 0 {
		goto L3
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	if v39 != v40 {
		v59 = v39
		v60 = v40
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v44 = v22
	v45 = v33
	goto L15
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v49 == int32(0) {
		v59 = v48
		v60 = v49
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v59 = v48
	v60 = v49
	goto L12
L17:
	;
	v52 = int32(1)
	if v48 == v49 {
		v44 = v44 + v52
		v45 = v45 + v52
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L2
L20:
	;
	if v86-v85 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v70 = v22
	v71 = v62
	goto L24
L24:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v85 = v74
	v86 = v75
	goto L21
L26:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L3
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v93
	if l2 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v98 = int32(702482)
	goto L32
L31:
	;
	v98 = int32(704419)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v98
	F_appendStringInfo(m, v92, int32(165743), v10)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v32
	F_errmsg_internal(m, int32(43821), v10+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(471101), int32(13138), int32(360234))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_rte_attribute_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		v57 = int32(628769)
		m.G0 = v7 + int32(16)
		return v57
	} else {
		if l1 <= int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v31 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v36 = F_get_attname(m, v34, l1, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v57 = v36
					m.G0 = v7 + int32(16)
					return v57
				}
			} else {
				if l1 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg_internal(m, int32(167472), v7)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(473353), int32(3387), int32(360618))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
					if v43 != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
						v46 = v44
					} else {
						v46 = int32(0)
					}
					if v46 < l1 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errmsg_internal(m, int32(167472), v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473353), int32(3387), int32(360618))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+l1<<(uint(int32(2))%32)-int32(4))))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
						v57 = v55
						m.G0 = v7 + int32(16)
						return v57
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v14 == int32(0) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v31 == int32(0) {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v36 = F_get_attname(m, v34, l1, int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v57 = v36
						m.G0 = v7 + int32(16)
						return v57
					}
				} else {
					if l1 <= int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errmsg_internal(m, int32(167472), v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473353), int32(3387), int32(360618))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						if v43 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
							v46 = v44
						} else {
							v46 = int32(0)
						}
						if v46 < l1 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
								F_errmsg_internal(m, int32(167472), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(473353), int32(3387), int32(360618))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+l1<<(uint(int32(2))%32)-int32(4))))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
							v57 = v55
							m.G0 = v7 + int32(16)
							return v57
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v17 != 0 {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					v20 = v18
				} else {
					v20 = int32(0)
				}
				if v20 < l1 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v31 == int32(0) {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v36 = F_get_attname(m, v34, l1, int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v57 = v36
							m.G0 = v7 + int32(16)
							return v57
						}
					} else {
						if l1 <= int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
								F_errmsg_internal(m, int32(167472), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(473353), int32(3387), int32(360618))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							if v43 != 0 {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
								v46 = v44
							} else {
								v46 = int32(0)
							}
							if v46 < l1 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v68
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
									F_errmsg_internal(m, int32(167472), v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(473353), int32(3387), int32(360618))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+l1<<(uint(int32(2))%32)-int32(4))))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
								v57 = v55
								m.G0 = v7 + int32(16)
								return v57
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+l1<<(uint(int32(2))%32)-int32(4))))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					v57 = v29
					m.G0 = v7 + int32(16)
					return v57
				}
			}
		}
	}
}
func F_scanRTEForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v326
L2:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v20 <= v19 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v269 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v270 != 0 {
		v326 = v269
		goto L1
	} else {
		goto L81
	}
L5:
	;
	if v251 != 0 {
		v326 = v251
		goto L1
	} else {
		goto L80
	}
L6:
	;
	v251 = v8
	goto L5
L7:
	;
	goto L8
L8:
	;
	v25 = v19
	v31 = v8
	goto L9
L9:
	;
	v37 = v25 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v25<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v47 == int32(0) {
		v66 = v46
		v67 = v47
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L63
	} else {
		goto L75
	}
L11:
	;
	goto L10
L12:
	;
	if l6 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if v67-v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	goto L13
L15:
	;
	if v46 != v47 {
		v66 = v46
		v67 = v47
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v51 = v43
	v52 = l3
	goto L17
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v55
		v67 = v56
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v66 = v55
	v67 = v56
	goto L14
L19:
	;
	v59 = int32(1)
	if v55 == v56 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v69 = v31
	goto L12
L22:
	;
	goto L23
L23:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v69 = v37
	goto L12
L25:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v37 < v223 {
		v25 = v37
		v31 = v69
		goto L9
	} else {
		goto L74
	}
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v72 < l5 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v74 == int32(0) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if l3&int32(3) == int32(0) {
		v100 = l3
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v43&int32(3) == int32(0) {
		v157 = v43
		goto L48
	} else {
		goto L49
	}
L30:
	;
	v133 = v125 - l3
	goto L29
L31:
	;
	v104 = v100
	goto L40
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v84 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = int32(0)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v89 = l3
	goto L36
L36:
	;
	v93 = v89 + int32(1)
	if v93&int32(3) == int32(0) {
		v100 = v93
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v125 = v93
	goto L30
L38:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v98 != 0 {
		v89 = v93
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v113 = int32(-2139062144)
	if (int32(16843008)-v110|v110)&v113 == v113 {
		v104 = v104 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v119 = v104
	goto L43
L42:
	;
	goto L41
L43:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v123 != 0 {
		v119 = v119 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v125 = v119
	goto L30
L45:
	;
	goto L44
L46:
	;
	v194 = F_varstr_levenshtein_less_equal(m, v43, v190, l3, v133, v72-l5+int32(1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L63
	} else {
		goto L64
	}
L47:
	;
	v190 = v182 - v43
	goto L46
L48:
	;
	v161 = v157
	goto L57
L49:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v141 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v190 = int32(0)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v146 = v43
	goto L53
L53:
	;
	v150 = v146 + int32(1)
	if v150&int32(3) == int32(0) {
		v157 = v150
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v182 = v150
	goto L47
L55:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v155 != 0 {
		v146 = v150
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v170 = int32(-2139062144)
	if (int32(16843008)-v167|v167)&v170 == v170 {
		v161 = v161 + int32(4)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v176 = v161
	goto L60
L59:
	;
	goto L58
L60:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v180 != 0 {
		v176 = v176 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v182 = v176
	goto L47
L62:
	;
	goto L61
L63:
	;
	return int32(0)
L64:
	;
	v199 = base.I32_div_s(v133, int32(2))
	if v199 < v194 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	v201 = l5 + v194
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v201 < v202 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+8)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v201
	goto L25
L67:
	;
	goto L68
L68:
	;
	if v201 != v202 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v210 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v211
	goto L25
L71:
	;
	goto L72
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v215 == int32(0) {
		goto L25
	} else {
		goto L73
	}
L73:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+16)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = l1
	goto L25
L74:
	;
	v251 = v69
	goto L5
L75:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L63
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l3
	F_errmsg(m, int32(107050), v16)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L63
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(473353), int32(850), int32(261481))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L63
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	goto L4
L81:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v271 == int32(99) {
		v326 = v269
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v275 = F_strcmp(m, int32(723484), l3)
	mBase = m.M
	if v275 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v304 == int32(0) {
		v326 = v269
		goto L1
	} else {
		goto L102
	}
L84:
	;
	v304 = int32(723480)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v280 = F_strcmp(m, int32(723584), l3)
	mBase = m.M
	if v280 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v304 = int32(723580)
	goto L83
L88:
	;
	goto L89
L89:
	;
	v285 = F_strcmp(m, int32(723684), l3)
	mBase = m.M
	if v285 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v304 = int32(723680)
	goto L83
L91:
	;
	goto L92
L92:
	;
	v290 = F_strcmp(m, int32(723784), l3)
	mBase = m.M
	if v290 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v304 = int32(723780)
	goto L83
L94:
	;
	goto L95
L95:
	;
	v295 = F_strcmp(m, int32(723884), l3)
	mBase = m.M
	if v295 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v304 = int32(723880)
	goto L83
L97:
	;
	goto L98
L98:
	;
	v302 = F_strcmp(m, int32(723984), l3)
	mBase = m.M
	if v302 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v303 = int32(0)
	goto L101
L100:
	;
	v303 = int32(723980)
	goto L101
L101:
	;
	v304 = v303
	goto L83
L102:
	;
	v307 = int32(*(*int16)(unsafe.Add(mBase, uint32(v304)+74)))
	if v307 == int32(0) {
		v326 = v269
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v310 = int32(0)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v315 = F_SearchSysCacheExists(m, int32(7), v312, v307, v310, v310)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L63
	} else {
		goto L104
	}
L104:
	;
	if v315 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v317 = v307
	goto L107
L106:
	;
	v317 = v310
	goto L107
L107:
	;
	v326 = v317
	goto L1
}
