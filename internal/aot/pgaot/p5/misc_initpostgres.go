package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitPostgres(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int64
	_ = v174
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v213 int32
	_ = v213
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int64
	_ = v302
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v401 int32
	_ = v401
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v653 int32
	_ = v653
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v803 int32
	_ = v803
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
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
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v962 int64
	_ = v962
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1149 int32
	_ = v1149
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1181 int32
	_ = v1181
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1242 int32
	_ = v1242
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1289 int32
	_ = v1289
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1444 int32
	_ = v1444
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1473 int32
	_ = v1473
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1534 int32
	_ = v1534
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1549 int64
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1643 int32
	_ = v1643
	var v1656 int32
	_ = v1656
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1701 int32
	_ = v1701
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1726 int32
	_ = v1726
	var v1739 int32
	_ = v1739
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1780 int32
	_ = v1780
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1838 int64
	_ = v1838
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1986 int32
	_ = v1986
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2001 int32
	_ = v2001
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2074 int32
	_ = v2074
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2135 int32
	_ = v2135
	var v2168 int32
	_ = v2168
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2569 int32
	_ = v2569
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2641 int32
	_ = v2641
	var v2651 int32
	_ = v2651
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2675 int32
	_ = v2675
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2705 int32
	_ = v2705
	var v2711 int32
	_ = v2711
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2721 int32
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2741 int32
	_ = v2741
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2755 int32
	_ = v2755
	var v2758 int32
	_ = v2758
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2769 int32
	_ = v2769
	var v2773 int32
	_ = v2773
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2785 int32
	_ = v2785
	var v2794 int32
	_ = v2794
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2822 int32
	_ = v2822
	var v2828 int32
	_ = v2828
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2838 int32
	_ = v2838
	var v2840 int32
	_ = v2840
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2850 int32
	_ = v2850
	var v2854 int32
	_ = v2854
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2862 int32
	_ = v2862
	var v2872 int32
	_ = v2872
	var v2875 int32
	_ = v2875
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2911 int32
	_ = v2911
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2986 int32
	_ = v2986
	var v2992 int32
	_ = v2992
	var v2997 int32
	_ = v2997
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3005 int32
	_ = v3005
	var v3010 int32
	_ = v3010
	var v3014 int32
	_ = v3014
	var v3016 int32
	_ = v3016
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3046 int64
	_ = v3046
	var v3052 int32
	_ = v3052
	var v3054 int32
	_ = v3054
	var v3057 int32
	_ = v3057
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3071 int32
	_ = v3071
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3077 int64
	_ = v3077
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3102 int32
	_ = v3102
	var v3106 int32
	_ = v3106
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3132 int32
	_ = v3132
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3193 int32
	_ = v3193
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3256 int32
	_ = v3256
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3266 int32
	_ = v3266
	var v3277 int32
	_ = v3277
	var v3280 int32
	_ = v3280
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3290 int32
	_ = v3290
	var v3301 int32
	_ = v3301
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3310 int32
	_ = v3310
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3316 int32
	_ = v3316
	var v3322 int32
	_ = v3322
	var v3332 int32
	_ = v3332
	var v3366 int32
	_ = v3366
	var v3380 int32
	_ = v3380
	var v3415 int32
	_ = v3415
	var v3416 int32
	_ = v3416
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3466 int32
	_ = v3466
	var v3472 int32
	_ = v3472
	var v3484 int32
	_ = v3484
	var v3486 int32
	_ = v3486
	var v3491 int32
	_ = v3491
	var v3500 int32
	_ = v3500
	var v3536 int32
	_ = v3536
	var v3544 int32
	_ = v3544
	var v3545 int32
	_ = v3545
	var v3548 int32
	_ = v3548
	var v3552 int32
	_ = v3552
	var v3563 int32
	_ = v3563
	var v3596 int32
	_ = v3596
	var v3613 int32
	_ = v3613
	var v3648 int32
	_ = v3648
	var v3656 int32
	_ = v3656
	var v3667 int32
	_ = v3667
	var v3668 int32
	_ = v3668
	var v3669 int32
	_ = v3669
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3722 int32
	_ = v3722
	var v3754 int32
	_ = v3754
	var v3767 int32
	_ = v3767
	var v3801 int32
	_ = v3801
	var v3804 int32
	_ = v3804
	var v3805 int32
	_ = v3805
	var v3815 int32
	_ = v3815
	var v3832 int32
	_ = v3832
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3877 int32
	_ = v3877
	var v3903 int32
	_ = v3903
	var v3919 int32
	_ = v3919
	var v3945 int32
	_ = v3945
	var v3955 int32
	_ = v3955
	var v3961 int32
	_ = v3961
	var v3987 int32
	_ = v3987
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4010 int32
	_ = v4010
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4023 int32
	_ = v4023
	var v4025 int32
	_ = v4025
	var v4026 int32
	_ = v4026
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4031 int32
	_ = v4031
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4039 int32
	_ = v4039
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4055 int32
	_ = v4055
	var v4059 int32
	_ = v4059
	var v4061 int32
	_ = v4061
	var v4067 int32
	_ = v4067
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4079 int32
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4088 int32
	_ = v4088
	var v4093 int32
	_ = v4093
	var v4095 int32
	_ = v4095
	var v4100 int32
	_ = v4100
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4116 int32
	_ = v4116
	var v4121 int32
	_ = v4121
	var v4127 int32
	_ = v4127
	var v4128 int32
	_ = v4128
	var v4129 int32
	_ = v4129
	var v4136 int32
	_ = v4136
	var v4137 int32
	_ = v4137
	var v4140 int32
	_ = v4140
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4148 int32
	_ = v4148
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4157 int32
	_ = v4157
	var v4161 int32
	_ = v4161
	var v4162 int32
	_ = v4162
	var v4166 int32
	_ = v4166
	var v4168 int32
	_ = v4168
	var v4169 int64
	_ = v4169
	var v4177 int32
	_ = v4177
	var v4181 int32
	_ = v4181
	var v4185 int32
	_ = v4185
	var v4191 int32
	_ = v4191
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4203 int32
	_ = v4203
	var v4204 int32
	_ = v4204
	var v4205 int32
	_ = v4205
	var v4209 int32
	_ = v4209
	var v4212 int32
	_ = v4212
	var v4216 int32
	_ = v4216
	var v4217 int32
	_ = v4217
	var v4225 int32
	_ = v4225
	var v4231 int32
	_ = v4231
	var v4233 int32
	_ = v4233
	var v4235 int32
	_ = v4235
	var v4245 int32
	_ = v4245
	var v4253 int32
	_ = v4253
	var v4254 int32
	_ = v4254
	var v4257 int32
	_ = v4257
	var v4258 int32
	_ = v4258
	var v4260 int32
	_ = v4260
	var v4263 int32
	_ = v4263
	var v4266 int32
	_ = v4266
	var v4269 int32
	_ = v4269
	var v4270 int32
	_ = v4270
	var v4273 int32
	_ = v4273
	var v4274 int32
	_ = v4274
	var v4277 int32
	_ = v4277
	var v4284 int32
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4306 int32
	_ = v4306
	var v4307 int32
	_ = v4307
	var v4312 int32
	_ = v4312
	var v4313 int32
	_ = v4313
	var v4315 int32
	_ = v4315
	var v4319 int32
	_ = v4319
	var v4324 int32
	_ = v4324
	var v4328 int32
	_ = v4328
	var v4329 int32
	_ = v4329
	var v4333 int32
	_ = v4333
	var v4336 int32
	_ = v4336
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4342 int32
	_ = v4342
	var v4344 int32
	_ = v4344
	var v4346 int32
	_ = v4346
	var v4349 int32
	_ = v4349
	var v4352 int32
	_ = v4352
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4362 int32
	_ = v4362
	var v4364 int32
	_ = v4364
	var v4365 int32
	_ = v4365
	var v4366 int32
	_ = v4366
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4371 int32
	_ = v4371
	var v4372 int32
	_ = v4372
	var v4373 int32
	_ = v4373
	var v4377 int32
	_ = v4377
	var v4380 int32
	_ = v4380
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4386 int32
	_ = v4386
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4392 int32
	_ = v4392
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4397 int32
	_ = v4397
	var v4399 int32
	_ = v4399
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4411 int32
	_ = v4411
	var v4416 int32
	_ = v4416
	var v4417 int32
	_ = v4417
	var v4422 int32
	_ = v4422
	var v4423 int32
	_ = v4423
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4436 int32
	_ = v4436
	var v4438 int32
	_ = v4438
	var v4440 int32
	_ = v4440
	var v4443 int32
	_ = v4443
	var v4446 int32
	_ = v4446
	var v4447 int32
	_ = v4447
	var v4448 int32
	_ = v4448
	var v4456 int32
	_ = v4456
	var v4458 int32
	_ = v4458
	var v4459 int32
	_ = v4459
	var v4460 int32
	_ = v4460
	var v4461 int32
	_ = v4461
	var v4463 int32
	_ = v4463
	var v4465 int32
	_ = v4465
	var v4466 int32
	_ = v4466
	var v4467 int32
	_ = v4467
	var v4471 int32
	_ = v4471
	var v4476 int32
	_ = v4476
	var v4477 int32
	_ = v4477
	var v4484 int32
	_ = v4484
	var v4489 int32
	_ = v4489
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4496 int32
	_ = v4496
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4499 int32
	_ = v4499
	var v4500 int32
	_ = v4500
	var v4502 int32
	_ = v4502
	var v4503 int32
	_ = v4503
	var v4506 int32
	_ = v4506
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4519 int32
	_ = v4519
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4522 int32
	_ = v4522
	var v4523 int32
	_ = v4523
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4530 int32
	_ = v4530
	var v4531 int32
	_ = v4531
	var v4535 int32
	_ = v4535
	var v4537 int32
	_ = v4537
	var v4540 int32
	_ = v4540
	var v4541 int32
	_ = v4541
	var v4542 int32
	_ = v4542
	var v4544 int32
	_ = v4544
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4550 int32
	_ = v4550
	var v4551 int32
	_ = v4551
	var v4552 int32
	_ = v4552
	var v4558 int32
	_ = v4558
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4566 int32
	_ = v4566
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4570 int32
	_ = v4570
	var v4571 int64
	_ = v4571
	var v4582 int32
	_ = v4582
	var v4586 int32
	_ = v4586
	var v4591 int32
	_ = v4591
	var v4592 int32
	_ = v4592
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4600 int32
	_ = v4600
	var v4601 int32
	_ = v4601
	var v4602 int32
	_ = v4602
	var v4603 int32
	_ = v4603
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4607 int32
	_ = v4607
	var v4610 int32
	_ = v4610
	var v4611 int32
	_ = v4611
	var v4612 int32
	_ = v4612
	var v4614 int32
	_ = v4614
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4621 int32
	_ = v4621
	var v4624 int32
	_ = v4624
	var v4630 int32
	_ = v4630
	var v4635 int32
	_ = v4635
	var v4638 int32
	_ = v4638
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4645 int32
	_ = v4645
	var v4647 int32
	_ = v4647
	var v4651 int32
	_ = v4651
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4657 int32
	_ = v4657
	var v4661 int32
	_ = v4661
	var v4663 int32
	_ = v4663
	var v4665 int32
	_ = v4665
	var v4667 int32
	_ = v4667
	var v4669 int32
	_ = v4669
	var v4679 int32
	_ = v4679
	var v4684 int32
	_ = v4684
	var v4685 int32
	_ = v4685
	var v4688 int32
	_ = v4688
	var v4694 int32
	_ = v4694
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4704 int32
	_ = v4704
	var v4706 int32
	_ = v4706
	var v4708 int32
	_ = v4708
	var v4709 int32
	_ = v4709
	var v4713 int32
	_ = v4713
	var v4715 int32
	_ = v4715
	var v4721 int32
	_ = v4721
	var v4727 int32
	_ = v4727
	var v4728 int32
	_ = v4728
	var v4733 int32
	_ = v4733
	var v4737 int32
	_ = v4737
	var v4742 int32
	_ = v4742
	var v4747 int32
	_ = v4747
	var v4749 int32
	_ = v4749
	var v4754 int32
	_ = v4754
	var v4763 int32
	_ = v4763
	var v4764 int32
	_ = v4764
	var v4768 int32
	_ = v4768
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4775 int32
	_ = v4775
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4783 int32
	_ = v4783
	var v4791 int32
	_ = v4791
	var v4793 int32
	_ = v4793
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4808 int32
	_ = v4808
	var v4813 int32
	_ = v4813
	var v4816 int32
	_ = v4816
	var v4819 int32
	_ = v4819
	var v4821 int32
	_ = v4821
	var v4823 int32
	_ = v4823
	var v4827 int32
	_ = v4827
	var v4828 int32
	_ = v4828
	var v4829 int32
	_ = v4829
	var v4835 int32
	_ = v4835
	var v4836 int32
	_ = v4836
	var v4846 int32
	_ = v4846
	var v4851 int32
	_ = v4851
	var v4854 int32
	_ = v4854
	var v4856 int32
	_ = v4856
	var v4858 int32
	_ = v4858
	var v4863 int32
	_ = v4863
	var v4866 int32
	_ = v4866
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4877 int32
	_ = v4877
	var v4878 int32
	_ = v4878
	var v4888 int32
	_ = v4888
	var v4893 int32
	_ = v4893
	var v4896 int32
	_ = v4896
	var v4898 int32
	_ = v4898
	var v4900 int32
	_ = v4900
	var v4905 int32
	_ = v4905
	var v4908 int32
	_ = v4908
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4917 int32
	_ = v4917
	var v4918 int32
	_ = v4918
	var v4919 int32
	_ = v4919
	var v4922 int32
	_ = v4922
	var v4924 int32
	_ = v4924
	var v4935 int32
	_ = v4935
	var v4936 int32
	_ = v4936
	var v4938 int32
	_ = v4938
	var v4971 int32
	_ = v4971
	var v4972 int32
	_ = v4972
	var v4973 int64
	_ = v4973
	var v4975 int64
	_ = v4975
	var v4977 int32
	_ = v4977
	var v4982 int32
	_ = v4982
	var v4985 int32
	_ = v4985
	var v4986 int32
	_ = v4986
	var v4989 int32
	_ = v4989
	var v4999 int32
	_ = v4999
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5042 int32
	_ = v5042
	var v5046 int32
	_ = v5046
	var v5047 int32
	_ = v5047
	var v5049 int32
	_ = v5049
	var v5050 int32
	_ = v5050
	var v5054 int32
	_ = v5054
	var v5056 int32
	_ = v5056
	var v5059 int32
	_ = v5059
	var v5103 int32
	_ = v5103
	var v5104 int32
	_ = v5104
	var v5110 int32
	_ = v5110
	var v5111 int32
	_ = v5111
	var v5121 int32
	_ = v5121
	var v5126 int32
	_ = v5126
	var v5129 int32
	_ = v5129
	var v5130 int32
	_ = v5130
	var v5131 int32
	_ = v5131
	var v5137 int32
	_ = v5137
	var v5142 int32
	_ = v5142
	var v5144 int32
	_ = v5144
	var v5145 int32
	_ = v5145
	var v5146 int32
	_ = v5146
	var v5152 int32
	_ = v5152
	var v5153 int32
	_ = v5153
	var v5154 int32
	_ = v5154
	var v5164 int32
	_ = v5164
	var v5165 int32
	_ = v5165
	var v5166 int32
	_ = v5166
	var v5175 int32
	_ = v5175
	var v5178 int32
	_ = v5178
	var v5182 int32
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5187 int32
	_ = v5187
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5193 int32
	_ = v5193
	var v5198 int32
	_ = v5198
	var v5199 int32
	_ = v5199
	var v5203 int32
	_ = v5203
	var v5208 int32
	_ = v5208
	var v5209 int32
	_ = v5209
	var v5210 int32
	_ = v5210
	var v5216 int32
	_ = v5216
	var v5217 int32
	_ = v5217
	var v5218 int32
	_ = v5218
	var v5226 int64
	_ = v5226
	var v5234 int32
	_ = v5234
	var v5235 int32
	_ = v5235
	var v5238 int64
	_ = v5238
	var v5241 int64
	_ = v5241
	var v5249 int32
	_ = v5249
	var v5253 int32
	_ = v5253
	var v5256 int32
	_ = v5256
	var v5257 int32
	_ = v5257
	var v5261 int32
	_ = v5261
	var v5266 int32
	_ = v5266
	var v5267 int32
	_ = v5267
	var v5268 int32
	_ = v5268
	var v5269 int32
	_ = v5269
	var v5275 int32
	_ = v5275
	var v5276 int32
	_ = v5276
	var v5277 int32
	_ = v5277
	var v5287 int32
	_ = v5287
	var v5288 int32
	_ = v5288
	var v5289 int32
	_ = v5289
	var v5290 int32
	_ = v5290
	var v5295 int32
	_ = v5295
	var v5296 int32
	_ = v5296
	var v5300 int32
	_ = v5300
	var v5305 int32
	_ = v5305
	var v5306 int32
	_ = v5306
	var v5307 int32
	_ = v5307
	var v5308 int32
	_ = v5308
	var v5314 int32
	_ = v5314
	var v5315 int32
	_ = v5315
	var v5316 int32
	_ = v5316
	var v5324 int32
	_ = v5324
	var v5325 int32
	_ = v5325
	var v5331 int32
	_ = v5331
	var v5332 int32
	_ = v5332
	var v5333 int32
	_ = v5333
	var v5344 int64
	_ = v5344
	var v5351 int64
	_ = v5351
	var v5352 int64
	_ = v5352
	var v5355 int64
	_ = v5355
	var v5356 int64
	_ = v5356
	var v5360 int64
	_ = v5360
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5373 int32
	_ = v5373
	var v5374 int32
	_ = v5374
	var v5420 int64
	_ = v5420
	var v5424 int64
	_ = v5424
	var v5425 int64
	_ = v5425
	var v5429 int64
	_ = v5429
	var v5432 int32
	_ = v5432
	var v5436 int32
	_ = v5436
	var v5440 int32
	_ = v5440
	var v5441 int32
	_ = v5441
	var v5442 int64
	_ = v5442
	var v5445 int32
	_ = v5445
	var v5449 int32
	_ = v5449
	var v5453 int32
	_ = v5453
	var v5454 int32
	_ = v5454
	var v5456 int32
	_ = v5456
	var v5458 int32
	_ = v5458
	var v5461 int32
	_ = v5461
	var v5462 int32
	_ = v5462
	var v5464 int32
	_ = v5464
	var v5470 int64
	_ = v5470
	var v5471 int32
	_ = v5471
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5483 int32
	_ = v5483
	var v5488 int32
	_ = v5488
	var v5503 int32
	_ = v5503
	var v5512 int32
	_ = v5512
	var v5544 int32
	_ = v5544
	var v5552 int32
	_ = v5552
	var v5553 int32
	_ = v5553
	var v5555 int32
	_ = v5555
	var v5560 int32
	_ = v5560
	var v5561 int32
	_ = v5561
	var v5565 int32
	_ = v5565
	var v5567 int32
	_ = v5567
	var v5585 int32
	_ = v5585
	var v5610 int32
	_ = v5610
	var v5621 int32
	_ = v5621
	var v5635 int32
	_ = v5635
	var v5660 int32
	_ = v5660
	var v5702 int32
	_ = v5702
	var v5717 int32
	_ = v5717
	var v5754 int32
	_ = v5754
	var v5758 int32
	_ = v5758
	var v5760 int32
	_ = v5760
	var v5769 int32
	_ = v5769
	var v5771 int32
	_ = v5771
	var v5805 int32
	_ = v5805
	var v5806 int32
	_ = v5806
	var v5807 int32
	_ = v5807
	var v5808 int32
	_ = v5808
	var v5819 int32
	_ = v5819
	var v5820 int32
	_ = v5820
	var v5821 int32
	_ = v5821
	var v5828 int32
	_ = v5828
	var v5829 int32
	_ = v5829
	var v5833 int32
	_ = v5833
	var v5846 int32
	_ = v5846
	var v5920 int32
	_ = v5920
	var v5928 int32
	_ = v5928
	var v5969 int32
	_ = v5969
	var v5973 int32
	_ = v5973
	var v5978 int32
	_ = v5978
	var v5979 int32
	_ = v5979
	var v5985 int32
	_ = v5985
	var v5991 int32
	_ = v5991
	var v5992 int32
	_ = v5992
	var v6000 int32
	_ = v6000
	var v6007 int32
	_ = v6007
	var v6012 int32
	_ = v6012
	var v6017 int32
	_ = v6017
	var v6018 int32
	_ = v6018
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6030 int32
	_ = v6030
	var v6031 int32
	_ = v6031
	var v6035 int32
	_ = v6035
	var v6036 int32
	_ = v6036
	var v6048 int32
	_ = v6048
	var v6053 int32
	_ = v6053
	var v6058 int32
	_ = v6058
	var v6059 int32
	_ = v6059
	var v6068 int32
	_ = v6068
	var v6073 int32
	_ = v6073
	var v6074 int32
	_ = v6074
	var v6075 int32
	_ = v6075
	var v6085 int32
	_ = v6085
	var v6086 int32
	_ = v6086
	var v6091 int32
	_ = v6091
	var v6092 int32
	_ = v6092
	var v6104 int32
	_ = v6104
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6120 int32
	_ = v6120
	var v6122 int32
	_ = v6122
	var v6128 int32
	_ = v6128
	var v6133 int32
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6136 int32
	_ = v6136
	var v6137 int32
	_ = v6137
	var v6138 int32
	_ = v6138
	var v6140 int64
	_ = v6140
	var v6142 int64
	_ = v6142
	var v6147 int32
	_ = v6147
	var v6154 int32
	_ = v6154
	var v6157 int32
	_ = v6157
	var v6163 int32
	_ = v6163
	var v6164 int32
	_ = v6164
	var v6169 int32
	_ = v6169
	var v6170 int32
	_ = v6170
	var v6171 int32
	_ = v6171
	var v6177 int32
	_ = v6177
	var v6182 int32
	_ = v6182
	var v6184 int32
	_ = v6184
	var v6186 int32
	_ = v6186
	var v6187 int64
	_ = v6187
	var v6188 int64
	_ = v6188
	var v6190 int64
	_ = v6190
	var v6191 int64
	_ = v6191
	var v6198 int32
	_ = v6198
	var v6199 int32
	_ = v6199
	var v6207 int32
	_ = v6207
	var v6212 int32
	_ = v6212
	var v6213 int32
	_ = v6213
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6223 int32
	_ = v6223
	var v6230 int32
	_ = v6230
	var v6235 int32
	_ = v6235
	var v6244 int64
	_ = v6244
	var v6246 int64
	_ = v6246
	var v6250 int64
	_ = v6250
	var v6253 int32
	_ = v6253
	var v6254 int32
	_ = v6254
	var v6256 int32
	_ = v6256
	var v6258 int32
	_ = v6258
	var v6260 int32
	_ = v6260
	var v6262 int32
	_ = v6262
	var v6306 int32
	_ = v6306
	var v6307 int32
	_ = v6307
	var v6315 int32
	_ = v6315
	var v6358 int32
	_ = v6358
	var v6361 int32
	_ = v6361
	var v6362 int32
	_ = v6362
	var v6363 int32
	_ = v6363
	var v6364 int32
	_ = v6364
	var v6365 int32
	_ = v6365
	var v6366 int32
	_ = v6366
	var v6367 int32
	_ = v6367
	var v6368 int32
	_ = v6368
	var v6373 int32
	_ = v6373
	var v6374 int32
	_ = v6374
	var v6375 int32
	_ = v6375
	var v6376 int32
	_ = v6376
	var v6380 int32
	_ = v6380
	var v6382 int32
	_ = v6382
	var v6385 int32
	_ = v6385
	var v6386 int32
	_ = v6386
	var v6387 int32
	_ = v6387
	var v6389 int32
	_ = v6389
	var v6393 int32
	_ = v6393
	var v6394 int32
	_ = v6394
	var v6395 int32
	_ = v6395
	var v6396 int32
	_ = v6396
	var v6397 int32
	_ = v6397
	var v6403 int32
	_ = v6403
	var v6404 int32
	_ = v6404
	var v6405 int32
	_ = v6405
	var v6406 int32
	_ = v6406
	var v6407 int32
	_ = v6407
	var v6408 int32
	_ = v6408
	var v6409 int32
	_ = v6409
	var v6410 int32
	_ = v6410
	var v6415 int32
	_ = v6415
	var v6416 int32
	_ = v6416
	var v6417 int32
	_ = v6417
	var v6418 int32
	_ = v6418
	var v6422 int32
	_ = v6422
	var v6424 int32
	_ = v6424
	var v6427 int32
	_ = v6427
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6431 int32
	_ = v6431
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6437 int32
	_ = v6437
	var v6438 int32
	_ = v6438
	var v6439 int32
	_ = v6439
	var v6445 int32
	_ = v6445
	var v6446 int32
	_ = v6446
	var v6449 int32
	_ = v6449
	var v6453 int32
	_ = v6453
	var v6455 int32
	_ = v6455
	var v6460 int32
	_ = v6460
	var v6463 int32
	_ = v6463
	var v6464 int32
	_ = v6464
	var v6467 int32
	_ = v6467
	var v6471 int32
	_ = v6471
	var v6473 int32
	_ = v6473
	var v6478 int32
	_ = v6478
	var v6481 int32
	_ = v6481
	var v6482 int32
	_ = v6482
	var v6485 int32
	_ = v6485
	var v6489 int32
	_ = v6489
	var v6491 int32
	_ = v6491
	var v6496 int32
	_ = v6496
	var v6499 int32
	_ = v6499
	var v6501 int32
	_ = v6501
	var v6502 int32
	_ = v6502
	var v6504 int32
	_ = v6504
	var v6505 int32
	_ = v6505
	var v6506 int32
	_ = v6506
	var v6507 int32
	_ = v6507
	var v6508 int32
	_ = v6508
	var v6509 int32
	_ = v6509
	var v6510 int32
	_ = v6510
	var v6515 int32
	_ = v6515
	var v6516 int32
	_ = v6516
	var v6517 int32
	_ = v6517
	var v6522 int32
	_ = v6522
	var v6528 int32
	_ = v6528
	var v6537 int32
	_ = v6537
	var v6546 int32
	_ = v6546
	var v6557 int32
	_ = v6557
	var v6562 int32
	_ = v6562
	var v6566 int32
	_ = v6566
	var v6569 int32
	_ = v6569
	var v6573 int32
	_ = v6573
	var v6578 int32
	_ = v6578
	var v6580 int32
	_ = v6580
	var v6582 int32
	_ = v6582
	var v6583 int32
	_ = v6583
	var v6585 int32
	_ = v6585
	var v6586 int32
	_ = v6586
	var v6587 int32
	_ = v6587
	var v6588 int32
	_ = v6588
	var v6589 int32
	_ = v6589
	var v6590 int32
	_ = v6590
	var v6591 int32
	_ = v6591
	var v6596 int32
	_ = v6596
	var v6597 int32
	_ = v6597
	var v6598 int32
	_ = v6598
	var v6609 int32
	_ = v6609
	var v6618 int32
	_ = v6618
	var v6627 int32
	_ = v6627
	var v6634 int32
	_ = v6634
	var v6637 int32
	_ = v6637
	var v6638 int32
	_ = v6638
	var v6641 int32
	_ = v6641
	var v6642 int32
	_ = v6642
	var v6643 int32
	_ = v6643
	var v6646 int32
	_ = v6646
	var v6647 int32
	_ = v6647
	var v6648 int64
	_ = v6648
	var v6656 int32
	_ = v6656
	var v6661 int32
	_ = v6661
	var v6666 int32
	_ = v6666
	var v6668 int32
	_ = v6668
	var v6672 int32
	_ = v6672
	var v6674 int32
	_ = v6674
	var v6676 int32
	_ = v6676
	var v6679 int32
	_ = v6679
	var v6682 int32
	_ = v6682
	var v6683 int32
	_ = v6683
	var v6684 int32
	_ = v6684
	var v6692 int32
	_ = v6692
	var v6694 int32
	_ = v6694
	var v6696 int32
	_ = v6696
	var v6702 int32
	_ = v6702
	var v6703 int32
	_ = v6703
	var v6704 int32
	_ = v6704
	var v6710 int32
	_ = v6710
	var v6711 int32
	_ = v6711
	var v6712 int32
	_ = v6712
	var v6713 int32
	_ = v6713
	var v6715 int32
	_ = v6715
	var v6716 int64
	_ = v6716
	var v6717 int32
	_ = v6717
	var v6723 int32
	_ = v6723
	var v6724 int32
	_ = v6724
	var v6730 int32
	_ = v6730
	var v6731 int32
	_ = v6731
	var v6732 int32
	_ = v6732
	var v6736 int32
	_ = v6736
	var v6738 int32
	_ = v6738
	var v6739 int32
	_ = v6739
	var v6744 int32
	_ = v6744
	var v6748 int32
	_ = v6748
	var v6753 int32
	_ = v6753
	var v6756 int32
	_ = v6756
	var v6759 int32
	_ = v6759
	var v6764 int32
	_ = v6764
	var v6765 int32
	_ = v6765
	var v6766 int32
	_ = v6766
	var v6769 int64
	_ = v6769
	var v6770 int64
	_ = v6770
	var v6781 int32
	_ = v6781
	var v6785 int32
	_ = v6785
	var v6787 int32
	_ = v6787
	var v6788 int32
	_ = v6788
	var v6793 int32
	_ = v6793
	var v6794 int32
	_ = v6794
	var v6798 int32
	_ = v6798
	var v6800 int32
	_ = v6800
	var v6803 int32
	_ = v6803
	var v6809 int32
	_ = v6809
	var v6810 int32
	_ = v6810
	var v6818 int32
	_ = v6818
	var v6821 int32
	_ = v6821
	var v6822 int32
	_ = v6822
	var v6823 int32
	_ = v6823
	var v6829 int32
	_ = v6829
	var v6834 int32
	_ = v6834
	var v6835 int32
	_ = v6835
	var v6837 int32
	_ = v6837
	var v6841 int32
	_ = v6841
	var v6845 int32
	_ = v6845
	var v6847 int32
	_ = v6847
	var v6851 int32
	_ = v6851
	var v6854 int32
	_ = v6854
	var v6856 int32
	_ = v6856
	var v6857 int32
	_ = v6857
	var v6858 int32
	_ = v6858
	var v6861 int32
	_ = v6861
	var v6862 int32
	_ = v6862
	var v6869 int32
	_ = v6869
	var v6881 int32
	_ = v6881
	var v6890 int32
	_ = v6890
	var v6898 int32
	_ = v6898
	var v6899 int32
	_ = v6899
	var v6900 int32
	_ = v6900
	var v6901 int32
	_ = v6901
	var v6903 int32
	_ = v6903
	var v6904 int32
	_ = v6904
	var v6905 int32
	_ = v6905
	var v6912 int32
	_ = v6912
	var v6924 int32
	_ = v6924
	var v6933 int32
	_ = v6933
	var v6942 int32
	_ = v6942
	var v6943 int32
	_ = v6943
	var v6945 int32
	_ = v6945
	var v6946 int32
	_ = v6946
	var v6950 int32
	_ = v6950
	var v6951 int32
	_ = v6951
	var v6955 int32
	_ = v6955
	var v6965 int32
	_ = v6965
	var v6971 int32
	_ = v6971
	var v6978 int32
	_ = v6978
	var v6979 int32
	_ = v6979
	var v6985 int32
	_ = v6985
	var v6987 int32
	_ = v6987
	var v6988 int32
	_ = v6988
	var v6992 int32
	_ = v6992
	var v6993 int32
	_ = v6993
	var v6997 int32
	_ = v6997
	var v7002 int32
	_ = v7002
	var v7004 int32
	_ = v7004
	var v7005 int32
	_ = v7005
	var v7008 int32
	_ = v7008
	var v7009 int32
	_ = v7009
	var v7013 int32
	_ = v7013
	var v7022 int32
	_ = v7022
	var v7048 int32
	_ = v7048
	var v7057 int32
	_ = v7057
	var v7060 int32
	_ = v7060
	var v7104 int32
	_ = v7104
	var v7107 int32
	_ = v7107
	var v7109 int32
	_ = v7109
	var v7111 int32
	_ = v7111
	var v7114 int32
	_ = v7114
	var v7116 int32
	_ = v7116
	var v7117 int32
	_ = v7117
	var v7162 int32
	_ = v7162
	var v7166 int32
	_ = v7166
	var v7167 int32
	_ = v7167
	var v7168 int32
	_ = v7168
	var v7172 int32
	_ = v7172
	var v7176 int32
	_ = v7176
	var v7180 int32
	_ = v7180
	var v7182 int32
	_ = v7182
	var v7184 int32
	_ = v7184
	var v7190 int32
	_ = v7190
	var v7192 int32
	_ = v7192
	var v7196 int32
	_ = v7196
	var v7201 int32
	_ = v7201
	var v7205 int32
	_ = v7205
	var v7206 int32
	_ = v7206
	var v7209 int32
	_ = v7209
	var v7212 int32
	_ = v7212
	var v7213 int32
	_ = v7213
	var v7214 int32
	_ = v7214
	var v7215 int32
	_ = v7215
	var v7216 int32
	_ = v7216
	var v7217 int32
	_ = v7217
	var v7218 int32
	_ = v7218
	var v7220 int32
	_ = v7220
	var v7223 int32
	_ = v7223
	var v7226 int32
	_ = v7226
	var v7227 int32
	_ = v7227
	var v7229 int32
	_ = v7229
	var v7232 int32
	_ = v7232
	var v7240 int32
	_ = v7240
	var v7242 int32
	_ = v7242
	var v7247 int32
	_ = v7247
	var v7250 int32
	_ = v7250
	var v7251 int32
	_ = v7251
	var v7254 int32
	_ = v7254
	var v7257 int32
	_ = v7257
	var v7258 int32
	_ = v7258
	var v7259 int32
	_ = v7259
	var v7260 int32
	_ = v7260
	var v7261 int32
	_ = v7261
	var v7262 int32
	_ = v7262
	var v7263 int32
	_ = v7263
	var v7265 int32
	_ = v7265
	var v7268 int32
	_ = v7268
	var v7269 int32
	_ = v7269
	var v7270 int32
	_ = v7270
	var v7271 int32
	_ = v7271
	var v7275 int32
	_ = v7275
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7283 int32
	_ = v7283
	var v7289 int32
	_ = v7289
	var v7290 int32
	_ = v7290
	var v7296 int32
	_ = v7296
	var v7299 int32
	_ = v7299
	var v7305 int32
	_ = v7305
	var v7310 int32
	_ = v7310
	var v7312 int32
	_ = v7312
	var v7314 int32
	_ = v7314
	var v7321 int32
	_ = v7321
	var v7325 int32
	_ = v7325
	var v7337 int32
	_ = v7337
	var v7338 int32
	_ = v7338
	var v7339 int32
	_ = v7339
	var v7341 int32
	_ = v7341
	var v7345 int32
	_ = v7345
	var v7346 int32
	_ = v7346
	var v7348 int32
	_ = v7348
	var v7349 int32
	_ = v7349
	var v7350 int32
	_ = v7350
	var v7352 int32
	_ = v7352
	var v7358 int32
	_ = v7358
	var v7359 int32
	_ = v7359
	var v7360 int32
	_ = v7360
	var v7361 int32
	_ = v7361
	var v7364 int32
	_ = v7364
	var v7371 int32
	_ = v7371
	var v7372 int32
	_ = v7372
	var v7373 int32
	_ = v7373
	var v7376 int32
	_ = v7376
	var v7379 int32
	_ = v7379
	var v7384 int32
	_ = v7384
	var v7385 int32
	_ = v7385
	var v7387 int32
	_ = v7387
	var v7389 int32
	_ = v7389
	var v7393 int32
	_ = v7393
	var v7394 int32
	_ = v7394
	var v7395 int32
	_ = v7395
	var v7400 int32
	_ = v7400
	var v7401 int32
	_ = v7401
	var v7402 int32
	_ = v7402
	var v7405 int32
	_ = v7405
	var v7406 int32
	_ = v7406
	var v7407 int32
	_ = v7407
	var v7409 int32
	_ = v7409
	var v7413 int32
	_ = v7413
	var v7414 int32
	_ = v7414
	var v7416 int32
	_ = v7416
	var v7418 int32
	_ = v7418
	var v7420 int32
	_ = v7420
	var v7421 int32
	_ = v7421
	var v7424 int32
	_ = v7424
	var v7431 int32
	_ = v7431
	var v7434 int32
	_ = v7434
	var v7438 int32
	_ = v7438
	var v7441 int32
	_ = v7441
	var v7450 int32
	_ = v7450
	var v7454 int32
	_ = v7454
	var v7456 int32
	_ = v7456
	var v7457 int32
	_ = v7457
	var v7461 int32
	_ = v7461
	var v7462 int32
	_ = v7462
	var v7464 int32
	_ = v7464
	var v7468 int32
	_ = v7468
	var v7470 int32
	_ = v7470
	var v7472 int32
	_ = v7472
	var v7475 int32
	_ = v7475
	var v7480 int32
	_ = v7480
	var v7481 int32
	_ = v7481
	var v7482 int32
	_ = v7482
	var v7484 int32
	_ = v7484
	var v7485 int32
	_ = v7485
	var v7487 int32
	_ = v7487
	var v7489 int32
	_ = v7489
	var v7492 int32
	_ = v7492
	var v7497 int32
	_ = v7497
	var v7498 int32
	_ = v7498
	var v7499 int32
	_ = v7499
	var v7506 int32
	_ = v7506
	var v7508 int32
	_ = v7508
	var v7509 int32
	_ = v7509
	var v7511 int32
	_ = v7511
	var v7522 int32
	_ = v7522
	var v7525 int32
	_ = v7525
	var v7529 int32
	_ = v7529
	var v7534 int32
	_ = v7534
	var v7538 int32
	_ = v7538
	var v7541 int32
	_ = v7541
	var v7548 int32
	_ = v7548
	var v7553 int32
	_ = v7553
	var v7557 int32
	_ = v7557
	var v7560 int32
	_ = v7560
	var v7567 int32
	_ = v7567
	var v7572 int32
	_ = v7572
	var v7576 int32
	_ = v7576
	var v7579 int32
	_ = v7579
	var v7583 int32
	_ = v7583
	var v7590 int32
	_ = v7590
	var v7595 int32
	_ = v7595
	var v7599 int32
	_ = v7599
	var v7602 int32
	_ = v7602
	var v7608 int32
	_ = v7608
	var v7613 int32
	_ = v7613
	var v7621 int32
	_ = v7621
	var v7624 int32
	_ = v7624
	var v7632 int32
	_ = v7632
	var v7636 int32
	_ = v7636
	var v7641 int32
	_ = v7641
	var v7643 int32
	_ = v7643
	var v7651 int32
	_ = v7651
	var v7654 int32
	_ = v7654
	var v7656 int32
	_ = v7656
	var v7658 int32
	_ = v7658
	var v7659 int32
	_ = v7659
	var v7660 int32
	_ = v7660
	var v7662 int32
	_ = v7662
	var v7666 int32
	_ = v7666
	var v7670 int32
	_ = v7670
	var v7674 int32
	_ = v7674
	var v7680 int32
	_ = v7680
	var v7685 int32
	_ = v7685
	var v7687 int32
	_ = v7687
	var v7689 int32
	_ = v7689
	var v7691 int32
	_ = v7691
	var v7693 int32
	_ = v7693
	var v7695 int32
	_ = v7695
	var v7698 int32
	_ = v7698
	var v7699 int32
	_ = v7699
	var v7700 int32
	_ = v7700
	var v7704 int32
	_ = v7704
	var v7705 int32
	_ = v7705
	var v7706 int32
	_ = v7706
	var v7707 int32
	_ = v7707
	var v7709 int32
	_ = v7709
	var v7712 int32
	_ = v7712
	var v7715 int32
	_ = v7715
	var v7718 int32
	_ = v7718
	var v7719 int32
	_ = v7719
	var v7722 int32
	_ = v7722
	var v7723 int32
	_ = v7723
	var v7726 int32
	_ = v7726
	var v7733 int32
	_ = v7733
	var v7734 int32
	_ = v7734
	var v7737 int32
	_ = v7737
	var v7741 int32
	_ = v7741
	var v7744 int32
	_ = v7744
	var v7749 int32
	_ = v7749
	var v7756 int32
	_ = v7756
	var v7758 int32
	_ = v7758
	var v7760 int32
	_ = v7760
	var v7761 int32
	_ = v7761
	var v7763 int32
	_ = v7763
	var v7766 int32
	_ = v7766
	var v7772 int32
	_ = v7772
	var v7773 int32
	_ = v7773
	var v7775 int32
	_ = v7775
	var v7777 int32
	_ = v7777
	var v7781 int32
	_ = v7781
	var v7782 int32
	_ = v7782
	var v7783 int32
	_ = v7783
	var v7789 int32
	_ = v7789
	var v7791 int32
	_ = v7791
	var v7793 int32
	_ = v7793
	var v7835 int32
	_ = v7835
	var v7838 int32
	_ = v7838
	var v7839 int32
	_ = v7839
	var v7842 int32
	_ = v7842
	var v7845 int32
	_ = v7845
	var v7849 int32
	_ = v7849
	var v7851 int32
	_ = v7851
	var v7855 int32
	_ = v7855
	var v7895 int32
	_ = v7895
	var v7899 int32
	_ = v7899
	var v7900 int32
	_ = v7900
	var v7943 int32
	_ = v7943
	var v7944 int32
	_ = v7944
	var v7946 int32
	_ = v7946
	var v7953 int32
	_ = v7953
	var v7957 int32
	_ = v7957
	var v7962 int32
	_ = v7962
	var v7965 int32
	_ = v7965
	var v7975 int32
	_ = v7975
	var v7979 int32
	_ = v7979
	var v7982 int32
	_ = v7982
	var v7983 int32
	_ = v7983
	var v7987 int32
	_ = v7987
	var v7990 int32
	_ = v7990
	var v7991 int32
	_ = v7991
	var v7992 int32
	_ = v7992
	var v7993 int32
	_ = v7993
	var v7996 int32
	_ = v7996
	var v7997 int32
	_ = v7997
	var v7998 int32
	_ = v7998
	var v7999 int32
	_ = v7999
	var v8001 int32
	_ = v8001
	var v8002 int32
	_ = v8002
	var v8006 int32
	_ = v8006
	var v8007 int32
	_ = v8007
	var v8010 int32
	_ = v8010
	var v8013 int32
	_ = v8013
	var v8016 int32
	_ = v8016
	var v8019 int32
	_ = v8019
	var v8022 int32
	_ = v8022
	var v8025 int32
	_ = v8025
	var v8026 int32
	_ = v8026
	var v8029 int32
	_ = v8029
	var v8030 int32
	_ = v8030
	var v8033 int32
	_ = v8033
	var v8040 int32
	_ = v8040
	var v8041 int32
	_ = v8041
	var v8044 int32
	_ = v8044
	var v8046 int32
	_ = v8046
	var v8048 int32
	_ = v8048
	var v8052 int32
	_ = v8052
	var v8053 int32
	_ = v8053
	var v8054 int32
	_ = v8054
	var v8055 int32
	_ = v8055
	var v8056 int32
	_ = v8056
	var v8057 int32
	_ = v8057
	var v8058 int32
	_ = v8058
	var v8063 int32
	_ = v8063
	var v8064 int32
	_ = v8064
	var v8067 int32
	_ = v8067
	var v8068 int32
	_ = v8068
	var v8069 int32
	_ = v8069
	var v8073 int32
	_ = v8073
	var v8074 int32
	_ = v8074
	var v8082 int32
	_ = v8082
	var v8087 int32
	_ = v8087
	var v8090 int32
	_ = v8090
	var v8091 int32
	_ = v8091
	var v8092 int32
	_ = v8092
	var v8093 int32
	_ = v8093
	var v8094 int32
	_ = v8094
	var v8097 int32
	_ = v8097
	var v8106 int32
	_ = v8106
	var v8108 int32
	_ = v8108
	var v8112 int32
	_ = v8112
	var v8117 int32
	_ = v8117
	var v8122 int32
	_ = v8122
	var v8123 int32
	_ = v8123
	var v8124 int32
	_ = v8124
	var v8125 int32
	_ = v8125
	var v8126 int32
	_ = v8126
	var v8128 int32
	_ = v8128
	var v8133 int32
	_ = v8133
	var v8134 int32
	_ = v8134
	var v8135 int32
	_ = v8135
	var v8136 int32
	_ = v8136
	var v8137 int32
	_ = v8137
	var v8139 int32
	_ = v8139
	var v8140 int32
	_ = v8140
	var v8142 int32
	_ = v8142
	var v8143 int32
	_ = v8143
	var v8148 int32
	_ = v8148
	var v8149 int32
	_ = v8149
	var v8159 int32
	_ = v8159
	var v8163 int32
	_ = v8163
	var v8166 int32
	_ = v8166
	var v8169 int32
	_ = v8169
	var v8170 int32
	_ = v8170
	var v8173 int32
	_ = v8173
	var v8174 int32
	_ = v8174
	var v8177 int32
	_ = v8177
	var v8184 int32
	_ = v8184
	var v8185 int32
	_ = v8185
	var v8191 int32
	_ = v8191
	var v8192 int32
	_ = v8192
	var v8196 int32
	_ = v8196
	var v8202 int32
	_ = v8202
	var v8209 int32
	_ = v8209
	var v8210 int32
	_ = v8210
	var v8211 int32
	_ = v8211
	var v8217 int32
	_ = v8217
	var v8220 int32
	_ = v8220
	var v8223 int32
	_ = v8223
	var v8228 int32
	_ = v8228
	var v8230 int32
	_ = v8230
	var v8232 int32
	_ = v8232
	var v8234 int32
	_ = v8234
	var v8236 int32
	_ = v8236
	var v8279 int32
	_ = v8279
	var v8281 int32
	_ = v8281
	var v8283 int32
	_ = v8283
	var v8285 int32
	_ = v8285
	var v8287 int32
	_ = v8287
	var v8292 int32
	_ = v8292
	var v8293 int32
	_ = v8293
	var v8295 int32
	_ = v8295
	var v8296 int32
	_ = v8296
	var v8297 int32
	_ = v8297
	var v8298 int32
	_ = v8298
	var v8301 int32
	_ = v8301
	var v8305 int32
	_ = v8305
	var v8309 int32
	_ = v8309
	var v8310 int32
	_ = v8310
	var v8314 int32
	_ = v8314
	var v8316 int32
	_ = v8316
	var v8319 int32
	_ = v8319
	var v8323 int32
	_ = v8323
	var v8329 int32
	_ = v8329
	var v8330 int32
	_ = v8330
	var v8332 int32
	_ = v8332
	var v8335 int32
	_ = v8335
	var v8338 int32
	_ = v8338
	var v8339 int32
	_ = v8339
	var v8342 int32
	_ = v8342
	var v8344 int32
	_ = v8344
	var v8351 int32
	_ = v8351
	var v8352 int32
	_ = v8352
	var v8355 int32
	_ = v8355
	var v8357 int32
	_ = v8357
	var v8360 int32
	_ = v8360
	var v8361 int32
	_ = v8361
	var v8368 int32
	_ = v8368
	var v8372 int32
	_ = v8372
	var v8376 int32
	_ = v8376
	var v8380 int32
	_ = v8380
	var v8382 int32
	_ = v8382
	var v8384 int64
	_ = v8384
	var v8392 int32
	_ = v8392
	var v8397 int32
	_ = v8397
	var v8402 int32
	_ = v8402
	var v8407 int32
	_ = v8407
	var v8409 int32
	_ = v8409
	var v8412 int32
	_ = v8412
	var v8420 int32
	_ = v8420
	var v8423 int32
	_ = v8423
	var v8425 int32
	_ = v8425
	var v8426 int32
	_ = v8426
	var v8431 int32
	_ = v8431
	var v8435 int32
	_ = v8435
	var v8437 int32
	_ = v8437
	var v8441 int32
	_ = v8441
	var v8486 int32
	_ = v8486
	var v8488 int32
	_ = v8488
	var v8501 int32
	_ = v8501
	var v8535 int32
	_ = v8535
	var v8543 int32
	_ = v8543
	var v8547 int32
	_ = v8547
	var v8552 int32
	_ = v8552
	var v8556 int32
	_ = v8556
	var v8558 int32
	_ = v8558
	var v8564 int32
	_ = v8564
	var v8569 int32
	_ = v8569
	var v8573 int32
	_ = v8573
	var v8576 int32
	_ = v8576
	var v8584 int32
	_ = v8584
	var v8587 int32
	_ = v8587
	var v8593 int32
	_ = v8593
	var v8598 int32
	_ = v8598
	var v8602 int32
	_ = v8602
	var v8605 int32
	_ = v8605
	var v8613 int32
	_ = v8613
	var v8618 int32
	_ = v8618
	var v8622 int32
	_ = v8622
	var v8625 int32
	_ = v8625
	var v8633 int32
	_ = v8633
	var v8637 int32
	_ = v8637
	var v8642 int32
	_ = v8642
	var v8646 int32
	_ = v8646
	var v8649 int32
	_ = v8649
	var v8657 int32
	_ = v8657
	var v8662 int32
	_ = v8662
	var v8666 int32
	_ = v8666
	var v8670 int32
	_ = v8670
	var v8676 int32
	_ = v8676
	var v8680 int32
	_ = v8680
	var v8685 int32
	_ = v8685
	var v8689 int32
	_ = v8689
	var v8693 int32
	_ = v8693
	var v8699 int32
	_ = v8699
	var v8703 int32
	_ = v8703
	var v8708 int32
	_ = v8708
	var v8713 int32
	_ = v8713
	var v8716 int32
	_ = v8716
	var v8722 int32
	_ = v8722
	var v8726 int32
	_ = v8726
	var v8731 int32
	_ = v8731
	v7 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(512)
	m.G0 = v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+428)) = v7
	v52 = F_errstart(m, int32(12), v7)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v52 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(_a_F_InitPostgres_0), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[1]))
	F_ProcArrayAdd(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(729), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	F_on_shmem_exit(m, int32(1117), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_pgstat_beinit(m)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v123 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	F_pgstat_bestart_initial(m)
	mBase = m.M
	F_SharedInvalBackendInit(m, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_SharedInvalBackendInit(m, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[3]))
	F_ProcSignalInit(m, int32(_a_F_InitPostgres_2), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_RegisterTimeout(m, int32(1), int32(1624))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_RegisterTimeout(m, int32(3), int32(1625))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_RegisterTimeout(m, int32(2), int32(1626))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_RegisterTimeout(m, int32(7), int32(1627))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_RegisterTimeout(m, int32(8), int32(1628))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_RegisterTimeout(m, int32(9), int32(1629))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_RegisterTimeout(m, int32(11), int32(1630))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_RegisterTimeout(m, int32(10), int32(1631))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L11
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[3]))
	F_ProcSignalInit(m, int32(_a_F_InitPostgres_2), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L11
L27:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v144 = m.G0
	v146 = v144 - int32(48)
	m.G0 = v146
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4]))
	if v149 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	F_StartupXLOG(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_ReleaseAuxProcessResources(m, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[5])) = v134
	F_before_shmem_exit(m, int32(930), v134)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_before_shmem_exit(m, int32(1632), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v146)+16)) = int64(34359738372)
	v160 = F_hash_create(m, int32(_a_F_InitPostgres_3), int32(400), v146, int32(40))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[6])) = v160
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4]))
	v166 = F_MemoryContextAlloc(m, v164, int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[7])) = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[8])) = v166
	v174 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[9])) = v174
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[10])) = v174
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[11])) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[12])) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[13])) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[14])) = v180
	m.G0 = v146 + int32(48)
	v195 = m.G0
	v197 = v195 - int32(16)
	m.G0 = v197
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v180
	v213 = v180
	goto L43
