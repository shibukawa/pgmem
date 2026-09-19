package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_TablesyncWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v468 int64
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v523 int64
	_ = v523
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v660 int32
	_ = v660
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int64
	_ = v695
	var v698 int32
	_ = v698
	var v701 int64
	_ = v701
	var v704 int64
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1031 int32
	_ = v1031
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1046 int64
	_ = v1046
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1074 int32
	_ = v1074
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1132 int32
	_ = v1132
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1290 int32
	_ = v1290
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1346 int32
	_ = v1346
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1379 int32
	_ = v1379
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1399 int32
	_ = v1399
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1475 int32
	_ = v1475
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1557 int32
	_ = v1557
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1583 int32
	_ = v1583
	var v1594 int32
	_ = v1594
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1753 int32
	_ = v1753
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1809 int32
	_ = v1809
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1923 int32
	_ = v1923
	var v1929 int32
	_ = v1929
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1977 int32
	_ = v1977
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2061 int32
	_ = v2061
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2131 int32
	_ = v2131
	var v2152 int32
	_ = v2152
	var v2156 int32
	_ = v2156
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2202 int32
	_ = v2202
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2218 int32
	_ = v2218
	var v2224 int32
	_ = v2224
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2270 int32
	_ = v2270
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2313 int32
	_ = v2313
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2351 int32
	_ = v2351
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2388 int32
	_ = v2388
	var v2424 int32
	_ = v2424
	var v2434 int32
	_ = v2434
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2486 int32
	_ = v2486
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2548 int32
	_ = v2548
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2561 int64
	_ = v2561
	var v2568 int32
	_ = v2568
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2598 int32
	_ = v2598
	var v2604 int32
	_ = v2604
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2663 int32
	_ = v2663
	var v2667 int32
	_ = v2667
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2715 int32
	_ = v2715
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2725 int64
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2733 int32
	_ = v2733
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2764 int32
	_ = v2764
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2782 int32
	_ = v2782
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2812 int32
	_ = v2812
	var v2818 int32
	_ = v2818
	var v2828 int32
	_ = v2828
	var v2835 int32
	_ = v2835
	var v2841 int32
	_ = v2841
	var v2843 int32
	_ = v2843
	var v2844 int64
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2854 int32
	_ = v2854
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2887 int32
	_ = v2887
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int64
	_ = v2896
	var v2903 int64
	_ = v2903
	var v2912 int32
	_ = v2912
	var v2921 int32
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2927 int32
	_ = v2927
	var v2933 int32
	_ = v2933
	var v2940 int32
	_ = v2940
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2945 int64
	_ = v2945
	var v2947 int32
	_ = v2947
	var v2978 int32
	_ = v2978
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2995 int32
	_ = v2995
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3008 int32
	_ = v3008
	var v3014 int32
	_ = v3014
	var v3018 int32
	_ = v3018
	var v3024 int32
	_ = v3024
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3041 int32
	_ = v3041
	var v3046 int32
	_ = v3046
	var v3050 int32
	_ = v3050
	var v3055 int32
	_ = v3055
	var v3064 int32
	_ = v3064
	var v3068 int32
	_ = v3068
	var v3074 int32
	_ = v3074
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3089 int32
	_ = v3089
	var v3097 int32
	_ = v3097
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3117 int32
	_ = v3117
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3147 int32
	_ = v3147
	var v3175 int32
	_ = v3175
	var v3176 int64
	_ = v3176
	var v3180 int32
	_ = v3180
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3186 int32
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3192 int32
	_ = v3192
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3196 int32
	_ = v3196
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3221 int64
	_ = v3221
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3248 int32
	_ = v3248
	var v3253 int32
	_ = v3253
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3283 int32
	_ = v3283
	var v3285 int32
	_ = v3285
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3292 int32
	_ = v3292
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int64
	_ = v3300
	var v3302 int32
	_ = v3302
	var v3307 int32
	_ = v3307
	v2 = int32(0)
	F_SetupApplyOrSyncWorker(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = m.G0
	v32 = v30 - int32(128)
	m.G0 = v32
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = int32(0)
	v39 = v32 + int32(56)
	v42 = m.G0
	v44 = v42 - int32(752)
	m.G0 = v44
	v49 = v2
	v50 = v2
	v51 = v2
	v52 = v2
	v53 = v2
	v54 = int32(-1)
	v67 = v2
	v68 = v2
	goto L3
L3:
	;
	if v54 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v44 + int32(752)
	v3202 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(v3202)))
	v3205 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+36))
	v3208 = v32 - int32(-64)
	F_ReplicationOriginNameForLogicalRep(m, v3203, v3206, v3208)
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L1
	} else {
		goto L549
	}
L5:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2]))
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3]))
	v82 = v44 + int32(352)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v44 + int32(348)
	goto L8
L6:
	;
	v88 = v53
	v89 = v67
	v90 = v68
	goto L7
L7:
	;
	goto L9
L8:
	;
	v88 = int32(0)
	v89 = v78
	v90 = v80
	goto L7
L9:
	;
	if v88 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	goto L4
L11:
	;
	v3175 = int32(m.ExcTag)
	v3176 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3175 == int32(0) {
		goto L539
	} else {
		goto L540
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3104
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3105
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3106
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v3107)
	v3138 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[4]))
	v3139 = F_MemoryContextStrdup(m, v3138, v3117)
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L11
	} else {
		goto L537
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L11
	} else {
		goto L495
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2457 = v44 + int32(616)
	F_appendStringInfoString(m, v2457, v2434)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L11
	} else {
		goto L420
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoChar(m, v44+int32(616), int32(41))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L11
	} else {
		goto L419
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2257 = v44 + int32(616)
	F_appendStringInfoString(m, v2257, int32(_a_F_TablesyncWorkerMain_0))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L11
	} else {
		goto L402
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v44 + int32(352)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v90
	v2192 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2192)+29)))
	if v2193 == int32(1) {
		goto L395
	} else {
		goto L396
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v113 = F_GetSubscriptionRelState(m, v106, v105, v44+int32(600))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+30)))
	if v124 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+24)))
	v130 = v127 ^ int32(1)
	goto L25
