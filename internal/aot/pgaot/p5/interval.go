package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_interval_cmp_lower_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_range_cmp_bounds(m, l2, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_interval_cmp_upper_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v4 float64
	_ = v4
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v3) & v9
	v13 = base.I64_reinterpret_f64(v4) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
		v22 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
	} else {
		v18 = int32(1)
		if base.F64_gt(v3, v4) != 0 {
			v30 = v18
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v30 = v18
			} else {
				v22 = v18
				v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
			}
		}
	}
	return v30
}
func F_interval_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v9 int32
	_ = v9
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v173 int32
	_ = v173
	var v191 int32
	_ = v191
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v311 int64
	_ = v311
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v325 int64
	_ = v325
	var v332 int64
	_ = v332
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v349 int64
	_ = v349
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v416 int64
	_ = v416
	var v419 int64
	_ = v419
	var v422 int64
	_ = v422
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v430 int64
	_ = v430
	var v437 int64
	_ = v437
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v453 int64
	_ = v453
	var v459 int64
	_ = v459
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v468 int64
	_ = v468
	var v470 int64
	_ = v470
	var v471 int64
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v485 int64
	_ = v485
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v716 int32
	_ = v716
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v922 int32
	_ = v922
	var v952 int32
	_ = v952
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int64
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 float64
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int64
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1067 int64
	_ = v1067
	var v1070 int64
	_ = v1070
	var v1073 int64
	_ = v1073
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1081 int64
	_ = v1081
	var v1088 int64
	_ = v1088
	var v1099 int64
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int64
	_ = v1113
	var v1116 int64
	_ = v1116
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int64
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1268 int32
	_ = v1268
	var v1275 float64
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1285 float64
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 float64
	_ = v1291
	var v1292 int64
	_ = v1292
	var v1294 float64
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1301 int64
	_ = v1301
	var v1303 int64
	_ = v1303
	var v1309 int64
	_ = v1309
	var v1311 float64
	_ = v1311
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1328 int64
	_ = v1328
	var v1330 int64
	_ = v1330
	var v1332 float64
	_ = v1332
	var v1343 int64
	_ = v1343
	var v1344 int64
	_ = v1344
	var v1360 int32
	_ = v1360
	var v1363 int64
	_ = v1363
	var v1364 int64
	_ = v1364
	var v1369 int64
	_ = v1369
	var v1372 int64
	_ = v1372
	var v1375 int64
	_ = v1375
	var v1378 int64
	_ = v1378
	var v1379 int64
	_ = v1379
	var v1383 int64
	_ = v1383
	var v1390 int64
	_ = v1390
	var v1401 int64
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1407 int64
	_ = v1407
	var v1408 int64
	_ = v1408
	var v1418 float64
	_ = v1418
	var v1422 int64
	_ = v1422
	var v1424 int64
	_ = v1424
	var v1426 float64
	_ = v1426
	var v1437 int64
	_ = v1437
	var v1438 int64
	_ = v1438
	var v1454 int32
	_ = v1454
	var v1457 int64
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1463 int64
	_ = v1463
	var v1466 int64
	_ = v1466
	var v1469 int64
	_ = v1469
	var v1472 int64
	_ = v1472
	var v1473 int64
	_ = v1473
	var v1477 int64
	_ = v1477
	var v1484 int64
	_ = v1484
	var v1495 int64
	_ = v1495
	var v1496 int64
	_ = v1496
	var v1501 int64
	_ = v1501
	var v1502 int64
	_ = v1502
	var v1512 float64
	_ = v1512
	var v1516 int64
	_ = v1516
	var v1518 int64
	_ = v1518
	var v1520 float64
	_ = v1520
	var v1531 int64
	_ = v1531
	var v1532 int64
	_ = v1532
	var v1550 int32
	_ = v1550
	var v1553 int64
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1559 int64
	_ = v1559
	var v1562 int64
	_ = v1562
	var v1565 int64
	_ = v1565
	var v1568 int64
	_ = v1568
	var v1569 int64
	_ = v1569
	var v1573 int64
	_ = v1573
	var v1580 int64
	_ = v1580
	var v1591 int64
	_ = v1591
	var v1592 int64
	_ = v1592
	var v1597 int64
	_ = v1597
	var v1598 int64
	_ = v1598
	var v1608 float64
	_ = v1608
	var v1612 int64
	_ = v1612
	var v1614 int64
	_ = v1614
	var v1616 float64
	_ = v1616
	var v1627 int64
	_ = v1627
	var v1628 int64
	_ = v1628
	var v1644 int32
	_ = v1644
	var v1647 int64
	_ = v1647
	var v1648 int64
	_ = v1648
	var v1653 int64
	_ = v1653
	var v1656 int64
	_ = v1656
	var v1659 int64
	_ = v1659
	var v1662 int64
	_ = v1662
	var v1663 int64
	_ = v1663
	var v1667 int64
	_ = v1667
	var v1674 int64
	_ = v1674
	var v1685 int64
	_ = v1685
	var v1686 int64
	_ = v1686
	var v1691 int64
	_ = v1691
	var v1692 int64
	_ = v1692
	var v1702 float64
	_ = v1702
	var v1706 int64
	_ = v1706
	var v1708 int64
	_ = v1708
	var v1710 float64
	_ = v1710
	var v1721 int64
	_ = v1721
	var v1722 int64
	_ = v1722
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1752 float64
	_ = v1752
	var v1756 int64
	_ = v1756
	var v1758 int64
	_ = v1758
	var v1760 float64
	_ = v1760
	var v1771 int64
	_ = v1771
	var v1772 int64
	_ = v1772
	var v1773 int64
	_ = v1773
	var v1793 int64
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1813 float64
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1828 float64
	_ = v1828
	var v1832 float64
	_ = v1832
	var v1836 int64
	_ = v1836
	var v1838 int64
	_ = v1838
	var v1840 float64
	_ = v1840
	var v1851 int64
	_ = v1851
	var v1852 int64
	_ = v1852
	var v1853 int64
	_ = v1853
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1887 float64
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1903 float64
	_ = v1903
	var v1907 float64
	_ = v1907
	var v1911 int64
	_ = v1911
	var v1913 int64
	_ = v1913
	var v1915 float64
	_ = v1915
	var v1926 int64
	_ = v1926
	var v1927 int64
	_ = v1927
	var v1928 int64
	_ = v1928
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1962 float64
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1987 int64
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2008 float64
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2033 int64
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2054 float64
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2079 int64
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2100 float64
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2155 int32
	_ = v2155
	var v2165 int32
	_ = v2165
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2218 int64
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2230 int32
	_ = v2230
	var v2251 int32
	_ = v2251
	var v2296 int32
	_ = v2296
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2332 int32
	_ = v2332
	var v2335 int64
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2373 int32
	_ = v2373
	var v2376 int32
	_ = v2376
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2415 int32
	_ = v2415
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2471 float64
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2477 float64
	_ = v2477
	var v2486 int32
	_ = v2486
	var v2494 float64
	_ = v2494
	var v2498 int64
	_ = v2498
	var v2500 int64
	_ = v2500
	var v2503 float64
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2521 int64
	_ = v2521
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2543 int32
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2577 int32
	_ = v2577
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2597 int32
	_ = v2597
	var v2604 int64
	_ = v2604
	var v2615 int32
	_ = v2615
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2626 float64
	_ = v2626
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2639 int32
	_ = v2639
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2659 int32
	_ = v2659
	var v2664 float64
	_ = v2664
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2680 float64
	_ = v2680
	var v2685 float64
	_ = v2685
	var v2689 int64
	_ = v2689
	var v2691 int64
	_ = v2691
	var v2693 float64
	_ = v2693
	var v2704 int64
	_ = v2704
	var v2705 int64
	_ = v2705
	var v2706 int64
	_ = v2706
	var v2722 int64
	_ = v2722
	var v2726 int32
	_ = v2726
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2739 int32
	_ = v2739
	var v2744 float64
	_ = v2744
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2759 float64
	_ = v2759
	var v2764 float64
	_ = v2764
	var v2768 int64
	_ = v2768
	var v2770 int64
	_ = v2770
	var v2772 float64
	_ = v2772
	var v2783 int64
	_ = v2783
	var v2784 int64
	_ = v2784
	var v2785 int64
	_ = v2785
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2809 int32
	_ = v2809
	var v2814 float64
	_ = v2814
	var v2818 int64
	_ = v2818
	var v2820 int64
	_ = v2820
	var v2822 float64
	_ = v2822
	var v2833 int64
	_ = v2833
	var v2834 int64
	_ = v2834
	var v2835 int64
	_ = v2835
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2854 int32
	_ = v2854
	var v2855 int64
	_ = v2855
	var v2856 int64
	_ = v2856
	var v2858 int64
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2867 int32
	_ = v2867
	var v2869 int64
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2881 float64
	_ = v2881
	var v2885 int64
	_ = v2885
	var v2887 int64
	_ = v2887
	var v2889 float64
	_ = v2889
	var v2900 int64
	_ = v2900
	var v2901 int64
	_ = v2901
	var v2902 int64
	_ = v2902
	var v2914 int32
	_ = v2914
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2936 float64
	_ = v2936
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2951 int32
	_ = v2951
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2965 int64
	_ = v2965
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2981 float64
	_ = v2981
	var v2985 float64
	_ = v2985
	var v2989 int32
	_ = v2989
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v3002 float64
	_ = v3002
	var v3006 float64
	_ = v3006
	var v3010 int64
	_ = v3010
	var v3012 int64
	_ = v3012
	var v3014 float64
	_ = v3014
	var v3025 int64
	_ = v3025
	var v3026 int64
	_ = v3026
	var v3027 int64
	_ = v3027
	var v3042 int32
	_ = v3042
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3056 int32
	_ = v3056
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3066 int64
	_ = v3066
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3081 float64
	_ = v3081
	var v3085 float64
	_ = v3085
	var v3089 int64
	_ = v3089
	var v3091 int64
	_ = v3091
	var v3093 float64
	_ = v3093
	var v3104 int64
	_ = v3104
	var v3105 int64
	_ = v3105
	var v3106 int64
	_ = v3106
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3128 int32
	_ = v3128
	var v3131 int64
	_ = v3131
	var v3132 int64
	_ = v3132
	var v3137 int64
	_ = v3137
	var v3140 int64
	_ = v3140
	var v3143 int64
	_ = v3143
	var v3146 int64
	_ = v3146
	var v3147 int64
	_ = v3147
	var v3151 int64
	_ = v3151
	var v3158 int64
	_ = v3158
	var v3169 int64
	_ = v3169
	var v3170 int64
	_ = v3170
	var v3175 int64
	_ = v3175
	var v3176 int64
	_ = v3176
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3188 float64
	_ = v3188
	var v3192 int64
	_ = v3192
	var v3194 int64
	_ = v3194
	var v3196 float64
	_ = v3196
	var v3207 int64
	_ = v3207
	var v3208 int64
	_ = v3208
	var v3218 int32
	_ = v3218
	var v3221 int64
	_ = v3221
	var v3222 int64
	_ = v3222
	var v3227 int64
	_ = v3227
	var v3230 int64
	_ = v3230
	var v3233 int64
	_ = v3233
	var v3236 int64
	_ = v3236
	var v3237 int64
	_ = v3237
	var v3241 int64
	_ = v3241
	var v3248 int64
	_ = v3248
	var v3259 int64
	_ = v3259
	var v3260 int64
	_ = v3260
	var v3265 int64
	_ = v3265
	var v3266 int64
	_ = v3266
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3278 float64
	_ = v3278
	var v3282 int64
	_ = v3282
	var v3284 int64
	_ = v3284
	var v3286 float64
	_ = v3286
	var v3297 int64
	_ = v3297
	var v3298 int64
	_ = v3298
	var v3308 int32
	_ = v3308
	var v3311 int64
	_ = v3311
	var v3312 int64
	_ = v3312
	var v3317 int64
	_ = v3317
	var v3320 int64
	_ = v3320
	var v3323 int64
	_ = v3323
	var v3326 int64
	_ = v3326
	var v3327 int64
	_ = v3327
	var v3331 int64
	_ = v3331
	var v3338 int64
	_ = v3338
	var v3349 int64
	_ = v3349
	var v3350 int64
	_ = v3350
	var v3355 int64
	_ = v3355
	var v3356 int64
	_ = v3356
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3368 float64
	_ = v3368
	var v3372 int64
	_ = v3372
	var v3374 int64
	_ = v3374
	var v3376 float64
	_ = v3376
	var v3387 int64
	_ = v3387
	var v3388 int64
	_ = v3388
	var v3403 int32
	_ = v3403
	var v3404 int32
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3407 int32
	_ = v3407
	var v3446 int32
	_ = v3446
	var v3449 int32
	_ = v3449
	var v3450 int32
	_ = v3450
	var v3454 int32
	_ = v3454
	var v3456 int32
	_ = v3456
	var v3457 int64
	_ = v3457
	var v3465 int32
	_ = v3465
	var v3469 int32
	_ = v3469
	var v3473 int32
	_ = v3473
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3497 int32
	_ = v3497
	var v3500 int32
	_ = v3500
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3513 int32
	_ = v3513
	var v3519 int32
	_ = v3519
	var v3521 int32
	_ = v3521
	var v3525 int32
	_ = v3525
	var v3533 int32
	_ = v3533
	var v3540 int32
	_ = v3540
	var v3542 int64
	_ = v3542
	var v3545 int64
	_ = v3545
	var v3546 int64
	_ = v3546
	var v3551 int64
	_ = v3551
	var v3554 int64
	_ = v3554
	var v3557 int64
	_ = v3557
	var v3560 int64
	_ = v3560
	var v3561 int64
	_ = v3561
	var v3565 int64
	_ = v3565
	var v3572 int64
	_ = v3572
	var v3583 int32
	_ = v3583
	var v3584 int64
	_ = v3584
	var v3585 int64
	_ = v3585
	var v3589 int64
	_ = v3589
	var v3590 int64
	_ = v3590
	var v3596 int64
	_ = v3596
	var v3597 int64
	_ = v3597
	var v3599 int64
	_ = v3599
	var v3601 int64
	_ = v3601
	var v3602 int64
	_ = v3602
	var v3612 int64
	_ = v3612
	var v3613 int64
	_ = v3613
	var v3624 int64
	_ = v3624
	var v3626 int64
	_ = v3626
	var v3628 float64
	_ = v3628
	var v3639 int64
	_ = v3639
	var v3640 int64
	_ = v3640
	var v3653 int32
	_ = v3653
	var v3656 int64
	_ = v3656
	var v3657 int64
	_ = v3657
	var v3662 int64
	_ = v3662
	var v3665 int64
	_ = v3665
	var v3668 int64
	_ = v3668
	var v3671 int64
	_ = v3671
	var v3672 int64
	_ = v3672
	var v3676 int64
	_ = v3676
	var v3683 int64
	_ = v3683
	var v3694 int64
	_ = v3694
	var v3695 int64
	_ = v3695
	var v3700 int64
	_ = v3700
	var v3701 int64
	_ = v3701
	var v3711 float64
	_ = v3711
	var v3715 int64
	_ = v3715
	var v3717 int64
	_ = v3717
	var v3719 float64
	_ = v3719
	var v3730 int64
	_ = v3730
	var v3731 int64
	_ = v3731
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3754 int64
	_ = v3754
	var v3755 float64
	_ = v3755
	var v3756 int64
	_ = v3756
	var v3757 int32
	_ = v3757
	var v3761 int32
	_ = v3761
	var v3763 int32
	_ = v3763
	var v3765 int64
	_ = v3765
	var v3770 int64
	_ = v3770
	var v3771 int64
	_ = v3771
	var v3775 int64
	_ = v3775
	var v3776 int64
	_ = v3776
	var v3786 float64
	_ = v3786
	var v3790 int64
	_ = v3790
	var v3792 int64
	_ = v3792
	var v3794 float64
	_ = v3794
	var v3805 int64
	_ = v3805
	var v3806 int64
	_ = v3806
	var v3817 int32
	_ = v3817
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3834 int32
	_ = v3834
	var v3842 int32
	_ = v3842
	var v3843 int32
	_ = v3843
	var v3845 int64
	_ = v3845
	var v3846 float64
	_ = v3846
	var v3847 int64
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3852 int32
	_ = v3852
	var v3854 int32
	_ = v3854
	var v3856 int64
	_ = v3856
	var v3861 int64
	_ = v3861
	var v3862 int64
	_ = v3862
	var v3866 int64
	_ = v3866
	var v3867 int64
	_ = v3867
	var v3877 float64
	_ = v3877
	var v3881 int64
	_ = v3881
	var v3883 int64
	_ = v3883
	var v3885 float64
	_ = v3885
	var v3896 int64
	_ = v3896
	var v3897 int64
	_ = v3897
	var v3908 int32
	_ = v3908
	var v3918 int32
	_ = v3918
	var v3919 int32
	_ = v3919
	var v3920 int32
	_ = v3920
	var v3929 int32
	_ = v3929
	var v3946 int32
	_ = v3946
	var v3987 int32
	_ = v3987
	var v4014 int32
	_ = v4014
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4024 int32
	_ = v4024
	var v4025 int32
	_ = v4025
	var v4026 int32
	_ = v4026
	var v4029 int64
	_ = v4029
	var v4030 int64
	_ = v4030
	var v4033 int64
	_ = v4033
	var v4039 int32
	_ = v4039
	var v4041 int64
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int32
	_ = v4045
	var v4050 int32
	_ = v4050
	var v4054 int32
	_ = v4054
	var v4059 int32
	_ = v4059
	var v4067 int32
	_ = v4067
	var v4069 int32
	_ = v4069
	var v4073 int32
	_ = v4073
	var v4078 int32
	_ = v4078
	var v4085 int32
	_ = v4085
	var v4087 int32
	_ = v4087
	v2 = int64(0)
	v9 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(528)
	m.G0 = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v40)+512)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v40)+520)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v40)+504)) = v2
	v60 = F_ParseDateTime(m, v44, v40+int32(16), int32(256), v40+int32(384), v40+int32(272), v40+int32(496))
	mBase = m.M
	if v60 == v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v64 = v40 + int32(384)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+496))
	if v42 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v2296 = v60
	goto L3
