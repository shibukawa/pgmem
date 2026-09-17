package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ComputeIndexAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	v18 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(128)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l7 == v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l14 != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	if l13 == int32(0) {
		v69 = v18
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v53 = v34 << (uint(int32(2)) % 32)
	v54 = F_palloc(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L10
	}
L5:
	;
	v40 = v34 << (uint(int32(2)) % 32)
	v41 = F_palloc(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v41
	v44 = F_palloc(m, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v44
	v49 = F_palloc(m, v34<<(uint(int32(1))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v49
	v69 = v18
	goto L1
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v54
	v57 = F_palloc(m, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v57
	v62 = F_palloc(m, v34<<(uint(int32(1))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v62
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	if l13 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v67 = int32(0)
	goto L15
L14:
	;
	v67 = v66
	goto L15
L15:
	;
	v69 = v67
	goto L1
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(124)))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(120)))) = v78
	goto L19
L17:
	;
	goto L18
L18:
	;
	if l6 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L200
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L6
	} else {
		goto L193
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L6
	} else {
		goto L187
	}
L23:
	;
	m.G0 = v32 + int32(128)
	return
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v82 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v88 = l0 + int32(12)
	v111 = v18
	v114 = v69
	goto L26
L26:
	;
	v119 = v111 << (uint(int32(2)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+v120)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v123 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v119))) = v281
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	if v34 <= v111 {
		goto L77
	} else {
		goto L78
	}
L29:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+22)))
	v255 = v253 + v254
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v255)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v88+v111<<(uint(int32(1))%32)))) = uint16(v256)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+96))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+68))
	F_ReleaseCatCache(m, v124)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L67
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v133
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_0), v32+int32(80))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L65
	}
L31:
	;
	v124 = F_SearchSysCacheAttName(m, l8, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v111 < v34 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	if v124 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if l12 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v133
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_1), v32+int32(96))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(1960), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v147 = F_exprType(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L61
	}
L44:
	;
	v149 = F_exprCollation(m, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v168 = v146
	goto L46
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v180 != int32(31) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v180 != int32(6) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v168 = v222
	goto L46
L51:
	;
	v196 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v88+v111<<(uint(int32(1))%32)))) = uint16(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v199 = F_lappend(m, v198, v168)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L54
	}
L52:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+8)))
	if v185 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v88+v111<<(uint(int32(1))%32)))) = uint16(v185)
	v281 = v147
	v282 = v149
	goto L28
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v199
	v202 = F_contain_mutable_functions_after_planning(m, v168)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	if v202 == int32(0) {
		v281 = v147
		v282 = v149
		goto L28
	} else {
		goto L56
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_4), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2019), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_5), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(1978), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(1955), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v281 = v259
	v282 = v258
	goto L28
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637+v119))) = v633
	v641 = v111 + int32(1)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v641 < v642 {
		v111 = v641
		v114 = v636
		goto L26
	} else {
		goto L186
	}
L69:
	;
	v587 = l5 + v111<<(uint(int32(1))%32)
	v588 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v587))) = uint16(v588)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	if l11 != 0 {
		goto L169
	} else {
		goto L170
	}
L70:
	;
	v548 = F_get_commutator(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L157
	}
L71:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l16)))
	F_AtEOXact_GUC(m, int32(0), v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
	} else {
		goto L151
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L146
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L6
	} else {
		goto L142
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L138
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L6
	} else {
		goto L134
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L6
	} else {
		goto L130
	}
L77:
	;
	if v293 != 0 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v293 == int32(0) {
		v345 = v282
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	if v295 != 0 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	if v296 != 0 {
		goto L74
	} else {
		goto L82
	}
L82:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	if v297 != 0 {
		goto L73
	} else {
		goto L83
	}
L83:
	;
	v298 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v119))) = v298
	*(*int32)(unsafe.Add(mBase, uint32(l4+v119))) = v298
	*(*uint16)(unsafe.Add(mBase, uint32(l5+v111<<(uint(int32(1))%32)))) = uint16(v298)
	v633 = v298
	v636 = v114
	v637 = l2
	goto L68
L84:
	;
	v346 = F_type_is_collatable(m, v281)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L97
	}
L85:
	;
	if l14 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v315 = F_get_collation_oid(m, v293, int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l16)))
	F_AtEOXact_GUC(m, int32(0), v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L90
	}
L89:
	;
	v345 = v315
	goto L84
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l15
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l14
	goto L91
L91:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v327 = F_get_collation_oid(m, v325, int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v32)+124))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v32)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v330
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v329
	goto L93
L93:
	;
	v336 = int32(_a_F_ComputeIndexAttrs_6)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v340 = v338 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v340
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v340
	F_RestrictSearchPath(m)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v345 = v327
	goto L84
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v119))) = v345
	if l14 != 0 {
		goto L108
	} else {
		goto L109
	}
L97:
	;
	if v346 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v345 != 0 {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v345 != 0 {
		goto L72
	} else {
		goto L107
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_7), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_errhint(m, int32(_a_F_ComputeIndexAttrs_8), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2090), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	goto L96
L108:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l16)))
	F_AtEOXact_GUC(m, int32(0), v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v378 = l3 + v119
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	v380 = F_ResolveOpClass(m, v379, v281, l9, l10)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l15
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l14
	goto L112
L112:
	;
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = v380
	if l14 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v406 = int32(0)
	if l13 == v406 {
		v584 = v406
		goto L69
	} else {
		goto L124
	}
L115:
	;
	if v114 == int32(0) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v32)+124))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v32)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v391
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v390
	goto L120
L118:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v388 = F_compatible_oper_opid(m, v387, v281, v281)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	v547 = v388
	goto L70
L120:
	;
	v397 = int32(_a_F_ComputeIndexAttrs_6)
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v401 = v399 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v401
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v401
	F_RestrictSearchPath(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	if v114 != 0 {
		goto L71
	} else {
		goto L123
	}
L123:
	;
	goto L114
L124:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if v111 == v34-int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v414 = int32(7)
	goto L127
L126:
	;
	v414 = int32(3)
	goto L127
L127:
	;
	F_GetOperatorFromCompareType(m, v409, int32(0), v414, v32+int32(112), v32+int32(118))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v421+v119))) = v423
	v425 = F_get_opcode(m, v423)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v427+v119))) = v425
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+118)))
	*(*uint16)(unsafe.Add(mBase, uint32(v430+v111<<(uint(int32(1))%32)))) = uint16(v434)
	v584 = v406
	goto L69
L130:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_9), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2034), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_10), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2038), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_11), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2042), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_12), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2046), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	v507 = F_format_type_be(m, v281)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v507
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_13), v32-int32(-64))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2098), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l15
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l14
	goto L152
L152:
	;
	v529 = F_compatible_oper_opid(m, v520, v281, v281)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v32)+124))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v32)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v532
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v531
	goto L154
L154:
	;
	v538 = int32(_a_F_ComputeIndexAttrs_6)
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v542 = v540 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v542
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l16))) = v542
	F_RestrictSearchPath(m)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v547 = v529
	goto L70
