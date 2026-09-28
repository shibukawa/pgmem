package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_grouping_conflict_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
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
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v686 int32
	_ = v686
	var v698 int32
	_ = v698
	var v721 int32
	_ = v721
	v3 = int32(0)
	if l0 == v3 {
		v721 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v721
L2:
	;
	v18 = l1 + int32(8)
	v19 = l0
	goto L4
L3:
	;
	v721 = int32(1)
	goto L1
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	switch v33 - int32(6) {
	case 0:
		goto L16
	default:
		goto L9
	case 11:
		goto L15
	case 14:
		goto L14
	case 23:
		goto L12
	case 26:
		goto L10
	case 28:
		v686 = v18
		goto L7
	case 31:
		goto L13
	case 39:
		goto L11
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v552
	goto L3
L6:
	;
	goto L5
L7:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v686)))
	if v698 != 0 {
		v19 = v698
		goto L4
	} else {
		goto L188
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v552
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v635 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L9:
	;
	v617 = F_expression_tree_walker_impl(m, v19, int32(925), l1)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L17
	} else {
		goto L179
	}
L10:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v549 == int32(0) {
		goto L9
	} else {
		goto L160
	}
L11:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v535 = int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v537 = F_grouping_conflict_walker(m, v536, l1)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L17
	} else {
		goto L155
	}
L12:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v525 = F_grouping_conflict_walker(m, v524, l1)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L17
	} else {
		goto L152
	}
L13:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v294 = int32(0)
	goto L89
