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
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1193 int32
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(528)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[763]))
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
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L6
	} else {
		goto L400
	}
L2:
	;
	m.G0 = v16 + int32(528)
	return v1207
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+320)) = l1
	v873 = F_pg_snprintf(m, v16+int32(464), int32(64), int32(41528), v16+int32(320))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L6
	} else {
		goto L285
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+192)) = l1
	v480 = F_pg_snprintf(m, v16+int32(464), int32(64), int32(64245), v16+int32(192))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L142
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l1
	v29 = F_pg_snprintf(m, v16+int32(464), int32(64), int32(41548), v16+int32(112))
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
		v1207 = v43
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v49 = m.G0
	v51 = v49 - int32(272)
	m.G0 = v51
	v53 = F___shm_mapname(m, v16+int32(464), v51)
	mBase = m.M
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v56 == int32(0) {
		v1207 = v43
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
		v1207 = v62
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v69 != int32(48))&base.B2i32(v69 != int32(22)) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(464)
	F_errmsg(m, int32(308607), v16+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	F_errcode(m, int32(8389))
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
	F_errfinish(m, int32(517369), int32(242), int32(27575))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v1207 = v62
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
	v110 = v16 + int32(464)
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
	v147 = F___memcpy(m, v107, int32(592728), int32(9))
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(37)
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
	v162 = F_open(m, v154, v99|int32(657408), v102)
	mBase = m.M
	v164 = v162
	goto L34
L51:
	;
	v170 = int32(4464764)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	*(*int32)(unsafe.Add(mBase, _consts[757])) = v172 - int32(1)
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
	v179 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v179 == int32(20) {
		v1207 = v8
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
		v1207 = v8
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v188 != int32(48))&base.B2i32(v188 != int32(22)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v16 + int32(464)
	F_errmsg(m, int32(308561), v16+int32(48))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L67
	}
L62:
	;
	F_errcode(m, int32(8389))
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
	F_errfinish(m, int32(517369), int32(266), int32(27575))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v1207 = v8
	goto L2
L69:
	;
	v399 = int32(1)
	v401 = F___mmap(m, v388, v399, v164)
	mBase = m.M
	if v401 == int32(-1) {
		goto L120
	} else {
		goto L121
	}
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
	v388 = v385
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
	v273 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v273 == int32(1) {
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
	v224 = F___fstatat(m, v164, int32(785340), v16+int32(336), int32(4096))
	mBase = m.M
	v225 = v224
	goto L74
L78:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v230 = F_close(m, v164)
	mBase = m.M
	v231 = int32(4464764)
	v233 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	*(*int32)(unsafe.Add(mBase, _consts[757])) = v233 - int32(1)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v229
	v239 = int32(0)
	v241 = F_errstart(m, l6, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	if v241 == int32(0) {
		v1207 = v239
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v246 != int32(48))&base.B2i32(v246 != int32(22)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v16 + int32(464)
	F_errmsg(m, int32(308423), v16+int32(80))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L88
	}
L83:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
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
	v258 = m.ExcPending
	if v258 != 0 {
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
	F_errfinish(m, int32(517369), int32(291), int32(27575))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	v1207 = v239
	goto L2
L90:
	;
	F_sigprocmask(m, int32(4456120), v16+int32(336))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = int32(167772185)
	goto L94
L93:
	;
	goto L92
L94:
	;
	v299 = F_ftruncate(m, v164, base.I64_extend_i32_u(l2))
	mBase = m.M
	if v299 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = int32(0)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v311 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L96:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v303 == int32(27) {
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
	v315 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	F_sigprocmask(m, v16+int32(336), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v299 == int32(0) {
		v388 = l2
		goto L69
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v315
	goto L102
L104:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v328 = F_close(m, v164)
	mBase = m.M
	v329 = int32(4464764)
	v331 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	*(*int32)(unsafe.Add(mBase, _consts[757])) = v331 - int32(1)
	goto L105
L105:
	;
	v338 = m.G0
	v340 = v338 - int32(272)
	m.G0 = v340
	v342 = F___shm_mapname(m, v16+int32(464), v340)
	mBase = m.M
	if v342 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v327
	v351 = int32(0)
	v353 = F_errstart(m, l6, v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L110
	}
L107:
	;
	v343 = F_unlink(m, v342)
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	m.G0 = v340 + int32(272)
	goto L106
L110:
	;
	if v353 == int32(0) {
		v1207 = v351
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v358 != int32(48))&base.B2i32(v358 != int32(22)) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v16 + int32(464)
	F_errmsg(m, int32(304852), v16+int32(96))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L118
	}
L113:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
	v370 = m.ExcPending
	if v370 != 0 {
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
	F_errfinish(m, int32(517369), int32(310), int32(27575))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	v1207 = v351
	goto L2
L120:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v406 = F_close(m, v164)
	mBase = m.M
	v407 = int32(4464764)
	v409 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	*(*int32)(unsafe.Add(mBase, _consts[757])) = v409 - int32(1)
	goto L123
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v401
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v388
	v466 = F_close(m, v164)
	mBase = m.M
	v467 = int32(4464764)
	v469 = *(*int32)(unsafe.Add(mBase, _consts[757]))
	*(*int32)(unsafe.Add(mBase, _consts[757])) = v469 - int32(1)
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
	v418 = m.G0
	v420 = v418 - int32(272)
	m.G0 = v420
	v422 = F___shm_mapname(m, v16+int32(464), v420)
	mBase = m.M
	if v422 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v405
	v431 = int32(0)
	v433 = F_errstart(m, l6, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L6
	} else {
		goto L131
	}
L127:
	;
	goto L126
L128:
	;
	v423 = F_unlink(m, v422)
	mBase = m.M
	goto L130
L129:
	;
	goto L130
L130:
	;
	m.G0 = v420 + int32(272)
	goto L127
L131:
	;
	if v433 == int32(0) {
		v1207 = v431
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v438 != int32(48))&base.B2i32(v438 != int32(22)) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v16 + int32(464)
	F_errmsg(m, int32(308516), v16-int32(-64))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L139
	}
L134:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
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
	v450 = m.ExcPending
	if v450 != 0 {
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
	F_errfinish(m, int32(517369), int32(332), int32(27575))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	v1207 = v431
	goto L2
L141:
	;
	v1207 = v399
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
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v504 != 0 {
		goto L153
	} else {
		goto L154
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(20)
	v1207 = v8
	goto L2
L147:
	;
	v488 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	if v488 == int32(0) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	F_errmsg_internal(m, int32(559219), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(517369), int32(470), int32(33642))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
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
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v639 = v505
	v640 = v504
	goto L152
L154:
	;
	goto L155
L155:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v509 = F_MemoryContextAlloc(m, v507, int32(4))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v512 = l1 >> (uint(int32(31)) % 32)
	v514 = l1 ^ v512 - v512
	if l0 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v516 = int32(0)
	goto L159
L158:
	;
	v516 = l2
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
	v519 = int32(384)
	goto L162
L161:
	;
	v519 = int32(1920)
	goto L162
L162:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	if v524 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v595 == int32(-1) {
		goto L185
	} else {
		goto L186
	}
L164:
	;
	v528 = v524
	goto L167
L165:
	;
	goto L166
L166:
	;
	if v519&int32(512) != 0 {
		goto L173
	} else {
		goto L174
	}
L167:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	if v514 == v531 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L166
L169:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v595 = v533
	goto L163
L170:
	;
	goto L171
L171:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528)+20))
	if v534 != 0 {
		v528 = v534
		goto L167
	} else {
		goto L172
	}
L172:
	;
	goto L168
L173:
	;
	v547 = int32(65536)
	goto L176
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(44)
	v595 = int32(-1)
	goto L163
L176:
	;
	if base.Ui32(v547) < base.Ui32(v516) {
		v547 = v547 << (uint(int32(1)) % 32)
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v553 = F_emscripten_builtin_malloc(m, v547)
	mBase = m.M
	if v553 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(48)
	v595 = int32(-1)
	goto L163
L180:
	;
	goto L181
L181:
	;
	v561 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v561 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	F_emscripten_builtin_free(m, v553)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(48)
	v595 = int32(-1)
	goto L163
L183:
	;
	goto L184
L184:
	;
	v569 = int32(4207164)
	v571 = *(*int32)(unsafe.Add(mBase, _consts[765]))
	*(*int32)(unsafe.Add(mBase, _consts[765])) = v571 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v561)+16)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v561)+12)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v561)+8)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v561)+4)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v571
	v580 = int32(4633208)
	v581 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	*(*int32)(unsafe.Add(mBase, uint32(v561)+20)) = v581
	*(*int32)(unsafe.Add(mBase, _consts[764])) = v561
	v595 = v571
	goto L163
