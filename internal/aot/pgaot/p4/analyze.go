package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_analyzeCTE(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
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
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
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
	var v345 int32
	_ = v345
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
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
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
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
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
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
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
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(224)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	if v18|v17 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L2:
	;
	if v414 != 0 {
		goto L125
	} else {
		goto L126
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L13
	} else {
		goto L116
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L13
	} else {
		goto L106
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L101
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L98
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L95
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L13
	} else {
		goto L90
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L13
	} else {
		goto L85
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v21 = F_transformExpr(m, l0, v19, int32(44))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v103 = F_parse_sub_analyze(m, v101, l0, l1, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L28
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v26 = F_transformExpr(m, l0, v24, int32(44))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v26
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+216)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v15)+196)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v26
	v38 = F_list_make2_impl(m, v15+int32(196), v15+int32(192))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v42 = F_select_common_type(m, l0, v38, int32(516275), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v47 = F_coerce_to_common_type(m, l0, v45, v42, int32(502965))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v53 = F_coerce_to_common_type(m, l0, v50, v51, int32(496432))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+188)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+184)) = v53
	v65 = F_list_make2_impl(m, v15+int32(188), v15+int32(184))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v68 = F_select_common_typmod(m, v65, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+204)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+200)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v15)+180)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v73
	v81 = F_list_make2_impl(m, v15+int32(180), v15+int32(176))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v84 = F_select_common_collation(m, l0, v81, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v89 = F_lookup_type_cache(m, v87, int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+52))
	if v91 == int32(0) {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v94 = F_get_negator(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	if v94 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v94
	goto L12
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v106 != int32(67) {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)+28))
	if v109 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v110 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v113 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+24)) = uint8(v114)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v117 == v114 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v123 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v131 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v126 = int32(76)
	goto L40
L39:
	;
	v126 = int32(96)
	goto L40
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120+v126)))
	F_analyzeCTETargetList(m, l0, l1, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v133 = v132
	goto L44
L43:
	;
	v133 = v114
	goto L44
L44:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v134 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v136 = v135
	goto L47
L46:
	;
	v136 = v3
	goto L47
L47:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v137 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = v138
	goto L50
L49:
	;
	v139 = v3
	goto L50
L50:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v143 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v146 = int32(76)
	goto L53
L52:
	;
	v146 = int32(96)
	goto L53
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140+v146)))
	if v148 == int32(0) {
		v414 = v133
		v415 = v139
		v416 = v136
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v151 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v152 <= v151 {
		v414 = v133
		v415 = v139
		v416 = v136
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v157 = v151
	v160 = v133
	v161 = v139
	v162 = v136
	v165 = v3
	goto L56
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v157<<(uint(int32(2))%32))))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+26)))
	if v172 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v414 = v240
	v415 = v241
	v416 = v242
	goto L2
L58:
	;
	if v160 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v240 = v160
	v241 = v161
	v242 = v162
	v244 = v165
	goto L60
L60:
	;
	v246 = v157 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v246 < v247 {
		v157 = v246
		v160 = v240
		v161 = v241
		v162 = v242
		v165 = v244
		goto L56
	} else {
		goto L84
	}
L61:
	;
	v193 = v165 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v195 = F_exprType(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L69
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L13
	} else {
		goto L66
	}
