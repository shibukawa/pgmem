package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterConstrUpdateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v6 = F_heap_copytuple(m, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
		v10 = v8 + v9
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v11 == int32(1) {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)) = uint8(v14)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+75)) = uint8(v14)
		} else {
		}
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
		if v18 == int32(1) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+73)) = uint8(v21)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+74)) = uint8(v23)
		} else {
		}
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
		if v25 == int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+106)) = uint8(v28)
		} else {
		}
		F_CatalogTupleUpdate(m, l1, v6+int32(4), v6)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _consts[444]))
			if v35 != 0 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v38 = int32(0)
				F_RunObjectPostAlterHook(m, int32(2606), v37, v38, v38, v38)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					F_CacheInvalidateRelcacheByRelid(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_pfree(m, v6)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				F_CacheInvalidateRelcacheByRelid(m, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_pfree(m, v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_AlterSystemSetConfigFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v36 int32
	_ = v36
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
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v350 int32
	_ = v350
	var v364 int32
	_ = v364
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v415 int32
	_ = v415
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v523 int32
	_ = v523
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v728 int32
	_ = v728
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v872 int32
	_ = v872
	var v885 int32
	_ = v885
	var v899 int32
	_ = v899
	var v914 int32
	_ = v914
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1005 int32
	_ = v1005
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1036 int32
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1064 int32
	_ = v1064
	var v1079 int32
	_ = v1079
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1109 int32
	_ = v1109
	var v1122 int32
	_ = v1122
	var v1138 int32
	_ = v1138
	var v1153 int32
	_ = v1153
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1240 int32
	_ = v1240
	var v1249 int32
	_ = v1249
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1304 int32
	_ = v1304
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1482 int32
	_ = v1482
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1511 int32
	_ = v1511
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1543 int32
	_ = v1543
	var v1555 int32
	_ = v1555
	var v1571 int32
	_ = v1571
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
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
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1638 int32
	_ = v1638
	var v1651 int32
	_ = v1651
	var v1664 int32
	_ = v1664
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1705 int32
	_ = v1705
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1734 int32
	_ = v1734
	var v1746 int32
	_ = v1746
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1879 int32
	_ = v1879
	var v1891 int32
	_ = v1891
	var v1911 int32
	_ = v1911
	var v1921 int32
	_ = v1921
	var v1962 int32
	_ = v1962
	var v1975 int32
	_ = v1975
	var v1989 int32
	_ = v1989
	var v2004 int32
	_ = v2004
	var v2016 int32
	_ = v2016
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2069 int32
	_ = v2069
	var v2081 int32
	_ = v2081
	var v2097 int32
	_ = v2097
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2150 int32
	_ = v2150
	var v2155 int32
	_ = v2155
	var v2159 int32
	_ = v2159
	var v2164 int32
	_ = v2164
	var v2178 int32
	_ = v2178
	var v2190 int32
	_ = v2190
	var v2204 int32
	_ = v2204
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2244 int32
	_ = v2244
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2276 int32
	_ = v2276
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2314 int32
	_ = v2314
	var v2326 int32
	_ = v2326
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2359 int32
	_ = v2359
	var v2371 int32
	_ = v2371
	var v2387 int32
	_ = v2387
	var v2402 int32
	_ = v2402
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2428 int64
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(256)
	m.G0 = v27
	v30 = l0
	v31 = v27
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = v2
	v38 = v2
	v39 = v2
	v40 = v2
	v41 = v2
	v42 = v2
	v43 = int32(-1)
	v51 = v27
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
	if v43 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v2427 = int32(m.ExcTag)
	v2428 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2427 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L7:
	;
	v56 = int32(16)
	v57 = v51 - v56
	m.G0 = v57
	v60 = v57 - v56
	m.G0 = v60
	v63 = v60 - v56
	m.G0 = v63
	v66 = v63 - v56
	m.G0 = v66
	v68 = int32(1024)
	v69 = v66 - v68
	m.G0 = v69
	v72 = v69 - v68
	m.G0 = v72
	v75 = v72 - v56
	m.G0 = v75
	v78 = v75 - v56
	m.G0 = v78
	v81 = v78 - int32(96)
	m.G0 = v81
	v84 = v81 - int32(160)
	m.G0 = v84
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v86
	v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1149])))
	if v91 == v86 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v1599 = v32
	v1600 = v33
	v1601 = v34
	v1602 = v35
	v1603 = v36
	v1604 = v37
	v1605 = v38
	v1606 = v39
	v1607 = v40
	v1608 = v41
	v1609 = v42
	v1618 = v51
	goto L9
L9:
	;
	if v1609 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	switch v153 {
	case 0:
		goto L22
	case 1, 4:
		v213 = v41
		v214 = int32(0)
		goto L21
	default:
		goto L23
	case 5:
		goto L20
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(1088))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errmsg(m, int32(101354), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4628), int32(409627))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, _consts[1150]))
	if v1496 != 0 {
		goto L201
	} else {
		goto L202
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+208)) = int32(356454)
	v1433 = F_pg_snprintf(m, v69, int32(1024), int32(217224), v31+int32(208))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L198
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v394 = F_find_option(m, v151, int32(0), int32(1), int32(10))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L44
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v322 = F_superuser(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L36
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v225 = F_superuser(m)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L28
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v211 = F_ExtractSetVariableArgs(m, v150)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v169
	F_errmsg_internal(m, int32(508426), v31+int32(48))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4648), int32(409627))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	v213 = v211
	v214 = v211
	goto L21
L28:
	;
	if v225 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v238 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v250 = F_pg_parameter_aclcheck(m, v151, v238, int64(8192))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v250 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(16797828))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v151
	F_errmsg(m, int32(736270), v31+int32(176))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4671), int32(409627))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L3
L36:
	;
	if v322 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(16797828))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errmsg(m, int32(560563), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4660), int32(409627))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = int32(356454)
	v955 = F_pg_snprintf(m, v69, int32(1024), int32(217224), v31+int32(128))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L127
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v849 = int32(10)
	v850 = F___strchrnul(m, v214, v849)
	mBase = m.M
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	if v852 == v849 {
		goto L119
	} else {
		goto L120
	}
L44:
	;
	if v394 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v396 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	if v214 != 0 {
		goto L73
	} else {
		goto L74
	}
L48:
	;
	if v214 == int32(0) {
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+21)))
	if v397&int32(33) == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(33685829))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = v151
	F_errmsg(m, int32(482191), v31+int32(144))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4696), int32(409627))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L3
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v476 = F_parse_and_validate_value(m, v394, v214, int32(3), int32(21), v75, v78)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v476 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v394)+24))
	if v539 != int32(3) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(50856066))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = v151
	F_errmsg(m, int32(763900), v31+int32(160))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4712), int32(409627))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L3
L66:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v558 == int32(0) {
		goto L43
	} else {
		goto L70
	}
