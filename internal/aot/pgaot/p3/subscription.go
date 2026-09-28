package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateSubscription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v116 int32
	_ = v116
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v192 int32
	_ = v192
	var v209 int32
	_ = v209
	var v227 int32
	_ = v227
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v267 int32
	_ = v267
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v411 int32
	_ = v411
	var v430 int32
	_ = v430
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v462 int32
	_ = v462
	var v468 int64
	_ = v468
	var v469 int64
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v490 int32
	_ = v490
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v528 int32
	_ = v528
	var v547 int32
	_ = v547
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v783 int32
	_ = v783
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v821 int32
	_ = v821
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int64
	_ = v908
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int64
	_ = v949
	var v953 int64
	_ = v953
	var v970 int64
	_ = v970
	var v971 int32
	_ = v971
	var v973 int64
	_ = v973
	var v975 int64
	_ = v975
	var v977 int64
	_ = v977
	var v979 int64
	_ = v979
	var v981 int64
	_ = v981
	var v983 int64
	_ = v983
	var v985 int64
	_ = v985
	var v991 int32
	_ = v991
	var v992 int64
	_ = v992
	var v994 int64
	_ = v994
	var v999 int64
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1042 int64
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1099 int64
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1157 int32
	_ = v1157
	var v1173 int32
	_ = v1173
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1220 int32
	_ = v1220
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1322 int32
	_ = v1322
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1494 int32
	_ = v1494
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1617 int32
	_ = v1617
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1667 int32
	_ = v1667
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1727 int32
	_ = v1727
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1766 int32
	_ = v1766
	var v1785 int32
	_ = v1785
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1813 int32
	_ = v1813
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1840 int32
	_ = v1840
	var v1856 int32
	_ = v1856
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1894 int32
	_ = v1894
	var v1912 int32
	_ = v1912
	var v1931 int32
	_ = v1931
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1982 int32
	_ = v1982
	var v1999 int64
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2099 int32
	_ = v2099
	var v2100 int64
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
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
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	v5 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(656)
	m.G0 = v37
	v40 = v37 + int32(278)
	v42 = v37 + int32(288)
	v44 = v37 + int32(277)
	v46 = v37 + int32(287)
	v48 = v37 + int32(296)
	v55 = v5
	v56 = v5
	v57 = v5
	v58 = v5
	v59 = v5
	v60 = v5
	v61 = v5
	v62 = v5
	v63 = v5
	v64 = v5
	v65 = v5
	v67 = v5
	v68 = v5
	v69 = v5
	v72 = v5
	v74 = int32(-1)
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v37 + int32(656)
	return
L4:
	;
	goto L3
L5:
	;
	if v74 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L4
L7:
	;
	v2099 = int32(m.ExcTag)
	v2100 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2099 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1943
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1937
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1940
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1950
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1951
	F_relation_close(m, v1941, int32(3))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L7
	} else {
		goto L184
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1873 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L7
	} else {
		goto L177
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v69
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[0]))
	v102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+312)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+304)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+296)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+288)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+280)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+272)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v37)+264)) = v102
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v69
	F_parse_subscription_options(m, l1, v116, int32(_a_F_CreateSubscription_0), v37+int32(264))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v1394 = v55
	v1395 = v56
	v1396 = v57
	v1397 = v58
	v1398 = v59
	v1399 = v60
	v1400 = v61
	v1401 = v62
	v1402 = v63
	v1403 = v64
	v1404 = v65
	v1405 = v67
	v1406 = v68
	v1407 = v69
	v1408 = v72
	goto L12
