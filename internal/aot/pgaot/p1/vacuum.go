package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v327 int32
	_ = v327
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v368 int32
	_ = v368
	var v383 int32
	_ = v383
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v514 int32
	_ = v514
	var v525 int32
	_ = v525
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v655 int32
	_ = v655
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v877 int32
	_ = v877
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v946 int32
	_ = v946
	var v959 int32
	_ = v959
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int64
	_ = v1082
	var v1084 int64
	_ = v1084
	var v1086 int64
	_ = v1086
	var v1088 int64
	_ = v1088
	var v1090 int64
	_ = v1090
	var v1092 int64
	_ = v1092
	var v1094 int64
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1132 int32
	_ = v1132
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1272 int32
	_ = v1272
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1388 int32
	_ = v1388
	var v1397 int32
	_ = v1397
	var v1410 int32
	_ = v1410
	var v1421 int32
	_ = v1421
	var v1432 int32
	_ = v1432
	var v1443 int32
	_ = v1443
	var v1451 int32
	_ = v1451
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1493 int32
	_ = v1493
	var v1504 int32
	_ = v1504
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1531 int32
	_ = v1531
	var v1540 int32
	_ = v1540
	var v1561 int32
	_ = v1561
	var v1587 int32
	_ = v1587
	var v1588 int64
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	v6 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(336)
	m.G0 = v28
	v37 = v6
	v38 = v6
	v39 = v6
	v40 = v6
	v41 = v6
	v42 = v6
	v43 = v6
	v44 = v6
	v45 = v6
	v46 = int32(-1)
	v50 = v6
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v46 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1587 = int32(m.ExcTag)
	v1588 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1587 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L7:
	;
	if v1002 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L8:
	;
	v989 = v37
	v990 = v38
	v991 = v39
	v993 = v41
	v994 = v42
	v995 = v43
	v996 = v44
	v997 = v45
	v998 = v40
	v1002 = v50
	goto L7
L9:
	;
	goto L10
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v60 = v58 & int32(1)
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+299)) = uint8(v96)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])))
	if v99 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	F_PreventInTransactionBlock(m, l4, int32(_a_F_vacuum_0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[1]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	if base.Ui32(v85) <= base.Ui32(int32(1)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v96 = int32(0)
	goto L11
L16:
	;
	v88 = int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+28))
	v95 = l4 ^ v88 | base.B2i32(v88 < v90)
	goto L18
L17:
	;
	v95 = int32(1)
	goto L18
L18:
	;
	v96 = v95
	goto L11
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v155&int32(1024) != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	F_errcode(m, int32(1088))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v136 = int32(_a_F_vacuum_0)
	goto L26
L25:
	;
	v136 = int32(_a_F_vacuum_1)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v136
	F_errmsg(m, int32(_a_F_vacuum_2), v28)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	F_errfinish(m, int32(_a_F_vacuum_3), int32(538), int32(_a_F_vacuum_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v915 = int32(1)
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v916&v915 != 0 {
		v931 = v915
		goto L120
	} else {
		goto L121
	}
L30:
	;
	v898 = v39
	v900 = v41
	v901 = v42
	v902 = v43
	v903 = v44
	v904 = v45
	v905 = l0
	goto L29
L31:
	;
	goto L32
L32:
	;
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v158 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v713 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L102
	}
L36:
	;
	v898 = v39
	v900 = v41
	v901 = v42
	v902 = v43
	v903 = v44
	v904 = v45
	v905 = int32(0)
	goto L29
L37:
	;
	goto L38
L38:
	;
	v162 = int32(0)
	v172 = v39
	v176 = v43
	v177 = v44
	v178 = v45
	v179 = v162
	v182 = v162
	goto L39
L39:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v179<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	if v194 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v898 = v694
	v900 = v41
	v901 = v42
	v902 = v668
	v903 = v669
	v904 = v670
	v905 = v694
	goto L29
L41:
	;
	v681 = int32(_a_F_vacuum_5)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v670
	v694 = F_list_concat(m, v182, v673)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L6
	} else {
		goto L100
	}
L42:
	;
	v195 = int32(_a_F_vacuum_5)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v209 = F_lappend(m, int32(0), v193)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v229 = int32(0)
	v231 = F_RangeVarGetRelidExtended(m, v214, int32(1), int32(base.Ui32(v213)>>(uint(int32(3))%32))&int32(4), v229, v229)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v196
	v668 = v176
	v669 = v177
	v670 = v209
	v673 = v209
	goto L41
L46:
	;
	if v231 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v246 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v352 = F_SearchSysCache1(m, int32(57), v231)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L62
	}
L50:
	;
	if v213&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v250 = int32(0)
	if v246 == v250 {
		v668 = v176
		v669 = v177
		v670 = v178
		v673 = v250
		goto L41
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v296 = int32(0)
	if v246 == v296 {
		v668 = v176
		v669 = v177
		v670 = v178
		v673 = v296
		goto L41
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errcode(m, int32(50463045))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v266
	F_errmsg(m, int32(_a_F_vacuum_6), v28+int32(32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errfinish(m, int32(_a_F_vacuum_3), int32(952), int32(_a_F_vacuum_7))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v668 = v176
	v669 = v177
	v670 = v178
	v673 = v250
	goto L41
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errcode(m, int32(50463045))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v312
	F_errmsg(m, int32(_a_F_vacuum_8), v28+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errfinish(m, int32(_a_F_vacuum_3), int32(957), int32(_a_F_vacuum_7))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v668 = v176
	v669 = v177
	v670 = v178
	v673 = v296
	goto L41
L62:
	;
	if v352 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v352)+16))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v409 = v398 + v399
	v410 = F_vacuum_is_permitted_for_relation(m, v231, v409, v213)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L6
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v231
	F_errmsg_internal(m, int32(_a_F_vacuum_9), v28+int32(48))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errfinish(m, int32(_a_F_vacuum_3), int32(967), int32(_a_F_vacuum_7))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L3
L69:
	;
	if v410 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v413 = int32(_a_F_vacuum_5)
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v428 = F_makeVacuumRelation(m, v418, v231, v417)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	v444 = v177
	v445 = int32(0)
	goto L72
L72:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+16)))
	if v213&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v440 = F_lappend(m, int32(0), v428)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v414
	v444 = v440
	v445 = v440
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_UnlockRelationOid(m, v231, int32(1))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L6
	} else {
		goto L99
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_ReleaseCatCache(m, v352)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L6
	} else {
		goto L86
	}
