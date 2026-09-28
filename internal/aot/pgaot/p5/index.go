package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ComputeIndexAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
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
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v670 int64
	_ = v670
	var v671 int32
	_ = v671
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	v19 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(128)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if l8 == v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l15 != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	if l14 == int32(0) {
		v67 = v19
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v53 = F_palloc_mul(m, int32(4), v35)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L10
	}
L5:
	;
	v41 = F_palloc_mul(m, int32(4), v35)
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v41
	v45 = F_palloc_mul(m, int32(4), v35)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v45
	v49 = F_palloc_mul(m, int32(2), v35)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v49
	v67 = v19
	goto L1
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v53
	v57 = F_palloc_mul(m, int32(4), v35)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v57
	v61 = F_palloc_mul(m, int32(2), v35)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	if l14 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = int32(0)
	goto L15
L14:
	;
	v66 = v65
	goto L15
L15:
	;
	v67 = v66
	goto L1
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(124)))) = v73
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(120)))) = v76
	goto L19
L17:
	;
	goto L18
L18:
	;
	if l7 == int32(0) {
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
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L213
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L6
	} else {
		goto L205
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L6
	} else {
		goto L198
	}
L23:
	;
	m.G0 = v33 + int32(128)
	return
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v80 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v86 = l1 + int32(12)
	v109 = v19
	v113 = v67
	goto L26
L26:
	;
	v118 = v109 << (uint(int32(2)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118+v119)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v122 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v118))) = v295
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	if v35 <= v109 {
		goto L81
	} else {
		goto L82
	}
L29:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+22)))
	v267 = v265 + v266
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267)+74)))
	*(*uint16)(unsafe.Add(mBase, uint32(v86+v109<<(uint(int32(1))%32)))) = uint16(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)+96))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)+68))
	F_ReleaseCatCache(m, v123)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L71
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v132
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_0), v33+int32(80))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L68
	}
L31:
	;
	v123 = F_SearchSysCacheAttName(m, l9, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v109 < v35 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	if v123 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if l13 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v132
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_1), v33+int32(96))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(1990), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	v149 = F_exprType(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L63
	}
L45:
	;
	v151 = F_exprCollation(m, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v171 = v148
	goto L47
L47:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v183 != int32(31) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v183 != int32(6) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v171 = v228
	goto L47
L52:
	;
	v199 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86+v109<<(uint(int32(1))%32)))) = uint16(v199)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v202 = F_lappend(m, v201, v171)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+8)))
	if v188 == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v86+v109<<(uint(int32(1))%32)))) = uint16(v188)
	v295 = v149
	v298 = v151
	goto L28
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v202
	v205 = F_contain_mutable_functions_after_planning(m, v171)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	if v205 == int32(0) {
		v295 = v149
		v298 = v151
		goto L28
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_4), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2051), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_5), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2009), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
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
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(1984), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v295 = v271
	v298 = v270
	goto L28
L72:
	;
	v685 = v109 + int32(1)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v685 < v686 {
		v109 = v685
		v113 = v683
		goto L26
	} else {
		goto L197
	}
L73:
	;
	v622 = l6 + v109<<(uint(int32(1))%32)
	v623 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v622))) = uint16(v623)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	if l12 != 0 {
		goto L179
	} else {
		goto L180
	}
L74:
	;
	v583 = F_get_commutator(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L167
	}
L75:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l17)))
	F_AtEOXact_GUC(m, int32(0), v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L161
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L155
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L6
	} else {
		goto L150
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L6
	} else {
		goto L145
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L140
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L6
	} else {
		goto L135
	}
L81:
	;
	if v306 != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v306 == int32(0) {
		v362 = v298
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	if v308 != 0 {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	if v309 != 0 {
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	if v310 != 0 {
		goto L77
	} else {
		goto L87
	}
L87:
	;
	v312 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4+v118))) = v312
	*(*int64)(unsafe.Add(mBase, uint32(l5+v109<<(uint(int32(3))%32)))) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l6+v109<<(uint(int32(1))%32)))) = uint16(v312)
	*(*int32)(unsafe.Add(mBase, uint32(l3+v118))) = v312
	v683 = v113
	goto L72
L88:
	;
	v363 = F_type_is_collatable(m, v295)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L6
	} else {
		goto L101
	}
L89:
	;
	if l15 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v332 = F_get_collation_oid(m, v306, int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l17)))
	F_AtEOXact_GUC(m, int32(0), v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L94
	}
L93:
	;
	v362 = v332
	goto L88
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l16
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l15
	goto L95
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v344 = F_get_collation_oid(m, v342, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v347
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v346
	goto L97
L97:
	;
	v353 = int32(_a_F_ComputeIndexAttrs_6)
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v357 = v355 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v357
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l17))) = v357
	F_RestrictSearchPath(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	v362 = v344
	goto L88
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3+v118))) = v362
	if l15 != 0 {
		goto L113
	} else {
		goto L114
	}
L101:
	;
	if v363 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v362 != 0 {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v362 != 0 {
		goto L76
	} else {
		goto L112
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_7), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	F_errhint(m, int32(_a_F_ComputeIndexAttrs_8), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2127), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	goto L100
L113:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l17)))
	F_AtEOXact_GUC(m, int32(0), v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v398 = l4 + v118
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v121)+20))
	v400 = F_ResolveOpClass(m, v399, v295, l10, l11)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l16
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l15
	goto L117
L117:
	;
	goto L115
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v398))) = v400
	if l15 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v426 = int32(0)
	if l14 == v426 {
		v619 = v426
		goto L73
	} else {
		goto L129
	}
L120:
	;
	if v113 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v411
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v410
	goto L125
L123:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v408 = F_compatible_oper_opid(m, v407, v295, v295)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v582 = v408
	goto L74
L125:
	;
	v417 = int32(_a_F_ComputeIndexAttrs_6)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v421 = v419 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v421
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l17))) = v421
	F_RestrictSearchPath(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	if v113 != 0 {
		goto L75
	} else {
		goto L128
	}
L128:
	;
	goto L119
L129:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if v109 == v35-int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v434 = int32(7)
	goto L132
L131:
	;
	v434 = int32(3)
	goto L132
L132:
	;
	F_GetOperatorFromCompareType(m, v429, int32(0), v434, v33+int32(112), v33+int32(118))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v441+v118))) = v443
	v445 = F_get_opcode(m, v443)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v447+v118))) = v445
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+118)))
	*(*uint16)(unsafe.Add(mBase, uint32(v450+v109<<(uint(int32(1))%32)))) = uint16(v454)
	v619 = v426
	goto L73
L135:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_9), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2067), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_10), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2072), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_11), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2077), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_12), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2082), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v539 = F_format_type_be(m, v295)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v539
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_13), v33-int32(-64))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2136), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = l16
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = l15
	goto L162
L162:
	;
	v564 = F_compatible_oper_opid(m, v555, v295, v295)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[1])) = v567
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[0])) = v566
	goto L164
L164:
	;
	v573 = int32(_a_F_ComputeIndexAttrs_6)
	v575 = *(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2]))
	v577 = v575 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ComputeIndexAttrs[2])) = v577
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l17))) = v577
	F_RestrictSearchPath(m)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	v582 = v564
	goto L74
L167:
	;
	if v583 != v582 {
		goto L22
	} else {
		goto L168
	}
