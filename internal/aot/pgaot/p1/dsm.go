package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_impl_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1086 int32
	_ = v1086
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(496)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[0]))
	switch v19 - int32(1) {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L1
	case 3:
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L6
	} else {
		goto L356
	}
L2:
	;
	m.G0 = v16 + int32(496)
	return v1086
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+320)) = l1
	v753 = F_pg_snprintf(m, v16+int32(432), int32(64), int32(_a_F_dsm_impl_op_0), v16+int32(320))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L6
	} else {
		goto L241
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = l1
	v479 = F_pg_snprintf(m, v16+int32(432), int32(64), int32(_a_F_dsm_impl_op_1), v16+int32(192))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L6
	} else {
		goto L142
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l1
	v29 = F_pg_snprintf(m, v16+int32(432), int32(64), int32(_a_F_dsm_impl_op_2), v16+int32(112))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if l0&int32(-2) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_ReserveExternalFD(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L30
	}
L11:
	;
	goto L13
L12:
	;
	goto L13
L13:
	;
	v39 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v39
	v43 = int32(1)
	if l0 != int32(3) {
		v1086 = v43
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v49 = m.G0
	v51 = v49 - int32(272)
	m.G0 = v51
	v53 = F___shm_mapname(m, v16+int32(432), v51)
	mBase = m.M
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v56 == int32(0) {
		v1086 = v43
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v54 = F_unlink(m, v53)
	mBase = m.M
	v56 = v54
	goto L18
L17:
	;
	v56 = int32(-1)
	goto L18
L18:
	;
	m.G0 = v51 + int32(272)
	goto L15
L19:
	;
	v62 = int32(0)
	v64 = F_errstart(m, l6, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v64 == int32(0) {
		v1086 = v62
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v69 != int32(48))&base.B2i32(v69 != int32(22)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_3), v16+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L27
	}
L26:
	;
	goto L22
L27:
	;
	goto L22
L28:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(242), int32(_a_F_dsm_impl_op_6))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v1086 = v62
	goto L2
L30:
	;
	if l0 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v99 = int32(2)
	goto L33
L32:
	;
	v99 = int32(194)
	goto L33
L33:
	;
	v100 = m.G0
	v102 = v100 - int32(288)
	m.G0 = v102
	v107 = v102 + int32(16)
	v110 = v16 + int32(432)
	goto L36
L34:
	;
	m.G0 = v102 + int32(288)
	if v164 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L35:
	;
	if v154 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v116 == int32(47) {
		v110 = v110 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v120 = F___strchrnul(m, v110, int32(47))
	mBase = m.M
	if v120 == v110 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L37
L39:
	;
	v147 = F___memcpy(m, v107, int32(_a_F_dsm_impl_op_7), int32(9))
	mBase = m.M
	v152 = F___memcpy(m, v102+int32(25), v110, v123+int32(1))
	mBase = m.M
	v154 = v107
	goto L35
L40:
	;
	if base.Ui32(v123) < base.Ui32(int32(256)) {
		goto L39
	} else {
		goto L47
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(28)
	v154 = int32(0)
	goto L35
L42:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v122 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v123 = v120 - v110
	if int32(2) < v123 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v126 != int32(46) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120-int32(1)))))
	if v131 != int32(46) {
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(37)
	v154 = int32(0)
	goto L35
L48:
	;
	v164 = int32(-1)
	goto L34
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(384)
	v162 = F_open(m, v154, v99|int32(_a_F_dsm_impl_op_8), v102)
	mBase = m.M
	v164 = v162
	goto L34
L51:
	;
	v170 = int32(_a_F_dsm_impl_op_9)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2])) = v172 - int32(1)
	goto L54
L52:
	;
	goto L53
L53:
	;
	if l0 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L54:
	;
	if l0 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if v179 == int32(20) {
		v1086 = v8
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v183 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	if v183 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v188 != int32(48))&base.B2i32(v188 != int32(22)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_10), v16+int32(48))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L67
	}
L62:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L66
	}
L65:
	;
	goto L61
L66:
	;
	goto L61
L67:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(266), int32(_a_F_dsm_impl_op_6))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v1086 = v8
	goto L2
