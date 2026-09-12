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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v174 int32
	_ = v174
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v221 int32
	_ = v221
	var v243 int32
	_ = v243
	var v266 int32
	_ = v266
	var v292 int32
	_ = v292
	var v316 int32
	_ = v316
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v432 int32
	_ = v432
	var v454 int32
	_ = v454
	var v477 int32
	_ = v477
	var v500 int32
	_ = v500
	var v524 int32
	_ = v524
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v599 int32
	_ = v599
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v647 int32
	_ = v647
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v761 int32
	_ = v761
	var v762 int64
	_ = v762
	var v774 int32
	_ = v774
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1036 int32
	_ = v1036
	var v1057 int32
	_ = v1057
	var v1079 int32
	_ = v1079
	var v1101 int32
	_ = v1101
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1202 int32
	_ = v1202
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1252 int32
	_ = v1252
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1419 int32
	_ = v1419
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1500 int32
	_ = v1500
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1624 int32
	_ = v1624
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1700 int32
	_ = v1700
	var v1724 int32
	_ = v1724
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1755 int32
	_ = v1755
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1787 int32
	_ = v1787
	var v1808 int32
	_ = v1808
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1856 int32
	_ = v1856
	var v1879 int32
	_ = v1879
	var v1903 int32
	_ = v1903
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1958 int32
	_ = v1958
	var v1980 int64
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2017 int32
	_ = v2017
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2052 int32
	_ = v2052
	var v2055 int32
	_ = v2055
	var v2085 int32
	_ = v2085
	var v2089 int32
	_ = v2089
	var v2090 int64
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
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
	var v2125 int32
	_ = v2125
	v5 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(144)
	m.G0 = v36
	v44 = v5
	v45 = v5
	v46 = v5
	v47 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	v51 = v5
	v52 = v5
	v53 = v5
	v54 = v5
	v55 = v5
	v56 = v5
	v57 = v5
	v58 = v5
	v59 = v5
	v60 = v5
	v61 = v5
	v62 = v5
	v63 = int32(-1)
	v64 = v5
	v68 = v36
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
	m.G0 = v36 + int32(144)
	return
L4:
	;
	goto L3
L5:
	;
	if v63 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L4
L7:
	;
	v2089 = int32(m.ExcTag)
	v2090 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2089 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1919
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1913
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1909
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1914
	F_sequence_close(m, v1918, int32(3))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		v2085 = v1933
		goto L7
	} else {
		goto L140
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v1830 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L135
	}
L10:
	;
	if v1307 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L11:
	;
	v1287 = v44
	v1288 = v45
	v1289 = v46
	v1290 = v47
	v1291 = v48
	v1292 = v49
	v1293 = v50
	v1294 = v51
	v1295 = v52
	v1296 = v53
	v1297 = v54
	v1298 = v55
	v1299 = v56
	v1300 = v57
	v1301 = v58
	v1302 = v59
	v1303 = v60
	v1304 = v61
	v1306 = v62
	v1307 = v64
	v1311 = v68
	goto L10
L12:
	;
	goto L13
L13:
	;
	v75 = v68 - int32(32)
	m.G0 = v75
	v78 = v75 - int32(80)
	m.G0 = v78
	v81 = v78 + int32(-64)
	m.G0 = v81
	v84 = v81 - int32(48)
	m.G0 = v84
	v87 = v84 - int32(16)
	m.G0 = v87
	v90 = v87 - int32(160)
	m.G0 = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v112 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v84)+32)) = v113
	v118 = v84 + int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v84)+16)) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v113
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_parse_subscription_options(m, l1, v125, int32(49087), v84)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v149 = v84 + int32(14)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+14)))
	if v150 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_PreventInTransactionBlock(m, l3, int32(676516))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v195 = F_has_privs_of_role(m, v112, int32(6304))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v195 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	v336 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v340 = F_object_aclcheck(m, int32(1262), v336, v112, int64(512))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L28
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errcode(m, int32(16797828))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errmsg(m, int32(247227), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = int32(247114)
	F_errdetail(m, int32(590126), v36+int32(48))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errfinish(m, int32(494261), int32(588), int32(247459))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L1
L28:
	;
	if v340 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	v361 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v363 = F_get_database_name(m, v361)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+21)))
	if v388 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_aclcheck_error(m, v340, int32(9), v363)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v546 = F_table_open(m, int32(6100), int32(3))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L43
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v408 = F_superuser_arg(m, v112)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L36
	}