L168:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v587 = F_get_opclass_family(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	v589 = F_get_op_opfamily_strategy(m, v582, v587)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	if v589 == int32(0) {
		goto L21
	} else {
		goto L171
	}
L171:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v593+v118))) = v582
	v596 = F_get_opcode(m, v582)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L6
	} else {
		goto L172
	}
L172:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v118))) = v596
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*uint16)(unsafe.Add(mBase, uint32(v601+v109<<(uint(int32(1))%32)))) = uint16(v589)
	v607 = v113 + int32(4)
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l8)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l8)+4))
	if base.Ui32(v607) < base.Ui32(v609+v610<<(uint(int32(2))%32)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v615 = v607
	goto L175
L174:
	;
	v615 = int32(0)
	goto L175
L175:
	;
	v619 = v615
	goto L73
L176:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v121)+24))
	if v661 != 0 {
		goto L193
	} else {
		goto L194
	}
L177:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v622))) = uint16(v632)
	goto L176
L178:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v121)+28))
	if v656 != int32(2) {
		goto L176
	} else {
		goto L192
	}
L179:
	;
	v626 = int32(2)
	if v625 == v626 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	if v625 != 0 {
		goto L20
	} else {
		goto L185
	}
L182:
	;
	v629 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v622))) = uint16(v629)
	v632 = int32(3)
	goto L184
L183:
	;
	v632 = v626
	goto L184
L184:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	switch v633 {
	case 0:
		goto L178
	case 1:
		goto L177
	default:
		goto L176
	}
L185:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v121)+32))
	if v634 == int32(0) {
		goto L176
	} else {
		goto L186
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = l10
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_14), v33)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2276), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	goto L177
L193:
	;
	v666 = int32(0)
	v670 = F_transformRelOptions(m, int64(0), v661, v666, v666, v666, v666)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5+v109<<(uint(int32(3))%32)))) = int64(0)
	v683 = v619
	goto L72
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5+v109<<(uint(int32(3))%32)))) = v670
	v683 = v619
	goto L72
L197:
	;
	goto L27
L198:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v728 = F_format_operator(m, v582)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v728
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_15), v33+int32(48))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	v738 = F_errdetail(m, int32(_a_F_ComputeIndexAttrs_16), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L202
	}
L202:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2205), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L6
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L6
	} else {
		goto L206
	}
L206:
	;
	v755 = F_format_operator(m, v582)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v757 = F_get_opfamily_name(m, v587)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L6
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v757
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v755
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_17), v33+int32(32))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v768 = F_errdetail(m, int32(_a_F_ComputeIndexAttrs_18), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L6
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2219), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L6
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L6
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = l10
	F_errmsg(m, int32(_a_F_ComputeIndexAttrs_19), v33+int32(16))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v121)+36))
	F_parser_errposition(m, l0, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_ComputeIndexAttrs_2), int32(2270), int32(_a_F_ComputeIndexAttrs_3))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(304)
	m.G0 = v21
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)) = uint16(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(-1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v31 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v38 = v31
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = l1
	if int32(0) < v30 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = v34
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L61
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L56
	}
L8:
	;
	m.G0 = v21 + int32(304)
	return v207
L9:
	;
	v43 = int32(0)
	v51 = v7
	goto L12
L10:
	;
	v182 = v7
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
	v62 = v43 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v29+v62)))
	if v64 == int32(0) {
		v169 = v51
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v182 = v169
	goto L11
L14:
	;
	v172 = v43 + int32(1)
	if v172 != v30 {
		v43 = v172
		v51 = v169
		goto L12
	} else {
		goto L51
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+v28)))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+116)))
	if v69 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+92))
	if v72 == int32(0) {
		v169 = v51
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+118)))
	if v75 != int32(1) {
		v169 = v51
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+192))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v80 = int32(0)
	if l5 == v80 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v64)+192))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+16)))
	if v122 == int32(0) {
		goto L7
	} else {
		goto L38
	}
L24:
	;
	if v118 == int32(0) {
		v169 = v51
		goto L14
	} else {
		goto L37
	}
L25:
	;
	v118 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v86 <= int32(0) {
		v112 = v80
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = v112
	goto L24
L29:
	;
	v89 = int32(0)
	if v89 < v86 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v92 = v86
	goto L32
L31:
	;
	v92 = v89
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v95 = int32(0)
	goto L33
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93+v95<<(uint(int32(2))%32))))
	v104 = base.B2i32(v103 == v79)
	if v103 == v79 {
		v112 = v104
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v112 = v104
	goto L28
L35:
	;
	v106 = v95 + int32(1)
	if v106 != v92 {
		v95 = v106
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
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v68)+84))
	if v125 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v155 = v21 + int32(32)
	F_FormIndexDatum(m, v68, l1, l2, v155, v21)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L48
	}
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v68)+88))
	if v128 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v131 = F_ExecPrepareQual(m, v125, l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v136 = v128
	goto L43
L43:
	;
	v137 = int32(_a_F_ExecCheckIndexConstraints_0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0])) = v140
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v145 = m.T0[v144].(func(*base.Module, int32, int32, int32) int64)(m, v136, v38, v21+int32(303))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+88)) = v131
	if v131 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v131
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecCheckIndexConstraints[0])) = v138
	if v145 != int64(0) {
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v169 = int32(1)
	goto L14
L48:
	;
	v158 = int32(0)
	v159 = int32(1)
	v163 = F_check_exclusion_or_unique_constraint(m, v27, v64, v68, l4, v155, v21, l2, v158, v158, v159, l3)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v163 == int32(0) {
		v207 = v158
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v169 = v159
	goto L14
L51:
	;
	goto L13
L52:
	;
	v207 = int32(1)
	goto L8
L53:
	;
	goto L54
L54:
	;
	v195 = int32(1)
	if v182&v195 == int32(0) {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v207 = v195
	goto L8
L56:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_ExecCheckIndexConstraints_1), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	F_errtableconstraint(m, v27, v233+int32(4))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ExecCheckIndexConstraints_2), int32(611), int32(_a_F_ExecCheckIndexConstraints_3))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
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
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_ExecCheckIndexConstraints_2), int32(657), int32(_a_F_ExecCheckIndexConstraints_3))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
func F_ExecIndexOnlyScanInstrumentEstimate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v6 == int32(0) {
			return
		} else {
			v9 = int32(8)
			v11 = F_mul_size(m, v6, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = F_add_size(m, v9, v11)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v20 = F_add_size(m, v15, (v13+int32(31))&int32(-32))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v20
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v25 = F_add_size(m, v23, int32(1))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v25
							return
						}
					}
				}
			}
		}
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
			F_errfinish(m, int32(_a_F_IndexOnlyRecheck_1), int32(336), int32(_a_F_IndexOnlyRecheck_2))
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
func F_index_beginscan_parallel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v11 = F_RestoreSnapshot(m, l5+int32(28))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_RegisterSnapshot(m, v11)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = F_index_beginscan_internal(m, l1, l3, l4, v11, l5, int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v11
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_index_beginscan_parallel[0]))
				if v24 == int32(0) {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
					v46 = m.T0[v45].(func(*base.Module, int32, int32) int32)(m, l0, l6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v46
						return v18
					}
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_beginscan_parallel[1])))
					if v28&int32(1) != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
						v46 = m.T0[v45].(func(*base.Module, int32, int32) int32)(m, l0, l6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v46
							return v18
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_index_beginscan_parallel_0), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_index_beginscan_parallel_1), int32(1256), int32(_a_F_index_beginscan_parallel_2))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
	}
}
func F_index_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int64
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int64
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v243 int64
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int64
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 float64
	_ = v433
	var v435 int32
	_ = v435
	var v437 float64
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
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
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
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
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	v16 = m.G0
	v18 = v16 - int32(400)
	m.G0 = v18
	if l4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	v38 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[0]))
	if v23 != int32(2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+25)))
	if v27 != int32(1) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v32 = F_plan_create_index_workers(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+128)) = v32
	goto L1