L69:
	;
	v398 = int32(1)
	v400 = F_mmap(m, v387, v398, v164)
	mBase = m.M
	if v400 == int32(-1) {
		goto L120
	} else {
		goto L121
	}
L70:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
	v387 = v384
	goto L69
L71:
	;
	if v164 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_impl_op[3])))
	if v272 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L74:
	;
	if v225 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L75:
	;
	v221 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v225 = v221
	goto L74
L76:
	;
	goto L77
L77:
	;
	v224 = F___fstatat(m, v164, int32(_a_F_dsm_impl_op_11), v16+int32(336), int32(_a_F_dsm_impl_op_12))
	mBase = m.M
	v225 = v224
	goto L74
L78:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v230 = F_close(m, v164)
	mBase = m.M
	v231 = int32(_a_F_dsm_impl_op_9)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2])) = v233 - int32(1)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v229
	v240 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	if v240 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v245 != int32(48))&base.B2i32(v245 != int32(22)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_13), v16+int32(80))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L88
	}
L83:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	goto L82
L88:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(291), int32(_a_F_dsm_impl_op_6))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	v1086 = v8
	goto L2
L90:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_dsm_impl_op_14), v16+int32(336))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = int32(167772185)
	goto L94
L93:
	;
	goto L92
L94:
	;
	v298 = F_ftruncate(m, v164, base.I64_extend_i32_u(l2))
	mBase = m.M
	if v298 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = int32(0)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsm_impl_op[3])))
	if v310 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if v302 == int32(27) {
		goto L94
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	goto L98
L100:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	F_pgmem_sigprocmask(m, v16+int32(336), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v298 == int32(0) {
		v387 = l2
		goto L69
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v314
	goto L102
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v327 = F_close(m, v164)
	mBase = m.M
	v328 = int32(_a_F_dsm_impl_op_9)
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2])) = v330 - int32(1)
	goto L105
L105:
	;
	v337 = m.G0
	v339 = v337 - int32(272)
	m.G0 = v339
	v341 = F___shm_mapname(m, v16+int32(432), v339)
	mBase = m.M
	if v341 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v326
	v350 = int32(0)
	v352 = F_errstart(m, l6, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L110
	}
L107:
	;
	v342 = F_unlink(m, v341)
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	m.G0 = v339 + int32(272)
	goto L106
L110:
	;
	if v352 == int32(0) {
		v1086 = v350
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v357 != int32(48))&base.B2i32(v357 != int32(22)) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_15), v16+int32(96))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L118
	}
L113:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L6
	} else {
		goto L117
	}
L116:
	;
	goto L112
L117:
	;
	goto L112
L118:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(310), int32(_a_F_dsm_impl_op_6))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	v1086 = v350
	goto L2
L120:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v405 = F_close(m, v164)
	mBase = m.M
	v406 = int32(_a_F_dsm_impl_op_9)
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2])) = v408 - int32(1)
	goto L123
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v400
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v387
	v465 = F_close(m, v164)
	mBase = m.M
	v466 = int32(_a_F_dsm_impl_op_9)
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[2])) = v468 - int32(1)
	goto L141
L123:
	;
	if l0 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v417 = m.G0
	v419 = v417 - int32(272)
	m.G0 = v419
	v421 = F___shm_mapname(m, v16+int32(432), v419)
	mBase = m.M
	if v421 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v404
	v430 = int32(0)
	v432 = F_errstart(m, l6, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L131
	}
L127:
	;
	goto L126
L128:
	;
	v422 = F_unlink(m, v421)
	mBase = m.M
	goto L130
L129:
	;
	goto L130
L130:
	;
	m.G0 = v419 + int32(272)
	goto L127
L131:
	;
	if v432 == int32(0) {
		v1086 = v430
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v437 != int32(48))&base.B2i32(v437 != int32(22)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_16), v16-int32(-64))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L6
	} else {
		goto L139
	}
L134:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L138
	}
L137:
	;
	goto L133
L138:
	;
	goto L133
L139:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(332), int32(_a_F_dsm_impl_op_6))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	v1086 = v430
	goto L2
L141:
	;
	v1086 = v398
	goto L2
