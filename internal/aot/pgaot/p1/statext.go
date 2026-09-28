package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_clauselist_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 float64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v302 float64
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v474 int32
	_ = v474
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v576 int32
	_ = v576
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v775 int32
	_ = v775
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v904 float64
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int64
	_ = v927
	var v932 int32
	_ = v932
	var v948 int32
	_ = v948
	var v971 float64
	_ = v971
	var v980 int32
	_ = v980
	var v981 float64
	_ = v981
	var v982 float64
	_ = v982
	var v986 int32
	_ = v986
	var v989 float64
	_ = v989
	var v990 float64
	_ = v990
	var v993 float64
	_ = v993
	var v995 float64
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1041 float64
	_ = v1041
	var v1042 float64
	_ = v1042
	var v1043 float64
	_ = v1043
	var v1044 float64
	_ = v1044
	var v1046 float64
	_ = v1046
	var v1054 float64
	_ = v1054
	var v1056 float64
	_ = v1056
	var v1058 float64
	_ = v1058
	var v1059 float64
	_ = v1059
	var v1067 float64
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 float64
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1094 int32
	_ = v1094
	var v1118 float64
	_ = v1118
	var v1120 float64
	_ = v1120
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 float64
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 float64
	_ = v1135
	var v1137 float64
	_ = v1137
	var v1145 float64
	_ = v1145
	var v1146 float64
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int64
	_ = v1182
	var v1190 int32
	_ = v1190
	var v1204 int32
	_ = v1204
	var v1230 float64
	_ = v1230
	var v1236 int32
	_ = v1236
	var v1237 float64
	_ = v1237
	var v1238 float64
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 float64
	_ = v1245
	var v1246 float64
	_ = v1246
	var v1247 float64
	_ = v1247
	var v1250 float64
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1256 float64
	_ = v1256
	var v1257 float64
	_ = v1257
	var v1260 float64
	_ = v1260
	var v1261 float64
	_ = v1261
	var v1264 float64
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1318 float64
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 float64
	_ = v1330
	var v1331 float64
	_ = v1331
	var v1332 float64
	_ = v1332
	var v1334 float64
	_ = v1334
	var v1342 float64
	_ = v1342
	var v1344 float64
	_ = v1344
	var v1346 float64
	_ = v1346
	var v1347 float64
	_ = v1347
	var v1355 float64
	_ = v1355
	var v1356 float64
	_ = v1356
	var v1357 float64
	_ = v1357
	var v1358 float64
	_ = v1358
	var v1359 float64
	_ = v1359
	var v1360 float64
	_ = v1360
	var v1362 float64
	_ = v1362
	var v1370 float64
	_ = v1370
	var v1372 float64
	_ = v1372
	var v1374 float64
	_ = v1374
	var v1375 float64
	_ = v1375
	var v1383 float64
	_ = v1383
	var v1385 float64
	_ = v1385
	var v1393 float64
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1431 float64
	_ = v1431
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1458 int32
	_ = v1458
	var v1476 float64
	_ = v1476
	var v1485 int32
	_ = v1485
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 float64
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1586 int32
	_ = v1586
	var v1660 int32
	_ = v1660
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1695 int32
	_ = v1695
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1864 int32
	_ = v1864
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1913 int32
	_ = v1913
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1998 int32
	_ = v1998
	var v2004 int32
	_ = v2004
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2020 int32
	_ = v2020
	var v2027 int32
	_ = v2027
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2109 int32
	_ = v2109
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2204 float64
	_ = v2204
	var v2205 float64
	_ = v2205
	var v2211 int32
	_ = v2211
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2232 int32
	_ = v2232
	var v2235 int32
	_ = v2235
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2281 int32
	_ = v2281
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2390 int32
	_ = v2390
	var v2402 int32
	_ = v2402
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2476 int32
	_ = v2476
	var v2479 int32
	_ = v2479
	var v2494 int32
	_ = v2494
	var v2500 int32
	_ = v2500
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2554 int32
	_ = v2554
	var v2582 int32
	_ = v2582
	var v2598 int32
	_ = v2598
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2643 int64
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2670 int32
	_ = v2670
	var v2674 int32
	_ = v2674
	var v2682 int32
	_ = v2682
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2763 int32
	_ = v2763
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2810 int32
	_ = v2810
	var v2820 int32
	_ = v2820
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2962 int32
	_ = v2962
	var v2977 int32
	_ = v2977
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3022 int32
	_ = v3022
	var v3064 int32
	_ = v3064
	var v3102 int32
	_ = v3102
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3111 int32
	_ = v3111
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3131 int32
	_ = v3131
	var v3138 float64
	_ = v3138
	var v3139 float64
	_ = v3139
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3159 int32
	_ = v3159
	var v3161 int32
	_ = v3161
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3183 float64
	_ = v3183
	var v3184 float64
	_ = v3184
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3214 int32
	_ = v3214
	var v3224 float64
	_ = v3224
	var v3225 float64
	_ = v3225
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3241 int64
	_ = v3241
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3253 int32
	_ = v3253
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3260 int64
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3262 int64
	_ = v3262
	var v3263 int32
	_ = v3263
	var v3264 int64
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3266 int64
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int64
	_ = v3268
	var v3272 int64
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3276 int32
	_ = v3276
	var v3277 int64
	_ = v3277
	var v3280 int64
	_ = v3280
	var v3285 int32
	_ = v3285
	var v3294 int32
	_ = v3294
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3330 int32
	_ = v3330
	var v3333 int32
	_ = v3333
	var v3341 int32
	_ = v3341
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3387 int32
	_ = v3387
	var v3392 int32
	_ = v3392
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3428 int32
	_ = v3428
	var v3431 float64
	_ = v3431
	var v3432 float64
	_ = v3432
	var v3448 int32
	_ = v3448
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3495 int32
	_ = v3495
	var v3497 int32
	_ = v3497
	var v3504 int32
	_ = v3504
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3542 int32
	_ = v3542
	var v3583 int32
	_ = v3583
	var v3589 int32
	_ = v3589
	var v3593 int32
	_ = v3593
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3601 int64
	_ = v3601
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3610 int32
	_ = v3610
	var v3613 int32
	_ = v3613
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3620 int64
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int64
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int64
	_ = v3624
	var v3625 int32
	_ = v3625
	var v3626 int64
	_ = v3626
	var v3627 int32
	_ = v3627
	var v3628 int64
	_ = v3628
	var v3632 int64
	_ = v3632
	var v3633 int32
	_ = v3633
	var v3636 int32
	_ = v3636
	var v3637 int64
	_ = v3637
	var v3640 int64
	_ = v3640
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3650 int32
	_ = v3650
	var v3655 int32
	_ = v3655
	var v3691 int32
	_ = v3691
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3699 int32
	_ = v3699
	var v3707 int32
	_ = v3707
	var v3713 int32
	_ = v3713
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3793 int32
	_ = v3793
	var v3795 int32
	_ = v3795
	var v3838 int64
	_ = v3838
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3847 int32
	_ = v3847
	var v3850 int32
	_ = v3850
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3857 int64
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3859 int64
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3861 int64
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3863 int64
	_ = v3863
	var v3864 int32
	_ = v3864
	var v3865 int64
	_ = v3865
	var v3869 int64
	_ = v3869
	var v3870 int32
	_ = v3870
	var v3873 int32
	_ = v3873
	var v3874 int64
	_ = v3874
	var v3877 int64
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3892 int32
	_ = v3892
	var v3895 int32
	_ = v3895
	var v3898 int32
	_ = v3898
	var v3902 int32
	_ = v3902
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3910 int32
	_ = v3910
	var v3917 int32
	_ = v3917
	var v3919 int32
	_ = v3919
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3941 int32
	_ = v3941
	var v3950 int32
	_ = v3950
	var v3958 int32
	_ = v3958
	var v3989 int32
	_ = v3989
	var v3992 int32
	_ = v3992
	var v3996 int32
	_ = v3996
	var v4003 int32
	_ = v4003
	var v4006 int32
	_ = v4006
	var v4009 int32
	_ = v4009
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4041 int32
	_ = v4041
	var v4043 int32
	_ = v4043
	var v4047 int32
	_ = v4047
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4050 int32
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4054 int32
	_ = v4054
	var v4055 int32
	_ = v4055
	var v4056 int32
	_ = v4056
	var v4058 int32
	_ = v4058
	var v4071 int32
	_ = v4071
	var v4105 float64
	_ = v4105
	var v4106 int32
	_ = v4106
	var v4116 int32
	_ = v4116
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4122 int32
	_ = v4122
	var v4126 int32
	_ = v4126
	var v4129 int32
	_ = v4129
	var v4131 int32
	_ = v4131
	var v4134 int32
	_ = v4134
	var v4141 int32
	_ = v4141
	var v4143 int32
	_ = v4143
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4165 int32
	_ = v4165
	var v4210 int32
	_ = v4210
	var v4218 int32
	_ = v4218
	var v4254 float64
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4256 int32
	_ = v4256
	var v4259 int32
	_ = v4259
	var v4260 int32
	_ = v4260
	var v4273 int32
	_ = v4273
	var v4300 float64
	_ = v4300
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4315 float64
	_ = v4315
	var v4316 float64
	_ = v4316
	var v4317 int32
	_ = v4317
	var v4318 int32
	_ = v4318
	var v4319 int32
	_ = v4319
	var v4331 int32
	_ = v4331
	var v4358 float64
	_ = v4358
	var v4367 int32
	_ = v4367
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4370 float64
	_ = v4370
	var v4373 int32
	_ = v4373
	var v4374 float64
	_ = v4374
	var v4380 float64
	_ = v4380
	var v4434 int32
	_ = v4434
	var v4435 float64
	_ = v4435
	var v4436 int32
	_ = v4436
	var v4448 int32
	_ = v4448
	var v4451 int32
	_ = v4451
	var v4476 float64
	_ = v4476
	var v4486 int32
	_ = v4486
	var v4487 float64
	_ = v4487
	var v4489 float64
	_ = v4489
	var v4491 float64
	_ = v4491
	var v4493 float64
	_ = v4493
	var v4494 float64
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4496 int32
	_ = v4496
	var v4498 int32
	_ = v4498
	var v4510 int32
	_ = v4510
	var v4535 float64
	_ = v4535
	var v4544 int32
	_ = v4544
	var v4551 int32
	_ = v4551
	var v4576 float64
	_ = v4576
	var v4587 float64
	_ = v4587
	var v4588 float64
	_ = v4588
	var v4589 int32
	_ = v4589
	var v4592 int32
	_ = v4592
	var v4627 float64
	_ = v4627
	var v4635 float64
	_ = v4635
	var v4725 float64
	_ = v4725
	var v4727 int32
	_ = v4727
	var v4729 int32
	_ = v4729
	var v4765 float64
	_ = v4765
	var v4780 int32
	_ = v4780
	var v4816 int32
	_ = v4816
	var v4818 int32
	_ = v4818
	var v4820 int32
	_ = v4820
	var v4837 int32
	_ = v4837
	var v4857 float64
	_ = v4857
	var v4864 int32
	_ = v4864
	var v4866 int32
	_ = v4866
	var v4868 int32
	_ = v4868
	var v4870 int32
	_ = v4870
	var v4878 int32
	_ = v4878
	var v4887 int32
	_ = v4887
	var v4889 int32
	_ = v4889
	var v4905 float64
	_ = v4905
	var v4906 float64
	_ = v4906
	var v4913 int32
	_ = v4913
	var v4930 int32
	_ = v4930
	var v4932 int32
	_ = v4932
	var v4948 float64
	_ = v4948
	var v4949 float64
	_ = v4949
	var v4975 int32
	_ = v4975
	var v4993 float64
	_ = v4993
	v42 = m.G0
	v44 = v42 - int32(48)
	m.G0 = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v46 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l7 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l5)+76))
	v60 = v46 + v47<<(uint(int32(2))%32)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+76))
	v60 = v53 + v54<<(uint(int32(2))%32) - int32(4)
	goto L1
