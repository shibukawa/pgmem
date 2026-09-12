package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitOnLock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v212 int64
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int64
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v302 int64
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v320 int64
	_ = v320
	var v322 int32
	_ = v322
	var v323 int64
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v572 int32
	_ = v572
	var v716 int32
	_ = v716
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int64
	_ = v734
	var v735 int64
	_ = v735
	var v743 int64
	_ = v743
	var v745 int32
	_ = v745
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v938 int32
	_ = v938
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v1006 int32
	_ = v1006
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1246 int32
	_ = v1246
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1465 int32
	_ = v1465
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int64
	_ = v1508
	var v1510 int64
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1530 int32
	_ = v1530
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1697 int64
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int64
	_ = v1706
	var v1707 int64
	_ = v1707
	var v1722 int64
	_ = v1722
	var v1726 int64
	_ = v1726
	var v1727 int64
	_ = v1727
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1804 int32
	_ = v1804
	var v1813 int32
	_ = v1813
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1839 int32
	_ = v1839
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2086 int64
	_ = v2086
	var v2087 int64
	_ = v2087
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2136 int32
	_ = v2136
	var v2145 int32
	_ = v2145
	var v2158 int32
	_ = v2158
	var v2159 int64
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	v3 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = l0
	v36 = l1
	v37 = v3
	v38 = int32(-1)
	v42 = v32
	v44 = v3
	v47 = v3
	v48 = v3
	v51 = v32
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v38 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v2158 = int32(m.ExcTag)
	v2159 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2158 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L6:
	;
	v67 = v51 - int32(160)
	m.G0 = v67
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v67
	*(*int32)(unsafe.Add(mBase, _consts[717])) = v36
	*(*int32)(unsafe.Add(mBase, _consts[718])) = v35
	v78 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v80 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	goto L9
L7:
	;
	v84 = v37
	v85 = v44
	v86 = v47
	v87 = v48
	v88 = v51
	goto L8
L8:
	;
	if v84 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v42
	goto L12
L10:
	;
	v84 = int32(0)
	v85 = v67
	v86 = v78
	v87 = v80
	v88 = v67
	goto L8
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v85
	v96 = int32(0)
	v97 = m.G0
	v99 = v97 - int32(400)
	m.G0 = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v107 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v112 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v87
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v85
	F_pg_re_throw(m)
	mBase = m.M
	v2128 = m.ExcPending
	if v2128 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L338
	}
L16:
	;
	v151 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[719])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[720])) = v151
	v157 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	if base.Ui32(v157) <= base.Ui32(int32(1)) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	if v122 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L18:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+316))
	v120 = base.B2i32(v118 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v120)
	v122 = v120
	goto L20
L19:
	;
	v122 = v96
	goto L20
L20:
	;
	goto L17
L21:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[309])))
	if v126 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v127 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v127 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L16
L27:
	;
	F_errcode(m, int32(16908292))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(14214), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(572448), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(492372), int32(922), int32(316206))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v218 = v35
	v219 = v36
	v220 = v99
	v225 = v42
	v227 = v85
	v230 = v86
	v231 = v87
	v234 = v88
	v235 = v96
	v236 = v108
	v237 = int32(1)
	v240 = v109
	v241 = v101&int32(15)<<(uint(int32(7))%32) + v107 + int32(23296)
	v242 = v99 + int32(32)
	v244 = v212
	goto L44
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[721]))
	if int32(0) < v161 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _consts[722])))
	if v190 != int32(1) {
		v212 = int64(0)
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v185 = *(*int64)(unsafe.Add(mBase, _consts[723]))
	*(*int64)(unsafe.Add(mBase, uint32(v183)+112)) = v185
	v212 = int64(0)
	goto L32
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v99)+352)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+384)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v99)+376)) = int64(2)
	v170 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+360)) = v170
	F_enable_timeouts(m, v99+int32(352), int32(2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	F_enable_timeout_after(m, int32(1), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		v2129 = v35
		v2130 = v36
		v2136 = v42
		v2145 = v88
		goto L5
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	goto L36
L42:
	;
	v196 = m.G0
	v197 = int32(16)
	v198 = v196 - v197
	m.G0 = v198
	F___gettimeofday(m, v198)
	mBase = m.M
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	v202 = int64(*(*int32)(unsafe.Add(mBase, uint32(v198)+8)))
	m.G0 = v198 + v197
	goto L43
L43:
	;
	v212 = v202 + v201*int64(1000000) - int64(946684800000000)
	goto L32
L44:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	if base.Ui32(int32(2)) <= base.Ui32(v248) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	if base.Ui32(int32(1)) < base.Ui32(v2050) {
		goto L326
	} else {
		goto L327
	}
L46:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1478)+16))
	v1481 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v1483 = base.B2i32(v1481 != int32(4))
	v1484 = v1483 & v237
	if v1481 != int32(4) {
		v1658 = v1481
		v1661 = v1484
		goto L213
	} else {
		goto L214
	}
