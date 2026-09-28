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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int64
	_ = v210
	var v211 int64
	_ = v211
	var v221 int64
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int64
	_ = v253
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v310 int64
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v328 int64
	_ = v328
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v432 int64
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v479 int32
	_ = v479
	var v623 int32
	_ = v623
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v651 int64
	_ = v651
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v851 int32
	_ = v851
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1151 int32
	_ = v1151
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1194 int32
	_ = v1194
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1321 int32
	_ = v1321
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int64
	_ = v1373
	var v1375 int64
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1511 int32
	_ = v1511
	var v1516 int32
	_ = v1516
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1538 int32
	_ = v1538
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1578 int64
	_ = v1578
	var v1581 int64
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int64
	_ = v1590
	var v1591 int64
	_ = v1591
	var v1606 int64
	_ = v1606
	var v1610 int64
	_ = v1610
	var v1611 int64
	_ = v1611
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1624 int64
	_ = v1624
	var v1625 int64
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1632 int64
	_ = v1632
	var v1636 int64
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1711 int32
	_ = v1711
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1740 int32
	_ = v1740
	var v1749 int32
	_ = v1749
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1774 int32
	_ = v1774
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1881 int32
	_ = v1881
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1907 int32
	_ = v1907
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1970 int32
	_ = v1970
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2016 int64
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2040 int32
	_ = v2040
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2058 int64
	_ = v2058
	var v2059 int64
	_ = v2059
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2085 int32
	_ = v2085
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2107 int32
	_ = v2107
	var v2129 int32
	_ = v2129
	var v2130 int64
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	v3 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(192)
	m.G0 = v31
	v34 = l0
	v35 = l1
	v37 = v3
	v40 = v31
	v43 = int32(-1)
	v46 = v3
	v47 = v3
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
	if v43 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v2129 = int32(m.ExcTag)
	v2130 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2129 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+180)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v40)+176)) = int32(1216)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v47
	v69 = int32(_a_F_WaitOnLock_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v40 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+172)) = v70
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[1])) = v35
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v34
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3]))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0]))
	goto L9
L7:
	;
	v92 = v37
	v93 = v46
	v94 = v47
	goto L8
L8:
	;
	if v92 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v86 = v40 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v40 + int32(12)
	goto L12
L10:
	;
	v92 = int32(0)
	v93 = v84
	v94 = v82
	goto L8
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v93
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v40 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v94
	v103 = int32(0)
	v106 = m.G0
	v108 = v106 - int32(400)
	m.G0 = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])))
	if v121 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v93
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v93
	F_pg_re_throw(m)
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L341
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6])) = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if base.Ui32(v165) <= base.Ui32(int32(1)) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	if v131 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L18:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+308))
	v129 = base.B2i32(v127 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])) = uint8(v129)
	v131 = v129
	goto L20
L19:
	;
	v131 = v103
	goto L20
L20:
	;
	goto L17
L21:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[9])))
	if v135&int32(1) != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v138 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v138 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
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
	v146 = m.ExcPending
	if v146 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_WaitOnLock_1), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v153 = F_errdetail(m, int32(_a_F_WaitOnLock_2), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_3), int32(924), int32(_a_F_WaitOnLock_4))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
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
	v227 = v34
	v228 = v35
	v229 = v108
	v233 = v40
	v234 = v103
	v236 = v103
	v239 = v93
	v240 = v94
	v242 = v103
	v244 = v117
	v246 = v118
	v247 = v110&int32(15)<<(uint(int32(7))%32) + v116 + int32(_a_F_WaitOnLock_5)
	v248 = int32(1)
	v249 = v108 + int32(32)
	v253 = v221
	goto L44
