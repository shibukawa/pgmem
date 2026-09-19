package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1105 int64
	_ = v1105
	var v1106 int64
	_ = v1106
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1169 int32
	_ = v1169
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1201 int32
	_ = v1201
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1262 int32
	_ = v1262
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1309 int32
	_ = v1309
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1493 int32
	_ = v1493
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1554 int32
	_ = v1554
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1569 int64
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1630 int32
	_ = v1630
	var v1638 int32
	_ = v1638
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1653 int32
	_ = v1653
	var v1663 int32
	_ = v1663
	var v1676 int32
	_ = v1676
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1746 int32
	_ = v1746
	var v1759 int32
	_ = v1759
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1800 int32
	_ = v1800
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1858 int64
	_ = v1858
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1902 int32
	_ = v1902
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1960 int32
	_ = v1960
	var v2006 int32
	_ = v2006
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2021 int32
	_ = v2021
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2094 int32
	_ = v2094
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2155 int32
	_ = v2155
	var v2188 int32
	_ = v2188
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2297 int32
	_ = v2297
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2332 int32
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2454 int32
	_ = v2454
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2589 int32
	_ = v2589
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2661 int32
	_ = v2661
	var v2671 int32
	_ = v2671
	var v2676 int32
	_ = v2676
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2695 int32
	_ = v2695
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2725 int32
	_ = v2725
	var v2731 int32
	_ = v2731
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2753 int32
	_ = v2753
	var v2757 int32
	_ = v2757
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2775 int32
	_ = v2775
	var v2778 int32
	_ = v2778
	var v2781 int32
	_ = v2781
	var v2783 int32
	_ = v2783
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2793 int32
	_ = v2793
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2814 int32
	_ = v2814
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2842 int32
	_ = v2842
	var v2848 int32
	_ = v2848
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2870 int32
	_ = v2870
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2882 int32
	_ = v2882
	var v2892 int32
	_ = v2892
	var v2895 int32
	_ = v2895
	var v2898 int32
	_ = v2898
	var v2900 int32
	_ = v2900
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2914 int32
	_ = v2914
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2922 int32
	_ = v2922
	var v2931 int32
	_ = v2931
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3006 int32
	_ = v3006
	var v3012 int32
	_ = v3012
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3046 int32
	_ = v3046
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3057 int32
	_ = v3057
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3066 int64
	_ = v3066
	var v3072 int32
	_ = v3072
	var v3074 int32
	_ = v3074
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3091 int32
	_ = v3091
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3097 int64
	_ = v3097
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3126 int32
	_ = v3126
	var v3131 int32
	_ = v3131
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3152 int32
	_ = v3152
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3264 int32
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3274 int32
	_ = v3274
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3308 int32
	_ = v3308
	var v3319 int32
	_ = v3319
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3328 int32
	_ = v3328
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3334 int32
	_ = v3334
	var v3340 int32
	_ = v3340
	var v3350 int32
	_ = v3350
	var v3384 int32
	_ = v3384
	var v3398 int32
	_ = v3398
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3484 int32
	_ = v3484
	var v3490 int32
	_ = v3490
	var v3502 int32
	_ = v3502
	var v3504 int32
	_ = v3504
	var v3509 int32
	_ = v3509
	var v3518 int32
	_ = v3518
	var v3554 int32
	_ = v3554
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3566 int32
	_ = v3566
	var v3570 int32
	_ = v3570
	var v3581 int32
	_ = v3581
	var v3614 int32
	_ = v3614
	var v3631 int32
	_ = v3631
	var v3666 int32
	_ = v3666
	var v3674 int32
	_ = v3674
	var v3685 int32
	_ = v3685
	var v3686 int32
	_ = v3686
	var v3687 int32
	_ = v3687
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3740 int32
	_ = v3740
	var v3772 int32
	_ = v3772
	var v3785 int32
	_ = v3785
	var v3819 int32
	_ = v3819
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3833 int32
	_ = v3833
	var v3850 int32
	_ = v3850
	var v3876 int32
	_ = v3876
	var v3879 int32
	_ = v3879
	var v3895 int32
	_ = v3895
	var v3921 int32
	_ = v3921
	var v3937 int32
	_ = v3937
	var v3963 int32
	_ = v3963
	var v3973 int32
	_ = v3973
	var v3979 int32
	_ = v3979
	var v4005 int32
	_ = v4005
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4057 int32
	_ = v4057
	var v4058 int32
	_ = v4058
	var v4059 int32
	_ = v4059
	var v4060 int32
	_ = v4060
	var v4061 int32
	_ = v4061
	var v4062 int32
	_ = v4062
	var v4073 int32
	_ = v4073
	var v4077 int32
	_ = v4077
	var v4079 int32
	_ = v4079
	var v4085 int32
	_ = v4085
	var v4091 int32
	_ = v4091
	var v4092 int32
	_ = v4092
	var v4097 int32
	_ = v4097
	var v4101 int32
	_ = v4101
	var v4106 int32
	_ = v4106
	var v4111 int32
	_ = v4111
	var v4113 int32
	_ = v4113
	var v4118 int32
	_ = v4118
	var v4127 int32
	_ = v4127
	var v4128 int32
	_ = v4128
	var v4134 int32
	_ = v4134
	var v4139 int32
	_ = v4139
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4154 int32
	_ = v4154
	var v4155 int32
	_ = v4155
	var v4158 int32
	_ = v4158
	var v4160 int32
	_ = v4160
	var v4162 int32
	_ = v4162
	var v4166 int32
	_ = v4166
	var v4169 int32
	_ = v4169
	var v4172 int32
	_ = v4172
	var v4175 int32
	_ = v4175
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4187 int64
	_ = v4187
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4203 int32
	_ = v4203
	var v4209 int32
	_ = v4209
	var v4213 int32
	_ = v4213
	var v4214 int32
	_ = v4214
	var v4221 int32
	_ = v4221
	var v4222 int32
	_ = v4222
	var v4223 int32
	_ = v4223
	var v4227 int32
	_ = v4227
	var v4230 int32
	_ = v4230
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4243 int32
	_ = v4243
	var v4249 int32
	_ = v4249
	var v4251 int32
	_ = v4251
	var v4253 int32
	_ = v4253
	var v4263 int32
	_ = v4263
	var v4271 int32
	_ = v4271
	var v4272 int32
	_ = v4272
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4278 int32
	_ = v4278
	var v4281 int32
	_ = v4281
	var v4284 int32
	_ = v4284
	var v4287 int32
	_ = v4287
	var v4288 int32
	_ = v4288
	var v4291 int32
	_ = v4291
	var v4292 int32
	_ = v4292
	var v4295 int32
	_ = v4295
	var v4302 int32
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4324 int32
	_ = v4324
	var v4325 int32
	_ = v4325
	var v4330 int32
	_ = v4330
	var v4331 int32
	_ = v4331
	var v4333 int32
	_ = v4333
	var v4337 int32
	_ = v4337
	var v4342 int32
	_ = v4342
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4351 int32
	_ = v4351
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4357 int32
	_ = v4357
	var v4360 int32
	_ = v4360
	var v4362 int32
	_ = v4362
	var v4364 int32
	_ = v4364
	var v4367 int32
	_ = v4367
	var v4370 int32
	_ = v4370
	var v4371 int32
	_ = v4371
	var v4372 int32
	_ = v4372
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4383 int32
	_ = v4383
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4387 int32
	_ = v4387
	var v4389 int32
	_ = v4389
	var v4390 int32
	_ = v4390
	var v4391 int32
	_ = v4391
	var v4395 int32
	_ = v4395
	var v4398 int32
	_ = v4398
	var v4399 int32
	_ = v4399
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4407 int32
	_ = v4407
	var v4408 int32
	_ = v4408
	var v4410 int32
	_ = v4410
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4415 int32
	_ = v4415
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
	var v4435 int32
	_ = v4435
	var v4440 int32
	_ = v4440
	var v4441 int32
	_ = v4441
	var v4447 int32
	_ = v4447
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4456 int32
	_ = v4456
	var v4458 int32
	_ = v4458
	var v4461 int32
	_ = v4461
	var v4464 int32
	_ = v4464
	var v4465 int32
	_ = v4465
	var v4466 int32
	_ = v4466
	var v4474 int32
	_ = v4474
	var v4476 int32
	_ = v4476
	var v4477 int32
	_ = v4477
	var v4478 int32
	_ = v4478
	var v4479 int32
	_ = v4479
	var v4481 int32
	_ = v4481
	var v4483 int32
	_ = v4483
	var v4484 int32
	_ = v4484
	var v4485 int32
	_ = v4485
	var v4489 int32
	_ = v4489
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4502 int32
	_ = v4502
	var v4507 int32
	_ = v4507
	var v4509 int32
	_ = v4509
	var v4510 int32
	_ = v4510
	var v4511 int32
	_ = v4511
	var v4512 int32
	_ = v4512
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4515 int32
	_ = v4515
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4520 int32
	_ = v4520
	var v4521 int32
	_ = v4521
	var v4524 int32
	_ = v4524
	var v4535 int32
	_ = v4535
	var v4536 int32
	_ = v4536
	var v4537 int32
	_ = v4537
	var v4538 int32
	_ = v4538
	var v4539 int32
	_ = v4539
	var v4540 int32
	_ = v4540
	var v4541 int32
	_ = v4541
	var v4546 int32
	_ = v4546
	var v4547 int32
	_ = v4547
	var v4548 int32
	_ = v4548
	var v4549 int32
	_ = v4549
	var v4553 int32
	_ = v4553
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4560 int32
	_ = v4560
	var v4562 int32
	_ = v4562
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4569 int32
	_ = v4569
	var v4570 int32
	_ = v4570
	var v4576 int32
	_ = v4576
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4583 int32
	_ = v4583
	var v4584 int32
	_ = v4584
	var v4586 int32
	_ = v4586
	var v4587 int32
	_ = v4587
	var v4588 int32
	_ = v4588
	var v4589 int64
	_ = v4589
	var v4600 int32
	_ = v4600
	var v4604 int32
	_ = v4604
	var v4609 int32
	_ = v4609
	var v4610 int32
	_ = v4610
	var v4611 int32
	_ = v4611
	var v4612 int32
	_ = v4612
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4620 int32
	_ = v4620
	var v4621 int32
	_ = v4621
	var v4622 int32
	_ = v4622
	var v4623 int32
	_ = v4623
	var v4625 int32
	_ = v4625
	var v4628 int32
	_ = v4628
	var v4629 int32
	_ = v4629
	var v4630 int32
	_ = v4630
	var v4632 int32
	_ = v4632
	var v4634 int32
	_ = v4634
	var v4635 int32
	_ = v4635
	var v4639 int32
	_ = v4639
	var v4642 int32
	_ = v4642
	var v4648 int32
	_ = v4648
	var v4653 int32
	_ = v4653
	var v4656 int32
	_ = v4656
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4663 int32
	_ = v4663
	var v4665 int32
	_ = v4665
	var v4669 int32
	_ = v4669
	var v4670 int32
	_ = v4670
	var v4671 int32
	_ = v4671
	var v4675 int32
	_ = v4675
	var v4679 int32
	_ = v4679
	var v4681 int32
	_ = v4681
	var v4683 int32
	_ = v4683
	var v4685 int32
	_ = v4685
	var v4687 int32
	_ = v4687
	var v4697 int32
	_ = v4697
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4706 int32
	_ = v4706
	var v4712 int32
	_ = v4712
	var v4713 int32
	_ = v4713
	var v4714 int32
	_ = v4714
	var v4722 int32
	_ = v4722
	var v4724 int32
	_ = v4724
	var v4726 int32
	_ = v4726
	var v4727 int32
	_ = v4727
	var v4731 int32
	_ = v4731
	var v4733 int32
	_ = v4733
	var v4739 int32
	_ = v4739
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4751 int32
	_ = v4751
	var v4755 int32
	_ = v4755
	var v4760 int32
	_ = v4760
	var v4765 int32
	_ = v4765
	var v4767 int32
	_ = v4767
	var v4772 int32
	_ = v4772
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4786 int32
	_ = v4786
	var v4791 int32
	_ = v4791
	var v4792 int32
	_ = v4792
	var v4793 int32
	_ = v4793
	var v4799 int32
	_ = v4799
	var v4800 int32
	_ = v4800
	var v4801 int32
	_ = v4801
	var v4809 int32
	_ = v4809
	var v4811 int32
	_ = v4811
	var v4816 int32
	_ = v4816
	var v4817 int32
	_ = v4817
	var v4826 int32
	_ = v4826
	var v4831 int32
	_ = v4831
	var v4834 int32
	_ = v4834
	var v4837 int32
	_ = v4837
	var v4839 int32
	_ = v4839
	var v4841 int32
	_ = v4841
	var v4845 int32
	_ = v4845
	var v4846 int32
	_ = v4846
	var v4847 int32
	_ = v4847
	var v4853 int32
	_ = v4853
	var v4854 int32
	_ = v4854
	var v4864 int32
	_ = v4864
	var v4869 int32
	_ = v4869
	var v4872 int32
	_ = v4872
	var v4874 int32
	_ = v4874
	var v4876 int32
	_ = v4876
	var v4881 int32
	_ = v4881
	var v4884 int32
	_ = v4884
	var v4888 int32
	_ = v4888
	var v4889 int32
	_ = v4889
	var v4895 int32
	_ = v4895
	var v4896 int32
	_ = v4896
	var v4906 int32
	_ = v4906
	var v4911 int32
	_ = v4911
	var v4914 int32
	_ = v4914
	var v4916 int32
	_ = v4916
	var v4918 int32
	_ = v4918
	var v4923 int32
	_ = v4923
	var v4926 int32
	_ = v4926
	var v4930 int32
	_ = v4930
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4935 int32
	_ = v4935
	var v4936 int32
	_ = v4936
	var v4937 int32
	_ = v4937
	var v4940 int32
	_ = v4940
	var v4942 int32
	_ = v4942
	var v4953 int32
	_ = v4953
	var v4954 int32
	_ = v4954
	var v4956 int32
	_ = v4956
	var v4989 int32
	_ = v4989
	var v4990 int32
	_ = v4990
	var v4991 int64
	_ = v4991
	var v4993 int64
	_ = v4993
	var v4995 int32
	_ = v4995
	var v5000 int32
	_ = v5000
	var v5003 int32
	_ = v5003
	var v5004 int32
	_ = v5004
	var v5007 int32
	_ = v5007
	var v5017 int32
	_ = v5017
	var v5052 int32
	_ = v5052
	var v5053 int32
	_ = v5053
	var v5055 int32
	_ = v5055
	var v5056 int32
	_ = v5056
	var v5060 int32
	_ = v5060
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5067 int32
	_ = v5067
	var v5068 int32
	_ = v5068
	var v5072 int32
	_ = v5072
	var v5074 int32
	_ = v5074
	var v5077 int32
	_ = v5077
	var v5121 int32
	_ = v5121
	var v5122 int32
	_ = v5122
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5139 int32
	_ = v5139
	var v5144 int32
	_ = v5144
	var v5147 int32
	_ = v5147
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5155 int32
	_ = v5155
	var v5160 int32
	_ = v5160
	var v5162 int32
	_ = v5162
	var v5163 int32
	_ = v5163
	var v5164 int32
	_ = v5164
	var v5170 int32
	_ = v5170
	var v5171 int32
	_ = v5171
	var v5172 int32
	_ = v5172
	var v5182 int32
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5184 int32
	_ = v5184
	var v5193 int32
	_ = v5193
	var v5196 int32
	_ = v5196
	var v5200 int32
	_ = v5200
	var v5201 int32
	_ = v5201
	var v5205 int32
	_ = v5205
	var v5207 int32
	_ = v5207
	var v5208 int32
	_ = v5208
	var v5211 int32
	_ = v5211
	var v5216 int32
	_ = v5216
	var v5217 int32
	_ = v5217
	var v5221 int32
	_ = v5221
	var v5226 int32
	_ = v5226
	var v5227 int32
	_ = v5227
	var v5228 int32
	_ = v5228
	var v5234 int32
	_ = v5234
	var v5235 int32
	_ = v5235
	var v5236 int32
	_ = v5236
	var v5244 int64
	_ = v5244
	var v5252 int32
	_ = v5252
	var v5253 int32
	_ = v5253
	var v5256 int64
	_ = v5256
	var v5259 int64
	_ = v5259
	var v5267 int32
	_ = v5267
	var v5271 int32
	_ = v5271
	var v5274 int32
	_ = v5274
	var v5275 int32
	_ = v5275
	var v5279 int32
	_ = v5279
	var v5284 int32
	_ = v5284
	var v5285 int32
	_ = v5285
	var v5286 int32
	_ = v5286
	var v5287 int32
	_ = v5287
	var v5293 int32
	_ = v5293
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5305 int32
	_ = v5305
	var v5306 int32
	_ = v5306
	var v5307 int32
	_ = v5307
	var v5308 int32
	_ = v5308
	var v5313 int32
	_ = v5313
	var v5314 int32
	_ = v5314
	var v5318 int32
	_ = v5318
	var v5323 int32
	_ = v5323
	var v5324 int32
	_ = v5324
	var v5325 int32
	_ = v5325
	var v5326 int32
	_ = v5326
	var v5332 int32
	_ = v5332
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5342 int32
	_ = v5342
	var v5343 int32
	_ = v5343
	var v5349 int32
	_ = v5349
	var v5350 int32
	_ = v5350
	var v5351 int32
	_ = v5351
	var v5362 int64
	_ = v5362
	var v5369 int64
	_ = v5369
	var v5370 int64
	_ = v5370
	var v5373 int64
	_ = v5373
	var v5374 int64
	_ = v5374
	var v5378 int64
	_ = v5378
	var v5381 int32
	_ = v5381
	var v5382 int32
	_ = v5382
	var v5391 int32
	_ = v5391
	var v5392 int32
	_ = v5392
	var v5438 int64
	_ = v5438
	var v5442 int64
	_ = v5442
	var v5443 int64
	_ = v5443
	var v5447 int64
	_ = v5447
	var v5450 int32
	_ = v5450
	var v5454 int32
	_ = v5454
	var v5458 int32
	_ = v5458
	var v5459 int32
	_ = v5459
	var v5460 int64
	_ = v5460
	var v5463 int32
	_ = v5463
	var v5467 int32
	_ = v5467
	var v5471 int32
	_ = v5471
	var v5472 int32
	_ = v5472
	var v5474 int32
	_ = v5474
	var v5476 int32
	_ = v5476
	var v5479 int32
	_ = v5479
	var v5480 int32
	_ = v5480
	var v5482 int32
	_ = v5482
	var v5488 int64
	_ = v5488
	var v5489 int32
	_ = v5489
	var v5497 int32
	_ = v5497
	var v5500 int32
	_ = v5500
	var v5501 int32
	_ = v5501
	var v5506 int32
	_ = v5506
	var v5521 int32
	_ = v5521
	var v5530 int32
	_ = v5530
	var v5562 int32
	_ = v5562
	var v5570 int32
	_ = v5570
	var v5571 int32
	_ = v5571
	var v5573 int32
	_ = v5573
	var v5578 int32
	_ = v5578
	var v5579 int32
	_ = v5579
	var v5583 int32
	_ = v5583
	var v5585 int32
	_ = v5585
	var v5603 int32
	_ = v5603
	var v5628 int32
	_ = v5628
	var v5639 int32
	_ = v5639
	var v5653 int32
	_ = v5653
	var v5678 int32
	_ = v5678
	var v5720 int32
	_ = v5720
	var v5735 int32
	_ = v5735
	var v5772 int32
	_ = v5772
	var v5776 int32
	_ = v5776
	var v5778 int32
	_ = v5778
	var v5787 int32
	_ = v5787
	var v5789 int32
	_ = v5789
	var v5823 int32
	_ = v5823
	var v5824 int32
	_ = v5824
	var v5825 int32
	_ = v5825
	var v5826 int32
	_ = v5826
	var v5837 int32
	_ = v5837
	var v5838 int32
	_ = v5838
	var v5839 int32
	_ = v5839
	var v5846 int32
	_ = v5846
	var v5847 int32
	_ = v5847
	var v5851 int32
	_ = v5851
	var v5864 int32
	_ = v5864
	var v5938 int32
	_ = v5938
	var v5946 int32
	_ = v5946
	var v5987 int32
	_ = v5987
	var v5991 int32
	_ = v5991
	var v5996 int32
	_ = v5996
	var v5997 int32
	_ = v5997
	var v6003 int32
	_ = v6003
	var v6009 int32
	_ = v6009
	var v6010 int32
	_ = v6010
	var v6018 int32
	_ = v6018
	var v6025 int32
	_ = v6025
	var v6030 int32
	_ = v6030
	var v6035 int32
	_ = v6035
	var v6036 int32
	_ = v6036
	var v6042 int32
	_ = v6042
	var v6044 int32
	_ = v6044
	var v6048 int32
	_ = v6048
	var v6049 int32
	_ = v6049
	var v6053 int32
	_ = v6053
	var v6054 int32
	_ = v6054
	var v6066 int32
	_ = v6066
	var v6071 int32
	_ = v6071
	var v6076 int32
	_ = v6076
	var v6077 int32
	_ = v6077
	var v6086 int32
	_ = v6086
	var v6091 int32
	_ = v6091
	var v6092 int32
	_ = v6092
	var v6093 int32
	_ = v6093
	var v6103 int32
	_ = v6103
	var v6104 int32
	_ = v6104
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6122 int32
	_ = v6122
	var v6127 int32
	_ = v6127
	var v6128 int32
	_ = v6128
	var v6129 int32
	_ = v6129
	var v6133 int32
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6138 int32
	_ = v6138
	var v6140 int32
	_ = v6140
	var v6146 int32
	_ = v6146
	var v6151 int32
	_ = v6151
	var v6152 int32
	_ = v6152
	var v6154 int32
	_ = v6154
	var v6155 int32
	_ = v6155
	var v6156 int32
	_ = v6156
	var v6158 int64
	_ = v6158
	var v6160 int64
	_ = v6160
	var v6165 int32
	_ = v6165
	var v6172 int32
	_ = v6172
	var v6175 int32
	_ = v6175
	var v6181 int32
	_ = v6181
	var v6182 int32
	_ = v6182
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6189 int32
	_ = v6189
	var v6195 int32
	_ = v6195
	var v6200 int32
	_ = v6200
	var v6202 int32
	_ = v6202
	var v6204 int32
	_ = v6204
	var v6205 int64
	_ = v6205
	var v6206 int64
	_ = v6206
	var v6208 int64
	_ = v6208
	var v6209 int64
	_ = v6209
	var v6216 int32
	_ = v6216
	var v6217 int32
	_ = v6217
	var v6225 int32
	_ = v6225
	var v6230 int32
	_ = v6230
	var v6231 int32
	_ = v6231
	var v6236 int32
	_ = v6236
	var v6237 int32
	_ = v6237
	var v6241 int32
	_ = v6241
	var v6248 int32
	_ = v6248
	var v6253 int32
	_ = v6253
	var v6262 int64
	_ = v6262
	var v6264 int64
	_ = v6264
	var v6268 int64
	_ = v6268
	var v6271 int32
	_ = v6271
	var v6272 int32
	_ = v6272
	var v6274 int32
	_ = v6274
	var v6276 int32
	_ = v6276
	var v6278 int32
	_ = v6278
	var v6280 int32
	_ = v6280
	var v6324 int32
	_ = v6324
	var v6325 int32
	_ = v6325
	var v6333 int32
	_ = v6333
	var v6376 int32
	_ = v6376
	var v6379 int32
	_ = v6379
	var v6380 int32
	_ = v6380
	var v6381 int32
	_ = v6381
	var v6382 int32
	_ = v6382
	var v6383 int32
	_ = v6383
	var v6384 int32
	_ = v6384
	var v6385 int32
	_ = v6385
	var v6386 int32
	_ = v6386
	var v6391 int32
	_ = v6391
	var v6392 int32
	_ = v6392
	var v6393 int32
	_ = v6393
	var v6394 int32
	_ = v6394
	var v6398 int32
	_ = v6398
	var v6400 int32
	_ = v6400
	var v6403 int32
	_ = v6403
	var v6404 int32
	_ = v6404
	var v6405 int32
	_ = v6405
	var v6407 int32
	_ = v6407
	var v6411 int32
	_ = v6411
	var v6412 int32
	_ = v6412
	var v6413 int32
	_ = v6413
	var v6414 int32
	_ = v6414
	var v6415 int32
	_ = v6415
	var v6421 int32
	_ = v6421
	var v6422 int32
	_ = v6422
	var v6423 int32
	_ = v6423
	var v6424 int32
	_ = v6424
	var v6425 int32
	_ = v6425
	var v6426 int32
	_ = v6426
	var v6427 int32
	_ = v6427
	var v6428 int32
	_ = v6428
	var v6433 int32
	_ = v6433
	var v6434 int32
	_ = v6434
	var v6435 int32
	_ = v6435
	var v6436 int32
	_ = v6436
	var v6440 int32
	_ = v6440
	var v6442 int32
	_ = v6442
	var v6445 int32
	_ = v6445
	var v6446 int32
	_ = v6446
	var v6447 int32
	_ = v6447
	var v6449 int32
	_ = v6449
	var v6453 int32
	_ = v6453
	var v6454 int32
	_ = v6454
	var v6455 int32
	_ = v6455
	var v6456 int32
	_ = v6456
	var v6457 int32
	_ = v6457
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
	var v6500 int32
	_ = v6500
	var v6503 int32
	_ = v6503
	var v6507 int32
	_ = v6507
	var v6509 int32
	_ = v6509
	var v6514 int32
	_ = v6514
	var v6517 int32
	_ = v6517
	var v6519 int32
	_ = v6519
	var v6520 int32
	_ = v6520
	var v6522 int32
	_ = v6522
	var v6523 int32
	_ = v6523
	var v6524 int32
	_ = v6524
	var v6525 int32
	_ = v6525
	var v6526 int32
	_ = v6526
	var v6527 int32
	_ = v6527
	var v6528 int32
	_ = v6528
	var v6533 int32
	_ = v6533
	var v6534 int32
	_ = v6534
	var v6535 int32
	_ = v6535
	var v6540 int32
	_ = v6540
	var v6546 int32
	_ = v6546
	var v6555 int32
	_ = v6555
	var v6564 int32
	_ = v6564
	var v6575 int32
	_ = v6575
	var v6580 int32
	_ = v6580
	var v6584 int32
	_ = v6584
	var v6587 int32
	_ = v6587
	var v6591 int32
	_ = v6591
	var v6596 int32
	_ = v6596
	var v6598 int32
	_ = v6598
	var v6600 int32
	_ = v6600
	var v6601 int32
	_ = v6601
	var v6603 int32
	_ = v6603
	var v6604 int32
	_ = v6604
	var v6605 int32
	_ = v6605
	var v6606 int32
	_ = v6606
	var v6607 int32
	_ = v6607
	var v6608 int32
	_ = v6608
	var v6609 int32
	_ = v6609
	var v6614 int32
	_ = v6614
	var v6615 int32
	_ = v6615
	var v6616 int32
	_ = v6616
	var v6627 int32
	_ = v6627
	var v6636 int32
	_ = v6636
	var v6645 int32
	_ = v6645
	var v6652 int32
	_ = v6652
	var v6655 int32
	_ = v6655
	var v6656 int32
	_ = v6656
	var v6659 int32
	_ = v6659
	var v6660 int32
	_ = v6660
	var v6661 int32
	_ = v6661
	var v6664 int32
	_ = v6664
	var v6665 int32
	_ = v6665
	var v6666 int64
	_ = v6666
	var v6674 int32
	_ = v6674
	var v6679 int32
	_ = v6679
	var v6684 int32
	_ = v6684
	var v6686 int32
	_ = v6686
	var v6690 int32
	_ = v6690
	var v6692 int32
	_ = v6692
	var v6694 int32
	_ = v6694
	var v6697 int32
	_ = v6697
	var v6700 int32
	_ = v6700
	var v6701 int32
	_ = v6701
	var v6702 int32
	_ = v6702
	var v6710 int32
	_ = v6710
	var v6712 int32
	_ = v6712
	var v6714 int32
	_ = v6714
	var v6720 int32
	_ = v6720
	var v6721 int32
	_ = v6721
	var v6722 int32
	_ = v6722
	var v6728 int32
	_ = v6728
	var v6729 int32
	_ = v6729
	var v6730 int32
	_ = v6730
	var v6731 int32
	_ = v6731
	var v6733 int32
	_ = v6733
	var v6734 int64
	_ = v6734
	var v6735 int32
	_ = v6735
	var v6741 int32
	_ = v6741
	var v6742 int32
	_ = v6742
	var v6748 int32
	_ = v6748
	var v6749 int32
	_ = v6749
	var v6750 int32
	_ = v6750
	var v6754 int32
	_ = v6754
	var v6756 int32
	_ = v6756
	var v6757 int32
	_ = v6757
	var v6762 int32
	_ = v6762
	var v6766 int32
	_ = v6766
	var v6771 int32
	_ = v6771
	var v6774 int32
	_ = v6774
	var v6777 int32
	_ = v6777
	var v6782 int32
	_ = v6782
	var v6783 int32
	_ = v6783
	var v6784 int32
	_ = v6784
	var v6787 int64
	_ = v6787
	var v6788 int64
	_ = v6788
	var v6799 int32
	_ = v6799
	var v6803 int32
	_ = v6803
	var v6805 int32
	_ = v6805
	var v6806 int32
	_ = v6806
	var v6811 int32
	_ = v6811
	var v6812 int32
	_ = v6812
	var v6816 int32
	_ = v6816
	var v6818 int32
	_ = v6818
	var v6821 int32
	_ = v6821
	var v6827 int32
	_ = v6827
	var v6828 int32
	_ = v6828
	var v6836 int32
	_ = v6836
	var v6839 int32
	_ = v6839
	var v6840 int32
	_ = v6840
	var v6841 int32
	_ = v6841
	var v6847 int32
	_ = v6847
	var v6852 int32
	_ = v6852
	var v6853 int32
	_ = v6853
	var v6855 int32
	_ = v6855
	var v6859 int32
	_ = v6859
	var v6863 int32
	_ = v6863
	var v6865 int32
	_ = v6865
	var v6869 int32
	_ = v6869
	var v6872 int32
	_ = v6872
	var v6874 int32
	_ = v6874
	var v6875 int32
	_ = v6875
	var v6876 int32
	_ = v6876
	var v6879 int32
	_ = v6879
	var v6880 int32
	_ = v6880
	var v6887 int32
	_ = v6887
	var v6899 int32
	_ = v6899
	var v6908 int32
	_ = v6908
	var v6916 int32
	_ = v6916
	var v6917 int32
	_ = v6917
	var v6918 int32
	_ = v6918
	var v6919 int32
	_ = v6919
	var v6921 int32
	_ = v6921
	var v6922 int32
	_ = v6922
	var v6923 int32
	_ = v6923
	var v6930 int32
	_ = v6930
	var v6942 int32
	_ = v6942
	var v6951 int32
	_ = v6951
	var v6960 int32
	_ = v6960
	var v6961 int32
	_ = v6961
	var v6963 int32
	_ = v6963
	var v6964 int32
	_ = v6964
	var v6968 int32
	_ = v6968
	var v6969 int32
	_ = v6969
	var v6973 int32
	_ = v6973
	var v6983 int32
	_ = v6983
	var v6989 int32
	_ = v6989
	var v6998 int32
	_ = v6998
	var v6999 int32
	_ = v6999
	var v7005 int32
	_ = v7005
	var v7007 int32
	_ = v7007
	var v7008 int32
	_ = v7008
	var v7012 int32
	_ = v7012
	var v7015 int32
	_ = v7015
	var v7017 int32
	_ = v7017
	var v7022 int32
	_ = v7022
	var v7024 int32
	_ = v7024
	var v7025 int32
	_ = v7025
	var v7028 int32
	_ = v7028
	var v7029 int32
	_ = v7029
	var v7033 int32
	_ = v7033
	var v7043 int32
	_ = v7043
	var v7068 int32
	_ = v7068
	var v7077 int32
	_ = v7077
	var v7080 int32
	_ = v7080
	var v7124 int32
	_ = v7124
	var v7125 int32
	_ = v7125
	var v7128 int32
	_ = v7128
	var v7130 int32
	_ = v7130
	var v7132 int32
	_ = v7132
	var v7135 int32
	_ = v7135
	var v7137 int32
	_ = v7137
	var v7138 int32
	_ = v7138
	var v7183 int32
	_ = v7183
	var v7187 int32
	_ = v7187
	var v7188 int32
	_ = v7188
	var v7189 int32
	_ = v7189
	var v7193 int32
	_ = v7193
	var v7197 int32
	_ = v7197
	var v7201 int32
	_ = v7201
	var v7203 int32
	_ = v7203
	var v7205 int32
	_ = v7205
	var v7212 int32
	_ = v7212
	var v7216 int32
	_ = v7216
	var v7221 int32
	_ = v7221
	var v7225 int32
	_ = v7225
	var v7226 int32
	_ = v7226
	var v7229 int32
	_ = v7229
	var v7232 int32
	_ = v7232
	var v7233 int32
	_ = v7233
	var v7234 int32
	_ = v7234
	var v7235 int32
	_ = v7235
	var v7236 int32
	_ = v7236
	var v7237 int32
	_ = v7237
	var v7238 int32
	_ = v7238
	var v7240 int32
	_ = v7240
	var v7243 int32
	_ = v7243
	var v7246 int32
	_ = v7246
	var v7247 int32
	_ = v7247
	var v7249 int32
	_ = v7249
	var v7252 int32
	_ = v7252
	var v7260 int32
	_ = v7260
	var v7262 int32
	_ = v7262
	var v7267 int32
	_ = v7267
	var v7270 int32
	_ = v7270
	var v7271 int32
	_ = v7271
	var v7274 int32
	_ = v7274
	var v7277 int32
	_ = v7277
	var v7278 int32
	_ = v7278
	var v7279 int32
	_ = v7279
	var v7280 int32
	_ = v7280
	var v7281 int32
	_ = v7281
	var v7282 int32
	_ = v7282
	var v7283 int32
	_ = v7283
	var v7285 int32
	_ = v7285
	var v7288 int32
	_ = v7288
	var v7289 int32
	_ = v7289
	var v7290 int32
	_ = v7290
	var v7291 int32
	_ = v7291
	var v7295 int32
	_ = v7295
	var v7301 int32
	_ = v7301
	var v7302 int32
	_ = v7302
	var v7303 int32
	_ = v7303
	var v7309 int32
	_ = v7309
	var v7310 int32
	_ = v7310
	var v7316 int32
	_ = v7316
	var v7319 int32
	_ = v7319
	var v7325 int32
	_ = v7325
	var v7330 int32
	_ = v7330
	var v7332 int32
	_ = v7332
	var v7334 int32
	_ = v7334
	var v7341 int32
	_ = v7341
	var v7345 int32
	_ = v7345
	var v7357 int32
	_ = v7357
	var v7358 int32
	_ = v7358
	var v7359 int32
	_ = v7359
	var v7361 int32
	_ = v7361
	var v7365 int32
	_ = v7365
	var v7366 int32
	_ = v7366
	var v7368 int32
	_ = v7368
	var v7369 int32
	_ = v7369
	var v7370 int32
	_ = v7370
	var v7372 int32
	_ = v7372
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7380 int32
	_ = v7380
	var v7381 int32
	_ = v7381
	var v7384 int32
	_ = v7384
	var v7391 int32
	_ = v7391
	var v7392 int32
	_ = v7392
	var v7393 int32
	_ = v7393
	var v7396 int32
	_ = v7396
	var v7399 int32
	_ = v7399
	var v7404 int32
	_ = v7404
	var v7405 int32
	_ = v7405
	var v7407 int32
	_ = v7407
	var v7409 int32
	_ = v7409
	var v7413 int32
	_ = v7413
	var v7414 int32
	_ = v7414
	var v7415 int32
	_ = v7415
	var v7420 int32
	_ = v7420
	var v7421 int32
	_ = v7421
	var v7422 int32
	_ = v7422
	var v7425 int32
	_ = v7425
	var v7426 int32
	_ = v7426
	var v7427 int32
	_ = v7427
	var v7429 int32
	_ = v7429
	var v7433 int32
	_ = v7433
	var v7434 int32
	_ = v7434
	var v7436 int32
	_ = v7436
	var v7438 int32
	_ = v7438
	var v7440 int32
	_ = v7440
	var v7441 int32
	_ = v7441
	var v7444 int32
	_ = v7444
	var v7451 int32
	_ = v7451
	var v7454 int32
	_ = v7454
	var v7458 int32
	_ = v7458
	var v7461 int32
	_ = v7461
	var v7470 int32
	_ = v7470
	var v7474 int32
	_ = v7474
	var v7476 int32
	_ = v7476
	var v7477 int32
	_ = v7477
	var v7481 int32
	_ = v7481
	var v7482 int32
	_ = v7482
	var v7484 int32
	_ = v7484
	var v7488 int32
	_ = v7488
	var v7490 int32
	_ = v7490
	var v7492 int32
	_ = v7492
	var v7495 int32
	_ = v7495
	var v7500 int32
	_ = v7500
	var v7501 int32
	_ = v7501
	var v7502 int32
	_ = v7502
	var v7504 int32
	_ = v7504
	var v7505 int32
	_ = v7505
	var v7507 int32
	_ = v7507
	var v7509 int32
	_ = v7509
	var v7512 int32
	_ = v7512
	var v7517 int32
	_ = v7517
	var v7518 int32
	_ = v7518
	var v7519 int32
	_ = v7519
	var v7526 int32
	_ = v7526
	var v7528 int32
	_ = v7528
	var v7529 int32
	_ = v7529
	var v7531 int32
	_ = v7531
	var v7542 int32
	_ = v7542
	var v7544 int32
	_ = v7544
	var v7550 int32
	_ = v7550
	var v7555 int32
	_ = v7555
	var v7559 int32
	_ = v7559
	var v7562 int32
	_ = v7562
	var v7566 int32
	_ = v7566
	var v7571 int32
	_ = v7571
	var v7575 int32
	_ = v7575
	var v7578 int32
	_ = v7578
	var v7585 int32
	_ = v7585
	var v7590 int32
	_ = v7590
	var v7594 int32
	_ = v7594
	var v7597 int32
	_ = v7597
	var v7604 int32
	_ = v7604
	var v7609 int32
	_ = v7609
	var v7613 int32
	_ = v7613
	var v7616 int32
	_ = v7616
	var v7620 int32
	_ = v7620
	var v7627 int32
	_ = v7627
	var v7632 int32
	_ = v7632
	var v7636 int32
	_ = v7636
	var v7639 int32
	_ = v7639
	var v7645 int32
	_ = v7645
	var v7650 int32
	_ = v7650
	var v7658 int32
	_ = v7658
	var v7661 int32
	_ = v7661
	var v7669 int32
	_ = v7669
	var v7673 int32
	_ = v7673
	var v7678 int32
	_ = v7678
	var v7680 int32
	_ = v7680
	var v7688 int32
	_ = v7688
	var v7691 int32
	_ = v7691
	var v7693 int32
	_ = v7693
	var v7695 int32
	_ = v7695
	var v7696 int32
	_ = v7696
	var v7697 int32
	_ = v7697
	var v7699 int32
	_ = v7699
	var v7703 int32
	_ = v7703
	var v7707 int32
	_ = v7707
	var v7711 int32
	_ = v7711
	var v7717 int32
	_ = v7717
	var v7722 int32
	_ = v7722
	var v7724 int32
	_ = v7724
	var v7726 int32
	_ = v7726
	var v7728 int32
	_ = v7728
	var v7730 int32
	_ = v7730
	var v7732 int32
	_ = v7732
	var v7735 int32
	_ = v7735
	var v7736 int32
	_ = v7736
	var v7737 int32
	_ = v7737
	var v7741 int32
	_ = v7741
	var v7742 int32
	_ = v7742
	var v7743 int32
	_ = v7743
	var v7744 int32
	_ = v7744
	var v7746 int32
	_ = v7746
	var v7749 int32
	_ = v7749
	var v7752 int32
	_ = v7752
	var v7755 int32
	_ = v7755
	var v7756 int32
	_ = v7756
	var v7759 int32
	_ = v7759
	var v7760 int32
	_ = v7760
	var v7763 int32
	_ = v7763
	var v7770 int32
	_ = v7770
	var v7771 int32
	_ = v7771
	var v7774 int32
	_ = v7774
	var v7778 int32
	_ = v7778
	var v7781 int32
	_ = v7781
	var v7786 int32
	_ = v7786
	var v7793 int32
	_ = v7793
	var v7795 int32
	_ = v7795
	var v7797 int32
	_ = v7797
	var v7798 int32
	_ = v7798
	var v7800 int32
	_ = v7800
	var v7803 int32
	_ = v7803
	var v7809 int32
	_ = v7809
	var v7810 int32
	_ = v7810
	var v7812 int32
	_ = v7812
	var v7814 int32
	_ = v7814
	var v7818 int32
	_ = v7818
	var v7819 int32
	_ = v7819
	var v7820 int32
	_ = v7820
	var v7826 int32
	_ = v7826
	var v7828 int32
	_ = v7828
	var v7830 int32
	_ = v7830
	var v7872 int32
	_ = v7872
	var v7875 int32
	_ = v7875
	var v7876 int32
	_ = v7876
	var v7879 int32
	_ = v7879
	var v7882 int32
	_ = v7882
	var v7886 int32
	_ = v7886
	var v7888 int32
	_ = v7888
	var v7892 int32
	_ = v7892
	var v7932 int32
	_ = v7932
	var v7936 int32
	_ = v7936
	var v7937 int32
	_ = v7937
	var v7980 int32
	_ = v7980
	var v7981 int32
	_ = v7981
	var v7983 int32
	_ = v7983
	var v7990 int32
	_ = v7990
	var v7994 int32
	_ = v7994
	var v7999 int32
	_ = v7999
	var v8002 int32
	_ = v8002
	var v8012 int32
	_ = v8012
	var v8016 int32
	_ = v8016
	var v8019 int32
	_ = v8019
	var v8020 int32
	_ = v8020
	var v8024 int32
	_ = v8024
	var v8027 int32
	_ = v8027
	var v8028 int32
	_ = v8028
	var v8029 int32
	_ = v8029
	var v8030 int32
	_ = v8030
	var v8033 int32
	_ = v8033
	var v8034 int32
	_ = v8034
	var v8035 int32
	_ = v8035
	var v8036 int32
	_ = v8036
	var v8038 int32
	_ = v8038
	var v8039 int32
	_ = v8039
	var v8043 int32
	_ = v8043
	var v8044 int32
	_ = v8044
	var v8047 int32
	_ = v8047
	var v8050 int32
	_ = v8050
	var v8053 int32
	_ = v8053
	var v8056 int32
	_ = v8056
	var v8059 int32
	_ = v8059
	var v8062 int32
	_ = v8062
	var v8063 int32
	_ = v8063
	var v8066 int32
	_ = v8066
	var v8067 int32
	_ = v8067
	var v8070 int32
	_ = v8070
	var v8077 int32
	_ = v8077
	var v8078 int32
	_ = v8078
	var v8081 int32
	_ = v8081
	var v8083 int32
	_ = v8083
	var v8085 int32
	_ = v8085
	var v8089 int32
	_ = v8089
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
	var v8095 int32
	_ = v8095
	var v8100 int32
	_ = v8100
	var v8101 int32
	_ = v8101
	var v8104 int32
	_ = v8104
	var v8105 int32
	_ = v8105
	var v8106 int32
	_ = v8106
	var v8110 int32
	_ = v8110
	var v8111 int32
	_ = v8111
	var v8119 int32
	_ = v8119
	var v8124 int32
	_ = v8124
	var v8127 int32
	_ = v8127
	var v8128 int32
	_ = v8128
	var v8129 int32
	_ = v8129
	var v8130 int32
	_ = v8130
	var v8131 int32
	_ = v8131
	var v8134 int32
	_ = v8134
	var v8143 int32
	_ = v8143
	var v8145 int32
	_ = v8145
	var v8149 int32
	_ = v8149
	var v8154 int32
	_ = v8154
	var v8159 int32
	_ = v8159
	var v8160 int32
	_ = v8160
	var v8161 int32
	_ = v8161
	var v8162 int32
	_ = v8162
	var v8163 int32
	_ = v8163
	var v8165 int32
	_ = v8165
	var v8170 int32
	_ = v8170
	var v8171 int32
	_ = v8171
	var v8172 int32
	_ = v8172
	var v8173 int32
	_ = v8173
	var v8174 int32
	_ = v8174
	var v8176 int32
	_ = v8176
	var v8177 int32
	_ = v8177
	var v8179 int32
	_ = v8179
	var v8180 int32
	_ = v8180
	var v8185 int32
	_ = v8185
	var v8186 int32
	_ = v8186
	var v8196 int32
	_ = v8196
	var v8200 int32
	_ = v8200
	var v8203 int32
	_ = v8203
	var v8206 int32
	_ = v8206
	var v8207 int32
	_ = v8207
	var v8210 int32
	_ = v8210
	var v8211 int32
	_ = v8211
	var v8214 int32
	_ = v8214
	var v8221 int32
	_ = v8221
	var v8222 int32
	_ = v8222
	var v8228 int32
	_ = v8228
	var v8229 int32
	_ = v8229
	var v8233 int32
	_ = v8233
	var v8239 int32
	_ = v8239
	var v8246 int32
	_ = v8246
	var v8247 int32
	_ = v8247
	var v8248 int32
	_ = v8248
	var v8254 int32
	_ = v8254
	var v8257 int32
	_ = v8257
	var v8260 int32
	_ = v8260
	var v8265 int32
	_ = v8265
	var v8267 int32
	_ = v8267
	var v8269 int32
	_ = v8269
	var v8271 int32
	_ = v8271
	var v8273 int32
	_ = v8273
	var v8316 int32
	_ = v8316
	var v8318 int32
	_ = v8318
	var v8320 int32
	_ = v8320
	var v8322 int32
	_ = v8322
	var v8324 int32
	_ = v8324
	var v8329 int32
	_ = v8329
	var v8330 int32
	_ = v8330
	var v8332 int32
	_ = v8332
	var v8333 int32
	_ = v8333
	var v8334 int32
	_ = v8334
	var v8335 int32
	_ = v8335
	var v8338 int32
	_ = v8338
	var v8342 int32
	_ = v8342
	var v8346 int32
	_ = v8346
	var v8347 int32
	_ = v8347
	var v8351 int32
	_ = v8351
	var v8353 int32
	_ = v8353
	var v8356 int32
	_ = v8356
	var v8360 int32
	_ = v8360
	var v8366 int32
	_ = v8366
	var v8368 int32
	_ = v8368
	var v8371 int32
	_ = v8371
	var v8374 int32
	_ = v8374
	var v8375 int32
	_ = v8375
	var v8378 int32
	_ = v8378
	var v8380 int32
	_ = v8380
	var v8387 int32
	_ = v8387
	var v8388 int32
	_ = v8388
	var v8391 int32
	_ = v8391
	var v8393 int32
	_ = v8393
	var v8396 int32
	_ = v8396
	var v8397 int32
	_ = v8397
	var v8404 int32
	_ = v8404
	var v8408 int32
	_ = v8408
	var v8412 int32
	_ = v8412
	var v8416 int32
	_ = v8416
	var v8418 int32
	_ = v8418
	var v8420 int64
	_ = v8420
	var v8428 int32
	_ = v8428
	var v8433 int32
	_ = v8433
	var v8438 int32
	_ = v8438
	var v8443 int32
	_ = v8443
	var v8445 int32
	_ = v8445
	var v8448 int32
	_ = v8448
	var v8456 int32
	_ = v8456
	var v8459 int32
	_ = v8459
	var v8461 int32
	_ = v8461
	var v8462 int32
	_ = v8462
	var v8467 int32
	_ = v8467
	var v8471 int32
	_ = v8471
	var v8473 int32
	_ = v8473
	var v8477 int32
	_ = v8477
	var v8522 int32
	_ = v8522
	var v8524 int32
	_ = v8524
	var v8537 int32
	_ = v8537
	var v8571 int32
	_ = v8571
	var v8579 int32
	_ = v8579
	var v8583 int32
	_ = v8583
	var v8588 int32
	_ = v8588
	var v8592 int32
	_ = v8592
	var v8594 int32
	_ = v8594
	var v8600 int32
	_ = v8600
	var v8605 int32
	_ = v8605
	var v8609 int32
	_ = v8609
	var v8612 int32
	_ = v8612
	var v8620 int32
	_ = v8620
	var v8623 int32
	_ = v8623
	var v8629 int32
	_ = v8629
	var v8634 int32
	_ = v8634
	var v8638 int32
	_ = v8638
	var v8641 int32
	_ = v8641
	var v8649 int32
	_ = v8649
	var v8654 int32
	_ = v8654
	var v8658 int32
	_ = v8658
	var v8661 int32
	_ = v8661
	var v8669 int32
	_ = v8669
	var v8673 int32
	_ = v8673
	var v8678 int32
	_ = v8678
	var v8682 int32
	_ = v8682
	var v8685 int32
	_ = v8685
	var v8693 int32
	_ = v8693
	var v8698 int32
	_ = v8698
	var v8702 int32
	_ = v8702
	var v8706 int32
	_ = v8706
	var v8712 int32
	_ = v8712
	var v8716 int32
	_ = v8716
	var v8721 int32
	_ = v8721
	var v8725 int32
	_ = v8725
	var v8729 int32
	_ = v8729
	var v8735 int32
	_ = v8735
	var v8739 int32
	_ = v8739
	var v8744 int32
	_ = v8744
	var v8749 int32
	_ = v8749
	var v8752 int32
	_ = v8752
	var v8758 int32
	_ = v8758
	var v8762 int32
	_ = v8762
	var v8767 int32
	_ = v8767
	v7 = int32(0)
	v42 = m.G0
	v44 = v42 - int32(528)
	m.G0 = v44
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+444)) = v7
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
	v8749 = m.ExcPending
	if v8749 != 0 {
		goto L1
	} else {
		goto L1827
	}
