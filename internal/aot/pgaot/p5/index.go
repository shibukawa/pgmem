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
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
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
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(124)))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, _consts[240]))
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
	v731 = m.ExcPending
	if v731 != 0 {
		goto L6
	} else {
		goto L200
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L6
	} else {
		goto L193
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l1+v119))) = v282
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
	F_errmsg(m, int32(68207), v32+int32(80))
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
	F_errmsg(m, int32(71384), v32+int32(96))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(489273), int32(1960), int32(129842))
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
	v282 = v147
	v288 = v149
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
		v282 = v147
		v288 = v149
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
	F_errmsg(m, int32(535246), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(489273), int32(2019), int32(129842))
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
	F_errmsg(m, int32(147340), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(489273), int32(1978), int32(129842))
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
	F_errfinish(m, int32(489273), int32(1955), int32(129842))
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
	v282 = v259
	v288 = v258
	goto L28
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v638+v119))) = v633
	v642 = v111 + int32(1)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v642 < v643 {
		v111 = v642
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
		v345 = v288
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
	v638 = l2
	goto L68
L84:
	;
	v346 = F_type_is_collatable(m, v282)
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
	*(*int32)(unsafe.Add(mBase, _consts[240])) = l15
	*(*int32)(unsafe.Add(mBase, _consts[239])) = l14
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
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v330
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v329
	goto L93
L93:
	;
	v336 = int32(4478344)
	v338 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	v340 = v338 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[241])) = v340
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
	F_errmsg(m, int32(266776), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_errhint(m, int32(554374), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(489273), int32(2090), int32(129842))
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
	v380 = F_ResolveOpClass(m, v379, v282, l9, l10)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[240])) = l15
	*(*int32)(unsafe.Add(mBase, _consts[239])) = l14
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
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v391
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v390
	goto L120
L118:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v388 = F_compatible_oper_opid(m, v387, v282, v282)
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
	v397 = int32(4478344)
	v399 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	v401 = v399 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[241])) = v401
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
	F_errmsg(m, int32(260398), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(489273), int32(2034), int32(129842))
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
	F_errmsg(m, int32(129380), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(489273), int32(2038), int32(129842))
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
	F_errmsg(m, int32(136892), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(489273), int32(2042), int32(129842))
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
	F_errmsg(m, int32(136706), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(489273), int32(2046), int32(129842))
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
	v507 = F_format_type_be(m, v282)
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
	F_errmsg(m, int32(186815), v32-int32(-64))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(489273), int32(2098), int32(129842))
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
	*(*int32)(unsafe.Add(mBase, _consts[240])) = l15
	*(*int32)(unsafe.Add(mBase, _consts[239])) = l14
	goto L152
L152:
	;
	v529 = F_compatible_oper_opid(m, v520, v282, v282)
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
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v532
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v531
	goto L154
L154:
	;
	v538 = int32(4478344)
	v540 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	v542 = v540 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[241])) = v542
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
	F_errmsg(m, int32(136765), v32)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(489273), int32(2234), int32(129842))
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
	v638 = l4
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
	v683 = m.ExcPending
	if v683 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	v684 = F_format_operator(m, v547)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v684
	F_errmsg(m, int32(339673), v32+int32(48))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	F_errdetail(m, int32(564909), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(489273), int32(2166), int32(129842))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
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
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	v708 = F_format_operator(m, v547)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	v710 = F_get_opfamily_name(m, v552)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v708
	F_errmsg(m, int32(672027), v32+int32(32))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	F_errdetail(m, int32(559497), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(489273), int32(2179), int32(129842))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
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
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = l9
	F_errmsg(m, int32(136943), v32+int32(16))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(489273), int32(2229), int32(129842))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(176)
	m.G0 = v19
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v29 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v36 = v29
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = l1
	if int32(0) < v28 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	v36 = v32
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L61
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L56
	}
L8:
	;
	m.G0 = v19 + int32(176)
	return v201
L9:
	;
	v41 = int32(0)
	v47 = v7
	goto L12
L10:
	;
	v175 = v7
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
	v58 = v41 << (uint(int32(2)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v27+v58)))
	if v60 == int32(0) {
		v163 = v47
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v175 = v163
	goto L11
L14:
	;
	v167 = v41 + int32(1)
	if v167 != v28 {
		v41 = v167
		v47 = v163
		goto L12
	} else {
		goto L51
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+v26)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+116)))
	if v65 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+92))
	if v68 == int32(0) {
		v163 = v47
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+118)))
	if v71 != int32(1) {
		v163 = v47
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+192))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v76 = int32(0)
	if l5 == v76 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v60)+192))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+16)))
	if v118 == int32(0) {
		goto L7
	} else {
		goto L38
	}
L24:
	;
	if v114 == int32(0) {
		v163 = v47
		goto L14
	} else {
		goto L37
	}
L25:
	;
	v114 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v82 <= int32(0) {
		v107 = v76
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = v107
	goto L24
L29:
	;
	v85 = int32(0)
	if v85 < v82 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v88 = v82
	goto L32
L31:
	;
	v88 = v85
	goto L32
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v91 = int32(0)
	goto L33
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89+v91<<(uint(int32(2))%32))))
	v100 = base.B2i32(v99 == v75)
	if v99 == v75 {
		v107 = v100
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v107 = v100
	goto L28
L35:
	;
	v102 = v91 + int32(1)
	if v102 != v88 {
		v91 = v102
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v64)+84))
	if v121 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_FormIndexDatum(m, v64, l1, l2, v19+int32(32), v19)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L48
	}
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v64)+88))
	if v124 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v127 = F_ExecPrepareQual(m, v121, l2)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v132 = v124
	goto L43
