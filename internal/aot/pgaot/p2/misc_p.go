package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseISO8601Number(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v62 int32
	_ = v62
	v10 = int32(-1)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(int32(10)) <= base.Ui32((v11-int32(48))&int32(255)) {
		if base.Ui32(int32(1)) < base.Ui32((v11-int32(45))&int32(255)) {
			v62 = v10
			return v62
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
			v27 = F_strtod(m, l0, l1)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v31 == l0 {
					v62 = v10
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[155]))
					if v34 != 0 {
						v62 = v10
					} else {
						v35 = int32(-2)
						v36 = base.F64_abs(v27)
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v36)) {
							v62 = v35
						} else {
							if base.F64_gt(v36, float64(1e+15)) != 0 {
								v62 = v35
							} else {
								if base.F64_ge(v27, float64(0)) != 0 {
									v48 = base.F64_floor(v27)
								} else {
									v48 = base.F64_neg(base.F64_floor(base.F64_neg(v27)))
								}
								if base.F64_lt(base.F64_abs(v48), float64(9.223372036854776e+18)) != 0 {
									v52 = base.I64_trunc_f64_s(v48)
									v54 = v52
								} else {
									v54 = int64(-9223372036854775807 - 1)
								}
								*(*int64)(unsafe.Add(mBase, uint32(l2))) = v54
								*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_sub(v27, base.F64_convert_i64_s(v54))
								v62 = int32(0)
							}
						}
					}
				}
				return v62
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
		v27 = F_strtod(m, l0, l1)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v31 == l0 {
				v62 = v10
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[155]))
				if v34 != 0 {
					v62 = v10
				} else {
					v35 = int32(-2)
					v36 = base.F64_abs(v27)
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v36)) {
						v62 = v35
					} else {
						if base.F64_gt(v36, float64(1e+15)) != 0 {
							v62 = v35
						} else {
							if base.F64_ge(v27, float64(0)) != 0 {
								v48 = base.F64_floor(v27)
							} else {
								v48 = base.F64_neg(base.F64_floor(base.F64_neg(v27)))
							}
							if base.F64_lt(base.F64_abs(v48), float64(9.223372036854776e+18)) != 0 {
								v52 = base.I64_trunc_f64_s(v48)
								v54 = v52
							} else {
								v54 = int64(-9223372036854775807 - 1)
							}
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v54
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_sub(v27, base.F64_convert_i64_s(v54))
							v62 = int32(0)
						}
					}
				}
			}
			return v62
		}
	}
}
func F_PlannedStmtRequiresSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v4 == int32(0) {
		v12 = int32(1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		switch v8 - int32(158) {
		case 0, 1, 45, 64, 65, 66, 67, 86, 88, 89:
			v12 = int32(0)
		default:
			v12 = int32(1)
		}
	}
	return v12
}
func F_PostgresMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int64
	_ = v202
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int64
	_ = v299
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v521 int32
	_ = v521
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v595 int32
	_ = v595
	var v607 int32
	_ = v607
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v642 int32
	_ = v642
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v693 int64
	_ = v693
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v790 int64
	_ = v790
	var v811 int32
	_ = v811
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v832 int32
	_ = v832
	var v844 int32
	_ = v844
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v875 int32
	_ = v875
	var v887 int32
	_ = v887
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v918 int32
	_ = v918
	var v930 int32
	_ = v930
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v961 int32
	_ = v961
	var v973 int32
	_ = v973
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1004 int32
	_ = v1004
	var v1016 int32
	_ = v1016
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int64
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1243 int64
	_ = v1243
	var v1256 int32
	_ = v1256
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1679 int32
	_ = v1679
	var v1692 int32
	_ = v1692
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1716 int32
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1766 int32
	_ = v1766
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int64
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v19 = v3
	v20 = v3
	v22 = int32(-1)
	v24 = v14
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v22 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v1778 = int32(m.ExcTag)
	v1779 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1778 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L6:
	;
	v30 = int32(16)
	v31 = v24 - v30
	m.G0 = v31
	v34 = v31 - v30
	m.G0 = v34
	v37 = v34 - v30
	m.G0 = v37
	v40 = v37 - v30
	m.G0 = v40
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v42)
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v44)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	if v49 == v42 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v1732 = v20
	v1736 = v24
	v1740 = v19
	goto L8
L8:
	;
	if v1740 != 0 {
		goto L392
	} else {
		goto L393
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_BaseInit(m)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L236
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v54 = int32(913)
	v56 = m.G0
	v58 = v56 - int32(144)
	m.G0 = v58
	switch int32(915) {
	case 0, 2:
		v68 = v54
		goto L14
	default:
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v495 = int32(913)
	v497 = m.G0
	v499 = v497 - int32(144)
	m.G0 = v499
	switch int32(915) {
	case 0, 2:
		v509 = v495
		goto L111
	default:
		goto L112
	}
L13:
	;
	v96 = int32(914)
	v98 = m.G0
	v100 = v98 - int32(144)
	m.G0 = v100
	switch int32(916) {
	case 0, 2:
		v110 = v96
		goto L27
	default:
		goto L28
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v68
	F_sigemptyset(m, v58+int32(8))
	mBase = m.M
	goto L17
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[628])) = v54
	v68 = int32(4729)
	goto L14
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+136)) = int32(268435456)
	v80 = v58 + int32(4)
	goto L21
L19:
	;
	m.G0 = v58 + int32(144)
	goto L13
L21:
	;
	goto L22
L22:
	;
	if v80 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v91 = F___memcpy(m, int32(4608908), v80, int32(140))
	mBase = m.M
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L19
L26:
	;
	v138 = int32(295)
	v140 = m.G0
	v142 = v140 - int32(144)
	m.G0 = v142
	switch int32(297) {
	case 0, 2:
		v152 = v138
		goto L40
	default:
		goto L41
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v110
	F_sigemptyset(m, v100+int32(8))
	mBase = m.M
	goto L30
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[853])) = v96
	v110 = int32(4729)
	goto L27
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+136)) = int32(268435456)
	v122 = v100 + int32(4)
	goto L34
L32:
	;
	m.G0 = v100 + int32(144)
	goto L26
L34:
	;
	goto L35
L35:
	;
	if v122 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v133 = F___memcpy(m, int32(4609048), v122, int32(140))
	mBase = m.M
	goto L38
L37:
	;
	goto L38
L38:
	;
	goto L32
L39:
	;
	v179 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v179
	*(*int32)(unsafe.Add(mBase, _consts[525])) = v179
	v189 = v179
	goto L53
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v152
	F_sigemptyset(m, v142+int32(8))
	mBase = m.M
	goto L43
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[629])) = v138
	v152 = int32(4729)
	goto L40
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+136)) = int32(268435456)
	v164 = v142 + int32(4)
	goto L47
L45:
	;
	m.G0 = v142 + int32(144)
	goto L39
L47:
	;
	goto L48
L48:
	;
	if v164 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v175 = F___memcpy(m, int32(4610868), v164, int32(140))
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L45
L52:
	;
	v326 = int32(-2)
	v328 = m.G0
	v330 = v328 - int32(144)
	m.G0 = v330
	switch int32(0) {
	case 0, 2:
		v340 = v326
		goto L59
	default:
		goto L60
	}
L53:
	;
	v191 = int32(40)
	v192 = v189 * v191
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[526]))) = uint8(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[854]))) = v189
	v202 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[855]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[856]))) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[857]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[858]))) = v195
	*(*uint8)(unsafe.Add(mBase, uint32(v192)+uint32(_consts[527]))) = uint8(v195)
	v221 = v189 | int32(1)
	v223 = v221 * v191
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[526]))) = uint8(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[854]))) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[855]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[856]))) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[857]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[858]))) = v195
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+uint32(_consts[527]))) = uint8(v195)
	v252 = v189 | int32(2)
	v254 = v252 * v191
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[526]))) = uint8(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[854]))) = v252
	*(*int64)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[855]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[856]))) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[857]))) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[858]))) = v195
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+uint32(_consts[527]))) = uint8(v195)
	if base.B2i32(v189 == int32(20)) == v195 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v320 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[859])) = uint8(v320)
	F_pqsignal_be(m, int32(14), int32(1784))
	mBase = m.M
	goto L52
L55:
	;
	v287 = v189 | int32(3)
	v289 = v287 * int32(40)
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[526]))) = uint8(v292)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[854]))) = v287
	v299 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[855]))) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[856]))) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[857]))) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[858]))) = v292
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+uint32(_consts[527]))) = uint8(v292)
	v189 = v189 + int32(4)
	goto L53
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v368 = int32(916)
	v370 = m.G0
	v372 = v370 - int32(144)
	m.G0 = v372
	switch int32(918) {
	case 0, 2:
		v382 = v368
		goto L72
	default:
		goto L73
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v340
	F_sigemptyset(m, v330+int32(8))
	mBase = m.M
	goto L62
L60:
	;
	*(*int32)(unsafe.Add(mBase, _consts[860])) = v326
	v340 = int32(4729)
	goto L59
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+136)) = int32(268435456)
	v352 = v330 + int32(4)
	goto L66
L64:
	;
	m.G0 = v330 + int32(144)
	goto L58
L66:
	;
	goto L67
L67:
	;
	if v352 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v363 = F___memcpy(m, int32(4610588), v352, int32(140))
	mBase = m.M
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L64
L71:
	;
	v410 = int32(1037)
	v412 = m.G0
	v414 = v412 - int32(144)
	m.G0 = v414
	switch int32(1039) {
	case 0, 2:
		v424 = v410
		goto L85
	default:
		goto L86
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v372)+4)) = v382
	F_sigemptyset(m, v372+int32(8))
	mBase = m.M
	goto L75
L73:
	;
	*(*int32)(unsafe.Add(mBase, _consts[861])) = v368
	v382 = int32(4729)
	goto L72
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v372)+136)) = int32(268435456)
	v394 = v372 + int32(4)
	goto L79
L77:
	;
	m.G0 = v372 + int32(144)
	goto L71
L79:
	;
	goto L80
L80:
	;
	if v394 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v405 = F___memcpy(m, int32(4610168), v394, int32(140))
	mBase = m.M
	goto L83
L82:
	;
	goto L83
L83:
	;
	goto L77
L84:
	;
	v452 = int32(0)
	v454 = m.G0
	v456 = v454 - int32(144)
	m.G0 = v456
	switch int32(2) {
	case 0, 2:
		v466 = v452
		goto L98
	default:
		goto L99
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414)+4)) = v424
	F_sigemptyset(m, v414+int32(8))
	mBase = m.M
	goto L88
L86:
	;
	*(*int32)(unsafe.Add(mBase, _consts[630])) = v410
	v424 = int32(4729)
	goto L85
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414)+136)) = int32(268435456)
	v436 = v414 + int32(4)
	goto L92
L90:
	;
	m.G0 = v414 + int32(144)
	goto L84
L92:
	;
	goto L93
L93:
	;
	if v436 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v447 = F___memcpy(m, int32(4610448), v436, int32(140))
	mBase = m.M
	goto L96
L95:
	;
	goto L96
L96:
	;
	goto L90
L97:
	;
	goto L9
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v466
	F_sigemptyset(m, v456+int32(8))
	mBase = m.M
	goto L100
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[862])) = v452
	v466 = int32(4729)
	goto L98
L100:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+136)) = int32(268435457)
	v478 = v456 + int32(4)
	goto L105
L103:
	;
	m.G0 = v456 + int32(144)
	goto L97
L105:
	;
	goto L106
L106:
	;
	if v478 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v489 = F___memcpy(m, int32(4611148), v478, int32(140))
	mBase = m.M
	goto L109
L108:
	;
	goto L109
L109:
	;
	goto L103
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v538 = int32(914)
	v540 = m.G0
	v542 = v540 - int32(144)
	m.G0 = v542
	switch int32(916) {
	case 0, 2:
		v552 = v538
		goto L124
	default:
		goto L125
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+4)) = v509
	F_sigemptyset(m, v499+int32(8))
	mBase = m.M
	goto L114
L112:
	;
	*(*int32)(unsafe.Add(mBase, _consts[628])) = v495
	v509 = int32(4729)
	goto L111
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v499)+136)) = int32(268435456)
	v521 = v499 + int32(4)
	goto L118
L116:
	;
	m.G0 = v499 + int32(144)
	goto L110
L118:
	;
	goto L119
L119:
	;
	if v521 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v532 = F___memcpy(m, int32(4608908), v521, int32(140))
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L116
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v581 = int32(295)
	v583 = m.G0
	v585 = v583 - int32(144)
	m.G0 = v585
	switch int32(297) {
	case 0, 2:
		v595 = v581
		goto L137
	default:
		goto L138
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+4)) = v552
	F_sigemptyset(m, v542+int32(8))
	mBase = m.M
	goto L127
L125:
	;
	*(*int32)(unsafe.Add(mBase, _consts[853])) = v538
	v552 = int32(4729)
	goto L124
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+136)) = int32(268435456)
	v564 = v542 + int32(4)
	goto L131
L129:
	;
	m.G0 = v542 + int32(144)
	goto L123
L131:
	;
	goto L132
L132:
	;
	if v564 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v575 = F___memcpy(m, int32(4609048), v564, int32(140))
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L129
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v627 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	if v627 != 0 {
		goto L149
	} else {
		goto L150
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585)+4)) = v595
	F_sigemptyset(m, v585+int32(8))
	mBase = m.M
	goto L140
L138:
	;
	*(*int32)(unsafe.Add(mBase, _consts[629])) = v581
	v595 = int32(4729)
	goto L137
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585)+136)) = int32(268435456)
	v607 = v585 + int32(4)
	goto L144
L142:
	;
	m.G0 = v585 + int32(144)
	goto L136
L144:
	;
	goto L145
L145:
	;
	if v607 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v618 = F___memcpy(m, int32(4610868), v607, int32(140))
	mBase = m.M
	goto L148
L147:
	;
	goto L148
L148:
	;
	goto L142
L149:
	;
	v628 = int32(1158)
	goto L151
L150:
	;
	v628 = int32(295)
	goto L151
L151:
	;
	v630 = m.G0
	v632 = v630 - int32(144)
	m.G0 = v632
	switch v628 + int32(2) {
	case 0, 2:
		v642 = v628
		goto L153
	default:
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v670 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v670
	*(*int32)(unsafe.Add(mBase, _consts[525])) = v670
	v680 = v670
	goto L166
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+4)) = v642
	F_sigemptyset(m, v632+int32(8))
	mBase = m.M
	goto L156
L154:
	;
	*(*int32)(unsafe.Add(mBase, _consts[863])) = v628
	v642 = int32(4729)
	goto L153
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632)+136)) = int32(268435456)
	v654 = v632 + int32(4)
	goto L160
L158:
	;
	m.G0 = v632 + int32(144)
	goto L152
L160:
	;
	goto L161
L161:
	;
	if v654 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v665 = F___memcpy(m, int32(4609188), v654, int32(140))
	mBase = m.M
	goto L164
L163:
	;
	goto L164
L164:
	;
	goto L158
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v818 = int32(-2)
	v820 = m.G0
	v822 = v820 - int32(144)
	m.G0 = v822
	switch int32(0) {
	case 0, 2:
		v832 = v818
		goto L172
	default:
		goto L173
	}
L166:
	;
	v682 = int32(40)
	v683 = v680 * v682
	v686 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[526]))) = uint8(v686)
	*(*int32)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[854]))) = v680
	v693 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[855]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[856]))) = v686
	*(*int64)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[857]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[858]))) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v683)+uint32(_consts[527]))) = uint8(v686)
	v712 = v680 | int32(1)
	v714 = v712 * v682
	*(*uint8)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[526]))) = uint8(v686)
	*(*int32)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[854]))) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[855]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[856]))) = v686
	*(*int64)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[857]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[858]))) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v714)+uint32(_consts[527]))) = uint8(v686)
	v743 = v680 | int32(2)
	v745 = v743 * v682
	*(*uint8)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[526]))) = uint8(v686)
	*(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[854]))) = v743
	*(*int64)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[855]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[856]))) = v686
	*(*int64)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[857]))) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[858]))) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v745)+uint32(_consts[527]))) = uint8(v686)
	if base.B2i32(v680 == int32(20)) == v686 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v811 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[859])) = uint8(v811)
	F_pqsignal_be(m, int32(14), int32(1784))
	mBase = m.M
	goto L165
L168:
	;
	v778 = v680 | int32(3)
	v780 = v778 * int32(40)
	v783 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[526]))) = uint8(v783)
	*(*int32)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[854]))) = v778
	v790 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[855]))) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[856]))) = v783
	*(*int64)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[857]))) = v790
	*(*int32)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[858]))) = v783
	*(*uint8)(unsafe.Add(mBase, uint32(v780)+uint32(_consts[527]))) = uint8(v783)
	v680 = v680 + int32(4)
	goto L166
L169:
	;
	goto L170
L170:
	;
	goto L167
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v861 = int32(916)
	v863 = m.G0
	v865 = v863 - int32(144)
	m.G0 = v865
	switch int32(918) {
	case 0, 2:
		v875 = v861
		goto L185
	default:
		goto L186
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+4)) = v832
	F_sigemptyset(m, v822+int32(8))
	mBase = m.M
	goto L175
L173:
	;
	*(*int32)(unsafe.Add(mBase, _consts[860])) = v818
	v832 = int32(4729)
	goto L172
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v822)+136)) = int32(268435456)
	v844 = v822 + int32(4)
	goto L179
L177:
	;
	m.G0 = v822 + int32(144)
	goto L171
L179:
	;
	goto L180
L180:
	;
	if v844 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v855 = F___memcpy(m, int32(4610588), v844, int32(140))
	mBase = m.M
	goto L183
L182:
	;
	goto L183
L183:
	;
	goto L177
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v904 = int32(-2)
	v906 = m.G0
	v908 = v906 - int32(144)
	m.G0 = v908
	switch int32(0) {
	case 0, 2:
		v918 = v904
		goto L198
	default:
		goto L199
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+4)) = v875
	F_sigemptyset(m, v865+int32(8))
	mBase = m.M
	goto L188
L186:
	;
	*(*int32)(unsafe.Add(mBase, _consts[861])) = v861
	v875 = int32(4729)
	goto L185
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865)+136)) = int32(268435456)
	v887 = v865 + int32(4)
	goto L192
L190:
	;
	m.G0 = v865 + int32(144)
	goto L184
L192:
	;
	goto L193
L193:
	;
	if v887 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v898 = F___memcpy(m, int32(4610168), v887, int32(140))
	mBase = m.M
	goto L196
L195:
	;
	goto L196
L196:
	;
	goto L190
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v947 = int32(918)
	v949 = m.G0
	v951 = v949 - int32(144)
	m.G0 = v951
	switch int32(920) {
	case 0, 2:
		v961 = v947
		goto L211
	default:
		goto L212
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v908)+4)) = v918
	F_sigemptyset(m, v908+int32(8))
	mBase = m.M
	goto L201
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[630])) = v904
	v918 = int32(4729)
	goto L198
L201:
	;
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v908)+136)) = int32(268435456)
	v930 = v908 + int32(4)
	goto L205
L203:
	;
	m.G0 = v908 + int32(144)
	goto L197
L205:
	;
	goto L206
L206:
	;
	if v930 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v941 = F___memcpy(m, int32(4610448), v930, int32(140))
	mBase = m.M
	goto L209
L208:
	;
	goto L209
L209:
	;
	goto L203
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v990 = int32(0)
	v992 = m.G0
	v994 = v992 - int32(144)
	m.G0 = v994
	switch int32(2) {
	case 0, 2:
		v1004 = v990
		goto L224
	default:
		goto L225
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v951)+4)) = v961
	F_sigemptyset(m, v951+int32(8))
	mBase = m.M
	goto L214
L212:
	;
	*(*int32)(unsafe.Add(mBase, _consts[864])) = v947
	v961 = int32(4729)
	goto L211
L214:
	;
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v951)+136)) = int32(268435456)
	v973 = v951 + int32(4)
	goto L218
L216:
	;
	m.G0 = v951 + int32(144)
	goto L210
L218:
	;
	goto L219
L219:
	;
	if v973 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v984 = F___memcpy(m, int32(4609888), v973, int32(140))
	mBase = m.M
	goto L222
L221:
	;
	goto L222
L222:
	;
	goto L216
L223:
	;
	goto L9
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v994)+4)) = v1004
	F_sigemptyset(m, v994+int32(8))
	mBase = m.M
	goto L226
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[862])) = v990
	v1004 = int32(4729)
	goto L224
L226:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v994)+136)) = int32(268435457)
	v1016 = v994 + int32(4)
	goto L231
L229:
	;
	m.G0 = v994 + int32(144)
	goto L223
L231:
	;
	goto L232
L232:
	;
	if v1016 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1027 = F___memcpy(m, int32(4611148), v1016, int32(140))
	mBase = m.M
	goto L235
L234:
	;
	goto L235
L235:
	;
	goto L229
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_sigprocmask(m, int32(4351080), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L237
	}
L237:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v1040 == int32(2) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, _consts[471]))
	if v1045 != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1132 = int32(0)
	v1135 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	F_InitPostgres(m, l0, v1132, l1, v1132, v1135^int32(1), v1132)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L267
	}
L241:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+8))
	if base.Ui32(int32(196609)) < base.Ui32(v1048) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v1052 = int32(32)
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1055 = int32(0)
	v1059 = m.G0
	v1061 = v1059 - int32(16)
	m.G0 = v1061
	*(*int32)(unsafe.Add(mBase, uint32(v1061))) = v1055
	v1067 = F_open(m, int32(275211), v1055, v1061)
	mBase = m.M
	if v1067 != int32(-1) {
		goto L248
	} else {
		goto L249
	}
L244:
	;
	v1051 = int32(32)
	goto L246
L245:
	;
	v1051 = int32(4)
	goto L246
L246:
	;
	v1052 = v1051
	goto L243
L247:
	;
	if v1100 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L248:
	;
	v1070 = int32(1)
	if v1052 == int32(0) {
		v1093 = v1070
		goto L251
	} else {
		goto L252
	}
L249:
	;
	v1100 = v1055
	goto L250
L250:
	;
	m.G0 = v1061 + int32(16)
	goto L247
L251:
	;
	v1095 = F_close(m, v1067)
	mBase = m.M
	v1100 = v1093
	goto L250
L252:
	;
	v1073 = int32(4438640)
	v1074 = v1052
	goto L253
L253:
	;
	v1079 = F_read(m, v1067, v1073, v1074)
	mBase = m.M
	if v1079 <= int32(0) {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1093 = v1070
	goto L251
L255:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v1083 == int32(27) {
		goto L253
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1088 = v1074 - v1079
	if v1088 != 0 {
		v1073 = v1073 + v1079
		v1074 = v1088
		goto L253
	} else {
		goto L259
	}
L258:
	;
	v1093 = int32(0)
	goto L251
L259:
	;
	goto L254
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, _consts[865])) = v1052
	goto L240
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_errcode(m, int32(2600))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_errmsg(m, int32(20788), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_errfinish(m, int32(472649), int32(5124), int32(265807))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, _consts[866]))
	if v1142 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_MemoryContextDelete(m, v1142)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, _consts[71])) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_BeginReportingGUCOptions(m)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L272
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, _consts[866])) = int32(0)
	goto L270
L272:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	if v1156 != int32(1) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_pgstat_report_connect(m)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L277
	}
L274:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[867])))
	if v1160 != int32(1) {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_on_proc_exit(m, int32(1159))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L276
	}
L276:
	;
	goto L273
L277:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	if v1173 == int32(1) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1177 = int32(0)
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v1181 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	goto L280
L280:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v1350 == int32(2) {
		goto L306
	} else {
		goto L307
	}
L281:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[683])) = uint8(v1191)
	v1194 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	if v1194 <= int32(0) {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+316))
	v1189 = base.B2i32(v1187 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v1189)
	v1191 = v1189
	goto L284
L283:
	;
	v1191 = v1177
	goto L284
L284:
	;
	goto L281
L285:
	;
	F_on_shmem_exit(m, int32(1029), int32(0))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L297
	}
L286:
	;
	v1203 = v1177
	goto L287
L287:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v1212 = v1209 + v1203*int32(96)
	v1214 = v1212 + int32(164)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	*(*int32)(unsafe.Add(mBase, uint32(v1214))) = int32(1)
	if v1215 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L285
L289:
	;
	F_s_lock(m, v1214, int32(474026), int32(2958), int32(82017))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1224 = v1212 + int32(88)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1224)))
	if v1225 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L291
L293:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	v1230 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+24)) = v1230
	v1232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1224)+16)) = uint8(v1232)
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+8)) = v1230
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+4)) = v1232
	*(*int32)(unsafe.Add(mBase, uint32(v1224))) = v1229
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+32)) = v1230
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+40)) = v1230
	v1243 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+48)) = v1243
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+56)) = v1243
	*(*int64)(unsafe.Add(mBase, uint32(v1224-int32(-64)))) = v1243
	*(*int64)(unsafe.Add(mBase, uint32(v1224)+80)) = v1230
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+72)) = v1232
	v1256 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+76)) = v1232
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+88)) = base.B2i32(v1256 != v1232)
	*(*int32)(unsafe.Add(mBase, _consts[684])) = v1224
	goto L285
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1214))) = int32(0)
	v1267 = v1203 + int32(1)
	v1269 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	if v1267 < v1269 {
		v1203 = v1267
		goto L287
	} else {
		goto L296
	}
L296:
	;
	goto L288
L297:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, _consts[868]))
	v1291 = *(*int32)(unsafe.Add(mBase, _consts[869]))
	*(*int32)(unsafe.Add(mBase, uint32(v1289+v1291<<(uint(int32(2))%32))+44)) = int32(3)
	F_SendPostmasterSignal(m, int32(8))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1301 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1309 = F_LWLockAcquire(m, v1305+int32(512), int32(0))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	v1335 = F_MemoryContextAllocZero(m, v1333, int32(131192))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L305
	}
L303:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1312)+124)))
	v1315 = v1313 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1312)+124)) = uint8(v1315)
	v1318 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1318)+12))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v1319+v1320))) = uint8(v1315)
	v1324 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1324+int32(512))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	*(*int32)(unsafe.Add(mBase, _consts[703])) = v1335
	goto L280
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_pq_beginmessage(m, v40, int32(75))
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L309
	}
L307:
	;
	v1398 = v1350
	goto L308
L308:
	;
	if v1398 == int32(1) {
		goto L313
	} else {
		goto L314
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1359 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	F_enlargeStringInfo(m, v40, int32(4))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L310
	}
L310:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v1366 = int32(24)
	v1368 = int32(65280)
	v1370 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1363+v1364))) = v1359<<(uint(v1366)%32) | v1359&v1368<<(uint(v1370)%32) | (int32(base.Ui32(v1359)>>(uint(v1370)%32))&v1368 | int32(base.Ui32(v1359)>>(uint(v1366)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v1363 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1388 = *(*int32)(unsafe.Add(mBase, _consts[865]))
	F_pq_sendbytes(m, v40, int32(4438640), v1388)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_pq_endmessage(m, v40)
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L312
	}
L312:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	v1398 = v1395
	goto L308
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(529566)
	F_pg_printf(m, int32(707751), v14)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1410 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	v1415 = F_AllocSetContextCreateInternal(m, v1410, int32(60039), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L317
	}
L316:
	;
	goto L315
L317:
	;
	*(*int32)(unsafe.Add(mBase, _consts[871])) = v1415
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1421 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	v1426 = F_AllocSetContextCreateInternal(m, v1421, int32(59717), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1426
	*(*int32)(unsafe.Add(mBase, _consts[872])) = v1426
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	F_initStringInfo(m, int32(4367480))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L319
	}
L319:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1437
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v31
	v1440 = int32(0)
	v1442 = m.G0
	v1444 = v1442 - int32(80)
	m.G0 = v1444
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	if v1447 != int32(1) {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	goto L388
L321:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L385
	}
L322:
	;
	m.G0 = v1444 + int32(80)
	goto L320
L323:
	;
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, _consts[873])))
	if v1451 != int32(1) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1455 == int32(0) {
		goto L322
	} else {
		goto L325
	}
L325:
	;
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, _consts[874])))
	if v1459 != int32(1) {
		goto L322
	} else {
		goto L326
	}
L326:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L327
	}
L327:
	;
	v1465 = F_EventCacheLookup(m, int32(4))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L330
	}
L328:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L384
	}
L329:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v1549 = m.G0
	v1551 = v1549 - int32(32)
	m.G0 = v1551
	v1553 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v1551)+30)) = uint16(v1553)
	v1555 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1551)+28)) = uint16(v1555)
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+24)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+20)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v1551)+16)) = v1555
	v1570 = F_LockAcquireExtended(m, v1551+int32(16), int32(8), v1555, int32(1), v1551+int32(12), v1555)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L357
	}
L330:
	;
	if v1465 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+4))
	if v1469 <= int32(0) {
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1478 = v1440
	v1481 = v1440
	goto L333
L333:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+12))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1483+v1478<<(uint(int32(2))%32))))
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1487)+4)))
	v1490 = *(*int32)(unsafe.Add(mBase, _consts[384]))
	if v1490 == int32(1) {
		goto L337
	} else {
		goto L338
	}
L334:
	;
	if v1511 == int32(0) {
		goto L329
	} else {
		goto L349
	}
L335:
	;
	v1513 = v1478 + int32(1)
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1465)+4))
	if v1513 < v1514 {
		v1478 = v1513
		v1481 = v1511
		goto L333
	} else {
		goto L348
	}
L336:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+8))
	if v1501 != 0 {
		goto L342
	} else {
		goto L343
	}
L337:
	;
	if v1488&int32(255) != int32(79) {
		goto L336
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	if v1488&int32(255) == int32(82) {
		v1511 = v1481
		goto L335
	} else {
		goto L341
	}
L340:
	;
	v1511 = v1481
	goto L335
L341:
	;
	goto L336
L342:
	;
	v1503 = F_bms_is_member(m, int32(162), v1501)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1487)))
	v1508 = F_lappend_oid(m, v1481, v1507)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L347
	}
L345:
	;
	if v1503 == int32(0) {
		v1511 = v1481
		goto L335
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	v1511 = v1508
	goto L335
L348:
	;
	goto L334
L349:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1444)+24)) = int64(695784701952)
	*(*int32)(unsafe.Add(mBase, uint32(v1444)+20)) = int32(263885)
	*(*int32)(unsafe.Add(mBase, uint32(v1444)+16)) = int32(441)
	v1524 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L350
	}
L350:
	;
	F_PushActiveSnapshot(m, v1524)
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L351
	}
L351:
	;
	F_EventTriggerInvoke(m, v1511, v1444+int32(16))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L352
	}
L352:
	;
	F_list_free(m, v1511)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L353
	}
L353:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L354
	}
L354:
	;
	goto L328
L355:
	;
	m.G0 = v1551 + int32(32)
	if v1570 == int32(0) {
		goto L328
	} else {
		goto L360
	}
L356:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L358
	}
L357:
	;
	switch v1570 {
	case 0, 3:
		goto L355
	default:
		goto L356
	}
L358:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+12))
	v1575 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1574)+53)) = uint8(v1575)
	goto L359
L359:
	;
	goto L355
L360:
	;
	v1583 = F_EventCacheLookup(m, int32(4))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L363
	}
L361:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1444)+24)) = int64(695784701952)
	*(*int32)(unsafe.Add(mBase, uint32(v1444)+20)) = int32(263885)
	*(*int32)(unsafe.Add(mBase, uint32(v1444)+16)) = int32(441)
	F_list_free(m, v1609)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L383
	}
L362:
	;
	v1628 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L371
	}
L363:
	;
	if v1583 == int32(0) {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+4))
	if v1587 <= int32(0) {
		goto L362
	} else {
		goto L365
	}
L365:
	;
	v1590 = int32(0)
	v1598 = v1590
	v1600 = v1590
	goto L366
L366:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+12))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1603+v1598<<(uint(int32(2))%32))))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1607)))
	v1609 = F_lappend_oid(m, v1600, v1608)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L368
	}
L367:
	;
	if v1609 != 0 {
		goto L361
	} else {
		goto L370
	}
L368:
	;
	v1612 = v1598 + int32(1)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+4))
	if v1612 < v1613 {
		v1598 = v1612
		v1600 = v1609
		goto L366
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	goto L362
L371:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	F_ScanKeyInit(m, v1444+int32(16), int32(1), int32(3), int32(184), v1636)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L372
	}
L372:
	;
	F_systable_inplace_update_begin(m, v1628, int32(2672), v1444+int32(16), v1444+int32(76), v1444+int32(72))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L373
	}
L373:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+76))
	if v1648 == int32(0) {
		goto L321
	} else {
		goto L374
	}
L374:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+16))
	v1652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651)+22)))
	v1653 = v1651 + v1652
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653)+79)))
	if v1654 == int32(1) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	F_sequence_close(m, v1628, int32(3))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L381
	}
