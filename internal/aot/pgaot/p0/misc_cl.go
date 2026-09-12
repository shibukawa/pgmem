package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CloneForeignKeyConstraints(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v576 int32
	_ = v576
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v668 int32
	_ = v668
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(912)
	m.G0 = v24
	v26 = F_RelationGetFKeyList(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L5
	} else {
		goto L218
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L5
	} else {
		goto L214
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L5
	} else {
		goto L211
	}
L4:
	;
	v692 = F_table_open(m, int32(2606), int32(2))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L94
	}
L5:
	;
	return
L6:
	;
	if v26 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v30 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = v4
	v41 = v4
	goto L9
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37<<(uint(int32(2))%32))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v59 == v60 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	if v63 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v63 = F_lappend_oid(m, v41, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v66 = v37 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v66 < v67 {
		v37 = v66
		v41 = v63
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	if v72 == int32(102) {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v77 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v82 = F_build_attrmap_by_name(m, v79, v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v84 = F_RelationGetFKeyList(m, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v86 = F_copyObjectImpl(m, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if int32(0) < v88 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v110 = v4
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_sequence_close(m, v77, int32(3))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L5
	} else {
		goto L93
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v110<<(uint(int32(2))%32))))
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+172)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v117
	v122 = F_SearchSysCache1(m, int32(19), v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v122 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)))
	v128 = v126 + v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+92))
	v130 = int32(0)
	if v63 == v130 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v642 = v110 + int32(1)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v642 < v643 {
		v110 = v642
		goto L23
	} else {
		goto L92
	}
L28:
	;
	if v168 != 0 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v168 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v136 <= int32(0) {
		v161 = v130
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v168 = v161
	goto L28
L33:
	;
	v139 = int32(0)
	if v139 < v136 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v142 = v136
	goto L36
L35:
	;
	v142 = v139
	goto L36
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v145 = int32(0)
	goto L37
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143+v145<<(uint(int32(2))%32))))
	v154 = base.B2i32(v153 == v129)
	if v153 == v129 {
		v161 = v154
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v161 = v154
	goto L32
L39:
	;
	v156 = v145 + int32(1)
	if v156 != v142 {
		v145 = v156
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v173 = F_table_open(m, v171, int32(6))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	goto L27
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+48))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+119)))
	if v176 == int32(112) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
	v182 = F_find_all_inheritors(m, v179, int32(6), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_DeconstructFkConstraintRow(m, v122, v24+int32(864), v24+int32(768), v24+int32(624), v24+int32(432), v24+int32(304), v24+int32(176), v24+int32(764), v24+int32(560))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v202 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	if v383 != 0 {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v206 = v202 & int32(3)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v209 = v207 - int32(2)
	v210 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v202) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v221 = v210
	v233 = int32(0)
	goto L56
L54:
	;
	v303 = v210
	goto L55
L55:
	;
	if v206 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v238 = int32(1)
	v239 = v221 << (uint(v238) % 32)
	v241 = v24 + int32(688)
	v244 = v24 + int32(768)
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v239))))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v246<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v239+v241))) = uint16(v250)
	v253 = v239 | int32(2)
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v253))))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v260<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v253+v241))) = uint16(v264)
	v266 = int32(4)
	v267 = v239 | v266
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v267))))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v274<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v267+v241))) = uint16(v278)
	v281 = v239 | int32(6)
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244+v281))))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v288<<(uint(v238)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v281+v241))) = uint16(v292)
	v295 = v221 + v266
	v297 = v233 + v266
	if v297 != v202&int32(2147483644) {
		v221 = v295
		v233 = v297
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v303 = v295
	goto L55
L58:
	;
	goto L57
L59:
	;
	v326 = v303
	v327 = v210
	goto L60
L60:
	;
	v343 = int32(1)
	v344 = v326 << (uint(v343) % 32)
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(768)+v344))))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209+v351<<(uint(v343)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v344+(v24+int32(688))))) = uint16(v355)
	v360 = v327 + v343
	if v360 != v206 {
		v326 = v326 + v343
		v327 = v360
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v128)+96))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v128)+80))
	F_GetForeignKeyCheckTriggers(m, v77, v384, v385, v386, v24+int32(172), v24+int32(92))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v86 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v469 = F_palloc0(m, int32(108))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L79
	}
L68:
	;
	v395 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v396 <= v395 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v405 = v395
	goto L70
L70:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v422+v405<<(uint(int32(2))%32))))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v434 = F_tryAttachPartitionForeignKey(m, l0, v426, l2, v116, v427, v24+int32(688), v24+int32(624), v24+int32(432), v400, v399, v77)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L72
	}
L71:
	;
	F_sequence_close(m, v173, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L77
	}
L72:
	;
	if v434 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v439 = v405 + int32(1)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v439 < v440 {
		v405 = v439
		goto L70
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L71
L76:
	;
	goto L67
L77:
	;
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	goto L27
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v469))) = int64(438086664353)
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+12)) = uint8(v473)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+13)) = uint8(v475)
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v469)+80)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v469)+72)) = v479
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+86)) = uint8(v483)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+87)) = uint8(v485)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+100)) = v479
	*(*int64)(unsafe.Add(mBase, uint32(v469)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+88)) = uint8(v487)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+15)) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+14)) = uint8(v493)
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+16)) = uint8(v497)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v479 < v499 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v469)+76))
	v508 = int32(0)
	v510 = v502
	goto L83
L81:
	;
	v557 = v499
	goto L82
