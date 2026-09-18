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
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1199 int32
	_ = v1199
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
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
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L6
	} else {
		goto L400
	}
L2:
	;
	m.G0 = v16 + int32(496)
	return v1199
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+320)) = l1
	v866 = F_pg_snprintf(m, v16+int32(432), int32(64), int32(_a_F_dsm_impl_op_0), v16+int32(320))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L6
	} else {
		goto L285
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
		v1199 = v43
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
		v1199 = v43
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
		v1199 = v62
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
	v1199 = v62
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
		v1199 = v8
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
		v1199 = v8
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
	v1199 = v8
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
		v1199 = v8
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
	v1199 = v8
	goto L2
L90:
	;
	F_sigprocmask(m, int32(_a_F_dsm_impl_op_14), v16+int32(336))
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
	F_sigprocmask(m, v16+int32(336), int32(0))
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
		v1199 = v350
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
	v1199 = v350
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
		v1199 = v430
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
	v1199 = v430
	goto L2
L141:
	;
	v1199 = v398
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
	v1199 = v8
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
		goto L200
	} else {
		goto L201
	}
L153:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v638 = v504
	v639 = v503
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
	v513 = l1 ^ v511 - v511
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
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[6]))
	if v523 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v594 == int32(-1) {
		goto L185
	} else {
		goto L186
	}
L164:
	;
	v527 = v523
	goto L167
L165:
	;
	goto L166
L166:
	;
	if v518&int32(512) != 0 {
		goto L173
	} else {
		goto L174
	}
L167:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	if v513 == v530 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v594 = v532
	goto L163
L170:
	;
	goto L171
L171:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v527)+20))
	if v533 != 0 {
		v527 = v533
		goto L167
	} else {
		goto L172
	}
L172:
	;
	goto L168
L173:
	;
	v546 = int32(_a_F_dsm_impl_op_19)
	goto L176
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(44)
	v594 = int32(-1)
	goto L163
L176:
	;
	if base.Ui32(v546) < base.Ui32(v515) {
		v546 = v546 << (uint(int32(1)) % 32)
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v552 = F_emscripten_builtin_malloc(m, v546)
	mBase = m.M
	if v552 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(48)
	v594 = int32(-1)
	goto L163
L180:
	;
	goto L181
L181:
	;
	v560 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v560 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	F_emscripten_builtin_free(m, v552)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(48)
	v594 = int32(-1)
	goto L163
L183:
	;
	goto L184
L184:
	;
	v568 = int32(_a_F_dsm_impl_op_20)
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[7])) = v570 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v560)+16)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v560)+12)) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v560)+8)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v570
	v579 = int32(_a_F_dsm_impl_op_21)
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+20)) = v580
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[6])) = v560
	v594 = v570
	goto L163
L185:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v598 == int32(20))&base.B2i32(l0 != int32(1)) != 0 {
		v1199 = v8
		goto L2
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = v594
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v508
	v638 = v594
	v639 = v508
	goto L152
L188:
	;
	F_pfree(m, v508)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v598
	v609 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	if v609 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L191
	}
L191:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v614 != int32(48))&base.B2i32(v614 != int32(22)) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	F_errmsg(m, int32(_a_F_dsm_impl_op_22), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L6
	} else {
		goto L198
	}
L193:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L197
	}
L196:
	;
	goto L192
L197:
	;
	goto L192
L198:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(519), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v1199 = v8
	goto L2
L200:
	;
	F_pfree(m, v639)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L6
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if l0 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L203:
	;
	v647 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v647
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v649 == v647 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v706 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v706
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v706
	v710 = int32(1)
	if l0 != int32(3) {
		v1199 = v710
		goto L2
	} else {
		goto L226
	}
L205:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[6]))
	if v654 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v671 == int32(0) {
		goto L204
	} else {
		goto L215
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(28)
	v671 = int32(-1)
	goto L206
L208:
	;
	v658 = v654
	goto L209
L209:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+12))
	if v649 != v659 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v671 = int32(0)
	goto L206