L120:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8725 = m.ExcPending
	if v8725 != 0 {
		goto L1
	} else {
		goto L1822
	}
L121:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8702 = m.ExcPending
	if v8702 != 0 {
		goto L1
	} else {
		goto L1817
	}
L122:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8682 = m.ExcPending
	if v8682 != 0 {
		goto L1
	} else {
		goto L1813
	}
L123:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8658 = m.ExcPending
	if v8658 != 0 {
		goto L1
	} else {
		goto L1808
	}
L124:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8638 = m.ExcPending
	if v8638 != 0 {
		goto L1
	} else {
		goto L1804
	}
L125:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v8609 = m.ExcPending
	if v8609 != 0 {
		goto L1
	} else {
		goto L1799
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8592 = m.ExcPending
	if v8592 != 0 {
		goto L1
	} else {
		goto L1796
	}
L127:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8571 = m.ExcPending
	if v8571 != 0 {
		goto L1
	} else {
		goto L1792
	}
L128:
	;
	m.G0 = v8537 + int32(528)
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
		goto L147
	} else {
		goto L148
	}
L132:
	;
	v8537 = v44
	goto L128
L133:
	;
	F_pgstat_bestart_final(m)
	mBase = m.M
	v8522 = m.ExcPending
	if v8522 != 0 {
		goto L1
	} else {
		goto L1790
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31])) = v7680
	v7688 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v7688)+60)) = v7680
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v7691 = m.ExcPending
	if v7691 != 0 {
		goto L1
	} else {
		goto L1608
	}
L135:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7658 = m.ExcPending
	if v7658 != 0 {
		goto L1
	} else {
		goto L1603
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32])) = int32(1663)
	v7680 = int32(1)
	goto L134
L137:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7636 = m.ExcPending
	if v7636 != 0 {
		goto L1
	} else {
		goto L1599
	}
L138:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7613 = m.ExcPending
	if v7613 != 0 {
		goto L1
	} else {
		goto L1594
	}
L139:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7594 = m.ExcPending
	if v7594 != 0 {
		goto L1
	} else {
		goto L1590
	}
L140:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7575 = m.ExcPending
	if v7575 != 0 {
		goto L1
	} else {
		goto L1586
	}
L141:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7559 = m.ExcPending
	if v7559 != 0 {
		goto L1
	} else {
		goto L1582
	}
L142:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7542 = m.ExcPending
	if v7542 != 0 {
		goto L1
	} else {
		goto L1579
	}
L143:
	;
	v6960 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v6960 != 0 {
		goto L1439
	} else {
		goto L1440
	}
L144:
	;
	v6916 = F_superuser(m)
	mBase = m.M
	v6917 = m.ExcPending
	if v6917 != 0 {
		goto L1
	} else {
		goto L1438
	}
L145:
	;
	v1072 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[34])) = uint8(v1072)
	v1075 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[35]))
	if v1077 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L146:
	;
	F_InitializeSessionUserId(m, l2, l3, int32(base.Ui32(l4&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L187
	}
L147:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L186
	}
L148:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[36]))
	if v958 < int32(0) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L153
	}
L150:
	;
	v962 = F_GetCurrentTimestamp(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[37])) = v962
	goto L152
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[38])) = int32(1)
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	switch v970 - int32(4) {
	case 0, 3:
		goto L147
	default:
		goto L154
	}
L154:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v974 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v977 = int32(_a_F_InitPostgres_32)
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v983 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[39])))
	if base.B2i32(v980 == int32(0))|base.B2i32(v980 != v983) != 0 {
		v1001 = v980
		v1002 = v983
		goto L160
	} else {
		goto L161
	}
L156:
	;
	goto L157
L157:
	;
	if v970 != int32(5) {
		goto L145
	} else {
		goto L184
	}
L158:
	;
	v1018 = F_table_open(m, int32(1260), int32(1))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L172
	}
L159:
	;
	if v1001-v1002 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	goto L159
L161:
	;
	v986 = l2
	v987 = v977
	goto L162
L162:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v987)+1)))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+1)))
	if v991 == int32(0) {
		v1001 = v991
		v1002 = v990
		goto L160
	} else {
		goto L164
	}
L163:
	;
	v1001 = v991
	v1002 = v990
	goto L160
L164:
	;
	v994 = int32(1)
	if v991 == v990 {
		v986 = v986 + v994
		v987 = v987 + v994
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	F_InitializeSessionUserIdStandalone(m)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v1009 = int32(0)
	F_InitializeSessionUserId(m, l2, v1009, v1009)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	v1015 = int32(1)
	goto L158
L170:
	;
	v1013 = F_superuser(m)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v1015 = v1013
	goto L158
L172:
	;
	v1020 = int32(0)
	v1022 = F_table_beginscan_catalog(m, v1018, v1020, v1020)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v1024 = F_heap_getnext(m, v1022)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
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
		goto L175
	}
L175:
	;
	F_relation_close(m, v1018, int32(1))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	if v1024 != 0 {
		v6918 = l0
		v6919 = l1
		v6921 = v1015
		v6922 = l4
		v6923 = l5
		v6930 = v44
		v6942 = v47
		v6951 = v7
		goto L143
	} else {
		goto L177
	}
L177:
	;
	v1036 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	if v1036 == int32(0) {
		v6918 = l0
		v6919 = l1
		v6921 = v1015
		v6922 = l4
		v6923 = l5
		v6930 = v44
		v6942 = v47
		v6951 = v7
		goto L143
	} else {
		goto L179
	}
L179:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_33), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+432)) = l2
	F_errhint(m, int32(_a_F_InitPostgres_34), v44+int32(432))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(896), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v6918 = l0
	v6919 = l1
	v6921 = v1015
	v6922 = l4
	v6923 = l5
	v6930 = v44
	v6942 = v47
	v6951 = v7
	goto L143
L184:
	;
	if l2|l3 != 0 {
		goto L146
	} else {
		goto L185
	}
L185:
	;
	goto L147
L186:
	;
	v6918 = l0
	v6919 = l1
	v6921 = int32(1)
	v6922 = l4
	v6923 = l5
	v6930 = v44
	v6942 = v47
	v6951 = v7
	goto L143
L187:
	;
	v6875 = l0
	v6876 = l1
	v6879 = l4
	v6880 = l5
	v6887 = v44
	v6899 = v47
	v6908 = v7
	goto L144
L188:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v1087 = F_AllocSetContextCreateInternal(m, v1082, int32(_a_F_InitPostgres_35), int32(0), int32(_a_F_InitPostgres_6), int32(_a_F_InitPostgres_7))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1090 = F_load_hba(m)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[35])) = v1087
	goto L190
L192:
	;
	if v1090 == int32(0) {
		goto L142
	} else {
		goto L193
	}
L193:
	;
	v1094 = F_load_ident(m)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v1100 = m.G0
	v1101 = int32(16)
	v1102 = v1100 - v1101
	m.G0 = v1102
	F_gettimeofday(m, v1102)
	mBase = m.M
	v1105 = *(*int64)(unsafe.Add(mBase, uint32(v1102)))
	v1106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1102)+8)))
	m.G0 = v1102 + v1101
	goto L195
L195:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[40])) = v1106 + v1105*int64(1000000) - int64(946684800000000)
	v1118 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[41]))
	F_enable_timeout_after(m, int32(3), v1118*int32(1000))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v1123 = int32(0)
	v1124 = m.G0
	v1126 = v1124 - int32(3792)
	m.G0 = v1126
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+700)) = v1123
	v1130 = m.G0
	v1132 = v1130 - int32(288)
	m.G0 = v1132
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v1136 = F_get_role_oid(m, v1134, int32(1))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[42]))
	if v1139 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+380)) = v2589
	m.G0 = v1132 + int32(288)
	v2622 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v2622 != 0 {
		goto L514
	} else {
		goto L515
	}
L199:
	;
	v2572 = F_palloc0(m, int32(420))
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L1
	} else {
		goto L513
	}
L200:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+4))
	if v1142 <= int32(0) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v1146 = v1075 + int32(144)
	v1169 = v7
	goto L202
L202:
	;
	v1188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1146))))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+12))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1169<<(uint(int32(2))%32))))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+12))
	if v1194 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L203:
	;
	goto L199
L204:
	;
	v2527 = v1169 + int32(1)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+4))
	if v2527 < v2528 {
		v1169 = v2527
		goto L202
	} else {
		goto L512
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+288)) = v1227
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-2)
	goto L204
L206:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+16))
	if v2138 == int32(0) {
		goto L204
	} else {
		goto L414
	}
L207:
	;
	if v1188 == int32(1) {
		goto L206
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	if v1188 == int32(1) {
		goto L204
	} else {
		goto L211
	}
L210:
	;
	goto L204
L211:
	;
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+488)))
	if v1201 == int32(1) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+288))
	switch v1210 {
	case 0:
		goto L219
	case 1, 2:
		goto L218
	case 3:
		goto L206
	default:
		goto L204
	}
L213:
	;
	if base.Ui32(v1194-int32(3)) < base.Ui32(int32(2)) {
		goto L204
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	switch v1194 - int32(2) {
	case 0, 2:
		goto L204
	default:
		goto L212
	}
L216:
	;
	goto L212
L217:
	;
	v2030 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1193)+24)))
	if v1188 != v2030 {
		goto L204
	} else {
		goto L403
	}
L218:
	;
	v1464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1132)+24)) = uint8(v1464)
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+20)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+16)) = v1210
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44])) = v1464
	v1472 = v1132 + int32(16)
	v1473 = m.G0
	v1475 = v1473 - int32(144)
	m.G0 = v1475
	v1479 = m.G0
	v1481 = v1479 - int32(272)
	m.G0 = v1481
	v1484 = v1481 + int32(8)
	goto L288
L219:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+292))
	if v1211 == int32(0) {
		goto L217
	} else {
		goto L220
	}
L220:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	if v1214 < int32(0) {
		goto L204
	} else {
		goto L221
	}
L221:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v1217 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v1222 = v1132 + int32(16)
	v1224 = int32(0)
	v1227 = F_pg_getnameinfo_all(m, v1146, v1220, v1222, int32(255), v1224, v1224, int32(8))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	v1233 = v1217
	goto L224
L224:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	if v1234 == int32(46) {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	if v1227 != 0 {
		goto L205
	} else {
		goto L226
	}
L226:
	;
	v1229 = F_pstrdup(m, v1222)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+280)) = v1229
	v1233 = v1229
	goto L224
L228:
	;
	v1237 = F_strlen(m, v1211)
	mBase = m.M
	v1238 = F_strlen(m, v1233)
	mBase = m.M
	if base.Ui32(v1238) < base.Ui32(v1237) {
		goto L204
	} else {
		goto L231
	}
L229:
	;
	v1244 = v1233
	goto L230
L230:
	;
	v1247 = v1211
	v1248 = v1244
	goto L233
L231:
	;
	v1244 = v1233 + (v1238 - v1237)
	goto L230
L232:
	;
	if v1285 != 0 {
		goto L204
	} else {
		goto L245
	}
L233:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247))))
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1248))))
	if v1251 == v1252 {
		v1274 = v1251
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v1285 = int32(0)
	goto L232
L235:
	;
	v1276 = int32(1)
	if v1274 != 0 {
		v1247 = v1247 + v1276
		v1248 = v1248 + v1276
		goto L233
	} else {
		goto L244
	}
L236:
	;
	if base.Ui32((v1251-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1262 = v1251 | int32(32)
	goto L239
L238:
	;
	v1262 = v1251
	goto L239
L239:
	;
	if base.Ui32((v1252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1271 = v1252 | int32(32)
	goto L242
L241:
	;
	v1271 = v1252
	goto L242
L242:
	;
	if v1262 == v1271 {
		v1274 = v1262
		goto L235
	} else {
		goto L243
	}
L243:
	;
	v1285 = v1262 - v1271
	goto L232
L244:
	;
	goto L234
L245:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	if v1286 == int32(1) {
		goto L206
	} else {
		goto L246
	}
L246:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	v1290 = int32(0)
	v1294 = m.Env.Getaddrinfo(m, v1289, v1290, v1290, v1132+int32(284))
	mBase = m.M
	if v1294 == v1290 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+284))
	if v1297 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+288)) = v1294
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-2)
	goto L204
L250:
	;
	v1298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1146))))
	v1309 = v1297
	goto L253
L251:
	;
	goto L252
L252:
	;
	v1448 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L1
	} else {
		goto L280
	}
L253:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+20))
	v1343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1342))))
	if v1343 != v1298 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+20))
	F_emscripten_builtin_free(m, v1402)
	mBase = m.M
	F_emscripten_builtin_free(m, v1297)
	mBase = m.M
	goto L279
L255:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+28))
	if v1401 != 0 {
		v1309 = v1401
		goto L253
	} else {
		goto L278
	}
L256:
	;
	switch v1298 - int32(2) {
	case 0:
		goto L258
	default:
		goto L255
	case 8:
		goto L259
	}
L257:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+20))
	F_emscripten_builtin_free(m, v1396)
	mBase = m.M
	F_emscripten_builtin_free(m, v1297)
	mBase = m.M
	goto L277
L258:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1342)+4))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+148))
	if v1393 != v1394 {
		goto L255
	} else {
		goto L276
	}
L259:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+8)))
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+152)))
	if v1345 != v1346 {
		goto L255
	} else {
		goto L260
	}
L260:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+9)))
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+153)))
	if v1348 != v1349 {
		goto L255
	} else {
		goto L261
	}
L261:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+10)))
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+154)))
	if v1351 != v1352 {
		goto L255
	} else {
		goto L262
	}
L262:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+11)))
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+155)))
	if v1354 != v1355 {
		goto L255
	} else {
		goto L263
	}
L263:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+12)))
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+156)))
	if v1357 != v1358 {
		goto L255
	} else {
		goto L264
	}
L264:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+13)))
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+157)))
	if v1360 != v1361 {
		goto L255
	} else {
		goto L265
	}
L265:
	;
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+14)))
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+158)))
	if v1363 != v1364 {
		goto L255
	} else {
		goto L266
	}
L266:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+15)))
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+159)))
	if v1366 != v1367 {
		goto L255
	} else {
		goto L267
	}
L267:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+16)))
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+160)))
	if v1369 != v1370 {
		goto L255
	} else {
		goto L268
	}
L268:
	;
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+17)))
	v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+161)))
	if v1372 != v1373 {
		goto L255
	} else {
		goto L269
	}
L269:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+18)))
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+162)))
	if v1375 != v1376 {
		goto L255
	} else {
		goto L270
	}
L270:
	;
	v1378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+19)))
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+163)))
	if v1378 != v1379 {
		goto L255
	} else {
		goto L271
	}
L271:
	;
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+20)))
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+164)))
	if v1381 != v1382 {
		goto L255
	} else {
		goto L272
	}
L272:
	;
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+21)))
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+165)))
	if v1384 != v1385 {
		goto L255
	} else {
		goto L273
	}
L273:
	;
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+22)))
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+166)))
	if v1387 != v1388 {
		goto L255
	} else {
		goto L274
	}
L274:
	;
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+23)))
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+167)))
	if v1390 == v1391 {
		goto L257
	} else {
		goto L275
	}
L275:
	;
	goto L255
L276:
	;
	goto L257
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(1)
	goto L206
L278:
	;
	goto L254
L279:
	;
	goto L252
L280:
	;
	if v1448 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1132))) = v1211
	F_errmsg_internal(m, int32(_a_F_InitPostgres_36), v1132)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+284)) = int32(-1)
	goto L204
L284:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_37), int32(1158), int32(_a_F_InitPostgres_38))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	v1595 = int32(0)
	v1596 = F_socket(m, int32(16), int32(_a_F_InitPostgres_39), v1595)
	mBase = m.M
	if v1596 < v1595 {
		goto L298
	} else {
		goto L299
	}
L287:
	;
	goto L286
L288:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1484))) = uint8(v1464)
	v1493 = v1481 + int32(272)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493-int32(1)))) = uint8(v1464)
	goto L289
L289:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1484)+2)) = uint8(v1464)
	*(*uint8)(unsafe.Add(mBase, uint32(v1484)+1)) = uint8(v1464)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493-int32(3)))) = uint8(v1464)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493-int32(2)))) = uint8(v1464)
	goto L290
L290:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1484)+3)) = uint8(v1464)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493-int32(4)))) = uint8(v1464)
	goto L291
L291:
	;
	v1515 = int32(0)
	v1518 = (v1515 - v1484) & int32(3)
	v1519 = v1484 + v1518
	*(*int32)(unsafe.Add(mBase, uint32(v1519))) = v1515
	v1527 = (int32(264) - v1518) & int32(-4)
	v1528 = v1519 + v1527
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(4)))) = v1515
	if base.Ui32(v1527) < base.Ui32(int32(9)) {
		goto L287
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+8)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+4)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(8)))) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(12)))) = v1515
	if base.Ui32(v1527) < base.Ui32(int32(25)) {
		goto L287
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+24)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+20)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+16)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1519)+12)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(16)))) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(20)))) = v1515
	v1554 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1528-v1554))) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v1528-int32(28)))) = v1515
	v1563 = v1519&int32(4) | v1554
	v1564 = v1527 - v1563
	if base.Ui32(v1564) < base.Ui32(int32(32)) {
		goto L287
	} else {
		goto L294
	}
L294:
	;
	v1569 = base.I64_extend_i32_u(v1515) * int64(4294967297)
	v1572 = v1563 + v1519
	v1573 = v1564
	goto L295
L295:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1572)+24)) = v1569
	*(*int64)(unsafe.Add(mBase, uint32(v1572)+16)) = v1569
	*(*int64)(unsafe.Add(mBase, uint32(v1572)+8)) = v1569
	*(*int64)(unsafe.Add(mBase, uint32(v1572))) = v1569
	v1581 = int32(32)
	v1584 = v1573 - v1581
	if base.Ui32(int32(31)) < base.Ui32(v1584) {
		v1572 = v1572 + v1581
		v1573 = v1584
		goto L295
	} else {
		goto L297
	}
L296:
	;
	goto L287
L297:
	;
	goto L296