L82:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v128)+88))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+107)))
	F_addFkConstraint(m, v24+int32(96), int32(1), v128+int32(4), v469, l2, v173, v576, v116, v557, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v587, v24+int32(560), int32(0), v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L5
	} else {
		goto L88
	}
L83:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v508<<(uint(int32(1))%32)))))
	v541 = F_makeString(m, v525+v526<<(uint(int32(4))%32)+v535*int32(100)-int32(76))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L5
	} else {
		goto L85
	}
L84:
	;
	v557 = v548
	goto L82
L85:
	;
	v543 = F_lappend(m, v510, v541)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+76)) = v543
	v547 = v508 + int32(1)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	if v547 < v548 {
		v508 = v547
		v510 = v543
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v24)+100))
	F_ReleaseCatCache(m, v122)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v24)+864))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
	F_addFkRecurseReferencing(m, l0, v469, l2, v173, v576, v594, v597, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v608, v24+int32(560), int32(0), int32(8), v613, v614, v591)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_sequence_close(m, v173, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	goto L27
L92:
	;
	goto L24
L93:
	;
	goto L4
L94:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v24+int32(768), int32(13), int32(3), int32(184), v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	F_ScanKeyInit(m, v24+int32(816), int32(4), int32(3), int32(61), int32(102))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v710 = int32(0)
	v717 = F_systable_beginscan(m, v692, v710, int32(1), v710, int32(2), v24+int32(768))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v719 = F_systable_getnext(m, v717)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	if v719 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v725 = v719
	v731 = v710
	goto L102
L100:
	;
	v760 = v710
	goto L101
L101:
	;
	F_systable_endscan(m, v717)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L5
	} else {
		goto L107
	}
L102:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v725)+16))
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+22)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v742+v743)))
	v746 = F_lappend_oid(m, v731, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L104
	}
L103:
	;
	v760 = v746
	goto L101
L104:
	;
	v748 = F_systable_getnext(m, v717)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	if v748 != 0 {
		v725 = v748
		v731 = v746
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	F_sequence_close(m, v692, int32(2))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v778 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v783 = F_build_attrmap_by_name(m, v780, v781, int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	if v760 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_sequence_close(m, v778, int32(3))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L5
	} else {
		goto L210
	}
L112:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v787 <= int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v809 = int32(0)
	goto L114
L114:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813+v809<<(uint(int32(2))%32))))
	v818 = F_SearchSysCache1(m, int32(19), v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L118
	}
L115:
	;
	goto L111
L116:
	;
	F_ReleaseCatCache(m, v818)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L5
	} else {
		goto L208
	}
L117:
	;
	v1334 = int32(0)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+8))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+107)))
	F_addFkConstraint(m, v24+int32(864), v1334, v1335, v1019, v865, l2, v1124, v817, v1336, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v1347, v24+int32(96), v1334, v1351)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L205
	}
L118:
	;
	if v818 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+22)))
	v822 = v820 + v821
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+92))
	v824 = int32(0)
	if v760 == v824 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L121
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L5
	} else {
		goto L202
	}
L122:
	;
	if v862 != 0 {
		goto L116
	} else {
		goto L135
	}
L123:
	;
	v862 = int32(0)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v830 <= int32(0) {
		v855 = v824
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v862 = v855
	goto L122
L127:
	;
	v833 = int32(0)
	if v833 < v830 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v836 = v830
	goto L130
L129:
	;
	v836 = v833
	goto L130
L130:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v839 = int32(0)
	goto L131
L131:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v837+v839<<(uint(int32(2))%32))))
	v848 = base.B2i32(v847 == v823)
	if v847 == v823 {
		v855 = v848
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v855 = v848
	goto L126
L133:
	;
	v850 = v839 + int32(1)
	if v850 != v836 {
		v839 = v850
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v822)+80))
	v865 = F_table_open(m, v863, int32(6))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v822)+88))
	F_DeconstructFkConstraintRow(m, v818, v24+int32(764), v24+int32(688), v24+int32(560), v24+int32(432), v24+int32(304), v24+int32(176), v24+int32(172), v24+int32(96))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v886 <= int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v1019 = F_palloc0(m, int32(108))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L147
	}
L139:
	;
	v889 = int32(1)
	v891 = int32(0)
	if v886 != v889 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v901 = v891
	v904 = int32(0)
	goto L143
L141:
	;
	v961 = v891
	goto L142
L142:
	;
	if v886&v889 == int32(0) {
		goto L138
	} else {
		goto L146
	}
L143:
	;
	v918 = int32(1)
	v919 = v901 << (uint(v918) % 32)
	v921 = v24 + int32(624)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v925 = v24 + int32(560)
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v925+v919))))
	v931 = int32(2)
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v923+v927<<(uint(v918)%32)-v931))))
	*(*uint16)(unsafe.Add(mBase, uint32(v919+v921))) = uint16(v933)
	v936 = v919 | v931
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v944 = int32(*(*int16)(unsafe.Add(mBase, uint32(v925+v936))))
	v950 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v940+v944<<(uint(v918)%32)-v931))))
	*(*uint16)(unsafe.Add(mBase, uint32(v936+v921))) = uint16(v950)
	v953 = v901 + v931
	v955 = v904 + v931
	if v955 != v886&int32(2147483646) {
		v901 = v953
		v904 = v955
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v961 = v953
	goto L142
L145:
	;
	goto L144
L146:
	;
	v980 = int32(1)
	v981 = v961 << (uint(v980) % 32)
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v989 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(560)+v981))))
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985+v989<<(uint(v980)%32)-int32(2)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v981+(v24+int32(624))))) = uint16(v995)
	goto L138
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+8)) = v822 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v1019))) = int64(438086664353)
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+12)) = uint8(v1026)
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+104)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+13)) = uint8(v1028)
	v1032 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+80)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+72)) = v1032
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+102)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+86)) = uint8(v1036)
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+87)) = uint8(v1038)
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+101)))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+100)) = v1032
	*(*int64)(unsafe.Add(mBase, uint32(v1019)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+88)) = uint8(v1040)
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+15)) = uint8(v1032)
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+14)) = uint8(v1046)
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1019)+16)) = uint8(v1050)
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v1032 < v1052 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+76))
	v1061 = int32(0)
	v1063 = v1055
	goto L151