L33:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[10]))
	if int32(0) < v169 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[11])))
	if v199 != int32(1) {
		v221 = int64(0)
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v193 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v195 = base.AtomicRmwXchg64(m, v191, int32(408), v193)
	v221 = int64(0)
	goto L32
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108)+352)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+384)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v108)+376)) = int64(2)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+360)) = v178
	F_enable_timeouts(m, v108+int32(352), int32(2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	F_enable_timeout_after(m, int32(1), v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		v2101 = v34
		v2102 = v35
		v2107 = v40
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
	v205 = m.G0
	v206 = int32(16)
	v207 = v205 - v206
	m.G0 = v207
	F_gettimeofday(m, v207)
	mBase = m.M
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	v211 = int64(*(*int32)(unsafe.Add(mBase, uint32(v207)+8)))
	m.G0 = v207 + v206
	goto L43
L43:
	;
	v221 = v211 + v210*int64(1000000) - int64(946684800000000)
	goto L32
L44:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if base.Ui32(int32(2)) <= base.Ui32(v256) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if base.Ui32(int32(1)) < base.Ui32(v2022) {
		goto L329
	} else {
		goto L330
	}
L46:
	;
	if v1998 == int32(1) {
		v227 = v1990
		v228 = v1991
		v229 = v1992
		v233 = v1996
		v234 = v2018
		v236 = v1999
		v239 = v2002
		v240 = v2003
		v242 = v2005
		v244 = v2007
		v246 = v2009
		v247 = v2010
		v248 = v2011
		v249 = v2012
		v253 = v2016
		goto L44
	} else {
		goto L328
	}
L47:
	;
	v1581 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1585 = m.G0
	v1586 = int32(16)
	v1587 = v1585 - v1586
	m.G0 = v1587
	F_gettimeofday(m, v1587)
	mBase = m.M
	v1590 = *(*int64)(unsafe.Add(mBase, uint32(v1587)))
	v1591 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1587)+8)))
	m.G0 = v1587 + v1586
	goto L248
L48:
	;
	if v1530 != 0 {
		v1552 = v227
		v1553 = v228
		v1554 = v229
		v1558 = v233
		v1559 = v1530
		v1560 = v1531
		v1564 = v239
		v1565 = v240
		v1567 = v1538
		v1569 = v244
		v1571 = v246
		v1572 = v247
		v1573 = v248
		v1574 = v249
		v1578 = v253
		goto L47
	} else {
		goto L247
	}
L49:
	;
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v227)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+232)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v227)))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+224)) = v261
	v264 = v229 + int32(224)
	v269 = base.B2i32(v242 == int32(0)) & base.B2i32(v253 != int64(0))
	v270 = m.G0
	v272 = v270 - int32(80)
	m.G0 = v272
	v279 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v272+int32(16)))) = v279
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[16]))
	*(*uint8)(unsafe.Add(mBase, uint32(v272+int32(79)))) = uint8(base.B2i32(v282 == int32(3)))
	goto L52
L50:
	;
	goto L51
L51:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[17]))
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+14)))
	v688 = F_WaitLatch(m, v682, int32(33), int32(0), v685|int32(50331648))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L97
	}
L52:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+79)))
	if v286 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v314 = m.G0
	v315 = int32(16)
	v316 = v314 - v315
	m.G0 = v316
	F_gettimeofday(m, v316)
	mBase = m.M
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v316)))
	v320 = int64(*(*int32)(unsafe.Add(mBase, uint32(v316)+8)))
	m.G0 = v316 + v315
	v328 = v320 + v319*int64(1000000) - int64(946684800000000)
	goto L59
L54:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	if v291 < int32(0) {
		v310 = int64(0)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[19]))
	if v301 < int32(0) {
		v310 = int64(0)
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v272)+16))
	v310 = v294 + base.I64_extend_i32_u(v291)*int64(1000)
	goto L53
L58:
	;
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v272)+16))
	v310 = v304 + base.I64_extend_i32_u(v301)*int64(1000)
	goto L53
L59:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v331 = int64(0)
	v334 = base.AtomicRmwCmpxchg64(m, v330, int32(408), v331, v331)
	if v334 == v331 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v340 = base.AtomicRmwXchg64(m, v338, int32(408), v328)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v342 = base.B2i32(v310 == int64(0))
	if v342|base.B2i32(v328 < v310) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+14)))
	F_ProcWaitForSignal(m, v386|int32(50331648))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L74
	}
L64:
	;
	v349 = F_GetLockConflicts(m, v264, int32(8), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v310 == int64(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+14)))
	F_ResolveRecoveryConflictWithVirtualXIDs(m, v349, int32(2), v352|int32(50331648), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L68
	}
L68:
	;
	goto L63
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v371))) = int64(4)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+8)) = v378
	F_enable_timeouts(m, v272+int32(16), v370)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L73
	}
L70:
	;
	v370 = int32(1)
	v371 = v272 + int32(16)
	goto L69
L71:
	;
	goto L72
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v272)+32)) = v310
	*(*int64)(unsafe.Add(mBase, uint32(v272)+16)) = int64(4294967302)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = int32(0)
	v370 = int32(2)
	v371 = v272 + int32(40)
	goto L69
L73:
	;
	goto L63
