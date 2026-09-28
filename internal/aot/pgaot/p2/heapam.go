package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_index_build_range_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v210 float64
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
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
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
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
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v760 float64
	_ = v760
	var v768 int32
	_ = v768
	var v770 float64
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v798 int32
	_ = v798
	var v820 float64
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v917 int32
	_ = v917
	var v947 float64
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v980 float64
	_ = v980
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int64
	_ = v989
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	v12 = int32(0)
	v30 = float64(0)
	v32 = m.G0
	v34 = v32 - int32(928)
	m.G0 = v34
	v37 = int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v38) < base.Ui32(int32(_a_F_heapam_index_build_range_scan_0)) {
		v47 = v37
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v49 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	if v42 == int32(99) {
		v47 = v37
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v45 = F_isTempToastNamespace(m, v42)
	mBase = m.M
	v47 = v45
	goto L2
L5:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	v55 = base.B2i32(v52 != int32(0))
	goto L7
L6:
	;
	v55 = int32(1)
	goto L7
L7:
	;
	v56 = F_CreateExecutorState(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return float64(0)
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+152))
	if v60 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = F_MakePerTupleExprContext(m, v56)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v65 = v60
	goto L12
L12:
	;
	v67 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v65 = v63
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v71 = F_ExecPrepareQual(m, v70, v56)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[0]))
	if v74 == int32(0) {
		v80 = v12
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l10 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v77 != 0 {
		v80 = v12
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v78 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v80 = v78
	goto L16
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L8
	} else {
		goto L282
	}
L21:
	;
	if l5 != 0 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	if v80 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	v112 = l10
	v113 = v111
	v114 = v12
	goto L21
L25:
	;
	v86 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	v90 = int32(_a_F_heapam_index_build_range_scan_1)
	goto L27
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v88 = F_RegisterSnapshot(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v90 = v88
	goto L27
L30:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v94&int32(1) == int32(0) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v99 = int32(0)
	if l3 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v106 = int32(449)
	goto L36
L35:
	;
	v106 = int32(321)
	goto L36
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v109 = m.T0[v108].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v90, v99, v99, v99, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v112 = v109
	v113 = v90
	v114 = base.B2i32(v80 == v99)
	goto L21
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
	if v116 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if l3 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	v121 = v116 + int32(20)
	goto L43
L42:
	;
	v121 = v112 + int32(40)
	goto L43
L43:
	;
	v122 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v121))))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	if v125 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L40
L45:
	;
	goto L44
L46:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4])))
	if v129&int32(1) == int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v134 = int32(_a_F_heapam_index_build_range_scan_2)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v137 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v136 + v137
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v140 + v137
	v144 = int32(0)
	v146 = int32(_a_F_heapam_index_build_range_scan_3)
	v147 = base.AtomicRmwOr32(m, v144, v146, v144)
	*(*int64)(unsafe.Add(mBase, uint32(v125+int32(120))+232)) = v122
	v155 = base.AtomicRmwOr32(m, v144, v146, v144)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v156 + v137
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v162 - v137
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+48)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v112)+44)) = l6
	goto L50
L49:
	;
	goto L50
L50:
	;
	v171 = F_heap_getnext(m, v112)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L52
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L8
	} else {
		goto L278
	}
L52:
	;
	if v171 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v179 = int32(-1)
	v184 = v171
	v198 = v179
	v202 = v179
	v210 = v30
	goto L56
L54:
	;
	v980 = v30
	goto L55
L55:
	;
	if l5 != 0 {
		goto L261
	} else {
		goto L262
	}
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	if v213 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v980 = v947
	goto L55
L58:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	if l5 == int32(0) {
		v281 = v216
		v283 = v198
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	if v281 != v202 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
	if v219 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v198 == v232 {
		v281 = v216
		v283 = v198
		goto L62
	} else {
		goto L74
	}
L65:
	;
	v222 = v219 + int32(28)
	goto L67
L66:
	;
	v222 = v112 + int32(44)
	goto L67
L67:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if base.Ui32(v223) < base.Ui32(v216) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v232 = v216 - v223
	goto L64
L69:
	;
	goto L70
L70:
	;
	if v219 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v228 = v219 + int32(20)
	goto L73
L72:
	;
	v228 = v112 + int32(40)
	goto L73
L73:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v232 = v229 + (v216 - v223)
	goto L64
L74:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	if v238 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	v281 = v279
	v283 = v232
	goto L62
L76:
	;
	goto L75
L77:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4])))
	if v242&int32(1) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v247 = int32(_a_F_heapam_index_build_range_scan_2)
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v250 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v249 + v250
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v253 + v250
	v257 = int32(0)
	v259 = int32(_a_F_heapam_index_build_range_scan_3)
	v260 = base.AtomicRmwOr32(m, v257, v259, v257)
	*(*int64)(unsafe.Add(mBase, uint32(v238+int32(128))+232)) = base.I64_extend_i32_u(v232)
	v268 = base.AtomicRmwOr32(m, v257, v259, v257)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v269 + v250
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v275 - v250
	goto L76
L79:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	if v285 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v316 = v202
	goto L81
L81:
	;
	if base.B2i32(v113 != int32(_a_F_heapam_index_build_range_scan_1)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	F_LockBufferInternal(m, v285, int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L8
	} else {
		goto L86
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v289+(v285^int32(-1))<<(uint(int32(2))%32))))
	v303 = v295
	goto L82
L84:
	;
	goto L85
L85:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	v303 = v297 + v285<<(uint(int32(13))%32) + int32(-8192)
	goto L82
L86:
	;
	F_heap_get_root_tuples(m, v303, v34+int32(48))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_UnlockBuffer(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	v316 = v314
	goto L81
L89:
	;
	v949 = F_heap_getnext(m, v112)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L8
	} else {
		goto L259
	}
L90:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	F_MemoryContextReset(m, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L8
	} else {
		goto L235
	}
L91:
	;
	goto L94
L92:
	;
	goto L93
L93:
	;
	v798 = int32(1)
	v820 = base.F64_add(v210, float64(1))
	goto L90