L77:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+119)))
	if v448&int32(1)|base.B2i32(v455 != int32(112)) != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v470 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v470 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v473
	F_errmsg(m, int32(_a_F_vacuum_10), v28-int32(-64))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_ReleaseCatCache(m, v352)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L6
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	F_errfinish(m, int32(_a_F_vacuum_3), int32(993), int32(_a_F_vacuum_7))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v631 = v176
	v636 = v445
	goto L75
L86:
	;
	if v448&int32(1) == int32(0) {
		v631 = v176
		v636 = v445
		goto L75
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v539 = int32(0)
	v541 = F_find_all_inheritors(m, v231, v539, v539)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	if v541 == int32(0) {
		v631 = v176
		v636 = v445
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v545 = int32(0)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v546 <= v545 {
		v631 = v176
		v636 = v445
		goto L75
	} else {
		goto L90
	}
L90:
	;
	v561 = v176
	v566 = v445
	v569 = v545
	v570 = v546
	goto L91
L91:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v574+v569<<(uint(int32(2))%32))))
	if v231 != v578 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v631 = v611
	v636 = v612
	goto L75
L93:
	;
	v580 = int32(_a_F_vacuum_5)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v595 = F_makeVacuumRelation(m, int32(0), v578, v584)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L96
	}
L94:
	;
	v611 = v561
	v612 = v566
	v613 = v570
	goto L95
L95:
	;
	v617 = v569 + int32(1)
	if v617 < v613 {
		v561 = v611
		v566 = v612
		v569 = v617
		v570 = v613
		goto L91
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v561
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v178
	v606 = F_lappend(m, v566, v595)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v581
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	v611 = v606
	v612 = v606
	v613 = v610
	goto L95
L98:
	;
	goto L92
L99:
	;
	v668 = v631
	v669 = v444
	v670 = v178
	v673 = v636
	goto L41
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v682
	v699 = v179 + int32(1)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v699 < v700 {
		v172 = v694
		v176 = v668
		v177 = v669
		v178 = v670
		v179 = v699
		v182 = v694
		goto L39
	} else {
		goto L101
	}
L101:
	;
	goto L40
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v724 = int32(0)
	v726 = F_table_beginscan_catalog(m, v713, v724, v724)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v737 = F_heap_getnext(m, v726)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v739 = int32(0)
	if v737 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v750 = v41
	v751 = v42
	v755 = v739
	v759 = v737
	goto L108
L106:
	;
	v849 = v41
	v850 = v42
	v854 = v739
	goto L107
L107:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+188))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	m.T0[v866].(func(*base.Module, int32))(m, v726)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L6
	} else {
		goto L118
	}
L108:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v759)+16))
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+22)))
	v767 = v765 + v766
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+119)))
	v770 = v768 - int32(109)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v770))|base.B2i32(int32(1)<<(uint(v770)%32)&int32(41) == int32(0)) != 0 {
		v824 = v751
		v825 = v755
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v849 = v837
	v850 = v824
	v854 = v825
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v824
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v837 = F_heap_getnext(m, v726)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L6
	} else {
		goto L116
	}
L111:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v790 = F_vacuum_is_permitted_for_relation(m, v780, v767, v155)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	if v790 == int32(0) {
		v824 = v751
		v825 = v755
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v794 = int32(_a_F_vacuum_5)
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v807 = int32(0)
	v809 = F_makeVacuumRelation(m, v807, v780, v807)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v751
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	v820 = F_lappend(m, v755, v809)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[2])) = v795
	v824 = v820
	v825 = v820
	goto L110
L116:
	;
	if v837 != 0 {
		v750 = v837
		v751 = v824
		v755 = v825
		v759 = v837
		goto L108
	} else {
		goto L117
	}
L117:
	;
	goto L109
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v45
	F_relation_close(m, v713, int32(1))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	v898 = v39
	v900 = v849
	v901 = v850
	v902 = v43
	v903 = v44
	v904 = v45
	v905 = v854
	goto L29
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+298)) = uint8(v931)
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+298)))
	if v933 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L121:
	;
	v920 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[3]))
	if v920 == int32(4) {
		v931 = v915
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+299)))
	if v924 != 0 {
		v931 = int32(0)
		goto L120
	} else {
		goto L123
	}
L123:
	;
	if v905 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v925 = int32(1)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v905)+4))
	if v925 < v926 {
		v931 = v925
		goto L120
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v931 = int32(0)
	goto L120
L127:
	;
	goto L126
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v904
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[4]))
	goto L131
L129:
	;
	goto L130
L130:
	;
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[5]))
	v975 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[6]))
	goto L137
L131:
	;
	if v946 != int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v904
	F_PopActiveSnapshot(m)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L6
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v898
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v904
	F_CommitTransactionCommand(m)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L6
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	goto L130
L137:
	;
	v977 = v28 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v977)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v977))) = v28 + int32(68)
	goto L140
L138:
	;
	v989 = v975
	v990 = v973
	v991 = v898
	v993 = v900
	v994 = v901
	v995 = v902
	v996 = v903
	v997 = v904
	v998 = v905
	v1002 = int32(0)
	goto L7
L140:
	;
	goto L138