L149:
	;
	goto L150
L150:
	;
	v1124 = F_index_get_partition(m, l2, v867)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L5
	} else {
		goto L156
	}
L151:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v865)+52))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1078)))
	v1088 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+int32(688)+v1061<<(uint(int32(1))%32)))))
	v1094 = F_makeString(m, v1078+v1079<<(uint(int32(4))%32)+v1088*int32(100)-int32(76))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	v1096 = F_lappend(m, v1063, v1094)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L5
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+76)) = v1096
	v1100 = v1061 + int32(1)
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	if v1100 < v1101 {
		v1061 = v1100
		v1063 = v1096
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	if v1124 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1126 = int32(0)
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+75)))
	if v1128 != int32(1) {
		v1315 = v1126
		v1322 = v1126
		goto L117
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L5
	} else {
		goto L199
	}
L160:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v822)+80))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v822)+96))
	F_ScanKeyInit(m, v24+int32(864), int32(11), int32(3), int32(184), v817)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L5
	} else {
		goto L161
	}
L161:
	;
	v1140 = int32(0)
	v1142 = int32(1)
	v1147 = F_systable_beginscan(m, v778, int32(2699), v1142, v1140, v1142, v24+int32(864))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L5
	} else {
		goto L164
	}
L162:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L5
	} else {
		goto L196
	}
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L5
	} else {
		goto L193
	}
L164:
	;
	v1149 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if v1149 == int32(0) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v1153 = v1140
	v1157 = v1149
	v1164 = v1126
	goto L167
L167:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+16))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174)+22)))
	v1176 = v1174 + v1175
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+84))
	if v1177 != v1131 {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L172
	}
L168:
	;
	goto L163
L169:
	;
	v1226 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L5
	} else {
		goto L191
	}
L170:
	;
	F_systable_endscan(m, v1147)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L190
	}
L171:
	;
	v1213 = F_systable_getnext(m, v1147)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L5
	} else {
		goto L186
	}
L172:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+4))
	if v1179 != v1132 {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+76))
	v1184 = v1181 - int32(1644)
	if base.Ui32(v1184) <= base.Ui32(int32(11)) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v1192 != int32(1) {
		v1209 = v1153
		v1211 = v1164
		goto L171
	} else {
		goto L178
	}
L175:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1184<<(uint(int32(2))%32))+uint32(_c_F_CloneForeignKeyConstraints[0])))
	v1192 = v1191
	goto L177
L176:
	;
	v1192 = int32(0)
	goto L177
L177:
	;
	goto L174
L178:
	;
	v1195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1176)+80)))
	if v1195&int32(8) != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v1204 == int32(0) {
		goto L169
	} else {
		goto L184
	}
L180:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	v1204 = v1198
	v1205 = v1153
	goto L179
L181:
	;
	goto L182