L7:
	;
	if v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(92)))) = v90
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(88)))) = v93
	goto L18
L9:
	;
	F_errfinish(m, int32(_a_F_index_build_0), v79, int32(_a_F_index_build_1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L17
	}
L10:
	;
	if v38 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v38 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v46 + v47
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v45 + v47
	F_errmsg_internal(m, int32(_a_F_index_build_2), v18+int32(48))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v79 = int32(3158)
	goto L9
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v63
	v65 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v62 + v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v61 + v65
	F_errmsg_internal(m, int32(_a_F_index_build_3), v18-int32(-64))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v79 = int32(3164)
	goto L9
L17:
	;
	goto L8
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[2])) = v97 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[1])) = v96
	goto L19
L19:
	;
	v105 = int32(_a_F_index_build_4)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[3]))
	v109 = v107 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[3])) = v109
	goto L20
L20:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	if l5 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v114 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v114
	v117 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v117
	v120 = *(*int64)(unsafe.Add(mBase, _c_F_index_build[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v120
	v122 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+168)) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v18)+160)) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v122
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = int64(2)
	v136 = v18 + int32(96)
	v138 = v18 + int32(128)
	goto L27
L23:
	;
	goto L24
L24:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+36))
	v322 = m.T0[v321].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L42
	}
L25:
	;
	goto L24
L26:
	;
	goto L25
L27:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[7]))
	if v148 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_build[8])))
	if v152&int32(1) == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v157 = int32(_a_F_index_build_5)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[9]))
	v160 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[9])) = v159 + v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v163 + v160
	v167 = int32(0)
	v170 = base.AtomicRmwOr32(m, v167, int32(_a_F_index_build_6), v167)
	goto L31
L30:
	;
	v297 = int32(0)
	v300 = base.AtomicRmwOr32(m, v297, int32(_a_F_index_build_6), v297)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v301 + v302
	v305 = int32(_a_F_index_build_5)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[9])) = v307 - v302
	goto L26
L31:
	;
	v176 = v148 + int32(232)
	goto L32
L32:
	;
	v182 = int32(0)
	v185 = int32(0)
	goto L35
L34:
	;
	v262 = int32(0)
	v265 = v246
	goto L39
L35:
	;
	v191 = int32(2)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v136+v185<<(uint(v191)%32))))
	v195 = int32(3)
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v138+v185<<(uint(v195)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+v194<<(uint(v195)%32)))) = v201
	v204 = v185 | int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v136+v204<<(uint(v191)%32))))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v138+v204<<(uint(v195)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+v208<<(uint(v195)%32)))) = v215
	v218 = v185 | v191
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v136+v218<<(uint(v191)%32))))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v138+v218<<(uint(v195)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+v222<<(uint(v195)%32)))) = v229
	v232 = v185 | v195
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v136+v232<<(uint(v191)%32))))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v138+v232<<(uint(v195)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+v236<<(uint(v195)%32)))) = v243
	v245 = int32(4)
	v246 = v185 + v245
	v248 = v182 + v245
	if v248 != int32(4) {
		v182 = v248
		v185 = v246
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L38
L37:
	;
	goto L36
L38:
	;
	goto L34
L39:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v136+v265<<(uint(int32(2))%32))))
	v275 = int32(3)
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v138+v265<<(uint(v275)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+v274<<(uint(v275)%32)))) = v281
	v283 = int32(1)
	v286 = v262 + v283
	if v286 != int32(2) {
		v262 = v286
		v265 = v265 + v283
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L30
L41:
	;
	goto L40
L42:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324)+118)))
	if v325 != int32(117) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l3 != 0 {
		goto L68
	} else {
		goto L69
	}
L44:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v328 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v354 = v328
	goto L47
L46:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v332
	v336 = F_smgropen(m, v18+int32(32), v329)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L48
	}
L47:
	;
	v356 = F_smgrexists(m, v354, int32(3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L53
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v336
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336)+72))
	if v340 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v354 = v352
	goto L47
L50:
	;
	v348 = v340
	goto L52
L51:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v336)+76))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v336)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v341)+4)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v336)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v336)+72))
	v348 = v346
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336)+72)) = v348 + int32(1)
	goto L49
L53:
	;
	if v356 != 0 {
		goto L43
	} else {
		goto L54
	}
L54:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v358 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v384 = v358
	goto L57
L56:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v360
	v362 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v362
	v366 = F_smgropen(m, v18+int32(16), v359)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L58
	}
L57:
	;
	F_smgrcreate(m, v384, int32(3), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L63
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v366
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366)+72))
	if v370 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v384 = v382
	goto L57
L60:
	;
	v378 = v370
	goto L62
L61:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v366)+76))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v366)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v366)+72))
	v378 = v376
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+72)) = v378 + int32(1)
	goto L59
L63:
	;
	F_log_smgrcreate(m, l1, int32(3))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+40))
	m.T0[v393].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L43
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L5
	} else {
		goto L131
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L128
	}
L68:
	;
	v433 = *(*float64)(unsafe.Add(mBase, uint32(v322)))
	F_index_update_stats(m, l0, int32(1), v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L78
	}
L69:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)))
	if v397&int32(1) == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v402 != 0 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v406 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v411 = F_SearchSysCacheCopy(m, int32(34), base.I64_extend_i32_u(v403), int64(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	if v411 == int32(0) {
		goto L67
	} else {
		goto L74
	}
L74:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v411)+16))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+22)))
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v415+v416)+19)) = uint8(v418)
	F_CatalogTupleUpdate(m, v406, v411+int32(4), v411)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	F_pfree(m, v411)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	F_relation_close(m, v406, int32(3))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	goto L68
L78:
	;
	v437 = *(*float64)(unsafe.Add(mBase, uint32(v322)+8))
	F_index_update_stats(m, l1, int32(0), v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if v442 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[10]))
	if v443 == v445 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	F_AtEOXact_GUC(m, int32(0), v109)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L5
	} else {
		goto L126
	}
L84:
	;
	v448 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[10])) = v448
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[11])) = v448
	goto L86
L85:
	;
	goto L86
L86:
	;
	v453 = F_CreateExecutorState(m)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v453)+152))
	if v455 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v458 = F_MakePerTupleExprContext(m, v453)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L5
	} else {
		goto L91
	}
L89:
	;
	v460 = v455
	goto L90
L90:
	;
	v462 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L5
	} else {
		goto L92
	}
L91:
	;
	v460 = v458
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v460)+4)) = v462
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v466 = F_ExecPrepareQual(m, v465, v453)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v468 = F_GetLatestSnapshot(m)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v470 = F_RegisterSnapshot(m, v468)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[12]))
	if v473 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_build[13])))
	if v475&int32(1) == int32(0) {
		goto L66
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v480 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v486 = m.T0[v485].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v470, v480, v480, v480, int32(449))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L5
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+40)) = v489
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+188))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+20))
	v495 = m.T0[v494].(func(*base.Module, int32, int32, int32) int32)(m, v486, int32(1), v462)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v495 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	goto L105