L14:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v171 = F_op_is_safe_index_member(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L17
	} else {
		goto L56
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v52 = F_op_is_safe_index_member(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L17
	} else {
		goto L23
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v38 = m.T0[v37].(func(*base.Module, int32, int32) int32)(m, v19, v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v38 == int32(0) {
		v721 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v44 == int32(0) {
		v721 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v47 = F_get_collation_isdeterministic(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v47 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v721 = v3
	goto L1
L23:
	;
	if v52 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v56 == int32(0) {
		v721 = v3
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v59 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int32(0)
L27:
	;
	goto L28
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v67 = int32(0)
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v67<<(uint(int32(2))%32))))
	if v85 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	return int32(0)
L31:
	;
	v165 = v67 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v165 < v166 {
		v67 = v165
		goto L29
	} else {
		goto L55
	}
L32:
	;
	v148 = F_grouping_conflict_walker(m, v85, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L53
	}
L33:
	;
	v90 = v85
	goto L34
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v102 != int32(27) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L32
L36:
	;
	if v102 == int32(34) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v133 != 0 {
		v90 = v133
		goto L34
	} else {
		goto L52
	}
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v107 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L40:
	;
	v111 = v90
	v112 = v102
	goto L41
L41:
	;
	if v112 != int32(6) {
		goto L32
	} else {
		goto L43
	}
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v111 = v107
	v112 = v110
	goto L41
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v117 = m.T0[v116].(func(*base.Module, int32, int32) int32)(m, v111, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	if v117 == int32(0) {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v121 = int32(1)
	v122 = F_equality_ops_are_compatible(m, v65, v117)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	if v122 == int32(0) {
		v721 = v121
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v126 == int32(0) {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	v129 = F_get_collation_isdeterministic(m, v126)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	if v129 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v64 == v131 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v721 = v121
	goto L1
L52:
	;
	goto L35
L53:
	;
	if v148 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	goto L31
L55:
	;
	goto L30
L56:
	;
	if v171 == int32(0) {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v175 == int32(0) {
		v721 = v3
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v178 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	return int32(0)
L60:
	;
	goto L61
L61:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v186 = int32(0)
	goto L62
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v186<<(uint(int32(2))%32))))
	if v204 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	return int32(0)
L64:
	;
	v284 = v186 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v284 < v285 {
		v186 = v284
		goto L62
	} else {
		goto L88
	}
L65:
	;
	v267 = F_grouping_conflict_walker(m, v204, l1)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L17
	} else {
		goto L86
	}
L66:
	;
	v209 = v204
	goto L67
L67:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v221 != int32(27) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L65
L69:
	;
	if v221 == int32(34) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v252 != 0 {
		v209 = v252
		goto L67
	} else {
		goto L85
	}
L72:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v226 == int32(0) {
		goto L65
	} else {
		goto L75
	}
L73:
	;
	v230 = v209
	v231 = v221
	goto L74
L74:
	;
	if v231 != int32(6) {
		goto L65
	} else {
		goto L76
	}
L75:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v230 = v226
	v231 = v229
	goto L74
L76:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v236 = m.T0[v235].(func(*base.Module, int32, int32) int32)(m, v230, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	if v236 == int32(0) {
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v240 = int32(1)
	v241 = F_equality_ops_are_compatible(m, v184, v236)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	if v241 == int32(0) {
		v721 = v240
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v245 == int32(0) {
		goto L64
	} else {
		goto L81
	}
L81:
	;
	v248 = F_get_collation_isdeterministic(m, v245)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L17
	} else {
		goto L82
	}
L82:
	;
	if v248 != 0 {
		goto L64
	} else {
		goto L83
	}
L83:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v183 == v250 {
		goto L64
	} else {
		goto L84
	}
L84:
	;
	v721 = v240
	goto L1
L85:
	;
	goto L68
L86:
	;
	if v267 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	goto L64
L88:
	;
	goto L63
L89:
	;
	v308 = int32(0)
	if v292 == v308 {
		v318 = v308
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v319 = int32(0)
	if v291 == v319 {
		v330 = v319
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v312 <= v294 {
		v318 = int32(0)
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v318 = v314 + v294<<(uint(int32(2))%32)
	goto L91
L94:
	;
	if v290 == int32(0) {
		v339 = v319
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v324 <= v294 {
		v330 = int32(0)
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v330 = v326 + v294<<(uint(int32(2))%32)
	goto L94
L97:
	;
	v340 = int32(0)
	if v289 == v340 {
		v349 = v340
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	if v333 <= v294 {
		v339 = v319
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v339 = v335 + v294<<(uint(int32(2))%32)
	goto L97
L100:
	;
	v350 = int32(0)
	v360 = base.B2i32(v318 != v350) & base.B2i32(v330 != v350) & base.B2i32(v339 != v350) & base.B2i32(v349 != v350)
	if v360 == v350 {
		v721 = v360
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v343 <= v294 {
		v349 = v340
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v349 = v345 + v294<<(uint(int32(2))%32)
	goto L100
L103:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	if v365 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v443 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L105:
	;
	v427 = F_grouping_conflict_walker(m, v365, l1)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L17
	} else {
		goto L126
	}
L106:
	;
	v370 = v365
	goto L107
L107:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	if v382 != int32(27) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L105
L109:
	;
	if v382 == int32(34) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v412 != 0 {
		v370 = v412
		goto L107
	} else {
		goto L125
	}
L112:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v387 == int32(0) {
		goto L105
	} else {
		goto L115
	}
L113:
	;
	v391 = v370
	v392 = v382
	goto L114
L114:
	;
	if v392 != int32(6) {
		goto L105
	} else {
		goto L116
	}
L115:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	v391 = v387
	v392 = v390
	goto L114
L116:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v397 = m.T0[v396].(func(*base.Module, int32, int32) int32)(m, v391, v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L117
	}
L117:
	;
	if v397 == int32(0) {
		goto L104
	} else {
		goto L118
	}
L118:
	;
	v401 = F_equality_ops_are_compatible(m, v364, v397)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	if v401 == int32(0) {
		v721 = v360
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	if v405 == int32(0) {
		goto L104
	} else {
		goto L121
	}
L121:
	;
	v408 = F_get_collation_isdeterministic(m, v405)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L17
	} else {
		goto L122
	}
L122:
	;
	if v408 != 0 {
		goto L104
	} else {
		goto L123
	}
L123:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	if v363 == v410 {
		goto L104
	} else {
		goto L124
	}
L124:
	;
	v721 = v360
	goto L1
L125:
	;
	goto L108
L126:
	;
	if v427 != 0 {
		v721 = v360
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L104
L128:
	;
	v294 = v294 + int32(1)
	goto L89
L129:
	;
	v505 = F_grouping_conflict_walker(m, v443, l1)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L17
	} else {
		goto L150
	}
L130:
	;
	v448 = v443
	goto L131
L131:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v460 != int32(27) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L129
L133:
	;
	if v460 == int32(34) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v490 != 0 {
		v448 = v490
		goto L131
	} else {
		goto L149
	}
L136:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v465 == int32(0) {
		goto L129
	} else {
		goto L139
	}
L137:
	;
	v469 = v448
	v470 = v460
	goto L138
L138:
	;
	if v470 != int32(6) {
		goto L129
	} else {
		goto L140
	}
L139:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	v469 = v465
	v470 = v468
	goto L138
L140:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v475 = m.T0[v474].(func(*base.Module, int32, int32) int32)(m, v469, v473)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	if v475 == int32(0) {
		goto L128
	} else {
		goto L142
	}
L142:
	;
	v479 = F_equality_ops_are_compatible(m, v364, v475)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	if v479 == int32(0) {
		v721 = v360
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v469)+20))
	if v483 == int32(0) {
		goto L128
	} else {
		goto L145
	}
L145:
	;
	v486 = F_get_collation_isdeterministic(m, v483)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	if v486 != 0 {
		goto L128
	} else {
		goto L147
	}
L147:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v469)+20))
	if v363 == v488 {
		goto L128
	} else {
		goto L148
	}
L148:
	;
	v721 = v360
	goto L1
L149:
	;
	goto L132
L150:
	;
	if v505 != 0 {
		v721 = v360
		goto L1
	} else {
		goto L151
	}
L151:
	;
	goto L128
L152:
	;
	if v525 != 0 {
		goto L3
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v530 = F_grouping_conflict_walker(m, v529, l1)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v523
	return v530
L155:
	;
	if v537 != 0 {
		v721 = v535
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v540 = F_grouping_conflict_walker(m, v539, l1)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	if v540 != 0 {
		v721 = v535
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v545 = F_grouping_conflict_walker(m, v544, l1)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v534
	return v545
L160:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v555 = v549
	goto L163
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v580
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v582 == int32(0) {
		goto L8
	} else {
		goto L172
	}
L162:
	;
	v577 = F_grouping_conflict_walker(m, v574, l1)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L17
	} else {
		goto L170
	}
L163:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v567 != int32(27) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v574 = int32(0)
	goto L162
L165:
	;
	if v567 != int32(6) {
		v574 = v555
		goto L162
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v572 != 0 {
		v555 = v572
		goto L163
	} else {
		goto L169
	}
L168:
	;
	v580 = v555
	goto L161
L169:
	;
	goto L164
L170:
	;
	if v577 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	v580 = int32(0)
	goto L161
L172:
	;
	v585 = int32(0)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v586 <= v585 {
		goto L8
	} else {
		goto L173
	}
L173:
	;
	v591 = v585
	goto L174
L174:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v582)+12))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603+v591<<(uint(int32(2))%32))))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	v609 = F_grouping_conflict_walker(m, v608, l1)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L17
	} else {
		goto L176
	}
L175:
	;
	goto L8
L176:
	;
	if v609 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v612 = v591 + int32(1)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v582)+4))
	if v612 < v613 {
		v591 = v612
		goto L174
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	return v617
L180:
	;
	v686 = v19 + int32(20)
	goto L7
L181:
	;
	v638 = int32(0)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v639 <= v638 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v644 = v638
	goto L183
L183:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656+v644<<(uint(int32(2))%32))))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	v662 = F_grouping_conflict_walker(m, v661, l1)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L17
	} else {
		goto L185
	}
L184:
	;
	goto L180
L185:
	;
	if v662 != 0 {
		goto L3
	} else {
		goto L186
	}
L186:
	;
	v665 = v644 + int32(1)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v665 < v666 {
		v644 = v665
		goto L183
	} else {
		goto L187
	}
L187:
	;
	goto L184
L188:
	;
	v721 = v3
	goto L1
}
func F_grouping_is_hashable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		v34 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v34
L5:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v16 = v10
	goto L8