L67:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v542 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v542)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v558)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L43
L72:
	;
	if v214 == int32(0) {
		goto L42
	} else {
		goto L117
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v787 = F_assignable_custom_variable_name(m, v151, int32(0), int32(21))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L116
	}
L74:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v573 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v590 = v573
	v594 = v151
	v598 = int32(1)
	v600 = int32(0)
	goto L76
L76:
	;
	v603 = v590 & int32(255)
	v605 = base.B2i32(v603 != int32(46))
	if v605 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v605&v744 != 0 {
		goto L72
	} else {
		goto L115
	}
L78:
	;
	v748 = v594 + int32(1)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748))))
	if v749 != 0 {
		v590 = v749
		v594 = v748
		v598 = base.B2i32(v603 == int32(46))
		v600 = v744
		goto L76
	} else {
		goto L114
	}
L79:
	;
	v608 = int32(1)
	if v598&v608 == int32(0) {
		v744 = v608
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v623 = int32(534348)
	v624 = base.I32_extend8_s(v590)
	v625 = int32(54)
	goto L86
L82:
	;
	goto L73
L83:
	;
	if v624 < int32(0) {
		v744 = v600
		goto L78
	} else {
		goto L109
	}
L84:
	;
	v728 = int32(0)
	goto L83
L85:
	;
	v706 = v699
	v708 = v701
	goto L103
L86:
	;
	goto L94
L94:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1151])))
	if v662 == v624&int32(255) {
		v692 = v623
		v694 = v625
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v694 == int32(0) {
		goto L84
	} else {
		goto L102
	}
L96:
	;
	goto L97
L97:
	;
	v672 = v623
	v674 = v625
	goto L98
L98:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	v679 = v678 ^ v624&int32(255)*int32(16843009)
	v682 = int32(-2139062144)
	if (int32(16843008)-v679|v679)&v682 != v682 {
		v699 = v672
		v701 = v674
		goto L85
	} else {
		goto L100
	}
L99:
	;
	v692 = v687
	v694 = v689
	goto L95
L100:
	;
	v686 = int32(4)
	v687 = v672 + v686
	v689 = v674 - v686
	if base.Ui32(int32(3)) < base.Ui32(v689) {
		v672 = v687
		v674 = v689
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v699 = v692
	v701 = v694
	goto L85
L103:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v624&int32(255) == v711 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	goto L84
L105:
	;
	v728 = v706
	goto L83
L106:
	;
	goto L107
L107:
	;
	v713 = int32(1)
	v716 = v708 - v713
	if v716 != 0 {
		v706 = v706 + v713
		v708 = v716
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	if v728 != 0 {
		v744 = v600
		goto L78
	} else {
		goto L110
	}
L110:
	;
	if v598&int32(1) != 0 {
		goto L73
	} else {
		goto L111
	}
L111:
	;
	if base.Ui32(int32(63)) < base.Ui32(v603) {
		goto L73
	} else {
		goto L112
	}
L112:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v624))%64)&int64(287948969894477825) == int64(0) {
		goto L73
	} else {
		goto L113
	}
L113:
	;
	v744 = v600
	goto L78
L114:
	;
	goto L77
L115:
	;
	goto L73
L116:
	;
	goto L72
L117:
	;
	goto L43
L118:
	;
	if v856 == int32(0) {
		goto L42
	} else {
		goto L122
	}
L119:
	;
	v856 = v850
	goto L121
L120:
	;
	v856 = int32(0)
	goto L121
L121:
	;
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(50856066))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errmsg(m, int32(392433), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4744), int32(409627))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L126
	}
L126:
	;
	goto L3
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+116)) = int32(247949)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v69
	v974 = F_pg_snprintf(m, v72, int32(1024), int32(188250), v31+int32(112))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v987 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v991 = F_LWLockAcquire(m, v987+int32(4480), int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1005 = F___fstatat(m, int32(-100), v69, v81, int32(0))
	mBase = m.M
	goto L130
L130:
	;
	if v1005 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1019 = F_AllocateFile(m, v69, int32(242988))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v1168 != 0 {
		goto L151
	} else {
		goto L152
	}
L134:
	;
	if v1019 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1092 = F_ParseConfigFp(m, v1019, v69, int32(0), int32(15), v60, v63)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L142
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode_for_file_access(m)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v69
	F_errmsg(m, int32(313457), v31-int32(-64))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4782), int32(409627))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L141
	}
L141:
	;
	goto L3
L142:
	;
	if v1092 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1164 = F_FreeFile(m, v1019)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L150
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode(m, int32(22))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = v69
	F_errmsg(m, int32(752926), v31+int32(96))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4790), int32(409627))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L149
	}
L149:
	;
	goto L3
L150:
	;
	goto L133
L151:
	;
	v1186 = int32(0)
	v1187 = v1168
	goto L154
L152:
	;
	goto L153
L153:
	;
	if v214 == int32(0) {
		v1482 = v213
		goto L17
	} else {
		goto L190
	}
L154:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+24))
	v1211 = v1193
	v1215 = v151
	goto L158
L155:
	;
	goto L153
L156:
	;
	if v1194 != 0 {
		v1186 = v1320
		v1187 = v1194
		goto L154
	} else {
		goto L189
	}
L157:
	;
	if v1260 != 0 {
		goto L175
	} else {
		goto L176
	}
L158:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1215))))
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	if v1220 != 0 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v1260 = base.I32_extend8_s(v1240) - base.I32_extend8_s(v1249)
	goto L157
L160:
	;
	v1228 = int32(1)
	if base.Ui32((v1220-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L161:
	;
	if v1219&int32(255) != 0 {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v1219&int32(255) != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1320 = v1187
	goto L156
L165:
	;
	v1227 = int32(-1)
	goto L167
L166:
	;
	v1227 = int32(0)
	goto L167
L167:
	;
	v1260 = v1227
	goto L157
L168:
	;
	v1240 = v1220 | int32(32)
	goto L170
L169:
	;
	v1240 = v1220
	goto L170
L170:
	;
	if base.Ui32((v1219-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1249 = v1219 | int32(32)
	goto L173
L172:
	;
	v1249 = v1219
	goto L173
L173:
	;
	if v1240 == v1249&int32(255) {
		v1211 = v1211 + v1228
		v1215 = v1215 + v1228
		goto L158
	} else {
		goto L174
	}
L174:
	;
	goto L159
L175:
	;
	v1320 = v1187
	goto L156
L176:
	;
	goto L177
L177:
	;
	if v1186 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v1194 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+24)) = v1194
	goto L178
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v1194
	goto L178
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v1186
	goto L184
L183:
	;
	goto L184
L184:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1187)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v1266)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L185
	}
L185:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v1279)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L186
	}
L186:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v1292)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_pfree(m, v1187)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L188
	}
L188:
	;
	v1320 = v1186
	goto L156
L189:
	;
	goto L155
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1359 = F_palloc(m, int32(28))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1371 = F_pstrdup(m, v151)
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1359))) = v1371
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1384 = F_pstrdup(m, v214)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+4)) = v1384
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1400 = F_pstrdup(m, int32(793540))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L194
	}
