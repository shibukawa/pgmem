package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NameListToString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	m.G0 = v9 + int32(32)
	return v80
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v24 != int32(476) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v37 = int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= v37 {
		goto L4
	} else {
		goto L14
	}
L8:
	;
	if v24 != int32(77) {
		v87 = v23
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_appendStringInfoString(m, v9+int32(16), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_appendStringInfoChar(m, v12, int32(42))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	goto L7
L14:
	;
	v44 = v37
	goto L15
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
	v53 = v9 + int32(16)
	F_appendStringInfoChar(m, v53, int32(46))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L4
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v57 != int32(77) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v71 = v44 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v44 = v71
		goto L15
	} else {
		goto L25
	}
L19:
	;
	if v57 != int32(476) {
		v87 = v51
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoChar(m, v9+int32(16), int32(42))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	F_appendStringInfoString(m, v53, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L18
L24:
	;
	goto L18
L25:
	;
	goto L16
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v95
	F_errmsg_internal(m, int32(_a_F_NameListToString_0), v9)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_NameListToString_1), int32(3686), int32(_a_F_NameListToString_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_name_for_var_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
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
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	v14 = m.G0
	v16 = v14 - int32(256)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != int32(8) {
		goto L15
	} else {
		goto L16
	}
L1:
	;
	m.G0 = v16 + int32(256)
	return v860
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L22
	} else {
		goto L249
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L22
	} else {
		goto L246
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L22
	} else {
		goto L243
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L22
	} else {
		goto L240
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L22
	} else {
		goto L237
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L22
	} else {
		goto L234
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L22
	} else {
		goto L231
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L22
	} else {
		goto L228
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L22
	} else {
		goto L225
	}
L11:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v16)+248))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)))
	v695 = v16 + int32(168)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v16)+252))
	base.MemoryCopy(m, v695, v696, int32(80))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v696)+44))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
	v706 = F_list_copy_tail(m, v699, (v692-v700)>>(uint(int32(2))%32)+int32(1))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L22
	} else {
		goto L221
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v68 = v67 + l2
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v69 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v56 = F_get_expr_result_tupdesc(m, l0, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L22
	} else {
		goto L27
	}
L14:
	;
	if v47 != int32(6) {
		goto L13
	} else {
		goto L25
	}
L15:
	;
	if v18 != int32(36) {
		v47 = v18
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v42 = F_find_param_referent(m, l0, l3, v16+int32(252), v16+int32(248))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if l1 <= int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v25 == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 < l1 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+l1<<(uint(int32(2))%32)-int32(4))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v860 = v37
	goto L1
L22:
	;
	return int32(0)
L23:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = v46
	goto L14
L25:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 == int32(2249) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v860 = v56 + v58<<(uint(int32(3))%32) + l1*int32(100) - int32(68)
	goto L1
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = v70
	goto L30
L29:
	;
	v72 = int32(0)
	goto L30
L30:
	;
	if v72 <= v68 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v68<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+252)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v80 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88+l0))))
	if v87 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = v85
	v88 = int32(8)
	goto L32
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	if v83 != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v87 = v80
	v88 = int32(40)
	goto L32
L36:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	switch v294 - int32(1) {
	case 0:
		goto L106
	case 1:
		goto L105
	default:
		v667 = l0
		goto L103
	case 5:
		goto L104
	}
L37:
	;
	switch v87 + int32(3) {
	case 0:
		goto L44
	case 1:
		goto L46
	case 2:
		goto L45
	default:
		goto L43
	}
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v93 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 < v87 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v87<<(uint(int32(2))%32)-int32(4))))
	if v90 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v105 = F_get_rte_attribute_name(m, v104, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v860 = v105
	goto L1
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L22
	} else {
		goto L100
	}
L44:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v78)+64))
	if v233 == int32(0) {
		goto L43
	} else {
		goto L84
	}
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v168 == int32(0) {
		goto L43
	} else {
		goto L65
	}
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v78)+56))
	if v110 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	if v110 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v150 == int32(0) {
		goto L9
	} else {
		goto L61
	}
L49:
	;
	goto L48