L74:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21]))
	if v392 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22])) = v479
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[23])) = v479
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[24])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[25])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[26])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[27])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[28])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[29])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[30])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[31])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[32])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[33])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[34])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[35])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[36])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[37])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[38])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[39])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[40])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[41])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[42])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[43])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[44])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[45])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[46])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[47])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[48])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[49])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[50])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[51])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[52])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[53])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[54])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[55])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[56])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[57])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[58])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[59])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[60])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[61])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[62])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[63])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[64])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[65])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[66])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[67])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[68])) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[69])) = uint8(v479)
	goto L86
L76:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20]))
	if v394 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v399 = F_GetLockConflicts(m, v264, int32(8), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if v401 == int32(0) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v408 = v399
	goto L80
L80:
	;
	v432 = *(*int64)(unsafe.Add(mBase, uint32(v408)))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v432
	v437 = F_SignalRecoveryConflictWithVirtualXID(m, v272+int32(8), int32(6))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L82
	}
L81:
	;
	if v269 != 0 {
		goto L75
	} else {
		goto L84
	}
L82:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v408)+12))
	if v439 != 0 {
		v408 = v408 + int32(8)
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20])) = int32(0)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+14)))
	F_ProcWaitForSignal(m, v445|int32(50331648))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L75
L86:
	;
	v623 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = v623
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20])) = v623
	m.G0 = v272 + int32(80)
	if v269 == v623 {
		v676 = v242
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v679)+416))
	v1530 = v234
	v1531 = v680
	v1538 = v676
	goto L48
L88:
	;
	v637 = m.G0
	v638 = int32(16)
	v639 = v637 - v638
	m.G0 = v639
	F_gettimeofday(m, v639)
	mBase = m.M
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v639)))
	v643 = int64(*(*int32)(unsafe.Add(mBase, uint32(v639)+8)))
	m.G0 = v639 + v638
	v651 = v643 + v642*int64(1000000) - int64(946684800000000)
	goto L89
L89:
	;
	v653 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	goto L90
L90:
	;
	if base.B2i32(base.I64_extend_i32_s(v653)*int64(1000) <= v651-v253) == int32(0) {
		v676 = int32(0)
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v666 = F_GetLockConflicts(m, v227, int32(8), v229+int32(352))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v668 = int32(0)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v229)+352))
	if v668 < v669 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v672 = v666
	goto L95
L94:
	;
	v672 = v668
	goto L95
L95:
	;
	F_LogRecoveryConflict(m, int32(2), v253, v651, v672, int32(1))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v676 = int32(1)
	goto L87
L97:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[17]))
	v692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v691))) = v692
	v697 = base.AtomicRmwOr32(m, v692, int32(_a_F_WaitOnLock_6), v692)
	goto L98
L98:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6]))
	if v699 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v705 = F_LWLockAcquire(m, v701+int32(_a_F_WaitOnLock_5), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L102
	}
L100:
	;
	v1321 = v234
	goto L101
L101:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[70]))
	if v1343 != 0 {
		goto L202
	} else {
		goto L203
	}
L102:
	;
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v712 = F_LWLockAcquire(m, v708+int32(_a_F_WaitOnLock_7), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v719 = F_LWLockAcquire(m, v715+int32(_a_F_WaitOnLock_8), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v726 = F_LWLockAcquire(m, v722+int32(_a_F_WaitOnLock_9), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v733 = F_LWLockAcquire(m, v729+int32(_a_F_WaitOnLock_10), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v740 = F_LWLockAcquire(m, v736+int32(_a_F_WaitOnLock_11), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v747 = F_LWLockAcquire(m, v743+int32(_a_F_WaitOnLock_12), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v754 = F_LWLockAcquire(m, v750+int32(_a_F_WaitOnLock_13), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v761 = F_LWLockAcquire(m, v757+int32(_a_F_WaitOnLock_14), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v768 = F_LWLockAcquire(m, v764+int32(_a_F_WaitOnLock_15), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v775 = F_LWLockAcquire(m, v771+int32(_a_F_WaitOnLock_16), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v782 = F_LWLockAcquire(m, v778+int32(_a_F_WaitOnLock_17), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v785 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v789 = F_LWLockAcquire(m, v785+int32(_a_F_WaitOnLock_18), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v796 = F_LWLockAcquire(m, v792+int32(_a_F_WaitOnLock_19), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v803 = F_LWLockAcquire(m, v799+int32(_a_F_WaitOnLock_20), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v810 = F_LWLockAcquire(m, v806+int32(_a_F_WaitOnLock_21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)+392))
	if v814 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1216+int32(_a_F_WaitOnLock_21))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L186
	}
L119:
	;
	v1194 = int32(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v818 = int32(0)
	v819 = m.G0
	v821 = v819 - int32(16)
	m.G0 = v821
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[71])) = v818
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[72])) = v818
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73])) = v818
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74])) = v818
	v835 = F_DeadLockCheckRecurse(m, v813)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L123
	}
L122:
	;
	m.G0 = v821 + int32(16)
	if v1151 != int32(3) {
		v1194 = v1151
		goto L118
	} else {
		goto L183
	}
L123:
	;
	if v835 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73]))
	if int32(0) < v840 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v1017 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73])) = v1017
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75])) = v1017
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v821)+12)) = v1017
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	v1031 = v821 + int32(12)
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v813)+364))
	if v1035 != 0 {
		goto L148
	} else {
		goto L149
	}