L194:
	;
	v1402 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+24)) = v1402
	*(*uint16)(unsafe.Add(mBase, uint32(v1359)+20)) = uint16(v1402)
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+16)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+12)) = v1400
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v1409 == v1402 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v1359
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v1359
	v1482 = v213
	goto L17
L196:
	;
	goto L197
L197:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v1414)+24)) = v1359
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v1359
	v1482 = v213
	goto L17
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+196)) = int32(247949)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v69
	v1452 = F_pg_snprintf(m, v72, int32(1024), int32(188250), v31+int32(192))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1465 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v1469 = F_LWLockAcquire(m, v1465+int32(4480), int32(0))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L200
	}
L200:
	;
	v1482 = v41
	goto L17
L201:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1497)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_RunObjectPostAlterHookStr(m, v151, int32(8192), v1498)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	v1524 = F_BasicOpenFile(m, v72, int32(578))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1524
	if v1524 < int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	v1591 = *(*int32)(unsafe.Add(mBase, _consts[53]))
	goto L213
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errcode_for_file_access(m)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v72
	F_errmsg(m, int32(313457), v31+int32(80))
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1482
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v57
	F_errfinish(m, int32(525932), int32(4832), int32(409627))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		v2424 = v84
		goto L6
	} else {
		goto L212
	}
L212:
	;
	goto L3
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v31 + int32(212)
	goto L216
L214:
	;
	v1599 = v57
	v1600 = v69
	v1601 = v72
	v1602 = v60
	v1603 = v66
	v1604 = v63
	v1605 = v84
	v1606 = v1591
	v1607 = v1589
	v1608 = v1482
	v1609 = int32(0)
	v1618 = v84
	goto L9
L216:
	;
	goto L214
L217:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2340 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1605
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1603)))
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1602)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_initStringInfo(m, v1599)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v1607
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1606
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v1603)))
	if int32(0) <= v2300 {
		goto L292
	} else {
		goto L293
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, int32(792997))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, int32(791576))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+4))
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v1690 = F_write(m, v1625, v1679, v1678)
	mBase = m.M
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+4))
	if v1690 != v1691 {
		goto L217
	} else {
		goto L224
	}
L224:
	;
	if v1626 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1705 = v1626
	goto L228
L226:
	;
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2150 = int32(*(*uint8)(unsafe.Add(mBase, _consts[44])))
	if v2150 != int32(1) {
		v2164 = int32(0)
		goto L275
	} else {
		goto L276
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	v1728 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1727))) = uint8(v1728)
	*(*int32)(unsafe.Add(mBase, uint32(v1599)+12)) = v1728
	*(*int32)(unsafe.Add(mBase, uint32(v1599)+4)) = v1728
	goto L230
L229:
	;
	goto L227
L230:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1705)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, v1734)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, int32(723082))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L232
	}
L232:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1705)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v1771 = int32(0)
	v1774 = F_strlen(m, v1760)
	mBase = m.M
	v1775 = int32(1)
	v1779 = F_emscripten_builtin_malloc(m, v1774<<(uint(v1775)%32)|v1775)
	mBase = m.M
	if v1779 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v1774 <= int32(0) {
		v1911 = v1771
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L235
L235:
	;
	if v1779 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L236:
	;
	v1921 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1911+v1779))) = uint8(v1921)
	goto L235
L237:
	;
	v1782 = int32(1)
	if v1774 != v1782 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v1803 = v1771
	v1806 = v1771
	v1807 = v1771
	goto L241
L239:
	;
	v1867 = v1771
	v1870 = v1771
	goto L240
L240:
	;
	if v1774&v1782 == int32(0) {
		v1911 = v1867
		goto L236
	} else {
		goto L250
	}
L241:
	;
	v1812 = v1760 + v1806
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1812))))
	if base.B2i32(v1813 != int32(92))&base.B2i32(v1813 != int32(39)) == int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1867 = v1848
	v1870 = v1846
	goto L240
L243:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1803+v1779))) = uint8(v1813)
	v1825 = v1803 + int32(1)
	goto L245
L244:
	;
	v1825 = v1803
	goto L245
L245:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1825+v1779))) = uint8(v1813)
	v1829 = v1825 + int32(1)
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1812)+1)))
	if base.B2i32(v1830 != int32(92))&base.B2i32(v1830 != int32(39)) == int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1829+v1779))) = uint8(v1830)
	v1842 = v1825 + int32(2)
	goto L248
L247:
	;
	v1842 = v1829
	goto L248
L248:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1842+v1779))) = uint8(v1830)
	v1845 = int32(2)
	v1846 = v1806 + v1845
	v1848 = v1842 + int32(1)
	v1850 = v1807 + v1845
	if v1850 != v1774&int32(2147483646) {
		v1803 = v1848
		v1806 = v1846
		v1807 = v1850
		goto L241
	} else {
		goto L249
	}
L249:
	;
	goto L242
L250:
	;
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1760+v1870))))
	if base.B2i32(v1879 != int32(92))&base.B2i32(v1879 != int32(39)) == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1867+v1779))) = uint8(v1879)
	v1891 = v1867 + int32(1)
	goto L253
L252:
	;
	v1891 = v1867
	goto L253
L253:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1891+v1779))) = uint8(v1879)
	v1911 = v1891 + int32(1)
	goto L236
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, v1779)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L261
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errcode(m, int32(8389))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errmsg(m, int32(14086), int32(0))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errfinish(m, int32(525932), int32(4507), int32(406303))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L260
	}
L260:
	;
	goto L3
L261:
	;
	F_emscripten_builtin_free(m, v1779)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_appendStringInfoString(m, v1599, int32(792806))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(0)
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+4))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2046 = F_write(m, v1625, v2035, v2034)
	mBase = m.M
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1599)+4))
	if v2046 != v2047 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2050 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v1705)+24))
	if v2113 != 0 {
		v1705 = v2113
		goto L228
	} else {
		goto L273
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L268
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errcode_for_file_access(m)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v1601
	F_errmsg(m, int32(313140), v31+int32(16))
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errfinish(m, int32(525932), int32(4521), int32(406303))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L272
	}
L272:
	;
	goto L3
L273:
	;
	goto L229
L274:
	;
	if v2164 != 0 {
		goto L281
	} else {
		goto L282
	}
L275:
	;
	goto L274
L276:
	;
	goto L277
L277:
	;
	v2155 = F_fsync(m, v1625)
	mBase = m.M
	if v2155 != int32(-1) {
		v2164 = v2155
		goto L275
	} else {
		goto L279
	}
L278:
	;
	v2164 = int32(-1)
	goto L275
L279:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	if v2159 == int32(27) {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v1599)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_pfree(m, v2220)
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L288
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errcode_for_file_access(m)
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v1601
	F_errmsg(m, int32(314405), v31)
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errfinish(m, int32(525932), int32(4529), int32(406303))
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L287
	}
L287:
	;
	goto L3