L94:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_LockBufferInternal(m, v352, int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	v357 = F_HeapTupleSatisfiesVacuum(m, v184, v80, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L108
	}
L97:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_UnlockBuffer(m, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L8
	} else {
		goto L231
	}
L98:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_UnlockBuffer(m, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L8
	} else {
		goto L230
	}
L99:
	;
	v768 = int32(0)
	v770 = base.F64_add(v210, float64(1))
	goto L98
L100:
	;
	v768 = int32(1)
	v770 = v760
	goto L98
L101:
	;
	v760 = base.F64_add(v210, float64(1))
	goto L100
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L8
	} else {
		goto L227
	}
L103:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_UnlockBuffer(m, v738)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L8
	} else {
		goto L226
	}
L104:
	;
	v734 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v734)
	goto L103
L105:
	;
	if l4 != 0 {
		goto L99
	} else {
		goto L164
	}
L106:
	;
	if l4 != 0 {
		goto L101
	} else {
		goto L112
	}
L107:
	;
	v359 = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+19)))
	if v361&int32(64) == v359 {
		v768 = v359
		v770 = v210
		goto L98
	} else {
		goto L109
	}
L108:
	;
	switch v357 {
	case 0:
		goto L103
	case 1:
		goto L101
	case 2:
		goto L107
	case 3:
		goto L106
	case 4:
		goto L105
	default:
		goto L102
	}
L109:
	;
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v360)+20)))
	if v366&int32(2048) != 0 {
		v768 = v359
		v770 = v210
		goto L98
	} else {
		goto L110
	}
L110:
	;
	if v366&int32(768) != int32(512) {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v768 = v359
	v770 = v210
	goto L98
L112:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v374)+20)))
	v376 = int32(768)
	if v375&v376 != v376 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v381 = v380
	goto L115
L114:
	;
	v381 = int32(2)
	goto L115
L115:
	;
	if base.Ui32(v381) < base.Ui32(int32(3)) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v513 != 0 {
		goto L101
	} else {
		goto L156
	}
L117:
	;
	v513 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v393 == v381 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v513 = int32(1)
	goto L116
L121:
	;
	goto L122
L122:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	if v397 <= int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v513 = v503
	goto L116
L124:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11]))
	if v401 == int32(0) {
		v503 = int32(0)
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[12]))
	v473 = int32(0)
	v476 = v397 - int32(1)
	goto L146
L127:
	;
	v406 = v401
	goto L128
L128:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406)+20))
	if v412 == int32(4) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v503 = int32(0)
	goto L123
L130:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v406)+80))
	if v466 != 0 {
		v406 = v466
		goto L128
	} else {
		goto L145
	}
L131:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	if v415 == int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v418 = int32(1)
	if v381 == v415 {
		v503 = v418
		goto L123
	} else {
		goto L133
	}
L133:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v406)+52))
	v422 = v420 - int32(1)
	if v422 < int32(0) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v406)+48))
	v428 = int32(0)
	v431 = v422
	goto L135
L135:
	;
	v436 = int32(2)
	v437 = base.I32_div_s(v431-v428, v436)
	v438 = v437 + v428
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v425+v438<<(uint(v436)%32))))
	if v442 == v381 {
		v503 = v418
		goto L123
	} else {
		goto L137
	}
L136:
	;
	goto L130
L137:
	;
	v451 = base.B2i32(v442-v381 < int32(0)) | base.B2i32(base.Ui32(v442) < base.Ui32(int32(3)))
	if v451 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v452 = v438 + int32(1)
	goto L140
L139:
	;
	v452 = v428
	goto L140
L140:
	;
	if v451 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v455 = v431
	goto L143
L142:
	;
	v455 = v438 - int32(1)
	goto L143
L143:
	;
	if v452 <= v455 {
		v428 = v452
		v431 = v455
		goto L135
	} else {
		goto L144
	}
L144:
	;
	goto L136
L145:
	;
	goto L129
L146:
	;
	v481 = int32(2)
	v482 = base.I32_div_s(v476-v473, v481)
	v483 = v482 + v473
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v471+v483<<(uint(v481)%32))))
	v488 = base.B2i32(v487 == v381)
	if v487 == v381 {
		v503 = v488
		goto L123
	} else {
		goto L148
	}
L147:
	;
	v503 = v488
	goto L123
L148:
	;
	v491 = base.B2i32(base.Ui32(v487) < base.Ui32(v381))
	if base.Ui32(v487) < base.Ui32(v381) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v492 = v483 + int32(1)
	goto L151
L150:
	;
	v492 = v473
	goto L151
L151:
	;
	if base.Ui32(v487) < base.Ui32(v381) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v495 = v476
	goto L154
L153:
	;
	v495 = v483 - int32(1)
	goto L154
L154:
	;
	if v492 <= v495 {
		v473 = v492
		v476 = v495
		goto L146
	} else {
		goto L155
	}
L155:
	;
	goto L147
L156:
	;
	if v47 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	if v55 == int32(0) {
		v760 = v210
		goto L100
	} else {
		goto L163
	}
L158:
	;
	v516 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	if v516 == int32(0) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v520 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_4), v34+int32(16))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L8
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1447), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	goto L157
L163:
	;
	v774 = v381
	goto L97
L164:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v536)+20)))
	if v537&int32(_a_F_heapam_index_build_range_scan_7) == int32(_a_F_heapam_index_build_range_scan_8) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	if base.Ui32(v545) < base.Ui32(int32(3)) {
		goto L171
	} else {
		goto L172
	}
L166:
	;
	v542 = F_HeapTupleGetUpdateXid(m, v536)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L8
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v545 = v544
	goto L165
L169:
	;
	v545 = v542
	goto L165
L170:
	;
	if v677 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L171:
	;
	v677 = int32(0)
	goto L170
L172:
	;
	goto L173
L173:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v557 == v545 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v677 = int32(1)
	goto L170
L175:
	;
	goto L176
L176:
	;
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	if v561 <= int32(0) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v677 = v667
	goto L170