L43:
	;
	v133 = int32(4480304)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v136
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)+20))
	v141 = m.T0[v140].(func(*base.Module, int32, int32, int32) int32)(m, v132, v36, v19+int32(175))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+88)) = v127
	if v127 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v132 = v127
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v134
	if v141 != 0 {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v163 = int32(1)
	goto L14
L48:
	;
	v152 = int32(0)
	v153 = int32(1)
	v159 = F_check_exclusion_or_unique_constraint(m, v25, v60, v64, l4, v19+int32(32), v19, l2, v152, v152, v153, l3)
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
		v201 = v152
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v163 = v153
	goto L14
L51:
	;
	goto L13
L52:
	;
	v201 = int32(1)
	goto L8
L53:
	;
	goto L54
L54:
	;
	v188 = int32(1)
	if v175&v188 == int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v201 = v188
	goto L8
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(130925), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	F_errtableconstraint(m, v25, v224+int32(4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(493158), int32(610), int32(119134))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	F_errmsg_internal(m, int32(28192), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(493158), int32(656), int32(119134))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
	v4 = F_GetPageWithFreeSpace(m, l0, int32(4096))
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
		F_errmsg_internal(m, int32(148746), int32(0))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(491749), int32(328), int32(315076))
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int64
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v514 int64
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v528 int64
	_ = v528
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v542 int64
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v580 int64
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
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
	var v627 int64
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int64
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v727 float64
	_ = v727
	var v729 int32
	_ = v729
	var v731 float64
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
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
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	v341 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L8
	} else {
		goto L87
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[231]))
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
	v31 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(48)
	m.G0 = v34
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v37 != int32(1) {
		v310 = v31
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v34 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+128)) = v310
	goto L1
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[236]))
	if v41 == int32(0) {
		v310 = v31
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = F_palloc0(m, int32(168))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(4294967363)
	v50 = F_palloc0(m, int32(92))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(266)
	v55 = F_palloc0(m, int32(384))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(267)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+344)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+280)) = v64
	v69 = F_palloc0(m, int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v69
	v78 = F_list_make1_impl(m, int32(1), v34+int32(12))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+84)) = v78
	v82 = F_palloc0(m, int32(136))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v84 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(101)
	v91 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+124)) = uint16(v91)
	v93 = int32(29185)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+20)) = uint16(v93)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v82
	v100 = F_list_make1_impl(m, v84, v34+int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+52)) = v100
	v105 = F_addRTEPermissionInfo(m, v45+int32(56), v82)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	F_setup_simple_rel_arrays(m, v55)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v111 = F_build_simple_rel(m, v55, int32(1), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v114 = F_table_open(m, v29, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v117 = F_index_open(m, v30, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v119 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+48))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+118)))
	if v121 == int32(116) {
		v290 = v119
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_relation_close(m, v117, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L8
	} else {
		goto L85
	}
L22:
	;
	v124 = F_RelationGetIndexExpressions(m, v117)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v126 = F_is_parallel_safe(m, v55, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	if v126 == int32(0) {
		v290 = v119
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v130 = F_RelationGetIndexPredicate(m, v117)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v132 = F_is_parallel_safe(m, v55, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v132 == int32(0) {
		v290 = v119
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+148))
	if v136 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[236]))
	if v136 < v140 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	F_estimate_rel_size(m, v114, int32(0), v34+int32(44), v34+int32(32), v34+int32(24))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L35
	}
L32:
	;
	v142 = v136
	goto L34
L33:
	;
	v142 = v140
	goto L34
L34:
	;
	v290 = v142
	goto L21
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v153 = base.F64_convert_i32_u(v152)
	v154 = float64(-1)
	v156 = *(*int32)(unsafe.Add(mBase, _consts[236]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v111)+148))
	if v159 != int32(-1) {
		v248 = v159
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v258 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L37:
	;
	goto L36
L38:
	;
	if v248 < v156 {
		goto L73
	} else {
		goto L74
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v162 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v179 = int32(0)
	if base.F64_ge(v153, float64(0)) == v179 {
		v211 = v179
		goto L48
	} else {
		goto L49
	}
L41:
	;
	if base.F64_ge(v153, float64(0)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if base.F64_lt(v153, base.F64_convert_i32_s(v167)) != 0 {
		v258 = int32(0)
		goto L37
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.F64_ge(v154, float64(0)) == int32(0) {
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[238]))
	if base.F64_lt(v154, base.F64_convert_i32_s(v176)) != 0 {
		v258 = int32(0)
		goto L37
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	if base.F64_ge(v154, float64(0)) == int32(0) {
		v248 = v211
		goto L38
	} else {
		goto L57
	}
L49:
	;
	v184 = int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	if v186 <= v184 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v189 = v184
	goto L52
L51:
	;
	v189 = v186
	goto L52
L52:
	;
	v191 = v189
	v195 = int32(1)
	goto L53
L53:
	;
	v198 = v191 * int32(3)
	if base.F64_ge(v153, base.F64_convert_i32_u(v198)) == int32(0) {
		v211 = v195
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v211 = v204
	goto L48
L55:
	;
	v204 = v195 + int32(1)
	if v198 < int32(715827883) {
		v191 = v198
		v195 = v204
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v217 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, _consts[238]))
	if v219 <= v217 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v222 = v217
	goto L60
L59:
	;
	v222 = v219
	goto L60
L60:
	;
	v224 = v222
	v229 = int32(1)
	goto L61
L61:
	;
	v231 = v224 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v231), v154) != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v211 < v238 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v235 = v229 + int32(1)
	if v231 < int32(715827883) {
		v224 = v231
		v229 = v235
		goto L61
	} else {
		goto L66
	}
L64:
	;
	v238 = v229
	goto L65
L65:
	;
	goto L62
L66:
	;
	v238 = v235
	goto L65
L67:
	;
	v240 = v211
	goto L69
L68:
	;
	v240 = v238
	goto L69
L69:
	;
	if int32(0) < v211 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v243 = v240
	goto L72
L71:
	;
	v243 = v238
	goto L72
L72:
	;
	v248 = v243
	goto L38
L73:
	;
	v251 = v248
	goto L75
L74:
	;
	v251 = v156
	goto L75
L75:
	;
	v258 = v251
	goto L37
L76:
	;
	v290 = v258
	goto L21
L77:
	;
	goto L78
L78:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[18]))
	v268 = v258
	goto L79
L79:
	;
	v279 = base.I32_div_s(v262, v268+int32(1))
	if int32(32767) < v279 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v290 = v119
	goto L21
L81:
	;
	v290 = v268
	goto L21
L82:
	;
	goto L83
L83:
	;
	v282 = int32(1)
	if v282 < v268 {
		v268 = v268 - v282
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	F_sequence_close(m, v114, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v310 = v290
	goto L5
L87:
	;
	if v338 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(92)))) = v393
	v396 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(88)))) = v396
	goto L98
L89:
	;
	F_errfinish(m, int32(487518), v382, int32(426253))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L97
	}
L90:
	;
	if v341 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v341 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L93:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v350 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v349 + v350
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v348 + v350
	F_errmsg_internal(m, int32(19693), v17+int32(48))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v382 = int32(3037)
	goto L89
L95:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v366
	v368 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v365 + v368
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v364 + v368
	F_errmsg_internal(m, int32(132796), v17-int32(-64))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	v382 = int32(3043)
	goto L89
L97:
	;
	goto L88
L98:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+80))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v400 | int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v399
	goto L99
L99:
	;
	v408 = int32(4478344)
	v410 = *(*int32)(unsafe.Add(mBase, _consts[241]))
	v412 = v410 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[241])) = v412
	goto L100
L100:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v417 = *(*int64)(unsafe.Add(mBase, _consts[242]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v417
	v420 = *(*int64)(unsafe.Add(mBase, _consts[243]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v420
	v423 = *(*int64)(unsafe.Add(mBase, _consts[244]))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v423
	v425 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+168)) = v425
	*(*int64)(unsafe.Add(mBase, uint32(v17)+160)) = v425
	*(*int64)(unsafe.Add(mBase, uint32(v17)+152)) = v425
	*(*int64)(unsafe.Add(mBase, uint32(v17)+144)) = v425
	*(*int64)(unsafe.Add(mBase, uint32(v17)+136)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = int64(2)
	v439 = v17 + int32(96)
	v441 = v17 + int32(128)
	v442 = int32(0)
	v449 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v449 == v442 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+36))
	v617 = m.T0[v616].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L119
	}