L141:
	;
	v1011 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1011)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[6])) = v28 + int32(128)
	v1018 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[7])) = uint8(v1018)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L6
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[5])) = v990
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[6])) = v989
	v1540 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1540)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[8])) = uint8(v1540)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[7])) = uint8(v1540)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[9])) = v1540
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_pg_re_throw(m)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L6
	} else {
		goto L243
	}
L144:
	;
	v1032 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[10])) = v1032
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[9])) = v1032
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[11])) = v1032
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[12])) = v1032
	if v998 == v1032 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[6])) = v989
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[5])) = v990
	v1493 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[8])) = uint8(v1493)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[0])) = uint8(v1493)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[7])) = uint8(v1493)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[9])) = v1493
	v1504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+298)))
	if v1504 != 0 {
		goto L235
	} else {
		goto L236
	}
L146:
	;
	v1045 = int32(0)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v998)+4))
	if v1046 <= v1045 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v1068 = v1045
	goto L148
L148:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v998)+12))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1074+v1068<<(uint(int32(2))%32))))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1079&int32(1) != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	goto L145
L150:
	;
	v1460 = v1068 + int32(1)
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v998)+4))
	if v1460 < v1461 {
		v1068 = v1460
		goto L148
	} else {
		goto L234
	}
L151:
	;
	v1451 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[7])) = uint8(v1451)
	goto L150
L152:
	;
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+120)) = v1082
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+112)) = v1084
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+104)) = v1086
	v1088 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+96)) = v1088
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+88)) = v1090
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+80)) = v1092
	v1094 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+72)) = v1094
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+8))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	v1109 = F_vacuum_rel(m, v1096, v1097, v28+int32(72), l2)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L6
	} else {
		goto L155
	}
L153:
	;
	v1114 = v1079
	goto L154
L154:
	;
	if v1114&int32(2) == int32(0) {
		goto L151
	} else {
		goto L157
	}
L155:
	;
	if v1109 == int32(0) {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1114 = v1113
	goto L154
L157:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+298)))
	if v1119 == int32(1) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_StartTransactionCommand(m)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L6
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+12))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+299)))
	v1168 = m.G0
	v1170 = v1168 - int32(32)
	m.G0 = v1170
	v1172 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+28)) = v1172
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+24)) = v1172
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[13])) = l2
	v1180 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[14]))
	if v1180 != 0 {
		goto L165
	} else {
		goto L166
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	v1142 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_PushActiveSnapshot(m, v1142)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	m.G0 = v1170 + int32(32)
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+298)))
	if v1397 == int32(1) {
		goto L227
	} else {
		goto L228
	}
L165:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L6
	} else {
		goto L168
	}
L166:
	;
	v1184 = v1176
	goto L167
L167:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1193 = F_vacuum_open_relation(m, v1157, v1156, v1184&int32(-2), int32(base.Ui32(v1187^int32(-1))>>(uint(int32(31))%32)), int32(4))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L6
	} else {
		goto L169
	}
L168:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1184 = v1183
	goto L167
L169:
	;
	if v1193 == int32(0) {
		goto L164
	} else {
		goto L170
	}
L170:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+56))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1202 = F_vacuum_is_permitted_for_relation(m, v1197, v1198, v1199&int32(-2))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L6
	} else {
		goto L171
	}
L171:
	;
	if v1202 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	F_relation_close(m, v1193, int32(4))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L6
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+118)))
	if v1210 != int32(116) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	goto L164
L176:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+56))
	if v1217 == int32(2619) {
		goto L180
	} else {
		goto L181
	}
L177:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193)+24)))
	if v1213 != 0 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	F_relation_close(m, v1193, int32(4))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L6
	} else {
		goto L179
	}
L179:
	;
	goto L164
L180:
	;
	F_relation_close(m, v1193, int32(4))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L6
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+119)))
	switch v1223 - int32(102) {
	case 0:
		goto L187
	default:
		goto L186
	case 7, 12:
		goto L185
	case 10:
		goto L184
	}
L183:
	;
	goto L164
L184:
	;
	v1291 = v1176&int32(4) + int32(13)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+56))
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[15]))
	if v1296 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+28)) = int32(501)
	v1284 = F_RelationGetNumberOfBlocksInFork(m, v1193, int32(0))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L6
	} else {
		goto L208
	}
L186:
	;
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1257&int32(1) != 0 {
		goto L201
	} else {
		goto L202
	}
L187:
	;
	v1227 = F_GetFdwRoutineForRelation(m, v1193, int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+128))
	if v1229 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1234 = m.T0[v1229].(func(*base.Module, int32, int32, int32) int32)(m, v1193, v1170+int32(28), v1170+int32(24))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L6
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1238 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L6
	} else {
		goto L194
	}
L192:
	;
	if v1234 != 0 {
		goto L184
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	if v1238 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+16)) = v1240 + int32(4)
	F_errmsg(m, int32(_a_F_vacuum_11), v1170+int32(16))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L6
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	F_relation_close(m, v1193, int32(4))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L6
	} else {
		goto L200
	}
L198:
	;
	F_errfinish(m, int32(_a_F_vacuum_12), int32(216), int32(_a_F_vacuum_13))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	goto L164
L201:
	;
	F_relation_close(m, v1193, int32(4))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L6
	} else {
		goto L207
	}
L202:
	;
	v1262 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L6
	} else {
		goto L203
	}
L203:
	;
	if v1262 == int32(0) {
		goto L201
	} else {
		goto L204
	}
L204:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1170))) = v1266 + int32(4)
	F_errmsg(m, int32(_a_F_vacuum_14), v1170)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L6
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_vacuum_12), int32(233), int32(_a_F_vacuum_13))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L6
	} else {
		goto L206
	}
L206:
	;
	goto L201
L207:
	;
	goto L164
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+24)) = v1284
	goto L184
L209:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332)+119)))
	if v1333 != int32(112) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	goto L209
