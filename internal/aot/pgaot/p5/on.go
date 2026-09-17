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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
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
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v209 int64
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
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
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v288 int32
	_ = v288
	var v291 int64
	_ = v291
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v318 int64
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v415 int64
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v560 int32
	_ = v560
	var v704 int32
	_ = v704
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int64
	_ = v723
	var v724 int64
	_ = v724
	var v732 int64
	_ = v732
	var v734 int32
	_ = v734
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v923 int32
	_ = v923
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v991 int32
	_ = v991
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1122 int32
	_ = v1122
	var v1129 int32
	_ = v1129
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1436 int32
	_ = v1436
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1477 int64
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1664 int64
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int64
	_ = v1673
	var v1674 int64
	_ = v1674
	var v1689 int64
	_ = v1689
	var v1693 int64
	_ = v1693
	var v1694 int64
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1769 int32
	_ = v1769
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1802 int32
	_ = v1802
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
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
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2029 int32
	_ = v2029
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2047 int64
	_ = v2047
	var v2048 int64
	_ = v2048
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2094 int32
	_ = v2094
	var v2114 int32
	_ = v2114
	var v2115 int64
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(192)
	m.G0 = v30
	v33 = l0
	v34 = l1
	v36 = int32(-1)
	v38 = v3
	v40 = v30
	v44 = v3
	v45 = v3
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
	if v36 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v2114 = int32(m.ExcTag)
	v2115 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2114 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v45
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[1])) = v33
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2]))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3]))
	goto L9
L7:
	;
	v80 = v38
	v81 = v44
	v82 = v45
	goto L8
L8:
	;
	if v80 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v74 = v40 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v40 + int32(12)
	goto L12
L10:
	;
	v80 = int32(0)
	v81 = v70
	v82 = v72
	goto L8
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v82
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v40 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v81
	v91 = int32(0)
	v92 = m.G0
	v94 = v92 - int32(400)
	m.G0 = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])))
	if v107 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v82
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v40)+188)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v40)+184)) = v82
	F_pg_re_throw(m)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L336
	}
L16:
	;
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6])) = v148
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = v148
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(v154) <= base.Ui32(int32(1)) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	if v117 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L18:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[9]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+316))
	v115 = base.B2i32(v113 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])) = uint8(v115)
	v117 = v115
	goto L20
L19:
	;
	v117 = v91
	goto L20
L20:
	;
	goto L17
L21:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[10])))
	if v121&int32(1) != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v124 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v124 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
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
	v132 = m.ExcPending
	if v132 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_WaitOnLock_0), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(_a_F_WaitOnLock_1), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_2), int32(922), int32(_a_F_WaitOnLock_3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
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
	v215 = v33
	v216 = v34
	v217 = v94
	v222 = v40
	v226 = v81
	v227 = v82
	v230 = v91
	v231 = v103
	v234 = v104
	v235 = v96&int32(15)<<(uint(int32(7))%32) + v102 + int32(_a_F_WaitOnLock_4)
	v236 = int32(1)
	v237 = v94 + int32(32)
	v239 = v209
	goto L44
L33:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[11]))
	if int32(0) < v158 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[12])))
	if v187 != int32(1) {
		v209 = int64(0)
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v182 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+112)) = v182
	v209 = int64(0)
	goto L32
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94)+352)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+384)) = v158
	*(*int64)(unsafe.Add(mBase, uint32(v94)+376)) = int64(2)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+360)) = v167
	F_enable_timeouts(m, v94+int32(352), int32(2))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	F_enable_timeout_after(m, int32(1), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		v2087 = v33
		v2088 = v34
		v2094 = v40
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
	v193 = m.G0
	v194 = int32(16)
	v195 = v193 - v194
	m.G0 = v195
	F_gettimeofday(m, v195)
	mBase = m.M
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	v199 = int64(*(*int32)(unsafe.Add(mBase, uint32(v195)+8)))
	m.G0 = v195 + v194
	goto L43
L43:
	;
	v209 = v199 + v198*int64(1000000) - int64(946684800000000)
	goto L32
L44:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(int32(2)) <= base.Ui32(v243) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(int32(1)) < base.Ui32(v2011) {
		goto L324
	} else {
		goto L325
	}
L46:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1449)+16))
	v1452 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v236&base.B2i32(v1452 == int32(4)) != 0 {
		goto L213
	} else {
		goto L214
	}
