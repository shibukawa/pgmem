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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v430 int32
	_ = v430
	var v431 float64
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v53 = v38 + v52
	v54 = F_palloc_mul(m, int32(8), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	if v52 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v59 = v50 + int32(5)
	if v38 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = int32(0)
	v72 = v6
	goto L13
L11:
	;
	v129 = v6
	goto L12
L12:
	;
	v148 = int32(3)
	v150 = v54 + v129<<(uint(v148)%32)
	v153 = v59 + v129*v148
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+2)) = uint8(v154)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	*(*uint16)(unsafe.Add(mBase, uint32(v150))) = uint16(v156)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(-1)
	goto L8
L13:
	;
	v91 = int32(3)
	v93 = v54 + v72<<(uint(v91)%32)
	v96 = v59 + v72*v91
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+2)) = uint8(v97)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
	*(*uint16)(unsafe.Add(mBase, uint32(v93))) = uint16(v99)
	v101 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v101
	v104 = v72 | int32(1)
	v107 = v54 + v104<<(uint(v91)%32)
	v110 = v59 + v104*v91
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)) = uint8(v111)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110))))
	*(*uint16)(unsafe.Add(mBase, uint32(v107))) = uint16(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v101
	v117 = int32(2)
	v118 = v72 + v117
	v120 = v70 + v117
	if v120 != v38&int32(2147483646) {
		v70 = v120
		v72 = v118
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
	v129 = v118
	goto L12
L17:
	;
	F_pg_qsort(m, v54, v53, int32(8), int32(_a_F_calc_word_similarity_0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L26
	}
L18:
	;
	v187 = v49 + int32(5)
	v190 = v54 + v38<<(uint(int32(3))%32)
	v191 = int32(0)
	if v52 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v199 = int32(0)
	v204 = v191
	goto L22
L20:
	;
	v259 = v191
	goto L21
L21:
	;
	v278 = int32(3)
	v280 = v190 + v259<<(uint(v278)%32)
	v283 = v187 + v259*v278
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v280)+2)) = uint8(v284)
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283))))
	*(*uint16)(unsafe.Add(mBase, uint32(v280))) = uint16(v286)
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v259
	goto L17
L22:
	;
	v223 = int32(3)
	v225 = v190 + v204<<(uint(v223)%32)
	v228 = v187 + v204*v223
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+2)) = uint8(v229)
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228))))
	*(*uint16)(unsafe.Add(mBase, uint32(v225))) = uint16(v231)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v204
	v235 = v204 | int32(1)
	v238 = v190 + v235<<(uint(v223)%32)
	v241 = v187 + v235*v223
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)) = uint8(v242)
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241))))
	*(*uint16)(unsafe.Add(mBase, uint32(v238))) = uint16(v244)
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = v235
	v247 = int32(2)
	v248 = v204 + v247
	v250 = v199 + v247
	if v250 != v52&int32(2147483646) {
		v199 = v250
		v204 = v248
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v52&int32(1) == int32(0) {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v259 = v248
	goto L21
L26:
	;
	F_pfree(m, v50)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v49)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v322 = F_palloc_mul(m, int32(4), v52)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v325 = F_palloc0_mul(m, int32(1), v53)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v53 <= int32(0) {
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
	v329 = int32(0)
	v400 = v329
	v405 = v329
	goto L31
L33:
	;
	goto L34
L34:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v331 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v341 = int32(1)
	v342 = int32(0)
	if v53 == v341 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v325))) = uint8(v334)
	goto L35
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322+v331<<(uint(int32(2))%32)))) = int32(0)
	goto L35
L39:
	;
	v400 = int32(0)
	v405 = v342
	goto L31
L40:
	;
	goto L41
L41:
	;
	v347 = int32(0)
	v348 = v341
	v352 = v342
	goto L42
L42:
	;
	v373 = v54 + v348<<(uint(int32(3))%32)
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_calc_word_similarity[0]))
	v378 = m.T0[v377].(func(*base.Module, int32, int32) int32)(m, v373-int32(8), v373)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v400 = v385
	v405 = v386
	goto L31
L44:
	;
	if v378 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352+v325))))
	v385 = v347 + v381
	v386 = v352 + int32(1)
	goto L47
L46:
	;
	v385 = v347
	v386 = v352
	goto L47
L47:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if int32(0) <= v387 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v398 = v348 + int32(1)
	if v398 != v53 {
		v347 = v385
		v348 = v398
		v352 = v386
		goto L42
	} else {
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322+v387<<(uint(int32(2))%32)))) = v386
	goto L48
L50:
	;
	goto L51
L51:
	;
	v395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v386+v325))) = uint8(v395)
	goto L48
L52:
	;
	goto L43
L53:
	;
	v430 = int32(_a_F_calc_word_similarity_1)
	goto L55
L54:
	;
	v430 = int32(_a_F_calc_word_similarity_2)
	goto L55
L55:
	;
	v431 = *(*float64)(unsafe.Add(mBase, uint32(v430)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405+v325))))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v436 = F_palloc_mul(m, int32(4), v53)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v439 = v53 << (uint(int32(2)) % 32)
	if v439 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryFill(m, v436, int32(255), v439)
	goto L59
L58:
	;
	goto L59
L59:
	;
	if v52 <= int32(0) {
		v750 = v21
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_pfree(m, v436)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L132
	}
L61:
	;
	v444 = v400 + v433
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
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v322+v458<<(uint(int32(2))%32))))
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
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v325))))
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
	v494 = v436 + v485<<(uint(int32(2))%32)
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
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v325))))
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
	if v727 != v52 {
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
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v434))))
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
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v325))))
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
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v322+v535<<(uint(v581)%32))))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v436+v584<<(uint(v581)%32))))
	if v535 == v588 {
		goto L106
	} else {
		goto L107
	}
L97:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535+v434))))
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
	if base.F64_le(v431, base.F64_promote_f32(v571)) != 0 {
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
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325+v584))))
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
	if base.F64_le(v431, base.F64_promote_f32(v625))&v451 != 0 {
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
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v322+v528<<(uint(v635)%32))))
	v641 = v436 + v638<<(uint(v635)%32)
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
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v322+v650<<(uint(v673)%32))))
	v679 = v436 + v676<<(uint(v673)%32)
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
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v322+v685<<(uint(v686)%32))))
	v692 = v436 + v689<<(uint(v686)%32)
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
	F_pfree(m, v322)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v325)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_pfree(m, v54)
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
					F_errfinish(m, int32(_a_F_word_is_not_variable_2), int32(2639), int32(_a_F_word_is_not_variable_3))
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
