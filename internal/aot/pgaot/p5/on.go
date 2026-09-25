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
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v924 int32
	_ = v924
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
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1119 int32
	_ = v1119
	var v1126 int32
	_ = v1126
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
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
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1430 int32
	_ = v1430
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int64
	_ = v1468
	var v1470 int64
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1490 int32
	_ = v1490
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1656 int64
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int64
	_ = v1665
	var v1666 int64
	_ = v1666
	var v1681 int64
	_ = v1681
	var v1685 int64
	_ = v1685
	var v1686 int64
	_ = v1686
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1760 int32
	_ = v1760
	var v1769 int32
	_ = v1769
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1914 int32
	_ = v1914
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2011 int32
	_ = v2011
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2036 int64
	_ = v2036
	var v2037 int64
	_ = v2037
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2083 int32
	_ = v2083
	var v2102 int32
	_ = v2102
	var v2103 int64
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
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
	v2102 = int32(m.ExcTag)
	v2103 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2102 == int32(0) {
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
	v2075 = m.ExcPending
	if v2075 != 0 {
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
		v2076 = v32
		v2077 = v33
		v2083 = v39
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
	v2000 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if base.Ui32(int32(1)) < base.Ui32(v2000) {
		goto L323
	} else {
		goto L324
	}
L46:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+16))
	v1445 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v235&base.B2i32(v1445 == int32(4)) != 0 {
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v1430 = v229
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
		v1430 = int32(0)
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v1430 = int32(1)
	goto L46
L106:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[18]))
	v765 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = v765
	v770 = base.AtomicRmwOr32(m, v765, int32(_a_F_WaitOnLock_5), v765)
	goto L107
L107:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[6]))
	if v772 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v778 = F_LWLockAcquire(m, v774+int32(_a_F_WaitOnLock_4), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[73]))
	if v1410 == int32(0) {
		v1430 = v229
		goto L46
	} else {
		goto L211
	}
L111:
	;
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v785 = F_LWLockAcquire(m, v781+int32(_a_F_WaitOnLock_6), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v792 = F_LWLockAcquire(m, v788+int32(_a_F_WaitOnLock_7), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L113
	}
L113:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v799 = F_LWLockAcquire(m, v795+int32(_a_F_WaitOnLock_8), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v802 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v806 = F_LWLockAcquire(m, v802+int32(_a_F_WaitOnLock_9), int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L115
	}
L115:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v813 = F_LWLockAcquire(m, v809+int32(_a_F_WaitOnLock_10), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v816 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v820 = F_LWLockAcquire(m, v816+int32(_a_F_WaitOnLock_11), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L117
	}
L117:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v827 = F_LWLockAcquire(m, v823+int32(_a_F_WaitOnLock_12), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L118
	}
L118:
	;
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v834 = F_LWLockAcquire(m, v830+int32(_a_F_WaitOnLock_13), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v837 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v841 = F_LWLockAcquire(m, v837+int32(_a_F_WaitOnLock_14), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L120
	}
L120:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v848 = F_LWLockAcquire(m, v844+int32(_a_F_WaitOnLock_15), int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L121
	}
L121:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v855 = F_LWLockAcquire(m, v851+int32(_a_F_WaitOnLock_16), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v862 = F_LWLockAcquire(m, v858+int32(_a_F_WaitOnLock_17), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L123
	}
L123:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v869 = F_LWLockAcquire(m, v865+int32(_a_F_WaitOnLock_18), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v876 = F_LWLockAcquire(m, v872+int32(_a_F_WaitOnLock_19), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L125
	}
L125:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v883 = F_LWLockAcquire(m, v879+int32(_a_F_WaitOnLock_20), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L126
	}
L126:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	if v887 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1285+int32(_a_F_WaitOnLock_20))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L195
	}
L128:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	if v890 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v893 = int32(0)
	v894 = m.G0
	v896 = v894 - int32(16)
	m.G0 = v896
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[74])) = v893
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[75])) = v893
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v893
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v893
	v910 = F_DeadLockCheckRecurse(m, v886)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L131
	}