L103:
	;
	goto L104
L104:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+188))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	m.T0[v573].(func(*base.Module, int32))(m, v486)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L5
	} else {
		goto L122
	}
L105:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[14]))
	if v515 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L104
L107:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L5
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if v466 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L109
L111:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+40)) = v548
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+188))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+20))
	v554 = m.T0[v553].(func(*base.Module, int32, int32, int32) int32)(m, v486, int32(1), v462)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L5
	} else {
		goto L120
	}
L112:
	;
	v518 = int32(_a_F_index_build_7)
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_index_build[15]))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v460)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[15])) = v521
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v466)+24))
	v526 = m.T0[v525].(func(*base.Module, int32, int32, int32) int64)(m, v466, v460, v18+int32(399))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v534 = v18 + int32(128)
	v536 = v18 + int32(96)
	F_FormIndexDatum(m, l2, v462, v453, v534, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L5
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[15])) = v519
	if v526 == int64(0) {
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	F_check_exclusion_constraint(m, l0, l1, l2, v462+int32(32), v534, v536, v453, int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v460)+20))
	F_MemoryContextReset(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	goto L111
L120:
	;
	if v554 != 0 {
		goto L105
	} else {
		goto L121
	}
L121:
	;
	goto L106
L122:
	;
	F_UnregisterSnapshot(m, v470)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_ExecDropSingleTupleTableSlot(m, v462)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_FreeExecutorState(m, v453)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v582
	goto L83
L126:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v18)+92))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[2])) = v605
	*(*int32)(unsafe.Add(mBase, _c_F_index_build[1])) = v604
	goto L127
L127:
	;
	m.G0 = v18 + int32(400)
	return
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v403
	F_errmsg_internal(m, int32(_a_F_index_build_8), v18)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_index_build_0), int32(3261), int32(_a_F_index_build_1))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errmsg_internal(m, int32(_a_F_index_build_9), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_index_build_10), int32(931), int32(_a_F_index_build_11))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
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
					F_errfinish(m, int32(_a_F_index_bulk_delete_2), int32(781), int32(_a_F_index_bulk_delete_3))
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
					F_errfinish(m, int32(_a_F_index_bulk_delete_2), int32(780), int32(_a_F_index_bulk_delete_3))
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
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L19
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L19
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
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
	v188 = F_SearchSysCacheCopy(m, int32(34), base.I64_extend_i32_u(l2), int64(0))
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
	v195 = int32(0)
	if l7&int32(1) == v195 {
		v204 = v195
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
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+14)))
	if v200 != 0 {
		v204 = v195
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+14)) = uint8(v201)
	v204 = v201
	goto L43
L46:
	;
	F_pfree(m, v188)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L19
	} else {
		goto L59
	}
L47:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_index_constraint_create[0]))
	if v225 == int32(0) {
		goto L46
	} else {
		goto L57
	}
L48:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L56
	}
L49:
	;
	if v204 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+16)))
	if v207 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+16)) = uint8(v210)
	F_CatalogTupleUpdate(m, v183, v188+int32(4), v188)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	if v204 != 0 {
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
	v221 = m.ExcPending
	if v221 != 0 {
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
	v229 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), l2, v229, v229, l9)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
	v237 = m.ExcPending
	if v237 != 0 {
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
	v251 = m.ExcPending
	if v251 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(_a_F_index_constraint_create_9), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(1954), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L19
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(1959), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
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
	v281 = m.ExcPending
	if v281 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_index_constraint_create_1), int32(2103), int32(_a_F_index_constraint_create_2))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
func F_index_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int64, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
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
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v857 int32
	_ = v857
	var v873 int32
	_ = v873
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v945 int32
	_ = v945
	var v1000 int32
	_ = v1000
	var v1014 int32
	_ = v1014
	var v1046 int32
	_ = v1046
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1148 int32
	_ = v1148
	var v1183 int32
	_ = v1183
	var v1187 int64
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int64
	_ = v1196
	var v1198 int64
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1215 int32
	_ = v1215
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1323 int32
	_ = v1323
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1440 int64
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1452 int64
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1467 int64
	_ = v1467
	var v1469 int64
	_ = v1469
	var v1471 int64
	_ = v1471
	var v1473 int64
	_ = v1473
	var v1482 int64
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1599 int32
	_ = v1599
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1698 int32
	_ = v1698
	var v1702 int32
	_ = v1702
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1740 int32
	_ = v1740
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1816 int32
	_ = v1816
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1986 int32
	_ = v1986
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2237 int32
	_ = v2237
	var v2275 int32
	_ = v2275
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2343 int32
	_ = v2343
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2362 int32
	_ = v2362
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2463 int32
	_ = v2463
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2502 int64
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
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
	var v2627 int32
	_ = v2627
	var v2634 int32
	_ = v2634
	var v2641 int32
	_ = v2641
	var v2644 int32
	_ = v2644
	v22 = int32(0)
	v56 = m.G0
	v58 = v56 - int32(368)
	m.G0 = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l6)+92))
	v64 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	switch v69 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L4
	default:
		v75 = v22
		goto L3
	}
L3:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v76 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+88))
	v75 = base.B2i32(v72 == int32(0))
	goto L3
L5:
	;
	m.G0 = v58 + int32(368)
	return v2644
L6:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[0]))
	if v2419 != 0 {
		goto L371
	} else {
		goto L372
	}
L7:
	;
	if l3 != 0 {
		goto L343
	} else {
		goto L344
	}
L8:
	;
	F_record_object_address_dependencies(m, v58+int32(192), v1568, int32(97))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L1
	} else {
		goto L341
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(1259)
	F_add_exact_object_address(m, v58+int32(160), v1568)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L340
	}
L10:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+117)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)+68))
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+118)))
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
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L337
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L333
	}
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if int32(0) < v100 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	v83 = int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v84) < base.Ui32(int32(_a_F_index_create_0)) {
		v93 = v83
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v93 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+68))
	if v88 == int32(99) {
		v93 = v83
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v91 = F_isTempToastNamespace(m, v88)
	mBase = m.M
	v93 = v91
	goto L17
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v97 == int32(2) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = v166
	F_errmsg_internal(m, int32(_a_F_index_create_1), v58+int32(128))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L1
	} else {
		goto L331
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L1
	} else {
		goto L327
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L323
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L1
	} else {
		goto L319
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L1
	} else {
		goto L316
	}
L27:
	;
	v328 = l16 & int32(2)
	if v328 != 0 {
		goto L74
	} else {
		goto L75
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L69
	}
L29:
	;
	v282 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v166))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L63
	}
L30:
	;
	v126 = int32(0)
	v127 = v100
	goto L33
L31:
	;
	goto L32
