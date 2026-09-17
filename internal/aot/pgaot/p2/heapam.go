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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 float64
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v711 float64
	_ = v711
	var v719 int32
	_ = v719
	var v721 float64
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v751 int32
	_ = v751
	var v773 float64
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v900 float64
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v932 float64
	_ = v932
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v940 int64
	_ = v940
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	v12 = int32(0)
	v30 = float64(0)
	v31 = m.G0
	v33 = v31 - int32(800)
	m.G0 = v33
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v37) < base.Ui32(int32(_a_F_heapam_index_build_range_scan_0)) {
		v46 = v36
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	if v48 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 == int32(99) {
		v46 = v36
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v44 = F_isTempToastNamespace(m, v41)
	mBase = m.M
	v46 = v44
	goto L2
L5:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	v54 = base.B2i32(v51 != int32(0))
	goto L7
L6:
	;
	v54 = int32(1)
	goto L7
L7:
	;
	v55 = F_CreateExecutorState(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return float64(0)
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v62 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v64 = v59
	goto L12
L12:
	;
	v66 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v64 = v62
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v70 = F_ExecPrepareQual(m, v69, v55)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[0]))
	if v73 == int32(0) {
		v79 = v12
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l10 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	if v76 != 0 {
		v79 = v12
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v77 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v79 = v77
	goto L16
L20:
	;
	if l5 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	v103 = l10
	v104 = v102
	v105 = v12
	goto L20
L24:
	;
	v85 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	v89 = int32(_a_F_heapam_index_build_range_scan_1)
	goto L26
L26:
	;
	v90 = int32(0)
	if l3 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v87 = F_RegisterSnapshot(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v89 = v87
	goto L26
L29:
	;
	v97 = int32(449)
	goto L31
L30:
	;
	v97 = int32(321)
	goto L31
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v100 = m.T0[v99].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v89, v90, v90, v90, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v103 = v100
	v104 = v89
	v105 = base.B2i32(v79 == v90)
	goto L20
L33:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v107 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if l3 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v112 = v107 + int32(20)
	goto L38
L37:
	;
	v112 = v103 + int32(36)
	goto L38
L38:
	;
	v113 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v112))))
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v116 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L35
L40:
	;
	goto L39
L41:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v120&int32(1) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v125 = int32(_a_F_heapam_index_build_range_scan_2)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v127 + v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v131 + v128
	*(*int64)(unsafe.Add(mBase, uint32(v116+int32(120))+232)) = v113
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v139 + v128
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v145 - v128
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = l6
	goto L45
L44:
	;
	goto L45
L45:
	;
	v154 = F_heap_getnext(m, v103)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L8
	} else {
		goto L273
	}
L47:
	;
	if v154 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v162 = int32(-1)
	v167 = v154
	v181 = v162
	v185 = v162
	v193 = v30
	goto L51
L49:
	;
	v932 = v30
	goto L50
L50:
	;
	if l5 != 0 {
		goto L256
	} else {
		goto L257
	}
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4]))
	if v195 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v932 = v900
	goto L50
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	if l5 == int32(0) {
		v255 = v198
		v257 = v181
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	if v255 != v185 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v201 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v181 == v214 {
		v255 = v198
		v257 = v181
		goto L57
	} else {
		goto L69
	}
L60:
	;
	v204 = v201 + int32(28)
	goto L62
L61:
	;
	v204 = v103 + int32(40)
	goto L62
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if base.Ui32(v205) < base.Ui32(v198) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v214 = v198 - v205
	goto L59
L64:
	;
	goto L65
L65:
	;
	if v201 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v210 = v201 + int32(20)
	goto L68
L67:
	;
	v210 = v103 + int32(36)
	goto L68
L68:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v214 = v211 + (v198 - v205)
	goto L59
L69:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v220 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v255 = v253
	v257 = v214
	goto L57
L71:
	;
	goto L70
L72:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v224&int32(1) == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v229 = int32(_a_F_heapam_index_build_range_scan_2)
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v232 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v231 + v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v235 + v232
	*(*int64)(unsafe.Add(mBase, uint32(v220+int32(128))+232)) = base.I64_extend_i32_u(v214)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v243 + v232
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v249 - v232
	goto L71
L74:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v259 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v291 = v185
	goto L76
L76:
	;
	if base.B2i32(v104 != int32(_a_F_heapam_index_build_range_scan_1)) == int32(0) {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	F_LockBuffer(m, v259, int32(1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L81
	}
L78:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v263+(v259^int32(-1))<<(uint(int32(2))%32))))
	v277 = v269
	goto L77
L79:
	;
	goto L80
L80:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	v277 = v271 + v259<<(uint(int32(13))%32) + int32(-8192)
	goto L77
L81:
	;
	F_heap_get_root_tuples(m, v277, v33+int32(48))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v285, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v103)+52))
	v291 = v289
	goto L76
L84:
	;
	v901 = F_heap_getnext(m, v103)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L8
	} else {
		goto L254
	}
L85:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	F_MemoryContextReset(m, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L8
	} else {
		goto L230
	}
L86:
	;
	goto L89
