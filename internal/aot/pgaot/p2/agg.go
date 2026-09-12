package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildAggTrans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v508 int64
	_ = v508
	var v510 int64
	_ = v510
	var v512 int64
	_ = v512
	var v514 int64
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v829 int64
	_ = v829
	var v831 int64
	_ = v831
	var v833 int64
	_ = v833
	var v835 int64
	_ = v835
	var v837 int64
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	v26 = F_palloc0(m, int32(68))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(380)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v32
	v39 = v21 + int32(-32)
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v32
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = l0
	v52 = v26 + int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v52
	v55 = v26 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v57 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v68 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ExecPushExprSetupSteps(m, v26, v21+int32(-56))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v83 = v80 + v68*int32(224)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v88 = F_expr_setup_walker(m, v85, v21+int32(-56))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	v94 = F_expr_setup_walker(m, v91, v21+int32(-56))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+36))
	v100 = F_expr_setup_walker(m, v97, v21+int32(-56))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
	v106 = F_expr_setup_walker(m, v103, v21+int32(-56))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+44))
	v112 = F_expr_setup_walker(m, v109, v21+int32(-56))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v115 = v68 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v115 < v116 {
		v68 = v115
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v142 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v146 = v44 & int32(1)
	v162 = v6
	goto L18
L16:
	;
	goto L17
L17:
	;
	v796 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v796
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(1)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v800 == v796 {
		goto L159
	} else {
		goto L160
	}
L18:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v170 = v167 + v162*int32(224)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+212))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+44))
	if v146 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	goto L17
L20:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+10)))
	if v456 != int32(1) {
		v523 = v445
		goto L84
	} else {
		goto L85
	}
L21:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	if v377 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L22:
	;
	v331 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v332 <= v331 {
		v445 = v247
		v447 = v250
		v448 = v331
		goto L20
	} else {
		goto L66
	}
L23:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v170)+220))
	F_ExecInitExprRec(m, v254, v26, v255+int32(20), v255+int32(24))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
	if v248 != 0 {
		goto L21
	} else {
		goto L46
	}
L25:
	;
	v175 = int32(0)
	goto L27
L26:
	;
	v175 = v174
	goto L27
L27:
	;
	if v175 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v178 = int32(0)
	if v146 == v178 {
		v247 = v178
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_ExecInitExprRec(m, v174, v26, v55, v52)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v182 = v171 + int32(28)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v170)+24))
	if v186 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	F_ExecInitExprRec(m, v187, v26, v182, v171+int32(32))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v192 = int32(0)
	v445 = v192
	v447 = v182
	v448 = v192
	goto L20
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(43)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v200 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v223 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v222 + v223
	v228 = v221 + v222*int32(40)
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+32)) = v229
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+24)) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+16)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v243 = F_lappend_int(m, int32(0), v240-v223)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v219
	v221 = v219
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v206 = F_palloc(m, int32(640))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v208 != v200 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v219 = v206
	goto L36
L41:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v221 = v210
	goto L35
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v200 << (uint(int32(1)) % 32)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v217 = F_repalloc(m, v214, v200*int32(80))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v219 = v217
	goto L36
L45:
	;
	v247 = v243
	goto L24
L46:
	;
	v250 = v171 + int32(28)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+32))
	if v252 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v445 = v247
	v447 = v250
	v448 = int32(0)
	goto L20
L48:
	;
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+32)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+28)) = v262
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+98)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v171 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v182
	if v266 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v276 = int32(104)
	goto L51
L50:
	;
	v276 = int32(105)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v278 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v301 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v300 + v301
	v306 = v299 + v300*int32(40)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+32)) = v307
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+24)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+16)) = v311
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v315
	v317 = int32(0)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+98)))
	if v319 == v301 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v297
	v299 = v297
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v284 = F_palloc(m, int32(640))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v286 != v278 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v297 = v284
	goto L53