L182:
	;
	if v1195&int32(16) == int32(0) {
		v1204 = v1164
		v1205 = v1153
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	v1204 = v1164
	v1205 = v1203
	goto L179
L184:
	;
	if v1205 != 0 {
		v1220 = v1205
		v1221 = v1204
		goto L170
	} else {
		goto L185
	}
L185:
	;
	v1209 = int32(0)
	v1211 = v1204
	goto L171
L186:
	;
	if v1213 != 0 {
		v1153 = v1209
		v1157 = v1213
		v1164 = v1211
		goto L167
	} else {
		goto L187
	}
L187:
	;
	if v1211 == int32(0) {
		goto L163
	} else {
		goto L188
	}
L188:
	;
	if v1209 == int32(0) {
		goto L162
	} else {
		goto L189
	}
L189:
	;
	v1220 = v1209
	v1221 = v1211
	goto L170
L190:
	;
	v1315 = v1220
	v1322 = v1221
	goto L117
L191:
	;
	if v1226 != 0 {
		v1153 = v1205
		v1157 = v1226
		v1164 = int32(0)
		goto L167
	} else {
		goto L192
	}
L192:
	;
	goto L168
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v817
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_0), v24+int32(32))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_2), int32(_a_F_CloneForeignKeyConstraints_3))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v817
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_4), v24+int32(48))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_5), int32(_a_F_CloneForeignKeyConstraints_3))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v1283 + int32(4)
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_6), v24+int32(16))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_7), int32(_a_F_CloneForeignKeyConstraints_8))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L5
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
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v817
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_9), v24)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_10), int32(_a_F_CloneForeignKeyConstraints_8))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L5
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
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v24)+868))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v24)+764))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822)+107)))
	F_addFkRecurseReferenced(m, v1019, v865, l2, v1124, v1354, v1355, v24+int32(624), v24+int32(688), v24+int32(432), v24+int32(304), v24+int32(176), v1366, v24+int32(96), v1322, v1315, v1369)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	F_sequence_close(m, v865, int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	goto L116
L208:
	;
	v1399 = v809 + int32(1)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v1399 < v1400 {
		v809 = v1399
		goto L114
	} else {
		goto L209
	}
L209:
	;
	goto L115
L210:
	;
	m.G0 = v24 + int32(912)
	return
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v116
	F_errmsg_internal(m, int32(_a_F_CloneForeignKeyConstraints_9), v24+int32(80))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_11), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	F_errmsg(m, int32(_a_F_CloneForeignKeyConstraints_13), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_14), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L5
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
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v1469 = F_get_constraint_name(m, v1468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v1469
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v1467 + int32(4)
	F_errmsg(m, int32(_a_F_CloneForeignKeyConstraints_15), v24-int32(-64))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_CloneForeignKeyConstraints_1), int32(_a_F_CloneForeignKeyConstraints_16), int32(_a_F_CloneForeignKeyConstraints_12))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CloseTransientFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CloseTransientFile[0]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CloseTransientFile[1]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v40 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v19 = v12 + v14*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(3) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v14 {
		v14 = v14 - int32(1)
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v23 != l0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = F_FreeDesc(m, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v25
L11:
	;
	goto L5
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_CloseTransientFile_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_pgaio_closing_fd(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(_a_F_CloseTransientFile_1), int32(2892), int32(_a_F_CloseTransientFile_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v53 = F_close(m, l0)
	mBase = m.M
	return v53
}
func F___clock_gettime(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if base.Ui32(int32(4)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F___clock_gettime[0])) = int32(28)
	} else {
		v18 = m.Wasi_snapshot_preview1.Clock_time_get(m, l0, int64(1), v8+int32(24))
		mBase = m.M
		if v18 == int32(0) {
			v25 = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F___clock_gettime[0])) = v18
			v25 = int32(-1)
		}
		if v25 != 0 {
		} else {
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			v28 = v8 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
			v31 = int64(1000000000)
			v32 = base.I64_div_u_s(v26, v31)
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v32
			v36 = v26 - v32*v31
			*(*uint32)(unsafe.Add(mBase, uint32(v28)+8)) = uint32(v36)
			v38 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v38
			v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v40
		}
	}
	m.G0 = v8 + int32(32)
	return
}
func F_cleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int64
	_ = v104
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int64
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_markreachable(m, l0, v10, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_markcanreach(m, l0, v13, v14, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_cleartraverse(m, l0, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L88
	}
L8:
	;
	v24 = v17
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v26 != 0 {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	goto L7
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 == v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v27 != 0 {
		v24 = v27
		goto L9
	} else {
		goto L87
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
	if v31 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v37 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L47
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
	if v45 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L15
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v79 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v50 = v48 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v50) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if int32(1)<<(uint(v50)%32)&int32(_a_F_cleanup_0) == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v59 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	if v60 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v72 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v45*int32(24))+12)) = v68
	v72 = v68
	goto L26
L28:
	;
	goto L29
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v70
	v72 = v70
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+36)) = v60
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+32)) = int64(0)
	goto L21
L33:
	;
	if v78 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v78
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v78
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v79
	goto L39
L38:
	;
	goto L39
L39:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v85 - int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	if v90 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v96 = v37 + int32(8)
	if v89 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v89
	goto L40
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v89
	goto L40
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v90
	goto L46
L45:
	;
	goto L46
L46:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v98 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = v104
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v37
	goto L20
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v118 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(-1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v199 != 0 {
		goto L80
	} else {
		goto L81
	}
L49:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118)+4)))
	if v126 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L47
L53:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
	if v160 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v131 = v129 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v131) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if int32(1)<<(uint(v131)%32)&int32(_a_F_cleanup_0) == int32(0) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v140 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	if v141 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v153 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v145+v126*int32(24))+12)) = v149
	v153 = v149
	goto L58
L60:
	;
	goto L61
L61:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v151
	v153 = v151
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+36)) = v141
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118)+32)) = int64(0)
	goto L53
L65:
	;
	if v159 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v159
	goto L65
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v159
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = v160
	goto L71
L70:
	;
	goto L71
L71:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+12)) = v166 - int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v118)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v118)+28))
	if v171 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v177 = v118 + int32(8)
	if v170 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+16)) = v170
	goto L72
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v170
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+28)) = v171
	goto L78
L77:
	;
	goto L78
L78:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = v179 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(0)
	v185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v177)+16)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v177)+8)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v177))) = v185
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v118
	goto L52
L79:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v198 != 0 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+32)) = v198
	goto L79
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v198
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
	goto L12
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v202
	goto L83
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v202
	goto L83
L87:
	;
	goto L10
L88:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v226 = v223
	v227 = v224
	goto L92
L90:
	;
	v235 = v223
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v235
	goto L3
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v226
	v232 = v226 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+28))
	if v233 != 0 {
		v226 = v232
		v227 = v233
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v235 = v232
	goto L91
L94:
	;
	goto L93
}
func F_cloneouts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+4)))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_cloneouts[0]))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v21 <= v22 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v78 != 0 {
		v14 = v78
		goto L4
	} else {
		goto L33
	}