L24:
	;
	v130 = int32(0)
	goto L25
L25:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v135 = base.AtomicRmwXchg32(m, v132, int32(56), int32(1))
	if v135 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v141+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1344), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+40)) = uint8(v113)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v44)+600))
	*(*int64)(unsafe.Add(mBase, uint32(v150)+48)) = v152
	v154 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v150)+56)), uint32(v154))
	v158 = v113 & int32(255)
	if v158 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v164 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v158-int32(114)))
	goto L32
L31:
	;
	v164 = v154
	goto L32
L32:
	;
	if v164 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_finish_sync_worker(m)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v178 = F_palloc(m, int32(64))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L11
	} else {
		goto L37
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+36))
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[5]))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+328)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v44)+324)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v44)+320)) = v185
	v204 = F_pg_snprintf(m, v178, int32(64), int32(_a_F_TablesyncWorkerMain_3), v44+int32(320))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+36))
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v217 = int32(1)
	v223 = m.T0[v211].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v208, v217, v217, v130&v217, v178, v44+int32(608))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7])) = v223
	if v223 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+36))
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_ReplicationOriginNameForLogicalRep(m, v270, v267, v44+int32(528))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L11
	} else {
		goto L48
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v44)+608))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v245
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_4), v44)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1381), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+40)))
	switch v281 - int32(100) {
	case 0:
		goto L52
	default:
		v295 = v280
		goto L51
	case 2:
		goto L50
	}
L49:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	if v471 != 0 {
		goto L79
	} else {
		goto L80
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L11
	} else {
		goto L75
	}
L51:
	;
	v298 = base.AtomicRmwXchg32(m, v295, int32(56), int32(1))
	if v298 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	F_ReplicationSlotDropAtPubNode(m, v289, v178, int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v295 = v294
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v304+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1430), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L11
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v313)+48)) = int64(0)
	v316 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+40)) = uint8(v316)
	v318 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v313)+56)), uint32(v318))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L11
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v328)+48))
	v330 = int32(*(*int8)(unsafe.Add(mBase, uint32(v328)+40)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_UpdateSubscriptionRelState(m, v332, v331, v330, v329, int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v345 = v44 + int32(528)
	v347 = F_replorigin_by_name(m, v345, int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	if v347 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v355 = F_replorigin_create(m, v345)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	v357 = v52
	v358 = v347
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L65
	}
L64:
	;
	v357 = v355
	v358 = v355
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v370 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_StartTransactionCommand(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v386 = F_table_open(m, v380, int32(3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v398 = int32(0)
	v400 = m.T0[v390].(func(*base.Module, int32, int32, int32, int32) int32)(m, v396, int32(_a_F_TablesyncWorkerMain_5), v398, v398)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v402 == int32(1) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+304)) = v420
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_6), v44+int32(304))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1481), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v453 = F_replorigin_by_name(m, v44+int32(528), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_replorigin_session_setup(m, v453, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[8])) = uint16(v453)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v468 = F_replorigin_session_get_progress(m)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v468
	v2857 = v49
	v2858 = v50
	v2859 = v51
	v2860 = v52
	goto L13
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v471)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L11
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	if v478 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v478)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v400)+16))
	if v485 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v485)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L11
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v400)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+32)))
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v510 = int32(0)
	v513 = m.T0[v503].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v509, v178, v510, v510, v500, int32(2), v39)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_LockRelationOid(m, int32(_a_F_TablesyncWorkerMain_7), int32(3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v529 = int32(1)
	F_replorigin_advance(m, v358, v523, int64(0), v529, v529)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_UnlockRelationOid(m, int32(_a_F_TablesyncWorkerMain_7), int32(3))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_replorigin_session_setup(m, v358, int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[8])) = uint16(v358)
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+31)))
	if v552 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_SwitchToUntrustedUser(m, v556, v44+int32(516))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L11
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v386)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v577 = F_pg_class_aclcheck(m, v565, v571, int64(1))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L11
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	if v577 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	v580 = int32(*(*int8)(unsafe.Add(mBase, uint32(v579)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	switch v580 - int32(73) {
	case 0, 32:
		v594 = int32(20)
		goto L106
	default:
		goto L107
	case 10:
		goto L111
	case 29:
		goto L108
	case 36:
		goto L109
	case 45:
		goto L110
	}
L103:
	;
	goto L104
L104:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v386)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v613 = int32(0)
	v615 = F_check_enable_rls(m, v608, v613, v613)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L11
	} else {
		goto L113
	}
L105:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_aclcheck_error(m, v577, v596, v597+int32(4))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L11
	} else {
		goto L112
	}
L106:
	;
	v596 = v594
	goto L105
L107:
	;
	v594 = int32(41)
	goto L106
L108:
	;
	v596 = int32(18)
	goto L105
L109:
	;
	v596 = int32(23)
	goto L105
L110:
	;
	v596 = int32(51)
	goto L105
L111:
	;
	v596 = int32(37)
	goto L105
L112:
	;
	goto L104