L211:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v658)+20))
	if v661 != 0 {
		v658 = v661
		goto L209
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	goto L210
L214:
	;
	goto L207
L215:
	;
	v675 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	if v675 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L217
	}
L217:
	;
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v680 != int32(48))&base.B2i32(v680 != int32(22)) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_23), v16+int32(144))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L6
	} else {
		goto L224
	}
L219:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L6
	} else {
		goto L223
	}
L222:
	;
	goto L218
L223:
	;
	goto L218
L224:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(538), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	v1199 = v8
	goto L2
L226:
	;
	v713 = int32(0)
	v715 = F_pgl_shmctl(m, v638, v713, v713)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	if int32(0) <= v715 {
		v1199 = v710
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v720 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	if v720 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1199 = int32(0)
	goto L2
L231:
	;
	goto L232
L232:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v726 != int32(48))&base.B2i32(v726 != int32(22)) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_3), v16+int32(128))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L6
	} else {
		goto L239
	}
L234:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
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
	v738 = m.ExcPending
	if v738 != 0 {
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
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(548), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	v1199 = int32(0)
	goto L2
L241:
	;
	v758 = F_pgl_shmctl(m, v638, int32(2), v16+int32(336))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L6
	} else {
		goto L244
	}
L242:
	;
	v793 = l2
	goto L243
L243:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[6]))
	if v796 != 0 {
		goto L259
	} else {
		goto L260
	}
L244:
	;
	if v758 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v761 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L6
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v16)+372))
	v793 = v792
	goto L243
L248:
	;
	if v761 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L249
	}
L249:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v766 != int32(48))&base.B2i32(v766 != int32(22)) == int32(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_13), v16+int32(176))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L6
	} else {
		goto L256
	}
L251:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L6
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L6
	} else {
		goto L255
	}
L254:
	;
	goto L250
L255:
	;
	goto L250
L256:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(564), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L6
	} else {
		goto L257
	}
L257:
	;
	v1199 = v8
	goto L2
L258:
	;
	if v811 == int32(-1) {
		goto L268
	} else {
		goto L269
	}
L259:
	;
	v798 = v796
	goto L262
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = int32(28)
	v811 = int32(-1)
	goto L258
L262:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)))
	if v638 == v799 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	goto L261
L264:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798)+12))
	v811 = v801
	goto L258
L265:
	;
	goto L266
L266:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v798)+20))
	if v802 != 0 {
		v798 = v802
		goto L262
	} else {
		goto L267
	}
L267:
	;
	goto L263
L268:
	;
	v815 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if l0 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v811
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v793
	v1199 = int32(1)
	goto L2
L271:
	;
	v818 = int32(0)
	v820 = F_pgl_shmctl(m, v638, v818, v818)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v815
	v825 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L6
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	if v825 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L276
	}
L276:
	;
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v830 != int32(48))&base.B2i32(v830 != int32(22)) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_16), v16+int32(160))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L283
	}
L278:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L6
	} else {
		goto L282
	}
L281:
	;
	goto L277
L282:
	;
	goto L277
L283:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(585), int32(_a_F_dsm_impl_op_18))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L6
	} else {
		goto L284
	}
L284:
	;
	v1199 = v8
	goto L2
L285:
	;
	if l0&int32(-2) == int32(2) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v872 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	goto L288
L288:
	;
	if l0 != 0 {
		goto L304
	} else {
		goto L305
	}
L289:
	;
	goto L291
L290:
	;
	goto L291
L291:
	;
	v874 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v874
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v874
	v878 = int32(1)
	if l0 != int32(3) {
		v1199 = v878
		goto L2
	} else {
		goto L292
	}
L292:
	;
	v883 = F_unlink(m, v16+int32(432))
	mBase = m.M
	if v883 == int32(0) {
		v1199 = v878
		goto L2
	} else {
		goto L293
	}