L157:
	;
	if v548 != v547 {
		goto L22
	} else {
		goto L158
	}
L158:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v552 = F_get_opclass_family(m, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	v554 = F_get_op_opfamily_strategy(m, v547, v552)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	if v554 == int32(0) {
		goto L21
	} else {
		goto L161
	}
L161:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v558+v119))) = v547
	v561 = F_get_opcode(m, v547)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v563+v119))) = v561
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*uint16)(unsafe.Add(mBase, uint32(v566+v111<<(uint(int32(1))%32)))) = uint16(v554)
	v572 = v114 + int32(4)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if base.Ui32(v572) < base.Ui32(v574+v575<<(uint(int32(2))%32)) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v580 = v572
	goto L165
L164:
	;
	v580 = int32(0)
	goto L165
L165:
	;
	v584 = v580
	goto L69
L166:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v122)+24))
	if v623 != 0 {
		goto L182
	} else {
		goto L183
	}
L167:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v587))) = uint16(v597)
	goto L166
L168:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v122)+28))
	if v618 != int32(2) {
		goto L166
	} else {
		goto L181
	}
L169:
	;
	v591 = int32(2)
	if v590 == v591 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	if v590 != 0 {
		goto L20
	} else {
		goto L175
	}
L172:
	;
	v594 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v587))) = uint16(v594)
	v597 = int32(3)
	goto L174
L173:
	;
	v597 = v591
	goto L174
L174:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	switch v598 {
	case 0:
		goto L168
	case 1:
		goto L167
	default:
		goto L166
	}
L175:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v122)+32))
	if v599 == int32(0) {
		goto L166
	} else {
		goto L176
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = l9
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_14), v32)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2234), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L6
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	goto L167
L182:
	;
	v624 = int32(0)
	v629 = F_transformRelOptions(m, v624, v623, v624, v624, v624, v624)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L6
	} else {
		goto L185
	}
L183:
	;
	v632 = int32(0)
	goto L184
L184:
	;
	v633 = v632
	v636 = v584
	v637 = l4
	goto L68
L185:
	;
	v632 = v629
	goto L184
L186:
	;
	goto L27
L187:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	v683 = F_format_operator(m, v547)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v683
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_15), v32+int32(48))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	F_errdetail(m, int32(_a_F_ComputeIndexAttrs_16), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2166), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	v707 = F_format_operator(m, v547)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	v709 = F_get_opfamily_name(m, v552)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v707
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_17), v32+int32(32))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errdetail(m, int32(_a_F_ComputeIndexAttrs_18), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2179), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = l9
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_19), v32+int32(16))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2229), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCheckIndexConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(176)
	m.G0 = v20
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v30 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v37 = v30
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l1
	if int32(0) < v29 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	v37 = v33
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L61
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L56
	}
L8:
	;
	m.G0 = v20 + int32(176)
	return v203
L9:
	;
	v42 = int32(0)
	v49 = v7
	goto L12
L10:
	;
	v177 = v7
	goto L11
L11:
	;
	if l5 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L12:
	;
	v60 = v42 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v28+v60)))
	if v62 == int32(0) {
		v164 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v177 = v164
	goto L11
L14:
	;
	v168 = v42 + int32(1)
	if v168 != v29 {
		v42 = v168
		v49 = v164
		goto L12
	} else {
		goto L51
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+v27)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+116)))
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+92))
	if v70 == int32(0) {
		v164 = v49
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+118)))
	if v73 != int32(1) {
		v164 = v49
		goto L14
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	if l5 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = int32(0)
	if l5 == v78 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v62)+192))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+16)))
	if v120 == int32(0) {
		goto L7
	} else {
		goto L38
	}
L24:
	;
	if v116 == int32(0) {
		v164 = v49
		goto L14
	} else {
		goto L37
	}
L25:
	;
	v116 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v84 <= int32(0) {
		v110 = v78
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = v110
	goto L24
L29:
	;
	v87 = int32(0)
	if v87 < v84 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v90 = v84
	goto L32
L31:
	;
	v90 = v87
	goto L32
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v93 = int32(0)
	goto L33
L33:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91+v93<<(uint(int32(2))%32))))
	v102 = base.B2i32(v101 == v77)
	if v101 == v77 {
		v110 = v102
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v110 = v102
	goto L28
L35:
	;
	v104 = v93 + int32(1)
	if v104 != v90 {
		v93 = v104
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L23
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v66)+84))
	if v123 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v151 = v20 + int32(32)
	F_FormIndexDatum(m, v66, l1, l2, v151, v20)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L48
	}
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	if v126 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v129 = F_ExecPrepareQual(m, v123, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v134 = v126
	goto L43
L43:
	;
	v135 = int32(_a_F_ExecCheckIndexConstraints_0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0])) = v138
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	v143 = m.T0[v142].(func(*base.Module, int32, int32, int32) int32)(m, v134, v37, v20+int32(175))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+88)) = v129
	if v129 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v129
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0])) = v136
	if v143 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v164 = int32(1)
	goto L14
L48:
	;
	v154 = int32(0)
	v155 = int32(1)
	v159 = F_check_exclusion_or_unique_constraint(m, v26, v62, v66, l4, v151, v20, l2, v154, v154, v155, l3)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v159 == int32(0) {
		v203 = v154
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v164 = v155
	goto L14
L51:
	;
	goto L13
L52:
	;
	v203 = int32(1)
	goto L8
L53:
	;
	goto L54
L54:
	;
	v190 = int32(1)
	if v177&v190 == int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v203 = v190
	goto L8
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_ExecCheckIndexConstraints_1), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	F_errtableconstraint(m, v26, v227+int32(4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ExecCheckIndexConstraints_2), int32(610), int32(_a_F_ExecCheckIndexConstraints_3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F_ExecCheckIndexConstraints_4), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_ExecCheckIndexConstraints_2), int32(656), int32(_a_F_ExecCheckIndexConstraints_3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetFreeIndexPage(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = F_GetPageWithFreeSpace(m, l0, int32(_a_F_GetFreeIndexPage_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != int32(-1) {
			F_RecordPageWithFreeSpace(m, l0, v4, int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v4
			}
		} else {
			return v4
		}
	}
}
func F_IndexOnlyRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(_a_F_IndexOnlyRecheck_0), int32(0))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_IndexOnlyRecheck_1), int32(328), int32(_a_F_IndexOnlyRecheck_2))
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_index_beginscan_parallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	v10 = F_RestoreSnapshot(m, l5+int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_RegisterSnapshot(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = F_index_beginscan_internal(m, l1, l3, l4, v10, l5, int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v10
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
				v24 = m.T0[v23].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v24
					return v17
				}
			}
		}
	}
}
func F_index_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v210 int64
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int64
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
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
	var v323 int64
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int64
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
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
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
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
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v423 float64
	_ = v423
	var v425 int32
	_ = v425
	var v427 float64
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
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
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
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
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	v15 = m.G0
	v17 = v15 - int32(272)
	m.G0 = v17
	if l4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	v37 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[0]))
	if v22 != int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+25)))
	if v26 != int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v31 = F_plan_create_index_workers(m, v29, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+128)) = v31
	goto L1
