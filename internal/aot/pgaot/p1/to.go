package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginCopyTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v309 int32
	_ = v309
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v427 int32
	_ = v427
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v548 int32
	_ = v548
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v584 int32
	_ = v584
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v636 int32
	_ = v636
	var v649 int32
	_ = v649
	var v661 int32
	_ = v661
	var v672 int32
	_ = v672
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v709 int32
	_ = v709
	var v720 int32
	_ = v720
	var v732 int32
	_ = v732
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v760 int32
	_ = v760
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v798 int32
	_ = v798
	var v810 int32
	_ = v810
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v839 int32
	_ = v839
	var v850 int32
	_ = v850
	var v862 int32
	_ = v862
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v950 int32
	_ = v950
	var v961 int32
	_ = v961
	var v973 int32
	_ = v973
	var v986 int32
	_ = v986
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1007 int32
	_ = v1007
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1147 int32
	_ = v1147
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1170 int32
	_ = v1170
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1224 int32
	_ = v1224
	var v1235 int32
	_ = v1235
	var v1259 int32
	_ = v1259
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1306 int32
	_ = v1306
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1345 int32
	_ = v1345
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1397 int32
	_ = v1397
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1422 int32
	_ = v1422
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1452 int32
	_ = v1452
	var v1463 int32
	_ = v1463
	var v1475 int32
	_ = v1475
	var v1488 int32
	_ = v1488
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1542 int32
	_ = v1542
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1582 int32
	_ = v1582
	var v1594 int32
	_ = v1594
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1617 int32
	_ = v1617
	var v1636 int32
	_ = v1636
	var v1649 int32
	_ = v1649
	var v1662 int32
	_ = v1662
	var v1672 int32
	_ = v1672
	var v1681 int32
	_ = v1681
	var v1688 int32
	_ = v1688
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1718 int32
	_ = v1718
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1743 int32
	_ = v1743
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1773 int32
	_ = v1773
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1799 int32
	_ = v1799
	var v1812 int32
	_ = v1812
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2040 int64
	_ = v2040
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2100 int32
	_ = v2100
	var v2101 int64
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	v6 = l5
	v9 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(496)
	m.G0 = v28
	v42 = v9
	v43 = v9
	v44 = v9
	v45 = v9
	v46 = v9
	v47 = v9
	v48 = v9
	v49 = v9
	v50 = int32(-1)
	v51 = v9
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1822)+176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[0])) = v1828
	m.G0 = v28 + int32(496)
	return v1822
L4:
	;
	if v50 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v2100 = int32(m.ExcTag)
	v2101 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2100 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L7:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1827)))
	if v1838 != 0 {
		goto L236
	} else {
		goto L237
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+456)) = int64(21474836484)
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyTo[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+440)) = v63
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyTo[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+432)) = v66
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v1520 = v42
	v1521 = v43
	v1522 = v44
	v1523 = v45
	v1524 = v46
	v1525 = v47
	v1526 = v48
	v1527 = v49
	v1529 = v51
	goto L10
L10:
	;
	if v1529 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	v451 = F_palloc0(m, int32(184))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L6
	} else {
		goto L48
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+119)))
	switch v71 - int32(83) {
	case 0:
		goto L15
	default:
		goto L13
	case 19:
		goto L16
	case 26:
		goto L17
	case 29:
		goto L14
	case 31:
		goto L11
	case 35:
		goto L18
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L44
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L6
	} else {
		goto L39
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L35
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L30
	}
