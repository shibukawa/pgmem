package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CreateStandaloneExprContext(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	v4 = F_palloc0(m, int32(88))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(388)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_CreateStandaloneExprContext[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v13
		v19 = F_AllocSetContextCreateInternal(m, v13, int32(_a_F_CreateStandaloneExprContext_0), int32(0), int32(_a_F_CreateStandaloneExprContext_1), int32(_a_F_CreateStandaloneExprContext_2))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v4)+32)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v4)+40)) = v21
			*(*int64)(unsafe.Add(mBase, uint32(v4)+76)) = v21
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+64)) = uint8(v30)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v21
			*(*uint8)(unsafe.Add(mBase, uint32(v4)+48)) = uint8(v30)
			return v4
		}
	}
}
func F_CreateTemplateTupleDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	v7 = F_palloc(m, l0*int32(108)+int32(28))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v14 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(2249)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v14
		return v7
	}
}
func F_create_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v212 float64
	_ = v212
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v250 float64
	_ = v250
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v581 int32
	_ = v581
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
	var v596 int32
	_ = v596
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v681 float64
	_ = v681
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 float64
	_ = v708
	var v710 float64
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v826 float64
	_ = v826
	var v849 int32
	_ = v849
	var v860 float64
	_ = v860
	var v886 int32
	_ = v886
	var v926 float64
	_ = v926
	var v927 int32
	_ = v927
	var v936 int32
	_ = v936
	var v942 float64
	_ = v942
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1027 float64
	_ = v1027
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 float64
	_ = v1051
	var v1052 float64
	_ = v1052
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1071 float64
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 float64
	_ = v1075
	var v1077 float64
	_ = v1077
	var v1078 float64
	_ = v1078
	var v1082 float64
	_ = v1082
	var v1085 float64
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1101 float64
	_ = v1101
	var v1124 float64
	_ = v1124
	var v1125 float64
	_ = v1125
	var v1134 float64
	_ = v1134
	var v1138 float64
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1144 float64
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 float64
	_ = v1150
	var v1151 float64
	_ = v1151
	var v1154 float64
	_ = v1154
	var v1157 float64
	_ = v1157
	var v1158 int64
	_ = v1158
	var v1163 float64
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 float64
	_ = v1218
	var v1219 float64
	_ = v1219
	var v1233 float64
	_ = v1233
	var v1254 float64
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int64
	_ = v1256
	var v1265 int32
	_ = v1265
	var v1272 int32
	_ = v1272
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 float64
	_ = v1316
	var v1317 float64
	_ = v1317
	var v1341 float64
	_ = v1341
	var v1352 float64
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1361 float64
	_ = v1361
	var v1362 float64
	_ = v1362
	var v1364 float64
	_ = v1364
	var v1366 float64
	_ = v1366
	var v1369 float64
	_ = v1369
	var v1372 float64
	_ = v1372
	var v1376 float64
	_ = v1376
	var v1385 float64
	_ = v1385
	var v1389 float64
	_ = v1389
	var v1394 float64
	_ = v1394
	var v1403 float64
	_ = v1403
	var v1407 float64
	_ = v1407
	var v1410 float64
	_ = v1410
	var v1416 float64
	_ = v1416
	var v1417 float64
	_ = v1417
	var v1418 float64
	_ = v1418
	var v1427 float64
	_ = v1427
	var v1431 float64
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 float64
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 float64
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int64
	_ = v1442
	var v1462 float64
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1470 int32
	_ = v1470
	var v1478 float64
	_ = v1478
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1510 float64
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 float64
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1528 float64
	_ = v1528
	var v1553 float64
	_ = v1553
	var v1555 float64
	_ = v1555
	var v1564 float64
	_ = v1564
	var v1568 float64
	_ = v1568
	var v1582 float64
	_ = v1582
	var v1604 float64
	_ = v1604
	var v1606 float64
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 float64
	_ = v1608
	var v1620 float64
	_ = v1620
	var v1624 float64
	_ = v1624
	var v1625 float64
	_ = v1625
	var v1627 float64
	_ = v1627
	v12 = float64(0)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = l8
	v41 = F_palloc0(m, int32(112))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(1559073128750)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v53 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v50, l9, v37+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v53
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	v57 = l7 & v56
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)) = uint8(v57)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v60 != int32(1) {
		v68 = int32(0)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v70 = v68 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+21)) = uint8(v70)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v72
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+76)) = uint8(v77)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v81
	v84 = m.G0
	v86 = v84 + int32(-64)
	m.G0 = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+88))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l5)+32))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v100 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v64 != int32(1) {
		v68 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v68 = v67
	goto L4
L7:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v106)))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if int32(0) < v109 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v106 = v100 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v106 = v103 + int32(16)
	goto L7
L11:
	;
	v112 = base.F64_convert_i32_u(v109)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_hashjoin_path[0])))
	if v114 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v41)+104)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = v88
	v150 = base.F64_mul(base.F64_convert_i32_s(v89), base.F64_convert_i32_s(v88))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v151 != 0 {
		goto L24
	} else {
		goto L25
	}
L14:
	;
	v120 = base.F64_add(base.F64_mul(v112, float64(-0.3)), float64(1))
	if base.F64_gt(v120, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v126 = v112
	goto L16
L16:
	;
	v128 = float64(1e+100)
	v129 = base.F64_div(v107, v126)
	if base.F64_gt(v129, v128)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v129)&int64(9223372036854775807))) != 0 {
		v142 = v128
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v124 = v120
	goto L19
L18:
	;
	v124 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v126 = base.F64_add(v124, v112)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v142
	goto L13
L21:
	;
	v138 = float64(1)
	if base.F64_le(v129, v138) != 0 {
		v142 = v138
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v142 = base.F64_nearest(v129)
	goto L20
L23:
	;
	v1124 = float64(1e+100)
	v1125 = base.F64_mul(v95, v1101)
	if base.F64_gt(v1125, v1124)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1125)&int64(9223372036854775807))) != 0 {
		v1138 = v1124
		goto L205
	} else {
		goto L206
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = int64(4607182418800017408)
	v221 = m.G0
	v223 = v221 - int32(16)
	m.G0 = v223
	if l10 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v153 != int32(4) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v159 = int32(0)
	if base.B2i32(v156 == v159)|base.B2i32(v158 == v159) != 0 {
		v205 = base.B2i32(v156|v158 == v159)
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v205 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L28:
	;
	goto L27
L29:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v173 != v174 {
		v205 = int32(0)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v176 = int32(1)
	if v173 <= v176 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v179 = v176
	goto L33
L32:
	;
	v179 = v173
	goto L33
L33:
	;
	v180 = int32(8)
	v185 = int32(0)
	goto L34
L34:
	;
	v193 = v185 << (uint(int32(2)) % 32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v156+v180+v193)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v158+v180+v193)))
	v198 = base.B2i32(v195 == v197)
	if v195 != v197 {
		v205 = v198
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v205 = v198
	goto L28
L36:
	;
	v201 = v185 + int32(1)
	if v201 != v179 {
		v185 = v201
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v212 = float64(1)
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = base.F64_div(v212, v150)
	v1101 = base.F64_div(v212, v93)
	goto L23
L39:
	;
	m.G0 = v223 + int32(16)
	if v886 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L40:
	;
	v886 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v228 < int32(2) {
		v886 = l10
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v231 = F_list_copy(m, l10)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = base.F64_div(float64(1), v860)
	v886 = v849
	goto L39
L45:
	;
	if v231 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v849 = int32(0)
	v860 = float64(1)
	goto L44
L47:
	;
	goto L48
L48:
	;
	v239 = int32(0)
	v250 = float64(1)
	v270 = v231
	goto L49
L49:
	;
	v273 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v273
	v275 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v275
	v281 = v239
	v282 = v273
	v283 = v275
	v287 = v275
	v311 = v275
	v312 = v270
	goto L51
L50:
	;
	v849 = v815
	v860 = v826
	goto L44
L51:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v287 < v315 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	if v659 != 0 {
		goto L114
	} else {
		goto L115
	}
L53:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317+v287<<(uint(int32(2))%32))))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+120)))
	if v324 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v625 = v281
	v627 = v283
	v655 = v311
	v656 = v312
	goto L55
L55:
	;
	goto L52
L56:
	;
	v325 = int32(48)
	goto L58
L57:
	;
	v325 = int32(44)
	goto L58
L58:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v321+v325)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	if v324 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v346 = int32(0)
	if v327 == v346 {
		goto L73
	} else {
		goto L74
	}
L60:
	;
	v330 = int32(0)
	if v329 == v330 {
		v343 = v330
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v329 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v333 < int32(2) {
		v343 = v330
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v343 = v337
	goto L59
L65:
	;
	v343 = int32(0)
	goto L59
L66:
	;
	goto L67
L67:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v343 = v342
	goto L59
L68:
	;
	if v620 != 0 {
		v281 = v589
		v282 = v590
		v283 = v591
		v287 = v596 + int32(1)
		v311 = v619
		v312 = v620
		goto L51
	} else {
		goto L111
	}
L69:
	;
	v587 = F_list_delete_nth_cell(m, v312, v287)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L110
	}
L70:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v444 = F_remove_nulling_relids(m, v343, v442, int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L96
	}
L71:
	;
	v437 = F_lappend(m, v281, v321)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L95
	}