L58:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v299 = v288
	goto L52
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v278 << (uint(int32(1)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v295 = F_repalloc(m, v292, v278*int32(80))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v297 = v295
	goto L53
L62:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v326 = F_lappend_int(m, int32(0), v323-int32(1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v328 = v317
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v55
	v445 = v328
	v447 = v182
	v448 = v317
	goto L20
L65:
	;
	v328 = v326
	goto L64
L66:
	;
	v343 = int32(0)
	goto L67
L67:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	if v343 == v358 {
		v445 = v247
		v447 = v250
		v448 = v331
		goto L20
	} else {
		goto L69
	}
L68:
	;
	v445 = v247
	v447 = v250
	v448 = v331
	goto L20
L69:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360+v343<<(uint(int32(2))%32))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v367 = v343 + int32(1)
	v370 = v171 + int32(20) + v367<<(uint(int32(3))%32)
	F_ExecInitExprRec(m, v365, v26, v370, v370+int32(4))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v367 < v375 {
		v343 = v367
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+32))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	F_ExecInitExprRec(m, v384, v26, v55, v52)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v170)+188))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)+20))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+32))
	if v391 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v445 = v247
	v447 = int32(0)
	v448 = v52
	goto L20
L76:
	;
	v445 = v247
	v447 = int32(0)
	v448 = v389
	goto L20
L77:
	;
	goto L78
L78:
	;
	v395 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v396 <= v395 {
		v445 = v247
		v447 = v395
		v448 = v389
		goto L20
	} else {
		goto L79
	}
L79:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	v406 = int32(0)
	goto L80
L80:
	;
	v422 = v406 << (uint(int32(2)) % 32)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422+v423)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	F_ExecInitExprRec(m, v426, v26, v422+v399, v406+v389)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v445 = v247
	v447 = v395
	v448 = v389
	goto L20
L82:
	;
	v432 = v406 + int32(1)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v432 < v433 {
		v406 = v432
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v170)+124))
	if v524 <= int32(0) {
		v583 = v523
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	if v459 <= int32(0) {
		v523 = v445
		goto L84
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v448
	if v459 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v472 = int32(107)
	goto L89
L88:
	;
	v472 = int32(106)
	goto L89
L89:
	;
	if v447 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v474 = v472
	goto L92
L91:
	;
	v474 = int32(106)
	goto L92
L92:
	;
	if v448 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v475 = int32(108)
	goto L95
L94:
	;
	v475 = v474
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v477 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v500 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v499 + v500
	v505 = v498 + v499*int32(40)
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v505)+32)) = v506
	v508 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v505)+24)) = v508
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v505)+16)) = v510
	v512 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v505)+8)) = v512
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v505))) = v514
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v519 = F_lappend_int(m, v445, v516-v500)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L106
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v496
	v498 = v496
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v483 = F_palloc(m, int32(640))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v485 != v477 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v496 = v483
	goto L97
L102:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v498 = v487
	goto L96
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v477 << (uint(int32(1)) % 32)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v494 = F_repalloc(m, v491, v477*int32(80))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v496 = v494
	goto L97
L106:
	;
	v523 = v519
	goto L84
L107:
	;
	if l2 != 0 {
		goto L124
	} else {
		goto L125
	}
L108:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+5)))
	if v527 != 0 {
		v583 = v523
		goto L107
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v170
	if v524 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v535 = int32(116)
	goto L112
L111:
	;
	v535 = int32(117)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v537 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v560 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v559 + v560
	v565 = v558 + v559*int32(40)
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+32)) = v566
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+24)) = v568
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+16)) = v570
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+8)) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v565))) = v574
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v579 = F_lappend_int(m, v523, v576-v560)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L123
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v556
	v558 = v556
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v543 = F_palloc(m, int32(640))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v545 != v537 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v556 = v543
	goto L114
L119:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v558 = v547
	goto L113
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v537 << (uint(int32(1)) % 32)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v554 = F_repalloc(m, v551, v537*int32(80))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v556 = v554
	goto L114