L127:
	;
	v851 = v818
	goto L130
L128:
	;
	goto L129
L129:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74]))
	if v1014 != 0 {
		goto L144
	} else {
		goto L145
	}
L130:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78]))
	v875 = v872 + v851*int32(12)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v875)+8))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v879 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v878)+40)) = v879
	v882 = v878 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v878)+36)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v878)+32)) = v882
	if v877 <= v879 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if int32(0) < v978 {
		v1151 = int32(2)
		goto L122
	} else {
		goto L143
	}
L132:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878)+15)))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v969<<(uint(int32(2))%32))+uint32(_c_F_WaitOnLock[79])))
	goto L140
L133:
	;
	v892 = v882
	v895 = v879
	goto L134
L134:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v876+v895<<(uint(int32(2))%32))))
	v921 = v919 + int32(388)
	if v892 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v878)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v878)+36)) = v882
	*(*int32)(unsafe.Add(mBase, uint32(v878)+32)) = v882
	goto L138
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v919)+392)) = v882
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v878)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v919)+388)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v929)+4)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v878)+32)) = v921
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v878)+40))
	v934 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v878)+40)) = v933 + v934
	v938 = v895 + v934
	if v938 == v877 {
		goto L132
	} else {
		goto L139
	}
L139:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v878)+36))
	v892 = v940
	v895 = v938
	goto L134
L140:
	;
	F_ProcLockWakeup(m, v972, v878)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L141
	}
L141:
	;
	v976 = v851 + int32(1)
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73]))
	if v976 < v978 {
		v851 = v976
		goto L130
	} else {
		goto L142
	}
L142:
	;
	goto L131
L143:
	;
	goto L129
L144:
	;
	v1015 = int32(4)
	goto L146
L145:
	;
	v1015 = int32(1)
	goto L146
L146:
	;
	v1151 = v1015
	goto L122
L147:
	;
	if v1129 != 0 {
		goto L177
	} else {
		goto L178
	}
L148:
	;
	v1036 = v1035
	goto L150
L149:
	;
	v1036 = v813
	goto L150
L150:
	;
	v1037 = int32(0)
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[80]))
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75]))
	if v1037 < v1041 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1044 = v1037
	goto L154
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75])) = v1041 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1039+v1041<<(uint(int32(2))%32)))) = v1036
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+392))
	if v1078 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L154:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1039+v1044<<(uint(int32(2))%32))))
	if v1036 == v1054 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L153
L156:
	;
	if v1044 != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	v1061 = v1044 + int32(1)
	if v1061 != v1041 {
		v1044 = v1061
		goto L154
	} else {
		goto L162
	}
L159:
	;
	v1129 = int32(0)
	goto L147
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v1017
	v1129 = int32(1)
	goto L147
L162:
	;
	goto L155
L163:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+372))
	if v1085 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v1081 = F_FindLockCycleRecurseMember(m, v1036, v1036, v1017, v1029, v1031)
	mBase = m.M
	if v1081 == int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v1129 = int32(1)
	goto L147
L166:
	;
	v1129 = int32(0)
	goto L147
L167:
	;
	v1089 = v1036 + int32(368)
	if v1085 == v1089 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v1091 = v1085
	goto L169
L169:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+16))
	if v1098 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L166
L171:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+4))
	if v1112 != v1089 {
		v1091 = v1112
		goto L169
	} else {
		goto L176
	}
L172:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+8))
	if v1101 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1105 = v1091 - int32(376)
	if v1105 == v1036 {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v1107 = F_FindLockCycleRecurseMember(m, v1105, v1036, v1017, v1029, v1031)
	mBase = m.M
	if v1107 == int32(0) {
		goto L171
	} else {
		goto L175
	}
L175:
	;
	v1129 = int32(1)
	goto L147
L176:
	;
	goto L170
L177:
	;
	v1151 = int32(3)
	goto L122
L178:
	;
	goto L179
L179:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L180
	}
