package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F___floatscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v22 int64
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v188 int32
	_ = v188
	var v195 int64
	_ = v195
	var v198 int32
	_ = v198
	var v216 int32
	_ = v216
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v305 int32
	_ = v305
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v323 int64
	_ = v323
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v336 int64
	_ = v336
	var v339 int64
	_ = v339
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v356 int64
	_ = v356
	var v374 int32
	_ = v374
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int64
	_ = v487
	var v488 int64
	_ = v488
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v557 int64
	_ = v557
	var v560 int64
	_ = v560
	var v563 int32
	_ = v563
	var v576 int32
	_ = v576
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int64
	_ = v608
	var v611 int32
	_ = v611
	var v639 int64
	_ = v639
	var v647 int64
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v780 int64
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int64
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v825 int64
	_ = v825
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v852 int64
	_ = v852
	var v853 int64
	_ = v853
	var v854 int64
	_ = v854
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v857 int64
	_ = v857
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int64
	_ = v902
	var v903 int32
	_ = v903
	var v907 int64
	_ = v907
	var v920 int64
	_ = v920
	var v922 int64
	_ = v922
	var v925 int64
	_ = v925
	var v926 int64
	_ = v926
	var v939 int64
	_ = v939
	var v940 int64
	_ = v940
	var v941 int64
	_ = v941
	var v942 int64
	_ = v942
	var v944 int64
	_ = v944
	var v945 int64
	_ = v945
	var v947 int64
	_ = v947
	var v948 int64
	_ = v948
	var v959 int64
	_ = v959
	var v960 int64
	_ = v960
	var v963 int64
	_ = v963
	var v964 int64
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int64
	_ = v967
	var v968 int64
	_ = v968
	var v969 int64
	_ = v969
	var v970 int64
	_ = v970
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int64
	_ = v981
	var v982 int64
	_ = v982
	var v983 int64
	_ = v983
	var v984 int64
	_ = v984
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int64
	_ = v998
	var v1001 int32
	_ = v1001
	var v1016 int64
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1041 int64
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int64
	_ = v1051
	var v1053 int64
	_ = v1053
	var v1057 int64
	_ = v1057
	var v1075 int64
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1083 int64
	_ = v1083
	var v1089 int64
	_ = v1089
	var v1090 int64
	_ = v1090
	var v1091 int64
	_ = v1091
	var v1093 int64
	_ = v1093
	var v1105 int64
	_ = v1105
	var v1106 int64
	_ = v1106
	var v1118 int32
	_ = v1118
	var v1131 int64
	_ = v1131
	var v1139 int32
	_ = v1139
	var v1141 int64
	_ = v1141
	var v1153 int32
	_ = v1153
	var v1177 int64
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int64
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1202 int32
	_ = v1202
	var v1205 int64
	_ = v1205
	var v1206 int64
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1216 int64
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1224 int64
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int64
	_ = v1234
	var v1236 int64
	_ = v1236
	var v1240 int64
	_ = v1240
	var v1258 int64
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1266 int64
	_ = v1266
	var v1272 int64
	_ = v1272
	var v1273 int64
	_ = v1273
	var v1274 int64
	_ = v1274
	var v1276 int64
	_ = v1276
	var v1288 int64
	_ = v1288
	var v1289 int64
	_ = v1289
	var v1290 int64
	_ = v1290
	var v1295 int64
	_ = v1295
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1317 int64
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1322 int64
	_ = v1322
	var v1335 int64
	_ = v1335
	var v1337 int64
	_ = v1337
	var v1340 int64
	_ = v1340
	var v1341 int64
	_ = v1341
	var v1349 int64
	_ = v1349
	var v1350 int64
	_ = v1350
	var v1351 int64
	_ = v1351
	var v1352 int64
	_ = v1352
	var v1356 int64
	_ = v1356
	var v1357 int64
	_ = v1357
	var v1361 int64
	_ = v1361
	var v1362 int64
	_ = v1362
	var v1378 int32
	_ = v1378
	var v1390 int64
	_ = v1390
	var v1392 int64
	_ = v1392
	var v1394 int64
	_ = v1394
	var v1400 int64
	_ = v1400
	var v1403 int64
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1411 int64
	_ = v1411
	var v1412 int64
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1455 int64
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1458 int64
	_ = v1458
	var v1459 int64
	_ = v1459
	var v1460 int64
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int64
	_ = v1466
	var v1467 int64
	_ = v1467
	var v1468 int64
	_ = v1468
	var v1480 int32
	_ = v1480
	var v1492 int64
	_ = v1492
	var v1494 int64
	_ = v1494
	var v1496 int64
	_ = v1496
	var v1503 int64
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int64
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int64
	_ = v1533
	var v1546 int64
	_ = v1546
	var v1548 int64
	_ = v1548
	var v1551 int64
	_ = v1551
	var v1552 int64
	_ = v1552
	var v1558 int64
	_ = v1558
	var v1559 int64
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1563 float64
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1569 float64
	_ = v1569
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1585 float64
	_ = v1585
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1598 float64
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1607 int64
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int64
	_ = v1617
	var v1619 int64
	_ = v1619
	var v1623 int64
	_ = v1623
	var v1641 int64
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1649 int64
	_ = v1649
	var v1655 int64
	_ = v1655
	var v1656 int64
	_ = v1656
	var v1657 int64
	_ = v1657
	var v1659 int64
	_ = v1659
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int64
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1690 int64
	_ = v1690
	var v1703 int64
	_ = v1703
	var v1705 int64
	_ = v1705
	var v1708 int64
	_ = v1708
	var v1709 int64
	_ = v1709
	var v1715 int64
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1719 int64
	_ = v1719
	var v1720 int64
	_ = v1720
	var v1726 int64
	_ = v1726
	var v1740 int64
	_ = v1740
	var v1741 int64
	_ = v1741
	var v1742 int64
	_ = v1742
	var v1743 int64
	_ = v1743
	var v1744 int64
	_ = v1744
	var v1745 int64
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1752 int64
	_ = v1752
	var v1761 int64
	_ = v1761
	var v1762 int64
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1788 int32
	_ = v1788
	var v1798 int32
	_ = v1798
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1831 int64
	_ = v1831
	var v1840 int64
	_ = v1840
	var v1843 int64
	_ = v1843
	var v1844 int64
	_ = v1844
	var v1852 int64
	_ = v1852
	var v1853 int64
	_ = v1853
	var v1857 int64
	_ = v1857
	var v1858 int64
	_ = v1858
	var v1863 int64
	_ = v1863
	var v1865 int64
	_ = v1865
	var v1869 int64
	_ = v1869
	var v1870 int64
	_ = v1870
	var v1871 int64
	_ = v1871
	var v1872 int64
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1876 int64
	_ = v1876
	var v1877 int64
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1886 int64
	_ = v1886
	var v1887 int64
	_ = v1887
	var v1893 int64
	_ = v1893
	var v1894 int64
	_ = v1894
	var v1895 int64
	_ = v1895
	var v1904 int64
	_ = v1904
	var v1905 int64
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1931 int32
	_ = v1931
	var v1941 int32
	_ = v1941
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1973 int64
	_ = v1973
	var v1974 int64
	_ = v1974
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1990 int64
	_ = v1990
	var v1991 int64
	_ = v1991
	var v1999 int64
	_ = v1999
	var v2000 int64
	_ = v2000
	var v2010 int32
	_ = v2010
	var v2013 int32
	_ = v2013
	var v2016 int64
	_ = v2016
	var v2017 int64
	_ = v2017
	var v2018 int64
	_ = v2018
	var v2019 int64
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2028 int64
	_ = v2028
	var v2030 int64
	_ = v2030
	var v2035 int64
	_ = v2035
	var v2036 int64
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2054 int64
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2059 int64
	_ = v2059
	var v2072 int64
	_ = v2072
	var v2074 int64
	_ = v2074
	var v2077 int64
	_ = v2077
	var v2078 int64
	_ = v2078
	var v2086 int64
	_ = v2086
	var v2087 int64
	_ = v2087
	var v2088 int64
	_ = v2088
	var v2089 int64
	_ = v2089
	var v2093 int64
	_ = v2093
	var v2094 int64
	_ = v2094
	var v2098 int64
	_ = v2098
	var v2099 int64
	_ = v2099
	var v2123 int64
	_ = v2123
	var v2129 int64
	_ = v2129
	var v2135 int64
	_ = v2135
	var v2136 int64
	_ = v2136
	var v2137 int64
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2233 int64
	_ = v2233
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2252 int64
	_ = v2252
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2277 int64
	_ = v2277
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2298 int32
	_ = v2298
	var v2307 int64
	_ = v2307
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2345 int64
	_ = v2345
	var v2346 int64
	_ = v2346
	var v2361 int64
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int64
	_ = v2402
	var v2403 int64
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2442 int64
	_ = v2442
	var v2443 int64
	_ = v2443
	var v2450 int64
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2460 int64
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2466 int64
	_ = v2466
	var v2467 int64
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2474 int64
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2481 int32
	_ = v2481
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2501 int64
	_ = v2501
	var v2502 int64
	_ = v2502
	var v2509 int64
	_ = v2509
	var v2512 int32
	_ = v2512
	var v2517 int32
	_ = v2517
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2537 int64
	_ = v2537
	var v2538 int64
	_ = v2538
	var v2579 int64
	_ = v2579
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2590 int32
	_ = v2590
	var v2597 int32
	_ = v2597
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2621 int64
	_ = v2621
	var v2622 int64
	_ = v2622
	var v2629 int32
	_ = v2629
	var v2635 int64
	_ = v2635
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2645 int64
	_ = v2645
	var v2647 int64
	_ = v2647
	var v2651 int64
	_ = v2651
	var v2669 int64
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2677 int64
	_ = v2677
	var v2683 int64
	_ = v2683
	var v2684 int64
	_ = v2684
	var v2685 int64
	_ = v2685
	var v2687 int64
	_ = v2687
	var v2699 int64
	_ = v2699
	var v2700 int64
	_ = v2700
	var v2709 int32
	_ = v2709
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2727 int64
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2732 int64
	_ = v2732
	var v2745 int64
	_ = v2745
	var v2747 int64
	_ = v2747
	var v2750 int64
	_ = v2750
	var v2751 int64
	_ = v2751
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2768 int32
	_ = v2768
	var v2773 int64
	_ = v2773
	var v2782 int64
	_ = v2782
	var v2785 int64
	_ = v2785
	var v2786 int64
	_ = v2786
	var v2794 int64
	_ = v2794
	var v2795 int64
	_ = v2795
	var v2796 int64
	_ = v2796
	var v2797 int64
	_ = v2797
	var v2799 int64
	_ = v2799
	var v2800 int64
	_ = v2800
	var v2809 int32
	_ = v2809
	var v2813 int32
	_ = v2813
	var v2815 int32
	_ = v2815
	var v2818 int32
	_ = v2818
	var v2820 int32
	_ = v2820
	var v2822 int64
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2827 int64
	_ = v2827
	var v2840 int64
	_ = v2840
	var v2842 int64
	_ = v2842
	var v2845 int64
	_ = v2845
	var v2846 int64
	_ = v2846
	var v2854 int64
	_ = v2854
	var v2855 int64
	_ = v2855
	var v2856 int64
	_ = v2856
	var v2857 int64
	_ = v2857
	var v2861 int64
	_ = v2861
	var v2862 int64
	_ = v2862
	var v2866 int64
	_ = v2866
	var v2867 int64
	_ = v2867
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2889 int64
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2894 int64
	_ = v2894
	var v2907 int64
	_ = v2907
	var v2909 int64
	_ = v2909
	var v2912 int64
	_ = v2912
	var v2913 int64
	_ = v2913
	var v2921 int64
	_ = v2921
	var v2922 int64
	_ = v2922
	var v2923 int64
	_ = v2923
	var v2924 int64
	_ = v2924
	var v2928 int64
	_ = v2928
	var v2929 int64
	_ = v2929
	var v2933 int64
	_ = v2933
	var v2934 int64
	_ = v2934
	var v2941 int32
	_ = v2941
	var v2942 int32
	_ = v2942
	var v2948 int32
	_ = v2948
	var v2951 int32
	_ = v2951
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v3017 int32
	_ = v3017
	var v3044 int32
	_ = v3044
	var v3050 int32
	_ = v3050
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int64
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3068 int64
	_ = v3068
	var v3081 int64
	_ = v3081
	var v3083 int64
	_ = v3083
	var v3086 int64
	_ = v3086
	var v3087 int64
	_ = v3087
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3105 int32
	_ = v3105
	var v3110 int64
	_ = v3110
	var v3119 int64
	_ = v3119
	var v3122 int64
	_ = v3122
	var v3123 int64
	_ = v3123
	var v3131 int64
	_ = v3131
	var v3132 int64
	_ = v3132
	var v3133 int64
	_ = v3133
	var v3134 int64
	_ = v3134
	var v3136 int64
	_ = v3136
	var v3137 int64
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3154 int64
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3159 int64
	_ = v3159
	var v3172 int64
	_ = v3172
	var v3174 int64
	_ = v3174
	var v3177 int64
	_ = v3177
	var v3178 int64
	_ = v3178
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3189 int32
	_ = v3189
	var v3191 int32
	_ = v3191
	var v3196 int32
	_ = v3196
	var v3201 int64
	_ = v3201
	var v3210 int64
	_ = v3210
	var v3213 int64
	_ = v3213
	var v3214 int64
	_ = v3214
	var v3222 int64
	_ = v3222
	var v3223 int64
	_ = v3223
	var v3224 int64
	_ = v3224
	var v3225 int64
	_ = v3225
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3237 int32
	_ = v3237
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3244 int32
	_ = v3244
	var v3246 int64
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3251 int64
	_ = v3251
	var v3264 int64
	_ = v3264
	var v3266 int64
	_ = v3266
	var v3269 int64
	_ = v3269
	var v3270 int64
	_ = v3270
	var v3278 int64
	_ = v3278
	var v3279 int64
	_ = v3279
	var v3280 int64
	_ = v3280
	var v3281 int64
	_ = v3281
	var v3283 int64
	_ = v3283
	var v3284 int64
	_ = v3284
	var v3289 int32
	_ = v3289
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3306 int32
	_ = v3306
	var v3308 int32
	_ = v3308
	var v3310 int64
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3315 int64
	_ = v3315
	var v3328 int64
	_ = v3328
	var v3330 int64
	_ = v3330
	var v3333 int64
	_ = v3333
	var v3334 int64
	_ = v3334
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3351 int32
	_ = v3351
	var v3356 int64
	_ = v3356
	var v3365 int64
	_ = v3365
	var v3368 int64
	_ = v3368
	var v3369 int64
	_ = v3369
	var v3377 int64
	_ = v3377
	var v3378 int64
	_ = v3378
	var v3379 int64
	_ = v3379
	var v3380 int64
	_ = v3380
	var v3383 int32
	_ = v3383
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3394 int32
	_ = v3394
	var v3397 int32
	_ = v3397
	var v3399 int32
	_ = v3399
	var v3401 int64
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3406 int64
	_ = v3406
	var v3419 int64
	_ = v3419
	var v3421 int64
	_ = v3421
	var v3424 int64
	_ = v3424
	var v3425 int64
	_ = v3425
	var v3433 int64
	_ = v3433
	var v3434 int64
	_ = v3434
	var v3435 int64
	_ = v3435
	var v3436 int64
	_ = v3436
	var v3438 int64
	_ = v3438
	var v3439 int64
	_ = v3439
	var v3449 int32
	_ = v3449
	var v3477 int32
	_ = v3477
	var v3480 int32
	_ = v3480
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3493 int32
	_ = v3493
	var v3496 int32
	_ = v3496
	var v3499 int32
	_ = v3499
	var v3505 int32
	_ = v3505
	var v3506 int32
	_ = v3506
	var v3513 int32
	_ = v3513
	var v3515 int32
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3559 int32
	_ = v3559
	var v3562 int32
	_ = v3562
	var v3564 int32
	_ = v3564
	var v3572 int32
	_ = v3572
	var v3574 int32
	_ = v3574
	var v3579 int32
	_ = v3579
	var v3604 int32
	_ = v3604
	var v3606 int32
	_ = v3606
	var v3611 int32
	_ = v3611
	var v3633 int32
	_ = v3633
	var v3635 int32
	_ = v3635
	var v3640 int32
	_ = v3640
	var v3643 int32
	_ = v3643
	var v3669 int32
	_ = v3669
	var v3679 int32
	_ = v3679
	var v3701 int32
	_ = v3701
	var v3708 int32
	_ = v3708
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
	var v3740 int32
	_ = v3740
	var v3743 int32
	_ = v3743
	var v3744 int64
	_ = v3744
	var v3747 int64
	_ = v3747
	var v3751 int64
	_ = v3751
	var v3752 int64
	_ = v3752
	var v3757 int64
	_ = v3757
	var v3759 int32
	_ = v3759
	var v3763 int32
	_ = v3763
	var v3765 int32
	_ = v3765
	var v3769 int32
	_ = v3769
	var v3771 int32
	_ = v3771
	var v3776 int32
	_ = v3776
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3790 int32
	_ = v3790
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3797 int32
	_ = v3797
	var v3800 int32
	_ = v3800
	var v3811 int32
	_ = v3811
	var v3813 int32
	_ = v3813
	var v3818 int32
	_ = v3818
	var v3821 int32
	_ = v3821
	var v3839 int32
	_ = v3839
	var v3841 int32
	_ = v3841
	var v3842 int32
	_ = v3842
	var v3851 int32
	_ = v3851
	var v3855 int32
	_ = v3855
	var v3860 int32
	_ = v3860
	var v3863 int32
	_ = v3863
	var v3885 int32
	_ = v3885
	var v3889 int32
	_ = v3889
	var v3897 int32
	_ = v3897
	var v3921 int32
	_ = v3921
	var v3947 int32
	_ = v3947
	var v3951 int32
	_ = v3951
	var v3954 int32
	_ = v3954
	var v3957 int32
	_ = v3957
	var v3961 int32
	_ = v3961
	var v3969 int64
	_ = v3969
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3993 int64
	_ = v3993
	var v3994 int64
	_ = v3994
	var v4003 int32
	_ = v4003
	var v4008 int32
	_ = v4008
	var v4014 int32
	_ = v4014
	var v4016 int32
	_ = v4016
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4027 int32
	_ = v4027
	var v4032 int32
	_ = v4032
	var v4037 int64
	_ = v4037
	var v4046 int64
	_ = v4046
	var v4049 int64
	_ = v4049
	var v4050 int64
	_ = v4050
	var v4063 int64
	_ = v4063
	var v4064 int64
	_ = v4064
	var v4065 int64
	_ = v4065
	var v4066 int64
	_ = v4066
	var v4068 int64
	_ = v4068
	var v4069 int64
	_ = v4069
	var v4071 int32
	_ = v4071
	var v4075 int32
	_ = v4075
	var v4079 int32
	_ = v4079
	var v4081 int32
	_ = v4081
	var v4084 int32
	_ = v4084
	var v4086 int32
	_ = v4086
	var v4088 int64
	_ = v4088
	var v4089 int32
	_ = v4089
	var v4093 int64
	_ = v4093
	var v4106 int64
	_ = v4106
	var v4108 int64
	_ = v4108
	var v4111 int64
	_ = v4111
	var v4112 int64
	_ = v4112
	var v4120 int64
	_ = v4120
	var v4121 int64
	_ = v4121
	var v4124 int64
	_ = v4124
	var v4125 int64
	_ = v4125
	var v4127 int32
	_ = v4127
	var v4128 int32
	_ = v4128
	var v4129 int32
	_ = v4129
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4134 int32
	_ = v4134
	var v4140 int32
	_ = v4140
	var v4144 int32
	_ = v4144
	var v4152 int32
	_ = v4152
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4157 int32
	_ = v4157
	var v4179 int32
	_ = v4179
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4185 int32
	_ = v4185
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4198 int32
	_ = v4198
	var v4200 int32
	_ = v4200
	var v4204 int32
	_ = v4204
	var v4213 int32
	_ = v4213
	var v4218 int32
	_ = v4218
	var v4219 float64
	_ = v4219
	var v4221 int32
	_ = v4221
	var v4225 float64
	_ = v4225
	var v4232 int32
	_ = v4232
	var v4235 int32
	_ = v4235
	var v4241 float64
	_ = v4241
	var v4248 int32
	_ = v4248
	var v4251 int32
	_ = v4251
	var v4254 float64
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4263 int64
	_ = v4263
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4273 int64
	_ = v4273
	var v4275 int64
	_ = v4275
	var v4279 int64
	_ = v4279
	var v4297 int64
	_ = v4297
	var v4301 int32
	_ = v4301
	var v4305 int64
	_ = v4305
	var v4311 int64
	_ = v4311
	var v4312 int64
	_ = v4312
	var v4313 int64
	_ = v4313
	var v4315 int64
	_ = v4315
	var v4328 int32
	_ = v4328
	var v4329 int64
	_ = v4329
	var v4330 int64
	_ = v4330
	var v4336 int64
	_ = v4336
	var v4350 int64
	_ = v4350
	var v4351 int64
	_ = v4351
	var v4353 int32
	_ = v4353
	var v4354 float64
	_ = v4354
	var v4356 int32
	_ = v4356
	var v4360 float64
	_ = v4360
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4376 float64
	_ = v4376
	var v4383 int32
	_ = v4383
	var v4386 int32
	_ = v4386
	var v4389 float64
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4398 int64
	_ = v4398
	var v4404 int32
	_ = v4404
	var v4406 int32
	_ = v4406
	var v4408 int64
	_ = v4408
	var v4410 int64
	_ = v4410
	var v4414 int64
	_ = v4414
	var v4432 int64
	_ = v4432
	var v4436 int32
	_ = v4436
	var v4440 int64
	_ = v4440
	var v4446 int64
	_ = v4446
	var v4447 int64
	_ = v4447
	var v4448 int64
	_ = v4448
	var v4450 int64
	_ = v4450
	var v4464 int64
	_ = v4464
	var v4465 int64
	_ = v4465
	var v4468 int32
	_ = v4468
	var v4469 int64
	_ = v4469
	var v4470 int64
	_ = v4470
	var v4472 int32
	_ = v4472
	var v4473 int32
	_ = v4473
	var v4474 int32
	_ = v4474
	var v4479 int64
	_ = v4479
	var v4480 int64
	_ = v4480
	var v4488 int64
	_ = v4488
	var v4489 int64
	_ = v4489
	var v4491 int64
	_ = v4491
	var v4492 int64
	_ = v4492
	var v4493 int64
	_ = v4493
	var v4494 int64
	_ = v4494
	var v4495 int64
	_ = v4495
	var v4496 int64
	_ = v4496
	var v4497 int64
	_ = v4497
	var v4498 int64
	_ = v4498
	var v4502 int32
	_ = v4502
	var v4509 int32
	_ = v4509
	var v4521 int32
	_ = v4521
	var v4525 int64
	_ = v4525
	var v4531 int32
	_ = v4531
	var v4533 int32
	_ = v4533
	var v4535 int64
	_ = v4535
	var v4537 int64
	_ = v4537
	var v4541 int64
	_ = v4541
	var v4559 int64
	_ = v4559
	var v4563 int32
	_ = v4563
	var v4567 int64
	_ = v4567
	var v4573 int64
	_ = v4573
	var v4574 int64
	_ = v4574
	var v4575 int64
	_ = v4575
	var v4577 int64
	_ = v4577
	var v4591 int64
	_ = v4591
	var v4592 int64
	_ = v4592
	var v4594 int64
	_ = v4594
	var v4595 int64
	_ = v4595
	var v4599 int32
	_ = v4599
	var v4603 int64
	_ = v4603
	var v4609 int32
	_ = v4609
	var v4611 int32
	_ = v4611
	var v4613 int64
	_ = v4613
	var v4615 int64
	_ = v4615
	var v4619 int64
	_ = v4619
	var v4637 int64
	_ = v4637
	var v4641 int32
	_ = v4641
	var v4645 int64
	_ = v4645
	var v4651 int64
	_ = v4651
	var v4652 int64
	_ = v4652
	var v4653 int64
	_ = v4653
	var v4655 int64
	_ = v4655
	var v4669 int64
	_ = v4669
	var v4670 int64
	_ = v4670
	var v4672 int64
	_ = v4672
	var v4673 int64
	_ = v4673
	var v4674 float64
	_ = v4674
	var v4681 int32
	_ = v4681
	var v4684 int64
	_ = v4684
	var v4690 int32
	_ = v4690
	var v4692 int32
	_ = v4692
	var v4694 int64
	_ = v4694
	var v4696 int64
	_ = v4696
	var v4700 int64
	_ = v4700
	var v4718 int64
	_ = v4718
	var v4722 int32
	_ = v4722
	var v4726 int64
	_ = v4726
	var v4732 int64
	_ = v4732
	var v4733 int64
	_ = v4733
	var v4734 int64
	_ = v4734
	var v4736 int64
	_ = v4736
	var v4750 int64
	_ = v4750
	var v4751 int64
	_ = v4751
	var v4753 int64
	_ = v4753
	var v4754 int64
	_ = v4754
	var v4756 int32
	_ = v4756
	var v4759 int64
	_ = v4759
	var v4765 int32
	_ = v4765
	var v4767 int32
	_ = v4767
	var v4769 int64
	_ = v4769
	var v4771 int64
	_ = v4771
	var v4775 int64
	_ = v4775
	var v4793 int64
	_ = v4793
	var v4797 int32
	_ = v4797
	var v4801 int64
	_ = v4801
	var v4807 int64
	_ = v4807
	var v4808 int64
	_ = v4808
	var v4809 int64
	_ = v4809
	var v4811 int64
	_ = v4811
	var v4825 int64
	_ = v4825
	var v4826 int64
	_ = v4826
	var v4828 int64
	_ = v4828
	var v4829 int64
	_ = v4829
	var v4830 int64
	_ = v4830
	var v4831 int64
	_ = v4831
	var v4837 int64
	_ = v4837
	var v4840 int64
	_ = v4840
	var v4841 int64
	_ = v4841
	var v4851 int64
	_ = v4851
	var v4852 int64
	_ = v4852
	var v4856 int32
	_ = v4856
	var v4878 int32
	_ = v4878
	var v4888 int32
	_ = v4888
	var v4895 int32
	_ = v4895
	var v4899 int32
	_ = v4899
	var v4905 int64
	_ = v4905
	var v4906 int64
	_ = v4906
	var v4908 int64
	_ = v4908
	var v4909 int64
	_ = v4909
	var v4915 int32
	_ = v4915
	var v4916 int64
	_ = v4916
	var v4917 int64
	_ = v4917
	var v4919 int32
	_ = v4919
	var v4920 int32
	_ = v4920
	var v4921 int32
	_ = v4921
	var v4926 int64
	_ = v4926
	var v4927 int64
	_ = v4927
	var v4933 int64
	_ = v4933
	var v4934 int64
	_ = v4934
	var v4941 int32
	_ = v4941
	var v4942 int64
	_ = v4942
	var v4948 int64
	_ = v4948
	var v4951 int64
	_ = v4951
	var v4952 int64
	_ = v4952
	var v4953 int64
	_ = v4953
	var v4957 int32
	_ = v4957
	var v4961 int64
	_ = v4961
	var v4962 int64
	_ = v4962
	var v4966 int32
	_ = v4966
	var v4993 int32
	_ = v4993
	var v4998 int32
	_ = v4998
	var v5002 int32
	_ = v5002
	var v5003 int64
	_ = v5003
	var v5005 int32
	_ = v5005
	var v5006 int64
	_ = v5006
	var v5007 int64
	_ = v5007
	var v5008 int64
	_ = v5008
	var v5014 int64
	_ = v5014
	var v5023 int64
	_ = v5023
	var v5024 int64
	_ = v5024
	var v5028 int32
	_ = v5028
	var v5050 int32
	_ = v5050
	var v5060 int32
	_ = v5060
	var v5067 int32
	_ = v5067
	var v5071 int32
	_ = v5071
	var v5072 int32
	_ = v5072
	var v5077 int32
	_ = v5077
	var v5087 int32
	_ = v5087
	var v5088 int64
	_ = v5088
	var v5089 int64
	_ = v5089
	var v5091 int32
	_ = v5091
	var v5093 int32
	_ = v5093
	var v5095 int32
	_ = v5095
	var v5104 int64
	_ = v5104
	var v5105 int64
	_ = v5105
	var v5115 int32
	_ = v5115
	var v5118 int32
	_ = v5118
	var v5121 int64
	_ = v5121
	var v5122 int64
	_ = v5122
	var v5130 int64
	_ = v5130
	var v5131 int64
	_ = v5131
	var v5141 int32
	_ = v5141
	var v5144 int32
	_ = v5144
	var v5147 int64
	_ = v5147
	var v5148 int64
	_ = v5148
	var v5149 int64
	_ = v5149
	var v5150 int64
	_ = v5150
	var v5151 int32
	_ = v5151
	var v5159 int64
	_ = v5159
	var v5161 int64
	_ = v5161
	var v5166 int64
	_ = v5166
	var v5167 int64
	_ = v5167
	var v5190 int64
	_ = v5190
	var v5197 int64
	_ = v5197
	var v5203 int64
	_ = v5203
	var v5204 int64
	_ = v5204
	var v5256 int64
	_ = v5256
	var v5286 int64
	_ = v5286
	var v5287 int64
	_ = v5287
	v5 = int32(0)
	v22 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(48)
	m.G0 = v32
	if base.Ui32(int32(2)) < base.Ui32(l2) {
		v5256 = v22
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v5286
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v5287
	m.G0 = v32 + int32(48)
	return
L2:
	;
	v5286 = v5256
	v5287 = int64(0)
	goto L1
L3:
	;
	v37 = l2 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F___floatscan[0])))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F___floatscan[1])))
	goto L4
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v69 != v70 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	switch v78 - int32(43) {
	case 0, 2:
		goto L15
	default:
		v103 = v78
		v104 = int32(1)
		goto L14
	}