L41:
	;
	F_CacheRegisterRelcacheCallback(m, int32(1585))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L92
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L89
	}
L43:
	;
	v247 = v213 << (uint(int32(5)) % 32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_InitPostgres[17])))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_InitPostgres[18])))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_InitPostgres[19])))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_InitPostgres[20])))
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4]))
	if v253 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	F_pg_qsort(m, int32(_a_F_InitPostgres_4), v547, int32(4), int32(1596))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L69
	}
L45:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v263 = F_AllocSetContextCreateInternal(m, v258, int32(_a_F_InitPostgres_5), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v266 = v253
	goto L47
L47:
	;
	v267 = int32(_a_F_InitPostgres_8)
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v266
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23]))
	if v272 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4])) = v263
	v266 = v263
	goto L47
L49:
	;
	v277 = F_palloc(m, int32(8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v284 = v247 + int32(_a_F_InitPostgres_9)
	v288 = F_palloc_aligned(m, int32(296), int32(128), int32(4))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23])) = v277
	*(*int64)(unsafe.Add(mBase, uint32(v277))) = int64(0)
	goto L51
L53:
	;
	v292 = F_palloc0(m, v251<<(uint(int32(3))%32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+12)) = v292
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+96)) = uint8(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v288)+92)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v288)+88)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v288)+84)) = int32(_a_F_InitPostgres_10)
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v213
	v302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v288)+68)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v288)+8)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v288)+76)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v288)+64)) = v250
	if v250 <= v295 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[23]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+100)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v288 + int32(100)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v213<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[24]))) = v288
	if v288 == int32(0) {
		goto L42
	} else {
		goto L67
	}
L56:
	;
	v313 = v250 & int32(3)
	v315 = v288 + int32(48)
	v316 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v250) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v329 = v316
	v335 = int32(0)
	goto L60
L58:
	;
	v401 = v316
	goto L59
L59:
	;
	v443 = v401
	v446 = int32(0)
	goto L64
L60:
	;
	v364 = v329 << (uint(int32(2)) % 32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364+v284)))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v364))) = v367
	v369 = int32(4)
	v370 = v364 | v369
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v284+v370)))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v370))) = v373
	v376 = v364 | int32(8)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v284+v376)))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v376))) = v379
	v382 = v364 | int32(12)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382+v284)))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v382))) = v385
	v388 = v329 + v369
	v390 = v335 + v369
	if v390 != v250&int32(2147483644) {
		v329 = v388
		v335 = v390
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if v313 == int32(0) {
		goto L55
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v401 = v388
	goto L59
L64:
	;
	v478 = v443 << (uint(int32(2)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v478+v284)))
	*(*int32)(unsafe.Add(mBase, uint32(v315+v478))) = v481
	v483 = int32(1)
	v486 = v446 + v483
	if v486 != v313 {
		v443 = v443 + v483
		v446 = v486
		goto L64
	} else {
		goto L66
	}
L65:
	;
	goto L55
L66:
	;
	goto L65
L67:
	;
	v543 = int32(_a_F_InitPostgres_11)
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15]))
	v546 = int32(1)
	v547 = v545 + v546
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v547
	v549 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v545<<(uint(v549)%32))+uint32(_c_F_InitPostgres[25]))) = v248
	v554 = int32(_a_F_InitPostgres_12)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	v557 = v555 << (uint(v549) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v557)+uint32(_c_F_InitPostgres[26]))) = v248
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v555 + v549
	*(*int32)(unsafe.Add(mBase, uint32(v557)+uint32(_c_F_InitPostgres[27]))) = v249
	v569 = v213 + v546
	if v569 != int32(85) {
		v213 = v569
		goto L43
	} else {
		goto L68
	}
L68:
	;
	goto L44
L69:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15]))
	if base.Ui32(int32(2)) <= base.Ui32(v579) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v590 = int32(0)
	v591 = int32(1)
	goto L73
L71:
	;
	v653 = v579
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[15])) = v653
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	F_pg_qsort(m, int32(_a_F_InitPostgres_13), v690, int32(4), int32(1596))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L79
	}
L73:
	;
	v624 = int32(2)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v591<<(uint(v624)%32))+uint32(_c_F_InitPostgres[25])))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v590<<(uint(v624)%32))+uint32(_c_F_InitPostgres[25])))
	if v626 == v629 {
		v637 = v590
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v653 = v637 + int32(1)
	goto L72
L75:
	;
	v640 = v591 + int32(1)
	if v640 != v579 {
		v590 = v637
		v591 = v640
		goto L73
	} else {
		goto L78
	}
L76:
	;
	v632 = v590 + int32(1)
	if v632 == v591 {
		v637 = v591
		goto L75
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[25]))) = v626
	v637 = v632
	goto L75
L78:
	;
	goto L74
L79:
	;
	v695 = int32(_a_F_InitPostgres_12)
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16]))
	if base.Ui32(int32(2)) <= base.Ui32(v697) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v708 = int32(0)
	v709 = int32(1)
	goto L83
L81:
	;
	v803 = v697
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[16])) = v803
	m.G0 = v197 + int32(16)
	goto L41
L83:
	;
	v742 = int32(2)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v709<<(uint(v742)%32))+uint32(_c_F_InitPostgres[26])))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v708<<(uint(v742)%32))+uint32(_c_F_InitPostgres[26])))
	if v744 == v747 {
		v755 = v708
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v803 = v755 + int32(1)
	goto L82
L85:
	;
	v758 = v709 + int32(1)
	if v758 != v697 {
		v708 = v755
		v709 = v758
		goto L83
	} else {
		goto L88
	}
L86:
	;
	v750 = v708 + int32(1)
	if v750 == v709 {
		v755 = v709
		goto L85
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v750<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[26]))) = v744
	v755 = v750
	goto L85
L88:
	;
	goto L84
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v248
	F_errmsg_internal(m, int32(_a_F_InitPostgres_14), v197)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_15), int32(136), int32(_a_F_InitPostgres_16))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_CacheRegisterSyscacheCallback(m, int32(47), int32(1586), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_CacheRegisterSyscacheCallback(m, int32(82), int32(1586), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_CacheRegisterSyscacheCallback(m, int32(38), int32(1587), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_CacheRegisterSyscacheCallback(m, int32(40), int32(1587), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_CacheRegisterSyscacheCallback(m, int32(3), int32(1587), int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_CacheRegisterSyscacheCallback(m, int32(32), int32(1587), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_CacheRegisterSyscacheCallback(m, int32(30), int32(1587), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v860 = m.G0
	v862 = v860 - int32(48)
	m.G0 = v862
	v866 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v871 = F_AllocSetContextCreateInternal(m, v866, int32(_a_F_InitPostgres_17), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[28])) = v871
	*(*int64)(unsafe.Add(mBase, uint32(v862)+16)) = int64(292057776192)
	v880 = F_hash_create(m, int32(_a_F_InitPostgres_18), int32(16), v862, int32(24))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[29])) = v880
	m.G0 = v862 + int32(48)
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v887 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_read_relmap_file(m, int32(_a_F_InitPostgres_19), int32(_a_F_InitPostgres_20), int32(0), int32(22))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v895 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v895 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v896 = int32(_a_F_InitPostgres_8)
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v900
	v903 = F_load_relcache_init_file(m, int32(1))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_before_shmem_exit(m, int32(1633), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L118
	}
L109:
	;
	if v903 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_21), int32(1248), int32(1), int32(18), int32(_a_F_InitPostgres_22))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v897
	goto L108
L113:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_23), int32(2842), int32(1), int32(12), int32(_a_F_InitPostgres_24))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_25), int32(2843), int32(1), int32(7), int32(_a_F_InitPostgres_26))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_27), int32(4066), int32(1), int32(4), int32(_a_F_InitPostgres_28))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_formrdesc(m, int32(_a_F_InitPostgres_29), int32(_a_F_InitPostgres_30), int32(1), int32(18), int32(_a_F_InitPostgres_31))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L112
L118:
	;
	v950 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	if v950 == int32(3) {
		goto L129
	} else {
		goto L130
	}
L119:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8713 = m.ExcPending
	if v8713 != 0 {
		goto L1
	} else {
		goto L1821
	}
L120:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8689 = m.ExcPending
	if v8689 != 0 {
		goto L1
	} else {
		goto L1816
	}
L121:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8666 = m.ExcPending
	if v8666 != 0 {
		goto L1
	} else {
		goto L1811
	}
L122:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8646 = m.ExcPending
	if v8646 != 0 {
		goto L1
	} else {
		goto L1807
	}
L123:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8622 = m.ExcPending
	if v8622 != 0 {
		goto L1
	} else {
		goto L1802
	}
L124:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8602 = m.ExcPending
	if v8602 != 0 {
		goto L1
	} else {
		goto L1798
	}
L125:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8573 = m.ExcPending
	if v8573 != 0 {
		goto L1
	} else {
		goto L1793
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8556 = m.ExcPending
	if v8556 != 0 {
		goto L1
	} else {
		goto L1790
	}
L127:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8535 = m.ExcPending
	if v8535 != 0 {
		goto L1
	} else {
		goto L1786
	}
L128:
	;
	m.G0 = v8501 + int32(512)
	return
L129:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v47 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v8501 = v44
	goto L128
L133:
	;
	v6942 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	if v6942 != 0 {
		goto L1424
	} else {
		goto L1425
	}
L134:
	;
	v6898 = F_superuser(m)
	mBase = m.M
	v6899 = m.ExcPending
	if v6899 != 0 {
		goto L1
	} else {
		goto L1423
	}
L135:
	;
	v1072 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[32])) = uint8(v1072)
	v1075 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v1080 = m.G0
	v1081 = int32(16)
	v1082 = v1080 - v1081
	m.G0 = v1082
	F_gettimeofday(m, v1082)
	mBase = m.M
	v1085 = *(*int64)(unsafe.Add(mBase, uint32(v1082)))
	v1086 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1082)+8)))
	m.G0 = v1082 + v1081
	goto L178
L136:
	;
	F_InitializeSessionUserId(m, l2, l3, int32(base.Ui32(l4&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L177
	}
L137:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L176
	}
L138:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v958 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L143
	}
L140:
	;
	v962 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[34])) = v962
	goto L142
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[35])) = int32(1)
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	switch v970 - int32(4) {
	case 0, 3:
		goto L137
	default:
		goto L144
	}
L144:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v974 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v977 = int32(_a_F_InitPostgres_32)
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[36])))
	if base.B2i32(v980 == int32(0))|base.B2i32(v980 != v983) != 0 {
		v1001 = v980
		v1002 = v983
		goto L150
	} else {
		goto L151
	}
L146:
	;
	goto L147
L147:
	;
	if v970 != int32(5) {
		goto L135
	} else {
		goto L174
	}
L148:
	;
	v1018 = F_table_open(m, int32(1260), int32(1))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L162
	}
L149:
	;
	if v1001-v1002 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L150:
	;
	goto L149
L151:
	;
	v986 = l2
	v987 = v977
	goto L152
L152:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	if v991 == int32(0) {
		v1001 = v991
		v1002 = v990
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v1001 = v991
	v1002 = v990
	goto L150
L154:
	;
	v994 = int32(1)
	if v991 == v990 {
		v986 = v986 + v994
		v987 = v987 + v994
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v1009 = int32(0)
	F_InitializeSessionUserId(m, l2, v1009, v1009)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	v1015 = int32(1)
	goto L148
L160:
	;
	v1013 = F_superuser(m)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v1015 = v1013
	goto L148
L162:
	;
	v1020 = int32(0)
	v1022 = F_table_beginscan_catalog(m, v1018, v1020, v1020)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v1024 = F_heap_getnext(m, v1022)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1022)))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+188))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+12))
	m.T0[v1028].(func(*base.Module, int32))(m, v1022)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_relation_close(m, v1018, int32(1))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	if v1024 != 0 {
		v6900 = l0
		v6901 = l1
		v6903 = v1015
		v6904 = l4
		v6905 = l5
		v6912 = v44
		v6924 = v47
		v6933 = v7
		goto L133
	} else {
		goto L167
	}
L167:
	;
	v1036 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	if v1036 == int32(0) {
		v6900 = l0
		v6901 = l1
		v6903 = v1015
		v6904 = l4
		v6905 = l5
		v6912 = v44
		v6924 = v47
		v6933 = v7
		goto L133
	} else {
		goto L169
	}
L169:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_33), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+416)) = l2
	F_errhint(m, int32(_a_F_InitPostgres_34), v44+int32(416))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(896), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v6900 = l0
	v6901 = l1
	v6903 = v1015
	v6904 = l4
	v6905 = l5
	v6912 = v44
	v6924 = v47
	v6933 = v7
	goto L133
L174:
	;
	if l2|l3 != 0 {
		goto L136
	} else {
		goto L175
	}
L175:
	;
	goto L137
L176:
	;
	v6900 = l0
	v6901 = l1
	v6903 = int32(1)
	v6904 = l4
	v6905 = l5
	v6912 = v44
	v6924 = v47
	v6933 = v7
	goto L133
L177:
	;
	v6857 = l0
	v6858 = l1
	v6861 = l4
	v6862 = l5
	v6869 = v44
	v6881 = v47
	v6890 = v7
	goto L134
L178:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[37])) = v1086 + v1085*int64(1000000) - int64(946684800000000)
	v1098 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[38]))
	F_enable_timeout_after(m, int32(3), v1098*int32(1000))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v1103 = int32(0)
	v1104 = m.G0
	v1106 = v1104 - int32(3792)
	m.G0 = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+700)) = v1103
	v1110 = m.G0
	v1112 = v1110 - int32(288)
	m.G0 = v1112
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v1116 = F_get_role_oid(m, v1114, int32(1))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[39]))
	if v1119 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+380)) = v2569
	m.G0 = v1112 + int32(288)
	v2602 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v2602 != 0 {
		goto L497
	} else {
		goto L498
	}
L182:
	;
	v2552 = F_palloc0(m, int32(420))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L496
	}
L183:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	if v1122 <= int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v1126 = v1075 + int32(144)
	v1149 = v7
	goto L185
L185:
	;
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126))))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+12))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1169+v1149<<(uint(int32(2))%32))))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+12))
	if v1174 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	goto L182
L187:
	;
	v2507 = v1149 + int32(1)
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	if v2507 < v2508 {
		v1149 = v2507
		goto L185
	} else {
		goto L495
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+288)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-2)
	goto L187
L189:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+16))
	if v2118 == int32(0) {
		goto L187
	} else {
		goto L397
	}
L190:
	;
	if v1168 == int32(1) {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if v1168 == int32(1) {
		goto L187
	} else {
		goto L194
	}
L193:
	;
	goto L187
L194:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+488)))
	if v1181 == int32(1) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+288))
	switch v1190 {
	case 0:
		goto L202
	case 1, 2:
		goto L201
	case 3:
		goto L189
	default:
		goto L187
	}
L196:
	;
	if base.Ui32(v1174-int32(3)) < base.Ui32(int32(2)) {
		goto L187
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	switch v1174 - int32(2) {
	case 0, 2:
		goto L187
	default:
		goto L195
	}
L199:
	;
	goto L195
L200:
	;
	v2010 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1173)+24)))
	if v1168 != v2010 {
		goto L187
	} else {
		goto L386
	}
L201:
	;
	v1444 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1112)+24)) = uint8(v1444)
	*(*int32)(unsafe.Add(mBase, uint32(v1112)+20)) = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v1112)+16)) = v1190
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41])) = v1444
	v1452 = v1112 + int32(16)
	v1453 = m.G0
	v1455 = v1453 - int32(144)
	m.G0 = v1455
	v1459 = m.G0
	v1461 = v1459 - int32(272)
	m.G0 = v1461
	v1464 = v1461 + int32(8)
	goto L271
L202:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+292))
	if v1191 == int32(0) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	if v1194 < int32(0) {
		goto L187
	} else {
		goto L204
	}
L204:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v1197 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v1202 = v1112 + int32(16)
	v1204 = int32(0)
	v1207 = F_pg_getnameinfo_all(m, v1126, v1200, v1202, int32(255), v1204, v1204, int32(8))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	v1213 = v1197
	goto L207
L207:
	;
	v1214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191))))
	if v1214 == int32(46) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	if v1207 != 0 {
		goto L188
	} else {
		goto L209
	}
L209:
	;
	v1209 = F_pstrdup(m, v1202)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+280)) = v1209
	v1213 = v1209
	goto L207
L211:
	;
	v1217 = F_strlen(m, v1191)
	mBase = m.M
	v1218 = F_strlen(m, v1213)
	mBase = m.M
	if base.Ui32(v1218) < base.Ui32(v1217) {
		goto L187
	} else {
		goto L214
	}
L212:
	;
	v1224 = v1213
	goto L213
L213:
	;
	v1227 = v1191
	v1228 = v1224
	goto L216
L214:
	;
	v1224 = v1213 + (v1218 - v1217)
	goto L213
L215:
	;
	if v1265 != 0 {
		goto L187
	} else {
		goto L228
	}