L7:
	;
	if v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(92)))) = v89
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(88)))) = v92
	goto L18
L9:
	;
	F_errfinish(m, int32(_a_F_index_build_0), v78, int32(_a_F_index_build_1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L17
	}
L10:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v46 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v45 + v46
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v44 + v46
	F_errmsg_internal(m, int32(_a_F_index_build_2), v17+int32(48))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v78 = int32(3037)
	goto L9
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v62
	v64 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v61 + v64
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v60 + v64
	F_errmsg_internal(m, int32(_a_F_index_build_3), v17-int32(-64))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v78 = int32(3043)
	goto L9
L17:
	;
	goto L8
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+80))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[2])) = v96 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[1])) = v95
	goto L19
L19:
	;
	v104 = int32(_a_F_index_build_4)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[3]))
	v108 = v106 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[3])) = v108
	goto L20
L20:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v113 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v113
	v116 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v116
	v119 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v119
	v121 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v17)+152)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v17)+144)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v17)+136)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = int64(2)
	v135 = v17 + int32(96)
	v137 = v17 + int32(128)
	goto L24
L22:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+36))
	v313 = m.T0[v312].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L39
	}
L23:
	;
	goto L22
L24:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[7]))
	if v147 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_build[8])))
	if v151&int32(1) == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v156 = int32(_a_F_index_build_5)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[9]))
	v159 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[9])) = v158 + v159
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v162 + v159
	goto L28
L27:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v293 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v292 + v293
	v296 = int32(_a_F_index_build_5)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[9])) = v298 - v293
	goto L23
L28:
	;
	v171 = v147 + int32(232)
	goto L29
L29:
	;
	v177 = int32(0)
	v180 = int32(0)
	goto L32
L31:
	;
	v257 = int32(0)
	v260 = v241
	goto L36
L32:
	;
	v186 = int32(2)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v135+v180<<(uint(v186)%32))))
	v190 = int32(3)
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v137+v180<<(uint(v190)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v171+v189<<(uint(v190)%32)))) = v196
	v199 = v180 | int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v135+v199<<(uint(v186)%32))))
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v137+v199<<(uint(v190)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v171+v203<<(uint(v190)%32)))) = v210
	v213 = v180 | v186
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v135+v213<<(uint(v186)%32))))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v137+v213<<(uint(v190)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v171+v217<<(uint(v190)%32)))) = v224
	v227 = v180 | v190
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v135+v227<<(uint(v186)%32))))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v137+v227<<(uint(v190)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v171+v231<<(uint(v190)%32)))) = v238
	v240 = int32(4)
	v241 = v180 + v240
	v243 = v177 + v240
	if v243 != int32(4) {
		v177 = v243
		v180 = v241
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L35
L34:
	;
	goto L33
L35:
	;
	goto L31
L36:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v135+v260<<(uint(int32(2))%32))))
	v270 = int32(3)
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v137+v260<<(uint(v270)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v171+v269<<(uint(v270)%32)))) = v276
	v278 = int32(1)
	v281 = v257 + v278
	if v281 != int32(2) {
		v257 = v281
		v260 = v260 + v278
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L27
L38:
	;
	goto L37
L39:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+118)))
	if v316 != int32(117) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if l3 != 0 {
		goto L64
	} else {
		goto L65
	}
L41:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v319 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v345 = v319
	goto L44
L43:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v323
	v327 = F_smgropen(m, v17+int32(32), v320)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	v347 = F_smgrexists(m, v345, int32(3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L50
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v327
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327)+72))
	if v331 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v345 = v343
	goto L44
L47:
	;
	v339 = v331
	goto L49
L48:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327)+76))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v327)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v327)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v327)+72))
	v339 = v337
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+72)) = v339 + int32(1)
	goto L46
L50:
	;
	if v347 != 0 {
		goto L40
	} else {
		goto L51
	}
L51:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v349 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v375 = v349
	goto L54
L53:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v351
	v353 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v353
	v357 = F_smgropen(m, v17+int32(16), v350)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L55
	}
L54:
	;
	F_smgrcreate(m, v375, int32(3), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L5
	} else {
		goto L60
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v357
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357)+72))
	if v361 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v375 = v373
	goto L54
L57:
	;
	v369 = v361
	goto L59
L58:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v357)+76))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v363
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v357)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v357)+72))
	v369 = v367
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+72)) = v369 + int32(1)
	goto L56
L60:
	;
	F_log_smgrcreate(m, l1, int32(3))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+40))
	m.T0[v384].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L40
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L127
	}
L64:
	;
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v313)))
	F_index_update_stats(m, l0, int32(1), v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L74
	}
L65:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)))
	if v388&int32(1) == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v393 != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v397 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v401 = F_SearchSysCacheCopy(m, int32(34), v394, int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	if v401 == int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)+16))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+22)))
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v405+v406)+19)) = uint8(v408)
	F_CatalogTupleUpdate(m, v397, v401+int32(4), v401)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v401)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	F_relation_close(m, v397, int32(3))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	goto L64
L74:
	;
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v313)+8))
	F_index_update_stats(m, l1, int32(0), v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v432 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[10]))
	if v433 == v435 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	F_AtEOXact_GUC(m, int32(0), v108)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L5
	} else {
		goto L125
	}
L80:
	;
	v438 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[10])) = v438
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[11])) = v438
	goto L82
L81:
	;
	goto L82
L82:
	;
	v443 = F_CreateExecutorState(m)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v443)+152))
	if v445 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v448 = F_MakePerTupleExprContext(m, v443)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L87
	}
L85:
	;
	v450 = v445
	goto L86
L86:
	;
	v452 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L88
	}
L87:
	;
	v450 = v448
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v450)+4)) = v452
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v456 = F_ExecPrepareQual(m, v455, v443)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	v458 = F_GetLatestSnapshot(m)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v460 = F_RegisterSnapshot(m, v458)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	v462 = int32(0)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+8))
	v468 = m.T0[v467].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v460, v462, v462, v462, int32(449))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v452)+36)) = v471
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[12]))
	if v474 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+188))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+12))
	m.T0[v578].(func(*base.Module, int32))(m, v468)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L5
	} else {
		goto L121
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L5
	} else {
		goto L118
	}
L95:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_build[13])))
	if v476&int32(1) == int32(0) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L99
L98:
	;
	goto L97
L99:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+188))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+20))
	v501 = m.T0[v500].(func(*base.Module, int32, int32, int32) int32)(m, v468, int32(1), v452)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L5
	} else {
		goto L101
	}
L100:
	;
	goto L94
L101:
	;
	if v501 == int32(0) {
		goto L93
	} else {
		goto L102
	}