L6:
	;
	goto L12
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v69 + int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v78 = v75
	goto L6
L8:
	;
	goto L9
L9:
	;
	v76 = F___shgetc(m, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v78 = v76
	goto L6
L12:
	;
	if base.B2i32(v78 == int32(32))|base.B2i32(base.Ui32(v78-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	if v103&int32(-33) == int32(73) {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	if v78 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v93 = int32(-1)
	goto L18
L17:
	;
	v93 = int32(1)
	goto L18
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v94 != v95 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v94 + int32(1)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v103 = v100
	v104 = v93
	goto L14
L20:
	;
	goto L21
L21:
	;
	v101 = F___shgetc(m, l1)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v103 = v101
	v104 = v93
	goto L14
L23:
	;
	if v163 != 0 {
		v418 = v158
		v423 = v163
		goto L72
	} else {
		goto L73
	}
L24:
	;
	v278 = m.G0
	v280 = v278 - int32(16)
	m.G0 = v280
	v285 = base.I32_reinterpret_f32(base.F32_mul(base.F32_convert_i32_s(v104), math.Float32frombits(uint32(0x7f800000))))
	v287 = v285 & int32(_a_F___floatscan_0)
	v289 = int32(base.Ui32(v285) >> (uint(int32(23)) % 32))
	v291 = v289 & int32(255)
	if v291 != 0 {
		goto L54
	} else {
		goto L55
	}
L25:
	;
	v114 = v5
	goto L28
L26:
	;
	v158 = v103
	v163 = v5
	goto L27
L27:
	;
	if v163 != int32(3) {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	if v114 == int32(7) {
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v158 = v149
	v163 = v152
	goto L27
L30:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v140 != v141 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F___floatscan[2]))))
	v152 = v114 + int32(1)
	if v150 == v149|int32(32) {
		v114 = v152
		goto L28
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v140 + int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v149 = v146
	goto L31
L33:
	;
	goto L34
L34:
	;
	v147 = F___shgetc(m, l1)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v149 = v147
	goto L31
L36:
	;
	goto L29
L37:
	;
	v188 = base.B2i32(v163 == int32(8))
	if v163 == int32(8) {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v195 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if base.B2i32(l3 == int32(0))|base.B2i32(base.Ui32(v163) < base.Ui32(int32(4))) != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	if v163 == int32(8) {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v198 - int32(1)
	goto L45
L44:
	;
	goto L45
L45:
	;
	if base.B2i32(l3 == int32(0))|base.B2i32(base.Ui32(v163) < base.Ui32(int32(4))) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v216 = v163
	goto L47
L47:
	;
	if base.B2i32(v195 < int64(0)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L24
L49:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v240 - int32(1)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v245 = v216 - int32(1)
	if base.Ui32(int32(3)) < base.Ui32(v245) {
		v216 = v245
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v344
	v356 = base.I64_extend_i32_u(v345)<<(uint(int64(48))%64) | base.I64_extend_i32_u(int32(base.Ui32(v285)>>(uint(int32(31))%32)))<<(uint(int64(63))%64) | v343
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v356
	m.G0 = v280 + int32(16)
	v5286 = v344
	v5287 = v356
	goto L1
L54:
	;
	if v291 != int32(255) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v305 = int32(0)
	if v287 == v305 {
		v343 = int64(0)
		v344 = v22
		v345 = v305
		goto L53
	} else {
		goto L60
	}
L57:
	;
	v343 = base.I64_extend_i32_u(v287) << (uint(int64(25)) % 64)
	v344 = v22
	v345 = v289&int32(255) + int32(_a_F___floatscan_1)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v343 = base.I64_extend_i32_u(v287) << (uint(int64(25)) % 64)
	v344 = v22
	v345 = int32(_a_F___floatscan_2)
	goto L53
L60:
	;
	v308 = base.I64_extend_i32_u(v287)
	v309 = int64(0)
	v310 = base.I32_clz(v287)
	v312 = v310 + int32(81)
	if v312&int32(64) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v280)+8))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	v343 = v336 ^ int64(281474976710656)
	v344 = v339
	v345 = int32(_a_F___floatscan_3) - v310
	goto L53
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v331
	*(*int64)(unsafe.Add(mBase, uint32(v280)+8)) = v332
	goto L61
L63:
	;
	v331 = int64(0)
	v332 = v308 << (uint(base.I64_extend_i32_u(v310+int32(17))) % 64)
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v312 == int32(0) {
		v331 = v308
		v332 = v309
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v323 = base.I64_extend_i32_u(v312)
	v331 = v308 << (uint(v323) % 64)
	v332 = v309<<(uint(v323)%64) | int64(base.Ui64(v308)>>(uint(base.I64_extend_i32_u(int32(64)-v312))%64))
	goto L62
L67:
	;
	v5286 = int64(0)
	v5287 = v557
	goto L1
L68:
	;
	if v418 != int32(48) {
		goto L128
	} else {
		goto L129
	}
L69:
	;
	v647 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v647
	v650 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v650 - v651)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v657-v651) <= v647) != 0 {
		goto L125
	} else {
		goto L126
	}
L70:
	;
	v608 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v608 {
		goto L121
	} else {
		goto L122
	}
L71:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v474 != v475 {
		goto L88
	} else {
		goto L89
	}
L72:
	;
	switch v423 {
	case 0:
		goto L68
	default:
		goto L70
	case 3:
		goto L71
	}
L73:
	;
	if v158&int32(-33) != int32(78) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v418 = v158
	v423 = int32(0)
	goto L72
L75:
	;
	goto L76
L76:
	;
	v374 = int32(0)
	goto L77
L77:
	;
	if v374 == int32(2) {
		goto L71
	} else {
		goto L79
	}
L78:
	;
	v418 = v409
	v423 = v412
	goto L72
L79:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v400 != v401 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v410 = int32(*(*int8)(unsafe.Add(mBase, uint32(v374)+uint32(_c_F___floatscan[3]))))
	v412 = v374 + int32(1)
	if v410 == v409|int32(32) {
		v374 = v412
		goto L77
	} else {
		goto L85
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v400 + int32(1)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	v409 = v406
	goto L80
L82:
	;
	goto L83
L83:
	;
	v407 = F___shgetc(m, l1)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v409 = v407
	goto L80
L85:
	;
	goto L78
L86:
	;
	v500 = int32(1)
	goto L96
L87:
	;
	if v483 == int32(40) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v474 + int32(1)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v483 = v480
	goto L87
L89:
	;
	goto L90
L90:
	;
	v481 = F___shgetc(m, l1)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v483 = v481
	goto L87
L92:
	;
	goto L86
L93:
	;
	goto L94
L94:
	;
	v487 = int64(9223231299366420480)
	v488 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v488 < int64(0) {
		v5286 = v22
		v5287 = v487
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v491 - int32(1)
	v5286 = v22
	v5287 = v487
	goto L1
L96:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v524 != v525 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v557 = int64(9223231299366420480)
	if v533 == int32(41) {
		v5286 = v22
		v5287 = v557
		goto L1
	} else {
		goto L106
	}
L98:
	;
	v540 = int32(26)
	v546 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(v533-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v533-int32(65)) < base.Ui32(v540))|base.B2i32(v533 == int32(95)) == v546)&base.B2i32(base.Ui32(v540) <= base.Ui32(v533-int32(97))) == v546 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v524 + int32(1)
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v533 = v530
	goto L98
L100:
	;
	goto L101
L101:
	;
	v531 = F___shgetc(m, l1)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v533 = v531
	goto L98
L103:
	;
	v500 = v500 + int32(1)
	goto L96
L104:
	;
	goto L105
L105:
	;
	goto L97
L106:
	;
	v560 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v560 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v563 - int32(1)
	goto L109
L108:
	;
	goto L109
L109:
	;
	if l3 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v576 = v500
	goto L115
L111:
	;
	if v500 != 0 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(28)
	v639 = int64(0)
	goto L69
L114:
	;
	goto L67
L115:
	;
	if int64(0) <= v560 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L67
L117:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v602 - int32(1)
	goto L119
L118:
	;
	goto L119
L119:
	;
	v607 = v576 - int32(1)
	if v607 != 0 {
		v576 = v607
		goto L115
	} else {
		goto L120
	}
L120:
	;
	goto L116
L121:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v611 - int32(1)
	goto L123
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(28)
	v639 = v22
	goto L69
L124:
	;
	v5256 = v639
	goto L2
L125:
	;
	v664 = v657
	goto L127
L126:
	;
	v664 = v651 + base.I32_wrap_i64(v647)
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v664
	goto L124
L128:
	;
	v2146 = v32 + int32(32)
	v2147 = int32(0)
	v2149 = m.G0
	v2151 = v2149 - int32(_a_F___floatscan_4)
	m.G0 = v2151
	v2154 = v2147 - v38
	v2155 = v2154 - v39
	v2158 = v418
	v2162 = v2147
	goto L462
L129:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v668 != v669 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v677&int32(-33) == int32(88) {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v668 + int32(1)
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	v677 = v674
	goto L130
L132:
	;
	goto L133
L133:
	;
	v675 = F___shgetc(m, l1)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	v677 = v675
	goto L130
L135:
	;
	v683 = v32 + int32(16)
	v685 = m.G0
	v687 = v685 - int32(432)
	m.G0 = v687
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v689 != v690 {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	v2137 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2137 < int64(0) {
		goto L128
	} else {
		goto L459
	}
L138:
	;
	v701 = v698
	v703 = v5
	goto L145
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v689 + int32(1)
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689))))
	v698 = v695
	goto L138
L140:
	;
	goto L141
L141:
	;
	v696 = F___shgetc(m, l1)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	v698 = v696
	goto L138
L143:
	;
	v833 = v803
	v835 = v805
	v839 = v5
	v840 = int32(0)
	v841 = v811
	v852 = v22
	v853 = int64(4611404543450677248)
	v854 = v22
	v855 = v825
	v856 = v22
	v857 = v22
	goto L169
L144:
	;
	if v752 != int32(48) {
		goto L158
	} else {
		goto L159
	}
L145:
	;
	if v701 != int32(48) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v750 = F___shgetc(m, l1)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L10
	} else {
		goto L157
	}
L147:
	;
	goto L146
L148:
	;
	if v701 != int32(46) {
		v803 = v701
		v805 = v703
		v811 = v5
		v825 = v22
		goto L143
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v739 != v740 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v732 == v733 {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v732 + int32(1)
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732))))
	v752 = v738
	goto L144
L153:
	;
	v742 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v739 + v742
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	v701 = v746
	v703 = v742
	goto L145
L154:
	;
	goto L155
L155:
	;
	v748 = F___shgetc(m, l1)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L156
	}
L156:
	;
	v701 = v748
	v703 = int32(1)
	goto L145
L157:
	;
	v752 = v750
	goto L144
L158:
	;
	v803 = v752
	v805 = v703
	v811 = int32(1)
	v825 = v22
	goto L143
L159:
	;
	goto L160
L160:
	;
	v780 = v22
	goto L161
