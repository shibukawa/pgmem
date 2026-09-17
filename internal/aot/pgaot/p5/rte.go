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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
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
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
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
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
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
	var v529 int32
	_ = v529
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
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
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v651 == int32(0) {
		goto L7
	} else {
		goto L185
	}
L9:
	;
	v650 = v347 + int32(1)
	goto L8
L10:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v529 = v9
	goto L147
L11:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	if v397 != 0 {
		goto L109
	} else {
		goto L110
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L31
	} else {
		goto L106
	}
L13:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v372 = F_relation_open(m, v370, int32(1))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L31
	} else {
		goto L103
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L31
	} else {
		goto L100
	}
L15:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v125 == int32(0) {
		v650 = int32(1)
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
	v48 = v35
	v51 = v9
	v53 = v9
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v51<<(uint(int32(2))%32))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+26)))
	if v69 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L7
L24:
	;
	if v48 == int32(0) {
		goto L14
	} else {
		goto L27
	}
L25:
	;
	v117 = v48
	v118 = v53
	goto L26
L26:
	;
	v121 = v51 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v121 < v122 {
		v48 = v117
		v51 = v121
		v53 = v118
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v85 = v53 + int32(1)
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
	v79 = F_makeString(m, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v81 = F_lappend(m, v74, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v81
	goto L30
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v88 = F_exprType(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L31
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v106 = v48 + int32(4)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if base.Ui32(v106) < base.Ui32(v110+v111<<(uint(int32(2))%32)) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v91 = F_exprTypmod(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v94 = F_exprCollation(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v96 = F_makeVar(m, l1, base.I32_extend16_s(v85), v88, v91, v94, l2)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v96)+32)) = l3
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v101 = F_lappend(m, v100, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v101
	goto L37
L43:
	;
	v116 = v106
	goto L45
L44:
	;
	v116 = int32(0)
	goto L45
L45:
	;
	v117 = v116
	v118 = v85
	goto L26
L46:
	;
	goto L23
L47:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v129 <= int32(0) {
		v650 = int32(1)
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v141 = v9
	v150 = v9
	goto L49
L49:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v150<<(uint(int32(2))%32))))
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
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v347 = v346 + v141
	v349 = v150 + int32(1)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v349 < v350 {
		v141 = v347
		v150 = v349
		goto L49
	} else {
		goto L99
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
	F_expandTupleDesc(m, v174, v175, v176, v141, l1, l2, l3, l4, l5, l6, l7)
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
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v141<<(uint(int32(2))%32))))
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
	v202 = F_makeVar(m, l1, base.I32_extend16_s(v141+int32(1)), v195, v197, v200, l2)
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
	F_errmsg_internal(m, int32(_a_F_expandRTE_0), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_expandRTE_1), int32(2926), int32(_a_F_expandRTE_2))
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
	v226 = F_list_copy_tail(m, v225, v141)
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
	if base.B2i32(v226 == v229)|base.B2i32(v228 <= v229) != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v241 = F_list_concat(m, v240, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L31
	} else {
		goto L84
	}
L78:
	;
	v239 = int32(0)
	goto L80
L79:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v228 < v236 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v228
	goto L83
L82:
	;
	goto L83
L83:
	;
	v239 = v226
	goto L80
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v241
	goto L75
L85:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v157)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v259 = int32(0)
	v263 = v141
	goto L86
L86:
	;
	v272 = int32(0)
	if v249 == v272 {
		v282 = v272
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v283 = int32(0)
	if v248 == v283 {
		v292 = v283
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	if v276 <= v259 {
		v282 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v282 = v278 + v259<<(uint(int32(2))%32)
	goto L88
L91:
	;
	if v247 == int32(0) {
		goto L51
	} else {
		goto L94
	}
L92:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v286 <= v259 {
		v292 = v283
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v292 = v288 + v259<<(uint(int32(2))%32)
	goto L91
L94:
	;
	v295 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	if base.B2i32(v292 == v295)|(base.B2i32(v282 == v295)|base.B2i32(v299 <= v259)) != 0 {
		goto L51
	} else {
		goto L95
	}
L95:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v303 == int32(0) {
		goto L51
	} else {
		goto L96
	}
L96:
	;
	v307 = v263 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v303+v259<<(uint(int32(2))%32))))
	v315 = F_makeVar(m, l1, base.I32_extend16_s(v307), v309, v310, v314, l2)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L31
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v315)+32)) = l3
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v320 = F_lappend(m, v319, v315)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L31
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v320
	v259 = v259 + int32(1)
	v263 = v307
	goto L86