L12:
	;
	F_createarc(m, l0, l4, v16, l2, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L32
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v24 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L16:
	;
	v28 = v24
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v34 != l3 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L12
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v42 != 0 {
		v28 = v42
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if v36 != v16&int32(_a_F_cloneouts_0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v40 == l4 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	goto L18
L24:
	;
	v47 = v43
	goto L25
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v53 != l2 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L12
L27:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	if v61 != 0 {
		v47 = v61
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v55 != v16&int32(_a_F_cloneouts_0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v59 == l4 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L26
L32:
	;
	goto L11
L33:
	;
	goto L5
}
func F_close_ls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 float64
	_ = v68
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = v12 + int32(16)
	v15 = F_point_sl(m, v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v20 = base.F64_abs(v19)
		if base.F64_le(v20, float64(1e-06)) != 0 {
			v41 = float64(0)
			if base.F64_eq(v41, v15) != 0 {
				v45 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
				v86 = int32(0)
				return v86
			} else {
				v49 = F_palloc(m, int32(16))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = F_lseg_interpt_line(m, v49, v12, v11)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v68 = float64(0)
							if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
								v86 = v49
							} else {
								v76 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
								v86 = int32(0)
							}
							return v86
						} else {
							v54 = F_line_closept_point(m, int32(0), v11, v12)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v57 = F_line_closept_point(m, int32(0), v11, v14)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									v59 = base.F64_lt(v54, v57)
									if v59 != 0 {
										v60 = v54
									} else {
										v60 = v57
									}
									if v49 == int32(0) {
										v68 = v60
									} else {
										if v59 != 0 {
											v63 = v12
										} else {
											v63 = v14
										}
										v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
										*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
										v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
										v68 = v60
									}
									if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v86 = v49
									} else {
										v76 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
										v86 = int32(0)
									}
									return v86
								}
							}
						}
					}
				}
			}
		} else {
			v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			v25 = base.F64_abs(v24)
			if base.F64_le(v25, float64(1e-06)) != 0 {
				v41 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v41, v15) != 0 {
					v45 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
					v86 = int32(0)
					return v86
				} else {
					v49 = F_palloc(m, int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = F_lseg_interpt_line(m, v49, v12, v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if v51 != 0 {
								v68 = float64(0)
								if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
									v86 = v49
								} else {
									v76 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
									v86 = int32(0)
								}
								return v86
							} else {
								v54 = F_line_closept_point(m, int32(0), v11, v12)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v57 = F_line_closept_point(m, int32(0), v11, v14)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = base.F64_lt(v54, v57)
										if v59 != 0 {
											v60 = v54
										} else {
											v60 = v57
										}
										if v49 == int32(0) {
											v68 = v60
										} else {
											if v59 != 0 {
												v63 = v12
											} else {
												v63 = v14
											}
											v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
											*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
											v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
											v68 = v60
										}
										if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v86 = v49
										} else {
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											v86 = int32(0)
										}
										return v86
									}
								}
							}
						}
					}
				}
			} else {
				v30 = base.F64_div(v19, base.F64_neg(v24))
				v32 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v30), v32)&base.F64_ne(v20, v32) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if base.F64_ne(v30, float64(0)) != 0 {
						v41 = v30
						if base.F64_eq(v41, v15) != 0 {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v86 = int32(0)
							return v86
						} else {
							v49 = F_palloc(m, int32(16))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = F_lseg_interpt_line(m, v49, v12, v11)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v51 != 0 {
										v68 = float64(0)
										if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
											v86 = v49
										} else {
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											v86 = int32(0)
										}
										return v86
									} else {
										v54 = F_line_closept_point(m, int32(0), v11, v12)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v57 = F_line_closept_point(m, int32(0), v11, v14)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v59 = base.F64_lt(v54, v57)
												if v59 != 0 {
													v60 = v54
												} else {
													v60 = v57
												}
												if v49 == int32(0) {
													v68 = v60
												} else {
													if v59 != 0 {
														v63 = v12
													} else {
														v63 = v14
													}
													v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
													*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
													v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
													v68 = v60
												}
												if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
													v86 = v49
												} else {
													v76 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
													v86 = int32(0)
												}
												return v86
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_ne(v25, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v41 = v30
							if base.F64_eq(v41, v15) != 0 {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v86 = int32(0)
								return v86
							} else {
								v49 = F_palloc(m, int32(16))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = F_lseg_interpt_line(m, v49, v12, v11)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v51 != 0 {
											v68 = float64(0)
											if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
												v86 = v49
											} else {
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												v86 = int32(0)
											}
											return v86
										} else {
											v54 = F_line_closept_point(m, int32(0), v11, v12)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												v57 = F_line_closept_point(m, int32(0), v11, v14)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													v59 = base.F64_lt(v54, v57)
													if v59 != 0 {
														v60 = v54
													} else {
														v60 = v57
													}
													if v49 == int32(0) {
														v68 = v60
													} else {
														if v59 != 0 {
															v63 = v12
														} else {
															v63 = v14
														}
														v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
														*(*int64)(unsafe.Add(mBase, uint32(v49))) = v64
														v66 = *(*int64)(unsafe.Add(mBase, uint32(v63)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v66
														v68 = v60
													}
													if base.Ui64(base.I64_reinterpret_f64(v68)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
														v86 = v49
													} else {
														v76 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
														v86 = int32(0)
													}
													return v86
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
}
func F_close_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_box_closept_lseg(m, v8, v5, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				v22 = int32(0)
			} else {
				v22 = v8
			}
			return v22
		}
	}
}
func F_closelog(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = m.G0
	v3 = int32(16)
	v4 = v2 - v3
	m.G0 = v4
	v6 = int32(_a_F_closelog_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_closelog[0]))
	v8 = F_close(m, v7)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_closelog[0])) = int32(-1)
	m.G0 = v4 + v3
	return
}
func F_cluster_rel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
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
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
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
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 float64
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 float64
	_ = v531
	var v532 float64
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 float64
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 float64
	_ = v557
	var v560 int32
	_ = v560
	var v561 float64
	_ = v561
	var v567 float64
	_ = v567
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 float64
	_ = v584
	var v585 float64
	_ = v585
	var v598 int32
	_ = v598
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v745 int32
	_ = v745
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 float64
	_ = v771
	var v772 float64
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 float64
	_ = v823
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v919 int32
	_ = v919
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(400)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[0]))
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v39 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	if l1 != 0 {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	goto L6
L8:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v43 != int32(1) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(_a_F_cluster_rel_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v49 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v48 + v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v52 + v49
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = v28
	v59 = v39 + int32(232)
	if v59&int32(3) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v86 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v85 + v86
	v89 = int32(_a_F_cluster_rel_0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v91 - v86
	goto L7
L11:
	;
	v65 = v39 + int32(392)
	if base.Ui32(v65) <= base.Ui32(v59) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v82 = F___memset(m, v59, int32(0), int32(160))
	mBase = m.M
	goto L10
L14:
	;
	v69 = v39 + int32(236)
	if base.Ui32(v69) < base.Ui32(v65) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = v65
	goto L17
L16:
	;
	v71 = v69
	goto L17
L17:
	;
	v79 = F___memset(m, v59, int32(0), (v71-v39-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L10
L18:
	;
	v100 = int64(1)
	goto L20
L19:
	;
	v100 = int64(2)
	goto L20
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v103 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(124)))) = v139
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v26+int32(120)))) = v142
	goto L25