L161:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v785 != v786 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v799 = int32(1)
	v803 = v794
	v805 = v799
	v811 = v799
	v825 = v796
	goto L143
L163:
	;
	v796 = v780 - int64(1)
	if v794 == int32(48) {
		v780 = v796
		goto L161
	} else {
		goto L168
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v785 + int32(1)
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785))))
	v794 = v791
	goto L163
L165:
	;
	goto L166
L166:
	;
	v792 = F___shgetc(m, l1)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	v794 = v792
	goto L163
L168:
	;
	goto L162
L169:
	;
	v861 = v833 - int32(48)
	if base.Ui32(v861) < base.Ui32(int32(10)) {
		v874 = v833
		goto L173
	} else {
		goto L174
	}
L170:
	;
	if v835 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L171:
	;
	goto L170
L172:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v987 != v988 {
		goto L196
	} else {
		goto L197
	}
L173:
	;
	if int32(57) < v833 {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v865 = base.B2i32(v833 != int32(46))
	v867 = v833 | int32(32)
	if v865&base.B2i32(base.Ui32(int32(5)) < base.Ui32(v867-int32(97))) != 0 {
		goto L171
	} else {
		goto L175
	}
L175:
	;
	if v833 != int32(46) {
		v874 = v867
		goto L173
	} else {
		goto L176
	}
L176:
	;
	if v841 != 0 {
		goto L171
	} else {
		goto L177
	}
L177:
	;
	v975 = v835
	v977 = v839
	v978 = v840
	v979 = int32(1)
	v981 = v852
	v982 = v853
	v983 = v854
	v984 = v852
	v985 = v856
	v986 = v857
	goto L172
L178:
	;
	v880 = v874 - int32(87)
	goto L180
L179:
	;
	v880 = v861
	goto L180
L180:
	;
	if v852 <= int64(7) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v975 = int32(1)
	v977 = v965
	v978 = v966
	v979 = v841
	v981 = v852 + int64(1)
	v982 = v967
	v983 = v968
	v984 = v855
	v985 = v969
	v986 = v970
	goto L172
L182:
	;
	v965 = v839
	v966 = v880 + v840<<(uint(int32(4))%32)
	v967 = v853
	v968 = v854
	v969 = v856
	v970 = v857
	goto L181
L183:
	;
	goto L184
L184:
	;
	if base.Ui64(v852) <= base.Ui64(int64(28)) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v889 = v687 + int32(48)
	v893 = m.G0
	v895 = v893 - int32(16)
	m.G0 = v895
	if v880 != 0 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	if base.B2i32(v880 == int32(0))|v839 != 0 {
		v965 = v839
		v966 = v840
		v967 = v853
		v968 = v854
		v969 = v856
		v970 = v857
		goto L181
	} else {
		goto L195
	}
L188:
	;
	F___multf3(m, v687+int32(32), v857, v853, int64(0), int64(4610278643543834624))
	mBase = m.M
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v687)+48))
	v940 = *(*int64)(unsafe.Add(mBase, uint32(v687)+56))
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v687)+32))
	v942 = *(*int64)(unsafe.Add(mBase, uint32(v687)+40))
	F___multf3(m, v687+int32(16), v939, v940, v941, v942)
	mBase = m.M
	v944 = *(*int64)(unsafe.Add(mBase, uint32(v687)+16))
	v945 = *(*int64)(unsafe.Add(mBase, uint32(v687)+24))
	F___addtf3(m, v687, v944, v945, v854, v856)
	mBase = m.M
	v947 = *(*int64)(unsafe.Add(mBase, uint32(v687)+8))
	v948 = *(*int64)(unsafe.Add(mBase, uint32(v687)))
	v965 = v839
	v966 = v840
	v967 = v942
	v968 = v948
	v969 = v947
	v970 = v941
	goto L181
L189:
	;
	v898 = v880 >> (uint(int32(31)) % 32)
	v900 = v880 ^ v898 - v898
	v902 = int64(0)
	v903 = base.I32_clz(v900)
	F___ashlti3(m, v895, base.I64_extend_i32_u(v900), v902, v903+int32(81))
	mBase = m.M
	v907 = *(*int64)(unsafe.Add(mBase, uint32(v895)+8))
	if v880 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v925 = int64(0)
	v926 = int64(0)
	goto L191
L191:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v889))) = v926
	*(*int64)(unsafe.Add(mBase, uint32(v889)+8)) = v925
	m.G0 = v895 + int32(16)
	goto L188
L192:
	;
	v920 = int64(-9223372036854775807 - 1)
	goto L194
L193:
	;
	v920 = v902
	goto L194
L194:
	;
	v922 = *(*int64)(unsafe.Add(mBase, uint32(v895)))
	v925 = v907 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v903)<<(uint(int64(48))%64) | v920
	v926 = v922
	goto L191
L195:
	;
	F___multf3(m, v687+int32(80), v857, v853, int64(0), int64(4611123068473966592))
	mBase = m.M
	v959 = *(*int64)(unsafe.Add(mBase, uint32(v687)+80))
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v687)+88))
	F___addtf3(m, v687-int32(-64), v959, v960, v854, v856)
	mBase = m.M
	v963 = *(*int64)(unsafe.Add(mBase, uint32(v687)+72))
	v964 = *(*int64)(unsafe.Add(mBase, uint32(v687)+64))
	v965 = int32(1)
	v966 = v840
	v967 = v853
	v968 = v964
	v969 = v963
	v970 = v857
	goto L181
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v987 + int32(1)
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987))))
	v833 = v993
	v835 = v975
	v839 = v977
	v840 = v978
	v841 = v979
	v852 = v981
	v853 = v982
	v854 = v983
	v855 = v984
	v856 = v985
	v857 = v986
	goto L169
L197:
	;
	goto L198
L198:
	;
	v994 = F___shgetc(m, l1)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	v833 = v994
	v835 = v975
	v839 = v977
	v840 = v978
	v841 = v979
	v852 = v981
	v853 = v982
	v854 = v983
	v855 = v984
	v856 = v985
	v857 = v986
	goto L169
L200:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v683))) = v2123
	*(*int64)(unsafe.Add(mBase, uint32(v683)+8)) = v2129
	m.G0 = v687 + int32(432)
	v2135 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
	v2136 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	v5286 = v2136
	v5287 = v2135
	goto L1
L201:
	;
	v998 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v998 {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	goto L203
L203:
	;
	if v852 <= int64(7) {
		goto L227
	} else {
		goto L228
	}
L204:
	;
	v1037 = v687 + int32(96)
	v1041 = int64(0)
	v1047 = m.G0
	v1049 = v1047 - int32(16)
	m.G0 = v1049
	v1051 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v104)))
	v1053 = v1051 & int64(4503599627370495)
	v1057 = int64(base.Ui64(v1051)>>(uint(int64(52))%64)) & int64(2047)
	if v1057 != v1041 {
		goto L218
	} else {
		goto L219
	}
L205:
	;
	v1016 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v1016
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v1019 - v1020)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v1026-v1020) <= v1016) != 0 {
		goto L213
	} else {
		goto L214
	}
L206:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1001 - int32(1)
	if l3 == int32(0) {
		goto L205
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	if l3 != 0 {
		goto L204
	} else {
		goto L211
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1001 - int32(2)
	if v841 == int32(0) {
		goto L204
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1001 - int32(3)
	goto L204
L211:
	;
	goto L205
L212:
	;
	goto L204
L213:
	;
	v1033 = v1026
	goto L215
L214:
	;
	v1033 = v1020 + base.I32_wrap_i64(v1016)
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v1033
	goto L212
L216:
	;
	v1105 = *(*int64)(unsafe.Add(mBase, uint32(v687)+96))
	v1106 = *(*int64)(unsafe.Add(mBase, uint32(v687)+104))
	v2123 = v1105
	v2129 = v1106
	goto L200
L217:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1037))) = v1093
	*(*int64)(unsafe.Add(mBase, uint32(v1037)+8)) = v1051&int64(-9223372036854775807-1) | v1090<<(uint(int64(48))%64) | v1091
	m.G0 = v1049 + int32(16)
	goto L216
L218:
	;
	if v1057 != int64(2047) {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	if v1053 == int64(0) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	v1090 = v1057 + int64(15360)
	v1091 = int64(base.Ui64(v1053) >> (uint(int64(4)) % 64))
	v1093 = v1053 << (uint(int64(60)) % 64)
	goto L217
L222:
	;
	goto L223
L223:
	;
	v1090 = int64(32767)
	v1091 = int64(base.Ui64(v1053) >> (uint(int64(4)) % 64))
	v1093 = v1053 << (uint(int64(60)) % 64)
	goto L217
L224:
	;
	v1075 = int64(0)
	v1090 = v1075
	v1091 = v1041
	v1093 = v1075
	goto L217
L225:
	;
	goto L226
L226:
	;
	v1079 = base.I32_wrap_i64(base.I64_clz(v1053))
	F___ashlti3(m, v1049, v1053, int64(0), v1079+int32(49))
	mBase = m.M
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v1049)+8))
	v1089 = *(*int64)(unsafe.Add(mBase, uint32(v1049)))
	v1090 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v1079)
	v1091 = v1083 ^ int64(281474976710656)
	v1093 = v1089
	goto L217
L227:
	;
	v1118 = v840
	v1131 = v852
	goto L230
L228:
	;
	v1153 = v840
	goto L229
L229:
	;
	if v833&int32(-33) == int32(80) {
		goto L236
	} else {
		goto L237
	}
L230:
	;
	v1139 = v1118 << (uint(int32(4)) % 32)
	v1141 = v1131 + int64(1)
	if v1141 != int64(8) {
		v1118 = v1139
		v1131 = v1141
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v1153 = v1139
	goto L229
L232:
	;
	goto L231
L233:
	;
	if v1153 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L234:
	;
	v1216 = int64(0)
	goto L233
L235:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1210 - int32(1)
	goto L234
L236:
	;
	v1177 = F_scanexp(m, l1, l3)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L10
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v1205 = int64(0)
	v1206 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v1206 < v1205 {
		v1216 = v1205
		goto L233
	} else {
		goto L249
	}
L239:
	;
	if v1177 != int64(-9223372036854775807-1) {
		v1216 = v1177
		goto L233
	} else {
		goto L240
	}
L240:
	;
	if l3 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v1181 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v1181 {
		goto L235
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1184 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v1184
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v1188 - v1189)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v1195-v1189) <= v1184) != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L234
L245:
	;
	v2123 = v1184
	v2129 = int64(0)
	goto L200
L246:
	;
	v1202 = v1195
	goto L248
L247:
	;
	v1202 = v1189 + base.I32_wrap_i64(v1184)
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v1202
	goto L245
L249:
	;
	goto L235
L250:
	;
	v1220 = v687 + int32(112)
	v1224 = int64(0)
	v1230 = m.G0
	v1232 = v1230 - int32(16)
	m.G0 = v1232
	v1234 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v104)))
	v1236 = v1234 & int64(4503599627370495)
	v1240 = int64(base.Ui64(v1234)>>(uint(int64(52))%64)) & int64(2047)
	if v1240 != v1224 {
		goto L255
	} else {
		goto L256
	}
L251:
	;
	goto L252
L252:
	;
	if v841 != 0 {
		goto L264
	} else {
		goto L265
	}
L253:
	;
	v1288 = *(*int64)(unsafe.Add(mBase, uint32(v687)+112))
	v1289 = *(*int64)(unsafe.Add(mBase, uint32(v687)+120))
	v2123 = v1288
	v2129 = v1289
	goto L200
L254:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1220))) = v1276
	*(*int64)(unsafe.Add(mBase, uint32(v1220)+8)) = v1234&int64(-9223372036854775807-1) | v1273<<(uint(int64(48))%64) | v1274
	m.G0 = v1232 + int32(16)
	goto L253
L255:
	;
	if v1240 != int64(2047) {
		goto L258
	} else {
		goto L259
	}
L256:
	;
	goto L257
L257:
	;
	if v1236 == int64(0) {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	v1273 = v1240 + int64(15360)
	v1274 = int64(base.Ui64(v1236) >> (uint(int64(4)) % 64))
	v1276 = v1236 << (uint(int64(60)) % 64)
	goto L254
L259:
	;
	goto L260
L260:
	;
	v1273 = int64(32767)
	v1274 = int64(base.Ui64(v1236) >> (uint(int64(4)) % 64))
	v1276 = v1236 << (uint(int64(60)) % 64)
	goto L254
L261:
	;
	v1258 = int64(0)
	v1273 = v1258
	v1274 = v1224
	v1276 = v1258
	goto L254
L262:
	;
	goto L263
L263:
	;
	v1262 = base.I32_wrap_i64(base.I64_clz(v1236))
	F___ashlti3(m, v1232, v1236, int64(0), v1262+int32(49))
	mBase = m.M
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(v1232)+8))
	v1272 = *(*int64)(unsafe.Add(mBase, uint32(v1232)))
	v1273 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v1262)
	v1274 = v1266 ^ int64(281474976710656)
	v1276 = v1272
	goto L254
L264:
	;
	v1290 = v855
	goto L266
L265:
	;
	v1290 = v852
	goto L266
L266:
	;
	v1295 = v1290<<(uint(int64(2))%64) + v1216 - int64(32)
	if base.I64_extend_i32_u(int32(0)-v38) < v1295 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	v1304 = v687 + int32(160)
	v1308 = m.G0
	v1310 = v1308 - int32(16)
	m.G0 = v1310
	if v104 != 0 {
		goto L271
	} else {
		goto L272
	}
L268:
	;
	goto L269
L269:
	;
	if base.I64_extend_i32_s(v38-int32(226)) <= v1295 {
		goto L277
	} else {
		goto L278
	}
L270:
	;
	v1349 = *(*int64)(unsafe.Add(mBase, uint32(v687)+160))
	v1350 = *(*int64)(unsafe.Add(mBase, uint32(v687)+168))
	v1351 = int64(-1)
	v1352 = int64(9223090561878065151)
	F___multf3(m, v687+int32(144), v1349, v1350, v1351, v1352)
	mBase = m.M
	v1356 = *(*int64)(unsafe.Add(mBase, uint32(v687)+144))
	v1357 = *(*int64)(unsafe.Add(mBase, uint32(v687)+152))
	F___multf3(m, v687+int32(128), v1356, v1357, v1351, v1352)
	mBase = m.M
	v1361 = *(*int64)(unsafe.Add(mBase, uint32(v687)+128))
	v1362 = *(*int64)(unsafe.Add(mBase, uint32(v687)+136))
	v2123 = v1361
	v2129 = v1362
	goto L200
L271:
	;
	v1313 = v104 >> (uint(int32(31)) % 32)
	v1315 = v104 ^ v1313 - v1313
	v1317 = int64(0)
	v1318 = base.I32_clz(v1315)
	F___ashlti3(m, v1310, base.I64_extend_i32_u(v1315), v1317, v1318+int32(81))
	mBase = m.M
	v1322 = *(*int64)(unsafe.Add(mBase, uint32(v1310)+8))
	if v104 < int32(0) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1340 = int64(0)
	v1341 = int64(0)
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1304))) = v1341
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+8)) = v1340
	m.G0 = v1310 + int32(16)
	goto L270
L274:
	;
	v1335 = int64(-9223372036854775807 - 1)
	goto L276
L275:
	;
	v1335 = v1317
	goto L276
L276:
	;
	v1337 = *(*int64)(unsafe.Add(mBase, uint32(v1310)))
	v1340 = v1322 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v1318)<<(uint(int64(48))%64) | v1335
	v1341 = v1337
	goto L273
L277:
	;
	if int32(0) <= v1153 {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	v2041 = v687 + int32(208)
	v2045 = m.G0
	v2047 = v2045 - int32(16)
	m.G0 = v2047
	if v104 != 0 {
		goto L453
	} else {
		goto L454
	}
L280:
	;
	v1378 = v1153
	v1390 = v1295
	v1392 = v854
	v1394 = v856
	goto L283
L281:
	;
	v1480 = v1153
	v1492 = v1295
	v1494 = v854
	v1496 = v856
	goto L282
L282:
	;
	v1503 = v1492 + base.I64_extend_i32_u(int32(32)-v38)
	v1504 = base.I32_wrap_i64(v1503)
	v1505 = int32(0)
	if v1505 < v1504 {
		goto L311
	} else {
		goto L312
	}
L283:
	;
	v1400 = int64(0)
	F___addtf3(m, v687+int32(416), v1392, v1394, v1400, int64(-4611967493404098560))
	mBase = m.M
	v1403 = int64(4611123068473966592)
	v1407 = int32(-1)
	v1411 = v1394 & int64(9223372036854775807)
	v1412 = int64(9223090561878065152)
	if v1411 == v1412 {
		goto L287
	} else {
		goto L288
	}
L284:
	;
	v1480 = v1464
	v1492 = v1466
	v1494 = v1468
	v1496 = v1467
	goto L282
L285:
	;
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(v687)+416))
	v1457 = base.B2i32(int32(0) <= v1452)
	if int32(0) <= v1452 {
		goto L303
	} else {
		goto L304
	}
L286:
	;
	v1452 = v1448
	goto L285
L287:
	;
	v1416 = base.B2i32(v1392 != v1400)
	goto L289
L288:
	;
	v1416 = base.B2i32(base.Ui64(v1412) < base.Ui64(v1411))
	goto L289
L289:
	;
	if v1416 != 0 {
		v1448 = v1407
		goto L286
	} else {
		goto L290
	}
L290:
	;
	goto L291
L291:
	;
	if v1392|(v1411|int64(4611123068473966592)) == int64(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1452 = int32(0)
	goto L285
L293:
	;
	goto L294
L294:
	;
	if int64(0) <= v1394&v1403 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	if base.B2i32(v1394 != v1403)&base.B2i32(v1394 < v1403) != 0 {
		v1448 = v1407
		goto L286
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	if v1394 == v1403 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1452 = base.B2i32(v1392|(v1394^v1403) != int64(0))
	goto L285
L299:
	;
	v1443 = base.B2i32(v1392 != int64(0))
	goto L301
L300:
	;
	v1443 = base.B2i32(v1403 < v1394)
	goto L301
L301:
	;
	if v1443 != 0 {
		v1448 = v1407
		goto L286
	} else {
		goto L302
	}
L302:
	;
	v1448 = base.B2i32(v1392|(v1394^v1403) != int64(0))
	goto L286
L303:
	;
	v1458 = v1455
	goto L305
L304:
	;
	v1458 = v1392
	goto L305
L305:
	;
	v1459 = *(*int64)(unsafe.Add(mBase, uint32(v687)+424))
	if int32(0) <= v1452 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1460 = v1459
	goto L308
L307:
	;
	v1460 = v1394
	goto L308
L308:
	;
	F___addtf3(m, v687+int32(400), v1392, v1394, v1458, v1460)
	mBase = m.M
	v1463 = v1378 << (uint(int32(1)) % 32)
	v1464 = v1457 | v1463
	v1466 = v1390 - int64(1)
	v1467 = *(*int64)(unsafe.Add(mBase, uint32(v687)+408))
	v1468 = *(*int64)(unsafe.Add(mBase, uint32(v687)+400))
	if int32(0) <= v1463 {
		v1378 = v1464
		v1390 = v1466
		v1392 = v1468
		v1394 = v1467
		goto L283
	} else {
		goto L309
	}
L309:
	;
	goto L284
L310:
	;
	v1747 = v687 + int32(320)
	v1748 = int32(1)
	v1752 = int64(0)
	v1761 = v1496 & int64(9223372036854775807)
	v1762 = int64(9223090561878065152)
	if v1761 == v1762 {
		goto L366
	} else {
		goto L367
	}
L311:
	;
	v1508 = v1504
	goto L313
L312:
	;
	v1508 = v1505
	goto L313
L313:
	;
	if v1503 < base.I64_extend_i32_u(v39) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1511 = v1508
	goto L316
L315:
	;
	v1511 = v39
	goto L316
L316:
	;
	if base.Ui32(int32(113)) <= base.Ui32(v1511) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1515 = v687 + int32(384)
	v1519 = m.G0
	v1521 = v1519 - int32(16)
	m.G0 = v1521
	if v104 != 0 {
		goto L321
	} else {
		goto L322
	}
L318:
	;
	goto L319
L319:
	;
	v1562 = v687 + int32(352)
	v1563 = float64(1)
	v1565 = int32(144) - v1511
	if int32(1024) <= v1565 {
		goto L329
	} else {
		goto L330
	}
L320:
	;
	v1558 = *(*int64)(unsafe.Add(mBase, uint32(v687)+392))
	v1559 = *(*int64)(unsafe.Add(mBase, uint32(v687)+384))
	v1742 = v1558
	v1743 = v1559
	v1744 = v22
	v1745 = int64(0)
	goto L310
L321:
	;
	v1524 = v104 >> (uint(int32(31)) % 32)
	v1526 = v104 ^ v1524 - v1524
	v1528 = int64(0)
	v1529 = base.I32_clz(v1526)
	F___ashlti3(m, v1521, base.I64_extend_i32_u(v1526), v1528, v1529+int32(81))
	mBase = m.M
	v1533 = *(*int64)(unsafe.Add(mBase, uint32(v1521)+8))
	if v104 < int32(0) {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	v1551 = int64(0)
	v1552 = int64(0)
	goto L323
L323:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1515))) = v1552
	*(*int64)(unsafe.Add(mBase, uint32(v1515)+8)) = v1551
	m.G0 = v1521 + int32(16)
	goto L320