L142:
	;
	if l1 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if l0 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v503 != 0 {
		goto L153
	} else {
		goto L154
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(20)
	v1086 = v8
	goto L2
L147:
	;
	v487 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	if v487 == int32(0) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	F_errmsg_internal(m, int32(_a_F_dsm_impl_op_17), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(470), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	goto L146
L152:
	;
	if l0&int32(-2) == int32(2) {
		goto L178
	} else {
		goto L179
	}
L153:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v563 = v504
	v564 = v503
	goto L152
L154:
	;
	goto L155
L155:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[5]))
	v508 = F_MemoryContextAlloc(m, v506, int32(4))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v511 = l1 >> (uint(int32(31)) % 32)
	if l0 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v515 = int32(0)
	goto L159
L158:
	;
	v515 = l2
	goto L159
L159:
	;
	if l0 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v518 = int32(384)
	goto L162
L161:
	;
	v518 = int32(1920)
	goto L162
L162:
	;
	v519 = F_pgmem_shmget(m, l1^v511-v511, v515, v518)
	mBase = m.M
	if v519 == int32(-1) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v523 == int32(20))&base.B2i32(l0 != int32(1)) != 0 {
		v1086 = v8
		goto L2
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = v519
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v508
	v563 = v519
	v564 = v508
	goto L152
L166:
	;
	F_pfree(m, v508)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v523
	v534 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v534 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L169
	}
L169:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v539 != int32(48))&base.B2i32(v539 != int32(22)) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	F_errmsg(m, int32(_a_F_dsm_impl_op_19), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L6
	} else {
		goto L176
	}
L171:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L6
	} else {
		goto L175
	}
L174:
	;
	goto L170
L175:
	;
	goto L170
L176:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(519), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	v1086 = v8
	goto L2
L178:
	;
	F_pfree(m, v564)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L6
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if l0 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L181:
	;
	v572 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v572
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v574 == v572 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v612
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v612
	v616 = int32(1)
	if l0 != int32(3) {
		v1086 = v616
		goto L2
	} else {
		goto L195
	}
L183:
	;
	v577 = F_pgmem_shmdt(m, v574)
	mBase = m.M
	if v577 == int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v581 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	if v581 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L186
	}
L186:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v586 != int32(48))&base.B2i32(v586 != int32(22)) == int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_20), v16+int32(144))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L6
	} else {
		goto L193
	}
L188:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L192
	}
L191:
	;
	goto L187
L192:
	;
	goto L187
L193:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(538), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	v1086 = v8
	goto L2
L195:
	;
	v619 = int32(0)
	v621 = F_pgmem_shmctl(m, v563, v619, v619)
	mBase = m.M
	if v619 <= v621 {
		v1086 = v616
		goto L2
	} else {
		goto L196
	}
L196:
	;
	v625 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	if v625 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1086 = int32(0)
	goto L2
L199:
	;
	goto L200
L200:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v631 != int32(48))&base.B2i32(v631 != int32(22)) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_3), v16+int32(128))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L6
	} else {
		goto L207
	}
L202:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L6
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L6
	} else {
		goto L206
	}
L205:
	;
	goto L201
L206:
	;
	goto L201
L207:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(548), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L208
	}
L208:
	;
	v1086 = int32(0)
	goto L2
L209:
	;
	v663 = F_pgmem_shmctl(m, v563, int32(2), v16+int32(336))
	mBase = m.M
	if v663 != 0 {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	v697 = l2
	goto L211
L211:
	;
	v699 = F_pgmem_shmat(m, v563, int32(0))
	mBase = m.M
	if v699 == int32(-1) {
		goto L225
	} else {
		goto L226
	}
L212:
	;
	v665 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L6
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v16)+372))
	v697 = v696
	goto L211
L215:
	;
	if v665 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L216
	}
L216:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v670 != int32(48))&base.B2i32(v670 != int32(22)) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_13), v16+int32(176))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L223
	}
L218:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L6
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L6
	} else {
		goto L222
	}
L221:
	;
	goto L217
L222:
	;
	goto L217
L223:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(564), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	v1086 = v8
	goto L2
L225:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if l0 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v699
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v697
	v1086 = int32(1)
	goto L2