L17:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+129)))
	if v139 != 0 {
		goto L11
	} else {
		goto L24
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(151027844))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v97 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_0), v28+int32(96))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errhint(m, int32(_a_F_BeginCopyTo_1), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(654), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(1088))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v163 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_4), v28+int32(112))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errhint(m, int32(_a_F_BeginCopyTo_5), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(662), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(151027844))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v228 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_6), v28+int32(128))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errhint(m, int32(_a_F_BeginCopyTo_1), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(669), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(151027844))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = v293 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_7), v28+int32(144))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(674), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(151027844))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v346 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_8), v28+int32(160))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errhint(m, int32(_a_F_BeginCopyTo_1), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(680), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errcode(m, int32(151027844))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v411 + int32(4)
	F_errmsg(m, int32(_a_F_BeginCopyTo_9), v28+int32(80))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v42
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(685), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[0]))
	v467 = F_AllocSetContextCreateInternal(m, v462, int32(_a_F_BeginCopyTo_10), int32(0), int32(_a_F_BeginCopyTo_11), int32(_a_F_BeginCopyTo_12))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+164)) = v467
	v470 = int32(_a_F_BeginCopyTo_13)
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[0])) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	v483 = v451 + int32(48)
	F_ProcessCopyOptions(m, l0, v483, int32(0), l7)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+54)))
	if v487 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v493 = int32(_a_F_BeginCopyTo_14)
	goto L53
L52:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+52)))
	if v491 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451))) = v493
	if l1 != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v492 = int32(_a_F_BeginCopyTo_15)
	goto L56
L55:
	;
	v492 = int32(_a_F_BeginCopyTo_16)
	goto L56
L56:
	;
	v493 = v492
	goto L53
L57:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1076)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1086 = v451 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	v1088 = F_CopyGetAttnums(m, v1077, v1072, l6)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L6
	} else {
		goto L137
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+24)) = l1
	v1072 = l1
	v1076 = l1 + int32(52)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v496 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v451)+24)) = v496
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v510 = F_pg_analyze_and_rewrite_fixedparams(m, l2, v498, v496, v496, v496)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v510 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if int32(2) <= v563 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_17), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(739), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L1
L69:
	;
	v584 = int32(0)
	goto L74
L70:
	;
	goto L71
L71:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+28))
	if v747 != 0 {
		goto L90
	} else {
		goto L91
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L6
	} else {
		goto L86
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L6
	} else {
		goto L82
	}
L74:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v562+v584<<(uint(int32(2))%32))))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+8))
	switch v596 - int32(3) {
	case 0:
		goto L73
	case 1:
		goto L72
	default:
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L78
	}
L76:
	;
	v600 = v584 + int32(1)
	if v600 != v563 {
		v584 = v600
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_18), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(762), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_19), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(753), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	goto L1
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_20), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(757), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	goto L1
L90:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L6
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v824 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	if v748 == int32(242) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_22), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(772), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	goto L1
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(778), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	goto L1
L102:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v887 = F_pg_plan_query(m, v746, v876, int32(2048), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L6
	} else {
		goto L109
	}
L103:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v746)+96))
	if v827 != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(1088))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_23), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(794), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	goto L1
L109:
	;
	if l3 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[3]))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)))
	goto L130
L111:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v887)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v900 = int32(0)
	if v891 == v900 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v938 != 0 {
		goto L110
	} else {
		goto L125
	}
L113:
	;
	v938 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v891)+4))
	if v906 <= int32(0) {
		v932 = v900
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v938 = v932
	goto L112
L117:
	;
	v909 = int32(0)
	if v909 < v906 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v912 = v906
	goto L120
L119:
	;
	v912 = v909
	goto L120
L120:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	v915 = int32(0)
	goto L121
L121:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v913+v915<<(uint(int32(2))%32))))
	v924 = base.B2i32(v923 == l3)
	if v923 == l3 {
		v932 = v924
		goto L116
	} else {
		goto L123
	}
L122:
	;
	v932 = v924
	goto L116
L123:
	;
	v926 = v915 + int32(1)
	if v926 != v912 {
		v915 = v926
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(325))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_24), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(823), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	goto L1
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_PushCopiedSnapshot(m, v997)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1027 = F_CreateDestReceiver(m, int32(8))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1027)+20)) = v451
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1040 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[3]))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1050 = int32(0)
	v1054 = F_CreateQueryDesc(m, v887, v1030, v1041, v1050, v1027, v1050, v1050, v1050)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+28)) = v1054
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_ExecutorStart(m, v1054, int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v451)+24))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v451)+28))
	v1072 = v1068
	v1076 = v1069 + int32(36)
	goto L57
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+32)) = v1088
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1100 = F_palloc0(m, v1091)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+100)) = v1100
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+96)))
	if v1103 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	if v1306 < int32(0) {
		goto L171
	} else {
		goto L172
	}
