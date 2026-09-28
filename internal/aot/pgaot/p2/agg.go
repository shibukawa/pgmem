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
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v503 int64
	_ = v503
	var v505 int64
	_ = v505
	var v507 int64
	_ = v507
	var v509 int64
	_ = v509
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int64
	_ = v562
	var v564 int64
	_ = v564
	var v566 int64
	_ = v566
	var v568 int64
	_ = v568
	var v570 int64
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v826 int64
	_ = v826
	var v828 int64
	_ = v828
	var v830 int64
	_ = v830
	var v832 int64
	_ = v832
	var v834 int64
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	v26 = F_palloc0(m, int32(72))
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
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(386)
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v32
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = l0
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
	v81 = v78 + v66*int32(240)
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
	v792 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v792
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(1)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v796 == v792 {
		goto L159
	} else {
		goto L160
	}
L18:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v160 = v157 + v152*int32(240)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+224))
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
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449)+10)))
	if v450 != int32(1) {
		v518 = v439
		goto L84
	} else {
		goto L85
	}
L21:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v369 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L22:
	;
	v323 = int32(0)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v324 <= v323 {
		v437 = v323
		v439 = v238
		v440 = v241
		goto L20
	} else {
		goto L66
	}
L23:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v160)+232))
	F_ExecInitExprRec(m, v245, v26, v246+int32(24), v246+int32(32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
	if v239 != 0 {
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
		v238 = v168
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
	v172 = v161 + int32(40)
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
	F_ExecInitExprRec(m, v177, v26, v172, v161+int32(48))
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
	v437 = v182
	v439 = v182
	v440 = v172
	goto L20
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(43)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v190 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v214 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v213 + v214
	v219 = v212 + v213*int32(40)
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+32)) = v220
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+24)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+16)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v219)+8)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v219))) = v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v234 = F_lappend_int(m, int32(0), v231-v214)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v210
	v212 = v210
	goto L35
L37:
	;
	v193 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v193
	v197 = F_palloc_mul(m, int32(40), v193)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v199 != v190 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v210 = v197
	goto L36
L41:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v212 = v201
	goto L35
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v190 << (uint(int32(1)) % 32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v208 = F_repalloc(m, v205, v190*int32(80))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v210 = v208
	goto L36
L45:
	;
	v238 = v234
	goto L24
L46:
	;
	v241 = v161 + int32(40)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+32))
	if v243 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v437 = int32(0)
	v439 = v238
	v440 = v241
	goto L20
L48:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+48)) = uint8(v253)
	*(*int64)(unsafe.Add(mBase, uint32(v246)+40)) = int64(0)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+98)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v161 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v172
	if v257 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v267 = int32(104)
	goto L51
L50:
	;
	v267 = int32(105)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v269 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v293 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v292 + v293
	v298 = v291 + v292*int32(40)
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+32)) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+24)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+16)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+8)) = v305
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v307
	v309 = int32(0)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+98)))
	if v311 == v293 {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v289
	v291 = v289
	goto L52
L54:
	;
	v272 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v272
	v276 = F_palloc_mul(m, int32(40), v272)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v278 != v269 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v289 = v276
	goto L53
L58:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v291 = v280
	goto L52
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v269 << (uint(int32(1)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v287 = F_repalloc(m, v284, v269*int32(80))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v289 = v287
	goto L53
L62:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v318 = F_lappend_int(m, int32(0), v315-int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v320 = v309
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v53
	v437 = v309
	v439 = v320
	v440 = v172
	goto L20
L65:
	;
	v320 = v318
	goto L64
L66:
	;
	v335 = int32(0)
	goto L67
L67:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v335 == v350 {
		v437 = v323
		v439 = v238
		v440 = v241
		goto L20
	} else {
		goto L69
	}
L68:
	;
	v437 = v323
	v439 = v238
	v440 = v241
	goto L20
L69:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v352+v335<<(uint(int32(2))%32))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v359 = v335 + int32(1)
	v362 = v161 + int32(24) + v359<<(uint(int32(4))%32)
	F_ExecInitExprRec(m, v357, v26, v362, v362+int32(8))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	if v359 < v367 {
		v335 = v359
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+32))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	F_ExecInitExprRec(m, v376, v26, v53, v50)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v160)+192))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+20))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+32))
	if v383 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v437 = v50
	v439 = v238
	v440 = int32(0)
	goto L20
L76:
	;
	v437 = v381
	v439 = v238
	v440 = int32(0)
	goto L20
L77:
	;
	goto L78
L78:
	;
	v387 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v388 <= v387 {
		v437 = v381
		v439 = v238
		v440 = v387
		goto L20
	} else {
		goto L79
	}