L47:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v220)+304)) = v251
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	*(*int64)(unsafe.Add(mBase, uint32(v220)+296)) = v253
	v256 = v220 + int32(296)
	v261 = base.B2i32(v235 == int32(0)) & base.B2i32(v244 != int64(0))
	v262 = m.G0
	v264 = v262 - int32(80)
	m.G0 = v264
	v271 = *(*int64)(unsafe.Add(mBase, _consts[725]))
	*(*int64)(unsafe.Add(mBase, uint32(v264+int32(16)))) = v271
	v274 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	*(*uint8)(unsafe.Add(mBase, uint32(v264+int32(79)))) = uint8(base.B2i32(v274 == int32(3)))
	goto L50
L48:
	;
	goto L49
L49:
	;
	v769 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+14)))
	v775 = F_WaitLatch(m, v769, int32(33), int32(0), v772|int32(50331648))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L106
	}
L50:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+79)))
	if v278 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v306 = m.G0
	v307 = int32(16)
	v308 = v306 - v307
	m.G0 = v308
	F___gettimeofday(m, v308)
	mBase = m.M
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v308)))
	v312 = int64(*(*int32)(unsafe.Add(mBase, uint32(v308)+8)))
	m.G0 = v308 + v307
	v320 = v312 + v311*int64(1000000) - int64(946684800000000)
	goto L57
L52:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[727]))
	if v283 < int32(0) {
		v302 = int64(0)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[728]))
	if v293 < int32(0) {
		v302 = int64(0)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v264)+16))
	v302 = v286 + base.I64_extend_i32_u(v283)*int64(1000)
	goto L51
L56:
	;
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v264)+16))
	v302 = v296 + base.I64_extend_i32_u(v293)*int64(1000)
	goto L51
L57:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v322)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v322)+112)) = v323
	if v323 == int64(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	*(*int64)(unsafe.Add(mBase, uint32(v328)+112)) = v320
	goto L60
L59:
	;
	goto L60
L60:
	;
	v331 = base.B2i32(v302 == int64(0))
	if v302 == int64(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+14)))
	F_ProcWaitForSignal(m, v372|int32(50331648))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L72
	}
L62:
	;
	if v302 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	if v320 < v302 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v335 = F_GetLockConflicts(m, v256, int32(8), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+14)))
	F_ResolveRecoveryConflictWithVirtualXIDs(m, v335, int32(9), v338|int32(50331648), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L66
	}
L66:
	;
	goto L61
L67:
	;
	*(*int32)(unsafe.Add(mBase, _consts[729])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v357))) = int64(4)
	v364 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	*(*int32)(unsafe.Add(mBase, uint32(v357)+8)) = v364
	F_enable_timeouts(m, v264+int32(16), v356)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L71
	}
L68:
	;
	v356 = int32(1)
	v357 = v264 + int32(16)
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+32)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v264)+16)) = int64(4294967302)
	*(*int32)(unsafe.Add(mBase, _consts[730])) = int32(0)
	v356 = int32(2)
	v357 = v264 + int32(40)
	goto L67
L71:
	;
	goto L61