L376:
	;
	v1657 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1653)+79)) = uint8(v1657)
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+72))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+76))
	F_systable_inplace_update_finish(m, v1659, v1660)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+72))
	F_systable_inplace_update_cancel(m, v1663)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L380
	}
L379:
	;
	goto L375
L380:
	;
	goto L375
L381:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1444)+76))
	F_pfree(m, v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L382
	}
L382:
	;
	goto L328
L383:
	;
	goto L328
L384:
	;
	goto L322
L385:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1444))) = v1712
	F_errmsg_internal(m, int32(47070), v1444)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(473994), int32(970), int32(263891))
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		v1774 = v40
		goto L5
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1722 = int32(4527952)
	*(*int32)(unsafe.Add(mBase, _consts[875])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[876])) = v14 + int32(8)
	goto L391
L389:
	;
	v1732 = v31
	v1736 = v40
	v1740 = int32(0)
	goto L8
L391:
	;
	goto L389
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1732
	F_PostgresMainLongJmp(m)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		v1774 = v1736
		goto L5
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	*(*int32)(unsafe.Add(mBase, _consts[261])) = int32(4527952)
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, _consts[877])))
	if v1748 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L395:
	;
	goto L394
L396:
	;
	v1751 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1732))) = uint8(v1751)
	goto L398
L397:
	;
	goto L398
L398:
	;
	goto L399
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v1732
	F_PostgresMainLoopOnce(m)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		v1774 = v1736
		goto L5
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	v1783 = int32(v1779)
	m.G0 = v1774
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+4))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1783)))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1786)))
	if v14+int32(8) == v1790 {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	m.ExcPending = 1
	goto L411
L404:
	;
	if v1793 != 0 {
		goto L408
	} else {
		goto L409
	}
L405:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1786)+4))
	v1793 = v1792
	goto L407
L406:
	;
	v1793 = int32(0)
	goto L407
L407:
	;
	goto L404
L408:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v19 = v1785
	v20 = v1794
	v22 = v1793
	v24 = v1774
	goto L1
L409:
	;
	goto L410
L410:
	;
	F___wasm_longjmp(m, v1786, v1785)
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	return
L412:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PrepareRedoAdd(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v28 int64
	_ = v28
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(1120)
	m.G0 = v14
	v17 = base.B2i32(v2 == int64(0))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L58
	}
L3:
	;
	m.G0 = v14 + int32(1120)
	return
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99 == int32(0) {
		goto L2
	} else {
		goto L28
	}
L7:
	;
	return
L8:
	;
	if base.Ui32(v20) <= base.Ui32(int32(2)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = base.I64_extend_i32_u(v20)
	goto L11
L10:
	;
	v28 = int64(base.Ui64(v21) >> (uint(int64(32)) % 64))
	if base.Ui32(base.I32_wrap_i64(v21)) < base.Ui32(v20) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+84)) = uint32(v40)
	v43 = int64(base.Ui64(v40) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+80)) = uint32(v43)
	v51 = F_pg_snprintf(m, v14+int32(96), int32(1024), int32(487943), v14+int32(80))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L15
	}
L12:
	;
	v35 = (v28 - int64(1)) & int64(4294967295)
	goto L14
L13:
	;
	v35 = v28
	goto L14
L14:
	;
	v40 = base.I64_extend_i32_u(v20) | v35<<(uint(int64(32))%64)
	goto L11
L15:
	;
	v55 = int32(0)
	v56 = F_access(m, v14+int32(96), v55)
	mBase = m.M
	if v56 == v55 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[154])))
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v91 != int32(44) {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v63 = int32(21)
	goto L21
L20:
	;
	v63 = int32(19)
	goto L21
L21:
	;
	v65 = F_errstart(m, v63, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if v65 == int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v69
	F_errmsg(m, int32(42745), v14+int32(48))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+36)) = uint32(v2)
	v78 = int64(base.Ui64(v2) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v14)+32)) = uint32(v78)
	F_errdetail(m, int32(584081), v14+int32(32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(477000), int32(2516), int32(443577))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	goto L6
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v99)+16)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+32)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+46)) = uint8(v111)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+45)) = uint8(v17)
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+44)) = uint8(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+40)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v110
	v120 = v99 + int32(47)
	v122 = l0 + int32(72)
	if (v122^v120)&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v197 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98+v197<<(uint(int32(2))%32))+8)) = v99
	if l3 != 0 {
		goto L50
	} else {
		goto L51
	}
L30:
	;
	goto L29
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v176)
	if v176&int32(255) == int32(0) {
		goto L30
	} else {
		goto L46
	}
L32:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v175 = v122
	v176 = v128
	v177 = v120
	goto L31
L33:
	;
	goto L34
L34:
	;
	if v122&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = v122
	v134 = v120
	goto L38
L36:
	;
	v146 = v122
	v148 = v120
	goto L37
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v153 = int32(-2139062144)
	if (int32(16843008)-v150|v150)&v153 != v153 {
		v175 = v146
		v176 = v150
		v177 = v148
		goto L31
	} else {
		goto L42
	}
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v135)
	if v135 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L39:
	;
	v146 = v142
	v148 = v140
	goto L37
L40:
	;
	v139 = int32(1)
	v140 = v134 + v139
	v142 = v132 + v139
	if v142&int32(3) != 0 {
		v132 = v142
		v134 = v140
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v158 = v146
	v159 = v150
	v160 = v148
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v159
	v162 = int32(4)
	v163 = v160 + v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v166 = v158 + v162
	v170 = int32(-2139062144)
	if (v164|(int32(16843008)-v164))&v170 == v170 {
		v158 = v166
		v159 = v164
		v160 = v163
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v175 = v166
	v176 = v164
	v177 = v163
	goto L31
L45:
	;
	goto L44
L46:
	;
	v184 = v175
	v186 = v177
	goto L47
L47:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)) = uint8(v187)
	v189 = int32(1)
	if v187 != 0 {
		v184 = v184 + v189
		v186 = v186 + v189
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L30
L49:
	;
	goto L48
L50:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v206 = int32(0)
	F_replorigin_advance(m, l3, v205, l2, v206, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v212 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	if v212 == int32(0) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v216
	F_errmsg_internal(m, int32(42642), v14+int32(16))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(477000), int32(2558), int32(443577))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L3
L58:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(439555), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v249
	F_errhint(m, int32(622390), v14)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(477000), int32(2532), int32(443577))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(96)
	F_errmsg(m, int32(284451), v14-int32(-64))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(477000), int32(2523), int32(443577))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PrescanPreparedTransactions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v19 = F_LWLockAcquire(m, v15+int32(2304), v3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v110+int32(2304))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L4:
	;
	v101 = v3
	v103 = v3
	v104 = v13
	goto L3
L5:
	;
	goto L6
L6:
	;
	v30 = v3
	v31 = v24
	v32 = v3
	v33 = v13
	v34 = v3
	v35 = v3
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v34<<(uint(int32(2))%32))+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v41)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+45)))
	v47 = F_ProcessTwoPhaseBuffer(m, v42, v43, v44, int32(0), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v101 = v88
	v103 = v89
	v104 = v90
	goto L3
L9:
	;
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v33))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v88 = v30
	v89 = v32
	v90 = v33
	v91 = v35
	goto L12
L12:
	;
	v94 = v34 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v94 < v97 {
		v30 = v88
		v31 = v96
		v32 = v89
		v33 = v90
		v34 = v94
		v35 = v91
		goto L7
	} else {
		goto L31
	}
L13:
	;
	if l0 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v60 = base.B2i32(base.Ui32(v42) < base.Ui32(v33))
	goto L13
L15:
	;
	goto L16
L16:
	;
	v60 = int32(base.Ui32(v42-v33) >> (uint(int32(31)) % 32))
	goto L13
L17:
	;
	if v30 != v35 {
		v74 = v32
		v75 = v35
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v82 = v30
	v83 = v32
	v84 = v35
	goto L19
L19:
	;
	if v60 != 0 {
		goto L27
	} else {
		goto L28
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74+v30<<(uint(int32(2))%32)))) = v42
	v82 = v30 + int32(1)
	v83 = v74
	v84 = v75
	goto L19
L21:
	;
	if v30 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = F_palloc(m, int32(40))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = F_repalloc(m, v32, v30<<(uint(int32(3))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v74 = v66
	v75 = int32(10)
	goto L20
L26:
	;
	v74 = v72
	v75 = v30 << (uint(int32(1)) % 32)
	goto L20
L27:
	;
	v85 = v42
	goto L29
L28:
	;
	v85 = v33
	goto L29
L29:
	;
	F_pfree(m, v47)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v88 = v82
	v89 = v83
	v90 = v85
	v91 = v84
	goto L12
L31:
	;
	goto L8
L32:
	;
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v103
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v101
	goto L35
L34:
	;
	goto L35
L35:
	;
	return v104
}
func F_ProcessCatchupInterrupt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, _consts[793]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	goto L6
L5:
	;
	goto L3
L6:
	;
	v14 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if v9 != int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[793]))
	if v41 != 0 {
		goto L4
	} else {
		goto L26
	}
L10:
	;
	if v14 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if v14 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	F_errmsg_internal(m, int32(245831), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(475854), int32(193), int32(79570))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L9
L19:
	;
	F_errmsg_internal(m, int32(245791), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	F_errfinish(m, int32(475854), int32(198), int32(79570))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	goto L9
L26:
	;
	goto L5
}
func F_ProcessMainLoopInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _consts[518]))
			if v6 != 0 {
				*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
					if v14 == int32(0) {
						v18 = *(*int32)(unsafe.Add(mBase, _consts[520]))
						if v18 != 0 {
							F_ProcessLogMemoryContextInterrupt(m)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					} else {
						F_proc_exit(m, int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
				if v14 == int32(0) {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[520]))
					if v18 != 0 {
						F_ProcessLogMemoryContextInterrupt(m)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				} else {
					F_proc_exit(m, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[518]))
		if v6 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[518])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
				if v14 == int32(0) {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[520]))
					if v18 != 0 {
						F_ProcessLogMemoryContextInterrupt(m)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				} else {
					F_proc_exit(m, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _consts[519]))
			if v14 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[520]))
				if v18 != 0 {
					F_ProcessLogMemoryContextInterrupt(m)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				F_proc_exit(m, int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ProcessRepliesIfAny(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int64
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v322 int64
	_ = v322
	var v326 int64
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v359 int64
	_ = v359
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int64
	_ = v388
	var v391 int64
	_ = v391
	var v396 int64
	_ = v396
	var v399 int64
	_ = v399
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int64
	_ = v456
	var v468 int64
	_ = v468
	var v483 int32
	_ = v483
	var v490 int64
	_ = v490
	var v495 int64
	_ = v495
	var v505 int64
	_ = v505
	var v510 int64
	_ = v510
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v521 int64
	_ = v521
	var v527 float64
	_ = v527
	var v531 int64
	_ = v531
	var v536 int64
	_ = v536
	var v539 int64
	_ = v539
	var v547 int64
	_ = v547
	var v566 int64
	_ = v566
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v596 int64
	_ = v596
	var v599 int64
	_ = v599
	var v604 int64
	_ = v604
	var v607 int64
	_ = v607
	var v611 int64
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int64
	_ = v629
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v651 int64
	_ = v651
	var v653 int64
	_ = v653
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int64
	_ = v664
	var v676 int64
	_ = v676
	var v691 int32
	_ = v691
	var v698 int64
	_ = v698
	var v703 int64
	_ = v703
	var v713 int64
	_ = v713
	var v718 int64
	_ = v718
	var v722 int32
	_ = v722
	var v723 int64
	_ = v723
	var v729 int64
	_ = v729
	var v735 float64
	_ = v735
	var v739 int64
	_ = v739
	var v744 int64
	_ = v744
	var v747 int64
	_ = v747
	var v755 int64
	_ = v755
	var v774 int64
	_ = v774
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v804 int64
	_ = v804
	var v807 int64
	_ = v807
	var v812 int64
	_ = v812
	var v815 int64
	_ = v815
	var v819 int64
	_ = v819
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int64
	_ = v830
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int64
	_ = v837
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v859 int64
	_ = v859
	var v861 int64
	_ = v861
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int64
	_ = v872
	var v884 int64
	_ = v884
	var v899 int32
	_ = v899
	var v906 int64
	_ = v906
	var v911 int64
	_ = v911
	var v921 int64
	_ = v921
	var v926 int64
	_ = v926
	var v930 int32
	_ = v930
	var v931 int64
	_ = v931
	var v937 int64
	_ = v937
	var v943 float64
	_ = v943
	var v947 int64
	_ = v947
	var v952 int64
	_ = v952
	var v955 int64
	_ = v955
	var v963 int64
	_ = v963
	var v982 int64
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int64
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1006 int32
	_ = v1006
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int64
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1054 int64
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1163 int32
	_ = v1163
	var v1174 int32
	_ = v1174
	var v1183 int64
	_ = v1183
	var v1184 int64
	_ = v1184
	var v1185 int64
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1192 int64
	_ = v1192
	var v1196 int64
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1202 int64
	_ = v1202
	var v1206 int64
	_ = v1206
	var v1207 int64
	_ = v1207
	var v1211 int64
	_ = v1211
	var v1212 int64
	_ = v1212
	var v1216 int64
	_ = v1216
	var v1217 int64
	_ = v1217
	var v1221 int64
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1226 int64
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1252 int64
	_ = v1252
	var v1253 int64
	_ = v1253
	var v1254 int64
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1263 int64
	_ = v1263
	var v1267 int64
	_ = v1267
	var v1268 int64
	_ = v1268
	var v1272 int64
	_ = v1272
	var v1273 int64
	_ = v1273
	var v1277 int64
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1328 int64
	_ = v1328
	var v1331 int64
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int64
	_ = v1341
	var v1344 int64
	_ = v1344
	var v1347 int64
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int64
	_ = v1386
	var v1389 int64
	_ = v1389
	var v1392 int64
	_ = v1392
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1413 int32
	_ = v1413
	var v1415 int64
	_ = v1415
	var v1417 int64
	_ = v1417
	var v1419 int64
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1447 int64
	_ = v1447
	var v1448 int64
	_ = v1448
	var v1449 int64
	_ = v1449
	var v1459 int32
	_ = v1459
	var v1474 int64
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1476 int64
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1543 int32
	_ = v1543
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1574 int64
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1609 int32
	_ = v1609
	var v1610 int64
	_ = v1610
	var v1613 int64
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1638 int32
	_ = v1638
	var v1659 int32
	_ = v1659
	var v1660 int64
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1699 int32
	_ = v1699
	var v1700 int64
	_ = v1700
	var v1703 int64
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1727 int32
	_ = v1727
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1751 int64
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1771 int32
	_ = v1771
	var v1790 int32
	_ = v1790
	var v1791 int64
	_ = v1791
	var v1794 int64
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1820 int32
	_ = v1820
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1852 int64
	_ = v1852
	var v1853 int64
	_ = v1853
	var v1858 int64
	_ = v1858
	var v1864 int64
	_ = v1864
	var v1868 int32
	_ = v1868
	var v1873 int32
	_ = v1873
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1946 int32
	_ = v1946
	var v1947 int64
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2020 int64
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2045 int32
	_ = v2045
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2111 int32
	_ = v2111
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2132 int64
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2157 int32
	_ = v2157
	var v2165 int64
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2222 int32
	_ = v2222
	var v2229 int32
	_ = v2229
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2244 int32
	_ = v2244
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2285 int32
	_ = v2285
	var v2303 int32
	_ = v2303
	var v2313 int32
	_ = v2313
	var v2338 int64
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2346 int32
	_ = v2346
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v35 = m.G0
	v36 = int32(16)
	v37 = v35 - v36
	m.G0 = v37
	F___gettimeofday(m, v37)
	mBase = m.M
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+8)))
	m.G0 = v37 + v36
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, _consts[689])) = v41 + v40*int64(1000000) - int64(946684800000000)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[690])))
	if v52 != 0 {
		v2346 = v29
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v2346 + int32(96)
	return
L3:
	;
	v64 = v29
	v70 = int32(0)
	v76 = v29 + int32(80)
	v77 = v29 + int32(60)
	v78 = v29 + int32(56)
	v79 = v29 + int32(52)
	goto L5
L4:
	;
	v2338 = *(*int64)(unsafe.Add(mBase, _consts[689]))
	*(*int64)(unsafe.Add(mBase, _consts[691])) = v2338
	v2341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[692])) = uint8(v2341)
	v2346 = v2313
	goto L2
L5:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v2285&int32(1) == int32(0) {
		v2346 = v64
		goto L2
	} else {
		goto L481
	}
L7:
	;
	return
L8:
	;
	v90 = v64 + int32(95)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[693]))
	v94 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	if v92 < v94 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v161 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L10:
	;
	v97 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[693])) = v92 + v97
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[695]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v102)
	v161 = v97
	goto L9
L11:
	;
	goto L12
L12:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[471]))
	if v106 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)) = uint8(v107)
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v109
	v114 = *(*int32)(unsafe.Add(mBase, _consts[471]))
	v116 = F_secure_read(m, v114, v90, v107)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L31
	}
L16:
	;
	v161 = v142
	goto L9
L17:
	;
	if v116 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	switch v121 {
	case 0:
		goto L21
	default:
		goto L22
	case 6, 27:
		v142 = v109
		goto L16
	}
L19:
	;
	goto L20
L20:
	;
	if v116 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v161 = int32(-1)
	goto L9
L22:
	;
	v124 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if v124 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(279255), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(475487), int32(1043), int32(377916))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v141 = v116
	goto L30
L29:
	;
	v141 = int32(-1)
	goto L30
L30:
	;
	v142 = v141
	goto L16
L31:
	;
	F_errcode(m, int32(50332160))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(243544), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(475487), int32(886), int32(320271))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v166 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v161 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	if v166 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L45
	}
L42:
	;
	F_errmsg(m, int32(243507), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(474026), int32(2263), int32(17212))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[472])) = uint8(v186)
	goto L49
L47:
	;
	goto L48
L48:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+95)))
	switch v191 - int32(88) {
	case 0, 11:
		goto L52
	default:
		goto L53
	case 12:
		v212 = int32(1073741822)
		goto L51
	}
L49:
	;
	if v70&int32(1) != 0 {
		v2313 = v64
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v2346 = v64
	goto L2
L51:
	;
	v213 = int32(4354320)
	v214 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, _consts[697])) = v215
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v215
	goto L58
L52:
	;
	v212 = int32(10000)
	goto L51
L53:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+95)))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v201
	F_errmsg(m, int32(689466), v64)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(474026), int32(2287), int32(17212))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
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
	v222 = F_pq_getmessage(m, int32(4354320), v212)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	if v222 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v226 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+95)))
	switch v243 - int32(88) {
	case 0:
		goto L73
	default:
		v2285 = v70
		goto L71
	case 11:
		goto L74
	case 12:
		goto L75
	}
L63:
	;
	if v226 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L70
	}
L67:
	;
	F_errmsg(m, int32(243507), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(474026), int32(2298), int32(17212))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v2303 = int32(*(*uint8)(unsafe.Add(mBase, _consts[690])))
	if v2303 == int32(0) {
		v70 = v2285
		goto L5
	} else {
		goto L480
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1931))) = int32(0)
	v2285 = v1929
	goto L71
L73:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L7
	} else {
		goto L479
	}
L74:
	;
	v2254 = int32(*(*uint8)(unsafe.Add(mBase, _consts[699])))
	if v2254 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L75:
	;
	v247 = F_pq_getmsgbyte(m, int32(4354320))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L79
	}
L76:
	;
	v2234 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L7
	} else {
		goto L467
	}
L77:
	;
	v2020 = F_pq_getmsgint64(m, int32(4354320))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L7
	} else {
		goto L388
	}
L78:
	;
	v253 = F_pq_getmsgint64(m, int32(4354320))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L80
	}
L79:
	;
	v249 = base.I32_extend8_s(v247)
	switch v249 - int32(104) {
	case 0:
		goto L77
	default:
		goto L76
	case 10:
		goto L78
	}
L80:
	;
	v256 = F_pq_getmsgint64(m, int32(4354320))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L7
	} else {
		goto L81
	}
L81:
	;
	v259 = F_pq_getmsgint64(m, int32(4354320))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	v262 = F_pq_getmsgint64(m, int32(4354320))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	v265 = F_pq_getmsgbyte(m, int32(4354320))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	v267 = int32(13)
	goto L87
L85:
	;
	if v301 != 0 {
		goto L99
	} else {
		goto L100
	}
L86:
	;
	goto L85
L87:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	goto L90
L88:
	;
	v286 = int32(0)
	goto L96
L90:
	;
	goto L91
L91:
	;
	goto L93
L93:
	;
	if v274 == int32(15) {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	if v274 <= v267 {
		v301 = int32(1)
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L88
L96:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v290 != int32(2) {
		v301 = v286
		goto L86
	} else {
		goto L97
	}
L97:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, _consts[701])))
	if v294 != 0 {
		v301 = v286
		goto L86
	} else {
		goto L98
	}
L98:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	v301 = int32(0) | base.B2i32(v298 <= v267)
	goto L86
L99:
	;
	v303 = F_timestamptz_to_str(m, v262)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v345 = m.G0
	v346 = int32(16)
	v347 = v345 - v346
	m.G0 = v347
	F___gettimeofday(m, v347)
	mBase = m.M
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v347)))
	v351 = int64(*(*int32)(unsafe.Add(mBase, uint32(v347)+8)))
	m.G0 = v347 + v346
	v359 = v351 + v350*int64(1000000) - int64(946684800000000)
	goto L114
L102:
	;
	v305 = F_pstrdup(m, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v309 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	if v309 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v305
	if v265 != 0 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	F_pfree(m, v305)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L113
	}
L108:
	;
	v314 = int32(638111)
	goto L110
L109:
	;
	v314 = int32(717063)
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v314
	*(*uint32)(unsafe.Add(mBase, uint32(v79))) = uint32(v259)
	v317 = int64(32)
	v318 = int64(base.Ui64(v259) >> (uint(v317) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v64)+48)) = uint32(v318)
	*(*uint32)(unsafe.Add(mBase, uint32(v64)+44)) = uint32(v256)
	v322 = int64(base.Ui64(v256) >> (uint(v317) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v64)+40)) = uint32(v322)
	*(*uint32)(unsafe.Add(mBase, uint32(v64)+36)) = uint32(v253)
	v326 = int64(base.Ui64(v253) >> (uint(v317) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v64)+32)) = uint32(v326)
	F_errmsg_internal(m, int32(184321), v64+int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(474026), int32(2451), int32(387459))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	goto L107
L113:
	;
	goto L101
L114:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[704])))
	if v379 != int32(-1) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[705])))
	if v587 != int32(-1) {
		goto L149
	} else {
		goto L150
	}
L116:
	;
	if v411 == v413 {
		v468 = v414
		goto L127
	} else {
		goto L128
	}
L117:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[706])))
	v411 = v379
	v413 = v382
	v414 = int64(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[707])))
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[708])))
	if base.Ui64(v253) < base.Ui64(v391) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if v359 < v388 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[707])))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[709]))) = v399
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[708])))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[710]))) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[706])))
	v409 = base.I32_rem_s(v405+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[704]))) = v409
	v411 = v409
	v413 = v405
	v414 = v388
	goto L116
L123:
	;
	v396 = int64(-1)
	goto L125
L124:
	;
	v396 = v359 - v388
	goto L125
L125:
	;
	v566 = v396
	goto L115
L126:
	;
	v495 = int64(-1)
	if v359 < v490 {
		v547 = v495
		goto L134
	} else {
		goto L135
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[709]))) = int64(0)
	v483 = v413
	v490 = v468
	goto L126
L128:
	;
	v417 = v373 + int32(8)
	v420 = v417 + v411<<(uint(int32(4))%32)
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v420)))
	if base.Ui64(v253) < base.Ui64(v421) {
		v483 = v411
		v490 = v414
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v431 = v411
	v433 = v420
	goto L130
L130:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v433)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[709]))) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v433)))
	*(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[710]))) = v445
	v450 = base.I32_rem_s(v431+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[704]))) = v450
	if v450 == v413 {
		v468 = v443
		goto L127
	} else {
		goto L132
	}
L131:
	;
	v483 = v450
	v490 = v443
	goto L126
L132:
	;
	v455 = v417 + v450<<(uint(int32(4))%32)
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v455)))
	if base.Ui64(v456) <= base.Ui64(v253) {
		v431 = v450
		v433 = v455
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v566 = v547
	goto L115
L135:
	;
	if v490 != int64(0) {
		v539 = v490
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v547 = v359 - v539
	goto L134
L137:
	;
	if v483 == v413 {
		v547 = v495
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v505 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[709])))
	if v505 != int64(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[710])))
	if base.Ui64(v253) < base.Ui64(v510) {
		v547 = v495
		goto L134
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v373+v483<<(uint(int32(4))%32))+16))
	v539 = v536
	goto L136
L142:
	;
	v514 = v373 + v483<<(uint(int32(4))%32)
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v514)+16))
	if v515 < v505 {
		v547 = v495
		goto L134
	} else {
		goto L143
	}
L143:
	;
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v514)+8))
	v527 = base.F64_add(base.F64_mul(base.F64_convert_i64_s(v515-v505), base.F64_div(base.F64_convert_i64_u(v253-v510), base.F64_convert_i64_u(v521-v510))), base.F64_convert_i64_s(v505))
	if base.F64_lt(base.F64_abs(v527), float64(9.223372036854776e+18)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v531 = base.I64_trunc_f64_s(v527)
	v539 = v531
	goto L136
L145:
	;
	goto L146
L146:
	;
	v539 = int64(-9223372036854775807 - 1)
	goto L136
L147:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[711])))
	if v795 != int32(-1) {
		goto L181
	} else {
		goto L182
	}
L148:
	;
	if v619 == v621 {
		v676 = v622
		goto L159
	} else {
		goto L160
	}
L149:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[706])))
	v619 = v587
	v621 = v590
	v622 = int64(0)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[712])))
	v599 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[713])))
	if base.Ui64(v256) < base.Ui64(v599) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v359 < v596 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[712])))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[714]))) = v607
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[713])))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[715]))) = v611
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[706])))
	v617 = base.I32_rem_s(v613+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[705]))) = v617
	v619 = v617
	v621 = v613
	v622 = v596
	goto L148
L155:
	;
	v604 = int64(-1)
	goto L157
L156:
	;
	v604 = v359 - v596
	goto L157
L157:
	;
	v774 = v604
	goto L147
L158:
	;
	v703 = int64(-1)
	if v359 < v698 {
		v755 = v703
		goto L166
	} else {
		goto L167
	}
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[714]))) = int64(0)
	v691 = v621
	v698 = v676
	goto L158
L160:
	;
	v625 = v581 + int32(8)
	v628 = v625 + v619<<(uint(int32(4))%32)
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v628)))
	if base.Ui64(v256) < base.Ui64(v629) {
		v691 = v619
		v698 = v622
		goto L158
	} else {
		goto L161
	}
L161:
	;
	v639 = v619
	v641 = v628
	goto L162
L162:
	;
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v641)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[714]))) = v651
	v653 = *(*int64)(unsafe.Add(mBase, uint32(v641)))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[715]))) = v653
	v658 = base.I32_rem_s(v639+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[705]))) = v658
	if v658 == v621 {
		v676 = v651
		goto L159
	} else {
		goto L164
	}
L163:
	;
	v691 = v658
	v698 = v651
	goto L158
L164:
	;
	v663 = v625 + v658<<(uint(int32(4))%32)
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v663)))
	if base.Ui64(v664) <= base.Ui64(v256) {
		v639 = v658
		v641 = v663
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v774 = v755
	goto L147
L167:
	;
	if v698 != int64(0) {
		v747 = v698
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v755 = v359 - v747
	goto L166
L169:
	;
	if v691 == v621 {
		v755 = v703
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v713 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[714])))
	if v713 != int64(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_consts[715])))
	if base.Ui64(v256) < base.Ui64(v718) {
		v755 = v703
		goto L166
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v581+v691<<(uint(int32(4))%32))+16))
	v747 = v744
	goto L168
L174:
	;
	v722 = v581 + v691<<(uint(int32(4))%32)
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v722)+16))
	if v723 < v713 {
		v755 = v703
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v729 = *(*int64)(unsafe.Add(mBase, uint32(v722)+8))
	v735 = base.F64_add(base.F64_mul(base.F64_convert_i64_s(v723-v713), base.F64_div(base.F64_convert_i64_u(v256-v718), base.F64_convert_i64_u(v729-v718))), base.F64_convert_i64_s(v713))
	if base.F64_lt(base.F64_abs(v735), float64(9.223372036854776e+18)) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v739 = base.I64_trunc_f64_s(v735)
	v747 = v739
	goto L168
L177:
	;
	goto L178
L178:
	;
	v747 = int64(-9223372036854775807 - 1)
	goto L168
L179:
	;
	v983 = int32(4354424)
	v984 = int32(*(*uint8)(unsafe.Add(mBase, _consts[716])))
	v987 = *(*int64)(unsafe.Add(mBase, _consts[717]))
	v988 = base.B2i32(v259 == v987)
	*(*uint8)(unsafe.Add(mBase, _consts[716])) = uint8(v988)
	if v265 != 0 {
		goto L211
	} else {
		goto L212
	}
L180:
	;
	if v827 == v829 {
		v884 = v830
		goto L191
	} else {
		goto L192
	}
L181:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[706])))
	v827 = v795
	v829 = v798
	v830 = int64(0)
	goto L180
L182:
	;
	goto L183
L183:
	;
	v804 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[718])))
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[719])))
	if base.Ui64(v259) < base.Ui64(v807) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if v359 < v804 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	v815 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[718])))
	*(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[720]))) = v815
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[719])))
	*(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[721]))) = v819
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[706])))
	v825 = base.I32_rem_s(v821+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[711]))) = v825
	v827 = v825
	v829 = v821
	v830 = v804
	goto L180
L187:
	;
	v812 = int64(-1)
	goto L189
L188:
	;
	v812 = v359 - v804
	goto L189
L189:
	;
	v982 = v812
	goto L179
L190:
	;
	v911 = int64(-1)
	if v359 < v906 {
		v963 = v911
		goto L198
	} else {
		goto L199
	}
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[720]))) = int64(0)
	v899 = v829
	v906 = v884
	goto L190
L192:
	;
	v833 = v789 + int32(8)
	v836 = v833 + v827<<(uint(int32(4))%32)
	v837 = *(*int64)(unsafe.Add(mBase, uint32(v836)))
	if base.Ui64(v259) < base.Ui64(v837) {
		v899 = v827
		v906 = v830
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v847 = v827
	v849 = v836
	goto L194
L194:
	;
	v859 = *(*int64)(unsafe.Add(mBase, uint32(v849)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[720]))) = v859
	v861 = *(*int64)(unsafe.Add(mBase, uint32(v849)))
	*(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[721]))) = v861
	v866 = base.I32_rem_s(v847+int32(1), int32(8192))
	*(*int32)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[711]))) = v866
	if v866 == v829 {
		v884 = v859
		goto L191
	} else {
		goto L196
	}
L195:
	;
	v899 = v866
	v906 = v859
	goto L190
L196:
	;
	v871 = v833 + v866<<(uint(int32(4))%32)
	v872 = *(*int64)(unsafe.Add(mBase, uint32(v871)))
	if base.Ui64(v872) <= base.Ui64(v259) {
		v847 = v866
		v849 = v871
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v982 = v963
	goto L179
L199:
	;
	if v906 != int64(0) {
		v955 = v906
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v963 = v359 - v955
	goto L198
L201:
	;
	if v899 == v829 {
		v963 = v911
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v921 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[720])))
	if v921 != int64(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v926 = *(*int64)(unsafe.Add(mBase, uint32(v789)+uint32(_consts[721])))
	if base.Ui64(v259) < base.Ui64(v926) {
		v963 = v911
		goto L198
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v952 = *(*int64)(unsafe.Add(mBase, uint32(v789+v899<<(uint(int32(4))%32))+16))
	v955 = v952
	goto L200
L206:
	;
	v930 = v789 + v899<<(uint(int32(4))%32)
	v931 = *(*int64)(unsafe.Add(mBase, uint32(v930)+16))
	if v931 < v921 {
		v963 = v911
		goto L198
	} else {
		goto L207
	}
L207:
	;
	v937 = *(*int64)(unsafe.Add(mBase, uint32(v930)+8))
	v943 = base.F64_add(base.F64_mul(base.F64_convert_i64_s(v931-v921), base.F64_div(base.F64_convert_i64_u(v259-v926), base.F64_convert_i64_u(v937-v926))), base.F64_convert_i64_s(v921))
	if base.F64_lt(base.F64_abs(v943), float64(9.223372036854776e+18)) != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v947 = base.I64_trunc_f64_s(v943)
	v955 = v947
	goto L200
L209:
	;
	goto L210
L210:
	;
	v955 = int64(-9223372036854775807 - 1)
	goto L200
L211:
	;
	F_WalSndKeepalive(m, int32(0), int64(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L7
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v994 = v984 & v988
	v996 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v996)+76)) = int32(1)
	if v997 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L213
L215:
	;
	F_s_lock(m, v996+int32(76), int32(474026), int32(2491), int32(387459))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L7
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v996)+40)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v996)+32)) = v256
	*(*int64)(unsafe.Add(mBase, uint32(v996)+24)) = v253
	if (base.B2i32(v566 != int64(-1))|v994)&int32(1) != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L217
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v996)+48)) = v566
	goto L221