L99:
	;
	goto L50
L100:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v357
	F_errmsg_internal(m, int32(_a_F_expandRTE_3), v24+int32(16))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L31
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_expandRTE_1), int32(2797), int32(_a_F_expandRTE_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L31
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+52))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	F_expandTupleDesc(m, v374, v369, v375, int32(0), l1, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L31
	} else {
		goto L104
	}
L104:
	;
	F_relation_close(m, v372, int32(1))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L31
	} else {
		goto L105
	}
L105:
	;
	goto L7
L106:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v386
	F_errmsg_internal(m, int32(_a_F_expandRTE_4), v24)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L31
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_expandRTE_1), int32(3110), int32(_a_F_expandRTE_2))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L31
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v399 = v398
	goto L111
L110:
	;
	v399 = v9
	goto L111
L111:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v412 = int32(0)
	v413 = v399
	goto L112
L112:
	;
	v425 = int32(0)
	if v402 == v425 {
		v436 = v425
		goto L114
	} else {
		goto L115
	}
L114:
	;
	if v401 == int32(0) {
		v445 = v425
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	if v430 <= v412 {
		v436 = int32(0)
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v402)+12))
	v436 = v432 + v412<<(uint(int32(2))%32)
	goto L114
L117:
	;
	if v400 == int32(0) {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v439 <= v412 {
		v445 = v425
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v445 = v441 + v412<<(uint(int32(2))%32)
	goto L117
L120:
	;
	v448 = int32(0)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if base.B2i32(v445 == v448)|(base.B2i32(v436 == v448)|base.B2i32(v452 <= v412)) != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	if v456 == int32(0) {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v456+v412<<(uint(int32(2))%32))))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	if l6 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if v464 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v492 = v413
	goto L125
L125:
	;
	v495 = v412 + int32(1)
	if l7 == int32(0) {
		v412 = v495
		v413 = v492
		goto L112
	} else {
		goto L138
	}
L126:
	;
	v481 = v413 + int32(4)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	if base.Ui32(v481) < base.Ui32(v485+v486<<(uint(int32(2))%32)) {
		goto L135
	} else {
		goto L136
	}
L127:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v472 = F_pstrdup(m, v470)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L31
	} else {
		goto L132
	}
L128:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	v470 = v466
	goto L127
L129:
	;
	goto L130
L130:
	;
	if l5 == int32(0) {
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v470 = int32(_a_F_expandRTE_5)
	goto L127
L132:
	;
	v474 = F_makeString(m, v472)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L31
	} else {
		goto L133
	}
L133:
	;
	v476 = F_lappend(m, v471, v474)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L31
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v476
	goto L126
L135:
	;
	v491 = v481
	goto L137
L136:
	;
	v491 = int32(0)
	goto L137
L137:
	;
	v492 = v491
	goto L125
L138:
	;
	if v464 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v499 = F_makeVar(m, l1, base.I32_extend16_s(v495), v464, v463, v462, l2)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L31
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if l5 == int32(0) {
		v412 = v495
		v413 = v492
		goto L112
	} else {
		goto L144
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v499)+32)) = l3
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v504 = F_lappend(m, v503, v499)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L31
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v504
	v412 = v495
	v413 = v492
	goto L112
L144:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v513 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L31
	} else {
		goto L145
	}
L145:
	;
	v515 = F_lappend(m, v509, v513)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L31
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v515
	v412 = v495
	v413 = v492
	goto L112
L147:
	;
	v542 = int32(0)
	if v520 == v542 {
		v552 = v542
		goto L149
	} else {
		goto L150
	}
L149:
	;
	if v518 == int32(0) {
		goto L7
	} else {
		goto L152
	}
L150:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v546 <= v529 {
		v552 = int32(0)
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v552 = v548 + v529<<(uint(int32(2))%32)
	goto L149
L152:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if base.B2i32(v552 == int32(0))|base.B2i32(v557 <= v529) != 0 {
		goto L7
	} else {
		goto L153
	}
L153:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if v560 == int32(0) {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	v564 = v529 + int32(1)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v560+v529<<(uint(int32(2))%32))))
	if v568 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	if l5 == int32(0) {
		v529 = v564
		goto L147
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if l6 != 0 {
		goto L168
	} else {
		goto L169
	}
L158:
	;
	if l6 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v575 = F_pstrdup(m, int32(_a_F_expandRTE_5))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L31
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if l7 == int32(0) {
		v529 = v564
		goto L147
	} else {
		goto L165
	}
L162:
	;
	v577 = F_makeString(m, v575)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L31
	} else {
		goto L163
	}
