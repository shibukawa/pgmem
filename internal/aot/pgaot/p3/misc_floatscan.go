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
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v197 int64
	_ = v197
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v322 int64
	_ = v322
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v338 int64
	_ = v338
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v355 int64
	_ = v355
	var v362 int32
	_ = v362
	var v373 int32
	_ = v373
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v547 int64
	_ = v547
	var v550 int64
	_ = v550
	var v553 int32
	_ = v553
	var v567 int32
	_ = v567
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int64
	_ = v597
	var v600 int32
	_ = v600
	var v627 int64
	_ = v627
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v765 int64
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int64
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v809 int64
	_ = v809
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v835 int64
	_ = v835
	var v836 int64
	_ = v836
	var v837 int64
	_ = v837
	var v838 int64
	_ = v838
	var v839 int64
	_ = v839
	var v840 int64
	_ = v840
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int64
	_ = v894
	var v909 int64
	_ = v909
	var v911 int64
	_ = v911
	var v912 int64
	_ = v912
	var v925 int64
	_ = v925
	var v926 int64
	_ = v926
	var v927 int64
	_ = v927
	var v928 int64
	_ = v928
	var v930 int64
	_ = v930
	var v931 int64
	_ = v931
	var v933 int64
	_ = v933
	var v934 int64
	_ = v934
	var v944 int64
	_ = v944
	var v945 int64
	_ = v945
	var v947 int64
	_ = v947
	var v949 int64
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int64
	_ = v952
	var v953 int64
	_ = v953
	var v954 int64
	_ = v954
	var v955 int64
	_ = v955
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int64
	_ = v965
	var v966 int64
	_ = v966
	var v967 int64
	_ = v967
	var v968 int64
	_ = v968
	var v969 int64
	_ = v969
	var v970 int64
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int64
	_ = v982
	var v985 int32
	_ = v985
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1021 int32
	_ = v1021
	var v1025 int64
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int64
	_ = v1035
	var v1037 int64
	_ = v1037
	var v1041 int64
	_ = v1041
	var v1059 int64
	_ = v1059
	var v1072 int32
	_ = v1072
	var v1079 int64
	_ = v1079
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1086 int64
	_ = v1086
	var v1098 int64
	_ = v1098
	var v1099 int64
	_ = v1099
	var v1111 int32
	_ = v1111
	var v1123 int64
	_ = v1123
	var v1131 int32
	_ = v1131
	var v1133 int64
	_ = v1133
	var v1145 int32
	_ = v1145
	var v1168 int64
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int64
	_ = v1172
	var v1175 int64
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1196 int64
	_ = v1196
	var v1197 int64
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1207 int64
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1215 int64
	_ = v1215
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int64
	_ = v1225
	var v1227 int64
	_ = v1227
	var v1231 int64
	_ = v1231
	var v1249 int64
	_ = v1249
	var v1262 int32
	_ = v1262
	var v1269 int64
	_ = v1269
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
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1325 int64
	_ = v1325
	var v1340 int64
	_ = v1340
	var v1342 int64
	_ = v1342
	var v1343 int64
	_ = v1343
	var v1351 int64
	_ = v1351
	var v1352 int64
	_ = v1352
	var v1353 int64
	_ = v1353
	var v1354 int64
	_ = v1354
	var v1358 int64
	_ = v1358
	var v1359 int64
	_ = v1359
	var v1363 int64
	_ = v1363
	var v1364 int64
	_ = v1364
	var v1380 int32
	_ = v1380
	var v1391 int64
	_ = v1391
	var v1393 int64
	_ = v1393
	var v1395 int64
	_ = v1395
	var v1401 int64
	_ = v1401
	var v1404 int64
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1412 int64
	_ = v1412
	var v1413 int64
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1456 int64
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int64
	_ = v1459
	var v1460 int64
	_ = v1460
	var v1461 int64
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1467 int64
	_ = v1467
	var v1468 int64
	_ = v1468
	var v1469 int64
	_ = v1469
	var v1481 int32
	_ = v1481
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
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1536 int64
	_ = v1536
	var v1551 int64
	_ = v1551
	var v1553 int64
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1560 int64
	_ = v1560
	var v1561 int64
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 float64
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1571 float64
	_ = v1571
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1587 float64
	_ = v1587
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1600 float64
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1609 int64
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1619 int64
	_ = v1619
	var v1621 int64
	_ = v1621
	var v1625 int64
	_ = v1625
	var v1643 int64
	_ = v1643
	var v1656 int32
	_ = v1656
	var v1663 int64
	_ = v1663
	var v1666 int64
	_ = v1666
	var v1667 int64
	_ = v1667
	var v1668 int64
	_ = v1668
	var v1670 int64
	_ = v1670
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1704 int64
	_ = v1704
	var v1719 int64
	_ = v1719
	var v1721 int64
	_ = v1721
	var v1722 int64
	_ = v1722
	var v1728 int64
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int64
	_ = v1731
	var v1732 int64
	_ = v1732
	var v1733 int64
	_ = v1733
	var v1739 int64
	_ = v1739
	var v1753 int64
	_ = v1753
	var v1754 int64
	_ = v1754
	var v1755 int64
	_ = v1755
	var v1756 int64
	_ = v1756
	var v1757 int64
	_ = v1757
	var v1758 int64
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1765 int64
	_ = v1765
	var v1774 int64
	_ = v1774
	var v1775 int64
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1801 int32
	_ = v1801
	var v1811 int32
	_ = v1811
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1842 int32
	_ = v1842
	var v1847 int64
	_ = v1847
	var v1856 int64
	_ = v1856
	var v1858 int64
	_ = v1858
	var v1859 int64
	_ = v1859
	var v1867 int64
	_ = v1867
	var v1868 int64
	_ = v1868
	var v1872 int64
	_ = v1872
	var v1873 int64
	_ = v1873
	var v1878 int64
	_ = v1878
	var v1880 int64
	_ = v1880
	var v1884 int64
	_ = v1884
	var v1885 int64
	_ = v1885
	var v1886 int64
	_ = v1886
	var v1887 int64
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1891 int64
	_ = v1891
	var v1892 int64
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1901 int64
	_ = v1901
	var v1902 int64
	_ = v1902
	var v1908 int64
	_ = v1908
	var v1909 int64
	_ = v1909
	var v1910 int64
	_ = v1910
	var v1919 int64
	_ = v1919
	var v1920 int64
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1946 int32
	_ = v1946
	var v1956 int32
	_ = v1956
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1988 int64
	_ = v1988
	var v1989 int64
	_ = v1989
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2005 int64
	_ = v2005
	var v2006 int64
	_ = v2006
	var v2014 int64
	_ = v2014
	var v2015 int64
	_ = v2015
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2031 int64
	_ = v2031
	var v2032 int64
	_ = v2032
	var v2033 int64
	_ = v2033
	var v2034 int64
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2043 int64
	_ = v2043
	var v2045 int64
	_ = v2045
	var v2050 int64
	_ = v2050
	var v2051 int64
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2077 int64
	_ = v2077
	var v2092 int64
	_ = v2092
	var v2094 int64
	_ = v2094
	var v2095 int64
	_ = v2095
	var v2103 int64
	_ = v2103
	var v2104 int64
	_ = v2104
	var v2105 int64
	_ = v2105
	var v2106 int64
	_ = v2106
	var v2110 int64
	_ = v2110
	var v2111 int64
	_ = v2111
	var v2115 int64
	_ = v2115
	var v2116 int64
	_ = v2116
	var v2139 int64
	_ = v2139
	var v2145 int64
	_ = v2145
	var v2151 int64
	_ = v2151
	var v2152 int64
	_ = v2152
	var v2153 int64
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2247 int64
	_ = v2247
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2266 int64
	_ = v2266
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
	var v2290 int64
	_ = v2290
	var v2301 int32
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2311 int32
	_ = v2311
	var v2319 int64
	_ = v2319
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2355 int64
	_ = v2355
	var v2356 int64
	_ = v2356
	var v2371 int64
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int64
	_ = v2413
	var v2414 int64
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2451 int64
	_ = v2451
	var v2452 int64
	_ = v2452
	var v2459 int64
	_ = v2459
	var v2466 int64
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2472 int64
	_ = v2472
	var v2473 int64
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2480 int64
	_ = v2480
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2506 int64
	_ = v2506
	var v2507 int64
	_ = v2507
	var v2514 int64
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2522 int32
	_ = v2522
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2541 int64
	_ = v2541
	var v2542 int64
	_ = v2542
	var v2582 int64
	_ = v2582
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2623 int64
	_ = v2623
	var v2624 int64
	_ = v2624
	var v2631 int32
	_ = v2631
	var v2637 int64
	_ = v2637
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2647 int64
	_ = v2647
	var v2649 int64
	_ = v2649
	var v2653 int64
	_ = v2653
	var v2671 int64
	_ = v2671
	var v2684 int32
	_ = v2684
	var v2691 int64
	_ = v2691
	var v2694 int64
	_ = v2694
	var v2695 int64
	_ = v2695
	var v2696 int64
	_ = v2696
	var v2698 int64
	_ = v2698
	var v2710 int64
	_ = v2710
	var v2711 int64
	_ = v2711
	var v2719 int32
	_ = v2719
	var v2721 int32
	_ = v2721
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2738 int32
	_ = v2738
	var v2742 int64
	_ = v2742
	var v2757 int64
	_ = v2757
	var v2759 int64
	_ = v2759
	var v2760 int64
	_ = v2760
	var v2767 int32
	_ = v2767
	var v2770 int32
	_ = v2770
	var v2772 int32
	_ = v2772
	var v2780 int32
	_ = v2780
	var v2785 int64
	_ = v2785
	var v2794 int64
	_ = v2794
	var v2796 int64
	_ = v2796
	var v2797 int64
	_ = v2797
	var v2805 int64
	_ = v2805
	var v2806 int64
	_ = v2806
	var v2807 int64
	_ = v2807
	var v2808 int64
	_ = v2808
	var v2810 int64
	_ = v2810
	var v2811 int64
	_ = v2811
	var v2820 int32
	_ = v2820
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
	var v2837 int32
	_ = v2837
	var v2841 int64
	_ = v2841
	var v2856 int64
	_ = v2856
	var v2858 int64
	_ = v2858
	var v2859 int64
	_ = v2859
	var v2867 int64
	_ = v2867
	var v2868 int64
	_ = v2868
	var v2869 int64
	_ = v2869
	var v2870 int64
	_ = v2870
	var v2874 int64
	_ = v2874
	var v2875 int64
	_ = v2875
	var v2879 int64
	_ = v2879
	var v2880 int64
	_ = v2880
	var v2889 int32
	_ = v2889
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2906 int32
	_ = v2906
	var v2910 int64
	_ = v2910
	var v2925 int64
	_ = v2925
	var v2927 int64
	_ = v2927
	var v2928 int64
	_ = v2928
	var v2936 int64
	_ = v2936
	var v2937 int64
	_ = v2937
	var v2938 int64
	_ = v2938
	var v2939 int64
	_ = v2939
	var v2943 int64
	_ = v2943
	var v2944 int64
	_ = v2944
	var v2948 int64
	_ = v2948
	var v2949 int64
	_ = v2949
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2987 int32
	_ = v2987
	var v2989 int32
	_ = v2989
	var v3028 int32
	_ = v3028
	var v3051 int32
	_ = v3051
	var v3060 int32
	_ = v3060
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3081 int64
	_ = v3081
	var v3096 int64
	_ = v3096
	var v3098 int64
	_ = v3098
	var v3099 int64
	_ = v3099
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3120 int32
	_ = v3120
	var v3125 int64
	_ = v3125
	var v3134 int64
	_ = v3134
	var v3136 int64
	_ = v3136
	var v3137 int64
	_ = v3137
	var v3145 int64
	_ = v3145
	var v3146 int64
	_ = v3146
	var v3147 int64
	_ = v3147
	var v3148 int64
	_ = v3148
	var v3150 int64
	_ = v3150
	var v3151 int64
	_ = v3151
	var v3155 int32
	_ = v3155
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3167 int32
	_ = v3167
	var v3169 int32
	_ = v3169
	var v3172 int32
	_ = v3172
	var v3176 int64
	_ = v3176
	var v3191 int64
	_ = v3191
	var v3193 int64
	_ = v3193
	var v3194 int64
	_ = v3194
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3215 int32
	_ = v3215
	var v3220 int64
	_ = v3220
	var v3229 int64
	_ = v3229
	var v3231 int64
	_ = v3231
	var v3232 int64
	_ = v3232
	var v3240 int64
	_ = v3240
	var v3241 int64
	_ = v3241
	var v3242 int64
	_ = v3242
	var v3243 int64
	_ = v3243
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3253 int32
	_ = v3253
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3265 int32
	_ = v3265
	var v3267 int32
	_ = v3267
	var v3270 int32
	_ = v3270
	var v3274 int64
	_ = v3274
	var v3289 int64
	_ = v3289
	var v3291 int64
	_ = v3291
	var v3292 int64
	_ = v3292
	var v3300 int64
	_ = v3300
	var v3301 int64
	_ = v3301
	var v3302 int64
	_ = v3302
	var v3303 int64
	_ = v3303
	var v3305 int64
	_ = v3305
	var v3306 int64
	_ = v3306
	var v3311 int32
	_ = v3311
	var v3315 int32
	_ = v3315
	var v3317 int32
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3323 int32
	_ = v3323
	var v3325 int32
	_ = v3325
	var v3331 int32
	_ = v3331
	var v3333 int32
	_ = v3333
	var v3336 int32
	_ = v3336
	var v3340 int64
	_ = v3340
	var v3355 int64
	_ = v3355
	var v3357 int64
	_ = v3357
	var v3358 int64
	_ = v3358
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3370 int32
	_ = v3370
	var v3378 int32
	_ = v3378
	var v3383 int64
	_ = v3383
	var v3392 int64
	_ = v3392
	var v3394 int64
	_ = v3394
	var v3395 int64
	_ = v3395
	var v3403 int64
	_ = v3403
	var v3404 int64
	_ = v3404
	var v3405 int64
	_ = v3405
	var v3406 int64
	_ = v3406
	var v3409 int32
	_ = v3409
	var v3414 int32
	_ = v3414
	var v3418 int32
	_ = v3418
	var v3420 int32
	_ = v3420
	var v3426 int32
	_ = v3426
	var v3428 int32
	_ = v3428
	var v3431 int32
	_ = v3431
	var v3435 int64
	_ = v3435
	var v3450 int64
	_ = v3450
	var v3452 int64
	_ = v3452
	var v3453 int64
	_ = v3453
	var v3461 int64
	_ = v3461
	var v3462 int64
	_ = v3462
	var v3463 int64
	_ = v3463
	var v3464 int64
	_ = v3464
	var v3466 int64
	_ = v3466
	var v3467 int64
	_ = v3467
	var v3475 int32
	_ = v3475
	var v3501 int32
	_ = v3501
	var v3505 int32
	_ = v3505
	var v3508 int32
	_ = v3508
	var v3510 int32
	_ = v3510
	var v3518 int32
	_ = v3518
	var v3521 int32
	_ = v3521
	var v3524 int32
	_ = v3524
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3541 int32
	_ = v3541
	var v3543 int32
	_ = v3543
	var v3567 int32
	_ = v3567
	var v3568 int32
	_ = v3568
	var v3569 int32
	_ = v3569
	var v3570 int32
	_ = v3570
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3583 int32
	_ = v3583
	var v3586 int32
	_ = v3586
	var v3588 int32
	_ = v3588
	var v3601 int32
	_ = v3601
	var v3603 int32
	_ = v3603
	var v3608 int32
	_ = v3608
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3639 int32
	_ = v3639
	var v3660 int32
	_ = v3660
	var v3662 int32
	_ = v3662
	var v3667 int32
	_ = v3667
	var v3669 int32
	_ = v3669
	var v3695 int32
	_ = v3695
	var v3704 int32
	_ = v3704
	var v3726 int32
	_ = v3726
	var v3733 int32
	_ = v3733
	var v3737 int32
	_ = v3737
	var v3739 int32
	_ = v3739
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3768 int64
	_ = v3768
	var v3771 int64
	_ = v3771
	var v3775 int64
	_ = v3775
	var v3776 int64
	_ = v3776
	var v3781 int64
	_ = v3781
	var v3783 int32
	_ = v3783
	var v3787 int32
	_ = v3787
	var v3789 int32
	_ = v3789
	var v3793 int32
	_ = v3793
	var v3795 int32
	_ = v3795
	var v3800 int32
	_ = v3800
	var v3806 int32
	_ = v3806
	var v3809 int32
	_ = v3809
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3823 int32
	_ = v3823
	var v3826 int32
	_ = v3826
	var v3837 int32
	_ = v3837
	var v3839 int32
	_ = v3839
	var v3844 int32
	_ = v3844
	var v3846 int32
	_ = v3846
	var v3864 int32
	_ = v3864
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3876 int32
	_ = v3876
	var v3880 int32
	_ = v3880
	var v3885 int32
	_ = v3885
	var v3887 int32
	_ = v3887
	var v3909 int32
	_ = v3909
	var v3913 int32
	_ = v3913
	var v3920 int32
	_ = v3920
	var v3945 int32
	_ = v3945
	var v3969 int32
	_ = v3969
	var v3973 int32
	_ = v3973
	var v3976 int32
	_ = v3976
	var v3981 int32
	_ = v3981
	var v3985 int32
	_ = v3985
	var v3993 int64
	_ = v3993
	var v3997 int32
	_ = v3997
	var v4002 int32
	_ = v4002
	var v4016 int64
	_ = v4016
	var v4017 int64
	_ = v4017
	var v4026 int32
	_ = v4026
	var v4031 int32
	_ = v4031
	var v4037 int32
	_ = v4037
	var v4039 int32
	_ = v4039
	var v4045 int32
	_ = v4045
	var v4048 int32
	_ = v4048
	var v4050 int32
	_ = v4050
	var v4058 int32
	_ = v4058
	var v4063 int64
	_ = v4063
	var v4072 int64
	_ = v4072
	var v4074 int64
	_ = v4074
	var v4075 int64
	_ = v4075
	var v4088 int64
	_ = v4088
	var v4089 int64
	_ = v4089
	var v4090 int64
	_ = v4090
	var v4091 int64
	_ = v4091
	var v4093 int64
	_ = v4093
	var v4094 int64
	_ = v4094
	var v4096 int32
	_ = v4096
	var v4100 int32
	_ = v4100
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4112 int32
	_ = v4112
	var v4114 int32
	_ = v4114
	var v4117 int32
	_ = v4117
	var v4121 int64
	_ = v4121
	var v4136 int64
	_ = v4136
	var v4138 int64
	_ = v4138
	var v4139 int64
	_ = v4139
	var v4147 int64
	_ = v4147
	var v4148 int64
	_ = v4148
	var v4150 int64
	_ = v4150
	var v4152 int64
	_ = v4152
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4156 int32
	_ = v4156
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4167 int32
	_ = v4167
	var v4171 int32
	_ = v4171
	var v4179 int32
	_ = v4179
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4184 int32
	_ = v4184
	var v4208 int32
	_ = v4208
	var v4209 int32
	_ = v4209
	var v4211 int32
	_ = v4211
	var v4220 int32
	_ = v4220
	var v4221 int32
	_ = v4221
	var v4224 int32
	_ = v4224
	var v4226 int32
	_ = v4226
	var v4230 int32
	_ = v4230
	var v4241 int32
	_ = v4241
	var v4246 int32
	_ = v4246
	var v4247 float64
	_ = v4247
	var v4249 int32
	_ = v4249
	var v4253 float64
	_ = v4253
	var v4260 int32
	_ = v4260
	var v4263 int32
	_ = v4263
	var v4269 float64
	_ = v4269
	var v4276 int32
	_ = v4276
	var v4279 int32
	_ = v4279
	var v4282 float64
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4291 int64
	_ = v4291
	var v4297 int32
	_ = v4297
	var v4299 int32
	_ = v4299
	var v4301 int64
	_ = v4301
	var v4303 int64
	_ = v4303
	var v4307 int64
	_ = v4307
	var v4325 int64
	_ = v4325
	var v4338 int32
	_ = v4338
	var v4345 int64
	_ = v4345
	var v4348 int64
	_ = v4348
	var v4349 int64
	_ = v4349
	var v4350 int64
	_ = v4350
	var v4352 int64
	_ = v4352
	var v4365 int32
	_ = v4365
	var v4366 int64
	_ = v4366
	var v4367 int64
	_ = v4367
	var v4373 int64
	_ = v4373
	var v4387 int64
	_ = v4387
	var v4388 int64
	_ = v4388
	var v4390 int32
	_ = v4390
	var v4391 float64
	_ = v4391
	var v4393 int32
	_ = v4393
	var v4397 float64
	_ = v4397
	var v4404 int32
	_ = v4404
	var v4407 int32
	_ = v4407
	var v4413 float64
	_ = v4413
	var v4420 int32
	_ = v4420
	var v4423 int32
	_ = v4423
	var v4426 float64
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4435 int64
	_ = v4435
	var v4441 int32
	_ = v4441
	var v4443 int32
	_ = v4443
	var v4445 int64
	_ = v4445
	var v4447 int64
	_ = v4447
	var v4451 int64
	_ = v4451
	var v4469 int64
	_ = v4469
	var v4482 int32
	_ = v4482
	var v4489 int64
	_ = v4489
	var v4492 int64
	_ = v4492
	var v4493 int64
	_ = v4493
	var v4494 int64
	_ = v4494
	var v4496 int64
	_ = v4496
	var v4510 int64
	_ = v4510
	var v4511 int64
	_ = v4511
	var v4514 int32
	_ = v4514
	var v4515 int64
	_ = v4515
	var v4516 int64
	_ = v4516
	var v4518 int32
	_ = v4518
	var v4519 int32
	_ = v4519
	var v4520 int32
	_ = v4520
	var v4525 int64
	_ = v4525
	var v4526 int64
	_ = v4526
	var v4534 int64
	_ = v4534
	var v4535 int64
	_ = v4535
	var v4537 int64
	_ = v4537
	var v4538 int64
	_ = v4538
	var v4539 int64
	_ = v4539
	var v4540 int64
	_ = v4540
	var v4541 int64
	_ = v4541
	var v4542 int64
	_ = v4542
	var v4543 int64
	_ = v4543
	var v4544 int64
	_ = v4544
	var v4548 int32
	_ = v4548
	var v4555 int32
	_ = v4555
	var v4566 int32
	_ = v4566
	var v4570 int64
	_ = v4570
	var v4576 int32
	_ = v4576
	var v4578 int32
	_ = v4578
	var v4580 int64
	_ = v4580
	var v4582 int64
	_ = v4582
	var v4586 int64
	_ = v4586
	var v4604 int64
	_ = v4604
	var v4617 int32
	_ = v4617
	var v4624 int64
	_ = v4624
	var v4627 int64
	_ = v4627
	var v4628 int64
	_ = v4628
	var v4629 int64
	_ = v4629
	var v4631 int64
	_ = v4631
	var v4645 int64
	_ = v4645
	var v4646 int64
	_ = v4646
	var v4648 int64
	_ = v4648
	var v4649 int64
	_ = v4649
	var v4653 int32
	_ = v4653
	var v4657 int64
	_ = v4657
	var v4663 int32
	_ = v4663
	var v4665 int32
	_ = v4665
	var v4667 int64
	_ = v4667
	var v4669 int64
	_ = v4669
	var v4673 int64
	_ = v4673
	var v4691 int64
	_ = v4691
	var v4704 int32
	_ = v4704
	var v4711 int64
	_ = v4711
	var v4714 int64
	_ = v4714
	var v4715 int64
	_ = v4715
	var v4716 int64
	_ = v4716
	var v4718 int64
	_ = v4718
	var v4732 int64
	_ = v4732
	var v4733 int64
	_ = v4733
	var v4735 int64
	_ = v4735
	var v4736 int64
	_ = v4736
	var v4737 float64
	_ = v4737
	var v4744 int32
	_ = v4744
	var v4747 int64
	_ = v4747
	var v4753 int32
	_ = v4753
	var v4755 int32
	_ = v4755
	var v4757 int64
	_ = v4757
	var v4759 int64
	_ = v4759
	var v4763 int64
	_ = v4763
	var v4781 int64
	_ = v4781
	var v4794 int32
	_ = v4794
	var v4801 int64
	_ = v4801
	var v4804 int64
	_ = v4804
	var v4805 int64
	_ = v4805
	var v4806 int64
	_ = v4806
	var v4808 int64
	_ = v4808
	var v4822 int64
	_ = v4822
	var v4823 int64
	_ = v4823
	var v4825 int64
	_ = v4825
	var v4826 int64
	_ = v4826
	var v4828 int32
	_ = v4828
	var v4831 int64
	_ = v4831
	var v4837 int32
	_ = v4837
	var v4839 int32
	_ = v4839
	var v4841 int64
	_ = v4841
	var v4843 int64
	_ = v4843
	var v4847 int64
	_ = v4847
	var v4865 int64
	_ = v4865
	var v4878 int32
	_ = v4878
	var v4885 int64
	_ = v4885
	var v4888 int64
	_ = v4888
	var v4889 int64
	_ = v4889
	var v4890 int64
	_ = v4890
	var v4892 int64
	_ = v4892
	var v4906 int64
	_ = v4906
	var v4907 int64
	_ = v4907
	var v4909 int64
	_ = v4909
	var v4910 int64
	_ = v4910
	var v4911 int64
	_ = v4911
	var v4912 int64
	_ = v4912
	var v4918 int64
	_ = v4918
	var v4921 int64
	_ = v4921
	var v4922 int64
	_ = v4922
	var v4932 int64
	_ = v4932
	var v4933 int64
	_ = v4933
	var v4937 int32
	_ = v4937
	var v4959 int32
	_ = v4959
	var v4969 int32
	_ = v4969
	var v4976 int32
	_ = v4976
	var v4980 int32
	_ = v4980
	var v4986 int64
	_ = v4986
	var v4987 int64
	_ = v4987
	var v4989 int64
	_ = v4989
	var v4990 int64
	_ = v4990
	var v4996 int32
	_ = v4996
	var v4997 int64
	_ = v4997
	var v4998 int64
	_ = v4998
	var v5000 int32
	_ = v5000
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5007 int64
	_ = v5007
	var v5008 int64
	_ = v5008
	var v5014 int64
	_ = v5014
	var v5015 int64
	_ = v5015
	var v5022 int32
	_ = v5022
	var v5023 int64
	_ = v5023
	var v5029 int64
	_ = v5029
	var v5032 int64
	_ = v5032
	var v5033 int64
	_ = v5033
	var v5034 int64
	_ = v5034
	var v5038 int32
	_ = v5038
	var v5042 int64
	_ = v5042
	var v5043 int64
	_ = v5043
	var v5047 int32
	_ = v5047
	var v5074 int32
	_ = v5074
	var v5079 int32
	_ = v5079
	var v5083 int32
	_ = v5083
	var v5084 int64
	_ = v5084
	var v5086 int32
	_ = v5086
	var v5087 int64
	_ = v5087
	var v5088 int64
	_ = v5088
	var v5089 int64
	_ = v5089
	var v5090 int64
	_ = v5090
	var v5099 int64
	_ = v5099
	var v5100 int64
	_ = v5100
	var v5104 int32
	_ = v5104
	var v5126 int32
	_ = v5126
	var v5136 int32
	_ = v5136
	var v5143 int32
	_ = v5143
	var v5147 int32
	_ = v5147
	var v5148 int32
	_ = v5148
	var v5153 int32
	_ = v5153
	var v5168 int32
	_ = v5168
	var v5169 int64
	_ = v5169
	var v5170 int64
	_ = v5170
	var v5172 int32
	_ = v5172
	var v5174 int32
	_ = v5174
	var v5176 int32
	_ = v5176
	var v5185 int64
	_ = v5185
	var v5186 int64
	_ = v5186
	var v5196 int32
	_ = v5196
	var v5199 int32
	_ = v5199
	var v5202 int64
	_ = v5202
	var v5203 int64
	_ = v5203
	var v5211 int64
	_ = v5211
	var v5212 int64
	_ = v5212
	var v5222 int32
	_ = v5222
	var v5225 int32
	_ = v5225
	var v5228 int64
	_ = v5228
	var v5229 int64
	_ = v5229
	var v5230 int64
	_ = v5230
	var v5231 int64
	_ = v5231
	var v5232 int32
	_ = v5232
	var v5240 int64
	_ = v5240
	var v5242 int64
	_ = v5242
	var v5247 int64
	_ = v5247
	var v5248 int64
	_ = v5248
	var v5270 int64
	_ = v5270
	var v5277 int64
	_ = v5277
	var v5283 int64
	_ = v5283
	var v5284 int64
	_ = v5284
	var v5334 int64
	_ = v5334
	var v5363 int64
	_ = v5363
	var v5364 int64
	_ = v5364
	v5 = int32(0)
	v21 = int64(0)
	v29 = m.G0
	v31 = v29 - int32(48)
	m.G0 = v31
	if base.Ui32(int32(2)) < base.Ui32(l2) {
		v5334 = v21
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v5363
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v5364
	m.G0 = v31 + int32(48)
	return
L2:
	;
	v5363 = v5334
	v5364 = int64(0)
	goto L1
L3:
	;
	v36 = l2 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1334])))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+uint32(_consts[1335])))
	goto L4
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v71 != v72 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	switch v80 - int32(43) {
	case 0, 2:
		goto L15
	default:
		v105 = v80
		v106 = int32(1)
		goto L14
	}