L5:
	;
	v63 = float64(0)
	goto L7
L6:
	;
	v63 = float64(1)
	goto L7
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l5)+120))
	if v64 == int32(0) {
		v1442 = l0
		v1443 = l1
		v1444 = l2
		v1445 = l3
		v1446 = l4
		v1447 = l5
		v1448 = l6
		v1449 = l7
		v1458 = v44
		v1476 = v63
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v1449 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 <= int32(0) {
		v1442 = l0
		v1443 = l1
		v1444 = l2
		v1445 = l3
		v1446 = l4
		v1447 = l5
		v1448 = l6
		v1449 = l7
		v1458 = v44
		v1476 = v63
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v81 = int32(0)
	goto L11
L11:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v71+v81<<(uint(int32(2))%32))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+16)))
	if v117 != int32(109) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if l1 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v121 = v81 + int32(1)
	if v121 != v67 {
		v81 = v121
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v1442 = l0
	v1443 = l1
	v1444 = l2
	v1445 = l3
	v1446 = l4
	v1447 = l5
	v1448 = l6
	v1449 = l7
	v1458 = v44
	v1476 = v63
	goto L8
L17:
	;
	v268 = l0
	v269 = l1
	v270 = l2
	v271 = l3
	v272 = l4
	v273 = l5
	v274 = l6
	v275 = l7
	v276 = v235
	v284 = v44
	v290 = v249
	v291 = v250
	v293 = v70
	v302 = v63
	goto L36
L18:
	;
	v125 = int32(4)
	v128 = F_palloc_mul(m, v125, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v136 = int32(4)
	v137 = l1 + v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v140 = F_palloc_mul(m, v136, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L21
	} else {
		goto L24
	}
L21:
	;
	return float64(0)
L22:
	;
	v134 = F_palloc_mul(m, int32(4), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v235 = v125
	v249 = v128
	v250 = v134
	goto L17
L24:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v144 = F_palloc_mul(m, int32(4), v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v146 <= int32(0) {
		v235 = v137
		v249 = v140
		v250 = v144
		goto L17
	} else {
		goto L26
	}
L26:
	;
	v159 = int32(0)
	goto L27
L27:
	;
	v192 = v159 << (uint(int32(2)) % 32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v192+v193)))
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v196
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v202 = F_bms_is_member(m, v159, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v235 = v137
	v249 = v140
	v250 = v144
	goto L17
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192+v144))) = v221
	v224 = v159 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v224 < v225 {
		v159 = v224
		goto L27
	} else {
		goto L35
	}
L30:
	;
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v140))) = v218
	v221 = v218
	goto L29
L31:
	;
	if v202 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l5)+76))
	v209 = F_statext_is_compatible_clause(m, l0, v195, v204, v44+int32(32), v44+int32(24))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v209 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v192+v140))) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v221 = v216
	goto L29
L35:
	;
	goto L28
L36:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+20)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v273)+120))
	if v269 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v867 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L39:
	;
	v313 = int32(0)
	v318 = F_choose_best_statistics(m, v310, v309&int32(1), v290, v291, v313)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v325 = F_choose_best_statistics(m, v310, v309&int32(1), v290, v291, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L21
	} else {
		goto L44
	}
L42:
	;
	if v318 == int32(0) {
		v1442 = v268
		v1443 = v269
		v1444 = v270
		v1445 = v271
		v1446 = v272
		v1447 = v273
		v1448 = v274
		v1449 = v275
		v1458 = v284
		v1476 = v302
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v860 = v268
	v861 = v269
	v862 = v270
	v863 = v271
	v864 = v272
	v865 = v273
	v866 = v274
	v867 = v275
	v868 = v276
	v872 = v318
	v873 = v313
	v876 = v284
	v881 = v313
	v882 = v290
	v883 = v291
	v885 = v293
	goto L38
L44:
	;
	if v325 == int32(0) {
		v1442 = v268
		v1443 = v269
		v1444 = v270
		v1445 = v271
		v1446 = v272
		v1447 = v273
		v1448 = v274
		v1449 = v275
		v1458 = v284
		v1476 = v302
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v329 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v330 <= v329 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v860 = v268
	v861 = v269
	v862 = v270
	v863 = v271
	v864 = v272
	v865 = v273
	v866 = v274
	v867 = v275
	v868 = v276
	v872 = v325
	v873 = int32(0)
	v876 = v284
	v881 = v329
	v882 = v290
	v883 = v291
	v885 = v293
	goto L38
L47:
	;
	goto L48
L48:
	;
	v335 = int32(0)
	v350 = v335
	v351 = v335
	v354 = int32(-1)
	v358 = v329
	goto L49
L49:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v380 = v354 + int32(1)
	v382 = v380 << (uint(int32(2)) % 32)
	v383 = v290 + v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v384 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v860 = v268
	v861 = v269
	v862 = v270
	v863 = v271
	v864 = v272
	v865 = v273
	v866 = v274
	v867 = v275
	v868 = v276
	v872 = v325
	v873 = v828
	v876 = v284
	v881 = v836
	v882 = v290
	v883 = v291
	v885 = v293
	goto L38
L51:
	;
	v857 = v351 + int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v857 < v858 {
		v350 = v828
		v351 = v857
		v354 = v380
		v358 = v836
		goto L49
	} else {
		goto L125
	}
L52:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382+v291)))
	if v388 == int32(0) {
		v828 = v350
		v836 = v358
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v325)+20))
	v392 = int32(0)
	if v384 == v392 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L54
L56:
	;
	if v445 == int32(0) {
		v828 = v350
		v836 = v358
		goto L51
	} else {
		goto L70
	}
L57:
	;
	v445 = int32(1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v391 == int32(0) {
		v438 = v392
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v445 = v438
	goto L56
L61:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v402 < v401 {
		v438 = v392
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v404 = int32(1)
	if v401 <= v404 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v407 = v404
	goto L65
L64:
	;
	v407 = v401
	goto L65
L65:
	;
	v408 = int32(8)
	v413 = int32(0)
	goto L66
L66:
	;
	v420 = v413 << (uint(int32(2)) % 32)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v384+v408+v420)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v391+v408+v420)))
	v427 = v422 & (v424 ^ int32(-1))
	v429 = base.B2i32(v427 == int32(0))
	if v427 != 0 {
		v438 = v429
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v438 = v429
	goto L60
L68:
	;
	v431 = v413 + int32(1)
	if v431 != v407 {
		v413 = v431
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v448 = v382 + v291
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v449 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v378+v351<<(uint(int32(2))%32))))
	v799 = F_lappend(m, v350, v798)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L21
	} else {
		goto L121
	}
L72:
	;
	if v350 != 0 {
		goto L117
	} else {
		goto L118
	}
L73:
	;
	v660 = int32(0)
	if v628 == v660 {
		goto L101
	} else {
		goto L102
	}
L74:
	;
	v450 = int32(0)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v450 < v451 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v628 = v618
	goto L73
L77:
	;
	v474 = v450
	goto L80
L78:
	;
	v576 = v449
	goto L79
L79:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v607 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v325)+24))
	if v495 == int32(0) {
		v828 = v350
		v836 = v358
		goto L51
	} else {
		goto L82
	}
L81:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v576 = v565
	goto L79
L82:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v498 <= int32(0) {
		v828 = v350
		v836 = v358
		goto L51
	} else {
		goto L83
	}
L83:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v474<<(uint(int32(2))%32))))
	v516 = int32(0)
	goto L84
L84:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v495)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548+v516<<(uint(int32(2))%32))))
	v553 = F_equal(m, v552, v505)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L21
	} else {
		goto L86
	}
L85:
	;
	v562 = v474 + int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v449)+4))
	if v562 < v563 {
		v474 = v562
		goto L80
	} else {
		goto L91
	}
L86:
	;
	if v553 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v558 = v516 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	if v558 < v559 {
		v516 = v558
		goto L84
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	goto L85
L90:
	;
	v828 = v350
	v836 = v358
	goto L51
L91:
	;
	goto L81
L92:
	;
	if v576 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	if v576 == int32(0) {
		v628 = v607
		goto L73
	} else {
		goto L99
	}
L95:
	;
	v628 = int32(0)
	goto L73
L96:
	;
	goto L97
L97:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v613 == int32(1) {
		goto L72
	} else {
		goto L98
	}
L98:
	;
	v775 = v358
	goto L71
L99:
	;
	v775 = v358
	goto L71
L100:
	;
	if v705 != int32(1) {
		v775 = v358
		goto L71
	} else {
		goto L116
	}
L101:
	;
	v705 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v668 = int32(1)
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v669 <= v668 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v672 = v668
	goto L106
L105:
	;
	v672 = v669
	goto L106
L106:
	;
	v676 = int32(0)
	v678 = v660
	goto L107
L107:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v628+int32(8)+v676<<(uint(int32(2))%32))))
	if v685 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v705 = v697
	goto L100
L109:
	;
	goto L108
L110:
	;
	v686 = int32(2)
	if v678 != 0 {
		v697 = v686
		goto L109
	} else {
		goto L113
	}
L111:
	;
	v692 = v678
	goto L112
L112:
	;
	v694 = v676 + int32(1)
	if v694 != v672 {
		v676 = v694
		v678 = v692
		goto L107
	} else {
		goto L115
	}
L113:
	;
	v687 = int32(1)
	if base.Ui32(v687) < base.Ui32(base.I32_popcnt(v685)) {
		v697 = v686
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v692 = v687
	goto L112
L115:
	;
	v697 = v692
	goto L109
L116:
	;
	goto L72
L117:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	v751 = v749
	goto L119
L118:
	;
	v751 = int32(0)
	goto L119
L119:
	;
	v752 = F_bms_add_member(m, v358, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	v775 = v752
	goto L71
L121:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v802 = F_bms_add_member(m, v801, v380)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v802
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	F_bms_free(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = int32(0)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	F_list_free(m, v810)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = int32(0)
	v828 = v799
	v836 = v775
	goto L51
L125:
	;
	goto L50
L126:
	;
	v904 = F_clauselist_selectivity_ext(m, v860, v873, v862, v863, v864, int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L21
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v876)+44)) = int32(0)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v872)+4))
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+20)))
	v1073 = F_statext_mcv_load(m, v1071, v1072)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L21
	} else {
		goto L150
	}