L47:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v215)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+304)) = v246
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v215)))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+296)) = v248
	v251 = v217 + int32(296)
	v256 = base.B2i32(v230 == int32(0)) & base.B2i32(v239 != int64(0))
	v257 = m.G0
	v259 = v257 - int32(80)
	m.G0 = v259
	v266 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v259+int32(16)))) = v266
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[17]))
	*(*uint8)(unsafe.Add(mBase, uint32(v259+int32(79)))) = uint8(base.B2i32(v269 == int32(3)))
	goto L50
L48:
	;
	goto L49
L49:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+14)))
	v764 = F_WaitLatch(m, v758, int32(33), int32(0), v761|int32(50331648))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L106
	}
L50:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+79)))
	if v273 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v301 = m.G0
	v302 = int32(16)
	v303 = v301 - v302
	m.G0 = v303
	F_gettimeofday(m, v303)
	mBase = m.M
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
	v307 = int64(*(*int32)(unsafe.Add(mBase, uint32(v303)+8)))
	m.G0 = v303 + v302
	v315 = v307 + v306*int64(1000000) - int64(946684800000000)
	goto L57
L52:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[19]))
	if v278 < int32(0) {
		v297 = int64(0)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20]))
	if v288 < int32(0) {
		v297 = int64(0)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
	v297 = v281 + base.I64_extend_i32_u(v278)*int64(1000)
	goto L51
L56:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v259)+16))
	v297 = v291 + base.I64_extend_i32_u(v288)*int64(1000)
	goto L51
L57:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v317)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v317)+112)) = v318
	if v318 == int64(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v323)+112)) = v315
	goto L60
L59:
	;
	goto L60
L60:
	;
	v326 = base.B2i32(v297 == int64(0))
	if v326|base.B2i32(v315 < v297) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+14)))
	F_ProcWaitForSignal(m, v370|int32(50331648))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L72
	}
L62:
	;
	v333 = F_GetLockConflicts(m, v251, int32(8), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v297 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+14)))
	F_ResolveRecoveryConflictWithVirtualXIDs(m, v333, int32(9), v336|int32(50331648), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L66
	}
L66:
	;
	goto L61
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v355))) = int64(4)
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v362
	F_enable_timeouts(m, v259+int32(16), v354)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L71
	}
L68:
	;
	v354 = int32(1)
	v355 = v259 + int32(16)
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v259)+32)) = v297
	*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = int64(4294967302)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22])) = int32(0)
	v354 = int32(2)
	v355 = v259 + int32(40)
	goto L67
L71:
	;
	goto L61
L72:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22]))
	if v376 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v560 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[23])) = v560
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[24])) = v560
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[25])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[26])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[27])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[28])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[29])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[30])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[31])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[32])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[33])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[34])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[35])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[36])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[37])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[38])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[39])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[40])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[41])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[42])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[43])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[44])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[45])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[46])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[47])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[48])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[49])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[50])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[51])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[52])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[53])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[54])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[55])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[56])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[57])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[58])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[59])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[60])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[61])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[62])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[63])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[64])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[65])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[66])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[67])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[68])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[69])) = uint8(v560)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[70])) = uint8(v560)
	goto L96
L74:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21]))
	if v378 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v383 = F_GetLockConflicts(m, v251, int32(8), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v385 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v392 = v383
	goto L78
L78:
	;
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v392)))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+8)) = v415
	v418 = v259 + int32(8)
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[71]))
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v427 = F_LWLockAcquire(m, v423+int32(512), int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L80
	}
