package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineQueryRewrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
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
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
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
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
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
	var v741 int32
	_ = v741
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	v15 = m.G0
	v17 = v15 - int32(240)
	m.G0 = v17
	v20 = F_table_open(m, l2, int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	if l5 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L9
	} else {
		goto L198
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L9
	} else {
		goto L194
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L9
	} else {
		goto L190
	}
L5:
	;
	v642 = int32(_a_F_DefineQueryRewrite_0)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v648 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[0])))
	if base.B2i32(v645 == int32(0))|base.B2i32(v645 != v648) != 0 {
		v666 = v645
		v667 = v648
		goto L179
	} else {
		goto L180
	}
L6:
	;
	v588 = int32(0)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v589 <= v588 {
		goto L5
	} else {
		goto L167
	}
L7:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+119)))
	switch v234 - int32(109) {
	case 0, 9:
		goto L70
	default:
		goto L71
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L66
	}
L9:
	;
	return
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
	v25 = v23 - int32(109)
	v32 = int32(0)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v25))|base.B2i32(int32(1)<<(uint(v25)%32)&int32(553) == v32) == v32 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[1])))
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L61
	}
L14:
	;
	v42 = int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if base.Ui32(v43) < base.Ui32(int32(_a_F_DefineQueryRewrite_1)) {
		v52 = v42
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[2]))
	v56 = F_object_ownercheck(m, int32(1259), l2, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L22
	}
L17:
	;
	if v52 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v47 == int32(99) {
		v52 = v42
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v50 = F_isTempToastNamespace(m, v47)
	mBase = m.M
	v52 = v50
	goto L18
L21:
	;
	goto L16
L22:
	;
	if v56 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+119)))
	switch v62 - int32(73) {
	case 0, 32:
		v72 = int32(20)
		goto L27
	default:
		goto L28
	case 10:
		goto L32
	case 29:
		goto L29
	case 36:
		goto L30
	case 45:
		goto L31
	}
L24:
	;
	goto L25
L25:
	;
	if l7 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	F_aclcheck_error(m, int32(2), v74, v75+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L33
	}
L27:
	;
	v74 = v72
	goto L26
L28:
	;
	v72 = int32(41)
	goto L27
L29:
	;
	v74 = int32(18)
	goto L26
L30:
	;
	v74 = int32(23)
	goto L26
L31:
	;
	v74 = int32(51)
	goto L26
L32:
	;
	v74 = int32(37)
	goto L26
L33:
	;
	goto L25
L34:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if int32(0) < v80 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if l4 == int32(1) {
		goto L7
	} else {
		goto L60
	}
L37:
	;
	v93 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	if l4 != int32(1) {
		goto L6
	} else {
		goto L59
	}
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v93<<(uint(int32(2))%32))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	if v103 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v154 = v93 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v154 < v155 {
		v93 = v154
		goto L40
	} else {
		goto L58
	}
L43:
	;
	v107 = F_getInsertSelectQuery(m, v102, int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	if v107 != v102 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	switch v110 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L42
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L53
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_2), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_3), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(293), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_6), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_7), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(298), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	goto L41
L59:
	;
	goto L7
L60:
	;
	goto L5
L61:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v182 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_8), v17)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v189)+119)))
	F_errdetail_relkind_not_supported(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(263), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v205 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_9), v17+int32(96))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(269), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L9
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
	if l7 != 0 {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v244 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_10), v17+int32(16))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v253)+119)))
	F_errdetail_relkind_not_supported(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(314), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L9
	} else {
		goto L163
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L9
	} else {
		goto L159
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L155
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L9
	} else {
		goto L151
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L9
	} else {
		goto L147
	}
L82:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if int32(2) <= v262 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L9
	} else {
		goto L142
	}
L85:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v269 != int32(1) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+42)))
	if v272 == int32(1) {
		goto L79
	} else {
		goto L88
	}
L88:
	;
	if l3 != 0 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v268)+76))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	F_checkRuleResultList(m, v275, v276, int32(1), base.B2i32(v234 != int32(109)))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L90
	}