L228:
	;
	v706 = int32(0)
	v708 = F_pgmem_shmctl(m, v563, v706, v706)
	mBase = m.M
	goto L230
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v703
	v712 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L6
	} else {
		goto L231
	}
L231:
	;
	if v712 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L232
	}
L232:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v717 != int32(48))&base.B2i32(v717 != int32(22)) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_16), v16+int32(160))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L6
	} else {
		goto L239
	}
L234:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L238
	}
L237:
	;
	goto L233
L238:
	;
	goto L233
L239:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(585), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	v1086 = v8
	goto L2
L241:
	;
	if l0&int32(-2) == int32(2) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v759 != 0 {
		goto L245
	} else {
		goto L246
	}
L243:
	;
	goto L244
L244:
	;
	if l0 != 0 {
		goto L260
	} else {
		goto L261
	}
L245:
	;
	goto L247
L246:
	;
	goto L247
L247:
	;
	v761 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v761
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v761
	v765 = int32(1)
	if l0 != int32(3) {
		v1086 = v765
		goto L2
	} else {
		goto L248
	}
L248:
	;
	v770 = F_unlink(m, v16+int32(432))
	mBase = m.M
	if v770 == int32(0) {
		v1086 = v765
		goto L2
	} else {
		goto L249
	}
L249:
	;
	v773 = int32(0)
	v775 = F_errstart(m, l6, v773)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L6
	} else {
		goto L250
	}
L250:
	;
	if v775 == int32(0) {
		v1086 = v773
		goto L2
	} else {
		goto L251
	}
L251:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v780 != int32(48))&base.B2i32(v780 != int32(22)) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+208)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_3), v16+int32(208))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L258
	}
L253:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
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
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(823), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L6
	} else {
		goto L259
	}
L259:
	;
	v1086 = v773
	goto L2
L260:
	;
	v810 = int32(2)
	goto L262
L261:
	;
	v810 = int32(194)
	goto L262
L262:
	;
	v811 = F_OpenTransientFile(m, v16+int32(432), v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	if v811 == int32(-1) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	if l0 != int32(1) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	if l0 == int32(1) {
		goto L283
	} else {
		goto L284
	}
L267:
	;
	v818 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if v818 == int32(20) {
		v1086 = v8
		goto L2
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v822 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	if v822 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L272
	}
L272:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v827 != int32(48))&base.B2i32(v827 != int32(22)) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_10), v16+int32(240))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L6
	} else {
		goto L279
	}
L274:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L6
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L6
	} else {
		goto L278
	}
L277:
	;
	goto L273
L278:
	;
	goto L273
L279:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(837), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L6
	} else {
		goto L280
	}
L280:
	;
	v1086 = v8
	goto L2
L281:
	;
	v1002 = int32(1)
	v1004 = F_mmap(m, v991, v1002, v811)
	mBase = m.M
	if v1004 == int32(-1) {
		goto L332
	} else {
		goto L333
	}
L282:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
	v991 = v988
	goto L281
L283:
	;
	if v811 < int32(0) {
		goto L287
	} else {
		goto L288
	}
L284:
	;
	goto L285
L285:
	;
	v906 = F_palloc0(m, int32(_a_F_dsm_impl_op_22))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L302
	}
L286:
	;
	if v864 == int32(0) {
		goto L282
	} else {
		goto L290
	}
L287:
	;
	v860 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v864 = v860
	goto L286
L288:
	;
	goto L289
L289:
	;
	v863 = F___fstatat(m, v811, int32(_a_F_dsm_impl_op_11), v16+int32(336), int32(_a_F_dsm_impl_op_12))
	mBase = m.M
	v864 = v863
	goto L286
L290:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v869 = F_CloseTransientFile(m, v811)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L6
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v868
	v874 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L6
	} else {
		goto L292
	}
L292:
	;
	if v874 == int32(0) {
		v1086 = v8
		goto L2
	} else {
		goto L293
	}
L293:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v879 != int32(48))&base.B2i32(v879 != int32(22)) == int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+288)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_13), v16+int32(288))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L6
	} else {
		goto L300
	}
L295:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L6
	} else {
		goto L299
	}
L298:
	;
	goto L294