L72:
	;
	if v400 == int32(0) {
		goto L71
	} else {
		goto L87
	}
L73:
	;
	v400 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v354 = int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v355 <= v354 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v358 = v354
	goto L78
L77:
	;
	v358 = v355
	goto L78
L78:
	;
	v363 = int32(0)
	v365 = int32(-1)
	goto L80
L79:
	;
	v400 = v392
	goto L72
L80:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v327+int32(8)+v363<<(uint(int32(2))%32))))
	if v373 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223+int32(12)))) = v384
	v392 = int32(1)
	goto L79
L82:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v373)))|base.B2i32(int32(0) <= v365) != 0 {
		v392 = v346
		goto L79
	} else {
		goto L85
	}
L83:
	;
	v384 = v365
	goto L84
L84:
	;
	v386 = v363 + int32(1)
	if v386 != v358 {
		v363 = v386
		v365 = v384
		goto L80
	} else {
		goto L86
	}
L85:
	;
	v384 = base.I32_ctz(v373) | v363<<(uint(int32(5))%32)
	goto L84
L86:
	;
	goto L81
L87:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v405 = v403 << (uint(int32(2)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405+v406)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+120))
	if v409 == int32(0) {
		goto L71
	} else {
		goto L88
	}
L88:
	;
	if v282 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v414+v405)))
	if v416 == int32(0) {
		goto L71
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if base.B2i32(v282 != v403) == int32(0) {
		v439 = v283
		v440 = v282
		goto L70
	} else {
		goto L94
	}
L92:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+21)))
	v421 = v419 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v421))|base.B2i32(int32(1)<<(uint(v421)%32)&int32(_a_F_create_hashjoin_path_0) == int32(0)) != 0 {
		goto L71
	} else {
		goto L93
	}
L93:
	;
	v439 = v408
	v440 = v403
	goto L70
L94:
	;
	v589 = v281
	v590 = v282
	v591 = v283
	v596 = v287
	v619 = v311
	v620 = v312
	goto L68
L95:
	;
	v551 = v437
	v552 = v282
	v553 = v283
	v581 = v311
	goto L69
L96:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	if v446 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v536 = F_palloc0(m, int32(24))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L107
	}
L98:
	;
	v449 = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v450 <= v449 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v455 = v449
	goto L100
L100:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v487+v455<<(uint(int32(2))%32))))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v493 = F_equal(m, v444, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	v589 = v281
	v590 = v440
	v591 = v439
	v596 = v287
	v619 = v311
	v620 = v312
	goto L68
L102:
	;
	if v493 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v498 = v455 + int32(1)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v498 < v499 {
		v455 = v498
		goto L100
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	goto L101
L106:
	;
	goto L97
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v444
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v539+v540<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+4)) = v544
	v546 = F_lappend(m, v446, v536)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v546
	v549 = F_lappend(m, v311, v321)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v551 = v281
	v552 = v440
	v553 = v439
	v581 = v549
	goto L69
L110:
	;
	v589 = v551
	v590 = v552
	v591 = v553
	v596 = v287 - int32(1)
	v619 = v581
	v620 = v587
	goto L68
L111:
	;
	v625 = v589
	v627 = v591
	v655 = v619
	v656 = v620
	goto L55
L112:
	;
	if v656 != 0 {
		v239 = v815
		v250 = v826
		v270 = v656
		goto L49
	} else {
		goto L156
	}
L113:
	;
	v681 = v250
	goto L121
L114:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if int32(2) <= v660 {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v663 = F_list_concat(m, v625, v655)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	goto L116
L118:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	F_list_free_deep(m, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_list_free(m, v655)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v815 = v663
	v826 = v250
	goto L112
L121:
	;
	v706 = F_estimate_multivariate_ndistinct(m, l0, v627, v223+int32(8), v223)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	v717 = v625
	v718 = int32(0)
	goto L128
L123:
	;
	v708 = *(*float64)(unsafe.Add(mBase, uint32(v223)))
	if base.F64_lt(v681, v708) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v710 = v708
	goto L126
L125:
	;
	v710 = v681
	goto L126
L126:
	;
	if v706 != 0 {
		v681 = v710
		goto L121
	} else {
		goto L127
	}
L127:
	;
	goto L122
L128:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v718 < v747 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v815 = v717
	v826 = v681
	goto L112
L130:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v659)+12))
	v753 = v749 + v718<<(uint(int32(2))%32)
	goto L132
L131:
	;
	v753 = int32(0)
	goto L132
L132:
	;
	if v655 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v815 = v625
	v826 = v681
	goto L112
L134:
	;
	goto L135
L135:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	if base.B2i32(v753 == int32(0))|base.B2i32(v758 <= v718) != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	goto L129
L137:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v655)+12))
	if v761 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	v766 = int32(0)
	if v764 == v766 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v804 != 0 {
		goto L152
	} else {
		goto L153
	}
L140:
	;
	v804 = int32(0)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v764)+4))
	if v772 <= int32(0) {
		v798 = v766
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v804 = v798
	goto L139
L144:
	;
	v775 = int32(0)
	if v775 < v772 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v778 = v772
	goto L147
L146:
	;
	v778 = v775
	goto L147
L147:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v764)+12))
	v781 = int32(0)
	goto L148
L148:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v779+v781<<(uint(int32(2))%32))))
	v790 = base.B2i32(v789 == v765)
	if v789 == v765 {
		v798 = v790
		goto L143
	} else {
		goto L150
	}
L149:
	;
	v798 = v790
	goto L143
L150:
	;
	v792 = v781 + int32(1)
	if v792 != v778 {
		v781 = v792
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v761+v718<<(uint(int32(2))%32))))
	v809 = F_lappend(m, v717, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	v811 = v717
	goto L154
L154:
	;
	v717 = v811
	v718 = v718 + int32(1)
	goto L128
L155:
	;
	v811 = v809
	goto L154
L156:
	;
	goto L50
L157:
	;
	v1101 = float64(1)
	goto L23
L158:
	;
	goto L159
L159:
	;
	v926 = float64(1)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	if v927 <= int32(0) {
		v1101 = v926
		goto L23
	} else {
		goto L160
	}
L160:
	;
	v936 = int32(0)
	v942 = v926
	goto L161
L161:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v886)+12))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v965+v936<<(uint(int32(2))%32))))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+48))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v971)+8))
	v973 = int32(0)
	if v970 == v973 {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v1101 = v1085
	goto L23
L163:
	;
	v1077 = *(*float64)(unsafe.Add(mBase, uint32(v1072+v969)))
	v1078 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	if base.F64_lt(v1075, v1078) != 0 {
		goto L195
	} else {
		goto L196
	}
L164:
	;
	if v1026 != 0 {
		goto L178
	} else {
		goto L179
	}
L165:
	;
	v1026 = int32(1)
	goto L164
L166:
	;
	goto L167
L167:
	;
	if v972 == int32(0) {
		v1019 = v973
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v1026 = v1019
	goto L164
L169:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v970)+4))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v972)+4))
	if v983 < v982 {
		v1019 = v973
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v985 = int32(1)
	if v982 <= v985 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v988 = v985
	goto L173
L172:
	;
	v988 = v982
	goto L173
L173:
	;
	v989 = int32(8)
	v994 = int32(0)
	goto L174
L174:
	;
	v1001 = v994 << (uint(int32(2)) % 32)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v970+v989+v1001)))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v972+v989+v1001)))
	v1008 = v1003 & (v1005 ^ int32(-1))
	v1010 = base.B2i32(v1008 == int32(0))
	if v1008 != 0 {
		v1019 = v1010
		goto L168
	} else {
		goto L176
	}
L175:
	;
	v1019 = v1010
	goto L168
L176:
	;
	v1012 = v994 + int32(1)
	if v1012 != v988 {
		v994 = v1012
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v1027 = *(*float64)(unsafe.Add(mBase, uint32(v969)+136))
	if base.F64_lt(v1027, float64(0)) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	v1052 = *(*float64)(unsafe.Add(mBase, uint32(v969)+128))
	if base.F64_lt(v1052, float64(0)) == int32(0) {
		goto L188
	} else {
		goto L189
	}
L181:
	;
	v1072 = int32(152)
	v1075 = v1027
	goto L163
L182:
	;
	goto L183
L183:
	;
	v1035 = int32(0)
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+28))
	if v1037 == v1035 {
		v1045 = v1035
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1046 = int32(152)
	F_estimate_hash_bucket_stats(m, l0, v1045, v150, v969+v1046, v969+int32(136))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	if v1040 < int32(2) {
		v1045 = v1035
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+12))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	v1045 = v1044
	goto L184
L187:
	;
	v1051 = *(*float64)(unsafe.Add(mBase, uint32(v969)+136))
	v1072 = v1046
	v1075 = v1051
	goto L163
L188:
	;
	v1072 = int32(144)
	v1075 = v1052
	goto L163
L189:
	;
	goto L190
L190:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v969)+4))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+28))
	if v1062 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+12))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	v1066 = v1064
	goto L193
L192:
	;
	v1066 = int32(0)
	goto L193