L90:
	;
	if l6 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v349 = int32(105)
	v350 = int32(_a_F_DefineQueryRewrite_0)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[0])))
	if base.B2i32(v353 == int32(0))|base.B2i32(v353 != v356) != 0 {
		v374 = v353
		v375 = v356
		goto L106
	} else {
		goto L107
	}
L92:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	if v282 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v285 <= int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v299 = int32(0)
	goto L95
L95:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v288+v299<<(uint(int32(2))%32))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	if v308 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L9
	} else {
		goto L101
	}
L97:
	;
	v312 = v299 + int32(1)
	if v285 != v312 {
		v299 = v312
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L91
L101:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v321 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_11), v17-int32(-64))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(385), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	if v374-v375 == int32(0) {
		v741 = l1
		v752 = v349
		goto L1
	} else {
		goto L112
	}
L106:
	;
	goto L105
L107:
	;
	v359 = l1
	v360 = v350
	goto L108
L108:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)))
	if v364 == int32(0) {
		v374 = v364
		v375 = v363
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v374 = v364
	v375 = v363
	goto L106
L110:
	;
	v367 = int32(1)
	if v364 == v363 {
		v359 = v359 + v367
		v360 = v360 + v367
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v379 = int32(_a_F_DefineQueryRewrite_12)
	goto L115
L113:
	;
	if v417-v418 != 0 {
		goto L77
	} else {
		goto L126
	}
L115:
	;
	goto L116
L116:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v386 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v387 = l1
	v388 = v379
	v389 = int32(4)
	v390 = v386
	goto L121
L118:
	;
	v413 = v379
	v417 = int32(0)
	goto L119
L119:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	goto L113
L120:
	;
	v413 = v408
	v417 = v410
	goto L119
L121:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if base.B2i32(v390 != v392)|base.B2i32(v392 == int32(0)) != 0 {
		v408 = v388
		v410 = v390
		goto L120
	} else {
		goto L123
	}
L122:
	;
	v408 = v402
	v410 = int32(0)
	goto L120
L123:
	;
	v398 = v389 - int32(1)
	if v398 == int32(0) {
		v408 = v388
		v410 = v390
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v401 = int32(1)
	v402 = v388 + v401
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+1)))
	if v403 != 0 {
		v387 = v387 + v401
		v388 = v402
		v389 = v398
		v390 = v403
		goto L121
	} else {
		goto L125
	}
L125:
	;
	goto L122
L126:
	;
	v426 = int32(4)
	v427 = l1 + v426
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v430 = v428 + v426
	goto L129
L127:
	;
	if v468-v469 != 0 {
		goto L77
	} else {
		goto L140
	}
L129:
	;
	goto L130
L130:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v437 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v438 = v427
	v439 = v430
	v440 = int32(56)
	v441 = v437
	goto L135
L132:
	;
	v464 = v430
	v468 = int32(0)
	goto L133
L133:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464))))
	goto L127
L134:
	;
	v464 = v459
	v468 = v461
	goto L133
L135:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	if base.B2i32(v441 != v443)|base.B2i32(v443 == int32(0)) != 0 {
		v459 = v439
		v461 = v441
		goto L134
	} else {
		goto L137
	}
L136:
	;
	v459 = v453
	v461 = int32(0)
	goto L134
L137:
	;
	v449 = v440 - int32(1)
	if v449 == int32(0) {
		v459 = v439
		v461 = v441
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v452 = int32(1)
	v453 = v439 + v452
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
	if v454 != 0 {
		v438 = v438 + v452
		v439 = v453
		v440 = v449
		v441 = v454
		goto L135
	} else {
		goto L139
	}
L139:
	;
	goto L136
L140:
	;
	v478 = F_pstrdup(m, int32(_a_F_DefineQueryRewrite_0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L9
	} else {
		goto L141
	}
L141:
	;
	v741 = v478
	v752 = v349
	goto L1
L142:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_13), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_14), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(323), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_15), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(331), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L9
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_16), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L9
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(341), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L9
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_17), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(349), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_18), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(357), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(_a_F_DefineQueryRewrite_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v572 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_19), v17+int32(48))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(410), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	v601 = v588
	v604 = int32(0)
	goto L168