L79:
	;
	if v256 != 0 {
		goto L73
	} else {
		goto L94
	}
L80:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	if v429 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v516+int32(512))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L92
	}
L82:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[72]))
	v444 = int32(0)
	goto L83
L83:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v421+int32(36)+v444<<(uint(int32(2))%32))))
	v471 = v437 + v468*int32(640)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+52))
	if v472 != v435 {
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
	v485 = v444 + int32(1)
	if v485 != v429 {
		v444 = v485
		goto L83
	} else {
		goto L91
	}
L87:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)+56))
	if v474 != v434 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v476 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v471)+73)) = uint8(v476)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v471)+44))
	if v478 == v476 {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v482 = F_SendProcSignal(m, v478, int32(13), v435)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
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
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	if v521 != 0 {
		v392 = v392 + int32(8)
		goto L78
	} else {
		goto L93
	}
L93:
	;
	goto L79
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = int32(0)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+14)))
	F_ProcWaitForSignal(m, v527|int32(50331648))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L95
	}
L95:
	;
	goto L73
L96:
	;
	v704 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22])) = v704
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = v704
	m.G0 = v259 + int32(80)
	if v256 == v704 {
		v1436 = v230
		goto L46
	} else {
		goto L97
	}
L97:
	;
	v718 = m.G0
	v719 = int32(16)
	v720 = v718 - v719
	m.G0 = v720
	F_gettimeofday(m, v720)
	mBase = m.M
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v720)))
	v724 = int64(*(*int32)(unsafe.Add(mBase, uint32(v720)+8)))
	m.G0 = v720 + v719
	v732 = v724 + v723*int64(1000000) - int64(946684800000000)
	goto L98
L98:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	goto L99
L99:
	;
	if base.B2i32(base.I64_extend_i32_s(v734)*int64(1000) <= v732-v239) == int32(0) {
		v1436 = int32(0)
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v747 = F_GetLockConflicts(m, v215, int32(8), v217+int32(352))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v749 = int32(0)
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	if v749 < v750 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v753 = v747
	goto L104
L103:
	;
	v753 = v749
	goto L104
L104:
	;
	F_LogRecoveryConflict(m, int32(9), v239, v732, v753, int32(1))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v1436 = int32(1)
	goto L46
L106:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = int32(0)
	goto L107
L107:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6]))
	if v771 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v773 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v777 = F_LWLockAcquire(m, v773+int32(_a_F_WaitOnLock_4), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73]))
	if v1416 == int32(0) {
		v1436 = v230
		goto L46
	} else {
		goto L211
	}
L111:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v784 = F_LWLockAcquire(m, v780+int32(_a_F_WaitOnLock_5), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v791 = F_LWLockAcquire(m, v787+int32(_a_F_WaitOnLock_6), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v794 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v798 = F_LWLockAcquire(m, v794+int32(_a_F_WaitOnLock_7), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v805 = F_LWLockAcquire(m, v801+int32(_a_F_WaitOnLock_8), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v812 = F_LWLockAcquire(m, v808+int32(_a_F_WaitOnLock_9), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v815 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v819 = F_LWLockAcquire(m, v815+int32(_a_F_WaitOnLock_10), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v826 = F_LWLockAcquire(m, v822+int32(_a_F_WaitOnLock_11), int32(0))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v833 = F_LWLockAcquire(m, v829+int32(_a_F_WaitOnLock_12), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v840 = F_LWLockAcquire(m, v836+int32(_a_F_WaitOnLock_13), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v847 = F_LWLockAcquire(m, v843+int32(_a_F_WaitOnLock_14), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L121
	}
L121:
	;
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v854 = F_LWLockAcquire(m, v850+int32(_a_F_WaitOnLock_15), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v861 = F_LWLockAcquire(m, v857+int32(_a_F_WaitOnLock_16), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v868 = F_LWLockAcquire(m, v864+int32(_a_F_WaitOnLock_17), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v875 = F_LWLockAcquire(m, v871+int32(_a_F_WaitOnLock_18), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v882 = F_LWLockAcquire(m, v878+int32(_a_F_WaitOnLock_19), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v885 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)))
	if v886 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1290+int32(_a_F_WaitOnLock_19))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L195
	}
L128:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	if v889 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v892 = int32(0)
	v893 = m.G0
	v895 = v893 - int32(16)
	m.G0 = v895
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74])) = v892
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75])) = v892
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v892
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v892
	v909 = F_DeadLockCheckRecurse(m, v885)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L131
	}
L130:
	;
	m.G0 = v895 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = v1223
	if v1223 != int32(3) {
		goto L127
	} else {
		goto L192
	}
L131:
	;
	if v909 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if int32(0) < v914 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v1085 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v1085
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1085
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1085
	*(*int32)(unsafe.Add(mBase, uint32(v895)+12)) = v1085
	v1097 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[80]))
	v1099 = v895 + int32(12)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v885)+616))
	if v1103 != 0 {
		goto L156
	} else {
		goto L157
	}