L50:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v116 <= int32(0) {
		v150 = int32(0)
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v150 = int32(0)
	goto L49
L53:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v122 = v116
	goto L56
L55:
	;
	v122 = v119
	goto L56
L56:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v127 = int32(0)
	goto L57
L57:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123+v127<<(uint(int32(2))%32))))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+8)))
	if v136 == v90&int32(_a_F_get_name_for_var_field_0) {
		v150 = v135
		goto L49
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v139 = v127 + int32(1)
	if v139 != v122 {
		v127 = v139
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
	v156 = v16 + int32(168)
	F_push_child_plan(m, v78, v154, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v160 = F_get_name_for_var_field(m, v159, l1, l2, l3)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v163 = F_list_delete_first(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	base.MemoryCopy(m, v78, v156, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v163
	v860 = v160
	goto L1
L65:
	;
	if v168 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v208 == int32(0) {
		goto L8
	} else {
		goto L79
	}
L67:
	;
	goto L66
L68:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v174 <= int32(0) {
		v208 = int32(0)
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v208 = int32(0)
	goto L67
L71:
	;
	v177 = int32(0)
	if v177 < v174 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v180 = v174
	goto L74
L73:
	;
	v180 = v177
	goto L74
L74:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v185 = int32(0)
	goto L75
L75:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v181+v185<<(uint(int32(2))%32))))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+8)))
	if v194 == v90&int32(_a_F_get_name_for_var_field_0) {
		v208 = v193
		goto L67
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v197 = v185 + int32(1)
	if v197 != v180 {
		v185 = v197
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v214 = v16 + int32(168)
	base.MemoryCopy(m, v214, v78, int32(80))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v219 = F_lcons(m, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L22
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v219
	F_set_deparse_plan(m, v78, v212)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L22
	} else {
		goto L81
	}
L81:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v225 = F_get_name_for_var_field(m, v224, l1, l2, l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L22
	} else {
		goto L82
	}
L82:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v228 = F_list_delete_first(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	base.MemoryCopy(m, v78, v214, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v228
	v860 = v225
	goto L1
L84:
	;
	if v233 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	if v273 == int32(0) {
		goto L7
	} else {
		goto L98
	}
L86:
	;
	goto L85
L87:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v239 <= int32(0) {
		v273 = int32(0)
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v273 = int32(0)
	goto L86
L90:
	;
	v242 = int32(0)
	if v242 < v239 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v245 = v239
	goto L93
L92:
	;
	v245 = v242
	goto L93
L93:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v250 = int32(0)
	goto L94
L94:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v246+v250<<(uint(int32(2))%32))))
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+8)))
	if v259 == v90&int32(_a_F_get_name_for_var_field_0) {
		v273 = v258
		goto L86
	} else {
		goto L96
	}
L95:
	;
	goto L89
L96:
	;
	v262 = v250 + int32(1)
	if v262 != v245 {
		v250 = v262
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v278 = F_get_name_for_var_field(m, v277, l1, l2, l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	v860 = v278
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v87
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_1), v16)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L22
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_3), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L22
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v681 = F_get_expr_result_tupdesc(m, v667, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L22
	} else {
		goto L220
	}
L104:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v104)+88))
	v447 = v446 + v68
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(v448) <= base.Ui32(v447) {
		goto L155
	} else {
		goto L156
	}
L105:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v104)+52))
	if v431 == int32(0) {
		goto L4
	} else {
		goto L152
	}
L106:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	if v297 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+76))
	if v298 != 0 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	goto L109
L109:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	if v362 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L110:
	;
	if v336 == int32(0) {
		goto L6
	} else {
		goto L123
	}
L111:
	;
	goto L110
L112:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v302 <= int32(0) {
		v336 = int32(0)
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v336 = int32(0)
	goto L111
L115:
	;
	v305 = int32(0)
	if v305 < v302 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v308 = v302
	goto L118
L117:
	;
	v308 = v305
	goto L118
L118:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v313 = int32(0)
	goto L119
L119:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v309+v313<<(uint(int32(2))%32))))
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v321)+8)))
	if v322 == v90&int32(_a_F_get_name_for_var_field_0) {
		v336 = v321
		goto L111
	} else {
		goto L121
	}
L120:
	;
	goto L114