L32:
	;
	v237 = l16 & int32(8)
	if v237 != 0 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v160 = v126 << (uint(int32(2)) % 32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l10+v160)))
	if v162 == int32(0) {
		v176 = v127
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v179 = v126 + int32(1)
	if v179 < v176 {
		v126 = v179
		v127 = v176
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l11+v160)))
	if base.Ui32(int32(2)) < base.Ui32(v166-int32(_a_F_index_create_2)) {
		v176 = v127
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v171 = F_get_collation_isdeterministic(m, v162)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v171 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v176 = v175
	goto L35
L40:
	;
	goto L34
L41:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v79&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v238) < base.Ui32(int32(_a_F_index_create_0)) {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v61 != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v244 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v79&int32(1)&base.B2i32(l9 != int32(1664)) != 0 {
		goto L26
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v250 = F_get_relname_relid(m, l1, v80)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v250 == int32(0) {
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
	v258 = int32(0)
	v261 = F_errstart(m, int32(18), v258)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v261 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errcode(m, int32(117571716))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_relation_close(m, v64, int32(3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+112)) = l1
	F_errmsg(m, int32(_a_F_index_create_3), v58+int32(112))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(905), int32(_a_F_index_create_5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v2644 = v258
	goto L5
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v282 == int32(0) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+16))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+144)) = v293 + v294 + int32(8)
	F_errmsg(m, int32(_a_F_index_create_6), v58+int32(144))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(852), int32(_a_F_index_create_5))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
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
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+96)) = l1
	F_errmsg(m, int32(_a_F_index_create_7), v58+int32(96))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(913), int32(_a_F_index_create_5))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
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
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L1
	} else {
		goto L312
	}
L74:
	;
	v330 = F_ConstraintNameIsUsed(m, int32(0), v60, l1)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	if v330 != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v334 = v333
	goto L81
L80:
	;
	v334 = int32(0)
	goto L81
L81:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v336 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	v338 = v337
	goto L84
L83:
	;
	v338 = v22
	goto L84
L84:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v341 = F_GetIndexAmRoutineByAmId(m, l8, int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v345 = int32(*(*int16)(unsafe.Add(mBase, uint32(v344)+120)))
	v346 = F_CreateTemplateTupleDesc(m, v335)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if int32(0) < v335 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L309
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L1
	} else {
		goto L306
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L303
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L1
	} else {
		goto L300
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L1
	} else {
		goto L297
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L1
	} else {
		goto L294
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L291
	}
L94:
	;
	v377 = v22
	v378 = v334
	v382 = v338
	goto L97
L95:
	;
	goto L96
L96:
	;
	v664 = l16 & int32(32)
	v665 = int32(0)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	if v665 < v674 {
		goto L145
	} else {
		goto L146
	}
L97:
	;
	v407 = int32(1)
	v410 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v377<<(uint(v407)%32)))))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v415 = int32(100)
	v417 = v346 + v411<<(uint(int32(3))%32) + v377*v415
	v419 = v417 + int32(28)
	base.MemoryFill(m, v419, int32(0), v415)
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+120)) = uint8(v407)
	v426 = v377 + v407
	*(*uint16)(unsafe.Add(mBase, uint32(v417)+102)) = uint16(v426)
	if v377 < v339 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l10+v377<<(uint(int32(2))%32))))
	v434 = v432
	goto L101
L100:
	;
	v434 = int32(0)
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419)+96)) = v434
	if v378 == int32(0) {
		goto L87
	} else {
		goto L102
	}
L102:
	;
	v439 = v417 + int32(32)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v442 = F_strncpy(m, v439, v440, int32(64))
	mBase = m.M
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+63)) = uint8(v443)
	goto L103
L103:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	if v410 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = int32(0)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v341)+32))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v377 < v531 {
		goto L119
	} else {
		goto L120
	}
L105:
	;
	if v345 < v410 {
		goto L88
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v382 == int32(0) {
		goto L89
	} else {
		goto L109
	}
L108:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v454 = v343 + v448<<(uint(int32(3))%32) + v410*int32(100)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v454-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+68)) = v457
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454))))
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+72)) = uint16(v459)
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+80)) = uint16(v461)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+76)) = v463
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+82)) = uint8(v465)
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+83)) = uint8(v467)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+84)) = uint8(v469)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+85)) = uint8(v471)
	v522 = v382
	goto L104
L109:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v480 = F_exprType(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v483 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v480))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v483 == int32(0) {
		goto L90
	} else {
		goto L112
	}
L112:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+68)) = v480
	v490 = v487 + v488
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v490)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+72)) = uint16(v491)
	v493 = F_exprTypmod(m, v479)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419)+76)) = v493
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+82)) = uint8(v496)
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+83)) = uint8(v498)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+129)))
	v501 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+85)) = uint8(v501)
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+84)) = uint8(v500)
	F_ReleaseCatCache(m, v483)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v419)+68))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v419)+96))
	v508 = int32(0)
	F_CheckAttributeType(m, v439, v506, v507, v508, v508)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v513 = v382 + int32(4)
	if base.Ui32(v513) < base.Ui32(v477+v476<<(uint(int32(2))%32)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v519 = v513
	goto L118
L117:
	;
	v519 = int32(0)
	goto L118
L118:
	;
	v522 = v519
	goto L104
L119:
	;
	v536 = l11 + v377<<(uint(int32(2))%32)
	v537 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v536))))
	v538 = F_SearchSysCache1(m, int32(14), v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	v561 = v530
	goto L121
L121:
	;
	if v561 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	if v538 == int32(0) {
		goto L91
	} else {
		goto L123
	}
L123:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+22)))
	v544 = v542 + v543
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+92))
	if v545 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	F_ReleaseCatCache(m, v538)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L132
	}
L125:
	;
	v546 = v545
	goto L127
L126:
	;
	v546 = v530
	goto L127
L127:
	;
	if v546 != int32(2283) {
		v558 = v546
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v544)+84))
	if v550 != int32(2277) {
		v558 = int32(2283)
		goto L124
	} else {
		goto L129
	}
L129:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v419)+68))
	v554 = F_get_base_element_type(m, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	if v554 == int32(0) {
		goto L92
	} else {
		goto L131
	}
L131:
	;
	v558 = v554
	goto L124
L132:
	;
	v561 = v558
	goto L121
L133:
	;
	v598 = v378 + int32(4)
	if base.Ui32(v598) < base.Ui32(v446+v445<<(uint(int32(2))%32)) {
		goto L139
	} else {
		goto L140
	}
L134:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v419)+68))
	if v561 == v567 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v571 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v561))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v571 == int32(0) {
		goto L93
	} else {
		goto L137
	}
L137:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v571)+16))
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+76)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+68)) = v561
	v580 = v575 + v576
	v581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+72)) = uint16(v581)
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+82)) = uint8(v583)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+83)) = uint8(v585)
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+129)))
	v588 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+85)) = uint8(v588)
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+84)) = uint8(v587)
	F_ReleaseCatCache(m, v571)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L133
L139:
	;
	v604 = v598
	goto L141
L140:
	;
	v604 = int32(0)
	goto L141
L141:
	;
	F_populate_compact_attribute(m, v346, v377)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v426 != v335 {
		v377 = v426
		v378 = v604
		v382 = v522
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
		v773 = l2
		v774 = l5
		goto L165
	} else {
		goto L166
	}
L145:
	;
	v678 = v346 + int32(28)
	v685 = v665
	v686 = v674
	v688 = v665
	goto L149
L146:
	;
	v742 = v665
	v749 = v674
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+20)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(v346)+16)) = v742
	goto L144
L148:
	;
	v742 = v736
	v749 = v715
	goto L147