L103:
	;
	goto L102
L104:
	;
	goto L105
L105:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v455&int32(1) == int32(0) {
		goto L103
	} else {
		goto L106
	}
L106:
	;
	v460 = int32(4474964)
	v462 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v463 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v462 + v463
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v466 + v463
	goto L108
L107:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v597 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = v596 + v597
	v600 = int32(4474964)
	v602 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v602 - v597
	goto L103
L108:
	;
	v475 = v449 + int32(232)
	goto L109
L109:
	;
	v481 = int32(0)
	v484 = v442
	goto L112
L111:
	;
	goto L115
L112:
	;
	v490 = int32(2)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v439+v484<<(uint(v490)%32))))
	v494 = int32(3)
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v441+v484<<(uint(v494)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v493<<(uint(v494)%32)))) = v500
	v503 = v484 | int32(1)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v439+v503<<(uint(v490)%32))))
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v441+v503<<(uint(v494)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v507<<(uint(v494)%32)))) = v514
	v517 = v484 | v490
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v439+v517<<(uint(v490)%32))))
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v441+v517<<(uint(v494)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v521<<(uint(v494)%32)))) = v528
	v531 = v484 | v494
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v439+v531<<(uint(v490)%32))))
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v441+v531<<(uint(v494)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v535<<(uint(v494)%32)))) = v542
	v544 = int32(4)
	v545 = v484 + v544
	v547 = v481 + v544
	if v547 != int32(4) {
		v481 = v547
		v484 = v545
		goto L112
	} else {
		goto L114
	}
L113:
	;
	goto L111
L114:
	;
	goto L113
L115:
	;
	v561 = int32(0)
	v564 = v545
	goto L116
L116:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v439+v564<<(uint(int32(2))%32))))
	v574 = int32(3)
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v441+v564<<(uint(v574)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v475+v573<<(uint(v574)%32)))) = v580
	v582 = int32(1)
	v585 = v561 + v582
	if v585 != int32(2) {
		v561 = v585
		v564 = v564 + v582
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L107
L118:
	;
	goto L117
L119:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+118)))
	if v620 != int32(117) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if l3 != 0 {
		goto L144
	} else {
		goto L145
	}
L121:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v623 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v649 = v623
	goto L124
L123:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v625
	v627 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = v627
	v631 = F_smgropen(m, v17+int32(32), v624)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L8
	} else {
		goto L125
	}
L124:
	;
	v651 = F_smgrexists(m, v649, int32(3))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L8
	} else {
		goto L130
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v631
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631)+72))
	if v635 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v649 = v647
	goto L124
L127:
	;
	v643 = v635
	goto L129
L128:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v631)+76))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v631)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v636)+4)) = v637
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v631)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v637))) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v631)+72))
	v643 = v641
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v631)+72)) = v643 + int32(1)
	goto L126
L130:
	;
	if v651 != 0 {
		goto L120
	} else {
		goto L131
	}
L131:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v653 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v679 = v653
	goto L134
L133:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v655
	v657 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v657
	v661 = F_smgropen(m, v17+int32(16), v654)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L8
	} else {
		goto L135
	}
L134:
	;
	F_smgrcreate(m, v679, int32(3), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L8
	} else {
		goto L140
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v661
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661)+72))
	if v665 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v679 = v677
	goto L134
L137:
	;
	v673 = v665
	goto L139
L138:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v661)+76))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v661)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v666)+4)) = v667
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v661)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v667))) = v669
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v661)+72))
	v673 = v671
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661)+72)) = v673 + int32(1)
	goto L136
L140:
	;
	F_log_smgrcreate(m, l1, int32(3))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L8
	} else {
		goto L141
	}
L141:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)+40))
	m.T0[v688].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L8
	} else {
		goto L142
	}
L142:
	;
	goto L120
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L8
	} else {
		goto L207
	}
L144:
	;
	v727 = *(*float64)(unsafe.Add(mBase, uint32(v617)))
	F_index_update_stats(m, l0, int32(1), v727)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L8
	} else {
		goto L154
	}
L145:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)))
	if v692&int32(1) == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v697 != 0 {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v701 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L8
	} else {
		goto L148
	}
L148:
	;
	v705 = F_SearchSysCacheCopy(m, int32(34), v698, int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L8
	} else {
		goto L149
	}
L149:
	;
	if v705 == int32(0) {
		goto L143
	} else {
		goto L150
	}
L150:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v705)+16))
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709)+22)))
	v712 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v709+v710)+19)) = uint8(v712)
	F_CatalogTupleUpdate(m, v701, v705+int32(4), v705)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L8
	} else {
		goto L151
	}
L151:
	;
	F_pfree(m, v705)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	F_sequence_close(m, v701, int32(3))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	goto L144
L154:
	;
	v731 = *(*float64)(unsafe.Add(mBase, uint32(v617)+8))
	F_index_update_stats(m, l1, int32(0), v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v736 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v739 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	if v737 == v739 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	goto L159
L159:
	;
	F_AtEOXact_GUC(m, int32(0), v412)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L8
	} else {
		goto L205
	}
L160:
	;
	v742 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[54])) = v742
	*(*int32)(unsafe.Add(mBase, _consts[245])) = v742
	goto L162
L161:
	;
	goto L162
L162:
	;
	v747 = F_CreateExecutorState(m)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v747)+152))
	if v749 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v752 = F_MakePerTupleExprContext(m, v747)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L8
	} else {
		goto L167
	}
L165:
	;
	v754 = v749
	goto L166
L166:
	;
	v756 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L8
	} else {
		goto L168
	}
L167:
	;
	v754 = v752
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v754)+4)) = v756
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v760 = F_ExecPrepareQual(m, v759, v747)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L8
	} else {
		goto L169
	}
L169:
	;
	v762 = F_GetLatestSnapshot(m)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L8
	} else {
		goto L170
	}
L170:
	;
	v764 = F_RegisterSnapshot(m, v762)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L8
	} else {
		goto L171
	}
L171:
	;
	v766 = int32(0)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+8))
	v772 = m.T0[v771].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v764, v766, v766, v766, int32(449))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L8
	} else {
		goto L172
	}
L172:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+36)) = v775
	v778 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v778 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)+188))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)+12))
	m.T0[v887].(func(*base.Module, int32))(m, v772)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L8
	} else {
		goto L201
	}
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L8
	} else {
		goto L198
	}
L175:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v780&int32(1) == int32(0) {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	goto L179
L178:
	;
	goto L177
L179:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)+188))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	v805 = m.T0[v804].(func(*base.Module, int32, int32, int32) int32)(m, v772, int32(1), v756)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L8
	} else {
		goto L181
	}