L135:
	;
	v923 = v892
	goto L138
L136:
	;
	goto L137
L137:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	if v1082 != 0 {
		goto L152
	} else {
		goto L153
	}
L138:
	;
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[81]))
	v948 = v945 + v923*int32(12)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v948)+8))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
	v952 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+40)) = v952
	v955 = v951 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v951)+32)) = v955
	if v950 <= v952 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if int32(0) < v1047 {
		v1223 = int32(2)
		goto L130
	} else {
		goto L151
	}
L140:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+15)))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1038<<(uint(int32(2))%32))+uint32(_c_F_WaitOnLock[82])))
	goto L148
L141:
	;
	v966 = v955
	v971 = v952
	goto L142
L142:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v949+v971<<(uint(int32(2))%32))))
	if v966 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v951)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+36)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v951)+32)) = v955
	goto L146
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v991)+4)) = v955
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v951)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v991))) = v999
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v951)+32)) = v991
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v951)+40))
	v1004 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+40)) = v1003 + v1004
	v1008 = v971 + v1004
	if v1008 == v950 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v951)+36))
	v966 = v1010
	v971 = v1008
	goto L142
L148:
	;
	F_ProcLockWakeup(m, v1041, v951)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v1045 = v923 + int32(1)
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if v1045 < v1047 {
		v923 = v1045
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
	v1083 = int32(4)
	goto L154
L153:
	;
	v1083 = int32(1)
	goto L154
L154:
	;
	v1223 = v1083
	goto L130
L155:
	;
	if v1204 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	v1104 = v1103
	goto L158
L157:
	;
	v1104 = v885
	goto L158
L158:
	;
	v1105 = int32(0)
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[83]))
	v1109 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78]))
	if v1105 < v1109 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1112 = v1105
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1109 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1107+v1109<<(uint(int32(2))%32)))) = v1104
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+4))
	if v1146 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1107+v1112<<(uint(int32(2))%32))))
	if v1104 == v1122 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	if v1112 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v1129 = v1112 + int32(1)
	if v1129 != v1109 {
		v1112 = v1129
		goto L162
	} else {
		goto L170
	}
L167:
	;
	v1204 = int32(0)
	goto L155
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1085
	v1204 = int32(1)
	goto L155
L170:
	;
	goto L163
L171:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+624))
	if v1156 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+92))
	if v1149 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1152 = F_FindLockCycleRecurseMember(m, v1104, v1104, v1085, v1097, v1099)
	mBase = m.M
	if v1152 == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v1204 = int32(1)
	goto L155
L175:
	;
	v1204 = int32(0)
	goto L155
L176:
	;
	v1160 = v1104 + int32(620)
	if v1156 == v1160 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1162 = v1156
	goto L178