L102:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[14]))
	if v506 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L5
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v456 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L105
L107:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v452)+36)) = v539
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[12]))
	if v542 == int32(0) {
		goto L99
	} else {
		goto L116
	}
L108:
	;
	v509 = int32(_a_F_index_build_6)
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[15]))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v450)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[15])) = v512
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v456)+20))
	v517 = m.T0[v516].(func(*base.Module, int32, int32, int32) int32)(m, v456, v450, v17+int32(271))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v525 = v17 + int32(128)
	v527 = v17 + int32(96)
	F_FormIndexDatum(m, l2, v452, v443, v525, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[15])) = v510
	if v517 == int32(0) {
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	F_check_exclusion_constraint(m, l0, l1, l2, v452+int32(28), v525, v527, v443, int32(1))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v450)+20))
	F_MemoryContextReset(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	goto L107
L116:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_build[13])))
	if v546&int32(1) != 0 {
		goto L99
	} else {
		goto L117
	}
L117:
	;
	goto L100
L118:
	;
	F_errmsg_internal(m, int32(_a_F_index_build_7), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_index_build_8), int32(1034), int32(_a_F_index_build_9))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_UnregisterSnapshot(m, v460)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	F_ExecDropSingleTupleTableSlot(m, v452)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_FreeExecutorState(m, v443)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v587 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v587
	goto L79
L125:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[2])) = v609
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[1])) = v608
	goto L126
L126:
	;
	m.G0 = v17 + int32(272)
	return
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v394
	F_errmsg_internal(m, int32(_a_F_index_build_10), v17)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_index_build_0), int32(3139), int32(_a_F_index_build_1))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_bulk_delete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_index_bulk_delete[0]))
	if v15 != v13 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_bulk_delete[1]))
		v19 = F_list_member_ptr(m, v18, v13)
		mBase = m.M
		v21 = v19
	} else {
		v21 = int32(1)
	}
	if v21 == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
		if v25 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_index_bulk_delete_0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v59 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_bulk_delete_1), v10+int32(16))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_bulk_delete_2), int32(803), int32(_a_F_index_bulk_delete_3))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = m.T0[v25].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(32)
				return v28
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43 + int32(4)
				F_errmsg(m, int32(_a_F_index_bulk_delete_4), v10)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_bulk_delete_2), int32(802), int32(_a_F_index_bulk_delete_3))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
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
func F_index_constraint_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	if l8 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L19
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L61
	}
L4:
	;
	if l6 != int32(120) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v20 = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if base.Ui32(v21) < base.Ui32(int32(_a_F_index_constraint_create_8)) {
		v30 = v20
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v30 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L7:
	;
	goto L6
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v25 == int32(99) {
		v30 = v20
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v28 = F_isTempToastNamespace(m, v25)
	mBase = m.M
	v30 = v28
	goto L7
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_index_constraint_create[1]))
	if v34 == int32(2) {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L4
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+76))
	if v39 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if l7&int32(16) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v42 = int32(1259)
	v45 = F_deleteDependencyRecordsForClass(m, v42, l2, v42, int32(97))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v48 = l7 & int32(2)
	v49 = int32(0)
	v52 = l7 & int32(4)
	v55 = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v69 = int32(32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v78 = base.B2i32(l3 == v49)
	v85 = F_CreateConstraintEntry(m, l5, v18, l6, base.B2i32(v48 != v49), base.B2i32(v52 != v49), v55, v55, l3, v57, l4+int32(12), v60, v61, v49, l2, v49, v49, v49, v49, v49, v49, v69, v69, v49, v49, v69, v74, v49, v49, v78, base.B2i32(l3 != v49), v78, base.B2i32(l7&v69 != v49), l9)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L19
	} else {
		goto L21
	}
L19:
	;
	return
L20:
	;
	goto L18
L21:
	;
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1259)
	F_recordDependencyOn(m, v15+int32(20), l0, int32(105))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(2606)
	v108 = v15 + int32(8)
	F_recordDependencyOn(m, l0, v108, int32(80))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v48 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(1259)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v114
	F_recordDependencyOn(m, l0, v108, int32(83))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v124 = F_palloc0(m, int32(52))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v172 = int32(0)
	if base.B2i32(l7&int32(8) == v172)|base.B2i32(l7&int32(3) == v172) == v172 {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = int32(0)
	v128 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)) = uint16(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(181)
	if l6 == int32(112) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v136 = int32(_a_F_index_constraint_create_3)
	goto L34
L33:
	;
	v136 = int32(_a_F_index_constraint_create_4)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v136
	v139 = F_SystemFuncName(m, int32(_a_F_index_constraint_create_5))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+48)) = v141
	v144 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+45)) = uint8(v144)
	v146 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+44)) = uint8(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+40)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v124)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+26)) = int32(_a_F_index_constraint_create_6)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)) = uint8(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = v139
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_CreateTrigger(m, v15+int32(8), v124, v141, v162, v141, v85, l2, v141, v146)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	v183 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L19
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	m.G0 = v15 + int32(32)
	return
L40:
	;
	v185 = int32(0)
	v188 = F_SearchSysCacheCopy(m, int32(34), l2, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v188 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+22)))
	v194 = v192 + v193
	if l7&int32(1) == int32(0) {
		v203 = v185
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v48 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+14)))
	if v199 != 0 {
		v203 = v185
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+14)) = uint8(v200)
	v203 = v200
	goto L43
L46:
	;
	F_pfree(m, v188)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L19
	} else {
		goto L59
	}
L47:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_index_constraint_create[0]))
	if v224 == int32(0) {
		goto L46
	} else {
		goto L57
	}
L48:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L19
	} else {
		goto L56
	}
L49:
	;
	if v203 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L50:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+16)))
	if v206 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+16)) = uint8(v209)
	F_CatalogTupleUpdate(m, v183, v188+int32(4), v188)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	if v203 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	F_CatalogTupleUpdate(m, v183, v188+int32(4), v188)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	goto L48
L56:
	;
	goto L47
L57:
	;
	v228 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), l2, v228, v228, l9)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	goto L46
L59:
	;
	F_relation_close(m, v183, int32(3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	goto L39
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_index_constraint_create_9), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(1921), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_index_constraint_create_0), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(1926), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l2
	F_errmsg_internal(m, int32(_a_F_index_constraint_create_7), v15)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(2070), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
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
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
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
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v758 int32
	_ = v758
	var v775 int32
	_ = v775
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v844 int32
	_ = v844
	var v897 int32
	_ = v897
	var v913 int32
	_ = v913
	var v942 int32
	_ = v942
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1043 int32
	_ = v1043
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1087 int64
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1103 int32
	_ = v1103
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int64
	_ = v1165
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1205 int32
	_ = v1205
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1357 int32
	_ = v1357
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1466 int32
	_ = v1466
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1769 int32
	_ = v1769
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1849 int32
	_ = v1849
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2094 int32
	_ = v2094
	var v2130 int32
	_ = v2130
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2215 int32
	_ = v2215
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2314 int32
	_ = v2314
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2351 int32
	_ = v2351
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2411 int32
	_ = v2411
	var v2415 int32
	_ = v2415
	var v2419 int32
	_ = v2419
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2477 int32
	_ = v2477
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	v22 = int32(0)
	v54 = m.G0
	v56 = v54 - int32(288)
	m.G0 = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l6)+92))
	v62 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	switch v67 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L4
	default:
		v73 = v22
		goto L3
	}