L129:
	;
	v907 = v876 + int32(32)
	v909 = v876 + int32(24)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v872)+4))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v860)+44))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v865)+76))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v912+v913<<(uint(int32(2))%32))))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917)+20)))
	v919 = F_statext_mcv_load(m, v911, v918)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L21
	} else {
		goto L131
	}
L130:
	;
	v1042 = *(*float64)(unsafe.Add(mBase, uint32(v876)+32))
	v1043 = *(*float64)(unsafe.Add(mBase, uint32(v876)+24))
	v1044 = float64(0)
	v1046 = base.F64_sub(v904, v1042)
	if base.F64_lt(v1046, v1044) != 0 {
		v1054 = v1044
		goto L141
	} else {
		goto L142
	}
L131:
	;
	v921 = int32(0)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v872)+20))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v872)+24))
	v925 = F_mcv_get_match_bitmap(m, v873, v922, v923, v919, v921)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	v927 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v907))) = v927
	*(*int64)(unsafe.Add(mBase, uint32(v909))) = v927
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v919)+8))
	if v932 == int32(0) {
		v1041 = float64(0)
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v948 = v921
	v971 = float64(0)
	goto L134
L134:
	;
	v980 = v919 + int32(48) + v948*int32(24)
	v981 = *(*float64)(unsafe.Add(mBase, uint32(v980)))
	v982 = *(*float64)(unsafe.Add(mBase, uint32(v909)))
	*(*float64)(unsafe.Add(mBase, uint32(v909))) = base.F64_add(v981, v982)
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948+v925))))
	if v986 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v1041 = v995
	goto L130
L136:
	;
	v989 = *(*float64)(unsafe.Add(mBase, uint32(v980)+8))
	v990 = *(*float64)(unsafe.Add(mBase, uint32(v907)))
	*(*float64)(unsafe.Add(mBase, uint32(v907))) = base.F64_add(v989, v990)
	v993 = *(*float64)(unsafe.Add(mBase, uint32(v980)))
	v995 = base.F64_add(v971, v993)
	goto L138
L137:
	;
	v995 = v971
	goto L138
L138:
	;
	v997 = v948 + int32(1)
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v919)+8))
	if base.Ui32(v997) < base.Ui32(v998) {
		v948 = v997
		v971 = v995
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	v268 = v860
	v269 = v861
	v270 = v862
	v271 = v863
	v272 = v864
	v273 = v865
	v274 = v866
	v275 = v867
	v276 = v868
	v284 = v876
	v290 = v882
	v291 = v883
	v293 = v885
	v302 = base.F64_mul(v302, v1067)
	goto L36
L141:
	;
	v1056 = base.F64_sub(float64(1), v1043)
	if base.F64_lt(v1056, v1054) != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	if base.F64_gt(v1046, float64(1)) == int32(0) {
		v1054 = v1046
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v1054 = float64(1)
	goto L141
L144:
	;
	goto L140
L145:
	;
	v1058 = v1056
	goto L147
L146:
	;
	v1058 = v1054
	goto L147
L147:
	;
	v1059 = base.F64_add(v1041, v1058)
	if base.F64_lt(v1059, float64(0)) != 0 {
		v1067 = v1044
		goto L144
	} else {
		goto L148
	}
L148:
	;
	if base.F64_gt(v1059, float64(1)) == int32(0) {
		v1067 = v1059
		goto L144
	} else {
		goto L149
	}
L149:
	;
	v1067 = float64(1)
	goto L144
L150:
	;
	if v873 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v268 = v860
	v269 = v861
	v270 = v862
	v271 = v863
	v272 = v864
	v273 = v865
	v274 = v866
	v275 = v867
	v276 = v868
	v284 = v876
	v290 = v882
	v291 = v883
	v293 = v885
	v302 = base.F64_sub(base.F64_add(v302, v1431), base.F64_mul(v302, v1431))
	goto L36
L152:
	;
	v1094 = v1075
	v1118 = v1076
	v1120 = v1076
	goto L157
L153:
	;
	v1075 = int32(0)
	v1076 = float64(0)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v1075 < v1078 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v1431 = float64(0)
	goto L151
L156:
	;
	goto L155
L157:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v873)+12))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1127+v1094<<(uint(int32(2))%32))))
	v1133 = F_clause_selectivity_ext(m, v860, v1131, v862, v863, v864, int32(0))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L21
	} else {
		goto L160
	}
L158:
	;
	v1431 = v1393
	goto L151
L159:
	;
	v1146 = float64(0)
	v1147 = m.G0
	v1149 = v1147 - int32(16)
	m.G0 = v1149
	v1152 = v876 + int32(44)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	if v1153 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v1135 = base.F64_mul(v1120, v1133)
	v1137 = base.F64_add(v1120, base.F64_sub(v1133, v1135))
	if base.F64_lt(v1137, float64(0)) != 0 {
		v1145 = float64(0)
		goto L159
	} else {
		goto L161
	}
L161:
	;
	if base.F64_gt(v1137, float64(1)) == int32(0) {
		v1145 = v1137
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v1145 = float64(1)
	goto L159
L163:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	v1158 = F_palloc0_mul(m, int32(1), v1157)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L21
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v1162 = v876 + int32(32)
	v1164 = v876 + int32(24)
	v1166 = v876 + int32(16)
	v1167 = int32(8)
	v1168 = v876 + v1167
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+8)) = v1131
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+12)) = v1131
	v1175 = F_list_make1_impl(m, int32(1), v1149+v1167)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L21
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1152))) = v1158
	goto L165
L167:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v872)+20))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v872)+24))
	v1180 = F_mcv_get_match_bitmap(m, v1175, v1177, v1178, v1073, int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L21
	} else {
		goto L168
	}
L168:
	;
	v1182 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1162))) = v1182
	*(*int64)(unsafe.Add(mBase, uint32(v1164))) = v1182
	*(*int64)(unsafe.Add(mBase, uint32(v1166))) = v1182
	*(*int64)(unsafe.Add(mBase, uint32(v1168))) = v1182
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	if v1190 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1204 = int32(0)
	v1230 = v1146
	goto L172
L170:
	;
	v1318 = v1146
	goto L171
L171:
	;
	F_pfree(m, v1180)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L21
	} else {
		goto L181
	}
L172:
	;
	v1236 = v1073 + int32(48) + v1204*int32(24)
	v1237 = *(*float64)(unsafe.Add(mBase, uint32(v1236)))
	v1238 = *(*float64)(unsafe.Add(mBase, uint32(v1168)))
	*(*float64)(unsafe.Add(mBase, uint32(v1168))) = base.F64_add(v1237, v1238)
	v1241 = v1204 + v1180
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241))))
	if v1242 != int32(1) {
		v1264 = v1230
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1318 = v1264
	goto L171
L174:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1268 = v1267 + v1204
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268))))
	if v1269 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v1245 = *(*float64)(unsafe.Add(mBase, uint32(v1236)))
	v1246 = *(*float64)(unsafe.Add(mBase, uint32(v1236)+8))
	v1247 = *(*float64)(unsafe.Add(mBase, uint32(v1162)))
	*(*float64)(unsafe.Add(mBase, uint32(v1162))) = base.F64_add(v1246, v1247)
	v1250 = base.F64_add(v1230, v1245)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1251+v1204))))
	if v1253 != int32(1) {
		v1264 = v1250
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v1256 = *(*float64)(unsafe.Add(mBase, uint32(v1236)))
	v1257 = *(*float64)(unsafe.Add(mBase, uint32(v1164)))
	*(*float64)(unsafe.Add(mBase, uint32(v1164))) = base.F64_add(v1256, v1257)
	v1260 = *(*float64)(unsafe.Add(mBase, uint32(v1236)+8))
	v1261 = *(*float64)(unsafe.Add(mBase, uint32(v1166)))
	*(*float64)(unsafe.Add(mBase, uint32(v1166))) = base.F64_add(v1260, v1261)
	v1264 = v1250
	goto L174
L177:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241))))
	v1273 = v1272
	goto L179
L178:
	;
	v1273 = int32(1)
	goto L179
L179:
	;
	v1274 = int32(1)
	v1275 = v1273 & v1274
	*(*uint8)(unsafe.Add(mBase, uint32(v1268))) = uint8(v1275)
	v1278 = v1204 + v1274
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1073)+8))
	if base.Ui32(v1278) < base.Ui32(v1279) {
		v1204 = v1278
		v1230 = v1264
		goto L172
	} else {
		goto L180
	}
L180:
	;
	goto L173
L181:
	;
	m.G0 = v1149 + int32(16)
	v1328 = F_bms_is_member(m, v1094, v881)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L21
	} else {
		goto L183
	}
L182:
	;
	v1395 = v1094 + int32(1)
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v873)+4))
	if v1395 < v1396 {
		v1094 = v1395
		v1118 = v1393
		v1120 = v1145
		goto L157
	} else {
		goto L209
	}
L183:
	;
	if v1328 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1356 = v1133
	goto L186
L185:
	;
	v1330 = *(*float64)(unsafe.Add(mBase, uint32(v876)+32))
	v1331 = *(*float64)(unsafe.Add(mBase, uint32(v876)+8))
	v1332 = float64(0)
	v1334 = base.F64_sub(v1133, v1330)
	if base.F64_lt(v1334, v1332) != 0 {
		v1342 = v1332
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1357 = *(*float64)(unsafe.Add(mBase, uint32(v876)+24))
	v1358 = *(*float64)(unsafe.Add(mBase, uint32(v876)+16))
	v1359 = *(*float64)(unsafe.Add(mBase, uint32(v876)+8))
	v1360 = float64(0)
	v1362 = base.F64_sub(v1135, v1358)
	if base.F64_lt(v1362, v1360) != 0 {
		v1370 = v1360
		goto L198
	} else {
		goto L199
	}
L187:
	;
	v1356 = v1355
	goto L186
L188:
	;
	v1344 = base.F64_sub(float64(1), v1331)
	if base.F64_lt(v1344, v1342) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if base.F64_gt(v1334, float64(1)) == int32(0) {
		v1342 = v1334
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1342 = float64(1)
	goto L188
L191:
	;
	goto L187
L192:
	;
	v1346 = v1344
	goto L194
L193:
	;
	v1346 = v1342
	goto L194
L194:
	;
	v1347 = base.F64_add(v1318, v1346)
	if base.F64_lt(v1347, float64(0)) != 0 {
		v1355 = v1332
		goto L191
	} else {
		goto L195
	}
L195:
	;
	if base.F64_gt(v1347, float64(1)) == int32(0) {
		v1355 = v1347
		goto L191
	} else {
		goto L196
	}
L196:
	;
	v1355 = float64(1)
	goto L191
L197:
	;
	v1385 = base.F64_add(v1118, base.F64_sub(v1356, v1383))
	if base.F64_lt(v1385, float64(0)) != 0 {
		v1393 = float64(0)
		goto L182
	} else {
		goto L207
	}
L198:
	;
	v1372 = base.F64_sub(float64(1), v1359)
	if base.F64_lt(v1372, v1370) != 0 {
		goto L202
	} else {
		goto L203
	}
L199:
	;
	if base.F64_gt(v1362, float64(1)) == int32(0) {
		v1370 = v1362
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1370 = float64(1)
	goto L198
L201:
	;
	goto L197
L202:
	;
	v1374 = v1372
	goto L204
L203:
	;
	v1374 = v1370
	goto L204
L204:
	;
	v1375 = base.F64_add(v1357, v1374)
	if base.F64_lt(v1375, float64(0)) != 0 {
		v1383 = v1360
		goto L201
	} else {
		goto L205
	}
L205:
	;
	if base.F64_gt(v1375, float64(1)) == int32(0) {
		v1383 = v1375
		goto L201
	} else {
		goto L206
	}
L206:
	;
	v1383 = float64(1)
	goto L201
L207:
	;
	if base.F64_gt(v1385, float64(1)) == int32(0) {
		v1393 = v1385
		goto L182
	} else {
		goto L208
	}
L208:
	;
	v1393 = float64(1)
	goto L182
L209:
	;
	goto L158
L210:
	;
	v1485 = int32(0)
	v1492 = m.G0
	v1494 = v1492 - int32(16)
	m.G0 = v1494
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+44))
	if v1496 != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	v4975 = v1458
	v4993 = v1476
	goto L212
L212:
	;
	m.G0 = v4975 + int32(48)
	return v4993
L213:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1510)))
	v1512 = float64(1)
	v1513 = int32(0)
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+120))
	if v1515 == v1513 {
		v1660 = v1513
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+76))
	v1510 = v1496 + v1497<<(uint(int32(2))%32)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1501)+52))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+12))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+76))
	v1510 = v1503 + v1504<<(uint(int32(2))%32) - int32(4)
	goto L213