L178:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1162-int32(624))))
	if v1171 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L175
L180:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1187 != v1160 {
		v1162 = v1187
		goto L178
	} else {
		goto L185
	}
L181:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1162-int32(536))))
	if v1176 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v1180 = v1162 - int32(628)
	if v1180 == v1104 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v1182 = F_FindLockCycleRecurseMember(m, v1180, v1104, v1085, v1097, v1099)
	mBase = m.M
	if v1182 == int32(0) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1204 = int32(1)
	goto L155
L185:
	;
	goto L179
L186:
	;
	v1223 = int32(3)
	goto L130
L187:
	;
	goto L188
L188:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_20), int32(0))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_21), int32(243), int32(_a_F_WaitOnLock_22))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
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
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1256 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[84]))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+92))
	v1258 = F_get_hash_value(m, v1256, v1257)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L193
	}
L193:
	;
	F_RemoveFromWaitQueue(m, v1254, v1258)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L194
	}
L194:
	;
	goto L127
L195:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1296+int32(_a_F_WaitOnLock_18))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1302+int32(_a_F_WaitOnLock_17))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1308+int32(_a_F_WaitOnLock_16))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1314+int32(_a_F_WaitOnLock_15))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1320+int32(_a_F_WaitOnLock_14))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1326+int32(_a_F_WaitOnLock_13))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1332+int32(_a_F_WaitOnLock_12))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1338+int32(_a_F_WaitOnLock_11))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1344+int32(_a_F_WaitOnLock_10))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1350+int32(_a_F_WaitOnLock_9))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L205
	}
L205:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1356+int32(_a_F_WaitOnLock_8))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1362+int32(_a_F_WaitOnLock_7))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1368+int32(_a_F_WaitOnLock_6))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1374+int32(_a_F_WaitOnLock_5))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1380+int32(_a_F_WaitOnLock_4))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6])) = int32(0)
	goto L110
L211:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v1436 = v230
	goto L46
L213:
	;
	v1456 = int32(_a_F_WaitOnLock_23)
	v1457 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	v1459 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v1459
	v1462 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v1466 = F_LWLockAcquire(m, v1462+int32(512), v1459)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L216
	}
L214:
	;
	v1628 = v236
	v1629 = v1452
	goto L215
L215:
	;
	v1630 = int32(0)
	v1633 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[85])))
	if base.B2i32(v1629 == v1630)|base.B2i32(v1633 != int32(1)) == v1630 {
		goto L257
	} else {
		goto L258
	}
L216:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[86]))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+12))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+48))
	v1473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470+v1471))))
	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+15)))
	v1475 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+352)) = v1475
	v1477 = *(*int64)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+360)) = v1477
	v1480 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1480+int32(512))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L217
	}
L217:
	;
	if v1473&int32(9) != int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	v1628 = int32(0)
	v1629 = v1623
	goto L215
L219:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+44))
	v1490 = int32(14)
	goto L222
L220:
	;
	if v1527 != 0 {
		goto L233
	} else {
		goto L234
	}
L221:
	;
	goto L220
L222:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[87]))
	goto L225
L223:
	;
	v1510 = int32(0)
	goto L230
L225:
	;
	goto L226
L226:
	;
	if int32(0)|base.B2i32(v1497 == int32(15)) != 0 {
		goto L223
	} else {
		goto L228
	}
L228:
	;
	if v1497 <= v1490 {
		v1527 = int32(1)
		goto L221
	} else {
		goto L229
	}
L229:
	;
	goto L223
L230:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	if v1514 != int32(2) {
		v1527 = v1510
		goto L221
	} else {
		goto L231
	}
L231:
	;
	v1518 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[89])))
	if v1518&int32(1) != 0 {
		v1527 = v1510
		goto L221
	} else {
		goto L232
	}
L232:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[90]))
	v1527 = int32(0) | base.B2i32(v1524 <= v1490)
	goto L221
