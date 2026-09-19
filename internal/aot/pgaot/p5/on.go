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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v208 int64
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v241 int32
	_ = v241
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int64
	_ = v279
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v295 int64
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v313 int64
	_ = v313
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
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
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v557 int32
	_ = v557
	var v701 int32
	_ = v701
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int64
	_ = v720
	var v721 int64
	_ = v721
	var v729 int64
	_ = v729
	var v731 int32
	_ = v731
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
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
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v986 int32
	_ = v986
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1426 int32
	_ = v1426
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int64
	_ = v1464
	var v1466 int64
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1486 int32
	_ = v1486
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1652 int64
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1661 int64
	_ = v1661
	var v1662 int64
	_ = v1662
	var v1677 int64
	_ = v1677
	var v1681 int64
	_ = v1681
	var v1682 int64
	_ = v1682
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1756 int32
	_ = v1756
	var v1765 int32
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1858 int32
	_ = v1858
	var v1862 int32
	_ = v1862
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2032 int64
	_ = v2032
	var v2033 int64
	_ = v2033
	var v2042 int32
	_ = v2042
	var v2045 int32
	_ = v2045
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2079 int32
	_ = v2079
	var v2098 int32
	_ = v2098
	var v2099 int64
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	v3 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(192)
	m.G0 = v29
	v32 = l0
	v33 = l1
	v35 = int32(-1)
	v37 = v3
	v39 = v29
	v43 = v3
	v44 = v3
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
	if v35 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v2098 = int32(m.ExcTag)
	v2099 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2098 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v44
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[0])) = v33
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[1])) = v32
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2]))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3]))
	goto L9
L7:
	;
	v78 = v37
	v79 = v43
	v80 = v44
	goto L8
L8:
	;
	if v78 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v72 = v39 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v39 + int32(12)
	goto L12
L10:
	;
	v78 = int32(0)
	v79 = v68
	v80 = v70
	goto L8
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v80
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v39 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = v79
	v89 = int32(0)
	v90 = m.G0
	v92 = v90 - int32(400)
	m.G0 = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])))
	if v105 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v80
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v39)+184)) = v80
	F_pg_re_throw(m)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L335
	}
L16:
	;
	v146 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6])) = v146
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = v146
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(v152) <= base.Ui32(int32(1)) {
		goto L33
	} else {
		goto L34
	}
L17:
	;
	if v115 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L18:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[9]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+316))
	v113 = base.B2i32(v111 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[5])) = uint8(v113)
	v115 = v113
	goto L20
L19:
	;
	v115 = v89
	goto L20
L20:
	;
	goto L17
L21:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[10])))
	if v119&int32(1) != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v122 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v122 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
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
	v130 = m.ExcPending
	if v130 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_WaitOnLock_0), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errdetail(m, int32(_a_F_WaitOnLock_1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_2), int32(922), int32(_a_F_WaitOnLock_3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
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
	v214 = v32
	v215 = v33
	v216 = v92
	v221 = v39
	v225 = v79
	v226 = v80
	v229 = v89
	v230 = v101
	v233 = v102
	v234 = v94&int32(15)<<(uint(int32(7))%32) + v100 + int32(_a_F_WaitOnLock_4)
	v235 = int32(1)
	v236 = v92 + int32(32)
	v238 = v208
	goto L44
L33:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[11]))
	if int32(0) < v156 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[12])))
	if v186 != int32(1) {
		v208 = int64(0)
		goto L32
	} else {
		goto L42
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v180 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	v182 = base.AtomicRmwXchg64(m, v178, int32(112), v180)
	v208 = int64(0)
	goto L32
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v92)+352)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+384)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v92)+376)) = int64(2)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+360)) = v165
	F_enable_timeouts(m, v92+int32(352), int32(2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	F_enable_timeout_after(m, int32(1), v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v2072 = v32
		v2073 = v33
		v2079 = v39
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
	v192 = m.G0
	v193 = int32(16)
	v194 = v192 - v193
	m.G0 = v194
	F_gettimeofday(m, v194)
	mBase = m.M
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	v198 = int64(*(*int32)(unsafe.Add(mBase, uint32(v194)+8)))
	m.G0 = v194 + v193
	goto L43
L43:
	;
	v208 = v198 + v197*int64(1000000) - int64(946684800000000)
	goto L32
L44:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(int32(2)) <= base.Ui32(v241) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(int32(1)) < base.Ui32(v1996) {
		goto L323
	} else {
		goto L324
	}
L46:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+16))
	v1441 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v235&base.B2i32(v1441 == int32(4)) != 0 {
		goto L213
	} else {
		goto L214
	}
L47:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v214)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+304)) = v244
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v214)))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+296)) = v246
	v249 = v216 + int32(296)
	v254 = base.B2i32(v229 == int32(0)) & base.B2i32(v238 != int64(0))
	v255 = m.G0
	v257 = v255 - int32(80)
	m.G0 = v257
	v264 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v257+int32(16)))) = v264
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[17]))
	*(*uint8)(unsafe.Add(mBase, uint32(v257+int32(79)))) = uint8(base.B2i32(v267 == int32(3)))
	goto L50