L217:
	;
	if v1660 != 0 {
		goto L228
	} else {
		goto L229
	}
L218:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+4))
	if v1518 <= int32(0) {
		v1586 = v1513
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1660 = v1586
	goto L217
L220:
	;
	v1521 = int32(0)
	if v1521 < v1518 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1524 = v1518
	goto L223
L222:
	;
	v1524 = v1521
	goto L223
L223:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+12))
	v1527 = int32(0)
	goto L224
L224:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1525+v1527<<(uint(int32(2))%32))))
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1571)+16)))
	v1574 = base.B2i32(v1572 == int32(102))
	if v1572 == int32(102) {
		v1586 = v1574
		goto L219
	} else {
		goto L226
	}
L225:
	;
	v1586 = v1574
	goto L219
L226:
	;
	v1576 = v1527 + int32(1)
	if v1576 != v1524 {
		v1527 = v1576
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	if v1443 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v4930 = v1458
	v4932 = v1494
	v4948 = v1476
	v4949 = v1512
	goto L230
L230:
	;
	m.G0 = v4932 + int32(16)
	v4975 = v4930
	v4993 = base.F64_mul(v4948, v4949)
	goto L212
L231:
	;
	v2109 = int32(0)
	if v2083 == v2109 {
		goto L278
	} else {
		goto L279
	}
L232:
	;
	v1667 = F_palloc_mul(m, int32(2), int32(0))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L21
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1674 = v1443 + int32(4)
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+4))
	v1677 = F_palloc_mul(m, int32(2), v1676)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L21
	} else {
		goto L237
	}
L235:
	;
	v1671 = F_palloc_mul(m, int32(4), int32(0))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L21
	} else {
		goto L236
	}
L236:
	;
	v2069 = v1485
	v2075 = v1671
	v2078 = v1667
	v2081 = v1485
	v2083 = v1485
	v2089 = v1485
	v2090 = v1443 + int32(4)
	goto L231
L237:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+4))
	v1681 = F_palloc_mul(m, int32(4), v1680)
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L21
	} else {
		goto L238
	}
L238:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+4))
	if int32(0) < v1683 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1688 = v1485
	v1695 = int32(0)
	goto L242
L240:
	;
	v1958 = v1485
	goto L241
L241:
	;
	v1998 = int32(16)
	v2004 = int32(0)
	v2006 = base.B2i32(v2004 < v1958)
	if v2004 < v1958 {
		goto L265
	} else {
		goto L266
	}
L242:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+12))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1728+v1695<<(uint(int32(2))%32))))
	v1733 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1494)+8)) = v1733
	v1737 = v1677 + v1695<<(uint(int32(1))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1737))) = uint16(v1733)
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1448)))
	v1741 = F_bms_is_member(m, v1695, v1740)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L21
	} else {
		goto L245
	}
L243:
	;
	v1958 = v1913
	goto L241
L244:
	;
	v1954 = v1695 + int32(1)
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1674)))
	if v1954 < v1955 {
		v1688 = v1913
		v1695 = v1954
		goto L242
	} else {
		goto L264
	}
L245:
	;
	if v1741 != 0 {
		v1913 = v1688
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+76))
	v1746 = F_dependency_is_compatible_clause(m, v1732, v1743, v1494+int32(14))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L21
	} else {
		goto L248
	}
L247:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1737))) = uint16(v1870)
	v1913 = v1871
	goto L244
L248:
	;
	if v1746 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1494)+14)))
	v1870 = v1748
	v1871 = v1688
	goto L247
L250:
	;
	goto L251
L251:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+120))
	v1752 = F_dependency_is_compatible_expression(m, v1732, v1749, v1494+int32(8))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L21
	} else {
		goto L252
	}
L252:
	;
	if v1752 == int32(0) {
		v1913 = v1688
		goto L244
	} else {
		goto L253
	}
L253:
	;
	v1756 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1494)+14)) = uint16(v1756)
	if v1688 <= v1756 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1681+v1688<<(uint(int32(2))%32)))) = v1864
	v1870 = v1688 ^ int32(-1)
	v1871 = v1688 + int32(1)
	goto L247
L255:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1494)+8))
	v1762 = int32(0)
	goto L256
L256:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1681+v1762<<(uint(int32(2))%32))))
	v1807 = F_equal(m, v1806, v1761)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L21
	} else {
		goto L258
	}
L257:
	;
	v1814 = int32(_a_F_statext_clauselist_selectivity_0)
	if v1762&v1814 == v1814 {
		goto L254
	} else {
		goto L263
	}
L258:
	;
	if v1807 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1812 = v1762 + int32(1)
	if v1812 != v1688 {
		v1762 = v1812
		goto L256
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	goto L257
L262:
	;
	goto L254
L263:
	;
	v1870 = v1762 ^ int32(-1)
	v1871 = v1688
	goto L247
L264:
	;
	goto L243
L265:
	;
	v2007 = (v1958<<(uint(v1998)%32) + int32(_a_F_statext_clauselist_selectivity_1)) >> (uint(v1998) % 32)
	goto L267
L266:
	;
	v2007 = v2004
	goto L267
L267:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v1674)))
	if v2008 <= int32(0) {
		v2069 = v1958
		v2075 = v1681
		v2078 = v1677
		v2081 = v2007
		v2083 = v1485
		v2089 = v2006
		v2090 = v1674
		goto L231
	} else {
		goto L268
	}
L268:
	;
	v2020 = int32(0)
	v2027 = v1485
	goto L269
L269:
	;
	v2055 = v1677 + v2020<<(uint(int32(1))%32)
	v2056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2055))))
	if v2056 != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v2069 = v1958
	v2075 = v1681
	v2078 = v1677
	v2081 = v2007
	v2083 = v2063
	v2089 = v2006
	v2090 = v1674
	goto L231
L271:
	;
	v2057 = v2056 + v2007
	*(*uint16)(unsafe.Add(mBase, uint32(v2055))) = uint16(v2057)
	v2060 = F_bms_add_member(m, v2027, base.I32_extend16_s(v2057))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L21
	} else {
		goto L274
	}
L272:
	;
	v2063 = v2027
	goto L273
L273:
	;
	v2065 = v2020 + int32(1)
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1674)))
	if v2065 < v2066 {
		v2020 = v2065
		v2027 = v2063
		goto L269
	} else {
		goto L275
	}
L274:
	;
	v2063 = v2060
	goto L273
L275:
	;
	goto L270
L276:
	;
	F_pfree(m, v4878)
	mBase = m.M
	v4913 = m.ExcPending
	if v4913 != 0 {
		goto L21
	} else {
		goto L601
	}
L277:
	;
	if v2154 != int32(2) {
		goto L293
	} else {
		goto L294
	}
L278:
	;
	v2154 = int32(0)
	goto L277
L279:
	;
	goto L280
L280:
	;
	v2117 = int32(1)
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2083)+4))
	if v2118 <= v2117 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v2121 = v2117
	goto L283
L282:
	;
	v2121 = v2118
	goto L283
L283:
	;
	v2125 = int32(0)
	v2127 = v2109
	goto L284
L284:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2083+int32(8)+v2125<<(uint(int32(2))%32))))
	if v2134 != 0 {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v2154 = v2146
	goto L277
L286:
	;
	goto L285
L287:
	;
	v2135 = int32(2)
	if v2127 != 0 {
		v2146 = v2135
		goto L286
	} else {
		goto L290
	}
L288:
	;
	v2141 = v2127
	goto L289
L289:
	;
	v2143 = v2125 + int32(1)
	if v2143 != v2121 {
		v2125 = v2143
		v2127 = v2141
		goto L284
	} else {
		goto L292
	}
L290:
	;
	v2136 = int32(1)
	if base.Ui32(v2136) < base.Ui32(base.I32_popcnt(v2134)) {
		v2146 = v2135
		goto L286
	} else {
		goto L291
	}
L291:
	;
	v2141 = v2136
	goto L289
L292:
	;
	v2146 = v2141
	goto L286
L293:
	;
	F_bms_free(m, v2083)
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L21
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+120))
	if v2160 != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v4878 = v2078
	v4887 = v1458
	v4889 = v1494
	v4905 = v1476
	v4906 = v1512
	goto L276
L297:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2160)+4))
	v2163 = v2161
	goto L299
L298:
	;
	v2163 = int32(0)
	goto L299
L299:
	;
	v2164 = F_palloc_mul(m, int32(4), v2163)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L21
	} else {
		goto L300
	}
L300:
	;
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v1447)+120))
	if v2166 != 0 {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v3238 = F_palloc_mul(m, int32(4), v3168)
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L21
	} else {
		goto L418
	}