L140:
	;
	if v1091 == int32(0) {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v451)+92))
	if v1110 == int32(0) {
		goto L139
	} else {
		goto L144
	}
L143:
	;
	base.MemoryFill(m, v1100, int32(1), v1091)
	goto L139
L144:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1122 = F_CopyGetAttnums(m, v1077, v1113, v1110)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	if v1122 == int32(0) {
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v1126 = int32(0)
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+4))
	if v1127 <= v1126 {
		goto L139
	} else {
		goto L147
	}
L147:
	;
	v1147 = v1126
	goto L148
L148:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+12))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1155+v1147<<(uint(int32(2))%32))))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v451)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1170 = int32(0)
	if v1161 == v1170 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L139
L150:
	;
	v1210 = v1159 - int32(1)
	if v1208 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L151:
	;
	v1208 = int32(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+4))
	if v1176 <= int32(0) {
		v1202 = v1170
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1208 = v1202
	goto L150
L155:
	;
	v1179 = int32(0)
	if v1179 < v1176 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1182 = v1176
	goto L158
L157:
	;
	v1182 = v1179
	goto L158
L158:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+12))
	v1185 = int32(0)
	goto L159
L159:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1183+v1185<<(uint(int32(2))%32))))
	v1194 = base.B2i32(v1193 == v1159)
	if v1193 == v1159 {
		v1202 = v1194
		goto L154
	} else {
		goto L161
	}
L160:
	;
	v1202 = v1194
	goto L154
L161:
	;
	v1196 = v1185 + int32(1)
	if v1196 != v1182 {
		v1185 = v1196
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L6
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v451)+100))
	v1275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1273+v1210))) = uint8(v1275)
	v1278 = v1147 + v1275
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+4))
	if v1278 < v1279 {
		v1147 = v1278
		goto L148
	} else {
		goto L170
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(_a_F_BeginCopyTo_25))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v1077 + v1160<<(uint(int32(4))%32) + v1210*int32(100) + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = int32(_a_F_BeginCopyTo_26)
	F_errmsg(m, int32(_a_F_BeginCopyTo_27), v28-int32(-64))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(881), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L6
	} else {
		goto L169
	}
L169:
	;
	goto L1
L170:
	;
	goto L149
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1318 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[4]))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+4))
	goto L174
L172:
	;
	v1320 = v49
	v1321 = v1306
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+16)) = v1321
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1332 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[5]))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+4))
	goto L175
L174:
	;
	v1320 = v1319
	v1321 = v1319
	goto L173
L175:
	;
	v1334 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v451)+4)) = v1334
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v451)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+21)) = uint8(base.B2i32(base.Ui32(v1336-int32(35)) < base.Ui32(int32(7))))
	v1345 = base.B2i32(v1321 != v1333) & base.B2i32(v1336 != v1334)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+20)) = uint8(v1345)
	if l4 == v1334 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+440)) = int64(3)
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[6]))
	if v1352 == int32(2) {
		v1822 = v451
		v1823 = v43
		v1824 = v44
		v1825 = v45
		v1826 = v46
		v1827 = v1086
		v1828 = v471
		v1829 = v1320
		goto L7
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1366 = F_pstrdup(m, l4)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L6
	} else {
		goto L180
	}
L179:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v451)+8)) = v1356
	v1822 = v451
	v1823 = v43
	v1824 = v44
	v1825 = v45
	v1826 = v46
	v1827 = v1086
	v1828 = v471
	v1829 = v1320
	goto L7
L180:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+40)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v451)+36)) = v1366
	v1371 = v451 + int32(36)
	if v6 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v28)+440)) = int64(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1383 = F_OpenPipeStream(m, v1366, int32(_a_F_BeginCopyTo_28))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L6
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+440)) = int64(1)
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v1438 != int32(47) {
		goto L190
	} else {
		goto L191
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v451)+8)) = v1383
	if v1383 != 0 {
		v1822 = v451
		v1823 = v1371
		v1824 = v44
		v1825 = v45
		v1826 = v46
		v1827 = v1086
		v1828 = v471
		v1829 = v1320
		goto L7
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L6
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode_for_file_access(m)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1371)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v1408
	F_errmsg(m, int32(_a_F_BeginCopyTo_29), v28+int32(48))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(934), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L6
	} else {
		goto L189
	}