L216:
	;
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227))))
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228))))
	if v1231 == v1232 {
		v1254 = v1231
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1265 = int32(0)
	goto L215
L218:
	;
	v1256 = int32(1)
	if v1254 != 0 {
		v1227 = v1227 + v1256
		v1228 = v1228 + v1256
		goto L216
	} else {
		goto L227
	}
L219:
	;
	if base.Ui32((v1231-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1242 = v1231 | int32(32)
	goto L222
L221:
	;
	v1242 = v1231
	goto L222
L222:
	;
	if base.Ui32((v1232-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1251 = v1232 | int32(32)
	goto L225
L224:
	;
	v1251 = v1232
	goto L225
L225:
	;
	if v1242 == v1251 {
		v1254 = v1242
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v1265 = v1242 - v1251
	goto L215
L227:
	;
	goto L217
L228:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	if v1266 == int32(1) {
		goto L189
	} else {
		goto L229
	}
L229:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	v1270 = int32(0)
	v1274 = m.Env.Getaddrinfo(m, v1269, v1270, v1270, v1112+int32(284))
	mBase = m.M
	if v1274 == v1270 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1112)+284))
	if v1277 != 0 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+288)) = v1274
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-2)
	goto L187
L233:
	;
	v1278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126))))
	v1289 = v1277
	goto L236
L234:
	;
	goto L235
L235:
	;
	v1428 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L263
	}
L236:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+20))
	v1323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1322))))
	if v1323 != v1278 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+20))
	F_emscripten_builtin_free(m, v1382)
	mBase = m.M
	F_emscripten_builtin_free(m, v1277)
	mBase = m.M
	goto L262
L238:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1289)+28))
	if v1381 != 0 {
		v1289 = v1381
		goto L236
	} else {
		goto L261
	}
L239:
	;
	switch v1278 - int32(2) {
	case 0:
		goto L241
	default:
		goto L238
	case 8:
		goto L242
	}
L240:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1277)+20))
	F_emscripten_builtin_free(m, v1376)
	mBase = m.M
	F_emscripten_builtin_free(m, v1277)
	mBase = m.M
	goto L260
L241:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1322)+4))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+148))
	if v1373 != v1374 {
		goto L238
	} else {
		goto L259
	}
L242:
	;
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+8)))
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+152)))
	if v1325 != v1326 {
		goto L238
	} else {
		goto L243
	}
L243:
	;
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+9)))
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+153)))
	if v1328 != v1329 {
		goto L238
	} else {
		goto L244
	}
L244:
	;
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+10)))
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+154)))
	if v1331 != v1332 {
		goto L238
	} else {
		goto L245
	}
L245:
	;
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+11)))
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+155)))
	if v1334 != v1335 {
		goto L238
	} else {
		goto L246
	}
L246:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+12)))
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+156)))
	if v1337 != v1338 {
		goto L238
	} else {
		goto L247
	}
L247:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+13)))
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+157)))
	if v1340 != v1341 {
		goto L238
	} else {
		goto L248
	}
L248:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+14)))
	v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+158)))
	if v1343 != v1344 {
		goto L238
	} else {
		goto L249
	}
L249:
	;
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+15)))
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+159)))
	if v1346 != v1347 {
		goto L238
	} else {
		goto L250
	}
L250:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+16)))
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+160)))
	if v1349 != v1350 {
		goto L238
	} else {
		goto L251
	}
L251:
	;
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+17)))
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+161)))
	if v1352 != v1353 {
		goto L238
	} else {
		goto L252
	}
L252:
	;
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+18)))
	v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+162)))
	if v1355 != v1356 {
		goto L238
	} else {
		goto L253
	}
L253:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+19)))
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+163)))
	if v1358 != v1359 {
		goto L238
	} else {
		goto L254
	}
L254:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+20)))
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+164)))
	if v1361 != v1362 {
		goto L238
	} else {
		goto L255
	}
L255:
	;
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+21)))
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+165)))
	if v1364 != v1365 {
		goto L238
	} else {
		goto L256
	}
L256:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+22)))
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+166)))
	if v1367 != v1368 {
		goto L238
	} else {
		goto L257
	}
L257:
	;
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+23)))
	v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+167)))
	if v1370 == v1371 {
		goto L240
	} else {
		goto L258
	}
L258:
	;
	goto L238
L259:
	;
	goto L240
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(1)
	goto L189
L261:
	;
	goto L237
L262:
	;
	goto L235
L263:
	;
	if v1428 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1112))) = v1191
	F_errmsg_internal(m, int32(_a_F_InitPostgres_35), v1112)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-1)
	goto L187
L267:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_36), int32(1158), int32(_a_F_InitPostgres_37))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v1575 = int32(0)
	v1576 = F_socket(m, int32(16), int32(_a_F_InitPostgres_38), v1575)
	mBase = m.M
	if v1576 < v1575 {
		goto L281
	} else {
		goto L282
	}
L270:
	;
	goto L269
L271:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1464))) = uint8(v1444)
	v1473 = v1461 + int32(272)
	*(*uint8)(unsafe.Add(mBase, uint32(v1473-int32(1)))) = uint8(v1444)
	goto L272
L272:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1464)+2)) = uint8(v1444)
	*(*uint8)(unsafe.Add(mBase, uint32(v1464)+1)) = uint8(v1444)
	*(*uint8)(unsafe.Add(mBase, uint32(v1473-int32(3)))) = uint8(v1444)
	*(*uint8)(unsafe.Add(mBase, uint32(v1473-int32(2)))) = uint8(v1444)
	goto L273
L273:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1464)+3)) = uint8(v1444)
	*(*uint8)(unsafe.Add(mBase, uint32(v1473-int32(4)))) = uint8(v1444)
	goto L274
L274:
	;
	v1495 = int32(0)
	v1498 = (v1495 - v1464) & int32(3)
	v1499 = v1464 + v1498
	*(*int32)(unsafe.Add(mBase, uint32(v1499))) = v1495
	v1507 = (int32(264) - v1498) & int32(-4)
	v1508 = v1499 + v1507
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(4)))) = v1495
	if base.Ui32(v1507) < base.Ui32(int32(9)) {
		goto L270
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+8)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+4)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(8)))) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(12)))) = v1495
	if base.Ui32(v1507) < base.Ui32(int32(25)) {
		goto L270
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+24)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+20)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+16)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1499)+12)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(16)))) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(20)))) = v1495
	v1534 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1508-v1534))) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v1508-int32(28)))) = v1495
	v1543 = v1499&int32(4) | v1534
	v1544 = v1507 - v1543
	if base.Ui32(v1544) < base.Ui32(int32(32)) {
		goto L270
	} else {
		goto L277
	}
L277:
	;
	v1549 = base.I64_extend_i32_u(v1495) * int64(4294967297)
	v1552 = v1543 + v1499
	v1553 = v1544
	goto L278
L278:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1552)+24)) = v1549
	*(*int64)(unsafe.Add(mBase, uint32(v1552)+16)) = v1549
	*(*int64)(unsafe.Add(mBase, uint32(v1552)+8)) = v1549
	*(*int64)(unsafe.Add(mBase, uint32(v1552))) = v1549
	v1561 = int32(32)
	v1564 = v1553 - v1561
	if base.Ui32(int32(31)) < base.Ui32(v1564) {
		v1552 = v1552 + v1561
		v1553 = v1564
		goto L278
	} else {
		goto L280
	}
L279:
	;
	goto L270
L280:
	;
	goto L279
L281:
	;
	v1750 = int32(-1)
	goto L283
L282:
	;
	v1581 = int32(18)
	v1583 = v1461 + int32(8)
	v1584 = int32(0)
	v1586 = m.G0
	v1588 = v1586 + int32(-8192)
	m.G0 = v1588
	v1591 = int32(20)
	F___memset(m, v1588, v1584, v1591)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1586)+uint32(_c_F_InitPostgres[42]))) = uint8(v1584)
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+uint32(_c_F_InitPostgres[43]))) = int32(1)
	v1596 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1586)+uint32(_c_F_InitPostgres[44]))) = uint16(v1596)
	*(*uint16)(unsafe.Add(mBase, uint32(v1586)+uint32(_c_F_InitPostgres[45]))) = uint16(v1581)
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+uint32(_c_F_InitPostgres[46]))) = v1591
	v1602 = F_send(m, v1576, v1588, v1591)
	mBase = m.M
	if v1602 < v1584 {
		v1656 = v1602
		goto L285
	} else {
		goto L286
	}
L283:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+8))
	if v1750 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L284:
	;
	if v1656 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L285:
	;
	m.G0 = v1588 - int32(-8192)
	goto L284
L286:
	;
	v1605 = F_recv(m, v1576, v1588)
	mBase = m.M
	if v1605 <= int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1656 = int32(-1)
	goto L285
L288:
	;
	v1610 = v1605
	goto L289
L289:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1610) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v1656 = int32(0)
	goto L285
L291:
	;
	goto L290
L292:
	;
	v1618 = v1588
	goto L295
L293:
	;
	goto L294
L294:
	;
	v1643 = F_recv(m, v1576, v1588)
	mBase = m.M
	if int32(0) < v1643 {
		v1610 = v1643
		goto L289
	} else {
		goto L300
	}
L295:
	;
	v1624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1618)+4)))
	switch v1624 - int32(2) {
	case 0:
		v1656 = int32(-1)
		goto L285
	case 1:
		goto L291
	default:
		goto L297
	}
L296:
	;
	goto L294
L297:
	;
	v1627 = F_netlink_msg_to_ifaddr(m, v1583, v1618)
	mBase = m.M
	if v1627 != 0 {
		v1656 = v1627
		goto L285
	} else {
		goto L298
	}
L298:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1618)))
	v1633 = v1618 + (v1628+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1610+v1588-v1633) {
		v1618 = v1633
		goto L295
	} else {
		goto L299
	}
L299:
	;
	goto L296
L300:
	;
	goto L287
L301:
	;
	v1666 = int32(22)
	v1667 = int32(0)
	v1669 = m.G0
	v1671 = v1669 + int32(-8192)
	m.G0 = v1671
	v1674 = int32(20)
	F___memset(m, v1671, v1667, v1674)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1669)+uint32(_c_F_InitPostgres[42]))) = uint8(v1667)
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+uint32(_c_F_InitPostgres[43]))) = int32(2)
	v1679 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+uint32(_c_F_InitPostgres[44]))) = uint16(v1679)
	*(*uint16)(unsafe.Add(mBase, uint32(v1669)+uint32(_c_F_InitPostgres[45]))) = uint16(v1666)
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+uint32(_c_F_InitPostgres[46]))) = v1674
	v1685 = F_send(m, v1576, v1671, v1674)
	mBase = m.M
	if v1685 < v1667 {
		v1739 = v1685
		goto L305
	} else {
		goto L306
	}
L302:
	;
	v1746 = v1656
	goto L303
L303:
	;
	v1747 = m.Wasi_snapshot_preview1.Fd_close(m, v1576)
	mBase = m.M
	v1750 = v1746
	goto L283
L304:
	;
	v1746 = v1739
	goto L303
L305:
	;
	m.G0 = v1671 - int32(-8192)
	goto L304
L306:
	;
	v1688 = F_recv(m, v1576, v1671)
	mBase = m.M
	if v1688 <= int32(0) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1739 = int32(-1)
	goto L305
L308:
	;
	v1693 = v1688
	goto L309
L309:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1693) {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1739 = int32(0)
	goto L305
L311:
	;
	goto L310
L312:
	;
	v1701 = v1671
	goto L315
L313:
	;
	goto L314
L314:
	;
	v1726 = F_recv(m, v1576, v1671)
	mBase = m.M
	if int32(0) < v1726 {
		v1693 = v1726
		goto L309
	} else {
		goto L320
	}
L315:
	;
	v1707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1701)+4)))
	switch v1707 - int32(2) {
	case 0:
		v1739 = int32(-1)
		goto L305
	case 1:
		goto L311
	default:
		goto L317
	}
L316:
	;
	goto L314
L317:
	;
	v1710 = F_netlink_msg_to_ifaddr(m, v1583, v1701)
	mBase = m.M
	if v1710 != 0 {
		v1739 = v1710
		goto L305
	} else {
		goto L318
	}
L318:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1701)))
	v1716 = v1701 + (v1711+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1693+v1671-v1716) {
		v1701 = v1716
		goto L315
	} else {
		goto L319
	}
L319:
	;
	goto L316
L320:
	;
	goto L307
L321:
	;
	m.G0 = v1461 + int32(272)
	if v1750 < int32(0) {
		goto L332
	} else {
		goto L333
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1455+int32(12)))) = v1751
	goto L321
L323:
	;
	goto L324
L324:
	;
	if v1751 != 0 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	goto L321
L326:
	;
	v1756 = v1751
	goto L329
L327:
	;
	goto L328
L328:
	;
	goto L325
L329:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1756)))
	F_emscripten_builtin_free(m, v1756)
	mBase = m.M
	if v1758 != 0 {
		v1756 = v1758
		goto L329
	} else {
		goto L331
	}
L330:
	;
	goto L328
L331:
	;
	goto L330
L332:
	;
	v1986 = int32(-1)
	goto L334
L333:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+12))
	if v1768 != 0 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	m.G0 = v1455 + int32(144)
	if v1986 < int32(0) {
		goto L378
	} else {
		goto L379
	}
L335:
	;
	v1770 = v1455 + int32(24)
	v1780 = v1768
	goto L338
L336:
	;
	v1936 = int32(0)
	goto L337
L337:
	;
	if v1936 != 0 {
		goto L372
	} else {
		goto L373
	}
L338:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+12))
	if v1812 != 0 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1455)+12))
	v1936 = v1893
	goto L337
L340:
	;
	v1813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1812))))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+16))
	if v1814 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L341:
	;
	goto L342
L342:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	if v1892 != 0 {
		v1780 = v1892
		goto L338
	} else {
		goto L370
	}
L343:
	;
	v1851 = m.G0
	v1853 = v1851 - int32(128)
	m.G0 = v1853
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1452)+8)))
	if v1855 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L344:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1455)+16)) = uint16(v1813)
	v1848 = v1455 + int32(16)
	goto L343
L345:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1455)+16)) = int64(0)
	v1838 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1770)+8)) = v1838
	*(*int64)(unsafe.Add(mBase, uint32(v1770))) = v1838
	*(*int32)(unsafe.Add(mBase, uint32(v1455)+40)) = int32(0)
	goto L344
L346:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1455)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1455)+16)) = int64(-4294967296)
	goto L344
L347:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+4))
	if v1829 != 0 {
		v1848 = v1814
		goto L343
	} else {
		goto L356
	}
L348:
	;
	switch v1813 - int32(2) {
	case 0:
		goto L346
	default:
		v1848 = v1455 + int32(16)
		goto L343
	case 8:
		goto L345
	}
L349:
	;
	v1817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1814))))
	if v1817 != v1813 {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	switch v1813 - int32(2) {
	case 0:
		goto L347
	default:
		v1848 = v1814
		goto L343
	case 8:
		goto L351
	}
L351:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+8))
	if v1821 != 0 {
		v1848 = v1814
		goto L343
	} else {
		goto L352
	}
L352:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+12))
	if v1822 != 0 {
		v1848 = v1814
		goto L343
	} else {
		goto L353
	}
L353:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+16))
	if v1823 != 0 {
		v1848 = v1814
		goto L343
	} else {
		goto L354
	}
L354:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+20))
	if v1824 != 0 {
		v1848 = v1814
		goto L343
	} else {
		goto L355
	}
L355:
	;
	goto L345
L356:
	;
	goto L346
L357:
	;
	goto L342
L358:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	if v1858 == int32(1) {
		goto L362
	} else {
		goto L363
	}
L359:
	;
	goto L360
L360:
	;
	m.G0 = v1853 + int32(128)
	goto L357
L361:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1452)+8)) = uint8(v1882)
	goto L360
L362:
	;
	v1861 = int32(0)
	v1863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1812))))
	v1864 = F_pg_sockaddr_cidr_mask(m, v1853, v1861, v1863)
	mBase = m.M
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+4))
	v1866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1865))))
	v1867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1812))))
	if v1866 != v1867 {
		v1882 = v1861
		goto L361
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+4))
	v1874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1873))))
	v1875 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1812))))
	if v1874 != v1875 {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	v1869 = F_pg_range_sockaddr(m, v1865, v1812, v1853)
	mBase = m.M
	if v1869 == int32(0) {
		v1882 = v1861
		goto L361
	} else {
		goto L366
	}
L366:
	;
	v1882 = int32(1)
	goto L361
L367:
	;
	v1882 = int32(0)
	goto L361
L368:
	;
	v1877 = F_pg_range_sockaddr(m, v1873, v1812, v1848)
	mBase = m.M
	if v1877 == int32(0) {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v1882 = int32(1)
	goto L361
L370:
	;
	goto L339
L371:
	;
	v1986 = int32(0)
	goto L334
L372:
	;
	v1938 = v1936
	goto L375
L373:
	;
	goto L374
L374:
	;
	goto L371
L375:
	;
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1938)))
	F_emscripten_builtin_free(m, v1938)
	mBase = m.M
	if v1940 != 0 {
		v1938 = v1940
		goto L375
	} else {
		goto L377
	}
L376:
	;
	goto L374
L377:
	;
	goto L376
L378:
	;
	v1994 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112)+24)))
	if v2007 == int32(0) {
		goto L187
	} else {
		goto L385
	}
L381:
	;
	if v1994 == int32(0) {
		goto L187
	} else {
		goto L382
	}
L382:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_39), int32(0))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_36), int32(1222), int32(_a_F_InitPostgres_40))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	goto L187
L385:
	;
	goto L189
L386:
	;
	v2018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126))))
	switch v2018 - int32(2) {
	case 0:
		goto L390
	default:
		goto L388
	case 8:
		goto L389
	}
L387:
	;
	if v2074 == int32(0) {
		goto L187
	} else {
		goto L396
	}
L388:
	;
	v2074 = int32(0)
	goto L387
L389:
	;
	v2029 = v1173 + int32(164)
	v2031 = v1173 + int32(32)
	v2033 = v1075 + int32(152)
	v2035 = int32(0)
	goto L391
L390:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v1173+int32(156))+4))
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1173+int32(24))+4))
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+4))
	v2074 = base.B2i32(v2021&(v2022^v2023) == int32(0))
	goto L387
L391:
	;
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2035+v2029))))
	v2043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2035+v2031))))
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2035+v2033))))
	if v2041&(v2043^v2045) != 0 {
		goto L388
	} else {
		goto L393
	}
L392:
	;
	v2074 = int32(1)
	goto L387
L393:
	;
	v2049 = v2035 | int32(1)
	v2051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031+v2049))))
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033+v2049))))
	v2056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2029+v2049))))
	if (v2051^v2053)&v2056 != 0 {
		goto L388
	} else {
		goto L394
	}
L394:
	;
	v2059 = v2035 + int32(2)
	if v2059 != int32(16) {
		v2035 = v2059
		goto L391
	} else {
		goto L395
	}
L395:
	;
	goto L392
L396:
	;
	goto L189
L397:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v2118)+4))
	if v2121 <= int32(0) {
		goto L187
	} else {
		goto L398
	}
L398:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	v2135 = int32(0)
	goto L399
L399:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v2118)+12))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2168+v2135<<(uint(int32(2))%32))))
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2172)+4)))
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	if v2175 != int32(1) {
		goto L403
	} else {
		goto L404
	}
L400:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+20))
	v2458 = F_check_role_2(m, v2456, v1116, v2457)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L1
	} else {
		goto L493
	}
L401:
	;
	goto L400
L402:
	;
	v2450 = v2135 + int32(1)
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2118)+4))
	if v2450 < v2451 {
		v2135 = v2450
		goto L399
	} else {
		goto L492
	}
L403:
	;
	if v2173&int32(1) == int32(0) {
		goto L415
	} else {
		goto L416
	}
L404:
	;
	v2179 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[48])))
	if v2179&int32(1) != 0 {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	if v2173&int32(1) != 0 {
		goto L402
	} else {
		goto L406
	}
L406:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2172)))
	v2185 = int32(_a_F_InitPostgres_41)
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184))))
	v2191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[49])))
	if base.B2i32(v2188 == int32(0))|base.B2i32(v2188 != v2191) != 0 {
		v2209 = v2188
		v2210 = v2191
		goto L408
	} else {
		goto L409
	}
L407:
	;
	if v2209-v2210 != 0 {
		goto L402
	} else {
		goto L414
	}
L408:
	;
	goto L407
L409:
	;
	v2194 = v2184
	v2195 = v2185
	goto L410
L410:
	;
	v2198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2195)+1)))
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194)+1)))
	if v2199 == int32(0) {
		v2209 = v2199
		v2210 = v2198
		goto L408
	} else {
		goto L412
	}
L411:
	;
	v2209 = v2199
	v2210 = v2198
	goto L408
L412:
	;
	v2202 = int32(1)
	if v2199 == v2198 {
		v2194 = v2194 + v2202
		v2195 = v2195 + v2202
		goto L410
	} else {
		goto L413
	}
L413:
	;
	goto L411
L414:
	;
	goto L401
L415:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2172)))
	v2217 = int32(_a_F_InitPostgres_42)
	v2220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2216))))
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if base.B2i32(v2220 == int32(0))|base.B2i32(v2220 != v2223) != 0 {
		v2241 = v2220
		v2242 = v2223
		goto L419
	} else {
		goto L420
	}
L416:
	;
	goto L417
L417:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2172)+8))
	if v2396 != 0 {
		goto L476
	} else {
		goto L477
	}
L418:
	;
	if v2241-v2242 == int32(0) {
		goto L401
	} else {
		goto L425
	}
L419:
	;
	goto L418
L420:
	;
	v2226 = v2216
	v2227 = v2217
	goto L421
L421:
	;
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2227)+1)))
	v2231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2226)+1)))
	if v2231 == int32(0) {
		v2241 = v2231
		v2242 = v2230
		goto L419
	} else {
		goto L423
	}
L422:
	;
	v2241 = v2231
	v2242 = v2230
	goto L419
L423:
	;
	v2234 = int32(1)
	if v2231 == v2230 {
		v2226 = v2226 + v2234
		v2227 = v2227 + v2234
		goto L421
	} else {
		goto L424
	}
L424:
	;
	goto L422
L425:
	;
	v2246 = int32(_a_F_InitPostgres_43)
	v2249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2216))))
	v2252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	if base.B2i32(v2249 == int32(0))|base.B2i32(v2249 != v2252) != 0 {
		v2270 = v2249
		v2271 = v2252
		goto L427
	} else {
		goto L428
	}
L426:
	;
	if v2270-v2271 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L427:
	;
	goto L426
L428:
	;
	v2255 = v2216
	v2256 = v2246
	goto L429
L429:
	;
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2256)+1)))
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2255)+1)))
	if v2260 == int32(0) {
		v2270 = v2260
		v2271 = v2259
		goto L427
	} else {
		goto L431
	}
L430:
	;
	v2270 = v2260
	v2271 = v2259
	goto L427
L431:
	;
	v2263 = int32(1)
	if v2260 == v2259 {
		v2255 = v2255 + v2263
		v2256 = v2256 + v2263
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	v2277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125))))
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2124))))
	if base.B2i32(v2277 == int32(0))|base.B2i32(v2277 != v2280) != 0 {
		v2298 = v2277
		v2299 = v2280
		goto L437
	} else {
		goto L438
	}
L434:
	;
	goto L435
L435:
	;
	v2303 = int32(_a_F_InitPostgres_44)
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2216))))
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[52])))
	if base.B2i32(v2306 == int32(0))|base.B2i32(v2306 != v2309) != 0 {
		v2327 = v2306
		v2328 = v2309
		goto L446
	} else {
		goto L447
	}
L436:
	;
	if v2298-v2299 == int32(0) {
		goto L401
	} else {
		goto L443
	}
L437:
	;
	goto L436
L438:
	;
	v2283 = v2125
	v2284 = v2124
	goto L439
L439:
	;
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2284)+1)))
	v2288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2283)+1)))
	if v2288 == int32(0) {
		v2298 = v2288
		v2299 = v2287
		goto L437
	} else {
		goto L441
	}
L440:
	;
	v2298 = v2288
	v2299 = v2287
	goto L437
L441:
	;
	v2291 = int32(1)
	if v2288 == v2287 {
		v2283 = v2283 + v2291
		v2284 = v2284 + v2291
		goto L439
	} else {
		goto L442
	}
L442:
	;
	goto L440
L443:
	;
	goto L402
L444:
	;
	v2366 = int32(_a_F_InitPostgres_41)
	v2369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2216))))
	v2372 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[49])))
	if base.B2i32(v2369 == int32(0))|base.B2i32(v2369 != v2372) != 0 {
		v2390 = v2369
		v2391 = v2372
		goto L469
	} else {
		goto L470
	}
L445:
	;
	if v2327-v2328 != 0 {
		goto L452
	} else {
		goto L453
	}
L446:
	;
	goto L445
L447:
	;
	v2312 = v2216
	v2313 = v2303
	goto L448
L448:
	;
	v2316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2313)+1)))
	v2317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2312)+1)))
	if v2317 == int32(0) {
		v2327 = v2317
		v2328 = v2316
		goto L446
	} else {
		goto L450
	}
L449:
	;
	v2327 = v2317
	v2328 = v2316
	goto L446
L450:
	;
	v2320 = int32(1)
	if v2317 == v2316 {
		v2312 = v2312 + v2320
		v2313 = v2313 + v2320
		goto L448
	} else {
		goto L451
	}
L451:
	;
	goto L449
L452:
	;
	v2330 = int32(_a_F_InitPostgres_45)
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2216))))
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[53])))
	if base.B2i32(v2333 == int32(0))|base.B2i32(v2333 != v2336) != 0 {
		v2354 = v2333
		v2355 = v2336
		goto L456
	} else {
		goto L457
	}
L453:
	;
	goto L454
L454:
	;
	if v1116 == int32(0) {
		goto L402
	} else {
		goto L463
	}
L455:
	;
	if v2354-v2355 != 0 {
		goto L444
	} else {
		goto L462
	}
L456:
	;
	goto L455
L457:
	;
	v2339 = v2216
	v2340 = v2330
	goto L458
L458:
	;
	v2343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2340)+1)))
	v2344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2339)+1)))
	if v2344 == int32(0) {
		v2354 = v2344
		v2355 = v2343
		goto L456
	} else {
		goto L460
	}
L459:
	;
	v2354 = v2344
	v2355 = v2343
	goto L456
L460:
	;
	v2347 = int32(1)
	if v2344 == v2343 {
		v2339 = v2339 + v2347
		v2340 = v2340 + v2347
		goto L458
	} else {
		goto L461
	}
L461:
	;
	goto L459
L462:
	;
	goto L454
L463:
	;
	v2360 = F_get_role_oid(m, v2125, int32(1))
	mBase = m.M
	v2361 = m.ExcPending
	if v2361 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	if v2360 == int32(0) {
		goto L402
	} else {
		goto L465
	}
L465:
	;
	v2364 = F_is_member_of_role_nosuper(m, v1116, v2360)
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	if v2364 != 0 {
		goto L401
	} else {
		goto L467
	}
L467:
	;
	goto L402
L468:
	;
	if v2390-v2391 == int32(0) {
		goto L402
	} else {
		goto L475
	}
L469:
	;
	goto L468
L470:
	;
	v2375 = v2216
	v2376 = v2366
	goto L471
L471:
	;
	v2379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2376)+1)))
	v2380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2375)+1)))
	if v2380 == int32(0) {
		v2390 = v2380
		v2391 = v2379
		goto L469
	} else {
		goto L473
	}
L472:
	;
	v2390 = v2380
	v2391 = v2379
	goto L469
L473:
	;
	v2383 = int32(1)
	if v2380 == v2379 {
		v2375 = v2375 + v2383
		v2376 = v2376 + v2383
		goto L471
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	goto L417
L476:
	;
	v2397 = F_strlen(m, v2125)
	mBase = m.M
	v2402 = F_palloc(m, v2397<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L1
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v2172)))
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2417))))
	v2423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2125))))
	if base.B2i32(v2420 == int32(0))|base.B2i32(v2420 != v2423) != 0 {
		v2441 = v2420
		v2442 = v2423
		goto L485
	} else {
		goto L486
	}
L479:
	;
	v2404 = F_strlen(m, v2125)
	mBase = m.M
	v2405 = F_pg_mb2wchar_with_len(m, v2125, v2402, v2404)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2172)+8))
	v2408 = int32(0)
	v2411 = F_pg_regexec(m, v2407, v2402, v2405, v2408, v2408, v2408)
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	F_pfree(m, v2402)
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	if v2411 == int32(0) {
		goto L401
	} else {
		goto L483
	}
L483:
	;
	goto L402
L484:
	;
	if v2441-v2442 == int32(0) {
		goto L401
	} else {
		goto L491
	}
L485:
	;
	goto L484
L486:
	;
	v2426 = v2417
	v2427 = v2125
	goto L487
L487:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2427)+1)))
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2426)+1)))
	if v2431 == int32(0) {
		v2441 = v2431
		v2442 = v2430
		goto L485
	} else {
		goto L489
	}
L488:
	;
	v2441 = v2431
	v2442 = v2430
	goto L485
L489:
	;
	v2434 = int32(1)
	if v2431 == v2430 {
		v2426 = v2426 + v2434
		v2427 = v2427 + v2434
		goto L487
	} else {
		goto L490
	}
L490:
	;
	goto L488
L491:
	;
	goto L402
L492:
	;
	goto L187
L493:
	;
	if v2458 == int32(0) {
		goto L187
	} else {
		goto L494
	}
L494:
	;
	v2569 = v1173
	goto L181
L495:
	;
	goto L186
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2552)+296)) = int32(1)
	v2569 = v2552
	goto L181
L497:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+356))
	if v2606 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L500:
	;
	goto L499
L501:
	;
	v6627 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[54])))
	if base.B2i32(v6627&int32(2) == int32(0))|v6598 != 0 {
		goto L1346
	} else {
		goto L1347
	}
L502:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = int32(0)
	v6609 = v47
	v6618 = v7
	goto L501
L503:
	;
	v6580 = int32(0)
	v6582 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_46), v1075, v6580, v6580)
	mBase = m.M
	v6583 = m.ExcPending
	if v6583 != 0 {
		goto L1
	} else {
		goto L1345
	}
L504:
	;
	v2609 = int32(-1)
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+296))
	switch v2610 {
	case 0:
		goto L514
	case 1:
		goto L513
	case 2, 12:
		goto L502
	case 3:
		goto L511
	case 4:
		goto L509
	case 5, 6:
		goto L510
	default:
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v2609
		v6609 = v47
		v6618 = v7
		goto L501
	case 13:
		goto L508
	case 14:
		goto L512
	case 15:
		goto L503
	}
L505:
	;
	goto L506
L506:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6566 = m.ExcPending
	if v6566 != 0 {
		goto L1
	} else {
		goto L1341
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+104)) = int32(_a_F_InitPostgres_47)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+100)) = v2633
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+96)) = v1106 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_48), v1106+int32(96))
	mBase = m.M
	v6557 = m.ExcPending
	if v6557 != 0 {
		goto L1
	} else {
		goto L1339
	}
L508:
	;
	v4399 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+372))
	if v4399 == int32(0) {
		goto L865
	} else {
		goto L866
	}
L509:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v4342 != 0 {
		goto L839
	} else {
		goto L840
	}
L510:
	;
	v4031 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4034 = F_get_role_password(m, v4031, v1106+int32(700))
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L1
	} else {
		goto L751
	}
L511:
	;
	v2999 = v1106 + int32(1568)
	v3002 = int32(132)
	base.MemoryCopy(m, v2999, v1075+int32(144), v3002)
	v3005 = v1106 + int32(1432)
	base.MemoryCopy(m, v3005, v1075+int32(12), v3002)
	v3010 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+1708)) = v3010
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+1704)) = v3010
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1696))
	v3016 = v1106 + int32(1168)
	v3022 = F_pg_getnameinfo_all(m, v2999, v3014, v3016, int32(255), v1106+int32(1136), int32(32), int32(3))
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L1
	} else {
		goto L621
	}
L512:
	;
	v2918 = m.G0
	v2920 = v2918 - int32(16)
	m.G0 = v2920
	*(*int32)(unsafe.Add(mBase, uint32(v2920))) = int32(12)
	goto L599
L513:
	;
	v2659 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v2661 = v1106 + int32(2752)
	v2663 = int32(0)
	v2666 = F_pg_getnameinfo_all(m, v1075+int32(144), v2659, v2661, int32(255), v2663, v2663, int32(1))
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L1
	} else {
		goto L521
	}
L514:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v2615 = v1106 + int32(2752)
	v2617 = int32(0)
	v2620 = F_pg_getnameinfo_all(m, v1075+int32(144), v2613, v2615, int32(255), v2617, v2617, int32(1))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[48])))
	v2625 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2634 = int32(1)
	if base.B2i32(v2623&v2634 == int32(0))&base.B2i32(v2625 == v2634) != 0 {
		goto L507
	} else {
		goto L518
	}
L518:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+92)) = int32(_a_F_InitPostgres_47)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+88)) = v2641
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+84)) = v2633
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+80)) = v2615
	F_errmsg(m, int32(_a_F_InitPostgres_49), v1106+int32(80))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(468), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L521:
	;
	v2669 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[48])))
	v2671 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2675 = m.ExcPending
	if v2675 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2680 = int32(1)
	if v2669&v2680|base.B2i32(v2671 != v2680) == int32(0) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+296)) = int32(_a_F_InitPostgres_47)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+292)) = v2679
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+288)) = v2661
	F_errmsg(m, int32(_a_F_InitPostgres_52), v1106+int32(288))
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L1
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+204)) = int32(_a_F_InitPostgres_47)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+200)) = v2800
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+196)) = v2679
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+192)) = v1106 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_53), v1106+int32(192))
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L1
	} else {
		goto L563
	}
L527:
	;
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	v2697 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v2697 != 0 {
		goto L529
	} else {
		goto L530
	}
L528:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(528), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L562
	}
L529:
	;
	switch v2696 + int32(2) {
	case 0:
		goto L532
	case 1:
		goto L533
	case 2:
		goto L534
	case 3:
		goto L535
	default:
		goto L528
	}
L530:
	;
	goto L531
L531:
	;
	if v2696 != int32(-2) {
		goto L528
	} else {
		goto L550
	}
L532:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2721 = int32(_a_F_InitPostgres_54)
	v2723 = v2718 + int32(1)
	if v2723 == int32(0) {
		v2743 = v2721
		goto L540
	} else {
		goto L541
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+256)) = v2697
	F_errdetail_log(m, int32(_a_F_InitPostgres_55), v1106+int32(256))
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L1
	} else {
		goto L538
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+240)) = v2697
	F_errdetail_log(m, int32(_a_F_InitPostgres_56), v1106+int32(240))
	mBase = m.M
	v2711 = m.ExcPending
	if v2711 != 0 {
		goto L1
	} else {
		goto L537
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+224)) = v2697
	F_errdetail_log(m, int32(_a_F_InitPostgres_57), v1106+int32(224))
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	goto L528
L537:
	;
	goto L528
L538:
	;
	goto L528
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+276)) = v2743 + base.B2i32(v2745 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+272)) = v2697
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v1106+int32(272))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L549
	}