L72:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	if v378 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v572 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[475])) = v572
	*(*int32)(unsafe.Add(mBase, _consts[476])) = v572
	*(*uint8)(unsafe.Add(mBase, _consts[477])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[483])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[731])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[732])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[733])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[734])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[735])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[736])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[737])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[738])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[739])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[740])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[741])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[742])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[743])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[744])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[745])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[746])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[747])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[748])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[749])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[750])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[751])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[752])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[753])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[754])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[755])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[756])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[757])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[758])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[759])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[760])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[761])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[762])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[763])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[764])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[765])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[766])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[767])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[768])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[769])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[770])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[771])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[772])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[773])) = uint8(v572)
	*(*uint8)(unsafe.Add(mBase, _consts[774])) = uint8(v572)
	goto L96
L74:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	if v380 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v385 = F_GetLockConflicts(m, v256, int32(8), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v387 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v394 = v385
	goto L78
L78:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v394)))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v419
	v422 = v264 + int32(8)
	v425 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	v427 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v431 = F_LWLockAcquire(m, v427+int32(512), int32(1))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L80
	}
L79:
	;
	if v261 != 0 {
		goto L73
	} else {
		goto L94
	}
L80:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v433 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v524+int32(512))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L92
	}
L82:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v441 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v452 = int32(0)
	goto L83
L83:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v425+int32(36)+v452<<(uint(int32(2))%32))))
	v477 = v441 + v474*int32(640)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)+52))
	if v478 != v439 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L81
L85:
	;
	goto L84
L86:
	;
	v491 = v452 + int32(1)
	if v491 != v433 {
		v452 = v491
		goto L83
	} else {
		goto L91
	}
L87:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v477)+56))
	if v480 != v438 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v482 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v477)+73)) = uint8(v482)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v477)+44))
	if v484 == v482 {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v488 = F_SendProcSignal(m, v484, int32(13), v439)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L90
	}
L90:
	;
	goto L81
L91:
	;
	goto L85
L92:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v394+int32(12))))
	if v533 != 0 {
		v394 = v394 + int32(8)
		goto L78
	} else {
		goto L93
	}
L93:
	;
	goto L79
L94:
	;
	*(*int32)(unsafe.Add(mBase, _consts[729])) = int32(0)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+14)))
	F_ProcWaitForSignal(m, v537|int32(50331648))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L95
	}
L95:
	;
	goto L73
L96:
	;
	v716 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[730])) = v716
	*(*int32)(unsafe.Add(mBase, _consts[729])) = v716
	m.G0 = v264 + int32(80)
	if v261 != int32(1) {
		v1465 = v235
		goto L46
	} else {
		goto L97
	}
L97:
	;
	v729 = m.G0
	v730 = int32(16)
	v731 = v729 - v730
	m.G0 = v731
	F___gettimeofday(m, v731)
	mBase = m.M
	v734 = *(*int64)(unsafe.Add(mBase, uint32(v731)))
	v735 = int64(*(*int32)(unsafe.Add(mBase, uint32(v731)+8)))
	m.G0 = v731 + v730
	v743 = v735 + v734*int64(1000000) - int64(946684800000000)
	goto L98
L98:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _consts[724]))
	goto L99
L99:
	;
	if base.B2i32(base.I64_extend_i32_s(v745)*int64(1000) <= v743-v244) == int32(0) {
		v1465 = v235
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v758 = F_GetLockConflicts(m, v218, int32(8), v220+int32(352))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v760 = int32(0)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	if v760 < v761 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v764 = v758
	goto L104
L103:
	;
	v764 = v760
	goto L104
L104:
	;
	F_LogRecoveryConflict(m, int32(9), v244, v743, v764, int32(1))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v1465 = int32(1)
	goto L46
L106:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, uint32(v778))) = int32(0)
	goto L107