L185:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v599 == int32(20))&base.B2i32(l0 != int32(1)) != 0 {
		v1207 = v8
		goto L2
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = v595
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v509
	v639 = v595
	v640 = v509
	goto L152
L188:
	;
	F_pfree(m, v509)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v599
	v610 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	if v610 == int32(0) {
		v1207 = v8
		goto L2
	} else {
		goto L191
	}
L191:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v615 != int32(48))&base.B2i32(v615 != int32(22)) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	F_errmsg(m, int32(304176), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L198
	}
L193:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
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
	v627 = m.ExcPending
	if v627 != 0 {
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
	F_errfinish(m, int32(517369), int32(519), int32(33642))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v1207 = v8
	goto L2
L200:
	;
	F_pfree(m, v640)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
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
	v648 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v648
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v650 == v648 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v708 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v708
	v712 = int32(1)
	if l0 != int32(3) {
		v1207 = v712
		goto L2
	} else {
		goto L226
	}
L205:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	if v655 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v672 == int32(0) {
		goto L204
	} else {
		goto L215
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	v672 = int32(-1)
	goto L206
L208:
	;
	v659 = v655
	goto L209
L209:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+12))
	if v650 != v660 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v672 = int32(0)
	goto L206
L211:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v659)+20))
	if v662 != 0 {
		v659 = v662
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
	v675 = int32(0)
	v677 = F_errstart(m, l6, v675)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	if v677 == int32(0) {
		v1207 = v675
		goto L2
	} else {
		goto L217
	}
L217:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v682 != int32(48))&base.B2i32(v682 != int32(22)) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+144)) = v16 + int32(464)
	F_errmsg(m, int32(308469), v16+int32(144))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L224
	}