L540:
	;
	v2745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2743))))
	goto L539
L541:
	;
	v2727 = v2721
	v2728 = v2723
	goto L542
L542:
	;
	v2729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2727))))
	if v2729 == int32(0) {
		v2743 = v2727
		goto L540
	} else {
		goto L544
	}
L543:
	;
	v2743 = v2739
	goto L540
L544:
	;
	v2733 = v2727
	goto L545
L545:
	;
	v2737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2733)+1)))
	if v2737 != 0 {
		v2733 = v2733 + int32(1)
		goto L545
	} else {
		goto L547
	}
L546:
	;
	v2739 = v2733 + int32(2)
	v2741 = v2728 + int32(1)
	if v2741 != 0 {
		v2727 = v2739
		v2728 = v2741
		goto L542
	} else {
		goto L548
	}
L547:
	;
	goto L546
L548:
	;
	goto L543
L549:
	;
	goto L528
L550:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2761 = int32(_a_F_InitPostgres_54)
	v2763 = v2758 + int32(1)
	if v2763 == int32(0) {
		v2783 = v2761
		goto L552
	} else {
		goto L553
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+208)) = v2783 + base.B2i32(v2785 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v1106+int32(208))
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
		goto L1
	} else {
		goto L561
	}
L552:
	;
	v2785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2783))))
	goto L551
L553:
	;
	v2767 = v2761
	v2768 = v2763
	goto L554
L554:
	;
	v2769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2767))))
	if v2769 == int32(0) {
		v2783 = v2767
		goto L552
	} else {
		goto L556
	}
L555:
	;
	v2783 = v2779
	goto L552
L556:
	;
	v2773 = v2767
	goto L557
L557:
	;
	v2777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2773)+1)))
	if v2777 != 0 {
		v2773 = v2773 + int32(1)
		goto L557
	} else {
		goto L559
	}
L558:
	;
	v2779 = v2773 + int32(2)
	v2781 = v2768 + int32(1)
	if v2781 != 0 {
		v2767 = v2779
		v2768 = v2781
		goto L554
	} else {
		goto L560
	}
L559:
	;
	goto L558
L560:
	;
	goto L555
L561:
	;
	goto L528
L562:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L563:
	;
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v2814 != 0 {
		goto L565
	} else {
		goto L566
	}
L564:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(537), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L1
	} else {
		goto L598
	}
L565:
	;
	switch v2813 + int32(2) {
	case 0:
		goto L568
	case 1:
		goto L569
	case 2:
		goto L570
	case 3:
		goto L571
	default:
		goto L564
	}
L566:
	;
	goto L567
L567:
	;
	if v2813 != int32(-2) {
		goto L564
	} else {
		goto L586
	}
L568:
	;
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2838 = int32(_a_F_InitPostgres_54)
	v2840 = v2835 + int32(1)
	if v2840 == int32(0) {
		v2860 = v2838
		goto L576
	} else {
		goto L577
	}
L569:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+160)) = v2814
	F_errdetail_log(m, int32(_a_F_InitPostgres_55), v1106+int32(160))
	mBase = m.M
	v2834 = m.ExcPending
	if v2834 != 0 {
		goto L1
	} else {
		goto L574
	}
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+144)) = v2814
	F_errdetail_log(m, int32(_a_F_InitPostgres_56), v1106+int32(144))
	mBase = m.M
	v2828 = m.ExcPending
	if v2828 != 0 {
		goto L1
	} else {
		goto L573
	}
L571:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+128)) = v2814
	F_errdetail_log(m, int32(_a_F_InitPostgres_57), v1106+int32(128))
	mBase = m.M
	v2822 = m.ExcPending
	if v2822 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	goto L564
L573:
	;
	goto L564
L574:
	;
	goto L564
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+180)) = v2860 + base.B2i32(v2862 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+176)) = v2814
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v1106+int32(176))
	mBase = m.M
	v2872 = m.ExcPending
	if v2872 != 0 {
		goto L1
	} else {
		goto L585
	}
L576:
	;
	v2862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2860))))
	goto L575
L577:
	;
	v2844 = v2838
	v2845 = v2840
	goto L578
L578:
	;
	v2846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2844))))
	if v2846 == int32(0) {
		v2860 = v2844
		goto L576
	} else {
		goto L580
	}
L579:
	;
	v2860 = v2856
	goto L576
L580:
	;
	v2850 = v2844
	goto L581
L581:
	;
	v2854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2850)+1)))
	if v2854 != 0 {
		v2850 = v2850 + int32(1)
		goto L581
	} else {
		goto L583
	}
L582:
	;
	v2856 = v2850 + int32(2)
	v2858 = v2845 + int32(1)
	if v2858 != 0 {
		v2844 = v2856
		v2845 = v2858
		goto L578
	} else {
		goto L584
	}
L583:
	;
	goto L582
L584:
	;
	goto L579
L585:
	;
	goto L564
L586:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2878 = int32(_a_F_InitPostgres_54)
	v2880 = v2875 + int32(1)
	if v2880 == int32(0) {
		v2900 = v2878
		goto L588
	} else {
		goto L589
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+112)) = v2900 + base.B2i32(v2902 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v1106+int32(112))
	mBase = m.M
	v2911 = m.ExcPending
	if v2911 != 0 {
		goto L1
	} else {
		goto L597
	}
L588:
	;
	v2902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2900))))
	goto L587
L589:
	;
	v2884 = v2878
	v2885 = v2880
	goto L590
L590:
	;
	v2886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2884))))
	if v2886 == int32(0) {
		v2900 = v2884
		goto L588
	} else {
		goto L592
	}
L591:
	;
	v2900 = v2896
	goto L588
L592:
	;
	v2890 = v2884
	goto L593
L593:
	;
	v2894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2890)+1)))
	if v2894 != 0 {
		v2890 = v2890 + int32(1)
		goto L593
	} else {
		goto L595
	}
L594:
	;
	v2896 = v2890 + int32(2)
	v2898 = v2885 + int32(1)
	if v2898 != 0 {
		v2884 = v2896
		v2885 = v2898
		goto L590
	} else {
		goto L596
	}
L595:
	;
	goto L594
L596:
	;
	goto L591
L597:
	;
	goto L564
L598:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L599:
	;
	v2930 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1106+int32(1168)))) = v2930
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1106+int32(880)))) = v2934
	goto L601
L601:
	;
	m.G0 = v2920 + int32(16)
	goto L603
L603:
	;
	goto L604
L604:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41])) = int32(44)
	v2982 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	if v2982 == int32(0) {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v2609
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L618
	}
L618:
	;
	v2986 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1168))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+320)) = v2986
	F_errmsg(m, int32(_a_F_InitPostgres_60), v1106+int32(320))
	mBase = m.M
	v2992 = m.ExcPending
	if v2992 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(1888), int32(_a_F_InitPostgres_61))
	mBase = m.M
	v2997 = m.ExcPending
	if v2997 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v2609
	v6609 = v47
	v6618 = v7
	goto L501
L621:
	;
	v3024 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1560))
	v3026 = v1106 + int32(880)
	v3032 = F_pg_getnameinfo_all(m, v3005, v3024, v3026, int32(255), v1106+int32(848), int32(32), int32(3))
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+432)) = int32(113)
	v3037 = v1106 + int32(816)
	v3042 = F_pg_snprintf(m, v3037, int32(32), int32(_a_F_InitPostgres_62), v1106+int32(432))
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L1
	} else {
		goto L623
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+704)) = int32(4)
	v3046 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1106)+716)) = v3046
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+712)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1106)+724)) = v3046
	v3052 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+732)) = v3052
	v3054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106)+1568)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+708)) = v3054
	v3057 = v1106 + int32(704)
	v3060 = F_pg_getaddrinfo_all(m, v3016, v3037, v3057, v1106+int32(1708))
	mBase = m.M
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1708))
	if v3060|base.B2i32(v3061 == v3052) == v3052 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+704)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+712)) = int32(1)
	v3071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106)+1432)))
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+708)) = v3071
	v3074 = v1106 + int32(716)
	v3075 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3074)+16)) = v3075
	v3077 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3074)+8)) = v3077
	*(*int64)(unsafe.Add(mBase, uint32(v3074))) = v3077
	v3084 = F_pg_getaddrinfo_all(m, v3026, v3075, v3057, v1106+int32(1704))
	mBase = m.M
	if v3084 != 0 {
		v3919 = v1103
		goto L627
	} else {
		goto L628
	}
L625:
	;
	v3955 = v3061
	v3961 = v1103
	goto L626
L626:
	;
	if v3955 != 0 {
		goto L721
	} else {
		goto L722
	}
L627:
	;
	v3945 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1708))
	v3955 = v3945
	v3961 = v3919
	goto L626
L628:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1704))
	if v3085 == int32(0) {
		v3919 = v1103
		goto L627
	} else {
		goto L629
	}
L629:
	;
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1708))
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3088)+4))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3088)+8))
	v3091 = *(*int32)(unsafe.Add(mBase, uint32(v3088)+12))
	v3092 = F_socket(m, v3089, v3090, v3091)
	mBase = m.M
	if v3092 == int32(-1) {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	v3097 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3098 = m.ExcPending
	if v3098 != 0 {
		goto L1
	} else {
		goto L633
	}
L631:
	;
	goto L632
L632:
	;
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1704))
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v3113)+20))
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v3113)+16))
	v3116 = F_bind(m, v3092, v3114, v3115)
	mBase = m.M
	if v3116 != 0 {
		goto L640
	} else {
		goto L641
	}
L633:
	;
	if v3097 == int32(0) {
		v3919 = v1103
		goto L627
	} else {
		goto L634
	}
L634:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3102 = m.ExcPending
	if v3102 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_63), int32(0))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(1742), int32(_a_F_InitPostgres_64))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	v3919 = v1103
	goto L627
L638:
	;
	v3903 = F_close(m, v3092)
	mBase = m.M
	v3919 = v3877
	goto L627
L639:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), v3858, int32(_a_F_InitPostgres_64))
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L1
	} else {
		goto L720
	}
L640:
	;
	v3119 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3120 = m.ExcPending
	if v3120 != 0 {
		goto L1
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+388)) = v1106 + int32(848)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+384)) = v1106 + int32(1136)
	v3149 = F_pg_snprintf(m, v1106+int32(736), int32(80), int32(_a_F_InitPostgres_65), v1106+int32(384))
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L1
	} else {
		goto L647
	}
L643:
	;
	if v3119 == int32(0) {
		v3877 = v1103
		goto L638
	} else {
		goto L644
	}
L644:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+416)) = v1106 + int32(880)
	F_errmsg(m, int32(_a_F_InitPostgres_66), v1106+int32(416))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	v3832 = v1103
	v3858 = int32(1758)
	goto L639
L647:
	;
	goto L649
L648:
	;
	v3304 = v1106 + int32(2752)
	v3306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3304+v3251))) = uint8(v3306)
	v3310 = v1106 + int32(1712)
	v3312 = m.G0
	v3314 = v3312 - int32(80)
	m.G0 = v3314
	v3316 = F_strlen(m, v3304)
	mBase = m.M
	if base.Ui32(v3316) < base.Ui32(int32(2)) {
		v3767 = v3306
		goto L677
	} else {
		goto L678
	}
L649:
	;
	v3193 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v3193 != 0 {
		goto L651
	} else {
		goto L652
	}
L650:
	;
	v3285 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L1
	} else {
		goto L673
	}
L651:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3195 = m.ExcPending
	if v3195 != 0 {
		goto L1
	} else {
		goto L654
	}
L652:
	;
	goto L653
L653:
	;
	v3197 = v1106 + int32(736)
	v3198 = F_strlen(m, v3197)
	mBase = m.M
	v3199 = F_pgl_send(m, v3197, v3198)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L1
	} else {
		goto L655
	}
L654:
	;
	goto L653
L655:
	;
	if int32(0) <= v3199 {
		goto L656
	} else {
		goto L657
	}
L656:
	;
	goto L659
L657:
	;
	goto L658
L658:
	;
	v3280 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	if v3280 == int32(27) {
		goto L649
	} else {
		goto L672
	}
L659:
	;
	v3245 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v3245 != 0 {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	v3261 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3262 = m.ExcPending
	if v3262 != 0 {
		goto L1
	} else {
		goto L668
	}
L661:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L1
	} else {
		goto L664
	}
L662:
	;
	goto L663
L663:
	;
	v3251 = F_pgl_recv(m, v1106+int32(2752), int32(591))
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L1
	} else {
		goto L665
	}
L664:
	;
	goto L663
L665:
	;
	if int32(0) <= v3251 {
		goto L648
	} else {
		goto L666
	}
L666:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	if v3256 == int32(27) {
		goto L659
	} else {
		goto L667
	}
L667:
	;
	goto L660
L668:
	;
	if v3261 == int32(0) {
		v3877 = v1103
		goto L638
	} else {
		goto L669
	}
L669:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+356)) = v1106 + int32(816)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+352)) = v1106 + int32(1168)
	F_errmsg(m, int32(_a_F_InitPostgres_67), v1106+int32(352))
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	v3832 = v1103
	v3858 = int32(1809)
	goto L639
L672:
	;
	goto L650
L673:
	;
	if v3285 == int32(0) {
		v3877 = v1103
		goto L638
	} else {
		goto L674
	}
L674:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3290 = m.ExcPending
	if v3290 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+340)) = v1106 + int32(816)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+336)) = v1106 + int32(1168)
	F_errmsg(m, int32(_a_F_InitPostgres_68), v1106+int32(336))
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	v3832 = v1103
	v3858 = int32(1792)
	goto L639
L677:
	;
	m.G0 = v3314 + int32(80)
	if v3767 != 0 {
		v3877 = int32(1)
		goto L638
	} else {
		goto L716
	}
L678:
	;
	v3322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3304+v3316-int32(2)))))
	if v3322 != int32(13) {
		v3767 = v3306
		goto L677
	} else {
		goto L679
	}
L679:
	;
	v3332 = v3304
	goto L680
L680:
	;
	v3366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3332))))
	if v3366 == int32(13) {
		v3767 = v3306
		goto L677
	} else {
		goto L682
	}
L681:
	;
	v3380 = v3332
	goto L686
L682:
	;
	if v3366 != int32(58) {
		goto L683
	} else {
		goto L684
	}
L683:
	;
	v3332 = v3332 + int32(1)
	goto L680
L684:
	;
	goto L685
L685:
	;
	goto L681
L686:
	;
	v3415 = v3380 + int32(1)
	v3416 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3380)+1)))
	goto L688
L687:
	;
	v3432 = v3415
	v3433 = int32(0)
	goto L690
L688:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3416))%64)))&base.B2i32(base.Ui32(v3416) < base.Ui32(int32(33))) != 0 {
		v3380 = v3415
		goto L686
	} else {
		goto L689
	}
L689:
	;
	goto L687
L690:
	;
	v3466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3432))))
	if base.B2i32(v3466 == int32(13))|base.B2i32(v3466 == int32(58)) != 0 {
		goto L692
	} else {
		goto L693
	}
L691:
	;
	v3491 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3433+v3314))) = uint8(v3491)
	v3500 = v3432
	goto L696
L692:
	;
	goto L691
L693:
	;
	v3472 = base.I32_extend8_s(v3466)
	goto L694
L694:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3472))%64)))&base.B2i32(base.Ui32(v3472) < base.Ui32(int32(33)))|base.B2i32(base.Ui32(int32(78)) < base.Ui32(v3433)) != 0 {
		goto L692
	} else {
		goto L695
	}
L695:
	;
	v3484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3432))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3433+v3314))) = uint8(v3484)
	v3486 = int32(1)
	v3432 = v3432 + v3486
	v3433 = v3433 + v3486
	goto L690
L696:
	;
	v3536 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3500))))
	goto L698
L697:
	;
	v3544 = int32(0)
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v3314)))
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3314)+3))
	if v3545^int32(1380275029)|(v3548^int32(_a_F_InitPostgres_69)) != 0 {
		v3767 = v3544
		goto L677
	} else {
		goto L700
	}
L698:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3536))%64)))&base.B2i32(base.Ui32(v3536) < base.Ui32(int32(33))) != 0 {
		v3500 = v3500 + int32(1)
		goto L696
	} else {
		goto L699
	}
L699:
	;
	goto L697
L700:
	;
	v3552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3500))))
	if v3552 != int32(58) {
		v3767 = v3544
		goto L677
	} else {
		goto L701
	}
L701:
	;
	v3563 = v3500
	goto L702
L702:
	;
	v3596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3563)+1)))
	if v3596 == int32(13) {
		v3767 = v3544
		goto L677
	} else {
		goto L704
	}
L703:
	;
	v3613 = v3563 + int32(2)
	goto L706
L704:
	;
	if v3596 != int32(58) {
		v3563 = v3563 + int32(1)
		goto L702
	} else {
		goto L705
	}
L705:
	;
	goto L703
L706:
	;
	v3648 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3613))))
	goto L708
L707:
	;
	v3656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3613))))
	if v3656 == int32(13) {
		v3722 = v3544
		goto L710
	} else {
		goto L711
	}
L708:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3648))%64)))&base.B2i32(base.Ui32(v3648) < base.Ui32(int32(33))) != 0 {
		v3613 = v3613 + int32(1)
		goto L706
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	v3754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3722+v3310))) = uint8(v3754)
	v3767 = int32(1)
	goto L677
L711:
	;
	v3667 = v3613
	v3668 = int32(0)
	v3669 = v3656
	goto L712
L712:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3668+v3310))) = uint8(v3669)
	v3704 = v3668 + int32(1)
	v3705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3667)+1)))
	if v3705 == int32(13) {
		v3722 = v3704
		goto L710
	} else {
		goto L714
	}
L713:
	;
	v3722 = v3704
	goto L710
L714:
	;
	if base.Ui32(v3668) < base.Ui32(int32(511)) {
		v3667 = v3667 + int32(1)
		v3668 = v3704
		v3669 = v3705
		goto L712
	} else {
		goto L715
	}
L715:
	;
	goto L713
L716:
	;
	v3801 = int32(0)
	v3804 = F_errstart(m, int32(15), v3801)
	mBase = m.M
	v3805 = m.ExcPending
	if v3805 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	if v3804 == int32(0) {
		v3877 = v3801
		goto L638
	} else {
		goto L718
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+368)) = v1106 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_70), v1106+int32(368))
	mBase = m.M
	v3815 = m.ExcPending
	if v3815 != 0 {
		goto L1
	} else {
		goto L719
	}
L719:
	;
	v3832 = v3801
	v3858 = int32(1819)
	goto L639
L720:
	;
	v3877 = v3832
	goto L638
L721:
	;
	v3987 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106)+1568)))
	if v3987 == int32(1) {
		goto L726
	} else {
		goto L727
	}
L722:
	;
	goto L723
L723:
	;
	v4003 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+1704))
	if v4003 != 0 {
		goto L734
	} else {
		goto L735
	}
L724:
	;
	goto L723
L725:
	;
	goto L724
L726:
	;
	if v3955 == int32(0) {
		goto L725
	} else {
		goto L729
	}
L727:
	;
	goto L728
L728:
	;
	if v3955 == int32(0) {
		goto L725
	} else {
		goto L733
	}
L729:
	;
	v3993 = v3955
	goto L730
L730:
	;
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(v3993)+28))
	v3995 = *(*int32)(unsafe.Add(mBase, uint32(v3993)+20))
	F_emscripten_builtin_free(m, v3995)
	mBase = m.M
	F_emscripten_builtin_free(m, v3993)
	mBase = m.M
	if v3994 != 0 {
		v3993 = v3994
		goto L730
	} else {
		goto L732
	}
L731:
	;
	goto L725
L732:
	;
	goto L731
L733:
	;
	F_freeaddrinfo(m, v3955)
	mBase = m.M
	goto L725
L734:
	;
	v4004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1106)+1432)))
	if v4004 == int32(1) {
		goto L739
	} else {
		goto L740
	}
L735:
	;
	goto L736
L736:
	;
	if v3961 == int32(0) {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v2609
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L747
	}
L737:
	;
	goto L736
L738:
	;
	goto L737
L739:
	;
	if v4003 == int32(0) {
		goto L738
	} else {
		goto L742
	}
L740:
	;
	goto L741
L741:
	;
	if v4003 == int32(0) {
		goto L738
	} else {
		goto L746
	}
L742:
	;
	v4010 = v4003
	goto L743
L743:
	;
	v4011 = *(*int32)(unsafe.Add(mBase, uint32(v4010)+28))
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v4010)+20))
	F_emscripten_builtin_free(m, v4012)
	mBase = m.M
	F_emscripten_builtin_free(m, v4010)
	mBase = m.M
	if v4011 != 0 {
		v4010 = v4011
		goto L743
	} else {
		goto L745
	}
L744:
	;
	goto L738
L745:
	;
	goto L744
L746:
	;
	F_freeaddrinfo(m, v4003)
	mBase = m.M
	goto L738
L747:
	;
	v4023 = v1106 + int32(1712)
	F_set_authn_id(m, v1075, v4023)
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L1
	} else {
		goto L748
	}
L748:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4027 = *(*int32)(unsafe.Add(mBase, uint32(v4026)+300))
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4029 = F_check_usermap(m, v4027, v4028, v4023)
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L1
	} else {
		goto L749
	}
L749:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v4029
	v6609 = v47
	v6618 = v7
	goto L501
L750:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v4043)+296))
	if base.B2i32(v4044 != int32(5))|base.B2i32(v4042 != int32(1)) == int32(0) {
		goto L757
	} else {
		goto L758
	}
L751:
	;
	if v4034 == int32(0) {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v4039 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[55]))
	v4042 = v4039
	goto L750
L753:
	;
	goto L754
L754:
	;
	v4040 = F_get_password_type(m, v4034)
	mBase = m.M
	v4041 = m.ExcPending
	if v4041 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	v4042 = v4040
	goto L750
L756:
	;
	if v4034 != 0 {
		goto L833
	} else {
		goto L834
	}
L757:
	;
	v4055 = int32(0)
	v4059 = m.G0
	v4061 = v4059 - int32(16)
	m.G0 = v4061
	*(*int32)(unsafe.Add(mBase, uint32(v4061))) = v4055
	v4067 = F_open(m, int32(_a_F_InitPostgres_71), v4055, v4061)
	mBase = m.M
	if v4067 != int32(-1) {
		goto L761
	} else {
		goto L762
	}
L758:
	;
	goto L759
L759:
	;
	v4328 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_72), v1075, v4034, v1106+int32(700))
	mBase = m.M
	v4329 = m.ExcPending
	if v4329 != 0 {
		goto L1
	} else {
		goto L832
	}
L760:
	;
	if v4100 == int32(0) {
		goto L773
	} else {
		goto L774
	}
L761:
	;
	goto L765
L762:
	;
	v4100 = v4055
	goto L763
L763:
	;
	m.G0 = v4061 + int32(16)
	goto L760
L764:
	;
	v4095 = F_close(m, v4067)
	mBase = m.M
	v4100 = v4093
	goto L763
L765:
	;
	v4073 = v1106 + int32(2752)
	v4074 = int32(4)
	goto L766
L766:
	;
	v4079 = F_read(m, v4067, v4073, v4074)
	mBase = m.M
	if v4079 <= int32(0) {
		goto L768
	} else {
		goto L769
	}
L767:
	;
	v4093 = int32(1)
	goto L764
L768:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	if v4083 == int32(27) {
		goto L766
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	v4088 = v4074 - v4079
	if v4088 != 0 {
		v4073 = v4073 + v4079
		v4074 = v4088
		goto L766
	} else {
		goto L772
	}
L771:
	;
	v4093 = int32(0)
	goto L764
L772:
	;
	goto L767
L773:
	;
	v4109 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4110 = m.ExcPending
	if v4110 != 0 {
		goto L1
	} else {
		goto L776
	}
L774:
	;
	goto L775
L775:
	;
	F_sendAuthRequest(m, int32(5), v1106+int32(2752), int32(4))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L1
	} else {
		goto L780
	}
L776:
	;
	if v4109 == int32(0) {
		v4333 = v2609
		goto L756
	} else {
		goto L777
	}
L777:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_73), int32(0))
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(893), int32(_a_F_InitPostgres_74))
	mBase = m.M
	v4121 = m.ExcPending
	if v4121 != 0 {
		goto L1
	} else {
		goto L779
	}
L779:
	;
	v4333 = v2609
	goto L756
L780:
	;
	v4128 = F_recv_password_packet(m)
	mBase = m.M
	v4129 = m.ExcPending
	if v4129 != 0 {
		goto L1
	} else {
		goto L781
	}
L781:
	;
	if v4128 == int32(0) {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v4333 = int32(-2)
	goto L756
L783:
	;
	goto L784
L784:
	;
	if v4034 == int32(0) {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	F_pfree(m, v4128)
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L1
	} else {
		goto L788
	}
L786:
	;
	goto L787
L787:
	;
	v4137 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4140 = m.G0
	v4142 = v4140 - int32(128)
	m.G0 = v4142
	v4144 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4142)+28)) = v4144
	*(*int32)(unsafe.Add(mBase, uint32(v4142)+116)) = v4144
	v4148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4034))))
	if v4148 != int32(109) {
		goto L791
	} else {
		goto L792
	}
L788:
	;
	v4333 = v2609
	goto L756
L789:
	;
	m.G0 = v4142 + int32(128)
	F_pfree(m, v4128)
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L1
	} else {
		goto L831
	}
L790:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+700)) = v4315
	v4319 = int32(-1)
	goto L789
L791:
	;
	v4306 = F_parse_scram_secret(m, v4034, v4142+int32(120), v4142+int32(112), v4142+int32(116), v4142+int32(124), v4142+int32(32), v4142+int32(80))
	mBase = m.M
	v4307 = m.ExcPending
	if v4307 != 0 {
		goto L1
	} else {
		goto L829
	}
L792:
	;
	v4151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4034)+1)))
	if v4151 != int32(100) {
		goto L791
	} else {
		goto L793
	}
L793:
	;
	v4154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4034)+2)))
	if v4154 != int32(53) {
		goto L791
	} else {
		goto L794
	}
L794:
	;
	v4157 = F_strlen(m, v4034)
	mBase = m.M
	if v4157 != int32(35) {
		goto L791
	} else {
		goto L795
	}
L795:
	;
	v4161 = v4034 + int32(3)
	v4162 = int32(_a_F_InitPostgres_75)
	v4166 = m.G0
	v4168 = v4166 - int32(32)
	v4169 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4168)+24)) = v4169
	*(*int64)(unsafe.Add(mBase, uint32(v4168)+16)) = v4169
	*(*int64)(unsafe.Add(mBase, uint32(v4168)+8)) = v4169
	*(*int64)(unsafe.Add(mBase, uint32(v4168))) = v4169
	v4177 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[56])))
	if v4177 == int32(0) {
		goto L797
	} else {
		goto L798
	}
L796:
	;
	if v4245 != int32(32) {
		goto L791
	} else {
		goto L815
	}
L797:
	;
	v4245 = int32(0)
	goto L796
L798:
	;
	goto L799
L799:
	;
	v4181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[57])))
	if v4181 == int32(0) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	v4185 = v4161
	goto L803
L801:
	;
	goto L802
L802:
	;
	v4195 = v4162
	v4196 = v4177
	goto L806
L803:
	;
	v4191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4185))))
	if v4191 == v4177 {
		v4185 = v4185 + int32(1)
		goto L803
	} else {
		goto L805
	}
L804:
	;
	v4245 = v4185 - v4161
	goto L796
L805:
	;
	goto L804
L806:
	;
	v4203 = v4168 + int32(base.Ui32(v4196)>>(uint(int32(3))%32))&int32(28)
	v4204 = *(*int32)(unsafe.Add(mBase, uint32(v4203)))
	v4205 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4203))) = v4204 | v4205<<(uint(v4196)%32)
	v4209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4195)+1)))
	if v4209 != 0 {
		v4195 = v4195 + v4205
		v4196 = v4209
		goto L806
	} else {
		goto L808
	}
L807:
	;
	v4212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4161))))
	if v4212 == int32(0) {
		v4235 = v4161
		goto L809
	} else {
		goto L810
	}
L808:
	;
	goto L807
L809:
	;
	v4245 = v4235 - v4161
	goto L796
L810:
	;
	v4216 = v4161
	v4217 = v4212
	goto L811
L811:
	;
	v4225 = *(*int32)(unsafe.Add(mBase, uint32(v4168+int32(base.Ui32(v4217)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v4225)>>(uint(v4217)%32))&int32(1) == int32(0) {
		v4235 = v4216
		goto L809
	} else {
		goto L813
	}
L812:
	;
	v4235 = v4233
	goto L809
L813:
	;
	v4231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4216)+1)))
	v4233 = v4216 + int32(1)
	if v4231 != 0 {
		v4216 = v4233
		v4217 = v4231
		goto L811
	} else {
		goto L814
	}
L814:
	;
	goto L812
L815:
	;
	v4253 = F_pg_md5_encrypt(m, v4161, v1106+int32(2752), int32(4), v4142+int32(32), v4142+int32(28))
	mBase = m.M
	v4254 = m.ExcPending
	if v4254 != 0 {
		goto L1
	} else {
		goto L816
	}
L816:
	;
	if v4253 == int32(0) {
		goto L817
	} else {
		goto L818
	}
L817:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, uint32(v4142)+28))
	v4315 = v4257
	goto L790
L818:
	;
	goto L819
L819:
	;
	v4258 = int32(0)
	v4260 = v4142 + int32(32)
	v4263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4128))))
	v4266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4260))))
	if base.B2i32(v4263 == v4258)|base.B2i32(v4263 != v4266) != 0 {
		v4284 = v4263
		v4285 = v4266
		goto L821
	} else {
		goto L822
	}
L820:
	;
	if v4284-v4285 == int32(0) {
		v4319 = v4258
		goto L789
	} else {
		goto L827
	}
L821:
	;
	goto L820
L822:
	;
	v4269 = v4128
	v4270 = v4260
	goto L823
L823:
	;
	v4273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4270)+1)))
	v4274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4269)+1)))
	if v4274 == int32(0) {
		v4284 = v4274
		v4285 = v4273
		goto L821
	} else {
		goto L825
	}
L824:
	;
	v4284 = v4274
	v4285 = v4273
	goto L821
L825:
	;
	v4277 = int32(1)
	if v4274 == v4273 {
		v4269 = v4269 + v4277
		v4270 = v4270 + v4277
		goto L823
	} else {
		goto L826
	}
L826:
	;
	goto L824
L827:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4142))) = v4137
	v4291 = F_psprintf(m, int32(_a_F_InitPostgres_76), v4142)
	mBase = m.M
	v4292 = m.ExcPending
	if v4292 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	v4315 = v4291
	goto L790
L829:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4142)+16)) = v4137
	v4312 = F_psprintf(m, int32(_a_F_InitPostgres_77), v4142+int32(16))
	mBase = m.M
	v4313 = m.ExcPending
	if v4313 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	v4315 = v4312
	goto L790
L831:
	;
	v4333 = v4319
	goto L756
L832:
	;
	v4333 = v4328
	goto L756
L833:
	;
	F_pfree(m, v4034)
	mBase = m.M
	v4336 = m.ExcPending
	if v4336 != 0 {
		goto L1
	} else {
		goto L836
	}
L834:
	;
	goto L835
L835:
	;
	if v4333 != 0 {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v4333
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L837
	}
L836:
	;
	goto L835
L837:
	;
	v4337 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	F_set_authn_id(m, v1075, v4337)
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L1
	} else {
		goto L838
	}
L838:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = int32(0)
	v6609 = v47
	v6618 = v7
	goto L501
L839:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L1
	} else {
		goto L842
	}
L840:
	;
	goto L841
L841:
	;
	v4346 = v1106 + int32(2752)
	F_pq_beginmessage(m, v4346, int32(82))
	mBase = m.M
	v4349 = m.ExcPending
	if v4349 != 0 {
		goto L1
	} else {
		goto L843
	}
L842:
	;
	goto L841
L843:
	;
	F_enlargeStringInfo(m, v4346, int32(4))
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	v4353 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+2756))
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v4353+v4354))) = int32(50331648)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+2756)) = v4353 + int32(4)
	F_pq_endmessage(m, v4346)
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L845
	}
L845:
	;
	v4364 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[58]))
	v4365 = *(*int32)(unsafe.Add(mBase, uint32(v4364)+4))
	v4366 = m.T0[v4365].(func(*base.Module) int32)(m)
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L846
	}
L846:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v4369 != 0 {
		goto L847
	} else {
		goto L848
	}
L847:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4371 = m.ExcPending
	if v4371 != 0 {
		goto L1
	} else {
		goto L850
	}
L848:
	;
	goto L849
L849:
	;
	v4372 = F_recv_password_packet(m)
	mBase = m.M
	v4373 = m.ExcPending
	if v4373 != 0 {
		goto L1
	} else {
		goto L851
	}
L850:
	;
	goto L849
L851:
	;
	if v4372 == int32(0) {
		goto L852
	} else {
		goto L853
	}
L852:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = int32(-2)
	v6609 = v47
	v6618 = v7
	goto L501
L853:
	;
	goto L854
L854:
	;
	v4377 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4380 = F_get_role_password(m, v4377, v1106+int32(700))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L1
	} else {
		goto L855
	}
L855:
	;
	if v4380 == int32(0) {
		goto L856
	} else {
		goto L857
	}
L856:
	;
	F_pfree(m, v4372)
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L1
	} else {
		goto L859
	}
L857:
	;
	goto L858
L858:
	;
	v4386 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4389 = F_plain_crypt_verify(m, v4386, v4380, v4372, v1106+int32(700))
	mBase = m.M
	v4390 = m.ExcPending
	if v4390 != 0 {
		goto L1
	} else {
		goto L860
	}
L859:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v2609
	v6609 = v47
	v6618 = v7
	goto L501
L860:
	;
	F_pfree(m, v4380)
	mBase = m.M
	v4392 = m.ExcPending
	if v4392 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	F_pfree(m, v4372)
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	if v4389 != 0 {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v4389
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L863
	}
L863:
	;
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	F_set_authn_id(m, v1075, v4395)
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = int32(0)
	v6609 = v47
	v6618 = v7
	goto L501
L865:
	;
	v4404 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L1
	} else {
		goto L868
	}
L866:
	;
	goto L867
L867:
	;
	v4417 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+380))
	if v4417 == int32(0) {
		goto L872
	} else {
		goto L873
	}
L868:
	;
	if v4404 == int32(0) {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v2609
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L869
	}
L869:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_78), int32(0))
	mBase = m.M
	v4411 = m.ExcPending
	if v4411 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2860), int32(_a_F_InitPostgres_79))
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L1
	} else {
		goto L871
	}
L871:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v2609
	v6609 = v47
	v6618 = v7
	goto L501