L180:
	;
	goto L174
L181:
	;
	if v805 == int32(0) {
		goto L173
	} else {
		goto L182
	}
L182:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v810 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L8
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	if v760 != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L185
L187:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v772)))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v756)+36)) = v848
	v851 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	if v851 == int32(0) {
		goto L179
	} else {
		goto L196
	}
L188:
	;
	v813 = int32(4480304)
	v814 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v754)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v816
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v760)+20))
	v821 = m.T0[v820].(func(*base.Module, int32, int32, int32) int32)(m, v760, v754, v17+int32(271))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L8
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	F_FormIndexDatum(m, l2, v756, v747, v17+int32(128), v17+int32(96))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L8
	} else {
		goto L193
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v814
	if v821 == int32(0) {
		goto L187
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	F_check_exclusion_constraint(m, l0, l1, l2, v756+int32(28), v17+int32(128), v17+int32(96), v747, int32(1))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L8
	} else {
		goto L194
	}
L194:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v754)+20))
	F_MemoryContextReset(m, v842)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L8
	} else {
		goto L195
	}
L195:
	;
	goto L187
L196:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, _consts[247])))
	if v855&int32(1) != 0 {
		goto L179
	} else {
		goto L197
	}
L197:
	;
	goto L180
L198:
	;
	F_errmsg_internal(m, int32(332904), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L8
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(323556), int32(1034), int32(84060))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L8
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_UnregisterSnapshot(m, v764)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L8
	} else {
		goto L202
	}
L202:
	;
	F_ExecDropSingleTupleTableSlot(m, v756)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L8
	} else {
		goto L203
	}
L203:
	;
	F_FreeExecutorState(m, v747)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L8
	} else {
		goto L204
	}
L204:
	;
	v896 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v896
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v896
	goto L159
L205:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, _consts[240])) = v918
	*(*int32)(unsafe.Add(mBase, _consts[239])) = v917
	goto L206
L206:
	;
	m.G0 = v17 + int32(272)
	return
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v698
	F_errmsg_internal(m, int32(39884), v17)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L8
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(487518), int32(3139), int32(426253))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L8
	} else {
		goto L209
	}
L209:
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	if v17 != v13 {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[55]))
		v21 = F_list_member_ptr(m, v20, v13)
		mBase = m.M
		v22 = v21
	} else {
		v22 = int32(1)
	}
	if v22 == int32(0) {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+204))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
		if v26 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(347654)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v60 + int32(4)
				F_errmsg_internal(m, int32(672497), v10+int32(16))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492203), int32(803), int32(347722))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = m.T0[v26].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(32)
				return v29
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v44 + int32(4)
				F_errmsg(m, int32(434269), v10)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492203), int32(802), int32(347722))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
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
	v274 = m.ExcPending
	if v274 != 0 {
		goto L19
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L19
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
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
	if base.Ui32(v21) < base.Ui32(int32(12000)) {
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
	v34 = *(*int32)(unsafe.Add(mBase, _consts[231]))
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
	v41 = l7 & int32(2)
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
	v44 = int32(1259)
	v47 = F_deleteDependencyRecordsForClass(m, v44, l2, v44, int32(97))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v49 = int32(0)
	v52 = l7 & int32(4)
	v55 = int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v69 = int32(32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+92))
	v78 = base.B2i32(l3 == v49)
	v85 = F_CreateConstraintEntry(m, l5, v18, l6, base.B2i32(v41 != v49), base.B2i32(v52 != v49), v55, v55, l3, v57, l4+int32(12), v60, v61, v49, l2, v49, v49, v49, v49, v49, v49, v69, v69, v49, v49, v69, v74, v49, v49, v78, base.B2i32(l3 != v49), v78, base.B2i32(l7&v69 != v49), l9)
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
	F_recordDependencyOn(m, l0, v15+int32(8), int32(80))
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
	if v41 != 0 {
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
	F_recordDependencyOn(m, l0, v15+int32(8), int32(83))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v125 = F_palloc0(m, int32(52))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L19
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if l7&int32(8) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = int32(0)
	v129 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v125)+4)) = uint16(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = int32(181)
	if l6 == int32(112) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v137 = int32(222730)
	goto L34
L33:
	;
	v137 = int32(222705)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+8)) = v137
	v140 = F_SystemFuncName(m, int32(314808))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v142 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+48)) = v142
	v145 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+45)) = uint8(v145)
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+44)) = uint8(v147)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+40)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v125)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+26)) = int32(1310720)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+24)) = uint8(v147)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v140
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_CreateTrigger(m, v15+int32(8), v125, v142, v163, v142, v85, l2, v142, v147)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	m.G0 = v15 + int32(32)
	return
L38:
	;
	if l7&int32(3) == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v181 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v183 = int32(0)
	v186 = F_SearchSysCacheCopy(m, int32(34), l2, v183)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L19
	} else {
		goto L41
	}
L41:
	;
	if v186 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+22)))
	v192 = v190 + v191
	if l7&int32(1) == int32(0) {
		v201 = v183
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v41 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+14)))
	if v197 != 0 {
		v201 = v183
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+14)) = uint8(v198)
	v201 = v198
	goto L43
L46:
	;
	F_pfree(m, v186)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L59
	}
L47:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v222 == int32(0) {
		goto L46
	} else {
		goto L57
	}
L48:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L19
	} else {
		goto L56
	}
L49:
	;
	if v201 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L50:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)))
	if v204 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)) = uint8(v207)
	F_CatalogTupleUpdate(m, v181, v186+int32(4), v186)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	if v201 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L47
L54:
	;
	F_CatalogTupleUpdate(m, v181, v186+int32(4), v186)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
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
	v226 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), l2, v226, v226, l9)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	goto L46
L59:
	;
	F_sequence_close(m, v181, int32(3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	goto L37
L61:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(439460), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(487518), int32(1921), int32(351732))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
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
	F_errmsg_internal(m, int32(143893), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(487518), int32(1926), int32(351732))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
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
	F_errmsg_internal(m, int32(39884), v15)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(487518), int32(2070), int32(351732))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
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
	var v352 int32
	_ = v352
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
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
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v881 int32
	_ = v881
	var v891 int32
	_ = v891
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v927 int32
	_ = v927
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v1022 int32
	_ = v1022
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1126 int32
	_ = v1126
	var v1165 int32
	_ = v1165
	var v1173 int32
	_ = v1173
	var v1181 int64
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1218 int32
	_ = v1218
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int64
	_ = v1266
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1305 int32
	_ = v1305
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1466 int32
	_ = v1466
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1704 int32
	_ = v1704
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1796 int32
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1868 int32
	_ = v1868
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1968 int32
	_ = v1968
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2101 int32
	_ = v2101
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2180 int32
	_ = v2180
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2223 int32
	_ = v2223
	var v2264 int32
	_ = v2264
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2355 int32
	_ = v2355
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2457 int32
	_ = v2457
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2499 int32
	_ = v2499
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2571 int32
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2624 int32
	_ = v2624
	var v2626 int32
	_ = v2626
	var v2630 int32
	_ = v2630
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	v22 = int32(0)
	v58 = m.G0
	v60 = v58 - int32(288)
	m.G0 = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l6)+92))
	v66 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
	switch v71 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L4
	default:
		v77 = v22
		goto L3
	}