L87:
	;
	goto L88
L88:
	;
	v751 = int32(1)
	v773 = base.F64_add(v193, float64(1))
	goto L85
L89:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v326, int32(1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	v331 = F_HeapTupleSatisfiesVacuum(m, v167, v79, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L103
	}
L92:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v728, int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L8
	} else {
		goto L226
	}
L93:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v722, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L8
	} else {
		goto L225
	}
L94:
	;
	v719 = int32(0)
	v721 = base.F64_add(v193, float64(1))
	goto L93
L95:
	;
	v719 = int32(1)
	v721 = v711
	goto L93
L96:
	;
	v711 = base.F64_add(v193, float64(1))
	goto L95
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L8
	} else {
		goto L222
	}
L98:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v688, int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L8
	} else {
		goto L221
	}
L99:
	;
	v684 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+122)) = uint8(v684)
	goto L98
L100:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L159
	}
L101:
	;
	if l4 != 0 {
		goto L96
	} else {
		goto L107
	}
L102:
	;
	v333 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+19)))
	if v335&int32(64) == v333 {
		v719 = v333
		v721 = v193
		goto L93
	} else {
		goto L104
	}
L103:
	;
	switch v331 {
	case 0:
		goto L98
	case 1:
		goto L96
	case 2:
		goto L102
	case 3:
		goto L101
	case 4:
		goto L100
	default:
		goto L97
	}
L104:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+20)))
	if v340&int32(2048) != 0 {
		v719 = v333
		v721 = v193
		goto L93
	} else {
		goto L105
	}
L105:
	;
	if v340&int32(768) != int32(512) {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v719 = v333
	v721 = v193
	goto L93
L107:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+20)))
	v350 = int32(768)
	if v349&v350 != v350 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v355 = v354
	goto L110
L109:
	;
	v355 = int32(2)
	goto L110
L110:
	;
	if base.Ui32(v355) < base.Ui32(int32(3)) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v475 != 0 {
		goto L96
	} else {
		goto L151
	}
L112:
	;
	v475 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	if v366 == v355 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v475 = int32(1)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	if v370 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v475 = v467
	goto L111
L119:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v374 == int32(0) {
		v467 = int32(0)
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	v438 = int32(0)
	v440 = v370 - int32(1)
	goto L141
L122:
	;
	v379 = v374
	goto L123
L123:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v379)+20))
	if v384 == int32(4) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v467 = int32(0)
	goto L118
L125:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v379)+80))
	if v431 != 0 {
		v379 = v431
		goto L123
	} else {
		goto L140
	}
L126:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	if v387 == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v390 = int32(1)
	if v355 == v387 {
		v467 = v390
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v379)+52))
	v394 = v392 - int32(1)
	if v394 < int32(0) {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v399 = int32(0)
	v401 = v394
	goto L130
L130:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v379)+48))
	v407 = int32(2)
	v408 = base.I32_div_s(v401-v399, v407)
	v409 = v408 + v399
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v405+v409<<(uint(v407)%32))))
	if v413 == v355 {
		v467 = v390
		goto L118
	} else {
		goto L132
	}
L131:
	;
	goto L125
L132:
	;
	v417 = F_TransactionIdPrecedes(m, v413, v355)
	mBase = m.M
	if v417 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v418 = v409 + int32(1)
	goto L135
L134:
	;
	v418 = v399
	goto L135
L135:
	;
	if v417 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v421 = v401
	goto L138
L137:
	;
	v421 = v409 - int32(1)
	goto L138
L138:
	;
	if v418 <= v421 {
		v399 = v418
		v401 = v421
		goto L130
	} else {
		goto L139
	}
L139:
	;
	goto L131
L140:
	;
	goto L124
L141:
	;
	v445 = int32(2)
	v446 = base.I32_div_s(v440-v438, v445)
	v447 = v446 + v438
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v436+v447<<(uint(v445)%32))))
	v452 = base.B2i32(v451 == v355)
	if v451 == v355 {
		v467 = v452
		goto L118
	} else {
		goto L143
	}
L142:
	;
	v467 = v452
	goto L118
L143:
	;
	v455 = base.B2i32(base.Ui32(v451) < base.Ui32(v355))
	if base.Ui32(v451) < base.Ui32(v355) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v456 = v447 + int32(1)
	goto L146
L145:
	;
	v456 = v438
	goto L146
L146:
	;
	if base.Ui32(v451) < base.Ui32(v355) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v459 = v440
	goto L149
L148:
	;
	v459 = v447 - int32(1)
	goto L149
L149:
	;
	if v456 <= v459 {
		v438 = v456
		v440 = v459
		goto L141
	} else {
		goto L150
	}
L150:
	;
	goto L142
L151:
	;
	if v46 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v54 == int32(0) {
		v711 = v193
		goto L95
	} else {
		goto L158
	}
L153:
	;
	v478 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L154
	}
L154:
	;
	if v478 == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v482 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_3), v33+int32(16))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_4), int32(1484), int32(_a_F_heapam_index_build_range_scan_5))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L157
	}
L157:
	;
	goto L152