L6:
	;
	goto L12
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v71 + int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v80 = v77
	goto L6
L8:
	;
	goto L9
L9:
	;
	v78 = F___shgetc(m, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v80 = v78
	goto L6
L12:
	;
	if base.B2i32(v80 == int32(32))|base.B2i32(base.Ui32(v80-int32(9)) < base.Ui32(int32(5))) != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	if v105&int32(-33) == int32(73) {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	if v80 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v95 = int32(-1)
	goto L18
L17:
	;
	v95 = int32(1)
	goto L18
L18:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v96 != v97 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v96 + int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v105 = v102
	v106 = v95
	goto L14
L20:
	;
	goto L21
L21:
	;
	v103 = F___shgetc(m, l1)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v105 = v103
	v106 = v95
	goto L14
L23:
	;
	if v165 != 0 {
		v417 = v161
		v421 = v165
		goto L74
	} else {
		goto L75
	}
L24:
	;
	v277 = m.G0
	v279 = v277 - int32(16)
	m.G0 = v279
	v284 = base.I32_reinterpret_f32(base.F32_mul(base.F32_convert_i32_s(v106), math.Float32frombits(uint32(0x7f800000))))
	v286 = v284 & int32(8388607)
	v288 = int32(base.Ui32(v284) >> (uint(int32(23)) % 32))
	v290 = v288 & int32(255)
	if v290 != 0 {
		goto L56
	} else {
		goto L57
	}
L25:
	;
	v117 = v5
	goto L28
L26:
	;
	v161 = v105
	v165 = v5
	goto L27
L27:
	;
	if v165 != int32(3) {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	if v117 == int32(7) {
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v161 = v150
	v165 = v154
	goto L27
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v141 != v142 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v154 = v117 + int32(1)
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117)+uint32(_consts[1336]))))
	if v155 == v150|int32(32) {
		v117 = v154
		goto L28
	} else {
		goto L36
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v141 + int32(1)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v150 = v147
	goto L31
L33:
	;
	goto L34
L34:
	;
	v148 = F___shgetc(m, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v150 = v148
	goto L31
L36:
	;
	goto L29
L37:
	;
	if v165 == int32(8) {
		goto L24
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v197 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	if l3 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(v165) < base.Ui32(int32(4)) {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if v165 == int32(8) {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v200 - int32(1)
	goto L46
L45:
	;
	goto L46
L46:
	;
	if l3 == int32(0) {
		goto L24
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v165) < base.Ui32(int32(4)) {
		goto L24
	} else {
		goto L48
	}
L48:
	;
	v216 = v165
	goto L49
L49:
	;
	if base.B2i32(v197 < int64(0)) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L24
L51:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v240 - int32(1)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v245 = v216 - int32(1)
	if base.Ui32(int32(3)) < base.Ui32(v245) {
		v216 = v245
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v343
	v355 = base.I64_extend_i32_u(v344)<<(uint(int64(48))%64) | base.I64_extend_i32_u(int32(base.Ui32(v284)>>(uint(int32(31))%32)))<<(uint(int64(63))%64) | v342
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v355
	m.G0 = v279 + int32(16)
	v5363 = v343
	v5364 = v355
	goto L1
L56:
	;
	if v290 != int32(255) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v304 = int32(0)
	if v286 == v304 {
		v342 = int64(0)
		v343 = v21
		v344 = v304
		goto L55
	} else {
		goto L62
	}
L59:
	;
	v342 = base.I64_extend_i32_u(v286) << (uint(int64(25)) % 64)
	v343 = v21
	v344 = v288&int32(255) + int32(16256)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v342 = base.I64_extend_i32_u(v286) << (uint(int64(25)) % 64)
	v343 = v21
	v344 = int32(32767)
	goto L55
L62:
	;
	v307 = base.I64_extend_i32_u(v286)
	v308 = int64(0)
	v309 = base.I32_clz(v286)
	v311 = v309 + int32(81)
	if v311&int32(64) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v279)+8))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v279)))
	v342 = v335 ^ int64(281474976710656)
	v343 = v338
	v344 = int32(16265) - v309
	goto L55
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v279))) = v330
	*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v331
	goto L63
L65:
	;
	v330 = int64(0)
	v331 = v307 << (uint(base.I64_extend_i32_u(v309+int32(17))) % 64)
	goto L64
L66:
	;
	goto L67
L67:
	;
	if v311 == int32(0) {
		v330 = v307
		v331 = v308
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v322 = base.I64_extend_i32_u(v311)
	v330 = v307 << (uint(v322) % 64)
	v331 = v308<<(uint(v322)%64) | int64(base.Ui64(v307)>>(uint(base.I64_extend_i32_u(int32(64)-v311))%64))
	goto L64
L69:
	;
	v5363 = int64(0)
	v5364 = v547
	goto L1
L70:
	;
	if v417 != int32(48) {
		goto L131
	} else {
		goto L132
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = int64(0)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v639 - v640)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L128
L72:
	;
	v597 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v597 {
		goto L124
	} else {
		goto L125
	}
L73:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v471 != v472 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	switch v421 {
	case 0:
		goto L70
	default:
		goto L72
	case 3:
		goto L73
	}
L75:
	;
	v362 = int32(0)
	if v161&int32(-33) != int32(78) {
		v417 = v161
		v421 = v362
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v373 = v362
	goto L77
L77:
	;
	if v373 == int32(2) {
		goto L73
	} else {
		goto L79
	}
L78:
	;
	v417 = v406
	v421 = v410
	goto L74
L79:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v397 != v398 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v410 = v373 + int32(1)
	v411 = int32(*(*int8)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[1337]))))
	if v411 == v406|int32(32) {
		v373 = v410
		goto L77
	} else {
		goto L85
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v397 + int32(1)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	v406 = v403
	goto L80
L82:
	;
	goto L83
L83:
	;
	v404 = F___shgetc(m, l1)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v406 = v404
	goto L80
L85:
	;
	goto L78
L86:
	;
	v498 = int32(1)
	goto L96
L87:
	;
	if v480 == int32(40) {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v471 + int32(1)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471))))
	v480 = v477
	goto L87
L89:
	;
	goto L90
L90:
	;
	v478 = F___shgetc(m, l1)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v480 = v478
	goto L87
L92:
	;
	goto L86
L93:
	;
	goto L94
L94:
	;
	v484 = int64(9223231299366420480)
	v485 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v485 < int64(0) {
		v5363 = v21
		v5364 = v484
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v488 - int32(1)
	v5363 = v21
	v5364 = v484
	goto L1
L96:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v520 != v521 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v547 = int64(9223231299366420480)
	if v529 == int32(41) {
		v5363 = v21
		v5364 = v547
		goto L1
	} else {
		goto L109
	}
L98:
	;
	if base.Ui32(v529-int32(48)) < base.Ui32(int32(10)) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v520 + int32(1)
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	v529 = v526
	goto L98
L100:
	;
	goto L101
L101:
	;
	v527 = F___shgetc(m, l1)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v529 = v527
	goto L98
L103:
	;
	goto L97
L104:
	;
	v498 = v498 + int32(1)
	goto L96
L105:
	;
	if base.Ui32(v529-int32(65)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	if v529 == int32(95) {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	if base.Ui32(int32(26)) <= base.Ui32(v529-int32(97)) {
		goto L103
	} else {
		goto L108
	}
L108:
	;
	goto L104
L109:
	;
	v550 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v550 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v553 - int32(1)
	goto L112
L111:
	;
	goto L112
L112:
	;
	if l3 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v567 = v498
	goto L118
L114:
	;
	if v498 != 0 {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	v627 = int64(0)
	goto L71
L117:
	;
	goto L69
L118:
	;
	if int64(0) <= v550 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L69
L120:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v591 - int32(1)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v596 = v567 - int32(1)
	if v596 != 0 {
		v567 = v596
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L119
L124:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v600 - int32(1)
	goto L126
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	v627 = v21
	goto L71
L127:
	;
	v5334 = v627
	goto L2
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v644
	goto L127
L131:
	;
	v2162 = v31 + int32(32)
	v2163 = int32(0)
	v2165 = m.G0
	v2167 = v2165 - int32(8976)
	m.G0 = v2167
	v2170 = v2163 - v39
	v2171 = v2170 - v42
	v2174 = v417
	v2179 = v2163
	goto L466
L132:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v656 != v657 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v665&int32(-33) == int32(88) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v656 + int32(1)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	v665 = v662
	goto L133
L135:
	;
	goto L136
L136:
	;
	v663 = F___shgetc(m, l1)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	v665 = v663
	goto L133
L138:
	;
	v671 = v31 + int32(16)
	v672 = m.G0
	v674 = v672 - int32(432)
	m.G0 = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v676 != v677 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	goto L140
L140:
	;
	v2153 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2153 < int64(0) {
		goto L131
	} else {
		goto L463
	}
L141:
	;
	v688 = v685
	v703 = v5
	goto L148
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v676 + int32(1)
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	v685 = v682
	goto L141
L143:
	;
	goto L144
L144:
	;
	v683 = F___shgetc(m, l1)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L10
	} else {
		goto L145
	}
L145:
	;
	v685 = v683
	goto L141
L146:
	;
	v817 = v788
	v823 = v5
	v824 = v5
	v826 = v797
	v832 = v803
	v835 = v21
	v836 = int64(4611404543450677248)
	v837 = v21
	v838 = v809
	v839 = v21
	v840 = v21
	goto L172
L147:
	;
	if v738 != int32(48) {
		goto L161
	} else {
		goto L162
	}
L148:
	;
	if v688 != int32(48) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v736 = F___shgetc(m, l1)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L10
	} else {
		goto L160
	}
L150:
	;
	goto L149
L151:
	;
	if v688 != int32(46) {
		v788 = v688
		v797 = v5
		v803 = v703
		v809 = v21
		goto L146
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v725 != v726 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v718 == v719 {
		goto L150
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v718 + int32(1)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	v738 = v724
	goto L147
L156:
	;
	v728 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v725 + v728
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725))))
	v688 = v732
	v703 = v728
	goto L148
L157:
	;
	goto L158
L158:
	;
	v734 = F___shgetc(m, l1)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	v688 = v734
	v703 = int32(1)
	goto L148
L160:
	;
	v738 = v736
	goto L147
L161:
	;
	v788 = v738
	v797 = int32(1)
	v803 = v703
	v809 = v21
	goto L146
L162:
	;
	goto L163
L163:
	;
	v765 = v21
	goto L164
L164:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v770 != v771 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v784 = int32(1)
	v788 = v779
	v797 = v784
	v803 = v784
	v809 = v781
	goto L146
L166:
	;
	v781 = v765 - int64(1)
	if v779 == int32(48) {
		v765 = v781
		goto L164
	} else {
		goto L171
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v770 + int32(1)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	v779 = v776
	goto L166
L168:
	;
	goto L169
L169:
	;
	v777 = F___shgetc(m, l1)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v779 = v777
	goto L166
L171:
	;
	goto L165
L172:
	;
	v844 = v817 - int32(48)
	if base.Ui32(v844) < base.Ui32(int32(10)) {
		v859 = v817
		goto L176
	} else {
		goto L177
	}
L173:
	;
	if v832 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L174:
	;
	goto L173
L175:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v971 != v972 {
		goto L198
	} else {
		goto L199
	}
L176:
	;
	if int32(57) < v817 {
		goto L181
	} else {
		goto L182
	}
L177:
	;
	v850 = v817 | int32(32)
	if base.B2i32(v817 != int32(46))&base.B2i32(base.Ui32(int32(5)) < base.Ui32(v850-int32(97))) != 0 {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	if v817 != int32(46) {
		v859 = v850
		goto L176
	} else {
		goto L179
	}
L179:
	;
	if v826 != 0 {
		goto L174
	} else {
		goto L180
	}
L180:
	;
	v960 = v823
	v961 = v824
	v962 = int32(1)
	v964 = v832
	v965 = v835
	v966 = v836
	v967 = v837
	v968 = v835
	v969 = v839
	v970 = v840
	goto L175
L181:
	;
	v864 = v859 - int32(87)
	goto L183
L182:
	;
	v864 = v844
	goto L183
L183:
	;
	if v835 <= int64(7) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v960 = v950
	v961 = v951
	v962 = v826
	v964 = int32(1)
	v965 = v835 + int64(1)
	v966 = v952
	v967 = v953
	v968 = v838
	v969 = v954
	v970 = v955
	goto L175
L185:
	;
	v950 = v823
	v951 = v864 + v824<<(uint(int32(4))%32)
	v952 = v836
	v953 = v837
	v954 = v839
	v955 = v840
	goto L184
L186:
	;
	goto L187
L187:
	;
	if base.Ui64(v835) <= base.Ui64(int64(28)) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v873 = v674 + int32(48)
	v877 = m.G0
	v879 = v877 - int32(16)
	m.G0 = v879
	if v864 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	goto L190
L190:
	;
	if v864 == int32(0) {
		v950 = v823
		v951 = v824
		v952 = v836
		v953 = v837
		v954 = v839
		v955 = v840
		goto L184
	} else {
		goto L196
	}
L191:
	;
	F___multf3(m, v674+int32(32), v840, v836, int64(0), int64(4610278643543834624))
	mBase = m.M
	v925 = *(*int64)(unsafe.Add(mBase, uint32(v674)+48))
	v926 = *(*int64)(unsafe.Add(mBase, uint32(v674)+56))
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v674)+32))
	v928 = *(*int64)(unsafe.Add(mBase, uint32(v674)+40))
	F___multf3(m, v674+int32(16), v925, v926, v927, v928)
	mBase = m.M
	v930 = *(*int64)(unsafe.Add(mBase, uint32(v674)+16))
	v931 = *(*int64)(unsafe.Add(mBase, uint32(v674)+24))
	F___addtf3(m, v674, v930, v931, v837, v839)
	mBase = m.M
	v933 = *(*int64)(unsafe.Add(mBase, uint32(v674)+8))
	v934 = *(*int64)(unsafe.Add(mBase, uint32(v674)))
	v950 = v823
	v951 = v824
	v952 = v928
	v953 = v934
	v954 = v933
	v955 = v927
	goto L184
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v873))) = v912
	*(*int64)(unsafe.Add(mBase, uint32(v873)+8)) = v911
	m.G0 = v879 + int32(16)
	goto L191