L220:
	;
	goto L221
L221:
	;
	if (base.B2i32(v774 != int64(-1))|v994)&int32(1) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v996)+56)) = v774
	goto L224
L223:
	;
	goto L224
L224:
	;
	if (base.B2i32(v982 != int64(-1))|v994)&int32(1) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v996)+64)) = v982
	goto L227
L226:
	;
	goto L227
L227:
	;
	v1028 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v996)+76)) = v1028
	*(*int64)(unsafe.Add(mBase, uint32(v996)+80)) = v262
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, _consts[683])))
	if v1032 == v1028 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1035 = int64(0)
	v1038 = int32(0)
	v1040 = m.G0
	v1042 = v1040 - int32(80)
	m.G0 = v1042
	v1045 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+72))
	if v1046 == v1038 {
		goto L233
	} else {
		goto L234
	}
L229:
	;
	goto L230
L230:
	;
	v1929 = int32(1)
	v1931 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v1931 == int32(0) {
		v2285 = v1929
		goto L71
	} else {
		goto L360
	}
L231:
	;
	m.G0 = v1042 + int32(80)
	goto L230
L232:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v1062 = int32(0)
	v1064 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1068 = F_LWLockAcquire(m, v1064+int32(4096), v1062)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L7
	} else {
		goto L237
	}
L233:
	;
	v1058 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[722])) = uint8(v1058)
	goto L231
L234:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
	if base.Ui32(int32(1)) < base.Ui32(v1049-int32(3)) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v1045)+32))
	if v1054 != int64(0) {
		goto L232
	} else {
		goto L236
	}
L236:
	;
	goto L233
L237:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	if v1071 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	v1572 = int32(0)
	v1574 = *(*int64)(unsafe.Add(mBase, uint32(v1061)+24))
	if base.Ui64(v1476) <= base.Ui64(v1574) {
		v1638 = v1572
		goto L328
	} else {
		goto L329
	}
L239:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1565+int32(4096))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L7
	} else {
		goto L327
	}
L240:
	;
	v1543 = int32(1)
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1077 = F_SyncRepGetCandidateStandbys(m, v1042+int32(76))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L7
	} else {
		goto L243
	}
L243:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+76))
	if v1077 <= int32(0) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	F_pfree(m, v1079)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L7
	} else {
		goto L312
	}
L245:
	;
	v1146 = int32(0)
	v1149 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+4))
	if v1077 < v1150 {
		v1459 = v1146
		v1474 = v1035
		v1475 = v1035
		v1476 = v1035
		v1480 = int32(1)
		goto L244
	} else {
		goto L254
	}
L246:
	;
	v1459 = v1124
	v1474 = v1035
	v1475 = v1035
	v1476 = v1035
	v1480 = int32(1)
	goto L244
L247:
	;
	v1124 = int32(1)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1083 = v1062
	goto L250
L250:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079+v1083*int32(48))+40)))
	if v1112 == int32(1) {
		goto L245
	} else {
		goto L252
	}
L251:
	;
	v1124 = v1115
	goto L246
L252:
	;
	v1115 = int32(1)
	v1117 = v1083 + v1115
	if v1117 != v1077 {
		v1083 = v1117
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1149)+8)))
	if v1152 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1459 = v1432
	v1474 = v1447
	v1475 = v1448
	v1476 = v1449
	v1480 = int32(0)
	goto L244
L256:
	;
	v1155 = int32(1)
	if v1077 == v1155 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	goto L258
L258:
	;
	v1280 = int32(0)
	v1282 = v1077 << (uint(int32(3)) % 32)
	v1283 = F_palloc(m, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L7
	} else {
		goto L294
	}
L259:
	;
	if v1077&v1155 == int32(0) {
		v1432 = v1146
		v1447 = v1252
		v1448 = v1253
		v1449 = v1254
		goto L255
	} else {
		goto L284
	}
L260:
	;
	v1232 = int32(0)
	v1252 = v1035
	v1253 = v1035
	v1254 = v1035
	goto L259
L261:
	;
	goto L262
L262:
	;
	v1163 = int32(0)
	v1174 = v1038
	v1183 = v1035
	v1184 = v1035
	v1185 = v1035
	goto L263
L263:
	;
	v1191 = v1079 + v1163*int32(48)
	v1192 = *(*int64)(unsafe.Add(mBase, uint32(v1191)+24))
	if base.Ui64(v1184-int64(1)) < base.Ui64(v1192) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1232 = v1228
	v1252 = v1216
	v1253 = v1206
	v1254 = v1226
	goto L259
L265:
	;
	v1196 = v1184
	goto L267
L266:
	;
	v1196 = v1192
	goto L267
L267:
	;
	v1201 = v1079 + (v1163|int32(1))*int32(48)
	v1202 = *(*int64)(unsafe.Add(mBase, uint32(v1201)+24))
	if base.Ui64(v1196-int64(1)) < base.Ui64(v1202) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1206 = v1196
	goto L270
L269:
	;
	v1206 = v1202
	goto L270
L270:
	;
	v1207 = *(*int64)(unsafe.Add(mBase, uint32(v1191)+16))
	if base.Ui64(v1183-int64(1)) < base.Ui64(v1207) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1211 = v1183
	goto L273
L272:
	;
	v1211 = v1207
	goto L273
L273:
	;
	v1212 = *(*int64)(unsafe.Add(mBase, uint32(v1201)+16))
	if base.Ui64(v1211-int64(1)) < base.Ui64(v1212) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1216 = v1211
	goto L276
L275:
	;
	v1216 = v1212
	goto L276
L276:
	;
	v1217 = *(*int64)(unsafe.Add(mBase, uint32(v1191)+8))
	if base.Ui64(v1185-int64(1)) < base.Ui64(v1217) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1221 = v1185
	goto L279
L278:
	;
	v1221 = v1217
	goto L279
L279:
	;
	v1222 = *(*int64)(unsafe.Add(mBase, uint32(v1201)+8))
	if base.Ui64(v1221-int64(1)) < base.Ui64(v1222) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1226 = v1221
	goto L282
L281:
	;
	v1226 = v1222
	goto L282
L282:
	;
	v1227 = int32(2)
	v1228 = v1163 + v1227
	v1230 = v1174 + v1227
	if v1230 != v1077&int32(2147483646) {
		v1163 = v1228
		v1174 = v1230
		v1183 = v1216
		v1184 = v1206
		v1185 = v1226
		goto L263
	} else {
		goto L283
	}
L283:
	;
	goto L264
L284:
	;
	v1262 = v1079 + v1232*int32(48)
	v1263 = *(*int64)(unsafe.Add(mBase, uint32(v1262)+24))
	if base.Ui64(v1253-int64(1)) < base.Ui64(v1263) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1267 = v1253
	goto L287
L286:
	;
	v1267 = v1263
	goto L287
L287:
	;
	v1268 = *(*int64)(unsafe.Add(mBase, uint32(v1262)+16))
	if base.Ui64(v1252-int64(1)) < base.Ui64(v1268) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1272 = v1252
	goto L290
L289:
	;
	v1272 = v1268
	goto L290
L290:
	;
	v1273 = *(*int64)(unsafe.Add(mBase, uint32(v1262)+8))
	if base.Ui64(v1254-int64(1)) < base.Ui64(v1273) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1277 = v1254
	goto L293
L292:
	;
	v1277 = v1273
	goto L293
L293:
	;
	v1432 = v1146
	v1447 = v1272
	v1448 = v1267
	v1449 = v1277
	goto L255
L294:
	;
	v1285 = F_palloc(m, v1282)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L7
	} else {
		goto L295
	}
L295:
	;
	v1287 = F_palloc(m, v1282)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L7
	} else {
		goto L296
	}
L296:
	;
	if v1077 != int32(1) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1293 = v1280
	v1297 = v1038
	goto L300
L298:
	;
	v1354 = v1280
	goto L299
L299:
	;
	if v1077&int32(1) != 0 {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1319 = int32(3)
	v1320 = v1293 << (uint(v1319) % 32)
	v1322 = int32(48)
	v1324 = v1079 + v1293*v1322
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v1324)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1283+v1320))) = v1325
	v1328 = *(*int64)(unsafe.Add(mBase, uint32(v1324)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1320+v1285))) = v1328
	v1331 = *(*int64)(unsafe.Add(mBase, uint32(v1324)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1320+v1287))) = v1331
	v1334 = v1293 | int32(1)
	v1336 = v1334 << (uint(v1319) % 32)
	v1340 = v1079 + v1334*v1322
	v1341 = *(*int64)(unsafe.Add(mBase, uint32(v1340)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1283+v1336))) = v1341
	v1344 = *(*int64)(unsafe.Add(mBase, uint32(v1340)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1336+v1285))) = v1344
	v1347 = *(*int64)(unsafe.Add(mBase, uint32(v1340)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1336+v1287))) = v1347
	v1349 = int32(2)
	v1350 = v1293 + v1349
	v1352 = v1297 + v1349
	if v1352 != v1077&int32(2147483646) {
		v1293 = v1350
		v1297 = v1352
		goto L300
	} else {
		goto L302
	}
L301:
	;
	v1354 = v1350
	goto L299
L302:
	;
	goto L301
L303:
	;
	v1381 = v1354 << (uint(int32(3)) % 32)
	v1385 = v1079 + v1354*int32(48)
	v1386 = *(*int64)(unsafe.Add(mBase, uint32(v1385)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1283+v1381))) = v1386
	v1389 = *(*int64)(unsafe.Add(mBase, uint32(v1385)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1381+v1285))) = v1389
	v1392 = *(*int64)(unsafe.Add(mBase, uint32(v1385)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1381+v1287))) = v1392
	goto L305
L304:
	;
	goto L305
L305:
	;
	F_pg_qsort(m, v1283, v1077, int32(8), int32(1026))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L7
	} else {
		goto L306
	}
L306:
	;
	F_pg_qsort(m, v1285, v1077, int32(8), int32(1026))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L7
	} else {
		goto L307
	}
L307:
	;
	F_pg_qsort(m, v1287, v1077, int32(8), int32(1026))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L7
	} else {
		goto L308
	}
L308:
	;
	v1413 = v1150&int32(255)<<(uint(int32(3))%32) - int32(8)
	v1415 = *(*int64)(unsafe.Add(mBase, uint32(v1287+v1413)))
	v1417 = *(*int64)(unsafe.Add(mBase, uint32(v1413+v1285)))
	v1419 = *(*int64)(unsafe.Add(mBase, uint32(v1413+v1283)))
	F_pfree(m, v1283)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L7
	} else {
		goto L309
	}
L309:
	;
	F_pfree(m, v1285)
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L7
	} else {
		goto L310
	}
L310:
	;
	F_pfree(m, v1287)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L7
	} else {
		goto L311
	}
L311:
	;
	v1432 = int32(0)
	v1447 = v1417
	v1448 = v1415
	v1449 = v1419
	goto L255
L312:
	;
	if v1459 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	if v1480|v1459 != int32(1) {
		goto L238
	} else {
		goto L326
	}
L314:
	;
	v1484 = int32(*(*uint8)(unsafe.Add(mBase, _consts[722])))
	if v1484&int32(1) == int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[722])) = uint8(v1490)
	v1493 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1493)+8)))
	v1497 = F_errstart(m, int32(15), v1490)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L7
	} else {
		goto L316
	}
L316:
	;
	if v1494 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	F_errfinish(m, int32(474371), v1529, int32(123514))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L7
	} else {
		goto L325
	}
L318:
	;
	if v1497 == int32(0) {
		goto L313
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	if v1497 == int32(0) {
		goto L313
	} else {
		goto L323
	}
L321:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+72))
	v1508 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+48)) = v1508
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+52)) = v1506
	F_errmsg(m, int32(446994), v1042+int32(48))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L7
	} else {
		goto L322
	}
L322:
	;
	v1529 = int32(529)
	goto L317
L323:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+64)) = v1520
	F_errmsg(m, int32(21983), v1042-int32(-64))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L7
	} else {
		goto L324
	}
L324:
	;
	v1529 = int32(533)
	goto L317
L325:
	;
	goto L313
L326:
	;
	v1543 = v1459
	goto L239
L327:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[722])) = uint8(v1543)
	goto L231
L328:
	;
	v1659 = v1061 + int32(32)
	v1660 = *(*int64)(unsafe.Add(mBase, uint32(v1659)))
	if base.Ui64(v1474) <= base.Ui64(v1660) {
		v1727 = v1572
		goto L337
	} else {
		goto L338
	}
L329:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1061)+24)) = v1476
	v1578 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1578)+4))
	if v1579 == int32(0) {
		v1638 = v1572
		goto L328
	} else {
		goto L330
	}
L330:
	;
	if v1579 == v1578 {
		v1638 = v1572
		goto L328
	} else {
		goto L331
	}
L331:
	;
	v1583 = v1579
	v1589 = v1572
	goto L332
L332:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+4))
	v1610 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+24))
	v1613 = *(*int64)(unsafe.Add(mBase, uint32(v1583-int32(12))))
	if base.Ui64(v1610) < base.Ui64(v1613) {
		v1638 = v1589
		goto L328
	} else {
		goto L334
	}
L333:
	;
	v1638 = v1630
	goto L328
L334:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1583)))
	*(*int32)(unsafe.Add(mBase, uint32(v1615)+4)) = v1609
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1583)))
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1617
	*(*int64)(unsafe.Add(mBase, uint32(v1583))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1583-int32(4)))) = int32(2)
	F_SetLatch(m, v1583-int32(120))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	v1630 = v1589 + int32(1)
	if v1578 != v1609 {
		v1583 = v1609
		v1589 = v1630
		goto L332
	} else {
		goto L336
	}
L336:
	;
	goto L333
L337:
	;
	v1748 = int32(0)
	v1750 = v1061 + int32(40)
	v1751 = *(*int64)(unsafe.Add(mBase, uint32(v1750)))
	if base.Ui64(v1475) <= base.Ui64(v1751) {
		v1820 = v1748
		goto L346
	} else {
		goto L347
	}
L338:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1659))) = v1474
	v1664 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1664)+12))
	if v1665 == int32(0) {
		v1727 = v1572
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1669 = v1664 + int32(8)
	if v1665 == v1669 {
		v1727 = v1572
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v1673 = v1665
	v1678 = v1572
	goto L341
L341:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1673)+4))
	v1700 = *(*int64)(unsafe.Add(mBase, uint32(v1664+int32(32))))
	v1703 = *(*int64)(unsafe.Add(mBase, uint32(v1673-int32(12))))
	if base.Ui64(v1700) < base.Ui64(v1703) {
		v1727 = v1678
		goto L337
	} else {
		goto L343
	}
L342:
	;
	v1727 = v1720
	goto L337
L343:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1673)))
	*(*int32)(unsafe.Add(mBase, uint32(v1705)+4)) = v1699
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1673)))
	*(*int32)(unsafe.Add(mBase, uint32(v1699))) = v1707
	*(*int64)(unsafe.Add(mBase, uint32(v1673))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1673-int32(4)))) = int32(2)
	F_SetLatch(m, v1673-int32(120))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L7
	} else {
		goto L344
	}
L344:
	;
	v1720 = v1678 + int32(1)
	if v1669 != v1699 {
		v1673 = v1699
		v1678 = v1720
		goto L341
	} else {
		goto L345
	}
L345:
	;
	goto L342
L346:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v1840+int32(4096))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L7
	} else {
		goto L355
	}
L347:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1750))) = v1475
	v1755 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1755)+20))
	if v1756 == int32(0) {
		v1820 = v1748
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1760 = v1755 + int32(16)
	if v1756 == v1760 {
		v1820 = v1748
		goto L346
	} else {
		goto L349
	}
L349:
	;
	v1764 = v1756
	v1771 = v1748
	goto L350
L350:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1764)+4))
	v1791 = *(*int64)(unsafe.Add(mBase, uint32(v1755+int32(40))))
	v1794 = *(*int64)(unsafe.Add(mBase, uint32(v1764-int32(12))))
	if base.Ui64(v1791) < base.Ui64(v1794) {
		v1820 = v1771
		goto L346
	} else {
		goto L352
	}
L351:
	;
	v1820 = v1811
	goto L346
L352:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1764)))
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+4)) = v1790
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1764)))
	*(*int32)(unsafe.Add(mBase, uint32(v1790))) = v1798
	*(*int64)(unsafe.Add(mBase, uint32(v1764))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1764-int32(4)))) = int32(2)
	F_SetLatch(m, v1764-int32(120))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L7
	} else {
		goto L353
	}
L353:
	;
	v1811 = v1771 + int32(1)
	if v1760 != v1790 {
		v1764 = v1790
		v1771 = v1811
		goto L350
	} else {
		goto L354
	}
L354:
	;
	goto L351
L355:
	;
	v1847 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L7
	} else {
		goto L356
	}
L356:
	;
	if v1847 == int32(0) {
		goto L231
	} else {
		goto L357
	}
L357:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+32)) = uint32(v1475)
	v1852 = int64(32)
	v1853 = int64(base.Ui64(v1475) >> (uint(v1852) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+28)) = uint32(v1853)
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+24)) = v1820
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+20)) = uint32(v1474)
	v1858 = int64(base.Ui64(v1474) >> (uint(v1852) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+16)) = uint32(v1858)
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+12)) = v1727
	*(*int32)(unsafe.Add(mBase, uint32(v1042))) = v1638
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+8)) = uint32(v1476)
	v1864 = int64(base.Ui64(v1476) >> (uint(v1852) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1042)+4)) = uint32(v1864)
	F_errmsg_internal(m, int32(488879), v1042)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L7
	} else {
		goto L358
	}
L358:
	;
	F_errfinish(m, int32(474371), int32(572), int32(123514))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L7
	} else {
		goto L359
	}
L359:
	;
	goto L231
L360:
	;
	if v256 == int64(0) {
		v2285 = v1929
		goto L71
	} else {
		goto L361
	}
L361:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+88))
	if v1936 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	F_LogicalConfirmReceivedLocation(m, v256)
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L7
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1931)))
	*(*int32)(unsafe.Add(mBase, uint32(v1931))) = int32(1)
	if v1939 != 0 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	v2285 = v1929
	goto L71
L366:
	;
	F_s_lock(m, v1931, int32(474026), int32(2390), int32(253652))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L7
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1947 = *(*int64)(unsafe.Add(mBase, uint32(v1931)+104))
	if v1947 == v256 {
		goto L72
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1931))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1931)+104)) = v256
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L7
	} else {
		goto L371
	}
L371:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L7
	} else {
		goto L372
	}
L372:
	;
	v1958 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v1958 == int32(1) {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	if v1968 != 0 {
		v2285 = v1929
		goto L71
	} else {
		goto L377
	}
L374:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1963)+316))
	v1966 = base.B2i32(v1964 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v1966)
	v1968 = v1966
	goto L376
L375:
	;
	v1968 = int32(0)
	goto L376
L376:
	;
	goto L373
L377:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v1973 = int32(0)
	v1979 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	if v1979 == v1973 {
		v2007 = v1973
		goto L379
	} else {
		goto L380
	}
L378:
	;
	if v2007 == int32(0) {
		v2285 = v1929
		goto L71
	} else {
		goto L386
	}
L379:
	;
	goto L378
L380:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1979)))
	if v1982 <= int32(0) {
		v2007 = v1973
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1988 = v1979 + int32(4)
	v1990 = v1973
	goto L382
L382:
	;
	v1993 = F_strcmp(m, v1988, v1970+int32(24))
	mBase = m.M
	v1995 = base.B2i32(v1993 == int32(0))
	if v1993 == int32(0) {
		v2007 = v1995
		goto L379
	} else {
		goto L384
	}
L383:
	;
	v2007 = v1995
	goto L379
L384:
	;
	v1998 = F_strlen(m, v1988)
	mBase = m.M
	v2000 = int32(1)
	v2003 = v1990 + v2000
	if v2003 != v1982 {
		v1988 = v1998 + v1988 + v2000
		v1990 = v2003
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	F_ConditionVariableBroadcast(m, v2014+int32(76))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L7
	} else {
		goto L387
	}
L387:
	;
	v2285 = v1929
	goto L71
L388:
	;
	v2024 = F_pq_getmsgint(m, int32(4354320), int32(4))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L7
	} else {
		goto L389
	}
L389:
	;
	v2028 = F_pq_getmsgint(m, int32(4354320), int32(4))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L7
	} else {
		goto L390
	}
L390:
	;
	v2032 = F_pq_getmsgint(m, int32(4354320), int32(4))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L7
	} else {
		goto L391
	}
L391:
	;
	v2036 = F_pq_getmsgint(m, int32(4354320), int32(4))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L7
	} else {
		goto L392
	}
L392:
	;
	v2038 = int32(13)
	goto L395
L393:
	;
	if v2072 != 0 {
		goto L407
	} else {
		goto L408
	}
L394:
	;
	goto L393
L395:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	goto L398
L396:
	;
	v2057 = int32(0)
	goto L404
L398:
	;
	goto L399
L399:
	;
	goto L401
L401:
	;
	if v2045 == int32(15) {
		goto L396
	} else {
		goto L402
	}
L402:
	;
	if v2045 <= v2038 {
		v2072 = int32(1)
		goto L394
	} else {
		goto L403
	}
L403:
	;
	goto L396
L404:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	if v2061 != int32(2) {
		v2072 = v2057
		goto L394
	} else {
		goto L405
	}
L405:
	;
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, _consts[701])))
	if v2065 != 0 {
		v2072 = v2057
		goto L394
	} else {
		goto L406
	}
L406:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	v2072 = int32(0) | base.B2i32(v2069 <= v2038)
	goto L394
L407:
	;
	v2074 = F_timestamptz_to_str(m, v2020)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L7
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	v2101 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v2101)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+76)) = int32(1)
	if v2102 != 0 {
		goto L419
	} else {
		goto L420
	}
L410:
	;
	v2076 = F_pstrdup(m, v2074)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L7
	} else {
		goto L411
	}
L411:
	;
	v2080 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L7
	} else {
		goto L412
	}
L412:
	;
	if v2080 != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v2076
	*(*int32)(unsafe.Add(mBase, uint32(v64)+76)) = v2036
	*(*int32)(unsafe.Add(mBase, uint32(v64)+72)) = v2032
	*(*int32)(unsafe.Add(mBase, uint32(v64)+68)) = v2028
	*(*int32)(unsafe.Add(mBase, uint32(v64)+64)) = v2024
	F_errmsg_internal(m, int32(184243), v64-int32(-64))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L7
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	F_pfree(m, v2076)
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L7
	} else {
		goto L418
	}
L416:
	;
	F_errfinish(m, int32(474026), int32(2633), int32(387623))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L7
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	goto L409
L419:
	;
	F_s_lock(m, v2101+int32(76), int32(474026), int32(2645), int32(387623))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L7
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2101)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2101)+80)) = v2020
	if base.Ui32(int32(2)) < base.Ui32(v2024) {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	goto L421
L423:
	;
	if base.Ui32(v2024) < base.Ui32(int32(3)) {
		goto L428
	} else {
		goto L429
	}
L424:
	;
	if base.Ui32(int32(2)) < base.Ui32(v2032) {
		goto L423
	} else {
		goto L425
	}
L425:
	;
	v2120 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v2121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2120)+40)) = v2121
	v2123 = int32(1)
	v2125 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v2125 == v2121 {
		v2285 = v2123
		goto L71
	} else {
		goto L426
	}
L426:
	;
	F_PhysicalReplicationSlotNewXmin(m, v2024, v2032)
	mBase = m.M
	v2129 = m.ExcPending
	if v2129 != 0 {
		goto L7
	} else {
		goto L427
	}
L427:
	;
	v2285 = v2123
	goto L71
L428:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v2032) {
		goto L444
	} else {
		goto L445
	}
L429:
	;
	v2132 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L7
	} else {
		goto L430
	}
L430:
	;
	v2136 = base.I32_wrap_i64(int64(base.Ui64(v2132) >> (uint(int64(32)) % 64)))
	v2137 = base.I32_wrap_i64(v2132)
	if base.Ui32(v2024) <= base.Ui32(v2137) {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2137))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2024)) == int32(0) {
		goto L438
	} else {
		goto L439
	}
L432:
	;
	if v2136 == v2028 {
		goto L431
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	v2141 = int32(1)
	if v2028+v2141 != v2136 {
		v2285 = v2141
		goto L71
	} else {
		goto L436
	}
L435:
	;
	v2285 = int32(1)
	goto L71
L436:
	;
	goto L431
L437:
	;
	if v2157 != 0 {
		goto L428
	} else {
		goto L441
	}
L438:
	;
	v2157 = base.B2i32(base.Ui32(v2024) <= base.Ui32(v2137))
	goto L437
L439:
	;
	goto L440
L440:
	;
	v2157 = base.B2i32(v2024-v2137 <= int32(0))
	goto L437
L441:
	;
	v2285 = int32(1)
	goto L71
L442:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	*(*int32)(unsafe.Add(mBase, uint32(v2229)+40)) = v2024
	v2285 = int32(1)
	goto L71
L443:
	;
	F_PhysicalReplicationSlotNewXmin(m, v2024, v2032)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L7
	} else {
		goto L466
	}
L444:
	;
	v2165 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L7
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v2214 == int32(0) {
		goto L442
	} else {
		goto L465
	}
L447:
	;
	v2169 = base.I32_wrap_i64(int64(base.Ui64(v2165) >> (uint(int64(32)) % 64)))
	v2170 = base.I32_wrap_i64(v2165)
	if base.Ui32(v2032) <= base.Ui32(v2170) {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	v2179 = int32(1)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2170))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2032)) == int32(0) {
		goto L455
	} else {
		goto L456
	}
L449:
	;
	if v2036 == v2169 {
		goto L448
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v2174 = int32(1)
	if v2036+v2174 != v2169 {
		v2285 = v2174
		goto L71
	} else {
		goto L453
	}
L452:
	;
	v2285 = int32(1)
	goto L71
L453:
	;
	goto L448
L454:
	;
	if v2191 == int32(0) {
		v2285 = v2179
		goto L71
	} else {
		goto L458
	}
L455:
	;
	v2191 = base.B2i32(base.Ui32(v2032) <= base.Ui32(v2170))
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2191 = base.B2i32(v2032-v2170 <= int32(0))
	goto L454
L458:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	if v2195 != 0 {
		goto L443
	} else {
		goto L459
	}
L459:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2024))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v2032)) == int32(0) {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	if v2207 == int32(0) {
		goto L442
	} else {
		goto L464
	}
L461:
	;
	v2207 = base.B2i32(base.Ui32(v2032) < base.Ui32(v2024))
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2207 = int32(base.Ui32(v2032-v2024) >> (uint(int32(31)) % 32))
	goto L460
L464:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	*(*int32)(unsafe.Add(mBase, uint32(v2211)+40)) = v2032
	v2285 = v2179
	goto L71
L465:
	;
	goto L443
L466:
	;
	v2285 = int32(1)
	goto L71
L467:
	;
	if v2234 != 0 {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L7
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L7
	} else {
		goto L474
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v249
	F_errmsg(m, int32(689500), v64+int32(16))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L7
	} else {
		goto L472
	}
L472:
	;
	F_errfinish(m, int32(474026), int32(2375), int32(387486))
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L7
	} else {
		goto L473
	}
L473:
	;
	goto L470
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	v2258 = int32(0)
	v2261 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2261)+20))
	m.T0[v2262].(func(*base.Module, int32, int32, int32))(m, int32(99), v2258, v2258)
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L7
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v2269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[690])) = uint8(v2269)
	v2313 = v64
	goto L4
L478:
	;
	v2266 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[699])) = uint8(v2266)
	goto L477
L479:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L480:
	;
	goto L6
L481:
	;
	v2313 = v64
	goto L4
}
func F_ProcessUtilityForAlterTable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	F_EventTriggerAlterTableEnd(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = F_palloc0(m, int32(104))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = l0
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v14)
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(25769804106)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v22
			v25 = *(*int32)(unsafe.Add(mBase, _consts[881]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v30 = *(*int32)(unsafe.Add(mBase, _consts[885]))
			if v30 != 0 {
				v31 = int32(0)
				m.T0[v30].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, v11, v28, v31, int32(3), v27, v26, v25, v31)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
					F_EventTriggerAlterTableStart(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v48 = *(*int32)(unsafe.Add(mBase, _consts[385]))
						if v48 == int32(0) {
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
							if v51 != 0 {
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v45
							}
						}
						return
					}
				}
			} else {
				v36 = int32(0)
				F_standard_ProcessUtility(m, v11, v28, v36, int32(3), v27, v26, v25, v36)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
					F_EventTriggerAlterTableStart(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v48 = *(*int32)(unsafe.Add(mBase, _consts[385]))
						if v48 == int32(0) {
						} else {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
							if v51 != 0 {
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v45
							}
						}
						return
					}
				}
			}
		}
	}
}
func F_PushActiveSnapshot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	F_PushActiveSnapshotWithLevel(m, l0, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_PushActiveSnapshotWithLevel(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	v10 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v12 = F_MemoryContextAlloc(m, v10, int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[1247]))
		if l0 == v15 {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
			v26 = l0 + int32(24)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v29 = l0 + int32(16)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
			v32 = int32(2)
			v34 = int32(72)
			v39 = v30<<(uint(v32)%32) + v34
			if int32(0) < v27 {
				v42 = (v27+v30)<<(uint(v32)%32) + v34
			} else {
				v42 = v39
			}
			v43 = F_MemoryContextAlloc(m, v24, v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v46 = v43 + int32(48)
				v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
				v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+40)) = v49
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v51
				v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+56)) = v53
				v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = v55
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v57
				v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v59
				v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v61
				*(*int64)(unsafe.Add(mBase, uint32(v43)+64)) = int64(0)
				v65 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v46))) = v65
				*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v65
				v69 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+30)) = uint8(v69)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				if v71 != 0 {
					v73 = v43 + int32(72)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v73
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v78 = v76 << (uint(int32(2)) % 32)
					if v78 != 0 {
						v79 = F__emscripten_memcpy_bulkmem(m, v73, v75, v78)
						mBase = m.M
					} else {
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = int32(0)
				}
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v84 <= int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
					v103 = v43
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v87 == int32(1) {
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
						if v90 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
							v103 = v43
						} else {
							v93 = v43 + v39
							*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v98 = v96 << (uint(int32(2)) % 32)
							if v98 != 0 {
								v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
								mBase = m.M
							} else {
							}
							v103 = v43
						}
					} else {
						v93 = v43 + v39
						*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v98 = v96 << (uint(int32(2)) % 32)
						if v98 != 0 {
							v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
							mBase = m.M
						} else {
						}
						v103 = v43
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
				v109 = int32(4443972)
				v110 = *(*int32)(unsafe.Add(mBase, _consts[131]))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v110
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
				*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = v113 + int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[131])) = v12
				return
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[1248]))
			if l0 == v18 {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
				v26 = l0 + int32(24)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v29 = l0 + int32(16)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v32 = int32(2)
				v34 = int32(72)
				v39 = v30<<(uint(v32)%32) + v34
				if int32(0) < v27 {
					v42 = (v27+v30)<<(uint(v32)%32) + v34
				} else {
					v42 = v39
				}
				v43 = F_MemoryContextAlloc(m, v24, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v46 = v43 + int32(48)
					v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
					v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+40)) = v49
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v51
					v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+56)) = v53
					v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = v55
					v57 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v57
					v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v59
					v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = v61
					*(*int64)(unsafe.Add(mBase, uint32(v43)+64)) = int64(0)
					v65 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = v65
					*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v65
					v69 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v43)+30)) = uint8(v69)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					if v71 != 0 {
						v73 = v43 + int32(72)
						*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v73
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v78 = v76 << (uint(int32(2)) % 32)
						if v78 != 0 {
							v79 = F__emscripten_memcpy_bulkmem(m, v73, v75, v78)
							mBase = m.M
						} else {
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = int32(0)
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v84 <= int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
						v103 = v43
					} else {
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v87 == int32(1) {
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
							if v90 != int32(1) {
								*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
								v103 = v43
							} else {
								v93 = v43 + v39
								*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v98 = v96 << (uint(int32(2)) % 32)
								if v98 != 0 {
									v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
									mBase = m.M
								} else {
								}
								v103 = v43
							}
						} else {
							v93 = v43 + v39
							*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v98 = v96 << (uint(int32(2)) % 32)
							if v98 != 0 {
								v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
								mBase = m.M
							} else {
							}
							v103 = v43
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
					v109 = int32(4443972)
					v110 = *(*int32)(unsafe.Add(mBase, _consts[131]))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v110
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = v113 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[131])) = v12
					return
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
				if v20 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, _consts[68]))
					v26 = l0 + int32(24)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v29 = l0 + int32(16)
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v32 = int32(2)
					v34 = int32(72)
					v39 = v30<<(uint(v32)%32) + v34
					if int32(0) < v27 {
						v42 = (v27+v30)<<(uint(v32)%32) + v34
					} else {
						v42 = v39
					}
					v43 = F_MemoryContextAlloc(m, v24, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v46 = v43 + int32(48)
						v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
						v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+40)) = v49
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+24)) = v51
						v53 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+56)) = v53
						v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = v55
						v57 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = v57
						v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v43)+8)) = v59
						v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v43))) = v61
						*(*int64)(unsafe.Add(mBase, uint32(v43)+64)) = int64(0)
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v46))) = v65
						*(*int32)(unsafe.Add(mBase, uint32(v43)+44)) = v65
						v69 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v43)+30)) = uint8(v69)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						if v71 != 0 {
							v73 = v43 + int32(72)
							*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v73
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v78 = v76 << (uint(int32(2)) % 32)
							if v78 != 0 {
								v79 = F__emscripten_memcpy_bulkmem(m, v73, v75, v78)
								mBase = m.M
							} else {
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = int32(0)
						}
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v84 <= int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
							v103 = v43
						} else {
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v87 == int32(1) {
								v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
								if v90 != int32(1) {
									*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = int32(0)
									v103 = v43
								} else {
									v93 = v43 + v39
									*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v98 = v96 << (uint(int32(2)) % 32)
									if v98 != 0 {
										v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
										mBase = m.M
									} else {
									}
									v103 = v43
								}
							} else {
								v93 = v43 + v39
								*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v93
								v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v98 = v96 << (uint(int32(2)) % 32)
								if v98 != 0 {
									v99 = F__emscripten_memcpy_bulkmem(m, v93, v95, v98)
									mBase = m.M
								} else {
								}
								v103 = v43
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
						v109 = int32(4443972)
						v110 = *(*int32)(unsafe.Add(mBase, _consts[131]))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v110
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
						*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = v113 + int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[131])) = v12
						return
					}
				} else {
					v103 = l0
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
					v109 = int32(4443972)
					v110 = *(*int32)(unsafe.Add(mBase, _consts[131]))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v110
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = v113 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[131])) = v12
					return
				}
			}
		}
	}
}
func F_p_iseqC(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9))))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		v14 = base.B2i32(v11 == v12)
	} else {
		v14 = int32(0)
	}
	return v14
}
func F_p_isxdigit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
			if base.Ui32(int32(127)) < base.Ui32(v13) {
				v65 = int32(0)
				return v65
			} else {
				return base.B2i32(base.Ui32(v13-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v13|int32(32)-int32(97)) < base.Ui32(int32(6)))
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(v34-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v34|int32(32)-int32(97)) < base.Ui32(int32(6)))
		}
	} else {
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v49))))
		v65 = base.B2i32(base.Ui32(v51-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v51|int32(32)-int32(97)) < base.Ui32(int32(6)))
		return v65
	}
}
func F_parse_ident(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
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
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_text_to_cstring(m, v13)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v18
	goto L4
L4:
	;
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20))))
	goto L6