L168:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606+v601<<(uint(int32(2))%32))))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+96))
	if v611 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L5
L170:
	;
	if v604&int32(1) != 0 {
		goto L4
	} else {
		goto L173
	}
L171:
	;
	v622 = v604
	goto L172
L172:
	;
	v624 = v601 + int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v624 < v625 {
		v601 = v624
		v604 = v622
		goto L168
	} else {
		goto L177
	}
L173:
	;
	if l3 != 0 {
		goto L3
	} else {
		goto L174
	}
L174:
	;
	if l5 == int32(0) {
		goto L2
	} else {
		goto L175
	}
L175:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v617 = int32(0)
	F_checkRuleResultList(m, v611, v616, v617, v617)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	v622 = int32(1)
	goto L172
L177:
	;
	goto L169
L178:
	;
	if v666-v667 != 0 {
		v741 = l1
		v752 = int32(97)
		goto L1
	} else {
		goto L185
	}
L179:
	;
	goto L178
L180:
	;
	v651 = l1
	v652 = v642
	goto L181
L181:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+1)))
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+1)))
	if v656 == int32(0) {
		v666 = v656
		v667 = v655
		goto L179
	} else {
		goto L183
	}
L182:
	;
	v666 = v656
	v667 = v655
	goto L179
L183:
	;
	v659 = int32(1)
	if v656 == v655 {
		v651 = v651 + v659
		v652 = v652 + v659
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L9
	} else {
		goto L186
	}
L186:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L9
	} else {
		goto L187
	}
L187:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = int32(_a_F_DefineQueryRewrite_0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v676 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_20), v17+int32(80))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(460), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L9
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L9
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(435), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L9
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L9
	} else {
		goto L195
	}
L195:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_22), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L9
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(440), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L9
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L9
	} else {
		goto L199
	}
L199:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_23), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L9
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(444), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L9
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	v956 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v953
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2618)
	F_relation_close(m, v20, v956)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L9
	} else {
		goto L263
	}
L203:
	;
	v756 = int32(0)
	if l7 == v756 {
		v953 = v756
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v760 = F_nodeToString(m, l3)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L9
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v762 = F_nodeToString(m, l7)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L208
	}
L208:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+200)) = int64(0)
	v767 = v17 + int32(136)
	v769 = F_strncpy(m, v767, v741, int32(64))
	mBase = m.M
	v770 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v769)+63)) = uint8(v770)
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+228)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v17)+224)) = int32(79)
	v775 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = (l4<<(uint(v775)%32) + int32(805306368)) >> (uint(v775) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+216)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v767
	v784 = F_cstring_to_text(m, v760)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L9
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+232)) = v784
	v787 = F_cstring_to_text(m, v762)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L9
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+236)) = v787
	v792 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L9
	} else {
		goto L212
	}
L212:
	;
	v795 = F_SearchSysCache2(m, int32(60), l2, v741)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L9
	} else {
		goto L215
	}
L213:
	;
	v864 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v17)+124)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = int32(2618)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+108)) = int32(1259)
	v875 = v17 + int32(120)
	F_recordDependencyOn(m, v875, v17+int32(108), v752)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L9
	} else {
		goto L234
	}
L214:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v792)+52))
	v843 = F_heap_modify_tuple(m, v795, v836, v17+int32(208), v17+int32(200), v17+int32(120))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L9
	} else {
		goto L229
	}
L215:
	;
	if v795 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = int64(72340168543043584)
	if l6 != 0 {
		goto L214
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v822 = F_GetNewOidWithIndex(m, v792, int32(2692), int32(1))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L9
	} else {
		goto L225
	}
L219:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L9
	} else {
		goto L220
	}
L220:
	;
	F_errcode(m, int32(_a_F_DefineQueryRewrite_24))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L9
	} else {
		goto L221
	}