L113:
	;
	if v615 == int32(2) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L11
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v674 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L11
	} else {
		goto L122
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(1088))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v645 = F_GetUserNameFromId(m, v639, int32(1))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v647 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v645
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_8), v44+int32(16))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1542), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_PushActiveSnapshot(m, v674)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v688 = F_get_namespace_name(m, v683)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v386)+48))
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v692
	v695 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v695
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+688)) = v698
	v701 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+680)) = v701
	v704 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+672)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v44)+668)) = int32(25)
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v717 = m.T0[v710].(func(*base.Module, int32) int32)(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v720 = v690 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+640)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+636)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v728 = v44 + int32(720)
	F_initStringInfo(m, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v735 = F_quote_literal_cstr(m, v688)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v741 = F_quote_literal_cstr(m, v720)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+292)) = v741
	*(*int32)(unsafe.Add(mBase, uint32(v44)+288)) = v735
	F_appendStringInfo(m, v728, int32(_a_F_TablesyncWorkerMain_9), v44+int32(288))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v767 = m.T0[v756].(func(*base.Module, int32, int32, int32, int32) int32)(m, v762, v763, int32(3), v44+int32(704))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	if v769 != int32(2) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L11
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v767)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v815 = F_MakeTupleTableSlot(m, v809, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L11
	} else {
		goto L138
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v767)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+280)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v44)+276)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_11), v44+int32(272))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(860), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L11
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v767)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v824 = F_tuplestore_gettupleslot(m, v817, int32(1), int32(0), v815)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L11
	} else {
		goto L139
	}
L139:
	;
	if v824 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L11
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v863 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815)+6)))
	if v863 <= int32(0) {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+260)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+256)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_13), v44+int32(256))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(867), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L11
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v815, int32(1))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L11
	} else {
		goto L150
	}
L148:
	;
	v874 = v863
	goto L149
L149:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+632)) = v876
	if base.I32_extend16_s(v874) <= int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v815)+6)))
	v874 = v873
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v815, int32(2))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L11
	} else {
		goto L154
	}
L152:
	;
	v889 = v875
	goto L153
L153:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+656)) = uint8(v890)
	v892 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815)+6)))
	if v892 <= int32(2) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v889 = v888
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v815, int32(3))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L11
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v815)+16))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)) = uint8(v903)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_ExecDropSingleTupleTableSlot(m, v815)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L11
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v767)+8))
	if v911 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v911)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L11
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v767)+12))
	if v918 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v918)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L11
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v767)+16))
	if v925 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v925)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L11
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v767)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L11
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v938 = int32(0)
	v940 = base.B2i32(v717 < int32(_a_F_TablesyncWorkerMain_14))
	if v717 < int32(_a_F_TablesyncWorkerMain_14) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1304 = v44 + int32(720)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)))
	v1306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1305))) = uint8(v1306)
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+12)) = v1306
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+4)) = v1306
	goto L231
L174:
	;
	v1276 = v51
	v1280 = v938
	v1290 = int32(0)
	goto L173
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+664)) = int32(22)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v948 = F_makeStringInfo(m)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L11
	} else {
		goto L177
	}
L177:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_GetPublicationsStr(m, v952, v948, int32(1))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v965 = v44 + int32(720)
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	v967 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v966))) = uint8(v967)
	*(*int32)(unsafe.Add(mBase, uint32(v965)+12)) = v967
	*(*int32)(unsafe.Add(mBase, uint32(v965)+4)) = v967
	goto L179
L179:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+244)) = v973
	*(*int32)(unsafe.Add(mBase, uint32(v44)+240)) = v978
	F_appendStringInfo(m, v965, int32(_a_F_TablesyncWorkerMain_15), v44+int32(240))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L11
	} else {
		goto L180
	}
L180:
	;
	v987 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v987)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v994 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v999 = m.T0[v988].(func(*base.Module, int32, int32, int32, int32) int32)(m, v994, v995, int32(1), v44+int32(664))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L11
	} else {
		goto L181
	}
L181:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v999)))
	if v1001 != int32(2) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L11
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(v1041)+40))
	if int64(2) <= v1046 {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L11
	} else {
		goto L186
	}
L186:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v999)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+232)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v44)+228)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+224)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_16), v44+int32(224))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L11
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(920), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L11
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L11
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v999)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1090 = F_MakeTupleTableSlot(m, v1084, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L11
	} else {
		goto L196
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(1088))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_17), v44+int32(32))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L11
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(934), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1099 = F_tuplestore_gettupleslot(m, v1092, int32(1), int32(0), v1090)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L11
	} else {
		goto L197
	}
L197:
	;
	if v1099 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1090)+6)))
	if v1101 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1216 = v51
	v1220 = v938
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1216
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_ExecDropSingleTupleTableSlot(m, v1090)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L11
	} else {
		goto L217
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1090, int32(1))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L11
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+20))
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	if v1112 != 0 {
		v1181 = v51
		v1185 = v938
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+8))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1181
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	m.T0[v1205].(func(*base.Module, int32))(m, v1090)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L11
	} else {
		goto L216
	}
L206:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+16))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1119 = F_pg_detoast_datum(m, v1114)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L11
	} else {
		goto L207
	}
L207:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+16))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+8))
	if v1122 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	v1132 = (v1125<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L210
L209:
	;
	v1132 = v1122
	goto L210
L210:
	;
	if v1121 <= int32(0) {
		v1181 = v51
		v1185 = v938
		goto L205
	} else {
		goto L211
	}
L211:
	;
	v1141 = v51
	v1143 = int32(0)
	v1145 = v938
	goto L212
L212:
	;
	v1167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1119+v1132+v1143<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1141
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1172 = F_bms_add_member(m, v1145, v1167)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L11
	} else {
		goto L214
	}
L213:
	;
	v1181 = v1172
	v1185 = v1172
	goto L205
