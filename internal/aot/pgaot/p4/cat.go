package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatCacheRemoveCTup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+52)) = uint8(v9)
	F_CatCacheRemoveCList(m, l0, v8)
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+53)))
	if v18 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	F_pfree(m, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v21 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = int32(0)
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v32<<(uint(int32(2))%32))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v40<<(uint(int32(3))%32))+24)))
	if v44 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v32<<(uint(int32(3))%32))))
	F_pfree(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = v32 + int32(1)
	if v54 != v21 {
		v32 = v54
		goto L9
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L10
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v65 - v66
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_CatCacheRemoveCTup[0]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v71 - v66
	return
}
func F_SearchCatCache2(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = int64(0)
	v7 = F_SearchCatCacheInternal(m, l0, int32(2), l1, l2, v5, v5)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_SearchCatCacheList(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v164 int32
	_ = v164
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v325 int32
	_ = v325
	var v341 int32
	_ = v341
	var v376 int32
	_ = v376
	var v388 int32
	_ = v388
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v518 int32
	_ = v518
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v586 int32
	_ = v586
	var v588 int64
	_ = v588
	var v592 int32
	_ = v592
	var v606 int64
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v647 int32
	_ = v647
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v713 int32
	_ = v713
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
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v746 int32
	_ = v746
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v817 int32
	_ = v817
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v983 int32
	_ = v983
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1149 int32
	_ = v1149
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1264 int32
	_ = v1264
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1320 int32
	_ = v1320
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int64
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int64
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1400 int64
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1451 int32
	_ = v1451
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1487 int32
	_ = v1487
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1602 int32
	_ = v1602
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1795 int32
	_ = v1795
	var v1809 int32
	_ = v1809
	var v1824 int32
	_ = v1824
	var v1858 int32
	_ = v1858
	var v1859 int64
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	v6 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(592)
	m.G0 = v36
	v49 = l0
	v50 = l1
	v51 = l2
	v52 = l3
	v53 = l4
	v54 = v36
	v55 = v6
	v57 = v6
	v58 = v6
	v59 = v6
	v60 = v6
	v61 = v6
	v62 = v6
	v63 = v6
	v64 = v6
	v65 = v6
	v66 = v6
	v67 = int32(-1)
	v69 = l0 + int32(8)
	v75 = l0 + int32(104)
	v76 = l0 + int32(48)
	v77 = v36 + int32(437)
	v78 = l0 + int32(32)
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
	if v67 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v1858 = int32(m.ExcTag)
	v1859 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1858 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L6
	} else {
		goto L182
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v1746
	v1773 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerRemember(m, v1773, base.I64_extend_i32_u(v1736), int32(_a_F_SearchCatCacheList_0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L6
	} else {
		goto L181
	}
L9:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v518 != v1694 {
		goto L174
	} else {
		goto L175
	}
L10:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v84 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v713 = v55
	v715 = v57
	v716 = v58
	v717 = v59
	v718 = v60
	v719 = v61
	v720 = v62
	v721 = v63
	v724 = v66
	goto L12
L12:
	;
	if v713 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	F_CatalogCacheInitializeCache(m, v49)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54)+472)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+464)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v54)+456)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v54)+448)) = v51
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	if v104 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v424 = int32(0)
	switch v50 - int32(1) {
	case 0:
		v478 = v424
		goto L46
	case 1:
		v460 = v424
		goto L47
	case 2:
		v443 = v424
		goto L48
	case 3:
		goto L49
	default:
		goto L7
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	v120 = F_MemoryContextAllocZero(m, v118, int32(128))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v49)+72))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	if v125 <= v126<<(uint(int32(1))%32) {
		goto L17
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v120
	goto L17
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v142 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if v142 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v49)+84))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v49)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v144
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheList_1), v54+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	v198 = F_MemoryContextAllocZero(m, v195, v183<<(uint(int32(4))%32))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	F_errfinish(m, int32(_a_F_SearchCatCacheList_2), int32(1055), int32(_a_F_SearchCatCacheList_3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v201 = v183 << (uint(int32(1)) % 32)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	if int32(0) < v202 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v226 = v202
	v227 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	F_pfree(m, v376)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L45
	}
L33:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	v244 = v241 + v227<<(uint(int32(3))%32)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v246 = int32(0)
	if base.B2i32(v245 == v246)|base.B2i32(v245 == v244) == v246 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v258 = v245
	goto L38
L36:
	;
	v325 = v226
	goto L37
L37:
	;
	v341 = v227 + int32(1)
	if v341 < v325 {
		v226 = v325
		v227 = v341
		goto L33
	} else {
		goto L44
	}
L38:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v286)+4)) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v287))) = v289
	v294 = v198 + v285&(v201-int32(1))<<(uint(int32(3))%32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	if v295 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v325 = v306
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v294
	goto L42
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+4)) = v294
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v301)+4)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v258
	if v287 != v244 {
		v258 = v287
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	goto L34
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v201
	goto L17
L46:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v491 = m.T0[v480].(func(*base.Module, int64) int32)(m, v51)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L53
	}