L158:
	;
	v726 = v355
	goto L92
L159:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v498)+20)))
	if v499&int32(_a_F_heapam_index_build_range_scan_6) == int32(_a_F_heapam_index_build_range_scan_7) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if base.Ui32(v507) < base.Ui32(int32(3)) {
		goto L166
	} else {
		goto L167
	}
L161:
	;
	v504 = F_HeapTupleGetUpdateXid(m, v498)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L8
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	v507 = v506
	goto L160
L164:
	;
	v507 = v504
	goto L160
L165:
	;
	if v627 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L166:
	;
	v627 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[7]))
	if v518 == v507 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v627 = int32(1)
	goto L165
L170:
	;
	goto L171
L171:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[8]))
	if v522 <= int32(0) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v627 = v619
	goto L165
L173:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[9]))
	if v526 == int32(0) {
		v619 = int32(0)
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[10]))
	v590 = int32(0)
	v592 = v522 - int32(1)
	goto L195
L176:
	;
	v531 = v526
	goto L177
L177:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v531)+20))
	if v536 == int32(4) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v619 = int32(0)
	goto L172
L179:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v531)+80))
	if v583 != 0 {
		v531 = v583
		goto L177
	} else {
		goto L194
	}
L180:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v539 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v542 = int32(1)
	if v507 == v539 {
		v619 = v542
		goto L172
	} else {
		goto L182
	}
L182:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v531)+52))
	v546 = v544 - int32(1)
	if v546 < int32(0) {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v551 = int32(0)
	v553 = v546
	goto L184
L184:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v531)+48))
	v559 = int32(2)
	v560 = base.I32_div_s(v553-v551, v559)
	v561 = v560 + v551
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v557+v561<<(uint(v559)%32))))
	if v565 == v507 {
		v619 = v542
		goto L172
	} else {
		goto L186
	}
L185:
	;
	goto L179
L186:
	;
	v569 = F_TransactionIdPrecedes(m, v565, v507)
	mBase = m.M
	if v569 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v570 = v561 + int32(1)
	goto L189
L188:
	;
	v570 = v551
	goto L189
L189:
	;
	if v569 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v573 = v553
	goto L192
L191:
	;
	v573 = v561 - int32(1)
	goto L192
L192:
	;
	if v570 <= v573 {
		v551 = v570
		v553 = v573
		goto L184
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	goto L178
L195:
	;
	v597 = int32(2)
	v598 = base.I32_div_s(v592-v590, v597)
	v599 = v598 + v590
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v588+v599<<(uint(v597)%32))))
	v604 = base.B2i32(v603 == v507)
	if v603 == v507 {
		v619 = v604
		goto L172
	} else {
		goto L197
	}
L196:
	;
	v619 = v604
	goto L172
L197:
	;
	v607 = base.B2i32(base.Ui32(v603) < base.Ui32(v507))
	if base.Ui32(v603) < base.Ui32(v507) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v608 = v599 + int32(1)
	goto L200
L199:
	;
	v608 = v590
	goto L200
L200:
	;
	if base.Ui32(v603) < base.Ui32(v507) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v611 = v592
	goto L203
L202:
	;
	v611 = v599 - int32(1)
	goto L203
L203:
	;
	if v608 <= v611 {
		v590 = v608
		v592 = v611
		goto L195
	} else {
		goto L204
	}
L204:
	;
	goto L196
L205:
	;
	if v46 != 0 {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	v667 = int32(0)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+19)))
	if v669&int32(64) == v667 {
		v719 = v667
		v721 = v193
		goto L93
	} else {
		goto L219
	}
L208:
	;
	if v54 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v632 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L8
	} else {
		goto L210
	}
L210:
	;
	if v632 == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v636 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_8), v33+int32(32))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L8
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_4), int32(1543), int32(_a_F_heapam_index_build_range_scan_5))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L8
	} else {
		goto L213
	}
L213:
	;
	goto L208
L214:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+19)))
	if v653&int32(64) == int32(0) {
		goto L94
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v726 = v507
	goto L92
L217:
	;
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v652)+20)))
	if v658&int32(2048)|base.B2i32(v658&int32(768) == int32(512)) != 0 {
		goto L94
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v668)+20)))
	if v674&int32(2048)|base.B2i32(v674&int32(768) == int32(512)) != 0 {
		v719 = v667
		v721 = v193
		goto L93
	} else {
		goto L220
	}
L220:
	;
	goto L99
L221:
	;
	v900 = v193
	goto L84
L222:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_9), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L8
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_4), int32(1613), int32(_a_F_heapam_index_build_range_scan_5))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L8
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
	v751 = v719
	v773 = v721
	goto L85
L226:
	;
	F_XactLockTableWait(m, v726, l0, v167+int32(4), int32(6))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L8
	} else {
		goto L227
	}
L227:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[4]))
	if v736 == int32(0) {
		goto L89
	} else {
		goto L228
	}
L228:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L8
	} else {
		goto L229
	}
L229:
	;
	goto L89
L230:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_ExecStoreBufferHeapTuple(m, v167, v66, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L8
	} else {
		goto L231
	}