L180:
	;
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_22), int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_23), int32(243), int32(_a_F_WaitOnLock_24))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v1180 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[81]))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1178)+384))
	v1182 = F_get_hash_value(m, v1180, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L184
	}
L184:
	;
	F_RemoveFromWaitQueue(m, v1178, v1182)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L185
	}
L185:
	;
	v1194 = int32(3)
	goto L118
L186:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1222+int32(_a_F_WaitOnLock_20))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L187
	}
L187:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1228+int32(_a_F_WaitOnLock_19))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L188
	}
L188:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1234+int32(_a_F_WaitOnLock_18))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L189
	}
L189:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1240+int32(_a_F_WaitOnLock_17))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L190
	}
L190:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1246+int32(_a_F_WaitOnLock_16))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L191
	}
L191:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1252+int32(_a_F_WaitOnLock_15))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1258+int32(_a_F_WaitOnLock_14))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L193
	}
L193:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1264+int32(_a_F_WaitOnLock_13))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L194
	}
L194:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1270+int32(_a_F_WaitOnLock_12))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L195
	}
L195:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1276+int32(_a_F_WaitOnLock_11))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1282+int32(_a_F_WaitOnLock_10))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1288+int32(_a_F_WaitOnLock_9))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1294+int32(_a_F_WaitOnLock_8))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1300+int32(_a_F_WaitOnLock_7))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1306+int32(_a_F_WaitOnLock_5))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6])) = int32(0)
	v1321 = v1194
	goto L101
L202:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[12]))
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+416))
	if base.B2i32(v1321 == int32(4))&v248 == int32(0) {
		v1530 = v1321
		v1531 = v1348
		v1538 = v242
		goto L48
	} else {
		goto L206
	}
L205:
	;
	goto L204
L206:
	;
	v1354 = int32(_a_F_WaitOnLock_25)
	v1355 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74]))
	v1357 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74])) = v1357
	v1360 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v1364 = F_LWLockAcquire(m, v1360+int32(512), v1357)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[82]))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+32))
	v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368+v1369))))
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+15)))
	v1373 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+352)) = v1373
	v1375 = *(*int64)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+360)) = v1375
	v1378 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1378+int32(512))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L208
	}
L208:
	;
	if v1371&int32(9) != int32(1) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1552 = v227
	v1553 = v228
	v1554 = v229
	v1558 = v233
	v1559 = int32(4)
	v1560 = v1348
	v1564 = v239
	v1565 = v240
	v1567 = v242
	v1569 = v244
	v1571 = v246
	v1572 = v247
	v1573 = int32(0)
	v1574 = v249
	v1578 = v253
	goto L47
L210:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+12))
	v1388 = int32(14)
	goto L213
L211:
	;
	if v1428 != 0 {
		goto L224
	} else {
		goto L225
	}
L212:
	;
	goto L211
L213:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[83]))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1395<<(uint(int32(2))%32))+uint32(_c_F_WaitOnLock[84])))
	goto L216
L214:
	;
	v1411 = int32(0)
	goto L221
L216:
	;
	goto L217
L217:
	;
	if int32(0)|base.B2i32(v1398 == int32(15)) != 0 {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	if v1398 <= v1388 {
		v1428 = int32(1)
		goto L212
	} else {
		goto L220
	}
L220:
	;
	goto L214
L221:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[85]))
	if v1415 != int32(2) {
		v1428 = v1411
		goto L212
	} else {
		goto L222
	}
L222:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[86])))
	if v1419&int32(1) != 0 {
		v1428 = v1411
		goto L212
	} else {
		goto L223
	}
L223:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[87]))
	v1428 = int32(0) | base.B2i32(v1425 <= v1388)
	goto L212
L224:
	;
	v1431 = v229 + int32(336)
	F_initStringInfo(m, v1431)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1493 = F_pgmem_kill(m, v1387, int32(2))
	mBase = m.M
	if int32(0) <= v1493 {
		goto L209
	} else {
		goto L241
	}
L227:
	;
	v1435 = v229 + int32(320)
	F_initStringInfo(m, v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L228
	}
L228:
	;
	F_DescribeLockTag(m, v1431, v229+int32(352))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L229
	}
L229:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	v1444 = int32(2)
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1372<<(uint(v1444)%32))+uint32(_c_F_WaitOnLock[79])))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+8))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1447+v246<<(uint(v1444)%32))))
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+288)) = v1443
	*(*int32)(unsafe.Add(mBase, uint32(v229)+292)) = v1451
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v229)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+296)) = v1454
	F_appendStringInfo(m, v1435, int32(_a_F_WaitOnLock_26), v229+int32(288))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L231
	}
L231:
	;
	v1463 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L232
	}