L3:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v78 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	v77 = base.B2i32(v74 == int32(0))
	goto L3
L5:
	;
	m.G0 = v60 + int32(288)
	return v2641
L6:
	;
	v2414 = *(*int32)(unsafe.Add(mBase, _consts[230]))
	if v2414 != 0 {
		goto L363
	} else {
		goto L364
	}
L7:
	;
	if l3 != 0 {
		goto L335
	} else {
		goto L336
	}
L8:
	;
	F_record_object_address_dependencies(m, v60+int32(192), v1544, int32(97))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L333
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(1259)
	F_add_exact_object_address(m, v60+int32(160), v1544)
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L1
	} else {
		goto L332
	}
L10:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+117)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+118)))
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
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L329
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L325
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if int32(0) < v102 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	v85 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v86) < base.Ui32(int32(12000)) {
		v95 = v85
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v95 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+68))
	if v90 == int32(99) {
		v95 = v85
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v93 = F_isTempToastNamespace(m, v90)
	mBase = m.M
	v95 = v93
	goto L17
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v99 == int32(2) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+128)) = v170
	F_errmsg_internal(m, int32(42149), v60+int32(128))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L323
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L1
	} else {
		goto L319
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L315
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L311
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L308
	}
L27:
	;
	v333 = l16 & int32(2)
	if v333 != 0 {
		goto L74
	} else {
		goto L75
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L69
	}
L29:
	;
	v287 = F_SearchSysCache1(m, int32(14), v170)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L63
	}
L30:
	;
	v127 = int32(0)
	v129 = v102
	goto L33
L31:
	;
	goto L32
L32:
	;
	v243 = l16 & int32(8)
	if v243 != 0 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v164 = v127 << (uint(int32(2)) % 32)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l10+v164)))
	if v166 == int32(0) {
		v180 = v129
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v183 = v127 + int32(1)
	if v183 < v180 {
		v127 = v183
		v129 = v180
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l11+v164)))
	if base.Ui32(int32(2)) < base.Ui32(v170-int32(4217)) {
		v180 = v129
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v175 = F_get_collation_isdeterministic(m, v166)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v175 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v180 = v179
	goto L35
L40:
	;
	goto L34
L41:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v81&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v244) < base.Ui32(int32(12000)) {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v250 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v81&int32(1)&base.B2i32(l9 != int32(1664)) != 0 {
		goto L26
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v256 = F_get_relname_relid(m, l1, v82)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v256 == int32(0) {
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
	v264 = int32(0)
	v267 = F_errstart(m, int32(18), v264)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v267 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_sequence_close(m, v66, int32(3))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+112)) = l1
	F_errmsg(m, int32(330021), v60+int32(112))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(487518), int32(894), int32(351719))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v2641 = v264
	goto L5
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v287 == int32(0) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+144)) = v298 + v299 + int32(8)
	F_errmsg(m, int32(679036), v60+int32(144))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(487518), int32(841), int32(351719))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+96)) = l1
	F_errmsg(m, int32(115592), v60+int32(96))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(487518), int32(902), int32(351719))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L1
	} else {
		goto L304
	}
L74:
	;
	v335 = F_ConstraintNameIsUsed(m, int32(0), v62, l1)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
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
	if v335 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v339 = v338
	goto L81
L80:
	;
	v339 = int32(0)
	goto L81
L81:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v341 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v343 = v342
	goto L84
L83:
	;
	v343 = v22
	goto L84
L84:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v346 = F_GetIndexAmRoutineByAmId(m, l8, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349)+120)))
	v351 = F_CreateTemplateTupleDesc(m, v340)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if int32(0) < v340 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L301
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L298
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L295
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L292
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L1
	} else {
		goto L289
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L1
	} else {
		goto L286
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L283
	}
L94:
	;
	v390 = v22
	v391 = v339
	v397 = v343
	goto L97
L95:
	;
	goto L96
L96:
	;
	v705 = l16 & int32(32)
	F_pfree(m, v346)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L154
	}
L97:
	;
	v425 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v390<<(uint(int32(1))%32)))))
	v427 = v390 * int32(100)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v430 = v428 << (uint(int32(4)) % 32)
	v432 = v427 + (v351 + int32(20) + v430)
	if v432&int32(3) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L96
L99:
	;
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+92)) = uint8(v466)
	v469 = v390 + v466
	*(*uint16)(unsafe.Add(mBase, uint32(v432)+74)) = uint16(v469)
	if v390 < v344 {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	if base.Ui32(v432+int32(100)) <= base.Ui32(v432) {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v463 = F__emscripten_memset_bulkmem(m, v432, base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	goto L108
L103:
	;
	v446 = v427 + (v351 + int32(120)) + v430
	v448 = v427 + (v351 + int32(24)) + v430
	if base.Ui32(v448) < base.Ui32(v446) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v450 = v446
	goto L106
L105:
	;
	v450 = v448
	goto L106
L106:
	;
	v459 = F__emscripten_memset_bulkmem(m, v432, base.I32_extend8_s(int32(0)), (v390*int32(-100)-(v430+v351)+v450-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L107
L107:
	;
	goto L99
L108:
	;
	goto L99
L109:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l10+v390<<(uint(int32(2))%32))))
	v477 = v475
	goto L111
L110:
	;
	v477 = int32(0)
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+96)) = v477
	if v391 == int32(0) {
		goto L87
	} else {
		goto L112
	}
L112:
	;
	v482 = v432 + int32(4)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v485 = F_strncpy(m, v482, v483, int32(64))
	mBase = m.M
	v486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v485)+63)) = uint8(v486)
	goto L113
L113:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	if v425 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = int32(0)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v346)+32))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v390 < v571 {
		goto L129
	} else {
		goto L130
	}
L115:
	;
	if v350 < v425 {
		goto L88
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v397 == int32(0) {
		goto L89
	} else {
		goto L119
	}
L118:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v497 = v348 - int32(80) + v491<<(uint(int32(4))%32) + v425*int32(100)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+68)) = v498
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497)+72)))
	*(*uint16)(unsafe.Add(mBase, uint32(v432)+72)) = uint16(v500)
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v432)+80)) = uint16(v502)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v497)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+76)) = v504
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+82)) = uint8(v506)
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+83)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+83)) = uint8(v508)
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+84)) = uint8(v510)
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+85)) = uint8(v512)
	v562 = v397
	goto L114