L214:
	;
	v1175 = v1143 + int32(1)
	if v1175 != v1121 {
		v1141 = v1172
		v1143 = v1175
		v1145 = v1172
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v1216 = v1181
	v1220 = v1185
	goto L200
L217:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v999)+8))
	if v1245 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1216
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v1245)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L11
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	if v1252 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L220
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1216
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v1252)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L11
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v999)+16))
	if v1259 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1216
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v1259)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L11
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1216
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v999)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L11
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v1276 = v1216
	v1280 = v1220
	v1290 = v948
	goto L173
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v1304, int32(_a_F_TablesyncWorkerMain_18))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	v1320 = base.B2i32(v717 < int32(_a_F_TablesyncWorkerMain_19))
	if v717 < int32(_a_F_TablesyncWorkerMain_19) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1332 = int32(4)
	goto L235
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v44+int32(720), int32(_a_F_TablesyncWorkerMain_20))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L11
	} else {
		goto L236
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+216)) = v1337
	*(*int32)(unsafe.Add(mBase, uint32(v44)+208)) = v1337
	if base.Ui32(v717-int32(_a_F_TablesyncWorkerMain_21)) < base.Ui32(int32(_a_F_TablesyncWorkerMain_22)) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1332 = int32(5)
	goto L235
L237:
	;
	v1346 = int32(_a_F_TablesyncWorkerMain_23)
	goto L239
L238:
	;
	v1346 = int32(_a_F_TablesyncWorkerMain_24)
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+212)) = v1346
	F_appendStringInfo(m, v44+int32(720), int32(_a_F_TablesyncWorkerMain_25), v44+int32(208))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L11
	} else {
		goto L240
	}
L240:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1356)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1363 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v1367 = m.T0[v1357].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1363, v1364, v1332, v44+int32(672))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L11
	} else {
		goto L241
	}
L241:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	if v1369 != int32(2) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L11
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1414 = F_palloc0(m, int32(_a_F_TablesyncWorkerMain_26))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L11
	} else {
		goto L249
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L11
	} else {
		goto L246
	}
L246:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+200)) = v1387
	*(*int32)(unsafe.Add(mBase, uint32(v44)+196)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+192)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_11), v44+int32(192))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1001), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L11
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+648)) = v1414
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1422 = F_palloc0(m, int32(_a_F_TablesyncWorkerMain_26))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L11
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+660)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+652)) = v1422
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1433 = F_MakeTupleTableSlot(m, v1427, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L11
	} else {
		goto L251
	}
L251:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1442 = F_tuplestore_gettupleslot(m, v1435, int32(1), int32(0), v1433)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L11
	} else {
		goto L252
	}
L252:
	;
	v1444 = int32(0)
	if v1442 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1455 = v1444
	v1457 = v1444
	goto L256
L254:
	;
	v1632 = v1444
	v1634 = v1444
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_ExecDropSingleTupleTableSlot(m, v1433)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L11
	} else {
		goto L299
	}
L256:
	;
	v1475 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1475 <= int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1632 = v1604
	v1634 = v1605
	goto L255
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1433, int32(1))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L11
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if v1280 != 0 {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	goto L260
L262:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+8))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	m.T0[v1609].(func(*base.Module, int32))(m, v1433)
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L11
	} else {
		goto L296
	}
L263:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+16))
	v1486 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1485))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1491 = F_bms_is_member(m, v1486, v1280)
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L11
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1495 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1495 <= int32(1) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	if v1491 == int32(0) {
		v1604 = v1455
		v1605 = v1457
		goto L262
	} else {
		goto L267
	}
L267:
	;
	goto L265
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1433, int32(2))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L11
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+16))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1511 = F_text_to_cstring(m, v1506)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L11
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	v1513 = int32(2)
	v1514 = v1455 << (uint(v1513) % 32)
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	*(*int32)(unsafe.Add(mBase, uint32(v1514+v1515))) = v1511
	v1518 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1518 <= v1513 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1433, int32(3))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L11
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v44)+652))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+16))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1528+v1514))) = v1531
	v1533 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1533 <= int32(3) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1433, int32(4))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L11
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+16))
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+12))
	if v1544 != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	goto L279
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v44)+660))
	v1550 = F_bms_add_member(m, v1549, v1455)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L11
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	if (v1457|v1320)&int32(1) != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+660)) = v1550
	goto L283
L285:
	;
	v1571 = v1457 | base.B2i32(int32(_a_F_TablesyncWorkerMain_27) < v717)
	goto L287
L286:
	;
	v1557 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1433)+6)))
	if v1557 <= int32(4) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1573 = v1455 + int32(1)
	if v1573 < int32(1664) {
		v1604 = v1573
		v1605 = v1571
		goto L262
	} else {
		goto L292
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1433, int32(5))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L11
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+16))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+16))
	v1571 = base.B2i32(v1568 != int32(0))
	goto L287
L291:
	;
	goto L290
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L11
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v688
	F_errmsg_internal(m, int32(_a_F_TablesyncWorkerMain_28), v44+int32(48))
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L11
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1049), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L11
	} else {
		goto L295
	}
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1623 = F_tuplestore_gettupleslot(m, v1616, int32(1), int32(0), v1433)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L11
	} else {
		goto L297
	}
L297:
	;
	if v1623 != 0 {
		v1455 = v1604
		v1457 = v1605
		goto L256
	} else {
		goto L298
	}
L298:
	;
	goto L257
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+644)) = v1632
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+8))
	if v1659 != 0 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v1659)
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L11
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	if v1666 != 0 {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	goto L302
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v1666)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L11
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+16))
	if v1673 != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	goto L306
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v1673)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L11
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v1367)
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L11
	} else {
		goto L312
	}
L311:
	;
	goto L310
L312:
	;
	v1686 = int32(0)
	if v940 == v1686 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1694 = v44 + int32(720)
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1694)))
	v1696 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1695))) = uint8(v1696)
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+12)) = v1696
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+4)) = v1696
	goto L316
L314:
	;
	v1936 = v1686
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	F_pfree(m, v1967)
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L11
	} else {
		goto L363
	}