L221:
	;
	v806 = F_get_rel_name(m, l2)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L9
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v806
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v741
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_25), v17+int32(32))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L9
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(105), int32(_a_F_DefineQueryRewrite_26))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L9
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v822
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v792)+52))
	v830 = F_heap_form_tuple(m, v825, v17+int32(208), v17+int32(200))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L9
	} else {
		goto L226
	}
L226:
	;
	F_CatalogTupleInsert(m, v792, v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L9
	} else {
		goto L227
	}
L227:
	;
	F_pfree(m, v830)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L9
	} else {
		goto L228
	}
L228:
	;
	v863 = v822
	goto L213
L229:
	;
	F_CatalogTupleUpdate(m, v792, v843+int32(4), v843)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L9
	} else {
		goto L230
	}
L230:
	;
	F_ReleaseCatCache(m, v795)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L9
	} else {
		goto L231
	}
L231:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v843)+16))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851)+22)))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v851+v852)))
	F_pfree(m, v843)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L9
	} else {
		goto L232
	}
L232:
	;
	v859 = F_deleteDependencyRecordsFor(m, int32(2618), v854, int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	v863 = v854
	goto L213
L234:
	;
	F_recordDependencyOnExpr(m, v875, l7, int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L9
	} else {
		goto L235
	}
L235:
	;
	if l3 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v886 = F_getInsertSelectQuery(m, v884, int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[3]))
	if v892 != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v886)+52))
	F_recordDependencyOnExpr(m, v875, l3, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L9
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	v894 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2618), v863, v894, v894)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L9
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	F_relation_close(m, v792, int32(3))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L9
	} else {
		goto L245
	}
L244:
	;
	goto L243
L245:
	;
	v901 = m.G0
	v903 = v901 - int32(16)
	m.G0 = v903
	v907 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L9
	} else {
		goto L246
	}
L246:
	;
	v911 = F_SearchSysCacheCopy(m, int32(57), l2, int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L9
	} else {
		goto L248
	}
L247:
	;
	v953 = v863
	goto L202
L248:
	;
	if v911 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v911)+16))
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913)+22)))
	v915 = v913 + v914
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+124)))
	if v916 != int32(1) {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	goto L251
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L9
	} else {
		goto L260
	}
L252:
	;
	F_pfree(m, v911)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L9
	} else {
		goto L258
	}
L253:
	;
	v919 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v915)+124)) = uint8(v919)
	F_CatalogTupleUpdate(m, v907, v911+int32(4), v911)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L9
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	F_CacheInvalidateRelcacheByTuple(m, v911)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L9
	} else {
		goto L257
	}
L256:
	;
	goto L252
L257:
	;
	goto L252