L36:
	;
	if v408 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errcode(m, int32(16797828))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errmsg(m, int32(19608), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errhint(m, int32(609352), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errfinish(m, int32(494261), int32(609), int32(247459))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L1
L43:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	v560 = l2 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	v568 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v573 = int32(0)
	v575 = F_GetSysCacheOid(m, int32(66), v568, v548, v573, v573)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L44
	}
L44:
	;
	if v575 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v673 = v84 + int32(4)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v674&int32(8) != 0 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errcode(m, int32(290948))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v622
	F_errmsg(m, int32(116435), v36+int32(32))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errfinish(m, int32(494261), int32(630), int32(247459))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L51
	}
L51:
	;
	goto L1
L52:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v680 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v677 != 0 {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(v673))) = v678
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = int32(338926)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_load_file(m, int32(214903), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+21)))
	if v714 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v736 = F_superuser(m)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L62
	}
L60:
	;
	v740 = int32(0)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	m.T0[v713].(func(*base.Module, int32, int32))(m, v685, v740)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L63
	}
L62:
	;
	v740 = v736 ^ int32(1)
	goto L61
L63:
	;
	v762 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v78)+48)) = v762
	*(*int64)(unsafe.Add(mBase, uint32(v78-int32(-64)))) = v762
	*(*int64)(unsafe.Add(mBase, uint32(v78)+56)) = v762
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v762
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v762
	v774 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+16)) = uint16(v774)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v797 = F_GetNewOidWithIndex(m, v546, int32(6114), int32(1))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v797
	v801 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v801
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v823 = F_Int64GetDatum(m, int64(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v823
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v848 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v826)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v848
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+13)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v852
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+17)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v854
	v856 = int32(*(*int8)(unsafe.Add(mBase, uint32(v84)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v856
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+19)))
	if v860 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v861 = int32(112)
	goto L69
L68:
	;
	v861 = int32(100)
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v861
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = v863
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+21)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v865
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v867
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+23)))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = v869
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	v876 = v84 + int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	v879 = v84 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	v882 = v84 + int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v896 = F_cstring_to_text(m, v685)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+52)) = v896
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	if v899 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v946 = F_cstring_to_text(m, v926)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L76
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v921 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v899)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v924 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+14)) = uint8(v924)
	goto L71
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+56)) = v921
	goto L71
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+60)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v968 = F_publicationListToArray(m, v686)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+64)) = v968
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v991 = F_cstring_to_text(m, v971)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+68)) = v991
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v546)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v1014 = F_heap_form_tuple(m, v994, v78, v75)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_CatalogTupleInsert(m, v546, v1014)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_pfree(m, v1014)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_recordDependencyOnOwner(m, int32(6100), v797, v112)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_ReplicationOriginNameForLogicalRep(m, v797, int32(0), v81)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v1121 = F_replorigin_create(m, v81)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L84
	}
L84:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+12)))
	if v1123 != int32(1) {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v1145 = F_superuser_arg(m, v112)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L86
	}
L86:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v1149 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)))
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+21)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	v1171 = int32(1)
	v1176 = m.T0[v1150].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v685, v1171, v1171, v1151&(v1145^v1171), v1147, v87)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L87
	}
L87:
	;
	if v1176 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v1281 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	goto L95
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L92
	}
L92:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1226
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v1225
	F_errmsg(m, int32(201609), v36+int32(16))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errfinish(m, int32(494261), int32(719), int32(247459))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L94
	}
L94:
	;
	goto L1
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v36 - int32(-64)
	goto L98
L96:
	;
	v1287 = v84
	v1288 = v78
	v1289 = v1176
	v1290 = v797
	v1291 = v90
	v1292 = v75
	v1293 = v560
	v1294 = v686
	v1295 = v673
	v1296 = v546
	v1297 = v1281
	v1298 = v1279
	v1299 = v81
	v1300 = v87
	v1301 = v882
	v1302 = v149
	v1303 = v876
	v1304 = v879
	v1306 = v118
	v1307 = int32(0)
	v1311 = v90
	goto L10
L98:
	;
	goto L96
L99:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_check_publications(m, v1289, v1294)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v1298
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v1297
	v1765 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1765)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	m.T0[v1766].(func(*base.Module, int32))(m, v1289)
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L133
	}
L102:
	;
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1293)))
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1359 = int32(0)
	F_check_publications_origin(m, v1289, v1294, v1339, v1338, v1359, v1359, v1337)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1383 = F_fetch_table_list(m, v1289, v1294)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L105
	}
L104:
	;
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302))))
	if v1561 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L105:
	;
	if v1383 == int32(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v1387 = int32(0)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+4))
	if v1388 <= v1387 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if v1363 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v1393 = int32(105)
	goto L110
L109:
	;
	v1393 = int32(114)
	goto L110
L110:
	;
	v1419 = v1387
	goto L111
L111:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+12))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1427+v1419<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1452 = int32(0)
	v1455 = F_RangeVarGetRelidExtended(m, v1431, int32(1), v1452, v1452, v1452)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L113
	}