L193:
	;
	F_estimate_hash_bucket_stats(m, l0, v1066, v150, v969+int32(144), v969+int32(128))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v1071 = *(*float64)(unsafe.Add(mBase, uint32(v969)+128))
	v1072 = int32(144)
	v1075 = v1071
	goto L163
L195:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v86))) = v1075
	goto L197
L196:
	;
	goto L197
L197:
	;
	if base.F64_gt(v942, v1077) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1082 = v1077
	goto L200
L199:
	;
	v1082 = v942
	goto L200
L200:
	;
	if base.F64_gt(v1077, float64(0)) != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1085 = v1082
	goto L203
L202:
	;
	v1085 = v942
	goto L203
L203:
	;
	v1087 = v936 + int32(1)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	if v1087 < v1088 {
		v936 = v1087
		v942 = v1085
		goto L161
	} else {
		goto L204
	}
L204:
	;
	goto L162
L205:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+32))
	v1144 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[1]))
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_create_hashjoin_path[2]))
	v1150 = base.F64_mul(base.F64_mul(v1144, base.F64_convert_i32_s(v1146)), float64(1024))
	v1151 = float64(4.294967295e+09)
	if base.F64_lt(v1150, v1151) != 0 {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v1134 = float64(1)
	if base.F64_le(v1125, v1134) != 0 {
		v1138 = v1134
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v1138 = base.F64_nearest(v1125)
	goto L205
L208:
	;
	v1157 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[3]))
	v1158 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v1158
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v1158
	v1163 = float64(0)
	if l10 == int32(0) {
		v1233 = v1163
		v1254 = v1163
		goto L212
	} else {
		goto L213
	}
L209:
	;
	v1154 = v1150
	goto L211
L210:
	;
	v1154 = v1151
	goto L211
L211:
	;
	goto L208
L212:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	v1256 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v1256
	if v1255 == int32(0) {
		v1341 = v12
		v1352 = float64(0)
		goto L219
	} else {
		goto L220
	}
L213:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1168 <= int32(0) {
		v1233 = v1163
		v1254 = float64(0)
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v1174 = int32(0)
	goto L215
L215:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1205+v1174<<(uint(int32(2))%32))))
	v1212 = F_cost_qual_eval_walker(m, v1209, v86+int32(8))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L217
	}
L216:
	;
	v1218 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v1219 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v1233 = v1218
	v1254 = v1219
	goto L212
L217:
	;
	v1215 = v1174 + int32(1)
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1215 < v1216 {
		v1174 = v1215
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v1353&int32(-2) != int32(4) {
		goto L228
	} else {
		goto L229
	}
L220:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+4))
	if v1265 <= int32(0) {
		v1341 = v12
		v1352 = float64(0)
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1272 = int32(0)
	goto L222
L222:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+12))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1303+v1272<<(uint(int32(2))%32))))
	v1310 = F_cost_qual_eval_walker(m, v1307, v86+int32(8))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L224
	}
L223:
	;
	v1316 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v1317 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v1341 = v1316
	v1352 = v1317
	goto L219
L224:
	;
	v1313 = v1272 + int32(1)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+4))
	if v1313 < v1314 {
		v1272 = v1313
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v1606 = *(*float64)(unsafe.Add(mBase, _c_F_create_hashjoin_path[4]))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v1608 = *(*float64)(unsafe.Add(mBase, uint32(v1607)+24))
	if base.F64_lt(base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v1154)), base.F64_mul(v1138, base.F64_convert_i32_u((v1140+int32(7))&int32(-8)+int32(24)))) != 0 {
		goto L258
	} else {
		goto L259
	}
L227:
	;
	v1416 = float64(1e+100)
	v1417 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	v1418 = base.F64_mul(v95, v1417)
	if base.F64_gt(v1418, v1416)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1418)&int64(9223372036854775807))) != 0 {
		v1431 = v1416
		goto L242
	} else {
		goto L243
	}
L228:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v1358 != int32(1) {
		goto L227
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1361 = float64(1e+100)
	v1362 = *(*float64)(unsafe.Add(mBase, uint32(l4)+16))
	v1364 = base.F64_nearest(base.F64_mul(v97, v1362))
	v1366 = base.F64_sub(v97, v1364)
	v1369 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	v1372 = *(*float64)(unsafe.Add(mBase, uint32(l4)+24))
	v1376 = base.F64_mul(base.F64_mul(v95, v1369), base.F64_div(float64(2), base.F64_add(v1372, float64(1))))
	if base.F64_gt(v1376, v1361) != 0 {
		v1389 = v1361
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	v1394 = base.F64_div(v95, v150)
	if base.F64_gt(v1394, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1394)&int64(9223372036854775807))) != 0 {
		v1407 = v1361
		goto L236
	} else {
		goto L237
	}
L233:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1376)&int64(9223372036854775807)) {
		v1389 = float64(1e+100)
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1385 = float64(1)
	if base.F64_le(v1376, v1385) != 0 {
		v1389 = v1385
		goto L232
	} else {
		goto L235
	}
L235:
	;
	v1389 = base.F64_nearest(v1376)
	goto L232
L236:
	;
	if v1353 == int32(5) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v1403 = float64(1)
	if base.F64_le(v1394, v1403) != 0 {
		v1407 = v1403
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1407 = base.F64_nearest(v1394)
	goto L236
L239:
	;
	v1410 = v1366
	goto L241
L240:
	;
	v1410 = v1364
	goto L241
L241:
	;
	v1582 = v1410
	v1604 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1233, v1366), v1407), float64(0.05)), base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v1233, v1364), v1389), float64(0.5)), v90))
	goto L226
L242:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
	v1433 = *(*float64)(unsafe.Add(mBase, uint32(v1432)+32))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v1435 = *(*float64)(unsafe.Add(mBase, uint32(v1434)+32))
	v1437 = v86 + int32(8)
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+8))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+8))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+8))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+8))
	v1442 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1437)+48)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+16)) = v1441
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+12)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+8)) = v1441
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+4)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v1437))) = int32(322)
	*(*int64)(unsafe.Add(mBase, uint32(v1437)+20)) = v1442
	*(*int64)(unsafe.Add(mBase, uint32(v1437)+28)) = v1442
	*(*int64)(unsafe.Add(mBase, uint32(v1437)+36)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+43)) = int32(0)
	goto L245
L243:
	;
	v1427 = float64(1)
	if base.F64_le(v1418, v1427) != 0 {
		v1431 = v1427
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v1431 = base.F64_nearest(v1418)
	goto L242
L245:
	;
	if l10 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1553 = float64(1e+100)
	v1555 = base.F64_mul(v1433, base.F64_mul(v1435, v1528))
	if base.F64_gt(v1555, v1553)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v1555)&int64(9223372036854775807))) != 0 {
		v1568 = v1553
		goto L255
	} else {
		goto L256
	}
L247:
	;
	v1528 = float64(1)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1462 = float64(1)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1463 <= int32(0) {
		v1528 = v1462
		goto L246
	} else {
		goto L250
	}
L250:
	;
	v1470 = int32(0)
	v1478 = v1462
	goto L251
L251:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1501+v1470<<(uint(int32(2))%32))))
	v1506 = int32(0)
	v1510 = F_clause_selectivity(m, l0, v1505, v1506, v1506, v86+int32(8))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	v1528 = v1512
	goto L246
L253:
	;
	v1512 = base.F64_mul(v1478, v1510)
	v1514 = v1470 + int32(1)
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v1514 < v1515 {
		v1470 = v1514
		v1478 = v1512
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1582 = v1568
	v1604 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_mul(v97, v1233), v1431), float64(0.5)), v90)
	goto L226
L256:
	;
	v1564 = float64(1)
	if base.F64_le(v1555, v1564) != 0 {
		v1568 = v1564
		goto L255
	} else {
		goto L257
	}
L257:
	;
	v1568 = base.F64_nearest(v1555)
	goto L255
L258:
	;
	v1620 = base.F64_add(v91, v1157)
	goto L260
L259:
	;
	v1620 = v91
	goto L260