L63:
	;
	if v162 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if v161 != 0 {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	F_errmsg_internal(m, int32(511382), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(475120), int32(379), int32(514148))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
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
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v195 != v197 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v199 = F_exprTypmod(m, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v199 != v201 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v203 = F_exprCollation(m, v194)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v203 != v205 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v208 = v161 + int32(4)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if base.Ui32(v208) < base.Ui32(v211+v212<<(uint(int32(2))%32)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v217 = v208
	goto L77
L76:
	;
	v217 = int32(0)
	goto L77
L77:
	;
	v219 = v162 + int32(4)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if base.Ui32(v219) < base.Ui32(v222+v223<<(uint(int32(2))%32)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v228 = v219
	goto L80
L79:
	;
	v228 = int32(0)
	goto L80
L80:
	;
	v230 = v160 + int32(4)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if base.Ui32(v230) < base.Ui32(v233+v234<<(uint(int32(2))%32)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v239 = v230
	goto L83
L82:
	;
	v239 = int32(0)
	goto L83
L83:
	;
	v240 = v239
	v241 = v217
	v242 = v228
	v244 = v193
	goto L60
L84:
	;
	goto L57
L85:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v257 = F_format_type_be(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v257
	F_errmsg(m, int32(179173), v15)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(475120), int32(304), int32(514148))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v276 = F_format_type_be(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v276
	F_errmsg(m, int32(179119), v15+int32(16))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(475120), int32(310), int32(514148))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L13
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errmsg_internal(m, int32(510580), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(475120), int32(324), int32(514148))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errmsg_internal(m, int32(510543), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(475120), int32(326), int32(514148))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(291424), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(475120), int32(337), int32(514148))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v344 = F_format_type_with_typemod(m, v342, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	v346 = F_exprType(m, v194)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v348 = F_exprTypmod(m, v194)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	v350 = F_format_type_with_typemod(m, v346, v348)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v341
	F_errmsg(m, int32(289400), v15+int32(160))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	F_errhint(m, int32(592520), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	v365 = F_exprLocation(m, v194)
	mBase = m.M
	F_parser_errposition(m, l0, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(475120), int32(392), int32(514148))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L117
	}
L117:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v382 = F_get_collation_name(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	v384 = F_exprCollation(m, v194)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L119
	}
L119:
	;
	v386 = F_get_collation_name(m, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v380
	F_errmsg(m, int32(289485), v15+int32(144))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	F_errhint(m, int32(580334), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	v401 = F_exprLocation(m, v194)
	mBase = m.M
	F_parser_errposition(m, l0, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(475120), int32(401), int32(514148))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L13
	} else {
		goto L129
	}
L126:
	;
	if v416 != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	if v415 == int32(0) {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	F_errmsg_internal(m, int32(511382), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(475120), int32(407), int32(514148))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L13
	} else {
		goto L271
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L13
	} else {
		goto L266
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L13
	} else {
		goto L261
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L13
	} else {
		goto L256
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L13
	} else {
		goto L251
	}
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L13
	} else {
		goto L246
	}
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L13
	} else {
		goto L241
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L13
	} else {
		goto L236
	}
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L13
	} else {
		goto L231
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L13
	} else {
		goto L226
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L13
	} else {
		goto L222
	}
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L13
	} else {
		goto L218
	}
L144:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L13
	} else {
		goto L213
	}
L145:
	;
	m.G0 = v15 + int32(224)
	return
L146:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v451 == int32(0) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+144))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	if v457 != int32(63) {
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v455)+16))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	if v461 != int32(63) {
		goto L142
	} else {
		goto L149
	}
L149:
	;
	if v17 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v464 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	if v18 == int32(0) {
		goto L145
	} else {
		goto L167
	}
L153:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v516 = F_makeString(m, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L13
	} else {
		goto L164
	}
L154:
	;
	v467 = int32(0)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v468 <= v467 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v474 = v467
	v477 = int32(0)
	goto L156
L156:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485+v477<<(uint(int32(2))%32))))
	v490 = F_list_member(m, v484, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L13
	} else {
		goto L158
	}
L157:
	;
	goto L153
L158:
	;
	if v490 == int32(0) {
		goto L141
	} else {
		goto L159
	}
L159:
	;
	v494 = F_list_member(m, v474, v489)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	if v494 != 0 {
		goto L140
	} else {
		goto L161
	}
L161:
	;
	v496 = F_lappend(m, v474, v489)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	v499 = v477 + int32(1)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v499 < v500 {
		v474 = v496
		v477 = v499
		goto L156
	} else {
		goto L163
	}
L163:
	;
	goto L157
L164:
	;
	v518 = F_list_member(m, v514, v516)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	if v518 != 0 {
		goto L139
	} else {
		goto L166
	}
L166:
	;
	goto L152
L167:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v534 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v586 = F_makeString(m, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L13
	} else {
		goto L179
	}
L169:
	;
	v537 = int32(0)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v538 <= v537 {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v544 = v537
	v547 = int32(0)
	goto L171
L171:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v547<<(uint(int32(2))%32))))
	v560 = F_list_member(m, v554, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L13
	} else {
		goto L173
	}
L172:
	;
	goto L168
L173:
	;
	if v560 == int32(0) {
		goto L138
	} else {
		goto L174
	}
L174:
	;
	v564 = F_list_member(m, v544, v559)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	if v564 != 0 {
		goto L137
	} else {
		goto L176
	}
L176:
	;
	v566 = F_lappend(m, v544, v559)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L13
	} else {
		goto L177
	}
L177:
	;
	v569 = v547 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v569 < v570 {
		v544 = v566
		v547 = v569
		goto L171
	} else {
		goto L178
	}
L178:
	;
	goto L172
L179:
	;
	v588 = F_list_member(m, v584, v586)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	if v588 != 0 {
		goto L136
	} else {
		goto L181
	}
L181:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v592 = F_makeString(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	v594 = F_list_member(m, v590, v592)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L13
	} else {
		goto L183
	}
L183:
	;
	if v594 != 0 {
		goto L135
	} else {
		goto L184
	}
L184:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v601 == int32(0) {
		v620 = v600
		v621 = v601
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v621-v620 == int32(0) {
		goto L134
	} else {
		goto L193
	}
L186:
	;
	goto L185
L187:
	;
	if v600 != v601 {
		v620 = v600
		v621 = v601
		goto L186
	} else {
		goto L188
	}
L188:
	;
	v605 = v596
	v606 = v597
	goto L189
L189:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)))
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605)+1)))
	if v610 == int32(0) {
		v620 = v609
		v621 = v610
		goto L186
	} else {
		goto L191
	}