L123:
	;
	v583 = v579
	goto L107
L124:
	;
	v584 = int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v585 <= v584 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	if l3 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L127:
	;
	v588 = v584
	goto L129
L128:
	;
	v588 = v585
	goto L129
L129:
	;
	v595 = int32(0)
	goto L130
L130:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v171, v170, v162, v595, v595, int32(0), l4)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L126
L132:
	;
	v616 = v595 + int32(1)
	if v616 != v588 {
		v595 = v616
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	if v583 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L135:
	;
	v640 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v642 != int32(2) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v646 = v645
	goto L138
L137:
	;
	v646 = v640
	goto L138
L138:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v647 <= int32(0) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v655 = v646
	v659 = v640
	goto L140
L140:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v171, v170, v162, v659, v655, int32(1), l4)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L134
L142:
	;
	v675 = int32(1)
	v678 = v659 + v675
	if v678 != v647 {
		v655 = v655 + v675
		v659 = v678
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v773 = v162 + int32(1)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v773 < v774 {
		v162 = v773
		goto L18
	} else {
		goto L156
	}
L145:
	;
	v702 = int32(0)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	if v703 <= v702 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v711 = v702
	goto L147
L147:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v711<<(uint(int32(2))%32))))
	v734 = v726 + v731*int32(40)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	switch v735 - int32(104) {
	case 0:
		goto L151
	case 1, 5, 6, 7, 8, 9, 10, 11:
		goto L149
	case 2, 3, 4:
		goto L152
	case 12, 13:
		goto L150
	default:
		goto L153
	}
L148:
	;
	goto L144
L149:
	;
	v749 = v711 + int32(1)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	if v749 < v750 {
		v711 = v749
		goto L147
	} else {
		goto L155
	}
L150:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+24)) = v746
	goto L149
L151:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+20)) = v744
	goto L149
L152:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+28)) = v742
	goto L149
L153:
	;
	if v735 != int32(43) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v734)+16)) = v740
	goto L149
L155:
	;
	goto L148
L156:
	;
	goto L19
L157:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v822 + int32(1)
	v828 = v821 + v822*int32(40)
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+32)) = v829
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+24)) = v831
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+16)) = v833
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v828)+8)) = v835
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v828))) = v837
	v839 = F_jit_compile_expr(m, v26)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L167
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v819
	v821 = v819
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v806 = F_palloc(m, int32(640))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v808 != v800 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v819 = v806
	goto L158
L163:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v821 = v810
	goto L157
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v800 << (uint(int32(1)) % 32)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v817 = F_repalloc(m, v814, v800*int32(80))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v819 = v817
	goto L158
L167:
	;
	if v839 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_ExecReadyInterpretedExpr(m, v26)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	m.G0 = v23 - int32(-64)
	return v26