L260:
	;
	v1624 = *(*float64)(unsafe.Add(mBase, uint32(v1607)+16))
	v1625 = base.F64_add(base.F64_add(base.F64_add(v1620, v1254), base.F64_sub(v1352, v1254)), v1624)
	*(*float64)(unsafe.Add(mBase, uint32(v41)+48)) = v1625
	v1627 = *(*float64)(unsafe.Add(mBase, uint32(v41)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+56)) = base.F64_add(v1625, base.F64_add(base.F64_mul(v1608, v1627), base.F64_add(base.F64_mul(base.F64_add(v1606, base.F64_sub(v1341, v1233)), v1582), v1604)))
	m.G0 = v86 - int32(-64)
	m.G0 = v37 + int32(16)
	return v41
}
func F_create_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 float64
	_ = v15
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v114 int32
	_ = v114
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v146 int64
	_ = v146
	var v151 float64
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 float64
	_ = v207
	var v208 float64
	_ = v208
	var v226 float64
	_ = v226
	var v243 float64
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 float64
	_ = v305
	var v306 float64
	_ = v306
	var v335 float64
	_ = v335
	var v341 float64
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 float64
	_ = v367
	var v368 int32
	_ = v368
	var v369 float64
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v395 float64
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v414 float64
	_ = v414
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 float64
	_ = v443
	var v444 int32
	_ = v444
	var v445 float64
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v464 float64
	_ = v464
	var v486 float64
	_ = v486
	var v488 float64
	_ = v488
	var v497 float64
	_ = v497
	var v501 float64
	_ = v501
	var v503 float64
	_ = v503
	var v504 float64
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v566 float64
	_ = v566
	var v573 float64
	_ = v573
	var v574 int32
	_ = v574
	var v577 float64
	_ = v577
	var v578 float64
	_ = v578
	var v579 int64
	_ = v579
	var v580 int32
	_ = v580
	var v582 float64
	_ = v582
	var v585 float64
	_ = v585
	var v586 int64
	_ = v586
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v723 int32
	_ = v723
	var v735 float64
	_ = v735
	var v751 int64
	_ = v751
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v762 int64
	_ = v762
	var v768 int32
	_ = v768
	var v769 float64
	_ = v769
	var v771 float64
	_ = v771
	var v779 float64
	_ = v779
	var v780 float64
	_ = v780
	var v782 float64
	_ = v782
	v15 = float64(0)
	v33 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = l7
	v41 = F_palloc0(m, int32(120))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(1554778161453)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v53 = F_get_joinrel_parampathinfo(m, l0, l1, l5, l6, v50, l9, v37+int32(12))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v53
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v58 != int32(1) {
		v65 = v33
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v67 = v65 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+21)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v41)+64)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v69
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+84)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v41)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+76)) = uint8(v73)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v41)+104)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v41)+100)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v41)+96)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v77
	v84 = m.G0
	v86 = v84 + int32(-64)
	m.G0 = v86
	v89 = *(*float64)(unsafe.Add(mBase, uint32(l6)+32))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l3)+72))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(l3)+64))
	v92 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l3)+48))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v100 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+21)))
	if v61 != int32(1) {
		v65 = v33
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+21)))
	v65 = v64
	goto L4
L7:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v106)))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if int32(0) < v109 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v106 = v100 + int32(8)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v106 = v103 + int32(16)
	goto L7
L11:
	;
	v112 = base.F64_convert_i32_u(v109)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_mergejoin_path[0])))
	if v114 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v146
	v151 = float64(0)
	if l10 == int32(0) {
		v226 = v151
		v243 = v151
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v120 = base.F64_add(base.F64_mul(v112, float64(-0.3)), float64(1))
	if base.F64_gt(v120, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v126 = v112
	goto L16
L16:
	;
	v128 = float64(1e+100)
	v129 = base.F64_div(v107, v126)
	if base.F64_gt(v129, v128)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v129)&int64(9223372036854775807))) != 0 {
		v142 = v128
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v124 = v120
	goto L19
L18:
	;
	v124 = math.Float64frombits(uint64(0x8000000000000000))
	goto L19
L19:
	;
	v126 = base.F64_add(v124, v112)
	goto L16
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v41)+32)) = v142
	goto L13
L21:
	;
	v138 = float64(1)
	if base.F64_le(v129, v138) != 0 {
		v142 = v138
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v142 = base.F64_nearest(v129)
	goto L20
L23:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	v245 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v245
	if v244 == int32(0) {
		v335 = v15
		v341 = float64(0)
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v156 <= int32(0) {
		v226 = v151
		v243 = float64(0)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v161 = int32(0)
	goto L26
L26:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v161<<(uint(int32(2))%32))))
	v201 = F_cost_qual_eval_walker(m, v198, v84+int32(-56))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v207 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v226 = v207
	v243 = v208
	goto L23
L28:
	;
	v204 = v161 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v204 < v205 {
		v161 = v204
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v342&int32(-2) != int32(4) {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v254 <= int32(0) {
		v335 = v15
		v341 = float64(0)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v259 = int32(0)
	goto L33
L33:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v259<<(uint(int32(2))%32))))
	v299 = F_cost_qual_eval_walker(m, v296, v84+int32(-56))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v305 = *(*float64)(unsafe.Add(mBase, uint32(v86)+24))
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v86)+16))
	v335 = v305
	v341 = v306
	goto L30
L35:
	;
	v302 = v259 + int32(1)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v302 < v303 {
		v259 = v302
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+112)) = uint8(v364)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v41)+84))
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v366)+32))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
	v369 = *(*float64)(unsafe.Add(mBase, uint32(v368)+32))
	v371 = v84 + int32(-56)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+8))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	v376 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v371)+48)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v371)+12)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v371)+8)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(322)
	*(*int64)(unsafe.Add(mBase, uint32(v371)+20)) = v376
	*(*int64)(unsafe.Add(mBase, uint32(v371)+28)) = v376
	*(*int64)(unsafe.Add(mBase, uint32(v371)+36)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v371)+43)) = int32(0)
	goto L50
L38:
	;
	v364 = int32(0)
	goto L37
L39:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)))
	if v347 != int32(1) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	if v350 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v352 = v351
	goto L45
L44:
	;
	v352 = int32(0)
	goto L45
L45:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v41)+96))
	if v354 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v357 = v355
	goto L48
L47:
	;
	v357 = int32(0)
	goto L48
L48:
	;
	if v357 == v352 {
		v364 = int32(1)
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	if l10 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v486 = float64(1e+100)
	v488 = base.F64_mul(v367, base.F64_mul(v369, v464))
	if base.F64_gt(v488, v486)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v488)&int64(9223372036854775807))) != 0 {
		v501 = v486
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v464 = float64(1)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v395 = float64(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v396 <= int32(0) {
		v464 = v395
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v401 = int32(0)
	v414 = v395
	goto L56
L56:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l10)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434+v401<<(uint(int32(2))%32))))
	v439 = int32(0)
	v443 = F_clause_selectivity(m, l0, v438, v439, v439, v84+int32(-56))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v464 = v445
	goto L51
L58:
	;
	v445 = base.F64_mul(v414, v443)
	v447 = v401 + int32(1)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l10)+4))
	if v447 < v448 {
		v401 = v447
		v414 = v445
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if base.F64_le(v89, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v497 = float64(1)
	if base.F64_le(v488, v497) != 0 {
		v501 = v497
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v501 = base.F64_nearest(v488)
	goto L60
L63:
	;
	v503 = float64(1)
	goto L65
L64:
	;
	v503 = v89
	goto L65
L65:
	;
	v504 = float64(0)
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+112)))
	if v505 != 0 {
		v573 = v504
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v574 = int32(0)
	v577 = base.F64_add(base.F64_div(v573, v92), float64(1))
	v578 = base.F64_mul(v94, v577)
	v579 = int64(64)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+112)))
	if v580 != 0 {
		v723 = v574
		v735 = v578
		v751 = v579
		goto L84
	} else {
		goto L85
	}
L67:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v506 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v566 = base.F64_sub(v501, v503)
	if base.F64_lt(v566, float64(0)) == int32(0) {
		v573 = v566
		goto L66
	} else {
		goto L83
	}
L69:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+20))
	if v508 != int32(4) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v514 = int32(0)
	if base.B2i32(v511 == v514)|base.B2i32(v513 == v514) != 0 {
		v560 = base.B2i32(v511|v513 == v514)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v560 != 0 {
		v573 = v504
		goto L66
	} else {
		goto L82
	}
L72:
	;
	goto L71
L73:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	if v528 != v529 {
		v560 = int32(0)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v531 = int32(1)
	if v528 <= v531 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v534 = v531
	goto L77
L76:
	;
	v534 = v528
	goto L77
L77:
	;
	v535 = int32(8)
	v540 = int32(0)
	goto L78
L78:
	;
	v548 = v540 << (uint(int32(2)) % 32)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v511+v535+v548)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v513+v535+v548)))
	v553 = base.B2i32(v550 == v552)
	if v550 != v552 {
		v560 = v553
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v560 = v553
	goto L72
L80:
	;
	v556 = v540 + int32(1)
	if v556 != v534 {
		v540 = v556
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L68
L83:
	;
	v573 = float64(0)
	goto L66
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+113)) = uint8(v723)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v756
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v760 != 0 {
		goto L110
	} else {
		goto L111
	}
L85:
	;
	v582 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[1]))
	v585 = base.F64_add(base.F64_mul(base.F64_mul(v92, v582), v577), v94)
	v586 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	if v586&int64(128) != int64(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v723 = int32(1)
	v735 = v585
	v751 = int64(128)
	goto L84
L87:
	;
	v615 = int32(0)
	v621 = l6
	goto L96
L88:
	;
	if base.B2i32(v586&int64(64) == int64(0))|base.F64_gt(v578, v585) != 0 {
		goto L86
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if l12 != 0 {
		v723 = v574
		v735 = v578
		v751 = v579
		goto L84
	} else {
		goto L94
	}
L91:
	;
	if l12 == int32(0) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_create_mergejoin_path[2]))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v604)+32))
	if base.F64_lt(base.F64_convert_i32_u(v600<<(uint(int32(10))%32)), base.F64_mul(v503, base.F64_convert_i32_u((v605+int32(7))&int32(-8)+int32(24)))) != 0 {
		goto L86
	} else {
		goto L93
	}