L193:
	;
	v911 = int64(0)
	v912 = int64(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v885 = v864 >> (uint(int32(31)) % 32)
	v887 = v864 ^ v885 - v885
	v890 = base.I32_clz(v887)
	F___ashlti3(m, v879, base.I64_extend_i32_u(v887), int64(0), v890+int32(81))
	mBase = m.M
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v879)+8))
	v909 = *(*int64)(unsafe.Add(mBase, uint32(v879)))
	v911 = v894 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v890)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v864&int32(-2147483648))<<(uint(int64(32))%64)
	v912 = v909
	goto L192
L196:
	;
	if v823 != 0 {
		v950 = v823
		v951 = v824
		v952 = v836
		v953 = v837
		v954 = v839
		v955 = v840
		goto L184
	} else {
		goto L197
	}
L197:
	;
	F___multf3(m, v674+int32(80), v840, v836, int64(0), int64(4611123068473966592))
	mBase = m.M
	v944 = *(*int64)(unsafe.Add(mBase, uint32(v674)+80))
	v945 = *(*int64)(unsafe.Add(mBase, uint32(v674)+88))
	F___addtf3(m, v674-int32(-64), v944, v945, v837, v839)
	mBase = m.M
	v947 = *(*int64)(unsafe.Add(mBase, uint32(v674)+72))
	v949 = *(*int64)(unsafe.Add(mBase, uint32(v674)+64))
	v950 = int32(1)
	v951 = v824
	v952 = v836
	v953 = v949
	v954 = v947
	v955 = v840
	goto L184
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v971 + int32(1)
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v817 = v977
	v823 = v960
	v824 = v961
	v826 = v962
	v832 = v964
	v835 = v965
	v836 = v966
	v837 = v967
	v838 = v968
	v839 = v969
	v840 = v970
	goto L172
L199:
	;
	goto L200
L200:
	;
	v978 = F___shgetc(m, l1)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	v817 = v978
	v823 = v960
	v824 = v961
	v826 = v962
	v832 = v964
	v835 = v965
	v836 = v966
	v837 = v967
	v838 = v968
	v839 = v969
	v840 = v970
	goto L172
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = v2139
	*(*int64)(unsafe.Add(mBase, uint32(v671)+8)) = v2145
	m.G0 = v674 + int32(432)
	v2151 = *(*int64)(unsafe.Add(mBase, uint32(v31)+24))
	v2152 = *(*int64)(unsafe.Add(mBase, uint32(v31)+16))
	v5363 = v2152
	v5364 = v2151
	goto L1
L203:
	;
	v982 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v982 {
		goto L208
	} else {
		goto L209
	}
L204:
	;
	goto L205
L205:
	;
	if v835 <= int64(7) {
		goto L232
	} else {
		goto L233
	}
L206:
	;
	v1021 = v674 + int32(96)
	v1025 = int64(0)
	v1031 = m.G0
	v1033 = v1031 - int32(16)
	m.G0 = v1033
	v1035 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v106)))
	v1037 = v1035 & int64(4503599627370495)
	v1041 = int64(base.Ui64(v1035)>>(uint(int64(52))%64)) & int64(2047)
	if v1041 != v1025 {
		goto L220
	} else {
		goto L221
	}
L207:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = int64(0)
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v1004 - v1005)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L215
L208:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v985 - int32(1)
	if l3 == int32(0) {
		goto L207
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if l3 != 0 {
		goto L206
	} else {
		goto L213
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v985 - int32(2)
	if v826 == int32(0) {
		goto L206
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v985 - int32(3)
	goto L206
L213:
	;
	goto L207
L214:
	;
	goto L206
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v1009
	goto L214
L218:
	;
	v1098 = *(*int64)(unsafe.Add(mBase, uint32(v674)+96))
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v674)+104))
	v2139 = v1098
	v2145 = v1099
	goto L202
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1021))) = v1086
	*(*int64)(unsafe.Add(mBase, uint32(v1021)+8)) = v1035&int64(-9223372036854775807-1) | v1083<<(uint(int64(48))%64) | v1084
	m.G0 = v1033 + int32(16)
	goto L218
L220:
	;
	if v1041 != int64(2047) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	if v1037 == int64(0) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v1083 = v1041 + int64(15360)
	v1084 = int64(base.Ui64(v1037) >> (uint(int64(4)) % 64))
	v1086 = v1037 << (uint(int64(60)) % 64)
	goto L219
L224:
	;
	goto L225
L225:
	;
	v1083 = int64(32767)
	v1084 = int64(base.Ui64(v1037) >> (uint(int64(4)) % 64))
	v1086 = v1037 << (uint(int64(60)) % 64)
	goto L219
L226:
	;
	v1059 = int64(0)
	v1083 = v1059
	v1084 = v1025
	v1086 = v1059
	goto L219
L227:
	;
	goto L228
L228:
	;
	if base.Ui64(v1037) < base.Ui64(int64(4294967296)) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1072 = base.I32_clz(base.I32_wrap_i64(v1035)) | int32(32)
	goto L231
L230:
	;
	v1072 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v1037) >> (uint(int64(32)) % 64))))
	goto L231
L231:
	;
	F___ashlti3(m, v1033, v1037, int64(0), v1072+int32(49))
	mBase = m.M
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v1033)+8))
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v1033)))
	v1083 = base.I64_extend_i32_u(int32(15372) - v1072)
	v1084 = v1079 ^ int64(281474976710656)
	v1086 = v1082
	goto L219
L232:
	;
	v1111 = v824
	v1123 = v835
	goto L235
L233:
	;
	v1145 = v824
	goto L234
L234:
	;
	if v817&int32(-33) == int32(80) {
		goto L241
	} else {
		goto L242
	}
L235:
	;
	v1131 = v1111 << (uint(int32(4)) % 32)
	v1133 = v1123 + int64(1)
	if v1133 != int64(8) {
		v1111 = v1131
		v1123 = v1133
		goto L235
	} else {
		goto L237
	}
L236:
	;
	v1145 = v1131
	goto L234
L237:
	;
	goto L236
L238:
	;
	if v1145 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L239:
	;
	v1207 = int64(0)
	goto L238
L240:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1201 - int32(1)
	goto L239
L241:
	;
	v1168 = F_scanexp(m, l1, l3)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L10
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1196 = int64(0)
	v1197 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v1197 < v1196 {
		v1207 = v1196
		goto L238
	} else {
		goto L254
	}
L244:
	;
	if v1168 != int64(-9223372036854775807-1) {
		v1207 = v1168
		goto L238
	} else {
		goto L245
	}
L245:
	;
	if l3 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1172 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if int64(0) <= v1172 {
		goto L240
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v1175 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v1175
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v1180 - v1181)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L251
L249:
	;
	goto L239
L250:
	;
	v2139 = v1175
	v2145 = int64(0)
	goto L202
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v1185
	goto L250
L254:
	;
	goto L240
L255:
	;
	v1211 = v674 + int32(112)
	v1215 = int64(0)
	v1221 = m.G0
	v1223 = v1221 - int32(16)
	m.G0 = v1223
	v1225 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v106)))
	v1227 = v1225 & int64(4503599627370495)
	v1231 = int64(base.Ui64(v1225)>>(uint(int64(52))%64)) & int64(2047)
	if v1231 != v1215 {
		goto L260
	} else {
		goto L261
	}
L256:
	;
	goto L257
L257:
	;
	if v826 != 0 {
		goto L272
	} else {
		goto L273
	}
L258:
	;
	v1288 = *(*int64)(unsafe.Add(mBase, uint32(v674)+112))
	v1289 = *(*int64)(unsafe.Add(mBase, uint32(v674)+120))
	v2139 = v1288
	v2145 = v1289
	goto L202
L259:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1211))) = v1276
	*(*int64)(unsafe.Add(mBase, uint32(v1211)+8)) = v1225&int64(-9223372036854775807-1) | v1273<<(uint(int64(48))%64) | v1274
	m.G0 = v1223 + int32(16)
	goto L258
L260:
	;
	if v1231 != int64(2047) {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	goto L262
L262:
	;
	if v1227 == int64(0) {
		goto L266
	} else {
		goto L267
	}
L263:
	;
	v1273 = v1231 + int64(15360)
	v1274 = int64(base.Ui64(v1227) >> (uint(int64(4)) % 64))
	v1276 = v1227 << (uint(int64(60)) % 64)
	goto L259
L264:
	;
	goto L265
L265:
	;
	v1273 = int64(32767)
	v1274 = int64(base.Ui64(v1227) >> (uint(int64(4)) % 64))
	v1276 = v1227 << (uint(int64(60)) % 64)
	goto L259
L266:
	;
	v1249 = int64(0)
	v1273 = v1249
	v1274 = v1215
	v1276 = v1249
	goto L259
L267:
	;
	goto L268
L268:
	;
	if base.Ui64(v1227) < base.Ui64(int64(4294967296)) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1262 = base.I32_clz(base.I32_wrap_i64(v1225)) | int32(32)
	goto L271
L270:
	;
	v1262 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v1227) >> (uint(int64(32)) % 64))))
	goto L271
L271:
	;
	F___ashlti3(m, v1223, v1227, int64(0), v1262+int32(49))
	mBase = m.M
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v1223)+8))
	v1272 = *(*int64)(unsafe.Add(mBase, uint32(v1223)))
	v1273 = base.I64_extend_i32_u(int32(15372) - v1262)
	v1274 = v1269 ^ int64(281474976710656)
	v1276 = v1272
	goto L259
L272:
	;
	v1290 = v838
	goto L274
L273:
	;
	v1290 = v835
	goto L274
L274:
	;
	v1295 = v1290<<(uint(int64(2))%64) + v1207 - int64(32)
	if base.I64_extend_i32_u(int32(0)-v39) < v1295 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	v1304 = v674 + int32(160)
	v1308 = m.G0
	v1310 = v1308 - int32(16)
	m.G0 = v1310
	if v106 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L276:
	;
	goto L277
L277:
	;
	if base.I64_extend_i32_s(v39-int32(226)) <= v1295 {
		goto L283
	} else {
		goto L284
	}
L278:
	;
	v1351 = *(*int64)(unsafe.Add(mBase, uint32(v674)+160))
	v1352 = *(*int64)(unsafe.Add(mBase, uint32(v674)+168))
	v1353 = int64(-1)
	v1354 = int64(9223090561878065151)
	F___multf3(m, v674+int32(144), v1351, v1352, v1353, v1354)
	mBase = m.M
	v1358 = *(*int64)(unsafe.Add(mBase, uint32(v674)+144))
	v1359 = *(*int64)(unsafe.Add(mBase, uint32(v674)+152))
	F___multf3(m, v674+int32(128), v1358, v1359, v1353, v1354)
	mBase = m.M
	v1363 = *(*int64)(unsafe.Add(mBase, uint32(v674)+128))
	v1364 = *(*int64)(unsafe.Add(mBase, uint32(v674)+136))
	v2139 = v1363
	v2145 = v1364
	goto L202
L279:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1304))) = v1343
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+8)) = v1342
	m.G0 = v1310 + int32(16)
	goto L278
L280:
	;
	v1342 = int64(0)
	v1343 = int64(0)
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1316 = v106 >> (uint(int32(31)) % 32)
	v1318 = v106 ^ v1316 - v1316
	v1321 = base.I32_clz(v1318)
	F___ashlti3(m, v1310, base.I64_extend_i32_u(v1318), int64(0), v1321+int32(81))
	mBase = m.M
	v1325 = *(*int64)(unsafe.Add(mBase, uint32(v1310)+8))
	v1340 = *(*int64)(unsafe.Add(mBase, uint32(v1310)))
	v1342 = v1325 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1321)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v1343 = v1340
	goto L279
L283:
	;
	if int32(0) <= v1145 {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	v2056 = v674 + int32(208)
	v2060 = m.G0
	v2062 = v2060 - int32(16)
	m.G0 = v2062
	if v106 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L286:
	;
	v1380 = v1145
	v1391 = v1295
	v1393 = v837
	v1395 = v839
	goto L289
L287:
	;
	v1481 = v1145
	v1492 = v1295
	v1494 = v837
	v1496 = v839
	goto L288
L288:
	;
	v1503 = v1492 + base.I64_extend_i32_u(int32(32)-v39)
	v1504 = base.I32_wrap_i64(v1503)
	v1505 = int32(0)
	if v1505 < v1504 {
		goto L317
	} else {
		goto L318
	}
L289:
	;
	v1401 = int64(0)
	F___addtf3(m, v674+int32(416), v1393, v1395, v1401, int64(-4611967493404098560))
	mBase = m.M
	v1404 = int64(4611123068473966592)
	v1408 = int32(-1)
	v1412 = v1395 & int64(9223372036854775807)
	v1413 = int64(9223090561878065152)
	if v1412 == v1413 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	v1481 = v1465
	v1492 = v1467
	v1494 = v1469
	v1496 = v1468
	goto L288
L291:
	;
	v1456 = *(*int64)(unsafe.Add(mBase, uint32(v674)+416))
	v1458 = base.B2i32(int32(0) <= v1453)
	if int32(0) <= v1453 {
		goto L309
	} else {
		goto L310
	}
L292:
	;
	v1453 = v1449
	goto L291
L293:
	;
	v1417 = base.B2i32(v1393 != v1401)
	goto L295
L294:
	;
	v1417 = base.B2i32(base.Ui64(v1413) < base.Ui64(v1412))
	goto L295
L295:
	;
	if v1417 != 0 {
		v1449 = v1408
		goto L292
	} else {
		goto L296
	}
L296:
	;
	goto L297
L297:
	;
	if v1393|(v1412|int64(4611123068473966592)) == int64(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1453 = int32(0)
	goto L291
L299:
	;
	goto L300
L300:
	;
	if int64(0) <= v1395&v1404 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if base.B2i32(v1395 != v1404)&base.B2i32(v1395 < v1404) != 0 {
		v1449 = v1408
		goto L292
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	if v1395 == v1404 {
		goto L305
	} else {
		goto L306
	}
L304:
	;
	v1453 = base.B2i32(v1393|(v1395^v1404) != int64(0))
	goto L291
L305:
	;
	v1444 = base.B2i32(v1393 != int64(0))
	goto L307
L306:
	;
	v1444 = base.B2i32(v1404 < v1395)
	goto L307
L307:
	;
	if v1444 != 0 {
		v1449 = v1408
		goto L292
	} else {
		goto L308
	}
L308:
	;
	v1449 = base.B2i32(v1393|(v1395^v1404) != int64(0))
	goto L292
L309:
	;
	v1459 = v1456
	goto L311
L310:
	;
	v1459 = v1393
	goto L311
L311:
	;
	v1460 = *(*int64)(unsafe.Add(mBase, uint32(v674)+424))
	if int32(0) <= v1453 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1461 = v1460
	goto L314
L313:
	;
	v1461 = v1395
	goto L314
L314:
	;
	F___addtf3(m, v674+int32(400), v1393, v1395, v1459, v1461)
	mBase = m.M
	v1464 = v1380 << (uint(int32(1)) % 32)
	v1465 = v1464 | v1458
	v1467 = v1391 - int64(1)
	v1468 = *(*int64)(unsafe.Add(mBase, uint32(v674)+408))
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v674)+400))
	if int32(0) <= v1464 {
		v1380 = v1465
		v1391 = v1467
		v1393 = v1469
		v1395 = v1468
		goto L289
	} else {
		goto L315
	}
L315:
	;
	goto L290
L316:
	;
	v1760 = v674 + int32(320)
	v1761 = int32(1)
	v1765 = int64(0)
	v1774 = v1496 & int64(9223372036854775807)
	v1775 = int64(9223090561878065152)
	if v1774 == v1775 {
		goto L371
	} else {
		goto L372
	}
L317:
	;
	v1508 = v1504
	goto L319
L318:
	;
	v1508 = v1505
	goto L319
L319:
	;
	if v1503 < base.I64_extend_i32_u(v42) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1511 = v1508
	goto L322
L321:
	;
	v1511 = v42
	goto L322
L322:
	;
	if base.Ui32(int32(113)) <= base.Ui32(v1511) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1515 = v674 + int32(384)
	v1519 = m.G0
	v1521 = v1519 - int32(16)
	m.G0 = v1521
	if v106 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L324:
	;
	goto L325
L325:
	;
	v1564 = v674 + int32(352)
	v1565 = float64(1)
	v1567 = int32(144) - v1511
	if int32(1024) <= v1567 {
		goto L333
	} else {
		goto L334
	}
L326:
	;
	v1560 = *(*int64)(unsafe.Add(mBase, uint32(v674)+392))
	v1561 = *(*int64)(unsafe.Add(mBase, uint32(v674)+384))
	v1755 = v1560
	v1756 = v1561
	v1757 = v21
	v1758 = int64(0)
	goto L316
L327:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1515))) = v1554
	*(*int64)(unsafe.Add(mBase, uint32(v1515)+8)) = v1553
	m.G0 = v1521 + int32(16)
	goto L326
L328:
	;
	v1553 = int64(0)
	v1554 = int64(0)
	goto L327
L329:
	;
	goto L330