L211:
	;
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[16])))
	if v1300&int32(1) == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1305 = int32(_a_F_vacuum_15)
	v1307 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[17]))
	v1308 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[17])) = v1307 + v1308
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1296)))
	*(*int32)(unsafe.Add(mBase, uint32(v1296))) = v1311 + v1308
	*(*int32)(unsafe.Add(mBase, uint32(v1296)+220)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1296)+224)) = v1293
	base.MemoryFill(m, v1296+int32(232), int32(0), int32(160))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1296)))
	*(*int32)(unsafe.Add(mBase, uint32(v1296))) = v1322 + v1308
	v1328 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[17])) = v1328 - v1308
	goto L210
L213:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+28))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+24))
	F_do_analyze_rel(m, v1193, l1, v1155, v1336, v1337, int32(0), v1167, v1291)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L6
	} else {
		goto L216
	}
L214:
	;
	v1342 = v1332
	goto L215
L215:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+126)))
	if v1343 == int32(1) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+48))
	v1342 = v1341
	goto L215
L217:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+28))
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1170)+24))
	F_do_analyze_rel(m, v1193, l1, v1155, v1346, v1347, int32(1), v1167, v1291)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L6
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	F_relation_close(m, v1193, int32(0))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L6
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[15]))
	if v1356 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L164
L223:
	;
	goto L222
L224:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_vacuum[16])))
	if v1360&int32(1) == int32(0) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+220))
	if v1365 == int32(0) {
		goto L223
	} else {
		goto L226
	}
L226:
	;
	v1368 = int32(_a_F_vacuum_15)
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[17]))
	v1371 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[17])) = v1370 + v1371
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1356)))
	*(*int32)(unsafe.Add(mBase, uint32(v1356))) = v1374 + v1371
	v1378 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1356)+220)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1356)+224)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v1356))) = v1374 + int32(2)
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum[17])) = v1388 - v1371
	goto L223
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L6
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L6
	} else {
		goto L233
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L6
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	goto L151
L233:
	;
	goto L151
L234:
	;
	goto L149
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_StartTransactionCommand(m)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L6
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1516&int32(513) == int32(1) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	goto L237
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+304)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v28)+300)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v28)+308)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v28)+316)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v28)+320)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v28)+324)) = v995
	*(*int32)(unsafe.Add(mBase, uint32(v28)+328)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v28)+332)) = v997
	F_vac_update_datfrozenxid(m)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L6
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	m.G0 = v28 + int32(336)
	return
L242:
	;
	goto L241
L243:
	;
	goto L5
L244:
	;
	v1592 = int32(v1588)
	m.G0 = v28
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1592)+4))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1592)))
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1595)))
	if v28+int32(68) == v1598 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	m.ExcPending = 1
	goto L253
L246:
	;
	if v1602 != 0 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	v1602 = v1600
	goto L249
L248:
	;
	v1602 = int32(0)
	goto L249
L249:
	;
	goto L246
L250:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v28)+332))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v28)+328))
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v28)+324))
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v28)+320))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v28)+316))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v28)+312))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v28)+308))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v28)+304))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v28)+300))
	v37 = v1610
	v38 = v1611
	v39 = v1606
	v40 = v1609
	v41 = v1608
	v42 = v1607
	v43 = v1605
	v44 = v1604
	v45 = v1603
	v46 = v1602
	v50 = v1594
	goto L1
L251:
	;
	goto L252
