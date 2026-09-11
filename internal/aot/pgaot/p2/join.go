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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
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
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
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
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v706 int32
	_ = v706
	var v734 int32
	_ = v734
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
	v32 = v22
	v33 = v20
	v34 = v6
	v38 = v6
	v39 = v6
	goto L7
L7:
	;
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v356 = int32(0)
	if v39 == v356 {
		v576 = v356
		goto L89
	} else {
		goto L90
	}
L9:
	;
	goto L8
L10:
	;
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v151 == int32(0) {
		goto L9
	} else {
		goto L31
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v138 = v31
	v139 = v32
	v140 = v33
	v142 = v46
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
	v139 = v132
	v140 = v129
	v142 = v132
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
	v155 = v139 + int32(4)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if base.Ui32(v155) < base.Ui32(v142+v157<<(uint(int32(2))%32)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v162 = v155
	goto L34
L33:
	;
	v162 = int32(0)
	goto L34
L34:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v164 = int32(0)
	if v163 == v164 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v217 == int32(0) {
		v31 = v138
		v32 = v162
		v33 = v140
		goto L7
	} else {
		goto L49
	}
L36:
	;
	v217 = int32(1)
	goto L35
L37:
	;
	goto L38
L38:
	;
	if l2 == int32(0) {
		v208 = v164
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v217 = v208
	goto L35
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v174 < v173 {
		v208 = v164
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v176 = int32(1)
	if v173 <= v176 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v179 = v176
	goto L44
L43:
	;
	v179 = v173
	goto L44
L44:
	;
	v180 = int32(8)
	v185 = int32(0)
	goto L45
L45:
	;
	v192 = v185 << (uint(int32(2)) % 32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v163+v180+v192)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+(l2+v180))))
	v199 = v194 & (v196 ^ int32(-1))
	v201 = base.B2i32(v199 == int32(0))
	if v199 != 0 {
		v208 = v201
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v208 = v201
	goto L39
L47:
	;
	v203 = v185 + int32(1)
	if v203 != v179 {
		v185 = v203
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v221 = int32(0)
	if v220 == v221 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v274 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v274 = int32(1)
	goto L50
L52:
	;
	goto L53
L53:
	;
	if l3 == int32(0) {
		v265 = v221
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v274 = v265
	goto L50
L55:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v231 < v230 {
		v265 = v221
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v233 = int32(1)
	if v230 <= v233 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v236 = v233
	goto L59
L58:
	;
	v236 = v230
	goto L59
L59:
	;
	v237 = int32(8)
	v242 = int32(0)
	goto L60
L60:
	;
	v249 = v242 << (uint(int32(2)) % 32)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v220+v237+v249)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+(l3+v237))))
	v256 = v251 & (v253 ^ int32(-1))
	v258 = base.B2i32(v256 == int32(0))
	if v256 != 0 {
		v265 = v258
		goto L54
	} else {
		goto L62
	}
L61:
	;
	v265 = v258
	goto L54
L62:
	;
	v260 = v242 + int32(1)
	if v260 != v236 {
		v242 = v260
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v275 = F_lappend(m, v39, v151)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v280 = int32(0)
	if v279 == v280 {
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
	v32 = v162
	v33 = v140
	v39 = v275
	goto L7
L69:
	;
	if v333 != 0 {
		goto L83
	} else {
		goto L84
	}
L70:
	;
	v333 = int32(1)
	goto L69
L71:
	;
	goto L72
L72:
	;
	if l4 == int32(0) {
		v324 = v280
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v333 = v324
	goto L69
L74:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v290 < v289 {
		v324 = v280
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v292 = int32(1)
	if v289 <= v292 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v295 = v292
	goto L78
L77:
	;
	v295 = v289
	goto L78
L78:
	;
	v296 = int32(8)
	v301 = int32(0)
	goto L79
L79:
	;
	v308 = v301 << (uint(int32(2)) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v279+v296+v308)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+(l4+v296))))
	v315 = v310 & (v312 ^ int32(-1))
	v317 = base.B2i32(v315 == int32(0))
	if v315 != 0 {
		v324 = v317
		goto L73
	} else {
		goto L81
	}
L80:
	;
	v324 = v317
	goto L73
L81:
	;
	v319 = v301 + int32(1)
	if v319 != v295 {
		v301 = v319
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v334 = F_lappend(m, v34, v151)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L67
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v336 = F_lappend(m, v38, v151)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L67
	} else {
		goto L87
	}
L86:
	;
	v31 = v138
	v32 = v162
	v33 = v140
	v34 = v334
	goto L7
L87:
	;
	v31 = v138
	v32 = v162
	v33 = v140
	v38 = v336
	goto L7
L88:
	;
	v734 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)) = uint8(v734)
	return int32(0)
L89:
	;
	if v38 == int32(0) {
		v706 = v576
		goto L138
	} else {
		goto L139
	}
L90:
	;
	if v34 == int32(0) {
		v576 = v356
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v361 <= int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v370 = int32(-1)
	v371 = v364
	v380 = v6
	v381 = v6
	v382 = v6
	v383 = v6
	goto L94
L93:
	;
	v563 = F_create_join_clause(m, l0, l1, v561, v559, v560, l1)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L67
	} else {
		goto L136
	}
L94:
	;
	if int32(0) < v371 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v524 < int32(0) {
		goto L88
	} else {
		goto L135
	}
L96:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386+v380<<(uint(int32(2))%32))))
	v396 = v370
	v401 = int32(0)
	v407 = v381
	v408 = v382
	v409 = v383
	goto L99
L97:
	;
	v524 = v370
	v525 = v371
	v535 = v381
	v536 = v382
	v537 = v383
	goto L98