L872:
	;
	v4422 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L875
	}
L873:
	;
	goto L874
L874:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v4436 != 0 {
		goto L879
	} else {
		goto L880
	}
L875:
	;
	if v4422 == int32(0) {
		v6585 = l0
		v6586 = l1
		v6587 = l2
		v6588 = l3
		v6589 = l4
		v6590 = l5
		v6591 = v1106
		v6596 = v1075
		v6597 = v44
		v6598 = v2609
		v6609 = v47
		v6618 = v7
		goto L501
	} else {
		goto L876
	}
L876:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_80), int32(0))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L1
	} else {
		goto L877
	}
L877:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2867), int32(_a_F_InitPostgres_79))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L1
	} else {
		goto L878
	}
L878:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v2609
	v6609 = v47
	v6618 = v7
	goto L501
L879:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4438 = m.ExcPending
	if v4438 != 0 {
		goto L1
	} else {
		goto L882
	}
L880:
	;
	goto L881
L881:
	;
	v4440 = v1106 + int32(2752)
	F_pq_beginmessage(m, v4440, int32(82))
	mBase = m.M
	v4443 = m.ExcPending
	if v4443 != 0 {
		goto L1
	} else {
		goto L883
	}
L882:
	;
	goto L881
L883:
	;
	F_enlargeStringInfo(m, v4440, int32(4))
	mBase = m.M
	v4446 = m.ExcPending
	if v4446 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	v4447 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+2756))
	v4448 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v4447+v4448))) = int32(50331648)
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+2756)) = v4447 + int32(4)
	F_pq_endmessage(m, v4440)
	mBase = m.M
	v4456 = m.ExcPending
	if v4456 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	v4458 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[58]))
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(v4458)+4))
	v4460 = m.T0[v4459].(func(*base.Module) int32)(m)
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	v4463 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v4463 != 0 {
		goto L887
	} else {
		goto L888
	}
L887:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L890
	}
L888:
	;
	goto L889
L889:
	;
	v4466 = F_recv_password_packet(m)
	mBase = m.M
	v4467 = m.ExcPending
	if v4467 != 0 {
		goto L1
	} else {
		goto L891
	}
L890:
	;
	goto L889
L891:
	;
	if v4466 == int32(0) {
		goto L892
	} else {
		goto L893
	}
L892:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = int32(-2)
	v6609 = v47
	v6618 = v7
	goto L501
L893:
	;
	goto L894
L894:
	;
	v4471 = F_strlen(m, v4466)
	mBase = m.M
	if base.Ui32(int32(129)) <= base.Ui32(v4471) {
		goto L895
	} else {
		goto L896
	}
L895:
	;
	v4476 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4477 = m.ExcPending
	if v4477 != 0 {
		goto L1
	} else {
		goto L898
	}
L896:
	;
	goto L897
L897:
	;
	v4492 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v4492)+380))
	if v4493 != 0 {
		goto L905
	} else {
		goto L906
	}
L898:
	;
	if v4476 != 0 {
		goto L899
	} else {
		goto L900
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1106)+448)) = int32(128)
	F_errmsg(m, int32(_a_F_InitPostgres_81), v1106+int32(448))
	mBase = m.M
	v4484 = m.ExcPending
	if v4484 != 0 {
		goto L1
	} else {
		goto L902
	}
L900:
	;
	goto L901
L901:
	;
	F_pfree(m, v4466)
	mBase = m.M
	v4491 = m.ExcPending
	if v4491 != 0 {
		goto L1
	} else {
		goto L904
	}
L902:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2881), int32(_a_F_InitPostgres_79))
	mBase = m.M
	v4489 = m.ExcPending
	if v4489 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	goto L901
L904:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v2609
	v6609 = v47
	v6618 = v7
	goto L501
L905:
	;
	v4494 = *(*int32)(unsafe.Add(mBase, uint32(v4493)+12))
	v4495 = v4494
	goto L907
L906:
	;
	v4495 = v7
	goto L907
L907:
	;
	v4496 = *(*int32)(unsafe.Add(mBase, uint32(v4492)+396))
	if v4496 != 0 {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v4497 = *(*int32)(unsafe.Add(mBase, uint32(v4496)+12))
	v4498 = v4497
	goto L910
L909:
	;
	v4498 = v7
	goto L910
L910:
	;
	v4499 = *(*int32)(unsafe.Add(mBase, uint32(v4492)+388))
	if v4499 != 0 {
		goto L911
	} else {
		goto L912
	}
L911:
	;
	v4500 = *(*int32)(unsafe.Add(mBase, uint32(v4499)+12))
	v4502 = v4500
	goto L913
L912:
	;
	v4502 = int32(0)
	goto L913
L913:
	;
	v4503 = *(*int32)(unsafe.Add(mBase, uint32(v4492)+372))
	if v4503 == int32(0) {
		v6504 = l0
		v6505 = l1
		v6506 = l2
		v6507 = l3
		v6508 = l4
		v6509 = l5
		v6510 = v1106
		v6515 = v1075
		v6516 = v44
		v6517 = v2609
		v6522 = v4466
		v6528 = v47
		v6537 = v7
		goto L914
	} else {
		goto L915
	}
L914:
	;
	F_pfree(m, v6522)
	mBase = m.M
	v6546 = m.ExcPending
	if v6546 != 0 {
		goto L1
	} else {
		goto L1338
	}
L915:
	;
	v4506 = *(*int32)(unsafe.Add(mBase, uint32(v4503)+4))
	if v4506 <= int32(0) {
		v6504 = l0
		v6505 = l1
		v6506 = l2
		v6507 = l3
		v6508 = l4
		v6509 = l5
		v6510 = v1106
		v6515 = v1075
		v6516 = v44
		v6517 = v2609
		v6522 = v4466
		v6528 = v47
		v6537 = v7
		goto L914
	} else {
		goto L916
	}
L916:
	;
	v4517 = l0
	v4518 = l1
	v4519 = l2
	v4520 = l3
	v4521 = l4
	v4522 = l5
	v4523 = v1106
	v4528 = v1075
	v4529 = v44
	v4530 = v2609
	v4531 = v1106 + int32(2756)
	v4535 = v4466
	v4537 = v4502
	v4540 = v4503
	v4541 = v47
	v4542 = v4498
	v4544 = v4495
	v4548 = v1106 + int32(1440)
	v4549 = v7
	v4550 = v7
	v4551 = v1106 + int32(1732)
	v4552 = v1106 + int32(1724)
	goto L917
L917:
	;
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v4540)+12))
	if v4542 != 0 {
		goto L919
	} else {
		goto L920
	}
L918:
	;
	v6504 = v6404
	v6505 = v6405
	v6506 = v6406
	v6507 = v6407
	v6508 = v6408
	v6509 = v6409
	v6510 = v6410
	v6515 = v6415
	v6516 = v6416
	v6517 = v6417
	v6522 = v6422
	v6528 = v6428
	v6537 = v6437
	goto L914
L919:
	;
	v4563 = *(*int32)(unsafe.Add(mBase, uint32(v4542)))
	v4564 = v4563
	goto L921
L920:
	;
	v4564 = int32(0)
	goto L921
L921:
	;
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v4558+v4549<<(uint(int32(2))%32))))
	if v4537 != 0 {
		goto L922
	} else {
		goto L923
	}
L922:
	;
	v4566 = *(*int32)(unsafe.Add(mBase, uint32(v4537)))
	v4568 = v4566
	goto L924
L923:
	;
	v4568 = int32(0)
	goto L924
L924:
	;
	v4569 = *(*int32)(unsafe.Add(mBase, uint32(v4544)))
	v4570 = *(*int32)(unsafe.Add(mBase, uint32(v4528)+364))
	v4571 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1576)) = v4571
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1568)) = v4571
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1584)) = v4571
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1592)) = v4571
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1576)) = int32(2)
	if v4564 != 0 {
		goto L925
	} else {
		goto L926
	}
L925:
	;
	v4582 = v4564
	goto L927
L926:
	;
	v4582 = int32(_a_F_InitPostgres_82)
	goto L927
L927:
	;
	v4586 = v4582
	goto L929
L928:
	;
	v4635 = F_pg_getaddrinfo_all(m, v4565, v4582, v4523+int32(1568), v4523+int32(704))
	mBase = m.M
	if v4635 == int32(0) {
		goto L946
	} else {
		goto L947
	}
L929:
	;
	v4591 = v4586 + int32(1)
	v4592 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4586))))
	v4593 = F___isspace(m, v4592)
	mBase = m.M
	if v4593 != 0 {
		v4586 = v4591
		goto L929
	} else {
		goto L931
	}
L930:
	;
	v4594 = int32(1)
	switch v4592&int32(255) - int32(43) {
	case 0:
		v4600 = v4594
		goto L933
	default:
		v4602 = v4592
		v4603 = v4586
		v4604 = v4594
		goto L932
	case 2:
		goto L934
	}
L931:
	;
	goto L930
L932:
	;
	v4605 = int32(0)
	v4607 = v4602 - int32(48)
	if base.Ui32(v4607) <= base.Ui32(int32(9)) {
		goto L935
	} else {
		goto L936
	}
L933:
	;
	v4601 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4591))))
	v4602 = v4601
	v4603 = v4591
	v4604 = v4600
	goto L932
L934:
	;
	v4600 = int32(0)
	goto L933
L935:
	;
	v4610 = v4605
	v4611 = v4607
	v4612 = v4603
	goto L938
L936:
	;
	v4624 = v4605
	goto L937
L937:
	;
	if v4604 != 0 {
		goto L941
	} else {
		goto L942
	}
L938:
	;
	v4614 = int32(10)
	v4616 = v4610*v4614 - v4611
	v4617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4612)+1)))
	v4621 = v4617 - int32(48)
	if base.Ui32(v4621) < base.Ui32(v4614) {
		v4610 = v4616
		v4611 = v4621
		v4612 = v4612 + int32(1)
		goto L938
	} else {
		goto L940
	}
L939:
	;
	v4624 = v4616
	goto L937
L940:
	;
	goto L939
L941:
	;
	v4630 = int32(0) - v4624
	goto L943
L942:
	;
	v4630 = v4624
	goto L943
L943:
	;
	goto L928
L944:
	;
	v6445 = *(*int32)(unsafe.Add(mBase, uint32(v6415)+380))
	v6446 = *(*int32)(unsafe.Add(mBase, uint32(v6445)+380))
	if v6446 == int32(0) {
		v6463 = v6431
		goto L1319
	} else {
		goto L1320
	}
L945:
	;
	v4704 = int32(20)
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v4704)
	v4706 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4523)+2752)) = uint8(v4706)
	v4708 = int32(16)
	v4709 = int32(0)
	v4713 = m.G0
	v4715 = v4713 - v4708
	m.G0 = v4715
	*(*int32)(unsafe.Add(mBase, uint32(v4715))) = v4709
	v4721 = F_open(m, int32(_a_F_InitPostgres_71), v4709, v4715)
	mBase = m.M
	if v4721 != int32(-1) {
		goto L978
	} else {
		goto L979
	}
L946:
	;
	v4638 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v4638 != 0 {
		goto L945
	} else {
		goto L949
	}
L947:
	;
	goto L948
L948:
	;
	v4641 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		goto L1
	} else {
		goto L950
	}
L949:
	;
	goto L948
L950:
	;
	if v4641 != 0 {
		goto L951
	} else {
		goto L952
	}
L951:
	;
	v4645 = int32(_a_F_InitPostgres_54)
	v4647 = v4635 + int32(1)
	if v4647 == int32(0) {
		v4667 = v4645
		goto L955
	} else {
		goto L956
	}
L952:
	;
	goto L953
L953:
	;
	v4685 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v4685 == int32(0) {
		v6404 = v4517
		v6405 = v4518
		v6406 = v4519
		v6407 = v4520
		v6408 = v4521
		v6409 = v4522
		v6410 = v4523
		v6415 = v4528
		v6416 = v4529
		v6417 = v4530
		v6418 = v4531
		v6422 = v4535
		v6424 = v4537
		v6427 = v4540
		v6428 = v4541
		v6429 = v4542
		v6431 = v4544
		v6435 = v4548
		v6436 = v4549
		v6437 = v4550
		v6438 = v4551
		v6439 = v4552
		goto L944
	} else {
		goto L966
	}
L954:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+692)) = v4667 + base.B2i32(v4669 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+688)) = v4565
	F_errmsg(m, int32(_a_F_InitPostgres_83), v4523+int32(688))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L1
	} else {
		goto L964
	}
L955:
	;
	v4669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4667))))
	goto L954
L956:
	;
	v4651 = v4645
	v4652 = v4647
	goto L957
L957:
	;
	v4653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4651))))
	if v4653 == int32(0) {
		v4667 = v4651
		goto L955
	} else {
		goto L959
	}
L958:
	;
	v4667 = v4663
	goto L955
L959:
	;
	v4657 = v4651
	goto L960
L960:
	;
	v4661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4657)+1)))
	if v4661 != 0 {
		v4657 = v4657 + int32(1)
		goto L960
	} else {
		goto L962
	}
L961:
	;
	v4663 = v4657 + int32(2)
	v4665 = v4652 + int32(1)
	if v4665 != 0 {
		v4651 = v4663
		v4652 = v4665
		goto L957
	} else {
		goto L963
	}
L962:
	;
	goto L961
L963:
	;
	goto L958
L964:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2984), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v4684 = m.ExcPending
	if v4684 != 0 {
		goto L1
	} else {
		goto L965
	}
L965:
	;
	goto L953
L966:
	;
	v4688 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	if v4688 == int32(1) {
		goto L969
	} else {
		goto L970
	}
L967:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L968:
	;
	goto L967
L969:
	;
	if v4685 == int32(0) {
		goto L968
	} else {
		goto L972
	}
L970:
	;
	goto L971
L971:
	;
	if v4685 == int32(0) {
		goto L968
	} else {
		goto L976
	}
L972:
	;
	v4694 = v4685
	goto L973
L973:
	;
	v4695 = *(*int32)(unsafe.Add(mBase, uint32(v4694)+28))
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v4694)+20))
	F_emscripten_builtin_free(m, v4696)
	mBase = m.M
	F_emscripten_builtin_free(m, v4694)
	mBase = m.M
	if v4695 != 0 {
		v4694 = v4695
		goto L973
	} else {
		goto L975
	}
L974:
	;
	goto L968
L975:
	;
	goto L974
L976:
	;
	F_freeaddrinfo(m, v4685)
	mBase = m.M
	goto L968
L977:
	;
	if v4754 == int32(0) {
		goto L990
	} else {
		goto L991
	}
L978:
	;
	goto L982
L979:
	;
	v4754 = v4709
	goto L980
L980:
	;
	m.G0 = v4715 + int32(16)
	goto L977
L981:
	;
	v4749 = F_close(m, v4721)
	mBase = m.M
	v4754 = v4747
	goto L980
L982:
	;
	v4727 = v4531
	v4728 = v4708
	goto L983
L983:
	;
	v4733 = F_read(m, v4721, v4727, v4728)
	mBase = m.M
	if v4733 <= int32(0) {
		goto L985
	} else {
		goto L986
	}
L984:
	;
	v4747 = int32(1)
	goto L981
L985:
	;
	v4737 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	if v4737 == int32(27) {
		goto L983
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	v4742 = v4728 - v4733
	if v4742 != 0 {
		v4727 = v4727 + v4733
		v4728 = v4742
		goto L983
	} else {
		goto L989
	}
L988:
	;
	v4747 = int32(0)
	goto L981
L989:
	;
	goto L984
L990:
	;
	v4763 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4764 = m.ExcPending
	if v4764 != 0 {
		goto L1
	} else {
		goto L993
	}
L991:
	;
	goto L992
L992:
	;
	v4791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+2756)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4523)+2753)) = uint8(v4791)
	v4793 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	if base.Ui32(int32(1021)) <= base.Ui32(v4793) {
		goto L1010
	} else {
		goto L1011
	}
L993:
	;
	if v4763 != 0 {
		goto L994
	} else {
		goto L995
	}
L994:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_85), int32(0))
	mBase = m.M
	v4768 = m.ExcPending
	if v4768 != 0 {
		goto L1
	} else {
		goto L997
	}
L995:
	;
	goto L996
L996:
	;
	v4774 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v4775 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v4774 == int32(1) {
		goto L1001
	} else {
		goto L1002
	}
L997:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2997), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v4773 = m.ExcPending
	if v4773 != 0 {
		goto L1
	} else {
		goto L998
	}
L998:
	;
	goto L996
L999:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L1000:
	;
	goto L999
L1001:
	;
	if v4775 == int32(0) {
		goto L1000
	} else {
		goto L1004
	}
L1002:
	;
	goto L1003
L1003:
	;
	if v4775 == int32(0) {
		goto L1000
	} else {
		goto L1008
	}
L1004:
	;
	v4781 = v4775
	goto L1005
L1005:
	;
	v4782 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+28))
	v4783 = *(*int32)(unsafe.Add(mBase, uint32(v4781)+20))
	F_emscripten_builtin_free(m, v4783)
	mBase = m.M
	F_emscripten_builtin_free(m, v4781)
	mBase = m.M
	if v4782 != 0 {
		v4781 = v4782
		goto L1005
	} else {
		goto L1007
	}
L1006:
	;
	goto L1000
L1007:
	;
	goto L1006
L1008:
	;
	F_freeaddrinfo(m, v4775)
	mBase = m.M
	goto L1000
L1009:
	;
	if v4568 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1010:
	;
	v4798 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1011:
	;
	goto L1012
L1012:
	;
	v4816 = v4523 + int32(2752) + v4793
	*(*int32)(unsafe.Add(mBase, uint32(v4816)+2)) = int32(134217728)
	v4819 = int32(1542)
	*(*uint16)(unsafe.Add(mBase, uint32(v4816))) = uint16(v4819)
	v4821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	v4823 = v4821 + int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v4823)
	goto L1009
L1013:
	;
	if v4798 == int32(0) {
		goto L1009
	} else {
		goto L1014
	}
L1014:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+672)) = int64(17179869190)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_86), v4523+int32(672))
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2833), int32(_a_F_InitPostgres_87))
	mBase = m.M
	v4813 = m.ExcPending
	if v4813 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	goto L1009
L1017:
	;
	v4827 = v4568
	goto L1019
L1018:
	;
	v4827 = int32(_a_F_InitPostgres_88)
	goto L1019
L1019:
	;
	v4828 = F_strlen(m, v4570)
	mBase = m.M
	v4829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	if int32(1025) <= v4828+v4829 {
		goto L1021
	} else {
		goto L1022
	}
L1020:
	;
	v4870 = F_strlen(m, v4827)
	mBase = m.M
	v4871 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	if int32(1025) <= v4870+v4871 {
		goto L1032
	} else {
		goto L1033
	}
L1021:
	;
	v4835 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4836 = m.ExcPending
	if v4836 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1022:
	;
	goto L1023
L1023:
	;
	v4854 = v4523 + int32(2752) + v4829
	v4856 = v4828 + int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4854)+1)) = uint8(v4856)
	v4858 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4854))) = uint8(v4858)
	if v4828 != 0 {
		goto L1028
	} else {
		goto L1029
	}
L1024:
	;
	if v4835 == int32(0) {
		goto L1020
	} else {
		goto L1025
	}
L1025:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+660)) = v4828
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+656)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_86), v4523+int32(656))
	mBase = m.M
	v4846 = m.ExcPending
	if v4846 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2833), int32(_a_F_InitPostgres_87))
	mBase = m.M
	v4851 = m.ExcPending
	if v4851 != 0 {
		goto L1
	} else {
		goto L1027
	}
L1027:
	;
	goto L1020
L1028:
	;
	base.MemoryCopy(m, v4854+int32(2), v4570, v4828)
	goto L1030
L1029:
	;
	goto L1030
L1030:
	;
	v4863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	v4866 = v4863 + v4856&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v4866)
	goto L1020
L1031:
	;
	v4912 = int32(16)
	v4913 = F_strlen(m, v4535)
	mBase = m.M
	v4914 = F_strlen(m, v4569)
	mBase = m.M
	v4917 = F_palloc(m, v4914+v4912)
	mBase = m.M
	v4918 = m.ExcPending
	if v4918 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1032:
	;
	v4877 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4878 = m.ExcPending
	if v4878 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1033:
	;
	goto L1034
L1034:
	;
	v4896 = v4523 + int32(2752) + v4871
	v4898 = v4870 + int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4896)+1)) = uint8(v4898)
	v4900 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v4896))) = uint8(v4900)
	if v4870 != 0 {
		goto L1039
	} else {
		goto L1040
	}
L1035:
	;
	if v4877 == int32(0) {
		goto L1031
	} else {
		goto L1036
	}
L1036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+644)) = v4870
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+640)) = int32(32)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_86), v4523+int32(640))
	mBase = m.M
	v4888 = m.ExcPending
	if v4888 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2833), int32(_a_F_InitPostgres_87))
	mBase = m.M
	v4893 = m.ExcPending
	if v4893 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	goto L1031
L1039:
	;
	base.MemoryCopy(m, v4896+int32(2), v4827, v4870)
	goto L1041
L1040:
	;
	goto L1041
L1041:
	;
	v4905 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	v4908 = v4905 + v4898&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v4908)
	goto L1031
L1042:
	;
	v4919 = F_strlen(m, v4569)
	mBase = m.M
	if v4919 != 0 {
		goto L1043
	} else {
		goto L1044
	}
L1043:
	;
	base.MemoryCopy(m, v4917, v4569, v4919)
	goto L1045
L1044:
	;
	goto L1045
L1045:
	;
	v4922 = v4913 + int32(15)
	v4924 = v4922 & int32(-16)
	if int32(16) <= v4922 {
		goto L1049
	} else {
		goto L1050
	}
L1046:
	;
	v5182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	v5183 = int32(8)
	v5187 = v5182<<(uint(v5183)%32) | int32(base.Ui32(v5182)>>(uint(v5183)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v5187)
	v5189 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	v5190 = *(*int32)(unsafe.Add(mBase, uint32(v5189)+4))
	v5193 = F_socket(m, v5190, int32(2), int32(0))
	mBase = m.M
	if v5193 == int32(-1) {
		goto L1092
	} else {
		goto L1093
	}
L1047:
	;
	v5164 = v4523 + int32(2752) + v5104
	v5165 = int32(2)
	v5166 = v4924 | v5165
	*(*uint8)(unsafe.Add(mBase, uint32(v5164)+1)) = uint8(v5166)
	*(*uint8)(unsafe.Add(mBase, uint32(v5164))) = uint8(v5165)
	if v4924 != 0 {
		goto L1089
	} else {
		goto L1090
	}
L1048:
	;
	v5129 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5130 = m.ExcPending
	if v5130 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1049:
	;
	v4935 = v4531
	v4936 = v4912
	v4938 = int32(0)
	goto L1052
L1050:
	;
	goto L1051
L1051:
	;
	F_pfree(m, v4917)
	mBase = m.M
	v5103 = m.ExcPending
	if v5103 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1052:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+848)) = int32(0)
	v4971 = F_strlen(m, v4569)
	mBase = m.M
	v4972 = v4971 + v4917
	v4973 = *(*int64)(unsafe.Add(mBase, uint32(v4935)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4972)+8)) = v4973
	v4975 = *(*int64)(unsafe.Add(mBase, uint32(v4935)))
	*(*int64)(unsafe.Add(mBase, uint32(v4972))) = v4975
	v4977 = F_strlen(m, v4569)
	mBase = m.M
	v4982 = v4523 + int32(1168) + v4938
	v4985 = F_pg_md5_binary(m, v4917, v4977+int32(16), v4982, v4523+int32(848))
	mBase = m.M
	v4986 = m.ExcPending
	if v4986 != 0 {
		goto L1
	} else {
		goto L1054
	}
L1053:
	;
	goto L1051
L1054:
	;
	if v4985 == int32(0) {
		goto L1048
	} else {
		goto L1055
	}
L1055:
	;
	v4989 = F_strlen(m, v4535)
	mBase = m.M
	v4999 = v4938
	goto L1056
L1056:
	;
	if base.Ui32(v4999) < base.Ui32(v4989) {
		goto L1058
	} else {
		goto L1059
	}
L1057:
	;
	v5056 = int32(16)
	v5059 = v4938 + v5056
	if v5059 < v4924 {
		v4935 = v4982
		v4936 = v4936 + v5056
		v4938 = v5059
		goto L1052
	} else {
		goto L1065
	}
L1058:
	;
	v5034 = v4523 + int32(1168) + v4999
	v5035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5034))))
	v5037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4999+v4535))))
	v5038 = v5035 ^ v5037
	*(*uint8)(unsafe.Add(mBase, uint32(v5034))) = uint8(v5038)
	goto L1060
L1059:
	;
	goto L1060
L1060:
	;
	v5042 = v4999 | int32(1)
	if base.Ui32(v5042) < base.Ui32(v4989) {
		goto L1061
	} else {
		goto L1062
	}
L1061:
	;
	v5046 = v4523 + int32(1168) + v5042
	v5047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5046))))
	v5049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4535+v5042))))
	v5050 = v5047 ^ v5049
	*(*uint8)(unsafe.Add(mBase, uint32(v5046))) = uint8(v5050)
	goto L1063
L1062:
	;
	goto L1063
L1063:
	;
	v5054 = v4999 + int32(2)
	if v5054 != v4936 {
		v4999 = v5054
		goto L1056
	} else {
		goto L1064
	}
L1064:
	;
	goto L1057
L1065:
	;
	goto L1053
L1066:
	;
	v5104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	if v4924+v5104 < int32(1025) {
		goto L1047
	} else {
		goto L1067
	}
L1067:
	;
	v5110 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5111 = m.ExcPending
	if v5111 != 0 {
		goto L1
	} else {
		goto L1068
	}
L1068:
	;
	if v5110 == int32(0) {
		goto L1046
	} else {
		goto L1069
	}
L1069:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+612)) = v4924
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+608)) = int32(2)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_86), v4523+int32(608))
	mBase = m.M
	v5121 = m.ExcPending
	if v5121 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1070:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(2833), int32(_a_F_InitPostgres_87))
	mBase = m.M
	v5126 = m.ExcPending
	if v5126 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1071:
	;
	goto L1046
L1072:
	;
	if v5129 != 0 {
		goto L1073
	} else {
		goto L1074
	}
L1073:
	;
	v5131 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+848))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+624)) = v5131
	F_errmsg(m, int32(_a_F_InitPostgres_89), v4523+int32(624))
	mBase = m.M
	v5137 = m.ExcPending
	if v5137 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1074:
	;
	goto L1075
L1075:
	;
	F_pfree(m, v4917)
	mBase = m.M
	v5144 = m.ExcPending
	if v5144 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1076:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3035), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v5142 = m.ExcPending
	if v5142 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	goto L1075
L1078:
	;
	v5145 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v5146 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v5145 == int32(1) {
		goto L1081
	} else {
		goto L1082
	}
L1079:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L1080:
	;
	goto L1079
L1081:
	;
	if v5146 == int32(0) {
		goto L1080
	} else {
		goto L1084
	}
L1082:
	;
	goto L1083
L1083:
	;
	if v5146 == int32(0) {
		goto L1080
	} else {
		goto L1088
	}
L1084:
	;
	v5152 = v5146
	goto L1085
L1085:
	;
	v5153 = *(*int32)(unsafe.Add(mBase, uint32(v5152)+28))
	v5154 = *(*int32)(unsafe.Add(mBase, uint32(v5152)+20))
	F_emscripten_builtin_free(m, v5154)
	mBase = m.M
	F_emscripten_builtin_free(m, v5152)
	mBase = m.M
	if v5153 != 0 {
		v5152 = v5153
		goto L1085
	} else {
		goto L1087
	}
L1086:
	;
	goto L1080
L1087:
	;
	goto L1086
L1088:
	;
	F_freeaddrinfo(m, v5146)
	mBase = m.M
	goto L1080
L1089:
	;
	base.MemoryCopy(m, v5164+int32(2), v4523+int32(1168), v4924)
	goto L1091
L1090:
	;
	goto L1091
L1091:
	;
	v5175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)))
	v5178 = v5175 + v5166&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+2754)) = uint16(v5178)
	goto L1046
L1092:
	;
	v5198 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5199 = m.ExcPending
	if v5199 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1093:
	;
	goto L1094
L1094:
	;
	v5226 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1448)) = v5226
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1440)) = v5226
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+1432)) = v5226
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1456)) = int32(0)
	v5234 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	v5235 = *(*int32)(unsafe.Add(mBase, uint32(v5234)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v4523)+1432)) = uint16(v5235)
	v5238 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[59]))
	*(*int64)(unsafe.Add(mBase, uint32(v4548)+8)) = v5238
	v5241 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[60]))
	*(*int64)(unsafe.Add(mBase, uint32(v4548))) = v5241
	if v5235&int32(_a_F_InitPostgres_90) == int32(10) {
		goto L1111
	} else {
		goto L1112
	}
L1095:
	;
	if v5198 != 0 {
		goto L1096
	} else {
		goto L1097
	}
L1096:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_91), int32(0))
	mBase = m.M
	v5203 = m.ExcPending
	if v5203 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1097:
	;
	goto L1098
L1098:
	;
	v5209 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v5210 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v5209 == int32(1) {
		goto L1103
	} else {
		goto L1104
	}
L1099:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3061), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v5208 = m.ExcPending
	if v5208 != 0 {
		goto L1
	} else {
		goto L1100
	}
L1100:
	;
	goto L1098
L1101:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L1102:
	;
	goto L1101
L1103:
	;
	if v5210 == int32(0) {
		goto L1102
	} else {
		goto L1106
	}
L1104:
	;
	goto L1105
L1105:
	;
	if v5210 == int32(0) {
		goto L1102
	} else {
		goto L1110
	}
L1106:
	;
	v5216 = v5210
	goto L1107
L1107:
	;
	v5217 = *(*int32)(unsafe.Add(mBase, uint32(v5216)+28))
	v5218 = *(*int32)(unsafe.Add(mBase, uint32(v5216)+20))
	F_emscripten_builtin_free(m, v5218)
	mBase = m.M
	F_emscripten_builtin_free(m, v5216)
	mBase = m.M
	if v5217 != 0 {
		v5216 = v5217
		goto L1107
	} else {
		goto L1109
	}
L1108:
	;
	goto L1102
L1109:
	;
	goto L1108
L1110:
	;
	F_freeaddrinfo(m, v5210)
	mBase = m.M
	goto L1102
L1111:
	;
	v5249 = int32(28)
	goto L1113
L1112:
	;
	v5249 = int32(16)
	goto L1113
L1113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1708)) = v5249
	v5253 = F_bind(m, v5193, v4523+int32(1432), v5249)
	mBase = m.M
	if v5253 != 0 {
		goto L1114
	} else {
		goto L1115
	}
L1114:
	;
	v5256 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5257 = m.ExcPending
	if v5257 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1115:
	;
	goto L1116
L1116:
	;
	v5287 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	v5288 = *(*int32)(unsafe.Add(mBase, uint32(v5287)+20))
	v5289 = *(*int32)(unsafe.Add(mBase, uint32(v5287)+16))
	v5290 = F_sendto(m, v5193, v4523+int32(2752), v5182, v5288, v5289)
	mBase = m.M
	if v5290 < int32(0) {
		goto L1133
	} else {
		goto L1134
	}
L1117:
	;
	if v5256 != 0 {
		goto L1118
	} else {
		goto L1119
	}
L1118:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_92), int32(0))
	mBase = m.M
	v5261 = m.ExcPending
	if v5261 != 0 {
		goto L1
	} else {
		goto L1121
	}
L1119:
	;
	goto L1120
L1120:
	;
	v5267 = F_close(m, v5193)
	mBase = m.M
	v5268 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v5269 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v5268 == int32(1) {
		goto L1125
	} else {
		goto L1126
	}
L1121:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3077), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v5266 = m.ExcPending
	if v5266 != 0 {
		goto L1
	} else {
		goto L1122
	}
L1122:
	;
	goto L1120
L1123:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L1124:
	;
	goto L1123
L1125:
	;
	if v5269 == int32(0) {
		goto L1124
	} else {
		goto L1128
	}
L1126:
	;
	goto L1127
L1127:
	;
	if v5269 == int32(0) {
		goto L1124
	} else {
		goto L1132
	}
L1128:
	;
	v5275 = v5269
	goto L1129
L1129:
	;
	v5276 = *(*int32)(unsafe.Add(mBase, uint32(v5275)+28))
	v5277 = *(*int32)(unsafe.Add(mBase, uint32(v5275)+20))
	F_emscripten_builtin_free(m, v5277)
	mBase = m.M
	F_emscripten_builtin_free(m, v5275)
	mBase = m.M
	if v5276 != 0 {
		v5275 = v5276
		goto L1129
	} else {
		goto L1131
	}
L1130:
	;
	goto L1124
L1131:
	;
	goto L1130
L1132:
	;
	F_freeaddrinfo(m, v5269)
	mBase = m.M
	goto L1124
L1133:
	;
	v5295 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5296 = m.ExcPending
	if v5296 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1134:
	;
	goto L1135
L1135:
	;
	v5324 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v5325 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v5324 == int32(1) {
		goto L1154
	} else {
		goto L1155
	}
L1136:
	;
	if v5295 != 0 {
		goto L1137
	} else {
		goto L1138
	}
L1137:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_93), int32(0))
	mBase = m.M
	v5300 = m.ExcPending
	if v5300 != 0 {
		goto L1
	} else {
		goto L1140
	}
L1138:
	;
	goto L1139
L1139:
	;
	v5306 = F_close(m, v5193)
	mBase = m.M
	v5307 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1572))
	v5308 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+704))
	if v5307 == int32(1) {
		goto L1144
	} else {
		goto L1145
	}
L1140:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3087), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v5305 = m.ExcPending
	if v5305 != 0 {
		goto L1
	} else {
		goto L1141
	}
L1141:
	;
	goto L1139
L1142:
	;
	v6404 = v4517
	v6405 = v4518
	v6406 = v4519
	v6407 = v4520
	v6408 = v4521
	v6409 = v4522
	v6410 = v4523
	v6415 = v4528
	v6416 = v4529
	v6417 = v4530
	v6418 = v4531
	v6422 = v4535
	v6424 = v4537
	v6427 = v4540
	v6428 = v4541
	v6429 = v4542
	v6431 = v4544
	v6435 = v4548
	v6436 = v4549
	v6437 = v4550
	v6438 = v4551
	v6439 = v4552
	goto L944
L1143:
	;
	goto L1142
L1144:
	;
	if v5308 == int32(0) {
		goto L1143
	} else {
		goto L1147
	}
L1145:
	;
	goto L1146