L79:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	v398 = int32(0)
	goto L80
L80:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413+v398<<(uint(int32(2))%32))))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	F_ExecInitExprRec(m, v418, v26, v391+v398<<(uint(int32(3))%32), v398+v381)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	v437 = v381
	v439 = v238
	v440 = v387
	goto L20
L82:
	;
	v426 = v398 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v426 < v427 {
		v398 = v426
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v160)+124))
	if v519 <= int32(0) {
		v579 = v518
		goto L107
	} else {
		goto L108
	}
L85:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v453 <= int32(0) {
		v518 = v439
		goto L84
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v437
	if v453 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v466 = int32(107)
	goto L89
L88:
	;
	v466 = int32(106)
	goto L89
L89:
	;
	if v440 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v468 = v466
	goto L92
L91:
	;
	v468 = int32(106)
	goto L92
L92:
	;
	if v437 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v469 = int32(108)
	goto L95
L94:
	;
	v469 = v468
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v471 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v495 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v494 + v495
	v500 = v493 + v494*int32(40)
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v500)+32)) = v501
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v500)+24)) = v503
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v500)+16)) = v505
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v500)+8)) = v507
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v500))) = v509
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v514 = F_lappend_int(m, v439, v511-v495)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L106
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v491
	v493 = v491
	goto L96
L98:
	;
	v474 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v474
	v478 = F_palloc_mul(m, int32(40), v474)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v480 != v471 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v491 = v478
	goto L97
L102:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v493 = v482
	goto L96
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v471 << (uint(int32(1)) % 32)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v489 = F_repalloc(m, v486, v471*int32(80))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v491 = v489
	goto L97
L106:
	;
	v518 = v514
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
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+5)))
	if v522 != 0 {
		v579 = v518
		goto L107
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v160
	if v519 == int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v530 = int32(116)
	goto L112
L111:
	;
	v530 = int32(117)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v530
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	if v532 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v556 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v555 + v556
	v561 = v554 + v555*int32(40)
	v562 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+32)) = v562
	v564 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+24)) = v564
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+16)) = v566
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v561)+8)) = v568
	v570 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v561))) = v570
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v575 = F_lappend_int(m, v518, v572-v556)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L123
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v552
	v554 = v552
	goto L113
L115:
	;
	v535 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v535
	v539 = F_palloc_mul(m, int32(40), v535)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v541 != v532 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v552 = v539
	goto L114
L119:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v554 = v543
	goto L113
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v532 << (uint(int32(1)) % 32)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v550 = F_repalloc(m, v547, v532*int32(80))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v552 = v550
	goto L114
L123:
	;
	v579 = v575
	goto L107
L124:
	;
	v580 = int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v581 <= v580 {
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
	v584 = v580
	goto L129
L128:
	;
	v584 = v581
	goto L129
L129:
	;
	v591 = int32(0)
	goto L130
L130:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v161, v160, v152, v591, v591, int32(0), l4)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L126
L132:
	;
	v612 = v591 + int32(1)
	if v612 != v584 {
		v591 = v612
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	if v579 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L135:
	;
	v636 = int32(0)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v638 != int32(2) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v642 = v641
	goto L138
L137:
	;
	v642 = v636
	goto L138
L138:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v643 <= int32(0) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v651 = v642
	v654 = v636
	goto L140
L140:
	;
	F_ExecBuildAggTransCall(m, v26, l0, v21+int32(-40), v161, v160, v152, v654, v651, int32(1), l4)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L134
L142:
	;
	v671 = int32(1)
	v674 = v654 + v671
	if v674 != v643 {
		v651 = v651 + v671
		v654 = v674
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v769 = v152 + int32(1)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v769 < v770 {
		v152 = v769
		goto L18
	} else {
		goto L156
	}
L145:
	;
	v698 = int32(0)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	if v699 <= v698 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v707 = v698
	goto L147
L147:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v579)+12))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v723+v707<<(uint(int32(2))%32))))
	v730 = v722 + v727*int32(40)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	switch v731 - int32(104) {
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
	v745 = v707 + int32(1)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	if v745 < v746 {
		v707 = v745
		goto L147
	} else {
		goto L155
	}
L150:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+24)) = v742
	goto L149
L151:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+20)) = v740
	goto L149
L152:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+28)) = v738
	goto L149
L153:
	;
	if v731 != int32(43) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+16)) = v736
	goto L149
L155:
	;
	goto L148
L156:
	;
	goto L19