L233:
	;
	v1530 = v217 + int32(336)
	F_initStringInfo(m, v1530)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1592 = F_kill(m, v1489, int32(2))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L250
	}
L236:
	;
	v1534 = v217 + int32(320)
	F_initStringInfo(m, v1534)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L237
	}
L237:
	;
	F_DescribeLockTag(m, v1530, v217+int32(352))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L238
	}
L238:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	v1543 = int32(2)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1474<<(uint(v1543)%32))+uint32(_c_F_WaitOnLock[82])))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+8))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v234<<(uint(v1543)%32))))
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+272)) = v1542
	*(*int32)(unsafe.Add(mBase, uint32(v217)+276)) = v1550
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+280)) = v1553
	F_appendStringInfo(m, v1534, int32(_a_F_WaitOnLock_24), v217+int32(272))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L240
	}
L240:
	;
	v1562 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L241
	}
L241:
	;
	if v1562 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+256)) = v1489
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_25), v217+int32(256))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	F_pfree(m, v1582)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L248
	}
L245:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+240)) = v1570
	F_errdetail_log(m, int32(_a_F_WaitOnLock_26), v217+int32(240))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), int32(1525), int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	F_pfree(m, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L249
	}
L249:
	;
	goto L235
L250:
	;
	if int32(0) <= v1592 {
		goto L218
	} else {
		goto L251
	}
L251:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[92]))
	if v1597 == int32(71) {
		goto L218
	} else {
		goto L252
	}
L252:
	;
	v1602 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L253
	}
L253:
	;
	if v1602 == int32(0) {
		goto L218
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+224)) = v1489
	F_errmsg(m, int32(_a_F_WaitOnLock_29), v217+int32(224))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), int32(1547), int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L256
	}
L256:
	;
	goto L218
L257:
	;
	v1640 = v217 + int32(352)
	F_initStringInfo(m, v1640)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	if v1450 == int32(1) {
		v230 = v1436
		v236 = v1628
		goto L44
	} else {
		goto L323
	}
L260:
	;
	F_initStringInfo(m, v217+int32(336))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_initStringInfo(m, v217+int32(320))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L262
	}
L262:
	;
	F_DescribeLockTag(m, v1640, v215)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L263
	}
L263:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+15)))
	v1655 = int32(2)
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1654<<(uint(v1655)%32))+uint32(_c_F_WaitOnLock[82])))
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1657)+8))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1658+v234<<(uint(v1655)%32))))
	goto L264
L264:
	;
	v1664 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	v1668 = m.G0
	v1669 = int32(16)
	v1670 = v1668 - v1669
	m.G0 = v1670
	F_gettimeofday(m, v1670)
	mBase = m.M
	v1673 = *(*int64)(unsafe.Add(mBase, uint32(v1670)))
	v1674 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1670)+8)))
	m.G0 = v1670 + v1669
	goto L265
L265:
	;
	v1689 = v1674 + v1673*int64(1000000) - int64(946684800000000) - v1664
	if v1689 <= int64(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v217)+312))
	v1706 = int32(1000)
	v1707 = base.I32_div_s(v1705, v1706)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+312)) = v1705 - v1707*v1706
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v217)+316))
	v1714 = F_LWLockAcquire(m, v235, int32(1))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L270
	}
L267:
	;
	v1701 = int32(0)
	v1702 = int32(0)
	goto L269
L268:
	;
	v1693 = int64(1000000)
	v1694 = base.I64_div_u_s(v1689, v1693)
	v1701 = base.I32_wrap_i64(v1694)
	v1702 = base.I32_wrap_i64(v1689 - v1694*v1693)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217+int32(316)))) = v1701
	*(*int32)(unsafe.Add(mBase, uint32(v217+int32(312)))) = v1702
	goto L266
L270:
	;
	v1718 = int32(0)
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+28))
	if v1720 == v1718 {
		v1802 = v1718
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1825 = v1707 + v1712*int32(1000)
	F_LWLockRelease(m, v235)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L290
	}