L302:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2166)+4))
	if int32(0) < v2167 {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	v3197 = v2075
	v3200 = v2078
	v3205 = v2083
	v3206 = v1458
	v3208 = v1494
	v3214 = v2164
	v3224 = v1476
	v3225 = v1512
	goto L304
L304:
	;
	F_pfree(m, v3214)
	mBase = m.M
	v3232 = m.ExcPending
	if v3232 != 0 {
		goto L21
	} else {
		goto L415
	}
L305:
	;
	v2171 = v2069
	v2172 = v1444
	v2173 = v1445
	v2174 = v1446
	v2176 = v1448
	v2177 = v2075
	v2180 = v2078
	v2182 = v1443
	v2183 = v2081
	v2185 = v2083
	v2186 = v1458
	v2187 = v1485
	v2188 = v1494
	v2189 = v1485
	v2190 = v1485
	v2191 = v2089
	v2192 = v2090
	v2193 = v1442
	v2194 = v2164
	v2195 = v2166
	v2197 = v1511
	v2204 = v1476
	v2205 = v1512
	goto L308
L306:
	;
	v3151 = v1444
	v3152 = v1445
	v3153 = v1446
	v3155 = v1448
	v3156 = v2075
	v3159 = v2078
	v3161 = v1443
	v3164 = v2083
	v3165 = v1458
	v3167 = v1494
	v3168 = v1485
	v3169 = v1485
	v3171 = v2090
	v3172 = v1442
	v3173 = v2164
	v3183 = v1476
	v3184 = v1512
	goto L307
L307:
	;
	if v3169 != 0 {
		goto L301
	} else {
		goto L414
	}
L308:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2195)+12))
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2211+v2187<<(uint(int32(2))%32))))
	v2216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2215)+16)))
	if v2216 != int32(102) {
		v3105 = v2171
		v3106 = v2172
		v3107 = v2173
		v3108 = v2174
		v3110 = v2176
		v3111 = v2177
		v3114 = v2180
		v3116 = v2182
		v3117 = v2183
		v3119 = v2185
		v3120 = v2186
		v3122 = v2188
		v3123 = v2189
		v3124 = v2190
		v3125 = v2191
		v3126 = v2192
		v3127 = v2193
		v3128 = v2194
		v3129 = v2195
		v3131 = v2197
		v3138 = v2204
		v3139 = v2205
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v3151 = v3106
	v3152 = v3107
	v3153 = v3108
	v3155 = v3110
	v3156 = v3111
	v3159 = v3114
	v3161 = v3116
	v3164 = v3119
	v3165 = v3120
	v3167 = v3122
	v3168 = v3123
	v3169 = v3124
	v3171 = v3126
	v3172 = v3127
	v3173 = v3128
	v3183 = v3138
	v3184 = v3139
	goto L307
L310:
	;
	v3146 = v2187 + int32(1)
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v3129)+4))
	if v3146 < v3147 {
		v2171 = v3105
		v2172 = v3106
		v2173 = v3107
		v2174 = v3108
		v2176 = v3110
		v2177 = v3111
		v2180 = v3114
		v2182 = v3116
		v2183 = v3117
		v2185 = v3119
		v2186 = v3120
		v2187 = v3146
		v2188 = v3122
		v2189 = v3123
		v2190 = v3124
		v2191 = v3125
		v2192 = v3126
		v2193 = v3127
		v2194 = v3128
		v2195 = v3129
		v2197 = v3131
		v2204 = v3138
		v2205 = v3139
		goto L308
	} else {
		goto L413
	}
L311:
	;
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2215)+8)))
	v2220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2197)+20)))
	if v2219 != v2220 {
		v3105 = v2171
		v3106 = v2172
		v3107 = v2173
		v3108 = v2174
		v3110 = v2176
		v3111 = v2177
		v3114 = v2180
		v3116 = v2182
		v3117 = v2183
		v3119 = v2185
		v3120 = v2186
		v3122 = v2188
		v3123 = v2189
		v3124 = v2190
		v3125 = v2191
		v3126 = v2192
		v3127 = v2193
		v3128 = v2194
		v3129 = v2195
		v3131 = v2197
		v3138 = v2204
		v3139 = v2205
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v2222 = int32(0)
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+20))
	if v2224 == v2222 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	if int32(0) <= v2281 {
		goto L324
	} else {
		goto L325
	}
L314:
	;
	v2281 = base.I32_ctz(v2267) | v2268<<(uint(int32(5))%32)
	goto L313
L315:
	;
	v2281 = int32(-2)
	goto L313
L316:
	;
	v2232 = int32(0)
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+4))
	if v2235 <= v2232 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v2238 = v2224 + int32(8)
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2238)))
	v2245 = v2242 & int32(-1)
	if v2245 != 0 {
		v2267 = v2245
		v2268 = v2232
		goto L314
	} else {
		goto L318
	}
L318:
	;
	v2246 = int32(1)
	if v2246 == v2235 {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	v2250 = v2246
	goto L320
L320:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2238+v2250<<(uint(int32(2))%32))))
	if v2257 != 0 {
		v2267 = v2257
		v2268 = v2250
		goto L314
	} else {
		goto L322
	}
L321:
	;
	goto L315
L322:
	;
	v2259 = v2250 + int32(1)
	if v2259 != v2235 {
		v2250 = v2259
		goto L320
	} else {
		goto L323
	}
L323:
	;
	goto L321
L324:
	;
	v2292 = v2281
	v2293 = v2222
	goto L327
L325:
	;
	v2402 = v2222
	goto L326
L326:
	;
	if v2191 != 0 {
		goto L345
	} else {
		goto L346
	}
L327:
	;
	if int32(0) < base.I32_extend16_s(v2292) {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v2402 = v2333
	goto L326
L329:
	;
	v2330 = F_bms_is_member(m, base.I32_extend16_s(v2292+v2183), v2185)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L21
	} else {
		goto L332
	}
L330:
	;
	v2333 = v2293
	goto L331
L331:
	;
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+20))
	if v2334 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	v2333 = v2330 + v2293
	goto L331
L333:
	;
	if int32(0) <= v2390 {
		v2292 = v2390
		v2293 = v2333
		goto L327
	} else {
		goto L344
	}
L334:
	;
	v2390 = base.I32_ctz(v2376) | v2377<<(uint(int32(5))%32)
	goto L333
L335:
	;
	v2390 = int32(-2)
	goto L333
L336:
	;
	v2341 = v2292 + int32(1)
	v2343 = int32(base.Ui32(v2341) >> (uint(int32(5)) % 32))
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2334)+4))
	if v2344 <= v2343 {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v2347 = v2334 + int32(8)
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2347+v2343<<(uint(int32(2))%32))))
	v2354 = v2351 & (int32(-1) << (uint(v2341) % 32))
	if v2354 != 0 {
		v2376 = v2354
		v2377 = v2343
		goto L334
	} else {
		goto L338
	}
L338:
	;
	v2356 = v2343 + int32(1)
	if v2356 == v2344 {
		goto L335
	} else {
		goto L339
	}
L339:
	;
	v2359 = v2356
	goto L340
L340:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v2347+v2359<<(uint(int32(2))%32))))
	if v2366 != 0 {
		v2376 = v2366
		v2377 = v2359
		goto L334
	} else {
		goto L342
	}
L341:
	;
	goto L335
L342:
	;
	v2368 = v2359 + int32(1)
	if v2368 != v2344 {
		v2359 = v2368
		goto L340
	} else {
		goto L343
	}
L343:
	;
	goto L341
L344:
	;
	goto L328
L345:
	;
	v2446 = int32(0)
	v2449 = v2222
	goto L348
L346:
	;
	v2598 = v2222
	goto L347
L347:
	;
	if v2402+v2598 < int32(2) {
		v3105 = v2171
		v3106 = v2172
		v3107 = v2173
		v3108 = v2174
		v3110 = v2176
		v3111 = v2177
		v3114 = v2180
		v3116 = v2182
		v3117 = v2183
		v3119 = v2185
		v3120 = v2186
		v3122 = v2188
		v3123 = v2189
		v3124 = v2190
		v3125 = v2191
		v3126 = v2192
		v3127 = v2193
		v3128 = v2194
		v3129 = v2195
		v3131 = v2197
		v3138 = v2204
		v3139 = v2205
		goto L310
	} else {
		goto L358
	}
L348:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+24))
	if v2476 == int32(0) {
		v2554 = v2449
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v2598 = v2554
	goto L347
L350:
	;
	v2582 = v2446 + int32(1)
	if v2582 != v2171 {
		v2446 = v2582
		v2449 = v2554
		goto L348
	} else {
		goto L357
	}
L351:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2476)+4))
	if v2479 <= int32(0) {
		v2554 = v2449
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v2494 = int32(0)
	v2500 = v2449
	goto L353
L353:
	;
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v2476)+12))
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(v2527+v2494<<(uint(int32(2))%32))))
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v2177+v2446<<(uint(int32(2))%32))))
	v2533 = F_equal(m, v2531, v2532)
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L21
	} else {
		goto L355
	}
L354:
	;
	v2554 = v2535
	goto L350
L355:
	;
	v2535 = v2533 + v2500
	v2537 = v2494 + int32(1)
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2476)+4))
	if v2537 < v2538 {
		v2494 = v2537
		v2500 = v2535
		goto L353
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	goto L349
L358:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+4))
	v2629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2197)+20)))
	v2630 = m.G0
	v2632 = v2630 - int32(32)
	m.G0 = v2632
	v2637 = F_SearchSysCache2(m, int32(62), base.I64_extend_i32_u(v2628), base.I64_extend_i32_u(v2629))
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L21
	} else {
		goto L361
	}
L359:
	;
	if v2191 != 0 {
		goto L377
	} else {
		goto L378
	}
L360:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L21
	} else {
		goto L373
	}
L361:
	;
	if v2637 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2643 = F_SysCacheGetAttr(m, int32(62), v2637, int32(4), v2632+int32(31))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L21
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2661 = m.ExcPending
	if v2661 != 0 {
		goto L21
	} else {
		goto L370
	}
L365:
	;
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2632)+31)))
	if v2645 == int32(1) {
		goto L360
	} else {
		goto L366
	}
L366:
	;
	v2649 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v2643))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L21
	} else {
		goto L367
	}
L367:
	;
	v2651 = F_statext_dependencies_deserialize(m, v2649)
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L21
	} else {
		goto L368
	}
L368:
	;
	F_ReleaseCatCache(m, v2637)
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L21
	} else {
		goto L369
	}
L369:
	;
	m.G0 = v2632 + int32(32)
	goto L359
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2632))) = v2628
	F_errmsg_internal(m, int32(_a_F_statext_clauselist_selectivity_2), v2632)
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L21
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(_a_F_statext_clauselist_selectivity_3), int32(699), int32(_a_F_statext_clauselist_selectivity_4))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L21
	} else {
		goto L372
	}
L372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2632)+20)) = v2628
	*(*int32)(unsafe.Add(mBase, uint32(v2632)+16)) = int32(102)
	F_errmsg_internal(m, int32(_a_F_statext_clauselist_selectivity_5), v2632+int32(16))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L21
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_statext_clauselist_selectivity_3), int32(706), int32(_a_F_statext_clauselist_selectivity_4))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L21
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	if v3064 == int32(0) {
		v3105 = v2171
		v3106 = v2172
		v3107 = v2173
		v3108 = v2174
		v3110 = v2176
		v3111 = v2177
		v3114 = v2180
		v3116 = v2182
		v3117 = v2183
		v3119 = v2185
		v3120 = v2186
		v3122 = v2188
		v3123 = v2189
		v3124 = v2190
		v3125 = v2191
		v3126 = v2192
		v3127 = v2193
		v3128 = v2194
		v3129 = v2195
		v3131 = v2197
		v3138 = v2204
		v3139 = v2205
		goto L310
	} else {
		goto L412
	}