L178:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11]))
	if v565 == int32(0) {
		v667 = int32(0)
		goto L177
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[12]))
	v637 = int32(0)
	v640 = v561 - int32(1)
	goto L200
L181:
	;
	v570 = v565
	goto L182
L182:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v570)+20))
	if v576 == int32(4) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v667 = int32(0)
	goto L177
L184:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v570)+80))
	if v630 != 0 {
		v570 = v630
		goto L182
	} else {
		goto L199
	}
L185:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	if v579 == int32(0) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v582 = int32(1)
	if v545 == v579 {
		v667 = v582
		goto L177
	} else {
		goto L187
	}
L187:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v570)+52))
	v586 = v584 - int32(1)
	if v586 < int32(0) {
		goto L184
	} else {
		goto L188
	}
L188:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v570)+48))
	v592 = int32(0)
	v595 = v586
	goto L189
L189:
	;
	v600 = int32(2)
	v601 = base.I32_div_s(v595-v592, v600)
	v602 = v601 + v592
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v589+v602<<(uint(v600)%32))))
	if v606 == v545 {
		v667 = v582
		goto L177
	} else {
		goto L191
	}
L190:
	;
	goto L184
L191:
	;
	v615 = base.B2i32(v606-v545 < int32(0)) | base.B2i32(base.Ui32(v606) < base.Ui32(int32(3)))
	if v615 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v616 = v602 + int32(1)
	goto L194
L193:
	;
	v616 = v592
	goto L194
L194:
	;
	if v615 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v619 = v595
	goto L197
L196:
	;
	v619 = v602 - int32(1)
	goto L197
L197:
	;
	if v616 <= v619 {
		v592 = v616
		v595 = v619
		goto L189
	} else {
		goto L198
	}
L198:
	;
	goto L190
L199:
	;
	goto L183
L200:
	;
	v645 = int32(2)
	v646 = base.I32_div_s(v640-v637, v645)
	v647 = v646 + v637
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v635+v647<<(uint(v645)%32))))
	v652 = base.B2i32(v651 == v545)
	if v651 == v545 {
		v667 = v652
		goto L177
	} else {
		goto L202
	}
L201:
	;
	v667 = v652
	goto L177
L202:
	;
	v655 = base.B2i32(base.Ui32(v651) < base.Ui32(v545))
	if base.Ui32(v651) < base.Ui32(v545) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v656 = v647 + int32(1)
	goto L205
L204:
	;
	v656 = v637
	goto L205
L205:
	;
	if base.Ui32(v651) < base.Ui32(v545) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v659 = v640
	goto L208
L207:
	;
	v659 = v647 - int32(1)
	goto L208
L208:
	;
	if v656 <= v659 {
		v637 = v656
		v640 = v659
		goto L200
	} else {
		goto L209
	}
L209:
	;
	goto L201
L210:
	;
	if v47 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v717 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718)+19)))
	if v719&int32(64) == v717 {
		v768 = v717
		v770 = v210
		goto L98
	} else {
		goto L224
	}
L213:
	;
	if v55 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	v682 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L8
	} else {
		goto L215
	}
L215:
	;
	if v682 == int32(0) {
		goto L213
	} else {
		goto L216
	}
L216:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v686 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_9), v34+int32(32))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L8
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1506), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L8
	} else {
		goto L218
	}
L218:
	;
	goto L213
L219:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702)+19)))
	if v703&int32(64) == int32(0) {
		goto L99
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v774 = v545
	goto L97
L222:
	;
	v708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v702)+20)))
	if v708&int32(2048)|base.B2i32(v708&int32(768) == int32(512)) != 0 {
		goto L99
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718)+20)))
	if v724&int32(2048)|base.B2i32(v724&int32(768) == int32(512)) != 0 {
		v768 = v717
		v770 = v210
		goto L98
	} else {
		goto L225
	}
L225:
	;
	goto L104
L226:
	;
	v947 = v210
	goto L89
L227:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_10), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L8
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1576), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L8
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v798 = v768
	v820 = v770
	goto L90
L231:
	;
	F_XactLockTableWait(m, v774, l0, v184+int32(4), int32(6))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L8
	} else {
		goto L232
	}
L232:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	if v783 == int32(0) {
		goto L94
	} else {
		goto L233
	}
L233:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L8
	} else {
		goto L234
	}
L234:
	;
	goto L94
L235:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_ExecStoreBufferHeapTuple(m, v184, v67, v825)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L8
	} else {
		goto L236
	}
L236:
	;
	if v71 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v828 = int32(_a_F_heapam_index_build_range_scan_11)
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[13]))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[13])) = v831
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v71)+24))
	v836 = m.T0[v835].(func(*base.Module, int32, int32, int32) int64)(m, v71, v65, v34+int32(42))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L8
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	F_FormIndexDatum(m, l2, v67, v56, v34+int32(672), v34+int32(640))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L8
	} else {
		goto L242
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[13])) = v829
	if v836 == int64(0) {
		v947 = v820
		goto L89
	} else {
		goto L241
	}
L241:
	;
	goto L239
L242:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v850 = int32(*(*int16)(unsafe.Add(mBase, uint32(v849)+18)))
	if v850 < int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+8)))
	v858 = v853<<(uint(int32(1))%32) + v34 + int32(46)
	v859 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v858))))
	if v859 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v184+int32(4), v34+int32(672), v34+int32(640), v798, l9)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L8
	} else {
		goto L258
	}
L246:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	if v862 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v892 = v859
	goto L248
L248:
	;
	if base.Ui32(int32(2048)) <= base.Ui32((v892-int32(1))&int32(_a_F_heapam_index_build_range_scan_12)) {
		goto L51
	} else {
		goto L256
	}
L249:
	;
	F_LockBufferInternal(m, v862, int32(1))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L8
	} else {
		goto L253
	}
L250:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v866+(v862^int32(-1))<<(uint(int32(2))%32))))
	v880 = v872
	goto L249
L251:
	;
	goto L252