L3:
	;
	if v2296 == int32(-1) {
		goto L499
	} else {
		goto L500
	}
L4:
	;
	v73 = int32(32767)
	goto L6
L5:
	;
	v73 = int32(base.Ui32(v42) >> (uint(int32(16)) % 32))
	goto L6
L6:
	;
	v74 = m.G0
	v76 = v74 - int32(160)
	m.G0 = v76
	v79 = v40 + int32(500)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = int32(17)
	v83 = v40 + int32(504)
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v84
	v86 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v86
	if v67 <= v84 {
		v173 = v9
		goto L10
	} else {
		goto L11
	}
L7:
	;
	m.G0 = v76 + int32(160)
	v2296 = v2251
	goto L3
L8:
	;
	v257 = v9
	v258 = int32(8)
	v261 = v214
	v273 = v9
	v275 = v9
	goto L20
L9:
	;
	v214 = v67 - int32(1)
	v219 = base.B2i32(v97 == int32(45))
	goto L8
L10:
	;
	v191 = v67 - int32(1)
	if v191 < int32(0) {
		v2251 = int32(-1)
		goto L7
	} else {
		goto L19
	}
L11:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1114]))
	if v93 != int32(2) {
		v173 = v9
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97 != int32(45) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v67) < base.Ui32(int32(2)) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v111 = int32(1)
	goto L15
L15:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v64+v111<<(uint(int32(2))%32))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	switch v145 - int32(43) {
	case 0, 2:
		v173 = int32(0)
		goto L10
	default:
		goto L17
	}
L16:
	;
	v173 = v148
	goto L10
L17:
	;
	v148 = int32(1)
	v150 = v111 + v148
	if v150 != v67 {
		v111 = v150
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v214 = v191
	v219 = v173
	goto L8
L20:
	;
	v285 = int32(-1)
	v287 = v261 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(272)+v287)))
	switch v289 {
	case 0, 2:
		goto L27
	case 1, 6:
		goto L26
	case 3:
		goto L29
	case 4:
		goto L28
	default:
		v2251 = v285
		goto L7
	}
L21:
	;
	v2210 = int32(0)
	v2213 = v2178 | base.B2i32(v2194 == v2210)
	v2214 = v2210 - v2213
	if v2213 != 0 {
		v2251 = v2214
		goto L7
	} else {
		goto L493
	}
L22:
	;
	if int32(0) < v261 {
		v257 = v2178
		v258 = v2179
		v261 = v261 - int32(1)
		v273 = v2194
		v275 = v2196
		goto L20
	} else {
		goto L492
	}
L23:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v76)+112))
	if v2165&v273 != 0 {
		goto L489
	} else {
		goto L490
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v1020 = v287 + v64
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1026 = F_strtox_2(m, v1021, v76+int32(116), int32(10), int64(-9223372036854775807-1))
	mBase = m.M
	goto L174
L25:
	;
	v1016 = int32(18)
	goto L24
L26:
	;
	if v257&int32(1) != 0 {
		v2251 = v285
		goto L7
	} else {
		goto L72
	}
L27:
	;
	if v258 != int32(8) {
		v1016 = v258
		goto L24
	} else {
		goto L58
	}
L28:
	;
	v383 = v287 + v64
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v386 = v384 + int32(1)
	v387 = int32(58)
	v388 = F___strchrnul(m, v386, v387)
	mBase = m.M
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v390 == v387 {
		goto L40
	} else {
		goto L41
	}
L29:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v64)))
	v296 = F_DecodeTimeCommon(m, v291, v73, v76+int32(112), v76+int32(120))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	if v296 != 0 {
		v2251 = v296
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v300 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v300
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v76)+136))
	v305 = int64(3600000000)
	v306 = int64(0)
	v311 = int64(32)
	v314 = int64(base.Ui64(v302) >> (uint(v311) % 64))
	v317 = int64(4294967295)
	v320 = v302 & v317
	v321 = v305 * v320
	v325 = int64(base.Ui64(v321)>>(uint(v311)%64)) + v305*v314
	v332 = v320*v306 + v325&v317
	*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v302*v306 + v302>>(uint(int64(63))%64)*v305 + v306*v314 + int64(base.Ui64(v325)>>(uint(v311)%64)) + int64(base.Ui64(v332)>>(uint(v311)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = v321&v317 | v332<<(uint(v311)%64)
	goto L33
L33:
	;
	v343 = int32(-2)
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	if v344 != v345>>(uint(int64(63))%64) {
		v2251 = v343
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v349 = v300 + v345
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v349
	if base.B2i32(v345 < int64(0))^base.B2i32(v349 < v300) != 0 {
		v2251 = v343
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v355 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+128)))
	v357 = v355 * int64(60000000)
	v358 = v349 + v357
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v358
	if base.B2i32(v357 < int64(0))^base.B2i32(v358 < v349) != 0 {
		v2251 = v343
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v364 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+124)))
	v366 = v364 * int64(1000000)
	v367 = v358 + v366
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v367
	if base.B2i32(v366 < int64(0))^base.B2i32(v367 < v358) != 0 {
		v2251 = v343
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v373 = int32(21)
	v374 = int32(0)
	if base.B2i32(v219 == v374)|base.B2i32(v367 <= int64(0)) != 0 {
		v2137 = v374
		v2138 = v373
		v2155 = v275
		goto L23
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v367
	v2137 = v374
	v2138 = v373
	v2155 = v275
	goto L23
L39:
	;
	if v394 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L40:
	;
	v394 = v388
	goto L42
L41:
	;
	v394 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v401 = F_DecodeTimeCommon(m, v386, v73, v76+int32(112), v76+int32(120))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	if v401 != 0 {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v403 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v403
	v406 = v76 + int32(96)
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v76)+136))
	v410 = int64(3600000000)
	v411 = int64(0)
	v416 = int64(32)
	v419 = int64(base.Ui64(v407) >> (uint(v416) % 64))
	v422 = int64(4294967295)
	v425 = v407 & v422
	v426 = v410 * v425
	v430 = int64(base.Ui64(v426)>>(uint(v416)%64)) + v410*v419
	v437 = v425*v411 + v430&v422
	*(*int64)(unsafe.Add(mBase, uint32(v406)+8)) = v407*v411 + v407>>(uint(int64(63))%64)*v410 + v411*v419 + int64(base.Ui64(v430)>>(uint(v416)%64)) + int64(base.Ui64(v437)>>(uint(v416)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v406))) = v426&v422 | v437<<(uint(v416)%64)
	goto L46
L46:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
	if v448 != v449>>(uint(int64(63))%64) {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v453 = v403 + v449
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v453
	if base.B2i32(v449 < int64(0))^base.B2i32(v453 < v403) != 0 {
		goto L27
	} else {
		goto L48
	}
L48:
	;
	v459 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+128)))
	v461 = v459 * int64(60000000)
	v462 = v453 + v461
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v462
	if base.B2i32(v461 < int64(0))^base.B2i32(v462 < v453) != 0 {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	v468 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+124)))
	v470 = v468 * int64(1000000)
	v471 = v462 + v470
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v471
	if base.B2i32(v470 < int64(0))^base.B2i32(v471 < v462) != 0 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v478 == int32(45) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v471 == int64(-9223372036854775807-1) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v487 = v471
	goto L53
L53:
	;
	v488 = int32(21)
	v489 = int32(0)
	if v219&base.B2i32(int64(0) < v487) == v489 {
		v2137 = v489
		v2138 = v488
		v2155 = v275
		goto L23
	} else {
		goto L57
	}
L54:
	;
	v2251 = int32(-2)
	goto L7
L55:
	;
	goto L56
L56:
	;
	v485 = int64(0) - v471
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v485
	v487 = v485
	goto L53
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v487
	v2137 = v489
	v2138 = v488
	v2155 = v275
	goto L23
L58:
	;
	if base.B2i32(int32(3071) < v73) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v1016 = int32(19)
	goto L24
L60:
	;
	if base.B2i32(v73 == int32(2048)) == int32(0) {
		goto L25
	} else {
		goto L71
	}
L61:
	;
	v1016 = int32(20)
	goto L24
L62:
	;
	v1016 = int32(21)
	goto L24
L63:
	;
	v1016 = int32(23)
	goto L24
L64:
	;
	switch v73 - int32(2) {
	case 0, 4:
		goto L63
	case 1, 3, 5:
		goto L25
	case 2:
		v1016 = int32(25)
		goto L24
	case 6:
		goto L62
	default:
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if base.B2i32(int32(6143) < v73) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	switch v73 - int32(1024) {
	case 0, 8:
		goto L61
	case 1, 2, 3, 4, 5, 6, 7:
		goto L25
	default:
		goto L60
	}
L68:
	;
	switch v73 - int32(3072) {
	case 0, 8:
		goto L59
	default:
		goto L25
	}
L69:
	;
	goto L70
L70:
	;
	goto L25
L71:
	;
	goto L59
L72:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v287+v64)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1115])))
	if v523 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v990 = int32(8)
	v993 = v967 & int32(255)
	if v993 == v990 {
		v2178 = int32(0)
		v2179 = v990
		v2194 = v273
		v2196 = v275
		goto L22
	} else {
		goto L168
	}
L74:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1042])))
	if v758 != 0 {
		goto L123
	} else {
		goto L124
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1115]))) = v686
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+11)))
	if v716 != int32(31) {
		v961 = v686
		v967 = v716
		goto L73
	} else {
		goto L121
	}
L76:
	;
	goto L81
L77:
	;
	goto L78
L78:
	;
	v571 = int32(*(*int8)(unsafe.Add(mBase, uint32(v520))))
	v583 = int32(1611360)
	v584 = int32(1610400)
	goto L94
L79:
	;
	if v560-v561 == int32(0) {
		v686 = v523
		goto L75
	} else {
		goto L93
	}
L81:
	;
	goto L82
L82:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v530 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v531 = v520
	v532 = v523
	v533 = int32(10)
	v534 = v530
	goto L87
L84:
	;
	v556 = v523
	v560 = int32(0)
	goto L85
L85:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	goto L79
L86:
	;
	v556 = v551
	v560 = v553
	goto L85
L87:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	if v534 != v536 {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v551 = v545
	v553 = int32(0)
	goto L86
L89:
	;
	if v536 == int32(0) {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v541 = v533 - int32(1)
	if v541 == int32(0) {
		v551 = v532
		v553 = v534
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v544 = int32(1)
	v545 = v532 + v544
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)))
	if v546 != 0 {
		v531 = v531 + v544
		v532 = v545
		v533 = v541
		v534 = v546
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L78
L94:
	;
	v616 = v584 + (v583-v584)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v616))))
	v618 = v571 - v617
	if v618 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L74
L96:
	;
	goto L101
L97:
	;
	v668 = v618
	goto L98
L98:
	;
	v672 = base.B2i32(v668 < int32(0))
	if v668 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L99:
	;
	if v659 == int32(0) {
		v686 = v616
		goto L75
	} else {
		goto L113
	}
L101:
	;
	goto L102
L102:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v627 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v628 = v520
	v629 = v616
	v630 = int32(10)
	v631 = v627
	goto L107
L104:
	;
	v653 = v616
	v657 = int32(0)
	goto L105
L105:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653))))
	v659 = v657 - v658
	goto L99
L106:
	;
	v653 = v648
	v657 = v650
	goto L105
L107:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	if v631 != v633 {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L109
	}
L108:
	;
	v648 = v642
	v650 = int32(0)
	goto L106
L109:
	;
	if v633 == int32(0) {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v638 = v630 - int32(1)
	if v638 == int32(0) {
		v648 = v629
		v650 = v631
		goto L106
	} else {
		goto L111
	}
L111:
	;
	v641 = int32(1)
	v642 = v629 + v641
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628)+1)))
	if v643 != 0 {
		v628 = v628 + v641
		v629 = v642
		v630 = v638
		v631 = v643
		goto L107
	} else {
		goto L112
	}
L112:
	;
	goto L108
L113:
	;
	v668 = v659
	goto L98
L114:
	;
	v673 = v616 - int32(16)
	goto L116
L115:
	;
	v673 = v583
	goto L116
L116:
	;
	if v668 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v676 = v584
	goto L119
L118:
	;
	v676 = v616 + int32(16)
	goto L119
L119:
	;
	if base.Ui32(v676) <= base.Ui32(v673) {
		v583 = v673
		v584 = v676
		goto L94
	} else {
		goto L120
	}
L120:
	;
	goto L95
L121:
	;
	goto L74
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+uint32(_consts[1042]))) = v922
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+11)))
	v961 = v922
	v967 = v952
	goto L73
L123:
	;
	goto L128
L124:
	;
	goto L125
L125:
	;
	v806 = int32(*(*int8)(unsafe.Add(mBase, uint32(v520))))
	v818 = int32(1610384)
	v819 = int32(1609248)
	goto L141
L126:
	;
	if v795-v796 == int32(0) {
		v922 = v758
		goto L122
	} else {
		goto L140
	}
L128:
	;
	goto L129
L129:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v765 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v766 = v520
	v767 = v758
	v768 = int32(10)
	v769 = v765
	goto L134
L131:
	;
	v791 = v758
	v795 = int32(0)
	goto L132
L132:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791))))
	goto L126
L133:
	;
	v791 = v786
	v795 = v788
	goto L132
L134:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if v769 != v771 {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L136
	}
L135:
	;
	v786 = v780
	v788 = int32(0)
	goto L133
L136:
	;
	if v771 == int32(0) {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v776 = v768 - int32(1)
	if v776 == int32(0) {
		v786 = v767
		v788 = v769
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v779 = int32(1)
	v780 = v767 + v779
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	if v781 != 0 {
		v766 = v766 + v779
		v767 = v780
		v768 = v776
		v769 = v781
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	goto L125
L141:
	;
	v851 = v819 + (v818-v819)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v852 = int32(*(*int8)(unsafe.Add(mBase, uint32(v851))))
	v853 = v806 - v852
	if v853 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v2251 = int32(-1)
	goto L7
L143:
	;
	goto L148
L144:
	;
	v903 = v853
	goto L145
L145:
	;
	v907 = base.B2i32(v903 < int32(0))
	if v903 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L146:
	;
	if v894 == int32(0) {
		v922 = v851
		goto L122
	} else {
		goto L160
	}
L148:
	;
	goto L149
L149:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v862 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v863 = v520
	v864 = v851
	v865 = int32(10)
	v866 = v862
	goto L154
L151:
	;
	v888 = v851
	v892 = int32(0)
	goto L152
L152:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v894 = v892 - v893
	goto L146
L153:
	;
	v888 = v883
	v892 = v885
	goto L152
L154:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864))))
	if v866 != v868 {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L156
	}
L155:
	;
	v883 = v877
	v885 = int32(0)
	goto L153
L156:
	;
	if v868 == int32(0) {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L157
	}
L157:
	;
	v873 = v865 - int32(1)
	if v873 == int32(0) {
		v883 = v864
		v885 = v866
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v876 = int32(1)
	v877 = v864 + v876
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863)+1)))
	if v878 != 0 {
		v863 = v863 + v876
		v864 = v877
		v865 = v873
		v866 = v878
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v903 = v894
	goto L145
L161:
	;
	v908 = v851 - int32(16)
	goto L163
L162:
	;
	v908 = v818
	goto L163
L163:
	;
	if v903 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v911 = v819
	goto L166
L165:
	;
	v911 = v851 + int32(16)
	goto L166
L166:
	;
	if base.Ui32(v911) <= base.Ui32(v908) {
		v818 = v908
		v819 = v911
		goto L141
	} else {
		goto L167
	}
L167:
	;
	goto L142
L168:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v961)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(0)
	v1000 = int32(-1)
	switch v993 {
	case 0:
		goto L169
	default:
		v2251 = v1000
		goto L7
	case 17:
		v2137 = int32(1)
		v2138 = v996
		v2155 = v275
		goto L23
	case 19:
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(31758)
	if v261 != v214 {
		v2251 = v1000
		goto L7
	} else {
		goto L172
	}
L170:
	;
	if v261 != v214 {
		v2251 = v1000
		goto L7
	} else {
		goto L171
	}
L171:
	;
	v2137 = int32(0)
	v2138 = v996
	v2155 = int32(1)
	goto L23
L172:
	;
	if base.Ui32(int32(1)) < base.Ui32(v996-int32(9)) {
		v2251 = v1000
		goto L7
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v996
	v1012 = int32(0)
	v2137 = v1012
	v2138 = v1012
	v2155 = v275
	goto L23
L174:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1028 == int32(68) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v2251 = int32(-2)
	goto L7
L176:
	;
	goto L177
L177:
	;
	v1032 = float64(0)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v76)+116))
	v1034 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	switch v1034 - int32(45) {
	case 0:
		goto L180
	case 1:
		goto L179
	default:
		goto L181
	}
L178:
	;
	if v219 == int32(0) {
		v1309 = v1292
		v1311 = v1294
		goto L249
	} else {
		goto L250
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+120)) = v1033
	v1124 = v1033 + int32(1)
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1125 == int32(0) {
		v1285 = v1032
		goto L201
	} else {
		goto L202
	}
L180:
	;
	v1045 = F_strtol(m, v1033+int32(1), v76+int32(116), int32(10))
	mBase = m.M
	goto L183
