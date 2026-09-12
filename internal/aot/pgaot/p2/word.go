package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_calc_word_similarity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 float32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v429 int32
	_ = v429
	var v430 float64
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 float32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 float32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 float32
	_ = v550
	var v556 int32
	_ = v556
	var v565 float32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 float32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 float32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 float32
	_ = v620
	var v625 float32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 float32
	_ = v720
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v749 float32
	_ = v749
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	v6 = int32(0)
	v21 = float32(0)
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v6
	F_generate_trgm_only(m, v27+int32(20), l0, l1, v6)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return float32(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if base.Ui32(int32(2)) <= base.Ui32(l4) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v46 = v27 + int32(4)
	goto L5
L4:
	;
	v46 = int32(0)
	goto L5
L5:
	;
	F_generate_trgm_only(m, v27+int32(8), l2, l3, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v52 = v38 + v51
	v55 = F_palloc(m, v52<<(uint(int32(3))%32))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v38 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v51 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v60 = v50 + int32(5)
	v61 = int32(1)
	if v38 != v61 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v71 = int32(0)
	v73 = v6
	goto L13
L11:
	;
	v128 = v6
	goto L12
L12:
	;
	if v38&v61 == int32(0) {
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v92 = int32(3)
	v94 = v55 + v73<<(uint(v92)%32)
	v97 = v60 + v73*v92
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	*(*uint16)(unsafe.Add(mBase, uint32(v94))) = uint16(v98)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)) = uint8(v100)
	v102 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v102
	v105 = v73 | int32(1)
	v108 = v55 + v105<<(uint(v92)%32)
	v111 = v60 + v105*v92
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v108)+2)) = uint8(v112)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111))))
	*(*uint16)(unsafe.Add(mBase, uint32(v108))) = uint16(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v102
	v118 = int32(2)
	v119 = v73 + v118
	v121 = v71 + v118
	if v121 != v38&int32(2147483646) {
		v71 = v121
		v73 = v119
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v128 = v119
	goto L12
L15:
	;
	goto L14
L16:
	;
	v149 = int32(3)
	v151 = v55 + v128<<(uint(v149)%32)
	v154 = v60 + v128*v149
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
	*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)) = uint8(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = int32(-1)
	goto L8
L17:
	;
	F_pg_qsort(m, v55, v52, int32(8), int32(6922))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v188 = v49 + int32(5)
	v191 = v55 + v38<<(uint(int32(3))%32)
	v192 = int32(1)
	v194 = int32(0)
	if v51 != v192 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v205 = v194
	v207 = int32(0)
	goto L22
L20:
	;
	v258 = v194
	goto L21
L21:
	;
	if v51&v192 == int32(0) {
		goto L17
	} else {
		goto L25
	}
L22:
	;
	v224 = int32(3)
	v226 = v191 + v205<<(uint(v224)%32)
	v229 = v188 + v205*v224
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	*(*uint16)(unsafe.Add(mBase, uint32(v226))) = uint16(v230)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226)+2)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v226)+4)) = v205
	v236 = v205 | int32(1)
	v239 = v191 + v236<<(uint(v224)%32)
	v242 = v188 + v236*v224
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+2)) = uint8(v243)
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242))))
	*(*uint16)(unsafe.Add(mBase, uint32(v239))) = uint16(v245)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v236
	v248 = int32(2)
	v249 = v205 + v248
	v251 = v207 + v248
	if v251 != v51&int32(2147483646) {
		v205 = v249
		v207 = v251
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v258 = v249
	goto L21
L24:
	;
	goto L23
L25:
	;
	v279 = int32(3)
	v281 = v191 + v258<<(uint(v279)%32)
	v284 = v188 + v258*v279
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	*(*uint16)(unsafe.Add(mBase, uint32(v281))) = uint16(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+2)) = uint8(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v258
	goto L17
L26:
	;
	F_pfree(m, v50)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v49)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v324 = F_palloc(m, v51<<(uint(int32(2))%32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v326 = F_palloc0(m, v52)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v52 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if base.Ui32(l4) < base.Ui32(int32(2)) {
		goto L53
	} else {
		goto L54
	}
L32:
	;
	v330 = int32(0)
	v406 = v330
	v408 = v330
	goto L31
L33:
	;
	goto L34
L34:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v332 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v342 = int32(1)
	v343 = int32(0)
	if v52 == v342 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v335 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v326))) = uint8(v335)
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v332<<(uint(int32(2))%32)))) = int32(0)
	goto L35