L288:
	;
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v1603)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2244 = F_close(m, v2233)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1603))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2258 = F_durable_rename(m, v1601, v1600, int32(21))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v1607
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v1606
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v1602)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_FreeConfigVariables(m, v2264)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2288 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v2288+int32(4480))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L291
	}
L291:
	;
	m.G0 = v31 + int32(256)
	return
L292:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v1603)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2314 = F_close(m, v2303)
	mBase = m.M
	goto L294
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	v2326 = F_unlink(m, v1601)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_pg_re_throw(m)
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L295
	}
L295:
	;
	goto L3
L296:
	;
	*(*int32)(unsafe.Add(mBase, _consts[43])) = int32(51)
	goto L298
L297:
	;
	goto L298
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errcode_for_file_access(m)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v1601
	F_errmsg(m, int32(313140), v31+int32(32))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+220)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v31)+216)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v31)+224)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v31)+228)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v31)+232)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v31)+236)) = v1600
	*(*int32)(unsafe.Add(mBase, uint32(v31)+240)) = v1603
	*(*int32)(unsafe.Add(mBase, uint32(v31)+244)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v31)+248)) = v1602
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1599
	F_errfinish(m, int32(525932), int32(4490), int32(406303))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		v2424 = v1618
		goto L6
	} else {
		goto L302
	}
L302:
	;
	goto L5
L303:
	;
	v2432 = int32(v2428)
	m.G0 = v2424
	v2434 = *(*int32)(unsafe.Add(mBase, uint32(v2432)+4))
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2432)))
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2435)))
	if v31+int32(212) == v2439 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	m.ExcPending = 1
	goto L312
L305:
	;
	if v2442 != 0 {
		goto L309
	} else {
		goto L310
	}
L306:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2435)+4))
	v2442 = v2441
	goto L308
L307:
	;
	v2442 = int32(0)
	goto L308
L308:
	;
	goto L305
L309:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v31)+252))
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v31)+248))
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v31)+244))
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v31)+240))
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v31)+236))
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v31)+232))
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v31)+228))
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v31)+224))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v31)+220))
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v31)+216))
	v32 = v2443
	v33 = v2447
	v34 = v2448
	v35 = v2444
	v36 = v2446
	v37 = v2445
	v38 = v2449
	v39 = v2451
	v40 = v2452
	v41 = v2450
	v42 = v2434
	v43 = v2442
	v51 = v2424
	goto L1
L310:
	;
	goto L311
L311:
	;
	F___wasm_longjmp(m, v2435, v2434)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	return
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecAlterExtensionContentsRecurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v624 int32
	_ = v624
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v818 int32
	_ = v818
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int64
	_ = v899
	var v900 int32
	_ = v900
	var v903 int64
	_ = v903
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v923 int64
	_ = v923
	var v924 int32
	_ = v924
	var v927 int64
	_ = v927
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v955 int64
	_ = v955
	var v956 int32
	_ = v956
	var v959 int64
	_ = v959
	var v966 int32
	_ = v966
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
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
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	v16 = m.G0
	v18 = v16 - int32(288)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v22 = F_getExtensionOfObject(m, v20, v21)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v24 {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L257
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L254
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L251
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L248
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L245
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L242
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L239
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L234
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L229
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L223
	}
L13:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v886 == int32(1247) {
		goto L204
	} else {
		goto L205
	}
L14:
	;
	if v22 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 != v280 {
		goto L10
	} else {
		goto L102
	}
L17:
	;
	if v20 == int32(2615) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = F_SearchSysCache1(m, int32(28), v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_recordDependencyOn(m, l2, l1, int32(101))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	if v21 == v42 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	if v31 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v42 = int32(0)
	goto L21
L24:
	;
	goto L25
L25:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37)+72))
	F_ReleaseCatCache(m, v31)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v42 = v39
	goto L21
L27:
	;
	goto L20
L28:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v51 = m.G0
	v53 = v51 - int32(112)
	m.G0 = v53
	if v50 != int32(2613) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L13
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L98
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L95
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L92
	}
L33:
	;
	m.G0 = v53 + int32(112)
	goto L29
L34:
	;
	v190 = F_get_object_attnum_acl(m, v50)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L79
	}
L35:
	;
	if v50 != int32(1259) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v152 = F_table_open(m, int32(2995), int32(3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L67
	}
L38:
	;
	v60 = F_SearchSysCache1(m, int32(57), v49)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v60 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	switch v67 - int32(73) {
	case 0, 26, 32:
		goto L43
	default:
		goto L42
	case 10:
		goto L41
	}
L41:
	;
	v137 = F_SysCacheGetAttr(m, int32(57), v60, int32(32), v53+int32(48))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L60
	}
L42:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66)+120)))
	if v73 <= int32(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	F_ReleaseCatCache(m, v60)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L33
L45:
	;
	v80 = int32(1)
	goto L46
L46:
	;
	v92 = F_SearchSysCache2(m, int32(7), v49, v80)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	if v92 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+22)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v95)+91)))
	if v97 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v116 = base.I32_extend16_s(v80 + int32(1))
	if v116 <= v73 {
		v80 = v116
		goto L46
	} else {
		goto L59
	}
L52:
	;
	F_ReleaseCatCache(m, v92)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L58
	}
L53:
	;
	v102 = F_SysCacheGetAttr(m, int32(7), v92, int32(22), v53+int32(48))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+48)))
	if v104 != 0 {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v106 = F_pg_detoast_datum(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_recordExtensionInitPrivWorker(m, v49, int32(1259), v80, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	goto L51
L59:
	;
	goto L47
L60:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+48)))
	if v139 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v144 = F_pg_detoast_datum(m, v137)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_ReleaseCatCache(m, v60)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	F_recordExtensionInitPrivWorker(m, v49, int32(1259), int32(0), v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L33
L67:
	;
	F_ScanKeyInit(m, v53+int32(48), int32(1), int32(3), int32(184), v49)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v162 = int32(1)
	v167 = F_systable_beginscan(m, v152, int32(2996), v162, int32(0), v162, v53+int32(48))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v169 = F_systable_getnext(m, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v169 == int32(0) {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v152)+52))
	v177 = F_heap_getattr_2(m, v169, int32(3), v174, v53+int32(111))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+111)))
	if v179 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v184 = F_pg_detoast_datum(m, v177)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_systable_endscan(m, v167)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	F_recordExtensionInitPrivWorker(m, v49, int32(2613), int32(0), v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L33
L79:
	;
	if v190 == int32(0) {
		goto L33
	} else {
		goto L80
	}
L80:
	;
	v194 = F_get_object_catcache_oid(m, v50)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v196 = F_SearchSysCache1(m, v194, v49)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v196 == int32(0) {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	v200 = F_get_object_attnum_acl(m, v50)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v204 = F_SysCacheGetAttr(m, v194, v196, v200, v53+int32(48))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+48)))
	if v206 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v210 = F_pg_detoast_datum(m, v204)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_ReleaseCatCache(m, v196)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	F_recordExtensionInitPrivWorker(m, v49, v50, int32(0), v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L33
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v49
	F_errmsg_internal(m, int32(50136), v53+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(523480), int32(4373), int32(36925))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+32)) = v49
	F_errmsg_internal(m, int32(45807), v53+int32(32))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(523480), int32(4474), int32(36925))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v268 = F_get_object_class_descr(m, v50)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v268
	F_errmsg_internal(m, int32(46411), v53)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(523480), int32(4499), int32(36925))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v284 = F_deleteDependencyRecordsForClass(m, v20, v21, int32(3079), int32(101))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v284 != int32(1) {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v288 == int32(1259) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v295 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v741 = v288
	goto L107
L107:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v753 = m.G0
	v755 = v753 - int32(16)
	m.G0 = v755
	if v741 == int32(1259) {
		goto L180
	} else {
		goto L181
	}
L108:
	;
	F_ScanKeyInit(m, v18+int32(240), int32(1), int32(3), int32(184), v292)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v305 = int32(1)
	v310 = F_systable_beginscan(m, v295, int32(3080), v305, int32(0), v305, v18+int32(240))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v312 = F_systable_getnext(m, v310)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v312 == int32(0) {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v295)+52))
	v320 = F_heap_getattr_7(m, v312, int32(7), v317, v18+int32(239))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+239)))
	if v322 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_systable_endscan(m, v310)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L175
	}