L47:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v473 = m.T0[v462].(func(*base.Module, int64) int32)(m, v52)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L6
	} else {
		goto L52
	}
L48:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v455 = m.T0[v444].(func(*base.Module, int64) int32)(m, v53)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v439 = m.T0[v427].(func(*base.Module, int64) int32)(m, int64(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v443 = base.I32_rotl(v439, int32(24))
	goto L48
L51:
	;
	v460 = base.I32_rotl(v455, int32(16)) ^ v443
	goto L47
L52:
	;
	v478 = base.I32_rotl(v473, int32(8)) ^ v460
	goto L46
L53:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v49)+80))
	v494 = v478 ^ v491
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v501 = v493 + v494&(v495-int32(1))<<(uint(int32(3))%32)
	v503 = v501 + int32(4)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v501)+4))
	v505 = int32(0)
	if base.B2i32(v504 == v505)|base.B2i32(v504 == v501) == v505 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v518 = v504
	goto L57
L55:
	;
	goto L56
L56:
	;
	v682 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+444)) = v682
	v685 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+436)) = uint16(v685)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+432)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v54)+428)) = v49
	v689 = int32(_a_F_SearchCatCacheList_4)
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v54 + int32(428)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+440)) = v690
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3]))
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4]))
	goto L69
L57:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+52)))
	if v544 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v647 != v501 {
		v518 = v647
		goto L57
	} else {
		goto L68
	}
L60:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if v545 != v494 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v518)+54)))
	if v50 != v547 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v558 = int32(0)
	goto L63
L63:
	;
	v586 = v558 << (uint(int32(3)) % 32)
	v588 = *(*int64)(unsafe.Add(mBase, uint32(v518+int32(16)+v586)))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v78+v558<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v54+int32(448)+v586)))
	v607 = m.T0[v592].(func(*base.Module, int64, int64) int32)(m, v588, v606)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L65
	}
L64:
	;
	goto L9
L65:
	;
	if v607 == int32(0) {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v612 = v558 + int32(1)
	if v50 != v612 {
		v558 = v612
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L58
L69:
	;
	v701 = v54 + int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v54 + int32(44)
	goto L72
L70:
	;
	v713 = v682
	v715 = v501
	v716 = v77
	v717 = v503
	v718 = v690
	v719 = v494
	v720 = v699
	v721 = v697
	v724 = v69
	goto L12
L72:
	;
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v54 + int32(272)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v49)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v758 = F_table_open(m, v746, int32(1))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4])) = v720
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v721
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v718
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v54)+444))
	if v1577 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L76:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v762 = v760 * int32(56)
	if v762 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	base.MemoryCopy(m, v54+int32(48), v75, v762)
	goto L79