L5:
	;
	v40 = int32(0)
	v41 = int32(0)
	v42 = v20
	goto L11
L6:
	;
	if base.B2i32(v29 == int32(32))|base.B2i32(base.Ui32((v29-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v20 = v20 + int32(1)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	F_errdetail(m, int32(629864), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L153
	}
L9:
	;
	F_errdetail(m, int32(629895), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L151
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L145
	}
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 == int32(95) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v473 = F_makeArrayResult(m, v381, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L144
	}
L13:
	;
	goto L12
L14:
	;
	v453 = v383
	goto L140
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L133
	}
L16:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v381 = F_accumArrayResult(m, v41, v376, int32(0), int32(25), v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L120
	}
L17:
	;
	v331 = v42
	goto L111
L18:
	;
	if v47 == int32(34) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v53 = v42 + int32(1)
	v54 = int32(34)
	v55 = F___strchrnul(m, v53, v54)
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v57 == v54 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	goto L21
L21:
	;
	v319 = base.I32_extend8_s(v47)
	if v319 < int32(0) {
		goto L17
	} else {
		goto L109
	}
L22:
	;
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v314)
	if v63 == v53 {
		goto L10
	} else {
		goto L107
	}
L23:
	;
	if v61 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v61 = v55
	goto L26
L25:
	;
	v61 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	v63 = v61
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L101
	}
L30:
	;
	v70 = v63 + int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v71 != int32(34) {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v63&int32(3) == int32(0) {
		v97 = v63
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v63 == v70 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	v130 = v122 - v63
	goto L33
L35:
	;
	v101 = v97
	goto L44
L36:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v81 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v130 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v86 = v63
	goto L40
L40:
	;
	v90 = v86 + int32(1)
	if v90&int32(3) == int32(0) {
		v97 = v90
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v122 = v90
	goto L34
L42:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v95 != 0 {
		v86 = v90
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v110 = int32(-2139062144)
	if (int32(16843008)-v107|v107)&v110 == v110 {
		v101 = v101 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v116 = v101
	goto L47
L46:
	;
	goto L45
L47:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v120 != 0 {
		v116 = v116 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v122 = v116
	goto L34
L49:
	;
	goto L48
L50:
	;
	v275 = int32(34)
	v276 = F___strchrnul(m, v70, v275)
	mBase = m.M
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v278 == v275 {
		goto L97
	} else {
		goto L98
	}
L51:
	;
	goto L50
L52:
	;
	v134 = v63 + v130
	if base.Ui32(v70-v134) <= base.Ui32(int32(0)-v130<<(uint(int32(1))%32)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v141 = F___memcpy(m, v63, v70, v130)
	mBase = m.M
	goto L50
L54:
	;
	goto L55
L55:
	;
	v144 = (v63 ^ v70) & int32(3)
	if base.Ui32(v63) < base.Ui32(v70) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v246 == int32(0) {
		goto L51
	} else {
		goto L92
	}
L57:
	;
	if base.Ui32(v224) <= base.Ui32(int32(3)) {
		v245 = v223
		v246 = v224
		v247 = v225
		goto L56
	} else {
		goto L88
	}
L58:
	;
	if v144 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	if v144 != 0 {
		v206 = v130
		goto L71
	} else {
		goto L72
	}
L61:
	;
	v245 = v70
	v246 = v130
	v247 = v63
	goto L56
L62:
	;
	goto L63
L63:
	;
	if v63&int32(3) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v223 = v70
	v224 = v130
	v225 = v63
	goto L57
L65:
	;
	goto L66
L66:
	;
	v151 = v70
	v152 = v130
	v153 = v63
	goto L67
L67:
	;
	if v152 == int32(0) {
		goto L51
	} else {
		goto L69
	}
L68:
	;
	v223 = v160
	v224 = v162
	v225 = v164
	goto L57
L69:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v157)
	v159 = int32(1)
	v160 = v151 + v159
	v162 = v152 - v159
	v164 = v153 + v159
	if v164&int32(3) != 0 {
		v151 = v160
		v152 = v162
		v153 = v164
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	if v206 == int32(0) {
		goto L51
	} else {
		goto L84
	}
L72:
	;
	if v134&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v171 = v130
	goto L76
L74:
	;
	v186 = v130
	goto L75
L75:
	;
	if base.Ui32(v186) <= base.Ui32(int32(3)) {
		v206 = v186
		goto L71
	} else {
		goto L80
	}
L76:
	;
	if v171 == int32(0) {
		goto L51
	} else {
		goto L78
	}
L77:
	;
	v186 = v177
	goto L75
L78:
	;
	v177 = v171 - int32(1)
	v178 = v63 + v177
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v177))))
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v180)
	if v178&int32(3) != 0 {
		v171 = v177
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v193 = v186
	goto L81
L81:
	;
	v197 = v193 - int32(4)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v70+v197)))
	*(*int32)(unsafe.Add(mBase, uint32(v63+v197))) = v200
	if base.Ui32(int32(3)) < base.Ui32(v197) {
		v193 = v197
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v206 = v197
	goto L71
L83:
	;
	goto L82
L84:
	;
	v213 = v206
	goto L85
L85:
	;
	v217 = v213 - int32(1)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63+v217))) = uint8(v220)
	if v217 != 0 {
		v213 = v217
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L51
L87:
	;
	goto L86
L88:
	;
	v230 = v223
	v231 = v224
	v232 = v225
	goto L89
L89:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v234
	v236 = int32(4)
	v237 = v230 + v236
	v239 = v232 + v236
	v241 = v231 - v236
	if base.Ui32(int32(3)) < base.Ui32(v241) {
		v230 = v237
		v231 = v241
		v232 = v239
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v245 = v237
	v246 = v241
	v247 = v239
	goto L56
L91:
	;
	goto L90
L92:
	;
	v252 = v245
	v253 = v246
	v254 = v247
	goto L93
L93:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v256)
	v258 = int32(1)
	v263 = v253 - v258
	if v263 != 0 {
		v252 = v252 + v258
		v253 = v263
		v254 = v254 + v258
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L51
L95:
	;
	goto L94
L96:
	;
	if v282 != 0 {
		v63 = v282
		goto L30
	} else {
		goto L100
	}
L97:
	;
	v282 = v276
	goto L99
L98:
	;
	v282 = int32(0)
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L31
L101:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v297 = F_text_to_cstring(m, v13)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v297
	F_errmsg(m, int32(685744), v8+int32(-32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errdetail(m, int32(558757), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(478044), int32(899), int32(91187))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v317 = F_cstring_to_text(m, v53)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v369 = v70
	v376 = v317
	goto L16
L109:
	;
	if base.Ui32(int32(25)) < base.Ui32((v319&int32(-33)-int32(65))&int32(255)) {
		goto L15
	} else {
		goto L110
	}
L110:
	;
	goto L17
L111:
	;
	v339 = v331 + int32(1)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if base.Ui32((v340-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v331 = v339
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v362 = v339 - v42
	v363 = int32(0)
	v365 = F_downcase_identifier(m, v42, v362, v363, v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L118
	}
L113:
	;
	if v340 == int32(36) {
		v331 = v339
		goto L111
	} else {
		goto L114
	}
L114:
	;
	if v340 == int32(95) {
		v331 = v339
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v351 = base.I32_extend8_s(v340)
	if v351 < int32(0) {
		v331 = v339
		goto L111
	} else {
		goto L116
	}
L116:
	;
	if base.Ui32((v351&int32(-33)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		v331 = v339
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L112
L118:
	;
	v367 = F_cstring_to_text_with_len(m, v365, v362)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v369 = v339
	v376 = v367
	goto L16
L120:
	;
	v383 = v369
	goto L121
L121:
	;
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v383))))
	goto L123
L122:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	if v402 == int32(46) {
		goto L14
	} else {
		goto L125
	}
L123:
	;
	if base.B2i32(v392 == int32(32))|base.B2i32(base.Ui32((v392-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v383 = v383 + int32(1)
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v402 == int32(0) {
		goto L13
	} else {
		goto L126
	}
L126:
	;
	if v17 == int32(0) {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v416 = F_text_to_cstring(m, v13)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v416
	F_errmsg(m, int32(685744), v10)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(478044), int32(986), int32(91187))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v434 = F_text_to_cstring(m, v13)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v434
	F_errmsg(m, int32(685744), v8+int32(-48))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v319 == int32(46) {
		goto L9
	} else {
		goto L137
	}
L137:
	;
	if v40&int32(1) != 0 {
		goto L8
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(478044), int32(963), int32(91187))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	v458 = int32(1)
	v460 = v453 + v458
	v461 = int32(*(*int8)(unsafe.Add(mBase, uint32(v460))))
	goto L142
L141:
	;
	v40 = v458
	v41 = v381
	v42 = v460
	goto L11
L142:
	;
	if base.B2i32(v461 == int32(32))|base.B2i32(base.Ui32((v461-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v453 = v460
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	m.G0 = v10 - int32(-64)
	return v473
L145:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v486 = F_text_to_cstring(m, v13)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v486
	F_errmsg(m, int32(685744), v8+int32(-16))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errdetail(m, int32(532492), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(478044), int32(913), int32(91187))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
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
	F_errfinish(m, int32(478044), int32(952), int32(91187))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errfinish(m, int32(478044), int32(958), int32(91187))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_real(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 float64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v182 int32
	_ = v182
	var v185 float64
	_ = v185
	var v195 int32
	_ = v195
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
	v22 = F_strtod(m, l0, v11+int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v22
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v27 == l0 {
		v195 = v5
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v11 + int32(16)
	return v195
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v30 == int32(68) {
		v195 = v5
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v22)&int64(9223372036854775807)) {
		v195 = v5
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v43 = v27
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if base.Ui32(v46-int32(9)) < base.Ui32(int32(5)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = v43 + int32(1)
	goto L13
L16:
	;
	if v46 == int32(32) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v46 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v182 = int32(1)
	if l1 == int32(0) {
		v195 = v182
		goto L9
	} else {
		goto L51
	}
L19:
	;
	v56 = l2 & int32(2130706432)
	if v56 == int32(0) {
		v195 = v5
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v66 = m.G0
	v68 = v66 - int32(16)
	m.G0 = v68
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	switch v72 {
	case 0, 9, 10, 11, 12, 13, 32:
		v91 = v43
		v92 = v68 + int32(12)
		goto L22
	default:
		goto L23
	}
L21:
	;
	if v163 != 0 {
		goto L18
	} else {
		goto L46
	}
L22:
	;
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v93)
	v102 = v91
	goto L26
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)) = uint8(v72)
	v77 = v43 + int32(1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	switch v78 {
	case 0, 9, 10, 11, 12, 13, 32:
		v91 = v77
		v92 = v68 + int32(13)
		goto L22
	default:
		goto L24
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+13)) = uint8(v78)
	v83 = v43 + int32(2)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	switch v84 {
	case 0, 9, 10, 11, 12, 13, 32:
		v91 = v83
		v92 = v68 + int32(14)
		goto L22
	default:
		goto L25
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+14)) = uint8(v84)
	v91 = v43 + int32(3)
	v92 = v68 + int32(15)
	goto L22
L26:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if base.Ui32(v105-int32(9)) < base.Ui32(int32(5)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v102 = v102 + int32(1)
	goto L26
L29:
	;
	if v105 == int32(32) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v105 != 0 {
		v163 = v93
		goto L31
	} else {
		goto L32
	}
L31:
	;
	m.G0 = v68 + int32(16)
	goto L21
L32:
	;
	if v56&int32(251658240) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = int32(1722560)
	goto L35
L34:
	;
	v116 = int32(1722976)
	goto L35
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v117 == int32(0) {
		v163 = v93
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v124 = v93
	goto L37
L37:
	;
	v131 = v116 + v124<<(uint(int32(4))%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v56 != v132 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v163 = int32(0)
	goto L31
L39:
	;
	v153 = v124 + int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116+v153<<(uint(int32(4))%32)))))
	if v157 != 0 {
		v124 = v153
		goto L37
	} else {
		goto L45
	}
L40:
	;
	v136 = F_strcmp(m, v68+int32(12), v131)
	mBase = m.M
	if v136 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v131)+8))
	v138 = base.F64_mul(v22, v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+16)))
	if v139 == int32(0) {
		v148 = v138
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11+int32(8)))) = v148
	v163 = int32(1)
	goto L31
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
	if v56 != v142 {
		v148 = v138
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v131)+24))
	v148 = base.F64_mul(v144, base.F64_nearest(base.F64_div(v138, v144)))
	goto L42
L45:
	;
	goto L38
L46:
	;
	if l3 == int32(0) {
		v195 = v5
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if l2&int32(251658240) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(629732)
	v195 = v5
	goto L9
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(629353)
	v195 = v5
	goto L9
L51:
	;
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v185
	v195 = v182
	goto L9
}
func F_performMultipleDeletions(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v4 < v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v20
	v24 = F_palloc(m, int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(137438953472)
	v29 = F_palloc(m, int32(384))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v34 <= v31 {
		v88 = v4
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_reportDependentObjects(m, v24, l1, l2, v88)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L26
	}
L9:
	;
	v43 = v4
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = v49 + v43*int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	switch v54 - int32(1259) {
	case 0:
		goto L15
	default:
		goto L13
	case 2:
		goto L14
	}
L11:
	;
	if v78 != int32(1) {
		v88 = int32(0)
		goto L8
	} else {
		goto L25
	}
L12:
	;
	F_findDependentObjects(m, v52, int32(1), l2, int32(0), v24, l0, v13+int32(12))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L23
	}
L13:
	;
	F_LockDatabaseObject(m, v54, v53, int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L22
	}
L14:
	;
	F_LockSharedObject(m, int32(1261), v53, int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	if l2&int32(2) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_LockRelationOid(m, v53, int32(4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_LockRelationOid(m, v53, int32(8))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L12
L20:
	;
	goto L12
L21:
	;
	goto L12
L22:
	;
	goto L12
L23:
	;
	v77 = v43 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v77 < v78 {
		v43 = v77
		goto L10
	} else {
		goto L24
	}
L24:
	;
	goto L11
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = v83
	goto L8
L26:
	;
	F_deleteObjectsInList(m, v24, v13+int32(12), l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	F_pfree(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v103 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_pfree(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v24)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_sequence_close(m, v108, int32(3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L3
}
func F_perform_pullup_replace_vars(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v15 = F_replace_rte_variables(m, v10, v11, v8, int32(850), l1, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v25 = F_replace_rte_variables(m, v20, v21, int32(0), int32(850), l1, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v7
	return
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v33 = F_replace_rte_variables(m, v28, v29, int32(0), int32(850), l1, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v42 = F_replace_rte_variables(m, v37, v38, int32(0), int32(850), l1, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v52 = F_replace_rte_variables(m, v47, v48, int32(0), int32(850), l1, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = v52
	goto L10
L13:
	;
	v100 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v106 = F_replace_rte_variables(m, v101, v102, v100, int32(850), l1, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L21
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v67 = int32(0)
	goto L16
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v67<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v79 = F_replace_rte_variables(m, v74, v75, int32(0), int32(850), l1, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v87 = F_replace_rte_variables(m, v82, v83, int32(0), int32(850), l1, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v87
	v91 = v67 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v91 < v92 {
		v67 = v91
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_replace_vars_in_jointree(m, v109, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v117 = F_replace_rte_variables(m, v112, v113, int32(0), int32(850), l1, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v120 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	if v155 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v123 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v128 = v100
	goto L27
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v142 = F_replace_rte_variables(m, v137, v138, int32(0), int32(850), l1, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = v142
	v146 = v128 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v146 < v147 {
		v128 = v146
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return
L32:
	;
	v158 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v159 <= v158 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v164 = v158
	goto L34
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v169 = int32(2)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v164<<(uint(v169)%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	switch v173 - v169 {
	case 0:
		goto L38
	default:
		goto L36
	case 7:
		goto L37
	}
L35:
	;
	goto L31
L36:
	;
	v193 = v164 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v193 < v194 {
		v164 = v193
		goto L34
	} else {
		goto L41
	}
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v189 = F_replace_rte_variables(m, v184, v185, int32(0), int32(850), l1, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)+52))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v181 = F_replace_rte_variables(m, v176, v177, int32(0), int32(850), l1, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+52)) = v181
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = v189
	goto L36
L41:
	;
	goto L35
}
func F_perform_spin_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v60 float64
	_ = v60
	var v67 float64
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[736]))
	if v9 <= v6 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v13 = v11 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
		if int32(1001) <= v13 {
			v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			F_s_lock_stuck(m, v82, v83, v84)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v17 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1000)
			} else {
			}
			v23 = *(*int32)(unsafe.Add(mBase, _consts[189]))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(150994950)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_pg_usleep(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[189]))
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v34 = int32(4527360)
				v37 = *(*int64)(unsafe.Add(mBase, _consts[87]))
				v38 = *(*int64)(unsafe.Add(mBase, _consts[88]))
				v39 = v37 ^ v38
				*(*int64)(unsafe.Add(mBase, _consts[88])) = base.I64_rotl(v39, int64(37))
				*(*int64)(unsafe.Add(mBase, _consts[87])) = v39<<(uint(int64(16))%64) ^ base.I64_rotl(v37, int64(24)) ^ v39
				v60 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v37*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v67 = base.F64_add(base.F64_mul(v60, base.F64_convert_i32_s(v33)), float64(0.5))
				if base.F64_lt(base.F64_abs(v67), float64(2.147483648e+09)) != 0 {
					v71 = base.I32_trunc_f64_s(v67)
					v73 = v71
				} else {
					v73 = int32(-2147483648)
				}
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v75 = v73 + v74
				if int32(1000000) < v75 {
					v78 = int32(1000)
				} else {
					v78 = v75
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78
				return
			}
		}
	} else {
		return
	}
}
func F_pgarch_archiveXlog(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int64
	_ = v488
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v17 = v2
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = int32(-1)
	v23 = v12
	goto L1
L1:
	;
	goto L3
L2:
	;
	m.G0 = v12 - int32(-64)
	return v405
L3:
	;
	if v22 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v487 = int32(m.ExcTag)
	v488 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v487 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L6:
	;
	v27 = v23 - int32(160)
	m.G0 = v27
	v29 = int32(1024)
	v30 = v27 - v29
	m.G0 = v30
	v33 = v30 - int32(80)
	m.G0 = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	v44 = F_pg_snprintf(m, v30, v29, int32(167530), v10+int32(-32))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		v486 = v33
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v134 = v17
	v135 = v18
	v136 = v19
	v137 = v20
	v138 = v21
	v139 = v23
	goto L8
L8:
	;
	if v138 != 0 {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	v55 = F_pg_snprintf(m, v33, int32(80), int32(176773), v10+int32(-48))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		v486 = v33
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v27
	if v33&int32(3) == int32(0) {
		v84 = v33
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v27
	v123 = int32(4443856)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v127 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v127
	goto L28
L12:
	;
	goto L11
L13:
	;
	v88 = v84
	goto L22
L14:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v68 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	goto L17
L17:
	;
	v73 = v33
	goto L18
L18:
	;
	v77 = v73 + int32(1)
	if v77&int32(3) == int32(0) {
		v84 = v77
		goto L13
	} else {
		goto L20
	}
L19:
	;
	goto L12
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v82 != 0 {
		v73 = v77
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v97 = int32(-2139062144)
	if (int32(16843008)-v94|v94)&v97 == v97 {
		v88 = v88 + int32(4)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v103 = v88
	goto L25
L24:
	;
	goto L23
L25:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v107 != 0 {
		v103 = v103 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L12
L27:
	;
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v10 + int32(-20)
	goto L31
L29:
	;
	v134 = v33
	v135 = v27
	v136 = v30
	v137 = v124
	v138 = int32(0)
	v139 = v33
	goto L8
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	v413 = F_pg_snprintf(m, v134, int32(80), v406, v12)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L50
	}
L33:
	;
	v405 = int32(0)
	v406 = int32(174881)
	goto L32
L34:
	;
	v140 = int32(4438508)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v142 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[84])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_EmitErrorReport(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[261])) = v135
	v374 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	v381 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v382 = m.T0[v375].(func(*base.Module, int32, int32, int32) int32)(m, v381, l0, v136)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L47
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[524])) = v160
	*(*int32)(unsafe.Add(mBase, _consts[525])) = v160
	*(*uint8)(unsafe.Add(mBase, _consts[526])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[527])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[528])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[529])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[530])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[531])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[532])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[533])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[534])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[535])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[536])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[537])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[538])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[539])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[540])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[541])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[542])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[543])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[544])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[545])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[546])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[547])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[548])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[549])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[550])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[551])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[552])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[553])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[554])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[555])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[556])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[557])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[558])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[559])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[560])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[561])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[562])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[563])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[564])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[565])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[566])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[567])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[568])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[569])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[570])) = uint8(v160)
	*(*uint8)(unsafe.Add(mBase, _consts[571])) = uint8(v160)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_LWLockReleaseAll(m)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	F_FlushErrorState(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	v359 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_MemoryContextReset(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v362 = int32(4438508)
	v364 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v364 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[261])) = int32(0)
	goto L33
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v137
	*(*int32)(unsafe.Add(mBase, _consts[261])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	v394 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_MemoryContextReset(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		v486 = v139
		goto L5
	} else {
		goto L48
	}
L48:
	;
	if v382 == int32(0) {
		goto L33
	} else {
		goto L49
	}
L49:
	;
	v405 = int32(1)
	v406 = int32(170978)
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	if v134&int32(3) == int32(0) {
		v442 = v134
		goto L53
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v135
	goto L4
L52:
	;
	goto L51
L53:
	;
	v446 = v442
	goto L62
L54:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v426 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	goto L57
L57:
	;
	v431 = v134
	goto L58
L58:
	;
	v435 = v431 + int32(1)
	if v435&int32(3) == int32(0) {
		v442 = v435
		goto L53
	} else {
		goto L60
	}
L59:
	;
	goto L52
L60:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v440 != 0 {
		v431 = v435
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v455 = int32(-2139062144)
	if (int32(16843008)-v452|v452)&v455 == v455 {
		v446 = v446 + int32(4)
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v461 = v446
	goto L65
L64:
	;
	goto L63
L65:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v465 != 0 {
		v461 = v461 + int32(1)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L52
L67:
	;
	goto L66
L68:
	;
	v492 = int32(v488)
	m.G0 = v486
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v492)))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v10+int32(-20) == v499 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	m.ExcPending = 1
	goto L77
L70:
	;
	if v502 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v502 = v501
	goto L73
L72:
	;
	v502 = int32(0)
	goto L73
L73:
	;
	goto L70
L74:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v17 = v505
	v18 = v503
	v19 = v504
	v20 = v506
	v21 = v494
	v22 = v502
	v23 = v486
	goto L1
L75:
	;
	goto L76
L76:
	;
	F___wasm_longjmp(m, v495, v494)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	return int32(0)
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pglz_compress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v387 int32
	_ = v387
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
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
	var v447 int32
	_ = v447
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	v5 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v5)
	v34 = int32(-1)
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v30 + int32(16)
	return v742
L2:
	;
	v36 = l3
	goto L4
L3:
	;
	v36 = int32(1822924)
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v37 <= int32(0) {
		v742 = v34
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if l1 < v40 {
		v742 = v34
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v42 < l1 {
		v742 = v34
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v44 = l0 + l1
	v45 = int32(100)
	v47 = int32(99)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v47 <= v48 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = v47
	goto L10
L9:
	;
	v51 = v48
	goto L10
L10:
	;
	if v48 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v45
	goto L13
L12:
	;
	v55 = v45 - v51
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if int32(21474837) <= l1 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v87 = F__emscripten_memset_bulkmem(m, int32(4445424), base.I32_extend8_s(int32(0)), v83<<(uint(int32(1))%32))
	mBase = m.M
	goto L25
L15:
	;
	if base.Ui32(l1) < base.Ui32(int32(1024)) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v62 = base.I32_div_u_s(l1, int32(100))
	v76 = v55 * v62
	goto L15
L17:
	;
	goto L18
L18:
	;
	v66 = base.I32_div_s(l1*v55, int32(100))
	if l1 < int32(128) {
		v82 = v66
		v83 = int32(512)
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(l1) < base.Ui32(int32(256)) {
		v82 = v66
		v83 = int32(1024)
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(int32(512)) <= base.Ui32(l1) {
		v76 = v66
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v82 = v66
	v83 = int32(2048)
	goto L14
L22:
	;
	v81 = int32(4096)
	goto L24
L23:
	;
	v81 = int32(8192)
	goto L24
L24:
	;
	v82 = v76
	v83 = v81
	goto L14
L25:
	;
	if base.Ui32(v44) <= base.Ui32(l0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v719))) = uint8(v714)
	v735 = v713 - l2
	if v82 <= v735 {
		goto L144
	} else {
		goto L145
	}
L27:
	;
	v713 = l2
	v714 = v5
	v719 = v30 + int32(15)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v92 = int32(17)
	if base.Ui32(v37) <= base.Ui32(v92) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v95 = v92
	goto L32
L31:
	;
	v95 = v37
	goto L32
L32:
	;
	if base.Ui32(int32(273)) <= base.Ui32(v95) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = int32(273)
	goto L35
L34:
	;
	v98 = v95
	goto L35
L35:
	;
	v99 = int32(100)
	if v99 <= v56 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = v99
	goto L38
L37:
	;
	v102 = v56
	goto L38
L38:
	;
	v103 = int32(0)
	if v103 < v102 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v106 = v102
	goto L41
L40:
	;
	v106 = v103
	goto L41
L41:
	;
	v107 = int32(1)
	v108 = v83 - v107
	v112 = l0
	v119 = l2
	v120 = v5
	v121 = v107
	v123 = v5
	v125 = v30 + int32(15)
	v126 = v5
	v134 = v5
	goto L42
L42:
	;
	v139 = int32(-1)
	v140 = v119 - l2
	if v82 <= v140 {
		v742 = v139
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v713 = v683
	v714 = v684
	v719 = v689
	goto L26
L44:
	;
	if v134 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v144 <= v140 {
		v742 = v139
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	v147 = v44 - v112
	v149 = base.B2i32(v147 < int32(4))
	if v147 < int32(4) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L47
L49:
	;
	if base.Ui32(v676) < base.Ui32(v44) {
		v112 = v676
		v119 = v683
		v120 = v684
		v121 = v685
		v123 = v687 << (uint(int32(1)) % 32)
		v125 = v689
		v126 = v690
		v134 = v698
		goto L42
	} else {
		goto L143
	}
L50:
	;
	if v123&int32(255) != 0 {
		goto L127
	} else {
		goto L128
	}
L51:
	;
	v162 = v146
	goto L53
L52:
	;
	v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+3)))
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+1)))
	v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+2)))
	v162 = v150 ^ (v151<<(uint(int32(4))%32) ^ v146<<(uint(int32(6))%32) ^ v157<<(uint(int32(2))%32))
	goto L53
L53:
	;
	v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v162&v108<<(uint(int32(1))%32)+v87))))
	if v167 == int32(0) {
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v171 = v167 << (uint(int32(4)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_consts[1274])))
	v175 = v112 - v174
	if int32(4094) < v175 {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v180 = int32(0)
	v185 = v174
	v187 = v180
	v192 = v98
	v200 = v180
	v203 = v171 + int32(4461808)
	v205 = v175
	goto L56
L56:
	;
	if v187 < int32(16) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	if v412 < int32(3) {
		goto L50
	} else {
		goto L103
	}
L58:
	;
	v410 = base.B2i32(v187 < v387)
	if v187 < v387 {
		goto L93
	} else {
		goto L94
	}
L59:
	;
	v387 = v147
	goto L58
L60:
	;
	v213 = v112
	v215 = v185
	v216 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v187) {
		goto L71
	} else {
		goto L72
	}
L63:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v239 != v240 {
		v387 = v216
		goto L58
	} else {
		goto L65
	}
L64:
	;
	goto L59