L316:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+180)) = v1702
	*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = v1707
	F_appendStringInfo(m, v1694, int32(_a_F_TablesyncWorkerMain_29), v44+int32(176))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L11
	} else {
		goto L317
	}
L317:
	;
	v1716 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1723 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v1728 = m.T0[v1717].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1723, v1724, int32(1), v44+int32(668))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L11
	} else {
		goto L318
	}
L318:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1728)))
	if v1730 != int32(2) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L11
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1769 = F_MakeTupleTableSlot(m, v1763, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L11
	} else {
		goto L325
	}
L322:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+168)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v44)+164)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v688
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_30), v44+int32(160))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L11
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1099), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L11
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1778 = F_tuplestore_gettupleslot(m, v1771, int32(1), int32(0), v1769)
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L11
	} else {
		goto L327
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_ExecDropSingleTupleTableSlot(m, v1769)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L11
	} else {
		goto L348
	}
L327:
	;
	if v1778 == int32(0) {
		v1870 = v1686
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1782 = v1686
	goto L329
L329:
	;
	v1809 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1769)+6)))
	if v1809 <= int32(0) {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1870 = v1851
	goto L326
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_slot_getsomeattrs_int(m, v1769, int32(1))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L11
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1769)+20))
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819))))
	if v1820 == int32(1) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	goto L333
L335:
	;
	if v1782 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L336:
	;
	goto L337
L337:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1769)+16))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1833)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1839 = F_text_to_cstring(m, v1834)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L11
	} else {
		goto L342
	}
L338:
	;
	v1870 = int32(0)
	goto L326
L339:
	;
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_list_free_deep(m, v1782)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L11
	} else {
		goto L341
	}
L341:
	;
	v1870 = int32(0)
	goto L326
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1845 = F_makeString(m, v1839)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L11
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1851 = F_lappend(m, v1782, v1845)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L11
	} else {
		goto L344
	}
L344:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1769)+8))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	m.T0[v1854].(func(*base.Module, int32))(m, v1769)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L11
	} else {
		goto L345
	}
L345:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1868 = F_tuplestore_gettupleslot(m, v1861, int32(1), int32(0), v1769)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L11
	} else {
		goto L346
	}
L346:
	;
	if v1868 != 0 {
		v1782 = v1851
		goto L329
	} else {
		goto L347
	}
L347:
	;
	goto L330
L348:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+8))
	if v1903 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v1903)
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L11
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+12))
	if v1910 != 0 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	goto L351
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v1910)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L11
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+16))
	if v1917 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	goto L355
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v1917)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L11
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v1728)
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L11
	} else {
		goto L361
	}
L360:
	;
	goto L359
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_free_attrmap(m, v1290)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L11
	} else {
		goto L362
	}
L362:
	;
	v1936 = v1870
	goto L315
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_logicalrep_relmap_update(m, v44+int32(632))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L11
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	v1984 = F_logicalrep_rel_open(m, v1982, int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L11
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v1991 = v44 + int32(616)
	F_initStringInfo(m, v1991)
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L11
	} else {
		goto L366
	}
L366:
	;
	v1994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)))
	v1997 = int32(0)
	if (base.B2i32(v1994 != int32(114))|base.B2i32(v1936 != v1997)|v1634)&int32(1) == v1997 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v44)+636))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v44)+640))
	v2011 = F_quote_qualified_identifier(m, v2009, v2010)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L11
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v44+int32(616), int32(_a_F_TablesyncWorkerMain_31))
	mBase = m.M
	v2120 = m.ExcPending
	if v2120 != 0 {
		goto L11
	} else {
		goto L384
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+144)) = v2011
	F_appendStringInfo(m, v1991, int32(_a_F_TablesyncWorkerMain_32), v44+int32(144))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L11
	} else {
		goto L371
	}
L371:
	;
	v2023 = int32(_a_F_TablesyncWorkerMain_33)
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2024 == int32(0) {
		v2434 = v2023
		goto L14
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v1991, int32(_a_F_TablesyncWorkerMain_34))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L11
	} else {
		goto L373
	}
L373:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2034 <= int32(0) {
		goto L15
	} else {
		goto L374
	}
L374:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2037)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2043 = F_quote_identifier(m, v2038)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L11
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v1991, v2043)
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L11
	} else {
		goto L376
	}
L376:
	;
	v2051 = int32(1)
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2052 <= v2051 {
		goto L15
	} else {
		goto L377
	}
L377:
	;
	v2061 = v2051
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2087 = v44 + int32(616)
	F_appendStringInfoString(m, v2087, int32(_a_F_TablesyncWorkerMain_35))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L11
	} else {
		goto L380
	}
L379:
	;
	goto L15
L380:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2091+v2061<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2100 = F_quote_identifier(m, v2095)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L11
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v2087, v2100)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L11
	} else {
		goto L382
	}
L382:
	;
	v2109 = v2061 + int32(1)
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2109 < v2110 {
		v2061 = v2109
		goto L378
	} else {
		goto L383
	}
L383:
	;
	goto L379
L384:
	;
	v2121 = int32(0)
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2122 <= v2121 {
		goto L16
	} else {
		goto L385
	}
L385:
	;
	v2131 = v2121
	goto L386
L386:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2152+v2131<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2161 = F_quote_identifier(m, v2156)
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L11
	} else {
		goto L388
	}
L387:
	;
	goto L16
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2168 = v44 + int32(616)
	F_appendStringInfoString(m, v2168, v2161)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2131 < v2171-int32(1) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v2168, int32(_a_F_TablesyncWorkerMain_35))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L11
	} else {
		goto L393
	}
L391:
	;
	v2183 = v2171
	goto L392
L392:
	;
	v2185 = v2131 + int32(1)
	if v2185 < v2183 {
		v2131 = v2185
		goto L386
	} else {
		goto L394
	}