L112:
	;
	goto L104
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1476 = F_get_rel_relkind(m, v1455)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+12))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_CheckSubscriptionRelkind(m, v1476, v1479, v1478)
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_AddSubscriptionRelState(m, v1290, v1455, v1393, int64(0), int32(1))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v1525 = v1419 + int32(1)
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+4))
	if v1525 < v1526 {
		v1419 = v1525
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L112
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v1298
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v1297
	v1733 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1733)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	m.T0[v1734].(func(*base.Module, int32))(m, v1289)
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L132
	}
L119:
	;
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303))))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1567 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+48))
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304))))
	if v1569 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1673 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L128
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1644 = int32(0)
	v1646 = int32(1)
	v1650 = m.T0[v1568].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1289, v1565, v1644, v1644, v1564&v1646, v1646, v1644)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L127
	}
L122:
	;
	if v1383 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287)+15)))
	if v1574&int32(1) != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	v1596 = int32(0)
	v1597 = int32(1)
	v1602 = m.T0[v1568].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1289, v1565, v1596, v1597, v1564&v1597, v1597, v1596)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_UpdateTwoPhaseState(m, v1290)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L126
	}
L126:
	;
	goto L120
L127:
	;
	goto L120
L128:
	;
	if v1673 == int32(0) {
		goto L118
	} else {
		goto L129
	}
L129:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v1677
	F_errmsg(m, int32(222961), v36)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_errfinish(m, int32(494261), int32(791), int32(247459))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L131
	}
L131:
	;
	goto L118
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v1298
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v1297
	v1909 = v1287
	v1910 = v1288
	v1911 = v1289
	v1912 = v1290
	v1913 = v1291
	v1914 = v1292
	v1915 = v1293
	v1916 = v1294
	v1917 = v1295
	v1918 = v1296
	v1919 = v1297
	v1920 = v1298
	v1921 = v1299
	v1922 = v1300
	v1923 = v1301
	v1924 = v1302
	v1925 = v1303
	v1926 = v1304
	v1928 = v1306
	v1933 = v1311
	goto L8
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1297
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1298
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1289
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1303
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1301
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1290
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1293
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1296
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1302
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1291
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1287
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1299
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1292
	F_pg_re_throw(m)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		v2085 = v1311
		goto L7
	} else {
		goto L134
	}
L134:
	;
	goto L1
L135:
	;
	if v1830 == int32(0) {
		v1909 = v84
		v1910 = v78
		v1911 = v46
		v1912 = v797
		v1913 = v90
		v1914 = v75
		v1915 = v560
		v1916 = v686
		v1917 = v673
		v1918 = v546
		v1919 = v54
		v1920 = v55
		v1921 = v81
		v1922 = v87
		v1923 = v882
		v1924 = v149
		v1925 = v876
		v1926 = v879
		v1928 = v118
		v1933 = v90
		goto L8
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errmsg(m, int32(447307), int32(0))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errhint(m, int32(613392), int32(0))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v75
	F_errfinish(m, int32(494261), int32(803), int32(247459))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		v2085 = v90
		goto L7
	} else {
		goto L139
	}
L139:
	;
	v1909 = v84
	v1910 = v78
	v1911 = v46
	v1912 = v797
	v1913 = v90
	v1914 = v75
	v1915 = v560
	v1916 = v686
	v1917 = v673
	v1918 = v546
	v1919 = v54
	v1920 = v55
	v1921 = v81
	v1922 = v87
	v1923 = v882
	v1924 = v149
	v1925 = v876
	v1926 = v879
	v1928 = v118
	v1933 = v90
	goto L8
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1919
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1913
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1909
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1914
	v1980 = base.I64_extend_i32_u(v1912)
	F_pgstat_create_transactional(m, int32(5), int32(0), v1980)
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		v2085 = v1933
		goto L7
	} else {
		goto L141
	}
L141:
	;
	v1984 = int32(0)
	v1987 = F_pgstat_get_entry_ref(m, int32(5), v1984, v1980, int32(1), v1984)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		v2085 = v1933
		goto L7
	} else {
		goto L142
	}
L142:
	;
	F_pgstat_reset_entry(m, int32(5), int32(0), v1980, int64(0))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		v2085 = v1933
		goto L7
	} else {
		goto L143
	}
L143:
	;
	v1994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1923))))
	if v1994 == int32(1) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1919
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1913
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1909
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1914
	v2017 = int32(*(*uint8)(unsafe.Add(mBase, _consts[494])))
	if v2017 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	v2023 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2023
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(6100)
	v2029 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	if v2029 == v2023 {
		goto L4
	} else {
		goto L151
	}
L147:
	;
	goto L146