L232:
	;
	if v1463 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+272)) = v1387
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_27), v229+int32(272))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v229)+336))
	F_pfree(m, v1483)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L239
	}
L236:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v229)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+256)) = v1471
	F_errdetail_log(m, int32(_a_F_WaitOnLock_28), v229+int32(256))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_29), int32(1560), int32(_a_F_WaitOnLock_30))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v229)+320))
	F_pfree(m, v1486)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L240
	}
L240:
	;
	goto L226
L241:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[89]))
	if v1497 == int32(71) {
		goto L209
	} else {
		goto L242
	}
L242:
	;
	v1502 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L243
	}
L243:
	;
	if v1502 == int32(0) {
		goto L209
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+240)) = v1387
	F_errmsg(m, int32(_a_F_WaitOnLock_31), v229+int32(240))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_29), int32(1582), int32(_a_F_WaitOnLock_30))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		v2101 = v227
		v2102 = v228
		v2107 = v233
		goto L5
	} else {
		goto L246
	}
L246:
	;
	goto L209
L247:
	;
	v1990 = v227
	v1991 = v228
	v1992 = v229
	v1996 = v233
	v1998 = v1531
	v1999 = v236
	v2002 = v239
	v2003 = v240
	v2005 = v1538
	v2007 = v244
	v2009 = v246
	v2010 = v247
	v2011 = v248
	v2012 = v249
	v2016 = v253
	v2018 = int32(0)
	goto L46
L248:
	;
	v1606 = v1591 + v1590*int64(1000000) - int64(946684800000000) - v1581
	if v1606 <= int64(0) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	if v1560 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	v1618 = int32(0)
	v1619 = int32(0)
	goto L252
L251:
	;
	v1610 = int64(1000000)
	v1611 = base.I64_div_u_s(v1606, v1610)
	v1618 = base.I32_wrap_i64(v1611)
	v1619 = base.I32_wrap_i64(v1606 - v1611*v1610)
	goto L252
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1554+int32(316)))) = v1618
	*(*int32)(unsafe.Add(mBase, uint32(v1554+int32(312)))) = v1619
	goto L249
L253:
	;
	v1624 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1554)+312)))
	v1625 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1554)+316)))
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552)+14)))
	v1631 = v1629 * int32(24)
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v1631)+uint32(_c_F_WaitOnLock[90])))
	*(*int64)(unsafe.Add(mBase, uint32(v1631)+uint32(_c_F_WaitOnLock[90]))) = v1632 + int64(1)
	v1636 = *(*int64)(unsafe.Add(mBase, uint32(v1631)+uint32(_c_F_WaitOnLock[91])))
	*(*int64)(unsafe.Add(mBase, uint32(v1631)+uint32(_c_F_WaitOnLock[91]))) = v1636 + (v1624 + v1625*int64(1000000))
	v1640 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[92])) = uint8(v1640)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[93])) = uint8(v1640)
	goto L255
L254:
	;
	goto L255
L255:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+312))
	v1648 = int32(1000)
	v1649 = base.I32_div_s(v1647, v1648)
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+312)) = v1647 - v1649*v1648
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[94])))
	if v1655 == int32(1) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+316))
	v1662 = v1554 + int32(352)
	F_initStringInfo(m, v1662)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L259
	}
L257:
	;
	v1970 = v236
	goto L258
L258:
	;
	v1990 = v1552
	v1991 = v1553
	v1992 = v1554
	v1996 = v1558
	v1998 = v1560
	v1999 = v1970
	v2002 = v1564
	v2003 = v1565
	v2005 = v1567
	v2007 = v1569
	v2009 = v1571
	v2010 = v1572
	v2011 = v1573
	v2012 = v1574
	v2016 = v1578
	v2018 = int32(1)
	goto L46
L259:
	;
	F_initStringInfo(m, v1554+int32(336))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_initStringInfo(m, v1554+int32(320))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_DescribeLockTag(m, v1662, v1552)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L262
	}
L262:
	;
	v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552)+15)))
	v1677 = int32(2)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1676<<(uint(v1677)%32))+uint32(_c_F_WaitOnLock[79])))
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1679)+8))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1680+v1571<<(uint(v1677)%32))))
	goto L263
L263:
	;
	v1686 = F_LWLockAcquire(m, v1572, int32(1))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L264
	}
L264:
	;
	v1688 = int32(0)
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+24))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+28))
	if v1690 == v1688 {
		v1774 = v1688
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1797 = v1649 + v1658*int32(1000)
	F_LWLockRelease(m, v1572)
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L284
	}