L48:
	;
	goto L49
L49:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+14)))
	v761 = F_WaitLatch(m, v755, int32(33), int32(0), v758|int32(50331648))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L106
	}
L50:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+79)))
	if v271 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v299 = m.G0
	v300 = int32(16)
	v301 = v299 - v300
	m.G0 = v301
	F_gettimeofday(m, v301)
	mBase = m.M
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	v305 = int64(*(*int32)(unsafe.Add(mBase, uint32(v301)+8)))
	m.G0 = v301 + v300
	v313 = v305 + v304*int64(1000000) - int64(946684800000000)
	goto L57
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[19]))
	if v276 < int32(0) {
		v295 = int64(0)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[20]))
	if v286 < int32(0) {
		v295 = int64(0)
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v257)+16))
	v295 = v279 + base.I64_extend_i32_u(v276)*int64(1000)
	goto L51
L56:
	;
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v257)+16))
	v295 = v289 + base.I64_extend_i32_u(v286)*int64(1000)
	goto L51
L57:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v316 = int64(0)
	v319 = base.AtomicRmwCmpxchg64(m, v315, int32(112), v316, v316)
	if v319 == v316 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v325 = base.AtomicRmwXchg64(m, v323, int32(112), v313)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v327 = base.B2i32(v295 == int64(0))
	if v327|base.B2i32(v313 < v295) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+14)))
	F_ProcWaitForSignal(m, v371|int32(50331648))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L72
	}
L62:
	;
	v334 = F_GetLockConflicts(m, v249, int32(8), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v295 == int64(0) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+14)))
	F_ResolveRecoveryConflictWithVirtualXIDs(m, v334, int32(9), v337|int32(50331648), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
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
	*(*int64)(unsafe.Add(mBase, uint32(v356))) = int64(4)
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v363
	F_enable_timeouts(m, v257+int32(16), v355)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L71
	}
L68:
	;
	v355 = int32(1)
	v356 = v257 + int32(16)
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v257)+32)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v257)+16)) = int64(4294967302)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22])) = int32(0)
	v355 = int32(2)
	v356 = v257 + int32(40)
	goto L67
L71:
	;
	goto L61
L72:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22]))
	if v377 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v557 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[23])) = v557
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[24])) = v557
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[25])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[26])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[27])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[28])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[29])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[30])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[31])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[32])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[33])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[34])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[35])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[36])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[37])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[38])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[39])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[40])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[41])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[42])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[43])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[44])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[45])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[46])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[47])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[48])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[49])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[50])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[51])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[52])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[53])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[54])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[55])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[56])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[57])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[58])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[59])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[60])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[61])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[62])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[63])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[64])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[65])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[66])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[67])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[68])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[69])) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[70])) = uint8(v557)
	goto L96