L39:
	;
	v406 = v343
	v408 = int32(0)
	goto L31
L40:
	;
	goto L41
L41:
	;
	v349 = v342
	v353 = v343
	v355 = int32(0)
	goto L42
L42:
	;
	v374 = v55 + v349<<(uint(int32(3))%32)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[1395]))
	v379 = m.T0[v378].(func(*base.Module, int32, int32) int32)(m, v374-int32(8), v374)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v406 = v386
	v408 = v387
	goto L31
L44:
	;
	if v379 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v326))))
	v386 = v353 + int32(1)
	v387 = v355 + v382
	goto L47
L46:
	;
	v386 = v353
	v387 = v355
	goto L47
L47:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if int32(0) <= v388 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v399 = v349 + int32(1)
	if v399 != v52 {
		v349 = v399
		v353 = v386
		v355 = v387
		goto L42
	} else {
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v388<<(uint(int32(2))%32)))) = v386
	goto L48
L50:
	;
	goto L51
L51:
	;
	v396 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v386+v326))) = uint8(v396)
	goto L48
L52:
	;
	goto L43
L53:
	;
	v429 = int32(4396080)
	goto L55
L54:
	;
	v429 = int32(4396088)
	goto L55
L55:
	;
	v430 = *(*float64)(unsafe.Add(mBase, uint32(v429)))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v326))))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v435 = v52 << (uint(int32(2)) % 32)
	v436 = F_palloc(m, v435)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v440 = F__emscripten_memset_bulkmem(m, v436, base.I32_extend8_s(int32(255)), v435)
	mBase = m.M
	goto L57
L57:
	;
	if v51 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_pfree(m, v440)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L135
	}
L59:
	;
	v749 = v21
	goto L58
L60:
	;
	goto L61
L61:
	;
	v443 = v432 + v408
	v444 = int32(0)
	v448 = base.B2i32(base.Ui32(l4) < base.Ui32(int32(2)))
	if base.Ui32(l4) < base.Ui32(int32(2)) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v449 = int32(-1)
	goto L64
L63:
	;
	v449 = v444
	goto L64
L64:
	;
	v450 = int32(1)
	v451 = l4 & v450
	v459 = v449
	v463 = v444
	v467 = v450
	v468 = v6
	v469 = v6
	v473 = v21
	goto L65
L65:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v478 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v749 = v720
	goto L58
L67:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v324+v463<<(uint(int32(2))%32))))
	if v459 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L69
L71:
	;
	if v448 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v326))))
	if v488 != int32(1) {
		v506 = v468
		v507 = v469
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v493 = v440 + v484<<(uint(int32(2))%32)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	if v494 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v326))))
	v502 = v468 + int32(1)
	v503 = v469 + v498
	goto L78
L77:
	;
	v502 = v468
	v503 = v469
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v463
	v506 = v502
	v507 = v503
	goto L71
L79:
	;
	v724 = int32(1)
	v727 = v463 + v724
	if v727 != v51 {
		v459 = v706
		v463 = v727
		v467 = v467 + v724
		v468 = v715
		v469 = v716
		v473 = v720
		goto L65
	} else {
		goto L134
	}
L80:
	;
	v521 = base.B2i32(v459 == int32(-1))
	if v459 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v433))))
	if v511&int32(2) != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v326))))
	if v515 != int32(1) {
		v706 = v459
		v715 = v506
		v716 = v507
		v720 = v473
		goto L79
	} else {
		goto L85
	}
L84:
	;
	v706 = v459
	v715 = v506
	v716 = v507
	v720 = v473
	goto L79
L85:
	;
	goto L80
L86:
	;
	v522 = int32(1)
	goto L88
L87:
	;
	v522 = v506
	goto L88