L231:
	;
	if v70 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v780 = int32(_a_F_heapam_index_build_range_scan_10)
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11]))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11])) = v783
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v788 = m.T0[v787].(func(*base.Module, int32, int32, int32) int32)(m, v70, v64, v33+int32(42))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L8
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	F_FormIndexDatum(m, l2, v66, v55, v33+int32(672), v33+int32(640))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L8
	} else {
		goto L237
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[11])) = v781
	if v788 == int32(0) {
		v900 = v773
		goto L84
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v802 = int32(*(*int16)(unsafe.Add(mBase, uint32(v801)+18)))
	if v802 < int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+8)))
	v810 = v805<<(uint(int32(1))%32) + v33 + int32(46)
	v811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v810))))
	if v811 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v167+int32(4), v33+int32(672), v33+int32(640), v751, l9)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L8
	} else {
		goto L253
	}
L241:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	if v814 < int32(0) {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v845 = v811
	goto L243
L243:
	;
	if base.Ui32(int32(2048)) <= base.Ui32((v845-int32(1))&int32(_a_F_heapam_index_build_range_scan_11)) {
		goto L46
	} else {
		goto L251
	}
L244:
	;
	F_LockBuffer(m, v814, int32(1))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L8
	} else {
		goto L248
	}
L245:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[5]))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818+(v814^int32(-1))<<(uint(int32(2))%32))))
	v832 = v824
	goto L244
L246:
	;
	goto L247
L247:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[6]))
	v832 = v826 + v814<<(uint(int32(13))%32) + int32(-8192)
	goto L244
L248:
	;
	F_heap_get_root_tuples(m, v832, v33+int32(48))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L8
	} else {
		goto L249
	}
L249:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	F_LockBuffer(m, v840, int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L8
	} else {
		goto L250
	}
L250:
	;
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v810))))
	v845 = v844
	goto L243
L251:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+46)) = uint16(v845)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+42)) = v852
	m.T0[l8].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, l1, v33+int32(42), v33+int32(672), v33+int32(640), v751, l9)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L8
	} else {
		goto L252
	}
L252:
	;
	v900 = v773
	goto L84
L253:
	;
	v900 = v773
	goto L84
L254:
	;
	if v901 != 0 {
		v167 = v901
		v181 = v257
		v185 = v291
		v193 = v900
		goto L51
	} else {
		goto L255
	}
L255:
	;
	goto L52
L256:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v103)+32))
	if v934 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+188))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v978)+12))
	m.T0[v979].(func(*base.Module, int32))(m, v103)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L8
	} else {
		goto L266
	}
L259:
	;
	v939 = v934 + int32(20)
	goto L261
L260:
	;
	v939 = v103 + int32(36)
	goto L261
L261:
	;
	v940 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v939))))
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[1]))
	if v943 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L258
L263:
	;
	goto L262
L264:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[2])))
	if v947&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v952 = int32(_a_F_heapam_index_build_range_scan_2)
	v954 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	v955 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v954 + v955
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v958 + v955
	*(*int64)(unsafe.Add(mBase, uint32(v943+int32(128))+232)) = v940
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	*(*int32)(unsafe.Add(mBase, uint32(v943))) = v966 + v955
	v972 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_build_range_scan[3])) = v972 - v955
	goto L263
L266:
	;
	if v105 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_UnregisterSnapshot(m, v104)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L8
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_ExecDropSingleTupleTableSlot(m, v66)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L8
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	F_FreeExecutorState(m, v55)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L8
	} else {
		goto L272
	}
L272:
	;
	v988 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v988
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v988
	m.G0 = v33 + int32(800)
	return v932
L273:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L8
	} else {
		goto L274
	}
L274:
	;
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+6)))
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v1005 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1003 | v1004<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_build_range_scan_12), v33)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L8
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_heapam_index_build_range_scan_4), int32(1693), int32(_a_F_heapam_index_build_range_scan_5))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L8
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapam_index_fetch_begin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		return v4
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 float64
	_ = v108
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v277 int64
	_ = v277
	var v280 int64
	_ = v280
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v387 int32
	_ = v387
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 float64
	_ = v521
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	v6 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1088)
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
	v32 = F_MakeTupleTableSlot(m, v30, int32(_a_F_heapam_index_validate_scan_0))
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
	v39 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v45 = m.T0[v44].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3, v39, v39, v39, int32(321))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v45)+36)))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[0]))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v83 = F_heap_getnext(m, v45)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[1])))
	if v54&int32(1) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(_a_F_heapam_index_validate_scan_1)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v61 + v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v65 + v62
	*(*int64)(unsafe.Add(mBase, uint32(v50+int32(120))+232)) = v47
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v73 + v62
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v79 - v62
	goto L11
L14:
	;
	if v83 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = int32(-1)
	v90 = v83
	v95 = v85
	v96 = v6
	v99 = v85
	v102 = v6
	goto L18
L16:
	;
	goto L17
L17:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+188))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+12))
	m.T0[v563].(func(*base.Module, int32))(m, v45)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L99
	}