L252:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	v880 = v874 + v862<<(uint(int32(13))%32) + int32(-8192)
	goto L249
L253:
	;
	F_heap_get_root_tuples(m, v880, v34+int32(48))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L8
	} else {
		goto L254
	}
L254:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v112)+60))
	F_UnlockBuffer(m, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L8
	} else {
		goto L255
	}
L255:
	;
	v891 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v858))))
	v892 = v891
	goto L248
L256:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+46)) = uint16(v892)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+42)) = v899
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v34+int32(42), v34+int32(672), v34+int32(640), v798, l9)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L8
	} else {
		goto L257
	}
L257:
	;
	v947 = v820
	goto L89
L258:
	;
	v947 = v820
	goto L89
L259:
	;
	if v949 != 0 {
		v184 = v949
		v198 = v283
		v202 = v316
		v210 = v947
		goto L56
	} else {
		goto L260
	}
L260:
	;
	goto L57
L261:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v112)+32))
	if v983 != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+188))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+12))
	m.T0[v1036].(func(*base.Module, int32))(m, v112)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L8
	} else {
		goto L271
	}
L264:
	;
	v988 = v983 + int32(20)
	goto L266
L265:
	;
	v988 = v112 + int32(40)
	goto L266
L266:
	;
	v989 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v988))))
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	if v992 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L263
L268:
	;
	goto L267
L269:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4])))
	if v996&int32(1) == int32(0) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1001 = int32(_a_F_heapam_index_build_range_scan_2)
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v1004 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v1003 + v1004
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1007 + v1004
	v1011 = int32(0)
	v1013 = int32(_a_F_heapam_index_build_range_scan_3)
	v1014 = base.AtomicRmwOr32(m, v1011, v1013, v1011)
	*(*int64)(unsafe.Add(mBase, uint32(v992+int32(128))+232)) = v989
	v1022 = base.AtomicRmwOr32(m, v1011, v1013, v1011)
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1023 + v1004
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5])) = v1029 - v1004
	goto L268
L271:
	;
	if v114 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	F_UnregisterSnapshot(m, v113)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L8
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	F_ExecDropSingleTupleTableSlot(m, v67)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L8
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	F_FreeExecutorState(m, v56)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	v1045 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v1045
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v1045
	m.G0 = v34 + int32(928)
	return v980
L278:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L8
	} else {
		goto L279
	}
L279:
	;
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+6)))
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+4)))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v853
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v1062 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1060 | v1061<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_13), v34)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L8
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_5), int32(1656), int32(_a_F_heapam_index_build_range_scan_6))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_14), int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L8
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_15), int32(931), int32(_a_F_heapam_index_build_range_scan_16))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L8
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_index_fetch_begin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(-4294967296)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		return v5
	}
}
func F_heapam_index_validate_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v376 int64
	_ = v376
	var v379 int64
	_ = v379
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
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
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 float64
	_ = v542
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1232)
	m.G0 = v20
	v22 = F_CreateExecutorState(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+152))
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = F_MakePerTupleExprContext(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v29 = v24
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = F_MakeSingleTupleTableSlot(m, v30, int32(_a_F_heapam_index_validate_scan_0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v29 = v27
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v36 = F_ExecPrepareQual(m, v35, v22)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[0]))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L107
	}
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[1])))
	if v41&int32(1) == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v47 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3, v47, v47, v47, int32(321))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+40)))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	if v58 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v99 = F_heap_getnext(m, v53)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[3])))
	if v62&int32(1) == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v67 = int32(_a_F_heapam_index_validate_scan_1)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	v70 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4])) = v69 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v73 + v70
	v77 = int32(0)
	v79 = int32(_a_F_heapam_index_validate_scan_2)
	v80 = base.AtomicRmwOr32(m, v77, v79, v77)
	*(*int64)(unsafe.Add(mBase, uint32(v58+int32(120))+232)) = v55
	v88 = base.AtomicRmwOr32(m, v77, v79, v77)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v89 + v70
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4])) = v95 - v70
	goto L16
L19:
	;
	if v99 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v101 = int32(-1)
	v106 = v99
	v111 = v101
	v112 = v6
	v115 = v101
	v118 = v6
	goto L23
L21:
	;
	goto L22
L22:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+188))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	m.T0[v584].(func(*base.Module, int32))(m, v53)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L104
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[5]))
	if v121 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v124 = *(*float64)(unsafe.Add(mBase, uint32(l4)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+8)) = base.F64_add(v124, float64(1))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	if base.B2i32(v128 == v115)&base.B2i32(v115 != int32(-1)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	if v139 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v181 = v115
	goto L31
L31:
	;
	if v111 != v181 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	v181 = v180
	goto L31
L33:
	;
	goto L32
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[3])))
	if v143&int32(1) == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v148 = int32(_a_F_heapam_index_validate_scan_1)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	v151 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4])) = v150 + v151
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v154 + v151
	v158 = int32(0)
	v160 = int32(_a_F_heapam_index_validate_scan_2)
	v161 = base.AtomicRmwOr32(m, v158, v160, v158)
	*(*int64)(unsafe.Add(mBase, uint32(v139+int32(128))+232)) = base.I64_extend_i32_u(v128)
	v169 = base.AtomicRmwOr32(m, v158, v160, v158)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v170 + v151
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4])) = v176 - v151
	goto L33
L36:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
	if v183 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v218 = v111
	goto L38
L38:
	;
	v220 = v106 + int32(4)
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+40)) = uint16(v221)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v223
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v226)+18)))
	if v227 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	F_LockBufferInternal(m, v183, int32(1))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187+(v183^int32(-1))<<(uint(int32(2))%32))))
	v201 = v193
	goto L39
L41:
	;
	goto L42
L42:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[7]))
	v201 = v195 + v183<<(uint(int32(13))%32) + int32(-8192)
	goto L39
L43:
	;
	F_heap_get_root_tuples(m, v201, v20+int32(352))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
	F_UnlockBuffer(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.MemoryFill(m, v20+int32(48), int32(0), int32(291))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	v218 = v217
	goto L38
L46:
	;
	v563 = F_heap_getnext(m, v53)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L102
	}