L181:
	;
	if v1034 == int32(0) {
		v1292 = v1026
		v1294 = v1032
		v1296 = v1016
		goto L178
	} else {
		goto L182
	}
L182:
	;
	v2251 = int32(-1)
	goto L7
L183:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1047 == int32(68) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v2251 = int32(-2)
	goto L7
L185:
	;
	goto L186
L186:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1045) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v2251 = int32(-2)
	goto L7
L188:
	;
	goto L189
L189:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v76)+116))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054))))
	if v1055 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v2251 = int32(-1)
	goto L7
L191:
	;
	goto L192
L192:
	;
	v1058 = v76 + int32(80)
	v1061 = int64(12)
	v1062 = int64(0)
	v1067 = int64(32)
	v1070 = int64(base.Ui64(v1026) >> (uint(v1067) % 64))
	v1073 = int64(4294967295)
	v1076 = v1026 & v1073
	v1077 = v1061 * v1076
	v1081 = int64(base.Ui64(v1077)>>(uint(v1067)%64)) + v1061*v1070
	v1088 = v1076*v1062 + v1081&v1073
	*(*int64)(unsafe.Add(mBase, uint32(v1058)+8)) = v1026*v1062 + v1026>>(uint(int64(63))%64)*v1061 + v1062*v1070 + int64(base.Ui64(v1081)>>(uint(v1067)%64)) + int64(base.Ui64(v1088)>>(uint(v1067)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1058))) = v1077&v1073 | v1088<<(uint(v1067)%64)
	goto L193
L193:
	;
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
	if v1099 != v1100>>(uint(int64(63))%64) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v2251 = int32(-2)
	goto L7
L195:
	;
	goto L196
L196:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	if v1109 == int32(45) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1112 = int32(0) - v1045
	goto L199
L198:
	;
	v1112 = v1045
	goto L199
L199:
	;
	v1113 = base.I64_extend_i32_s(v1112)
	v1116 = v1113 + v1100
	if base.B2i32(v1113 < int64(0))^base.B2i32(v1116 < v1100) == int32(0) {
		v1292 = v1116
		v1294 = v1032
		v1296 = int32(23)
		goto L178
	} else {
		goto L200
	}
L200:
	;
	v2251 = int32(-2)
	goto L7
L201:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287))))
	if v1288 == int32(45) {
		goto L246
	} else {
		goto L247
	}
L202:
	;
	v1128 = int32(515387)
	v1132 = m.G0
	v1134 = v1132 - int32(32)
	v1135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+24)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+16)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+8)) = v1135
	*(*int64)(unsafe.Add(mBase, uint32(v1134))) = v1135
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v1143 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v2251 = int32(-1)
	goto L7
L204:
	;
	if v1124&int32(3) == int32(0) {
		v1235 = v1124
		goto L227
	} else {
		goto L228
	}
L205:
	;
	v1211 = int32(0)
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v1147 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1151 = v1124
	goto L211
L209:
	;
	goto L210
L210:
	;
	v1161 = v1128
	v1162 = v1143
	goto L214
L211:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	if v1157 == v1143 {
		v1151 = v1151 + int32(1)
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v1211 = v1151 - v1124
	goto L204
L213:
	;
	goto L212
L214:
	;
	v1169 = v1134 + int32(base.Ui32(v1162)>>(uint(int32(3))%32))&int32(28)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	v1171 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1169))) = v1170 | v1171<<(uint(v1162)%32)
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1161)+1)))
	if v1175 != 0 {
		v1161 = v1161 + v1171
		v1162 = v1175
		goto L214
	} else {
		goto L216
	}
L215:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1178 == int32(0) {
		v1203 = v1124
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L215
L217:
	;
	v1211 = v1203 - v1124
	goto L204
L218:
	;
	v1182 = v1124
	v1183 = v1178
	goto L219
L219:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1134+int32(base.Ui32(v1183)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v1191)>>(uint(v1183)%32))&int32(1) == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1203 = v1199
	goto L217
L221:
	;
	v1203 = v1182
	goto L217
L222:
	;
	goto L223
L223:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182)+1)))
	v1199 = v1182 + int32(1)
	if v1197 != 0 {
		v1182 = v1199
		v1183 = v1197
		goto L219
	} else {
		goto L224
	}
L224:
	;
	goto L220
L225:
	;
	if v1211 != v1268 {
		goto L203
	} else {
		goto L242
	}
L226:
	;
	v1268 = v1260 - v1124
	goto L225
L227:
	;
	v1239 = v1235
	goto L236
L228:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	if v1219 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1268 = int32(0)
	goto L225
L230:
	;
	goto L231
L231:
	;
	v1224 = v1124
	goto L232
L232:
	;
	v1228 = v1224 + int32(1)
	if v1228&int32(3) == int32(0) {
		v1235 = v1228
		goto L227
	} else {
		goto L234
	}
L233:
	;
	v1260 = v1228
	goto L226
L234:
	;
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228))))
	if v1233 != 0 {
		v1224 = v1228
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	v1248 = int32(-2139062144)
	if (int32(16843008)-v1245|v1245)&v1248 == v1248 {
		v1239 = v1239 + int32(4)
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v1254 = v1239
	goto L239
L238:
	;
	goto L237
L239:
	;
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	if v1258 != 0 {
		v1254 = v1254 + int32(1)
		goto L239
	} else {
		goto L241
	}
L240:
	;
	v1260 = v1254
	goto L226
L241:
	;
	goto L240
L242:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v1275 = F_strtod(m, v1033, v76+int32(120))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L30
	} else {
		goto L243
	}
L243:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v76)+120))
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1277))))
	if v1278 != 0 {
		goto L203
	} else {
		goto L244
	}
L244:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1280 == int32(0) {
		v1285 = v1275
		goto L201
	} else {
		goto L245
	}
L245:
	;
	goto L203
L246:
	;
	v1291 = base.F64_neg(v1285)
	goto L248
L247:
	;
	v1291 = v1285
	goto L248
L248:
	;
	v1292 = v1026
	v1294 = v1291
	v1296 = v1016
	goto L178
L249:
	;
	switch v1296 - int32(18) {
	case 0:
		goto L262
	case 1:
		goto L261
	case 2:
		goto L260
	case 3:
		goto L259
	case 4:
		goto L258
	case 5:
		goto L257
	default:
		v2251 = int32(-1)
		goto L7
	case 7:
		goto L256
	case 8:
		goto L255
	case 9:
		goto L254
	case 10:
		goto L253
	case 11:
		goto L263
	case 12:
		goto L264
	}
L250:
	;
	v1301 = v1292 >> (uint(int64(63)) % 64)
	v1303 = v1301 - (v1292 ^ v1301)
	if base.F64_gt(v1294, float64(0)) == int32(0) {
		v1309 = v1303
		v1311 = v1294
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v1309 = v1303
	v1311 = base.F64_neg(v1294)
	goto L249
L252:
	;
	v2137 = int32(0)
	v2138 = int32(21)
	v2155 = v275
	goto L23
L253:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L475
	} else {
		goto L476
	}
L254:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L461
	} else {
		goto L462
	}
L255:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L447
	} else {
		goto L448
	}
L256:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L436
	} else {
		goto L437
	}
L257:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L410
	} else {
		goto L411
	}
L258:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L381
	} else {
		goto L382
	}
L259:
	;
	if base.Ui64(v1309-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L362
	} else {
		goto L363
	}
L260:
	;
	v1644 = v76 - int32(-64)
	v1647 = int64(3600000000)
	v1648 = int64(0)
	v1653 = int64(32)
	v1656 = int64(base.Ui64(v1309) >> (uint(v1653) % 64))
	v1659 = int64(4294967295)
	v1662 = v1309 & v1659
	v1663 = v1647 * v1662
	v1667 = int64(base.Ui64(v1663)>>(uint(v1653)%64)) + v1647*v1656
	v1674 = v1662*v1648 + v1667&v1659
	*(*int64)(unsafe.Add(mBase, uint32(v1644)+8)) = v1309*v1648 + v1309>>(uint(int64(63))%64)*v1647 + v1648*v1656 + int64(base.Ui64(v1667)>>(uint(v1653)%64)) + int64(base.Ui64(v1674)>>(uint(v1653)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1644))) = v1663&v1659 | v1674<<(uint(v1653)%64)
	goto L342
L261:
	;
	v1550 = v76 + int32(48)
	v1553 = int64(60000000)
	v1554 = int64(0)
	v1559 = int64(32)
	v1562 = int64(base.Ui64(v1309) >> (uint(v1559) % 64))
	v1565 = int64(4294967295)
	v1568 = v1309 & v1565
	v1569 = v1553 * v1568
	v1573 = int64(base.Ui64(v1569)>>(uint(v1559)%64)) + v1553*v1562
	v1580 = v1568*v1554 + v1573&v1565
	*(*int64)(unsafe.Add(mBase, uint32(v1550)+8)) = v1309*v1554 + v1309>>(uint(int64(63))%64)*v1553 + v1554*v1562 + int64(base.Ui64(v1573)>>(uint(v1559)%64)) + int64(base.Ui64(v1580)>>(uint(v1559)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1550))) = v1569&v1565 | v1580<<(uint(v1559)%64)
	goto L322
L262:
	;
	v1454 = v76 + int32(32)
	v1457 = int64(1000000)
	v1458 = int64(0)
	v1463 = int64(32)
	v1466 = int64(base.Ui64(v1309) >> (uint(v1463) % 64))
	v1469 = int64(4294967295)
	v1472 = v1309 & v1469
	v1473 = v1457 * v1472
	v1477 = int64(base.Ui64(v1473)>>(uint(v1463)%64)) + v1457*v1466
	v1484 = v1472*v1458 + v1477&v1469
	*(*int64)(unsafe.Add(mBase, uint32(v1454)+8)) = v1309*v1458 + v1309>>(uint(int64(63))%64)*v1457 + v1458*v1466 + int64(base.Ui64(v1477)>>(uint(v1463)%64)) + int64(base.Ui64(v1484)>>(uint(v1463)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1454))) = v1473&v1469 | v1484<<(uint(v1463)%64)
	goto L301
L263:
	;
	v1360 = v76 + int32(16)
	v1363 = int64(1000)
	v1364 = int64(0)
	v1369 = int64(32)
	v1372 = int64(base.Ui64(v1309) >> (uint(v1369) % 64))
	v1375 = int64(4294967295)
	v1378 = v1309 & v1375
	v1379 = v1363 * v1378
	v1383 = int64(base.Ui64(v1379)>>(uint(v1369)%64)) + v1363*v1372
	v1390 = v1378*v1364 + v1383&v1375
	*(*int64)(unsafe.Add(mBase, uint32(v1360)+8)) = v1309*v1364 + v1309>>(uint(int64(63))%64)*v1363 + v1364*v1372 + int64(base.Ui64(v1383)>>(uint(v1369)%64)) + int64(base.Ui64(v1390)>>(uint(v1369)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1360))) = v1379&v1375 | v1390<<(uint(v1369)%64)
	goto L281
L264:
	;
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1316 = v1315 + v1309
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1316
	if base.B2i32(v1309 < int64(0))^base.B2i32(v1316 < v1315) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v2251 = int32(-2)
	goto L7
L266:
	;
	goto L267
L267:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if base.F64_lt(base.F64_abs(v1311), float64(9.223372036854776e+18)) != 0 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(16384)
	v2137 = int32(0)
	v2138 = int32(30)
	v2155 = v275
	goto L23
L271:
	;
	v1344 = v1343 + v1316
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1344
	if base.B2i32(v1343 < int64(0))^base.B2i32(v1344 < v1316) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L280
	}
L272:
	;
	v1332 = base.F64_sub(v1311, base.F64_convert_i64_s(v1330))
	if base.F64_gt(v1332, float64(0.5)) != 0 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	v1328 = base.I64_trunc_f64_s(v1311)
	v1330 = v1328
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1330 = int64(-9223372036854775807 - 1)
	goto L272
L276:
	;
	v1343 = v1330 + int64(1)
	goto L271
L277:
	;
	goto L278
L278:
	;
	if base.F64_lt(v1332, float64(-0.5)) == int32(0) {
		v1343 = v1330
		goto L271
	} else {
		goto L279
	}
L279:
	;
	v1343 = v1330 - int64(1)
	goto L271
L280:
	;
	goto L270
L281:
	;
	v1401 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
	v1402 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
	if v1401 != v1402>>(uint(int64(63))%64) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v2251 = int32(-2)
	goto L7
L283:
	;
	goto L284
L284:
	;
	v1407 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1408 = v1407 + v1402
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1408
	if base.B2i32(v1402 < int64(0))^base.B2i32(v1408 < v1407) != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v2251 = int32(-2)
	goto L7
L286:
	;
	goto L287
L287:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1418 = base.F64_mul(v1311, float64(1000))
	if base.F64_lt(base.F64_abs(v1418), float64(9.223372036854776e+18)) != 0 {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(8192)
	v2137 = int32(0)
	v2138 = int32(29)
	v2155 = v275
	goto L23
L291:
	;
	v1426 = base.F64_sub(v1418, base.F64_convert_i64_s(v1424))
	if base.F64_gt(v1426, float64(0.5)) != 0 {
		goto L296
	} else {
		goto L297
	}
L292:
	;
	v1422 = base.I64_trunc_f64_s(v1418)
	v1424 = v1422
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1424 = int64(-9223372036854775807 - 1)
	goto L291
L295:
	;
	v1438 = v1437 + v1408
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1438
	if base.B2i32(v1437 < int64(0))^base.B2i32(v1438 < v1408) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L300
	}
L296:
	;
	v1437 = v1424 + int64(1)
	goto L295
L297:
	;
	goto L298
L298:
	;
	if base.F64_lt(v1426, float64(-0.5)) == int32(0) {
		v1437 = v1424
		goto L295
	} else {
		goto L299
	}
L299:
	;
	v1437 = v1424 - int64(1)
	goto L295
L300:
	;
	goto L290
L301:
	;
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
	v1496 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
	if v1495 != v1496>>(uint(int64(63))%64) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v2251 = int32(-2)
	goto L7
L303:
	;
	goto L304
L304:
	;
	v1501 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1502 = v1501 + v1496
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1502
	if base.B2i32(v1496 < int64(0))^base.B2i32(v1502 < v1501) != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v2251 = int32(-2)
	goto L7
L306:
	;
	goto L307
L307:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v2137 = int32(0)
	v2138 = int32(18)
	v2155 = v275
	goto L23
L309:
	;
	v1512 = base.F64_mul(v1311, float64(1e+06))
	if base.F64_lt(base.F64_abs(v1512), float64(9.223372036854776e+18)) != 0 {
		goto L313
	} else {
		goto L314
	}
L310:
	;
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(4096)
	goto L308
L312:
	;
	v1520 = base.F64_sub(v1512, base.F64_convert_i64_s(v1518))
	if base.F64_gt(v1520, float64(0.5)) != 0 {
		goto L317
	} else {
		goto L318
	}
L313:
	;
	v1516 = base.I64_trunc_f64_s(v1512)
	v1518 = v1516
	goto L312
L314:
	;
	goto L315
L315:
	;
	v1518 = int64(-9223372036854775807 - 1)
	goto L312
L316:
	;
	v1532 = v1531 + v1502
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1532
	if base.B2i32(v1531 < int64(0))^base.B2i32(v1532 < v1502) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L321
	}
L317:
	;
	v1531 = v1518 + int64(1)
	goto L316
L318:
	;
	goto L319
L319:
	;
	if base.F64_lt(v1520, float64(-0.5)) == int32(0) {
		v1531 = v1518
		goto L316
	} else {
		goto L320
	}
L320:
	;
	v1531 = v1518 - int64(1)
	goto L316
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(28672)
	goto L308
L322:
	;
	v1591 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
	v1592 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
	if v1591 != v1592>>(uint(int64(63))%64) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v2251 = int32(-2)
	goto L7
L324:
	;
	goto L325
L325:
	;
	v1597 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1598 = v1597 + v1592
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1598
	if base.B2i32(v1592 < int64(0))^base.B2i32(v1598 < v1597) != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2251 = int32(-2)
	goto L7
L327:
	;
	goto L328
L328:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1608 = base.F64_mul(v1311, float64(6e+07))
	if base.F64_lt(base.F64_abs(v1608), float64(9.223372036854776e+18)) != 0 {
		goto L333
	} else {
		goto L334
	}
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(2048)
	v2137 = int32(0)
	v2138 = int32(19)
	v2155 = v275
	goto L23
L332:
	;
	v1616 = base.F64_sub(v1608, base.F64_convert_i64_s(v1614))
	if base.F64_gt(v1616, float64(0.5)) != 0 {
		goto L337
	} else {
		goto L338
	}
L333:
	;
	v1612 = base.I64_trunc_f64_s(v1608)
	v1614 = v1612
	goto L332
L334:
	;
	goto L335
L335:
	;
	v1614 = int64(-9223372036854775807 - 1)
	goto L332
L336:
	;
	v1628 = v1627 + v1598
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1628
	if base.B2i32(v1627 < int64(0))^base.B2i32(v1628 < v1598) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L341
	}
L337:
	;
	v1627 = v1614 + int64(1)
	goto L336
L338:
	;
	goto L339
L339:
	;
	if base.F64_lt(v1616, float64(-0.5)) == int32(0) {
		v1627 = v1614
		goto L336
	} else {
		goto L340
	}
L340:
	;
	v1627 = v1614 - int64(1)
	goto L336
L341:
	;
	goto L331
L342:
	;
	v1685 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
	v1686 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
	if v1685 != v1686>>(uint(int64(63))%64) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2251 = int32(-2)
	goto L7
L344:
	;
	goto L345
L345:
	;
	v1691 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1692 = v1691 + v1686
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1692
	if base.B2i32(v1686 < int64(0))^base.B2i32(v1692 < v1691) != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v2251 = int32(-2)
	goto L7