L88:
	;
	v526 = base.F32_div(base.F32_convert_i32_s(v507), base.F32_convert_i32_s(v522+(v443-v507)))
	if v459 == int32(-1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v527 = v463
	goto L91
L90:
	;
	v527 = v459
	goto L91
L91:
	;
	if v463 < v527 {
		v605 = v527
		v614 = v522
		v615 = v507
		v620 = v526
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if base.F32_gt(v473, v620) != 0 {
		goto L110
	} else {
		goto L111
	}
L93:
	;
	v532 = v507
	v534 = v527
	v535 = v527
	v536 = v522
	v544 = v522
	v545 = v507
	v550 = v526
	goto L94
L94:
	;
	if v448 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v605 = v575
	v614 = v576
	v615 = v577
	v620 = v578
	goto L92
L96:
	;
	v580 = int32(2)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v324+v534<<(uint(v580)%32))))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v440+v583<<(uint(v580)%32))))
	if v534 == v587 {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534+v433))))
	if v556&int32(1) == int32(0) {
		v575 = v535
		v576 = v544
		v577 = v545
		v578 = v550
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v565 = base.F32_div(base.F32_convert_i32_s(v532), base.F32_convert_i32_s(v443-v532+v536))
	if base.F32_lt(v550, v565) != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v567 = v534
	v568 = v536
	v569 = v532
	v570 = v565
	goto L103
L102:
	;
	v567 = v535
	v568 = v544
	v569 = v545
	v570 = v550
	goto L103
L103:
	;
	if v451 == int32(0) {
		v575 = v567
		v576 = v568
		v577 = v569
		v578 = v570
		goto L96
	} else {
		goto L104
	}
L104:
	;
	if base.F64_le(v430, base.F64_promote_f32(v570)) != 0 {
		v605 = v567
		v614 = v568
		v615 = v569
		v620 = v570
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v575 = v567
	v576 = v568
	v577 = v569
	v578 = v570
	goto L96
L106:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v583))))
	v594 = v532 - v592
	v595 = v536 - int32(1)
	goto L108
L107:
	;
	v594 = v532
	v595 = v536
	goto L108
L108:
	;
	v597 = v534 + int32(1)
	if v597 != v467 {
		v532 = v594
		v534 = v597
		v535 = v575
		v536 = v595
		v544 = v576
		v545 = v577
		v550 = v578
		goto L94
	} else {
		goto L109
	}
L109:
	;
	goto L95
L110:
	;
	v625 = v473
	goto L112
L111:
	;
	v625 = v620
	goto L112
L112:
	;
	if base.F64_le(v430, base.F64_promote_f32(v625)) != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v628 = v451
	goto L115
L114:
	;
	v628 = int32(0)
	goto L115
L115:
	;
	if v628 != 0 {
		v749 = v625
		goto L58
	} else {
		goto L116
	}
L116:
	;
	if v605 <= v527 {
		v706 = v605
		v715 = v614
		v716 = v615
		v720 = v625
		goto L79
	} else {
		goto L117
	}
L117:
	;
	v630 = int32(1)
	v631 = v527 + v630
	if (v605-v527)&v630 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v635 = int32(2)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v324+v527<<(uint(v635)%32))))
	v641 = v440 + v638<<(uint(v635)%32)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v527 == v642 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v646 = v527
	goto L120
L120:
	;
	if v631 == v605 {
		v706 = v605
		v715 = v614
		v716 = v615
		v720 = v625
		goto L79
	} else {
		goto L124
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641))) = int32(-1)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v646 = v631
	goto L120
L124:
	;
	v650 = v646
	goto L125
L125:
	;
	v673 = int32(2)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v324+v650<<(uint(v673)%32))))
	v679 = v440 + v676<<(uint(v673)%32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	if v650 == v680 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v706 = v605
	v715 = v614
	v716 = v615
	v720 = v625
	goto L79
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679))) = int32(-1)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v685 = v650 + int32(1)
	v686 = int32(2)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v324+v685<<(uint(v686)%32))))
	v692 = v440 + v689<<(uint(v686)%32)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	if v693 == v685 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692))) = int32(-1)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v698 = v650 + int32(2)
	if v698 != v605 {
		v650 = v698
		goto L125
	} else {
		goto L133
	}
L133:
	;
	goto L126
L134:
	;
	goto L66
L135:
	;
	F_pfree(m, v324)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_pfree(m, v326)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_pfree(m, v55)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	m.G0 = v27 + int32(32)
	return v749
}
func F_word_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
			F_errmsg(m, int32(396900), v8)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = F_plpgsql_scanner_errposition(m, l1, l2)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(26975), int32(2635), int32(396280))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
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
