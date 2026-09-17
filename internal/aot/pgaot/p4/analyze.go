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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
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
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
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
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v519 int32
	_ = v519
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
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
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
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
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
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
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
		goto L143
	} else {
		goto L144
	}
L2:
	;
	if v416|v419|v421 == int32(0) {
		goto L1
	} else {
		goto L126
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L117
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L107
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L102
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L99
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L13
	} else {
		goto L96
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L13
	} else {
		goto L91
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L86
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
	v42 = F_select_common_type(m, l0, v38, int32(_a_F_analyzeCTE_0), int32(0))
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
	v47 = F_coerce_to_common_type(m, l0, v45, v42, int32(_a_F_analyzeCTE_1))
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
	v53 = F_coerce_to_common_type(m, l0, v50, v51, int32(_a_F_analyzeCTE_2))
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
		v416 = v133
		v419 = v136
		v421 = v139
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v151 <= int32(0) {
		v416 = v133
		v419 = v136
		v421 = v139
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v159 = v133
	v160 = int32(0)
	v162 = v136
	v164 = v139
	v166 = v3
	goto L56
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v160<<(uint(int32(2))%32))))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+26)))
	if v172 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v416 = v243
	v419 = v245
	v421 = v246
	goto L2
L58:
	;
	v175 = int32(0)
	if v164 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v243 = v159
	v245 = v162
	v246 = v164
	v247 = v166
	goto L60
L60:
	;
	v249 = v160 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v249 < v250 {
		v159 = v243
		v160 = v249
		v162 = v245
		v164 = v246
		v166 = v247
		goto L56
	} else {
		goto L85
	}
L61:
	;
	v181 = base.B2i32(v159 == v175) | base.B2i32(v162 == v175)
	goto L63
L62:
	;
	v181 = int32(1)
	goto L63
L63:
	;
	if v181 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L13
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v196 = v166 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v198 = F_exprType(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L70
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_analyzeCTE_3), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(379), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v198 != v200 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v202 = F_exprTypmod(m, v197)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v202 != v204 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v206 = F_exprCollation(m, v197)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v206 != v208 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v211 = v164 + int32(4)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if base.Ui32(v211) < base.Ui32(v214+v215<<(uint(int32(2))%32)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v220 = v211
	goto L78
L77:
	;
	v220 = int32(0)
	goto L78
L78:
	;
	v222 = v162 + int32(4)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if base.Ui32(v222) < base.Ui32(v225+v226<<(uint(int32(2))%32)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v231 = v222
	goto L81
L80:
	;
	v231 = int32(0)
	goto L81
L81:
	;
	v233 = v159 + int32(4)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if base.Ui32(v233) < base.Ui32(v236+v237<<(uint(int32(2))%32)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v242 = v233
	goto L84
L83:
	;
	v242 = int32(0)
	goto L84
L84:
	;
	v243 = v242
	v245 = v231
	v246 = v220
	v247 = v196
	goto L60
L85:
	;
	goto L57
L86:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L87
	}
L87:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v260 = F_format_type_be(m, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v260
	F_errmsg(m, int32(_a_F_analyzeCTE_6), v15)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(304), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v279 = F_format_type_be(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v279
	F_errmsg(m, int32(_a_F_analyzeCTE_7), v15+int32(16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L13
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(310), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errmsg_internal(m, int32(_a_F_analyzeCTE_8), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(324), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errmsg_internal(m, int32(_a_F_analyzeCTE_9), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(326), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_10), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(337), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v347 = F_format_type_with_typemod(m, v345, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v349 = F_exprType(m, v197)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L110
	}
L110:
	;
	v351 = F_exprTypmod(m, v197)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	v353 = F_format_type_with_typemod(m, v349, v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+172)) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v344
	F_errmsg(m, int32(_a_F_analyzeCTE_11), v15+int32(160))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	F_errhint(m, int32(_a_F_analyzeCTE_12), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	v368 = F_exprLocation(m, v197)
	mBase = m.M
	F_parser_errposition(m, l0, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L13
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(392), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L13
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v385 = F_get_collation_name(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L13
	} else {
		goto L119
	}
L119:
	;
	v387 = F_exprCollation(m, v197)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	v389 = F_get_collation_name(m, v387)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v15)+148)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v383
	F_errmsg(m, int32(_a_F_analyzeCTE_13), v15+int32(144))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	F_errhint(m, int32(_a_F_analyzeCTE_14), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	v404 = F_exprLocation(m, v197)
	mBase = m.M
	F_parser_errposition(m, l0, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(401), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	F_errmsg_internal(m, int32(_a_F_analyzeCTE_3), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(407), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L13
	} else {
		goto L266
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L13
	} else {
		goto L261
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L13
	} else {
		goto L256
	}
L133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L13
	} else {
		goto L251
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L13
	} else {
		goto L246
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L13
	} else {
		goto L241
	}
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L13
	} else {
		goto L236
	}
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L13
	} else {
		goto L231
	}
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L13
	} else {
		goto L226
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L13
	} else {
		goto L221
	}
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L13
	} else {
		goto L217
	}
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L13
	} else {
		goto L213
	}
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L13
	} else {
		goto L208
	}
L143:
	;
	m.G0 = v15 + int32(224)
	return
L144:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v456 == int32(0) {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+144))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	if v462 != int32(63) {
		goto L141
	} else {
		goto L146
	}
