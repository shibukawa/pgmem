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
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v98 int32
	_ = v98
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
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
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int64
	_ = v496
	var v498 int64
	_ = v498
	var v500 int64
	_ = v500
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int64
	_ = v556
	var v558 int64
	_ = v558
	var v560 int64
	_ = v560
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v819 int64
	_ = v819
	var v821 int64
	_ = v821
	var v823 int64
	_ = v823
	var v825 int64
	_ = v825
	var v827 int64
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
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
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v32
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = l0
	v50 = v26 + int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v50
	v53 = v26 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v55 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v66 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ExecPushExprSetupSteps(m, v26, v21+int32(-56))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v81 = v78 + v66*int32(224)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v85 = v21 + int32(-56)
	v86 = F_expr_setup_walker(m, v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+32))
	v90 = F_expr_setup_walker(m, v89, v85)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+36))
	v94 = F_expr_setup_walker(m, v93, v85)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	v98 = F_expr_setup_walker(m, v97, v85)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	v102 = F_expr_setup_walker(m, v101, v85)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v105 = v66 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v105 < v106 {
		v66 = v105
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v132 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v136 = v42 & int32(1)
	v152 = v6
	goto L18
L16:
	;
	goto L17
L17:
	;
	v786 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v786
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(1)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v790 == v786 {
		goto L159
	} else {
		goto L160
	}
L18:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v160 = v157 + v152*int32(224)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+212))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+44))
	if v136 != 0 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	goto L17
L20:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+10)))
	if v446 != int32(1) {
		v513 = v435
		goto L84
	} else {
		goto L85
	}
L21:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v367 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L22:
	;
	v321 = int32(0)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v322 <= v321 {
		v433 = v321
		v435 = v237
		v436 = v240
		goto L20
	} else {
		goto L66
	}
L23:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v160)+220))
	F_ExecInitExprRec(m, v244, v26, v245+int32(20), v245+int32(24))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
	if v238 != 0 {
		goto L21
	} else {
		goto L46
	}
L25:
	;
	v165 = int32(0)
	goto L27
L26:
	;
	v165 = v164
	goto L27
L27:
	;
	if v165 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v168 = int32(0)
	if v136 == v168 {
		v237 = v168
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_ExecInitExprRec(m, v164, v26, v53, v50)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v172 = v161 + int32(28)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163)+32))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v160)+24))
	if v176 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	F_ExecInitExprRec(m, v177, v26, v172, v161+int32(32))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v182 = int32(0)
	v433 = v182
	v435 = v182
	v436 = v172
	goto L20
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(43)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v190 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v213 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v212 + v213
	v218 = v211 + v212*int32(40)
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v218)+32)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v218)+24)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v218)+16)) = v223
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v218)+8)) = v225
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v218))) = v227
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v233 = F_lappend_int(m, int32(0), v230-v213)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v209
	v211 = v209
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v196 = F_palloc(m, int32(640))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v198 != v190 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v209 = v196
	goto L36
L41:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v211 = v200
	goto L35
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v190 << (uint(int32(1)) % 32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v207 = F_repalloc(m, v204, v190*int32(80))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v209 = v207
	goto L36
L45:
	;
	v237 = v233
	goto L24
L46:
	;
	v240 = v161 + int32(28)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+32))
	if v242 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v433 = int32(0)
	v435 = v237
	v436 = v240
	goto L20
L48:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+32)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+28)) = v252
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+98)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v161 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v172
	if v256 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v266 = int32(104)
	goto L51
L50:
	;
	v266 = int32(105)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v268 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v291 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v290 + v291
	v296 = v289 + v290*int32(40)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+32)) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+24)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+16)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v296)+8)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v296))) = v305
	v307 = int32(0)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+98)))
	if v309 == v291 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v287
	v289 = v287
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v274 = F_palloc(m, int32(640))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v276 != v268 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v287 = v274
	goto L53
L58:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v289 = v278
	goto L52
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v268 << (uint(int32(1)) % 32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v285 = F_repalloc(m, v282, v268*int32(80))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v287 = v285
	goto L53
L62:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v316 = F_lappend_int(m, int32(0), v313-int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v318 = v307
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v53
	v433 = v307
	v435 = v318
	v436 = v172
	goto L20
L65:
	;
	v318 = v316
	goto L64
L66:
	;
	v333 = int32(0)
	goto L67
L67:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v333 == v348 {
		v433 = v321
		v435 = v237
		v436 = v240
		goto L20
	} else {
		goto L69
	}
L68:
	;
	v433 = v321
	v435 = v237
	v436 = v240
	goto L20
L69:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350+v333<<(uint(int32(2))%32))))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v357 = v333 + int32(1)
	v360 = v161 + int32(20) + v357<<(uint(int32(3))%32)
	F_ExecInitExprRec(m, v355, v26, v360, v360+int32(4))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v357 < v365 {
		v333 = v357
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+32))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	F_ExecInitExprRec(m, v374, v26, v53, v50)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v160)+188))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+20))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+32))
	if v381 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v433 = v50
	v435 = v237
	v436 = int32(0)
	goto L20
L76:
	;
	v433 = v379
	v435 = v237
	v436 = int32(0)
	goto L20
L77:
	;
	goto L78
L78:
	;
	v385 = int32(0)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v386 <= v385 {
		v433 = v379
		v435 = v237
		v436 = v385
		goto L20
	} else {
		goto L79
	}
L79:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	v396 = int32(0)
	goto L80
L80:
	;
	v412 = v396 << (uint(int32(2)) % 32)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v412+v413)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	F_ExecInitExprRec(m, v416, v26, v389+v412, v396+v379)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v433 = v379
	v435 = v237
	v436 = v385
	goto L20