L347:
	;
	goto L348
L348:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1702 = base.F64_mul(v1311, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v1702), float64(9.223372036854776e+18)) != 0 {
		goto L353
	} else {
		goto L354
	}
L350:
	;
	goto L351
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(1024)
	goto L252
L352:
	;
	v1710 = base.F64_sub(v1702, base.F64_convert_i64_s(v1708))
	if base.F64_gt(v1710, float64(0.5)) != 0 {
		goto L357
	} else {
		goto L358
	}
L353:
	;
	v1706 = base.I64_trunc_f64_s(v1702)
	v1708 = v1706
	goto L352
L354:
	;
	goto L355
L355:
	;
	v1708 = int64(-9223372036854775807 - 1)
	goto L352
L356:
	;
	v1722 = v1721 + v1692
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1722
	if base.B2i32(v1721 < int64(0))^base.B2i32(v1722 < v1692) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L361
	}
L357:
	;
	v1721 = v1708 + int64(1)
	goto L356
L358:
	;
	goto L359
L359:
	;
	if base.F64_lt(v1710, float64(-0.5)) == int32(0) {
		v1721 = v1708
		goto L356
	} else {
		goto L360
	}
L360:
	;
	v1721 = v1708 - int64(1)
	goto L356
L361:
	;
	goto L351
L362:
	;
	v2251 = int32(-2)
	goto L7
L363:
	;
	goto L364
L364:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1741 = base.I32_wrap_i64(v1309)
	v1742 = v1740 + v1741
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1742
	if base.B2i32(v1741 < int32(0))^base.B2i32(v1742 < v1740) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v2251 = int32(-2)
	goto L7
L366:
	;
	goto L367
L367:
	;
	if base.F64_ne(v1311, float64(0)) != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1752 = base.F64_mul(v1311, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1752), float64(9.223372036854776e+18)) != 0 {
		goto L372
	} else {
		goto L373
	}
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(8)
	goto L252
L371:
	;
	v1760 = base.F64_sub(v1752, base.F64_convert_i64_s(v1758))
	if base.F64_gt(v1760, float64(0.5)) != 0 {
		goto L376
	} else {
		goto L377
	}
L372:
	;
	v1756 = base.I64_trunc_f64_s(v1752)
	v1758 = v1756
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1758 = int64(-9223372036854775807 - 1)
	goto L371
L375:
	;
	v1772 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1773 = v1772 + v1771
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1773
	if base.B2i32(v1771 < int64(0))^base.B2i32(v1773 < v1772) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L380
	}
L376:
	;
	v1771 = v1758 + int64(1)
	goto L375
L377:
	;
	goto L378
L378:
	;
	if base.F64_lt(v1760, float64(-0.5)) == int32(0) {
		v1771 = v1758
		goto L375
	} else {
		goto L379
	}
L379:
	;
	v1771 = v1758 - int64(1)
	goto L375
L380:
	;
	goto L370
L381:
	;
	v2251 = int32(-2)
	goto L7
L382:
	;
	goto L383
L383:
	;
	v1793 = v1309 * int64(7)
	v1797 = base.I32_wrap_i64(v1793)
	if base.I32_wrap_i64(int64(base.Ui64(v1793)>>(uint(int64(32))%64))) != v1797>>(uint(int32(31))%32) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v2251 = int32(-2)
	goto L7
L385:
	;
	goto L386
L386:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1803 = v1802 + v1797
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1803
	if base.B2i32(v1797 < int32(0))^base.B2i32(v1803 < v1802) != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v2251 = int32(-2)
	goto L7
L388:
	;
	goto L389
L389:
	;
	if base.F64_eq(v1311, float64(0)) != 0 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(16777216)
	v2137 = int32(0)
	v2138 = int32(22)
	v2155 = v275
	goto L23
L391:
	;
	v1813 = base.F64_mul(v1311, float64(7))
	if base.F64_lt(base.F64_abs(v1813), float64(2.147483648e+09)) != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v1820 = v1819 + v1803
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1820
	if base.B2i32(v1819 < int32(0))^base.B2i32(v1820 < v1803) != 0 {
		goto L396
	} else {
		goto L397
	}
L393:
	;
	v1817 = base.I32_trunc_f64_s(v1813)
	v1819 = v1817
	goto L392
L394:
	;
	goto L395
L395:
	;
	v1819 = int32(-2147483648)
	goto L392
L396:
	;
	v2251 = int32(-2)
	goto L7
L397:
	;
	goto L398
L398:
	;
	v1828 = base.F64_sub(v1813, base.F64_convert_i32_s(v1819))
	if base.F64_eq(v1828, float64(0)) != 0 {
		goto L390
	} else {
		goto L399
	}
L399:
	;
	v1832 = base.F64_mul(v1828, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1832), float64(9.223372036854776e+18)) != 0 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1840 = base.F64_sub(v1832, base.F64_convert_i64_s(v1838))
	if base.F64_gt(v1840, float64(0.5)) != 0 {
		goto L405
	} else {
		goto L406
	}
L401:
	;
	v1836 = base.I64_trunc_f64_s(v1832)
	v1838 = v1836
	goto L400
L402:
	;
	goto L403
L403:
	;
	v1838 = int64(-9223372036854775807 - 1)
	goto L400
L404:
	;
	v1852 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1853 = v1852 + v1851
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1853
	if base.B2i32(v1851 < int64(0))^base.B2i32(v1853 < v1852) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L409
	}
L405:
	;
	v1851 = v1838 + int64(1)
	goto L404
L406:
	;
	goto L407
L407:
	;
	if base.F64_lt(v1840, float64(-0.5)) == int32(0) {
		v1851 = v1838
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v1851 = v1838 - int64(1)
	goto L404
L409:
	;
	goto L390
L410:
	;
	v2251 = int32(-2)
	goto L7
L411:
	;
	goto L412
L412:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1876 = base.I32_wrap_i64(v1309)
	v1877 = v1875 + v1876
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v1877
	if base.B2i32(v1876 < int32(0))^base.B2i32(v1877 < v1875) != 0 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v2251 = int32(-2)
	goto L7
L414:
	;
	goto L415
L415:
	;
	if base.F64_eq(v1311, float64(0)) != 0 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(2)
	v2137 = int32(0)
	v2138 = int32(23)
	v2155 = v275
	goto L23
L417:
	;
	v1887 = base.F64_mul(v1311, float64(30))
	if base.F64_lt(base.F64_abs(v1887), float64(2.147483648e+09)) != 0 {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v1895 = v1893 + v1894
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v1895
	if base.B2i32(v1893 < int32(0))^base.B2i32(v1895 < v1894) != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	v1891 = base.I32_trunc_f64_s(v1887)
	v1893 = v1891
	goto L418
L420:
	;
	goto L421
L421:
	;
	v1893 = int32(-2147483648)
	goto L418
L422:
	;
	v2251 = int32(-2)
	goto L7
L423:
	;
	goto L424
L424:
	;
	v1903 = base.F64_sub(v1887, base.F64_convert_i32_s(v1893))
	if base.F64_eq(v1903, float64(0)) != 0 {
		goto L416
	} else {
		goto L425
	}
L425:
	;
	v1907 = base.F64_mul(v1903, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v1907), float64(9.223372036854776e+18)) != 0 {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v1915 = base.F64_sub(v1907, base.F64_convert_i64_s(v1913))
	if base.F64_gt(v1915, float64(0.5)) != 0 {
		goto L431
	} else {
		goto L432
	}
L427:
	;
	v1911 = base.I64_trunc_f64_s(v1907)
	v1913 = v1911
	goto L426
L428:
	;
	goto L429
L429:
	;
	v1913 = int64(-9223372036854775807 - 1)
	goto L426
L430:
	;
	v1927 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v1928 = v1927 + v1926
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v1928
	if base.B2i32(v1926 < int64(0))^base.B2i32(v1928 < v1927) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L435
	}
L431:
	;
	v1926 = v1913 + int64(1)
	goto L430
L432:
	;
	goto L433
L433:
	;
	if base.F64_lt(v1915, float64(-0.5)) == int32(0) {
		v1926 = v1913
		goto L430
	} else {
		goto L434
	}
L434:
	;
	v1926 = v1913 - int64(1)
	goto L430
L435:
	;
	goto L416
L436:
	;
	v2251 = int32(-2)
	goto L7
L437:
	;
	goto L438
L438:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v1952 = base.I32_wrap_i64(v1309)
	v1953 = v1951 + v1952
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1953
	if base.B2i32(v1952 < int32(0))^base.B2i32(v1953 < v1951) != 0 {
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2251 = int32(-2)
	goto L7
L440:
	;
	goto L441
L441:
	;
	v1962 = base.F64_nearest(base.F64_mul(v1311, float64(12)))
	if base.F64_lt(base.F64_abs(v1962), float64(2.147483648e+09)) != 0 {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1970 = v1968 + v1969
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v1970
	if base.B2i32(v1968 < int32(0))^base.B2i32(v1970 < v1969) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L446
	}
L443:
	;
	v1966 = base.I32_trunc_f64_s(v1962)
	v1968 = v1966
	goto L442
L444:
	;
	goto L445
L445:
	;
	v1968 = int32(-2147483648)
	goto L442
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(4)
	v2137 = int32(0)
	v2138 = int32(25)
	v2155 = v275
	goto L23
L447:
	;
	v2251 = int32(-2)
	goto L7
L448:
	;
	goto L449
L449:
	;
	v1987 = v1309 * int64(10)
	v1991 = base.I32_wrap_i64(v1987)
	if base.I32_wrap_i64(int64(base.Ui64(v1987)>>(uint(int64(32))%64))) != v1991>>(uint(int32(31))%32) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2251 = int32(-2)
	goto L7
L451:
	;
	goto L452
L452:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v1997 = v1996 + v1991
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v1997
	if base.B2i32(v1991 < int32(0))^base.B2i32(v1997 < v1996) != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v2251 = int32(-2)
	goto L7
L454:
	;
	goto L455
L455:
	;
	v2008 = base.F64_nearest(base.F64_mul(base.F64_mul(v1311, float64(10)), float64(12)))
	if base.F64_lt(base.F64_abs(v2008), float64(2.147483648e+09)) != 0 {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v2016 = v2014 + v2015
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2016
	if base.B2i32(v2014 < int32(0))^base.B2i32(v2016 < v2015) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L460
	}
L457:
	;
	v2012 = base.I32_trunc_f64_s(v2008)
	v2014 = v2012
	goto L456
L458:
	;
	goto L459
L459:
	;
	v2014 = int32(-2147483648)
	goto L456
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(33554432)
	v2137 = int32(0)
	v2138 = int32(26)
	v2155 = v275
	goto L23
L461:
	;
	v2251 = int32(-2)
	goto L7
L462:
	;
	goto L463
L463:
	;
	v2033 = v1309 * int64(100)
	v2037 = base.I32_wrap_i64(v2033)
	if base.I32_wrap_i64(int64(base.Ui64(v2033)>>(uint(int64(32))%64))) != v2037>>(uint(int32(31))%32) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v2251 = int32(-2)
	goto L7
L465:
	;
	goto L466
L466:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v2043 = v2042 + v2037
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v2043
	if base.B2i32(v2037 < int32(0))^base.B2i32(v2043 < v2042) != 0 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v2251 = int32(-2)
	goto L7
L468:
	;
	goto L469
L469:
	;
	v2054 = base.F64_nearest(base.F64_mul(base.F64_mul(v1311, float64(100)), float64(12)))
	if base.F64_lt(base.F64_abs(v2054), float64(2.147483648e+09)) != 0 {
		goto L471
	} else {
		goto L472
	}
L470:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v2062 = v2060 + v2061
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2062
	if base.B2i32(v2060 < int32(0))^base.B2i32(v2062 < v2061) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L474
	}
L471:
	;
	v2058 = base.I32_trunc_f64_s(v2054)
	v2060 = v2058
	goto L470
L472:
	;
	goto L473
L473:
	;
	v2060 = int32(-2147483648)
	goto L470
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(67108864)
	v2137 = int32(0)
	v2138 = int32(27)
	v2155 = v275
	goto L23
L475:
	;
	v2251 = int32(-2)
	goto L7
L476:
	;
	goto L477
L477:
	;
	v2079 = v1309 * int64(1000)
	v2083 = base.I32_wrap_i64(v2079)
	if base.I32_wrap_i64(int64(base.Ui64(v2079)>>(uint(int64(32))%64))) != v2083>>(uint(int32(31))%32) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2251 = int32(-2)
	goto L7
L479:
	;
	goto L480
L480:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v2089 = v2088 + v2083
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v2089
	if base.B2i32(v2083 < int32(0))^base.B2i32(v2089 < v2088) != 0 {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v2251 = int32(-2)
	goto L7
L482:
	;
	goto L483
L483:
	;
	v2100 = base.F64_nearest(base.F64_mul(base.F64_mul(v1311, float64(1000)), float64(12)))
	if base.F64_lt(base.F64_abs(v2100), float64(2.147483648e+09)) != 0 {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v2108 = v2106 + v2107
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2108
	if base.B2i32(v2106 < int32(0))^base.B2i32(v2108 < v2107) != 0 {
		v2251 = int32(-2)
		goto L7
	} else {
		goto L488
	}
L485:
	;
	v2104 = base.I32_trunc_f64_s(v2100)
	v2106 = v2104
	goto L484
L486:
	;
	goto L487
L487:
	;
	v2106 = int32(-2147483648)
	goto L484
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+112)) = int32(134217728)
	v2137 = int32(0)
	v2138 = int32(28)
	v2155 = v275
	goto L23
L489:
	;
	v2251 = int32(-1)
	goto L7
L490:
	;
	goto L491
L491:
	;
	v2178 = v2137
	v2179 = v2138
	v2194 = v2165 | v273
	v2196 = v2155
	goto L22
L492:
	;
	goto L21
L493:
	;
	if v2196 == int32(0) {
		v2251 = v2214
		goto L7
	} else {
		goto L494
	}
L494:
	;
	v2217 = int32(-2)
	v2218 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	if v2218 == int64(-9223372036854775807-1) {
		v2251 = v2217
		goto L7
	} else {
		goto L495
	}
L495:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v2221 == int32(-2147483648) {
		v2251 = v2217
		goto L7
	} else {
		goto L496
	}
L496:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v2224 == int32(-2147483648) {
		v2251 = v2217
		goto L7
	} else {
		goto L497
	}
L497:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v2227 == int32(-2147483648) {
		v2251 = v2217
		goto L7
	} else {
		goto L498
	}
L498:
	;
	v2230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v2230 - v2227
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v2230 - v2224
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v2230 - v2221
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = int64(0) - v2218
	v2251 = v2230
	goto L7
L499:
	;
	v2322 = int32(0)
	v2323 = m.G0
	v2325 = v2323 - int32(112)
	m.G0 = v2325
	*(*int32)(unsafe.Add(mBase, uint32(v40+int32(500)))) = int32(17)
	v2332 = v40 + int32(504)
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+16)) = v2322
	v2335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2332)+8)) = v2335
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v2335
	v2339 = int32(-1)
	if v44&int32(3) == v2322 {
		v2363 = v44
		goto L505
	} else {
		goto L506
	}
L500:
	;
	v3987 = v2296
	goto L501
L501:
	;
	if v3987 != 0 {
		goto L944
	} else {
		goto L945
	}
L502:
	;
	m.G0 = v2325 + int32(112)
	v3987 = v3946
	goto L501
L503:
	;
	if base.Ui32(v2396) < base.Ui32(int32(2)) {
		v3946 = v2339
		goto L502
	} else {
		goto L520
	}
L504:
	;
	v2396 = v2388 - v44
	goto L503
L505:
	;
	v2367 = v2363
	goto L514
L506:
	;
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v2347 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2396 = int32(0)
	goto L503
L508:
	;
	goto L509
L509:
	;
	v2352 = v44
	goto L510
L510:
	;
	v2356 = v2352 + int32(1)
	if v2356&int32(3) == int32(0) {
		v2363 = v2356
		goto L505
	} else {
		goto L512
	}
L511:
	;
	v2388 = v2356
	goto L504
L512:
	;
	v2361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2356))))
	if v2361 != 0 {
		v2352 = v2356
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
L514:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2367)))
	v2376 = int32(-2139062144)
	if (int32(16843008)-v2373|v2373)&v2376 == v2376 {
		v2367 = v2367 + int32(4)
		goto L514
	} else {
		goto L516
	}
L515:
	;
	v2382 = v2367
	goto L517
L516:
	;
	goto L515
L517:
	;
	v2386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2382))))
	if v2386 != 0 {
		v2382 = v2382 + int32(1)
		goto L517
	} else {
		goto L519
	}
L518:
	;
	v2388 = v2382
	goto L504
L519:
	;
	goto L518
L520:
	;
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v2399 != int32(80) {
		v3946 = v2339
		goto L502
	} else {
		goto L521
	}
L521:
	;
	v2403 = v44 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2325)+108)) = v2403
	v2405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2403))))
	if v2405 != 0 {
		goto L526
	} else {
		goto L527
	}
L522:
	;
	v3946 = v3929
	goto L502
L523:
	;
	v3929 = int32(0)
	goto L522
L524:
	;
	if v2422&int32(1) != 0 {
		v3946 = v2339
		goto L502
	} else {
		goto L877
	}
L525:
	;
	if v2505 != 0 {
		v3929 = v2453
		goto L522
	} else {
		goto L836
	}
L526:
	;
	v2415 = v2405
	v2421 = int32(1)
	v2422 = v2322
	v2423 = v2403
	goto L529
L527:
	;
	goto L528
L528:
	;
	v3946 = int32(0)
	goto L502