L3:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v74 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	v73 = base.B2i32(v70 == int32(0))
	goto L3
L5:
	;
	m.G0 = v56 + int32(288)
	return v2487
L6:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[0]))
	if v2270 != 0 {
		goto L353
	} else {
		goto L354
	}
L7:
	;
	if l3 != 0 {
		goto L325
	} else {
		goto L326
	}
L8:
	;
	F_record_object_address_dependencies(m, v56+int32(192), v1435, int32(97))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L323
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(1259)
	F_add_exact_object_address(m, v56+int32(160), v1435)
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L1
	} else {
		goto L322
	}
L10:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+117)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v66)+68))
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+118)))
	if l18 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L1
	} else {
		goto L319
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L315
	}
L14:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if int32(0) < v98 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	v81 = int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v82) < base.Ui32(int32(_a_F_index_create_0)) {
		v91 = v81
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v91 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+68))
	if v86 == int32(99) {
		v91 = v81
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v89 = F_isTempToastNamespace(m, v86)
	mBase = m.M
	v91 = v89
	goto L17
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v95 == int32(2) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+128)) = v162
	F_errmsg_internal(m, int32(_a_F_index_create_1), v56+int32(128))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L313
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L309
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L1
	} else {
		goto L305
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L1
	} else {
		goto L301
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L298
	}
L27:
	;
	v321 = l16 & int32(2)
	if v321 != 0 {
		goto L74
	} else {
		goto L75
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L69
	}
L29:
	;
	v275 = F_SearchSysCache1(m, int32(14), v162)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L63
	}
L30:
	;
	v124 = int32(0)
	v125 = v98
	goto L33
L31:
	;
	goto L32
L32:
	;
	v231 = l16 & int32(8)
	if v231 != 0 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v156 = v124 << (uint(int32(2)) % 32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l10+v156)))
	if v158 == int32(0) {
		v172 = v125
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v175 = v124 + int32(1)
	if v175 < v172 {
		v124 = v175
		v125 = v172
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l11+v156)))
	if base.Ui32(int32(2)) < base.Ui32(v162-int32(_a_F_index_create_2)) {
		v172 = v125
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v167 = F_get_collation_isdeterministic(m, v158)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v167 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v172 = v171
	goto L35
L40:
	;
	goto L34
L41:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v77&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v232) < base.Ui32(int32(_a_F_index_create_0)) {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v59 != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v238 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v77&int32(1)&base.B2i32(l9 != int32(1664)) != 0 {
		goto L26
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v244 = F_get_relname_relid(m, l1, v78)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v244 == int32(0) {
		goto L27
	} else {
		goto L53
	}
L53:
	;
	if l16&int32(16) == int32(0) {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v252 = int32(0)
	v255 = F_errstart(m, int32(18), v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v255 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_relation_close(m, v62, int32(3))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+112)) = l1
	F_errmsg(m, int32(_a_F_index_create_3), v56+int32(112))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(894), int32(_a_F_index_create_5))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v2487 = v252
	goto L5
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v275 == int32(0) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+144)) = v286 + v287 + int32(8)
	F_errmsg(m, int32(_a_F_index_create_6), v56+int32(144))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(841), int32(_a_F_index_create_5))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+96)) = l1
	F_errmsg(m, int32(_a_F_index_create_7), v56+int32(96))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(902), int32(_a_F_index_create_5))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L294
	}
L74:
	;
	v323 = F_ConstraintNameIsUsed(m, int32(0), v58, l1)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if l7 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v323 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v327 = v326
	goto L81
L80:
	;
	v327 = int32(0)
	goto L81
L81:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v329 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v331 = v330
	goto L84
L83:
	;
	v331 = v22
	goto L84
L84:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v334 = F_GetIndexAmRoutineByAmId(m, l8, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v338 = int32(*(*int16)(unsafe.Add(mBase, uint32(v337)+120)))
	v339 = F_CreateTemplateTupleDesc(m, v328)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if int32(0) < v328 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L1
	} else {
		goto L291
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L288
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L1
	} else {
		goto L285
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L282
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L279
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L1
	} else {
		goto L276
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L273
	}
L94:
	;
	v370 = v22
	v372 = v327
	v375 = v331
	goto L97
L95:
	;
	goto L96
L96:
	;
	v651 = l16 & int32(32)
	F_pfree(m, v334)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L144
	}
L97:
	;
	v398 = int32(1)
	v401 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v370<<(uint(v398)%32)))))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v406 = int32(100)
	v408 = v339 + v402<<(uint(int32(4))%32) + v370*v406
	v410 = v408 + int32(20)
	base.MemoryFill(m, v410, int32(0), v406)
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+112)) = uint8(v398)
	v417 = v370 + v398
	*(*uint16)(unsafe.Add(mBase, uint32(v408)+94)) = uint16(v417)
	if v370 < v332 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l10+v370<<(uint(int32(2))%32))))
	v425 = v423
	goto L101
L100:
	;
	v425 = int32(0)
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+96)) = v425
	if v372 == int32(0) {
		goto L87
	} else {
		goto L102
	}
L102:
	;
	v430 = v408 + int32(24)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v433 = F_strncpy(m, v430, v431, int32(64))
	mBase = m.M
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+63)) = uint8(v434)
	goto L103
L103:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	if v401 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v334)+32))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v370 < v521 {
		goto L119
	} else {
		goto L120
	}
L105:
	;
	if v338 < v401 {
		goto L88
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v375 == int32(0) {
		goto L89
	} else {
		goto L109
	}
L108:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v445 = v336 + v439<<(uint(int32(4))%32) + v401*int32(100)
	v447 = v445 - int32(80)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+68)) = v448
	v450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v447)+72)))
	*(*uint16)(unsafe.Add(mBase, uint32(v410)+72)) = uint16(v450)
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v445))))
	*(*uint16)(unsafe.Add(mBase, uint32(v410)+80)) = uint16(v452)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+76)) = v454
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+82)) = uint8(v456)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+83)) = uint8(v458)
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+84)) = uint8(v460)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+85)) = uint8(v462)
	v512 = v375
	goto L104