L12:
	;
	if v1408 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L13:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+278)))
	if v136 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_PreventInTransactionBlock(m, l3, int32(_a_F_CreateSubscription_1))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v171 = F_has_privs_of_role(m, v101, int32(_a_F_CreateSubscription_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	if v171 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[1]))
	v286 = F_object_aclcheck(m, int32(1262), v284, v101, int64(512))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errcode(m, int32(16797828))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errmsg(m, int32(_a_F_CreateSubscription_3), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = int32(_a_F_CreateSubscription_4)
	v247 = F_errdetail(m, int32(_a_F_CreateSubscription_5), v37-int32(-64))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(700), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L1
L27:
	;
	if v286 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[1]))
	v304 = F_get_database_name(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+285)))
	if v324 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_aclcheck_error(m, v286, int32(9), v304)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v447 = F_table_open(m, int32(_a_F_CreateSubscription_8), int32(3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L42
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v339 = F_superuser_arg(m, v101)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	if v339 != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errcode(m, int32(16797828))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errmsg(m, int32(_a_F_CreateSubscription_9), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errhint(m, int32(_a_F_CreateSubscription_10), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(721), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	v449 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	v462 = l2 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v468 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_CreateSubscription[1])))
	v469 = int64(0)
	v471 = F_GetSysCacheOid(m, int32(66), v468, v449, v469, v469)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	if v471 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L7
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	v562 = int32(1)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+277)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+288)))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v37)+292))
	F_CheckSubDeadTupleRetention(m, v562, (v563^int32(-1))&v562, int32(19), v569, v569, base.B2i32(int32(0) < v570))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L7
	} else {
		goto L51
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errcode(m, int32(_a_F_CreateSubscription_11))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v508
	F_errmsg(m, int32(_a_F_CreateSubscription_12), v37+int32(48))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(742), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	goto L1
L51:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+264)))
	if v575&int32(8) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
	if v581 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v37)+268))
	if v578 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+268)) = v579
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+272)) = int32(_a_F_CreateSubscription_13)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v37)+312))
	if v586 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+312)) = int32(_a_F_CreateSubscription_14)
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	v602 = v37 + int32(264) | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_load_file(m, int32(_a_F_CreateSubscription_15), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v613 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v908 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+560)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+552)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+544)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+536)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+528)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+520)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+512)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+576)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+584)) = v908
	*(*int64)(unsafe.Add(mBase, uint32(v37)+591)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v944 = F_GetNewOidWithIndex(m, v447, int32(_a_F_CreateSubscription_16), int32(1))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L7
	} else {
		goto L94
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v629 = F_GetForeignServerByName(m, v613, int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v860 = int32(0)
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+285)))
	if v865 == int32(1) {
		goto L89
	} else {
		goto L90
	}
L66:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v648 = F_object_aclcheck(m, int32(1417), v631, v101, int64(256))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	if v648 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v629)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_aclcheck_error(m, v648, int32(17), v650)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L7
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_GetUserMappingExtended(m, v101, v669)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L7
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+276)))
	if v686 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v703 = F_ForeignServerConnectionString(m, v101, v629)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L7
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v762 = F_GetForeignDataWrapper(m, v747)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L7
	} else {
		goto L82
	}
L76:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+285)))
	if v709 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v726 = F_superuser(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	v730 = int32(0)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	m.T0[v708].(func(*base.Module, int32, int32))(m, v703, v730)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L7
	} else {
		goto L81
	}
L80:
	;
	v730 = v726 ^ int32(1)
	goto L79
L81:
	;
	v903 = v631
	v904 = v703
	goto L62
L82:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v762)+20))
	if v765 != 0 {
		v903 = v631
		v904 = int32(0)
		goto L62
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errcode(m, int32(1088))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v762)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v801
	F_errmsg(m, int32(_a_F_CreateSubscription_17), v37+int32(32))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L7
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v838 = F_errdetail(m, int32(_a_F_CreateSubscription_18), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(807), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	goto L1
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v882 = F_superuser(m)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	v886 = v860
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	m.T0[v863].(func(*base.Module, int32, int32))(m, v859, v886)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L7
	} else {
		goto L93
	}
L92:
	;
	v886 = v882 ^ int32(1)
	goto L91