L130:
	;
	m.G0 = v896 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7])) = v1220
	if v1220 != int32(3) {
		goto L127
	} else {
		goto L192
	}
L131:
	;
	if v910 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if int32(0) < v915 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v1082 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76])) = v1082
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1082
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1082
	*(*int32)(unsafe.Add(mBase, uint32(v896)+12)) = v1082
	v1094 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[80]))
	v1096 = v896 + int32(12)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v886)+616))
	if v1100 != 0 {
		goto L156
	} else {
		goto L157
	}
L135:
	;
	v924 = v893
	goto L138
L136:
	;
	goto L137
L137:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	if v1079 != 0 {
		goto L152
	} else {
		goto L153
	}
L138:
	;
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[81]))
	v948 = v945 + v924*int32(12)
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
	if int32(0) < v1045 {
		v1220 = int32(2)
		goto L130
	} else {
		goto L151
	}
L140:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+15)))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1036<<(uint(int32(2))%32))+uint32(_c_F_WaitOnLock[82])))
	goto L148
L141:
	;
	v966 = v955
	v971 = v952
	goto L142
L142:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v949+v971<<(uint(int32(2))%32))))
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
	*(*int32)(unsafe.Add(mBase, uint32(v990)+4)) = v955
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v951)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v998)+4)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v951)+32)) = v990
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v951)+40))
	v1003 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v951)+40)) = v1002 + v1003
	v1007 = v971 + v1003
	if v1007 == v950 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v951)+36))
	v966 = v1009
	v971 = v1007
	goto L142
L148:
	;
	F_ProcLockWakeup(m, v1039, v951)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L149
	}
L149:
	;
	v1043 = v924 + int32(1)
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[76]))
	if v1043 < v1045 {
		v924 = v1043
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
	v1080 = int32(4)
	goto L154
L153:
	;
	v1080 = int32(1)
	goto L154
L154:
	;
	v1220 = v1080
	goto L130
L155:
	;
	if v1201 != 0 {
		goto L186
	} else {
		goto L187
	}
L156:
	;
	v1101 = v1100
	goto L158
L157:
	;
	v1101 = v886
	goto L158
L158:
	;
	v1102 = int32(0)
	v1104 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[83]))
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78]))
	if v1102 < v1106 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1109 = v1102
	goto L162
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[78])) = v1106 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1104+v1106<<(uint(int32(2))%32)))) = v1101
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+4))
	if v1143 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1104+v1109<<(uint(int32(2))%32))))
	if v1101 == v1119 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L161
L164:
	;
	if v1109 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	v1126 = v1109 + int32(1)
	if v1126 != v1106 {
		v1109 = v1126
		goto L162
	} else {
		goto L170
	}
L167:
	;
	v1201 = int32(0)
	goto L155
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[79])) = v1082
	v1201 = int32(1)
	goto L155
L170:
	;
	goto L163
L171:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+624))
	if v1153 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+92))
	if v1146 == int32(0) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v1149 = F_FindLockCycleRecurseMember(m, v1101, v1101, v1082, v1094, v1096)
	mBase = m.M
	if v1149 == int32(0) {
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v1201 = int32(1)
	goto L155
L175:
	;
	v1201 = int32(0)
	goto L155
L176:
	;
	v1157 = v1101 + int32(620)
	if v1153 == v1157 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1159 = v1153
	goto L178
L178:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1159-int32(624))))
	if v1168 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L175
L180:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+4))
	if v1184 != v1157 {
		v1159 = v1184
		goto L178
	} else {
		goto L185
	}
L181:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1159-int32(536))))
	if v1173 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v1177 = v1159 - int32(628)
	if v1177 == v1101 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v1179 = F_FindLockCycleRecurseMember(m, v1177, v1101, v1082, v1094, v1096)
	mBase = m.M
	if v1179 == int32(0) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1201 = int32(1)
	goto L155