L393:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	v2183 = v2182
	goto L392
L394:
	;
	goto L387
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L11
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L11
	} else {
		goto L399
	}
L398:
	;
	v3104 = v49
	v3105 = v50
	v3106 = v51
	v3107 = v52
	v3117 = int32(0)
	goto L12
L399:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2210)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_pgstat_report_subscription_error(m, v2211, int32(0))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_pg_re_throw(m)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L11
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)))
	if v2261 == int32(114) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v2257, int32(_a_F_TablesyncWorkerMain_36))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L11
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v44)+636))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v44)+640))
	v2277 = F_quote_qualified_identifier(m, v2275, v2276)
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L11
	} else {
		goto L407
	}
L406:
	;
	goto L405
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2284 = v44 + int32(616)
	F_appendStringInfoString(m, v2284, v2277)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L11
	} else {
		goto L408
	}
L408:
	;
	v2287 = int32(_a_F_TablesyncWorkerMain_37)
	if v1936 == int32(0) {
		v2434 = v2287
		goto L14
	} else {
		goto L409
	}
L409:
	;
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2290)))
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v2292
	F_appendStringInfo(m, v2284, int32(_a_F_TablesyncWorkerMain_38), v44+int32(128))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L11
	} else {
		goto L410
	}
L410:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+4))
	if int32(2) <= v2304 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v2313 = int32(1)
	goto L414
L412:
	;
	goto L413
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_list_free_deep(m, v1936)
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L11
	} else {
		goto L418
	}
L414:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+12))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2334+v2313<<(uint(int32(2))%32))))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2338)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v2339
	F_appendStringInfo(m, v44+int32(616), int32(_a_F_TablesyncWorkerMain_39), v44+int32(112))
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L11
	} else {
		goto L416
	}
L415:
	;
	goto L413
L416:
	;
	v2353 = v2313 + int32(1)
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(v1936)+4))
	if v2353 < v2354 {
		v2313 = v2353
		goto L414
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	v2434 = v2287
	goto L14
L419:
	;
	v2434 = v2023
	goto L14
L420:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2461)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2469 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2470 = m.T0[v2462].(func(*base.Module, int32) int32)(m, v2469)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L11
	} else {
		goto L422
	}
L421:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(v2518)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2525 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v44)+616))
	v2527 = int32(0)
	v2529 = m.T0[v2519].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2525, v2526, v2527, v2527)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L11
	} else {
		goto L429
	}
L422:
	;
	if v2470 < int32(_a_F_TablesyncWorkerMain_40) {
		v2515 = v50
		v2516 = int32(0)
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476)+26)))
	if v2477 != int32(1) {
		v2515 = v50
		v2516 = int32(0)
		goto L421
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_appendStringInfoString(m, v2457, int32(_a_F_TablesyncWorkerMain_41))
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L11
	} else {
		goto L425
	}
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2492 = F_makeString(m, int32(_a_F_TablesyncWorkerMain_42))
	mBase = m.M
	v2493 = m.ExcPending
	if v2493 != 0 {
		goto L11
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2500 = F_makeDefElem(m, int32(_a_F_TablesyncWorkerMain_43), v2492, int32(-1))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L11
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+612)) = v2500
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = v2500
	v2512 = F_list_make1_impl(m, int32(1), v44+int32(108))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L11
	} else {
		goto L428
	}
L428:
	;
	v2515 = v2512
	v2516 = v2512
	goto L421
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v44)+616))
	F_pfree(m, v2535)
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L11
	} else {
		goto L430
	}
L430:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2529)))
	if v2538 != int32(4) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L11
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+8))
	if v2578 != 0 {
		goto L438
	} else {
		goto L439
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L11
	} else {
		goto L435
	}
L435:
	;
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2561 = *(*int64)(unsafe.Add(mBase, uint32(v44)+636))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+104)) = v2556
	*(*int64)(unsafe.Add(mBase, uint32(v44)+96)) = v2561
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_44), v44+int32(96))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L11
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1266), int32(_a_F_TablesyncWorkerMain_45))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L11
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v2578)
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		goto L11
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+12))
	if v2585 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	goto L440
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v2585)
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L11
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+16))
	if v2592 != 0 {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	goto L444
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v2592)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L11
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v2529)
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L11
	} else {
		goto L450
	}
L449:
	;
	goto L448
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2610 = F_makeStringInfo(m)
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L11
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[15])) = v2610
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2618 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		goto L11
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2625 = int32(0)
	v2628 = F_addRangeTableEntryForRelation(m, v2618, v386, int32(1), v2625, v2625, v2625)
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L11
	} else {
		goto L453
	}
L453:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v1984)+12))
	if v2630 <= int32(0) {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2715 = int32(0)
	v2719 = F_BeginCopyFrom(m, v2618, v386, v2715, v2715, v2715, int32(1023), v2684, v2516)
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L11
	} else {
		goto L463
	}
L455:
	;
	v2684 = int32(0)
	v2686 = v49
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2634 = int32(0)
	v2638 = v49
	v2642 = v2634
	v2644 = v2634
	goto L458
L458:
	;
	v2663 = *(*int32)(unsafe.Add(mBase, uint32(v1984)+16))
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v2663+v2642<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2638
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2672 = F_makeString(m, v2667)
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L11
	} else {
		goto L460
	}
L459:
	;
	v2684 = v2678
	v2686 = v2678
	goto L454
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2638
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2678 = F_lappend(m, v2644, v2672)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L11
	} else {
		goto L461
	}
L461:
	;
	v2681 = v2642 + int32(1)
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v1984)+12))
	if v2681 < v2682 {
		v2638 = v2678
		v2642 = v2681
		v2644 = v2678
		goto L458
	} else {
		goto L462
	}