L7:
	;
	v16 = v13
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v21 = v2
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17+v21<<(uint(int32(2))%32))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v26 != int32(1) {
		v34 = v26
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v34 = v26
	goto L4
L11:
	;
	v30 = v21 + int32(1)
	if v30 != v16 {
		v21 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14328(m, l0, l1, l2, int32(107))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_show_grouping_set_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
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
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainOpenGroup(m, int32(_a_F_show_grouping_set_keys_0), v8, int32(1), l6)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = base.B2i32(v24&int32(-2) == int32(2))
	if v24&int32(-2) == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = int32(_a_F_show_grouping_set_keys_1)
	goto L5
L4:
	;
	v36 = int32(_a_F_show_grouping_set_keys_2)
	goto L5
L5:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_ExplainOpenGroup(m, v36, v36, int32(0), l6)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	F_show_sort_group_keys(m, l0, int32(_a_F_show_grouping_set_keys_3), v40, int32(0), v42, v43, v44, v45, l5, l6)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v48 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v49 + int32(1)
	goto L6
L10:
	;
	if v22 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L98
	}
L12:
	;
	F_ExplainCloseGroup(m, v36, int32(0), l6)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L93
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v58 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v24&int32(-2) == int32(2) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = int32(_a_F_show_grouping_set_keys_4)
	goto L17
L16:
	;
	v63 = int32(_a_F_show_grouping_set_keys_5)
	goto L17
L17:
	;
	v75 = v8
	goto L18
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v75<<(uint(int32(2))%32))))
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L12
L20:
	;
	v432 = v75 + int32(1)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v432 < v433 {
		v75 = v432
		goto L18
	} else {
		goto L92
	}