L298:
	;
	v1770 = int32(-1)
	goto L300
L299:
	;
	v1601 = int32(18)
	v1603 = v1481 + int32(8)
	v1604 = int32(0)
	v1606 = m.G0
	v1608 = v1606 + int32(-8192)
	m.G0 = v1608
	v1611 = int32(20)
	F___memset(m, v1608, v1604, v1611)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_InitPostgres[45]))) = uint8(v1604)
	*(*int32)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_InitPostgres[46]))) = int32(1)
	v1616 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_InitPostgres[47]))) = uint16(v1616)
	*(*uint16)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_InitPostgres[48]))) = uint16(v1601)
	*(*int32)(unsafe.Add(mBase, uint32(v1606)+uint32(_c_F_InitPostgres[49]))) = v1611
	v1622 = F_send(m, v1596, v1608, v1611)
	mBase = m.M
	if v1622 < v1604 {
		v1676 = v1622
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1481)+8))
	if v1770 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L301:
	;
	if v1676 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L302:
	;
	m.G0 = v1608 - int32(-8192)
	goto L301
L303:
	;
	v1625 = F_recv(m, v1596, v1608)
	mBase = m.M
	if v1625 <= int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1676 = int32(-1)
	goto L302
L305:
	;
	v1630 = v1625
	goto L306
L306:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1630) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	v1676 = int32(0)
	goto L302
L308:
	;
	goto L307
L309:
	;
	v1638 = v1608
	goto L312
L310:
	;
	goto L311
L311:
	;
	v1663 = F_recv(m, v1596, v1608)
	mBase = m.M
	if int32(0) < v1663 {
		v1630 = v1663
		goto L306
	} else {
		goto L317
	}
L312:
	;
	v1644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1638)+4)))
	switch v1644 - int32(2) {
	case 0:
		v1676 = int32(-1)
		goto L302
	case 1:
		goto L308
	default:
		goto L314
	}
L313:
	;
	goto L311
L314:
	;
	v1647 = F_netlink_msg_to_ifaddr(m, v1603, v1638)
	mBase = m.M
	if v1647 != 0 {
		v1676 = v1647
		goto L302
	} else {
		goto L315
	}
L315:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1638)))
	v1653 = v1638 + (v1648+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1630+v1608-v1653) {
		v1638 = v1653
		goto L312
	} else {
		goto L316
	}
L316:
	;
	goto L313
L317:
	;
	goto L304
L318:
	;
	v1686 = int32(22)
	v1687 = int32(0)
	v1689 = m.G0
	v1691 = v1689 + int32(-8192)
	m.G0 = v1691
	v1694 = int32(20)
	F___memset(m, v1691, v1687, v1694)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1689)+uint32(_c_F_InitPostgres[45]))) = uint8(v1687)
	*(*int32)(unsafe.Add(mBase, uint32(v1689)+uint32(_c_F_InitPostgres[46]))) = int32(2)
	v1699 = int32(769)
	*(*uint16)(unsafe.Add(mBase, uint32(v1689)+uint32(_c_F_InitPostgres[47]))) = uint16(v1699)
	*(*uint16)(unsafe.Add(mBase, uint32(v1689)+uint32(_c_F_InitPostgres[48]))) = uint16(v1686)
	*(*int32)(unsafe.Add(mBase, uint32(v1689)+uint32(_c_F_InitPostgres[49]))) = v1694
	v1705 = F_send(m, v1596, v1691, v1694)
	mBase = m.M
	if v1705 < v1687 {
		v1759 = v1705
		goto L322
	} else {
		goto L323
	}
L319:
	;
	v1766 = v1676
	goto L320
L320:
	;
	v1767 = m.Wasi_snapshot_preview1.Fd_close(m, v1596)
	mBase = m.M
	v1770 = v1766
	goto L300
L321:
	;
	v1766 = v1759
	goto L320
L322:
	;
	m.G0 = v1691 - int32(-8192)
	goto L321
L323:
	;
	v1708 = F_recv(m, v1596, v1691)
	mBase = m.M
	if v1708 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1759 = int32(-1)
	goto L322
L325:
	;
	v1713 = v1708
	goto L326
L326:
	;
	if base.Ui32(int32(16)) <= base.Ui32(v1713) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	v1759 = int32(0)
	goto L322
L328:
	;
	goto L327
L329:
	;
	v1721 = v1691
	goto L332
L330:
	;
	goto L331
L331:
	;
	v1746 = F_recv(m, v1596, v1691)
	mBase = m.M
	if int32(0) < v1746 {
		v1713 = v1746
		goto L326
	} else {
		goto L337
	}
L332:
	;
	v1727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1721)+4)))
	switch v1727 - int32(2) {
	case 0:
		v1759 = int32(-1)
		goto L322
	case 1:
		goto L328
	default:
		goto L334
	}
L333:
	;
	goto L331
L334:
	;
	v1730 = F_netlink_msg_to_ifaddr(m, v1603, v1721)
	mBase = m.M
	if v1730 != 0 {
		v1759 = v1730
		goto L322
	} else {
		goto L335
	}
L335:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1721)))
	v1736 = v1721 + (v1731+int32(3))&int32(-4)
	if base.Ui32(int32(15)) < base.Ui32(v1713+v1691-v1736) {
		v1721 = v1736
		goto L332
	} else {
		goto L336
	}
L336:
	;
	goto L333
L337:
	;
	goto L324
L338:
	;
	m.G0 = v1481 + int32(272)
	if v1770 < int32(0) {
		goto L349
	} else {
		goto L350
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1475+int32(12)))) = v1771
	goto L338
L340:
	;
	goto L341
L341:
	;
	if v1771 != 0 {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	goto L338
L343:
	;
	v1776 = v1771
	goto L346
L344:
	;
	goto L345
L345:
	;
	goto L342
L346:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1776)))
	F_emscripten_builtin_free(m, v1776)
	mBase = m.M
	if v1778 != 0 {
		v1776 = v1778
		goto L346
	} else {
		goto L348
	}
L347:
	;
	goto L345
L348:
	;
	goto L347
L349:
	;
	v2006 = int32(-1)
	goto L351
L350:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+12))
	if v1788 != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	m.G0 = v1475 + int32(144)
	if v2006 < int32(0) {
		goto L395
	} else {
		goto L396
	}
L352:
	;
	v1790 = v1475 + int32(24)
	v1800 = v1788
	goto L355
L353:
	;
	v1956 = int32(0)
	goto L354
L354:
	;
	if v1956 != 0 {
		goto L389
	} else {
		goto L390
	}
L355:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+12))
	if v1832 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+12))
	v1956 = v1913
	goto L354
L357:
	;
	v1833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1832))))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+16))
	if v1834 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L358:
	;
	goto L359
L359:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1800)))
	if v1912 != 0 {
		v1800 = v1912
		goto L355
	} else {
		goto L387
	}
L360:
	;
	v1871 = m.G0
	v1873 = v1871 - int32(128)
	m.G0 = v1873
	v1875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472)+8)))
	if v1875 == int32(0) {
		goto L375
	} else {
		goto L376
	}
L361:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1475)+16)) = uint16(v1833)
	v1868 = v1475 + int32(16)
	goto L360
L362:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1475)+16)) = int64(0)
	v1858 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1790)+8)) = v1858
	*(*int64)(unsafe.Add(mBase, uint32(v1790))) = v1858
	*(*int32)(unsafe.Add(mBase, uint32(v1475)+40)) = int32(0)
	goto L361
L363:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1475)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1475)+16)) = int64(-4294967296)
	goto L361
L364:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+4))
	if v1849 != 0 {
		v1868 = v1834
		goto L360
	} else {
		goto L373
	}
L365:
	;
	switch v1833 - int32(2) {
	case 0:
		goto L363
	default:
		v1868 = v1475 + int32(16)
		goto L360
	case 8:
		goto L362
	}
L366:
	;
	v1837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1834))))
	if v1837 != v1833 {
		goto L365
	} else {
		goto L367
	}
L367:
	;
	switch v1833 - int32(2) {
	case 0:
		goto L364
	default:
		v1868 = v1834
		goto L360
	case 8:
		goto L368
	}
L368:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+8))
	if v1841 != 0 {
		v1868 = v1834
		goto L360
	} else {
		goto L369
	}
L369:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+12))
	if v1842 != 0 {
		v1868 = v1834
		goto L360
	} else {
		goto L370
	}
L370:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+16))
	if v1843 != 0 {
		v1868 = v1834
		goto L360
	} else {
		goto L371
	}
L371:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1834)+20))
	if v1844 != 0 {
		v1868 = v1834
		goto L360
	} else {
		goto L372
	}
L372:
	;
	goto L362
L373:
	;
	goto L363
L374:
	;
	goto L359
L375:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1472)))
	if v1878 == int32(1) {
		goto L379
	} else {
		goto L380
	}
L376:
	;
	goto L377
L377:
	;
	m.G0 = v1873 + int32(128)
	goto L374
L378:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1472)+8)) = uint8(v1902)
	goto L377
L379:
	;
	v1881 = int32(0)
	v1883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1832))))
	v1884 = F_pg_sockaddr_cidr_mask(m, v1873, v1881, v1883)
	mBase = m.M
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	v1886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1885))))
	v1887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1832))))
	if v1886 != v1887 {
		v1902 = v1881
		goto L378
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	v1894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1893))))
	v1895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1832))))
	if v1894 != v1895 {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	v1889 = F_pg_range_sockaddr(m, v1885, v1832, v1873)
	mBase = m.M
	if v1889 == int32(0) {
		v1902 = v1881
		goto L378
	} else {
		goto L383
	}
L383:
	;
	v1902 = int32(1)
	goto L378
L384:
	;
	v1902 = int32(0)
	goto L378
L385:
	;
	v1897 = F_pg_range_sockaddr(m, v1893, v1832, v1868)
	mBase = m.M
	if v1897 == int32(0) {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1902 = int32(1)
	goto L378
L387:
	;
	goto L356
L388:
	;
	v2006 = int32(0)
	goto L351
L389:
	;
	v1958 = v1956
	goto L392
L390:
	;
	goto L391
L391:
	;
	goto L388
L392:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1958)))
	F_emscripten_builtin_free(m, v1958)
	mBase = m.M
	if v1960 != 0 {
		v1958 = v1960
		goto L392
	} else {
		goto L394
	}
L393:
	;
	goto L391
L394:
	;
	goto L393
L395:
	;
	v2014 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v2027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132)+24)))
	if v2027 == int32(0) {
		goto L204
	} else {
		goto L402
	}
L398:
	;
	if v2014 == int32(0) {
		goto L204
	} else {
		goto L399
	}
L399:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_40), int32(0))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_37), int32(1222), int32(_a_F_InitPostgres_41))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	goto L204
L402:
	;
	goto L206
L403:
	;
	v2038 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1146))))
	switch v2038 - int32(2) {
	case 0:
		goto L407
	default:
		goto L405
	case 8:
		goto L406
	}
L404:
	;
	if v2094 == int32(0) {
		goto L204
	} else {
		goto L413
	}
L405:
	;
	v2094 = int32(0)
	goto L404
L406:
	;
	v2049 = v1193 + int32(164)
	v2051 = v1193 + int32(32)
	v2053 = v1075 + int32(152)
	v2055 = int32(0)
	goto L408
L407:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1193+int32(156))+4))
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1193+int32(24))+4))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	v2094 = base.B2i32(v2041&(v2042^v2043) == int32(0))
	goto L404
L408:
	;
	v2061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055+v2049))))
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055+v2051))))
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2055+v2053))))
	if v2061&(v2063^v2065) != 0 {
		goto L405
	} else {
		goto L410
	}
L409:
	;
	v2094 = int32(1)
	goto L404
L410:
	;
	v2069 = v2055 | int32(1)
	v2071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2051+v2069))))
	v2073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2053+v2069))))
	v2076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2049+v2069))))
	if (v2071^v2073)&v2076 != 0 {
		goto L405
	} else {
		goto L411
	}
L411:
	;
	v2079 = v2055 + int32(2)
	if v2079 != int32(16) {
		v2055 = v2079
		goto L408
	} else {
		goto L412
	}
L412:
	;
	goto L409
L413:
	;
	goto L206
L414:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2138)+4))
	if v2141 <= int32(0) {
		goto L204
	} else {
		goto L415
	}
L415:
	;
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	v2155 = int32(0)
	goto L416
L416:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v2138)+12))
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2188+v2155<<(uint(int32(2))%32))))
	v2193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2192)+4)))
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v2195 != int32(1) {
		goto L420
	} else {
		goto L421
	}
L417:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+20))
	v2478 = F_check_role_2(m, v2476, v1136, v2477)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L1
	} else {
		goto L510
	}
L418:
	;
	goto L417
L419:
	;
	v2470 = v2155 + int32(1)
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2138)+4))
	if v2470 < v2471 {
		v2155 = v2470
		goto L416
	} else {
		goto L509
	}
L420:
	;
	if v2193&int32(1) == int32(0) {
		goto L432
	} else {
		goto L433
	}
L421:
	;
	v2199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	if v2199&int32(1) != 0 {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	if v2193&int32(1) != 0 {
		goto L419
	} else {
		goto L423
	}
L423:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	v2205 = int32(_a_F_InitPostgres_42)
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2204))))
	v2211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[52])))
	if base.B2i32(v2208 == int32(0))|base.B2i32(v2208 != v2211) != 0 {
		v2229 = v2208
		v2230 = v2211
		goto L425
	} else {
		goto L426
	}
L424:
	;
	if v2229-v2230 != 0 {
		goto L419
	} else {
		goto L431
	}
L425:
	;
	goto L424
L426:
	;
	v2214 = v2204
	v2215 = v2205
	goto L427
L427:
	;
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2215)+1)))
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2214)+1)))
	if v2219 == int32(0) {
		v2229 = v2219
		v2230 = v2218
		goto L425
	} else {
		goto L429
	}
L428:
	;
	v2229 = v2219
	v2230 = v2218
	goto L425
L429:
	;
	v2222 = int32(1)
	if v2219 == v2218 {
		v2214 = v2214 + v2222
		v2215 = v2215 + v2222
		goto L427
	} else {
		goto L430
	}
L430:
	;
	goto L428
L431:
	;
	goto L418
L432:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	v2237 = int32(_a_F_InitPostgres_43)
	v2240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236))))
	v2243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[53])))
	if base.B2i32(v2240 == int32(0))|base.B2i32(v2240 != v2243) != 0 {
		v2261 = v2240
		v2262 = v2243
		goto L436
	} else {
		goto L437
	}
L433:
	;
	goto L434
L434:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+8))
	if v2416 != 0 {
		goto L493
	} else {
		goto L494
	}
L435:
	;
	if v2261-v2262 == int32(0) {
		goto L418
	} else {
		goto L442
	}
L436:
	;
	goto L435
L437:
	;
	v2246 = v2236
	v2247 = v2237
	goto L438
L438:
	;
	v2250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2247)+1)))
	v2251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2246)+1)))
	if v2251 == int32(0) {
		v2261 = v2251
		v2262 = v2250
		goto L436
	} else {
		goto L440
	}
L439:
	;
	v2261 = v2251
	v2262 = v2250
	goto L436
L440:
	;
	v2254 = int32(1)
	if v2251 == v2250 {
		v2246 = v2246 + v2254
		v2247 = v2247 + v2254
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	v2266 = int32(_a_F_InitPostgres_44)
	v2269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236))))
	v2272 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[54])))
	if base.B2i32(v2269 == int32(0))|base.B2i32(v2269 != v2272) != 0 {
		v2290 = v2269
		v2291 = v2272
		goto L444
	} else {
		goto L445
	}
L443:
	;
	if v2290-v2291 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L444:
	;
	goto L443
L445:
	;
	v2275 = v2236
	v2276 = v2266
	goto L446
L446:
	;
	v2279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2276)+1)))
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2275)+1)))
	if v2280 == int32(0) {
		v2290 = v2280
		v2291 = v2279
		goto L444
	} else {
		goto L448
	}
L447:
	;
	v2290 = v2280
	v2291 = v2279
	goto L444
L448:
	;
	v2283 = int32(1)
	if v2280 == v2279 {
		v2275 = v2275 + v2283
		v2276 = v2276 + v2283
		goto L446
	} else {
		goto L449
	}
L449:
	;
	goto L447
L450:
	;
	v2297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2145))))
	v2300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2144))))
	if base.B2i32(v2297 == int32(0))|base.B2i32(v2297 != v2300) != 0 {
		v2318 = v2297
		v2319 = v2300
		goto L454
	} else {
		goto L455
	}
L451:
	;
	goto L452
L452:
	;
	v2323 = int32(_a_F_InitPostgres_45)
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236))))
	v2329 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[55])))
	if base.B2i32(v2326 == int32(0))|base.B2i32(v2326 != v2329) != 0 {
		v2347 = v2326
		v2348 = v2329
		goto L463
	} else {
		goto L464
	}
L453:
	;
	if v2318-v2319 == int32(0) {
		goto L418
	} else {
		goto L460
	}
L454:
	;
	goto L453
L455:
	;
	v2303 = v2145
	v2304 = v2144
	goto L456
L456:
	;
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2304)+1)))
	v2308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303)+1)))
	if v2308 == int32(0) {
		v2318 = v2308
		v2319 = v2307
		goto L454
	} else {
		goto L458
	}
L457:
	;
	v2318 = v2308
	v2319 = v2307
	goto L454
L458:
	;
	v2311 = int32(1)
	if v2308 == v2307 {
		v2303 = v2303 + v2311
		v2304 = v2304 + v2311
		goto L456
	} else {
		goto L459
	}
L459:
	;
	goto L457
L460:
	;
	goto L419
L461:
	;
	v2386 = int32(_a_F_InitPostgres_42)
	v2389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236))))
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[52])))
	if base.B2i32(v2389 == int32(0))|base.B2i32(v2389 != v2392) != 0 {
		v2410 = v2389
		v2411 = v2392
		goto L486
	} else {
		goto L487
	}
L462:
	;
	if v2347-v2348 != 0 {
		goto L469
	} else {
		goto L470
	}
L463:
	;
	goto L462
L464:
	;
	v2332 = v2236
	v2333 = v2323
	goto L465
L465:
	;
	v2336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2333)+1)))
	v2337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2332)+1)))
	if v2337 == int32(0) {
		v2347 = v2337
		v2348 = v2336
		goto L463
	} else {
		goto L467
	}
L466:
	;
	v2347 = v2337
	v2348 = v2336
	goto L463
L467:
	;
	v2340 = int32(1)
	if v2337 == v2336 {
		v2332 = v2332 + v2340
		v2333 = v2333 + v2340
		goto L465
	} else {
		goto L468
	}
L468:
	;
	goto L466
L469:
	;
	v2350 = int32(_a_F_InitPostgres_46)
	v2353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2236))))
	v2356 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[56])))
	if base.B2i32(v2353 == int32(0))|base.B2i32(v2353 != v2356) != 0 {
		v2374 = v2353
		v2375 = v2356
		goto L473
	} else {
		goto L474
	}
L470:
	;
	goto L471
L471:
	;
	if v1136 == int32(0) {
		goto L419
	} else {
		goto L480
	}
L472:
	;
	if v2374-v2375 != 0 {
		goto L461
	} else {
		goto L479
	}
L473:
	;
	goto L472
L474:
	;
	v2359 = v2236
	v2360 = v2350
	goto L475
L475:
	;
	v2363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2360)+1)))
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2359)+1)))
	if v2364 == int32(0) {
		v2374 = v2364
		v2375 = v2363
		goto L473
	} else {
		goto L477
	}
L476:
	;
	v2374 = v2364
	v2375 = v2363
	goto L473
L477:
	;
	v2367 = int32(1)
	if v2364 == v2363 {
		v2359 = v2359 + v2367
		v2360 = v2360 + v2367
		goto L475
	} else {
		goto L478
	}
L478:
	;
	goto L476
L479:
	;
	goto L471
L480:
	;
	v2380 = F_get_role_oid(m, v2145, int32(1))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L1
	} else {
		goto L481
	}
L481:
	;
	if v2380 == int32(0) {
		goto L419
	} else {
		goto L482
	}
L482:
	;
	v2384 = F_is_member_of_role_nosuper(m, v1136, v2380)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L483
	}
L483:
	;
	if v2384 != 0 {
		goto L418
	} else {
		goto L484
	}
L484:
	;
	goto L419
L485:
	;
	if v2410-v2411 == int32(0) {
		goto L419
	} else {
		goto L492
	}
L486:
	;
	goto L485
L487:
	;
	v2395 = v2236
	v2396 = v2386
	goto L488
L488:
	;
	v2399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2396)+1)))
	v2400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2395)+1)))
	if v2400 == int32(0) {
		v2410 = v2400
		v2411 = v2399
		goto L486
	} else {
		goto L490
	}
L489:
	;
	v2410 = v2400
	v2411 = v2399
	goto L486
L490:
	;
	v2403 = int32(1)
	if v2400 == v2399 {
		v2395 = v2395 + v2403
		v2396 = v2396 + v2403
		goto L488
	} else {
		goto L491
	}
L491:
	;
	goto L489
L492:
	;
	goto L434
L493:
	;
	v2417 = F_strlen(m, v2145)
	mBase = m.M
	v2422 = F_palloc(m, v2417<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L496
	}
L494:
	;
	goto L495
L495:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2192)))
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2437))))
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2145))))
	if base.B2i32(v2440 == int32(0))|base.B2i32(v2440 != v2443) != 0 {
		v2461 = v2440
		v2462 = v2443
		goto L502
	} else {
		goto L503
	}
L496:
	;
	v2424 = F_strlen(m, v2145)
	mBase = m.M
	v2425 = F_pg_mb2wchar_with_len(m, v2145, v2422, v2424)
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2192)+8))
	v2428 = int32(0)
	v2431 = F_pg_regexec(m, v2427, v2422, v2425, v2428, v2428, v2428)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	F_pfree(m, v2422)
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	if v2431 == int32(0) {
		goto L418
	} else {
		goto L500
	}
L500:
	;
	goto L419
L501:
	;
	if v2461-v2462 == int32(0) {
		goto L418
	} else {
		goto L508
	}
L502:
	;
	goto L501
L503:
	;
	v2446 = v2437
	v2447 = v2145
	goto L504
L504:
	;
	v2450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2447)+1)))
	v2451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2446)+1)))
	if v2451 == int32(0) {
		v2461 = v2451
		v2462 = v2450
		goto L502
	} else {
		goto L506
	}
L505:
	;
	v2461 = v2451
	v2462 = v2450
	goto L502
L506:
	;
	v2454 = int32(1)
	if v2451 == v2450 {
		v2446 = v2446 + v2454
		v2447 = v2447 + v2454
		goto L504
	} else {
		goto L507
	}
L507:
	;
	goto L505
L508:
	;
	goto L419
L509:
	;
	goto L204
L510:
	;
	if v2478 == int32(0) {
		goto L204
	} else {
		goto L511
	}
L511:
	;
	v2589 = v1193
	goto L198
L512:
	;
	goto L203
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2572)+296)) = int32(1)
	v2589 = v2572
	goto L198
L514:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v2626 = *(*int32)(unsafe.Add(mBase, uint32(v2625)+356))
	if v2626 == int32(0) {
		goto L521
	} else {
		goto L522
	}
L517:
	;
	goto L516
L518:
	;
	v6645 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[57])))
	if base.B2i32(v6645&int32(2) == int32(0))|v6616 != 0 {
		goto L1361
	} else {
		goto L1362
	}
L519:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = int32(0)
	v6627 = v47
	v6636 = v7
	goto L518
L520:
	;
	v6598 = int32(0)
	v6600 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_47), v1075, v6598, v6598)
	mBase = m.M
	v6601 = m.ExcPending
	if v6601 != 0 {
		goto L1
	} else {
		goto L1360
	}
L521:
	;
	v2629 = int32(-1)
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2625)+296))
	switch v2630 {
	case 0:
		goto L531
	case 1:
		goto L530
	case 2, 12:
		goto L519
	case 3:
		goto L528
	case 4:
		goto L526
	case 5, 6:
		goto L527
	default:
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v2629
		v6627 = v47
		v6636 = v7
		goto L518
	case 13:
		goto L525
	case 14:
		goto L529
	case 15:
		goto L520
	}
L522:
	;
	goto L523
L523:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6584 = m.ExcPending
	if v6584 != 0 {
		goto L1
	} else {
		goto L1356
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+104)) = int32(_a_F_InitPostgres_48)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+100)) = v2653
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+96)) = v1126 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_49), v1126+int32(96))
	mBase = m.M
	v6575 = m.ExcPending
	if v6575 != 0 {
		goto L1
	} else {
		goto L1354
	}
L525:
	;
	v4417 = *(*int32)(unsafe.Add(mBase, uint32(v2625)+372))
	if v4417 == int32(0) {
		goto L880
	} else {
		goto L881
	}
L526:
	;
	v4360 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4360 != 0 {
		goto L854
	} else {
		goto L855
	}
L527:
	;
	v4049 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4052 = F_get_role_password(m, v4049, v1126+int32(700))
	mBase = m.M
	v4053 = m.ExcPending
	if v4053 != 0 {
		goto L1
	} else {
		goto L766
	}
L528:
	;
	v3019 = v1126 + int32(1568)
	v3022 = int32(132)
	base.MemoryCopy(m, v3019, v1075+int32(144), v3022)
	v3025 = v1126 + int32(1432)
	base.MemoryCopy(m, v3025, v1075+int32(12), v3022)
	v3030 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+1708)) = v3030
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+1704)) = v3030
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1696))
	v3036 = v1126 + int32(1168)
	v3042 = F_pg_getnameinfo_all(m, v3019, v3034, v3036, int32(255), v1126+int32(1136), int32(32), int32(3))
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L1
	} else {
		goto L638
	}
L529:
	;
	v2938 = m.G0
	v2940 = v2938 - int32(16)
	m.G0 = v2940
	*(*int32)(unsafe.Add(mBase, uint32(v2940))) = int32(12)
	goto L616
L530:
	;
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v2681 = v1126 + int32(2752)
	v2683 = int32(0)
	v2686 = F_pg_getnameinfo_all(m, v1075+int32(144), v2679, v2681, int32(255), v2683, v2683, int32(1))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L1
	} else {
		goto L538
	}
L531:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+272))
	v2635 = v1126 + int32(2752)
	v2637 = int32(0)
	v2640 = F_pg_getnameinfo_all(m, v1075+int32(144), v2633, v2635, int32(255), v2637, v2637, int32(1))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v2643 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L1
	} else {
		goto L533
	}
L533:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2654 = int32(1)
	if base.B2i32(v2643&v2654 == int32(0))&base.B2i32(v2645 == v2654) != 0 {
		goto L524
	} else {
		goto L535
	}
L535:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+92)) = int32(_a_F_InitPostgres_48)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+88)) = v2661
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+84)) = v2653
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+80)) = v2635
	F_errmsg(m, int32(_a_F_InitPostgres_50), v1126+int32(80))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(468), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L1
	} else {
		goto L537
	}
L537:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L538:
	;
	v2689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	v2691 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v2698 = m.ExcPending
	if v2698 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	v2699 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v2700 = int32(1)
	if v2689&v2700|base.B2i32(v2691 != v2700) == int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+296)) = int32(_a_F_InitPostgres_48)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+292)) = v2699
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+288)) = v2681
	F_errmsg(m, int32(_a_F_InitPostgres_53), v1126+int32(288))
	mBase = m.M
	v2715 = m.ExcPending
	if v2715 != 0 {
		goto L1
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+204)) = int32(_a_F_InitPostgres_48)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+200)) = v2820
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+196)) = v2699
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+192)) = v1126 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_54), v1126+int32(192))
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L580
	}
L544:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v2717 != 0 {
		goto L546
	} else {
		goto L547
	}
L545:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(528), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v2819 = m.ExcPending
	if v2819 != 0 {
		goto L1
	} else {
		goto L579
	}
L546:
	;
	switch v2716 + int32(2) {
	case 0:
		goto L549
	case 1:
		goto L550
	case 2:
		goto L551
	case 3:
		goto L552
	default:
		goto L545
	}
L547:
	;
	goto L548
L548:
	;
	if v2716 != int32(-2) {
		goto L545
	} else {
		goto L567
	}
L549:
	;
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2741 = int32(_a_F_InitPostgres_55)
	v2743 = v2738 + int32(1)
	if v2743 == int32(0) {
		v2763 = v2741
		goto L557
	} else {
		goto L558
	}
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+256)) = v2717
	F_errdetail_log(m, int32(_a_F_InitPostgres_56), v1126+int32(256))
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L1
	} else {
		goto L555
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+240)) = v2717
	F_errdetail_log(m, int32(_a_F_InitPostgres_57), v1126+int32(240))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L1
	} else {
		goto L554
	}
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+224)) = v2717
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v1126+int32(224))
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	goto L545
L554:
	;
	goto L545
L555:
	;
	goto L545
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+276)) = v2763 + base.B2i32(v2765 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+272)) = v2717
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v1126+int32(272))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L1
	} else {
		goto L566
	}
L557:
	;
	v2765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2763))))
	goto L556
L558:
	;
	v2747 = v2741
	v2748 = v2743
	goto L559
L559:
	;
	v2749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747))))
	if v2749 == int32(0) {
		v2763 = v2747
		goto L557
	} else {
		goto L561
	}
L560:
	;
	v2763 = v2759
	goto L557
L561:
	;
	v2753 = v2747
	goto L562
L562:
	;
	v2757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2753)+1)))
	if v2757 != 0 {
		v2753 = v2753 + int32(1)
		goto L562
	} else {
		goto L564
	}
L563:
	;
	v2759 = v2753 + int32(2)
	v2761 = v2748 + int32(1)
	if v2761 != 0 {
		v2747 = v2759
		v2748 = v2761
		goto L559
	} else {
		goto L565
	}
L564:
	;
	goto L563
L565:
	;
	goto L560
L566:
	;
	goto L545
L567:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2781 = int32(_a_F_InitPostgres_55)
	v2783 = v2778 + int32(1)
	if v2783 == int32(0) {
		v2803 = v2781
		goto L569
	} else {
		goto L570
	}
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+208)) = v2803 + base.B2i32(v2805 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_60), v1126+int32(208))
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L1
	} else {
		goto L578
	}
L569:
	;
	v2805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2803))))
	goto L568
L570:
	;
	v2787 = v2781
	v2788 = v2783
	goto L571
L571:
	;
	v2789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2787))))
	if v2789 == int32(0) {
		v2803 = v2787
		goto L569
	} else {
		goto L573
	}
L572:
	;
	v2803 = v2799
	goto L569
L573:
	;
	v2793 = v2787
	goto L574
L574:
	;
	v2797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2793)+1)))
	if v2797 != 0 {
		v2793 = v2793 + int32(1)
		goto L574
	} else {
		goto L576
	}
L575:
	;
	v2799 = v2793 + int32(2)
	v2801 = v2788 + int32(1)
	if v2801 != 0 {
		v2787 = v2799
		v2788 = v2801
		goto L571
	} else {
		goto L577
	}
L576:
	;
	goto L575
L577:
	;
	goto L572
L578:
	;
	goto L545
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+284))
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+280))
	if v2834 != 0 {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(537), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v2936 = m.ExcPending
	if v2936 != 0 {
		goto L1
	} else {
		goto L615
	}
L582:
	;
	switch v2833 + int32(2) {
	case 0:
		goto L585
	case 1:
		goto L586
	case 2:
		goto L587
	case 3:
		goto L588
	default:
		goto L581
	}
L583:
	;
	goto L584
L584:
	;
	if v2833 != int32(-2) {
		goto L581
	} else {
		goto L603
	}
L585:
	;
	v2855 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2858 = int32(_a_F_InitPostgres_55)
	v2860 = v2855 + int32(1)
	if v2860 == int32(0) {
		v2880 = v2858
		goto L593
	} else {
		goto L594
	}
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+160)) = v2834
	F_errdetail_log(m, int32(_a_F_InitPostgres_56), v1126+int32(160))
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L1
	} else {
		goto L591
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+144)) = v2834
	F_errdetail_log(m, int32(_a_F_InitPostgres_57), v1126+int32(144))
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		goto L1
	} else {
		goto L590
	}
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+128)) = v2834
	F_errdetail_log(m, int32(_a_F_InitPostgres_58), v1126+int32(128))
	mBase = m.M
	v2842 = m.ExcPending
	if v2842 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	goto L581
L590:
	;
	goto L581
L591:
	;
	goto L581
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+180)) = v2880 + base.B2i32(v2882 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+176)) = v2834
	F_errdetail_log(m, int32(_a_F_InitPostgres_59), v1126+int32(176))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L1
	} else {
		goto L602
	}
L593:
	;
	v2882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880))))
	goto L592
L594:
	;
	v2864 = v2858
	v2865 = v2860
	goto L595
L595:
	;
	v2866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2864))))
	if v2866 == int32(0) {
		v2880 = v2864
		goto L593
	} else {
		goto L597
	}
L596:
	;
	v2880 = v2876
	goto L593
L597:
	;
	v2870 = v2864
	goto L598
L598:
	;
	v2874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2870)+1)))
	if v2874 != 0 {
		v2870 = v2870 + int32(1)
		goto L598
	} else {
		goto L600
	}
L599:
	;
	v2876 = v2870 + int32(2)
	v2878 = v2865 + int32(1)
	if v2878 != 0 {
		v2864 = v2876
		v2865 = v2878
		goto L595
	} else {
		goto L601
	}
L600:
	;
	goto L599
L601:
	;
	goto L596
L602:
	;
	goto L581
L603:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+288))
	v2898 = int32(_a_F_InitPostgres_55)
	v2900 = v2895 + int32(1)
	if v2900 == int32(0) {
		v2920 = v2898
		goto L605
	} else {
		goto L606
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+112)) = v2920 + base.B2i32(v2922 == int32(0))
	F_errdetail_log(m, int32(_a_F_InitPostgres_60), v1126+int32(112))
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L1
	} else {
		goto L614
	}
L605:
	;
	v2922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2920))))
	goto L604