L529:
	;
	if v2415&int32(255) == int32(84) {
		goto L532
	} else {
		goto L533
	}
L530:
	;
	goto L528
L531:
	;
	v3407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3405))))
	if v3407 != 0 {
		v2415 = v3407
		v2421 = v3403
		v2422 = v3404
		v2423 = v3405
		goto L529
	} else {
		goto L835
	}
L532:
	;
	v2449 = v2423 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2325)+108)) = v2449
	v2451 = int32(0)
	v3403 = v2451
	v3404 = v2451
	v3405 = v2449
	goto L531
L533:
	;
	goto L534
L534:
	;
	v2453 = int32(-1)
	if base.Ui32(int32(10)) <= base.Ui32((v2415-int32(48))&int32(255)) {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	if base.Ui32(int32(1)) < base.Ui32((v2415-int32(45))&int32(255)) {
		v3929 = v2453
		goto L522
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v2471 = F_strtod(m, v2423, v2325+int32(108))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L30
	} else {
		goto L539
	}
L538:
	;
	goto L537
L539:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+108))
	if v2473 == v2423 {
		v3929 = v2453
		goto L522
	} else {
		goto L540
	}
L540:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v2476 != 0 {
		v3929 = v2453
		goto L522
	} else {
		goto L541
	}
L541:
	;
	v2477 = base.F64_abs(v2471)
	if base.F64_gt(v2477, float64(1e+15)) != 0 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v3929 = int32(-2)
	goto L522
L543:
	;
	goto L544
L544:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v2477)) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v3929 = int32(-2)
	goto L522
L546:
	;
	goto L547
L547:
	;
	v2486 = v2473 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2325)+108)) = v2486
	if base.F64_ge(v2471, float64(0)) != 0 {
		goto L549
	} else {
		goto L550
	}
L548:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2325)+96)) = v2500
	v2503 = base.F64_sub(v2471, base.F64_convert_i64_s(v2500))
	*(*float64)(unsafe.Add(mBase, uint32(v2325)+88)) = v2503
	v2505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2473))))
	if v2421&int32(1) != 0 {
		goto L555
	} else {
		goto L556
	}
L549:
	;
	v2494 = base.F64_floor(v2471)
	goto L551
L550:
	;
	v2494 = base.F64_neg(base.F64_floor(base.F64_neg(v2471)))
	goto L551
L551:
	;
	if base.F64_lt(base.F64_abs(v2494), float64(9.223372036854776e+18)) != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v2498 = base.I64_trunc_f64_s(v2494)
	v2500 = v2498
	goto L548
L553:
	;
	goto L554
L554:
	;
	v2500 = int64(-9223372036854775807 - 1)
	goto L548
L555:
	;
	switch v2505 - int32(45) {
	case 0:
		goto L558
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 40, 41, 43:
		v3929 = v2453
		goto L522
	case 23:
		goto L560
	case 32:
		goto L562
	case 39:
		goto L564
	case 42:
		goto L561
	case 44:
		goto L563
	default:
		goto L565
	}
L556:
	;
	goto L557
L557:
	;
	switch v2505 - int32(58) {
	case 0:
		goto L524
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 16, 17, 18, 20, 21, 22, 23, 24:
		v3929 = v2453
		goto L522
	case 14:
		goto L780
	case 19:
		goto L779
	case 25:
		goto L778
	default:
		goto L525
	}
L558:
	;
	if v2422&int32(1) != 0 {
		v3929 = v2453
		goto L522
	} else {
		goto L704
	}
L559:
	;
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+16))
	v2846 = base.I32_wrap_i64(v2604)
	v2847 = v2845 + v2846
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+16)) = v2847
	if base.B2i32(v2846 < int32(0))^base.B2i32(v2847 < v2845) != 0 {
		goto L681
	} else {
		goto L682
	}
L560:
	;
	if base.Ui64(v2500-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L662
	} else {
		goto L663
	}
L561:
	;
	if base.Ui64(v2500-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L630
	} else {
		goto L631
	}
L562:
	;
	if base.Ui64(v2500-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L601
	} else {
		goto L602
	}
L563:
	;
	if base.Ui64(v2500-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L590
	} else {
		goto L591
	}
L564:
	;
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	v2513 = v2423 + base.B2i32(v2510 == int32(45))
	v2514 = int32(515387)
	v2518 = m.G0
	v2520 = v2518 - int32(32)
	v2521 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2520)+24)) = v2521
	*(*int64)(unsafe.Add(mBase, uint32(v2520)+16)) = v2521
	*(*int64)(unsafe.Add(mBase, uint32(v2520)+8)) = v2521
	*(*int64)(unsafe.Add(mBase, uint32(v2520))) = v2521
	v2529 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v2529 == int32(0) {
		goto L568
	} else {
		goto L569
	}
L565:
	;
	if v2505 != 0 {
		v3929 = v2453
		goto L522
	} else {
		goto L566
	}
L566:
	;
	goto L564
L567:
	;
	if (v2422|base.B2i32(v2597 != int32(8)))&int32(1) != 0 {
		goto L558
	} else {
		goto L588
	}
L568:
	;
	v2597 = int32(0)
	goto L567
L569:
	;
	goto L570
L570:
	;
	v2533 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v2533 == int32(0) {
		goto L571
	} else {
		goto L572
	}
L571:
	;
	v2537 = v2513
	goto L574
L572:
	;
	goto L573
L573:
	;
	v2547 = v2514
	v2548 = v2529
	goto L577
L574:
	;
	v2543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2537))))
	if v2543 == v2529 {
		v2537 = v2537 + int32(1)
		goto L574
	} else {
		goto L576
	}
L575:
	;
	v2597 = v2537 - v2513
	goto L567
L576:
	;
	goto L575
L577:
	;
	v2555 = v2520 + int32(base.Ui32(v2548)>>(uint(int32(3))%32))&int32(28)
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2555)))
	v2557 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2555))) = v2556 | v2557<<(uint(v2548)%32)
	v2561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+1)))
	if v2561 != 0 {
		v2547 = v2547 + v2557
		v2548 = v2561
		goto L577
	} else {
		goto L579
	}
L578:
	;
	v2564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2513))))
	if v2564 == int32(0) {
		v2589 = v2513
		goto L580
	} else {
		goto L581
	}
L579:
	;
	goto L578
L580:
	;
	v2597 = v2589 - v2513
	goto L567
L581:
	;
	v2568 = v2513
	v2569 = v2564
	goto L582
L582:
	;
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(v2520+int32(base.Ui32(v2569)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v2577)>>(uint(v2569)%32))&int32(1) == int32(0) {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v2589 = v2585
	goto L580
L584:
	;
	v2589 = v2568
	goto L580
L585:
	;
	goto L586
L586:
	;
	v2583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2568)+1)))
	v2585 = v2568 + int32(1)
	if v2583 != 0 {
		v2568 = v2585
		v2569 = v2583
		goto L582
	} else {
		goto L587
	}
L587:
	;
	goto L583
L588:
	;
	v2604 = base.I64_div_s(v2500, int64(10000))
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v2604-int64(2147483648)) {
		goto L559
	} else {
		goto L589
	}
L589:
	;
	v3929 = int32(-2)
	goto L522
L590:
	;
	v3929 = int32(-2)
	goto L522
L591:
	;
	goto L592
L592:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+16))
	v2616 = base.I32_wrap_i64(v2500)
	v2617 = v2615 + v2616
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+16)) = v2617
	if base.B2i32(v2616 < int32(0))^base.B2i32(v2617 < v2615) != 0 {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v3929 = int32(-2)
	goto L522
L594:
	;
	goto L595
L595:
	;
	v2626 = base.F64_nearest(base.F64_mul(v2503, float64(12)))
	if base.F64_lt(base.F64_abs(v2626), float64(2.147483648e+09)) != 0 {
		goto L597
	} else {
		goto L598
	}
L596:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	v2634 = v2632 + v2633
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+12)) = v2634
	v2636 = int32(1)
	v2639 = int32(0)
	if base.B2i32(v2632 < v2639)^base.B2i32(v2634 < v2633) == v2639 {
		v3403 = v2636
		v3404 = v2636
		v3405 = v2486
		goto L531
	} else {
		goto L600
	}
L597:
	;
	v2630 = base.I32_trunc_f64_s(v2626)
	v2632 = v2630
	goto L596
L598:
	;
	goto L599
L599:
	;
	v2632 = int32(-2147483648)
	goto L596
L600:
	;
	v3929 = int32(-2)
	goto L522
L601:
	;
	v3929 = int32(-2)
	goto L522
L602:
	;
	goto L603
L603:
	;
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	v2651 = base.I32_wrap_i64(v2500)
	v2652 = v2650 + v2651
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+12)) = v2652
	if base.B2i32(v2651 < int32(0))^base.B2i32(v2652 < v2650) != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3929 = int32(-2)
	goto L522
L605:
	;
	goto L606
L606:
	;
	v2659 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		goto L607
	} else {
		goto L608
	}
L607:
	;
	v3403 = int32(1)
	v3404 = v2659
	v3405 = v2486
	goto L531
L608:
	;
	goto L609
L609:
	;
	v2664 = base.F64_mul(v2503, float64(30))
	if base.F64_lt(base.F64_abs(v2664), float64(2.147483648e+09)) != 0 {
		goto L611
	} else {
		goto L612
	}
L610:
	;
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v2672 = v2670 + v2671
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2672
	if base.B2i32(v2670 < int32(0))^base.B2i32(v2672 < v2671) != 0 {
		goto L614
	} else {
		goto L615
	}
L611:
	;
	v2668 = base.I32_trunc_f64_s(v2664)
	v2670 = v2668
	goto L610
L612:
	;
	goto L613
L613:
	;
	v2670 = int32(-2147483648)
	goto L610
L614:
	;
	v3929 = int32(-2)
	goto L522
L615:
	;
	goto L616
L616:
	;
	v2680 = base.F64_sub(v2664, base.F64_convert_i32_s(v2670))
	if base.F64_eq(v2680, float64(0)) != 0 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v3403 = int32(1)
	v3404 = v2659
	v3405 = v2486
	goto L531
L618:
	;
	goto L619
L619:
	;
	v2685 = base.F64_mul(v2680, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2685), float64(9.223372036854776e+18)) != 0 {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	v2693 = base.F64_sub(v2685, base.F64_convert_i64_s(v2691))
	if base.F64_gt(v2693, float64(0.5)) != 0 {
		goto L625
	} else {
		goto L626
	}
L621:
	;
	v2689 = base.I64_trunc_f64_s(v2685)
	v2691 = v2689
	goto L620
L622:
	;
	goto L623
L623:
	;
	v2691 = int64(-9223372036854775807 - 1)
	goto L620
L624:
	;
	v2705 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v2706 = v2705 + v2704
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v2706
	if base.B2i32(v2704 < int64(0))^base.B2i32(v2706 < v2705) == int32(0) {
		v3403 = int32(1)
		v3404 = v2659
		v3405 = v2486
		goto L531
	} else {
		goto L629
	}
L625:
	;
	v2704 = v2691 + int64(1)
	goto L624
L626:
	;
	goto L627
L627:
	;
	if base.F64_lt(v2693, float64(-0.5)) == int32(0) {
		v2704 = v2691
		goto L624
	} else {
		goto L628
	}
L628:
	;
	v2704 = v2691 - int64(1)
	goto L624
L629:
	;
	v3929 = int32(-2)
	goto L522
L630:
	;
	v3929 = int32(-2)
	goto L522
L631:
	;
	goto L632
L632:
	;
	v2722 = v2500 * int64(7)
	v2726 = base.I32_wrap_i64(v2722)
	if base.I32_wrap_i64(int64(base.Ui64(v2722)>>(uint(int64(32))%64))) != v2726>>(uint(int32(31))%32) {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3929 = int32(-2)
	goto L522
L634:
	;
	goto L635
L635:
	;
	v2731 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v2732 = v2731 + v2726
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2732
	if base.B2i32(v2726 < int32(0))^base.B2i32(v2732 < v2731) != 0 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v3929 = int32(-2)
	goto L522
L637:
	;
	goto L638
L638:
	;
	v2739 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	v3403 = int32(1)
	v3404 = v2739
	v3405 = v2486
	goto L531
L640:
	;
	goto L641
L641:
	;
	v2744 = base.F64_mul(v2503, float64(7))
	if base.F64_lt(base.F64_abs(v2744), float64(2.147483648e+09)) != 0 {
		goto L643
	} else {
		goto L644
	}
L642:
	;
	v2751 = v2732 + v2750
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2751
	if base.B2i32(v2750 < int32(0))^base.B2i32(v2751 < v2732) != 0 {
		goto L646
	} else {
		goto L647
	}
L643:
	;
	v2748 = base.I32_trunc_f64_s(v2744)
	v2750 = v2748
	goto L642
L644:
	;
	goto L645
L645:
	;
	v2750 = int32(-2147483648)
	goto L642
L646:
	;
	v3929 = int32(-2)
	goto L522
L647:
	;
	goto L648
L648:
	;
	v2759 = base.F64_sub(v2744, base.F64_convert_i32_s(v2750))
	if base.F64_eq(v2759, float64(0)) != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v3403 = int32(1)
	v3404 = v2739
	v3405 = v2486
	goto L531
L650:
	;
	goto L651
L651:
	;
	v2764 = base.F64_mul(v2759, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2764), float64(9.223372036854776e+18)) != 0 {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	v2772 = base.F64_sub(v2764, base.F64_convert_i64_s(v2770))
	if base.F64_gt(v2772, float64(0.5)) != 0 {
		goto L657
	} else {
		goto L658
	}
L653:
	;
	v2768 = base.I64_trunc_f64_s(v2764)
	v2770 = v2768
	goto L652
L654:
	;
	goto L655
L655:
	;
	v2770 = int64(-9223372036854775807 - 1)
	goto L652
L656:
	;
	v2784 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v2785 = v2784 + v2783
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v2785
	if base.B2i32(v2783 < int64(0))^base.B2i32(v2785 < v2784) == int32(0) {
		v3403 = int32(1)
		v3404 = v2739
		v3405 = v2486
		goto L531
	} else {
		goto L661
	}
L657:
	;
	v2783 = v2770 + int64(1)
	goto L656
L658:
	;
	goto L659
L659:
	;
	if base.F64_lt(v2772, float64(-0.5)) == int32(0) {
		v2783 = v2770
		goto L656
	} else {
		goto L660
	}
L660:
	;
	v2783 = v2770 - int64(1)
	goto L656
L661:
	;
	v3929 = int32(-2)
	goto L522
L662:
	;
	v3929 = int32(-2)
	goto L522
L663:
	;
	goto L664
L664:
	;
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v2801 = base.I32_wrap_i64(v2500)
	v2802 = v2800 + v2801
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2802
	if base.B2i32(v2801 < int32(0))^base.B2i32(v2802 < v2800) != 0 {
		goto L665
	} else {
		goto L666
	}
L665:
	;
	v3929 = int32(-2)
	goto L522
L666:
	;
	goto L667
L667:
	;
	v2809 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v3403 = int32(1)
	v3404 = v2809
	v3405 = v2486
	goto L531
L669:
	;
	goto L670
L670:
	;
	v2814 = base.F64_mul(v2503, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2814), float64(9.223372036854776e+18)) != 0 {
		goto L672
	} else {
		goto L673
	}
L671:
	;
	v2822 = base.F64_sub(v2814, base.F64_convert_i64_s(v2820))
	if base.F64_gt(v2822, float64(0.5)) != 0 {
		goto L676
	} else {
		goto L677
	}
L672:
	;
	v2818 = base.I64_trunc_f64_s(v2814)
	v2820 = v2818
	goto L671
L673:
	;
	goto L674
L674:
	;
	v2820 = int64(-9223372036854775807 - 1)
	goto L671
L675:
	;
	v2834 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v2835 = v2834 + v2833
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v2835
	if base.B2i32(v2833 < int64(0))^base.B2i32(v2835 < v2834) == int32(0) {
		v3403 = int32(1)
		v3404 = v2809
		v3405 = v2486
		goto L531
	} else {
		goto L680
	}
L676:
	;
	v2833 = v2820 + int64(1)
	goto L675
L677:
	;
	goto L678
L678:
	;
	if base.F64_lt(v2822, float64(-0.5)) == int32(0) {
		v2833 = v2820
		goto L675
	} else {
		goto L679
	}
L679:
	;
	v2833 = v2820 - int64(1)
	goto L675
L680:
	;
	v3929 = int32(-2)
	goto L522
L681:
	;
	v3929 = int32(-2)
	goto L522
L682:
	;
	goto L683
L683:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	v2855 = int64(100)
	v2856 = base.I64_div_s(v2500, v2855)
	v2858 = base.I64_rem_s(v2856, v2855)
	v2859 = base.I32_wrap_i64(v2858)
	v2860 = v2854 + v2859
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+12)) = v2860
	if base.B2i32(v2859 < int32(0))^base.B2i32(v2860 < v2854) != 0 {
		goto L684
	} else {
		goto L685
	}
L684:
	;
	v3929 = int32(-2)
	goto L522
L685:
	;
	goto L686
L686:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v2869 = base.I64_rem_s(v2500, int64(100))
	v2870 = base.I32_wrap_i64(v2869)
	v2871 = v2867 + v2870
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2871
	if base.B2i32(v2870 < int32(0))^base.B2i32(v2871 < v2867) != 0 {
		goto L687
	} else {
		goto L688
	}
L687:
	;
	v3929 = int32(-2)
	goto L522
L688:
	;
	goto L689
L689:
	;
	if base.F64_ne(v2503, float64(0)) != 0 {
		goto L690
	} else {
		goto L691
	}
L690:
	;
	v2881 = base.F64_mul(v2503, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v2881), float64(9.223372036854776e+18)) != 0 {
		goto L694
	} else {
		goto L695
	}