L115:
	;
	v323 = F_pg_detoast_datum(m, v320)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v325 != int32(1) {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v323)+20))
	if v328 != int32(1) {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	if v331 < int32(0) {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	if v334 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	if v335 != int32(26) {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if v331 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v347 = int32(0)
	goto L123
L123:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v323+int32(24)+v347<<(uint(int32(2))%32))))
	if v291 != v361 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v366 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+216)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+208)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+200)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+192)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+184)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+176)) = int64(72339069014638592)
	if v331 == int32(1) {
		goto L130
	} else {
		goto L131
	}
L125:
	;
	v364 = v347 + int32(1)
	if v364 != v331 {
		v347 = v364
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L114
L129:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v295)+52))
	v533 = F_heap_getattr_7(m, v312, int32(8), v530, v18+int32(239))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L147
	}
L130:
	;
	v380 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+190)) = uint8(v380)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v382 = int32(0)
	F_deconstruct_array_builtin(m, v323, int32(26), v18+int32(172), v382, v18+int32(168))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v392 = v331 - int32(1)
	if v392 <= v347 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v511 = F_construct_array_builtin(m, v509, v392, int32(26))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L146
	}
L135:
	;
	v396 = (v392 - v347) & int32(3)
	if v396 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v402 = v347
	v404 = v382
	goto L139
L137:
	;
	v431 = v347
	goto L138
L138:
	;
	if base.Ui32(v331-v347-int32(2)) < base.Ui32(int32(3)) {
		goto L134
	} else {
		goto L142
	}
L139:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v413 = int32(2)
	v416 = int32(1)
	v417 = v402 + v416
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v412+v417<<(uint(v413)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v402<<(uint(v413)%32)))) = v421
	v424 = v404 + v416
	if v424 != v396 {
		v402 = v417
		v404 = v424
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v431 = v417
	goto L138
L141:
	;
	goto L140
L142:
	;
	v451 = v431
	goto L143
L143:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v462 = int32(2)
	v463 = v451 << (uint(v462) % 32)
	v465 = int32(4)
	v466 = v463 + v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v461+v466)))
	*(*int32)(unsafe.Add(mBase, uint32(v461+v463))) = v468
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v473 = v463 + int32(8)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v470+v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v470+v466))) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v480 = v463 + int32(12)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v477+v480)))
	*(*int32)(unsafe.Add(mBase, uint32(v477+v473))) = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v487 = v451 + v465
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v484+v487<<(uint(v462)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v484+v480))) = v491
	if v487 != v392 {
		v451 = v487
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L134
L145:
	;
	goto L144
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+216)) = v511
	goto L129
L147:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+239)))
	if v535 == int32(1) {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v538 = F_pg_detoast_datum(m, v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v540 != int32(1) {
		goto L5
	} else {
		goto L150
	}
L150:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	if v543 != int32(1) {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v538)+8))
	if v546 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	if v547 != int32(25) {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v538)+16))
	if v550 != v331 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	if v331 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v295)+52))
	v710 = F_heap_modify_tuple(m, v312, v703, v18+int32(192), v18+int32(184), v18+int32(176))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L173
	}
L156:
	;
	v554 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+191)) = uint8(v554)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v556 = int32(0)
	F_deconstruct_array_builtin(m, v538, int32(25), v18+int32(172), v556, v18+int32(168))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v566 = v331 - int32(1)
	if v566 <= v347 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v685 = F_construct_array_builtin(m, v683, v566, int32(25))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L172
	}
L161:
	;
	v573 = (v566 - v347) & int32(3)
	if v573 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v578 = v347
	v579 = v556
	goto L165
L163:
	;
	v607 = v347
	goto L164
L164:
	;
	if base.Ui32(v331-v347-int32(2)) < base.Ui32(int32(3)) {
		goto L160
	} else {
		goto L168
	}
L165:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v590 = int32(2)
	v593 = int32(1)
	v594 = v578 + v593
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v589+v594<<(uint(v590)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v589+v578<<(uint(v590)%32)))) = v598
	v601 = v579 + v593
	if v601 != v573 {
		v578 = v594
		v579 = v601
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v607 = v594
	goto L164
L167:
	;
	goto L166
L168:
	;
	v624 = v607
	goto L169
L169:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v636 = int32(2)
	v637 = v624 << (uint(v636) % 32)
	v639 = int32(4)
	v640 = v637 + v639
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v635+v640)))
	*(*int32)(unsafe.Add(mBase, uint32(v635+v637))) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v647 = v637 + int32(8)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v644+v647)))
	*(*int32)(unsafe.Add(mBase, uint32(v644+v640))) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v654 = v637 + int32(12)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v651+v654)))
	*(*int32)(unsafe.Add(mBase, uint32(v651+v647))) = v656
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v661 = v624 + v639
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v658+v661<<(uint(v636)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v658+v654))) = v665
	if v661 != v566 {
		v624 = v661
		goto L169
	} else {
		goto L171
	}
L170:
	;
	goto L160
L171:
	;
	goto L170
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v685
	goto L155
L173:
	;
	F_CatalogTupleUpdate(m, v295, v710+int32(4), v710)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	goto L114
L175:
	;
	F_sequence_close(m, v295, int32(3))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v741 = v736
	goto L107
L177:
	;
	m.G0 = v755 + int32(16)
	goto L13
L178:
	;
	F_ReleaseCatCache(m, v760)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L202
	}
L179:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L199
	}
L180:
	;
	v760 = F_SearchSysCache1(m, int32(57), v752)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v834 = int32(0)
	F_recordExtensionInitPrivWorker(m, v752, v741, v834, v834)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L198
	}