L324:
	;
	v1546 = int64(-9223372036854775807 - 1)
	goto L326
L325:
	;
	v1546 = v1528
	goto L326
L326:
	;
	v1548 = *(*int64)(unsafe.Add(mBase, uint32(v1521)))
	v1551 = v1533 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v1529)<<(uint(int64(48))%64) | v1546
	v1552 = v1548
	goto L323
L327:
	;
	v1607 = int64(0)
	v1613 = m.G0
	v1615 = v1613 - int32(16)
	m.G0 = v1615
	v1617 = base.I64_reinterpret_f64(base.F64_mul(v1598, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1599+int32(1023))<<(uint(int64(52))%64))))
	v1619 = v1617 & int64(4503599627370495)
	v1623 = int64(base.Ui64(v1617)>>(uint(int64(52))%64)) & int64(2047)
	if v1623 != v1607 {
		goto L347
	} else {
		goto L348
	}
L328:
	;
	goto L327
L329:
	;
	v1569 = base.F64_mul(v1563, float64(8.98846567431158e+307))
	if base.Ui32(v1565) < base.Ui32(int32(2047)) {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	goto L331
L331:
	;
	if int32(-1023) < v1565 {
		v1598 = v1563
		v1599 = v1565
		goto L328
	} else {
		goto L338
	}
L332:
	;
	v1598 = v1569
	v1599 = v1565 - int32(1023)
	goto L328
L333:
	;
	goto L334
L334:
	;
	v1576 = int32(3069)
	if base.Ui32(v1576) <= base.Ui32(v1565) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1579 = v1576
	goto L337
L336:
	;
	v1579 = v1565
	goto L337
L337:
	;
	v1598 = base.F64_mul(v1569, float64(8.98846567431158e+307))
	v1599 = v1579 - int32(2046)
	goto L328
L338:
	;
	v1585 = base.F64_mul(v1563, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1565) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1598 = v1585
	v1599 = v1565 + int32(969)
	goto L328
L340:
	;
	goto L341
L341:
	;
	v1592 = int32(-2960)
	if base.Ui32(v1565) <= base.Ui32(v1592) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1595 = v1592
	goto L344
L343:
	;
	v1595 = v1565
	goto L344
L344:
	;
	v1598 = base.F64_mul(v1585, float64(2.004168360008973e-292))
	v1599 = v1595 + int32(1938)
	goto L328
L345:
	;
	v1672 = v687 + int32(336)
	v1676 = m.G0
	v1678 = v1676 - int32(16)
	m.G0 = v1678
	if v104 != 0 {
		goto L357
	} else {
		goto L358
	}
L346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1562))) = v1659
	*(*int64)(unsafe.Add(mBase, uint32(v1562)+8)) = v1617&int64(-9223372036854775807-1) | v1656<<(uint(int64(48))%64) | v1657
	m.G0 = v1615 + int32(16)
	goto L345
L347:
	;
	if v1623 != int64(2047) {
		goto L350
	} else {
		goto L351
	}
L348:
	;
	goto L349
L349:
	;
	if v1619 == int64(0) {
		goto L353
	} else {
		goto L354
	}
L350:
	;
	v1656 = v1623 + int64(15360)
	v1657 = int64(base.Ui64(v1619) >> (uint(int64(4)) % 64))
	v1659 = v1619 << (uint(int64(60)) % 64)
	goto L346
L351:
	;
	goto L352
L352:
	;
	v1656 = int64(32767)
	v1657 = int64(base.Ui64(v1619) >> (uint(int64(4)) % 64))
	v1659 = v1619 << (uint(int64(60)) % 64)
	goto L346
L353:
	;
	v1641 = int64(0)
	v1656 = v1641
	v1657 = v1607
	v1659 = v1641
	goto L346
L354:
	;
	goto L355
L355:
	;
	v1645 = base.I32_wrap_i64(base.I64_clz(v1619))
	F___ashlti3(m, v1615, v1619, int64(0), v1645+int32(49))
	mBase = m.M
	v1649 = *(*int64)(unsafe.Add(mBase, uint32(v1615)+8))
	v1655 = *(*int64)(unsafe.Add(mBase, uint32(v1615)))
	v1656 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v1645)
	v1657 = v1649 ^ int64(281474976710656)
	v1659 = v1655
	goto L346
L356:
	;
	v1715 = *(*int64)(unsafe.Add(mBase, uint32(v687)+336))
	v1717 = v687 + int32(368)
	v1718 = *(*int64)(unsafe.Add(mBase, uint32(v687)+352))
	v1719 = *(*int64)(unsafe.Add(mBase, uint32(v687)+360))
	v1720 = *(*int64)(unsafe.Add(mBase, uint32(v687)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1717))) = v1718
	v1726 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v1717)+8)) = v1719&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v1719&int64(9223090561878065152))>>(uint(v1726)%64)))|base.I32_wrap_i64(int64(base.Ui64(v1720)>>(uint(v1726)%64)))&int32(_a_F___floatscan_7))<<(uint(v1726)%64)
	goto L363
L357:
	;
	v1681 = v104 >> (uint(int32(31)) % 32)
	v1683 = v104 ^ v1681 - v1681
	v1685 = int64(0)
	v1686 = base.I32_clz(v1683)
	F___ashlti3(m, v1678, base.I64_extend_i32_u(v1683), v1685, v1686+int32(81))
	mBase = m.M
	v1690 = *(*int64)(unsafe.Add(mBase, uint32(v1678)+8))
	if v104 < int32(0) {
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1708 = int64(0)
	v1709 = int64(0)
	goto L359
L359:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1672))) = v1709
	*(*int64)(unsafe.Add(mBase, uint32(v1672)+8)) = v1708
	m.G0 = v1678 + int32(16)
	goto L356
L360:
	;
	v1703 = int64(-9223372036854775807 - 1)
	goto L362
L361:
	;
	v1703 = v1685
	goto L362
L362:
	;
	v1705 = *(*int64)(unsafe.Add(mBase, uint32(v1678)))
	v1708 = v1690 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v1686)<<(uint(int64(48))%64) | v1703
	v1709 = v1705
	goto L359
L363:
	;
	v1740 = *(*int64)(unsafe.Add(mBase, uint32(v687)+376))
	v1741 = *(*int64)(unsafe.Add(mBase, uint32(v687)+368))
	v1742 = v1720
	v1743 = v1715
	v1744 = v1740
	v1745 = v1741
	goto L310
L364:
	;
	v1815 = base.B2i32(v1480&v1748 == int32(0)) & (base.B2i32(v1809 != int32(0)) & base.B2i32(base.Ui32(v1511) < base.Ui32(int32(32))))
	v1816 = v1480 | v1815
	v1819 = m.G0
	v1821 = v1819 - int32(16)
	m.G0 = v1821
	if v1816 != 0 {
		goto L393
	} else {
		goto L394
	}
L365:
	;
	v1809 = v1805
	goto L364
L366:
	;
	v1766 = base.B2i32(v1494 != v1752)
	goto L368
L367:
	;
	v1766 = base.B2i32(base.Ui64(v1762) < base.Ui64(v1761))
	goto L368
L368:
	;
	if v1766 != 0 {
		v1805 = v1748
		goto L365
	} else {
		goto L369
	}
L369:
	;
	goto L371
L371:
	;
	goto L372
L372:
	;
	goto L373
L373:
	;
	if v1494|v1752|(v1761|int64(0)) == int64(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1809 = int32(0)
	goto L364
L375:
	;
	goto L376
L376:
	;
	if int64(0) <= v1496&v1752 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	if v1496 == v1752 {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	goto L379
L379:
	;
	if v1496 == v1752 {
		goto L386
	} else {
		goto L387
	}
L380:
	;
	v1788 = base.B2i32(base.Ui64(v1494) < base.Ui64(v1752))
	goto L382
L381:
	;
	v1788 = base.B2i32(v1496 < v1752)
	goto L382
L382:
	;
	if v1788 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1809 = int32(-1)
	goto L364
L384:
	;
	goto L385
L385:
	;
	v1809 = base.B2i32(v1494^v1752|(v1496^v1752) != int64(0))
	goto L364
L386:
	;
	v1798 = base.B2i32(base.Ui64(v1752) < base.Ui64(v1494))
	goto L388
L387:
	;
	v1798 = base.B2i32(v1752 < v1496)
	goto L388
L388:
	;
	if v1798 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1809 = int32(-1)
	goto L364
L390:
	;
	goto L391
L391:
	;
	v1805 = base.B2i32(v1494^v1752|(v1496^v1752) != int64(0))
	goto L365
L392:
	;
	v1852 = *(*int64)(unsafe.Add(mBase, uint32(v687)+320))
	v1853 = *(*int64)(unsafe.Add(mBase, uint32(v687)+328))
	F___multf3(m, v687+int32(304), v1743, v1742, v1852, v1853)
	mBase = m.M
	v1857 = *(*int64)(unsafe.Add(mBase, uint32(v687)+304))
	v1858 = *(*int64)(unsafe.Add(mBase, uint32(v687)+312))
	F___addtf3(m, v687+int32(272), v1857, v1858, v1745, v1744)
	mBase = m.M
	if v1815 != 0 {
		goto L396
	} else {
		goto L397
	}
L393:
	;
	v1826 = base.I32_clz(v1816)
	F___ashlti3(m, v1821, base.I64_extend_i32_u(v1816), int64(0), int32(112)-(v1826^int32(31)))
	mBase = m.M
	v1831 = *(*int64)(unsafe.Add(mBase, uint32(v1821)+8))
	v1840 = *(*int64)(unsafe.Add(mBase, uint32(v1821)))
	v1843 = v1831 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v1826)<<(uint(int64(48))%64)
	v1844 = v1840
	goto L395
L394:
	;
	v1843 = int64(0)
	v1844 = int64(0)
	goto L395
L395:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1747))) = v1844
	*(*int64)(unsafe.Add(mBase, uint32(v1747)+8)) = v1843
	m.G0 = v1821 + int32(16)
	goto L392
L396:
	;
	v1863 = int64(0)
	goto L398
L397:
	;
	v1863 = v1494
	goto L398
L398:
	;
	if v1815 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v1865 = int64(0)
	goto L401
L400:
	;
	v1865 = v1496
	goto L401
L401:
	;
	F___multf3(m, v687+int32(288), v1743, v1742, v1863, v1865)
	mBase = m.M
	v1869 = *(*int64)(unsafe.Add(mBase, uint32(v687)+288))
	v1870 = *(*int64)(unsafe.Add(mBase, uint32(v687)+296))
	v1871 = *(*int64)(unsafe.Add(mBase, uint32(v687)+272))
	v1872 = *(*int64)(unsafe.Add(mBase, uint32(v687)+280))
	F___addtf3(m, v687+int32(256), v1869, v1870, v1871, v1872)
	mBase = m.M
	v1875 = v687 + int32(240)
	v1876 = *(*int64)(unsafe.Add(mBase, uint32(v687)+256))
	v1877 = *(*int64)(unsafe.Add(mBase, uint32(v687)+264))
	v1879 = m.G0
	v1880 = int32(16)
	v1881 = v1879 - v1880
	m.G0 = v1881
	F___addtf3(m, v1881, v1876, v1877, v1745, v1744^int64(-9223372036854775807-1))
	mBase = m.M
	v1886 = *(*int64)(unsafe.Add(mBase, uint32(v1881)))
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(v1881)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1875)+8)) = v1887
	*(*int64)(unsafe.Add(mBase, uint32(v1875))) = v1886
	m.G0 = v1881 + v1880
	goto L402
L402:
	;
	v1893 = *(*int64)(unsafe.Add(mBase, uint32(v687)+240))
	v1894 = *(*int64)(unsafe.Add(mBase, uint32(v687)+248))
	v1895 = int64(0)
	v1904 = v1894 & int64(9223372036854775807)
	v1905 = int64(9223090561878065152)
	if v1904 == v1905 {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	if v1952 == int32(0) {
		goto L431
	} else {
		goto L432
	}
L404:
	;
	v1952 = v1948
	goto L403
L405:
	;
	v1909 = base.B2i32(v1893 != v1895)
	goto L407
L406:
	;
	v1909 = base.B2i32(base.Ui64(v1905) < base.Ui64(v1904))
	goto L407
L407:
	;
	if v1909 != 0 {
		v1948 = int32(1)
		goto L404
	} else {
		goto L408
	}
L408:
	;
	goto L410
L410:
	;
	goto L411
L411:
	;
	goto L412
L412:
	;
	if v1893|v1895|(v1904|int64(0)) == int64(0) {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1952 = int32(0)
	goto L403
L414:
	;
	goto L415
L415:
	;
	if int64(0) <= v1894&v1895 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	if v1894 == v1895 {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	goto L418
L418:
	;
	if v1894 == v1895 {
		goto L425
	} else {
		goto L426
	}
L419:
	;
	v1931 = base.B2i32(base.Ui64(v1893) < base.Ui64(v1895))
	goto L421
L420:
	;
	v1931 = base.B2i32(v1894 < v1895)
	goto L421
L421:
	;
	if v1931 != 0 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1952 = int32(-1)
	goto L403
L423:
	;
	goto L424
L424:
	;
	v1952 = base.B2i32(v1893^v1895|(v1894^v1895) != int64(0))
	goto L403
L425:
	;
	v1941 = base.B2i32(base.Ui64(v1895) < base.Ui64(v1893))
	goto L427
L426:
	;
	v1941 = base.B2i32(v1895 < v1894)
	goto L427
L427:
	;
	if v1941 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1952 = int32(-1)
	goto L403
L429:
	;
	goto L430
L430:
	;
	v1948 = base.B2i32(v1893^v1895|(v1894^v1895) != int64(0))
	goto L404
L431:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	goto L433
L432:
	;
	goto L433
L433:
	;
	v1959 = v687 + int32(224)
	v1960 = base.I32_wrap_i64(v1492)
	v1962 = m.G0
	v1964 = v1962 - int32(80)
	m.G0 = v1964
	if int32(_a_F___floatscan_8) <= v1960 {
		goto L436
	} else {
		goto L437
	}
L434:
	;
	v2035 = *(*int64)(unsafe.Add(mBase, uint32(v687)+224))
	v2036 = *(*int64)(unsafe.Add(mBase, uint32(v687)+232))
	v2123 = v2035
	v2129 = v2036
	goto L200
L435:
	;
	F___multf3(m, v1964, v2018, v2019, int64(0), base.I64_extend_i32_u(v2020+int32(_a_F___floatscan_9))<<(uint(int64(48))%64))
	mBase = m.M
	v2028 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1959)+8)) = v2028
	v2030 = *(*int64)(unsafe.Add(mBase, uint32(v1964)))
	*(*int64)(unsafe.Add(mBase, uint32(v1959))) = v2030
	m.G0 = v1964 + int32(80)
	goto L434
L436:
	;
	F___multf3(m, v1964+int32(32), v1893, v1894, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1973 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+40))
	v1974 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+32))
	if base.Ui32(v1960) < base.Ui32(int32(_a_F___floatscan_2)) {
		goto L439
	} else {
		goto L440
	}
L437:
	;
	goto L438
L438:
	;
	if int32(-16383) < v1960 {
		v2018 = v1893
		v2019 = v1894
		v2020 = v1960
		goto L435
	} else {
		goto L445
	}
L439:
	;
	v2018 = v1974
	v2019 = v1973
	v2020 = v1960 - int32(_a_F___floatscan_9)
	goto L435
L440:
	;
	goto L441
L441:
	;
	F___multf3(m, v1964+int32(16), v1974, v1973, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1984 = int32(_a_F___floatscan_10)
	if base.Ui32(v1984) <= base.Ui32(v1960) {
		goto L442
	} else {
		goto L443
	}
L442:
	;
	v1987 = v1984
	goto L444
L443:
	;
	v1987 = v1960
	goto L444
L444:
	;
	v1990 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+24))
	v1991 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+16))
	v2018 = v1991
	v2019 = v1990
	v2020 = v1987 - int32(_a_F___floatscan_11)
	goto L435
L445:
	;
	F___multf3(m, v1964-int32(-64), v1893, v1894, int64(0), int64(32088147345014784))
	mBase = m.M
	v1999 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+72))
	v2000 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+64))
	if base.Ui32(int32(-32652)) < base.Ui32(v1960) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2018 = v2000
	v2019 = v1999
	v2020 = v1960 + int32(_a_F___floatscan_12)
	goto L435
L447:
	;
	goto L448
L448:
	;
	F___multf3(m, v1964+int32(48), v2000, v1999, int64(0), int64(32088147345014784))
	mBase = m.M
	v2010 = int32(-48920)
	if base.Ui32(v1960) <= base.Ui32(v2010) {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v2013 = v2010
	goto L451
L450:
	;
	v2013 = v1960
	goto L451
L451:
	;
	v2016 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+56))
	v2017 = *(*int64)(unsafe.Add(mBase, uint32(v1964)+48))
	v2018 = v2017
	v2019 = v2016
	v2020 = v2013 + int32(_a_F___floatscan_13)
	goto L435
L452:
	;
	v2086 = *(*int64)(unsafe.Add(mBase, uint32(v687)+208))
	v2087 = *(*int64)(unsafe.Add(mBase, uint32(v687)+216))
	v2088 = int64(0)
	v2089 = int64(281474976710656)
	F___multf3(m, v687+int32(192), v2086, v2087, v2088, v2089)
	mBase = m.M
	v2093 = *(*int64)(unsafe.Add(mBase, uint32(v687)+192))
	v2094 = *(*int64)(unsafe.Add(mBase, uint32(v687)+200))
	F___multf3(m, v687+int32(176), v2093, v2094, v2088, v2089)
	mBase = m.M
	v2098 = *(*int64)(unsafe.Add(mBase, uint32(v687)+176))
	v2099 = *(*int64)(unsafe.Add(mBase, uint32(v687)+184))
	v2123 = v2098
	v2129 = v2099
	goto L200
L453:
	;
	v2050 = v104 >> (uint(int32(31)) % 32)
	v2052 = v104 ^ v2050 - v2050
	v2054 = int64(0)
	v2055 = base.I32_clz(v2052)
	F___ashlti3(m, v2047, base.I64_extend_i32_u(v2052), v2054, v2055+int32(81))
	mBase = m.M
	v2059 = *(*int64)(unsafe.Add(mBase, uint32(v2047)+8))
	if v104 < int32(0) {
		goto L456
	} else {
		goto L457
	}
L454:
	;
	v2077 = int64(0)
	v2078 = int64(0)
	goto L455
L455:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2041))) = v2078
	*(*int64)(unsafe.Add(mBase, uint32(v2041)+8)) = v2077
	m.G0 = v2047 + int32(16)
	goto L452
L456:
	;
	v2072 = int64(-9223372036854775807 - 1)
	goto L458
L457:
	;
	v2072 = v2054
	goto L458
L458:
	;
	v2074 = *(*int64)(unsafe.Add(mBase, uint32(v2047)))
	v2077 = v2059 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v2055)<<(uint(int64(48))%64) | v2072
	v2078 = v2074
	goto L455
L459:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2140 - int32(1)
	goto L128
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2151)+784)) = int32(0)
	v2318 = base.B2i32(v2288 == int32(46))
	v2320 = v2288 - int32(48)
	if v2318|base.B2i32(base.Ui32(v2320) <= base.Ui32(int32(9))) != 0 {
		goto L491
	} else {
		goto L492
	}
L461:
	;
	if v2209 == int32(48) {
		goto L475
	} else {
		goto L476
	}
L462:
	;
	if v2158 != int32(48) {
		goto L465
	} else {
		goto L466
	}