L149:
	;
	v694 = v678 + v674<<(uint(int32(3))%32) + v685*int32(100)
	v697 = v678 + v685<<(uint(int32(3))%32)
	if v674 != v686 {
		v715 = v686
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v736 = v674
	goto L148
L151:
	;
	v716 = int32(*(*int16)(unsafe.Add(mBase, uint32(v697)+2)))
	if v716 <= int32(0) {
		v736 = v685
		goto L148
	} else {
		goto L159
	}
L152:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+7)))
	if v699 != int32(118) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v715 = v685
	goto L151
L154:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+4)))
	if v702 != int32(1) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+6)))
	if v705&int32(6) != 0 {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	v708 = int32(*(*int16)(unsafe.Add(mBase, uint32(v697)+2)))
	if v708 <= int32(0) {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+90)))
	if v711 != int32(118) {
		v715 = v674
		goto L151
	} else {
		goto L158
	}
L158:
	;
	goto L153
L159:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+90)))
	if v719 == int32(118) {
		v736 = v685
		goto L148
	} else {
		goto L160
	}
L160:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+5)))
	v728 = (v688 + v722 - int32(1)) & (int32(0) - v722)
	if int32(_a_F_index_create_8) < v728 {
		v736 = v685
		goto L148
	} else {
		goto L161
	}
L161:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v697))) = uint16(v728)
	v734 = v685 + int32(1)
	if v734 != v674 {
		v685 = v734
		v686 = v715
		v688 = v728 + v716
		goto L149
	} else {
		goto L162
	}
L162:
	;
	goto L150
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L287
	}
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L1
	} else {
		goto L283
	}
L165:
	;
	v785 = F_heap_create(m, l1, v80, l9, v773, v774, l8, v346, v664^int32(105), v81, v79&int32(1), v75, l18, v58+int32(156), v58+int32(152), base.B2i32(l5 == int32(0)))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L173
	}
L166:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_create[2])))
	if v753 == int32(1) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[3]))
	if v757 == int32(0) {
		goto L163
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v771 = F_GetNewRelFileNumber(m, l9, v64, v81)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	v761 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[3])) = v761
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[4]))
	if v664|v764 == v761 {
		goto L164
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[4])) = int32(0)
	v773 = v757
	v774 = v764
	goto L165
L172:
	;
	v773 = v771
	v774 = l5
	goto L165
L173:
	;
	v787 = m.G0
	v789 = v787 - int32(32)
	m.G0 = v789
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v785)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v789)+16)) = v791
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v785)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v789)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v789)+20)) = v793
	v800 = int32(0)
	v805 = F_LockAcquireExtended(m, v789+int32(16), int32(8), v800, v800, v789+int32(12), v800)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v805 != int32(3) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	m.G0 = v789 + int32(32)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v785)+48))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v818)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v817)+80)) = v819
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v785)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v821)+84)) = l8
	v823 = int32(0)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v785)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v824)+131)) = uint8(base.B2i32(l3 != v823))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v785)+56))
	F_InsertPgClassTuple(m, v64, v785, v828, int64(0), l15)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v789)+12))
	v812 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v811)+53)) = uint8(v812)
	goto L179
L179:
	;
	goto L177
L180:
	;
	F_relation_close(m, v64, int32(3))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v835 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if l12 == int32(0) {
		v1215 = v823
		goto L194
	} else {
		goto L195
	}
L183:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v785)+52))
	v840 = v835 & int32(3)
	v841 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v835) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v857 = v841
	v873 = int32(0)
	goto L187
L185:
	;
	v945 = v841
	goto L186
L186:
	;
	v1000 = v945
	v1014 = v841
	goto L191
L187:
	;
	v904 = v857 * int32(100)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	v906 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v904+(v838+v905<<(uint(v906)%32)))+28)) = v773
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838+v911<<(uint(v906)%32)+v904)+128)) = v773
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838+v917<<(uint(v906)%32)+v904)+228)) = v773
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838+v923<<(uint(v906)%32)+v904)+328)) = v773
	v929 = int32(4)
	v930 = v857 + v929
	v932 = v873 + v929
	if v932 != v835&int32(2147483644) {
		v857 = v930
		v873 = v932
		goto L187
	} else {
		goto L189
	}
L188:
	;
	if v840 == int32(0) {
		goto L182
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	v945 = v930
	goto L186
L191:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v838)))
	*(*int32)(unsafe.Add(mBase, uint32(v838+v1046<<(uint(int32(3))%32)+v1000*int32(100))+28)) = v773
	v1054 = int32(1)
	v1057 = v1014 + v1054
	if v1057 != v840 {
		v1000 = v1000 + v1054
		v1014 = v1057
		goto L191
	} else {
		goto L193
	}
L192:
	;
	goto L182
L193:
	;
	goto L192
L194:
	;
	v1265 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L209
	}
L195:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v785)+52))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1117)))
	v1119 = F_palloc0_mul(m, int32(32), v1118)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v785)+52))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)))
	if v1122 <= int32(0) {
		v1215 = v1119
		goto L194
	} else {
		goto L197
	}
L197:
	;
	v1148 = int32(0)
	goto L198
L198:
	;
	v1183 = v1119 + v1148<<(uint(int32(5))%32)
	v1187 = *(*int64)(unsafe.Add(mBase, uint32(l12+v1148<<(uint(int32(3))%32))))
	if v1187 != int64(0) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1215 = v1119
	goto L194
L200:
	;
	if l14 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1183)+16)) = v1187
	goto L200
L202:
	;
	goto L203
L203:
	;
	v1191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1183)+24)) = uint8(v1191)
	goto L200
L204:
	;
	v1204 = v1148 + int32(1)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v785)+52))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	if v1204 < v1206 {
		v1148 = v1204
		goto L198
	} else {
		goto L208
	}
L205:
	;
	v1195 = l14 + v1148<<(uint(int32(4))%32)
	v1196 = *(*int64)(unsafe.Add(mBase, uint32(v1195)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1183)+8)) = v1196
	v1198 = *(*int64)(unsafe.Add(mBase, uint32(v1195)))
	*(*int64)(unsafe.Add(mBase, uint32(v1183))) = v1198
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1183)+8)) = uint8(v1200)
	goto L204
L208:
	;
	goto L199
L209:
	;
	v1267 = F_CatalogOpenIndexes(m, v1265)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1269 = int32(0)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v785)+52))
	F_InsertPgAttributeTuples(m, v1265, v1270, v1269, v1215, v1267)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_CatalogCloseIndexes(m, v1267)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_relation_close(m, v1265, int32(3))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v1279 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+173)) = v1279
	*(*int64)(unsafe.Add(mBase, uint32(v58)+168)) = v1279
	*(*int64)(unsafe.Add(mBase, uint32(v58)+160)) = v1279
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v1292 = F_buildint2vector(m, int32(0), v1291)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if int32(0) < v1294 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1323 = v1269
	goto L218
L216:
	;
	goto L217
L217:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1423 = F_buildoidvector(m, l10, v1422)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L1
	} else {
		goto L221
	}
L218:
	;
	v1356 = int32(1)
	v1357 = v1323 << (uint(v1356) % 32)
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1357+(l6+int32(12))))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1292+int32(24)+v1357))) = uint16(v1360)
	v1363 = v1323 + v1356
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1363 < v1364 {
		v1323 = v1363
		goto L218
	} else {
		goto L220
	}
L219:
	;
	goto L217
L220:
	;
	goto L219