L185:
	;
	goto L179
L186:
	;
	v1220 = int32(3)
	goto L130
L187:
	;
	goto L188
L188:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L189
	}
L189:
	;
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_21), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_22), int32(243), int32(_a_F_WaitOnLock_23))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v1250 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[13]))
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[84]))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+92))
	v1254 = F_get_hash_value(m, v1252, v1253)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L193
	}
L193:
	;
	F_RemoveFromWaitQueue(m, v1250, v1254)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L194
	}
L194:
	;
	goto L127
L195:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1291+int32(_a_F_WaitOnLock_19))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L196
	}
L196:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1297+int32(_a_F_WaitOnLock_18))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1303+int32(_a_F_WaitOnLock_17))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1309+int32(_a_F_WaitOnLock_16))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1315+int32(_a_F_WaitOnLock_15))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L200
	}
L200:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1321+int32(_a_F_WaitOnLock_14))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1327+int32(_a_F_WaitOnLock_13))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L202
	}
L202:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1333+int32(_a_F_WaitOnLock_12))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1339+int32(_a_F_WaitOnLock_11))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L204
	}
L204:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1345+int32(_a_F_WaitOnLock_10))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L205
	}
L205:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1351+int32(_a_F_WaitOnLock_9))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L206
	}
L206:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1357+int32(_a_F_WaitOnLock_8))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L207
	}
L207:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1363+int32(_a_F_WaitOnLock_7))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1369+int32(_a_F_WaitOnLock_6))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1375+int32(_a_F_WaitOnLock_4))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v1414 = m.ExcPending
	if v1414 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v1430 = v229
	goto L46
L213:
	;
	v1449 = int32(_a_F_WaitOnLock_24)
	v1450 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77]))
	v1452 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[77])) = v1452
	v1455 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	v1459 = F_LWLockAcquire(m, v1455+int32(512), v1452)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L216
	}
L214:
	;
	v1620 = v235
	v1621 = v1445
	goto L215
L215:
	;
	v1622 = int32(0)
	v1625 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[85])))
	if base.B2i32(v1621 == v1622)|base.B2i32(v1625 != int32(1)) == v1622 {
		goto L256
	} else {
		goto L257
	}
L216:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[86]))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+12))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+48))
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463+v1464))))
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+15)))
	v1468 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+352)) = v1468
	v1470 = *(*int64)(unsafe.Add(mBase, uint32(v230)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+360)) = v1470
	v1473 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[4]))
	F_LWLockRelease(m, v1473+int32(512))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L217
	}
L217:
	;
	if v1466&int32(9) != int32(1) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	v1620 = int32(0)
	v1621 = v1615
	goto L215
L219:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+44))
	v1483 = int32(14)
	goto L222
L220:
	;
	if v1520 != 0 {
		goto L233
	} else {
		goto L234
	}
L221:
	;
	goto L220
L222:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[87]))
	goto L225
L223:
	;
	v1503 = int32(0)
	goto L230
L225:
	;
	goto L226
L226:
	;
	if int32(0)|base.B2i32(v1490 == int32(15)) != 0 {
		goto L223
	} else {
		goto L228
	}
L228:
	;
	if v1490 <= v1483 {
		v1520 = int32(1)
		goto L221
	} else {
		goto L229
	}
L229:
	;
	goto L223
L230:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[88]))
	if v1507 != int32(2) {
		v1520 = v1503
		goto L221
	} else {
		goto L231
	}
L231:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitOnLock[89])))
	if v1511&int32(1) != 0 {
		v1520 = v1503
		goto L221
	} else {
		goto L232
	}
L232:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[90]))
	v1520 = int32(0) | base.B2i32(v1517 <= v1483)
	goto L221