L74:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21]))
	if v379 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v384 = F_GetLockConflicts(m, v249, int32(8), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v386 == int32(0) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v393 = v384
	goto L78
L78:
	;
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v393)))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+8)) = v415
	v418 = v257 + int32(8)
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[71]))
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v427 = F_LWLockAcquire(m, v423+int32(512), int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L80
	}
L79:
	;
	if v254 != 0 {
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
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v514+int32(512))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
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
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v421+int32(36)+v444<<(uint(int32(2))%32))))
	v470 = v437 + v467*int32(640)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+52))
	if v471 != v435 {
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
	v484 = v444 + int32(1)
	if v484 != v429 {
		v444 = v484
		goto L83
	} else {
		goto L91
	}
L87:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v470)+56))
	if v473 != v434 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v475 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v470)+73)) = uint8(v475)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v470)+44))
	if v477 == v475 {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v481 = F_SendProcSignal(m, v477, int32(13), v435)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
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
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	if v519 != 0 {
		v393 = v393 + int32(8)
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
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+14)))
	F_ProcWaitForSignal(m, v525|int32(50331648))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L95
	}
L95:
	;
	goto L73
L96:
	;
	v701 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[22])) = v701
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[21])) = v701
	m.G0 = v257 + int32(80)
	if v254 == v701 {
		v1426 = v229
		goto L46
	} else {
		goto L97
	}
L97:
	;
	v715 = m.G0
	v716 = int32(16)
	v717 = v715 - v716
	m.G0 = v717
	F_gettimeofday(m, v717)
	mBase = m.M
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v717)))
	v721 = int64(*(*int32)(unsafe.Add(mBase, uint32(v717)+8)))
	m.G0 = v717 + v716
	v729 = v721 + v720*int64(1000000) - int64(946684800000000)
	goto L98
L98:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[15]))
	goto L99
L99:
	;
	if base.B2i32(base.I64_extend_i32_s(v731)*int64(1000) <= v729-v238) == int32(0) {
		v1426 = int32(0)
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v744 = F_GetLockConflicts(m, v214, int32(8), v216+int32(352))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v746 = int32(0)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	if v746 < v747 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v750 = v744
	goto L104
L103:
	;
	v750 = v746
	goto L104
L104:
	;
	F_LogRecoveryConflict(m, int32(9), v238, v729, v750, int32(1))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v1426 = int32(1)
	goto L46
L106:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = int32(0)
	goto L107
L107:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6]))
	if v768 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v774 = F_LWLockAcquire(m, v770+int32(_a_F_WaitOnLock_4), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73]))
	if v1406 == int32(0) {
		v1426 = v229
		goto L46
	} else {
		goto L211
	}