L93:
	;
	v723 = v574
	v735 = v578
	v751 = v579
	goto L84
L94:
	;
	goto L87
L95:
	;
	if v681&int32(1) != 0 {
		v723 = v574
		v735 = v578
		v751 = v579
		goto L84
	} else {
		goto L109
	}
L96:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v621)+4))
	switch v650 - int32(335) {
	case 0:
		goto L101
	default:
		v681 = v615
		goto L95
	case 3:
		goto L100
	case 4:
		goto L99
	case 10, 11:
		goto L103
	case 24:
		goto L102
	case 29, 31:
		goto L98
	}
L97:
	;
	v681 = int32(1)
	goto L95
L98:
	;
	goto L97
L99:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v621)+72))
	if v672 == int32(0) {
		v681 = v615
		goto L95
	} else {
		goto L107
	}
L100:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v621)+72))
	if v664 == int32(0) {
		v681 = v615
		goto L95
	} else {
		goto L105
	}
L101:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	if v660 != int32(303) {
		v681 = v615
		goto L95
	} else {
		goto L104
	}
L102:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+72)))
	v681 = int32(base.Ui32(v655&int32(2)) >> (uint(int32(1)) % 32))
	goto L95
L103:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v621)+72))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+113)))
	v681 = v654
	goto L95
L104:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v621)+72))
	v621 = v663
	goto L96
L105:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v664)+4))
	if v667 != int32(1) {
		v681 = v615
		goto L95
	} else {
		goto L106
	}
L106:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v664)+12))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	v621 = v671
	goto L96
L107:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v675 != int32(1) {
		v681 = v615
		goto L95
	} else {
		goto L108
	}
L108:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)))
	v621 = v679
	goto L96
L109:
	;
	goto L86
L110:
	;
	v761 = v751
	goto L112
L111:
	;
	v761 = v751 | int64(262144)
	goto L112
L112:
	;
	v762 = *(*int64)(unsafe.Add(mBase, uint32(l4)+40))
	if v761&v762 != v761 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v756 + int32(1)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v769 = *(*float64)(unsafe.Add(mBase, uint32(v768)+24))
	v771 = *(*float64)(unsafe.Add(mBase, _c_F_create_mergejoin_path[3]))
	v779 = *(*float64)(unsafe.Add(mBase, uint32(v768)+16))
	v780 = base.F64_add(base.F64_add(base.F64_sub(v341, v243), base.F64_add(base.F64_mul(v226, base.F64_add(base.F64_mul(v90, v577), v91)), base.F64_add(v96, v243))), v779)
	*(*float64)(unsafe.Add(mBase, uint32(v41)+48)) = v780
	v782 = *(*float64)(unsafe.Add(mBase, uint32(v41)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v41)+56)) = base.F64_add(v780, base.F64_add(base.F64_mul(v769, v782), base.F64_add(base.F64_mul(base.F64_add(v771, base.F64_sub(v335, v226)), v501), base.F64_add(base.F64_mul(v226, base.F64_add(base.F64_mul(base.F64_sub(v92, v90), v577), base.F64_sub(v93, v91))), base.F64_add(v95, v735)))))
	m.G0 = v86 - int32(-64)
	m.G0 = v37 + int32(16)
	return v41
}
func F_create_unique_paths(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v296 int32
	_ = v296
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
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
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
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
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int64
	_ = v769
	var v775 int32
	_ = v775
	var v776 float64
	_ = v776
	var v777 int32
	_ = v777
	var v779 float64
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v807 int32
	_ = v807
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 float64
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 float64
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1019 int32
	_ = v1019
	var v1029 int32
	_ = v1029
	var v1054 int32
	_ = v1054
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(48)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	if v28 != 0 {
		v1054 = v28
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v26 + int32(48)
	return v1054
L2:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+45)))
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v34))|base.B2i32(int32(1)<<(uint(v34)%32)&int32(44) == int32(0)) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+46)))
	if v30 == int32(1) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v1054 = int32(0)
	goto L1
L6:
	;
	v47 = F_GetMemoryChunkContext(m, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+192))
	if v45 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v1054 = int32(0)
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v51 = int32(_a_F_create_unique_paths_0)
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_create_unique_paths[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_create_unique_paths[0])) = v47
	v56 = F_palloc0(m, int32(304))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(270)
	base.MemoryCopy(m, v56, l1, int32(304))
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+60)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v56)+52)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v56)+44)) = v62
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v75 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v68))|base.B2i32(int32(1)<<(uint(v68)%32)&int32(44) == v75) == v75 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_create_unique_paths[0])) = v52
	v1054 = v1029
	goto L1
L13:
	;
	F_create_final_unique_paths(m, l0, l1, v729, v734, l2, v56)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L184
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+192))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+40))
	v83 = F_copy_pathtarget(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v96 = F_make_tlist_from_pathtarget(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v87 = F_adjust_appendrel_attrs_multilevel(m, l0, v85, l1, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v83
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+200))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+196))
	v729 = v93
	v734 = v92
	goto L13
L19:
	;
	if v96 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v101 = v98 + int32(1)
	goto L22
L21:
	;
	v101 = int32(1)
	goto L22
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v109 = v101
	v111 = v4
	v113 = v96
	v116 = v4
	v120 = v4
	v124 = v4
	goto L25
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L9
	} else {
		goto L181
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L9
	} else {
		goto L178
	}
L25:
	;
	v127 = int32(0)
	if v103 == v127 {
		v137 = v127
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L9
	} else {
		goto L175
	}
L27:
	;
	v138 = int32(0)
	if v102 == v138 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v131 <= v111 {
		v137 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v137 = v133 + v111<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v150+v111<<(uint(int32(2))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v172 = F_tlist_member(m, v171, v113)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L46
	}
L31:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+45)))
	if v157 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v141 = int32(0)
	v153 = v141
	v154 = v96
	v155 = v141
	goto L31
L33:
	;
	goto L34
L34:
	;
	v143 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if base.B2i32(v137 == v143)|base.B2i32(v145 <= v111) == v143 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	if v150 != 0 {
		goto L30
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v153 = v116
	v154 = v113
	v155 = v124
	goto L31
L38:
	;
	goto L37
L39:
	;
	v158 = v155
	goto L41
L40:
	;
	v158 = int32(0)
	goto L41
L41:
	;
	if v158|v153 == int32(0) {
		v1029 = v138
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v162 = F_make_pathtarget_from_tlist(m, v154)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v164 = F_set_pathtarget_cost_width(m, l0, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+40)) = v164
	v729 = v158
	v734 = v153
	goto L13
L45:
	;
	v190 = int64(0)
	v192 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(v170), v190, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L9
	} else {
		goto L53
	}
L46:
	;
	if v172 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v183 = v172
	v184 = v109
	v185 = v113
	goto L45
L48:
	;
	goto L49
L49:
	;
	v177 = int32(0)
	v179 = F_makeTargetEntry(m, v171, base.I32_extend16_s(v109), v177, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v181 = F_lappend(m, v113, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v183 = v179
	v184 = v109 + int32(1)
	v185 = v181
	goto L45
L52:
	;
	F_ReleaseCatCacheList(m, v192)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L75
	}
L53:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)+56))
	if int32(0) < v194 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v210 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v296 = int32(0)
	goto L52
L57:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v192-int32(-64)+v210<<(uint(int32(2))%32))))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+72))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+22)))
	v228 = v226 + v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+24))
	if v229 <= int32(2741) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	goto L56
L59:
	;
	v265 = v210 + int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v192)+56))
	if v265 < v266 {
		v210 = v265
		goto L57
	} else {
		goto L74
	}
L60:
	;
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+16)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v254 = F_IndexAmTranslateStrategy(m, v252, v251, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L9
	} else {
		goto L70
	}
L61:
	;
	v245 = F_GetIndexAmRoutineByAmId(m, v229, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L68
	}
L62:
	;
	switch v229 - int32(403) {
	case 0:
		v251 = v229
		goto L60
	case 1:
		goto L61
	case 2:
		goto L59
	default:
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.B2i32(v229 == int32(2742))|base.B2i32(v229 == int32(3580))|base.B2i32(v229 == int32(4000)) != 0 {
		goto L59
	} else {
		goto L67
	}
L65:
	;
	if v229 != int32(783) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	goto L61
L68:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+10)))
	if v247 != int32(1) {
		goto L59
	} else {
		goto L69
	}
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v228)+24))
	v251 = v250
	goto L60
L70:
	;
	if v254 != int32(3) {
		goto L59
	} else {
		goto L71
	}
L71:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v261 = F_get_opfamily_member_for_cmptype(m, v258, v259, v259, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	if v261 != 0 {
		v296 = v261
		goto L52
	} else {
		goto L73
	}
L73:
	;
	goto L59
L74:
	;
	goto L58
L75:
	;
	if v296 != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L26
L77:
	;
	v109 = v184
	v111 = v111 + int32(1)
	v113 = v185
	v116 = v672
	v120 = v674
	v124 = v675
	goto L25
L78:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+46)))
	if v503 != int32(1) {
		v672 = v116
		v674 = v500
		v675 = v501
		goto L77
	} else {
		goto L134
	}