L463:
	;
	v2207 = F___shgetc(m, l1)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L10
	} else {
		goto L474
	}
L464:
	;
	goto L463
L465:
	;
	if v2158 != int32(46) {
		v2288 = v2158
		v2292 = v2162
		v2298 = v5
		v2307 = v22
		goto L460
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2196 != v2197 {
		goto L470
	} else {
		goto L471
	}
L468:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2189 == v2190 {
		goto L464
	} else {
		goto L469
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2189 + int32(1)
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2189))))
	v2209 = v2195
	goto L461
L470:
	;
	v2199 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2196 + v2199
	v2203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196))))
	v2158 = v2203
	v2162 = v2199
	goto L462
L471:
	;
	goto L472
L472:
	;
	v2205 = F___shgetc(m, l1)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L10
	} else {
		goto L473
	}
L473:
	;
	v2158 = v2205
	v2162 = int32(1)
	goto L462
L474:
	;
	v2209 = v2207
	goto L461
L475:
	;
	v2233 = v22
	goto L478
L476:
	;
	v2258 = v2209
	v2262 = v2162
	v2277 = v22
	goto L477
L477:
	;
	v2288 = v2258
	v2292 = v2262
	v2298 = int32(1)
	v2307 = v2277
	goto L460
L478:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2241 != v2242 {
		goto L481
	} else {
		goto L482
	}
L479:
	;
	v2258 = v2250
	v2262 = int32(1)
	v2277 = v2252
	goto L477
L480:
	;
	v2252 = v2233 - int64(1)
	if v2250 == int32(48) {
		v2233 = v2252
		goto L478
	} else {
		goto L485
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2241 + int32(1)
	v2247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241))))
	v2250 = v2247
	goto L480
L482:
	;
	goto L483
L483:
	;
	v2248 = F___shgetc(m, l1)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L10
	} else {
		goto L484
	}
L484:
	;
	v2250 = v2248
	goto L480
L485:
	;
	goto L479
L486:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2146)+8)) = v5197
	*(*int64)(unsafe.Add(mBase, uint32(v2146))) = v5190
	m.G0 = v2151 + int32(_a_F___floatscan_4)
	v5203 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
	v5204 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
	v5286 = v5204
	v5287 = v5203
	goto L1
L487:
	;
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+784))
	if v2629 == int32(0) {
		goto L540
	} else {
		goto L541
	}
L488:
	;
	v2579 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v2579
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v2583 - v2584)
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(1)|base.B2i32(base.I64_extend_i32_s(v2590-v2584) <= v2579) != 0 {
		goto L537
	} else {
		goto L538
	}
L489:
	;
	if v2517 == int32(0) {
		v2607 = v2523
		v2608 = v2524
		v2611 = v2527
		v2621 = v2537
		v2622 = v2538
		goto L487
	} else {
		goto L535
	}
L490:
	;
	v2509 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2509 < int64(0) {
		v2517 = v2481
		v2523 = v2487
		v2524 = v2488
		v2527 = v2491
		v2537 = v2501
		v2538 = v2502
		goto L489
	} else {
		goto L534
	}
L491:
	;
	v2325 = v2318
	v2326 = v2288
	v2327 = v2320
	v2330 = v2292
	v2331 = v2147
	v2332 = v5
	v2335 = v5
	v2336 = v2298
	v2345 = v2307
	v2346 = v22
	goto L494
L492:
	;
	v2423 = v2288
	v2427 = v2292
	v2428 = v2147
	v2429 = v5
	v2432 = v5
	v2433 = v2298
	v2442 = v2307
	v2443 = v22
	goto L493
L493:
	;
	if v2433 != 0 {
		goto L522
	} else {
		goto L523
	}
L494:
	;
	if v2325&int32(1) != 0 {
		goto L497
	} else {
		goto L498
	}
L495:
	;
	v2423 = v2413
	v2427 = v2397
	v2428 = v2398
	v2429 = v2399
	v2432 = v2400
	v2433 = v2401
	v2442 = v2402
	v2443 = v2403
	goto L493
L496:
	;
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2404 != v2405 {
		goto L517
	} else {
		goto L518
	}
L497:
	;
	if v2336 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L498:
	;
	goto L499
L499:
	;
	v2361 = v2346 + int64(1)
	if v2331 <= int32(2044) {
		goto L503
	} else {
		goto L504
	}
L500:
	;
	v2397 = v2330
	v2398 = v2331
	v2399 = v2332
	v2400 = v2335
	v2401 = int32(1)
	v2402 = v2346
	v2403 = v2346
	goto L496
L501:
	;
	goto L502
L502:
	;
	v2481 = base.B2i32(v2330 == int32(0))
	v2487 = v2331
	v2488 = v2332
	v2491 = v2335
	v2501 = v2345
	v2502 = v2346
	goto L490
L503:
	;
	if v2326 == int32(48) {
		goto L506
	} else {
		goto L507
	}
L504:
	;
	goto L505
L505:
	;
	if v2326 == int32(48) {
		v2397 = v2330
		v2398 = v2331
		v2399 = v2332
		v2400 = v2335
		v2401 = v2336
		v2402 = v2345
		v2403 = v2361
		goto L496
	} else {
		goto L515
	}
L506:
	;
	v2367 = v2335
	goto L508
L507:
	;
	v2367 = base.I32_wrap_i64(v2361)
	goto L508
L508:
	;
	v2372 = v2151 + int32(784) + v2331<<(uint(int32(2))%32)
	if v2332 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2372)))
	v2379 = v2326 + v2373*int32(10) - int32(48)
	goto L511
L510:
	;
	v2379 = v2327
	goto L511
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = v2379
	v2381 = int32(1)
	v2384 = v2332 + v2381
	v2386 = base.B2i32(v2384 == int32(9))
	if v2384 == int32(9) {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v2387 = int32(0)
	goto L514
L513:
	;
	v2387 = v2384
	goto L514
L514:
	;
	v2397 = v2381
	v2398 = v2386 + v2331
	v2399 = v2387
	v2400 = v2367
	v2401 = v2336
	v2402 = v2345
	v2403 = v2361
	goto L496
L515:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+uint32(_c_F___floatscan[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v2151)+uint32(_c_F___floatscan[5]))) = v2391 | int32(1)
	v2397 = v2330
	v2398 = v2331
	v2399 = v2332
	v2400 = int32(_a_F___floatscan_14)
	v2401 = v2336
	v2402 = v2345
	v2403 = v2361
	goto L496
L516:
	;
	v2415 = v2413 - int32(48)
	v2417 = base.B2i32(v2413 == int32(46))
	if v2417|base.B2i32(base.Ui32(v2415) < base.Ui32(int32(10))) != 0 {
		v2325 = v2417
		v2326 = v2413
		v2327 = v2415
		v2330 = v2397
		v2331 = v2398
		v2332 = v2399
		v2335 = v2400
		v2336 = v2401
		v2345 = v2402
		v2346 = v2403
		goto L494
	} else {
		goto L521
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2404 + int32(1)
	v2410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2404))))
	v2413 = v2410
	goto L516
L518:
	;
	goto L519
L519:
	;
	v2411 = F___shgetc(m, l1)
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L10
	} else {
		goto L520
	}
L520:
	;
	v2413 = v2411
	goto L516
L521:
	;
	goto L495
L522:
	;
	v2450 = v2442
	goto L524
L523:
	;
	v2450 = v2443
	goto L524
L524:
	;
	v2451 = int32(0)
	if base.B2i32(v2427 == v2451)|base.B2i32(v2423&int32(-33) != int32(69)) == v2451 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2460 = F_scanexp(m, l1, l3)
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L10
	} else {
		goto L529
	}
L526:
	;
	goto L527
L527:
	;
	v2476 = int32(0)
	v2477 = base.B2i32(v2427 == v2476)
	if v2423 < v2476 {
		v2517 = v2477
		v2523 = v2428
		v2524 = v2429
		v2527 = v2432
		v2537 = v2450
		v2538 = v2443
		goto L489
	} else {
		goto L533
	}
L528:
	;
	v2607 = v2428
	v2608 = v2429
	v2611 = v2432
	v2621 = v2450 + v2474
	v2622 = v2443
	goto L487
L529:
	;
	if v2460 != int64(-9223372036854775807-1) {
		v2474 = v2460
		goto L528
	} else {
		goto L530
	}
L530:
	;
	if l3 == int32(0) {
		goto L488
	} else {
		goto L531
	}
L531:
	;
	v2466 = int64(0)
	v2467 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2467 < v2466 {
		v2474 = v2466
		goto L528
	} else {
		goto L532
	}
L532:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2470 - int32(1)
	v2474 = v2466
	goto L528
L533:
	;
	v2481 = v2477
	v2487 = v2428
	v2488 = v2429
	v2491 = v2432
	v2501 = v2450
	v2502 = v2443
	goto L490
L534:
	;
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2512 - int32(1)
	v2517 = v2481
	v2523 = v2487
	v2524 = v2488
	v2527 = v2491
	v2537 = v2501
	v2538 = v2502
	goto L489
L535:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(28)
	goto L488
L536:
	;
	v5190 = v2579
	v5197 = int64(0)
	goto L486
L537:
	;
	v2597 = v2590
	goto L539
L538:
	;
	v2597 = v2584 + base.I32_wrap_i64(v2579)
	goto L539
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v2597
	goto L536
L540:
	;
	v2635 = int64(0)
	v2641 = m.G0
	v2643 = v2641 - int32(16)
	m.G0 = v2643
	v2645 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v104)))
	v2647 = v2645 & int64(4503599627370495)
	v2651 = int64(base.Ui64(v2645)>>(uint(int64(52))%64)) & int64(2047)
	if v2651 != v2635 {
		goto L545
	} else {
		goto L546
	}
L541:
	;
	goto L542
L542:
	;
	if int32(base.Ui32(v2629)>>(uint(v39)%32)) != 0 {
		goto L554
	} else {
		goto L555
	}
L543:
	;
	v2699 = *(*int64)(unsafe.Add(mBase, uint32(v2151)))
	v2700 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+8))
	v5190 = v2699
	v5197 = v2700
	goto L486
L544:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2151))) = v2687
	*(*int64)(unsafe.Add(mBase, uint32(v2151)+8)) = v2645&int64(-9223372036854775807-1) | v2684<<(uint(int64(48))%64) | v2685
	m.G0 = v2643 + int32(16)
	goto L543
L545:
	;
	if v2651 != int64(2047) {
		goto L548
	} else {
		goto L549
	}
L546:
	;
	goto L547
L547:
	;
	if v2647 == int64(0) {
		goto L551
	} else {
		goto L552
	}
L548:
	;
	v2684 = v2651 + int64(15360)
	v2685 = int64(base.Ui64(v2647) >> (uint(int64(4)) % 64))
	v2687 = v2647 << (uint(int64(60)) % 64)
	goto L544
L549:
	;
	goto L550
L550:
	;
	v2684 = int64(32767)
	v2685 = int64(base.Ui64(v2647) >> (uint(int64(4)) % 64))
	v2687 = v2647 << (uint(int64(60)) % 64)
	goto L544
L551:
	;
	v2669 = int64(0)
	v2684 = v2669
	v2685 = v2635
	v2687 = v2669
	goto L544
L552:
	;
	goto L553
L553:
	;
	v2673 = base.I32_wrap_i64(base.I64_clz(v2647))
	F___ashlti3(m, v2643, v2647, int64(0), v2673+int32(49))
	mBase = m.M
	v2677 = *(*int64)(unsafe.Add(mBase, uint32(v2643)+8))
	v2683 = *(*int64)(unsafe.Add(mBase, uint32(v2643)))
	v2684 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v2673)
	v2685 = v2677 ^ int64(281474976710656)
	v2687 = v2683
	goto L544
L554:
	;
	v2709 = base.B2i32(base.Ui32(v39) <= base.Ui32(int32(30)))
	goto L556
L555:
	;
	v2709 = int32(0)
	goto L556
L556:
	;
	if base.B2i32(v2621 != v2622)|base.B2i32(int64(9) < v2622)|v2709 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v2714 = v2151 + int32(48)
	v2718 = m.G0
	v2720 = v2718 - int32(16)
	m.G0 = v2720
	if v104 != 0 {
		goto L561
	} else {
		goto L562
	}
L558:
	;
	goto L559
L559:
	;
	if base.I64_extend_i32_u(int32(base.Ui32(v2154)>>(uint(int32(1))%32))) < v2621 {
		goto L571
	} else {
		goto L572
	}
L560:
	;
	v2758 = v2151 + int32(32)
	v2761 = m.G0
	v2763 = v2761 - int32(16)
	m.G0 = v2763
	if v2629 != 0 {
		goto L568
	} else {
		goto L569
	}
L561:
	;
	v2723 = v104 >> (uint(int32(31)) % 32)
	v2725 = v104 ^ v2723 - v2723
	v2727 = int64(0)
	v2728 = base.I32_clz(v2725)
	F___ashlti3(m, v2720, base.I64_extend_i32_u(v2725), v2727, v2728+int32(81))
	mBase = m.M
	v2732 = *(*int64)(unsafe.Add(mBase, uint32(v2720)+8))
	if v104 < int32(0) {
		goto L564
	} else {
		goto L565
	}
L562:
	;
	v2750 = int64(0)
	v2751 = int64(0)
	goto L563
L563:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2714))) = v2751
	*(*int64)(unsafe.Add(mBase, uint32(v2714)+8)) = v2750
	m.G0 = v2720 + int32(16)
	goto L560
L564:
	;
	v2745 = int64(-9223372036854775807 - 1)
	goto L566
L565:
	;
	v2745 = v2727
	goto L566
L566:
	;
	v2747 = *(*int64)(unsafe.Add(mBase, uint32(v2720)))
	v2750 = v2732 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v2728)<<(uint(int64(48))%64) | v2745
	v2751 = v2747
	goto L563
L567:
	;
	v2794 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+48))
	v2795 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+56))
	v2796 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+32))
	v2797 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+40))
	F___multf3(m, v2151+int32(16), v2794, v2795, v2796, v2797)
	mBase = m.M
	v2799 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+16))
	v2800 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+24))
	v5190 = v2799
	v5197 = v2800
	goto L486
L568:
	;
	v2768 = base.I32_clz(v2629)
	F___ashlti3(m, v2763, base.I64_extend_i32_u(v2629), int64(0), int32(112)-(v2768^int32(31)))
	mBase = m.M
	v2773 = *(*int64)(unsafe.Add(mBase, uint32(v2763)+8))
	v2782 = *(*int64)(unsafe.Add(mBase, uint32(v2763)))
	v2785 = v2773 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v2768)<<(uint(int64(48))%64)
	v2786 = v2782
	goto L570
L569:
	;
	v2785 = int64(0)
	v2786 = int64(0)
	goto L570
L570:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2758))) = v2786
	*(*int64)(unsafe.Add(mBase, uint32(v2758)+8)) = v2785
	m.G0 = v2763 + int32(16)
	goto L567
L571:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	v2809 = v2151 + int32(96)
	v2813 = m.G0
	v2815 = v2813 - int32(16)
	m.G0 = v2815
	if v104 != 0 {
		goto L575
	} else {
		goto L576
	}
L572:
	;
	goto L573
L573:
	;
	if v2621 < base.I64_extend_i32_s(v38-int32(226)) {
		goto L581
	} else {
		goto L582
	}
L574:
	;
	v2854 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+96))
	v2855 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+104))
	v2856 = int64(-1)
	v2857 = int64(9223090561878065151)
	F___multf3(m, v2151+int32(80), v2854, v2855, v2856, v2857)
	mBase = m.M
	v2861 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+80))
	v2862 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+88))
	F___multf3(m, v2151-int32(-64), v2861, v2862, v2856, v2857)
	mBase = m.M
	v2866 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+64))
	v2867 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+72))
	v5190 = v2866
	v5197 = v2867
	goto L486
L575:
	;
	v2818 = v104 >> (uint(int32(31)) % 32)
	v2820 = v104 ^ v2818 - v2818
	v2822 = int64(0)
	v2823 = base.I32_clz(v2820)
	F___ashlti3(m, v2815, base.I64_extend_i32_u(v2820), v2822, v2823+int32(81))
	mBase = m.M
	v2827 = *(*int64)(unsafe.Add(mBase, uint32(v2815)+8))
	if v104 < int32(0) {
		goto L578
	} else {
		goto L579
	}
L576:
	;
	v2845 = int64(0)
	v2846 = int64(0)
	goto L577
L577:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2809))) = v2846
	*(*int64)(unsafe.Add(mBase, uint32(v2809)+8)) = v2845
	m.G0 = v2815 + int32(16)
	goto L574
L578:
	;
	v2840 = int64(-9223372036854775807 - 1)
	goto L580
L579:
	;
	v2840 = v2822
	goto L580
L580:
	;
	v2842 = *(*int64)(unsafe.Add(mBase, uint32(v2815)))
	v2845 = v2827 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v2823)<<(uint(int64(48))%64) | v2840
	v2846 = v2842
	goto L577
L581:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	v2876 = v2151 + int32(144)
	v2880 = m.G0
	v2882 = v2880 - int32(16)
	m.G0 = v2882
	if v104 != 0 {
		goto L585
	} else {
		goto L586
	}
L582:
	;
	goto L583
L583:
	;
	if v2608 != 0 {
		goto L591
	} else {
		goto L592
	}
L584:
	;
	v2921 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+144))
	v2922 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+152))
	v2923 = int64(0)
	v2924 = int64(281474976710656)
	F___multf3(m, v2151+int32(128), v2921, v2922, v2923, v2924)
	mBase = m.M
	v2928 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+128))
	v2929 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+136))
	F___multf3(m, v2151+int32(112), v2928, v2929, v2923, v2924)
	mBase = m.M
	v2933 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+112))
	v2934 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+120))
	v5190 = v2933
	v5197 = v2934
	goto L486
L585:
	;
	v2885 = v104 >> (uint(int32(31)) % 32)
	v2887 = v104 ^ v2885 - v2885
	v2889 = int64(0)
	v2890 = base.I32_clz(v2887)
	F___ashlti3(m, v2882, base.I64_extend_i32_u(v2887), v2889, v2890+int32(81))
	mBase = m.M
	v2894 = *(*int64)(unsafe.Add(mBase, uint32(v2882)+8))
	if v104 < int32(0) {
		goto L588
	} else {
		goto L589
	}
L586:
	;
	v2912 = int64(0)
	v2913 = int64(0)
	goto L587
L587:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2876))) = v2913
	*(*int64)(unsafe.Add(mBase, uint32(v2876)+8)) = v2912
	m.G0 = v2882 + int32(16)
	goto L584
L588:
	;
	v2907 = int64(-9223372036854775807 - 1)
	goto L590
L589:
	;
	v2907 = v2889
	goto L590
L590:
	;
	v2909 = *(*int64)(unsafe.Add(mBase, uint32(v2882)))
	v2912 = v2894 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v2890)<<(uint(int64(48))%64) | v2907
	v2913 = v2909
	goto L587
L591:
	;
	if v2608 <= int32(8) {
		goto L594
	} else {
		goto L595
	}
L592:
	;
	v3017 = v2607
	goto L593
L593:
	;
	v3044 = base.I32_wrap_i64(v2621)
	if base.B2i32(int32(9) <= v2611)|base.B2i32(int64(17) < v2621)|base.B2i32(v3044 < v2611) != 0 {
		goto L600
	} else {
		goto L601
	}
L594:
	;
	v2941 = v2151 + int32(784) + v2607<<(uint(int32(2))%32)
	v2942 = *(*int32)(unsafe.Add(mBase, uint32(v2941)))
	v2948 = v2942
	v2951 = v2608
	goto L597
L595:
	;
	goto L596
L596:
	;
	v3017 = v2607 + int32(1)
	goto L593
L597:
	;
	v2973 = v2948 * int32(10)
	v2975 = v2951 + int32(1)
	if v2975 != int32(9) {
		v2948 = v2973
		v2951 = v2975
		goto L597
	} else {
		goto L599
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2941))) = v2973
	goto L596
L599:
	;
	goto L598
L600:
	;
	v3449 = v3017
	goto L659
L601:
	;
	if v2621 == int64(9) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v3050 = v2151 + int32(192)
	v3054 = m.G0
	v3056 = v3054 - int32(16)
	m.G0 = v3056
	if v104 != 0 {
		goto L606
	} else {
		goto L607
	}
L603:
	;
	goto L604
L604:
	;
	if v2621 <= int64(8) {
		goto L616
	} else {
		goto L617
	}