L183:
	;
	if v760 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v760)+16))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764)+22)))
	v766 = v764 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+119)))
	switch v767 - int32(73) {
	case 0, 26, 32:
		goto L178
	default:
		goto L186
	case 10:
		goto L185
	}
L185:
	;
	F_ReleaseCatCache(m, v760)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L197
	}
L186:
	;
	v771 = int32(*(*int16)(unsafe.Add(mBase, uint32(v766)+120)))
	if v771 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v779 = int32(1)
	goto L188
L188:
	;
	v790 = F_SearchSysCache2(m, int32(7), v752, v779)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	goto L185
L190:
	;
	if v790 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_recordExtensionInitPrivWorker(m, v752, int32(1259), v779, int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v800 = base.I32_extend16_s(v779 + int32(1))
	if v800 <= v771 {
		v779 = v800
		goto L188
	} else {
		goto L196
	}
L194:
	;
	F_ReleaseCatCache(m, v790)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	goto L189
L197:
	;
	goto L182
L198:
	;
	goto L177
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = v752
	F_errmsg_internal(m, int32(50136), v755)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(523480), int32(4533), int32(36904))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
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
	goto L177
L203:
	;
	m.G0 = v18 + int32(288)
	return
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+240)) = int32(1247)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v894 = F_get_array_type(m, v893)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	v940 = v886
	goto L206
L206:
	;
	if v940 != int32(1259) {
		goto L203
	} else {
		goto L219
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+244)) = v894
	if v894 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v897
	v899 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v18)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v900
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v899
	v903 = *(*int64)(unsafe.Add(mBase, uint32(v18)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v903
	F_ExecAlterExtensionContentsRecurse(m, l0, v18+int32(112), v18+int32(96))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v913 = F_type_is_range(m, v912)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L212
	}
L211:
	;
	goto L210
L212:
	;
	if v913 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v916 = F_get_range_multirange(m, v915)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v940 = v937
	goto L206
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+244)) = v916
	if v916 == int32(0) {
		goto L3
	} else {
		goto L217
	}
L217:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v921
	v923 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v18)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v924
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v923
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v18)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v927
	F_ExecAlterExtensionContentsRecurse(m, l0, v18+int32(80), v18-int32(-64))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+240)) = int32(1247)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v948 = F_get_rel_type_id(m, v947)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+244)) = v948
	if v948 == int32(0) {
		goto L203
	} else {
		goto L221
	}
L221:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v953
	v955 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v18)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v956
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v955
	v959 = *(*int64)(unsafe.Add(mBase, uint32(v18)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v959
	F_ExecAlterExtensionContentsRecurse(m, l0, v18+int32(48), v18+int32(32))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	goto L203
L223:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v980 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v982 = F_get_extension_name(m, v22)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v982
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v980
	F_errmsg(m, int32(745451), v18+int32(128))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(521848), int32(3820), int32(379393))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1004 = F_get_namespace_name(m, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1004
	F_errmsg(m, int32(285179), v18)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(521848), int32(3833), int32(379393))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
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
	F_errcode(m, int32(325))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v1025 = F_getObjectDescription(m, l2, int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+164)) = v1027
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v1025
	F_errmsg(m, int32(745492), v18+int32(160))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(521848), int32(3860), int32(379393))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_errmsg_internal(m, int32(182832), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(521848), int32(3868), int32(379393))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v292
	F_errmsg_internal(m, int32(50745), v18+int32(144))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(521848), int32(3058), int32(360690))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errmsg_internal(m, int32(26590), int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(521848), int32(3083), int32(360690))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(342177), int32(0))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(521848), int32(3140), int32(360690))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	F_errmsg_internal(m, int32(25364), int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(521848), int32(3150), int32(360690))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F_errmsg_internal(m, int32(342177), int32(0))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(521848), int32(3152), int32(360690))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1129 = F_format_type_be(m, v1128)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1129
	F_errmsg(m, int32(204478), v18+int32(16))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(521848), int32(3913), int32(379393))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecAlterObjectSchemaStmt(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v303 int32
	_ = v303
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v690 int64
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v925 int32
	_ = v925
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v26 - int32(1) {
	case 0, 6, 7, 18, 23, 24, 25, 28, 33, 38, 44, 45, 46, 47:
		goto L3
	default:
		goto L4
	case 11, 48:
		goto L5
	case 14:
		goto L7
	case 17, 22, 36, 40, 50:
		goto L6
	}
L1:
	;
	if l2 != 0 {
		goto L244
	} else {
		goto L245
	}
L2:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v905 = v900
	v906 = v899
	v908 = v901
	goto L1
L3:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v858 = int32(0)
	F_get_object_address(m, v24+int32(16), v26, v857, v858, int32(8), v858)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L19
	} else {
		goto L239
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L19
	} else {
		goto L236
	}
L5:
	;
	v783 = v24 + int32(16)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L212
	} else {
		goto L213
	}
L6:
	;
	v657 = v24 + int32(16)
	if l2 != 0 {
		goto L172
	} else {
		goto L173
	}
L7:
	;
	v30 = v24 + int32(16)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v24 + int32(28)
	goto L10
L9:
	;
	v37 = int32(0)
	goto L10
L10:
	;
	v38 = m.G0
	v40 = v38 - int32(240)
	m.G0 = v40
	v43 = int32(0)
	v46 = F_GetSysCacheOid(m, int32(27), v32, v43, v43, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	goto L2
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L19
	} else {
		goto L167
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L19
	} else {
		goto L164
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L19
	} else {
		goto L157
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L19
	} else {
		goto L154
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L19
	} else {
		goto L150
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L19
	} else {
		goto L147
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L19
	} else {
		goto L143
	}
L19:
	;
	return
L20:
	;
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v48 = F_LookupCreationNamespace(m, v33)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L19
	} else {
		goto L139
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v53 = F_object_ownercheck(m, int32(3079), v46, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	if v53 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_aclcheck_error(m, int32(2), int32(15), v32)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L19
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v65 = F_object_aclcheck(m, int32(2615), v48, v63, int64(512))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v65 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_aclcheck_error(m, v65, int32(36), v33)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v71 = F_getExtensionOfObject(m, int32(2615), v48)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L19
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if v71 == v46 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v76 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_ScanKeyInit(m, v40+int32(144), int32(1), int32(3), int32(184), v46)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v86 = int32(1)
	v91 = F_systable_beginscan(m, v76, int32(3080), v86, int32(0), v86, v40+int32(144))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v93 = F_systable_getnext(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	if v93 == int32(0) {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	v97 = F_heap_copytuple(m, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+22)))
	F_systable_endscan(m, v91)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v103 = v99 + v100
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+72))
	if v48 == v104 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	m.G0 = v40 + int32(240)
	goto L11
L45:
	;
	F_sequence_close(m, v76, int32(3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L19
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+76)))
	if v115 == int32(0) {
		goto L16
	} else {
		goto L49
	}
L48:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v110
	v113 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v113
	goto L44