L258:
	;
	F_relation_close(m, v907, int32(3))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	m.G0 = v903 + int32(16)
	goto L247
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v903))) = l2
	F_errmsg_internal(m, int32(_a_F_DefineQueryRewrite_27), v903)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L9
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_28), int32(65), int32(_a_F_DefineQueryRewrite_29))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L9
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	m.G0 = v17 + int32(240)
	return
}
func F_assign_query_collations(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = F_query_tree_walker_impl(m, l1, int32(482), l0, int32(10))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_extract_query_dependencies_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	v3 = int32(0)
	if l0 == v3 {
		v157 = v3
		v158 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v157 == int32(0)) & v158
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(67) {
		v84 = l0
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+44)))
	if v103 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L4:
	;
	F_fix_expr_common(m, l1, v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L38
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != int32(6) {
		v97 = l0
		v101 = v3
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 != int32(213) {
		v50 = v15
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v57 = base.B2i32(v16 == int32(213))
	v59 = v50
	goto L24
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v20 = F_extract_query_dependencies_walker(m, v19, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v24 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v157 = int32(1)
	v158 = v3
	goto L1
L12:
	;
	goto L13
L13:
	;
	v28 = v24
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(67) {
		v84 = v28
		goto L4
	} else {
		goto L16
	}
L15:
	;
	v157 = int32(1)
	v158 = v3
	goto L1
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v37 != int32(6) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v97 = v28
	v101 = int32(1)
	goto L3
L18:
	;
	goto L19
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != int32(213) {
		v50 = v41
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = F_extract_query_dependencies_walker(m, v45, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v49 != 0 {
		v28 = v49
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	if v83 != 0 {
		v97 = v83
		v101 = v57
		goto L3
	} else {
		goto L37
	}
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v61 - int32(241) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		goto L30
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v59 = v81
	goto L24
L27:
	;
	v83 = v79
	goto L23
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v76 == int32(6) {
		v80 = v75
		goto L26
	} else {
		goto L36
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 == int32(6) {
		v80 = v71
		goto L26
	} else {
		goto L35
	}
L30:
	;
	if v61 != int32(201) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v83 = int32(0)
	goto L23
L32:
	;
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 != int32(6) {
		v79 = v67
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v80 = v67
	goto L26
L35:
	;
	v79 = v71
	goto L27
L36:
	;
	v79 = v75
	goto L27
L37:
	;
	v157 = v57
	v158 = v3
	goto L1
L38:
	;
	v95 = F_expression_tree_walker_impl(m, v84, int32(838), l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v157 = base.B2i32(v9 == int32(67))
	v158 = v95
	goto L1
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+81)) = uint8(v107)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	if v109 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v151 = F_query_tree_walker_impl(m, v97, int32(838), l1, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L55
	}
L44:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v120 = v3
	goto L46
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v120<<(uint(int32(2))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	switch v126 {
	case 0:
		goto L51
	case 1, 7:
		goto L50
	default:
		goto L48
	}
L47:
	;
	goto L43
L48:
	;
	v140 = v120 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v140 < v141 {
		v120 = v140
		goto L46
	} else {
		goto L54
	}
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+56))
	v134 = F_lappend_oid(m, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L53
	}
L50:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v128 == int32(0) {
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v131 = v127
	goto L49
L52:
	;
	v131 = v128
	goto L49
L53:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+56)) = v134
	goto L48
L54:
	;
	goto L47
L55:
	;
	v157 = v101
	v158 = v151
	goto L1
}
func F_query_to_xml_and_xmlschema(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_text_to_cstring(m, v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v22 = F_pg_detoast_datum_packed(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = F_text_to_cstring(m, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = int32(0)
						v31 = F_SPI_prepare(m, v18, v29, v29)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							if v31 != 0 {
								v33 = F_SPI_cursor_open(m, v31)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									if v33 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v18
											F_errmsg_internal(m, int32(_a_F_query_to_xml_and_xmlschema_0), v11+int32(16))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_query_to_xml_and_xmlschema_1), int32(3161), int32(_a_F_query_to_xml_and_xmlschema_2))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
										v38 = int32(0)
										v40 = base.B2i32(v20 != v38)
										v41 = F_map_sql_table_to_xmlschema(m, v37, v38, v40, v24)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											v43 = F_strlen(m, v41)
											mBase = m.M
											v45 = v43 + int32(1)
											v46 = F_SPI_palloc(m, v45)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												if v45 != 0 {
													base.MemoryCopy(m, v46, v41, v45)
												} else {
												}
												F_SPI_cursor_close(m, v33)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return int32(0)
												} else {
													v51 = F_SPI_finish(m)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
														return int32(0)
													} else {
														v54 = F_query_to_xml_internal(m, v18, int32(0), v46, v40, v24)
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return int32(0)
														} else {
															v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
															v58 = F_cstring_to_text_with_len(m, v56, v57)
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v58
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v18
									F_errmsg_internal(m, int32(_a_F_query_to_xml_and_xmlschema_3), v11)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_query_to_xml_and_xmlschema_1), int32(3158), int32(_a_F_query_to_xml_and_xmlschema_2))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
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
	}
}
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
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
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v291 int32
	_ = v291
	v5 = int32(0)
	if l3&int32(64) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v24 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v23, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v22 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = F_palloc(m, int32(168))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	base.MemoryCopy(m, v16, l0, int32(168))
	v22 = v16
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+152))
	v28 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v27, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+152)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v36 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v35, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v44 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v43, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v48 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v47, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+144))
	v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+112))
	v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v55, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v22)+128))
	v60 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v59, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	v64 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v63, l2)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v64
	if l3&int32(128) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if l3&int32(2) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v70 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v69, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	if v85 == int32(0) {
		v145 = v5
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	v74 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v73, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v78 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v77, l2)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	v82 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v81, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v82
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v145
	goto L18
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v88 <= int32(0) {
		v145 = v5
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = v5
	v99 = v5
	goto L29
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v99<<(uint(int32(2))%32))))
	v109 = F_palloc(m, int32(56))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v145 = v133
	goto L26