L1146:
	;
	if v5308 == int32(0) {
		goto L1143
	} else {
		goto L1151
	}
L1147:
	;
	v5314 = v5308
	goto L1148
L1148:
	;
	v5315 = *(*int32)(unsafe.Add(mBase, uint32(v5314)+28))
	v5316 = *(*int32)(unsafe.Add(mBase, uint32(v5314)+20))
	F_emscripten_builtin_free(m, v5316)
	mBase = m.M
	F_emscripten_builtin_free(m, v5314)
	mBase = m.M
	if v5315 != 0 {
		v5314 = v5315
		goto L1148
	} else {
		goto L1150
	}
L1149:
	;
	goto L1143
L1150:
	;
	goto L1149
L1151:
	;
	F_freeaddrinfo(m, v5308)
	mBase = m.M
	goto L1143
L1152:
	;
	F_gettimeofday(m, v4523+int32(1136))
	mBase = m.M
	v5344 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+1136))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1704)) = int32(0)
	F_gettimeofday(m, v4523+int32(816))
	mBase = m.M
	v5351 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4523)+1144)))
	v5352 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4523)+824)))
	v5355 = v5344 + int64(3)
	v5356 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+816))
	v5360 = v5351 - v5352 + (v5355-v5356)*int64(1000000)
	if v5360 <= int64(0) {
		goto L1164
	} else {
		goto L1165
	}
L1153:
	;
	goto L1152
L1154:
	;
	if v5325 == int32(0) {
		goto L1153
	} else {
		goto L1157
	}
L1155:
	;
	goto L1156
L1156:
	;
	if v5325 == int32(0) {
		goto L1153
	} else {
		goto L1161
	}
L1157:
	;
	v5331 = v5325
	goto L1158
L1158:
	;
	v5332 = *(*int32)(unsafe.Add(mBase, uint32(v5331)+28))
	v5333 = *(*int32)(unsafe.Add(mBase, uint32(v5331)+20))
	F_emscripten_builtin_free(m, v5333)
	mBase = m.M
	F_emscripten_builtin_free(m, v5331)
	mBase = m.M
	if v5332 != 0 {
		v5331 = v5332
		goto L1158
	} else {
		goto L1160
	}
L1159:
	;
	goto L1153
L1160:
	;
	goto L1159
L1161:
	;
	F_freeaddrinfo(m, v5325)
	mBase = m.M
	goto L1153
L1162:
	;
	v6403 = F_close(m, v5193)
	mBase = m.M
	v6404 = v6362
	v6405 = v6363
	v6406 = v6364
	v6407 = v6365
	v6408 = v6366
	v6409 = v6367
	v6410 = v6368
	v6415 = v6373
	v6416 = v6374
	v6417 = v6375
	v6418 = v6376
	v6422 = v6380
	v6424 = v6382
	v6427 = v6385
	v6428 = v6386
	v6429 = v6387
	v6431 = v6389
	v6435 = v6393
	v6436 = v6394
	v6437 = v6395
	v6438 = v6396
	v6439 = v6397
	goto L944
L1163:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), v6358, int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6361 = m.ExcPending
	if v6361 != 0 {
		goto L1
	} else {
		goto L1318
	}
L1164:
	;
	v6306 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6307 = m.ExcPending
	if v6307 != 0 {
		goto L1
	} else {
		goto L1315
	}
L1165:
	;
	v5363 = int32(1)
	v5364 = v5193 + v5363
	v5373 = v4523 + int32(880) + int32(base.Ui32(v5193)>>(uint(int32(3))%32))&int32(536870908)
	v5374 = int32(8)
	v5420 = v5360
	goto L1167
L1166:
	;
	v6260 = F_close(m, v5193)
	mBase = m.M
	F_pfree(m, v4535)
	mBase = m.M
	v6262 = m.ExcPending
	if v6262 != 0 {
		goto L1
	} else {
		goto L1314
	}
L1167:
	;
	v5424 = int64(1000000)
	v5425 = base.I64_div_u_s(v5420, v5424)
	*(*int64)(unsafe.Add(mBase, uint32(v4523)+848)) = v5425
	v5429 = v5420 - v5425*v5424
	*(*uint32)(unsafe.Add(mBase, uint32(v4523)+856)) = uint32(v5429)
	v5432 = v4523 + int32(880)
	base.MemoryFill(m, v5432, int32(0), int32(128))
	v5436 = *(*int32)(unsafe.Add(mBase, uint32(v5373)))
	*(*int32)(unsafe.Add(mBase, uint32(v5373))) = v5436 | v5363<<(uint(v5193)%32)
	v5440 = v4523 + int32(848)
	if v5440 != 0 {
		goto L1171
	} else {
		goto L1172
	}
L1168:
	;
	v6253 = F_close(m, v5193)
	mBase = m.M
	v6254 = *(*int32)(unsafe.Add(mBase, uint32(v4528)+364))
	F_set_authn_id(m, v4528, v6254)
	mBase = m.M
	v6256 = m.ExcPending
	if v6256 != 0 {
		goto L1
	} else {
		goto L1312
	}
L1169:
	;
	goto L1168
L1170:
	;
	if base.Ui32(int32(-4095)) <= base.Ui32(v5928) {
		goto L1230
	} else {
		goto L1231
	}
L1171:
	;
	v5441 = int32(-28)
	v5442 = *(*int64)(unsafe.Add(mBase, uint32(v5440)))
	if v5442 < int64(0) {
		v5928 = v5441
		goto L1170
	} else {
		goto L1174
	}
L1172:
	;
	goto L1173
L1173:
	;
	v5449 = int32(0)
	if v5364 < v5449 {
		v5920 = int32(-28)
		goto L1176
	} else {
		goto L1177
	}
L1174:
	;
	v5445 = *(*int32)(unsafe.Add(mBase, uint32(v5440)+8))
	if v5445 < int32(0) {
		v5928 = v5441
		goto L1170
	} else {
		goto L1175
	}
L1175:
	;
	goto L1173
L1176:
	;
	v5928 = v5920
	goto L1170
L1177:
	;
	if v5440 != 0 {
		goto L1178
	} else {
		goto L1179
	}
L1178:
	;
	v5453 = *(*int32)(unsafe.Add(mBase, uint32(v5440)))
	v5454 = int32(1000)
	v5456 = *(*int32)(unsafe.Add(mBase, uint32(v5440)+8))
	v5458 = base.I32_div_s(v5456, v5454)
	v5461 = v5453*v5454 + v5458
	goto L1180
L1179:
	;
	v5461 = int32(-1)
	goto L1180
L1180:
	;
	v5462 = int32(8)
	v5464 = int32(0)
	if v5364 == v5464 {
		v5482 = v5464
		goto L1183
	} else {
		goto L1184
	}
L1181:
	;
	if v5364 != 0 {
		goto L1191
	} else {
		goto L1192
	}
L1182:
	;
	goto L1181
L1183:
	;
	v5483 = F_emscripten_builtin_malloc(m, v5482)
	mBase = m.M
	if v5483 == int32(0) {
		goto L1182
	} else {
		goto L1189
	}
L1184:
	;
	v5470 = base.I64_extend_i32_u(v5364) * base.I64_extend_i32_u(v5462)
	v5471 = base.I32_wrap_i64(v5470)
	if base.Ui32(v5364|v5462) < base.Ui32(int32(_a_F_InitPostgres_94)) {
		v5482 = v5471
		goto L1183
	} else {
		goto L1185
	}
L1185:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v5470)>>(uint(int64(32))%64))) != 0 {
		goto L1186
	} else {
		goto L1187
	}
L1186:
	;
	v5479 = int32(-1)
	goto L1188
L1187:
	;
	v5479 = v5471
	goto L1188
L1188:
	;
	v5482 = v5479
	goto L1183
L1189:
	;
	v5488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5483-int32(4)))))
	if v5488&int32(3) == int32(0) {
		goto L1182
	} else {
		goto L1190
	}
L1190:
	;
	F___memset(m, v5483, int32(0), v5482)
	mBase = m.M
	goto L1182
L1191:
	;
	v5503 = int32(0)
	v5512 = v5449
	goto L1194
L1192:
	;
	v5585 = v5449
	goto L1193
L1193:
	;
	v5610 = m.Env.X__syscall_poll(m, v5483, v5585, v5461)
	mBase = m.M
	if v5610 < int32(0) {
		v5846 = v5610
		goto L1203
	} else {
		goto L1204
	}
L1194:
	;
	if v5432 == int32(0) {
		goto L1196
	} else {
		goto L1197
	}
L1195:
	;
	v5585 = v5565
	goto L1193
L1196:
	;
	v5560 = v5483 + v5512<<(uint(int32(3))%32)
	v5561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5560)+4)))
	if v5561 != 0 {
		goto L1199
	} else {
		goto L1200
	}
L1197:
	;
	v5544 = *(*int32)(unsafe.Add(mBase, uint32(v5432+int32(base.Ui32(v5503)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v5544)>>(uint(v5503)%32))&int32(1) == int32(0) {
		goto L1196
	} else {
		goto L1198
	}
L1198:
	;
	v5552 = v5483 + v5512<<(uint(int32(3))%32)
	v5553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5552)+4)))
	v5555 = v5553 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5552)+4)) = uint16(v5555)
	goto L1196
L1199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5560))) = v5503
	v5565 = v5512 + int32(1)
	goto L1201
L1200:
	;
	v5565 = v5512
	goto L1201
L1201:
	;
	v5567 = v5503 + int32(1)
	if v5567 != v5364 {
		v5503 = v5567
		v5512 = v5565
		goto L1194
	} else {
		goto L1202
	}
L1202:
	;
	goto L1195
L1203:
	;
	F_emscripten_builtin_free(m, v5483)
	mBase = m.M
	v5920 = v5846
	goto L1176
L1204:
	;
	if v5432 != 0 {
		goto L1205
	} else {
		goto L1206
	}
L1205:
	;
	v5621 = v5432
	v5635 = int32(32)
	goto L1208
L1206:
	;
	goto L1207
L1207:
	;
	v5702 = int32(0)
	if base.B2i32(v5610 == v5702)|base.B2i32(v5585 <= v5702) == v5702 {
		goto L1211
	} else {
		goto L1212
	}
L1208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5621))) = int32(0)
	v5660 = v5635 - int32(1)
	if v5660 != 0 {
		v5621 = v5621 + int32(4)
		v5635 = v5660
		goto L1208
	} else {
		goto L1210
	}
L1209:
	;
	goto L1207
L1210:
	;
	goto L1209
L1211:
	;
	v5717 = v5702
	goto L1215
L1212:
	;
	goto L1213
L1213:
	;
	v5846 = int32(0)
	goto L1203
L1214:
	;
	v5846 = int32(-8)
	goto L1203
L1215:
	;
	v5754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5483+v5717<<(uint(int32(3))%32))+6)))
	if v5754&int32(32) != 0 {
		goto L1214
	} else {
		goto L1217
	}
L1216:
	;
	v5760 = int32(0)
	v5769 = v5760
	v5771 = v5760
	goto L1219
L1217:
	;
	v5758 = v5717 + int32(1)
	if v5758 != v5585 {
		v5717 = v5758
		goto L1215
	} else {
		goto L1218
	}
L1218:
	;
	goto L1216
L1219:
	;
	v5805 = v5483 + v5769<<(uint(int32(3))%32)
	v5806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5805)+6)))
	if v5806 != 0 {
		goto L1221
	} else {
		goto L1222
	}
L1220:
	;
	v5846 = v5829
	goto L1203
L1221:
	;
	v5807 = *(*int32)(unsafe.Add(mBase, uint32(v5805)))
	v5808 = int32(0)
	if base.B2i32(v5432 == v5808)|base.B2i32(v5806&int32(25) == v5808) != 0 {
		goto L1224
	} else {
		goto L1225
	}
L1222:
	;
	v5829 = v5771
	goto L1223
L1223:
	;
	v5833 = v5769 + int32(1)
	if v5833 != v5585 {
		v5769 = v5833
		v5771 = v5829
		goto L1219
	} else {
		goto L1227
	}
L1224:
	;
	v5828 = v5771
	goto L1226
L1225:
	;
	v5819 = v5432 + int32(base.Ui32(v5807)>>(uint(int32(3))%32))&int32(536870908)
	v5820 = *(*int32)(unsafe.Add(mBase, uint32(v5819)))
	v5821 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5819))) = v5820 | v5821<<(uint(v5807)%32)
	v5828 = v5771 + v5821
	goto L1226
L1226:
	;
	v5829 = v5828
	goto L1223
L1227:
	;
	goto L1220
L1228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1704)) = int32(0)
	F_gettimeofday(m, v4523+int32(816))
	mBase = m.M
	v6244 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4523)+824)))
	v6246 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+816))
	v6250 = v5351 - v6244 + (v5355-v6246)*int64(1000000)
	if int64(0) < v6250 {
		v5420 = v6250
		goto L1167
	} else {
		goto L1311
	}
L1229:
	;
	if v5969 < int32(0) {
		goto L1233
	} else {
		goto L1234
	}
L1230:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41])) = int32(0) - v5928
	v5969 = int32(-1)
	goto L1232
L1231:
	;
	v5969 = v5928
	goto L1232
L1232:
	;
	goto L1229
L1233:
	;
	v5973 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	if v5973 == int32(27) {
		goto L1228
	} else {
		goto L1236
	}
L1234:
	;
	goto L1235
L1235:
	;
	if v5969 == int32(0) {
		goto L1240
	} else {
		goto L1241
	}
L1236:
	;
	v5978 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5979 = m.ExcPending
	if v5979 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	if v5978 == int32(0) {
		v6362 = v4517
		v6363 = v4518
		v6364 = v4519
		v6365 = v4520
		v6366 = v4521
		v6367 = v4522
		v6368 = v4523
		v6373 = v4528
		v6374 = v4529
		v6375 = v4530
		v6376 = v4531
		v6380 = v4535
		v6382 = v4537
		v6385 = v4540
		v6386 = v4541
		v6387 = v4542
		v6389 = v4544
		v6393 = v4548
		v6394 = v4549
		v6395 = v4550
		v6396 = v4551
		v6397 = v4552
		goto L1162
	} else {
		goto L1238
	}
L1238:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_95), int32(0))
	mBase = m.M
	v5985 = m.ExcPending
	if v5985 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	v6358 = int32(3140)
	goto L1163
L1240:
	;
	v5991 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5992 = m.ExcPending
	if v5992 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1241:
	;
	goto L1242
L1242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+1708)) = int32(28)
	v6007 = int32(0)
	v6012 = F_recvfrom(m, v5193, v4523+int32(1712), int32(1024), v6007, v4523+int32(736), v4523+int32(1708))
	mBase = m.M
	if v6012 < v6007 {
		goto L1246
	} else {
		goto L1247
	}
L1243:
	;
	if v5991 == int32(0) {
		v6362 = v4517
		v6363 = v4518
		v6364 = v4519
		v6365 = v4520
		v6366 = v4521
		v6367 = v4522
		v6368 = v4523
		v6373 = v4528
		v6374 = v4529
		v6375 = v4530
		v6376 = v4531
		v6380 = v4535
		v6382 = v4537
		v6385 = v4540
		v6386 = v4541
		v6387 = v4542
		v6389 = v4544
		v6393 = v4548
		v6394 = v4549
		v6395 = v4550
		v6396 = v4551
		v6397 = v4552
		goto L1162
	} else {
		goto L1244
	}
L1244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+480)) = v4565
	F_errmsg(m, int32(_a_F_InitPostgres_96), v4523+int32(480))
	mBase = m.M
	v6000 = m.ExcPending
	if v6000 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1245:
	;
	v6358 = int32(3148)
	goto L1163
L1246:
	;
	v6017 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6018 = m.ExcPending
	if v6018 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1247:
	;
	goto L1248
L1248:
	;
	v6026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+738)))
	if (v4630<<(uint(v5374)%32)|int32(base.Ui32(v4630&int32(_a_F_InitPostgres_97))>>(uint(v5374)%32)))&int32(_a_F_InitPostgres_90) != v6026 {
		goto L1252
	} else {
		goto L1253
	}
L1249:
	;
	if v6017 == int32(0) {
		v6362 = v4517
		v6363 = v4518
		v6364 = v4519
		v6365 = v4520
		v6366 = v4521
		v6367 = v4522
		v6368 = v4523
		v6373 = v4528
		v6374 = v4529
		v6375 = v4530
		v6376 = v4531
		v6380 = v4535
		v6382 = v4537
		v6385 = v4540
		v6386 = v4541
		v6387 = v4542
		v6389 = v4544
		v6393 = v4548
		v6394 = v4549
		v6395 = v4550
		v6396 = v4551
		v6397 = v4552
		goto L1162
	} else {
		goto L1250
	}
L1250:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_98), int32(0))
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	v6358 = int32(3170)
	goto L1163
L1252:
	;
	v6030 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6031 = m.ExcPending
	if v6031 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1253:
	;
	goto L1254
L1254:
	;
	if base.Ui32(v6012) <= base.Ui32(int32(19)) {
		goto L1259
	} else {
		goto L1260
	}
L1255:
	;
	if v6030 == int32(0) {
		goto L1228
	} else {
		goto L1256
	}
L1256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+592)) = v4565
	v6035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+738)))
	v6036 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+596)) = (v6035<<(uint(v6036)%32) | int32(base.Ui32(v6035)>>(uint(v6036)%32))) & int32(_a_F_InitPostgres_90)
	F_errmsg(m, int32(_a_F_InitPostgres_99), v4523+int32(592))
	mBase = m.M
	v6048 = m.ExcPending
	if v6048 != 0 {
		goto L1
	} else {
		goto L1257
	}
L1257:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3179), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6053 = m.ExcPending
	if v6053 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1258:
	;
	goto L1228
L1259:
	;
	v6058 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6059 = m.ExcPending
	if v6059 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1260:
	;
	goto L1261
L1261:
	;
	v6074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+1714)))
	v6075 = int32(8)
	if (v6074<<(uint(v6075)%32)|int32(base.Ui32(v6074)>>(uint(v6075)%32)))&int32(_a_F_InitPostgres_90) != v6012 {
		goto L1266
	} else {
		goto L1267
	}
L1262:
	;
	if v6058 == int32(0) {
		goto L1228
	} else {
		goto L1263
	}
L1263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+500)) = v6012
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+496)) = v4565
	F_errmsg(m, int32(_a_F_InitPostgres_100), v4523+int32(496))
	mBase = m.M
	v6068 = m.ExcPending
	if v6068 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3186), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6073 = m.ExcPending
	if v6073 != 0 {
		goto L1
	} else {
		goto L1265
	}
L1265:
	;
	goto L1228
L1266:
	;
	v6085 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6086 = m.ExcPending
	if v6086 != 0 {
		goto L1
	} else {
		goto L1269
	}
L1267:
	;
	goto L1268
L1268:
	;
	v6110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+2753)))
	v6111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+1713)))
	if v6110 != v6111 {
		goto L1273
	} else {
		goto L1274
	}
L1269:
	;
	if v6085 == int32(0) {
		goto L1228
	} else {
		goto L1270
	}
L1270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+576)) = v4565
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+584)) = v6012
	v6091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4523)+1714)))
	v6092 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+580)) = (v6091<<(uint(v6092)%32) | int32(base.Ui32(v6091)>>(uint(v6092)%32))) & int32(_a_F_InitPostgres_90)
	F_errmsg(m, int32(_a_F_InitPostgres_101), v4523+int32(576))
	mBase = m.M
	v6104 = m.ExcPending
	if v6104 != 0 {
		goto L1
	} else {
		goto L1271
	}
L1271:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3194), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6109 = m.ExcPending
	if v6109 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	goto L1228
L1273:
	;
	v6115 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6116 = m.ExcPending
	if v6116 != 0 {
		goto L1
	} else {
		goto L1276
	}
L1274:
	;
	goto L1275
L1275:
	;
	v6134 = F_strlen(m, v4569)
	mBase = m.M
	v6136 = F_palloc(m, v6134+v6012)
	mBase = m.M
	v6137 = m.ExcPending
	if v6137 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1276:
	;
	if v6115 == int32(0) {
		goto L1228
	} else {
		goto L1277
	}
L1277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+560)) = v4565
	v6120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+1713)))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+564)) = v6120
	v6122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+2753)))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+568)) = v6122
	F_errmsg(m, int32(_a_F_InitPostgres_102), v4523+int32(560))
	mBase = m.M
	v6128 = m.ExcPending
	if v6128 != 0 {
		goto L1
	} else {
		goto L1278
	}
L1278:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3202), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6133 = m.ExcPending
	if v6133 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1279:
	;
	goto L1228
L1280:
	;
	v6138 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1712))
	*(*int32)(unsafe.Add(mBase, uint32(v6136))) = v6138
	v6140 = *(*int64)(unsafe.Add(mBase, uint32(v4531)))
	*(*int64)(unsafe.Add(mBase, uint32(v6136)+4)) = v6140
	v6142 = *(*int64)(unsafe.Add(mBase, uint32(v4531)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v6136)+12)) = v6142
	if v6012 == int32(20) {
		goto L1281
	} else {
		goto L1282
	}
L1281:
	;
	v6154 = F_strlen(m, v4569)
	mBase = m.M
	if v6154 != 0 {
		goto L1284
	} else {
		goto L1285
	}
L1282:
	;
	v6147 = v6012 - int32(20)
	if v6147 == int32(0) {
		goto L1281
	} else {
		goto L1283
	}
L1283:
	;
	base.MemoryCopy(m, v6136+int32(20), v4551, v6147)
	goto L1281
L1284:
	;
	base.MemoryCopy(m, v6136+v6012, v4569, v6154)
	goto L1286
L1285:
	;
	goto L1286
L1286:
	;
	v6157 = F_strlen(m, v4569)
	mBase = m.M
	v6163 = F_pg_md5_binary(m, v6136, v6157+v6012, v4523+int32(1168), v4523+int32(1704))
	mBase = m.M
	v6164 = m.ExcPending
	if v6164 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1287:
	;
	if v6163 == int32(0) {
		goto L1288
	} else {
		goto L1289
	}
L1288:
	;
	v6169 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6170 = m.ExcPending
	if v6170 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1289:
	;
	goto L1290
L1290:
	;
	F_pfree(m, v6136)
	mBase = m.M
	v6186 = m.ExcPending
	if v6186 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1291:
	;
	if v6169 != 0 {
		goto L1292
	} else {
		goto L1293
	}
L1292:
	;
	v6171 = *(*int32)(unsafe.Add(mBase, uint32(v4523)+1704))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+544)) = v6171
	F_errmsg(m, int32(_a_F_InitPostgres_103), v4523+int32(544))
	mBase = m.M
	v6177 = m.ExcPending
	if v6177 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1293:
	;
	goto L1294
L1294:
	;
	F_pfree(m, v6136)
	mBase = m.M
	v6184 = m.ExcPending
	if v6184 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1295:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3229), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6182 = m.ExcPending
	if v6182 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	goto L1294
L1297:
	;
	goto L1228
L1298:
	;
	v6187 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+1716))
	v6188 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+1168))
	v6190 = *(*int64)(unsafe.Add(mBase, uint32(v4552)))
	v6191 = *(*int64)(unsafe.Add(mBase, uint32(v4523)+1176))
	if v6187^v6188|(v6190^v6191) != int64(0) {
		goto L1299
	} else {
		goto L1300
	}
L1299:
	;
	v6198 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6199 = m.ExcPending
	if v6199 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1300:
	;
	goto L1301
L1301:
	;
	v6213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+1712)))
	switch v6213 - int32(2) {
	case 0:
		goto L1169
	case 1:
		goto L1166
	default:
		goto L1306
	}
L1302:
	;
	if v6198 == int32(0) {
		goto L1228
	} else {
		goto L1303
	}
L1303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+528)) = v4565
	F_errmsg(m, int32(_a_F_InitPostgres_104), v4523+int32(528))
	mBase = m.M
	v6207 = m.ExcPending
	if v6207 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3239), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6212 = m.ExcPending
	if v6212 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1305:
	;
	goto L1228
L1306:
	;
	v6218 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6219 = m.ExcPending
	if v6219 != 0 {
		goto L1
	} else {
		goto L1307
	}
L1307:
	;
	if v6218 == int32(0) {
		goto L1228
	} else {
		goto L1308
	}
L1308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+512)) = v4565
	v6223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4523)+1712)))
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+516)) = v6223
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+520)) = v4570
	F_errmsg(m, int32(_a_F_InitPostgres_105), v4523+int32(512))
	mBase = m.M
	v6230 = m.ExcPending
	if v6230 != 0 {
		goto L1
	} else {
		goto L1309
	}
L1309:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(3257), int32(_a_F_InitPostgres_84))
	mBase = m.M
	v6235 = m.ExcPending
	if v6235 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1310:
	;
	goto L1228
L1311:
	;
	goto L1164
L1312:
	;
	F_pfree(m, v4535)
	mBase = m.M
	v6258 = m.ExcPending
	if v6258 != 0 {
		goto L1
	} else {
		goto L1313
	}
L1313:
	;
	v6585 = v4517
	v6586 = v4518
	v6587 = v4519
	v6588 = v4520
	v6589 = v4521
	v6590 = v4522
	v6591 = v4523
	v6596 = v4528
	v6597 = v4529
	v6598 = int32(0)
	v6609 = v4541
	v6618 = v4550
	goto L501
L1314:
	;
	v6585 = v4517
	v6586 = v4518
	v6587 = v4519
	v6588 = v4520
	v6589 = v4521
	v6590 = v4522
	v6591 = v4523
	v6596 = v4528
	v6597 = v4529
	v6598 = v4530
	v6609 = v4541
	v6618 = v4550
	goto L501
L1315:
	;
	if v6306 == int32(0) {
		v6362 = v4517
		v6363 = v4518
		v6364 = v4519
		v6365 = v4520
		v6366 = v4521
		v6367 = v4522
		v6368 = v4523
		v6373 = v4528
		v6374 = v4529
		v6375 = v4530
		v6376 = v4531
		v6380 = v4535
		v6382 = v4537
		v6385 = v4540
		v6386 = v4541
		v6387 = v4542
		v6389 = v4544
		v6393 = v4548
		v6394 = v4549
		v6395 = v4550
		v6396 = v4551
		v6397 = v4552
		goto L1162
	} else {
		goto L1316
	}
L1316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4523)+464)) = v4565
	F_errmsg(m, int32(_a_F_InitPostgres_96), v4523+int32(464))
	mBase = m.M
	v6315 = m.ExcPending
	if v6315 != 0 {
		goto L1
	} else {
		goto L1317
	}
L1317:
	;
	v6358 = int32(3122)
	goto L1163
L1318:
	;
	v6362 = v4517
	v6363 = v4518
	v6364 = v4519
	v6365 = v4520
	v6366 = v4521
	v6367 = v4522
	v6368 = v4523
	v6373 = v4528
	v6374 = v4529
	v6375 = v4530
	v6376 = v4531
	v6380 = v4535
	v6382 = v4537
	v6385 = v4540
	v6386 = v4541
	v6387 = v4542
	v6389 = v4544
	v6393 = v4548
	v6394 = v4549
	v6395 = v4550
	v6396 = v4551
	v6397 = v4552
	goto L1162
L1319:
	;
	v6464 = *(*int32)(unsafe.Add(mBase, uint32(v6445)+396))
	if v6464 == int32(0) {
		v6481 = v6429
		goto L1325
	} else {
		goto L1326
	}
L1320:
	;
	v6449 = *(*int32)(unsafe.Add(mBase, uint32(v6446)+4))
	if v6449 < int32(2) {
		v6463 = v6431
		goto L1319
	} else {
		goto L1321
	}
L1321:
	;
	v6453 = v6431 + int32(4)
	v6455 = *(*int32)(unsafe.Add(mBase, uint32(v6446)+12))
	if base.Ui32(v6453) < base.Ui32(v6455+v6449<<(uint(int32(2))%32)) {
		goto L1322
	} else {
		goto L1323
	}
L1322:
	;
	v6460 = v6453
	goto L1324
L1323:
	;
	v6460 = int32(0)
	goto L1324
L1324:
	;
	v6463 = v6460
	goto L1319
L1325:
	;
	v6482 = *(*int32)(unsafe.Add(mBase, uint32(v6445)+388))
	if v6482 == int32(0) {
		v6499 = v6424
		goto L1331
	} else {
		goto L1332
	}
L1326:
	;
	v6467 = *(*int32)(unsafe.Add(mBase, uint32(v6464)+4))
	if v6467 < int32(2) {
		v6481 = v6429
		goto L1325
	} else {
		goto L1327
	}
L1327:
	;
	v6471 = v6429 + int32(4)
	v6473 = *(*int32)(unsafe.Add(mBase, uint32(v6464)+12))
	if base.Ui32(v6471) < base.Ui32(v6473+v6467<<(uint(int32(2))%32)) {
		goto L1328
	} else {
		goto L1329
	}
L1328:
	;
	v6478 = v6471
	goto L1330
L1329:
	;
	v6478 = int32(0)
	goto L1330
L1330:
	;
	v6481 = v6478
	goto L1325
L1331:
	;
	v6501 = v6436 + int32(1)
	v6502 = *(*int32)(unsafe.Add(mBase, uint32(v6427)+4))
	if v6501 < v6502 {
		v4517 = v6404
		v4518 = v6405
		v4519 = v6406
		v4520 = v6407
		v4521 = v6408
		v4522 = v6409
		v4523 = v6410
		v4528 = v6415
		v4529 = v6416
		v4530 = v6417
		v4531 = v6418
		v4535 = v6422
		v4537 = v6499
		v4540 = v6427
		v4541 = v6428
		v4542 = v6481
		v4544 = v6463
		v4548 = v6435
		v4549 = v6501
		v4550 = v6437
		v4551 = v6438
		v4552 = v6439
		goto L917
	} else {
		goto L1337
	}
L1332:
	;
	v6485 = *(*int32)(unsafe.Add(mBase, uint32(v6482)+4))
	if v6485 < int32(2) {
		v6499 = v6424
		goto L1331
	} else {
		goto L1333
	}
L1333:
	;
	v6489 = v6424 + int32(4)
	v6491 = *(*int32)(unsafe.Add(mBase, uint32(v6482)+12))
	if base.Ui32(v6489) < base.Ui32(v6491+v6485<<(uint(int32(2))%32)) {
		goto L1334
	} else {
		goto L1335
	}
L1334:
	;
	v6496 = v6489
	goto L1336
L1335:
	;
	v6496 = int32(0)
	goto L1336
L1336:
	;
	v6499 = v6496
	goto L1331
L1337:
	;
	goto L918
L1338:
	;
	v6585 = v6504
	v6586 = v6505
	v6587 = v6506
	v6588 = v6507
	v6589 = v6508
	v6590 = v6509
	v6591 = v6510
	v6596 = v6515
	v6597 = v6516
	v6598 = v6517
	v6609 = v6528
	v6618 = v6537
	goto L501
L1339:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(460), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v6562 = m.ExcPending
	if v6562 != 0 {
		goto L1
	} else {
		goto L1340
	}
L1340:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1341:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v6569 = m.ExcPending
	if v6569 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1342:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_106), int32(0))
	mBase = m.M
	v6573 = m.ExcPending
	if v6573 != 0 {
		goto L1
	} else {
		goto L1343
	}
L1343:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(405), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v6578 = m.ExcPending
	if v6578 != 0 {
		goto L1
	} else {
		goto L1344
	}
L1344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1345:
	;
	v6585 = l0
	v6586 = l1
	v6587 = l2
	v6588 = l3
	v6589 = l4
	v6590 = l5
	v6591 = v1106
	v6596 = v1075
	v6597 = v44
	v6598 = v6582
	v6609 = v47
	v6618 = v7
	goto L501
L1346:
	;
	v6666 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[61]))
	if v6666 != 0 {
		goto L1354
	} else {
		goto L1355
	}
L1347:
	;
	v6634 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[62]))
	if v6634 != 0 {
		goto L1346
	} else {
		goto L1348
	}
L1348:
	;
	v6637 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6638 = m.ExcPending
	if v6638 != 0 {
		goto L1
	} else {
		goto L1349
	}
L1349:
	;
	if v6637 == int32(0) {
		goto L1346
	} else {
		goto L1350
	}
L1350:
	;
	v6641 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+364))
	v6642 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+380))
	v6643 = *(*int32)(unsafe.Add(mBase, uint32(v6642)+296))
	v6646 = *(*int32)(unsafe.Add(mBase, uint32(v6643<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[63])))
	goto L1351
L1351:
	;
	v6647 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+380))
	v6648 = *(*int64)(unsafe.Add(mBase, uint32(v6647)))
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+68)) = v6646
	*(*int64)(unsafe.Add(mBase, uint32(v6591)+72)) = v6648
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+64)) = v6641
	F_errmsg(m, int32(_a_F_InitPostgres_107), v6591-int32(-64))
	mBase = m.M
	v6656 = m.ExcPending
	if v6656 != 0 {
		goto L1
	} else {
		goto L1352
	}
L1352:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(660), int32(_a_F_InitPostgres_51))
	mBase = m.M
	v6661 = m.ExcPending
	if v6661 != 0 {
		goto L1
	} else {
		goto L1353
	}
L1353:
	;
	goto L1346
L1354:
	;
	m.T0[v6666].(func(*base.Module, int32, int32))(m, v6596, v6598)
	mBase = m.M
	v6668 = m.ExcPending
	if v6668 != 0 {
		goto L1
	} else {
		goto L1357
	}
L1355:
	;
	goto L1356
L1356:
	;
	if v6598 == int32(0) {
		goto L1359
	} else {
		goto L1360
	}
L1357:
	;
	goto L1356
L1358:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6759 = m.ExcPending
	if v6759 != 0 {
		goto L1
	} else {
		goto L1394
	}
L1359:
	;
	v6672 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v6672 != 0 {
		goto L1362
	} else {
		goto L1363
	}
L1360:
	;
	goto L1361
L1361:
	;
	if v6598 != int32(-2) {
		goto L1373
	} else {
		goto L1374
	}
L1362:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6674 = m.ExcPending
	if v6674 != 0 {
		goto L1
	} else {
		goto L1365
	}
L1363:
	;
	goto L1364
L1364:
	;
	v6676 = v6591 + int32(2752)
	F_pq_beginmessage(m, v6676, int32(82))
	mBase = m.M
	v6679 = m.ExcPending
	if v6679 != 0 {
		goto L1
	} else {
		goto L1366
	}
L1365:
	;
	goto L1364
L1366:
	;
	F_enlargeStringInfo(m, v6676, int32(4))
	mBase = m.M
	v6682 = m.ExcPending
	if v6682 != 0 {
		goto L1
	} else {
		goto L1367
	}