L65:
	;
	if base.Ui32(int32(272)) < base.Ui32(v216) {
		v387 = v216
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v244 = int32(1)
	v249 = v216 + v244
	if v249 != v147 {
		v213 = v213 + v244
		v215 = v215 + v244
		v216 = v249
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	if v313 != 0 {
		v387 = int32(0)
		goto L58
	} else {
		goto L86
	}
L69:
	;
	v313 = int32(0)
	goto L68
L70:
	;
	v287 = v282
	v288 = v283
	v289 = v284
	goto L80
L71:
	;
	if (v112|v185)&int32(3) != 0 {
		v282 = v112
		v283 = v185
		v284 = v187
		goto L70
	} else {
		goto L74
	}
L72:
	;
	v275 = v112
	v276 = v185
	v277 = v187
	goto L73
L73:
	;
	if v277 == int32(0) {
		goto L69
	} else {
		goto L79
	}
L74:
	;
	v259 = v112
	v260 = v185
	v261 = v187
	goto L75
L75:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v264 != v265 {
		v282 = v259
		v283 = v260
		v284 = v261
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v275 = v270
	v276 = v268
	v277 = v272
	goto L73
L77:
	;
	v267 = int32(4)
	v268 = v260 + v267
	v270 = v259 + v267
	v272 = v261 - v267
	if base.Ui32(int32(3)) < base.Ui32(v272) {
		v259 = v270
		v260 = v268
		v261 = v272
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v282 = v275
	v283 = v276
	v284 = v277
	goto L70
L80:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	if v292 == v293 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v313 = v292 - v293
	goto L68
L82:
	;
	v295 = int32(1)
	v300 = v289 - v295
	if v300 != 0 {
		v287 = v287 + v295
		v288 = v288 + v295
		v289 = v300
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	goto L69
L86:
	;
	v314 = v187 + v112
	if base.Ui32(v44) <= base.Ui32(v314) {
		v387 = v187
		goto L58
	} else {
		goto L87
	}
L87:
	;
	v318 = v314
	v320 = v185 + v187
	v321 = v187
	goto L88
L88:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	if v344 != v345 {
		v387 = v321
		goto L58
	} else {
		goto L90
	}
L89:
	;
	goto L59
L90:
	;
	if int32(272) < v321 {
		v387 = v321
		goto L58
	} else {
		goto L91
	}
L91:
	;
	v349 = int32(1)
	v354 = v318 + v349
	if base.Ui32(v354) < base.Ui32(v44) {
		v318 = v354
		v320 = v320 + v349
		v321 = v321 + v349
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	v411 = v205
	goto L95
L94:
	;
	v411 = v200
	goto L95
L95:
	;
	if v187 < v387 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v412 = v387
	goto L98
L97:
	;
	v412 = v187
	goto L98
L98:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v413 == int32(4461808) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	goto L57
L100:
	;
	if v192 <= v412 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v419 = base.I32_div_s(v192*v106, int32(-100))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v413)+12))
	v422 = v112 - v421
	if v422 < int32(4095) {
		v185 = v421
		v187 = v412
		v192 = v419 + v192
		v200 = v411
		v203 = v413
		v205 = v422
		goto L56
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	if v123&int32(255) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v437 = v120
	v438 = v123
	v439 = v125
	v440 = v119
	goto L106
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v120)
	v433 = int32(1)
	v437 = int32(0)
	v438 = v433
	v439 = v119
	v440 = v119 + v433
	goto L106
L106:
	;
	v442 = int32(base.Ui32(v411) >> (uint(int32(4)) % 32))
	if base.Ui32(int32(18)) <= base.Ui32(v412) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v461)
	*(*uint8)(unsafe.Add(mBase, uint32(v440)+1)) = uint8(v411)
	v464 = v112
	v469 = v412
	v473 = v121
	v478 = v126
	goto L111
L108:
	;
	v447 = v412 - int32(18)
	*(*uint8)(unsafe.Add(mBase, uint32(v440)+2)) = uint8(v447)
	v460 = v440 + int32(3)
	v461 = v442 | int32(15)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v460 = v440 + int32(2)
	v461 = v412 + int32(253) | v442&int32(240)
	goto L107
L111:
	;
	v491 = int32(*(*int8)(unsafe.Add(mBase, uint32(v464))))
	v492 = int32(4)
	v493 = v473 << (uint(v492) % 32)
	if v492 <= v44-v464 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v676 = v562
	v683 = v460
	v684 = v437 | v438
	v685 = v560
	v687 = v438
	v689 = v439
	v690 = v563
	v698 = v554
	goto L49
L113:
	;
	v499 = int32(*(*int8)(unsafe.Add(mBase, uint32(v464)+3)))
	v500 = int32(*(*int8)(unsafe.Add(mBase, uint32(v464)+1)))
	v506 = int32(*(*int8)(unsafe.Add(mBase, uint32(v464)+2)))
	v511 = v499 ^ (v500<<(uint(int32(4))%32) ^ v491<<(uint(int32(6))%32) ^ v506<<(uint(int32(2))%32))
	goto L115
L114:
	;
	v511 = v491
	goto L115
L115:
	;
	v512 = v511 & v108
	v513 = int32(1)
	v515 = v512<<(uint(v513)%32) + v87
	if v478&v513 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v540 = int32(*(*int16)(unsafe.Add(mBase, uint32(v515))))
	*(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1274]))) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1275]))) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1276]))) = int32(0)
	v546 = v540 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1277]))) = v546 + int32(4461808)
	*(*int32)(unsafe.Add(mBase, uint32(v546)+uint32(_consts[1276]))) = v493 + int32(4461808)
	*(*uint16)(unsafe.Add(mBase, uint32(v515))) = uint16(v473)
	v554 = int32(1)
	v557 = v473 + v554
	v559 = base.B2i32(int32(4096) < v557)
	if int32(4096) < v557 {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1277])))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1276])))
	if v521 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v520 == int32(0) {
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1275])))
	v531 = int32(base.Ui32(v520-int32(4461808)) >> (uint(int32(4)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v524<<(uint(int32(1))%32)+v87))) = uint16(v531)
	goto L118
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521))) = v520
	goto L118
L122:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v493)+uint32(_consts[1276])))
	*(*int32)(unsafe.Add(mBase, uint32(v520)+4)) = v536
	goto L116
L123:
	;
	v560 = v554
	goto L125
L124:
	;
	v560 = v557
	goto L125
L125:
	;
	v561 = int32(1)
	v562 = v464 + v561
	v563 = v559 | v478
	v565 = v469 - v561
	if v565 != 0 {
		v464 = v562
		v469 = v565
		v473 = v560
		v478 = v563
		goto L111
	} else {
		goto L126
	}
L126:
	;
	goto L112
L127:
	;
	v600 = v120
	v601 = v123
	v602 = v125
	v603 = v119
	goto L129
L128:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v120)
	v596 = int32(1)
	v600 = int32(0)
	v601 = v596
	v602 = v119
	v603 = v119 + v596
	goto L129
L129:
	;
	v604 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v604)
	v607 = v121 << (uint(int32(4)) % 32)
	if v147 < int32(4) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v622 = v604
	goto L132
L131:
	;
	v610 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+3)))
	v611 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+1)))
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112)+2)))
	v622 = v610 ^ (v611<<(uint(int32(4))%32) ^ v604<<(uint(int32(6))%32) ^ v617<<(uint(int32(2))%32))
	goto L132
L132:
	;
	v623 = v622 & v108
	v624 = int32(1)
	v626 = v623<<(uint(v624)%32) + v87
	if v126&v624 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v651 = int32(1)
	v653 = int32(*(*int16)(unsafe.Add(mBase, uint32(v626))))
	*(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1274]))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1275]))) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1276]))) = int32(0)
	v659 = v653 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1277]))) = v659 + int32(4461808)
	*(*int32)(unsafe.Add(mBase, uint32(v659)+uint32(_consts[1276]))) = v607 + int32(4461808)
	*(*uint16)(unsafe.Add(mBase, uint32(v626))) = uint16(v121)
	v669 = v121 + v651
	v671 = base.B2i32(int32(4096) < v669)
	if int32(4096) < v669 {
		goto L140
	} else {
		goto L141
	}
L134:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1277])))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1276])))
	if v632 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v631 == int32(0) {
		goto L133
	} else {
		goto L139
	}
L136:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1275])))
	v642 = int32(base.Ui32(v631-int32(4461808)) >> (uint(int32(4)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v635<<(uint(int32(1))%32)+v87))) = uint16(v642)
	goto L135
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = v631
	goto L135
L139:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v607)+uint32(_consts[1276])))
	*(*int32)(unsafe.Add(mBase, uint32(v631)+4)) = v647
	goto L133
L140:
	;
	v672 = v651
	goto L142
L141:
	;
	v672 = v669
	goto L142
L142:
	;
	v676 = v112 + int32(1)
	v683 = v603 + v651
	v684 = v600
	v685 = v672
	v687 = v601
	v689 = v602
	v690 = v671 | v126
	v698 = v134
	goto L49
L143:
	;
	goto L43
L144:
	;
	v737 = int32(-1)
	goto L146
L145:
	;
	v737 = v735
	goto L146
L146:
	;
	v742 = v737
	goto L1
}
func F_phraseto_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1176), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_pipe(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_pipe(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_pkt_stream_process(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v12 != 0 {
		v68 = int32(-12)
		m.G0 = v9 + int32(16)
		return v68
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v13 == v4 {
			v15 = int32(238)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v15)
			v54 = v9 + int32(9)
		} else {
			if v4 <= int32(191) {
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v4)
				v51 = v9 + int32(9)
			} else {
				if base.Ui32(v4) <= base.Ui32(int32(8383)) {
					v27 = v4 - int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+9)) = uint8(v27)
					v32 = int32(base.Ui32(v27)>>(uint(int32(8))%32)) + int32(-64)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v32)
					v51 = v9 + int32(10)
				} else {
					v36 = int32(255)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v4)
					v40 = int32(base.Ui32(v4) >> (uint(int32(8)) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+11)) = uint8(v40)
					v43 = int32(base.Ui32(v4) >> (uint(int32(16)) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+10)) = uint8(v43)
					v46 = int32(base.Ui32(v4) >> (uint(int32(24)) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+9)) = uint8(v46)
					v51 = v9 + int32(13)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
			v54 = v51
		}
		v56 = v9 + int32(8)
		v60 = F_pushf_write(m, l0, v56, v54-v56)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			if v60 < int32(0) {
				v68 = v60
				m.G0 = v9 + int32(16)
				return v68
			} else {
				v66 = F_pushf_write(m, l0, l2, v4)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v68 = v66
					m.G0 = v9 + int32(16)
					return v68
				}
			}
		}
	}
}
func F_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 float64
	_ = v96
	var v104 float64
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 float64
	_ = v122
	var v128 float64
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v196 float64
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v280 float64
	_ = v280
	var v282 float64
	_ = v282
	var v286 float64
	_ = v286
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v293 float64
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v312 float64
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 float64
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 float64
	_ = v347
	var v348 float64
	_ = v348
	var v351 int32
	_ = v351
	var v352 float64
	_ = v352
	var v353 float64
	_ = v353
	var v355 float64
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v378 float64
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 float64
	_ = v386
	var v387 float64
	_ = v387
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v405 float64
	_ = v405
	var v408 int32
	_ = v408
	var v410 float64
	_ = v410
	var v411 float64
	_ = v411
	var v414 float64
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int64
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
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
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 float64
	_ = v589
	var v594 float64
	_ = v594
	var v598 int32
	_ = v598
	var v602 float64
	_ = v602
	var v607 float64
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 float64
	_ = v616
	var v621 float64
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v672 int32
	_ = v672
	var v673 int64
	_ = v673
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int64
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v707 int32
	_ = v707
	v12 = float64(0)
	v15 = *(*int32)(unsafe.Add(mBase, _consts[494]))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v672)+16))
	v677 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v677 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L2:
	;
	v16 = m.T0[v15].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = F_palloc0(m, int32(92))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v672 = v16
	goto L1
L7:
	;
	v672 = v526
	goto L1
L8:
	;
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(266)
	v38 = F__emscripten_memset_bulkmem(m, v26+int32(8), base.I32_extend8_s(v28), int32(74))
	mBase = m.M
	goto L9
L9:
	;
	if l2&int32(2048) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v87 = int32(0)
	v89 = v83 & base.B2i32(v86 != v87)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+83)) = uint8(v89)
	if l2&int32(256) == v87 {
		v104 = v12
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+82)) = uint8(v79)
	v81 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+84)) = uint8(v81)
	v83 = int32(0)
	goto L10
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[92])))
	if v44 != int32(1) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 != int32(1) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	if v50 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	if v52 <= int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if int32(0) <= v56 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v59 = m.G0
	v61 = v59 - int32(16)
	m.G0 = v61
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = int32(0)
	v65 = int32(30067)
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+8)) = uint16(v65)
	v69 = F_max_parallel_hazard_walker(m, l0, v61+int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+8)))
	m.G0 = v61 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+84)) = uint8(v71)
	v77 = base.B2i32(v71 != int32(117))
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+82)) = uint8(v77)
	v83 = v77
	goto L10
L19:
	;
	v106 = int32(0)
	v109 = F_subquery_planner(m, v26, l0, v106, v106, v104, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L23
	}
L20:
	;
	v96 = *(*float64)(unsafe.Add(mBase, _consts[497]))
	if base.F64_ge(v96, float64(1)) != 0 {
		v104 = v12
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if base.F64_le(v96, float64(0)) == int32(0) {
		v104 = v96
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v104 = float64(1e-10)
	goto L19
L23:
	;
	v113 = F_fetch_upper_rel(m, v109, int32(7), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	if base.F64_le(v104, float64(0)) != 0 {
		v224 = v115
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v230 = F_create_plan(m, v109, v224)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L59
	}
L26:
	;
	if base.F64_ge(v104, float64(1)) == int32(0) {
		v128 = v104
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+32))
	if v130 == int32(0) {
		v224 = v115
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v115)+32))
	if base.F64_gt(v122, float64(0)) == int32(0) {
		v128 = v104
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v128 = base.F64_div(v104, v122)
	goto L27
L30:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v133 <= int32(0) {
		v224 = v115
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v137 = int32(0)
	v143 = v133
	v144 = v115
	goto L32
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v137<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	if v155 != 0 {
		v212 = v143
		v213 = v144
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v224 = v213
	goto L25
L34:
	;
	v215 = v137 + int32(1)
	if v215 < v212 {
		v137 = v215
		v143 = v212
		v144 = v213
		goto L32
	} else {
		goto L58
	}
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	if v154 == v156 {
		v212 = v143
		v213 = v144
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v144)+40))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+40))
	if v162 != v163 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v207 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L38:
	;
	if v162 < v163 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if base.F64_le(v128, float64(0))|base.F64_ge(v128, float64(1)) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v168 = int32(-1)
	goto L43
L42:
	;
	v168 = int32(1)
	goto L43
L43:
	;
	v207 = v168
	goto L37
L44:
	;
	v207 = v202
	goto L37
L45:
	;
	v174 = int32(-1)
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v144)+56))
	v176 = *(*float64)(unsafe.Add(mBase, uint32(v154)+56))
	if base.F64_lt(v175, v176) != 0 {
		v202 = v174
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v144)+56))
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v144)+48))
	v191 = base.F64_add(base.F64_mul(v128, base.F64_sub(v187, v188)), v188)
	v192 = *(*float64)(unsafe.Add(mBase, uint32(v154)+56))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v154)+48))
	v196 = base.F64_add(base.F64_mul(v128, base.F64_sub(v192, v193)), v193)
	if base.F64_lt(v191, v196) != 0 {
		v202 = int32(-1)
		goto L44
	} else {
		goto L54
	}
L48:
	;
	if base.F64_gt(v175, v176) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v207 = int32(1)
	goto L37
L50:
	;
	goto L51
L51:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v144)+48))
	v181 = *(*float64)(unsafe.Add(mBase, uint32(v154)+48))
	if base.F64_lt(v180, v181) != 0 {
		v202 = v174
		goto L44
	} else {
		goto L52
	}
L52:
	;
	if base.F64_gt(v180, v181) != 0 {
		v202 = int32(1)
		goto L44
	} else {
		goto L53
	}
L53:
	;
	v207 = int32(0)
	goto L37
L54:
	;
	v202 = base.F64_lt(v196, v191)
	goto L44
L55:
	;
	v210 = v144
	goto L57
L56:
	;
	v210 = v154
	goto L57
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v212 = v211
	v213 = v210
	goto L34
L58:
	;
	goto L33
L59:
	;
	if l2&int32(2) == int32(0) {
		v240 = v230
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	if v242 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v236 = F_ExecSupportsBackwardScan(m, v230)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	if v236 != 0 {
		v240 = v230
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v238 = F_materialize_finished_plan(m, v230)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v240 = v238
	goto L60
L65:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	if v423 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L66:
	;
	v420 = v240
	goto L65
L67:
	;
	goto L68
L68:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+37)))
	if v245 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v420 = v240
	goto L65
L70:
	;
	goto L71
L71:
	;
	if v242 != int32(2) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v254 = F_palloc0(m, int32(88))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v240)+60))
	if v250 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v420 = v240
	goto L65
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = int32(368)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+80)) = uint8(v259)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+72)) = v259
	v263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+56)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v254)+52)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v254)+48)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v254)+44)) = v258
	v270 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	*(*uint8)(unsafe.Add(mBase, uint32(v254)+81)) = uint8(base.B2i32(v270 == int32(2)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v240)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+60)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v240)+60)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v254)+76)) = int32(-1)
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v240)+8))
	v282 = *(*float64)(unsafe.Add(mBase, _consts[498]))
	*(*float64)(unsafe.Add(mBase, uint32(v254)+8)) = base.F64_add(v280, v282)
	v286 = *(*float64)(unsafe.Add(mBase, _consts[499]))
	v287 = *(*float64)(unsafe.Add(mBase, uint32(v240)+24))
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v240)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v254)+16)) = base.F64_add(base.F64_mul(v286, v287), base.F64_add(v282, v289))
	v293 = *(*float64)(unsafe.Add(mBase, uint32(v240)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v254)+24)) = v293
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v240)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v254)+36)) = uint16(v263)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+32)) = v295
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v254)+60))
	v312 = float64(0)
	if v299 == v263 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v410 = *(*float64)(unsafe.Add(mBase, uint32(v240)+8))
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v23)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v240)+8)) = base.F64_sub(v410, v411)
	v414 = *(*float64)(unsafe.Add(mBase, uint32(v240)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v240)+16)) = base.F64_sub(v414, v411)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+83)) = uint8(v418)
	v420 = v254
	goto L65
L77:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v23+int32(8)))) = v405
	v408 = v398 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(7)))) = uint8(v408)
	goto L76
L78:
	;
	v398 = v263
	v405 = v312
	goto L77
L79:
	;
	goto L80
L80:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if v315 <= int32(0) {
		v398 = v263
		v405 = v312
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v318 = int32(0)
	if v318 < v315 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v321 = v315
	goto L84
L83:
	;
	v321 = v318
	goto L84
L84:
	;
	v322 = int32(1)
	if v315 == v322 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v321&v322 == int32(0) {
		v398 = v371
		v405 = v378
		goto L77
	} else {
		goto L92
	}
L86:
	;
	v370 = int32(0)
	v371 = v263
	v378 = v312
	goto L85
L87:
	;
	goto L88
L88:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v334 = int32(0)
	v335 = v263
	v338 = v263
	v342 = v312
	goto L89
L89:
	;
	v343 = int32(2)
	v345 = v329 + v334<<(uint(v343)%32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v346)+56))
	v348 = *(*float64)(unsafe.Add(mBase, uint32(v346)+64))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v352 = *(*float64)(unsafe.Add(mBase, uint32(v351)+56))
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v351)+64))
	v355 = base.F64_add(base.F64_add(v342, base.F64_add(v347, v348)), base.F64_add(v352, v353))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+38)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+38)))
	v361 = v356&v357 ^ int32(1) | v335
	v363 = v334 + v343
	v365 = v338 + v343
	if v365 != v321&int32(2147483646) {
		v334 = v363
		v335 = v361
		v338 = v365
		v342 = v355
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v370 = v363
	v371 = v361
	v378 = v355
	goto L85
L91:
	;
	goto L90
L92:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381+v370<<(uint(int32(2))%32))))
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v385)+56))
	v387 = *(*float64)(unsafe.Add(mBase, uint32(v385)+64))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+38)))
	v398 = v390 ^ int32(1) | v371
	v405 = base.F64_add(v378, base.F64_add(v386, v387))
	goto L77
L93:
	;
	v485 = F_set_plan_references(m, v109, v420)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L108
	}
L94:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v429 = int32(0)
	goto L95
L95:
	;
	v442 = int32(0)
	if v427 == v442 {
		v452 = v442
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if v426 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v446 <= v429 {
		v452 = int32(0)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v452 = v448 + v429<<(uint(int32(2))%32)
	goto L97
L100:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	F_SS_finalize_plan(m, v466, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L107
	}
L101:
	;
	F_SS_finalize_plan(m, v109, v420)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L106
	}
L102:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	if v455 <= v429 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	if v452 == int32(0) {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v426)+12))
	v462 = v459 + v429<<(uint(int32(2))%32)
	if v462 != 0 {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	goto L93
L107:
	;
	v429 = v429 + int32(1)
	goto L95
L108:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v490 = int32(0)
	goto L109
L109:
	;
	v503 = int32(0)
	if v488 == v503 {
		v513 = v503
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v487 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v507 <= v490 {
		v513 = int32(0)
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v513 = v509 + v490<<(uint(int32(2))%32)
	goto L111
L114:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v654 = F_set_plan_references(m, v652, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L5
	} else {
		goto L140
	}
L115:
	;
	v526 = F_palloc0(m, int32(104))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L120
	}
L116:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v516 <= v490 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	if v513 == int32(0) {
		goto L115
	} else {
		goto L118
	}
L118:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v523 = v520 + v490<<(uint(int32(2))%32)
	if v523 != 0 {
		goto L114
	} else {
		goto L119
	}
L119:
	;
	goto L115
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = int32(330)
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v530
	v532 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v526)+8)) = v532
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+24)) = uint8(base.B2i32(v534 != int32(0)))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+25)) = uint8(v538)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+26)) = uint8(v540)
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+80)))
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+27)) = uint8(v542)
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+81)))
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+28)) = uint8(v544)
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+83)))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+36)) = v485
	*(*uint8)(unsafe.Add(mBase, uint32(v526)+29)) = uint8(v546)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+40)) = v549
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+44)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v555 = F_bms_difference(m, v553, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+48)) = v555
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+52)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+56)) = v560
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+60)) = v562
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+64)) = v564
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+68)) = v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+72)) = v568
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+76)) = v570
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v26)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+80)) = v572
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+84)) = v574
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+88)) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+92)) = v578
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+96)) = v580
	v585 = int32(*(*uint8)(unsafe.Add(mBase, _consts[500])))
	if v585 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	if v646 != 0 {
		goto L136
	} else {
		goto L137
	}
L123:
	;
	v589 = *(*float64)(unsafe.Add(mBase, _consts[501]))
	if base.F64_ge(v589, float64(0)) == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v594 = *(*float64)(unsafe.Add(mBase, uint32(v485)+16))
	if base.F64_gt(v594, v589) == int32(0) {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v598 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v598
	v602 = *(*float64)(unsafe.Add(mBase, _consts[502]))
	if base.F64_ge(v602, float64(0)) == int32(0) {
		v614 = v598
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v616 = *(*float64)(unsafe.Add(mBase, _consts[503]))
	if base.F64_ge(v616, float64(0)) == int32(0) {
		v628 = v614
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v607 = *(*float64)(unsafe.Add(mBase, uint32(v485)+16))
	if base.F64_gt(v607, v602) == int32(0) {
		v614 = v598
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v611 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v611
	v614 = v611
	goto L126
L129:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, _consts[504])))
	if v630 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v621 = *(*float64)(unsafe.Add(mBase, uint32(v485)+16))
	if base.F64_gt(v621, v616) == int32(0) {
		v628 = v614
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v626 = v614 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v626
	v628 = v626
	goto L129
L132:
	;
	v634 = v628 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v634
	v636 = v634
	goto L134
L133:
	;
	v636 = v628
	goto L134
L134:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, _consts[505])))
	if v638 != int32(1) {
		goto L122
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526)+32)) = v636 | int32(16)
	goto L122
L136:
	;
	F_DestroyPartitionDirectory(m, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	m.G0 = v23 + int32(16)
	goto L7
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v654
	v490 = v490 + int32(1)
	goto L109
L141:
	;
	return v672
L142:
	;
	goto L141
L143:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	if v681 != int32(1) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v686 = *(*int64)(unsafe.Add(mBase, uint32(v677)+400))
	if int32(1)&base.B2i32(v686 != int64(0)) != 0 {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	v690 = int32(4438516)
	v692 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v693 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v692 + v693
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v677)))
	*(*int32)(unsafe.Add(mBase, uint32(v677))) = v696 + v693
	*(*int64)(unsafe.Add(mBase, uint32(v677)+400)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v677))) = v696 + int32(2)
	v707 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v707 - v693
	goto L142
}
func F_pop_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 float64
	_ = v127
	var v130 int32
	_ = v130
	switch l1 - int32(9) {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v7 + int32(4)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
		return
	case 1:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13 + int32(4)
		v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v17
		return
	case 2:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19 + int32(4)
		v23 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
		return
	case 3:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v41 = (v37 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41 + int32(8)
		v45 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v45
		return
	case 4:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25 + int32(4)
		v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v25))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v29
		return
	case 5:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31 + int32(4)
		v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v35
		return
	case 6:
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v47 + int32(4)
		v51 = int64(*(*int16)(unsafe.Add(mBase, uint32(v47))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v51
		return
	case 7:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v53 + int32(4)
		v57 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v53))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v57
		return
	case 8:
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v59 + int32(4)
		v63 = int64(*(*int8)(unsafe.Add(mBase, uint32(v59))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v63
		return
	case 9:
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65 + int32(4)
		v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v69
		return
	case 10:
		v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v75 = (v71 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v75 + int32(8)
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v79
		return
	case 11:
		v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v81 + int32(4)
		v85 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v81))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v85
		return
	case 12:
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v91 = (v87 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v91 + int32(8)
		v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v95
		return
	case 13:
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v101 = (v97 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v101 + int32(8)
		v105 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
		return
	case 14:
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v107 + int32(4)
		v111 = int64(*(*int32)(unsafe.Add(mBase, uint32(v107))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v111
		return
	case 15:
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v113 + int32(4)
		v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v113))))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v117
		return
	case 16:
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v123 = (v119 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v123 + int32(8)
		v127 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
		*(*float64)(unsafe.Add(mBase, uint32(l0))) = v127
		return
	case 17:
		m.T0[l3].(func(*base.Module, int32, int32))(m, l0, l2)
		mBase = m.M
		v130 = m.ExcPending
		if v130 != 0 {
			return
		} else {
			return
		}
	default:
		return
	}
}
func F_preprocess_aggrefs(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_preprocess_aggrefs_walker(m, l1, l0)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_printTypmod(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l2 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v14 = F_psprintf(m, int32(638921), v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v28 = v14
			m.G0 = v7 + int32(32)
			return v28
		}
	} else {
		v19 = F_OidFunctionCall1Coll(m, l2, int32(0), l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			v26 = F_psprintf(m, int32(166263), v7+int32(16))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = v26
				m.G0 = v7 + int32(32)
				return v28
			}
		}
	}
}
func F_processCASbits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v17)
	goto L6
L5:
	;
	goto L6
L6:
	;
	if l6 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v19 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v19)
	goto L9
L8:
	;
	goto L9
L9:
	;
	if l5 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v21)
	goto L12
L11:
	;
	goto L12
L12:
	;
	if l0&int32(10) != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L43
	} else {
		goto L69
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L43
	} else {
		goto L64
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L43
	} else {
		goto L59
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L43
	} else {
		goto L54
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L43
	} else {
		goto L49
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L43
	} else {
		goto L44
	}
L19:
	;
	if l3 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l0&int32(8) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v27)
	goto L21
L23:
	;
	if l4 == int32(0) {
		goto L17
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if l0&int32(16) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v33)
	goto L25
L27:
	;
	if l6 == int32(0) {
		goto L16
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l0&int32(32) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v39)
	goto L29
L31:
	;
	if l7 == int32(0) {
		goto L15
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l0&int32(64) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v45)
	goto L33
L35:
	;
	if l0&int32(128) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if l5 == int32(0) {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v53)
	if l6 == v53 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v57)
	goto L35
L39:
	;
	if l5 == int32(0) {
		goto L13
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	m.G0 = v13 + int32(96)
	return
L42:
	;
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v63)
	goto L41
L43:
	;
	return
L44:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = l2
	F_errmsg(m, int32(518543), v13+int32(80))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(25284), int32(19471), int32(116246))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L43
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l2
	F_errmsg(m, int32(518543), v13-int32(-64))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(25284), int32(19484), int32(116246))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
	F_errmsg(m, int32(520775), v13+int32(48))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L43
	} else {
		goto L56
	}
L56:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(25284), int32(19497), int32(116246))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L43
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l2
	F_errmsg(m, int32(498466), v13+int32(32))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L43
	} else {
		goto L61
	}
L61:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(25284), int32(19510), int32(116246))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L43
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L43
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	F_errmsg(m, int32(521476), v13+int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L43
	} else {
		goto L66
	}
L66:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L43
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(25284), int32(19523), int32(116246))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L43
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(521435), v13)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L43
	} else {
		goto L71
	}
L71:
	;
	F_scanner_errposition(m, l1, l8)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L43
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(25284), int32(19545), int32(116246))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L43
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_process_equivalence(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int64
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v743 int32
	_ = v743
	v4 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v29 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v25 + int32(16)
	return v743
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v32 == int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v743 = int32(0)
	goto L1
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v52 = F_exprType(m, v47)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v46 = v4
	v47 = v4
	goto L5
L7:
	;
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v42 < int32(2) {
		v46 = v4
		v47 = v41
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v46 = v45
	v47 = v41
	goto L5
L10:
	;
	return int32(0)
L11:
	;
	v56 = F_canonicalize_ec_expression(m, v47, v52, v36)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = F_exprType(m, v46)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v60 = F_canonicalize_ec_expression(m, v46, v58, v36)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v62 = F_equal(m, v56, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	if v62 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_set_opfuncid(m, v28)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_op_input_types(m, v49, v25+int32(12), v25+int32(8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L26
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v67 = F_func_strict(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	if v67 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v743 = int32(0)
	goto L1
L22:
	;
	goto L23
L23:
	;
	v73 = F_palloc0(m, int32(20))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = int32(-1)
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+12)) = uint8(v77)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(52)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+11)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+12)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+10)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	v93 = F_make_restrictinfo(m, l0, v73, v85, v86, v87, v88, v89, v77, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v93
	v743 = v77
	goto L1
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v103 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v743 = int32(1)
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v276
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v288
	goto L28
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L10
	} else {
		goto L162
	}
L31:
	;
	v565 = F_palloc0(m, int32(60))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L10
	} else {
		goto L147
	}
L32:
	;
	v106 = int32(-1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v276 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L34:
	;
	v276 = int32(0)
	v278 = v4
	v288 = v4
	v289 = v4
	v292 = v106
	goto L33
L35:
	;
	goto L36
L36:
	;
	v111 = int32(0)
	v114 = v111
	v118 = v111
	v120 = v4
	v130 = v4
	v131 = v4
	v134 = v106
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v114<<(uint(int32(2))%32))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+41)))
	if v140 != 0 {
		v250 = v118
		v252 = v120
		v262 = v130
		v263 = v131
		v266 = v134
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v276 = v250
	v278 = v252
	v288 = v262
	v289 = v263
	v292 = v266
	goto L33
L39:
	;
	v268 = v114 + int32(1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v268 < v269 {
		v114 = v268
		v118 = v250
		v120 = v252
		v130 = v262
		v131 = v263
		v134 = v266
		goto L37
	} else {
		goto L70
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
	if v36 != v141 {
		v250 = v118
		v252 = v120
		v262 = v130
		v263 = v131
		v266 = v134
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v144 = F_equal(m, v102, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	if v144 == int32(0) {
		v250 = v118
		v252 = v120
		v262 = v130
		v263 = v131
		v266 = v134
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	if v148 == int32(0) {
		v226 = v118
		v228 = v120
		v238 = v130
		v239 = v131
		v242 = v134
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v226 == int32(0) {
		v250 = v226
		v252 = v228
		v262 = v238
		v263 = v239
		v266 = v242
		goto L39
	} else {
		goto L68
	}
L45:
	;
	v151 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v152 <= v151 {
		v226 = v118
		v228 = v120
		v238 = v130
		v239 = v131
		v242 = v134
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v158 = v151
	v160 = v118
	v162 = v120
	v172 = v130
	v173 = v131
	v176 = v134
	goto L47
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v158<<(uint(int32(2))%32))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+12)))
	if v182 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v226 = v211
	v228 = v212
	v238 = v213
	v239 = v214
	v242 = v216
	goto L44
L49:
	;
	v218 = v158 + int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v218 < v219 {
		v158 = v218
		v160 = v211
		v162 = v212
		v172 = v213
		v173 = v214
		v176 = v216
		goto L47
	} else {
		goto L67
	}
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v185 != l2 {
		v211 = v160
		v212 = v162
		v213 = v172
		v214 = v173
		v216 = v176
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v160 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L52
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	if v200 != v201 {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	if v162 != 0 {
		v211 = v160
		v212 = v162
		v213 = v172
		v214 = v173
		v216 = v176
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	if v187 != v188 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v191 = F_equal(m, v56, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	if v191 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	if v162 == int32(0) {
		v197 = v139
		v198 = v181
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v226 = v139
	v228 = v162
	v238 = v172
	v239 = v181
	v242 = v176
	goto L44
L61:
	;
	v197 = v160
	v198 = v173
	goto L54
L62:
	;
	v211 = v197
	v212 = int32(0)
	v213 = v172
	v214 = v198
	v216 = v176
	goto L49
L63:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v204 = F_equal(m, v60, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	if v204 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v208 = int32(0)
	if v197 == v208 {
		v211 = v208
		v212 = v139
		v213 = v181
		v214 = v198
		v216 = v114
		goto L49
	} else {
		goto L66
	}
L66:
	;
	v226 = v197
	v228 = v139
	v238 = v181
	v239 = v198
	v242 = v114
	goto L44
L67:
	;
	goto L48
L68:
	;
	if v228 != 0 {
		v276 = v226
		v278 = v228
		v288 = v238
		v289 = v239
		v292 = v242
		goto L33
	} else {
		goto L69
	}
L69:
	;
	v250 = v226
	v252 = v228
	v262 = v238
	v263 = v239
	v266 = v242
	goto L39
L70:
	;
	goto L38
L71:
	;
	if v276 != 0 {
		goto L117
	} else {
		goto L118
	}
L72:
	;
	if v278 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v276 == v278 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v299 = F_lappend(m, v298, v27)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L10
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v312 == int32(1) {
		goto L30
	} else {
		goto L84
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v276)+48))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v302) < base.Ui32(v303) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v305 = v302
	goto L80
L79:
	;
	v305 = v303
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+48)) = v305
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v308) < base.Ui32(v307) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v310 = v307
	goto L83
L82:
	;
	v310 = v308
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+52)) = v310
	goto L29
L84:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v276)+16))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v317 = F_list_concat(m, v315, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v317
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v278)+24))
	v322 = F_list_concat(m, v320, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v322
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v276)+28))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v278)+28))
	v327 = F_list_concat(m, v325, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+28)) = v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v276)+32))
	if v330 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v276)+36))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v278)+36))
	v396 = F_bms_join(m, v394, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L10
	} else {
		goto L96
	}
