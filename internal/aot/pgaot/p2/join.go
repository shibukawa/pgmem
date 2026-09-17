package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_join_implied_equalities_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
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
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
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
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
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
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v736 int32
	_ = v736
	v6 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v22 = v21
	goto L3
L2:
	;
	v22 = v6
	goto L3
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = l2
	goto L6
L5:
	;
	v24 = int32(0)
	goto L6
L6:
	;
	v31 = int32(-1)
	v33 = v22
	v34 = v20
	v35 = v6
	v37 = v6
	v38 = v6
	goto L7
L7:
	;
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v357 = int32(0)
	if base.B2i32(v38 == v357)|base.B2i32(v35 == v357) == v357 {
		goto L89
	} else {
		goto L90
	}
L9:
	;
	goto L8
L10:
	;
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v152 == int32(0) {
		goto L9
	} else {
		goto L31
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v138 = v31
	v140 = v33
	v141 = v34
	v151 = v46
	goto L11
L13:
	;
	goto L14
L14:
	;
	v52 = v31
	goto L15
L15:
	;
	if v24 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v138 = v120
	v140 = v132
	v141 = v129
	v151 = v132
	goto L11
L17:
	;
	if v120 <= int32(0) {
		goto L9
	} else {
		goto L28
	}
L18:
	;
	v120 = base.I32_ctz(v106) | v107<<(uint(int32(5))%32)
	goto L17
L19:
	;
	v120 = int32(-2)
	goto L17
L20:
	;
	v71 = v52 + int32(1)
	v73 = base.I32_div_s(v71, int32(32))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v74 <= v73 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v77 = v24 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v73<<(uint(int32(2))%32))))
	v84 = v81 & (int32(-1) << (uint(v71) % 32))
	if v84 != 0 {
		v106 = v84
		v107 = v73
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v86 = v73 + int32(1)
	if v86 == v74 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v89 = v86
	goto L24
L24:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77+v89<<(uint(int32(2))%32))))
	if v96 != 0 {
		v106 = v96
		v107 = v89
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v98 = v89 + int32(1)
	if v98 != v74 {
		v89 = v98
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v123 <= v120 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v120<<(uint(int32(2))%32))))
	if v129 == int32(0) {
		v52 = v120
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v156 = v140 + int32(4)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if base.Ui32(v156) < base.Ui32(v151+v158<<(uint(int32(2))%32)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v163 = v156
	goto L34
L33:
	;
	v163 = int32(0)
	goto L34
L34:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v165 = int32(0)
	if v164 == v165 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v218 == int32(0) {
		v31 = v138
		v33 = v163
		v34 = v141
		goto L7
	} else {
		goto L49
	}
L36:
	;
	v218 = int32(1)
	goto L35
L37:
	;
	goto L38
L38:
	;
	if l2 == int32(0) {
		v211 = v165
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v218 = v211
	goto L35
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v175 < v174 {
		v211 = v165
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v177 = int32(1)
	if v174 <= v177 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v180 = v177
	goto L44
L43:
	;
	v180 = v174
	goto L44
L44:
	;
	v181 = int32(8)
	v186 = int32(0)
	goto L45
L45:
	;
	v193 = v186 << (uint(int32(2)) % 32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v164+v181+v193)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l2+v181+v193)))
	v200 = v195 & (v197 ^ int32(-1))
	v202 = base.B2i32(v200 == int32(0))
	if v200 != 0 {
		v211 = v202
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v211 = v202
	goto L39
L47:
	;
	v204 = v186 + int32(1)
	if v204 != v180 {
		v186 = v204
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v222 = int32(0)
	if v221 == v222 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v275 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v275 = int32(1)
	goto L50
L52:
	;
	goto L53
L53:
	;
	if l3 == int32(0) {
		v268 = v222
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v275 = v268
	goto L50
L55:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v232 < v231 {
		v268 = v222
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v234 = int32(1)
	if v231 <= v234 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v237 = v234
	goto L59
L58:
	;
	v237 = v231
	goto L59
L59:
	;
	v238 = int32(8)
	v243 = int32(0)
	goto L60
L60:
	;
	v250 = v243 << (uint(int32(2)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v221+v238+v250)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l3+v238+v250)))
	v257 = v252 & (v254 ^ int32(-1))
	v259 = base.B2i32(v257 == int32(0))
	if v257 != 0 {
		v268 = v259
		goto L54
	} else {
		goto L62
	}
L61:
	;
	v268 = v259
	goto L54
L62:
	;
	v261 = v243 + int32(1)
	if v261 != v237 {
		v243 = v261
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v276 = F_lappend(m, v38, v152)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v281 = int32(0)
	if v280 == v281 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	return int32(0)
L68:
	;
	v31 = v138
	v33 = v163
	v34 = v141
	v38 = v276
	goto L7
L69:
	;
	if v334 != 0 {
		goto L83
	} else {
		goto L84
	}
L70:
	;
	v334 = int32(1)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if l4 == int32(0) {
		v327 = v281
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v334 = v327
	goto L69
L74:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v291 < v290 {
		v327 = v281
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v293 = int32(1)
	if v290 <= v293 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v296 = v293
	goto L78
L77:
	;
	v296 = v290
	goto L78
L78:
	;
	v297 = int32(8)
	v302 = int32(0)
	goto L79
L79:
	;
	v309 = v302 << (uint(int32(2)) % 32)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v280+v297+v309)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l4+v297+v309)))
	v316 = v311 & (v313 ^ int32(-1))
	v318 = base.B2i32(v316 == int32(0))
	if v316 != 0 {
		v327 = v318
		goto L73
	} else {
		goto L81
	}
L80:
	;
	v327 = v318
	goto L73
L81:
	;
	v320 = v302 + int32(1)
	if v320 != v296 {
		v302 = v320
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v335 = F_lappend(m, v35, v152)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L67
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v337 = F_lappend(m, v37, v152)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L67
	} else {
		goto L87
	}
L86:
	;
	v31 = v138
	v33 = v163
	v34 = v141
	v35 = v335
	goto L7
L87:
	;
	v31 = v138
	v33 = v163
	v34 = v141
	v37 = v337
	goto L7
L88:
	;
	v736 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)) = uint8(v736)
	return int32(0)