L1367:
	;
	v6683 = *(*int32)(unsafe.Add(mBase, uint32(v6591)+2756))
	v6684 = *(*int32)(unsafe.Add(mBase, uint32(v6591)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v6683+v6684))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+2756)) = v6683 + int32(4)
	F_pq_endmessage(m, v6676)
	mBase = m.M
	v6692 = m.ExcPending
	if v6692 != 0 {
		goto L1
	} else {
		goto L1368
	}
L1368:
	;
	v6694 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[40]))
	if v6694 != 0 {
		goto L1369
	} else {
		goto L1370
	}
L1369:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6696 = m.ExcPending
	if v6696 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1370:
	;
	goto L1371
L1371:
	;
	m.G0 = v6591 + int32(3792)
	goto L1358
L1372:
	;
	goto L1371
L1373:
	;
	v6702 = *(*int32)(unsafe.Add(mBase, uint32(v6591)+700))
	v6703 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+380))
	v6704 = *(*int32)(unsafe.Add(mBase, uint32(v6703)+296))
	if base.Ui32(int32(15)) < base.Ui32(v6704) {
		goto L1377
	} else {
		goto L1378
	}
L1374:
	;
	goto L1375
L1375:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6756 = m.ExcPending
	if v6756 != 0 {
		goto L1
	} else {
		goto L1393
	}
L1376:
	;
	v6716 = *(*int64)(unsafe.Add(mBase, uint32(v6703)))
	v6717 = *(*int32)(unsafe.Add(mBase, uint32(v6703)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+56)) = v6717
	*(*int64)(unsafe.Add(mBase, uint32(v6591)+48)) = v6716
	v6723 = F_psprintf(m, int32(_a_F_InitPostgres_108), v6591+int32(48))
	mBase = m.M
	v6724 = m.ExcPending
	if v6724 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1377:
	;
	v6713 = int32(514)
	v6715 = int32(_a_F_InitPostgres_109)
	goto L1376
L1378:
	;
	goto L1379
L1379:
	;
	v6710 = v6704 << (uint(int32(2)) % 32)
	v6711 = *(*int32)(unsafe.Add(mBase, uint32(v6710)+uint32(_c_F_InitPostgres[64])))
	v6712 = *(*int32)(unsafe.Add(mBase, uint32(v6710)+uint32(_c_F_InitPostgres[65])))
	v6713 = v6711
	v6715 = v6712
	goto L1376
L1380:
	;
	if v6702 != 0 {
		goto L1381
	} else {
		goto L1382
	}
L1381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+36)) = v6723
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+32)) = v6702
	v6730 = F_psprintf(m, int32(_a_F_InitPostgres_110), v6591+int32(32))
	mBase = m.M
	v6731 = m.ExcPending
	if v6731 != 0 {
		goto L1
	} else {
		goto L1384
	}
L1382:
	;
	v6732 = v6723
	goto L1383
L1383:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6736 = m.ExcPending
	if v6736 != 0 {
		goto L1
	} else {
		goto L1385
	}
L1384:
	;
	v6732 = v6730
	goto L1383
L1385:
	;
	F_errcode(m, v6713)
	mBase = m.M
	v6738 = m.ExcPending
	if v6738 != 0 {
		goto L1
	} else {
		goto L1386
	}
L1386:
	;
	v6739 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v6591)+16)) = v6739
	F_errmsg(m, v6715, v6591+int32(16))
	mBase = m.M
	v6744 = m.ExcPending
	if v6744 != 0 {
		goto L1
	} else {
		goto L1387
	}
L1387:
	;
	if v6732 != 0 {
		goto L1388
	} else {
		goto L1389
	}
L1388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6591))) = v6732
	F_errdetail_log(m, int32(_a_F_InitPostgres_111), v6591)
	mBase = m.M
	v6748 = m.ExcPending
	if v6748 != 0 {
		goto L1
	} else {
		goto L1391
	}
L1389:
	;
	goto L1390
L1390:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_50), int32(320), int32(_a_F_InitPostgres_112))
	mBase = m.M
	v6753 = m.ExcPending
	if v6753 != 0 {
		goto L1
	} else {
		goto L1392
	}
L1391:
	;
	goto L1390
L1392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1394:
	;
	v6764 = m.G0
	v6765 = int32(16)
	v6766 = v6764 - v6765
	m.G0 = v6766
	F_gettimeofday(m, v6766)
	mBase = m.M
	v6769 = *(*int64)(unsafe.Add(mBase, uint32(v6766)))
	v6770 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6766)+8)))
	m.G0 = v6766 + v6765
	goto L1395
L1395:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[66])) = v6770 + v6769*int64(1000000) - int64(946684800000000)
	v6781 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[54])))
	if v6781&int32(4) != 0 {
		goto L1396
	} else {
		goto L1397
	}
L1396:
	;
	v6785 = v6597 + int32(432)
	F_initStringInfo(m, v6785)
	mBase = m.M
	v6787 = m.ExcPending
	if v6787 != 0 {
		goto L1
	} else {
		goto L1399
	}
L1397:
	;
	goto L1398
L1398:
	;
	v6841 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[32])) = uint8(v6841)
	F_InitializeSessionUserId(m, v6587, v6588, v6841)
	mBase = m.M
	v6845 = m.ExcPending
	if v6845 != 0 {
		goto L1
	} else {
		goto L1419
	}
L1399:
	;
	v6788 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v6597)+400)) = v6788
	v6793 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	if v6793 != 0 {
		goto L1400
	} else {
		goto L1401
	}
L1400:
	;
	v6794 = int32(_a_F_InitPostgres_113)
	goto L1402
L1401:
	;
	v6794 = int32(_a_F_InitPostgres_114)
	goto L1402
L1402:
	;
	F_appendStringInfo(m, v6785, v6794, v6597+int32(400))
	mBase = m.M
	v6798 = m.ExcPending
	if v6798 != 0 {
		goto L1
	} else {
		goto L1403
	}
L1403:
	;
	v6800 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	if v6800 == int32(0) {
		goto L1404
	} else {
		goto L1405
	}
L1404:
	;
	v6803 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v6597)+384)) = v6803
	F_appendStringInfo(m, v6785, int32(_a_F_InitPostgres_115), v6597+int32(384))
	mBase = m.M
	v6809 = m.ExcPending
	if v6809 != 0 {
		goto L1
	} else {
		goto L1407
	}
L1405:
	;
	goto L1406
L1406:
	;
	v6810 = *(*int32)(unsafe.Add(mBase, uint32(v6596)+376))
	if v6810 != 0 {
		goto L1408
	} else {
		goto L1409
	}
L1407:
	;
	goto L1406
L1408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6597)+368)) = v6810
	F_appendStringInfo(m, v6597+int32(432), int32(_a_F_InitPostgres_116), v6597+int32(368))
	mBase = m.M
	v6818 = m.ExcPending
	if v6818 != 0 {
		goto L1
	} else {
		goto L1411
	}
L1409:
	;
	goto L1410
L1410:
	;
	v6821 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6822 = m.ExcPending
	if v6822 != 0 {
		goto L1
	} else {
		goto L1412
	}
L1411:
	;
	goto L1410
L1412:
	;
	if v6821 != 0 {
		goto L1413
	} else {
		goto L1414
	}
L1413:
	;
	v6823 = *(*int32)(unsafe.Add(mBase, uint32(v6597)+432))
	*(*int32)(unsafe.Add(mBase, uint32(v6597)+352)) = v6823
	F_errmsg_internal(m, int32(_a_F_InitPostgres_111), v6597+int32(352))
	mBase = m.M
	v6829 = m.ExcPending
	if v6829 != 0 {
		goto L1
	} else {
		goto L1416
	}
L1414:
	;
	goto L1415
L1415:
	;
	v6835 = *(*int32)(unsafe.Add(mBase, uint32(v6597)+432))
	F_pfree(m, v6835)
	mBase = m.M
	v6837 = m.ExcPending
	if v6837 != 0 {
		goto L1
	} else {
		goto L1418
	}
L1416:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(309), int32(_a_F_InitPostgres_117))
	mBase = m.M
	v6834 = m.ExcPending
	if v6834 != 0 {
		goto L1
	} else {
		goto L1417
	}
L1417:
	;
	goto L1415
L1418:
	;
	goto L1398
L1419:
	;
	v6847 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[62]))
	if v6847 == int32(0) {
		v6857 = v6585
		v6858 = v6586
		v6861 = v6589
		v6862 = v6590
		v6869 = v6597
		v6881 = v6609
		v6890 = v6618
		goto L134
	} else {
		goto L1420
	}
L1420:
	;
	v6851 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[67]))
	v6854 = *(*int32)(unsafe.Add(mBase, uint32(v6851<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[63])))
	goto L1421
L1421:
	;
	F_InitializeSystemUser(m, v6847, v6854)
	mBase = m.M
	v6856 = m.ExcPending
	if v6856 != 0 {
		goto L1
	} else {
		goto L1422
	}
L1422:
	;
	v6857 = v6585
	v6858 = v6586
	v6861 = v6589
	v6862 = v6590
	v6869 = v6597
	v6881 = v6609
	v6890 = v6618
	goto L134
L1423:
	;
	v6900 = v6857
	v6901 = v6858
	v6903 = v6898
	v6904 = v6861
	v6905 = v6862
	v6912 = v6869
	v6924 = v6881
	v6933 = v6890
	goto L133
L1424:
	;
	v6943 = int32(_a_F_InitPostgres_118)
	v6945 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[68]))
	v6946 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[68])) = v6945 + v6946
	v6950 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[69]))
	v6951 = *(*int32)(unsafe.Add(mBase, uint32(v6950)))
	*(*int32)(unsafe.Add(mBase, uint32(v6950))) = v6951 + v6946
	v6955 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6950)+192)) = uint8(v6955)
	*(*uint8)(unsafe.Add(mBase, uint32(v6950)+200)) = uint8(v6955)
	*(*int32)(unsafe.Add(mBase, uint32(v6950))) = v6951 + int32(2)
	v6965 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[68]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[68])) = v6965 - v6946
	goto L1426
L1425:
	;
	goto L1426
L1426:
	;
	v6971 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[70])))
	if (v6971^int32(-1)|v6903)&int32(1) != 0 {
		goto L1435
	} else {
		goto L1436
	}
L1427:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v8486 = m.ExcPending
	if v8486 != 0 {
		goto L1
	} else {
		goto L1784
	}
L1428:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71])) = v7643
	v7651 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v7651)+60)) = v7643
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v7654 = m.ExcPending
	if v7654 != 0 {
		goto L1
	} else {
		goto L1601
	}
L1429:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7621 = m.ExcPending
	if v7621 != 0 {
		goto L1
	} else {
		goto L1596
	}
L1430:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[72])) = int32(1663)
	v7643 = int32(1)
	goto L1428
L1431:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7599 = m.ExcPending
	if v7599 != 0 {
		goto L1
	} else {
		goto L1592
	}
L1432:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7576 = m.ExcPending
	if v7576 != 0 {
		goto L1
	} else {
		goto L1587
	}
L1433:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7557 = m.ExcPending
	if v7557 != 0 {
		goto L1
	} else {
		goto L1583
	}
L1434:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7538 = m.ExcPending
	if v7538 != 0 {
		goto L1
	} else {
		goto L1579
	}
L1435:
	;
	v6978 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v6979 = int32(1)
	if (base.B2i32(v6978 != v6979)|v6903)&v6979 != 0 {
		goto L1438
	} else {
		goto L1439
	}
L1436:
	;
	goto L1437
L1437:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7522 = m.ExcPending
	if v7522 != 0 {
		goto L1
	} else {
		goto L1575
	}
L1438:
	;
	v7162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	if v7162 != int32(1) {
		goto L1456
	} else {
		goto L1457
	}
L1439:
	;
	v6985 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[73]))
	v6987 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	v6988 = v6985 + v6987
	if v6988 <= int32(0) {
		goto L1438
	} else {
		goto L1440
	}
L1440:
	;
	v6992 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	v6993 = *(*int32)(unsafe.Add(mBase, uint32(v6992)))
	*(*int32)(unsafe.Add(mBase, uint32(v6992))) = int32(1)
	if v6993 != 0 {
		goto L1441
	} else {
		goto L1442
	}
L1441:
	;
	v6997 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	F_s_lock(m, v6997, int32(_a_F_InitPostgres_119), int32(789), int32(_a_F_InitPostgres_120))
	mBase = m.M
	v7002 = m.ExcPending
	if v7002 != 0 {
		goto L1
	} else {
		goto L1444
	}
L1442:
	;
	goto L1443
L1443:
	;
	v7004 = v6912 + int32(428)
	v7005 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7004))) = v7005
	v7008 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	v7009 = *(*int32)(unsafe.Add(mBase, uint32(v7008)+24))
	if v7009 == v7005 {
		goto L1445
	} else {
		goto L1446
	}
L1444:
	;
	goto L1443
L1445:
	;
	v7104 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	*(*int32)(unsafe.Add(mBase, uint32(v7104))) = int32(0)
	v7107 = *(*int32)(unsafe.Add(mBase, uint32(v7004)))
	if v7107 == v6988 {
		goto L1438
	} else {
		goto L1452
	}
L1446:
	;
	v7013 = v7008 + int32(20)
	if v7009 == v7013 {
		goto L1445
	} else {
		goto L1447
	}
L1447:
	;
	v7022 = v7009
	v7048 = v6933
	goto L1448
L1448:
	;
	v7057 = v7048 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7004))) = v7057
	if v6988 == v7057 {
		goto L1445
	} else {
		goto L1450
	}
L1449:
	;
	goto L1445
L1450:
	;
	v7060 = *(*int32)(unsafe.Add(mBase, uint32(v7022)+4))
	if v7060 != v7013 {
		v7022 = v7060
		v7048 = v7057
		goto L1448
	} else {
		goto L1451
	}
L1451:
	;
	goto L1449
L1452:
	;
	v7109 = *(*int32)(unsafe.Add(mBase, uint32(v6912)+428))
	v7111 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	if v7109 < v7111 {
		goto L1434
	} else {
		goto L1453
	}
L1453:
	;
	v7114 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	v7116 = F_has_privs_of_role(m, v7114, int32(_a_F_InitPostgres_121))
	mBase = m.M
	v7117 = m.ExcPending
	if v7117 != 0 {
		goto L1
	} else {
		goto L1454
	}
L1454:
	;
	if v7116 == int32(0) {
		goto L1433
	} else {
		goto L1455
	}
L1455:
	;
	goto L1438
L1456:
	;
	if v6924 == int32(0) {
		goto L1430
	} else {
		goto L1471
	}
L1457:
	;
	v7166 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	v7167 = F_has_rolreplication(m, v7166)
	mBase = m.M
	v7168 = m.ExcPending
	if v7168 != 0 {
		goto L1
	} else {
		goto L1458
	}
L1458:
	;
	if v7167 == int32(0) {
		goto L1432
	} else {
		goto L1459
	}
L1459:
	;
	v7172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[47])))
	if v7172 != int32(1) {
		goto L1456
	} else {
		goto L1460
	}
L1460:
	;
	v7176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[48])))
	if v7176&int32(1) != 0 {
		goto L1456
	} else {
		goto L1461
	}
L1461:
	;
	v7180 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	if v7180 != 0 {
		goto L1462
	} else {
		goto L1463
	}
L1462:
	;
	F_process_startup_options(m, v7180, v6903)
	mBase = m.M
	v7182 = m.ExcPending
	if v7182 != 0 {
		goto L1
	} else {
		goto L1465
	}
L1463:
	;
	goto L1464
L1464:
	;
	v7184 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	if int32(0) < v7184 {
		goto L1466
	} else {
		goto L1467
	}
L1465:
	;
	goto L1464
L1466:
	;
	F_pg_usleep(m, v7184*int32(_a_F_InitPostgres_122))
	mBase = m.M
	v7190 = m.ExcPending
	if v7190 != 0 {
		goto L1
	} else {
		goto L1469
	}
L1467:
	;
	goto L1468
L1468:
	;
	F_InitializeClientEncoding(m)
	mBase = m.M
	v7192 = m.ExcPending
	if v7192 != 0 {
		goto L1
	} else {
		goto L1470
	}
L1469:
	;
	goto L1468
L1470:
	;
	goto L1427
L1471:
	;
	if v6900 != 0 {
		goto L1473
	} else {
		goto L1474
	}
L1472:
	;
	F_LockSharedObject(m, int32(1262), v7232, int32(3))
	mBase = m.M
	v7240 = m.ExcPending
	if v7240 != 0 {
		goto L1
	} else {
		goto L1488
	}
L1473:
	;
	v7196 = v6912 + int32(432)
	F_ScanKeyInit(m, v7196, int32(2), int32(3), int32(62), v6900)
	mBase = m.M
	v7201 = m.ExcPending
	if v7201 != 0 {
		goto L1
	} else {
		goto L1476
	}
L1474:
	;
	goto L1475
L1475:
	;
	if v6901 == int32(0) {
		goto L1427
	} else {
		goto L1487
	}
L1476:
	;
	v7205 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v7206 = m.ExcPending
	if v7206 != 0 {
		goto L1
	} else {
		goto L1477
	}
L1477:
	;
	v7209 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[79])))
	v7212 = F_systable_beginscan(m, v7205, int32(2671), v7209, int32(0), int32(1), v7196)
	mBase = m.M
	v7213 = m.ExcPending
	if v7213 != 0 {
		goto L1
	} else {
		goto L1478
	}
L1478:
	;
	v7214 = F_systable_getnext(m, v7212)
	mBase = m.M
	v7215 = m.ExcPending
	if v7215 != 0 {
		goto L1
	} else {
		goto L1479
	}
L1479:
	;
	if v7214 != 0 {
		goto L1480
	} else {
		goto L1481
	}
L1480:
	;
	v7216 = F_heap_copytuple(m, v7214)
	mBase = m.M
	v7217 = m.ExcPending
	if v7217 != 0 {
		goto L1
	} else {
		goto L1483
	}
L1481:
	;
	v7218 = int32(0)
	goto L1482
L1482:
	;
	F_systable_endscan(m, v7212)
	mBase = m.M
	v7220 = m.ExcPending
	if v7220 != 0 {
		goto L1
	} else {
		goto L1484
	}
L1483:
	;
	v7218 = v7216
	goto L1482
L1484:
	;
	F_relation_close(m, v7205, int32(1))
	mBase = m.M
	v7223 = m.ExcPending
	if v7223 != 0 {
		goto L1
	} else {
		goto L1485
	}
L1485:
	;
	if v7218 == int32(0) {
		goto L1431
	} else {
		goto L1486
	}
L1486:
	;
	v7226 = *(*int32)(unsafe.Add(mBase, uint32(v7218)+16))
	v7227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7226)+22)))
	v7229 = *(*int32)(unsafe.Add(mBase, uint32(v7226+v7227)))
	v7232 = v7229
	goto L1472
L1487:
	;
	v7232 = v6901
	goto L1472
L1488:
	;
	v7242 = v6912 + int32(432)
	F_ScanKeyInit(m, v7242, int32(1), int32(3), int32(184), v7232)
	mBase = m.M
	v7247 = m.ExcPending
	if v7247 != 0 {
		goto L1
	} else {
		goto L1489
	}
L1489:
	;
	v7250 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v7251 = m.ExcPending
	if v7251 != 0 {
		goto L1
	} else {
		goto L1490
	}
L1490:
	;
	v7254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[79])))
	v7257 = F_systable_beginscan(m, v7250, int32(2672), v7254, int32(0), int32(1), v7242)
	mBase = m.M
	v7258 = m.ExcPending
	if v7258 != 0 {
		goto L1
	} else {
		goto L1491
	}
L1491:
	;
	v7259 = F_systable_getnext(m, v7257)
	mBase = m.M
	v7260 = m.ExcPending
	if v7260 != 0 {
		goto L1
	} else {
		goto L1492
	}
L1492:
	;
	if v7259 != 0 {
		goto L1493
	} else {
		goto L1494
	}
L1493:
	;
	v7261 = F_heap_copytuple(m, v7259)
	mBase = m.M
	v7262 = m.ExcPending
	if v7262 != 0 {
		goto L1
	} else {
		goto L1496
	}
L1494:
	;
	v7263 = int32(0)
	goto L1495
L1495:
	;
	F_systable_endscan(m, v7257)
	mBase = m.M
	v7265 = m.ExcPending
	if v7265 != 0 {
		goto L1
	} else {
		goto L1497
	}
L1496:
	;
	v7263 = v7261
	goto L1495
L1497:
	;
	F_relation_close(m, v7250, int32(1))
	mBase = m.M
	v7268 = m.ExcPending
	if v7268 != 0 {
		goto L1
	} else {
		goto L1498
	}
L1498:
	;
	if v7263 != 0 {
		goto L1500
	} else {
		goto L1501
	}
L1499:
	;
	v7312 = v6912 + int32(432)
	v7314 = v7271 + int32(4)
	goto L1523
L1500:
	;
	v7269 = *(*int32)(unsafe.Add(mBase, uint32(v7263)+16))
	v7270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7269)+22)))
	v7271 = v7269 + v7270
	if v6900 == int32(0) {
		goto L1499
	} else {
		goto L1503
	}
L1501:
	;
	goto L1502
L1502:
	;
	if v6900 != 0 {
		goto L119
	} else {
		goto L1515
	}
L1503:
	;
	v7275 = v7271 + int32(4)
	if v7275|v6900 != 0 {
		goto L1505
	} else {
		goto L1506
	}
L1504:
	;
	if v7290 == int32(0) {
		goto L1499
	} else {
		goto L1514
	}
L1505:
	;
	v7281 = int32(-1)
	goto L1507
L1506:
	;
	v7281 = int32(0)
	goto L1507
L1507:
	;
	if v7275 != 0 {
		goto L1508
	} else {
		goto L1509
	}
L1508:
	;
	v7282 = int32(1)
	goto L1510
L1509:
	;
	v7282 = v7281
	goto L1510
L1510:
	;
	v7283 = int32(0)
	if base.B2i32(v7275 == v7283)|base.B2i32(v6900 == v7283) != 0 {
		goto L1511
	} else {
		goto L1512
	}
L1511:
	;
	v7290 = v7282
	goto L1513
L1512:
	;
	v7289 = F_strncmp(m, v7275, v6900, int32(64))
	mBase = m.M
	v7290 = v7289
	goto L1513
L1513:
	;
	goto L1504
L1514:
	;
	goto L119
L1515:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7296 = m.ExcPending
	if v7296 != 0 {
		goto L1
	} else {
		goto L1516
	}
L1516:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v7299 = m.ExcPending
	if v7299 != 0 {
		goto L1
	} else {
		goto L1517
	}
L1517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+240)) = v7232
	F_errmsg(m, int32(_a_F_InitPostgres_123), v6912+int32(240))
	mBase = m.M
	v7305 = m.ExcPending
	if v7305 != 0 {
		goto L1
	} else {
		goto L1518
	}
L1518:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1101), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7310 = m.ExcPending
	if v7310 != 0 {
		goto L1
	} else {
		goto L1519
	}
L1519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1520:
	;
	v7434 = *(*int32)(unsafe.Add(mBase, uint32(v7271)+80))
	goto L1551
L1521:
	;
	v7431 = F_strlen(m, v7420)
	mBase = m.M
	goto L1520
L1523:
	;
	goto L1524
L1524:
	;
	v7321 = int32(63)
	if (v7312^v7314)&int32(3) != 0 {
		goto L1528
	} else {
		goto L1529
	}
L1525:
	;
	v7424 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7421))) = uint8(v7424)
	goto L1521
L1526:
	;
	v7405 = v7400
	v7406 = v7401
	v7407 = v7402
	goto L1547
L1527:
	;
	if v7395 == int32(0) {
		v7420 = v7393
		v7421 = v7394
		goto L1525
	} else {
		goto L1546
	}
L1528:
	;
	v7393 = v7314
	v7394 = v7312
	v7395 = v7321
	goto L1527
L1529:
	;
	goto L1530
L1530:
	;
	v7325 = int32(0)
	if base.B2i32(v7314&int32(3) == v7325)|int32(0) == v7325 {
		goto L1532
	} else {
		goto L1533
	}
L1531:
	;
	if v7361 == int32(0) {
		v7420 = v7358
		v7421 = v7359
		goto L1525
	} else {
		goto L1540
	}
L1532:
	;
	v7337 = v7314
	v7338 = v7312
	v7339 = v7321
	goto L1535
L1533:
	;
	goto L1534
L1534:
	;
	v7358 = v7314
	v7359 = v7312
	v7360 = v7321
	v7361 = int32(1)
	goto L1531
L1535:
	;
	v7341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7337))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7338))) = uint8(v7341)
	if v7341 == int32(0) {
		v7400 = v7337
		v7401 = v7338
		v7402 = v7339
		goto L1526
	} else {
		goto L1537
	}
L1536:
	;
	v7358 = v7352
	v7359 = v7346
	v7360 = v7348
	v7361 = v7350
	goto L1531
L1537:
	;
	v7345 = int32(1)
	v7346 = v7338 + v7345
	v7348 = v7339 - v7345
	v7349 = int32(0)
	v7350 = base.B2i32(v7348 != v7349)
	v7352 = v7337 + v7345
	if v7352&int32(3) == v7349 {
		v7358 = v7352
		v7359 = v7346
		v7360 = v7348
		v7361 = v7350
		goto L1531
	} else {
		goto L1538
	}
L1538:
	;
	if v7348 != 0 {
		v7337 = v7352
		v7338 = v7346
		v7339 = v7348
		goto L1535
	} else {
		goto L1539
	}
L1539:
	;
	goto L1536
L1540:
	;
	v7364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7358))))
	if base.B2i32(v7364 == int32(0))|base.B2i32(base.Ui32(v7360) < base.Ui32(int32(4))) != 0 {
		v7393 = v7358
		v7394 = v7359
		v7395 = v7360
		goto L1527
	} else {
		goto L1541
	}
L1541:
	;
	v7371 = v7358
	v7372 = v7359
	v7373 = v7360
	goto L1542
L1542:
	;
	v7376 = *(*int32)(unsafe.Add(mBase, uint32(v7371)))
	v7379 = int32(-2139062144)
	if (int32(16843008)-v7376|v7376)&v7379 != v7379 {
		v7400 = v7371
		v7401 = v7372
		v7402 = v7373
		goto L1526
	} else {
		goto L1544
	}
L1543:
	;
	v7393 = v7387
	v7394 = v7385
	v7395 = v7389
	goto L1527
L1544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7372))) = v7376
	v7384 = int32(4)
	v7385 = v7372 + v7384
	v7387 = v7371 + v7384
	v7389 = v7373 - v7384
	if base.Ui32(int32(3)) < base.Ui32(v7389) {
		v7371 = v7387
		v7372 = v7385
		v7373 = v7389
		goto L1542
	} else {
		goto L1545
	}
L1545:
	;
	goto L1543
L1546:
	;
	v7400 = v7393
	v7401 = v7394
	v7402 = v7395
	goto L1526
L1547:
	;
	v7409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7405))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7406))) = uint8(v7409)
	if v7409 == int32(0) {
		v7420 = v7405
		v7421 = v7406
		goto L1525
	} else {
		goto L1549
	}
L1548:
	;
	v7420 = v7416
	v7421 = v7414
	goto L1525
L1549:
	;
	v7413 = int32(1)
	v7414 = v7406 + v7413
	v7416 = v7405 + v7413
	v7418 = v7407 - v7413
	if v7418 != 0 {
		v7405 = v7416
		v7406 = v7414
		v7407 = v7418
		goto L1547
	} else {
		goto L1550
	}
L1550:
	;
	goto L1548
L1551:
	;
	if v7434 == int32(-2) {
		goto L1429
	} else {
		goto L1552
	}
L1552:
	;
	v7438 = *(*int32)(unsafe.Add(mBase, uint32(v7271)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[72])) = v7438
	v7441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7271)+79)))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[80])) = uint8(v7441)
	if v6905 == int32(0) {
		v7643 = v7232
		goto L1428
	} else {
		goto L1553
	}
L1553:
	;
	if (v7312^v6905)&int32(3) != 0 {
		goto L1557
	} else {
		goto L1558
	}
L1554:
	;
	v7643 = v7232
	goto L1428
L1555:
	;
	goto L1554
L1556:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7499))) = uint8(v7498)
	if v7498&int32(255) == int32(0) {
		goto L1555
	} else {
		goto L1571
	}
L1557:
	;
	v7450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7312))))
	v7497 = v7312
	v7498 = v7450
	v7499 = v6905
	goto L1556
L1558:
	;
	goto L1559
L1559:
	;
	if v7312&int32(3) != 0 {
		goto L1560
	} else {
		goto L1561
	}
L1560:
	;
	v7454 = v7312
	v7456 = v6905
	goto L1563
L1561:
	;
	v7468 = v7312
	v7470 = v6905
	goto L1562
L1562:
	;
	v7472 = *(*int32)(unsafe.Add(mBase, uint32(v7468)))
	v7475 = int32(-2139062144)
	if (int32(16843008)-v7472|v7472)&v7475 != v7475 {
		v7497 = v7468
		v7498 = v7472
		v7499 = v7470
		goto L1556
	} else {
		goto L1567
	}
L1563:
	;
	v7457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7454))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7456))) = uint8(v7457)
	if v7457 == int32(0) {
		goto L1555
	} else {
		goto L1565
	}
L1564:
	;
	v7468 = v7464
	v7470 = v7462
	goto L1562
L1565:
	;
	v7461 = int32(1)
	v7462 = v7456 + v7461
	v7464 = v7454 + v7461
	if v7464&int32(3) != 0 {
		v7454 = v7464
		v7456 = v7462
		goto L1563
	} else {
		goto L1566
	}
L1566:
	;
	goto L1564
L1567:
	;
	v7480 = v7468
	v7481 = v7472
	v7482 = v7470
	goto L1568
L1568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7482))) = v7481
	v7484 = int32(4)
	v7485 = v7482 + v7484
	v7487 = v7480 + v7484
	v7489 = *(*int32)(unsafe.Add(mBase, uint32(v7480)+4))
	v7492 = int32(-2139062144)
	if (int32(16843008)-v7489|v7489)&v7492 == v7492 {
		v7480 = v7487
		v7481 = v7489
		v7482 = v7485
		goto L1568
	} else {
		goto L1570
	}
L1569:
	;
	v7497 = v7487
	v7498 = v7489
	v7499 = v7485
	goto L1556
L1570:
	;
	goto L1569
L1571:
	;
	v7506 = v7497
	v7508 = v7499
	goto L1572
L1572:
	;
	v7509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7506)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7508)+1)) = uint8(v7509)
	v7511 = int32(1)
	if v7509 != 0 {
		v7506 = v7506 + v7511
		v7508 = v7508 + v7511
		goto L1572
	} else {
		goto L1574
	}
L1573:
	;
	goto L1555
L1574:
	;
	goto L1573
L1575:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v7525 = m.ExcPending
	if v7525 != 0 {
		goto L1
	} else {
		goto L1576
	}
L1576:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_124), int32(0))
	mBase = m.M
	v7529 = m.ExcPending
	if v7529 != 0 {
		goto L1
	} else {
		goto L1577
	}
L1577:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(940), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7534 = m.ExcPending
	if v7534 != 0 {
		goto L1
	} else {
		goto L1578
	}
L1578:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1579:
	;
	F_errcode(m, int32(_a_F_InitPostgres_125))
	mBase = m.M
	v7541 = m.ExcPending
	if v7541 != 0 {
		goto L1
	} else {
		goto L1580
	}
L1580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+320)) = int32(_a_F_InitPostgres_126)
	F_errmsg(m, int32(_a_F_InitPostgres_127), v6912+int32(320))
	mBase = m.M
	v7548 = m.ExcPending
	if v7548 != 0 {
		goto L1
	} else {
		goto L1581
	}
L1581:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(961), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7553 = m.ExcPending
	if v7553 != 0 {
		goto L1
	} else {
		goto L1582
	}
L1582:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1583:
	;
	F_errcode(m, int32(_a_F_InitPostgres_125))
	mBase = m.M
	v7560 = m.ExcPending
	if v7560 != 0 {
		goto L1
	} else {
		goto L1584
	}
L1584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+336)) = int32(_a_F_InitPostgres_128)
	F_errmsg(m, int32(_a_F_InitPostgres_129), v6912+int32(336))
	mBase = m.M
	v7567 = m.ExcPending
	if v7567 != 0 {
		goto L1
	} else {
		goto L1585
	}
L1585:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(967), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7572 = m.ExcPending
	if v7572 != 0 {
		goto L1
	} else {
		goto L1586
	}
L1586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1587:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v7579 = m.ExcPending
	if v7579 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1588:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_130), int32(0))
	mBase = m.M
	v7583 = m.ExcPending
	if v7583 != 0 {
		goto L1
	} else {
		goto L1589
	}
L1589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+304)) = int32(_a_F_InitPostgres_131)
	F_errdetail(m, int32(_a_F_InitPostgres_132), v6912+int32(304))
	mBase = m.M
	v7590 = m.ExcPending
	if v7590 != 0 {
		goto L1
	} else {
		goto L1590
	}
L1590:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(980), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7595 = m.ExcPending
	if v7595 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1591:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1592:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v7602 = m.ExcPending
	if v7602 != 0 {
		goto L1
	} else {
		goto L1593
	}
L1593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+288)) = v6900
	F_errmsg(m, int32(_a_F_InitPostgres_133), v6912+int32(288))
	mBase = m.M
	v7608 = m.ExcPending
	if v7608 != 0 {
		goto L1
	} else {
		goto L1594
	}
L1594:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1032), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7613 = m.ExcPending
	if v7613 != 0 {
		goto L1
	} else {
		goto L1595
	}
L1595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1596:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v7624 = m.ExcPending
	if v7624 != 0 {
		goto L1
	} else {
		goto L1597
	}
L1597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+272)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_134), v6912+int32(272))
	mBase = m.M
	v7632 = m.ExcPending
	if v7632 != 0 {
		goto L1
	} else {
		goto L1598
	}
L1598:
	;
	F_errhint(m, int32(_a_F_InitPostgres_135), int32(0))
	mBase = m.M
	v7636 = m.ExcPending
	if v7636 != 0 {
		goto L1
	} else {
		goto L1599
	}
L1599:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1111), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7641 = m.ExcPending
	if v7641 != 0 {
		goto L1
	} else {
		goto L1600
	}
L1600:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1601:
	;
	v7656 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v7658 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[72]))
	v7659 = F_GetDatabasePath(m, v7656, v7658)
	mBase = m.M
	v7660 = m.ExcPending
	if v7660 != 0 {
		goto L1
	} else {
		goto L1602
	}
L1602:
	;
	if v6924 != 0 {
		goto L1604
	} else {
		goto L1605
	}