L31:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v107)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+48)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v107)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+40)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v107)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+32)) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v107)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+24)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v107)+24))
	v126 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v125, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	v130 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v129, l2)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = v130
	v133 = F_lappend(m, v97, v109)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v136 = v99 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v136 < v137 {
		v97 = v133
		v99 = v136
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v173
	v175 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	if v178 == v175 {
		v291 = v175
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v169 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v164, l2)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v171 = F_copyObjectImpl(m, v164)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v173 = v169
	goto L36
L41:
	;
	v173 = v171
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v291
	return v22
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if int32(0) < v181 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v195 = v175
	v198 = v175
	goto L47
L45:
	;
	v272 = v175
	goto L46
L46:
	;
	v291 = v272
	goto L42
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v198<<(uint(int32(2))%32))))
	v208 = F_palloc(m, int32(136))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L49
	}
L48:
	;
	v272 = v261
	goto L46
L49:
	;
	base.MemoryCopy(m, v208, v206, int32(136))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	switch v212 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	case 5:
		goto L52
	default:
		goto L50
	case 9:
		goto L51
	}
L50:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v206)+128))
	v258 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v257, l2)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L77
	}
L51:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)+120))
	if l3&int32(256) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L52:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v206)+80))
	v244 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v243, l2)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L71
	}
L53:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	v240 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v239, l2)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L70
	}
L54:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	v236 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v235, l2)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L69
	}
L55:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v206)+52))
	if l3&int32(4) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	if l3&int32(1) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v214 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v213, l2)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = v214
	goto L50
L59:
	;
	v220 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v217, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v223 = F_copyObjectImpl(m, v217)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+36)) = v220
	goto L50
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+36)) = v223
	goto L50
L64:
	;
	v229 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v226, l2)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v232 = F_copyObjectImpl(m, v226)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+52)) = v229
	goto L50
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+52)) = v232
	goto L50
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+68)) = v236
	goto L50
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+76)) = v240
	goto L50
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+80)) = v244
	goto L50
L72:
	;
	v250 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v247, l2)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v253 = F_copyObjectImpl(m, v247)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+120)) = v250
	goto L50
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+120)) = v253
	goto L50
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+128)) = v258
	v261 = F_lappend(m, v195, v208)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v264 = v198 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v264 < v265 {
		v195 = v261
		v198 = v264
		goto L47
	} else {
		goto L79
	}
L79:
	;
	goto L48
}
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v9 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v8, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v9 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v14 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v13, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v17 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v16, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v17 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v23 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v22, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v26 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v25, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v26 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v29 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v29 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v38 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v37, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v38 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v41 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v40, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v41 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if l3&int32(128) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if l3&int32(2) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v45, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v57 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L29:
	;
	if v46 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v48, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v49 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v52 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v55 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v54, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v55 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v68 = int32(0)
	goto L39
L39:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v68<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v76 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v75, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L1
L41:
	;
	goto L40
L42:
	;
	if v76 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v79 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v79 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v82 = v68 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v82 < v83 {
		v68 = v82
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L25
L47:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v97 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v96, l2)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if l3&int32(8) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v97 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v104 = F_range_table_walker_impl(m, v103, l1, l2, l3)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	return int32(0)
L55:
	;
	if v104 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
}