L606:
	;
	v2904 = v2898
	v2905 = v2900
	goto L607
L607:
	;
	v2906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2904))))
	if v2906 == int32(0) {
		v2920 = v2904
		goto L605
	} else {
		goto L609
	}
L608:
	;
	v2920 = v2916
	goto L605
L609:
	;
	v2910 = v2904
	goto L610
L610:
	;
	v2914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2910)+1)))
	if v2914 != 0 {
		v2910 = v2910 + int32(1)
		goto L610
	} else {
		goto L612
	}
L611:
	;
	v2916 = v2910 + int32(2)
	v2918 = v2905 + int32(1)
	if v2918 != 0 {
		v2904 = v2916
		v2905 = v2918
		goto L607
	} else {
		goto L613
	}
L612:
	;
	goto L611
L613:
	;
	goto L608
L614:
	;
	goto L581
L615:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L616:
	;
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2940)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1126+int32(1168)))) = v2950
	v2954 = *(*int32)(unsafe.Add(mBase, uint32(v2940)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1126+int32(880)))) = v2954
	goto L618
L618:
	;
	m.G0 = v2940 + int32(16)
	goto L620
L620:
	;
	goto L621
L621:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44])) = int32(44)
	v3002 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3003 = m.ExcPending
	if v3003 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	if v3002 == int32(0) {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v2629
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L635
	}
L635:
	;
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1168))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+320)) = v3006
	F_errmsg(m, int32(_a_F_InitPostgres_61), v1126+int32(320))
	mBase = m.M
	v3012 = m.ExcPending
	if v3012 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(1888), int32(_a_F_InitPostgres_62))
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v2629
	v6627 = v47
	v6636 = v7
	goto L518
L638:
	;
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1560))
	v3046 = v1126 + int32(880)
	v3052 = F_pg_getnameinfo_all(m, v3025, v3044, v3046, int32(255), v1126+int32(848), int32(32), int32(3))
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L1
	} else {
		goto L639
	}
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+432)) = int32(113)
	v3057 = v1126 + int32(816)
	v3062 = F_pg_snprintf(m, v3057, int32(32), int32(_a_F_InitPostgres_63), v1126+int32(432))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+704)) = int32(4)
	v3066 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1126)+716)) = v3066
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+712)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v1126)+724)) = v3066
	v3072 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+732)) = v3072
	v3074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126)+1568)))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+708)) = v3074
	v3077 = v1126 + int32(704)
	v3080 = F_pg_getaddrinfo_all(m, v3036, v3057, v3077, v1126+int32(1708))
	mBase = m.M
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1708))
	if v3080|base.B2i32(v3081 == v3072) == v3072 {
		goto L641
	} else {
		goto L642
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+704)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+712)) = int32(1)
	v3091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126)+1432)))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+708)) = v3091
	v3094 = v1126 + int32(716)
	v3095 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3094)+16)) = v3095
	v3097 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3094)+8)) = v3097
	*(*int64)(unsafe.Add(mBase, uint32(v3094))) = v3097
	v3104 = F_pg_getaddrinfo_all(m, v3046, v3095, v3077, v1126+int32(1704))
	mBase = m.M
	if v3104 != 0 {
		v3937 = v1123
		goto L644
	} else {
		goto L645
	}
L642:
	;
	v3973 = v3081
	v3979 = v1123
	goto L643
L643:
	;
	if v3973 != 0 {
		goto L736
	} else {
		goto L737
	}
L644:
	;
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1708))
	v3973 = v3963
	v3979 = v3937
	goto L643
L645:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1704))
	if v3105 == int32(0) {
		v3937 = v1123
		goto L644
	} else {
		goto L646
	}
L646:
	;
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1708))
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v3108)+4))
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(v3108)+8))
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(v3108)+12))
	v3112 = F_socket(m, v3109, v3110, v3111)
	mBase = m.M
	if v3112 == int32(-1) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v3117 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L1
	} else {
		goto L650
	}
L648:
	;
	goto L649
L649:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1704))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v3133)+20))
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3133)+16))
	v3136 = F_bind(m, v3112, v3134, v3135)
	mBase = m.M
	if v3136 != 0 {
		goto L657
	} else {
		goto L658
	}
L650:
	;
	if v3117 == int32(0) {
		v3937 = v1123
		goto L644
	} else {
		goto L651
	}
L651:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3122 = m.ExcPending
	if v3122 != 0 {
		goto L1
	} else {
		goto L652
	}
L652:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_64), int32(0))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L1
	} else {
		goto L653
	}
L653:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(1742), int32(_a_F_InitPostgres_65))
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	v3937 = v1123
	goto L644
L655:
	;
	v3921 = F_close(m, v3112)
	mBase = m.M
	v3937 = v3895
	goto L644
L656:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), v3876, int32(_a_F_InitPostgres_65))
	mBase = m.M
	v3879 = m.ExcPending
	if v3879 != 0 {
		goto L1
	} else {
		goto L735
	}
L657:
	;
	v3139 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L660
	}
L658:
	;
	goto L659
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+388)) = v1126 + int32(848)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+384)) = v1126 + int32(1136)
	v3169 = F_pg_snprintf(m, v1126+int32(736), int32(80), int32(_a_F_InitPostgres_66), v1126+int32(384))
	mBase = m.M
	v3170 = m.ExcPending
	if v3170 != 0 {
		goto L1
	} else {
		goto L664
	}
L660:
	;
	if v3139 == int32(0) {
		v3895 = v1123
		goto L655
	} else {
		goto L661
	}
L661:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3144 = m.ExcPending
	if v3144 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+416)) = v1126 + int32(880)
	F_errmsg(m, int32(_a_F_InitPostgres_67), v1126+int32(416))
	mBase = m.M
	v3152 = m.ExcPending
	if v3152 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v3850 = v1123
	v3876 = int32(1758)
	goto L656
L664:
	;
	goto L666
L665:
	;
	v3322 = v1126 + int32(2752)
	v3324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3322+v3270))) = uint8(v3324)
	v3328 = v1126 + int32(1712)
	v3330 = m.G0
	v3332 = v3330 - int32(80)
	m.G0 = v3332
	v3334 = F_strlen(m, v3322)
	mBase = m.M
	if base.Ui32(v3334) < base.Ui32(int32(2)) {
		v3785 = v3324
		goto L692
	} else {
		goto L693
	}
L666:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v3213 != 0 {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	v3303 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L1
	} else {
		goto L688
	}
L668:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L1
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	v3217 = v1126 + int32(736)
	v3218 = F_strlen(m, v3217)
	mBase = m.M
	v3219 = F_pgmem_send(m, v3112, v3217, v3218)
	mBase = m.M
	if int32(0) <= v3219 {
		goto L672
	} else {
		goto L673
	}
L671:
	;
	goto L670
L672:
	;
	goto L675
L673:
	;
	goto L674
L674:
	;
	v3298 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v3298 == int32(27) {
		goto L666
	} else {
		goto L687
	}
L675:
	;
	v3264 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v3264 != 0 {
		goto L677
	} else {
		goto L678
	}
L676:
	;
	v3279 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L1
	} else {
		goto L683
	}
L677:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L1
	} else {
		goto L680
	}
L678:
	;
	goto L679
L679:
	;
	v3270 = F_pgmem_recv(m, v3112, v1126+int32(2752), int32(591))
	mBase = m.M
	if int32(0) <= v3270 {
		goto L665
	} else {
		goto L681
	}
L680:
	;
	goto L679
L681:
	;
	v3274 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v3274 == int32(27) {
		goto L675
	} else {
		goto L682
	}
L682:
	;
	goto L676
L683:
	;
	if v3279 == int32(0) {
		v3895 = v1123
		goto L655
	} else {
		goto L684
	}
L684:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+356)) = v1126 + int32(816)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+352)) = v1126 + int32(1168)
	F_errmsg(m, int32(_a_F_InitPostgres_68), v1126+int32(352))
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	v3850 = v1123
	v3876 = int32(1809)
	goto L656
L687:
	;
	goto L667
L688:
	;
	if v3303 == int32(0) {
		v3895 = v1123
		goto L655
	} else {
		goto L689
	}
L689:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v3308 = m.ExcPending
	if v3308 != 0 {
		goto L1
	} else {
		goto L690
	}
L690:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+340)) = v1126 + int32(816)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+336)) = v1126 + int32(1168)
	F_errmsg(m, int32(_a_F_InitPostgres_69), v1126+int32(336))
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L1
	} else {
		goto L691
	}
L691:
	;
	v3850 = v1123
	v3876 = int32(1792)
	goto L656
L692:
	;
	m.G0 = v3332 + int32(80)
	if v3785 != 0 {
		v3895 = int32(1)
		goto L655
	} else {
		goto L731
	}
L693:
	;
	v3340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3322+v3334-int32(2)))))
	if v3340 != int32(13) {
		v3785 = v3324
		goto L692
	} else {
		goto L694
	}
L694:
	;
	v3350 = v3322
	goto L695
L695:
	;
	v3384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3350))))
	if v3384 == int32(13) {
		v3785 = v3324
		goto L692
	} else {
		goto L697
	}
L696:
	;
	v3398 = v3350
	goto L701
L697:
	;
	if v3384 != int32(58) {
		goto L698
	} else {
		goto L699
	}
L698:
	;
	v3350 = v3350 + int32(1)
	goto L695
L699:
	;
	goto L700
L700:
	;
	goto L696
L701:
	;
	v3433 = v3398 + int32(1)
	v3434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3398)+1)))
	goto L703
L702:
	;
	v3450 = v3433
	v3451 = int32(0)
	goto L705
L703:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3434))%64)))&base.B2i32(base.Ui32(v3434) < base.Ui32(int32(33))) != 0 {
		v3398 = v3433
		goto L701
	} else {
		goto L704
	}
L704:
	;
	goto L702
L705:
	;
	v3484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3450))))
	if base.B2i32(v3484 == int32(13))|base.B2i32(v3484 == int32(58)) != 0 {
		goto L707
	} else {
		goto L708
	}
L706:
	;
	v3509 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3451+v3332))) = uint8(v3509)
	v3518 = v3450
	goto L711
L707:
	;
	goto L706
L708:
	;
	v3490 = base.I32_extend8_s(v3484)
	goto L709
L709:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3490))%64)))&base.B2i32(base.Ui32(v3490) < base.Ui32(int32(33)))|base.B2i32(base.Ui32(int32(78)) < base.Ui32(v3451)) != 0 {
		goto L707
	} else {
		goto L710
	}
L710:
	;
	v3502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3450))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3451+v3332))) = uint8(v3502)
	v3504 = int32(1)
	v3450 = v3450 + v3504
	v3451 = v3451 + v3504
	goto L705
L711:
	;
	v3554 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3518))))
	goto L713
L712:
	;
	v3562 = int32(0)
	v3563 = *(*int32)(unsafe.Add(mBase, uint32(v3332)))
	v3566 = *(*int32)(unsafe.Add(mBase, uint32(v3332)+3))
	if v3563^int32(1380275029)|(v3566^int32(_a_F_InitPostgres_70)) != 0 {
		v3785 = v3562
		goto L692
	} else {
		goto L715
	}
L713:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3554))%64)))&base.B2i32(base.Ui32(v3554) < base.Ui32(int32(33))) != 0 {
		v3518 = v3518 + int32(1)
		goto L711
	} else {
		goto L714
	}
L714:
	;
	goto L712
L715:
	;
	v3570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3518))))
	if v3570 != int32(58) {
		v3785 = v3562
		goto L692
	} else {
		goto L716
	}
L716:
	;
	v3581 = v3518
	goto L717
L717:
	;
	v3614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3581)+1)))
	if v3614 == int32(13) {
		v3785 = v3562
		goto L692
	} else {
		goto L719
	}
L718:
	;
	v3631 = v3581 + int32(2)
	goto L721
L719:
	;
	if v3614 != int32(58) {
		v3581 = v3581 + int32(1)
		goto L717
	} else {
		goto L720
	}
L720:
	;
	goto L718
L721:
	;
	v3666 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3631))))
	goto L723
L722:
	;
	v3674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3631))))
	if v3674 == int32(13) {
		v3740 = v3562
		goto L725
	} else {
		goto L726
	}
L723:
	;
	if base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(v3666))%64)))&base.B2i32(base.Ui32(v3666) < base.Ui32(int32(33))) != 0 {
		v3631 = v3631 + int32(1)
		goto L721
	} else {
		goto L724
	}
L724:
	;
	goto L722
L725:
	;
	v3772 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3740+v3328))) = uint8(v3772)
	v3785 = int32(1)
	goto L692
L726:
	;
	v3685 = v3631
	v3686 = int32(0)
	v3687 = v3674
	goto L727
L727:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3686+v3328))) = uint8(v3687)
	v3722 = v3686 + int32(1)
	v3723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3685)+1)))
	if v3723 == int32(13) {
		v3740 = v3722
		goto L725
	} else {
		goto L729
	}
L728:
	;
	v3740 = v3722
	goto L725
L729:
	;
	if base.Ui32(v3686) < base.Ui32(int32(511)) {
		v3685 = v3685 + int32(1)
		v3686 = v3722
		v3687 = v3723
		goto L727
	} else {
		goto L730
	}
L730:
	;
	goto L728
L731:
	;
	v3819 = int32(0)
	v3822 = F_errstart(m, int32(15), v3819)
	mBase = m.M
	v3823 = m.ExcPending
	if v3823 != 0 {
		goto L1
	} else {
		goto L732
	}
L732:
	;
	if v3822 == int32(0) {
		v3895 = v3819
		goto L655
	} else {
		goto L733
	}
L733:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+368)) = v1126 + int32(2752)
	F_errmsg(m, int32(_a_F_InitPostgres_71), v1126+int32(368))
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	v3850 = v3819
	v3876 = int32(1819)
	goto L656
L735:
	;
	v3895 = v3850
	goto L655
L736:
	;
	v4005 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126)+1568)))
	if v4005 == int32(1) {
		goto L741
	} else {
		goto L742
	}
L737:
	;
	goto L738
L738:
	;
	v4021 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+1704))
	if v4021 != 0 {
		goto L749
	} else {
		goto L750
	}
L739:
	;
	goto L738
L740:
	;
	goto L739
L741:
	;
	if v3973 == int32(0) {
		goto L740
	} else {
		goto L744
	}
L742:
	;
	goto L743
L743:
	;
	if v3973 == int32(0) {
		goto L740
	} else {
		goto L748
	}
L744:
	;
	v4011 = v3973
	goto L745
L745:
	;
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v4011)+28))
	v4013 = *(*int32)(unsafe.Add(mBase, uint32(v4011)+20))
	F_emscripten_builtin_free(m, v4013)
	mBase = m.M
	F_emscripten_builtin_free(m, v4011)
	mBase = m.M
	if v4012 != 0 {
		v4011 = v4012
		goto L745
	} else {
		goto L747
	}
L746:
	;
	goto L740
L747:
	;
	goto L746
L748:
	;
	F_freeaddrinfo(m, v3973)
	mBase = m.M
	goto L740
L749:
	;
	v4022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126)+1432)))
	if v4022 == int32(1) {
		goto L754
	} else {
		goto L755
	}
L750:
	;
	goto L751
L751:
	;
	if v3979 == int32(0) {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v2629
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L762
	}
L752:
	;
	goto L751
L753:
	;
	goto L752
L754:
	;
	if v4021 == int32(0) {
		goto L753
	} else {
		goto L757
	}
L755:
	;
	goto L756
L756:
	;
	if v4021 == int32(0) {
		goto L753
	} else {
		goto L761
	}
L757:
	;
	v4028 = v4021
	goto L758
L758:
	;
	v4029 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+28))
	v4030 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+20))
	F_emscripten_builtin_free(m, v4030)
	mBase = m.M
	F_emscripten_builtin_free(m, v4028)
	mBase = m.M
	if v4029 != 0 {
		v4028 = v4029
		goto L758
	} else {
		goto L760
	}
L759:
	;
	goto L753
L760:
	;
	goto L759
L761:
	;
	F_freeaddrinfo(m, v4021)
	mBase = m.M
	goto L753
L762:
	;
	v4041 = v1126 + int32(1712)
	F_set_authn_id(m, v1075, v4041)
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		goto L1
	} else {
		goto L763
	}
L763:
	;
	v4044 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(v4044)+300))
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4047 = F_check_usermap(m, v4045, v4046, v4041)
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		goto L1
	} else {
		goto L764
	}
L764:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v4047
	v6627 = v47
	v6636 = v7
	goto L518
L765:
	;
	v4061 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+296))
	if base.B2i32(v4062 != int32(5))|base.B2i32(v4060 != int32(1)) == int32(0) {
		goto L772
	} else {
		goto L773
	}
L766:
	;
	if v4052 == int32(0) {
		goto L767
	} else {
		goto L768
	}
L767:
	;
	v4057 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[58]))
	v4060 = v4057
	goto L765
L768:
	;
	goto L769
L769:
	;
	v4058 = F_get_password_type(m, v4052)
	mBase = m.M
	v4059 = m.ExcPending
	if v4059 != 0 {
		goto L1
	} else {
		goto L770
	}
L770:
	;
	v4060 = v4058
	goto L765
L771:
	;
	if v4052 != 0 {
		goto L848
	} else {
		goto L849
	}
L772:
	;
	v4073 = int32(0)
	v4077 = m.G0
	v4079 = v4077 - int32(16)
	m.G0 = v4079
	*(*int32)(unsafe.Add(mBase, uint32(v4079))) = v4073
	v4085 = F_open(m, int32(_a_F_InitPostgres_72), v4073, v4079)
	mBase = m.M
	if v4085 != int32(-1) {
		goto L776
	} else {
		goto L777
	}
L773:
	;
	goto L774
L774:
	;
	v4346 = F_CheckSASLAuth(m, int32(_a_F_InitPostgres_73), v1075, v4052, v1126+int32(700))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L1
	} else {
		goto L847
	}
L775:
	;
	if v4118 == int32(0) {
		goto L788
	} else {
		goto L789
	}
L776:
	;
	goto L780
L777:
	;
	v4118 = v4073
	goto L778
L778:
	;
	m.G0 = v4079 + int32(16)
	goto L775
L779:
	;
	v4113 = F_close(m, v4085)
	mBase = m.M
	v4118 = v4111
	goto L778
L780:
	;
	v4091 = v1126 + int32(2752)
	v4092 = int32(4)
	goto L781
L781:
	;
	v4097 = F_read(m, v4085, v4091, v4092)
	mBase = m.M
	if v4097 <= int32(0) {
		goto L783
	} else {
		goto L784
	}
L782:
	;
	v4111 = int32(1)
	goto L779
L783:
	;
	v4101 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v4101 == int32(27) {
		goto L781
	} else {
		goto L786
	}
L784:
	;
	goto L785
L785:
	;
	v4106 = v4092 - v4097
	if v4106 != 0 {
		v4091 = v4091 + v4097
		v4092 = v4106
		goto L781
	} else {
		goto L787
	}
L786:
	;
	v4111 = int32(0)
	goto L779
L787:
	;
	goto L782
L788:
	;
	v4127 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4128 = m.ExcPending
	if v4128 != 0 {
		goto L1
	} else {
		goto L791
	}
L789:
	;
	goto L790
L790:
	;
	F_sendAuthRequest(m, int32(5), v1126+int32(2752), int32(4))
	mBase = m.M
	v4145 = m.ExcPending
	if v4145 != 0 {
		goto L1
	} else {
		goto L795
	}
L791:
	;
	if v4127 == int32(0) {
		v4351 = v2629
		goto L771
	} else {
		goto L792
	}
L792:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_74), int32(0))
	mBase = m.M
	v4134 = m.ExcPending
	if v4134 != 0 {
		goto L1
	} else {
		goto L793
	}
L793:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(893), int32(_a_F_InitPostgres_75))
	mBase = m.M
	v4139 = m.ExcPending
	if v4139 != 0 {
		goto L1
	} else {
		goto L794
	}
L794:
	;
	v4351 = v2629
	goto L771
L795:
	;
	v4146 = F_recv_password_packet(m)
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L1
	} else {
		goto L796
	}
L796:
	;
	if v4146 == int32(0) {
		goto L797
	} else {
		goto L798
	}
L797:
	;
	v4351 = int32(-2)
	goto L771
L798:
	;
	goto L799
L799:
	;
	if v4052 == int32(0) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	F_pfree(m, v4146)
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L1
	} else {
		goto L803
	}
L801:
	;
	goto L802
L802:
	;
	v4155 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4158 = m.G0
	v4160 = v4158 - int32(128)
	m.G0 = v4160
	v4162 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4160)+28)) = v4162
	*(*int32)(unsafe.Add(mBase, uint32(v4160)+116)) = v4162
	v4166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4052))))
	if v4166 != int32(109) {
		goto L806
	} else {
		goto L807
	}
L803:
	;
	v4351 = v2629
	goto L771
L804:
	;
	m.G0 = v4160 + int32(128)
	F_pfree(m, v4146)
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L1
	} else {
		goto L846
	}
L805:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+700)) = v4333
	v4337 = int32(-1)
	goto L804
L806:
	;
	v4324 = F_parse_scram_secret(m, v4052, v4160+int32(120), v4160+int32(112), v4160+int32(116), v4160+int32(124), v4160+int32(32), v4160+int32(80))
	mBase = m.M
	v4325 = m.ExcPending
	if v4325 != 0 {
		goto L1
	} else {
		goto L844
	}
L807:
	;
	v4169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4052)+1)))
	if v4169 != int32(100) {
		goto L806
	} else {
		goto L808
	}
L808:
	;
	v4172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4052)+2)))
	if v4172 != int32(53) {
		goto L806
	} else {
		goto L809
	}
L809:
	;
	v4175 = F_strlen(m, v4052)
	mBase = m.M
	if v4175 != int32(35) {
		goto L806
	} else {
		goto L810
	}
L810:
	;
	v4179 = v4052 + int32(3)
	v4180 = int32(_a_F_InitPostgres_76)
	v4184 = m.G0
	v4186 = v4184 - int32(32)
	v4187 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4186)+24)) = v4187
	*(*int64)(unsafe.Add(mBase, uint32(v4186)+16)) = v4187
	*(*int64)(unsafe.Add(mBase, uint32(v4186)+8)) = v4187
	*(*int64)(unsafe.Add(mBase, uint32(v4186))) = v4187
	v4195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[59])))
	if v4195 == int32(0) {
		goto L812
	} else {
		goto L813
	}
L811:
	;
	if v4263 != int32(32) {
		goto L806
	} else {
		goto L830
	}
L812:
	;
	v4263 = int32(0)
	goto L811
L813:
	;
	goto L814
L814:
	;
	v4199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[60])))
	if v4199 == int32(0) {
		goto L815
	} else {
		goto L816
	}
L815:
	;
	v4203 = v4179
	goto L818
L816:
	;
	goto L817
L817:
	;
	v4213 = v4180
	v4214 = v4195
	goto L821
L818:
	;
	v4209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4203))))
	if v4209 == v4195 {
		v4203 = v4203 + int32(1)
		goto L818
	} else {
		goto L820
	}
L819:
	;
	v4263 = v4203 - v4179
	goto L811
L820:
	;
	goto L819
L821:
	;
	v4221 = v4186 + int32(base.Ui32(v4214)>>(uint(int32(3))%32))&int32(28)
	v4222 = *(*int32)(unsafe.Add(mBase, uint32(v4221)))
	v4223 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4221))) = v4222 | v4223<<(uint(v4214)%32)
	v4227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4213)+1)))
	if v4227 != 0 {
		v4213 = v4213 + v4223
		v4214 = v4227
		goto L821
	} else {
		goto L823
	}
L822:
	;
	v4230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4179))))
	if v4230 == int32(0) {
		v4253 = v4179
		goto L824
	} else {
		goto L825
	}
L823:
	;
	goto L822
L824:
	;
	v4263 = v4253 - v4179
	goto L811
L825:
	;
	v4234 = v4179
	v4235 = v4230
	goto L826
L826:
	;
	v4243 = *(*int32)(unsafe.Add(mBase, uint32(v4186+int32(base.Ui32(v4235)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v4243)>>(uint(v4235)%32))&int32(1) == int32(0) {
		v4253 = v4234
		goto L824
	} else {
		goto L828
	}
L827:
	;
	v4253 = v4251
	goto L824
L828:
	;
	v4249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4234)+1)))
	v4251 = v4234 + int32(1)
	if v4249 != 0 {
		v4234 = v4251
		v4235 = v4249
		goto L826
	} else {
		goto L829
	}
L829:
	;
	goto L827
L830:
	;
	v4271 = F_pg_md5_encrypt(m, v4179, v1126+int32(2752), int32(4), v4160+int32(32), v4160+int32(28))
	mBase = m.M
	v4272 = m.ExcPending
	if v4272 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	if v4271 == int32(0) {
		goto L832
	} else {
		goto L833
	}
L832:
	;
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(v4160)+28))
	v4333 = v4275
	goto L805
L833:
	;
	goto L834
L834:
	;
	v4276 = int32(0)
	v4278 = v4160 + int32(32)
	v4281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4146))))
	v4284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4278))))
	if base.B2i32(v4281 == v4276)|base.B2i32(v4281 != v4284) != 0 {
		v4302 = v4281
		v4303 = v4284
		goto L836
	} else {
		goto L837
	}
L835:
	;
	if v4302-v4303 == int32(0) {
		v4337 = v4276
		goto L804
	} else {
		goto L842
	}
L836:
	;
	goto L835
L837:
	;
	v4287 = v4146
	v4288 = v4278
	goto L838
L838:
	;
	v4291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4288)+1)))
	v4292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4287)+1)))
	if v4292 == int32(0) {
		v4302 = v4292
		v4303 = v4291
		goto L836
	} else {
		goto L840
	}
L839:
	;
	v4302 = v4292
	v4303 = v4291
	goto L836
L840:
	;
	v4295 = int32(1)
	if v4292 == v4291 {
		v4287 = v4287 + v4295
		v4288 = v4288 + v4295
		goto L838
	} else {
		goto L841
	}
L841:
	;
	goto L839
L842:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4160))) = v4155
	v4309 = F_psprintf(m, int32(_a_F_InitPostgres_77), v4160)
	mBase = m.M
	v4310 = m.ExcPending
	if v4310 != 0 {
		goto L1
	} else {
		goto L843
	}
L843:
	;
	v4333 = v4309
	goto L805
L844:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4160)+16)) = v4155
	v4330 = F_psprintf(m, int32(_a_F_InitPostgres_78), v4160+int32(16))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L1
	} else {
		goto L845
	}
L845:
	;
	v4333 = v4330
	goto L805
L846:
	;
	v4351 = v4337
	goto L771
L847:
	;
	v4351 = v4346
	goto L771
L848:
	;
	F_pfree(m, v4052)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L851
	}
L849:
	;
	goto L850
L850:
	;
	if v4351 != 0 {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v4351
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L852
	}
L851:
	;
	goto L850
L852:
	;
	v4355 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	F_set_authn_id(m, v1075, v4355)
	mBase = m.M
	v4357 = m.ExcPending
	if v4357 != 0 {
		goto L1
	} else {
		goto L853
	}
L853:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = int32(0)
	v6627 = v47
	v6636 = v7
	goto L518
L854:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L857
	}
L855:
	;
	goto L856
L856:
	;
	v4364 = v1126 + int32(2752)
	F_pq_beginmessage(m, v4364, int32(82))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L858
	}
L857:
	;
	goto L856
L858:
	;
	F_enlargeStringInfo(m, v4364, int32(4))
	mBase = m.M
	v4370 = m.ExcPending
	if v4370 != 0 {
		goto L1
	} else {
		goto L859
	}
L859:
	;
	v4371 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+2756))
	v4372 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v4371+v4372))) = int32(50331648)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+2756)) = v4371 + int32(4)
	F_pq_endmessage(m, v4364)
	mBase = m.M
	v4380 = m.ExcPending
	if v4380 != 0 {
		goto L1
	} else {
		goto L860
	}
L860:
	;
	v4382 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[61]))
	v4383 = *(*int32)(unsafe.Add(mBase, uint32(v4382)+4))
	v4384 = m.T0[v4383].(func(*base.Module) int32)(m)
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	v4387 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4387 != 0 {
		goto L862
	} else {
		goto L863
	}
L862:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4389 = m.ExcPending
	if v4389 != 0 {
		goto L1
	} else {
		goto L865
	}
L863:
	;
	goto L864
L864:
	;
	v4390 = F_recv_password_packet(m)
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L1
	} else {
		goto L866
	}
L865:
	;
	goto L864
L866:
	;
	if v4390 == int32(0) {
		goto L867
	} else {
		goto L868
	}
L867:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = int32(-2)
	v6627 = v47
	v6636 = v7
	goto L518
L868:
	;
	goto L869
L869:
	;
	v4395 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4398 = F_get_role_password(m, v4395, v1126+int32(700))
	mBase = m.M
	v4399 = m.ExcPending
	if v4399 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	if v4398 == int32(0) {
		goto L871
	} else {
		goto L872
	}
L871:
	;
	F_pfree(m, v4390)
	mBase = m.M
	v4403 = m.ExcPending
	if v4403 != 0 {
		goto L1
	} else {
		goto L874
	}
L872:
	;
	goto L873
L873:
	;
	v4404 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	v4407 = F_plain_crypt_verify(m, v4404, v4398, v4390, v1126+int32(700))
	mBase = m.M
	v4408 = m.ExcPending
	if v4408 != 0 {
		goto L1
	} else {
		goto L875
	}
L874:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v2629
	v6627 = v47
	v6636 = v7
	goto L518
L875:
	;
	F_pfree(m, v4398)
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L1
	} else {
		goto L876
	}
L876:
	;
	F_pfree(m, v4390)
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L1
	} else {
		goto L877
	}
L877:
	;
	if v4407 != 0 {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v4407
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L878
	}
L878:
	;
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+364))
	F_set_authn_id(m, v1075, v4413)
	mBase = m.M
	v4415 = m.ExcPending
	if v4415 != 0 {
		goto L1
	} else {
		goto L879
	}
L879:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = int32(0)
	v6627 = v47
	v6636 = v7
	goto L518
L880:
	;
	v4422 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L883
	}
L881:
	;
	goto L882
L882:
	;
	v4435 = *(*int32)(unsafe.Add(mBase, uint32(v2625)+380))
	if v4435 == int32(0) {
		goto L887
	} else {
		goto L888
	}
L883:
	;
	if v4422 == int32(0) {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v2629
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L884
	}
L884:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_79), int32(0))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2860), int32(_a_F_InitPostgres_80))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v2629
	v6627 = v47
	v6636 = v7
	goto L518
L887:
	;
	v4440 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4441 = m.ExcPending
	if v4441 != 0 {
		goto L1
	} else {
		goto L890
	}
L888:
	;
	goto L889
L889:
	;
	v4454 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4454 != 0 {
		goto L894
	} else {
		goto L895
	}
L890:
	;
	if v4440 == int32(0) {
		v6603 = l0
		v6604 = l1
		v6605 = l2
		v6606 = l3
		v6607 = l4
		v6608 = l5
		v6609 = v1126
		v6614 = v1075
		v6615 = v44
		v6616 = v2629
		v6627 = v47
		v6636 = v7
		goto L518
	} else {
		goto L891
	}
L891:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_81), int32(0))
	mBase = m.M
	v4447 = m.ExcPending
	if v4447 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2867), int32(_a_F_InitPostgres_80))
	mBase = m.M
	v4452 = m.ExcPending
	if v4452 != 0 {
		goto L1
	} else {
		goto L893
	}
L893:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v2629
	v6627 = v47
	v6636 = v7
	goto L518
L894:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4456 = m.ExcPending
	if v4456 != 0 {
		goto L1
	} else {
		goto L897
	}
L895:
	;
	goto L896
L896:
	;
	v4458 = v1126 + int32(2752)
	F_pq_beginmessage(m, v4458, int32(82))
	mBase = m.M
	v4461 = m.ExcPending
	if v4461 != 0 {
		goto L1
	} else {
		goto L898
	}
L897:
	;
	goto L896
L898:
	;
	F_enlargeStringInfo(m, v4458, int32(4))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L1
	} else {
		goto L899
	}
L899:
	;
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+2756))
	v4466 = *(*int32)(unsafe.Add(mBase, uint32(v1126)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v4465+v4466))) = int32(50331648)
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+2756)) = v4465 + int32(4)
	F_pq_endmessage(m, v4458)
	mBase = m.M
	v4474 = m.ExcPending
	if v4474 != 0 {
		goto L1
	} else {
		goto L900
	}
L900:
	;
	v4476 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[61]))
	v4477 = *(*int32)(unsafe.Add(mBase, uint32(v4476)+4))
	v4478 = m.T0[v4477].(func(*base.Module) int32)(m)
	mBase = m.M
	v4479 = m.ExcPending
	if v4479 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	v4481 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v4481 != 0 {
		goto L902
	} else {
		goto L903
	}
L902:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L1
	} else {
		goto L905
	}
L903:
	;
	goto L904
L904:
	;
	v4484 = F_recv_password_packet(m)
	mBase = m.M
	v4485 = m.ExcPending
	if v4485 != 0 {
		goto L1
	} else {
		goto L906
	}
L905:
	;
	goto L904
L906:
	;
	if v4484 == int32(0) {
		goto L907
	} else {
		goto L908
	}
L907:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = int32(-2)
	v6627 = v47
	v6636 = v7
	goto L518
L908:
	;
	goto L909
L909:
	;
	v4489 = F_strlen(m, v4484)
	mBase = m.M
	if base.Ui32(int32(129)) <= base.Ui32(v4489) {
		goto L910
	} else {
		goto L911
	}
L910:
	;
	v4494 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4495 = m.ExcPending
	if v4495 != 0 {
		goto L1
	} else {
		goto L913
	}
L911:
	;
	goto L912
L912:
	;
	v4510 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+380))
	v4511 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+380))
	if v4511 != 0 {
		goto L920
	} else {
		goto L921
	}
L913:
	;
	if v4494 != 0 {
		goto L914
	} else {
		goto L915
	}