L163:
	;
	v579 = F_lappend(m, v573, v577)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L31
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v579
	goto L161
L165:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v588 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L31
	} else {
		goto L166
	}
L166:
	;
	v590 = F_lappend(m, v584, v588)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L31
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v590
	v529 = v564
	goto L147
L168:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v596 = F_pstrdup(m, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L31
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	if l7 == int32(0) {
		v529 = v564
		goto L147
	} else {
		goto L174
	}
L171:
	;
	v598 = F_makeString(m, v596)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L31
	} else {
		goto L172
	}
L172:
	;
	v600 = F_lappend(m, v593, v598)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L31
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v600
	goto L170
L174:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v605 == int32(6) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v620)+32)) = l3
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v624 = F_lappend(m, v623, v620)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L31
	} else {
		goto L184
	}
L176:
	;
	v608 = F_copyObjectImpl(m, v568)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L31
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v612 = F_exprType(m, v568)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L31
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+28)) = l2
	v620 = v608
	goto L175
L180:
	;
	v614 = F_exprTypmod(m, v568)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L31
	} else {
		goto L181
	}
L181:
	;
	v616 = F_exprCollation(m, v568)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L31
	} else {
		goto L182
	}
L182:
	;
	v618 = F_makeVar(m, l1, base.I32_extend16_s(v564), v612, v614, v616, l2)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L31
	} else {
		goto L183
	}
L183:
	;
	v620 = v618
	goto L175
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v624
	v529 = v564
	goto L147
L185:
	;
	if l6 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+8))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v657+v658<<(uint(int32(2))%32)-int32(4))))
	v665 = F_lappend(m, v654, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L31
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if l7 == int32(0) {
		goto L7
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v665
	goto L188
L190:
	;
	v675 = F_makeVar(m, l1, base.I32_extend16_s(v650), int32(20), int32(-1), int32(0), l2)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L31
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v675)+32)) = l3
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v679 = F_lappend(m, v678, v675)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L31
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v679
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
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
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v108 = F_quote_identifier(m, v22)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L30
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+v13-int32(4))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+20)))
	if v30 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v31 {
	case 0:
		goto L7
	case 1, 3, 5:
		goto L2
	default:
		goto L1
	case 6:
		goto L6
	}