L299:
	;
	goto L294
L300:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(861), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L6
	} else {
		goto L301
	}
L301:
	;
	v1086 = v8
	goto L2
L302:
	;
	if l2 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v991 = int32(0)
	goto L281
L304:
	;
	goto L305
L305:
	;
	v912 = l2
	goto L306
L306:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v925))) = int32(167772186)
	v928 = int32(_a_F_dsm_impl_op_22)
	if base.Ui32(v928) <= base.Ui32(v912) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	if v931 == v932 {
		v991 = l2
		goto L281
	} else {
		goto L317
	}
L308:
	;
	goto L307
L309:
	;
	v931 = v928
	goto L311
L310:
	;
	v931 = v912
	goto L311
L311:
	;
	v932 = F_write(m, v811, v906, v931)
	mBase = m.M
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v934))) = int32(0)
	v937 = base.B2i32(v931 == v932)
	if v931 != v932 {
		goto L308
	} else {
		goto L312
	}
L312:
	;
	if v931 == v932 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v940 = v931
	goto L315
L314:
	;
	v940 = int32(0)
	goto L315
L315:
	;
	v941 = v912 - v940
	if v941 != 0 {
		v912 = v941
		goto L306
	} else {
		goto L316
	}
L316:
	;
	goto L308
L317:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v945 = F_CloseTransientFile(m, v811)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L6
	} else {
		goto L318
	}
L318:
	;
	v949 = F_unlink(m, v16+int32(432))
	mBase = m.M
	if v944 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v952 = v944
	goto L321
L320:
	;
	v952 = int32(51)
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v952
	v954 = int32(0)
	v956 = F_errstart(m, l6, v954)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L6
	} else {
		goto L322
	}
L322:
	;
	if v956 == int32(0) {
		v1086 = v954
		goto L2
	} else {
		goto L323
	}
L323:
	;
	v961 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v961 != int32(48))&base.B2i32(v961 != int32(22)) == int32(0) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+308)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+304)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_15), v16+int32(304))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L6
	} else {
		goto L330
	}
L325:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L6
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L6
	} else {
		goto L329
	}
L328:
	;
	goto L324
L329:
	;
	goto L324
L330:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(912), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L6
	} else {
		goto L331
	}
L331:
	;
	v1086 = v954
	goto L2
L332:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v1009 = F_CloseTransientFile(m, v811)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L6
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1004
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v991
	v1053 = F_CloseTransientFile(m, v811)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L6
	} else {
		goto L349
	}
L335:
	;
	if l0 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1015 = F_unlink(m, v16+int32(432))
	mBase = m.M
	goto L338
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v1008
	v1018 = int32(0)
	v1020 = F_errstart(m, l6, v1018)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L6
	} else {
		goto L339
	}
L339:
	;
	if v1020 == int32(0) {
		v1086 = v1018
		goto L2
	} else {
		goto L340
	}
L340:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v1025 != int32(48))&base.B2i32(v1025 != int32(22)) == int32(0) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+256)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_16), v16+int32(256))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L6
	} else {
		goto L347
	}
L342:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L6
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L6
	} else {
		goto L346
	}
L345:
	;
	goto L341
L346:
	;
	goto L341
L347:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(934), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L6
	} else {
		goto L348
	}
L348:
	;
	v1086 = v1018
	goto L2
L349:
	;
	if v1053 == int32(0) {
		v1086 = v1002
		goto L2
	} else {
		goto L350
	}
L350:
	;
	v1057 = int32(0)
	v1059 = F_errstart(m, l6, v1057)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L6
	} else {
		goto L351
	}
L351:
	;
	if v1059 == int32(0) {
		v1086 = v1057
		goto L2
	} else {
		goto L352
	}
L352:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L6
	} else {
		goto L353
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+272)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_23), v16+int32(272))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L6
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(945), int32(_a_F_dsm_impl_op_21))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L6
	} else {
		goto L355
	}
L355:
	;
	v1086 = v1057
	goto L2
L356:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v1100
	F_errmsg_internal(m, int32(_a_F_dsm_impl_op_24), v16)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L6
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(191), int32(_a_F_dsm_impl_op_25))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L6
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