L293:
	;
	v886 = int32(0)
	v888 = F_errstart(m, l6, v886)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	if v888 == int32(0) {
		v1199 = v886
		goto L2
	} else {
		goto L295
	}
L295:
	;
	v893 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v893 != int32(48))&base.B2i32(v893 != int32(22)) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+208)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_3), v16+int32(208))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L6
	} else {
		goto L302
	}
L297:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L6
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L6
	} else {
		goto L301
	}
L300:
	;
	goto L296
L301:
	;
	goto L296
L302:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(823), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L6
	} else {
		goto L303
	}
L303:
	;
	v1199 = v886
	goto L2
L304:
	;
	v923 = int32(2)
	goto L306
L305:
	;
	v923 = int32(194)
	goto L306
L306:
	;
	v924 = F_OpenTransientFile(m, v16+int32(432), v923)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L6
	} else {
		goto L307
	}
L307:
	;
	if v924 == int32(-1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	if l0 != int32(1) {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L310
L310:
	;
	if l0 == int32(1) {
		goto L327
	} else {
		goto L328
	}
L311:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if v931 == int32(20) {
		v1199 = v8
		goto L2
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v935 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L6
	} else {
		goto L315
	}
L314:
	;
	goto L313
L315:
	;
	if v935 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L316
	}
L316:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v940 != int32(48))&base.B2i32(v940 != int32(22)) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_10), v16+int32(240))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L6
	} else {
		goto L323
	}
L318:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L6
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L6
	} else {
		goto L322
	}
L321:
	;
	goto L317
L322:
	;
	goto L317
L323:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(837), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L6
	} else {
		goto L324
	}
L324:
	;
	v1199 = v8
	goto L2
L325:
	;
	v1115 = int32(1)
	v1117 = F_mmap(m, v1104, v1115, v924)
	mBase = m.M
	if v1117 == int32(-1) {
		goto L376
	} else {
		goto L377
	}
L326:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
	v1104 = v1101
	goto L325
L327:
	;
	if v924 < int32(0) {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	goto L329
L329:
	;
	v1019 = F_palloc0(m, int32(_a_F_dsm_impl_op_25))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L6
	} else {
		goto L346
	}
L330:
	;
	if v977 == int32(0) {
		goto L326
	} else {
		goto L334
	}
L331:
	;
	v973 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v977 = v973
	goto L330
L332:
	;
	goto L333
L333:
	;
	v976 = F___fstatat(m, v924, int32(_a_F_dsm_impl_op_11), v16+int32(336), int32(_a_F_dsm_impl_op_12))
	mBase = m.M
	v977 = v976
	goto L330
L334:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v982 = F_CloseTransientFile(m, v924)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L6
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v981
	v987 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L6
	} else {
		goto L336
	}
L336:
	;
	if v987 == int32(0) {
		v1199 = v8
		goto L2
	} else {
		goto L337
	}
L337:
	;
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v992 != int32(48))&base.B2i32(v992 != int32(22)) == int32(0) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+288)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_13), v16+int32(288))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L6
	} else {
		goto L344
	}
L339:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L6
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L6
	} else {
		goto L343
	}
L342:
	;
	goto L338
L343:
	;
	goto L338
L344:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(861), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L6
	} else {
		goto L345
	}
L345:
	;
	v1199 = v8
	goto L2
L346:
	;
	if l2 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1104 = int32(0)
	goto L325
L348:
	;
	goto L349
L349:
	;
	v1025 = l2
	goto L350
L350:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1038))) = int32(167772186)
	v1041 = int32(_a_F_dsm_impl_op_25)
	if base.Ui32(v1041) <= base.Ui32(v1025) {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	if v1044 == v1045 {
		v1104 = l2
		goto L325
	} else {
		goto L361
	}
L352:
	;
	goto L351
L353:
	;
	v1044 = v1041
	goto L355
L354:
	;
	v1044 = v1025
	goto L355
L355:
	;
	v1045 = F_write(m, v924, v1019, v1044)
	mBase = m.M
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1047))) = int32(0)
	v1050 = base.B2i32(v1044 == v1045)
	if v1044 != v1045 {
		goto L352
	} else {
		goto L356
	}