L189:
	;
	goto L1
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L6
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	v1498 = F_umask(m, int32(18))
	mBase = m.M
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[8]))
	v1503 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[9]))
	goto L197
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errcode(m, int32(33579140))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errmsg(m, int32(_a_F_BeginCopyTo_30), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L6
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1320
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1086
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v451
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(950), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	goto L1
L197:
	;
	v1505 = v28 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v1505)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1505))) = v28 + int32(172)
	goto L200
L198:
	;
	v1520 = v451
	v1521 = v1371
	v1522 = v1503
	v1523 = v1501
	v1524 = v1498
	v1525 = v1086
	v1526 = v471
	v1527 = v1320
	v1529 = int32(0)
	goto L10
L200:
	;
	goto L198
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+60))
	if v1681 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[9])) = v28 + int32(176)
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	v1552 = F_AllocateFile(m, v1542, int32(_a_F_BeginCopyTo_28))
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L6
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[8])) = v1523
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[9])) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	v1662 = F_umask(m, v1524)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_pg_re_throw(m)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L6
	} else {
		goto L215
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1520)+8)) = v1552
	v1555 = int32(_a_F_BeginCopyTo_31)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[8])) = v1523
	v1557 = int32(_a_F_BeginCopyTo_32)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[9])) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	v1567 = F_umask(m, v1524)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[8])) = v1523
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[9])) = v1522
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+8))
	if v1572 != 0 {
		goto L201
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	v1582 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L6
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errcode_for_file_access(m)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L6
	} else {
		goto L208
	}
L208:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v1605
	F_errmsg(m, int32(_a_F_BeginCopyTo_33), v28)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	if base.B2i32(v1582 != int32(44))&base.B2i32(v1582 != int32(2)) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errhint(m, int32(_a_F_BeginCopyTo_34), int32(0))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L6
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(973), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L6
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	goto L1
L215:
	;
	goto L1
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	if v1688 < int32(0) {
		goto L221
	} else {
		goto L222
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[10])) = int32(8)
	v1688 = int32(-1)
	goto L219
L218:
	;
	v1688 = v1681
	goto L219
L219:
	;
	goto L216
L220:
	;
	if v1706 != 0 {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	v1702 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v1706 = v1702
	goto L220
L222:
	;
	goto L223
L223:
	;
	v1705 = F___fstatat(m, v1688, int32(_a_F_BeginCopyTo_35), v28+int32(336), int32(_a_F_BeginCopyTo_36))
	mBase = m.M
	v1706 = v1705
	goto L220
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L6
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v28)+340))
	if v1757&int32(_a_F_BeginCopyTo_37) != int32(_a_F_BeginCopyTo_38) {
		v1822 = v1520
		v1823 = v1521
		v1824 = v1522
		v1825 = v1523
		v1826 = v1524
		v1827 = v1525
		v1828 = v1526
		v1829 = v1527
		goto L7
	} else {
		goto L231
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errcode_for_file_access(m)
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v1729
	F_errmsg(m, int32(_a_F_BeginCopyTo_39), v28+int32(32))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(980), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	goto L1
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L6
	} else {
		goto L233
	}
L233:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1521)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v1785
	F_errmsg(m, int32(_a_F_BeginCopyTo_40), v28+int32(16))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1522
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1523
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1521
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1525
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1520
	F_errfinish(m, int32(_a_F_BeginCopyTo_2), int32(985), int32(_a_F_BeginCopyTo_3))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	goto L1
L236:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+56))
	v1841 = v1839
	goto L238
L237:
	;
	v1841 = int32(0)
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1826
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1827
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1828
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1822
	v1853 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[11]))
	if v1853 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+468)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v28)+464)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v28)+472)) = v1826
	*(*int32)(unsafe.Add(mBase, uint32(v28)+476)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v28)+480)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v28)+484)) = v1827
	*(*int32)(unsafe.Add(mBase, uint32(v28)+488)) = v1828
	*(*int32)(unsafe.Add(mBase, uint32(v28)+492)) = v1822
	goto L245