L93:
	;
	v903 = v860
	v904 = v859
	goto L62
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+400)) = int64(0)
	v949 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_CreateSubscription[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+392)) = v949
	*(*int64)(unsafe.Add(mBase, uint32(v37)+384)) = base.I64_extend_i32_u(v944)
	v953 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v462))))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	v970 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), v953)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+408)) = v970
	v973 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+277)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+424)) = v973
	v975 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+281)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+432)) = v975
	v977 = int64(*(*int8)(unsafe.Add(mBase, uint32(v37)+282)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+440)) = v977
	v979 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+284)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+456)) = v979
	v981 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+285)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+464)) = v981
	v983 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+286)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+472)) = v983
	v985 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+287)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+480)) = v985
	*(*int64)(unsafe.Add(mBase, uint32(v37)+416)) = base.I64_extend_i32_u(v101)
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+283)))
	if v991 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v992 = int64(112)
	goto L98
L97:
	;
	v992 = int64(100)
	goto L98
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+448)) = v992
	v994 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+288)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+488)) = v994
	*(*int64)(unsafe.Add(mBase, uint32(v37)+504)) = v994
	*(*int64)(unsafe.Add(mBase, uint32(v37)+512)) = base.I64_extend_i32_u(v903)
	v999 = int64(*(*int32)(unsafe.Add(mBase, uint32(v37)+292)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+496)) = v999
	v1002 = v37 + int32(283)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v1003 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v37)+268))
	if v1024 != 0 {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1018 = F_cstring_to_text(m, v1003)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v1022 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+593)) = uint8(v1022)
	goto L99
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+520)) = base.I64_extend_i32_u(v1018)
	goto L99
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
	v1062 = F_cstring_to_text(m, v1061)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L7
	} else {
		goto L109
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1042 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), base.I64_extend_i32_u(v1024))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L7
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v1045 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+594)) = uint8(v1045)
	goto L104
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+528)) = v1042
	goto L104
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v37)+536)) = base.I64_extend_i32_u(v1062)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v37)+312))
	v1081 = F_cstring_to_text(m, v1080)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v37)+544)) = base.I64_extend_i32_u(v1081)
	v1099 = F_publicationListToArray(m, v907)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+552)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v37)+296))
	v1117 = F_cstring_to_text(m, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+560)) = base.I64_extend_i32_u(v1117)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v447)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1140 = F_heap_form_tuple(m, v1121, v37+int32(384), v37+int32(576))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_CatalogTupleInsert(m, v447, v1140)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_pfree(m, v1140)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_recordDependencyOnOwner(m, int32(_a_F_CreateSubscription_8), v944, v101)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(_a_F_CreateSubscription_8)
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1196 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+260)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+256)) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v37)+252)) = int32(1417)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_recordDependencyOn(m, l0, v37+int32(252), int32(110))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L7
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1237 = v37 + int32(320)
	F_ReplicationOriginNameForLogicalRep(m, v944, int32(0), v1237)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L7
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1254 = F_replorigin_create(m, v1237)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+276)))
	if v1256 != int32(1) {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1273 = F_superuser_arg(m, v101)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v1277 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1293 = int32(1)
	v1295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+285)))
	v1301 = m.T0[v1278].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v904, v1293, v1293, v1295&(v1273^v1293), v1275, v37+int32(248))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	if v1301 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L7
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[3]))
	v1386 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[4]))
	goto L133
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v37)+248))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = v1355
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v1340
	F_errmsg(m, int32(_a_F_CreateSubscription_19), v37+int32(16))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(926), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	goto L1
L133:
	;
	v1388 = v37 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1388))) = v37 + int32(76)
	goto L136
L134:
	;
	v1394 = v1301
	v1395 = v944
	v1396 = v462
	v1397 = v907
	v1398 = v447
	v1399 = v1384
	v1400 = v1386
	v1401 = v48
	v1402 = v1002
	v1403 = v602
	v1404 = v46
	v1405 = v42
	v1406 = v44
	v1407 = v40
	v1408 = int32(0)
	goto L12
L136:
	;
	goto L134
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[4])) = v37 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_check_publications(m, v1394, v1397)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L7
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[3])) = v1399
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[4])) = v1400
	v1823 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	m.T0[v1824].(func(*base.Module, int32))(m, v1394)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L7
	} else {
		goto L175
	}