L605:
	;
	v3094 = v2151 + int32(176)
	v3095 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+784))
	v3098 = m.G0
	v3100 = v3098 - int32(16)
	m.G0 = v3100
	if v3095 != 0 {
		goto L613
	} else {
		goto L614
	}
L606:
	;
	v3059 = v104 >> (uint(int32(31)) % 32)
	v3061 = v104 ^ v3059 - v3059
	v3063 = int64(0)
	v3064 = base.I32_clz(v3061)
	F___ashlti3(m, v3056, base.I64_extend_i32_u(v3061), v3063, v3064+int32(81))
	mBase = m.M
	v3068 = *(*int64)(unsafe.Add(mBase, uint32(v3056)+8))
	if v104 < int32(0) {
		goto L609
	} else {
		goto L610
	}
L607:
	;
	v3086 = int64(0)
	v3087 = int64(0)
	goto L608
L608:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3050))) = v3087
	*(*int64)(unsafe.Add(mBase, uint32(v3050)+8)) = v3086
	m.G0 = v3056 + int32(16)
	goto L605
L609:
	;
	v3081 = int64(-9223372036854775807 - 1)
	goto L611
L610:
	;
	v3081 = v3063
	goto L611
L611:
	;
	v3083 = *(*int64)(unsafe.Add(mBase, uint32(v3056)))
	v3086 = v3068 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3064)<<(uint(int64(48))%64) | v3081
	v3087 = v3083
	goto L608
L612:
	;
	v3131 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+192))
	v3132 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+200))
	v3133 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+176))
	v3134 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+184))
	F___multf3(m, v2151+int32(160), v3131, v3132, v3133, v3134)
	mBase = m.M
	v3136 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+160))
	v3137 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+168))
	v5190 = v3136
	v5197 = v3137
	goto L486
L613:
	;
	v3105 = base.I32_clz(v3095)
	F___ashlti3(m, v3100, base.I64_extend_i32_u(v3095), int64(0), int32(112)-(v3105^int32(31)))
	mBase = m.M
	v3110 = *(*int64)(unsafe.Add(mBase, uint32(v3100)+8))
	v3119 = *(*int64)(unsafe.Add(mBase, uint32(v3100)))
	v3122 = v3110 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3105)<<(uint(int64(48))%64)
	v3123 = v3119
	goto L615
L614:
	;
	v3122 = int64(0)
	v3123 = int64(0)
	goto L615
L615:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3094))) = v3123
	*(*int64)(unsafe.Add(mBase, uint32(v3094)+8)) = v3122
	m.G0 = v3100 + int32(16)
	goto L612
L616:
	;
	v3141 = v2151 + int32(272)
	v3145 = m.G0
	v3147 = v3145 - int32(16)
	m.G0 = v3147
	if v104 != 0 {
		goto L620
	} else {
		goto L621
	}
L617:
	;
	goto L618
L618:
	;
	v3289 = v39 + v3044*int32(-3) + int32(27)
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+784))
	if int32(base.Ui32(v3293)>>(uint(v3289)%32)) != 0 {
		goto L637
	} else {
		goto L638
	}
L619:
	;
	v3185 = v2151 + int32(256)
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v2151)+784))
	v3189 = m.G0
	v3191 = v3189 - int32(16)
	m.G0 = v3191
	if v3186 != 0 {
		goto L627
	} else {
		goto L628
	}
L620:
	;
	v3150 = v104 >> (uint(int32(31)) % 32)
	v3152 = v104 ^ v3150 - v3150
	v3154 = int64(0)
	v3155 = base.I32_clz(v3152)
	F___ashlti3(m, v3147, base.I64_extend_i32_u(v3152), v3154, v3155+int32(81))
	mBase = m.M
	v3159 = *(*int64)(unsafe.Add(mBase, uint32(v3147)+8))
	if v104 < int32(0) {
		goto L623
	} else {
		goto L624
	}
L621:
	;
	v3177 = int64(0)
	v3178 = int64(0)
	goto L622
L622:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3141))) = v3178
	*(*int64)(unsafe.Add(mBase, uint32(v3141)+8)) = v3177
	m.G0 = v3147 + int32(16)
	goto L619
L623:
	;
	v3172 = int64(-9223372036854775807 - 1)
	goto L625
L624:
	;
	v3172 = v3154
	goto L625
L625:
	;
	v3174 = *(*int64)(unsafe.Add(mBase, uint32(v3147)))
	v3177 = v3159 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3155)<<(uint(int64(48))%64) | v3172
	v3178 = v3174
	goto L622
L626:
	;
	v3222 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+272))
	v3223 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+280))
	v3224 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+256))
	v3225 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+264))
	F___multf3(m, v2151+int32(240), v3222, v3223, v3224, v3225)
	mBase = m.M
	v3228 = v2151 + int32(224)
	v3233 = *(*int32)(unsafe.Add(mBase, uint32((int32(8)-v3044)<<(uint(int32(2))%32))+uint32(_c_F___floatscan[6])))
	v3237 = m.G0
	v3239 = v3237 - int32(16)
	m.G0 = v3239
	if v3233 != 0 {
		goto L631
	} else {
		goto L632
	}
L627:
	;
	v3196 = base.I32_clz(v3186)
	F___ashlti3(m, v3191, base.I64_extend_i32_u(v3186), int64(0), int32(112)-(v3196^int32(31)))
	mBase = m.M
	v3201 = *(*int64)(unsafe.Add(mBase, uint32(v3191)+8))
	v3210 = *(*int64)(unsafe.Add(mBase, uint32(v3191)))
	v3213 = v3201 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3196)<<(uint(int64(48))%64)
	v3214 = v3210
	goto L629
L628:
	;
	v3213 = int64(0)
	v3214 = int64(0)
	goto L629
L629:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3185))) = v3214
	*(*int64)(unsafe.Add(mBase, uint32(v3185)+8)) = v3213
	m.G0 = v3191 + int32(16)
	goto L626
L630:
	;
	v3278 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+240))
	v3279 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+248))
	v3280 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+224))
	v3281 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+232))
	F___divtf3(m, v2151+int32(208), v3278, v3279, v3280, v3281)
	mBase = m.M
	v3283 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+208))
	v3284 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+216))
	v5190 = v3283
	v5197 = v3284
	goto L486
L631:
	;
	v3242 = v3233 >> (uint(int32(31)) % 32)
	v3244 = v3233 ^ v3242 - v3242
	v3246 = int64(0)
	v3247 = base.I32_clz(v3244)
	F___ashlti3(m, v3239, base.I64_extend_i32_u(v3244), v3246, v3247+int32(81))
	mBase = m.M
	v3251 = *(*int64)(unsafe.Add(mBase, uint32(v3239)+8))
	if v3233 < int32(0) {
		goto L634
	} else {
		goto L635
	}
L632:
	;
	v3269 = int64(0)
	v3270 = int64(0)
	goto L633
L633:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3228))) = v3270
	*(*int64)(unsafe.Add(mBase, uint32(v3228)+8)) = v3269
	m.G0 = v3239 + int32(16)
	goto L630
L634:
	;
	v3264 = int64(-9223372036854775807 - 1)
	goto L636
L635:
	;
	v3264 = v3246
	goto L636
L636:
	;
	v3266 = *(*int64)(unsafe.Add(mBase, uint32(v3239)))
	v3269 = v3251 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3247)<<(uint(int64(48))%64) | v3264
	v3270 = v3266
	goto L633
L637:
	;
	v3295 = base.B2i32(v3289 <= int32(30))
	goto L639
L638:
	;
	v3295 = int32(0)
	goto L639
L639:
	;
	if v3295 != 0 {
		goto L600
	} else {
		goto L640
	}
L640:
	;
	v3297 = v2151 + int32(352)
	v3301 = m.G0
	v3303 = v3301 - int32(16)
	m.G0 = v3303
	if v104 != 0 {
		goto L642
	} else {
		goto L643
	}
L641:
	;
	v3341 = v2151 + int32(336)
	v3344 = m.G0
	v3346 = v3344 - int32(16)
	m.G0 = v3346
	if v3293 != 0 {
		goto L649
	} else {
		goto L650
	}
L642:
	;
	v3306 = v104 >> (uint(int32(31)) % 32)
	v3308 = v104 ^ v3306 - v3306
	v3310 = int64(0)
	v3311 = base.I32_clz(v3308)
	F___ashlti3(m, v3303, base.I64_extend_i32_u(v3308), v3310, v3311+int32(81))
	mBase = m.M
	v3315 = *(*int64)(unsafe.Add(mBase, uint32(v3303)+8))
	if v104 < int32(0) {
		goto L645
	} else {
		goto L646
	}
L643:
	;
	v3333 = int64(0)
	v3334 = int64(0)
	goto L644
L644:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3297))) = v3334
	*(*int64)(unsafe.Add(mBase, uint32(v3297)+8)) = v3333
	m.G0 = v3303 + int32(16)
	goto L641
L645:
	;
	v3328 = int64(-9223372036854775807 - 1)
	goto L647
L646:
	;
	v3328 = v3310
	goto L647
L647:
	;
	v3330 = *(*int64)(unsafe.Add(mBase, uint32(v3303)))
	v3333 = v3315 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3311)<<(uint(int64(48))%64) | v3328
	v3334 = v3330
	goto L644
L648:
	;
	v3377 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+352))
	v3378 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+360))
	v3379 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+336))
	v3380 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+344))
	F___multf3(m, v2151+int32(320), v3377, v3378, v3379, v3380)
	mBase = m.M
	v3383 = v2151 + int32(304)
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v3044<<(uint(int32(2))%32))+uint32(_c_F___floatscan[7])))
	v3392 = m.G0
	v3394 = v3392 - int32(16)
	m.G0 = v3394
	if v3388 != 0 {
		goto L653
	} else {
		goto L654
	}
L649:
	;
	v3351 = base.I32_clz(v3293)
	F___ashlti3(m, v3346, base.I64_extend_i32_u(v3293), int64(0), int32(112)-(v3351^int32(31)))
	mBase = m.M
	v3356 = *(*int64)(unsafe.Add(mBase, uint32(v3346)+8))
	v3365 = *(*int64)(unsafe.Add(mBase, uint32(v3346)))
	v3368 = v3356 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3351)<<(uint(int64(48))%64)
	v3369 = v3365
	goto L651
L650:
	;
	v3368 = int64(0)
	v3369 = int64(0)
	goto L651
L651:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3341))) = v3369
	*(*int64)(unsafe.Add(mBase, uint32(v3341)+8)) = v3368
	m.G0 = v3346 + int32(16)
	goto L648
L652:
	;
	v3433 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+320))
	v3434 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+328))
	v3435 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+304))
	v3436 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+312))
	F___multf3(m, v2151+int32(288), v3433, v3434, v3435, v3436)
	mBase = m.M
	v3438 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+288))
	v3439 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+296))
	v5190 = v3438
	v5197 = v3439
	goto L486
L653:
	;
	v3397 = v3388 >> (uint(int32(31)) % 32)
	v3399 = v3388 ^ v3397 - v3397
	v3401 = int64(0)
	v3402 = base.I32_clz(v3399)
	F___ashlti3(m, v3394, base.I64_extend_i32_u(v3399), v3401, v3402+int32(81))
	mBase = m.M
	v3406 = *(*int64)(unsafe.Add(mBase, uint32(v3394)+8))
	if v3388 < int32(0) {
		goto L656
	} else {
		goto L657
	}
L654:
	;
	v3424 = int64(0)
	v3425 = int64(0)
	goto L655
L655:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3383))) = v3425
	*(*int64)(unsafe.Add(mBase, uint32(v3383)+8)) = v3424
	m.G0 = v3394 + int32(16)
	goto L652
L656:
	;
	v3419 = int64(-9223372036854775807 - 1)
	goto L658
L657:
	;
	v3419 = v3401
	goto L658
L658:
	;
	v3421 = *(*int64)(unsafe.Add(mBase, uint32(v3394)))
	v3424 = v3406 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v3402)<<(uint(int64(48))%64) | v3419
	v3425 = v3421
	goto L655
L659:
	;
	v3477 = v2151 + int32(784) + v3449<<(uint(int32(2))%32)
	v3480 = *(*int32)(unsafe.Add(mBase, uint32(v3477-int32(4))))
	if v3480 == int32(0) {
		v3449 = v3449 - int32(1)
		goto L659
	} else {
		goto L661
	}
L660:
	;
	v3483 = int32(0)
	v3485 = base.I32_rem_s(v3044, int32(9))
	if v3485 == v3483 {
		goto L663
	} else {
		goto L664
	}
L661:
	;
	goto L660
L662:
	;
	v3633 = v3604
	v3635 = v3606
	v3640 = v3611
	v3643 = v3483
	goto L683
L663:
	;
	v3604 = v3449
	v3606 = int32(0)
	v3611 = v3044
	goto L662
L664:
	;
	goto L665
L665:
	;
	if v2621 < int64(0) {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v3493 = v3485 + int32(9)
	goto L668
L667:
	;
	v3493 = v3485
	goto L668
L668:
	;
	if v3449 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L669:
	;
	v3604 = v3572
	v3606 = v3574
	v3611 = v3579 - v3493 + int32(9)
	goto L662
L670:
	;
	v3496 = int32(0)
	v3572 = v3496
	v3574 = v3496
	v3579 = v3044
	goto L669
L671:
	;
	goto L672
L672:
	;
	v3499 = int32(0)
	v3505 = *(*int32)(unsafe.Add(mBase, uint32((v3499-v3493)<<(uint(int32(2))%32))+uint32(_c_F___floatscan[1])))
	v3506 = base.I32_div_s(int32(1000000000), v3505)
	v3513 = v3499
	v3515 = v3499
	v3516 = v3499
	v3518 = v3044
	goto L673
L673:
	;
	v3543 = v2151 + int32(784) + v3515<<(uint(int32(2))%32)
	v3544 = *(*int32)(unsafe.Add(mBase, uint32(v3543)))
	v3545 = base.I32_div_u_s(v3544, v3505)
	v3546 = v3516 + v3545
	*(*int32)(unsafe.Add(mBase, uint32(v3543))) = v3546
	v3555 = base.B2i32(v3546 == int32(0)) & base.B2i32(v3513 == v3515)
	if v3555 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	if v3562 == int32(0) {
		v3572 = v3449
		v3574 = v3556
		v3579 = v3559
		goto L669
	} else {
		goto L682
	}
L675:
	;
	v3556 = (v3513 + int32(1)) & int32(2047)
	goto L677
L676:
	;
	v3556 = v3513
	goto L677
L677:
	;
	if v3555 != 0 {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	v3559 = v3518 - int32(9)
	goto L680
L679:
	;
	v3559 = v3518
	goto L680
L680:
	;
	v3562 = v3506 * (v3544 - v3505*v3545)
	v3564 = v3515 + int32(1)
	if v3564 != v3449 {
		v3513 = v3556
		v3515 = v3564
		v3516 = v3562
		v3518 = v3559
		goto L673
	} else {
		goto L681
	}
L681:
	;
	goto L674
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3477))) = v3562
	v3572 = v3449 + int32(1)
	v3574 = v3556
	v3579 = v3559
	goto L669
L683:
	;
	v3669 = v3633
	v3679 = v3643
	goto L686
L684:
	;
	v3811 = v3669
	v3813 = v3635
	v3818 = v3640
	v3821 = v3679
	goto L713
L685:
	;
	goto L684
L686:
	;
	if base.B2i32(v3640 < int32(36)) == int32(0) {
		goto L688
	} else {
		goto L689
	}
L687:
	;
	v3782 = (v3635 - int32(1)) & int32(2047)
	if v3782 == v3708 {
		goto L709
	} else {
		goto L710
	}
L688:
	;
	if v3640 != int32(36) {
		goto L685
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	v3708 = v3669
	v3713 = int32(0)
	v3714 = v3669 + int32(2047)
	goto L693
L691:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v2151+int32(784)+v3635<<(uint(int32(2))%32))))
	if base.Ui32(int32(_a_F___floatscan_15)) <= base.Ui32(v3701) {
		goto L685
	} else {
		goto L692
	}
L692:
	;
	goto L690
L693:
	;
	v3740 = v3714 & int32(2047)
	v3743 = v2151 + int32(784) + v3740<<(uint(int32(2))%32)
	v3744 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3743))))
	v3747 = base.I64_extend_i32_u(v3713) + v3744<<(uint(int64(29))%64)
	if base.Ui64(v3747) < base.Ui64(int64(1000000001)) {
		goto L695
	} else {
		goto L696
	}
L694:
	;
	v3776 = v3679 - int32(29)
	if v3759 == int32(0) {
		v3669 = v3708
		v3679 = v3776
		goto L686
	} else {
		goto L708
	}
L695:
	;
	v3757 = v3747
	v3759 = int32(0)
	goto L697
L696:
	;
	v3751 = int64(1000000000)
	v3752 = base.I64_div_u_s(v3747, v3751)
	v3757 = v3747 - v3752*v3751
	v3759 = base.I32_wrap_i64(v3752)
	goto L697
L697:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3743))) = uint32(v3757)
	if v3757 == int64(0) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v3763 = v3740
	goto L700
L699:
	;
	v3763 = v3708
	goto L700
L700:
	;
	if v3635 == v3740 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	v3765 = v3708
	goto L703
L702:
	;
	v3765 = v3763
	goto L703
L703:
	;
	v3769 = (v3708 - int32(1)) & int32(2047)
	if v3740 != v3769 {
		goto L704
	} else {
		goto L705
	}
L704:
	;
	v3771 = v3708
	goto L706
L705:
	;
	v3771 = v3765
	goto L706
L706:
	;
	if v3635 != v3740 {
		v3708 = v3771
		v3713 = v3759
		v3714 = v3740 - int32(1)
		goto L693
	} else {
		goto L707
	}
L707:
	;
	goto L694
L708:
	;
	goto L687
L709:
	;
	v3785 = v2151 + int32(784)
	v3790 = int32(2)
	v3792 = v3785 + (v3708+int32(2046))&int32(2047)<<(uint(v3790)%32)
	v3793 = *(*int32)(unsafe.Add(mBase, uint32(v3792)))
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(v3769<<(uint(v3790)%32)+v3785)))
	*(*int32)(unsafe.Add(mBase, uint32(v3792))) = v3793 | v3797
	v3800 = v3769
	goto L711
L710:
	;
	v3800 = v3708
	goto L711
L711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2151+int32(784)+v3782<<(uint(int32(2))%32)))) = v3759
	v3633 = v3800
	v3635 = v3782
	v3640 = v3640 + int32(9)
	v3643 = v3776
	goto L683
L712:
	;
	v4502 = (v3889 + int32(4)) & int32(2047)
	if v4502 == v4014 {
		v4908 = v4493
		v4909 = v4495
		goto L830
	} else {
		goto L831
	}
L713:
	;
	v3839 = int32(1)
	v3841 = int32(2047)
	v3842 = (v3811 + v3839) & v3841
	v3851 = v2151 + int32(784) + (v3811-v3839)&v3841<<(uint(int32(2))%32)
	v3855 = v3813
	v3860 = v3818
	v3863 = v3821
	goto L715
L714:
	;
	v4218 = v2151 + int32(656)
	v4219 = float64(1)
	v4221 = int32(225) - v4134
	if int32(1024) <= v4221 {
		goto L772
	} else {
		goto L773
	}
L715:
	;
	if int32(45) < v3860 {
		goto L717
	} else {
		goto L718
	}
L716:
	;
	goto L714
L717:
	;
	v3885 = int32(9)
	goto L719
L718:
	;
	v3885 = int32(1)
	goto L719
L719:
	;
	v3889 = v3855
	v3897 = v3863
	goto L721
L720:
	;
	goto L716
L721:
	;
	v3921 = int32(0)
	goto L724
L722:
	;
	v4144 = int32(-1)
	v4152 = v3889
	v4154 = int32(0)
	v4155 = v3889
	v4157 = v3860
	goto L757
L723:
	;
	v4140 = v3885 + v3897
	if v3811 == v3889 {
		v3889 = v3811
		v3897 = v4140
		goto L721
	} else {
		goto L756
	}
L724:
	;
	v3947 = (v3921 + v3889) & int32(2047)
	if v3947 == v3811 {
		goto L726
	} else {
		goto L727
	}
L725:
	;
	if v3860 != int32(36) {
		goto L723
	} else {
		goto L731
	}
L726:
	;
	goto L725