L240:
	;
	goto L239
L241:
	;
	v1857 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyTo[12])))
	if v1857&int32(1) == int32(0) {
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1862 = int32(_a_F_BeginCopyTo_41)
	v1864 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13]))
	v1865 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13])) = v1864 + v1865
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	*(*int32)(unsafe.Add(mBase, uint32(v1853))) = v1868 + v1865
	*(*int32)(unsafe.Add(mBase, uint32(v1853)+220)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v1853)+224)) = v1841
	base.MemoryFill(m, v1853+int32(232), int32(0), int32(160))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	*(*int32)(unsafe.Add(mBase, uint32(v1853))) = v1879 + v1865
	v1885 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13])) = v1885 - v1865
	goto L240
L243:
	;
	goto L5
L244:
	;
	goto L243
L245:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[11]))
	if v1911 == int32(0) {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1915 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyTo[12])))
	if v1915&int32(1) == int32(0) {
		goto L244
	} else {
		goto L247
	}
L247:
	;
	v1920 = int32(_a_F_BeginCopyTo_41)
	v1922 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13]))
	v1923 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13])) = v1922 + v1923
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1911)))
	*(*int32)(unsafe.Add(mBase, uint32(v1911))) = v1926 + v1923
	goto L249
L248:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1911)))
	v2057 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1911))) = v2056 + v2057
	v2060 = int32(_a_F_BeginCopyTo_41)
	v2062 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyTo[13])) = v2062 - v2057
	goto L244
L249:
	;
	goto L251
L251:
	;
	goto L252
L252:
	;
	v2021 = int32(0)
	v2024 = int32(0)
	goto L257
L257:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(456)+v2024<<(uint(int32(2))%32))))
	v2034 = int32(3)
	v2040 = *(*int64)(unsafe.Add(mBase, uint32(v28+int32(432)+v2024<<(uint(v2034)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1911+int32(232)+v2033<<(uint(v2034)%32)))) = v2040
	v2042 = int32(1)
	v2045 = v2021 + v2042
	if v2045 != int32(2) {
		v2021 = v2045
		v2024 = v2024 + v2042
		goto L257
	} else {
		goto L259
	}
L258:
	;
	goto L248
L259:
	;
	goto L258
L260:
	;
	v2105 = int32(v2101)
	m.G0 = v28
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2105)+4))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2105)))
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2108)))
	if v28+int32(172) == v2111 {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	m.ExcPending = 1
	goto L269
L262:
	;
	if v2115 != 0 {
		goto L266
	} else {
		goto L267
	}
L263:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2108)+4))
	v2115 = v2113
	goto L265
L264:
	;
	v2115 = int32(0)
	goto L265
L265:
	;
	goto L262
L266:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v28)+492))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v28)+488))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v28)+484))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v28)+480))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v28)+476))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v28)+472))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v28)+468))
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v28)+464))
	v42 = v2116
	v43 = v2120
	v44 = v2122
	v45 = v2123
	v46 = v2121
	v47 = v2118
	v48 = v2117
	v49 = v2119
	v50 = v2115
	v51 = v2107
	goto L2
L267:
	;
	goto L268
L268:
	;
	F___wasm_longjmp(m, v2108, v2107)
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	return int32(0)
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyToBinaryOneRow(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v16 = int32(8)
	v23 = v15<<(uint(v16)%32) | int32(base.Ui32(v15)>>(uint(v16)%32))
	goto L3
L2:
	;
	v23 = int32(0)
	goto L3
L3:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+6)) = uint16(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v25, v11+int32(6), int32(2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v31 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_CopySendEndOfRow(m, l0)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L20
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(0)
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v51 = v49 - int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v52))))
	if v54 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L6