L140:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1396)))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1405))))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+279)))
	v1451 = int32(0)
	F_check_publications_origin_tables(m, v1394, v1397, v1450, v1435, v1434, v1451, v1451, v1433)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1396)))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+279)))
	v1472 = int32(0)
	F_check_publications_origin_sequences(m, v1394, v1397, v1471, v1456, v1472, v1472, v1455)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1405))))
	if v1476 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_CheckPubDeadTupleRetention(m, v1394)
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L7
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+279)))
	v1510 = F_fetch_relation_list(m, v1394, v1397)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L7
	} else {
		goto L148
	}
L146:
	;
	goto L145
L147:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1407))))
	if v1677 != int32(1) {
		goto L163
	} else {
		goto L164
	}
L148:
	;
	if v1510 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v1667 = int32(0)
	goto L147
L150:
	;
	goto L151
L151:
	;
	v1515 = int32(0)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+4))
	if v1516 <= v1515 {
		v1667 = v1515
		goto L147
	} else {
		goto L152
	}
L152:
	;
	if v1509 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v1521 = int32(105)
	goto L155
L154:
	;
	v1521 = int32(114)
	goto L155
L155:
	;
	v1545 = int32(0)
	v1547 = v1515
	goto L156
L156:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+12))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1557+v1545<<(uint(int32(2))%32))))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1578 = int32(0)
	v1581 = F_RangeVarGetRelidExtended(m, v1562, int32(1), v1578, v1578, v1578)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L7
	} else {
		goto L158
	}
L157:
	;
	v1667 = v1638
	goto L147
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1597 = F_get_rel_relkind(m, v1581)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L7
	} else {
		goto L159
	}
L159:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+12))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+8))
	v1601 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1561)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_CheckSubscriptionRelkind(m, v1597, v1601, v1600, v1599)
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L7
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_AddSubscriptionRelState(m, v1395, v1581, v1521, int64(0), int32(1))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	v1638 = v1547 | base.B2i32(v1597 != int32(83))
	v1640 = v1545 + int32(1)
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1510)+4))
	if v1640 < v1641 {
		v1545 = v1640
		v1547 = v1638
		goto L156
	} else {
		goto L162
	}
L162:
	;
	goto L157
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[3])) = v1399
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[4])) = v1400
	v1796 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	m.T0[v1797].(func(*base.Module, int32))(m, v1394)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L7
	} else {
		goto L174
	}
L164:
	;
	v1680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404))))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1403)))
	v1683 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[2]))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1683)+48))
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402))))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1700 = int32(0)
	v1701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+279)))
	v1702 = int32(1)
	v1707 = v1685 & (v1701 ^ v1702) & v1667 & v1702
	v1710 = m.T0[v1684].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1394, v1681, v1700, v1707, v1680, v1702, v1700)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L7
	} else {
		goto L165
	}
L165:
	;
	if v1707 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_UpdateTwoPhaseState(m, v1395)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L7
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	v1744 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L7
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	if v1744 == int32(0) {
		goto L163
	} else {
		goto L171
	}
L171:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1403)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v1748
	F_errmsg(m, int32(_a_F_CreateSubscription_20), v37)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L7
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(1017), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	goto L163
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[3])) = v1399
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[4])) = v1400
	v1937 = v1394
	v1938 = v1395
	v1939 = v1396
	v1940 = v1397
	v1941 = v1398
	v1942 = v1399
	v1943 = v1400
	v1944 = v1401
	v1945 = v1402
	v1946 = v1403
	v1947 = v1404
	v1949 = v1405
	v1950 = v1406
	v1951 = v1407
	goto L8
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1400
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1401
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1404
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1403
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1396
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1407
	F_pg_re_throw(m)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L7
	} else {
		goto L176
	}
L176:
	;
	goto L1