L109:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v466)+12))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v471 = F_exprType(m, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v473 = F_SearchSysCache1(m, int32(82), v471)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v473 == int32(0) {
		goto L90
	} else {
		goto L112
	}
L112:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473)+16))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+68)) = v471
	v480 = v477 + v478
	v481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v480)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v410)+72)) = uint16(v481)
	v483 = F_exprTypmod(m, v470)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+76)) = v483
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+82)) = uint8(v486)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+83)) = uint8(v488)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+129)))
	v491 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+85)) = uint8(v491)
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+84)) = uint8(v490)
	F_ReleaseCatCache(m, v473)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v410)+68))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v410)+96))
	v498 = int32(0)
	F_CheckAttributeType(m, v430, v496, v497, v498, v498)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v503 = v375 + int32(4)
	if base.Ui32(v503) < base.Ui32(v468+v467<<(uint(int32(2))%32)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v509 = v503
	goto L118
L117:
	;
	v509 = int32(0)
	goto L118
L118:
	;
	v512 = v509
	goto L104
L119:
	;
	v526 = l11 + v370<<(uint(int32(2))%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v528 = F_SearchSysCache1(m, int32(14), v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	v551 = v520
	goto L121
L121:
	;
	if v551 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	if v528 == int32(0) {
		goto L91
	} else {
		goto L123
	}
L123:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528)+16))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532)+22)))
	v534 = v532 + v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+92))
	if v535 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	F_ReleaseCatCache(m, v528)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L132
	}
L125:
	;
	v536 = v535
	goto L127
L126:
	;
	v536 = v520
	goto L127
L127:
	;
	if v536 != int32(2283) {
		v548 = v536
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534)+84))
	if v540 != int32(2277) {
		v548 = int32(2283)
		goto L124
	} else {
		goto L129
	}
L129:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v410)+68))
	v544 = F_get_base_element_type(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v544 == int32(0) {
		goto L92
	} else {
		goto L131
	}
L131:
	;
	v548 = v544
	goto L124
L132:
	;
	v551 = v548
	goto L121
L133:
	;
	v587 = v372 + int32(4)
	if base.Ui32(v587) < base.Ui32(v437+v436<<(uint(int32(2))%32)) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v410)+68))
	if v551 == v557 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v560 = F_SearchSysCache1(m, int32(82), v551)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v560 == int32(0) {
		goto L93
	} else {
		goto L137
	}
L137:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560)+16))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v410)+68)) = v551
	v569 = v564 + v565
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v569)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v410)+72)) = uint16(v570)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+82)) = uint8(v572)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+83)) = uint8(v574)
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+129)))
	v577 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+85)) = uint8(v577)
	*(*uint8)(unsafe.Add(mBase, uint32(v410)+84)) = uint8(v576)
	F_ReleaseCatCache(m, v560)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L133
L139:
	;
	v593 = v587
	goto L141
L140:
	;
	v593 = int32(0)
	goto L141
L141:
	;
	F_populate_compact_attribute(m, v339, v370)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v417 != v328 {
		v370 = v417
		v372 = v593
		v375 = v512
		goto L97
	} else {
		goto L143
	}
L143:
	;
	goto L98
L144:
	;
	if l2 != 0 {
		v675 = l2
		v676 = l5
		goto L147
	} else {
		goto L148
	}
L145:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L269
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L265
	}
L147:
	;
	v687 = F_heap_create(m, l1, v78, l9, v675, v676, l8, v339, v651^int32(105), v79, v77&int32(1), v73, l18, v56+int32(156), v56+int32(152), base.B2i32(l5 == int32(0)))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L155
	}
L148:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_create[2])))
	if v655 == int32(1) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v659 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[3]))
	if v659 == int32(0) {
		goto L145
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v673 = F_GetNewRelFileNumber(m, l9, v62, v79)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v663 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[3])) = v663
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[4]))
	if v651|v666 == v663 {
		goto L146
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[4])) = int32(0)
	v675 = v659
	v676 = v666
	goto L147
L154:
	;
	v675 = v673
	v676 = l5
	goto L147
L155:
	;
	v689 = m.G0
	v691 = v689 - int32(32)
	m.G0 = v691
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v687)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v691)+16)) = v693
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v687)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v691)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v691)+20)) = v695
	v702 = int32(0)
	v707 = F_LockAcquireExtended(m, v691+int32(16), int32(8), v702, v702, v691+int32(12), v702)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v707 != int32(3) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	m.G0 = v691 + int32(32)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v687)+48))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v719)+80)) = v721
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v687)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v723)+84)) = l8
	v725 = int32(0)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v687)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v726)+131)) = uint8(base.B2i32(l3 != v725))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v687)+56))
	F_InsertPgClassTuple(m, v62, v687, v730, v725, l15)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v714 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v713)+53)) = uint8(v714)
	goto L161
L161:
	;
	goto L159
L162:
	;
	F_relation_close(m, v62, int32(3))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v737 <= int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	if l12 == int32(0) {
		v1103 = v725
		goto L176
	} else {
		goto L177
	}
L165:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v687)+52))
	v742 = v737 & int32(3)
	v743 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v737) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v758 = v743
	v775 = int32(0)
	goto L169
L167:
	;
	v844 = v743
	goto L168
L168:
	;
	v897 = v844
	v913 = v743
	goto L173
L169:
	;
	v804 = v758 * int32(100)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v806 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v804+(v740+v805<<(uint(v806)%32)))+20)) = v675
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v811<<(uint(v806)%32)+v804)+120)) = v675
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v817<<(uint(v806)%32)+v804)+220)) = v675
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v823<<(uint(v806)%32)+v804)+320)) = v675
	v830 = v758 + v806
	v832 = v775 + v806
	if v832 != v737&int32(2147483644) {
		v758 = v830
		v775 = v832
		goto L169
	} else {
		goto L171
	}
L170:
	;
	if v742 == int32(0) {
		goto L164
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v844 = v830
	goto L168
L173:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v942<<(uint(int32(4))%32)+v897*int32(100))+20)) = v675
	v950 = int32(1)
	v953 = v913 + v950
	if v953 != v742 {
		v897 = v897 + v950
		v913 = v953
		goto L173
	} else {
		goto L175
	}
L174:
	;
	goto L164
L175:
	;
	goto L174
L176:
	;
	v1151 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L191
	}
L177:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v687)+52))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1010)))
	v1014 = F_palloc0(m, v1011<<(uint(int32(4))%32))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v687)+52))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)))
	if v1017 <= int32(0) {
		v1103 = v1014
		goto L176
	} else {
		goto L179
	}
L179:
	;
	v1043 = int32(0)
	goto L180
L180:
	;
	v1076 = v1014 + v1043<<(uint(int32(4))%32)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l12+v1043<<(uint(int32(2))%32))))
	if v1080 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v1103 = v1014
	goto L176
L182:
	;
	if l14 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+8)) = v1080
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1082 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+12)) = uint8(v1082)
	goto L182
L186:
	;
	v1092 = v1043 + int32(1)
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v687)+52))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)))
	if v1092 < v1094 {
		v1043 = v1092
		goto L180
	} else {
		goto L190
	}
L187:
	;
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(l14+v1043<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1076))) = v1087
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1089 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+4)) = uint8(v1089)
	goto L186
L190:
	;
	goto L181