L1603:
	;
	v8279 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	if v8279 != 0 {
		goto L1747
	} else {
		goto L1748
	}
L1604:
	;
	v7662 = F_access(m, v7659, int32(0))
	mBase = m.M
	if v7662 == int32(-1) {
		goto L1607
	} else {
		goto L1608
	}
L1605:
	;
	goto L1606
L1606:
	;
	F_SetDatabasePath(m, v7659)
	mBase = m.M
	v8230 = m.ExcPending
	if v8230 != 0 {
		goto L1
	} else {
		goto L1743
	}
L1607:
	;
	v7666 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7670 = m.ExcPending
	if v7670 != 0 {
		goto L1
	} else {
		goto L1610
	}
L1608:
	;
	goto L1609
L1609:
	;
	F_ValidatePgVersion(m, v7659)
	mBase = m.M
	v7687 = m.ExcPending
	if v7687 != 0 {
		goto L1
	} else {
		goto L1615
	}
L1610:
	;
	if v7666 == int32(44) {
		goto L127
	} else {
		goto L1611
	}
L1611:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v7674 = m.ExcPending
	if v7674 != 0 {
		goto L1
	} else {
		goto L1612
	}
L1612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+32)) = v7659
	F_errmsg(m, int32(_a_F_InitPostgres_136), v6912+int32(32))
	mBase = m.M
	v7680 = m.ExcPending
	if v7680 != 0 {
		goto L1
	} else {
		goto L1613
	}
L1613:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1177), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7685 = m.ExcPending
	if v7685 != 0 {
		goto L1
	} else {
		goto L1614
	}
L1614:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1615:
	;
	F_SetDatabasePath(m, v7659)
	mBase = m.M
	v7689 = m.ExcPending
	if v7689 != 0 {
		goto L1
	} else {
		goto L1616
	}
L1616:
	;
	F_pfree(m, v7659)
	mBase = m.M
	v7691 = m.ExcPending
	if v7691 != 0 {
		goto L1
	} else {
		goto L1617
	}
L1617:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v7693 = m.ExcPending
	if v7693 != 0 {
		goto L1
	} else {
		goto L1618
	}
L1618:
	;
	F_initialize_acl(m)
	mBase = m.M
	v7695 = m.ExcPending
	if v7695 != 0 {
		goto L1
	} else {
		goto L1619
	}
L1619:
	;
	v7698 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v7699 = F_SearchSysCache1(m, int32(21), v7698)
	mBase = m.M
	v7700 = m.ExcPending
	if v7700 != 0 {
		goto L1
	} else {
		goto L1620
	}
L1620:
	;
	if v7699 == int32(0) {
		goto L126
	} else {
		goto L1621
	}
L1621:
	;
	v7704 = v6912 + int32(432)
	v7705 = *(*int32)(unsafe.Add(mBase, uint32(v7699)+16))
	v7706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7705)+22)))
	v7707 = v7705 + v7706
	v7709 = v7707 + int32(4)
	v7712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7704))))
	v7715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7709))))
	if base.B2i32(v7712 == int32(0))|base.B2i32(v7712 != v7715) != 0 {
		v7733 = v7712
		v7734 = v7715
		goto L1623
	} else {
		goto L1624
	}
L1622:
	;
	if v7733-v7734 != 0 {
		goto L125
	} else {
		goto L1629
	}
L1623:
	;
	goto L1622
L1624:
	;
	v7718 = v7704
	v7719 = v7709
	goto L1625
L1625:
	;
	v7722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7719)+1)))
	v7723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7718)+1)))
	if v7723 == int32(0) {
		v7733 = v7723
		v7734 = v7722
		goto L1623
	} else {
		goto L1627
	}
L1626:
	;
	v7733 = v7723
	v7734 = v7722
	goto L1623
L1627:
	;
	v7726 = int32(1)
	if v7723 == v7722 {
		v7718 = v7718 + v7726
		v7719 = v7719 + v7726
		goto L1625
	} else {
		goto L1628
	}
L1628:
	;
	goto L1626
L1629:
	;
	v7737 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v7737 != int32(1) {
		goto L1630
	} else {
		goto L1631
	}
L1630:
	;
	v7943 = *(*int32)(unsafe.Add(mBase, uint32(v7707)+72))
	v7944 = m.G0
	v7946 = v7944 - int32(16)
	m.G0 = v7946
	if base.Ui32(int32(35)) <= base.Ui32(v7943) {
		goto L1658
	} else {
		goto L1659
	}
L1631:
	;
	v7741 = v6904 & int32(2)
	if v7741 == int32(0) {
		goto L1632
	} else {
		goto L1633
	}
L1632:
	;
	v7744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7707)+78)))
	if v7744&int32(1) == int32(0) {
		goto L124
	} else {
		goto L1635
	}
L1633:
	;
	goto L1634
L1634:
	;
	v7749 = int32(0)
	if base.B2i32(v7741 != v7749)|v6903 == v7749 {
		goto L1636
	} else {
		goto L1637
	}
L1635:
	;
	goto L1634
L1636:
	;
	v7756 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v7758 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	v7760 = F_object_aclcheck(m, int32(1262), v7756, v7758, int64(2048))
	mBase = m.M
	v7761 = m.ExcPending
	if v7761 != 0 {
		goto L1
	} else {
		goto L1639
	}
L1637:
	;
	goto L1638
L1638:
	;
	v7763 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v7766 = *(*int32)(unsafe.Add(mBase, uint32(v7707)+80))
	if v6903|(base.B2i32(v7763 != int32(1))|base.B2i32(v7766 < int32(0))) != 0 {
		goto L1630
	} else {
		goto L1641
	}
L1639:
	;
	if v7760 != 0 {
		goto L123
	} else {
		goto L1640
	}
L1640:
	;
	goto L1638
L1641:
	;
	v7772 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v7773 = int32(0)
	v7775 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[81]))
	v7777 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[82]))
	v7781 = F_LWLockAcquire(m, v7777+int32(512), int32(1))
	mBase = m.M
	v7782 = m.ExcPending
	if v7782 != 0 {
		goto L1
	} else {
		goto L1642
	}
L1642:
	;
	v7783 = *(*int32)(unsafe.Add(mBase, uint32(v7775)))
	if int32(0) < v7783 {
		goto L1643
	} else {
		goto L1644
	}
L1643:
	;
	v7789 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[83]))
	v7791 = int32(0)
	v7793 = v7773
	goto L1646
L1644:
	;
	v7855 = v7773
	goto L1645
L1645:
	;
	v7895 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[82]))
	F_LWLockRelease(m, v7895+int32(512))
	mBase = m.M
	v7899 = m.ExcPending
	if v7899 != 0 {
		goto L1
	} else {
		goto L1656
	}
L1646:
	;
	v7835 = *(*int32)(unsafe.Add(mBase, uint32(v7775+int32(36)+v7791<<(uint(int32(2))%32))))
	v7838 = v7789 + v7835*int32(640)
	v7839 = *(*int32)(unsafe.Add(mBase, uint32(v7838)+44))
	if v7839 == int32(0) {
		v7849 = v7793
		goto L1648
	} else {
		goto L1649
	}
L1647:
	;
	v7855 = v7849
	goto L1645
L1648:
	;
	v7851 = v7791 + int32(1)
	if v7851 != v7783 {
		v7791 = v7851
		v7793 = v7849
		goto L1646
	} else {
		goto L1655
	}
L1649:
	;
	v7842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7838)+72)))
	if v7842 != int32(1) {
		v7849 = v7793
		goto L1648
	} else {
		goto L1650
	}
L1650:
	;
	if v7772 != 0 {
		goto L1651
	} else {
		goto L1652
	}
L1651:
	;
	v7845 = *(*int32)(unsafe.Add(mBase, uint32(v7838)+60))
	if v7845 != v7772 {
		v7849 = v7793
		goto L1648
	} else {
		goto L1654
	}
L1652:
	;
	goto L1653
L1653:
	;
	v7849 = v7793 + int32(1)
	goto L1648
L1654:
	;
	goto L1653
L1655:
	;
	goto L1647
L1656:
	;
	v7900 = *(*int32)(unsafe.Add(mBase, uint32(v7707)+80))
	if v7900 < v7855 {
		goto L122
	} else {
		goto L1657
	}
L1657:
	;
	goto L1630
L1658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7953 = m.ExcPending
	if v7953 != 0 {
		goto L1
	} else {
		goto L1661
	}
L1659:
	;
	goto L1660
L1660:
	;
	v7965 = v7943 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84])) = v7965 + int32(_a_F_InitPostgres_137)
	m.G0 = v7946 + int32(16)
	v7975 = *(*int32)(unsafe.Add(mBase, uint32(v7965)+uint32(_c_F_InitPostgres[85])))
	goto L1664
L1661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7946))) = v7943
	F_errmsg_internal(m, int32(_a_F_InitPostgres_138), v7946)
	mBase = m.M
	v7957 = m.ExcPending
	if v7957 != 0 {
		goto L1
	} else {
		goto L1662
	}
L1662:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_139), int32(1290), int32(_a_F_InitPostgres_140))
	mBase = m.M
	v7962 = m.ExcPending
	if v7962 != 0 {
		goto L1
	} else {
		goto L1663
	}
L1663:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1664:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_141), v7975, int32(0), int32(1))
	mBase = m.M
	v7979 = m.ExcPending
	if v7979 != 0 {
		goto L1
	} else {
		goto L1665
	}
L1665:
	;
	v7982 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84]))
	v7983 = *(*int32)(unsafe.Add(mBase, uint32(v7982)))
	goto L1666
L1666:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_142), v7983, int32(4), int32(1))
	mBase = m.M
	v7987 = m.ExcPending
	if v7987 != 0 {
		goto L1
	} else {
		goto L1667
	}
L1667:
	;
	v7990 = F_SysCacheGetAttrNotNull(m, int32(21), v7699, int32(13))
	mBase = m.M
	v7991 = m.ExcPending
	if v7991 != 0 {
		goto L1
	} else {
		goto L1668
	}
L1668:
	;
	v7992 = F_text_to_cstring(m, v7990)
	mBase = m.M
	v7993 = m.ExcPending
	if v7993 != 0 {
		goto L1
	} else {
		goto L1669
	}
L1669:
	;
	v7996 = F_SysCacheGetAttrNotNull(m, int32(21), v7699, int32(14))
	mBase = m.M
	v7997 = m.ExcPending
	if v7997 != 0 {
		goto L1
	} else {
		goto L1670
	}
L1670:
	;
	v7998 = F_text_to_cstring(m, v7996)
	mBase = m.M
	v7999 = m.ExcPending
	if v7999 != 0 {
		goto L1
	} else {
		goto L1671
	}
L1671:
	;
	v8001 = F_pg_perm_setlocale(m, int32(3), v7992)
	mBase = m.M
	v8002 = m.ExcPending
	if v8002 != 0 {
		goto L1
	} else {
		goto L1672
	}
L1672:
	;
	if v8001 == int32(0) {
		goto L121
	} else {
		goto L1673
	}
L1673:
	;
	v8006 = F_pg_perm_setlocale(m, int32(0), v7998)
	mBase = m.M
	v8007 = m.ExcPending
	if v8007 != 0 {
		goto L1
	} else {
		goto L1674
	}
L1674:
	;
	if v8006 == int32(0) {
		goto L120
	} else {
		goto L1675
	}
L1675:
	;
	v8010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7998))))
	if v8010 == int32(67) {
		goto L1678
	} else {
		goto L1679
	}
L1676:
	;
	v8046 = m.G0
	v8048 = v8046 - int32(32)
	m.G0 = v8048
	v8052 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v8053 = F_SearchSysCache1(m, int32(21), v8052)
	mBase = m.M
	v8054 = m.ExcPending
	if v8054 != 0 {
		goto L1
	} else {
		goto L1691
	}
L1677:
	;
	v8044 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[86])) = uint8(v8044)
	goto L1676
L1678:
	;
	v8013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7998)+1)))
	if v8013 == int32(0) {
		goto L1677
	} else {
		goto L1681
	}
L1679:
	;
	goto L1680
L1680:
	;
	v8016 = int32(_a_F_InitPostgres_143)
	v8019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7998))))
	v8022 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[87])))
	if base.B2i32(v8019 == int32(0))|base.B2i32(v8019 != v8022) != 0 {
		v8040 = v8019
		v8041 = v8022
		goto L1683
	} else {
		goto L1684
	}
L1681:
	;
	goto L1680
L1682:
	;
	if v8040-v8041 != 0 {
		goto L1676
	} else {
		goto L1689
	}
L1683:
	;
	goto L1682
L1684:
	;
	v8025 = v7998
	v8026 = v8016
	goto L1685
L1685:
	;
	v8029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8026)+1)))
	v8030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8025)+1)))
	if v8030 == int32(0) {
		v8040 = v8030
		v8041 = v8029
		goto L1683
	} else {
		goto L1687
	}
L1686:
	;
	v8040 = v8030
	v8041 = v8029
	goto L1683
L1687:
	;
	v8033 = int32(1)
	if v8030 == v8029 {
		v8025 = v8025 + v8033
		v8026 = v8026 + v8033
		goto L1685
	} else {
		goto L1688
	}
L1688:
	;
	goto L1686
L1689:
	;
	goto L1677
L1690:
	;
	v8122 = F_SysCacheGetAttr(m, int32(21), v7699, int32(17), v6912+int32(511))
	mBase = m.M
	v8123 = m.ExcPending
	if v8123 != 0 {
		goto L1
	} else {
		goto L1710
	}
L1691:
	;
	if v8053 != 0 {
		goto L1692
	} else {
		goto L1693
	}
L1692:
	;
	v8055 = *(*int32)(unsafe.Add(mBase, uint32(v8053)+16))
	v8056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8055)+22)))
	v8057 = v8055 + v8056
	v8058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8057)+76)))
	switch v8058 - int32(98) {
	case 0:
		goto L1696
	case 1:
		goto L1698
	default:
		goto L1697
	case 7:
		goto L1699
	}
L1693:
	;
	goto L1694
L1694:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8106 = m.ExcPending
	if v8106 != 0 {
		goto L1
	} else {
		goto L1707
	}
L1695:
	;
	v8094 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8093)+4)) = uint8(v8094)
	F_ReleaseCatCache(m, v8053)
	mBase = m.M
	v8097 = m.ExcPending
	if v8097 != 0 {
		goto L1
	} else {
		goto L1706
	}
L1696:
	;
	v8090 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8091 = F_create_pg_locale_builtin(m, int32(100), v8090)
	mBase = m.M
	v8092 = m.ExcPending
	if v8092 != 0 {
		goto L1
	} else {
		goto L1705
	}
L1697:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8073 = m.ExcPending
	if v8073 != 0 {
		goto L1
	} else {
		goto L1702
	}
L1698:
	;
	v8067 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8068 = F_create_pg_locale_libc(m, int32(100), v8067)
	mBase = m.M
	v8069 = m.ExcPending
	if v8069 != 0 {
		goto L1
	} else {
		goto L1701
	}
L1699:
	;
	v8063 = F_create_pg_locale_icu(m)
	mBase = m.M
	v8064 = m.ExcPending
	if v8064 != 0 {
		goto L1
	} else {
		goto L1700
	}
L1700:
	;
	v8093 = v8063
	goto L1695
L1701:
	;
	v8093 = v8068
	goto L1695
L1702:
	;
	v8074 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8057)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v8048)+20)) = v8074
	*(*int32)(unsafe.Add(mBase, uint32(v8048)+16)) = int32(_a_F_InitPostgres_144)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_145), v8048+int32(16))
	mBase = m.M
	v8082 = m.ExcPending
	if v8082 != 0 {
		goto L1
	} else {
		goto L1703
	}
L1703:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_146), int32(1179), int32(_a_F_InitPostgres_144))
	mBase = m.M
	v8087 = m.ExcPending
	if v8087 != 0 {
		goto L1
	} else {
		goto L1704
	}
L1704:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1705:
	;
	v8093 = v8091
	goto L1695
L1706:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[88])) = v8093
	m.G0 = v8048 + int32(32)
	goto L1690
L1707:
	;
	v8108 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	*(*int32)(unsafe.Add(mBase, uint32(v8048))) = v8108
	F_errmsg_internal(m, int32(_a_F_InitPostgres_147), v8048)
	mBase = m.M
	v8112 = m.ExcPending
	if v8112 != 0 {
		goto L1
	} else {
		goto L1708
	}
L1708:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_146), int32(1165), int32(_a_F_InitPostgres_144))
	mBase = m.M
	v8117 = m.ExcPending
	if v8117 != 0 {
		goto L1
	} else {
		goto L1709
	}
L1709:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1710:
	;
	v8124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6912)+511)))
	if v8124 != 0 {
		goto L1711
	} else {
		goto L1712
	}
L1711:
	;
	F_ReleaseCatCache(m, v7699)
	mBase = m.M
	v8228 = m.ExcPending
	if v8228 != 0 {
		goto L1
	} else {
		goto L1742
	}
L1712:
	;
	v8125 = F_text_to_cstring(m, v8122)
	mBase = m.M
	v8126 = m.ExcPending
	if v8126 != 0 {
		goto L1
	} else {
		goto L1713
	}
L1713:
	;
	v8128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7707)+76)))
	if v8128 != int32(99) {
		goto L1715
	} else {
		goto L1716
	}
L1714:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), v8220, int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8223 = m.ExcPending
	if v8223 != 0 {
		goto L1
	} else {
		goto L1741
	}
L1715:
	;
	v8133 = F_SysCacheGetAttrNotNull(m, int32(21), v7699, int32(15))
	mBase = m.M
	v8134 = m.ExcPending
	if v8134 != 0 {
		goto L1
	} else {
		goto L1718
	}
L1716:
	;
	v8139 = v7992
	v8140 = int32(99)
	goto L1717
L1717:
	;
	v8142 = F_get_collation_actual_version(m, base.I32_extend8_s(v8140), v8139)
	mBase = m.M
	v8143 = m.ExcPending
	if v8143 != 0 {
		goto L1
	} else {
		goto L1720
	}
L1718:
	;
	v8135 = F_text_to_cstring(m, v8133)
	mBase = m.M
	v8136 = m.ExcPending
	if v8136 != 0 {
		goto L1
	} else {
		goto L1719
	}
L1719:
	;
	v8137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7707)+76)))
	v8139 = v8135
	v8140 = v8137
	goto L1717
L1720:
	;
	if v8142 == int32(0) {
		goto L1721
	} else {
		goto L1722
	}
L1721:
	;
	v8148 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v8149 = m.ExcPending
	if v8149 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1722:
	;
	goto L1723
L1723:
	;
	v8163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8142))))
	v8166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8125))))
	if base.B2i32(v8163 == int32(0))|base.B2i32(v8163 != v8166) != 0 {
		v8184 = v8163
		v8185 = v8166
		goto L1728
	} else {
		goto L1729
	}
L1724:
	;
	if v8148 == int32(0) {
		goto L1711
	} else {
		goto L1725
	}
L1725:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+96)) = v6912 + int32(432)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_149), v6912+int32(96))
	mBase = m.M
	v8159 = m.ExcPending
	if v8159 != 0 {
		goto L1
	} else {
		goto L1726
	}
L1726:
	;
	v8220 = int32(468)
	goto L1714
L1727:
	;
	if v8184-v8185 == int32(0) {
		goto L1711
	} else {
		goto L1734
	}
L1728:
	;
	goto L1727
L1729:
	;
	v8169 = v8142
	v8170 = v8125
	goto L1730
L1730:
	;
	v8173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8170)+1)))
	v8174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8169)+1)))
	if v8174 == int32(0) {
		v8184 = v8174
		v8185 = v8173
		goto L1728
	} else {
		goto L1732
	}
L1731:
	;
	v8184 = v8174
	v8185 = v8173
	goto L1728
L1732:
	;
	v8177 = int32(1)
	if v8174 == v8173 {
		v8169 = v8169 + v8177
		v8170 = v8170 + v8177
		goto L1730
	} else {
		goto L1733
	}
L1733:
	;
	goto L1731
L1734:
	;
	v8191 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v8192 = m.ExcPending
	if v8192 != 0 {
		goto L1
	} else {
		goto L1735
	}
L1735:
	;
	if v8191 == int32(0) {
		goto L1711
	} else {
		goto L1736
	}
L1736:
	;
	v8196 = v6912 + int32(432)
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+144)) = v8196
	F_errmsg(m, int32(_a_F_InitPostgres_150), v6912+int32(144))
	mBase = m.M
	v8202 = m.ExcPending
	if v8202 != 0 {
		goto L1
	} else {
		goto L1737
	}
L1737:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+132)) = v8142
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+128)) = v8125
	F_errdetail(m, int32(_a_F_InitPostgres_151), v6912+int32(128))
	mBase = m.M
	v8209 = m.ExcPending
	if v8209 != 0 {
		goto L1
	} else {
		goto L1738
	}
L1738:
	;
	v8210 = F_quote_identifier(m, v8196)
	mBase = m.M
	v8211 = m.ExcPending
	if v8211 != 0 {
		goto L1
	} else {
		goto L1739
	}
L1739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+112)) = v8210
	F_errhint(m, int32(_a_F_InitPostgres_152), v6912+int32(112))
	mBase = m.M
	v8217 = m.ExcPending
	if v8217 != 0 {
		goto L1
	} else {
		goto L1740
	}
L1740:
	;
	v8220 = int32(479)
	goto L1714
L1741:
	;
	goto L1711
L1742:
	;
	goto L1603
L1743:
	;
	F_pfree(m, v7659)
	mBase = m.M
	v8232 = m.ExcPending
	if v8232 != 0 {
		goto L1
	} else {
		goto L1744
	}
L1744:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v8234 = m.ExcPending
	if v8234 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1745:
	;
	F_initialize_acl(m)
	mBase = m.M
	v8236 = m.ExcPending
	if v8236 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	goto L1603
L1747:
	;
	F_process_startup_options(m, v8279, v6903)
	mBase = m.M
	v8281 = m.ExcPending
	if v8281 != 0 {
		goto L1
	} else {
		goto L1750
	}
L1748:
	;
	goto L1749
L1749:
	;
	v8283 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v8285 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[89]))
	v8287 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v8287 == int32(1) {
		goto L1751
	} else {
		goto L1752
	}
L1750:
	;
	goto L1749
L1751:
	;
	v8292 = F_table_open(m, int32(2964), int32(1))
	mBase = m.M
	v8293 = m.ExcPending
	if v8293 != 0 {
		goto L1
	} else {
		goto L1754
	}
L1752:
	;
	goto L1753
L1753:
	;
	v8323 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	if int32(0) < v8323 {
		goto L1763
	} else {
		goto L1764
	}
L1754:
	;
	v8295 = F_GetCatalogSnapshot(m, int32(2964))
	mBase = m.M
	v8296 = m.ExcPending
	if v8296 != 0 {
		goto L1
	} else {
		goto L1755
	}
L1755:
	;
	v8297 = F_RegisterSnapshot(m, v8295)
	mBase = m.M
	v8298 = m.ExcPending
	if v8298 != 0 {
		goto L1
	} else {
		goto L1756
	}
L1756:
	;
	F_ApplySetting(m, v8297, v8283, v8285, v8292, int32(8))
	mBase = m.M
	v8301 = m.ExcPending
	if v8301 != 0 {
		goto L1
	} else {
		goto L1757
	}
L1757:
	;
	F_ApplySetting(m, v8297, int32(0), v8285, v8292, int32(7))
	mBase = m.M
	v8305 = m.ExcPending
	if v8305 != 0 {
		goto L1
	} else {
		goto L1758
	}
L1758:
	;
	F_ApplySetting(m, v8297, v8283, int32(0), v8292, int32(6))
	mBase = m.M
	v8309 = m.ExcPending
	if v8309 != 0 {
		goto L1
	} else {
		goto L1759
	}
L1759:
	;
	v8310 = int32(0)
	F_ApplySetting(m, v8297, v8310, v8310, v8292, int32(5))
	mBase = m.M
	v8314 = m.ExcPending
	if v8314 != 0 {
		goto L1
	} else {
		goto L1760
	}
L1760:
	;
	F_UnregisterSnapshot(m, v8297)
	mBase = m.M
	v8316 = m.ExcPending
	if v8316 != 0 {
		goto L1
	} else {
		goto L1761
	}
L1761:
	;
	F_relation_close(m, v8292, int32(1))
	mBase = m.M
	v8319 = m.ExcPending
	if v8319 != 0 {
		goto L1
	} else {
		goto L1762
	}
L1762:
	;
	goto L1753
L1763:
	;
	F_pg_usleep(m, v8323*int32(_a_F_InitPostgres_122))
	mBase = m.M
	v8329 = m.ExcPending
	if v8329 != 0 {
		goto L1
	} else {
		goto L1766
	}
L1764:
	;
	goto L1765
L1765:
	;
	v8330 = m.G0
	v8332 = v8330 - int32(16)
	m.G0 = v8332
	v8335 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v8335 == int32(0) {
		goto L1768
	} else {
		goto L1769
	}
L1766:
	;
	goto L1765
L1767:
	;
	m.G0 = v8332 + int32(16)
	F_InitializeClientEncoding(m)
	mBase = m.M
	v8420 = m.ExcPending
	if v8420 != 0 {
		goto L1
	} else {
		goto L1776
	}
L1768:
	;
	v8338 = int32(_a_F_InitPostgres_8)
	v8339 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v8342 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v8342
	v8344 = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v8332)+12)) = v8344
	*(*int32)(unsafe.Add(mBase, uint32(v8332)+8)) = v8344
	v8351 = F_list_make1_impl(m, int32(472), v8332+int32(8))
	mBase = m.M
	v8352 = m.ExcPending
	if v8352 != 0 {
		goto L1
	} else {
		goto L1771
	}
L1769:
	;
	goto L1770
L1770:
	;
	F_CacheRegisterSyscacheCallback(m, int32(38), int32(470), int32(0))
	mBase = m.M
	v8392 = m.ExcPending
	if v8392 != 0 {
		goto L1
	} else {
		goto L1772
	}
L1771:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v8339
	v8355 = int32(_a_F_InitPostgres_153)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[90])) = v8351
	v8357 = int32(_a_F_InitPostgres_154)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[91])) = int32(11)
	v8360 = int32(_a_F_InitPostgres_155)
	v8361 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[92])) = uint8(v8361)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[93])) = uint8(v8361)
	v8368 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[94])) = v8368
	v8372 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[90]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[95])) = v8372
	v8376 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[91]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[96])) = v8376
	v8380 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[92])))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[97])) = uint8(v8380)
	v8382 = int32(_a_F_InitPostgres_156)
	v8384 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[98]))
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[98])) = v8384 + int64(1)
	goto L1767
L1772:
	;
	F_CacheRegisterSyscacheCallback(m, int32(11), int32(470), int32(0))
	mBase = m.M
	v8397 = m.ExcPending
	if v8397 != 0 {
		goto L1
	} else {
		goto L1773
	}
L1773:
	;
	F_CacheRegisterSyscacheCallback(m, int32(9), int32(470), int32(0))
	mBase = m.M
	v8402 = m.ExcPending
	if v8402 != 0 {
		goto L1
	} else {
		goto L1774
	}
L1774:
	;
	F_CacheRegisterSyscacheCallback(m, int32(21), int32(470), int32(0))
	mBase = m.M
	v8407 = m.ExcPending
	if v8407 != 0 {
		goto L1
	} else {
		goto L1775
	}
L1775:
	;
	v8409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[93])) = uint8(v8409)
	v8412 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[99])) = uint8(v8412)
	goto L1767
L1776:
	;
	v8423 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8425 = F_MemoryContextAllocZero(m, v8423, int32(20))
	mBase = m.M
	v8426 = m.ExcPending
	if v8426 != 0 {
		goto L1
	} else {
		goto L1777
	}
L1777:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[100])) = v8425
	if v6904&int32(1) != 0 {
		goto L1778
	} else {
		goto L1779
	}
L1778:
	;
	v8431 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[101]))
	F_load_libraries(m, v8431, int32(_a_F_InitPostgres_157), int32(0))
	mBase = m.M
	v8435 = m.ExcPending
	if v8435 != 0 {
		goto L1
	} else {
		goto L1781
	}
L1779:
	;
	goto L1780
L1780:
	;
	if v6924 == int32(0) {
		v8501 = v6912
		goto L128
	} else {
		goto L1783
	}
L1781:
	;
	v8437 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[102]))
	F_load_libraries(m, v8437, int32(_a_F_InitPostgres_158), int32(1))
	mBase = m.M
	v8441 = m.ExcPending
	if v8441 != 0 {
		goto L1
	} else {
		goto L1782
	}
L1782:
	;
	goto L1780
L1783:
	;
	goto L1427
L1784:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8488 = m.ExcPending
	if v8488 != 0 {
		goto L1
	} else {
		goto L1785
	}
L1785:
	;
	v8501 = v6912
	goto L128
L1786:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+16)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_133), v6912+int32(16))
	mBase = m.M
	v8543 = m.ExcPending
	if v8543 != 0 {
		goto L1
	} else {
		goto L1787
	}
L1787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912))) = v7659
	F_errdetail(m, int32(_a_F_InitPostgres_159), v6912)
	mBase = m.M
	v8547 = m.ExcPending
	if v8547 != 0 {
		goto L1
	} else {
		goto L1788
	}
L1788:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1172), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v8552 = m.ExcPending
	if v8552 != 0 {
		goto L1
	} else {
		goto L1789
	}
L1789:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1790:
	;
	v8558 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+48)) = v8558
	F_errmsg_internal(m, int32(_a_F_InitPostgres_147), v6912+int32(48))
	mBase = m.M
	v8564 = m.ExcPending
	if v8564 != 0 {
		goto L1
	} else {
		goto L1791
	}
L1791:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(335), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8569 = m.ExcPending
	if v8569 != 0 {
		goto L1
	} else {
		goto L1792
	}
L1792:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1793:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8576 = m.ExcPending
	if v8576 != 0 {
		goto L1
	} else {
		goto L1794
	}
L1794:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+224)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_160), v6912+int32(224))
	mBase = m.M
	v8584 = m.ExcPending
	if v8584 != 0 {
		goto L1
	} else {
		goto L1795
	}
L1795:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+212)) = v7709
	v8587 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+208)) = v8587
	F_errdetail(m, int32(_a_F_InitPostgres_161), v6912+int32(208))
	mBase = m.M
	v8593 = m.ExcPending
	if v8593 != 0 {
		goto L1
	} else {
		goto L1796
	}
L1796:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(345), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8598 = m.ExcPending
	if v8598 != 0 {
		goto L1
	} else {
		goto L1797
	}
L1797:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1798:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v8605 = m.ExcPending
	if v8605 != 0 {
		goto L1
	} else {
		goto L1799
	}
L1799:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+192)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_162), v6912+int32(192))
	mBase = m.M
	v8613 = m.ExcPending
	if v8613 != 0 {
		goto L1
	} else {
		goto L1800
	}
L1800:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(365), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8618 = m.ExcPending
	if v8618 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1801:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1802:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v8625 = m.ExcPending
	if v8625 != 0 {
		goto L1
	} else {
		goto L1803
	}
L1803:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+176)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_163), v6912+int32(176))
	mBase = m.M
	v8633 = m.ExcPending
	if v8633 != 0 {
		goto L1
	} else {
		goto L1804
	}
L1804:
	;
	F_errdetail(m, int32(_a_F_InitPostgres_164), int32(0))
	mBase = m.M
	v8637 = m.ExcPending
	if v8637 != 0 {
		goto L1
	} else {
		goto L1805
	}
L1805:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(378), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8642 = m.ExcPending
	if v8642 != 0 {
		goto L1
	} else {
		goto L1806
	}
L1806:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1807:
	;
	F_errcode(m, int32(_a_F_InitPostgres_125))
	mBase = m.M
	v8649 = m.ExcPending
	if v8649 != 0 {
		goto L1
	} else {
		goto L1808
	}
L1808:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+160)) = v6912 + int32(432)
	F_errmsg(m, int32(_a_F_InitPostgres_165), v6912+int32(160))
	mBase = m.M
	v8657 = m.ExcPending
	if v8657 != 0 {
		goto L1
	} else {
		goto L1809
	}
L1809:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(399), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8662 = m.ExcPending
	if v8662 != 0 {
		goto L1
	} else {
		goto L1810
	}
L1810:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1811:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_166), int32(0))
	mBase = m.M
	v8670 = m.ExcPending
	if v8670 != 0 {
		goto L1
	} else {
		goto L1812
	}
L1812:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+64)) = v7992
	F_errdetail(m, int32(_a_F_InitPostgres_167), v6912-int32(-64))
	mBase = m.M
	v8676 = m.ExcPending
	if v8676 != 0 {
		goto L1
	} else {
		goto L1813
	}
L1813:
	;
	F_errhint(m, int32(_a_F_InitPostgres_168), int32(0))
	mBase = m.M
	v8680 = m.ExcPending
	if v8680 != 0 {
		goto L1
	} else {
		goto L1814
	}
L1814:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(425), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8685 = m.ExcPending
	if v8685 != 0 {
		goto L1
	} else {
		goto L1815
	}
L1815:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1816:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_166), int32(0))
	mBase = m.M
	v8693 = m.ExcPending
	if v8693 != 0 {
		goto L1
	} else {
		goto L1817
	}
L1817:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+80)) = v7998
	F_errdetail(m, int32(_a_F_InitPostgres_169), v6912+int32(80))
	mBase = m.M
	v8699 = m.ExcPending
	if v8699 != 0 {
		goto L1
	} else {
		goto L1818
	}
L1818:
	;
	F_errhint(m, int32(_a_F_InitPostgres_168), int32(0))
	mBase = m.M
	v8703 = m.ExcPending
	if v8703 != 0 {
		goto L1
	} else {
		goto L1819
	}
L1819:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(432), int32(_a_F_InitPostgres_148))
	mBase = m.M
	v8708 = m.ExcPending
	if v8708 != 0 {
		goto L1
	} else {
		goto L1820
	}
L1820:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1821:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8716 = m.ExcPending
	if v8716 != 0 {
		goto L1
	} else {
		goto L1822
	}
L1822:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6912)+256)) = v6900
	F_errmsg(m, int32(_a_F_InitPostgres_133), v6912+int32(256))
	mBase = m.M
	v8722 = m.ExcPending
	if v8722 != 0 {
		goto L1
	} else {
		goto L1823
	}
L1823:
	;
	F_errdetail(m, int32(_a_F_InitPostgres_170), int32(0))
	mBase = m.M
	v8726 = m.ExcPending
	if v8726 != 0 {
		goto L1
	} else {
		goto L1824
	}
L1824:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1097), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v8731 = m.ExcPending
	if v8731 != 0 {
		goto L1
	} else {
		goto L1825
	}
L1825:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