L252:
	;
	F___wasm_longjmp(m, v1595, v1594)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	return
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_get_cutoffs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 float64
	_ = v201
	var v204 float64
	_ = v204
	var v206 float64
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v22
	v24 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
		v29 = F_GetOldestMultiXactId(m)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v29
			v32 = F_ReadNextFullTransactionId(m)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v35 = F_ReadNextMultiXactId(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = F_MultiXactMemberFreezeThreshold(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v35 == v37 {
							v41 = int32(1)
						} else {
							v41 = v35 - v37
						}
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v43 = int32(3)
						v44 = base.I32_wrap_i64(v32)
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
						v47 = v44 - v46
						if base.Ui32(v47) <= base.Ui32(v43) {
							v50 = v43
						} else {
							v50 = v47
						}
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v50))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
							v62 = base.B2i32(base.Ui32(v42) < base.Ui32(v50))
						} else {
							v62 = int32(base.Ui32(v42-v50) >> (uint(int32(31)) % 32))
						}
						if v62 == int32(0) {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
								if v17 < int32(0) {
									v114 = v111
								} else {
									v114 = v17
								}
								v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
								v118 = base.I32_div_s(v116, int32(2))
								if v114 < v118 {
									v120 = v114
								} else {
									v120 = v118
								}
								v121 = v44 - v120
								if base.Ui32(v121) <= base.Ui32(int32(3)) {
									v124 = int32(3)
								} else {
									v124 = v121
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
									v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
								} else {
									v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
								}
								if v138 != 0 {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
								} else {
								}
								v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
								if v16 < int32(0) {
									v146 = v143
								} else {
									v146 = v16
								}
								v148 = base.I32_div_s(v37, int32(2))
								if v146 < v148 {
									v150 = v146
								} else {
									v150 = v148
								}
								if v35 == v150 {
									v153 = int32(1)
								} else {
									v153 = v35 - v150
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
								v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
									v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
								} else {
								}
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
								if v15 < int32(0) {
									v167 = v164
								} else {
									v167 = v15
								}
								v168 = base.F64_convert_i32_s(v167)
								v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
								v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
								if base.F64_lt(v168, v173) != 0 {
									v175 = v168
								} else {
									v175 = v173
								}
								v177 = v44 - base.I32_trunc_sat_f64_s(v175)
								if base.Ui32(v177) <= base.Ui32(int32(3)) {
									v180 = int32(3)
								} else {
									v180 = v177
								}
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
									v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
								} else {
									v192 = base.B2i32(v161-v180 <= int32(0))
								}
								if v192 != 0 {
									v217 = int32(1)
								} else {
									v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
									if v14 < int32(0) {
										v200 = v197
									} else {
										v200 = v14
									}
									v201 = base.F64_convert_i32_s(v200)
									v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
									if base.F64_lt(v201, v204) != 0 {
										v206 = v201
									} else {
										v206 = v204
									}
									v207 = base.I32_trunc_sat_f64_s(v206)
									if v35 == v207 {
										v210 = int32(1)
									} else {
										v210 = v35 - v207
									}
									v217 = base.B2i32(v194-v210 <= int32(0))
								}
								return v217
							} else {
								v92 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 == int32(0) {
										v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
										if v17 < int32(0) {
											v114 = v111
										} else {
											v114 = v17
										}
										v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
										v118 = base.I32_div_s(v116, int32(2))
										if v114 < v118 {
											v120 = v114
										} else {
											v120 = v118
										}
										v121 = v44 - v120
										if base.Ui32(v121) <= base.Ui32(int32(3)) {
											v124 = int32(3)
										} else {
											v124 = v121
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
											v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
										} else {
											v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
										}
										if v138 != 0 {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
										} else {
										}
										v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
										if v16 < int32(0) {
											v146 = v143
										} else {
											v146 = v16
										}
										v148 = base.I32_div_s(v37, int32(2))
										if v146 < v148 {
											v150 = v146
										} else {
											v150 = v148
										}
										if v35 == v150 {
											v153 = int32(1)
										} else {
											v153 = v35 - v150
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
										} else {
										}
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
										if v15 < int32(0) {
											v167 = v164
										} else {
											v167 = v15
										}
										v168 = base.F64_convert_i32_s(v167)
										v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
										v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
										if base.F64_lt(v168, v173) != 0 {
											v175 = v168
										} else {
											v175 = v173
										}
										v177 = v44 - base.I32_trunc_sat_f64_s(v175)
										if base.Ui32(v177) <= base.Ui32(int32(3)) {
											v180 = int32(3)
										} else {
											v180 = v177
										}
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
											v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
										} else {
											v192 = base.B2i32(v161-v180 <= int32(0))
										}
										if v192 != 0 {
											v217 = int32(1)
										} else {
											v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
											if v14 < int32(0) {
												v200 = v197
											} else {
												v200 = v14
											}
											v201 = base.F64_convert_i32_s(v200)
											v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
											if base.F64_lt(v201, v204) != 0 {
												v206 = v201
											} else {
												v206 = v204
											}
											v207 = base.I32_trunc_sat_f64_s(v206)
											if v35 == v207 {
												v210 = int32(1)
											} else {
												v210 = v35 - v207
											}
											v217 = base.B2i32(v194-v210 <= int32(0))
										}
										return v217
									} else {
										F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_0), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(_a_F_vacuum_get_cutoffs_1), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1190), int32(_a_F_vacuum_get_cutoffs_3))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
													if v17 < int32(0) {
														v114 = v111
													} else {
														v114 = v17
													}
													v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
													v118 = base.I32_div_s(v116, int32(2))
													if v114 < v118 {
														v120 = v114
													} else {
														v120 = v118
													}
													v121 = v44 - v120
													if base.Ui32(v121) <= base.Ui32(int32(3)) {
														v124 = int32(3)
													} else {
														v124 = v121
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
														v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
													} else {
														v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
													}
													if v138 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
													} else {
													}
													v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
													if v16 < int32(0) {
														v146 = v143
													} else {
														v146 = v16
													}
													v148 = base.I32_div_s(v37, int32(2))
													if v146 < v148 {
														v150 = v146
													} else {
														v150 = v148
													}
													if v35 == v150 {
														v153 = int32(1)
													} else {
														v153 = v35 - v150
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
													} else {
													}
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
													if v15 < int32(0) {
														v167 = v164
													} else {
														v167 = v15
													}
													v168 = base.F64_convert_i32_s(v167)
													v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
													v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
													if base.F64_lt(v168, v173) != 0 {
														v175 = v168
													} else {
														v175 = v173
													}
													v177 = v44 - base.I32_trunc_sat_f64_s(v175)
													if base.Ui32(v177) <= base.Ui32(int32(3)) {
														v180 = int32(3)
													} else {
														v180 = v177
													}
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
														v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
													} else {
														v192 = base.B2i32(v161-v180 <= int32(0))
													}
													if v192 != 0 {
														v217 = int32(1)
													} else {
														v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
														v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
														if v14 < int32(0) {
															v200 = v197
														} else {
															v200 = v14
														}
														v201 = base.F64_convert_i32_s(v200)
														v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
														if base.F64_lt(v201, v204) != 0 {
															v206 = v201
														} else {
															v206 = v204
														}
														v207 = base.I32_trunc_sat_f64_s(v206)
														if v35 == v207 {
															v210 = int32(1)
														} else {
															v210 = v35 - v207
														}
														v217 = base.B2i32(v194-v210 <= int32(0))
													}
													return v217
												}
											}
										}
									}
								}
							}
						} else {
							v67 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								if v67 == int32(0) {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
										v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
										if v17 < int32(0) {
											v114 = v111
										} else {
											v114 = v17
										}
										v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
										v118 = base.I32_div_s(v116, int32(2))
										if v114 < v118 {
											v120 = v114
										} else {
											v120 = v118
										}
										v121 = v44 - v120
										if base.Ui32(v121) <= base.Ui32(int32(3)) {
											v124 = int32(3)
										} else {
											v124 = v121
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
											v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
										} else {
											v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
										}
										if v138 != 0 {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
										} else {
										}
										v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
										if v16 < int32(0) {
											v146 = v143
										} else {
											v146 = v16
										}
										v148 = base.I32_div_s(v37, int32(2))
										if v146 < v148 {
											v150 = v146
										} else {
											v150 = v148
										}
										if v35 == v150 {
											v153 = int32(1)
										} else {
											v153 = v35 - v150
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
										} else {
										}
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
										if v15 < int32(0) {
											v167 = v164
										} else {
											v167 = v15
										}
										v168 = base.F64_convert_i32_s(v167)
										v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
										v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
										if base.F64_lt(v168, v173) != 0 {
											v175 = v168
										} else {
											v175 = v173
										}
										v177 = v44 - base.I32_trunc_sat_f64_s(v175)
										if base.Ui32(v177) <= base.Ui32(int32(3)) {
											v180 = int32(3)
										} else {
											v180 = v177
										}
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
											v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
										} else {
											v192 = base.B2i32(v161-v180 <= int32(0))
										}
										if v192 != 0 {
											v217 = int32(1)
										} else {
											v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
											if v14 < int32(0) {
												v200 = v197
											} else {
												v200 = v14
											}
											v201 = base.F64_convert_i32_s(v200)
											v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
											if base.F64_lt(v201, v204) != 0 {
												v206 = v201
											} else {
												v206 = v204
											}
											v207 = base.I32_trunc_sat_f64_s(v206)
											if v35 == v207 {
												v210 = int32(1)
											} else {
												v210 = v35 - v207
											}
											v217 = base.B2i32(v194-v210 <= int32(0))
										}
										return v217
									} else {
										v92 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											if v92 == int32(0) {
												v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
												if v17 < int32(0) {
													v114 = v111
												} else {
													v114 = v17
												}
												v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
												v118 = base.I32_div_s(v116, int32(2))
												if v114 < v118 {
													v120 = v114
												} else {
													v120 = v118
												}
												v121 = v44 - v120
												if base.Ui32(v121) <= base.Ui32(int32(3)) {
													v124 = int32(3)
												} else {
													v124 = v121
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
													v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
												} else {
													v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
												}
												if v138 != 0 {
													v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
												} else {
												}
												v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
												if v16 < int32(0) {
													v146 = v143
												} else {
													v146 = v16
												}
												v148 = base.I32_div_s(v37, int32(2))
												if v146 < v148 {
													v150 = v146
												} else {
													v150 = v148
												}
												if v35 == v150 {
													v153 = int32(1)
												} else {
													v153 = v35 - v150
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
												v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
												} else {
												}
												v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
												if v15 < int32(0) {
													v167 = v164
												} else {
													v167 = v15
												}
												v168 = base.F64_convert_i32_s(v167)
												v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
												v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
												if base.F64_lt(v168, v173) != 0 {
													v175 = v168
												} else {
													v175 = v173
												}
												v177 = v44 - base.I32_trunc_sat_f64_s(v175)
												if base.Ui32(v177) <= base.Ui32(int32(3)) {
													v180 = int32(3)
												} else {
													v180 = v177
												}
												if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
													v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
												} else {
													v192 = base.B2i32(v161-v180 <= int32(0))
												}
												if v192 != 0 {
													v217 = int32(1)
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
													v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
													if v14 < int32(0) {
														v200 = v197
													} else {
														v200 = v14
													}
													v201 = base.F64_convert_i32_s(v200)
													v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
													if base.F64_lt(v201, v204) != 0 {
														v206 = v201
													} else {
														v206 = v204
													}
													v207 = base.I32_trunc_sat_f64_s(v206)
													if v35 == v207 {
														v210 = int32(1)
													} else {
														v210 = v35 - v207
													}
													v217 = base.B2i32(v194-v210 <= int32(0))
												}
												return v217
											} else {
												F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_0), int32(0))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_vacuum_get_cutoffs_1), int32(0))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1190), int32(_a_F_vacuum_get_cutoffs_3))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int32(0)
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
															if v17 < int32(0) {
																v114 = v111
															} else {
																v114 = v17
															}
															v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
															v118 = base.I32_div_s(v116, int32(2))
															if v114 < v118 {
																v120 = v114
															} else {
																v120 = v118
															}
															v121 = v44 - v120
															if base.Ui32(v121) <= base.Ui32(int32(3)) {
																v124 = int32(3)
															} else {
																v124 = v121
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
															} else {
																v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
															}
															if v138 != 0 {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
															} else {
															}
															v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
															if v16 < int32(0) {
																v146 = v143
															} else {
																v146 = v16
															}
															v148 = base.I32_div_s(v37, int32(2))
															if v146 < v148 {
																v150 = v146
															} else {
																v150 = v148
															}
															if v35 == v150 {
																v153 = int32(1)
															} else {
																v153 = v35 - v150
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
															v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
															if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
															} else {
															}
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
															v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
															if v15 < int32(0) {
																v167 = v164
															} else {
																v167 = v15
															}
															v168 = base.F64_convert_i32_s(v167)
															v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
															v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
															if base.F64_lt(v168, v173) != 0 {
																v175 = v168
															} else {
																v175 = v173
															}
															v177 = v44 - base.I32_trunc_sat_f64_s(v175)
															if base.Ui32(v177) <= base.Ui32(int32(3)) {
																v180 = int32(3)
															} else {
																v180 = v177
															}
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
															} else {
																v192 = base.B2i32(v161-v180 <= int32(0))
															}
															if v192 != 0 {
																v217 = int32(1)
															} else {
																v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
																if v14 < int32(0) {
																	v200 = v197
																} else {
																	v200 = v14
																}
																v201 = base.F64_convert_i32_s(v200)
																v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																if base.F64_lt(v201, v204) != 0 {
																	v206 = v201
																} else {
																	v206 = v204
																}
																v207 = base.I32_trunc_sat_f64_s(v206)
																if v35 == v207 {
																	v210 = int32(1)
																} else {
																	v210 = v35 - v207
																}
																v217 = base.B2i32(v194-v210 <= int32(0))
															}
															return v217
														}
													}
												}
											}
										}
									}
								} else {
									F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_4), int32(0))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(_a_F_vacuum_get_cutoffs_1), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1185), int32(_a_F_vacuum_get_cutoffs_3))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
													if v17 < int32(0) {
														v114 = v111
													} else {
														v114 = v17
													}
													v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
													v118 = base.I32_div_s(v116, int32(2))
													if v114 < v118 {
														v120 = v114
													} else {
														v120 = v118
													}
													v121 = v44 - v120
													if base.Ui32(v121) <= base.Ui32(int32(3)) {
														v124 = int32(3)
													} else {
														v124 = v121
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
														v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
													} else {
														v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
													}
													if v138 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
													} else {
													}
													v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
													if v16 < int32(0) {
														v146 = v143
													} else {
														v146 = v16
													}
													v148 = base.I32_div_s(v37, int32(2))
													if v146 < v148 {
														v150 = v146
													} else {
														v150 = v148
													}
													if v35 == v150 {
														v153 = int32(1)
													} else {
														v153 = v35 - v150
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
													} else {
													}
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
													if v15 < int32(0) {
														v167 = v164
													} else {
														v167 = v15
													}
													v168 = base.F64_convert_i32_s(v167)
													v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
													v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
													if base.F64_lt(v168, v173) != 0 {
														v175 = v168
													} else {
														v175 = v173
													}
													v177 = v44 - base.I32_trunc_sat_f64_s(v175)
													if base.Ui32(v177) <= base.Ui32(int32(3)) {
														v180 = int32(3)
													} else {
														v180 = v177
													}
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
														v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
													} else {
														v192 = base.B2i32(v161-v180 <= int32(0))
													}
													if v192 != 0 {
														v217 = int32(1)
													} else {
														v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
														v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
														if v14 < int32(0) {
															v200 = v197
														} else {
															v200 = v14
														}
														v201 = base.F64_convert_i32_s(v200)
														v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
														if base.F64_lt(v201, v204) != 0 {
															v206 = v201
														} else {
															v206 = v204
														}
														v207 = base.I32_trunc_sat_f64_s(v206)
														if v35 == v207 {
															v210 = int32(1)
														} else {
															v210 = v35 - v207
														}
														v217 = base.B2i32(v194-v210 <= int32(0))
													}
													return v217
												} else {
													v92 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														if v92 == int32(0) {
															v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
															if v17 < int32(0) {
																v114 = v111
															} else {
																v114 = v17
															}
															v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
															v118 = base.I32_div_s(v116, int32(2))
															if v114 < v118 {
																v120 = v114
															} else {
																v120 = v118
															}
															v121 = v44 - v120
															if base.Ui32(v121) <= base.Ui32(int32(3)) {
																v124 = int32(3)
															} else {
																v124 = v121
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
															} else {
																v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
															}
															if v138 != 0 {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
															} else {
															}
															v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
															if v16 < int32(0) {
																v146 = v143
															} else {
																v146 = v16
															}
															v148 = base.I32_div_s(v37, int32(2))
															if v146 < v148 {
																v150 = v146
															} else {
																v150 = v148
															}
															if v35 == v150 {
																v153 = int32(1)
															} else {
																v153 = v35 - v150
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
															v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
															if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
															} else {
															}
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
															v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
															if v15 < int32(0) {
																v167 = v164
															} else {
																v167 = v15
															}
															v168 = base.F64_convert_i32_s(v167)
															v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
															v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
															if base.F64_lt(v168, v173) != 0 {
																v175 = v168
															} else {
																v175 = v173
															}
															v177 = v44 - base.I32_trunc_sat_f64_s(v175)
															if base.Ui32(v177) <= base.Ui32(int32(3)) {
																v180 = int32(3)
															} else {
																v180 = v177
															}
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
															} else {
																v192 = base.B2i32(v161-v180 <= int32(0))
															}
															if v192 != 0 {
																v217 = int32(1)
															} else {
																v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
																if v14 < int32(0) {
																	v200 = v197
																} else {
																	v200 = v14
																}
																v201 = base.F64_convert_i32_s(v200)
																v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																if base.F64_lt(v201, v204) != 0 {
																	v206 = v201
																} else {
																	v206 = v204
																}
																v207 = base.I32_trunc_sat_f64_s(v206)
																if v35 == v207 {
																	v210 = int32(1)
																} else {
																	v210 = v35 - v207
																}
																v217 = base.B2i32(v194-v210 <= int32(0))
															}
															return v217
														} else {
															F_errmsg(m, int32(_a_F_vacuum_get_cutoffs_0), int32(0))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(_a_F_vacuum_get_cutoffs_1), int32(0))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_vacuum_get_cutoffs_2), int32(1190), int32(_a_F_vacuum_get_cutoffs_3))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v111 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[1]))
																		if v17 < int32(0) {
																			v114 = v111
																		} else {
																			v114 = v17
																		}
																		v116 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
																		v118 = base.I32_div_s(v116, int32(2))
																		if v114 < v118 {
																			v120 = v114
																		} else {
																			v120 = v118
																		}
																		v121 = v44 - v120
																		if base.Ui32(v121) <= base.Ui32(int32(3)) {
																			v124 = int32(3)
																		} else {
																			v124 = v121
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
																		v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																			v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
																		} else {
																			v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
																		}
																		if v138 != 0 {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
																		} else {
																		}
																		v143 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[2]))
																		if v16 < int32(0) {
																			v146 = v143
																		} else {
																			v146 = v16
																		}
																		v148 = base.I32_div_s(v37, int32(2))
																		if v146 < v148 {
																			v150 = v146
																		} else {
																			v150 = v148
																		}
																		if v35 == v150 {
																			v153 = int32(1)
																		} else {
																			v153 = v35 - v150
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
																		v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																		if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																			v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
																		} else {
																		}
																		v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
																		v164 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[3]))
																		if v15 < int32(0) {
																			v167 = v164
																		} else {
																			v167 = v15
																		}
																		v168 = base.F64_convert_i32_s(v167)
																		v170 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[0]))
																		v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
																		if base.F64_lt(v168, v173) != 0 {
																			v175 = v168
																		} else {
																			v175 = v173
																		}
																		v177 = v44 - base.I32_trunc_sat_f64_s(v175)
																		if base.Ui32(v177) <= base.Ui32(int32(3)) {
																			v180 = int32(3)
																		} else {
																			v180 = v177
																		}
																		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v180))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																			v192 = base.B2i32(base.Ui32(v161) <= base.Ui32(v180))
																		} else {
																			v192 = base.B2i32(v161-v180 <= int32(0))
																		}
																		if v192 != 0 {
																			v217 = int32(1)
																		} else {
																			v194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																			v197 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_get_cutoffs[4]))
																			if v14 < int32(0) {
																				v200 = v197
																			} else {
																				v200 = v14
																			}
																			v201 = base.F64_convert_i32_s(v200)
																			v204 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																			if base.F64_lt(v201, v204) != 0 {
																				v206 = v201
																			} else {
																				v206 = v204
																			}
																			v207 = base.I32_trunc_sat_f64_s(v206)
																			if v35 == v207 {
																				v210 = int32(1)
																			} else {
																				v210 = v35 - v207
																			}
																			v217 = base.B2i32(v194-v210 <= int32(0))
																		}
																		return v217
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
				}
			}
		}
	}
}
func F_vacuum_rel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 float64
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v27
	F_StartTransactionCommand(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v33&int32(16) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[0]))
	v43 = F_LWLockAcquire(m, v39+int32(512), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v73 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[1]))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)))
	v49 = v47 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v51 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = v47 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)) = uint8(v55)
	v57 = v55
	goto L9