L89:
	;
	if v326 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v335 <= int32(0) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v342 = int32(0)
	goto L92
L92:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361+v342<<(uint(int32(2))%32))))
	F_ec_add_clause_to_derives_hash(m, v276, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L94
	}
L93:
	;
	goto L88
L94:
	;
	v369 = v342 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v369 < v370 {
		v342 = v369
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+36)) = v396
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+40)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+40)))
	v401 = v399 | v400
	*(*uint8)(unsafe.Add(mBase, uint32(v276)+40)) = uint8(v401)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v276)+48))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v278)+48))
	if base.Ui32(v403) < base.Ui32(v404) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v406 = v403
	goto L99
L98:
	;
	v406 = v404
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+48)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	if base.Ui32(v409) < base.Ui32(v408) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v411 = v408
	goto L102
L101:
	;
	v411 = v409
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+52)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v278)+56)) = v276
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v415 = F_list_delete_nth_cell(m, v414, v292)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v415
	v418 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v418
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v278)+28))
	F_list_free(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+28)) = int32(0)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v278)+32))
	if v427 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+20))
	F_pfree(m, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+36)) = int32(0)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v438 = F_lappend(m, v437, v27)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L10
	} else {
		goto L110
	}
L108:
	;
	F_pfree(m, v427)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+32)) = int32(0)
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v438
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v276)+48))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v441) < base.Ui32(v442) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v444 = v441
	goto L113
L112:
	;
	v444 = v442
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+48)) = v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v447) < base.Ui32(v446) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v449 = v446
	goto L116
L115:
	;
	v449 = v447
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+52)) = v449
	goto L29
L117:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v453 = F_palloc0(m, int32(28))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L10
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v278 == int32(0) {
		goto L31
	} else {
		goto L133
	}
L120:
	;
	v455 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+24)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v453)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v453)+16)) = v451
	*(*uint16)(unsafe.Add(mBase, uint32(v453)+12)) = uint16(v455)
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = int32(274)
	if v50 == v455 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v467 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v453)+12)) = uint8(v467)
	*(*uint8)(unsafe.Add(mBase, uint32(v276)+40)) = uint8(v467)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v276)+16))
	v472 = F_lappend(m, v471, v453)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v472
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v276)+36))
	v476 = F_bms_add_members(m, v475, v50)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L10
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+36)) = v476
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v276)+24))
	v480 = F_lappend(m, v479, v27)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+24)) = v480
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v276)+48))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v483) < base.Ui32(v484) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v486 = v483
	goto L129
L128:
	;
	v486 = v484
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+48)) = v486
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v489) < base.Ui32(v488) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v491 = v488
	goto L132
L131:
	;
	v491 = v489
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+52)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v453
	goto L28
L133:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v498 = F_palloc0(m, int32(28))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	v500 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v498)+24)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v498)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v498)+16)) = v496
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+12)) = uint16(v500)
	*(*int32)(unsafe.Add(mBase, uint32(v498)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v498)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v498))) = int32(274)
	if v51 == v500 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+12)) = uint8(v512)
	*(*uint8)(unsafe.Add(mBase, uint32(v278)+40)) = uint8(v512)
	goto L137
L136:
	;
	goto L137
L137:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v517 = F_lappend(m, v516, v498)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L10
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v517
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v278)+36))
	v521 = F_bms_add_members(m, v520, v51)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+36)) = v521
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v278)+24))
	v525 = F_lappend(m, v524, v27)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+24)) = v525
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v278)+48))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v528) < base.Ui32(v529) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v531 = v528
	goto L143
L142:
	;
	v531 = v529
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+48)) = v531
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if base.Ui32(v534) < base.Ui32(v533) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v536 = v533
	goto L146
L145:
	;
	v536 = v534
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+52)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v278
	goto L27
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+20)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v565)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v565)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v565)+4)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = int32(273)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v27
	v578 = F_list_make1_impl(m, int32(1), v25)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	v580 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v565)+44)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v565)+24)) = v578
	v583 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v565)+28)) = v583
	*(*int64)(unsafe.Add(mBase, uint32(v565)+35)) = v583
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v565)+56)) = v580
	*(*int32)(unsafe.Add(mBase, uint32(v565)+52)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v565)+48)) = v587
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v594 = F_palloc0(m, int32(28))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	v596 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v594)+24)) = v596
	*(*int32)(unsafe.Add(mBase, uint32(v594)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v594)+16)) = v592
	*(*uint16)(unsafe.Add(mBase, uint32(v594)+12)) = uint16(v596)
	*(*int32)(unsafe.Add(mBase, uint32(v594)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v594)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = int32(274)
	if v51 == v596 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v608 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v594)+12)) = uint8(v608)
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+40)) = uint8(v608)
	goto L152
L151:
	;
	goto L152
L152:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v613 = F_lappend(m, v612, v594)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+16)) = v613
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v565)+36))
	v617 = F_bms_add_members(m, v616, v51)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L10
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+36)) = v617
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v622 = F_palloc0(m, int32(28))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	v624 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v622)+24)) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v622)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v622)+16)) = v620
	*(*uint16)(unsafe.Add(mBase, uint32(v622)+12)) = uint16(v624)
	*(*int32)(unsafe.Add(mBase, uint32(v622)+8)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v622)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = int32(274)
	if v50 == v624 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v636 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v622)+12)) = uint8(v636)
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+40)) = uint8(v636)
	goto L158
L157:
	;
	goto L158
L158:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v641 = F_lappend(m, v640, v622)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+16)) = v641
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v565)+36))
	v645 = F_bms_add_members(m, v644, v50)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L10
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v565)+36)) = v645
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v649 = F_lappend(m, v648, v565)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L10
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v649
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v622
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v565
	goto L27
L162:
	;
	F_errmsg_internal(m, int32(151896), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(472266), int32(400), int32(398753))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_process_sublinks_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 float64
	_ = v91
	var v97 float64
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
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
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
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
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 float64
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v428 float64
	_ = v428
	var v430 int32
	_ = v430
	var v434 float64
	_ = v434
	var v435 float64
	_ = v435
	var v438 float64
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v24
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v644
L2:
	;
	v644 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v29 - int32(9) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		goto L5
	case 12:
		goto L6
	case 13:
		goto L10
	default:
		goto L11
	}
L5:
	;
	v637 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v637)
	v642 = F_expression_tree_mutator_impl(m, l0, int32(846), v22+int32(16))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L222
	}
L6:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v483 {
	case 0:
		goto L187
	case 1:
		goto L186
	default:
		goto L5
	}
L7:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v480 {
		v644 = l0
		goto L1
	} else {
		goto L184
	}
L8:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v477 == int32(0) {
		goto L5
	} else {
		goto L183
	}
L9:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v474 == int32(0) {
		goto L5
	} else {
		goto L182
	}
L10:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = F_process_sublinks_mutator(m, v41, v22+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v29 == int32(61) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v29 != int32(319) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v36 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v644 = l0
	goto L1
L15:
	;
	return int32(0)
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = F_copyObjectImpl(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = F_simplify_EXISTS_query(m, v51, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v59 = v3
	goto L20
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if base.Ui32(v50) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v59 = v57
	goto L20
L22:
	;
	v66 = float64(0.5)
	goto L24
L23:
	;
	v66 = float64(0)
	goto L24
L24:
	;
	if v50 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = v66
	goto L27
L26:
	;
	v68 = float64(1)
	goto L27
L27:
	;
	v70 = F_subquery_planner(m, v60, v53, v51, int32(0), v68, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v73
	v77 = F_fetch_upper_rel(m, v70, int32(7), v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
	if base.F64_le(v68, float64(0)) != 0 {
		v131 = v84
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v136 = F_create_plan(m, v70, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L15
	} else {
		goto L47
	}
L31:
	;
	goto L30
L32:
	;
	if base.F64_ge(v68, float64(1)) == int32(0) {
		v97 = v68
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+32))
	if v99 == int32(0) {
		v131 = v84
		goto L31
	} else {
		goto L36
	}
L34:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v84)+32))
	if base.F64_gt(v91, float64(0)) == int32(0) {
		v97 = v68
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v97 = base.F64_div(v68, v91)
	goto L33
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 <= int32(0) {
		v131 = v84
		goto L31
	} else {
		goto L37
	}
L37:
	;
	v107 = v84
	v109 = int32(0)
	goto L38
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v109<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	if v117 != 0 {
		v124 = v107
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v131 = v124
	goto L31
L40:
	;
	v126 = v109 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v126 < v127 {
		v107 = v124
		v109 = v126
		goto L38
	} else {
		goto L46
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
	if v116 == v118 {
		v124 = v107
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v120 = F_compare_fractional_path_costs(m, v107, v116, v97)
	mBase = m.M
	if v120 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v123 = v107
	goto L45
L44:
	;
	v123 = v116
	goto L45
L45:
	;
	v124 = v123
	goto L40
L46:
	;
	goto L39
L47:
	;
	v141 = F_build_subplan(m, v51, v136, v131, v70, v72, v50, v49, v44, int32(0), v48&int32(1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	if v59 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v644 = v141
	goto L1
L50:
	;
	goto L51
L51:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v145 != int32(23) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v644 = v141
	goto L1
L53:
	;
	goto L54
L54:
	;
	v148 = F_copyObjectImpl(m, v52)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v150 = F_simplify_EXISTS_query(m, v51, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+60))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+8)) = int32(0)
	v157 = F_contain_vars_of_level(m, v148, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	if v157 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v644 = v141
	goto L1
L59:
	;
	goto L60
L60:
	;
	v159 = F_contain_volatile_functions(m, v153)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	if v159 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v644 = v141
	goto L1
L63:
	;
	goto L64
L64:
	;
	v162 = F_eval_const_expressions(m, v51, v153)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v165 = F_canonicalize_qual(m, v162, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	v167 = F_make_ands_implicit(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	if v167 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v644 = v141
	goto L1
L69:
	;
	goto L70
L70:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v171 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v644 = v141
	goto L1
L72:
	;
	goto L73
L73:
	;
	v174 = int32(0)
	v179 = int32(0)
	v182 = v174
	v183 = v174
	v186 = v3
	v188 = v3
	v189 = v3
	goto L74
L74:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v179<<(uint(int32(2))%32))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v200 != int32(17) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if v254 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L76:
	;
	v262 = v179 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v262 < v263 {
		v179 = v262
		v182 = v253
		v183 = v254
		v186 = v257
		v188 = v258
		v189 = v259
		goto L74
	} else {
		goto L104
	}
L77:
	;
	v251 = F_lappend(m, v182, v199)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L15
	} else {
		goto L103
	}
L78:
	;
	v203 = F_hash_ok_operator(m, v199)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	if v203 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199)+28))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v212 = F_contain_vars_of_level(m, v210, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	if v212 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v214 = F_lappend(m, v183, v210)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L15
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v225 = F_contain_vars_of_level(m, v209, int32(1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L15
	} else {
		goto L89
	}
L85:
	;
	v216 = F_lappend(m, v186, v209)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v219 = F_lappend_oid(m, v188, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v222 = F_lappend_oid(m, v189, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L15
	} else {
		goto L88
	}
L88:
	;
	v253 = v182
	v254 = v214
	v257 = v216
	v258 = v219
	v259 = v222
	goto L76
L89:
	;
	if v225 == int32(0) {
		goto L77
	} else {
		goto L90
	}
L90:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v230 = F_get_commutator(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L15
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v230
	if v230 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v644 = v141
	goto L1
L93:
	;
	goto L94
L94:
	;
	v235 = F_hash_ok_operator(m, v199)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L15
	} else {
		goto L95
	}
L95:
	;
	if v235 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v644 = v141
	goto L1
L97:
	;
	goto L98
L98:
	;
	v239 = F_lappend(m, v183, v209)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L15
	} else {
		goto L99
	}
L99:
	;
	v241 = F_lappend(m, v186, v210)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v244 = F_lappend_oid(m, v188, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v199)+24))
	v247 = F_lappend_oid(m, v189, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L102
	}
L102:
	;
	v253 = v182
	v254 = v239
	v257 = v241
	v258 = v244
	v259 = v247
	goto L76
L103:
	;
	v253 = v251
	v254 = v183
	v257 = v186
	v258 = v188
	v259 = v189
	goto L76
L104:
	;
	goto L75
L105:
	;
	v644 = v141
	goto L1
L106:
	;
	goto L107
L107:
	;
	v268 = F_contain_vars_of_level(m, v253, int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	if v268 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v644 = v141
	goto L1
L110:
	;
	goto L111
L111:
	;
	v271 = F_contain_vars_of_level(m, v257, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L15
	} else {
		goto L112
	}
L112:
	;
	if v271 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v644 = v141
	goto L1
L114:
	;
	goto L115
L115:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+36)))
	if v274 != int32(1) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v286 = F_contain_vars_of_level(m, v254, int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L15
	} else {
		goto L124
	}
L117:
	;
	v278 = F_contain_aggs_of_level(m, v253, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	if v278 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v644 = v141
	goto L1
L120:
	;
	goto L121
L121:
	;
	v281 = F_contain_aggs_of_level(m, v257, int32(1))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L15
	} else {
		goto L122
	}
L122:
	;
	if v281 == int32(0) {
		goto L116
	} else {
		goto L123
	}
L123:
	;
	v644 = v141
	goto L1
L124:
	;
	if v286 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v644 = v141
	goto L1
L126:
	;
	goto L127
L127:
	;
	v288 = F_contain_subplans(m, v254)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	if v288 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v644 = v141
	goto L1
L130:
	;
	goto L131
L131:
	;
	F_IncrementVarSublevelsUp(m, v254, int32(-1), int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	if v253 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v294 = F_make_ands_explicit(m, v253)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L15
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v300 = int32(0)
	v302 = v300
	v311 = int32(1)
	v316 = v300
	v317 = v3
	v318 = v3
	goto L137
L136:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v148)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+8)) = v294
	goto L135
L137:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v302 < v322 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+76)) = v316
	v399 = F_make_ands_explicit(m, v317)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L15
	} else {
		goto L165
	}
L139:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v328 = v324 + v302<<(uint(int32(2))%32)
	goto L141
L140:
	;
	v328 = int32(0)
	goto L141
L141:
	;
	v329 = int32(0)
	if v257 == v329 {
		v340 = v329
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v258 == int32(0) {
		v349 = v329
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v334 <= v302 {
		v340 = int32(0)
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v340 = v336 + v302<<(uint(int32(2))%32)
	goto L142
L145:
	;
	v350 = int32(0)
	if v259 == v350 {
		v359 = v350
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v343 <= v302 {
		v349 = v329
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v349 = v345 + v302<<(uint(int32(2))%32)
	goto L145
L148:
	;
	if v328 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v353 <= v302 {
		v359 = v350
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v359 = v355 + v302<<(uint(int32(2))%32)
	goto L148
L151:
	;
	goto L138
L152:
	;
	if v340 == int32(0) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	if v349 == int32(0) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	if v359 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v372 = F_exprType(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L15
	} else {
		goto L156
	}
L156:
	;
	v374 = F_exprTypmod(m, v371)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L15
	} else {
		goto L157
	}
L157:
	;
	v376 = F_exprCollation(m, v371)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	v378 = F_generate_new_exec_param(m, v51, v372, v374, v376)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L15
	} else {
		goto L159
	}
L159:
	;
	v381 = int32(0)
	v383 = F_makeTargetEntry(m, v371, base.I32_extend16_s(v311), v381, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L15
	} else {
		goto L160
	}
L160:
	;
	v385 = F_lappend(m, v316, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L15
	} else {
		goto L161
	}
L161:
	;
	v387 = F_make_opclause(m, v369, v370, v378, v368)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L15
	} else {
		goto L162
	}
L162:
	;
	v389 = F_lappend(m, v317, v387)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L15
	} else {
		goto L163
	}
L163:
	;
	v391 = int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v396 = F_lappend_int(m, v318, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L15
	} else {
		goto L164
	}
L164:
	;
	v302 = v302 + v391
	v311 = v311 + v391
	v316 = v385
	v317 = v389
	v318 = v396
	goto L137
L165:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v402 = int32(0)
	v405 = F_subquery_planner(m, v401, v148, v51, v402, float64(0), v402)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L15
	} else {
		goto L166
	}
L166:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v408 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+20)) = v408
	v412 = F_fetch_upper_rel(m, v405, int32(7), v408)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L15
	} else {
		goto L167
	}
L167:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v412)+48))
	v415 = *(*float64)(unsafe.Add(mBase, uint32(v414)+32))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v414)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+32))
	v428 = *(*float64)(unsafe.Add(mBase, _consts[507]))
	v430 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v434 = base.F64_mul(base.F64_mul(v428, base.F64_convert_i32_s(v430)), float64(1024))
	v435 = float64(4.294967295e+09)
	if base.F64_lt(v434, v435) != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if base.F64_gt(base.F64_mul(v415, base.F64_convert_i32_u((v417+int32(7))&int32(-8)+int32(24))), base.F64_convert_i32_u(v446)) != 0 {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	v438 = v434
	goto L171
L170:
	;
	v438 = v435
	goto L171
L171:
	;
	if base.F64_lt(v438, float64(4.294967296e+09))&base.F64_ge(v438, float64(0)) != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v444 = base.I32_trunc_f64_u(v438)
	v446 = v444
	goto L168
L173:
	;
	goto L174
L174:
	;
	v446 = int32(0)
	goto L168
L175:
	;
	v644 = v141
	goto L1
L176:
	;
	goto L177
L177:
	;
	v449 = F_create_plan(m, v405, v414)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L15
	} else {
		goto L178
	}
L178:
	;
	v454 = F_build_subplan(m, v51, v449, v414, v405, v407, int32(2), int32(0), v399, v318, int32(1))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L15
	} else {
		goto L179
	}
L179:
	;
	v457 = F_palloc0(m, int32(8))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L15
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v454
	v469 = F_list_make2_impl(m, v22+int32(12), v22+int32(8))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L15
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457)+4)) = v469
	v472 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+320)) = uint8(v472)
	v644 = v457
	goto L1
L182:
	;
	v644 = l0
	goto L1
L183:
	;
	v644 = l0
	goto L1
L184:
	;
	goto L5
L185:
	;
	v635 = F_make_andclause(m, v618)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L15
	} else {
		goto L221
	}
L186:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v538)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v540 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L187:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v484)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v486 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v494 = v487
	v496 = int32(0)
	goto L193
L189:
	;
	v487 = int32(0)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if v487 < v488 {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v618 = int32(0)
	goto L185
L192:
	;
	goto L191
L193:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513+v494<<(uint(int32(2))%32))))
	v520 = F_process_sublinks_mutator(m, v517, v22+int32(16))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L15
	} else {
		goto L197
	}
L194:
	;
	v618 = v533
	goto L185
L195:
	;
	v535 = v494 + int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	if v535 < v536 {
		v494 = v535
		v496 = v533
		goto L193
	} else {
		goto L203
	}
L196:
	;
	v531 = F_lappend(m, v496, v520)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L15
	} else {
		goto L202
	}
L197:
	;
	if v520 == int32(0) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	if v524 != int32(21) {
		goto L196
	} else {
		goto L199
	}
L199:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v527 != 0 {
		goto L196
	} else {
		goto L200
	}
L200:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v520)+8))
	v529 = F_list_concat(m, v496, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L15
	} else {
		goto L201
	}
L201:
	;
	v533 = v529
	goto L195
L202:
	;
	v533 = v531
	goto L195
L203:
	;
	goto L194
L204:
	;
	v614 = F_make_orclause(m, v597)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L15
	} else {
		goto L220
	}
L205:
	;
	v597 = int32(0)
	goto L204
L206:
	;
	goto L207
L207:
	;
	v544 = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v545 <= v544 {
		v597 = v544
		goto L204
	} else {
		goto L208
	}
L208:
	;
	v549 = int32(0)
	v551 = v544
	goto L209
L209:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v540)+12))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v568+v549<<(uint(int32(2))%32))))
	v575 = F_process_sublinks_mutator(m, v572, v22+int32(16))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L15
	} else {
		goto L213
	}
L210:
	;
	v597 = v590
	goto L204
L211:
	;
	v592 = v549 + int32(1)
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v592 < v593 {
		v549 = v592
		v551 = v590
		goto L209
	} else {
		goto L219
	}
L212:
	;
	v588 = F_lappend(m, v551, v575)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L15
	} else {
		goto L218
	}
L213:
	;
	if v575 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	if v579 != int32(21) {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v582 != int32(1) {
		goto L212
	} else {
		goto L216
	}
L216:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v575)+8))
	v586 = F_list_concat(m, v551, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L15
	} else {
		goto L217
	}
L217:
	;
	v590 = v586
	goto L211
L218:
	;
	v590 = v588
	goto L211
L219:
	;
	goto L210
L220:
	;
	v644 = v614
	goto L1
L221:
	;
	v644 = v635
	goto L1
L222:
	;
	v644 = v642
	goto L1
}
func F_provider_init(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v1 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(1056)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[500])))
	if v8 != int32(1) {
		v105 = v1
		m.G0 = v5 + int32(1056)
		return v105
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1258])))
		if v12 != 0 {
			v105 = v1
			m.G0 = v5 + int32(1056)
			return v105
		} else {
			v13 = int32(1)
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1259])))
			if v15 != 0 {
				v105 = v13
				m.G0 = v5 + int32(1056)
				return v105
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(4440736)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = int32(228192)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[1260]))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v21
				v29 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(166154), v5+int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v35 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(32)
							F_errmsg_internal(m, int32(169897), v5)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(471807), int32(91), int32(94191))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v50 = F_pg_file_exists(m, v5+int32(32))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										if v50 == int32(0) {
											v54 = int32(0)
											v57 = F_errstart(m, int32(14), v54)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												if v57 != 0 {
													F_errmsg_internal(m, int32(256558), int32(0))
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(471807), int32(95), int32(94191))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v69 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v69)
															v105 = v54
															m.G0 = v5 + int32(1056)
															return v105
														}
													}
												} else {
													v69 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v69)
													v105 = v54
													m.G0 = v5 + int32(1056)
													return v105
												}
											}
										} else {
											v72 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v72)
											v80 = F_load_external_function(m, v5+int32(32), int32(94183), v72, int32(0))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												m.T0[v80].(func(*base.Module, int32))(m, int32(4445304))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v85 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[1259])) = uint8(v85)
													v88 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v88)
													v92 = F_errstart(m, int32(14), v88)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														if v92 == int32(0) {
															v105 = v13
															m.G0 = v5 + int32(1056)
															return v105
														} else {
															F_errmsg_internal(m, int32(256616), int32(0))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(471807), int32(117), int32(94191))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	v105 = v13
																	m.G0 = v5 + int32(1056)
																	return v105
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
						} else {
							v50 = F_pg_file_exists(m, v5+int32(32))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if v50 == int32(0) {
									v54 = int32(0)
									v57 = F_errstart(m, int32(14), v54)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if v57 != 0 {
											F_errmsg_internal(m, int32(256558), int32(0))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(471807), int32(95), int32(94191))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v69 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v69)
													v105 = v54
													m.G0 = v5 + int32(1056)
													return v105
												}
											}
										} else {
											v69 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v69)
											v105 = v54
											m.G0 = v5 + int32(1056)
											return v105
										}
									}
								} else {
									v72 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v72)
									v80 = F_load_external_function(m, v5+int32(32), int32(94183), v72, int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										m.T0[v80].(func(*base.Module, int32))(m, int32(4445304))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											v85 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _consts[1259])) = uint8(v85)
											v88 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _consts[1258])) = uint8(v88)
											v92 = F_errstart(m, int32(14), v88)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												if v92 == int32(0) {
													v105 = v13
													m.G0 = v5 + int32(1056)
													return v105
												} else {
													F_errmsg_internal(m, int32(256616), int32(0))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(471807), int32(117), int32(94191))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															v105 = v13
															m.G0 = v5 + int32(1056)
															return v105
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
func F_prsd_headline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v564 int32
	_ = v564
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1305 int32
	_ = v1305
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1456 int32
	_ = v1456
	var v1465 int32
	_ = v1465
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1546 int32
	_ = v1546
	var v1551 int32
	_ = v1551
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1629 int32
	_ = v1629
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1668 int32
	_ = v1668
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1825 int32
	_ = v1825
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1882 int32
	_ = v1882
	var v1901 int32
	_ = v1901
	var v1909 int32
	_ = v1909
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1939 int32
	_ = v1939
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1961 int32
	_ = v1961
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2061 int32
	_ = v2061
	var v2082 int32
	_ = v2082
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2196 int32
	_ = v2196
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2229 int32
	_ = v2229
	var v2237 int32
	_ = v2237
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2269 int32
	_ = v2269
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2303 int32
	_ = v2303
	var v2320 int32
	_ = v2320
	var v2328 int32
	_ = v2328
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2354 int32
	_ = v2354
	var v2359 int32
	_ = v2359
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2404 int32
	_ = v2404
	var v2412 int32
	_ = v2412
	var v2419 int32
	_ = v2419
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2451 int32
	_ = v2451
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2475 int32
	_ = v2475
	var v2485 int32
	_ = v2485
	var v2503 int32
	_ = v2503
	var v2510 int32
	_ = v2510
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2552 int32
	_ = v2552
	var v2559 int32
	_ = v2559
	var v2582 int32
	_ = v2582
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2605 int32
	_ = v2605
	var v2616 int32
	_ = v2616
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2637 int32
	_ = v2637
	var v2665 int32
	_ = v2665
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2709 int32
	_ = v2709
	var v2715 int32
	_ = v2715
	var v2718 int32
	_ = v2718
	var v2738 int32
	_ = v2738
	var v2774 int32
	_ = v2774
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2811 int32
	_ = v2811
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2823 int32
	_ = v2823
	var v2827 int32
	_ = v2827
	var v2851 int32
	_ = v2851
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2874 int32
	_ = v2874
	var v2885 int32
	_ = v2885
	var v2888 int32
	_ = v2888
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2899 int32
	_ = v2899
	var v2938 int32
	_ = v2938
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2971 int32
	_ = v2971
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2997 int32
	_ = v2997
	var v3002 int32
	_ = v3002
	var v3006 int32
	_ = v3006
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3023 int32
	_ = v3023
	var v3026 int32
	_ = v3026
	var v3032 int32
	_ = v3032
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3046 int32
	_ = v3046
	var v3055 int32
	_ = v3055
	var v3060 int32
	_ = v3060
	var v3064 int32
	_ = v3064
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3075 int32
	_ = v3075
	var v3081 int32
	_ = v3081
	var v3084 int32
	_ = v3084
	var v3090 int32
	_ = v3090
	var v3094 int32
	_ = v3094
	var v3096 int32
	_ = v3096
	var v3104 int32
	_ = v3104
	var v3113 int32
	_ = v3113
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3133 int32
	_ = v3133
	var v3139 int32
	_ = v3139
	var v3142 int32
	_ = v3142
	var v3148 int32
	_ = v3148
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3162 int32
	_ = v3162
	v2 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = int64(0)
	if v31 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if int32(0) < v887 {
		goto L260
	} else {
		goto L261
	}
L2:
	;
	v874 = int32(35)
	v875 = int32(15)
	v876 = v2
	v879 = int32(3)
	v882 = v2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = int32(15)
	v43 = int32(35)
	v44 = int32(3)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v45 <= int32(0) {
		v754 = v2
		v762 = v43
		v763 = v42
		v764 = v2
		v767 = v44
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v775 = int32(1)
	if v754&v775 != 0 {
		v874 = v762
		v875 = v763
		v876 = v764
		v879 = v767
		v882 = v775
		goto L1
	} else {
		goto L234
	}
L6:
	;
	v49 = v2
	v52 = v2
	v60 = v43
	v61 = v42
	v62 = v2
	v65 = v44
	goto L7
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v49<<(uint(int32(2))%32))))
	v78 = F_defGetString(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L10
	} else {
		goto L230
	}
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v86 = v82
	v87 = int32(162639)
	goto L14
L12:
	;
	v728 = v49 + int32(1)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v728 < v729 {
		v49 = v728
		v52 = v722
		v60 = v723
		v61 = v724
		v62 = v725
		v65 = v726
		goto L7
	} else {
		goto L229
	}