L111:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v781 = F_LWLockAcquire(m, v777+int32(_a_F_WaitOnLock_5), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v788 = F_LWLockAcquire(m, v784+int32(_a_F_WaitOnLock_6), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v795 = F_LWLockAcquire(m, v791+int32(_a_F_WaitOnLock_7), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v798 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v802 = F_LWLockAcquire(m, v798+int32(_a_F_WaitOnLock_8), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v805 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v809 = F_LWLockAcquire(m, v805+int32(_a_F_WaitOnLock_9), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v816 = F_LWLockAcquire(m, v812+int32(_a_F_WaitOnLock_10), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v823 = F_LWLockAcquire(m, v819+int32(_a_F_WaitOnLock_11), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v830 = F_LWLockAcquire(m, v826+int32(_a_F_WaitOnLock_12), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v837 = F_LWLockAcquire(m, v833+int32(_a_F_WaitOnLock_13), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v844 = F_LWLockAcquire(m, v840+int32(_a_F_WaitOnLock_14), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L121
	}
L121:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v851 = F_LWLockAcquire(m, v847+int32(_a_F_WaitOnLock_15), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v858 = F_LWLockAcquire(m, v854+int32(_a_F_WaitOnLock_16), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v865 = F_LWLockAcquire(m, v861+int32(_a_F_WaitOnLock_17), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v872 = F_LWLockAcquire(m, v868+int32(_a_F_WaitOnLock_18), int32(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v879 = F_LWLockAcquire(m, v875+int32(_a_F_WaitOnLock_19), int32(0))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	if v883 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1281+int32(_a_F_WaitOnLock_19))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L195
	}
L128:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v882)+4))
	if v886 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v889 = int32(0)
	v890 = m.G0
	v892 = v890 - int32(16)
	m.G0 = v892
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74])) = v889
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75])) = v889
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v889
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v889
	v906 = F_DeadLockCheckRecurse(m, v882)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L131
	}
L130:
	;
	m.G0 = v892 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = v1216
	if v1216 != int32(3) {
		goto L127
	} else {
		goto L192
	}
L131:
	;
	if v906 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if int32(0) < v911 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v1078 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v1078
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1078
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1078
	*(*int32)(unsafe.Add(mBase, uint32(v892)+12)) = v1078
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[80]))
	v1092 = v892 + int32(12)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v882)+616))
	if v1096 != 0 {
		goto L156
	} else {
		goto L157
	}
L135:
	;
	v920 = v889
	goto L138
L136:
	;
	goto L137
L137:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	if v1075 != 0 {
		goto L152
	} else {
		goto L153
	}
L138:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[81]))
	v944 = v941 + v920*int32(12)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v944)+8))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v944)))
	v948 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v947)+40)) = v948
	v951 = v947 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v947)+36)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v947)+32)) = v951
	if v946 <= v948 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if int32(0) < v1041 {
		v1216 = int32(2)
		goto L130
	} else {
		goto L151
	}
L140:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+15)))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1032<<(uint(int32(2))%32))+uint32(_c_F_WaitOnLock[82])))
	goto L148
L141:
	;
	v962 = v951
	v967 = v948
	goto L142
L142:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v945+v967<<(uint(int32(2))%32))))
	if v962 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v947)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v947)+36)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(v947)+32)) = v951
	goto L146
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v986)+4)) = v951
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v947)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v986))) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v994)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v947)+32)) = v986
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v947)+40))
	v999 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v947)+40)) = v998 + v999
	v1003 = v967 + v999
	if v1003 == v946 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v947)+36))
	v962 = v1005
	v967 = v1003
	goto L142
L148:
	;
	F_ProcLockWakeup(m, v1035, v947)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v1039 = v920 + int32(1)
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if v1039 < v1041 {
		v920 = v1039
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
	v1076 = int32(4)
	goto L154
L153:
	;
	v1076 = int32(1)
	goto L154
L154:
	;
	v1216 = v1076
	goto L130
L155:
	;
	if v1197 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	v1097 = v1096
	goto L158
L157:
	;
	v1097 = v882
	goto L158
L158:
	;
	v1098 = int32(0)
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[83]))
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78]))
	if v1098 < v1102 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1105 = v1098
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1102 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1100+v1102<<(uint(int32(2))%32)))) = v1097
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1139 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1100+v1105<<(uint(int32(2))%32))))
	if v1097 == v1115 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	if v1105 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v1122 = v1105 + int32(1)
	if v1122 != v1102 {
		v1105 = v1122
		goto L162
	} else {
		goto L170
	}
L167:
	;
	v1197 = int32(0)
	goto L155
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1078
	v1197 = int32(1)
	goto L155
L170:
	;
	goto L163
L171:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+624))
	if v1149 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+92))
	if v1142 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1145 = F_FindLockCycleRecurseMember(m, v1097, v1097, v1078, v1090, v1092)
	mBase = m.M
	if v1145 == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v1197 = int32(1)
	goto L155
L175:
	;
	v1197 = int32(0)
	goto L155
L176:
	;
	v1153 = v1097 + int32(620)
	if v1149 == v1153 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1155 = v1149
	goto L178