L146:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v460)+16))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	if v466 != int32(63) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	if v17 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v469 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	if v18 == int32(0) {
		goto L143
	} else {
		goto L165
	}
L151:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v521 = F_makeString(m, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L13
	} else {
		goto L162
	}
L152:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v472 <= int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v475 = int32(0)
	v481 = v475
	v482 = v475
	goto L154
L154:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v469)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v490+v482<<(uint(int32(2))%32))))
	v495 = F_list_member(m, v489, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L13
	} else {
		goto L156
	}
L155:
	;
	goto L151
L156:
	;
	if v495 == int32(0) {
		goto L139
	} else {
		goto L157
	}
L157:
	;
	v499 = F_list_member(m, v481, v494)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	if v499 != 0 {
		goto L138
	} else {
		goto L159
	}
L159:
	;
	v501 = F_lappend(m, v481, v494)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	v504 = v482 + int32(1)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v504 < v505 {
		v481 = v501
		v482 = v504
		goto L154
	} else {
		goto L161
	}
L161:
	;
	goto L155
L162:
	;
	v523 = F_list_member(m, v519, v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	if v523 != 0 {
		goto L137
	} else {
		goto L164
	}
L164:
	;
	goto L150
L165:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v539 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v591 = F_makeString(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L13
	} else {
		goto L177
	}
L167:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v542 <= int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v545 = int32(0)
	v551 = v545
	v552 = v545
	goto L169
L169:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560+v552<<(uint(int32(2))%32))))
	v565 = F_list_member(m, v559, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L13
	} else {
		goto L171
	}
L170:
	;
	goto L166
L171:
	;
	if v565 == int32(0) {
		goto L136
	} else {
		goto L172
	}
L172:
	;
	v569 = F_list_member(m, v551, v564)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L13
	} else {
		goto L173
	}
L173:
	;
	if v569 != 0 {
		goto L135
	} else {
		goto L174
	}
L174:
	;
	v571 = F_lappend(m, v551, v564)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	v574 = v552 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v574 < v575 {
		v551 = v571
		v552 = v574
		goto L169
	} else {
		goto L176
	}
L176:
	;
	goto L170
L177:
	;
	v593 = F_list_member(m, v589, v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L13
	} else {
		goto L178
	}
L178:
	;
	if v593 != 0 {
		goto L134
	} else {
		goto L179
	}
L179:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v597 = F_makeString(m, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	v599 = F_list_member(m, v595, v597)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L13
	} else {
		goto L181
	}
L181:
	;
	if v599 != 0 {
		goto L133
	} else {
		goto L182
	}
L182:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	if base.B2i32(v605 == int32(0))|base.B2i32(v605 != v608) != 0 {
		v626 = v605
		v627 = v608
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v626-v627 == int32(0) {
		goto L132
	} else {
		goto L190
	}
L184:
	;
	goto L183
L185:
	;
	v611 = v601
	v612 = v602
	goto L186
L186:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+1)))
	if v616 == int32(0) {
		v626 = v616
		v627 = v615
		goto L184
	} else {
		goto L188
	}
L187:
	;
	v626 = v616
	v627 = v615
	goto L184