L107:
	;
	v782 = *(*int32)(unsafe.Add(mBase, _consts[719]))
	if v782 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v788 = F_LWLockAcquire(m, v784+int32(23296), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v1443 == int32(0) {
		v1465 = v235
		goto L46
	} else {
		goto L211
	}
L111:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v795 = F_LWLockAcquire(m, v791+int32(23424), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v798 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v802 = F_LWLockAcquire(m, v798+int32(23552), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v805 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v809 = F_LWLockAcquire(m, v805+int32(23680), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v816 = F_LWLockAcquire(m, v812+int32(23808), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v819 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v823 = F_LWLockAcquire(m, v819+int32(23936), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v830 = F_LWLockAcquire(m, v826+int32(24064), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v837 = F_LWLockAcquire(m, v833+int32(24192), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v844 = F_LWLockAcquire(m, v840+int32(24320), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v851 = F_LWLockAcquire(m, v847+int32(24448), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v858 = F_LWLockAcquire(m, v854+int32(24576), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L121
	}
L121:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v865 = F_LWLockAcquire(m, v861+int32(24704), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v872 = F_LWLockAcquire(m, v868+int32(24832), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v879 = F_LWLockAcquire(m, v875+int32(24960), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v886 = F_LWLockAcquire(m, v882+int32(25088), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v893 = F_LWLockAcquire(m, v889+int32(25216), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)))
	if v897 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1315+int32(25216))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L195
	}
L128:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	if v900 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v903 = int32(0)
	v904 = m.G0
	v906 = v904 - int32(16)
	m.G0 = v906
	*(*int32)(unsafe.Add(mBase, _consts[775])) = v903
	*(*int32)(unsafe.Add(mBase, _consts[776])) = v903
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v903
	*(*int32)(unsafe.Add(mBase, _consts[713])) = v903
	v920 = F_DeadLockCheckRecurse(m, v896)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L131
	}
L130:
	;
	m.G0 = v906 + int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[720])) = v1246
	if v1246 != int32(3) {
		goto L127
	} else {
		goto L192
	}
L131:
	;
	if v920 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v925 = *(*int32)(unsafe.Add(mBase, _consts[708]))
	if int32(0) < v925 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v1106 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v1106
	*(*int32)(unsafe.Add(mBase, _consts[710])) = v1106
	*(*int32)(unsafe.Add(mBase, _consts[711])) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v906)+12)) = v1106
	v1118 = *(*int32)(unsafe.Add(mBase, _consts[777]))
	v1120 = v906 + int32(12)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v896)+616))
	if v1124 != 0 {
		goto L156
	} else {
		goto L157
	}
L135:
	;
	v938 = v903
	goto L138
L136:
	;
	goto L137
L137:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	if v1103 != 0 {
		goto L152
	} else {
		goto L153
	}
L138:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[714]))
	v961 = v958 + v938*int32(12)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v961)+4))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v961)+8))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v961)))
	v965 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+40)) = v965
	v968 = v964 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+36)) = v968
	*(*int32)(unsafe.Add(mBase, uint32(v964)+32)) = v968
	if v963 <= v965 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if int32(0) < v1066 {
		v1246 = int32(2)
		goto L130
	} else {
		goto L151
	}
L140:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+15)))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1055<<(uint(int32(2))%32))+uint32(_consts[707])))
	goto L148
L141:
	;
	v979 = v968
	v982 = v965
	goto L142
L142:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v962+v982<<(uint(int32(2))%32))))
	if v979 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v964)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+36)) = v968
	*(*int32)(unsafe.Add(mBase, uint32(v964)+32)) = v968
	goto L146
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1006)+4)) = v968
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v964)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1006))) = v1014
	*(*int32)(unsafe.Add(mBase, uint32(v1014)+4)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v964)+32)) = v1006
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v964)+40))
	v1019 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v964)+40)) = v1018 + v1019
	v1023 = v982 + v1019
	if v1023 == v963 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v964)+36))
	v979 = v1025
	v982 = v1023
	goto L142
L148:
	;
	F_ProcLockWakeup(m, v1060, v964)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v1064 = v938 + int32(1)
	v1066 = *(*int32)(unsafe.Add(mBase, _consts[708]))
	if v1064 < v1066 {
		v938 = v1064
		goto L138
	} else {
		goto L150
	}
L150:
	;
	goto L139
L151:
	;
	goto L137
L152:
	;
	v1104 = int32(4)
	goto L154
L153:
	;
	v1104 = int32(1)
	goto L154
L154:
	;
	v1246 = v1104
	goto L130
L155:
	;
	if v1225 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	v1125 = v1124
	goto L158
L157:
	;
	v1125 = v896
	goto L158
L158:
	;
	v1126 = int32(0)
	v1128 = *(*int32)(unsafe.Add(mBase, _consts[709]))
	v1130 = *(*int32)(unsafe.Add(mBase, _consts[710]))
	if v1126 < v1130 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1133 = v1126
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[710])) = v1130 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1128+v1130<<(uint(int32(2))%32)))) = v1125
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1167 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1128+v1133<<(uint(int32(2))%32))))
	if v1125 == v1143 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	if v1133 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v1150 = v1133 + int32(1)
	if v1150 != v1130 {
		v1133 = v1150
		goto L162
	} else {
		goto L170
	}