L13:
	;
	if v124 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v90 == v91 {
		v113 = v90
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v124 = int32(0)
	goto L13
L16:
	;
	v115 = int32(1)
	if v113 != 0 {
		v86 = v86 + v115
		v87 = v87 + v115
		goto L14
	} else {
		goto L25
	}
L17:
	;
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v101 = v90 | int32(32)
	goto L20
L19:
	;
	v101 = v90
	goto L20
L20:
	;
	if base.Ui32((v91-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v110 = v91 | int32(32)
	goto L23
L22:
	;
	v110 = v91
	goto L23
L23:
	;
	if v101 == v110 {
		v113 = v101
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v124 = v101 - v110
	goto L13
L25:
	;
	goto L15
L26:
	;
	v127 = F_pg_strtoint32(m, v78)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v133 = v129
	v134 = int32(162648)
	goto L31
L29:
	;
	v722 = v52
	v723 = v127
	v724 = v61
	v725 = v62
	v726 = v65
	goto L12
L30:
	;
	if v171 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v137 == v138 {
		v160 = v137
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v171 = int32(0)
	goto L30
L33:
	;
	v162 = int32(1)
	if v160 != 0 {
		v133 = v133 + v162
		v134 = v134 + v162
		goto L31
	} else {
		goto L42
	}
L34:
	;
	if base.Ui32((v137-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v148 = v137 | int32(32)
	goto L37
L36:
	;
	v148 = v137
	goto L37
L37:
	;
	if base.Ui32((v138-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v157 = v138 | int32(32)
	goto L40
L39:
	;
	v157 = v138
	goto L40
L40:
	;
	if v148 == v157 {
		v160 = v148
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v171 = v148 - v157
	goto L30
L42:
	;
	goto L32
L43:
	;
	v174 = F_pg_strtoint32(m, v78)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v180 = v176
	v181 = int32(402956)
	goto L48
L46:
	;
	v722 = v52
	v723 = v60
	v724 = v174
	v725 = v62
	v726 = v65
	goto L12
L47:
	;
	if v218 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v184 == v185 {
		v207 = v184
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v218 = int32(0)
	goto L47
L50:
	;
	v209 = int32(1)
	if v207 != 0 {
		v180 = v180 + v209
		v181 = v181 + v209
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v184-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v195 = v184 | int32(32)
	goto L54
L53:
	;
	v195 = v184
	goto L54
L54:
	;
	if base.Ui32((v185-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v204 = v185 | int32(32)
	goto L57
L56:
	;
	v204 = v185
	goto L57
L57:
	;
	if v195 == v204 {
		v207 = v195
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v218 = v195 - v204
	goto L47
L59:
	;
	goto L49
L60:
	;
	v221 = F_pg_strtoint32(m, v78)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v227 = v223
	v228 = int32(115312)
	goto L65
L63:
	;
	v722 = v52
	v723 = v60
	v724 = v61
	v725 = v62
	v726 = v221
	goto L12
L64:
	;
	if v265 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v231 == v232 {
		v254 = v231
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v265 = int32(0)
	goto L64
L67:
	;
	v256 = int32(1)
	if v254 != 0 {
		v227 = v227 + v256
		v228 = v228 + v256
		goto L65
	} else {
		goto L76
	}
L68:
	;
	if base.Ui32((v231-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v242 = v231 | int32(32)
	goto L71
L70:
	;
	v242 = v231
	goto L71
L71:
	;
	if base.Ui32((v232-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v251 = v232 | int32(32)
	goto L74
L73:
	;
	v251 = v232
	goto L74
L74:
	;
	if v242 == v251 {
		v254 = v242
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v265 = v242 - v251
	goto L64
L76:
	;
	goto L66
L77:
	;
	v268 = F_pg_strtoint32(m, v78)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v274 = v270
	v275 = int32(294346)
	goto L82
L80:
	;
	v722 = v52
	v723 = v60
	v724 = v61
	v725 = v268
	v726 = v65
	goto L12
L81:
	;
	if v312 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v278 == v279 {
		v301 = v278
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v312 = int32(0)
	goto L81
L84:
	;
	v303 = int32(1)
	if v301 != 0 {
		v274 = v274 + v303
		v275 = v275 + v303
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if base.Ui32((v278-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v289 = v278 | int32(32)
	goto L88
L87:
	;
	v289 = v278
	goto L88
L88:
	;
	if base.Ui32((v279-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v298 = v279 | int32(32)
	goto L91
L90:
	;
	v298 = v279
	goto L91
L91:
	;
	if v289 == v298 {
		v301 = v289
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v312 = v289 - v298
	goto L81
L93:
	;
	goto L83
L94:
	;
	v315 = F_pstrdup(m, v78)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v322 = v318
	v323 = int32(294355)
	goto L99
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v315
	v722 = v52
	v723 = v60
	v724 = v61
	v725 = v62
	v726 = v65
	goto L12
L98:
	;
	if v360 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v326 == v327 {
		v349 = v326
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v360 = int32(0)
	goto L98
L101:
	;
	v351 = int32(1)
	if v349 != 0 {
		v322 = v322 + v351
		v323 = v323 + v351
		goto L99
	} else {
		goto L110
	}
L102:
	;
	if base.Ui32((v326-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v337 = v326 | int32(32)
	goto L105
L104:
	;
	v337 = v326
	goto L105
L105:
	;
	if base.Ui32((v327-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v346 = v327 | int32(32)
	goto L108
L107:
	;
	v346 = v327
	goto L108
L108:
	;
	if v337 == v346 {
		v349 = v337
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v360 = v337 - v346
	goto L98
L110:
	;
	goto L100
L111:
	;
	v363 = F_pstrdup(m, v78)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L10
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v370 = v366
	v371 = int32(206047)
	goto L116
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v363
	v722 = v52
	v723 = v60
	v724 = v61
	v725 = v62
	v726 = v65
	goto L12
L115:
	;
	if v408 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L116:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if v374 == v375 {
		v397 = v374
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v408 = int32(0)
	goto L115
L118:
	;
	v399 = int32(1)
	if v397 != 0 {
		v370 = v370 + v399
		v371 = v371 + v399
		goto L116
	} else {
		goto L127
	}
L119:
	;
	if base.Ui32((v374-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v385 = v374 | int32(32)
	goto L122
L121:
	;
	v385 = v374
	goto L122
L122:
	;
	if base.Ui32((v375-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v394 = v375 | int32(32)
	goto L125
L124:
	;
	v394 = v375
	goto L125
L125:
	;
	if v385 == v394 {
		v397 = v385
		goto L118
	} else {
		goto L126
	}
L126:
	;
	v408 = v385 - v394
	goto L115
L127:
	;
	goto L117
L128:
	;
	v411 = F_pstrdup(m, v78)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L10
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v418 = v414
	v419 = int32(291677)
	goto L133
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v411
	v722 = v52
	v723 = v60
	v724 = v61
	v725 = v62
	v726 = v65
	goto L12
L132:
	;
	if v456 != 0 {
		goto L9
	} else {
		goto L145
	}
L133:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419))))
	if v422 == v423 {
		v445 = v422
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v456 = int32(0)
	goto L132
L135:
	;
	v447 = int32(1)
	if v445 != 0 {
		v418 = v418 + v447
		v419 = v419 + v447
		goto L133
	} else {
		goto L144
	}
L136:
	;
	if base.Ui32((v422-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v433 = v422 | int32(32)
	goto L139
L138:
	;
	v433 = v422
	goto L139
L139:
	;
	if base.Ui32((v423-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v442 = v423 | int32(32)
	goto L142
L141:
	;
	v442 = v423
	goto L142
L142:
	;
	if v433 == v442 {
		v445 = v433
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v456 = v433 - v442
	goto L132
L144:
	;
	goto L134
L145:
	;
	v457 = int32(1)
	v461 = v78
	v462 = int32(531289)
	goto L147
L146:
	;
	if v499 == int32(0) {
		v722 = v457
		v723 = v60
		v724 = v61
		v725 = v62
		v726 = v65
		goto L12
	} else {
		goto L159
	}
L147:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v465 == v466 {
		v488 = v465
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v499 = int32(0)
	goto L146
L149:
	;
	v490 = int32(1)
	if v488 != 0 {
		v461 = v461 + v490
		v462 = v462 + v490
		goto L147
	} else {
		goto L158
	}
L150:
	;
	if base.Ui32((v465-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v476 = v465 | int32(32)
	goto L153
L152:
	;
	v476 = v465
	goto L153
L153:
	;
	if base.Ui32((v466-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v485 = v466 | int32(32)
	goto L156
L155:
	;
	v485 = v466
	goto L156
L156:
	;
	if v476 == v485 {
		v488 = v476
		goto L149
	} else {
		goto L157
	}
L157:
	;
	v499 = v476 - v485
	goto L146
L158:
	;
	goto L148
L159:
	;
	v505 = v78
	v506 = int32(260641)
	goto L161
L160:
	;
	if v543 == int32(0) {
		v722 = v457
		v723 = v60
		v724 = v61
		v725 = v62
		v726 = v65
		goto L12
	} else {
		goto L173
	}
L161:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v509 == v510 {
		v532 = v509
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v543 = int32(0)
	goto L160
L163:
	;
	v534 = int32(1)
	if v532 != 0 {
		v505 = v505 + v534
		v506 = v506 + v534
		goto L161
	} else {
		goto L172
	}
L164:
	;
	if base.Ui32((v509-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v520 = v509 | int32(32)
	goto L167
L166:
	;
	v520 = v509
	goto L167
L167:
	;
	if base.Ui32((v510-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v529 = v510 | int32(32)
	goto L170
L169:
	;
	v529 = v510
	goto L170
L170:
	;
	if v520 == v529 {
		v532 = v520
		goto L163
	} else {
		goto L171
	}
L171:
	;
	v543 = v520 - v529
	goto L160
L172:
	;
	goto L162
L173:
	;
	v549 = v78
	v550 = int32(328733)
	goto L175
L174:
	;
	if v587 == int32(0) {
		v722 = v457
		v723 = v60
		v724 = v61
		v725 = v62
		v726 = v65
		goto L12
	} else {
		goto L187
	}
L175:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549))))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	if v553 == v554 {
		v576 = v553
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v587 = int32(0)
	goto L174
L177:
	;
	v578 = int32(1)
	if v576 != 0 {
		v549 = v549 + v578
		v550 = v550 + v578
		goto L175
	} else {
		goto L186
	}
L178:
	;
	if base.Ui32((v553-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v564 = v553 | int32(32)
	goto L181
L180:
	;
	v564 = v553
	goto L181
L181:
	;
	if base.Ui32((v554-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v573 = v554 | int32(32)
	goto L184
L183:
	;
	v573 = v554
	goto L184
L184:
	;
	if v564 == v573 {
		v576 = v564
		goto L177
	} else {
		goto L185
	}
L185:
	;
	v587 = v564 - v573
	goto L174
L186:
	;
	goto L176
L187:
	;
	v593 = v78
	v594 = int32(105346)
	goto L189
L188:
	;
	if v631 == int32(0) {
		v722 = v457
		v723 = v60
		v724 = v61
		v725 = v62
		v726 = v65
		goto L12
	} else {
		goto L201
	}
L189:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594))))
	if v597 == v598 {
		v620 = v597
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v631 = int32(0)
	goto L188
L191:
	;
	v622 = int32(1)
	if v620 != 0 {
		v593 = v593 + v622
		v594 = v594 + v622
		goto L189
	} else {
		goto L200
	}
L192:
	;
	if base.Ui32((v597-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v608 = v597 | int32(32)
	goto L195
L194:
	;
	v608 = v597
	goto L195
L195:
	;
	if base.Ui32((v598-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v617 = v598 | int32(32)
	goto L198
L197:
	;
	v617 = v598
	goto L198
L198:
	;
	if v608 == v617 {
		v620 = v608
		goto L191
	} else {
		goto L199
	}
L199:
	;
	v631 = v608 - v617
	goto L188
L200:
	;
	goto L190
L201:
	;
	v637 = v78
	v638 = int32(25337)
	goto L203
L202:
	;
	if v675 == int32(0) {
		v722 = v457
		v723 = v60
		v724 = v61
		v725 = v62
		v726 = v65
		goto L12
	} else {
		goto L215
	}
L203:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	if v641 == v642 {
		v664 = v641
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v675 = int32(0)
	goto L202
L205:
	;
	v666 = int32(1)
	if v664 != 0 {
		v637 = v637 + v666
		v638 = v638 + v666
		goto L203
	} else {
		goto L214
	}
L206:
	;
	if base.Ui32((v641-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v652 = v641 | int32(32)
	goto L209
L208:
	;
	v652 = v641
	goto L209
L209:
	;
	if base.Ui32((v642-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v661 = v642 | int32(32)
	goto L212
L211:
	;
	v661 = v642
	goto L212
L212:
	;
	if v652 == v661 {
		v664 = v652
		goto L205
	} else {
		goto L213
	}
L213:
	;
	v675 = v652 - v661
	goto L202
L214:
	;
	goto L204
L215:
	;
	v681 = v78
	v682 = int32(148083)
	goto L217
L216:
	;
	v722 = base.B2i32(v719 == int32(0))
	v723 = v60
	v724 = v61
	v725 = v62
	v726 = v65
	goto L12
L217:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
	if v685 == v686 {
		v708 = v685
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v719 = int32(0)
	goto L216
L219:
	;
	v710 = int32(1)
	if v708 != 0 {
		v681 = v681 + v710
		v682 = v682 + v710
		goto L217
	} else {
		goto L228
	}
L220:
	;
	if base.Ui32((v685-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v696 = v685 | int32(32)
	goto L223
L222:
	;
	v696 = v685
	goto L223
L223:
	;
	if base.Ui32((v686-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v705 = v686 | int32(32)
	goto L226
L225:
	;
	v705 = v686
	goto L226
L226:
	;
	if v696 == v705 {
		v708 = v696
		goto L219
	} else {
		goto L227
	}
L227:
	;
	v719 = v696 - v705
	goto L216
L228:
	;
	goto L218
L229:
	;
	v754 = v722
	v762 = v723
	v763 = v724
	v764 = v725
	v767 = v726
	goto L5
L230:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v738
	F_errmsg(m, int32(685588), v28-int32(-64))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(476617), int32(2666), int32(357001))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	if v763 < v762 {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L256
	}
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L10
	} else {
		goto L252
	}
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L10
	} else {
		goto L248
	}
L238:
	;
	if v763 <= int32(0) {
		goto L237
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L244
	}
L241:
	;
	if v767 < int32(0) {
		goto L236
	} else {
		goto L242
	}
L242:
	;
	if v764 < int32(0) {
		goto L235
	} else {
		goto L243
	}
L243:
	;
	v874 = v762
	v875 = v763
	v876 = v764
	v879 = v767
	v882 = int32(0)
	goto L1
L244:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L10
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = int32(162639)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = int32(162648)
	F_errmsg(m, int32(175804), v28+int32(48))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(476617), int32(2675), int32(357001))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(162648)
	F_errmsg(m, int32(327101), v28)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(476617), int32(2679), int32(357001))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(402956)
	F_errmsg(m, int32(532132), v28+int32(16))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L10
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(476617), int32(2683), int32(357001))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L10
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = int32(115312)
	F_errmsg(m, int32(532132), v28+int32(32))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L10
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(476617), int32(2687), int32(357001))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v890
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v892
	v894 = m.G0
	v896 = v894 - int32(16)
	m.G0 = v896
	v905 = F_TS_execute_locations_recurse(m, v30+int32(8), v28+int32(80), int32(1190), v896+int32(12))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L10
	} else {
		goto L263
	}
L261:
	;
	v916 = v2
	goto L262
L262:
	;
	if v876 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L263:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v896)+12))
	m.G0 = v896 + int32(16)
	if v905 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v912 = v907
	goto L266
L265:
	;
	v912 = int32(0)
	goto L266
L266:
	;
	v916 = v912
	goto L262
L267:
	;
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2964 == int32(0) {
		goto L641
	} else {
		goto L642
	}
L268:
	;
	v919 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v919
	if v882 == v919 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	goto L270
L270:
	;
	v1702 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v1702
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v1702
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v1702
	v1709 = F_palloc(m, int32(640))
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L10
	} else {
		goto L432
	}
L271:
	;
	v1629 = v1604
	goto L418
L272:
	;
	if v1583 < v1578 {
		goto L267
	} else {
		goto L417
	}
L273:
	;
	v933 = F_hlCover(m, v32, v30, v916, v28+int32(80), v28+int32(92), v28+int32(88))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L10
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v1578 = v2
	v1583 = v1566 - int32(1)
	goto L272
L276:
	;
	if v933 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v935 = int32(0)
	v936 = base.B2i32(v935 < v874)
	v939 = int32(-1)
	v951 = v939
	v956 = v939
	v963 = v939
	v964 = v2
	goto L280
L278:
	;
	goto L279
L279:
	;
	if v875 <= int32(0) {
		goto L267
	} else {
		goto L405
	}
L280:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v28)+88))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	v969 = int32(0)
	if v874 <= v935 {
		v1032 = v969
		v1033 = v968
		v1034 = v968
		v1035 = v969
		v1037 = v936
		goto L282
	} else {
		goto L283
	}
L281:
	;
	if int32(0) <= v1476 {
		v1578 = v1473
		v1583 = v1475
		goto L272
	} else {
		goto L404
	}
L282:
	;
	if v1037 != 0 {
		goto L300
	} else {
		goto L301
	}
L283:
	;
	if v967 < v968 {
		v1032 = v969
		v1033 = v968
		v1034 = v968
		v1035 = v969
		v1037 = v936
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v973 = v969
	v975 = v968
	v976 = v969
	goto L285
L285:
	;
	v998 = int32(1)
	v1002 = v972 + v975<<(uint(int32(4))%32)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1002)))
	v1007 = int32(base.Ui32(v1003)>>(uint(int32(8))%32)) & int32(255)
	if v998<<(uint(v1007)%32)&int32(241696) != 0 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1032 = v1016
	v1033 = v975
	v1034 = v1029
	v1035 = v1027
	v1037 = v1017
	goto L282
L287:
	;
	v1015 = base.B2i32(base.Ui32(v1007) <= base.Ui32(int32(17)))
	goto L289
L288:
	;
	v1015 = int32(0)
	goto L289
L289:
	;
	if v1015 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1016 = v973
	goto L292
L291:
	;
	v1016 = v973 + v998
	goto L292
L292:
	;
	v1017 = base.B2i32(v1016 < v874)
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+12))
	if v1025 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1026 = int32(base.Ui32(v1003^int32(-1))>>(uint(int32(3))%32)) & int32(1)
	goto L295
L294:
	;
	v1026 = int32(0)
	goto L295
L295:
	;
	v1027 = v1026 + v976
	v1029 = v975 + int32(1)
	if v967 < v1029 {
		v1032 = v1016
		v1033 = v975
		v1034 = v1029
		v1035 = v1027
		v1037 = v1017
		goto L282
	} else {
		goto L296
	}
L296:
	;
	if v1016 < v874 {
		v973 = v1016
		v975 = v1029
		v976 = v1027
		goto L285
	} else {
		goto L297
	}
L297:
	;
	goto L286
L298:
	;
	v1394 = base.B2i32(v1369 <= v968) & base.B2i32(v967 <= v1368)
	v1395 = int32(1)
	if v1394&((v964^v1395)&v1395) != 0 {
		goto L384
	} else {
		goto L385
	}
L299:
	;
	v1368 = v1343
	v1369 = v968
	v1370 = v1345
	goto L298
L300:
	;
	v1058 = v1034 - int32(1)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v1059 <= v1058 {
		v1150 = v1032
		v1151 = v1033
		v1153 = v1035
		goto L303
	} else {
		goto L304
	}
L301:
	;
	goto L302
L302:
	;
	if v1032 <= v875 {
		v1343 = v1033
		v1345 = v1035
		goto L299
	} else {
		goto L359
	}
L303:
	;
	if v875 <= v1150 {
		v1343 = v1151
		v1345 = v1153
		goto L299
	} else {
		goto L330
	}
L304:
	;
	if v874 <= v1032 {
		v1150 = v1032
		v1151 = v1033
		v1153 = v1035
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1063 = v1032
	v1065 = v1058
	v1066 = v1035
	goto L306
L306:
	;
	v1090 = v1062 + v1065<<(uint(int32(4))%32)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1090)))
	v1093 = int32(base.Ui32(v1091) >> (uint(int32(8)) % 32))
	if v1065 <= v967 {
		v1118 = v1063
		v1119 = v1066
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1150 = v1118
	v1151 = v1065
	v1153 = v1119
	goto L303
L308:
	;
	v1122 = v1093 & int32(255)
	if int32(1)<<(uint(v1122)%32)&int32(15987104) != 0 {
		goto L318
	} else {
		goto L319
	}
L309:
	;
	v1095 = int32(1)
	v1098 = v1093 & int32(255)
	if v1095<<(uint(v1098)%32)&int32(241696) != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1106 = base.B2i32(base.Ui32(v1098) <= base.Ui32(int32(17)))
	goto L312
L311:
	;
	v1106 = int32(0)
	goto L312
L312:
	;
	if v1106 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1107 = v1063
	goto L315
L314:
	;
	v1107 = v1063 + v1095
	goto L315
L315:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+12))
	if v1108 == int32(0) {
		v1118 = v1107
		v1119 = v1066
		goto L308
	} else {
		goto L316
	}
L316:
	;
	v1118 = v1107
	v1119 = int32(base.Ui32(v1091^int32(-1))>>(uint(int32(3))%32))&int32(1) + v1066
	goto L308
L317:
	;
	v1147 = v1065 + int32(1)
	if v1059 <= v1147 {
		v1150 = v1118
		v1151 = v1065
		v1153 = v1119
		goto L303
	} else {
		goto L328
	}
L318:
	;
	v1130 = base.B2i32(base.Ui32(v1122) <= base.Ui32(int32(23)))
	goto L320
L319:
	;
	v1130 = int32(0)
	goto L320
L320:
	;
	v1131 = int32(0)
	if base.B2i32(v1130 == v1131)&base.B2i32(v879 < int32(base.Ui32(v1091)>>(uint(int32(16))%32))) == v1131 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+12))
	if v1139 == int32(0) {
		goto L317
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	if v875 <= v1118 {
		v1150 = v1118
		v1151 = v1065
		v1153 = v1119
		goto L303
	} else {
		goto L327
	}
L324:
	;
	if v1091&int32(8) != 0 {
		goto L317
	} else {
		goto L325
	}
L325:
	;
	if v1118 < v875 {
		goto L317
	} else {
		goto L326
	}
L326:
	;
	v1150 = v1118
	v1151 = v1065
	v1153 = v1119
	goto L303
L327:
	;
	goto L317
L328:
	;
	if v1118 < v874 {
		v1063 = v1118
		v1065 = v1147
		v1066 = v1119
		goto L306
	} else {
		goto L329
	}
L329:
	;
	goto L307
L330:
	;
	v1176 = int32(0)
	v1178 = v968 - int32(1)
	if v1178 < v1176 {
		v1368 = v1151
		v1369 = v1176
		v1370 = v1153
		goto L298
	} else {
		goto L331
	}
L331:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1182 = v1150
	v1185 = v1153
	v1186 = v1178
	goto L332
L332:
	;
	v1207 = int32(1)
	v1211 = v1181 + v1186<<(uint(int32(4))%32)
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1211)))
	v1216 = int32(base.Ui32(v1212)>>(uint(int32(8))%32)) & int32(255)
	if v1207<<(uint(v1216)%32)&int32(241696) != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v1368 = v1151
	v1369 = int32(0)
	v1370 = v1235
	goto L298
L334:
	;
	v1224 = base.B2i32(base.Ui32(v1216) <= base.Ui32(int32(17)))
	goto L336
L335:
	;
	v1224 = int32(0)
	goto L336
L336:
	;
	if v1224 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1225 = v1182
	goto L339
L338:
	;
	v1225 = v1182 + v1207
	goto L339
L339:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+12))
	if v1233 != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1234 = int32(base.Ui32(v1212^int32(-1))>>(uint(int32(3))%32)) & int32(1)
	goto L342
L341:
	;
	v1234 = int32(0)
	goto L342
L342:
	;
	v1235 = v1234 + v1185
	if v874 <= v1225 {
		v1368 = v1151
		v1369 = v1186
		v1370 = v1235
		goto L298
	} else {
		goto L343
	}
L343:
	;
	if v879 < int32(base.Ui32(v1212)>>(uint(int32(16))%32)) {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if int32(0) < v1186 {
		v1182 = v1225
		v1185 = v1235
		v1186 = v1186 - int32(1)
		goto L332
	} else {
		goto L358
	}
L345:
	;
	if int32(1)<<(uint(v1216)%32)&int32(15987104) != 0 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	v1249 = int32(1)
	goto L347
L347:
	;
	if v1249 != 0 {
		goto L351
	} else {
		goto L352
	}
L348:
	;
	v1247 = base.B2i32(base.Ui32(v1216) <= base.Ui32(int32(23)))
	goto L350
L349:
	;
	v1247 = int32(0)
	goto L350
L350:
	;
	v1249 = v1247
	goto L347
L351:
	;
	if v1233 == int32(0) {
		goto L344
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	if v875 <= v1225 {
		v1368 = v1151
		v1369 = v1186
		v1370 = v1235
		goto L298
	} else {
		goto L357
	}
L354:
	;
	if v1212&int32(8) != 0 {
		goto L344
	} else {
		goto L355
	}
L355:
	;
	if v1225 < v875 {
		goto L344
	} else {
		goto L356
	}
L356:
	;
	v1368 = v1151
	v1369 = v1186
	v1370 = v1235
	goto L298
L357:
	;
	goto L344
L358:
	;
	goto L333
L359:
	;
	if v1034 < v967 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1263 = v1034
	goto L362
L361:
	;
	v1263 = v967
	goto L362
L362:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1265 = v1032
	v1266 = v1033
	v1268 = v1035
	v1270 = v1263
	goto L363
L363:
	;
	v1292 = v1264 + v1270<<(uint(int32(4))%32)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)))
	v1297 = int32(base.Ui32(v1293)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v1297)%32)&int32(15987104) != 0 {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	v1343 = v1329
	v1345 = v1327
	goto L299
L365:
	;
	v1305 = base.B2i32(base.Ui32(v1297) <= base.Ui32(int32(23)))
	goto L367
L366:
	;
	v1305 = int32(0)
	goto L367
L367:
	;
	if base.B2i32(v1305 == int32(0))&base.B2i32(v879 < int32(base.Ui32(v1293)>>(uint(int32(16))%32))) != 0 {
		v1343 = v1266
		v1345 = v1268
		goto L299
	} else {
		goto L368
	}
L368:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+12))
	if v1315 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1316 = v1293 & int32(8)
	goto L371
L370:
	;
	v1316 = int32(1)
	goto L371
L371:
	;
	if v1316 == int32(0) {
		v1343 = v1266
		v1345 = v1268
		goto L299
	} else {
		goto L372
	}
L372:
	;
	v1321 = int32(1)
	if v1315 != 0 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1326 = int32(base.Ui32(v1293)>>(uint(int32(3))%32))&v1321 - v1321
	goto L375
L374:
	;
	v1326 = int32(0)
	goto L375
L375:
	;
	v1327 = v1326 + v1268
	v1328 = int32(1)
	v1329 = v1270 - v1328
	if v1328<<(uint(v1297)%32)&int32(241696) != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1339 = base.B2i32(base.Ui32(v1297) <= base.Ui32(int32(17)))
	goto L378
L377:
	;
	v1339 = int32(0)
	goto L378
L378:
	;
	if v1339 != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1340 = v1265
	goto L381
L380:
	;
	v1340 = v1265 - v1328
	goto L381
L381:
	;
	if v875 < v1340 {
		v1265 = v1340
		v1266 = v1329
		v1268 = v1327
		v1270 = v1329
		goto L363
	} else {
		goto L382
	}
L382:
	;
	goto L364
L383:
	;
	v1484 = F_hlCover(m, v32, v30, v916, v28+int32(80), v28+int32(92), v28+int32(88))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L10
	} else {
		goto L402
	}
L384:
	;
	v1473 = v1369
	v1475 = v1368
	v1476 = v1370
	v1477 = v1394
	goto L383
L385:
	;
	v1400 = v1394 ^ v964
	if base.B2i32(v1400&int32(1) == int32(0))&base.B2i32(v963 < v1370) != 0 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	if (v1400|base.B2i32(v1370 != v963))&int32(1) != 0 {
		v1473 = v951
		v1475 = v956
		v1476 = v963
		v1477 = v964
		goto L383
	} else {
		goto L387
	}
L387:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1414 = v1411 + v1368<<(uint(int32(4))%32)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	v1419 = int32(base.Ui32(v1415)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v1419)%32)&int32(15987104) != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1427 = base.B2i32(base.Ui32(v1419) <= base.Ui32(int32(23)))
	goto L390
L389:
	;
	v1427 = int32(0)
	goto L390
L390:
	;
	v1428 = int32(0)
	if base.B2i32(v1427 == v1428)&base.B2i32(v879 < int32(base.Ui32(v1415)>>(uint(int32(16))%32))) == v1428 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	if v1415&int32(8) != 0 {
		v1473 = v951
		v1475 = v956
		v1476 = v963
		v1477 = v964
		goto L383
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	v1443 = v1411 + v956<<(uint(int32(4))%32)
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1443)))
	v1448 = int32(base.Ui32(v1444)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v1448)%32)&int32(15987104) != 0 {
		goto L396
	} else {
		goto L397
	}
L394:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1414)+12))
	if v1438 == int32(0) {
		v1473 = v951
		v1475 = v956
		v1476 = v963
		v1477 = v964
		goto L383
	} else {
		goto L395
	}
L395:
	;
	goto L393
L396:
	;
	v1456 = base.B2i32(base.Ui32(v1448) <= base.Ui32(int32(23)))
	goto L398
L397:
	;
	v1456 = int32(0)
	goto L398
L398:
	;
	if base.B2i32(v1456 == int32(0))&base.B2i32(v879 < int32(base.Ui32(v1444)>>(uint(int32(16))%32))) != 0 {
		v1473 = v951
		v1475 = v956
		v1476 = v963
		v1477 = v964
		goto L383
	} else {
		goto L399
	}
L399:
	;
	if v1444&int32(8) != 0 {
		goto L384
	} else {
		goto L400
	}
L400:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+12))
	if v1465 != 0 {
		v1473 = v951
		v1475 = v956
		v1476 = v963
		v1477 = v964
		goto L383
	} else {
		goto L401
	}
L401:
	;
	goto L384
L402:
	;
	if v1484 != 0 {
		v951 = v1473
		v956 = v1475
		v963 = v1476
		v964 = v1477
		goto L280
	} else {
		goto L403
	}
L403:
	;
	goto L281
L404:
	;
	goto L279
L405:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v1515 <= int32(0) {
		goto L267
	} else {
		goto L406
	}
L406:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1519 = int32(0)
	v1521 = v1519
	v1522 = v1519
	goto L407
L407:
	;
	v1546 = int32(1)
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1521<<(uint(int32(4))%32))+1)))
	if v1546<<(uint(v1551)%32)&int32(241696) != 0 {
		goto L409
	} else {
		goto L410
	}
L408:
	;
	v1604 = v1561
	v1609 = v1521
	goto L271
L409:
	;
	v1559 = base.B2i32(base.Ui32(v1551) <= base.Ui32(int32(17)))
	goto L411
L410:
	;
	v1559 = int32(0)
	goto L411
L411:
	;
	if v1559 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1560 = v1522
	goto L414
L413:
	;
	v1560 = v1522 + v1546
	goto L414
L414:
	;
	v1561 = int32(0)
	v1563 = v1521 + int32(1)
	if v1515 <= v1563 {
		v1604 = v1561
		v1609 = v1521
		goto L271
	} else {
		goto L415
	}
L415:
	;
	if v1560 < v875 {
		v1521 = v1563
		v1522 = v1560
		goto L407
	} else {
		goto L416
	}
L416:
	;
	goto L408
L417:
	;
	v1604 = v1578
	v1609 = v1583
	goto L271
L418:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1647 = v1629 << (uint(int32(4)) % 32)
	v1648 = v1645 + v1647
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+12))
	if v1649 != 0 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	goto L267
L420:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1648)))
	*(*int32)(unsafe.Add(mBase, uint32(v1648))) = v1650 | int32(1)
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1655 = v1654
	goto L422
L421:
	;
	v1655 = v1645
	goto L422
L422:
	;
	v1656 = v1655 + v1647
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)))
	v1659 = int32(base.Ui32(v1657) >> (uint(int32(8)) % 32))
	if v882 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L423:
	;
	v1689 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1686+v1647))) = int32(base.Ui32(v1685)>>(uint(v1689)%32))&v1689 | v1685&int32(-3) ^ v1689
	v1700 = v1629 + int32(1)
	if v1700 <= v1609 {
		v1629 = v1700
		goto L418
	} else {
		goto L431
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1656))) = v1657 | v1679
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1682+v1647)))
	v1685 = v1684
	v1686 = v1682
	goto L423
L425:
	;
	v1679 = int32(16)
	goto L424
L426:
	;
	switch v1659&int32(255) - int32(5) {
	case 0, 10, 11, 12:
		goto L425
	default:
		v1685 = v1657
		v1686 = v1655
		goto L423
	case 8:
		v1679 = int32(4)
		goto L424
	}
L427:
	;
	goto L428
L428:
	;
	v1668 = v1659 & int32(255)
	if base.Ui32(int32(17)) < base.Ui32(v1668) {
		v1685 = v1657
		v1686 = v1655
		goto L423
	} else {
		goto L429
	}
L429:
	;
	if int32(1)<<(uint(v1668)%32)&int32(229408) == int32(0) {
		v1685 = v1657
		v1686 = v1655
		goto L423
	} else {
		goto L430
	}
L430:
	;
	goto L425
L431:
	;
	goto L419
L432:
	;
	v1717 = F_hlCover(m, v32, v30, v916, v28+int32(80), v28+int32(92), v28+int32(88))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L10
	} else {
		goto L433
	}
L433:
	;
	if v1717 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1731 = int32(32)
	v1735 = v1709
	v1736 = v2
	goto L437
L435:
	;
	v2049 = v1709
	v2050 = v2
	goto L436
L436:
	;
	if v876 <= int32(0) {
		goto L489
	} else {
		goto L490
	}
L437:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v28)+88))
	if v1745 <= v1746 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	v2049 = v2016
	v2050 = v2017
	goto L436
L439:
	;
	v1749 = v1745
	v1751 = v1746
	v1759 = v1731
	v1763 = v1735
	v1764 = v1736
	goto L442
L440:
	;
	v2012 = v1731
	v2016 = v1735
	v2017 = v1736
	goto L441
L441:
	;
	v2032 = F_hlCover(m, v32, v30, v916, v28+int32(80), v28+int32(92), v28+int32(88))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L10
	} else {
		goto L486
	}
L442:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v1775 = v1749
	goto L444
L443:
	;
	v2012 = v1984
	v2016 = v1985
	v2017 = v1996
	goto L441
L444:
	;
	v1801 = v1773 + v1775<<(uint(int32(4))%32)
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1801)+12))
	if v1802 != 0 {
		goto L447
	} else {
		goto L448
	}
L445:
	;
	v1813 = int32(0)
	if v1751 < v1775 {
		v1874 = v1775
		v1875 = v1813
		v1882 = v1813
		goto L453
	} else {
		goto L454
	}
L446:
	;
	goto L445
L447:
	;
	v1804 = v1775 + int32(1)
	if v1751 < v1804 {
		goto L446
	} else {
		goto L450
	}
L448:
	;
	goto L449
L449:
	;
	v1810 = v1775 + int32(1)
	if v1810 <= v1751 {
		v1775 = v1810
		goto L444
	} else {
		goto L452
	}
L450:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1801)))
	if v1806&int32(8) != 0 {
		v1775 = v1804
		goto L444
	} else {
		goto L451
	}