L914:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+448)) = int32(128)
	F_errmsg(m, int32(_a_F_InitPostgres_82), v1126+int32(448))
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		goto L1
	} else {
		goto L917
	}
L915:
	;
	goto L916
L916:
	;
	F_pfree(m, v4484)
	mBase = m.M
	v4509 = m.ExcPending
	if v4509 != 0 {
		goto L1
	} else {
		goto L919
	}
L917:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2881), int32(_a_F_InitPostgres_80))
	mBase = m.M
	v4507 = m.ExcPending
	if v4507 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	goto L916
L919:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v2629
	v6627 = v47
	v6636 = v7
	goto L518
L920:
	;
	v4512 = *(*int32)(unsafe.Add(mBase, uint32(v4511)+12))
	v4513 = v4512
	goto L922
L921:
	;
	v4513 = v7
	goto L922
L922:
	;
	v4514 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+396))
	if v4514 != 0 {
		goto L923
	} else {
		goto L924
	}
L923:
	;
	v4515 = *(*int32)(unsafe.Add(mBase, uint32(v4514)+12))
	v4516 = v4515
	goto L925
L924:
	;
	v4516 = v7
	goto L925
L925:
	;
	v4517 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+388))
	if v4517 != 0 {
		goto L926
	} else {
		goto L927
	}
L926:
	;
	v4518 = *(*int32)(unsafe.Add(mBase, uint32(v4517)+12))
	v4520 = v4518
	goto L928
L927:
	;
	v4520 = int32(0)
	goto L928
L928:
	;
	v4521 = *(*int32)(unsafe.Add(mBase, uint32(v4510)+372))
	if v4521 == int32(0) {
		v6522 = l0
		v6523 = l1
		v6524 = l2
		v6525 = l3
		v6526 = l4
		v6527 = l5
		v6528 = v1126
		v6533 = v1075
		v6534 = v44
		v6535 = v2629
		v6540 = v4484
		v6546 = v47
		v6555 = v7
		goto L929
	} else {
		goto L930
	}
L929:
	;
	F_pfree(m, v6540)
	mBase = m.M
	v6564 = m.ExcPending
	if v6564 != 0 {
		goto L1
	} else {
		goto L1353
	}
L930:
	;
	v4524 = *(*int32)(unsafe.Add(mBase, uint32(v4521)+4))
	if v4524 <= int32(0) {
		v6522 = l0
		v6523 = l1
		v6524 = l2
		v6525 = l3
		v6526 = l4
		v6527 = l5
		v6528 = v1126
		v6533 = v1075
		v6534 = v44
		v6535 = v2629
		v6540 = v4484
		v6546 = v47
		v6555 = v7
		goto L929
	} else {
		goto L931
	}
L931:
	;
	v4535 = l0
	v4536 = l1
	v4537 = l2
	v4538 = l3
	v4539 = l4
	v4540 = l5
	v4541 = v1126
	v4546 = v1075
	v4547 = v44
	v4548 = v2629
	v4549 = v1126 + int32(2756)
	v4553 = v4484
	v4555 = v4520
	v4558 = v4521
	v4559 = v47
	v4560 = v4516
	v4562 = v4513
	v4566 = v1126 + int32(1440)
	v4567 = v7
	v4568 = v7
	v4569 = v1126 + int32(1732)
	v4570 = v1126 + int32(1724)
	goto L932
L932:
	;
	v4576 = *(*int32)(unsafe.Add(mBase, uint32(v4558)+12))
	if v4560 != 0 {
		goto L934
	} else {
		goto L935
	}
L933:
	;
	v6522 = v6422
	v6523 = v6423
	v6524 = v6424
	v6525 = v6425
	v6526 = v6426
	v6527 = v6427
	v6528 = v6428
	v6533 = v6433
	v6534 = v6434
	v6535 = v6435
	v6540 = v6440
	v6546 = v6446
	v6555 = v6455
	goto L929
L934:
	;
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v4560)))
	v4582 = v4581
	goto L936
L935:
	;
	v4582 = int32(0)
	goto L936
L936:
	;
	v4583 = *(*int32)(unsafe.Add(mBase, uint32(v4576+v4567<<(uint(int32(2))%32))))
	if v4555 != 0 {
		goto L937
	} else {
		goto L938
	}
L937:
	;
	v4584 = *(*int32)(unsafe.Add(mBase, uint32(v4555)))
	v4586 = v4584
	goto L939
L938:
	;
	v4586 = int32(0)
	goto L939
L939:
	;
	v4587 = *(*int32)(unsafe.Add(mBase, uint32(v4562)))
	v4588 = *(*int32)(unsafe.Add(mBase, uint32(v4546)+364))
	v4589 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1576)) = v4589
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1568)) = v4589
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1584)) = v4589
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1592)) = v4589
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1576)) = int32(2)
	if v4582 != 0 {
		goto L940
	} else {
		goto L941
	}
L940:
	;
	v4600 = v4582
	goto L942
L941:
	;
	v4600 = int32(_a_F_InitPostgres_83)
	goto L942
L942:
	;
	v4604 = v4600
	goto L944
L943:
	;
	v4653 = F_pg_getaddrinfo_all(m, v4583, v4600, v4541+int32(1568), v4541+int32(704))
	mBase = m.M
	if v4653 == int32(0) {
		goto L961
	} else {
		goto L962
	}
L944:
	;
	v4609 = v4604 + int32(1)
	v4610 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4604))))
	v4611 = F___isspace(m, v4610)
	mBase = m.M
	if v4611 != 0 {
		v4604 = v4609
		goto L944
	} else {
		goto L946
	}
L945:
	;
	v4612 = int32(1)
	switch v4610&int32(255) - int32(43) {
	case 0:
		v4618 = v4612
		goto L948
	default:
		v4620 = v4610
		v4621 = v4604
		v4622 = v4612
		goto L947
	case 2:
		goto L949
	}
L946:
	;
	goto L945
L947:
	;
	v4623 = int32(0)
	v4625 = v4620 - int32(48)
	if base.Ui32(v4625) <= base.Ui32(int32(9)) {
		goto L950
	} else {
		goto L951
	}
L948:
	;
	v4619 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4609))))
	v4620 = v4619
	v4621 = v4609
	v4622 = v4618
	goto L947
L949:
	;
	v4618 = int32(0)
	goto L948
L950:
	;
	v4628 = v4623
	v4629 = v4625
	v4630 = v4621
	goto L953
L951:
	;
	v4642 = v4623
	goto L952
L952:
	;
	if v4622 != 0 {
		goto L956
	} else {
		goto L957
	}
L953:
	;
	v4632 = int32(10)
	v4634 = v4628*v4632 - v4629
	v4635 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4630)+1)))
	v4639 = v4635 - int32(48)
	if base.Ui32(v4639) < base.Ui32(v4632) {
		v4628 = v4634
		v4629 = v4639
		v4630 = v4630 + int32(1)
		goto L953
	} else {
		goto L955
	}
L954:
	;
	v4642 = v4634
	goto L952
L955:
	;
	goto L954
L956:
	;
	v4648 = int32(0) - v4642
	goto L958
L957:
	;
	v4648 = v4642
	goto L958
L958:
	;
	goto L943
L959:
	;
	v6463 = *(*int32)(unsafe.Add(mBase, uint32(v6433)+380))
	v6464 = *(*int32)(unsafe.Add(mBase, uint32(v6463)+380))
	if v6464 == int32(0) {
		v6481 = v6449
		goto L1334
	} else {
		goto L1335
	}
L960:
	;
	v4722 = int32(20)
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v4722)
	v4724 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4541)+2752)) = uint8(v4724)
	v4726 = int32(16)
	v4727 = int32(0)
	v4731 = m.G0
	v4733 = v4731 - v4726
	m.G0 = v4733
	*(*int32)(unsafe.Add(mBase, uint32(v4733))) = v4727
	v4739 = F_open(m, int32(_a_F_InitPostgres_72), v4727, v4733)
	mBase = m.M
	if v4739 != int32(-1) {
		goto L993
	} else {
		goto L994
	}
L961:
	;
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v4656 != 0 {
		goto L960
	} else {
		goto L964
	}
L962:
	;
	goto L963
L963:
	;
	v4659 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L1
	} else {
		goto L965
	}
L964:
	;
	goto L963
L965:
	;
	if v4659 != 0 {
		goto L966
	} else {
		goto L967
	}
L966:
	;
	v4663 = int32(_a_F_InitPostgres_55)
	v4665 = v4653 + int32(1)
	if v4665 == int32(0) {
		v4685 = v4663
		goto L970
	} else {
		goto L971
	}
L967:
	;
	goto L968
L968:
	;
	v4703 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v4703 == int32(0) {
		v6422 = v4535
		v6423 = v4536
		v6424 = v4537
		v6425 = v4538
		v6426 = v4539
		v6427 = v4540
		v6428 = v4541
		v6433 = v4546
		v6434 = v4547
		v6435 = v4548
		v6436 = v4549
		v6440 = v4553
		v6442 = v4555
		v6445 = v4558
		v6446 = v4559
		v6447 = v4560
		v6449 = v4562
		v6453 = v4566
		v6454 = v4567
		v6455 = v4568
		v6456 = v4569
		v6457 = v4570
		goto L959
	} else {
		goto L981
	}
L969:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+692)) = v4685 + base.B2i32(v4687 == int32(0))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+688)) = v4583
	F_errmsg(m, int32(_a_F_InitPostgres_84), v4541+int32(688))
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L1
	} else {
		goto L979
	}
L970:
	;
	v4687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4685))))
	goto L969
L971:
	;
	v4669 = v4663
	v4670 = v4665
	goto L972
L972:
	;
	v4671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4669))))
	if v4671 == int32(0) {
		v4685 = v4669
		goto L970
	} else {
		goto L974
	}
L973:
	;
	v4685 = v4681
	goto L970
L974:
	;
	v4675 = v4669
	goto L975
L975:
	;
	v4679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4675)+1)))
	if v4679 != 0 {
		v4675 = v4675 + int32(1)
		goto L975
	} else {
		goto L977
	}
L976:
	;
	v4681 = v4675 + int32(2)
	v4683 = v4670 + int32(1)
	if v4683 != 0 {
		v4669 = v4681
		v4670 = v4683
		goto L972
	} else {
		goto L978
	}
L977:
	;
	goto L976
L978:
	;
	goto L973
L979:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2984), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	goto L968
L981:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	if v4706 == int32(1) {
		goto L984
	} else {
		goto L985
	}
L982:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L983:
	;
	goto L982
L984:
	;
	if v4703 == int32(0) {
		goto L983
	} else {
		goto L987
	}
L985:
	;
	goto L986
L986:
	;
	if v4703 == int32(0) {
		goto L983
	} else {
		goto L991
	}
L987:
	;
	v4712 = v4703
	goto L988
L988:
	;
	v4713 = *(*int32)(unsafe.Add(mBase, uint32(v4712)+28))
	v4714 = *(*int32)(unsafe.Add(mBase, uint32(v4712)+20))
	F_emscripten_builtin_free(m, v4714)
	mBase = m.M
	F_emscripten_builtin_free(m, v4712)
	mBase = m.M
	if v4713 != 0 {
		v4712 = v4713
		goto L988
	} else {
		goto L990
	}
L989:
	;
	goto L983
L990:
	;
	goto L989
L991:
	;
	F_freeaddrinfo(m, v4703)
	mBase = m.M
	goto L983
L992:
	;
	if v4772 == int32(0) {
		goto L1005
	} else {
		goto L1006
	}
L993:
	;
	goto L997
L994:
	;
	v4772 = v4727
	goto L995
L995:
	;
	m.G0 = v4733 + int32(16)
	goto L992
L996:
	;
	v4767 = F_close(m, v4739)
	mBase = m.M
	v4772 = v4765
	goto L995
L997:
	;
	v4745 = v4549
	v4746 = v4726
	goto L998
L998:
	;
	v4751 = F_read(m, v4739, v4745, v4746)
	mBase = m.M
	if v4751 <= int32(0) {
		goto L1000
	} else {
		goto L1001
	}
L999:
	;
	v4765 = int32(1)
	goto L996
L1000:
	;
	v4755 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v4755 == int32(27) {
		goto L998
	} else {
		goto L1003
	}
L1001:
	;
	goto L1002
L1002:
	;
	v4760 = v4746 - v4751
	if v4760 != 0 {
		v4745 = v4745 + v4751
		v4746 = v4760
		goto L998
	} else {
		goto L1004
	}
L1003:
	;
	v4765 = int32(0)
	goto L996
L1004:
	;
	goto L999
L1005:
	;
	v4781 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1006:
	;
	goto L1007
L1007:
	;
	v4809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+2756)))
	*(*uint8)(unsafe.Add(mBase, uint32(v4541)+2753)) = uint8(v4809)
	v4811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	if base.Ui32(int32(1021)) <= base.Ui32(v4811) {
		goto L1025
	} else {
		goto L1026
	}
L1008:
	;
	if v4781 != 0 {
		goto L1009
	} else {
		goto L1010
	}
L1009:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_86), int32(0))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L1
	} else {
		goto L1012
	}
L1010:
	;
	goto L1011
L1011:
	;
	v4792 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v4793 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v4792 == int32(1) {
		goto L1016
	} else {
		goto L1017
	}
L1012:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2997), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v4791 = m.ExcPending
	if v4791 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1013:
	;
	goto L1011
L1014:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L1015:
	;
	goto L1014
L1016:
	;
	if v4793 == int32(0) {
		goto L1015
	} else {
		goto L1019
	}
L1017:
	;
	goto L1018
L1018:
	;
	if v4793 == int32(0) {
		goto L1015
	} else {
		goto L1023
	}
L1019:
	;
	v4799 = v4793
	goto L1020
L1020:
	;
	v4800 = *(*int32)(unsafe.Add(mBase, uint32(v4799)+28))
	v4801 = *(*int32)(unsafe.Add(mBase, uint32(v4799)+20))
	F_emscripten_builtin_free(m, v4801)
	mBase = m.M
	F_emscripten_builtin_free(m, v4799)
	mBase = m.M
	if v4800 != 0 {
		v4799 = v4800
		goto L1020
	} else {
		goto L1022
	}
L1021:
	;
	goto L1015
L1022:
	;
	goto L1021
L1023:
	;
	F_freeaddrinfo(m, v4793)
	mBase = m.M
	goto L1015
L1024:
	;
	if v4586 != 0 {
		goto L1032
	} else {
		goto L1033
	}
L1025:
	;
	v4816 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L1028
	}
L1026:
	;
	goto L1027
L1027:
	;
	v4834 = v4541 + int32(2752) + v4811
	*(*int32)(unsafe.Add(mBase, uint32(v4834)+2)) = int32(134217728)
	v4837 = int32(1542)
	*(*uint16)(unsafe.Add(mBase, uint32(v4834))) = uint16(v4837)
	v4839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	v4841 = v4839 + int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v4841)
	goto L1024
L1028:
	;
	if v4816 == int32(0) {
		goto L1024
	} else {
		goto L1029
	}
L1029:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+672)) = int64(17179869190)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_87), v4541+int32(672))
	mBase = m.M
	v4826 = m.ExcPending
	if v4826 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2833), int32(_a_F_InitPostgres_88))
	mBase = m.M
	v4831 = m.ExcPending
	if v4831 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	goto L1024
L1032:
	;
	v4845 = v4586
	goto L1034
L1033:
	;
	v4845 = int32(_a_F_InitPostgres_89)
	goto L1034
L1034:
	;
	v4846 = F_strlen(m, v4588)
	mBase = m.M
	v4847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	if int32(1025) <= v4846+v4847 {
		goto L1036
	} else {
		goto L1037
	}
L1035:
	;
	v4888 = F_strlen(m, v4845)
	mBase = m.M
	v4889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	if int32(1025) <= v4888+v4889 {
		goto L1047
	} else {
		goto L1048
	}
L1036:
	;
	v4853 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4854 = m.ExcPending
	if v4854 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1037:
	;
	goto L1038
L1038:
	;
	v4872 = v4541 + int32(2752) + v4847
	v4874 = v4846 + int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4872)+1)) = uint8(v4874)
	v4876 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4872))) = uint8(v4876)
	if v4846 != 0 {
		goto L1043
	} else {
		goto L1044
	}
L1039:
	;
	if v4853 == int32(0) {
		goto L1035
	} else {
		goto L1040
	}
L1040:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+660)) = v4846
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+656)) = int32(1)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_87), v4541+int32(656))
	mBase = m.M
	v4864 = m.ExcPending
	if v4864 != 0 {
		goto L1
	} else {
		goto L1041
	}
L1041:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2833), int32(_a_F_InitPostgres_88))
	mBase = m.M
	v4869 = m.ExcPending
	if v4869 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1042:
	;
	goto L1035
L1043:
	;
	base.MemoryCopy(m, v4872+int32(2), v4588, v4846)
	goto L1045
L1044:
	;
	goto L1045
L1045:
	;
	v4881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	v4884 = v4881 + v4874&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v4884)
	goto L1035
L1046:
	;
	v4930 = int32(16)
	v4931 = F_strlen(m, v4553)
	mBase = m.M
	v4932 = F_strlen(m, v4587)
	mBase = m.M
	v4935 = F_palloc(m, v4932+v4930)
	mBase = m.M
	v4936 = m.ExcPending
	if v4936 != 0 {
		goto L1
	} else {
		goto L1057
	}
L1047:
	;
	v4895 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4896 = m.ExcPending
	if v4896 != 0 {
		goto L1
	} else {
		goto L1050
	}
L1048:
	;
	goto L1049
L1049:
	;
	v4914 = v4541 + int32(2752) + v4889
	v4916 = v4888 + int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v4914)+1)) = uint8(v4916)
	v4918 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v4914))) = uint8(v4918)
	if v4888 != 0 {
		goto L1054
	} else {
		goto L1055
	}
L1050:
	;
	if v4895 == int32(0) {
		goto L1046
	} else {
		goto L1051
	}
L1051:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+644)) = v4888
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+640)) = int32(32)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_87), v4541+int32(640))
	mBase = m.M
	v4906 = m.ExcPending
	if v4906 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2833), int32(_a_F_InitPostgres_88))
	mBase = m.M
	v4911 = m.ExcPending
	if v4911 != 0 {
		goto L1
	} else {
		goto L1053
	}
L1053:
	;
	goto L1046
L1054:
	;
	base.MemoryCopy(m, v4914+int32(2), v4845, v4888)
	goto L1056
L1055:
	;
	goto L1056
L1056:
	;
	v4923 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	v4926 = v4923 + v4916&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v4926)
	goto L1046
L1057:
	;
	v4937 = F_strlen(m, v4587)
	mBase = m.M
	if v4937 != 0 {
		goto L1058
	} else {
		goto L1059
	}
L1058:
	;
	base.MemoryCopy(m, v4935, v4587, v4937)
	goto L1060
L1059:
	;
	goto L1060
L1060:
	;
	v4940 = v4931 + int32(15)
	v4942 = v4940 & int32(-16)
	if int32(16) <= v4940 {
		goto L1064
	} else {
		goto L1065
	}
L1061:
	;
	v5200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	v5201 = int32(8)
	v5205 = v5200<<(uint(v5201)%32) | int32(base.Ui32(v5200)>>(uint(v5201)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v5205)
	v5207 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	v5208 = *(*int32)(unsafe.Add(mBase, uint32(v5207)+4))
	v5211 = F_socket(m, v5208, int32(2), int32(0))
	mBase = m.M
	if v5211 == int32(-1) {
		goto L1107
	} else {
		goto L1108
	}
L1062:
	;
	v5182 = v4541 + int32(2752) + v5122
	v5183 = int32(2)
	v5184 = v4942 | v5183
	*(*uint8)(unsafe.Add(mBase, uint32(v5182)+1)) = uint8(v5184)
	*(*uint8)(unsafe.Add(mBase, uint32(v5182))) = uint8(v5183)
	if v4942 != 0 {
		goto L1104
	} else {
		goto L1105
	}
L1063:
	;
	v5147 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5148 = m.ExcPending
	if v5148 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1064:
	;
	v4953 = v4549
	v4954 = v4930
	v4956 = int32(0)
	goto L1067
L1065:
	;
	goto L1066
L1066:
	;
	F_pfree(m, v4935)
	mBase = m.M
	v5121 = m.ExcPending
	if v5121 != 0 {
		goto L1
	} else {
		goto L1081
	}
L1067:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+848)) = int32(0)
	v4989 = F_strlen(m, v4587)
	mBase = m.M
	v4990 = v4989 + v4935
	v4991 = *(*int64)(unsafe.Add(mBase, uint32(v4953)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v4990)+8)) = v4991
	v4993 = *(*int64)(unsafe.Add(mBase, uint32(v4953)))
	*(*int64)(unsafe.Add(mBase, uint32(v4990))) = v4993
	v4995 = F_strlen(m, v4587)
	mBase = m.M
	v5000 = v4541 + int32(1168) + v4956
	v5003 = F_pg_md5_binary(m, v4935, v4995+int32(16), v5000, v4541+int32(848))
	mBase = m.M
	v5004 = m.ExcPending
	if v5004 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1068:
	;
	goto L1066
L1069:
	;
	if v5003 == int32(0) {
		goto L1063
	} else {
		goto L1070
	}
L1070:
	;
	v5007 = F_strlen(m, v4553)
	mBase = m.M
	v5017 = v4956
	goto L1071
L1071:
	;
	if base.Ui32(v5017) < base.Ui32(v5007) {
		goto L1073
	} else {
		goto L1074
	}
L1072:
	;
	v5074 = int32(16)
	v5077 = v4956 + v5074
	if v5077 < v4942 {
		v4953 = v5000
		v4954 = v4954 + v5074
		v4956 = v5077
		goto L1067
	} else {
		goto L1080
	}
L1073:
	;
	v5052 = v4541 + int32(1168) + v5017
	v5053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5052))))
	v5055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5017+v4553))))
	v5056 = v5053 ^ v5055
	*(*uint8)(unsafe.Add(mBase, uint32(v5052))) = uint8(v5056)
	goto L1075
L1074:
	;
	goto L1075
L1075:
	;
	v5060 = v5017 | int32(1)
	if base.Ui32(v5060) < base.Ui32(v5007) {
		goto L1076
	} else {
		goto L1077
	}
L1076:
	;
	v5064 = v4541 + int32(1168) + v5060
	v5065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5064))))
	v5067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4553+v5060))))
	v5068 = v5065 ^ v5067
	*(*uint8)(unsafe.Add(mBase, uint32(v5064))) = uint8(v5068)
	goto L1078
L1077:
	;
	goto L1078
L1078:
	;
	v5072 = v5017 + int32(2)
	if v5072 != v4954 {
		v5017 = v5072
		goto L1071
	} else {
		goto L1079
	}
L1079:
	;
	goto L1072
L1080:
	;
	goto L1068
L1081:
	;
	v5122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	if v4942+v5122 < int32(1025) {
		goto L1062
	} else {
		goto L1082
	}
L1082:
	;
	v5128 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5129 = m.ExcPending
	if v5129 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1083:
	;
	if v5128 == int32(0) {
		goto L1061
	} else {
		goto L1084
	}
L1084:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+612)) = v4942
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+608)) = int32(2)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_87), v4541+int32(608))
	mBase = m.M
	v5139 = m.ExcPending
	if v5139 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1085:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(2833), int32(_a_F_InitPostgres_88))
	mBase = m.M
	v5144 = m.ExcPending
	if v5144 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	goto L1061
L1087:
	;
	if v5147 != 0 {
		goto L1088
	} else {
		goto L1089
	}
L1088:
	;
	v5149 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+848))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+624)) = v5149
	F_errmsg(m, int32(_a_F_InitPostgres_90), v4541+int32(624))
	mBase = m.M
	v5155 = m.ExcPending
	if v5155 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1089:
	;
	goto L1090
L1090:
	;
	F_pfree(m, v4935)
	mBase = m.M
	v5162 = m.ExcPending
	if v5162 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1091:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3035), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v5160 = m.ExcPending
	if v5160 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	goto L1090
L1093:
	;
	v5163 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v5164 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v5163 == int32(1) {
		goto L1096
	} else {
		goto L1097
	}
L1094:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L1095:
	;
	goto L1094
L1096:
	;
	if v5164 == int32(0) {
		goto L1095
	} else {
		goto L1099
	}
L1097:
	;
	goto L1098
L1098:
	;
	if v5164 == int32(0) {
		goto L1095
	} else {
		goto L1103
	}
L1099:
	;
	v5170 = v5164
	goto L1100
L1100:
	;
	v5171 = *(*int32)(unsafe.Add(mBase, uint32(v5170)+28))
	v5172 = *(*int32)(unsafe.Add(mBase, uint32(v5170)+20))
	F_emscripten_builtin_free(m, v5172)
	mBase = m.M
	F_emscripten_builtin_free(m, v5170)
	mBase = m.M
	if v5171 != 0 {
		v5170 = v5171
		goto L1100
	} else {
		goto L1102
	}
L1101:
	;
	goto L1095
L1102:
	;
	goto L1101
L1103:
	;
	F_freeaddrinfo(m, v5164)
	mBase = m.M
	goto L1095
L1104:
	;
	base.MemoryCopy(m, v5182+int32(2), v4541+int32(1168), v4942)
	goto L1106
L1105:
	;
	goto L1106
L1106:
	;
	v5193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)))
	v5196 = v5193 + v5184&int32(255)
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+2754)) = uint16(v5196)
	goto L1061
L1107:
	;
	v5216 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5217 = m.ExcPending
	if v5217 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1108:
	;
	goto L1109
L1109:
	;
	v5244 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1448)) = v5244
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1440)) = v5244
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+1432)) = v5244
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1456)) = int32(0)
	v5252 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	v5253 = *(*int32)(unsafe.Add(mBase, uint32(v5252)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v4541)+1432)) = uint16(v5253)
	v5256 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[62]))
	*(*int64)(unsafe.Add(mBase, uint32(v4566)+8)) = v5256
	v5259 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[63]))
	*(*int64)(unsafe.Add(mBase, uint32(v4566))) = v5259
	if v5253&int32(_a_F_InitPostgres_91) == int32(10) {
		goto L1126
	} else {
		goto L1127
	}
L1110:
	;
	if v5216 != 0 {
		goto L1111
	} else {
		goto L1112
	}
L1111:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_92), int32(0))
	mBase = m.M
	v5221 = m.ExcPending
	if v5221 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1112:
	;
	goto L1113
L1113:
	;
	v5227 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v5228 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v5227 == int32(1) {
		goto L1118
	} else {
		goto L1119
	}
L1114:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3061), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v5226 = m.ExcPending
	if v5226 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	goto L1113
L1116:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L1117:
	;
	goto L1116
L1118:
	;
	if v5228 == int32(0) {
		goto L1117
	} else {
		goto L1121
	}
L1119:
	;
	goto L1120
L1120:
	;
	if v5228 == int32(0) {
		goto L1117
	} else {
		goto L1125
	}
L1121:
	;
	v5234 = v5228
	goto L1122
L1122:
	;
	v5235 = *(*int32)(unsafe.Add(mBase, uint32(v5234)+28))
	v5236 = *(*int32)(unsafe.Add(mBase, uint32(v5234)+20))
	F_emscripten_builtin_free(m, v5236)
	mBase = m.M
	F_emscripten_builtin_free(m, v5234)
	mBase = m.M
	if v5235 != 0 {
		v5234 = v5235
		goto L1122
	} else {
		goto L1124
	}
L1123:
	;
	goto L1117
L1124:
	;
	goto L1123
L1125:
	;
	F_freeaddrinfo(m, v5228)
	mBase = m.M
	goto L1117
L1126:
	;
	v5267 = int32(28)
	goto L1128
L1127:
	;
	v5267 = int32(16)
	goto L1128
L1128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1708)) = v5267
	v5271 = F_bind(m, v5211, v4541+int32(1432), v5267)
	mBase = m.M
	if v5271 != 0 {
		goto L1129
	} else {
		goto L1130
	}
L1129:
	;
	v5274 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5275 = m.ExcPending
	if v5275 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1130:
	;
	goto L1131
L1131:
	;
	v5305 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	v5306 = *(*int32)(unsafe.Add(mBase, uint32(v5305)+20))
	v5307 = *(*int32)(unsafe.Add(mBase, uint32(v5305)+16))
	v5308 = F_sendto(m, v5211, v4541+int32(2752), v5200, v5306, v5307)
	mBase = m.M
	if v5308 < int32(0) {
		goto L1148
	} else {
		goto L1149
	}
L1132:
	;
	if v5274 != 0 {
		goto L1133
	} else {
		goto L1134
	}
L1133:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_93), int32(0))
	mBase = m.M
	v5279 = m.ExcPending
	if v5279 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1134:
	;
	goto L1135
L1135:
	;
	v5285 = F_close(m, v5211)
	mBase = m.M
	v5286 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v5287 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v5286 == int32(1) {
		goto L1140
	} else {
		goto L1141
	}
L1136:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3077), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v5284 = m.ExcPending
	if v5284 != 0 {
		goto L1
	} else {
		goto L1137
	}
L1137:
	;
	goto L1135
L1138:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L1139:
	;
	goto L1138
L1140:
	;
	if v5287 == int32(0) {
		goto L1139
	} else {
		goto L1143
	}
L1141:
	;
	goto L1142
L1142:
	;
	if v5287 == int32(0) {
		goto L1139
	} else {
		goto L1147
	}
L1143:
	;
	v5293 = v5287
	goto L1144
L1144:
	;
	v5294 = *(*int32)(unsafe.Add(mBase, uint32(v5293)+28))
	v5295 = *(*int32)(unsafe.Add(mBase, uint32(v5293)+20))
	F_emscripten_builtin_free(m, v5295)
	mBase = m.M
	F_emscripten_builtin_free(m, v5293)
	mBase = m.M
	if v5294 != 0 {
		v5293 = v5294
		goto L1144
	} else {
		goto L1146
	}
L1145:
	;
	goto L1139
L1146:
	;
	goto L1145
L1147:
	;
	F_freeaddrinfo(m, v5287)
	mBase = m.M
	goto L1139
L1148:
	;
	v5313 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1149:
	;
	goto L1150
L1150:
	;
	v5342 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v5343 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v5342 == int32(1) {
		goto L1169
	} else {
		goto L1170
	}
L1151:
	;
	if v5313 != 0 {
		goto L1152
	} else {
		goto L1153
	}
L1152:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_94), int32(0))
	mBase = m.M
	v5318 = m.ExcPending
	if v5318 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1153:
	;
	goto L1154
L1154:
	;
	v5324 = F_close(m, v5211)
	mBase = m.M
	v5325 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1572))
	v5326 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+704))
	if v5325 == int32(1) {
		goto L1159
	} else {
		goto L1160
	}
L1155:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3087), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v5323 = m.ExcPending
	if v5323 != 0 {
		goto L1
	} else {
		goto L1156
	}
L1156:
	;
	goto L1154
L1157:
	;
	v6422 = v4535
	v6423 = v4536
	v6424 = v4537
	v6425 = v4538
	v6426 = v4539
	v6427 = v4540
	v6428 = v4541
	v6433 = v4546
	v6434 = v4547
	v6435 = v4548
	v6436 = v4549
	v6440 = v4553
	v6442 = v4555
	v6445 = v4558
	v6446 = v4559
	v6447 = v4560
	v6449 = v4562
	v6453 = v4566
	v6454 = v4567
	v6455 = v4568
	v6456 = v4569
	v6457 = v4570
	goto L959
L1158:
	;
	goto L1157
L1159:
	;
	if v5326 == int32(0) {
		goto L1158
	} else {
		goto L1162
	}
L1160:
	;
	goto L1161
L1161:
	;
	if v5326 == int32(0) {
		goto L1158
	} else {
		goto L1166
	}
L1162:
	;
	v5332 = v5326
	goto L1163
L1163:
	;
	v5333 = *(*int32)(unsafe.Add(mBase, uint32(v5332)+28))
	v5334 = *(*int32)(unsafe.Add(mBase, uint32(v5332)+20))
	F_emscripten_builtin_free(m, v5334)
	mBase = m.M
	F_emscripten_builtin_free(m, v5332)
	mBase = m.M
	if v5333 != 0 {
		v5332 = v5333
		goto L1163
	} else {
		goto L1165
	}
L1164:
	;
	goto L1158
L1165:
	;
	goto L1164
L1166:
	;
	F_freeaddrinfo(m, v5326)
	mBase = m.M
	goto L1158
L1167:
	;
	F_gettimeofday(m, v4541+int32(1136))
	mBase = m.M
	v5362 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+1136))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1704)) = int32(0)
	F_gettimeofday(m, v4541+int32(816))
	mBase = m.M
	v5369 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4541)+1144)))
	v5370 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4541)+824)))
	v5373 = v5362 + int64(3)
	v5374 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+816))
	v5378 = v5369 - v5370 + (v5373-v5374)*int64(1000000)
	if v5378 <= int64(0) {
		goto L1179
	} else {
		goto L1180
	}
L1168:
	;
	goto L1167
L1169:
	;
	if v5343 == int32(0) {
		goto L1168
	} else {
		goto L1172
	}
L1170:
	;
	goto L1171
L1171:
	;
	if v5343 == int32(0) {
		goto L1168
	} else {
		goto L1176
	}
L1172:
	;
	v5349 = v5343
	goto L1173
L1173:
	;
	v5350 = *(*int32)(unsafe.Add(mBase, uint32(v5349)+28))
	v5351 = *(*int32)(unsafe.Add(mBase, uint32(v5349)+20))
	F_emscripten_builtin_free(m, v5351)
	mBase = m.M
	F_emscripten_builtin_free(m, v5349)
	mBase = m.M
	if v5350 != 0 {
		v5349 = v5350
		goto L1173
	} else {
		goto L1175
	}
L1174:
	;
	goto L1168
L1175:
	;
	goto L1174
L1176:
	;
	F_freeaddrinfo(m, v5343)
	mBase = m.M
	goto L1168