L330:
	;
	v1527 = v106 >> (uint(int32(31)) % 32)
	v1529 = v106 ^ v1527 - v1527
	v1532 = base.I32_clz(v1529)
	F___ashlti3(m, v1521, base.I64_extend_i32_u(v1529), int64(0), v1532+int32(81))
	mBase = m.M
	v1536 = *(*int64)(unsafe.Add(mBase, uint32(v1521)+8))
	v1551 = *(*int64)(unsafe.Add(mBase, uint32(v1521)))
	v1553 = v1536 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1532)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v1554 = v1551
	goto L327
L331:
	;
	v1609 = int64(0)
	v1615 = m.G0
	v1617 = v1615 - int32(16)
	m.G0 = v1617
	v1619 = base.I64_reinterpret_f64(base.F64_mul(v1600, base.F64_reinterpret_i64(base.I64_extend_i32_u(v1601+int32(1023))<<(uint(int64(52))%64))))
	v1621 = v1619 & int64(4503599627370495)
	v1625 = int64(base.Ui64(v1619)>>(uint(int64(52))%64)) & int64(2047)
	if v1625 != v1609 {
		goto L351
	} else {
		goto L352
	}
L332:
	;
	goto L331
L333:
	;
	v1571 = base.F64_mul(v1565, float64(8.98846567431158e+307))
	if base.Ui32(v1567) < base.Ui32(int32(2047)) {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	goto L335
L335:
	;
	if int32(-1023) < v1567 {
		v1600 = v1565
		v1601 = v1567
		goto L332
	} else {
		goto L342
	}
L336:
	;
	v1600 = v1571
	v1601 = v1567 - int32(1023)
	goto L332
L337:
	;
	goto L338
L338:
	;
	v1578 = int32(3069)
	if base.Ui32(v1578) <= base.Ui32(v1567) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1581 = v1578
	goto L341
L340:
	;
	v1581 = v1567
	goto L341
L341:
	;
	v1600 = base.F64_mul(v1571, float64(8.98846567431158e+307))
	v1601 = v1581 - int32(2046)
	goto L332
L342:
	;
	v1587 = base.F64_mul(v1565, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v1567) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1600 = v1587
	v1601 = v1567 + int32(969)
	goto L332
L344:
	;
	goto L345
L345:
	;
	v1594 = int32(-2960)
	if base.Ui32(v1567) <= base.Ui32(v1594) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1597 = v1594
	goto L348
L347:
	;
	v1597 = v1567
	goto L348
L348:
	;
	v1600 = base.F64_mul(v1587, float64(2.004168360008973e-292))
	v1601 = v1597 + int32(1938)
	goto L332
L349:
	;
	v1683 = v674 + int32(336)
	v1687 = m.G0
	v1689 = v1687 - int32(16)
	m.G0 = v1689
	if v106 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L350:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1564))) = v1670
	*(*int64)(unsafe.Add(mBase, uint32(v1564)+8)) = v1619&int64(-9223372036854775807-1) | v1667<<(uint(int64(48))%64) | v1668
	m.G0 = v1617 + int32(16)
	goto L349
L351:
	;
	if v1625 != int64(2047) {
		goto L354
	} else {
		goto L355
	}
L352:
	;
	goto L353
L353:
	;
	if v1621 == int64(0) {
		goto L357
	} else {
		goto L358
	}
L354:
	;
	v1667 = v1625 + int64(15360)
	v1668 = int64(base.Ui64(v1621) >> (uint(int64(4)) % 64))
	v1670 = v1621 << (uint(int64(60)) % 64)
	goto L350
L355:
	;
	goto L356
L356:
	;
	v1667 = int64(32767)
	v1668 = int64(base.Ui64(v1621) >> (uint(int64(4)) % 64))
	v1670 = v1621 << (uint(int64(60)) % 64)
	goto L350
L357:
	;
	v1643 = int64(0)
	v1667 = v1643
	v1668 = v1609
	v1670 = v1643
	goto L350
L358:
	;
	goto L359
L359:
	;
	if base.Ui64(v1621) < base.Ui64(int64(4294967296)) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1656 = base.I32_clz(base.I32_wrap_i64(v1619)) | int32(32)
	goto L362
L361:
	;
	v1656 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v1621) >> (uint(int64(32)) % 64))))
	goto L362
L362:
	;
	F___ashlti3(m, v1617, v1621, int64(0), v1656+int32(49))
	mBase = m.M
	v1663 = *(*int64)(unsafe.Add(mBase, uint32(v1617)+8))
	v1666 = *(*int64)(unsafe.Add(mBase, uint32(v1617)))
	v1667 = base.I64_extend_i32_u(int32(15372) - v1656)
	v1668 = v1663 ^ int64(281474976710656)
	v1670 = v1666
	goto L350
L363:
	;
	v1728 = *(*int64)(unsafe.Add(mBase, uint32(v674)+336))
	v1730 = v674 + int32(368)
	v1731 = *(*int64)(unsafe.Add(mBase, uint32(v674)+352))
	v1732 = *(*int64)(unsafe.Add(mBase, uint32(v674)+360))
	v1733 = *(*int64)(unsafe.Add(mBase, uint32(v674)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v1730))) = v1731
	v1739 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v1730)+8)) = v1732&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v1732&int64(9223090561878065152))>>(uint(v1739)%64)))|base.I32_wrap_i64(int64(base.Ui64(v1733)>>(uint(v1739)%64)))&int32(32768))<<(uint(v1739)%64)
	goto L368
L364:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1683))) = v1722
	*(*int64)(unsafe.Add(mBase, uint32(v1683)+8)) = v1721
	m.G0 = v1689 + int32(16)
	goto L363
L365:
	;
	v1721 = int64(0)
	v1722 = int64(0)
	goto L364
L366:
	;
	goto L367
L367:
	;
	v1695 = v106 >> (uint(int32(31)) % 32)
	v1697 = v106 ^ v1695 - v1695
	v1700 = base.I32_clz(v1697)
	F___ashlti3(m, v1689, base.I64_extend_i32_u(v1697), int64(0), v1700+int32(81))
	mBase = m.M
	v1704 = *(*int64)(unsafe.Add(mBase, uint32(v1689)+8))
	v1719 = *(*int64)(unsafe.Add(mBase, uint32(v1689)))
	v1721 = v1704 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1700)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v1722 = v1719
	goto L364
L368:
	;
	v1753 = *(*int64)(unsafe.Add(mBase, uint32(v674)+376))
	v1754 = *(*int64)(unsafe.Add(mBase, uint32(v674)+368))
	v1755 = v1733
	v1756 = v1728
	v1757 = v1753
	v1758 = v1754
	goto L316
L369:
	;
	v1823 = int32(0)
	v1828 = base.B2i32(v1481&v1761 == int32(0)) & (base.B2i32(v1822 != v1823) & base.B2i32(base.Ui32(v1511) < base.Ui32(int32(32))))
	v1829 = v1481 | v1828
	v1832 = m.G0
	v1834 = v1832 - int32(16)
	m.G0 = v1834
	if v1829 == v1823 {
		goto L399
	} else {
		goto L400
	}
L370:
	;
	v1822 = v1818
	goto L369
L371:
	;
	v1779 = base.B2i32(v1494 != v1765)
	goto L373
L372:
	;
	v1779 = base.B2i32(base.Ui64(v1775) < base.Ui64(v1774))
	goto L373
L373:
	;
	if v1779 != 0 {
		v1818 = v1761
		goto L370
	} else {
		goto L374
	}
L374:
	;
	goto L376
L376:
	;
	goto L377
L377:
	;
	goto L378
L378:
	;
	if v1494|v1765|(v1774|int64(0)) == int64(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1822 = int32(0)
	goto L369
L380:
	;
	goto L381
L381:
	;
	if int64(0) <= v1496&v1765 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	if v1496 == v1765 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	goto L384
L384:
	;
	if v1496 == v1765 {
		goto L391
	} else {
		goto L392
	}
L385:
	;
	v1801 = base.B2i32(base.Ui64(v1494) < base.Ui64(v1765))
	goto L387
L386:
	;
	v1801 = base.B2i32(v1496 < v1765)
	goto L387
L387:
	;
	if v1801 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1822 = int32(-1)
	goto L369
L389:
	;
	goto L390
L390:
	;
	v1822 = base.B2i32(v1494^v1765|(v1496^v1765) != int64(0))
	goto L369
L391:
	;
	v1811 = base.B2i32(base.Ui64(v1765) < base.Ui64(v1494))
	goto L393
L392:
	;
	v1811 = base.B2i32(v1765 < v1496)
	goto L393
L393:
	;
	if v1811 != 0 {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1822 = int32(-1)
	goto L369
L395:
	;
	goto L396
L396:
	;
	v1818 = base.B2i32(v1494^v1765|(v1496^v1765) != int64(0))
	goto L370
L397:
	;
	v1867 = *(*int64)(unsafe.Add(mBase, uint32(v674)+320))
	v1868 = *(*int64)(unsafe.Add(mBase, uint32(v674)+328))
	F___multf3(m, v674+int32(304), v1756, v1755, v1867, v1868)
	mBase = m.M
	v1872 = *(*int64)(unsafe.Add(mBase, uint32(v674)+304))
	v1873 = *(*int64)(unsafe.Add(mBase, uint32(v674)+312))
	F___addtf3(m, v674+int32(272), v1872, v1873, v1758, v1757)
	mBase = m.M
	if v1828 != 0 {
		goto L402
	} else {
		goto L403
	}
L398:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1760))) = v1859
	*(*int64)(unsafe.Add(mBase, uint32(v1760)+8)) = v1858
	m.G0 = v1834 + int32(16)
	goto L397
L399:
	;
	v1858 = int64(0)
	v1859 = int64(0)
	goto L398
L400:
	;
	goto L401
L401:
	;
	v1842 = base.I32_clz(v1829)
	F___ashlti3(m, v1834, base.I64_extend_i32_u(v1829), int64(0), int32(112)-(v1842^int32(31)))
	mBase = m.M
	v1847 = *(*int64)(unsafe.Add(mBase, uint32(v1834)+8))
	v1856 = *(*int64)(unsafe.Add(mBase, uint32(v1834)))
	v1858 = v1847 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v1842)<<(uint(int64(48))%64)
	v1859 = v1856
	goto L398
L402:
	;
	v1878 = int64(0)
	goto L404
L403:
	;
	v1878 = v1494
	goto L404
L404:
	;
	if v1828 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1880 = int64(0)
	goto L407
L406:
	;
	v1880 = v1496
	goto L407
L407:
	;
	F___multf3(m, v674+int32(288), v1756, v1755, v1878, v1880)
	mBase = m.M
	v1884 = *(*int64)(unsafe.Add(mBase, uint32(v674)+288))
	v1885 = *(*int64)(unsafe.Add(mBase, uint32(v674)+296))
	v1886 = *(*int64)(unsafe.Add(mBase, uint32(v674)+272))
	v1887 = *(*int64)(unsafe.Add(mBase, uint32(v674)+280))
	F___addtf3(m, v674+int32(256), v1884, v1885, v1886, v1887)
	mBase = m.M
	v1890 = v674 + int32(240)
	v1891 = *(*int64)(unsafe.Add(mBase, uint32(v674)+256))
	v1892 = *(*int64)(unsafe.Add(mBase, uint32(v674)+264))
	v1894 = m.G0
	v1895 = int32(16)
	v1896 = v1894 - v1895
	m.G0 = v1896
	F___addtf3(m, v1896, v1891, v1892, v1758, v1757^int64(-9223372036854775807-1))
	mBase = m.M
	v1901 = *(*int64)(unsafe.Add(mBase, uint32(v1896)))
	v1902 = *(*int64)(unsafe.Add(mBase, uint32(v1896)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1890)+8)) = v1902
	*(*int64)(unsafe.Add(mBase, uint32(v1890))) = v1901
	m.G0 = v1896 + v1895
	goto L408
L408:
	;
	v1908 = *(*int64)(unsafe.Add(mBase, uint32(v674)+240))
	v1909 = *(*int64)(unsafe.Add(mBase, uint32(v674)+248))
	v1910 = int64(0)
	v1919 = v1909 & int64(9223372036854775807)
	v1920 = int64(9223090561878065152)
	if v1919 == v1920 {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	if v1967 == int32(0) {
		goto L437
	} else {
		goto L438
	}
L410:
	;
	v1967 = v1963
	goto L409
L411:
	;
	v1924 = base.B2i32(v1908 != v1910)
	goto L413
L412:
	;
	v1924 = base.B2i32(base.Ui64(v1920) < base.Ui64(v1919))
	goto L413
L413:
	;
	if v1924 != 0 {
		v1963 = int32(1)
		goto L410
	} else {
		goto L414
	}
L414:
	;
	goto L416
L416:
	;
	goto L417
L417:
	;
	goto L418
L418:
	;
	if v1908|v1910|(v1919|int64(0)) == int64(0) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1967 = int32(0)
	goto L409
L420:
	;
	goto L421
L421:
	;
	if int64(0) <= v1909&v1910 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if v1909 == v1910 {
		goto L425
	} else {
		goto L426
	}
L423:
	;
	goto L424
L424:
	;
	if v1909 == v1910 {
		goto L431
	} else {
		goto L432
	}
L425:
	;
	v1946 = base.B2i32(base.Ui64(v1908) < base.Ui64(v1910))
	goto L427
L426:
	;
	v1946 = base.B2i32(v1909 < v1910)
	goto L427
L427:
	;
	if v1946 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1967 = int32(-1)
	goto L409
L429:
	;
	goto L430
L430:
	;
	v1967 = base.B2i32(v1908^v1910|(v1909^v1910) != int64(0))
	goto L409
L431:
	;
	v1956 = base.B2i32(base.Ui64(v1910) < base.Ui64(v1908))
	goto L433
L432:
	;
	v1956 = base.B2i32(v1910 < v1909)
	goto L433
L433:
	;
	if v1956 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1967 = int32(-1)
	goto L409
L435:
	;
	goto L436
L436:
	;
	v1963 = base.B2i32(v1908^v1910|(v1909^v1910) != int64(0))
	goto L410
L437:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	goto L439
L438:
	;
	goto L439
L439:
	;
	v1974 = v674 + int32(224)
	v1975 = base.I32_wrap_i64(v1492)
	v1977 = m.G0
	v1979 = v1977 - int32(80)
	m.G0 = v1979
	if int32(16384) <= v1975 {
		goto L442
	} else {
		goto L443
	}
L440:
	;
	v2050 = *(*int64)(unsafe.Add(mBase, uint32(v674)+224))
	v2051 = *(*int64)(unsafe.Add(mBase, uint32(v674)+232))
	v2139 = v2050
	v2145 = v2051
	goto L202
L441:
	;
	F___multf3(m, v1979, v2033, v2034, int64(0), base.I64_extend_i32_u(v2035+int32(16383))<<(uint(int64(48))%64))
	mBase = m.M
	v2043 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1974)+8)) = v2043
	v2045 = *(*int64)(unsafe.Add(mBase, uint32(v1979)))
	*(*int64)(unsafe.Add(mBase, uint32(v1974))) = v2045
	m.G0 = v1979 + int32(80)
	goto L440
L442:
	;
	F___multf3(m, v1979+int32(32), v1908, v1909, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1988 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+40))
	v1989 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+32))
	if base.Ui32(v1975) < base.Ui32(int32(32767)) {
		goto L445
	} else {
		goto L446
	}
L443:
	;
	goto L444
L444:
	;
	if int32(-16383) < v1975 {
		v2033 = v1908
		v2034 = v1909
		v2035 = v1975
		goto L441
	} else {
		goto L451
	}
L445:
	;
	v2033 = v1989
	v2034 = v1988
	v2035 = v1975 - int32(16383)
	goto L441
L446:
	;
	goto L447
L447:
	;
	F___multf3(m, v1979+int32(16), v1989, v1988, int64(0), int64(9222809086901354496))
	mBase = m.M
	v1999 = int32(49149)
	if base.Ui32(v1999) <= base.Ui32(v1975) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2002 = v1999
	goto L450
L449:
	;
	v2002 = v1975
	goto L450
L450:
	;
	v2005 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+24))
	v2006 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+16))
	v2033 = v2006
	v2034 = v2005
	v2035 = v2002 - int32(32766)
	goto L441
L451:
	;
	F___multf3(m, v1979-int32(-64), v1908, v1909, int64(0), int64(32088147345014784))
	mBase = m.M
	v2014 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+72))
	v2015 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+64))
	if base.Ui32(int32(-32652)) < base.Ui32(v1975) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v2033 = v2015
	v2034 = v2014
	v2035 = v1975 + int32(16269)
	goto L441
L453:
	;
	goto L454
L454:
	;
	F___multf3(m, v1979+int32(48), v2015, v2014, int64(0), int64(32088147345014784))
	mBase = m.M
	v2025 = int32(-48920)
	if base.Ui32(v1975) <= base.Ui32(v2025) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2028 = v2025
	goto L457
L456:
	;
	v2028 = v1975
	goto L457
L457:
	;
	v2031 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+56))
	v2032 = *(*int64)(unsafe.Add(mBase, uint32(v1979)+48))
	v2033 = v2032
	v2034 = v2031
	v2035 = v2028 + int32(32538)
	goto L441
L458:
	;
	v2103 = *(*int64)(unsafe.Add(mBase, uint32(v674)+208))
	v2104 = *(*int64)(unsafe.Add(mBase, uint32(v674)+216))
	v2105 = int64(0)
	v2106 = int64(281474976710656)
	F___multf3(m, v674+int32(192), v2103, v2104, v2105, v2106)
	mBase = m.M
	v2110 = *(*int64)(unsafe.Add(mBase, uint32(v674)+192))
	v2111 = *(*int64)(unsafe.Add(mBase, uint32(v674)+200))
	F___multf3(m, v674+int32(176), v2110, v2111, v2105, v2106)
	mBase = m.M
	v2115 = *(*int64)(unsafe.Add(mBase, uint32(v674)+176))
	v2116 = *(*int64)(unsafe.Add(mBase, uint32(v674)+184))
	v2139 = v2115
	v2145 = v2116
	goto L202
L459:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2056))) = v2095
	*(*int64)(unsafe.Add(mBase, uint32(v2056)+8)) = v2094
	m.G0 = v2062 + int32(16)
	goto L458
L460:
	;
	v2094 = int64(0)
	v2095 = int64(0)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v2068 = v106 >> (uint(int32(31)) % 32)
	v2070 = v106 ^ v2068 - v2068
	v2073 = base.I32_clz(v2070)
	F___ashlti3(m, v2062, base.I64_extend_i32_u(v2070), int64(0), v2073+int32(81))
	mBase = m.M
	v2077 = *(*int64)(unsafe.Add(mBase, uint32(v2062)+8))
	v2092 = *(*int64)(unsafe.Add(mBase, uint32(v2062)))
	v2094 = v2077 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2073)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v2095 = v2092
	goto L459
L463:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2156 - int32(1)
	goto L131
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2167)+784)) = int32(0)
	v2330 = v2301 - int32(48)
	v2332 = base.B2i32(v2301 == int32(46))
	if v2301 == int32(46) {
		goto L496
	} else {
		goto L497
	}
L465:
	;
	if v2224 == int32(48) {
		goto L479
	} else {
		goto L480
	}
L466:
	;
	if v2174 != int32(48) {
		goto L469
	} else {
		goto L470
	}
L467:
	;
	v2222 = F___shgetc(m, l1)
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L10
	} else {
		goto L478
	}
L468:
	;
	goto L467
L469:
	;
	if v2174 != int32(46) {
		v2301 = v2174
		v2306 = v2179
		v2311 = v5
		v2319 = v21
		goto L464
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2211 != v2212 {
		goto L474
	} else {
		goto L475
	}
L472:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2204 == v2205 {
		goto L468
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2204 + int32(1)
	v2210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2204))))
	v2224 = v2210
	goto L465
L474:
	;
	v2214 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2211 + v2214
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2211))))
	v2174 = v2218
	v2179 = v2214
	goto L466
L475:
	;
	goto L476
L476:
	;
	v2220 = F___shgetc(m, l1)
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L10
	} else {
		goto L477
	}
L477:
	;
	v2174 = v2220
	v2179 = int32(1)
	goto L466
L478:
	;
	v2224 = v2222
	goto L465
L479:
	;
	v2247 = v21
	goto L482
L480:
	;
	v2272 = v2224
	v2277 = v2179
	v2290 = v21
	goto L481