L221:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1426 = F_buildoidvector(m, l11, v1425)
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v1429 = F_buildint2vector(m, l13, v1428)
	mBase = m.M
	v1430 = m.ExcPending
	if v1430 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v1431 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1432 = F_nodeToString(m, v1431)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	v1440 = v1279
	goto L226
L226:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v1441 != 0 {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v1434 = F_cstring_to_text(m, v1432)
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_pfree(m, v1432)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v1440 = base.I64_extend_i32_u(v1434)
	goto L226
L230:
	;
	v1442 = F_make_ands_explicit(m, v1441)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	v1452 = int64(0)
	goto L232
L232:
	;
	v1454 = l16 & int32(1)
	v1461 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L237
	}
L233:
	;
	v1444 = F_nodeToString(m, v1442)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v1446 = F_cstring_to_text(m, v1444)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_pfree(m, v1444)
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v1452 = base.I64_extend_i32_u(v1446)
	goto L232
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+200)) = base.I64_extend_i32_u(v60)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+192)) = base.I64_extend_i32_u(v773)
	v1467 = int64(*(*int16)(unsafe.Add(mBase, uint32(l6)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+208)) = v1467
	v1469 = int64(*(*int16)(unsafe.Add(mBase, uint32(l6)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+216)) = v1469
	v1471 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+224)) = v1471
	v1473 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l6)+117)))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+336)) = base.I64_extend_i32_u(v1429)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+328)) = base.I64_extend_i32_u(v1426)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+320)) = base.I64_extend_i32_u(v1423)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+312)) = base.I64_extend_i32_u(v1292)
	v1482 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+304)) = v1482
	*(*int64)(unsafe.Add(mBase, uint32(v58)+296)) = int64(1)
	v1486 = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+288)) = base.I64_extend_i32_u(base.B2i32(v237 == v1486))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+280)) = v1482
	*(*int64)(unsafe.Add(mBase, uint32(v58)+272)) = base.I64_extend_i32_u(base.B2i32(l16&int32(72) == v1486))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+264)) = v1482
	*(*int64)(unsafe.Add(mBase, uint32(v58)+256)) = base.I64_extend_i32_u(base.B2i32(l16&int32(256)|l17&int32(2) == int32(0)))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+248)) = base.I64_extend_i32_u(base.B2i32(v61 != v1486))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+240)) = base.I64_extend_i32_u(v1454)
	*(*int64)(unsafe.Add(mBase, uint32(v58)+232)) = v1473
	*(*int64)(unsafe.Add(mBase, uint32(v58)+344)) = v1440
	if v1440 == v1482 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+179)) = uint8(v1510)
	goto L240
L239:
	;
	goto L240
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v58)+352)) = v1452
	if v1452 == int64(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1515 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+180)) = uint8(v1515)
	goto L243
L242:
	;
	goto L243
L243:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+52))
	v1522 = F_heap_form_tuple(m, v1517, v58+int32(192), v58+int32(160))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_CatalogTupleInsert(m, v1461, v1522)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_relation_close(m, v1461, int32(3))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_pfree(m, v1522)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	if l3 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	F_StoreSingleInheritance(m, v773, l3, int32(1))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v1543 == int32(0) {
		goto L6
	} else {
		goto L255
	}
L252:
	;
	F_LockRelationOid(m, l3, int32(4))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_SetRelationHasSubclass(m, l3, int32(1))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	goto L251
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+200)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+196)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v58)+192)) = int32(1259)
	if v328 != 0 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L1
	} else {
		goto L280
	}
L257:
	;
	if v1454 != 0 {
		v1559 = int32(112)
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v1568 = F_new_object_addresses(m)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L1
	} else {
		goto L269
	}
L260:
	;
	F_index_constraint_create(m, v58+int32(160), l0, v773, l4, l6, l1, v1559, l17, l18, l19)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L267
	}
L261:
	;
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+116)))
	if v1554 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1555 = int32(117)
	goto L264
L263:
	;
	v1555 = int32(120)
	goto L264
L264:
	;
	if v1554 != 0 {
		v1559 = v1555
		goto L260
	} else {
		goto L265
	}
L265:
	;
	if v61 == int32(0) {
		goto L256
	} else {
		goto L266
	}
L266:
	;
	v1559 = v1555
	goto L260
L267:
	;
	if l20 == int32(0) {
		goto L7
	} else {
		goto L268
	}
L268:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v58)+164))
	*(*int32)(unsafe.Add(mBase, uint32(l20))) = v1566
	goto L7
L269:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1570 <= int32(0) {
		goto L9
	} else {
		goto L270
	}
L270:
	;
	v1575 = int32(0)
	v1581 = v1575
	v1586 = v1570
	v1599 = v1575
	goto L271
L271:
	;
	v1635 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6+int32(12)+v1599<<(uint(int32(1))%32)))))
	if v1635 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	goto L8
L273:
	;
	v1639 = v1599 + int32(1)
	if v1639 < v1586 {
		v1599 = v1639
		goto L271
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = v1635
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(1259)
	F_add_exact_object_address(m, v58+int32(160), v1568)
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	if v1581 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	goto L9
L278:
	;
	v1649 = int32(1)
	v1651 = v1599 + v1649
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v1651 < v1652 {
		v1581 = v1649
		v1586 = v1652
		v1599 = v1651
		goto L271
	} else {
		goto L279
	}
L279:
	;
	goto L272
L280:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_9), int32(0))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(1114), int32(_a_F_index_create_5))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(_a_F_index_create_10), int32(0))
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(964), int32(_a_F_index_create_5))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errmsg(m, int32(_a_F_index_create_11), int32(0))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(954), int32(_a_F_index_create_5))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v58)+48)) = v561
	F_errmsg_internal(m, int32(_a_F_index_create_12), v58+int32(48))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(467), int32(_a_F_index_create_13))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
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
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v419)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v1718
	F_errmsg_internal(m, int32(_a_F_index_create_14), v58+int32(32))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(453), int32(_a_F_index_create_13))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v1734
	F_errmsg_internal(m, int32(_a_F_index_create_15), v58+int32(16))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(434), int32(_a_F_index_create_13))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v480
	F_errmsg_internal(m, int32(_a_F_index_create_12), v58)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(378), int32(_a_F_index_create_13))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_16), int32(0))
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(368), int32(_a_F_index_create_13))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+64)) = v410
	F_errmsg_internal(m, int32(_a_F_index_create_17), v58-int32(-64))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(349), int32(_a_F_index_create_13))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_index_create_18), int32(0))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(332), int32(_a_F_index_create_13))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L312:
	;
	F_errcode(m, int32(_a_F_index_create_19))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v58)+84)) = v1807 + int32(4)
	F_errmsg(m, int32(_a_F_index_create_20), v58+int32(80))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(927), int32(_a_F_index_create_5))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_21), int32(0))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(890), int32(_a_F_index_create_5))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
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
	F_errcode(m, int32(325))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_index_create_22), int32(0))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(884), int32(_a_F_index_create_5))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_errmsg(m, int32(_a_F_index_create_23), int32(0))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(875), int32(_a_F_index_create_5))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	F_errmsg(m, int32(_a_F_index_create_24), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(866), int32(_a_F_index_create_5))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(848), int32(_a_F_index_create_5))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	F_errmsg(m, int32(_a_F_index_create_25), int32(0))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(811), int32(_a_F_index_create_5))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_errmsg_internal(m, int32(_a_F_index_create_26), int32(0))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_index_create_4), int32(804), int32(_a_F_index_create_5))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	goto L8