L462:
	;
	goto L459
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2725 = F_CopyFrom(m, v2719)
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L11
	} else {
		goto L464
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_logicalrep_rel_close(m, v1984, int32(0))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L11
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L11
	} else {
		goto L466
	}
L466:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v2741)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	v2748 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2750 = int32(0)
	v2752 = m.T0[v2742].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2748, int32(_a_F_TablesyncWorkerMain_46), v2750, v2750)
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L11
	} else {
		goto L467
	}
L467:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v2752)))
	if v2754 != int32(1) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L11
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2752)+8))
	if v2792 != 0 {
		goto L475
	} else {
		goto L476
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2771 = m.ExcPending
	if v2771 != 0 {
		goto L11
	} else {
		goto L472
	}
L472:
	;
	v2772 = *(*int32)(unsafe.Add(mBase, uint32(v2752)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v2772
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_47), v44+int32(80))
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L11
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1554), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2791 = m.ExcPending
	if v2791 != 0 {
		goto L11
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v2792)
	mBase = m.M
	v2798 = m.ExcPending
	if v2798 != 0 {
		goto L11
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2752)+12))
	if v2799 != 0 {
		goto L479
	} else {
		goto L480
	}
L478:
	;
	goto L477
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_tuplestore_end(m, v2799)
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L11
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2752)+16))
	if v2806 != 0 {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	goto L481
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_FreeTupleDesc(m, v2806)
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L11
	} else {
		goto L486
	}
L484:
	;
	goto L485
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_pfree(m, v2752)
	mBase = m.M
	v2818 = m.ExcPending
	if v2818 != 0 {
		goto L11
	} else {
		goto L487
	}
L486:
	;
	goto L485
L487:
	;
	if v552 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_RestoreUserContext(m, v44+int32(516))
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L11
	} else {
		goto L491
	}
L489:
	;
	goto L490
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_relation_close(m, v386, int32(0))
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L11
	} else {
		goto L492
	}
L491:
	;
	goto L490
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2841 = m.ExcPending
	if v2841 != 0 {
		goto L11
	} else {
		goto L493
	}
L493:
	;
	v2843 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2844 = *(*int64)(unsafe.Add(mBase, uint32(v2843)+48))
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2843)+36))
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2843)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1276
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v357)
	F_UpdateSubscriptionRelState(m, v2846, v2845, int32(102), v2844, int32(0))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L11
	} else {
		goto L494
	}
L494:
	;
	v2857 = v2686
	v2858 = v2515
	v2859 = v1276
	v2860 = v357
	goto L13
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v2894 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		goto L11
	} else {
		goto L496
	}
L496:
	;
	if v2894 != 0 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2896 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+72)) = uint32(v2896)
	v2903 = int64(base.Ui64(v2896) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+68)) = uint32(v2903)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = v44 + int32(528)
	F_errmsg_internal(m, int32(_a_F_TablesyncWorkerMain_48), v44-int32(-64))
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L11
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2927 = base.AtomicRmwXchg32(m, v2924, int32(56), int32(1))
	if v2927 != 0 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1581), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2921 = m.ExcPending
	if v2921 != 0 {
		goto L11
	} else {
		goto L501
	}
L501:
	;
	goto L499
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v2933 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v2933+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1586), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L11
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2942 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2943 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v2942)+40)) = uint8(v2943)
	v2945 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v2942)+48)) = v2945
	v2947 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2942)+56)), uint32(v2947))
	goto L506
L505:
	;
	goto L504
L506:
	;
	v2978 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[16]))
	if v2978 != 0 {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3097 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	F_LWLockRelease(m, v3097+int32(_a_F_TablesyncWorkerMain_49))
	mBase = m.M
	v3101 = m.ExcPending
	if v3101 != 0 {
		goto L11
	} else {
		goto L536
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	F_ProcessInterrupts(m)
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L11
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	v2986 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2986)+40)))
	if v2987 == int32(99) {
		v3104 = v2857
		v3105 = v2858
		v3106 = v2859
		v3107 = v2860
		v3117 = v178
		goto L12
	} else {
		goto L512
	}
L511:
	;
	goto L510
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v2995 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	v2999 = F_LWLockAcquire(m, v2995+int32(_a_F_TablesyncWorkerMain_49), int32(1))
	mBase = m.M
	v3000 = m.ExcPending
	if v3000 != 0 {
		goto L11
	} else {
		goto L513
	}
L513:
	;
	v3002 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v3003 = *(*int32)(unsafe.Add(mBase, uint32(v3002)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3008 = int32(0)
	v3014 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[18]))
	if v3014 <= v3008 {
		v3046 = v3008
		goto L515
	} else {
		goto L516
	}
L514:
	;
	if v3046 != 0 {
		goto L525
	} else {
		goto L526
	}
L515:
	;
	goto L514
L516:
	;
	v3018 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[19]))
	v3024 = v3008
	goto L517
L517:
	;
	v3029 = v3018 + int32(16) + v3024*int32(112)
	v3030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3029)+16)))
	if v3030 != int32(1) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v3046 = int32(0)
	goto L515
L519:
	;
	v3041 = v3024 + int32(1)
	if v3041 != v3014 {
		v3024 = v3041
		goto L517
	} else {
		goto L524
	}
L520:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3029)))
	if v3033 == int32(3) {
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v3029)+32))
	if v3036 != v3003 {
		goto L519
	} else {
		goto L522
	}
L522:
	;
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(v3029)+36))
	if v3038 != v3008 {
		goto L519
	} else {
		goto L523
	}
L523:
	;
	v3046 = v3029
	goto L515
L524:
	;
	goto L518
L525:
	;
	v3050 = *(*int32)(unsafe.Add(mBase, uint32(v3046)+20))
	if v3050 != 0 {
		goto L528
	} else {
		goto L529
	}