L727:
	;
	v3951 = int32(2)
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(v2151+int32(784)+v3947<<(uint(v3951)%32))))
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3921<<(uint(v3951)%32))+uint32(_c_F___floatscan[8])))
	if base.Ui32(v3954) < base.Ui32(v3957) {
		goto L726
	} else {
		goto L728
	}
L728:
	;
	if base.Ui32(v3957) < base.Ui32(v3954) {
		goto L723
	} else {
		goto L729
	}
L729:
	;
	v3961 = v3921 + int32(1)
	if v3961 != int32(4) {
		v3921 = v3961
		goto L724
	} else {
		goto L730
	}
L730:
	;
	goto L726
L731:
	;
	v3969 = int64(0)
	v3973 = v3811
	v3977 = int32(0)
	v3993 = v3969
	v3994 = v3969
	goto L732
L732:
	;
	v4003 = (v3977 + v3889) & int32(2047)
	if v3973 == v4003 {
		goto L734
	} else {
		goto L735
	}
L733:
	;
	v4075 = v2151 + int32(720)
	v4079 = m.G0
	v4081 = v4079 - int32(16)
	m.G0 = v4081
	if v104 != 0 {
		goto L743
	} else {
		goto L744
	}
L734:
	;
	v4008 = (v3973 + int32(1)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v4008<<(uint(int32(2))%32)+v2151)+780)) = int32(0)
	v4014 = v4008
	goto L736
L735:
	;
	v4014 = v3973
	goto L736
L736:
	;
	v4016 = v2151 + int32(768)
	v4022 = *(*int32)(unsafe.Add(mBase, uint32(v2151+int32(784)+v4003<<(uint(int32(2))%32))))
	v4025 = m.G0
	v4027 = v4025 - int32(16)
	m.G0 = v4027
	if v4022 != 0 {
		goto L738
	} else {
		goto L739
	}
L737:
	;
	F___multf3(m, v2151+int32(752), v3994, v3993, int64(0), int64(4619810130798575616))
	mBase = m.M
	v4063 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+752))
	v4064 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+760))
	v4065 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+768))
	v4066 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+776))
	F___addtf3(m, v2151+int32(736), v4063, v4064, v4065, v4066)
	mBase = m.M
	v4068 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+744))
	v4069 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+736))
	v4071 = v3977 + int32(1)
	if v4071 != int32(4) {
		v3973 = v4014
		v3977 = v4071
		v3993 = v4068
		v3994 = v4069
		goto L732
	} else {
		goto L741
	}
L738:
	;
	v4032 = base.I32_clz(v4022)
	F___ashlti3(m, v4027, base.I64_extend_i32_u(v4022), int64(0), int32(112)-(v4032^int32(31)))
	mBase = m.M
	v4037 = *(*int64)(unsafe.Add(mBase, uint32(v4027)+8))
	v4046 = *(*int64)(unsafe.Add(mBase, uint32(v4027)))
	v4049 = v4037 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v4032)<<(uint(int64(48))%64)
	v4050 = v4046
	goto L740
L739:
	;
	v4049 = int64(0)
	v4050 = int64(0)
	goto L740
L740:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4016))) = v4050
	*(*int64)(unsafe.Add(mBase, uint32(v4016)+8)) = v4049
	m.G0 = v4027 + int32(16)
	goto L737
L741:
	;
	goto L733
L742:
	;
	v4120 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+720))
	v4121 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+728))
	F___multf3(m, v2151+int32(704), v4069, v4068, v4120, v4121)
	mBase = m.M
	v4124 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+712))
	v4125 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+704))
	v4127 = v3897 + int32(113)
	v4128 = v4127 - v38
	v4129 = int32(0)
	if v4129 < v4128 {
		goto L749
	} else {
		goto L750
	}
L743:
	;
	v4084 = v104 >> (uint(int32(31)) % 32)
	v4086 = v104 ^ v4084 - v4084
	v4088 = int64(0)
	v4089 = base.I32_clz(v4086)
	F___ashlti3(m, v4081, base.I64_extend_i32_u(v4086), v4088, v4089+int32(81))
	mBase = m.M
	v4093 = *(*int64)(unsafe.Add(mBase, uint32(v4081)+8))
	if v104 < int32(0) {
		goto L746
	} else {
		goto L747
	}
L744:
	;
	v4111 = int64(0)
	v4112 = int64(0)
	goto L745
L745:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4075))) = v4112
	*(*int64)(unsafe.Add(mBase, uint32(v4075)+8)) = v4111
	m.G0 = v4081 + int32(16)
	goto L742
L746:
	;
	v4106 = int64(-9223372036854775807 - 1)
	goto L748
L747:
	;
	v4106 = v4088
	goto L748
L748:
	;
	v4108 = *(*int64)(unsafe.Add(mBase, uint32(v4081)))
	v4111 = v4093 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatscan_5)-v4089)<<(uint(int64(48))%64) | v4106
	v4112 = v4108
	goto L745
L749:
	;
	v4132 = v4128
	goto L751
L750:
	;
	v4132 = v4129
	goto L751
L751:
	;
	v4133 = base.B2i32(v4128 < v39)
	if v4128 < v39 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v4134 = v4132
	goto L754
L753:
	;
	v4134 = v39
	goto L754
L754:
	;
	if base.Ui32(v4134) <= base.Ui32(int32(112)) {
		goto L720
	} else {
		goto L755
	}
L755:
	;
	v4493 = int64(0)
	v4494 = v4124
	v4495 = v22
	v4496 = v4125
	v4497 = v22
	v4498 = v22
	goto L712
L756:
	;
	goto L722
L757:
	;
	v4179 = v2151 + int32(784)
	v4182 = v4179 + v4155<<(uint(int32(2))%32)
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4182)))
	v4185 = v4154 + int32(base.Ui32(v4183)>>(uint(v3885)%32))
	*(*int32)(unsafe.Add(mBase, uint32(v4182))) = v4185
	v4194 = base.B2i32(v4185 == int32(0)) & base.B2i32(v4152 == v4155)
	if v4194 != 0 {
		goto L759
	} else {
		goto L760
	}
L758:
	;
	if v4200 == int32(0) {
		v3855 = v4195
		v3860 = v4198
		v3863 = v4140
		goto L715
	} else {
		goto L766
	}
L759:
	;
	v4195 = (v4152 + int32(1)) & int32(2047)
	goto L761
L760:
	;
	v4195 = v4152
	goto L761
L761:
	;
	if v4194 != 0 {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v4198 = v4157 - int32(9)
	goto L764
L763:
	;
	v4198 = v4157
	goto L764
L764:
	;
	v4200 = v4183 & (v4144<<(uint(v3885)%32) ^ v4144) * int32(base.Ui32(int32(1000000000))>>(uint(v3885)%32))
	v4204 = (v4155 + int32(1)) & int32(2047)
	if v4204 != v3811 {
		v4152 = v4195
		v4154 = v4200
		v4155 = v4204
		v4157 = v4198
		goto L757
	} else {
		goto L765
	}
L765:
	;
	goto L758
L766:
	;
	if v3842 != v4195 {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3811<<(uint(int32(2))%32)+v4179))) = v4200
	v3811 = v3842
	v3813 = v4195
	v3818 = v4198
	v3821 = v4140
	goto L713
L768:
	;
	goto L769
L769:
	;
	v4213 = *(*int32)(unsafe.Add(mBase, uint32(v3851)))
	*(*int32)(unsafe.Add(mBase, uint32(v3851))) = v4213 | int32(1)
	v3855 = v4195
	v3860 = v4198
	v3863 = v4140
	goto L715
L770:
	;
	v4263 = int64(0)
	v4269 = m.G0
	v4271 = v4269 - int32(16)
	m.G0 = v4271
	v4273 = base.I64_reinterpret_f64(base.F64_mul(v4254, base.F64_reinterpret_i64(base.I64_extend_i32_u(v4255+int32(1023))<<(uint(int64(52))%64))))
	v4275 = v4273 & int64(4503599627370495)
	v4279 = int64(base.Ui64(v4273)>>(uint(int64(52))%64)) & int64(2047)
	if v4279 != v4263 {
		goto L790
	} else {
		goto L791
	}
L771:
	;
	goto L770
L772:
	;
	v4225 = base.F64_mul(v4219, float64(8.98846567431158e+307))
	if base.Ui32(v4221) < base.Ui32(int32(2047)) {
		goto L775
	} else {
		goto L776
	}
L773:
	;
	goto L774
L774:
	;
	if int32(-1023) < v4221 {
		v4254 = v4219
		v4255 = v4221
		goto L771
	} else {
		goto L781
	}
L775:
	;
	v4254 = v4225
	v4255 = v4221 - int32(1023)
	goto L771
L776:
	;
	goto L777
L777:
	;
	v4232 = int32(3069)
	if base.Ui32(v4232) <= base.Ui32(v4221) {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v4235 = v4232
	goto L780
L779:
	;
	v4235 = v4221
	goto L780
L780:
	;
	v4254 = base.F64_mul(v4225, float64(8.98846567431158e+307))
	v4255 = v4235 - int32(2046)
	goto L771
L781:
	;
	v4241 = base.F64_mul(v4219, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v4221) {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v4254 = v4241
	v4255 = v4221 + int32(969)
	goto L771
L783:
	;
	goto L784
L784:
	;
	v4248 = int32(-2960)
	if base.Ui32(v4221) <= base.Ui32(v4248) {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v4251 = v4248
	goto L787
L786:
	;
	v4251 = v4221
	goto L787
L787:
	;
	v4254 = base.F64_mul(v4241, float64(2.004168360008973e-292))
	v4255 = v4251 + int32(1938)
	goto L771
L788:
	;
	v4328 = v2151 + int32(688)
	v4329 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+656))
	v4330 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+664))
	*(*int64)(unsafe.Add(mBase, uint32(v4328))) = v4329
	v4336 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v4328)+8)) = v4330&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v4330&int64(9223090561878065152))>>(uint(v4336)%64)))|base.I32_wrap_i64(int64(base.Ui64(v4124)>>(uint(v4336)%64)))&int32(_a_F___floatscan_7))<<(uint(v4336)%64)
	goto L799
L789:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4218))) = v4315
	*(*int64)(unsafe.Add(mBase, uint32(v4218)+8)) = v4273&int64(-9223372036854775807-1) | v4312<<(uint(int64(48))%64) | v4313
	m.G0 = v4271 + int32(16)
	goto L788
L790:
	;
	if v4279 != int64(2047) {
		goto L793
	} else {
		goto L794
	}
L791:
	;
	goto L792
L792:
	;
	if v4275 == int64(0) {
		goto L796
	} else {
		goto L797
	}
L793:
	;
	v4312 = v4279 + int64(15360)
	v4313 = int64(base.Ui64(v4275) >> (uint(int64(4)) % 64))
	v4315 = v4275 << (uint(int64(60)) % 64)
	goto L789
L794:
	;
	goto L795
L795:
	;
	v4312 = int64(32767)
	v4313 = int64(base.Ui64(v4275) >> (uint(int64(4)) % 64))
	v4315 = v4275 << (uint(int64(60)) % 64)
	goto L789
L796:
	;
	v4297 = int64(0)
	v4312 = v4297
	v4313 = v4263
	v4315 = v4297
	goto L789
L797:
	;
	goto L798
L798:
	;
	v4301 = base.I32_wrap_i64(base.I64_clz(v4275))
	F___ashlti3(m, v4271, v4275, int64(0), v4301+int32(49))
	mBase = m.M
	v4305 = *(*int64)(unsafe.Add(mBase, uint32(v4271)+8))
	v4311 = *(*int64)(unsafe.Add(mBase, uint32(v4271)))
	v4312 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4301)
	v4313 = v4305 ^ int64(281474976710656)
	v4315 = v4311
	goto L789
L799:
	;
	v4350 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+696))
	v4351 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+688))
	v4353 = v2151 + int32(640)
	v4354 = float64(1)
	v4356 = int32(113) - v4134
	if int32(1024) <= v4356 {
		goto L802
	} else {
		goto L803
	}
L800:
	;
	v4398 = int64(0)
	v4404 = m.G0
	v4406 = v4404 - int32(16)
	m.G0 = v4406
	v4408 = base.I64_reinterpret_f64(base.F64_mul(v4389, base.F64_reinterpret_i64(base.I64_extend_i32_u(v4390+int32(1023))<<(uint(int64(52))%64))))
	v4410 = v4408 & int64(4503599627370495)
	v4414 = int64(base.Ui64(v4408)>>(uint(int64(52))%64)) & int64(2047)
	if v4414 != v4398 {
		goto L820
	} else {
		goto L821
	}
L801:
	;
	goto L800
L802:
	;
	v4360 = base.F64_mul(v4354, float64(8.98846567431158e+307))
	if base.Ui32(v4356) < base.Ui32(int32(2047)) {
		goto L805
	} else {
		goto L806
	}
L803:
	;
	goto L804
L804:
	;
	if int32(-1023) < v4356 {
		v4389 = v4354
		v4390 = v4356
		goto L801
	} else {
		goto L811
	}
L805:
	;
	v4389 = v4360
	v4390 = v4356 - int32(1023)
	goto L801
L806:
	;
	goto L807
L807:
	;
	v4367 = int32(3069)
	if base.Ui32(v4367) <= base.Ui32(v4356) {
		goto L808
	} else {
		goto L809
	}
L808:
	;
	v4370 = v4367
	goto L810
L809:
	;
	v4370 = v4356
	goto L810
L810:
	;
	v4389 = base.F64_mul(v4360, float64(8.98846567431158e+307))
	v4390 = v4370 - int32(2046)
	goto L801
L811:
	;
	v4376 = base.F64_mul(v4354, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v4356) {
		goto L812
	} else {
		goto L813
	}
L812:
	;
	v4389 = v4376
	v4390 = v4356 + int32(969)
	goto L801
L813:
	;
	goto L814
L814:
	;
	v4383 = int32(-2960)
	if base.Ui32(v4356) <= base.Ui32(v4383) {
		goto L815
	} else {
		goto L816
	}
L815:
	;
	v4386 = v4383
	goto L817
L816:
	;
	v4386 = v4356
	goto L817
L817:
	;
	v4389 = base.F64_mul(v4376, float64(2.004168360008973e-292))
	v4390 = v4386 + int32(1938)
	goto L801
L818:
	;
	v4464 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+640))
	v4465 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+648))
	F_fmodl(m, v2151+int32(672), v4125, v4124, v4464, v4465)
	mBase = m.M
	v4468 = v2151 + int32(624)
	v4469 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+672))
	v4470 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+680))
	v4472 = m.G0
	v4473 = int32(16)
	v4474 = v4472 - v4473
	m.G0 = v4474
	F___addtf3(m, v4474, v4125, v4124, v4469, v4470^int64(-9223372036854775807-1))
	mBase = m.M
	v4479 = *(*int64)(unsafe.Add(mBase, uint32(v4474)))
	v4480 = *(*int64)(unsafe.Add(mBase, uint32(v4474)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4468)+8)) = v4480
	*(*int64)(unsafe.Add(mBase, uint32(v4468))) = v4479
	m.G0 = v4474 + v4473
	goto L829
L819:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4353))) = v4450
	*(*int64)(unsafe.Add(mBase, uint32(v4353)+8)) = v4408&int64(-9223372036854775807-1) | v4447<<(uint(int64(48))%64) | v4448
	m.G0 = v4406 + int32(16)
	goto L818
L820:
	;
	if v4414 != int64(2047) {
		goto L823
	} else {
		goto L824
	}
L821:
	;
	goto L822
L822:
	;
	if v4410 == int64(0) {
		goto L826
	} else {
		goto L827
	}
L823:
	;
	v4447 = v4414 + int64(15360)
	v4448 = int64(base.Ui64(v4410) >> (uint(int64(4)) % 64))
	v4450 = v4410 << (uint(int64(60)) % 64)
	goto L819
L824:
	;
	goto L825
L825:
	;
	v4447 = int64(32767)
	v4448 = int64(base.Ui64(v4410) >> (uint(int64(4)) % 64))
	v4450 = v4410 << (uint(int64(60)) % 64)
	goto L819
L826:
	;
	v4432 = int64(0)
	v4447 = v4432
	v4448 = v4398
	v4450 = v4432
	goto L819
L827:
	;
	goto L828
L828:
	;
	v4436 = base.I32_wrap_i64(base.I64_clz(v4410))
	F___ashlti3(m, v4406, v4410, int64(0), v4436+int32(49))
	mBase = m.M
	v4440 = *(*int64)(unsafe.Add(mBase, uint32(v4406)+8))
	v4446 = *(*int64)(unsafe.Add(mBase, uint32(v4406)))
	v4447 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4436)
	v4448 = v4440 ^ int64(281474976710656)
	v4450 = v4446
	goto L819
L829:
	;
	v4488 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+624))
	v4489 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+632))
	F___addtf3(m, v2151+int32(608), v4351, v4350, v4488, v4489)
	mBase = m.M
	v4491 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+616))
	v4492 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+608))
	v4493 = v4469
	v4494 = v4491
	v4495 = v4470
	v4496 = v4492
	v4497 = v4351
	v4498 = v4350
	goto L712
L830:
	;
	F___addtf3(m, v2151+int32(432), v4496, v4494, v4908, v4909)
	mBase = m.M
	v4915 = v2151 + int32(416)
	v4916 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+432))
	v4917 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+440))
	v4919 = m.G0
	v4920 = int32(16)
	v4921 = v4919 - v4920
	m.G0 = v4921
	F___addtf3(m, v4921, v4916, v4917, v4497, v4498^int64(-9223372036854775807-1))
	mBase = m.M
	v4926 = *(*int64)(unsafe.Add(mBase, uint32(v4921)))
	v4927 = *(*int64)(unsafe.Add(mBase, uint32(v4921)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4915)+8)) = v4927
	*(*int64)(unsafe.Add(mBase, uint32(v4915))) = v4926
	m.G0 = v4921 + v4920
	goto L917
L831:
	;
	v4509 = *(*int32)(unsafe.Add(mBase, uint32(v2151+int32(784)+v4502<<(uint(int32(2))%32))))
	if base.Ui32(v4509) <= base.Ui32(int32(499999999)) {
		goto L833
	} else {
		goto L834
	}
L832:
	;
	if base.Ui32(int32(111)) < base.Ui32(v4134) {
		v4908 = v4830
		v4909 = v4831
		goto L830
	} else {
		goto L887
	}
L833:
	;
	if base.B2i32(v4509 == int32(0))&base.B2i32((v3889+int32(5))&int32(2047) == v4014) != 0 {
		v4830 = v4493
		v4831 = v4495
		goto L832
	} else {
		goto L836
	}
L834:
	;
	goto L835
L835:
	;
	if v4509 != int32(500000000) {
		goto L848
	} else {
		goto L849
	}
L836:
	;
	v4521 = v2151 + int32(496)
	v4525 = int64(0)
	v4531 = m.G0
	v4533 = v4531 - int32(16)
	m.G0 = v4533
	v4535 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(v104), float64(0.25)))
	v4537 = v4535 & int64(4503599627370495)
	v4541 = int64(base.Ui64(v4535)>>(uint(int64(52))%64)) & int64(2047)
	if v4541 != v4525 {
		goto L839
	} else {
		goto L840
	}
L837:
	;
	v4591 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+496))
	v4592 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+504))
	F___addtf3(m, v2151+int32(480), v4493, v4495, v4591, v4592)
	mBase = m.M
	v4594 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+488))
	v4595 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+480))
	v4830 = v4595
	v4831 = v4594
	goto L832
L838:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4521))) = v4577
	*(*int64)(unsafe.Add(mBase, uint32(v4521)+8)) = v4535&int64(-9223372036854775807-1) | v4574<<(uint(int64(48))%64) | v4575
	m.G0 = v4533 + int32(16)
	goto L837
L839:
	;
	if v4541 != int64(2047) {
		goto L842
	} else {
		goto L843
	}
L840:
	;
	goto L841
L841:
	;
	if v4537 == int64(0) {
		goto L845
	} else {
		goto L846
	}
L842:
	;
	v4574 = v4541 + int64(15360)
	v4575 = int64(base.Ui64(v4537) >> (uint(int64(4)) % 64))
	v4577 = v4537 << (uint(int64(60)) % 64)
	goto L838
L843:
	;
	goto L844
L844:
	;
	v4574 = int64(32767)
	v4575 = int64(base.Ui64(v4537) >> (uint(int64(4)) % 64))
	v4577 = v4537 << (uint(int64(60)) % 64)
	goto L838