L178:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1155-int32(624))))
	if v1164 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L175
L180:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+4))
	if v1180 != v1153 {
		v1155 = v1180
		goto L178
	} else {
		goto L185
	}
L181:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1155-int32(536))))
	if v1169 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v1173 = v1155 - int32(628)
	if v1173 == v1097 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v1175 = F_FindLockCycleRecurseMember(m, v1173, v1097, v1078, v1090, v1092)
	mBase = m.M
	if v1175 == int32(0) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1197 = int32(1)
	goto L155
L185:
	;
	goto L179
L186:
	;
	v1216 = int32(3)
	goto L130
L187:
	;
	goto L188
L188:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_20), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_21), int32(243), int32(_a_F_WaitOnLock_22))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
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
	v1246 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[84]))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+92))
	v1250 = F_get_hash_value(m, v1248, v1249)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L193
	}
L193:
	;
	F_RemoveFromWaitQueue(m, v1246, v1250)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L194
	}
L194:
	;
	goto L127
L195:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1287+int32(_a_F_WaitOnLock_18))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1293+int32(_a_F_WaitOnLock_17))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1299+int32(_a_F_WaitOnLock_16))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1305+int32(_a_F_WaitOnLock_15))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1311+int32(_a_F_WaitOnLock_14))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1317+int32(_a_F_WaitOnLock_13))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1323+int32(_a_F_WaitOnLock_12))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1329+int32(_a_F_WaitOnLock_11))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1335+int32(_a_F_WaitOnLock_10))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1341+int32(_a_F_WaitOnLock_9))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L205
	}
L205:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1347+int32(_a_F_WaitOnLock_8))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1353+int32(_a_F_WaitOnLock_7))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1359+int32(_a_F_WaitOnLock_6))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1365+int32(_a_F_WaitOnLock_5))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1371+int32(_a_F_WaitOnLock_4))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
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
	v1410 = m.ExcPending
	if v1410 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v1426 = v229
	goto L46
L213:
	;
	v1445 = int32(_a_F_WaitOnLock_23)
	v1446 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	v1448 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v1448
	v1451 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v1455 = F_LWLockAcquire(m, v1451+int32(512), v1448)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L216
	}
L214:
	;
	v1616 = v235
	v1617 = v1441
	goto L215
L215:
	;
	v1618 = int32(0)
	v1621 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[85])))
	if base.B2i32(v1617 == v1618)|base.B2i32(v1621 != int32(1)) == v1618 {
		goto L256
	} else {
		goto L257
	}
L216:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[86]))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+12))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+48))
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459+v1460))))
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+15)))
	v1464 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+352)) = v1464
	v1466 = *(*int64)(unsafe.Add(mBase, uint32(v230)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+360)) = v1466
	v1469 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1469+int32(512))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L217
	}
L217:
	;
	if v1462&int32(9) != int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	v1616 = int32(0)
	v1617 = v1611
	goto L215
L219:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1446)+44))
	v1479 = int32(14)
	goto L222
L220:
	;
	if v1516 != 0 {
		goto L233
	} else {
		goto L234
	}
L221:
	;
	goto L220
L222:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[87]))
	goto L225
L223:
	;
	v1499 = int32(0)
	goto L230
L225:
	;
	goto L226
L226:
	;
	if int32(0)|base.B2i32(v1486 == int32(15)) != 0 {
		goto L223
	} else {
		goto L228
	}
L228:
	;
	if v1486 <= v1479 {
		v1516 = int32(1)
		goto L221
	} else {
		goto L229
	}
L229:
	;
	goto L223
L230:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	if v1503 != int32(2) {
		v1516 = v1499
		goto L221
	} else {
		goto L231
	}
L231:
	;
	v1507 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[89])))
	if v1507&int32(1) != 0 {
		v1516 = v1499
		goto L221
	} else {
		goto L232
	}