L526:
	;
	goto L527
L527:
	;
	goto L507
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v3046)+20))
	F_SetLatch(m, v3055+int32(20))
	mBase = m.M
	goto L531
L529:
	;
	goto L530
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3064 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	F_LWLockRelease(m, v3064+int32(_a_F_TablesyncWorkerMain_49))
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L11
	} else {
		goto L532
	}
L531:
	;
	goto L530
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3074 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[20]))
	v3078 = F_WaitLatch(m, v3074, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v3079 = m.ExcPending
	if v3079 != 0 {
		goto L11
	} else {
		goto L533
	}
L533:
	;
	if v3078&int32(1) == int32(0) {
		goto L506
	} else {
		goto L534
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2857
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2858
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2859
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2860)
	v3089 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v3089))) = int32(0)
	goto L535
L535:
	;
	goto L506
L536:
	;
	v3104 = v2857
	v3105 = v2858
	v3106 = v2859
	v3107 = v2860
	v3117 = v178
	goto L12
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(52)))) = v3139
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3105
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3104
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3106
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v3107)
	F_pfree(m, v3117)
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L11
	} else {
		goto L538
	}
L538:
	;
	goto L10
L539:
	;
	v3180 = int32(v3176)
	m.G0 = v44
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v3180)+4))
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3180)))
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3183)))
	if v44+int32(348) == v3186 {
		goto L542
	} else {
		goto L543
	}
L540:
	;
	m.ExcPending = 1
	goto L1
L541:
	;
	if v3190 != 0 {
		goto L545
	} else {
		goto L546
	}
L542:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v3183)+4))
	v3190 = v3188
	goto L544
L543:
	;
	v3190 = int32(0)
	goto L544
L544:
	;
	goto L541
L545:
	;
	v3191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)))
	v3192 = *(*int32)(unsafe.Add(mBase, uint32(v44)+744))
	v3193 = *(*int32)(unsafe.Add(mBase, uint32(v44)+740))
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v44)+736))
	v49 = v3194
	v50 = v3193
	v51 = v3192
	v52 = v3191
	v53 = v3182
	v54 = v3190
	v67 = v89
	v68 = v90
	goto L3
L546:
	;
	goto L547
L547:
	;
	F___wasm_longjmp(m, v3183, v3182)
	mBase = m.M
	v3196 = m.ExcPending
	if v3196 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L549:
	;
	F_set_apply_error_context_origin(m, v3208)
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v3214 = int32(1)
	v3216 = v32 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v3216))) = uint8(v3214)
	v3221 = *(*int64)(unsafe.Add(mBase, uint32(v32+int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+4)) = v3213
	*(*int64)(unsafe.Add(mBase, uint32(v3216)+8)) = v3221
	v3226 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v3228 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v3229 = *(*int32)(unsafe.Add(mBase, uint32(v3228)+24))
	v3230 = m.T0[v3229].(func(*base.Module, int32) int32)(m, v3226)
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L1
	} else {
		goto L553
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+28)) = v3280
	v3283 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3283)+68)) = uint8(v3279)
	v3285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3216)+32)) = uint8(v3285)
	v3287 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+52))
	v3288 = F_pstrdup(m, v3287)
	mBase = m.M
	v3289 = m.ExcPending
	if v3289 != 0 {
		goto L1
	} else {
		goto L570
	}
L552:
	;
	v3269 = int32(0)
	if v3266&int32(255) != int32(102) {
		goto L567
	} else {
		goto L568
	}
L553:
	;
	if v3230 <= int32(_a_F_TablesyncWorkerMain_50) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	if int32(_a_F_TablesyncWorkerMain_51) < v3230 {
		goto L557
	} else {
		goto L558
	}
L555:
	;
	goto L556
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+16)) = int32(4)
	v3257 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3258 = *(*int32)(unsafe.Add(mBase, uint32(v3257)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+20)) = v3258
	v3260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3257)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3216)+24)) = uint8(v3260)
	v3263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3257)+27)))
	if v3263 == int32(112) {
		v3278 = v3257
		v3279 = v3214
		v3280 = int32(_a_F_TablesyncWorkerMain_52)
		goto L551
	} else {
		goto L566
	}
L557:
	;
	v3239 = int32(2)
	goto L559
L558:
	;
	v3239 = int32(1)
	goto L559
L559:
	;
	if int32(_a_F_TablesyncWorkerMain_53) < v3230 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3242 = int32(3)
	goto L562
L561:
	;
	v3242 = v3239
	goto L562
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+16)) = v3242
	v3245 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3245)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+20)) = v3246
	v3248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3245)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3216)+24)) = uint8(v3248)
	if v3230 < int32(_a_F_TablesyncWorkerMain_54) {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v3278 = v3245
	v3279 = int32(0)
	v3280 = int32(0)
	goto L551
L564:
	;
	goto L565
L565:
	;
	v3253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3245)+27)))
	v3266 = v3253
	v3267 = v3245
	goto L552
L566:
	;
	v3266 = v3263
	v3267 = v3257
	goto L552
L567:
	;
	v3276 = int32(_a_F_TablesyncWorkerMain_55)
	goto L569
L568:
	;
	v3276 = v3269
	goto L569
L569:
	;
	v3278 = v3267
	v3279 = v3269
	v3280 = v3276
	goto L551
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3216)+36)) = v3288
	v3292 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v3296 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(v3296)+32))
	v3298 = m.T0[v3297].(func(*base.Module, int32, int32) int32)(m, v3292, v32+int32(8))
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	v3300 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
	F_start_apply(m, v3300)
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	m.G0 = v32 + int32(128)
	F_finish_sync_worker(m)
	mBase = m.M
	v3307 = m.ExcPending
	if v3307 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