L272:
	;
	v1725 = v1719 + int32(24)
	if v1720 == v1725 {
		v1802 = v1718
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1730 = v1720
	v1731 = v1718
	v1732 = int32(1)
	v1737 = int32(1)
	goto L274
L274:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1730-int32(16))))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+44))
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+96))
	if v1758 == v1730-int32(20) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v1802 = v1793
	goto L271
L276:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1730)+4))
	if v1796 != v1725 {
		v1730 = v1796
		v1731 = v1793
		v1732 = v1794
		v1737 = v1795
		goto L274
	} else {
		goto L289
	}
L277:
	;
	if v1732 != 0 {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+208)) = v1757
	if v1737 != 0 {
		goto L285
	} else {
		goto L286
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+176)) = v1757
	F_appendStringInfo(m, v217+int32(336), int32(_a_F_WaitOnLock_30), v217+int32(176))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+192)) = v1757
	F_appendStringInfo(m, v217+int32(336), int32(_a_F_WaitOnLock_31), v217+int32(192))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L284
	}
L283:
	;
	v1793 = v1731
	v1794 = int32(0)
	v1795 = v1737
	goto L276
L284:
	;
	v1793 = v1731
	v1794 = int32(0)
	v1795 = v1737
	goto L276
L285:
	;
	v1785 = int32(_a_F_WaitOnLock_30)
	goto L287
L286:
	;
	v1785 = int32(_a_F_WaitOnLock_31)
	goto L287
L287:
	;
	F_appendStringInfo(m, v217+int32(320), v1785, v217+int32(208))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L288
	}
L288:
	;
	v1793 = v1731 + int32(1)
	v1794 = v1732
	v1795 = int32(0)
	goto L276
L289:
	;
	goto L275
L290:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	switch v1829 - int32(2) {
	case 0:
		goto L294
	case 1:
		goto L293
	default:
		goto L291
	}
L291:
	;
	switch v1450 {
	case 0:
		goto L305
	case 1:
		goto L306
	default:
		goto L304
	}
L292:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v217)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+160)) = v1850
	*(*int32)(unsafe.Add(mBase, uint32(v217)+148)) = v1662
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+152)) = v1853
	*(*int32)(unsafe.Add(mBase, uint32(v217)+156)) = v1825
	v1857 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+144)) = v1857
	F_errmsg(m, v1848, v217+int32(144))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L299
	}
L293:
	;
	v1842 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L297
	}
L294:
	;
	v1834 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L295
	}
L295:
	;
	if v1834 == int32(0) {
		goto L291
	} else {
		goto L296
	}
L296:
	;
	v1848 = int32(_a_F_WaitOnLock_32)
	v1849 = int32(1595)
	goto L292
L297:
	;
	if v1842 == int32(0) {
		goto L291
	} else {
		goto L298
	}
L298:
	;
	v1848 = int32(_a_F_WaitOnLock_33)
	v1849 = int32(1610)
	goto L292
L299:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+128)) = v1863
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+132)) = v1865
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1802, v217+int32(128))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), v1849, int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L301
	}
L301:
	;
	goto L291
L302:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = int32(1)
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	F_pfree(m, v1972)
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L320
	}
L303:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), v1965, int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L319
	}
L304:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v1933 == int32(3) {
		goto L302
	} else {
		goto L314
	}
L305:
	;
	v1913 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L311
	}
L306:
	;
	v1882 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L307
	}
L307:
	;
	if v1882 == int32(0) {
		goto L302
	} else {
		goto L308
	}
L308:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v217)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+80)) = v1886
	*(*int32)(unsafe.Add(mBase, uint32(v217)+68)) = v1662
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+72)) = v1889
	*(*int32)(unsafe.Add(mBase, uint32(v217)+76)) = v1825
	v1893 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+64)) = v1893
	F_errmsg(m, int32(_a_F_WaitOnLock_36), v217-int32(-64))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L309
	}