L167:
	;
	v1225 = int32(0)
	goto L155
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _consts[711])) = v1106
	v1225 = int32(1)
	goto L155
L170:
	;
	goto L163
L171:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+624))
	if v1177 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+92))
	if v1170 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1173 = F_FindLockCycleRecurseMember(m, v1125, v1125, v1106, v1118, v1120)
	mBase = m.M
	if v1173 == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v1225 = int32(1)
	goto L155
L175:
	;
	v1225 = int32(0)
	goto L155
L176:
	;
	v1181 = v1125 + int32(620)
	if v1177 == v1181 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1183 = v1177
	goto L178
L178:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1183-int32(624))))
	if v1192 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L175
L180:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+4))
	if v1208 != v1181 {
		v1183 = v1208
		goto L178
	} else {
		goto L185
	}
L181:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1183-int32(536))))
	if v1197 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v1201 = v1183 - int32(628)
	if v1201 == v1125 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v1203 = F_FindLockCycleRecurseMember(m, v1201, v1125, v1106, v1118, v1120)
	mBase = m.M
	if v1203 == int32(0) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1225 = int32(1)
	goto L155
L185:
	;
	goto L179
L186:
	;
	v1246 = int32(3)
	goto L130
L187:
	;
	goto L188
L188:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errmsg_internal(m, int32(451398), int32(0))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(497609), int32(243), int32(318327))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v1279 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+92))
	v1281 = F_get_hash_value(m, v1279, v1280)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L193
	}
L193:
	;
	F_RemoveFromWaitQueue(m, v1277, v1281)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L194
	}
L194:
	;
	goto L127
L195:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1321+int32(25088))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1327+int32(24960))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1333+int32(24832))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1339+int32(24704))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1345+int32(24576))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1351+int32(24448))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1357+int32(24320))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1363+int32(24192))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1369+int32(24064))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1375+int32(23936))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L205
	}
L205:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1381+int32(23808))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1387+int32(23680))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1393+int32(23552))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1399+int32(23424))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1405+int32(23296))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _consts[719])) = int32(0)
	goto L110
L211:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v1465 = v235
	goto L46
L213:
	;
	v1663 = int32(*(*uint8)(unsafe.Add(mBase, _consts[779])))
	if v1663 != int32(1) {
		goto L258
	} else {
		goto L259
	}
L214:
	;
	v1485 = int32(1)
	if (v237^v1485)&v1485 != 0 {
		v1658 = v1481
		v1661 = v1484
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v1489 = int32(4432864)
	v1490 = *(*int32)(unsafe.Add(mBase, _consts[713]))
	v1492 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[713])) = v1492
	v1495 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1499 = F_LWLockAcquire(m, v1495+int32(512), v1492)
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L216
	}
L216:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+12))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+48))
	v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1503+v1504))))
	v1507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+15)))
	v1508 = *(*int64)(unsafe.Add(mBase, uint32(v236)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v220)+360)) = v1508
	v1510 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	*(*int64)(unsafe.Add(mBase, uint32(v220)+352)) = v1510
	v1513 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v1513+int32(512))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L217
	}
L217:
	;
	if v1506&int32(9) != int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v1658 = v1656
	v1661 = int32(0)
	goto L213
L219:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1490)+44))
	v1523 = int32(14)
	goto L222
L220:
	;
	if v1557 != 0 {
		goto L234
	} else {
		goto L235
	}
L221:
	;
	goto L220
L222:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	goto L225
L223:
	;
	v1542 = int32(0)
	goto L231
L225:
	;
	goto L226
L226:
	;
	goto L228
L228:
	;
	if v1530 == int32(15) {
		goto L223
	} else {
		goto L229
	}
L229:
	;
	if v1530 <= v1523 {
		v1557 = int32(1)
		goto L221
	} else {
		goto L230
	}
L230:
	;
	goto L223
L231:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v1546 != int32(2) {
		v1557 = v1542
		goto L221
	} else {
		goto L232
	}