L47:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v241)+47)))
	if v508 != 0 {
		v555 = v499
		v561 = v505
		goto L46
	} else {
		goto L92
	}
L48:
	;
	v461 = int32(0)
	v463 = v20 + int32(36)
	v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v453)+2)))
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v453))))
	v469 = int32(16)
	v471 = v467 | v468<<(uint(v469)%32)
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v463)+2)))
	v473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v463))))
	v476 = v472 | v473<<(uint(v469)%32)
	if base.Ui32(v471) < base.Ui32(v476) {
		v487 = int32(-1)
		goto L87
	} else {
		goto L88
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L82
	}
L50:
	;
	v230 = int32(1)
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225<<(uint(v230)%32)+v20)+350)))
	if base.Ui32(int32(2048)) <= base.Ui32((v233-v230)&int32(_a_F_heapam_index_validate_scan_3)) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	v241 = v225
	goto L52
L52:
	;
	if v118 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+40)) = uint16(v233)
	v241 = v233
	goto L52
L54:
	;
	if v112 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v408 = v112
	goto L56
L56:
	;
	v499 = v408
	v505 = int32(1)
	goto L47
L57:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v286 = int32(0)
	v292 = F_tuplesort_getdatum(m, v284, int32(1), v286, v20+int32(24), v20+int32(23), v286)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L67
	}
L58:
	;
	v247 = v20 + int32(36)
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
	v253 = int32(16)
	v255 = v251 | v252<<(uint(v253)%32)
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+2)))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247))))
	v260 = v256 | v257<<(uint(v253)%32)
	if base.Ui32(v255) < base.Ui32(v260) {
		v271 = int32(-1)
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if int32(0) <= v271 {
		v453 = v112
		goto L48
	} else {
		goto L64
	}
L60:
	;
	goto L59
L61:
	;
	if base.Ui32(v260) < base.Ui32(v255) {
		v271 = int32(1)
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if base.Ui32(v265) < base.Ui32(v266) {
		v271 = int32(-1)
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v271 = base.B2i32(base.Ui32(v266) < base.Ui32(v265))
	goto L60
L64:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
	if v274|v275<<(uint(int32(16))%32) != v218 {
		goto L57
	} else {
		goto L65
	}
L65:
	;
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	v282 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v280)+47)) = uint8(v282)
	goto L57
L66:
	;
	v408 = int32(0)
	goto L56
L67:
	;
	if v292 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+46)) = uint16(v296)
	v299 = int64(base.Ui64(v296) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+44)) = uint16(v299)
	v302 = int64(base.Ui64(v296) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+42)) = uint16(v302)
	goto L69
L69:
	;
	v322 = v20 + int32(42)
	v324 = v20 + int32(36)
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322)+2)))
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322))))
	v330 = int32(16)
	v332 = v328 | v329<<(uint(v330)%32)
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+2)))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324))))
	v337 = v333 | v334<<(uint(v330)%32)
	if base.Ui32(v332) < base.Ui32(v337) {
		v348 = int32(-1)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if int32(0) <= v348 {
		v453 = v322
		goto L48
	} else {
		goto L76
	}
L72:
	;
	goto L71
L73:
	;
	if base.Ui32(v337) < base.Ui32(v332) {
		v348 = int32(1)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322)+4)))
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v324)+4)))
	if base.Ui32(v342) < base.Ui32(v343) {
		v348 = int32(-1)
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v348 = base.B2i32(base.Ui32(v343) < base.Ui32(v342))
	goto L72
L76:
	;
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+44)))
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+42)))
	if v218 == v351|v352<<(uint(int32(16))%32) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+46)))
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v357)+47)) = uint8(v359)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v363 = int32(0)
	v369 = F_tuplesort_getdatum(m, v361, int32(1), v363, v20+int32(24), v20+int32(23), v363)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v369 == int32(0) {
		goto L66
	} else {
		goto L81
	}
L81:
	;
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+46)) = uint16(v373)
	v376 = int64(base.Ui64(v373) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+44)) = uint16(v376)
	v379 = int64(base.Ui64(v373) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+42)) = uint16(v379)
	goto L69
L82:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+6)))
	v425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v426 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v424 | v425<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_validate_scan_4), v20)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_heapam_index_validate_scan_5), int32(1833), int32(_a_F_heapam_index_validate_scan_6))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	if v487 <= int32(0) {
		v555 = v453
		v561 = v461
		goto L46
	} else {
		goto L91
	}
L87:
	;
	goto L86
L88:
	;
	if base.Ui32(v476) < base.Ui32(v471) {
		v487 = int32(1)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v453)+4)))
	v482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v463)+4)))
	if base.Ui32(v481) < base.Ui32(v482) {
		v487 = int32(-1)
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v487 = base.B2i32(base.Ui32(v482) < base.Ui32(v481))
	goto L87
L91:
	;
	v499 = v453
	v505 = v461
	goto L47
L92:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	F_MemoryContextReset(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v513 = F_ExecStoreHeapTuple(m, v106, v32, int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v36 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v515 = int32(_a_F_heapam_index_validate_scan_7)
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[8]))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[8])) = v518
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v523 = m.T0[v522].(func(*base.Module, int32, int32, int32) int64)(m, v36, v29, v20+int32(24))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v531 = v20 + int32(976)
	v533 = v20 + int32(944)
	F_FormIndexDatum(m, l2, v32, v22, v531, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[8])) = v516
	if v523 == int64(0) {
		v555 = v499
		v561 = v505
		goto L46
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	v540 = F_index_insert(m, l1, v531, v533, v20+int32(36), l0, v538, int32(0), l2)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v542 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+24)) = base.F64_add(v542, float64(1))
	v555 = v499
	v561 = v505
	goto L46
L102:
	;
	if v563 != 0 {
		v106 = v563
		v111 = v218
		v112 = v555
		v115 = v181
		v118 = v561
		goto L23
	} else {
		goto L103
	}