L232:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[90]))
	v1516 = int32(0) | base.B2i32(v1513 <= v1479)
	goto L221
L233:
	;
	v1519 = v216 + int32(336)
	F_initStringInfo(m, v1519)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1581 = F_pgmem_kill(m, v1478, int32(2))
	mBase = m.M
	if int32(0) <= v1581 {
		goto L218
	} else {
		goto L250
	}
L236:
	;
	v1523 = v216 + int32(320)
	F_initStringInfo(m, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L237
	}
L237:
	;
	F_DescribeLockTag(m, v1519, v216+int32(352))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L238
	}
L238:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	v1532 = int32(2)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1463<<(uint(v1532)%32))+uint32(_c_F_WaitOnLock[82])))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+8))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1535+v233<<(uint(v1532)%32))))
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+272)) = v1531
	*(*int32)(unsafe.Add(mBase, uint32(v216)+276)) = v1539
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+280)) = v1542
	F_appendStringInfo(m, v1523, int32(_a_F_WaitOnLock_24), v216+int32(272))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L240
	}
L240:
	;
	v1551 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L241
	}
L241:
	;
	if v1551 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+256)) = v1478
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_25), v216+int32(256))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	F_pfree(m, v1571)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L248
	}
L245:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+240)) = v1559
	F_errdetail_log(m, int32(_a_F_WaitOnLock_26), v216+int32(240))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), int32(1525), int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	F_pfree(m, v1574)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L249
	}
L249:
	;
	goto L235
L250:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[92]))
	if v1585 == int32(71) {
		goto L218
	} else {
		goto L251
	}
L251:
	;
	v1590 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L252
	}
L252:
	;
	if v1590 == int32(0) {
		goto L218
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+224)) = v1478
	F_errmsg(m, int32(_a_F_WaitOnLock_29), v216+int32(224))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), int32(1547), int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L255
	}
L255:
	;
	goto L218
L256:
	;
	v1628 = v216 + int32(352)
	F_initStringInfo(m, v1628)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v1439 == int32(1) {
		v229 = v1426
		v235 = v1616
		goto L44
	} else {
		goto L322
	}
L259:
	;
	F_initStringInfo(m, v216+int32(336))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_initStringInfo(m, v216+int32(320))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_DescribeLockTag(m, v1628, v214)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L262
	}
L262:
	;
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+15)))
	v1643 = int32(2)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1642<<(uint(v1643)%32))+uint32(_c_F_WaitOnLock[82])))
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1645)+8))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1646+v233<<(uint(v1643)%32))))
	goto L263
L263:
	;
	v1652 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	v1656 = m.G0
	v1657 = int32(16)
	v1658 = v1656 - v1657
	m.G0 = v1658
	F_gettimeofday(m, v1658)
	mBase = m.M
	v1661 = *(*int64)(unsafe.Add(mBase, uint32(v1658)))
	v1662 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1658)+8)))
	m.G0 = v1658 + v1657
	goto L264
L264:
	;
	v1677 = v1662 + v1661*int64(1000000) - int64(946684800000000) - v1652
	if v1677 <= int64(0) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	v1694 = int32(1000)
	v1695 = base.I32_div_s(v1693, v1694)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+312)) = v1693 - v1695*v1694
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v216)+316))
	v1702 = F_LWLockAcquire(m, v234, int32(1))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L269
	}
L266:
	;
	v1689 = int32(0)
	v1690 = int32(0)
	goto L268
L267:
	;
	v1681 = int64(1000000)
	v1682 = base.I64_div_u_s(v1677, v1681)
	v1689 = base.I32_wrap_i64(v1682)
	v1690 = base.I32_wrap_i64(v1677 - v1682*v1681)
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216+int32(316)))) = v1689
	*(*int32)(unsafe.Add(mBase, uint32(v216+int32(312)))) = v1690
	goto L265