L232:
	;
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, _consts[781])))
	if v1550 != 0 {
		v1557 = v1542
		goto L221
	} else {
		goto L233
	}
L233:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v1557 = int32(0) | base.B2i32(v1554 <= v1523)
	goto L221
L234:
	;
	F_initStringInfo(m, v220+int32(336))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1627 = F_kill(m, v1522, int32(2))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L251
	}
L237:
	;
	F_initStringInfo(m, v220+int32(320))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L238
	}
L238:
	;
	F_DescribeLockTag(m, v220+int32(336), v220+int32(352))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L239
	}
L239:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	v1575 = int32(2)
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1507<<(uint(v1575)%32))+uint32(_consts[707])))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+8))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1580+v240<<(uint(v1575)%32))))
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+272)) = v1574
	*(*int32)(unsafe.Add(mBase, uint32(v220)+276)) = v1584
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+280)) = v1587
	F_appendStringInfo(m, v220+int32(320), int32(604929), v220+int32(272))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L241
	}
L241:
	;
	v1598 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L242
	}
L242:
	;
	if v1598 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+256)) = v1522
	F_errmsg_internal(m, int32(479688), v220+int32(256))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	F_pfree(m, v1618)
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L249
	}
L246:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+240)) = v1606
	F_errdetail_log(m, int32(206113), v220+int32(240))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(499857), int32(1525), int32(237819))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L248
	}
L248:
	;
	goto L245
L249:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	F_pfree(m, v1621)
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L250
	}
L250:
	;
	goto L236
L251:
	;
	if int32(0) <= v1627 {
		goto L218
	} else {
		goto L252
	}
L252:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v1632 == int32(71) {
		goto L218
	} else {
		goto L253
	}
L253:
	;
	v1637 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L254
	}
L254:
	;
	if v1637 == int32(0) {
		goto L218
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+224)) = v1522
	F_errmsg(m, int32(295782), v220+int32(224))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(499857), int32(1547), int32(237819))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L257
	}
L257:
	;
	goto L218
L258:
	;
	if v1479 == int32(1) {
		v235 = v1465
		v237 = v1661
		goto L44
	} else {
		goto L325
	}
L259:
	;
	if v1658 == int32(0) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	F_initStringInfo(m, v220+int32(352))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_initStringInfo(m, v220+int32(336))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L262
	}
L262:
	;
	F_initStringInfo(m, v220+int32(320))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L263
	}
L263:
	;
	F_DescribeLockTag(m, v220+int32(352), v218)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L264
	}
L264:
	;
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+15)))
	v1686 = int32(2)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1685<<(uint(v1686)%32))+uint32(_consts[707])))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+8))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v240<<(uint(v1686)%32))))
	goto L265
L265:
	;
	v1697 = *(*int64)(unsafe.Add(mBase, _consts[723]))
	v1701 = m.G0
	v1702 = int32(16)
	v1703 = v1701 - v1702
	m.G0 = v1703
	F___gettimeofday(m, v1703)
	mBase = m.M
	v1706 = *(*int64)(unsafe.Add(mBase, uint32(v1703)))
	v1707 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1703)+8)))
	m.G0 = v1703 + v1702
	goto L266
L266:
	;
	v1722 = v1707 + v1706*int64(1000000) - int64(946684800000000) - v1697
	if v1722 <= int64(0) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v220)+312))
	v1739 = int32(1000)
	v1740 = base.I32_div_s(v1738, v1739)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+312)) = v1738 - v1740*v1739
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v220)+316))
	v1747 = F_LWLockAcquire(m, v241, int32(1))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L272
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220+int32(316)))) = v1734
	*(*int32)(unsafe.Add(mBase, uint32(v220+int32(312)))) = v1735
	goto L267
L269:
	;
	v1734 = int32(0)
	v1735 = int32(0)
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1726 = int64(1000000)
	v1727 = base.I64_div_u_s(v1722, v1726)
	v1734 = base.I32_wrap_i64(v1727)
	v1735 = base.I32_wrap_i64(v1722 - v1727*v1726)
	goto L268