L481:
	;
	v2301 = v2272
	v2306 = v2277
	v2311 = int32(1)
	v2319 = v2290
	goto L464
L482:
	;
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2255 != v2256 {
		goto L485
	} else {
		goto L486
	}
L483:
	;
	v2272 = v2264
	v2277 = int32(1)
	v2290 = v2266
	goto L481
L484:
	;
	v2266 = v2247 - int64(1)
	if v2264 == int32(48) {
		v2247 = v2266
		goto L482
	} else {
		goto L489
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2255 + int32(1)
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255))))
	v2264 = v2261
	goto L484
L486:
	;
	goto L487
L487:
	;
	v2262 = F___shgetc(m, l1)
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L10
	} else {
		goto L488
	}
L488:
	;
	v2264 = v2262
	goto L484
L489:
	;
	goto L483
L490:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2162)+8)) = v5277
	*(*int64)(unsafe.Add(mBase, uint32(v2162))) = v5270
	m.G0 = v2167 + int32(8976)
	v5283 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
	v5284 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
	v5363 = v5284
	v5364 = v5283
	goto L1
L491:
	;
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+784))
	if v2631 == int32(0) {
		goto L546
	} else {
		goto L547
	}
L492:
	;
	v2582 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v2582
	v2587 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = base.I64_extend_i32_s(v2587 - v2588)
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	goto L543
L493:
	;
	if v2522 == int32(0) {
		v2608 = v2526
		v2611 = v2529
		v2613 = v2531
		v2623 = v2541
		v2624 = v2542
		goto L491
	} else {
		goto L541
	}
L494:
	;
	v2514 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2514 < int64(0) {
		v2522 = v2487
		v2526 = v2491
		v2529 = v2494
		v2531 = v2496
		v2541 = v2506
		v2542 = v2507
		goto L493
	} else {
		goto L540
	}
L495:
	;
	if v2443 != 0 {
		goto L528
	} else {
		goto L529
	}
L496:
	;
	v2336 = v2332
	v2337 = v2301
	v2338 = v2330
	v2340 = v2163
	v2342 = v2306
	v2343 = v5
	v2345 = v5
	v2347 = v2311
	v2355 = v2319
	v2356 = v21
	goto L499
L497:
	;
	if base.Ui32(v2330) <= base.Ui32(int32(9)) {
		goto L496
	} else {
		goto L498
	}
L498:
	;
	v2433 = v2301
	v2436 = v2163
	v2438 = v2306
	v2439 = v5
	v2441 = v5
	v2443 = v2311
	v2451 = v2319
	v2452 = v21
	goto L495
L499:
	;
	if v2336&int32(1) != 0 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	v2433 = v2424
	v2436 = v2408
	v2438 = v2409
	v2439 = v2410
	v2441 = v2411
	v2443 = v2412
	v2451 = v2413
	v2452 = v2414
	goto L495
L501:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	if v2415 != v2416 {
		goto L522
	} else {
		goto L523
	}
L502:
	;
	if v2347 == int32(0) {
		goto L505
	} else {
		goto L506
	}
L503:
	;
	goto L504
L504:
	;
	v2371 = v2356 + int64(1)
	if v2340 <= int32(2044) {
		goto L508
	} else {
		goto L509
	}
L505:
	;
	v2408 = v2340
	v2409 = v2342
	v2410 = v2343
	v2411 = v2345
	v2412 = int32(1)
	v2413 = v2356
	v2414 = v2356
	goto L501
L506:
	;
	goto L507
L507:
	;
	v2487 = base.B2i32(v2342 == int32(0))
	v2491 = v2340
	v2494 = v2343
	v2496 = v2345
	v2506 = v2355
	v2507 = v2356
	goto L494
L508:
	;
	if v2337 == int32(48) {
		goto L511
	} else {
		goto L512
	}
L509:
	;
	goto L510
L510:
	;
	if v2337 == int32(48) {
		v2408 = v2340
		v2409 = v2342
		v2410 = v2343
		v2411 = v2345
		v2412 = v2347
		v2413 = v2355
		v2414 = v2371
		goto L501
	} else {
		goto L520
	}
L511:
	;
	v2377 = v2345
	goto L513
L512:
	;
	v2377 = base.I32_wrap_i64(v2371)
	goto L513
L513:
	;
	v2382 = v2167 + int32(784) + v2340<<(uint(int32(2))%32)
	if v2343 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2382)))
	v2389 = v2337 + v2383*int32(10) - int32(48)
	goto L516
L515:
	;
	v2389 = v2338
	goto L516
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2382))) = v2389
	v2391 = int32(1)
	v2394 = v2343 + v2391
	v2396 = base.B2i32(v2394 == int32(9))
	if v2394 == int32(9) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2397 = int32(0)
	goto L519
L518:
	;
	v2397 = v2394
	goto L519
L519:
	;
	v2408 = v2396 + v2340
	v2409 = v2391
	v2410 = v2397
	v2411 = v2377
	v2412 = v2347
	v2413 = v2355
	v2414 = v2371
	goto L501
L520:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+uint32(_consts[1338])))
	*(*int32)(unsafe.Add(mBase, uint32(v2167)+uint32(_consts[1338]))) = v2401 | int32(1)
	v2408 = v2340
	v2409 = v2342
	v2410 = v2343
	v2411 = int32(18396)
	v2412 = v2347
	v2413 = v2355
	v2414 = v2371
	goto L501
L521:
	;
	v2426 = v2424 - int32(48)
	v2428 = base.B2i32(v2424 == int32(46))
	if v2424 == int32(46) {
		v2336 = v2428
		v2337 = v2424
		v2338 = v2426
		v2340 = v2408
		v2342 = v2409
		v2343 = v2410
		v2345 = v2411
		v2347 = v2412
		v2355 = v2413
		v2356 = v2414
		goto L499
	} else {
		goto L526
	}
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2415 + int32(1)
	v2421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2415))))
	v2424 = v2421
	goto L521
L523:
	;
	goto L524
L524:
	;
	v2422 = F___shgetc(m, l1)
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L10
	} else {
		goto L525
	}
L525:
	;
	v2424 = v2422
	goto L521
L526:
	;
	if base.Ui32(v2426) < base.Ui32(int32(10)) {
		v2336 = v2428
		v2337 = v2424
		v2338 = v2426
		v2340 = v2408
		v2342 = v2409
		v2343 = v2410
		v2345 = v2411
		v2347 = v2412
		v2355 = v2413
		v2356 = v2414
		goto L499
	} else {
		goto L527
	}
L527:
	;
	goto L500
L528:
	;
	v2459 = v2451
	goto L530
L529:
	;
	v2459 = v2452
	goto L530
L530:
	;
	if v2438 == int32(0) {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v2482 = int32(0)
	v2483 = base.B2i32(v2438 == v2482)
	if v2433 < v2482 {
		v2522 = v2483
		v2526 = v2436
		v2529 = v2439
		v2531 = v2441
		v2541 = v2459
		v2542 = v2452
		goto L493
	} else {
		goto L539
	}
L532:
	;
	if v2433&int32(-33) != int32(69) {
		goto L531
	} else {
		goto L533
	}
L533:
	;
	v2466 = F_scanexp(m, l1, l3)
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L10
	} else {
		goto L535
	}
L534:
	;
	v2608 = v2436
	v2611 = v2439
	v2613 = v2441
	v2623 = v2459 + v2480
	v2624 = v2452
	goto L491
L535:
	;
	if v2466 != int64(-9223372036854775807-1) {
		v2480 = v2466
		goto L534
	} else {
		goto L536
	}
L536:
	;
	if l3 == int32(0) {
		goto L492
	} else {
		goto L537
	}
L537:
	;
	v2472 = int64(0)
	v2473 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	if v2473 < v2472 {
		v2480 = v2472
		goto L534
	} else {
		goto L538
	}
L538:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2476 - int32(1)
	v2480 = v2472
	goto L534
L539:
	;
	v2487 = v2483
	v2491 = v2436
	v2494 = v2439
	v2496 = v2441
	v2506 = v2459
	v2507 = v2452
	goto L494
L540:
	;
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2517 - int32(1)
	v2522 = v2487
	v2526 = v2491
	v2529 = v2494
	v2531 = v2496
	v2541 = v2506
	v2542 = v2507
	goto L493
L541:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(28)
	goto L492
L542:
	;
	v5270 = v2582
	v5277 = int64(0)
	goto L490
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+104)) = v2592
	goto L542
L546:
	;
	v2637 = int64(0)
	v2643 = m.G0
	v2645 = v2643 - int32(16)
	m.G0 = v2645
	v2647 = base.I64_reinterpret_f64(base.F64_copysign(float64(0), base.F64_convert_i32_s(v106)))
	v2649 = v2647 & int64(4503599627370495)
	v2653 = int64(base.Ui64(v2647)>>(uint(int64(52))%64)) & int64(2047)
	if v2653 != v2637 {
		goto L551
	} else {
		goto L552
	}
L547:
	;
	goto L548
L548:
	;
	if int64(9) < v2624 {
		goto L563
	} else {
		goto L564
	}
L549:
	;
	v2710 = *(*int64)(unsafe.Add(mBase, uint32(v2167)))
	v2711 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+8))
	v5270 = v2710
	v5277 = v2711
	goto L490
L550:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2167))) = v2698
	*(*int64)(unsafe.Add(mBase, uint32(v2167)+8)) = v2647&int64(-9223372036854775807-1) | v2695<<(uint(int64(48))%64) | v2696
	m.G0 = v2645 + int32(16)
	goto L549
L551:
	;
	if v2653 != int64(2047) {
		goto L554
	} else {
		goto L555
	}
L552:
	;
	goto L553
L553:
	;
	if v2649 == int64(0) {
		goto L557
	} else {
		goto L558
	}
L554:
	;
	v2695 = v2653 + int64(15360)
	v2696 = int64(base.Ui64(v2649) >> (uint(int64(4)) % 64))
	v2698 = v2649 << (uint(int64(60)) % 64)
	goto L550
L555:
	;
	goto L556
L556:
	;
	v2695 = int64(32767)
	v2696 = int64(base.Ui64(v2649) >> (uint(int64(4)) % 64))
	v2698 = v2649 << (uint(int64(60)) % 64)
	goto L550
L557:
	;
	v2671 = int64(0)
	v2695 = v2671
	v2696 = v2637
	v2698 = v2671
	goto L550
L558:
	;
	goto L559
L559:
	;
	if base.Ui64(v2649) < base.Ui64(int64(4294967296)) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2684 = base.I32_clz(base.I32_wrap_i64(v2647)) | int32(32)
	goto L562
L561:
	;
	v2684 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v2649) >> (uint(int64(32)) % 64))))
	goto L562
L562:
	;
	F___ashlti3(m, v2645, v2649, int64(0), v2684+int32(49))
	mBase = m.M
	v2691 = *(*int64)(unsafe.Add(mBase, uint32(v2645)+8))
	v2694 = *(*int64)(unsafe.Add(mBase, uint32(v2645)))
	v2695 = base.I64_extend_i32_u(int32(15372) - v2684)
	v2696 = v2691 ^ int64(281474976710656)
	v2698 = v2694
	goto L550
L563:
	;
	if base.I64_extend_i32_u(int32(base.Ui32(v2170)>>(uint(int32(1))%32))) < v2623 {
		goto L580
	} else {
		goto L581
	}
L564:
	;
	if v2623 != v2624 {
		goto L563
	} else {
		goto L565
	}
L565:
	;
	if int32(base.Ui32(v2631)>>(uint(v42)%32)) != 0 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v2719 = base.B2i32(base.Ui32(v42) <= base.Ui32(int32(30)))
	goto L568
L567:
	;
	v2719 = int32(0)
	goto L568
L568:
	;
	if v2719 != 0 {
		goto L563
	} else {
		goto L569
	}
L569:
	;
	v2721 = v2167 + int32(48)
	v2725 = m.G0
	v2727 = v2725 - int32(16)
	m.G0 = v2727
	if v106 == int32(0) {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	v2767 = v2167 + int32(32)
	v2770 = m.G0
	v2772 = v2770 - int32(16)
	m.G0 = v2772
	if v2631 == int32(0) {
		goto L577
	} else {
		goto L578
	}
L571:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2721))) = v2760
	*(*int64)(unsafe.Add(mBase, uint32(v2721)+8)) = v2759
	m.G0 = v2727 + int32(16)
	goto L570
L572:
	;
	v2759 = int64(0)
	v2760 = int64(0)
	goto L571
L573:
	;
	goto L574
L574:
	;
	v2733 = v106 >> (uint(int32(31)) % 32)
	v2735 = v106 ^ v2733 - v2733
	v2738 = base.I32_clz(v2735)
	F___ashlti3(m, v2727, base.I64_extend_i32_u(v2735), int64(0), v2738+int32(81))
	mBase = m.M
	v2742 = *(*int64)(unsafe.Add(mBase, uint32(v2727)+8))
	v2757 = *(*int64)(unsafe.Add(mBase, uint32(v2727)))
	v2759 = v2742 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2738)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v2760 = v2757
	goto L571
L575:
	;
	v2805 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+48))
	v2806 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+56))
	v2807 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+32))
	v2808 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+40))
	F___multf3(m, v2167+int32(16), v2805, v2806, v2807, v2808)
	mBase = m.M
	v2810 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+16))
	v2811 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+24))
	v5270 = v2810
	v5277 = v2811
	goto L490
L576:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2767))) = v2797
	*(*int64)(unsafe.Add(mBase, uint32(v2767)+8)) = v2796
	m.G0 = v2772 + int32(16)
	goto L575
L577:
	;
	v2796 = int64(0)
	v2797 = int64(0)
	goto L576
L578:
	;
	goto L579
L579:
	;
	v2780 = base.I32_clz(v2631)
	F___ashlti3(m, v2772, base.I64_extend_i32_u(v2631), int64(0), int32(112)-(v2780^int32(31)))
	mBase = m.M
	v2785 = *(*int64)(unsafe.Add(mBase, uint32(v2772)+8))
	v2794 = *(*int64)(unsafe.Add(mBase, uint32(v2772)))
	v2796 = v2785 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2780)<<(uint(int64(48))%64)
	v2797 = v2794
	goto L576
L580:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	v2820 = v2167 + int32(96)
	v2824 = m.G0
	v2826 = v2824 - int32(16)
	m.G0 = v2826
	if v106 == int32(0) {
		goto L585
	} else {
		goto L586
	}
L581:
	;
	goto L582
L582:
	;
	if v2623 < base.I64_extend_i32_s(v39-int32(226)) {
		goto L588
	} else {
		goto L589
	}
L583:
	;
	v2867 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+96))
	v2868 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+104))
	v2869 = int64(-1)
	v2870 = int64(9223090561878065151)
	F___multf3(m, v2167+int32(80), v2867, v2868, v2869, v2870)
	mBase = m.M
	v2874 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+80))
	v2875 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+88))
	F___multf3(m, v2167-int32(-64), v2874, v2875, v2869, v2870)
	mBase = m.M
	v2879 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+64))
	v2880 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+72))
	v5270 = v2879
	v5277 = v2880
	goto L490
L584:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2820))) = v2859
	*(*int64)(unsafe.Add(mBase, uint32(v2820)+8)) = v2858
	m.G0 = v2826 + int32(16)
	goto L583
L585:
	;
	v2858 = int64(0)
	v2859 = int64(0)
	goto L584
L586:
	;
	goto L587
L587:
	;
	v2832 = v106 >> (uint(int32(31)) % 32)
	v2834 = v106 ^ v2832 - v2832
	v2837 = base.I32_clz(v2834)
	F___ashlti3(m, v2826, base.I64_extend_i32_u(v2834), int64(0), v2837+int32(81))
	mBase = m.M
	v2841 = *(*int64)(unsafe.Add(mBase, uint32(v2826)+8))
	v2856 = *(*int64)(unsafe.Add(mBase, uint32(v2826)))
	v2858 = v2841 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2837)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v2859 = v2856
	goto L584
L588:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	v2889 = v2167 + int32(144)
	v2893 = m.G0
	v2895 = v2893 - int32(16)
	m.G0 = v2895
	if v106 == int32(0) {
		goto L593
	} else {
		goto L594
	}
L589:
	;
	goto L590
L590:
	;
	if v2611 != 0 {
		goto L596
	} else {
		goto L597
	}
L591:
	;
	v2936 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+144))
	v2937 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+152))
	v2938 = int64(0)
	v2939 = int64(281474976710656)
	F___multf3(m, v2167+int32(128), v2936, v2937, v2938, v2939)
	mBase = m.M
	v2943 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+128))
	v2944 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+136))
	F___multf3(m, v2167+int32(112), v2943, v2944, v2938, v2939)
	mBase = m.M
	v2948 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+112))
	v2949 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+120))
	v5270 = v2948
	v5277 = v2949
	goto L490
L592:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2889))) = v2928
	*(*int64)(unsafe.Add(mBase, uint32(v2889)+8)) = v2927
	m.G0 = v2895 + int32(16)
	goto L591
L593:
	;
	v2927 = int64(0)
	v2928 = int64(0)
	goto L592
L594:
	;
	goto L595
L595:
	;
	v2901 = v106 >> (uint(int32(31)) % 32)
	v2903 = v106 ^ v2901 - v2901
	v2906 = base.I32_clz(v2903)
	F___ashlti3(m, v2895, base.I64_extend_i32_u(v2903), int64(0), v2906+int32(81))
	mBase = m.M
	v2910 = *(*int64)(unsafe.Add(mBase, uint32(v2895)+8))
	v2925 = *(*int64)(unsafe.Add(mBase, uint32(v2895)))
	v2927 = v2910 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v2906)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v2928 = v2925
	goto L592
L596:
	;
	if v2611 <= int32(8) {
		goto L599
	} else {
		goto L600
	}
L597:
	;
	v3028 = v2608
	goto L598
L598:
	;
	v3051 = base.I32_wrap_i64(v2623)
	if int32(9) <= v2613 {
		goto L605
	} else {
		goto L606
	}
L599:
	;
	v2956 = v2167 + int32(784) + v2608<<(uint(int32(2))%32)
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2956)))
	v2964 = v2957
	v2966 = v2611
	goto L602
L600:
	;
	goto L601
L601:
	;
	v3028 = v2608 + int32(1)
	goto L598
L602:
	;
	v2987 = v2964 * int32(10)
	v2989 = v2966 + int32(1)
	if v2989 != int32(9) {
		v2964 = v2987
		v2966 = v2989
		goto L602
	} else {
		goto L604
	}
L603:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2956))) = v2987
	goto L601
L604:
	;
	goto L603
L605:
	;
	v3475 = v3028
	goto L659
L606:
	;
	if int64(17) < v2623 {
		goto L605
	} else {
		goto L607
	}
L607:
	;
	if v3051 < v2613 {
		goto L605
	} else {
		goto L608
	}
L608:
	;
	if v2623 == int64(9) {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v3060 = v2167 + int32(192)
	v3064 = m.G0
	v3066 = v3064 - int32(16)
	m.G0 = v3066
	if v106 == int32(0) {
		goto L614
	} else {
		goto L615
	}
L610:
	;
	goto L611
L611:
	;
	if v2623 <= int64(8) {
		goto L622
	} else {
		goto L623
	}
L612:
	;
	v3106 = v2167 + int32(176)
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+784))
	v3110 = m.G0
	v3112 = v3110 - int32(16)
	m.G0 = v3112
	if v3107 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L613:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3060))) = v3099
	*(*int64)(unsafe.Add(mBase, uint32(v3060)+8)) = v3098
	m.G0 = v3066 + int32(16)
	goto L612
L614:
	;
	v3098 = int64(0)
	v3099 = int64(0)
	goto L613
L615:
	;
	goto L616
L616:
	;
	v3072 = v106 >> (uint(int32(31)) % 32)
	v3074 = v106 ^ v3072 - v3072
	v3077 = base.I32_clz(v3074)
	F___ashlti3(m, v3066, base.I64_extend_i32_u(v3074), int64(0), v3077+int32(81))
	mBase = m.M
	v3081 = *(*int64)(unsafe.Add(mBase, uint32(v3066)+8))
	v3096 = *(*int64)(unsafe.Add(mBase, uint32(v3066)))
	v3098 = v3081 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3077)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v3099 = v3096
	goto L613