L341:
	;
	F_free_object_addresses(m, v1568)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	goto L7
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(1259)
	v2110 = v58 + int32(192)
	v2112 = v58 + int32(160)
	F_recordDependencyOn(m, v2110, v2112, int32(80))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L1
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v2126 = F_new_object_addresses(m)
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(1259)
	F_recordDependencyOn(m, v2110, v2112, int32(83))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2128 <= int32(0) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v2343 = v58 + int32(192)
	F_record_object_address_dependencies(m, v2343, v2126, int32(110))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L363
	}
L350:
	;
	v2154 = int32(0)
	v2155 = v2128
	goto L351
L351:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(l10+v2154<<(uint(int32(2))%32))))
	v2191 = int32(0)
	if base.B2i32(v2190 == v2191)|base.B2i32(v2190 == int32(100)) == v2191 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	if v2208 <= int32(0) {
		goto L349
	} else {
		goto L358
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v2190
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(3456)
	F_add_exact_object_address(m, v58+int32(160), v2126)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L1
	} else {
		goto L356
	}
L354:
	;
	v2208 = v2155
	goto L355
L355:
	;
	v2210 = v2154 + int32(1)
	if v2210 < v2208 {
		v2154 = v2210
		v2155 = v2208
		goto L351
	} else {
		goto L357
	}
L356:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v2208 = v2207
	goto L355
L357:
	;
	goto L352
L358:
	;
	v2237 = int32(0)
	goto L359
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+160)) = int32(2616)
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l11+v2237<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+168)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+164)) = v2275
	F_add_exact_object_address(m, v58+int32(160), v2126)
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L1
	} else {
		goto L361
	}
L360:
	;
	goto L349
L361:
	;
	v2284 = v2237 + int32(1)
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2284 < v2285 {
		v2237 = v2284
		goto L359
	} else {
		goto L362
	}
L362:
	;
	goto L360
L363:
	;
	F_free_object_addresses(m, v2126)
	mBase = m.M
	v2348 = m.ExcPending
	if v2348 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	if v2349 != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	F_recordDependencyOnSingleRelExpr(m, v2343, v2349, v60, int32(97), int32(0))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L1
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	if v2354 == int32(0) {
		goto L6
	} else {
		goto L369
	}
L368:
	;
	goto L367
L369:
	;
	F_recordDependencyOnSingleRelExpr(m, v58+int32(192), v2354, v60, int32(97), int32(0))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	goto L6
L371:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v773, int32(0), l19)
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L1
	} else {
		goto L375
	}
L374:
	;
	goto L373
L375:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v2427 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	F_RelationInitIndexAccessInfo(m, v785)
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L1
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v2432 = *(*int32)(unsafe.Add(mBase, uint32(v785)+192))
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v2432)+10)) = uint16(v2433)
	if l12 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	goto L378
L380:
	;
	v2564 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[1]))
	if v2564 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L381:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2437 <= int32(0) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v2463 = int32(0)
	goto L383
L383:
	;
	v2496 = int32(1)
	v2497 = v2463 + v2496
	v2502 = *(*int64)(unsafe.Add(mBase, uint32(l12+v2463<<(uint(int32(3))%32))))
	v2504 = F_index_opclass_options(m, v785, base.I32_extend16_s(v2497), v2502, v2496)
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L1
	} else {
		goto L385
	}
L384:
	;
	goto L380
L385:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v2497 < v2506 {
		v2463 = v2497
		goto L383
	} else {
		goto L386
	}
L386:
	;
	goto L384
L387:
	;
	F_relation_close(m, v785, int32(0))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L1
	} else {
		goto L405
	}
L388:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[5]))
	if v2568 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L389:
	;
	goto L390
L390:
	;
	if l16&int32(4) != 0 {
		goto L399
	} else {
		goto L400
	}
L391:
	;
	v2572 = int32(0)
	v2577 = F_AllocSetContextCreateInternal(m, v2572, int32(_a_F_index_create_27), v2572, int32(_a_F_index_create_28), int32(_a_F_index_create_29))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L1
	} else {
		goto L394
	}
L392:
	;
	v2580 = v2568
	goto L393
L393:
	;
	v2581 = int32(_a_F_index_create_30)
	v2582 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[6])) = v2580
	v2586 = F_palloc(m, int32(16))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L1
	} else {
		goto L395
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[5])) = v2577
	v2580 = v2577
	goto L393
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2586)+4)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v2586))) = v60
	v2591 = F_palloc(m, int32(144))
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2586)+8)) = v2591
	base.MemoryCopy(m, v2591, l6, int32(144))
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(l6)+76))
	v2597 = F_copyObjectImpl(m, v2596)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2599)+76)) = v2597
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2601)+80)) = int32(0)
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(l6)+84))
	v2605 = F_copyObjectImpl(m, v2604)
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2607)+84)) = v2605
	v2609 = *(*int32)(unsafe.Add(mBase, uint32(v2586)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2609)+88)) = int32(0)
	v2612 = int32(_a_F_index_create_31)
	v2613 = *(*int32)(unsafe.Add(mBase, _c_F_index_create[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v2586)+12)) = v2613
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[6])) = v2582
	*(*int32)(unsafe.Add(mBase, _c_F_index_create[7])) = v2586
	goto L387
L399:
	;
	F_index_update_stats(m, l0, int32(1), float64(-1))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v2627 = int32(0)
	F_index_build(m, l0, v785, l6, v2627, int32(1), base.B2i32(l16&int32(128) == v2627))
	mBase = m.M
	v2634 = m.ExcPending
	if v2634 != 0 {
		goto L1
	} else {
		goto L404
	}
L402:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	goto L387
L404:
	;
	goto L387
L405:
	;
	v2644 = v773
	goto L5
}
func F_index_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if int32(0) <= base.I32_extend16_s(v7) {
		v11 = int32(8)
	} else {
		v11 = int32(16)
	}
	F_index_deform_tuple_internal(m, l1, l2, l3, l0+v11, l0+int32(8), int32(base.Ui32(v7)>>(uint(int32(15))%32)))
	mBase = m.M
	return
}
func F_index_getnext_slot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = l0 + int32(66)
	goto L2
L1:
	;
	m.G0 = v12 + int32(16)
	return v75
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v75 = int32(1)
	goto L1
L4:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+188))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+56))
	v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v35, l0+int32(60), v36, l2, v17, v12+int32(15))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L10
	}
L5:
	;
	v28 = F_index_getnext_tid(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v28 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v75 = int32(0)
	goto L1
L9:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v64 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if v42 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
	if v47 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+268)))
	if v50 != int32(1) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	v57 = v47
	goto L14
L14:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+32)) = v58 + int64(1)
	goto L9
L15:
	;
	F_pgstat_assoc_relation(m, v46)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+272))
	v57 = v56
	goto L14
L17:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v67)
	goto L19
L18:
	;
	goto L19
L19:
	;
	if v42 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	goto L3
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
					F_errfinish(m, int32(_a_F_index_insert_cleanup_1), int32(245), int32(_a_F_index_insert_cleanup_2))
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
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v5 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v8 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	F_tuplesort_putdatum(m, v3, v4|(v5<<(uint(int64(16))%64)|v8<<(uint(int64(32))%64)), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v18, float64(1))
		return int32(0)
	}
}