L103:
	;
	goto L24
L104:
	;
	F_ExecDropSingleTupleTableSlot(m, v32)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_FreeExecutorState(m, v22)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v591 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v591
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v591
	m.G0 = v20 + int32(1232)
	return
L107:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_index_validate_scan_8), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_heapam_index_validate_scan_9), int32(931), int32(_a_F_heapam_index_validate_scan_10))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_relation_needs_toast_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v172 int32
	_ = v172
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 <= v2 {
		v172 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v172 & int32(1)
L2:
	;
	v20 = int32(0)
	v21 = v2
	v22 = v16
	v25 = v2
	v28 = v2
	goto L3
L3:
	;
	v35 = v15 + v22<<(uint(int32(3))%32) + v20*int32(100)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+119)))
	if v36 != 0 {
		v137 = v21
		v138 = v22
		v141 = v25
		v142 = v28
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if (v141^int32(-1)|v142)&int32(1) != 0 {
		v172 = v141
		goto L1
	} else {
		goto L36
	}
L5:
	;
	v145 = v20 + int32(1)
	if v145 < v138 {
		v20 = v145
		v21 = v137
		v22 = v138
		v25 = v141
		v28 = v142
		goto L3
	} else {
		goto L35
	}
L6:
	;
	v38 = v35 + int32(28)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+90)))
	if v39 == int32(118) {
		v137 = v21
		v138 = v22
		v141 = v25
		v142 = v28
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+83)))
	switch v43 - int32(99) {
	case 0:
		v71 = v21
		v72 = int32(-1)
		goto L8
	case 1:
		goto L9
	default:
		goto L12
	case 6:
		goto L10
	case 16:
		goto L11
	}
L8:
	;
	v73 = v71 & v72
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+72)))
	if int32(0) < v74 {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v71 = v21 + int32(7)
	v72 = int32(-8)
	goto L8
L10:
	;
	v71 = v21 + int32(3)
	v72 = int32(-4)
	goto L8
L11:
	;
	v71 = v21 + int32(1)
	v72 = int32(-2)
	goto L8
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = base.I32_extend8_s(v43)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_needs_toast_table_0), v13)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_needs_toast_table_1), int32(322), int32(_a_F_heapam_relation_needs_toast_table_2))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v137 = v73 + v74
	v138 = v22
	v141 = v25
	v142 = v28
	goto L5
L18:
	;
	goto L19
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v81 = int32(-1)
	if v79 < int32(0) {
		v121 = v81
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v124 = int32(0)
	if v124 < v123 {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v123 = v121
	goto L20
L22:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v78-int32(1042)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v117 = int32(8)
	v118 = base.I32_div_s(v79+int32(7), v117)
	v121 = v118 + v117
	goto L21
L24:
	;
	if v78 != int32(1700) {
		v121 = v81
		goto L21
	} else {
		goto L28
	}
L25:
	;
	switch v78 - int32(1560) {
	case 0, 2:
		goto L23
	case 1:
		v121 = v81
		goto L21
	default:
		goto L24
	}
L26:
	;
	goto L27
L27:
	;
	v90 = F_GetDatabaseEncoding(m)
	mBase = m.M
	v91 = F_pg_encoding_max_length(m, v90)
	mBase = m.M
	v92 = int32(4)
	v123 = v91*(v79-v92) + v92
	goto L20
L28:
	;
	v100 = int32(4)
	if v79 < v100 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v114 = int32(-1)
	goto L31
L30:
	;
	v114 = int32(base.Ui32(int32(base.Ui32(v79-v100)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_heapam_relation_needs_toast_table_3) + int32(8)
	goto L31
L31:
	;
	v123 = v114
	goto L20
L32:
	;
	v127 = v123
	goto L34
L33:
	;
	v127 = v124
	goto L34
L34:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+84)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v137 = v127 + v73
	v138 = v136
	v141 = base.B2i32(v132 != int32(112)) | v25
	v142 = base.B2i32(v123 < int32(0)) | v28
	goto L5
L35:
	;
	goto L4
L36:
	;
	v152 = int32(7)
	v155 = base.I32_div_s(v138+v152, int32(8))
	v158 = int32(-8)
	v172 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32((v155+int32(30))&v158+(v137+v152)&v158))
	goto L1
}
func F_heapam_scan_analyze_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v386 float64
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 float64
	_ = v400
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v19 < v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L21
	} else {
		goto L117
	}
L2:
	;
	m.G0 = v16 + int32(16)
	return v446
L3:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v38) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+(v19^int32(-1))<<(uint(int32(2))%32))))
	v37 = v29
	goto L3
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[1]))
	v37 = v31 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	v48 = int32(base.Ui32(v38+int32(_a_F_heapam_scan_analyze_next_tuple_0))>>(uint(int32(2))%32)) & int32(_a_F_heapam_scan_analyze_next_tuple_1)
	goto L9
L8:
	;
	v48 = int32(0)
	goto L9
L9:
	;
	if base.Ui32(v18) <= base.Ui32(v48) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = l3 + int32(52)
	v58 = v18
	goto L13
L11:
	;
	v425 = v19
	goto L12
L12:
	;
	F_UnlockReleaseBuffer(m, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L21
	} else {
		goto L115
	}
L13:
	;
	v69 = v37 + int32(20) + v58<<(uint(int32(2))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch int32(base.Ui32(v70)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L17
	default:
		goto L15
	case 2:
		goto L16
	}
L14:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v425 = v411
	goto L12
L15:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v408 = v406 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v408
	if base.Ui32(v408) <= base.Ui32(v48) {
		v58 = v408
		goto L13
	} else {
		goto L114
	}
L16:
	;
	v400 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v400, float64(1))
	goto L15
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+60)) = uint16(v58)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+58)) = uint16(v77)
	v81 = int32(base.Ui32(v77) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v37 + v86&int32(_a_F_heapam_scan_analyze_next_tuple_2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(base.Ui32(v91) >> (uint(int32(17)) % 32))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v98 = F_HeapTupleSatisfiesVacuumHorizon(m, v53, v95, v16+int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v386 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = base.F64_add(v386, float64(1))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_ExecStoreBufferHeapTuple(m, v53, l3, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L21
	} else {
		goto L113
	}
L19:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+20)))
	if v244&int32(_a_F_heapam_scan_analyze_next_tuple_3) == int32(_a_F_heapam_scan_analyze_next_tuple_4) {
		goto L68
	} else {
		goto L69
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+20)))
	v104 = int32(768)
	if v103&v104 != v104 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	return int32(0)