L177:
	;
	if v1873 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1937 = v55
	v1938 = v944
	v1939 = v462
	v1940 = v907
	v1941 = v447
	v1942 = v60
	v1943 = v61
	v1944 = v48
	v1945 = v1002
	v1946 = v602
	v1947 = v46
	v1949 = v42
	v1950 = v44
	v1951 = v40
	goto L8
L179:
	;
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errmsg(m, int32(_a_F_CreateSubscription_21), int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errhint(m, int32(_a_F_CreateSubscription_22), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L7
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v907
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v462
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v40
	F_errfinish(m, int32(_a_F_CreateSubscription_6), int32(1029), int32(_a_F_CreateSubscription_7))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L7
	} else {
		goto L183
	}
L183:
	;
	v1937 = v55
	v1938 = v944
	v1939 = v462
	v1940 = v907
	v1941 = v447
	v1942 = v60
	v1943 = v61
	v1944 = v48
	v1945 = v1002
	v1946 = v602
	v1947 = v46
	v1949 = v42
	v1950 = v44
	v1951 = v40
	goto L8
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1943
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1937
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1940
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1950
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1951
	v1999 = base.I64_extend_i32_u(v1938)
	F_pgstat_create_transactional(m, int32(5), int32(0), v1999)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	v2003 = int32(0)
	v2006 = F_pgstat_get_entry_ref(m, int32(5), v2003, v1999, int32(1), v2003)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	F_pgstat_reset_entry(m, int32(5), int32(0), v1999, int64(0))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1950))))
	if v2013 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSubscription[5]))
	if v2043 == int32(0) {
		goto L4
	} else {
		goto L197
	}
L189:
	;
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1949))))
	if v2016&int32(1) == int32(0) {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1943
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1937
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1940
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1950
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1951
	v2036 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateSubscription[6])))
	if v2036 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	goto L191
L193:
	;
	goto L188
L194:
	;
	v2040 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CreateSubscription[6])) = uint8(v2040)
	goto L196
L195:
	;
	goto L196
L196:
	;
	goto L193
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+604)) = v1943
	*(*int32)(unsafe.Add(mBase, uint32(v37)+600)) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v37)+608)) = v1937
	*(*int32)(unsafe.Add(mBase, uint32(v37)+612)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v37)+616)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v37)+620)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(v37)+624)) = v1938
	*(*int32)(unsafe.Add(mBase, uint32(v37)+628)) = v1940
	*(*int32)(unsafe.Add(mBase, uint32(v37)+632)) = v1946
	*(*int32)(unsafe.Add(mBase, uint32(v37)+636)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(v37)+640)) = v1950
	*(*int32)(unsafe.Add(mBase, uint32(v37)+644)) = v1939
	*(*int32)(unsafe.Add(mBase, uint32(v37)+648)) = v1941
	*(*int32)(unsafe.Add(mBase, uint32(v37)+652)) = v1951
	v2061 = int32(0)
	F_RunObjectPostCreateHook(m, int32(_a_F_CreateSubscription_8), v1938, v2061, v2061)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L7
	} else {
		goto L198
	}
L198:
	;
	goto L6
L199:
	;
	v2104 = int32(v2100)
	m.G0 = v37
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+4))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2104)))
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v2107)))
	if v37+int32(76) == v2110 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	m.ExcPending = 1
	goto L208
L201:
	;
	if v2114 != 0 {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+4))
	v2114 = v2112
	goto L204
L203:
	;
	v2114 = int32(0)
	goto L204
L204:
	;
	goto L201
L205:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v37)+652))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v37)+648))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v37)+644))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v37)+640))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v37)+636))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v37)+632))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v37)+628))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v37)+624))
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v37)+620))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v37)+616))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v37)+612))
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v37)+608))
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v37)+604))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v37)+600))
	v55 = v2126
	v56 = v2122
	v57 = v2117
	v58 = v2121
	v59 = v2116
	v60 = v2128
	v61 = v2127
	v62 = v2125
	v63 = v2123
	v64 = v2120
	v65 = v2124
	v67 = v2119
	v68 = v2118
	v69 = v2115
	v72 = v2106
	v74 = v2114
	goto L2