L309:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+48)) = v1900
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+52)) = v1902
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1802, v217+int32(48))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L310
	}
L310:
	;
	v1965 = int32(1619)
	goto L303
L311:
	;
	if v1913 == int32(0) {
		goto L302
	} else {
		goto L312
	}
L312:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v217)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+112)) = v1917
	*(*int32)(unsafe.Add(mBase, uint32(v217)+100)) = v1662
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+104)) = v1920
	*(*int32)(unsafe.Add(mBase, uint32(v217)+108)) = v1825
	v1924 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+96)) = v1924
	F_errmsg(m, int32(_a_F_WaitOnLock_37), v217+int32(96))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L313
	}
L313:
	;
	v1965 = int32(1623)
	goto L303
L314:
	;
	v1938 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L315
	}
L315:
	;
	if v1938 == int32(0) {
		goto L302
	} else {
		goto L316
	}
L316:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v217)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v217)+20)) = v1662
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v217)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = v1945
	*(*int32)(unsafe.Add(mBase, uint32(v217)+28)) = v1825
	v1949 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v1949
	F_errmsg(m, int32(_a_F_WaitOnLock_38), v217+int32(16))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L317
	}
L317:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v1956
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v1958
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1802, v217)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L318
	}
L318:
	;
	v1965 = int32(1643)
	goto L303
L319:
	;
	goto L302
L320:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v217)+320))
	F_pfree(m, v1975)
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L321
	}
L321:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v217)+336))
	F_pfree(m, v1978)
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L322
	}
L322:
	;
	goto L259
L323:
	;
	goto L45
L324:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if v1436&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2034)) != 0 {
		goto L331
	} else {
		goto L332
	}
L325:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[11]))
	if int32(0) < v2015 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2018 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+364)) = uint8(v2018)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+360)) = int32(2)
	v2022 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+356)) = uint8(v2022)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+352)) = v2018
	F_disable_timeouts(m, v217+int32(352))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	F_disable_timeout(m, int32(1))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L330
	}
L329:
	;
	goto L324
L330:
	;
	goto L324
L331:
	;
	v2042 = m.G0
	v2043 = int32(16)
	v2044 = v2042 - v2043
	m.G0 = v2044
	F_gettimeofday(m, v2044)
	mBase = m.M
	v2047 = *(*int64)(unsafe.Add(mBase, uint32(v2044)))
	v2048 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2044)+8)))
	m.G0 = v2044 + v2043
	goto L334
L332:
	;
	goto L333
L333:
	;
	m.G0 = v217 + int32(400)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v227
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v226
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[1])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+184)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v222)+188)) = v226
	m.G0 = v222 + int32(192)
	return v1450
L334:
	;
	v2057 = int32(0)
	F_LogRecoveryConflict(m, int32(9), v239, v2048+v2047*int64(1000000)-int64(946684800000000), v2057, v2057)
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		v2087 = v215
		v2088 = v216
		v2094 = v222
		goto L5
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	goto L4
L337:
	;
	v2119 = int32(v2115)
	m.G0 = v2094
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2119)+4))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2119)))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2122)))
	if v2094+int32(12) == v2125 {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	m.ExcPending = 1
	goto L346
L339:
	;
	if v2129 != 0 {
		goto L343
	} else {
		goto L344
	}
L340:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v2122)+4))
	v2129 = v2127
	goto L342
L341:
	;
	v2129 = int32(0)
	goto L342
L342:
	;
	goto L339
L343:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+188))
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2094)+184))
	v33 = v2087
	v34 = v2088
	v36 = v2129
	v38 = v2121
	v40 = v2094
	v44 = v2130
	v45 = v2131
	goto L1
L344:
	;
	goto L345
L345:
	;
	F___wasm_longjmp(m, v2122, v2121)
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	return int32(0)
L347:
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