L148:
	;
	v2021 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[494])) = uint8(v2021)
	goto L150
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+72)) = v1919
	*(*int32)(unsafe.Add(mBase, uint32(v36)+68)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1916
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v36)+112)) = v1918
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v36)+120)) = v1913
	*(*int32)(unsafe.Add(mBase, uint32(v36)+124)) = v1922
	*(*int32)(unsafe.Add(mBase, uint32(v36)+128)) = v1909
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = v1921
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v36)+140)) = v1914
	v2052 = int32(0)
	F_RunObjectPostCreateHook(m, int32(6100), v1912, v2052, v2052)
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		v2085 = v1933
		goto L7
	} else {
		goto L152
	}
L152:
	;
	goto L6
L153:
	;
	v2094 = int32(v2090)
	m.G0 = v2085
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+4))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2094)))
	v2101 = *(*int32)(unsafe.Add(mBase, uint32(v2097)))
	if v36-int32(-64) == v2101 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	m.ExcPending = 1
	goto L162
L155:
	;
	if v2104 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2097)+4))
	v2104 = v2103
	goto L158
L157:
	;
	v2104 = int32(0)
	goto L158
L158:
	;
	goto L155
L159:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v36)+140))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v36)+136))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v36)+132))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v36)+128))
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v36)+124))
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v36)+120))
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v36)+116))
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+112))
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v36)+108))
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	v44 = v2108
	v45 = v2106
	v46 = v2121
	v47 = v2116
	v48 = v2110
	v49 = v2105
	v50 = v2113
	v51 = v2115
	v52 = v2114
	v53 = v2112
	v54 = v2122
	v55 = v2123
	v56 = v2107
	v57 = v2109
	v58 = v2117
	v59 = v2111
	v60 = v2119
	v61 = v2118
	v62 = v2120
	v63 = v2104
	v64 = v2096
	v68 = v2085
	goto L2
L160:
	;
	goto L161
L161:
	;
	F___wasm_longjmp(m, v2097, v2096)
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	return
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_GetSubscription(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(67), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v153
L2:
	;
	return int32(0)
L3:
	;
	if v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		v153 = int32(0)
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v34 = F_palloc(m, int32(56))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(44593), v9)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(496196), int32(87), int32(247408))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = l0
	v37 = v31 + v32
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v40
	v44 = F_pstrdup(m, v37+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v47
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+25)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+85)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)) = uint8(v51)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+86)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+27)) = uint8(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+28)) = uint8(v55)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+88)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+29)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+89)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+30)) = uint8(v59)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+90)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+31)) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+91)))
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+32)) = uint8(v63)
	v67 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(14))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v69 = F_text_to_cstring(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v69
	v72 = int32(0)
	v77 = F_SysCacheGetAttr(m, int32(67), v12, int32(15), v9+int32(7))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v84 = int32(0)
	goto L18
L17:
	;
	v82 = F_pstrdup(m, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v84
	v88 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L20
	}
L19:
	;
	v84 = v82
	goto L18
L20:
	;
	v90 = F_text_to_cstring(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v90
	v95 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(17))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v97 = F_pg_detoast_datum(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_deconstruct_array_builtin(m, v97, int32(25), v9+int32(12), int32(0), v9+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v107 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v110 = int32(0)
	v111 = v72
	goto L28
L26:
	;
	v132 = v72
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v132
	v140 = F_SysCacheGetAttrNotNull(m, int32(67), v12, int32(18))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L34
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v110<<(uint(int32(2))%32))))
	v121 = F_text_to_cstring(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L30
	}
L29:
	;
	v132 = v125
	goto L27
L30:
	;
	v123 = F_makeString(m, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v125 = F_lappend(m, v111, v123)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v128 = v110 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v128 < v129 {
		v110 = v128
		v111 = v125
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v142 = F_text_to_cstring(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v146 = F_superuser_arg(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)) = uint8(v146)
	F_ReleaseCatCache(m, v12)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v153 = v34
	goto L1
}
func F_GetSubscriptionRelState(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(6102), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = F_SearchSysCache2(m, int32(68), l1, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				F_sequence_close(m, v13, int32(1))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
					v48 = int32(0)
					m.G0 = v9 + int32(16)
					return base.I32_extend8_s(v48)
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+8)))
				v36 = F_SysCacheGetAttr(m, int32(68), v18, int32(4), v9+int32(15))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v38 != 0 {
						v41 = int64(0)
					} else {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
						v41 = v40
					}
					*(*int64)(unsafe.Add(mBase, uint32(l2))) = v41
					F_ReleaseCatCache(m, v18)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v13, int32(1))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v31
							m.G0 = v9 + int32(16)
							return base.I32_extend8_s(v48)
						}
					}
				}
			}
		}
	}
}