L22:
	;
	switch v98 {
	case 0, 2:
		goto L16
	case 1:
		goto L18
	case 3:
		goto L20
	case 4:
		goto L19
	default:
		goto L1
	}
L23:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v110 = v108
	goto L25
L24:
	;
	v110 = int32(2)
	goto L25
L25:
	;
	if base.Ui32(v110) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v242 != 0 {
		goto L18
	} else {
		goto L66
	}
L27:
	;
	v242 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v122 == v110 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v242 = int32(1)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v126 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v242 = v232
	goto L26
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v130 == int32(0) {
		v232 = int32(0)
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v202 = int32(0)
	v205 = v126 - int32(1)
	goto L56
L37:
	;
	v135 = v130
	goto L38
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)+20))
	if v141 == int32(4) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v232 = int32(0)
	goto L33
L40:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v135)+80))
	if v195 != 0 {
		v135 = v195
		goto L38
	} else {
		goto L55
	}
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v144 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v147 = int32(1)
	if v110 == v144 {
		v232 = v147
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v135)+52))
	v151 = v149 - int32(1)
	if v151 < int32(0) {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	v157 = int32(0)
	v160 = v151
	goto L45
L45:
	;
	v165 = int32(2)
	v166 = base.I32_div_s(v160-v157, v165)
	v167 = v166 + v157
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v154+v167<<(uint(v165)%32))))
	if v171 == v110 {
		v232 = v147
		goto L33
	} else {
		goto L47
	}
L46:
	;
	goto L40
L47:
	;
	v180 = base.B2i32(v171-v110 < int32(0)) | base.B2i32(base.Ui32(v171) < base.Ui32(int32(3)))
	if v180 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v181 = v167 + int32(1)
	goto L50
L49:
	;
	v181 = v157
	goto L50
L50:
	;
	if v180 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v184 = v160
	goto L53
L52:
	;
	v184 = v167 - int32(1)
	goto L53
L53:
	;
	if v181 <= v184 {
		v157 = v181
		v160 = v184
		goto L45
	} else {
		goto L54
	}
L54:
	;
	goto L46
L55:
	;
	goto L39
L56:
	;
	v210 = int32(2)
	v211 = base.I32_div_s(v205-v202, v210)
	v212 = v211 + v202
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v200+v212<<(uint(v210)%32))))
	v217 = base.B2i32(v216 == v110)
	if v216 == v110 {
		v232 = v217
		goto L33
	} else {
		goto L58
	}
L57:
	;
	v232 = v217
	goto L33
L58:
	;
	v220 = base.B2i32(base.Ui32(v216) < base.Ui32(v110))
	if base.Ui32(v216) < base.Ui32(v110) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v221 = v212 + int32(1)
	goto L61
L60:
	;
	v221 = v202
	goto L61
L61:
	;
	if base.Ui32(v216) < base.Ui32(v110) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v224 = v205
	goto L64
L63:
	;
	v224 = v212 - int32(1)
	goto L64
L64:
	;
	if v221 <= v224 {
		v202 = v221
		v205 = v224
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L57
L66:
	;
	goto L15
L67:
	;
	if base.Ui32(v252) < base.Ui32(int32(3)) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	v249 = F_HeapTupleGetUpdateXid(m, v243)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	v252 = v251
	goto L67
L71:
	;
	v252 = v249
	goto L67
L72:
	;
	if v384 != 0 {
		goto L16
	} else {
		goto L112
	}
L73:
	;
	v384 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v264 == v252 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v384 = int32(1)
	goto L72
L77:
	;
	goto L78
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v268 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v384 = v374
	goto L72
L80:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v272 == int32(0) {
		v374 = int32(0)
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v344 = int32(0)
	v347 = v268 - int32(1)
	goto L102
L83:
	;
	v277 = v272
	goto L84
L84:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	if v283 == int32(4) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v374 = int32(0)
	goto L79
L86:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v277)+80))
	if v337 != 0 {
		v277 = v337
		goto L84
	} else {
		goto L101
	}
L87:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	if v286 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v289 = int32(1)
	if v252 == v286 {
		v374 = v289
		goto L79
	} else {
		goto L89
	}
L89:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v277)+52))
	v293 = v291 - int32(1)
	if v293 < int32(0) {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v277)+48))
	v299 = int32(0)
	v302 = v293
	goto L91
L91:
	;
	v307 = int32(2)
	v308 = base.I32_div_s(v302-v299, v307)
	v309 = v308 + v299
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v296+v309<<(uint(v307)%32))))
	if v313 == v252 {
		v374 = v289
		goto L79
	} else {
		goto L93
	}
L92:
	;
	goto L86
L93:
	;
	v322 = base.B2i32(v313-v252 < int32(0)) | base.B2i32(base.Ui32(v313) < base.Ui32(int32(3)))
	if v322 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v323 = v309 + int32(1)
	goto L96
L95:
	;
	v323 = v299
	goto L96
L96:
	;
	if v322 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v326 = v302
	goto L99
L98:
	;
	v326 = v309 - int32(1)
	goto L99
L99:
	;
	if v323 <= v326 {
		v299 = v323
		v302 = v326
		goto L91
	} else {
		goto L100
	}
L100:
	;
	goto L92
L101:
	;
	goto L85
L102:
	;
	v352 = int32(2)
	v353 = base.I32_div_s(v347-v344, v352)
	v354 = v353 + v344
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v342+v354<<(uint(v352)%32))))
	v359 = base.B2i32(v358 == v252)
	if v358 == v252 {
		v374 = v359
		goto L79
	} else {
		goto L104
	}