L157:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v819 + int32(1)
	v825 = v818 + v819*int32(40)
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v825)+32)) = v826
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v825)+24)) = v828
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v23)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v825)+16)) = v830
	v832 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v825)+8)) = v832
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v825))) = v834
	v836 = F_jit_compile_expr(m, v26)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L167
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v816
	v818 = v816
	goto L157
L159:
	;
	v799 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v799
	v803 = F_palloc_mul(m, int32(40), v799)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v805 != v796 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v816 = v803
	goto L158
L163:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v818 = v807
	goto L157
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v796 << (uint(int32(1)) % 32)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v814 = F_repalloc(m, v811, v796*int32(80))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v816 = v814
	goto L158
L167:
	;
	if v836 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_ExecReadyInterpretedExpr(m, v26)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
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
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
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
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v25 == int32(0) {
			v28 = int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v28
			v32 = F_palloc_mul(m, int32(40), v28)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v45 = v32
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v45
				v47 = v45
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v49 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48 + v49
				v54 = v47 + v48*int32(40)
				v55 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v55
				v57 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v57
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v59
				v61 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v61
				v63 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v54))) = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v68 = v65 - v49
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
				if v70 == int32(0) {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+191)))
					if v75 == int32(1) {
						if v74&int32(1) == int32(0) {
							v103 = int32(112)
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
							if v85 != 0 {
								v86 = int32(110)
							} else {
								v86 = int32(111)
							}
							v103 = v86
						}
					} else {
						if v74&int32(1) == int32(0) {
							v103 = int32(115)
						} else {
							v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
							if v94 != 0 {
								v95 = int32(113)
							} else {
								v95 = int32(114)
							}
							v103 = v95
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					if v98 == int32(1) {
						v101 = int32(118)
					} else {
						v101 = int32(119)
					}
					v103 = v101
				}
				*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v103
				*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v110 == int32(0) {
					v113 = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v113
					v117 = F_palloc_mul(m, int32(40), v113)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						v130 = v117
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
						v132 = v130
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
						v139 = v132 + v133*int32(40)
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
						v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
						if v68 != int32(-1) {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
						} else {
						}
						return
					}
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v119 != v110 {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v132 = v121
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
						v139 = v132 + v133*int32(40)
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
						v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
						if v68 != int32(-1) {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
						} else {
						}
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v110 << (uint(int32(1)) % 32)
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v128 = F_repalloc(m, v125, v110*int32(80))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							v130 = v128
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
							v132 = v130
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
							v139 = v132 + v133*int32(40)
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
							v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
							if v68 != int32(-1) {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
							} else {
							}
							return
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v34 != v25 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v47 = v36
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v49 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48 + v49
				v54 = v47 + v48*int32(40)
				v55 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v55
				v57 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v57
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v59
				v61 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v61
				v63 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v54))) = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v68 = v65 - v49
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
				if v70 == int32(0) {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+191)))
					if v75 == int32(1) {
						if v74&int32(1) == int32(0) {
							v103 = int32(112)
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
							if v85 != 0 {
								v86 = int32(110)
							} else {
								v86 = int32(111)
							}
							v103 = v86
						}
					} else {
						if v74&int32(1) == int32(0) {
							v103 = int32(115)
						} else {
							v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
							if v94 != 0 {
								v95 = int32(113)
							} else {
								v95 = int32(114)
							}
							v103 = v95
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					if v98 == int32(1) {
						v101 = int32(118)
					} else {
						v101 = int32(119)
					}
					v103 = v101
				}
				*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v103
				*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
				v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v110 == int32(0) {
					v113 = int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v113
					v117 = F_palloc_mul(m, int32(40), v113)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						v130 = v117
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
						v132 = v130
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
						v139 = v132 + v133*int32(40)
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
						v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
						if v68 != int32(-1) {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
						} else {
						}
						return
					}
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					if v119 != v110 {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v132 = v121
						v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
						v139 = v132 + v133*int32(40)
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
						v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
						v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
						v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
						v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
						if v68 != int32(-1) {
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
						} else {
						}
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v110 << (uint(int32(1)) % 32)
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v128 = F_repalloc(m, v125, v110*int32(80))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							v130 = v128
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
							v132 = v130
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
							v139 = v132 + v133*int32(40)
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
							v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
							if v68 != int32(-1) {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
							} else {
							}
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v25 << (uint(int32(1)) % 32)
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v43 = F_repalloc(m, v40, v25*int32(80))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = v43
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v45
					v47 = v45
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v49 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48 + v49
					v54 = v47 + v48*int32(40)
					v55 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v55
					v57 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v57
					v59 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v59
					v61 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v61
					v63 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v54))) = v63
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v68 = v65 - v49
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
					if v70 == int32(0) {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+191)))
						if v75 == int32(1) {
							if v74&int32(1) == int32(0) {
								v103 = int32(112)
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
								if v85 != 0 {
									v86 = int32(110)
								} else {
									v86 = int32(111)
								}
								v103 = v86
							}
						} else {
							if v74&int32(1) == int32(0) {
								v103 = int32(115)
							} else {
								v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
								if v94 != 0 {
									v95 = int32(113)
								} else {
									v95 = int32(114)
								}
								v103 = v95
							}
						}
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
						if v98 == int32(1) {
							v101 = int32(118)
						} else {
							v101 = int32(119)
						}
						v103 = v101
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v103
					*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
					v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v110 == int32(0) {
						v113 = int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v113
						v117 = F_palloc_mul(m, int32(40), v113)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							v130 = v117
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
							v132 = v130
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
							v139 = v132 + v133*int32(40)
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
							v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
							if v68 != int32(-1) {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
							} else {
							}
							return
						}
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v119 != v110 {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v132 = v121
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
							v139 = v132 + v133*int32(40)
							v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
							v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
							v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
							v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
							v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
							if v68 != int32(-1) {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
							} else {
							}
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v110 << (uint(int32(1)) % 32)
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v128 = F_repalloc(m, v125, v110*int32(80))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return
							} else {
								v130 = v128
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
								v132 = v130
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
								v139 = v132 + v133*int32(40)
								v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
								v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
								v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
								v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
								v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
								*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
								if v68 != int32(-1) {
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
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
		v68 = int32(-1)
		v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
		if v70 == int32(0) {
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+10)))
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+191)))
			if v75 == int32(1) {
				if v74&int32(1) == int32(0) {
					v103 = int32(112)
				} else {
					v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
					if v85 != 0 {
						v86 = int32(110)
					} else {
						v86 = int32(111)
					}
					v103 = v86
				}
			} else {
				if v74&int32(1) == int32(0) {
					v103 = int32(115)
				} else {
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+184)))
					if v94 != 0 {
						v95 = int32(113)
					} else {
						v95 = int32(114)
					}
					v103 = v95
				}
			}
		} else {
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
			if v98 == int32(1) {
				v101 = int32(118)
			} else {
				v101 = int32(119)
			}
			v103 = v101
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v103
		*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v18
		v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v110 == int32(0) {
			v113 = int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v113
			v117 = F_palloc_mul(m, int32(40), v113)
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				v130 = v117
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
				v132 = v130
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
				v139 = v132 + v133*int32(40)
				v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
				v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
				v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
				v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
				if v68 != int32(-1) {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
				} else {
				}
				return
			}
		} else {
			v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v119 != v110 {
				v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v132 = v121
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
				v139 = v132 + v133*int32(40)
				v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
				v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
				v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
				v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
				v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
				if v68 != int32(-1) {
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
				} else {
				}
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v110 << (uint(int32(1)) % 32)
				v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v128 = F_repalloc(m, v125, v110*int32(80))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					v130 = v128
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v130
					v132 = v130
					v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v133 + int32(1)
					v139 = v132 + v133*int32(40)
					v140 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v139)+32)) = v140
					v142 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v139)+24)) = v142
					v144 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v139)+16)) = v144
					v146 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v139)+8)) = v146
					v148 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(v139))) = v148
					if v68 != int32(-1) {
						v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v152+v68*int32(40))+20)) = v156
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
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								F_parser_errposition(m, v75, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(822), int32(_a_F_check_agg_arguments_walker_2))
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
					v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(816), int32(_a_F_check_agg_arguments_walker_2))
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(816), int32(_a_F_check_agg_arguments_walker_2))
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
						v106 = F_query_tree_walker_impl(m, l0, int32(511), l1, int32(16))
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
							v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
					v106 = F_query_tree_walker_impl(m, l0, int32(511), l1, int32(16))
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								F_parser_errposition(m, v75, v76)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(822), int32(_a_F_check_agg_arguments_walker_2))
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
					v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(816), int32(_a_F_check_agg_arguments_walker_2))
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
											F_errfinish(m, int32(_a_F_check_agg_arguments_walker_1), int32(816), int32(_a_F_check_agg_arguments_walker_2))
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
						v106 = F_query_tree_walker_impl(m, l0, int32(511), l1, int32(16))
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
							v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
					v106 = F_query_tree_walker_impl(m, l0, int32(511), l1, int32(16))
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
						v115 = F_expression_tree_walker_impl(m, l0, int32(511), l1)
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