L18:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[3]))
	if v105 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v108 = *(*float64)(unsafe.Add(mBase, uint32(l4)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+8)) = base.F64_add(v108, float64(1))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	if base.B2i32(v112 == v99)&base.B2i32(v99 != int32(-1)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[0]))
	if v123 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v157 = v99
	goto L26
L26:
	;
	if v95 != v157 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v157 = v156
	goto L26
L28:
	;
	goto L27
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[1])))
	if v127&int32(1) == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v132 = int32(_a_F_heapam_index_validate_scan_1)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	v135 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v134 + v135
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v138 + v135
	*(*int64)(unsafe.Add(mBase, uint32(v123+int32(128))+232)) = base.I64_extend_i32_u(v112)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v146 + v135
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[2])) = v152 - v135
	goto L28
L31:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	if v159 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v195 = v95
	goto L33
L33:
	;
	v197 = v90 + int32(4)
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v200
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+8)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v203)+18)))
	if v204 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	F_LockBuffer(m, v159, int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[4]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163+(v159^int32(-1))<<(uint(int32(2))%32))))
	v177 = v169
	goto L34
L36:
	;
	goto L37
L37:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[5]))
	v177 = v171 + v159<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	F_heap_get_root_tuples(m, v177, v20+int32(336))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	F_LockBuffer(m, v185, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.MemoryFill(m, v20+int32(32), int32(0), int32(291))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v195 = v194
	goto L33
L41:
	;
	v542 = F_heap_getnext(m, v45)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L97
	}
L42:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v218)+31)))
	if v487 != 0 {
		v534 = v478
		v540 = v484
		goto L41
	} else {
		goto L87
	}
L43:
	;
	v440 = int32(0)
	v442 = v20 + int32(20)
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+2)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432))))
	v448 = int32(16)
	v450 = v446 | v447<<(uint(v448)%32)
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442)+2)))
	v452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442))))
	v455 = v451 | v452<<(uint(v448)%32)
	if base.Ui32(v450) < base.Ui32(v455) {
		v466 = int32(-1)
		goto L82
	} else {
		goto L83
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L77
	}
L45:
	;
	v207 = int32(1)
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202<<(uint(v207)%32)+v20)+334)))
	if base.Ui32(int32(2048)) <= base.Ui32((v210-v207)&int32(_a_F_heapam_index_validate_scan_2)) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v218 = v202
	goto L47
L47:
	;
	if v102 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v210)
	v218 = v210
	goto L47
L49:
	;
	if v96 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v387 = v96
	goto L51
L51:
	;
	v478 = v387
	v484 = int32(1)
	goto L42
L52:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v263 = int32(0)
	v269 = F_tuplesort_getdatum(m, v261, int32(1), v263, v20+int32(16), v20+int32(15), v263)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L62
	}
L53:
	;
	v224 = v20 + int32(20)
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+2)))
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
	v230 = int32(16)
	v232 = v228 | v229<<(uint(v230)%32)
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+2)))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224))))
	v237 = v233 | v234<<(uint(v230)%32)
	if base.Ui32(v232) < base.Ui32(v237) {
		v248 = int32(-1)
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if int32(0) <= v248 {
		v432 = v96
		goto L43
	} else {
		goto L59
	}
L55:
	;
	goto L54
L56:
	;
	if base.Ui32(v237) < base.Ui32(v232) {
		v248 = int32(1)
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)))
	if base.Ui32(v242) < base.Ui32(v243) {
		v248 = int32(-1)
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v248 = base.B2i32(base.Ui32(v243) < base.Ui32(v242))
	goto L55
L59:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+2)))
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96))))
	if v251|v252<<(uint(int32(16))%32) != v195 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+4)))
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v257)+31)) = uint8(v259)
	goto L52
L61:
	;
	v387 = int32(0)
	goto L51
L62:
	;
	if v269 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v273)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v274)
	v277 = int64(base.Ui64(v274) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v277)
	v280 = int64(base.Ui64(v274) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v280)
	goto L64
L64:
	;
	v300 = v20 + int32(26)
	v302 = v20 + int32(20)
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+2)))
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300))))
	v308 = int32(16)
	v310 = v306 | v307<<(uint(v308)%32)
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+2)))
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302))))
	v315 = v311 | v312<<(uint(v308)%32)
	if base.Ui32(v310) < base.Ui32(v315) {
		v326 = int32(-1)
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if int32(0) <= v326 {
		v432 = v300
		goto L43
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	if base.Ui32(v315) < base.Ui32(v310) {
		v326 = int32(1)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+4)))
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+4)))
	if base.Ui32(v320) < base.Ui32(v321) {
		v326 = int32(-1)
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v326 = base.B2i32(base.Ui32(v321) < base.Ui32(v320))
	goto L67
L71:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)))
	if v195 == v329|v330<<(uint(int32(16))%32) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)))
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v335)+31)) = uint8(v337)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v341 = int32(0)
	v347 = F_tuplesort_getdatum(m, v339, int32(1), v341, v20+int32(16), v20+int32(15), v341)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v347 == int32(0) {
		goto L61
	} else {
		goto L76
	}