L1177:
	;
	v6421 = F_close(m, v5211)
	mBase = m.M
	v6422 = v6380
	v6423 = v6381
	v6424 = v6382
	v6425 = v6383
	v6426 = v6384
	v6427 = v6385
	v6428 = v6386
	v6433 = v6391
	v6434 = v6392
	v6435 = v6393
	v6436 = v6394
	v6440 = v6398
	v6442 = v6400
	v6445 = v6403
	v6446 = v6404
	v6447 = v6405
	v6449 = v6407
	v6453 = v6411
	v6454 = v6412
	v6455 = v6413
	v6456 = v6414
	v6457 = v6415
	goto L959
L1178:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), v6376, int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6379 = m.ExcPending
	if v6379 != 0 {
		goto L1
	} else {
		goto L1333
	}
L1179:
	;
	v6324 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6325 = m.ExcPending
	if v6325 != 0 {
		goto L1
	} else {
		goto L1330
	}
L1180:
	;
	v5381 = int32(1)
	v5382 = v5211 + v5381
	v5391 = v4541 + int32(880) + int32(base.Ui32(v5211)>>(uint(int32(3))%32))&int32(536870908)
	v5392 = int32(8)
	v5438 = v5378
	goto L1182
L1181:
	;
	v6278 = F_close(m, v5211)
	mBase = m.M
	F_pfree(m, v4553)
	mBase = m.M
	v6280 = m.ExcPending
	if v6280 != 0 {
		goto L1
	} else {
		goto L1329
	}
L1182:
	;
	v5442 = int64(1000000)
	v5443 = base.I64_div_u_s(v5438, v5442)
	*(*int64)(unsafe.Add(mBase, uint32(v4541)+848)) = v5443
	v5447 = v5438 - v5443*v5442
	*(*uint32)(unsafe.Add(mBase, uint32(v4541)+856)) = uint32(v5447)
	v5450 = v4541 + int32(880)
	base.MemoryFill(m, v5450, int32(0), int32(128))
	v5454 = *(*int32)(unsafe.Add(mBase, uint32(v5391)))
	*(*int32)(unsafe.Add(mBase, uint32(v5391))) = v5454 | v5381<<(uint(v5211)%32)
	v5458 = v4541 + int32(848)
	if v5458 != 0 {
		goto L1186
	} else {
		goto L1187
	}
L1183:
	;
	v6271 = F_close(m, v5211)
	mBase = m.M
	v6272 = *(*int32)(unsafe.Add(mBase, uint32(v4546)+364))
	F_set_authn_id(m, v4546, v6272)
	mBase = m.M
	v6274 = m.ExcPending
	if v6274 != 0 {
		goto L1
	} else {
		goto L1327
	}
L1184:
	;
	goto L1183
L1185:
	;
	if base.Ui32(int32(-4095)) <= base.Ui32(v5946) {
		goto L1245
	} else {
		goto L1246
	}
L1186:
	;
	v5459 = int32(-28)
	v5460 = *(*int64)(unsafe.Add(mBase, uint32(v5458)))
	if v5460 < int64(0) {
		v5946 = v5459
		goto L1185
	} else {
		goto L1189
	}
L1187:
	;
	goto L1188
L1188:
	;
	v5467 = int32(0)
	if v5382 < v5467 {
		v5938 = int32(-28)
		goto L1191
	} else {
		goto L1192
	}
L1189:
	;
	v5463 = *(*int32)(unsafe.Add(mBase, uint32(v5458)+8))
	if v5463 < int32(0) {
		v5946 = v5459
		goto L1185
	} else {
		goto L1190
	}
L1190:
	;
	goto L1188
L1191:
	;
	v5946 = v5938
	goto L1185
L1192:
	;
	if v5458 != 0 {
		goto L1193
	} else {
		goto L1194
	}
L1193:
	;
	v5471 = *(*int32)(unsafe.Add(mBase, uint32(v5458)))
	v5472 = int32(1000)
	v5474 = *(*int32)(unsafe.Add(mBase, uint32(v5458)+8))
	v5476 = base.I32_div_s(v5474, v5472)
	v5479 = v5471*v5472 + v5476
	goto L1195
L1194:
	;
	v5479 = int32(-1)
	goto L1195
L1195:
	;
	v5480 = int32(8)
	v5482 = int32(0)
	if v5382 == v5482 {
		v5500 = v5482
		goto L1198
	} else {
		goto L1199
	}
L1196:
	;
	if v5382 != 0 {
		goto L1206
	} else {
		goto L1207
	}
L1197:
	;
	goto L1196
L1198:
	;
	v5501 = F_emscripten_builtin_malloc(m, v5500)
	mBase = m.M
	if v5501 == int32(0) {
		goto L1197
	} else {
		goto L1204
	}
L1199:
	;
	v5488 = base.I64_extend_i32_u(v5382) * base.I64_extend_i32_u(v5480)
	v5489 = base.I32_wrap_i64(v5488)
	if base.Ui32(v5382|v5480) < base.Ui32(int32(_a_F_InitPostgres_95)) {
		v5500 = v5489
		goto L1198
	} else {
		goto L1200
	}
L1200:
	;
	if base.I32_wrap_i64(int64(base.Ui64(v5488)>>(uint(int64(32))%64))) != 0 {
		goto L1201
	} else {
		goto L1202
	}
L1201:
	;
	v5497 = int32(-1)
	goto L1203
L1202:
	;
	v5497 = v5489
	goto L1203
L1203:
	;
	v5500 = v5497
	goto L1198
L1204:
	;
	v5506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5501-int32(4)))))
	if v5506&int32(3) == int32(0) {
		goto L1197
	} else {
		goto L1205
	}
L1205:
	;
	F___memset(m, v5501, int32(0), v5500)
	mBase = m.M
	goto L1197
L1206:
	;
	v5521 = int32(0)
	v5530 = v5467
	goto L1209
L1207:
	;
	v5603 = v5467
	goto L1208
L1208:
	;
	v5628 = m.Env.X__syscall_poll(m, v5501, v5603, v5479)
	mBase = m.M
	if v5628 < int32(0) {
		v5864 = v5628
		goto L1218
	} else {
		goto L1219
	}
L1209:
	;
	if v5450 == int32(0) {
		goto L1211
	} else {
		goto L1212
	}
L1210:
	;
	v5603 = v5583
	goto L1208
L1211:
	;
	v5578 = v5501 + v5530<<(uint(int32(3))%32)
	v5579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5578)+4)))
	if v5579 != 0 {
		goto L1214
	} else {
		goto L1215
	}
L1212:
	;
	v5562 = *(*int32)(unsafe.Add(mBase, uint32(v5450+int32(base.Ui32(v5521)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v5562)>>(uint(v5521)%32))&int32(1) == int32(0) {
		goto L1211
	} else {
		goto L1213
	}
L1213:
	;
	v5570 = v5501 + v5530<<(uint(int32(3))%32)
	v5571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5570)+4)))
	v5573 = v5571 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5570)+4)) = uint16(v5573)
	goto L1211
L1214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5578))) = v5521
	v5583 = v5530 + int32(1)
	goto L1216
L1215:
	;
	v5583 = v5530
	goto L1216
L1216:
	;
	v5585 = v5521 + int32(1)
	if v5585 != v5382 {
		v5521 = v5585
		v5530 = v5583
		goto L1209
	} else {
		goto L1217
	}
L1217:
	;
	goto L1210
L1218:
	;
	F_emscripten_builtin_free(m, v5501)
	mBase = m.M
	v5938 = v5864
	goto L1191
L1219:
	;
	if v5450 != 0 {
		goto L1220
	} else {
		goto L1221
	}
L1220:
	;
	v5639 = v5450
	v5653 = int32(32)
	goto L1223
L1221:
	;
	goto L1222
L1222:
	;
	v5720 = int32(0)
	if base.B2i32(v5628 == v5720)|base.B2i32(v5603 <= v5720) == v5720 {
		goto L1226
	} else {
		goto L1227
	}
L1223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5639))) = int32(0)
	v5678 = v5653 - int32(1)
	if v5678 != 0 {
		v5639 = v5639 + int32(4)
		v5653 = v5678
		goto L1223
	} else {
		goto L1225
	}
L1224:
	;
	goto L1222
L1225:
	;
	goto L1224
L1226:
	;
	v5735 = v5720
	goto L1230
L1227:
	;
	goto L1228
L1228:
	;
	v5864 = int32(0)
	goto L1218
L1229:
	;
	v5864 = int32(-8)
	goto L1218
L1230:
	;
	v5772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5501+v5735<<(uint(int32(3))%32))+6)))
	if v5772&int32(32) != 0 {
		goto L1229
	} else {
		goto L1232
	}
L1231:
	;
	v5778 = int32(0)
	v5787 = v5778
	v5789 = v5778
	goto L1234
L1232:
	;
	v5776 = v5735 + int32(1)
	if v5776 != v5603 {
		v5735 = v5776
		goto L1230
	} else {
		goto L1233
	}
L1233:
	;
	goto L1231
L1234:
	;
	v5823 = v5501 + v5787<<(uint(int32(3))%32)
	v5824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5823)+6)))
	if v5824 != 0 {
		goto L1236
	} else {
		goto L1237
	}
L1235:
	;
	v5864 = v5847
	goto L1218
L1236:
	;
	v5825 = *(*int32)(unsafe.Add(mBase, uint32(v5823)))
	v5826 = int32(0)
	if base.B2i32(v5450 == v5826)|base.B2i32(v5824&int32(25) == v5826) != 0 {
		goto L1239
	} else {
		goto L1240
	}
L1237:
	;
	v5847 = v5789
	goto L1238
L1238:
	;
	v5851 = v5787 + int32(1)
	if v5851 != v5603 {
		v5787 = v5851
		v5789 = v5847
		goto L1234
	} else {
		goto L1242
	}
L1239:
	;
	v5846 = v5789
	goto L1241
L1240:
	;
	v5837 = v5450 + int32(base.Ui32(v5825)>>(uint(int32(3))%32))&int32(536870908)
	v5838 = *(*int32)(unsafe.Add(mBase, uint32(v5837)))
	v5839 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v5837))) = v5838 | v5839<<(uint(v5825)%32)
	v5846 = v5789 + v5839
	goto L1241
L1241:
	;
	v5847 = v5846
	goto L1238
L1242:
	;
	goto L1235
L1243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1704)) = int32(0)
	F_gettimeofday(m, v4541+int32(816))
	mBase = m.M
	v6262 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4541)+824)))
	v6264 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+816))
	v6268 = v5369 - v6262 + (v5373-v6264)*int64(1000000)
	if int64(0) < v6268 {
		v5438 = v6268
		goto L1182
	} else {
		goto L1326
	}
L1244:
	;
	if v5987 < int32(0) {
		goto L1248
	} else {
		goto L1249
	}
L1245:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44])) = int32(0) - v5946
	v5987 = int32(-1)
	goto L1247
L1246:
	;
	v5987 = v5946
	goto L1247
L1247:
	;
	goto L1244
L1248:
	;
	v5991 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	if v5991 == int32(27) {
		goto L1243
	} else {
		goto L1251
	}
L1249:
	;
	goto L1250
L1250:
	;
	if v5987 == int32(0) {
		goto L1255
	} else {
		goto L1256
	}
L1251:
	;
	v5996 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v5997 = m.ExcPending
	if v5997 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	if v5996 == int32(0) {
		v6380 = v4535
		v6381 = v4536
		v6382 = v4537
		v6383 = v4538
		v6384 = v4539
		v6385 = v4540
		v6386 = v4541
		v6391 = v4546
		v6392 = v4547
		v6393 = v4548
		v6394 = v4549
		v6398 = v4553
		v6400 = v4555
		v6403 = v4558
		v6404 = v4559
		v6405 = v4560
		v6407 = v4562
		v6411 = v4566
		v6412 = v4567
		v6413 = v4568
		v6414 = v4569
		v6415 = v4570
		goto L1177
	} else {
		goto L1253
	}
L1253:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_96), int32(0))
	mBase = m.M
	v6003 = m.ExcPending
	if v6003 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	v6376 = int32(3140)
	goto L1178
L1255:
	;
	v6009 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6010 = m.ExcPending
	if v6010 != 0 {
		goto L1
	} else {
		goto L1258
	}
L1256:
	;
	goto L1257
L1257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+1708)) = int32(28)
	v6025 = int32(0)
	v6030 = F_recvfrom(m, v5211, v4541+int32(1712), int32(1024), v6025, v4541+int32(736), v4541+int32(1708))
	mBase = m.M
	if v6030 < v6025 {
		goto L1261
	} else {
		goto L1262
	}
L1258:
	;
	if v6009 == int32(0) {
		v6380 = v4535
		v6381 = v4536
		v6382 = v4537
		v6383 = v4538
		v6384 = v4539
		v6385 = v4540
		v6386 = v4541
		v6391 = v4546
		v6392 = v4547
		v6393 = v4548
		v6394 = v4549
		v6398 = v4553
		v6400 = v4555
		v6403 = v4558
		v6404 = v4559
		v6405 = v4560
		v6407 = v4562
		v6411 = v4566
		v6412 = v4567
		v6413 = v4568
		v6414 = v4569
		v6415 = v4570
		goto L1177
	} else {
		goto L1259
	}
L1259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+480)) = v4583
	F_errmsg(m, int32(_a_F_InitPostgres_97), v4541+int32(480))
	mBase = m.M
	v6018 = m.ExcPending
	if v6018 != 0 {
		goto L1
	} else {
		goto L1260
	}
L1260:
	;
	v6376 = int32(3148)
	goto L1178
L1261:
	;
	v6035 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6036 = m.ExcPending
	if v6036 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1262:
	;
	goto L1263
L1263:
	;
	v6044 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+738)))
	if (v4648<<(uint(v5392)%32)|int32(base.Ui32(v4648&int32(_a_F_InitPostgres_98))>>(uint(v5392)%32)))&int32(_a_F_InitPostgres_91) != v6044 {
		goto L1267
	} else {
		goto L1268
	}
L1264:
	;
	if v6035 == int32(0) {
		v6380 = v4535
		v6381 = v4536
		v6382 = v4537
		v6383 = v4538
		v6384 = v4539
		v6385 = v4540
		v6386 = v4541
		v6391 = v4546
		v6392 = v4547
		v6393 = v4548
		v6394 = v4549
		v6398 = v4553
		v6400 = v4555
		v6403 = v4558
		v6404 = v4559
		v6405 = v4560
		v6407 = v4562
		v6411 = v4566
		v6412 = v4567
		v6413 = v4568
		v6414 = v4569
		v6415 = v4570
		goto L1177
	} else {
		goto L1265
	}
L1265:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_99), int32(0))
	mBase = m.M
	v6042 = m.ExcPending
	if v6042 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1266:
	;
	v6376 = int32(3170)
	goto L1178
L1267:
	;
	v6048 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6049 = m.ExcPending
	if v6049 != 0 {
		goto L1
	} else {
		goto L1270
	}
L1268:
	;
	goto L1269
L1269:
	;
	if base.Ui32(v6030) <= base.Ui32(int32(19)) {
		goto L1274
	} else {
		goto L1275
	}
L1270:
	;
	if v6048 == int32(0) {
		goto L1243
	} else {
		goto L1271
	}
L1271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+592)) = v4583
	v6053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+738)))
	v6054 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+596)) = (v6053<<(uint(v6054)%32) | int32(base.Ui32(v6053)>>(uint(v6054)%32))) & int32(_a_F_InitPostgres_91)
	F_errmsg(m, int32(_a_F_InitPostgres_100), v4541+int32(592))
	mBase = m.M
	v6066 = m.ExcPending
	if v6066 != 0 {
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3179), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6071 = m.ExcPending
	if v6071 != 0 {
		goto L1
	} else {
		goto L1273
	}
L1273:
	;
	goto L1243
L1274:
	;
	v6076 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6077 = m.ExcPending
	if v6077 != 0 {
		goto L1
	} else {
		goto L1277
	}
L1275:
	;
	goto L1276
L1276:
	;
	v6092 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+1714)))
	v6093 = int32(8)
	if (v6092<<(uint(v6093)%32)|int32(base.Ui32(v6092)>>(uint(v6093)%32)))&int32(_a_F_InitPostgres_91) != v6030 {
		goto L1281
	} else {
		goto L1282
	}
L1277:
	;
	if v6076 == int32(0) {
		goto L1243
	} else {
		goto L1278
	}
L1278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+500)) = v6030
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+496)) = v4583
	F_errmsg(m, int32(_a_F_InitPostgres_101), v4541+int32(496))
	mBase = m.M
	v6086 = m.ExcPending
	if v6086 != 0 {
		goto L1
	} else {
		goto L1279
	}
L1279:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3186), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6091 = m.ExcPending
	if v6091 != 0 {
		goto L1
	} else {
		goto L1280
	}
L1280:
	;
	goto L1243
L1281:
	;
	v6103 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6104 = m.ExcPending
	if v6104 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1282:
	;
	goto L1283
L1283:
	;
	v6128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+2753)))
	v6129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+1713)))
	if v6128 != v6129 {
		goto L1288
	} else {
		goto L1289
	}
L1284:
	;
	if v6103 == int32(0) {
		goto L1243
	} else {
		goto L1285
	}
L1285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+576)) = v4583
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+584)) = v6030
	v6109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4541)+1714)))
	v6110 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+580)) = (v6109<<(uint(v6110)%32) | int32(base.Ui32(v6109)>>(uint(v6110)%32))) & int32(_a_F_InitPostgres_91)
	F_errmsg(m, int32(_a_F_InitPostgres_102), v4541+int32(576))
	mBase = m.M
	v6122 = m.ExcPending
	if v6122 != 0 {
		goto L1
	} else {
		goto L1286
	}
L1286:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3194), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6127 = m.ExcPending
	if v6127 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1287:
	;
	goto L1243
L1288:
	;
	v6133 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6134 = m.ExcPending
	if v6134 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1289:
	;
	goto L1290
L1290:
	;
	v6152 = F_strlen(m, v4587)
	mBase = m.M
	v6154 = F_palloc(m, v6152+v6030)
	mBase = m.M
	v6155 = m.ExcPending
	if v6155 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1291:
	;
	if v6133 == int32(0) {
		goto L1243
	} else {
		goto L1292
	}
L1292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+560)) = v4583
	v6138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+1713)))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+564)) = v6138
	v6140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+2753)))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+568)) = v6140
	F_errmsg(m, int32(_a_F_InitPostgres_103), v4541+int32(560))
	mBase = m.M
	v6146 = m.ExcPending
	if v6146 != 0 {
		goto L1
	} else {
		goto L1293
	}
L1293:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3202), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6151 = m.ExcPending
	if v6151 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1294:
	;
	goto L1243
L1295:
	;
	v6156 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1712))
	*(*int32)(unsafe.Add(mBase, uint32(v6154))) = v6156
	v6158 = *(*int64)(unsafe.Add(mBase, uint32(v4549)))
	*(*int64)(unsafe.Add(mBase, uint32(v6154)+4)) = v6158
	v6160 = *(*int64)(unsafe.Add(mBase, uint32(v4549)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v6154)+12)) = v6160
	if v6030 == int32(20) {
		goto L1296
	} else {
		goto L1297
	}
L1296:
	;
	v6172 = F_strlen(m, v4587)
	mBase = m.M
	if v6172 != 0 {
		goto L1299
	} else {
		goto L1300
	}
L1297:
	;
	v6165 = v6030 - int32(20)
	if v6165 == int32(0) {
		goto L1296
	} else {
		goto L1298
	}
L1298:
	;
	base.MemoryCopy(m, v6154+int32(20), v4569, v6165)
	goto L1296
L1299:
	;
	base.MemoryCopy(m, v6154+v6030, v4587, v6172)
	goto L1301
L1300:
	;
	goto L1301
L1301:
	;
	v6175 = F_strlen(m, v4587)
	mBase = m.M
	v6181 = F_pg_md5_binary(m, v6154, v6175+v6030, v4541+int32(1168), v4541+int32(1704))
	mBase = m.M
	v6182 = m.ExcPending
	if v6182 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1302:
	;
	if v6181 == int32(0) {
		goto L1303
	} else {
		goto L1304
	}
L1303:
	;
	v6187 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6188 = m.ExcPending
	if v6188 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1304:
	;
	goto L1305
L1305:
	;
	F_pfree(m, v6154)
	mBase = m.M
	v6204 = m.ExcPending
	if v6204 != 0 {
		goto L1
	} else {
		goto L1313
	}
L1306:
	;
	if v6187 != 0 {
		goto L1307
	} else {
		goto L1308
	}
L1307:
	;
	v6189 = *(*int32)(unsafe.Add(mBase, uint32(v4541)+1704))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+544)) = v6189
	F_errmsg(m, int32(_a_F_InitPostgres_104), v4541+int32(544))
	mBase = m.M
	v6195 = m.ExcPending
	if v6195 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1308:
	;
	goto L1309
L1309:
	;
	F_pfree(m, v6154)
	mBase = m.M
	v6202 = m.ExcPending
	if v6202 != 0 {
		goto L1
	} else {
		goto L1312
	}
L1310:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3229), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6200 = m.ExcPending
	if v6200 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1311:
	;
	goto L1309
L1312:
	;
	goto L1243
L1313:
	;
	v6205 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+1716))
	v6206 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+1168))
	v6208 = *(*int64)(unsafe.Add(mBase, uint32(v4570)))
	v6209 = *(*int64)(unsafe.Add(mBase, uint32(v4541)+1176))
	if v6205^v6206|(v6208^v6209) != int64(0) {
		goto L1314
	} else {
		goto L1315
	}
L1314:
	;
	v6216 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6217 = m.ExcPending
	if v6217 != 0 {
		goto L1
	} else {
		goto L1317
	}
L1315:
	;
	goto L1316
L1316:
	;
	v6231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+1712)))
	switch v6231 - int32(2) {
	case 0:
		goto L1184
	case 1:
		goto L1181
	default:
		goto L1321
	}
L1317:
	;
	if v6216 == int32(0) {
		goto L1243
	} else {
		goto L1318
	}
L1318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+528)) = v4583
	F_errmsg(m, int32(_a_F_InitPostgres_105), v4541+int32(528))
	mBase = m.M
	v6225 = m.ExcPending
	if v6225 != 0 {
		goto L1
	} else {
		goto L1319
	}
L1319:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3239), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6230 = m.ExcPending
	if v6230 != 0 {
		goto L1
	} else {
		goto L1320
	}
L1320:
	;
	goto L1243
L1321:
	;
	v6236 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6237 = m.ExcPending
	if v6237 != 0 {
		goto L1
	} else {
		goto L1322
	}
L1322:
	;
	if v6236 == int32(0) {
		goto L1243
	} else {
		goto L1323
	}
L1323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+512)) = v4583
	v6241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4541)+1712)))
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+516)) = v6241
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+520)) = v4588
	F_errmsg(m, int32(_a_F_InitPostgres_106), v4541+int32(512))
	mBase = m.M
	v6248 = m.ExcPending
	if v6248 != 0 {
		goto L1
	} else {
		goto L1324
	}
L1324:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(3257), int32(_a_F_InitPostgres_85))
	mBase = m.M
	v6253 = m.ExcPending
	if v6253 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	goto L1243
L1326:
	;
	goto L1179
L1327:
	;
	F_pfree(m, v4553)
	mBase = m.M
	v6276 = m.ExcPending
	if v6276 != 0 {
		goto L1
	} else {
		goto L1328
	}
L1328:
	;
	v6603 = v4535
	v6604 = v4536
	v6605 = v4537
	v6606 = v4538
	v6607 = v4539
	v6608 = v4540
	v6609 = v4541
	v6614 = v4546
	v6615 = v4547
	v6616 = int32(0)
	v6627 = v4559
	v6636 = v4568
	goto L518
L1329:
	;
	v6603 = v4535
	v6604 = v4536
	v6605 = v4537
	v6606 = v4538
	v6607 = v4539
	v6608 = v4540
	v6609 = v4541
	v6614 = v4546
	v6615 = v4547
	v6616 = v4548
	v6627 = v4559
	v6636 = v4568
	goto L518
L1330:
	;
	if v6324 == int32(0) {
		v6380 = v4535
		v6381 = v4536
		v6382 = v4537
		v6383 = v4538
		v6384 = v4539
		v6385 = v4540
		v6386 = v4541
		v6391 = v4546
		v6392 = v4547
		v6393 = v4548
		v6394 = v4549
		v6398 = v4553
		v6400 = v4555
		v6403 = v4558
		v6404 = v4559
		v6405 = v4560
		v6407 = v4562
		v6411 = v4566
		v6412 = v4567
		v6413 = v4568
		v6414 = v4569
		v6415 = v4570
		goto L1177
	} else {
		goto L1331
	}
L1331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4541)+464)) = v4583
	F_errmsg(m, int32(_a_F_InitPostgres_97), v4541+int32(464))
	mBase = m.M
	v6333 = m.ExcPending
	if v6333 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1332:
	;
	v6376 = int32(3122)
	goto L1178
L1333:
	;
	v6380 = v4535
	v6381 = v4536
	v6382 = v4537
	v6383 = v4538
	v6384 = v4539
	v6385 = v4540
	v6386 = v4541
	v6391 = v4546
	v6392 = v4547
	v6393 = v4548
	v6394 = v4549
	v6398 = v4553
	v6400 = v4555
	v6403 = v4558
	v6404 = v4559
	v6405 = v4560
	v6407 = v4562
	v6411 = v4566
	v6412 = v4567
	v6413 = v4568
	v6414 = v4569
	v6415 = v4570
	goto L1177
L1334:
	;
	v6482 = *(*int32)(unsafe.Add(mBase, uint32(v6463)+396))
	if v6482 == int32(0) {
		v6499 = v6447
		goto L1340
	} else {
		goto L1341
	}
L1335:
	;
	v6467 = *(*int32)(unsafe.Add(mBase, uint32(v6464)+4))
	if v6467 < int32(2) {
		v6481 = v6449
		goto L1334
	} else {
		goto L1336
	}
L1336:
	;
	v6471 = v6449 + int32(4)
	v6473 = *(*int32)(unsafe.Add(mBase, uint32(v6464)+12))
	if base.Ui32(v6471) < base.Ui32(v6473+v6467<<(uint(int32(2))%32)) {
		goto L1337
	} else {
		goto L1338
	}
L1337:
	;
	v6478 = v6471
	goto L1339
L1338:
	;
	v6478 = int32(0)
	goto L1339
L1339:
	;
	v6481 = v6478
	goto L1334
L1340:
	;
	v6500 = *(*int32)(unsafe.Add(mBase, uint32(v6463)+388))
	if v6500 == int32(0) {
		v6517 = v6442
		goto L1346
	} else {
		goto L1347
	}
L1341:
	;
	v6485 = *(*int32)(unsafe.Add(mBase, uint32(v6482)+4))
	if v6485 < int32(2) {
		v6499 = v6447
		goto L1340
	} else {
		goto L1342
	}
L1342:
	;
	v6489 = v6447 + int32(4)
	v6491 = *(*int32)(unsafe.Add(mBase, uint32(v6482)+12))
	if base.Ui32(v6489) < base.Ui32(v6491+v6485<<(uint(int32(2))%32)) {
		goto L1343
	} else {
		goto L1344
	}
L1343:
	;
	v6496 = v6489
	goto L1345
L1344:
	;
	v6496 = int32(0)
	goto L1345
L1345:
	;
	v6499 = v6496
	goto L1340
L1346:
	;
	v6519 = v6454 + int32(1)
	v6520 = *(*int32)(unsafe.Add(mBase, uint32(v6445)+4))
	if v6519 < v6520 {
		v4535 = v6422
		v4536 = v6423
		v4537 = v6424
		v4538 = v6425
		v4539 = v6426
		v4540 = v6427
		v4541 = v6428
		v4546 = v6433
		v4547 = v6434
		v4548 = v6435
		v4549 = v6436
		v4553 = v6440
		v4555 = v6517
		v4558 = v6445
		v4559 = v6446
		v4560 = v6499
		v4562 = v6481
		v4566 = v6453
		v4567 = v6519
		v4568 = v6455
		v4569 = v6456
		v4570 = v6457
		goto L932
	} else {
		goto L1352
	}
L1347:
	;
	v6503 = *(*int32)(unsafe.Add(mBase, uint32(v6500)+4))
	if v6503 < int32(2) {
		v6517 = v6442
		goto L1346
	} else {
		goto L1348
	}
L1348:
	;
	v6507 = v6442 + int32(4)
	v6509 = *(*int32)(unsafe.Add(mBase, uint32(v6500)+12))
	if base.Ui32(v6507) < base.Ui32(v6509+v6503<<(uint(int32(2))%32)) {
		goto L1349
	} else {
		goto L1350
	}
L1349:
	;
	v6514 = v6507
	goto L1351
L1350:
	;
	v6514 = int32(0)
	goto L1351
L1351:
	;
	v6517 = v6514
	goto L1346
L1352:
	;
	goto L933
L1353:
	;
	v6603 = v6522
	v6604 = v6523
	v6605 = v6524
	v6606 = v6525
	v6607 = v6526
	v6608 = v6527
	v6609 = v6528
	v6614 = v6533
	v6615 = v6534
	v6616 = v6535
	v6627 = v6546
	v6636 = v6555
	goto L518
L1354:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(460), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v6580 = m.ExcPending
	if v6580 != 0 {
		goto L1
	} else {
		goto L1355
	}
L1355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1356:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v6587 = m.ExcPending
	if v6587 != 0 {
		goto L1
	} else {
		goto L1357
	}
L1357:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_107), int32(0))
	mBase = m.M
	v6591 = m.ExcPending
	if v6591 != 0 {
		goto L1
	} else {
		goto L1358
	}
L1358:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(405), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v6596 = m.ExcPending
	if v6596 != 0 {
		goto L1
	} else {
		goto L1359
	}
L1359:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1360:
	;
	v6603 = l0
	v6604 = l1
	v6605 = l2
	v6606 = l3
	v6607 = l4
	v6608 = l5
	v6609 = v1126
	v6614 = v1075
	v6615 = v44
	v6616 = v6600
	v6627 = v47
	v6636 = v7
	goto L518
L1361:
	;
	v6684 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[64]))
	if v6684 != 0 {
		goto L1369
	} else {
		goto L1370
	}
L1362:
	;
	v6652 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[65]))
	if v6652 != 0 {
		goto L1361
	} else {
		goto L1363
	}
L1363:
	;
	v6655 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6656 = m.ExcPending
	if v6656 != 0 {
		goto L1
	} else {
		goto L1364
	}
L1364:
	;
	if v6655 == int32(0) {
		goto L1361
	} else {
		goto L1365
	}
L1365:
	;
	v6659 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+364))
	v6660 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+380))
	v6661 = *(*int32)(unsafe.Add(mBase, uint32(v6660)+296))
	v6664 = *(*int32)(unsafe.Add(mBase, uint32(v6661<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[66])))
	goto L1366
L1366:
	;
	v6665 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+380))
	v6666 = *(*int64)(unsafe.Add(mBase, uint32(v6665)))
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+68)) = v6664
	*(*int64)(unsafe.Add(mBase, uint32(v6609)+72)) = v6666
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+64)) = v6659
	F_errmsg(m, int32(_a_F_InitPostgres_108), v6609-int32(-64))
	mBase = m.M
	v6674 = m.ExcPending
	if v6674 != 0 {
		goto L1
	} else {
		goto L1367
	}
L1367:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(660), int32(_a_F_InitPostgres_52))
	mBase = m.M
	v6679 = m.ExcPending
	if v6679 != 0 {
		goto L1
	} else {
		goto L1368
	}
L1368:
	;
	goto L1361
L1369:
	;
	m.T0[v6684].(func(*base.Module, int32, int32))(m, v6614, v6616)
	mBase = m.M
	v6686 = m.ExcPending
	if v6686 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1370:
	;
	goto L1371
L1371:
	;
	if v6616 == int32(0) {
		goto L1374
	} else {
		goto L1375
	}
L1372:
	;
	goto L1371
L1373:
	;
	F_disable_timeout(m, int32(3))
	mBase = m.M
	v6777 = m.ExcPending
	if v6777 != 0 {
		goto L1
	} else {
		goto L1409
	}
L1374:
	;
	v6690 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v6690 != 0 {
		goto L1377
	} else {
		goto L1378
	}
L1375:
	;
	goto L1376
L1376:
	;
	if v6616 != int32(-2) {
		goto L1388
	} else {
		goto L1389
	}
L1377:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6692 = m.ExcPending
	if v6692 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1378:
	;
	goto L1379
L1379:
	;
	v6694 = v6609 + int32(2752)
	F_pq_beginmessage(m, v6694, int32(82))
	mBase = m.M
	v6697 = m.ExcPending
	if v6697 != 0 {
		goto L1
	} else {
		goto L1381
	}
L1380:
	;
	goto L1379
L1381:
	;
	F_enlargeStringInfo(m, v6694, int32(4))
	mBase = m.M
	v6700 = m.ExcPending
	if v6700 != 0 {
		goto L1
	} else {
		goto L1382
	}
L1382:
	;
	v6701 = *(*int32)(unsafe.Add(mBase, uint32(v6609)+2756))
	v6702 = *(*int32)(unsafe.Add(mBase, uint32(v6609)+2752))
	*(*int32)(unsafe.Add(mBase, uint32(v6701+v6702))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+2756)) = v6701 + int32(4)
	F_pq_endmessage(m, v6694)
	mBase = m.M
	v6710 = m.ExcPending
	if v6710 != 0 {
		goto L1
	} else {
		goto L1383
	}
L1383:
	;
	v6712 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[43]))
	if v6712 != 0 {
		goto L1384
	} else {
		goto L1385
	}
L1384:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6714 = m.ExcPending
	if v6714 != 0 {
		goto L1
	} else {
		goto L1387
	}
L1385:
	;
	goto L1386
L1386:
	;
	m.G0 = v6609 + int32(3792)
	goto L1373
L1387:
	;
	goto L1386
L1388:
	;
	v6720 = *(*int32)(unsafe.Add(mBase, uint32(v6609)+700))
	v6721 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+380))
	v6722 = *(*int32)(unsafe.Add(mBase, uint32(v6721)+296))
	if base.Ui32(int32(15)) < base.Ui32(v6722) {
		goto L1392
	} else {
		goto L1393
	}