L266:
	;
	v1695 = v1689 + int32(24)
	if v1690 == v1695 {
		v1774 = v1688
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v1700 = v1690
	v1701 = int32(1)
	v1702 = v1688
	v1711 = int32(1)
	goto L268
L268:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1700-int32(16))))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1727)+12))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1727)+396))
	if v1729 == v1700-int32(20) {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	v1774 = v1765
	goto L265
L270:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1700)+4))
	if v1767 != v1695 {
		v1700 = v1767
		v1701 = v1764
		v1702 = v1765
		v1711 = v1766
		goto L268
	} else {
		goto L283
	}
L271:
	;
	if v1701 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+208)) = v1728
	if v1711 != 0 {
		goto L279
	} else {
		goto L280
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+176)) = v1728
	F_appendStringInfo(m, v1554+int32(336), int32(_a_F_WaitOnLock_32), v1554+int32(176))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+192)) = v1728
	F_appendStringInfo(m, v1554+int32(336), int32(_a_F_WaitOnLock_33), v1554+int32(192))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L278
	}
L277:
	;
	v1764 = int32(0)
	v1765 = v1702
	v1766 = v1711
	goto L270
L278:
	;
	v1764 = int32(0)
	v1765 = v1702
	v1766 = v1711
	goto L270
L279:
	;
	v1756 = int32(_a_F_WaitOnLock_32)
	goto L281
L280:
	;
	v1756 = int32(_a_F_WaitOnLock_33)
	goto L281
L281:
	;
	F_appendStringInfo(m, v1554+int32(320), v1756, v1554+int32(208))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L282
	}
L282:
	;
	v1764 = v1701
	v1765 = v1702 + int32(1)
	v1766 = int32(0)
	goto L270
L283:
	;
	goto L269
L284:
	;
	switch v1559 - int32(2) {
	case 0:
		goto L288
	case 1:
		goto L287
	default:
		goto L285
	}
L285:
	;
	if v1560 == int32(1) {
		goto L298
	} else {
		goto L299
	}
L286:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+160)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+148)) = v1684
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+152)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+156)) = v1797
	v1827 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+144)) = v1827
	F_errmsg(m, v1818, v1554+int32(144))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L293
	}
L287:
	;
	v1812 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L291
	}
L288:
	;
	v1804 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L289
	}
L289:
	;
	if v1804 == int32(0) {
		goto L285
	} else {
		goto L290
	}
L290:
	;
	v1818 = int32(_a_F_WaitOnLock_34)
	v1819 = int32(1641)
	goto L286
L291:
	;
	if v1812 == int32(0) {
		goto L285
	} else {
		goto L292
	}
L292:
	;
	v1818 = int32(_a_F_WaitOnLock_35)
	v1819 = int32(1656)
	goto L286
L293:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+128)) = v1833
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+132)) = v1835
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_36), int32(_a_F_WaitOnLock_37), v1774, v1554+int32(128))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L294
	}
L294:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_29), v1819, int32(_a_F_WaitOnLock_30))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L295
	}
L295:
	;
	goto L285
L296:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+352))
	F_pfree(m, v1952)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L325
	}
L297:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_29), v1944, int32(_a_F_WaitOnLock_30))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L324
	}
L298:
	;
	v1851 = int32(1)
	if v236&v1851 != 0 {
		v1950 = v1851
		goto L296
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	if v1560 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L301:
	;
	v1856 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L302
	}
L302:
	;
	if v1856 == int32(0) {
		v1950 = v1851
		goto L296
	} else {
		goto L303
	}
L303:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v1574))) = v1860
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+20)) = v1684
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+24)) = v1863
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+28)) = v1797
	v1867 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+16)) = v1867
	F_errmsg(m, int32(_a_F_WaitOnLock_38), v1554+int32(16))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L304
	}
L304:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v1554))) = v1874
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+4)) = v1876
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_36), int32(_a_F_WaitOnLock_37), v1774, v1554)
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L305
	}
L305:
	;
	v1944 = int32(1680)
	v1945 = int32(1)
	goto L297
L306:
	;
	v1944 = v1942
	v1945 = v236
	goto L297
L307:
	;
	v1888 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	if v1559 == int32(3) {
		goto L315
	} else {
		goto L316
	}
L310:
	;
	if v1888 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1950 = v236
	goto L296
L312:
	;
	goto L313
L313:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v1554-int32(-64)))) = v1894
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+52)) = v1684
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+56)) = v1897
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+60)) = v1797
	v1901 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+48)) = v1901
	F_errmsg(m, int32(_a_F_WaitOnLock_39), v1554+int32(48))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L314
	}