L79:
	;
	v318 = F_get_equality_op_for_ordering_op(m, v296, int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+45)))
	if v495 == int32(1) {
		goto L24
	} else {
		goto L133
	}
L82:
	;
	if v318 == int32(0) {
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v323 = F_palloc0(m, int32(20))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = int32(106)
	v327 = int32(0)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	if v336 == v327 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323)+18)) = uint8(v469)
	*(*uint16)(unsafe.Add(mBase, uint32(v323)+16)) = uint16(v469)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v460
	v476 = F_lappend(m, v120, v323)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L9
	} else {
		goto L121
	}
L86:
	;
	if v185 == int32(0) {
		v456 = int32(1)
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v460 = v336
	goto L88
L88:
	;
	goto L85
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+16)) = v456
	v460 = v456
	goto L88
L90:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v343 <= int32(0) {
		v456 = int32(1)
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v346 = int32(0)
	if v346 < v343 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v349 = v343
	goto L94
L93:
	;
	v349 = v346
	goto L94
L94:
	;
	v351 = v349 & int32(3)
	v352 = int32(0)
	if int32(4) <= v343 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v456 = v434 + int32(1)
	goto L89
L96:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v361 = v352
	v362 = int32(0)
	v363 = v327
	goto L99
L97:
	;
	v398 = v352
	v400 = v327
	goto L98
L98:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v409 = int32(0)
	v411 = v398
	v413 = v400
	goto L115
L99:
	;
	v372 = v357 + v363<<(uint(int32(2))%32)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+16))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+16))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+16))
	if base.Ui32(v361) < base.Ui32(v380) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v351 == int32(0) {
		v434 = v388
		goto L95
	} else {
		goto L114
	}
L101:
	;
	v382 = v380
	goto L103
L102:
	;
	v382 = v361
	goto L103
L103:
	;
	if base.Ui32(v382) < base.Ui32(v378) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v384 = v378
	goto L106
L105:
	;
	v384 = v382
	goto L106
L106:
	;
	if base.Ui32(v384) < base.Ui32(v376) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v386 = v376
	goto L109
L108:
	;
	v386 = v384
	goto L109
L109:
	;
	if base.Ui32(v386) < base.Ui32(v374) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v388 = v374
	goto L112
L111:
	;
	v388 = v386
	goto L112
L112:
	;
	v389 = int32(4)
	v390 = v363 + v389
	v392 = v362 + v389
	if v392 != v349&int32(2147483644) {
		v361 = v388
		v362 = v392
		v363 = v390
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	v398 = v388
	v400 = v390
	goto L98
L115:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v407+v413<<(uint(int32(2))%32))))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+16))
	if base.Ui32(v411) < base.Ui32(v424) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v434 = v426
	goto L95
L117:
	;
	v426 = v424
	goto L119
L118:
	;
	v426 = v411
	goto L119
L119:
	;
	v427 = int32(1)
	v430 = v409 + v427
	if v430 != v351 {
		v409 = v430
		v411 = v426
		v413 = v413 + v427
		goto L115
	} else {
		goto L120
	}
L120:
	;
	goto L116
L121:
	;
	v478 = F_make_pathkeys_for_sortclauses(m, l0, v476, v185)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	if v478 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	v482 = v480
	goto L125
L124:
	;
	v482 = int32(0)
	goto L125
L125:
	;
	if v476 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v485 = v483
	goto L128
L127:
	;
	v485 = int32(0)
	goto L128
L128:
	;
	if v482 == v485 {
		v500 = v476
		v501 = v478
		goto L78
	} else {
		goto L129
	}
L129:
	;
	v487 = F_list_delete_last(m, v476)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	if v172 != 0 {
		v672 = v116
		v674 = v487
		v675 = v478
		goto L77
	} else {
		goto L131
	}
L131:
	;
	v491 = F_list_delete_last(m, v185)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L9
	} else {
		goto L132
	}
L132:
	;
	v109 = v184 - int32(1)
	v111 = v111 + int32(1)
	v113 = v491
	v120 = v487
	v124 = v478
	goto L25
L133:
	;
	v500 = v120
	v501 = v124
	goto L78
L134:
	;
	v508 = F_get_compatible_hash_operators(m, v170, v26+int32(44))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	if v508 == int32(0) {
		goto L23
	} else {
		goto L136
	}
L136:
	;
	v513 = F_palloc0(m, int32(20))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = int32(106)
	v517 = int32(0)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	if v526 == v517 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v650
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v661 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v513)+18)) = uint8(v661)
	v663 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v513)+16)) = uint16(v663)
	*(*int32)(unsafe.Add(mBase, uint32(v513)+12)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v513)+8)) = v660
	v667 = F_lappend(m, v116, v513)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L9
	} else {
		goto L174
	}
L139:
	;
	if v185 == int32(0) {
		v646 = int32(1)
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v650 = v526
	goto L141
L141:
	;
	goto L138
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+16)) = v646
	v650 = v646
	goto L141
L143:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v533 <= int32(0) {
		v646 = int32(1)
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v536 = int32(0)
	if v536 < v533 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v539 = v533
	goto L147
L146:
	;
	v539 = v536
	goto L147
L147:
	;
	v541 = v539 & int32(3)
	v542 = int32(0)
	if int32(4) <= v533 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v646 = v624 + int32(1)
	goto L142
L149:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v551 = v542
	v552 = int32(0)
	v553 = v517
	goto L152
L150:
	;
	v588 = v542
	v590 = v517
	goto L151
L151:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v599 = int32(0)
	v601 = v588
	v603 = v590
	goto L168
L152:
	;
	v562 = v547 + v553<<(uint(int32(2))%32)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+16))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+16))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v569)+16))
	if base.Ui32(v551) < base.Ui32(v570) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v541 == int32(0) {
		v624 = v578
		goto L148
	} else {
		goto L167
	}
L154:
	;
	v572 = v570
	goto L156
L155:
	;
	v572 = v551
	goto L156
L156:
	;
	if base.Ui32(v572) < base.Ui32(v568) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v574 = v568
	goto L159
L158:
	;
	v574 = v572
	goto L159
L159:
	;
	if base.Ui32(v574) < base.Ui32(v566) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v576 = v566
	goto L162
L161:
	;
	v576 = v574
	goto L162
L162:
	;
	if base.Ui32(v576) < base.Ui32(v564) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v578 = v564
	goto L165
L164:
	;
	v578 = v576
	goto L165
L165:
	;
	v579 = int32(4)
	v580 = v553 + v579
	v582 = v552 + v579
	if v582 != v539&int32(2147483644) {
		v551 = v578
		v552 = v582
		v553 = v580
		goto L152
	} else {
		goto L166
	}
L166:
	;
	goto L153
L167:
	;
	v588 = v578
	v590 = v580
	goto L151
L168:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v597+v603<<(uint(int32(2))%32))))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	if base.Ui32(v601) < base.Ui32(v614) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v624 = v616
	goto L148
L170:
	;
	v616 = v614
	goto L172
L171:
	;
	v616 = v601
	goto L172
L172:
	;
	v617 = int32(1)
	v620 = v599 + v617
	if v620 != v541 {
		v599 = v620
		v601 = v616
		v603 = v603 + v617
		goto L168
	} else {
		goto L173
	}
L173:
	;
	goto L169
L174:
	;
	v672 = v667
	v674 = v500
	v675 = v501
	goto L77
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v296
	F_errmsg_internal(m, int32(_a_F_create_unique_paths_1), v26+int32(32))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_create_unique_paths_2), int32(_a_F_create_unique_paths_3), int32(_a_F_create_unique_paths_4))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L9
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v170
	F_errmsg_internal(m, int32(_a_F_create_unique_paths_5), v26)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L9
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_create_unique_paths_2), int32(_a_F_create_unique_paths_6), int32(_a_F_create_unique_paths_4))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v170
	F_errmsg_internal(m, int32(_a_F_create_unique_paths_7), v26+int32(16))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_create_unique_paths_2), int32(_a_F_create_unique_paths_8), int32(_a_F_create_unique_paths_4))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L9
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v747 != int32(1) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_set_cheapest(m, v56)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L9
	} else {
		goto L262
	}
L186:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v750 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	v755 = F_is_parallel_safe(m, l0, v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	if v755 == int32(0) {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+12))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v763 = F_palloc0(m, int32(304))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = int32(270)
	base.MemoryCopy(m, v763, l1, int32(304))
	v769 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v763)+60)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v763)+52)) = v769
	*(*int64)(unsafe.Add(mBase, uint32(v763)+44)) = v769
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v776 = *(*float64)(unsafe.Add(mBase, uint32(v761)+32))
	v777 = int32(0)
	v779 = F_estimate_num_groups(m, l0, v775, v776, v777, v777)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v763)+16)) = v779
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v763)+40)) = v782
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+45)))
	if v784 != int32(1) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+46)))
	if v969 == int32(1) {
		goto L252
	} else {
		goto L253
	}
L193:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v787 == int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	if v790 <= int32(0) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	v807 = int32(0)
	goto L196
L196:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v787)+12))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v817+v807<<(uint(int32(2))%32))))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+64))
	v824 = v26 + int32(44)
	if v729 == v822 {
		goto L201
	} else {
		goto L202
	}