L78:
	;
	goto L79
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v54)+264)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v54)+208)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v54)+152)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v54)+96)) = v51
	v786 = v64
	v787 = v65
	goto L80
L80:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v54)+444))
	if v804 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	F_relation_close(m, v758, int32(1))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L6
	} else {
		goto L131
	}
L82:
	;
	v890 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+444)) = v890
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v890)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v49)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	switch v906 - int32(1) {
	case 0, 1:
		v916 = v890
		goto L89
	default:
		goto L90
	case 7, 9, 10, 20, 42, 43:
		goto L91
	case 33:
		goto L92
	}
L83:
	;
	v807 = int32(0)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v808 <= v807 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v817 = v807
	goto L85
L85:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v804)+12))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v844+v817<<(uint(int32(2))%32))))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+48))
	v850 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v848)+48)) = v849 - v850
	v854 = v817 + v850
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v854 < v855 {
		v817 = v854
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v930 = F_systable_beginscan(m, v758, v894, v916, int32(0), v50, v54+int32(48))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L6
	} else {
		goto L95
	}
L89:
	;
	goto L88
L90:
	;
	v916 = int32(1)
	goto L89
L91:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheList[5])))
	if v912 != int32(1) {
		v916 = v890
		goto L89
	} else {
		goto L94
	}
L92:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheList[6])))
	if v910 != 0 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	v916 = v890
	goto L89
L94:
	;
	goto L90
L95:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v930)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v786
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v943 = F_systable_getnext(m, v930)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L6
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	F_systable_endscan(m, v930)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L6
	} else {
		goto L129
	}
L97:
	;
	if v943 == int32(0) {
		v1221 = v786
		v1222 = v787
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	if v947&int32(1) != 0 {
		v1221 = v786
		v1222 = v787
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v957 = v943
	v965 = v786
	v966 = v787
	goto L100
L100:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v994 = F_CatalogCacheComputeTupleHashValue(m, v49, v983, v957)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L6
	} else {
		goto L102
	}
L101:
	;
	v1204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v716))) = uint8(v1204)
	v1221 = v1130
	v1222 = v966
	goto L96
L102:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v1000 = (v997 - int32(1)) & v994
	v1003 = v996 + v1000<<(uint(int32(3))%32)
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	v1005 = int32(0)
	if base.B2i32(v1004 == v1005)|base.B2i32(v1004 == v1003) == v1005 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+444))
	v1178 = F_lappend(m, v1177, v1140)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L6
	} else {
		goto L125
	}
L105:
	;
	v1012 = v957 + int32(4)
	v1019 = v1004
	goto L108
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1130 = F_CatalogCacheCreateEntry(m, v49, v957, int32(0), v994, v1000)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L6
	} else {
		goto L123
	}
L108:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+52)))
	if v1046 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L107
L110:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+4))
	if v1084 != v1003 {
		v1019 = v1084
		goto L108
	} else {
		goto L122
	}
L111:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+53)))
	if v1047 != 0 {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+12))
	if v1048 != v994 {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1061 = v1019 + int32(60)
	v1062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1061)+2)))
	v1063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1061))))
	v1064 = int32(16)
	v1067 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+2)))
	v1068 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012))))
	if v1062|v1063<<(uint(v1064)%32) == v1067|v1068<<(uint(v1064)%32) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	if v1078 == int32(0) {
		goto L110
	} else {
		goto L120
	}
L115:
	;
	goto L114
L116:
	;
	v1074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1061)+4)))
	v1075 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1012)+4)))
	if v1074 == v1075 {
		v1078 = int32(1)
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v1078 = int32(0)
	goto L115
L119:
	;
	goto L118
L120:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+76))
	if v1081 == int32(0) {
		v1140 = v1019
		v1149 = v965
		goto L104
	} else {
		goto L121
	}
L121:
	;
	goto L110
L122:
	;
	goto L109