L1389:
	;
	goto L1390
L1390:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6774 = m.ExcPending
	if v6774 != 0 {
		goto L1
	} else {
		goto L1408
	}
L1391:
	;
	v6734 = *(*int64)(unsafe.Add(mBase, uint32(v6721)))
	v6735 = *(*int32)(unsafe.Add(mBase, uint32(v6721)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+56)) = v6735
	*(*int64)(unsafe.Add(mBase, uint32(v6609)+48)) = v6734
	v6741 = F_psprintf(m, int32(_a_F_InitPostgres_109), v6609+int32(48))
	mBase = m.M
	v6742 = m.ExcPending
	if v6742 != 0 {
		goto L1
	} else {
		goto L1395
	}
L1392:
	;
	v6731 = int32(514)
	v6733 = int32(_a_F_InitPostgres_110)
	goto L1391
L1393:
	;
	goto L1394
L1394:
	;
	v6728 = v6722 << (uint(int32(2)) % 32)
	v6729 = *(*int32)(unsafe.Add(mBase, uint32(v6728)+uint32(_c_F_InitPostgres[67])))
	v6730 = *(*int32)(unsafe.Add(mBase, uint32(v6728)+uint32(_c_F_InitPostgres[68])))
	v6731 = v6729
	v6733 = v6730
	goto L1391
L1395:
	;
	if v6720 != 0 {
		goto L1396
	} else {
		goto L1397
	}
L1396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+36)) = v6741
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+32)) = v6720
	v6748 = F_psprintf(m, int32(_a_F_InitPostgres_111), v6609+int32(32))
	mBase = m.M
	v6749 = m.ExcPending
	if v6749 != 0 {
		goto L1
	} else {
		goto L1399
	}
L1397:
	;
	v6750 = v6741
	goto L1398
L1398:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6754 = m.ExcPending
	if v6754 != 0 {
		goto L1
	} else {
		goto L1400
	}
L1399:
	;
	v6750 = v6748
	goto L1398
L1400:
	;
	F_errcode(m, v6731)
	mBase = m.M
	v6756 = m.ExcPending
	if v6756 != 0 {
		goto L1
	} else {
		goto L1401
	}
L1401:
	;
	v6757 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v6609)+16)) = v6757
	F_errmsg(m, v6733, v6609+int32(16))
	mBase = m.M
	v6762 = m.ExcPending
	if v6762 != 0 {
		goto L1
	} else {
		goto L1402
	}
L1402:
	;
	if v6750 != 0 {
		goto L1403
	} else {
		goto L1404
	}
L1403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6609))) = v6750
	F_errdetail_log(m, int32(_a_F_InitPostgres_112), v6609)
	mBase = m.M
	v6766 = m.ExcPending
	if v6766 != 0 {
		goto L1
	} else {
		goto L1406
	}
L1404:
	;
	goto L1405
L1405:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_51), int32(320), int32(_a_F_InitPostgres_113))
	mBase = m.M
	v6771 = m.ExcPending
	if v6771 != 0 {
		goto L1
	} else {
		goto L1407
	}
L1406:
	;
	goto L1405
L1407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1409:
	;
	v6782 = m.G0
	v6783 = int32(16)
	v6784 = v6782 - v6783
	m.G0 = v6784
	F_gettimeofday(m, v6784)
	mBase = m.M
	v6787 = *(*int64)(unsafe.Add(mBase, uint32(v6784)))
	v6788 = int64(*(*int32)(unsafe.Add(mBase, uint32(v6784)+8)))
	m.G0 = v6784 + v6783
	goto L1410
L1410:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[69])) = v6788 + v6787*int64(1000000) - int64(946684800000000)
	v6799 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[57])))
	if v6799&int32(4) != 0 {
		goto L1411
	} else {
		goto L1412
	}
L1411:
	;
	v6803 = v6615 + int32(448)
	F_initStringInfo(m, v6803)
	mBase = m.M
	v6805 = m.ExcPending
	if v6805 != 0 {
		goto L1
	} else {
		goto L1414
	}
L1412:
	;
	goto L1413
L1413:
	;
	v6859 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[34])) = uint8(v6859)
	F_InitializeSessionUserId(m, v6605, v6606, v6859)
	mBase = m.M
	v6863 = m.ExcPending
	if v6863 != 0 {
		goto L1
	} else {
		goto L1434
	}
L1414:
	;
	v6806 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v6615)+400)) = v6806
	v6811 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v6811 != 0 {
		goto L1415
	} else {
		goto L1416
	}
L1415:
	;
	v6812 = int32(_a_F_InitPostgres_114)
	goto L1417
L1416:
	;
	v6812 = int32(_a_F_InitPostgres_115)
	goto L1417
L1417:
	;
	F_appendStringInfo(m, v6803, v6812, v6615+int32(400))
	mBase = m.M
	v6816 = m.ExcPending
	if v6816 != 0 {
		goto L1
	} else {
		goto L1418
	}
L1418:
	;
	v6818 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v6818 == int32(0) {
		goto L1419
	} else {
		goto L1420
	}
L1419:
	;
	v6821 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v6615)+384)) = v6821
	F_appendStringInfo(m, v6803, int32(_a_F_InitPostgres_116), v6615+int32(384))
	mBase = m.M
	v6827 = m.ExcPending
	if v6827 != 0 {
		goto L1
	} else {
		goto L1422
	}
L1420:
	;
	goto L1421
L1421:
	;
	v6828 = *(*int32)(unsafe.Add(mBase, uint32(v6614)+376))
	if v6828 != 0 {
		goto L1423
	} else {
		goto L1424
	}
L1422:
	;
	goto L1421
L1423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6615)+368)) = v6828
	F_appendStringInfo(m, v6615+int32(448), int32(_a_F_InitPostgres_117), v6615+int32(368))
	mBase = m.M
	v6836 = m.ExcPending
	if v6836 != 0 {
		goto L1
	} else {
		goto L1426
	}
L1424:
	;
	goto L1425
L1425:
	;
	v6839 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v6840 = m.ExcPending
	if v6840 != 0 {
		goto L1
	} else {
		goto L1427
	}
L1426:
	;
	goto L1425
L1427:
	;
	if v6839 != 0 {
		goto L1428
	} else {
		goto L1429
	}
L1428:
	;
	v6841 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+448))
	*(*int32)(unsafe.Add(mBase, uint32(v6615)+352)) = v6841
	F_errmsg_internal(m, int32(_a_F_InitPostgres_112), v6615+int32(352))
	mBase = m.M
	v6847 = m.ExcPending
	if v6847 != 0 {
		goto L1
	} else {
		goto L1431
	}
L1429:
	;
	goto L1430
L1430:
	;
	v6853 = *(*int32)(unsafe.Add(mBase, uint32(v6615)+448))
	F_pfree(m, v6853)
	mBase = m.M
	v6855 = m.ExcPending
	if v6855 != 0 {
		goto L1
	} else {
		goto L1433
	}
L1431:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(309), int32(_a_F_InitPostgres_118))
	mBase = m.M
	v6852 = m.ExcPending
	if v6852 != 0 {
		goto L1
	} else {
		goto L1432
	}
L1432:
	;
	goto L1430
L1433:
	;
	goto L1413
L1434:
	;
	v6865 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[65]))
	if v6865 == int32(0) {
		v6875 = v6603
		v6876 = v6604
		v6879 = v6607
		v6880 = v6608
		v6887 = v6615
		v6899 = v6627
		v6908 = v6636
		goto L144
	} else {
		goto L1435
	}
L1435:
	;
	v6869 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[70]))
	v6872 = *(*int32)(unsafe.Add(mBase, uint32(v6869<<(uint(int32(2))%32))+uint32(_c_F_InitPostgres[66])))
	goto L1436
L1436:
	;
	F_InitializeSystemUser(m, v6865, v6872)
	mBase = m.M
	v6874 = m.ExcPending
	if v6874 != 0 {
		goto L1
	} else {
		goto L1437
	}
L1437:
	;
	v6875 = v6603
	v6876 = v6604
	v6879 = v6607
	v6880 = v6608
	v6887 = v6615
	v6899 = v6627
	v6908 = v6636
	goto L144
L1438:
	;
	v6918 = v6875
	v6919 = v6876
	v6921 = v6916
	v6922 = v6879
	v6923 = v6880
	v6930 = v6887
	v6942 = v6899
	v6951 = v6908
	goto L143
L1439:
	;
	v6961 = int32(_a_F_InitPostgres_119)
	v6963 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	v6964 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71])) = v6963 + v6964
	v6968 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[72]))
	v6969 = *(*int32)(unsafe.Add(mBase, uint32(v6968)))
	*(*int32)(unsafe.Add(mBase, uint32(v6968))) = v6969 + v6964
	v6973 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6968)+192)) = uint8(v6973)
	*(*uint8)(unsafe.Add(mBase, uint32(v6968)+200)) = uint8(v6973)
	*(*int32)(unsafe.Add(mBase, uint32(v6968))) = v6969 + int32(2)
	v6983 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[71])) = v6983 - v6964
	goto L1441
L1440:
	;
	goto L1441
L1441:
	;
	v6989 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[73])))
	if (v6989^int32(-1)|v6921)&int32(1) == int32(0) {
		goto L141
	} else {
		goto L1442
	}
L1442:
	;
	v6998 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v6999 = int32(1)
	if (base.B2i32(v6998 != v6999)|v6921)&v6999 != 0 {
		goto L1443
	} else {
		goto L1444
	}
L1443:
	;
	v7183 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v7183 != int32(1) {
		goto L1461
	} else {
		goto L1462
	}
L1444:
	;
	v7005 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[74]))
	v7007 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	v7008 = v7005 + v7007
	if v7008 <= int32(0) {
		goto L1443
	} else {
		goto L1445
	}
L1445:
	;
	v7012 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	v7015 = base.AtomicRmwXchg32(m, v7012, int32(0), int32(1))
	if v7015 != 0 {
		goto L1446
	} else {
		goto L1447
	}
L1446:
	;
	v7017 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	F_s_lock(m, v7017, int32(_a_F_InitPostgres_120), int32(789), int32(_a_F_InitPostgres_121))
	mBase = m.M
	v7022 = m.ExcPending
	if v7022 != 0 {
		goto L1
	} else {
		goto L1449
	}
L1447:
	;
	goto L1448
L1448:
	;
	v7024 = v6930 + int32(444)
	v7025 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7024))) = v7025
	v7028 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[77]))
	v7029 = *(*int32)(unsafe.Add(mBase, uint32(v7028)+24))
	if v7029 == v7025 {
		goto L1450
	} else {
		goto L1451
	}
L1449:
	;
	goto L1448
L1450:
	;
	v7124 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[76]))
	v7125 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7124))), uint32(v7125))
	v7128 = *(*int32)(unsafe.Add(mBase, uint32(v7024)))
	if v7128 == v7008 {
		goto L1443
	} else {
		goto L1457
	}
L1451:
	;
	v7033 = v7028 + int32(20)
	if v7029 == v7033 {
		goto L1450
	} else {
		goto L1452
	}
L1452:
	;
	v7043 = v7029
	v7068 = v6951
	goto L1453
L1453:
	;
	v7077 = v7068 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7024))) = v7077
	if v7008 == v7077 {
		goto L1450
	} else {
		goto L1455
	}
L1454:
	;
	goto L1450
L1455:
	;
	v7080 = *(*int32)(unsafe.Add(mBase, uint32(v7043)+4))
	if v7080 != v7033 {
		v7043 = v7080
		v7068 = v7077
		goto L1453
	} else {
		goto L1456
	}
L1456:
	;
	goto L1454
L1457:
	;
	v7130 = *(*int32)(unsafe.Add(mBase, uint32(v6930)+444))
	v7132 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[75]))
	if v7130 < v7132 {
		goto L140
	} else {
		goto L1458
	}
L1458:
	;
	v7135 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	v7137 = F_has_privs_of_role(m, v7135, int32(_a_F_InitPostgres_122))
	mBase = m.M
	v7138 = m.ExcPending
	if v7138 != 0 {
		goto L1
	} else {
		goto L1459
	}
L1459:
	;
	if v7137 == int32(0) {
		goto L139
	} else {
		goto L1460
	}
L1460:
	;
	goto L1443
L1461:
	;
	if v6942 == int32(0) {
		goto L136
	} else {
		goto L1475
	}
L1462:
	;
	v7187 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	v7188 = F_has_rolreplication(m, v7187)
	mBase = m.M
	v7189 = m.ExcPending
	if v7189 != 0 {
		goto L1
	} else {
		goto L1463
	}
L1463:
	;
	if v7188 == int32(0) {
		goto L138
	} else {
		goto L1464
	}
L1464:
	;
	v7193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[50])))
	if v7193 != int32(1) {
		goto L1461
	} else {
		goto L1465
	}
L1465:
	;
	v7197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[51])))
	if v7197&int32(1) != 0 {
		goto L1461
	} else {
		goto L1466
	}
L1466:
	;
	v7201 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v7201 != 0 {
		goto L1467
	} else {
		goto L1468
	}
L1467:
	;
	F_process_startup_options(m, v7201, v6921)
	mBase = m.M
	v7203 = m.ExcPending
	if v7203 != 0 {
		goto L1
	} else {
		goto L1470
	}
L1468:
	;
	goto L1469
L1469:
	;
	v7205 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[79]))
	if int32(0) < v7205 {
		goto L1471
	} else {
		goto L1472
	}
L1470:
	;
	goto L1469
L1471:
	;
	F_pg_usleep(m, v7205*int32(_a_F_InitPostgres_123))
	mBase = m.M
	goto L1473
L1472:
	;
	goto L1473
L1473:
	;
	F_InitializeClientEncoding(m)
	mBase = m.M
	v7212 = m.ExcPending
	if v7212 != 0 {
		goto L1
	} else {
		goto L1474
	}
L1474:
	;
	goto L133
L1475:
	;
	if v6918 != 0 {
		goto L1477
	} else {
		goto L1478
	}
L1476:
	;
	F_LockSharedObject(m, int32(1262), v7252, int32(3))
	mBase = m.M
	v7260 = m.ExcPending
	if v7260 != 0 {
		goto L1
	} else {
		goto L1492
	}
L1477:
	;
	v7216 = v6930 + int32(448)
	F_ScanKeyInit(m, v7216, int32(2), int32(3), int32(62), v6918)
	mBase = m.M
	v7221 = m.ExcPending
	if v7221 != 0 {
		goto L1
	} else {
		goto L1480
	}
L1478:
	;
	goto L1479
L1479:
	;
	if v6919 == int32(0) {
		goto L133
	} else {
		goto L1491
	}
L1480:
	;
	v7225 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v7226 = m.ExcPending
	if v7226 != 0 {
		goto L1
	} else {
		goto L1481
	}
L1481:
	;
	v7229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[80])))
	v7232 = F_systable_beginscan(m, v7225, int32(2671), v7229, int32(0), int32(1), v7216)
	mBase = m.M
	v7233 = m.ExcPending
	if v7233 != 0 {
		goto L1
	} else {
		goto L1482
	}
L1482:
	;
	v7234 = F_systable_getnext(m, v7232)
	mBase = m.M
	v7235 = m.ExcPending
	if v7235 != 0 {
		goto L1
	} else {
		goto L1483
	}
L1483:
	;
	if v7234 != 0 {
		goto L1484
	} else {
		goto L1485
	}
L1484:
	;
	v7236 = F_heap_copytuple(m, v7234)
	mBase = m.M
	v7237 = m.ExcPending
	if v7237 != 0 {
		goto L1
	} else {
		goto L1487
	}
L1485:
	;
	v7238 = int32(0)
	goto L1486
L1486:
	;
	F_systable_endscan(m, v7232)
	mBase = m.M
	v7240 = m.ExcPending
	if v7240 != 0 {
		goto L1
	} else {
		goto L1488
	}
L1487:
	;
	v7238 = v7236
	goto L1486
L1488:
	;
	F_relation_close(m, v7225, int32(1))
	mBase = m.M
	v7243 = m.ExcPending
	if v7243 != 0 {
		goto L1
	} else {
		goto L1489
	}
L1489:
	;
	if v7238 == int32(0) {
		goto L137
	} else {
		goto L1490
	}
L1490:
	;
	v7246 = *(*int32)(unsafe.Add(mBase, uint32(v7238)+16))
	v7247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7246)+22)))
	v7249 = *(*int32)(unsafe.Add(mBase, uint32(v7246+v7247)))
	v7252 = v7249
	goto L1476
L1491:
	;
	v7252 = v6919
	goto L1476
L1492:
	;
	v7262 = v6930 + int32(448)
	F_ScanKeyInit(m, v7262, int32(1), int32(3), int32(184), v7252)
	mBase = m.M
	v7267 = m.ExcPending
	if v7267 != 0 {
		goto L1
	} else {
		goto L1493
	}
L1493:
	;
	v7270 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v7271 = m.ExcPending
	if v7271 != 0 {
		goto L1
	} else {
		goto L1494
	}
L1494:
	;
	v7274 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[80])))
	v7277 = F_systable_beginscan(m, v7270, int32(2672), v7274, int32(0), int32(1), v7262)
	mBase = m.M
	v7278 = m.ExcPending
	if v7278 != 0 {
		goto L1
	} else {
		goto L1495
	}
L1495:
	;
	v7279 = F_systable_getnext(m, v7277)
	mBase = m.M
	v7280 = m.ExcPending
	if v7280 != 0 {
		goto L1
	} else {
		goto L1496
	}
L1496:
	;
	if v7279 != 0 {
		goto L1497
	} else {
		goto L1498
	}
L1497:
	;
	v7281 = F_heap_copytuple(m, v7279)
	mBase = m.M
	v7282 = m.ExcPending
	if v7282 != 0 {
		goto L1
	} else {
		goto L1500
	}
L1498:
	;
	v7283 = int32(0)
	goto L1499
L1499:
	;
	F_systable_endscan(m, v7277)
	mBase = m.M
	v7285 = m.ExcPending
	if v7285 != 0 {
		goto L1
	} else {
		goto L1501
	}
L1500:
	;
	v7283 = v7281
	goto L1499
L1501:
	;
	F_relation_close(m, v7270, int32(1))
	mBase = m.M
	v7288 = m.ExcPending
	if v7288 != 0 {
		goto L1
	} else {
		goto L1502
	}
L1502:
	;
	if v7283 != 0 {
		goto L1504
	} else {
		goto L1505
	}
L1503:
	;
	v7332 = v6930 + int32(448)
	v7334 = v7291 + int32(4)
	goto L1527
L1504:
	;
	v7289 = *(*int32)(unsafe.Add(mBase, uint32(v7283)+16))
	v7290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7289)+22)))
	v7291 = v7289 + v7290
	if v6918 == int32(0) {
		goto L1503
	} else {
		goto L1507
	}
L1505:
	;
	goto L1506
L1506:
	;
	if v6918 != 0 {
		goto L119
	} else {
		goto L1519
	}
L1507:
	;
	v7295 = v7291 + int32(4)
	if v7295|v6918 != 0 {
		goto L1509
	} else {
		goto L1510
	}
L1508:
	;
	if v7310 == int32(0) {
		goto L1503
	} else {
		goto L1518
	}
L1509:
	;
	v7301 = int32(-1)
	goto L1511
L1510:
	;
	v7301 = int32(0)
	goto L1511
L1511:
	;
	if v7295 != 0 {
		goto L1512
	} else {
		goto L1513
	}
L1512:
	;
	v7302 = int32(1)
	goto L1514
L1513:
	;
	v7302 = v7301
	goto L1514
L1514:
	;
	v7303 = int32(0)
	if base.B2i32(v7295 == v7303)|base.B2i32(v6918 == v7303) != 0 {
		goto L1515
	} else {
		goto L1516
	}
L1515:
	;
	v7310 = v7302
	goto L1517
L1516:
	;
	v7309 = F_strncmp(m, v7295, v6918, int32(64))
	mBase = m.M
	v7310 = v7309
	goto L1517
L1517:
	;
	goto L1508
L1518:
	;
	goto L119
L1519:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7316 = m.ExcPending
	if v7316 != 0 {
		goto L1
	} else {
		goto L1520
	}
L1520:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v7319 = m.ExcPending
	if v7319 != 0 {
		goto L1
	} else {
		goto L1521
	}
L1521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+240)) = v7252
	F_errmsg(m, int32(_a_F_InitPostgres_124), v6930+int32(240))
	mBase = m.M
	v7325 = m.ExcPending
	if v7325 != 0 {
		goto L1
	} else {
		goto L1522
	}
L1522:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1101), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7330 = m.ExcPending
	if v7330 != 0 {
		goto L1
	} else {
		goto L1523
	}
L1523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1524:
	;
	v7454 = *(*int32)(unsafe.Add(mBase, uint32(v7291)+80))
	goto L1555
L1525:
	;
	v7451 = F_strlen(m, v7440)
	mBase = m.M
	goto L1524
L1527:
	;
	goto L1528
L1528:
	;
	v7341 = int32(63)
	if (v7332^v7334)&int32(3) != 0 {
		goto L1532
	} else {
		goto L1533
	}
L1529:
	;
	v7444 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7441))) = uint8(v7444)
	goto L1525
L1530:
	;
	v7425 = v7420
	v7426 = v7421
	v7427 = v7422
	goto L1551
L1531:
	;
	if v7415 == int32(0) {
		v7440 = v7413
		v7441 = v7414
		goto L1529
	} else {
		goto L1550
	}
L1532:
	;
	v7413 = v7334
	v7414 = v7332
	v7415 = v7341
	goto L1531
L1533:
	;
	goto L1534
L1534:
	;
	v7345 = int32(0)
	if base.B2i32(v7334&int32(3) == v7345)|int32(0) == v7345 {
		goto L1536
	} else {
		goto L1537
	}
L1535:
	;
	if v7381 == int32(0) {
		v7440 = v7378
		v7441 = v7379
		goto L1529
	} else {
		goto L1544
	}
L1536:
	;
	v7357 = v7334
	v7358 = v7332
	v7359 = v7341
	goto L1539
L1537:
	;
	goto L1538
L1538:
	;
	v7378 = v7334
	v7379 = v7332
	v7380 = v7341
	v7381 = int32(1)
	goto L1535
L1539:
	;
	v7361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7357))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7358))) = uint8(v7361)
	if v7361 == int32(0) {
		v7420 = v7357
		v7421 = v7358
		v7422 = v7359
		goto L1530
	} else {
		goto L1541
	}
L1540:
	;
	v7378 = v7372
	v7379 = v7366
	v7380 = v7368
	v7381 = v7370
	goto L1535
L1541:
	;
	v7365 = int32(1)
	v7366 = v7358 + v7365
	v7368 = v7359 - v7365
	v7369 = int32(0)
	v7370 = base.B2i32(v7368 != v7369)
	v7372 = v7357 + v7365
	if v7372&int32(3) == v7369 {
		v7378 = v7372
		v7379 = v7366
		v7380 = v7368
		v7381 = v7370
		goto L1535
	} else {
		goto L1542
	}
L1542:
	;
	if v7368 != 0 {
		v7357 = v7372
		v7358 = v7366
		v7359 = v7368
		goto L1539
	} else {
		goto L1543
	}
L1543:
	;
	goto L1540
L1544:
	;
	v7384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7378))))
	if base.B2i32(v7384 == int32(0))|base.B2i32(base.Ui32(v7380) < base.Ui32(int32(4))) != 0 {
		v7413 = v7378
		v7414 = v7379
		v7415 = v7380
		goto L1531
	} else {
		goto L1545
	}
L1545:
	;
	v7391 = v7378
	v7392 = v7379
	v7393 = v7380
	goto L1546
L1546:
	;
	v7396 = *(*int32)(unsafe.Add(mBase, uint32(v7391)))
	v7399 = int32(-2139062144)
	if (int32(16843008)-v7396|v7396)&v7399 != v7399 {
		v7420 = v7391
		v7421 = v7392
		v7422 = v7393
		goto L1530
	} else {
		goto L1548
	}
L1547:
	;
	v7413 = v7407
	v7414 = v7405
	v7415 = v7409
	goto L1531
L1548:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7392))) = v7396
	v7404 = int32(4)
	v7405 = v7392 + v7404
	v7407 = v7391 + v7404
	v7409 = v7393 - v7404
	if base.Ui32(int32(3)) < base.Ui32(v7409) {
		v7391 = v7407
		v7392 = v7405
		v7393 = v7409
		goto L1546
	} else {
		goto L1549
	}
L1549:
	;
	goto L1547
L1550:
	;
	v7420 = v7413
	v7421 = v7414
	v7422 = v7415
	goto L1530
L1551:
	;
	v7429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7425))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7426))) = uint8(v7429)
	if v7429 == int32(0) {
		v7440 = v7425
		v7441 = v7426
		goto L1529
	} else {
		goto L1553
	}
L1552:
	;
	v7440 = v7436
	v7441 = v7434
	goto L1529
L1553:
	;
	v7433 = int32(1)
	v7434 = v7426 + v7433
	v7436 = v7425 + v7433
	v7438 = v7427 - v7433
	if v7438 != 0 {
		v7425 = v7436
		v7426 = v7434
		v7427 = v7438
		goto L1551
	} else {
		goto L1554
	}
L1554:
	;
	goto L1552
L1555:
	;
	if v7454 == int32(-2) {
		goto L135
	} else {
		goto L1556
	}
L1556:
	;
	v7458 = *(*int32)(unsafe.Add(mBase, uint32(v7291)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32])) = v7458
	v7461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7291)+79)))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[81])) = uint8(v7461)
	if v6923 == int32(0) {
		v7680 = v7252
		goto L134
	} else {
		goto L1557
	}
L1557:
	;
	if (v7332^v6923)&int32(3) != 0 {
		goto L1561
	} else {
		goto L1562
	}
L1558:
	;
	v7680 = v7252
	goto L134
L1559:
	;
	goto L1558
L1560:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7519))) = uint8(v7518)
	if v7518&int32(255) == int32(0) {
		goto L1559
	} else {
		goto L1575
	}
L1561:
	;
	v7470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7332))))
	v7517 = v7332
	v7518 = v7470
	v7519 = v6923
	goto L1560
L1562:
	;
	goto L1563
L1563:
	;
	if v7332&int32(3) != 0 {
		goto L1564
	} else {
		goto L1565
	}
L1564:
	;
	v7474 = v7332
	v7476 = v6923
	goto L1567
L1565:
	;
	v7488 = v7332
	v7490 = v6923
	goto L1566
L1566:
	;
	v7492 = *(*int32)(unsafe.Add(mBase, uint32(v7488)))
	v7495 = int32(-2139062144)
	if (int32(16843008)-v7492|v7492)&v7495 != v7495 {
		v7517 = v7488
		v7518 = v7492
		v7519 = v7490
		goto L1560
	} else {
		goto L1571
	}
L1567:
	;
	v7477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7474))))
	*(*uint8)(unsafe.Add(mBase, uint32(v7476))) = uint8(v7477)
	if v7477 == int32(0) {
		goto L1559
	} else {
		goto L1569
	}
L1568:
	;
	v7488 = v7484
	v7490 = v7482
	goto L1566
L1569:
	;
	v7481 = int32(1)
	v7482 = v7476 + v7481
	v7484 = v7474 + v7481
	if v7484&int32(3) != 0 {
		v7474 = v7484
		v7476 = v7482
		goto L1567
	} else {
		goto L1570
	}
L1570:
	;
	goto L1568
L1571:
	;
	v7500 = v7488
	v7501 = v7492
	v7502 = v7490
	goto L1572
L1572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7502))) = v7501
	v7504 = int32(4)
	v7505 = v7502 + v7504
	v7507 = v7500 + v7504
	v7509 = *(*int32)(unsafe.Add(mBase, uint32(v7500)+4))
	v7512 = int32(-2139062144)
	if (int32(16843008)-v7509|v7509)&v7512 == v7512 {
		v7500 = v7507
		v7501 = v7509
		v7502 = v7505
		goto L1572
	} else {
		goto L1574
	}
L1573:
	;
	v7517 = v7507
	v7518 = v7509
	v7519 = v7505
	goto L1560
L1574:
	;
	goto L1573
L1575:
	;
	v7526 = v7517
	v7528 = v7519
	goto L1576
L1576:
	;
	v7529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7526)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7528)+1)) = uint8(v7529)
	v7531 = int32(1)
	if v7529 != 0 {
		v7526 = v7526 + v7531
		v7528 = v7528 + v7531
		goto L1576
	} else {
		goto L1578
	}
L1577:
	;
	goto L1559
L1578:
	;
	goto L1577
L1579:
	;
	v7544 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[82]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+416)) = v7544
	F_errmsg(m, int32(_a_F_InitPostgres_125), v44+int32(416))
	mBase = m.M
	v7550 = m.ExcPending
	if v7550 != 0 {
		goto L1
	} else {
		goto L1580
	}
L1580:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(225), int32(_a_F_InitPostgres_118))
	mBase = m.M
	v7555 = m.ExcPending
	if v7555 != 0 {
		goto L1
	} else {
		goto L1581
	}
L1581:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1582:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v7562 = m.ExcPending
	if v7562 != 0 {
		goto L1
	} else {
		goto L1583
	}
L1583:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_126), int32(0))
	mBase = m.M
	v7566 = m.ExcPending
	if v7566 != 0 {
		goto L1
	} else {
		goto L1584
	}
L1584:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(940), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7571 = m.ExcPending
	if v7571 != 0 {
		goto L1
	} else {
		goto L1585
	}
L1585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1586:
	;
	F_errcode(m, int32(_a_F_InitPostgres_127))
	mBase = m.M
	v7578 = m.ExcPending
	if v7578 != 0 {
		goto L1
	} else {
		goto L1587
	}
L1587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+320)) = int32(_a_F_InitPostgres_128)
	F_errmsg(m, int32(_a_F_InitPostgres_129), v6930+int32(320))
	mBase = m.M
	v7585 = m.ExcPending
	if v7585 != 0 {
		goto L1
	} else {
		goto L1588
	}
L1588:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(961), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7590 = m.ExcPending
	if v7590 != 0 {
		goto L1
	} else {
		goto L1589
	}
L1589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1590:
	;
	F_errcode(m, int32(_a_F_InitPostgres_127))
	mBase = m.M
	v7597 = m.ExcPending
	if v7597 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+336)) = int32(_a_F_InitPostgres_130)
	F_errmsg(m, int32(_a_F_InitPostgres_131), v6930+int32(336))
	mBase = m.M
	v7604 = m.ExcPending
	if v7604 != 0 {
		goto L1
	} else {
		goto L1592
	}
L1592:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(967), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7609 = m.ExcPending
	if v7609 != 0 {
		goto L1
	} else {
		goto L1593
	}
L1593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1594:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v7616 = m.ExcPending
	if v7616 != 0 {
		goto L1
	} else {
		goto L1595
	}
L1595:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_132), int32(0))
	mBase = m.M
	v7620 = m.ExcPending
	if v7620 != 0 {
		goto L1
	} else {
		goto L1596
	}
L1596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+304)) = int32(_a_F_InitPostgres_133)
	F_errdetail(m, int32(_a_F_InitPostgres_134), v6930+int32(304))
	mBase = m.M
	v7627 = m.ExcPending
	if v7627 != 0 {
		goto L1
	} else {
		goto L1597
	}
L1597:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(980), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7632 = m.ExcPending
	if v7632 != 0 {
		goto L1
	} else {
		goto L1598
	}
L1598:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1599:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v7639 = m.ExcPending
	if v7639 != 0 {
		goto L1
	} else {
		goto L1600
	}
L1600:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+288)) = v6918
	F_errmsg(m, int32(_a_F_InitPostgres_135), v6930+int32(288))
	mBase = m.M
	v7645 = m.ExcPending
	if v7645 != 0 {
		goto L1
	} else {
		goto L1601
	}
L1601:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1032), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7650 = m.ExcPending
	if v7650 != 0 {
		goto L1
	} else {
		goto L1602
	}
L1602:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1603:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v7661 = m.ExcPending
	if v7661 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+272)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_136), v6930+int32(272))
	mBase = m.M
	v7669 = m.ExcPending
	if v7669 != 0 {
		goto L1
	} else {
		goto L1605
	}
L1605:
	;
	F_errhint(m, int32(_a_F_InitPostgres_137), int32(0))
	mBase = m.M
	v7673 = m.ExcPending
	if v7673 != 0 {
		goto L1
	} else {
		goto L1606
	}
L1606:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1111), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7678 = m.ExcPending
	if v7678 != 0 {
		goto L1
	} else {
		goto L1607
	}
L1607:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1608:
	;
	v7693 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v7695 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[32]))
	v7696 = F_GetDatabasePath(m, v7693, v7695)
	mBase = m.M
	v7697 = m.ExcPending
	if v7697 != 0 {
		goto L1
	} else {
		goto L1609
	}
L1609:
	;
	if v6942 != 0 {
		goto L1611
	} else {
		goto L1612
	}
L1610:
	;
	v8316 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[33]))
	if v8316 != 0 {
		goto L1754
	} else {
		goto L1755
	}
L1611:
	;
	v7699 = F_access(m, v7696, int32(0))
	mBase = m.M
	if v7699 == int32(-1) {
		goto L1614
	} else {
		goto L1615
	}
L1612:
	;
	goto L1613
L1613:
	;
	F_SetDatabasePath(m, v7696)
	mBase = m.M
	v8267 = m.ExcPending
	if v8267 != 0 {
		goto L1
	} else {
		goto L1750
	}
L1614:
	;
	v7703 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[44]))
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v7707 = m.ExcPending
	if v7707 != 0 {
		goto L1
	} else {
		goto L1617
	}
L1615:
	;
	goto L1616
L1616:
	;
	F_ValidatePgVersion(m, v7696)
	mBase = m.M
	v7724 = m.ExcPending
	if v7724 != 0 {
		goto L1
	} else {
		goto L1622
	}
L1617:
	;
	if v7703 == int32(44) {
		goto L127
	} else {
		goto L1618
	}
L1618:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v7711 = m.ExcPending
	if v7711 != 0 {
		goto L1
	} else {
		goto L1619
	}