L691:
	;
	goto L692
L692:
	;
	v2914 = int32(0)
	if v2505 == v2914 {
		goto L523
	} else {
		goto L703
	}
L693:
	;
	v2889 = base.F64_sub(v2881, base.F64_convert_i64_s(v2887))
	if base.F64_gt(v2889, float64(0.5)) != 0 {
		goto L698
	} else {
		goto L699
	}
L694:
	;
	v2885 = base.I64_trunc_f64_s(v2881)
	v2887 = v2885
	goto L693
L695:
	;
	goto L696
L696:
	;
	v2887 = int64(-9223372036854775807 - 1)
	goto L693
L697:
	;
	v2901 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v2902 = v2901 + v2900
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v2902
	if base.B2i32(v2900 < int64(0))^base.B2i32(v2902 < v2901) != 0 {
		v3929 = int32(-2)
		goto L522
	} else {
		goto L702
	}
L698:
	;
	v2900 = v2887 + int64(1)
	goto L697
L699:
	;
	goto L700
L700:
	;
	if base.F64_lt(v2889, float64(-0.5)) == int32(0) {
		v2900 = v2887
		goto L697
	} else {
		goto L701
	}
L701:
	;
	v2900 = v2887 - int64(1)
	goto L697
L702:
	;
	goto L692
L703:
	;
	v3403 = int32(0)
	v3404 = v2914
	v3405 = v2486
	goto L531
L704:
	;
	if base.Ui64(v2500-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L705
	} else {
		goto L706
	}
L705:
	;
	v3929 = int32(-2)
	goto L522
L706:
	;
	goto L707
L707:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+16))
	v2926 = base.I32_wrap_i64(v2500)
	v2927 = v2925 + v2926
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+16)) = v2927
	if base.B2i32(v2926 < int32(0))^base.B2i32(v2927 < v2925) != 0 {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	v3929 = int32(-2)
	goto L522
L709:
	;
	goto L710
L710:
	;
	v2936 = base.F64_nearest(base.F64_mul(v2503, float64(12)))
	if base.F64_lt(base.F64_abs(v2936), float64(2.147483648e+09)) != 0 {
		goto L712
	} else {
		goto L713
	}
L711:
	;
	v2943 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	v2944 = v2942 + v2943
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+12)) = v2944
	if base.B2i32(v2942 < int32(0))^base.B2i32(v2944 < v2943) != 0 {
		goto L715
	} else {
		goto L716
	}
L712:
	;
	v2940 = base.I32_trunc_f64_s(v2936)
	v2942 = v2940
	goto L711
L713:
	;
	goto L714
L714:
	;
	v2942 = int32(-2147483648)
	goto L711
L715:
	;
	v3929 = int32(-2)
	goto L522
L716:
	;
	goto L717
L717:
	;
	v2951 = int32(0)
	if v2505 == int32(84) {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	v3403 = int32(0)
	v3404 = v2951
	v3405 = v2486
	goto L531
L719:
	;
	goto L720
L720:
	;
	if v2505 == int32(0) {
		v3929 = v2505
		goto L522
	} else {
		goto L721
	}
L721:
	;
	v2963 = F_ParseISO8601Number(m, v2486, v2325+int32(108), v2325+int32(96), v2325+int32(88))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L30
	} else {
		goto L722
	}
L722:
	;
	if v2963 != 0 {
		v3929 = v2963
		goto L522
	} else {
		goto L723
	}
L723:
	;
	v2965 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+96))
	if base.Ui64(v2965-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L724
	} else {
		goto L725
	}
L724:
	;
	v3929 = int32(-2)
	goto L522
L725:
	;
	goto L726
L726:
	;
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+12))
	v2972 = base.I32_wrap_i64(v2965)
	v2973 = v2971 + v2972
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+12)) = v2973
	if base.B2i32(v2972 < int32(0))^base.B2i32(v2973 < v2971) != 0 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	v3929 = int32(-2)
	goto L522
L728:
	;
	goto L729
L729:
	;
	v2981 = *(*float64)(unsafe.Add(mBase, uint32(v2325)+88))
	if base.F64_eq(v2981, float64(0)) != 0 {
		v3042 = int32(1)
		goto L730
	} else {
		goto L731
	}
L730:
	;
	if v3042 == int32(0) {
		goto L747
	} else {
		goto L748
	}
L731:
	;
	v2985 = base.F64_mul(v2981, float64(30))
	if base.F64_lt(base.F64_abs(v2985), float64(2.147483648e+09)) != 0 {
		goto L733
	} else {
		goto L734
	}
L732:
	;
	v2992 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v2993 = v2991 + v2992
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v2993
	v2995 = int32(0)
	if base.B2i32(v2991 < v2995)^base.B2i32(v2993 < v2992) != 0 {
		v3042 = v2995
		goto L730
	} else {
		goto L736
	}
L733:
	;
	v2989 = base.I32_trunc_f64_s(v2985)
	v2991 = v2989
	goto L732
L734:
	;
	goto L735
L735:
	;
	v2991 = int32(-2147483648)
	goto L732
L736:
	;
	v3002 = base.F64_sub(v2985, base.F64_convert_i32_s(v2991))
	if base.F64_eq(v3002, float64(0)) != 0 {
		v3042 = int32(1)
		goto L730
	} else {
		goto L737
	}
L737:
	;
	v3006 = base.F64_mul(v3002, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v3006), float64(9.223372036854776e+18)) != 0 {
		goto L739
	} else {
		goto L740
	}
L738:
	;
	v3014 = base.F64_sub(v3006, base.F64_convert_i64_s(v3012))
	if base.F64_gt(v3014, float64(0.5)) != 0 {
		goto L743
	} else {
		goto L744
	}
L739:
	;
	v3010 = base.I64_trunc_f64_s(v3006)
	v3012 = v3010
	goto L738
L740:
	;
	goto L741
L741:
	;
	v3012 = int64(-9223372036854775807 - 1)
	goto L738
L742:
	;
	v3026 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3027 = v3026 + v3025
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3027
	v3042 = base.B2i32(base.B2i32(v3025 < int64(0))^base.B2i32(v3027 < v3026) == int32(0))
	goto L730
L743:
	;
	v3025 = v3012 + int64(1)
	goto L742
L744:
	;
	goto L745
L745:
	;
	if base.F64_lt(v3014, float64(-0.5)) == int32(0) {
		v3025 = v3012
		goto L742
	} else {
		goto L746
	}
L746:
	;
	v3025 = v3012 - int64(1)
	goto L742
L747:
	;
	v3929 = int32(-2)
	goto L522
L748:
	;
	goto L749
L749:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+108))
	v3047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3046))))
	if v3047 != int32(45) {
		goto L750
	} else {
		goto L751
	}
L750:
	;
	if v3047 == int32(84) {
		v3403 = int32(0)
		v3404 = v2951
		v3405 = v3046
		goto L531
	} else {
		goto L753
	}
L751:
	;
	goto L752
L752:
	;
	v3056 = v3046 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2325)+108)) = v3056
	v3064 = F_ParseISO8601Number(m, v3056, v2325+int32(108), v2325+int32(96), v2325+int32(88))
	mBase = m.M
	v3065 = m.ExcPending
	if v3065 != 0 {
		goto L30
	} else {
		goto L755
	}
L753:
	;
	if v3047 == int32(0) {
		v3929 = v3047
		goto L522
	} else {
		goto L754
	}
L754:
	;
	v3946 = v2339
	goto L502
L755:
	;
	if v3064 != 0 {
		v3929 = v3064
		goto L522
	} else {
		goto L756
	}
L756:
	;
	v3066 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+96))
	if base.Ui64(v3066-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v3929 = int32(-2)
	goto L522
L758:
	;
	goto L759
L759:
	;
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v2332)+8))
	v3073 = base.I32_wrap_i64(v3066)
	v3074 = v3072 + v3073
	*(*int32)(unsafe.Add(mBase, uint32(v2332)+8)) = v3074
	if base.B2i32(v3073 < int32(0))^base.B2i32(v3074 < v3072) != 0 {
		goto L760
	} else {
		goto L761
	}
L760:
	;
	v3929 = int32(-2)
	goto L522
L761:
	;
	goto L762
L762:
	;
	v3081 = *(*float64)(unsafe.Add(mBase, uint32(v2325)+88))
	if base.F64_ne(v3081, float64(0)) != 0 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	v3085 = base.F64_mul(v3081, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v3085), float64(9.223372036854776e+18)) != 0 {
		goto L767
	} else {
		goto L768
	}
L764:
	;
	goto L765
L765:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+108))
	v3120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3119))))
	if v3120 == int32(84) {
		v3403 = int32(0)
		v3404 = v2951
		v3405 = v3119
		goto L531
	} else {
		goto L776
	}
L766:
	;
	v3093 = base.F64_sub(v3085, base.F64_convert_i64_s(v3091))
	if base.F64_gt(v3093, float64(0.5)) != 0 {
		goto L771
	} else {
		goto L772
	}
L767:
	;
	v3089 = base.I64_trunc_f64_s(v3085)
	v3091 = v3089
	goto L766
L768:
	;
	goto L769
L769:
	;
	v3091 = int64(-9223372036854775807 - 1)
	goto L766
L770:
	;
	v3105 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3106 = v3105 + v3104
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3106
	if base.B2i32(v3104 < int64(0))^base.B2i32(v3106 < v3105) != 0 {
		v3929 = int32(-2)
		goto L522
	} else {
		goto L775
	}
L771:
	;
	v3104 = v3091 + int64(1)
	goto L770
L772:
	;
	goto L773
L773:
	;
	if base.F64_lt(v3093, float64(-0.5)) == int32(0) {
		v3104 = v3091
		goto L770
	} else {
		goto L774
	}
L774:
	;
	v3104 = v3091 - int64(1)
	goto L770
L775:
	;
	goto L765
L776:
	;
	if v3120 == int32(0) {
		v3929 = v3120
		goto L522
	} else {
		goto L777
	}
L777:
	;
	v3946 = v2339
	goto L502
L778:
	;
	v3308 = v2325 + int32(40)
	v3311 = int64(1000000)
	v3312 = int64(0)
	v3317 = int64(32)
	v3320 = int64(base.Ui64(v2500) >> (uint(v3317) % 64))
	v3323 = int64(4294967295)
	v3326 = v2500 & v3323
	v3327 = v3311 * v3326
	v3331 = int64(base.Ui64(v3327)>>(uint(v3317)%64)) + v3311*v3320
	v3338 = v3326*v3312 + v3331&v3323
	*(*int64)(unsafe.Add(mBase, uint32(v3308)+8)) = v2500*v3312 + v2500>>(uint(int64(63))%64)*v3311 + v3312*v3320 + int64(base.Ui64(v3331)>>(uint(v3317)%64)) + int64(base.Ui64(v3338)>>(uint(v3317)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3308))) = v3327&v3323 | v3338<<(uint(v3317)%64)
	goto L817
L779:
	;
	v3218 = v2325 + int32(24)
	v3221 = int64(60000000)
	v3222 = int64(0)
	v3227 = int64(32)
	v3230 = int64(base.Ui64(v2500) >> (uint(v3227) % 64))
	v3233 = int64(4294967295)
	v3236 = v2500 & v3233
	v3237 = v3221 * v3236
	v3241 = int64(base.Ui64(v3237)>>(uint(v3227)%64)) + v3221*v3230
	v3248 = v3236*v3222 + v3241&v3233
	*(*int64)(unsafe.Add(mBase, uint32(v3218)+8)) = v2500*v3222 + v2500>>(uint(int64(63))%64)*v3221 + v3222*v3230 + int64(base.Ui64(v3241)>>(uint(v3227)%64)) + int64(base.Ui64(v3248)>>(uint(v3227)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3218))) = v3237&v3233 | v3248<<(uint(v3227)%64)
	goto L799
L780:
	;
	v3128 = v2325 + int32(8)
	v3131 = int64(3600000000)
	v3132 = int64(0)
	v3137 = int64(32)
	v3140 = int64(base.Ui64(v2500) >> (uint(v3137) % 64))
	v3143 = int64(4294967295)
	v3146 = v2500 & v3143
	v3147 = v3131 * v3146
	v3151 = int64(base.Ui64(v3147)>>(uint(v3137)%64)) + v3131*v3140
	v3158 = v3146*v3132 + v3151&v3143
	*(*int64)(unsafe.Add(mBase, uint32(v3128)+8)) = v2500*v3132 + v2500>>(uint(int64(63))%64)*v3131 + v3132*v3140 + int64(base.Ui64(v3151)>>(uint(v3137)%64)) + int64(base.Ui64(v3158)>>(uint(v3137)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3128))) = v3147&v3143 | v3158<<(uint(v3137)%64)
	goto L781
L781:
	;
	v3169 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+16))
	v3170 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+8))
	if v3169 != v3170>>(uint(int64(63))%64) {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v3929 = int32(-2)
	goto L522
L783:
	;
	goto L784
L784:
	;
	v3175 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3176 = v3175 + v3170
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3176
	if base.B2i32(v3170 < int64(0))^base.B2i32(v3176 < v3175) != 0 {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v3929 = int32(-2)
	goto L522
L786:
	;
	goto L787
L787:
	;
	v3183 = int32(0)
	v3184 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		v3403 = v3183
		v3404 = v3184
		v3405 = v2486
		goto L531
	} else {
		goto L788
	}
L788:
	;
	v3188 = base.F64_mul(v2503, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v3188), float64(9.223372036854776e+18)) != 0 {
		goto L790
	} else {
		goto L791
	}
L789:
	;
	v3196 = base.F64_sub(v3188, base.F64_convert_i64_s(v3194))
	if base.F64_gt(v3196, float64(0.5)) != 0 {
		goto L794
	} else {
		goto L795
	}
L790:
	;
	v3192 = base.I64_trunc_f64_s(v3188)
	v3194 = v3192
	goto L789
L791:
	;
	goto L792
L792:
	;
	v3194 = int64(-9223372036854775807 - 1)
	goto L789
L793:
	;
	v3208 = v3207 + v3176
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3208
	if base.B2i32(v3207 < int64(0))^base.B2i32(v3208 < v3176) == int32(0) {
		v3403 = v3183
		v3404 = v3184
		v3405 = v2486
		goto L531
	} else {
		goto L798
	}
L794:
	;
	v3207 = v3194 + int64(1)
	goto L793
L795:
	;
	goto L796
L796:
	;
	if base.F64_lt(v3196, float64(-0.5)) == int32(0) {
		v3207 = v3194
		goto L793
	} else {
		goto L797
	}
L797:
	;
	v3207 = v3194 - int64(1)
	goto L793
L798:
	;
	v3929 = int32(-2)
	goto L522
L799:
	;
	v3259 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+32))
	v3260 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+24))
	if v3259 != v3260>>(uint(int64(63))%64) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	v3929 = int32(-2)
	goto L522
L801:
	;
	goto L802
L802:
	;
	v3265 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3266 = v3265 + v3260
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3266
	if base.B2i32(v3260 < int64(0))^base.B2i32(v3266 < v3265) != 0 {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v3929 = int32(-2)
	goto L522
L804:
	;
	goto L805
L805:
	;
	v3273 = int32(0)
	v3274 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		v3403 = v3273
		v3404 = v3274
		v3405 = v2486
		goto L531
	} else {
		goto L806
	}
L806:
	;
	v3278 = base.F64_mul(v2503, float64(6e+07))
	if base.F64_lt(base.F64_abs(v3278), float64(9.223372036854776e+18)) != 0 {
		goto L808
	} else {
		goto L809
	}
L807:
	;
	v3286 = base.F64_sub(v3278, base.F64_convert_i64_s(v3284))
	if base.F64_gt(v3286, float64(0.5)) != 0 {
		goto L812
	} else {
		goto L813
	}
L808:
	;
	v3282 = base.I64_trunc_f64_s(v3278)
	v3284 = v3282
	goto L807
L809:
	;
	goto L810
L810:
	;
	v3284 = int64(-9223372036854775807 - 1)
	goto L807
L811:
	;
	v3298 = v3297 + v3266
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3298
	if base.B2i32(v3297 < int64(0))^base.B2i32(v3298 < v3266) == int32(0) {
		v3403 = v3273
		v3404 = v3274
		v3405 = v2486
		goto L531
	} else {
		goto L816
	}
L812:
	;
	v3297 = v3284 + int64(1)
	goto L811
L813:
	;
	goto L814
L814:
	;
	if base.F64_lt(v3286, float64(-0.5)) == int32(0) {
		v3297 = v3284
		goto L811
	} else {
		goto L815
	}
L815:
	;
	v3297 = v3284 - int64(1)
	goto L811
L816:
	;
	v3929 = int32(-2)
	goto L522
L817:
	;
	v3349 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+48))
	v3350 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+40))
	if v3349 != v3350>>(uint(int64(63))%64) {
		goto L818
	} else {
		goto L819
	}
L818:
	;
	v3929 = int32(-2)
	goto L522
L819:
	;
	goto L820
L820:
	;
	v3355 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3356 = v3355 + v3350
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3356
	if base.B2i32(v3350 < int64(0))^base.B2i32(v3356 < v3355) != 0 {
		goto L821
	} else {
		goto L822
	}
L821:
	;
	v3929 = int32(-2)
	goto L522
L822:
	;
	goto L823
L823:
	;
	v3363 = int32(0)
	v3364 = int32(1)
	if base.F64_eq(v2503, float64(0)) != 0 {
		v3403 = v3363
		v3404 = v3364
		v3405 = v2486
		goto L531
	} else {
		goto L824
	}
L824:
	;
	v3368 = base.F64_mul(v2503, float64(1e+06))
	if base.F64_lt(base.F64_abs(v3368), float64(9.223372036854776e+18)) != 0 {
		goto L826
	} else {
		goto L827
	}