L119:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v521 = F_exprType(m, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v523 = F_SearchSysCache1(m, int32(82), v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v523 == int32(0) {
		goto L90
	} else {
		goto L122
	}
L122:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523)+16))
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+68)) = v521
	v530 = v527 + v528
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v432)+72)) = uint16(v531)
	v533 = F_exprTypmod(m, v520)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+76)) = v533
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+82)) = uint8(v536)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+83)) = uint8(v538)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+129)))
	v541 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+85)) = uint8(v541)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+84)) = uint8(v540)
	F_ReleaseCatCache(m, v523)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v432)+96))
	v548 = int32(0)
	F_CheckAttributeType(m, v482, v546, v547, v548, v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v553 = v397 + int32(4)
	if base.Ui32(v553) < base.Ui32(v518+v517<<(uint(int32(2))%32)) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v559 = v553
	goto L128
L127:
	;
	v559 = int32(0)
	goto L128
L128:
	;
	v562 = v559
	goto L114
L129:
	;
	v576 = l11 + v390<<(uint(int32(2))%32)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v578 = F_SearchSysCache1(m, int32(14), v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	v601 = v570
	goto L131
L131:
	;
	if v601 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L132:
	;
	if v578 == int32(0) {
		goto L91
	} else {
		goto L133
	}
L133:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+22)))
	v584 = v582 + v583
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)+92))
	if v585 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	F_ReleaseCatCache(m, v578)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L142
	}
L135:
	;
	v586 = v585
	goto L137
L136:
	;
	v586 = v570
	goto L137
L137:
	;
	if v586 != int32(2283) {
		v598 = v586
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v584)+84))
	if v590 != int32(2277) {
		v598 = int32(2283)
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	v594 = F_get_base_element_type(m, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v594 == int32(0) {
		goto L92
	} else {
		goto L141
	}
L141:
	;
	v598 = v594
	goto L134
L142:
	;
	v601 = v598
	goto L131
L143:
	;
	v637 = v391 + int32(4)
	if base.Ui32(v637) < base.Ui32(v489+v488<<(uint(int32(2))%32)) {
		goto L149
	} else {
		goto L150
	}
L144:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	if v601 == v607 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v610 = F_SearchSysCache1(m, int32(82), v601)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v610 == int32(0) {
		goto L93
	} else {
		goto L147
	}
L147:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610)+16))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v432)+68)) = v601
	v619 = v614 + v615
	v620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v619)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v432)+72)) = uint16(v620)
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+82)) = uint8(v622)
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+83)) = uint8(v624)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+129)))
	v627 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+85)) = uint8(v627)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+84)) = uint8(v626)
	F_ReleaseCatCache(m, v610)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	goto L143
L149:
	;
	v643 = v637
	goto L151
L150:
	;
	v643 = int32(0)
	goto L151
L151:
	;
	F_populate_compact_attribute(m, v351, v390)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v469 != v340 {
		v390 = v469
		v391 = v643
		v397 = v562
		goto L97
	} else {
		goto L153
	}
L153:
	;
	goto L98
L154:
	;
	if l2 != 0 {
		v729 = l2
		v730 = l5
		goto L157
	} else {
		goto L158
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L1
	} else {
		goto L279
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L275
	}
L157:
	;
	v741 = F_heap_create(m, l1, v82, l9, v729, v730, l8, v351, v705^int32(105), v83, v81&int32(1), v77, l18, v60+int32(156), v60+int32(152), base.B2i32(l5 == int32(0)))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L165
	}
L158:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, _consts[79])))
	if v709 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _consts[232]))
	if v713 == int32(0) {
		goto L155
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v727 = F_GetNewRelFileNumber(m, l9, v66, v83)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L164
	}
L162:
	;
	v717 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[232])) = v717
	v720 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	if v705|v720 == v717 {
		goto L156
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[233])) = int32(0)
	v729 = v713
	v730 = v720
	goto L157
L164:
	;
	v729 = v727
	v730 = l5
	goto L157
L165:
	;
	v743 = m.G0
	v745 = v743 - int32(32)
	m.G0 = v745
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v741)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+16)) = v747
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v741)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v745)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v745)+20)) = v749
	v756 = int32(0)
	v761 = F_LockAcquireExtended(m, v745+int32(16), int32(8), v756, v756, v745+int32(12), v756)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	if v761 != int32(3) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	m.G0 = v745 + int32(32)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v741)+48))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v773)+80)) = v775
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v741)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+84)) = l8
	v779 = int32(0)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v741)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v780)+131)) = uint8(base.B2i32(l3 != v779))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v741)+56))
	F_InsertPgClassTuple(m, v66, v741, v784, v779, l15)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v745)+12))
	v768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+53)) = uint8(v768)
	goto L171
L171:
	;
	goto L169
L172:
	;
	F_sequence_close(m, v66, int32(3))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v791 <= int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	if l12 == int32(0) {
		v1218 = v779
		goto L186
	} else {
		goto L187
	}
L175:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v741)+52))
	v796 = v794 + int32(20)
	v798 = v791 & int32(3)
	v799 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v791) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v827 = v799
	v833 = int32(0)
	goto L179
L177:
	;
	v927 = v799
	goto L178
L178:
	;
	if v798 == int32(0) {
		goto L174
	} else {
		goto L182
	}
L179:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	v864 = int32(4)
	v867 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v796+v863<<(uint(v864)%32)+v827*v867))) = v729
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v796+v871<<(uint(v864)%32)+(v827|int32(1))*v867))) = v729
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v796+v881<<(uint(v864)%32)+(v827|int32(2))*v867))) = v729
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v796+v891<<(uint(v864)%32)+(v827|int32(3))*v867))) = v729
	v902 = v827 + v864
	v904 = v833 + v864
	if v904 != v791&int32(2147483644) {
		v827 = v902
		v833 = v904
		goto L179
	} else {
		goto L181
	}
L180:
	;
	v927 = v902
	goto L178
L181:
	;
	goto L180
L182:
	;
	v986 = v927
	v993 = v799
	goto L183
L183:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	*(*int32)(unsafe.Add(mBase, uint32(v796+v1022<<(uint(int32(4))%32)+v986*int32(100)))) = v729
	v1030 = int32(1)
	v1033 = v993 + v1030
	if v1033 != v798 {
		v986 = v986 + v1030
		v993 = v1033
		goto L183
	} else {
		goto L185
	}
L184:
	;
	goto L174
L185:
	;
	goto L184
L186:
	;
	v1252 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L201
	}
L187:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v741)+52))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)))
	v1098 = F_palloc0(m, v1095<<(uint(int32(4))%32))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v741)+52))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)))
	if v1101 <= int32(0) {
		v1218 = v1098
		goto L186
	} else {
		goto L189
	}
L189:
	;
	v1126 = int32(0)
	goto L190
L190:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l12+v1126<<(uint(int32(2))%32))))
	if v1165 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v1218 = v1098
	goto L186
L192:
	;
	if l14 != 0 {
		goto L197
	} else {
		goto L198
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1098+v1126<<(uint(int32(4))%32))+8)) = v1165
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1098+v1126<<(uint(int32(4))%32))+12)) = uint8(v1173)
	goto L192