L377:
	;
	v2690 = *(*int32)(unsafe.Add(mBase, uint32(v2651)+8))
	if v2690 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+24))
	if v2688 != 0 {
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v2651)+8))
	v3064 = v2689
	goto L376
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2651)+8)) = v3022
	v3064 = v3022
	goto L376
L381:
	;
	v3022 = int32(0)
	goto L380
L382:
	;
	goto L383
L383:
	;
	v2695 = v2651 + int32(12)
	v2696 = int32(0)
	v2707 = v2696
	v2709 = v2696
	goto L384
L384:
	;
	v2741 = v2695 + v2707<<(uint(int32(2))%32)
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v2741)))
	v2743 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2742)+8)))
	if int32(0) < v2743 {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	v3022 = v2977
	goto L380
L386:
	;
	v3008 = v2707 + int32(1)
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v2651)+8))
	if base.Ui32(v3008) < base.Ui32(v3009) {
		v2707 = v3008
		v2709 = v2977
		goto L384
	} else {
		goto L411
	}
L387:
	;
	v2763 = int32(0)
	goto L390
L388:
	;
	goto L389
L389:
	;
	if v2707 != v2709 {
		goto L408
	} else {
		goto L409
	}
L390:
	;
	v2792 = v2742 + int32(10) + v2763<<(uint(int32(1))%32)
	v2793 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2792))))
	if int32(0) < v2793 {
		goto L393
	} else {
		goto L394
	}
L391:
	;
	goto L389
L392:
	;
	v2914 = v2763 + int32(1)
	v2915 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2742)+8)))
	if v2914 < v2915 {
		v2763 = v2914
		goto L390
	} else {
		goto L407
	}
L393:
	;
	v2796 = v2793 + v2183
	*(*uint16)(unsafe.Add(mBase, uint32(v2792))) = uint16(v2796)
	v2799 = F_bms_is_member(m, base.I32_extend16_s(v2796), v2185)
	mBase = m.M
	v2800 = m.ExcPending
	if v2800 != 0 {
		goto L21
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	if v2191 == int32(0) {
		v2977 = v2709
		goto L386
	} else {
		goto L398
	}
L396:
	;
	if v2799 != 0 {
		goto L392
	} else {
		goto L397
	}
L397:
	;
	v2977 = v2709
	goto L386
L398:
	;
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2215)+24))
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(v2803)+12))
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2804+(v2793^int32(-1))<<(uint(int32(2))%32))))
	v2820 = int32(0)
	goto L399
L399:
	;
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2177+v2820<<(uint(int32(2))%32))))
	v2857 = F_equal(m, v2856, v2810)
	mBase = m.M
	v2858 = m.ExcPending
	if v2858 != 0 {
		goto L21
	} else {
		goto L401
	}
L400:
	;
	v2866 = v2183 + (v2820 ^ int32(-1))
	if v2866&int32(_a_F_statext_clauselist_selectivity_0) == int32(0) {
		v2977 = v2709
		goto L386
	} else {
		goto L406
	}
L401:
	;
	if v2857 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v2862 = v2820 + int32(1)
	if v2862 != v2171 {
		v2820 = v2862
		goto L399
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	goto L400
L405:
	;
	v2977 = v2709
	goto L386
L406:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2792))) = uint16(v2866)
	goto L392
L407:
	;
	goto L391
L408:
	;
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v2741)))
	*(*int32)(unsafe.Add(mBase, uint32(v2695+v2709<<(uint(int32(2))%32)))) = v2962
	goto L410
L409:
	;
	goto L410
L410:
	;
	v2977 = v2709 + int32(1)
	goto L386
L411:
	;
	goto L385
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2194+v2190<<(uint(int32(2))%32)))) = v2651
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(v2651)+8))
	v3105 = v2171
	v3106 = v2172
	v3107 = v2173
	v3108 = v2174
	v3110 = v2176
	v3111 = v2177
	v3114 = v2180
	v3116 = v2182
	v3117 = v2183
	v3119 = v2185
	v3120 = v2186
	v3122 = v2188
	v3123 = v3102 + v2189
	v3124 = v2190 + int32(1)
	v3125 = v2191
	v3126 = v2192
	v3127 = v2193
	v3128 = v2194
	v3129 = v2195
	v3131 = v2197
	v3138 = v2204
	v3139 = v2205
	goto L310
L413:
	;
	goto L309
L414:
	;
	v3197 = v3156
	v3200 = v3159
	v3205 = v3164
	v3206 = v3165
	v3208 = v3167
	v3214 = v3173
	v3224 = v3183
	v3225 = v3184
	goto L304
L415:
	;
	F_bms_free(m, v3205)
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		goto L21
	} else {
		goto L416
	}
L416:
	;
	F_pfree(m, v3200)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L21
	} else {
		goto L417
	}
L417:
	;
	v4878 = v3197
	v4887 = v3206
	v4889 = v3208
	v4905 = v3224
	v4906 = v3225
	goto L276
L418:
	;
	v3241 = int64(0)
	if v3164 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	if int32(0) < v3169 {
		goto L434
	} else {
		goto L435
	}
L420:
	;
	v3285 = int32(0)
	goto L419
L421:
	;
	goto L422
L422:
	;
	v3246 = v3164 + int32(8)
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3164)+4))
	if v3247 == int32(1) {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(v3246)))
	v3285 = base.I32_popcnt(v3250)
	goto L419
L424:
	;
	goto L425
L425:
	;
	v3253 = v3247 << (uint(int32(2)) % 32)
	if v3253 <= int32(7) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v3285 = base.I32_wrap_i64(v3280)
	goto L419
L427:
	;
	if v3253 == int32(0) {
		v3280 = v3241
		goto L426
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v3277 = F_pg_popcount_optimized(m, v3246, v3253)
	mBase = m.M
	v3280 = v3277
	goto L426
L430:
	;
	v3258 = v3253
	v3259 = v3246
	v3260 = v3241
	goto L431
L431:
	;
	v3261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3259)+3)))
	v3262 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3261)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3259)+2)))
	v3264 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3263)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3259)+1)))
	v3266 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3265)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3259))))
	v3268 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3267)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3272 = v3262 + (v3264 + (v3266 + (v3260 + v3268)))
	v3273 = int32(4)
	v3276 = v3258 - v3273
	if v3276 != 0 {
		v3258 = v3276
		v3259 = v3259 + v3273
		v3260 = v3272
		goto L431
	} else {
		goto L433
	}
L432:
	;
	v3280 = v3272
	goto L426
L433:
	;
	goto L432
L434:
	;
	v3294 = v3285
	v3304 = v3164
	v3306 = int32(0)
	goto L437
L435:
	;
	v4837 = v3164
	v4857 = v3184
	goto L436
L436:
	;
	F_pfree(m, v3238)
	mBase = m.M
	v4864 = m.ExcPending
	if v4864 != 0 {
		goto L21
	} else {
		goto L597
	}
L437:
	;
	v3330 = int32(0)
	v3333 = v3330
	v3341 = v3330
	goto L439
L438:
	;
	if v3306 != 0 {
		goto L483
	} else {
		goto L484
	}
L439:
	;
	v3376 = v3173 + v3341<<(uint(int32(2))%32)
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3376)))
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v3377)+8))
	if v3378 != 0 {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	if v3542 != 0 {
		goto L464
	} else {
		goto L465
	}
L441:
	;
	v3380 = v3333
	v3387 = v3377
	v3392 = int32(0)
	goto L444
L442:
	;
	v3542 = v3333
	goto L443
L443:
	;
	v3583 = v3341 + int32(1)
	if v3583 != v3169 {
		v3333 = v3542
		v3341 = v3583
		goto L439
	} else {
		goto L463
	}
L444:
	;
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3387+v3392<<(uint(int32(2))%32))+12))
	v3424 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3423)+8)))
	if v3294 < v3424 {
		v3497 = v3380
		v3504 = v3387
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v3542 = v3497
	goto L443
L446:
	;
	v3538 = v3392 + int32(1)
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v3504)+8))
	if base.Ui32(v3538) < base.Ui32(v3539) {
		v3380 = v3497
		v3387 = v3504
		v3392 = v3538
		goto L444
	} else {
		goto L462
	}
L447:
	;
	if v3380 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	if v3424 <= int32(0) {
		goto L453
	} else {
		goto L454
	}
L449:
	;
	v3428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3380)+8)))
	if v3424 < v3428 {
		v3497 = v3380
		v3504 = v3387
		goto L446
	} else {
		goto L450
	}
L450:
	;
	if v3424 != v3428 {
		goto L448
	} else {
		goto L451
	}
L451:
	;
	v3431 = *(*float64)(unsafe.Add(mBase, uint32(v3380)))
	v3432 = *(*float64)(unsafe.Add(mBase, uint32(v3423)))
	if base.F64_gt(v3431, v3432) != 0 {
		v3497 = v3380
		v3504 = v3387
		goto L446
	} else {
		goto L452
	}
L452:
	;
	goto L448
L453:
	;
	v3497 = v3423
	v3504 = v3387
	goto L446
L454:
	;
	goto L455
L455:
	;
	v3448 = int32(0)
	goto L457
L456:
	;
	v3495 = *(*int32)(unsafe.Add(mBase, uint32(v3376)))
	v3497 = v3493
	v3504 = v3495
	goto L446
L457:
	;
	v3484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3423+int32(10)+v3448<<(uint(int32(1))%32)))))
	v3485 = F_bms_is_member(m, v3484, v3304)
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L21
	} else {
		goto L459
	}
L458:
	;
	v3493 = v3423
	goto L456
L459:
	;
	if v3485 == int32(0) {
		v3493 = v3380
		goto L456
	} else {
		goto L460
	}
L460:
	;
	v3490 = v3448 + int32(1)
	v3491 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3423)+8)))
	if v3490 < v3491 {
		v3448 = v3490
		goto L457
	} else {
		goto L461
	}
L461:
	;
	goto L458
L462:
	;
	goto L445
L463:
	;
	goto L440
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3238+v3306<<(uint(int32(2))%32)))) = v3542
	v3589 = int32(1)
	v3593 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3542)+8)))
	v3597 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3542+int32(8)+v3593<<(uint(v3589)%32)))))
	v3598 = F_bms_del_member(m, v3304, v3597)
	mBase = m.M
	v3599 = m.ExcPending
	if v3599 != 0 {
		goto L21
	} else {
		goto L467
	}
L465:
	;
	goto L466
L466:
	;
	goto L438
L467:
	;
	v3601 = int64(0)
	if v3598 == int32(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v3294 = v3645
	v3304 = v3598
	v3306 = v3306 + v3589
	goto L437
L469:
	;
	v3645 = int32(0)
	goto L468
L470:
	;
	goto L471
L471:
	;
	v3606 = v3598 + int32(8)
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v3598)+4))
	if v3607 == int32(1) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(v3606)))
	v3645 = base.I32_popcnt(v3610)
	goto L468