L233:
	;
	v1523 = v216 + int32(336)
	F_initStringInfo(m, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1585 = F_pgmem_kill(m, v1482, int32(2))
	mBase = m.M
	if int32(0) <= v1585 {
		goto L218
	} else {
		goto L250
	}
L236:
	;
	v1527 = v216 + int32(320)
	F_initStringInfo(m, v1527)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L237
	}
L237:
	;
	F_DescribeLockTag(m, v1523, v216+int32(352))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L238
	}
L238:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	v1536 = int32(2)
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1467<<(uint(v1536)%32))+uint32(_c_F_WaitOnLock[82])))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+8))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1539+v233<<(uint(v1536)%32))))
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+272)) = v1535
	*(*int32)(unsafe.Add(mBase, uint32(v216)+276)) = v1543
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+280)) = v1546
	F_appendStringInfo(m, v1527, int32(_a_F_WaitOnLock_25), v216+int32(272))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L240
	}
L240:
	;
	v1555 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L241
	}
L241:
	;
	if v1555 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+256)) = v1482
	F_errmsg_internal(m, int32(_a_F_WaitOnLock_26), v216+int32(256))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	F_pfree(m, v1575)
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L248
	}
L245:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+240)) = v1563
	F_errdetail_log(m, int32(_a_F_WaitOnLock_27), v216+int32(240))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_28), int32(1525), int32(_a_F_WaitOnLock_29))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L247
	}
L247:
	;
	goto L244
L248:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	F_pfree(m, v1578)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L249
	}
L249:
	;
	goto L235
L250:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[92]))
	if v1589 == int32(71) {
		goto L218
	} else {
		goto L251
	}
L251:
	;
	v1594 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L252
	}
L252:
	;
	if v1594 == int32(0) {
		goto L218
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+224)) = v1482
	F_errmsg(m, int32(_a_F_WaitOnLock_30), v216+int32(224))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_28), int32(1547), int32(_a_F_WaitOnLock_29))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L255
	}
L255:
	;
	goto L218
L256:
	;
	v1632 = v216 + int32(352)
	F_initStringInfo(m, v1632)
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	if v1443 == int32(1) {
		v229 = v1430
		v235 = v1620
		goto L44
	} else {
		goto L322
	}
L259:
	;
	F_initStringInfo(m, v216+int32(336))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_initStringInfo(m, v216+int32(320))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_DescribeLockTag(m, v1632, v214)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L262
	}
L262:
	;
	v1646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+15)))
	v1647 = int32(2)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1646<<(uint(v1647)%32))+uint32(_c_F_WaitOnLock[82])))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+8))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v233<<(uint(v1647)%32))))
	goto L263
L263:
	;
	v1656 = *(*int64)(unsafe.Add(mBase, _c_F_WaitOnLock[14]))
	v1660 = m.G0
	v1661 = int32(16)
	v1662 = v1660 - v1661
	m.G0 = v1662
	F_gettimeofday(m, v1662)
	mBase = m.M
	v1665 = *(*int64)(unsafe.Add(mBase, uint32(v1662)))
	v1666 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1662)+8)))
	m.G0 = v1662 + v1661
	goto L264
L264:
	;
	v1681 = v1666 + v1665*int64(1000000) - int64(946684800000000) - v1656
	if v1681 <= int64(0) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	v1698 = int32(1000)
	v1699 = base.I32_div_s(v1697, v1698)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+312)) = v1697 - v1699*v1698
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v216)+316))
	v1706 = F_LWLockAcquire(m, v234, int32(1))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L269
	}
L266:
	;
	v1693 = int32(0)
	v1694 = int32(0)
	goto L268
L267:
	;
	v1685 = int64(1000000)
	v1686 = base.I64_div_u_s(v1681, v1685)
	v1693 = base.I32_wrap_i64(v1686)
	v1694 = base.I32_wrap_i64(v1681 - v1686*v1685)
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216+int32(316)))) = v1693
	*(*int32)(unsafe.Add(mBase, uint32(v216+int32(312)))) = v1694
	goto L265