L89:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v365 <= int32(0) {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	v575 = v357
	goto L91
L91:
	;
	if v37 == int32(0) {
		v705 = v575
		goto L137
	} else {
		goto L138
	}
L92:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v374 = int32(-1)
	v375 = v368
	v384 = v6
	v385 = v6
	v386 = v6
	v387 = v6
	goto L94
L93:
	;
	v565 = F_create_join_clause(m, l0, l1, v560, v563, v561, l1)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L67
	} else {
		goto L135
	}
L94:
	;
	if int32(0) < v375 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v526 < int32(0) {
		goto L88
	} else {
		goto L134
	}
L96:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v390+v386<<(uint(int32(2))%32))))
	v400 = v374
	v402 = int32(0)
	v410 = v384
	v411 = v385
	v413 = v387
	goto L99
L97:
	;
	v526 = v374
	v527 = v375
	v536 = v384
	v537 = v385
	v539 = v387
	goto L98
L98:
	;
	v541 = v386 + int32(1)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v541 < v542 {
		v374 = v526
		v375 = v527
		v384 = v536
		v385 = v537
		v386 = v541
		v387 = v539
		goto L94
	} else {
		goto L133
	}
L99:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v414 == int32(0) {
		v504 = v400
		v514 = v410
		v515 = v411
		v517 = v413
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v526 = v504
	v527 = v520
	v536 = v514
	v537 = v515
	v539 = v517
	goto L98
L101:
	;
	v519 = v402 + int32(1)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v519 < v520 {
		v400 = v504
		v402 = v519
		v410 = v514
		v411 = v515
		v413 = v517
		goto L99
	} else {
		goto L132
	}
L102:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v417 <= int32(0) {
		v504 = v400
		v514 = v410
		v515 = v411
		v517 = v413
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v402<<(uint(int32(2))%32))))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+16))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	v433 = int32(0)
	goto L104
L104:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v414)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446+v433<<(uint(int32(2))%32))))
	v452 = F_get_opfamily_member_for_cmptype(m, v450, v426, v425, int32(3))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L67
	} else {
		goto L107
	}
L105:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v466 != int32(6) {
		goto L117
	} else {
		goto L118
	}
L106:
	;
	goto L105