L190:
	;
	v620 = v609
	v621 = v610
	goto L186
L191:
	;
	v613 = int32(1)
	if v609 == v610 {
		v605 = v605 + v613
		v606 = v606 + v613
		goto L189
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	if v17 == int32(0) {
		goto L145
	} else {
		goto L194
	}
L194:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	if v631 == int32(0) {
		v650 = v630
		v651 = v631
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v651-v650 == int32(0) {
		goto L133
	} else {
		goto L203
	}
L196:
	;
	goto L195
L197:
	;
	if v630 != v631 {
		v650 = v630
		v651 = v631
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v635 = v627
	v636 = v596
	goto L199
L199:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+1)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635)+1)))
	if v640 == int32(0) {
		v650 = v639
		v651 = v640
		goto L196
	} else {
		goto L201
	}
L200:
	;
	v650 = v639
	v651 = v640
	goto L196
L201:
	;
	v643 = int32(1)
	if v639 == v640 {
		v635 = v635 + v643
		v636 = v636 + v643
		goto L199
	} else {
		goto L202
	}
L202:
	;
	goto L200
L203:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	if v658 == int32(0) {
		v677 = v657
		v678 = v658
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v678-v677 == int32(0) {
		goto L132
	} else {
		goto L212
	}
L205:
	;
	goto L204
L206:
	;
	if v657 != v658 {
		v677 = v657
		v678 = v658
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v662 = v627
	v663 = v597
	goto L208
L208:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+1)))
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	if v667 == int32(0) {
		v677 = v666
		v678 = v667
		goto L205
	} else {
		goto L210
	}
L209:
	;
	v677 = v666
	v678 = v667
	goto L205
L210:
	;
	v670 = int32(1)
	if v666 == v667 {
		v662 = v662 + v670
		v663 = v663 + v670
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	goto L145
L213:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L13
	} else {
		goto L214
	}
L214:
	;
	F_errmsg(m, int32(326766), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L13
	} else {
		goto L215
	}
L215:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v708)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L13
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(475120), int32(422), int32(514148))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L13
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L13
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(498102), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L13
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(475120), int32(452), int32(514148))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L13
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L13
	} else {
		goto L223
	}
L223:
	;
	F_errmsg(m, int32(498026), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L13
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(475120), int32(457), int32(514148))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L13
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L13
	} else {
		goto L227
	}
L227:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v755
	F_errmsg(m, int32(71471), v15+int32(128))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L13
	} else {
		goto L228
	}
L228:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L13
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(475120), int32(474), int32(514148))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L13
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L13
	} else {
		goto L232
	}
L232:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v777
	F_errmsg(m, int32(395745), v15+int32(112))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L13
	} else {
		goto L233
	}
L233:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L13
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(475120), int32(481), int32(514148))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L13
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L13
	} else {
		goto L237
	}
L237:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v799
	F_errmsg(m, int32(71702), v15+int32(32))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L13
	} else {
		goto L238
	}
L238:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L13
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(475120), int32(490), int32(514148))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L13
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L13
	} else {
		goto L242
	}
L242:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v821
	F_errmsg(m, int32(71520), v15+int32(96))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L13
	} else {
		goto L243
	}
L243:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L13
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(475120), int32(507), int32(514148))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L13
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L13
	} else {
		goto L247
	}
L247:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v843
	F_errmsg(m, int32(395789), v15+int32(80))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L13
	} else {
		goto L248
	}
L248:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v850)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L13
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(475120), int32(514), int32(514148))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L13
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L13
	} else {
		goto L252
	}
L252:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v865
	F_errmsg(m, int32(71568), v15+int32(48))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L13
	} else {
		goto L253
	}
L253:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v872)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L13
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(475120), int32(523), int32(514148))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L13
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L13
	} else {
		goto L257
	}
L257:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v887
	F_errmsg(m, int32(71635), v15-int32(-64))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L13
	} else {
		goto L258
	}
L258:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L13
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(475120), int32(530), int32(514148))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L13
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L13
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(358516), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L13
	} else {
		goto L263
	}
L263:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L13
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(475120), int32(537), int32(514148))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L13
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L13
	} else {
		goto L267
	}
L267:
	;
	F_errmsg(m, int32(358448), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L13
	} else {
		goto L268
	}
L268:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v932)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L13
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(475120), int32(547), int32(514148))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L13
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L13
	} else {
		goto L272
	}
L272:
	;
	F_errmsg(m, int32(358579), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L13
	} else {
		goto L273
	}
L273:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L13
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(475120), int32(554), int32(514148))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L13
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_serializeAnalyzeShutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v8 != 0 {
				F_pfree(m, v8)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v13 != 0 {
						F_MemoryContextDelete(m, v13)
						mBase = m.M
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v13 != 0 {
					F_MemoryContextDelete(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					return
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v8 != 0 {
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v13 != 0 {
					F_MemoryContextDelete(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					return
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v13 != 0 {
				F_MemoryContextDelete(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				return
			}
		}
	}
}