L269:
	;
	v1706 = int32(0)
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+28))
	if v1708 == v1706 {
		v1789 = v1706
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1811 = v1695 + v1700*int32(1000)
	F_LWLockRelease(m, v234)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L289
	}
L271:
	;
	v1713 = v1707 + int32(24)
	if v1708 == v1713 {
		v1789 = v1706
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1718 = v1708
	v1719 = v1706
	v1720 = int32(1)
	v1725 = int32(1)
	goto L273
L273:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1718-int32(16))))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+44))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+96))
	if v1745 == v1718-int32(20) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1789 = v1780
	goto L270
L275:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+4))
	if v1783 != v1713 {
		v1718 = v1783
		v1719 = v1780
		v1720 = v1781
		v1725 = v1782
		goto L273
	} else {
		goto L288
	}
L276:
	;
	if v1720 != 0 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+208)) = v1744
	if v1725 != 0 {
		goto L284
	} else {
		goto L285
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+176)) = v1744
	F_appendStringInfo(m, v216+int32(336), int32(_a_F_WaitOnLock_30), v216+int32(176))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+192)) = v1744
	F_appendStringInfo(m, v216+int32(336), int32(_a_F_WaitOnLock_31), v216+int32(192))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L283
	}
L282:
	;
	v1780 = v1719
	v1781 = int32(0)
	v1782 = v1725
	goto L275
L283:
	;
	v1780 = v1719
	v1781 = int32(0)
	v1782 = v1725
	goto L275
L284:
	;
	v1772 = int32(_a_F_WaitOnLock_30)
	goto L286
L285:
	;
	v1772 = int32(_a_F_WaitOnLock_31)
	goto L286
L286:
	;
	F_appendStringInfo(m, v216+int32(320), v1772, v216+int32(208))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v1780 = v1719 + int32(1)
	v1781 = v1720
	v1782 = int32(0)
	goto L275
L288:
	;
	goto L274
L289:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	switch v1815 - int32(2) {
	case 0:
		goto L293
	case 1:
		goto L292
	default:
		goto L290
	}
L290:
	;
	switch v1439 {
	case 0:
		goto L304
	case 1:
		goto L305
	default:
		goto L303
	}
L291:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+160)) = v1836
	*(*int32)(unsafe.Add(mBase, uint32(v216)+148)) = v1650
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+152)) = v1839
	*(*int32)(unsafe.Add(mBase, uint32(v216)+156)) = v1811
	v1843 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+144)) = v1843
	F_errmsg(m, v1834, v216+int32(144))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L298
	}
L292:
	;
	v1828 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L296
	}
L293:
	;
	v1820 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L294
	}
L294:
	;
	if v1820 == int32(0) {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	v1834 = int32(_a_F_WaitOnLock_32)
	v1835 = int32(1595)
	goto L291
L296:
	;
	if v1828 == int32(0) {
		goto L290
	} else {
		goto L297
	}
L297:
	;
	v1834 = int32(_a_F_WaitOnLock_33)
	v1835 = int32(1610)
	goto L291
L298:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+128)) = v1849
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+132)) = v1851
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1789, v216+int32(128))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), v1835, int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L300
	}
L300:
	;
	goto L290
L301:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = int32(1)
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	F_pfree(m, v1958)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L319
	}
L302:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_27), v1951, int32(_a_F_WaitOnLock_28))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L318
	}
L303:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v1919 == int32(3) {
		goto L301
	} else {
		goto L313
	}
L304:
	;
	v1899 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L310
	}
L305:
	;
	v1868 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L306
	}
L306:
	;
	if v1868 == int32(0) {
		goto L301
	} else {
		goto L307
	}
L307:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+80)) = v1872
	*(*int32)(unsafe.Add(mBase, uint32(v216)+68)) = v1650
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+72)) = v1875
	*(*int32)(unsafe.Add(mBase, uint32(v216)+76)) = v1811
	v1879 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+64)) = v1879
	F_errmsg(m, int32(_a_F_WaitOnLock_36), v216-int32(-64))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L308
	}