L191:
	;
	v1153 = F_CatalogOpenIndexes(m, v1151)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v1155 = int32(0)
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v687)+52))
	F_InsertPgAttributeTuples(m, v1151, v1156, v1155, v1103, v1153)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_CatalogCloseIndexes(m, v1153)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_relation_close(m, v1151, int32(3))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v1165 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+173)) = v1165
	*(*int64)(unsafe.Add(mBase, uint32(v56)+168)) = v1165
	*(*int64)(unsafe.Add(mBase, uint32(v56)+160)) = v1165
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v1173 = F_buildint2vector(m, int32(0), v1172)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v1175 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1205 = int32(0)
	goto L200
L198:
	;
	goto L199
L199:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1300 = F_buildoidvector(m, l10, v1299)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L203
	}
L200:
	;
	v1236 = int32(1)
	v1237 = v1205 << (uint(v1236) % 32)
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6+int32(12)+v1237))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1173+int32(24)+v1237))) = uint16(v1240)
	v1243 = v1205 + v1236
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1243 < v1244 {
		v1205 = v1243
		goto L200
	} else {
		goto L202
	}
L201:
	;
	goto L199
L202:
	;
	goto L201
L203:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1303 = F_buildoidvector(m, l11, v1302)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1306 = F_buildint2vector(m, l13, v1305)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v1308 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1309 = F_nodeToString(m, v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	v1316 = v1155
	goto L208
L208:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v1318 != 0 {
		goto L212
	} else {
		goto L213
	}
L209:
	;
	v1311 = F_cstring_to_text(m, v1309)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_pfree(m, v1309)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v1316 = v1311
	goto L208
L212:
	;
	v1319 = F_make_ands_explicit(m, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L215
	}
L213:
	;
	v1328 = int32(0)
	goto L214
L214:
	;
	v1330 = l16 & int32(1)
	v1337 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L219
	}
L215:
	;
	v1321 = F_nodeToString(m, v1319)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v1323 = F_cstring_to_text(m, v1321)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_pfree(m, v1321)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v1328 = v1323
	goto L214
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+196)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v56)+192)) = v675
	v1341 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+200)) = v1341
	v1343 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+204)) = v1343
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+208)) = v1345
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+268)) = v1316
	*(*int32)(unsafe.Add(mBase, uint32(v56)+264)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v56)+260)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v56)+256)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v56)+252)) = v1173
	*(*int64)(unsafe.Add(mBase, uint32(v56)+244)) = int64(1)
	v1357 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+240)) = base.B2i32(int32(base.Ui32(v231)>>(uint(int32(3))%32)) == v1357)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+236)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v56)+232)) = base.B2i32(l16&int32(72) == v1357)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+228)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v56)+224)) = base.B2i32(l17&int32(2) == v1357)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+220)) = base.B2i32(v59 != v1357)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+216)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v56)+212)) = v1347
	if v1316 == v1357 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1377 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+179)) = uint8(v1377)
	goto L222
L221:
	;
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+272)) = v1328
	if v1328 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1382 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+180)) = uint8(v1382)
	goto L225
L224:
	;
	goto L225
L225:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1337)+52))
	v1389 = F_heap_form_tuple(m, v1384, v56+int32(192), v56+int32(160))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_CatalogTupleInsert(m, v1337, v1389)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_relation_close(m, v1337, int32(3))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_pfree(m, v1389)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if l3 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	F_StoreSingleInheritance(m, v675, l3, int32(1))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v1410 == int32(0) {
		goto L6
	} else {
		goto L237
	}
L234:
	;
	F_LockRelationOid(m, l3, int32(4))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_SetRelationHasSubclass(m, l3, int32(1))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	goto L233
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+196)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v56)+192)) = int32(1259)
	if v321 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L262
	}
L239:
	;
	if v1330 != 0 {
		v1426 = int32(112)
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v1435 = F_new_object_addresses(m)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L251
	}
L242:
	;
	F_index_constraint_create(m, v56+int32(160), l0, v675, l4, l6, l1, v1426, l17, l18, l19)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L1
	} else {
		goto L249
	}
L243:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	if v1421 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1422 = int32(117)
	goto L246
L245:
	;
	v1422 = int32(120)
	goto L246
L246:
	;
	if v1421 != 0 {
		v1426 = v1422
		goto L242
	} else {
		goto L247
	}
L247:
	;
	if v59 == int32(0) {
		goto L238
	} else {
		goto L248
	}
L248:
	;
	v1426 = v1422
	goto L242
L249:
	;
	if l20 == int32(0) {
		goto L7
	} else {
		goto L250
	}
L250:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v56)+164))
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v1433
	goto L7
L251:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1437 <= int32(0) {
		goto L9
	} else {
		goto L252
	}
L252:
	;
	v1442 = int32(0)
	v1448 = v1442
	v1452 = v1437
	v1466 = v1442
	goto L253
L253:
	;
	v1500 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v1466<<(uint(int32(1))%32)))))
	if v1500 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	goto L8
L255:
	;
	v1504 = v1466 + int32(1)
	if v1504 < v1452 {
		v1466 = v1504
		goto L253
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(1259)
	F_add_exact_object_address(m, v56+int32(160), v1435)
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	if v1448 != 0 {
		goto L8
	} else {
		goto L259
	}
L259:
	;
	goto L9
L260:
	;
	v1514 = int32(1)
	v1516 = v1466 + v1514
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1516 < v1517 {
		v1448 = v1514
		v1452 = v1517
		v1466 = v1516
		goto L253
	} else {
		goto L261
	}
L261:
	;
	goto L254
L262:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_8), int32(0))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(1102), int32(_a_F_index_create_5))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_index_create_9), int32(0))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(953), int32(_a_F_index_create_5))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(_a_F_index_create_10), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(943), int32(_a_F_index_create_5))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+48)) = v551
	F_errmsg_internal(m, int32(_a_F_index_create_11), v56+int32(48))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(466), int32(_a_F_index_create_12))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v410)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+32)) = v1583
	F_errmsg_internal(m, int32(_a_F_index_create_13), v56+int32(32))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(452), int32(_a_F_index_create_12))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = v1599
	F_errmsg_internal(m, int32(_a_F_index_create_14), v56+int32(16))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(433), int32(_a_F_index_create_12))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v471
	F_errmsg_internal(m, int32(_a_F_index_create_11), v56)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(377), int32(_a_F_index_create_12))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_15), int32(0))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(367), int32(_a_F_index_create_12))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+64)) = v401
	F_errmsg_internal(m, int32(_a_F_index_create_16), v56-int32(-64))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(348), int32(_a_F_index_create_12))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_17), int32(0))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(331), int32(_a_F_index_create_12))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	F_errcode(m, int32(_a_F_index_create_18))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v56)+84)) = v1672 + int32(4)
	F_errmsg(m, int32(_a_F_index_create_19), v56+int32(80))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(916), int32(_a_F_index_create_5))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_20), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(879), int32(_a_F_index_create_5))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_errmsg(m, int32(_a_F_index_create_21), int32(0))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(873), int32(_a_F_index_create_5))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L305:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errmsg(m, int32(_a_F_index_create_22), int32(0))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(864), int32(_a_F_index_create_5))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_errmsg(m, int32(_a_F_index_create_23), int32(0))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(855), int32(_a_F_index_create_5))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(837), int32(_a_F_index_create_5))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	F_errmsg(m, int32(_a_F_index_create_24), int32(0))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(800), int32(_a_F_index_create_5))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_25), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(793), int32(_a_F_index_create_5))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	goto L8