L123:
	;
	if v1130 == int32(0) {
		goto L103
	} else {
		goto L124
	}
L124:
	;
	v1140 = v1130
	v1149 = v1130
	goto L104
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+444)) = v1178
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1140)+48)) = v1181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1195 = F_systable_getnext(m, v930)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	if v1195 == int32(0) {
		v1221 = v1149
		v1222 = v1195
		goto L96
	} else {
		goto L127
	}
L127:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	if v1199&int32(1) == int32(0) {
		v957 = v1195
		v965 = v1149
		v966 = v1195
		goto L100
	} else {
		goto L128
	}
L128:
	;
	v1221 = v1149
	v1222 = v1195
	goto L96
L129:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	if v1251 != 0 {
		v786 = v1221
		v787 = v1222
		goto L80
	} else {
		goto L130
	}
L130:
	;
	goto L81
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1276 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerEnlarge(m, v1276)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	v1280 = int32(_a_F_SearchCatCacheList_5)
	v1281 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7]))
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7])) = v1284
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v54)+444))
	if v1286 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	v1288 = v1287
	goto L135
L134:
	;
	v1288 = int32(0)
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1303 = F_palloc(m, v1288<<(uint(int32(2))%32)-int32(-64))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	if int32(0) < v50 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v1320 = int32(0)
	goto L140
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7])) = v1281
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v721
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4])) = v720
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+60)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+8)) = int32(1383485699)
	*(*uint16)(unsafe.Add(mBase, uint32(v1303)+54)) = uint16(v50)
	v1451 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+53)) = uint8(base.B2i32(v932 != v1451))
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+52)) = uint8(v1451)
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+48)) = v1451
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+56)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+12)) = v719
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v54)+444))
	if v1460 == v1451 {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	v1347 = int32(3)
	v1348 = v1320 << (uint(v1347) % 32)
	v1352 = *(*int64)(unsafe.Add(mBase, uint32(v1348+(v54+int32(448)))))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1309)))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v76+v1320<<(uint(int32(2))%32))))
	v1365 = v1309 + v1353<<(uint(v1347)%32) + v1360*int32(100) - int32(72)
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+68))
	if v1366 == int32(19) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1383 = F_strncpy(m, v54+int32(488), base.I32_wrap_i64(v1352), int32(64))
	mBase = m.M
	v1384 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1383)+63)) = uint8(v1384)
	goto L145
L143:
	;
	v1386 = v1352
	goto L144
L144:
	;
	v1387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1365)+72)))
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1365)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v1221
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v1222
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	v1400 = F_datumCopy(m, v1386, v1388, v1387)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L6
	} else {
		goto L146
	}
L145:
	;
	v1386 = base.I64_extend_i32_u(v54 + int32(488))
	goto L144
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1303+int32(16)+v1348))) = v1400
	v1404 = v1320 + int32(1)
	if v1404 != v50 {
		v1320 = v1404
		goto L140
	} else {
		goto L147
	}
L147:
	;
	goto L141
L148:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	if v1554 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v1463 = int32(0)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+4))
	if v1464 <= v1463 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v1487 = v1463
	goto L151
L151:
	;
	v1503 = v1487 << (uint(int32(2)) % 32)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1505+v1503)))
	*(*int32)(unsafe.Add(mBase, uint32(v1303-int32(-64)+v1503))) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+76)) = v1303
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+48)) = v1510 - int32(1)
	v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507)+52)))
	if v1514 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L148
L153:
	;
	v1515 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303)+52)) = uint8(v1515)
	goto L155
L154:
	;
	goto L155
L155:
	;
	v1518 = v1487 + int32(1)
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+4))
	if v1518 < v1519 {
		v1487 = v1518
		goto L151
	} else {
		goto L156
	}
L156:
	;
	goto L152
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v715))) = v715
	v1558 = v715
	goto L159