L82:
	;
	v422 = v396 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	if v422 < v423 {
		v396 = v422
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v160)+124))
	if v514 <= int32(0) {
		v573 = v513
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v449 <= int32(0) {
		v513 = v435
		goto L84
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v433
	if v449 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v462 = int32(107)
	goto L89
L88:
	;
	v462 = int32(106)
	goto L89
L89:
	;
	if v436 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v464 = v462
	goto L92
L91:
	;
	v464 = int32(106)
	goto L92
L92:
	;
	if v433 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v465 = int32(108)
	goto L95
L94:
	;
	v465 = v464
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v465
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v467 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v490 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v489 + v490
	v495 = v488 + v489*int32(40)
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v495)+32)) = v496
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v495)+24)) = v498
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v495)+16)) = v500
	v502 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v495)+8)) = v502
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v495))) = v504
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v509 = F_lappend_int(m, v435, v506-v490)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L106
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v486
	v488 = v486
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v473 = F_palloc(m, int32(640))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v475 != v467 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v486 = v473
	goto L97
L102:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v488 = v477
	goto L96
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v467 << (uint(int32(1)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v484 = F_repalloc(m, v481, v467*int32(80))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v486 = v484
	goto L97
L106:
	;
	v513 = v509
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
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
	if v517 != 0 {
		v573 = v513
		goto L107
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v160
	if v514 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v525 = int32(116)
	goto L112
L111:
	;
	v525 = int32(117)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v525
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v527 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v550 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v549 + v550
	v555 = v548 + v549*int32(40)
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v555)+32)) = v556
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v555)+24)) = v558
	v560 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v555)+16)) = v560
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v555)+8)) = v562
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v555))) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v569 = F_lappend_int(m, v513, v566-v550)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L123
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v546
	v548 = v546
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v533 = F_palloc(m, int32(640))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v535 != v527 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v546 = v533
	goto L114
L119:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v548 = v537
	goto L113
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v527 << (uint(int32(1)) % 32)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v544 = F_repalloc(m, v541, v527*int32(80))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v546 = v544
	goto L114
L123:
	;
	v573 = v569
	goto L107
L124:
	;
	v574 = int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v575 <= v574 {
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
	v578 = v574
	goto L129
L128:
	;
	v578 = v575
	goto L129
L129:
	;
	v585 = int32(0)
	goto L130
L130:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v161, v160, v152, v585, v585, int32(0), l4)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L126
L132:
	;
	v606 = v585 + int32(1)
	if v606 != v578 {
		v585 = v606
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	if v573 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L135:
	;
	v630 = int32(0)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v632 != int32(2) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v636 = v635
	goto L138
L137:
	;
	v636 = v630
	goto L138
L138:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v637 <= int32(0) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v645 = v636
	v648 = v630
	goto L140
L140:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v161, v160, v152, v648, v645, int32(1), l4)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L134
L142:
	;
	v665 = int32(1)
	v668 = v648 + v665
	if v668 != v637 {
		v645 = v645 + v665
		v648 = v668
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v763 = v152 + int32(1)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v763 < v764 {
		v152 = v763
		goto L18
	} else {
		goto L156
	}
L145:
	;
	v692 = int32(0)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v693 <= v692 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v701 = v692
	goto L147
L147:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v717+v701<<(uint(int32(2))%32))))
	v724 = v716 + v721*int32(40)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	switch v725 - int32(104) {
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
	v739 = v701 + int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v739 < v740 {
		v701 = v739
		goto L147
	} else {
		goto L155
	}
L150:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+24)) = v736
	goto L149
L151:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+20)) = v734
	goto L149
L152:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+28)) = v732
	goto L149
L153:
	;
	if v725 != int32(43) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+16)) = v730
	goto L149
L155:
	;
	goto L148
L156:
	;
	goto L19
L157:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v812 + int32(1)
	v818 = v811 + v812*int32(40)
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+32)) = v819
	v821 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+24)) = v821
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+16)) = v823
	v825 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+8)) = v825
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v818))) = v827
	v829 = F_jit_compile_expr(m, v26)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L167
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v809
	v811 = v809
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(16)
	v796 = F_palloc(m, int32(640))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	if v798 != v790 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v809 = v796
	goto L158
L163:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v811 = v800
	goto L157
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v790 << (uint(int32(1)) % 32)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v807 = F_repalloc(m, v804, v790*int32(80))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v809 = v807
	goto L158
L167:
	;
	if v829 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_ExecReadyInterpretedExpr(m, v26)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
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
							F_errmsg(m, int32(_a_F_check_agg_arguments_walker_0), int32(0))
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
									F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(821), int32(_a_F_check_agg_arguments_walker_2))
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
								F_errmsg(m, int32(_a_F_check_agg_arguments_walker_3), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_check_agg_arguments_walker_4), int32(0))
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(815), int32(_a_F_check_agg_arguments_walker_2))
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
								F_errmsg(m, int32(_a_F_check_agg_arguments_walker_3), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_check_agg_arguments_walker_4), int32(0))
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(815), int32(_a_F_check_agg_arguments_walker_2))
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
							F_errmsg(m, int32(_a_F_check_agg_arguments_walker_0), int32(0))
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
									F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(821), int32(_a_F_check_agg_arguments_walker_2))
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
								F_errmsg(m, int32(_a_F_check_agg_arguments_walker_3), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_check_agg_arguments_walker_4), int32(0))
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(815), int32(_a_F_check_agg_arguments_walker_2))
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
								F_errmsg(m, int32(_a_F_check_agg_arguments_walker_3), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_check_agg_arguments_walker_4), int32(0))
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(815), int32(_a_F_check_agg_arguments_walker_2))
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