L22:
	;
	goto L21
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v107 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v110 = int32(_a_F_cluster_rel_0)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v112 + v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v116 + v113
	*(*int64)(unsafe.Add(mBase, uint32(v103+int32(0))+232)) = v100
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v124 + v113
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v130 - v113
	goto L22
L25:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+80))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5])) = v146 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4])) = v145
	goto L26
L26:
	;
	v154 = int32(_a_F_cluster_rel_1)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[6]))
	v158 = v156 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[6])) = v158
	goto L27
L27:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	if v29&int32(2) != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L211
	}
L30:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_2), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L209
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L4
	} else {
		goto L205
	}
L32:
	;
	F_AtEOXact_GUC(m, int32(0), v158)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L198
	}
L33:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+118)))
	if v224 != int32(116) {
		goto L66
	} else {
		goto L67
	}
L34:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+117)))
	if v220 == int32(1) {
		goto L31
	} else {
		goto L65
	}
L35:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v164 = F_pg_class_aclcheck(m, v28, v162, int64(16384))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l1 != 0 {
		goto L34
	} else {
		goto L64
	}
L38:
	;
	if v164 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v168 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+118)))
	if v187 != int32(116) {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	if v168 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v170 = F_get_rel_name(m, v28)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+112)) = v170
	F_errmsg(m, int32(_a_F_cluster_rel_3), v26+int32(112))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1752), int32(_a_F_cluster_rel_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L32
L50:
	;
	if l1 == int32(0) {
		v223 = v186
		goto L33
	} else {
		goto L54
	}
L51:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v190 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	goto L32
L54:
	;
	v197 = int32(0)
	v200 = F_SearchSysCacheExists(m, int32(57), l1, v197, v197, v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v200 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v207&int32(4) == int32(0) {
		goto L34
	} else {
		goto L60
	}
L59:
	;
	goto L32
L60:
	;
	v212 = F_get_index_isclustered(m, l1)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v212 != 0 {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	goto L32
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v223 = v217
	goto L33
L65:
	;
	v223 = v219
	goto L33
L66:
	;
	if l1 != 0 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v227 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	if l1 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_6), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(425), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v246 = int32(_a_F_cluster_rel_8)
	goto L76
L75:
	;
	v246 = int32(_a_F_cluster_rel_9)
	goto L76
L76:
	;
	F_CheckTableNotInUse(m, l0, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	if l1 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_check_index_is_clusterable(m, l0, l1, int32(8))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	v256 = int32(0)
	goto L80
L80:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+119)))
	if v258 != int32(109) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v253 = F_index_open(m, l1, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v256 = v253
	goto L80
L83:
	;
	F_TransferPredicateLocksToHeapRelation(m, l0)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L87
	}
L84:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+129)))
	if v261 != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	F_relation_close(m, l0, int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	goto L32
L87:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+92))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v268)+84))
	if v256 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	F_mark_index_clustered(m, l0, v271, int32(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v276 = v268
	goto L90
L90:
	;
	v277 = int32(*(*int8)(unsafe.Add(mBase, uint32(v276)+118)))
	v279 = int32(1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v280) < base.Ui32(int32(_a_F_cluster_rel_10)) {
		v289 = v279
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v276 = v275
	goto L90
L92:
	;
	v291 = F_make_new_heap(m, v267, v269, v270, v277, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L96
	}
L93:
	;
	goto L92
L94:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+68))
	if v284 == int32(99) {
		v289 = v279
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v287 = F_isTempToastNamespace(m, v284)
	mBase = m.M
	v289 = v287
	goto L93
L96:
	;
	v294 = F_table_open(m, v291, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v296 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+312)) = v296
	*(*int64)(unsafe.Add(mBase, uint32(v26)+304)) = v296
	*(*int64)(unsafe.Add(mBase, uint32(v26)+296)) = v296
	F_getrusage(m, v26+int32(144))
	mBase = m.M
	F___gettimeofday(m, v26+int32(128))
	mBase = m.M
	goto L98
L98:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+68))
	v310 = F_get_namespace_name(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+112))
	if v313 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v331 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+392)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+384)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+376)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+368)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+360)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+352)) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v26)+344)) = v331
	v349 = F_vacuum_get_cutoffs(m, l0, v26+int32(344), v26+int32(320))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L105
	}