L206:
	;
	goto L207
L207:
	;
	F___wasm_longjmp(m, v2107, v2106)
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	return
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSubscription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = F_SearchSysCache1(m, int32(67), base.I64_extend_i32_u(l0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v192
L2:
	;
	return int32(0)
L3:
	;
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		v192 = int32(0)
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v36 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_GetSubscription[0]))
	v43 = F_AllocSetContextCreateInternal(m, v38, int32(_a_F_GetSubscription_0), v36, int32(1024), int32(_a_F_GetSubscription_1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg_internal(m, int32(_a_F_GetSubscription_2), v13)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_GetSubscription_3), int32(102), int32(_a_F_GetSubscription_4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	v45 = int32(_a_F_GetSubscription_5)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_GetSubscription[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_GetSubscription[0])) = v43
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v52 = F_palloc0(m, int32(72))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v43
	v56 = v49 + v50
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52)+16)) = v59
	v63 = F_pstrdup(m, v56+int32(16))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+33)) = uint8(v68)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+34)) = uint8(v70)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+86)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+35)) = uint8(v72)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+36)) = uint8(v74)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+37)) = uint8(v76)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+38)) = uint8(v78)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+39)) = uint8(v80)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+40)) = uint8(v82)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+41)) = uint8(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+44)) = v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+48)) = uint8(v88)
	v94 = F_SysCacheGetAttr(m, int32(67), v17, int32(19), v13+int32(7))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)))
	if v97 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v102 = int32(0)
	goto L17
L16:
	;
	v100 = F_pstrdup(m, base.I32_wrap_i64(v94))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+52)) = v102
	v106 = F_SysCacheGetAttrNotNull(m, int32(67), v17, int32(20))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v102 = v100
	goto L17
L19:
	;
	v109 = F_text_to_cstring(m, base.I32_wrap_i64(v106))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+56)) = v109
	v114 = F_SysCacheGetAttrNotNull(m, int32(67), v17, int32(21))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v117 = F_text_to_cstring(m, base.I32_wrap_i64(v114))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+60)) = v117
	v122 = F_SysCacheGetAttrNotNull(m, int32(67), v17, int32(22))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v125 = F_pg_detoast_datum(m, base.I32_wrap_i64(v122))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_deconstruct_array_builtin(m, v125, int32(25), v13+int32(12), int32(0), v13+int32(8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if int32(0) < v135 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v138 = int32(0)
	v139 = v36
	goto L29
L27:
	;
	v164 = v36
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+64)) = v164
	v176 = F_SysCacheGetAttrNotNull(m, int32(67), v17, int32(23))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L35
	}
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v138<<(uint(int32(3))%32))))
	v153 = F_text_to_cstring(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v164 = v157
	goto L28
L31:
	;
	v155 = F_makeString(m, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v157 = F_lappend(m, v139, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v160 = v138 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v160 < v161 {
		v138 = v160
		v139 = v157
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v179 = F_text_to_cstring(m, base.I32_wrap_i64(v176))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+68)) = v179
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v183 = F_superuser_arg(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+32)) = uint8(v183)
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_GetSubscription[0])) = v46
	v192 = v52
	goto L1
}
func F_GetSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(_a_F_GetSubscriptionRelState_0), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v19 = F_SearchSysCache2(m, int32(68), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 == int32(0) {
				F_relation_close(m, v12, int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v48 = int32(0)
					m.G0 = v8 + int32(16)
					return base.I32_extend8_s(v48)
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+8)))
				v38 = F_SysCacheGetAttr(m, int32(68), v19, int32(4), v8+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v40 != 0 {
						v41 = int64(0)
					} else {
						v41 = v38
					}
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v41
					F_ReleaseCatCache(m, v19)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v12, int32(1))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v32
							m.G0 = v8 + int32(16)
							return base.I32_extend8_s(v48)
						}
					}
				}
			}
		}
	}
}