L272:
	;
	v1751 = int32(0)
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v218)+24))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1752)+28))
	if v1753 == v1751 {
		v1839 = v1751
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1862 = v1740 + v1745*int32(1000)
	F_LWLockRelease(m, v241)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L292
	}
L274:
	;
	v1758 = v1752 + int32(24)
	if v1753 == v1758 {
		v1839 = v1751
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1763 = v1753
	v1766 = v1751
	v1771 = int32(1)
	v1774 = int32(1)
	goto L276
L276:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1763-int32(16))))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+44))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+96))
	if v1793 == v1763-int32(20) {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	v1839 = v1828
	goto L273
L278:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1763)+4))
	if v1831 != v1758 {
		v1763 = v1831
		v1766 = v1828
		v1771 = v1829
		v1774 = v1830
		goto L276
	} else {
		goto L291
	}
L279:
	;
	if v1771 != 0 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+208)) = v1792
	if v1774 != 0 {
		goto L287
	} else {
		goto L288
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+176)) = v1792
	F_appendStringInfo(m, v220+int32(336), int32(488195), v220+int32(176))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+192)) = v1792
	F_appendStringInfo(m, v220+int32(336), int32(488181), v220+int32(192))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L286
	}
L285:
	;
	v1828 = v1766
	v1829 = int32(0)
	v1830 = v1774
	goto L278
L286:
	;
	v1828 = v1766
	v1829 = int32(0)
	v1830 = v1774
	goto L278
L287:
	;
	v1820 = int32(488195)
	goto L289
L288:
	;
	v1820 = int32(488181)
	goto L289
L289:
	;
	F_appendStringInfo(m, v220+int32(320), v1820, v220+int32(208))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L290
	}
L290:
	;
	v1828 = v1766 + int32(1)
	v1829 = v1771
	v1830 = int32(0)
	goto L278
L291:
	;
	goto L277
L292:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	switch v1866 - int32(2) {
	case 0:
		goto L296
	case 1:
		goto L295
	default:
		goto L293
	}
L293:
	;
	switch v1479 {
	case 0:
		goto L307
	case 1:
		goto L308
	default:
		goto L306
	}
L294:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v220)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+160)) = v1887
	*(*int32)(unsafe.Add(mBase, uint32(v220)+148)) = v1695
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+152)) = v1890
	*(*int32)(unsafe.Add(mBase, uint32(v220)+156)) = v1862
	v1894 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+144)) = v1894
	F_errmsg(m, v1886, v220+int32(144))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L301
	}
L295:
	;
	v1879 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L299
	}
L296:
	;
	v1871 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L297
	}
L297:
	;
	if v1871 == int32(0) {
		goto L293
	} else {
		goto L298
	}
L298:
	;
	v1885 = int32(1595)
	v1886 = int32(151436)
	goto L294
L299:
	;
	if v1879 == int32(0) {
		goto L293
	} else {
		goto L300
	}
L300:
	;
	v1885 = int32(1610)
	v1886 = int32(151259)
	goto L294
L301:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+128)) = v1900
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+132)) = v1902
	F_errdetail_log_plural(m, int32(606256), int32(606302), v1839, v220+int32(128))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(499857), v1885, int32(237819))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L303
	}
L303:
	;
	goto L293
L304:
	;
	*(*int32)(unsafe.Add(mBase, _consts[720])) = int32(1)
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	F_pfree(m, v2009)
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L322
	}
L305:
	;
	F_errfinish(m, int32(499857), v2002, int32(237819))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L321
	}
L306:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	if v1970 == int32(3) {
		goto L304
	} else {
		goto L316
	}
L307:
	;
	v1950 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L313
	}
L308:
	;
	v1919 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L309
	}
L309:
	;
	if v1919 == int32(0) {
		goto L304
	} else {
		goto L310
	}
L310:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v220)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+80)) = v1923
	*(*int32)(unsafe.Add(mBase, uint32(v220)+68)) = v1695
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+72)) = v1926
	*(*int32)(unsafe.Add(mBase, uint32(v220)+76)) = v1862
	v1930 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+64)) = v1930
	F_errmsg(m, int32(151203), v220-int32(-64))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L311
	}
L311:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+48)) = v1937
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+52)) = v1939
	F_errdetail_log_plural(m, int32(606256), int32(606302), v1839, v220+int32(48))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L312
	}