L158:
	;
	v1558 = v1554
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1303))) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+4)) = v1558
	*(*int32)(unsafe.Add(mBase, uint32(v1558))) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v717))) = v1303
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v49)+72))
	v1564 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+72)) = v1563 + v1564
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1303)+48)) = v1567 + v1564
	v1736 = v1303
	v1737 = v715
	v1738 = v716
	v1739 = v717
	v1740 = v718
	v1741 = v719
	v1742 = v720
	v1743 = v721
	v1744 = v1221
	v1745 = v1222
	v1746 = v724
	goto L8
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	F_pg_re_throw(m)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L6
	} else {
		goto L173
	}
L161:
	;
	v1580 = int32(0)
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+4))
	if v1581 <= v1580 {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v1602 = v1580
	goto L163
L163:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+12))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1617+v1602<<(uint(int32(2))%32))))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+48))
	v1623 = int32(1)
	v1624 = v1622 - v1623
	*(*int32)(unsafe.Add(mBase, uint32(v1621)+48)) = v1624
	v1626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1621)+52)))
	if base.B2i32(v1626 != v1623)|v1624 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L160
L165:
	;
	v1646 = v1602 + int32(1)
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1577)+4))
	if v1646 < v1647 {
		v1602 = v1646
		goto L163
	} else {
		goto L172
	}
L166:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+76))
	if v1630 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+48))
	if v1631 != 0 {
		goto L165
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v717
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v715
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v724
	F_CatCacheRemoveCTup(m, v49, v1621)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L6
	} else {
		goto L171
	}
L170:
	;
	goto L169
L171:
	;
	goto L165
L172:
	;
	goto L164
L173:
	;
	goto L3
L174:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1696)+4)) = v1697
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	*(*int32)(unsafe.Add(mBase, uint32(v1697))) = v1699
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v1701 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	v1722 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerEnlarge(m, v1722)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L6
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v501))) = v501
	v1705 = v501
	goto L179
L178:
	;
	v1705 = v1701
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v1705
	*(*int32)(unsafe.Add(mBase, uint32(v1705))) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v501)+4)) = v518
	goto L176
L180:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v518)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+48)) = v1725 + int32(1)
	v1736 = v518
	v1737 = v501
	v1738 = v58
	v1739 = v503
	v1740 = v60
	v1741 = v494
	v1742 = v62
	v1743 = v63
	v1744 = v64
	v1745 = v65
	v1746 = v69
	goto L8
L181:
	;
	m.G0 = v54 + int32(592)
	return v1736
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v50
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheList_6), v54)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L6
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+556)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v54)+552)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v54)+560)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v54)+564)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v54)+568)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+572)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v54)+576)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v54)+580)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v54)+584)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v54)+588)) = v69
	F_errfinish(m, int32(_a_F_SearchCatCacheList_2), int32(385), int32(_a_F_SearchCatCacheList_7))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L6
	} else {
		goto L184
	}
L184:
	;
	goto L5
L185:
	;
	v1863 = int32(v1859)
	m.G0 = v54
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1863)+4))
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1863)))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1866)))
	if v54+int32(44) == v1869 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	m.ExcPending = 1
	goto L194
L187:
	;
	if v1873 != 0 {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1866)+4))
	v1873 = v1871
	goto L190
L189:
	;
	v1873 = int32(0)
	goto L190
L190:
	;
	goto L187
L191:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v54)+588))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v54)+584))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v54)+580))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v54)+576))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v54)+572))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v54)+568))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v54)+564))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v54)+560))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v54)+556))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v54)+552))
	v55 = v1865
	v57 = v1876
	v58 = v1879
	v59 = v1877
	v60 = v1878
	v61 = v1875
	v62 = v1881
	v63 = v1880
	v64 = v1882
	v65 = v1883
	v66 = v1874
	v67 = v1873
	goto L1
L192:
	;
	goto L193
L193:
	;
	F___wasm_longjmp(m, v1866, v1865)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	return int32(0)
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