L197:
	;
	goto L192
L198:
	;
	v943 = v807 + int32(1)
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	if v943 < v944 {
		v807 = v943
		goto L196
	} else {
		goto L251
	}
L199:
	;
	if v902|base.B2i32(v761 == v821) == int32(0) {
		goto L231
	} else {
		goto L232
	}
L200:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v890
	v902 = int32(1)
	goto L199
L201:
	;
	if v729 != 0 {
		goto L200
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v729 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = int32(0)
	v902 = int32(1)
	goto L199
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = int32(0)
	v902 = int32(1)
	goto L199
L206:
	;
	goto L207
L207:
	;
	if v822 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v842 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v842
	v902 = v842
	goto L199
L209:
	;
	goto L210
L210:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v822)+4))
	v846 = int32(0)
	if v846 < v845 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v849 = v845
	goto L213
L212:
	;
	v849 = v846
	goto L213
L213:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v855 = int32(0)
	goto L214
L214:
	;
	if v855 < v850 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	v866 = v862 + v855<<(uint(int32(2))%32)
	goto L218
L217:
	;
	v866 = int32(0)
	goto L218
L218:
	;
	if v855 == v849 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v849
	v902 = base.B2i32(v866 == int32(0))
	goto L199
L220:
	;
	goto L221
L221:
	;
	v872 = base.B2i32(v866 == int32(0))
	if v866 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v855
	v902 = v872
	goto L199
L223:
	;
	goto L224
L224:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v822)+12))
	if v876 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v855
	v902 = v872
	goto L199
L226:
	;
	goto L227
L227:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v876+v855<<(uint(int32(2))%32))))
	if v880 != v884 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v855
	v902 = int32(0)
	goto L199
L229:
	;
	v855 = v855 + int32(1)
	goto L214
L231:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v907 == int32(0) {
		goto L198
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v763)+40))
	v917 = F_create_projection_path(m, l0, v763, v821, v916)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L9
	} else {
		goto L237
	}
L234:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_unique_paths[1])))
	if v911&int32(1) == int32(0) {
		goto L198
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	if v729 != 0 {
		goto L246
	} else {
		goto L247
	}
L237:
	;
	if v902 != 0 {
		v931 = v917
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v919 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v928 = F_create_incremental_sort_path(m, l0, v763, v917, v729, v919, float64(-1))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L9
	} else {
		goto L245
	}
L240:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_unique_paths[1])))
	if v921&int32(1) != 0 {
		goto L239
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v925 = F_create_sort_path(m, v763, v917, v729, float64(-1))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L9
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	v931 = v925
	goto L236
L245:
	;
	v931 = v928
	goto L236
L246:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v934 = v932
	goto L248
L247:
	;
	v934 = int32(0)
	goto L248
L248:
	;
	v935 = *(*float64)(unsafe.Add(mBase, uint32(v763)+16))
	v936 = F_create_unique_path(m, v763, v931, v934, v935)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L9
	} else {
		goto L249
	}
L249:
	;
	F_add_partial_path(m, v763, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L9
	} else {
		goto L250
	}
L250:
	;
	goto L198
L251:
	;
	goto L197
L252:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v763)+40))
	v973 = F_create_projection_path(m, l0, v763, v761, v972)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L9
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v763)+52))
	if v985 == int32(0) {
		goto L185
	} else {
		goto L258
	}
L255:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	v977 = int32(0)
	v980 = *(*float64)(unsafe.Add(mBase, uint32(v763)+16))
	v981 = F_create_agg_path(m, l0, v763, v973, v975, int32(2), v977, v734, v977, v977, v980)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L9
	} else {
		goto L256
	}
L256:
	;
	F_add_partial_path(m, v763, v981)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L9
	} else {
		goto L257
	}
L257:
	;
	goto L254
L258:
	;
	F_generate_useful_gather_paths(m, l0, v763, int32(1))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	F_set_cheapest(m, v763)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L9
	} else {
		goto L260
	}
L260:
	;
	F_create_final_unique_paths(m, l0, v763, v729, v734, l2, v56)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L9
	} else {
		goto L261
	}
L261:
	;
	goto L185
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+200)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(l1)+196)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(l1)+192)) = v56
	v1029 = v56
	goto L12
}
func F_crlf_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	if l3 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v10 = l2 + l3
	v13 = l2
	goto L4
L4:
	;
	v17 = v10 - v13
	v18 = int32(0)
	if base.B2i32(v13&int32(3) == v18)|base.B2i32(v17 == v18) != 0 {
		v48 = v13
		v50 = v17
		v51 = base.B2i32(v17 != v18)
		goto L11
	} else {
		goto L12
	}
L5:
	;
	return v165
L6:
	;
	goto L5
L7:
	;
	if base.Ui32(v10) <= base.Ui32(v136) {
		v159 = v135
		v160 = v136
		goto L42
	} else {
		goto L43
	}
L8:
	;
	if v122 != 0 {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v122 = int32(0)
	goto L8
L10:
	;
	v100 = v93
	v102 = v95
	goto L27
L11:
	;
	if v51 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L12:
	;
	v31 = v13
	v33 = v17
	goto L13
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v36 == int32(10) {
		v93 = v31
		v95 = v33
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v48 = v43
	v50 = v39
	v51 = v41
	goto L11
L15:
	;
	v38 = int32(1)
	v39 = v33 - v38
	v40 = int32(0)
	v41 = base.B2i32(v39 != v40)
	v43 = v31 + v38
	if v43&int32(3) == v40 {
		v48 = v43
		v50 = v39
		v51 = v41
		goto L11
	} else {
		goto L16
	}
L16:
	;
	if v39 != 0 {
		v31 = v43
		v33 = v39
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(int32(10) == v57)|base.B2i32(base.Ui32(v50) < base.Ui32(int32(4))) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = v48
	v68 = v50
	goto L22
L20:
	;
	v86 = v48
	v88 = v50
	goto L21
L21:
	;
	if v88 == int32(0) {
		goto L9
	} else {
		goto L26
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v73 = v72 ^ int32(168430090)
	v76 = int32(-2139062144)
	if (int32(16843008)-v73|v73)&v76 != v76 {
		v93 = v66
		v95 = v68
		goto L10
	} else {
		goto L24
	}
L23:
	;
	v86 = v81
	v88 = v83
	goto L21
L24:
	;
	v80 = int32(4)
	v81 = v66 + v80
	v83 = v68 - v80
	if base.Ui32(int32(3)) < base.Ui32(v83) {
		v66 = v81
		v68 = v83
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v93 = v86
	v95 = v88
	goto L10
L27:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if int32(10) == v105 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L9
L29:
	;
	v122 = v100
	goto L8
L30:
	;
	goto L31
L31:
	;
	v107 = int32(1)
	v110 = v102 - v107
	if v110 != 0 {
		v100 = v100 + v107
		v102 = v110
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v123 = v122
	goto L35
L34:
	;
	v123 = v10
	goto L35
L35:
	;
	v124 = v123 - v13
	if v124 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v135 = int32(0)
	v136 = v13
	goto L7
L37:
	;
	goto L38
L38:
	;
	v128 = F_pushf_write(m, l0, v13, v124)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int32(0)
L40:
	;
	if v128 < int32(0) {
		v165 = v128
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v135 = v128
	v136 = v13 + v124
	goto L7
L42:
	;
	if base.Ui32(v160) < base.Ui32(v10) {
		v13 = v160
		goto L4
	} else {
		goto L53
	}
L43:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v138 != int32(10) {
		v159 = v135
		v160 = v136
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v143 = v136
	goto L45
L45:
	;
	v148 = F_pushf_write(m, l0, int32(_a_F_crlf_process_0), int32(2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L39
	} else {
		goto L47
	}
L46:
	;
	v159 = v148
	v160 = v10
	goto L42
L47:
	;
	if v148 < int32(0) {
		v159 = v148
		v160 = v143
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v153 = v143 + int32(1)
	if base.Ui32(v153) < base.Ui32(v10) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v155 != int32(10) {
		v159 = v148
		v160 = v153
		goto L42
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L46
L52:
	;
	v143 = v153
	goto L45
L53:
	;
	v165 = v159
	goto L6
}
func F_crosstab(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
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
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v231 int64
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v335 int32
	_ = v335
	var v350 int64
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v427 int64
	_ = v427
	var v431 int32
	_ = v431
	var v450 int64
	_ = v450
	var v471 int64
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v498 int64
	_ = v498
	var v502 int32
	_ = v502
	var v519 int64
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v539 int64
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v597 int32
	_ = v597
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
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
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = F_text_to_cstring(m, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L174
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != int32(389) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)))
	if v37&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L170
	}
L10:
	;
	v47 = F_SPI_execute(m, v29, int32(1), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L1
	} else {
		goto L165
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L161
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L156
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L149
	}
L15:
	;
	m.G0 = v22 + int32(32)
	return int64(0)
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 != int32(3) {
		goto L11
	} else {
		goto L23
	}
L17:
	;
	if v47 == int32(5) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_crosstab[1]))
	if v52 != int64(0) {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v56 = F_SPI_finish(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(2)
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
	goto L15
L23:
	;
	v71 = F_get_call_result_type(m, l0, int32(0), v22+int32(28))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v71 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v71 == int32(3) {
		goto L12
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 <= int32(1) {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_crosstab_0), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(444), int32(_a_F_crosstab_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v97 = int32(3)
	v99 = v93 + v94<<(uint(v97)%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+104))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v104 = v64 + v101<<(uint(v97)%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+104))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v99)+96))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+96))
	if base.B2i32(v100 != v105)&base.B2i32(int32(0) <= v100)|base.B2i32(v110 != v111) != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104)+304))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v104)+296))
	v119 = int32(1)
	goto L35