L312:
	;
	v2002 = int32(1619)
	goto L305
L313:
	;
	if v1950 == int32(0) {
		goto L304
	} else {
		goto L314
	}
L314:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v220)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+112)) = v1954
	*(*int32)(unsafe.Add(mBase, uint32(v220)+100)) = v1695
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+104)) = v1957
	*(*int32)(unsafe.Add(mBase, uint32(v220)+108)) = v1862
	v1961 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+96)) = v1961
	F_errmsg(m, int32(151389), v220+int32(96))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L315
	}
L315:
	;
	v2002 = int32(1623)
	goto L305
L316:
	;
	v1975 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L317
	}
L317:
	;
	if v1975 == int32(0) {
		goto L304
	} else {
		goto L318
	}
L318:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v220)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v1979
	*(*int32)(unsafe.Add(mBase, uint32(v220)+20)) = v1695
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v220)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = v1982
	*(*int32)(unsafe.Add(mBase, uint32(v220)+28)) = v1862
	v1986 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = v1986
	F_errmsg(m, int32(151333), v220+int32(16))
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L319
	}
L319:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v1993
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v1995
	F_errdetail_log_plural(m, int32(606256), int32(606302), v1839, v220)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L320
	}
L320:
	;
	v2002 = int32(1643)
	goto L305
L321:
	;
	goto L304
L322:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v220)+320))
	F_pfree(m, v2012)
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L323
	}
L323:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v220)+336))
	F_pfree(m, v2015)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L324
	}
L324:
	;
	goto L258
L325:
	;
	goto L45
L326:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, _consts[32]))
	if v1465&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2073)) != 0 {
		goto L333
	} else {
		goto L334
	}
L327:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, _consts[721]))
	if int32(0) < v2054 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v2057 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+364)) = uint8(v2057)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+360)) = int32(2)
	v2061 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+356)) = uint8(v2061)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+352)) = v2057
	F_disable_timeouts(m, v220+int32(352))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	F_disable_timeout(m, int32(1))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L332
	}
L331:
	;
	goto L326
L332:
	;
	goto L326
L333:
	;
	v2081 = m.G0
	v2082 = int32(16)
	v2083 = v2081 - v2082
	m.G0 = v2083
	F___gettimeofday(m, v2083)
	mBase = m.M
	v2086 = *(*int64)(unsafe.Add(mBase, uint32(v2083)))
	v2087 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2083)+8)))
	m.G0 = v2083 + v2082
	goto L336
L334:
	;
	goto L335
L335:
	;
	m.G0 = v220 + int32(400)
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v231
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v230
	*(*int32)(unsafe.Add(mBase, _consts[718])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v227
	m.G0 = v225 + int32(16)
	return v1479
L336:
	;
	v2096 = int32(0)
	F_LogRecoveryConflict(m, int32(9), v244, v2087+v2086*int64(1000000)-int64(946684800000000), v2096, v2096)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		v2129 = v218
		v2130 = v219
		v2136 = v225
		v2145 = v234
		goto L5
	} else {
		goto L337
	}
L337:
	;
	goto L335
L338:
	;
	goto L4
L339:
	;
	v2163 = int32(v2159)
	m.G0 = v2145
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2163)+4))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2163)))
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2166)))
	if v2136 == v2168 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	m.ExcPending = 1
	goto L348
L341:
	;
	if v2171 != 0 {
		goto L345
	} else {
		goto L346
	}
L342:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	v2171 = v2170
	goto L344
L343:
	;
	v2171 = int32(0)
	goto L344
L344:
	;
	goto L341
L345:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+12))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+8))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2136)+4))
	v35 = v2129
	v36 = v2130
	v37 = v2165
	v38 = v2171
	v42 = v2136
	v44 = v2172
	v47 = v2173
	v48 = v2174
	v51 = v2145
	goto L1
L346:
	;
	goto L347
L347:
	;
	F___wasm_longjmp(m, v2166, v2165)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	return int32(0)
L349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_on_shmem_exit_lists_are_empty(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v2 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	if v2 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[658]))
		if v6 != 0 {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(20285), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(499754), int32(444), int32(8535))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(20327), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(499754), int32(442), int32(8535))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
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