L121:
	;
	v325 = v313 + int32(1)
	if v325 != v308 {
		v313 = v325
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+26)))
	if v340 == int32(1) {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v344 != int32(6) {
		v667 = v343
		goto L103
	} else {
		goto L125
	}
L125:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v348 = F_list_copy_tail(m, v347, v68)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	v351 = v16 + int32(168)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	F_set_deparse_for_query(m, v351, v352, v348)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	v355 = F_lcons(m, v351, v348)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v355
	v359 = F_get_name_for_var_field(m, v343, l1, int32(0), l3)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v347
	v860 = v359
	goto L1
L130:
	;
	v366 = F_palloc(m, int32(32))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L22
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v375 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l1
	v373 = F_pg_snprintf(m, v366, int32(32), int32(_a_F_get_name_for_var_field_5), v16+int32(112))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L22
	} else {
		goto L134
	}
L134:
	;
	v860 = v366
	goto L1
L135:
	;
	if v413 == int32(0) {
		goto L5
	} else {
		goto L148
	}
L136:
	;
	goto L135
L137:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v379 <= int32(0) {
		v413 = int32(0)
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v413 = int32(0)
	goto L136
L140:
	;
	v382 = int32(0)
	if v382 < v379 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v385 = v379
	goto L143
L142:
	;
	v385 = v382
	goto L143
L143:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v390 = int32(0)
	goto L144
L144:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v386+v390<<(uint(int32(2))%32))))
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+8)))
	if v399 == v90&int32(_a_F_get_name_for_var_field_0) {
		v413 = v398
		goto L136
	} else {
		goto L146
	}
L145:
	;
	goto L139
L146:
	;
	v402 = v390 + int32(1)
	if v402 != v385 {
		v390 = v402
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v419 = v16 + int32(168)
	F_push_child_plan(m, v78, v417, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L22
	} else {
		goto L149
	}
L149:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v423 = F_get_name_for_var_field(m, v422, l1, l2, l3)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L22
	} else {
		goto L150
	}
L150:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v426 = F_list_delete_first(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L151
	}
L151:
	;
	base.MemoryCopy(m, v78, v419, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v426
	v860 = v423
	goto L1
L152:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v434+v90<<(uint(int32(2))%32)-int32(4))))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	if v441 != int32(6) {
		v667 = v440
		goto L103
	} else {
		goto L153
	}
L153:
	;
	v444 = F_get_name_for_var_field(m, v440, l1, v68, l3)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L22
	} else {
		goto L154
	}
L154:
	;
	v860 = v444
	goto L1
L155:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	if v598 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L156:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v74+v447<<(uint(int32(2))%32))))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+16))
	if v454 == int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v457 <= int32(0) {
		goto L155
	} else {
		goto L158
	}
L158:
	;
	v460 = int32(0)
	if v460 < v457 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v464 = v457
	goto L161
L160:
	;
	v464 = v460
	goto L161
L161:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v104)+84))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v454)+12))
	v467 = v460
	goto L162
L162:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v466+v467<<(uint(int32(2))%32))))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+4))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if base.B2i32(v487 == int32(0))|base.B2i32(v487 != v490) != 0 {
		v508 = v487
		v509 = v490
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v517 == int32(1) {
		goto L175
	} else {
		goto L176
	}
L164:
	;
	if v508-v509 != 0 {
		goto L171
	} else {
		goto L172
	}
L165:
	;
	goto L164
L166:
	;
	v493 = v484
	v494 = v465
	goto L167
L167:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
	if v498 == int32(0) {
		v508 = v498
		v509 = v497
		goto L165
	} else {
		goto L169
	}
L168:
	;
	v508 = v498
	v509 = v497
	goto L165
L169:
	;
	v501 = int32(1)
	if v498 == v497 {
		v493 = v493 + v501
		v494 = v494 + v501
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v512 = v467 + int32(1)
	if v464 != v512 {
		v467 = v512
		goto L162
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	goto L163
L174:
	;
	goto L155
L175:
	;
	v520 = int32(76)
	goto L177
L176:
	;
	v520 = int32(96)
	goto L177
L177:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514+v520)))
	if v522 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v560 == int32(0) {
		goto L3
	} else {
		goto L191
	}
L179:
	;
	goto L178