L101:
	;
	F_LockRelationOid(m, v313, int32(8))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+112))
	if v320 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v294)+48))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+112))
	if v324 == int32(0) {
		v330 = v4
		goto L100
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+264)) = v320
	v330 = int32(1)
	goto L100
L105:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+136))
	if v352 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+140))
	if v374 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v352))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v355)) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if v367 == int32(0) {
		goto L106
	} else {
		goto L112
	}
L109:
	;
	v367 = base.B2i32(base.Ui32(v355) < base.Ui32(v352))
	goto L108
L110:
	;
	goto L111
L111:
	;
	v367 = int32(base.Ui32(v355-v352) >> (uint(int32(31)) % 32))
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+336)) = v352
	goto L106
L113:
	;
	if v29&int32(1) != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	goto L115
L115:
	;
	if int32(base.Ui32(v377-v374)>>(uint(int32(31))%32)) == int32(0) {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+340)) = v374
	goto L113
L117:
	;
	v386 = int32(17)
	goto L119
L118:
	;
	v386 = int32(13)
	goto L119
L119:
	;
	if v256 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v26)+328))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)+124))
	m.T0[v757].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v294, v256, v726, v745, v26+int32(336), v26+int32(340), v26+int32(312), v26+int32(304), v26+int32(296))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L168
	}
L121:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), v700, int32(_a_F_cluster_rel_11))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L167
	}
L122:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+84))
	if v388 == int32(403) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	v676 = int32(0)
	v678 = F_errstart(m, v386, v676)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L164
	}
L125:
	;
	v660 = F_errstart(m, v386, int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L161
	}
L126:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	v393 = m.G0
	v395 = v393 - int32(112)
	m.G0 = v395
	v397 = int32(1)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[7])))
	if v399 != v397 {
		v598 = v397
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	v636 = int32(0)
	v638 = F_errstart(m, v386, v636)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L4
	} else {
		goto L158
	}
L129:
	;
	m.G0 = v395 + int32(112)
	if v598 != 0 {
		goto L125
	} else {
		goto L157
	}
L130:
	;
	v403 = F_palloc0(m, int32(168))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v403))) = int64(4294967363)
	v408 = F_palloc0(m, int32(92))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = int32(266)
	v413 = F_palloc0(m, int32(384))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+8)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v413)+4)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = int32(267)
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+344)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+280)) = v422
	v427 = F_palloc0(m, int32(8))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v395)+20)) = v427
	v436 = F_list_make1_impl(m, int32(1), v395+int32(12))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+84)) = v436
	v440 = F_palloc0(m, int32(136))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v442 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v440)+24)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v440)+16)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v440)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(101)
	v449 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+124)) = uint16(v449)
	v451 = int32(_a_F_cluster_rel_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v440)+20)) = uint16(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v440
	v458 = F_list_make1_impl(m, v442, v395+int32(8))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+52)) = v458
	v463 = F_addRTEPermissionInfo(m, v403+int32(56), v440)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_setup_simple_rel_arrays(m, v413)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v469 = F_build_simple_rel(m, v413, int32(1), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v469)+108))
	if v471 == int32(0) {
		v598 = v397
		goto L129
	} else {
		goto L141
	}
L141:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v474 <= int32(0) {
		v598 = v397
		goto L129
	} else {
		goto L142
	}
L142:
	;
	v477 = int32(0)
	if v477 < v474 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v481 = v474
	goto L145
L144:
	;
	v481 = v477
	goto L145
L145:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v487 = v477
	goto L146
L146:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v482+v487<<(uint(int32(2))%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v392 != v510 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v516 = *(*float64)(unsafe.Add(mBase, uint32(v469)+120))
	*(*float64)(unsafe.Add(mBase, uint32(v469)+16)) = v516
	v519 = F_get_relation_data_width(m, v391, int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L152
	}
L148:
	;
	v512 = int32(1)
	v514 = v487 + v512
	if v481 != v514 {
		v487 = v514
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	v598 = v512
	goto L129
L152:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v469)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v521)+32)) = v519
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v469)+116))
	*(*float64)(unsafe.Add(mBase, uint32(v413)+288)) = base.F64_convert_i32_u(v523)
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v509)+84))
	F_cost_qual_eval(m, v395+int32(96), v528, v413)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v531 = *(*float64)(unsafe.Add(mBase, uint32(v395)+104))
	v532 = *(*float64)(unsafe.Add(mBase, uint32(v395)+96))
	v534 = v395 + int32(24)
	v535 = int32(0)
	v537 = F_create_seqscan_path(m, v413, v469, v535, v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v537)+40))
	v540 = *(*float64)(unsafe.Add(mBase, uint32(v537)+56))
	v541 = *(*float64)(unsafe.Add(mBase, uint32(v469)+120))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v469)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+32))
	v544 = base.F64_add(v532, v531)
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[9]))
	v550 = m.G0
	v551 = int32(16)
	v552 = v550 - v551
	m.G0 = v552
	F_cost_tuplesort(m, v552+int32(8), v552, v541, v543, base.F64_add(v544, v544), v547, float64(-1))
	mBase = m.M
	v557 = *(*float64)(unsafe.Add(mBase, uint32(v552)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+32)) = v541
	v560 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[10])))
	v561 = base.F64_add(v540, v557)
	*(*float64)(unsafe.Add(mBase, uint32(v534)+48)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v534)+40)) = v539 + (v560 ^ int32(1))
	v567 = *(*float64)(unsafe.Add(mBase, uint32(v552)))
	*(*float64)(unsafe.Add(mBase, uint32(v534)+56)) = base.F64_add(v561, v567)
	m.G0 = v552 + v551
	goto L155