L451:
	;
	goto L446
L452:
	;
	goto L446
L453:
	;
	if v1751 <= v1874 {
		v1955 = v1751
		v1961 = v1882
		goto L465
	} else {
		goto L466
	}
L454:
	;
	v1817 = v1775
	v1818 = v1813
	v1825 = v1813
	goto L455
L455:
	;
	if v874 <= v1825 {
		v1874 = v1817
		v1875 = v1818
		v1882 = v1825
		goto L453
	} else {
		goto L457
	}
L456:
	;
	v1874 = v1871
	v1875 = v1869
	v1882 = v1860
	goto L453
L457:
	;
	v1842 = int32(1)
	v1846 = v1773 + v1817<<(uint(int32(4))%32)
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1846)))
	v1851 = int32(base.Ui32(v1847)>>(uint(int32(8))%32)) & int32(255)
	if v1842<<(uint(v1851)%32)&int32(241696) != 0 {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v1859 = base.B2i32(base.Ui32(v1851) <= base.Ui32(int32(17)))
	goto L460
L459:
	;
	v1859 = int32(0)
	goto L460
L460:
	;
	if v1859 != 0 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1860 = v1825
	goto L463
L462:
	;
	v1860 = v1825 + v1842
	goto L463
L463:
	;
	v1863 = int32(0)
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1846)+12))
	v1869 = v1818 + base.B2i32(v1847&int32(8) == v1863)&base.B2i32(v1865 != v1863)
	v1871 = v1817 + int32(1)
	if v1871 <= v1751 {
		v1817 = v1871
		v1818 = v1869
		v1825 = v1860
		goto L455
	} else {
		goto L464
	}
L464:
	;
	goto L456
L465:
	;
	if v1759 <= v1764 {
		goto L481
	} else {
		goto L482
	}
L466:
	;
	if v1874 < v1775 {
		v1955 = v1874
		v1961 = v1882
		goto L465
	} else {
		goto L467
	}
L467:
	;
	v1901 = v1874
	v1909 = v1882
	goto L468
L468:
	;
	v1927 = v1773 + v1901<<(uint(int32(4))%32)
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1927)))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1927)+12))
	if v1928&int32(8) != 0 {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	v1955 = v1901
	v1961 = v1948
	goto L465
L470:
	;
	v1933 = int32(0)
	goto L472
L471:
	;
	v1933 = v1930
	goto L472
L472:
	;
	if v1933 != 0 {
		v1955 = v1901
		v1961 = v1909
		goto L465
	} else {
		goto L473
	}
L473:
	;
	v1934 = int32(1)
	v1939 = int32(base.Ui32(v1928)>>(uint(int32(8))%32)) & int32(255)
	if v1934<<(uint(v1939)%32)&int32(241696) != 0 {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v1947 = base.B2i32(base.Ui32(v1939) <= base.Ui32(int32(17)))
	goto L476
L475:
	;
	v1947 = int32(0)
	goto L476
L476:
	;
	if v1947 != 0 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1948 = v1909
	goto L479
L478:
	;
	v1948 = v1909 - v1934
	goto L479
L479:
	;
	v1950 = v1901 - int32(1)
	if v1775 <= v1950 {
		v1901 = v1950
		v1909 = v1948
		goto L468
	} else {
		goto L480
	}
L480:
	;
	goto L469
L481:
	;
	v1980 = F_repalloc(m, v1763, v1759*int32(40))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L10
	} else {
		goto L484
	}
L482:
	;
	v1984 = v1759
	v1985 = v1763
	goto L483
L483:
	;
	v1988 = v1985 + v1764*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v1988)+12)) = v1961
	*(*int32)(unsafe.Add(mBase, uint32(v1988)+4)) = v1955
	*(*int32)(unsafe.Add(mBase, uint32(v1988))) = v1775
	v1992 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1988)+16)) = uint16(v1992)
	*(*int32)(unsafe.Add(mBase, uint32(v1988)+8)) = v1875
	v1995 = int32(1)
	v1996 = v1764 + v1995
	v1998 = v1955 + v1995
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v28)+88))
	if v1998 <= v1999 {
		v1749 = v1998
		v1751 = v1999
		v1759 = v1984
		v1763 = v1985
		v1764 = v1996
		goto L442
	} else {
		goto L485
	}
L484:
	;
	v1984 = v1759 << (uint(int32(1)) % 32)
	v1985 = v1980
	goto L483
L485:
	;
	goto L443
L486:
	;
	if v2032 != 0 {
		v1731 = v2012
		v1735 = v2016
		v1736 = v2017
		goto L437
	} else {
		goto L487
	}
L487:
	;
	goto L438
L488:
	;
	F_pfree(m, v2049)
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L10
	} else {
		goto L640
	}
L489:
	;
	if v875 <= int32(0) {
		goto L488
	} else {
		goto L615
	}
L490:
	;
	v2061 = int32(0)
	v2082 = v2061
	goto L492
L491:
	;
	if int32(0) < v2738 {
		goto L488
	} else {
		goto L614
	}
L492:
	;
	v2090 = int32(0)
	if v2050 <= v2061 {
		goto L489
	} else {
		goto L494
	}
L493:
	;
	v2738 = v876
	goto L491
L494:
	;
	v2093 = v2090
	v2096 = v2090
	v2097 = int32(-1)
	v2098 = int32(2147483647)
	goto L495
L495:
	;
	v2120 = v2049 + v2093*int32(20)
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2120)+16)))
	if v2121 != 0 {
		v2134 = v2096
		v2135 = v2097
		v2136 = v2098
		goto L497
	} else {
		goto L498
	}
L496:
	;
	if v2135 < int32(0) {
		v2738 = v2082
		goto L491
	} else {
		goto L511
	}
L497:
	;
	v2138 = v2093 + int32(1)
	if v2138 != v2050 {
		v2093 = v2138
		v2096 = v2134
		v2097 = v2135
		v2098 = v2136
		goto L495
	} else {
		goto L510
	}
L498:
	;
	v2122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2120)+17)))
	if v2122 != 0 {
		v2134 = v2096
		v2135 = v2097
		v2136 = v2098
		goto L497
	} else {
		goto L499
	}
L499:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+8))
	if v2096 < v2123 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+12))
	v2134 = v2123
	v2135 = v2093
	v2136 = v2125
	goto L497
L501:
	;
	goto L502
L502:
	;
	if v2123 != v2096 {
		v2134 = v2096
		v2135 = v2097
		v2136 = v2098
		goto L497
	} else {
		goto L503
	}
L503:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2120)+12))
	if v2127 < v2098 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v2129 = v2093
	goto L506
L505:
	;
	v2129 = v2097
	goto L506
L506:
	;
	if v2098 < v2127 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2131 = v2098
	goto L509
L508:
	;
	v2131 = v2127
	goto L509
L509:
	;
	v2134 = v2096
	v2135 = v2129
	v2136 = v2131
	goto L497
L510:
	;
	goto L496
L511:
	;
	v2144 = v2049 + v2135*int32(20)
	v2145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2144)+16)) = uint8(v2145)
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2144)+4))
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2144)))
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v2144)+12))
	if v874 <= v2149 {
		v2532 = v2147
		v2535 = v2149
		v2552 = v2148
		goto L512
	} else {
		goto L513
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2144)+12)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(v2144)+4)) = v2532
	*(*int32)(unsafe.Add(mBase, uint32(v2144))) = v2552
	if v2552 <= v2532 {
		goto L584
	} else {
		goto L585
	}
L513:
	;
	v2151 = v874 - v2149
	v2153 = base.I32_div_s(v2151, int32(2))
	v2155 = v2148 - int32(1)
	if v2155 < int32(0) {
		v2303 = v2149
		goto L515
	} else {
		goto L516
	}
L514:
	;
	v2346 = v2147 + int32(1)
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v2347 <= v2346 {
		v2485 = v2328
		goto L550
	} else {
		goto L551
	}
L515:
	;
	v2320 = v2148
	v2328 = v2303
	goto L514
L516:
	;
	if v2151 < int32(2) {
		v2303 = v2149
		goto L515
	} else {
		goto L517
	}
L517:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2161+v2155<<(uint(int32(4))%32))))
	if v2165&int32(2) != 0 {
		v2303 = v2149
		goto L515
	} else {
		goto L518
	}
L518:
	;
	v2169 = v2155
	v2170 = v2165
	v2176 = v2149
	v2179 = int32(0)
	goto L519
L519:
	;
	v2196 = int32(base.Ui32(v2170)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v2196)%32)&int32(241696) != 0 {
		goto L521
	} else {
		goto L522
	}
L520:
	;
	if v2148 <= v2169 {
		v2320 = v2169
		v2328 = v2211
		goto L514
	} else {
		goto L531
	}
L521:
	;
	v2204 = base.B2i32(base.Ui32(v2196) <= base.Ui32(int32(17)))
	goto L523
L522:
	;
	v2204 = int32(0)
	goto L523
L523:
	;
	if v2204 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2207 = int32(1)
	v2211 = v2176 + v2207
	v2212 = v2179 + v2207
	goto L526
L525:
	;
	v2211 = v2176
	v2212 = v2179
	goto L526
L526:
	;
	if v2169 <= int32(0) {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	goto L520
L528:
	;
	if v2153 <= v2212 {
		goto L527
	} else {
		goto L529
	}
L529:
	;
	v2217 = v2169 - int32(1)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2161+v2217<<(uint(int32(4))%32))))
	if v2221&int32(2) == int32(0) {
		v2169 = v2217
		v2170 = v2221
		v2176 = v2211
		v2179 = v2212
		goto L519
	} else {
		goto L530
	}
L530:
	;
	goto L527
L531:
	;
	v2229 = v2169
	v2237 = v2211
	goto L532
L532:
	;
	v2256 = v2161 + v2229<<(uint(int32(4))%32)
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2256)))
	v2261 = int32(base.Ui32(v2257)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v2261)%32)&int32(15987104) != 0 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v2303 = v2291
	goto L515
L534:
	;
	v2269 = base.B2i32(base.Ui32(v2261) <= base.Ui32(int32(23)))
	goto L536
L535:
	;
	v2269 = int32(0)
	goto L536
L536:
	;
	if base.B2i32(v2269 == int32(0))&base.B2i32(v879 < int32(base.Ui32(v2257)>>(uint(int32(16))%32))) != 0 {
		v2320 = v2229
		v2328 = v2237
		goto L514
	} else {
		goto L537
	}
L537:
	;
	if v2257&int32(8) == int32(0) {
		goto L538
	} else {
		goto L539
	}
L538:
	;
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+12))
	if v2280 != 0 {
		v2320 = v2229
		v2328 = v2237
		goto L514
	} else {
		goto L541
	}
L539:
	;
	goto L540
L540:
	;
	v2281 = int32(1)
	if v2281<<(uint(v2261)%32)&int32(241696) != 0 {
		goto L542
	} else {
		goto L543
	}
L541:
	;
	goto L540
L542:
	;
	v2290 = base.B2i32(base.Ui32(v2261) <= base.Ui32(int32(17)))
	goto L544
L543:
	;
	v2290 = int32(0)
	goto L544
L544:
	;
	if v2290 != 0 {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2291 = v2237
	goto L547
L546:
	;
	v2291 = v2237 - v2281
	goto L547
L547:
	;
	v2293 = v2229 + int32(1)
	if v2293 != v2148 {
		v2229 = v2293
		v2237 = v2291
		goto L532
	} else {
		goto L548
	}
L548:
	;
	goto L533
L549:
	;
	v2532 = v2503
	v2535 = v2510
	v2552 = v2320
	goto L512
L550:
	;
	v2532 = v2147
	v2535 = v2485
	v2552 = v2320
	goto L512
L551:
	;
	if v874 <= v2328 {
		v2485 = v2328
		goto L550
	} else {
		goto L552
	}
L552:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v2350+v2346<<(uint(int32(4))%32))))
	if v2354&int32(2) != 0 {
		v2485 = v2328
		goto L550
	} else {
		goto L553
	}
L553:
	;
	v2359 = v2346
	v2365 = v2328
	v2367 = v2354
	goto L554
L554:
	;
	v2382 = int32(1)
	v2387 = int32(base.Ui32(v2367)>>(uint(int32(8))%32)) & int32(255)
	if v2382<<(uint(v2387)%32)&int32(241696) != 0 {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	if v2359 <= v2147 {
		v2503 = v2359
		v2510 = v2396
		goto L549
	} else {
		goto L566
	}
L556:
	;
	v2395 = base.B2i32(base.Ui32(v2387) <= base.Ui32(int32(17)))
	goto L558
L557:
	;
	v2395 = int32(0)
	goto L558
L558:
	;
	if v2395 != 0 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v2396 = v2365
	goto L561
L560:
	;
	v2396 = v2365 + v2382
	goto L561
L561:
	;
	v2398 = v2359 + int32(1)
	if v2347 <= v2398 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	goto L555
L563:
	;
	if v874 <= v2396 {
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2350+v2398<<(uint(int32(4))%32))))
	if v2404&int32(2) == int32(0) {
		v2359 = v2398
		v2365 = v2396
		v2367 = v2404
		goto L554
	} else {
		goto L565
	}
L565:
	;
	goto L562
L566:
	;
	v2412 = v2359
	v2419 = v2396
	goto L567
L567:
	;
	v2438 = v2350 + v2412<<(uint(int32(4))%32)
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2438)))
	v2443 = int32(base.Ui32(v2439)>>(uint(int32(8))%32)) & int32(255)
	if int32(1)<<(uint(v2443)%32)&int32(15987104) != 0 {
		goto L569
	} else {
		goto L570
	}
L568:
	;
	v2485 = v2473
	goto L550
L569:
	;
	v2451 = base.B2i32(base.Ui32(v2443) <= base.Ui32(int32(23)))
	goto L571
L570:
	;
	v2451 = int32(0)
	goto L571
L571:
	;
	if base.B2i32(v2451 == int32(0))&base.B2i32(v879 < int32(base.Ui32(v2439)>>(uint(int32(16))%32))) != 0 {
		v2503 = v2412
		v2510 = v2419
		goto L549
	} else {
		goto L572
	}
L572:
	;
	if v2439&int32(8) == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2438)+12))
	if v2462 != 0 {
		v2503 = v2412
		v2510 = v2419
		goto L549
	} else {
		goto L576
	}
L574:
	;
	goto L575
L575:
	;
	v2463 = int32(1)
	if v2463<<(uint(v2443)%32)&int32(241696) != 0 {
		goto L577
	} else {
		goto L578
	}
L576:
	;
	goto L575
L577:
	;
	v2472 = base.B2i32(base.Ui32(v2443) <= base.Ui32(int32(17)))
	goto L579
L578:
	;
	v2472 = int32(0)
	goto L579
L579:
	;
	if v2472 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v2473 = v2419
	goto L582
L581:
	;
	v2473 = v2419 - v2463
	goto L582
L582:
	;
	v2475 = v2412 - int32(1)
	if v2147 < v2475 {
		v2412 = v2475
		v2419 = v2473
		goto L567
	} else {
		goto L583
	}
L583:
	;
	goto L568
L584:
	;
	v2559 = v2552
	goto L587
L585:
	;
	goto L586
L586:
	;
	v2665 = int32(0)
	goto L601
L587:
	;
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2584 = v2559 << (uint(int32(4)) % 32)
	v2585 = v2582 + v2584
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2585)+12))
	if v2586 != 0 {
		goto L589
	} else {
		goto L590
	}
L588:
	;
	goto L586
L589:
	;
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(v2585)))
	*(*int32)(unsafe.Add(mBase, uint32(v2585))) = v2587 | int32(1)
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2592 = v2591
	goto L591
L590:
	;
	v2592 = v2582
	goto L591
L591:
	;
	v2593 = v2592 + v2584
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(v2593)))
	v2596 = int32(base.Ui32(v2594) >> (uint(int32(8)) % 32))
	if v882 == int32(0) {
		goto L595
	} else {
		goto L596
	}
L592:
	;
	v2626 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2623+v2584))) = int32(base.Ui32(v2622)>>(uint(v2626)%32))&v2626 | v2622&int32(-3) ^ v2626
	v2637 = v2559 + int32(1)
	if v2637 <= v2532 {
		v2559 = v2637
		goto L587
	} else {
		goto L600
	}
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2593))) = v2594 | v2616
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2619+v2584)))
	v2622 = v2621
	v2623 = v2619
	goto L592
L594:
	;
	v2616 = int32(16)
	goto L593
L595:
	;
	switch v2596&int32(255) - int32(5) {
	case 0, 10, 11, 12:
		goto L594
	default:
		v2622 = v2594
		v2623 = v2592
		goto L592
	case 8:
		v2616 = int32(4)
		goto L593
	}
L596:
	;
	goto L597
L597:
	;
	v2605 = v2596 & int32(255)
	if base.Ui32(int32(17)) < base.Ui32(v2605) {
		v2622 = v2594
		v2623 = v2592
		goto L592
	} else {
		goto L598
	}
L598:
	;
	if int32(1)<<(uint(v2605)%32)&int32(229408) == int32(0) {
		v2622 = v2594
		v2623 = v2592
		goto L592
	} else {
		goto L599
	}
L599:
	;
	goto L594
L600:
	;
	goto L588
L601:
	;
	if v2665 == v2135 {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	v2718 = v2082 + int32(1)
	if v2718 != v876 {
		v2082 = v2718
		goto L492
	} else {
		goto L613
	}
L603:
	;
	v2715 = v2665 + int32(1)
	if v2715 != v2050 {
		v2665 = v2715
		goto L601
	} else {
		goto L612
	}
L604:
	;
	v2693 = v2049 + v2665*int32(20)
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2693)))
	v2695 = base.B2i32(v2694 < v2552)
	v2696 = int32(0)
	if base.B2i32(v2695 == v2696)&base.B2i32(v2694 <= v2532) == v2696 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+4))
	if v2532 < v2702 {
		goto L608
	} else {
		goto L609
	}
L606:
	;
	goto L607
L607:
	;
	v2709 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2693)+17)) = uint8(v2709)
	goto L603
L608:
	;
	v2705 = v2695
	goto L610
L609:
	;
	v2705 = base.B2i32(v2552 <= v2702)
	goto L610
L610:
	;
	if v2705 != int32(1) {
		goto L603
	} else {
		goto L611
	}
L611:
	;
	goto L607
L612:
	;
	goto L602
L613:
	;
	goto L493
L614:
	;
	goto L489
L615:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v2774 <= int32(0) {
		goto L488
	} else {
		goto L616
	}
L616:
	;
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2778 = int32(0)
	v2780 = v2778
	v2783 = v2778
	goto L617
L617:
	;
	v2805 = int32(0)
	v2806 = int32(1)
	v2811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2777+v2780<<(uint(int32(4))%32))+1)))
	if v2806<<(uint(v2811)%32)&int32(241696) != 0 {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	v2827 = v2805
	goto L626
L619:
	;
	v2819 = base.B2i32(base.Ui32(v2811) <= base.Ui32(int32(17)))
	goto L621
L620:
	;
	v2819 = v2805
	goto L621
L621:
	;
	if v2819 != 0 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v2820 = v2783
	goto L624
L623:
	;
	v2820 = v2783 + v2806
	goto L624
L624:
	;
	v2823 = v2780 + int32(1)
	if base.B2i32(v2820 < v875)&base.B2i32(v2823 < v2774) != 0 {
		v2780 = v2823
		v2783 = v2820
		goto L617
	} else {
		goto L625
	}
L625:
	;
	goto L618
L626:
	;
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2853 = v2827 << (uint(int32(4)) % 32)
	v2854 = v2851 + v2853
	v2855 = *(*int32)(unsafe.Add(mBase, uint32(v2854)+12))
	if v2855 != 0 {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L488
L628:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2854)))
	*(*int32)(unsafe.Add(mBase, uint32(v2854))) = v2856 | int32(1)
	v2860 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2861 = v2860
	goto L630
L629:
	;
	v2861 = v2851
	goto L630
L630:
	;
	v2862 = v2861 + v2853
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2862)))
	v2865 = int32(base.Ui32(v2863) >> (uint(int32(8)) % 32))
	if v882 == int32(0) {
		goto L634
	} else {
		goto L635
	}
L631:
	;
	v2899 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v2894+v2827<<(uint(int32(4))%32)))) = int32(base.Ui32(v2893)>>(uint(v2899)%32))&v2899 | v2893&int32(-3) ^ v2899
	if v2827 != v2780 {
		v2827 = v2827 + int32(1)
		goto L626
	} else {
		goto L639
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2862))) = v2863 | v2885
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v2888+v2827<<(uint(int32(4))%32))))
	v2893 = v2892
	v2894 = v2888
	goto L631
L633:
	;
	v2885 = int32(16)
	goto L632
L634:
	;
	switch v2865&int32(255) - int32(5) {
	case 0, 10, 11, 12:
		goto L633
	default:
		v2893 = v2863
		v2894 = v2861
		goto L631
	case 8:
		v2885 = int32(4)
		goto L632
	}
L635:
	;
	goto L636
L636:
	;
	v2874 = v2865 & int32(255)
	if base.Ui32(int32(17)) < base.Ui32(v2874) {
		v2893 = v2863
		v2894 = v2861
		goto L631
	} else {
		goto L637
	}
L637:
	;
	if int32(1)<<(uint(v2874)%32)&int32(229408) == int32(0) {
		v2893 = v2863
		v2894 = v2861
		goto L631
	} else {
		goto L638
	}
L638:
	;
	goto L633
L639:
	;
	goto L627
L640:
	;
	goto L267
L641:
	;
	v2968 = F_pstrdup(m, int32(523713))
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L10
	} else {
		goto L644
	}
L642:
	;
	goto L643
L643:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v2971 == int32(0) {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v2968
	goto L643
L645:
	;
	v2975 = F_pstrdup(m, int32(523717))
	mBase = m.M
	v2976 = m.ExcPending
	if v2976 != 0 {
		goto L10
	} else {
		goto L648
	}
L646:
	;
	v2978 = v2971
	goto L647
L647:
	;
	v2979 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v2979 == int32(0) {
		goto L649
	} else {
		goto L650
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v2975
	v2978 = v2975
	goto L647
L649:
	;
	v2983 = F_pstrdup(m, int32(705587))
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L10
	} else {
		goto L652
	}
L650:
	;
	v2987 = v2978
	v2988 = v2979
	goto L651
L651:
	;
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2989&int32(3) == int32(0) {
		v3013 = v2989
		goto L655
	} else {
		goto L656
	}
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v2983
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v2987 = v2986
	v2988 = v2983
	goto L651
L653:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+28)) = uint16(v3046)
	if v2987&int32(3) == int32(0) {
		v3071 = v2987
		goto L672
	} else {
		goto L673
	}
L654:
	;
	v3046 = v3038 - v2989
	goto L653
L655:
	;
	v3017 = v3013
	goto L664
L656:
	;
	v2997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2989))))
	if v2997 == int32(0) {
		goto L657
	} else {
		goto L658
	}
L657:
	;
	v3046 = int32(0)
	goto L653
L658:
	;
	goto L659
L659:
	;
	v3002 = v2989
	goto L660
L660:
	;
	v3006 = v3002 + int32(1)
	if v3006&int32(3) == int32(0) {
		v3013 = v3006
		goto L655
	} else {
		goto L662
	}
L661:
	;
	v3038 = v3006
	goto L654
L662:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3006))))
	if v3011 != 0 {
		v3002 = v3006
		goto L660
	} else {
		goto L663
	}
L663:
	;
	goto L661
L664:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(v3017)))
	v3026 = int32(-2139062144)
	if (int32(16843008)-v3023|v3023)&v3026 == v3026 {
		v3017 = v3017 + int32(4)
		goto L664
	} else {
		goto L666
	}
L665:
	;
	v3032 = v3017
	goto L667
L666:
	;
	goto L665
L667:
	;
	v3036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3032))))
	if v3036 != 0 {
		v3032 = v3032 + int32(1)
		goto L667
	} else {
		goto L669
	}
L668:
	;
	v3038 = v3032
	goto L654
L669:
	;
	goto L668
L670:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+30)) = uint16(v3104)
	if v2988&int32(3) == int32(0) {
		v3129 = v2988
		goto L689
	} else {
		goto L690
	}
L671:
	;
	v3104 = v3096 - v2987
	goto L670
L672:
	;
	v3075 = v3071
	goto L681
L673:
	;
	v3055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2987))))
	if v3055 == int32(0) {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v3104 = int32(0)
	goto L670
L675:
	;
	goto L676
L676:
	;
	v3060 = v2987
	goto L677
L677:
	;
	v3064 = v3060 + int32(1)
	if v3064&int32(3) == int32(0) {
		v3071 = v3064
		goto L672
	} else {
		goto L679
	}
L678:
	;
	v3096 = v3064
	goto L671
L679:
	;
	v3069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3064))))
	if v3069 != 0 {
		v3060 = v3064
		goto L677
	} else {
		goto L680
	}
L680:
	;
	goto L678
L681:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v3075)))
	v3084 = int32(-2139062144)
	if (int32(16843008)-v3081|v3081)&v3084 == v3084 {
		v3075 = v3075 + int32(4)
		goto L681
	} else {
		goto L683
	}
L682:
	;
	v3090 = v3075
	goto L684
L683:
	;
	goto L682
L684:
	;
	v3094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3090))))
	if v3094 != 0 {
		v3090 = v3090 + int32(1)
		goto L684
	} else {
		goto L686
	}
L685:
	;
	v3096 = v3090
	goto L671
L686:
	;
	goto L685
L687:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+32)) = uint16(v3162)
	m.G0 = v28 + int32(96)
	return v32
L688:
	;
	v3162 = v3154 - v2988
	goto L687
L689:
	;
	v3133 = v3129
	goto L698
L690:
	;
	v3113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2988))))
	if v3113 == int32(0) {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v3162 = int32(0)
	goto L687
L692:
	;
	goto L693
L693:
	;
	v3118 = v2988
	goto L694
L694:
	;
	v3122 = v3118 + int32(1)
	if v3122&int32(3) == int32(0) {
		v3129 = v3122
		goto L689
	} else {
		goto L696
	}
L695:
	;
	v3154 = v3122
	goto L688
L696:
	;
	v3127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3122))))
	if v3127 != 0 {
		v3118 = v3122
		goto L694
	} else {
		goto L697
	}
L697:
	;
	goto L695
L698:
	;
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(v3133)))
	v3142 = int32(-2139062144)
	if (int32(16843008)-v3139|v3139)&v3142 == v3142 {
		v3133 = v3133 + int32(4)
		goto L698
	} else {
		goto L700
	}
L699:
	;
	v3148 = v3133
	goto L701
L700:
	;
	goto L699
L701:
	;
	v3152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3148))))
	if v3152 != 0 {
		v3148 = v3148 + int32(1)
		goto L701
	} else {
		goto L703
	}
L702:
	;
	v3154 = v3148
	goto L688
L703:
	;
	goto L702
}
func F_prsd_nexttoken(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_TParserGet(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
			v17 = v16
		} else {
			v17 = int32(0)
		}
		return v17
	}
}
func F_prsd_start(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc0(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[358]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_consts[1003])))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v5
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v6
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v19
		if int32(2) <= v19 {
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v25)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1002])))
			v33 = F_palloc(m, v5<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v28 == int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v33
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v40 = F_pg_mb2wchar_with_len(m, v38, v33, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v56 = F_palloc(m, int32(32))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v58 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+24)) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(0)
							return v8
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v33
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					F_char2wchar(m, v33, v43+int32(1), v46, v43, int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v56 = F_palloc(m, int32(32))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v58 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+24)) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v58
							*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(0)
							return v8
						}
					}
				}
			}
		} else {
			v50 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v50)
			v56 = F_palloc(m, int32(32))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v56))) = v58
				*(*int64)(unsafe.Add(mBase, uint32(v56)+24)) = v58
				*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v58
				*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(0)
				return v8
			}
		}
	}
}
func F_pts_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errcontext_msg(m, int32(674763), v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_pull_ors(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v2 = int32(0)
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v40
L2:
	;
	v9 = v2
	v11 = v2
	goto L7
L3:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v5 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v40 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v11<<(uint(int32(2))%32))))
	if v16 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v40 = v34
	goto L1
L9:
	;
	v36 = v11 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v36 < v37 {
		v9 = v34
		v11 = v36
		goto L7
	} else {
		goto L18
	}
L10:
	;
	v32 = F_lappend(m, v9, v16)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L14
	} else {
		goto L17
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(21) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != int32(1) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v26 = F_pull_ors(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v30 = F_list_concat(m, v9, v26)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v34 = v30
	goto L9
L17:
	;
	v34 = v32
	goto L9
L18:
	;
	goto L8
}
func F_pull_paramids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v4 == int32(8) {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v9 = F_bms_add_member(m, v7, v8)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
				return int32(0)
			}
		} else {
			v17 = F_expression_tree_walker_impl(m, l0, int32(875), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = v17
				return v19
			}
		}
	} else {
		v19 = int32(0)
		return v19
	}
}
func F_pull_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
	if l0 == int32(0) {
		v32 = v10
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
		m.G0 = v8 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v15 == int32(6) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v18 != l1 {
				v32 = v10
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
				m.G0 = v8 + int32(16)
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != 0 {
					v32 = v10
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
					m.G0 = v8 + int32(16)
					return
				} else {
					v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
					v24 = F_bms_add_member(m, v10, v21+int32(7))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v32 = v24
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			v29 = F_expression_tree_walker_impl(m, l0, int32(896), v8+int32(8))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v32 = v31
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v32
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pull_varnos_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v3
	v17 = F_query_or_expression_tree_walker_impl(m, l1, int32(895), v6+int32(4), v3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
		m.G0 = v6 + int32(16)
		return v21
	}
}
func F_pushStop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_palloc0(m, int32(12))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = int32(3)
		*(*uint8)(unsafe.Add(mBase, uint32(v4))) = uint8(v6)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = F_lcons(m, v4, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
			return
		}
	}
}
func F_pushf_create_mbuf_writer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	v5 = F_palloc0(m, int32(24))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(4335472)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(0)
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
		return v14
	}
}
func F_pushf_write(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v8 - v27
	if int32(0) < v28 {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	if int32(0) < v22 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = m.T0[v14].(func(*base.Module, int32, int32, int32, int32) int32)(m, v11, v15, l1, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v20 = F_pushf_write(m, v11, l1, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	v22 = v16
	goto L4
L10:
	;
	v22 = v20
	goto L4
L11:
	;
	v25 = int32(-12)
	goto L13
L12:
	;
	v25 = v22
	goto L13
L13:
	;
	return v25
L14:
	;
	return v116
L15:
	;
	v116 = int32(0)
	goto L14
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v31 + v27
	if l2 < v28 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v44 = l1
	v45 = l2
	v46 = v8
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v51 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	if v28 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v36 + l2
	goto L15
L23:
	;
	v34 = F__emscripten_memcpy_bulkmem(m, v32, l1, l2)
	mBase = m.M
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = l1 + v28
	v45 = l2 - v28
	v46 = v42
	goto L18
L27:
	;
	v39 = F__emscripten_memcpy_bulkmem(m, v32, l1, v28)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	if int32(0) < v57 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = m.T0[v51].(func(*base.Module, int32, int32, int32, int32) int32)(m, v47, v52, v48, v46)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v55 = F_pushf_write(m, v47, v48, v46)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	v57 = v53
	goto L30
L35:
	;
	v57 = v55
	goto L30
L36:
	;
	v60 = int32(-12)
	goto L38
L37:
	;
	v60 = v57
	goto L38
L38:
	;
	if v60 < int32(0) {
		v116 = v60
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v63
	if v45 <= v63 {
		v116 = v63
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = v44
	v71 = v45
	v72 = v68
	goto L41
L41:
	;
	if v72 < v71 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v71 != 0 {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v80 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	if int32(0) < v86 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = m.T0[v80].(func(*base.Module, int32, int32, int32, int32) int32)(m, v77, v81, v70, v72)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v84 = F_pushf_write(m, v77, v70, v72)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L51
	}
L50:
	;
	v86 = v82
	goto L46
L51:
	;
	v86 = v84
	goto L46
L52:
	;
	v89 = int32(-12)
	goto L54
L53:
	;
	v89 = v86
	goto L54
L54:
	;
	if v89 < int32(0) {
		v116 = v89
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v94 = int32(0)
	v95 = v71 - v92
	if v94 < v95 {
		v70 = v70 + v92
		v71 = v95
		v72 = v92
		goto L41
	} else {
		goto L56
	}
L56:
	;
	v116 = v94
	goto L14
L57:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v101 + v71
	goto L15
L58:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v98, v70, v71)
	mBase = m.M
	goto L60
L59:
	;
	goto L60
L60:
	;
	goto L57
}
func F_pushval_asis(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	var v8 int32
	_ = v8
	F_pushValue(m, l1, l2, l3, l4, l5)
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
