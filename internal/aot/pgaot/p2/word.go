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
	var v130 int32
	_ = v130
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
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
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
	var v260 int32
	_ = v260
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
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
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
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v431 int32
	_ = v431
	var v432 float64
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
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
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 float32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 float32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 float32
	_ = v550
	var v557 int32
	_ = v557
	var v566 float32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 float32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 float32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 float32
	_ = v620
	var v625 float32
	_ = v625
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
	var v702 int32
	_ = v702
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 float32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v750 float32
	_ = v750
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
	if v38 != int32(1) {
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
	v130 = v6
	goto L12
L12:
	;
	v149 = int32(3)
	v151 = v55 + v130<<(uint(v149)%32)
	v154 = v60 + v130*v149
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+2)) = uint8(v155)
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
	*(*uint16)(unsafe.Add(mBase, uint32(v151))) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = int32(-1)
	goto L8
L13:
	;
	v92 = int32(3)
	v94 = v55 + v73<<(uint(v92)%32)
	v97 = v60 + v73*v92
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)) = uint8(v98)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
	*(*uint16)(unsafe.Add(mBase, uint32(v94))) = uint16(v100)
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
	if v38&int32(1) == int32(0) {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v130 = v119
	goto L12
L17:
	;
	F_pg_qsort(m, v55, v52, int32(8), int32(_a_F_calc_word_similarity_0))
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
	v192 = int32(0)
	if v51 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v200 = int32(0)
	v205 = v192
	goto L22
L20:
	;
	v260 = v192
	goto L21
L21:
	;
	v279 = int32(3)
	v281 = v191 + v260<<(uint(v279)%32)
	v284 = v188 + v260*v279
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v281)+2)) = uint8(v285)
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	*(*uint16)(unsafe.Add(mBase, uint32(v281))) = uint16(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v260
	goto L17
L22:
	;
	v224 = int32(3)
	v226 = v191 + v205<<(uint(v224)%32)
	v229 = v188 + v205*v224
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v226)+2)) = uint8(v230)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	*(*uint16)(unsafe.Add(mBase, uint32(v226))) = uint16(v232)
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
	v251 = v200 + v248
	if v251 != v51&int32(2147483646) {
		v200 = v251
		v205 = v249
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v51&int32(1) == int32(0) {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v260 = v249
	goto L21
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
	if base.Ui32(l4&int32(255)) < base.Ui32(int32(2)) {
		goto L53
	} else {
		goto L54
	}
L32:
	;
	v330 = int32(0)
	v401 = v330
	v406 = v330
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
	v401 = int32(0)
	v406 = v343
	goto L31
L40:
	;
	goto L41
L41:
	;
	v348 = int32(0)
	v349 = v342
	v353 = v343
	goto L42
L42:
	;
	v374 = v55 + v349<<(uint(int32(3))%32)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_calc_word_similarity[0]))
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
	v401 = v386
	v406 = v387
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
	v386 = v348 + v382
	v387 = v353 + int32(1)
	goto L47
L46:
	;
	v386 = v348
	v387 = v353
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
		v348 = v386
		v349 = v399
		v353 = v387
		goto L42
	} else {
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v388<<(uint(int32(2))%32)))) = v387
	goto L48
L50:
	;
	goto L51
L51:
	;
	v396 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v387+v326))) = uint8(v396)
	goto L48
L52:
	;
	goto L43
L53:
	;
	v431 = int32(_a_F_calc_word_similarity_1)
	goto L55
L54:
	;
	v431 = int32(_a_F_calc_word_similarity_2)
	goto L55
L55:
	;
	v432 = *(*float64)(unsafe.Add(mBase, uint32(v431)))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v326))))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v437 = v52 << (uint(int32(2)) % 32)
	v438 = F_palloc(m, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v437 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryFill(m, v438, int32(255), v437)
	goto L59
L58:
	;
	goto L59
L59:
	;
	if v51 <= int32(0) {
		v750 = v21
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_pfree(m, v438)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L132
	}
L61:
	;
	v444 = v401 + v434
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
	v449 = int32(0)
	goto L64
L64:
	;
	v450 = int32(1)
	v451 = l4 & v450
	v456 = v449
	v458 = int32(0)
	v464 = v450
	v469 = v6
	v470 = v6
	v475 = v21
	goto L65
L65:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_calc_word_similarity[1]))
	if v479 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v750 = v721
	goto L60
L67:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v324+v458<<(uint(int32(2))%32))))
	if v456 < int32(0) {
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
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v326))))
	if v489 != int32(1) {
		v507 = v469
		v508 = v470
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v494 = v438 + v485<<(uint(int32(2))%32)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v495 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v326))))
	v503 = v469 + v501
	v504 = v470 + int32(1)
	goto L78
L77:
	;
	v503 = v469
	v504 = v470
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494))) = v458
	v507 = v503
	v508 = v504
	goto L71
L79:
	;
	v724 = int32(1)
	v727 = v458 + v724
	if v727 != v51 {
		v456 = v702
		v458 = v727
		v464 = v464 + v724
		v469 = v715
		v470 = v716
		v475 = v721
		goto L65
	} else {
		goto L131
	}