L617:
	;
	v3145 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+192))
	v3146 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+200))
	v3147 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+176))
	v3148 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+184))
	F___multf3(m, v2167+int32(160), v3145, v3146, v3147, v3148)
	mBase = m.M
	v3150 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+160))
	v3151 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+168))
	v5270 = v3150
	v5277 = v3151
	goto L490
L618:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3106))) = v3137
	*(*int64)(unsafe.Add(mBase, uint32(v3106)+8)) = v3136
	m.G0 = v3112 + int32(16)
	goto L617
L619:
	;
	v3136 = int64(0)
	v3137 = int64(0)
	goto L618
L620:
	;
	goto L621
L621:
	;
	v3120 = base.I32_clz(v3107)
	F___ashlti3(m, v3112, base.I64_extend_i32_u(v3107), int64(0), int32(112)-(v3120^int32(31)))
	mBase = m.M
	v3125 = *(*int64)(unsafe.Add(mBase, uint32(v3112)+8))
	v3134 = *(*int64)(unsafe.Add(mBase, uint32(v3112)))
	v3136 = v3125 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3120)<<(uint(int64(48))%64)
	v3137 = v3134
	goto L618
L622:
	;
	v3155 = v2167 + int32(272)
	v3159 = m.G0
	v3161 = v3159 - int32(16)
	m.G0 = v3161
	if v106 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L623:
	;
	goto L624
L624:
	;
	v3311 = v42 + v3051*int32(-3) + int32(27)
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+784))
	if int32(base.Ui32(v3315)>>(uint(v3311)%32)) != 0 {
		goto L640
	} else {
		goto L641
	}
L625:
	;
	v3201 = v2167 + int32(256)
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(v2167)+784))
	v3205 = m.G0
	v3207 = v3205 - int32(16)
	m.G0 = v3207
	if v3202 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L626:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3155))) = v3194
	*(*int64)(unsafe.Add(mBase, uint32(v3155)+8)) = v3193
	m.G0 = v3161 + int32(16)
	goto L625
L627:
	;
	v3193 = int64(0)
	v3194 = int64(0)
	goto L626
L628:
	;
	goto L629
L629:
	;
	v3167 = v106 >> (uint(int32(31)) % 32)
	v3169 = v106 ^ v3167 - v3167
	v3172 = base.I32_clz(v3169)
	F___ashlti3(m, v3161, base.I64_extend_i32_u(v3169), int64(0), v3172+int32(81))
	mBase = m.M
	v3176 = *(*int64)(unsafe.Add(mBase, uint32(v3161)+8))
	v3191 = *(*int64)(unsafe.Add(mBase, uint32(v3161)))
	v3193 = v3176 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3172)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v3194 = v3191
	goto L626
L630:
	;
	v3240 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+272))
	v3241 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+280))
	v3242 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+256))
	v3243 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+264))
	F___multf3(m, v2167+int32(240), v3240, v3241, v3242, v3243)
	mBase = m.M
	v3246 = v2167 + int32(224)
	v3247 = int32(0)
	v3253 = *(*int32)(unsafe.Add(mBase, uint32((v3247-v3051)<<(uint(int32(2))%32))+uint32(_consts[1335])))
	v3257 = m.G0
	v3259 = v3257 - int32(16)
	m.G0 = v3259
	if v3253 == v3247 {
		goto L637
	} else {
		goto L638
	}
L631:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3201))) = v3232
	*(*int64)(unsafe.Add(mBase, uint32(v3201)+8)) = v3231
	m.G0 = v3207 + int32(16)
	goto L630
L632:
	;
	v3231 = int64(0)
	v3232 = int64(0)
	goto L631
L633:
	;
	goto L634
L634:
	;
	v3215 = base.I32_clz(v3202)
	F___ashlti3(m, v3207, base.I64_extend_i32_u(v3202), int64(0), int32(112)-(v3215^int32(31)))
	mBase = m.M
	v3220 = *(*int64)(unsafe.Add(mBase, uint32(v3207)+8))
	v3229 = *(*int64)(unsafe.Add(mBase, uint32(v3207)))
	v3231 = v3220 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3215)<<(uint(int64(48))%64)
	v3232 = v3229
	goto L631
L635:
	;
	v3300 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+240))
	v3301 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+248))
	v3302 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+224))
	v3303 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+232))
	F___divtf3(m, v2167+int32(208), v3300, v3301, v3302, v3303)
	mBase = m.M
	v3305 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+208))
	v3306 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+216))
	v5270 = v3305
	v5277 = v3306
	goto L490
L636:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3246))) = v3292
	*(*int64)(unsafe.Add(mBase, uint32(v3246)+8)) = v3291
	m.G0 = v3259 + int32(16)
	goto L635
L637:
	;
	v3291 = int64(0)
	v3292 = int64(0)
	goto L636
L638:
	;
	goto L639
L639:
	;
	v3265 = v3253 >> (uint(int32(31)) % 32)
	v3267 = v3253 ^ v3265 - v3265
	v3270 = base.I32_clz(v3267)
	F___ashlti3(m, v3259, base.I64_extend_i32_u(v3267), int64(0), v3270+int32(81))
	mBase = m.M
	v3274 = *(*int64)(unsafe.Add(mBase, uint32(v3259)+8))
	v3289 = *(*int64)(unsafe.Add(mBase, uint32(v3259)))
	v3291 = v3274 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3270)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v3253&int32(-2147483648))<<(uint(int64(32))%64)
	v3292 = v3289
	goto L636
L640:
	;
	v3317 = base.B2i32(v3311 <= int32(30))
	goto L642
L641:
	;
	v3317 = int32(0)
	goto L642
L642:
	;
	if v3317 != 0 {
		goto L605
	} else {
		goto L643
	}
L643:
	;
	v3319 = v2167 + int32(352)
	v3323 = m.G0
	v3325 = v3323 - int32(16)
	m.G0 = v3325
	if v106 == int32(0) {
		goto L646
	} else {
		goto L647
	}
L644:
	;
	v3365 = v2167 + int32(336)
	v3368 = m.G0
	v3370 = v3368 - int32(16)
	m.G0 = v3370
	if v3315 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L645:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3319))) = v3358
	*(*int64)(unsafe.Add(mBase, uint32(v3319)+8)) = v3357
	m.G0 = v3325 + int32(16)
	goto L644
L646:
	;
	v3357 = int64(0)
	v3358 = int64(0)
	goto L645
L647:
	;
	goto L648
L648:
	;
	v3331 = v106 >> (uint(int32(31)) % 32)
	v3333 = v106 ^ v3331 - v3331
	v3336 = base.I32_clz(v3333)
	F___ashlti3(m, v3325, base.I64_extend_i32_u(v3333), int64(0), v3336+int32(81))
	mBase = m.M
	v3340 = *(*int64)(unsafe.Add(mBase, uint32(v3325)+8))
	v3355 = *(*int64)(unsafe.Add(mBase, uint32(v3325)))
	v3357 = v3340 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3336)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v3358 = v3355
	goto L645
L649:
	;
	v3403 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+352))
	v3404 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+360))
	v3405 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+336))
	v3406 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+344))
	F___multf3(m, v2167+int32(320), v3403, v3404, v3405, v3406)
	mBase = m.M
	v3409 = v2167 + int32(304)
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v3051<<(uint(int32(2))%32))+uint32(_consts[1339])))
	v3418 = m.G0
	v3420 = v3418 - int32(16)
	m.G0 = v3420
	if v3414 == int32(0) {
		goto L656
	} else {
		goto L657
	}
L650:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3365))) = v3395
	*(*int64)(unsafe.Add(mBase, uint32(v3365)+8)) = v3394
	m.G0 = v3370 + int32(16)
	goto L649
L651:
	;
	v3394 = int64(0)
	v3395 = int64(0)
	goto L650
L652:
	;
	goto L653
L653:
	;
	v3378 = base.I32_clz(v3315)
	F___ashlti3(m, v3370, base.I64_extend_i32_u(v3315), int64(0), int32(112)-(v3378^int32(31)))
	mBase = m.M
	v3383 = *(*int64)(unsafe.Add(mBase, uint32(v3370)+8))
	v3392 = *(*int64)(unsafe.Add(mBase, uint32(v3370)))
	v3394 = v3383 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3378)<<(uint(int64(48))%64)
	v3395 = v3392
	goto L650
L654:
	;
	v3461 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+320))
	v3462 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+328))
	v3463 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+304))
	v3464 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+312))
	F___multf3(m, v2167+int32(288), v3461, v3462, v3463, v3464)
	mBase = m.M
	v3466 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+288))
	v3467 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+296))
	v5270 = v3466
	v5277 = v3467
	goto L490
L655:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3409))) = v3453
	*(*int64)(unsafe.Add(mBase, uint32(v3409)+8)) = v3452
	m.G0 = v3420 + int32(16)
	goto L654
L656:
	;
	v3452 = int64(0)
	v3453 = int64(0)
	goto L655
L657:
	;
	goto L658
L658:
	;
	v3426 = v3414 >> (uint(int32(31)) % 32)
	v3428 = v3414 ^ v3426 - v3426
	v3431 = base.I32_clz(v3428)
	F___ashlti3(m, v3420, base.I64_extend_i32_u(v3428), int64(0), v3431+int32(81))
	mBase = m.M
	v3435 = *(*int64)(unsafe.Add(mBase, uint32(v3420)+8))
	v3450 = *(*int64)(unsafe.Add(mBase, uint32(v3420)))
	v3452 = v3435 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v3431)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v3414&int32(-2147483648))<<(uint(int64(32))%64)
	v3453 = v3450
	goto L655
L659:
	;
	v3501 = v3475 - int32(1)
	v3505 = *(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3501<<(uint(int32(2))%32))))
	if v3505 == int32(0) {
		v3475 = v3501
		goto L659
	} else {
		goto L661
	}
L660:
	;
	v3508 = int32(0)
	v3510 = base.I32_rem_s(v3051, int32(9))
	if v3510 == v3508 {
		goto L663
	} else {
		goto L664
	}
L661:
	;
	goto L660
L662:
	;
	v3660 = v3632
	v3662 = v3634
	v3667 = v3639
	v3669 = v3508
	goto L683
L663:
	;
	v3632 = v3475
	v3634 = int32(0)
	v3639 = v3051
	goto L662
L664:
	;
	goto L665
L665:
	;
	if v2623 < int64(0) {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v3518 = v3510 + int32(9)
	goto L668
L667:
	;
	v3518 = v3510
	goto L668
L668:
	;
	if v3475 == int32(0) {
		goto L670
	} else {
		goto L671
	}
L669:
	;
	v3632 = v3601
	v3634 = v3603
	v3639 = v3608 - v3518 + int32(9)
	goto L662
L670:
	;
	v3521 = int32(0)
	v3601 = v3521
	v3603 = v3521
	v3608 = v3051
	goto L669
L671:
	;
	goto L672
L672:
	;
	v3524 = int32(0)
	v3530 = *(*int32)(unsafe.Add(mBase, uint32((v3524-v3518)<<(uint(int32(2))%32))+uint32(_consts[1335])))
	v3531 = base.I32_div_s(int32(1000000000), v3530)
	v3537 = v3524
	v3538 = v3524
	v3541 = v3524
	v3543 = v3051
	goto L673
L673:
	;
	v3567 = v2167 + int32(784) + v3541<<(uint(int32(2))%32)
	v3568 = *(*int32)(unsafe.Add(mBase, uint32(v3567)))
	v3569 = base.I32_div_u_s(v3568, v3530)
	v3570 = v3569 + v3537
	*(*int32)(unsafe.Add(mBase, uint32(v3567))) = v3570
	v3579 = base.B2i32(v3570 == int32(0)) & base.B2i32(v3538 == v3541)
	if v3579 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	if v3586 == int32(0) {
		v3601 = v3475
		v3603 = v3580
		v3608 = v3583
		goto L669
	} else {
		goto L682
	}
L675:
	;
	v3580 = (v3538 + int32(1)) & int32(2047)
	goto L677
L676:
	;
	v3580 = v3538
	goto L677
L677:
	;
	if v3579 != 0 {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	v3583 = v3543 - int32(9)
	goto L680
L679:
	;
	v3583 = v3543
	goto L680
L680:
	;
	v3586 = v3531 * (v3568 - v3530*v3569)
	v3588 = v3541 + int32(1)
	if v3588 != v3475 {
		v3537 = v3586
		v3538 = v3580
		v3541 = v3588
		v3543 = v3583
		goto L673
	} else {
		goto L681
	}
L681:
	;
	goto L674
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3475<<(uint(int32(2))%32)))) = v3586
	v3601 = v3475 + int32(1)
	v3603 = v3580
	v3608 = v3583
	goto L669
L683:
	;
	v3695 = v3660
	v3704 = v3669
	goto L686
L684:
	;
	v3837 = v3695
	v3839 = v3662
	v3844 = v3667
	v3846 = v3704
	goto L713
L685:
	;
	goto L684
L686:
	;
	if base.B2i32(v3667 < int32(36)) == int32(0) {
		goto L688
	} else {
		goto L689
	}
L687:
	;
	v3806 = (v3662 - int32(1)) & int32(2047)
	if v3806 == v3733 {
		goto L709
	} else {
		goto L710
	}
L688:
	;
	if v3667 != int32(36) {
		goto L685
	} else {
		goto L691
	}
L689:
	;
	goto L690
L690:
	;
	v3733 = v3695
	v3737 = v3695 + int32(2047)
	v3739 = int32(0)
	goto L693
L691:
	;
	v3726 = *(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3662<<(uint(int32(2))%32))))
	if base.Ui32(int32(10384593)) <= base.Ui32(v3726) {
		goto L685
	} else {
		goto L692
	}
L692:
	;
	goto L690
L693:
	;
	v3764 = v3737 & int32(2047)
	v3767 = v2167 + int32(784) + v3764<<(uint(int32(2))%32)
	v3768 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v3767))))
	v3771 = base.I64_extend_i32_u(v3739) + v3768<<(uint(int64(29))%64)
	if base.Ui64(v3771) < base.Ui64(int64(1000000001)) {
		goto L695
	} else {
		goto L696
	}
L694:
	;
	v3800 = v3704 - int32(29)
	if v3783 == int32(0) {
		v3695 = v3733
		v3704 = v3800
		goto L686
	} else {
		goto L708
	}
L695:
	;
	v3781 = v3771
	v3783 = int32(0)
	goto L697
L696:
	;
	v3775 = int64(1000000000)
	v3776 = base.I64_div_u_s(v3771, v3775)
	v3781 = v3771 - v3776*v3775
	v3783 = base.I32_wrap_i64(v3776)
	goto L697
L697:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v3767))) = uint32(v3781)
	if v3781 == int64(0) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v3787 = v3764
	goto L700
L699:
	;
	v3787 = v3733
	goto L700
L700:
	;
	if v3662 == v3764 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	v3789 = v3733
	goto L703
L702:
	;
	v3789 = v3787
	goto L703
L703:
	;
	v3793 = (v3733 - int32(1)) & int32(2047)
	if v3764 != v3793 {
		goto L704
	} else {
		goto L705
	}
L704:
	;
	v3795 = v3733
	goto L706
L705:
	;
	v3795 = v3789
	goto L706
L706:
	;
	if v3662 != v3764 {
		v3733 = v3795
		v3737 = v3764 - int32(1)
		v3739 = v3783
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
	v3809 = v2167 + int32(784)
	v3814 = int32(2)
	v3816 = v3809 + (v3733+int32(2046))&int32(2047)<<(uint(v3814)%32)
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(v3816)))
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3809+v3793<<(uint(v3814)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3816))) = v3817 | v3823
	v3826 = v3793
	goto L711
L710:
	;
	v3826 = v3733
	goto L711
L711:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3806<<(uint(int32(2))%32)))) = v3783
	v3660 = v3826
	v3662 = v3806
	v3667 = v3667 + int32(9)
	v3669 = v3800
	goto L683
L712:
	;
	v4548 = (v3913 + int32(4)) & int32(2047)
	if v4548 == v4037 {
		v4989 = v4539
		v4990 = v4541
		goto L835
	} else {
		goto L836
	}
L713:
	;
	v3864 = int32(1)
	v3866 = int32(2047)
	v3867 = (v3837 + v3864) & v3866
	v3876 = v2167 + int32(784) + (v3837-v3864)&v3866<<(uint(int32(2))%32)
	v3880 = v3839
	v3885 = v3844
	v3887 = v3846
	goto L715
L714:
	;
	v4246 = v2167 + int32(656)
	v4247 = float64(1)
	v4249 = int32(225) - v4161
	if int32(1024) <= v4249 {
		goto L771
	} else {
		goto L772
	}
L715:
	;
	if int32(45) < v3885 {
		goto L717
	} else {
		goto L718
	}
L716:
	;
	goto L714
L717:
	;
	v3909 = int32(9)
	goto L719
L718:
	;
	v3909 = int32(1)
	goto L719
L719:
	;
	v3913 = v3880
	v3920 = v3887
	goto L721
L720:
	;
	goto L716
L721:
	;
	v3945 = int32(0)
	goto L724
L722:
	;
	v4171 = int32(-1)
	v4179 = v3913
	v4182 = int32(0)
	v4183 = v3913
	v4184 = v3885
	goto L756
L723:
	;
	v4167 = v3909 + v3920
	if v3837 == v3913 {
		v3913 = v3837
		v3920 = v4167
		goto L721
	} else {
		goto L755
	}
L724:
	;
	v3969 = (v3945 + v3913) & int32(2047)
	if v3969 == v3837 {
		goto L726
	} else {
		goto L727
	}
L725:
	;
	if v3885 != int32(36) {
		goto L723
	} else {
		goto L731
	}
L726:
	;
	goto L725
L727:
	;
	v3973 = int32(2)
	v3976 = *(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3969<<(uint(v3973)%32))))
	v3981 = *(*int32)(unsafe.Add(mBase, uint32(v3945<<(uint(v3973)%32))+uint32(_consts[1340])))
	if base.Ui32(v3976) < base.Ui32(v3981) {
		goto L726
	} else {
		goto L728
	}
L728:
	;
	if base.Ui32(v3981) < base.Ui32(v3976) {
		goto L723
	} else {
		goto L729
	}
L729:
	;
	v3985 = v3945 + int32(1)
	if v3985 != int32(4) {
		v3945 = v3985
		goto L724
	} else {
		goto L730
	}
L730:
	;
	goto L726
L731:
	;
	v3993 = int64(0)
	v3997 = v3837
	v4002 = int32(0)
	v4016 = v3993
	v4017 = v3993
	goto L732
L732:
	;
	v4026 = (v4002 + v3913) & int32(2047)
	if v3997 == v4026 {
		goto L734
	} else {
		goto L735
	}
L733:
	;
	v4100 = v2167 + int32(720)
	v4104 = m.G0
	v4106 = v4104 - int32(16)
	m.G0 = v4106
	if v106 == int32(0) {
		goto L745
	} else {
		goto L746
	}
L734:
	;
	v4031 = (v3997 + int32(1)) & int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v4031<<(uint(int32(2))%32)+v2167)+780)) = int32(0)
	v4037 = v4031
	goto L736
L735:
	;
	v4037 = v3997
	goto L736
L736:
	;
	v4039 = v2167 + int32(768)
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v4026<<(uint(int32(2))%32))))
	v4048 = m.G0
	v4050 = v4048 - int32(16)
	m.G0 = v4050
	if v4045 == int32(0) {
		goto L739
	} else {
		goto L740
	}
L737:
	;
	F___multf3(m, v2167+int32(752), v4016, v4017, int64(0), int64(4619810130798575616))
	mBase = m.M
	v4088 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+752))
	v4089 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+760))
	v4090 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+768))
	v4091 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+776))
	F___addtf3(m, v2167+int32(736), v4088, v4089, v4090, v4091)
	mBase = m.M
	v4093 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+744))
	v4094 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+736))
	v4096 = v4002 + int32(1)
	if v4096 != int32(4) {
		v3997 = v4037
		v4002 = v4096
		v4016 = v4094
		v4017 = v4093
		goto L732
	} else {
		goto L742
	}
L738:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4039))) = v4075
	*(*int64)(unsafe.Add(mBase, uint32(v4039)+8)) = v4074
	m.G0 = v4050 + int32(16)
	goto L737
L739:
	;
	v4074 = int64(0)
	v4075 = int64(0)
	goto L738
L740:
	;
	goto L741