L473:
	;
	goto L474
L474:
	;
	v3613 = v3607 << (uint(int32(2)) % 32)
	if v3613 <= int32(7) {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	v3645 = base.I32_wrap_i64(v3640)
	goto L468
L476:
	;
	if v3613 == int32(0) {
		v3640 = v3601
		goto L475
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v3637 = F_pg_popcount_optimized(m, v3606, v3613)
	mBase = m.M
	v3640 = v3637
	goto L475
L479:
	;
	v3618 = v3613
	v3619 = v3606
	v3620 = v3601
	goto L480
L480:
	;
	v3621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3619)+3)))
	v3622 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3621)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3619)+2)))
	v3624 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3623)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3619)+1)))
	v3626 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3625)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3619))))
	v3628 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3627)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3632 = v3622 + (v3624 + (v3626 + (v3620 + v3628)))
	v3633 = int32(4)
	v3636 = v3618 - v3633
	if v3636 != 0 {
		v3618 = v3636
		v3619 = v3619 + v3633
		v3620 = v3632
		goto L480
	} else {
		goto L482
	}
L481:
	;
	v3640 = v3632
	goto L475
L482:
	;
	goto L481
L483:
	;
	v3646 = int32(0)
	if v3646 < v3306 {
		goto L486
	} else {
		goto L487
	}
L484:
	;
	v4765 = v3184
	goto L485
L485:
	;
	v4780 = int32(0)
	goto L593
L486:
	;
	v3650 = v3646
	v3655 = int32(0)
	goto L489
L487:
	;
	v3795 = v3646
	goto L488
L488:
	;
	v3838 = int64(0)
	if v3795 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L489:
	;
	v3691 = int32(0)
	v3694 = v3238 + v3655<<(uint(int32(2))%32)
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v3694)))
	v3696 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3695)+8)))
	if v3691 < v3696 {
		goto L491
	} else {
		goto L492
	}
L490:
	;
	v3795 = v3751
	goto L488
L491:
	;
	v3699 = v3650
	v3707 = v3691
	v3713 = v3695
	goto L494
L492:
	;
	v3751 = v3650
	goto L493
L493:
	;
	v3793 = v3655 + int32(1)
	if v3793 != v3306 {
		v3650 = v3751
		v3655 = v3793
		goto L489
	} else {
		goto L498
	}
L494:
	;
	v3743 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3713+v3707<<(uint(int32(1))%32))+10)))
	v3744 = F_bms_add_member(m, v3699, v3743)
	mBase = m.M
	v3745 = m.ExcPending
	if v3745 != 0 {
		goto L21
	} else {
		goto L496
	}
L495:
	;
	v3751 = v3744
	goto L493
L496:
	;
	v3747 = v3707 + int32(1)
	v3748 = *(*int32)(unsafe.Add(mBase, uint32(v3694)))
	v3749 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3748)+8)))
	if v3747 < v3749 {
		v3699 = v3744
		v3707 = v3747
		v3713 = v3748
		goto L494
	} else {
		goto L497
	}
L497:
	;
	goto L495
L498:
	;
	goto L490
L499:
	;
	v3883 = F_palloc_mul(m, int32(8), v3882)
	mBase = m.M
	v3884 = m.ExcPending
	if v3884 != 0 {
		goto L21
	} else {
		goto L514
	}
L500:
	;
	v3882 = int32(0)
	goto L499
L501:
	;
	goto L502
L502:
	;
	v3843 = v3795 + int32(8)
	v3844 = *(*int32)(unsafe.Add(mBase, uint32(v3795)+4))
	if v3844 == int32(1) {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v3843)))
	v3882 = base.I32_popcnt(v3847)
	goto L499
L504:
	;
	goto L505
L505:
	;
	v3850 = v3844 << (uint(int32(2)) % 32)
	if v3850 <= int32(7) {
		goto L507
	} else {
		goto L508
	}
L506:
	;
	v3882 = base.I32_wrap_i64(v3877)
	goto L499
L507:
	;
	if v3850 == int32(0) {
		v3877 = v3838
		goto L506
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v3874 = F_pg_popcount_optimized(m, v3843, v3850)
	mBase = m.M
	v3877 = v3874
	goto L506
L510:
	;
	v3855 = v3850
	v3856 = v3843
	v3857 = v3838
	goto L511
L511:
	;
	v3858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3856)+3)))
	v3859 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3858)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3856)+2)))
	v3861 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3860)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3856)+1)))
	v3863 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3862)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3856))))
	v3865 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3864)+uint32(_c_F_statext_clauselist_selectivity[0]))))
	v3869 = v3859 + (v3861 + (v3863 + (v3857 + v3865)))
	v3870 = int32(4)
	v3873 = v3855 - v3870
	if v3873 != 0 {
		v3855 = v3873
		v3856 = v3856 + v3870
		v3857 = v3869
		goto L511
	} else {
		goto L513
	}
L512:
	;
	v3877 = v3869
	goto L506
L513:
	;
	goto L512
L514:
	;
	if v3795 == int32(0) {
		goto L517
	} else {
		goto L518
	}
L515:
	;
	if int32(0) <= v3941 {
		goto L526
	} else {
		goto L527
	}
L516:
	;
	v3941 = base.I32_ctz(v3927) | v3928<<(uint(int32(5))%32)
	goto L515
L517:
	;
	v3941 = int32(-2)
	goto L515
L518:
	;
	v3892 = int32(0)
	v3895 = *(*int32)(unsafe.Add(mBase, uint32(v3795)+4))
	if v3895 <= v3892 {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v3898 = v3795 + int32(8)
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(v3898)))
	v3905 = v3902 & int32(-1)
	if v3905 != 0 {
		v3927 = v3905
		v3928 = v3892
		goto L516
	} else {
		goto L520
	}
L520:
	;
	v3906 = int32(1)
	if v3906 == v3895 {
		goto L517
	} else {
		goto L521
	}
L521:
	;
	v3910 = v3906
	goto L522
L522:
	;
	v3917 = *(*int32)(unsafe.Add(mBase, uint32(v3898+v3910<<(uint(int32(2))%32))))
	if v3917 != 0 {
		v3927 = v3917
		v3928 = v3910
		goto L516
	} else {
		goto L524
	}
L523:
	;
	goto L517
L524:
	;
	v3919 = v3910 + int32(1)
	if v3919 != v3895 {
		v3910 = v3919
		goto L522
	} else {
		goto L525
	}
L525:
	;
	goto L523
L526:
	;
	v3950 = v3941
	v3958 = int32(0)
	goto L529
L527:
	;
	goto L528
L528:
	;
	v4210 = v3306 - int32(1)
	if int32(0) <= v4210 {
		goto L557
	} else {
		goto L558
	}
L529:
	;
	if v3161 == int32(0) {
		goto L532
	} else {
		goto L533
	}
L530:
	;
	goto L528
L531:
	;
	v4105 = F_clauselist_selectivity_ext(m, v3172, v4071, v3151, v3152, v3153, int32(0))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L21
	} else {
		goto L544
	}
L532:
	;
	v4071 = int32(0)
	goto L531
L533:
	;
	goto L534
L534:
	;
	v3989 = int32(0)
	v3992 = *(*int32)(unsafe.Add(mBase, uint32(v3171)))
	if v3992 <= v3989 {
		v4071 = v3989
		goto L531
	} else {
		goto L535
	}
L535:
	;
	v3996 = v3992
	v4003 = v3989
	v4006 = v3989
	v4009 = int32(-1)
	goto L536
L536:
	;
	v4036 = int32(1)
	v4037 = v4009 + v4036
	v4041 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3159+v4037<<(uint(v4036)%32)))))
	if v4041 == v3950 {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	v4071 = v4056
	goto L531
L538:
	;
	v4043 = *(*int32)(unsafe.Add(mBase, uint32(v3161)+12))
	v4047 = *(*int32)(unsafe.Add(mBase, uint32(v4043+v4003<<(uint(int32(2))%32))))
	v4048 = F_lappend(m, v4006, v4047)
	mBase = m.M
	v4049 = m.ExcPending
	if v4049 != 0 {
		goto L21
	} else {
		goto L541
	}
L539:
	;
	v4055 = v3996
	v4056 = v4006
	goto L540
L540:
	;
	v4058 = v4003 + int32(1)
	if v4058 < v4055 {
		v3996 = v4055
		v4003 = v4058
		v4006 = v4056
		v4009 = v4037
		goto L536
	} else {
		goto L543
	}
L541:
	;
	v4050 = *(*int32)(unsafe.Add(mBase, uint32(v3155)))
	v4051 = F_bms_add_member(m, v4050, v4037)
	mBase = m.M
	v4052 = m.ExcPending
	if v4052 != 0 {
		goto L21
	} else {
		goto L542
	}
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3155))) = v4051
	v4054 = *(*int32)(unsafe.Add(mBase, uint32(v3161)+4))
	v4055 = v4054
	v4056 = v4048
	goto L540
L543:
	;
	goto L537
L544:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v3883+v3958<<(uint(int32(3))%32)))) = v4105
	if v3795 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L545:
	;
	if int32(0) <= v4165 {
		v3950 = v4165
		v3958 = v3958 + int32(1)
		goto L529
	} else {
		goto L556
	}
L546:
	;
	v4165 = base.I32_ctz(v4151) | v4152<<(uint(int32(5))%32)
	goto L545
L547:
	;
	v4165 = int32(-2)
	goto L545
L548:
	;
	v4116 = v3950 + int32(1)
	v4118 = int32(base.Ui32(v4116) >> (uint(int32(5)) % 32))
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v3795)+4))
	if v4119 <= v4118 {
		goto L547
	} else {
		goto L549
	}
L549:
	;
	v4122 = v3795 + int32(8)
	v4126 = *(*int32)(unsafe.Add(mBase, uint32(v4122+v4118<<(uint(int32(2))%32))))
	v4129 = v4126 & (int32(-1) << (uint(v4116) % 32))
	if v4129 != 0 {
		v4151 = v4129
		v4152 = v4118
		goto L546
	} else {
		goto L550
	}
L550:
	;
	v4131 = v4118 + int32(1)
	if v4131 == v4119 {
		goto L547
	} else {
		goto L551
	}
L551:
	;
	v4134 = v4131
	goto L552
L552:
	;
	v4141 = *(*int32)(unsafe.Add(mBase, uint32(v4122+v4134<<(uint(int32(2))%32))))
	if v4141 != 0 {
		v4151 = v4141
		v4152 = v4134
		goto L546
	} else {
		goto L554
	}
L553:
	;
	goto L547
L554:
	;
	v4143 = v4134 + int32(1)
	if v4143 != v4119 {
		v4134 = v4143
		goto L552
	} else {
		goto L555
	}
L555:
	;
	goto L553
L556:
	;
	goto L530
L557:
	;
	v4218 = v4210
	goto L560
L558:
	;
	goto L559
L559:
	;
	if int32(0) < v3882 {
		goto L575
	} else {
		goto L576
	}