L76:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v351)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+30)) = uint16(v352)
	v355 = int64(base.Ui64(v352) >> (uint(int64(16)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v355)
	v358 = int64(base.Ui64(v352) >> (uint(int64(32)) % 64))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+26)) = uint16(v358)
	goto L64
L77:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+4)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v405 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v403 | v404<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F_heapam_index_validate_scan_3), v20)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_heapam_index_validate_scan_4), int32(1870), int32(_a_F_heapam_index_validate_scan_5))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	if v466 <= int32(0) {
		v534 = v432
		v540 = v440
		goto L41
	} else {
		goto L86
	}
L82:
	;
	goto L81
L83:
	;
	if base.Ui32(v455) < base.Ui32(v450) {
		v466 = int32(1)
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+4)))
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v442)+4)))
	if base.Ui32(v460) < base.Ui32(v461) {
		v466 = int32(-1)
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v466 = base.B2i32(base.Ui32(v461) < base.Ui32(v460))
	goto L82
L86:
	;
	v478 = v432
	v484 = v440
	goto L42
L87:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	F_MemoryContextReset(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v492 = F_ExecStoreHeapTuple(m, v90, v32, int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v36 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v494 = int32(_a_F_heapam_index_validate_scan_6)
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6])) = v497
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v502 = m.T0[v501].(func(*base.Module, int32, int32, int32) int32)(m, v36, v29, v20+int32(16))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v510 = v20 + int32(960)
	v512 = v20 + int32(928)
	F_FormIndexDatum(m, l2, v32, v22, v510, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_index_validate_scan[6])) = v495
	if v502 == int32(0) {
		v534 = v478
		v540 = v484
		goto L41
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+116)))
	v519 = F_index_insert(m, l1, v510, v512, v20+int32(20), l0, v517, int32(0), l2)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v521 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l4)+24)) = base.F64_add(v521, float64(1))
	v534 = v478
	v540 = v484
	goto L41
L97:
	;
	if v542 != 0 {
		v90 = v542
		v95 = v195
		v96 = v534
		v99 = v157
		v102 = v540
		goto L18
	} else {
		goto L98
	}
L98:
	;
	goto L19
L99:
	;
	F_ExecDropSingleTupleTableSlot(m, v32)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_FreeExecutorState(m, v22)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+88)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v570
	m.G0 = v20 + int32(1088)
	return
}
func F_heapam_relation_needs_toast_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 <= v2 {
		v146 = v2
	} else {
		v14 = int32(0)
		v15 = v2
		v16 = v10
		v18 = v2
		v20 = v2
		for {
			v27 = v9 + v16<<(uint(int32(4))%32) + v14*int32(100)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+111)))
			if v28 != 0 {
				v113 = v15
				v114 = v16
				v116 = v18
				v117 = v20
			} else {
				v30 = v27 + int32(20)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+90)))
				if v31 == int32(118) {
					v113 = v15
					v114 = v16
					v116 = v18
					v117 = v20
				} else {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+83)))
					switch v34 - int32(99) {
					case 0:
						v49 = v15
					case 1:
						v49 = (v15 + int32(7)) & int32(-8)
					default:
						v49 = (v15 + int32(1)) & int32(-2)
					case 6:
						v49 = (v15 + int32(3)) & int32(-4)
					}
					v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30)+72)))
					if int32(0) < v50 {
						v113 = v49 + v50
						v114 = v16
						v116 = v18
						v117 = v20
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
						v57 = int32(-1)
						if v55 < int32(0) {
							v97 = v57
							v99 = v97
						} else {
							if base.Ui32(int32(2)) <= base.Ui32(v54-int32(1042)) {
								switch v54 - int32(1560) {
								case 0, 2:
									v93 = int32(8)
									v94 = base.I32_div_s(v55+int32(7), v93)
									v97 = v94 + v93
									v99 = v97
								case 1:
									v97 = v57
									v99 = v97
								default:
									if v54 != int32(1700) {
										v97 = v57
										v99 = v97
									} else {
										v76 = int32(4)
										if v55 < v76 {
											v90 = int32(-1)
										} else {
											v90 = int32(base.Ui32(int32(base.Ui32(v55-v76)>>(uint(int32(16))%32))+int32(6))>>(uint(int32(1))%32))&int32(_a_F_heapam_relation_needs_toast_table_0) + int32(8)
										}
										v99 = v90
									}
								}
							} else {
								v66 = F_GetDatabaseEncoding(m)
								mBase = m.M
								v67 = F_pg_encoding_max_length(m, v66)
								mBase = m.M
								v68 = int32(4)
								v99 = v67*(v55-v68) + v68
							}
						}
						v100 = int32(0)
						if v100 < v99 {
							v103 = v99
						} else {
							v103 = v100
						}
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+84)))
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v113 = v103 + v49
						v114 = v112
						v116 = base.B2i32(v108 != int32(112)) | v18
						v117 = base.B2i32(v99 < int32(0)) | v20
					}
				}
			}
			v120 = v14 + int32(1)
			if v120 < v114 {
				v14 = v120
				v15 = v113
				v16 = v114
				v18 = v116
				v20 = v117
				continue
			} else {
				break
			}
			break
		}
		if (v116^int32(-1)|v117)&int32(1) != 0 {
			v146 = v116
		} else {
			v127 = int32(7)
			v130 = base.I32_div_s(v114+v127, int32(8))
			v133 = int32(-8)
			v146 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32((v130+int32(30))&v133+(v113+v127)&v133))
		}
	}
	return v146 & int32(1)
}
func F_heapam_scan_analyze_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v354 float64
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 float64
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L20
	} else {
		goto L116
	}