L825:
	;
	v3376 = base.F64_sub(v3368, base.F64_convert_i64_s(v3374))
	if base.F64_gt(v3376, float64(0.5)) != 0 {
		goto L830
	} else {
		goto L831
	}
L826:
	;
	v3372 = base.I64_trunc_f64_s(v3368)
	v3374 = v3372
	goto L825
L827:
	;
	goto L828
L828:
	;
	v3374 = int64(-9223372036854775807 - 1)
	goto L825
L829:
	;
	v3388 = v3387 + v3356
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3388
	if base.B2i32(v3387 < int64(0))^base.B2i32(v3388 < v3356) != 0 {
		v3929 = int32(-2)
		goto L522
	} else {
		goto L834
	}
L830:
	;
	v3387 = v3374 + int64(1)
	goto L829
L831:
	;
	goto L832
L832:
	;
	if base.F64_lt(v3376, float64(-0.5)) == int32(0) {
		v3387 = v3374
		goto L829
	} else {
		goto L833
	}
L833:
	;
	v3387 = v3374 - int64(1)
	goto L829
L834:
	;
	v3403 = v3363
	v3404 = v3364
	v3405 = v2486
	goto L531
L835:
	;
	goto L530
L836:
	;
	v3446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2423))))
	v3449 = v2423 + base.B2i32(v3446 == int32(45))
	v3450 = int32(515387)
	v3454 = m.G0
	v3456 = v3454 - int32(32)
	v3457 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3456)+24)) = v3457
	*(*int64)(unsafe.Add(mBase, uint32(v3456)+16)) = v3457
	*(*int64)(unsafe.Add(mBase, uint32(v3456)+8)) = v3457
	*(*int64)(unsafe.Add(mBase, uint32(v3456))) = v3457
	v3465 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v3465 == int32(0) {
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if (v2422|base.B2i32(v3533 != int32(6)))&int32(1) != 0 {
		goto L524
	} else {
		goto L858
	}
L838:
	;
	v3533 = int32(0)
	goto L837
L839:
	;
	goto L840
L840:
	;
	v3469 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v3469 == int32(0) {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	v3473 = v3449
	goto L844
L842:
	;
	goto L843
L843:
	;
	v3483 = v3450
	v3484 = v3465
	goto L847
L844:
	;
	v3479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3473))))
	if v3479 == v3465 {
		v3473 = v3473 + int32(1)
		goto L844
	} else {
		goto L846
	}
L845:
	;
	v3533 = v3473 - v3449
	goto L837
L846:
	;
	goto L845
L847:
	;
	v3491 = v3456 + int32(base.Ui32(v3484)>>(uint(int32(3))%32))&int32(28)
	v3492 = *(*int32)(unsafe.Add(mBase, uint32(v3491)))
	v3493 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3491))) = v3492 | v3493<<(uint(v3484)%32)
	v3497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483)+1)))
	if v3497 != 0 {
		v3483 = v3483 + v3493
		v3484 = v3497
		goto L847
	} else {
		goto L849
	}
L848:
	;
	v3500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3449))))
	if v3500 == int32(0) {
		v3525 = v3449
		goto L850
	} else {
		goto L851
	}
L849:
	;
	goto L848
L850:
	;
	v3533 = v3525 - v3449
	goto L837
L851:
	;
	v3504 = v3449
	v3505 = v3500
	goto L852
L852:
	;
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v3456+int32(base.Ui32(v3505)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v3513)>>(uint(v3505)%32))&int32(1) == int32(0) {
		goto L854
	} else {
		goto L855
	}
L853:
	;
	v3525 = v3521
	goto L850
L854:
	;
	v3525 = v3504
	goto L850
L855:
	;
	goto L856
L856:
	;
	v3519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3504)+1)))
	v3521 = v3504 + int32(1)
	if v3519 != 0 {
		v3504 = v3521
		v3505 = v3519
		goto L852
	} else {
		goto L857
	}
L857:
	;
	goto L853
L858:
	;
	v3540 = v2325 + int32(72)
	v3542 = base.I64_div_s(v2500, int64(10000))
	v3545 = int64(3600000000)
	v3546 = int64(0)
	v3551 = int64(32)
	v3554 = int64(base.Ui64(v3542) >> (uint(v3551) % 64))
	v3557 = int64(4294967295)
	v3560 = v3542 & v3557
	v3561 = v3545 * v3560
	v3565 = int64(base.Ui64(v3561)>>(uint(v3551)%64)) + v3545*v3554
	v3572 = v3560*v3546 + v3565&v3557
	*(*int64)(unsafe.Add(mBase, uint32(v3540)+8)) = v3542*v3546 + v3542>>(uint(int64(63))%64)*v3545 + v3546*v3554 + int64(base.Ui64(v3565)>>(uint(v3551)%64)) + int64(base.Ui64(v3572)>>(uint(v3551)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3540))) = v3561&v3557 | v3572<<(uint(v3551)%64)
	goto L859
L859:
	;
	v3583 = int32(-2)
	v3584 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+80))
	v3585 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+72))
	if v3584 != v3585>>(uint(int64(63))%64) {
		v3946 = v3583
		goto L502
	} else {
		goto L860
	}
L860:
	;
	v3589 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3590 = v3589 + v3585
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3590
	if base.B2i32(v3585 < int64(0))^base.B2i32(v3590 < v3589) != 0 {
		v3946 = v3583
		goto L502
	} else {
		goto L861
	}
L861:
	;
	v3596 = int64(100)
	v3597 = base.I64_div_s(v2500, v3596)
	v3599 = base.I64_rem_s(v3597, v3596)
	v3601 = v3599 * int64(60000000)
	v3602 = v3590 + v3601
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3602
	if base.B2i32(v3601 < int64(0))^base.B2i32(v3602 < v3590) != 0 {
		v3946 = v3583
		goto L502
	} else {
		goto L862
	}
L862:
	;
	v3612 = (v2500 - v3597*int64(100)) * int64(1000000)
	v3613 = v3602 + v3612
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3613
	if base.B2i32(v3612 < int64(0))^base.B2i32(v3613 < v3602) != 0 {
		v3946 = v3583
		goto L502
	} else {
		goto L863
	}
L863:
	;
	if base.F64_ne(v2503, float64(0)) != 0 {
		goto L864
	} else {
		goto L865
	}
L864:
	;
	if base.F64_lt(base.F64_abs(v2503), float64(9.223372036854776e+18)) != 0 {
		goto L869
	} else {
		goto L870
	}
L865:
	;
	goto L866
L866:
	;
	v3946 = int32(0)
	goto L502
L867:
	;
	v3640 = v3613 + v3639
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3640
	if base.B2i32(v3639 < int64(0))^base.B2i32(v3640 < v3613) != 0 {
		v3946 = v3583
		goto L502
	} else {
		goto L876
	}
L868:
	;
	v3628 = base.F64_sub(v2503, base.F64_convert_i64_s(v3626))
	if base.F64_gt(v3628, float64(0.5)) != 0 {
		goto L872
	} else {
		goto L873
	}
L869:
	;
	v3624 = base.I64_trunc_f64_s(v2503)
	v3626 = v3624
	goto L868
L870:
	;
	goto L871
L871:
	;
	v3626 = int64(-9223372036854775807 - 1)
	goto L868
L872:
	;
	v3639 = v3626 + int64(1)
	goto L867
L873:
	;
	goto L874
L874:
	;
	if base.F64_lt(v3628, float64(-0.5)) == int32(0) {
		v3639 = v3626
		goto L867
	} else {
		goto L875
	}
L875:
	;
	v3639 = v3626 - int64(1)
	goto L867
L876:
	;
	goto L866
L877:
	;
	v3653 = v2325 + int32(56)
	v3656 = int64(3600000000)
	v3657 = int64(0)
	v3662 = int64(32)
	v3665 = int64(base.Ui64(v2500) >> (uint(v3662) % 64))
	v3668 = int64(4294967295)
	v3671 = v2500 & v3668
	v3672 = v3656 * v3671
	v3676 = int64(base.Ui64(v3672)>>(uint(v3662)%64)) + v3656*v3665
	v3683 = v3671*v3657 + v3676&v3668
	*(*int64)(unsafe.Add(mBase, uint32(v3653)+8)) = v2500*v3657 + v2500>>(uint(int64(63))%64)*v3656 + v3657*v3665 + int64(base.Ui64(v3676)>>(uint(v3662)%64)) + int64(base.Ui64(v3683)>>(uint(v3662)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3653))) = v3672&v3668 | v3683<<(uint(v3662)%64)
	goto L878
L878:
	;
	v3694 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+64))
	v3695 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+56))
	if v3694 != v3695>>(uint(int64(63))%64) {
		goto L879
	} else {
		goto L880
	}
L879:
	;
	v3946 = int32(-2)
	goto L502
L880:
	;
	goto L881
L881:
	;
	v3700 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3701 = v3700 + v3695
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3701
	if base.B2i32(v3695 < int64(0))^base.B2i32(v3701 < v3700) != 0 {
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v3946 = int32(-2)
	goto L502
L883:
	;
	goto L884
L884:
	;
	if base.F64_eq(v2503, float64(0)) != 0 {
		goto L885
	} else {
		goto L886
	}
L885:
	;
	if v2505 == int32(0) {
		goto L897
	} else {
		goto L898
	}
L886:
	;
	v3711 = base.F64_mul(v2503, float64(3.6e+09))
	if base.F64_lt(base.F64_abs(v3711), float64(9.223372036854776e+18)) != 0 {
		goto L888
	} else {
		goto L889
	}
L887:
	;
	v3719 = base.F64_sub(v3711, base.F64_convert_i64_s(v3717))
	if base.F64_gt(v3719, float64(0.5)) != 0 {
		goto L892
	} else {
		goto L893
	}
L888:
	;
	v3715 = base.I64_trunc_f64_s(v3711)
	v3717 = v3715
	goto L887
L889:
	;
	goto L890
L890:
	;
	v3717 = int64(-9223372036854775807 - 1)
	goto L887
L891:
	;
	v3731 = v3730 + v3701
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3731
	if base.B2i32(v3730 < int64(0))^base.B2i32(v3731 < v3701) == int32(0) {
		goto L885
	} else {
		goto L896
	}
L892:
	;
	v3730 = v3717 + int64(1)
	goto L891
L893:
	;
	goto L894
L894:
	;
	if base.F64_lt(v3719, float64(-0.5)) == int32(0) {
		v3730 = v3717
		goto L891
	} else {
		goto L895
	}
L895:
	;
	v3730 = v3717 - int64(1)
	goto L891
L896:
	;
	v3946 = int32(-2)
	goto L502
L897:
	;
	v3946 = int32(0)
	goto L502
L898:
	;
	goto L899
L899:
	;
	v3752 = F_ParseISO8601Number(m, v2486, v2325+int32(108), v2325+int32(96), v2325+int32(88))
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L30
	} else {
		goto L900
	}
L900:
	;
	if v3752 != 0 {
		v3946 = v3752
		goto L502
	} else {
		goto L901
	}
L901:
	;
	v3754 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+96))
	v3755 = *(*float64)(unsafe.Add(mBase, uint32(v2325)+88))
	v3756 = int64(60000000)
	v3757 = int32(0)
	v3761 = m.G0
	v3763 = v3761 - int32(16)
	m.G0 = v3763
	v3765 = int64(63)
	F___multi3(m, v3763, v3754, v3754>>(uint(v3765)%64), v3756, int64(0))
	mBase = m.M
	v3770 = *(*int64)(unsafe.Add(mBase, uint32(v3763)+8))
	v3771 = *(*int64)(unsafe.Add(mBase, uint32(v3763)))
	if v3770 != v3771>>(uint(v3765)%64) {
		v3817 = v3757
		goto L903
	} else {
		goto L904
	}
L902:
	;
	if v3817 == int32(0) {
		goto L916
	} else {
		goto L917
	}
L903:
	;
	m.G0 = v3763 + int32(16)
	goto L902
L904:
	;
	v3775 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3776 = v3775 + v3771
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3776
	if base.B2i32(v3771 < int64(0))^base.B2i32(v3776 < v3775) != 0 {
		v3817 = v3757
		goto L903
	} else {
		goto L905
	}
L905:
	;
	if base.F64_eq(v3755, float64(0)) != 0 {
		v3817 = int32(1)
		goto L903
	} else {
		goto L906
	}
L906:
	;
	v3786 = base.F64_mul(v3755, base.F64_convert_i64_u(v3756))
	if base.F64_lt(base.F64_abs(v3786), float64(9.223372036854776e+18)) != 0 {
		goto L908
	} else {
		goto L909
	}
L907:
	;
	v3794 = base.F64_sub(v3786, base.F64_convert_i64_s(v3792))
	if base.F64_gt(v3794, float64(0.5)) != 0 {
		goto L912
	} else {
		goto L913
	}
L908:
	;
	v3790 = base.I64_trunc_f64_s(v3786)
	v3792 = v3790
	goto L907
L909:
	;
	goto L910
L910:
	;
	v3792 = int64(-9223372036854775807 - 1)
	goto L907
L911:
	;
	v3806 = v3805 + v3776
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3806
	v3817 = base.B2i32(base.B2i32(v3805 < int64(0))^base.B2i32(v3806 < v3776) == int32(0))
	goto L903
L912:
	;
	v3805 = v3792 + int64(1)
	goto L911
L913:
	;
	goto L914
L914:
	;
	if base.F64_lt(v3794, float64(-0.5)) == int32(0) {
		v3805 = v3792
		goto L911
	} else {
		goto L915
	}
L915:
	;
	v3805 = v3792 - int64(1)
	goto L911
L916:
	;
	v3946 = int32(-2)
	goto L502
L917:
	;
	goto L918
L918:
	;
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+108))
	v3827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3826))))
	if v3827 == int32(0) {
		v3946 = v3827
		goto L502
	} else {
		goto L919
	}
L919:
	;
	if v3827 != int32(58) {
		goto L920
	} else {
		goto L921
	}
L920:
	;
	v3946 = int32(-1)
	goto L502
L921:
	;
	goto L922
L922:
	;
	v3834 = v3826 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2325)+108)) = v3834
	v3842 = F_ParseISO8601Number(m, v3834, v2325+int32(108), v2325+int32(96), v2325+int32(88))
	mBase = m.M
	v3843 = m.ExcPending
	if v3843 != 0 {
		goto L30
	} else {
		goto L923
	}
L923:
	;
	if v3842 != 0 {
		v3946 = v3842
		goto L502
	} else {
		goto L924
	}
L924:
	;
	v3845 = *(*int64)(unsafe.Add(mBase, uint32(v2325)+96))
	v3846 = *(*float64)(unsafe.Add(mBase, uint32(v2325)+88))
	v3847 = int64(1000000)
	v3848 = int32(0)
	v3852 = m.G0
	v3854 = v3852 - int32(16)
	m.G0 = v3854
	v3856 = int64(63)
	F___multi3(m, v3854, v3845, v3845>>(uint(v3856)%64), v3847, int64(0))
	mBase = m.M
	v3861 = *(*int64)(unsafe.Add(mBase, uint32(v3854)+8))
	v3862 = *(*int64)(unsafe.Add(mBase, uint32(v3854)))
	if v3861 != v3862>>(uint(v3856)%64) {
		v3908 = v3848
		goto L926
	} else {
		goto L927
	}
L925:
	;
	if v3908 == int32(0) {
		v3946 = int32(-2)
		goto L502
	} else {
		goto L939
	}
L926:
	;
	m.G0 = v3854 + int32(16)
	goto L925
L927:
	;
	v3866 = *(*int64)(unsafe.Add(mBase, uint32(v2332)))
	v3867 = v3866 + v3862
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3867
	if base.B2i32(v3862 < int64(0))^base.B2i32(v3867 < v3866) != 0 {
		v3908 = v3848
		goto L926
	} else {
		goto L928
	}
L928:
	;
	if base.F64_eq(v3846, float64(0)) != 0 {
		v3908 = int32(1)
		goto L926
	} else {
		goto L929
	}
L929:
	;
	v3877 = base.F64_mul(v3846, base.F64_convert_i64_u(v3847))
	if base.F64_lt(base.F64_abs(v3877), float64(9.223372036854776e+18)) != 0 {
		goto L931
	} else {
		goto L932
	}
L930:
	;
	v3885 = base.F64_sub(v3877, base.F64_convert_i64_s(v3883))
	if base.F64_gt(v3885, float64(0.5)) != 0 {
		goto L935
	} else {
		goto L936
	}
L931:
	;
	v3881 = base.I64_trunc_f64_s(v3877)
	v3883 = v3881
	goto L930
L932:
	;
	goto L933
L933:
	;
	v3883 = int64(-9223372036854775807 - 1)
	goto L930
L934:
	;
	v3897 = v3896 + v3867
	*(*int64)(unsafe.Add(mBase, uint32(v2332))) = v3897
	v3908 = base.B2i32(base.B2i32(v3896 < int64(0))^base.B2i32(v3897 < v3867) == int32(0))
	goto L926
L935:
	;
	v3896 = v3883 + int64(1)
	goto L934
L936:
	;
	goto L937
L937:
	;
	if base.F64_lt(v3885, float64(-0.5)) == int32(0) {
		v3896 = v3883
		goto L934
	} else {
		goto L938
	}
L938:
	;
	v3896 = v3883 - int64(1)
	goto L934
L939:
	;
	v3918 = *(*int32)(unsafe.Add(mBase, uint32(v2325)+108))
	v3919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3918))))
	if v3919 != 0 {
		goto L940
	} else {
		goto L941
	}
L940:
	;
	v3920 = int32(-1)
	goto L942
L941:
	;
	v3920 = int32(0)
	goto L942
L942:
	;
	v3946 = v3920
	goto L502