L269:
	;
	v1710 = int32(0)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+28))
	if v1712 == v1710 {
		v1793 = v1710
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1815 = v1699 + v1704*int32(1000)
	F_LWLockRelease(m, v234)
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L289
	}
L271:
	;
	v1717 = v1711 + int32(24)
	if v1712 == v1717 {
		v1793 = v1710
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1722 = v1712
	v1723 = v1710
	v1724 = int32(1)
	v1729 = int32(1)
	goto L273
L273:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1722-int32(16))))
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1747)+44))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1747)+96))
	if v1749 == v1722-int32(20) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	v1793 = v1784
	goto L270
L275:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1722)+4))
	if v1787 != v1717 {
		v1722 = v1787
		v1723 = v1784
		v1724 = v1785
		v1729 = v1786
		goto L273
	} else {
		goto L288
	}
L276:
	;
	if v1724 != 0 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+208)) = v1748
	if v1729 != 0 {
		goto L284
	} else {
		goto L285
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+176)) = v1748
	F_appendStringInfo(m, v216+int32(336), int32(_a_F_WaitOnLock_31), v216+int32(176))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+192)) = v1748
	F_appendStringInfo(m, v216+int32(336), int32(_a_F_WaitOnLock_32), v216+int32(192))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L283
	}
L282:
	;
	v1784 = v1723
	v1785 = int32(0)
	v1786 = v1729
	goto L275
L283:
	;
	v1784 = v1723
	v1785 = int32(0)
	v1786 = v1729
	goto L275
L284:
	;
	v1776 = int32(_a_F_WaitOnLock_31)
	goto L286
L285:
	;
	v1776 = int32(_a_F_WaitOnLock_32)
	goto L286
L286:
	;
	F_appendStringInfo(m, v216+int32(320), v1776, v216+int32(208))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L287
	}
L287:
	;
	v1784 = v1723 + int32(1)
	v1785 = v1724
	v1786 = int32(0)
	goto L275
L288:
	;
	goto L274
L289:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	switch v1819 - int32(2) {
	case 0:
		goto L293
	case 1:
		goto L292
	default:
		goto L290
	}
L290:
	;
	switch v1443 {
	case 0:
		goto L304
	case 1:
		goto L305
	default:
		goto L303
	}
L291:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+160)) = v1840
	*(*int32)(unsafe.Add(mBase, uint32(v216)+148)) = v1654
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+152)) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v216)+156)) = v1815
	v1847 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+144)) = v1847
	F_errmsg(m, v1838, v216+int32(144))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L298
	}
L292:
	;
	v1832 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L296
	}
L293:
	;
	v1824 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L294
	}
L294:
	;
	if v1824 == int32(0) {
		goto L290
	} else {
		goto L295
	}
L295:
	;
	v1838 = int32(_a_F_WaitOnLock_33)
	v1839 = int32(1595)
	goto L291
L296:
	;
	if v1832 == int32(0) {
		goto L290
	} else {
		goto L297
	}
L297:
	;
	v1838 = int32(_a_F_WaitOnLock_34)
	v1839 = int32(1610)
	goto L291
L298:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+128)) = v1853
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+132)) = v1855
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_35), int32(_a_F_WaitOnLock_36), v1793, v216+int32(128))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_28), v1839, int32(_a_F_WaitOnLock_29))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	F_pfree(m, v1962)
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L319
	}
L302:
	;
	F_errfinish(m, int32(_a_F_WaitOnLock_28), v1955, int32(_a_F_WaitOnLock_29))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L318
	}
L303:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[7]))
	if v1923 == int32(3) {
		goto L301
	} else {
		goto L313
	}
L304:
	;
	v1903 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L310
	}
L305:
	;
	v1872 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L306
	}
L306:
	;
	if v1872 == int32(0) {
		goto L301
	} else {
		goto L307
	}