L171:
	;
	goto L170
}
func F_ExecBuildAggTransCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	if l8 != 0 {
		v17 = l1 + int32(156)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
		v17 = v13 + l6<<(uint(int32(2))%32)
	}
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if l9 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(109)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v25 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
			v31 = F_palloc(m, int32(640))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v44 = v31
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
				v46 = v44
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v48 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v47 + v48
				v53 = v46 + v47*int32(40)
				v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+32)) = v54
				v56 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v56
				v58 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v58
				v60 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v60
				v62 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v53))) = v62
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v67 = v64 - v48
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
				if v69 == int32(0) {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+187)))
					if v74 == int32(1) {
						if v73&int32(1) == int32(0) {
							v102 = int32(112)
						} else {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
							if v84 != 0 {
								v85 = int32(110)
							} else {
								v85 = int32(111)
							}
							v102 = v85
						}
					} else {
						if v73&int32(1) == int32(0) {
							v102 = int32(115)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
							if v93 != 0 {
								v94 = int32(113)
							} else {
								v94 = int32(114)
							}
							v102 = v94
						}
					}
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					if v97 == int32(1) {
						v100 = int32(118)
					} else {
						v100 = int32(119)
					}
					v102 = v100
				}
				*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
				*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v109 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
					v115 = F_palloc(m, int32(640))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v128 = v115
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
						v130 = v128
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
						v137 = v130 + v131*int32(40)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
						if v67 != int32(-1) {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
						} else {
						}
						return
					}
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v117 != v109 {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v130 = v119
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
						v137 = v130 + v131*int32(40)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
						if v67 != int32(-1) {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
						} else {
						}
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v109 << (uint(int32(1)) % 32)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v126 = F_repalloc(m, v123, v109*int32(80))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v128 = v126
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
							v130 = v128
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
							v137 = v130 + v131*int32(40)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
							if v67 != int32(-1) {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
							} else {
							}
							return
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v33 != v25 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v46 = v35
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v48 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v47 + v48
				v53 = v46 + v47*int32(40)
				v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+32)) = v54
				v56 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v56
				v58 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v58
				v60 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v60
				v62 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v53))) = v62
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v67 = v64 - v48
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
				if v69 == int32(0) {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+187)))
					if v74 == int32(1) {
						if v73&int32(1) == int32(0) {
							v102 = int32(112)
						} else {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
							if v84 != 0 {
								v85 = int32(110)
							} else {
								v85 = int32(111)
							}
							v102 = v85
						}
					} else {
						if v73&int32(1) == int32(0) {
							v102 = int32(115)
						} else {
							v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
							if v93 != 0 {
								v94 = int32(113)
							} else {
								v94 = int32(114)
							}
							v102 = v94
						}
					}
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					if v97 == int32(1) {
						v100 = int32(118)
					} else {
						v100 = int32(119)
					}
					v102 = v100
				}
				*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
				*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v109 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
					v115 = F_palloc(m, int32(640))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return
					} else {
						v128 = v115
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
						v130 = v128
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
						v137 = v130 + v131*int32(40)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
						if v67 != int32(-1) {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
						} else {
						}
						return
					}
				} else {
					v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v117 != v109 {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v130 = v119
						v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
						v137 = v130 + v131*int32(40)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
						if v67 != int32(-1) {
							v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
						} else {
						}
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v109 << (uint(int32(1)) % 32)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v126 = F_repalloc(m, v123, v109*int32(80))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							v128 = v126
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
							v130 = v128
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
							v137 = v130 + v131*int32(40)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
							if v67 != int32(-1) {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
							} else {
							}
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v25 << (uint(int32(1)) % 32)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v42 = F_repalloc(m, v39, v25*int32(80))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = v42
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
					v46 = v44
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v48 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v47 + v48
					v53 = v46 + v47*int32(40)
					v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+32)) = v54
					v56 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+24)) = v56
					v58 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v58
					v60 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v60
					v62 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v53))) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v67 = v64 - v48
					v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
					if v69 == int32(0) {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+187)))
						if v74 == int32(1) {
							if v73&int32(1) == int32(0) {
								v102 = int32(112)
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
								if v84 != 0 {
									v85 = int32(110)
								} else {
									v85 = int32(111)
								}
								v102 = v85
							}
						} else {
							if v73&int32(1) == int32(0) {
								v102 = int32(115)
							} else {
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
								if v93 != 0 {
									v94 = int32(113)
								} else {
									v94 = int32(114)
								}
								v102 = v94
							}
						}
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						if v97 == int32(1) {
							v100 = int32(118)
						} else {
							v100 = int32(119)
						}
						v102 = v100
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
					*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v109 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
						v115 = F_palloc(m, int32(640))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v128 = v115
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
							v130 = v128
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
							v137 = v130 + v131*int32(40)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
							if v67 != int32(-1) {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
							} else {
							}
							return
						}
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						if v117 != v109 {
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v130 = v119
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
							v137 = v130 + v131*int32(40)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
							if v67 != int32(-1) {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
							} else {
							}
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v109 << (uint(int32(1)) % 32)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v126 = F_repalloc(m, v123, v109*int32(80))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return
							} else {
								v128 = v126
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
								v130 = v128
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
								v137 = v130 + v131*int32(40)
								v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
								v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
								v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
								v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
								v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
								if v67 != int32(-1) {
									v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
								} else {
								}
								return
							}
						}
					}
				}
			}
		}
	} else {
		v67 = int32(-1)
		v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
		if v69 == int32(0) {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+187)))
			if v74 == int32(1) {
				if v73&int32(1) == int32(0) {
					v102 = int32(112)
				} else {
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
					if v84 != 0 {
						v85 = int32(110)
					} else {
						v85 = int32(111)
					}
					v102 = v85
				}
			} else {
				if v73&int32(1) == int32(0) {
					v102 = int32(115)
				} else {
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+180)))
					if v93 != 0 {
						v94 = int32(113)
					} else {
						v94 = int32(114)
					}
					v102 = v94
				}
			}
		} else {
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			if v97 == int32(1) {
				v100 = int32(118)
			} else {
				v100 = int32(119)
			}
			v102 = v100
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
		*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
		v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v109 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(16)
			v115 = F_palloc(m, int32(640))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				v128 = v115
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
				v130 = v128
				v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
				v137 = v130 + v131*int32(40)
				v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
				v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
				v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
				v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
				if v67 != int32(-1) {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
				} else {
				}
				return
			}
		} else {
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v117 != v109 {
				v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v130 = v119
				v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
				v137 = v130 + v131*int32(40)
				v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
				v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
				v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
				v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
				if v67 != int32(-1) {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
				} else {
				}
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v109 << (uint(int32(1)) % 32)
				v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v126 = F_repalloc(m, v123, v109*int32(80))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return
				} else {
					v128 = v126
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v128
					v130 = v128
					v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v131 + int32(1)
					v137 = v130 + v131*int32(40)
					v138 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v137)+32)) = v138
					v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v137)+24)) = v140
					v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v137)+16)) = v142
					v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = v144
					v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v137))) = v146
					if v67 != int32(-1) {
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v150+v67*int32(40))+20)) = v154
					} else {
					}
					return
				}
			}
		}
	}
}
func F_check_agg_arguments_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v3 = int32(0)
	if l0 == v3 {
		v120 = v3
		return v120
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(6) {
		case 0:
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v13 = v11 - v12
			if v13 < int32(0) {
				v120 = v3
				return v120
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v16) <= base.Ui32(v13) {
					v120 = v3
					return v120
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v13
					return int32(0)
				}
			}
		default:
			v31 = v8
			if v31 == int32(10) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v40 = v38 - v39
				if v40 < int32(0) {
					v49 = int32(10)
					v50 = v39
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v44) <= base.Ui32(v40) {
						v49 = int32(10)
						v50 = v39
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v40
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v49 = v48
						v50 = v39
					}
				}
			} else {
				v35 = v31
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v49 = v35
				v50 = v37
			}
			if v50 == int32(0) {
				switch v49 - int32(11) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50364548))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(152898), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_parser_errposition(m, v75, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498669), int32(821), int32(221495))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 1, 2, 3, 5:
					v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v120 = v115
						return v120
					}
				case 4:
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					if v113 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(153026), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(621903), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v138 = F_exprLocation(m, l0)
										mBase = m.M
										F_parser_errposition(m, v137, v138)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(498669), int32(815), int32(221495))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					}
				case 6:
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					if v59 != int32(1) {
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(153026), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(621903), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v138 = F_exprLocation(m, l0)
										mBase = m.M
										F_parser_errposition(m, v137, v138)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(498669), int32(815), int32(221495))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
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
					}
				default:
					if v49 == int32(67) {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v50 + int32(1)
						v106 = F_query_tree_walker_impl(m, l0, int32(476), l1, int32(16))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v108 - int32(1)
							return v106
						}
					} else {
						if v49 == int32(101) {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v88 != int32(6) {
								v120 = v3
								return v120
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								v92 = v91 - v50
								if v92 < int32(0) {
									v120 = v3
									return v120
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									if base.Ui32(v95) <= base.Ui32(v92) {
										v120 = v3
										return v120
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l0
										*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v92
										return int32(0)
									}
								}
							}
						} else {
							v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v120 = v115
								return v120
							}
						}
					}
				}
			} else {
				if v49 == int32(67) {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v50 + int32(1)
					v106 = F_query_tree_walker_impl(m, l0, int32(476), l1, int32(16))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v108 - int32(1)
						return v106
					}
				} else {
					if v49 != int32(101) {
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v88 != int32(6) {
							v120 = v3
							return v120
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							v92 = v91 - v50
							if v92 < int32(0) {
								v120 = v3
								return v120
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if base.Ui32(v95) <= base.Ui32(v92) {
									v120 = v3
									return v120
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v92
									return int32(0)
								}
							}
						}
					}
				}
			}
		case 3:
			v21 = int32(9)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v24 = v22 - v23
			if v24 < int32(0) {
				v35 = v21
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v49 = v35
				v50 = v37
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if base.Ui32(v27) <= base.Ui32(v24) {
					v35 = v21
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v49 = v35
					v50 = v37
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v24
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v31 = v30
					if v31 == int32(10) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v40 = v38 - v39
						if v40 < int32(0) {
							v49 = int32(10)
							v50 = v39
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if base.Ui32(v44) <= base.Ui32(v40) {
								v49 = int32(10)
								v50 = v39
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v40
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = v48
								v50 = v39
							}
						}
					} else {
						v35 = v31
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v49 = v35
						v50 = v37
					}
				}
			}
			if v50 == int32(0) {
				switch v49 - int32(11) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50364548))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(152898), int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								F_parser_errposition(m, v75, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498669), int32(821), int32(221495))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 1, 2, 3, 5:
					v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v120 = v115
						return v120
					}
				case 4:
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					if v113 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(153026), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(621903), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v138 = F_exprLocation(m, l0)
										mBase = m.M
										F_parser_errposition(m, v137, v138)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(498669), int32(815), int32(221495))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					}
				case 6:
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
					if v59 != int32(1) {
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(153026), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(621903), int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v138 = F_exprLocation(m, l0)
										mBase = m.M
										F_parser_errposition(m, v137, v138)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(498669), int32(815), int32(221495))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int32(0)
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
					}
				default:
					if v49 == int32(67) {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v50 + int32(1)
						v106 = F_query_tree_walker_impl(m, l0, int32(476), l1, int32(16))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v108 - int32(1)
							return v106
						}
					} else {
						if v49 == int32(101) {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v88 != int32(6) {
								v120 = v3
								return v120
							} else {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
								v92 = v91 - v50
								if v92 < int32(0) {
									v120 = v3
									return v120
								} else {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									if base.Ui32(v95) <= base.Ui32(v92) {
										v120 = v3
										return v120
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l0
										*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v92
										return int32(0)
									}
								}
							}
						} else {
							v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v120 = v115
								return v120
							}
						}
					}
				}
			} else {
				if v49 == int32(67) {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v50 + int32(1)
					v106 = F_query_tree_walker_impl(m, l0, int32(476), l1, int32(16))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v108 - int32(1)
						return v106
					}
				} else {
					if v49 != int32(101) {
						v115 = F_expression_tree_walker_impl(m, l0, int32(476), l1)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = v115
							return v120
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v88 != int32(6) {
							v120 = v3
							return v120
						} else {
							v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
							v92 = v91 - v50
							if v92 < int32(0) {
								v120 = v3
								return v120
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if base.Ui32(v95) <= base.Ui32(v92) {
									v120 = v3
									return v120
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l0
									*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v92
									return int32(0)
								}
							}
						}
					}
				}
			}
		}
	}
}