L2:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v33) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(2))%32))))
	v32 = v24
	goto L2
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[1]))
	v32 = v26 + v14<<(uint(int32(13))%32) + int32(-8192)
	goto L2
L6:
	;
	v43 = int32(base.Ui32(v33+int32(_a_F_heapam_scan_analyze_next_tuple_0))>>(uint(int32(2))%32)) & int32(_a_F_heapam_scan_analyze_next_tuple_1)
	goto L8
L7:
	;
	v43 = int32(0)
	goto L8
L8:
	;
	if base.Ui32(v13) <= base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = l4 + int32(48)
	v54 = v13
	goto L12
L10:
	;
	v393 = v14
	goto L11
L11:
	;
	F_UnlockReleaseBuffer(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L114
	}
L12:
	;
	v63 = v32 + int32(20) + v54<<(uint(int32(2))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	switch int32(base.Ui32(v64)>>(uint(int32(15))%32))&int32(3) - int32(1) {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L13:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v393 = v380
	goto L11
L14:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v377 = v375 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v377
	if base.Ui32(v377) <= base.Ui32(v43) {
		v54 = v377
		goto L12
	} else {
		goto L113
	}
L15:
	;
	v369 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v369, float64(1))
	goto L14
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+56)) = uint16(v54)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+54)) = uint16(v71)
	v75 = int32(base.Ui32(v71) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+52)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+60)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+64)) = v32 + v80&int32(_a_F_heapam_scan_analyze_next_tuple_2)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+48)) = int32(base.Ui32(v85) >> (uint(int32(17)) % 32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v90 = F_HeapTupleSatisfiesVacuum(m, v48, l1, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v354 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = base.F64_add(v354, float64(1))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v48, l4, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L20
	} else {
		goto L112
	}
L18:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223)+20)))
	if v224&int32(_a_F_heapam_scan_analyze_next_tuple_3) == int32(_a_F_heapam_scan_analyze_next_tuple_4) {
		goto L67
	} else {
		goto L68
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l4)+64))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+20)))
	v96 = int32(768)
	if v95&v96 != v96 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return int32(0)
L21:
	;
	switch v90 {
	case 0, 2:
		goto L15
	case 1:
		goto L17
	case 3:
		goto L19
	case 4:
		goto L18
	default:
		goto L1
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = v100
	goto L24
L23:
	;
	v102 = int32(2)
	goto L24
L24:
	;
	if base.Ui32(v102) < base.Ui32(int32(3)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v222 != 0 {
		goto L17
	} else {
		goto L65
	}
L26:
	;
	v222 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v113 == v102 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v222 = int32(1)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v117 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v222 = v214
	goto L25
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v121 == int32(0) {
		v214 = int32(0)
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v185 = int32(0)
	v187 = v117 - int32(1)
	goto L55
L36:
	;
	v126 = v121
	goto L37
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	if v131 == int32(4) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v214 = int32(0)
	goto L32
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v126)+80))
	if v178 != 0 {
		v126 = v178
		goto L37
	} else {
		goto L54
	}
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v134 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v137 = int32(1)
	if v102 == v134 {
		v214 = v137
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+52))
	v141 = v139 - int32(1)
	if v141 < int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v146 = int32(0)
	v148 = v141
	goto L44
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v154 = int32(2)
	v155 = base.I32_div_s(v148-v146, v154)
	v156 = v155 + v146
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v152+v156<<(uint(v154)%32))))
	if v160 == v102 {
		v214 = v137
		goto L32
	} else {
		goto L46
	}
L45:
	;
	goto L39
L46:
	;
	v164 = F_TransactionIdPrecedes(m, v160, v102)
	mBase = m.M
	if v164 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v165 = v156 + int32(1)
	goto L49
L48:
	;
	v165 = v146
	goto L49
L49:
	;
	if v164 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v168 = v148
	goto L52
L51:
	;
	v168 = v156 - int32(1)
	goto L52
L52:
	;
	if v165 <= v168 {
		v146 = v165
		v148 = v168
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L45
L54:
	;
	goto L38
L55:
	;
	v192 = int32(2)
	v193 = base.I32_div_s(v187-v185, v192)
	v194 = v193 + v185
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183+v194<<(uint(v192)%32))))
	v199 = base.B2i32(v198 == v102)
	if v198 == v102 {
		v214 = v199
		goto L32
	} else {
		goto L57
	}
L56:
	;
	v214 = v199
	goto L32