L308:
	;
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+48)) = v1886
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+52)) = v1888
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1789, v216+int32(48))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L309
	}
L309:
	;
	v1951 = int32(1619)
	goto L302
L310:
	;
	if v1899 == int32(0) {
		goto L301
	} else {
		goto L311
	}
L311:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+112)) = v1903
	*(*int32)(unsafe.Add(mBase, uint32(v216)+100)) = v1650
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+104)) = v1906
	*(*int32)(unsafe.Add(mBase, uint32(v216)+108)) = v1811
	v1910 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+96)) = v1910
	F_errmsg(m, int32(_a_F_WaitOnLock_37), v216+int32(96))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L312
	}
L312:
	;
	v1951 = int32(1623)
	goto L302
L313:
	;
	v1924 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L314
	}
L314:
	;
	if v1924 == int32(0) {
		goto L301
	} else {
		goto L315
	}
L315:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v1650
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = v1931
	*(*int32)(unsafe.Add(mBase, uint32(v216)+28)) = v1811
	v1935 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v1935
	F_errmsg(m, int32(_a_F_WaitOnLock_38), v216+int32(16))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L316
	}
L316:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v1942
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v1944
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_34), int32(_a_F_WaitOnLock_35), v1789, v216)
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L317
	}
L317:
	;
	v1951 = int32(1643)
	goto L302
L318:
	;
	goto L301
L319:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	F_pfree(m, v1961)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L320
	}
L320:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	F_pfree(m, v1964)
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L321
	}
L321:
	;
	goto L258
L322:
	;
	goto L45
L323:
	;
	v2019 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if v1426&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2019)) != 0 {
		goto L330
	} else {
		goto L331
	}
L324:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[11]))
	if int32(0) < v2000 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v2003 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+364)) = uint8(v2003)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+360)) = int32(2)
	v2007 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+356)) = uint8(v2007)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+352)) = v2003
	F_disable_timeouts(m, v216+int32(352))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	F_disable_timeout(m, int32(1))
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L329
	}
L328:
	;
	goto L323
L329:
	;
	goto L323
L330:
	;
	v2027 = m.G0
	v2028 = int32(16)
	v2029 = v2027 - v2028
	m.G0 = v2029
	F_gettimeofday(m, v2029)
	mBase = m.M
	v2032 = *(*int64)(unsafe.Add(mBase, uint32(v2029)))
	v2033 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2029)+8)))
	m.G0 = v2029 + v2028
	goto L333
L331:
	;
	goto L332
L332:
	;
	m.G0 = v216 + int32(400)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[3])) = v226
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[2])) = v225
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[1])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+184)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v221)+188)) = v225
	m.G0 = v221 + int32(192)
	return v1439
L333:
	;
	v2042 = int32(0)
	F_LogRecoveryConflict(m, int32(9), v238, v2033+v2032*int64(1000000)-int64(946684800000000), v2042, v2042)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		v2072 = v214
		v2073 = v215
		v2079 = v221
		goto L5
	} else {
		goto L334
	}
L334:
	;
	goto L332
L335:
	;
	goto L4
L336:
	;
	v2103 = int32(v2099)
	m.G0 = v2079
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+4))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2103)))
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2106)))
	if v2079+int32(12) == v2109 {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	m.ExcPending = 1
	goto L345
L338:
	;
	if v2113 != 0 {
		goto L342
	} else {
		goto L343
	}
L339:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+4))
	v2113 = v2111
	goto L341
L340:
	;
	v2113 = int32(0)
	goto L341
L341:
	;
	goto L338
L342:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+188))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2079)+184))
	v32 = v2072
	v33 = v2073
	v35 = v2113
	v37 = v2105
	v39 = v2079
	v43 = v2114
	v44 = v2115
	goto L1
L343:
	;
	goto L344
L344:
	;
	F___wasm_longjmp(m, v2106, v2105)
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	return int32(0)
L346:
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