L80:
	;
	v522 = base.B2i32(v456 == int32(-1))
	if v456 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v435))))
	if v512&int32(2) != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v326))))
	if v516 != int32(1) {
		v702 = v456
		v715 = v507
		v716 = v508
		v721 = v475
		goto L79
	} else {
		goto L85
	}
L84:
	;
	v702 = v456
	v715 = v507
	v716 = v508
	v721 = v475
	goto L79
L85:
	;
	goto L80
L86:
	;
	v523 = int32(1)
	goto L88
L87:
	;
	v523 = v508
	goto L88
L88:
	;
	v527 = base.F32_div(base.F32_convert_i32_s(v507), base.F32_convert_i32_s(v523+(v444-v507)))
	if v456 == int32(-1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v528 = v458
	goto L91
L90:
	;
	v528 = v456
	goto L91
L91:
	;
	if v458 < v528 {
		v602 = v528
		v615 = v507
		v616 = v523
		v620 = v527
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if base.F32_lt(v620, v475) != 0 {
		goto L110
	} else {
		goto L111
	}
L93:
	;
	v530 = v523
	v532 = v528
	v533 = v507
	v535 = v528
	v545 = v507
	v546 = v523
	v550 = v527
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
	v602 = v576
	v615 = v577
	v616 = v578
	v620 = v579
	goto L92
L96:
	;
	v581 = int32(2)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v324+v535<<(uint(v581)%32))))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v438+v584<<(uint(v581)%32))))
	if v535 == v588 {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535+v435))))
	if v557&int32(1) == int32(0) {
		v576 = v532
		v577 = v545
		v578 = v546
		v579 = v550
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v566 = base.F32_div(base.F32_convert_i32_s(v533), base.F32_convert_i32_s(v444-v533+v530))
	if base.F32_lt(v550, v566) != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v568 = v535
	v569 = v533
	v570 = v530
	v571 = v566
	goto L103
L102:
	;
	v568 = v532
	v569 = v545
	v570 = v546
	v571 = v550
	goto L103
L103:
	;
	if v451 == int32(0) {
		v576 = v568
		v577 = v569
		v578 = v570
		v579 = v571
		goto L96
	} else {
		goto L104
	}
L104:
	;
	if base.F64_le(v432, base.F64_promote_f32(v571)) != 0 {
		v602 = v568
		v615 = v569
		v616 = v570
		v620 = v571
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v576 = v568
	v577 = v569
	v578 = v570
	v579 = v571
	goto L96
L106:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v584))))
	v595 = v530 - int32(1)
	v596 = v533 - v591
	goto L108
L107:
	;
	v595 = v530
	v596 = v533
	goto L108
L108:
	;
	v598 = v535 + int32(1)
	if v598 != v464 {
		v530 = v595
		v532 = v576
		v533 = v596
		v535 = v598
		v545 = v577
		v546 = v578
		v550 = v579
		goto L94
	} else {
		goto L109
	}
L109:
	;
	goto L95
L110:
	;
	v625 = v475
	goto L112
L111:
	;
	v625 = v620
	goto L112
L112:
	;
	if base.F64_le(v432, base.F64_promote_f32(v625))&v451 != 0 {
		v750 = v625
		goto L60
	} else {
		goto L113
	}
L113:
	;
	if v602 <= v528 {
		v702 = v602
		v715 = v615
		v716 = v616
		v721 = v625
		goto L79
	} else {
		goto L114
	}
L114:
	;
	v630 = int32(1)
	v631 = v528 + v630
	if (v602-v528)&v630 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v635 = int32(2)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v324+v528<<(uint(v635)%32))))
	v641 = v438 + v638<<(uint(v635)%32)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v528 == v642 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v646 = v528
	goto L117
L117:
	;
	if v631 == v602 {
		v702 = v602
		v715 = v615
		v716 = v616
		v721 = v625
		goto L79
	} else {
		goto L121
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641))) = int32(-1)
	goto L120
L119:
	;
	goto L120
L120:
	;
	v646 = v631
	goto L117
L121:
	;
	v650 = v646
	goto L122
L122:
	;
	v673 = int32(2)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v324+v650<<(uint(v673)%32))))
	v679 = v438 + v676<<(uint(v673)%32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	if v650 == v680 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v702 = v602
	v715 = v615
	v716 = v616
	v721 = v625
	goto L79
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v679))) = int32(-1)
	goto L126
L125:
	;
	goto L126
L126:
	;
	v685 = v650 + int32(1)
	v686 = int32(2)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v324+v685<<(uint(v686)%32))))
	v692 = v438 + v689<<(uint(v686)%32)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	if v693 == v685 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v692))) = int32(-1)
	goto L129
L128:
	;
	goto L129
L129:
	;
	v698 = v650 + int32(2)
	if v698 != v602 {
		v650 = v698
		goto L122
	} else {
		goto L130
	}
L130:
	;
	goto L123
L131:
	;
	goto L66
L132:
	;
	F_pfree(m, v324)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v326)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_pfree(m, v55)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	m.G0 = v27 + int32(32)
	return v750
}
func F_word_is_not_variable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_errstart_cold(m, int32(21), int32(_a_F_word_is_not_variable_0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
			F_errmsg(m, int32(_a_F_word_is_not_variable_1), v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = F_plpgsql_scanner_errposition(m, l1, l2)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_word_is_not_variable_2), int32(2635), int32(_a_F_word_is_not_variable_3))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