L314:
	;
	v1942 = int32(1687)
	goto L306
L315:
	;
	v1950 = v236
	goto L296
L316:
	;
	goto L317
L317:
	;
	v1913 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L318
	}
L318:
	;
	if v1913 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1950 = v236
	goto L296
L320:
	;
	goto L321
L321:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+112)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+100)) = v1684
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+104)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+108)) = v1797
	v1924 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+96)) = v1924
	F_errmsg(m, int32(_a_F_WaitOnLock_40), v1554+int32(96))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L322
	}
L322:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+80)) = v1931
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+84)) = v1933
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_36), int32(_a_F_WaitOnLock_37), v1774, v1554+int32(80))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L323
	}
L323:
	;
	v1942 = int32(1707)
	goto L306
L324:
	;
	v1950 = v1945
	goto L296
L325:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+320))
	F_pfree(m, v1955)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+336))
	F_pfree(m, v1958)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		v2101 = v1552
		v2102 = v1553
		v2107 = v1558
		goto L5
	} else {
		goto L327
	}
L327:
	;
	v1970 = v1950
	goto L258
L328:
	;
	goto L45
L329:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v2005&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2045)) != 0 {
		goto L336
	} else {
		goto L337
	}
L330:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[10]))
	if int32(0) < v2026 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v2029 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1992)+364)) = uint8(v2029)
	*(*int32)(unsafe.Add(mBase, uint32(v1992)+360)) = int32(2)
	v2033 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1992)+356)) = uint8(v2033)
	*(*int32)(unsafe.Add(mBase, uint32(v1992)+352)) = v2029
	F_disable_timeouts(m, v1992+int32(352))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		v2101 = v1990
		v2102 = v1991
		v2107 = v1996
		goto L5
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	F_disable_timeout(m, int32(1))
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		v2101 = v1990
		v2102 = v1991
		v2107 = v1996
		goto L5
	} else {
		goto L335
	}
L334:
	;
	goto L329
L335:
	;
	goto L329
L336:
	;
	v2053 = m.G0
	v2054 = int32(16)
	v2055 = v2053 - v2054
	m.G0 = v2055
	F_gettimeofday(m, v2055)
	mBase = m.M
	v2058 = *(*int64)(unsafe.Add(mBase, uint32(v2055)))
	v2059 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2055)+8)))
	m.G0 = v2055 + v2054
	goto L339
L337:
	;
	goto L338
L338:
	;
	m.G0 = v1992 + int32(400)
	v2075 = int32(_a_F_WaitOnLock_0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v2002
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v2003
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1996)+184)) = v2002
	*(*int32)(unsafe.Add(mBase, uint32(v1996)+188)) = v2003
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v1996)+172))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v2085
	m.G0 = v1996 + int32(192)
	return v1998
L339:
	;
	v2068 = int32(0)
	F_LogRecoveryConflict(m, int32(2), v2016, v2059+v2058*int64(1000000)-int64(946684800000000), v2068, v2068)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		v2101 = v1990
		v2102 = v1991
		v2107 = v1996
		goto L5
	} else {
		goto L340
	}
L340:
	;
	goto L338
L341:
	;
	goto L4
L342:
	;
	v2134 = int32(v2130)
	m.G0 = v2107
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2134)+4))
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2134)))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2137)))
	if v2107+int32(12) == v2140 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	m.ExcPending = 1
	goto L351
L344:
	;
	if v2144 != 0 {
		goto L348
	} else {
		goto L349
	}
L345:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v2137)+4))
	v2144 = v2142
	goto L347
L346:
	;
	v2144 = int32(0)
	goto L347
L347:
	;
	goto L344
L348:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+188))
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+184))
	v34 = v2101
	v35 = v2102
	v37 = v2136
	v40 = v2107
	v43 = v2144
	v46 = v2146
	v47 = v2145
	goto L1
L349:
	;
	goto L350
L350:
	;
	F___wasm_longjmp(m, v2137, v2136)
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	return int32(0)
L352:
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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_check_on_shmem_exit_lists_are_empty[0]))
	if v2 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_check_on_shmem_exit_lists_are_empty[1]))
		if v6 != 0 {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_check_on_shmem_exit_lists_are_empty_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_on_shmem_exit_lists_are_empty_1), int32(444), int32(_a_F_check_on_shmem_exit_lists_are_empty_2))
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
			F_errmsg_internal(m, int32(_a_F_check_on_shmem_exit_lists_are_empty_3), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_check_on_shmem_exit_lists_are_empty_1), int32(442), int32(_a_F_check_on_shmem_exit_lists_are_empty_2))
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