L35:
	;
	v140 = v99 + int32(28) + v119*int32(100)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+76))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)+68))
	if v117 == v142 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v180 = int32(_a_F_crosstab_3)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v41
	v184 = F_CreateTupleDescCopy(m, v93)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L37:
	;
	v178 = v119 + int32(1)
	if v178 != v94 {
		v119 = v178
		goto L35
	} else {
		goto L49
	}
L38:
	;
	if base.B2i32(v141 == v116)|base.B2i32(v141 < int32(0)) != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v148 = v142
	goto L40
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v148 = v117
	goto L40
L42:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v160 = F_format_type_with_typemod(m, v117, v116)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v162 = F_format_type_with_typemod(m, v148, v141)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v119 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v160
	v170 = F_errdetail(m, int32(_a_F_crosstab_5), v22)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1571), int32(_a_F_crosstab_6))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	goto L36
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v184
	v187 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_crosstab[3]))
	v196 = F_tuplestore_begin_heap(m, int32(base.Ui32(v188&int32(4))>>(uint(int32(2))%32)), v187, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_crosstab[2])) = v181
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v201 = F_TupleDescGetAttInMetadata(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = int32(2)
	v209 = int32(1)
	v219 = v187
	v221 = v209
	v231 = int64(0)
	goto L53
L53:
	;
	v234 = F_palloc0(m, v204<<(uint(v205)%32))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = int32(2)
	v607 = F_SPI_finish(m)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L148
	}
L55:
	;
	if v204 < int32(2) {
		v471 = v231
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v542 = int32(0)
	if v524 != 0 {
		goto L132
	} else {
		goto L133
	}
L57:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v524 = v522
	v539 = v519
	goto L56
L58:
	;
	F_pfree(m, v219)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L131
	}
L59:
	;
	v474 = F_BuildTupleFromCStrings(m, v201, v234)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L127
	}
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238+base.I32_wrap_i64(v231)<<(uint(int32(2))%32))))
	v245 = F_SPI_getvalue(m, v243, v64, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v245 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v247 = F_pstrdup(m, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v250 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v250
	if v221 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v250 = v247
	goto L64
L66:
	;
	v471 = v450 - int64(1)
	goto L59
L67:
	;
	F_pfree(m, v411)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L126
	}
L68:
	;
	v322 = F_SPI_getvalue(m, v243, v64, int32(3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L95
	}
L69:
	;
	if v250 != 0 {
		v450 = v231
		goto L66
	} else {
		goto L94
	}
L70:
	;
	if v250 == int32(0) {
		v411 = v245
		v427 = v231
		goto L67
	} else {
		goto L85
	}
L71:
	;
	if v245 == int32(0) {
		goto L69
	} else {
		goto L84
	}
L72:
	;
	if v245|v219 == int32(0) {
		v524 = v250
		v539 = v231
		goto L56
	} else {
		goto L73
	}
L73:
	;
	v255 = int32(0)
	if base.B2i32(v219 == v255)|base.B2i32(v245 == v255) != 0 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if base.B2i32(v262 == int32(0))|base.B2i32(v262 != v265) != 0 {
		v283 = v262
		v284 = v265
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v283-v284 != 0 {
		goto L70
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v268 = v219
	v269 = v245
	goto L78
L78:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)))
	if v273 == int32(0) {
		v283 = v273
		v284 = v272
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v283 = v273
	v284 = v272
	goto L76
L80:
	;
	v276 = int32(1)
	if v273 == v272 {
		v268 = v268 + v276
		v269 = v269 + v276
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_pfree(m, v245)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v498 = v231
	goto L58
L84:
	;
	goto L70
L85:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if base.B2i32(v294 == int32(0))|base.B2i32(v294 != v297) != 0 {
		v315 = v294
		v316 = v297
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v315-v316 != 0 {
		v411 = v245
		v427 = v231
		goto L67
	} else {
		goto L93
	}
L87:
	;
	goto L86
L88:
	;
	v300 = v245
	v301 = v250
	goto L89
L89:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v305 == int32(0) {
		v315 = v305
		v316 = v304
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v315 = v305
	v316 = v304
	goto L87
L91:
	;
	v308 = int32(1)
	if v305 == v304 {
		v300 = v300 + v308
		v301 = v301 + v308
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v320 = int32(0)
	goto L68
L94:
	;
	v320 = int32(1)
	goto L68
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v322
	if v320 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v245)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v329 = v231 + base.I64_extend_i32_u(base.B2i32(v204 != v205))
	if v204 == int32(2) {
		v471 = v329
		goto L59
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	if base.Ui64(v52) <= base.Ui64(v329) {
		v471 = v329
		goto L59
	} else {
		goto L101
	}
L101:
	;
	v335 = int32(1)
	v350 = v329
	goto L102
L102:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353+base.I32_wrap_i64(v350)<<(uint(int32(2))%32))))
	v360 = F_SPI_getvalue(m, v358, v64, int32(1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v471 = v406
	goto L59
L104:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	if v360 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v399 = F_SPI_getvalue(m, v358, v64, int32(3))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L119
	}
L106:
	;
	if v362 == int32(0) {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v362 == int32(0) {
		v411 = v360
		v427 = v350
		goto L67
	} else {
		goto L110
	}
L109:
	;
	v450 = v350
	goto L66
L110:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if base.B2i32(v371 == int32(0))|base.B2i32(v371 != v374) != 0 {
		v392 = v371
		v393 = v374
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v392-v393 != 0 {
		v411 = v360
		v427 = v350
		goto L67
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v377 = v360
	v378 = v362
	goto L114
L114:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378)+1)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+1)))
	if v382 == int32(0) {
		v392 = v382
		v393 = v381
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v392 = v382
	v393 = v381
	goto L112
L116:
	;
	v385 = int32(1)
	if v382 == v381 {
		v377 = v377 + v385
		v378 = v378 + v385
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	goto L105
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234+v335<<(uint(int32(2))%32))+4)) = v399
	if v360 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_pfree(m, v360)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v406 = base.I64_extend_i32_u(base.B2i32(v335 < v204-v205)) + v350
	v408 = v335 + int32(1)
	if v204-v209 <= v408 {
		v471 = v406
		goto L59
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	if base.Ui64(v406) < base.Ui64(v52) {
		v335 = v408
		v350 = v406
		goto L102
	} else {
		goto L125
	}
L125:
	;
	goto L103
L126:
	;
	v471 = v427 - int64(1)
	goto L59
L127:
	;
	F_tuplestore_puttuple(m, v196, v474)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_pfree(m, v474)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v219 == int32(0) {
		v519 = v471
		goto L57
	} else {
		goto L130
	}
L130:
	;
	v498 = v471
	goto L58
L131:
	;
	v519 = v498
	goto L57
L132:
	;
	v544 = F_pstrdup(m, v524)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v546 = v542
	goto L134
L134:
	;
	if int32(0) < v204 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v546 = v544
	goto L134
L136:
	;
	v549 = v542
	goto L139
L137:
	;
	goto L138
L138:
	;
	F_pfree(m, v234)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L146
	}
L139:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v234+v549<<(uint(int32(2))%32))))
	if v571 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L138
L141:
	;
	F_pfree(m, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v575 = v549 + int32(1)
	if v575 != v204 {
		v549 = v575
		goto L139
	} else {
		goto L145
	}
L144:
	;
	goto L143
L145:
	;
	goto L140
L146:
	;
	v600 = v539 + int64(1)
	if base.Ui64(v600) < base.Ui64(v52) {
		v219 = v546
		v221 = int32(0)
		v231 = v600
		goto L53
	} else {
		goto L147
	}
L147:
	;
	goto L54
L148:
	;
	goto L15
L149:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v644 = F_format_type_with_typemod(m, v111, v105)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v646 = F_format_type_with_typemod(m, v110, v100)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v646
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v644
	v653 = F_errdetail(m, int32(_a_F_crosstab_7), v22+int32(16))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1549), int32(_a_F_crosstab_6))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_crosstab_4), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v673 = F_errdetail(m, int32(_a_F_crosstab_8), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(1534), int32(_a_F_crosstab_6))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_crosstab_9), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(438), int32(_a_F_crosstab_2))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errmsg(m, int32(_a_F_crosstab_10), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v709 = F_errdetail(m, int32(_a_F_crosstab_11), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(425), int32(_a_F_crosstab_2))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errmsg(m, int32(_a_F_crosstab_12), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(388), int32(_a_F_crosstab_2))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errmsg(m, int32(_a_F_crosstab_13), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_crosstab_1), int32(384), int32(_a_F_crosstab_2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