L49:
	;
	v118 = F_new_object_addresses(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v103)+72))
	v123 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	F_ScanKeyInit(m, v40+int32(144), int32(4), int32(3), int32(184), int32(3079))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L52
	}
L52:
	;
	F_ScanKeyInit(m, v40+int32(192), int32(5), int32(3), int32(184), v46)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	v146 = F_systable_beginscan(m, v123, int32(2674), int32(1), int32(0), int32(2), v40+int32(144))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	v148 = F_systable_getnext(m, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	if v148 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v151 = v103 + int32(4)
	v153 = v148
	goto L59
L57:
	;
	goto L58
L58:
	;
	if v37 != 0 {
		goto L126
	} else {
		goto L127
	}
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+22)))
	v175 = v173 + v174
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+24)))
	if v176 == int32(110) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L58
L61:
	;
	v419 = F_systable_getnext(m, v146)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L19
	} else {
		goto L124
	}
L62:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v179 != int32(3079) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	v325 = v176
	goto L64
L64:
	;
	if v325&int32(255) != int32(101) {
		goto L61
	} else {
		goto L94
	}
L65:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v184 = F_SearchSysCache1(m, int32(28), v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L19
	} else {
		goto L67
	}
L66:
	;
	v201 = F_palloc0(m, int32(48))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L19
	} else {
		goto L73
	}
L67:
	;
	if v184 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v199 = int32(0)
	goto L66
L69:
	;
	goto L70
L70:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v194 = F_pstrdup(m, v189+v190+int32(4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L19
	} else {
		goto L71
	}
L71:
	;
	F_ReleaseCatCache(m, v184)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L19
	} else {
		goto L72
	}
L72:
	;
	v199 = v194
	goto L66
L73:
	;
	v203 = F_pstrdup(m, v199)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L19
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+36)) = int32(-1)
	v207 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+34)) = uint8(v207)
	v209 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v201)+32)) = uint16(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v203
	F_parse_extension_control_file(m, v201, v207)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v201)+44))
	if v215 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+24)))
	v325 = v303
	goto L64
L77:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v218 <= int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v221 = int32(0)
	if v221 < v218 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v224 = v218
	goto L81
L80:
	;
	v224 = v221
	goto L81
L81:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v228 = int32(0)
	goto L82
L82:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v225+v228<<(uint(int32(2))%32))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v255 == int32(0) {
		v274 = v254
		v275 = v255
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L76
L84:
	;
	if v275-v274 == int32(0) {
		goto L12
	} else {
		goto L92
	}
L85:
	;
	goto L84
L86:
	;
	if v254 != v255 {
		v274 = v254
		v275 = v255
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v259 = v251
	v260 = v151
	goto L88
L88:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	if v264 == int32(0) {
		v274 = v263
		v275 = v264
		goto L85
	} else {
		goto L90
	}
L89:
	;
	v274 = v263
	v275 = v264
	goto L85
L90:
	;
	v267 = int32(1)
	if v263 == v264 {
		v259 = v259 + v267
		v260 = v260 + v267
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v280 = v228 + int32(1)
	if v280 != v224 {
		v228 = v280
		goto L82
	} else {
		goto L93
	}
L93:
	;
	goto L83
L94:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+132)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+136)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+140)) = v334
	if v334 != 0 {
		goto L15
	} else {
		goto L95
	}
L95:
	;
	v336 = int32(0)
	if v330 <= int32(2752) {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	if v394 == int32(0) {
		goto L61
	} else {
		goto L122
	}
L97:
	;
	v382 = F_table_open(m, v330, int32(3))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L19
	} else {
		goto L119
	}
L98:
	;
	v394 = v379
	goto L96
L99:
	;
	v376 = F_AlterTypeNamespace_oid(m, v332, v48, int32(1), v118)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L19
	} else {
		goto L118
	}
L100:
	;
	v366 = F_relation_open(m, v332, int32(8))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L19
	} else {
		goto L115
	}
L101:
	;
	switch v330 - int32(1247) {
	case 0:
		goto L99
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11:
		v379 = v336
		goto L98
	case 8:
		goto L97
	case 12:
		goto L100
	default:
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v330 <= int32(3599) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v342 = v330 - int32(2607)
	if base.Ui32(int32(10)) < base.Ui32(v342) {
		v379 = v336
		goto L98
	} else {
		goto L105
	}
L105:
	;
	if int32(1)<<(uint(v342)%32)&int32(1537) == int32(0) {
		v379 = v336
		goto L98
	} else {
		goto L106
	}
L106:
	;
	goto L97
L107:
	;
	if v330 == int32(2753) {
		goto L97
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	if base.Ui32(v330-int32(3600)) < base.Ui32(int32(3)) {
		goto L97
	} else {
		goto L113
	}
L110:
	;
	if v330 == int32(3381) {
		goto L97
	} else {
		goto L111
	}
L111:
	;
	if v330 == int32(3456) {
		goto L97
	} else {
		goto L112
	}
L112:
	;
	v379 = v336
	goto L98
L113:
	;
	if v330 == int32(3764) {
		goto L97
	} else {
		goto L114
	}
L114:
	;
	v379 = v336
	goto L98
L115:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v366)+48))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+68))
	F_AlterTableNamespaceInternal(m, v366, v369, v48, v118)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	F_relation_close(m, v366, int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	v394 = v369
	goto L96
L118:
	;
	v379 = v376
	goto L98
L119:
	;
	v384 = F_AlterObjectNamespace_internal(m, v382, v332, v48)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	F_sequence_close(m, v382, int32(3))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L19
	} else {
		goto L121
	}
L121:
	;
	v394 = v384
	goto L96
L122:
	;
	if v394 != v120 {
		goto L14
	} else {
		goto L123
	}
L123:
	;
	goto L61
L124:
	;
	if v419 != 0 {
		v153 = v419
		goto L59
	} else {
		goto L125
	}
L125:
	;
	goto L60
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v120
	goto L128
L127:
	;
	goto L128
L128:
	;
	F_systable_endscan(m, v146)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L19
	} else {
		goto L129
	}
L129:
	;
	F_relation_close(m, v123, int32(1))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L19
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+72)) = v48
	F_CatalogTupleUpdate(m, v76, v97+int32(4), v97)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	F_sequence_close(m, v76, int32(3))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	v458 = F_changeDependencyFor(m, int32(3079), v46, int32(2615), v120, v48)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L19
	} else {
		goto L133
	}
L133:
	;
	if v458 != int32(1) {
		goto L13
	} else {
		goto L134
	}
L134:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	if v463 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v465 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), v46, v465, v465, v465)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L19
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(3079)
	goto L44
L138:
	;
	goto L137