L943:
	;
	m.G0 = v40 + int32(528)
	return v4087
L944:
	;
	if v3987 == int32(-2) {
		goto L947
	} else {
		goto L948
	}
L945:
	;
	goto L946
L946:
	;
	v4024 = F_palloc(m, int32(16))
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L30
	} else {
		goto L951
	}
L947:
	;
	v4014 = int32(-4)
	goto L949
L948:
	;
	v4014 = v3987
	goto L949
L949:
	;
	F_DateTimeParseError(m, v4014, v40+int32(8), v44, int32(290483), v43)
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L30
	} else {
		goto L950
	}
L950:
	;
	v4020 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v4020)
	v4087 = int32(0)
	goto L943
L951:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v40)+500))
	switch v4026 - int32(9) {
	case 0:
		goto L955
	case 1:
		goto L953
	default:
		goto L954
	case 8:
		goto L956
	}
L952:
	;
	F_AdjustIntervalForTypmod(m, v4024, v42, v43)
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L30
	} else {
		goto L968
	}
L953:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4024)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v4024))) = int64(9223372036854775807)
	goto L952
L954:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4067 = m.ExcPending
	if v4067 != 0 {
		goto L30
	} else {
		goto L965
	}
L955:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4024)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v4024))) = int64(-9223372036854775807 - 1)
	goto L952
L956:
	;
	v4029 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+516)))
	v4030 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+520)))
	v4033 = v4029 + v4030*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v4033-int64(2147483648)) {
		goto L957
	} else {
		goto L958
	}
L957:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v4024)+12)) = uint32(v4033)
	v4039 = *(*int32)(unsafe.Add(mBase, uint32(v40)+512))
	*(*int32)(unsafe.Add(mBase, uint32(v4024)+8)) = v4039
	v4041 = *(*int64)(unsafe.Add(mBase, uint32(v40)+504))
	*(*int64)(unsafe.Add(mBase, uint32(v4024))) = v4041
	goto L952
L958:
	;
	goto L959
L959:
	;
	v4043 = int32(0)
	v4044 = F_errsave_start(m, v43)
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L30
	} else {
		goto L960
	}
L960:
	;
	if v4044 == int32(0) {
		v4087 = v4043
		goto L943
	} else {
		goto L961
	}
L961:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v4050 = m.ExcPending
	if v4050 != 0 {
		goto L30
	} else {
		goto L962
	}
L962:
	;
	F_errmsg(m, int32(378436), int32(0))
	mBase = m.M
	v4054 = m.ExcPending
	if v4054 != 0 {
		goto L30
	} else {
		goto L963
	}
L963:
	;
	F_errsave_finish(m, v43, int32(465436), int32(948), int32(262591))
	mBase = m.M
	v4059 = m.ExcPending
	if v4059 != 0 {
		goto L30
	} else {
		goto L964
	}
L964:
	;
	v4087 = v4043
	goto L943
L965:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v44
	v4069 = *(*int32)(unsafe.Add(mBase, uint32(v40)+500))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v4069
	F_errmsg_internal(m, int32(662007), v40)
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L30
	} else {
		goto L966
	}
L966:
	;
	F_errfinish(m, int32(465436), int32(961), int32(262591))
	mBase = m.M
	v4078 = m.ExcPending
	if v4078 != 0 {
		goto L30
	} else {
		goto L967
	}
L967:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L968:
	;
	v4087 = v4024
	goto L943
}
func F_interval_lerp(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	v16 = F_DirectFunctionCall2Coll(m, int32(1475), v4, l1, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v22 = F_DirectFunctionCall2Coll(m, int32(1474), v4, v16, v7+int32(8))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_DirectFunctionCall2Coll(m, int32(1473), v4, v22, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v24
			}
		}
	}
}
func F_interval_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	return base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64)))) != int64(0))
}
func F_interval_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v228 float64
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v321 int64
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
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 float64
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v436 int64
	_ = v436
	var v441 int64
	_ = v441
	var v444 int64
	_ = v444
	var v445 int64
	_ = v445
	var v450 int64
	_ = v450
	var v453 int64
	_ = v453
	var v456 int64
	_ = v456
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v464 int64
	_ = v464
	var v471 int64
	_ = v471
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v484 int64
	_ = v484
	var v490 int64
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v520 int64
	_ = v520
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v556 int64
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(1)
	v27 = v22 + v26
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v32 = v30 & v26
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v27
	goto L5
L4:
	;
	v33 = v22 + int32(4)
	goto L5
L5:
	;
	if v30 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v64 = F_downcase_truncate_identifier(m, v33, v61, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v36 = int32(4)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v38&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v51 = int32(1)
	if v32 != 0 {
		v61 = int32(base.Ui32(v30)>>(uint(v51)%32)) - v51
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v47 = v36
	goto L12
L11:
	;
	v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
	goto L12
L12:
	;
	if v38 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = v36
	goto L15
L14:
	;
	v50 = v47
	goto L15
L15:
	;
	v61 = v50
	goto L6
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v67 = v17 + int32(76)
	v74 = *(*int32)(unsafe.Add(mBase, _consts[1115]))
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v135 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1115])) = v118
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v118)+11)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v126
	v135 = v125
	goto L18
L20:
	;
	v76 = F_strncmp(m, v64, v74, int32(10))
	mBase = m.M
	if v76 == int32(0) {
		v118 = v74
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
	v85 = int32(1610400)
	v87 = int32(1611360)
	goto L24
L23:
	;
	goto L22
L24:
	;
	v94 = v85 + (v87-v85)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94))))
	v96 = v79 - v95
	if v96 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = int32(0)
	v135 = int32(31)
	goto L18
L26:
	;
	v100 = F_strncmp(m, v64, v94, int32(10))
	mBase = m.M
	if v100 == int32(0) {
		v118 = v94
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v103 = v96
	goto L28
L28:
	;
	v107 = base.B2i32(v103 < int32(0))
	if v103 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v103 = v100
	goto L28
L30:
	;
	v108 = v94 - int32(16)
	goto L32
L31:
	;
	v108 = v87
	goto L32
L32:
	;
	if v103 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v111 = v85
	goto L35
L34:
	;
	v111 = v94 + int32(16)
	goto L35
L35:
	;
	if base.Ui32(v111) <= base.Ui32(v108) {
		v85 = v111
		v87 = v108
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v139 = v17 + int32(76)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[1042]))
	if v146 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v208 = v135
	goto L39
L39:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v209 != int32(2147483647) {
		goto L62
	} else {
		goto L63
	}
L40:
	;
	v208 = v207
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1042])) = v190
	v197 = int32(*(*int8)(unsafe.Add(mBase, uint32(v190)+11)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v198
	v207 = v197
	goto L40
L42:
	;
	v148 = F_strncmp(m, v64, v146, int32(10))
	mBase = m.M
	if v148 == int32(0) {
		v190 = v146
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v64))))
	v157 = int32(1609248)
	v159 = int32(1610384)
	goto L46
L45:
	;
	goto L44
L46:
	;
	v166 = v157 + (v159-v157)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v167 = int32(*(*int8)(unsafe.Add(mBase, uint32(v166))))
	v168 = v151 - v167
	if v168 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(0)
	v207 = int32(31)
	goto L40
L48:
	;
	v172 = F_strncmp(m, v64, v166, int32(10))
	mBase = m.M
	if v172 == int32(0) {
		v190 = v166
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v175 = v168
	goto L50
L50:
	;
	v179 = base.B2i32(v175 < int32(0))
	if v175 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v175 = v172
	goto L50
L52:
	;
	v180 = v166 - int32(16)
	goto L54
L53:
	;
	v180 = v159
	goto L54
L54:
	;
	if v175 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = v157
	goto L57
L56:
	;
	v183 = v166 + int32(16)
	goto L57
L57:
	;
	if base.Ui32(v183) <= base.Ui32(v180) {
		v157 = v183
		v159 = v180
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	m.G0 = v17 + int32(80)
	return v571
L60:
	;
	if v208 == int32(17) {
		goto L100
	} else {
		goto L101
	}
L61:
	;
	switch v208 {
	case 0, 17:
		goto L70
	default:
		goto L71
	}
L62:
	;
	if v209 != int32(-2147483648) {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v221 != int32(2147483647) {
		goto L60
	} else {
		goto L68
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v214 != int32(-2147483648) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	if v217 != int64(-9223372036854775807-1) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v228 = math.Float64frombits(uint64(0xfff0000000000000))
	goto L61
L68:
	;
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	if v224 != int64(9223372036854775807) {
		goto L60
	} else {
		goto L69
	}
L69:
	;
	v228 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L61
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if base.Ui32(v251) <= base.Ui32(int32(30)) {
		goto L79
	} else {
		goto L80
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v237 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v64
	F_errmsg(m, int32(178101), v17+int32(48))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(465436), int32(6058), int32(78092))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v301 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v301)
	v571 = int32(0)
	goto L59
L78:
	;
	if l1 != 0 {
		goto L89
	} else {
		goto L90
	}
L79:
	;
	v255 = int32(1) << (uint(v251) % 32)
	if v255&int32(506464256) != 0 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	if v255&int32(1640759296) != 0 {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v269 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v64
	F_errmsg(m, int32(178138), v17-int32(-64))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(465436), int32(6089), int32(78092))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	if base.F64_lt(v228, float64(0)) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v299 = F_Float8GetDatum(m, v228)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L97
	}
L92:
	;
	v286 = int32(0)
	v290 = F_DirectFunctionCall3Coll(m, int32(408), v286, int32(10332), v286, int32(-1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v293 = int32(0)
	v297 = F_DirectFunctionCall3Coll(m, int32(408), v293, int32(10343), v293, int32(-1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	v571 = v290
	goto L59
L96:
	;
	v571 = v297
	goto L59
L97:
	;
	v571 = v299
	goto L59
L98:
	;
	if l1 != 0 {
		goto L153
	} else {
		goto L154
	}
L99:
	;
	v556 = base.I64_extend32_s(v321) + base.I64_extend32_s(v318)*int64(1000000)
	goto L98
L100:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v308 = base.I64_div_s(v306, int64(3600000000))
	v311 = v308*int64(-3600000000) + v306
	v313 = base.I64_div_s(v311, int64(60000000))
	v316 = v313*int64(-60000000) + v311
	v318 = base.I64_div_s(v316, int64(1000000))
	v321 = v318*int64(4293967296) + v316
	v322 = base.I32_wrap_i64(v321)
	v323 = base.I32_wrap_i64(v318)
	v325 = base.I32_rem_s(v209, int32(12))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	switch v327 - int32(18) {
	case 0:
		goto L113
	case 1:
		goto L112
	case 2:
		v556 = v308
		goto L98
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	case 8:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L114
	case 12:
		goto L99
	default:
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v208 != 0 {
		goto L133
	} else {
		goto L134
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L128
	}
L104:
	;
	v398 = base.I32_div_s(v209, int32(12000))
	v556 = base.I64_extend_i32_s(v398)
	goto L98
L105:
	;
	v395 = base.I32_div_s(v209, int32(1200))
	v556 = base.I64_extend_i32_s(v395)
	goto L98
L106:
	;
	v392 = base.I32_div_s(v209, int32(120))
	v556 = base.I64_extend_i32_s(v392)
	goto L98
L107:
	;
	v389 = base.I32_div_s(v209, int32(12))
	v556 = base.I64_extend_i32_s(v389)
	goto L98
L108:
	;
	if int32(0) <= v209 {
		goto L125
	} else {
		goto L126
	}
L109:
	;
	v556 = base.I64_extend_i32_s(v325)
	goto L98
L110:
	;
	v365 = base.I32_div_s(v326, int32(7))
	v556 = base.I64_extend_i32_s(v365)
	goto L98
L111:
	;
	v556 = base.I64_extend_i32_s(v326)
	goto L98
L112:
	;
	v556 = base.I64_extend32_s(v313)
	goto L98
L113:
	;
	if l1 != 0 {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	if l1 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v336 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v321)+base.I64_extend32_s(v318)*int64(1000000), int32(3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v339 = float64(1000)
	v345 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v323), v339), base.F64_div(base.F64_convert_i32_s(v322), v339)))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v571 = v336
	goto L59
L119:
	;
	v571 = v345
	goto L59
L120:
	;
	v353 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v321)+base.I64_extend32_s(v318)*int64(1000000), int32(6))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v360 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v322), float64(1e+06)), base.F64_convert_i32_s(v323)))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	v571 = v353
	goto L59
L124:
	;
	v571 = v360
	goto L59
L125:
	;
	v373 = base.I32_div_u_s(v325&int32(255), int32(3))
	v556 = base.I64_extend_i32_u(v373 + int32(1))
	goto L98
L126:
	;
	goto L127
L127:
	;
	v380 = base.I32_rem_s(int32(0)-v209, int32(12))
	v383 = base.I32_div_s(base.I32_extend8_s(v380), int32(-3))
	v556 = base.I64_extend8_s(base.I64_extend_i32_u(v383 - int32(1)))
	goto L98
L128:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v408 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v64
	F_errmsg(m, int32(178138), v17)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(465436), int32(6233), int32(230156))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L148
	}
L134:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v420 != int32(11) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	if l1 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v424 = v17 + int32(16)
	v425 = int32(12)
	v426 = base.I32_rem_s(v209, v425)
	v431 = base.I32_div_s(v209, v425)
	v436 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+8)))
	v441 = (base.I64_extend_i32_s(v426*int32(120)) + base.I64_extend_i32_s(v431)*int64(1461) + v436<<(uint(int64(2))%64)) * int64(21600)
	v444 = int64(1000000)
	v445 = int64(0)
	v450 = int64(32)
	v453 = int64(base.Ui64(v441) >> (uint(v450) % 64))
	v456 = int64(4294967295)
	v459 = v441 & v456
	v460 = v444 * v459
	v464 = int64(base.Ui64(v460)>>(uint(v450)%64)) + v444*v453
	v471 = v459*v445 + v464&v456
	*(*int64)(unsafe.Add(mBase, uint32(v424)+8)) = v441*v445 + v441>>(uint(int64(63))%64)*v444 + v445*v453 + int64(base.Ui64(v464)>>(uint(v450)%64)) + int64(base.Ui64(v471)>>(uint(v450)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v424))) = v460&v456 | v471<<(uint(v450)%64)
	goto L139
L137:
	;
	goto L138
L138:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v509 = int32(12)
	v510 = base.I32_div_s(v209, v509)
	v520 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v527 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v505), float64(86400)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v209-v510*v509), float64(2.592e+06)), base.F64_add(base.F64_mul(base.F64_convert_i32_s(v510), float64(3.15576e+07)), base.F64_div(base.F64_convert_i64_s(v520), float64(1e+06))))))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L147
	}
L139:
	;
	v482 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	v484 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	if v483 != v484>>(uint(int64(63))%64) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v498 = F_int64_div_fast_to_numeric(m, v482, int32(6))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L144
	}
L141:
	;
	v490 = v482 + v484
	if base.B2i32(v482 < int64(0))^base.B2i32(v490 < v484) != 0 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v494 = F_int64_div_fast_to_numeric(m, v490, int32(6))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v571 = v494
	goto L59
L144:
	;
	v500 = F_int64_to_numeric(m, v441)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v503 = F_numeric_add_opt_error(m, v498, v500, int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v571 = v503
	goto L59
L147:
	;
	v571 = v527
	goto L59
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v537 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v64
	F_errmsg(m, int32(178101), v17+int32(32))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(465436), int32(6294), int32(230156))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
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
	v557 = F_int64_to_numeric(m, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v560 = F_Float8GetDatum(m, base.F64_convert_i64_s(v556))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v571 = v557
	goto L59
L157:
	;
	v571 = v560
	goto L59
}
func F_interval_sum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v43 int64
	_ = v43
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != 0 {
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 == int32(0) {
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			return int32(0)
		} else {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
			v10 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
			if v9+v10 != int64(0)-v13 {
				v23 = int64(0)
				if base.B2i32(v23 < v9)&base.B2i32(v23 < v13) == int32(0) {
					v31 = F_palloc(m, int32(16))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
						if int64(0) < v35 {
							*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(9223372036854775807)
							return v31
						} else {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
							if int64(0) < v43 {
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(-9223372036854775807 - 1)
								return v31
							} else {
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v51
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v6)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v53
								return v31
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(378436), int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(465436), int32(4285), int32(269321))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				return int32(0)
			}
		}
	}
}
func F_interval_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != int32(457) {
		v49 = v2
		return v49
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 != int32(7) {
			v49 = v2
			return v49
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
			if v17 != 0 {
				v49 = v2
				return v49
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				if v19 < int32(0) {
					v44 = F_relabel_to_typmod(m, v18, v19)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v49 = v44
						return v49
					}
				} else {
					v22 = F_exprTypmod(m, v18)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_intervaltypmodleastfield(m, v22)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = F_intervaltypmodleastfield(m, v19)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v26) < base.Ui32(v28) {
									v49 = v2
									return v49
								} else {
									if v26 != 0 {
										v44 = F_relabel_to_typmod(m, v18, v19)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v49 = v44
											return v49
										}
									} else {
										v32 = v19 & int32(65535)
										if base.Ui32(int32(5)) < base.Ui32(v32) {
											v44 = F_relabel_to_typmod(m, v18, v19)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v49 = v44
												return v49
											}
										} else {
											if v22 < int32(0) {
												v38 = int32(-1)
											} else {
												v38 = v22
											}
											if base.Ui32(v32) < base.Ui32(v38&int32(65535)) {
												v49 = v2
												return v49
											} else {
												v44 = F_relabel_to_typmod(m, v18, v19)
												mBase = m.M
												v45 = m.ExcPending
												if v45 != 0 {
													return int32(0)
												} else {
													v49 = v44
													return v49
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