L741:
	;
	v4058 = base.I32_clz(v4045)
	F___ashlti3(m, v4050, base.I64_extend_i32_u(v4045), int64(0), int32(112)-(v4058^int32(31)))
	mBase = m.M
	v4063 = *(*int64)(unsafe.Add(mBase, uint32(v4050)+8))
	v4072 = *(*int64)(unsafe.Add(mBase, uint32(v4050)))
	v4074 = v4063 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v4058)<<(uint(int64(48))%64)
	v4075 = v4072
	goto L738
L742:
	;
	goto L733
L743:
	;
	v4147 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+720))
	v4148 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+728))
	F___multf3(m, v2167+int32(704), v4094, v4093, v4147, v4148)
	mBase = m.M
	v4150 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+712))
	v4152 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+704))
	v4154 = v3920 + int32(113)
	v4155 = v4154 - v39
	v4156 = int32(0)
	if v4156 < v4155 {
		goto L748
	} else {
		goto L749
	}
L744:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4100))) = v4139
	*(*int64)(unsafe.Add(mBase, uint32(v4100)+8)) = v4138
	m.G0 = v4106 + int32(16)
	goto L743
L745:
	;
	v4138 = int64(0)
	v4139 = int64(0)
	goto L744
L746:
	;
	goto L747
L747:
	;
	v4112 = v106 >> (uint(int32(31)) % 32)
	v4114 = v106 ^ v4112 - v4112
	v4117 = base.I32_clz(v4114)
	F___ashlti3(m, v4106, base.I64_extend_i32_u(v4114), int64(0), v4117+int32(81))
	mBase = m.M
	v4121 = *(*int64)(unsafe.Add(mBase, uint32(v4106)+8))
	v4136 = *(*int64)(unsafe.Add(mBase, uint32(v4106)))
	v4138 = v4121 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(16414)-v4117)<<(uint(int64(48))%64) | base.I64_extend_i32_u(v106&int32(-2147483648))<<(uint(int64(32))%64)
	v4139 = v4136
	goto L744
L748:
	;
	v4159 = v4155
	goto L750
L749:
	;
	v4159 = v4156
	goto L750
L750:
	;
	v4160 = base.B2i32(v4155 < v42)
	if v4155 < v42 {
		goto L751
	} else {
		goto L752
	}
L751:
	;
	v4161 = v4159
	goto L753
L752:
	;
	v4161 = v42
	goto L753
L753:
	;
	if base.Ui32(v4161) <= base.Ui32(int32(112)) {
		goto L720
	} else {
		goto L754
	}
L754:
	;
	v4539 = int64(0)
	v4540 = v4150
	v4541 = v21
	v4542 = v4152
	v4543 = v21
	v4544 = v21
	goto L712
L755:
	;
	goto L722
L756:
	;
	v4208 = v2167 + int32(784) + v4183<<(uint(int32(2))%32)
	v4209 = *(*int32)(unsafe.Add(mBase, uint32(v4208)))
	v4211 = int32(base.Ui32(v4209)>>(uint(v3909)%32)) + v4182
	*(*int32)(unsafe.Add(mBase, uint32(v4208))) = v4211
	v4220 = base.B2i32(v4211 == int32(0)) & base.B2i32(v4179 == v4183)
	if v4220 != 0 {
		goto L758
	} else {
		goto L759
	}
L757:
	;
	if v4226 == int32(0) {
		v3880 = v4221
		v3885 = v4224
		v3887 = v4167
		goto L715
	} else {
		goto L765
	}
L758:
	;
	v4221 = (v4179 + int32(1)) & int32(2047)
	goto L760
L759:
	;
	v4221 = v4179
	goto L760
L760:
	;
	if v4220 != 0 {
		goto L761
	} else {
		goto L762
	}
L761:
	;
	v4224 = v4184 - int32(9)
	goto L763
L762:
	;
	v4224 = v4184
	goto L763
L763:
	;
	v4226 = v4209 & (v4171<<(uint(v3909)%32) ^ v4171) * int32(base.Ui32(int32(1000000000))>>(uint(v3909)%32))
	v4230 = (v4183 + int32(1)) & int32(2047)
	if v4230 != v3837 {
		v4179 = v4221
		v4182 = v4226
		v4183 = v4230
		v4184 = v4224
		goto L756
	} else {
		goto L764
	}
L764:
	;
	goto L757
L765:
	;
	if v4221 != v3867 {
		goto L766
	} else {
		goto L767
	}
L766:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v3837<<(uint(int32(2))%32)))) = v4226
	v3837 = v3867
	v3839 = v4221
	v3844 = v4224
	v3846 = v4167
	goto L713
L767:
	;
	goto L768
L768:
	;
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v3876)))
	*(*int32)(unsafe.Add(mBase, uint32(v3876))) = v4241 | int32(1)
	v3880 = v4221
	v3885 = v4224
	v3887 = v4167
	goto L715
L769:
	;
	v4291 = int64(0)
	v4297 = m.G0
	v4299 = v4297 - int32(16)
	m.G0 = v4299
	v4301 = base.I64_reinterpret_f64(base.F64_mul(v4282, base.F64_reinterpret_i64(base.I64_extend_i32_u(v4283+int32(1023))<<(uint(int64(52))%64))))
	v4303 = v4301 & int64(4503599627370495)
	v4307 = int64(base.Ui64(v4301)>>(uint(int64(52))%64)) & int64(2047)
	if v4307 != v4291 {
		goto L789
	} else {
		goto L790
	}
L770:
	;
	goto L769
L771:
	;
	v4253 = base.F64_mul(v4247, float64(8.98846567431158e+307))
	if base.Ui32(v4249) < base.Ui32(int32(2047)) {
		goto L774
	} else {
		goto L775
	}
L772:
	;
	goto L773
L773:
	;
	if int32(-1023) < v4249 {
		v4282 = v4247
		v4283 = v4249
		goto L770
	} else {
		goto L780
	}
L774:
	;
	v4282 = v4253
	v4283 = v4249 - int32(1023)
	goto L770
L775:
	;
	goto L776
L776:
	;
	v4260 = int32(3069)
	if base.Ui32(v4260) <= base.Ui32(v4249) {
		goto L777
	} else {
		goto L778
	}
L777:
	;
	v4263 = v4260
	goto L779
L778:
	;
	v4263 = v4249
	goto L779
L779:
	;
	v4282 = base.F64_mul(v4253, float64(8.98846567431158e+307))
	v4283 = v4263 - int32(2046)
	goto L770
L780:
	;
	v4269 = base.F64_mul(v4247, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v4249) {
		goto L781
	} else {
		goto L782
	}
L781:
	;
	v4282 = v4269
	v4283 = v4249 + int32(969)
	goto L770
L782:
	;
	goto L783
L783:
	;
	v4276 = int32(-2960)
	if base.Ui32(v4249) <= base.Ui32(v4276) {
		goto L784
	} else {
		goto L785
	}
L784:
	;
	v4279 = v4276
	goto L786
L785:
	;
	v4279 = v4249
	goto L786
L786:
	;
	v4282 = base.F64_mul(v4269, float64(2.004168360008973e-292))
	v4283 = v4279 + int32(1938)
	goto L770
L787:
	;
	v4365 = v2167 + int32(688)
	v4366 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+656))
	v4367 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+664))
	*(*int64)(unsafe.Add(mBase, uint32(v4365))) = v4366
	v4373 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(v4365)+8)) = v4367&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(v4367&int64(9223090561878065152))>>(uint(v4373)%64)))|base.I32_wrap_i64(int64(base.Ui64(v4150)>>(uint(v4373)%64)))&int32(32768))<<(uint(v4373)%64)
	goto L801
L788:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4246))) = v4352
	*(*int64)(unsafe.Add(mBase, uint32(v4246)+8)) = v4301&int64(-9223372036854775807-1) | v4349<<(uint(int64(48))%64) | v4350
	m.G0 = v4299 + int32(16)
	goto L787
L789:
	;
	if v4307 != int64(2047) {
		goto L792
	} else {
		goto L793
	}
L790:
	;
	goto L791
L791:
	;
	if v4303 == int64(0) {
		goto L795
	} else {
		goto L796
	}
L792:
	;
	v4349 = v4307 + int64(15360)
	v4350 = int64(base.Ui64(v4303) >> (uint(int64(4)) % 64))
	v4352 = v4303 << (uint(int64(60)) % 64)
	goto L788
L793:
	;
	goto L794
L794:
	;
	v4349 = int64(32767)
	v4350 = int64(base.Ui64(v4303) >> (uint(int64(4)) % 64))
	v4352 = v4303 << (uint(int64(60)) % 64)
	goto L788
L795:
	;
	v4325 = int64(0)
	v4349 = v4325
	v4350 = v4291
	v4352 = v4325
	goto L788
L796:
	;
	goto L797
L797:
	;
	if base.Ui64(v4303) < base.Ui64(int64(4294967296)) {
		goto L798
	} else {
		goto L799
	}
L798:
	;
	v4338 = base.I32_clz(base.I32_wrap_i64(v4301)) | int32(32)
	goto L800
L799:
	;
	v4338 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4303) >> (uint(int64(32)) % 64))))
	goto L800
L800:
	;
	F___ashlti3(m, v4299, v4303, int64(0), v4338+int32(49))
	mBase = m.M
	v4345 = *(*int64)(unsafe.Add(mBase, uint32(v4299)+8))
	v4348 = *(*int64)(unsafe.Add(mBase, uint32(v4299)))
	v4349 = base.I64_extend_i32_u(int32(15372) - v4338)
	v4350 = v4345 ^ int64(281474976710656)
	v4352 = v4348
	goto L788
L801:
	;
	v4387 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+696))
	v4388 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+688))
	v4390 = v2167 + int32(640)
	v4391 = float64(1)
	v4393 = int32(113) - v4161
	if int32(1024) <= v4393 {
		goto L804
	} else {
		goto L805
	}
L802:
	;
	v4435 = int64(0)
	v4441 = m.G0
	v4443 = v4441 - int32(16)
	m.G0 = v4443
	v4445 = base.I64_reinterpret_f64(base.F64_mul(v4426, base.F64_reinterpret_i64(base.I64_extend_i32_u(v4427+int32(1023))<<(uint(int64(52))%64))))
	v4447 = v4445 & int64(4503599627370495)
	v4451 = int64(base.Ui64(v4445)>>(uint(int64(52))%64)) & int64(2047)
	if v4451 != v4435 {
		goto L822
	} else {
		goto L823
	}
L803:
	;
	goto L802
L804:
	;
	v4397 = base.F64_mul(v4391, float64(8.98846567431158e+307))
	if base.Ui32(v4393) < base.Ui32(int32(2047)) {
		goto L807
	} else {
		goto L808
	}
L805:
	;
	goto L806
L806:
	;
	if int32(-1023) < v4393 {
		v4426 = v4391
		v4427 = v4393
		goto L803
	} else {
		goto L813
	}
L807:
	;
	v4426 = v4397
	v4427 = v4393 - int32(1023)
	goto L803
L808:
	;
	goto L809
L809:
	;
	v4404 = int32(3069)
	if base.Ui32(v4404) <= base.Ui32(v4393) {
		goto L810
	} else {
		goto L811
	}
L810:
	;
	v4407 = v4404
	goto L812
L811:
	;
	v4407 = v4393
	goto L812
L812:
	;
	v4426 = base.F64_mul(v4397, float64(8.98846567431158e+307))
	v4427 = v4407 - int32(2046)
	goto L803
L813:
	;
	v4413 = base.F64_mul(v4391, float64(2.004168360008973e-292))
	if base.Ui32(int32(-1992)) < base.Ui32(v4393) {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v4426 = v4413
	v4427 = v4393 + int32(969)
	goto L803
L815:
	;
	goto L816
L816:
	;
	v4420 = int32(-2960)
	if base.Ui32(v4393) <= base.Ui32(v4420) {
		goto L817
	} else {
		goto L818
	}
L817:
	;
	v4423 = v4420
	goto L819
L818:
	;
	v4423 = v4393
	goto L819
L819:
	;
	v4426 = base.F64_mul(v4413, float64(2.004168360008973e-292))
	v4427 = v4423 + int32(1938)
	goto L803
L820:
	;
	v4510 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+640))
	v4511 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+648))
	F_fmodl(m, v2167+int32(672), v4152, v4150, v4510, v4511)
	mBase = m.M
	v4514 = v2167 + int32(624)
	v4515 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+672))
	v4516 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+680))
	v4518 = m.G0
	v4519 = int32(16)
	v4520 = v4518 - v4519
	m.G0 = v4520
	F___addtf3(m, v4520, v4152, v4150, v4515, v4516^int64(-9223372036854775807-1))
	mBase = m.M
	v4525 = *(*int64)(unsafe.Add(mBase, uint32(v4520)))
	v4526 = *(*int64)(unsafe.Add(mBase, uint32(v4520)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4514)+8)) = v4526
	*(*int64)(unsafe.Add(mBase, uint32(v4514))) = v4525
	m.G0 = v4520 + v4519
	goto L834
L821:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4390))) = v4496
	*(*int64)(unsafe.Add(mBase, uint32(v4390)+8)) = v4445&int64(-9223372036854775807-1) | v4493<<(uint(int64(48))%64) | v4494
	m.G0 = v4443 + int32(16)
	goto L820
L822:
	;
	if v4451 != int64(2047) {
		goto L825
	} else {
		goto L826
	}
L823:
	;
	goto L824
L824:
	;
	if v4447 == int64(0) {
		goto L828
	} else {
		goto L829
	}
L825:
	;
	v4493 = v4451 + int64(15360)
	v4494 = int64(base.Ui64(v4447) >> (uint(int64(4)) % 64))
	v4496 = v4447 << (uint(int64(60)) % 64)
	goto L821
L826:
	;
	goto L827
L827:
	;
	v4493 = int64(32767)
	v4494 = int64(base.Ui64(v4447) >> (uint(int64(4)) % 64))
	v4496 = v4447 << (uint(int64(60)) % 64)
	goto L821
L828:
	;
	v4469 = int64(0)
	v4493 = v4469
	v4494 = v4435
	v4496 = v4469
	goto L821
L829:
	;
	goto L830
L830:
	;
	if base.Ui64(v4447) < base.Ui64(int64(4294967296)) {
		goto L831
	} else {
		goto L832
	}
L831:
	;
	v4482 = base.I32_clz(base.I32_wrap_i64(v4445)) | int32(32)
	goto L833
L832:
	;
	v4482 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4447) >> (uint(int64(32)) % 64))))
	goto L833
L833:
	;
	F___ashlti3(m, v4443, v4447, int64(0), v4482+int32(49))
	mBase = m.M
	v4489 = *(*int64)(unsafe.Add(mBase, uint32(v4443)+8))
	v4492 = *(*int64)(unsafe.Add(mBase, uint32(v4443)))
	v4493 = base.I64_extend_i32_u(int32(15372) - v4482)
	v4494 = v4489 ^ int64(281474976710656)
	v4496 = v4492
	goto L821
L834:
	;
	v4534 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+624))
	v4535 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+632))
	F___addtf3(m, v2167+int32(608), v4388, v4387, v4534, v4535)
	mBase = m.M
	v4537 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+616))
	v4538 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+608))
	v4539 = v4515
	v4540 = v4537
	v4541 = v4516
	v4542 = v4538
	v4543 = v4388
	v4544 = v4387
	goto L712
L835:
	;
	F___addtf3(m, v2167+int32(432), v4542, v4540, v4989, v4990)
	mBase = m.M
	v4996 = v2167 + int32(416)
	v4997 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+432))
	v4998 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+440))
	v5000 = m.G0
	v5001 = int32(16)
	v5002 = v5000 - v5001
	m.G0 = v5002
	F___addtf3(m, v5002, v4997, v4998, v4543, v4544^int64(-9223372036854775807-1))
	mBase = m.M
	v5007 = *(*int64)(unsafe.Add(mBase, uint32(v5002)))
	v5008 = *(*int64)(unsafe.Add(mBase, uint32(v5002)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4996)+8)) = v5008
	*(*int64)(unsafe.Add(mBase, uint32(v4996))) = v5007
	m.G0 = v5002 + v5001
	goto L937
L836:
	;
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v2167+int32(784)+v4548<<(uint(int32(2))%32))))
	if base.Ui32(v4555) <= base.Ui32(int32(499999999)) {
		goto L838
	} else {
		goto L839
	}
L837:
	;
	if base.Ui32(int32(111)) < base.Ui32(v4161) {
		v4989 = v4911
		v4990 = v4912
		goto L835
	} else {
		goto L907
	}
L838:
	;
	if v4555 == int32(0) {
		goto L841
	} else {
		goto L842
	}
L839:
	;
	goto L840
L840:
	;
	if v4555 != int32(500000000) {
		goto L859
	} else {
		goto L860
	}
L841:
	;
	if (v3913+int32(5))&int32(2047) == v4037 {
		v4911 = v4539
		v4912 = v4541
		goto L837
	} else {
		goto L844
	}
L842:
	;
	goto L843
L843:
	;
	v4566 = v2167 + int32(496)
	v4570 = int64(0)
	v4576 = m.G0
	v4578 = v4576 - int32(16)
	m.G0 = v4578
	v4580 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(v106), float64(0.25)))
	v4582 = v4580 & int64(4503599627370495)
	v4586 = int64(base.Ui64(v4580)>>(uint(int64(52))%64)) & int64(2047)
	if v4586 != v4570 {
		goto L847
	} else {
		goto L848
	}
L844:
	;
	goto L843
L845:
	;
	v4645 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+496))
	v4646 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+504))
	F___addtf3(m, v2167+int32(480), v4539, v4541, v4645, v4646)
	mBase = m.M
	v4648 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+488))
	v4649 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+480))
	v4911 = v4649
	v4912 = v4648
	goto L837
L846:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4566))) = v4631
	*(*int64)(unsafe.Add(mBase, uint32(v4566)+8)) = v4580&int64(-9223372036854775807-1) | v4628<<(uint(int64(48))%64) | v4629
	m.G0 = v4578 + int32(16)
	goto L845
L847:
	;
	if v4586 != int64(2047) {
		goto L850
	} else {
		goto L851
	}
L848:
	;
	goto L849
L849:
	;
	if v4582 == int64(0) {
		goto L853
	} else {
		goto L854
	}
L850:
	;
	v4628 = v4586 + int64(15360)
	v4629 = int64(base.Ui64(v4582) >> (uint(int64(4)) % 64))
	v4631 = v4582 << (uint(int64(60)) % 64)
	goto L846
L851:
	;
	goto L852
L852:
	;
	v4628 = int64(32767)
	v4629 = int64(base.Ui64(v4582) >> (uint(int64(4)) % 64))
	v4631 = v4582 << (uint(int64(60)) % 64)
	goto L846
L853:
	;
	v4604 = int64(0)
	v4628 = v4604
	v4629 = v4570
	v4631 = v4604
	goto L846
L854:
	;
	goto L855
L855:
	;
	if base.Ui64(v4582) < base.Ui64(int64(4294967296)) {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	v4617 = base.I32_clz(base.I32_wrap_i64(v4580)) | int32(32)
	goto L858
L857:
	;
	v4617 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4582) >> (uint(int64(32)) % 64))))
	goto L858
L858:
	;
	F___ashlti3(m, v4578, v4582, int64(0), v4617+int32(49))
	mBase = m.M
	v4624 = *(*int64)(unsafe.Add(mBase, uint32(v4578)+8))
	v4627 = *(*int64)(unsafe.Add(mBase, uint32(v4578)))
	v4628 = base.I64_extend_i32_u(int32(15372) - v4617)
	v4629 = v4624 ^ int64(281474976710656)
	v4631 = v4627
	goto L846
L859:
	;
	v4653 = v2167 + int32(592)
	v4657 = int64(0)
	v4663 = m.G0
	v4665 = v4663 - int32(16)
	m.G0 = v4665
	v4667 = base.I64_reinterpret_f64(base.F64_mul(base.F64_convert_i32_s(v106), float64(0.75)))
	v4669 = v4667 & int64(4503599627370495)
	v4673 = int64(base.Ui64(v4667)>>(uint(int64(52))%64)) & int64(2047)
	if v4673 != v4657 {
		goto L864
	} else {
		goto L865
	}
L860:
	;
	goto L861