L180:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v526 <= int32(0) {
		v560 = int32(0)
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v560 = int32(0)
	goto L179
L183:
	;
	v529 = int32(0)
	if v529 < v526 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v532 = v526
	goto L186
L185:
	;
	v532 = v529
	goto L186
L186:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v537 = int32(0)
	goto L187
L187:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v533+v537<<(uint(int32(2))%32))))
	v546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545)+8)))
	if v546 == v90&int32(_a_F_get_name_for_var_field_0) {
		v560 = v545
		goto L179
	} else {
		goto L189
	}
L188:
	;
	goto L182
L189:
	;
	v549 = v537 + int32(1)
	if v549 != v532 {
		v537 = v549
		goto L187
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+26)))
	if v564 == int32(1) {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v568 != int32(6) {
		v667 = v567
		goto L103
	} else {
		goto L193
	}
L193:
	;
	v572 = v16 + int32(168)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v574 = F_list_copy_tail(m, v573, v447)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L22
	} else {
		goto L194
	}
L194:
	;
	F_set_deparse_for_query(m, v572, v514, v574)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L22
	} else {
		goto L195
	}
L195:
	;
	v578 = F_lcons(m, v572, v574)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L22
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v578
	v582 = F_get_name_for_var_field(m, v567, l1, int32(0), l3)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L22
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v573
	v860 = v582
	goto L1
L198:
	;
	v602 = F_palloc(m, int32(32))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L22
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	if v611 != 0 {
		goto L205
	} else {
		goto L206
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l1
	v609 = F_pg_snprintf(m, v602, int32(32), int32(_a_F_get_name_for_var_field_5), v16-int32(-64))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L22
	} else {
		goto L202
	}
L202:
	;
	v860 = v602
	goto L1
L203:
	;
	if v649 == int32(0) {
		goto L2
	} else {
		goto L216
	}
L204:
	;
	goto L203
L205:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v615 <= int32(0) {
		v649 = int32(0)
		goto L204
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v649 = int32(0)
	goto L204
L208:
	;
	v618 = int32(0)
	if v618 < v615 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v621 = v615
	goto L211
L210:
	;
	v621 = v618
	goto L211
L211:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v611)+12))
	v626 = int32(0)
	goto L212
L212:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v622+v626<<(uint(int32(2))%32))))
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634)+8)))
	if v635 == v90&int32(_a_F_get_name_for_var_field_0) {
		v649 = v634
		goto L204
	} else {
		goto L214
	}
L213:
	;
	goto L207
L214:
	;
	v638 = v626 + int32(1)
	if v638 != v621 {
		v626 = v638
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v655 = v16 + int32(168)
	F_push_child_plan(m, v78, v653, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L22
	} else {
		goto L217
	}
L217:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	v659 = F_get_name_for_var_field(m, v658, l1, l2, l3)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L22
	} else {
		goto L218
	}
L218:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v662 = F_list_delete_first(m, v661)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L22
	} else {
		goto L219
	}
L219:
	;
	base.MemoryCopy(m, v78, v655, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v662
	v860 = v659
	goto L1
L220:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v860 = v681 + v683<<(uint(int32(3))%32) + l1*int32(100) - int32(68)
	goto L1
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+44)) = v706
	F_set_deparse_plan(m, v696, v693)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L22
	} else {
		goto L222
	}
L222:
	;
	v712 = F_get_name_for_var_field(m, v42, l1, int32(0), l3)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L22
	} else {
		goto L223
	}
L223:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v696)+44))
	F_list_free(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L22
	} else {
		goto L224
	}
L224:
	;
	base.MemoryCopy(m, v696, v695, int32(80))
	v860 = v712
	goto L1
L225:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+164)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v723
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_6), v16+int32(160))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L22
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_7), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L22
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_8), v16+int32(16))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L22
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_9), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L22
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_10), v16+int32(32))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L22
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_11), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L22
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_12), v16+int32(48))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L22
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_13), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L22
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+148)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v786
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_14), v16+int32(144))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L22
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_15), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L22
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_16), v16+int32(128))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L22
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_17), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L22
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_18), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L22
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_19), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L22
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L246:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v832
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_20), v16+int32(96))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L22
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_21), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L22
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v90
	F_errmsg_internal(m, int32(_a_F_get_name_for_var_field_16), v16+int32(80))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L22
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_get_name_for_var_field_2), int32(_a_F_get_name_for_var_field_22), int32(_a_F_get_name_for_var_field_4))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L22
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