L8:
	;
	v57 = v49
	goto L9
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[2]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v61))) = uint8(v57)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[0]))
	F_LWLockRelease(m, v65+int32(512))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	F_PushActiveSnapshot(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[3]))
	if v78 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v81&int32(16) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v13 + int32(80)
	return v320
L18:
	;
	v91 = int32(8)
	goto L20
L19:
	;
	v91 = int32(4)
	goto L20
L20:
	;
	v92 = F_vacuum_open_relation(m, l0, l1, v81, int32(base.Ui32(v82^int32(-1))>>(uint(int32(31))%32)), v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v92 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L115
	}
L25:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L114
	}
L26:
	;
	v96 = v94
	goto L28
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+56))
	v96 = v95
	goto L28
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v101 = F_vacuum_is_permitted_for_relation(m, v96, v97, v98&int32(-3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v101 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+119)))
	v108 = v106 - int32(109)
	if int32(1)<<(uint(v108)%32)&int32(169) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = base.B2i32(base.Ui32(v108) <= base.Ui32(int32(7)))
	goto L33
L32:
	;
	v116 = int32(0)
	goto L33
L33:
	;
	if v116 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v119 = int32(0)
	v122 = F_errstart(m, int32(19), v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+118)))
	if v142 == int32(116) {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v122 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v124 + int32(4)
	F_errmsg(m, int32(_a_F_vacuum_rel_0), v13)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(_a_F_vacuum_rel_1), int32(2144), int32(_a_F_vacuum_rel_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v320 = v119
	goto L17
L46:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)))
	if v145 == int32(0) {
		goto L25
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v106 == int32(112) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v320 = int32(1)
	goto L17
L51:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v92)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v156
	F_LockRelationIdForSession(m, v13+int32(72), v91)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v162 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v165 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v179 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	v167 = int32(2)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v168 == v167 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v176 = int32(1)
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v176
	goto L60