L98:
	;
	v539 = v380 + int32(1)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v539 < v540 {
		v370 = v524
		v371 = v525
		v380 = v539
		v381 = v535
		v382 = v536
		v383 = v537
		goto L94
	} else {
		goto L134
	}
L99:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v410 == int32(0) {
		v500 = v396
		v511 = v407
		v512 = v408
		v513 = v409
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v500 == int32(3) {
		v559 = v511
		v560 = v512
		v561 = v513
		goto L93
	} else {
		goto L133
	}
L101:
	;
	v515 = v401 + int32(1)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v515 < v516 {
		v396 = v500
		v401 = v515
		v407 = v511
		v408 = v512
		v409 = v513
		goto L99
	} else {
		goto L132
	}
L102:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v413 <= int32(0) {
		v500 = v396
		v511 = v407
		v512 = v408
		v513 = v409
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416+v401<<(uint(int32(2))%32))))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+16))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v429 = int32(0)
	goto L104
L104:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v442+v429<<(uint(int32(2))%32))))
	v448 = F_get_opfamily_member_for_cmptype(m, v446, v422, v421, int32(3))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L67
	} else {
		goto L107
	}
L105:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	if v462 != int32(6) {
		goto L117
	} else {
		goto L118
	}
L106:
	;
	goto L105
L107:
	;
	if v448 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v450 == int32(0) {
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v458 = v429 + int32(1)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v458 < v459 {
		v429 = v458
		goto L104
	} else {
		goto L115
	}
L111:
	;
	v453 = F_get_opcode(m, v448)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L67
	} else {
		goto L112
	}
L112:
	;
	v455 = F_get_func_leakproof(m, v453)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L67
	} else {
		goto L113
	}
L113:
	;
	if v455 != 0 {
		goto L106
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v500 = v396
	v511 = v407
	v512 = v408
	v513 = v409
	goto L101
L116:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)))
	if v476 != int32(6) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	v465 = int32(0)
	if v462 != int32(27) {
		v474 = v465
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v474 = int32(1)
	goto L116
L120:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	if v469 != int32(6) {
		v474 = v465
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v488 = F_exprType(m, v461)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L67
	} else {
		goto L128
	}
L123:
	;
	if v476 != int32(27) {
		v487 = v474
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v487 = v474 + int32(1)
	goto L122
L126:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	if v482 != int32(6) {
		v487 = v474
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v490 = F_op_hashjoinable(m, v448, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L67
	} else {
		goto L129
	}
L129:
	;
	v492 = v490 + v487
	if v492 <= v396 {
		v500 = v396
		v511 = v407
		v512 = v408
		v513 = v409
		goto L101
	} else {
		goto L130
	}
L130:
	;
	if v492 != int32(3) {
		v500 = v492
		v511 = v390
		v512 = v420
		v513 = v448
		goto L101
	} else {
		goto L131
	}
L131:
	;
	v559 = v390
	v560 = v420
	v561 = v448
	goto L93
L132:
	;
	goto L100
L133:
	;
	v524 = v500
	v525 = v516
	v535 = v511
	v536 = v512
	v537 = v513
	goto L98
L134:
	;
	goto L95
L135:
	;
	v559 = v535
	v560 = v536
	v561 = v537
	goto L93
L136:
	;
	v565 = F_lappend(m, int32(0), v563)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L67
	} else {
		goto L137
	}
L137:
	;
	v576 = v565
	goto L89
L138:
	;
	return v706
L139:
	;
	v587 = F_list_concat(m, v39, v34)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L67
	} else {
		goto L140
	}
L140:
	;
	if v587 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v591 = F_lappend(m, v38, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L67
	} else {
		goto L144
	}
L142:
	;
	v595 = v38
	goto L143
L143:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v596 <= int32(0) {
		v706 = v576
		goto L138
	} else {
		goto L146
	}
L144:
	;
	if v591 == int32(0) {
		v706 = v576
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v595 = v591
	goto L143
L146:
	;
	v599 = int32(0)
	v603 = v599
	v609 = v599
	v610 = v576
	goto L147
L147:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v619+v603<<(uint(int32(2))%32))))
	if v609 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v706 = v684
	goto L138
L149:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v624 == int32(0) {
		goto L88
	} else {
		goto L152
	}
L150:
	;
	v684 = v610
	goto L151
L151:
	;
	v694 = v603 + int32(1)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v694 < v695 {
		v603 = v694
		v609 = v623
		v610 = v684
		goto L147
	} else {
		goto L168
	}
L152:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v624)+4))
	if v627 <= int32(0) {
		goto L88
	} else {
		goto L153
	}
L153:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v623)+16))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v609)+16))
	v638 = int32(0)
	goto L154
L154:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v651+v638<<(uint(int32(2))%32))))
	v657 = F_get_opfamily_member_for_cmptype(m, v655, v631, v630, int32(3))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L67
	} else {
		goto L157
	}
L155:
	;
	v671 = F_create_join_clause(m, l0, l1, v657, v609, v623, int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L67
	} else {
		goto L166
	}
L156:
	;
	goto L155
L157:
	;
	if v657 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v659 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v667 = v638 + int32(1)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v624)+4))
	if v667 < v668 {
		v638 = v667
		goto L154
	} else {
		goto L165
	}
L161:
	;
	v662 = F_get_opcode(m, v657)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L67
	} else {
		goto L162
	}
L162:
	;
	v664 = F_get_func_leakproof(m, v662)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L67
	} else {
		goto L163
	}
L163:
	;
	if v664 != 0 {
		goto L156
	} else {
		goto L164
	}
L164:
	;
	goto L160
L165:
	;
	goto L88
L166:
	;
	v673 = F_lappend(m, v610, v671)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L67
	} else {
		goto L167
	}
L167:
	;
	v684 = v673
	goto L151
L168:
	;
	goto L148
}