L21:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	switch v220 {
	case 0, 1:
		goto L51
	case 2:
		goto L53
	case 3:
		goto L52
	default:
		goto L50
	}
L22:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v85 < v87 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v200 != 0 {
		v204 = int32(0)
		goto L21
	} else {
		goto L48
	}
L25:
	;
	v90 = v85
	v91 = v85
	goto L28
L26:
	;
	v167 = v85
	goto L27
L27:
	;
	if v167 != 0 {
		v204 = v167
		goto L21
	} else {
		goto L47
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v91<<(uint(int32(2))%32))))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v111<<(uint(int32(1))%32)))))
	if v106 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v167 = v161
	goto L27
L30:
	;
	if v153 == int32(0) {
		goto L11
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v119 <= int32(0) {
		v153 = int32(0)
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v153 = int32(0)
	goto L31
L35:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v119
	goto L38
L37:
	;
	v125 = v122
	goto L38
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v130 = int32(0)
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126+v130<<(uint(int32(2))%32))))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+8)))
	if v139 == v115&int32(_a_F_show_grouping_set_keys_6) {
		v153 = v138
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v142 = v130 + int32(1)
	if v142 != v125 {
		v130 = v142
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v159 = F_deparse_expression(m, v157, l3, l4, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v161 = F_lappend(m, v90, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v164 = v91 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v164 < v165 {
		v90 = v161
		v91 = v164
		goto L28
	} else {
		goto L46
	}
L46:
	;
	goto L29
L47:
	;
	goto L24
L48:
	;
	F_ExplainPropertyText(m, v63, int32(_a_F_show_grouping_set_keys_7), l6)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L20
L50:
	;
	goto L20
L51:
	;
	F_ExplainPropertyList(m, v63, v204, l6)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L91
	}
L52:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v312 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v223 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v230, int32(10))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v224, int32(44))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = int32(1)
	goto L54
L58:
	;
	goto L54
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v234, v235<<(uint(int32(1))%32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v240, int32(91))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v204 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v306, int32(93))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L72
	}
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v246 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	F_escape_json(m, v249, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v254 = int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v255 <= v254 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v259 = v254
	goto L67
L67:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v275, int32(_a_F_show_grouping_set_keys_8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v274+v259<<(uint(int32(2))%32))))
	F_escape_json(m, v279, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v287 = v259 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v287 < v288 {
		v259 = v287
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L50
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v327, int32(_a_F_show_grouping_set_keys_9))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L79
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = int32(1)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v317, int32(10))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v321, v322<<(uint(int32(1))%32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	if v204 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v393, int32(93))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L90
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v334 <= int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	F_escape_json(m, v337, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v342 <= int32(1) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v346 = int32(1)
	goto L85
L85:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v362, int32(_a_F_show_grouping_set_keys_8))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L80
L87:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v361+v346<<(uint(int32(2))%32))))
	F_escape_json(m, v366, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v374 = v346 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v374 < v375 {
		v346 = v374
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	goto L50
L91:
	;
	goto L50
L92:
	;
	goto L19
L93:
	;
	if l2 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_grouping_set_keys_0), int32(1), l6)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v456 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v457 - int32(1)
	goto L94
L97:
	;
	m.G0 = v19 + int32(16)
	return
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v115
	F_errmsg_internal(m, int32(_a_F_show_grouping_set_keys_10), v19)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_show_grouping_set_keys_11), int32(2739), int32(_a_F_show_grouping_set_keys_12))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