L57:
	;
	v202 = base.B2i32(base.Ui32(v198) < base.Ui32(v102))
	if base.Ui32(v198) < base.Ui32(v102) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v203 = v194 + int32(1)
	goto L60
L59:
	;
	v203 = v185
	goto L60
L60:
	;
	if base.Ui32(v198) < base.Ui32(v102) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v206 = v187
	goto L63
L62:
	;
	v206 = v194 - int32(1)
	goto L63
L63:
	;
	if v203 <= v206 {
		v185 = v203
		v187 = v206
		goto L55
	} else {
		goto L64
	}
L64:
	;
	goto L56
L65:
	;
	goto L14
L66:
	;
	if base.Ui32(v232) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v229 = F_HeapTupleGetUpdateXid(m, v223)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v232 = v231
	goto L66
L70:
	;
	v232 = v229
	goto L66
L71:
	;
	if v352 != 0 {
		goto L15
	} else {
		goto L111
	}
L72:
	;
	v352 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[2]))
	if v243 == v232 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v352 = int32(1)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[3]))
	if v247 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v352 = v344
	goto L71
L79:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[4]))
	if v251 == int32(0) {
		v344 = int32(0)
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_tuple[5]))
	v315 = int32(0)
	v317 = v247 - int32(1)
	goto L101
L82:
	;
	v256 = v251
	goto L83
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	if v261 == int32(4) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v344 = int32(0)
	goto L78
L85:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v256)+80))
	if v308 != 0 {
		v256 = v308
		goto L83
	} else {
		goto L100
	}
L86:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v264 == int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v267 = int32(1)
	if v232 == v264 {
		v344 = v267
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v256)+52))
	v271 = v269 - int32(1)
	if v271 < int32(0) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v276 = int32(0)
	v278 = v271
	goto L90
L90:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	v284 = int32(2)
	v285 = base.I32_div_s(v278-v276, v284)
	v286 = v285 + v276
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v282+v286<<(uint(v284)%32))))
	if v290 == v232 {
		v344 = v267
		goto L78
	} else {
		goto L92
	}
L91:
	;
	goto L85
L92:
	;
	v294 = F_TransactionIdPrecedes(m, v290, v232)
	mBase = m.M
	if v294 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v295 = v286 + int32(1)
	goto L95
L94:
	;
	v295 = v276
	goto L95
L95:
	;
	if v294 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v298 = v278
	goto L98
L97:
	;
	v298 = v286 - int32(1)
	goto L98
L98:
	;
	if v295 <= v298 {
		v276 = v295
		v278 = v298
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L91
L100:
	;
	goto L84
L101:
	;
	v322 = int32(2)
	v323 = base.I32_div_s(v317-v315, v322)
	v324 = v323 + v315
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v313+v324<<(uint(v322)%32))))
	v329 = base.B2i32(v328 == v232)
	if v328 == v232 {
		v344 = v329
		goto L78
	} else {
		goto L103
	}
L102:
	;
	v344 = v329
	goto L78
L103:
	;
	v332 = base.B2i32(base.Ui32(v328) < base.Ui32(v232))
	if base.Ui32(v328) < base.Ui32(v232) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v333 = v324 + int32(1)
	goto L106
L105:
	;
	v333 = v315
	goto L106
L106:
	;
	if base.Ui32(v328) < base.Ui32(v232) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v336 = v317
	goto L109
L108:
	;
	v336 = v324 - int32(1)
	goto L109
L109:
	;
	if v333 <= v336 {
		v315 = v333
		v317 = v336
		goto L101
	} else {
		goto L110
	}
L110:
	;
	goto L102
L111:
	;
	goto L17
L112:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v362 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v361 + v362
	return v362
L113:
	;
	goto L13
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	m.T0[v399].(func(*base.Module, int32))(m, l4)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	return int32(0)
L116:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_scan_analyze_next_tuple_5), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_heapam_scan_analyze_next_tuple_6), int32(1147), int32(_a_F_heapam_scan_analyze_next_tuple_7))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
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
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	v4 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v28, int32(1))
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v68 = l0 - int32(-64)
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
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ExecStoreBufferHeapTuple(m, v68, l2, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v52 + v116&int32(_a_F_heapam_scan_sample_next_tuple_4)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v105)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v21)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v64)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v125) >> (uint(int32(17)) % 32))
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
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
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(108)+v165<<(uint(v163)%32)))))
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
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v208, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L23
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v214, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	m.T0[v219].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
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
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, int32(0), v223, v68, v224, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+272))
	if v234 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+268)))
	if v237 != int32(1) {
		goto L16
	} else {
		goto L64
	}
L62:
	;
	v244 = v234
	goto L63
L63:
	;
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v244)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v244)+24)) = v245 + int64(1)
	goto L16
L64:
	;
	F_pgstat_assoc_relation(m, v233)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+272))
	v244 = v243
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
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_LockBuffer(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
		v12 = F_HeapTupleSatisfiesVisibility(m, v10, l2, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			F_LockBuffer(m, v14, int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
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
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v16 = base.B2i32(base.Ui32(v10|v11<<(uint(int32(16))%32)) < base.Ui32(v9))
		}
	}
	return v16
}