L356:
	;
	if v1044 == v1045 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1053 = v1044
	goto L359
L358:
	;
	v1053 = int32(0)
	goto L359
L359:
	;
	v1054 = v1025 - v1053
	if v1054 != 0 {
		v1025 = v1054
		goto L350
	} else {
		goto L360
	}
L360:
	;
	goto L352
L361:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v1058 = F_CloseTransientFile(m, v924)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L6
	} else {
		goto L362
	}
L362:
	;
	v1062 = F_unlink(m, v16+int32(432))
	mBase = m.M
	if v1057 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1065 = v1057
	goto L365
L364:
	;
	v1065 = int32(51)
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v1065
	v1067 = int32(0)
	v1069 = F_errstart(m, l6, v1067)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L6
	} else {
		goto L366
	}
L366:
	;
	if v1069 == int32(0) {
		v1199 = v1067
		goto L2
	} else {
		goto L367
	}
L367:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v1074 != int32(48))&base.B2i32(v1074 != int32(22)) == int32(0) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+308)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+304)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_15), v16+int32(304))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L6
	} else {
		goto L374
	}
L369:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L6
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L6
	} else {
		goto L373
	}
L372:
	;
	goto L368
L373:
	;
	goto L368
L374:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(912), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L6
	} else {
		goto L375
	}
L375:
	;
	v1199 = v1067
	goto L2
L376:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	v1122 = F_CloseTransientFile(m, v924)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L6
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1104
	v1166 = F_CloseTransientFile(m, v924)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L6
	} else {
		goto L393
	}
L379:
	;
	if l0 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1128 = F_unlink(m, v16+int32(432))
	mBase = m.M
	goto L382
L381:
	;
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1])) = v1121
	v1131 = int32(0)
	v1133 = F_errstart(m, l6, v1131)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L6
	} else {
		goto L383
	}
L383:
	;
	if v1133 == int32(0) {
		v1199 = v1131
		goto L2
	} else {
		goto L384
	}
L384:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[1]))
	if base.B2i32(v1138 != int32(48))&base.B2i32(v1138 != int32(22)) == int32(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+256)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_16), v16+int32(256))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L6
	} else {
		goto L391
	}
L386:
	;
	F_errcode(m, int32(_a_F_dsm_impl_op_4))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L6
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L6
	} else {
		goto L390
	}
L389:
	;
	goto L385
L390:
	;
	goto L385
L391:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(934), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L6
	} else {
		goto L392
	}
L392:
	;
	v1199 = v1131
	goto L2
L393:
	;
	if v1166 == int32(0) {
		v1199 = v1115
		goto L2
	} else {
		goto L394
	}
L394:
	;
	v1170 = int32(0)
	v1172 = F_errstart(m, l6, v1170)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L6
	} else {
		goto L395
	}
L395:
	;
	if v1172 == int32(0) {
		v1199 = v1170
		goto L2
	} else {
		goto L396
	}
L396:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L6
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+272)) = v16 + int32(432)
	F_errmsg(m, int32(_a_F_dsm_impl_op_26), v16+int32(272))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L6
	} else {
		goto L398
	}
L398:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(945), int32(_a_F_dsm_impl_op_24))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L6
	} else {
		goto L399
	}
L399:
	;
	v1199 = v1170
	goto L2
L400:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, _c_F_dsm_impl_op[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v1213
	F_errmsg_internal(m, int32(_a_F_dsm_impl_op_27), v16)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L6
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(_a_F_dsm_impl_op_5), int32(191), int32(_a_F_dsm_impl_op_28))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L6
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