L219:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
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
	v694 = m.ExcPending
	if v694 != 0 {
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
	F_errfinish(m, int32(517369), int32(538), int32(33642))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	v1207 = v675
	goto L2
L226:
	;
	v715 = int32(0)
	v717 = F_pgl_shmctl(m, v639, v715, v715)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	if int32(0) <= v717 {
		v1207 = v712
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v722 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	if v722 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1207 = int32(0)
	goto L2
L231:
	;
	goto L232
L232:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v728 != int32(48))&base.B2i32(v728 != int32(22)) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v16 + int32(464)
	F_errmsg(m, int32(308607), v16+int32(128))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L6
	} else {
		goto L239
	}
L234:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
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
	v740 = m.ExcPending
	if v740 != 0 {
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
	F_errfinish(m, int32(517369), int32(548), int32(33642))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	v1207 = int32(0)
	goto L2
L241:
	;
	v760 = F_pgl_shmctl(m, v639, int32(2), v16+int32(336))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L6
	} else {
		goto L244
	}
L242:
	;
	v795 = l2
	goto L243
L243:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _consts[764]))
	if v799 != 0 {
		goto L259
	} else {
		goto L260
	}
L244:
	;
	if v760 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v763 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L6
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v16)+372))
	v795 = v794
	goto L243
L248:
	;
	if v763 == int32(0) {
		v1207 = v8
		goto L2
	} else {
		goto L249
	}
L249:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v768 != int32(48))&base.B2i32(v768 != int32(22)) == int32(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+176)) = v16 + int32(464)
	F_errmsg(m, int32(308423), v16+int32(176))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L256
	}
L251:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
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
	v780 = m.ExcPending
	if v780 != 0 {
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
	F_errfinish(m, int32(517369), int32(564), int32(33642))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L257
	}
L257:
	;
	v1207 = v8
	goto L2
L258:
	;
	if v817 == int32(-1) {
		goto L268
	} else {
		goto L269
	}
L259:
	;
	v801 = v799
	goto L262
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(28)
	v817 = int32(-1)
	goto L258
L262:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	if v639 == v803 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	goto L261
L264:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v801)+12))
	v817 = v805
	goto L258
L265:
	;
	goto L266
L266:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v801)+20))
	if v806 != 0 {
		v801 = v806
		goto L262
	} else {
		goto L267
	}
L267:
	;
	goto L263
L268:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _consts[137]))
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v817
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v795
	v1207 = int32(1)
	goto L2