L196:
	;
	v1189 = v1126 + int32(1)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v741)+52))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	if v1189 < v1191 {
		v1126 = v1189
		goto L190
	} else {
		goto L200
	}
L197:
	;
	v1181 = *(*int64)(unsafe.Add(mBase, uint32(l14+v1126<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1098+v1126<<(uint(int32(4))%32)))) = v1181
	goto L196
L198:
	;
	goto L199
L199:
	;
	v1186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1098+v1126<<(uint(int32(4))%32))+4)) = uint8(v1186)
	goto L196
L200:
	;
	goto L191
L201:
	;
	v1254 = F_CatalogOpenIndexes(m, v1252)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v1256 = int32(0)
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v741)+52))
	F_InsertPgAttributeTuples(m, v1252, v1257, v1256, v1218, v1254)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_CatalogCloseIndexes(m, v1254)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_sequence_close(m, v1252, int32(3))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1266 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+173)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v60)+168)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v60)+160)) = v1266
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v1274 = F_buildint2vector(m, int32(0), v1273)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v1276 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1305 = int32(0)
	goto L210
L208:
	;
	goto L209
L209:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1409 = F_buildoidvector(m, l10, v1408)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L213
	}
L210:
	;
	v1341 = int32(1)
	v1342 = v1305 << (uint(v1341) % 32)
	v1345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1342+(l6+int32(12))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1274+int32(24)+v1342))) = uint16(v1345)
	v1348 = v1305 + v1341
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1348 < v1349 {
		v1305 = v1348
		goto L210
	} else {
		goto L212
	}
L211:
	;
	goto L209
L212:
	;
	goto L211
L213:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1412 = F_buildoidvector(m, l11, v1411)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1415 = F_buildint2vector(m, l13, v1414)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v1417 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1418 = F_nodeToString(m, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	v1425 = v1256
	goto L218
L218:
	;
	v1427 = l16 & int32(1)
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v1433 != 0 {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1420 = F_cstring_to_text(m, v1418)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_pfree(m, v1418)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1425 = v1420
	goto L218
L222:
	;
	v1434 = F_make_ands_explicit(m, v1433)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	v1443 = int32(0)
	goto L224
L224:
	;
	v1446 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L229
	}
L225:
	;
	v1436 = F_nodeToString(m, v1434)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v1438 = F_cstring_to_text(m, v1436)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_pfree(m, v1436)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v1443 = v1438
	goto L224
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+196)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = v729
	v1450 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+200)) = v1450
	v1452 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+204)) = v1452
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+208)) = v1454
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+268)) = v1425
	*(*int32)(unsafe.Add(mBase, uint32(v60)+264)) = v1415
	*(*int32)(unsafe.Add(mBase, uint32(v60)+260)) = v1412
	*(*int32)(unsafe.Add(mBase, uint32(v60)+256)) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v60)+252)) = v1274
	*(*int64)(unsafe.Add(mBase, uint32(v60)+244)) = int64(1)
	v1466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+240)) = base.B2i32(int32(base.Ui32(v243)>>(uint(int32(3))%32)) == v1466)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+236)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(v60)+232)) = base.B2i32(l16&int32(72) == v1466)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+228)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(v60)+224)) = base.B2i32(l17&int32(2) == v1466)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+220)) = base.B2i32(v63 != v1466)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+216)) = v1427
	*(*int32)(unsafe.Add(mBase, uint32(v60)+212)) = v1456
	if v1425 == v1466 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1486 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+179)) = uint8(v1486)
	goto L232
L231:
	;
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+272)) = v1443
	if v1443 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1491 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+180)) = uint8(v1491)
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+52))
	v1498 = F_heap_form_tuple(m, v1493, v60+int32(192), v60+int32(160))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_CatalogTupleInsert(m, v1446, v1498)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_sequence_close(m, v1446, int32(3))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_pfree(m, v1498)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	if l3 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	F_StoreSingleInheritance(m, v729, l3, int32(1))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v1519 == int32(0) {
		goto L6
	} else {
		goto L247
	}
L244:
	;
	F_LockRelationOid(m, l3, int32(4))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_SetRelationHasSubclass(m, l3, int32(1))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	goto L243
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+196)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v60)+192)) = int32(1259)
	if v333 != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L272
	}
L249:
	;
	if v1427 != 0 {
		v1534 = int32(112)
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	v1544 = F_new_object_addresses(m)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L261
	}
L252:
	;
	F_index_constraint_create(m, v60+int32(160), l0, v729, l4, l6, l1, v1534, l17, l18, l19)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L1
	} else {
		goto L259
	}
L253:
	;
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	if v1530 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1531 = int32(117)
	goto L256
L255:
	;
	v1531 = int32(120)
	goto L256
L256:
	;
	if v1530 != 0 {
		v1534 = v1531
		goto L252
	} else {
		goto L257
	}
L257:
	;
	if v63 == int32(0) {
		goto L248
	} else {
		goto L258
	}
L258:
	;
	v1534 = v1531
	goto L252
L259:
	;
	if l20 == int32(0) {
		goto L7
	} else {
		goto L260
	}
L260:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v60)+164))
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v1542
	goto L7
L261:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1546 <= int32(0) {
		goto L9
	} else {
		goto L262
	}
L262:
	;
	v1551 = int32(0)
	v1574 = v1551
	v1577 = v1546
	v1580 = v1551
	goto L263
L263:
	;
	v1613 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v1574<<(uint(int32(1))%32)))))
	if v1613 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L8