L845:
	;
	v4559 = int64(0)
	v4574 = v4559
	v4575 = v4525
	v4577 = v4559
	goto L838
L846:
	;
	goto L847
L847:
	;
	v4563 = base.I32_wrap_i64(base.I64_clz(v4537))
	F___ashlti3(m, v4533, v4537, int64(0), v4563+int32(49))
	mBase = m.M
	v4567 = *(*int64)(unsafe.Add(mBase, uint32(v4533)+8))
	v4573 = *(*int64)(unsafe.Add(mBase, uint32(v4533)))
	v4574 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4563)
	v4575 = v4567 ^ int64(281474976710656)
	v4577 = v4573
	goto L838
L848:
	;
	v4599 = v2151 + int32(592)
	v4603 = int64(0)
	v4609 = m.G0
	v4611 = v4609 - int32(16)
	m.G0 = v4611
	v4613 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(v104), float64(0.75)))
	v4615 = v4613 & int64(4503599627370495)
	v4619 = int64(base.Ui64(v4613)>>(uint(int64(52))%64)) & int64(2047)
	if v4619 != v4603 {
		goto L853
	} else {
		goto L854
	}
L849:
	;
	goto L850
L850:
	;
	v4674 = base.F64_convert_i32_s(v104)
	if v4014 == (v3889+int32(5))&int32(2047) {
		goto L862
	} else {
		goto L863
	}
L851:
	;
	v4669 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+592))
	v4670 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+600))
	F___addtf3(m, v2151+int32(576), v4493, v4495, v4669, v4670)
	mBase = m.M
	v4672 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+584))
	v4673 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+576))
	v4830 = v4673
	v4831 = v4672
	goto L832
L852:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4599))) = v4655
	*(*int64)(unsafe.Add(mBase, uint32(v4599)+8)) = v4613&int64(-9223372036854775807-1) | v4652<<(uint(int64(48))%64) | v4653
	m.G0 = v4611 + int32(16)
	goto L851
L853:
	;
	if v4619 != int64(2047) {
		goto L856
	} else {
		goto L857
	}
L854:
	;
	goto L855
L855:
	;
	if v4615 == int64(0) {
		goto L859
	} else {
		goto L860
	}
L856:
	;
	v4652 = v4619 + int64(15360)
	v4653 = int64(base.Ui64(v4615) >> (uint(int64(4)) % 64))
	v4655 = v4615 << (uint(int64(60)) % 64)
	goto L852
L857:
	;
	goto L858
L858:
	;
	v4652 = int64(32767)
	v4653 = int64(base.Ui64(v4615) >> (uint(int64(4)) % 64))
	v4655 = v4615 << (uint(int64(60)) % 64)
	goto L852
L859:
	;
	v4637 = int64(0)
	v4652 = v4637
	v4653 = v4603
	v4655 = v4637
	goto L852
L860:
	;
	goto L861
L861:
	;
	v4641 = base.I32_wrap_i64(base.I64_clz(v4615))
	F___ashlti3(m, v4611, v4615, int64(0), v4641+int32(49))
	mBase = m.M
	v4645 = *(*int64)(unsafe.Add(mBase, uint32(v4611)+8))
	v4651 = *(*int64)(unsafe.Add(mBase, uint32(v4611)))
	v4652 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4641)
	v4653 = v4645 ^ int64(281474976710656)
	v4655 = v4651
	goto L852
L862:
	;
	v4681 = v2151 + int32(528)
	v4684 = int64(0)
	v4690 = m.G0
	v4692 = v4690 - int32(16)
	m.G0 = v4692
	v4694 = base.I64_reinterpret_f64(base.F64_mul(v4674, float64(0.5)))
	v4696 = v4694 & int64(4503599627370495)
	v4700 = int64(base.Ui64(v4694)>>(uint(int64(52))%64)) & int64(2047)
	if v4700 != v4684 {
		goto L867
	} else {
		goto L868
	}
L863:
	;
	goto L864
L864:
	;
	v4756 = v2151 + int32(560)
	v4759 = int64(0)
	v4765 = m.G0
	v4767 = v4765 - int32(16)
	m.G0 = v4767
	v4769 = base.I64_reinterpret_f64(base.F64_mul(v4674, float64(0.75)))
	v4771 = v4769 & int64(4503599627370495)
	v4775 = int64(base.Ui64(v4769)>>(uint(int64(52))%64)) & int64(2047)
	if v4775 != v4759 {
		goto L878
	} else {
		goto L879
	}
L865:
	;
	v4750 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+528))
	v4751 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+536))
	F___addtf3(m, v2151+int32(512), v4493, v4495, v4750, v4751)
	mBase = m.M
	v4753 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+520))
	v4754 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+512))
	v4830 = v4754
	v4831 = v4753
	goto L832
L866:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4681))) = v4736
	*(*int64)(unsafe.Add(mBase, uint32(v4681)+8)) = v4694&int64(-9223372036854775807-1) | v4733<<(uint(int64(48))%64) | v4734
	m.G0 = v4692 + int32(16)
	goto L865
L867:
	;
	if v4700 != int64(2047) {
		goto L870
	} else {
		goto L871
	}
L868:
	;
	goto L869
L869:
	;
	if v4696 == int64(0) {
		goto L873
	} else {
		goto L874
	}
L870:
	;
	v4733 = v4700 + int64(15360)
	v4734 = int64(base.Ui64(v4696) >> (uint(int64(4)) % 64))
	v4736 = v4696 << (uint(int64(60)) % 64)
	goto L866
L871:
	;
	goto L872
L872:
	;
	v4733 = int64(32767)
	v4734 = int64(base.Ui64(v4696) >> (uint(int64(4)) % 64))
	v4736 = v4696 << (uint(int64(60)) % 64)
	goto L866
L873:
	;
	v4718 = int64(0)
	v4733 = v4718
	v4734 = v4684
	v4736 = v4718
	goto L866
L874:
	;
	goto L875
L875:
	;
	v4722 = base.I32_wrap_i64(base.I64_clz(v4696))
	F___ashlti3(m, v4692, v4696, int64(0), v4722+int32(49))
	mBase = m.M
	v4726 = *(*int64)(unsafe.Add(mBase, uint32(v4692)+8))
	v4732 = *(*int64)(unsafe.Add(mBase, uint32(v4692)))
	v4733 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4722)
	v4734 = v4726 ^ int64(281474976710656)
	v4736 = v4732
	goto L866
L876:
	;
	v4825 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+560))
	v4826 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+568))
	F___addtf3(m, v2151+int32(544), v4493, v4495, v4825, v4826)
	mBase = m.M
	v4828 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+552))
	v4829 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+544))
	v4830 = v4829
	v4831 = v4828
	goto L832
L877:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4756))) = v4811
	*(*int64)(unsafe.Add(mBase, uint32(v4756)+8)) = v4769&int64(-9223372036854775807-1) | v4808<<(uint(int64(48))%64) | v4809
	m.G0 = v4767 + int32(16)
	goto L876
L878:
	;
	if v4775 != int64(2047) {
		goto L881
	} else {
		goto L882
	}
L879:
	;
	goto L880
L880:
	;
	if v4771 == int64(0) {
		goto L884
	} else {
		goto L885
	}
L881:
	;
	v4808 = v4775 + int64(15360)
	v4809 = int64(base.Ui64(v4771) >> (uint(int64(4)) % 64))
	v4811 = v4771 << (uint(int64(60)) % 64)
	goto L877
L882:
	;
	goto L883
L883:
	;
	v4808 = int64(32767)
	v4809 = int64(base.Ui64(v4771) >> (uint(int64(4)) % 64))
	v4811 = v4771 << (uint(int64(60)) % 64)
	goto L877
L884:
	;
	v4793 = int64(0)
	v4808 = v4793
	v4809 = v4759
	v4811 = v4793
	goto L877
L885:
	;
	goto L886
L886:
	;
	v4797 = base.I32_wrap_i64(base.I64_clz(v4771))
	F___ashlti3(m, v4767, v4771, int64(0), v4797+int32(49))
	mBase = m.M
	v4801 = *(*int64)(unsafe.Add(mBase, uint32(v4767)+8))
	v4807 = *(*int64)(unsafe.Add(mBase, uint32(v4767)))
	v4808 = base.I64_extend_i32_u(int32(_a_F___floatscan_6) - v4797)
	v4809 = v4801 ^ int64(281474976710656)
	v4811 = v4807
	goto L877
L887:
	;
	v4837 = int64(0)
	F_fmodl(m, v2151+int32(464), v4830, v4831, v4837, int64(4611404543450677248))
	mBase = m.M
	v4840 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+464))
	v4841 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+472))
	v4851 = v4841 & int64(9223372036854775807)
	v4852 = int64(9223090561878065152)
	if v4851 == v4852 {
		goto L890
	} else {
		goto L891
	}
L888:
	;
	if v4899 != 0 {
		v4908 = v4830
		v4909 = v4831
		goto L830
	} else {
		goto L916
	}
L889:
	;
	v4899 = v4895
	goto L888
L890:
	;
	v4856 = base.B2i32(v4840 != v4837)
	goto L892
L891:
	;
	v4856 = base.B2i32(base.Ui64(v4852) < base.Ui64(v4851))
	goto L892
L892:
	;
	if v4856 != 0 {
		v4895 = int32(1)
		goto L889
	} else {
		goto L893
	}
L893:
	;
	goto L895
L895:
	;
	goto L896
L896:
	;
	goto L897
L897:
	;
	if v4840|v4837|(v4851|int64(0)) == int64(0) {
		goto L898
	} else {
		goto L899
	}
L898:
	;
	v4899 = int32(0)
	goto L888
L899:
	;
	goto L900
L900:
	;
	if int64(0) <= v4841&v4837 {
		goto L901
	} else {
		goto L902
	}
L901:
	;
	if v4841 == v4837 {
		goto L904
	} else {
		goto L905
	}
L902:
	;
	goto L903
L903:
	;
	if v4841 == v4837 {
		goto L910
	} else {
		goto L911
	}
L904:
	;
	v4878 = base.B2i32(base.Ui64(v4840) < base.Ui64(v4837))
	goto L906
L905:
	;
	v4878 = base.B2i32(v4841 < v4837)
	goto L906
L906:
	;
	if v4878 != 0 {
		goto L907
	} else {
		goto L908
	}
L907:
	;
	v4899 = int32(-1)
	goto L888
L908:
	;
	goto L909
L909:
	;
	v4899 = base.B2i32(v4840^v4837|(v4841^v4837) != int64(0))
	goto L888
L910:
	;
	v4888 = base.B2i32(base.Ui64(v4837) < base.Ui64(v4840))
	goto L912
L911:
	;
	v4888 = base.B2i32(v4837 < v4841)
	goto L912
L912:
	;
	if v4888 != 0 {
		goto L913
	} else {
		goto L914
	}
L913:
	;
	v4899 = int32(-1)
	goto L888
L914:
	;
	goto L915
L915:
	;
	v4895 = base.B2i32(v4840^v4837|(v4841^v4837) != int64(0))
	goto L889
L916:
	;
	F___addtf3(m, v2151+int32(448), v4830, v4831, int64(0), int64(4611404543450677248))
	mBase = m.M
	v4905 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+456))
	v4906 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+448))
	v4908 = v4906
	v4909 = v4905
	goto L830
L917:
	;
	v4933 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+424))
	v4934 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+416))
	if v4127&int32(2147483647) <= v2155-int32(2) {
		v5087 = v3897
		v5088 = v4933
		v5089 = v4934
		goto L918
	} else {
		goto L919
	}
L918:
	;
	v5091 = v2151 + int32(368)
	v5093 = m.G0
	v5095 = v5093 - int32(80)
	m.G0 = v5095
	if int32(_a_F___floatscan_8) <= v5087 {
		goto L975
	} else {
		goto L976
	}
L919:
	;
	v4941 = v2151 + int32(400)
	v4942 = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(v4941)+8)) = v4933 & v4942
	*(*int64)(unsafe.Add(mBase, uint32(v4941))) = v4934
	v4948 = int64(0)
	F___multf3(m, v2151+int32(384), v4934, v4933, v4948, int64(4611123068473966592))
	mBase = m.M
	v4951 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+400))
	v4952 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+408))
	v4953 = int64(4643211215818981376)
	v4957 = int32(-1)
	v4961 = v4952 & v4942
	v4962 = int64(9223090561878065152)
	if v4961 == v4962 {
		goto L922
	} else {
		goto L923
	}
L920:
	;
	v5003 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+392))
	v5005 = base.B2i32(int32(0) <= v5002)
	if int32(0) <= v5002 {
		goto L938
	} else {
		goto L939
	}
L921:
	;
	v5002 = v4998
	goto L920
L922:
	;
	v4966 = base.B2i32(v4951 != v4948)
	goto L924
L923:
	;
	v4966 = base.B2i32(base.Ui64(v4962) < base.Ui64(v4961))
	goto L924
L924:
	;
	if v4966 != 0 {
		v4998 = v4957
		goto L921
	} else {
		goto L925
	}
L925:
	;
	goto L926
L926:
	;
	if v4951|(v4961|int64(4643211215818981376)) == int64(0) {
		goto L927
	} else {
		goto L928
	}
L927:
	;
	v5002 = int32(0)
	goto L920
L928:
	;
	goto L929
L929:
	;
	if int64(0) <= v4952&v4953 {
		goto L930
	} else {
		goto L931
	}
L930:
	;
	if base.B2i32(v4952 != v4953)&base.B2i32(v4952 < v4953) != 0 {
		v4998 = v4957
		goto L921
	} else {
		goto L933
	}
L931:
	;
	goto L932
L932:
	;
	if v4952 == v4953 {
		goto L934
	} else {
		goto L935
	}
L933:
	;
	v5002 = base.B2i32(v4951|(v4952^v4953) != int64(0))
	goto L920
L934:
	;
	v4993 = base.B2i32(v4951 != int64(0))
	goto L936
L935:
	;
	v4993 = base.B2i32(v4953 < v4952)
	goto L936
L936:
	;
	if v4993 != 0 {
		v4998 = v4957
		goto L921
	} else {
		goto L937
	}
L937:
	;
	v4998 = base.B2i32(v4951|(v4952^v4953) != int64(0))
	goto L921
L938:
	;
	v5006 = v5003
	goto L940
L939:
	;
	v5006 = v4933
	goto L940
L940:
	;
	v5007 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+384))
	if int32(0) <= v5002 {
		goto L941
	} else {
		goto L942
	}
L941:
	;
	v5008 = v5007
	goto L943
L942:
	;
	v5008 = v4934
	goto L943
L943:
	;
	v5014 = int64(0)
	v5023 = v4909 & int64(9223372036854775807)
	v5024 = int64(9223090561878065152)
	if v5023 == v5024 {
		goto L946
	} else {
		goto L947
	}
L944:
	;
	v5072 = int32(0)
	v5077 = v5005 + v3897
	if base.B2i32(v4133&(base.B2i32(v4134 != v4128)|base.B2i32(v5002 < int32(0)))&base.B2i32(v5071 != v5072) == v5072)&base.B2i32(v5077+int32(110) <= v2155) != 0 {
		v5087 = v5077
		v5088 = v5006
		v5089 = v5008
		goto L918
	} else {
		goto L972
	}
L945:
	;
	v5071 = v5067
	goto L944
L946:
	;
	v5028 = base.B2i32(v4908 != v5014)
	goto L948
L947:
	;
	v5028 = base.B2i32(base.Ui64(v5024) < base.Ui64(v5023))
	goto L948
L948:
	;
	if v5028 != 0 {
		v5067 = int32(1)
		goto L945
	} else {
		goto L949
	}
L949:
	;
	goto L951
L951:
	;
	goto L952
L952:
	;
	goto L953
L953:
	;
	if v4908|v5014|(v5023|int64(0)) == int64(0) {
		goto L954
	} else {
		goto L955
	}
L954:
	;
	v5071 = int32(0)
	goto L944
L955:
	;
	goto L956
L956:
	;
	if int64(0) <= v4909&v5014 {
		goto L957
	} else {
		goto L958
	}
L957:
	;
	if v4909 == v5014 {
		goto L960
	} else {
		goto L961
	}
L958:
	;
	goto L959
L959:
	;
	if v4909 == v5014 {
		goto L966
	} else {
		goto L967
	}
L960:
	;
	v5050 = base.B2i32(base.Ui64(v4908) < base.Ui64(v5014))
	goto L962
L961:
	;
	v5050 = base.B2i32(v4909 < v5014)
	goto L962
L962:
	;
	if v5050 != 0 {
		goto L963
	} else {
		goto L964
	}
L963:
	;
	v5071 = int32(-1)
	goto L944
L964:
	;
	goto L965
L965:
	;
	v5071 = base.B2i32(v4908^v5014|(v4909^v5014) != int64(0))
	goto L944
L966:
	;
	v5060 = base.B2i32(base.Ui64(v5014) < base.Ui64(v4908))
	goto L968
L967:
	;
	v5060 = base.B2i32(v5014 < v4909)
	goto L968
L968:
	;
	if v5060 != 0 {
		goto L969
	} else {
		goto L970
	}
L969:
	;
	v5071 = int32(-1)
	goto L944
L970:
	;
	goto L971
L971:
	;
	v5067 = base.B2i32(v4908^v5014|(v4909^v5014) != int64(0))
	goto L945
L972:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___floatscan[4])) = int32(68)
	v5087 = v5077
	v5088 = v5006
	v5089 = v5008
	goto L918
L973:
	;
	v5166 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+368))
	v5167 = *(*int64)(unsafe.Add(mBase, uint32(v2151)+376))
	v5190 = v5166
	v5197 = v5167
	goto L486
L974:
	;
	F___multf3(m, v5095, v5149, v5150, int64(0), base.I64_extend_i32_u(v5151+int32(_a_F___floatscan_9))<<(uint(int64(48))%64))
	mBase = m.M
	v5159 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5091)+8)) = v5159
	v5161 = *(*int64)(unsafe.Add(mBase, uint32(v5095)))
	*(*int64)(unsafe.Add(mBase, uint32(v5091))) = v5161
	m.G0 = v5095 + int32(80)
	goto L973
L975:
	;
	F___multf3(m, v5095+int32(32), v5089, v5088, int64(0), int64(9222809086901354496))
	mBase = m.M
	v5104 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+40))
	v5105 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+32))
	if base.Ui32(v5087) < base.Ui32(int32(_a_F___floatscan_2)) {
		goto L978
	} else {
		goto L979
	}
L976:
	;
	goto L977
L977:
	;
	if int32(-16383) < v5087 {
		v5149 = v5089
		v5150 = v5088
		v5151 = v5087
		goto L974
	} else {
		goto L984
	}
L978:
	;
	v5149 = v5105
	v5150 = v5104
	v5151 = v5087 - int32(_a_F___floatscan_9)
	goto L974
L979:
	;
	goto L980
L980:
	;
	F___multf3(m, v5095+int32(16), v5105, v5104, int64(0), int64(9222809086901354496))
	mBase = m.M
	v5115 = int32(_a_F___floatscan_10)
	if base.Ui32(v5115) <= base.Ui32(v5087) {
		goto L981
	} else {
		goto L982
	}
L981:
	;
	v5118 = v5115
	goto L983
L982:
	;
	v5118 = v5087
	goto L983
L983:
	;
	v5121 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+24))
	v5122 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+16))
	v5149 = v5122
	v5150 = v5121
	v5151 = v5118 - int32(_a_F___floatscan_11)
	goto L974
L984:
	;
	F___multf3(m, v5095-int32(-64), v5089, v5088, int64(0), int64(32088147345014784))
	mBase = m.M
	v5130 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+72))
	v5131 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+64))
	if base.Ui32(int32(-32652)) < base.Ui32(v5087) {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	v5149 = v5131
	v5150 = v5130
	v5151 = v5087 + int32(_a_F___floatscan_12)
	goto L974
L986:
	;
	goto L987
L987:
	;
	F___multf3(m, v5095+int32(48), v5131, v5130, int64(0), int64(32088147345014784))
	mBase = m.M
	v5141 = int32(-48920)
	if base.Ui32(v5087) <= base.Ui32(v5141) {
		goto L988
	} else {
		goto L989
	}
L988:
	;
	v5144 = v5141
	goto L990
L989:
	;
	v5144 = v5087
	goto L990
L990:
	;
	v5147 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+56))
	v5148 = *(*int64)(unsafe.Add(mBase, uint32(v5095)+48))
	v5149 = v5148
	v5150 = v5147
	v5151 = v5144 + int32(_a_F___floatscan_13)
	goto L974
}