L323:
	;
	F_free_object_addresses(m, v1435)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	goto L7
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(1259)
	v1969 = v56 + int32(192)
	v1971 = v56 + int32(160)
	F_recordDependencyOn(m, v1969, v1971, int32(80))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1985 = F_new_object_addresses(m)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L330
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(1259)
	F_recordDependencyOn(m, v1969, v1971, int32(83))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	goto L327
L330:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v1987 <= int32(0) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v2196 = v56 + int32(192)
	F_record_object_address_dependencies(m, v2196, v1985, int32(110))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L1
	} else {
		goto L345
	}
L332:
	;
	v2013 = int32(0)
	v2014 = v1987
	goto L333
L333:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(l10+v2013<<(uint(int32(2))%32))))
	v2048 = int32(0)
	if base.B2i32(v2047 == v2048)|base.B2i32(v2047 == int32(100)) == v2048 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	if v2065 <= int32(0) {
		goto L331
	} else {
		goto L340
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = v2047
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(3456)
	F_add_exact_object_address(m, v56+int32(160), v1985)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L1
	} else {
		goto L338
	}
L336:
	;
	v2065 = v2014
	goto L337
L337:
	;
	v2067 = v2013 + int32(1)
	if v2067 < v2065 {
		v2013 = v2067
		v2014 = v2065
		goto L333
	} else {
		goto L339
	}
L338:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v2065 = v2064
	goto L337
L339:
	;
	goto L334
L340:
	;
	v2094 = int32(0)
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+160)) = int32(2616)
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(l11+v2094<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+164)) = v2130
	F_add_exact_object_address(m, v56+int32(160), v1985)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L1
	} else {
		goto L343
	}
L342:
	;
	goto L331
L343:
	;
	v2139 = v2094 + int32(1)
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2139 < v2140 {
		v2094 = v2139
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	F_free_object_addresses(m, v1985)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v2202 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	F_recordDependencyOnSingleRelExpr(m, v2196, v2202, v58, int32(97), int32(0))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v2207 == int32(0) {
		goto L6
	} else {
		goto L351
	}
L350:
	;
	goto L349
L351:
	;
	F_recordDependencyOnSingleRelExpr(m, v56+int32(192), v2207, v58, int32(97), int32(0))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	goto L6
L353:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v675, int32(0), l19)
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L1
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L1
	} else {
		goto L357
	}
L356:
	;
	goto L355
L357:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v2278 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	F_RelationInitIndexAccessInfo(m, v687)
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v687)+192))
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v2283)+10)) = uint16(v2284)
	if l12 == int32(0) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L360
L362:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v2411 == int32(0) {
		goto L370
	} else {
		goto L371
	}
L363:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2288 <= int32(0) {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v2314 = int32(0)
	goto L365
L365:
	;
	v2345 = int32(1)
	v2346 = v2314 + v2345
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(l12+v2314<<(uint(int32(2))%32))))
	v2353 = F_index_opclass_options(m, v687, base.I32_extend16_s(v2346), v2351, v2345)
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L1
	} else {
		goto L367
	}
L366:
	;
	goto L362
L367:
	;
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2346 < v2355 {
		v2314 = v2346
		goto L365
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	F_relation_close(m, v687, int32(0))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L1
	} else {
		goto L387
	}
L370:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[5]))
	if v2415 == int32(0) {
		goto L373
	} else {
		goto L374
	}
L371:
	;
	goto L372
L372:
	;
	if l16&int32(4) != 0 {
		goto L381
	} else {
		goto L382
	}
L373:
	;
	v2419 = int32(0)
	v2424 = F_AllocSetContextCreateInternal(m, v2419, int32(_a_F_index_create_26), v2419, int32(_a_F_index_create_27), int32(_a_F_index_create_28))
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L1
	} else {
		goto L376
	}
L374:
	;
	v2427 = v2415
	goto L375
L375:
	;
	v2428 = int32(_a_F_index_create_29)
	v2429 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[6])) = v2427
	v2433 = F_palloc(m, int32(16))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L1
	} else {
		goto L377
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[5])) = v2424
	v2427 = v2424
	goto L375
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2433)+4)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v2433))) = v58
	v2438 = F_palloc(m, int32(144))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2433)+8)) = v2438
	base.MemoryCopy(m, v2438, l6, int32(144))
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v2444 = F_copyObjectImpl(m, v2443)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2446)+76)) = v2444
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2448)+80)) = int32(0)
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	v2452 = F_copyObjectImpl(m, v2451)
	mBase = m.M
	v2453 = m.ExcPending
	if v2453 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2454)+84)) = v2452
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2433)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2456)+88)) = int32(0)
	v2459 = int32(_a_F_index_create_30)
	v2460 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v2433)+12)) = v2460
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[6])) = v2429
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[7])) = v2433
	goto L369
L381:
	;
	F_index_update_stats(m, l0, int32(1), float64(-1))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	F_index_build(m, l0, v687, l6, int32(0), int32(1))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L1
	} else {
		goto L386
	}
L384:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	goto L369
L386:
	;
	goto L369
L387:
	;
	v2487 = v675
	goto L5
}
func F_index_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if int32(0) <= base.I32_extend16_s(v7) {
		v11 = int32(8)
	} else {
		v11 = int32(16)
	}
	F_index_deform_tuple_internal(m, l1, l2, l3, l0+v11, l0+int32(8), int32(base.Ui32(v7)>>(uint(int32(15))%32)))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		return
	}
}
func F_index_getnext_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	goto L1
L1:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(1)
L3:
	;
	v14 = F_index_fetch_heap(m, l0, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L8
	}
L4:
	;
	v8 = F_index_getnext_tid(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	return int32(0)
L8:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
}
func F_index_insert_cleanup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_index_insert_cleanup[0]))
	if v11 != v9 {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_index_insert_cleanup[1]))
		v15 = F_list_member_ptr(m, v14, v9)
		mBase = m.M
		v17 = v15
	} else {
		v17 = int32(1)
	}
	if v17 == int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
		if v21 != 0 {
			m.T0[v21].(func(*base.Module, int32, int32))(m, l0, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34 + int32(4)
				F_errmsg(m, int32(_a_F_index_insert_cleanup_0), v7)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_index_insert_cleanup_1), int32(244), int32(_a_F_index_insert_cleanup_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
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
func F_validate_index_callback(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v8 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v13 = F_Int64GetDatum(m, v4|(v5<<(uint(int64(16))%64)|v8<<(uint(int64(32))%64)))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_tuplesort_putdatum(m, v3, v13, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v20, float64(1))
			return int32(0)
		}
	}
}