L139:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v32
	F_errmsg(m, int32(77888), v40)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L19
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(521848), int32(199), int32(456147))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L19
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v32
	F_errmsg(m, int32(533315), v40+int32(16))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L19
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(521848), int32(3236), int32(439940))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L19
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v46
	F_errmsg_internal(m, int32(50745), v40+int32(32))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L19
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(521848), int32(3253), int32(439940))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L19
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L19
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+128)) = v103 + int32(4)
	F_errmsg(m, int32(572835), v40+int32(128))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L19
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(521848), int32(3276), int32(439940))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L19
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errmsg_internal(m, int32(23806), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L19
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(521848), int32(3354), int32(439940))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L19
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = v151
	F_errmsg(m, int32(572835), v40+int32(80))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	v598 = F_getObjectDescription(m, v40+int32(132), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	v600 = F_get_namespace_name(m, v120)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v598
	F_errdetail(m, int32(759402), v40-int32(-64))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(521848), int32(3373), int32(439940))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v103 + int32(4)
	F_errmsg_internal(m, int32(195238), v40+int32(48))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(521848), int32(3395), int32(439940))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+112)) = v151
	F_errmsg(m, int32(110702), v40+int32(112))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+100)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = v199
	F_errdetail(m, int32(698501), v40+int32(96))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L19
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(521848), int32(3336), int32(439940))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L19
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	v661 = v24 + int32(28)
	goto L174
L173:
	;
	v661 = int32(0)
	goto L174
L174:
	;
	v662 = m.G0
	v664 = v662 - int32(32)
	m.G0 = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v670 = F_RangeVarGetRelidExtended(m, v666, int32(8), v668, int32(576), l1)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L19
	} else {
		goto L178
	}
L175:
	;
	goto L2
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L19
	} else {
		goto L206
	}
L177:
	;
	m.G0 = v664 + int32(32)
	goto L175
L178:
	;
	if v670 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v676 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L19
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v696 = F_relation_open(m, v670, int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L19
	} else {
		goto L188
	}
L182:
	;
	if v676 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v664))) = v679
	F_errmsg(m, int32(350183), v664)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L19
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v690 = *(*int64)(unsafe.Add(mBase, _consts[458]))
	*(*int64)(unsafe.Add(mBase, uint32(v657))) = v690
	v693 = *(*int32)(unsafe.Add(mBase, _consts[457]))
	*(*int32)(unsafe.Add(mBase, uint32(v657)+8)) = v693
	goto L177
L186:
	;
	F_errfinish(m, int32(519627), int32(18965), int32(440091))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L19
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v696)+48))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+68))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698)+119)))
	if v700 == int32(83) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v708 = F_sequenceIsOwned(m, v670, int32(97), v664+int32(28), v664+int32(24))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L19
	} else {
		goto L192
	}
L190:
	;
	v718 = v698
	goto L191
L191:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v723 = F_makeRangeVar(m, v719, v718+int32(4), int32(-1))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L19
	} else {
		goto L196
	}
L192:
	;
	if v708 != 0 {
		goto L176
	} else {
		goto L193
	}
L193:
	;
	v715 = F_sequenceIsOwned(m, v670, int32(105), v664+int32(28), v664+int32(24))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L19
	} else {
		goto L194
	}
L194:
	;
	if v715 != 0 {
		goto L176
	} else {
		goto L195
	}
L195:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v696)+48))
	v718 = v717
	goto L191
L196:
	;
	v725 = int32(0)
	v727 = F_RangeVarGetAndCheckCreationNamespace(m, v723, v725, v725)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L19
	} else {
		goto L197
	}
L197:
	;
	F_CheckSetNamespace(m, v699, v727)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L19
	} else {
		goto L198
	}
L198:
	;
	v731 = F_new_object_addresses(m)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L19
	} else {
		goto L199
	}
L199:
	;
	F_AlterTableNamespaceInternal(m, v696, v699, v727, v731)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L19
	} else {
		goto L200
	}
L200:
	;
	F_free_object_addresses(m, v731)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L19
	} else {
		goto L201
	}
L201:
	;
	if v661 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v699
	goto L204
L203:
	;
	goto L204
L204:
	;
	F_relation_close(m, v696, int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L19
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v657)+4)) = v670
	*(*int32)(unsafe.Add(mBase, uint32(v657))) = int32(1259)
	goto L177
L206:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L19
	} else {
		goto L207
	}
L207:
	;
	F_errmsg(m, int32(533265), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L19
	} else {
		goto L208
	}
L208:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v696)+48))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v664)+28))
	v766 = F_get_rel_name(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L19
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v664)+20)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v664)+16)) = v764 + int32(4)
	F_errdetail(m, int32(699250), v664+int32(16))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L19
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(519627), int32(18986), int32(440091))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L19
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	v789 = v24 + int32(28)
	goto L214
L213:
	;
	v789 = int32(0)
	goto L214
L214:
	;
	v790 = m.G0
	v792 = v790 - int32(16)
	m.G0 = v792
	v795 = F_makeTypeNameFromNameList(m, v784)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v797 = F_typenameTypeId(m, int32(0), v795)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L19
	} else {
		goto L216
	}
L216:
	;
	if v26 == int32(12) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L2
L218:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L19
	} else {
		goto L231
	}
L219:
	;
	v801 = F_get_typtype(m, v797)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L19
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v805 = F_LookupCreationNamespace(m, v785)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L19
	} else {
		goto L224
	}
L222:
	;
	if v801 != int32(100) {
		goto L218
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v808 = F_new_object_addresses(m)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L19
	} else {
		goto L225
	}
L225:
	;
	v810 = F_AlterTypeNamespace_oid(m, v797, v805, int32(0), v808)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L19
	} else {
		goto L226
	}
L226:
	;
	F_free_object_addresses(m, v808)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L19
	} else {
		goto L227
	}
L227:
	;
	if v789 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v789))) = v810
	goto L230
L229:
	;
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v783)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v783)+4)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = int32(1247)
	m.G0 = v792 + int32(16)
	goto L217
L231:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L19
	} else {
		goto L232
	}
L232:
	;
	v830 = F_format_type_be(m, v797)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L19
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v792))) = v830
	F_errmsg(m, int32(291651), v792)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L19
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(519616), int32(4073), int32(440072))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L19
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v845
	F_errmsg_internal(m, int32(508571), v24)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L19
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(520586), int32(600), int32(104507))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L19
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v867 = F_table_open(m, v865, int32(3))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L19
	} else {
		goto L240
	}
L240:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v870 = F_LookupCreationNamespace(m, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L19
	} else {
		goto L241
	}
L241:
	;
	v872 = F_AlterObjectNamespace_internal(m, v867, v864, v870)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L19
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v872
	F_sequence_close(m, v867, int32(3))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L19
	} else {
		goto L243
	}
L243:
	;
	v905 = v864
	v906 = v863
	v908 = v865
	goto L1
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2615)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v925
	goto L246
L245:
	;
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v906
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v908
	m.G0 = v24 + int32(32)
	return
}
func F__equalAlterTableSpaceOptionsStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v51
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = F_equal(m, v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v51 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 != 0 {
		v51 = v3
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	return int32(0)
L18:
	;
	if v40 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	v51 = base.B2i32(v48 == v49)
	goto L1
}