L64:
	;
	v171 = int32(3)
	goto L66
L65:
	;
	v171 = v167
	goto L66
L66:
	;
	if v168 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v173 = v171
	goto L69
L68:
	;
	v173 = int32(1)
	goto L69
L69:
	;
	v176 = v173
	goto L63
L70:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v189 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v179)+120))
	if base.F64_ge(v182, float64(0)) == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2)+40)) = v182
	goto L70
L73:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v194 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v212 = int32(80)
	if base.B2i32(v207&int32(128) == v206)|base.B2i32(v207&v212 == v212) == v206 {
		goto L85
	} else {
		goto L86
	}
L76:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+117)))
	if v198 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v201 = int32(_a_F_vacuum_rel_3)
	goto L78
L78:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v202 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v199 = v194 + int32(116)
	goto L81
L80:
	;
	v199 = int32(_a_F_vacuum_rel_3)
	goto L81
L81:
	;
	v201 = v199
	goto L78
L82:
	;
	v203 = int32(3)
	goto L84
L83:
	;
	v203 = int32(2)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v203
	goto L75
L85:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+112))
	v221 = v220
	goto L87
L86:
	;
	v221 = v206
	goto L87
L87:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v227
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)))) = v230
	goto L88
L88:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+80))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5])) = v234 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[4])) = v233
	goto L89
L89:
	;
	v242 = int32(_a_F_vacuum_rel_4)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[6]))
	v246 = v244 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[6])) = v246
	goto L90
L90:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v250&int32(64) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	F_AtEOXact_GUC(m, int32(0), v246)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L101
	}
L93:
	;
	if v250&int32(16) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v270 = v92
	goto L92
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(base.Ui32(v250)>>(uint(int32(2))%32)) & int32(1)
	v260 = int32(0)
	F_cluster_rel(m, v92, v260, v13+int32(4))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v92)+188))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+128))
	m.T0[v267].(func(*base.Module, int32, int32, int32))(m, v92, l2, l3)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v270 = v260
	goto L92
L100:
	;
	goto L95
L101:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[5])) = v275
	*(*int32)(unsafe.Add(mBase, _c_F_vacuum_rel[4])) = v274
	goto L102
L102:
	;
	if v270 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_relation_close(m, v270, int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v221 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l0
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v288 | int32(64)
	v295 = F_vacuum_rel(m, v221, int32(0), v13+int32(8), l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_UnlockRelationIdForSession(m, v13+int32(72), v91)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	goto L50
L114:
	;
	goto L24
L115:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v320 = int32(0)
	goto L17
}