L11:
	;
	v107 = v42 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v107 < v108 {
		v42 = v107
		goto L9
	} else {
		goto L19
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v57, v11+int32(8), int32(4))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v51<<(uint(int32(2))%32))))
	v71 = F_SendFunctionCall(m, v13+v51*int32(28), v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v76 = int32(4)
	v77 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - v76
	v78 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = base.I32_rotr(v77&v78, int32(8)) | base.I32_rotr(v77, int32(24))&v78
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v88, v11+int32(12), v76)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v95 = int32(4)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	F_appendBinaryStringInfo(m, v94, v71+v95, int32(base.Ui32(v97)>>(uint(int32(2))%32))-v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	goto L10
L20:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_coerce_to_specific_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_exprType(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l2 == v12 {
			v24 = l1
			v25 = F_expression_returns_set(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
							F_errmsg(m, int32(_a_F_coerce_to_specific_type_typmod_0), v10)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = F_exprLocation(m, v24)
								mBase = m.M
								F_parser_errposition(m, l0, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_coerce_to_specific_type_typmod_1), int32(1239), int32(_a_F_coerce_to_specific_type_typmod_2))
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
				} else {
					m.G0 = v10 + int32(32)
					return v24
				}
			}
		} else {
			v20 = F_coerce_to_target_type(m, l0, l1, v12, l2, l3, int32(1), int32(2), int32(-1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = F_format_type_be(m, l2)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_format_type_be(m, v12)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v40
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
									F_errmsg(m, int32(_a_F_coerce_to_specific_type_typmod_3), v10+int32(16))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = F_exprLocation(m, l1)
										mBase = m.M
										F_parser_errposition(m, l0, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_coerce_to_specific_type_typmod_1), int32(1229), int32(_a_F_coerce_to_specific_type_typmod_2))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
				} else {
					v24 = v20
					v25 = F_expression_returns_set(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
									F_errmsg(m, int32(_a_F_coerce_to_specific_type_typmod_0), v10)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v69 = F_exprLocation(m, v24)
										mBase = m.M
										F_parser_errposition(m, l0, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_coerce_to_specific_type_typmod_1), int32(1239), int32(_a_F_coerce_to_specific_type_typmod_2))
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
						} else {
							m.G0 = v10 + int32(32)
							return v24
						}
					}
				}
			}
		}
	}
}
func F_to_ascii_default(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_to_ascii_default[0]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v10 = F_encode_to_ascii(m, v3, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_to_bin64(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14006(m, l0, int64(1), int64(2), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_to_oct32(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14005(m, l0, int64(3), int64(8), int32(7))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_to_tsvector_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v17 == int32(1) {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
			if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
			} else {
				if v20 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				v44 = v31
				v46 = base.I32_div_u_s(v44, int32(6))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v46
				if base.Ui32(int32(11)) < base.Ui32(v44) {
					if base.Ui32(v44) < base.Ui32(int32(402653184)) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(67108863)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
				}
			}
		} else {
			v32 = int32(1)
			if v17&v32 != 0 {
				v44 = int32(base.Ui32(v17)>>(uint(v32)%32)) - v32
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
			}
			v46 = base.I32_div_u_s(v44, int32(6))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v46
			if base.Ui32(int32(11)) < base.Ui32(v44) {
				if base.Ui32(v44) < base.Ui32(int32(402653184)) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(67108863)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v63 = F_palloc(m, v60<<(uint(int32(4))%32))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v63
			v66 = int32(1)
			v67 = v13 + v66
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v72 = v70 & v66
			if v72 != 0 {
				v73 = v67
			} else {
				v73 = v13 + int32(4)
			}
			if v70 == int32(1) {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
				if v79 == int32(18) {
					v82 = int32(16)
				} else {
					v82 = int32(0)
				}
				if base.Ui32((v79-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v89 = int32(4)
				} else {
					v89 = v82
				}
				v100 = v89
			} else {
				v90 = int32(1)
				if v72 != 0 {
					v100 = int32(base.Ui32(v70)>>(uint(v90)%32)) - v90
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			F_parsetext(m, v11, v9, v73, v100)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v103 != v13 {
					F_pfree(m, v13)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = F_make_tsvector(m, v9)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v107
						}
					}
				} else {
					v107 = F_make_tsvector(m, v9)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v107
					}
				}
			}
		}
	}
}