L107:
	;
	if v452 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v454 == int32(0) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v462 = v433 + int32(1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v462 < v463 {
		v433 = v462
		goto L104
	} else {
		goto L115
	}
L111:
	;
	v457 = F_get_opcode(m, v452)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L67
	} else {
		goto L112
	}
L112:
	;
	v459 = F_get_func_leakproof(m, v457)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L67
	} else {
		goto L113
	}
L113:
	;
	if v459 != 0 {
		goto L106
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v504 = v400
	v514 = v410
	v515 = v411
	v517 = v413
	goto L101
L116:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if v480 != int32(6) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	v469 = int32(0)
	if v466 != int32(27) {
		v478 = v469
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v478 = int32(1)
	goto L116
L120:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v465)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	if v473 != int32(6) {
		v478 = v469
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v492 = F_exprType(m, v465)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L67
	} else {
		goto L128
	}
L123:
	;
	if v480 != int32(27) {
		v491 = v478
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v491 = v478 + int32(1)
	goto L122
L126:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	if v486 != int32(6) {
		v491 = v478
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v494 = F_op_hashjoinable(m, v452, v492)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L67
	} else {
		goto L129
	}
L129:
	;
	v496 = v494 + v491
	if v496 <= v400 {
		v504 = v400
		v514 = v410
		v515 = v411
		v517 = v413
		goto L101
	} else {
		goto L130
	}
L130:
	;
	if v496 != int32(3) {
		v504 = v496
		v514 = v452
		v515 = v424
		v517 = v394
		goto L101
	} else {
		goto L131
	}
L131:
	;
	v560 = v452
	v561 = v424
	v563 = v394
	goto L93
L132:
	;
	goto L100
L133:
	;
	goto L95
L134:
	;
	v560 = v536
	v561 = v537
	v563 = v539
	goto L93
L135:
	;
	v567 = F_lappend(m, int32(0), v565)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L67
	} else {
		goto L136
	}
L136:
	;
	v575 = v567
	goto L91
L137:
	;
	return v705
L138:
	;
	v589 = F_list_concat(m, v38, v35)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L67
	} else {
		goto L139
	}
L139:
	;
	if v589 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	v593 = F_lappend(m, v37, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L67
	} else {
		goto L143
	}
L141:
	;
	v597 = v37
	goto L142
L142:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v598 <= int32(0) {
		v705 = v575
		goto L137
	} else {
		goto L145
	}
L143:
	;
	if v593 == int32(0) {
		v705 = v575
		goto L137
	} else {
		goto L144
	}
L144:
	;
	v597 = v593
	goto L142
L145:
	;
	v601 = int32(0)
	v605 = v601
	v609 = v575
	v612 = v601
	goto L146
L146:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v621+v605<<(uint(int32(2))%32))))
	if v612 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v705 = v683
	goto L137
L148:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v626 == int32(0) {
		goto L88
	} else {
		goto L151
	}
L149:
	;
	v683 = v609
	goto L150
L150:
	;
	v696 = v605 + int32(1)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v696 < v697 {
		v605 = v696
		v609 = v683
		v612 = v625
		goto L146
	} else {
		goto L167
	}
L151:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	if v629 <= int32(0) {
		goto L88
	} else {
		goto L152
	}
L152:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v625)+16))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v612)+16))
	v640 = int32(0)
	goto L153
L153:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v653+v640<<(uint(int32(2))%32))))
	v659 = F_get_opfamily_member_for_cmptype(m, v657, v633, v632, int32(3))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L67
	} else {
		goto L156
	}
L154:
	;
	v673 = F_create_join_clause(m, l0, l1, v659, v612, v625, int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L67
	} else {
		goto L165
	}
L155:
	;
	goto L154
L156:
	;
	if v659 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v661 == int32(0) {
		goto L155
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v669 = v640 + int32(1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	if v669 < v670 {
		v640 = v669
		goto L153
	} else {
		goto L164
	}
L160:
	;
	v664 = F_get_opcode(m, v659)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L67
	} else {
		goto L161
	}
L161:
	;
	v666 = F_get_func_leakproof(m, v664)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L67
	} else {
		goto L162
	}
L162:
	;
	if v666 != 0 {
		goto L155
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	goto L88
L165:
	;
	v675 = F_lappend(m, v609, v673)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L67
	} else {
		goto L166
	}
L166:
	;
	v683 = v675
	goto L150
L167:
	;
	goto L147
}