L560:
	;
	v4254 = float64(1)
	v4255 = int32(0)
	v4256 = int32(2)
	v4259 = *(*int32)(unsafe.Add(mBase, uint32(v3238+v4218<<(uint(v4256)%32))))
	v4260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4259)+8)))
	if v4256 <= v4260 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	goto L559
L562:
	;
	v4273 = v4255
	v4300 = v4254
	goto L565
L563:
	;
	v4331 = v4255
	v4358 = v4254
	goto L564
L564:
	;
	v4367 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4259+v4331<<(uint(int32(1))%32))+10)))
	v4368 = F_bms_member_index(m, v3795, v4367)
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L21
	} else {
		goto L569
	}
L565:
	;
	v4309 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4259+int32(10)+v4273<<(uint(int32(1))%32)))))
	v4310 = F_bms_member_index(m, v3795, v4309)
	mBase = m.M
	v4311 = m.ExcPending
	if v4311 != 0 {
		goto L21
	} else {
		goto L567
	}
L566:
	;
	v4331 = v4318
	v4358 = v4316
	goto L564
L567:
	;
	v4315 = *(*float64)(unsafe.Add(mBase, uint32(v3883+v4310<<(uint(int32(3))%32))))
	v4316 = base.F64_mul(v4300, v4315)
	v4317 = int32(1)
	v4318 = v4273 + v4317
	v4319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4259)+8)))
	if v4318 < v4319-v4317 {
		v4273 = v4318
		v4300 = v4316
		goto L565
	} else {
		goto L568
	}
L568:
	;
	goto L566
L569:
	;
	v4370 = *(*float64)(unsafe.Add(mBase, uint32(v4259)))
	v4373 = v3883 + v4368<<(uint(int32(3))%32)
	v4374 = *(*float64)(unsafe.Add(mBase, uint32(v4373)))
	if base.F64_le(v4358, v4374) == int32(0) {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v4380 = base.F64_div(base.F64_mul(v4374, v4370), v4358)
	goto L572
L571:
	;
	v4380 = v4370
	goto L572
L572:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v4373))) = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v4370), v4374), v4380)
	if int32(0) < v4218 {
		v4218 = v4218 - int32(1)
		goto L560
	} else {
		goto L573
	}
L573:
	;
	goto L561
L574:
	;
	F_pfree(m, v3883)
	mBase = m.M
	v4727 = m.ExcPending
	if v4727 != 0 {
		goto L21
	} else {
		goto L591
	}
L575:
	;
	v4434 = v3882 & int32(3)
	v4435 = float64(1)
	v4436 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v3882) {
		goto L579
	} else {
		goto L580
	}
L576:
	;
	goto L577
L577:
	;
	v4725 = float64(1)
	goto L574
L578:
	;
	v4635 = float64(0)
	if base.F64_lt(v4627, v4635) != 0 {
		v4725 = v4635
		goto L574
	} else {
		goto L589
	}
L579:
	;
	v4448 = int32(0)
	v4451 = v4436
	v4476 = v4435
	goto L582
L580:
	;
	v4510 = v4436
	v4535 = v4435
	goto L581
L581:
	;
	v4544 = v4436
	v4551 = v4510
	v4576 = v4535
	goto L586
L582:
	;
	v4486 = v3883 + v4451<<(uint(int32(3))%32)
	v4487 = *(*float64)(unsafe.Add(mBase, uint32(v4486)))
	v4489 = *(*float64)(unsafe.Add(mBase, uint32(v4486)+8))
	v4491 = *(*float64)(unsafe.Add(mBase, uint32(v4486)+16))
	v4493 = *(*float64)(unsafe.Add(mBase, uint32(v4486)+24))
	v4494 = base.F64_mul(base.F64_mul(base.F64_mul(base.F64_mul(v4476, v4487), v4489), v4491), v4493)
	v4495 = int32(4)
	v4496 = v4451 + v4495
	v4498 = v4448 + v4495
	if v4498 != v3882&int32(2147483644) {
		v4448 = v4498
		v4451 = v4496
		v4476 = v4494
		goto L582
	} else {
		goto L584
	}
L583:
	;
	if v4434 == int32(0) {
		v4627 = v4494
		goto L578
	} else {
		goto L585
	}
L584:
	;
	goto L583
L585:
	;
	v4510 = v4496
	v4535 = v4494
	goto L581
L586:
	;
	v4587 = *(*float64)(unsafe.Add(mBase, uint32(v3883+v4551<<(uint(int32(3))%32))))
	v4588 = base.F64_mul(v4576, v4587)
	v4589 = int32(1)
	v4592 = v4544 + v4589
	if v4592 != v4434 {
		v4544 = v4592
		v4551 = v4551 + v4589
		v4576 = v4588
		goto L586
	} else {
		goto L588
	}
L587:
	;
	v4627 = v4588
	goto L578
L588:
	;
	goto L587
L589:
	;
	if base.F64_gt(v4627, float64(1)) == int32(0) {
		v4725 = v4627
		goto L574
	} else {
		goto L590
	}
L590:
	;
	goto L577
L591:
	;
	F_bms_free(m, v3795)
	mBase = m.M
	v4729 = m.ExcPending
	if v4729 != 0 {
		goto L21
	} else {
		goto L592
	}
L592:
	;
	v4765 = v4725
	goto L485
L593:
	;
	v4816 = *(*int32)(unsafe.Add(mBase, uint32(v3173+v4780<<(uint(int32(2))%32))))
	F_pfree(m, v4816)
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L21
	} else {
		goto L595
	}
L594:
	;
	v4837 = v3304
	v4857 = v4765
	goto L436
L595:
	;
	v4820 = v4780 + int32(1)
	if v4820 != v3169 {
		v4780 = v4820
		goto L593
	} else {
		goto L596
	}
L596:
	;
	goto L594
L597:
	;
	F_pfree(m, v3173)
	mBase = m.M
	v4866 = m.ExcPending
	if v4866 != 0 {
		goto L21
	} else {
		goto L598
	}
L598:
	;
	F_bms_free(m, v4837)
	mBase = m.M
	v4868 = m.ExcPending
	if v4868 != 0 {
		goto L21
	} else {
		goto L599
	}
L599:
	;
	F_pfree(m, v3159)
	mBase = m.M
	v4870 = m.ExcPending
	if v4870 != 0 {
		goto L21
	} else {
		goto L600
	}
L600:
	;
	v4878 = v3156
	v4887 = v3165
	v4889 = v3167
	v4905 = v3183
	v4906 = v4857
	goto L276
L601:
	;
	v4930 = v4887
	v4932 = v4889
	v4948 = v4905
	v4949 = v4906
	goto L230
}
func F_statext_is_compatible_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v295 int32
	_ = v295
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 != int32(320) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v295
L2:
	;
	if v13 != int32(21) {
		v295 = v6
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)))
	if v50 != 0 {
		v295 = v6
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 != 0 {
		v295 = v6
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 == int32(0) {
		v295 = v19
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= v23 {
		v295 = v19
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v28 = v23
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v28<<(uint(int32(2))%32))))
	v40 = F_statext_is_compatible_clause(m, l0, v39, l2, l3, l4)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v295 = v40
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	if v40 == int32(0) {
		v295 = v40
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v47 = v28 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v47 < v48 {
		v28 = v47
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v54 = int32(0)
	if v51 == v54 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v108 == int32(0) {
		v295 = v6
		goto L1
	} else {
		goto L31
	}
L17:
	;
	v108 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v63 <= v62 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v66 = v62
	goto L22
L21:
	;
	v66 = v63
	goto L22
L22:
	;
	v71 = int32(0)
	v73 = int32(-1)
	goto L24
L23:
	;
	v108 = v100
	goto L16
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v51+int32(8)+v71<<(uint(int32(2))%32))))
	if v81 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(12)))) = v92
	v100 = int32(1)
	goto L23
L26:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v81)))|base.B2i32(int32(0) <= v73) != 0 {
		v100 = v54
		goto L23
	} else {
		goto L29
	}
L27:
	;
	v92 = v73
	goto L28
L28:
	;
	v94 = v71 + int32(1)
	if v94 != v66 {
		v71 = v94
		v73 = v92
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v92 = base.I32_ctz(v81) | v71<<(uint(int32(5))%32)
	goto L28
L30:
	;
	goto L25
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v111 != l2 {
		v295 = v6
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v118 = F_statext_is_compatible_clause_internal(m, v115, l2, l3, l4, v11+int32(11))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	if v118 == int32(0) {
		v295 = v6
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v125 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v128 == v125 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	goto L37
L37:
	;
	v295 = int32(1)
	goto L1
L38:
	;
	if int32(0) <= v185 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v185 = base.I32_ctz(v171) | v172<<(uint(int32(5))%32)
	goto L38
L40:
	;
	v185 = int32(-2)
	goto L38
L41:
	;
	v136 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v139 <= v136 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v142 = v128 + int32(8)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v149 = v146 & int32(-1)
	if v149 != 0 {
		v171 = v149
		v172 = v136
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v150 = int32(1)
	if v150 == v139 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v154 = v150
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v142+v154<<(uint(int32(2))%32))))
	if v161 != 0 {
		v171 = v161
		v172 = v154
		goto L39
	} else {
		goto L47
	}
L46:
	;
	goto L40
L47:
	;
	v163 = v154 + int32(1)
	if v163 != v139 {
		v154 = v163
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v189 = v185
	v194 = v125
	goto L52
L50:
	;
	v266 = v125
	goto L51
L51:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v268 != 0 {
		goto L67
	} else {
		goto L68
	}
L52:
	;
	v198 = F_bms_add_member(m, v194, v189+int32(7))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L11
	} else {
		goto L54
	}
L53:
	;
	v266 = v198
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v201 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if int32(0) <= v257 {
		v189 = v257
		v194 = v198
		goto L52
	} else {
		goto L66
	}
L56:
	;
	v257 = base.I32_ctz(v243) | v244<<(uint(int32(5))%32)
	goto L55
L57:
	;
	v257 = int32(-2)
	goto L55
L58:
	;
	v208 = v189 + int32(1)
	v210 = int32(base.Ui32(v208) >> (uint(int32(5)) % 32))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v211 <= v210 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v214 = v201 + int32(8)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v210<<(uint(int32(2))%32))))
	v221 = v218 & (int32(-1) << (uint(v208) % 32))
	if v221 != 0 {
		v243 = v221
		v244 = v210
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v223 = v210 + int32(1)
	if v223 == v211 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v226 = v223
	goto L62
L62:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v214+v226<<(uint(int32(2))%32))))
	if v233 != 0 {
		v243 = v233
		v244 = v226
		goto L56
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v235 = v226 + int32(1)
	if v235 != v211 {
		v226 = v235
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L53
L67:
	;
	F_pull_varattnos(m, v268, l2, v11+int32(4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L70
	}
L68:
	;
	v274 = v266
	goto L69
L69:
	;
	v275 = F_all_rows_selectable(m, l0, l2, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v274 = v273
	goto L69
L71:
	;
	if v275 == int32(0) {
		v295 = v6
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L37
}