L307:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+80)) = v1876
	*(*int32)(unsafe.Add(mBase, uint32(v216)+68)) = v1654
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+72)) = v1879
	*(*int32)(unsafe.Add(mBase, uint32(v216)+76)) = v1815
	v1883 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+64)) = v1883
	F_errmsg(m, int32(_a_F_WaitOnLock_37), v216-int32(-64))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L308
	}
L308:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+48)) = v1890
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+52)) = v1892
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_35), int32(_a_F_WaitOnLock_36), v1793, v216+int32(48))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L309
	}
L309:
	;
	v1955 = int32(1619)
	goto L302
L310:
	;
	if v1903 == int32(0) {
		goto L301
	} else {
		goto L311
	}
L311:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+112)) = v1907
	*(*int32)(unsafe.Add(mBase, uint32(v216)+100)) = v1654
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+104)) = v1910
	*(*int32)(unsafe.Add(mBase, uint32(v216)+108)) = v1815
	v1914 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+96)) = v1914
	F_errmsg(m, int32(_a_F_WaitOnLock_38), v216+int32(96))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L312
	}
L312:
	;
	v1955 = int32(1623)
	goto L302
L313:
	;
	v1928 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L314
	}
L314:
	;
	if v1928 == int32(0) {
		goto L301
	} else {
		goto L315
	}
L315:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v216)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v1932
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v1654
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v216)+352))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = v1935
	*(*int32)(unsafe.Add(mBase, uint32(v216)+28)) = v1815
	v1939 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[91]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v1939
	F_errmsg(m, int32(_a_F_WaitOnLock_39), v216+int32(16))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L316
	}
L316:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v1946
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v1948
	F_errdetail_log_plural(m, int32(_a_F_WaitOnLock_35), int32(_a_F_WaitOnLock_36), v1793, v216)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L317
	}
L317:
	;
	v1955 = int32(1643)
	goto L302
L318:
	;
	goto L301
L319:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v216)+320))
	F_pfree(m, v1965)
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
		goto L5
	} else {
		goto L320
	}
L320:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v216)+336))
	F_pfree(m, v1968)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v2023 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[8]))
	if v1430&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v2023)) != 0 {
		goto L330
	} else {
		goto L331
	}
L324:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, _c_F_WaitOnLock[11]))
	if int32(0) < v2004 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v2007 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+364)) = uint8(v2007)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+360)) = int32(2)
	v2011 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+356)) = uint8(v2011)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+352)) = v2007
	F_disable_timeouts(m, v216+int32(352))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v2021 = m.ExcPending
	if v2021 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v2031 = m.G0
	v2032 = int32(16)
	v2033 = v2031 - v2032
	m.G0 = v2033
	F_gettimeofday(m, v2033)
	mBase = m.M
	v2036 = *(*int64)(unsafe.Add(mBase, uint32(v2033)))
	v2037 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2033)+8)))
	m.G0 = v2033 + v2032
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
	return v1443
L333:
	;
	v2046 = int32(0)
	F_LogRecoveryConflict(m, int32(9), v238, v2037+v2036*int64(1000000)-int64(946684800000000), v2046, v2046)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		v2076 = v214
		v2077 = v215
		v2083 = v221
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
	v2107 = int32(v2103)
	m.G0 = v2083
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+4))
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v2107)))
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2110)))
	if v2083+int32(12) == v2113 {
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
	if v2117 != 0 {
		goto L342
	} else {
		goto L343
	}
L339:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+4))
	v2117 = v2115
	goto L341
L340:
	;
	v2117 = int32(0)
	goto L341
L341:
	;
	goto L338
L342:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2083)+188))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2083)+184))
	v32 = v2076
	v33 = v2077
	v35 = v2117
	v37 = v2109
	v39 = v2083
	v43 = v2118
	v44 = v2119
	goto L1
L343:
	;
	goto L344
L344:
	;
	F___wasm_longjmp(m, v2110, v2109)
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
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