L1619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+32)) = v7696
	F_errmsg(m, int32(_a_F_InitPostgres_138), v6930+int32(32))
	mBase = m.M
	v7717 = m.ExcPending
	if v7717 != 0 {
		goto L1
	} else {
		goto L1620
	}
L1620:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1177), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v7722 = m.ExcPending
	if v7722 != 0 {
		goto L1
	} else {
		goto L1621
	}
L1621:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1622:
	;
	F_SetDatabasePath(m, v7696)
	mBase = m.M
	v7726 = m.ExcPending
	if v7726 != 0 {
		goto L1
	} else {
		goto L1623
	}
L1623:
	;
	F_pfree(m, v7696)
	mBase = m.M
	v7728 = m.ExcPending
	if v7728 != 0 {
		goto L1
	} else {
		goto L1624
	}
L1624:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v7730 = m.ExcPending
	if v7730 != 0 {
		goto L1
	} else {
		goto L1625
	}
L1625:
	;
	F_initialize_acl(m)
	mBase = m.M
	v7732 = m.ExcPending
	if v7732 != 0 {
		goto L1
	} else {
		goto L1626
	}
L1626:
	;
	v7735 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v7736 = F_SearchSysCache1(m, int32(21), v7735)
	mBase = m.M
	v7737 = m.ExcPending
	if v7737 != 0 {
		goto L1
	} else {
		goto L1627
	}
L1627:
	;
	if v7736 == int32(0) {
		goto L126
	} else {
		goto L1628
	}
L1628:
	;
	v7741 = v6930 + int32(448)
	v7742 = *(*int32)(unsafe.Add(mBase, uint32(v7736)+16))
	v7743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7742)+22)))
	v7744 = v7742 + v7743
	v7746 = v7744 + int32(4)
	v7749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7741))))
	v7752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7746))))
	if base.B2i32(v7749 == int32(0))|base.B2i32(v7749 != v7752) != 0 {
		v7770 = v7749
		v7771 = v7752
		goto L1630
	} else {
		goto L1631
	}
L1629:
	;
	if v7770-v7771 != 0 {
		goto L125
	} else {
		goto L1636
	}
L1630:
	;
	goto L1629
L1631:
	;
	v7755 = v7741
	v7756 = v7746
	goto L1632
L1632:
	;
	v7759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7756)+1)))
	v7760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7755)+1)))
	if v7760 == int32(0) {
		v7770 = v7760
		v7771 = v7759
		goto L1630
	} else {
		goto L1634
	}
L1633:
	;
	v7770 = v7760
	v7771 = v7759
	goto L1630
L1634:
	;
	v7763 = int32(1)
	if v7760 == v7759 {
		v7755 = v7755 + v7763
		v7756 = v7756 + v7763
		goto L1632
	} else {
		goto L1635
	}
L1635:
	;
	goto L1633
L1636:
	;
	v7774 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v7774 != int32(1) {
		goto L1637
	} else {
		goto L1638
	}
L1637:
	;
	v7980 = *(*int32)(unsafe.Add(mBase, uint32(v7744)+72))
	v7981 = m.G0
	v7983 = v7981 - int32(16)
	m.G0 = v7983
	if base.Ui32(int32(35)) <= base.Ui32(v7980) {
		goto L1665
	} else {
		goto L1666
	}
L1638:
	;
	v7778 = v6922 & int32(2)
	if v7778 == int32(0) {
		goto L1639
	} else {
		goto L1640
	}
L1639:
	;
	v7781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7744)+78)))
	if v7781&int32(1) == int32(0) {
		goto L124
	} else {
		goto L1642
	}
L1640:
	;
	goto L1641
L1641:
	;
	v7786 = int32(0)
	if base.B2i32(v7778 != v7786)|v6921 == v7786 {
		goto L1643
	} else {
		goto L1644
	}
L1642:
	;
	goto L1641
L1643:
	;
	v7793 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v7795 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	v7797 = F_object_aclcheck(m, int32(1262), v7793, v7795, int64(2048))
	mBase = m.M
	v7798 = m.ExcPending
	if v7798 != 0 {
		goto L1
	} else {
		goto L1646
	}
L1644:
	;
	goto L1645
L1645:
	;
	v7800 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[30]))
	v7803 = *(*int32)(unsafe.Add(mBase, uint32(v7744)+80))
	if v6921|(base.B2i32(v7800 != int32(1))|base.B2i32(v7803 < int32(0))) != 0 {
		goto L1637
	} else {
		goto L1648
	}
L1646:
	;
	if v7797 != 0 {
		goto L123
	} else {
		goto L1647
	}
L1647:
	;
	goto L1645
L1648:
	;
	v7809 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v7810 = int32(0)
	v7812 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[83]))
	v7814 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84]))
	v7818 = F_LWLockAcquire(m, v7814+int32(512), int32(1))
	mBase = m.M
	v7819 = m.ExcPending
	if v7819 != 0 {
		goto L1
	} else {
		goto L1649
	}
L1649:
	;
	v7820 = *(*int32)(unsafe.Add(mBase, uint32(v7812)))
	if int32(0) < v7820 {
		goto L1650
	} else {
		goto L1651
	}
L1650:
	;
	v7826 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[85]))
	v7828 = int32(0)
	v7830 = v7810
	goto L1653
L1651:
	;
	v7892 = v7810
	goto L1652
L1652:
	;
	v7932 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[84]))
	F_LWLockRelease(m, v7932+int32(512))
	mBase = m.M
	v7936 = m.ExcPending
	if v7936 != 0 {
		goto L1
	} else {
		goto L1663
	}
L1653:
	;
	v7872 = *(*int32)(unsafe.Add(mBase, uint32(v7812+int32(36)+v7828<<(uint(int32(2))%32))))
	v7875 = v7826 + v7872*int32(640)
	v7876 = *(*int32)(unsafe.Add(mBase, uint32(v7875)+44))
	if v7876 == int32(0) {
		v7886 = v7830
		goto L1655
	} else {
		goto L1656
	}
L1654:
	;
	v7892 = v7886
	goto L1652
L1655:
	;
	v7888 = v7828 + int32(1)
	if v7888 != v7820 {
		v7828 = v7888
		v7830 = v7886
		goto L1653
	} else {
		goto L1662
	}
L1656:
	;
	v7879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7875)+72)))
	if v7879 != int32(1) {
		v7886 = v7830
		goto L1655
	} else {
		goto L1657
	}
L1657:
	;
	if v7809 != 0 {
		goto L1658
	} else {
		goto L1659
	}
L1658:
	;
	v7882 = *(*int32)(unsafe.Add(mBase, uint32(v7875)+60))
	if v7882 != v7809 {
		v7886 = v7830
		goto L1655
	} else {
		goto L1661
	}
L1659:
	;
	goto L1660
L1660:
	;
	v7886 = v7830 + int32(1)
	goto L1655
L1661:
	;
	goto L1660
L1662:
	;
	goto L1654
L1663:
	;
	v7937 = *(*int32)(unsafe.Add(mBase, uint32(v7744)+80))
	if v7937 < v7892 {
		goto L122
	} else {
		goto L1664
	}
L1664:
	;
	goto L1637
L1665:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v7990 = m.ExcPending
	if v7990 != 0 {
		goto L1
	} else {
		goto L1668
	}
L1666:
	;
	goto L1667
L1667:
	;
	v8002 = v7980 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[86])) = v8002 + int32(_a_F_InitPostgres_139)
	m.G0 = v7983 + int32(16)
	v8012 = *(*int32)(unsafe.Add(mBase, uint32(v8002)+uint32(_c_F_InitPostgres[87])))
	goto L1671
L1668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7983))) = v7980
	F_errmsg_internal(m, int32(_a_F_InitPostgres_140), v7983)
	mBase = m.M
	v7994 = m.ExcPending
	if v7994 != 0 {
		goto L1
	} else {
		goto L1669
	}
L1669:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_141), int32(1290), int32(_a_F_InitPostgres_142))
	mBase = m.M
	v7999 = m.ExcPending
	if v7999 != 0 {
		goto L1
	} else {
		goto L1670
	}
L1670:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1671:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_143), v8012, int32(0), int32(1))
	mBase = m.M
	v8016 = m.ExcPending
	if v8016 != 0 {
		goto L1
	} else {
		goto L1672
	}
L1672:
	;
	v8019 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[86]))
	v8020 = *(*int32)(unsafe.Add(mBase, uint32(v8019)))
	goto L1673
L1673:
	;
	F_SetConfigOption(m, int32(_a_F_InitPostgres_144), v8020, int32(4), int32(1))
	mBase = m.M
	v8024 = m.ExcPending
	if v8024 != 0 {
		goto L1
	} else {
		goto L1674
	}
L1674:
	;
	v8027 = F_SysCacheGetAttrNotNull(m, int32(21), v7736, int32(13))
	mBase = m.M
	v8028 = m.ExcPending
	if v8028 != 0 {
		goto L1
	} else {
		goto L1675
	}
L1675:
	;
	v8029 = F_text_to_cstring(m, v8027)
	mBase = m.M
	v8030 = m.ExcPending
	if v8030 != 0 {
		goto L1
	} else {
		goto L1676
	}
L1676:
	;
	v8033 = F_SysCacheGetAttrNotNull(m, int32(21), v7736, int32(14))
	mBase = m.M
	v8034 = m.ExcPending
	if v8034 != 0 {
		goto L1
	} else {
		goto L1677
	}
L1677:
	;
	v8035 = F_text_to_cstring(m, v8033)
	mBase = m.M
	v8036 = m.ExcPending
	if v8036 != 0 {
		goto L1
	} else {
		goto L1678
	}
L1678:
	;
	v8038 = F_pg_perm_setlocale(m, int32(3), v8029)
	mBase = m.M
	v8039 = m.ExcPending
	if v8039 != 0 {
		goto L1
	} else {
		goto L1679
	}
L1679:
	;
	if v8038 == int32(0) {
		goto L121
	} else {
		goto L1680
	}
L1680:
	;
	v8043 = F_pg_perm_setlocale(m, int32(0), v8035)
	mBase = m.M
	v8044 = m.ExcPending
	if v8044 != 0 {
		goto L1
	} else {
		goto L1681
	}
L1681:
	;
	if v8043 == int32(0) {
		goto L120
	} else {
		goto L1682
	}
L1682:
	;
	v8047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8035))))
	if v8047 == int32(67) {
		goto L1685
	} else {
		goto L1686
	}
L1683:
	;
	v8083 = m.G0
	v8085 = v8083 - int32(32)
	m.G0 = v8085
	v8089 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v8090 = F_SearchSysCache1(m, int32(21), v8089)
	mBase = m.M
	v8091 = m.ExcPending
	if v8091 != 0 {
		goto L1
	} else {
		goto L1698
	}
L1684:
	;
	v8081 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[88])) = uint8(v8081)
	goto L1683
L1685:
	;
	v8050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8035)+1)))
	if v8050 == int32(0) {
		goto L1684
	} else {
		goto L1688
	}
L1686:
	;
	goto L1687
L1687:
	;
	v8053 = int32(_a_F_InitPostgres_145)
	v8056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8035))))
	v8059 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[89])))
	if base.B2i32(v8056 == int32(0))|base.B2i32(v8056 != v8059) != 0 {
		v8077 = v8056
		v8078 = v8059
		goto L1690
	} else {
		goto L1691
	}
L1688:
	;
	goto L1687
L1689:
	;
	if v8077-v8078 != 0 {
		goto L1683
	} else {
		goto L1696
	}
L1690:
	;
	goto L1689
L1691:
	;
	v8062 = v8035
	v8063 = v8053
	goto L1692
L1692:
	;
	v8066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8063)+1)))
	v8067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8062)+1)))
	if v8067 == int32(0) {
		v8077 = v8067
		v8078 = v8066
		goto L1690
	} else {
		goto L1694
	}
L1693:
	;
	v8077 = v8067
	v8078 = v8066
	goto L1690
L1694:
	;
	v8070 = int32(1)
	if v8067 == v8066 {
		v8062 = v8062 + v8070
		v8063 = v8063 + v8070
		goto L1692
	} else {
		goto L1695
	}
L1695:
	;
	goto L1693
L1696:
	;
	goto L1684
L1697:
	;
	v8159 = F_SysCacheGetAttr(m, int32(21), v7736, int32(17), v6930+int32(527))
	mBase = m.M
	v8160 = m.ExcPending
	if v8160 != 0 {
		goto L1
	} else {
		goto L1717
	}
L1698:
	;
	if v8090 != 0 {
		goto L1699
	} else {
		goto L1700
	}
L1699:
	;
	v8092 = *(*int32)(unsafe.Add(mBase, uint32(v8090)+16))
	v8093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8092)+22)))
	v8094 = v8092 + v8093
	v8095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8094)+76)))
	switch v8095 - int32(98) {
	case 0:
		goto L1703
	case 1:
		goto L1705
	default:
		goto L1704
	case 7:
		goto L1706
	}
L1700:
	;
	goto L1701
L1701:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8143 = m.ExcPending
	if v8143 != 0 {
		goto L1
	} else {
		goto L1714
	}
L1702:
	;
	v8131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8130)+4)) = uint8(v8131)
	F_ReleaseCatCache(m, v8090)
	mBase = m.M
	v8134 = m.ExcPending
	if v8134 != 0 {
		goto L1
	} else {
		goto L1713
	}
L1703:
	;
	v8127 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8128 = F_create_pg_locale_builtin(m, int32(100), v8127)
	mBase = m.M
	v8129 = m.ExcPending
	if v8129 != 0 {
		goto L1
	} else {
		goto L1712
	}
L1704:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8110 = m.ExcPending
	if v8110 != 0 {
		goto L1
	} else {
		goto L1709
	}
L1705:
	;
	v8104 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8105 = F_create_pg_locale_libc(m, int32(100), v8104)
	mBase = m.M
	v8106 = m.ExcPending
	if v8106 != 0 {
		goto L1
	} else {
		goto L1708
	}
L1706:
	;
	v8100 = F_create_pg_locale_icu(m)
	mBase = m.M
	v8101 = m.ExcPending
	if v8101 != 0 {
		goto L1
	} else {
		goto L1707
	}
L1707:
	;
	v8130 = v8100
	goto L1702
L1708:
	;
	v8130 = v8105
	goto L1702
L1709:
	;
	v8111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8094)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v8085)+20)) = v8111
	*(*int32)(unsafe.Add(mBase, uint32(v8085)+16)) = int32(_a_F_InitPostgres_146)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_147), v8085+int32(16))
	mBase = m.M
	v8119 = m.ExcPending
	if v8119 != 0 {
		goto L1
	} else {
		goto L1710
	}
L1710:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_148), int32(1179), int32(_a_F_InitPostgres_146))
	mBase = m.M
	v8124 = m.ExcPending
	if v8124 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1712:
	;
	v8130 = v8128
	goto L1702
L1713:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[90])) = v8130
	m.G0 = v8085 + int32(32)
	goto L1697
L1714:
	;
	v8145 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v8085))) = v8145
	F_errmsg_internal(m, int32(_a_F_InitPostgres_149), v8085)
	mBase = m.M
	v8149 = m.ExcPending
	if v8149 != 0 {
		goto L1
	} else {
		goto L1715
	}
L1715:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_148), int32(1165), int32(_a_F_InitPostgres_146))
	mBase = m.M
	v8154 = m.ExcPending
	if v8154 != 0 {
		goto L1
	} else {
		goto L1716
	}
L1716:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1717:
	;
	v8161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6930)+527)))
	if v8161 != 0 {
		goto L1718
	} else {
		goto L1719
	}
L1718:
	;
	F_ReleaseCatCache(m, v7736)
	mBase = m.M
	v8265 = m.ExcPending
	if v8265 != 0 {
		goto L1
	} else {
		goto L1749
	}
L1719:
	;
	v8162 = F_text_to_cstring(m, v8159)
	mBase = m.M
	v8163 = m.ExcPending
	if v8163 != 0 {
		goto L1
	} else {
		goto L1720
	}
L1720:
	;
	v8165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7744)+76)))
	if v8165 != int32(99) {
		goto L1722
	} else {
		goto L1723
	}
L1721:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), v8257, int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8260 = m.ExcPending
	if v8260 != 0 {
		goto L1
	} else {
		goto L1748
	}
L1722:
	;
	v8170 = F_SysCacheGetAttrNotNull(m, int32(21), v7736, int32(15))
	mBase = m.M
	v8171 = m.ExcPending
	if v8171 != 0 {
		goto L1
	} else {
		goto L1725
	}
L1723:
	;
	v8176 = v8029
	v8177 = int32(99)
	goto L1724
L1724:
	;
	v8179 = F_get_collation_actual_version(m, base.I32_extend8_s(v8177), v8176)
	mBase = m.M
	v8180 = m.ExcPending
	if v8180 != 0 {
		goto L1
	} else {
		goto L1727
	}
L1725:
	;
	v8172 = F_text_to_cstring(m, v8170)
	mBase = m.M
	v8173 = m.ExcPending
	if v8173 != 0 {
		goto L1
	} else {
		goto L1726
	}
L1726:
	;
	v8174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7744)+76)))
	v8176 = v8172
	v8177 = v8174
	goto L1724
L1727:
	;
	if v8179 == int32(0) {
		goto L1728
	} else {
		goto L1729
	}
L1728:
	;
	v8185 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v8186 = m.ExcPending
	if v8186 != 0 {
		goto L1
	} else {
		goto L1731
	}
L1729:
	;
	goto L1730
L1730:
	;
	v8200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8179))))
	v8203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8162))))
	if base.B2i32(v8200 == int32(0))|base.B2i32(v8200 != v8203) != 0 {
		v8221 = v8200
		v8222 = v8203
		goto L1735
	} else {
		goto L1736
	}
L1731:
	;
	if v8185 == int32(0) {
		goto L1718
	} else {
		goto L1732
	}
L1732:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+96)) = v6930 + int32(448)
	F_errmsg_internal(m, int32(_a_F_InitPostgres_151), v6930+int32(96))
	mBase = m.M
	v8196 = m.ExcPending
	if v8196 != 0 {
		goto L1
	} else {
		goto L1733
	}
L1733:
	;
	v8257 = int32(468)
	goto L1721
L1734:
	;
	if v8221-v8222 == int32(0) {
		goto L1718
	} else {
		goto L1741
	}
L1735:
	;
	goto L1734
L1736:
	;
	v8206 = v8179
	v8207 = v8162
	goto L1737
L1737:
	;
	v8210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8207)+1)))
	v8211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8206)+1)))
	if v8211 == int32(0) {
		v8221 = v8211
		v8222 = v8210
		goto L1735
	} else {
		goto L1739
	}
L1738:
	;
	v8221 = v8211
	v8222 = v8210
	goto L1735
L1739:
	;
	v8214 = int32(1)
	if v8211 == v8210 {
		v8206 = v8206 + v8214
		v8207 = v8207 + v8214
		goto L1737
	} else {
		goto L1740
	}
L1740:
	;
	goto L1738
L1741:
	;
	v8228 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v8229 = m.ExcPending
	if v8229 != 0 {
		goto L1
	} else {
		goto L1742
	}
L1742:
	;
	if v8228 == int32(0) {
		goto L1718
	} else {
		goto L1743
	}
L1743:
	;
	v8233 = v6930 + int32(448)
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+144)) = v8233
	F_errmsg(m, int32(_a_F_InitPostgres_152), v6930+int32(144))
	mBase = m.M
	v8239 = m.ExcPending
	if v8239 != 0 {
		goto L1
	} else {
		goto L1744
	}
L1744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+132)) = v8179
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+128)) = v8162
	F_errdetail(m, int32(_a_F_InitPostgres_153), v6930+int32(128))
	mBase = m.M
	v8246 = m.ExcPending
	if v8246 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1745:
	;
	v8247 = F_quote_identifier(m, v8233)
	mBase = m.M
	v8248 = m.ExcPending
	if v8248 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+112)) = v8247
	F_errhint(m, int32(_a_F_InitPostgres_154), v6930+int32(112))
	mBase = m.M
	v8254 = m.ExcPending
	if v8254 != 0 {
		goto L1
	} else {
		goto L1747
	}
L1747:
	;
	v8257 = int32(479)
	goto L1721
L1748:
	;
	goto L1718
L1749:
	;
	goto L1610
L1750:
	;
	F_pfree(m, v7696)
	mBase = m.M
	v8269 = m.ExcPending
	if v8269 != 0 {
		goto L1
	} else {
		goto L1751
	}
L1751:
	;
	F_RelationCacheInitializePhase3(m)
	mBase = m.M
	v8271 = m.ExcPending
	if v8271 != 0 {
		goto L1
	} else {
		goto L1752
	}
L1752:
	;
	F_initialize_acl(m)
	mBase = m.M
	v8273 = m.ExcPending
	if v8273 != 0 {
		goto L1
	} else {
		goto L1753
	}
L1753:
	;
	goto L1610
L1754:
	;
	F_process_startup_options(m, v8316, v6921)
	mBase = m.M
	v8318 = m.ExcPending
	if v8318 != 0 {
		goto L1
	} else {
		goto L1757
	}
L1755:
	;
	goto L1756
L1756:
	;
	v8320 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	v8322 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[91]))
	v8324 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[2])))
	if v8324 == int32(1) {
		goto L1758
	} else {
		goto L1759
	}
L1757:
	;
	goto L1756
L1758:
	;
	v8329 = F_table_open(m, int32(2964), int32(1))
	mBase = m.M
	v8330 = m.ExcPending
	if v8330 != 0 {
		goto L1
	} else {
		goto L1761
	}
L1759:
	;
	goto L1760
L1760:
	;
	v8360 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[79]))
	if int32(0) < v8360 {
		goto L1770
	} else {
		goto L1771
	}
L1761:
	;
	v8332 = F_GetCatalogSnapshot(m, int32(2964))
	mBase = m.M
	v8333 = m.ExcPending
	if v8333 != 0 {
		goto L1
	} else {
		goto L1762
	}
L1762:
	;
	v8334 = F_RegisterSnapshot(m, v8332)
	mBase = m.M
	v8335 = m.ExcPending
	if v8335 != 0 {
		goto L1
	} else {
		goto L1763
	}
L1763:
	;
	F_ApplySetting(m, v8334, v8320, v8322, v8329, int32(8))
	mBase = m.M
	v8338 = m.ExcPending
	if v8338 != 0 {
		goto L1
	} else {
		goto L1764
	}
L1764:
	;
	F_ApplySetting(m, v8334, int32(0), v8322, v8329, int32(7))
	mBase = m.M
	v8342 = m.ExcPending
	if v8342 != 0 {
		goto L1
	} else {
		goto L1765
	}
L1765:
	;
	F_ApplySetting(m, v8334, v8320, int32(0), v8329, int32(6))
	mBase = m.M
	v8346 = m.ExcPending
	if v8346 != 0 {
		goto L1
	} else {
		goto L1766
	}
L1766:
	;
	v8347 = int32(0)
	F_ApplySetting(m, v8334, v8347, v8347, v8329, int32(5))
	mBase = m.M
	v8351 = m.ExcPending
	if v8351 != 0 {
		goto L1
	} else {
		goto L1767
	}
L1767:
	;
	F_UnregisterSnapshot(m, v8334)
	mBase = m.M
	v8353 = m.ExcPending
	if v8353 != 0 {
		goto L1
	} else {
		goto L1768
	}
L1768:
	;
	F_relation_close(m, v8329, int32(1))
	mBase = m.M
	v8356 = m.ExcPending
	if v8356 != 0 {
		goto L1
	} else {
		goto L1769
	}
L1769:
	;
	goto L1760
L1770:
	;
	F_pg_usleep(m, v8360*int32(_a_F_InitPostgres_123))
	mBase = m.M
	goto L1772
L1771:
	;
	goto L1772
L1772:
	;
	v8366 = m.G0
	v8368 = v8366 - int32(16)
	m.G0 = v8368
	v8371 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[0]))
	if v8371 == int32(0) {
		goto L1774
	} else {
		goto L1775
	}
L1773:
	;
	m.G0 = v8368 + int32(16)
	F_InitializeClientEncoding(m)
	mBase = m.M
	v8456 = m.ExcPending
	if v8456 != 0 {
		goto L1
	} else {
		goto L1782
	}
L1774:
	;
	v8374 = int32(_a_F_InitPostgres_8)
	v8375 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22]))
	v8378 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v8378
	v8380 = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v8368)+12)) = v8380
	*(*int32)(unsafe.Add(mBase, uint32(v8368)+8)) = v8380
	v8387 = F_list_make1_impl(m, int32(472), v8368+int32(8))
	mBase = m.M
	v8388 = m.ExcPending
	if v8388 != 0 {
		goto L1
	} else {
		goto L1777
	}
L1775:
	;
	goto L1776
L1776:
	;
	F_CacheRegisterSyscacheCallback(m, int32(38), int32(470), int32(0))
	mBase = m.M
	v8428 = m.ExcPending
	if v8428 != 0 {
		goto L1
	} else {
		goto L1778
	}
L1777:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[22])) = v8375
	v8391 = int32(_a_F_InitPostgres_155)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[92])) = v8387
	v8393 = int32(_a_F_InitPostgres_156)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[93])) = int32(11)
	v8396 = int32(_a_F_InitPostgres_157)
	v8397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[94])) = uint8(v8397)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[95])) = uint8(v8397)
	v8404 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[96])) = v8404
	v8408 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[92]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[97])) = v8408
	v8412 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[93]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[98])) = v8412
	v8416 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[94])))
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[99])) = uint8(v8416)
	v8418 = int32(_a_F_InitPostgres_158)
	v8420 = *(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[100]))
	*(*int64)(unsafe.Add(mBase, _c_F_InitPostgres[100])) = v8420 + int64(1)
	goto L1773
L1778:
	;
	F_CacheRegisterSyscacheCallback(m, int32(11), int32(470), int32(0))
	mBase = m.M
	v8433 = m.ExcPending
	if v8433 != 0 {
		goto L1
	} else {
		goto L1779
	}
L1779:
	;
	F_CacheRegisterSyscacheCallback(m, int32(9), int32(470), int32(0))
	mBase = m.M
	v8438 = m.ExcPending
	if v8438 != 0 {
		goto L1
	} else {
		goto L1780
	}
L1780:
	;
	F_CacheRegisterSyscacheCallback(m, int32(21), int32(470), int32(0))
	mBase = m.M
	v8443 = m.ExcPending
	if v8443 != 0 {
		goto L1
	} else {
		goto L1781
	}
L1781:
	;
	v8445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[95])) = uint8(v8445)
	v8448 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitPostgres[101])) = uint8(v8448)
	goto L1773
L1782:
	;
	v8459 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[21]))
	v8461 = F_MemoryContextAllocZero(m, v8459, int32(20))
	mBase = m.M
	v8462 = m.ExcPending
	if v8462 != 0 {
		goto L1
	} else {
		goto L1783
	}
L1783:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[102])) = v8461
	if v6922&int32(1) != 0 {
		goto L1784
	} else {
		goto L1785
	}
L1784:
	;
	v8467 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[103]))
	F_load_libraries(m, v8467, int32(_a_F_InitPostgres_159), int32(0))
	mBase = m.M
	v8471 = m.ExcPending
	if v8471 != 0 {
		goto L1
	} else {
		goto L1787
	}
L1785:
	;
	goto L1786
L1786:
	;
	if v6942 == int32(0) {
		v8537 = v6930
		goto L128
	} else {
		goto L1789
	}
L1787:
	;
	v8473 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[104]))
	F_load_libraries(m, v8473, int32(_a_F_InitPostgres_160), int32(1))
	mBase = m.M
	v8477 = m.ExcPending
	if v8477 != 0 {
		goto L1
	} else {
		goto L1788
	}
L1788:
	;
	goto L1786
L1789:
	;
	goto L133
L1790:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v8524 = m.ExcPending
	if v8524 != 0 {
		goto L1
	} else {
		goto L1791
	}
L1791:
	;
	v8537 = v6930
	goto L128
L1792:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+16)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_135), v6930+int32(16))
	mBase = m.M
	v8579 = m.ExcPending
	if v8579 != 0 {
		goto L1
	} else {
		goto L1793
	}
L1793:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930))) = v7696
	F_errdetail(m, int32(_a_F_InitPostgres_161), v6930)
	mBase = m.M
	v8583 = m.ExcPending
	if v8583 != 0 {
		goto L1
	} else {
		goto L1794
	}
L1794:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1172), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v8588 = m.ExcPending
	if v8588 != 0 {
		goto L1
	} else {
		goto L1795
	}
L1795:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1796:
	;
	v8594 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+48)) = v8594
	F_errmsg_internal(m, int32(_a_F_InitPostgres_149), v6930+int32(48))
	mBase = m.M
	v8600 = m.ExcPending
	if v8600 != 0 {
		goto L1
	} else {
		goto L1797
	}
L1797:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(335), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8605 = m.ExcPending
	if v8605 != 0 {
		goto L1
	} else {
		goto L1798
	}
L1798:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1799:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8612 = m.ExcPending
	if v8612 != 0 {
		goto L1
	} else {
		goto L1800
	}
L1800:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+224)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_162), v6930+int32(224))
	mBase = m.M
	v8620 = m.ExcPending
	if v8620 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1801:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+212)) = v7746
	v8623 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostgres[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+208)) = v8623
	F_errdetail(m, int32(_a_F_InitPostgres_163), v6930+int32(208))
	mBase = m.M
	v8629 = m.ExcPending
	if v8629 != 0 {
		goto L1
	} else {
		goto L1802
	}
L1802:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(345), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8634 = m.ExcPending
	if v8634 != 0 {
		goto L1
	} else {
		goto L1803
	}
L1803:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1804:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v8641 = m.ExcPending
	if v8641 != 0 {
		goto L1
	} else {
		goto L1805
	}
L1805:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+192)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_164), v6930+int32(192))
	mBase = m.M
	v8649 = m.ExcPending
	if v8649 != 0 {
		goto L1
	} else {
		goto L1806
	}
L1806:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(365), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8654 = m.ExcPending
	if v8654 != 0 {
		goto L1
	} else {
		goto L1807
	}
L1807:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1808:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v8661 = m.ExcPending
	if v8661 != 0 {
		goto L1
	} else {
		goto L1809
	}
L1809:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+176)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_165), v6930+int32(176))
	mBase = m.M
	v8669 = m.ExcPending
	if v8669 != 0 {
		goto L1
	} else {
		goto L1810
	}
L1810:
	;
	F_errdetail(m, int32(_a_F_InitPostgres_166), int32(0))
	mBase = m.M
	v8673 = m.ExcPending
	if v8673 != 0 {
		goto L1
	} else {
		goto L1811
	}
L1811:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(378), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8678 = m.ExcPending
	if v8678 != 0 {
		goto L1
	} else {
		goto L1812
	}
L1812:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1813:
	;
	F_errcode(m, int32(_a_F_InitPostgres_127))
	mBase = m.M
	v8685 = m.ExcPending
	if v8685 != 0 {
		goto L1
	} else {
		goto L1814
	}
L1814:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+160)) = v6930 + int32(448)
	F_errmsg(m, int32(_a_F_InitPostgres_167), v6930+int32(160))
	mBase = m.M
	v8693 = m.ExcPending
	if v8693 != 0 {
		goto L1
	} else {
		goto L1815
	}
L1815:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(399), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8698 = m.ExcPending
	if v8698 != 0 {
		goto L1
	} else {
		goto L1816
	}
L1816:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1817:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_168), int32(0))
	mBase = m.M
	v8706 = m.ExcPending
	if v8706 != 0 {
		goto L1
	} else {
		goto L1818
	}
L1818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+64)) = v8029
	F_errdetail(m, int32(_a_F_InitPostgres_169), v6930-int32(-64))
	mBase = m.M
	v8712 = m.ExcPending
	if v8712 != 0 {
		goto L1
	} else {
		goto L1819
	}
L1819:
	;
	F_errhint(m, int32(_a_F_InitPostgres_170), int32(0))
	mBase = m.M
	v8716 = m.ExcPending
	if v8716 != 0 {
		goto L1
	} else {
		goto L1820
	}
L1820:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(425), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8721 = m.ExcPending
	if v8721 != 0 {
		goto L1
	} else {
		goto L1821
	}
L1821:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1822:
	;
	F_errmsg(m, int32(_a_F_InitPostgres_168), int32(0))
	mBase = m.M
	v8729 = m.ExcPending
	if v8729 != 0 {
		goto L1
	} else {
		goto L1823
	}
L1823:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+80)) = v8035
	F_errdetail(m, int32(_a_F_InitPostgres_171), v6930+int32(80))
	mBase = m.M
	v8735 = m.ExcPending
	if v8735 != 0 {
		goto L1
	} else {
		goto L1824
	}
L1824:
	;
	F_errhint(m, int32(_a_F_InitPostgres_170), int32(0))
	mBase = m.M
	v8739 = m.ExcPending
	if v8739 != 0 {
		goto L1
	} else {
		goto L1825
	}
L1825:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(432), int32(_a_F_InitPostgres_150))
	mBase = m.M
	v8744 = m.ExcPending
	if v8744 != 0 {
		goto L1
	} else {
		goto L1826
	}
L1826:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1827:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v8752 = m.ExcPending
	if v8752 != 0 {
		goto L1
	} else {
		goto L1828
	}
L1828:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6930)+256)) = v6918
	F_errmsg(m, int32(_a_F_InitPostgres_135), v6930+int32(256))
	mBase = m.M
	v8758 = m.ExcPending
	if v8758 != 0 {
		goto L1
	} else {
		goto L1829
	}
L1829:
	;
	F_errdetail(m, int32(_a_F_InitPostgres_172), int32(0))
	mBase = m.M
	v8762 = m.ExcPending
	if v8762 != 0 {
		goto L1
	} else {
		goto L1830
	}
L1830:
	;
	F_errfinish(m, int32(_a_F_InitPostgres_1), int32(1097), int32(_a_F_InitPostgres_0))
	mBase = m.M
	v8767 = m.ExcPending
	if v8767 != 0 {
		goto L1
	} else {
		goto L1831
	}
L1831:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