L271:
	;
	v824 = int32(0)
	v826 = F_pgl_shmctl(m, v639, v824, v824)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L6
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v821
	v830 = int32(0)
	v832 = F_errstart(m, l6, v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L6
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	if v832 == int32(0) {
		v1207 = v830
		goto L2
	} else {
		goto L276
	}
L276:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v837 != int32(48))&base.B2i32(v837 != int32(22)) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+160)) = v16 + int32(464)
	F_errmsg(m, int32(308516), v16+int32(160))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L6
	} else {
		goto L283
	}
L278:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
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
	v849 = m.ExcPending
	if v849 != 0 {
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
	F_errfinish(m, int32(517369), int32(585), int32(33642))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L6
	} else {
		goto L284
	}
L284:
	;
	v1207 = v830
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
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v879 != 0 {
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
	v881 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v881
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v881
	v885 = int32(1)
	if l0 != int32(3) {
		v1207 = v885
		goto L2
	} else {
		goto L292
	}
L292:
	;
	v890 = F_unlink(m, v16+int32(464))
	mBase = m.M
	if v890 == int32(0) {
		v1207 = v885
		goto L2
	} else {
		goto L293
	}
L293:
	;
	v893 = int32(0)
	v895 = F_errstart(m, l6, v893)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	if v895 == int32(0) {
		v1207 = v893
		goto L2
	} else {
		goto L295
	}
L295:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v900 != int32(48))&base.B2i32(v900 != int32(22)) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+208)) = v16 + int32(464)
	F_errmsg(m, int32(308607), v16+int32(208))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L6
	} else {
		goto L302
	}
L297:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
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
	v912 = m.ExcPending
	if v912 != 0 {
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
	F_errfinish(m, int32(517369), int32(823), int32(248286))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L6
	} else {
		goto L303
	}
L303:
	;
	v1207 = v893
	goto L2
L304:
	;
	v930 = int32(2)
	goto L306
L305:
	;
	v930 = int32(194)
	goto L306
L306:
	;
	v931 = F_OpenTransientFile(m, v16+int32(464), v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L6
	} else {
		goto L307
	}
L307:
	;
	if v931 == int32(-1) {
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
	v938 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v938 == int32(20) {
		v1207 = v8
		goto L2
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v942 = F_errstart(m, l6, int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L6
	} else {
		goto L315
	}
L314:
	;
	goto L313
L315:
	;
	if v942 == int32(0) {
		v1207 = v8
		goto L2
	} else {
		goto L316
	}
L316:
	;
	v947 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v947 != int32(48))&base.B2i32(v947 != int32(22)) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+240)) = v16 + int32(464)
	F_errmsg(m, int32(308561), v16+int32(240))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L6
	} else {
		goto L323
	}
L318:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
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
	v959 = m.ExcPending
	if v959 != 0 {
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
	F_errfinish(m, int32(517369), int32(837), int32(248286))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L6
	} else {
		goto L324
	}
L324:
	;
	v1207 = v8
	goto L2
L325:
	;
	v1123 = int32(1)
	v1125 = F___mmap(m, v1112, v1123, v931)
	mBase = m.M
	if v1125 == int32(-1) {
		goto L376
	} else {
		goto L377
	}
L326:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
	v1112 = v1109
	goto L325
L327:
	;
	if v931 < int32(0) {
		goto L331
	} else {
		goto L332
	}
L328:
	;
	goto L329
L329:
	;
	v1027 = F_palloc0(m, int32(8192))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L6
	} else {
		goto L346
	}
L330:
	;
	if v984 == int32(0) {
		goto L326
	} else {
		goto L334
	}
L331:
	;
	v980 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v984 = v980
	goto L330
L332:
	;
	goto L333
L333:
	;
	v983 = F___fstatat(m, v931, int32(785340), v16+int32(336), int32(4096))
	mBase = m.M
	v984 = v983
	goto L330
L334:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v989 = F_CloseTransientFile(m, v931)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L6
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v988
	v993 = int32(0)
	v995 = F_errstart(m, l6, v993)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L6
	} else {
		goto L336
	}
L336:
	;
	if v995 == int32(0) {
		v1207 = v993
		goto L2
	} else {
		goto L337
	}
L337:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v1000 != int32(48))&base.B2i32(v1000 != int32(22)) == int32(0) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+288)) = v16 + int32(464)
	F_errmsg(m, int32(308423), v16+int32(288))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L6
	} else {
		goto L344
	}