L5:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if base.B2i32(v79 == int32(0))|base.B2i32(v79 != v82) != 0 {
		v100 = v79
		v101 = v82
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if base.B2i32(v53 == int32(0))|base.B2i32(v53 != v56) != 0 {
		v74 = v53
		v75 = v56
		goto L15
	} else {
		goto L16
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
	if v33 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v32
	F_errmsg_internal(m, int32(_a_F_get_rte_alias_0), v10+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_get_rte_alias_1), int32(_a_F_get_rte_alias_2), int32(_a_F_get_rte_alias_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	if v74-v75 != 0 {
		goto L2
	} else {
		goto L21
	}
L15:
	;
	goto L14
L16:
	;
	v59 = v22
	v60 = v50
	goto L17
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v64 == int32(0) {
		v74 = v64
		v75 = v63
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v74 = v64
	v75 = v63
	goto L15
L19:
	;
	v67 = int32(1)
	if v64 == v63 {
		v59 = v59 + v67
		v60 = v60 + v67
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L1
L22:
	;
	if v100-v101 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v85 = v22
	v86 = v33
	goto L25
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v90
		v101 = v89
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v100 = v90
	v101 = v89
	goto L23
L27:
	;
	v93 = int32(1)
	if v90 == v89 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L2
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v108
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = int32(_a_F_get_rte_alias_4)
	goto L33
L32:
	;
	v113 = int32(_a_F_get_rte_alias_5)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v113
	F_appendStringInfo(m, v107, int32(_a_F_get_rte_alias_6), v10)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L1
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
		v57 = int32(_a_F_get_rte_attribute_name_0)
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
						F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
					if v43 == int32(0) {
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
							F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
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
								F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
							F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
						if v43 == int32(0) {
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
								F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
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
									F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				if v17 == int32(0) {
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
								F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
							if v43 == int32(0) {
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
									F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
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
										F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
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
									F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
								if v43 == int32(0) {
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
										F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
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
											F_errmsg_internal(m, int32(_a_F_get_rte_attribute_name_1), v7)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_rte_attribute_name_2), int32(3387), int32(_a_F_get_rte_attribute_name_3))
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
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L27
	} else {
		goto L67
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v202
L3:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 < v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v145 != 0 {
		v202 = v144
		goto L2
	} else {
		goto L40
	}
L6:
	;
	v25 = v19
	v32 = v8
	goto L9
L7:
	;
	v127 = v8
	goto L8
L8:
	;
	if v127 != 0 {
		v202 = v127
		goto L2
	} else {
		goto L39
	}
L9:
	;
	v37 = v25 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v25<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v127 = v70
	goto L8
L11:
	;
	if l6 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v67-v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v52 = v43
	v53 = l3
	goto L15
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v67 = v57
	v68 = v56
	goto L13
L17:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v70 = v32
	goto L11
L20:
	;
	goto L21
L21:
	;
	if v32 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v70 = v37
	goto L11
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v37 < v116 {
		v25 = v37
		v32 = v70
		goto L9
	} else {
		goto L38
	}
L24:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v73 < l5 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v75 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v78 = F_strlen(m, l3)
	mBase = m.M
	v79 = F_strlen(m, v43)
	mBase = m.M
	v80 = int32(1)
	v87 = F_varstr_levenshtein_less_equal(m, v43, v79, l3, v78, v80, v80, v80, v73-l5+v80, v80)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	v92 = base.I32_div_s(v78, int32(2))
	if v92 < v87 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v94 = l5 + v87
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v94 < v95 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+8)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v94
	goto L23
L31:
	;
	goto L32
L32:
	;
	if v94 != v95 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v103 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v104
	goto L23
L35:
	;
	goto L36
L36:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v108 == int32(0) {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+16)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = l1
	goto L23
L38:
	;
	goto L10
L39:
	;
	goto L5
L40:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+21)))
	if v146 == int32(99) {
		v202 = v144
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v150 = F_strcmp(m, int32(_a_F_scanRTEForColumn_0), l3)
	mBase = m.M
	if v150 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v179 == int32(0) {
		v202 = v144
		goto L2
	} else {
		goto L61
	}
L43:
	;
	v179 = int32(_a_F_scanRTEForColumn_1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v155 = F_strcmp(m, int32(_a_F_scanRTEForColumn_2), l3)
	mBase = m.M
	if v155 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v179 = int32(_a_F_scanRTEForColumn_3)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v160 = F_strcmp(m, int32(_a_F_scanRTEForColumn_4), l3)
	mBase = m.M
	if v160 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v179 = int32(_a_F_scanRTEForColumn_5)
	goto L42
L50:
	;
	goto L51
L51:
	;
	v165 = F_strcmp(m, int32(_a_F_scanRTEForColumn_6), l3)
	mBase = m.M
	if v165 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v179 = int32(_a_F_scanRTEForColumn_7)
	goto L42
L53:
	;
	goto L54
L54:
	;
	v170 = F_strcmp(m, int32(_a_F_scanRTEForColumn_8), l3)
	mBase = m.M
	if v170 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v179 = int32(_a_F_scanRTEForColumn_9)
	goto L42
L56:
	;
	goto L57
L57:
	;
	v177 = F_strcmp(m, int32(_a_F_scanRTEForColumn_10), l3)
	mBase = m.M
	if v177 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v178 = int32(0)
	goto L60
L59:
	;
	v178 = int32(_a_F_scanRTEForColumn_11)
	goto L60
L60:
	;
	v179 = v178
	goto L42
L61:
	;
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179)+74)))
	if v182 == int32(0) {
		v202 = v144
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v185 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v190 = F_SearchSysCacheExists(m, int32(7), v187, v182, v185, v185)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L27
	} else {
		goto L63
	}
L63:
	;
	if v190 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v192 = v182
	goto L66
L65:
	;
	v192 = v185
	goto L66
L66:
	;
	v202 = v192
	goto L2
L67:
	;
	F_errcode(m, int32(33583236))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L27
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l3
	F_errmsg(m, int32(_a_F_scanRTEForColumn_12), v16)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_scanRTEForColumn_13), int32(850), int32(_a_F_scanRTEForColumn_14))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