L188:
	;
	v619 = int32(1)
	if v616 == v615 {
		v611 = v611 + v619
		v612 = v612 + v619
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	if v17 == int32(0) {
		goto L143
	} else {
		goto L191
	}
L191:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	if base.B2i32(v636 == int32(0))|base.B2i32(v636 != v639) != 0 {
		v657 = v636
		v658 = v639
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v657-v658 == int32(0) {
		goto L131
	} else {
		goto L199
	}
L193:
	;
	goto L192
L194:
	;
	v642 = v633
	v643 = v601
	goto L195
L195:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+1)))
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+1)))
	if v647 == int32(0) {
		v657 = v647
		v658 = v646
		goto L193
	} else {
		goto L197
	}
L196:
	;
	v657 = v647
	v658 = v646
	goto L193
L197:
	;
	v650 = int32(1)
	if v647 == v646 {
		v642 = v642 + v650
		v643 = v643 + v650
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	if base.B2i32(v664 == int32(0))|base.B2i32(v664 != v667) != 0 {
		v685 = v664
		v686 = v667
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v685-v686 == int32(0) {
		goto L130
	} else {
		goto L207
	}
L201:
	;
	goto L200
L202:
	;
	v670 = v633
	v671 = v602
	goto L203
L203:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+1)))
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670)+1)))
	if v675 == int32(0) {
		v685 = v675
		v686 = v674
		goto L201
	} else {
		goto L205
	}
L204:
	;
	v685 = v675
	v686 = v674
	goto L201
L205:
	;
	v678 = int32(1)
	if v675 == v674 {
		v670 = v670 + v678
		v671 = v671 + v678
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	goto L143
L208:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L13
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_15), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L13
	} else {
		goto L210
	}
L210:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L13
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(422), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L13
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
	v730 = m.ExcPending
	if v730 != 0 {
		goto L13
	} else {
		goto L214
	}
L214:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_16), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L13
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(452), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L13
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L13
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_17), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(457), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L13
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L13
	} else {
		goto L222
	}
L222:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v763
	F_errmsg(m, int32(_a_F_analyzeCTE_18), v15+int32(128))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L13
	} else {
		goto L223
	}
L223:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L13
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(474), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L13
	} else {
		goto L227
	}
L227:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v785
	F_errmsg(m, int32(_a_F_analyzeCTE_19), v15+int32(112))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L13
	} else {
		goto L228
	}
L228:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L13
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(481), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L13
	} else {
		goto L232
	}
L232:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v807
	F_errmsg(m, int32(_a_F_analyzeCTE_20), v15+int32(32))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L13
	} else {
		goto L233
	}
L233:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L13
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(490), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
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
	v828 = m.ExcPending
	if v828 != 0 {
		goto L13
	} else {
		goto L237
	}
L237:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v829
	F_errmsg(m, int32(_a_F_analyzeCTE_21), v15+int32(96))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L13
	} else {
		goto L238
	}
L238:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v836)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L13
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(507), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
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
	F_errcode(m, int32(16806020))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L13
	} else {
		goto L242
	}
L242:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v851
	F_errmsg(m, int32(_a_F_analyzeCTE_22), v15+int32(80))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L13
	} else {
		goto L243
	}
L243:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v858)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L13
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(514), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L13
	} else {
		goto L247
	}
L247:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v873
	F_errmsg(m, int32(_a_F_analyzeCTE_23), v15+int32(48))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L13
	} else {
		goto L248
	}
L248:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L13
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(523), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
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
	v894 = m.ExcPending
	if v894 != 0 {
		goto L13
	} else {
		goto L252
	}
L252:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v895
	F_errmsg(m, int32(_a_F_analyzeCTE_24), v15-int32(-64))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L13
	} else {
		goto L253
	}
L253:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L13
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(530), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
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
	v916 = m.ExcPending
	if v916 != 0 {
		goto L13
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_25), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L13
	} else {
		goto L258
	}
L258:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	F_parser_errposition(m, l0, v921)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L13
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(537), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
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
	v935 = m.ExcPending
	if v935 != 0 {
		goto L13
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_26), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L13
	} else {
		goto L263
	}
L263:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v940)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L13
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(547), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
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
	v954 = m.ExcPending
	if v954 != 0 {
		goto L13
	} else {
		goto L267
	}
L267:
	;
	F_errmsg(m, int32(_a_F_analyzeCTE_27), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L13
	} else {
		goto L268
	}
L268:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_parser_errposition(m, l0, v959)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L13
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_analyzeCTE_4), int32(554), int32(_a_F_analyzeCTE_5))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L13
	} else {
		goto L270
	}
L270:
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