L103:
	;
	v374 = v359
	goto L79
L104:
	;
	v362 = base.B2i32(base.Ui32(v358) < base.Ui32(v252))
	if base.Ui32(v358) < base.Ui32(v252) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v363 = v354 + int32(1)
	goto L107
L106:
	;
	v363 = v344
	goto L107
L107:
	;
	if base.Ui32(v358) < base.Ui32(v252) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v366 = v347
	goto L110
L109:
	;
	v366 = v354 - int32(1)
	goto L110
L110:
	;
	if v363 <= v366 {
		v344 = v363
		v347 = v366
		goto L102
	} else {
		goto L111
	}
L111:
	;
	goto L103
L112:
	;
	goto L18
L113:
	;
	v393 = int32(1)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v394 + v393
	v446 = v393
	goto L2
L114:
	;
	goto L14
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	m.T0[v431].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L21
	} else {
		goto L116
	}
L116:
	;
	v446 = v5
	goto L2
L117:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_scan_analyze_next_tuple_5), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_heapam_scan_analyze_next_tuple_6), int32(1110), int32(_a_F_heapam_scan_analyze_next_tuple_7))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_scan_sample_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
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
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	v4 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = v23 & int32(256)
	if v25 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_LockBufferInternal(m, v28, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v34 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+10)))
	if v53&int32(4) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(v34^int32(-1))<<(uint(int32(2))%32))))
	v52 = v44
	goto L6
L8:
	;
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[1]))
	v52 = v46 + v34<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+29)))
	v60 = v57 ^ int32(1)
	goto L12
L11:
	;
	v60 = v4
	goto L12
L12:
	;
	v64 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	v68 = l0 + int32(68)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v69) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v77 = int32(base.Ui32(v69+int32(_a_F_heapam_scan_sample_next_tuple_0)) >> (uint(int32(2)) % 32))
	goto L15
L14:
	;
	v77 = int32(0)
	goto L15
L15:
	;
	goto L17
L16:
	;
	return base.B2i32(base.Ui32(v108&int32(_a_F_heapam_scan_sample_next_tuple_1)) < base.Ui32(int32(2048)))
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_tuple[2]))
	if v101 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_ExecStoreBufferHeapTuple(m, v68, l2, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L60
	}
L19:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v105 = m.T0[v104].(func(*base.Module, int32, int32, int32) int32)(m, l1, v21, v77&int32(_a_F_heapam_scan_sample_next_tuple_1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L26
	}
L22:
	;
	goto L21
L23:
	;
	goto L18
L24:
	;
	if v185 == int32(0) {
		goto L17
	} else {
		goto L59
	}
L25:
	;
	if v25 != 0 {
		goto L17
	} else {
		goto L57
	}
L26:
	;
	v108 = v105 - int32(1)
	if base.Ui32(v108&int32(_a_F_heapam_scan_sample_next_tuple_1)) <= base.Ui32(int32(2047)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v115 = v52 + int32(20) + v105<<(uint(int32(2))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v116&int32(_a_F_heapam_scan_sample_next_tuple_2) != int32(_a_F_heapam_scan_sample_next_tuple_3) {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v25 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v52 + v116&int32(_a_F_heapam_scan_sample_next_tuple_4)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v105)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v64)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(base.Ui32(v125) >> (uint(int32(17)) % 32))
	v132 = int32(1)
	if v60&v132 != 0 {
		v185 = v132
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v25 != 0 {
		goto L24
	} else {
		goto L48
	}
L32:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v135&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v139 == v138 {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v179 = F_HeapTupleSatisfiesVisibility(m, v68, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L47
	}
L36:
	;
	v145 = v138
	v152 = v139
	goto L37
L37:
	;
	v163 = int32(1)
	v165 = int32(base.Ui32(v152-v145)>>(uint(v163)%32)) + v145
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(116)+v165<<(uint(v163)%32)))))
	v170 = base.B2i32(v105 == v169)
	if v105 == v169 {
		v185 = v170
		goto L31
	} else {
		goto L39
	}
L38:
	;
	v185 = v170
	goto L31
L39:
	;
	v173 = base.B2i32(base.Ui32(v105) < base.Ui32(v169))
	if base.Ui32(v105) < base.Ui32(v169) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v174 = v145
	goto L42
L41:
	;
	v174 = v165 + int32(1)
	goto L42
L42:
	;
	if base.Ui32(v105) < base.Ui32(v169) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = v165
	goto L45
L44:
	;
	v175 = v152
	goto L45
L45:
	;
	if base.Ui32(v174) < base.Ui32(v175) {
		v145 = v174
		v152 = v175
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v185 = v179
	goto L31
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, v185, v201, v68, v202, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v185 == int32(0) {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_UnlockBuffer(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_UnlockBuffer(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	m.T0[v217].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	goto L16
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, int32(0), v221, v68, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	goto L17
L59:
	;
	goto L23
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+272))
	if v232 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+268)))
	if v235 != int32(1) {
		goto L16
	} else {
		goto L64
	}
L62:
	;
	v242 = v232
	goto L63
L63:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v242)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v242)+24)) = v243 + int64(1)
	goto L16
L64:
	;
	F_pgstat_assoc_relation(m, v231)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+272))
	v242 = v241
	goto L63
}
func F_heapam_slot_callbacks(m *base.Module, l0 int32) int32 {
	return int32(_a_F_heapam_slot_callbacks_0)
}
func F_heapam_tuple_satisfies_snapshot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_LockBufferInternal(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
		v12 = F_HeapTupleSatisfiesVisibility(m, v10, l2, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			F_UnlockBuffer(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_heapam_tuple_tid_valid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	v3 = int32(0)
	if l1 == v3 {
		v16 = v3
	} else {
		v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v6 == int32(0) {
			v16 = v3
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v16 = base.B2i32(base.Ui32(v10|v11<<(uint(int32(16))%32)) < base.Ui32(v9))
		}
	}
	return v16
}