L265:
	;
	v1617 = v1574 + int32(1)
	if v1617 < v1577 {
		v1574 = v1617
		goto L263
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = v1613
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(1259)
	F_add_exact_object_address(m, v60+int32(160), v1544)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	if v1580&int32(1) != 0 {
		goto L8
	} else {
		goto L269
	}
L269:
	;
	goto L9
L270:
	;
	v1629 = int32(1)
	v1631 = v1574 + v1629
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1631 < v1632 {
		v1574 = v1631
		v1577 = v1632
		v1580 = v1629
		goto L263
	} else {
		goto L271
	}
L271:
	;
	goto L264
L272:
	;
	F_errmsg_internal(m, int32(536702), int32(0))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(487518), int32(1102), int32(351719))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_errmsg(m, int32(408522), int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(487518), int32(953), int32(351719))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	F_errmsg(m, int32(408838), int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(487518), int32(943), int32(351719))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+48)) = v601
	F_errmsg_internal(m, int32(50097), v60+int32(48))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(487518), int32(466), int32(206237))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v1698
	F_errmsg_internal(m, int32(49999), v60+int32(32))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(487518), int32(452), int32(206237))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v1714
	F_errmsg_internal(m, int32(42047), v60+int32(16))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(487518), int32(433), int32(206237))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v521
	F_errmsg_internal(m, int32(50097), v60)
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(487518), int32(377), int32(206237))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errmsg_internal(m, int32(73972), int32(0))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(487518), int32(367), int32(206237))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v60)+64)) = v425
	F_errmsg_internal(m, int32(467183), v60-int32(-64))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(487518), int32(348), int32(206237))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
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
	F_errmsg_internal(m, int32(74005), int32(0))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(487518), int32(331), int32(206237))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v60)+84)) = v1787 + int32(4)
	F_errmsg(m, int32(115439), v60+int32(80))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(487518), int32(916), int32(351719))
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	F_errmsg_internal(m, int32(414865), int32(0))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(487518), int32(879), int32(351719))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L311:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errmsg(m, int32(498028), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(487518), int32(873), int32(351719))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
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
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	F_errmsg(m, int32(437645), int32(0))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(487518), int32(864), int32(351719))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(437885), int32(0))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(487518), int32(855), int32(351719))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	F_errfinish(m, int32(487518), int32(837), int32(351719))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	F_errmsg(m, int32(439460), int32(0))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(487518), int32(800), int32(351719))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	F_errmsg_internal(m, int32(271423), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(487518), int32(793), int32(351719))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	goto L8
L333:
	;
	F_free_object_addresses(m, v1544)
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	goto L7
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(1259)
	F_recordDependencyOn(m, v60+int32(192), v60+int32(160), int32(80))
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L1
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	v2114 = F_new_object_addresses(m)
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L1
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(1259)
	F_recordDependencyOn(m, v60+int32(192), v60+int32(160), int32(83))
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	goto L337
L340:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2116 <= int32(0) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	F_record_object_address_dependencies(m, v60+int32(192), v2114, int32(110))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L1
	} else {
		goto L355
	}
L342:
	;
	v2141 = int32(0)
	v2143 = v2116
	goto L343
L343:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l10+v2141<<(uint(int32(2))%32))))
	if v2180 == int32(0) {
		v2195 = v2143
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if v2195 <= int32(0) {
		goto L341
	} else {
		goto L350
	}
L345:
	;
	v2197 = v2141 + int32(1)
	if v2197 < v2195 {
		v2141 = v2197
		v2143 = v2195
		goto L343
	} else {
		goto L349
	}
L346:
	;
	if v2180 == int32(100) {
		v2195 = v2143
		goto L345
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = v2180
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(3456)
	F_add_exact_object_address(m, v60+int32(160), v2114)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L1
	} else {
		goto L348
	}
L348:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v2195 = v2194
	goto L345
L349:
	;
	goto L344
L350:
	;
	v2223 = int32(0)
	goto L351
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+160)) = int32(2616)
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(l11+v2223<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+164)) = v2264
	F_add_exact_object_address(m, v60+int32(160), v2114)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L1
	} else {
		goto L353
	}
L352:
	;
	goto L341
L353:
	;
	v2273 = v2223 + int32(1)
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2273 < v2274 {
		v2223 = v2273
		goto L351
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	F_free_object_addresses(m, v2114)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v2340 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	F_recordDependencyOnSingleRelExpr(m, v60+int32(192), v2340, v62, int32(97), int32(0))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v2347 == int32(0) {
		goto L6
	} else {
		goto L361
	}
L360:
	;
	goto L359
L361:
	;
	F_recordDependencyOnSingleRelExpr(m, v60+int32(192), v2347, v62, int32(97), int32(0))
	mBase = m.M
	v2355 = m.ExcPending
	if v2355 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	goto L6
L363:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v729, int32(0), l19)
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L1
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L1
	} else {
		goto L367
	}
L366:
	;
	goto L365
L367:
	;
	v2422 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v2422 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	F_RelationInitIndexAccessInfo(m, v741)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v741)+192))
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v2427)+10)) = uint16(v2428)
	if l12 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L370
L372:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v2563 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L373:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2432 <= int32(0) {
		goto L372
	} else {
		goto L374
	}
L374:
	;
	v2457 = int32(0)
	goto L375
L375:
	;
	v2493 = int32(1)
	v2494 = v2457 + v2493
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(l12+v2457<<(uint(int32(2))%32))))
	v2501 = F_index_opclass_options(m, v741, base.I32_extend16_s(v2494), v2499, v2493)
	mBase = m.M
	v2502 = m.ExcPending
	if v2502 != 0 {
		goto L1
	} else {
		goto L377
	}
L376:
	;
	goto L372
L377:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2494 < v2503 {
		v2457 = v2494
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	F_relation_close(m, v741, int32(0))
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L1
	} else {
		goto L401
	}
L380:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, _consts[234]))
	if v2567 == int32(0) {
		goto L383
	} else {
		goto L384
	}
L381:
	;
	goto L382
L382:
	;
	if l16&int32(4) != 0 {
		goto L395
	} else {
		goto L396
	}
L383:
	;
	v2571 = int32(0)
	v2576 = F_AllocSetContextCreateInternal(m, v2571, int32(539167), v2571, int32(8192), int32(8388608))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L1
	} else {
		goto L386
	}
L384:
	;
	v2579 = v2567
	goto L385
L385:
	;
	v2580 = int32(4480304)
	v2581 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2579
	v2585 = F_palloc(m, int32(16))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L1
	} else {
		goto L387
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, _consts[234])) = v2576
	v2579 = v2576
	goto L385
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+4)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v2585))) = v62
	v2590 = F_palloc(m, int32(144))
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+8)) = v2590
	goto L390
L389:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v2597 = F_copyObjectImpl(m, v2596)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L393
	}
L390:
	;
	v2594 = F__emscripten_memcpy_bulkmem(m, v2590, l6, int32(144))
	mBase = m.M
	goto L392
L392:
	;
	goto L389
L393:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2599)+76)) = v2597
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2601)+80)) = int32(0)
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	v2605 = F_copyObjectImpl(m, v2604)
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2607)+84)) = v2605
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2609)+88)) = int32(0)
	v2612 = int32(4376996)
	v2613 = *(*int32)(unsafe.Add(mBase, _consts[235]))
	*(*int32)(unsafe.Add(mBase, uint32(v2585)+12)) = v2613
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2581
	*(*int32)(unsafe.Add(mBase, _consts[235])) = v2585
	goto L379
L395:
	;
	F_index_update_stats(m, l0, int32(1), float64(-1))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	F_index_build(m, l0, v741, l6, int32(0), int32(1))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L1
	} else {
		goto L400
	}
L398:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	goto L379
L400:
	;
	goto L379
L401:
	;
	v2641 = v729
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[54]))
	if v13 != v9 {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[55]))
		v17 = F_list_member_ptr(m, v16, v9)
		mBase = m.M
		v18 = v17
	} else {
		v18 = int32(1)
	}
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
		if v22 != 0 {
			m.T0[v22].(func(*base.Module, int32, int32))(m, l0, l1)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
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
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v35 + int32(4)
				F_errmsg(m, int32(434269), v7)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errfinish(m, int32(492203), int32(244), int32(230581))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
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