L155:
	;
	v573 = int32(0)
	v582 = F_create_index_path(m, v413, v509, v573, v573, v573, v573, int32(1), v573, v573, float64(1), v573)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v584 = *(*float64)(unsafe.Add(mBase, uint32(v395)+80))
	v585 = *(*float64)(unsafe.Add(mBase, uint32(v582)+56))
	v598 = base.F64_lt(v584, v585)
	goto L129
L157:
	;
	goto L128
L158:
	;
	if v638 == int32(0) {
		v726 = v636
		goto L120
	} else {
		goto L159
	}
L159:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v256)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v310
	v645 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v643 + v645
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v642 + v645
	F_errmsg(m, int32(_a_F_cluster_rel_13), v26+int32(96))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v700 = int32(964)
	v717 = int32(0)
	goto L121
L161:
	;
	if v660 == int32(0) {
		v726 = int32(1)
		goto L120
	} else {
		goto L162
	}
L162:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v664 + int32(4)
	F_errmsg(m, int32(_a_F_cluster_rel_14), v26+int32(80))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v700 = int32(969)
	v717 = int32(1)
	goto L121
L164:
	;
	if v678 == int32(0) {
		v726 = v676
		goto L120
	} else {
		goto L165
	}
L165:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v682 + int32(4)
	F_errmsg(m, int32(_a_F_cluster_rel_15), v26-int32(-64))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v700 = int32(974)
	v717 = int32(0)
	goto L121
L167:
	;
	v726 = v717
	goto L120
L168:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v26)+340))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v26)+336))
	v762 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+264)) = v762
	v765 = F_RelationGetNumberOfBlocksInFork(m, v294, v762)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v768 = F_errstart(m, v386, int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	if v768 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v771 = *(*float64)(unsafe.Add(mBase, uint32(v26)+304))
	v772 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	v774 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v810 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L179
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v774
	*(*float64)(unsafe.Add(mBase, uint32(v26)+48)) = v772
	*(*float64)(unsafe.Add(mBase, uint32(v26)+40)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v770 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v310
	F_errmsg(m, int32(_a_F_cluster_rel_16), v26+int32(32))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v26)+296))
	v791 = F_pg_rusage_show(m, v26+int32(128))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v791
	*(*float64)(unsafe.Add(mBase, uint32(v26)+16)) = v788
	F_errdetail(m, int32(_a_F_cluster_rel_17), v26+int32(16))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1007), int32(_a_F_cluster_rel_11))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	goto L173
L179:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v294)+56))
	v815 = F_SearchSysCacheCopy(m, int32(57), v813, int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	if v815 == int32(0) {
		goto L29
	} else {
		goto L181
	}
L181:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819)+22)))
	v821 = v819 + v820
	*(*int32)(unsafe.Add(mBase, uint32(v821)+96)) = v765
	v823 = *(*float64)(unsafe.Add(mBase, uint32(v26)+312))
	*(*float32)(unsafe.Add(mBase, uint32(v821)+100)) = base.F32_demote_f64(v823)
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v826 != int32(1259) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	F_pfree(m, v815)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L188
	}
L183:
	;
	F_CatalogTupleUpdate(m, v810, v815+int32(4), v815)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	F_CacheInvalidateRelcacheByTuple(m, v815)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L187
	}
L186:
	;
	goto L182
L187:
	;
	goto L182
L188:
	;
	F_sequence_close(m, v810, int32(3))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	F_sequence_close(m, l0, int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	if v256 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_relation_close(m, v256, int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L4
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	F_sequence_close(m, v294, int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	F_finish_heap_swap(m, v267, v291, v289, v330, int32(0), int32(1), v761, v760, v277)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	goto L32
L198:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v26)+124))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[5])) = v882
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[4])) = v881
	goto L199
L199:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[1]))
	if v889 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	m.G0 = v26 + int32(400)
	return
L201:
	;
	goto L200
L202:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cluster_rel[2])))
	if v893 != int32(1) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v889)+220))
	if v896 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v899 = int32(_a_F_cluster_rel_0)
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	v902 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v901 + v902
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v905 + v902
	v909 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v889)+220)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v889)+224)) = v909
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v905 + int32(2)
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_cluster_rel[3])) = v919 - v902
	goto L201
L205:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	F_errmsg(m, int32(_a_F_cluster_rel_18), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(410), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(421), int32(_a_F_cluster_rel_7))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v294)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v955
	F_errmsg_internal(m, int32(_a_F_cluster_rel_19), v26)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_cluster_rel_4), int32(1016), int32(_a_F_cluster_rel_11))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