L339:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
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
	v1012 = m.ExcPending
	if v1012 != 0 {
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
	F_errfinish(m, int32(517369), int32(861), int32(248286))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L6
	} else {
		goto L345
	}
L345:
	;
	v1207 = v993
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
	v1112 = int32(0)
	goto L325
L348:
	;
	goto L349
L349:
	;
	v1033 = l2
	goto L350
L350:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v1046))) = int32(167772186)
	v1049 = int32(8192)
	if base.Ui32(v1049) <= base.Ui32(v1033) {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	if v1053 == v1052 {
		v1112 = l2
		goto L325
	} else {
		goto L361
	}
L352:
	;
	goto L351
L353:
	;
	v1052 = v1049
	goto L355
L354:
	;
	v1052 = v1033
	goto L355
L355:
	;
	v1053 = F_write(m, v931, v1027, v1052)
	mBase = m.M
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v1055))) = int32(0)
	v1058 = base.B2i32(v1053 == v1052)
	if v1053 != v1052 {
		goto L352
	} else {
		goto L356
	}
L356:
	;
	if v1053 == v1052 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1061 = v1052
	goto L359
L358:
	;
	v1061 = int32(0)
	goto L359
L359:
	;
	v1062 = v1033 - v1061
	if v1062 != 0 {
		v1033 = v1062
		goto L350
	} else {
		goto L360
	}
L360:
	;
	goto L352
L361:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v1066 = F_CloseTransientFile(m, v931)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L362
	}
L362:
	;
	v1070 = F_unlink(m, v16+int32(464))
	mBase = m.M
	if v1065 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1073 = v1065
	goto L365
L364:
	;
	v1073 = int32(51)
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v1073
	v1075 = int32(0)
	v1077 = F_errstart(m, l6, v1075)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L6
	} else {
		goto L366
	}
L366:
	;
	if v1077 == int32(0) {
		v1207 = v1075
		goto L2
	} else {
		goto L367
	}
L367:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v1082 != int32(48))&base.B2i32(v1082 != int32(22)) == int32(0) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+308)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+304)) = v16 + int32(464)
	F_errmsg(m, int32(304852), v16+int32(304))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L6
	} else {
		goto L374
	}
L369:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
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
	v1094 = m.ExcPending
	if v1094 != 0 {
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
	F_errfinish(m, int32(517369), int32(912), int32(248286))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L6
	} else {
		goto L375
	}
L375:
	;
	v1207 = v1075
	goto L2
L376:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v1130 = F_CloseTransientFile(m, v931)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L6
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1112
	v1174 = F_CloseTransientFile(m, v931)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
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
	v1136 = F_unlink(m, v16+int32(464))
	mBase = m.M
	goto L382
L381:
	;
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v1129
	v1139 = int32(0)
	v1141 = F_errstart(m, l6, v1139)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L6
	} else {
		goto L383
	}
L383:
	;
	if v1141 == int32(0) {
		v1207 = v1139
		goto L2
	} else {
		goto L384
	}
L384:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if base.B2i32(v1146 != int32(48))&base.B2i32(v1146 != int32(22)) == int32(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+256)) = v16 + int32(464)
	F_errmsg(m, int32(308516), v16+int32(256))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L6
	} else {
		goto L391
	}
L386:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
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
	v1158 = m.ExcPending
	if v1158 != 0 {
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
	F_errfinish(m, int32(517369), int32(934), int32(248286))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L6
	} else {
		goto L392
	}
L392:
	;
	v1207 = v1139
	goto L2
L393:
	;
	if v1174 == int32(0) {
		v1207 = v1123
		goto L2
	} else {
		goto L394
	}
L394:
	;
	v1178 = int32(0)
	v1180 = F_errstart(m, l6, v1178)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L6
	} else {
		goto L395
	}
L395:
	;
	if v1180 == int32(0) {
		v1207 = v1178
		goto L2
	} else {
		goto L396
	}
L396:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L6
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+272)) = v16 + int32(464)
	F_errmsg(m, int32(308655), v16+int32(272))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L6
	} else {
		goto L398
	}
L398:
	;
	F_errfinish(m, int32(517369), int32(945), int32(248286))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L6
	} else {
		goto L399
	}
L399:
	;
	v1207 = v1178
	goto L2
L400:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, _consts[763]))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v1221
	F_errmsg_internal(m, int32(502702), v16)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L6
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(517369), int32(191), int32(245184))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
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