L861:
	;
	v4737 = base.F64_convert_i32_s(v106)
	if v4037 == (v3913+int32(5))&int32(2047) {
		goto L876
	} else {
		goto L877
	}
L862:
	;
	v4732 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+592))
	v4733 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+600))
	F___addtf3(m, v2167+int32(576), v4539, v4541, v4732, v4733)
	mBase = m.M
	v4735 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+584))
	v4736 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+576))
	v4911 = v4736
	v4912 = v4735
	goto L837
L863:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4653))) = v4718
	*(*int64)(unsafe.Add(mBase, uint32(v4653)+8)) = v4667&int64(-9223372036854775807-1) | v4715<<(uint(int64(48))%64) | v4716
	m.G0 = v4665 + int32(16)
	goto L862
L864:
	;
	if v4673 != int64(2047) {
		goto L867
	} else {
		goto L868
	}
L865:
	;
	goto L866
L866:
	;
	if v4669 == int64(0) {
		goto L870
	} else {
		goto L871
	}
L867:
	;
	v4715 = v4673 + int64(15360)
	v4716 = int64(base.Ui64(v4669) >> (uint(int64(4)) % 64))
	v4718 = v4669 << (uint(int64(60)) % 64)
	goto L863
L868:
	;
	goto L869
L869:
	;
	v4715 = int64(32767)
	v4716 = int64(base.Ui64(v4669) >> (uint(int64(4)) % 64))
	v4718 = v4669 << (uint(int64(60)) % 64)
	goto L863
L870:
	;
	v4691 = int64(0)
	v4715 = v4691
	v4716 = v4657
	v4718 = v4691
	goto L863
L871:
	;
	goto L872
L872:
	;
	if base.Ui64(v4669) < base.Ui64(int64(4294967296)) {
		goto L873
	} else {
		goto L874
	}
L873:
	;
	v4704 = base.I32_clz(base.I32_wrap_i64(v4667)) | int32(32)
	goto L875
L874:
	;
	v4704 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4669) >> (uint(int64(32)) % 64))))
	goto L875
L875:
	;
	F___ashlti3(m, v4665, v4669, int64(0), v4704+int32(49))
	mBase = m.M
	v4711 = *(*int64)(unsafe.Add(mBase, uint32(v4665)+8))
	v4714 = *(*int64)(unsafe.Add(mBase, uint32(v4665)))
	v4715 = base.I64_extend_i32_u(int32(15372) - v4704)
	v4716 = v4711 ^ int64(281474976710656)
	v4718 = v4714
	goto L863
L876:
	;
	v4744 = v2167 + int32(528)
	v4747 = int64(0)
	v4753 = m.G0
	v4755 = v4753 - int32(16)
	m.G0 = v4755
	v4757 = base.I64_reinterpret_f64(base.F64_mul(v4737, float64(0.5)))
	v4759 = v4757 & int64(4503599627370495)
	v4763 = int64(base.Ui64(v4757)>>(uint(int64(52))%64)) & int64(2047)
	if v4763 != v4747 {
		goto L881
	} else {
		goto L882
	}
L877:
	;
	goto L878
L878:
	;
	v4828 = v2167 + int32(560)
	v4831 = int64(0)
	v4837 = m.G0
	v4839 = v4837 - int32(16)
	m.G0 = v4839
	v4841 = base.I64_reinterpret_f64(base.F64_mul(v4737, float64(0.75)))
	v4843 = v4841 & int64(4503599627370495)
	v4847 = int64(base.Ui64(v4841)>>(uint(int64(52))%64)) & int64(2047)
	if v4847 != v4831 {
		goto L895
	} else {
		goto L896
	}
L879:
	;
	v4822 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+528))
	v4823 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+536))
	F___addtf3(m, v2167+int32(512), v4539, v4541, v4822, v4823)
	mBase = m.M
	v4825 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+520))
	v4826 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+512))
	v4911 = v4826
	v4912 = v4825
	goto L837
L880:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4744))) = v4808
	*(*int64)(unsafe.Add(mBase, uint32(v4744)+8)) = v4757&int64(-9223372036854775807-1) | v4805<<(uint(int64(48))%64) | v4806
	m.G0 = v4755 + int32(16)
	goto L879
L881:
	;
	if v4763 != int64(2047) {
		goto L884
	} else {
		goto L885
	}
L882:
	;
	goto L883
L883:
	;
	if v4759 == int64(0) {
		goto L887
	} else {
		goto L888
	}
L884:
	;
	v4805 = v4763 + int64(15360)
	v4806 = int64(base.Ui64(v4759) >> (uint(int64(4)) % 64))
	v4808 = v4759 << (uint(int64(60)) % 64)
	goto L880
L885:
	;
	goto L886
L886:
	;
	v4805 = int64(32767)
	v4806 = int64(base.Ui64(v4759) >> (uint(int64(4)) % 64))
	v4808 = v4759 << (uint(int64(60)) % 64)
	goto L880
L887:
	;
	v4781 = int64(0)
	v4805 = v4781
	v4806 = v4747
	v4808 = v4781
	goto L880
L888:
	;
	goto L889
L889:
	;
	if base.Ui64(v4759) < base.Ui64(int64(4294967296)) {
		goto L890
	} else {
		goto L891
	}
L890:
	;
	v4794 = base.I32_clz(base.I32_wrap_i64(v4757)) | int32(32)
	goto L892
L891:
	;
	v4794 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4759) >> (uint(int64(32)) % 64))))
	goto L892
L892:
	;
	F___ashlti3(m, v4755, v4759, int64(0), v4794+int32(49))
	mBase = m.M
	v4801 = *(*int64)(unsafe.Add(mBase, uint32(v4755)+8))
	v4804 = *(*int64)(unsafe.Add(mBase, uint32(v4755)))
	v4805 = base.I64_extend_i32_u(int32(15372) - v4794)
	v4806 = v4801 ^ int64(281474976710656)
	v4808 = v4804
	goto L880
L893:
	;
	v4906 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+560))
	v4907 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+568))
	F___addtf3(m, v2167+int32(544), v4539, v4541, v4906, v4907)
	mBase = m.M
	v4909 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+552))
	v4910 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+544))
	v4911 = v4910
	v4912 = v4909
	goto L837
L894:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4828))) = v4892
	*(*int64)(unsafe.Add(mBase, uint32(v4828)+8)) = v4841&int64(-9223372036854775807-1) | v4889<<(uint(int64(48))%64) | v4890
	m.G0 = v4839 + int32(16)
	goto L893
L895:
	;
	if v4847 != int64(2047) {
		goto L898
	} else {
		goto L899
	}
L896:
	;
	goto L897
L897:
	;
	if v4843 == int64(0) {
		goto L901
	} else {
		goto L902
	}
L898:
	;
	v4889 = v4847 + int64(15360)
	v4890 = int64(base.Ui64(v4843) >> (uint(int64(4)) % 64))
	v4892 = v4843 << (uint(int64(60)) % 64)
	goto L894
L899:
	;
	goto L900
L900:
	;
	v4889 = int64(32767)
	v4890 = int64(base.Ui64(v4843) >> (uint(int64(4)) % 64))
	v4892 = v4843 << (uint(int64(60)) % 64)
	goto L894
L901:
	;
	v4865 = int64(0)
	v4889 = v4865
	v4890 = v4831
	v4892 = v4865
	goto L894
L902:
	;
	goto L903
L903:
	;
	if base.Ui64(v4843) < base.Ui64(int64(4294967296)) {
		goto L904
	} else {
		goto L905
	}
L904:
	;
	v4878 = base.I32_clz(base.I32_wrap_i64(v4841)) | int32(32)
	goto L906
L905:
	;
	v4878 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v4843) >> (uint(int64(32)) % 64))))
	goto L906
L906:
	;
	F___ashlti3(m, v4839, v4843, int64(0), v4878+int32(49))
	mBase = m.M
	v4885 = *(*int64)(unsafe.Add(mBase, uint32(v4839)+8))
	v4888 = *(*int64)(unsafe.Add(mBase, uint32(v4839)))
	v4889 = base.I64_extend_i32_u(int32(15372) - v4878)
	v4890 = v4885 ^ int64(281474976710656)
	v4892 = v4888
	goto L894
L907:
	;
	v4918 = int64(0)
	F_fmodl(m, v2167+int32(464), v4911, v4912, v4918, int64(4611404543450677248))
	mBase = m.M
	v4921 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+464))
	v4922 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+472))
	v4932 = v4922 & int64(9223372036854775807)
	v4933 = int64(9223090561878065152)
	if v4932 == v4933 {
		goto L910
	} else {
		goto L911
	}
L908:
	;
	if v4980 != 0 {
		v4989 = v4911
		v4990 = v4912
		goto L835
	} else {
		goto L936
	}
L909:
	;
	v4980 = v4976
	goto L908
L910:
	;
	v4937 = base.B2i32(v4921 != v4918)
	goto L912
L911:
	;
	v4937 = base.B2i32(base.Ui64(v4933) < base.Ui64(v4932))
	goto L912
L912:
	;
	if v4937 != 0 {
		v4976 = int32(1)
		goto L909
	} else {
		goto L913
	}
L913:
	;
	goto L915
L915:
	;
	goto L916
L916:
	;
	goto L917
L917:
	;
	if v4921|v4918|(v4932|int64(0)) == int64(0) {
		goto L918
	} else {
		goto L919
	}
L918:
	;
	v4980 = int32(0)
	goto L908
L919:
	;
	goto L920
L920:
	;
	if int64(0) <= v4922&v4918 {
		goto L921
	} else {
		goto L922
	}
L921:
	;
	if v4922 == v4918 {
		goto L924
	} else {
		goto L925
	}
L922:
	;
	goto L923
L923:
	;
	if v4922 == v4918 {
		goto L930
	} else {
		goto L931
	}
L924:
	;
	v4959 = base.B2i32(base.Ui64(v4921) < base.Ui64(v4918))
	goto L926
L925:
	;
	v4959 = base.B2i32(v4922 < v4918)
	goto L926
L926:
	;
	if v4959 != 0 {
		goto L927
	} else {
		goto L928
	}
L927:
	;
	v4980 = int32(-1)
	goto L908
L928:
	;
	goto L929
L929:
	;
	v4980 = base.B2i32(v4921^v4918|(v4922^v4918) != int64(0))
	goto L908
L930:
	;
	v4969 = base.B2i32(base.Ui64(v4918) < base.Ui64(v4921))
	goto L932
L931:
	;
	v4969 = base.B2i32(v4918 < v4922)
	goto L932
L932:
	;
	if v4969 != 0 {
		goto L933
	} else {
		goto L934
	}
L933:
	;
	v4980 = int32(-1)
	goto L908
L934:
	;
	goto L935
L935:
	;
	v4976 = base.B2i32(v4921^v4918|(v4922^v4918) != int64(0))
	goto L909
L936:
	;
	F___addtf3(m, v2167+int32(448), v4911, v4912, int64(0), int64(4611404543450677248))
	mBase = m.M
	v4986 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+456))
	v4987 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+448))
	v4989 = v4987
	v4990 = v4986
	goto L835
L937:
	;
	v5014 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+424))
	v5015 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+416))
	if v4154&int32(2147483647) <= v2171-int32(2) {
		v5168 = v3920
		v5169 = v5014
		v5170 = v5015
		goto L938
	} else {
		goto L939
	}
L938:
	;
	v5172 = v2167 + int32(368)
	v5174 = m.G0
	v5176 = v5174 - int32(80)
	m.G0 = v5176
	if int32(16384) <= v5168 {
		goto L998
	} else {
		goto L999
	}
L939:
	;
	v5022 = v2167 + int32(400)
	v5023 = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(v5022)+8)) = v5014 & v5023
	*(*int64)(unsafe.Add(mBase, uint32(v5022))) = v5015
	v5029 = int64(0)
	F___multf3(m, v2167+int32(384), v5015, v5014, v5029, int64(4611123068473966592))
	mBase = m.M
	v5032 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+400))
	v5033 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+408))
	v5034 = int64(4643211215818981376)
	v5038 = int32(-1)
	v5042 = v5033 & v5023
	v5043 = int64(9223090561878065152)
	if v5042 == v5043 {
		goto L942
	} else {
		goto L943
	}
L940:
	;
	v5084 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+392))
	v5086 = base.B2i32(int32(0) <= v5083)
	if int32(0) <= v5083 {
		goto L958
	} else {
		goto L959
	}
L941:
	;
	v5083 = v5079
	goto L940
L942:
	;
	v5047 = base.B2i32(v5032 != v5029)
	goto L944
L943:
	;
	v5047 = base.B2i32(base.Ui64(v5043) < base.Ui64(v5042))
	goto L944
L944:
	;
	if v5047 != 0 {
		v5079 = v5038
		goto L941
	} else {
		goto L945
	}
L945:
	;
	goto L946
L946:
	;
	if v5032|(v5042|int64(4643211215818981376)) == int64(0) {
		goto L947
	} else {
		goto L948
	}
L947:
	;
	v5083 = int32(0)
	goto L940
L948:
	;
	goto L949
L949:
	;
	if int64(0) <= v5033&v5034 {
		goto L950
	} else {
		goto L951
	}
L950:
	;
	if base.B2i32(v5033 != v5034)&base.B2i32(v5033 < v5034) != 0 {
		v5079 = v5038
		goto L941
	} else {
		goto L953
	}
L951:
	;
	goto L952
L952:
	;
	if v5033 == v5034 {
		goto L954
	} else {
		goto L955
	}
L953:
	;
	v5083 = base.B2i32(v5032|(v5033^v5034) != int64(0))
	goto L940
L954:
	;
	v5074 = base.B2i32(v5032 != int64(0))
	goto L956
L955:
	;
	v5074 = base.B2i32(v5034 < v5033)
	goto L956
L956:
	;
	if v5074 != 0 {
		v5079 = v5038
		goto L941
	} else {
		goto L957
	}
L957:
	;
	v5079 = base.B2i32(v5032|(v5033^v5034) != int64(0))
	goto L941
L958:
	;
	v5087 = v5084
	goto L960
L959:
	;
	v5087 = v5014
	goto L960
L960:
	;
	v5088 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+384))
	if int32(0) <= v5083 {
		goto L961
	} else {
		goto L962
	}
L961:
	;
	v5089 = v5088
	goto L963
L962:
	;
	v5089 = v5015
	goto L963
L963:
	;
	v5090 = int64(0)
	v5099 = v4990 & int64(9223372036854775807)
	v5100 = int64(9223090561878065152)
	if v5099 == v5100 {
		goto L966
	} else {
		goto L967
	}
L964:
	;
	v5148 = v5086 + v3920
	if v5148+int32(110) <= v2171 {
		goto L992
	} else {
		goto L993
	}
L965:
	;
	v5147 = v5143
	goto L964
L966:
	;
	v5104 = base.B2i32(v4989 != v5090)
	goto L968
L967:
	;
	v5104 = base.B2i32(base.Ui64(v5100) < base.Ui64(v5099))
	goto L968
L968:
	;
	if v5104 != 0 {
		v5143 = int32(1)
		goto L965
	} else {
		goto L969
	}
L969:
	;
	goto L971
L971:
	;
	goto L972
L972:
	;
	goto L973
L973:
	;
	if v4989|v5090|(v5099|int64(0)) == int64(0) {
		goto L974
	} else {
		goto L975
	}
L974:
	;
	v5147 = int32(0)
	goto L964
L975:
	;
	goto L976
L976:
	;
	if int64(0) <= v4990&v5090 {
		goto L977
	} else {
		goto L978
	}
L977:
	;
	if v4990 == v5090 {
		goto L980
	} else {
		goto L981
	}
L978:
	;
	goto L979
L979:
	;
	if v4990 == v5090 {
		goto L986
	} else {
		goto L987
	}
L980:
	;
	v5126 = base.B2i32(base.Ui64(v4989) < base.Ui64(v5090))
	goto L982
L981:
	;
	v5126 = base.B2i32(v4990 < v5090)
	goto L982
L982:
	;
	if v5126 != 0 {
		goto L983
	} else {
		goto L984
	}
L983:
	;
	v5147 = int32(-1)
	goto L964
L984:
	;
	goto L985
L985:
	;
	v5147 = base.B2i32(v4989^v5090|(v4990^v5090) != int64(0))
	goto L964
L986:
	;
	v5136 = base.B2i32(base.Ui64(v5090) < base.Ui64(v4989))
	goto L988
L987:
	;
	v5136 = base.B2i32(v5090 < v4990)
	goto L988
L988:
	;
	if v5136 != 0 {
		goto L989
	} else {
		goto L990
	}
L989:
	;
	v5147 = int32(-1)
	goto L964
L990:
	;
	goto L991
L991:
	;
	v5143 = base.B2i32(v4989^v5090|(v4990^v5090) != int64(0))
	goto L965
L992:
	;
	v5153 = int32(0)
	if v4160&(base.B2i32(v4161 != v4155)|base.B2i32(v5083 < v5153))&base.B2i32(v5147 != v5153) == v5153 {
		v5168 = v5148
		v5169 = v5087
		v5170 = v5089
		goto L938
	} else {
		goto L995
	}
L993:
	;
	goto L994
L994:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(68)
	v5168 = v5148
	v5169 = v5087
	v5170 = v5089
	goto L938
L995:
	;
	goto L994
L996:
	;
	v5247 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+368))
	v5248 = *(*int64)(unsafe.Add(mBase, uint32(v2167)+376))
	v5270 = v5247
	v5277 = v5248
	goto L490
L997:
	;
	F___multf3(m, v5176, v5230, v5231, int64(0), base.I64_extend_i32_u(v5232+int32(16383))<<(uint(int64(48))%64))
	mBase = m.M
	v5240 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v5172)+8)) = v5240
	v5242 = *(*int64)(unsafe.Add(mBase, uint32(v5176)))
	*(*int64)(unsafe.Add(mBase, uint32(v5172))) = v5242
	m.G0 = v5176 + int32(80)
	goto L996
L998:
	;
	F___multf3(m, v5176+int32(32), v5170, v5169, int64(0), int64(9222809086901354496))
	mBase = m.M
	v5185 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+40))
	v5186 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+32))
	if base.Ui32(v5168) < base.Ui32(int32(32767)) {
		goto L1001
	} else {
		goto L1002
	}
L999:
	;
	goto L1000
L1000:
	;
	if int32(-16383) < v5168 {
		v5230 = v5170
		v5231 = v5169
		v5232 = v5168
		goto L997
	} else {
		goto L1007
	}
L1001:
	;
	v5230 = v5186
	v5231 = v5185
	v5232 = v5168 - int32(16383)
	goto L997
L1002:
	;
	goto L1003
L1003:
	;
	F___multf3(m, v5176+int32(16), v5186, v5185, int64(0), int64(9222809086901354496))
	mBase = m.M
	v5196 = int32(49149)
	if base.Ui32(v5196) <= base.Ui32(v5168) {
		goto L1004
	} else {
		goto L1005
	}
L1004:
	;
	v5199 = v5196
	goto L1006
L1005:
	;
	v5199 = v5168
	goto L1006
L1006:
	;
	v5202 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+24))
	v5203 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+16))
	v5230 = v5203
	v5231 = v5202
	v5232 = v5199 - int32(32766)
	goto L997
L1007:
	;
	F___multf3(m, v5176-int32(-64), v5170, v5169, int64(0), int64(32088147345014784))
	mBase = m.M
	v5211 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+72))
	v5212 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+64))
	if base.Ui32(int32(-32652)) < base.Ui32(v5168) {
		goto L1008
	} else {
		goto L1009
	}
L1008:
	;
	v5230 = v5212
	v5231 = v5211
	v5232 = v5168 + int32(16269)
	goto L997
L1009:
	;
	goto L1010
L1010:
	;
	F___multf3(m, v5176+int32(48), v5212, v5211, int64(0), int64(32088147345014784))
	mBase = m.M
	v5222 = int32(-48920)
	if base.Ui32(v5168) <= base.Ui32(v5222) {
		goto L1011
	} else {
		goto L1012
	}
L1011:
	;
	v5225 = v5222
	goto L1013
L1012:
	;
	v5225 = v5168
	goto L1013
L1013:
	;
	v5228 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+56))
	v5229 = *(*int64)(unsafe.Add(mBase, uint32(v5176)+48))
	v5230 = v5229
	v5231 = v5228
	v5232 = v5225 + int32(32538)
	goto L997
}
